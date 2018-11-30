// General types for MPLS / TE data model
package mpls_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_types"))
}

type LSPMETRICINHERITED struct {
}

func (id LSPMETRICINHERITED) String() string {
	return "openconfig-mpls-types:LSP_METRIC_INHERITED"
}

type DOWN struct {
}

func (id DOWN) String() string {
	return "openconfig-mpls-types:DOWN"
}

type LOCALLYCOMPUTED struct {
}

func (id LOCALLYCOMPUTED) String() string {
	return "openconfig-mpls-types:LOCALLY_COMPUTED"
}

type PROTECTIONTYPE struct {
}

func (id PROTECTIONTYPE) String() string {
	return "openconfig-mpls-types:PROTECTION_TYPE"
}

type P2P struct {
}

func (id P2P) String() string {
	return "openconfig-mpls-types:P2P"
}

type PATHSETUPSR struct {
}

func (id PATHSETUPSR) String() string {
	return "openconfig-mpls-types:PATH_SETUP_SR"
}

type INGRESS struct {
}

func (id INGRESS) String() string {
	return "openconfig-mpls-types:INGRESS"
}

type PATHSETUPRSVP struct {
}

func (id PATHSETUPRSVP) String() string {
	return "openconfig-mpls-types:PATH_SETUP_RSVP"
}

type PATHSETUPLDP struct {
}

func (id PATHSETUPLDP) String() string {
	return "openconfig-mpls-types:PATH_SETUP_LDP"
}

type EXPLICIT struct {
}

func (id EXPLICIT) String() string {
	return "openconfig-mpls-types:EXPLICIT"
}

type ADMINUP struct {
}

func (id ADMINUP) String() string {
	return "openconfig-mpls-types:ADMIN_UP"
}

type PATHSETUPPROTOCOL struct {
}

func (id PATHSETUPPROTOCOL) String() string {
	return "openconfig-mpls-types:PATH_SETUP_PROTOCOL"
}

type P2MP struct {
}

func (id P2MP) String() string {
	return "openconfig-mpls-types:P2MP"
}

type PATHCOMPUTATIONMETHOD struct {
}

func (id PATHCOMPUTATIONMETHOD) String() string {
	return "openconfig-mpls-types:PATH_COMPUTATION_METHOD"
}

type LINKPROTECTIONREQUIRED struct {
}

func (id LINKPROTECTIONREQUIRED) String() string {
	return "openconfig-mpls-types:LINK_PROTECTION_REQUIRED"
}

type TUNNELADMINSTATUS struct {
}

func (id TUNNELADMINSTATUS) String() string {
	return "openconfig-mpls-types:TUNNEL_ADMIN_STATUS"
}

type TUNNELTYPE struct {
}

func (id TUNNELTYPE) String() string {
	return "openconfig-mpls-types:TUNNEL_TYPE"
}

type EGRESS struct {
}

func (id EGRESS) String() string {
	return "openconfig-mpls-types:EGRESS"
}

type NULLLABELTYPE struct {
}

func (id NULLLABELTYPE) String() string {
	return "openconfig-mpls-types:NULL_LABEL_TYPE"
}

type LSPOPERSTATUS struct {
}

func (id LSPOPERSTATUS) String() string {
	return "openconfig-mpls-types:LSP_OPER_STATUS"
}

type LSPMETRICTYPE struct {
}

func (id LSPMETRICTYPE) String() string {
	return "openconfig-mpls-types:LSP_METRIC_TYPE"
}

type IMPLICIT struct {
}

func (id IMPLICIT) String() string {
	return "openconfig-mpls-types:IMPLICIT"
}

type LINKNODEPROTECTIONREQUESTED struct {
}

func (id LINKNODEPROTECTIONREQUESTED) String() string {
	return "openconfig-mpls-types:LINK_NODE_PROTECTION_REQUESTED"
}

type LSPMETRICABSOLUTE struct {
}

func (id LSPMETRICABSOLUTE) String() string {
	return "openconfig-mpls-types:LSP_METRIC_ABSOLUTE"
}

type EXTERNALLYQUERIED struct {
}

func (id EXTERNALLYQUERIED) String() string {
	return "openconfig-mpls-types:EXTERNALLY_QUERIED"
}

type ADMINDOWN struct {
}

func (id ADMINDOWN) String() string {
	return "openconfig-mpls-types:ADMIN_DOWN"
}

type EXPLICITLYDEFINED struct {
}

func (id EXPLICITLYDEFINED) String() string {
	return "openconfig-mpls-types:EXPLICITLY_DEFINED"
}

type TRANSIT struct {
}

func (id TRANSIT) String() string {
	return "openconfig-mpls-types:TRANSIT"
}

type UP struct {
}

func (id UP) String() string {
	return "openconfig-mpls-types:UP"
}

type LSPMETRICRELATIVE struct {
}

func (id LSPMETRICRELATIVE) String() string {
	return "openconfig-mpls-types:LSP_METRIC_RELATIVE"
}

type UNPROTECTED struct {
}

func (id UNPROTECTED) String() string {
	return "openconfig-mpls-types:UNPROTECTED"
}

type LSPROLE struct {
}

func (id LSPROLE) String() string {
	return "openconfig-mpls-types:LSP_ROLE"
}

// TunnelType represents defines the tunnel type for the LSP
type TunnelType string

const (
    // point-to-point label-switched-path
    TunnelType_P2P TunnelType = "P2P"

    // point-to-multipoint label-switched-path
    TunnelType_P2MP TunnelType = "P2MP"

    // multipoint-to-multipoint label-switched-path
    TunnelType_MP2MP TunnelType = "MP2MP"
)

// MplsLabel represents type for MPLS label value encoding
type MplsLabel string

const (
    // valid at the bottom of the label stack,
    // indicates that stack must be popped and packet forwarded
    // based on IPv4 header
    MplsLabel_IPV4_EXPLICIT_NULL MplsLabel = "IPV4_EXPLICIT_NULL"

    // allowed anywhere in the label stack except
    // the bottom, local router delivers packet to the local CPU
    // when this label is at the top of the stack
    MplsLabel_ROUTER_ALERT MplsLabel = "ROUTER_ALERT"

    // valid at the bottom of the label stack,
    // indicates that stack must be popped and packet forwarded
    // based on IPv6 header
    MplsLabel_IPV6_EXPLICIT_NULL MplsLabel = "IPV6_EXPLICIT_NULL"

    // assigned by local LSR but not carried in
    // packets
    MplsLabel_IMPLICIT_NULL MplsLabel = "IMPLICIT_NULL"

    // Entropy label indicator, to allow an LSR
    // to distinguish between entropy label and applicaiton
    // labels RFC 6790
    MplsLabel_ENTROPY_LABEL_INDICATOR MplsLabel = "ENTROPY_LABEL_INDICATOR"

    // This value is utilised to indicate that the packet that
    // is forwarded by the local system does not have an MPLS
    // header applied to it. Typically, this is used at the
    // egress of an LSP
    MplsLabel_NO_LABEL MplsLabel = "NO_LABEL"
)

