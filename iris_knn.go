package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	rawData, err := base.ParseCSVToInstances("iris.csv", false)
	if err != nil {
		panic(err)
	}

	// print summary
	fmt.Println(rawData)

	// initial new KNN classifier
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	// do train, test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	// calc euclidean distances and return most popular
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	// print precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}