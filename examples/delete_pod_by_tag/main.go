package main

import (
	"context"

	sdk "github.com/pigfall/k8ssdk"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func main() {
	config, err := sdk.LoadKubeConfig("/etc/rancher/k3s/k3s.yaml")
	if err != nil {
		panic(err)
	}

	clientSet, err := sdk.NewClientSet(config)
	if err != nil {
		panic(err)
	}

	if err := clientSet.CoreV1().Pods("tzz-dev").DeleteCollection(
			context.Background(),
			metav1.DeleteOptions{},
			metav1.ListOptions{TypeMeta: metav1.TypeMeta{Kind: "Pod", APIVersion:"v1"},LabelSelector: "app=nginx"},
	);err != nil{
		panic(err)
	}

}
