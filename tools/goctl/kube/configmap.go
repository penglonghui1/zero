package kube

import (
	"context"
	"log"
	"os"

	"k8s.io/client-go/tools/clientcmd"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GenConfigMap(businessConfig map[string]string) {
	cm := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "flyele-dev-config-2",
			Namespace: "flyele",
		},
		Data: businessConfig,
	}
	//b, err := cm.Marshal()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/rudy/.kube/config")
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = clientSet.CoreV1().ConfigMaps("flyele").Create(context.Background(), &cm, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile("configmap.yaml", []byte(cm.String()), 777)

}
