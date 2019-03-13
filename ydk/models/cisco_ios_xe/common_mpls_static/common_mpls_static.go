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

type BgpRouteNexthop struct {
}

func (id BgpRouteNexthop) String() string {
	return "common-mpls-static:bgp-route-nexthop"
}

type NexthopResolutionType struct {
}

func (id NexthopResolutionType) String() string {
	return "common-mpls-static:nexthop-resolution-type"
}

type IsisRouteNexthop struct {
}

func (id IsisRouteNexthop) String() string {
	return "common-mpls-static:isis-route-nexthop"
}

type LspType struct {
}

func (id LspType) String() string {
	return "common-mpls-static:lsp-type"
}

type StaticNexthop struct {
}

func (id StaticNexthop) String() string {
	return "common-mpls-static:static-nexthop"
}

type LspIPv6 struct {
}

func (id LspIPv6) String() string {
	return "common-mpls-static:lsp-IPv6"
}

type LspIPv4 struct {
}

func (id LspIPv4) String() string {
	return "common-mpls-static:lsp-IPv4"
}

type OspfRouteNexthop struct {
}

func (id OspfRouteNexthop) String() string {
	return "common-mpls-static:ospf-route-nexthop"
}

type Lsp struct {
}

func (id Lsp) String() string {
	return "common-mpls-static:lsp"
}

type LspVrf struct {
}

func (id LspVrf) String() string {
	return "common-mpls-static:lsp-vrf"
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
    mplsStatic.EntityData.AbsolutePath = mplsStatic.EntityData.SegmentPath
    mplsStatic.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsStatic.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsStatic.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsStatic.EntityData.Children = types.NewOrderedMap()
    mplsStatic.EntityData.Children.Append("mpls-static-cfg", types.YChild{"MplsStaticCfg", &mplsStatic.MplsStaticCfg})
    mplsStatic.EntityData.Children.Append("mpls-static-state", types.YChild{"MplsStaticState", &mplsStatic.MplsStaticState})
    mplsStatic.EntityData.Leafs = types.NewOrderedMap()

    mplsStatic.EntityData.YListKeys = []string {}

    return &(mplsStatic.EntityData)
}

