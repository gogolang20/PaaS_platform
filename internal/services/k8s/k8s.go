package k8s

import (
	"flag"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

const (
	defaultResync = 0
)

var (
	kubeconfig *string
	clientset  *kubernetes.Clientset
)

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "../testdata/gaia/testkubeconfig.yml", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logrus.Error("[k8s][init] get config error: ", err)
		return
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Error("[k8s][init] create clientset error: ", err)
		return
	}
}

func CreateDeployment() {
	factory := informers.NewSharedInformerFactory(clientset, defaultResync)

	services := factory.Core().V1().Services()
	services.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {

		},
		UpdateFunc: func(oldObj, newObj interface{}) {

		},
		DeleteFunc: func(obj interface{}) {

		},
	})
}
