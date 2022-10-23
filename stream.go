package k8ssdk

import(
	"k8s.io/client-go/transport/spdy"
	"net/http"
)

func RouteTriper(kubeconfig *KubeConfig)(http.RoundTripper,spdy.Upgrader,error){
	return spdy.RoundTripperFor(kubeconfig.config)
}
