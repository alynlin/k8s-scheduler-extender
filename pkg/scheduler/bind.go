package scheduler

import (
	"k8s.io/apimachinery/pkg/types"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

type Bind struct {
	Func func(podName string, podNamespace string, podUID types.UID, node string) error
}

func (b Bind) Handler(args schedulerapi.ExtenderBindingArgs) *schedulerapi.ExtenderBindingResult {
	err := b.Func(args.PodName, args.PodNamespace, args.PodUID, args.Node)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return &schedulerapi.ExtenderBindingResult{
		Error: errMsg,
	}
}
