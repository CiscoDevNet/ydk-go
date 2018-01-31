// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-ldp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mpls-ldp: MPLS LDP configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_ldp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_ldp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-ldp-cfg mpls-ldp}", reflect.TypeOf(MplsLdp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp", reflect.TypeOf(MplsLdp{}))
}

// MplsLdpLabelAdvertise represents Mpls ldp label advertise
type MplsLdpLabelAdvertise string

const (
    // Advertise label for prefix(es) permitted by
    // prefix ACL
    MplsLdpLabelAdvertise_for_ MplsLdpLabelAdvertise = "for"

    // Advertise label for prefix(es) permitted by
    // prefix ACL to peer(s) permitted by peer ACL
    MplsLdpLabelAdvertise_for_to MplsLdpLabelAdvertise = "for-to"
)

// MldpPolicyMode represents Mldp policy mode
type MldpPolicyMode string

const (
    // Inbound route policy
    MldpPolicyMode_inbound MldpPolicyMode = "inbound"

    // Outbound route policy
    MldpPolicyMode_outbound MldpPolicyMode = "outbound"
)

// MplsLdpLabelAllocation represents Mpls ldp label allocation
type MplsLdpLabelAllocation string

const (
    // Allocate label for prefixes permitted by ACL
    MplsLdpLabelAllocation_acl MplsLdpLabelAllocation = "acl"

    // Allocate label for host routes only
    MplsLdpLabelAllocation_host MplsLdpLabelAllocation = "host"
)

// MplsLdpTargetedAccept represents Mpls ldp targeted accept
type MplsLdpTargetedAccept string

const (
    // Accept targeted hello from all
    MplsLdpTargetedAccept_all MplsLdpTargetedAccept = "all"

    // Accept targeted hello from peer ACL
    MplsLdpTargetedAccept_from MplsLdpTargetedAccept = "from"
)

// MplsLdpNbrPassword represents Mpls ldp nbr password
type MplsLdpNbrPassword string

const (
    // Disable the global default password for this
    // neighbor
    MplsLdpNbrPassword_disable MplsLdpNbrPassword = "disable"

    // Specify a password for this neighbor
    MplsLdpNbrPassword_specified MplsLdpNbrPassword = "specified"
)

// MplsLdpDownstreamOnDemand represents Mpls ldp downstream on demand
type MplsLdpDownstreamOnDemand string

const (
    // Downstream on Demand peers permitted by ACL
    MplsLdpDownstreamOnDemand_peer_acl MplsLdpDownstreamOnDemand = "peer-acl"
)

// MplsLdpExpNull represents Mpls ldp exp null
type MplsLdpExpNull string

const (
    // Advertise explicit-null for all connected
    // prefixes to all peers
    MplsLdpExpNull_all MplsLdpExpNull = "all"

    // Advertise explicit-null for prefix(es)
    // permitted by prefix ACL
    MplsLdpExpNull_for_ MplsLdpExpNull = "for"

    // Advertise explicit-null for all connected
    // prefixes to peer(s) permitted by peer ACL
    MplsLdpExpNull_to MplsLdpExpNull = "to"

    // Advertise explicit-null for prefix(es)
    // permitted by prefix ACL to peer(s) permitted by
    // peer ACL
    MplsLdpExpNull_for_to MplsLdpExpNull = "for-to"
)

// MplsLdpAdvertiseBgpAcl represents Mpls ldp advertise bgp acl
type MplsLdpAdvertiseBgpAcl string

const (
    // BGP prefixes advertised to peers permitted by
    // ACL
    MplsLdpAdvertiseBgpAcl_peer_acl MplsLdpAdvertiseBgpAcl = "peer-acl"
)

// MplsLdpafName represents Mpls ldpaf name
type MplsLdpafName string

const (
    // IPv4
    MplsLdpafName_ipv4 MplsLdpafName = "ipv4"

    // IPv6
    MplsLdpafName_ipv6 MplsLdpafName = "ipv6"
)

// MplsLdpTransportAddress represents Mpls ldp transport address
type MplsLdpTransportAddress string

const (
    // Use interface IP address
    MplsLdpTransportAddress_interface_ MplsLdpTransportAddress = "interface"

    // Use given IP address
    MplsLdpTransportAddress_address MplsLdpTransportAddress = "address"
)

// MplsLdpSessionProtection represents Mpls ldp session protection
type MplsLdpSessionProtection string

const (
    // Protect all peer sessions
    MplsLdpSessionProtection_all MplsLdpSessionProtection = "all"

    // Protect peer session(s) permitted by peer ACL
    MplsLdpSessionProtection_for_ MplsLdpSessionProtection = "for"

    // Protect all peer sessions and holdup protection
    // for given duration
    MplsLdpSessionProtection_all_with_duration MplsLdpSessionProtection = "all-with-duration"

    // Protect peer session(s) permitted by peer ACL
    // and holdup protection for given duration
    MplsLdpSessionProtection_for_with_duration MplsLdpSessionProtection = "for-with-duration"

    // Protect all peer sessions and holdup protection
    // forever
    MplsLdpSessionProtection_all_with_forever MplsLdpSessionProtection = "all-with-forever"

    // Protect peer session(s) permitted by peer ACL
    // and holdup protection forever
    MplsLdpSessionProtection_for_with_forever MplsLdpSessionProtection = "for-with-forever"
)

// MplsLdp
// MPLS LDP configuration
type MplsLdp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Label Distribution Protocol (LDP) globally.Without creating this
    // object the LDP feature will not be enabled. Deleting this object will stop
    // the LDP feature. The type is interface{}.
    Enable interface{}

    // Global VRF attribute configuration for MPLS LDP.
    DefaultVrf MplsLdp_DefaultVrf

    // VRF Table attribute configuration for MPLS LDP.
    Vrfs MplsLdp_Vrfs

    // Global configuration for MPLS LDP.
    Global MplsLdp_Global
}

func (mplsLdp *MplsLdp) GetFilter() yfilter.YFilter { return mplsLdp.YFilter }

func (mplsLdp *MplsLdp) SetFilter(yf yfilter.YFilter) { mplsLdp.YFilter = yf }

func (mplsLdp *MplsLdp) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "default-vrf" { return "DefaultVrf" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "global" { return "Global" }
    return ""
}

func (mplsLdp *MplsLdp) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp"
}

func (mplsLdp *MplsLdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-vrf" {
        return &mplsLdp.DefaultVrf
    }
    if childYangName == "vrfs" {
        return &mplsLdp.Vrfs
    }
    if childYangName == "global" {
        return &mplsLdp.Global
    }
    return nil
}

func (mplsLdp *MplsLdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-vrf"] = &mplsLdp.DefaultVrf
    children["vrfs"] = &mplsLdp.Vrfs
    children["global"] = &mplsLdp.Global
    return children
}

func (mplsLdp *MplsLdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = mplsLdp.Enable
    return leafs
}

func (mplsLdp *MplsLdp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLdp *MplsLdp) GetYangName() string { return "mpls-ldp" }

func (mplsLdp *MplsLdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLdp *MplsLdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLdp *MplsLdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLdp *MplsLdp) SetParent(parent types.Entity) { mplsLdp.parent = parent }

func (mplsLdp *MplsLdp) GetParent() types.Entity { return mplsLdp.parent }

func (mplsLdp *MplsLdp) GetParentYangName() string { return "Cisco-IOS-XR-mpls-ldp-cfg" }

// MplsLdp_DefaultVrf
// Global VRF attribute configuration for MPLS LDP
type MplsLdp_DefaultVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family specific configuration for MPLS LDP.
    Afs MplsLdp_DefaultVrf_Afs

    // Default VRF Global configuration for MPLS LDP.
    Global MplsLdp_DefaultVrf_Global

    // MPLS LDP configuration pertaining to interfaces.
    Interfaces MplsLdp_DefaultVrf_Interfaces
}

func (defaultVrf *MplsLdp_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *MplsLdp_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *MplsLdp_DefaultVrf) GetGoName(yname string) string {
    if yname == "afs" { return "Afs" }
    if yname == "global" { return "Global" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (defaultVrf *MplsLdp_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *MplsLdp_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &defaultVrf.Afs
    }
    if childYangName == "global" {
        return &defaultVrf.Global
    }
    if childYangName == "interfaces" {
        return &defaultVrf.Interfaces
    }
    return nil
}

func (defaultVrf *MplsLdp_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &defaultVrf.Afs
    children["global"] = &defaultVrf.Global
    children["interfaces"] = &defaultVrf.Interfaces
    return children
}

func (defaultVrf *MplsLdp_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultVrf *MplsLdp_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *MplsLdp_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *MplsLdp_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *MplsLdp_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *MplsLdp_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *MplsLdp_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *MplsLdp_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *MplsLdp_DefaultVrf) GetParentYangName() string { return "mpls-ldp" }

// MplsLdp_DefaultVrf_Afs
// Address Family specific configuration for MPLS
// LDP
type MplsLdp_DefaultVrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af.
    Af []MplsLdp_DefaultVrf_Afs_Af
}

func (afs *MplsLdp_DefaultVrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_DefaultVrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_DefaultVrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_DefaultVrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_DefaultVrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_DefaultVrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_DefaultVrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_DefaultVrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsLdp_DefaultVrf_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_DefaultVrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsLdp_DefaultVrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsLdp_DefaultVrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsLdp_DefaultVrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_DefaultVrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_DefaultVrf_Afs) GetParentYangName() string { return "default-vrf" }

// MplsLdp_DefaultVrf_Afs_Af
// Configure data for given Address Family
type MplsLdp_DefaultVrf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family type. The type is MplsLdpafName.
    AfName interface{}

    // Enable Address Family. The type is interface{}.
    Enable interface{}

    // Configure Label policies and control.
    Label MplsLdp_DefaultVrf_Afs_Af_Label

    // Configure Discovery parameters.
    Discovery MplsLdp_DefaultVrf_Afs_Af_Discovery

    // MPLS Traffic Engingeering parameters for LDP.
    TrafficEngineering MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering

    // Configuration related to Neighbors.
    Neighbor MplsLdp_DefaultVrf_Afs_Af_Neighbor

    // MPLS LDP configuration for protocol redistribution.
    RedistributionProtocol MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol
}

func (af *MplsLdp_DefaultVrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_DefaultVrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_DefaultVrf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "enable" { return "Enable" }
    if yname == "label" { return "Label" }
    if yname == "discovery" { return "Discovery" }
    if yname == "traffic-engineering" { return "TrafficEngineering" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "redistribution-protocol" { return "RedistributionProtocol" }
    return ""
}

func (af *MplsLdp_DefaultVrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_DefaultVrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label" {
        return &af.Label
    }
    if childYangName == "discovery" {
        return &af.Discovery
    }
    if childYangName == "traffic-engineering" {
        return &af.TrafficEngineering
    }
    if childYangName == "neighbor" {
        return &af.Neighbor
    }
    if childYangName == "redistribution-protocol" {
        return &af.RedistributionProtocol
    }
    return nil
}

func (af *MplsLdp_DefaultVrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label"] = &af.Label
    children["discovery"] = &af.Discovery
    children["traffic-engineering"] = &af.TrafficEngineering
    children["neighbor"] = &af.Neighbor
    children["redistribution-protocol"] = &af.RedistributionProtocol
    return children
}

func (af *MplsLdp_DefaultVrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["enable"] = af.Enable
    return leafs
}

