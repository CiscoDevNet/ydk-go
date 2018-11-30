// Types associated with a network instance
package network_instance_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package network_instance_types"))
}

type INSTANCELABEL struct {
}

func (id INSTANCELABEL) String() string {
	return "openconfig-network-instance-types:INSTANCE_LABEL"
}

type L2VSI struct {
}

func (id L2VSI) String() string {
	return "openconfig-network-instance-types:L2VSI"
}

type L3VRF struct {
}

func (id L3VRF) String() string {
	return "openconfig-network-instance-types:L3VRF"
}

type REMOTE struct {
}

func (id REMOTE) String() string {
	return "openconfig-network-instance-types:REMOTE"
}

type LDP struct {
}

func (id LDP) String() string {
	return "openconfig-network-instance-types:LDP"
}

type PERPREFIX struct {
}

func (id PERPREFIX) String() string {
	return "openconfig-network-instance-types:PER_PREFIX"
}

type SIGNALLINGPROTOCOL struct {
}

func (id SIGNALLINGPROTOCOL) String() string {
	return "openconfig-network-instance-types:SIGNALLING_PROTOCOL"
}

type MPLS struct {
}

func (id MPLS) String() string {
	return "openconfig-network-instance-types:MPLS"
}

type BGPVPLS struct {
}

func (id BGPVPLS) String() string {
	return "openconfig-network-instance-types:BGP_VPLS"
}

type BGPEVPN struct {
}

func (id BGPEVPN) String() string {
	return "openconfig-network-instance-types:BGP_EVPN"
}

type LABELALLOCATIONMODE struct {
}

func (id LABELALLOCATIONMODE) String() string {
	return "openconfig-network-instance-types:LABEL_ALLOCATION_MODE"
}

type ENDPOINTTYPE struct {
}

func (id ENDPOINTTYPE) String() string {
	return "openconfig-network-instance-types:ENDPOINT_TYPE"
}

type ENCAPSULATION struct {
}

func (id ENCAPSULATION) String() string {
	return "openconfig-network-instance-types:ENCAPSULATION"
}

type L2P2P struct {
}

func (id L2P2P) String() string {
	return "openconfig-network-instance-types:L2P2P"
}

type NETWORKINSTANCETYPE struct {
}

func (id NETWORKINSTANCETYPE) String() string {
	return "openconfig-network-instance-types:NETWORK_INSTANCE_TYPE"
}

type DEFAULTINSTANCE struct {
}

func (id DEFAULTINSTANCE) String() string {
	return "openconfig-network-instance-types:DEFAULT_INSTANCE"
}

type LOCAL struct {
}

func (id LOCAL) String() string {
	return "openconfig-network-instance-types:LOCAL"
}

type L2L3 struct {
}

func (id L2L3) String() string {
	return "openconfig-network-instance-types:L2L3"
}

type VXLAN struct {
}

func (id VXLAN) String() string {
	return "openconfig-network-instance-types:VXLAN"
}

type PERNEXTHOP struct {
}

func (id PERNEXTHOP) String() string {
	return "openconfig-network-instance-types:PER_NEXTHOP"
}

