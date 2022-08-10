package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"
  // "os"

	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
  // "k8s.io/api"
)

func errCheck(err error) {
  if err != nil {
    panic(err.Error())
  }
}

func k8sInit() (*kubernetes.Clientset, string) {
  config, err := rest.InClusterConfig()
  errCheck(err)

  clientset, err := kubernetes.NewForConfig(config)
  errCheck(err)

  currentNamespaceBytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
  errCheck(err)

  currentNamespace := string(currentNamespaceBytes)

  return clientset, currentNamespace
}

func main() {

  labelMap := make(map[string]string)

  labelMap["app.kubernetes.io/instance"] = "rapidfort"

  podLabelSelector := metav1.LabelSelector{
    MatchLabels: labelMap,
  }

  podListOptions := metav1.ListOptions{
    LabelSelector: labels.Set(podLabelSelector.MatchLabels).String(),
  }
  println("Started")

  println(podListOptions.LabelSelector)

  clientset, currentNamespace := k8sInit()

  print("namespace: ")
  println(currentNamespace)

  ready := false

  for !ready {
    pods, err := clientset.CoreV1().Pods(currentNamespace).List(context.Background(), podListOptions)
    errCheck(err)

    fmt.Printf("%d pods were found\n", len(pods.Items))
    var readyPods int

    for _, pod := range pods.Items {
      totalContainers := len(pod.Spec.Containers)

      var readyContainers int

      for _, containerStatus := range pod.Status.ContainerStatuses {
        if containerStatus.Ready {
          readyContainers++
        }
      }

      if readyContainers == totalContainers {
        readyPods++
      }

      fmt.Printf("Pod %s container status: %d/%d ready\n", pod.Name, readyContainers, totalContainers)
    }

    if readyPods == len(pods.Items) {
      ready = true;
      continue
    }

    time.Sleep(5 * time.Second)
  }
}