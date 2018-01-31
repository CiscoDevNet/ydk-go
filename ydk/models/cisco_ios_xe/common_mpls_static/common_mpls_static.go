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
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Static configuration commands.
    MplsStaticCfg MplsStatic_MplsStaticCfg

    // MPLS static operational data.
    MplsStaticState MplsStatic_MplsStaticState
}

func (mplsStatic *MplsStatic) GetFilter() yfilter.YFilter { return mplsStatic.YFilter }

func (mplsStatic *MplsStatic) SetFilter(yf yfilter.YFilter) { mplsStatic.YFilter = yf }

func (mplsStatic *MplsStatic) GetGoName(yname string) string {
    if yname == "mpls-static-cfg" { return "MplsStaticCfg" }
    if yname == "mpls-static-state" { return "MplsStaticState" }
    return ""
}

func (mplsStatic *MplsStatic) GetSegmentPath() string {
    return "common-mpls-static:mpls-static"
}

func (mplsStatic *MplsStatic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpls-static-cfg" {
        return &mplsStatic.MplsStaticCfg
    }
    if childYangName == "mpls-static-state" {
        return &mplsStatic.MplsStaticState
    }
    return nil
}

func (mplsStatic *MplsStatic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpls-static-cfg"] = &mplsStatic.MplsStaticCfg
    children["mpls-static-state"] = &mplsStatic.MplsStaticState
    return children
}

func (mplsStatic *MplsStatic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsStatic *MplsStatic) GetBundleName() string { return "cisco_ios_xe" }

func (mplsStatic *MplsStatic) GetYangName() string { return "mpls-static" }

func (mplsStatic *MplsStatic) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsStatic *MplsStatic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsStatic *MplsStatic) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsStatic *MplsStatic) SetParent(parent types.Entity) { mplsStatic.parent = parent }

func (mplsStatic *MplsStatic) GetParent() types.Entity { return mplsStatic.parent }

func (mplsStatic *MplsStatic) GetParentYangName() string { return "common-mpls-static" }

// MplsStatic_MplsStaticCfg
// MPLS Static configuration commands
type MplsStatic_MplsStaticCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of interfaces configured with mpls.
    Interfaces MplsStatic_MplsStaticCfg_Interfaces

    // The LSPs indexed by ipv4 prefix.
    Ipv4IngressLsps MplsStatic_MplsStaticCfg_Ipv4IngressLsps

    // The LSPs indexed by ipv6 prefix.
    Ipv6IngressLsps MplsStatic_MplsStaticCfg_Ipv6IngressLsps

    // The LSPs indexed by name.
    NamedLsps MplsStatic_MplsStaticCfg_NamedLsps

    // The LSPs indexed by in-label.
    InLabelLsps MplsStatic_MplsStaticCfg_InLabelLsps
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetFilter() yfilter.YFilter { return mplsStaticCfg.YFilter }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) SetFilter(yf yfilter.YFilter) { mplsStaticCfg.YFilter = yf }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "ipv4-ingress-lsps" { return "Ipv4IngressLsps" }
    if yname == "ipv6-ingress-lsps" { return "Ipv6IngressLsps" }
    if yname == "named-lsps" { return "NamedLsps" }
    if yname == "in-label-lsps" { return "InLabelLsps" }
    return ""
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetSegmentPath() string {
    return "mpls-static-cfg"
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &mplsStaticCfg.Interfaces
    }
    if childYangName == "ipv4-ingress-lsps" {
        return &mplsStaticCfg.Ipv4IngressLsps
    }
    if childYangName == "ipv6-ingress-lsps" {
        return &mplsStaticCfg.Ipv6IngressLsps
    }
    if childYangName == "named-lsps" {
        return &mplsStaticCfg.NamedLsps
    }
    if childYangName == "in-label-lsps" {
        return &mplsStaticCfg.InLabelLsps
    }
    return nil
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &mplsStaticCfg.Interfaces
    children["ipv4-ingress-lsps"] = &mplsStaticCfg.Ipv4IngressLsps
    children["ipv6-ingress-lsps"] = &mplsStaticCfg.Ipv6IngressLsps
    children["named-lsps"] = &mplsStaticCfg.NamedLsps
    children["in-label-lsps"] = &mplsStaticCfg.InLabelLsps
    return children
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetBundleName() string { return "cisco_ios_xe" }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetYangName() string { return "mpls-static-cfg" }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) SetParent(parent types.Entity) { mplsStaticCfg.parent = parent }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetParent() types.Entity { return mplsStaticCfg.parent }

func (mplsStaticCfg *MplsStatic_MplsStaticCfg) GetParentYangName() string { return "mpls-static" }

// MplsStatic_MplsStaticCfg_Interfaces
// The list of interfaces configured with mpls
type MplsStatic_MplsStaticCfg_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of interfaces that can be enabled under mpls static. The type is slice
    // of MplsStatic_MplsStaticCfg_Interfaces_Interface.
    Interface []MplsStatic_MplsStaticCfg_Interfaces_Interface
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetBundleName() string { return "cisco_ios_xe" }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsStatic_MplsStaticCfg_Interfaces) GetParentYangName() string { return "mpls-static-cfg" }

