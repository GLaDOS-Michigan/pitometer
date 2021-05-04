import sys
import os
import csv
import statistics
import numpy as np
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
from scipy import stats
from scipy import signal
import seaborn as sns
import pickle

from conv import *

# Plotting constants
from plot_constants import *

TRAIN_SETS = ["train"]
TEST_SETS = ["test1", "test2", "test3"]


DELAYS = [0, 200, 1_000, 5_000]  # units of microseconds


def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)

    print("Collecting data")
    total_grant_data = dict()
    total_accept_data = dict()
    total_round_data = dict()

    # each toylock data is dict of (size -> delay -> node -> [ durs ])
    for train_set in TRAIN_SETS:
        with open("%s/%s/%s" %(exp_dir, train_set, 'total_grant_data.pickle'), 'rb') as handle:
            train_grant_data = pickle.load(handle)
            merge_maps(total_grant_data, train_grant_data)
        with open("%s/%s/%s" %(exp_dir, train_set, 'total_accept_data.pickle'), 'rb') as handle:
            train_accept_data = pickle.load(handle)
            merge_maps(total_accept_data, train_accept_data)
    for test_set in TEST_SETS:
        with open("%s/%s/%s" %(exp_dir, test_set, 'total_round_data.pickle'), 'rb') as handle:
            test_round_data = pickle.load(handle)
            merge_maps(total_round_data, test_round_data)

    with open("%s/../network/%s" %(exp_dir, 'total_payload16_data.pickle'), 'rb') as handle:
        total_network_data = pickle.load(handle)

    print("\nComputing graphs")

    # Plot Rounds
    plot_convolution("Convolutions", exp_dir, total_grant_data, total_accept_data)
    plot_micro_1_distr_fidelity("Micro-benchmark1", exp_dir, total_round_data, total_grant_data, total_accept_data, total_network_data)
    plot_micro_2_size_fidelity("Micro-benchmark2", exp_dir, total_round_data, total_grant_data, total_accept_data, total_network_data)
    plot_micro_1_distr_fidelity_FINAL("Micro-benchmark1", exp_dir, total_round_data, total_grant_data, total_accept_data, total_network_data)
    print("Done")

def merge_maps(map1, map2):
    """ Copy and merge map2 into map1, each map is of the form 
    (size -> delay -> node -> [ durs ])
    """
    for size in map2.keys():
        if size not in map1:
            map1[size] = dict()
        for delay in map2[size].keys():
            if delay not in map1[size]:
                map1[size][delay] = dict()
            for node in map2[size][delay].keys():
                if node not in map1[size][delay]:
                    map1[size][delay][node] = []
                map1[size][delay][node].extend(map2[size][delay][node])


