import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import statistics
from numpy.lib.function_base import hanning
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

HBPeriod=100
EpochLength=1000


THROW=1  # Ignore the first THROW requests in computing method latencies
CTHROW=1  # Ignore the THROW requests in computing client latencies

TRAIN_SET = "train"
TEST_SET = "test"
CD_VALUES = [0, 500, 1500, 2000, 2500, 3000]

START = datetime.fromisoformat("2021-11-14 00:00:01")
END = datetime.fromisoformat("2021-12-15 04:00:00")

CLIENT = "us-east-2a-C"
L = "us-east-2a-L"
B = "us-east-2b"
C = "us-east-2c"

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
    failure_data = parse_csvs(exp_dir)
    
    with open("%s/../../aux/total_f1_node_data.pickle" %exp_dir, 'rb') as handle:
        node_data = pickle.load(handle)
    
    with open("%s/../../aux/total_payload16_data.pickle" %exp_dir, 'rb') as handle:
        network_data = pickle.load(handle)
  

    print("\nPlotting graphs for experiment %s" %exp_dir)
    plot_graph(exp_dir, failure_data, node_data, network_data)
    print("Done")
    
    
def parse_csvs(exp_dir):
    failure_data = {key: [] for key in CD_VALUES}
    for cd in CD_VALUES:
        csv_file = "%s/cd_%d/cd_log_%d.csv" %(exp_dir, cd, cd)
        print(csv_file)
    
        with open(csv_file, 'r') as f:
            csvreader = csv.reader(f, delimiter=' ',)
            for row in csvreader:
                if len(row) < 4 or 'req0' not in row[0]:
                    continue
                start_time = float(row[1])
                end_time = float(row[2])
                dur = (end_time - start_time)  # duration in milliseconds
                failure_data[cd].append(dur)
    return failure_data 


def plot_graph(root, failure_data, node_data, network_data):   
    pred = compute_predicted(network_data, node_data)
    y_vals_predicted = [pred for x in CD_VALUES]
    y_vals = [ np.mean(failure_data[cd]) for cd in CD_VALUES ]
    errors = [ statistics.stdev(failure_data[cd]) if len(failure_data[cd]) > 1 else 0 for cd in CD_VALUES ]
    
    with PdfPages("%s/failure_graph.pdf" %root) as pp:
        fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
        fig.subplots_adjust(right=0.96, bottom=0.18, left=0.15)
        plt.plot(CD_VALUES, y_vals, label='Observed behavior', color='navy', marker='o')
        plt.plot(CD_VALUES, y_vals_predicted, label='Predicted behavior', color='firebrick', linestyle='dashed')
        # this_ax.errorbar(CD_VALUES, y_vals, yerr=errors, linestyle="None", marker="None", color="black")

        this_ax.set_xlabel('failure offset (μs)')
        this_ax.set_ylabel('request latency (ms)')
        this_ax.set_title('End-to-end behavior with 1 crash failure')
        # this_ax.set_ylim(0, 20)
        # this_ax.set_xlim(0, 1)
        # this_ax.grid()
        # this_ax.set_yscale("log")
        # this_ax.xaxis.set_ticks(np.arange(0, 1.1, 0.2))
        this_ax.legend()
        pp.savefig(fig)
        plt.close(fig)
        
        
def compute_predicted(total_network_data, node_data):
    """ return pdf, bins
    Arguments:
        actual_client_latencies -- list of actual client latencies
        total_network_data -- map[src node][target node] -> list of network tuples
        actual_method_latencies -- map of method name to list of latencies

        ReplyBound = Failover
        + D + ProcPkt + Q + D + Q + noop(0, 1) + Action(2)
        + Mbe2a + D + Proc + Q + D + Q + noop(0, 5) + Action(6)
    """
    work_actions_times, noop_action_times, max_queue_time = mean_action_times(node_data)
    # delay_c_oh = mean_network_delay(compute_actual_network(total_network_data, CLIENT, L))
    delay_oh_or = mean_network_delay(compute_actual_network(total_network_data, L, C))
    
    r1 = delay_oh_or + work_actions_times[0] + max_queue_time + delay_oh_or + max_queue_time + noop_actions_up_to(noop_action_times, 0, 2) + work_actions_times[2]
    r2 = work_actions_times[3] + delay_oh_or + work_actions_times[0] + max_queue_time + delay_oh_or + max_queue_time + noop_actions_up_to(noop_action_times, 0, 6) + work_actions_times[6]
    res = compute_predicted_Failover(total_network_data, node_data) + r1 + r2
    
    print("Total: " + str(res))
    return res
    

