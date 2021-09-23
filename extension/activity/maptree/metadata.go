package maptree

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

type Input struct {
	TransactionID string `md:"businessTransactionID"`
	Segments map[string]interface{} `md:"segments"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	trxID, err := coerce.ToString(values["businessTransactionID"])
	if err != nil {
		return err
	}
	segments,err := coerce.ToObject(values["segments"])
	if err != nil {
		return err
	}
	r.TransactionID = trxID
	r.Segments = segments
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"businessTransactionID": r.TransactionID,
		"segments": r.Segments,
	}
}

type Output struct {
	Products map[string]interface{} `md:"products"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	products, err := coerce.ToObject(values["products"])
	if err != nil {
		return err
	}
	o.Products = products
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"products": o.Products,
	}
}