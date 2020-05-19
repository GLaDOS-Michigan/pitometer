import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import seaborn as sns

F_VALUES = [1, 2, 3, 4, 5, 6]
THROWAWAY = 1000  # How many initial executions to ignore
METHODS = ["LReplicaNextProcessPacket",
           "LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a",
           "LReplicaNextSpontaneousMaybeEnterPhase2",
           "LReplicaNextReadClockMaybeNominateValueAndSend2a",
           "LReplicaNextSpontaneousTruncateLogBasedOnCheckpoints",
           "LReplicaNextSpontaneousMaybeMakeDecision",
           "LReplicaNextSpontaneousMaybeExecute",
           "LReplicaNextReadClockCheckForViewTimeout",
           "LReplicaNextReadClockCheckForQuorumOfViewSuspicions",
           "LReplicaNextReadClockMaybeSendHeartbeat"
           ]

plt.rc('xtick', labelsize=8)    # fontsize of the tick labels
plt.rc('ytick', labelsize=8) 

def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)

    total_data = dict()  # total_data[f][node_id][method_name] = list of durations

    for f in F_VALUES:
        total_data[f] = analyze_f(exp_dir, f)

    # Print graphs
    for f in F_VALUES:
        print("\tDrawing charts for f=%d" %f)
        plot_individual_figures("f_%d_individual_plots" %f, exp_dir, total_data[f])
        plot_overall_figures("f_%d_aggregate_plots" %f, exp_dir, total_data[f])
    print("Done")


def analyze_f(exp_dir, f):
    f_dir = "%s/f_%d" %(exp_dir, f)
    print("\tAnalyzing data for f=%d in %s" %(f, f_dir))

    f_data = dict()  # f_data[node_id][method_name] = list of durations

    # Gather a list of trial directories
    trial_dirs = []  # list of trial directories under f_dir
    for root, dirs, files in os.walk(f_dir):
        if dirs != []:
            trial_dirs.extend(["%s/%s" %(f_dir, d) for d in dirs])

    print("\t\tAnalyzing data for %d trials" %len(trial_dirs))

    # Look under each trial directory and analyze results
    for trial_dir in trial_dirs:
        trial_data = analyze_trial_dir(trial_dir)
        for node_id in trial_data:
            if node_id not in f_data:
                f_data[node_id] = dict()  # f_data[node_id][method_name] = list of durations
            for method_name in trial_data[node_id]:
                if method_name not in f_data[node_id]:
                    f_data[node_id][method_name] = []
                f_data[node_id][method_name].extend(trial_data[node_id][method_name])
    return f_data


def analyze_trial_dir(trial_dir):
    """
    Arguments:
        trial_dir {string} -- absolute directory of a trial
    Returns:
        dict[node_id] -> (dict[method_name] -> list of durations)
    """
    print("\t\tAnalyzing trial for %s" %trial_dir)
    trial_data = dict()   # trial_data[node_id][method_name] = list of durations
    for root, _, files in os.walk(trial_dir):
        files = [f for f in files if not f[0] == '.']  # ignore hidden files
        for csv in files:
            file_name, file_extension = os.path.splitext(csv)
            if file_extension == '.csv':
                if 'node' in file_name:
                    # Parse the csv file name
                    node_id = int(file_name.split('_')[1])
                    method_name = file_name.split('_')[2]
                    if method_name in METHODS:
                        if node_id not in trial_data:
                            trial_data[node_id] = dict()
                        trial_data[node_id][method_name] = analyze_csv("%s/%s" %(trial_dir, csv))
                else:
                    # TODO: Ignore the client for now
                    pass
    return trial_data


def analyze_csv(filepath):
    durations_nano = []
    with open(filepath, 'r') as node1:
        csvreader = csv.reader(node1, delimiter=',',)
        for row in csvreader:
            if len(row) > 2 and int(row[0]) > THROWAWAY:
                event_type = row[1]
                if event_type == 'Start':
                    prevStart = int(row[3])
                if event_type == 'End':
                    dur = int(row[3]) - prevStart  # duration in nanoseconds
                    durations_nano.append(dur)
    durations_milli = list(map(lambda x: x/1_000_000.0, durations_nano))
    return durations_milli


