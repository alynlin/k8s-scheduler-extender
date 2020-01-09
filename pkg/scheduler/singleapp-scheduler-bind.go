package scheduler

import (
	"github.com/k8s-scheduler-extender/pkg/informers"
	k8s "github.com/k8s-scheduler-extender/pkg/simple/client"
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func NewNoSingleapBind() *Bind {
	return &Bind{
		Func: func(podName string, podNamespace string, podUID types.UID, node string) error {
			return fmt.Errorf("This extender doesn't support Bind.  Please make 'BindVerb' be empty in your ExtenderConfig.")
		},
	}
}

func NewSingleapBind() *Bind {
	return &Bind{
		Func: func(podName string, podNamespace string, podUID types.UID, node string) error {

			/*pod, err := getPod(podName, podNamespace, podUID)
			if err != nil {
				log.Printf("warn: Failed to handle pod %s in ns %s due to error %v", podName, podNamespace, err)
				return err
			}*/

			//todo something

			return nil

		},
	}
}

func getPod(name string, namespace string, podUID types.UID) (pod *v1.Pod, err error) {
	pod, err = informers.SharedInformerFactory().Core().V1().Pods().Lister().Pods(namespace).Get(name);
	if errors.IsNotFound(err) {
		pod, err = k8s.Client().CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	if pod.UID != podUID {
		pod, err = k8s.Client().CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		if pod.UID != podUID {
			return nil, fmt.Errorf("The pod %s in ns %s's uid is %v, and it's not equal with expected %v",
				name,
				namespace,
				pod.UID,
				podUID)
		}
	}

	return pod, nil
}
