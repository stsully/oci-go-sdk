// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// IdpGroupMapping. A mapping between a single group defined by the identity provider (IdP) you're federating with
// and a single IAM Service Group in Oracle Bare Metal Cloud
// Services. For more information about group mappings and what they're for, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
// A given IdP group can be mapped to zero, one, or multiple IAM Service groups, and vice versa.
// But each `IdPGroupMapping` object is between only a single IdP group and IAM Service group.
// Each `IdPGroupMapping` object has its own OCID.
// **Note:** Any users who are in more than 50 IdP groups cannot be authenticated to use the Oracle
// Bare Metal Cloud Services Console.
type IdpGroupMapping struct {

	// The OCID of the `IdpGroupMapping`.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the `IdentityProvider` this mapping belongs to.
	IdpID *string `mandatory:"true" json:"idpId,omitempty"`

	// The name of the IdP group that is mapped to the IAM Service group.
	IdpGroupName *string `mandatory:"true" json:"idpGroupName,omitempty"`

	// The OCID of the IAM Service group that is mapped to the IdP group.
	GroupID *string `mandatory:"true" json:"groupId,omitempty"`

	// The OCID of the tenancy containing the `IdentityProvider`.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// Date and time the mapping was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The mapping's current state.  After creating a mapping object, make sure its `lifecycleState` changes
	// from CREATING to ACTIVE before using it.
	LifecycleState *string `mandatory:"true" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (model IdpGroupMapping) String() string {
	return common.PointerString(model)
}