def plot_convolution(name, root, total_grant_data, total_accept_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
    """
    print("Plotting convolutions")
    for delay in DELAYS:
        with PdfPages("%s/%s_%d.pdf" %(root, name, delay)) as pp:
            x_vals_ring_size = sorted(list(total_grant_data.keys()))
            for ring_size in x_vals_ring_size:
                actual_grant_latencies, actual_accept_latencies = compute_actual_grant_accept(total_grant_data, total_accept_data, delay, ring_size)
                fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
                fig.subplots_adjust(left=0.13, right=0.96, top=0.91, bottom=0.15 )
                # plot_convolution_ax(delay, ring_size, this_ax, "ring size %.d, workload %.1f ms" %(ring_size, delay/1000.0), actual_grant_latencies, actual_accept_latencies)
                plot_convolution_ax(delay, ring_size, this_ax, "Convolution of two CDFs", actual_grant_latencies, actual_accept_latencies)
                pp.savefig(fig)
                plt.close(fig)
        
def plot_convolution_ax(    
    delay,
    ring_size, 
    this_ax, 
    name, 
    actual_grant_latencies, 
    actual_accept_latencies):
    binsize = 1e-4
    grant_pdf, grant_bins = raw_data_to_pdf(actual_grant_latencies, binsize)
    accept_pdf, accept_bins = raw_data_to_pdf(actual_accept_latencies, binsize)
    sum_pdf, newstart, newbinsize = add_histograms(grant_pdf, accept_pdf, min(actual_grant_latencies), min(actual_accept_latencies), binsize, binsize)
    max_data = max(actual_grant_latencies) + max(actual_accept_latencies)
    grant_cdf = pdf_to_cdf(grant_pdf)
    accept_cdf = pdf_to_cdf(accept_pdf)
    sum_cdf = pdf_to_cdf(sum_pdf)
    bincount = int((max_data - newstart)/newbinsize)
    binrange = newbinsize * len(sum_pdf)
    sum_bins = np.linspace(newstart + newbinsize, newstart + binrange, len(sum_pdf))
    
    this_ax.plot(sum_cdf, sum_bins, color='navy', label="convolution")
    this_ax.plot(grant_cdf, grant_bins[:-1], color='forestgreen',label="Grant",linestyle='dotted')
    this_ax.plot(accept_cdf, accept_bins[:-1], color='firebrick',label="Accept",linestyle='dashed')
    this_ax.set_xlabel('cumulative probability')
    this_ax.set_ylabel('latency (ms)')
    this_ax.set_title(name)
    # this_ax.set_ylim(max(0, min(actual_round_latencies)-1), np.percentile(actual_round_latencies, 99.9)+1.5)
    this_ax.set_xlim(0, 1)
    # this_ax.set_yscale("log")
    this_ax.xaxis.set_ticks(np.arange(0, 1.1, 0.1))
    this_ax.legend()


def plot_micro_1_distr_fidelity(name, root, total_round_data, total_grant_data, total_accept_data, total_network_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
        total_network_data -- total_network_data[i][j] is the timings for node i to node j
    """
    print("Plotting graphs for Micro-benchmark 1")
    for delay in DELAYS:
        with PdfPages("%s/%s_%d.pdf" %(root, name, delay)) as pp:
            x_vals_ring_size = sorted(list(total_round_data.keys()))
            for ring_size in x_vals_ring_size:
                participants = sorted(list(total_round_data[ring_size][delay].keys()))
                leader_node = participants[0]
                actual_round_latencies = total_round_data[ring_size][delay][leader_node]
                actual_grant_latencies, actual_accept_latencies = compute_actual_grant_accept(total_grant_data, total_accept_data, delay, ring_size)
                actual_network_latencies = compute_actual_network(participants, total_network_data)
                fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
                fig.subplots_adjust(left=0.215, right=0.95, top=0.88, bottom=0.21 )
                plot_micro_1_distr_fidelity_ax(delay, ring_size, this_ax, "Ring size %.d, workload %.1f ms" %(ring_size, delay/1000.0), actual_round_latencies, actual_grant_latencies, actual_accept_latencies, actual_network_latencies)
                pp.savefig(fig)
                plt.close(fig)


def plot_micro_1_distr_fidelity_FINAL(name, root, total_round_data, total_grant_data, total_accept_data, total_network_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
        total_network_data -- total_network_data[i][j] is the timings for node i to node j
    """
    print("Plotting graphs for Micro-benchmark 1 FINAL")
    ring_size = 8
    with PdfPages("%s/%s_size%d.pdf" %(root, name, ring_size)) as pp:
        fig, axes = plt.subplots(1, len(DELAYS), figsize=(12, 2.5), sharex=False)
        col = 0
        for delay in DELAYS:
            participants = sorted(list(total_round_data[ring_size][delay].keys()))
            leader_node = participants[0]
            actual_round_latencies = total_round_data[ring_size][delay][leader_node]
            actual_grant_latencies, actual_accept_latencies = compute_actual_grant_accept(total_grant_data, total_accept_data, delay, ring_size)
            actual_network_latencies = compute_actual_network(participants, total_network_data)
            fig.subplots_adjust(left=0.05, right=0.96, top=0.88, bottom=0.19 )
            this_ax = axes[col]
            plot_micro_1_distr_fidelity_ax(delay, ring_size, this_ax, "workload %.1f ms" %(delay/1000.0), actual_round_latencies, actual_grant_latencies, actual_accept_latencies, actual_network_latencies)
            col += 1
        pp.savefig(fig)
        plt.close(fig)


def compute_actual_network(participants, total_network_data):
    """Compute the aggregate grant and accept latencies for this delay and ring_size
    Arguments:
        participants {list} -- list of participants in this ring
        total_network_data -- total_network_data[i][j] is the timings for node i to node j
    """
    # participants.sort()
    aggregate_network_latencies = []
    for k in total_network_data.keys():
        for j in total_network_data.keys():
            if k != j: 
                aggregate_network_latencies.extend(total_network_data[j][k])
    return [x/2.0 for x in aggregate_network_latencies]

    # for i in range(len(participants)):
    #     this = participants[i]
    #     succ = participants[(i+1)%len(participants)]
    #     aggregate_network_latencies.extend(total_network_data[this][succ])
    # return [x/2.0 for x in aggregate_network_latencies]


def compute_actual_grant_accept(total_grant_data, total_accept_data, delay, ring_size):
    """Compute the aggregate grant and accept latencies for this delay and ring_size
    Arguments:
        total_grant_data {dict} -- dict of (size -> delay -> node -> [ durs ])
        total_accept_data {dict} -- dict of (size -> delay -> node -> [ durs ])
        delay {int} -- [description]
        ring_size {int} -- [description]
    """
    aggregate_grant_latencies = []
    aggregate_accept_latencies = []
    # TONY: Use size only 2 for training data
    # for r in total_grant_data.keys():
    #     for node in total_grant_data[r][delay].keys():  
    #         aggregate_grant_latencies.extend(total_grant_data[r][delay][node])
    #     for node in total_accept_data[r][delay].keys():
    #         aggregate_accept_latencies.extend(total_accept_data[r][delay][node])
    for node in total_grant_data[2][delay].keys():  
        aggregate_grant_latencies.extend(total_grant_data[2][delay][node])
    for node in total_accept_data[2][delay].keys():
        aggregate_accept_latencies.extend(total_accept_data[2][delay][node])
    return aggregate_grant_latencies, aggregate_accept_latencies


def plot_micro_1_distr_fidelity_ax(
    delay,
    ring_size, 
    this_ax, 
    name, 
    actual_round_latencies, 
    actual_grant_latencies, 
    actual_accept_latencies,
    actual_network_latencies
):
    print("Delay: %d, ring_size: %d" %(delay, ring_size))
    round_cdf, round_bins = raw_data_to_cdf(actual_round_latencies)
    grant_cdf, grant_bins = raw_data_to_cdf(actual_grant_latencies)
    accept_cdf, accept_bins = raw_data_to_cdf(actual_accept_latencies)
    network_cdf, network_bins = raw_data_to_cdf(actual_network_latencies)
    predict_pdf, predict_bins = compute_predicted_toylock_pdf(ring_size, actual_grant_latencies, actual_accept_latencies, actual_network_latencies)
    predict_cdf = pdf_to_cdf(predict_pdf)
    print('name: ' + name)
    print('Pred average '+  str(np.average(predict_bins, weights=predict_pdf)))
    print('Real average ' + str(sum(actual_round_latencies)/ len(actual_round_latencies)))
    print()
    this_ax.plot(predict_cdf, predict_bins, label='predicted performance', color='firebrick', linestyle='dashed')
    this_ax.plot(round_cdf, round_bins, label='actual performance', color='navy')
    # plt.plot(network_cdf, network_bins[:-1], label='network', linestyle='dashed')
    # plt.plot(grant_cdf, grant_bins[:-1], label='grant', linestyle='dashdot')
    # plt.plot(accept_cdf, accept_bins[:-1], label='accept', linestyle='dotted')
    this_ax.set_xlabel('cumulative probability')
    this_ax.set_ylabel('round latency (ms)')
    this_ax.set_title(name)
    # this_ax.set_ylim(0, np.percentile(list(actual_round_latencies) + list(predict_bins), 99.9))
    this_ax.set_ylim(0, np.percentile(list(actual_round_latencies), 100)+30)
    this_ax.set_xlim(0, 1)
    # this_ax.set_yscale("log")
    this_ax.xaxis.set_ticks(np.arange(0, 1.1, 0.2))
    this_ax.legend()

def compute_predicted_toylock_pdf(ring_size, actual_grant_latencies, actual_accept_latencies, actual_network_latencies):
    initial_binsize = 1e-3
    grant_pdf, _ = raw_data_to_pdf(actual_grant_latencies, initial_binsize)
    accept_pdf, _ = raw_data_to_pdf(actual_accept_latencies, initial_binsize)
    network_pdf, _ = raw_data_to_pdf(actual_network_latencies, initial_binsize)
    grant_accept_pdf, newstart, newbinsize = add_histograms(grant_pdf, accept_pdf, min(actual_grant_latencies), min(actual_accept_latencies), initial_binsize, initial_binsize)
    grant_accept_network_pdf, newstart,  newbinsize = add_histograms(grant_accept_pdf, network_pdf, newstart, min(actual_network_latencies), newbinsize, initial_binsize)
    max_data = (max(actual_grant_latencies) + max(actual_accept_latencies) + max(actual_network_latencies))
    # print(len(grant_accept_network_pdf))
    grant_accept_net_start, grant_accept_net_binsize, grant_accept_net_max = newstart, newbinsize, max_data
    total_sum_pdf = grant_accept_network_pdf
    for i in range(ring_size-1):
        total_sum_pdf, newstart,  newbinsize = add_histograms(total_sum_pdf, grant_accept_network_pdf, newstart, grant_accept_net_start, newbinsize, grant_accept_net_binsize)
        max_data = max_data + grant_accept_net_max
    bincount = int((max_data - newstart)/newbinsize)
    # conv_bins = np.linspace(newstart, max_data, len(total_sum_pdf))
    # return total_sum_pdf, conv_bins
    binrange = newbinsize * len(total_sum_pdf)
    # Want the bin to hold the upper bound on each bin. E.g. a bin [0, 1)
    # should be plotted at the point 1
    conv_bins = np.linspace(newstart + newbinsize, newstart + binrange, len(total_sum_pdf))
    return total_sum_pdf, conv_bins


def plot_micro_2_size_fidelity(name, root, total_round_data, total_grant_data, total_accept_data, total_network_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
    """
    print("Plotting graphs for Micro-benchmark 2")
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        for delay in DELAYS:
            x_vals_ring_size = sorted(list(total_round_data.keys()))
            y_vals_observed_max = []
            y_vals_observed_ninety_nine_point_nine_percentiles = []
            y_vals_observed_mean = []
            y_vals_predicted_mean = []
            y_vals_predicted_ninety_nine_point_nine_percentiles = []
            y_vals_predicted_max = []
            for ring_size in x_vals_ring_size:
                participants = sorted(list(total_round_data[ring_size][delay].keys()))
                leader_node = participants[0]

                actual_round_latencies = total_round_data[ring_size][delay][leader_node]
                actual_grant_latencies, actual_accept_latencies = compute_actual_grant_accept(total_grant_data, total_accept_data, delay, ring_size)
                actual_network_latencies = compute_actual_network(participants, total_network_data)
                y_vals_observed_ninety_nine_point_nine_percentiles.append(np.percentile(actual_round_latencies, 99.9))
                y_vals_observed_mean.append(np.mean(actual_round_latencies))
                y_vals_observed_max.append(max(actual_round_latencies))

                predict_pdf, predict_bins = compute_predicted_toylock_pdf(ring_size, actual_grant_latencies, actual_accept_latencies, actual_network_latencies)
                predict_cdf = pdf_to_cdf(predict_pdf)
                y_vals_predicted_mean.append(get_mean(predict_pdf, predict_bins))
                y_vals_predicted_ninety_nine_point_nine_percentiles.append(get_percentile(predict_cdf, predict_bins, 99.9))
                y_vals_predicted_max.append(get_percentile(predict_cdf, predict_bins, 100))

                print("Delay %.1f, size %d" %(delay, ring_size))
                print("\tobserved max: %.3f" %max(actual_round_latencies))
                print("\tpredicted max: %.3f" %get_percentile(predict_cdf, predict_bins, 100))
                print()

            fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
            fig.subplots_adjust(left=0.12, right=0.96, top=0.91, bottom=0.13 )
            plot_micro_2_size_fidelity_ax(this_ax, "delay %.1f" %(delay/1000.0), x_vals_ring_size, 
                y_vals_observed_mean,
                y_vals_observed_ninety_nine_point_nine_percentiles, 
                y_vals_observed_max,
                y_vals_predicted_mean,
                y_vals_predicted_ninety_nine_point_nine_percentiles,
                y_vals_predicted_max)
            pp.savefig(fig)
            plt.close(fig)

            # Also, plot their ratios
            fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
            fig.subplots_adjust(left=0.12, right=0.96, top=0.91, bottom=0.13 )
            plot_micro_2_size_fidelity_ratio_ax(this_ax, "Ratio of predicted latency over observed latency", x_vals_ring_size, 
                y_vals_observed_mean,
                y_vals_observed_ninety_nine_point_nine_percentiles, 
                y_vals_observed_max,
                y_vals_predicted_mean,
                y_vals_predicted_ninety_nine_point_nine_percentiles,
                y_vals_predicted_max)
            pp.savefig(fig)
            plt.close(fig)

def plot_micro_2_size_fidelity_ax(
    this_ax, 
    title, 
    x_vals_ring_size, 
    y_vals_observed_mean,
    y_vals_observed_ninety_nine_point_nine_percentiles,
    y_vals_observed_max,
    y_vals_predicted_mean,
    y_vals_predicted_ninety_nine_point_nine_percentiles,
    y_vals_predicted_max):
    this_ax.set_title(title)
    this_ax.set_xlabel("ring size")
    this_ax.set_xticks(x_vals_ring_size)
    this_ax.set_ylabel("latency (ms)")
    this_ax.plot(x_vals_ring_size, y_vals_observed_mean, label='observed mean', marker='o', color='forestgreen')
    this_ax.plot(x_vals_ring_size, y_vals_predicted_mean, label='predicted mean', marker='v', color='forestgreen', linestyle='dashed')
    this_ax.plot(x_vals_ring_size, y_vals_observed_ninety_nine_point_nine_percentiles, label='observed 99.9 percentile', marker='o', color='navy')
    this_ax.plot(x_vals_ring_size, y_vals_predicted_ninety_nine_point_nine_percentiles, label='predicted 99.9 percentile', marker='v', color='navy', linestyle='dashed')
    this_ax.plot(x_vals_ring_size, y_vals_observed_max, label='observed max', marker='o', color='firebrick')
    this_ax.plot(x_vals_ring_size, y_vals_predicted_max, label='predicted max', marker='v', color='firebrick', linestyle='dashed')
    this_ax.set_yscale("log")
    this_ax.legend()

def plot_micro_2_size_fidelity_ratio_ax(
    this_ax, 
    title, 
    x_vals_ring_size, 
    y_vals_observed_mean,
    y_vals_observed_ninety_nine_point_nine_percentiles,
    y_vals_observed_max,
    y_vals_predicted_mean,
    y_vals_predicted_ninety_nine_point_nine_percentiles,
    y_vals_predicted_max):
    this_ax.set_title(title)
    this_ax.set_xticks(x_vals_ring_size)
    this_ax.set_xlabel("ring size")
    this_ax.set_ylabel("ratio of predicted/observed")

    mean_ratio = [y_vals_predicted_mean[i]/y_vals_observed_mean[i] for i in range(len(y_vals_predicted_mean))]
    ninety_nine_nine_percentile_ratio = [y_vals_predicted_ninety_nine_point_nine_percentiles[i]/y_vals_observed_ninety_nine_point_nine_percentiles[i] for i in range(len(y_vals_predicted_mean))]
    max_ratio =  [y_vals_predicted_max[i]/y_vals_observed_max[i] for i in range(len(y_vals_predicted_mean))]
    this_ax.set_yticks(range(-1,int(max(max_ratio))+2, 1))
    this_ax.plot(x_vals_ring_size, max_ratio, label='max', marker='x', color='firebrick')
    this_ax.plot(x_vals_ring_size, ninety_nine_nine_percentile_ratio, label='99.9 percentile', marker='v', color='navy')
    this_ax.plot(x_vals_ring_size, mean_ratio, label='mean', marker='o', color='forestgreen')
    this_ax.set_ylim(-1, int(max(max_ratio))+2)
    this_ax.legend()


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)