// MplsStatic_MplsStaticCfg_Interfaces_Interface
// List of interfaces that can be enabled under
// mpls static
type MplsStatic_MplsStaticCfg_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name This attribute is mandatory.
    Name interface{}

    // Interface Enabled boolean. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    Enabled interface{}
}

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["enabled"] = self.Enabled
    return leafs
}

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xe" }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsStatic_MplsStaticCfg_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps
// The LSPs indexed by ipv4 prefix
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Static IPv4 Label Switched Path Configuration at Ingress. The type is
    // slice of MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp.
    Ipv4IngressLsp []MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetFilter() yfilter.YFilter { return ipv4IngressLsps.YFilter }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) SetFilter(yf yfilter.YFilter) { ipv4IngressLsps.YFilter = yf }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetGoName(yname string) string {
    if yname == "ipv4-ingress-lsp" { return "Ipv4IngressLsp" }
    return ""
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetSegmentPath() string {
    return "ipv4-ingress-lsps"
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-ingress-lsp" {
        for _, c := range ipv4IngressLsps.Ipv4IngressLsp {
            if ipv4IngressLsps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp{}
        ipv4IngressLsps.Ipv4IngressLsp = append(ipv4IngressLsps.Ipv4IngressLsp, child)
        return &ipv4IngressLsps.Ipv4IngressLsp[len(ipv4IngressLsps.Ipv4IngressLsp)-1]
    }
    return nil
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4IngressLsps.Ipv4IngressLsp {
        children[ipv4IngressLsps.Ipv4IngressLsp[i].GetSegmentPath()] = &ipv4IngressLsps.Ipv4IngressLsp[i]
    }
    return children
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetBundleName() string { return "cisco_ios_xe" }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetYangName() string { return "ipv4-ingress-lsps" }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) SetParent(parent types.Entity) { ipv4IngressLsps.parent = parent }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetParent() types.Entity { return ipv4IngressLsps.parent }

func (ipv4IngressLsps *MplsStatic_MplsStaticCfg_Ipv4IngressLsps) GetParentYangName() string { return "mpls-static-cfg" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp
// MPLS Static IPv4 Label Switched
// Path Configuration at Ingress
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. IPv4 prefix of packets that will ingress on this
    // LSP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Prefix interface{}

    // Value of the local label. Optional for ingress. The type is one of the
    // following types: int with range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // Name of the LSP. The type is string.
    Name interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetFilter() yfilter.YFilter { return ipv4IngressLsp.YFilter }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) SetFilter(yf yfilter.YFilter) { ipv4IngressLsp.YFilter = yf }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "prefix" { return "Prefix" }
    if yname == "in-label" { return "InLabel" }
    if yname == "name" { return "Name" }
    if yname == "path" { return "Path" }
    return ""
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetSegmentPath() string {
    return "ipv4-ingress-lsp" + "[vrf-name='" + fmt.Sprintf("%v", ipv4IngressLsp.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", ipv4IngressLsp.Prefix) + "']"
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        return &ipv4IngressLsp.Path
    }
    return nil
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path"] = &ipv4IngressLsp.Path
    return children
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv4IngressLsp.VrfName
    leafs["prefix"] = ipv4IngressLsp.Prefix
    leafs["in-label"] = ipv4IngressLsp.InLabel
    leafs["name"] = ipv4IngressLsp.Name
    return leafs
}

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetBundleName() string { return "cisco_ios_xe" }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetYangName() string { return "ipv4-ingress-lsp" }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) SetParent(parent types.Entity) { ipv4IngressLsp.parent = parent }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetParent() types.Entity { return ipv4IngressLsp.parent }

