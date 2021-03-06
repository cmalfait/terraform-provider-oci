// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateNamespaceMetadataRequest wrapper for the UpdateNamespaceMetadata operation
type UpdateNamespaceMetadataRequest struct {

	// The top-level namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Request object for update NamespaceMetadata.
	UpdateNamespaceMetadataDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`
}

func (request UpdateNamespaceMetadataRequest) String() string {
	return common.PointerString(request)
}

// UpdateNamespaceMetadataResponse wrapper for the UpdateNamespaceMetadata operation
type UpdateNamespaceMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The NamespaceMetadata instance
	NamespaceMetadata `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateNamespaceMetadataResponse) String() string {
	return common.PointerString(response)
}
