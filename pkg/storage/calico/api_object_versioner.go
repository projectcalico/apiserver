// Copyright (c) 2019 Tigera, Inc. All rights reserved.

package calico

import (
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	etcd "k8s.io/apiserver/pkg/storage/etcd3"
)

// APIObjectVersioner implements versioning and extracting etcd node information
// for objects that have an embedded ObjectMeta or ListMeta field.
type APIObjectVersioner struct {
	*etcd.APIObjectVersioner
}

// ObjectResourceVersion implements Versioner
func (a APIObjectVersioner) ObjectResourceVersion(obj runtime.Object) (uint64, error) {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return 0, err
	}
	version := accessor.GetResourceVersion()
	if len(version) == 0 {
		return 0, nil
	}
	if strings.ContainsRune(version, '/') == true {
		revs := strings.Split(version, "/")
		crdNPRev, k8sNPRev := revs[0], revs[1]
		if crdNPRev == "" && k8sNPRev != "" {
			reason := "kubernetes network policies must be managed through the kubernetes API"
			return 0, errors.NewBadRequest(reason)
		}
		version = crdNPRev
	}
	return strconv.ParseUint(version, 10, 64)
}
