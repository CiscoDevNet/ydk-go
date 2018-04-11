// YANG module describing configuration and
// operational data relating to MPLS Static.
package common_mpls_static

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package common_mpls_static"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:common-mpls-static mpls-static}", reflect.TypeOf(MplsStatic{}))
    ydk.RegisterEntity("common-mpls-static:mpls-static", reflect.TypeOf(MplsStatic{}))
}

type LspType struct {
}

func (id LspType) String() string {
	return "common-mpls-static:lsp-type"
}

type LspIPv4 struct {
}

func (id LspIPv4) String() string {
	return "common-mpls-static:lsp-IPv4"
}

type LspIPv6 struct {
}

func (id LspIPv6) String() string {
	return "common-mpls-static:lsp-IPv6"
}

type LspVrf struct {
}

func (id LspVrf) String() string {
	return "common-mpls-static:lsp-vrf"
}

type Lsp struct {
}

func (id Lsp) String() string {
	return "common-mpls-static:lsp"
}

type NexthopResolutionType struct {
}

func (id NexthopResolutionType) String() string {
	return "common-mpls-static:nexthop-resolution-type"
}

type StaticNexthop struct {
}

func (id StaticNexthop) String() string {
	return "common-mpls-static:static-nexthop"
}

type BgpRouteNexthop struct {
}

func (id BgpRouteNexthop) String() string {
	return "common-mpls-static:bgp-route-nexthop"
}

type OspfRouteNexthop struct {
}

func (id OspfRouteNexthop) String() string {
	return "common-mpls-static:ospf-route-nexthop"
}

type IsisRouteNexthop struct {
}

func (id IsisRouteNexthop) String() string {
	return "common-mpls-static:isis-route-nexthop"
}

// Hoptype represents The Nexthop type
type Hoptype string

const (
    // Primary next hop
    Hoptype_PRIMARY Hoptype = "PRIMARY"

    // Backup next hop
    Hoptype_BACKUP Hoptype = "BACKUP"
)

// MplsStatic
// MPLS Static module
type MplsStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static configuration commands.
    MplsStaticCfg MplsStatic_MplsStaticCfg

    // MPLS static operational data.
    MplsStaticState MplsStatic_MplsStaticState
}

func (mplsStatic *MplsStatic) GetEntityData() *types.CommonEntityData {
    mplsStatic.EntityData.YFilter = mplsStatic.YFilter
    mplsStatic.EntityData.YangName = "mpls-static"
    mplsStatic.EntityData.BundleName = "cisco_ios_xe"
    mplsStatic.EntityData.ParentYangName = "common-mpls-static"
    mplsStatic.EntityData.SegmentPath = "common-mpls-static:mpls-static"
    mplsStatic.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsStatic.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsStatic.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsStatic.EntityData.Children = make(map[string]types.YChild)
    mplsStatic.EntityData.Children["mpls-static-cfg"] = types.YChild{"MplsStaticCfg", &mplsStatic.MplsStaticCfg}
    mplsStatic.EntityData.Children["mpls-static-state"] = types.YChild{"MplsStaticState", &mplsStatic.MplsStaticState}
    mplsStatic.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsStatic.EntityData)
}

// MplsStatic_MplsStaticCfg
// MPLS Static configuration commands
type MplsStatic_MplsStaticCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The LSPs indexed by ipv4 prefix.
    Ipv4IngressLsps MplsStatic_MplsStaticCfg_Ipv4IngressLsps

    // The LSPs indexed by in-label.
    InLabelLsps MplsStatic_MplsStaticCfg_InLabelLsps

    // The LSPs indexed by ipv6 prefix.
    Ipv6IngressLsps MplsStatic_MplsStaticCfg_Ipv6IngressLsps

    // The list of interfaces configured with mpls.
    Interfaces MplsStatic_MplsStaticCfg_Interfaces

    // The LSPs indexed by name.
    NamedLsps MplsStatic_MplsStaticCfg_NamedLsps
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetEntityData() *types.CommonEntityData {
    mplsStaticCfg.EntityData.YFilter = mplsStaticCfg.YFilter
    mplsStaticCfg.EntityData.YangName = "mpls-static-cfg"
    mplsStaticCfg.EntityData.BundleName = "cisco_ios_xe"
    mplsStaticCfg.EntityData.ParentYangName = "mpls-static"
    mplsStaticCfg.EntityData.SegmentPath = "mpls-static-cfg"
    mplsStaticCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsStaticCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsStaticCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsStaticCfg.EntityData.Children = make(map[string]types.YChild)
    mplsStaticCfg.EntityData.Children["ipv4-ingress-lsps"] = types.YChild{"Ipv4IngressLsps", &mplsStaticCfg.Ipv4IngressLsps}
    mplsStaticCfg.EntityData.Children["in-label-lsps"] = types.YChild{"InLabelLsps", &mplsStaticCfg.InLabelLsps}
    mplsStaticCfg.EntityData.Children["ipv6-ingress-lsps"] = types.YChild{"Ipv6IngressLsps", &mplsStaticCfg.Ipv6IngressLsps}
    mplsStaticCfg.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &mplsStaticCfg.Interfaces}
    mplsStaticCfg.EntityData.Children["named-lsps"] = types.YChild{"NamedLsps", &mplsStaticCfg.NamedLsps}
    mplsStaticCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsStaticCfg.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps
