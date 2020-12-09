import numpy as np
from scipy import signal

def get_mean(pdf, bins):
    sum = 0
    for i in range(len(pdf)):
        sum += pdf[i] * float(bins[i])
    return sum

def get_percentile(cdf, bins, percentile):
    assert percentile >= 0 and percentile <= 100
    for i in range(len(cdf)):
        if cdf[i] > percentile/100.0:
            return bins[i]
    return bins[-1]

def raw_data_to_cdf(data):
    binsize = 1e-4
    pdf, bins = raw_data_to_pdf(data, binsize)
    cdf, bins = pdf_to_cdf(pdf), bins.tolist()
    return [0] + cdf, [bins[0]] + bins

def raw_data_to_pdf(data, binsize):
    bincount = int((max(data) - min(data))/binsize)
    bins = np.linspace(min(data), max(data), bincount)
    pdf, bins = np.histogram(data, bins=bins)
    return pdf, bins

def pdf_to_cdf(pdf):
    cdf = np.cumsum(pdf/pdf.sum()).tolist()
    return cdf

def cdf_to_pdf(cdf):
    dx = 1.0 / len(cdf)
    pdf = [0]
    for i in range(len(cdf)-1):
        pdf.append((cdf[i+1]-cdf[i])/dx)
    
    return np.array(pdf), dx


def add_histograms(pdf1, pdf2, start1, start2, binsize1, binsize2):
    """
    pdf{j} should be a seq of numbers representing the histogram for the jth
    distribution. Assumes that the numbers in the sequence pdf{j} sum to 1
    start{j} is the starting of pmf1
    """
    conv_pdf = signal.fftconvolve(pdf1,pdf2,'full')
    conv_pdf = conv_pdf/float(conv_pdf.sum()) # This should be unnecessary, but
                                              # keeping it just in case pdf1 and pdf2 don't have a sum of 1
    binsize_out = binsize1 # + binsize2
    start_out = start1 + start2
    return conv_pdf, start_out, binsize_out