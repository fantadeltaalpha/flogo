package maptree

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	SourceObject string `md:"source"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	var err error
	r.SourceObject, err = coerce.ToString(values["source"])
	if err != nil {
		return err
	}
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"source": r.SourceObject,
	}
}

/*type Output struct {
	AnOutput string `md:"output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	obj, _ := coerce.ToString(values["output"])
	o.AnOutput = obj
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.AnOutput,
	}
}*/