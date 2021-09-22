package maptree

import (
	"github.com/project-flogo/core/data/coerce"
)

/*type Settings struct {
	ASetting string `md:"aSetting,required"`
}*/

type Input struct {
	AnInput string `md:"anInput,required"`
	Segments []interface{} `md:"segments"`
}

type Segment struct {
	Type string `md:"type"`
	Name string `md:"name"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anInput"])
	segments,_ := coerce.ToArray(values["segments"])
	r.AnInput = strVal
	r.Segments = segments
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anInput": r.AnInput,
		"segments": r.Segments,
	}
}

type Output struct {
	AnOutput interface{} `md:"anOutput"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anOutput"])
	o.AnOutput = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	//return structs.Map(o)
	return map[string]interface{}{
		"anOutput": o.AnOutput,
	}
}