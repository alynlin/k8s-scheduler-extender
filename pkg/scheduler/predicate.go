package scheduler

import (
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

type Predicate struct {
	Name string
	Func func(pod v1.Pod, node v1.Node, nodeName string) (bool, error)
}

func (p Predicate) Handler(args schedulerapi.ExtenderArgs) *schedulerapi.ExtenderFilterResult {
	pod := *args.Pod
	nodeNames := *args.NodeNames
	canSchedule := make([]string, 0, len(nodeNames))
	canNotSchedule := make(map[string]string)

	for _, nodeName := range nodeNames {
		node := v1.Node{}
		result, err := p.Func(pod, node, nodeName)
		if err != nil {
			canNotSchedule[nodeName] = err.Error()
		} else {
			if result {
				canSchedule = append(canSchedule, nodeName)
			}
		}
	}

	result := schedulerapi.ExtenderFilterResult{
		NodeNames:   &canSchedule,
		FailedNodes: canNotSchedule,
		Error:       "",
	}

	return &result
}