func (ipv4IngressLsp *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp) GetParentYangName() string { return "ipv4-ingress-lsps" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enables automatic protection if true. The type is bool. The default value
    // is false.
    AutoProtect interface{}

    // The incoming label processing.
    Operations MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations

    // next-hops list. The type is slice of
    // MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop.
    NextHop []MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetGoName(yname string) string {
    if yname == "auto-protect" { return "AutoProtect" }
    if yname == "operations" { return "Operations" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetSegmentPath() string {
    return "path"
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operations" {
        return &path.Operations
    }
    if childYangName == "next-hop" {
        for _, c := range path.NextHop {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop{}
        path.NextHop = append(path.NextHop, child)
        return &path.NextHop[len(path.NextHop)-1]
    }
    return nil
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["operations"] = &path.Operations
    for i := range path.NextHop {
        children[path.NextHop[i].GetSegmentPath()] = &path.NextHop[i]
    }
    return children
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-protect"] = path.AutoProtect
    return leafs
}

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetBundleName() string { return "cisco_ios_xe" }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetYangName() string { return "path" }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path) GetParentYangName() string { return "ipv4-ingress-lsp" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "type" { return "Type" }
    if yname == "protected-by" { return "ProtectedBy" }
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "operations" { return "Operations" }
    return ""
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetSegmentPath() string {
    return "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop-type" {
        return &nextHop.NextHopType
    }
    if childYangName == "operations" {
        return &nextHop.Operations
    }
    return nil
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop-type"] = &nextHop.NextHopType
    children["operations"] = &nextHop.Operations
    return children
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = nextHop.Index
    leafs["type"] = nextHop.Type
    leafs["protected-by"] = nextHop.ProtectedBy
    return leafs
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetBundleName() string { return "cisco_ios_xe" }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    IfIndex interface{}

    // IPv6 Address of the nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv6Address interface{}

    // MAC address of the nexthop. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This attribute is mandatory.
    MacAddress interface{}

    // IPv4 Address of the nexthop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ipv4Address interface{}

    // Name of the outgoing interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    OutInterfaceName interface{}
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetFilter() yfilter.YFilter { return nextHopType.YFilter }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) SetFilter(yf yfilter.YFilter) { nextHopType.YFilter = yf }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetGoName(yname string) string {
    if yname == "if-index" { return "IfIndex" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "out-interface-name" { return "OutInterfaceName" }
    return ""
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetSegmentPath() string {
    return "next-hop-type"
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-index"] = nextHopType.IfIndex
    leafs["ipv6-address"] = nextHopType.Ipv6Address
    leafs["mac-address"] = nextHopType.MacAddress
    leafs["ipv4-address"] = nextHopType.Ipv4Address
    leafs["out-interface-name"] = nextHopType.OutInterfaceName
    return leafs
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetBundleName() string { return "cisco_ios_xe" }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetYangName() string { return "next-hop-type" }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) SetParent(parent types.Entity) { nextHopType.parent = parent }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetParent() types.Entity { return nextHopType.parent }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_NextHopType) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pop the incoming label and forward. The type is interface{}. This attribute
    // is mandatory.
    PopAndForward interface{}

    // preserve incoming label stack. The type is interface{}. This attribute is
    // mandatory.
    Preserve interface{}

    // Push outgoing label stack.
    Push MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push

    // Push outgoing label stack.
    Swap MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetGoName(yname string) string {
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "preserve" { return "Preserve" }
    if yname == "push" { return "Push" }
    if yname == "swap" { return "Swap" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "push" {
        return &operations.Push
    }
    if childYangName == "swap" {
        return &operations.Swap
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["push"] = &operations.Push
    children["swap"] = &operations.Swap
    return children
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pop-and-forward"] = operations.PopAndForward
    leafs["preserve"] = operations.Preserve
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv4IngressLsps_Ipv4IngressLsp_Path_NextHop_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps
// The LSPs indexed by ipv6 prefix
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Static IPv6 Label Switched Path Configuration at Ingress. The type is
    // slice of MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp.
    Ipv6IngressLsp []MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetFilter() yfilter.YFilter { return ipv6IngressLsps.YFilter }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) SetFilter(yf yfilter.YFilter) { ipv6IngressLsps.YFilter = yf }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetGoName(yname string) string {
    if yname == "ipv6-ingress-lsp" { return "Ipv6IngressLsp" }
    return ""
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetSegmentPath() string {
    return "ipv6-ingress-lsps"
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-ingress-lsp" {
        for _, c := range ipv6IngressLsps.Ipv6IngressLsp {
            if ipv6IngressLsps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp{}
        ipv6IngressLsps.Ipv6IngressLsp = append(ipv6IngressLsps.Ipv6IngressLsp, child)
        return &ipv6IngressLsps.Ipv6IngressLsp[len(ipv6IngressLsps.Ipv6IngressLsp)-1]
    }
    return nil
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6IngressLsps.Ipv6IngressLsp {
        children[ipv6IngressLsps.Ipv6IngressLsp[i].GetSegmentPath()] = &ipv6IngressLsps.Ipv6IngressLsp[i]
    }
    return children
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetYangName() string { return "ipv6-ingress-lsps" }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) SetParent(parent types.Entity) { ipv6IngressLsps.parent = parent }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetParent() types.Entity { return ipv6IngressLsps.parent }

func (ipv6IngressLsps *MplsStatic_MplsStaticCfg_Ipv6IngressLsps) GetParentYangName() string { return "mpls-static-cfg" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp
// MPLS Static IPv6 Label Switched Path
// Configuration at Ingress
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetFilter() yfilter.YFilter { return ipv6IngressLsp.YFilter }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) SetFilter(yf yfilter.YFilter) { ipv6IngressLsp.YFilter = yf }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "prefix" { return "Prefix" }
    if yname == "name" { return "Name" }
    if yname == "in-label" { return "InLabel" }
    if yname == "path" { return "Path" }
    return ""
}

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetSegmentPath() string {
    return "ipv6-ingress-lsp" + "[vrf-name='" + fmt.Sprintf("%v", ipv6IngressLsp.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", ipv6IngressLsp.Prefix) + "']"
}

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        return &ipv6IngressLsp.Path
    }
    return nil
}

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path"] = &ipv6IngressLsp.Path
    return children
}

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv6IngressLsp.VrfName
    leafs["prefix"] = ipv6IngressLsp.Prefix
    leafs["name"] = ipv6IngressLsp.Name
    leafs["in-label"] = ipv6IngressLsp.InLabel
    return leafs
}

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetBundleName() string { return "cisco_ios_xe" }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetYangName() string { return "ipv6-ingress-lsp" }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) SetParent(parent types.Entity) { ipv6IngressLsp.parent = parent }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetParent() types.Entity { return ipv6IngressLsp.parent }

func (ipv6IngressLsp *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp) GetParentYangName() string { return "ipv6-ingress-lsps" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path struct {
    parent types.Entity
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

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetGoName(yname string) string {
    if yname == "auto-protect" { return "AutoProtect" }
    if yname == "operations" { return "Operations" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetSegmentPath() string {
    return "path"
}

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operations" {
        return &path.Operations
    }
    if childYangName == "next-hop" {
        for _, c := range path.NextHop {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop{}
        path.NextHop = append(path.NextHop, child)
        return &path.NextHop[len(path.NextHop)-1]
    }
    return nil
}

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["operations"] = &path.Operations
    for i := range path.NextHop {
        children[path.NextHop[i].GetSegmentPath()] = &path.NextHop[i]
    }
    return children
}

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-protect"] = path.AutoProtect
    return leafs
}

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetBundleName() string { return "cisco_ios_xe" }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetYangName() string { return "path" }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path) GetParentYangName() string { return "ipv6-ingress-lsp" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "type" { return "Type" }
    if yname == "protected-by" { return "ProtectedBy" }
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "operations" { return "Operations" }
    return ""
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetSegmentPath() string {
    return "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop-type" {
        return &nextHop.NextHopType
    }
    if childYangName == "operations" {
        return &nextHop.Operations
    }
    return nil
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop-type"] = &nextHop.NextHopType
    children["operations"] = &nextHop.Operations
    return children
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = nextHop.Index
    leafs["type"] = nextHop.Type
    leafs["protected-by"] = nextHop.ProtectedBy
    return leafs
}

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetBundleName() string { return "cisco_ios_xe" }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType struct {
    parent types.Entity
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

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetFilter() yfilter.YFilter { return nextHopType.YFilter }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) SetFilter(yf yfilter.YFilter) { nextHopType.YFilter = yf }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetGoName(yname string) string {
    if yname == "if-index" { return "IfIndex" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "out-interface-name" { return "OutInterfaceName" }
    return ""
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetSegmentPath() string {
    return "next-hop-type"
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-index"] = nextHopType.IfIndex
    leafs["ipv4-address"] = nextHopType.Ipv4Address
    leafs["ipv6-address"] = nextHopType.Ipv6Address
    leafs["mac-address"] = nextHopType.MacAddress
    leafs["out-interface-name"] = nextHopType.OutInterfaceName
    return leafs
}

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetBundleName() string { return "cisco_ios_xe" }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetYangName() string { return "next-hop-type" }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) SetParent(parent types.Entity) { nextHopType.parent = parent }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetParent() types.Entity { return nextHopType.parent }

func (nextHopType *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_NextHopType) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_Ipv6IngressLsps_Ipv6IngressLsp_Path_NextHop_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticCfg_NamedLsps
// The LSPs indexed by name
type MplsStatic_MplsStaticCfg_NamedLsps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Static Label Switched Path Configuration. The LSPs in this list are
    // referenced by a string name. The LSPs may be ingress/egress/crossconnect,
    // may have v4/v6 prefixes and may be associated with any VRF. The other
    // specialized lists above are for implemetations that are keyed on prefixes
    // or in-labels instead of the LSP name. The type is slice of
    // MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp.
    NamedLsp []MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetFilter() yfilter.YFilter { return namedLsps.YFilter }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) SetFilter(yf yfilter.YFilter) { namedLsps.YFilter = yf }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetGoName(yname string) string {
    if yname == "named-lsp" { return "NamedLsp" }
    return ""
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetSegmentPath() string {
    return "named-lsps"
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "named-lsp" {
        for _, c := range namedLsps.NamedLsp {
            if namedLsps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp{}
        namedLsps.NamedLsp = append(namedLsps.NamedLsp, child)
        return &namedLsps.NamedLsp[len(namedLsps.NamedLsp)-1]
    }
    return nil
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range namedLsps.NamedLsp {
        children[namedLsps.NamedLsp[i].GetSegmentPath()] = &namedLsps.NamedLsp[i]
    }
    return children
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetBundleName() string { return "cisco_ios_xe" }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetYangName() string { return "named-lsps" }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) SetParent(parent types.Entity) { namedLsps.parent = parent }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetParent() types.Entity { return namedLsps.parent }

func (namedLsps *MplsStatic_MplsStaticCfg_NamedLsps) GetParentYangName() string { return "mpls-static-cfg" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp
// MPLS Static Label Switched Path Configuration.
// The LSPs in this list are referenced by a string name.
// The LSPs may be ingress/egress/crossconnect,
// may have v4/v6 prefixes and may be associated with any
// VRF. The other specialized lists above are for
// implemetations that are keyed on prefixes or in-labels
// instead of the LSP name.
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetFilter() yfilter.YFilter { return namedLsp.YFilter }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) SetFilter(yf yfilter.YFilter) { namedLsp.YFilter = yf }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "name" { return "Name" }
    if yname == "lsp-type" { return "LspType" }
    if yname == "in-label" { return "InLabel" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    if yname == "path" { return "Path" }
    return ""
}

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetSegmentPath() string {
    return "named-lsp" + "[vrf-name='" + fmt.Sprintf("%v", namedLsp.VrfName) + "']" + "[name='" + fmt.Sprintf("%v", namedLsp.Name) + "']"
}

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        return &namedLsp.Path
    }
    return nil
}

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path"] = &namedLsp.Path
    return children
}

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = namedLsp.VrfName
    leafs["name"] = namedLsp.Name
    leafs["lsp-type"] = namedLsp.LspType
    leafs["in-label"] = namedLsp.InLabel
    leafs["ipv4-prefix"] = namedLsp.Ipv4Prefix
    leafs["ipv6-prefix"] = namedLsp.Ipv6Prefix
    return leafs
}

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetBundleName() string { return "cisco_ios_xe" }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetYangName() string { return "named-lsp" }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) SetParent(parent types.Entity) { namedLsp.parent = parent }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetParent() types.Entity { return namedLsp.parent }

