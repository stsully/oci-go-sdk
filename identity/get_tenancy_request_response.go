// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetTenancyRequest wrapper for the GetTenancy operation
type GetTenancyRequest struct {

	// The OCID of the tenancy.
	TenancyID *string `mandatory:"true" contributesTo:"path" name:"tenancyId"`
}

func (request GetTenancyRequest) String() string {
	return common.PointerString(request)
}

// GetTenancyResponse wrapper for the GetTenancy operation
type GetTenancyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Tenancy instance
	Tenancy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTenancyResponse) String() string {
	return common.PointerString(response)
}