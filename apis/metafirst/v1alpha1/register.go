package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "metafirst.cnvrg.io"

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group:   GroupName,
		Version: "v1alpha1",
	}

	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme

	types = []runtime.Object{
		&First{},
		&FirstList{},
	}
)

func init() {
	localSchemeBuilder.Register(addKnownTypes)
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, types...)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
