// This module contains a set of general type definitions that
// are used across OpenConfig models. It can be imported by modules
// that make use of these types.
package types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package types"))
}

type L2ETHERNET struct {
}

func (id L2ETHERNET) String() string {
	return "openconfig-types:L2_ETHERNET"
}

type IPV6 struct {
}

func (id IPV6) String() string {
	return "openconfig-types:IPV6"
}

type MPLS struct {
}

func (id MPLS) String() string {
	return "openconfig-types:MPLS"
}

type ADDRESSFAMILY struct {
}

func (id ADDRESSFAMILY) String() string {
	return "openconfig-types:ADDRESS_FAMILY"
}

type IPV4 struct {
}

func (id IPV4) String() string {
	return "openconfig-types:IPV4"
}

