# Experiments

This directory contains the scripts to run experiments, and stores the experiment results.
The scripts in **bold** are end-to-end push-button scripts for running and processing 
entire experiments

- data/
    * Directory that stores experimental results
- analyze_network.py
    * Python script to plot results from a network experiment
- analyze_rsl.py
    * Python script to plot results from a rsl experiment
- analyze_rsl_throughput.py
    * Python script to plot results from a rsl latency-throughput experiment
- analyze_toylock.py
    * Python script to plot results from a toylock experiment
- backup    
    * Back up experiment data to the pito-data repository, and other remote machines
- build_to_hosts
    * Script that builds rsl and toylock locally, and then copies the executables to the specified hosts
- clean_builds
    * Cleaning script that deletes executables and miscellaneous temp files.
- conv.py
    * Python library for adding probability distributions
- git_pull_all_hosts
    * Script that does a 'git pull' on all hosts for a specified branch for
- hosts.csv
    * A map from Skynode machine ID to their IP address
- kill_all_experiments
    * Kills all experiments running on all hosts using `kill -9`
- plot_cosntants.py
    * Python library containing constants for drawing graphs
- **process_network_results**
    * Converts network log file into separate csv files for each network client. It places each resulting csv in the directory of its corresponding log file.
    * Then, it calls analyze_network.py to generate graphs
- **process_rsl_results**
    * Converts rsl log file into separate csv files for each toylock stopwatch. It places each resulting csv in the directory of its corresponding log file.
    * Then, it calls analyze_rsl.py to generate graphs
- **process_toylock_results**
    * Converts toylock log file into separate csv files for each toylock stopwatch. It places each resulting csv in the directory of its corresponding log file.
    * Then, it calls analyze_toylock.py to generate graphs
- **run_network_experiment**
    * Start a network experiment, which consists of a set of trials
- **run_rsl_experiment**
    * Start an rsl experiment, which consists of a set of trials
- **run_rsl_throughput_experiment**
    * Start an rsl experiment that is designed to measure throughput-latency data
- **run_toylock_experiment**
    * Start a toylock experiment, which consists of a set of trials
- start_network_trial
    * Start a single network trial with remote hosts
- start_rsl_trial
    * Start a single rsl trial with the specified remote hosts
- start_rsl_trial-local
    * Start a single rsl trial with all nodes running locally, using the loopback IP addresses
- start_toylock_trial
    * Start a single toylock trial with the specified remote hosts
- start_toylock_trial-local
    * Start a single toylock trial with all nodes running locally, using the loopback IP addresses

    