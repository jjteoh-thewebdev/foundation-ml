const fs = require('fs'); // for reading csv

class KNNClassifier {
    constructor(k=3, distanceMetric='euclidean', weights='uniform') {
        this.k = k; // Number of nearest neighbors
        this.distanceMetric = distanceMetric; // Distance metric
        this.weights = weights; // Weighting method
        this.X_train = [];
        this.y_train = [];
    }

    // Add training data
    fit(X_train, y_train) {
        if (X_train.length !== y_train.length) {
            throw new Error("Data and labels must have the same length.");
        }
        this.X_train = X_train;
        this.y_train = y_train;
    }

    // Distance calculation
    _distance(a, b) {
        if (this.distanceMetric === 'euclidean') {
            return Math.sqrt(a.reduce((sum, val, i) => sum + (val - b[i]) ** 2, 0));
        } else if (this.distanceMetric === 'manhattan') {
            return a.reduce((sum, val, i) => sum + Math.abs(val - b[i]), 0);
        } else {
            throw new Error("Unsupported distance metric: use 'euclidean' or 'manhattan'");
        }
    }


    // Predict the label for a single data point
    predictSingle(x) {
        const distances = this.X_train.map((x_train, i) => ({
            label: this.y_train[i],
            distance: this._distance(x, x_train),
          }));
      
          // Sort by distance
          distances.sort((a, b) => a.distance - b.distance);
          const neighbors = distances.slice(0, this.k);
      
          if (this.weights === 'uniform') {
            // Count labels
            const counts = {};
            for (const neighbor of neighbors) {
              counts[neighbor.label] = (counts[neighbor.label] || 0) + 1;
            }
            // Return the most frequent
            return Object.entries(counts).reduce((a, b) => (a[1] > b[1] ? a : b))[0];
          }
      
          else if (this.weights === 'distance') {
            const weights = {};
            for (const neighbor of neighbors) {
              const weight = 1 / (neighbor.distance + 1e-5); // avoid div by 0
              weights[neighbor.label] = (weights[neighbor.label] || 0) + weight;
            }
            return Object.entries(weights).reduce((a, b) => (a[1] > b[1] ? a : b))[0];
          }
      
          else {
            throw new Error("Unsupported weight method: use 'uniform' or 'distance'");
          }
    }

    // Predict labels for multiple data points
    predict(X) {
        if (!this.X_train.length) {
            throw new Error("Model has not been trained yet.");
        }

        if (!Array.isArray(X)) {
            throw new Error("Input data must be an array.");
        }

        if (X.length === 0) {
            return [];
        }

        return X.map((point) => this.predictSingle(point));
    }

    score(X, y_true) {
        const y_pred = this.predict(X);
        let correct = 0;
        for (let i = 0; i < y_true.length; i++) {
          if (parseFloat(y_pred[i]) === parseFloat(y_true[i])) correct++;
        }
        return correct / y_true.length;
    }
}


function mulberry32(seed) {
    return function () {
        seed |= 0;
        seed = seed + 0x6D2B79F5 | 0;
        let t = Math.imul(seed ^ seed >>> 15, 1 | seed);
        t = t + Math.imul(t ^ t >>> 7, 61 | t) ^ t;
        return ((t ^ t >>> 14) >>> 0) / 4294967296;
    }
}

function shuffle(array, randomFn) {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(randomFn() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
}


// Load data from CSV
// make sure to run this from the js/ folder
const filePath = `../../../../dataset/customer_segmentation/clean/data.csv`;
const data = fs.readFileSync(filePath, 'utf8');
let rows = data.trim().split('\n').map(row => row.split(','));

console.log(`Total rows: ${rows.length}`);
console.log(`Total columns: ${rows[0].length}`);
console.log(`Columns`, rows[0])

// Remove the header row
rows = rows.slice(1);

// Feature selection
// Drop 'Work_Experience', 'Gender_Male', 'Gender_Female', 'Family_Size'
rows = rows.map(row => {
    return [
        parseFloat(row[0]), // Age
        row[5] === `True`? 1 : 0, // Ever_Married_No
        row[6] === `True`? 1 : 0, // Ever_Married_Yes
        row[7] === `True`? 1 : 0, // Graduated_No
        row[8] === `True`? 1 : 0, // Graduated_Yes
        parseFloat(row[9]), // Profession_Freq
        parseFloat(row[10]), // Spending_Score_Freq,
        parseFloat(row[11]), // Category_Freq
        parseFloat(row[12]), // Segmentation_Freq
    ]
})

// Shuffle data
// Set seed for reproducibility
const seed = 1234;
const rand = mulberry32(seed);
shuffle(rows, rand);

// Split data into training and test sets
const splitIndex = Math.floor(rows.length * 0.8);
const trainingData = rows.slice(0, splitIndex);
const testData = rows.slice(splitIndex);

// Remove labels from training and test data
// last column is label
const X_train = trainingData.map(row => row.slice(0, -1));
const X_test = testData.map(row => row.slice(0, -1));
const y_train = trainingData.map(row => row[row.length - 1]);  
const y_test = testData.map(row => row[row.length - 1]);

const knn = new KNNClassifier(25, 'euclidean', 'uniform');
knn.fit(X_train, y_train);

console.time("Prediction Time(Training)");
console.log("Training accuracy: ", knn.score(X_train, y_train));
console.timeEnd("Prediction Time(Training)");

console.time("Prediction Time(Testing)");
console.log("Test accuracy: ", knn.score(X_test, y_test));
console.timeEnd("Prediction Time(Testing)");

// y_pred = knn.predict(X_test);
// console.log("Predictions: ", X_test[0], y_pred[0], y_test[0]);
// console.log("Predictions: ", X_test[1], y_pred[1], y_test[1]);
// console.log("Predictions: ", X_test[2], y_pred[2], y_test[2]); 
// console.log("Predictions: ", X_test[3], y_pred[3], y_test[3]);
// console.log("Predictions: ", X_test[4], y_pred[4], y_test[4]);
// console.log("Predictions: ", X_test[5], y_pred[5], y_test[5]);
// console.log("Predictions: ", X_test[6], y_pred[6], y_test[6]);
// console.log("Predictions: ", X_test[7], y_pred[7], y_test[7]);
// console.log("Predictions: ", X_test[8], y_pred[8], y_test[8]);
// console.log("Predictions: ", X_test[9], y_pred[9], y_test[9]);
