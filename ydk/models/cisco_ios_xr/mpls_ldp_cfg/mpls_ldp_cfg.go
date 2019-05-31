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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// MplsLdpNbrPassword represents Mpls ldp nbr password
type MplsLdpNbrPassword string

const (
    // Disable the global default password for this
    // neighbor
    MplsLdpNbrPassword_disable MplsLdpNbrPassword = "disable"

    // Specify a password for this neighbor
    MplsLdpNbrPassword_specified MplsLdpNbrPassword = "specified"
)

// MplsLdpLabelAllocation represents Mpls ldp label allocation
type MplsLdpLabelAllocation string

const (
    // Allocate label for prefixes permitted by ACL
    MplsLdpLabelAllocation_acl MplsLdpLabelAllocation = "acl"

    // Allocate label for host routes only
    MplsLdpLabelAllocation_host MplsLdpLabelAllocation = "host"
)

// MplsLdpDownstreamOnDemand represents Mpls ldp downstream on demand
type MplsLdpDownstreamOnDemand string

const (
    // Downstream on Demand peers permitted by ACL
    MplsLdpDownstreamOnDemand_peer_acl MplsLdpDownstreamOnDemand = "peer-acl"
)

// MldpPolicyMode represents Mldp policy mode
type MldpPolicyMode string

const (
    // Inbound route policy
    MldpPolicyMode_inbound MldpPolicyMode = "inbound"

    // Outbound route policy
    MldpPolicyMode_outbound MldpPolicyMode = "outbound"
)

// MplsLdpTargetedAccept represents Mpls ldp targeted accept
type MplsLdpTargetedAccept string

const (
    // Accept targeted hello from all
    MplsLdpTargetedAccept_all MplsLdpTargetedAccept = "all"

    // Accept targeted hello from peer ACL
    MplsLdpTargetedAccept_from MplsLdpTargetedAccept = "from"
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

// MplsLdpAdvertiseBgpAcl represents Mpls ldp advertise bgp acl
type MplsLdpAdvertiseBgpAcl string

const (
    // BGP prefixes advertised to peers permitted by
    // ACL
    MplsLdpAdvertiseBgpAcl_peer_acl MplsLdpAdvertiseBgpAcl = "peer-acl"
)

// MplsLdp
// MPLS LDP configuration
// This type is a presence type.
type MplsLdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable Label Distribution Protocol (LDP) globally.Without creating this
    // object the LDP feature will not be enabled. Deleting this object will stop
    // the LDP feature. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Global VRF attribute configuration for MPLS LDP.
    DefaultVrf MplsLdp_DefaultVrf

    // VRF Table attribute configuration for MPLS LDP.
    Vrfs MplsLdp_Vrfs

    // Global configuration for MPLS LDP.
    Global MplsLdp_Global
}

func (mplsLdp *MplsLdp) GetEntityData() *types.CommonEntityData {
    mplsLdp.EntityData.YFilter = mplsLdp.YFilter
    mplsLdp.EntityData.YangName = "mpls-ldp"
    mplsLdp.EntityData.BundleName = "cisco_ios_xr"
    mplsLdp.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-ldp-cfg"
    mplsLdp.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp"
    mplsLdp.EntityData.AbsolutePath = mplsLdp.EntityData.SegmentPath
    mplsLdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLdp.EntityData.Children = types.NewOrderedMap()
    mplsLdp.EntityData.Children.Append("default-vrf", types.YChild{"DefaultVrf", &mplsLdp.DefaultVrf})
    mplsLdp.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &mplsLdp.Vrfs})
    mplsLdp.EntityData.Children.Append("global", types.YChild{"Global", &mplsLdp.Global})
    mplsLdp.EntityData.Leafs = types.NewOrderedMap()
    mplsLdp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mplsLdp.Enable})

    mplsLdp.EntityData.YListKeys = []string {}

    return &(mplsLdp.EntityData)
}

// MplsLdp_DefaultVrf
// Global VRF attribute configuration for MPLS LDP
type MplsLdp_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family specific configuration for MPLS LDP.
    Afs MplsLdp_DefaultVrf_Afs

    // Default VRF Global configuration for MPLS LDP.
    Global MplsLdp_DefaultVrf_Global

    // MPLS LDP configuration pertaining to interfaces.
    Interfaces MplsLdp_DefaultVrf_Interfaces
}

func (defaultVrf *MplsLdp_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "mpls-ldp"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/" + defaultVrf.EntityData.SegmentPath
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = types.NewOrderedMap()
    defaultVrf.EntityData.Children.Append("afs", types.YChild{"Afs", &defaultVrf.Afs})
    defaultVrf.EntityData.Children.Append("global", types.YChild{"Global", &defaultVrf.Global})
    defaultVrf.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &defaultVrf.Interfaces})
    defaultVrf.EntityData.Leafs = types.NewOrderedMap()

    defaultVrf.EntityData.YListKeys = []string {}

    return &(defaultVrf.EntityData)
}

// MplsLdp_DefaultVrf_Afs
// Address Family specific configuration for MPLS
// LDP
type MplsLdp_DefaultVrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af.
    Af []*MplsLdp_DefaultVrf_Afs_Af
}

func (afs *MplsLdp_DefaultVrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "default-vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af
// Configure data for given Address Family
type MplsLdp_DefaultVrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (af *MplsLdp_DefaultVrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("label", types.YChild{"Label", &af.Label})
    af.EntityData.Children.Append("discovery", types.YChild{"Discovery", &af.Discovery})
    af.EntityData.Children.Append("traffic-engineering", types.YChild{"TrafficEngineering", &af.TrafficEngineering})
    af.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &af.Neighbor})
    af.EntityData.Children.Append("redistribution-protocol", types.YChild{"RedistributionProtocol", &af.RedistributionProtocol})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label
// Configure Label policies and control
type MplsLdp_DefaultVrf_Afs_Af_Label struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure remote/peer label policies and control.
    Remote MplsLdp_DefaultVrf_Afs_Af_Label_Remote

    // Configure local label policies and control.
    Local MplsLdp_DefaultVrf_Afs_Af_Label_Local
}

