package first

import (
	"context"
	"fmt"
	"github.com/Dimss/first/apis/metafirst"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

type firstStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func newStrategy(typer runtime.ObjectTyper) firstStrategy {
	return firstStrategy{typer, names.SimpleNameGenerator}
}

func MatchFirst(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*metafirst.First)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a First")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *metafirst.First) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

func (f firstStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (f firstStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {

}

func (f firstStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (f firstStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (f firstStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (f firstStrategy) NamespaceScoped() bool {
	return false
}

func (f firstStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {

}

func (f firstStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (f firstStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func (f firstStrategy) Canonicalize(obj runtime.Object) {

}
