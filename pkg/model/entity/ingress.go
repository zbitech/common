package entity

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"time"
)

type Ingress struct {
	Name      string
	Namespace string
	Object    *unstructured.Unstructured
	Created   time.Time
	Updated   time.Time
}
