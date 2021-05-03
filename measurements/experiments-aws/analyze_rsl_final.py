import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import statistics
from scipy import stats
from scipy import signal
from scipy import ndimage
from scipy.interpolate import make_interp_spline, BSpline
from datetime import datetime
import seaborn as sns
import pickle

from conv import *

# Plotting constants
from plot_constants import *

THROW=100  # Ignore the first THROW requests in computing client latencies

TRAIN_SET = "test"
TEST_SET = "test"
F_VALUES = [1]

START = datetime.fromisoformat("2021-05-02 00:00:00")
END = datetime.fromisoformat("2021-05-02 10:00:00")

CLIENT = "us-east-2a"
OH = "us-east-2b"
OR = "us-west-2a"
CA = "us-west-1a"

WORK_METHODS = {0: "LReplicaNextProcessPacket",
           1: "LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a",
           2: "LReplicaNextSpontaneousMaybeEnterPhase2",
           3: "LReplicaNextReadClockMaybeNominateValueAndSend2a",
           4: "LReplicaNextSpontaneousTruncateLogBasedOnCheckpoints",
           5: "LReplicaNextSpontaneousMaybeMakeDecision",
           6: "LReplicaNextSpontaneousMaybeExecute",
           7: "LReplicaNextReadClockCheckForViewTimeout",
           8: "LReplicaNextReadClockCheckForQuorumOfViewSuspicions",
           9: "LReplicaNextReadClockMaybeSendHeartbeat"
}

NOOP_METHODS = {
           0: "LReplicaNextProcessPacket",  # this is the same as in WORK_METHODS
           1: "LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoop",
           2: "LReplicaNextSpontaneousMaybeEnterPhase2Noop",
           3: "LReplicaNextReadClockMaybeNominateValueAndSend2aNoop",
           4: "LReplicaNextSpontaneousTruncateLogBasedOnCheckpointsNoop",
           5: "LReplicaNextSpontaneousMaybeMakeDecisionNoop",
           6: "LReplicaNextSpontaneousMaybeExecuteNoop",
           7: "LReplicaNextReadClockCheckForViewTimeoutNoop",
           8: "LReplicaNextReadClockCheckForQuorumOfViewSuspicionsNoop",
           9: "LReplicaNextReadClockMaybeSendHeartbeatNoop"
}

MAX_QUEUE = "MaxQueueing"


def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)

    """
    total_node_data[f][node_id][method_name] = list of durations
    total_client_data[f][i] = list of client durations for trial i
    total_client_start_end[f][i] = (start, end) time of trial i, defined from start of first request to end of last request
    """

    total_node_data, total_client_data, total_client_start_end = dict(), dict(), dict()
    for f in F_VALUES:
        """
        total_f_node_data[node_id][method_name] = list of durations
        total_f_client_data[i] = list of client durations for trial i
        total_f_client_start_end[i] = (start, end) time of trial i, defined from start of first request to end of last request
        """
        try:
            # Training set in general may not contain data for all f
            with open("%s/%s/total_f%d_node_data.pickle" %(exp_dir, TRAIN_SET, f), 'rb') as handle:
                total_node_data[f] = pickle.load(handle)   
        except FileNotFoundError:
            print("%s/%s/total_f%d_node_data.pickle not found" %(exp_dir, TRAIN_SET, f))
        with open("%s/%s/total_f%d_client_data.pickle" %(exp_dir, TEST_SET, f), 'rb') as handle:
            total_client_data[f] = pickle.load(handle)
        # with open("%s/%s/total_f%d_client_start_end.pickle" %(exp_dir, TEST_SET, f), 'rb') as handle:
        #     total_client_start_end[f] = pickle.load(handle)

    # total_network_data[i][j] is the timings for node i to node j
    with open("%s/../network/%s" %(exp_dir, 'total_payload16_data.pickle'), 'rb') as handle:
        total_network_data = pickle.load(handle)
    # Note that total_client_start_end is currently not used in any computation

    # Plot graphs
    print("\nPlotting graphs for experiment %s" %exp_dir)
    plot_distributions("Paxos Distributions (simple)", exp_dir, total_network_data, total_node_data, total_client_data)
    plot_macro_1_bound_accuracy_simple("Macro-benchmark1_simple", exp_dir, total_network_data, total_node_data, total_client_data, total_client_start_end)
    print("Done")


