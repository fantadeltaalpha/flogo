package mapper_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/fantadeltaalpha/flogo/extension/activity/maptree/mapper"
)

func BenchmarkParseJSONIterJob(b *testing.B){
	file,err := os.Open("/Users/fanggara/Downloads/JSON/Response/response_OfferInternet.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data,err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//dataStr := string(data)
	mapper := mapper.New()
	for n := 0; n < b.N; n++ {
		_,err = mapper.ToProduct(data)
		if err != nil {
			panic(err)
		}
	}
}