func (af *MplsLdp_DefaultVrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsLdp_DefaultVrf_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_DefaultVrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsLdp_DefaultVrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsLdp_DefaultVrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsLdp_DefaultVrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_DefaultVrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_DefaultVrf_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_DefaultVrf_Afs_Af_Label
// Configure Label policies and control
type MplsLdp_DefaultVrf_Afs_Af_Label struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure remote/peer label policies and control.
    Remote MplsLdp_DefaultVrf_Afs_Af_Label_Remote

    // Configure local label policies and control.
    Local MplsLdp_DefaultVrf_Afs_Af_Label_Local
}

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetFilter() yfilter.YFilter { return label.YFilter }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) SetFilter(yf yfilter.YFilter) { label.YFilter = yf }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetGoName(yname string) string {
    if yname == "remote" { return "Remote" }
    if yname == "local" { return "Local" }
    return ""
}

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetSegmentPath() string {
    return "label"
}

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote" {
        return &label.Remote
    }
    if childYangName == "local" {
        return &label.Local
    }
    return nil
}

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote"] = &label.Remote
    children["local"] = &label.Local
    return children
}

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetBundleName() string { return "cisco_ios_xr" }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetYangName() string { return "label" }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) SetParent(parent types.Entity) { label.parent = parent }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetParent() types.Entity { return label.parent }

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote
// Configure remote/peer label policies and
// control
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure inbound label acceptance.
    Accept MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept
}

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetFilter() yfilter.YFilter { return remote.YFilter }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) SetFilter(yf yfilter.YFilter) { remote.YFilter = yf }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetGoName(yname string) string {
    if yname == "accept" { return "Accept" }
    return ""
}

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetSegmentPath() string {
    return "remote"
}

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accept" {
        return &remote.Accept
    }
    return nil
}

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accept"] = &remote.Accept
    return children
}

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetBundleName() string { return "cisco_ios_xr" }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetYangName() string { return "remote" }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) SetParent(parent types.Entity) { remote.parent = parent }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetParent() types.Entity { return remote.parent }

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetParentYangName() string { return "label" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept
// Configure inbound label acceptance
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration related to neighbors for inbound label acceptance.
    PeerAcceptPolicies MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
}

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetFilter() yfilter.YFilter { return accept.YFilter }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) SetFilter(yf yfilter.YFilter) { accept.YFilter = yf }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetGoName(yname string) string {
    if yname == "peer-accept-policies" { return "PeerAcceptPolicies" }
    return ""
}

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetSegmentPath() string {
    return "accept"
}

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-accept-policies" {
        return &accept.PeerAcceptPolicies
    }
    return nil
}

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-accept-policies"] = &accept.PeerAcceptPolicies
    return children
}

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetBundleName() string { return "cisco_ios_xr" }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetYangName() string { return "accept" }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) SetParent(parent types.Entity) { accept.parent = parent }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetParent() types.Entity { return accept.parent }

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetParentYangName() string { return "remote" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
// Configuration related to neighbors for
// inbound label acceptance
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control acceptance of labels from a neighbor for prefix(es) using ACL. The
    // type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy.
    PeerAcceptPolicy []MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
}

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetFilter() yfilter.YFilter { return peerAcceptPolicies.YFilter }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) SetFilter(yf yfilter.YFilter) { peerAcceptPolicies.YFilter = yf }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetGoName(yname string) string {
    if yname == "peer-accept-policy" { return "PeerAcceptPolicy" }
    return ""
}

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetSegmentPath() string {
    return "peer-accept-policies"
}

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-accept-policy" {
        for _, c := range peerAcceptPolicies.PeerAcceptPolicy {
            if peerAcceptPolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy{}
        peerAcceptPolicies.PeerAcceptPolicy = append(peerAcceptPolicies.PeerAcceptPolicy, child)
        return &peerAcceptPolicies.PeerAcceptPolicy[len(peerAcceptPolicies.PeerAcceptPolicy)-1]
    }
    return nil
}

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerAcceptPolicies.PeerAcceptPolicy {
        children[peerAcceptPolicies.PeerAcceptPolicy[i].GetSegmentPath()] = &peerAcceptPolicies.PeerAcceptPolicy[i]
    }
    return children
}

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetYangName() string { return "peer-accept-policies" }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) SetParent(parent types.Entity) { peerAcceptPolicies.parent = parent }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetParent() types.Entity { return peerAcceptPolicies.parent }

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetParentYangName() string { return "accept" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
// Control acceptance of labels from a
// neighbor for prefix(es) using ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..4294967295.
    LabelSpaceId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetFilter() yfilter.YFilter { return peerAcceptPolicy.YFilter }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) SetFilter(yf yfilter.YFilter) { peerAcceptPolicy.YFilter = yf }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetSegmentPath() string {
    return "peer-accept-policy" + "[lsr-id='" + fmt.Sprintf("%v", peerAcceptPolicy.LsrId) + "']" + "[label-space-id='" + fmt.Sprintf("%v", peerAcceptPolicy.LabelSpaceId) + "']"
}

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = peerAcceptPolicy.LsrId
    leafs["label-space-id"] = peerAcceptPolicy.LabelSpaceId
    leafs["prefix-acl-name"] = peerAcceptPolicy.PrefixAclName
    return leafs
}

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetYangName() string { return "peer-accept-policy" }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) SetParent(parent types.Entity) { peerAcceptPolicy.parent = parent }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetParent() types.Entity { return peerAcceptPolicy.parent }

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetParentYangName() string { return "peer-accept-policies" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local
// Configure local label policies and control
type MplsLdp_DefaultVrf_Afs_Af_Label_Local struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control use of implicit-null label for set of prefix(es). The type is
    // string.
    ImplicitNullOverride interface{}

    // Enable MPLS forwarding for default route. The type is interface{}.
    DefaultRoute interface{}

    // Configure outbound label advertisement.
    Advertise MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise

    // Control local label allocation for prefix(es).
    Allocate MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate
}

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetFilter() yfilter.YFilter { return local.YFilter }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) SetFilter(yf yfilter.YFilter) { local.YFilter = yf }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetGoName(yname string) string {
    if yname == "implicit-null-override" { return "ImplicitNullOverride" }
    if yname == "default-route" { return "DefaultRoute" }
    if yname == "advertise" { return "Advertise" }
    if yname == "allocate" { return "Allocate" }
    return ""
}

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetSegmentPath() string {
    return "local"
}

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "advertise" {
        return &local.Advertise
    }
    if childYangName == "allocate" {
        return &local.Allocate
    }
    return nil
}

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["advertise"] = &local.Advertise
    children["allocate"] = &local.Allocate
    return children
}

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["implicit-null-override"] = local.ImplicitNullOverride
    leafs["default-route"] = local.DefaultRoute
    return leafs
}

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetBundleName() string { return "cisco_ios_xr" }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetYangName() string { return "local" }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) SetParent(parent types.Entity) { local.parent = parent }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetParent() types.Entity { return local.parent }

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetParentYangName() string { return "label" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise
// Configure outbound label advertisement
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable label advertisement. The type is interface{}.
    Disable interface{}

    // Configure peer centric outbound label advertisement using ACL.
    PeerAdvertisePolicies MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies

    // Configure prefix centric outbound label advertisement using ACL.
    PrefixAdvertisePolicies MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies

    // Configure advertisment of explicit-null for connected prefixes.
    ExplicitNull MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull

    // Configure outbound label advertisement for an interface.
    Interfaces MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces
}

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetFilter() yfilter.YFilter { return advertise.YFilter }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) SetFilter(yf yfilter.YFilter) { advertise.YFilter = yf }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "peer-advertise-policies" { return "PeerAdvertisePolicies" }
    if yname == "prefix-advertise-policies" { return "PrefixAdvertisePolicies" }
    if yname == "explicit-null" { return "ExplicitNull" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetSegmentPath() string {
    return "advertise"
}

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-advertise-policies" {
        return &advertise.PeerAdvertisePolicies
    }
    if childYangName == "prefix-advertise-policies" {
        return &advertise.PrefixAdvertisePolicies
    }
    if childYangName == "explicit-null" {
        return &advertise.ExplicitNull
    }
    if childYangName == "interfaces" {
        return &advertise.Interfaces
    }
    return nil
}

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-advertise-policies"] = &advertise.PeerAdvertisePolicies
    children["prefix-advertise-policies"] = &advertise.PrefixAdvertisePolicies
    children["explicit-null"] = &advertise.ExplicitNull
    children["interfaces"] = &advertise.Interfaces
    return children
}

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = advertise.Disable
    return leafs
}

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetBundleName() string { return "cisco_ios_xr" }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetYangName() string { return "advertise" }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) SetParent(parent types.Entity) { advertise.parent = parent }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetParent() types.Entity { return advertise.parent }

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetParentYangName() string { return "local" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies
// Configure peer centric outbound label
// advertisement using ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control advertisement of prefix(es) using ACL. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy.
    PeerAdvertisePolicy []MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
}

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetFilter() yfilter.YFilter { return peerAdvertisePolicies.YFilter }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) SetFilter(yf yfilter.YFilter) { peerAdvertisePolicies.YFilter = yf }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetGoName(yname string) string {
    if yname == "peer-advertise-policy" { return "PeerAdvertisePolicy" }
    return ""
}

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetSegmentPath() string {
    return "peer-advertise-policies"
}

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-advertise-policy" {
        for _, c := range peerAdvertisePolicies.PeerAdvertisePolicy {
            if peerAdvertisePolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy{}
        peerAdvertisePolicies.PeerAdvertisePolicy = append(peerAdvertisePolicies.PeerAdvertisePolicy, child)
        return &peerAdvertisePolicies.PeerAdvertisePolicy[len(peerAdvertisePolicies.PeerAdvertisePolicy)-1]
    }
    return nil
}

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerAdvertisePolicies.PeerAdvertisePolicy {
        children[peerAdvertisePolicies.PeerAdvertisePolicy[i].GetSegmentPath()] = &peerAdvertisePolicies.PeerAdvertisePolicy[i]
    }
    return children
}

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetBundleName() string { return "cisco_ios_xr" }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetYangName() string { return "peer-advertise-policies" }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) SetParent(parent types.Entity) { peerAdvertisePolicies.parent = parent }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetParent() types.Entity { return peerAdvertisePolicies.parent }

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetParentYangName() string { return "advertise" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
// Control advertisement of prefix(es) using
// ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..4294967295.
    LabelSpaceId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetFilter() yfilter.YFilter { return peerAdvertisePolicy.YFilter }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) SetFilter(yf yfilter.YFilter) { peerAdvertisePolicy.YFilter = yf }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetSegmentPath() string {
    return "peer-advertise-policy" + "[lsr-id='" + fmt.Sprintf("%v", peerAdvertisePolicy.LsrId) + "']" + "[label-space-id='" + fmt.Sprintf("%v", peerAdvertisePolicy.LabelSpaceId) + "']"
}

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = peerAdvertisePolicy.LsrId
    leafs["label-space-id"] = peerAdvertisePolicy.LabelSpaceId
    leafs["prefix-acl-name"] = peerAdvertisePolicy.PrefixAclName
    return leafs
}

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetYangName() string { return "peer-advertise-policy" }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) SetParent(parent types.Entity) { peerAdvertisePolicy.parent = parent }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetParent() types.Entity { return peerAdvertisePolicy.parent }

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetParentYangName() string { return "peer-advertise-policies" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies
// Configure prefix centric outbound label
// advertisement using ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control advertisement of prefix(es) using ACL. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy.
    PrefixAdvertisePolicy []MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy
}

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetFilter() yfilter.YFilter { return prefixAdvertisePolicies.YFilter }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) SetFilter(yf yfilter.YFilter) { prefixAdvertisePolicies.YFilter = yf }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetGoName(yname string) string {
    if yname == "prefix-advertise-policy" { return "PrefixAdvertisePolicy" }
    return ""
}

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetSegmentPath() string {
    return "prefix-advertise-policies"
}

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-advertise-policy" {
        for _, c := range prefixAdvertisePolicies.PrefixAdvertisePolicy {
            if prefixAdvertisePolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy{}
        prefixAdvertisePolicies.PrefixAdvertisePolicy = append(prefixAdvertisePolicies.PrefixAdvertisePolicy, child)
        return &prefixAdvertisePolicies.PrefixAdvertisePolicy[len(prefixAdvertisePolicies.PrefixAdvertisePolicy)-1]
    }
    return nil
}

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixAdvertisePolicies.PrefixAdvertisePolicy {
        children[prefixAdvertisePolicies.PrefixAdvertisePolicy[i].GetSegmentPath()] = &prefixAdvertisePolicies.PrefixAdvertisePolicy[i]
    }
    return children
}

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetBundleName() string { return "cisco_ios_xr" }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetYangName() string { return "prefix-advertise-policies" }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) SetParent(parent types.Entity) { prefixAdvertisePolicies.parent = parent }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetParent() types.Entity { return prefixAdvertisePolicies.parent }

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetParentYangName() string { return "advertise" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy
// Control advertisement of prefix(es) using
// ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of prefix ACL. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PrefixAclName interface{}

    // Label advertise type. The type is MplsLdpLabelAdvertise.
    AdvertiseType interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetFilter() yfilter.YFilter { return prefixAdvertisePolicy.YFilter }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) SetFilter(yf yfilter.YFilter) { prefixAdvertisePolicy.YFilter = yf }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetGoName(yname string) string {
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    if yname == "advertise-type" { return "AdvertiseType" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    return ""
}

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetSegmentPath() string {
    return "prefix-advertise-policy" + "[prefix-acl-name='" + fmt.Sprintf("%v", prefixAdvertisePolicy.PrefixAclName) + "']"
}

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-acl-name"] = prefixAdvertisePolicy.PrefixAclName
    leafs["advertise-type"] = prefixAdvertisePolicy.AdvertiseType
    leafs["peer-acl-name"] = prefixAdvertisePolicy.PeerAclName
    return leafs
}

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetYangName() string { return "prefix-advertise-policy" }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) SetParent(parent types.Entity) { prefixAdvertisePolicy.parent = parent }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetParent() types.Entity { return prefixAdvertisePolicy.parent }

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetParentYangName() string { return "prefix-advertise-policies" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull
// Configure advertisment of explicit-null
// for connected prefixes.
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Explicit Null command variant. The type is MplsLdpExpNull.
    ExplicitNullType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetFilter() yfilter.YFilter { return explicitNull.YFilter }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) SetFilter(yf yfilter.YFilter) { explicitNull.YFilter = yf }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetGoName(yname string) string {
    if yname == "explicit-null-type" { return "ExplicitNullType" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    return ""
}

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetSegmentPath() string {
    return "explicit-null"
}

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["explicit-null-type"] = explicitNull.ExplicitNullType
    leafs["prefix-acl-name"] = explicitNull.PrefixAclName
    leafs["peer-acl-name"] = explicitNull.PeerAclName
    return leafs
}

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetBundleName() string { return "cisco_ios_xr" }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetYangName() string { return "explicit-null" }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) SetParent(parent types.Entity) { explicitNull.parent = parent }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetParent() types.Entity { return explicitNull.parent }

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetParentYangName() string { return "advertise" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces
// Configure outbound label advertisement for
// an interface
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control advertisement of interface's host IP address. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface.
    Interface []MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
}

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetParentYangName() string { return "advertise" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
// Control advertisement of interface's host
// IP address
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate
// Control local label allocation for
// prefix(es)
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label allocation type. The type is MplsLdpLabelAllocation.
    AllocationType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}
}

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetFilter() yfilter.YFilter { return allocate.YFilter }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) SetFilter(yf yfilter.YFilter) { allocate.YFilter = yf }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetGoName(yname string) string {
    if yname == "allocation-type" { return "AllocationType" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetSegmentPath() string {
    return "allocate"
}

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allocation-type"] = allocate.AllocationType
    leafs["prefix-acl-name"] = allocate.PrefixAclName
    return leafs
}

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetBundleName() string { return "cisco_ios_xr" }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetYangName() string { return "allocate" }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) SetParent(parent types.Entity) { allocate.parent = parent }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetParent() types.Entity { return allocate.parent }

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetParentYangName() string { return "local" }

// MplsLdp_DefaultVrf_Afs_Af_Discovery
// Configure Discovery parameters
type MplsLdp_DefaultVrf_Afs_Af_Discovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global discovery transport address for address family. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    TransportAddress interface{}

    // Configure acceptance from and responding to targeted hellos.
    TargetedHelloAccept MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept
}

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetGoName(yname string) string {
    if yname == "transport-address" { return "TransportAddress" }
    if yname == "targeted-hello-accept" { return "TargetedHelloAccept" }
    return ""
}

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "targeted-hello-accept" {
        return &discovery.TargetedHelloAccept
    }
    return nil
}

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["targeted-hello-accept"] = &discovery.TargetedHelloAccept
    return children
}

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport-address"] = discovery.TransportAddress
    return leafs
}

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetBundleName() string { return "cisco_ios_xr" }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept
// Configure acceptance from and responding to
// targeted hellos.
type MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of acceptance. The type is MplsLdpTargetedAccept.
    AcceptType interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetFilter() yfilter.YFilter { return targetedHelloAccept.YFilter }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) SetFilter(yf yfilter.YFilter) { targetedHelloAccept.YFilter = yf }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetGoName(yname string) string {
    if yname == "accept-type" { return "AcceptType" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    return ""
}

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetSegmentPath() string {
    return "targeted-hello-accept"
}

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accept-type"] = targetedHelloAccept.AcceptType
    leafs["peer-acl-name"] = targetedHelloAccept.PeerAclName
    return leafs
}

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetBundleName() string { return "cisco_ios_xr" }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetYangName() string { return "targeted-hello-accept" }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) SetParent(parent types.Entity) { targetedHelloAccept.parent = parent }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetParent() types.Entity { return targetedHelloAccept.parent }

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetParentYangName() string { return "discovery" }

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering
// MPLS Traffic Engingeering parameters for LDP
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Traffic Engineering auto-tunnel mesh parameters for LDP.
    AutoTunnelMesh MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh
}

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetFilter() yfilter.YFilter { return trafficEngineering.YFilter }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) SetFilter(yf yfilter.YFilter) { trafficEngineering.YFilter = yf }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetGoName(yname string) string {
    if yname == "auto-tunnel-mesh" { return "AutoTunnelMesh" }
    return ""
}

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetSegmentPath() string {
    return "traffic-engineering"
}

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto-tunnel-mesh" {
        return &trafficEngineering.AutoTunnelMesh
    }
    return nil
}

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto-tunnel-mesh"] = &trafficEngineering.AutoTunnelMesh
    return children
}

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetBundleName() string { return "cisco_ios_xr" }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetYangName() string { return "traffic-engineering" }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) SetParent(parent types.Entity) { trafficEngineering.parent = parent }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetParent() types.Entity { return trafficEngineering.parent }

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh
// MPLS Traffic Engineering auto-tunnel mesh
// parameters for LDP
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable all MPLS TE auto-tunnel mesh-group interfaces. The type is
    // interface{}.
    GroupAll interface{}

    // Enable interfaces in specific MPLS TE auto-tunnel mesh-groups.
    GroupIds MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds
}

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetFilter() yfilter.YFilter { return autoTunnelMesh.YFilter }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) SetFilter(yf yfilter.YFilter) { autoTunnelMesh.YFilter = yf }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetGoName(yname string) string {
    if yname == "group-all" { return "GroupAll" }
    if yname == "group-ids" { return "GroupIds" }
    return ""
}

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetSegmentPath() string {
    return "auto-tunnel-mesh"
}

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-ids" {
        return &autoTunnelMesh.GroupIds
    }
    return nil
}

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-ids"] = &autoTunnelMesh.GroupIds
    return children
}

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-all"] = autoTunnelMesh.GroupAll
    return leafs
}

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetBundleName() string { return "cisco_ios_xr" }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetYangName() string { return "auto-tunnel-mesh" }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) SetParent(parent types.Entity) { autoTunnelMesh.parent = parent }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetParent() types.Entity { return autoTunnelMesh.parent }

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetParentYangName() string { return "traffic-engineering" }

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds
// Enable interfaces in specific MPLS TE
// auto-tunnel mesh-groups
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Auto-mesh group identifier to enable. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId.
    GroupId []MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId
}

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetFilter() yfilter.YFilter { return groupIds.YFilter }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) SetFilter(yf yfilter.YFilter) { groupIds.YFilter = yf }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetSegmentPath() string {
    return "group-ids"
}

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-id" {
        for _, c := range groupIds.GroupId {
            if groupIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId{}
        groupIds.GroupId = append(groupIds.GroupId, child)
        return &groupIds.GroupId[len(groupIds.GroupId)-1]
    }
    return nil
}

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupIds.GroupId {
        children[groupIds.GroupId[i].GetSegmentPath()] = &groupIds.GroupId[i]
    }
    return children
}

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetBundleName() string { return "cisco_ios_xr" }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetYangName() string { return "group-ids" }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) SetParent(parent types.Entity) { groupIds.parent = parent }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetParent() types.Entity { return groupIds.parent }

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetParentYangName() string { return "auto-tunnel-mesh" }

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId
// Auto-mesh group identifier to enable
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Mesh group ID. The type is interface{} with range:
    // 0..4294967295.
    MeshGroupId interface{}
}

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetFilter() yfilter.YFilter { return groupId.YFilter }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) SetFilter(yf yfilter.YFilter) { groupId.YFilter = yf }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetGoName(yname string) string {
    if yname == "mesh-group-id" { return "MeshGroupId" }
    return ""
}

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetSegmentPath() string {
    return "group-id" + "[mesh-group-id='" + fmt.Sprintf("%v", groupId.MeshGroupId) + "']"
}

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mesh-group-id"] = groupId.MeshGroupId
    return leafs
}

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetBundleName() string { return "cisco_ios_xr" }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetYangName() string { return "group-id" }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) SetParent(parent types.Entity) { groupId.parent = parent }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetParent() types.Entity { return groupId.parent }

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetParentYangName() string { return "group-ids" }

// MplsLdp_DefaultVrf_Afs_Af_Neighbor
// Configuration related to Neighbors
type MplsLdp_DefaultVrf_Afs_Af_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration related to neighbors using neighbor address.
    Addresses MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses
}

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetGoName(yname string) string {
    if yname == "addresses" { return "Addresses" }
    return ""
}

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &neighbor.Addresses
    }
    return nil
}

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &neighbor.Addresses
    return children
}

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses
// Configuration related to neighbors using
// neighbor address
type MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address based configuration related to a neighbor. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address.
    Address []MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address
}

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetYangName() string { return "addresses" }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetParentYangName() string { return "neighbor" }

// MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address
// IP address based configuration related to a
// neighbor
type MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Establish targeted session with given address. The type is interface{}.
    Targeted interface{}
}

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "targeted" { return "Targeted" }
    return ""
}

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetSegmentPath() string {
    return "address" + "[ip-address='" + fmt.Sprintf("%v", address.IpAddress) + "']"
}

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = address.IpAddress
    leafs["targeted"] = address.Targeted
    return leafs
}

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetYangName() string { return "address" }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetParentYangName() string { return "addresses" }

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol
// MPLS LDP configuration for protocol
// redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP configuration for protocol redistribution.
    Bgp MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp
}

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetFilter() yfilter.YFilter { return redistributionProtocol.YFilter }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) SetFilter(yf yfilter.YFilter) { redistributionProtocol.YFilter = yf }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetGoName(yname string) string {
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetSegmentPath() string {
    return "redistribution-protocol"
}

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp" {
        return &redistributionProtocol.Bgp
    }
    return nil
}

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp"] = &redistributionProtocol.Bgp
    return children
}

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetBundleName() string { return "cisco_ios_xr" }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetYangName() string { return "redistribution-protocol" }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) SetParent(parent types.Entity) { redistributionProtocol.parent = parent }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetParent() types.Entity { return redistributionProtocol.parent }

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp
// MPLS LDP configuration for protocol
// redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP configuration for protocol redistribution.
    As MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As

    // ACL containing list of neighbors for BGP route redistribution.
    AdvertiseTo MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo
}

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "advertise-to" { return "AdvertiseTo" }
    return ""
}

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "as" {
        return &bgp.As
    }
    if childYangName == "advertise-to" {
        return &bgp.AdvertiseTo
    }
    return nil
}

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["as"] = &bgp.As
    children["advertise-to"] = &bgp.AdvertiseTo
    return children
}

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetYangName() string { return "bgp" }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetParentYangName() string { return "redistribution-protocol" }

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As
// MPLS LDP configuration for protocol
// redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First half of BGP AS number in XX.YY format.  Mandatory Must be a non-zero
    // value if second half is zero. The type is interface{} with range: 0..65535.
    AsXx interface{}

    // Second half of BGP AS number in XX.YY format. Mandatory Must be a non-zero
    // value if first half is zero. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}
}

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetFilter() yfilter.YFilter { return as.YFilter }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) SetFilter(yf yfilter.YFilter) { as.YFilter = yf }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    return ""
}

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetSegmentPath() string {
    return "as"
}

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = as.AsXx
    leafs["as-yy"] = as.AsYy
    return leafs
}

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetBundleName() string { return "cisco_ios_xr" }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetYangName() string { return "as" }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) SetParent(parent types.Entity) { as.parent = parent }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetParent() types.Entity { return as.parent }

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetParentYangName() string { return "bgp" }

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo
// ACL containing list of neighbors for BGP
// route redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // advertise to peer acl type. The type is MplsLdpAdvertiseBgpAcl.
    Type interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetFilter() yfilter.YFilter { return advertiseTo.YFilter }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) SetFilter(yf yfilter.YFilter) { advertiseTo.YFilter = yf }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    return ""
}

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetSegmentPath() string {
    return "advertise-to"
}

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = advertiseTo.Type
    leafs["peer-acl-name"] = advertiseTo.PeerAclName
    return leafs
}

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetBundleName() string { return "cisco_ios_xr" }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetYangName() string { return "advertise-to" }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) SetParent(parent types.Entity) { advertiseTo.parent = parent }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetParent() types.Entity { return advertiseTo.parent }

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetParentYangName() string { return "bgp" }

// MplsLdp_DefaultVrf_Global
// Default VRF Global configuration for MPLS LDP
type MplsLdp_DefaultVrf_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for LDP Router ID (LDP ID). The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // LDP Session parameters.
    Session MplsLdp_DefaultVrf_Global_Session

    // Configuration related to Neighbors.
    Neighbor MplsLdp_DefaultVrf_Global_Neighbor

    // Configuration for per-VRF LDP Graceful Restart parameters.
    GracefulRestart MplsLdp_DefaultVrf_Global_GracefulRestart
}

func (global *MplsLdp_DefaultVrf_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *MplsLdp_DefaultVrf_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *MplsLdp_DefaultVrf_Global) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "session" { return "Session" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    return ""
}

func (global *MplsLdp_DefaultVrf_Global) GetSegmentPath() string {
    return "global"
}

func (global *MplsLdp_DefaultVrf_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &global.Session
    }
    if childYangName == "neighbor" {
        return &global.Neighbor
    }
    if childYangName == "graceful-restart" {
        return &global.GracefulRestart
    }
    return nil
}

func (global *MplsLdp_DefaultVrf_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &global.Session
    children["neighbor"] = &global.Neighbor
    children["graceful-restart"] = &global.GracefulRestart
    return children
}

func (global *MplsLdp_DefaultVrf_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = global.RouterId
    return leafs
}

func (global *MplsLdp_DefaultVrf_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *MplsLdp_DefaultVrf_Global) GetYangName() string { return "global" }

func (global *MplsLdp_DefaultVrf_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *MplsLdp_DefaultVrf_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *MplsLdp_DefaultVrf_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *MplsLdp_DefaultVrf_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *MplsLdp_DefaultVrf_Global) GetParent() types.Entity { return global.parent }

func (global *MplsLdp_DefaultVrf_Global) GetParentYangName() string { return "default-vrf" }

// MplsLdp_DefaultVrf_Global_Session
// LDP Session parameters
type MplsLdp_DefaultVrf_Global_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Session Protection parameters.
    Protection MplsLdp_DefaultVrf_Global_Session_Protection

    // ACL with the list of neighbors configured for Downstream on Demand.
    DownstreamOnDemand MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand
}

func (session *MplsLdp_DefaultVrf_Global_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *MplsLdp_DefaultVrf_Global_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *MplsLdp_DefaultVrf_Global_Session) GetGoName(yname string) string {
    if yname == "protection" { return "Protection" }
    if yname == "downstream-on-demand" { return "DownstreamOnDemand" }
    return ""
}

func (session *MplsLdp_DefaultVrf_Global_Session) GetSegmentPath() string {
    return "session"
}

func (session *MplsLdp_DefaultVrf_Global_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protection" {
        return &session.Protection
    }
    if childYangName == "downstream-on-demand" {
        return &session.DownstreamOnDemand
    }
    return nil
}

func (session *MplsLdp_DefaultVrf_Global_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protection"] = &session.Protection
    children["downstream-on-demand"] = &session.DownstreamOnDemand
    return children
}

func (session *MplsLdp_DefaultVrf_Global_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *MplsLdp_DefaultVrf_Global_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *MplsLdp_DefaultVrf_Global_Session) GetYangName() string { return "session" }

func (session *MplsLdp_DefaultVrf_Global_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *MplsLdp_DefaultVrf_Global_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *MplsLdp_DefaultVrf_Global_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *MplsLdp_DefaultVrf_Global_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *MplsLdp_DefaultVrf_Global_Session) GetParent() types.Entity { return session.parent }

func (session *MplsLdp_DefaultVrf_Global_Session) GetParentYangName() string { return "global" }

// MplsLdp_DefaultVrf_Global_Session_Protection
// Configure Session Protection parameters
type MplsLdp_DefaultVrf_Global_Session_Protection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session protection type. The type is MplsLdpSessionProtection.
    ProtectionType interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}

    // Holdup duration. The type is interface{} with range: 30..2147483.
    Duration interface{}
}

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetFilter() yfilter.YFilter { return protection.YFilter }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) SetFilter(yf yfilter.YFilter) { protection.YFilter = yf }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetGoName(yname string) string {
    if yname == "protection-type" { return "ProtectionType" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    if yname == "duration" { return "Duration" }
    return ""
}

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetSegmentPath() string {
    return "protection"
}

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protection-type"] = protection.ProtectionType
    leafs["peer-acl-name"] = protection.PeerAclName
    leafs["duration"] = protection.Duration
    return leafs
}

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetBundleName() string { return "cisco_ios_xr" }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetYangName() string { return "protection" }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) SetParent(parent types.Entity) { protection.parent = parent }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetParent() types.Entity { return protection.parent }

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetParentYangName() string { return "session" }

// MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand
// ACL with the list of neighbors configured for
// Downstream on Demand
type MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Downstream on demand type. The type is MplsLdpDownstreamOnDemand.
    Type interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetFilter() yfilter.YFilter { return downstreamOnDemand.YFilter }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) SetFilter(yf yfilter.YFilter) { downstreamOnDemand.YFilter = yf }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    return ""
}

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetSegmentPath() string {
    return "downstream-on-demand"
}

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = downstreamOnDemand.Type
    leafs["peer-acl-name"] = downstreamOnDemand.PeerAclName
    return leafs
}

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetBundleName() string { return "cisco_ios_xr" }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetYangName() string { return "downstream-on-demand" }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) SetParent(parent types.Entity) { downstreamOnDemand.parent = parent }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetParent() types.Entity { return downstreamOnDemand.parent }

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetParentYangName() string { return "session" }

// MplsLdp_DefaultVrf_Global_Neighbor
// Configuration related to Neighbors
type MplsLdp_DefaultVrf_Global_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default password for all neigbors. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}

    // Configuration related to Neighbors using LDP Id.
    LdpIds MplsLdp_DefaultVrf_Global_Neighbor_LdpIds

    // Configuration related to neighbor transport.
    DualStack MplsLdp_DefaultVrf_Global_Neighbor_DualStack
}

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetGoName(yname string) string {
    if yname == "password" { return "Password" }
    if yname == "ldp-ids" { return "LdpIds" }
    if yname == "dual-stack" { return "DualStack" }
    return ""
}

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-ids" {
        return &neighbor.LdpIds
    }
    if childYangName == "dual-stack" {
        return &neighbor.DualStack
    }
    return nil
}

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ldp-ids"] = &neighbor.LdpIds
    children["dual-stack"] = &neighbor.DualStack
    return children
}

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["password"] = neighbor.Password
    return leafs
}

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetParentYangName() string { return "global" }

// MplsLdp_DefaultVrf_Global_Neighbor_LdpIds
// Configuration related to Neighbors using LDP
// Id
type MplsLdp_DefaultVrf_Global_Neighbor_LdpIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP ID based configuration related to a neigbor. The type is slice of
    // MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId.
    LdpId []MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId
}

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetFilter() yfilter.YFilter { return ldpIds.YFilter }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) SetFilter(yf yfilter.YFilter) { ldpIds.YFilter = yf }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetGoName(yname string) string {
    if yname == "ldp-id" { return "LdpId" }
    return ""
}

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetSegmentPath() string {
    return "ldp-ids"
}

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-id" {
        for _, c := range ldpIds.LdpId {
            if ldpIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId{}
        ldpIds.LdpId = append(ldpIds.LdpId, child)
        return &ldpIds.LdpId[len(ldpIds.LdpId)-1]
    }
    return nil
}

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ldpIds.LdpId {
        children[ldpIds.LdpId[i].GetSegmentPath()] = &ldpIds.LdpId[i]
    }
    return children
}

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetBundleName() string { return "cisco_ios_xr" }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetYangName() string { return "ldp-ids" }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) SetParent(parent types.Entity) { ldpIds.parent = parent }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetParent() types.Entity { return ldpIds.parent }

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetParentYangName() string { return "neighbor" }

// MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId
// LDP ID based configuration related to a
// neigbor
type MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..4294967295.
    LabelSpaceId interface{}

    // Password for MD5 authentication for this neighbor.
    Password MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password
}

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetFilter() yfilter.YFilter { return ldpId.YFilter }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) SetFilter(yf yfilter.YFilter) { ldpId.YFilter = yf }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    if yname == "password" { return "Password" }
    return ""
}

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetSegmentPath() string {
    return "ldp-id" + "[lsr-id='" + fmt.Sprintf("%v", ldpId.LsrId) + "']" + "[label-space-id='" + fmt.Sprintf("%v", ldpId.LabelSpaceId) + "']"
}

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "password" {
        return &ldpId.Password
    }
    return nil
}

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["password"] = &ldpId.Password
    return children
}

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = ldpId.LsrId
    leafs["label-space-id"] = ldpId.LabelSpaceId
    return leafs
}

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetBundleName() string { return "cisco_ios_xr" }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetYangName() string { return "ldp-id" }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) SetParent(parent types.Entity) { ldpId.parent = parent }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetParent() types.Entity { return ldpId.parent }

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetParentYangName() string { return "ldp-ids" }

// MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password
// Password for MD5 authentication for this
// neighbor
type MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Command type for password configuration. The type is MplsLdpNbrPassword.
    CommandType interface{}

    // The neighbor password. The type is string with pattern: (!.+)|([^!].+).
    Password interface{}
}

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetFilter() yfilter.YFilter { return password.YFilter }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) SetFilter(yf yfilter.YFilter) { password.YFilter = yf }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetGoName(yname string) string {
    if yname == "command-type" { return "CommandType" }
    if yname == "password" { return "Password" }
    return ""
}

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetSegmentPath() string {
    return "password"
}

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["command-type"] = password.CommandType
    leafs["password"] = password.Password
    return leafs
}

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetBundleName() string { return "cisco_ios_xr" }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetYangName() string { return "password" }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) SetParent(parent types.Entity) { password.parent = parent }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetParent() types.Entity { return password.parent }

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetParentYangName() string { return "ldp-id" }

// MplsLdp_DefaultVrf_Global_Neighbor_DualStack
// Configuration related to neighbor transport
type MplsLdp_DefaultVrf_Global_Neighbor_DualStack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration to enable neighbor dual-stack tlv-compliance. The type is
    // interface{}.
    TlvCompliance interface{}

    // Configuration related to neighbor transport.
    TransportConnection MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection
}

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetFilter() yfilter.YFilter { return dualStack.YFilter }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) SetFilter(yf yfilter.YFilter) { dualStack.YFilter = yf }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetGoName(yname string) string {
    if yname == "tlv-compliance" { return "TlvCompliance" }
    if yname == "transport-connection" { return "TransportConnection" }
    return ""
}

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetSegmentPath() string {
    return "dual-stack"
}

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport-connection" {
        return &dualStack.TransportConnection
    }
    return nil
}

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transport-connection"] = &dualStack.TransportConnection
    return children
}

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-compliance"] = dualStack.TlvCompliance
    return leafs
}

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetBundleName() string { return "cisco_ios_xr" }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetYangName() string { return "dual-stack" }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) SetParent(parent types.Entity) { dualStack.parent = parent }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetParent() types.Entity { return dualStack.parent }

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetParentYangName() string { return "neighbor" }

// MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection
// Configuration related to neighbor transport
type MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration related to neighbor dual-stack xport-connection max-wait. The
    // type is interface{} with range: 0..60. Units are second.
    MaxWait interface{}

    // Configuration related to neighbor dual-stack xport-connection preference.
    Prefer MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer
}

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetFilter() yfilter.YFilter { return transportConnection.YFilter }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) SetFilter(yf yfilter.YFilter) { transportConnection.YFilter = yf }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetGoName(yname string) string {
    if yname == "max-wait" { return "MaxWait" }
    if yname == "prefer" { return "Prefer" }
    return ""
}

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetSegmentPath() string {
    return "transport-connection"
}

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefer" {
        return &transportConnection.Prefer
    }
    return nil
}

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefer"] = &transportConnection.Prefer
    return children
}

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-wait"] = transportConnection.MaxWait
    return leafs
}

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetBundleName() string { return "cisco_ios_xr" }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetYangName() string { return "transport-connection" }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) SetParent(parent types.Entity) { transportConnection.parent = parent }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetParent() types.Entity { return transportConnection.parent }

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetParentYangName() string { return "dual-stack" }

// MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer
// Configuration related to neighbor
// dual-stack xport-connection preference
type MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration related to neighbor dual-stack xport-connection preference
    // ipv4. The type is interface{}.
    Ipv4 interface{}
}

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetFilter() yfilter.YFilter { return prefer.YFilter }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) SetFilter(yf yfilter.YFilter) { prefer.YFilter = yf }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetSegmentPath() string {
    return "prefer"
}

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4"] = prefer.Ipv4
    return leafs
}

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetBundleName() string { return "cisco_ios_xr" }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetYangName() string { return "prefer" }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) SetParent(parent types.Entity) { prefer.parent = parent }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetParent() types.Entity { return prefer.parent }

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetParentYangName() string { return "transport-connection" }

// MplsLdp_DefaultVrf_Global_GracefulRestart
// Configuration for per-VRF LDP Graceful Restart
// parameters
type MplsLdp_DefaultVrf_Global_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure parameters related to GR peer(s) opearating in helper mode.
    HelperPeer MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer
}

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetGoName(yname string) string {
    if yname == "helper-peer" { return "HelperPeer" }
    return ""
}

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-peer" {
        return &gracefulRestart.HelperPeer
    }
    return nil
}

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-peer"] = &gracefulRestart.HelperPeer
    return children
}

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetBundleName() string { return "cisco_ios_xr" }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetParentYangName() string { return "global" }

// MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer
// Configure parameters related to GR peer(s)
// opearating in helper mode
type MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maintain the state of a GR peer upon a local reset. The type is string.
    MaintainOnLocalReset interface{}
}

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetFilter() yfilter.YFilter { return helperPeer.YFilter }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) SetFilter(yf yfilter.YFilter) { helperPeer.YFilter = yf }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetGoName(yname string) string {
    if yname == "maintain-on-local-reset" { return "MaintainOnLocalReset" }
    return ""
}

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetSegmentPath() string {
    return "helper-peer"
}

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maintain-on-local-reset"] = helperPeer.MaintainOnLocalReset
    return leafs
}

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetBundleName() string { return "cisco_ios_xr" }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetYangName() string { return "helper-peer" }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) SetParent(parent types.Entity) { helperPeer.parent = parent }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetParent() types.Entity { return helperPeer.parent }

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetParentYangName() string { return "graceful-restart" }

// MplsLdp_DefaultVrf_Interfaces
// MPLS LDP configuration pertaining to interfaces
type MplsLdp_DefaultVrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP configuration for a particular interface. The type is slice of
    // MplsLdp_DefaultVrf_Interfaces_Interface.
    Interface []MplsLdp_DefaultVrf_Interfaces_Interface
}

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetParentYangName() string { return "default-vrf" }

