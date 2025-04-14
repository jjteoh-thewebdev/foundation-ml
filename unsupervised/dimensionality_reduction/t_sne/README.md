# t-SNE (t-Distributed Stochastic Neighbor Embedding)

## Overview
t-SNE is a dimensionality reduction technique primarily used for visualizing high-dimensional data in a lower-dimensional space, typically 2D or 3D. It is particularly effective for exploring patterns and clusters in datasets.

## Key Concepts
- **High-dimensional data**: Data with many features or dimensions.
- **Dimensionality reduction**: The process of reducing the number of features while preserving meaningful structure.
- **Similarity preservation**: t-SNE aims to maintain the relative distances between points in the high-dimensional and low-dimensional spaces.

## Algorithm
1. Compute pairwise similarities between data points in the high-dimensional space.
2. Define a similar probability distribution in the low-dimensional space.
3. Minimize the Kullback-Leibler (KL) divergence between the two distributions using gradient descent.

## Advantages
- Captures non-linear relationships in data.
- Effective for visualizing clusters and patterns.
- Works well with high-dimensional datasets.

## Disadvantages
- Computationally expensive for large datasets.
- Results can vary depending on hyperparameters (e.g., perplexity).
- Not suitable for all types of data.

## Use Cases
- Visualizing high-dimensional datasets in machine learning.
- Exploring clusters in biological data (e.g., gene expression).
- Analyzing customer segmentation in marketing.
- Understanding embeddings in natural language processing.
