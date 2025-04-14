# DBSCAN Clustering

## Overview
DBSCAN (Density-Based Spatial Clustering of Applications with Noise) is a popular clustering algorithm used to identify clusters of varying shapes and sizes in a dataset. It groups points that are closely packed together while marking points in low-density regions as outliers.

## Key Concepts
- **Epsilon (ε):** The maximum distance between two points to be considered neighbors.
- **MinPts:** The minimum number of points required to form a dense region (a cluster).
- **Core Point:** A point with at least `MinPts` neighbors within a distance of `ε`.
- **Border Point:** A point that is not a core point but is within `ε` distance of a core point.
- **Noise Point:** A point that does not belong to any cluster.

## Algorithm
1. Select an unvisited point and check its neighbors within distance `ε`.
2. If the point has at least `MinPts` neighbors, it becomes a core point, and a new cluster is formed.
3. Expand the cluster by recursively visiting all neighbors of core points.
4. Mark points that do not belong to any cluster as noise.
5. Repeat until all points are visited.

## Advantages
- Can find clusters of arbitrary shapes.
- Handles noise and outliers effectively.
- Does not require specifying the number of clusters in advance.

## Disadvantages
- Sensitive to the choice of `ε` and `MinPts`.
- Struggles with datasets of varying density.
- Computationally expensive for large datasets.

## Use Cases
- Geospatial data analysis (e.g., identifying regions of interest).
- Image segmentation.
- Anomaly detection in datasets.
- Customer segmentation in marketing.
