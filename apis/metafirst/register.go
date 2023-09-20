package metafirst

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "metafirst.cnvrg.io"

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group:   GroupName,
		Version: runtime.APIVersionInternal,
	}

	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme

	types = []runtime.Object{
		&First{},
		&FirstList{},
	}
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, types...)
	return nil
}
