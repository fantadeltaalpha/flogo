package maptree

import (
	"github.com/fantadeltaalpha/flogo/extension/activity/maptree/mapper"
	"github.com/fantadeltaalpha/flogo/extension/activity/maptree/model"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})
var logger = activity.GetLogger("maptree")

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	/*s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)*/

	act := &Activity{ m : mapper.New()} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	m *mapper.Mapper
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
		return true, err
	}

	logger.Debugf("Input: %s", input.AnInput)
	logger.Infof("segments: %v", input.Segments)

	attr := model.Attribute{Value: "Test",Language: "ID"}
	attrObj,_ := coerce.ToObject(attr)
	output := &Output{AnOutput: attrObj}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}