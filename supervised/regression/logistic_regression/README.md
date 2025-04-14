# Logistic Regression

## Overview
Logistic Regression is a statistical method used for binary classification problems. It predicts the probability of an outcome belonging to one of two categories based on one or more input features.

## Key Concepts
- **Sigmoid Function**: Converts linear predictions into probabilities between 0 and 1.
- **Decision Boundary**: A threshold used to classify probabilities into binary outcomes.
- **Cost Function**: Log-loss is used to optimize the model during training.
- **Weights and Bias**: Parameters learned during training to minimize the cost function.

## Algorithm
1. Initialize weights and bias.
2. Compute the linear combination of input features and weights.
3. Apply the sigmoid function to obtain probabilities.
4. Calculate the log-loss cost function.
5. Use gradient descent to update weights and bias iteratively.
6. Repeat until convergence or a stopping criterion is met.

## Advantages
- Simple to implement and interpret.
- Works well for linearly separable data.
- Outputs probabilities, which can be useful for decision-making.

## Disadvantages
- Assumes a linear relationship between input features and the log-odds of the outcome.
- Struggles with non-linear data unless features are transformed.
- Sensitive to outliers.

## Use Cases
- Spam email classification.
- Predicting customer churn.
- Medical diagnosis (e.g., disease presence/absence).
- Credit risk assessment.
