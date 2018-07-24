/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package polymorphichelpers

import (
	"bytes"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

func TestDefaultObjectResumer(t *testing.T) {
	tests := []struct {
		object    runtime.Object
		notHave   []byte
		expectErr bool
	}{
		{
			object: &extensions.Deployment{
				Spec: extensions.DeploymentSpec{
					Paused: true,
				},
			},
			notHave:   []byte(`paused":true`),
			expectErr: false,
		},
		{
			object: &extensions.Deployment{
				Spec: extensions.DeploymentSpec{
					Paused: false,
				},
			},
			expectErr: true,
		},
		{
			object:    &extensions.ReplicaSet{},
			expectErr: true,
		},
	}

	for _, test := range tests {
		actual, err := defaultObjectResumer(test.object)
		if test.expectErr {
			if err == nil {
				t.Error("unexpected non-error")
			}
			continue
		}
		if !test.expectErr && err != nil {
			t.Errorf("unexpected error: %v", err)
			continue
		}
		if bytes.Contains(actual, test.notHave) {
			t.Errorf("expected to not have %s, but got %s", test.notHave, actual)
		}
	}
}
