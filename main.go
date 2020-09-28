package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/rumpl/kwatch/informer"
	"github.com/rumpl/kwatch/store"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Not yet needed...
	// s := services.New()
	// server := NewServer()
	// v1 := server.e.Group("/api")
	// s.Register(v1)

	stopper := make(chan struct{})
	if err := informer.StartDeploymentsInformer(clientset, stopper); err != nil {
		return err
	}
	defer close(stopper)

	// Store test
	db, err := store.New()
	if err != nil {
		return err
	}

	dep := &store.Deployment{
		Name:  "test",
		Image: "image",
	}
	dep2 := &store.Deployment{
		Name:  "test2",
		Image: "image2",
	}
	db.Insert(dep)
	db.Insert(dep2)

	l, err := db.List()
	if err != nil {
		return err
	}

	for _, v := range l {
		fmt.Println(v.Name, v.Image)
	}

	return nil
}
