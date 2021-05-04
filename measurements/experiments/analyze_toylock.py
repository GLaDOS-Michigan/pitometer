import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import seaborn as sns
import pickle


NODES = list(range(1, 21))
DELAYS = [0, 200, 1_000, 5_000]  # units of microseconds


def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)

    # each data is dict of (size -> delay -> node -> [ durs ])
    total_grant_data, total_accept_data, total_round_data = parse_files(exp_dir)  
    
    with open("%s/%s" %(exp_dir, 'total_grant_data.pickle'), 'wb') as handle:
        pickle.dump(total_grant_data, handle)
    with open("%s/%s" %(exp_dir, 'total_accept_data.pickle'), 'wb') as handle:
        pickle.dump(total_accept_data, handle)
    with open("%s/%s" %(exp_dir, 'total_round_data.pickle'), 'wb') as handle:
        pickle.dump(total_round_data, handle)

    with open("%s/%s" %(exp_dir, 'total_grant_data.pickle'), 'rb') as handle:
        total_grant_data = pickle.load(handle)
    with open("%s/%s" %(exp_dir, 'total_accept_data.pickle'), 'rb') as handle:
        total_accept_data = pickle.load(handle)
    with open("%s/%s" %(exp_dir, 'total_round_data.pickle'), 'rb') as handle:
        total_round_data = pickle.load(handle)

    print("\nPlotting graphs for experiment %s" %exp_dir)
    # Plot Grant
    plot_grant_or_accept("nodeGrant", exp_dir, total_grant_data)

    # Plot Accept
    plot_grant_or_accept("nodeAccept", exp_dir, total_accept_data)

    # Plot Rounds
    plot_round("rounds", exp_dir, total_round_data)
    print("Done")


def parse_files(exp_dir):
    """ Parse all csv files under exp_dir into dict format

    Arguments:
        exp_dir {string} -- Root directory of the experiment

    Returns:
        total_grant_data -- dict of (size -> delay -> node -> [ durs ])
        total_accept_data -- dict of (size -> delay -> node -> [ durs ])
        total_round_data -- dict of (size -> delay -> node -> [ durs ])
    """
    exp_dir = os.path.abspath(exp_dir)

    total_grant_data = dict()    # size -> delay -> node -> [ durs ]
    total_accept_data = dict()   # size -> delay -> node -> [ durs ]
    total_round_data = dict()    # size -> delay -> node -> [ durs ]

    for root, _, files in os.walk(exp_dir):
        files = [f for f in files if not f[0] == '.']  # ignore hidden files
        if files != [] and 'delay' in root:
            # This is a leaf directory containing trial csv files
            print("\tAnalyzing trial %s" %root)
            size = int(root.split('size')[1].split('/')[0])
            delay = int(root.split('delay')[1].split('/')[0])
            
            # Make sure dict is initialized
            if size not in total_grant_data:
                total_grant_data[size] = dict()
            if size not in total_accept_data:
                total_accept_data[size] = dict()
            if size not in total_round_data:
                total_round_data[size] = dict()
            if delay not in total_grant_data[size]:
                total_grant_data[size][delay] = dict()
            if delay not in total_accept_data[size]:
                total_accept_data[size][delay] = dict()
            if delay not in total_round_data[size]:
                total_round_data[size][delay] = dict()
            
            for f in files:
                file_name, file_extension = os.path.splitext(f)
                if file_extension == '.csv':
                    node_id = int(file_name.split('skynode')[1].split('.')[0])
                    if "grant" in file_name:
                        if node_id not in total_grant_data[size][delay]:
                            total_grant_data[size][delay][node_id] = []
                        if node_id not in total_round_data[size][delay]:
                            total_round_data[size][delay][node_id] = []
                        total_grant_data[size][delay][node_id].extend(analyze_grant_or_accept_csv("%s/%s" %(root, f)))
                        total_round_data[size][delay][node_id].extend(analyze_round_csv("%s/%s" %(root, f)))
                    if "accept" in file_name:
                        if node_id not in total_accept_data[size][delay]:
                            total_accept_data[size][delay][node_id] = []
                        total_accept_data[size][delay][node_id].extend(analyze_grant_or_accept_csv("%s/%s" %(root, f)))
    return total_grant_data, total_accept_data, total_round_data


