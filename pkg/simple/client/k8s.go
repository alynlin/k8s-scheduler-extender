package client

import (
	"flag"
	"log"
	"os"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfigFile string
	k8sClient      *kubernetes.Clientset
	k8sClientOnce  sync.Once
	KubeConfig     *rest.Config
	MasterURL      string
)

func init() {
	flag.StringVar(&kubeConfigFile, "kubeconfig", "", "path to kubeconfig file")
	flag.StringVar(&MasterURL, "master-url", "", "kube-apiserver url, only needed when out of cluster")
}

func Client() *kubernetes.Clientset {

	k8sClientOnce.Do(func() {

		config, err := Config()

		if err != nil {
			log.Fatalln(err)
		}

		k8sClient = kubernetes.NewForConfigOrDie(config)

		KubeConfig = config
	})

	return k8sClient
}

func Config() (kubeConfig *rest.Config, err error) {

	if _, err = os.Stat(kubeConfigFile); err == nil {
		kubeConfig, err = clientcmd.BuildConfigFromFlags(MasterURL, kubeConfigFile)
	} else {
		kubeConfig, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	kubeConfig.QPS = 1e6
	kubeConfig.Burst = 1e6

	return kubeConfig, nil
}
