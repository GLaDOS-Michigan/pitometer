import sys
import csv
import matplotlib.pyplot as plt
import seaborn as sns


def main(argv):
    total_data = []
    for filepath in argv:
        total_data.append(analyze_file(filepath))
    f, axes = plt.subplots(len(total_data), 1, figsize=(7, 7), sharex=True)
    sns.despine(left=True)
    i = 0
    for durations_nano in total_data:
        filename = argv[i].split('/')[-1]
        this_ax = axes[i]
        this_ax.set_title(filename)
        durations_milli = list(map(lambda x: x/1_000_000, durations_nano))
        sns.distplot(durations_milli, kde=False, ax= this_ax)
        i += 1
    plt.tight_layout()
    plt.show()


def analyze_file(filepath):
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
    return durations_nano


if __name__ == "__main__":
    # positional arguments are <inputfile>
    main(sys.argv[1:])