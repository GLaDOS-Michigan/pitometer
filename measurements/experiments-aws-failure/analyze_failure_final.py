import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import statistics
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

THROW=1  # Ignore the first THROW requests in computing method latencies
CTHROW=1  # Ignore the THROW requests in computing client latencies

TRAIN_SET = "train"
TEST_SET = "test"
CD_VALUES = [0, 500, 1500, 2000, 2500, 3000]

START = datetime.fromisoformat("2021-11-14 00:00:01")
END = datetime.fromisoformat("2021-12-15 04:00:00")


def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)
    data = parse_csvs(exp_dir)

    print("\nPlotting graphs for experiment %s" %exp_dir)
    plot_graph(exp_dir, data)
    # print("Done")
    
    
def parse_csvs(exp_dir):
    data = {key: [] for key in CD_VALUES}
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
                data[cd].append(dur)
    return data  


def plot_graph(root, data):
    y_vals = [ np.mean(data[cd]) for cd in CD_VALUES ]
    errors = [ statistics.stdev(data[cd]) if len(data[cd]) > 1 else 0 for cd in CD_VALUES ]
    
    with PdfPages("%s/failure_graph.pdf" %root) as pp:
        fig, this_ax = plt.subplots(1, 1, figsize=(fig_width, fig_height), sharex=False)
        fig.subplots_adjust(right=0.96, bottom=0.18, left=0.15)
        plt.plot(CD_VALUES, y_vals, label='Observed behavior', color='navy', marker='o')
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



if __name__ == "__main__":
    # positional arguments <experiment_dir>
    if len(sys.argv) < 2 or len(sys.argv) > 2:
        print("Error: Script takes a single positional argument that is the directory containing the toylock trials")
        exit(1)
    exp_dir =sys.argv[1]
    main(exp_dir)
