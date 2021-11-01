import sys
import pickle
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import seaborn as sns

F_VALUES = [1]

METHODS = ["LReplicaNextProcessPacket",
            "LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a",
            "LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoop",
            "LReplicaNextSpontaneousMaybeEnterPhase2",
            "LReplicaNextSpontaneousMaybeEnterPhase2Noop",
            "LReplicaNextReadClockMaybeNominateValueAndSend2a",
            "LReplicaNextReadClockMaybeNominateValueAndSend2aNoop",
            "LReplicaNextSpontaneousTruncateLogBasedOnCheckpoints",
            "LReplicaNextSpontaneousTruncateLogBasedOnCheckpointsNoop",
            "LReplicaNextSpontaneousMaybeMakeDecision",
            "LReplicaNextSpontaneousMaybeMakeDecisionNoop",
            "LReplicaNextSpontaneousMaybeExecute",
            "LReplicaNextSpontaneousMaybeExecuteNoop",
            "LReplicaNextReadClockCheckForViewTimeout",
            "LReplicaNextReadClockCheckForViewTimeoutNoop",
            "LReplicaNextReadClockCheckForQuorumOfViewSuspicions",
            "LReplicaNextReadClockCheckForQuorumOfViewSuspicionsNoop",
            "LReplicaNextReadClockMaybeSendHeartbeat",
            "LReplicaNextReadClockMaybeSendHeartbeatNoop",
            "MaxQueueing"
           ]

plt.rc('xtick', labelsize=8)    # fontsize of the tick labels
plt.rc('ytick', labelsize=8) 

def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)

    for f in F_VALUES:
        """
        total_f_node_data[node_id][method_name][trial] = list of durations
        total_f_client_data[i] = list of client durations for trial i
        total_f_client_start_end[i] = (start, end) time of trial i, defined from start of first request to end of last request
        """
        total_f_node_data, total_f_client_data, total_f_client_start_end = analyze_f(exp_dir, f)

        # Save data
        with open("%s/total_f%d_node_data.pickle" %(exp_dir, f), 'wb') as handle:
            pickle.dump(total_f_node_data, handle)
        with open("%s/total_f%d_client_data.pickle" %(exp_dir, f), 'wb') as handle:
            pickle.dump(total_f_client_data, handle)
        with open("%s/total_f%d_client_start_end.pickle" %(exp_dir, f), 'wb') as handle:
            pickle.dump(total_f_client_start_end, handle)

        with open("%s/total_f%d_node_data.pickle" %(exp_dir, f), 'rb') as handle:
            total_f_node_data = pickle.load(handle)
        with open("%s/total_f%d_client_data.pickle" %(exp_dir, f), 'rb') as handle:
            total_f_client_data = pickle.load(handle)
        with open("%s/total_f%d_client_start_end.pickle" %(exp_dir, f), 'rb') as handle:
            total_f_client_start_end = pickle.load(handle)

        # Print graphs
        print("\tDrawing charts for f=%d" %f)
        plot_client_figures("f_%d_client_plots" %f, exp_dir, total_f_client_data, total_f_client_start_end)
        plot_individual_figures("f_%d_nodes_individual_plots" %f, exp_dir, total_f_node_data)
        # plot_overall_figures("f_%d_nodes_aggregate_plots" %f, exp_dir, total_f_node_data)
    print("Done")


def analyze_f(exp_dir, f):
    """Analyze all the trials for sub-experiment with f failures

    Arguments:
        exp_dir {string} -- root directory of the experiment
        f {int} -- f value for rsl
    Returns:
        dict[node_id] -> (dict[method_name] -> list of durations)
        dict[trial_num] -> list of durations
    """
    f_dir = "%s/f_%d" %(exp_dir, f)
    print("\tAnalyzing data for f=%d in %s" %(f, f_dir))

    """
    f_node_data[node_id][method_name] = list of durations
    f_client_data[i] = list of client durations for trial i
    f_client_start_end[i] = (start, end) time of trial i, defined from start of first request to end of last request
    """
    f_node_data = dict()  
    f_client_data = dict() 
    f_client_start_end = dict() 

    # Gather a list of trial directories
    trial_dirs = []  # list of trial directories under f_dir
    for root, dirs, files in os.walk(f_dir):
        if dirs != []:
            trial_dirs.extend(["%s/%s" %(f_dir, d) for d in dirs])

    print("\t\tAnalyzing data for %d trials" %len(trial_dirs))

    # Look under each trial directory and analyze results
    for trial_dir in trial_dirs:
        trial_num = int(trial_dir.split('trial')[1])
        trial_data, f_client_data[trial_num], f_client_start_end[trial_num] = analyze_trial_dir(trial_dir)
        # trial_data[node_id][method_name] = list of durations
        for node_id in trial_data:
            if node_id not in f_node_data:
                f_node_data[node_id] = dict()  # f_data[node_id][method_name][trial] = list of durations
            for method_name in trial_data[node_id]:
                if method_name not in f_node_data[node_id]:
                    f_node_data[node_id][method_name] = dict()
                f_node_data[node_id][method_name][trial_num] = []
                f_node_data[node_id][method_name][trial_num].extend(trial_data[node_id][method_name])
    return f_node_data, f_client_data, f_client_start_end


