// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

type CreateSwiftPasswordDetails struct {

	// The description you assign to the Swift password during creation. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`
}

func (m CreateSwiftPasswordDetails) String() string {
	return common.PointerString(m)
}
