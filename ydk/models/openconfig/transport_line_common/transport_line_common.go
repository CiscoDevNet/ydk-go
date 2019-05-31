// This module defines common data elements for OpenConfig data
// models for optical transport line system elements, such as
// amplifiers and ROADMs (wavelength routers).
package transport_line_common

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package transport_line_common"))
}

type OPTICALLINEPORTTYPE struct {
}

func (id OPTICALLINEPORTTYPE) String() string {
	return "openconfig-transport-line-common:OPTICAL_LINE_PORT_TYPE"
}

type INGRESS struct {
}

func (id INGRESS) String() string {
	return "openconfig-transport-line-common:INGRESS"
}

type EGRESS struct {
}

func (id EGRESS) String() string {
	return "openconfig-transport-line-common:EGRESS"
}

type ADD struct {
}

func (id ADD) String() string {
	return "openconfig-transport-line-common:ADD"
}

type DROP struct {
}

func (id DROP) String() string {
	return "openconfig-transport-line-common:DROP"
}

type MONITOR struct {
}

func (id MONITOR) String() string {
	return "openconfig-transport-line-common:MONITOR"
}