func (namedLsp *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp) GetParentYangName() string { return "named-lsps" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path struct {
    parent types.Entity
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

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetGoName(yname string) string {
    if yname == "auto-protect" { return "AutoProtect" }
    if yname == "operations" { return "Operations" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetSegmentPath() string {
    return "path"
}

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operations" {
        return &path.Operations
    }
    if childYangName == "next-hop" {
        for _, c := range path.NextHop {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop{}
        path.NextHop = append(path.NextHop, child)
        return &path.NextHop[len(path.NextHop)-1]
    }
    return nil
}

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["operations"] = &path.Operations
    for i := range path.NextHop {
        children[path.NextHop[i].GetSegmentPath()] = &path.NextHop[i]
    }
    return children
}

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-protect"] = path.AutoProtect
    return leafs
}

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetBundleName() string { return "cisco_ios_xe" }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetYangName() string { return "path" }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path) GetParentYangName() string { return "named-lsp" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "type" { return "Type" }
    if yname == "protected-by" { return "ProtectedBy" }
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "operations" { return "Operations" }
    return ""
}

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetSegmentPath() string {
    return "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
}

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop-type" {
        return &nextHop.NextHopType
    }
    if childYangName == "operations" {
        return &nextHop.Operations
    }
    return nil
}

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop-type"] = &nextHop.NextHopType
    children["operations"] = &nextHop.Operations
    return children
}

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = nextHop.Index
    leafs["type"] = nextHop.Type
    leafs["protected-by"] = nextHop.ProtectedBy
    return leafs
}

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetBundleName() string { return "cisco_ios_xe" }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType struct {
    parent types.Entity
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

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetFilter() yfilter.YFilter { return nextHopType.YFilter }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) SetFilter(yf yfilter.YFilter) { nextHopType.YFilter = yf }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetGoName(yname string) string {
    if yname == "if-index" { return "IfIndex" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "out-interface-name" { return "OutInterfaceName" }
    return ""
}

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetSegmentPath() string {
    return "next-hop-type"
}

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-index"] = nextHopType.IfIndex
    leafs["ipv4-address"] = nextHopType.Ipv4Address
    leafs["ipv6-address"] = nextHopType.Ipv6Address
    leafs["mac-address"] = nextHopType.MacAddress
    leafs["out-interface-name"] = nextHopType.OutInterfaceName
    return leafs
}

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetBundleName() string { return "cisco_ios_xe" }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetYangName() string { return "next-hop-type" }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) SetParent(parent types.Entity) { nextHopType.parent = parent }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetParent() types.Entity { return nextHopType.parent }