// MplsLdp_DefaultVrf_Interfaces_Interface
// MPLS LDP configuration for a particular
// interface
type MplsLdp_DefaultVrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enable Label Distribution Protocol (LDP) on thisinterface. The type is
    // interface{}.
    Enable interface{}

    // Address Family specific configuration for MPLS LDP intf.
    Afs MplsLdp_DefaultVrf_Interfaces_Interface_Afs

    // Per VRF interface Global configuration for MPLS LDP.
    Global MplsLdp_DefaultVrf_Interfaces_Interface_Global
}

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
    if yname == "afs" { return "Afs" }
    if yname == "global" { return "Global" }
    return ""
}

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &self.Afs
    }
    if childYangName == "global" {
        return &self.Global
    }
    return nil
}

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &self.Afs
    children["global"] = &self.Global
    return children
}

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["enable"] = self.Enable
    return leafs
}

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs
// Address Family specific configuration for
// MPLS LDP intf
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af.
    Af []MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af
}

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetParentYangName() string { return "interface" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af
// Configure data for given Address Family
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is MplsLdpafName.
    AfName interface{}

    // Enable Address Family. The type is interface{}.
    Enable interface{}

    // Configure interface discovery parameters.
    Discovery MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery

    // LDP interface IGP configuration.
    Igp MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp

    // Interface configuration parameters for mLDP.
    Mldp MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp
}

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "enable" { return "Enable" }
    if yname == "discovery" { return "Discovery" }
    if yname == "igp" { return "Igp" }
    if yname == "mldp" { return "Mldp" }
    return ""
}

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovery" {
        return &af.Discovery
    }
    if childYangName == "igp" {
        return &af.Igp
    }
    if childYangName == "mldp" {
        return &af.Mldp
    }
    return nil
}

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discovery"] = &af.Discovery
    children["igp"] = &af.Igp
    children["mldp"] = &af.Mldp
    return children
}

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["enable"] = af.Enable
    return leafs
}

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery
// Configure interface discovery parameters
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP configuration for interface discovery transportaddress.
    TransportAddress MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetGoName(yname string) string {
    if yname == "transport-address" { return "TransportAddress" }
    return ""
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport-address" {
        return &discovery.TransportAddress
    }
    return nil
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transport-address"] = &discovery.TransportAddress
    return children
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetBundleName() string { return "cisco_ios_xr" }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
// MPLS LDP configuration for interface
// discovery transportaddress.
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transport address option. The type is MplsLdpTransportAddress.
    AddressType interface{}

    // IP address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetFilter() yfilter.YFilter { return transportAddress.YFilter }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) SetFilter(yf yfilter.YFilter) { transportAddress.YFilter = yf }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "address" { return "Address" }
    return ""
}

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetSegmentPath() string {
    return "transport-address"
}

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = transportAddress.AddressType
    leafs["address"] = transportAddress.Address
    return leafs
}

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetBundleName() string { return "cisco_ios_xr" }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetYangName() string { return "transport-address" }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) SetParent(parent types.Entity) { transportAddress.parent = parent }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetParent() types.Entity { return transportAddress.parent }

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetParentYangName() string { return "discovery" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp
// LDP interface IGP configuration
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable IGP Auto-config on this interface. The type is interface{}.
    DisableAutoConfig interface{}
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetGoName(yname string) string {
    if yname == "disable-auto-config" { return "DisableAutoConfig" }
    return ""
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable-auto-config"] = igp.DisableAutoConfig
    return leafs
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetYangName() string { return "igp" }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetParent() types.Entity { return igp.parent }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp
// Interface configuration parameters for mLDP
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable mLDP on LDP enabled interface. The type is interface{}.
    Disable interface{}
}

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetFilter() yfilter.YFilter { return mldp.YFilter }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) SetFilter(yf yfilter.YFilter) { mldp.YFilter = yf }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetSegmentPath() string {
    return "mldp"
}

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = mldp.Disable
    return leafs
}

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetBundleName() string { return "cisco_ios_xr" }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetYangName() string { return "mldp" }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) SetParent(parent types.Entity) { mldp.parent = parent }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetParent() types.Entity { return mldp.parent }

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetParentYangName() string { return "af" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Global
// Per VRF interface Global configuration for
// MPLS LDP
type MplsLdp_DefaultVrf_Interfaces_Interface_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure interface discovery parameters.
    Discovery MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery

    // LDP IGP configuration.
    Igp MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp
}

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetGoName(yname string) string {
    if yname == "discovery" { return "Discovery" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetSegmentPath() string {
    return "global"
}

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovery" {
        return &global.Discovery
    }
    if childYangName == "igp" {
        return &global.Igp
    }
    return nil
}

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discovery"] = &global.Discovery
    children["igp"] = &global.Igp
    return children
}

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetYangName() string { return "global" }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetParent() types.Entity { return global.parent }

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetParentYangName() string { return "interface" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery
// Configure interface discovery parameters
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable discovery's quick start mode. The type is interface{}.
    DisableQuickStart interface{}

    // LDP Link Hellos.
    LinkHello MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetGoName(yname string) string {
    if yname == "disable-quick-start" { return "DisableQuickStart" }
    if yname == "link-hello" { return "LinkHello" }
    return ""
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-hello" {
        return &discovery.LinkHello
    }
    return nil
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-hello"] = &discovery.LinkHello
    return children
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable-quick-start"] = discovery.DisableQuickStart
    return leafs
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetBundleName() string { return "cisco_ios_xr" }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetParentYangName() string { return "global" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello
// LDP Link Hellos
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Hello interval. The type is interface{} with range: 1..65535. Units
    // are second. The default value is 5.
    Interval interface{}

    // Dual Stack Address Family Preference. The type is MplsLdpafName. The
    // default value is ipv4.
    DualStack interface{}

    // Time (seconds) - 65535 implies infinite. The type is interface{} with
    // range: 1..65535. Units are second. The default value is 15.
    HoldTime interface{}
}

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetFilter() yfilter.YFilter { return linkHello.YFilter }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) SetFilter(yf yfilter.YFilter) { linkHello.YFilter = yf }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "dual-stack" { return "DualStack" }
    if yname == "hold-time" { return "HoldTime" }
    return ""
}

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetSegmentPath() string {
    return "link-hello"
}

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = linkHello.Interval
    leafs["dual-stack"] = linkHello.DualStack
    leafs["hold-time"] = linkHello.HoldTime
    return leafs
}

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetBundleName() string { return "cisco_ios_xr" }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetYangName() string { return "link-hello" }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) SetParent(parent types.Entity) { linkHello.parent = parent }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetParent() types.Entity { return linkHello.parent }

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetParentYangName() string { return "discovery" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp
// LDP IGP configuration
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP IGP synchronization.
    Sync MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetGoName(yname string) string {
    if yname == "sync" { return "Sync" }
    return ""
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sync" {
        return &igp.Sync
    }
    return nil
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sync"] = &igp.Sync
    return children
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetYangName() string { return "igp" }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetParent() types.Entity { return igp.parent }

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetParentYangName() string { return "global" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync
// LDP IGP synchronization
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP IGP synchronization delay time.
    Delay MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay
}

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetFilter() yfilter.YFilter { return sync.YFilter }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) SetFilter(yf yfilter.YFilter) { sync.YFilter = yf }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetGoName(yname string) string {
    if yname == "delay" { return "Delay" }
    return ""
}

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetSegmentPath() string {
    return "sync"
}

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delay" {
        return &sync.Delay
    }
    return nil
}

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delay"] = &sync.Delay
    return children
}

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetBundleName() string { return "cisco_ios_xr" }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetYangName() string { return "sync" }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) SetParent(parent types.Entity) { sync.parent = parent }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetParent() types.Entity { return sync.parent }

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetParentYangName() string { return "igp" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay
// LDP IGP synchronization delay time
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface sync up delay after session up.
    OnSessionUp MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp
}

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetFilter() yfilter.YFilter { return delay.YFilter }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) SetFilter(yf yfilter.YFilter) { delay.YFilter = yf }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetGoName(yname string) string {
    if yname == "on-session-up" { return "OnSessionUp" }
    return ""
}

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetSegmentPath() string {
    return "delay"
}

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "on-session-up" {
        return &delay.OnSessionUp
    }
    return nil
}

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["on-session-up"] = &delay.OnSessionUp
    return children
}

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetBundleName() string { return "cisco_ios_xr" }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetYangName() string { return "delay" }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) SetParent(parent types.Entity) { delay.parent = parent }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetParent() types.Entity { return delay.parent }

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetParentYangName() string { return "sync" }

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp
// Interface sync up delay after session up
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable delay after session up. The type is interface{}.
    Disable interface{}

    // Time (seconds). The type is interface{} with range: 5..300. Units are
    // second.
    Timeout interface{}
}

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetFilter() yfilter.YFilter { return onSessionUp.YFilter }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) SetFilter(yf yfilter.YFilter) { onSessionUp.YFilter = yf }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetSegmentPath() string {
    return "on-session-up"
}

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = onSessionUp.Disable
    leafs["timeout"] = onSessionUp.Timeout
    return leafs
}

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetBundleName() string { return "cisco_ios_xr" }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetYangName() string { return "on-session-up" }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) SetParent(parent types.Entity) { onSessionUp.parent = parent }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetParent() types.Entity { return onSessionUp.parent }

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetParentYangName() string { return "delay" }

// MplsLdp_Vrfs
// VRF Table attribute configuration for MPLS LDP
type MplsLdp_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF attribute configuration for MPLS LDP. The type is slice of
    // MplsLdp_Vrfs_Vrf.
    Vrf []MplsLdp_Vrfs_Vrf
}

func (vrfs *MplsLdp_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *MplsLdp_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *MplsLdp_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *MplsLdp_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *MplsLdp_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *MplsLdp_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *MplsLdp_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *MplsLdp_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *MplsLdp_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *MplsLdp_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *MplsLdp_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *MplsLdp_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *MplsLdp_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *MplsLdp_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *MplsLdp_Vrfs) GetParentYangName() string { return "mpls-ldp" }

// MplsLdp_Vrfs_Vrf
// VRF attribute configuration for MPLS LDP
type MplsLdp_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Enable VRF. The type is interface{}.
    Enable interface{}

    // Per VRF Global configuration for MPLS LDP.
    Global MplsLdp_Vrfs_Vrf_Global

    // Address Family specific configuration for MPLS LDP vrf.
    Afs MplsLdp_Vrfs_Vrf_Afs

    // MPLS LDP configuration pertaining to interfaces.
    Interfaces MplsLdp_Vrfs_Vrf_Interfaces
}

func (vrf *MplsLdp_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *MplsLdp_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *MplsLdp_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "enable" { return "Enable" }
    if yname == "global" { return "Global" }
    if yname == "afs" { return "Afs" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (vrf *MplsLdp_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *MplsLdp_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global" {
        return &vrf.Global
    }
    if childYangName == "afs" {
        return &vrf.Afs
    }
    if childYangName == "interfaces" {
        return &vrf.Interfaces
    }
    return nil
}

func (vrf *MplsLdp_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global"] = &vrf.Global
    children["afs"] = &vrf.Afs
    children["interfaces"] = &vrf.Interfaces
    return children
}

func (vrf *MplsLdp_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["enable"] = vrf.Enable
    return leafs
}

func (vrf *MplsLdp_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *MplsLdp_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *MplsLdp_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *MplsLdp_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *MplsLdp_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *MplsLdp_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *MplsLdp_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *MplsLdp_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// MplsLdp_Vrfs_Vrf_Global
// Per VRF Global configuration for MPLS LDP
type MplsLdp_Vrfs_Vrf_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for LDP Router ID (LDP ID). The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // LDP Session parameters.
    Session MplsLdp_Vrfs_Vrf_Global_Session

    // Configuration related to Neighbors.
    Neighbor MplsLdp_Vrfs_Vrf_Global_Neighbor

    // Configuration for per-VRF LDP Graceful Restart parameters.
    GracefulRestart MplsLdp_Vrfs_Vrf_Global_GracefulRestart
}

func (global *MplsLdp_Vrfs_Vrf_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *MplsLdp_Vrfs_Vrf_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *MplsLdp_Vrfs_Vrf_Global) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "session" { return "Session" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    return ""
}

func (global *MplsLdp_Vrfs_Vrf_Global) GetSegmentPath() string {
    return "global"
}

func (global *MplsLdp_Vrfs_Vrf_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &global.Session
    }
    if childYangName == "neighbor" {
        return &global.Neighbor
    }
    if childYangName == "graceful-restart" {
        return &global.GracefulRestart
    }
    return nil
}

func (global *MplsLdp_Vrfs_Vrf_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &global.Session
    children["neighbor"] = &global.Neighbor
    children["graceful-restart"] = &global.GracefulRestart
    return children
}

func (global *MplsLdp_Vrfs_Vrf_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = global.RouterId
    return leafs
}

func (global *MplsLdp_Vrfs_Vrf_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *MplsLdp_Vrfs_Vrf_Global) GetYangName() string { return "global" }

func (global *MplsLdp_Vrfs_Vrf_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *MplsLdp_Vrfs_Vrf_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *MplsLdp_Vrfs_Vrf_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *MplsLdp_Vrfs_Vrf_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *MplsLdp_Vrfs_Vrf_Global) GetParent() types.Entity { return global.parent }

func (global *MplsLdp_Vrfs_Vrf_Global) GetParentYangName() string { return "vrf" }

// MplsLdp_Vrfs_Vrf_Global_Session
// LDP Session parameters
type MplsLdp_Vrfs_Vrf_Global_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL with the list of neighbors configured for Downstream on Demand.
    DownstreamOnDemand MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand
}

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetGoName(yname string) string {
    if yname == "downstream-on-demand" { return "DownstreamOnDemand" }
    return ""
}

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetSegmentPath() string {
    return "session"
}

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "downstream-on-demand" {
        return &session.DownstreamOnDemand
    }
    return nil
}

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["downstream-on-demand"] = &session.DownstreamOnDemand
    return children
}

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetYangName() string { return "session" }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetParent() types.Entity { return session.parent }

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetParentYangName() string { return "global" }

// MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand
// ACL with the list of neighbors configured
// for Downstream on Demand
type MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Downstream on demand type. The type is MplsLdpDownstreamOnDemand.
    Type interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetFilter() yfilter.YFilter { return downstreamOnDemand.YFilter }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) SetFilter(yf yfilter.YFilter) { downstreamOnDemand.YFilter = yf }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    return ""
}

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetSegmentPath() string {
    return "downstream-on-demand"
}

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = downstreamOnDemand.Type
    leafs["peer-acl-name"] = downstreamOnDemand.PeerAclName
    return leafs
}

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetBundleName() string { return "cisco_ios_xr" }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetYangName() string { return "downstream-on-demand" }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) SetParent(parent types.Entity) { downstreamOnDemand.parent = parent }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetParent() types.Entity { return downstreamOnDemand.parent }

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetParentYangName() string { return "session" }

// MplsLdp_Vrfs_Vrf_Global_Neighbor
// Configuration related to Neighbors
type MplsLdp_Vrfs_Vrf_Global_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default password for all neigbors. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}

    // Configuration related to Neighbors using LDP Id.
    LdpIds MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds
}

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetGoName(yname string) string {
    if yname == "password" { return "Password" }
    if yname == "ldp-ids" { return "LdpIds" }
    return ""
}

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-ids" {
        return &neighbor.LdpIds
    }
    return nil
}

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ldp-ids"] = &neighbor.LdpIds
    return children
}

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["password"] = neighbor.Password
    return leafs
}

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetParentYangName() string { return "global" }

// MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds
// Configuration related to Neighbors using LDP
// Id
type MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP ID based configuration related to a neigbor. The type is slice of
    // MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId.
    LdpId []MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId
}

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetFilter() yfilter.YFilter { return ldpIds.YFilter }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) SetFilter(yf yfilter.YFilter) { ldpIds.YFilter = yf }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetGoName(yname string) string {
    if yname == "ldp-id" { return "LdpId" }
    return ""
}

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetSegmentPath() string {
    return "ldp-ids"
}

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-id" {
        for _, c := range ldpIds.LdpId {
            if ldpIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId{}
        ldpIds.LdpId = append(ldpIds.LdpId, child)
        return &ldpIds.LdpId[len(ldpIds.LdpId)-1]
    }
    return nil
}

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ldpIds.LdpId {
        children[ldpIds.LdpId[i].GetSegmentPath()] = &ldpIds.LdpId[i]
    }
    return children
}

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetBundleName() string { return "cisco_ios_xr" }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetYangName() string { return "ldp-ids" }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) SetParent(parent types.Entity) { ldpIds.parent = parent }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetParent() types.Entity { return ldpIds.parent }

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetParentYangName() string { return "neighbor" }

// MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId
// LDP ID based configuration related to a
// neigbor
type MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..4294967295.
    LabelSpaceId interface{}

    // Password for MD5 authentication for this neighbor.
    Password MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password
}

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetFilter() yfilter.YFilter { return ldpId.YFilter }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) SetFilter(yf yfilter.YFilter) { ldpId.YFilter = yf }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    if yname == "password" { return "Password" }
    return ""
}

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetSegmentPath() string {
    return "ldp-id" + "[lsr-id='" + fmt.Sprintf("%v", ldpId.LsrId) + "']" + "[label-space-id='" + fmt.Sprintf("%v", ldpId.LabelSpaceId) + "']"
}

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "password" {
        return &ldpId.Password
    }
    return nil
}

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["password"] = &ldpId.Password
    return children
}

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = ldpId.LsrId
    leafs["label-space-id"] = ldpId.LabelSpaceId
    return leafs
}

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetBundleName() string { return "cisco_ios_xr" }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetYangName() string { return "ldp-id" }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) SetParent(parent types.Entity) { ldpId.parent = parent }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetParent() types.Entity { return ldpId.parent }

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetParentYangName() string { return "ldp-ids" }

// MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password
// Password for MD5 authentication for this
// neighbor
type MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Command type for password configuration. The type is MplsLdpNbrPassword.
    CommandType interface{}

    // The neighbor password. The type is string with pattern: (!.+)|([^!].+).
    Password interface{}
}

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetFilter() yfilter.YFilter { return password.YFilter }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) SetFilter(yf yfilter.YFilter) { password.YFilter = yf }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetGoName(yname string) string {
    if yname == "command-type" { return "CommandType" }
    if yname == "password" { return "Password" }
    return ""
}

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetSegmentPath() string {
    return "password"
}

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["command-type"] = password.CommandType
    leafs["password"] = password.Password
    return leafs
}

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetBundleName() string { return "cisco_ios_xr" }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetYangName() string { return "password" }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) SetParent(parent types.Entity) { password.parent = parent }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetParent() types.Entity { return password.parent }

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetParentYangName() string { return "ldp-id" }

// MplsLdp_Vrfs_Vrf_Global_GracefulRestart
// Configuration for per-VRF LDP Graceful
// Restart parameters
type MplsLdp_Vrfs_Vrf_Global_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure parameters related to GR peer(s) opearating in helper mode.
    HelperPeer MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer
}

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetGoName(yname string) string {
    if yname == "helper-peer" { return "HelperPeer" }
    return ""
}

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-peer" {
        return &gracefulRestart.HelperPeer
    }
    return nil
}

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-peer"] = &gracefulRestart.HelperPeer
    return children
}

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetBundleName() string { return "cisco_ios_xr" }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetParentYangName() string { return "global" }

// MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer
// Configure parameters related to GR peer(s)
// opearating in helper mode
type MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maintain the state of a GR peer upon a local reset. The type is string.
    MaintainOnLocalReset interface{}
}

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetFilter() yfilter.YFilter { return helperPeer.YFilter }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) SetFilter(yf yfilter.YFilter) { helperPeer.YFilter = yf }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetGoName(yname string) string {
    if yname == "maintain-on-local-reset" { return "MaintainOnLocalReset" }
    return ""
}

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetSegmentPath() string {
    return "helper-peer"
}

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maintain-on-local-reset"] = helperPeer.MaintainOnLocalReset
    return leafs
}

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetBundleName() string { return "cisco_ios_xr" }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetYangName() string { return "helper-peer" }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) SetParent(parent types.Entity) { helperPeer.parent = parent }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetParent() types.Entity { return helperPeer.parent }

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetParentYangName() string { return "graceful-restart" }

// MplsLdp_Vrfs_Vrf_Afs
// Address Family specific configuration for MPLS
// LDP vrf
type MplsLdp_Vrfs_Vrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af.
    Af []MplsLdp_Vrfs_Vrf_Afs_Af
}

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_Vrfs_Vrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsLdp_Vrfs_Vrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetParentYangName() string { return "vrf" }

// MplsLdp_Vrfs_Vrf_Afs_Af
// Configure data for given Address Family
type MplsLdp_Vrfs_Vrf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is MplsLdpafName.
    AfName interface{}

    // Enable Address Family. The type is interface{}.
    Enable interface{}

    // Configure Discovery parameters.
    Discovery MplsLdp_Vrfs_Vrf_Afs_Af_Discovery

    // Configure Label policies and control.
    Label MplsLdp_Vrfs_Vrf_Afs_Af_Label
}

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "enable" { return "Enable" }
    if yname == "discovery" { return "Discovery" }
    if yname == "label" { return "Label" }
    return ""
}

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovery" {
        return &af.Discovery
    }
    if childYangName == "label" {
        return &af.Label
    }
    return nil
}

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discovery"] = &af.Discovery
    children["label"] = &af.Label
    return children
}

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["enable"] = af.Enable
    return leafs
}

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Discovery
// Configure Discovery parameters
type MplsLdp_Vrfs_Vrf_Afs_Af_Discovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global discovery transport address for address family. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    TransportAddress interface{}
}

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetGoName(yname string) string {
    if yname == "transport-address" { return "TransportAddress" }
    return ""
}

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport-address"] = discovery.TransportAddress
    return leafs
}

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetBundleName() string { return "cisco_ios_xr" }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetParentYangName() string { return "af" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label
// Configure Label policies and control
type MplsLdp_Vrfs_Vrf_Afs_Af_Label struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure remote/peer label policies and control.
    Remote MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote

    // Configure local label policies and control.
    Local MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local
}

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetFilter() yfilter.YFilter { return label.YFilter }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) SetFilter(yf yfilter.YFilter) { label.YFilter = yf }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetGoName(yname string) string {
    if yname == "remote" { return "Remote" }
    if yname == "local" { return "Local" }
    return ""
}

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetSegmentPath() string {
    return "label"
}

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote" {
        return &label.Remote
    }
    if childYangName == "local" {
        return &label.Local
    }
    return nil
}

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote"] = &label.Remote
    children["local"] = &label.Local
    return children
}

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetBundleName() string { return "cisco_ios_xr" }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetYangName() string { return "label" }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) SetParent(parent types.Entity) { label.parent = parent }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetParent() types.Entity { return label.parent }

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetParentYangName() string { return "af" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote
// Configure remote/peer label policies and
// control
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure inbound label acceptance.
    Accept MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept
}

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetFilter() yfilter.YFilter { return remote.YFilter }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) SetFilter(yf yfilter.YFilter) { remote.YFilter = yf }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetGoName(yname string) string {
    if yname == "accept" { return "Accept" }
    return ""
}

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetSegmentPath() string {
    return "remote"
}

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accept" {
        return &remote.Accept
    }
    return nil
}

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accept"] = &remote.Accept
    return children
}

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetBundleName() string { return "cisco_ios_xr" }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetYangName() string { return "remote" }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) SetParent(parent types.Entity) { remote.parent = parent }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetParent() types.Entity { return remote.parent }

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetParentYangName() string { return "label" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept
// Configure inbound label acceptance
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration related to Neighbors for inbound label acceptance.
    PeerAcceptPolicies MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
}

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetFilter() yfilter.YFilter { return accept.YFilter }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) SetFilter(yf yfilter.YFilter) { accept.YFilter = yf }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetGoName(yname string) string {
    if yname == "peer-accept-policies" { return "PeerAcceptPolicies" }
    return ""
}

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetSegmentPath() string {
    return "accept"
}

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-accept-policies" {
        return &accept.PeerAcceptPolicies
    }
    return nil
}

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-accept-policies"] = &accept.PeerAcceptPolicies
    return children
}

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetBundleName() string { return "cisco_ios_xr" }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetYangName() string { return "accept" }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) SetParent(parent types.Entity) { accept.parent = parent }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetParent() types.Entity { return accept.parent }

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetParentYangName() string { return "remote" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
// Configuration related to Neighbors for
// inbound label acceptance
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control acceptasnce of labels from a neighbor for prefix(es) using ACL. The
    // type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy.
    PeerAcceptPolicy []MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
}

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetFilter() yfilter.YFilter { return peerAcceptPolicies.YFilter }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) SetFilter(yf yfilter.YFilter) { peerAcceptPolicies.YFilter = yf }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetGoName(yname string) string {
    if yname == "peer-accept-policy" { return "PeerAcceptPolicy" }
    return ""
}

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetSegmentPath() string {
    return "peer-accept-policies"
}

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-accept-policy" {
        for _, c := range peerAcceptPolicies.PeerAcceptPolicy {
            if peerAcceptPolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy{}
        peerAcceptPolicies.PeerAcceptPolicy = append(peerAcceptPolicies.PeerAcceptPolicy, child)
        return &peerAcceptPolicies.PeerAcceptPolicy[len(peerAcceptPolicies.PeerAcceptPolicy)-1]
    }
    return nil
}

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerAcceptPolicies.PeerAcceptPolicy {
        children[peerAcceptPolicies.PeerAcceptPolicy[i].GetSegmentPath()] = &peerAcceptPolicies.PeerAcceptPolicy[i]
    }
    return children
}

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetYangName() string { return "peer-accept-policies" }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) SetParent(parent types.Entity) { peerAcceptPolicies.parent = parent }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetParent() types.Entity { return peerAcceptPolicies.parent }

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetParentYangName() string { return "accept" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
// Control acceptasnce of labels from a
// neighbor for prefix(es) using ACL
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..4294967295.
    LabelSpaceId interface{}

    // Data container.
    PeerAcceptPolicyData MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData

    // keys: lsr-id. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId.
    LsrId []MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId
}

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetFilter() yfilter.YFilter { return peerAcceptPolicy.YFilter }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) SetFilter(yf yfilter.YFilter) { peerAcceptPolicy.YFilter = yf }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetGoName(yname string) string {
    if yname == "label-space-id" { return "LabelSpaceId" }
    if yname == "peer-accept-policy-data" { return "PeerAcceptPolicyData" }
    if yname == "lsr-id" { return "LsrId" }
    return ""
}

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetSegmentPath() string {
    return "peer-accept-policy" + "[label-space-id='" + fmt.Sprintf("%v", peerAcceptPolicy.LabelSpaceId) + "']"
}

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-accept-policy-data" {
        return &peerAcceptPolicy.PeerAcceptPolicyData
    }
    if childYangName == "lsr-id" {
        for _, c := range peerAcceptPolicy.LsrId {
            if peerAcceptPolicy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId{}
        peerAcceptPolicy.LsrId = append(peerAcceptPolicy.LsrId, child)
        return &peerAcceptPolicy.LsrId[len(peerAcceptPolicy.LsrId)-1]
    }
    return nil
}

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-accept-policy-data"] = &peerAcceptPolicy.PeerAcceptPolicyData
    for i := range peerAcceptPolicy.LsrId {
        children[peerAcceptPolicy.LsrId[i].GetSegmentPath()] = &peerAcceptPolicy.LsrId[i]
    }
    return children
}

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-space-id"] = peerAcceptPolicy.LabelSpaceId
    return leafs
}

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetYangName() string { return "peer-accept-policy" }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) SetParent(parent types.Entity) { peerAcceptPolicy.parent = parent }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetParent() types.Entity { return peerAcceptPolicy.parent }

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetParentYangName() string { return "peer-accept-policies" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData
// Data container.
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetFilter() yfilter.YFilter { return peerAcceptPolicyData.YFilter }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) SetFilter(yf yfilter.YFilter) { peerAcceptPolicyData.YFilter = yf }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetGoName(yname string) string {
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetSegmentPath() string {
    return "peer-accept-policy-data"
}

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-acl-name"] = peerAcceptPolicyData.PrefixAclName
    return leafs
}

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetBundleName() string { return "cisco_ios_xr" }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetYangName() string { return "peer-accept-policy-data" }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) SetParent(parent types.Entity) { peerAcceptPolicyData.parent = parent }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetParent() types.Entity { return peerAcceptPolicyData.parent }

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetParentYangName() string { return "peer-accept-policy" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId
// keys: lsr-id
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetFilter() yfilter.YFilter { return lsrId.YFilter }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) SetFilter(yf yfilter.YFilter) { lsrId.YFilter = yf }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetSegmentPath() string {
    return "lsr-id" + "[lsr-id='" + fmt.Sprintf("%v", lsrId.LsrId) + "']"
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = lsrId.LsrId
    leafs["prefix-acl-name"] = lsrId.PrefixAclName
    return leafs
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetBundleName() string { return "cisco_ios_xr" }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetYangName() string { return "lsr-id" }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) SetParent(parent types.Entity) { lsrId.parent = parent }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetParent() types.Entity { return lsrId.parent }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetParentYangName() string { return "peer-accept-policy" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local
// Configure local label policies and control
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control use of implicit-null label for set of prefix(es). The type is
    // string.
    ImplicitNullOverride interface{}

    // Enable MPLS forwarding for default route. The type is interface{}.
    DefaultRoute interface{}

    // Configure outbound label advertisement.
    Advertise MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise

    // Control local label allocation for prefix(es).
    Allocate MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate
}

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetFilter() yfilter.YFilter { return local.YFilter }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) SetFilter(yf yfilter.YFilter) { local.YFilter = yf }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetGoName(yname string) string {
    if yname == "implicit-null-override" { return "ImplicitNullOverride" }
    if yname == "default-route" { return "DefaultRoute" }
    if yname == "advertise" { return "Advertise" }
    if yname == "allocate" { return "Allocate" }
    return ""
}

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetSegmentPath() string {
    return "local"
}

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "advertise" {
        return &local.Advertise
    }
    if childYangName == "allocate" {
        return &local.Allocate
    }
    return nil
}

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["advertise"] = &local.Advertise
    children["allocate"] = &local.Allocate
    return children
}

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["implicit-null-override"] = local.ImplicitNullOverride
    leafs["default-route"] = local.DefaultRoute
    return leafs
}

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetBundleName() string { return "cisco_ios_xr" }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetYangName() string { return "local" }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) SetParent(parent types.Entity) { local.parent = parent }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetParent() types.Entity { return local.parent }

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetParentYangName() string { return "label" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise
// Configure outbound label advertisement
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable label advertisement. The type is interface{}.
    Disable interface{}

    // Configure peer centric outbound label advertisement using ACL.
    PeerAdvertisePolicies MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies

    // Configure outbound label advertisement for an interface.
    Interfaces MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces

    // Configure advertisment of explicit-null for connected prefixes.
    ExplicitNull MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull
}

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetFilter() yfilter.YFilter { return advertise.YFilter }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) SetFilter(yf yfilter.YFilter) { advertise.YFilter = yf }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "peer-advertise-policies" { return "PeerAdvertisePolicies" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "explicit-null" { return "ExplicitNull" }
    return ""
}

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetSegmentPath() string {
    return "advertise"
}

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-advertise-policies" {
        return &advertise.PeerAdvertisePolicies
    }
    if childYangName == "interfaces" {
        return &advertise.Interfaces
    }
    if childYangName == "explicit-null" {
        return &advertise.ExplicitNull
    }
    return nil
}

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-advertise-policies"] = &advertise.PeerAdvertisePolicies
    children["interfaces"] = &advertise.Interfaces
    children["explicit-null"] = &advertise.ExplicitNull
    return children
}

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = advertise.Disable
    return leafs
}

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetBundleName() string { return "cisco_ios_xr" }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetYangName() string { return "advertise" }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) SetParent(parent types.Entity) { advertise.parent = parent }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetParent() types.Entity { return advertise.parent }

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetParentYangName() string { return "local" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies
// Configure peer centric outbound label
// advertisement using ACL
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control advertisement of prefix(es) using ACL. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy.
    PeerAdvertisePolicy []MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
}

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetFilter() yfilter.YFilter { return peerAdvertisePolicies.YFilter }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) SetFilter(yf yfilter.YFilter) { peerAdvertisePolicies.YFilter = yf }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetGoName(yname string) string {
    if yname == "peer-advertise-policy" { return "PeerAdvertisePolicy" }
    return ""
}

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetSegmentPath() string {
    return "peer-advertise-policies"
}

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-advertise-policy" {
        for _, c := range peerAdvertisePolicies.PeerAdvertisePolicy {
            if peerAdvertisePolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy{}
        peerAdvertisePolicies.PeerAdvertisePolicy = append(peerAdvertisePolicies.PeerAdvertisePolicy, child)
        return &peerAdvertisePolicies.PeerAdvertisePolicy[len(peerAdvertisePolicies.PeerAdvertisePolicy)-1]
    }
    return nil
}

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerAdvertisePolicies.PeerAdvertisePolicy {
        children[peerAdvertisePolicies.PeerAdvertisePolicy[i].GetSegmentPath()] = &peerAdvertisePolicies.PeerAdvertisePolicy[i]
    }
    return children
}

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetBundleName() string { return "cisco_ios_xr" }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetYangName() string { return "peer-advertise-policies" }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) SetParent(parent types.Entity) { peerAdvertisePolicies.parent = parent }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetParent() types.Entity { return peerAdvertisePolicies.parent }

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetParentYangName() string { return "advertise" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
// Control advertisement of prefix(es)
// using ACL
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..4294967295.
    LabelSpaceId interface{}

    // Data container.
    PeerAdvertisePolicyData MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData

    // keys: lsr-id. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId.
    LsrId []MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId
}

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetFilter() yfilter.YFilter { return peerAdvertisePolicy.YFilter }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) SetFilter(yf yfilter.YFilter) { peerAdvertisePolicy.YFilter = yf }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetGoName(yname string) string {
    if yname == "label-space-id" { return "LabelSpaceId" }
    if yname == "peer-advertise-policy-data" { return "PeerAdvertisePolicyData" }
    if yname == "lsr-id" { return "LsrId" }
    return ""
}

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetSegmentPath() string {
    return "peer-advertise-policy" + "[label-space-id='" + fmt.Sprintf("%v", peerAdvertisePolicy.LabelSpaceId) + "']"
}

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-advertise-policy-data" {
        return &peerAdvertisePolicy.PeerAdvertisePolicyData
    }
    if childYangName == "lsr-id" {
        for _, c := range peerAdvertisePolicy.LsrId {
            if peerAdvertisePolicy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId{}
        peerAdvertisePolicy.LsrId = append(peerAdvertisePolicy.LsrId, child)
        return &peerAdvertisePolicy.LsrId[len(peerAdvertisePolicy.LsrId)-1]
    }
    return nil
}

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-advertise-policy-data"] = &peerAdvertisePolicy.PeerAdvertisePolicyData
    for i := range peerAdvertisePolicy.LsrId {
        children[peerAdvertisePolicy.LsrId[i].GetSegmentPath()] = &peerAdvertisePolicy.LsrId[i]
    }
    return children
}

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-space-id"] = peerAdvertisePolicy.LabelSpaceId
    return leafs
}

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetYangName() string { return "peer-advertise-policy" }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) SetParent(parent types.Entity) { peerAdvertisePolicy.parent = parent }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetParent() types.Entity { return peerAdvertisePolicy.parent }

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetParentYangName() string { return "peer-advertise-policies" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData
// Data container.
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetFilter() yfilter.YFilter { return peerAdvertisePolicyData.YFilter }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) SetFilter(yf yfilter.YFilter) { peerAdvertisePolicyData.YFilter = yf }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetGoName(yname string) string {
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetSegmentPath() string {
    return "peer-advertise-policy-data"
}

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-acl-name"] = peerAdvertisePolicyData.PrefixAclName
    return leafs
}

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetBundleName() string { return "cisco_ios_xr" }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetYangName() string { return "peer-advertise-policy-data" }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) SetParent(parent types.Entity) { peerAdvertisePolicyData.parent = parent }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetParent() types.Entity { return peerAdvertisePolicyData.parent }

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetParentYangName() string { return "peer-advertise-policy" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId
// keys: lsr-id
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetFilter() yfilter.YFilter { return lsrId.YFilter }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) SetFilter(yf yfilter.YFilter) { lsrId.YFilter = yf }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetSegmentPath() string {
    return "lsr-id" + "[lsr-id='" + fmt.Sprintf("%v", lsrId.LsrId) + "']"
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = lsrId.LsrId
    leafs["prefix-acl-name"] = lsrId.PrefixAclName
    return leafs
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetBundleName() string { return "cisco_ios_xr" }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetYangName() string { return "lsr-id" }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) SetParent(parent types.Entity) { lsrId.parent = parent }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetParent() types.Entity { return lsrId.parent }

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetParentYangName() string { return "peer-advertise-policy" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces
// Configure outbound label advertisement
// for an interface
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Control advertisement of interface's host IP address. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface.
    Interface []MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
}

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetParentYangName() string { return "advertise" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
// Control advertisement of interface's
// host IP address
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull
// Configure advertisment of explicit-null
// for connected prefixes.
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Explicit Null command variant. The type is MplsLdpExpNull.
    ExplicitNullType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetFilter() yfilter.YFilter { return explicitNull.YFilter }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) SetFilter(yf yfilter.YFilter) { explicitNull.YFilter = yf }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetGoName(yname string) string {
    if yname == "explicit-null-type" { return "ExplicitNullType" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    if yname == "peer-acl-name" { return "PeerAclName" }
    return ""
}

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetSegmentPath() string {
    return "explicit-null"
}

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["explicit-null-type"] = explicitNull.ExplicitNullType
    leafs["prefix-acl-name"] = explicitNull.PrefixAclName
    leafs["peer-acl-name"] = explicitNull.PeerAclName
    return leafs
}

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetBundleName() string { return "cisco_ios_xr" }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetYangName() string { return "explicit-null" }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) SetParent(parent types.Entity) { explicitNull.parent = parent }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetParent() types.Entity { return explicitNull.parent }

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetParentYangName() string { return "advertise" }

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate
// Control local label allocation for
// prefix(es)
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label allocation type. The type is MplsLdpLabelAllocation.
    AllocationType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}
}

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetFilter() yfilter.YFilter { return allocate.YFilter }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) SetFilter(yf yfilter.YFilter) { allocate.YFilter = yf }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetGoName(yname string) string {
    if yname == "allocation-type" { return "AllocationType" }
    if yname == "prefix-acl-name" { return "PrefixAclName" }
    return ""
}

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetSegmentPath() string {
    return "allocate"
}

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allocation-type"] = allocate.AllocationType
    leafs["prefix-acl-name"] = allocate.PrefixAclName
    return leafs
}

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetBundleName() string { return "cisco_ios_xr" }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetYangName() string { return "allocate" }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) SetParent(parent types.Entity) { allocate.parent = parent }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetParent() types.Entity { return allocate.parent }

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetParentYangName() string { return "local" }