def analyze_trial_dir(trial_dir):
    """
    Arguments:
        trial_dir {string} -- absolute directory of a trial
    Returns:
        dict of (node_id -> method_name -> [durations...])
        list of client durations for trial
        tuple (start, end) time of this client, defined from start of first request to end of last request
    """
    print("\t\tAnalyzing trial %s" %trial_dir)
    trial_data = dict()   # trial_data[node_id][method_name] = list of durations
    client_data = []
    client_start_end = (-1, -1)
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
                        trial_data[node_id][method_name] = analyze_node_csv("%s/%s" %(trial_dir, csv))
            elif file_extension == '.log' and 'client' in file_name:
                client_data, client_start_end = analyze_client_csv("%s/%s" %(trial_dir, csv))
    return trial_data, client_data, client_start_end


def analyze_node_csv(filepath):
    durations_milli = []
    with open(filepath, 'r') as node1:
        csvreader = csv.reader(node1, delimiter=',',)
        for row in csvreader:
            if 'init' in row[0]:
                continue
            if len(row) > 2 and int(row[0]) >= 0:
                start_time = int(row[2])
                end_time = int(row[3])
                dur = (end_time - start_time)/1_000_000.0  # duration in milliseconds
                durations_milli.append(dur)
    return durations_milli


def analyze_client_csv(filepath):
    """
    Arguments:
        filepath {string} -- path to client csv file

    Returns:
        list [durations ... ]
        tuple (start, end) time of this client, defined from start of first request to end of last request
    """
    durations_milli = []
    start = 999999999999
    end = 0
    with open(filepath, 'r') as client:
        csvreader = csv.reader(client, delimiter=' ',)
        for row in csvreader:
            if 'TIMEOUT' in row[0] or 'DEBUG:' in row:
                continue
            req_start = float(row[1])
            req_end = float(row[2])
            start = min(start, req_start)   # Note: this is rather inefficient
            end = max(end, req_end)
            durations_milli.append(req_end - req_start)
    return durations_milli, (start, end)


def plot_individual_figures(name, root, data):
    """ Plot a the data for each method, for each node
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        data -- data[node_id][method_name][trial] = list of durations
    """
    num_nodes = len(data)               # num rows
    nodes = list(data.keys())
    nodes.sort()
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        # Plot historgrams
        for method in METHODS:
            fig, axes = plt.subplots(num_nodes, 1, figsize=(8.5, 11), sharex=True)
            fig.suptitle(method, fontsize=12, fontweight='bold')
            sns.despine(left=True)
            row = 0
            for node in nodes:
                print("\t\tDrawing individual chart for node %d : %s" %(node, method))
                try:
                    durations_milli = []
                    for t in data[node][method]:
                        durations_milli.extend(data[node][method][t])
                except KeyError:
                    print("No data for method %s in node %d" %(method, node))
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
            for ax, row in zip(axes, nodes):
                ax.annotate("Node %d" %row, xy=(0, 0.5), xytext=(-ax.yaxis.labelpad - pad, 0),
                        xycoords=ax.yaxis.label, textcoords='offset points',
                        fontsize=10, ha='right', va='center')
            fig.tight_layout()
            fig.subplots_adjust(left=0.2, top=0.92, right=0.85)
            plt.subplots_adjust(hspace=0.2)
            pp.savefig(fig)
            plt.close(fig)
        
        # Plot time series
        for method in METHODS:
            for node in nodes:
                num_trials = len(data[node][method])
                fig, axes = plt.subplots(num_trials, 1, figsize=(8.5, 11), sharex=True)
                fig.suptitle(method, fontsize=12, fontweight='bold')
                sns.despine(left=True)

                row = 0
                print("\t\tDrawing time series chart for node %d : %s" %(node, method))
                for t in data[node][method]:
                    try:
                        durations_milli = data[node][method][t]
                    except KeyError:
                        print("No data for method %s in node %d" %(method, node))
                        continue
                    this_ax = axes[row]
                    # Plot the subfigure 
                    this_ax.grid()
                    this_ax.scatter(range(len(durations_milli)), durations_milli, marker='.', s=3)
                    if len(durations_milli) > 5:
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
                for ax, row in zip(axes, nodes):
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
        fig, axes = plt.subplots(len(METHODS), 1, figsize=(8.5, 20), sharex=True)
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