func (nextHopType *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_NextHopType) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_NamedLsps_NamedLsp_Path_NextHop_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticCfg_InLabelLsps
// The LSPs indexed by in-label
type MplsStatic_MplsStaticCfg_InLabelLsps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Non-ingress MPLS Static LSPs, keyed on the incoming label. The type is
    // slice of MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp.
    InLabelLsp []MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetFilter() yfilter.YFilter { return inLabelLsps.YFilter }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) SetFilter(yf yfilter.YFilter) { inLabelLsps.YFilter = yf }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetGoName(yname string) string {
    if yname == "in-label-lsp" { return "InLabelLsp" }
    return ""
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetSegmentPath() string {
    return "in-label-lsps"
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "in-label-lsp" {
        for _, c := range inLabelLsps.InLabelLsp {
            if inLabelLsps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp{}
        inLabelLsps.InLabelLsp = append(inLabelLsps.InLabelLsp, child)
        return &inLabelLsps.InLabelLsp[len(inLabelLsps.InLabelLsp)-1]
    }
    return nil
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inLabelLsps.InLabelLsp {
        children[inLabelLsps.InLabelLsp[i].GetSegmentPath()] = &inLabelLsps.InLabelLsp[i]
    }
    return children
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetBundleName() string { return "cisco_ios_xe" }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetYangName() string { return "in-label-lsps" }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) SetParent(parent types.Entity) { inLabelLsps.parent = parent }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetParent() types.Entity { return inLabelLsps.parent }

func (inLabelLsps *MplsStatic_MplsStaticCfg_InLabelLsps) GetParentYangName() string { return "mpls-static-cfg" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp
// Non-ingress MPLS Static LSPs,
// keyed on the incoming label
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string.
    VrfName interface{}

    // This attribute is a key. Value of the local label. The type is one of the
    // following types: int with range: 16..1048575, or enumeration IetfMplsLabel.
    InLabel interface{}

    // Fowarding path.
    Path MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path
}

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetFilter() yfilter.YFilter { return inLabelLsp.YFilter }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) SetFilter(yf yfilter.YFilter) { inLabelLsp.YFilter = yf }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "in-label" { return "InLabel" }
    if yname == "path" { return "Path" }
    return ""
}

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetSegmentPath() string {
    return "in-label-lsp" + "[vrf-name='" + fmt.Sprintf("%v", inLabelLsp.VrfName) + "']" + "[in-label='" + fmt.Sprintf("%v", inLabelLsp.InLabel) + "']"
}

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        return &inLabelLsp.Path
    }
    return nil
}

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path"] = &inLabelLsp.Path
    return children
}

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = inLabelLsp.VrfName
    leafs["in-label"] = inLabelLsp.InLabel
    return leafs
}

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetBundleName() string { return "cisco_ios_xe" }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetYangName() string { return "in-label-lsp" }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) SetParent(parent types.Entity) { inLabelLsp.parent = parent }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetParent() types.Entity { return inLabelLsp.parent }

