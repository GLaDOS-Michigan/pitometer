import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
import seaborn as sns


def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)
    for root, _, files in os.walk(exp_dir):
        files = [f for f in files if not f[0] == '.']  # ignore hidden files
        if files != []:
            # This is a leaf directory containing trial csv files
            print("\tAnalyzing trial %s" %root)
            
            total_grant_data = []
            total_accept_data = []
            total_round_data = []
            grant_titles = []
            accept_titles = []
            round_titles = []

            for f in files:
                file_name, file_extension = os.path.splitext(f)
                if file_extension == '.csv':
                    if "grant" in file_name:
                        total_grant_data.append(analyze_grant_or_accept_csv("%s/%s" %(root, f)))
                        total_round_data.append(analyze_round_csv("%s/%s" %(root, f)))
                        grant_titles.append(file_name)
                        round_titles.append(file_name)
                    if "accept" in file_name:
                        total_accept_data.append(analyze_grant_or_accept_csv("%s/%s" %(root, f)))
                        accept_titles.append(file_name)

            # Plot Grant
            plot_figures("nodeGrant", root, total_grant_data, grant_titles)

            # Plot Accept
            plot_figures("nodeAccept", root, total_accept_data, accept_titles)

            # Plot Rounds
            plot_figures("rounds", root, total_round_data, round_titles)
    print("Done")


def plot_figures(name, root, total_data, titles):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- list of lists of data
        titles -- list of titles for each subfigure
    """

    if len(total_data) <= 3:
        fig, axes = plt.subplots(len(total_data), 1, figsize=(7, 3*max(2, len(total_data))), sharex=True)
    else:
        fig, axes = plt.subplots(3, 2, figsize=(14, 9), sharex=True)
    fig.suptitle(name)
    sns.despine(left=True)
    i = 0
    for durations_milli in total_data:
        if len(total_data) == 1:
            this_ax = axes
        else:
            if len(total_data) <= 3:
                this_ax = axes[i]
            else:
                this_ax = axes[i%3][i//3]
        # Plot the subfigure
        this_ax.set_title(titles[i], fontsize=9)
        this_ax.grid()
        sns.distplot(durations_milli, kde=False, ax=this_ax, hist_kws=dict(edgecolor="k", linewidth=0.1))
        stats = AnchoredText(
            generate_statistics(durations_milli), 
            loc='upper right',  
            prop=dict(size=8),
            bbox_to_anchor=(1.1, 1),
            bbox_transform=this_ax.transAxes
        )
        this_ax.add_artist(stats)
        # this_ax.set_xlabel('time (ms)', fontsize=9)
        # this_ax.set_ylabel('count', fontsize=9)
        # this_ax.set_xlim(0, x_max)
        # this_ax.set_ylim(0, 1)
        i += 1

    # Display some global figures
    global_data = []
    for row in total_data:
        global_data.extend(row)

    global_stats =  "Global statistics:\n%s" %generate_statistics(global_data)
    plt.figtext(0.86, 0.11, global_stats,
            fontsize=8,
            bbox=dict(boxstyle="round", facecolor='#D8D8D8',
            ec="0.5", pad=0.3, alpha=1))
    
    # plt.tight_layout()
    plt.subplots_adjust(hspace=0.2)
    plt.xlabel('time (ms)', fontsize=10)
    plt.ylabel('count', fontsize=10)
    plt.savefig("%s/%s.pdf" %(root, name))
    plt.close(fig)


def generate_statistics(input):
    """
    Generates a string containing some statistics for the input
    Arguments:
        input -- list of numbers
    """
    res = []
    res.append("n = %d" %len(input))
    res.append("μ = %.3f" %statistics.mean(input))
    res.append("σ = %.4f" %statistics.stdev(input))
    res.append("99.9%% = %.3f" %np.percentile(input, 99.9))
    res.append("")
    res.append("max = %.3f" %np.max(input))
    res.append("min = %.3f" %np.min(input))
    return "\n".join(res)


def analyze_grant_or_accept_csv(filepath):
    durations_nano = []
    with open(filepath, 'r') as node1:
        csvreader = csv.reader(node1, delimiter=',',)
        for row in csvreader:
            if row != [] or int(row[0]) >= 0:
                event_type = row[1]
                if event_type == 'Start':
                    prevStart = int(row[3])
                if event_type == 'End':
                    dur = int(row[3]) - prevStart  # duration in nanoseconds
                    durations_nano.append(dur)
    durations_milli = list(map(lambda x: x/1_000_000.0, durations_nano))
    return durations_milli

def analyze_round_csv(filepath):
    """ Computes the time taken for each round from nodeGrant csv data.
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
                event_type = row[1]
                if event_type == 'End':
                    round_num = int(row[0])
                    if round_num == 0:
                        round_start_time = int(row[3])
                        continue
                    dur = int(row[3]) - round_start_time  # duration in nanoseconds
                    durations_nano.append(dur)
                    round_start_time = int(row[3])
    durations_milli = list(map(lambda x: x/1_000_000, durations_nano))
    return durations_milli


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)