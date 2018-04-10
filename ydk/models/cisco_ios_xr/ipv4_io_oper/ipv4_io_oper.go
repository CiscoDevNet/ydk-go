// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-io package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-network: IPv4 network operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_io_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_io_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-io-oper ipv4-network}", reflect.TypeOf(Ipv4Network{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-io-oper:ipv4-network", reflect.TypeOf(Ipv4Network{}))
}

// RpfMode represents Interface line states
type RpfMode string

const (
    // Strict RPF
    RpfMode_strict RpfMode = "strict"

    // Loose RPF
    RpfMode_loose RpfMode = "loose"
)

// Ipv4MaOperLineState represents Interface line states
type Ipv4MaOperLineState string

const (
    // Interface state is unknown
    Ipv4MaOperLineState_unknown Ipv4MaOperLineState = "unknown"

    // Interface has been shutdown
    Ipv4MaOperLineState_shutdown Ipv4MaOperLineState = "shutdown"

    // Interface state is down
    Ipv4MaOperLineState_down Ipv4MaOperLineState = "down"

    // Interface state is up
    Ipv4MaOperLineState_up Ipv4MaOperLineState = "up"
)

// Ipv4Network
// IPv4 network operational data
type Ipv4Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific IPv4 network operational data.
    Nodes Ipv4Network_Nodes

    // IPv4 network operational interface data.
    Interfaces Ipv4Network_Interfaces
}

func (ipv4Network *Ipv4Network) GetEntityData() *types.CommonEntityData {
    ipv4Network.EntityData.YFilter = ipv4Network.YFilter
    ipv4Network.EntityData.YangName = "ipv4-network"
    ipv4Network.EntityData.BundleName = "cisco_ios_xr"
    ipv4Network.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-io-oper"
    ipv4Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-io-oper:ipv4-network"
    ipv4Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Network.EntityData.Children = make(map[string]types.YChild)
    ipv4Network.EntityData.Children["nodes"] = types.YChild{"Nodes", &ipv4Network.Nodes}
    ipv4Network.EntityData.Children["Cisco-IOS-XR-ipv4-ma-oper:interfaces"] = types.YChild{"Interfaces", &ipv4Network.Interfaces}
    ipv4Network.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Network.EntityData)
}

// Ipv4Network_Nodes
// Node-specific IPv4 network operational data
type Ipv4Network_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 network operational data for a particular node. The type is slice of
    // Ipv4Network_Nodes_Node.
    Node []Ipv4Network_Nodes_Node
}

func (nodes *Ipv4Network_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ipv4-network"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Ipv4Network_Nodes_Node
// IPv4 network operational data for a particular
// node
type Ipv4Network_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // IPv4 network operational interface data.
    InterfaceData Ipv4Network_Nodes_Node_InterfaceData

    // Statistical IPv4 network operational data for a node.
    Statistics Ipv4Network_Nodes_Node_Statistics
}

func (node *Ipv4Network_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["interface-data"] = types.YChild{"InterfaceData", &node.InterfaceData}
    node.EntityData.Children["statistics"] = types.YChild{"Statistics", &node.Statistics}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData
// IPv4 network operational interface data
type Ipv4Network_Nodes_Node_InterfaceData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF specific IPv4 network operational interface data.
    Vrfs Ipv4Network_Nodes_Node_InterfaceData_Vrfs

    // Summary of IPv4 network operational interface data on a node.
    Summary Ipv4Network_Nodes_Node_InterfaceData_Summary
}

func (interfaceData *Ipv4Network_Nodes_Node_InterfaceData) GetEntityData() *types.CommonEntityData {
    interfaceData.EntityData.YFilter = interfaceData.YFilter
    interfaceData.EntityData.YangName = "interface-data"
    interfaceData.EntityData.BundleName = "cisco_ios_xr"
    interfaceData.EntityData.ParentYangName = "node"
    interfaceData.EntityData.SegmentPath = "interface-data"
    interfaceData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceData.EntityData.Children = make(map[string]types.YChild)
    interfaceData.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &interfaceData.Vrfs}
    interfaceData.EntityData.Children["summary"] = types.YChild{"Summary", &interfaceData.Summary}
    interfaceData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceData.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs
// VRF specific IPv4 network operational
// interface data
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name of an interface belong to. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf.
    Vrf []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf
}

func (vrfs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "interface-data"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf
// VRF name of an interface belong to
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Brief interface IPv4 network operational data for a node.
    Briefs Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs

    // Detail interface IPv4 network operational data for a node.
    Details Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
}

func (vrf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["briefs"] = types.YChild{"Briefs", &vrf.Briefs}
    vrf.EntityData.Children["details"] = types.YChild{"Details", &vrf.Details}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs
// Brief interface IPv4 network operational
// data for a node
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief interface IPv4 network operational data for an interface. The type is
    // slice of Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief.
    Brief []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
}