func (inLabelLsp *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp) GetParentYangName() string { return "in-label-lsps" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path
// Fowarding path
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path struct {
    parent types.Entity
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

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetGoName(yname string) string {
    if yname == "auto-protect" { return "AutoProtect" }
    if yname == "operations" { return "Operations" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetSegmentPath() string {
    return "path"
}

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operations" {
        return &path.Operations
    }
    if childYangName == "next-hop" {
        for _, c := range path.NextHop {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop{}
        path.NextHop = append(path.NextHop, child)
        return &path.NextHop[len(path.NextHop)-1]
    }
    return nil
}

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["operations"] = &path.Operations
    for i := range path.NextHop {
        children[path.NextHop[i].GetSegmentPath()] = &path.NextHop[i]
    }
    return children
}

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-protect"] = path.AutoProtect
    return leafs
}

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetBundleName() string { return "cisco_ios_xe" }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetYangName() string { return "path" }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path) GetParentYangName() string { return "in-label-lsp" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "type" { return "Type" }
    if yname == "protected-by" { return "ProtectedBy" }
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "operations" { return "Operations" }
    return ""
}

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetSegmentPath() string {
    return "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
}

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop-type" {
        return &nextHop.NextHopType
    }
    if childYangName == "operations" {
        return &nextHop.Operations
    }
    return nil
}

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop-type"] = &nextHop.NextHopType
    children["operations"] = &nextHop.Operations
    return children
}

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = nextHop.Index
    leafs["type"] = nextHop.Type
    leafs["protected-by"] = nextHop.ProtectedBy
    return leafs
}

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetBundleName() string { return "cisco_ios_xe" }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType struct {
    parent types.Entity
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

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetFilter() yfilter.YFilter { return nextHopType.YFilter }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) SetFilter(yf yfilter.YFilter) { nextHopType.YFilter = yf }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetGoName(yname string) string {
    if yname == "if-index" { return "IfIndex" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "out-interface-name" { return "OutInterfaceName" }
    return ""
}

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetSegmentPath() string {
    return "next-hop-type"
}

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-index"] = nextHopType.IfIndex
    leafs["ipv4-address"] = nextHopType.Ipv4Address
    leafs["ipv6-address"] = nextHopType.Ipv6Address
    leafs["mac-address"] = nextHopType.MacAddress
    leafs["out-interface-name"] = nextHopType.OutInterfaceName
    return leafs
}

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetBundleName() string { return "cisco_ios_xe" }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetYangName() string { return "next-hop-type" }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) SetParent(parent types.Entity) { nextHopType.parent = parent }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetParent() types.Entity { return nextHopType.parent }

func (nextHopType *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_NextHopType) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticCfg_InLabelLsps_InLabelLsp_Path_NextHop_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticState
// MPLS static operational data
type MplsStatic_MplsStaticState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Static Label Switched Paths.
    LabelSwitchedPaths MplsStatic_MplsStaticState_LabelSwitchedPaths
}

func (mplsStaticState *MplsStatic_MplsStaticState) GetFilter() yfilter.YFilter { return mplsStaticState.YFilter }

func (mplsStaticState *MplsStatic_MplsStaticState) SetFilter(yf yfilter.YFilter) { mplsStaticState.YFilter = yf }

func (mplsStaticState *MplsStatic_MplsStaticState) GetGoName(yname string) string {
    if yname == "label-switched-paths" { return "LabelSwitchedPaths" }
    return ""
}

func (mplsStaticState *MplsStatic_MplsStaticState) GetSegmentPath() string {
    return "mpls-static-state"
}

func (mplsStaticState *MplsStatic_MplsStaticState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-paths" {
        return &mplsStaticState.LabelSwitchedPaths
    }
    return nil
}

func (mplsStaticState *MplsStatic_MplsStaticState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-switched-paths"] = &mplsStaticState.LabelSwitchedPaths
    return children
}

func (mplsStaticState *MplsStatic_MplsStaticState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsStaticState *MplsStatic_MplsStaticState) GetBundleName() string { return "cisco_ios_xe" }

func (mplsStaticState *MplsStatic_MplsStaticState) GetYangName() string { return "mpls-static-state" }

func (mplsStaticState *MplsStatic_MplsStaticState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsStaticState *MplsStatic_MplsStaticState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsStaticState *MplsStatic_MplsStaticState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsStaticState *MplsStatic_MplsStaticState) SetParent(parent types.Entity) { mplsStaticState.parent = parent }

func (mplsStaticState *MplsStatic_MplsStaticState) GetParent() types.Entity { return mplsStaticState.parent }

func (mplsStaticState *MplsStatic_MplsStaticState) GetParentYangName() string { return "mpls-static" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths
// MPLS Static Label Switched Paths
type MplsStatic_MplsStaticState_LabelSwitchedPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LSPs list. The type is slice of
    // MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath.
    LabelSwitchedPath []MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetFilter() yfilter.YFilter { return labelSwitchedPaths.YFilter }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) SetFilter(yf yfilter.YFilter) { labelSwitchedPaths.YFilter = yf }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetGoName(yname string) string {
    if yname == "label-switched-path" { return "LabelSwitchedPath" }
    return ""
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetSegmentPath() string {
    return "label-switched-paths"
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-path" {
        for _, c := range labelSwitchedPaths.LabelSwitchedPath {
            if labelSwitchedPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath{}
        labelSwitchedPaths.LabelSwitchedPath = append(labelSwitchedPaths.LabelSwitchedPath, child)
        return &labelSwitchedPaths.LabelSwitchedPath[len(labelSwitchedPaths.LabelSwitchedPath)-1]
    }
    return nil
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range labelSwitchedPaths.LabelSwitchedPath {
        children[labelSwitchedPaths.LabelSwitchedPath[i].GetSegmentPath()] = &labelSwitchedPaths.LabelSwitchedPath[i]
    }
    return children
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetBundleName() string { return "cisco_ios_xe" }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetYangName() string { return "label-switched-paths" }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) SetParent(parent types.Entity) { labelSwitchedPaths.parent = parent }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetParent() types.Entity { return labelSwitchedPaths.parent }

func (labelSwitchedPaths *MplsStatic_MplsStaticState_LabelSwitchedPaths) GetParentYangName() string { return "mpls-static-state" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath
// MPLS LSPs list
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetFilter() yfilter.YFilter { return labelSwitchedPath.YFilter }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) SetFilter(yf yfilter.YFilter) { labelSwitchedPath.YFilter = yf }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "prefix" { return "Prefix" }
    if yname == "name" { return "Name" }
    if yname == "in-label-value" { return "InLabelValue" }
    if yname == "ingress-stats" { return "IngressStats" }
    if yname == "egress-stats" { return "EgressStats" }
    if yname == "path" { return "Path" }
    return ""
}

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetSegmentPath() string {
    return "label-switched-path" + "[vrf-name='" + fmt.Sprintf("%v", labelSwitchedPath.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", labelSwitchedPath.Prefix) + "']"
}

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ingress-stats" {
        return &labelSwitchedPath.IngressStats
    }
    if childYangName == "egress-stats" {
        return &labelSwitchedPath.EgressStats
    }
    if childYangName == "path" {
        return &labelSwitchedPath.Path
    }
    return nil
}

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ingress-stats"] = &labelSwitchedPath.IngressStats
    children["egress-stats"] = &labelSwitchedPath.EgressStats
    children["path"] = &labelSwitchedPath.Path
    return children
}

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = labelSwitchedPath.VrfName
    leafs["prefix"] = labelSwitchedPath.Prefix
    leafs["name"] = labelSwitchedPath.Name
    leafs["in-label-value"] = labelSwitchedPath.InLabelValue
    return leafs
}

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetBundleName() string { return "cisco_ios_xe" }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetYangName() string { return "label-switched-path" }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) SetParent(parent types.Entity) { labelSwitchedPath.parent = parent }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetParent() types.Entity { return labelSwitchedPath.parent }

