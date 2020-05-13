# Experiments

This directory contains the scripts to run experiments, and stores the experiment results.

- data/
    * Directory that stores experimental results
- analyze_toylock.py
    * Python script to plot results from a toylock experiment
- build_to_hosts
    * Script that builds rsl and toylock locally, and then copies the executables to the specified hosts
- clean_builds
    * Cleaning script that deletes executables and miscellaneous temp files.
- hosts.csv
    * A map from Skynode machine ID to their IP address
- run_toylock_experiment
    * Start a toylock experiment, which consists of a set of trials
- start_toylock_trial-local
    * Start a single toylock trial with the specified remote hosts
- start_toylock_trial-local
    * Start a single toylock trial with all nodes running locally, using the loopback IP addresses
- toylock_log_to_csv
    * Converts toylock log file into separate csv files for each toylock stopwatch. It places each resulting csv in the directory of its corresponding log file.

    