def plot_client_figures(name, root, data, start_end_data):
    """ Plot a the data for each method, for each node
    Arguments:
        name {string} -- name of this figure
        root {string} -- directory to save this figure
        data {dict} -- data[trial_id] = list of durations
        start_end_data {dict} = start_end_data[i] = (start, end) time of trial i, defined from start of first request to end of last request
    """
    print("\t\tDrawing chart for client")
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        # Draw aggregate
        aggregate_duration = []
        for durs in data.values():
            aggregate_duration.extend(durs)
        fig, this_ax = plt.subplots(1, 1, figsize=(8.5, 5), sharex=True)
        fig.suptitle("Aggregate client data over %d trials" %len(list(data.keys())), fontsize=12, fontweight='bold')
        this_ax.grid()
        sns.distplot(aggregate_duration, kde=False, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
        if len(aggregate_duration) > 0:
            stats = AnchoredText(
                generate_client_statistics(aggregate_duration), 
                loc='upper right',  
                prop=dict(size=8),
                bbox_to_anchor=(1.1, 1),
                bbox_transform=this_ax.transAxes
            )
            this_ax.add_artist(stats)
        fig.tight_layout()
        fig.subplots_adjust(left=0.2, top=0.92, right=0.85)
        plt.subplots_adjust(hspace=0.2)
        pp.savefig(fig)
        plt.close(fig)

        # Draw individual trial historgrams
        trials = list(data.keys())
        trials.sort()

        trials_per_page = 5
        trials_pages = [trials[i:i + trials_per_page] for i in range(0, len(trials), trials_per_page)]  

        for trial_page in trials_pages:
            fig, axes = plt.subplots(len(trial_page), 1, figsize=(8.5, 11), sharex=True)
            fig.suptitle("Client data for each trial", fontsize=12, fontweight='bold')
            sns.despine(left=True)

            row = 0
            for t in trial_page:
                durations_milli = data[t]
                if len(trial_page) == 1:
                    this_ax = axes
                else:
                    this_ax = axes[row]
                # Plot the subfigure 
                this_ax.grid()
                sns.distplot(durations_milli, kde=False, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
                if len(durations_milli) > 0:
                    stats = AnchoredText(
                        generate_client_statistics(durations_milli, start_end=start_end_data[t]), 
                        loc='upper right',  
                        prop=dict(size=8),
                        bbox_to_anchor=(1.1, 1),
                        bbox_transform=this_ax.transAxes
                    )
                    this_ax.add_artist(stats)
                row += 1
            pad = 5
            if len(trial_page) == 1:
                axes.annotate("Trial %d" %trial_page[0], xy=(0, 0.5), xytext=(-axes.yaxis.labelpad - pad, 0),
                            xycoords=axes.yaxis.label, textcoords='offset points',
                            fontsize=10, ha='right', va='center')
            else:
                for ax, t in zip(axes, trial_page):
                    ax.annotate("Trial %d" %t, xy=(0, 0.5), xytext=(-ax.yaxis.labelpad - pad, 0),
                            xycoords=ax.yaxis.label, textcoords='offset points',
                            fontsize=10, ha='right', va='center')
            fig.tight_layout()
            fig.subplots_adjust(left=0.2, top=0.92, right=0.85)
            plt.subplots_adjust(hspace=0.2)
            pp.savefig(fig)
            plt.close(fig)

        # Draw individual trial timeseries
        trials_per_page = 5
        trials_pages = [trials[i:i + trials_per_page] for i in range(0, len(trials), trials_per_page)]  

        for trial_page in trials_pages:
            fig, axes = plt.subplots(len(trial_page), 1, figsize=(8.5, 11), sharex=True)
            fig.suptitle("Client data timeseries for each trial", fontsize=12, fontweight='bold')
            sns.despine(left=True)

            row = 0
            for t in trial_page:
                durations_milli = data[t]
                if len(trial_page) == 1:
                    this_ax = axes
                else:
                    this_ax = axes[row]
                # Plot the subfigure 
                this_ax.grid()
                this_ax.scatter(range(len(durations_milli)), durations_milli, marker='.', s=3)
                row += 1
            pad = 5
            if len(trial_page) == 1:
                axes.annotate("Trial %d" %trial_page[0], xy=(0, 0.5), xytext=(-axes.yaxis.labelpad - pad, 0),
                            xycoords=axes.yaxis.label, textcoords='offset points',
                            fontsize=10, ha='right', va='center')
            else:
                for ax, t in zip(axes, trial_page):
                    ax.annotate("Trial %d" %t, xy=(0, 0.5), xytext=(-ax.yaxis.labelpad - pad, 0),
                            xycoords=ax.yaxis.label, textcoords='offset points',
                            fontsize=10, ha='right', va='center')
            fig.tight_layout()
            fig.subplots_adjust(left=0.2, top=0.92, right=0.85)
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
    if len(input) > 1:
        res.append("σ = %.4f" %statistics.stdev(input))
    res.append("99.9%% = %.3f" %np.percentile(input, 99.9))
    res.append("")
    res.append("max = %.3f" %np.max(input))
    res.append("min = %.3f" %np.min(input))
    return "\n".join(res)

def generate_client_statistics(input, start_end=None):
    """
    Generates a string containing some statistics for the input
    Arguments:
        input {list} -- list of numbers
        start_end {tuple} -- (start, end) time of trial i, defined from start of first request to end of last request
    """
    res = []
    if start_end is not None:
        start = start_end[0]
        end = start_end[1]
        assert start > 0 and end > 0
        res.append(f"rate = {'{:,}'.format((len(input)/float(end-start)*1000.0))} reqs/sec")
        res.append(f"duration = %.2f s" %(float(end-start)/1000.0))
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
    exp_dir =sys.argv[1]
    main(exp_dir)