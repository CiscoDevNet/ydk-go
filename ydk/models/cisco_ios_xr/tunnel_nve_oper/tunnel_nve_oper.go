// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-nve package operational data.
// 
// This module contains definitions
// for the following management objects:
//   nve: NVE operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_nve_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_nve_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-nve-oper nve}", reflect.TypeOf(Nve{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-nve-oper:nve", reflect.TypeOf(Nve{}))
}

// Nve
// NVE operational data
type Nve struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table for VNIs.
    Vnis Nve_Vnis

    // Table for NVE interface attributes.
    Interfaces Nve_Interfaces
}

func (nve *Nve) GetFilter() yfilter.YFilter { return nve.YFilter }

func (nve *Nve) SetFilter(yf yfilter.YFilter) { nve.YFilter = yf }

func (nve *Nve) GetGoName(yname string) string {
    if yname == "vnis" { return "Vnis" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (nve *Nve) GetSegmentPath() string {
    return "Cisco-IOS-XR-tunnel-nve-oper:nve"
}

func (nve *Nve) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vnis" {
        return &nve.Vnis
    }
    if childYangName == "interfaces" {
        return &nve.Interfaces
    }
    return nil
}

func (nve *Nve) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vnis"] = &nve.Vnis
    children["interfaces"] = &nve.Interfaces
    return children
}

func (nve *Nve) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nve *Nve) GetBundleName() string { return "cisco_ios_xr" }

func (nve *Nve) GetYangName() string { return "nve" }

func (nve *Nve) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nve *Nve) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nve *Nve) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nve *Nve) SetParent(parent types.Entity) { nve.parent = parent }

func (nve *Nve) GetParent() types.Entity { return nve.parent }

func (nve *Nve) GetParentYangName() string { return "Cisco-IOS-XR-tunnel-nve-oper" }

// Nve_Vnis
// Table for VNIs
type Nve_Vnis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The attributes for a particular VNI. The type is slice of Nve_Vnis_Vni.
    Vni []Nve_Vnis_Vni
}

func (vnis *Nve_Vnis) GetFilter() yfilter.YFilter { return vnis.YFilter }

func (vnis *Nve_Vnis) SetFilter(yf yfilter.YFilter) { vnis.YFilter = yf }

func (vnis *Nve_Vnis) GetGoName(yname string) string {
    if yname == "vni" { return "Vni" }
    return ""
}

func (vnis *Nve_Vnis) GetSegmentPath() string {
    return "vnis"
}

func (vnis *Nve_Vnis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vni" {
        for _, c := range vnis.Vni {
            if vnis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Nve_Vnis_Vni{}
        vnis.Vni = append(vnis.Vni, child)
        return &vnis.Vni[len(vnis.Vni)-1]
    }
    return nil
}

func (vnis *Nve_Vnis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vnis.Vni {
        children[vnis.Vni[i].GetSegmentPath()] = &vnis.Vni[i]
    }
    return children
}

func (vnis *Nve_Vnis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vnis *Nve_Vnis) GetBundleName() string { return "cisco_ios_xr" }

func (vnis *Nve_Vnis) GetYangName() string { return "vnis" }

func (vnis *Nve_Vnis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vnis *Nve_Vnis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vnis *Nve_Vnis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vnis *Nve_Vnis) SetParent(parent types.Entity) { vnis.parent = parent }

func (vnis *Nve_Vnis) GetParent() types.Entity { return vnis.parent }

func (vnis *Nve_Vnis) GetParentYangName() string { return "nve" }

