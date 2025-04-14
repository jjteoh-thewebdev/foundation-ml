# Principal Component Analysis (PCA)

## Overview
Principal Component Analysis (PCA) is a statistical technique used for dimensionality reduction. It transforms high-dimensional data into a lower-dimensional space while preserving as much variance as possible. PCA is widely used in machine learning, data analysis, and visualization.

## Key Concepts
- **Dimensionality Reduction**: Reducing the number of features in a dataset while retaining important information.
- **Principal Components**: New axes that maximize variance and are orthogonal to each other.
- **Variance**: Measure of the spread of data; PCA aims to retain the components with the highest variance.
- **Eigenvalues and Eigenvectors**: Used to compute the principal components from the covariance matrix.

## Algorithm
1. Standardize the dataset to have zero mean and unit variance.
2. Compute the covariance matrix of the data.
3. Calculate the eigenvalues and eigenvectors of the covariance matrix.
4. Sort the eigenvectors by their corresponding eigenvalues in descending order.
5. Select the top `k` eigenvectors to form the projection matrix.
6. Transform the original data into the lower-dimensional space using the projection matrix.

## Advantages
- Reduces computational complexity by lowering the number of features.
- Helps in visualizing high-dimensional data.
- Removes noise and redundant features.
- Improves model performance by reducing overfitting.

## Disadvantages
- Sensitive to scaling; requires data standardization.
- Assumes linear relationships between features.
- May lose interpretability of the original features.
- Not suitable for non-linear datasets.

## Use Cases
- Image compression and feature extraction.
- Preprocessing step for machine learning models.
- Exploratory data analysis and visualization.
- Reducing noise in datasets for better performance.
