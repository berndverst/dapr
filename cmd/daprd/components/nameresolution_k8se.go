// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation and Dapr Contributors.
// Licensed under the MIT License.
// ------------------------------------------------------------

package components

import (
	nrK8se "github.com/dapr/components-contrib/nameresolution/k8se"
	nrLoader "github.com/dapr/dapr/pkg/components/nameresolution"
)

func init() {
	nrLoader.DefaultRegistry.RegisterComponent(nrK8se.NewResolver, "k8se")
}
