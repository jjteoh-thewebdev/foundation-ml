# Hierarchical Clustering

## Overview
Hierarchical clustering is a method of cluster analysis that seeks to build a hierarchy of clusters. It is commonly used in unsupervised machine learning to group similar data points into clusters based on their features.

## Key Concepts
- **Dendrogram**: A tree-like diagram that records the sequences of merges or splits in hierarchical clustering.
- **Agglomerative Clustering**: A bottom-up approach where each data point starts as its own cluster, and pairs of clusters are merged as one moves up the hierarchy.
- **Divisive Clustering**: A top-down approach where all data points start in one cluster, and splits are performed recursively.

## Algorithm
### Agglomerative Clustering Steps:
1. Start with each data point as its own cluster.
2. Compute the distance between all pairs of clusters.
3. Merge the two closest clusters.
4. Repeat steps 2 and 3 until all points are in a single cluster or a stopping criterion is met.

### Divisive Clustering Steps:
1. Start with all data points in one cluster.
2. Split the cluster into two based on a criterion (e.g., distance or similarity).
3. Repeat step 2 recursively for each resulting cluster until each data point is its own cluster or a stopping criterion is met.

## Advantages
- Does not require specifying the number of clusters in advance.
- Produces a dendrogram, which provides a visual representation of the clustering process.
- Can capture nested clusters and hierarchical relationships.

## Disadvantages
- Computationally expensive for large datasets.
- Sensitive to noise and outliers.
- The choice of distance metric and linkage method can significantly affect results.

## Use Cases
- Gene expression analysis in bioinformatics.
- Document clustering in natural language processing.
- Market segmentation in business analytics.
- Image segmentation in computer vision.
