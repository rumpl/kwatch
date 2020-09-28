package main

import (
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

	db, err := store.New()
	if err != nil {
		return err
	}

	// Not yet needed...
	// s := services.New()
	// server := NewServer()
	// v1 := server.e.Group("/api")
	// s.Register(v1)

	inf := informer.NewDeploymentsInformer(db, clientset)
	stopper := make(chan struct{})
	defer close(stopper)

	if err := inf.Listen(stopper); err != nil {
		return err
	}

	return nil
}