def plot_individual_figures(name, root, data):
    """ Plot a the data for each method, for each node
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        data -- data[node_id][method_name] = list of durations
    """

    num_nodes = len(data)               # num rows

    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        for method in METHODS:
            fig, axes = plt.subplots(num_nodes, 1, figsize=(8.5, 11), sharex=True)
            fig.suptitle(method, fontsize=12, fontweight='bold')
            sns.despine(left=True)

            row = 0
            nodes = list(data.keys())
            nodes.sort()
            for node in nodes:
                print("\t\tDrawing individual chart for node %d : %s" %(node, method))
                try:
                    durations_milli = data[node][method]
                except KeyError:
                    # print("No data for method %s in node %d" %(method, node))
                    continue

                this_ax = axes[row]
                # Plot the subfigure 
                this_ax.grid()
                sns.distplot(durations_milli, kde=False, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
                if len(durations_milli) > 0:
                    stats = AnchoredText(
                        generate_statistics(durations_milli), 
                        loc='upper right',  
                        prop=dict(size=8),
                        bbox_to_anchor=(1.1, 1),
                        bbox_transform=this_ax.transAxes
                    )
                    this_ax.add_artist(stats)
                row += 1
            pad = 5
            for ax, row in zip(axes, list(data.keys())):
                ax.annotate("Node %d" %row, xy=(0, 0.5), xytext=(-ax.yaxis.labelpad - pad, 0),
                        xycoords=ax.yaxis.label, textcoords='offset points',
                        fontsize=10, ha='right', va='center')
            fig.tight_layout()
            fig.subplots_adjust(left=0.2, top=0.92, right=0.85)
            plt.subplots_adjust(hspace=0.2)
            pp.savefig(fig)
            plt.close(fig)


def plot_overall_figures(name, root, data):
    """ Plot a the data for each method, aggregated across all nodes
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        data -- data[node_id][method_name] = list of durations
    """

    # Collect and organize data
    # aggregated_method_data[method] is a list of all durations for method
    aggregated_method_data = dict()  
    for method_name in METHODS:
        if method_name not in aggregated_method_data:
            aggregated_method_data[method_name] = []
        for node_id in data.keys():
            aggregated_method_data[method_name].extend(data[node_id][method_name])

    # Plot the data
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        fig, axes = plt.subplots(len(METHODS), 1, figsize=(8.5, 11), sharex=True)
        sns.despine(left=True)

        row = 0
        for method in METHODS:
            print("\t\tDrawing overall chart for method %s" %(method))
            try:
                durations_milli = aggregated_method_data[method]
            except KeyError:
                # print("No data for method %s in node %d" %(method, node))
                row += 1
                continue

            this_ax = axes[row]
            # Plot the subfigure 
            this_ax.grid()
            sns.distplot(durations_milli, kde=False, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
            if len(durations_milli) > 0:
                stats = AnchoredText(
                    generate_statistics(durations_milli), 
                    loc='upper right',  
                    prop=dict(size=8),
                    bbox_to_anchor=(1.1, 1),
                    bbox_transform=this_ax.transAxes
                )
                this_ax.add_artist(stats)
            row += 1
        pad = 5
        for ax, row in zip(axes, METHODS):
            ax.annotate(row, xy=(0, 0.5), xytext=(-ax.yaxis.labelpad - pad, 0),
                    xycoords=ax.yaxis.label, textcoords='offset points',
                    fontsize=7, ha='right', va='center')
        fig.tight_layout()
        fig.subplots_adjust(left=0.4, top=0.92, right=0.85)
        plt.subplots_adjust(hspace=0.2)
        pp.savefig(fig)
        plt.close(fig)


def generate_statistics(input):
    """
    Generates a string containing some statistics for the input
    Arguments:
        input -- list of numbers
    """
    res = []
    res.append(f"n = {'{:,}'.format(len(input))}")
    res.append("μ = %.3f" %statistics.mean(input))
    res.append("σ = %.4f" %statistics.stdev(input))
    res.append("99.9%% = %.3f" %np.percentile(input, 99.9))
    res.append("")
    res.append("max = %.3f" %np.max(input))
    res.append("min = %.3f" %np.min(input))
    return "\n".join(res)


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)