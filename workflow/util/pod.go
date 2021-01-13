package util

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func PodToUnstructured(pod *corev1.Pod) (*unstructured.Unstructured, error) {
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(pod)
	un := &unstructured.Unstructured{Object: obj}
	un.SetKind("Pod")
	un.SetAPIVersion("v1")
	return un, err
}

func PodFromUnstructured(un *unstructured.Unstructured) (*corev1.Pod, error) {
	pod := &corev1.Pod{}
	return pod, runtime.DefaultUnstructuredConverter.FromUnstructured(un.Object, pod)
}
