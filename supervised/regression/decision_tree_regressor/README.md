# Decision Tree Regressor

This directory contains an implementation of a Decision Tree Regressor, a supervised learning algorithm used for regression tasks. Decision trees partition the data into subsets based on feature values, creating a tree-like model of decisions.

## Overview

The Decision Tree Regressor works by recursively splitting the dataset into smaller subsets based on the feature that minimizes the error (e.g., mean squared error). The process continues until a stopping criterion is met, such as a maximum depth or a minimum number of samples per leaf.

## Key Features

- Handles both numerical and categorical data.
- Non-parametric and does not assume a specific data distribution.
- Can model non-linear relationships effectively.

## Advantages

- Easy to interpret and visualize.
- Requires little data preprocessing (e.g., no need for feature scaling).
- Can handle both numerical and categorical data.
- Robust to outliers and irrelevant features.

## Disadvantages

- Prone to overfitting, especially with deep trees.
- Sensitive to small changes in the data, which can lead to different splits.
- Can be biased towards features with more levels (categorical data).
- May not perform well on very large datasets compared to other algorithms.

## References

- [Scikit-learn Decision Tree Regressor Documentation](https://scikit-learn.org/stable/modules/generated/sklearn.tree.DecisionTreeRegressor.html)
- [Introduction to Decision Trees](https://en.wikipedia.org/wiki/Decision_tree_learning)