func (briefs *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetEntityData() *types.CommonEntityData {
    briefs.EntityData.YFilter = briefs.YFilter
    briefs.EntityData.YangName = "briefs"
    briefs.EntityData.BundleName = "cisco_ios_xr"
    briefs.EntityData.ParentYangName = "vrf"
    briefs.EntityData.SegmentPath = "briefs"
    briefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefs.EntityData.Children = make(map[string]types.YChild)
    briefs.EntityData.Children["brief"] = types.YChild{"Brief", nil}
    for i := range briefs.Brief {
        briefs.EntityData.Children[types.GetSegmentPath(&briefs.Brief[i])] = types.YChild{"Brief", &briefs.Brief[i]}
    }
    briefs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(briefs.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
// Brief interface IPv4 network operational
// data for an interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Primary address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF name of the interface. The type is string.
    VrfName interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}
}

func (brief *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "briefs"
    brief.EntityData.SegmentPath = "brief" + "[interface-name='" + fmt.Sprintf("%v", brief.InterfaceName) + "']"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", brief.InterfaceName}
    brief.EntityData.Leafs["primary-address"] = types.YLeaf{"PrimaryAddress", brief.PrimaryAddress}
    brief.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", brief.VrfId}
    brief.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", brief.VrfName}
    brief.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", brief.LineState}
    return &(brief.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
// Detail interface IPv4 network operational
// data for a node
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail interface IPv4 network operational data for an interface. The type
    // is slice of Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail.
    Detail []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
}

func (details *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "vrf"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range details.Detail {
        details.EntityData.Children[types.GetSegmentPath(&details.Detail[i])] = types.YChild{"Detail", &details.Detail[i]}
    }
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
// Detail interface IPv4 network operational
// data for an interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Primary address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}

    // Prefix length of primary address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route tag associated with the primary address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}

    // IP MTU of the interface. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Are ICMP unreachables sent on the interface?. The type is bool.
    Unreachable interface{}

    // Are ICMP redirects sent on the interface?. The type is bool.
    Redirect interface{}

    // Are direct broadcasts sent on the interface?. The type is bool.
    DirectBroadcast interface{}

    // Are mask replies sent on the interface?. The type is bool.
    MaskReply interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Name of referenced interface (valid if unnumbered). The type is string.
    UnnumberedInterfaceName interface{}

    // Is Proxy ARP disabled on the interface?. The type is bool.
    ProxyArpDisabled interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // ACLs configured on the interface.
    Acl Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl

    // Multi ACLs configured on the interface.
    MultiAcl Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl

    // Helper Addresses configured on the interface.
    HelperAddress Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress

    // RPF config on the interface.
    Rpf Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa

    // Address Publish Time.
    PubUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime

    // IDB Create Time.
    IdbUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime

    // Multicast groups joined on the interface. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup.
    MulticastGroup []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup

    // Secondary addresses on the interface. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress.
    SecondaryAddress []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress
}

func (detail *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "details"
    detail.EntityData.SegmentPath = "detail" + "[interface-name='" + fmt.Sprintf("%v", detail.InterfaceName) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["acl"] = types.YChild{"Acl", &detail.Acl}
    detail.EntityData.Children["multi-acl"] = types.YChild{"MultiAcl", &detail.MultiAcl}
    detail.EntityData.Children["helper-address"] = types.YChild{"HelperAddress", &detail.HelperAddress}
    detail.EntityData.Children["rpf"] = types.YChild{"Rpf", &detail.Rpf}
    detail.EntityData.Children["bgp-pa"] = types.YChild{"BgpPa", &detail.BgpPa}
    detail.EntityData.Children["pub-utime"] = types.YChild{"PubUtime", &detail.PubUtime}
    detail.EntityData.Children["idb-utime"] = types.YChild{"IdbUtime", &detail.IdbUtime}
    detail.EntityData.Children["caps-utime"] = types.YChild{"CapsUtime", &detail.CapsUtime}
    detail.EntityData.Children["fwd-en-utime"] = types.YChild{"FwdEnUtime", &detail.FwdEnUtime}
    detail.EntityData.Children["fwd-dis-utime"] = types.YChild{"FwdDisUtime", &detail.FwdDisUtime}
    detail.EntityData.Children["multicast-group"] = types.YChild{"MulticastGroup", nil}
    for i := range detail.MulticastGroup {
        detail.EntityData.Children[types.GetSegmentPath(&detail.MulticastGroup[i])] = types.YChild{"MulticastGroup", &detail.MulticastGroup[i]}
    }
    detail.EntityData.Children["secondary-address"] = types.YChild{"SecondaryAddress", nil}
    for i := range detail.SecondaryAddress {
        detail.EntityData.Children[types.GetSegmentPath(&detail.SecondaryAddress[i])] = types.YChild{"SecondaryAddress", &detail.SecondaryAddress[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", detail.InterfaceName}
    detail.EntityData.Leafs["primary-address"] = types.YLeaf{"PrimaryAddress", detail.PrimaryAddress}
    detail.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", detail.VrfId}
    detail.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", detail.LineState}
    detail.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", detail.PrefixLength}
    detail.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", detail.RouteTag}
    detail.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", detail.Mtu}
    detail.EntityData.Leafs["unreachable"] = types.YLeaf{"Unreachable", detail.Unreachable}
    detail.EntityData.Leafs["redirect"] = types.YLeaf{"Redirect", detail.Redirect}
    detail.EntityData.Leafs["direct-broadcast"] = types.YLeaf{"DirectBroadcast", detail.DirectBroadcast}
    detail.EntityData.Leafs["mask-reply"] = types.YLeaf{"MaskReply", detail.MaskReply}
    detail.EntityData.Leafs["rg-id-exists"] = types.YLeaf{"RgIdExists", detail.RgIdExists}
    detail.EntityData.Leafs["mlacp-active"] = types.YLeaf{"MlacpActive", detail.MlacpActive}
    detail.EntityData.Leafs["unnumbered-interface-name"] = types.YLeaf{"UnnumberedInterfaceName", detail.UnnumberedInterfaceName}
    detail.EntityData.Leafs["proxy-arp-disabled"] = types.YLeaf{"ProxyArpDisabled", detail.ProxyArpDisabled}
    detail.EntityData.Leafs["flow-tag-src"] = types.YLeaf{"FlowTagSrc", detail.FlowTagSrc}
    detail.EntityData.Leafs["flow-tag-dst"] = types.YLeaf{"FlowTagDst", detail.FlowTagDst}
    return &(detail.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl
// ACLs configured on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    Inbound interface{}

    // ACL applied to outgoing packets. The type is string.
    Outbound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (acl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "cisco_ios_xr"
    acl.EntityData.ParentYangName = "detail"
    acl.EntityData.SegmentPath = "acl"
    acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acl.EntityData.Children = make(map[string]types.YChild)
    acl.EntityData.Leafs = make(map[string]types.YLeaf)
    acl.EntityData.Leafs["inbound"] = types.YLeaf{"Inbound", acl.Inbound}
    acl.EntityData.Leafs["outbound"] = types.YLeaf{"Outbound", acl.Outbound}
    acl.EntityData.Leafs["common-in-bound"] = types.YLeaf{"CommonInBound", acl.CommonInBound}
    acl.EntityData.Leafs["common-out-bound"] = types.YLeaf{"CommonOutBound", acl.CommonOutBound}
    return &(acl.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl
// Multi ACLs configured on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound.
    Inbound []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound

    // Outbound ACLs. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound.
    Outbound []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound

    // Common ACLs. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common.
    Common []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common
}

func (multiAcl *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl) GetEntityData() *types.CommonEntityData {
    multiAcl.EntityData.YFilter = multiAcl.YFilter
    multiAcl.EntityData.YangName = "multi-acl"
    multiAcl.EntityData.BundleName = "cisco_ios_xr"
    multiAcl.EntityData.ParentYangName = "detail"
    multiAcl.EntityData.SegmentPath = "multi-acl"
    multiAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multiAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multiAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multiAcl.EntityData.Children = make(map[string]types.YChild)
    multiAcl.EntityData.Children["inbound"] = types.YChild{"Inbound", nil}
    for i := range multiAcl.Inbound {
        multiAcl.EntityData.Children[types.GetSegmentPath(&multiAcl.Inbound[i])] = types.YChild{"Inbound", &multiAcl.Inbound[i]}
    }
    multiAcl.EntityData.Children["outbound"] = types.YChild{"Outbound", nil}
    for i := range multiAcl.Outbound {
        multiAcl.EntityData.Children[types.GetSegmentPath(&multiAcl.Outbound[i])] = types.YChild{"Outbound", &multiAcl.Outbound[i]}
    }
    multiAcl.EntityData.Children["common"] = types.YChild{"Common", nil}
    for i := range multiAcl.Common {
        multiAcl.EntityData.Children[types.GetSegmentPath(&multiAcl.Common[i])] = types.YChild{"Common", &multiAcl.Common[i]}
    }
    multiAcl.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multiAcl.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound
// Inbound ACLs
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (inbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "multi-acl"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = make(map[string]types.YChild)
    inbound.EntityData.Leafs = make(map[string]types.YLeaf)
    inbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", inbound.Entry}
    return &(inbound.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound
// Outbound ACLs
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (outbound *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "multi-acl"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = make(map[string]types.YChild)
    outbound.EntityData.Leafs = make(map[string]types.YLeaf)
    outbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", outbound.Entry}
    return &(outbound.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common
// Common ACLs
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (common *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAcl_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "multi-acl"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = make(map[string]types.YChild)
    common.EntityData.Leafs = make(map[string]types.YLeaf)
    common.EntityData.Leafs["entry"] = types.YLeaf{"Entry", common.Entry}
    return &(common.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress
// Helper Addresses configured on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Helper address. The type is slice of
    // Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray.
    AddressArray []Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray
}

func (helperAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "detail"
    helperAddress.EntityData.SegmentPath = "helper-address"
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = make(map[string]types.YChild)
    helperAddress.EntityData.Children["address-array"] = types.YChild{"AddressArray", nil}
    for i := range helperAddress.AddressArray {
        helperAddress.EntityData.Children[types.GetSegmentPath(&helperAddress.AddressArray[i])] = types.YChild{"AddressArray", &helperAddress.AddressArray[i]}
    }
    helperAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helperAddress.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray
// Helper address
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Entry interface{}
}

func (addressArray *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_HelperAddress_AddressArray) GetEntityData() *types.CommonEntityData {
    addressArray.EntityData.YFilter = addressArray.YFilter
    addressArray.EntityData.YangName = "address-array"
    addressArray.EntityData.BundleName = "cisco_ios_xr"
    addressArray.EntityData.ParentYangName = "helper-address"
    addressArray.EntityData.SegmentPath = "address-array"
    addressArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressArray.EntityData.Children = make(map[string]types.YChild)
    addressArray.EntityData.Leafs = make(map[string]types.YLeaf)
    addressArray.EntityData.Leafs["entry"] = types.YLeaf{"Entry", addressArray.Entry}
    return &(addressArray.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf
// RPF config on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is RpfMode.
    Mode interface{}
}

func (rpf *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "detail"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = make(map[string]types.YChild)
    rpf.EntityData.Leafs = make(map[string]types.YLeaf)
    rpf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", rpf.Enable}
    rpf.EntityData.Leafs["allow-default-route"] = types.YLeaf{"AllowDefaultRoute", rpf.AllowDefaultRoute}
    rpf.EntityData.Leafs["allow-self-ping"] = types.YLeaf{"AllowSelfPing", rpf.AllowSelfPing}
    rpf.EntityData.Leafs["mode"] = types.YLeaf{"Mode", rpf.Mode}
    return &(rpf.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa
// BGP PA config on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input

    // BGP PA output config.
    Output Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
}

func (bgpPa *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetEntityData() *types.CommonEntityData {
    bgpPa.EntityData.YFilter = bgpPa.YFilter
    bgpPa.EntityData.YangName = "bgp-pa"
    bgpPa.EntityData.BundleName = "cisco_ios_xr"
    bgpPa.EntityData.ParentYangName = "detail"
    bgpPa.EntityData.SegmentPath = "bgp-pa"
    bgpPa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpPa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpPa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpPa.EntityData.Children = make(map[string]types.YChild)
    bgpPa.EntityData.Children["input"] = types.YChild{"Input", &bgpPa.Input}
    bgpPa.EntityData.Children["output"] = types.YChild{"Output", &bgpPa.Output}
    bgpPa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpPa.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input
// BGP PA input config
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "bgp-pa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["enable"] = types.YLeaf{"Enable", input.Enable}
    input.EntityData.Leafs["source"] = types.YLeaf{"Source", input.Source}
    input.EntityData.Leafs["destination"] = types.YLeaf{"Destination", input.Destination}
    return &(input.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
// BGP PA output config
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "bgp-pa"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["enable"] = types.YLeaf{"Enable", output.Enable}
    output.EntityData.Leafs["source"] = types.YLeaf{"Source", output.Source}
    output.EntityData.Leafs["destination"] = types.YLeaf{"Destination", output.Destination}
    return &(output.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime
// Address Publish Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (pubUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_PubUtime) GetEntityData() *types.CommonEntityData {
    pubUtime.EntityData.YFilter = pubUtime.YFilter
    pubUtime.EntityData.YangName = "pub-utime"
    pubUtime.EntityData.BundleName = "cisco_ios_xr"
    pubUtime.EntityData.ParentYangName = "detail"
    pubUtime.EntityData.SegmentPath = "pub-utime"
    pubUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pubUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pubUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pubUtime.EntityData.Children = make(map[string]types.YChild)
    pubUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pubUtime.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime
// IDB Create Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetEntityData() *types.CommonEntityData {
    idbUtime.EntityData.YFilter = idbUtime.YFilter
    idbUtime.EntityData.YangName = "idb-utime"
    idbUtime.EntityData.BundleName = "cisco_ios_xr"
    idbUtime.EntityData.ParentYangName = "detail"
    idbUtime.EntityData.SegmentPath = "idb-utime"
    idbUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idbUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idbUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idbUtime.EntityData.Children = make(map[string]types.YChild)
    idbUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(idbUtime.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime
// CAPS Add Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetEntityData() *types.CommonEntityData {
    capsUtime.EntityData.YFilter = capsUtime.YFilter
    capsUtime.EntityData.YangName = "caps-utime"
    capsUtime.EntityData.BundleName = "cisco_ios_xr"
    capsUtime.EntityData.ParentYangName = "detail"
    capsUtime.EntityData.SegmentPath = "caps-utime"
    capsUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsUtime.EntityData.Children = make(map[string]types.YChild)
    capsUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(capsUtime.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime
// FWD ENABLE Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetEntityData() *types.CommonEntityData {
    fwdEnUtime.EntityData.YFilter = fwdEnUtime.YFilter
    fwdEnUtime.EntityData.YangName = "fwd-en-utime"
    fwdEnUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdEnUtime.EntityData.ParentYangName = "detail"
    fwdEnUtime.EntityData.SegmentPath = "fwd-en-utime"
    fwdEnUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdEnUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdEnUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdEnUtime.EntityData.Children = make(map[string]types.YChild)
    fwdEnUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdEnUtime.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime
// FWD DISABLE Time
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetEntityData() *types.CommonEntityData {
    fwdDisUtime.EntityData.YFilter = fwdDisUtime.YFilter
    fwdDisUtime.EntityData.YangName = "fwd-dis-utime"
    fwdDisUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdDisUtime.EntityData.ParentYangName = "detail"
    fwdDisUtime.EntityData.SegmentPath = "fwd-dis-utime"
    fwdDisUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdDisUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdDisUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdDisUtime.EntityData.Children = make(map[string]types.YChild)
    fwdDisUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdDisUtime.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup
// Multicast groups joined on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of multicast group. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    GroupAddress interface{}
}

func (multicastGroup *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetEntityData() *types.CommonEntityData {
    multicastGroup.EntityData.YFilter = multicastGroup.YFilter
    multicastGroup.EntityData.YangName = "multicast-group"
    multicastGroup.EntityData.BundleName = "cisco_ios_xr"
    multicastGroup.EntityData.ParentYangName = "detail"
    multicastGroup.EntityData.SegmentPath = "multicast-group"
    multicastGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastGroup.EntityData.Children = make(map[string]types.YChild)
    multicastGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    multicastGroup.EntityData.Leafs["group-address"] = types.YLeaf{"GroupAddress", multicastGroup.GroupAddress}
    return &(multicastGroup.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress
// Secondary addresses on the interface
type Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix length of address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route Tag associated with this address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}
}

func (secondaryAddress *Ipv4Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_SecondaryAddress) GetEntityData() *types.CommonEntityData {
    secondaryAddress.EntityData.YFilter = secondaryAddress.YFilter
    secondaryAddress.EntityData.YangName = "secondary-address"
    secondaryAddress.EntityData.BundleName = "cisco_ios_xr"
    secondaryAddress.EntityData.ParentYangName = "detail"
    secondaryAddress.EntityData.SegmentPath = "secondary-address"
    secondaryAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryAddress.EntityData.Children = make(map[string]types.YChild)
    secondaryAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", secondaryAddress.Address}
    secondaryAddress.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", secondaryAddress.PrefixLength}
    secondaryAddress.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", secondaryAddress.RouteTag}
    return &(secondaryAddress.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Summary
// Summary of IPv4 network operational interface
// data on a node
type Ipv4Network_Nodes_Node_InterfaceData_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces (up,down) with basecaps up. The type is interface{}
    // with range: 0..4294967295.
    IfUpDownBasecapsUp interface{}

    // Number of interfaces (up,up).
    IfUpUp Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp

    // Number of interfaces (up,down).
    IfUpDown Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown

    // Number of interfaces (down,down).
    IfDownDown Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown

    // Number of interfaces (shutdown,down).
    IfShutdownDown Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
}

func (summary *Ipv4Network_Nodes_Node_InterfaceData_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "interface-data"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["if-up-up"] = types.YChild{"IfUpUp", &summary.IfUpUp}
    summary.EntityData.Children["if-up-down"] = types.YChild{"IfUpDown", &summary.IfUpDown}
    summary.EntityData.Children["if-down-down"] = types.YChild{"IfDownDown", &summary.IfDownDown}
    summary.EntityData.Children["if-shutdown-down"] = types.YChild{"IfShutdownDown", &summary.IfShutdownDown}
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["if-up-down-basecaps-up"] = types.YLeaf{"IfUpDownBasecapsUp", summary.IfUpDownBasecapsUp}
    return &(summary.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp
// Number of interfaces (up,up)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpUp *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetEntityData() *types.CommonEntityData {
    ifUpUp.EntityData.YFilter = ifUpUp.YFilter
    ifUpUp.EntityData.YangName = "if-up-up"
    ifUpUp.EntityData.BundleName = "cisco_ios_xr"
    ifUpUp.EntityData.ParentYangName = "summary"
    ifUpUp.EntityData.SegmentPath = "if-up-up"
    ifUpUp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifUpUp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifUpUp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifUpUp.EntityData.Children = make(map[string]types.YChild)
    ifUpUp.EntityData.Leafs = make(map[string]types.YLeaf)
    ifUpUp.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifUpUp.IpAssigned}
    ifUpUp.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifUpUp.IpUnnumbered}
    ifUpUp.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifUpUp.IpUnassigned}
    return &(ifUpUp.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown
// Number of interfaces (up,down)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetEntityData() *types.CommonEntityData {
    ifUpDown.EntityData.YFilter = ifUpDown.YFilter
    ifUpDown.EntityData.YangName = "if-up-down"
    ifUpDown.EntityData.BundleName = "cisco_ios_xr"
    ifUpDown.EntityData.ParentYangName = "summary"
    ifUpDown.EntityData.SegmentPath = "if-up-down"
    ifUpDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifUpDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifUpDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifUpDown.EntityData.Children = make(map[string]types.YChild)
    ifUpDown.EntityData.Leafs = make(map[string]types.YLeaf)
    ifUpDown.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifUpDown.IpAssigned}
    ifUpDown.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifUpDown.IpUnnumbered}
    ifUpDown.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifUpDown.IpUnassigned}
    return &(ifUpDown.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown
// Number of interfaces (down,down)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifDownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetEntityData() *types.CommonEntityData {
    ifDownDown.EntityData.YFilter = ifDownDown.YFilter
    ifDownDown.EntityData.YangName = "if-down-down"
    ifDownDown.EntityData.BundleName = "cisco_ios_xr"
    ifDownDown.EntityData.ParentYangName = "summary"
    ifDownDown.EntityData.SegmentPath = "if-down-down"
    ifDownDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifDownDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifDownDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifDownDown.EntityData.Children = make(map[string]types.YChild)
    ifDownDown.EntityData.Leafs = make(map[string]types.YLeaf)
    ifDownDown.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifDownDown.IpAssigned}
    ifDownDown.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifDownDown.IpUnnumbered}
    ifDownDown.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifDownDown.IpUnassigned}
    return &(ifDownDown.EntityData)
}

// Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
// Number of interfaces (shutdown,down)
type Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifShutdownDown *Ipv4Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetEntityData() *types.CommonEntityData {
    ifShutdownDown.EntityData.YFilter = ifShutdownDown.YFilter
    ifShutdownDown.EntityData.YangName = "if-shutdown-down"
    ifShutdownDown.EntityData.BundleName = "cisco_ios_xr"
    ifShutdownDown.EntityData.ParentYangName = "summary"
    ifShutdownDown.EntityData.SegmentPath = "if-shutdown-down"
    ifShutdownDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifShutdownDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifShutdownDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifShutdownDown.EntityData.Children = make(map[string]types.YChild)
    ifShutdownDown.EntityData.Leafs = make(map[string]types.YLeaf)
    ifShutdownDown.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifShutdownDown.IpAssigned}
    ifShutdownDown.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifShutdownDown.IpUnnumbered}
    ifShutdownDown.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifShutdownDown.IpUnassigned}
    return &(ifShutdownDown.EntityData)
}

// Ipv4Network_Nodes_Node_Statistics
// Statistical IPv4 network operational data for
// a node
type Ipv4Network_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic statistics for a node.
    Traffic Ipv4Network_Nodes_Node_Statistics_Traffic
}

func (statistics *Ipv4Network_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["traffic"] = types.YChild{"Traffic", &statistics.Traffic}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Network_Nodes_Node_Statistics_Traffic
// Traffic statistics for a node
type Ipv4Network_Nodes_Node_Statistics_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Network Stats.
    Ipv4Stats Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats

    // ICMP Stats.
    IcmpStats Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats
}

func (traffic *Ipv4Network_Nodes_Node_Statistics_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "statistics"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = make(map[string]types.YChild)
    traffic.EntityData.Children["ipv4-stats"] = types.YChild{"Ipv4Stats", &traffic.Ipv4Stats}
    traffic.EntityData.Children["icmp-stats"] = types.YChild{"IcmpStats", &traffic.IcmpStats}
    traffic.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(traffic.EntityData)
}

// Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats
// IPv4 Network Stats
type Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Packets. The type is interface{} with range: 0..4294967295.
    InputPackets interface{}

    // Received Packets. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Format Errors. The type is interface{} with range: 0..4294967295.
    FormatErrors interface{}

    // Bad Hop Count. The type is interface{} with range: 0..4294967295.
    BadHopCount interface{}

    // Bad Source Address. The type is interface{} with range: 0..4294967295.
    BadSourceAddress interface{}

    // Bad Header. The type is interface{} with range: 0..4294967295.
    BadHeader interface{}

    // No Protocol. The type is interface{} with range: 0..4294967295.
    NoProtocol interface{}

    // No Gateway. The type is interface{} with range: 0..4294967295.
    NoGateway interface{}

    // RaInput. The type is interface{} with range: 0..4294967295.
    ReassembleInput interface{}

    // Reassembled. The type is interface{} with range: 0..4294967295.
    Reassembled interface{}

    // Reassembly Timeout. The type is interface{} with range: 0..4294967295.
    ReassembleTimeout interface{}

    // Reassembly Max Drop. The type is interface{} with range: 0..4294967295.
    ReassembleMaxDrop interface{}

    // Reassembly Failed. The type is interface{} with range: 0..4294967295.
    ReassembleFailed interface{}

    // IP Options Present. The type is interface{} with range: 0..4294967295.
    OptionsPresent interface{}

    // Bad Option. The type is interface{} with range: 0..4294967295.
    BadOption interface{}

    // Unknown Option. The type is interface{} with range: 0..4294967295.
    UnknownOption interface{}

    // Bad Security Option. The type is interface{} with range: 0..4294967295.
    BadSecurityOption interface{}

    // Basic Security Option. The type is interface{} with range: 0..4294967295.
    BasicSecurityOption interface{}

    // Extended Security Option. The type is interface{} with range:
    // 0..4294967295.
    ExtendedSecurityOption interface{}

    // Cipso Option. The type is interface{} with range: 0..4294967295.
    CipsoOption interface{}

    // Strict Source Route Option. The type is interface{} with range:
    // 0..4294967295.
    StrictSourceRouteOption interface{}

    // Loose Source Route Option. The type is interface{} with range:
    // 0..4294967295.
    LooseSourceRouteOption interface{}

    // Record Route Option. The type is interface{} with range: 0..4294967295.
    RecordRouteOption interface{}

    // SID Option. The type is interface{} with range: 0..4294967295.
    SidOption interface{}

    // Timestamp Option. The type is interface{} with range: 0..4294967295.
    TimestampOption interface{}

    // Router Alert Option. The type is interface{} with range: 0..4294967295.
    RouterAlertOption interface{}

    // Noop Option. The type is interface{} with range: 0..4294967295.
    NoopOption interface{}

    // End Option. The type is interface{} with range: 0..4294967295.
    EndOption interface{}

    // Packets Output. The type is interface{} with range: 0..4294967295.
    PacketsOutput interface{}

    // Packets Forwarded. The type is interface{} with range: 0..4294967295.
    PacketsForwarded interface{}

    // Packets Fragmented. The type is interface{} with range: 0..4294967295.
    PacketsFragmented interface{}

    // Fragment Count. The type is interface{} with range: 0..4294967295.
    FragmentCount interface{}

    // Encapsulation Failed. The type is interface{} with range: 0..4294967295.
    EncapsulationFailed interface{}

    // No Router. The type is interface{} with range: 0..4294967295.
    NoRouter interface{}

    // Packet Too Big. The type is interface{} with range: 0..4294967295.
    PacketTooBig interface{}

    // Multicast In. The type is interface{} with range: 0..4294967295.
    MulticastIn interface{}

    // Multicast Out. The type is interface{} with range: 0..4294967295.
    MulticastOut interface{}

    // Broadcast In. The type is interface{} with range: 0..4294967295.
    BroadcastIn interface{}

    // Broadcast Out. The type is interface{} with range: 0..4294967295.
    BroadcastOut interface{}

    // Lisp IPv4 encapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV4Encap interface{}

    // Lisp IPv4 decapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV4Decap interface{}

    // Lisp IPv6 encapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV6Encap interface{}

    // Lisp IPv6 decapped packets. The type is interface{} with range:
    // 0..4294967295.
    LispV6Decap interface{}

    // Lisp encap errors. The type is interface{} with range: 0..4294967295.
    LispEncapError interface{}

    // Lisp decap errors. The type is interface{} with range: 0..4294967295.
    LispDecapError interface{}
}

func (ipv4Stats *Ipv4Network_Nodes_Node_Statistics_Traffic_Ipv4Stats) GetEntityData() *types.CommonEntityData {
    ipv4Stats.EntityData.YFilter = ipv4Stats.YFilter
    ipv4Stats.EntityData.YangName = "ipv4-stats"
    ipv4Stats.EntityData.BundleName = "cisco_ios_xr"
    ipv4Stats.EntityData.ParentYangName = "traffic"
    ipv4Stats.EntityData.SegmentPath = "ipv4-stats"
    ipv4Stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Stats.EntityData.Children = make(map[string]types.YChild)
    ipv4Stats.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Stats.EntityData.Leafs["input-packets"] = types.YLeaf{"InputPackets", ipv4Stats.InputPackets}
    ipv4Stats.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", ipv4Stats.ReceivedPackets}
    ipv4Stats.EntityData.Leafs["format-errors"] = types.YLeaf{"FormatErrors", ipv4Stats.FormatErrors}
    ipv4Stats.EntityData.Leafs["bad-hop-count"] = types.YLeaf{"BadHopCount", ipv4Stats.BadHopCount}
    ipv4Stats.EntityData.Leafs["bad-source-address"] = types.YLeaf{"BadSourceAddress", ipv4Stats.BadSourceAddress}
    ipv4Stats.EntityData.Leafs["bad-header"] = types.YLeaf{"BadHeader", ipv4Stats.BadHeader}
    ipv4Stats.EntityData.Leafs["no-protocol"] = types.YLeaf{"NoProtocol", ipv4Stats.NoProtocol}
    ipv4Stats.EntityData.Leafs["no-gateway"] = types.YLeaf{"NoGateway", ipv4Stats.NoGateway}
    ipv4Stats.EntityData.Leafs["reassemble-input"] = types.YLeaf{"ReassembleInput", ipv4Stats.ReassembleInput}
    ipv4Stats.EntityData.Leafs["reassembled"] = types.YLeaf{"Reassembled", ipv4Stats.Reassembled}
    ipv4Stats.EntityData.Leafs["reassemble-timeout"] = types.YLeaf{"ReassembleTimeout", ipv4Stats.ReassembleTimeout}
    ipv4Stats.EntityData.Leafs["reassemble-max-drop"] = types.YLeaf{"ReassembleMaxDrop", ipv4Stats.ReassembleMaxDrop}
    ipv4Stats.EntityData.Leafs["reassemble-failed"] = types.YLeaf{"ReassembleFailed", ipv4Stats.ReassembleFailed}
    ipv4Stats.EntityData.Leafs["options-present"] = types.YLeaf{"OptionsPresent", ipv4Stats.OptionsPresent}
    ipv4Stats.EntityData.Leafs["bad-option"] = types.YLeaf{"BadOption", ipv4Stats.BadOption}
    ipv4Stats.EntityData.Leafs["unknown-option"] = types.YLeaf{"UnknownOption", ipv4Stats.UnknownOption}
    ipv4Stats.EntityData.Leafs["bad-security-option"] = types.YLeaf{"BadSecurityOption", ipv4Stats.BadSecurityOption}
    ipv4Stats.EntityData.Leafs["basic-security-option"] = types.YLeaf{"BasicSecurityOption", ipv4Stats.BasicSecurityOption}
    ipv4Stats.EntityData.Leafs["extended-security-option"] = types.YLeaf{"ExtendedSecurityOption", ipv4Stats.ExtendedSecurityOption}
    ipv4Stats.EntityData.Leafs["cipso-option"] = types.YLeaf{"CipsoOption", ipv4Stats.CipsoOption}
    ipv4Stats.EntityData.Leafs["strict-source-route-option"] = types.YLeaf{"StrictSourceRouteOption", ipv4Stats.StrictSourceRouteOption}
    ipv4Stats.EntityData.Leafs["loose-source-route-option"] = types.YLeaf{"LooseSourceRouteOption", ipv4Stats.LooseSourceRouteOption}
    ipv4Stats.EntityData.Leafs["record-route-option"] = types.YLeaf{"RecordRouteOption", ipv4Stats.RecordRouteOption}
    ipv4Stats.EntityData.Leafs["sid-option"] = types.YLeaf{"SidOption", ipv4Stats.SidOption}
    ipv4Stats.EntityData.Leafs["timestamp-option"] = types.YLeaf{"TimestampOption", ipv4Stats.TimestampOption}
    ipv4Stats.EntityData.Leafs["router-alert-option"] = types.YLeaf{"RouterAlertOption", ipv4Stats.RouterAlertOption}
    ipv4Stats.EntityData.Leafs["noop-option"] = types.YLeaf{"NoopOption", ipv4Stats.NoopOption}
    ipv4Stats.EntityData.Leafs["end-option"] = types.YLeaf{"EndOption", ipv4Stats.EndOption}
    ipv4Stats.EntityData.Leafs["packets-output"] = types.YLeaf{"PacketsOutput", ipv4Stats.PacketsOutput}
    ipv4Stats.EntityData.Leafs["packets-forwarded"] = types.YLeaf{"PacketsForwarded", ipv4Stats.PacketsForwarded}
    ipv4Stats.EntityData.Leafs["packets-fragmented"] = types.YLeaf{"PacketsFragmented", ipv4Stats.PacketsFragmented}
    ipv4Stats.EntityData.Leafs["fragment-count"] = types.YLeaf{"FragmentCount", ipv4Stats.FragmentCount}
    ipv4Stats.EntityData.Leafs["encapsulation-failed"] = types.YLeaf{"EncapsulationFailed", ipv4Stats.EncapsulationFailed}
    ipv4Stats.EntityData.Leafs["no-router"] = types.YLeaf{"NoRouter", ipv4Stats.NoRouter}
    ipv4Stats.EntityData.Leafs["packet-too-big"] = types.YLeaf{"PacketTooBig", ipv4Stats.PacketTooBig}
    ipv4Stats.EntityData.Leafs["multicast-in"] = types.YLeaf{"MulticastIn", ipv4Stats.MulticastIn}
    ipv4Stats.EntityData.Leafs["multicast-out"] = types.YLeaf{"MulticastOut", ipv4Stats.MulticastOut}
    ipv4Stats.EntityData.Leafs["broadcast-in"] = types.YLeaf{"BroadcastIn", ipv4Stats.BroadcastIn}
    ipv4Stats.EntityData.Leafs["broadcast-out"] = types.YLeaf{"BroadcastOut", ipv4Stats.BroadcastOut}
    ipv4Stats.EntityData.Leafs["lisp-v4-encap"] = types.YLeaf{"LispV4Encap", ipv4Stats.LispV4Encap}
    ipv4Stats.EntityData.Leafs["lisp-v4-decap"] = types.YLeaf{"LispV4Decap", ipv4Stats.LispV4Decap}
    ipv4Stats.EntityData.Leafs["lisp-v6-encap"] = types.YLeaf{"LispV6Encap", ipv4Stats.LispV6Encap}
    ipv4Stats.EntityData.Leafs["lisp-v6-decap"] = types.YLeaf{"LispV6Decap", ipv4Stats.LispV6Decap}
    ipv4Stats.EntityData.Leafs["lisp-encap-error"] = types.YLeaf{"LispEncapError", ipv4Stats.LispEncapError}
    ipv4Stats.EntityData.Leafs["lisp-decap-error"] = types.YLeaf{"LispDecapError", ipv4Stats.LispDecapError}
    return &(ipv4Stats.EntityData)
}

// Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats
// ICMP Stats
type Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICMP Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // ICMP Checksum Errors. The type is interface{} with range: 0..4294967295.
    ChecksumError interface{}

    // ICMP Unknown. The type is interface{} with range: 0..4294967295.
    Unknown interface{}

    // ICMP Transmitted. The type is interface{} with range: 0..4294967295.
    Output interface{}

    // ICMP Admin Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    AdminUnreachableSent interface{}

    // ICMP Network Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    NetworkUnreachableSent interface{}

    // ICMP Host Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    HostUnreachableSent interface{}

    // ICMP Protocol Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    ProtocolUnreachableSent interface{}

    // ICMP Port Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    PortUnreachableSent interface{}

    // ICMP Fragment Unreachable Sent. The type is interface{} with range:
    // 0..4294967295.
    FragmentUnreachableSent interface{}

    // ICMP Admin Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    AdminUnreachableReceived interface{}

    // ICMP Network Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    NetworkUnreachableReceived interface{}

    // ICMP Host Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    HostUnreachableReceived interface{}

    // ICMP Protocol Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    ProtocolUnreachableReceived interface{}

    // ICMP Port Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    PortUnreachableReceived interface{}

    // ICMP Fragment Unreachable Received. The type is interface{} with range:
    // 0..4294967295.
    FragmentUnreachableReceived interface{}

    // ICMP Hopcount Sent. The type is interface{} with range: 0..4294967295.
    HopcountSent interface{}

    // ICMP Reassembly Sent. The type is interface{} with range: 0..4294967295.
    ReassemblySent interface{}

    // ICMP Hopcount Received. The type is interface{} with range: 0..4294967295.
    HopcountReceived interface{}

    // ICMP Reassembly Received. The type is interface{} with range:
    // 0..4294967295.
    ReasseblyReceived interface{}

    // ICMP Parameter Error Received. The type is interface{} with range:
    // 0..4294967295.
    ParamErrorReceived interface{}

    // ICMP Parameter Error Sent. The type is interface{} with range:
    // 0..4294967295.
    ParamErrorSend interface{}

    // ICMP Echo Request Sent. The type is interface{} with range: 0..4294967295.
    EchoRequestSent interface{}

    // ICMP Echo Request Sent. The type is interface{} with range: 0..4294967295.
    EchoRequestReceived interface{}

    // ICMP Echo Reply Sent. The type is interface{} with range: 0..4294967295.
    EchoReplySent interface{}

    // ICMP Echo Reply Received. The type is interface{} with range:
    // 0..4294967295.
    EchoReplyReceived interface{}

    // ICMP Mask Sent. The type is interface{} with range: 0..4294967295.
    MaskRequestSent interface{}

    // ICMP Mask Received. The type is interface{} with range: 0..4294967295.
    MaskRequestReceived interface{}

    // ICMP Mask Sent. The type is interface{} with range: 0..4294967295.
    MaskReplySent interface{}

    // ICMP Mask Received. The type is interface{} with range: 0..4294967295.
    MaskReplyReceived interface{}

    // ICMP Source Quench. The type is interface{} with range: 0..4294967295.
    SourceQuenchReceived interface{}

    // ICMP Redirect Received. The type is interface{} with range: 0..4294967295.
    RedirectReceived interface{}

    // ICMP Redirect Sent. The type is interface{} with range: 0..4294967295.
    RedirectSend interface{}

    // ICMP Timestamp Received. The type is interface{} with range: 0..4294967295.
    TimestampReceived interface{}

    // ICMP Timestamp Reply Received. The type is interface{} with range:
    // 0..4294967295.
    TimestampReplyReceived interface{}

    // ICMP Router Advertisement Received. The type is interface{} with range:
    // 0..4294967295.
    RouterAdvertReceived interface{}

    // ICMP Router Solicited Received. The type is interface{} with range:
    // 0..4294967295.
    RouterSolicitReceived interface{}
}

func (icmpStats *Ipv4Network_Nodes_Node_Statistics_Traffic_IcmpStats) GetEntityData() *types.CommonEntityData {
    icmpStats.EntityData.YFilter = icmpStats.YFilter
    icmpStats.EntityData.YangName = "icmp-stats"
    icmpStats.EntityData.BundleName = "cisco_ios_xr"
    icmpStats.EntityData.ParentYangName = "traffic"
    icmpStats.EntityData.SegmentPath = "icmp-stats"
    icmpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpStats.EntityData.Children = make(map[string]types.YChild)
    icmpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpStats.EntityData.Leafs["received"] = types.YLeaf{"Received", icmpStats.Received}
    icmpStats.EntityData.Leafs["checksum-error"] = types.YLeaf{"ChecksumError", icmpStats.ChecksumError}
    icmpStats.EntityData.Leafs["unknown"] = types.YLeaf{"Unknown", icmpStats.Unknown}
    icmpStats.EntityData.Leafs["output"] = types.YLeaf{"Output", icmpStats.Output}
    icmpStats.EntityData.Leafs["admin-unreachable-sent"] = types.YLeaf{"AdminUnreachableSent", icmpStats.AdminUnreachableSent}
    icmpStats.EntityData.Leafs["network-unreachable-sent"] = types.YLeaf{"NetworkUnreachableSent", icmpStats.NetworkUnreachableSent}
    icmpStats.EntityData.Leafs["host-unreachable-sent"] = types.YLeaf{"HostUnreachableSent", icmpStats.HostUnreachableSent}
    icmpStats.EntityData.Leafs["protocol-unreachable-sent"] = types.YLeaf{"ProtocolUnreachableSent", icmpStats.ProtocolUnreachableSent}
    icmpStats.EntityData.Leafs["port-unreachable-sent"] = types.YLeaf{"PortUnreachableSent", icmpStats.PortUnreachableSent}
    icmpStats.EntityData.Leafs["fragment-unreachable-sent"] = types.YLeaf{"FragmentUnreachableSent", icmpStats.FragmentUnreachableSent}
    icmpStats.EntityData.Leafs["admin-unreachable-received"] = types.YLeaf{"AdminUnreachableReceived", icmpStats.AdminUnreachableReceived}
    icmpStats.EntityData.Leafs["network-unreachable-received"] = types.YLeaf{"NetworkUnreachableReceived", icmpStats.NetworkUnreachableReceived}
    icmpStats.EntityData.Leafs["host-unreachable-received"] = types.YLeaf{"HostUnreachableReceived", icmpStats.HostUnreachableReceived}
    icmpStats.EntityData.Leafs["protocol-unreachable-received"] = types.YLeaf{"ProtocolUnreachableReceived", icmpStats.ProtocolUnreachableReceived}
    icmpStats.EntityData.Leafs["port-unreachable-received"] = types.YLeaf{"PortUnreachableReceived", icmpStats.PortUnreachableReceived}
    icmpStats.EntityData.Leafs["fragment-unreachable-received"] = types.YLeaf{"FragmentUnreachableReceived", icmpStats.FragmentUnreachableReceived}
    icmpStats.EntityData.Leafs["hopcount-sent"] = types.YLeaf{"HopcountSent", icmpStats.HopcountSent}
    icmpStats.EntityData.Leafs["reassembly-sent"] = types.YLeaf{"ReassemblySent", icmpStats.ReassemblySent}
    icmpStats.EntityData.Leafs["hopcount-received"] = types.YLeaf{"HopcountReceived", icmpStats.HopcountReceived}
    icmpStats.EntityData.Leafs["reassebly-received"] = types.YLeaf{"ReasseblyReceived", icmpStats.ReasseblyReceived}
    icmpStats.EntityData.Leafs["param-error-received"] = types.YLeaf{"ParamErrorReceived", icmpStats.ParamErrorReceived}
    icmpStats.EntityData.Leafs["param-error-send"] = types.YLeaf{"ParamErrorSend", icmpStats.ParamErrorSend}
    icmpStats.EntityData.Leafs["echo-request-sent"] = types.YLeaf{"EchoRequestSent", icmpStats.EchoRequestSent}
    icmpStats.EntityData.Leafs["echo-request-received"] = types.YLeaf{"EchoRequestReceived", icmpStats.EchoRequestReceived}
    icmpStats.EntityData.Leafs["echo-reply-sent"] = types.YLeaf{"EchoReplySent", icmpStats.EchoReplySent}
    icmpStats.EntityData.Leafs["echo-reply-received"] = types.YLeaf{"EchoReplyReceived", icmpStats.EchoReplyReceived}
    icmpStats.EntityData.Leafs["mask-request-sent"] = types.YLeaf{"MaskRequestSent", icmpStats.MaskRequestSent}
    icmpStats.EntityData.Leafs["mask-request-received"] = types.YLeaf{"MaskRequestReceived", icmpStats.MaskRequestReceived}
    icmpStats.EntityData.Leafs["mask-reply-sent"] = types.YLeaf{"MaskReplySent", icmpStats.MaskReplySent}
    icmpStats.EntityData.Leafs["mask-reply-received"] = types.YLeaf{"MaskReplyReceived", icmpStats.MaskReplyReceived}
    icmpStats.EntityData.Leafs["source-quench-received"] = types.YLeaf{"SourceQuenchReceived", icmpStats.SourceQuenchReceived}
    icmpStats.EntityData.Leafs["redirect-received"] = types.YLeaf{"RedirectReceived", icmpStats.RedirectReceived}
    icmpStats.EntityData.Leafs["redirect-send"] = types.YLeaf{"RedirectSend", icmpStats.RedirectSend}
    icmpStats.EntityData.Leafs["timestamp-received"] = types.YLeaf{"TimestampReceived", icmpStats.TimestampReceived}
    icmpStats.EntityData.Leafs["timestamp-reply-received"] = types.YLeaf{"TimestampReplyReceived", icmpStats.TimestampReplyReceived}
    icmpStats.EntityData.Leafs["router-advert-received"] = types.YLeaf{"RouterAdvertReceived", icmpStats.RouterAdvertReceived}
    icmpStats.EntityData.Leafs["router-solicit-received"] = types.YLeaf{"RouterSolicitReceived", icmpStats.RouterSolicitReceived}
    return &(icmpStats.EntityData)
}

// Ipv4Network_Interfaces
// IPv4 network operational interface data
type Ipv4Network_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface names with VRF. The type is slice of
    // Ipv4Network_Interfaces_Interface_.
    Interface_ []Ipv4Network_Interfaces_Interface
}

func (interfaces *Ipv4Network_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv4-network"
    interfaces.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ma-oper:interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Ipv4Network_Interfaces_Interface
// Interface names with VRF
type Ipv4Network_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // List of VRF on the interface.
    Vrfs Ipv4Network_Interfaces_Interface_Vrfs
}

func (self *Ipv4Network_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &self.Vrfs}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs
// List of VRF on the interface
type Ipv4Network_Interfaces_Interface_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF information on the interface. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf.
    Vrf []Ipv4Network_Interfaces_Interface_Vrfs_Vrf
}

func (vrfs *Ipv4Network_Interfaces_Interface_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "interface"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf
// VRF information on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Detail IPv4 network operational data for an interface.
    Detail Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail

    // Brief IPv4 network operational data for an interface.
    Brief Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief
}

func (vrf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["detail"] = types.YChild{"Detail", &vrf.Detail}
    vrf.EntityData.Children["brief"] = types.YChild{"Brief", &vrf.Brief}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail
// Detail IPv4 network operational data for an
// interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}

    // Prefix length of primary address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route tag associated with the primary address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}

    // IP MTU of the interface. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Are ICMP unreachables sent on the interface?. The type is bool.
    Unreachable interface{}

    // Are ICMP redirects sent on the interface?. The type is bool.
    Redirect interface{}

    // Are direct broadcasts sent on the interface?. The type is bool.
    DirectBroadcast interface{}

    // Are mask replies sent on the interface?. The type is bool.
    MaskReply interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Name of referenced interface (valid if unnumbered). The type is string.
    UnnumberedInterfaceName interface{}

    // Is Proxy ARP disabled on the interface?. The type is bool.
    ProxyArpDisabled interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // ACLs configured on the interface.
    Acl Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl

    // Multi ACLs configured on the interface.
    MultiAcl Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl

    // Helper Addresses configured on the interface.
    HelperAddress Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress

    // RPF config on the interface.
    Rpf Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa

    // Address Publish Time.
    PubUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime

    // IDB Create Time.
    IdbUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime

    // Multicast groups joined on the interface. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup.
    MulticastGroup []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup

    // Secondary addresses on the interface. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress.
    SecondaryAddress []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress
}

func (detail *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "vrf"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["acl"] = types.YChild{"Acl", &detail.Acl}
    detail.EntityData.Children["multi-acl"] = types.YChild{"MultiAcl", &detail.MultiAcl}
    detail.EntityData.Children["helper-address"] = types.YChild{"HelperAddress", &detail.HelperAddress}
    detail.EntityData.Children["rpf"] = types.YChild{"Rpf", &detail.Rpf}
    detail.EntityData.Children["bgp-pa"] = types.YChild{"BgpPa", &detail.BgpPa}
    detail.EntityData.Children["pub-utime"] = types.YChild{"PubUtime", &detail.PubUtime}
    detail.EntityData.Children["idb-utime"] = types.YChild{"IdbUtime", &detail.IdbUtime}
    detail.EntityData.Children["caps-utime"] = types.YChild{"CapsUtime", &detail.CapsUtime}
    detail.EntityData.Children["fwd-en-utime"] = types.YChild{"FwdEnUtime", &detail.FwdEnUtime}
    detail.EntityData.Children["fwd-dis-utime"] = types.YChild{"FwdDisUtime", &detail.FwdDisUtime}
    detail.EntityData.Children["multicast-group"] = types.YChild{"MulticastGroup", nil}
    for i := range detail.MulticastGroup {
        detail.EntityData.Children[types.GetSegmentPath(&detail.MulticastGroup[i])] = types.YChild{"MulticastGroup", &detail.MulticastGroup[i]}
    }
    detail.EntityData.Children["secondary-address"] = types.YChild{"SecondaryAddress", nil}
    for i := range detail.SecondaryAddress {
        detail.EntityData.Children[types.GetSegmentPath(&detail.SecondaryAddress[i])] = types.YChild{"SecondaryAddress", &detail.SecondaryAddress[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["primary-address"] = types.YLeaf{"PrimaryAddress", detail.PrimaryAddress}
    detail.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", detail.VrfId}
    detail.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", detail.LineState}
    detail.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", detail.PrefixLength}
    detail.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", detail.RouteTag}
    detail.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", detail.Mtu}
    detail.EntityData.Leafs["unreachable"] = types.YLeaf{"Unreachable", detail.Unreachable}
    detail.EntityData.Leafs["redirect"] = types.YLeaf{"Redirect", detail.Redirect}
    detail.EntityData.Leafs["direct-broadcast"] = types.YLeaf{"DirectBroadcast", detail.DirectBroadcast}
    detail.EntityData.Leafs["mask-reply"] = types.YLeaf{"MaskReply", detail.MaskReply}
    detail.EntityData.Leafs["rg-id-exists"] = types.YLeaf{"RgIdExists", detail.RgIdExists}
    detail.EntityData.Leafs["mlacp-active"] = types.YLeaf{"MlacpActive", detail.MlacpActive}
    detail.EntityData.Leafs["unnumbered-interface-name"] = types.YLeaf{"UnnumberedInterfaceName", detail.UnnumberedInterfaceName}
    detail.EntityData.Leafs["proxy-arp-disabled"] = types.YLeaf{"ProxyArpDisabled", detail.ProxyArpDisabled}
    detail.EntityData.Leafs["flow-tag-src"] = types.YLeaf{"FlowTagSrc", detail.FlowTagSrc}
    detail.EntityData.Leafs["flow-tag-dst"] = types.YLeaf{"FlowTagDst", detail.FlowTagDst}
    return &(detail.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl
// ACLs configured on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    Inbound interface{}

    // ACL applied to outgoing packets. The type is string.
    Outbound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (acl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "cisco_ios_xr"
    acl.EntityData.ParentYangName = "detail"
    acl.EntityData.SegmentPath = "acl"
    acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acl.EntityData.Children = make(map[string]types.YChild)
    acl.EntityData.Leafs = make(map[string]types.YLeaf)
    acl.EntityData.Leafs["inbound"] = types.YLeaf{"Inbound", acl.Inbound}
    acl.EntityData.Leafs["outbound"] = types.YLeaf{"Outbound", acl.Outbound}
    acl.EntityData.Leafs["common-in-bound"] = types.YLeaf{"CommonInBound", acl.CommonInBound}
    acl.EntityData.Leafs["common-out-bound"] = types.YLeaf{"CommonOutBound", acl.CommonOutBound}
    return &(acl.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl
// Multi ACLs configured on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Inbound.
    Inbound []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Inbound

    // Outbound ACLs. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Outbound.
    Outbound []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Outbound

    // Common ACLs. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Common.
    Common []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Common
}

func (multiAcl *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl) GetEntityData() *types.CommonEntityData {
    multiAcl.EntityData.YFilter = multiAcl.YFilter
    multiAcl.EntityData.YangName = "multi-acl"
    multiAcl.EntityData.BundleName = "cisco_ios_xr"
    multiAcl.EntityData.ParentYangName = "detail"
    multiAcl.EntityData.SegmentPath = "multi-acl"
    multiAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multiAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multiAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multiAcl.EntityData.Children = make(map[string]types.YChild)
    multiAcl.EntityData.Children["inbound"] = types.YChild{"Inbound", nil}
    for i := range multiAcl.Inbound {
        multiAcl.EntityData.Children[types.GetSegmentPath(&multiAcl.Inbound[i])] = types.YChild{"Inbound", &multiAcl.Inbound[i]}
    }
    multiAcl.EntityData.Children["outbound"] = types.YChild{"Outbound", nil}
    for i := range multiAcl.Outbound {
        multiAcl.EntityData.Children[types.GetSegmentPath(&multiAcl.Outbound[i])] = types.YChild{"Outbound", &multiAcl.Outbound[i]}
    }
    multiAcl.EntityData.Children["common"] = types.YChild{"Common", nil}
    for i := range multiAcl.Common {
        multiAcl.EntityData.Children[types.GetSegmentPath(&multiAcl.Common[i])] = types.YChild{"Common", &multiAcl.Common[i]}
    }
    multiAcl.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multiAcl.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Inbound
// Inbound ACLs
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (inbound *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "multi-acl"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = make(map[string]types.YChild)
    inbound.EntityData.Leafs = make(map[string]types.YLeaf)
    inbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", inbound.Entry}
    return &(inbound.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Outbound
// Outbound ACLs
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (outbound *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "multi-acl"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = make(map[string]types.YChild)
    outbound.EntityData.Leafs = make(map[string]types.YLeaf)
    outbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", outbound.Entry}
    return &(outbound.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Common
// Common ACLs
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (common *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MultiAcl_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "multi-acl"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = make(map[string]types.YChild)
    common.EntityData.Leafs = make(map[string]types.YLeaf)
    common.EntityData.Leafs["entry"] = types.YLeaf{"Entry", common.Entry}
    return &(common.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress
// Helper Addresses configured on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Helper address. The type is slice of
    // Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress_AddressArray.
    AddressArray []Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress_AddressArray
}

func (helperAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "detail"
    helperAddress.EntityData.SegmentPath = "helper-address"
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = make(map[string]types.YChild)
    helperAddress.EntityData.Children["address-array"] = types.YChild{"AddressArray", nil}
    for i := range helperAddress.AddressArray {
        helperAddress.EntityData.Children[types.GetSegmentPath(&helperAddress.AddressArray[i])] = types.YChild{"AddressArray", &helperAddress.AddressArray[i]}
    }
    helperAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helperAddress.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress_AddressArray
// Helper address
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress_AddressArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Entry interface{}
}

func (addressArray *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_HelperAddress_AddressArray) GetEntityData() *types.CommonEntityData {
    addressArray.EntityData.YFilter = addressArray.YFilter
    addressArray.EntityData.YangName = "address-array"
    addressArray.EntityData.BundleName = "cisco_ios_xr"
    addressArray.EntityData.ParentYangName = "helper-address"
    addressArray.EntityData.SegmentPath = "address-array"
    addressArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressArray.EntityData.Children = make(map[string]types.YChild)
    addressArray.EntityData.Leafs = make(map[string]types.YLeaf)
    addressArray.EntityData.Leafs["entry"] = types.YLeaf{"Entry", addressArray.Entry}
    return &(addressArray.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf
// RPF config on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is RpfMode.
    Mode interface{}
}

func (rpf *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "detail"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = make(map[string]types.YChild)
    rpf.EntityData.Leafs = make(map[string]types.YLeaf)
    rpf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", rpf.Enable}
    rpf.EntityData.Leafs["allow-default-route"] = types.YLeaf{"AllowDefaultRoute", rpf.AllowDefaultRoute}
    rpf.EntityData.Leafs["allow-self-ping"] = types.YLeaf{"AllowSelfPing", rpf.AllowSelfPing}
    rpf.EntityData.Leafs["mode"] = types.YLeaf{"Mode", rpf.Mode}
    return &(rpf.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa
// BGP PA config on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input

    // BGP PA output config.
    Output Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output
}

func (bgpPa *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa) GetEntityData() *types.CommonEntityData {
    bgpPa.EntityData.YFilter = bgpPa.YFilter
    bgpPa.EntityData.YangName = "bgp-pa"
    bgpPa.EntityData.BundleName = "cisco_ios_xr"
    bgpPa.EntityData.ParentYangName = "detail"
    bgpPa.EntityData.SegmentPath = "bgp-pa"
    bgpPa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpPa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpPa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpPa.EntityData.Children = make(map[string]types.YChild)
    bgpPa.EntityData.Children["input"] = types.YChild{"Input", &bgpPa.Input}
    bgpPa.EntityData.Children["output"] = types.YChild{"Output", &bgpPa.Output}
    bgpPa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpPa.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input
// BGP PA input config
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "bgp-pa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["enable"] = types.YLeaf{"Enable", input.Enable}
    input.EntityData.Leafs["source"] = types.YLeaf{"Source", input.Source}
    input.EntityData.Leafs["destination"] = types.YLeaf{"Destination", input.Destination}
    return &(input.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output
// BGP PA output config
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is bool.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_BgpPa_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "bgp-pa"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["enable"] = types.YLeaf{"Enable", output.Enable}
    output.EntityData.Leafs["source"] = types.YLeaf{"Source", output.Source}
    output.EntityData.Leafs["destination"] = types.YLeaf{"Destination", output.Destination}
    return &(output.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime
// Address Publish Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (pubUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_PubUtime) GetEntityData() *types.CommonEntityData {
    pubUtime.EntityData.YFilter = pubUtime.YFilter
    pubUtime.EntityData.YangName = "pub-utime"
    pubUtime.EntityData.BundleName = "cisco_ios_xr"
    pubUtime.EntityData.ParentYangName = "detail"
    pubUtime.EntityData.SegmentPath = "pub-utime"
    pubUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pubUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pubUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pubUtime.EntityData.Children = make(map[string]types.YChild)
    pubUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pubUtime.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime
// IDB Create Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_IdbUtime) GetEntityData() *types.CommonEntityData {
    idbUtime.EntityData.YFilter = idbUtime.YFilter
    idbUtime.EntityData.YangName = "idb-utime"
    idbUtime.EntityData.BundleName = "cisco_ios_xr"
    idbUtime.EntityData.ParentYangName = "detail"
    idbUtime.EntityData.SegmentPath = "idb-utime"
    idbUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idbUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idbUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idbUtime.EntityData.Children = make(map[string]types.YChild)
    idbUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(idbUtime.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime
// CAPS Add Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_CapsUtime) GetEntityData() *types.CommonEntityData {
    capsUtime.EntityData.YFilter = capsUtime.YFilter
    capsUtime.EntityData.YangName = "caps-utime"
    capsUtime.EntityData.BundleName = "cisco_ios_xr"
    capsUtime.EntityData.ParentYangName = "detail"
    capsUtime.EntityData.SegmentPath = "caps-utime"
    capsUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsUtime.EntityData.Children = make(map[string]types.YChild)
    capsUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(capsUtime.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime
// FWD ENABLE Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdEnUtime) GetEntityData() *types.CommonEntityData {
    fwdEnUtime.EntityData.YFilter = fwdEnUtime.YFilter
    fwdEnUtime.EntityData.YangName = "fwd-en-utime"
    fwdEnUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdEnUtime.EntityData.ParentYangName = "detail"
    fwdEnUtime.EntityData.SegmentPath = "fwd-en-utime"
    fwdEnUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdEnUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdEnUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdEnUtime.EntityData.Children = make(map[string]types.YChild)
    fwdEnUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdEnUtime.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime
// FWD DISABLE Time
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_FwdDisUtime) GetEntityData() *types.CommonEntityData {
    fwdDisUtime.EntityData.YFilter = fwdDisUtime.YFilter
    fwdDisUtime.EntityData.YangName = "fwd-dis-utime"
    fwdDisUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdDisUtime.EntityData.ParentYangName = "detail"
    fwdDisUtime.EntityData.SegmentPath = "fwd-dis-utime"
    fwdDisUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdDisUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdDisUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdDisUtime.EntityData.Children = make(map[string]types.YChild)
    fwdDisUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdDisUtime.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup
// Multicast groups joined on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of multicast group. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    GroupAddress interface{}
}

func (multicastGroup *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_MulticastGroup) GetEntityData() *types.CommonEntityData {
    multicastGroup.EntityData.YFilter = multicastGroup.YFilter
    multicastGroup.EntityData.YangName = "multicast-group"
    multicastGroup.EntityData.BundleName = "cisco_ios_xr"
    multicastGroup.EntityData.ParentYangName = "detail"
    multicastGroup.EntityData.SegmentPath = "multicast-group"
    multicastGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastGroup.EntityData.Children = make(map[string]types.YChild)
    multicastGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    multicastGroup.EntityData.Leafs["group-address"] = types.YLeaf{"GroupAddress", multicastGroup.GroupAddress}
    return &(multicastGroup.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress
// Secondary addresses on the interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix length of address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Route Tag associated with this address (0 = no tag). The type is
    // interface{} with range: 0..4294967295.
    RouteTag interface{}
}

func (secondaryAddress *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Detail_SecondaryAddress) GetEntityData() *types.CommonEntityData {
    secondaryAddress.EntityData.YFilter = secondaryAddress.YFilter
    secondaryAddress.EntityData.YangName = "secondary-address"
    secondaryAddress.EntityData.BundleName = "cisco_ios_xr"
    secondaryAddress.EntityData.ParentYangName = "detail"
    secondaryAddress.EntityData.SegmentPath = "secondary-address"
    secondaryAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryAddress.EntityData.Children = make(map[string]types.YChild)
    secondaryAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", secondaryAddress.Address}
    secondaryAddress.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", secondaryAddress.PrefixLength}
    secondaryAddress.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", secondaryAddress.RouteTag}
    return &(secondaryAddress.EntityData)
}

// Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief
// Brief IPv4 network operational data for an
// interface
type Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddress interface{}

    // VRF ID of the interface. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF name of the interface. The type is string.
    VrfName interface{}

    // Line state of the interface. The type is Ipv4MaOperLineState.
    LineState interface{}
}

func (brief *Ipv4Network_Interfaces_Interface_Vrfs_Vrf_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "vrf"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["primary-address"] = types.YLeaf{"PrimaryAddress", brief.PrimaryAddress}
    brief.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", brief.VrfId}
    brief.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", brief.VrfName}
    brief.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", brief.LineState}
    return &(brief.EntityData)
}

