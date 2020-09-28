package informer

import (
	"fmt"

	"github.com/rumpl/kwatch/store"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// DeploymentsInformer is the informer for deployments
type DeploymentsInformer struct {
	store     store.Store
	clientset *kubernetes.Clientset
}

// NewDeploymentsInformer you know what it does
func NewDeploymentsInformer(store store.Store, clientset *kubernetes.Clientset) *DeploymentsInformer {
	return &DeploymentsInformer{
		store:     store,
		clientset: clientset,
	}
}

// Listen creates a new informer that looks for deployments
func (di *DeploymentsInformer) Listen(stopper chan struct{}) error {
	// panic
	defer runtime.HandleCrash()

	factory := informers.NewSharedInformerFactory(di.clientset, 0)
	informer := factory.Apps().V1().Deployments().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		DeleteFunc: onDelete,
	})

	go informer.Run(stopper)

	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		return fmt.Errorf("sync timeout")
	}

	return nil
}

func onAdd(obj interface{}) {
	node := obj.(*appsv1.Deployment)
	fmt.Println("add", node.Name)
	for _, c := range node.Spec.Template.Spec.Containers {
		fmt.Println(c.Image)
	}
}

func onDelete(obj interface{}) {
	node := obj.(*appsv1.Deployment)
	fmt.Println("delete", node.Name)
	for _, c := range node.Spec.Template.Spec.Containers {
		fmt.Println(c.Image)
	}
}