def plot_grant_or_accept(name, root, total_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
    """
    for delay in DELAYS:
        plot_grant_or_accept_delay(name, root, delay, total_data)


def plot_grant_or_accept_delay(name, root, delay, total_data):
    """
    Aggregate all data in total_data for this delay, since execution time should be 
    independent of node id or ring size
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        delay -- delay to plot, in units of microseconds
        total_data -- dict of (size -> delay -> node -> [ durs ])
    """
    print("\tPlot %s for delay %.1f ms" %(name, delay/1000.0))
    # Get total data for this delay, regardless of size or node
    total_aggregate_data = []
    for size in total_data.keys():
        for node in total_data[size][delay].keys():
            total_aggregate_data.extend(total_data[size][delay][node])
    # Plot graph
    with PdfPages("%s/delay%d_%s.pdf" %(root, delay, name)) as pp:
        fig, axes = plt.subplots(2, 1, figsize=(8.5, 11), sharex=False)
        fig.suptitle("%s, delay %.1f ms" %(name, delay/1000.0), fontweight='bold')
        plot_histogram(axes[0], total_aggregate_data)
        plot_cdf(axes[1], total_aggregate_data)
        pp.savefig(fig)
        plt.close(fig)


def plot_round(name, root, total_round_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
    """
    for delay in DELAYS:
        plot_round_delay(name, root, delay, total_round_data)
    return


def plot_round_delay(name, root, delay, total_round_data):
    """
    On each page, plot 1 graph for a single size, from the perspective of the leader node
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        delay -- delay to plot, in units of microseconds
        total_round_data -- dict of (size -> delay -> node -> [ durs ])
    """
    print("\tPlot %s for delay %.1f ms" %(name, delay/1000.0))
    sizes = list(total_round_data.keys())
    sizes.sort()

    with PdfPages("%s/delay%d_%s.pdf" %(root, delay, name)) as pp:
        for size in sizes:
            # all_nodes = list(total_round_data[size][delay].keys())
            # all_nodes.sort()
            leader_node = 20
            data = total_round_data[size][delay][leader_node]
            fig, axes = plt.subplots(2, 1, figsize=(8.5, 11), sharex=False)
            fig.suptitle("%s, delay %.1f ms, size %d" %(name, delay/1000.0, size), fontweight='bold')
            plot_histogram(axes[0], data)
            plot_cdf(axes[1], data)
            pp.savefig(fig)
            plt.close(fig)

def plot_cdf(this_ax, data, title=None):
    """Plot a histogram
    Arguments:
        this_ax {axes} -- axes on which to plot
        title {string}  -- title of this_ax
        data {list} -- list of data
    """
    kwargs = {'cumulative': True}
    sns.distplot(data, hist_kws=kwargs, kde_kws=kwargs, vertical=True)
    this_ax.set_xlim(0, 1)
    this_ax.xaxis.set_ticks(np.arange(0, 1, 0.1))
    this_ax.grid()
    if title is not None:
        this_ax.set_title(title)
    this_ax.set_xlabel('cumulative probability', fontsize=10)
    this_ax.set_ylabel('latency (ms)', fontsize=10)

def plot_histogram(this_ax, data, title=None, stats=True, kde=False):
    """Plot a histogram
    Arguments:
        this_ax {axes} -- axes on which to plot
        data {list} -- list of data
        title {string}  -- title of this_ax
        stats {bool} -- toggle statistics box
        kde {bool} -- toggle kde option
    """
    sns.distplot(data, kde=kde, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
    if len(data) > 0:
        stats = AnchoredText(
                generate_statistics(data), 
                loc='upper right',  
                prop=dict(size=8),
                bbox_to_anchor=(1.1, 1),
                bbox_transform=this_ax.transAxes
                )
        this_ax.add_artist(stats)
    this_ax.grid()
    if title is not None:
        this_ax.set_title(title)
    this_ax.set_xlabel('latency (ms)', fontsize=10)
    this_ax.set_ylabel('count', fontsize=10)



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


def analyze_grant_or_accept_csv(filepath):
    """ Computes the time taken for each nodeAccept or nodeGrant from the csv file
    Arguments:
        filepath -- path to a nodeGrant/nodeAccept csv file
    Returns:
        durations_milli -- A list of the times taken for each round in milliseconds
    """
    durations_nano = []
    with open(filepath, 'r') as node1:
        csvreader = csv.reader(node1, delimiter=',',)
        print(filepath)
        for row in csvreader:
            if row != []:
                start_time = int(row[2])
                end_time = int(row[3])
                dur = end_time - start_time  # duration in nanoseconds
                durations_nano.append(dur)
    durations_milli = list(map(lambda x: x/1_000_000.0, durations_nano))
    return durations_milli

def analyze_round_csv(filepath):
    """ Computes the time taken for each round from nodeGrant csv data.
        The time for a round is defined by the time between the end of a grant and the end of the next grant
    Arguments:
        filepath -- path to a nodeGrant csv file
    Returns:
        durations_milli -- A list of the times taken for each round in milliseconds
    """
    durations_nano = []
    with open(filepath, 'r') as node1:
        csvreader = csv.reader(node1, delimiter=',',)
        for row in csvreader:
            if row != [] or int(row[0]) >= 0:
                round_num = int(row[0])
                if round_num == 0:
                    round_start_time = int(row[3])
                    continue
                round_end_time = int(row[3])
                dur = round_end_time - round_start_time  # duration in nanoseconds
                durations_nano.append(dur)
                round_start_time = round_end_time
    durations_milli = list(map(lambda x: x/1_000_000, durations_nano))
    return durations_milli


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)