// Nve_Vnis_Vni
// The attributes for a particular VNI
type Nve_Vnis_Vni struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VNI ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Vni interface{}

    // NVE Interface name. The type is string.
    InterfaceName interface{}

    // VNI Number. The type is interface{} with range: 0..4294967295.
    VniXr interface{}

    // State. The type is interface{} with range: -128..127.
    State interface{}

    // MCAST IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    McastIpv4Address interface{}

    // Flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // VNI Min in Range. The type is interface{} with range: 0..4294967295.
    VniMin interface{}

    // VNI Max in Range. The type is interface{} with range: 0..4294967295.
    VniMax interface{}

    // McastFlags. The type is interface{} with range: 0..4294967295.
    McastFlags interface{}

    // UDP Port. The type is interface{} with range: 0..4294967295.
    UdpPort interface{}

    // BVI Interface Handle. The type is interface{} with range: 0..4294967295.
    BviIfh interface{}

    // BVI Interface Oper State. The type is interface{} with range: 0..255.
    BviState interface{}

    // BVI MAC address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    BviMac interface{}

    // L3 VRF Name. The type is string.
    VrfName interface{}

    // L3 VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // IPv4 Table ID. The type is interface{} with range: 0..4294967295.
    Ipv4TblId interface{}

    // IPv6 Table ID. The type is interface{} with range: 0..4294967295.
    Ipv6TblId interface{}

    // VRF VNI. The type is interface{} with range: 0..4294967295.
    VrfVni interface{}

    // TOPO ID valid flag. The type is bool.
    TopoValid interface{}

    // L2RIB Topology ID. The type is interface{} with range: 0..4294967295.
    TopoId interface{}

    // L2RIB Topology Name. The type is string with length: 0..50.
    TopoName interface{}
}

func (vni *Nve_Vnis_Vni) GetFilter() yfilter.YFilter { return vni.YFilter }

func (vni *Nve_Vnis_Vni) SetFilter(yf yfilter.YFilter) { vni.YFilter = yf }

func (vni *Nve_Vnis_Vni) GetGoName(yname string) string {
    if yname == "vni" { return "Vni" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "vni-xr" { return "VniXr" }
    if yname == "state" { return "State" }
    if yname == "mcast-ipv4-address" { return "McastIpv4Address" }
    if yname == "flags" { return "Flags" }
    if yname == "vni-min" { return "VniMin" }
    if yname == "vni-max" { return "VniMax" }
    if yname == "mcast-flags" { return "McastFlags" }
    if yname == "udp-port" { return "UdpPort" }
    if yname == "bvi-ifh" { return "BviIfh" }
    if yname == "bvi-state" { return "BviState" }
    if yname == "bvi-mac" { return "BviMac" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "ipv4-tbl-id" { return "Ipv4TblId" }
    if yname == "ipv6-tbl-id" { return "Ipv6TblId" }
    if yname == "vrf-vni" { return "VrfVni" }
    if yname == "topo-valid" { return "TopoValid" }
    if yname == "topo-id" { return "TopoId" }
    if yname == "topo-name" { return "TopoName" }
    return ""
}

func (vni *Nve_Vnis_Vni) GetSegmentPath() string {
    return "vni" + "[vni='" + fmt.Sprintf("%v", vni.Vni) + "']"
}

func (vni *Nve_Vnis_Vni) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vni *Nve_Vnis_Vni) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vni *Nve_Vnis_Vni) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vni"] = vni.Vni
    leafs["interface-name"] = vni.InterfaceName
    leafs["vni-xr"] = vni.VniXr
    leafs["state"] = vni.State
    leafs["mcast-ipv4-address"] = vni.McastIpv4Address
    leafs["flags"] = vni.Flags
    leafs["vni-min"] = vni.VniMin
    leafs["vni-max"] = vni.VniMax
    leafs["mcast-flags"] = vni.McastFlags
    leafs["udp-port"] = vni.UdpPort
    leafs["bvi-ifh"] = vni.BviIfh
    leafs["bvi-state"] = vni.BviState
    leafs["bvi-mac"] = vni.BviMac
    leafs["vrf-name"] = vni.VrfName
    leafs["vrf-id"] = vni.VrfId
    leafs["ipv4-tbl-id"] = vni.Ipv4TblId
    leafs["ipv6-tbl-id"] = vni.Ipv6TblId
    leafs["vrf-vni"] = vni.VrfVni
    leafs["topo-valid"] = vni.TopoValid
    leafs["topo-id"] = vni.TopoId
    leafs["topo-name"] = vni.TopoName
    return leafs
}

func (vni *Nve_Vnis_Vni) GetBundleName() string { return "cisco_ios_xr" }

func (vni *Nve_Vnis_Vni) GetYangName() string { return "vni" }

