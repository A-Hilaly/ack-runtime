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

package types

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcfg "github.com/aws-controllers-k8s/runtime/pkg/config"
	ackmetrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"
)

// AWSResourceManager is responsible for providing a consistent way to perform
// CRUD+L operations in a backend AWS service API for Kubernetes custom
// resources (CR) corresponding to those AWS service API resources.
//
// Use an AWSResourceManagerFactory to create an AWSResourceManager for a
// particular APIResource and AWS account.
type AWSResourceManager interface {
	ReadOne(context.Context, AWSResource) (AWSResource, error)
	Create(context.Context, AWSResource) (AWSResource, error)
	Update(context.Context, AWSResource, AWSResource, *ackcompare.Delta) (AWSResource, error)
	Delete(context.Context, AWSResource) (AWSResource, error)
	ARNFromName(string) string
	LateInitialize(context.Context, AWSResource) (AWSResource, error)
	ResolveReferences(context.Context, client.Reader, AWSResource) (AWSResource, error)
	IsSynced(context.Context, AWSResource) (bool, error)
	EnsureTags(context.Context, AWSResource, ServiceControllerMetadata) error
}

// AWSResourceManagerFactory returns an AWSResourceManager that can be used to
// manage AWS resources for a particular AWS account
// TODO(jaypipes): Move AWSResourceManagerFactory into its own file
type AWSResourceManagerFactory interface {
	// ResourceDescriptor returns an AWSResourceDescriptor that can be used by
	// the upstream controller-runtime to introspect the CRs that the resource
	// manager will manage as well as produce Kubernetes runtime object
	// prototypes
	ResourceDescriptor() AWSResourceDescriptor
	// ManagerFor returns an AWSResourceManager that manages AWS resources on
	// behalf of a particular AWS account and in a specific AWS region
	ManagerFor(
		ackcfg.Config, // passed by-value to avoid mutation by consumers
		logr.Logger,
		*ackmetrics.Metrics,
		Reconciler,
		*session.Session,
		ackv1alpha1.AWSAccountID,
		ackv1alpha1.AWSRegion,
	) (AWSResourceManager, error)
	// IsAdoptable returns true if the resource is able to be adopted
	IsAdoptable() bool
	// RequeueOnSuccessSeconds returns true if the resource should be requeued after specified seconds
	// Default is false which means resource will not be requeued after success.
	RequeueOnSuccessSeconds() int
}
