package first

import (
	"github.com/Dimss/first/apis/metafirst"
	"github.com/Dimss/first/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := newStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:     func() runtime.Object { return &metafirst.First{} },
		NewListFunc: func() runtime.Object { return &metafirst.FirstList{} },

		PredicateFunc:             MatchFirst,
		DefaultQualifiedResource:  metafirst.Resource("firsts"),
		SingularQualifiedResource: metafirst.Resource("first"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(metafirst.Resource("firsts")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{Store: store}, nil
}
