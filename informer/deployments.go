package informer

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// StartDeploymentsInformer creates a new informer that looks for deployments
func StartDeploymentsInformer(clientset *kubernetes.Clientset, stopper chan struct{}) error {
	// panic
	defer runtime.HandleCrash()

	factory := informers.NewSharedInformerFactory(clientset, 0)
	informer := factory.Apps().V1().Deployments().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		DeleteFunc: onDelete,
	})

	go informer.Run(stopper)

	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		return fmt.Errorf("Timed out")
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