// The LSPs indexed by ipv4 prefix
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static IPv4 Label Switched Path Configuration at Ingress. The type is
    // slice of MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp.
    Ipv4IngressLsp []MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetEntityData() *types.CommonEntityData {
    ipv4IngressLsps.EntityData.YFilter = ipv4IngressLsps.YFilter
    ipv4IngressLsps.EntityData.YangName = "ipv4-ingress-lsps"
    ipv4IngressLsps.EntityData.BundleName = "cisco_ios_xe"
    ipv4IngressLsps.EntityData.ParentYangName = "mpls-static-cfg"
    ipv4IngressLsps.EntityData.SegmentPath = "ipv4-ingress-lsps"
    ipv4IngressLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4IngressLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4IngressLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4IngressLsps.EntityData.Children = make(map[string]types.YChild)
    ipv4IngressLsps.EntityData.Children["ipv4-ingress-lsp"] = types.YChild{"Ipv4IngressLsp", nil}
    for i := range ipv4IngressLsps.Ipv4IngressLsp {
        ipv4IngressLsps.EntityData.Children[types.GetSegmentPath(&ipv4IngressLsps.Ipv4IngressLsp[i])] = types.YChild{"Ipv4IngressLsp", &ipv4IngressLsps.Ipv4IngressLsp[i]}
    }
    ipv4IngressLsps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4IngressLsps.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp
// MPLS Static IPv4 Label Switched
// Path Configuration at Ingress
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. IPv4 prefix of packets that will ingress on this
    // LSP. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))'.
    Prefix interface{}

    // Value of the local label. Optional for ingress. The type is one of the
    // following types: int with range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // Name of the LSP. The type is string.
    Name interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetEntityData() *types.CommonEntityData {
    ipv4IngressLsp.EntityData.YFilter = ipv4IngressLsp.YFilter
    ipv4IngressLsp.EntityData.YangName = "ipv4-ingress-lsp"
    ipv4IngressLsp.EntityData.BundleName = "cisco_ios_xe"
    ipv4IngressLsp.EntityData.ParentYangName = "ipv4-ingress-lsps"
    ipv4IngressLsp.EntityData.SegmentPath = "ipv4-ingress-lsp" + "[vrf-name='" + fmt.Sprintf("%v", ipv4IngressLsp.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", ipv4IngressLsp.Prefix) + "']"
    ipv4IngressLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4IngressLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4IngressLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4IngressLsp.EntityData.Children = make(map[string]types.YChild)
    ipv4IngressLsp.EntityData.Children["path"] = types.YChild{"Path", &ipv4IngressLsp.Path}
    ipv4IngressLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4IngressLsp.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4IngressLsp.VrfName}
    ipv4IngressLsp.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ipv4IngressLsp.Prefix}
    ipv4IngressLsp.EntityData.Leafs["in-label"] = types.YLeaf{"InLabel", ipv4IngressLsp.InLabel}
    ipv4IngressLsp.EntityData.Leafs["name"] = types.YLeaf{"Name", ipv4IngressLsp.Name}
    return &(ipv4IngressLsp.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables automatic protection if true. The type is bool. The default value
    // is false.
    AutoProtect interface{}

    // next-hops list. The type is slice of
    // MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop.
    NextHop []MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "ipv4-ingress-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range path.NextHop {
        path.EntityData.Children[types.GetSegmentPath(&path.NextHop[i])] = types.YChild{"NextHop", &path.NextHop[i]}
    }
    path.EntityData.Children["operations"] = types.YChild{"Operations", &path.Operations}
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["auto-protect"] = types.YLeaf{"AutoProtect", path.AutoProtect}
    return &(path.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // Index of the nexthop that protects this nexthop. The type is interface{}
    // with range: 0..4294967295.
    ProtectedBy interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type_ interface{}

    // Next-hop.
    NextHopType MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xe"
    nextHop.EntityData.ParentYangName = "path"
    nextHop.EntityData.SegmentPath = "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Children["next-hop-type"] = types.YChild{"NextHopType", &nextHop.NextHopType}
    nextHop.EntityData.Children["operations"] = types.YChild{"Operations", &nextHop.Operations}
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["index"] = types.YLeaf{"Index", nextHop.Index}
    nextHop.EntityData.Leafs["protected-by"] = types.YLeaf{"ProtectedBy", nextHop.ProtectedBy}
    nextHop.EntityData.Leafs["type"] = types.YLeaf{"Type_", nextHop.Type_}
    return &(nextHop.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    OutInterfaceName interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'. This attribute is mandatory.
    MacAddress interface{}

    // IPv4 Address of the nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // The interface index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    IfIndex interface{}
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetEntityData() *types.CommonEntityData {
    nextHopType.EntityData.YFilter = nextHopType.YFilter
    nextHopType.EntityData.YangName = "next-hop-type"
    nextHopType.EntityData.BundleName = "cisco_ios_xe"
    nextHopType.EntityData.ParentYangName = "next-hop"
    nextHopType.EntityData.SegmentPath = "next-hop-type"
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = make(map[string]types.YChild)
    nextHopType.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopType.EntityData.Leafs["out-interface-name"] = types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName}
    nextHopType.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address}
    nextHopType.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextHopType.MacAddress}
    nextHopType.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address}
    nextHopType.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", nextHopType.IfIndex}
    return &(nextHopType.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "next-hop"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "path"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps
// The LSPs indexed by in-label
type MplsStatic_MplsStaticCfg_InLabelLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Non-ingress MPLS Static LSPs, keyed on the incoming label. The type is
    // slice of MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp.
    InLabelLsp []MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetEntityData() *types.CommonEntityData {
    inLabelLsps.EntityData.YFilter = inLabelLsps.YFilter
    inLabelLsps.EntityData.YangName = "in-label-lsps"
    inLabelLsps.EntityData.BundleName = "cisco_ios_xe"
    inLabelLsps.EntityData.ParentYangName = "mpls-static-cfg"
    inLabelLsps.EntityData.SegmentPath = "in-label-lsps"
    inLabelLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    inLabelLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    inLabelLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    inLabelLsps.EntityData.Children = make(map[string]types.YChild)
    inLabelLsps.EntityData.Children["in-label-lsp"] = types.YChild{"InLabelLsp", nil}
    for i := range inLabelLsps.InLabelLsp {
        inLabelLsps.EntityData.Children[types.GetSegmentPath(&inLabelLsps.InLabelLsp[i])] = types.YChild{"InLabelLsp", &inLabelLsps.InLabelLsp[i]}
    }
    inLabelLsps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inLabelLsps.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp
// Non-ingress MPLS Static LSPs,
// keyed on the incoming label
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. Value of the local label. The type is one of the
    // following types: int with range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path
}

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetEntityData() *types.CommonEntityData {
    inLabelLsp.EntityData.YFilter = inLabelLsp.YFilter
    inLabelLsp.EntityData.YangName = "in-label-lsp"
    inLabelLsp.EntityData.BundleName = "cisco_ios_xe"
    inLabelLsp.EntityData.ParentYangName = "in-label-lsps"
    inLabelLsp.EntityData.SegmentPath = "in-label-lsp" + "[vrf-name='" + fmt.Sprintf("%v", inLabelLsp.VrfName) + "']" + "[in-label='" + fmt.Sprintf("%v", inLabelLsp.InLabel) + "']"
    inLabelLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    inLabelLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    inLabelLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    inLabelLsp.EntityData.Children = make(map[string]types.YChild)
    inLabelLsp.EntityData.Children["path"] = types.YChild{"Path", &inLabelLsp.Path}
    inLabelLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    inLabelLsp.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", inLabelLsp.VrfName}
    inLabelLsp.EntityData.Leafs["in-label"] = types.YLeaf{"InLabel", inLabelLsp.InLabel}
    return &(inLabelLsp.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables automatic protection if true. The type is bool. The default value
    // is false.
    AutoProtect interface{}

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations

    // next-hops list. The type is slice of
    // MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop.
    NextHop []MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "in-label-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["operations"] = types.YChild{"Operations", &path.Operations}
    path.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range path.NextHop {
        path.EntityData.Children[types.GetSegmentPath(&path.NextHop[i])] = types.YChild{"NextHop", &path.NextHop[i]}
    }
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["auto-protect"] = types.YLeaf{"AutoProtect", path.AutoProtect}
    return &(path.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "path"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type_ interface{}

    // Index of the nexthop that protects this nexthop. The type is interface{}
    // with range: 0..4294967295.
    ProtectedBy interface{}

    // Next-hop.
    NextHopType MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations
}

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xe"
    nextHop.EntityData.ParentYangName = "path"
    nextHop.EntityData.SegmentPath = "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Children["next-hop-type"] = types.YChild{"NextHopType", &nextHop.NextHopType}
    nextHop.EntityData.Children["operations"] = types.YChild{"Operations", &nextHop.Operations}
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["index"] = types.YLeaf{"Index", nextHop.Index}
    nextHop.EntityData.Leafs["type"] = types.YLeaf{"Type_", nextHop.Type_}
    nextHop.EntityData.Leafs["protected-by"] = types.YLeaf{"ProtectedBy", nextHop.ProtectedBy}
    return &(nextHop.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    IfIndex interface{}

    // IPv4 Address of the nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'. This attribute is mandatory.
    MacAddress interface{}

    // Name of the outgoing interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    OutInterfaceName interface{}
}

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetEntityData() *types.CommonEntityData {
    nextHopType.EntityData.YFilter = nextHopType.YFilter
    nextHopType.EntityData.YangName = "next-hop-type"
    nextHopType.EntityData.BundleName = "cisco_ios_xe"
    nextHopType.EntityData.ParentYangName = "next-hop"
    nextHopType.EntityData.SegmentPath = "next-hop-type"
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = make(map[string]types.YChild)
    nextHopType.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopType.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", nextHopType.IfIndex}
    nextHopType.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address}
    nextHopType.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address}
    nextHopType.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextHopType.MacAddress}
    nextHopType.EntityData.Leafs["out-interface-name"] = types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName}
    return &(nextHopType.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "next-hop"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps
// The LSPs indexed by ipv6 prefix
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static IPv6 Label Switched Path Configuration at Ingress. The type is
    // slice of MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp.
    Ipv6IngressLsp []MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetEntityData() *types.CommonEntityData {
    ipv6IngressLsps.EntityData.YFilter = ipv6IngressLsps.YFilter
    ipv6IngressLsps.EntityData.YangName = "ipv6-ingress-lsps"
    ipv6IngressLsps.EntityData.BundleName = "cisco_ios_xe"
    ipv6IngressLsps.EntityData.ParentYangName = "mpls-static-cfg"
    ipv6IngressLsps.EntityData.SegmentPath = "ipv6-ingress-lsps"
    ipv6IngressLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6IngressLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6IngressLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6IngressLsps.EntityData.Children = make(map[string]types.YChild)
    ipv6IngressLsps.EntityData.Children["ipv6-ingress-lsp"] = types.YChild{"Ipv6IngressLsp", nil}
    for i := range ipv6IngressLsps.Ipv6IngressLsp {
        ipv6IngressLsps.EntityData.Children[types.GetSegmentPath(&ipv6IngressLsps.Ipv6IngressLsp[i])] = types.YChild{"Ipv6IngressLsp", &ipv6IngressLsps.Ipv6IngressLsp[i]}
    }
    ipv6IngressLsps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6IngressLsps.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp
// MPLS Static IPv6 Label Switched Path
// Configuration at Ingress
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. IPv6 prefix of packets that will ingress on this
    // LSP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    Prefix interface{}

    // Name of the LSP. The type is string.
    Name interface{}

    // Value of the local label. Optional for ingress. The type is one of the
    // following types: int with range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path
}

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetEntityData() *types.CommonEntityData {
    ipv6IngressLsp.EntityData.YFilter = ipv6IngressLsp.YFilter
    ipv6IngressLsp.EntityData.YangName = "ipv6-ingress-lsp"
    ipv6IngressLsp.EntityData.BundleName = "cisco_ios_xe"
    ipv6IngressLsp.EntityData.ParentYangName = "ipv6-ingress-lsps"
    ipv6IngressLsp.EntityData.SegmentPath = "ipv6-ingress-lsp" + "[vrf-name='" + fmt.Sprintf("%v", ipv6IngressLsp.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", ipv6IngressLsp.Prefix) + "']"
    ipv6IngressLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6IngressLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6IngressLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6IngressLsp.EntityData.Children = make(map[string]types.YChild)
    ipv6IngressLsp.EntityData.Children["path"] = types.YChild{"Path", &ipv6IngressLsp.Path}
    ipv6IngressLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6IngressLsp.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6IngressLsp.VrfName}
    ipv6IngressLsp.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ipv6IngressLsp.Prefix}
    ipv6IngressLsp.EntityData.Leafs["name"] = types.YLeaf{"Name", ipv6IngressLsp.Name}
    ipv6IngressLsp.EntityData.Leafs["in-label"] = types.YLeaf{"InLabel", ipv6IngressLsp.InLabel}
    return &(ipv6IngressLsp.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables automatic protection if true. The type is bool. The default value
    // is false.
    AutoProtect interface{}

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations

    // next-hops list. The type is slice of
    // MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop.
    NextHop []MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "ipv6-ingress-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["operations"] = types.YChild{"Operations", &path.Operations}
    path.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range path.NextHop {
        path.EntityData.Children[types.GetSegmentPath(&path.NextHop[i])] = types.YChild{"NextHop", &path.NextHop[i]}
    }
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["auto-protect"] = types.YLeaf{"AutoProtect", path.AutoProtect}
    return &(path.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "path"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type_ interface{}

    // Index of the nexthop that protects this nexthop. The type is interface{}
    // with range: 0..4294967295.
    ProtectedBy interface{}

    // Next-hop.
    NextHopType MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xe"
    nextHop.EntityData.ParentYangName = "path"
    nextHop.EntityData.SegmentPath = "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Children["next-hop-type"] = types.YChild{"NextHopType", &nextHop.NextHopType}
    nextHop.EntityData.Children["operations"] = types.YChild{"Operations", &nextHop.Operations}
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["index"] = types.YLeaf{"Index", nextHop.Index}
    nextHop.EntityData.Leafs["type"] = types.YLeaf{"Type_", nextHop.Type_}
    nextHop.EntityData.Leafs["protected-by"] = types.YLeaf{"ProtectedBy", nextHop.ProtectedBy}
    return &(nextHop.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    IfIndex interface{}

    // IPv4 Address of the nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'. This attribute is mandatory.
    MacAddress interface{}

    // Name of the outgoing interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    OutInterfaceName interface{}
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetEntityData() *types.CommonEntityData {
    nextHopType.EntityData.YFilter = nextHopType.YFilter
    nextHopType.EntityData.YangName = "next-hop-type"
    nextHopType.EntityData.BundleName = "cisco_ios_xe"
    nextHopType.EntityData.ParentYangName = "next-hop"
    nextHopType.EntityData.SegmentPath = "next-hop-type"
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = make(map[string]types.YChild)
    nextHopType.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopType.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", nextHopType.IfIndex}
    nextHopType.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address}
    nextHopType.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address}
    nextHopType.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextHopType.MacAddress}
    nextHopType.EntityData.Leafs["out-interface-name"] = types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName}
    return &(nextHopType.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "next-hop"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Interfaces
// The list of interfaces configured with mpls
type MplsStatic_MplsStaticCfg_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interfaces that can be enabled under mpls static. The type is slice
    // of MplsStatic_MplsStaticCfg_Interfaces_Interface_.
    Interface_ []MplsStatic_MplsStaticCfg_Interfaces_Interface
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "mpls-static-cfg"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// MplsStatic_MplsStaticCfg_Interfaces_Interface
// List of interfaces that can be enabled under
// mpls static
type MplsStatic_MplsStaticCfg_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name This attribute is mandatory.
    Name interface{}

    // Interface Enabled boolean. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    Enabled interface{}
}

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    self.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", self.Enabled}
    return &(self.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps
// The LSPs indexed by name
type MplsStatic_MplsStaticCfg_NamedLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static Label Switched Path Configuration. The LSPs in this list are
    // referenced by a string name. The LSPs may be ingress/egress/crossconnect,
    // may have v4/v6 prefixes and may be associated with any VRF. The other
    // specialized lists above are for implemetations that are keyed on prefixes
    // or in-labels instead of the LSP name. The type is slice of
    // MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp.
    NamedLsp []MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetEntityData() *types.CommonEntityData {
    namedLsps.EntityData.YFilter = namedLsps.YFilter
    namedLsps.EntityData.YangName = "named-lsps"
    namedLsps.EntityData.BundleName = "cisco_ios_xe"
    namedLsps.EntityData.ParentYangName = "mpls-static-cfg"
    namedLsps.EntityData.SegmentPath = "named-lsps"
    namedLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    namedLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    namedLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    namedLsps.EntityData.Children = make(map[string]types.YChild)
    namedLsps.EntityData.Children["named-lsp"] = types.YChild{"NamedLsp", nil}
    for i := range namedLsps.NamedLsp {
        namedLsps.EntityData.Children[types.GetSegmentPath(&namedLsps.NamedLsp[i])] = types.YChild{"NamedLsp", &namedLsps.NamedLsp[i]}
    }
    namedLsps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(namedLsps.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp
// MPLS Static Label Switched Path Configuration.
// The LSPs in this list are referenced by a string name.
// The LSPs may be ingress/egress/crossconnect,
// may have v4/v6 prefixes and may be associated with any
// VRF. The other specialized lists above are for
// implemetations that are keyed on prefixes or in-labels
// instead of the LSP name.
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. Name of the LSP. The type is string. This
    // attribute is mandatory.
    Name interface{}

    // lsp type. The type is one of the following: LspIPv4LspIPv6LspVrfLsp. This
    // attribute is mandatory.
    LspType interface{}

    // Value of the local label. The type is one of the following types: int with
    // range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // ipv4 prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))'.
    Ipv4Prefix interface{}

    // ipv6 prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    Ipv6Prefix interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path
}

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetEntityData() *types.CommonEntityData {
    namedLsp.EntityData.YFilter = namedLsp.YFilter
    namedLsp.EntityData.YangName = "named-lsp"
    namedLsp.EntityData.BundleName = "cisco_ios_xe"
    namedLsp.EntityData.ParentYangName = "named-lsps"
    namedLsp.EntityData.SegmentPath = "named-lsp" + "[vrf-name='" + fmt.Sprintf("%v", namedLsp.VrfName) + "']" + "[name='" + fmt.Sprintf("%v", namedLsp.Name) + "']"
    namedLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    namedLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    namedLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    namedLsp.EntityData.Children = make(map[string]types.YChild)
    namedLsp.EntityData.Children["path"] = types.YChild{"Path", &namedLsp.Path}
    namedLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    namedLsp.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", namedLsp.VrfName}
    namedLsp.EntityData.Leafs["name"] = types.YLeaf{"Name", namedLsp.Name}
    namedLsp.EntityData.Leafs["lsp-type"] = types.YLeaf{"LspType", namedLsp.LspType}
    namedLsp.EntityData.Leafs["in-label"] = types.YLeaf{"InLabel", namedLsp.InLabel}
    namedLsp.EntityData.Leafs["ipv4-prefix"] = types.YLeaf{"Ipv4Prefix", namedLsp.Ipv4Prefix}
    namedLsp.EntityData.Leafs["ipv6-prefix"] = types.YLeaf{"Ipv6Prefix", namedLsp.Ipv6Prefix}
    return &(namedLsp.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables automatic protection if true. The type is bool. The default value
    // is false.
    AutoProtect interface{}

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations

    // next-hops list. The type is slice of
    // MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop.
    NextHop []MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "named-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["operations"] = types.YChild{"Operations", &path.Operations}
    path.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range path.NextHop {
        path.EntityData.Children[types.GetSegmentPath(&path.NextHop[i])] = types.YChild{"NextHop", &path.NextHop[i]}
    }
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["auto-protect"] = types.YLeaf{"AutoProtect", path.AutoProtect}
    return &(path.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "path"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type_ interface{}

    // Index of the nexthop that protects this nexthop. The type is interface{}
    // with range: 0..4294967295.
    ProtectedBy interface{}

    // Next-hop.
    NextHopType MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations
}

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xe"
    nextHop.EntityData.ParentYangName = "path"
    nextHop.EntityData.SegmentPath = "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Children["next-hop-type"] = types.YChild{"NextHopType", &nextHop.NextHopType}
    nextHop.EntityData.Children["operations"] = types.YChild{"Operations", &nextHop.Operations}
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["index"] = types.YLeaf{"Index", nextHop.Index}
    nextHop.EntityData.Leafs["type"] = types.YLeaf{"Type_", nextHop.Type_}
    nextHop.EntityData.Leafs["protected-by"] = types.YLeaf{"ProtectedBy", nextHop.ProtectedBy}
    return &(nextHop.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    IfIndex interface{}

    // IPv4 Address of the nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'. This attribute is mandatory.
    MacAddress interface{}

    // Name of the outgoing interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    OutInterfaceName interface{}
}

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetEntityData() *types.CommonEntityData {
    nextHopType.EntityData.YFilter = nextHopType.YFilter
    nextHopType.EntityData.YangName = "next-hop-type"
    nextHopType.EntityData.BundleName = "cisco_ios_xe"
    nextHopType.EntityData.ParentYangName = "next-hop"
    nextHopType.EntityData.SegmentPath = "next-hop-type"
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = make(map[string]types.YChild)
    nextHopType.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopType.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", nextHopType.IfIndex}
    nextHopType.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address}
    nextHopType.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address}
    nextHopType.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextHopType.MacAddress}
    nextHopType.EntityData.Leafs["out-interface-name"] = types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName}
    return &(nextHopType.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "next-hop"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticState
// MPLS static operational data
type MplsStatic_MplsStaticState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static Label Switched Paths.
    LabelSwitchedPaths MplsStatic_MplsStaticState_LabelSwitchedPaths
}

func (mplsStaticState *MplsStatic_MplsStaticState) GetEntityData() *types.CommonEntityData {
    mplsStaticState.EntityData.YFilter = mplsStaticState.YFilter
    mplsStaticState.EntityData.YangName = "mpls-static-state"
    mplsStaticState.EntityData.BundleName = "cisco_ios_xe"
    mplsStaticState.EntityData.ParentYangName = "mpls-static"
    mplsStaticState.EntityData.SegmentPath = "mpls-static-state"
    mplsStaticState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsStaticState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsStaticState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsStaticState.EntityData.Children = make(map[string]types.YChild)
    mplsStaticState.EntityData.Children["label-switched-paths"] = types.YChild{"LabelSwitchedPaths", &mplsStaticState.LabelSwitchedPaths}
    mplsStaticState.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsStaticState.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths
// MPLS Static Label Switched Paths
type MplsStatic_MplsStaticState_LabelSwitchedPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LSPs list. The type is slice of
    // MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath.
    LabelSwitchedPath []MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetEntityData() *types.CommonEntityData {
    labelSwitchedPaths.EntityData.YFilter = labelSwitchedPaths.YFilter
    labelSwitchedPaths.EntityData.YangName = "label-switched-paths"
    labelSwitchedPaths.EntityData.BundleName = "cisco_ios_xe"
    labelSwitchedPaths.EntityData.ParentYangName = "mpls-static-state"
    labelSwitchedPaths.EntityData.SegmentPath = "label-switched-paths"
    labelSwitchedPaths.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelSwitchedPaths.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelSwitchedPaths.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelSwitchedPaths.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedPaths.EntityData.Children["label-switched-path"] = types.YChild{"LabelSwitchedPath", nil}
    for i := range labelSwitchedPaths.LabelSwitchedPath {
        labelSwitchedPaths.EntityData.Children[types.GetSegmentPath(&labelSwitchedPaths.LabelSwitchedPath[i])] = types.YChild{"LabelSwitchedPath", &labelSwitchedPaths.LabelSwitchedPath[i]}
    }
    labelSwitchedPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(labelSwitchedPaths.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath
// MPLS LSPs list
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. IP v4/v6 prefix. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    Prefix interface{}

    // Name of the LSP. The type is string.
    Name interface{}

    // Value of the local label. The type is one of the following types: int with
    // range: 16..1048575, or enumeration IetfMplsLabel.
    InLabelValue interface{}

    // ingress stats.
    IngressStats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats

    // egress stats.
    EgressStats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats

    // Fowarding path.
    Path MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path
}

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetEntityData() *types.CommonEntityData {
    labelSwitchedPath.EntityData.YFilter = labelSwitchedPath.YFilter
    labelSwitchedPath.EntityData.YangName = "label-switched-path"
    labelSwitchedPath.EntityData.BundleName = "cisco_ios_xe"
    labelSwitchedPath.EntityData.ParentYangName = "label-switched-paths"
    labelSwitchedPath.EntityData.SegmentPath = "label-switched-path" + "[vrf-name='" + fmt.Sprintf("%v", labelSwitchedPath.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", labelSwitchedPath.Prefix) + "']"
    labelSwitchedPath.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelSwitchedPath.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelSwitchedPath.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelSwitchedPath.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedPath.EntityData.Children["ingress-stats"] = types.YChild{"IngressStats", &labelSwitchedPath.IngressStats}
    labelSwitchedPath.EntityData.Children["egress-stats"] = types.YChild{"EgressStats", &labelSwitchedPath.EgressStats}
    labelSwitchedPath.EntityData.Children["path"] = types.YChild{"Path", &labelSwitchedPath.Path}
    labelSwitchedPath.EntityData.Leafs = make(map[string]types.YLeaf)
    labelSwitchedPath.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", labelSwitchedPath.VrfName}
    labelSwitchedPath.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", labelSwitchedPath.Prefix}
    labelSwitchedPath.EntityData.Leafs["name"] = types.YLeaf{"Name", labelSwitchedPath.Name}
    labelSwitchedPath.EntityData.Leafs["in-label-value"] = types.YLeaf{"InLabelValue", labelSwitchedPath.InLabelValue}
    return &(labelSwitchedPath.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats
// ingress stats
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics.
    Stats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats
}

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetEntityData() *types.CommonEntityData {
    ingressStats.EntityData.YFilter = ingressStats.YFilter
    ingressStats.EntityData.YangName = "ingress-stats"
    ingressStats.EntityData.BundleName = "cisco_ios_xe"
    ingressStats.EntityData.ParentYangName = "label-switched-path"
    ingressStats.EntityData.SegmentPath = "ingress-stats"
    ingressStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ingressStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ingressStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ingressStats.EntityData.Children = make(map[string]types.YChild)
    ingressStats.EntityData.Children["stats"] = types.YChild{"Stats", &ingressStats.Stats}
    ingressStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ingressStats.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats
// Statistics
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stats for packet count. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // stats for byte count. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // stats for dropped-packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}

    // stats for dropped-bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedBytes interface{}
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xe"
    stats.EntityData.ParentYangName = "ingress-stats"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["packets"] = types.YLeaf{"Packets", stats.Packets}
    stats.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", stats.Bytes}
    stats.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", stats.DroppedPackets}
    stats.EntityData.Leafs["dropped-bytes"] = types.YLeaf{"DroppedBytes", stats.DroppedBytes}
    return &(stats.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats
// egress stats
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics.
    Stats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats
}

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetEntityData() *types.CommonEntityData {
    egressStats.EntityData.YFilter = egressStats.YFilter
    egressStats.EntityData.YangName = "egress-stats"
    egressStats.EntityData.BundleName = "cisco_ios_xe"
    egressStats.EntityData.ParentYangName = "label-switched-path"
    egressStats.EntityData.SegmentPath = "egress-stats"
    egressStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egressStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egressStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egressStats.EntityData.Children = make(map[string]types.YChild)
    egressStats.EntityData.Children["stats"] = types.YChild{"Stats", &egressStats.Stats}
    egressStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(egressStats.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats
// Statistics
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stats for packet count. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // stats for byte count. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // stats for dropped-packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}

    // stats for dropped-bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedBytes interface{}
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xe"
    stats.EntityData.ParentYangName = "egress-stats"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["packets"] = types.YLeaf{"Packets", stats.Packets}
    stats.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", stats.Bytes}
    stats.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", stats.DroppedPackets}
    stats.EntityData.Leafs["dropped-bytes"] = types.YLeaf{"DroppedBytes", stats.DroppedBytes}
    return &(stats.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path
// Fowarding path
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables automatic protection if true. The type is bool. The default value
    // is false.
    AutoProtect interface{}

    // The incoming label processing.
    Operations MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations

    // next-hops list. The type is slice of
    // MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop.
    NextHop []MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop
}

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "label-switched-path"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["operations"] = types.YChild{"Operations", &path.Operations}
    path.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range path.NextHop {
        path.EntityData.Children[types.GetSegmentPath(&path.NextHop[i])] = types.YChild{"NextHop", &path.NextHop[i]}
    }
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["auto-protect"] = types.YLeaf{"AutoProtect", path.AutoProtect}
    return &(path.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "path"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type_ interface{}

    // Index of the nexthop that protects this nexthop. The type is interface{}
    // with range: 0..4294967295.
    ProtectedBy interface{}

    // The origin of this nexthop. The type is one of the following:
    // StaticNexthopBgpRouteNexthopOspfRouteNexthopIsisRouteNexthop.
    Origin interface{}

    // Next-hop.
    NextHopType MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType

    // The incoming label processing.
    Operations MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations

    // lsp stats.
    NexthopStats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats
}

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xe"
    nextHop.EntityData.ParentYangName = "path"
    nextHop.EntityData.SegmentPath = "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Children["next-hop-type"] = types.YChild{"NextHopType", &nextHop.NextHopType}
    nextHop.EntityData.Children["operations"] = types.YChild{"Operations", &nextHop.Operations}
    nextHop.EntityData.Children["nexthop-stats"] = types.YChild{"NexthopStats", &nextHop.NexthopStats}
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["index"] = types.YLeaf{"Index", nextHop.Index}
    nextHop.EntityData.Leafs["type"] = types.YLeaf{"Type_", nextHop.Type_}
    nextHop.EntityData.Leafs["protected-by"] = types.YLeaf{"ProtectedBy", nextHop.ProtectedBy}
    nextHop.EntityData.Leafs["origin"] = types.YLeaf{"Origin", nextHop.Origin}
    return &(nextHop.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    IfIndex interface{}

    // IPv4 Address of the nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'. This attribute is mandatory.
    MacAddress interface{}

    // Name of the outgoing interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    OutInterfaceName interface{}
}

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetEntityData() *types.CommonEntityData {
    nextHopType.EntityData.YFilter = nextHopType.YFilter
    nextHopType.EntityData.YangName = "next-hop-type"
    nextHopType.EntityData.BundleName = "cisco_ios_xe"
    nextHopType.EntityData.ParentYangName = "next-hop"
    nextHopType.EntityData.SegmentPath = "next-hop-type"
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = make(map[string]types.YChild)
    nextHopType.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopType.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", nextHopType.IfIndex}
    nextHopType.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address}
    nextHopType.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address}
    nextHopType.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextHopType.MacAddress}
    nextHopType.EntityData.Leafs["out-interface-name"] = types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName}
    return &(nextHopType.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "next-hop"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["swap"] = types.YChild{"Swap", &operations.Swap}
    operations.EntityData.Children["push"] = types.YChild{"Push", &operations.Push}
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    operations.EntityData.Leafs["preserve"] = types.YLeaf{"Preserve", operations.Preserve}
    operations.EntityData.Leafs["pop-and-forward"] = types.YLeaf{"PopAndForward", operations.PopAndForward}
    return &(operations.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetEntityData() *types.CommonEntityData {
    swap.EntityData.YFilter = swap.YFilter
    swap.EntityData.YangName = "swap"
    swap.EntityData.BundleName = "cisco_ios_xe"
    swap.EntityData.ParentYangName = "operations"
    swap.EntityData.SegmentPath = "swap"
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = make(map[string]types.YChild)
    swap.EntityData.Children["stack"] = types.YChild{"Stack", &swap.Stack}
    swap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swap.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "swap"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetEntityData() *types.CommonEntityData {
    push.EntityData.YFilter = push.YFilter
    push.EntityData.YangName = "push"
    push.EntityData.BundleName = "cisco_ios_xe"
    push.EntityData.ParentYangName = "operations"
    push.EntityData.SegmentPath = "push"
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = make(map[string]types.YChild)
    push.EntityData.Children["stack"] = types.YChild{"Stack", &push.Stack}
    push.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(push.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xe"
    stack.EntityData.ParentYangName = "push"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["label-stack"] = types.YLeaf{"LabelStack", stack.LabelStack}
    return &(stack.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats
// lsp stats
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics.
    Stats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats
}

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetEntityData() *types.CommonEntityData {
    nexthopStats.EntityData.YFilter = nexthopStats.YFilter
    nexthopStats.EntityData.YangName = "nexthop-stats"
    nexthopStats.EntityData.BundleName = "cisco_ios_xe"
    nexthopStats.EntityData.ParentYangName = "next-hop"
    nexthopStats.EntityData.SegmentPath = "nexthop-stats"
    nexthopStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nexthopStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nexthopStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nexthopStats.EntityData.Children = make(map[string]types.YChild)
    nexthopStats.EntityData.Children["stats"] = types.YChild{"Stats", &nexthopStats.Stats}
    nexthopStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nexthopStats.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats
// Statistics
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stats for packet count. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // stats for byte count. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // stats for dropped-packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}

    // stats for dropped-bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedBytes interface{}
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xe"
    stats.EntityData.ParentYangName = "nexthop-stats"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["packets"] = types.YLeaf{"Packets", stats.Packets}
    stats.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", stats.Bytes}
    stats.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", stats.DroppedPackets}
    stats.EntityData.Leafs["dropped-bytes"] = types.YLeaf{"DroppedBytes", stats.DroppedBytes}
    return &(stats.EntityData)
}

