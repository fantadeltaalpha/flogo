package maptree

import (
	"encoding/json"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{},&Input{}, &Output{})

var logger = log.ChildLogger(log.RootLogger(), "maptree")

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	/*s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)*/

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	//m *mapper.Mapper
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	logger.Infof("Input.TrxID: %s", input.TransactionID)
	logger.Infof("Input.Segments: %v", input.Segments)

	product := Product{ID: "002992",Type: "PO"} 
	data,err := json.Marshal(product)
	if err != nil {
		return false, err
	}
	products:= make(map[string]interface{})
	err = json.Unmarshal(data, &products)
	
	output := &Output{Products: products}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

type Product struct{
	ID string `json:"id"`
	Type string `json:"type"`
}