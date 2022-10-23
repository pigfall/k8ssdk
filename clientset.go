package k8ssdk

import(
	k8s "k8s.io/client-go/kubernetes"
)


func NewClientSet(kubeconfig *KubeConfig)(k8s.Interface,error){
	return k8s.NewForConfig(kubeconfig.config)
}
