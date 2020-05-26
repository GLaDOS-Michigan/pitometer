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


DELAYS = [0, 200, 1_000, 5_000, 25_000]  # units of microseconds


def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)

    # each toylock data is dict of (size -> delay -> node -> [ durs ])
    with open("%s/%s" %(exp_dir, 'total_grant_data.pickle'), 'rb') as handle:
        total_grant_data = pickle.load(handle)
    with open("%s/%s" %(exp_dir, 'total_accept_data.pickle'), 'rb') as handle:
        total_accept_data = pickle.load(handle)
    with open("%s/%s" %(exp_dir, 'total_round_data.pickle'), 'rb') as handle:
        total_round_data = pickle.load(handle)
    # total_network_data[i][j] is the timings for node i to node j
    with open("%s/../../network/%s" %(exp_dir, 'total_payload16_data.pickle'), 'rb') as handle:
        total_network_data = pickle.load(handle)

    print("\nPlotting graphs for experiment %s" %exp_dir)

    # Plot Rounds
    plot_micro_1_distr_fidelity("Micro-benchmark1", exp_dir, total_round_data, total_grant_data, total_accept_data, total_network_data)
    plot_micro_2_size_fidelity("Micro-benchmark2", exp_dir, total_round_data, total_grant_data, total_accept_data, total_network_data)
    print("Done")


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
                fig.subplots_adjust(left=0.17, right=0.95, top=0.9, bottom=0.16 )
                plot_micro_1_distr_fidelity_ax(delay, ring_size, this_ax, "ring size %.d" %(ring_size), actual_round_latencies, actual_grant_latencies, actual_accept_latencies, actual_network_latencies)
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
    for r in total_grant_data.keys():
        for node in total_grant_data[r][delay].keys():
            aggregate_grant_latencies.extend(total_grant_data[r][delay][node])
        for node in total_accept_data[r][delay].keys():
            aggregate_accept_latencies.extend(total_accept_data[r][delay][node])
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
    plt.plot(round_cdf, round_bins[:-1], label='actual round', linewidth=0.7)
    plt.plot(predict_cdf, predict_bins, label='predicted')
    # plt.plot(network_cdf, network_bins[:-1], label='network', linestyle='dashed')
    # plt.plot(grant_cdf, grant_bins[:-1], label='grant', linestyle='dashdot')
    # plt.plot(accept_cdf, accept_bins[:-1], label='accept', linestyle='dotted')
    this_ax.set_xlabel('cumulative probability')
    this_ax.set_ylabel('latency (ms)')
    this_ax.set_title(name)
    # this_ax.set_ylim(min(actual_round_latencies), max(actual_round_latencies)+1)
    this_ax.set_xlim(0, 1)
    this_ax.set_yscale("log")
    this_ax.xaxis.set_ticks(np.arange(0, 1.1, 0.1))
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

            fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
            fig.subplots_adjust(left=0.17, right=0.95, top=0.9, bottom=0.16 )
            plot_micro_2_size_fidelity_ax(this_ax, "delay %.1f" %(delay/1000.0), x_vals_ring_size, 
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
    this_ax.set_ylabel("latency (ms)")
    this_ax.plot(x_vals_ring_size, y_vals_observed_mean, label='observed mean', marker='v', color='green')
    this_ax.plot(x_vals_ring_size, y_vals_predicted_mean, label='predicted mean', marker='v', color='green', linestyle='dashed')
    this_ax.plot(x_vals_ring_size, y_vals_observed_ninety_nine_point_nine_percentiles, label='observed 99.9 percentile', marker='o', color='blue')
    this_ax.plot(x_vals_ring_size, y_vals_predicted_ninety_nine_point_nine_percentiles, label='predicted 99.9 percentile', marker='x', color='blue', linestyle='dashed')
    this_ax.plot(x_vals_ring_size, y_vals_observed_max, label='observed max', marker='o', color='red')
    this_ax.plot(x_vals_ring_size, y_vals_predicted_max, label='predicted max', marker='x', color='red', linestyle='dashed')
    this_ax.legend()


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)