// MplsLdp_Vrfs_Vrf_Interfaces
// MPLS LDP configuration pertaining to
// interfaces
type MplsLdp_Vrfs_Vrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP configuration for a particular interface. The type is slice of
    // MplsLdp_Vrfs_Vrf_Interfaces_Interface.
    Interface []MplsLdp_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetParentYangName() string { return "vrf" }

// MplsLdp_Vrfs_Vrf_Interfaces_Interface
// MPLS LDP configuration for a particular
// interface
type MplsLdp_Vrfs_Vrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enable Label Distribution Protocol (LDP) on thisinterface. The type is
    // interface{}.
    Enable interface{}

    // Address Family specific configuration for MPLS LDP vrf intf.
    Afs MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs
}

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
    if yname == "afs" { return "Afs" }
    return ""
}

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &self.Afs
    }
    return nil
}

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &self.Afs
    return children
}

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["enable"] = self.Enable
    return leafs
}

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs
// Address Family specific configuration for
// MPLS LDP vrf intf
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af.
    Af []MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af
}

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetParentYangName() string { return "interface" }

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af
// Configure data for given Address Family
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is MplsLdpafName.
    AfName interface{}

    // Enable Address Family. The type is interface{}.
    Enable interface{}

    // Configure interface discovery parameters.
    Discovery MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery
}

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "enable" { return "Enable" }
    if yname == "discovery" { return "Discovery" }
    return ""
}

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovery" {
        return &af.Discovery
    }
    return nil
}

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discovery"] = &af.Discovery
    return children
}

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["enable"] = af.Enable
    return leafs
}

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery
// Configure interface discovery parameters
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP configuration for interface discovery transportaddress.
    TransportAddress MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
}

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetGoName(yname string) string {
    if yname == "transport-address" { return "TransportAddress" }
    return ""
}

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport-address" {
        return &discovery.TransportAddress
    }
    return nil
}

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transport-address"] = &discovery.TransportAddress
    return children
}

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetBundleName() string { return "cisco_ios_xr" }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetParentYangName() string { return "af" }

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
// MPLS LDP configuration for interface
// discovery transportaddress.
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transport address option. The type is MplsLdpTransportAddress.
    AddressType interface{}

    // IP address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetFilter() yfilter.YFilter { return transportAddress.YFilter }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) SetFilter(yf yfilter.YFilter) { transportAddress.YFilter = yf }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "address" { return "Address" }
    return ""
}

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetSegmentPath() string {
    return "transport-address"
}

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = transportAddress.AddressType
    leafs["address"] = transportAddress.Address
    return leafs
}

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetBundleName() string { return "cisco_ios_xr" }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetYangName() string { return "transport-address" }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) SetParent(parent types.Entity) { transportAddress.parent = parent }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetParent() types.Entity { return transportAddress.parent }

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetParentYangName() string { return "discovery" }

// MplsLdp_Global
// Global configuration for MPLS LDP
type MplsLdp_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable the implicit enabling for IPv4 address family. The type is
    // interface{}.
    DisableImplicitIpv4 interface{}

    // Configure Ltrace Buffer Multiplier. The type is interface{} with range:
    // 1..5. The default value is 1.
    LtraceBufMultiplier interface{}

    // Configure for LDP Entropy-Label.
    EntropyLabel MplsLdp_Global_EntropyLabel

    // LDP Session parameters.
    Session MplsLdp_Global_Session

    // LDP IGP configuration.
    Igp MplsLdp_Global_Igp

    // Enable logging of events.
    EnableLogging MplsLdp_Global_EnableLogging

    // Configure LDP signalling parameters.
    Signalling MplsLdp_Global_Signalling

    // Configure LDP Non-Stop Routing.
    Nsr MplsLdp_Global_Nsr

    // Configuration for LDP Graceful Restart parameters.
    GracefulRestart MplsLdp_Global_GracefulRestart

    // Configure Discovery parameters.
    Discovery MplsLdp_Global_Discovery

    // MPLS mLDP configuration.
    Mldp MplsLdp_Global_Mldp
}

func (global *MplsLdp_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *MplsLdp_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *MplsLdp_Global) GetGoName(yname string) string {
    if yname == "disable-implicit-ipv4" { return "DisableImplicitIpv4" }
    if yname == "ltrace-buf-multiplier" { return "LtraceBufMultiplier" }
    if yname == "entropy-label" { return "EntropyLabel" }
    if yname == "session" { return "Session" }
    if yname == "igp" { return "Igp" }
    if yname == "enable-logging" { return "EnableLogging" }
    if yname == "signalling" { return "Signalling" }
    if yname == "nsr" { return "Nsr" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "discovery" { return "Discovery" }
    if yname == "mldp" { return "Mldp" }
    return ""
}

func (global *MplsLdp_Global) GetSegmentPath() string {
    return "global"
}

func (global *MplsLdp_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entropy-label" {
        return &global.EntropyLabel
    }
    if childYangName == "session" {
        return &global.Session
    }
    if childYangName == "igp" {
        return &global.Igp
    }
    if childYangName == "enable-logging" {
        return &global.EnableLogging
    }
    if childYangName == "signalling" {
        return &global.Signalling
    }
    if childYangName == "nsr" {
        return &global.Nsr
    }
    if childYangName == "graceful-restart" {
        return &global.GracefulRestart
    }
    if childYangName == "discovery" {
        return &global.Discovery
    }
    if childYangName == "mldp" {
        return &global.Mldp
    }
    return nil
}

func (global *MplsLdp_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entropy-label"] = &global.EntropyLabel
    children["session"] = &global.Session
    children["igp"] = &global.Igp
    children["enable-logging"] = &global.EnableLogging
    children["signalling"] = &global.Signalling
    children["nsr"] = &global.Nsr
    children["graceful-restart"] = &global.GracefulRestart
    children["discovery"] = &global.Discovery
    children["mldp"] = &global.Mldp
    return children
}

func (global *MplsLdp_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable-implicit-ipv4"] = global.DisableImplicitIpv4
    leafs["ltrace-buf-multiplier"] = global.LtraceBufMultiplier
    return leafs
}

func (global *MplsLdp_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *MplsLdp_Global) GetYangName() string { return "global" }

func (global *MplsLdp_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *MplsLdp_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *MplsLdp_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *MplsLdp_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *MplsLdp_Global) GetParent() types.Entity { return global.parent }

func (global *MplsLdp_Global) GetParentYangName() string { return "mpls-ldp" }

// MplsLdp_Global_EntropyLabel
// Configure for LDP Entropy-Label
type MplsLdp_Global_EntropyLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is interface{}.
    Enable interface{}
}

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetFilter() yfilter.YFilter { return entropyLabel.YFilter }

func (entropyLabel *MplsLdp_Global_EntropyLabel) SetFilter(yf yfilter.YFilter) { entropyLabel.YFilter = yf }

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetSegmentPath() string {
    return "entropy-label"
}

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = entropyLabel.Enable
    return leafs
}

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetBundleName() string { return "cisco_ios_xr" }

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetYangName() string { return "entropy-label" }

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entropyLabel *MplsLdp_Global_EntropyLabel) SetParent(parent types.Entity) { entropyLabel.parent = parent }

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetParent() types.Entity { return entropyLabel.parent }

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetParentYangName() string { return "global" }

// MplsLdp_Global_Session
// LDP Session parameters
type MplsLdp_Global_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP Session holdtime. The type is interface{} with range: 15..65535. Units
    // are second. The default value is 180.
    HoldTime interface{}

    // Configure Session Backoff parameters.
    BackoffTime MplsLdp_Global_Session_BackoffTime
}

func (session *MplsLdp_Global_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *MplsLdp_Global_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *MplsLdp_Global_Session) GetGoName(yname string) string {
    if yname == "hold-time" { return "HoldTime" }
    if yname == "backoff-time" { return "BackoffTime" }
    return ""
}

func (session *MplsLdp_Global_Session) GetSegmentPath() string {
    return "session"
}

func (session *MplsLdp_Global_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "backoff-time" {
        return &session.BackoffTime
    }
    return nil
}

func (session *MplsLdp_Global_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["backoff-time"] = &session.BackoffTime
    return children
}

func (session *MplsLdp_Global_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hold-time"] = session.HoldTime
    return leafs
}

func (session *MplsLdp_Global_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *MplsLdp_Global_Session) GetYangName() string { return "session" }

func (session *MplsLdp_Global_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *MplsLdp_Global_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *MplsLdp_Global_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *MplsLdp_Global_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *MplsLdp_Global_Session) GetParent() types.Entity { return session.parent }

func (session *MplsLdp_Global_Session) GetParentYangName() string { return "global" }

// MplsLdp_Global_Session_BackoffTime
// Configure Session Backoff parameters
type MplsLdp_Global_Session_BackoffTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Initial session backoff time (seconds). The type is interface{} with range:
    // 5..2147483. Units are second. The default value is 15.
    InitialBackoffTime interface{}

    // Maximum session backoff time (seconds). The type is interface{} with range:
    // 5..2147483. Units are second. The default value is 120.
    MaxBackoffTime interface{}
}

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetFilter() yfilter.YFilter { return backoffTime.YFilter }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) SetFilter(yf yfilter.YFilter) { backoffTime.YFilter = yf }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetGoName(yname string) string {
    if yname == "initial-backoff-time" { return "InitialBackoffTime" }
    if yname == "max-backoff-time" { return "MaxBackoffTime" }
    return ""
}

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetSegmentPath() string {
    return "backoff-time"
}

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initial-backoff-time"] = backoffTime.InitialBackoffTime
    leafs["max-backoff-time"] = backoffTime.MaxBackoffTime
    return leafs
}

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetBundleName() string { return "cisco_ios_xr" }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetYangName() string { return "backoff-time" }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) SetParent(parent types.Entity) { backoffTime.parent = parent }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetParent() types.Entity { return backoffTime.parent }

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetParentYangName() string { return "session" }

// MplsLdp_Global_Igp
// LDP IGP configuration
type MplsLdp_Global_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP IGP synchronization.
    Sync MplsLdp_Global_Igp_Sync
}

func (igp *MplsLdp_Global_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *MplsLdp_Global_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *MplsLdp_Global_Igp) GetGoName(yname string) string {
    if yname == "sync" { return "Sync" }
    return ""
}

func (igp *MplsLdp_Global_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *MplsLdp_Global_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sync" {
        return &igp.Sync
    }
    return nil
}

func (igp *MplsLdp_Global_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sync"] = &igp.Sync
    return children
}

func (igp *MplsLdp_Global_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igp *MplsLdp_Global_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *MplsLdp_Global_Igp) GetYangName() string { return "igp" }

func (igp *MplsLdp_Global_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *MplsLdp_Global_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *MplsLdp_Global_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *MplsLdp_Global_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *MplsLdp_Global_Igp) GetParent() types.Entity { return igp.parent }

func (igp *MplsLdp_Global_Igp) GetParentYangName() string { return "global" }

// MplsLdp_Global_Igp_Sync
// LDP IGP synchronization
type MplsLdp_Global_Igp_Sync struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP IGP synchronization delay time.
    Delay MplsLdp_Global_Igp_Sync_Delay
}

func (sync *MplsLdp_Global_Igp_Sync) GetFilter() yfilter.YFilter { return sync.YFilter }

func (sync *MplsLdp_Global_Igp_Sync) SetFilter(yf yfilter.YFilter) { sync.YFilter = yf }

func (sync *MplsLdp_Global_Igp_Sync) GetGoName(yname string) string {
    if yname == "delay" { return "Delay" }
    return ""
}

func (sync *MplsLdp_Global_Igp_Sync) GetSegmentPath() string {
    return "sync"
}

func (sync *MplsLdp_Global_Igp_Sync) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delay" {
        return &sync.Delay
    }
    return nil
}

func (sync *MplsLdp_Global_Igp_Sync) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delay"] = &sync.Delay
    return children
}

func (sync *MplsLdp_Global_Igp_Sync) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sync *MplsLdp_Global_Igp_Sync) GetBundleName() string { return "cisco_ios_xr" }

func (sync *MplsLdp_Global_Igp_Sync) GetYangName() string { return "sync" }

func (sync *MplsLdp_Global_Igp_Sync) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sync *MplsLdp_Global_Igp_Sync) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sync *MplsLdp_Global_Igp_Sync) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sync *MplsLdp_Global_Igp_Sync) SetParent(parent types.Entity) { sync.parent = parent }

func (sync *MplsLdp_Global_Igp_Sync) GetParent() types.Entity { return sync.parent }

func (sync *MplsLdp_Global_Igp_Sync) GetParentYangName() string { return "igp" }

// MplsLdp_Global_Igp_Sync_Delay
// LDP IGP synchronization delay time
type MplsLdp_Global_Igp_Sync_Delay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface sync up delay after session up. The type is interface{} with
    // range: 5..300. Units are second.
    OnSessionUp interface{}

    // Global sync up delay to be used after process restart. The type is
    // interface{} with range: 60..600. Units are second.
    OnProcRestart interface{}
}

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetFilter() yfilter.YFilter { return delay.YFilter }

func (delay *MplsLdp_Global_Igp_Sync_Delay) SetFilter(yf yfilter.YFilter) { delay.YFilter = yf }

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetGoName(yname string) string {
    if yname == "on-session-up" { return "OnSessionUp" }
    if yname == "on-proc-restart" { return "OnProcRestart" }
    return ""
}

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetSegmentPath() string {
    return "delay"
}

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["on-session-up"] = delay.OnSessionUp
    leafs["on-proc-restart"] = delay.OnProcRestart
    return leafs
}

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetBundleName() string { return "cisco_ios_xr" }

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetYangName() string { return "delay" }

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delay *MplsLdp_Global_Igp_Sync_Delay) SetParent(parent types.Entity) { delay.parent = parent }

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetParent() types.Entity { return delay.parent }

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetParentYangName() string { return "sync" }

