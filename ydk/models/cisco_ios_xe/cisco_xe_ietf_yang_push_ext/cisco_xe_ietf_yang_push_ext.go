// This module contains augmentations on top of IETF yang push.
package cisco_xe_ietf_yang_push_ext

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_xe_ietf_yang_push_ext"))
}

type EncodeTdl struct {
}

func (id EncodeTdl) String() string {
	return "cisco-xe-ietf-yang-push-ext:encode-tdl"
}

type TdlStream struct {
}

func (id TdlStream) String() string {
	return "cisco-xe-ietf-yang-push-ext:tdl-stream"
}

