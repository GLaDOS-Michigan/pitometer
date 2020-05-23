import sys
import os
import csv
import statistics
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnchoredText
from matplotlib.backends.backend_pdf import PdfPages
import seaborn as sns

BATCH_SIZES = [1, 32]

def main(exp_dir):
    exp_dir = os.path.abspath(exp_dir)
    print("\nAnalyzing data for experiment %s" %exp_dir)
    for batch_size in BATCH_SIZES:
        # total_data is dict (f -> threads -> trials -> [ (start, end, client_id)... ]
        total_batch_data = parse_batch_files(exp_dir, batch_size)
        name = "batch_%d_latency_throughput" %batch_size
        gen_latency_throughput_graphs(name, exp_dir, total_batch_data)


def parse_batch_files(exp_dir, batch_size):
    """Collects the total data for a batch

    Arguments:
        exp_dir {string} -- root directory of experiment
        batch_size {int} -- batch size of which data to collect
    Returns:
        dict (f -> threads -> trial_num -> [ (start, end, client_id)... ]
    """
    # total_data is dict (f -> threads -> trials -> [ (start, end, client_id)... ]
    total_batch_data = dict()
    batch_dir = "%s/batch_%d" %(exp_dir, batch_size)
    for root, _, files in os.walk(batch_dir):
        files = [f for f in files if not f[0] == '.']  # ignore hidden files
        if files != [] and 'trial' in root:
            # This is a leaf directory containing trial csv files
            print("\tAnalyzing trial %s" %root)
            f, threads, trial = parse_trial_params(root)

            # Populate the dict
            if f not in total_batch_data:
                total_batch_data[f] = dict()
            if threads not in total_batch_data[f]:
                total_batch_data[f][threads] = dict()
            for file in files:
                if 'client' in file:
                    client_log = "%s/%s" %(root, file)
                    total_batch_data[f][threads][trial] = parse_client_log(client_log)
    return total_batch_data


def gen_latency_throughput_graphs(name, root, total_batch_data):
    """Plots the latency throughput graph for this batch_size in one pdf file

    Arguments:
        name {string} -- name of this file
        root -- directory to save this figure
        total_batch_data {dict} -- (f -> threads -> trials -> [ (start, end, client_id)... ]
    """
    with PdfPages("%s/%s.pdf" %(root, name)) as pp:
        fig, axes = plt.subplots(len(total_batch_data), 1, figsize=(8.5, 11), sharex=True)
        row = 0
        for f in sorted(list(total_batch_data.keys())):
            latencies = compute_average_latencies(total_batch_data[f])
            throughputs = compute_average_throughputs(total_batch_data[f])
            # Test run, plot the latencies only
            plot_latency_throughput(axes[row], "f = %d" %(f), latencies, throughputs)
            row += 1
        pp.savefig(fig)
        plt.close(fig)
            


def plot_latency_throughput(this_ax, title, latencies, throughputs):
    this_ax.set_title(title)
    this_ax.grid()
    this_ax.plot(throughputs, latencies, marker='x')


def compute_average_throughputs(total_f_data):
    """Computes the average throughputs for each num_threads for this f

    Arguments:
        total_f_data {dict} -- (threads -> trials -> [ (start, end, client_id)... ]
    returns
        {list} -- throughputs list
    """
    res = []
    for num_threads in sorted(list(total_f_data.keys())):
        avg_trial_throughputs = []  # list of average tp for each trial
        for trial in sorted(list(total_f_data[num_threads].keys())): 
            trial_data = total_f_data[num_threads][trial]

            trial_start_time = min([start for (start, end, client_id) in trial_data])
            trial_end_time = max([end for (start, end, client_id) in trial_data])
            num_requests = len(trial_data)
            avg_trial_throughputs.append(float(num_requests)/(trial_end_time-trial_start_time)*1000.0)
        res.append(np.mean(avg_trial_throughputs))
    return res


def compute_average_latencies(total_f_data):
    """Computes the average latencies for each num_threads for this f

    Arguments:
        total_f_data {dict} -- (threads -> trials -> [ (start, end, client_id)... ]
    returns
        {list} -- latencies list
    """
    res = []
    for num_threads in sorted(list(total_f_data.keys())):
        avg_trial_latencies = []  # list of average latency for each trial
        for trial in sorted(list(total_f_data[num_threads].keys())): 
            trial_latencies = []  # list of all latency for this trial
            trial_data = total_f_data[num_threads][trial]
            for req in trial_data:
                trial_latencies.append(req[1] - req[0])
            avg_trial_latencies.append(np.mean(trial_latencies))
        res.append(np.mean(avg_trial_latencies))
    return res


def parse_trial_params(path):
    """Returns the trial params given the path string
    Arguments:
        path {string} -- path to this trial
    Returns:
        {int} -- f
        {int} -- num theads
        {int} -- trial number
    """
    segments = path.split('/')
    f = int(segments[-3].split('_')[1])
    num_threads = int(segments[-2].split('_')[1])
    trial_num = int(segments[-1].split('ial')[1])
    return f, num_threads, trial_num


def parse_client_log(client_log):
    """Parses the client log into the format
    [ (start, end, client_id)... ]
    Arguments:
        client_log {string} -- path to a client log
    Returns:
        {list} -- [ (start, end, client_id)... ]
    """
    res = []
    with open(client_log, 'r') as client:
        csvreader = csv.reader(client, delimiter=' ')
        for row in csvreader:
            req_start = int(row[1])
            req_end = int(row[2])
            client_id = int(row[3])
            res.append((req_start, req_end, client_id))
    # Ignore the first and last 10%
    num_reqs = len(res)
    truncated_res = []
    for i in range(num_reqs):
        if i > num_reqs * 0.1 and i < num_reqs * 0.9:
            truncated_res.append(res[i])
    return truncated_res


if __name__ == "__main__":
    # positional arguments <experiment_dir>
    exp_dir =sys.argv[1]
    main(exp_dir)