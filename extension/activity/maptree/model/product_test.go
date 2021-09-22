package model_test

import (
	"testing"

	"github.com/fantadeltaalpha/flogo/extension/activity/maptree/model"
	jsoniter "github.com/json-iterator/go"
)

func TestJsonToMap(t *testing.T){
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	prod := model.Product{
		ID: "Test",
		Name: "Halo",
		Type: "PO",	
		//PCO: []string{"MR_1234"},
		ProductComprisedOf: []model.Product{
			{ID: "MR_1234",Name: "MR Test",Type: "MR"},
		},	
	}

	data,err := json.Marshal(prod)
	if err != nil {
		panic(err)
	}

	resp := make(map[string]interface{})

	err  = json.Unmarshal(data,&resp)
	if err != nil {
		panic(err)
	}

	t.Log(resp)
}