// MplsStatic_MplsStaticCfg
// MPLS Static configuration commands
type MplsStatic_MplsStaticCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The LSPs indexed by ipv6 prefix.
    Ipv6IngressLsps MplsStatic_MplsStaticCfg_Ipv6IngressLsps

    // The LSPs indexed by name.
    NamedLsps MplsStatic_MplsStaticCfg_NamedLsps

    // The list of interfaces configured with mpls.
    Interfaces MplsStatic_MplsStaticCfg_Interfaces

    // The LSPs indexed by in-label.
    InLabelLsps MplsStatic_MplsStaticCfg_InLabelLsps

    // The LSPs indexed by ipv4 prefix.
    Ipv4IngressLsps MplsStatic_MplsStaticCfg_Ipv4IngressLsps
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetEntityData() *types.CommonEntityData {
    mplsStaticCfg.EntityData.YFilter = mplsStaticCfg.YFilter
    mplsStaticCfg.EntityData.YangName = "mpls-static-cfg"
    mplsStaticCfg.EntityData.BundleName = "cisco_ios_xe"
    mplsStaticCfg.EntityData.ParentYangName = "mpls-static"
    mplsStaticCfg.EntityData.SegmentPath = "mpls-static-cfg"
    mplsStaticCfg.EntityData.AbsolutePath = "common-mpls-static:mpls-static/" + mplsStaticCfg.EntityData.SegmentPath
    mplsStaticCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsStaticCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsStaticCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsStaticCfg.EntityData.Children = types.NewOrderedMap()
    mplsStaticCfg.EntityData.Children.Append("ipv6-ingress-lsps", types.YChild{"Ipv6IngressLsps", &mplsStaticCfg.Ipv6IngressLsps})
    mplsStaticCfg.EntityData.Children.Append("named-lsps", types.YChild{"NamedLsps", &mplsStaticCfg.NamedLsps})
    mplsStaticCfg.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &mplsStaticCfg.Interfaces})
    mplsStaticCfg.EntityData.Children.Append("in-label-lsps", types.YChild{"InLabelLsps", &mplsStaticCfg.InLabelLsps})
    mplsStaticCfg.EntityData.Children.Append("ipv4-ingress-lsps", types.YChild{"Ipv4IngressLsps", &mplsStaticCfg.Ipv4IngressLsps})
    mplsStaticCfg.EntityData.Leafs = types.NewOrderedMap()

    mplsStaticCfg.EntityData.YListKeys = []string {}

    return &(mplsStaticCfg.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps
// The LSPs indexed by ipv6 prefix
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static IPv6 Label Switched Path Configuration at Ingress. The type is
    // slice of MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp.
    Ipv6IngressLsp []*MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetEntityData() *types.CommonEntityData {
    ipv6IngressLsps.EntityData.YFilter = ipv6IngressLsps.YFilter
    ipv6IngressLsps.EntityData.YangName = "ipv6-ingress-lsps"
    ipv6IngressLsps.EntityData.BundleName = "cisco_ios_xe"
    ipv6IngressLsps.EntityData.ParentYangName = "mpls-static-cfg"
    ipv6IngressLsps.EntityData.SegmentPath = "ipv6-ingress-lsps"
    ipv6IngressLsps.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/" + ipv6IngressLsps.EntityData.SegmentPath
    ipv6IngressLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6IngressLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6IngressLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6IngressLsps.EntityData.Children = types.NewOrderedMap()
    ipv6IngressLsps.EntityData.Children.Append("ipv6-ingress-lsp", types.YChild{"Ipv6IngressLsp", nil})
    for i := range ipv6IngressLsps.Ipv6IngressLsp {
        ipv6IngressLsps.EntityData.Children.Append(types.GetSegmentPath(ipv6IngressLsps.Ipv6IngressLsp[i]), types.YChild{"Ipv6IngressLsp", ipv6IngressLsps.Ipv6IngressLsp[i]})
    }
    ipv6IngressLsps.EntityData.Leafs = types.NewOrderedMap()

    ipv6IngressLsps.EntityData.YListKeys = []string {}

    return &(ipv6IngressLsps.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp
// MPLS Static IPv6 Label Switched Path
// Configuration at Ingress
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. IPv6 prefix of packets that will ingress on this
    // LSP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
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
    ipv6IngressLsp.EntityData.SegmentPath = "ipv6-ingress-lsp" + types.AddKeyToken(ipv6IngressLsp.VrfName, "vrf-name") + types.AddKeyToken(ipv6IngressLsp.Prefix, "prefix")
    ipv6IngressLsp.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/" + ipv6IngressLsp.EntityData.SegmentPath
    ipv6IngressLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6IngressLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6IngressLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6IngressLsp.EntityData.Children = types.NewOrderedMap()
    ipv6IngressLsp.EntityData.Children.Append("path", types.YChild{"Path", &ipv6IngressLsp.Path})
    ipv6IngressLsp.EntityData.Leafs = types.NewOrderedMap()
    ipv6IngressLsp.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6IngressLsp.VrfName})
    ipv6IngressLsp.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ipv6IngressLsp.Prefix})
    ipv6IngressLsp.EntityData.Leafs.Append("name", types.YLeaf{"Name", ipv6IngressLsp.Name})
    ipv6IngressLsp.EntityData.Leafs.Append("in-label", types.YLeaf{"InLabel", ipv6IngressLsp.InLabel})

    ipv6IngressLsp.EntityData.YListKeys = []string {"VrfName", "Prefix"}

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
    NextHop []*MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "ipv6-ingress-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/" + path.EntityData.SegmentPath
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("operations", types.YChild{"Operations", &path.Operations})
    path.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range path.NextHop {
        path.EntityData.Children.Append(types.GetSegmentPath(path.NextHop[i]), types.YChild{"NextHop", path.NextHop[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("auto-protect", types.YLeaf{"AutoProtect", path.AutoProtect})

    path.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type interface{}

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
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop-type", types.YChild{"NextHopType", &nextHop.NextHopType})
    nextHop.EntityData.Children.Append("operations", types.YChild{"Operations", &nextHop.Operations})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("protected-by", types.YLeaf{"ProtectedBy", nextHop.ProtectedBy})

    nextHop.EntityData.YListKeys = []string {"Index"}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This attribute is mandatory.
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
    nextHopType.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/next-hop/" + nextHopType.EntityData.SegmentPath
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = types.NewOrderedMap()
    nextHopType.EntityData.Leafs = types.NewOrderedMap()
    nextHopType.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", nextHopType.IfIndex})
    nextHopType.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address})
    nextHopType.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address})
    nextHopType.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", nextHopType.MacAddress})
    nextHopType.EntityData.Leafs.Append("out-interface-name", types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName})

    nextHopType.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/next-hop/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/next-hop/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/next-hop/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/next-hop/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv6-ingress-lsps/ipv6-ingress-lsp/path/next-hop/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
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
    NamedLsp []*MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetEntityData() *types.CommonEntityData {
    namedLsps.EntityData.YFilter = namedLsps.YFilter
    namedLsps.EntityData.YangName = "named-lsps"
    namedLsps.EntityData.BundleName = "cisco_ios_xe"
    namedLsps.EntityData.ParentYangName = "mpls-static-cfg"
    namedLsps.EntityData.SegmentPath = "named-lsps"
    namedLsps.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/" + namedLsps.EntityData.SegmentPath
    namedLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    namedLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    namedLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    namedLsps.EntityData.Children = types.NewOrderedMap()
    namedLsps.EntityData.Children.Append("named-lsp", types.YChild{"NamedLsp", nil})
    for i := range namedLsps.NamedLsp {
        namedLsps.EntityData.Children.Append(types.GetSegmentPath(namedLsps.NamedLsp[i]), types.YChild{"NamedLsp", namedLsps.NamedLsp[i]})
    }
    namedLsps.EntityData.Leafs = types.NewOrderedMap()

    namedLsps.EntityData.YListKeys = []string {}

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
    YListKey string

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. Name of the LSP. The type is string. This
    // attribute is mandatory.
    Name interface{}

    // lsp type. The type is one of the following: LspIPv6LspIPv4LspLspVrf. This
    // attribute is mandatory.
    LspType interface{}

    // Value of the local label. The type is one of the following types: int with
    // range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // ipv4 prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Ipv4Prefix interface{}

    // ipv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Ipv6Prefix interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path
}

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetEntityData() *types.CommonEntityData {
    namedLsp.EntityData.YFilter = namedLsp.YFilter
    namedLsp.EntityData.YangName = "named-lsp"
    namedLsp.EntityData.BundleName = "cisco_ios_xe"
    namedLsp.EntityData.ParentYangName = "named-lsps"
    namedLsp.EntityData.SegmentPath = "named-lsp" + types.AddKeyToken(namedLsp.VrfName, "vrf-name") + types.AddKeyToken(namedLsp.Name, "name")
    namedLsp.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/" + namedLsp.EntityData.SegmentPath
    namedLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    namedLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    namedLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    namedLsp.EntityData.Children = types.NewOrderedMap()
    namedLsp.EntityData.Children.Append("path", types.YChild{"Path", &namedLsp.Path})
    namedLsp.EntityData.Leafs = types.NewOrderedMap()
    namedLsp.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", namedLsp.VrfName})
    namedLsp.EntityData.Leafs.Append("name", types.YLeaf{"Name", namedLsp.Name})
    namedLsp.EntityData.Leafs.Append("lsp-type", types.YLeaf{"LspType", namedLsp.LspType})
    namedLsp.EntityData.Leafs.Append("in-label", types.YLeaf{"InLabel", namedLsp.InLabel})
    namedLsp.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", namedLsp.Ipv4Prefix})
    namedLsp.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", namedLsp.Ipv6Prefix})

    namedLsp.EntityData.YListKeys = []string {"VrfName", "Name"}

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
    NextHop []*MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "named-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/" + path.EntityData.SegmentPath
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("operations", types.YChild{"Operations", &path.Operations})
    path.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range path.NextHop {
        path.EntityData.Children.Append(types.GetSegmentPath(path.NextHop[i]), types.YChild{"NextHop", path.NextHop[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("auto-protect", types.YLeaf{"AutoProtect", path.AutoProtect})

    path.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type interface{}

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
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop-type", types.YChild{"NextHopType", &nextHop.NextHopType})
    nextHop.EntityData.Children.Append("operations", types.YChild{"Operations", &nextHop.Operations})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("protected-by", types.YLeaf{"ProtectedBy", nextHop.ProtectedBy})

    nextHop.EntityData.YListKeys = []string {"Index"}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This attribute is mandatory.
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
    nextHopType.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/next-hop/" + nextHopType.EntityData.SegmentPath
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = types.NewOrderedMap()
    nextHopType.EntityData.Leafs = types.NewOrderedMap()
    nextHopType.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", nextHopType.IfIndex})
    nextHopType.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address})
    nextHopType.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address})
    nextHopType.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", nextHopType.MacAddress})
    nextHopType.EntityData.Leafs.Append("out-interface-name", types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName})

    nextHopType.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/next-hop/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/next-hop/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/next-hop/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/next-hop/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/named-lsps/named-lsp/path/next-hop/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Interfaces