// MplsLdp_Global_EnableLogging
// Enable logging of events
type MplsLdp_Global_EnableLogging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable logging of NSR events. The type is interface{}.
    Nsr interface{}

    // Enable logging of neighbor events. The type is interface{}.
    NeighborChanges interface{}

    // Enable logging of adjacency events. The type is interface{}.
    Adjacency interface{}

    // Enable logging of session protection events. The type is interface{}.
    SessionProtection interface{}

    // Enable logging of Graceful Restart (GR) events. The type is interface{}.
    GrSessionChanges interface{}
}

func (enableLogging *MplsLdp_Global_EnableLogging) GetFilter() yfilter.YFilter { return enableLogging.YFilter }

func (enableLogging *MplsLdp_Global_EnableLogging) SetFilter(yf yfilter.YFilter) { enableLogging.YFilter = yf }

func (enableLogging *MplsLdp_Global_EnableLogging) GetGoName(yname string) string {
    if yname == "nsr" { return "Nsr" }
    if yname == "neighbor-changes" { return "NeighborChanges" }
    if yname == "adjacency" { return "Adjacency" }
    if yname == "session-protection" { return "SessionProtection" }
    if yname == "gr-session-changes" { return "GrSessionChanges" }
    return ""
}

func (enableLogging *MplsLdp_Global_EnableLogging) GetSegmentPath() string {
    return "enable-logging"
}

func (enableLogging *MplsLdp_Global_EnableLogging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (enableLogging *MplsLdp_Global_EnableLogging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (enableLogging *MplsLdp_Global_EnableLogging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nsr"] = enableLogging.Nsr
    leafs["neighbor-changes"] = enableLogging.NeighborChanges
    leafs["adjacency"] = enableLogging.Adjacency
    leafs["session-protection"] = enableLogging.SessionProtection
    leafs["gr-session-changes"] = enableLogging.GrSessionChanges
    return leafs
}

func (enableLogging *MplsLdp_Global_EnableLogging) GetBundleName() string { return "cisco_ios_xr" }

func (enableLogging *MplsLdp_Global_EnableLogging) GetYangName() string { return "enable-logging" }

func (enableLogging *MplsLdp_Global_EnableLogging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enableLogging *MplsLdp_Global_EnableLogging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enableLogging *MplsLdp_Global_EnableLogging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enableLogging *MplsLdp_Global_EnableLogging) SetParent(parent types.Entity) { enableLogging.parent = parent }

func (enableLogging *MplsLdp_Global_EnableLogging) GetParent() types.Entity { return enableLogging.parent }

func (enableLogging *MplsLdp_Global_EnableLogging) GetParentYangName() string { return "global" }

// MplsLdp_Global_Signalling
// Configure LDP signalling parameters
type MplsLdp_Global_Signalling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DSCP for control packets. The type is interface{} with range: 0..63. The
    // default value is 48.
    Dscp interface{}
}

func (signalling *MplsLdp_Global_Signalling) GetFilter() yfilter.YFilter { return signalling.YFilter }

func (signalling *MplsLdp_Global_Signalling) SetFilter(yf yfilter.YFilter) { signalling.YFilter = yf }

func (signalling *MplsLdp_Global_Signalling) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (signalling *MplsLdp_Global_Signalling) GetSegmentPath() string {
    return "signalling"
}

func (signalling *MplsLdp_Global_Signalling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (signalling *MplsLdp_Global_Signalling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (signalling *MplsLdp_Global_Signalling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = signalling.Dscp
    return leafs
}

func (signalling *MplsLdp_Global_Signalling) GetBundleName() string { return "cisco_ios_xr" }

func (signalling *MplsLdp_Global_Signalling) GetYangName() string { return "signalling" }

func (signalling *MplsLdp_Global_Signalling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (signalling *MplsLdp_Global_Signalling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (signalling *MplsLdp_Global_Signalling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (signalling *MplsLdp_Global_Signalling) SetParent(parent types.Entity) { signalling.parent = parent }

func (signalling *MplsLdp_Global_Signalling) GetParent() types.Entity { return signalling.parent }

func (signalling *MplsLdp_Global_Signalling) GetParentYangName() string { return "global" }

// MplsLdp_Global_Nsr
// Configure LDP Non-Stop Routing
type MplsLdp_Global_Nsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is interface{}.
    Enable interface{}
}

func (nsr *MplsLdp_Global_Nsr) GetFilter() yfilter.YFilter { return nsr.YFilter }

func (nsr *MplsLdp_Global_Nsr) SetFilter(yf yfilter.YFilter) { nsr.YFilter = yf }

func (nsr *MplsLdp_Global_Nsr) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (nsr *MplsLdp_Global_Nsr) GetSegmentPath() string {
    return "nsr"
}

func (nsr *MplsLdp_Global_Nsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsr *MplsLdp_Global_Nsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsr *MplsLdp_Global_Nsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = nsr.Enable
    return leafs
}

func (nsr *MplsLdp_Global_Nsr) GetBundleName() string { return "cisco_ios_xr" }

func (nsr *MplsLdp_Global_Nsr) GetYangName() string { return "nsr" }

func (nsr *MplsLdp_Global_Nsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nsr *MplsLdp_Global_Nsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nsr *MplsLdp_Global_Nsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nsr *MplsLdp_Global_Nsr) SetParent(parent types.Entity) { nsr.parent = parent }

func (nsr *MplsLdp_Global_Nsr) GetParent() types.Entity { return nsr.parent }

func (nsr *MplsLdp_Global_Nsr) GetParentYangName() string { return "global" }

// MplsLdp_Global_GracefulRestart
// Configuration for LDP Graceful Restart
// parameters
type MplsLdp_Global_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Graceful Restart Reconnect Timeout value. The type is interface{}
    // with range: 60..1800. Units are second. The default value is 120.
    ReconnectTimeout interface{}

    // none. The type is interface{}.
    Enable interface{}

    // Configure Graceful Restart Session holdtime. The type is interface{} with
    // range: 60..1800. Units are second. The default value is 180.
    ForwardingHoldTime interface{}
}

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetGoName(yname string) string {
    if yname == "reconnect-timeout" { return "ReconnectTimeout" }
    if yname == "enable" { return "Enable" }
    if yname == "forwarding-hold-time" { return "ForwardingHoldTime" }
    return ""
}

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reconnect-timeout"] = gracefulRestart.ReconnectTimeout
    leafs["enable"] = gracefulRestart.Enable
    leafs["forwarding-hold-time"] = gracefulRestart.ForwardingHoldTime
    return leafs
}

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetBundleName() string { return "cisco_ios_xr" }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetParentYangName() string { return "global" }

// MplsLdp_Global_Discovery
// Configure Discovery parameters
type MplsLdp_Global_Discovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable transmit and receive processing for private Instance TLV in LDP
    // discovery hello messages. The type is interface{}.
    DisableInstanceTlv interface{}

    // Disable discovery's quick start mode. The type is interface{}.
    DisableQuickStart interface{}

    // LDP Link Hellos.
    LinkHello MplsLdp_Global_Discovery_LinkHello

    // LDP Targeted Hellos.
    TargetedHello MplsLdp_Global_Discovery_TargetedHello
}

func (discovery *MplsLdp_Global_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_Global_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_Global_Discovery) GetGoName(yname string) string {
    if yname == "disable-instance-tlv" { return "DisableInstanceTlv" }
    if yname == "disable-quick-start" { return "DisableQuickStart" }
    if yname == "link-hello" { return "LinkHello" }
    if yname == "targeted-hello" { return "TargetedHello" }
    return ""
}

func (discovery *MplsLdp_Global_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_Global_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-hello" {
        return &discovery.LinkHello
    }
    if childYangName == "targeted-hello" {
        return &discovery.TargetedHello
    }
    return nil
}

func (discovery *MplsLdp_Global_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-hello"] = &discovery.LinkHello
    children["targeted-hello"] = &discovery.TargetedHello
    return children
}

func (discovery *MplsLdp_Global_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable-instance-tlv"] = discovery.DisableInstanceTlv
    leafs["disable-quick-start"] = discovery.DisableQuickStart
    return leafs
}

func (discovery *MplsLdp_Global_Discovery) GetBundleName() string { return "cisco_ios_xr" }

func (discovery *MplsLdp_Global_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_Global_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discovery *MplsLdp_Global_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discovery *MplsLdp_Global_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discovery *MplsLdp_Global_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_Global_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_Global_Discovery) GetParentYangName() string { return "global" }

// MplsLdp_Global_Discovery_LinkHello
// LDP Link Hellos
type MplsLdp_Global_Discovery_LinkHello struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Hello interval. The type is interface{} with range: 1..65535. Units
    // are second. The default value is 5.
    Interval interface{}

    // Time (seconds) - 65535 implies infinite. The type is interface{} with
    // range: 1..65535. Units are second. The default value is 15.
    HoldTime interface{}
}

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetFilter() yfilter.YFilter { return linkHello.YFilter }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) SetFilter(yf yfilter.YFilter) { linkHello.YFilter = yf }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "hold-time" { return "HoldTime" }
    return ""
}

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetSegmentPath() string {
    return "link-hello"
}

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = linkHello.Interval
    leafs["hold-time"] = linkHello.HoldTime
    return leafs
}

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetBundleName() string { return "cisco_ios_xr" }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetYangName() string { return "link-hello" }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) SetParent(parent types.Entity) { linkHello.parent = parent }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetParent() types.Entity { return linkHello.parent }

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetParentYangName() string { return "discovery" }

// MplsLdp_Global_Discovery_TargetedHello
// LDP Targeted Hellos
type MplsLdp_Global_Discovery_TargetedHello struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Targeted Hello interval. The type is interface{} with range: 1..65535.
    // Units are second. The default value is 10.
    Interval interface{}

    // Time (seconds) - 65535 implies infinite. The type is interface{} with
    // range: 1..65535. Units are second. The default value is 90.
    HoldTime interface{}
}

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetFilter() yfilter.YFilter { return targetedHello.YFilter }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) SetFilter(yf yfilter.YFilter) { targetedHello.YFilter = yf }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "hold-time" { return "HoldTime" }
    return ""
}

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetSegmentPath() string {
    return "targeted-hello"
}

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = targetedHello.Interval
    leafs["hold-time"] = targetedHello.HoldTime
    return leafs
}

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetBundleName() string { return "cisco_ios_xr" }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetYangName() string { return "targeted-hello" }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) SetParent(parent types.Entity) { targetedHello.parent = parent }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetParent() types.Entity { return targetedHello.parent }

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetParentYangName() string { return "discovery" }

// MplsLdp_Global_Mldp
// MPLS mLDP configuration
type MplsLdp_Global_Mldp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Multicast Label Distribution Protocol (mLDP). The type is
    // interface{}.
    Enable interface{}

    // VRF Table attribute configuration for MPLS LDP.
    Vrfs MplsLdp_Global_Mldp_Vrfs

    // Default VRF attribute configuration for mLDP.
    DefaultVrf MplsLdp_Global_Mldp_DefaultVrf

    // Global configuration for mLDP.
    MldpGlobal MplsLdp_Global_Mldp_MldpGlobal
}

func (mldp *MplsLdp_Global_Mldp) GetFilter() yfilter.YFilter { return mldp.YFilter }

func (mldp *MplsLdp_Global_Mldp) SetFilter(yf yfilter.YFilter) { mldp.YFilter = yf }

func (mldp *MplsLdp_Global_Mldp) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "default-vrf" { return "DefaultVrf" }
    if yname == "mldp-global" { return "MldpGlobal" }
    return ""
}

func (mldp *MplsLdp_Global_Mldp) GetSegmentPath() string {
    return "mldp"
}

func (mldp *MplsLdp_Global_Mldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &mldp.Vrfs
    }
    if childYangName == "default-vrf" {
        return &mldp.DefaultVrf
    }
    if childYangName == "mldp-global" {
        return &mldp.MldpGlobal
    }
    return nil
}

func (mldp *MplsLdp_Global_Mldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &mldp.Vrfs
    children["default-vrf"] = &mldp.DefaultVrf
    children["mldp-global"] = &mldp.MldpGlobal
    return children
}

func (mldp *MplsLdp_Global_Mldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = mldp.Enable
    return leafs
}

func (mldp *MplsLdp_Global_Mldp) GetBundleName() string { return "cisco_ios_xr" }

func (mldp *MplsLdp_Global_Mldp) GetYangName() string { return "mldp" }

func (mldp *MplsLdp_Global_Mldp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mldp *MplsLdp_Global_Mldp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mldp *MplsLdp_Global_Mldp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mldp *MplsLdp_Global_Mldp) SetParent(parent types.Entity) { mldp.parent = parent }

func (mldp *MplsLdp_Global_Mldp) GetParent() types.Entity { return mldp.parent }

func (mldp *MplsLdp_Global_Mldp) GetParentYangName() string { return "global" }

// MplsLdp_Global_Mldp_Vrfs
// VRF Table attribute configuration for MPLS LDP
type MplsLdp_Global_Mldp_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF attribute configuration for MPLS LDP. The type is slice of
    // MplsLdp_Global_Mldp_Vrfs_Vrf.
    Vrf []MplsLdp_Global_Mldp_Vrfs_Vrf
}

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Global_Mldp_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetParentYangName() string { return "mldp" }

