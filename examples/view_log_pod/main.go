package main

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	logspb "github.com/evo-cloud/logs/go/gen/proto/logs"
	"github.com/golang/protobuf/proto"
	sdk "github.com/pigfall/k8ssdk"
	v1 "k8s.io/api/core/v1"
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
	rd, err := clientSet.CoreV1().Pods("tzz-dev").GetLogs("nginx", &v1.PodLogOptions{Container: "nginx", Follow: true}).Stream(context.Background())
	if err != nil {
		panic(err)
	}
	defer rd.Close()
	bufRd := bufio.NewReader(rd)
	for {
		var entry logspb.LogEntry
		line, err := bufRd.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		in := strings.TrimSpace(string(line))
		if err := proto.Unmarshal([]byte(in), &entry); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(entry)
		fmt.Println(string(line))
	}
}
