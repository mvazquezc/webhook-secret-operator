package routes

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
)

// RouteGetter implementations get the URL for OpenShift Routes.
type RouteGetter interface {
	// RouteURL gets the full URL to access this Route, adding an optional path
	// element to the URL.
	RouteURL(context.Context, types.NamespacedName, string) (string, error)
}
