package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	knn "github.com/jj/foundation-ml/supervised/classification/knn/golang"
)

// readCSV reads a CSV file and returns its contents as a slice of string slices
func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// prepareData processes the CSV data and returns features and labels
func prepareData(data [][]string) ([][]float64, []float64, error) {
	// Skip header row
	rows := data[1:]

	X := make([][]float64, len(rows))
	y := make([]float64, len(rows))

	for i, row := range rows {
		// Feature selection
		// [Age, Ever_Married_No, Ever_Married_Yes, Graduated_No, Graduated_Yes, 
		// Profession_Freq, Spending_Score_Freq, Category_Freq, Segmentation_Freq]
		
		features := make([]float64, 8) // Excluding the label column

		// Age
		age, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing Age: %w", err)
		}
		features[0] = age

		// Ever_Married_No
		if row[5] == "True" {
			features[1] = 1.0
		} else {
			features[1] = 0.0
		}

		// Ever_Married_Yes
		if row[6] == "True" {
			features[2] = 1.0
		} else {
			features[2] = 0.0
		}

		// Graduated_No
		if row[7] == "True" {
			features[3] = 1.0
		} else {
			features[3] = 0.0
		}

		// Graduated_Yes
		if row[8] == "True" {
			features[4] = 1.0
		} else {
			features[4] = 0.0
		}

		// Profession_Freq
		profFreq, err := strconv.ParseFloat(row[9], 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing Profession_Freq: %w", err)
		}
		features[5] = profFreq

		// Spending_Score_Freq
		spendFreq, err := strconv.ParseFloat(row[10], 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing Spending_Score_Freq: %w", err)
		}
		features[6] = spendFreq

		// Category_Freq
		catFreq, err := strconv.ParseFloat(row[11], 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing Category_Freq: %w", err)
		}
		features[7] = catFreq

		// Segmentation_Freq (label)
		segFreq, err := strconv.ParseFloat(row[12], 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing Segmentation_Freq: %w", err)
		}
		y[i] = segFreq

		X[i] = features
	}

	return X, y, nil
}

// Function to shuffle data with the same seed for reproducibility
func shuffle(X [][]float64, y []float64, seed int64) {
	rand.Seed(seed)
	for i := len(X) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		X[i], X[j] = X[j], X[i]
		y[i], y[j] = y[j], y[i]
	}
}

// splitData splits data into training and test sets
func splitData(X [][]float64, y []float64, testRatio float64) ([][]float64, []float64, [][]float64, []float64) {
	splitIndex := int(float64(len(X)) * (1 - testRatio))
	return X[:splitIndex], y[:splitIndex], X[splitIndex:], y[splitIndex:]
}

func main() {
	// Set paths
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}


	var csvData [][]string
	var dataPath string

	// make sure to run this from the cmd/ folder
	path := "../../../../../dataset/customer_segmentation/clean/data.csv"
	absPath := filepath.Join(wd, path)
	if _, err := os.Stat(absPath); err == nil {
		csvData, err = readCSV(absPath)
		if err == nil {
			dataPath = absPath
		}
	} else if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err == nil {
			csvData, err = readCSV(path)
			if err == nil {
				dataPath = path
			}
		}
	}
	

	if csvData == nil {
		log.Fatalf("Could not find dataset file in any of the specified paths")
	}

	fmt.Printf("Found dataset at: %s\n", dataPath)
	fmt.Printf("Total rows: %d\n", len(csvData)-1) // Subtract header row
	fmt.Printf("Total columns: %d\n", len(csvData[0]))

	// Prepare data
	X, y, err := prepareData(csvData)
	if err != nil {
		log.Fatalf("Error preparing data: %v", err)
	}

	// Shuffle data with the same seed for reproducibility
	shuffle(X, y, 1234)

	// Split into training and test sets
	XTrain, yTrain, XTest, yTest := splitData(X, y, 0.2)

	fmt.Printf("Training set size: %d\n", len(XTrain))
	fmt.Printf("Test set size: %d\n", len(XTest))

	// Create and train the model
	fmt.Println("Training model...")
	model := knn.NewKNNClassifier(25, knn.Euclidean, knn.Uniform)
	err = model.Fit(XTrain, yTrain)
	if err != nil {
		log.Fatalf("Error fitting model: %v", err)
	}

	// Evaluate on training set
	start := time.Now()
	trainAcc, err := model.Score(XTrain, yTrain)
	if err != nil {
		log.Fatalf("Error scoring on training set: %v", err)
	}
	trainTime := time.Since(start)
	fmt.Printf("Training accuracy: %.4f\n", trainAcc)
	fmt.Printf("Prediction time for training dataset: %.4f seconds\n", trainTime.Seconds())

	// Evaluate on test set
	start = time.Now()
	testAcc, err := model.Score(XTest, yTest)
	if err != nil {
		log.Fatalf("Error scoring on test set: %v", err)
	}
	testTime := time.Since(start)
	fmt.Printf("Test accuracy: %.4f\n", testAcc)
	fmt.Printf("Prediction time for test dataset: %.4f seconds\n", testTime.Seconds())

	// Test predictions for a few samples
	fmt.Println("\nSample predictions:")
	predictions, err := model.Predict(XTest[:10])
	if err != nil {
		log.Fatalf("Error predicting: %v", err)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Sample %d:\n", i+1)
		fmt.Printf("  Features: %v\n", XTest[i])
		fmt.Printf("  Predicted: %.0f, Actual: %.0f\n", predictions[i], yTest[i])
	}
} 