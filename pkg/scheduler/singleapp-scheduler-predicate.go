package scheduler

import (
	"github.com/k8s-scheduler-extender/pkg/informers"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewSingleappPredicate() *Predicate {

	return &Predicate{
		Name: "singleappfilter",
		Func: func(pod v1.Pod, node v1.Node, nodeName string) (bool, error) {
			selector, _ := metav1.LabelSelectorAsSelector(metav1.SetAsLabelSelector(pod.Labels))
			pods, err := informers.SharedInformerFactory().Core().V1().Pods().Lister().Pods(pod.Namespace).List(selector)

			if err != nil {
				return false, err
			}

			for _, p := range pods {
				if p.Spec.NodeName == nodeName {
					return false, fmt.Errorf("pod already exists in node")
					break
				}
			}
			return true, nil
		},
	}
}
