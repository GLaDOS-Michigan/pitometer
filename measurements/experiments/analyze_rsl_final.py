import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
from scipy import stats
from scipy import signal
import seaborn as sns
import pickle

from conv import *

# Plotting constants
from plot_constants import *

F_VALUES = [1, 2, 3]

TRAIN_SET = "emergency_train"
TEST_SET = "emergency_train"


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
        with open("%s/%s/total_f%d_node_data.pickle" %(exp_dir, TRAIN_SET, f), 'rb') as handle:
            total_node_data[f] = pickle.load(handle)
        with open("%s/%s/total_f%d_client_data.pickle" %(exp_dir, TRAIN_SET, f), 'rb') as handle:
            total_client_data[f] = pickle.load(handle)
        with open("%s/%s/total_f%d_client_start_end.pickle" %(exp_dir, TEST_SET, f), 'rb') as handle:
            total_client_start_end[f] = pickle.load(handle)

    # total_network_data[i][j] is the timings for node i to node j
    with open("%s/../network/run1/%s" %(exp_dir, 'total_payload32_data.pickle'), 'rb') as handle:
        total_network_data = pickle.load(handle)

    # Plot graphs
    print("\nPlotting graphs for experiment %s" %exp_dir)
    plot_macro_1_bound_accuracy("Macro-benchmark1", exp_dir, total_network_data, total_node_data, total_client_data, total_client_start_end)
    print("Done")