// The list of interfaces configured with mpls
type MplsStatic_MplsStaticCfg_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interfaces that can be enabled under mpls static. The type is slice
    // of MplsStatic_MplsStaticCfg_Interfaces_Interface.
    Interface []*MplsStatic_MplsStaticCfg_Interfaces_Interface
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "mpls-static-cfg"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MplsStatic_MplsStaticCfg_Interfaces_Interface
// List of interfaces that can be enabled under
// mpls static
type MplsStatic_MplsStaticCfg_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Name, "name")
    self.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", self.Enabled})

    self.EntityData.YListKeys = []string {"Name"}

    return &(self.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps
// The LSPs indexed by in-label
type MplsStatic_MplsStaticCfg_InLabelLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Non-ingress MPLS Static LSPs, keyed on the incoming label. The type is
    // slice of MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp.
    InLabelLsp []*MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetEntityData() *types.CommonEntityData {
    inLabelLsps.EntityData.YFilter = inLabelLsps.YFilter
    inLabelLsps.EntityData.YangName = "in-label-lsps"
    inLabelLsps.EntityData.BundleName = "cisco_ios_xe"
    inLabelLsps.EntityData.ParentYangName = "mpls-static-cfg"
    inLabelLsps.EntityData.SegmentPath = "in-label-lsps"
    inLabelLsps.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/" + inLabelLsps.EntityData.SegmentPath
    inLabelLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    inLabelLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    inLabelLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    inLabelLsps.EntityData.Children = types.NewOrderedMap()
    inLabelLsps.EntityData.Children.Append("in-label-lsp", types.YChild{"InLabelLsp", nil})
    for i := range inLabelLsps.InLabelLsp {
        inLabelLsps.EntityData.Children.Append(types.GetSegmentPath(inLabelLsps.InLabelLsp[i]), types.YChild{"InLabelLsp", inLabelLsps.InLabelLsp[i]})
    }
    inLabelLsps.EntityData.Leafs = types.NewOrderedMap()

    inLabelLsps.EntityData.YListKeys = []string {}

    return &(inLabelLsps.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp
// Non-ingress MPLS Static LSPs,
// keyed on the incoming label
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    inLabelLsp.EntityData.SegmentPath = "in-label-lsp" + types.AddKeyToken(inLabelLsp.VrfName, "vrf-name") + types.AddKeyToken(inLabelLsp.InLabel, "in-label")
    inLabelLsp.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/" + inLabelLsp.EntityData.SegmentPath
    inLabelLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    inLabelLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    inLabelLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    inLabelLsp.EntityData.Children = types.NewOrderedMap()
    inLabelLsp.EntityData.Children.Append("path", types.YChild{"Path", &inLabelLsp.Path})
    inLabelLsp.EntityData.Leafs = types.NewOrderedMap()
    inLabelLsp.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", inLabelLsp.VrfName})
    inLabelLsp.EntityData.Leafs.Append("in-label", types.YLeaf{"InLabel", inLabelLsp.InLabel})

    inLabelLsp.EntityData.YListKeys = []string {"VrfName", "InLabel"}

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
    NextHop []*MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "in-label-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/" + path.EntityData.SegmentPath
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("operations", types.YChild{"Operations", &path.Operations})
    path.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range path.NextHop {
        path.EntityData.Children.Append(types.GetSegmentPath(path.NextHop[i]), types.YChild{"NextHop", path.NextHop[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("auto-protect", types.YLeaf{"AutoProtect", path.AutoProtect})

    path.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type interface{}

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
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop-type", types.YChild{"NextHopType", &nextHop.NextHopType})
    nextHop.EntityData.Children.Append("operations", types.YChild{"Operations", &nextHop.Operations})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("protected-by", types.YLeaf{"ProtectedBy", nextHop.ProtectedBy})

    nextHop.EntityData.YListKeys = []string {"Index"}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This attribute is mandatory.
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
    nextHopType.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/next-hop/" + nextHopType.EntityData.SegmentPath
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = types.NewOrderedMap()
    nextHopType.EntityData.Leafs = types.NewOrderedMap()
    nextHopType.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", nextHopType.IfIndex})
    nextHopType.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address})
    nextHopType.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address})
    nextHopType.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", nextHopType.MacAddress})
    nextHopType.EntityData.Leafs.Append("out-interface-name", types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName})

    nextHopType.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/next-hop/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/next-hop/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/next-hop/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/next-hop/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/in-label-lsps/in-label-lsp/path/next-hop/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps
// The LSPs indexed by ipv4 prefix
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static IPv4 Label Switched Path Configuration at Ingress. The type is
    // slice of MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp.
    Ipv4IngressLsp []*MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetEntityData() *types.CommonEntityData {
    ipv4IngressLsps.EntityData.YFilter = ipv4IngressLsps.YFilter
    ipv4IngressLsps.EntityData.YangName = "ipv4-ingress-lsps"
    ipv4IngressLsps.EntityData.BundleName = "cisco_ios_xe"
    ipv4IngressLsps.EntityData.ParentYangName = "mpls-static-cfg"
    ipv4IngressLsps.EntityData.SegmentPath = "ipv4-ingress-lsps"
    ipv4IngressLsps.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/" + ipv4IngressLsps.EntityData.SegmentPath
    ipv4IngressLsps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4IngressLsps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4IngressLsps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4IngressLsps.EntityData.Children = types.NewOrderedMap()
    ipv4IngressLsps.EntityData.Children.Append("ipv4-ingress-lsp", types.YChild{"Ipv4IngressLsp", nil})
    for i := range ipv4IngressLsps.Ipv4IngressLsp {
        ipv4IngressLsps.EntityData.Children.Append(types.GetSegmentPath(ipv4IngressLsps.Ipv4IngressLsp[i]), types.YChild{"Ipv4IngressLsp", ipv4IngressLsps.Ipv4IngressLsp[i]})
    }
    ipv4IngressLsps.EntityData.Leafs = types.NewOrderedMap()

    ipv4IngressLsps.EntityData.YListKeys = []string {}

    return &(ipv4IngressLsps.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp
// MPLS Static IPv4 Label Switched
// Path Configuration at Ingress
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. IPv4 prefix of packets that will ingress on this
    // LSP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Prefix interface{}

    // Name of the LSP. The type is string.
    Name interface{}

    // Value of the local label. Optional for ingress. The type is one of the
    // following types: int with range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetEntityData() *types.CommonEntityData {
    ipv4IngressLsp.EntityData.YFilter = ipv4IngressLsp.YFilter
    ipv4IngressLsp.EntityData.YangName = "ipv4-ingress-lsp"
    ipv4IngressLsp.EntityData.BundleName = "cisco_ios_xe"
    ipv4IngressLsp.EntityData.ParentYangName = "ipv4-ingress-lsps"
    ipv4IngressLsp.EntityData.SegmentPath = "ipv4-ingress-lsp" + types.AddKeyToken(ipv4IngressLsp.VrfName, "vrf-name") + types.AddKeyToken(ipv4IngressLsp.Prefix, "prefix")
    ipv4IngressLsp.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/" + ipv4IngressLsp.EntityData.SegmentPath
    ipv4IngressLsp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv4IngressLsp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv4IngressLsp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv4IngressLsp.EntityData.Children = types.NewOrderedMap()
    ipv4IngressLsp.EntityData.Children.Append("path", types.YChild{"Path", &ipv4IngressLsp.Path})
    ipv4IngressLsp.EntityData.Leafs = types.NewOrderedMap()
    ipv4IngressLsp.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4IngressLsp.VrfName})
    ipv4IngressLsp.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ipv4IngressLsp.Prefix})
    ipv4IngressLsp.EntityData.Leafs.Append("name", types.YLeaf{"Name", ipv4IngressLsp.Name})
    ipv4IngressLsp.EntityData.Leafs.Append("in-label", types.YLeaf{"InLabel", ipv4IngressLsp.InLabel})

    ipv4IngressLsp.EntityData.YListKeys = []string {"VrfName", "Prefix"}

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

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations

    // next-hops list. The type is slice of
    // MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop.
    NextHop []*MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "ipv4-ingress-lsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/" + path.EntityData.SegmentPath
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("operations", types.YChild{"Operations", &path.Operations})
    path.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range path.NextHop {
        path.EntityData.Children.Append(types.GetSegmentPath(path.NextHop[i]), types.YChild{"NextHop", path.NextHop[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("auto-protect", types.YLeaf{"AutoProtect", path.AutoProtect})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type interface{}

    // Index of the nexthop that protects this nexthop. The type is interface{}
    // with range: 0..4294967295.
    ProtectedBy interface{}

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
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop-type", types.YChild{"NextHopType", &nextHop.NextHopType})
    nextHop.EntityData.Children.Append("operations", types.YChild{"Operations", &nextHop.Operations})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("protected-by", types.YLeaf{"ProtectedBy", nextHop.ProtectedBy})

    nextHop.EntityData.YListKeys = []string {"Index"}

    return &(nextHop.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address of the nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // IPv4 Address of the nexthop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // The interface index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    IfIndex interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This attribute is mandatory.
    MacAddress interface{}

    // Name of the outgoing interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    OutInterfaceName interface{}
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetEntityData() *types.CommonEntityData {
    nextHopType.EntityData.YFilter = nextHopType.YFilter
    nextHopType.EntityData.YangName = "next-hop-type"
    nextHopType.EntityData.BundleName = "cisco_ios_xe"
    nextHopType.EntityData.ParentYangName = "next-hop"
    nextHopType.EntityData.SegmentPath = "next-hop-type"
    nextHopType.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/next-hop/" + nextHopType.EntityData.SegmentPath
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = types.NewOrderedMap()
    nextHopType.EntityData.Leafs = types.NewOrderedMap()
    nextHopType.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address})
    nextHopType.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address})
    nextHopType.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", nextHopType.IfIndex})
    nextHopType.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", nextHopType.MacAddress})
    nextHopType.EntityData.Leafs.Append("out-interface-name", types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName})

    nextHopType.EntityData.YListKeys = []string {}

    return &(nextHopType.EntityData)
}

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xe"
    operations.EntityData.ParentYangName = "next-hop"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/next-hop/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

    return &(operations.EntityData)
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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/next-hop/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/next-hop/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/next-hop/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-cfg/ipv4-ingress-lsps/ipv4-ingress-lsp/path/next-hop/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    mplsStaticState.EntityData.AbsolutePath = "common-mpls-static:mpls-static/" + mplsStaticState.EntityData.SegmentPath
    mplsStaticState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsStaticState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsStaticState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsStaticState.EntityData.Children = types.NewOrderedMap()
    mplsStaticState.EntityData.Children.Append("label-switched-paths", types.YChild{"LabelSwitchedPaths", &mplsStaticState.LabelSwitchedPaths})
    mplsStaticState.EntityData.Leafs = types.NewOrderedMap()

    mplsStaticState.EntityData.YListKeys = []string {}

    return &(mplsStaticState.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths
// MPLS Static Label Switched Paths
type MplsStatic_MplsStaticState_LabelSwitchedPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LSPs list. The type is slice of
    // MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath.
    LabelSwitchedPath []*MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetEntityData() *types.CommonEntityData {
    labelSwitchedPaths.EntityData.YFilter = labelSwitchedPaths.YFilter
    labelSwitchedPaths.EntityData.YangName = "label-switched-paths"
    labelSwitchedPaths.EntityData.BundleName = "cisco_ios_xe"
    labelSwitchedPaths.EntityData.ParentYangName = "mpls-static-state"
    labelSwitchedPaths.EntityData.SegmentPath = "label-switched-paths"
    labelSwitchedPaths.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/" + labelSwitchedPaths.EntityData.SegmentPath
    labelSwitchedPaths.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelSwitchedPaths.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelSwitchedPaths.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelSwitchedPaths.EntityData.Children = types.NewOrderedMap()
    labelSwitchedPaths.EntityData.Children.Append("label-switched-path", types.YChild{"LabelSwitchedPath", nil})
    for i := range labelSwitchedPaths.LabelSwitchedPath {
        labelSwitchedPaths.EntityData.Children.Append(types.GetSegmentPath(labelSwitchedPaths.LabelSwitchedPath[i]), types.YChild{"LabelSwitchedPath", labelSwitchedPaths.LabelSwitchedPath[i]})
    }
    labelSwitchedPaths.EntityData.Leafs = types.NewOrderedMap()

    labelSwitchedPaths.EntityData.YListKeys = []string {}

    return &(labelSwitchedPaths.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath
// MPLS LSPs list
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. IP v4/v6 prefix. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
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
    labelSwitchedPath.EntityData.SegmentPath = "label-switched-path" + types.AddKeyToken(labelSwitchedPath.VrfName, "vrf-name") + types.AddKeyToken(labelSwitchedPath.Prefix, "prefix")
    labelSwitchedPath.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/" + labelSwitchedPath.EntityData.SegmentPath
    labelSwitchedPath.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelSwitchedPath.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelSwitchedPath.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelSwitchedPath.EntityData.Children = types.NewOrderedMap()
    labelSwitchedPath.EntityData.Children.Append("ingress-stats", types.YChild{"IngressStats", &labelSwitchedPath.IngressStats})
    labelSwitchedPath.EntityData.Children.Append("egress-stats", types.YChild{"EgressStats", &labelSwitchedPath.EgressStats})
    labelSwitchedPath.EntityData.Children.Append("path", types.YChild{"Path", &labelSwitchedPath.Path})
    labelSwitchedPath.EntityData.Leafs = types.NewOrderedMap()
    labelSwitchedPath.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", labelSwitchedPath.VrfName})
    labelSwitchedPath.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", labelSwitchedPath.Prefix})
    labelSwitchedPath.EntityData.Leafs.Append("name", types.YLeaf{"Name", labelSwitchedPath.Name})
    labelSwitchedPath.EntityData.Leafs.Append("in-label-value", types.YLeaf{"InLabelValue", labelSwitchedPath.InLabelValue})

    labelSwitchedPath.EntityData.YListKeys = []string {"VrfName", "Prefix"}

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
    ingressStats.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/" + ingressStats.EntityData.SegmentPath
    ingressStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ingressStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ingressStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ingressStats.EntityData.Children = types.NewOrderedMap()
    ingressStats.EntityData.Children.Append("stats", types.YChild{"Stats", &ingressStats.Stats})
    ingressStats.EntityData.Leafs = types.NewOrderedMap()

    ingressStats.EntityData.YListKeys = []string {}

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
    stats.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/ingress-stats/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", stats.Packets})
    stats.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", stats.Bytes})
    stats.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", stats.DroppedPackets})
    stats.EntityData.Leafs.Append("dropped-bytes", types.YLeaf{"DroppedBytes", stats.DroppedBytes})

    stats.EntityData.YListKeys = []string {}

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
    egressStats.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/" + egressStats.EntityData.SegmentPath
    egressStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    egressStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    egressStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    egressStats.EntityData.Children = types.NewOrderedMap()
    egressStats.EntityData.Children.Append("stats", types.YChild{"Stats", &egressStats.Stats})
    egressStats.EntityData.Leafs = types.NewOrderedMap()

    egressStats.EntityData.YListKeys = []string {}

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
    stats.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/egress-stats/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", stats.Packets})
    stats.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", stats.Bytes})
    stats.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", stats.DroppedPackets})
    stats.EntityData.Leafs.Append("dropped-bytes", types.YLeaf{"DroppedBytes", stats.DroppedBytes})

    stats.EntityData.YListKeys = []string {}

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
    NextHop []*MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop
}

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "label-switched-path"
    path.EntityData.SegmentPath = "path"
    path.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/" + path.EntityData.SegmentPath
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("operations", types.YChild{"Operations", &path.Operations})
    path.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range path.NextHop {
        path.EntityData.Children.Append(types.GetSegmentPath(path.NextHop[i]), types.YChild{"NextHop", path.NextHop[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("auto-protect", types.YLeaf{"AutoProtect", path.AutoProtect})

    path.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the nexthop. The type is interface{} with
    // range: 0..4294967295. This attribute is mandatory.
    Index interface{}

    // The forwarding path's hoptype. The type is Hoptype. This attribute is
    // mandatory.
    Type interface{}

    // Index of the nexthop that protects this nexthop. The type is interface{}
    // with range: 0..4294967295.
    ProtectedBy interface{}

    // The origin of this nexthop. The type is one of the following:
    // BgpRouteNexthopIsisRouteNexthopStaticNexthopOspfRouteNexthop.
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
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("next-hop-type", types.YChild{"NextHopType", &nextHop.NextHopType})
    nextHop.EntityData.Children.Append("operations", types.YChild{"Operations", &nextHop.Operations})
    nextHop.EntityData.Children.Append("nexthop-stats", types.YChild{"NexthopStats", &nextHop.NexthopStats})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})
    nextHop.EntityData.Leafs.Append("type", types.YLeaf{"Type", nextHop.Type})
    nextHop.EntityData.Leafs.Append("protected-by", types.YLeaf{"ProtectedBy", nextHop.ProtectedBy})
    nextHop.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", nextHop.Origin})

    nextHop.EntityData.YListKeys = []string {"Index"}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This attribute is mandatory.
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
    nextHopType.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/" + nextHopType.EntityData.SegmentPath
    nextHopType.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHopType.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHopType.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHopType.EntityData.Children = types.NewOrderedMap()
    nextHopType.EntityData.Leafs = types.NewOrderedMap()
    nextHopType.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", nextHopType.IfIndex})
    nextHopType.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHopType.Ipv4Address})
    nextHopType.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHopType.Ipv6Address})
    nextHopType.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", nextHopType.MacAddress})
    nextHopType.EntityData.Leafs.Append("out-interface-name", types.YLeaf{"OutInterfaceName", nextHopType.OutInterfaceName})

    nextHopType.EntityData.YListKeys = []string {}

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
    operations.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/" + operations.EntityData.SegmentPath
    operations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("swap", types.YChild{"Swap", &operations.Swap})
    operations.EntityData.Children.Append("push", types.YChild{"Push", &operations.Push})
    operations.EntityData.Leafs = types.NewOrderedMap()
    operations.EntityData.Leafs.Append("preserve", types.YLeaf{"Preserve", operations.Preserve})
    operations.EntityData.Leafs.Append("pop-and-forward", types.YLeaf{"PopAndForward", operations.PopAndForward})

    operations.EntityData.YListKeys = []string {}

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
    swap.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/operations/" + swap.EntityData.SegmentPath
    swap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    swap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    swap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    swap.EntityData.Children = types.NewOrderedMap()
    swap.EntityData.Children.Append("stack", types.YChild{"Stack", &swap.Stack})
    swap.EntityData.Leafs = types.NewOrderedMap()

    swap.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/operations/swap/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    push.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/operations/" + push.EntityData.SegmentPath
    push.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    push.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    push.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    push.EntityData.Children = types.NewOrderedMap()
    push.EntityData.Children.Append("stack", types.YChild{"Stack", &push.Stack})
    push.EntityData.Leafs = types.NewOrderedMap()

    push.EntityData.YListKeys = []string {}

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
    stack.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/operations/push/" + stack.EntityData.SegmentPath
    stack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", stack.LabelStack})

    stack.EntityData.YListKeys = []string {}

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
    nexthopStats.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/" + nexthopStats.EntityData.SegmentPath
    nexthopStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nexthopStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nexthopStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nexthopStats.EntityData.Children = types.NewOrderedMap()
    nexthopStats.EntityData.Children.Append("stats", types.YChild{"Stats", &nexthopStats.Stats})
    nexthopStats.EntityData.Leafs = types.NewOrderedMap()

    nexthopStats.EntityData.YListKeys = []string {}

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
    stats.EntityData.AbsolutePath = "common-mpls-static:mpls-static/mpls-static-state/label-switched-paths/label-switched-path/path/next-hop/nexthop-stats/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", stats.Packets})
    stats.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", stats.Bytes})
    stats.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", stats.DroppedPackets})
    stats.EntityData.Leafs.Append("dropped-bytes", types.YLeaf{"DroppedBytes", stats.DroppedBytes})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

