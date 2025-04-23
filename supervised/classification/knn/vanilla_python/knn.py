import numpy as np
from collections import Counter

class KNNClassifier:
    def __init__(self, k=3, distance_metric='euclidean', weights='uniform'):
        self.k = k
        self.distance_metric = distance_metric
        self.weights = weights
    
    def fit(self, X_train, y_train):
        self.X_train = np.array(X_train, dtype=np.float32)
        self.y_train = np.array(y_train, dtype=np.float32)
    
    def _compute_distances(self, x):
        if self.distance_metric == 'euclidean':
            # Vectorized Euclidean distance
            # np.sum to vectorize the sum of squares
            # resulting in a 1D(scalar) array of distances
            return np.sqrt(np.sum((self.X_train - x) ** 2, axis=1))
        elif self.distance_metric == 'manhattan':
            # Vectorized Manhattan distance
            return np.sum(np.abs(self.X_train - x), axis=1)
        else:
            raise ValueError("Unsupported distance metric. Use 'euclidean' or 'manhattan'")
    
    def predict(self, X_test):
        X_test = np.array(X_test, dtype=np.float64)
        return np.array([self._predict_single(x) for x in X_test])
    
    def _predict_single(self, x):
        # Calculate distances
        # distances = [self._distance(x, x_train) for x_train in self.X_train]
        distances = self._compute_distances(x)
        
        # Get k nearest neighbors
        k_indices = np.argsort(distances)[:self.k] # argsort returns indices of the sorted array
        k_neighbor_labels = [self.y_train[i] for i in k_indices]
        k_neighbor_distances = [distances[i] for i in k_indices]

        # uniform: take majority vote(most frequent label)
        if self.weights == 'uniform':
            most_common = Counter(k_neighbor_labels).most_common(1)
            return most_common[0][0]
        
        # distance: take weighted vote
        # (closer neighbors have more influence)
        elif self.weights == 'distance':
            class_weights = {}
            for label, dist in zip(k_neighbor_labels, k_neighbor_distances):
                weight = 1 / (dist + 1e-5)  # add small epsilon to avoid division by zero
                class_weights[label] = class_weights.get(label, 0) + weight
            return max(class_weights.items(), key=lambda item: item[1])[0]
        
        else:
            raise ValueError("Unsupported weights. Use 'uniform' or 'distance'")
    
    def score(self, X_test, y_test):
        y_pred = self.predict(X_test)
        return np.mean(y_pred == y_test)
