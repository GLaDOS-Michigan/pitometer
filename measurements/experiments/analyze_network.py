import sys
import os
import csv
import pickle
import statistics
import numpy as np
import matplotlib.pyplot as plt
import textwrap as tw
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import seaborn as sns

from conv import *

NODES = [7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19 ,20]
# PAYLOADS = [4, 16, 32, 128, 512]
PAYLOADS = [16]
NODES.sort()
PAYLOADS.sort()
THROWAWAY = 10 # Number of starting readings to throw away

def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)
    for payload in PAYLOADS:
        print("\tAnalyzing payload %d" %payload)
        # {2D map} total_data[i][j] is the timings for node i to node j for this payload
        total_payload_data = dict()  
        for root, dirs, files in os.walk("%s/payload%d" %(exp_dir, payload)):
            files = [f for f in files if not f[0] == '.']  # ignore hidden files
            if dirs == []:
                # This is a leaf directory containing trial csv files
                print("\t\tAnalyzing trial %s" %root)
                
                for i in NODES:
                    grep_str = "node%d" %i
                    nodei_csvs= [f for f in files if grep_str in f.split('-')[0] and ".csv" in f]
                    nodei_csvs.sort()
                    if i not in total_payload_data:
                        total_payload_data[i] = dict()

                    for j in NODES:
                        # File for log of nodei->nodej
                        if j not in total_payload_data[i]:
                            total_payload_data[i][j] = []
                        # find the csv for nodei->nodej
                        grep_str = "node%d." %j
                        i_j_csv = None
                        for candidate in nodei_csvs:
                            if grep_str in candidate.split('-')[1]:
                                i_j_csv = candidate
                        if i_j_csv is not None:
                            total_payload_data[i][j].extend(analyze_csv("%s/%s" %(root, i_j_csv)))
        print("\tDrawing payload %d" %payload)
        with open("%s/total_payload%d_data.pickle" %(exp_dir, payload), 'wb') as handle:
            pickle.dump(total_payload_data, handle)
        with open("%s/total_payload%d_data.pickle" %(exp_dir, payload), 'rb') as handle:
            total_payload_data = pickle.load(handle)
        plot_figures("rtt_payload%d" %payload, exp_dir, total_payload_data)
    print("Done")


def plot_figures(name, root, total_data):
    """ Plot all network figures
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    # assert len(total_data) == len(NODES) and len(total_data[NODES[0]]) == len(NODES)
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        plot_aggregate(pp, name, root, total_data)
        # plot_individuals(pp, name, root, total_data)
        # plot_cdf(pp, name, root, total_data)
        # plot_correlations(pp, name, root, total_data)


def plot_cdf(pp, name, root, total_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        pp -- PdfPages object
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    # fig, axes = plt.subplots(len(NODES), len(NODES), figsize=(4*len(NODES), 4*len(NODES)), sharex=True)
    fig, axes = plt.subplots(2, 2, figsize=(5*2, 5*2))
    fig.suptitle(name)
    sns.despine(left=True)
    
    row = 0
    for i in total_data.keys():
        col = 0
        for j in total_data[i].keys():
            if i > 8 or j > 8:
                continue
            i_j_data = total_data[i][j]
            this_ax = axes[row][col]

            this_ax.set_title("node%d -> node%d" %(i, j), fontsize=9)
            this_ax.grid()

            cdf, bins = raw_data_to_cdf(i_j_data)

            this_ax.plot(cdf, bins[:-1])

            if i == len(NODES) - 1:
                this_ax.set_xlabel('round trip time (ms)', fontsize=9)
            col += 1
        row += 1
    # Draw plot
    plt.subplots_adjust(hspace=0.2, wspace=0.3)
    # plt.xlabel('latency (ms)', fontsize=10)
    # plt.ylabel('count', fontsize=10)
    pp.savefig(fig)
    plt.close(fig)


def plot_correlations(pp, name, root, total_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        pp -- PdfPages object
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    # fig, axes = plt.subplots(len(NODES), len(NODES), figsize=(4*len(NODES), 4*len(NODES)), sharex=True)
    fig, axes = plt.subplots(2, 2, figsize=(5*5, 5*2))
    fig.suptitle(name)
    sns.despine(left=True)
    
    row = 0
    for i in total_data.keys():
        col = 0
        for j in total_data[i].keys():
            if i > 8 or j > 8:
                continue
            i_j_data = total_data[i][j]
            this_ax = axes[row][col]

            this_ax.set_title("node%d -> node%d" %(i, j), fontsize=9)
            this_ax.grid()
            this_ax.scatter(range(len(i_j_data)), i_j_data, marker = '.')

            if i == len(NODES) - 1:
                this_ax.set_xlabel('round trip time (ms)', fontsize=9)
            col += 1
        row += 1

    # Draw plot
    plt.subplots_adjust(hspace=0.2, wspace=0.3)
    # plt.xlabel('latency (ms)', fontsize=10)
    # plt.ylabel('count', fontsize=10)
    pp.savefig(fig)
    plt.close(fig)




def plot_aggregate(pp, name, root, total_data):
    """ Plot the aggregated network behavior
    Arguments:
        pp -- PdfPages object
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    # First, collect list of ALL data
    aggregate_data = []  
    for i in total_data.keys():
        for durations in total_data[i].values():
            aggregate_data.extend(durations)
            

    # Next, draw the graph
    fig, this_ax = plt.subplots(1, 1, figsize=(8.5, 5), sharex=True)
    fig.suptitle(name)
    sns.despine(left=True)
    
    this_ax.set_title("Aggregate data across all nodes", fontsize=9)
    this_ax.grid()
    sns.distplot(aggregate_data, kde=False, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
    stats = AnchoredText(
        generate_statistics(aggregate_data), 
        loc='upper right',  
        prop=dict(size=8),
        bbox_to_anchor=(1.1, 1),
        bbox_transform=this_ax.transAxes
    )
    this_ax.add_artist(stats)
    this_ax.set_xlabel('round trip time (ms)', fontsize=10)
    this_ax.set_ylabel('count', fontsize=10)
    pp.savefig(fig)
    plt.close(fig)

    # Next, draw the cdf graph
    fig, this_ax = plt.subplots(1, 1, figsize=(8.5, 5), sharex=True)
    fig.suptitle(name)
    
    cdf, bins = raw_data_to_cdf(aggregate_data)

    this_ax.plot(cdf, bins)
    this_ax.set_title("Aggregate cdf across all nodes", fontsize=9)
    this_ax.grid()
    this_ax.set_xlabel('cumulative probability', fontsize=10)
    this_ax.set_ylabel('round trip time (ms)', fontsize=10)

    # Draw plot
    # plt.subplots_adjust(hspace=0.2, wspace=0.3)
    pp.savefig(fig)
    plt.close(fig)


def plot_individuals(pp, name, root, total_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        pp -- PdfPages object
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    fig, axes = plt.subplots(len(NODES), len(NODES), figsize=(4*len(NODES), 4*len(NODES)), sharex=True)
    fig.suptitle(name)
    sns.despine(left=True)
    
    row = 0
    for i in total_data.keys():
        col = 0
        for j in total_data[i].keys():
            i_j_data = total_data[i][j]
            this_ax = axes[row][col]

            this_ax.set_title("node%d -> node%d" %(i, j), fontsize=9)
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
            col += 1
        row += 1

    # Draw plot
    plt.subplots_adjust(hspace=0.2, wspace=0.3)
    # plt.xlabel('latency (ms)', fontsize=10)
    # plt.ylabel('count', fontsize=10)
    pp.savefig(fig)
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
    res.append(f"n = {'{:,}'.format(len(input))}")
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