def plot_client_data(name, root, total_client_data):
    print("Plotting graphs for clients")
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        for f in [2]:
            actual_client_latencies_trials = total_client_data[f]
            for i in range(len(actual_client_latencies_trials)):
                fig, this_ax = plt.subplots(1, 1, figsize=(fig_width+3, fig_height), sharex=False)
                fig.subplots_adjust(left=0.215, right=0.95, top=0.88, bottom=0.21 )
                this_ax.plot(actual_client_latencies_trials[i])
                this_ax.set_title("f=%d, trial %d" %(f, i))
                pp.savefig(fig)
                plt.close(fig)


def plot_distributions(name, root, total_network_data, total_node_data, total_client_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_node_data -- total_node_data[f][node_id][method_name] = list of durations
        total_client_data -- total_client_data[f][i] = list of client durations for trial i
    """
    
    # First attempt to plot client cdfs
    print("Plotting graphs for Paxos distributions")
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        for f in F_VALUES:
            actual_client_latencies = [t for i in total_client_data[f] for t in total_client_data[f][i][THROW:-THROW]]  # simply combine data from all trials
            actual_method_latencies = compute_actual_node(total_node_data[f])   
            fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
            fig.subplots_adjust(left=0.12, right=0.95, top=0.88, bottom=0.21 )
            plot_distributions_ax(f, this_ax, "f = %d" %(f), actual_client_latencies, total_network_data, actual_method_latencies)
            pp.savefig(fig)
            plt.close(fig)

def flatten_map_of_array(l):
    res = []
    for key in l:
        res.extend(l[key])
    return res


def plot_distributions_ax(f, this_ax, name, actual_client_latencies, total_network_data, actual_method_latencies):
    """ 
    Arguments:
        name -- name of this figure
        actual_client_latencies -- list of actual client latencies
        actual_network_latencies -- list of network latencies
        actual_method_latencies -- map of method name to list of latencies
    """
    print("Plotting distribution for f = %d" %(f))
    client_cdf, client_bins = raw_data_to_cdf(actual_client_latencies)
    client_cdf, client_bins = smooth(client_cdf, client_bins)
    sanity_check(actual_client_latencies, total_network_data, actual_method_latencies)
    predict_pdf, predict_bins = compute_predicted_rsl_pdf_simple(f, total_network_data, actual_method_latencies)
    predict_cdf = pdf_to_cdf(predict_pdf)

    plt.plot(predict_cdf, predict_bins, label='predicted performance', color='firebrick', linestyle='dashed')
    plt.plot(client_cdf, client_bins, label='actual performance', color='navy')

    this_ax.set_xlabel('cumulative probability')
    this_ax.set_ylabel('request latency (ms)')
    this_ax.set_title('Latency distributions of an IronRSL instance')
    # this_ax.set_ylim(0, np.percentile(list(actual_client_latencies) + list(predict_bins), 99.9))
    # this_ax.set_ylim(0, np.percentile(list(actual_client_latencies), 100)+30)
    this_ax.set_ylim(0, 120)
    this_ax.set_xlim(0, 1)
    # this_ax.grid()
    # this_ax.set_yscale("log")
    this_ax.xaxis.set_ticks(np.arange(0, 1.1, 0.2))
    this_ax.legend()

def smooth(x_vals, y_vals):
    x_res, y_res = [x_vals[0]], [y_vals[0]]
    curr_y = y_vals[0]
    for i in range(1, len(x_vals)):
        if y_vals[i] != curr_y:
            curr_y = y_vals[i]
            x_res.append(x_vals[i])
            y_res.append(y_vals[i]*1.0)
    return x_res, y_res


def sanity_check(actual_client_latencies, total_network_data, actual_method_latencies):
    # check that network minimums are sane
    print()
    print("SANITY CHECK")
    print("min/max for end-to-end client latency is %.3f/%.3f" %(min(actual_client_latencies), max(actual_client_latencies)))
    print()
    nodes = ["us-east-2a", "us-west-1a", "us-west-2a"]
    for src in nodes:
        for dst in nodes:
            data = [p[0]/2.0 for p in total_network_data[src][dst] if START < p[1] and p[1] < END and not p[2]]
            print("min/max from %s to %s is %.3f/%.3f" %(src, dst, min(data), max(data)))
    print()
    q_data = actual_method_latencies["MaxQueueing"]
    print(len(q_data))
    print(len(actual_client_latencies))
    print("min/max for queueing is %.3f/%.3f" %(min(q_data), max(q_data)))
    print("percentiles for queueing is p50:%.3f, p90:%.3f, p99:%.3f, p99.9:%.3f," %(np.percentile(q_data, 50), np.percentile(q_data, 90), np.percentile(q_data, 99), np.percentile(q_data, 99.9)))
    print()


def compute_predicted_rsl_pdf_simple(f, total_network_data, actual_method_latencies):
    """ return pdf, bins
    Arguments:
        actual_client_latencies -- list of actual client latencies
        total_network_data -- map[src node][target node] -> list of network tuples
        actual_method_latencies -- map of method name to list of latencies

        // ReplyBound = TB2b + MaxQueueTime + ProcessPacketFull(2b) + NoOps(1, 6) + ExecuteFull + D(OH->C)
    """
    initial_binsize = 1e-3
    tb2b_pdf, tb2b_start, tb2b_binsize = compute_TB2b_pdf_simple(f, total_network_data, actual_method_latencies, initial_binsize)
    (processPacketFull_pdf, _), processPacketFull_start = raw_data_to_pdf(actual_method_latencies["LReplicaNextProcessPacket"], initial_binsize), min(actual_method_latencies["LReplicaNextProcessPacket"])
    noop_1_10_pdf, noop_1_10_start, noop_1_10_binsize = convolve_noop_pdf(actual_method_latencies, 1, 10, initial_binsize)
    noop_1_6_pdf, noop_1_6_start, noop_1_6_binsize = convolve_noop_pdf(actual_method_latencies, 1, 6, initial_binsize)
    (executeFull_pdf, _), executeFull_start = raw_data_to_pdf(actual_method_latencies["LReplicaNextSpontaneousMaybeExecute"], initial_binsize), min(actual_method_latencies["LReplicaNextSpontaneousMaybeExecute"])
    net_C_OH_pdf, min_C_OH = network_to_pdf(total_network_data, CLIENT, OH, initial_binsize)

    q_data = actual_method_latencies["MaxQueueing"]
    # q_data.sort()
    # q_data = q_data[len(q_data)//10*9:]
    (maxQ_pdf, _), maxQ_start = raw_data_to_pdf(q_data, initial_binsize), min(q_data)

    sum_pdf, sum_start, sum_binsize = add_histograms(
            tb2b_pdf, maxQ_pdf, 
            tb2b_start, maxQ_start, 
            tb2b_binsize, initial_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
            sum_pdf, processPacketFull_pdf, 
            sum_start, processPacketFull_start, 
            sum_binsize, initial_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
            sum_pdf, noop_1_6_pdf, 
            sum_start, noop_1_6_start, 
            sum_binsize, noop_1_6_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
            sum_pdf, executeFull_pdf, 
            sum_start, executeFull_start, 
            sum_binsize, initial_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
            sum_pdf, net_C_OH_pdf, 
            sum_start, min_C_OH, 
            sum_binsize, initial_binsize)
    binrange = sum_binsize * len(sum_pdf)
    conv_bins = np.linspace(sum_start + sum_binsize, sum_start + binrange, len(sum_pdf))
    return sum_pdf, conv_bins 


def network_to_pdf(total_network_data, src, targ, initial_binsize):
    latencies = [p[0]/2.0 for p in total_network_data[src][targ] if START < p[1] and p[1] < END and not p[2]]
    latencies.extend([p[0]/2.0 for p in total_network_data[targ][src] if START < p[1] and p[1] < END and not p[2]])
    net_pdf, _ = raw_data_to_pdf(latencies, initial_binsize)
    return net_pdf, min(latencies)

def compute_TB2b_pdf_simple(f, total_network_data, actual_method_latencies, initial_binsize):
    """ 
    Arguments:
        total_network_data -- map[src node][target node] -> list of network tuples
        actual_method_latencies -- map of method name to list of latencies
        // NoOps(i, j) = no-op-action i + ... +  no-op-action j-1
    
        // TB2b = TB2a + MaxQueueTime + ProcessPacketFull(2a) + NoOps(0, 10) + D(CA->OH)
        // TB2a = ProcessPacketFull(request) + NoOps(1, 3) + NominateValueFull + NoOps(0, 10) + D(C->OH) + D(OH->CA)
    """
    processPacketFull_pdf, _ = raw_data_to_pdf(actual_method_latencies["LReplicaNextProcessPacket"], initial_binsize)
    noop_1_3_pdf, noop_1_3_start, noop_1_3_binsize = convolve_noop_pdf(actual_method_latencies, 1, 3, initial_binsize)
    nominateValueFull_pdf, _ = raw_data_to_pdf(actual_method_latencies["LReplicaNextReadClockMaybeNominateValueAndSend2a"], initial_binsize)
    noop_0_10_pdf, noop_0_10_start, noop_0_10_binsize = convolve_noop_pdf(actual_method_latencies, 0, 10, initial_binsize)

    net_C_OH_pdf, min_C_OH = network_to_pdf(total_network_data, CLIENT, OH, initial_binsize)
    net_OH_CA_pdf, min_OH_CA = network_to_pdf(total_network_data, OH, CA, initial_binsize)
    net_OH_OR_pdf, min_OH_OR = network_to_pdf(total_network_data, OH, OR, initial_binsize)

    q_data = actual_method_latencies["MaxQueueing"]
    # q_data.sort()
    # q_data = q_data[len(q_data)//10*9:]  
    (maxQ_pdf, _), maxQ_start = raw_data_to_pdf(q_data, initial_binsize), min(q_data)


    # TB2A
    sum_pdf, sum_start, sum_binsize = add_histograms(
        processPacketFull_pdf, noop_1_3_pdf, 
        min(actual_method_latencies["LReplicaNextProcessPacket"]), noop_1_3_start, 
        initial_binsize, noop_1_3_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, nominateValueFull_pdf, 
        sum_start, min(actual_method_latencies["LReplicaNextReadClockMaybeNominateValueAndSend2a"]), 
        sum_binsize, initial_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, noop_0_10_pdf, 
        sum_start, noop_0_10_start, 
        sum_binsize, noop_0_10_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, net_C_OH_pdf, 
        sum_start, min_C_OH, 
        sum_binsize, initial_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, net_OH_OR_pdf, 
        sum_start, min_OH_OR, 
        sum_binsize, initial_binsize)
    
    # TB2B
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, processPacketFull_pdf, 
        sum_start, min(actual_method_latencies["LReplicaNextProcessPacket"]), 
        sum_binsize, initial_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, maxQ_pdf, 
        sum_start, maxQ_start, 
        sum_binsize, initial_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, noop_0_10_pdf, 
        sum_start, noop_0_10_start, 
        sum_binsize, noop_0_10_binsize)
    sum_pdf, sum_start, sum_binsize = add_histograms(
        sum_pdf, net_OH_OR_pdf, 
        sum_start, min_OH_OR, 
        sum_binsize, initial_binsize)

    return sum_pdf, sum_start, sum_binsize

def convolve_noop_pdf(actual_method_latencies, i, j, init_binsize):
    sum_pdf, _ = raw_data_to_pdf(actual_method_latencies[NOOP_METHODS[i]], init_binsize)
    sum_start = min(actual_method_latencies[NOOP_METHODS[i]])
    sum_binsize = init_binsize
    for x in range(i+1, j):
        pdf, _ = raw_data_to_pdf(actual_method_latencies[NOOP_METHODS[x]], init_binsize)
        sum_pdf, sum_start, sum_binsize = add_histograms(sum_pdf, pdf, sum_start, min(actual_method_latencies[NOOP_METHODS[x]]), sum_binsize, init_binsize)
    return sum_pdf, sum_start, sum_binsize



def plot_macro_1_bound_accuracy_simple(name, root, total_network_data, total_node_data, total_client_data, total_client_start_end):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_node_data -- total_node_data[f][node_id][method_name] = list of durations
        total_client_data -- total_client_data[f][i] = list of client durations for trial i
        total_client_start_end -- total_client_start_end[f][i] = (start, end) time of trial i, defined from start of first request to end of last request
    """
    print("Plotting graphs for Micro-benchmark 1 (simplified model)")

    # Compute data points
    x_vals_f = sorted(list(total_client_data.keys()))
    y_vals_actual_max = [get_f_max(total_client_data[f]) for f in x_vals_f]
    y_vals_actual_999 = [get_f_999(total_client_data[f]) for f in x_vals_f]
    y_vals_actual_mean = [get_f_mean(total_client_data[f]) for f in x_vals_f]

    # y_vals_actual_median = [statistics.median(total_client_data[f]) for f in x_vals_f]
    y_vals_actual_median = [statistics.median(flatten_map_of_array(total_client_data[f])) for f in x_vals_f]

    y_vals_actual_errors = [get_f_error(total_client_data[f]) for f in x_vals_f]
    
    print("Computing predictions")
    # TONY: Always use total_node_data[1] to make predictions
    y_vals_predict_max = [predict_f_max_simple(total_network_data, total_node_data[f], f) for f in x_vals_f]
    y_vals_predict_999 = [predict_f_percentile_simple(total_network_data, total_node_data[f], f, 99.9) for f in x_vals_f]
    y_vals_predict_mean = [predict_f_mean_simple(total_network_data, total_node_data[f], f) for f in x_vals_f]

    print("Drawing graphs")
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        # Draw plot
        fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
        fig.subplots_adjust(left=0.15, right=0.95, top=0.9, bottom=0.16 )
        this_ax.set_title("Predictions of IronRSL performance")
        
        this_ax.plot(x_vals_f, y_vals_predict_mean, label='pred. mean', marker='o', color='blue', linestyle='dashed')
        this_ax.plot(x_vals_f, y_vals_actual_mean, label='obs. mean', marker='o', color='blue')
        
        this_ax.plot(x_vals_f, y_vals_predict_999, label='pred. 99.9%',marker='v', color='orange', linestyle='dashed')
        this_ax.plot(x_vals_f, y_vals_actual_999, label='obs. 99.9%', marker='v', color='orange')

        this_ax.plot(x_vals_f, y_vals_predict_max, label='pred. max', marker='x', color='firebrick', linestyle='dashed')
        this_ax.plot(x_vals_f, y_vals_actual_max, label='obs. max', marker='x', color='firebrick')
        
        # this_ax.errorbar(x_vals_f, y_vals_actual_mean, yerr=y_vals_actual_errors, linestyle="None", marker="None", color="black")
        this_ax.legend(loc='upper right', bbox_to_anchor=(0.99, 0.3), ncol=3, columnspacing=0.5, fontsize=6.5)
        this_ax.set_xlabel("f")
        this_ax.set_ylabel("request latency (ms)")
        this_ax.xaxis.set_ticks(x_vals_f)
        this_ax.set_yscale("log")
        this_ax.set_ylim(bottom=0.1)
        # this_ax.set_ylim(bottom=0)
        pp.savefig(fig)
        plt.close(fig)

        print("Predict max   :" + str(y_vals_predict_max) )
        print("Real max      :" + str(y_vals_actual_max) )
        print("Predicted 999 :" + str(y_vals_predict_999) )
        print("Real 999      :" + str(y_vals_actual_999) )
        # print("Real median   :" + str(y_vals_actual_median) )
        print("Predict mean  :" + str(y_vals_predict_mean) )
        print("Real mean     :" + str(y_vals_actual_mean) )
        # print("Real ratio    :" + str([(y_vals_predict_mean[i]-y_vals_actual_mean[i])/y_vals_actual_mean[i] for i in range(len(y_vals_actual_mean))]) )


def predict_f_max_simple(total_network_data, total_f_node_data, f):
    work_actions_times, noop_action_times, max_queue_time = max_action_times(total_f_node_data)
    delay_c_oh = max(compute_actual_network(total_network_data, CLIENT, OH))
    delay_oh_or = max(compute_actual_network(total_network_data, OH, OR))
    return sum_from_action_times_simple(work_actions_times, noop_action_times, max_queue_time, c_oh=delay_c_oh, oh_or=delay_oh_or)

def predict_f_percentile_simple(total_network_data, total_f_node_data, f, percentile):
    work_actions_times, noop_action_times, max_queue_time = percentile_action_times(total_f_node_data, percentile)
    delay_c_oh = np.percentile(compute_actual_network(total_network_data, CLIENT, OH), percentile)
    delay_oh_or = np.percentile(compute_actual_network(total_network_data, OH, OR), percentile)
    return sum_from_action_times_simple(work_actions_times, noop_action_times, max_queue_time, c_oh=delay_c_oh, oh_or=delay_oh_or)

def predict_f_mean_simple(total_network_data, total_f_node_data, f):
    work_actions_times, noop_action_times, max_queue_time = mean_action_times(total_f_node_data)
    delay_c_oh = mean_network_delay(compute_actual_network(total_network_data, CLIENT, OH))
    delay_oh_or = mean_network_delay(compute_actual_network(total_network_data, OH, OR))
    return sum_from_action_times_simple(work_actions_times, noop_action_times, max_queue_time, c_oh=delay_c_oh, oh_or=delay_oh_or)


def mean_network_delay(network_delays):
    return sum(network_delays)/len(network_delays)

def max_action_times(total_f_node_data):
    """
    Returns a dictionary mapping action id to the max completion time
    One dictionary for real and noop
    total_f_node_data[node_id][method_name] = list of durations
    """
    work_res, noop_res = dict(), dict()
    for work_method_id, work_name in WORK_METHODS.items():
        max_time = 0
        for node in total_f_node_data.keys():
            max_time = max(max_time, max([0] + total_f_node_data[node][work_name]))
        work_res[work_method_id] = max_time
    for noop_method_id, noop_name in NOOP_METHODS.items():
        max_time = 0
        for node in total_f_node_data.keys():
            max_time = max(max_time, max([0] + total_f_node_data[node][noop_name]))
        noop_res[noop_method_id] = max_time
    max_queue_res = 0
    for node in total_f_node_data.keys():
            max_queue_res = max(max_time, max([0] + total_f_node_data[node][MAX_QUEUE]))
    return work_res, noop_res, max_queue_res

def percentile_action_times(total_f_node_data, percentile):
    """
    Returns a dictionary mapping action id to the percentile completion time
    total_f_node_data[node_id][method_name] = list of durations
    """
    work_res, noop_res = dict(), dict()
    for work_method_id, work_name in WORK_METHODS.items():
        data = []
        for node in total_f_node_data.keys():
            data.extend(total_f_node_data[node][work_name])
        if len(data) == 0:
            work_res[work_method_id] = 0
        else:
            work_res[work_method_id] = np.percentile(data, percentile)
        
    for noop_method_id, noop_name in NOOP_METHODS.items():
        data = []
        for node in total_f_node_data.keys():
            data.extend(total_f_node_data[node][noop_name])
        if len(data) == 0:
            work_res[work_method_id] = 0
        else:
            noop_res[noop_method_id] = np.percentile(data, percentile)
    data = []
    for node in total_f_node_data.keys():
        data.extend(total_f_node_data[node][MAX_QUEUE])
    max_queue_res = np.percentile(data, percentile)
    return work_res, noop_res, max_queue_res


def mean_action_times(total_f_node_data):
    """
    Returns a dictionary mapping action id to the percentile completion time
    total_f_node_data[node_id][method_name] = list of durations
    """
    work_res, noop_res = dict(), dict()
    for method_id, name in WORK_METHODS.items():
        if method_id == 0:
            aggregate = []
            for node in total_f_node_data.keys():
                aggregate.extend(total_f_node_data[node][name])
            work_res[method_id] = np.mean(aggregate) * 10
        else:
            sum_times = 0
            count = 0
            for node in total_f_node_data.keys():
                sum_times += np.sum(total_f_node_data[node][name])
                count += len(total_f_node_data[node][name])
            if count > 0:
                work_res[method_id] = sum_times/float(count)
            else:
                work_res[method_id] = 0.0
    for method_id, name in NOOP_METHODS.items():
        sum_times = 0
        count = 0
        for node in total_f_node_data.keys():
            sum_times += np.sum(total_f_node_data[node][name])
            count += len(total_f_node_data[node][name])
        noop_res[method_id] = sum_times/float(count)
    sum_times = 0
    count = 0
    for node in total_f_node_data.keys():
        sum_times += np.sum(total_f_node_data[node][MAX_QUEUE])
        count += len(total_f_node_data[node][MAX_QUEUE])
    max_queue_res = sum_times/float(count)
    return work_res, noop_res, max_queue_res


def sum_from_action_times_simple(work_actions_times, noop_action_times, max_queue_time, c_oh=-10, oh_or=-10):
    """
    Computes the predicted RSL using actions_times
    // NoOps(i, j) = no-op-action i + ... +  no-op-action j-1
    
    // ReplyBound = TB2b + MaxQueueTime + NoOps(1, 6) + ExecuteFull + D
    // TB2b = TB2a + MaxQueueTime + ProcessPacketFull(2a) + NoOps(0, 10) + D
    // TB2a = ProcessPacketFull(request) + NoOps(1, 3) + NominateValueFull + NoOps(0, 10) + D + D
    Arguments:
        actions_times -- Map from each action id to the time it uses
    """
    TB2a = work_actions_times[0] + noop_actions_up_to(noop_action_times, 1, 3) + work_actions_times[3] + noop_actions_up_to(noop_action_times, 0, 10) + c_oh * 2
    TB2b = TB2a + max_queue_time + work_actions_times[0] + noop_actions_up_to(noop_action_times, 0, 10) + oh_or
    res = TB2b + max_queue_time + noop_actions_up_to(noop_action_times, 1, 6) + work_actions_times[6] + oh_or
    return res

def noop_actions_up_to(noop_actions_times, i, j):
    res = 0
    for x in range(i, j):
        res += noop_actions_times[x]
    return res

def compute_actual_network(total_network_data, src, targ):
    """
    Arguments:
        total_network_data -- total_network_data[i][j] is the timings for node i to node j
    """
    latencies = [p[0]/2.0 for p in total_network_data[src][targ] if START < p[1] and p[1] < END and not p[2]]
    latencies.extend([p[0]/2.0 for p in total_network_data[targ][src] if START < p[1] and p[1] < END and not p[2]])
    return latencies

def compute_actual_node(total_node_data_f):
    """maps total_node_data to res: method_name -> list of latencies
    Args:
        total_node_data : total_node_data_f[node_id][method_name] = list of durations
    """
    res = dict()
    for node in total_node_data_f:
        for method in total_node_data_f[node]:
            if method not in res:
                res[method] = []
            res[method].extend(total_node_data_f[node][method])
    return res


def get_f_max(total_f_client_data):
    """Get the maximum lantecy observed
    Arguments:
        total_f_client_data -- total_f_client_data[i] = list of client durations for trial i
    """
    res = 0
    for durs in total_f_client_data.values():
        res = max(res, max(durs[THROW:])) # Ignore the first 100 requests
    return res

def get_f_999(total_f_client_data):
    """Get the 99.9% lantecy observed
    Arguments:
        total_f_client_data -- total_f_client_data[i] = list of client durations for trial i
    """
    aggregate = []
    for durs in total_f_client_data.values():
        aggregate.extend(durs[THROW:]) # Ignore the first 100 requests
    return np.percentile(aggregate, 99.9)

def get_f_mean(total_f_client_data):
    """Get the mean lantecy observed
    Arguments:
        total_f_client_data -- total_f_client_data[i] = list of client durations for trial i
    """
    aggregate = []
    for durs in total_f_client_data.values():
        aggregate.extend(durs[THROW:])
    return np.mean(aggregate)

def get_f_error(total_f_client_data):
    """Get the stdev
    Arguments:
        total_f_client_data -- total_f_client_data[i] = list of client durations for trial i
    """
    aggregate = []
    for durs in total_f_client_data.values():
        aggregate.extend(durs[THROW:])
    return statistics.stdev(aggregate)


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)