func (vni *Nve_Vnis_Vni) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vni *Nve_Vnis_Vni) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vni *Nve_Vnis_Vni) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vni *Nve_Vnis_Vni) SetParent(parent types.Entity) { vni.parent = parent }

func (vni *Nve_Vnis_Vni) GetParent() types.Entity { return vni.parent }

func (vni *Nve_Vnis_Vni) GetParentYangName() string { return "vnis" }

// Nve_Interfaces
// Table for NVE interface attributes
type Nve_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The attributes for a particular interface. The type is slice of
    // Nve_Interfaces_Interface.
    Interface []Nve_Interfaces_Interface
}

func (interfaces *Nve_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Nve_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Nve_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Nve_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Nve_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Nve_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Nve_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Nve_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Nve_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Nve_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Nve_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Nve_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Nve_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Nve_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Nve_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Nve_Interfaces) GetParentYangName() string { return "nve" }

// Nve_Interfaces_Interface
// The attributes for a particular interface
type Nve_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // State. The type is interface{} with range: -128..127.
    State interface{}

    // Admin State. The type is interface{} with range: -128..127.
    AdminState interface{}

    // Flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Encap. The type is interface{} with range: -128..127.
    Encap interface{}

    // Source Interface name. The type is string.
    SourceInterfaceName interface{}

    // Source IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceIpv4Address interface{}

    // NVE IfHandle. The type is interface{} with range: 0..18446744073709551615.
    IfHandle interface{}

    // Source Intf State. The type is interface{} with range: -128..127.
    SourceState interface{}

    // UDP Port. The type is interface{} with range: 0..4294967295.
    UdpPort interface{}

    // Anycast Source Interface name. The type is string.
    AnyCastSourceInterfaceName interface{}

    // Anycast Source IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AnyCastSourceIpv4Address interface{}

    // Anycast Source Interface State. The type is interface{} with range:
    // -128..127.
    AnyCastSourceState interface{}

    // MCAST sync group IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SyncMcastIpv4Address interface{}

    // Sync McastFlags. The type is interface{} with range: 0..4294967295.
    SyncMcastFlags interface{}
}

func (self *Nve_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Nve_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Nve_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "state" { return "State" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "flags" { return "Flags" }
    if yname == "encap" { return "Encap" }
    if yname == "source-interface-name" { return "SourceInterfaceName" }
    if yname == "source-ipv4-address" { return "SourceIpv4Address" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "source-state" { return "SourceState" }
    if yname == "udp-port" { return "UdpPort" }
    if yname == "any-cast-source-interface-name" { return "AnyCastSourceInterfaceName" }
    if yname == "any-cast-source-ipv4-address" { return "AnyCastSourceIpv4Address" }
    if yname == "any-cast-source-state" { return "AnyCastSourceState" }
    if yname == "sync-mcast-ipv4-address" { return "SyncMcastIpv4Address" }
    if yname == "sync-mcast-flags" { return "SyncMcastFlags" }
    return ""
}

func (self *Nve_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Nve_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Nve_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Nve_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-name-xr"] = self.InterfaceNameXr
    leafs["state"] = self.State
    leafs["admin-state"] = self.AdminState
    leafs["flags"] = self.Flags
    leafs["encap"] = self.Encap
    leafs["source-interface-name"] = self.SourceInterfaceName
    leafs["source-ipv4-address"] = self.SourceIpv4Address
    leafs["if-handle"] = self.IfHandle
    leafs["source-state"] = self.SourceState
    leafs["udp-port"] = self.UdpPort
    leafs["any-cast-source-interface-name"] = self.AnyCastSourceInterfaceName
    leafs["any-cast-source-ipv4-address"] = self.AnyCastSourceIpv4Address
    leafs["any-cast-source-state"] = self.AnyCastSourceState
    leafs["sync-mcast-ipv4-address"] = self.SyncMcastIpv4Address
    leafs["sync-mcast-flags"] = self.SyncMcastFlags
    return leafs
}

func (self *Nve_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Nve_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Nve_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Nve_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Nve_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Nve_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Nve_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Nve_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

