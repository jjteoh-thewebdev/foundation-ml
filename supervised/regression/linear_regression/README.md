# Linear Regression

## Overview
Linear regression is a fundamental statistical and machine learning technique used to model the relationship between a dependent variable and one or more independent variables. It assumes a linear relationship between the variables and is widely used for predictive analysis.

## Key Concepts
- **Dependent Variable (Target):** The variable we aim to predict or explain.
- **Independent Variables (Features):** The variables used to predict the dependent variable.
- **Line of Best Fit:** A straight line that minimizes the difference between predicted and actual values.
- **Residuals:** The differences between observed and predicted values.

## Algorithm
1. Initialize the parameters (weights and bias).
2. Compute the predicted values using the linear equation:  
    \[
    y = w_1x_1 + w_2x_2 + \dots + w_nx_n + b
    \]
3. Calculate the loss using a cost function, typically Mean Squared Error (MSE):  
    \[
    MSE = \frac{1}{n} \sum_{i=1}^{n} (y_i - \hat{y}_i)^2
    \]
4. Optimize the parameters using gradient descent or another optimization algorithm to minimize the loss.
5. Repeat until convergence or a stopping criterion is met.

## Advantages
- Simple to implement and interpret.
- Computationally efficient for small to medium-sized datasets.
- Works well when the relationship between variables is linear.

## Disadvantages
- Assumes a linear relationship, which may not hold in real-world scenarios.
- Sensitive to outliers, which can significantly affect the model.
- Prone to overfitting when the number of features is large compared to the number of observations.

## Use Cases
- Predicting house prices based on features like size, location, and number of rooms.
- Estimating sales revenue based on advertising spend.
- Modeling the relationship between temperature and energy consumption.
