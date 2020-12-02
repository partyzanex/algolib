package knn

import (
	"bytes"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestKNN(t *testing.T) {
	client := http.DefaultClient
	resp, err := client.Get("https://raw.githubusercontent.com/sjwhitworth/golearn/master/examples/datasets/iris.csv")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	_ = resp.Body.Close()

	rawData, err := base.ParseCSVToInstancesFromReader(bytes.NewReader(b), false)
	if err != nil {
		t.Fatal(err)
	}

	cls := knn.NewKnnClassifier("euclidean", "kdtree", 2)

	train, test := base.InstancesTrainTestSplit(rawData, 0.11)

	err = cls.Fit(train)
	if err != nil {
		t.Fatal(err)
	}

	// _ = cls.Save("./model.ml")

	p, err := cls.Predict(test)
	if err != nil {
		t.Fatal(err)
	}

	confusionMat, err := evaluation.GetConfusionMatrix(test, p)
	if err != nil {
		t.Fatal(err)
	}

	// prec := evaluation.GetPrecision("Iris-setosa", confusionMat)
	// t.Log(prec)

	t.Log(evaluation.GetSummary(confusionMat))
}
