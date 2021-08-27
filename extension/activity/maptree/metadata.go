package maptree

import "github.com/project-flogo/core/data/coerce"

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