func (labelSwitchedPath *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath) GetParentYangName() string { return "label-switched-paths" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats
// ingress stats
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics.
    Stats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats
}

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetFilter() yfilter.YFilter { return ingressStats.YFilter }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) SetFilter(yf yfilter.YFilter) { ingressStats.YFilter = yf }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetGoName(yname string) string {
    if yname == "stats" { return "Stats" }
    return ""
}

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetSegmentPath() string {
    return "ingress-stats"
}

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &ingressStats.Stats
    }
    return nil
}

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &ingressStats.Stats
    return children
}

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetBundleName() string { return "cisco_ios_xe" }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetYangName() string { return "ingress-stats" }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) SetParent(parent types.Entity) { ingressStats.parent = parent }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetParent() types.Entity { return ingressStats.parent }

func (ingressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats
// Statistics
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats struct {
    parent types.Entity
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

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    if yname == "dropped-bytes" { return "DroppedBytes" }
    return ""
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = stats.Packets
    leafs["bytes"] = stats.Bytes
    leafs["dropped-packets"] = stats.DroppedPackets
    leafs["dropped-bytes"] = stats.DroppedBytes
    return leafs
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetBundleName() string { return "cisco_ios_xe" }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetYangName() string { return "stats" }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetParent() types.Entity { return stats.parent }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_IngressStats_Stats) GetParentYangName() string { return "ingress-stats" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats
// egress stats
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics.
    Stats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats
}

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetFilter() yfilter.YFilter { return egressStats.YFilter }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) SetFilter(yf yfilter.YFilter) { egressStats.YFilter = yf }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetGoName(yname string) string {
    if yname == "stats" { return "Stats" }
    return ""
}

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetSegmentPath() string {
    return "egress-stats"
}

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &egressStats.Stats
    }
    return nil
}

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &egressStats.Stats
    return children
}

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetBundleName() string { return "cisco_ios_xe" }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetYangName() string { return "egress-stats" }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) SetParent(parent types.Entity) { egressStats.parent = parent }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetParent() types.Entity { return egressStats.parent }

func (egressStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats
// Statistics
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats struct {
    parent types.Entity
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

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    if yname == "dropped-bytes" { return "DroppedBytes" }
    return ""
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = stats.Packets
    leafs["bytes"] = stats.Bytes
    leafs["dropped-packets"] = stats.DroppedPackets
    leafs["dropped-bytes"] = stats.DroppedBytes
    return leafs
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetBundleName() string { return "cisco_ios_xe" }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetYangName() string { return "stats" }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetParent() types.Entity { return stats.parent }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_EgressStats_Stats) GetParentYangName() string { return "egress-stats" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path
// Fowarding path
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path struct {
    parent types.Entity
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

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetGoName(yname string) string {
    if yname == "auto-protect" { return "AutoProtect" }
    if yname == "operations" { return "Operations" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetSegmentPath() string {
    return "path"
}

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operations" {
        return &path.Operations
    }
    if childYangName == "next-hop" {
        for _, c := range path.NextHop {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop{}
        path.NextHop = append(path.NextHop, child)
        return &path.NextHop[len(path.NextHop)-1]
    }
    return nil
}

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["operations"] = &path.Operations
    for i := range path.NextHop {
        children[path.NextHop[i].GetSegmentPath()] = &path.NextHop[i]
    }
    return children
}

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-protect"] = path.AutoProtect
    return leafs
}

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetBundleName() string { return "cisco_ios_xe" }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetYangName() string { return "path" }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations
// The incoming label processing
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop
// next-hops list
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "type" { return "Type" }
    if yname == "protected-by" { return "ProtectedBy" }
    if yname == "origin" { return "Origin" }
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "operations" { return "Operations" }
    if yname == "nexthop-stats" { return "NexthopStats" }
    return ""
}

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetSegmentPath() string {
    return "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
}

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop-type" {
        return &nextHop.NextHopType
    }
    if childYangName == "operations" {
        return &nextHop.Operations
    }
    if childYangName == "nexthop-stats" {
        return &nextHop.NexthopStats
    }
    return nil
}

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop-type"] = &nextHop.NextHopType
    children["operations"] = &nextHop.Operations
    children["nexthop-stats"] = &nextHop.NexthopStats
    return children
}

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = nextHop.Index
    leafs["type"] = nextHop.Type
    leafs["protected-by"] = nextHop.ProtectedBy
    leafs["origin"] = nextHop.Origin
    return leafs
}

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetBundleName() string { return "cisco_ios_xe" }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop) GetParentYangName() string { return "path" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType
// Next-hop
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType struct {
    parent types.Entity
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

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetFilter() yfilter.YFilter { return nextHopType.YFilter }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) SetFilter(yf yfilter.YFilter) { nextHopType.YFilter = yf }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetGoName(yname string) string {
    if yname == "if-index" { return "IfIndex" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "out-interface-name" { return "OutInterfaceName" }
    return ""
}

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetSegmentPath() string {
    return "next-hop-type"
}

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-index"] = nextHopType.IfIndex
    leafs["ipv4-address"] = nextHopType.Ipv4Address
    leafs["ipv6-address"] = nextHopType.Ipv6Address
    leafs["mac-address"] = nextHopType.MacAddress
    leafs["out-interface-name"] = nextHopType.OutInterfaceName
    return leafs
}

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetBundleName() string { return "cisco_ios_xe" }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetYangName() string { return "next-hop-type" }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) SetParent(parent types.Entity) { nextHopType.parent = parent }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetParent() types.Entity { return nextHopType.parent }

