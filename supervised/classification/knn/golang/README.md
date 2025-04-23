# Golang KNN Implementation

This is a pure Go implementation of the K-Nearest Neighbors (KNN) algorithm for classification tasks. It can be used for the same customer segmentation dataset as the JavaScript and Python implementations in the parent directories.

## Features

- Clean, idiomatic Go implementation
- Support for both Euclidean and Manhattan distance metrics
- Weighted or uniform voting methods
- Test accuracy similar to Python and JavaScript versions

## Package Structure

- `knn.go`: Main KNN classifier implementation
- `cmd/`: Contains the executable for training and evaluating on the customer segmentation dataset

## Usage

### Building and Running

Navigate to the `cmd` directory and run:

```bash
go run main.go
```

This will:
1. Load the customer segmentation dataset
2. Process the data with the same features as in the JS/Python versions
3. Train a KNN model with k=25 and evaluate it
4. Display the accuracy metrics and sample predictions


## API Reference

### Types

- `DistanceMetric`: String enum for distance metrics (`Euclidean`, `Manhattan`)
- `WeightMethod`: String enum for weighting methods (`Uniform`, `Distance`)
- `KNNClassifier`: The main classifier struct

### Constructor

```go
NewKNNClassifier(k int, distanceMetric DistanceMetric, weightMethod WeightMethod) *KNNClassifier
```

Creates a new KNN classifier with the specified parameters.

### Methods

```go
Fit(xTrain [][]float64, yTrain []float64) error
```

Trains the model on the provided data and labels.

```go
Predict(X [][]float64) ([]float64, error)
```

Makes predictions for multiple data points.

```go
PredictSingle(x []float64) (float64, error)
```

Makes a prediction for a single data point.

```go
Score(X [][]float64, yTrue []float64) (float64, error)
```

Calculates the accuracy of predictions compared to true labels. 