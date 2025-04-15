import pandas as pd
import time
from knn import KNNClassifier
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score

from matplotlib import pyplot as plt
from sklearn.metrics import ConfusionMatrixDisplay, confusion_matrix

# Load the dataset
data_src = '../../../../dataset/customer_segmentation'

data = pd.read_csv(f'{data_src}/clean/data.csv')

print(f"Total rows: {data.shape[0]}")
print(f"Total columns: {data.shape[1]}")
data.head()

# Feature selection
data = data.drop(columns=['Gender_Male', 'Gender_Female', 'Family_Size', 'Work_Experience'])

# Split the dataset into features and target variable
X = data.drop('Segmentation_Freq',axis=1)
y = data['Segmentation_Freq']

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)
print(X_train.shape,y_train.shape,X_test.shape,y_test.shape)


# Train model
print("Training model...")
model = KNNClassifier(k=25, distance_metric='euclidean')
model.fit(X_train, y_train)

start_time = time.time()
y_train_pred = model.predict(X_train)
end_time = time.time()
print(f"Prediction time for training dataset: {end_time - start_time:.4f} seconds")

start_time = time.time()
y_test_pred = model.predict(X_test)
end_time = time.time()
print(f"Prediction time for test dataset: {end_time - start_time:.4f} seconds")


print("train_accuracy",accuracy_score(y_train,y_train_pred))
print("test_accuracy",accuracy_score(y_test,y_test_pred))


cm=confusion_matrix(y_test,y_test_pred)

print("Confusion Matrix:")
print(cm)

# This may not work in devcontainer, anyway we will export to png for viewing
ConfusionMatrixDisplay(cm).plot()
plt.show()
plt.savefig("out/plot.png")  # export as png
plt.close()