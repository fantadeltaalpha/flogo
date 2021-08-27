package maptree

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)


func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestMap(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	iCtx := test.NewActivityInitContext(nil, nil)

	act, err := NewActivity(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("source", map[string]interface{}{"foo":"bar"})

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	fmt.Println(tc.GetOutput("output"))
}