/*


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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
)

// AzurePlatformSpecApplyConfiguration represents an declarative configuration of the AzurePlatformSpec type for use
// with apply.
type AzurePlatformSpecApplyConfiguration struct {
	Credentials       *v1.LocalObjectReference `json:"credentials,omitempty"`
	Cloud             *string                  `json:"cloud,omitempty"`
	Location          *string                  `json:"location,omitempty"`
	ResourceGroupName *string                  `json:"resourceGroup,omitempty"`
	VnetID            *string                  `json:"vnetID,omitempty"`
	SubnetID          *string                  `json:"subnetID,omitempty"`
	SubscriptionID    *string                  `json:"subscriptionID,omitempty"`
	SecurityGroupID   *string                  `json:"securityGroupID,omitempty"`
}

// AzurePlatformSpecApplyConfiguration constructs an declarative configuration of the AzurePlatformSpec type for use with
// apply.
func AzurePlatformSpec() *AzurePlatformSpecApplyConfiguration {
	return &AzurePlatformSpecApplyConfiguration{}
}

// WithCredentials sets the Credentials field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Credentials field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithCredentials(value v1.LocalObjectReference) *AzurePlatformSpecApplyConfiguration {
	b.Credentials = &value
	return b
}

// WithCloud sets the Cloud field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cloud field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithCloud(value string) *AzurePlatformSpecApplyConfiguration {
	b.Cloud = &value
	return b
}

// WithLocation sets the Location field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Location field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithLocation(value string) *AzurePlatformSpecApplyConfiguration {
	b.Location = &value
	return b
}

// WithResourceGroupName sets the ResourceGroupName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceGroupName field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithResourceGroupName(value string) *AzurePlatformSpecApplyConfiguration {
	b.ResourceGroupName = &value
	return b
}

// WithVnetID sets the VnetID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VnetID field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithVnetID(value string) *AzurePlatformSpecApplyConfiguration {
	b.VnetID = &value
	return b
}

// WithSubnetID sets the SubnetID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubnetID field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithSubnetID(value string) *AzurePlatformSpecApplyConfiguration {
	b.SubnetID = &value
	return b
}

// WithSubscriptionID sets the SubscriptionID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubscriptionID field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithSubscriptionID(value string) *AzurePlatformSpecApplyConfiguration {
	b.SubscriptionID = &value
	return b
}

// WithSecurityGroupID sets the SecurityGroupID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityGroupID field is set to the value of the last call.
func (b *AzurePlatformSpecApplyConfiguration) WithSecurityGroupID(value string) *AzurePlatformSpecApplyConfiguration {
	b.SecurityGroupID = &value
	return b
}