func (label *MplsLdp_DefaultVrf_Afs_Af_Label) GetEntityData() *types.CommonEntityData {
    label.EntityData.YFilter = label.YFilter
    label.EntityData.YangName = "label"
    label.EntityData.BundleName = "cisco_ios_xr"
    label.EntityData.ParentYangName = "af"
    label.EntityData.SegmentPath = "label"
    label.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/" + label.EntityData.SegmentPath
    label.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    label.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    label.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    label.EntityData.Children = types.NewOrderedMap()
    label.EntityData.Children.Append("remote", types.YChild{"Remote", &label.Remote})
    label.EntityData.Children.Append("local", types.YChild{"Local", &label.Local})
    label.EntityData.Leafs = types.NewOrderedMap()

    label.EntityData.YListKeys = []string {}

    return &(label.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote
// Configure remote/peer label policies and
// control
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure inbound label acceptance.
    Accept MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept
}

func (remote *MplsLdp_DefaultVrf_Afs_Af_Label_Remote) GetEntityData() *types.CommonEntityData {
    remote.EntityData.YFilter = remote.YFilter
    remote.EntityData.YangName = "remote"
    remote.EntityData.BundleName = "cisco_ios_xr"
    remote.EntityData.ParentYangName = "label"
    remote.EntityData.SegmentPath = "remote"
    remote.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/" + remote.EntityData.SegmentPath
    remote.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remote.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remote.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remote.EntityData.Children = types.NewOrderedMap()
    remote.EntityData.Children.Append("accept", types.YChild{"Accept", &remote.Accept})
    remote.EntityData.Leafs = types.NewOrderedMap()

    remote.EntityData.YListKeys = []string {}

    return &(remote.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept
// Configure inbound label acceptance
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to neighbors for inbound label acceptance.
    PeerAcceptPolicies MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
}

func (accept *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept) GetEntityData() *types.CommonEntityData {
    accept.EntityData.YFilter = accept.YFilter
    accept.EntityData.YangName = "accept"
    accept.EntityData.BundleName = "cisco_ios_xr"
    accept.EntityData.ParentYangName = "remote"
    accept.EntityData.SegmentPath = "accept"
    accept.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/remote/" + accept.EntityData.SegmentPath
    accept.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accept.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accept.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accept.EntityData.Children = types.NewOrderedMap()
    accept.EntityData.Children.Append("peer-accept-policies", types.YChild{"PeerAcceptPolicies", &accept.PeerAcceptPolicies})
    accept.EntityData.Leafs = types.NewOrderedMap()

    accept.EntityData.YListKeys = []string {}

    return &(accept.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
// Configuration related to neighbors for
// inbound label acceptance
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control acceptance of labels from a neighbor for prefix(es) using ACL. The
    // type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy.
    PeerAcceptPolicy []*MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
}

func (peerAcceptPolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetEntityData() *types.CommonEntityData {
    peerAcceptPolicies.EntityData.YFilter = peerAcceptPolicies.YFilter
    peerAcceptPolicies.EntityData.YangName = "peer-accept-policies"
    peerAcceptPolicies.EntityData.BundleName = "cisco_ios_xr"
    peerAcceptPolicies.EntityData.ParentYangName = "accept"
    peerAcceptPolicies.EntityData.SegmentPath = "peer-accept-policies"
    peerAcceptPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/remote/accept/" + peerAcceptPolicies.EntityData.SegmentPath
    peerAcceptPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAcceptPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAcceptPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAcceptPolicies.EntityData.Children = types.NewOrderedMap()
    peerAcceptPolicies.EntityData.Children.Append("peer-accept-policy", types.YChild{"PeerAcceptPolicy", nil})
    for i := range peerAcceptPolicies.PeerAcceptPolicy {
        peerAcceptPolicies.EntityData.Children.Append(types.GetSegmentPath(peerAcceptPolicies.PeerAcceptPolicy[i]), types.YChild{"PeerAcceptPolicy", peerAcceptPolicies.PeerAcceptPolicy[i]})
    }
    peerAcceptPolicies.EntityData.Leafs = types.NewOrderedMap()

    peerAcceptPolicies.EntityData.YListKeys = []string {}

    return &(peerAcceptPolicies.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
// Control acceptance of labels from a
// neighbor for prefix(es) using ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..0.
    LabelSpaceId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAcceptPolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetEntityData() *types.CommonEntityData {
    peerAcceptPolicy.EntityData.YFilter = peerAcceptPolicy.YFilter
    peerAcceptPolicy.EntityData.YangName = "peer-accept-policy"
    peerAcceptPolicy.EntityData.BundleName = "cisco_ios_xr"
    peerAcceptPolicy.EntityData.ParentYangName = "peer-accept-policies"
    peerAcceptPolicy.EntityData.SegmentPath = "peer-accept-policy" + types.AddKeyToken(peerAcceptPolicy.LsrId, "lsr-id") + types.AddKeyToken(peerAcceptPolicy.LabelSpaceId, "label-space-id")
    peerAcceptPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/remote/accept/peer-accept-policies/" + peerAcceptPolicy.EntityData.SegmentPath
    peerAcceptPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAcceptPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAcceptPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAcceptPolicy.EntityData.Children = types.NewOrderedMap()
    peerAcceptPolicy.EntityData.Leafs = types.NewOrderedMap()
    peerAcceptPolicy.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", peerAcceptPolicy.LsrId})
    peerAcceptPolicy.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", peerAcceptPolicy.LabelSpaceId})
    peerAcceptPolicy.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", peerAcceptPolicy.PrefixAclName})

    peerAcceptPolicy.EntityData.YListKeys = []string {"LsrId", "LabelSpaceId"}

    return &(peerAcceptPolicy.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local
// Configure local label policies and control
type MplsLdp_DefaultVrf_Afs_Af_Label_Local struct {
    EntityData types.CommonEntityData
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

func (local *MplsLdp_DefaultVrf_Afs_Af_Label_Local) GetEntityData() *types.CommonEntityData {
    local.EntityData.YFilter = local.YFilter
    local.EntityData.YangName = "local"
    local.EntityData.BundleName = "cisco_ios_xr"
    local.EntityData.ParentYangName = "label"
    local.EntityData.SegmentPath = "local"
    local.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/" + local.EntityData.SegmentPath
    local.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    local.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    local.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    local.EntityData.Children = types.NewOrderedMap()
    local.EntityData.Children.Append("advertise", types.YChild{"Advertise", &local.Advertise})
    local.EntityData.Children.Append("allocate", types.YChild{"Allocate", &local.Allocate})
    local.EntityData.Leafs = types.NewOrderedMap()
    local.EntityData.Leafs.Append("implicit-null-override", types.YLeaf{"ImplicitNullOverride", local.ImplicitNullOverride})
    local.EntityData.Leafs.Append("default-route", types.YLeaf{"DefaultRoute", local.DefaultRoute})

    local.EntityData.YListKeys = []string {}

    return &(local.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise
// Configure outbound label advertisement
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise struct {
    EntityData types.CommonEntityData
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

func (advertise *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise) GetEntityData() *types.CommonEntityData {
    advertise.EntityData.YFilter = advertise.YFilter
    advertise.EntityData.YangName = "advertise"
    advertise.EntityData.BundleName = "cisco_ios_xr"
    advertise.EntityData.ParentYangName = "local"
    advertise.EntityData.SegmentPath = "advertise"
    advertise.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/" + advertise.EntityData.SegmentPath
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = types.NewOrderedMap()
    advertise.EntityData.Children.Append("peer-advertise-policies", types.YChild{"PeerAdvertisePolicies", &advertise.PeerAdvertisePolicies})
    advertise.EntityData.Children.Append("prefix-advertise-policies", types.YChild{"PrefixAdvertisePolicies", &advertise.PrefixAdvertisePolicies})
    advertise.EntityData.Children.Append("explicit-null", types.YChild{"ExplicitNull", &advertise.ExplicitNull})
    advertise.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &advertise.Interfaces})
    advertise.EntityData.Leafs = types.NewOrderedMap()
    advertise.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", advertise.Disable})

    advertise.EntityData.YListKeys = []string {}

    return &(advertise.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies
// Configure peer centric outbound label
// advertisement using ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control advertisement of prefix(es) using ACL. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy.
    PeerAdvertisePolicy []*MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
}

func (peerAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetEntityData() *types.CommonEntityData {
    peerAdvertisePolicies.EntityData.YFilter = peerAdvertisePolicies.YFilter
    peerAdvertisePolicies.EntityData.YangName = "peer-advertise-policies"
    peerAdvertisePolicies.EntityData.BundleName = "cisco_ios_xr"
    peerAdvertisePolicies.EntityData.ParentYangName = "advertise"
    peerAdvertisePolicies.EntityData.SegmentPath = "peer-advertise-policies"
    peerAdvertisePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/advertise/" + peerAdvertisePolicies.EntityData.SegmentPath
    peerAdvertisePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAdvertisePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAdvertisePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAdvertisePolicies.EntityData.Children = types.NewOrderedMap()
    peerAdvertisePolicies.EntityData.Children.Append("peer-advertise-policy", types.YChild{"PeerAdvertisePolicy", nil})
    for i := range peerAdvertisePolicies.PeerAdvertisePolicy {
        peerAdvertisePolicies.EntityData.Children.Append(types.GetSegmentPath(peerAdvertisePolicies.PeerAdvertisePolicy[i]), types.YChild{"PeerAdvertisePolicy", peerAdvertisePolicies.PeerAdvertisePolicy[i]})
    }
    peerAdvertisePolicies.EntityData.Leafs = types.NewOrderedMap()

    peerAdvertisePolicies.EntityData.YListKeys = []string {}

    return &(peerAdvertisePolicies.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
// Control advertisement of prefix(es) using
// ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..0.
    LabelSpaceId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetEntityData() *types.CommonEntityData {
    peerAdvertisePolicy.EntityData.YFilter = peerAdvertisePolicy.YFilter
    peerAdvertisePolicy.EntityData.YangName = "peer-advertise-policy"
    peerAdvertisePolicy.EntityData.BundleName = "cisco_ios_xr"
    peerAdvertisePolicy.EntityData.ParentYangName = "peer-advertise-policies"
    peerAdvertisePolicy.EntityData.SegmentPath = "peer-advertise-policy" + types.AddKeyToken(peerAdvertisePolicy.LsrId, "lsr-id") + types.AddKeyToken(peerAdvertisePolicy.LabelSpaceId, "label-space-id")
    peerAdvertisePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/advertise/peer-advertise-policies/" + peerAdvertisePolicy.EntityData.SegmentPath
    peerAdvertisePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAdvertisePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAdvertisePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAdvertisePolicy.EntityData.Children = types.NewOrderedMap()
    peerAdvertisePolicy.EntityData.Leafs = types.NewOrderedMap()
    peerAdvertisePolicy.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", peerAdvertisePolicy.LsrId})
    peerAdvertisePolicy.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", peerAdvertisePolicy.LabelSpaceId})
    peerAdvertisePolicy.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", peerAdvertisePolicy.PrefixAclName})

    peerAdvertisePolicy.EntityData.YListKeys = []string {"LsrId", "LabelSpaceId"}

    return &(peerAdvertisePolicy.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies
// Configure prefix centric outbound label
// advertisement using ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control advertisement of prefix(es) using ACL. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy.
    PrefixAdvertisePolicy []*MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy
}

func (prefixAdvertisePolicies *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies) GetEntityData() *types.CommonEntityData {
    prefixAdvertisePolicies.EntityData.YFilter = prefixAdvertisePolicies.YFilter
    prefixAdvertisePolicies.EntityData.YangName = "prefix-advertise-policies"
    prefixAdvertisePolicies.EntityData.BundleName = "cisco_ios_xr"
    prefixAdvertisePolicies.EntityData.ParentYangName = "advertise"
    prefixAdvertisePolicies.EntityData.SegmentPath = "prefix-advertise-policies"
    prefixAdvertisePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/advertise/" + prefixAdvertisePolicies.EntityData.SegmentPath
    prefixAdvertisePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixAdvertisePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixAdvertisePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixAdvertisePolicies.EntityData.Children = types.NewOrderedMap()
    prefixAdvertisePolicies.EntityData.Children.Append("prefix-advertise-policy", types.YChild{"PrefixAdvertisePolicy", nil})
    for i := range prefixAdvertisePolicies.PrefixAdvertisePolicy {
        prefixAdvertisePolicies.EntityData.Children.Append(types.GetSegmentPath(prefixAdvertisePolicies.PrefixAdvertisePolicy[i]), types.YChild{"PrefixAdvertisePolicy", prefixAdvertisePolicies.PrefixAdvertisePolicy[i]})
    }
    prefixAdvertisePolicies.EntityData.Leafs = types.NewOrderedMap()

    prefixAdvertisePolicies.EntityData.YListKeys = []string {}

    return &(prefixAdvertisePolicies.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy
// Control advertisement of prefix(es) using
// ACL
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of prefix ACL. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PrefixAclName interface{}

    // Label advertise type. The type is MplsLdpLabelAdvertise.
    AdvertiseType interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (prefixAdvertisePolicy *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_PrefixAdvertisePolicies_PrefixAdvertisePolicy) GetEntityData() *types.CommonEntityData {
    prefixAdvertisePolicy.EntityData.YFilter = prefixAdvertisePolicy.YFilter
    prefixAdvertisePolicy.EntityData.YangName = "prefix-advertise-policy"
    prefixAdvertisePolicy.EntityData.BundleName = "cisco_ios_xr"
    prefixAdvertisePolicy.EntityData.ParentYangName = "prefix-advertise-policies"
    prefixAdvertisePolicy.EntityData.SegmentPath = "prefix-advertise-policy" + types.AddKeyToken(prefixAdvertisePolicy.PrefixAclName, "prefix-acl-name")
    prefixAdvertisePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/advertise/prefix-advertise-policies/" + prefixAdvertisePolicy.EntityData.SegmentPath
    prefixAdvertisePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixAdvertisePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixAdvertisePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixAdvertisePolicy.EntityData.Children = types.NewOrderedMap()
    prefixAdvertisePolicy.EntityData.Leafs = types.NewOrderedMap()
    prefixAdvertisePolicy.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", prefixAdvertisePolicy.PrefixAclName})
    prefixAdvertisePolicy.EntityData.Leafs.Append("advertise-type", types.YLeaf{"AdvertiseType", prefixAdvertisePolicy.AdvertiseType})
    prefixAdvertisePolicy.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", prefixAdvertisePolicy.PeerAclName})

    prefixAdvertisePolicy.EntityData.YListKeys = []string {"PrefixAclName"}

    return &(prefixAdvertisePolicy.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull
// Configure advertisment of explicit-null
// for connected prefixes.
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Explicit Null command variant. The type is MplsLdpExpNull.
    ExplicitNullType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (explicitNull *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetEntityData() *types.CommonEntityData {
    explicitNull.EntityData.YFilter = explicitNull.YFilter
    explicitNull.EntityData.YangName = "explicit-null"
    explicitNull.EntityData.BundleName = "cisco_ios_xr"
    explicitNull.EntityData.ParentYangName = "advertise"
    explicitNull.EntityData.SegmentPath = "explicit-null"
    explicitNull.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/advertise/" + explicitNull.EntityData.SegmentPath
    explicitNull.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitNull.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitNull.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitNull.EntityData.Children = types.NewOrderedMap()
    explicitNull.EntityData.Leafs = types.NewOrderedMap()
    explicitNull.EntityData.Leafs.Append("explicit-null-type", types.YLeaf{"ExplicitNullType", explicitNull.ExplicitNullType})
    explicitNull.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", explicitNull.PrefixAclName})
    explicitNull.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", explicitNull.PeerAclName})

    explicitNull.EntityData.YListKeys = []string {}

    return &(explicitNull.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces
// Configure outbound label advertisement for
// an interface
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control advertisement of interface's host IP address. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface.
    Interface []*MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
}

func (interfaces *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "advertise"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/advertise/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
// Control advertisement of interface's host
// IP address
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of interface. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (self *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/advertise/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate
// Control local label allocation for
// prefix(es)
type MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label allocation type. The type is MplsLdpLabelAllocation.
    AllocationType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}
}

func (allocate *MplsLdp_DefaultVrf_Afs_Af_Label_Local_Allocate) GetEntityData() *types.CommonEntityData {
    allocate.EntityData.YFilter = allocate.YFilter
    allocate.EntityData.YangName = "allocate"
    allocate.EntityData.BundleName = "cisco_ios_xr"
    allocate.EntityData.ParentYangName = "local"
    allocate.EntityData.SegmentPath = "allocate"
    allocate.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/label/local/" + allocate.EntityData.SegmentPath
    allocate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allocate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allocate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allocate.EntityData.Children = types.NewOrderedMap()
    allocate.EntityData.Leafs = types.NewOrderedMap()
    allocate.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", allocate.AllocationType})
    allocate.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", allocate.PrefixAclName})

    allocate.EntityData.YListKeys = []string {}

    return &(allocate.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Discovery
// Configure Discovery parameters
type MplsLdp_DefaultVrf_Afs_Af_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global discovery transport address for address family. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    TransportAddress interface{}

    // Configure acceptance from and responding to targeted hellos.
    TargetedHelloAccept MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept
}

func (discovery *MplsLdp_DefaultVrf_Afs_Af_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xr"
    discovery.EntityData.ParentYangName = "af"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/" + discovery.EntityData.SegmentPath
    discovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Children.Append("targeted-hello-accept", types.YChild{"TargetedHelloAccept", &discovery.TargetedHelloAccept})
    discovery.EntityData.Leafs = types.NewOrderedMap()
    discovery.EntityData.Leafs.Append("transport-address", types.YLeaf{"TransportAddress", discovery.TransportAddress})

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept
// Configure acceptance from and responding to
// targeted hellos.
type MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of acceptance. The type is MplsLdpTargetedAccept.
    AcceptType interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (targetedHelloAccept *MplsLdp_DefaultVrf_Afs_Af_Discovery_TargetedHelloAccept) GetEntityData() *types.CommonEntityData {
    targetedHelloAccept.EntityData.YFilter = targetedHelloAccept.YFilter
    targetedHelloAccept.EntityData.YangName = "targeted-hello-accept"
    targetedHelloAccept.EntityData.BundleName = "cisco_ios_xr"
    targetedHelloAccept.EntityData.ParentYangName = "discovery"
    targetedHelloAccept.EntityData.SegmentPath = "targeted-hello-accept"
    targetedHelloAccept.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/discovery/" + targetedHelloAccept.EntityData.SegmentPath
    targetedHelloAccept.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetedHelloAccept.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetedHelloAccept.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetedHelloAccept.EntityData.Children = types.NewOrderedMap()
    targetedHelloAccept.EntityData.Leafs = types.NewOrderedMap()
    targetedHelloAccept.EntityData.Leafs.Append("accept-type", types.YLeaf{"AcceptType", targetedHelloAccept.AcceptType})
    targetedHelloAccept.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", targetedHelloAccept.PeerAclName})

    targetedHelloAccept.EntityData.YListKeys = []string {}

    return &(targetedHelloAccept.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering
// MPLS Traffic Engingeering parameters for LDP
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Traffic Engineering auto-tunnel mesh parameters for LDP.
    AutoTunnelMesh MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh
}

func (trafficEngineering *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering) GetEntityData() *types.CommonEntityData {
    trafficEngineering.EntityData.YFilter = trafficEngineering.YFilter
    trafficEngineering.EntityData.YangName = "traffic-engineering"
    trafficEngineering.EntityData.BundleName = "cisco_ios_xr"
    trafficEngineering.EntityData.ParentYangName = "af"
    trafficEngineering.EntityData.SegmentPath = "traffic-engineering"
    trafficEngineering.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/" + trafficEngineering.EntityData.SegmentPath
    trafficEngineering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficEngineering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficEngineering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficEngineering.EntityData.Children = types.NewOrderedMap()
    trafficEngineering.EntityData.Children.Append("auto-tunnel-mesh", types.YChild{"AutoTunnelMesh", &trafficEngineering.AutoTunnelMesh})
    trafficEngineering.EntityData.Leafs = types.NewOrderedMap()

    trafficEngineering.EntityData.YListKeys = []string {}

    return &(trafficEngineering.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh
// MPLS Traffic Engineering auto-tunnel mesh
// parameters for LDP
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable all MPLS TE auto-tunnel mesh-group interfaces. The type is
    // interface{}.
    GroupAll interface{}

    // Enable interfaces in specific MPLS TE auto-tunnel mesh-groups.
    GroupIds MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds
}

func (autoTunnelMesh *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh) GetEntityData() *types.CommonEntityData {
    autoTunnelMesh.EntityData.YFilter = autoTunnelMesh.YFilter
    autoTunnelMesh.EntityData.YangName = "auto-tunnel-mesh"
    autoTunnelMesh.EntityData.BundleName = "cisco_ios_xr"
    autoTunnelMesh.EntityData.ParentYangName = "traffic-engineering"
    autoTunnelMesh.EntityData.SegmentPath = "auto-tunnel-mesh"
    autoTunnelMesh.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/traffic-engineering/" + autoTunnelMesh.EntityData.SegmentPath
    autoTunnelMesh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoTunnelMesh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoTunnelMesh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoTunnelMesh.EntityData.Children = types.NewOrderedMap()
    autoTunnelMesh.EntityData.Children.Append("group-ids", types.YChild{"GroupIds", &autoTunnelMesh.GroupIds})
    autoTunnelMesh.EntityData.Leafs = types.NewOrderedMap()
    autoTunnelMesh.EntityData.Leafs.Append("group-all", types.YLeaf{"GroupAll", autoTunnelMesh.GroupAll})

    autoTunnelMesh.EntityData.YListKeys = []string {}

    return &(autoTunnelMesh.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds
// Enable interfaces in specific MPLS TE
// auto-tunnel mesh-groups
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto-mesh group identifier to enable. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId.
    GroupId []*MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId
}

func (groupIds *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds) GetEntityData() *types.CommonEntityData {
    groupIds.EntityData.YFilter = groupIds.YFilter
    groupIds.EntityData.YangName = "group-ids"
    groupIds.EntityData.BundleName = "cisco_ios_xr"
    groupIds.EntityData.ParentYangName = "auto-tunnel-mesh"
    groupIds.EntityData.SegmentPath = "group-ids"
    groupIds.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/traffic-engineering/auto-tunnel-mesh/" + groupIds.EntityData.SegmentPath
    groupIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupIds.EntityData.Children = types.NewOrderedMap()
    groupIds.EntityData.Children.Append("group-id", types.YChild{"GroupId", nil})
    for i := range groupIds.GroupId {
        groupIds.EntityData.Children.Append(types.GetSegmentPath(groupIds.GroupId[i]), types.YChild{"GroupId", groupIds.GroupId[i]})
    }
    groupIds.EntityData.Leafs = types.NewOrderedMap()

    groupIds.EntityData.YListKeys = []string {}

    return &(groupIds.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId
// Auto-mesh group identifier to enable
type MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Mesh group ID. The type is interface{} with range:
    // 0..4294967295.
    MeshGroupId interface{}
}

func (groupId *MplsLdp_DefaultVrf_Afs_Af_TrafficEngineering_AutoTunnelMesh_GroupIds_GroupId) GetEntityData() *types.CommonEntityData {
    groupId.EntityData.YFilter = groupId.YFilter
    groupId.EntityData.YangName = "group-id"
    groupId.EntityData.BundleName = "cisco_ios_xr"
    groupId.EntityData.ParentYangName = "group-ids"
    groupId.EntityData.SegmentPath = "group-id" + types.AddKeyToken(groupId.MeshGroupId, "mesh-group-id")
    groupId.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/traffic-engineering/auto-tunnel-mesh/group-ids/" + groupId.EntityData.SegmentPath
    groupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupId.EntityData.Children = types.NewOrderedMap()
    groupId.EntityData.Leafs = types.NewOrderedMap()
    groupId.EntityData.Leafs.Append("mesh-group-id", types.YLeaf{"MeshGroupId", groupId.MeshGroupId})

    groupId.EntityData.YListKeys = []string {"MeshGroupId"}

    return &(groupId.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Neighbor
// Configuration related to Neighbors
type MplsLdp_DefaultVrf_Afs_Af_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to neighbors using neighbor address.
    Addresses MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses

    // Configuration related to SR policies.
    SegmentRoutingPolicies MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies
}

func (neighbor *MplsLdp_DefaultVrf_Afs_Af_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "af"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("addresses", types.YChild{"Addresses", &neighbor.Addresses})
    neighbor.EntityData.Children.Append("segment-routing-policies", types.YChild{"SegmentRoutingPolicies", &neighbor.SegmentRoutingPolicies})
    neighbor.EntityData.Leafs = types.NewOrderedMap()

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses
// Configuration related to neighbors using
// neighbor address
type MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address based configuration related to a neighbor. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address.
    Address []*MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address
}

func (addresses *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "neighbor"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/neighbor/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address
// IP address based configuration related to a
// neighbor
type MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Establish targeted session with given address. The type is interface{}.
    Targeted interface{}
}

func (address *MplsLdp_DefaultVrf_Afs_Af_Neighbor_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.IpAddress, "ip-address")
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/neighbor/addresses/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", address.IpAddress})
    address.EntityData.Leafs.Append("targeted", types.YLeaf{"Targeted", address.Targeted})

    address.EntityData.YListKeys = []string {"IpAddress"}

    return &(address.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies
// Configuration related to SR policies
type MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name based configuration related to a SR policy. The type is slice of
    // MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies_SegmentRoutingPolicy.
    SegmentRoutingPolicy []*MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies_SegmentRoutingPolicy
}

func (segmentRoutingPolicies *MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies) GetEntityData() *types.CommonEntityData {
    segmentRoutingPolicies.EntityData.YFilter = segmentRoutingPolicies.YFilter
    segmentRoutingPolicies.EntityData.YangName = "segment-routing-policies"
    segmentRoutingPolicies.EntityData.BundleName = "cisco_ios_xr"
    segmentRoutingPolicies.EntityData.ParentYangName = "neighbor"
    segmentRoutingPolicies.EntityData.SegmentPath = "segment-routing-policies"
    segmentRoutingPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/neighbor/" + segmentRoutingPolicies.EntityData.SegmentPath
    segmentRoutingPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRoutingPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRoutingPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRoutingPolicies.EntityData.Children = types.NewOrderedMap()
    segmentRoutingPolicies.EntityData.Children.Append("segment-routing-policy", types.YChild{"SegmentRoutingPolicy", nil})
    for i := range segmentRoutingPolicies.SegmentRoutingPolicy {
        segmentRoutingPolicies.EntityData.Children.Append(types.GetSegmentPath(segmentRoutingPolicies.SegmentRoutingPolicy[i]), types.YChild{"SegmentRoutingPolicy", segmentRoutingPolicies.SegmentRoutingPolicy[i]})
    }
    segmentRoutingPolicies.EntityData.Leafs = types.NewOrderedMap()

    segmentRoutingPolicies.EntityData.YListKeys = []string {}

    return &(segmentRoutingPolicies.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies_SegmentRoutingPolicy
// Name based configuration related to a SR
// policy
type MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies_SegmentRoutingPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. SR Policy Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Establish targeted session with given address. The type is interface{}.
    Targeted interface{}
}

func (segmentRoutingPolicy *MplsLdp_DefaultVrf_Afs_Af_Neighbor_SegmentRoutingPolicies_SegmentRoutingPolicy) GetEntityData() *types.CommonEntityData {
    segmentRoutingPolicy.EntityData.YFilter = segmentRoutingPolicy.YFilter
    segmentRoutingPolicy.EntityData.YangName = "segment-routing-policy"
    segmentRoutingPolicy.EntityData.BundleName = "cisco_ios_xr"
    segmentRoutingPolicy.EntityData.ParentYangName = "segment-routing-policies"
    segmentRoutingPolicy.EntityData.SegmentPath = "segment-routing-policy" + types.AddKeyToken(segmentRoutingPolicy.Name, "name")
    segmentRoutingPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/neighbor/segment-routing-policies/" + segmentRoutingPolicy.EntityData.SegmentPath
    segmentRoutingPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRoutingPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRoutingPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRoutingPolicy.EntityData.Children = types.NewOrderedMap()
    segmentRoutingPolicy.EntityData.Leafs = types.NewOrderedMap()
    segmentRoutingPolicy.EntityData.Leafs.Append("name", types.YLeaf{"Name", segmentRoutingPolicy.Name})
    segmentRoutingPolicy.EntityData.Leafs.Append("targeted", types.YLeaf{"Targeted", segmentRoutingPolicy.Targeted})

    segmentRoutingPolicy.EntityData.YListKeys = []string {"Name"}

    return &(segmentRoutingPolicy.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol
// MPLS LDP configuration for protocol
// redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP configuration for protocol redistribution.
    Bgp MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp
}

func (redistributionProtocol *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol) GetEntityData() *types.CommonEntityData {
    redistributionProtocol.EntityData.YFilter = redistributionProtocol.YFilter
    redistributionProtocol.EntityData.YangName = "redistribution-protocol"
    redistributionProtocol.EntityData.BundleName = "cisco_ios_xr"
    redistributionProtocol.EntityData.ParentYangName = "af"
    redistributionProtocol.EntityData.SegmentPath = "redistribution-protocol"
    redistributionProtocol.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/" + redistributionProtocol.EntityData.SegmentPath
    redistributionProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributionProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributionProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributionProtocol.EntityData.Children = types.NewOrderedMap()
    redistributionProtocol.EntityData.Children.Append("bgp", types.YChild{"Bgp", &redistributionProtocol.Bgp})
    redistributionProtocol.EntityData.Leafs = types.NewOrderedMap()

    redistributionProtocol.EntityData.YListKeys = []string {}

    return &(redistributionProtocol.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp
// MPLS LDP configuration for protocol
// redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP configuration for protocol redistribution.
    As MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As

    // ACL containing list of neighbors for BGP route redistribution.
    AdvertiseTo MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo
}

func (bgp *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "redistribution-protocol"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/redistribution-protocol/" + bgp.EntityData.SegmentPath
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("as", types.YChild{"As", &bgp.As})
    bgp.EntityData.Children.Append("advertise-to", types.YChild{"AdvertiseTo", &bgp.AdvertiseTo})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As
// MPLS LDP configuration for protocol
// redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First half of BGP AS number in XX.YY format.  Mandatory Must be a non-zero
    // value if second half is zero. The type is interface{} with range: 0..65535.
    AsXx interface{}

    // Second half of BGP AS number in XX.YY format. Mandatory Must be a non-zero
    // value if first half is zero. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}
}

func (as *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_As) GetEntityData() *types.CommonEntityData {
    as.EntityData.YFilter = as.YFilter
    as.EntityData.YangName = "as"
    as.EntityData.BundleName = "cisco_ios_xr"
    as.EntityData.ParentYangName = "bgp"
    as.EntityData.SegmentPath = "as"
    as.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/redistribution-protocol/bgp/" + as.EntityData.SegmentPath
    as.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    as.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    as.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    as.EntityData.Children = types.NewOrderedMap()
    as.EntityData.Leafs = types.NewOrderedMap()
    as.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", as.AsXx})
    as.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", as.AsYy})

    as.EntityData.YListKeys = []string {}

    return &(as.EntityData)
}

// MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo
// ACL containing list of neighbors for BGP
// route redistribution
type MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // advertise to peer acl type. The type is MplsLdpAdvertiseBgpAcl.
    Type interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (advertiseTo *MplsLdp_DefaultVrf_Afs_Af_RedistributionProtocol_Bgp_AdvertiseTo) GetEntityData() *types.CommonEntityData {
    advertiseTo.EntityData.YFilter = advertiseTo.YFilter
    advertiseTo.EntityData.YangName = "advertise-to"
    advertiseTo.EntityData.BundleName = "cisco_ios_xr"
    advertiseTo.EntityData.ParentYangName = "bgp"
    advertiseTo.EntityData.SegmentPath = "advertise-to"
    advertiseTo.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/afs/af/redistribution-protocol/bgp/" + advertiseTo.EntityData.SegmentPath
    advertiseTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertiseTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertiseTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertiseTo.EntityData.Children = types.NewOrderedMap()
    advertiseTo.EntityData.Leafs = types.NewOrderedMap()
    advertiseTo.EntityData.Leafs.Append("type", types.YLeaf{"Type", advertiseTo.Type})
    advertiseTo.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", advertiseTo.PeerAclName})

    advertiseTo.EntityData.YListKeys = []string {}

    return &(advertiseTo.EntityData)
}

// MplsLdp_DefaultVrf_Global
// Default VRF Global configuration for MPLS LDP
type MplsLdp_DefaultVrf_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for LDP Router ID (LDP ID). The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // LDP Session parameters.
    Session MplsLdp_DefaultVrf_Global_Session

    // Configuration related to Neighbors.
    Neighbor MplsLdp_DefaultVrf_Global_Neighbor

    // Configuration for per-VRF LDP Graceful Restart parameters.
    GracefulRestart MplsLdp_DefaultVrf_Global_GracefulRestart
}

func (global *MplsLdp_DefaultVrf_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "default-vrf"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("session", types.YChild{"Session", &global.Session})
    global.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &global.Neighbor})
    global.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &global.GracefulRestart})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", global.RouterId})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// MplsLdp_DefaultVrf_Global_Session
// LDP Session parameters
type MplsLdp_DefaultVrf_Global_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Session Protection parameters.
    Protection MplsLdp_DefaultVrf_Global_Session_Protection

    // ACL with the list of neighbors configured for Downstream on Demand.
    DownstreamOnDemand MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand
}

func (session *MplsLdp_DefaultVrf_Global_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "global"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("protection", types.YChild{"Protection", &session.Protection})
    session.EntityData.Children.Append("downstream-on-demand", types.YChild{"DownstreamOnDemand", &session.DownstreamOnDemand})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// MplsLdp_DefaultVrf_Global_Session_Protection
// Configure Session Protection parameters
type MplsLdp_DefaultVrf_Global_Session_Protection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session protection type. The type is MplsLdpSessionProtection.
    ProtectionType interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}

    // Holdup duration. The type is interface{} with range: 30..2147483.
    Duration interface{}
}

func (protection *MplsLdp_DefaultVrf_Global_Session_Protection) GetEntityData() *types.CommonEntityData {
    protection.EntityData.YFilter = protection.YFilter
    protection.EntityData.YangName = "protection"
    protection.EntityData.BundleName = "cisco_ios_xr"
    protection.EntityData.ParentYangName = "session"
    protection.EntityData.SegmentPath = "protection"
    protection.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/session/" + protection.EntityData.SegmentPath
    protection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protection.EntityData.Children = types.NewOrderedMap()
    protection.EntityData.Leafs = types.NewOrderedMap()
    protection.EntityData.Leafs.Append("protection-type", types.YLeaf{"ProtectionType", protection.ProtectionType})
    protection.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", protection.PeerAclName})
    protection.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", protection.Duration})

    protection.EntityData.YListKeys = []string {}

    return &(protection.EntityData)
}

// MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand
// ACL with the list of neighbors configured for
// Downstream on Demand
type MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Downstream on demand type. The type is MplsLdpDownstreamOnDemand.
    Type interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (downstreamOnDemand *MplsLdp_DefaultVrf_Global_Session_DownstreamOnDemand) GetEntityData() *types.CommonEntityData {
    downstreamOnDemand.EntityData.YFilter = downstreamOnDemand.YFilter
    downstreamOnDemand.EntityData.YangName = "downstream-on-demand"
    downstreamOnDemand.EntityData.BundleName = "cisco_ios_xr"
    downstreamOnDemand.EntityData.ParentYangName = "session"
    downstreamOnDemand.EntityData.SegmentPath = "downstream-on-demand"
    downstreamOnDemand.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/session/" + downstreamOnDemand.EntityData.SegmentPath
    downstreamOnDemand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    downstreamOnDemand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    downstreamOnDemand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    downstreamOnDemand.EntityData.Children = types.NewOrderedMap()
    downstreamOnDemand.EntityData.Leafs = types.NewOrderedMap()
    downstreamOnDemand.EntityData.Leafs.Append("type", types.YLeaf{"Type", downstreamOnDemand.Type})
    downstreamOnDemand.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", downstreamOnDemand.PeerAclName})

    downstreamOnDemand.EntityData.YListKeys = []string {}

    return &(downstreamOnDemand.EntityData)
}

// MplsLdp_DefaultVrf_Global_Neighbor
// Configuration related to Neighbors
type MplsLdp_DefaultVrf_Global_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default password for all neigbors. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}

    // Configuration related to Neighbors using LDP Id.
    LdpIds MplsLdp_DefaultVrf_Global_Neighbor_LdpIds

    // Configuration related to neighbor transport.
    DualStack MplsLdp_DefaultVrf_Global_Neighbor_DualStack
}

func (neighbor *MplsLdp_DefaultVrf_Global_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "global"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("ldp-ids", types.YChild{"LdpIds", &neighbor.LdpIds})
    neighbor.EntityData.Children.Append("dual-stack", types.YChild{"DualStack", &neighbor.DualStack})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("password", types.YLeaf{"Password", neighbor.Password})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// MplsLdp_DefaultVrf_Global_Neighbor_LdpIds
// Configuration related to Neighbors using LDP
// Id
type MplsLdp_DefaultVrf_Global_Neighbor_LdpIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP ID based configuration related to a neigbor. The type is slice of
    // MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId.
    LdpId []*MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId
}

func (ldpIds *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds) GetEntityData() *types.CommonEntityData {
    ldpIds.EntityData.YFilter = ldpIds.YFilter
    ldpIds.EntityData.YangName = "ldp-ids"
    ldpIds.EntityData.BundleName = "cisco_ios_xr"
    ldpIds.EntityData.ParentYangName = "neighbor"
    ldpIds.EntityData.SegmentPath = "ldp-ids"
    ldpIds.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/neighbor/" + ldpIds.EntityData.SegmentPath
    ldpIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpIds.EntityData.Children = types.NewOrderedMap()
    ldpIds.EntityData.Children.Append("ldp-id", types.YChild{"LdpId", nil})
    for i := range ldpIds.LdpId {
        ldpIds.EntityData.Children.Append(types.GetSegmentPath(ldpIds.LdpId[i]), types.YChild{"LdpId", ldpIds.LdpId[i]})
    }
    ldpIds.EntityData.Leafs = types.NewOrderedMap()

    ldpIds.EntityData.YListKeys = []string {}

    return &(ldpIds.EntityData)
}

// MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId
// LDP ID based configuration related to a
// neigbor
type MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..0.
    LabelSpaceId interface{}

    // Password for MD5 authentication for this neighbor.
    Password MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password
}

func (ldpId *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId) GetEntityData() *types.CommonEntityData {
    ldpId.EntityData.YFilter = ldpId.YFilter
    ldpId.EntityData.YangName = "ldp-id"
    ldpId.EntityData.BundleName = "cisco_ios_xr"
    ldpId.EntityData.ParentYangName = "ldp-ids"
    ldpId.EntityData.SegmentPath = "ldp-id" + types.AddKeyToken(ldpId.LsrId, "lsr-id") + types.AddKeyToken(ldpId.LabelSpaceId, "label-space-id")
    ldpId.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/neighbor/ldp-ids/" + ldpId.EntityData.SegmentPath
    ldpId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpId.EntityData.Children = types.NewOrderedMap()
    ldpId.EntityData.Children.Append("password", types.YChild{"Password", &ldpId.Password})
    ldpId.EntityData.Leafs = types.NewOrderedMap()
    ldpId.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", ldpId.LsrId})
    ldpId.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", ldpId.LabelSpaceId})

    ldpId.EntityData.YListKeys = []string {"LsrId", "LabelSpaceId"}

    return &(ldpId.EntityData)
}

// MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password
// Password for MD5 authentication for this
// neighbor
type MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Command type for password configuration. The type is MplsLdpNbrPassword.
    CommandType interface{}

    // The neighbor password. The type is string with pattern: b'(!.+)|([^!].+)'.
    Password interface{}
}

func (password *MplsLdp_DefaultVrf_Global_Neighbor_LdpIds_LdpId_Password) GetEntityData() *types.CommonEntityData {
    password.EntityData.YFilter = password.YFilter
    password.EntityData.YangName = "password"
    password.EntityData.BundleName = "cisco_ios_xr"
    password.EntityData.ParentYangName = "ldp-id"
    password.EntityData.SegmentPath = "password"
    password.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/neighbor/ldp-ids/ldp-id/" + password.EntityData.SegmentPath
    password.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    password.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    password.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    password.EntityData.Children = types.NewOrderedMap()
    password.EntityData.Leafs = types.NewOrderedMap()
    password.EntityData.Leafs.Append("command-type", types.YLeaf{"CommandType", password.CommandType})
    password.EntityData.Leafs.Append("password", types.YLeaf{"Password", password.Password})

    password.EntityData.YListKeys = []string {}

    return &(password.EntityData)
}

// MplsLdp_DefaultVrf_Global_Neighbor_DualStack
// Configuration related to neighbor transport
type MplsLdp_DefaultVrf_Global_Neighbor_DualStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration to enable neighbor dual-stack tlv-compliance. The type is
    // interface{}.
    TlvCompliance interface{}

    // Configuration related to neighbor transport.
    TransportConnection MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection
}

func (dualStack *MplsLdp_DefaultVrf_Global_Neighbor_DualStack) GetEntityData() *types.CommonEntityData {
    dualStack.EntityData.YFilter = dualStack.YFilter
    dualStack.EntityData.YangName = "dual-stack"
    dualStack.EntityData.BundleName = "cisco_ios_xr"
    dualStack.EntityData.ParentYangName = "neighbor"
    dualStack.EntityData.SegmentPath = "dual-stack"
    dualStack.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/neighbor/" + dualStack.EntityData.SegmentPath
    dualStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dualStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dualStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dualStack.EntityData.Children = types.NewOrderedMap()
    dualStack.EntityData.Children.Append("transport-connection", types.YChild{"TransportConnection", &dualStack.TransportConnection})
    dualStack.EntityData.Leafs = types.NewOrderedMap()
    dualStack.EntityData.Leafs.Append("tlv-compliance", types.YLeaf{"TlvCompliance", dualStack.TlvCompliance})

    dualStack.EntityData.YListKeys = []string {}

    return &(dualStack.EntityData)
}

// MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection
// Configuration related to neighbor transport
type MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to neighbor dual-stack xport-connection max-wait. The
    // type is interface{} with range: 0..60. Units are second.
    MaxWait interface{}

    // Configuration related to neighbor dual-stack xport-connection preference.
    Prefer MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer
}

func (transportConnection *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection) GetEntityData() *types.CommonEntityData {
    transportConnection.EntityData.YFilter = transportConnection.YFilter
    transportConnection.EntityData.YangName = "transport-connection"
    transportConnection.EntityData.BundleName = "cisco_ios_xr"
    transportConnection.EntityData.ParentYangName = "dual-stack"
    transportConnection.EntityData.SegmentPath = "transport-connection"
    transportConnection.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/neighbor/dual-stack/" + transportConnection.EntityData.SegmentPath
    transportConnection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportConnection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportConnection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportConnection.EntityData.Children = types.NewOrderedMap()
    transportConnection.EntityData.Children.Append("prefer", types.YChild{"Prefer", &transportConnection.Prefer})
    transportConnection.EntityData.Leafs = types.NewOrderedMap()
    transportConnection.EntityData.Leafs.Append("max-wait", types.YLeaf{"MaxWait", transportConnection.MaxWait})

    transportConnection.EntityData.YListKeys = []string {}

    return &(transportConnection.EntityData)
}

// MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer
// Configuration related to neighbor
// dual-stack xport-connection preference
type MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to neighbor dual-stack xport-connection preference
    // ipv4. The type is interface{}.
    Ipv4 interface{}
}

func (prefer *MplsLdp_DefaultVrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetEntityData() *types.CommonEntityData {
    prefer.EntityData.YFilter = prefer.YFilter
    prefer.EntityData.YangName = "prefer"
    prefer.EntityData.BundleName = "cisco_ios_xr"
    prefer.EntityData.ParentYangName = "transport-connection"
    prefer.EntityData.SegmentPath = "prefer"
    prefer.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/neighbor/dual-stack/transport-connection/" + prefer.EntityData.SegmentPath
    prefer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefer.EntityData.Children = types.NewOrderedMap()
    prefer.EntityData.Leafs = types.NewOrderedMap()
    prefer.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", prefer.Ipv4})

    prefer.EntityData.YListKeys = []string {}

    return &(prefer.EntityData)
}

// MplsLdp_DefaultVrf_Global_GracefulRestart
// Configuration for per-VRF LDP Graceful Restart
// parameters
type MplsLdp_DefaultVrf_Global_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure parameters related to GR peer(s) opearating in helper mode.
    HelperPeer MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer
}

func (gracefulRestart *MplsLdp_DefaultVrf_Global_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xr"
    gracefulRestart.EntityData.ParentYangName = "global"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("helper-peer", types.YChild{"HelperPeer", &gracefulRestart.HelperPeer})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer
// Configure parameters related to GR peer(s)
// opearating in helper mode
type MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintain the state of a GR peer upon a local reset. The type is string.
    MaintainOnLocalReset interface{}
}

func (helperPeer *MplsLdp_DefaultVrf_Global_GracefulRestart_HelperPeer) GetEntityData() *types.CommonEntityData {
    helperPeer.EntityData.YFilter = helperPeer.YFilter
    helperPeer.EntityData.YangName = "helper-peer"
    helperPeer.EntityData.BundleName = "cisco_ios_xr"
    helperPeer.EntityData.ParentYangName = "graceful-restart"
    helperPeer.EntityData.SegmentPath = "helper-peer"
    helperPeer.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/global/graceful-restart/" + helperPeer.EntityData.SegmentPath
    helperPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperPeer.EntityData.Children = types.NewOrderedMap()
    helperPeer.EntityData.Leafs = types.NewOrderedMap()
    helperPeer.EntityData.Leafs.Append("maintain-on-local-reset", types.YLeaf{"MaintainOnLocalReset", helperPeer.MaintainOnLocalReset})

    helperPeer.EntityData.YListKeys = []string {}

    return &(helperPeer.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces
// MPLS LDP configuration pertaining to interfaces
type MplsLdp_DefaultVrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP configuration for a particular interface. The type is slice of
    // MplsLdp_DefaultVrf_Interfaces_Interface.
    Interface []*MplsLdp_DefaultVrf_Interfaces_Interface
}

func (interfaces *MplsLdp_DefaultVrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "default-vrf"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface
// MPLS LDP configuration for a particular
// interface
type MplsLdp_DefaultVrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of interface. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable Label Distribution Protocol (LDP) on thisinterface. The type is
    // interface{}.
    Enable interface{}

    // Address Family specific configuration for MPLS LDP intf.
    Afs MplsLdp_DefaultVrf_Interfaces_Interface_Afs

    // Per VRF interface Global configuration for MPLS LDP.
    Global MplsLdp_DefaultVrf_Interfaces_Interface_Global
}

func (self *MplsLdp_DefaultVrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("afs", types.YChild{"Afs", &self.Afs})
    self.EntityData.Children.Append("global", types.YChild{"Global", &self.Global})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs
// Address Family specific configuration for
// MPLS LDP intf
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af.
    Af []*MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af
}

func (afs *MplsLdp_DefaultVrf_Interfaces_Interface_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "interface"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af
// Configure data for given Address Family
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (af *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("discovery", types.YChild{"Discovery", &af.Discovery})
    af.EntityData.Children.Append("igp", types.YChild{"Igp", &af.Igp})
    af.EntityData.Children.Append("mldp", types.YChild{"Mldp", &af.Mldp})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery
// Configure interface discovery parameters
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP configuration for interface discovery transportaddress.
    TransportAddress MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xr"
    discovery.EntityData.ParentYangName = "af"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/afs/af/" + discovery.EntityData.SegmentPath
    discovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Children.Append("transport-address", types.YChild{"TransportAddress", &discovery.TransportAddress})
    discovery.EntityData.Leafs = types.NewOrderedMap()

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
// MPLS LDP configuration for interface
// discovery transportaddress.
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport address option. The type is MplsLdpTransportAddress.
    AddressType interface{}

    // IP address. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (transportAddress *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetEntityData() *types.CommonEntityData {
    transportAddress.EntityData.YFilter = transportAddress.YFilter
    transportAddress.EntityData.YangName = "transport-address"
    transportAddress.EntityData.BundleName = "cisco_ios_xr"
    transportAddress.EntityData.ParentYangName = "discovery"
    transportAddress.EntityData.SegmentPath = "transport-address"
    transportAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/afs/af/discovery/" + transportAddress.EntityData.SegmentPath
    transportAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportAddress.EntityData.Children = types.NewOrderedMap()
    transportAddress.EntityData.Leafs = types.NewOrderedMap()
    transportAddress.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", transportAddress.AddressType})
    transportAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", transportAddress.Address})

    transportAddress.EntityData.YListKeys = []string {}

    return &(transportAddress.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp
// LDP interface IGP configuration
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable IGP Auto-config on this interface. The type is interface{}.
    DisableAutoConfig interface{}
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "af"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/afs/af/" + igp.EntityData.SegmentPath
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Leafs = types.NewOrderedMap()
    igp.EntityData.Leafs.Append("disable-auto-config", types.YLeaf{"DisableAutoConfig", igp.DisableAutoConfig})

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp
// Interface configuration parameters for mLDP
type MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable mLDP on LDP enabled interface. The type is interface{}.
    Disable interface{}
}

func (mldp *MplsLdp_DefaultVrf_Interfaces_Interface_Afs_Af_Mldp) GetEntityData() *types.CommonEntityData {
    mldp.EntityData.YFilter = mldp.YFilter
    mldp.EntityData.YangName = "mldp"
    mldp.EntityData.BundleName = "cisco_ios_xr"
    mldp.EntityData.ParentYangName = "af"
    mldp.EntityData.SegmentPath = "mldp"
    mldp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/afs/af/" + mldp.EntityData.SegmentPath
    mldp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mldp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mldp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mldp.EntityData.Children = types.NewOrderedMap()
    mldp.EntityData.Leafs = types.NewOrderedMap()
    mldp.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", mldp.Disable})

    mldp.EntityData.YListKeys = []string {}

    return &(mldp.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Global
// Per VRF interface Global configuration for
// MPLS LDP
type MplsLdp_DefaultVrf_Interfaces_Interface_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure interface discovery parameters.
    Discovery MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery

    // LDP IGP configuration.
    Igp MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp
}

func (global *MplsLdp_DefaultVrf_Interfaces_Interface_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "interface"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("discovery", types.YChild{"Discovery", &global.Discovery})
    global.EntityData.Children.Append("igp", types.YChild{"Igp", &global.Igp})
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery
// Configure interface discovery parameters
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable discovery's quick start mode. The type is interface{}.
    DisableQuickStart interface{}

    // LDP Link Hellos.
    LinkHello MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello
}

func (discovery *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xr"
    discovery.EntityData.ParentYangName = "global"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/global/" + discovery.EntityData.SegmentPath
    discovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Children.Append("link-hello", types.YChild{"LinkHello", &discovery.LinkHello})
    discovery.EntityData.Leafs = types.NewOrderedMap()
    discovery.EntityData.Leafs.Append("disable-quick-start", types.YLeaf{"DisableQuickStart", discovery.DisableQuickStart})

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello
// LDP Link Hellos
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello struct {
    EntityData types.CommonEntityData
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

func (linkHello *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Discovery_LinkHello) GetEntityData() *types.CommonEntityData {
    linkHello.EntityData.YFilter = linkHello.YFilter
    linkHello.EntityData.YangName = "link-hello"
    linkHello.EntityData.BundleName = "cisco_ios_xr"
    linkHello.EntityData.ParentYangName = "discovery"
    linkHello.EntityData.SegmentPath = "link-hello"
    linkHello.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/global/discovery/" + linkHello.EntityData.SegmentPath
    linkHello.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkHello.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkHello.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkHello.EntityData.Children = types.NewOrderedMap()
    linkHello.EntityData.Leafs = types.NewOrderedMap()
    linkHello.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", linkHello.Interval})
    linkHello.EntityData.Leafs.Append("dual-stack", types.YLeaf{"DualStack", linkHello.DualStack})
    linkHello.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", linkHello.HoldTime})

    linkHello.EntityData.YListKeys = []string {}

    return &(linkHello.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp
// LDP IGP configuration
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP IGP synchronization.
    Sync MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync
}

func (igp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "global"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/global/" + igp.EntityData.SegmentPath
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("sync", types.YChild{"Sync", &igp.Sync})
    igp.EntityData.Leafs = types.NewOrderedMap()

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync
// LDP IGP synchronization
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP IGP synchronization delay time.
    Delay MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay
}

func (sync *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync) GetEntityData() *types.CommonEntityData {
    sync.EntityData.YFilter = sync.YFilter
    sync.EntityData.YangName = "sync"
    sync.EntityData.BundleName = "cisco_ios_xr"
    sync.EntityData.ParentYangName = "igp"
    sync.EntityData.SegmentPath = "sync"
    sync.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/global/igp/" + sync.EntityData.SegmentPath
    sync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sync.EntityData.Children = types.NewOrderedMap()
    sync.EntityData.Children.Append("delay", types.YChild{"Delay", &sync.Delay})
    sync.EntityData.Leafs = types.NewOrderedMap()

    sync.EntityData.YListKeys = []string {}

    return &(sync.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay
// LDP IGP synchronization delay time
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface sync up delay after session up.
    OnSessionUp MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp
}

func (delay *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay) GetEntityData() *types.CommonEntityData {
    delay.EntityData.YFilter = delay.YFilter
    delay.EntityData.YangName = "delay"
    delay.EntityData.BundleName = "cisco_ios_xr"
    delay.EntityData.ParentYangName = "sync"
    delay.EntityData.SegmentPath = "delay"
    delay.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/global/igp/sync/" + delay.EntityData.SegmentPath
    delay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delay.EntityData.Children = types.NewOrderedMap()
    delay.EntityData.Children.Append("on-session-up", types.YChild{"OnSessionUp", &delay.OnSessionUp})
    delay.EntityData.Leafs = types.NewOrderedMap()

    delay.EntityData.YListKeys = []string {}

    return &(delay.EntityData)
}

// MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp
// Interface sync up delay after session up
type MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable delay after session up. The type is interface{}.
    Disable interface{}

    // Time (seconds). The type is interface{} with range: 5..300. Units are
    // second.
    Timeout interface{}
}

func (onSessionUp *MplsLdp_DefaultVrf_Interfaces_Interface_Global_Igp_Sync_Delay_OnSessionUp) GetEntityData() *types.CommonEntityData {
    onSessionUp.EntityData.YFilter = onSessionUp.YFilter
    onSessionUp.EntityData.YangName = "on-session-up"
    onSessionUp.EntityData.BundleName = "cisco_ios_xr"
    onSessionUp.EntityData.ParentYangName = "delay"
    onSessionUp.EntityData.SegmentPath = "on-session-up"
    onSessionUp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/default-vrf/interfaces/interface/global/igp/sync/delay/" + onSessionUp.EntityData.SegmentPath
    onSessionUp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onSessionUp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onSessionUp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onSessionUp.EntityData.Children = types.NewOrderedMap()
    onSessionUp.EntityData.Leafs = types.NewOrderedMap()
    onSessionUp.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", onSessionUp.Disable})
    onSessionUp.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", onSessionUp.Timeout})

    onSessionUp.EntityData.YListKeys = []string {}

    return &(onSessionUp.EntityData)
}

// MplsLdp_Vrfs
// VRF Table attribute configuration for MPLS LDP
type MplsLdp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF attribute configuration for MPLS LDP. The type is slice of
    // MplsLdp_Vrfs_Vrf.
    Vrf []*MplsLdp_Vrfs_Vrf
}

func (vrfs *MplsLdp_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "mpls-ldp"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// MplsLdp_Vrfs_Vrf
// VRF attribute configuration for MPLS LDP
type MplsLdp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (vrf *MplsLdp_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("global", types.YChild{"Global", &vrf.Global})
    vrf.EntityData.Children.Append("afs", types.YChild{"Afs", &vrf.Afs})
    vrf.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &vrf.Interfaces})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrf.Enable})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global
// Per VRF Global configuration for MPLS LDP
type MplsLdp_Vrfs_Vrf_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for LDP Router ID (LDP ID). The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // LDP Session parameters.
    Session MplsLdp_Vrfs_Vrf_Global_Session

    // Configuration related to Neighbors.
    Neighbor MplsLdp_Vrfs_Vrf_Global_Neighbor

    // Configuration for per-VRF LDP Graceful Restart parameters.
    GracefulRestart MplsLdp_Vrfs_Vrf_Global_GracefulRestart
}

func (global *MplsLdp_Vrfs_Vrf_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "vrf"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("session", types.YChild{"Session", &global.Session})
    global.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &global.Neighbor})
    global.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &global.GracefulRestart})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", global.RouterId})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Session
// LDP Session parameters
type MplsLdp_Vrfs_Vrf_Global_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL with the list of neighbors configured for Downstream on Demand.
    DownstreamOnDemand MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand
}

func (session *MplsLdp_Vrfs_Vrf_Global_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "global"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("downstream-on-demand", types.YChild{"DownstreamOnDemand", &session.DownstreamOnDemand})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand
// ACL with the list of neighbors configured
// for Downstream on Demand
type MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Downstream on demand type. The type is MplsLdpDownstreamOnDemand.
    Type interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (downstreamOnDemand *MplsLdp_Vrfs_Vrf_Global_Session_DownstreamOnDemand) GetEntityData() *types.CommonEntityData {
    downstreamOnDemand.EntityData.YFilter = downstreamOnDemand.YFilter
    downstreamOnDemand.EntityData.YangName = "downstream-on-demand"
    downstreamOnDemand.EntityData.BundleName = "cisco_ios_xr"
    downstreamOnDemand.EntityData.ParentYangName = "session"
    downstreamOnDemand.EntityData.SegmentPath = "downstream-on-demand"
    downstreamOnDemand.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/session/" + downstreamOnDemand.EntityData.SegmentPath
    downstreamOnDemand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    downstreamOnDemand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    downstreamOnDemand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    downstreamOnDemand.EntityData.Children = types.NewOrderedMap()
    downstreamOnDemand.EntityData.Leafs = types.NewOrderedMap()
    downstreamOnDemand.EntityData.Leafs.Append("type", types.YLeaf{"Type", downstreamOnDemand.Type})
    downstreamOnDemand.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", downstreamOnDemand.PeerAclName})

    downstreamOnDemand.EntityData.YListKeys = []string {}

    return &(downstreamOnDemand.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Neighbor
// Configuration related to Neighbors
type MplsLdp_Vrfs_Vrf_Global_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default password for all neigbors. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}

    // Configuration related to neighbor transport.
    DualStack MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack

    // Configuration related to Neighbors using LDP Id.
    LdpIds MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds
}

func (neighbor *MplsLdp_Vrfs_Vrf_Global_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "global"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("dual-stack", types.YChild{"DualStack", &neighbor.DualStack})
    neighbor.EntityData.Children.Append("ldp-ids", types.YChild{"LdpIds", &neighbor.LdpIds})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("password", types.YLeaf{"Password", neighbor.Password})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack
// Configuration related to neighbor transport
type MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to neighbor transport.
    TransportConnection MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection
}

func (dualStack *MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack) GetEntityData() *types.CommonEntityData {
    dualStack.EntityData.YFilter = dualStack.YFilter
    dualStack.EntityData.YangName = "dual-stack"
    dualStack.EntityData.BundleName = "cisco_ios_xr"
    dualStack.EntityData.ParentYangName = "neighbor"
    dualStack.EntityData.SegmentPath = "dual-stack"
    dualStack.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/neighbor/" + dualStack.EntityData.SegmentPath
    dualStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dualStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dualStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dualStack.EntityData.Children = types.NewOrderedMap()
    dualStack.EntityData.Children.Append("transport-connection", types.YChild{"TransportConnection", &dualStack.TransportConnection})
    dualStack.EntityData.Leafs = types.NewOrderedMap()

    dualStack.EntityData.YListKeys = []string {}

    return &(dualStack.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection
// Configuration related to neighbor transport
type MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to neighbor dual-stack xport-connection max-wait. The
    // type is interface{} with range: 0..60. Units are second.
    MaxWait interface{}

    // Configuration related to neighbor dual-stack xport-connection preference.
    Prefer MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection_Prefer
}

func (transportConnection *MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection) GetEntityData() *types.CommonEntityData {
    transportConnection.EntityData.YFilter = transportConnection.YFilter
    transportConnection.EntityData.YangName = "transport-connection"
    transportConnection.EntityData.BundleName = "cisco_ios_xr"
    transportConnection.EntityData.ParentYangName = "dual-stack"
    transportConnection.EntityData.SegmentPath = "transport-connection"
    transportConnection.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/neighbor/dual-stack/" + transportConnection.EntityData.SegmentPath
    transportConnection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportConnection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportConnection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportConnection.EntityData.Children = types.NewOrderedMap()
    transportConnection.EntityData.Children.Append("prefer", types.YChild{"Prefer", &transportConnection.Prefer})
    transportConnection.EntityData.Leafs = types.NewOrderedMap()
    transportConnection.EntityData.Leafs.Append("max-wait", types.YLeaf{"MaxWait", transportConnection.MaxWait})

    transportConnection.EntityData.YListKeys = []string {}

    return &(transportConnection.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection_Prefer
// Configuration related to neighbor
// dual-stack xport-connection preference
type MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection_Prefer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to neighbor dual-stack xport-connection preference
    // ipv4. The type is interface{}.
    Ipv4 interface{}
}

func (prefer *MplsLdp_Vrfs_Vrf_Global_Neighbor_DualStack_TransportConnection_Prefer) GetEntityData() *types.CommonEntityData {
    prefer.EntityData.YFilter = prefer.YFilter
    prefer.EntityData.YangName = "prefer"
    prefer.EntityData.BundleName = "cisco_ios_xr"
    prefer.EntityData.ParentYangName = "transport-connection"
    prefer.EntityData.SegmentPath = "prefer"
    prefer.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/neighbor/dual-stack/transport-connection/" + prefer.EntityData.SegmentPath
    prefer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefer.EntityData.Children = types.NewOrderedMap()
    prefer.EntityData.Leafs = types.NewOrderedMap()
    prefer.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", prefer.Ipv4})

    prefer.EntityData.YListKeys = []string {}

    return &(prefer.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds
// Configuration related to Neighbors using LDP
// Id
type MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP ID based configuration related to a neigbor. The type is slice of
    // MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId.
    LdpId []*MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId
}

func (ldpIds *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds) GetEntityData() *types.CommonEntityData {
    ldpIds.EntityData.YFilter = ldpIds.YFilter
    ldpIds.EntityData.YangName = "ldp-ids"
    ldpIds.EntityData.BundleName = "cisco_ios_xr"
    ldpIds.EntityData.ParentYangName = "neighbor"
    ldpIds.EntityData.SegmentPath = "ldp-ids"
    ldpIds.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/neighbor/" + ldpIds.EntityData.SegmentPath
    ldpIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpIds.EntityData.Children = types.NewOrderedMap()
    ldpIds.EntityData.Children.Append("ldp-id", types.YChild{"LdpId", nil})
    for i := range ldpIds.LdpId {
        ldpIds.EntityData.Children.Append(types.GetSegmentPath(ldpIds.LdpId[i]), types.YChild{"LdpId", ldpIds.LdpId[i]})
    }
    ldpIds.EntityData.Leafs = types.NewOrderedMap()

    ldpIds.EntityData.YListKeys = []string {}

    return &(ldpIds.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId
// LDP ID based configuration related to a
// neigbor
type MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LsrId interface{}

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..0.
    LabelSpaceId interface{}

    // Password for MD5 authentication for this neighbor.
    Password MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password
}

func (ldpId *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId) GetEntityData() *types.CommonEntityData {
    ldpId.EntityData.YFilter = ldpId.YFilter
    ldpId.EntityData.YangName = "ldp-id"
    ldpId.EntityData.BundleName = "cisco_ios_xr"
    ldpId.EntityData.ParentYangName = "ldp-ids"
    ldpId.EntityData.SegmentPath = "ldp-id" + types.AddKeyToken(ldpId.LsrId, "lsr-id") + types.AddKeyToken(ldpId.LabelSpaceId, "label-space-id")
    ldpId.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/neighbor/ldp-ids/" + ldpId.EntityData.SegmentPath
    ldpId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpId.EntityData.Children = types.NewOrderedMap()
    ldpId.EntityData.Children.Append("password", types.YChild{"Password", &ldpId.Password})
    ldpId.EntityData.Leafs = types.NewOrderedMap()
    ldpId.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", ldpId.LsrId})
    ldpId.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", ldpId.LabelSpaceId})

    ldpId.EntityData.YListKeys = []string {"LsrId", "LabelSpaceId"}

    return &(ldpId.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password
// Password for MD5 authentication for this
// neighbor
type MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Command type for password configuration. The type is MplsLdpNbrPassword.
    CommandType interface{}

    // The neighbor password. The type is string with pattern: b'(!.+)|([^!].+)'.
    Password interface{}
}

func (password *MplsLdp_Vrfs_Vrf_Global_Neighbor_LdpIds_LdpId_Password) GetEntityData() *types.CommonEntityData {
    password.EntityData.YFilter = password.YFilter
    password.EntityData.YangName = "password"
    password.EntityData.BundleName = "cisco_ios_xr"
    password.EntityData.ParentYangName = "ldp-id"
    password.EntityData.SegmentPath = "password"
    password.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/neighbor/ldp-ids/ldp-id/" + password.EntityData.SegmentPath
    password.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    password.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    password.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    password.EntityData.Children = types.NewOrderedMap()
    password.EntityData.Leafs = types.NewOrderedMap()
    password.EntityData.Leafs.Append("command-type", types.YLeaf{"CommandType", password.CommandType})
    password.EntityData.Leafs.Append("password", types.YLeaf{"Password", password.Password})

    password.EntityData.YListKeys = []string {}

    return &(password.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_GracefulRestart
// Configuration for per-VRF LDP Graceful
// Restart parameters
type MplsLdp_Vrfs_Vrf_Global_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure parameters related to GR peer(s) opearating in helper mode.
    HelperPeer MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer
}

func (gracefulRestart *MplsLdp_Vrfs_Vrf_Global_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xr"
    gracefulRestart.EntityData.ParentYangName = "global"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("helper-peer", types.YChild{"HelperPeer", &gracefulRestart.HelperPeer})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer
// Configure parameters related to GR peer(s)
// opearating in helper mode
type MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintain the state of a GR peer upon a local reset. The type is string.
    MaintainOnLocalReset interface{}
}

func (helperPeer *MplsLdp_Vrfs_Vrf_Global_GracefulRestart_HelperPeer) GetEntityData() *types.CommonEntityData {
    helperPeer.EntityData.YFilter = helperPeer.YFilter
    helperPeer.EntityData.YangName = "helper-peer"
    helperPeer.EntityData.BundleName = "cisco_ios_xr"
    helperPeer.EntityData.ParentYangName = "graceful-restart"
    helperPeer.EntityData.SegmentPath = "helper-peer"
    helperPeer.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/global/graceful-restart/" + helperPeer.EntityData.SegmentPath
    helperPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperPeer.EntityData.Children = types.NewOrderedMap()
    helperPeer.EntityData.Leafs = types.NewOrderedMap()
    helperPeer.EntityData.Leafs.Append("maintain-on-local-reset", types.YLeaf{"MaintainOnLocalReset", helperPeer.MaintainOnLocalReset})

    helperPeer.EntityData.YListKeys = []string {}

    return &(helperPeer.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs
// Address Family specific configuration for MPLS
// LDP vrf
type MplsLdp_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af.
    Af []*MplsLdp_Vrfs_Vrf_Afs_Af
}

func (afs *MplsLdp_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af
// Configure data for given Address Family
type MplsLdp_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family name. The type is MplsLdpafName.
    AfName interface{}

    // Enable Address Family. The type is interface{}.
    Enable interface{}

    // Configure Discovery parameters.
    Discovery MplsLdp_Vrfs_Vrf_Afs_Af_Discovery

    // Configure Label policies and control.
    Label MplsLdp_Vrfs_Vrf_Afs_Af_Label
}

func (af *MplsLdp_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("discovery", types.YChild{"Discovery", &af.Discovery})
    af.EntityData.Children.Append("label", types.YChild{"Label", &af.Label})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Discovery
// Configure Discovery parameters
type MplsLdp_Vrfs_Vrf_Afs_Af_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global discovery transport address for address family. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    TransportAddress interface{}
}

func (discovery *MplsLdp_Vrfs_Vrf_Afs_Af_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xr"
    discovery.EntityData.ParentYangName = "af"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/" + discovery.EntityData.SegmentPath
    discovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Leafs = types.NewOrderedMap()
    discovery.EntityData.Leafs.Append("transport-address", types.YLeaf{"TransportAddress", discovery.TransportAddress})

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label
// Configure Label policies and control
type MplsLdp_Vrfs_Vrf_Afs_Af_Label struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure remote/peer label policies and control.
    Remote MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote

    // Configure local label policies and control.
    Local MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local
}

func (label *MplsLdp_Vrfs_Vrf_Afs_Af_Label) GetEntityData() *types.CommonEntityData {
    label.EntityData.YFilter = label.YFilter
    label.EntityData.YangName = "label"
    label.EntityData.BundleName = "cisco_ios_xr"
    label.EntityData.ParentYangName = "af"
    label.EntityData.SegmentPath = "label"
    label.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/" + label.EntityData.SegmentPath
    label.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    label.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    label.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    label.EntityData.Children = types.NewOrderedMap()
    label.EntityData.Children.Append("remote", types.YChild{"Remote", &label.Remote})
    label.EntityData.Children.Append("local", types.YChild{"Local", &label.Local})
    label.EntityData.Leafs = types.NewOrderedMap()

    label.EntityData.YListKeys = []string {}

    return &(label.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote
// Configure remote/peer label policies and
// control
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure inbound label acceptance.
    Accept MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept
}

func (remote *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote) GetEntityData() *types.CommonEntityData {
    remote.EntityData.YFilter = remote.YFilter
    remote.EntityData.YangName = "remote"
    remote.EntityData.BundleName = "cisco_ios_xr"
    remote.EntityData.ParentYangName = "label"
    remote.EntityData.SegmentPath = "remote"
    remote.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/" + remote.EntityData.SegmentPath
    remote.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remote.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remote.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remote.EntityData.Children = types.NewOrderedMap()
    remote.EntityData.Children.Append("accept", types.YChild{"Accept", &remote.Accept})
    remote.EntityData.Leafs = types.NewOrderedMap()

    remote.EntityData.YListKeys = []string {}

    return &(remote.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept
// Configure inbound label acceptance
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration related to Neighbors for inbound label acceptance.
    PeerAcceptPolicies MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
}

func (accept *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept) GetEntityData() *types.CommonEntityData {
    accept.EntityData.YFilter = accept.YFilter
    accept.EntityData.YangName = "accept"
    accept.EntityData.BundleName = "cisco_ios_xr"
    accept.EntityData.ParentYangName = "remote"
    accept.EntityData.SegmentPath = "accept"
    accept.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/remote/" + accept.EntityData.SegmentPath
    accept.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accept.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accept.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accept.EntityData.Children = types.NewOrderedMap()
    accept.EntityData.Children.Append("peer-accept-policies", types.YChild{"PeerAcceptPolicies", &accept.PeerAcceptPolicies})
    accept.EntityData.Leafs = types.NewOrderedMap()

    accept.EntityData.YListKeys = []string {}

    return &(accept.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies
// Configuration related to Neighbors for
// inbound label acceptance
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control acceptasnce of labels from a neighbor for prefix(es) using ACL. The
    // type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy.
    PeerAcceptPolicy []*MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
}

func (peerAcceptPolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies) GetEntityData() *types.CommonEntityData {
    peerAcceptPolicies.EntityData.YFilter = peerAcceptPolicies.YFilter
    peerAcceptPolicies.EntityData.YangName = "peer-accept-policies"
    peerAcceptPolicies.EntityData.BundleName = "cisco_ios_xr"
    peerAcceptPolicies.EntityData.ParentYangName = "accept"
    peerAcceptPolicies.EntityData.SegmentPath = "peer-accept-policies"
    peerAcceptPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/remote/accept/" + peerAcceptPolicies.EntityData.SegmentPath
    peerAcceptPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAcceptPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAcceptPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAcceptPolicies.EntityData.Children = types.NewOrderedMap()
    peerAcceptPolicies.EntityData.Children.Append("peer-accept-policy", types.YChild{"PeerAcceptPolicy", nil})
    for i := range peerAcceptPolicies.PeerAcceptPolicy {
        peerAcceptPolicies.EntityData.Children.Append(types.GetSegmentPath(peerAcceptPolicies.PeerAcceptPolicy[i]), types.YChild{"PeerAcceptPolicy", peerAcceptPolicies.PeerAcceptPolicy[i]})
    }
    peerAcceptPolicies.EntityData.Leafs = types.NewOrderedMap()

    peerAcceptPolicies.EntityData.YListKeys = []string {}

    return &(peerAcceptPolicies.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy
// Control acceptasnce of labels from a
// neighbor for prefix(es) using ACL
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..0.
    LabelSpaceId interface{}

    // Data container.
    PeerAcceptPolicyData MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData

    // keys: lsr-id. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId.
    LsrId []*MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId
}

func (peerAcceptPolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy) GetEntityData() *types.CommonEntityData {
    peerAcceptPolicy.EntityData.YFilter = peerAcceptPolicy.YFilter
    peerAcceptPolicy.EntityData.YangName = "peer-accept-policy"
    peerAcceptPolicy.EntityData.BundleName = "cisco_ios_xr"
    peerAcceptPolicy.EntityData.ParentYangName = "peer-accept-policies"
    peerAcceptPolicy.EntityData.SegmentPath = "peer-accept-policy" + types.AddKeyToken(peerAcceptPolicy.LabelSpaceId, "label-space-id")
    peerAcceptPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/remote/accept/peer-accept-policies/" + peerAcceptPolicy.EntityData.SegmentPath
    peerAcceptPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAcceptPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAcceptPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAcceptPolicy.EntityData.Children = types.NewOrderedMap()
    peerAcceptPolicy.EntityData.Children.Append("peer-accept-policy-data", types.YChild{"PeerAcceptPolicyData", &peerAcceptPolicy.PeerAcceptPolicyData})
    peerAcceptPolicy.EntityData.Children.Append("lsr-id", types.YChild{"LsrId", nil})
    for i := range peerAcceptPolicy.LsrId {
        peerAcceptPolicy.EntityData.Children.Append(types.GetSegmentPath(peerAcceptPolicy.LsrId[i]), types.YChild{"LsrId", peerAcceptPolicy.LsrId[i]})
    }
    peerAcceptPolicy.EntityData.Leafs = types.NewOrderedMap()
    peerAcceptPolicy.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", peerAcceptPolicy.LabelSpaceId})

    peerAcceptPolicy.EntityData.YListKeys = []string {"LabelSpaceId"}

    return &(peerAcceptPolicy.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData
// Data container.
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAcceptPolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_PeerAcceptPolicyData) GetEntityData() *types.CommonEntityData {
    peerAcceptPolicyData.EntityData.YFilter = peerAcceptPolicyData.YFilter
    peerAcceptPolicyData.EntityData.YangName = "peer-accept-policy-data"
    peerAcceptPolicyData.EntityData.BundleName = "cisco_ios_xr"
    peerAcceptPolicyData.EntityData.ParentYangName = "peer-accept-policy"
    peerAcceptPolicyData.EntityData.SegmentPath = "peer-accept-policy-data"
    peerAcceptPolicyData.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/remote/accept/peer-accept-policies/peer-accept-policy/" + peerAcceptPolicyData.EntityData.SegmentPath
    peerAcceptPolicyData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAcceptPolicyData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAcceptPolicyData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAcceptPolicyData.EntityData.Children = types.NewOrderedMap()
    peerAcceptPolicyData.EntityData.Leafs = types.NewOrderedMap()
    peerAcceptPolicyData.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", peerAcceptPolicyData.PrefixAclName})

    peerAcceptPolicyData.EntityData.YListKeys = []string {}

    return &(peerAcceptPolicyData.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId
// keys: lsr-id
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LsrId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Remote_Accept_PeerAcceptPolicies_PeerAcceptPolicy_LsrId) GetEntityData() *types.CommonEntityData {
    lsrId.EntityData.YFilter = lsrId.YFilter
    lsrId.EntityData.YangName = "lsr-id"
    lsrId.EntityData.BundleName = "cisco_ios_xr"
    lsrId.EntityData.ParentYangName = "peer-accept-policy"
    lsrId.EntityData.SegmentPath = "lsr-id" + types.AddKeyToken(lsrId.LsrId, "lsr-id")
    lsrId.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/remote/accept/peer-accept-policies/peer-accept-policy/" + lsrId.EntityData.SegmentPath
    lsrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsrId.EntityData.Children = types.NewOrderedMap()
    lsrId.EntityData.Leafs = types.NewOrderedMap()
    lsrId.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", lsrId.LsrId})
    lsrId.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", lsrId.PrefixAclName})

    lsrId.EntityData.YListKeys = []string {"LsrId"}

    return &(lsrId.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local
// Configure local label policies and control
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local struct {
    EntityData types.CommonEntityData
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

func (local *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local) GetEntityData() *types.CommonEntityData {
    local.EntityData.YFilter = local.YFilter
    local.EntityData.YangName = "local"
    local.EntityData.BundleName = "cisco_ios_xr"
    local.EntityData.ParentYangName = "label"
    local.EntityData.SegmentPath = "local"
    local.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/" + local.EntityData.SegmentPath
    local.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    local.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    local.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    local.EntityData.Children = types.NewOrderedMap()
    local.EntityData.Children.Append("advertise", types.YChild{"Advertise", &local.Advertise})
    local.EntityData.Children.Append("allocate", types.YChild{"Allocate", &local.Allocate})
    local.EntityData.Leafs = types.NewOrderedMap()
    local.EntityData.Leafs.Append("implicit-null-override", types.YLeaf{"ImplicitNullOverride", local.ImplicitNullOverride})
    local.EntityData.Leafs.Append("default-route", types.YLeaf{"DefaultRoute", local.DefaultRoute})

    local.EntityData.YListKeys = []string {}

    return &(local.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise
// Configure outbound label advertisement
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise struct {
    EntityData types.CommonEntityData
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

func (advertise *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise) GetEntityData() *types.CommonEntityData {
    advertise.EntityData.YFilter = advertise.YFilter
    advertise.EntityData.YangName = "advertise"
    advertise.EntityData.BundleName = "cisco_ios_xr"
    advertise.EntityData.ParentYangName = "local"
    advertise.EntityData.SegmentPath = "advertise"
    advertise.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/" + advertise.EntityData.SegmentPath
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = types.NewOrderedMap()
    advertise.EntityData.Children.Append("peer-advertise-policies", types.YChild{"PeerAdvertisePolicies", &advertise.PeerAdvertisePolicies})
    advertise.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &advertise.Interfaces})
    advertise.EntityData.Children.Append("explicit-null", types.YChild{"ExplicitNull", &advertise.ExplicitNull})
    advertise.EntityData.Leafs = types.NewOrderedMap()
    advertise.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", advertise.Disable})

    advertise.EntityData.YListKeys = []string {}

    return &(advertise.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies
// Configure peer centric outbound label
// advertisement using ACL
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control advertisement of prefix(es) using ACL. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy.
    PeerAdvertisePolicy []*MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
}

func (peerAdvertisePolicies *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies) GetEntityData() *types.CommonEntityData {
    peerAdvertisePolicies.EntityData.YFilter = peerAdvertisePolicies.YFilter
    peerAdvertisePolicies.EntityData.YangName = "peer-advertise-policies"
    peerAdvertisePolicies.EntityData.BundleName = "cisco_ios_xr"
    peerAdvertisePolicies.EntityData.ParentYangName = "advertise"
    peerAdvertisePolicies.EntityData.SegmentPath = "peer-advertise-policies"
    peerAdvertisePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/advertise/" + peerAdvertisePolicies.EntityData.SegmentPath
    peerAdvertisePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAdvertisePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAdvertisePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAdvertisePolicies.EntityData.Children = types.NewOrderedMap()
    peerAdvertisePolicies.EntityData.Children.Append("peer-advertise-policy", types.YChild{"PeerAdvertisePolicy", nil})
    for i := range peerAdvertisePolicies.PeerAdvertisePolicy {
        peerAdvertisePolicies.EntityData.Children.Append(types.GetSegmentPath(peerAdvertisePolicies.PeerAdvertisePolicy[i]), types.YChild{"PeerAdvertisePolicy", peerAdvertisePolicies.PeerAdvertisePolicy[i]})
    }
    peerAdvertisePolicies.EntityData.Leafs = types.NewOrderedMap()

    peerAdvertisePolicies.EntityData.YListKeys = []string {}

    return &(peerAdvertisePolicies.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy
// Control advertisement of prefix(es)
// using ACL
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Label space ID of neighbor. The type is
    // interface{} with range: 0..0.
    LabelSpaceId interface{}

    // Data container.
    PeerAdvertisePolicyData MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData

    // keys: lsr-id. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId.
    LsrId []*MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId
}

func (peerAdvertisePolicy *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy) GetEntityData() *types.CommonEntityData {
    peerAdvertisePolicy.EntityData.YFilter = peerAdvertisePolicy.YFilter
    peerAdvertisePolicy.EntityData.YangName = "peer-advertise-policy"
    peerAdvertisePolicy.EntityData.BundleName = "cisco_ios_xr"
    peerAdvertisePolicy.EntityData.ParentYangName = "peer-advertise-policies"
    peerAdvertisePolicy.EntityData.SegmentPath = "peer-advertise-policy" + types.AddKeyToken(peerAdvertisePolicy.LabelSpaceId, "label-space-id")
    peerAdvertisePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/advertise/peer-advertise-policies/" + peerAdvertisePolicy.EntityData.SegmentPath
    peerAdvertisePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAdvertisePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAdvertisePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAdvertisePolicy.EntityData.Children = types.NewOrderedMap()
    peerAdvertisePolicy.EntityData.Children.Append("peer-advertise-policy-data", types.YChild{"PeerAdvertisePolicyData", &peerAdvertisePolicy.PeerAdvertisePolicyData})
    peerAdvertisePolicy.EntityData.Children.Append("lsr-id", types.YChild{"LsrId", nil})
    for i := range peerAdvertisePolicy.LsrId {
        peerAdvertisePolicy.EntityData.Children.Append(types.GetSegmentPath(peerAdvertisePolicy.LsrId[i]), types.YChild{"LsrId", peerAdvertisePolicy.LsrId[i]})
    }
    peerAdvertisePolicy.EntityData.Leafs = types.NewOrderedMap()
    peerAdvertisePolicy.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", peerAdvertisePolicy.LabelSpaceId})

    peerAdvertisePolicy.EntityData.YListKeys = []string {"LabelSpaceId"}

    return &(peerAdvertisePolicy.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData
// Data container.
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (peerAdvertisePolicyData *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_PeerAdvertisePolicyData) GetEntityData() *types.CommonEntityData {
    peerAdvertisePolicyData.EntityData.YFilter = peerAdvertisePolicyData.YFilter
    peerAdvertisePolicyData.EntityData.YangName = "peer-advertise-policy-data"
    peerAdvertisePolicyData.EntityData.BundleName = "cisco_ios_xr"
    peerAdvertisePolicyData.EntityData.ParentYangName = "peer-advertise-policy"
    peerAdvertisePolicyData.EntityData.SegmentPath = "peer-advertise-policy-data"
    peerAdvertisePolicyData.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/advertise/peer-advertise-policies/peer-advertise-policy/" + peerAdvertisePolicyData.EntityData.SegmentPath
    peerAdvertisePolicyData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAdvertisePolicyData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAdvertisePolicyData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAdvertisePolicyData.EntityData.Children = types.NewOrderedMap()
    peerAdvertisePolicyData.EntityData.Leafs = types.NewOrderedMap()
    peerAdvertisePolicyData.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", peerAdvertisePolicyData.PrefixAclName})

    peerAdvertisePolicyData.EntityData.YListKeys = []string {}

    return &(peerAdvertisePolicyData.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId
// keys: lsr-id
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSR ID of neighbor. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LsrId interface{}

    // Name of prefix ACL. The type is string. This attribute is mandatory.
    PrefixAclName interface{}
}

func (lsrId *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_PeerAdvertisePolicies_PeerAdvertisePolicy_LsrId) GetEntityData() *types.CommonEntityData {
    lsrId.EntityData.YFilter = lsrId.YFilter
    lsrId.EntityData.YangName = "lsr-id"
    lsrId.EntityData.BundleName = "cisco_ios_xr"
    lsrId.EntityData.ParentYangName = "peer-advertise-policy"
    lsrId.EntityData.SegmentPath = "lsr-id" + types.AddKeyToken(lsrId.LsrId, "lsr-id")
    lsrId.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/advertise/peer-advertise-policies/peer-advertise-policy/" + lsrId.EntityData.SegmentPath
    lsrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsrId.EntityData.Children = types.NewOrderedMap()
    lsrId.EntityData.Leafs = types.NewOrderedMap()
    lsrId.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", lsrId.LsrId})
    lsrId.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", lsrId.PrefixAclName})

    lsrId.EntityData.YListKeys = []string {"LsrId"}

    return &(lsrId.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces
// Configure outbound label advertisement
// for an interface
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control advertisement of interface's host IP address. The type is slice of
    // MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface.
    Interface []*MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
}

func (interfaces *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "advertise"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/advertise/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface
// Control advertisement of interface's
// host IP address
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of interface. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (self *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/advertise/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull
// Configure advertisment of explicit-null
// for connected prefixes.
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Explicit Null command variant. The type is MplsLdpExpNull.
    ExplicitNullType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}

    // Name of peer ACL. The type is string.
    PeerAclName interface{}
}

func (explicitNull *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Advertise_ExplicitNull) GetEntityData() *types.CommonEntityData {
    explicitNull.EntityData.YFilter = explicitNull.YFilter
    explicitNull.EntityData.YangName = "explicit-null"
    explicitNull.EntityData.BundleName = "cisco_ios_xr"
    explicitNull.EntityData.ParentYangName = "advertise"
    explicitNull.EntityData.SegmentPath = "explicit-null"
    explicitNull.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/advertise/" + explicitNull.EntityData.SegmentPath
    explicitNull.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitNull.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitNull.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitNull.EntityData.Children = types.NewOrderedMap()
    explicitNull.EntityData.Leafs = types.NewOrderedMap()
    explicitNull.EntityData.Leafs.Append("explicit-null-type", types.YLeaf{"ExplicitNullType", explicitNull.ExplicitNullType})
    explicitNull.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", explicitNull.PrefixAclName})
    explicitNull.EntityData.Leafs.Append("peer-acl-name", types.YLeaf{"PeerAclName", explicitNull.PeerAclName})

    explicitNull.EntityData.YListKeys = []string {}

    return &(explicitNull.EntityData)
}

// MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate
// Control local label allocation for
// prefix(es)
type MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label allocation type. The type is MplsLdpLabelAllocation.
    AllocationType interface{}

    // Name of prefix ACL. The type is string.
    PrefixAclName interface{}
}

func (allocate *MplsLdp_Vrfs_Vrf_Afs_Af_Label_Local_Allocate) GetEntityData() *types.CommonEntityData {
    allocate.EntityData.YFilter = allocate.YFilter
    allocate.EntityData.YangName = "allocate"
    allocate.EntityData.BundleName = "cisco_ios_xr"
    allocate.EntityData.ParentYangName = "local"
    allocate.EntityData.SegmentPath = "allocate"
    allocate.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/afs/af/label/local/" + allocate.EntityData.SegmentPath
    allocate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allocate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allocate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allocate.EntityData.Children = types.NewOrderedMap()
    allocate.EntityData.Leafs = types.NewOrderedMap()
    allocate.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", allocate.AllocationType})
    allocate.EntityData.Leafs.Append("prefix-acl-name", types.YLeaf{"PrefixAclName", allocate.PrefixAclName})

    allocate.EntityData.YListKeys = []string {}

    return &(allocate.EntityData)
}

// MplsLdp_Vrfs_Vrf_Interfaces
// MPLS LDP configuration pertaining to
// interfaces
type MplsLdp_Vrfs_Vrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP configuration for a particular interface. The type is slice of
    // MplsLdp_Vrfs_Vrf_Interfaces_Interface.
    Interface []*MplsLdp_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *MplsLdp_Vrfs_Vrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrf"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MplsLdp_Vrfs_Vrf_Interfaces_Interface
// MPLS LDP configuration for a particular
// interface
type MplsLdp_Vrfs_Vrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of interface. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable Label Distribution Protocol (LDP) on thisinterface. The type is
    // interface{}.
    Enable interface{}

    // Address Family specific configuration for MPLS LDP vrf intf.
    Afs MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs
}

func (self *MplsLdp_Vrfs_Vrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("afs", types.YChild{"Afs", &self.Afs})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs
// Address Family specific configuration for
// MPLS LDP vrf intf
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure data for given Address Family. The type is slice of
    // MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af.
    Af []*MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af
}

func (afs *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "interface"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/interfaces/interface/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af
// Configure data for given Address Family
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family name. The type is MplsLdpafName.
    AfName interface{}

    // Enable Address Family. The type is interface{}.
    Enable interface{}

    // Configure interface discovery parameters.
    Discovery MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery
}

func (af *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/interfaces/interface/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("discovery", types.YChild{"Discovery", &af.Discovery})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery
// Configure interface discovery parameters
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP configuration for interface discovery transportaddress.
    TransportAddress MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
}

func (discovery *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xr"
    discovery.EntityData.ParentYangName = "af"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/interfaces/interface/afs/af/" + discovery.EntityData.SegmentPath
    discovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Children.Append("transport-address", types.YChild{"TransportAddress", &discovery.TransportAddress})
    discovery.EntityData.Leafs = types.NewOrderedMap()

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress
// MPLS LDP configuration for interface
// discovery transportaddress.
type MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport address option. The type is MplsLdpTransportAddress.
    AddressType interface{}

    // IP address. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (transportAddress *MplsLdp_Vrfs_Vrf_Interfaces_Interface_Afs_Af_Discovery_TransportAddress) GetEntityData() *types.CommonEntityData {
    transportAddress.EntityData.YFilter = transportAddress.YFilter
    transportAddress.EntityData.YangName = "transport-address"
    transportAddress.EntityData.BundleName = "cisco_ios_xr"
    transportAddress.EntityData.ParentYangName = "discovery"
    transportAddress.EntityData.SegmentPath = "transport-address"
    transportAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/vrfs/vrf/interfaces/interface/afs/af/discovery/" + transportAddress.EntityData.SegmentPath
    transportAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportAddress.EntityData.Children = types.NewOrderedMap()
    transportAddress.EntityData.Leafs = types.NewOrderedMap()
    transportAddress.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", transportAddress.AddressType})
    transportAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", transportAddress.Address})

    transportAddress.EntityData.YListKeys = []string {}

    return &(transportAddress.EntityData)
}

// MplsLdp_Global
// Global configuration for MPLS LDP
type MplsLdp_Global struct {
    EntityData types.CommonEntityData
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

func (global *MplsLdp_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "mpls-ldp"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("entropy-label", types.YChild{"EntropyLabel", &global.EntropyLabel})
    global.EntityData.Children.Append("session", types.YChild{"Session", &global.Session})
    global.EntityData.Children.Append("igp", types.YChild{"Igp", &global.Igp})
    global.EntityData.Children.Append("enable-logging", types.YChild{"EnableLogging", &global.EnableLogging})
    global.EntityData.Children.Append("signalling", types.YChild{"Signalling", &global.Signalling})
    global.EntityData.Children.Append("nsr", types.YChild{"Nsr", &global.Nsr})
    global.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &global.GracefulRestart})
    global.EntityData.Children.Append("discovery", types.YChild{"Discovery", &global.Discovery})
    global.EntityData.Children.Append("mldp", types.YChild{"Mldp", &global.Mldp})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("disable-implicit-ipv4", types.YLeaf{"DisableImplicitIpv4", global.DisableImplicitIpv4})
    global.EntityData.Leafs.Append("ltrace-buf-multiplier", types.YLeaf{"LtraceBufMultiplier", global.LtraceBufMultiplier})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// MplsLdp_Global_EntropyLabel
// Configure for LDP Entropy-Label
type MplsLdp_Global_EntropyLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is interface{}.
    Enable interface{}
}

func (entropyLabel *MplsLdp_Global_EntropyLabel) GetEntityData() *types.CommonEntityData {
    entropyLabel.EntityData.YFilter = entropyLabel.YFilter
    entropyLabel.EntityData.YangName = "entropy-label"
    entropyLabel.EntityData.BundleName = "cisco_ios_xr"
    entropyLabel.EntityData.ParentYangName = "global"
    entropyLabel.EntityData.SegmentPath = "entropy-label"
    entropyLabel.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + entropyLabel.EntityData.SegmentPath
    entropyLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entropyLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entropyLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entropyLabel.EntityData.Children = types.NewOrderedMap()
    entropyLabel.EntityData.Leafs = types.NewOrderedMap()
    entropyLabel.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", entropyLabel.Enable})

    entropyLabel.EntityData.YListKeys = []string {}

    return &(entropyLabel.EntityData)
}

// MplsLdp_Global_Session
// LDP Session parameters
type MplsLdp_Global_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP Session holdtime. The type is interface{} with range: 15..65535. Units
    // are second. The default value is 180.
    HoldTime interface{}

    // Configure Session Backoff parameters.
    BackoffTime MplsLdp_Global_Session_BackoffTime
}

func (session *MplsLdp_Global_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "global"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("backoff-time", types.YChild{"BackoffTime", &session.BackoffTime})
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", session.HoldTime})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// MplsLdp_Global_Session_BackoffTime
// Configure Session Backoff parameters
type MplsLdp_Global_Session_BackoffTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initial session backoff time (seconds). The type is interface{} with range:
    // 5..2147483. Units are second. The default value is 15.
    InitialBackoffTime interface{}

    // Maximum session backoff time (seconds). The type is interface{} with range:
    // 5..2147483. Units are second. The default value is 120.
    MaxBackoffTime interface{}
}

func (backoffTime *MplsLdp_Global_Session_BackoffTime) GetEntityData() *types.CommonEntityData {
    backoffTime.EntityData.YFilter = backoffTime.YFilter
    backoffTime.EntityData.YangName = "backoff-time"
    backoffTime.EntityData.BundleName = "cisco_ios_xr"
    backoffTime.EntityData.ParentYangName = "session"
    backoffTime.EntityData.SegmentPath = "backoff-time"
    backoffTime.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/session/" + backoffTime.EntityData.SegmentPath
    backoffTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backoffTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backoffTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backoffTime.EntityData.Children = types.NewOrderedMap()
    backoffTime.EntityData.Leafs = types.NewOrderedMap()
    backoffTime.EntityData.Leafs.Append("initial-backoff-time", types.YLeaf{"InitialBackoffTime", backoffTime.InitialBackoffTime})
    backoffTime.EntityData.Leafs.Append("max-backoff-time", types.YLeaf{"MaxBackoffTime", backoffTime.MaxBackoffTime})

    backoffTime.EntityData.YListKeys = []string {}

    return &(backoffTime.EntityData)
}

// MplsLdp_Global_Igp
// LDP IGP configuration
type MplsLdp_Global_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP IGP synchronization.
    Sync MplsLdp_Global_Igp_Sync
}

func (igp *MplsLdp_Global_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "global"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + igp.EntityData.SegmentPath
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("sync", types.YChild{"Sync", &igp.Sync})
    igp.EntityData.Leafs = types.NewOrderedMap()

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// MplsLdp_Global_Igp_Sync
// LDP IGP synchronization
type MplsLdp_Global_Igp_Sync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP IGP synchronization delay time.
    Delay MplsLdp_Global_Igp_Sync_Delay
}

func (sync *MplsLdp_Global_Igp_Sync) GetEntityData() *types.CommonEntityData {
    sync.EntityData.YFilter = sync.YFilter
    sync.EntityData.YangName = "sync"
    sync.EntityData.BundleName = "cisco_ios_xr"
    sync.EntityData.ParentYangName = "igp"
    sync.EntityData.SegmentPath = "sync"
    sync.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/igp/" + sync.EntityData.SegmentPath
    sync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sync.EntityData.Children = types.NewOrderedMap()
    sync.EntityData.Children.Append("delay", types.YChild{"Delay", &sync.Delay})
    sync.EntityData.Leafs = types.NewOrderedMap()

    sync.EntityData.YListKeys = []string {}

    return &(sync.EntityData)
}

// MplsLdp_Global_Igp_Sync_Delay
// LDP IGP synchronization delay time
type MplsLdp_Global_Igp_Sync_Delay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface sync up delay after session up. The type is interface{} with
    // range: 5..300. Units are second.
    OnSessionUp interface{}

    // Global sync up delay to be used after process restart. The type is
    // interface{} with range: 60..600. Units are second.
    OnProcRestart interface{}
}

func (delay *MplsLdp_Global_Igp_Sync_Delay) GetEntityData() *types.CommonEntityData {
    delay.EntityData.YFilter = delay.YFilter
    delay.EntityData.YangName = "delay"
    delay.EntityData.BundleName = "cisco_ios_xr"
    delay.EntityData.ParentYangName = "sync"
    delay.EntityData.SegmentPath = "delay"
    delay.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/igp/sync/" + delay.EntityData.SegmentPath
    delay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delay.EntityData.Children = types.NewOrderedMap()
    delay.EntityData.Leafs = types.NewOrderedMap()
    delay.EntityData.Leafs.Append("on-session-up", types.YLeaf{"OnSessionUp", delay.OnSessionUp})
    delay.EntityData.Leafs.Append("on-proc-restart", types.YLeaf{"OnProcRestart", delay.OnProcRestart})

    delay.EntityData.YListKeys = []string {}

    return &(delay.EntityData)
}

// MplsLdp_Global_EnableLogging
// Enable logging of events
type MplsLdp_Global_EnableLogging struct {
    EntityData types.CommonEntityData
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

func (enableLogging *MplsLdp_Global_EnableLogging) GetEntityData() *types.CommonEntityData {
    enableLogging.EntityData.YFilter = enableLogging.YFilter
    enableLogging.EntityData.YangName = "enable-logging"
    enableLogging.EntityData.BundleName = "cisco_ios_xr"
    enableLogging.EntityData.ParentYangName = "global"
    enableLogging.EntityData.SegmentPath = "enable-logging"
    enableLogging.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + enableLogging.EntityData.SegmentPath
    enableLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enableLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enableLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enableLogging.EntityData.Children = types.NewOrderedMap()
    enableLogging.EntityData.Leafs = types.NewOrderedMap()
    enableLogging.EntityData.Leafs.Append("nsr", types.YLeaf{"Nsr", enableLogging.Nsr})
    enableLogging.EntityData.Leafs.Append("neighbor-changes", types.YLeaf{"NeighborChanges", enableLogging.NeighborChanges})
    enableLogging.EntityData.Leafs.Append("adjacency", types.YLeaf{"Adjacency", enableLogging.Adjacency})
    enableLogging.EntityData.Leafs.Append("session-protection", types.YLeaf{"SessionProtection", enableLogging.SessionProtection})
    enableLogging.EntityData.Leafs.Append("gr-session-changes", types.YLeaf{"GrSessionChanges", enableLogging.GrSessionChanges})

    enableLogging.EntityData.YListKeys = []string {}

    return &(enableLogging.EntityData)
}

// MplsLdp_Global_Signalling
// Configure LDP signalling parameters
type MplsLdp_Global_Signalling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DSCP for control packets. The type is interface{} with range: 0..63. The
    // default value is 48.
    Dscp interface{}
}

func (signalling *MplsLdp_Global_Signalling) GetEntityData() *types.CommonEntityData {
    signalling.EntityData.YFilter = signalling.YFilter
    signalling.EntityData.YangName = "signalling"
    signalling.EntityData.BundleName = "cisco_ios_xr"
    signalling.EntityData.ParentYangName = "global"
    signalling.EntityData.SegmentPath = "signalling"
    signalling.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + signalling.EntityData.SegmentPath
    signalling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalling.EntityData.Children = types.NewOrderedMap()
    signalling.EntityData.Leafs = types.NewOrderedMap()
    signalling.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", signalling.Dscp})

    signalling.EntityData.YListKeys = []string {}

    return &(signalling.EntityData)
}

// MplsLdp_Global_Nsr
// Configure LDP Non-Stop Routing
type MplsLdp_Global_Nsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is interface{}.
    Enable interface{}
}

func (nsr *MplsLdp_Global_Nsr) GetEntityData() *types.CommonEntityData {
    nsr.EntityData.YFilter = nsr.YFilter
    nsr.EntityData.YangName = "nsr"
    nsr.EntityData.BundleName = "cisco_ios_xr"
    nsr.EntityData.ParentYangName = "global"
    nsr.EntityData.SegmentPath = "nsr"
    nsr.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + nsr.EntityData.SegmentPath
    nsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nsr.EntityData.Children = types.NewOrderedMap()
    nsr.EntityData.Leafs = types.NewOrderedMap()
    nsr.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", nsr.Enable})

    nsr.EntityData.YListKeys = []string {}

    return &(nsr.EntityData)
}

// MplsLdp_Global_GracefulRestart
// Configuration for LDP Graceful Restart
// parameters
type MplsLdp_Global_GracefulRestart struct {
    EntityData types.CommonEntityData
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

func (gracefulRestart *MplsLdp_Global_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xr"
    gracefulRestart.EntityData.ParentYangName = "global"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs.Append("reconnect-timeout", types.YLeaf{"ReconnectTimeout", gracefulRestart.ReconnectTimeout})
    gracefulRestart.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", gracefulRestart.Enable})
    gracefulRestart.EntityData.Leafs.Append("forwarding-hold-time", types.YLeaf{"ForwardingHoldTime", gracefulRestart.ForwardingHoldTime})

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// MplsLdp_Global_Discovery
// Configure Discovery parameters
type MplsLdp_Global_Discovery struct {
    EntityData types.CommonEntityData
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

func (discovery *MplsLdp_Global_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xr"
    discovery.EntityData.ParentYangName = "global"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + discovery.EntityData.SegmentPath
    discovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Children.Append("link-hello", types.YChild{"LinkHello", &discovery.LinkHello})
    discovery.EntityData.Children.Append("targeted-hello", types.YChild{"TargetedHello", &discovery.TargetedHello})
    discovery.EntityData.Leafs = types.NewOrderedMap()
    discovery.EntityData.Leafs.Append("disable-instance-tlv", types.YLeaf{"DisableInstanceTlv", discovery.DisableInstanceTlv})
    discovery.EntityData.Leafs.Append("disable-quick-start", types.YLeaf{"DisableQuickStart", discovery.DisableQuickStart})

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_Global_Discovery_LinkHello
// LDP Link Hellos
type MplsLdp_Global_Discovery_LinkHello struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Hello interval. The type is interface{} with range: 1..65535. Units
    // are second. The default value is 5.
    Interval interface{}

    // Time (seconds) - 65535 implies infinite. The type is interface{} with
    // range: 1..65535. Units are second. The default value is 15.
    HoldTime interface{}
}

func (linkHello *MplsLdp_Global_Discovery_LinkHello) GetEntityData() *types.CommonEntityData {
    linkHello.EntityData.YFilter = linkHello.YFilter
    linkHello.EntityData.YangName = "link-hello"
    linkHello.EntityData.BundleName = "cisco_ios_xr"
    linkHello.EntityData.ParentYangName = "discovery"
    linkHello.EntityData.SegmentPath = "link-hello"
    linkHello.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/discovery/" + linkHello.EntityData.SegmentPath
    linkHello.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkHello.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkHello.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkHello.EntityData.Children = types.NewOrderedMap()
    linkHello.EntityData.Leafs = types.NewOrderedMap()
    linkHello.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", linkHello.Interval})
    linkHello.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", linkHello.HoldTime})

    linkHello.EntityData.YListKeys = []string {}

    return &(linkHello.EntityData)
}

// MplsLdp_Global_Discovery_TargetedHello
// LDP Targeted Hellos
type MplsLdp_Global_Discovery_TargetedHello struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Targeted Hello interval. The type is interface{} with range: 1..65535.
    // Units are second. The default value is 10.
    Interval interface{}

    // Time (seconds) - 65535 implies infinite. The type is interface{} with
    // range: 1..65535. Units are second. The default value is 90.
    HoldTime interface{}
}

func (targetedHello *MplsLdp_Global_Discovery_TargetedHello) GetEntityData() *types.CommonEntityData {
    targetedHello.EntityData.YFilter = targetedHello.YFilter
    targetedHello.EntityData.YangName = "targeted-hello"
    targetedHello.EntityData.BundleName = "cisco_ios_xr"
    targetedHello.EntityData.ParentYangName = "discovery"
    targetedHello.EntityData.SegmentPath = "targeted-hello"
    targetedHello.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/discovery/" + targetedHello.EntityData.SegmentPath
    targetedHello.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetedHello.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetedHello.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetedHello.EntityData.Children = types.NewOrderedMap()
    targetedHello.EntityData.Leafs = types.NewOrderedMap()
    targetedHello.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", targetedHello.Interval})
    targetedHello.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", targetedHello.HoldTime})

    targetedHello.EntityData.YListKeys = []string {}

    return &(targetedHello.EntityData)
}

// MplsLdp_Global_Mldp
// MPLS mLDP configuration
type MplsLdp_Global_Mldp struct {
    EntityData types.CommonEntityData
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

func (mldp *MplsLdp_Global_Mldp) GetEntityData() *types.CommonEntityData {
    mldp.EntityData.YFilter = mldp.YFilter
    mldp.EntityData.YangName = "mldp"
    mldp.EntityData.BundleName = "cisco_ios_xr"
    mldp.EntityData.ParentYangName = "global"
    mldp.EntityData.SegmentPath = "mldp"
    mldp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/" + mldp.EntityData.SegmentPath
    mldp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mldp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mldp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mldp.EntityData.Children = types.NewOrderedMap()
    mldp.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &mldp.Vrfs})
    mldp.EntityData.Children.Append("default-vrf", types.YChild{"DefaultVrf", &mldp.DefaultVrf})
    mldp.EntityData.Children.Append("mldp-global", types.YChild{"MldpGlobal", &mldp.MldpGlobal})
    mldp.EntityData.Leafs = types.NewOrderedMap()
    mldp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mldp.Enable})

    mldp.EntityData.YListKeys = []string {}

    return &(mldp.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs
// VRF Table attribute configuration for MPLS LDP
type MplsLdp_Global_Mldp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF attribute configuration for MPLS LDP. The type is slice of
    // MplsLdp_Global_Mldp_Vrfs_Vrf.
    Vrf []*MplsLdp_Global_Mldp_Vrfs_Vrf
}

func (vrfs *MplsLdp_Global_Mldp_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "mldp"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf
// VRF attribute configuration for MPLS LDP
type MplsLdp_Global_Mldp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with length: 1..32.
    VrfName interface{}

    // Enable Multicast Label Distribution Protocol (mLDP). The type is
    // interface{}.
    Enable interface{}

    // Address Family specific operational data.
    Afs MplsLdp_Global_Mldp_Vrfs_Vrf_Afs
}

func (vrf *MplsLdp_Global_Mldp_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("afs", types.YChild{"Afs", &vrf.Afs})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrf.Enable})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs
// Address Family specific operational data
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af.
    Af []*MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af
}

func (afs *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af
// Operational data for given Address Family
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (af *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("recursive-forwarding", types.YChild{"RecursiveForwarding", &af.RecursiveForwarding})
    af.EntityData.Children.Append("mldp-recursive-fec", types.YChild{"MldpRecursiveFec", &af.MldpRecursiveFec})
    af.EntityData.Children.Append("neighbor-policies", types.YChild{"NeighborPolicies", &af.NeighborPolicies})
    af.EntityData.Children.Append("mo-frr", types.YChild{"MoFrr", &af.MoFrr})
    af.EntityData.Children.Append("make-before-break", types.YChild{"MakeBeforeBreak", &af.MakeBeforeBreak})
    af.EntityData.Children.Append("csc", types.YChild{"Csc", &af.Csc})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})
    af.EntityData.Leafs.Append("mldp-rib-unicast-always", types.YLeaf{"MldpRibUnicastAlways", af.MldpRibUnicastAlways})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding
// Enable recursive forwarding
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable recursive forwarding. The type is interface{}.
    Enable interface{}

    // Recursive forwarding policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (recursiveForwarding *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_RecursiveForwarding) GetEntityData() *types.CommonEntityData {
    recursiveForwarding.EntityData.YFilter = recursiveForwarding.YFilter
    recursiveForwarding.EntityData.YangName = "recursive-forwarding"
    recursiveForwarding.EntityData.BundleName = "cisco_ios_xr"
    recursiveForwarding.EntityData.ParentYangName = "af"
    recursiveForwarding.EntityData.SegmentPath = "recursive-forwarding"
    recursiveForwarding.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/" + recursiveForwarding.EntityData.SegmentPath
    recursiveForwarding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    recursiveForwarding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    recursiveForwarding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    recursiveForwarding.EntityData.Children = types.NewOrderedMap()
    recursiveForwarding.EntityData.Leafs = types.NewOrderedMap()
    recursiveForwarding.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", recursiveForwarding.Enable})
    recursiveForwarding.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", recursiveForwarding.Policy})

    recursiveForwarding.EntityData.YListKeys = []string {}

    return &(recursiveForwarding.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec
// MPLS mLDP Recursive FEC
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS mLDP Recursive FEC. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MldpRecursiveFec) GetEntityData() *types.CommonEntityData {
    mldpRecursiveFec.EntityData.YFilter = mldpRecursiveFec.YFilter
    mldpRecursiveFec.EntityData.YangName = "mldp-recursive-fec"
    mldpRecursiveFec.EntityData.BundleName = "cisco_ios_xr"
    mldpRecursiveFec.EntityData.ParentYangName = "af"
    mldpRecursiveFec.EntityData.SegmentPath = "mldp-recursive-fec"
    mldpRecursiveFec.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/" + mldpRecursiveFec.EntityData.SegmentPath
    mldpRecursiveFec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mldpRecursiveFec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mldpRecursiveFec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mldpRecursiveFec.EntityData.Children = types.NewOrderedMap()
    mldpRecursiveFec.EntityData.Leafs = types.NewOrderedMap()
    mldpRecursiveFec.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mldpRecursiveFec.Enable})
    mldpRecursiveFec.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", mldpRecursiveFec.Policy})

    mldpRecursiveFec.EntityData.YListKeys = []string {}

    return &(mldpRecursiveFec.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies
// MLDP neighbor policies
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route Policy. The type is slice of
    // MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy.
    NeighborPolicy []*MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy
}

func (neighborPolicies *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies) GetEntityData() *types.CommonEntityData {
    neighborPolicies.EntityData.YFilter = neighborPolicies.YFilter
    neighborPolicies.EntityData.YangName = "neighbor-policies"
    neighborPolicies.EntityData.BundleName = "cisco_ios_xr"
    neighborPolicies.EntityData.ParentYangName = "af"
    neighborPolicies.EntityData.SegmentPath = "neighbor-policies"
    neighborPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/" + neighborPolicies.EntityData.SegmentPath
    neighborPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborPolicies.EntityData.Children = types.NewOrderedMap()
    neighborPolicies.EntityData.Children.Append("neighbor-policy", types.YChild{"NeighborPolicy", nil})
    for i := range neighborPolicies.NeighborPolicy {
        neighborPolicies.EntityData.Children.Append(types.GetSegmentPath(neighborPolicies.NeighborPolicy[i]), types.YChild{"NeighborPolicy", neighborPolicies.NeighborPolicy[i]})
    }
    neighborPolicies.EntityData.Leafs = types.NewOrderedMap()

    neighborPolicies.EntityData.YListKeys = []string {}

    return &(neighborPolicies.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy
// Route Policy
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RootAddress interface{}

    // This attribute is a key. Inbound/Outbound Policy. The type is
    // MldpPolicyMode.
    PolicyMode interface{}

    // Route policy name. The type is string with length: 1..64. This attribute is
    // mandatory.
    RoutePolicy interface{}
}

func (neighborPolicy *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetEntityData() *types.CommonEntityData {
    neighborPolicy.EntityData.YFilter = neighborPolicy.YFilter
    neighborPolicy.EntityData.YangName = "neighbor-policy"
    neighborPolicy.EntityData.BundleName = "cisco_ios_xr"
    neighborPolicy.EntityData.ParentYangName = "neighbor-policies"
    neighborPolicy.EntityData.SegmentPath = "neighbor-policy" + types.AddKeyToken(neighborPolicy.RootAddress, "root-address") + types.AddKeyToken(neighborPolicy.PolicyMode, "policy-mode")
    neighborPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/neighbor-policies/" + neighborPolicy.EntityData.SegmentPath
    neighborPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborPolicy.EntityData.Children = types.NewOrderedMap()
    neighborPolicy.EntityData.Leafs = types.NewOrderedMap()
    neighborPolicy.EntityData.Leafs.Append("root-address", types.YLeaf{"RootAddress", neighborPolicy.RootAddress})
    neighborPolicy.EntityData.Leafs.Append("policy-mode", types.YLeaf{"PolicyMode", neighborPolicy.PolicyMode})
    neighborPolicy.EntityData.Leafs.Append("route-policy", types.YLeaf{"RoutePolicy", neighborPolicy.RoutePolicy})

    neighborPolicy.EntityData.YListKeys = []string {"RootAddress", "PolicyMode"}

    return &(neighborPolicy.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr
// MPLS mLDP MoFRR
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS mLDP MoFRR. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (moFrr *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MoFrr) GetEntityData() *types.CommonEntityData {
    moFrr.EntityData.YFilter = moFrr.YFilter
    moFrr.EntityData.YangName = "mo-frr"
    moFrr.EntityData.BundleName = "cisco_ios_xr"
    moFrr.EntityData.ParentYangName = "af"
    moFrr.EntityData.SegmentPath = "mo-frr"
    moFrr.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/" + moFrr.EntityData.SegmentPath
    moFrr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moFrr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moFrr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moFrr.EntityData.Children = types.NewOrderedMap()
    moFrr.EntityData.Leafs = types.NewOrderedMap()
    moFrr.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", moFrr.Enable})
    moFrr.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", moFrr.Policy})

    moFrr.EntityData.YListKeys = []string {}

    return &(moFrr.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak
// MPLS mLDP Make-Before-Break configuration
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}

    // Enable MPLS mLDP MBB signaling.
    Signaling MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak) GetEntityData() *types.CommonEntityData {
    makeBeforeBreak.EntityData.YFilter = makeBeforeBreak.YFilter
    makeBeforeBreak.EntityData.YangName = "make-before-break"
    makeBeforeBreak.EntityData.BundleName = "cisco_ios_xr"
    makeBeforeBreak.EntityData.ParentYangName = "af"
    makeBeforeBreak.EntityData.SegmentPath = "make-before-break"
    makeBeforeBreak.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/" + makeBeforeBreak.EntityData.SegmentPath
    makeBeforeBreak.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    makeBeforeBreak.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    makeBeforeBreak.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    makeBeforeBreak.EntityData.Children = types.NewOrderedMap()
    makeBeforeBreak.EntityData.Children.Append("signaling", types.YChild{"Signaling", &makeBeforeBreak.Signaling})
    makeBeforeBreak.EntityData.Leafs = types.NewOrderedMap()
    makeBeforeBreak.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", makeBeforeBreak.Policy})

    makeBeforeBreak.EntityData.YListKeys = []string {}

    return &(makeBeforeBreak.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling
// Enable MPLS mLDP MBB signaling
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding Delay in Seconds. The type is interface{} with range: 0..600.
    // Units are second.
    ForwardDelay interface{}

    // Delete Delay in seconds. The type is interface{} with range: 0..60. Units
    // are second.
    DeleteDelay interface{}
}

func (signaling *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_MakeBeforeBreak_Signaling) GetEntityData() *types.CommonEntityData {
    signaling.EntityData.YFilter = signaling.YFilter
    signaling.EntityData.YangName = "signaling"
    signaling.EntityData.BundleName = "cisco_ios_xr"
    signaling.EntityData.ParentYangName = "make-before-break"
    signaling.EntityData.SegmentPath = "signaling"
    signaling.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/make-before-break/" + signaling.EntityData.SegmentPath
    signaling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signaling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signaling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signaling.EntityData.Children = types.NewOrderedMap()
    signaling.EntityData.Leafs = types.NewOrderedMap()
    signaling.EntityData.Leafs.Append("forward-delay", types.YLeaf{"ForwardDelay", signaling.ForwardDelay})
    signaling.EntityData.Leafs.Append("delete-delay", types.YLeaf{"DeleteDelay", signaling.DeleteDelay})

    signaling.EntityData.YListKeys = []string {}

    return &(signaling.EntityData)
}

// MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc
// MPLS mLDP CSC
type MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS mLDP CSC. The type is interface{}.
    Enable interface{}
}

func (csc *MplsLdp_Global_Mldp_Vrfs_Vrf_Afs_Af_Csc) GetEntityData() *types.CommonEntityData {
    csc.EntityData.YFilter = csc.YFilter
    csc.EntityData.YangName = "csc"
    csc.EntityData.BundleName = "cisco_ios_xr"
    csc.EntityData.ParentYangName = "af"
    csc.EntityData.SegmentPath = "csc"
    csc.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/vrfs/vrf/afs/af/" + csc.EntityData.SegmentPath
    csc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csc.EntityData.Children = types.NewOrderedMap()
    csc.EntityData.Leafs = types.NewOrderedMap()
    csc.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", csc.Enable})

    csc.EntityData.YListKeys = []string {}

    return &(csc.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf
// Default VRF attribute configuration for mLDP
type MplsLdp_Global_Mldp_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family specific operational data.
    Afs MplsLdp_Global_Mldp_DefaultVrf_Afs
}

func (defaultVrf *MplsLdp_Global_Mldp_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "mldp"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/" + defaultVrf.EntityData.SegmentPath
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = types.NewOrderedMap()
    defaultVrf.EntityData.Children.Append("afs", types.YChild{"Afs", &defaultVrf.Afs})
    defaultVrf.EntityData.Leafs = types.NewOrderedMap()

    defaultVrf.EntityData.YListKeys = []string {}

    return &(defaultVrf.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs
// Address Family specific operational data
type MplsLdp_Global_Mldp_DefaultVrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // MplsLdp_Global_Mldp_DefaultVrf_Afs_Af.
    Af []*MplsLdp_Global_Mldp_DefaultVrf_Afs_Af
}

func (afs *MplsLdp_Global_Mldp_DefaultVrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "default-vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af
// Operational data for given Address Family
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (af *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("recursive-forwarding", types.YChild{"RecursiveForwarding", &af.RecursiveForwarding})
    af.EntityData.Children.Append("mldp-recursive-fec", types.YChild{"MldpRecursiveFec", &af.MldpRecursiveFec})
    af.EntityData.Children.Append("neighbor-policies", types.YChild{"NeighborPolicies", &af.NeighborPolicies})
    af.EntityData.Children.Append("mo-frr", types.YChild{"MoFrr", &af.MoFrr})
    af.EntityData.Children.Append("make-before-break", types.YChild{"MakeBeforeBreak", &af.MakeBeforeBreak})
    af.EntityData.Children.Append("csc", types.YChild{"Csc", &af.Csc})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})
    af.EntityData.Leafs.Append("mldp-rib-unicast-always", types.YLeaf{"MldpRibUnicastAlways", af.MldpRibUnicastAlways})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding
// Enable recursive forwarding
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable recursive forwarding. The type is interface{}.
    Enable interface{}

    // Recursive forwarding policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (recursiveForwarding *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_RecursiveForwarding) GetEntityData() *types.CommonEntityData {
    recursiveForwarding.EntityData.YFilter = recursiveForwarding.YFilter
    recursiveForwarding.EntityData.YangName = "recursive-forwarding"
    recursiveForwarding.EntityData.BundleName = "cisco_ios_xr"
    recursiveForwarding.EntityData.ParentYangName = "af"
    recursiveForwarding.EntityData.SegmentPath = "recursive-forwarding"
    recursiveForwarding.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/" + recursiveForwarding.EntityData.SegmentPath
    recursiveForwarding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    recursiveForwarding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    recursiveForwarding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    recursiveForwarding.EntityData.Children = types.NewOrderedMap()
    recursiveForwarding.EntityData.Leafs = types.NewOrderedMap()
    recursiveForwarding.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", recursiveForwarding.Enable})
    recursiveForwarding.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", recursiveForwarding.Policy})

    recursiveForwarding.EntityData.YListKeys = []string {}

    return &(recursiveForwarding.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec
// MPLS mLDP Recursive FEC
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS mLDP Recursive FEC. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (mldpRecursiveFec *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MldpRecursiveFec) GetEntityData() *types.CommonEntityData {
    mldpRecursiveFec.EntityData.YFilter = mldpRecursiveFec.YFilter
    mldpRecursiveFec.EntityData.YangName = "mldp-recursive-fec"
    mldpRecursiveFec.EntityData.BundleName = "cisco_ios_xr"
    mldpRecursiveFec.EntityData.ParentYangName = "af"
    mldpRecursiveFec.EntityData.SegmentPath = "mldp-recursive-fec"
    mldpRecursiveFec.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/" + mldpRecursiveFec.EntityData.SegmentPath
    mldpRecursiveFec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mldpRecursiveFec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mldpRecursiveFec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mldpRecursiveFec.EntityData.Children = types.NewOrderedMap()
    mldpRecursiveFec.EntityData.Leafs = types.NewOrderedMap()
    mldpRecursiveFec.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mldpRecursiveFec.Enable})
    mldpRecursiveFec.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", mldpRecursiveFec.Policy})

    mldpRecursiveFec.EntityData.YListKeys = []string {}

    return &(mldpRecursiveFec.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies
// MLDP neighbor policies
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route Policy. The type is slice of
    // MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy.
    NeighborPolicy []*MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy
}

func (neighborPolicies *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies) GetEntityData() *types.CommonEntityData {
    neighborPolicies.EntityData.YFilter = neighborPolicies.YFilter
    neighborPolicies.EntityData.YangName = "neighbor-policies"
    neighborPolicies.EntityData.BundleName = "cisco_ios_xr"
    neighborPolicies.EntityData.ParentYangName = "af"
    neighborPolicies.EntityData.SegmentPath = "neighbor-policies"
    neighborPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/" + neighborPolicies.EntityData.SegmentPath
    neighborPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborPolicies.EntityData.Children = types.NewOrderedMap()
    neighborPolicies.EntityData.Children.Append("neighbor-policy", types.YChild{"NeighborPolicy", nil})
    for i := range neighborPolicies.NeighborPolicy {
        neighborPolicies.EntityData.Children.Append(types.GetSegmentPath(neighborPolicies.NeighborPolicy[i]), types.YChild{"NeighborPolicy", neighborPolicies.NeighborPolicy[i]})
    }
    neighborPolicies.EntityData.Leafs = types.NewOrderedMap()

    neighborPolicies.EntityData.YListKeys = []string {}

    return &(neighborPolicies.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy
// Route Policy
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RootAddress interface{}

    // This attribute is a key. Inbound/Outbound Policy. The type is
    // MldpPolicyMode.
    PolicyMode interface{}

    // Route policy name. The type is string with length: 1..64. This attribute is
    // mandatory.
    RoutePolicy interface{}
}

func (neighborPolicy *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_NeighborPolicies_NeighborPolicy) GetEntityData() *types.CommonEntityData {
    neighborPolicy.EntityData.YFilter = neighborPolicy.YFilter
    neighborPolicy.EntityData.YangName = "neighbor-policy"
    neighborPolicy.EntityData.BundleName = "cisco_ios_xr"
    neighborPolicy.EntityData.ParentYangName = "neighbor-policies"
    neighborPolicy.EntityData.SegmentPath = "neighbor-policy" + types.AddKeyToken(neighborPolicy.RootAddress, "root-address") + types.AddKeyToken(neighborPolicy.PolicyMode, "policy-mode")
    neighborPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/neighbor-policies/" + neighborPolicy.EntityData.SegmentPath
    neighborPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborPolicy.EntityData.Children = types.NewOrderedMap()
    neighborPolicy.EntityData.Leafs = types.NewOrderedMap()
    neighborPolicy.EntityData.Leafs.Append("root-address", types.YLeaf{"RootAddress", neighborPolicy.RootAddress})
    neighborPolicy.EntityData.Leafs.Append("policy-mode", types.YLeaf{"PolicyMode", neighborPolicy.PolicyMode})
    neighborPolicy.EntityData.Leafs.Append("route-policy", types.YLeaf{"RoutePolicy", neighborPolicy.RoutePolicy})

    neighborPolicy.EntityData.YListKeys = []string {"RootAddress", "PolicyMode"}

    return &(neighborPolicy.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr
// MPLS mLDP MoFRR
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS mLDP MoFRR. The type is interface{}.
    Enable interface{}

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}
}

func (moFrr *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MoFrr) GetEntityData() *types.CommonEntityData {
    moFrr.EntityData.YFilter = moFrr.YFilter
    moFrr.EntityData.YangName = "mo-frr"
    moFrr.EntityData.BundleName = "cisco_ios_xr"
    moFrr.EntityData.ParentYangName = "af"
    moFrr.EntityData.SegmentPath = "mo-frr"
    moFrr.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/" + moFrr.EntityData.SegmentPath
    moFrr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moFrr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moFrr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moFrr.EntityData.Children = types.NewOrderedMap()
    moFrr.EntityData.Leafs = types.NewOrderedMap()
    moFrr.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", moFrr.Enable})
    moFrr.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", moFrr.Policy})

    moFrr.EntityData.YListKeys = []string {}

    return &(moFrr.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak
// MPLS mLDP Make-Before-Break configuration
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy name. The type is string with length: 1..64.
    Policy interface{}

    // Enable MPLS mLDP MBB signaling.
    Signaling MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling
}

func (makeBeforeBreak *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak) GetEntityData() *types.CommonEntityData {
    makeBeforeBreak.EntityData.YFilter = makeBeforeBreak.YFilter
    makeBeforeBreak.EntityData.YangName = "make-before-break"
    makeBeforeBreak.EntityData.BundleName = "cisco_ios_xr"
    makeBeforeBreak.EntityData.ParentYangName = "af"
    makeBeforeBreak.EntityData.SegmentPath = "make-before-break"
    makeBeforeBreak.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/" + makeBeforeBreak.EntityData.SegmentPath
    makeBeforeBreak.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    makeBeforeBreak.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    makeBeforeBreak.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    makeBeforeBreak.EntityData.Children = types.NewOrderedMap()
    makeBeforeBreak.EntityData.Children.Append("signaling", types.YChild{"Signaling", &makeBeforeBreak.Signaling})
    makeBeforeBreak.EntityData.Leafs = types.NewOrderedMap()
    makeBeforeBreak.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", makeBeforeBreak.Policy})

    makeBeforeBreak.EntityData.YListKeys = []string {}

    return &(makeBeforeBreak.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling
// Enable MPLS mLDP MBB signaling
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding Delay in Seconds. The type is interface{} with range: 0..600.
    // Units are second.
    ForwardDelay interface{}

    // Delete Delay in seconds. The type is interface{} with range: 0..60. Units
    // are second.
    DeleteDelay interface{}
}

func (signaling *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_MakeBeforeBreak_Signaling) GetEntityData() *types.CommonEntityData {
    signaling.EntityData.YFilter = signaling.YFilter
    signaling.EntityData.YangName = "signaling"
    signaling.EntityData.BundleName = "cisco_ios_xr"
    signaling.EntityData.ParentYangName = "make-before-break"
    signaling.EntityData.SegmentPath = "signaling"
    signaling.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/make-before-break/" + signaling.EntityData.SegmentPath
    signaling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signaling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signaling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signaling.EntityData.Children = types.NewOrderedMap()
    signaling.EntityData.Leafs = types.NewOrderedMap()
    signaling.EntityData.Leafs.Append("forward-delay", types.YLeaf{"ForwardDelay", signaling.ForwardDelay})
    signaling.EntityData.Leafs.Append("delete-delay", types.YLeaf{"DeleteDelay", signaling.DeleteDelay})

    signaling.EntityData.YListKeys = []string {}

    return &(signaling.EntityData)
}

// MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc
// MPLS mLDP CSC
type MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS mLDP CSC. The type is interface{}.
    Enable interface{}
}

func (csc *MplsLdp_Global_Mldp_DefaultVrf_Afs_Af_Csc) GetEntityData() *types.CommonEntityData {
    csc.EntityData.YFilter = csc.YFilter
    csc.EntityData.YangName = "csc"
    csc.EntityData.BundleName = "cisco_ios_xr"
    csc.EntityData.ParentYangName = "af"
    csc.EntityData.SegmentPath = "csc"
    csc.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/default-vrf/afs/af/" + csc.EntityData.SegmentPath
    csc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csc.EntityData.Children = types.NewOrderedMap()
    csc.EntityData.Leafs = types.NewOrderedMap()
    csc.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", csc.Enable})

    csc.EntityData.YListKeys = []string {}

    return &(csc.EntityData)
}

// MplsLdp_Global_Mldp_MldpGlobal
// Global configuration for mLDP
type MplsLdp_Global_Mldp_MldpGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS mLDP logging.
    Logging MplsLdp_Global_Mldp_MldpGlobal_Logging
}

func (mldpGlobal *MplsLdp_Global_Mldp_MldpGlobal) GetEntityData() *types.CommonEntityData {
    mldpGlobal.EntityData.YFilter = mldpGlobal.YFilter
    mldpGlobal.EntityData.YangName = "mldp-global"
    mldpGlobal.EntityData.BundleName = "cisco_ios_xr"
    mldpGlobal.EntityData.ParentYangName = "mldp"
    mldpGlobal.EntityData.SegmentPath = "mldp-global"
    mldpGlobal.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/" + mldpGlobal.EntityData.SegmentPath
    mldpGlobal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mldpGlobal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mldpGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mldpGlobal.EntityData.Children = types.NewOrderedMap()
    mldpGlobal.EntityData.Children.Append("logging", types.YChild{"Logging", &mldpGlobal.Logging})
    mldpGlobal.EntityData.Leafs = types.NewOrderedMap()

    mldpGlobal.EntityData.YListKeys = []string {}

    return &(mldpGlobal.EntityData)
}

// MplsLdp_Global_Mldp_MldpGlobal_Logging
// MPLS mLDP logging
type MplsLdp_Global_Mldp_MldpGlobal_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS mLDP logging notifications. The type is interface{}.
    Notifications interface{}
}

func (logging *MplsLdp_Global_Mldp_MldpGlobal_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "mldp-global"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp/global/mldp/mldp-global/" + logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("notifications", types.YLeaf{"Notifications", logging.Notifications})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