// MplsLdp_Global_Mldp_Vrfs_Vrf
// VRF attribute configuration for MPLS LDP
type MplsLdp_Global_Mldp_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with length: 1..32.
    VrfName interface{}

    // Enable Multicast Label Distribution Protocol (mLDP). The type is
    // interface{}.
    Enable interface{}

    // Address Family specific operational data.
    Afs MplsLdp_Global_Mldp_Vrfs_Vrf_Afs
}

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "enable" { return "Enable" }
    if yname == "afs" { return "Afs" }
    return ""
}

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &vrf.Afs
    }
    return nil
}

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &vrf.Afs
    return children
}

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["enable"] = vrf.Enable
    return leafs
}

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs
// Address Family specific operational data
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af.
    Af []MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af
}

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetParentYangName() string { return "vrf" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af
// Operational data for given Address Family
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is MplsLdpafName.
    AfName interface{}

    // Enable Multicast Label Distribution Protocol (mLDP) under AF. The type is
    // interface{}.
    Enable interface{}

    // Enable MPLS MLDP RIB unicast-always configuration. The type is interface{}.
    MldpRibUnicastAlways interface{}

    // Enable recursive forwarding.
    RecursiveForwarding MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding

    // MPLS mLDP Recursive FEC.
    MldpRecursiveFec MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec

    // MLDP neighbor policies.
    NeighborPolicies MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies

    // MPLS mLDP MoFRR.
    MoFrr MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr

    // MPLS mLDP Make-Before-Break configuration.
    MakeBeforeBreak MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak

    // MPLS mLDP CSC.
    Csc MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc
}

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "enable" { return "Enable" }
    if yname == "mldp-rib-unicast-always" { return "MldpRibUnicastAlways" }
    if yname == "recursive-forwarding" { return "RecursiveForwarding" }
    if yname == "mldp-recursive-fec" { return "MldpRecursiveFec" }
    if yname == "neighbor-policies" { return "NeighborPolicies" }
    if yname == "mo-frr" { return "MoFrr" }
    if yname == "make-before-break" { return "MakeBeforeBreak" }
    if yname == "csc" { return "Csc" }
    return ""
}

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "recursive-forwarding" {
        return &af.RecursiveForwarding
    }
    if childYangName == "mldp-recursive-fec" {
        return &af.MldpRecursiveFec
    }
    if childYangName == "neighbor-policies" {
        return &af.NeighborPolicies
    }
    if childYangName == "mo-frr" {
        return &af.MoFrr
    }
    if childYangName == "make-before-break" {
        return &af.MakeBeforeBreak
    }
    if childYangName == "csc" {
        return &af.Csc
    }
    return nil
}

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["recursive-forwarding"] = &af.RecursiveForwarding
    children["mldp-recursive-fec"] = &af.MldpRecursiveFec
    children["neighbor-policies"] = &af.NeighborPolicies
    children["mo-frr"] = &af.MoFrr
    children["make-before-break"] = &af.MakeBeforeBreak
    children["csc"] = &af.Csc
    return children
}

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["enable"] = af.Enable
    leafs["mldp-rib-unicast-always"] = af.MldpRibUnicastAlways
    return leafs
}

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding
// Enable recursive forwarding
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable recursive forwarding. The type is interface{}.
    Enable interface{}

    // Recursive forwarding policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetFilter() yfilter.YFilter { return recursiveForwarding.YFilter }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) SetFilter(yf yfilter.YFilter) { recursiveForwarding.YFilter = yf }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetSegmentPath() string {
    return "recursive-forwarding"
}

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = recursiveForwarding.Enable
    leafs["policy"] = recursiveForwarding.Policy
    return leafs
}

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetBundleName() string { return "cisco_ios_xr" }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetYangName() string { return "recursive-forwarding" }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) SetParent(parent types.Entity) { recursiveForwarding.parent = parent }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetParent() types.Entity { return recursiveForwarding.parent }

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec
// MPLS mLDP Recursive FEC
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS mLDP Recursive FEC. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetFilter() yfilter.YFilter { return mldpRecursiveFec.YFilter }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) SetFilter(yf yfilter.YFilter) { mldpRecursiveFec.YFilter = yf }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetSegmentPath() string {
    return "mldp-recursive-fec"
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = mldpRecursiveFec.Enable
    leafs["policy"] = mldpRecursiveFec.Policy
    return leafs
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetBundleName() string { return "cisco_ios_xr" }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetYangName() string { return "mldp-recursive-fec" }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) SetParent(parent types.Entity) { mldpRecursiveFec.parent = parent }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetParent() types.Entity { return mldpRecursiveFec.parent }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies
// MLDP neighbor policies
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Policy. The type is slice of
    // MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy.
    NeighborPolicy []MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy
}

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetFilter() yfilter.YFilter { return neighborPolicies.YFilter }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) SetFilter(yf yfilter.YFilter) { neighborPolicies.YFilter = yf }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetGoName(yname string) string {
    if yname == "neighbor-policy" { return "NeighborPolicy" }
    return ""
}

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetSegmentPath() string {
    return "neighbor-policies"
}

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor-policy" {
        for _, c := range neighborPolicies.NeighborPolicy {
            if neighborPolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy{}
        neighborPolicies.NeighborPolicy = append(neighborPolicies.NeighborPolicy, child)
        return &neighborPolicies.NeighborPolicy[len(neighborPolicies.NeighborPolicy)-1]
    }
    return nil
}

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighborPolicies.NeighborPolicy {
        children[neighborPolicies.NeighborPolicy[i].GetSegmentPath()] = &neighborPolicies.NeighborPolicy[i]
    }
    return children
}

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetYangName() string { return "neighbor-policies" }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) SetParent(parent types.Entity) { neighborPolicies.parent = parent }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetParent() types.Entity { return neighborPolicies.parent }

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy
// Route Policy
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RootAddress interface{}

    // This attribute is a key. Inbound/Outbound Policy. The type is
    // MldpPolicyMode.
    PolicyMode interface{}

    // Route policy name. The type is string with length: 1..64. This attribute is
    // mandatory.
    RoutePolicy interface{}
}

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetFilter() yfilter.YFilter { return neighborPolicy.YFilter }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) SetFilter(yf yfilter.YFilter) { neighborPolicy.YFilter = yf }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetGoName(yname string) string {
    if yname == "root-address" { return "RootAddress" }
    if yname == "policy-mode" { return "PolicyMode" }
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetSegmentPath() string {
    return "neighbor-policy" + "[root-address='" + fmt.Sprintf("%v", neighborPolicy.RootAddress) + "']" + "[policy-mode='" + fmt.Sprintf("%v", neighborPolicy.PolicyMode) + "']"
}

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["root-address"] = neighborPolicy.RootAddress
    leafs["policy-mode"] = neighborPolicy.PolicyMode
    leafs["route-policy"] = neighborPolicy.RoutePolicy
    return leafs
}

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetYangName() string { return "neighbor-policy" }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) SetParent(parent types.Entity) { neighborPolicy.parent = parent }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetParent() types.Entity { return neighborPolicy.parent }

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetParentYangName() string { return "neighbor-policies" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr
// MPLS mLDP MoFRR
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS mLDP MoFRR. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetFilter() yfilter.YFilter { return moFrr.YFilter }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) SetFilter(yf yfilter.YFilter) { moFrr.YFilter = yf }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetSegmentPath() string {
    return "mo-frr"
}

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = moFrr.Enable
    leafs["policy"] = moFrr.Policy
    return leafs
}

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetBundleName() string { return "cisco_ios_xr" }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetYangName() string { return "mo-frr" }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) SetParent(parent types.Entity) { moFrr.parent = parent }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetParent() types.Entity { return moFrr.parent }

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak
// MPLS mLDP Make-Before-Break configuration
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}

    // Enable MPLS mLDP MBB signaling.
    Signaling MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetFilter() yfilter.YFilter { return makeBeforeBreak.YFilter }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) SetFilter(yf yfilter.YFilter) { makeBeforeBreak.YFilter = yf }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    if yname == "signaling" { return "Signaling" }
    return ""
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetSegmentPath() string {
    return "make-before-break"
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "signaling" {
        return &makeBeforeBreak.Signaling
    }
    return nil
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["signaling"] = &makeBeforeBreak.Signaling
    return children
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = makeBeforeBreak.Policy
    return leafs
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetBundleName() string { return "cisco_ios_xr" }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetYangName() string { return "make-before-break" }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) SetParent(parent types.Entity) { makeBeforeBreak.parent = parent }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetParent() types.Entity { return makeBeforeBreak.parent }

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling
// Enable MPLS mLDP MBB signaling
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Forwarding Delay in Seconds. The type is interface{} with range: 0..600.
    // Units are second.
    ForwardDelay interface{}

    // Delete Delay in seconds. The type is interface{} with range: 0..60. Units
    // are second.
    DeleteDelay interface{}
}

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetFilter() yfilter.YFilter { return signaling.YFilter }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) SetFilter(yf yfilter.YFilter) { signaling.YFilter = yf }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetGoName(yname string) string {
    if yname == "forward-delay" { return "ForwardDelay" }
    if yname == "delete-delay" { return "DeleteDelay" }
    return ""
}

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetSegmentPath() string {
    return "signaling"
}

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["forward-delay"] = signaling.ForwardDelay
    leafs["delete-delay"] = signaling.DeleteDelay
    return leafs
}

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetBundleName() string { return "cisco_ios_xr" }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetYangName() string { return "signaling" }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) SetParent(parent types.Entity) { signaling.parent = parent }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetParent() types.Entity { return signaling.parent }

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetParentYangName() string { return "make-before-break" }

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc
// MPLS mLDP CSC
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS mLDP CSC. The type is interface{}.
    Enable interface{}
}

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetFilter() yfilter.YFilter { return csc.YFilter }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) SetFilter(yf yfilter.YFilter) { csc.YFilter = yf }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetSegmentPath() string {
    return "csc"
}

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = csc.Enable
    return leafs
}

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetBundleName() string { return "cisco_ios_xr" }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetYangName() string { return "csc" }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) SetParent(parent types.Entity) { csc.parent = parent }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetParent() types.Entity { return csc.parent }

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_DefaultVrf
// Default VRF attribute configuration for mLDP
type MplsLdp_Global_Mldp_DefaultVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family specific operational data.
    Afs MplsLdp_Global_Mldp_DefaultVrf_Afs
}

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetGoName(yname string) string {
    if yname == "afs" { return "Afs" }
    return ""
}

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &defaultVrf.Afs
    }
    return nil
}

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &defaultVrf.Afs
    return children
}

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetParentYangName() string { return "mldp" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs
// Address Family specific operational data
type MplsLdp_Global_Mldp_DefaultVrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // MplsLdp_Global_Mldp_DefaultVrf_Afs_Af.
    Af []MplsLdp_Global_Mldp_DefaultVrf_Afs_Af
}

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Global_Mldp_DefaultVrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetParentYangName() string { return "default-vrf" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af
// Operational data for given Address Family
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is MplsLdpafName.
    AfName interface{}

    // Enable Multicast Label Distribution Protocol (mLDP) under AF. The type is
    // interface{}.
    Enable interface{}

    // Enable MPLS MLDP RIB unicast-always configuration. The type is interface{}.
    MldpRibUnicastAlways interface{}

    // Enable recursive forwarding.
    RecursiveForwarding MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding

    // MPLS mLDP Recursive FEC.
    MldpRecursiveFec MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec

    // MLDP neighbor policies.
    NeighborPolicies MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies

    // MPLS mLDP MoFRR.
    MoFrr MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr

    // MPLS mLDP Make-Before-Break configuration.
    MakeBeforeBreak MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak

    // MPLS mLDP CSC.
    Csc MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc
}

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "enable" { return "Enable" }
    if yname == "mldp-rib-unicast-always" { return "MldpRibUnicastAlways" }
    if yname == "recursive-forwarding" { return "RecursiveForwarding" }
    if yname == "mldp-recursive-fec" { return "MldpRecursiveFec" }
    if yname == "neighbor-policies" { return "NeighborPolicies" }
    if yname == "mo-frr" { return "MoFrr" }
    if yname == "make-before-break" { return "MakeBeforeBreak" }
    if yname == "csc" { return "Csc" }
    return ""
}

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "recursive-forwarding" {
        return &af.RecursiveForwarding
    }
    if childYangName == "mldp-recursive-fec" {
        return &af.MldpRecursiveFec
    }
    if childYangName == "neighbor-policies" {
        return &af.NeighborPolicies
    }
    if childYangName == "mo-frr" {
        return &af.MoFrr
    }
    if childYangName == "make-before-break" {
        return &af.MakeBeforeBreak
    }
    if childYangName == "csc" {
        return &af.Csc
    }
    return nil
}

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["recursive-forwarding"] = &af.RecursiveForwarding
    children["mldp-recursive-fec"] = &af.MldpRecursiveFec
    children["neighbor-policies"] = &af.NeighborPolicies
    children["mo-frr"] = &af.MoFrr
    children["make-before-break"] = &af.MakeBeforeBreak
    children["csc"] = &af.Csc
    return children
}

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["enable"] = af.Enable
    leafs["mldp-rib-unicast-always"] = af.MldpRibUnicastAlways
    return leafs
}

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding
// Enable recursive forwarding
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable recursive forwarding. The type is interface{}.
    Enable interface{}

    // Recursive forwarding policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetFilter() yfilter.YFilter { return recursiveForwarding.YFilter }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) SetFilter(yf yfilter.YFilter) { recursiveForwarding.YFilter = yf }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetSegmentPath() string {
    return "recursive-forwarding"
}

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = recursiveForwarding.Enable
    leafs["policy"] = recursiveForwarding.Policy
    return leafs
}

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetBundleName() string { return "cisco_ios_xr" }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetYangName() string { return "recursive-forwarding" }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) SetParent(parent types.Entity) { recursiveForwarding.parent = parent }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetParent() types.Entity { return recursiveForwarding.parent }

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec
// MPLS mLDP Recursive FEC
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS mLDP Recursive FEC. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetFilter() yfilter.YFilter { return mldpRecursiveFec.YFilter }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) SetFilter(yf yfilter.YFilter) { mldpRecursiveFec.YFilter = yf }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetSegmentPath() string {
    return "mldp-recursive-fec"
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = mldpRecursiveFec.Enable
    leafs["policy"] = mldpRecursiveFec.Policy
    return leafs
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetBundleName() string { return "cisco_ios_xr" }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetYangName() string { return "mldp-recursive-fec" }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) SetParent(parent types.Entity) { mldpRecursiveFec.parent = parent }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetParent() types.Entity { return mldpRecursiveFec.parent }

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies
// MLDP neighbor policies
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Policy. The type is slice of
    // MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy.
    NeighborPolicy []MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy
}

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetFilter() yfilter.YFilter { return neighborPolicies.YFilter }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) SetFilter(yf yfilter.YFilter) { neighborPolicies.YFilter = yf }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetGoName(yname string) string {
    if yname == "neighbor-policy" { return "NeighborPolicy" }
    return ""
}

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetSegmentPath() string {
    return "neighbor-policies"
}

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor-policy" {
        for _, c := range neighborPolicies.NeighborPolicy {
            if neighborPolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy{}
        neighborPolicies.NeighborPolicy = append(neighborPolicies.NeighborPolicy, child)
        return &neighborPolicies.NeighborPolicy[len(neighborPolicies.NeighborPolicy)-1]
    }
    return nil
}

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighborPolicies.NeighborPolicy {
        children[neighborPolicies.NeighborPolicy[i].GetSegmentPath()] = &neighborPolicies.NeighborPolicy[i]
    }
    return children
}

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetYangName() string { return "neighbor-policies" }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) SetParent(parent types.Entity) { neighborPolicies.parent = parent }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetParent() types.Entity { return neighborPolicies.parent }

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy
// Route Policy
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RootAddress interface{}

    // This attribute is a key. Inbound/Outbound Policy. The type is
    // MldpPolicyMode.
    PolicyMode interface{}

    // Route policy name. The type is string with length: 1..64. This attribute is
    // mandatory.
    RoutePolicy interface{}
}

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetFilter() yfilter.YFilter { return neighborPolicy.YFilter }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) SetFilter(yf yfilter.YFilter) { neighborPolicy.YFilter = yf }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetGoName(yname string) string {
    if yname == "root-address" { return "RootAddress" }
    if yname == "policy-mode" { return "PolicyMode" }
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetSegmentPath() string {
    return "neighbor-policy" + "[root-address='" + fmt.Sprintf("%v", neighborPolicy.RootAddress) + "']" + "[policy-mode='" + fmt.Sprintf("%v", neighborPolicy.PolicyMode) + "']"
}

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["root-address"] = neighborPolicy.RootAddress
    leafs["policy-mode"] = neighborPolicy.PolicyMode
    leafs["route-policy"] = neighborPolicy.RoutePolicy
    return leafs
}

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetYangName() string { return "neighbor-policy" }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) SetParent(parent types.Entity) { neighborPolicy.parent = parent }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetParent() types.Entity { return neighborPolicy.parent }

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetParentYangName() string { return "neighbor-policies" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr
// MPLS mLDP MoFRR
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS mLDP MoFRR. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetFilter() yfilter.YFilter { return moFrr.YFilter }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) SetFilter(yf yfilter.YFilter) { moFrr.YFilter = yf }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetSegmentPath() string {
    return "mo-frr"
}

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = moFrr.Enable
    leafs["policy"] = moFrr.Policy
    return leafs
}

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetBundleName() string { return "cisco_ios_xr" }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetYangName() string { return "mo-frr" }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) SetParent(parent types.Entity) { moFrr.parent = parent }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetParent() types.Entity { return moFrr.parent }

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak
// MPLS mLDP Make-Before-Break configuration
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}

    // Enable MPLS mLDP MBB signaling.
    Signaling MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetFilter() yfilter.YFilter { return makeBeforeBreak.YFilter }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) SetFilter(yf yfilter.YFilter) { makeBeforeBreak.YFilter = yf }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    if yname == "signaling" { return "Signaling" }
    return ""
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetSegmentPath() string {
    return "make-before-break"
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "signaling" {
        return &makeBeforeBreak.Signaling
    }
    return nil
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["signaling"] = &makeBeforeBreak.Signaling
    return children
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = makeBeforeBreak.Policy
    return leafs
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetBundleName() string { return "cisco_ios_xr" }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetYangName() string { return "make-before-break" }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) SetParent(parent types.Entity) { makeBeforeBreak.parent = parent }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetParent() types.Entity { return makeBeforeBreak.parent }

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling
// Enable MPLS mLDP MBB signaling
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Forwarding Delay in Seconds. The type is interface{} with range: 0..600.
    // Units are second.
    ForwardDelay interface{}

    // Delete Delay in seconds. The type is interface{} with range: 0..60. Units
    // are second.
    DeleteDelay interface{}
}

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetFilter() yfilter.YFilter { return signaling.YFilter }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) SetFilter(yf yfilter.YFilter) { signaling.YFilter = yf }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetGoName(yname string) string {
    if yname == "forward-delay" { return "ForwardDelay" }
    if yname == "delete-delay" { return "DeleteDelay" }
    return ""
}

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetSegmentPath() string {
    return "signaling"
}

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["forward-delay"] = signaling.ForwardDelay
    leafs["delete-delay"] = signaling.DeleteDelay
    return leafs
}

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetBundleName() string { return "cisco_ios_xr" }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetYangName() string { return "signaling" }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) SetParent(parent types.Entity) { signaling.parent = parent }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetParent() types.Entity { return signaling.parent }

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetParentYangName() string { return "make-before-break" }

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc
// MPLS mLDP CSC
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS mLDP CSC. The type is interface{}.
    Enable interface{}
}

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetFilter() yfilter.YFilter { return csc.YFilter }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) SetFilter(yf yfilter.YFilter) { csc.YFilter = yf }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetSegmentPath() string {
    return "csc"
}

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = csc.Enable
    return leafs
}

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetBundleName() string { return "cisco_ios_xr" }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetYangName() string { return "csc" }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) SetParent(parent types.Entity) { csc.parent = parent }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetParent() types.Entity { return csc.parent }

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetParentYangName() string { return "af" }

// MplsLdp_Global_Mldp_MldpGlobal
// Global configuration for mLDP
type MplsLdp_Global_Mldp_MldpGlobal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS mLDP logging.
    Logging MplsLdp_Global_Mldp_MldpGlobal_Logging
}

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetFilter() yfilter.YFilter { return mldpGlobal.YFilter }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) SetFilter(yf yfilter.YFilter) { mldpGlobal.YFilter = yf }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetGoName(yname string) string {
    if yname == "logging" { return "Logging" }
    return ""
}

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetSegmentPath() string {
    return "mldp-global"
}

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logging" {
        return &mldpGlobal.Logging
    }
    return nil
}

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["logging"] = &mldpGlobal.Logging
    return children
}

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetBundleName() string { return "cisco_ios_xr" }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetYangName() string { return "mldp-global" }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) SetParent(parent types.Entity) { mldpGlobal.parent = parent }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetParent() types.Entity { return mldpGlobal.parent }

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetParentYangName() string { return "mldp" }

// MplsLdp_Global_Mldp_MldpGlobal_Logging
// MPLS mLDP logging
type MplsLdp_Global_Mldp_MldpGlobal_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS mLDP logging notifications. The type is interface{}.
    Notifications interface{}
}

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetGoName(yname string) string {
    if yname == "notifications" { return "Notifications" }
    return ""
}

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["notifications"] = logging.Notifications
    return leafs
}

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetYangName() string { return "logging" }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetParent() types.Entity { return logging.parent }

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetParentYangName() string { return "mldp-global" }

