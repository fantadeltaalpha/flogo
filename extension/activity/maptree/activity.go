package maptree

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)



func init() {
	_ = activity.Register(&Activity{},New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityLogger = log.ChildLogger(log.RootLogger(), "maptree")
var activityMd = activity.ToMetadata(&Input{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

func New(ctx activity.InitContext) (activity.Activity, error){
	return &Activity{},nil
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

	activityLogger.Debugf("Input: %v", input.SourceObject)

	/*output := &Output{AnOutput: input.SourceObject}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}*/

	return true, nil
}