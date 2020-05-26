import matplotlib.pyplot as plt

# Global plotting params
# Fonts
base_size = 6
plt.rc('font', family='sans-serif') 
plt.rc('font', size=base_size)              # controls default text sizes
plt.rc('axes', titlesize=base_size)        # fontsize of the axes title
plt.rc('axes', labelsize=base_size)        # fontsize of the x and y labels
plt.rc('xtick', labelsize=base_size)       # fontsize of the tick labels
plt.rc('ytick', labelsize=base_size)       # fontsize of the tick labels
plt.rc('legend', fontsize=base_size)       # legend fontsize
# plt.rc('figure', titlesize=12)      # fontsize of the figure title

# Graphics
plt.rc('lines', linewidth=0.5)
plt.rc('lines', markersize=2)
plt.rc('legend', frameon=False)        # No frame around legend

# Figure size
fig_width = 3
fig_height = 2.5