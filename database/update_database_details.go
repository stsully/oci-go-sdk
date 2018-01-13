// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

type UpdateDatabaseDetails struct {
	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig,omitempty"`
}

func (m UpdateDatabaseDetails) String() string {
	return common.PointerString(m)
}
