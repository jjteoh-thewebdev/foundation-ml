# Gradient Boosting

## Overview

Gradient boosting is a supervised learning algorithm used for classification and regression tasks. It builds an ensemble of weak learners, typically decision trees, in a stage-wise fashion. Each tree is trained to correct the errors made by the previous trees. The final prediction is made by combining the predictions of all trees in the ensemble.

## Key Concepts

*   **Ensemble Learning:** Combining multiple models to improve overall performance.
*   **Weak Learners:** Models that perform slightly better than random guessing.
*   **Boosting:** Sequentially training models, where each model focuses on the mistakes of the previous ones.
*   **Gradient Descent:** Used to minimize the loss function by iteratively adjusting the model's parameters.

## Algorithm

1.  Initialize the ensemble with a simple model (e.g., a constant value).
2.  For each iteration:
    *   Compute the negative gradient of the loss function with respect to the current ensemble prediction. This gradient represents the "errors" or "residuals" that the current ensemble is making.
    *   Train a weak learner (e.g., a decision tree) to predict these negative gradients.
    *   Update the ensemble by adding the predictions of the weak learner, scaled by a learning rate (also known as shrinkage). The learning rate controls the contribution of each tree to the ensemble and helps prevent overfitting.
3.  Repeat until a stopping criterion is met (e.g., a maximum number of iterations or a minimum improvement in performance).

## Advantages

*   High accuracy and predictive power.
*   Handles missing data and mixed data types.
*   Feature importance estimation.

## Disadvantages

*   Sensitive to noisy data and outliers.
*   Can be computationally expensive and memory-intensive.
*   Prone to overfitting if not properly tuned.

## Use Cases

*   Fraud detection
*   Risk assessment
*   Image classification
*   Natural language processing