def compute_predicted_Failover(total_network_data, node_data):
    """ 
    Arguments:
        total_network_data -- map[src node][target node] -> list of network tuples
        actual_method_latencies -- map of method name to list of latencies
        // NoOps(i, j) = no-op-action i + ... +  no-op-action j-1
        
        TBNewView() + StepToTimeDelta(RslStep(9)) + StepToTimeDelta(RslStep(0)) + Timeout() + StepToTimeDelta(RslStep(1))
    """    
    work_actions_times, noop_action_times, max_queue_time = mean_action_times(node_data)
    # delay_c_oh = mean_network_delay(compute_actual_network(total_network_data, CLIENT, L))
    # delay_oh_or = mean_network_delay(compute_actual_network(total_network_data, L, C))
    
    res = compute_predicted_Failover_TBNewView(total_network_data, node_data) + work_actions_times[9] + work_actions_times[0] + 0 + work_actions_times[1]
    print("Failover: " + str(res))
    return res


def compute_predicted_Failover_TBNewView(total_network_data, node_data):
    """ 
    Arguments:
        total_network_data -- map[src node][target node] -> list of network tuples
        actual_method_latencies -- map of method name to list of latencies
        // NoOps(i, j) = no-op-action i + ... +  no-op-action j-1
          
        TBNewView() = MaxQueueTime + EpochLength + Timeout() + NoOps(0, 10) + StepToTimeDelta(RslStep(7)) + EpochLength
        +  Timeout() + NoOps(0, 10) + StepToTimeDelta(RslStep(7)) + HBPeriod + Timeout() 
        + NoOps(0, 9) + StepToTimeDelta(RslStep(9)) + D+ MaxQueueTime + NoOps(0, 8) + StepToTimeDelta(RslStep(8))
    """    
    work_actions_times, noop_action_times, max_queue_time = mean_action_times(node_data)
    # delay_c_oh = mean_network_delay(compute_actual_network(total_network_data, CLIENT, L))
    delay_oh_or = mean_network_delay(compute_actual_network(total_network_data, L, C))
    
    
    r1 = max_queue_time + EpochLength + 0 + noop_actions_up_to(noop_action_times, 0, 10) + work_actions_times[7] + EpochLength
    r2 = noop_actions_up_to(noop_action_times, 0, 10) + work_actions_times[7] + HBPeriod + 0
    r3 = noop_actions_up_to(noop_action_times, 0, 9) + work_actions_times[9] + delay_oh_or + max_queue_time + noop_actions_up_to(noop_action_times, 0, 8) + + work_actions_times[8]
    res = r1 + r2 + r3
    print("TBNewView: " + str(res))
    return res


def noop_actions_up_to(noop_actions_times, i, j):
    res = 0
    for x in range(i, j):
        res += noop_actions_times[x]
    return res



def compute_actual_node(total_node_data_f):
    """maps total_node_data to res: method_name -> list of latencies
    Args:
        total_node_data : total_node_data_f[node_id][method_name][trial] = list of durations
    """
    res = dict()
    for node in total_node_data_f:
        for method in total_node_data_f[node]:
            if method not in res:
                    res[method] = []
            for t in total_node_data_f[node][method]:
                res[method].extend(total_node_data_f[node][method][t][THROW:-THROW])
    return res
    
def mean_network_delay(network_delays):
    return sum(network_delays)/len(network_delays)

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
                for t in total_f_node_data[node][name]:
                # try:
                    aggregate.extend(total_f_node_data[node][name][t])
                # except:
                    # print("node " + str(node) + " has no data for " + name)
            work_res[method_id] = np.mean(aggregate) * 10
        else:
            sum_times = 0
            count = 0
            for node in total_f_node_data.keys():
                # print(total_f_node_data[node][name])
                for t in total_f_node_data[node][name]:
                    sum_times += np.sum(total_f_node_data[node][name][t])
                    count += len(total_f_node_data[node][name][t])
            if count > 0:
                work_res[method_id] = sum_times/float(count)
            else:
                work_res[method_id] = 0.0
    for method_id, name in NOOP_METHODS.items():
        sum_times = 0
        count = 0
        for node in total_f_node_data.keys():
            for t in total_f_node_data[node][name]:
                sum_times += np.sum(total_f_node_data[node][name][t])
                count += len(total_f_node_data[node][name][t])
        noop_res[method_id] = sum_times/float(count)
    sum_times = 0
    count = 0
    for node in total_f_node_data.keys():
        for t in total_f_node_data[node][name]:
            sum_times += np.sum(total_f_node_data[node][MAX_QUEUE][t])
            count += len(total_f_node_data[node][MAX_QUEUE][t])
    max_queue_res = sum_times/float(count)
    return work_res, noop_res, max_queue_res

def compute_actual_network(total_network_data, src, targ):
    """
    Arguments:
        total_network_data -- total_network_data[i][j] is the timings for node i to node j
    """
    latencies = [p[0]/2.0 for p in total_network_data[src][targ] if START < p[1] and p[1] < END and not p[2]]
    latencies.extend([p[0]/2.0 for p in total_network_data[targ][src] if START < p[1] and p[1] < END and not p[2]])
    return latencies

if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)
