package k8ssdk

import(
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
)

type KubeConfig struct{
	config *rest.Config
}

func LoadKubeConfig(configFilepath string)(*KubeConfig,error){
	config,err := clientcmd.BuildConfigFromFlags("", configFilepath)
	if err != nil{
		return nil,err
	}
	return &KubeConfig{config:config},nil
}

func (c *KubeConfig) GetRestConfig()*rest.Config{
	return c.config
}
