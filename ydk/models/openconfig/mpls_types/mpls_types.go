// General types for MPLS / TE data model
package mpls_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_types"))
}

type PathSetupProtocol struct {
}

func (id PathSetupProtocol) String() string {
	return "openconfig-mpls-types:path-setup-protocol"
}

type PathSetupRsvp struct {
}

func (id PathSetupRsvp) String() string {
	return "openconfig-mpls-types:path-setup-rsvp"
}

type PathSetupSr struct {
}

func (id PathSetupSr) String() string {
	return "openconfig-mpls-types:path-setup-sr"
}

type PathSetupLdp struct {
}

func (id PathSetupLdp) String() string {
	return "openconfig-mpls-types:path-setup-ldp"
}

type ProtectionType struct {
}

func (id ProtectionType) String() string {
	return "openconfig-mpls-types:protection-type"
}

type Unprotected struct {
}

func (id Unprotected) String() string {
	return "openconfig-mpls-types:unprotected"
}

type LinkProtectionRequested struct {
}

func (id LinkProtectionRequested) String() string {
	return "openconfig-mpls-types:link-protection-requested"
}

type LinkNodeProtectionRequested struct {
}

func (id LinkNodeProtectionRequested) String() string {
	return "openconfig-mpls-types:link-node-protection-requested"
}

type LspRole struct {
}

func (id LspRole) String() string {
	return "openconfig-mpls-types:lsp-role"
}

type INGRESS struct {
}

func (id INGRESS) String() string {
	return "openconfig-mpls-types:INGRESS"
}

type EGRESS struct {
}

func (id EGRESS) String() string {
	return "openconfig-mpls-types:EGRESS"
}

type TRANSIT struct {
}

func (id TRANSIT) String() string {
	return "openconfig-mpls-types:TRANSIT"
}

type TunnelType struct {
}

func (id TunnelType) String() string {
	return "openconfig-mpls-types:tunnel-type"
}

type P2P struct {
}

func (id P2P) String() string {
	return "openconfig-mpls-types:P2P"
}

type P2MP struct {
}

func (id P2MP) String() string {
	return "openconfig-mpls-types:P2MP"
}

type LspOperStatus struct {
}

func (id LspOperStatus) String() string {
	return "openconfig-mpls-types:lsp-oper-status"
}

type DOWN struct {
}

func (id DOWN) String() string {
	return "openconfig-mpls-types:DOWN"
}

type UP struct {
}

func (id UP) String() string {
	return "openconfig-mpls-types:UP"
}

type TunnelAdminStatus struct {
}

func (id TunnelAdminStatus) String() string {
	return "openconfig-mpls-types:tunnel-admin-status"
}

type ADMINDOWN struct {
}

func (id ADMINDOWN) String() string {
	return "openconfig-mpls-types:ADMIN_DOWN"
}

type ADMINUP struct {
}

func (id ADMINUP) String() string {
	return "openconfig-mpls-types:ADMIN_UP"
}

type NullLabelType struct {
}

func (id NullLabelType) String() string {
	return "openconfig-mpls-types:null-label-type"
}

type EXPLICIT struct {
}

func (id EXPLICIT) String() string {
	return "openconfig-mpls-types:EXPLICIT"
}

type IMPLICIT struct {
}

func (id IMPLICIT) String() string {
	return "openconfig-mpls-types:IMPLICIT"
}

type PATHSETUPPROTOCOL struct {
}

func (id PATHSETUPPROTOCOL) String() string {
	return "openconfig-mpls-types:PATH_SETUP_PROTOCOL"
}

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
)

// TunnelType_ represents defines the tunnel type for the LSP
type TunnelType_ string

const (
    // point-to-point label-switched-path
    TunnelType__P2P TunnelType_ = "P2P"

    // point-to-multipoint label-switched-path
    TunnelType__P2MP TunnelType_ = "P2MP"

    // multipoint-to-multipoint label-switched-path
    TunnelType__MP2MP TunnelType_ = "MP2MP"
)