def plot_macro_1_bound_accuracy(name, root, total_network_data, total_node_data, total_client_data, total_client_start_end):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_node_data -- total_node_data[f][node_id][method_name] = list of durations
        total_client_data -- total_client_data[f][i] = list of client durations for trial i
        total_client_start_end -- total_client_start_end[f][i] = (start, end) time of trial i, defined from start of first request to end of last request
    """
    print("Plotting graphs for Micro-benchmark 1")

    # Compute data points
    x_vals_f = sorted(list(total_node_data.keys()))
    y_vals_actual_max = [get_f_max(total_client_data[f]) for f in x_vals_f]
    # y_vals_actual_999 = [get_f_999(total_client_data[f]) for f in x_vals_f]
    y_vals_actual_mean = [get_f_mean(total_client_data[f]) for f in x_vals_f]
    
    print("Computing predictions")
    # TONY: Always use total_node_data[1] to make predictions
    y_vals_predict_max = [predict_f_max(total_network_data, total_node_data[f], f) for f in x_vals_f]
    # y_vals_predict_999 = [predict_f_percentile(total_network_data, total_node_data[f], f, 99.9) for f in x_vals_f]
    y_vals_predict_mean = [predict_f_mean(total_network_data, total_node_data[f], f) for f in x_vals_f]

    print("Drawing graphs")
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        # Draw plot
        fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
        fig.subplots_adjust(left=0.17, right=0.95, top=0.9, bottom=0.16 )

        # this_ax.plot(x_vals_f, y_vals_actual_max, label='observed max', marker='o', color='red')
        # this_ax.plot(x_vals_f, y_vals_predict_max, label='predicted max', marker='x', color='red', linestyle='dashed')

        # this_ax.plot(x_vals_f, y_vals_actual_999, label='observed 99.9', marker='o', color='blue')
        # this_ax.plot(x_vals_f, y_vals_predict_999, label='predicted 99.9', marker='x', color='blue', linestyle='dashed')

        this_ax.plot(x_vals_f, y_vals_actual_mean, label='observed mean', marker='o', color='green')
        this_ax.plot(x_vals_f, y_vals_predict_mean, label='predicted mean', marker='x', color='green', linestyle='dashed')

        this_ax.legend()
        this_ax.set_xlabel("f")
        this_ax.set_ylabel("latency (ms)")
        # this_ax.set_yscale("log")
        pp.savefig(fig)
        plt.close(fig)


def predict_f_max(total_network_data, total_f_node_data, f):
    work_actions_times, noop_action_times = max_action_times(total_f_node_data)
    network_delays = compute_actual_network(total_network_data)
    return sum_from_action_times(work_actions_times, noop_action_times, f, max(network_delays))


def predict_f_percentile(total_network_data, total_f_node_data, f, percentile):
    actions_times = percentile_action_times(total_f_node_data, percentile)
    network_delays = compute_actual_network(total_network_data)
    return sum_from_action_times(actions_times, f, np.percentile(network_delays, 99.9))

def predict_f_mean(total_network_data, total_f_node_data, f):
    work_actions_times, noop_action_times = mean_action_times(total_f_node_data)
    network_delays = compute_actual_network(total_network_data)
    return sum_from_action_times(work_actions_times, noop_action_times, f, mean_network_delay(network_delays, f))

def predict_f_mean_bad(total_network_data, total_f_node_data, f):
    actions_times = mean_action_times(total_f_node_data)
    network_delays = compute_actual_network(total_network_data)
    return sum_from_action_times(actions_times, f, np.mean(network_delays))


def mean_network_delay(network_delays, f):
    cdf, bins = raw_data_to_cdf(network_delays)
    total_cdf = cdf
    for q in range(f):
        for i in range(len(cdf)):
            total_cdf[i] = total_cdf[i] * cdf[i]
    mean = 0
    for i in range(len(bins)-1):
        binsize = bins[i+1] - bins[i]
        mean += binsize - cdf[i] * binsize
    return mean

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
    return work_res, noop_res

def percentile_action_times(total_f_node_data, percentile):
    """
    Returns a dictionary mapping action id to the percentile completion time
    total_f_node_data[node_id][method_name] = list of durations
    """
    method_ids = list(range(0, 10))
    res = dict()
    for method_id in method_ids:
        name = METHODS[method_id]
        aggregate_method_times = []
        for node in total_f_node_data.keys():
            aggregate_method_times.extend(total_f_node_data[node][name])
        res[method_id] = np.percentile(aggregate_method_times, percentile)
    return res

def mean_action_times(total_f_node_data):
    """
    Returns a dictionary mapping action id to the percentile completion time
    total_f_node_data[node_id][method_name] = list of durations
    """
    work_res, noop_res = dict(), dict()
    for method_id, name in WORK_METHODS.items():
        # if method_id == 0:
        #     aggregate = []
        #     for node in total_f_node_data.keys():
        #         aggregate.extend(total_f_node_data[node][name])
        #     work_res[method_id] = np.percentile(aggregate, 99.9)
        # else:
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
    print(work_res)
    print(noop_res)
    print()
    return work_res, noop_res

def sum_from_action_times(work_actions_times, noop_action_times, f, delay):
    """
    Computes the predicted RSL using actions_times
    // Bound with full vs no-op versions:
    // NoOps(i, j) = no-op-action i + ... +  no-op-action j-1
    
    // ReplyBound = TB2b + (f + 3) * (ProcessPacketFull(2b) + NoOps(1, 10)) + ProcessPacketFull + NoOps(1, 6) + ExecuteFull

    // TB2b = TB2a + ProcessPacketFull(2a) + NoOps(0, 10) + D
    // TB2a = ProcessPacketFull(request) + NoOps(1, 3) + NominateValueFull + NoOps(0, 10) + D
    Arguments:
        actions_times -- Map from each action id to the time it uses
    """
    TB2a = work_actions_times[0] + noop_actions_up_to(noop_action_times, 1, 3) + work_actions_times[3] + noop_actions_up_to(noop_action_times, 0, 10) + delay
    TB2b = TB2a + work_actions_times[0] + noop_actions_up_to(noop_action_times, 0, 10) + delay
    res = TB2b + (f+3) * (work_actions_times[0] + noop_actions_up_to(noop_action_times, 1, 10)) + work_actions_times[0] + noop_actions_up_to(noop_action_times, 1, 6) + work_actions_times[6]
    return res

def noop_actions_up_to(noop_actions_times, i, j):
    res = 0
    for x in range(i, j):
        res += noop_actions_times[x]
    return res

def compute_actual_network(total_network_data):
    """Compute the aggregate grant and accept latencies for this delay and ring_size
    Arguments:
        total_network_data -- total_network_data[i][j] is the timings for node i to node j
    """
    # participants.sort()
    aggregate_network_latencies = []
    for k in total_network_data.keys():
        for j in total_network_data.keys():
            if k != j: 
                aggregate_network_latencies.extend(total_network_data[j][k])
    return [x/2.0 for x in aggregate_network_latencies]


def get_f_max(total_f_client_data):
    """Get the maximum lantecy observed
    Arguments:
        total_f_client_data -- total_f_client_data[i] = list of client durations for trial i
    """
    res = 0
    for durs in total_f_client_data.values():
        res = max(res, max(durs))
    return res

def get_f_999(total_f_client_data):
    """Get the 99.9% lantecy observed
    Arguments:
        total_f_client_data -- total_f_client_data[i] = list of client durations for trial i
    """
    aggregate = []
    for durs in total_f_client_data.values():
        aggregate.extend(durs)
    return np.percentile(aggregate, 99.9)

def get_f_mean(total_f_client_data):
    """Get the mean lantecy observed
    Arguments:
        total_f_client_data -- total_f_client_data[i] = list of client durations for trial i
    """
    aggregate = []
    for durs in total_f_client_data.values():
        aggregate.extend(durs)
    return np.mean(aggregate)


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)
