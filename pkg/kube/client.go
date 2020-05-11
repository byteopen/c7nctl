package kube

import (
	"github.com/vinkdong/gox/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/helm/pkg/helm/portforwarder"
	"k8s.io/helm/pkg/kube"
	"os"
)

// TODO remove
func getConfig() (*rest.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	overrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmd.ClusterDefaults}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, overrides)

	config, err := clientConfig.ClientConfig()

	if err != nil {
		log.Error(err)
	}

	return config, err
}

func getClientset(c *rest.Config) (*kube.Tunnel, kubernetes.Interface, error) {
	client, err := kubernetes.NewForConfig(c)
	if err != nil {
		log.Error(err)
	}
	tunnel, err := portforwarder.New("kube-system", client, c)
	return tunnel, client, err
}

func GetTunnel() *kube.Tunnel {
	config, _ := getConfig()
	tunnel, _, err := getClientset(config)
	if err != nil {
		log.Error(err)
		os.Exit(129)
	}
	return tunnel
}

func GetConfig() (*rest.Config, error) {
	return getConfig()
}

func GetClient() kubernetes.Interface {
	config, _ := getConfig()
	_, client, _ := getClientset(config)
	return client
}
