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
from matplotlib import dates as mdates
from datetime import datetime
import seaborn as sns

from conv import *

NODES = [1, 2, 3, 4]
PAYLOADS = [16]
NODES.sort()
PAYLOADS.sort()
HOSTS = "aws-hosts_single-region.csv"


START = datetime.fromisoformat("2021-10-26 00:00:01")
END = datetime.fromisoformat("2021-12-15 04:00:00")

SAMPLE_EVERY = 1

def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    for payload in PAYLOADS:
        print("\nAnalyzing all network data (payload %d): %s" %(payload, exp_dir))
        total_payload_data = dict()  # {2D map} The timings for node i to node j for this payload
        for root, dirs, files in os.walk("%s" %exp_dir):
            if "payload%d" %payload not in root:
                # Only look at data for the specific payload
                continue    
            files = [f for f in files if not f[0] == '.']  # ignore hidden files
            if dirs == []:
                # This is a leaf directory containing trial csv files
                host_names = get_hosts()
                for src in host_names:
                    if src not in total_payload_data:
                        total_payload_data[src] = dict()
                    for target in host_names:
                        csv_name = "node_%s-%s.csv" %(src, target)
                        # print(csv_name)
                        if target not in total_payload_data[src]:
                            total_payload_data[src][target] = []
                        res = analyze_csv("%s/%s" %(root, csv_name))
                        total_payload_data[src][target].extend(res)
        
        # Save the data
        with open("%s/total_payload%d_data.pickle" %(exp_dir, payload), 'wb') as handle:
            pickle.dump(total_payload_data, handle)
        with open("%s/total_payload%d_data.pickle" %(exp_dir, payload), 'rb') as handle:
            total_payload_data = pickle.load(handle)
        
        # Draw the pictures
        print("\tDrawing payload %d" %payload)
        plot_figures("rtt_payload%d" %payload, exp_dir, total_payload_data)
    print("Done")


def get_hosts():
    """Returns a list of host names for the cluseter"""
    res = []
    with open(HOSTS, 'r') as hosts:
        csvreader = csv.reader(hosts, delimiter=',',)
        next(csvreader)  # skip the header
        for row in csvreader:
            if row != []:
                res.append(row[1])
    return res


def plot_figures(name, root, total_data):
    """ Plot all network figures
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    # assert len(total_data) == len(NODES) and len(total_data[NODES[0]]) == len(NODES)
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        # plot_individuals(pp, name, root, total_data)
        plot_cdf(pp, name, root, total_data)
        plot_time_series(pp, name, root, total_data)


def plot_cdf(pp, name, root, total_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        pp -- PdfPages object
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    fig, axes = plt.subplots(len(NODES), len(NODES), figsize=(4*len(NODES), 4*len(NODES)), sharex=True)
    # fig, axes = plt.subplots(2, 2, figsize=(5*2, 5*2))
    fig.suptitle(name)
    sns.despine(left=True)
    
    row = 0
    for i in total_data.keys():
        col = 0
        for j in total_data[i].keys():
            i_j_data = [p[0]/2 for p in total_data[i][j] if not p[2]]
            this_ax = axes[row][col]
            this_ax.set_title("node%s -> node%s" %(i, j), fontsize=9)
            this_ax.grid()

            cdf, bins = raw_data_to_cdf(i_j_data)

            this_ax.plot(cdf, bins)

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


def plot_time_series(pp, name, root, total_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        pp -- PdfPages object
        name -- name of this figure
        root -- directory to save this figure
        total_data {2D map} -- total_data[i][j] is the timings for node i to node j
    """
    fig, axes = plt.subplots(len(NODES)*len(NODES), 1, figsize=(4*len(NODES), 4*len(NODES)*len(NODES)), sharex=True)
    fig.suptitle(name)
    sns.despine(left=True)
    
    row = 0
    for i in total_data.keys():
        for j in total_data[i].keys():
            y_vals = [p[0] for p in total_data[i][j]]
            colors = ['red' if p[2] else 'blue' for p in total_data[i][j]]
            x_vals = mdates.date2num([p[1] for p in total_data[i][j]])
            this_ax = axes[row]

            this_ax.set_title("%s -> %s" %(i, j), fontsize=9)
            this_ax.grid()
            this_ax.scatter(x_vals, y_vals, c=colors, marker='.', s=3)

            # Major ticks every few hours.
            fmt_hrs = mdates.HourLocator(interval=6)
            this_ax.xaxis.set_major_locator(fmt_hrs)

            # Minor ticks every hour.
            fmt_hr = mdates.HourLocator(interval=1)
            this_ax.xaxis.set_minor_locator(fmt_hr)

            # See https://docs.python.org/3/library/datetime.html#strftime-strptime-behavior for formatting str
            this_ax.xaxis.set_major_formatter(mdates.DateFormatter('%m-%d %H:%M')) 

            # Statistics
            stats = AnchoredText(
                generate_statistics([y for y in y_vals if y != 0]), 
                loc='upper right',  
                prop=dict(size=8),
                bbox_to_anchor=(1.1, 1),
                bbox_transform=this_ax.transAxes
            )
            this_ax.add_artist(stats)
            if i == len(NODES) - 1:
                this_ax.set_xlabel('round trip time (ms)', fontsize=9)
            row += 1

    fig.autofmt_xdate()
    # Draw plot
    plt.subplots_adjust(hspace=0.2, wspace=0.3)
    # plt.xlabel('latency (ms)', fontsize=10)
    # plt.ylabel('count', fontsize=10)
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
            i_j_data = [p[0]/2 for p in total_data[i][j]]  # each elem of total_data is a pair
            this_ax = axes[row][col]

            this_ax.set_title("%s -> %s" %(i, j), fontsize=9)
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
    plt.xlabel('latency (ms)', fontsize=10)
    plt.ylabel('count', fontsize=10)
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
    durations_milli = []
    with open(filepath, 'r') as node1:
        csvreader = csv.reader(node1, delimiter=',',)
        i = 0
        count = 0  # count of valid, non-timeout entries
        sum = 0
        for row in csvreader:
            i += 1
            # Only look at every SAMPLE_EVERY row
            if (SAMPLE_EVERY <= 1 or i % SAMPLE_EVERY == 0) and  row != []:
                if "TIMEOUT" in row[1]:
                    timestamp = parse_go_timestamp(row[5])
                    if START < timestamp and timestamp < END:
                        durations_milli.append((0,timestamp, True))
                else:
                    start_time = int(row[3])
                    end_time = int(row[4])
                    timestamp = parse_go_timestamp(row[5])
                    dur = (end_time - start_time)/1_000_000.0  # duration in milliseconds
                    # Weed out anomalous data
                    if count > 0 and dur < sum / count / 10:
                        continue
                    # Consider desired time period
                    if START < timestamp and timestamp < END:
                        durations_milli.append((dur,timestamp, False))
                        count += 1
                        sum += dur
    return durations_milli


def parse_go_timestamp(time_str):
    """ Return Python datetime object representing time_str"""
    res = time_str[:23].strip()  # leave the seconds to 3dp
    return datetime.fromisoformat(res)

if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 3 or len(sys.argv) > 3:
        print("Error: Wrong number of arguments")
        exit(1)
    HOSTS = sys.argv[1]
    NODES = range(len(get_hosts()))
    exp_dir =sys.argv[2]
    main(exp_dir)