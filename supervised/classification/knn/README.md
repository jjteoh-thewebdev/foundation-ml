# K-Nearest Neighbors (KNN)

## Description
K-Nearest Neighbors is a simple, non-parametric algorithm used for classification and regression. It predicts the output based on the majority class or average of the nearest neighbors.

## Key Concepts
- **Distance Metrics**: Measures like Euclidean(√∑(xᵢ - yᵢ)²), Manhattan(∑|xᵢ - yᵢ|), or Minkowski distance to find nearest neighbors.
- **K Value**: The number of neighbors considered for prediction.
- **Lazy Learning**: No explicit training phase; predictions are made during inference.

## Algo
1. Compute distance between `X_test` and `X_train`
2. Sort the distance
3. Pick the top `k` nearest neighbors
4. Return majority class among those neighbors.

## Common Applications
- Pattern recognition
- Recommender systems
- Image classification
- Anomaly detection

## Advantages
- Simple to implement and understand.
- No assumptions about data distribution.
- Effective for small datasets.

## Disadvantages
- Computationally expensive for large datasets.
- Sensitive to the choice of K and distance metric.
- Struggles with high-dimensional data.
