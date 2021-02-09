// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package v1alpha1

// AWSIdentifiers provide all unique ways to reference an AWS resource.
type AWSIdentifiers struct {
	// ARN is the AWS Resource Name for the resource. It is a globally
	// unique identifier.
	ARN *AWSResourceName `json:"arn,omitempty"`
	// Name is a user-supplied string identifier for the resource. It may
	// or may not be globally unique, depending on the type of resource.
	Name *string `json:"name,omitempty"`
	// ID is an identifier that has been generated by the AWS service for
	// the resource. ID fields are *usually*, but not always, globally unique.
	// Sometimes ID fields are numeric, other times they are strings.
	ID *string `json:"id,omitempty"`
}

// TargetKubernetesResource provides all the values necessary to identify a given ACK type
// and override any metadata values when creating a resource of that type.
type TargetKubernetesResource struct {
	// +kubebuilder:validation:Required
	Group *string `json:"group"`
	// +kubebuilder:validation:Required
	Kind     *string            `json:"kind"`
	Metadata *PartialObjectMeta `json:"metadata,omitempty"`
}
