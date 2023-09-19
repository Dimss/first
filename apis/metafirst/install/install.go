package install

import (
	"github.com/Dimss/first/apis/metafirst"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(metafirst.AddToScheme(scheme))
}
