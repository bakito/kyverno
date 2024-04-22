/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// DryRunOptionApplyConfiguration represents an declarative configuration of the DryRunOption type for use
// with apply.
type DryRunOptionApplyConfiguration struct {
	Enable    *bool   `json:"enable,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// DryRunOptionApplyConfiguration constructs an declarative configuration of the DryRunOption type for use with
// apply.
func DryRunOption() *DryRunOptionApplyConfiguration {
	return &DryRunOptionApplyConfiguration{}
}

// WithEnable sets the Enable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enable field is set to the value of the last call.
func (b *DryRunOptionApplyConfiguration) WithEnable(value bool) *DryRunOptionApplyConfiguration {
	b.Enable = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *DryRunOptionApplyConfiguration) WithNamespace(value string) *DryRunOptionApplyConfiguration {
	b.Namespace = &value
	return b
}