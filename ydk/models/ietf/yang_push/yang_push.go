// This module contains conceptual YANG specifications
// for YANG push.
package yang_push

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package yang_push"))
}

type CustomStream struct {
}

func (id CustomStream) String() string {
	return "ietf-yang-push:custom-stream"
}

type YangPush struct {
}

func (id YangPush) String() string {
	return "ietf-yang-push:yang-push"
}

type ErrorDataNotAuthorized struct {
}

func (id ErrorDataNotAuthorized) String() string {
	return "ietf-yang-push:error-data-not-authorized"
}

type Http2 struct {
}

func (id Http2) String() string {
	return "ietf-yang-push:http2"
}

// ChangeType represents to a datastore.
type ChangeType string

const (
    // A new data node was created
    ChangeType_create ChangeType = "create"

    // A data node was deleted
    ChangeType_delete_ ChangeType = "delete"

    // The value of a data node has changed
    ChangeType_modify ChangeType = "modify"
)

