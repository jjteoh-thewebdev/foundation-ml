package knn

import (
	"fmt"
	"math"
	"sort"
)

// DistanceMetric enum type
type DistanceMetric string

const (
	Euclidean DistanceMetric = "euclidean"
	Manhattan DistanceMetric = "manhattan"
)

// WeightMethod enum type
type WeightMethod string

const (
	Uniform  WeightMethod = "uniform"
	Distance WeightMethod = "distance"
)

// KNNClassifier represents a K-Nearest Neighbors classifier
type KNNClassifier struct {
	k             int
	distanceMetric DistanceMetric
	weightMethod  WeightMethod
	xTrain        [][]float64
	yTrain        []float64
}

// NewKNNClassifier creates a new KNN classifier with the specified parameters
func NewKNNClassifier(k int, distanceMetric DistanceMetric, weightMethod WeightMethod) *KNNClassifier {
	return &KNNClassifier{
		k:             k,
		distanceMetric: distanceMetric,
		weightMethod:  weightMethod,
		xTrain:        [][]float64{},
		yTrain:        []float64{},
	}
}

// Fit trains the KNN model with the provided training data
func (knn *KNNClassifier) Fit(xTrain [][]float64, yTrain []float64) error {
	if len(xTrain) != len(yTrain) {
		return fmt.Errorf("data and labels must have the same length")
	}
	
	// Make deep copies of the data to avoid reference issues
	knn.xTrain = make([][]float64, len(xTrain))
	for i, x := range xTrain {
		knn.xTrain[i] = make([]float64, len(x))
		copy(knn.xTrain[i], x)
	}
	
	knn.yTrain = make([]float64, len(yTrain))
	copy(knn.yTrain, yTrain)
	
	return nil
}

// Calculate distance between two vectors
func (knn *KNNClassifier) distance(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vectors must have the same length")
	}
	
	switch knn.distanceMetric {
	case Euclidean:
		sum := 0.0
		for i := 0; i < len(a); i++ {
			sum += math.Pow(a[i]-b[i], 2)
		}
		return math.Sqrt(sum), nil
		
	case Manhattan:
		sum := 0.0
		for i := 0; i < len(a); i++ {
			sum += math.Abs(a[i] - b[i])
		}
		return sum, nil
		
	default:
		return 0, fmt.Errorf("unsupported distance metric: use 'euclidean' or 'manhattan'")
	}
}

// Neighbor represents a data point and its distance from a query point
type Neighbor struct {
	Label    float64
	Distance float64
}

// PredictSingle predicts the label for a single data point
func (knn *KNNClassifier) PredictSingle(x []float64) (float64, error) {
	if len(knn.xTrain) == 0 {
		return 0, fmt.Errorf("model has not been trained yet")
	}

	// Calculate distances to all training points
	neighbors := make([]Neighbor, len(knn.xTrain))
	for i, xTrain := range knn.xTrain {
		dist, err := knn.distance(x, xTrain)
		if err != nil {
			return 0, err
		}
		
		neighbors[i] = Neighbor{
			Label:    knn.yTrain[i],
			Distance: dist,
		}
	}
	
	// Sort by distance
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance < neighbors[j].Distance
	})
	
	// Get k-nearest neighbors
	k := knn.k
	if k > len(neighbors) {
		k = len(neighbors)
	}
	
	kNearest := neighbors[:k]
	
	// Determine the predicted label based on weight method
	switch knn.weightMethod {
	case Uniform:
		// Count labels
		counts := make(map[float64]int)
		for _, neighbor := range kNearest {
			counts[neighbor.Label]++
		}
		
		// Find the most frequent label
		var maxCount int
		var predictedLabel float64
		for label, count := range counts {
			if count > maxCount {
				maxCount = count
				predictedLabel = label
			}
		}
		
		return predictedLabel, nil
		
	case Distance:
		// Weighted voting based on distance
		weights := make(map[float64]float64)
		for _, neighbor := range kNearest {
			// Add small epsilon to avoid division by zero
			weight := 1.0 / (neighbor.Distance + 1e-5)
			weights[neighbor.Label] += weight
		}
		
		// Find the label with the highest weighted vote
		var maxWeight float64
		var predictedLabel float64
		for label, weight := range weights {
			if weight > maxWeight {
				maxWeight = weight
				predictedLabel = label
			}
		}
		
		return predictedLabel, nil
		
	default:
		return 0, fmt.Errorf("unsupported weight method: use 'uniform' or 'distance'")
	}
}

// Predict predicts labels for multiple data points
func (knn *KNNClassifier) Predict(X [][]float64) ([]float64, error) {
	if len(knn.xTrain) == 0 {
		return nil, fmt.Errorf("model has not been trained yet")
	}
	
	predictions := make([]float64, len(X))
	
	for i, x := range X {
		prediction, err := knn.PredictSingle(x)
		if err != nil {
			return nil, err
		}
		predictions[i] = prediction
	}
	
	return predictions, nil
}

// Score calculates the accuracy of predictions compared to true labels
func (knn *KNNClassifier) Score(X [][]float64, yTrue []float64) (float64, error) {
	if len(X) != len(yTrue) {
		return 0, fmt.Errorf("data and labels must have the same length")
	}
	
	yPred, err := knn.Predict(X)
	if err != nil {
		return 0, err
	}
	
	correct := 0
	for i := 0; i < len(yTrue); i++ {
		if yPred[i] == yTrue[i] {
			correct++
		}
	}
	
	return float64(correct) / float64(len(yTrue)), nil
} 