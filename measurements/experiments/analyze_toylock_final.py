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
DELAYS = [0, 200, 1_000, 5_000, 25_000]  # units of microseconds

# Global plotting params
# Fonts
plt.rc('font', family='serif') 
plt.rc('font', size=12)              # controls default text sizes
plt.rc('axes', titlesize=12)        # fontsize of the axes title
plt.rc('axes', labelsize=12)        # fontsize of the x and y labels
plt.rc('xtick', labelsize=12)       # fontsize of the tick labels
plt.rc('ytick', labelsize=12)       # fontsize of the tick labels
plt.rc('legend', fontsize=12)       # legend fontsize
plt.rc('figure', titlesize=15)      # fontsize of the figure title

# Lines
plt.rc('lines', linewidth=1)


def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)

    # each data is dict of (size -> delay -> node -> [ durs ])
    with open("%s/%s" %(exp_dir, 'total_grant_data.pickle'), 'rb') as handle:
        total_grant_data = pickle.load(handle)
    with open("%s/%s" %(exp_dir, 'total_accept_data.pickle'), 'rb') as handle:
        total_accept_data = pickle.load(handle)
    with open("%s/%s" %(exp_dir, 'total_round_data.pickle'), 'rb') as handle:
        total_round_data = pickle.load(handle)

    print("\nPlotting graphs for experiment %s" %exp_dir)

    # Plot Rounds
    plot_micro_2_size_fidelity("Micro-benchmark 2", exp_dir, total_round_data)
    print("Done")


def plot_micro_1_distr_fidelity(name, root, total_round_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
    """
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        for delay in DELAYS:
            x_vals_ring_size = sorted(list(total_round_data.keys()))
            y_vals_ninety_nine_point_nine_percentiles = []
            y_vals_fifty_percentiles = []
            for ring_size in x_vals_ring_size:
                leader_node = sorted(list(total_round_data[ring_size][delay].keys()))[0]
                y_vals_ninety_nine_point_nine_percentiles.append(np.percentile(total_round_data[ring_size][delay][leader_node], 99.9))
                y_vals_fifty_percentiles.append(np.percentile(total_round_data[ring_size][delay][leader_node], 50))
            fig, this_ax = plt.subplots(1, 1, figsize=(8.5, 5), sharex=False)
            plot_micro_2_size_fidelity_ax(this_ax, "delay %.1f" %(delay/1000.0), x_vals_ring_size, y_vals_fifty_percentiles, y_vals_ninety_nine_point_nine_percentiles)
            pp.savefig(fig)
            plt.close(fig)



def plot_micro_2_size_fidelity(name, root, total_round_data):
    """ Plot a figure where each subfigure is from an element in total_data
    Arguments:
        name -- name of this figure
        root -- directory to save this figure
        total_data -- dict of (size -> delay -> node -> [ durs ])
    """
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        for delay in DELAYS:
            x_vals_ring_size = sorted(list(total_round_data.keys()))
            y_vals_ninety_nine_point_nine_percentiles = []
            y_vals_fifty_percentiles = []
            for ring_size in x_vals_ring_size:
                leader_node = sorted(list(total_round_data[ring_size][delay].keys()))[0]
                y_vals_ninety_nine_point_nine_percentiles.append(np.percentile(total_round_data[ring_size][delay][leader_node], 99.9))
                y_vals_fifty_percentiles.append(np.percentile(total_round_data[ring_size][delay][leader_node], 50))
            fig, this_ax = plt.subplots(1, 1, figsize=(8.5, 5), sharex=False)
            plot_micro_2_size_fidelity_ax(this_ax, "delay %.1f" %(delay/1000.0), x_vals_ring_size, y_vals_fifty_percentiles, y_vals_ninety_nine_point_nine_percentiles)
            pp.savefig(fig)
            plt.close(fig)

def plot_micro_2_size_fidelity_ax(this_ax, title, x_vals_ring_size, y_vals_fifty_percentiles, y_vals_ninety_nine_point_nine_percentiles):
    this_ax.set_title(title)
    this_ax.set_xlabel("ring size")
    this_ax.set_ylabel("latency (ms)")
    this_ax.plot(x_vals_ring_size, y_vals_fifty_percentiles, label='observed 50%%-ile', marker='x', color='blue')
    this_ax.plot(x_vals_ring_size, y_vals_ninety_nine_point_nine_percentiles, label='observed 99.9%%-ile', marker='x', color='red')


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)