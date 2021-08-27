package maptree

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	SourceObject map[string]interface{} `md:"source,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	obj, _ := coerce.ToObject(values["source"])
	r.SourceObject = obj
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"source": r.SourceObject,
	}
}

type Output struct {
	AnOutput map[string]interface{} `md:"output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	obj, _ := coerce.ToObject(values["output"])
	o.AnOutput = obj
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.AnOutput,
	}
}

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity,error) {
	return &Activity{},nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
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

	ctx.Logger().Debugf("Input: %v", input.SourceObject)

	output := &Output{AnOutput: input.SourceObject}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}