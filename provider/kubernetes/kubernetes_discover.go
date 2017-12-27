package kubernetes

import (
    "fmt"
    "io/ioutil"
    meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/labels"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "log"
)

type Provider struct{}

func (p *Provider) Help() string {
    return `Kubernetes:

    provider:  "kubernetes"
    namespace: The Kubernetes namespace to filter on
    service:   The Kubernetes service name to filter on
`
}

func (p *Provider) Addrs(args map[string]string, l *log.Logger) ([]string, error) {
    if args["provider"] != "kubernetes" {
        return nil, fmt.Errorf("discover-kubernetes: invalid provider " + args["provider"])
    }

    if l == nil {
        l = log.New(ioutil.Discard, "", 0)
    }

    namespace := args["namespace"]
    name := args["service"]

    config, err := rest.InClusterConfig()
    if err != nil {
        return nil, fmt.Errorf("discover-kubernetes: %s", err)
    }
    client, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, fmt.Errorf("discover-kubernetes: %s", err)
    }

    service, err := client.Core().Services(namespace).Get(name, meta_v1.GetOptions{})
    if err != nil {
        return nil, fmt.Errorf("discover-kubernetes: %s", err)
    }

    var podIpAddresses []string
    set := labels.Set(service.Spec.Selector)
    pods, err := client.Core().Pods(namespace).List(meta_v1.ListOptions{LabelSelector: set.AsSelector().String()})
    if err != nil {
        return nil, fmt.Errorf("discover-kubernetes: listing pods of service [%s]: %v", service.GetName(), err)
    }
    for _, v := range pods.Items {
        podIpAddresses = append(podIpAddresses, v.Status.PodIP)
    }

    return podIpAddresses, nil
}
