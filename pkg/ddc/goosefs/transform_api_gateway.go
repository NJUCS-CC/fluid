package goosefs

import (
	"fmt"

	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
)

// transformAPIGateway decide whether to enable APIGateway in value according to GooseFSRuntime
func (e *GooseFSEngine) transformAPIGateway(runtime *datav1alpha1.GooseFSRuntime, value *GooseFS) (err error) {
	if runtime == nil || value == nil {
		err = fmt.Errorf("cannot transform because runtime or value will lead to nil pointer")
		return
	}
	value.APIGateway.Enabled = runtime.Spec.APIGateway.Enabled
	return
}
