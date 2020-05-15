import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
import textwrap as tw
from matplotlib.offsetbox import AnchoredText
import seaborn as sns

NODES = [1, 2, 3, 4, 5, 6]
NODES.sort()
THROWAWAY = 10 # Number of starting readings to throw away

def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)
    for root, _, files in os.walk(exp_dir):
        files = [f for f in files if not f[0] == '.']  # ignore hidden files
        if files != []:
            # This is a leaf directory containing trial csv files
            print("\tAnalyzing trial %s" %root)
            
            # total_node_data[i][j] is the timings for node i to node j
            total_node_data = [[None for j in NODES] for i in NODES] 
            titles = [["node%d -> node%d" %(node_i, node_j) for node_j in NODES] for node_i in NODES] 

            for i in range(len(NODES)):
                grep_str = "node%d" %NODES[i]
                nodei_csvs= [f for f in files if grep_str in f.split('-')[0] and ".csv" in f]
                nodei_csvs.sort()
                # print(nodei_csvs)

                for j in range(len(nodei_csvs)):
                    # File for log of nodei->nodej
                    # print(NODES[j], j)
                    i_j_csv = nodei_csvs[j]
                    grep_str = "node%d" %NODES[j]
                    assert grep_str in i_j_csv.split('-')[1]  # sanity check to make sure we have the right target file

                    total_node_data[i][j] = analyze_csv("%s/%s" %(root, i_j_csv))

            payload = root.split('/')[-1]
            plot_figures("rtt_%s" %payload, root, total_node_data, titles)
    print("Done")


def plot_figures(name, root, total_data, titles):
    """ Plot a figure where each subfigure is from an element in total_data
    Assumes total_data and titles are 2d arrays of same shape, len(NODES) * len(NODES)
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- list of lists of data
        titles -- list of titles for each subfigure
    """

    assert len(total_data) == len(NODES) and len(total_data[0]) == len(NODES)
    assert len(titles) == len(NODES) and len(titles[0]) == len(NODES)

    fig, axes = plt.subplots(len(NODES), len(NODES), figsize=(4*len(NODES), 4*len(NODES)), sharex=True)
    fig.suptitle(name)
    sns.despine(left=True)
    
    for i in range(len(NODES)):
        for j in range(len(NODES)):
            i_j_data = total_data[i][j]
            i_j_title = titles[i][j]
            this_ax = axes[i][j]

            this_ax.set_title(i_j_title, fontsize=9)
            this_ax.grid()
            sns.distplot(i_j_data, kde=False, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
            stats = AnchoredText(
                generate_statistics(i_j_data), 
                loc='upper right',  
                prop=dict(size=8),
                bbox_to_anchor=(1.1, 1),
                bbox_transform=this_ax.transAxes
            )
            this_ax.add_artist(stats)
            if i == len(NODES) - 1:
                this_ax.set_xlabel('round trip time (ms)', fontsize=9)
            # this_ax.set_ylabel('count', fontsize=9)
            # this_ax.set_xlim(0, x_max)
            # this_ax.set_ylim(0, 1)

    # Display some global figures
    global_data = []
    for row in total_data:
        for rtts in row:
            global_data.extend(rtts)

    global_stats =  "Global statistics:\n%s" %generate_statistics(global_data)
    plt.figtext(0.8, 0.91, global_stats,
            fontsize=12,
            bbox=dict(boxstyle="round", facecolor='#D8D8D8',
            ec="0.5", pad=0.5, alpha=1), fontweight='bold')

    # Draw plot
    plt.subplots_adjust(hspace=0.2, wspace=0.3)
    # plt.xlabel('latency (ms)', fontsize=10)
    # plt.ylabel('count', fontsize=10)
    plt.savefig("%s/%s.pdf" %(root, name))
    plt.close(fig)


def generate_statistics(input):
    """
    Generates a string containing some statistics for the input
    Arguments:
        input -- list of numbers
    """
    if len(input) == 0:
        return ""
    res = []
    res.append("n = %d" %len(input))
    res.append("μ = %.3f" %statistics.mean(input))
    res.append("σ = %.4f" %statistics.stdev(input))
    res.append("99.9%% = %.3f" %np.percentile(input, 99.9))
    res.append("")
    res.append("max = %.3f" %np.max(input))
    res.append("min = %.3f" %np.min(input))
    return "\n".join(res)


def analyze_csv(filepath):
    durations_nano = []
    with open(filepath, 'r') as node1:
        csvreader = csv.reader(node1, delimiter=',',)
        for row in csvreader:
            if int(row[0]) < THROWAWAY:
                continue
            if row != []:
                event_type = row[1]
                if event_type == 'Start':
                    prevStart = int(row[4])
                if event_type == 'End':
                    dur = int(row[4]) - prevStart  # duration in nanoseconds
                    durations_nano.append(dur)
    durations_milli = list(map(lambda x: x/1_000_000.0, durations_nano))
    return durations_milli


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)