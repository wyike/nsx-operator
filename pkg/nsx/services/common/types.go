/* Copyright © 2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

package common

import (
	"time"

	"github.com/openlyinc/pointy"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware-tanzu/nsx-operator/pkg/config"
	"github.com/vmware-tanzu/nsx-operator/pkg/nsx"
)

const (
	HashLength                   int    = 8
	MaxTagLength                 int    = 256
	TagScopeCluster              string = "nsx-op/cluster"
	TagScopeNamespace            string = "nsx-op/namespace"
	TagScopeSecurityPolicyCRName string = "nsx-op/security_policy_cr_name"
	TagScopeSecurityPolicyCRUID  string = "nsx-op/security_policy_cr_uid"
	TagScopeRuleID               string = "nsx-op/rule_id"
	TagScopeGroupType            string = "nsx-op/group_type"
	TagScopeSelectorHash         string = "nsx-op/selector_hash"
	TagScopeNCPCluster           string = "ncp/cluster"
	TagScopeNCPProject           string = "ncp/project"
	TagScopeNCPVIFProject        string = "ncp/vif_project"
	TagScopeNCPPod               string = "ncp/pod"
	TagScopeNCPVNETInterface     string = "ncp/vnet_interface"

	GCInterval    = 60 * time.Second
	FinalizerName = "securitypolicy.nsx.vmware.com/finalizer"
)

var (
	ResourceType               string = "resource_type"
	ResourceTypeSecurityPolicy string = "SecurityPolicy"
	ResourceTypeGroup          string = "Group"
	ResourceTypeRule           string = "Rule"
)

type Service struct {
	Client           client.Client
	NSXClient        *nsx.Client
	NSXConfig        *config.NSXOperatorConfig
	ResourceCacheMap map[string]cache.Indexer // Cache for all resources by resource type
}

func NewConverter() *bindings.TypeConverter {
	converter := bindings.NewTypeConverter()
	converter.SetMode(bindings.REST)
	return converter
}

var (
	String = pointy.String // address of string
	Int64  = pointy.Int64  // address of int64
)