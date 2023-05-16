package main

import (
	//"context"
	"fmt"
	"net/http"
	"strconv"

	sdk "github.com/pigfall/k8ssdk"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	//"github.com/telepresenceio/telepresence/v2/pkg/dnet"
)

func main() {
	config, err := sdk.LoadKubeConfig("/etc/rancher/k3s/k3s.yaml")
	if err != nil {
		panic(err)
	}
	tripper, upgrader, err := sdk.RouteTriper(config)
	if err != nil {
		panic(err)
	}

	clientSet, err := sdk.NewClientSet(config)
	if err != nil {
		panic(err)
	}

	reqURL := clientSet.CoreV1().RESTClient().
		Post().
		Resource("pods").
		Namespace("tzz-dev").
		Name("nginx").
		SubResource("portforward").
		URL()

	streamDialer := spdy.NewDialer(upgrader, &http.Client{Transport: tripper}, http.MethodPost, reqURL)
	conn, _, err := streamDialer.Dial(portforward.PortForwardProtocolV1Name)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	headers := http.Header{}
	headers.Set(core.PortHeader, strconv.FormatInt(int64(80), 10))
	headers.Set(core.PortForwardRequestIDHeader, strconv.FormatInt(1, 10))

	headers.Set(core.StreamType, core.StreamTypeError)
	errorStream, err := conn.CreateStream(headers)
	if err != nil {
		 panic(err)
	}
	_ = errorStream.Close()

	headers.Set(core.StreamType, core.StreamTypeData)
	stream, err := conn.CreateStream(headers)
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	//dialer,err := dnet.NewK8sPortForwardDialer(context.Background(),config.GetRestConfig(),clientSet)
	//if err != nil{
	//	panic(err)
	//}
	////conn2,err := dialer(context.Background(),"nginx-6595874d85-lzl2w.tzz:80")
	//conn2,err := dialer(context.Background(),"svc/nginx.tzz:80")
	//if err != nil{
	//	panic(err)
	//}
	//defer conn2.Close()
	_,err = stream.Write([]byte("GET\n\n"))
	if err != nil{
		panic(err)
	}
	buf := make([]byte,1024*8)
	fmt.Println("reading")
	n,err :=stream.Read(buf)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}