func (nextHopType *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NextHopType) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations
// The incoming label processing
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations struct {
    parent types.Entity
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

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetGoName(yname string) string {
    if yname == "preserve" { return "Preserve" }
    if yname == "pop-and-forward" { return "PopAndForward" }
    if yname == "swap" { return "Swap" }
    if yname == "push" { return "Push" }
    return ""
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "swap" {
        return &operations.Swap
    }
    if childYangName == "push" {
        return &operations.Push
    }
    return nil
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["swap"] = &operations.Swap
    children["push"] = &operations.Push
    return children
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preserve"] = operations.Preserve
    leafs["pop-and-forward"] = operations.PopAndForward
    return leafs
}

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetBundleName() string { return "cisco_ios_xe" }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetYangName() string { return "operations" }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetParent() types.Entity { return operations.parent }

func (operations *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetFilter() yfilter.YFilter { return swap.YFilter }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) SetFilter(yf yfilter.YFilter) { swap.YFilter = yf }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetSegmentPath() string {
    return "swap"
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &swap.Stack
    }
    return nil
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &swap.Stack
    return children
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetBundleName() string { return "cisco_ios_xe" }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetYangName() string { return "swap" }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) SetParent(parent types.Entity) { swap.parent = parent }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetParent() types.Entity { return swap.parent }

func (swap *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Swap_Stack) GetParentYangName() string { return "swap" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push
// Push outgoing label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The label stack.
    Stack MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetFilter() yfilter.YFilter { return push.YFilter }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) SetFilter(yf yfilter.YFilter) { push.YFilter = yf }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetGoName(yname string) string {
    if yname == "stack" { return "Stack" }
    return ""
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetSegmentPath() string {
    return "push"
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &push.Stack
    }
    return nil
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &push.Stack
    return children
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetBundleName() string { return "cisco_ios_xe" }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetYangName() string { return "push" }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) SetParent(parent types.Entity) { push.parent = parent }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetParent() types.Entity { return push.parent }

func (push *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push) GetParentYangName() string { return "operations" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack
// The label stack
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First label in the list is the top of the stack. The type is one of the
    // following types: slice of int with range: 16..1048575, or slice of  
    // :go:struct:`IetfMplsLabel
    // <ydk/models/cisco_ios_xe/common_mpls_types/IetfMplsLabel>`.
    LabelStack []interface{}
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetGoName(yname string) string {
    if yname == "label-stack" { return "LabelStack" }
    return ""
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-stack"] = stack.LabelStack
    return leafs
}

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetBundleName() string { return "cisco_ios_xe" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetYangName() string { return "stack" }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetParent() types.Entity { return stack.parent }

func (stack *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_Operations_Push_Stack) GetParentYangName() string { return "push" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats
// lsp stats
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics.
    Stats MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats
}

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetFilter() yfilter.YFilter { return nexthopStats.YFilter }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) SetFilter(yf yfilter.YFilter) { nexthopStats.YFilter = yf }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetGoName(yname string) string {
    if yname == "stats" { return "Stats" }
    return ""
}

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetSegmentPath() string {
    return "nexthop-stats"
}

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &nexthopStats.Stats
    }
    return nil
}

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &nexthopStats.Stats
    return children
}

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetBundleName() string { return "cisco_ios_xe" }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetYangName() string { return "nexthop-stats" }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) SetParent(parent types.Entity) { nexthopStats.parent = parent }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetParent() types.Entity { return nexthopStats.parent }

func (nexthopStats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats) GetParentYangName() string { return "next-hop" }

// MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats
// Statistics
type MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats struct {
    parent types.Entity
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

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    if yname == "dropped-bytes" { return "DroppedBytes" }
    return ""
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = stats.Packets
    leafs["bytes"] = stats.Bytes
    leafs["dropped-packets"] = stats.DroppedPackets
    leafs["dropped-bytes"] = stats.DroppedBytes
    return leafs
}

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetBundleName() string { return "cisco_ios_xe" }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetYangName() string { return "stats" }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetParent() types.Entity { return stats.parent }

func (stats *MplsStatic_MplsStaticState_LabelSwitchedPaths_LabelSwitchedPath_Path_NextHop_NexthopStats_Stats) GetParentYangName() string { return "nexthop-stats" }

