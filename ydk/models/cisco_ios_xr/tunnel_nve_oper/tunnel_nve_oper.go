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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table for VNIs.
    Vnis Nve_Vnis

    // Table for NVE interface attributes.
    Interfaces Nve_Interfaces
}

func (nve *Nve) GetEntityData() *types.CommonEntityData {
    nve.EntityData.YFilter = nve.YFilter
    nve.EntityData.YangName = "nve"
    nve.EntityData.BundleName = "cisco_ios_xr"
    nve.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-nve-oper"
    nve.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-nve-oper:nve"
    nve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nve.EntityData.Children = make(map[string]types.YChild)
    nve.EntityData.Children["vnis"] = types.YChild{"Vnis", &nve.Vnis}
    nve.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &nve.Interfaces}
    nve.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nve.EntityData)
}

// Nve_Vnis
// Table for VNIs
type Nve_Vnis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The attributes for a particular VNI. The type is slice of Nve_Vnis_Vni.
    Vni []Nve_Vnis_Vni
}

func (vnis *Nve_Vnis) GetEntityData() *types.CommonEntityData {
    vnis.EntityData.YFilter = vnis.YFilter
    vnis.EntityData.YangName = "vnis"
    vnis.EntityData.BundleName = "cisco_ios_xr"
    vnis.EntityData.ParentYangName = "nve"
    vnis.EntityData.SegmentPath = "vnis"
    vnis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vnis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vnis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vnis.EntityData.Children = make(map[string]types.YChild)
    vnis.EntityData.Children["vni"] = types.YChild{"Vni", nil}
    for i := range vnis.Vni {
        vnis.EntityData.Children[types.GetSegmentPath(&vnis.Vni[i])] = types.YChild{"Vni", &vnis.Vni[i]}
    }
    vnis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vnis.EntityData)
}

// Nve_Vnis_Vni
// The attributes for a particular VNI
type Nve_Vnis_Vni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VNI ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Vni interface{}

    // NVE Interface name. The type is string.
    InterfaceName interface{}

    // VNI Number. The type is interface{} with range: 0..4294967295.
    VniXr interface{}

    // State. The type is interface{} with range: -128..127.
    State interface{}

    // MCAST IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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

func (vni *Nve_Vnis_Vni) GetEntityData() *types.CommonEntityData {
    vni.EntityData.YFilter = vni.YFilter
    vni.EntityData.YangName = "vni"
    vni.EntityData.BundleName = "cisco_ios_xr"
    vni.EntityData.ParentYangName = "vnis"
    vni.EntityData.SegmentPath = "vni" + "[vni='" + fmt.Sprintf("%v", vni.Vni) + "']"
    vni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vni.EntityData.Children = make(map[string]types.YChild)
    vni.EntityData.Leafs = make(map[string]types.YLeaf)
    vni.EntityData.Leafs["vni"] = types.YLeaf{"Vni", vni.Vni}
    vni.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", vni.InterfaceName}
    vni.EntityData.Leafs["vni-xr"] = types.YLeaf{"VniXr", vni.VniXr}
    vni.EntityData.Leafs["state"] = types.YLeaf{"State", vni.State}
    vni.EntityData.Leafs["mcast-ipv4-address"] = types.YLeaf{"McastIpv4Address", vni.McastIpv4Address}
    vni.EntityData.Leafs["flags"] = types.YLeaf{"Flags", vni.Flags}
    vni.EntityData.Leafs["vni-min"] = types.YLeaf{"VniMin", vni.VniMin}
    vni.EntityData.Leafs["vni-max"] = types.YLeaf{"VniMax", vni.VniMax}
    vni.EntityData.Leafs["mcast-flags"] = types.YLeaf{"McastFlags", vni.McastFlags}
    vni.EntityData.Leafs["udp-port"] = types.YLeaf{"UdpPort", vni.UdpPort}
    vni.EntityData.Leafs["bvi-ifh"] = types.YLeaf{"BviIfh", vni.BviIfh}
    vni.EntityData.Leafs["bvi-state"] = types.YLeaf{"BviState", vni.BviState}
    vni.EntityData.Leafs["bvi-mac"] = types.YLeaf{"BviMac", vni.BviMac}
    vni.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vni.VrfName}
    vni.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", vni.VrfId}
    vni.EntityData.Leafs["ipv4-tbl-id"] = types.YLeaf{"Ipv4TblId", vni.Ipv4TblId}
    vni.EntityData.Leafs["ipv6-tbl-id"] = types.YLeaf{"Ipv6TblId", vni.Ipv6TblId}
    vni.EntityData.Leafs["vrf-vni"] = types.YLeaf{"VrfVni", vni.VrfVni}
    vni.EntityData.Leafs["topo-valid"] = types.YLeaf{"TopoValid", vni.TopoValid}
    vni.EntityData.Leafs["topo-id"] = types.YLeaf{"TopoId", vni.TopoId}
    vni.EntityData.Leafs["topo-name"] = types.YLeaf{"TopoName", vni.TopoName}
    return &(vni.EntityData)
}

// Nve_Interfaces
// Table for NVE interface attributes
type Nve_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The attributes for a particular interface. The type is slice of
    // Nve_Interfaces_Interface_.
    Interface_ []Nve_Interfaces_Interface
}

func (interfaces *Nve_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "nve"
    interfaces.EntityData.SegmentPath = "interfaces"
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

// Nve_Interfaces_Interface
// The attributes for a particular interface
type Nve_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AnyCastSourceIpv4Address interface{}

    // Anycast Source Interface State. The type is interface{} with range:
    // -128..127.
    AnyCastSourceState interface{}

    // MCAST sync group IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SyncMcastIpv4Address interface{}

    // Sync McastFlags. The type is interface{} with range: 0..4294967295.
    SyncMcastFlags interface{}
}

func (self *Nve_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", self.InterfaceNameXr}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", self.AdminState}
    self.EntityData.Leafs["flags"] = types.YLeaf{"Flags", self.Flags}
    self.EntityData.Leafs["encap"] = types.YLeaf{"Encap", self.Encap}
    self.EntityData.Leafs["source-interface-name"] = types.YLeaf{"SourceInterfaceName", self.SourceInterfaceName}
    self.EntityData.Leafs["source-ipv4-address"] = types.YLeaf{"SourceIpv4Address", self.SourceIpv4Address}
    self.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", self.IfHandle}
    self.EntityData.Leafs["source-state"] = types.YLeaf{"SourceState", self.SourceState}
    self.EntityData.Leafs["udp-port"] = types.YLeaf{"UdpPort", self.UdpPort}
    self.EntityData.Leafs["any-cast-source-interface-name"] = types.YLeaf{"AnyCastSourceInterfaceName", self.AnyCastSourceInterfaceName}
    self.EntityData.Leafs["any-cast-source-ipv4-address"] = types.YLeaf{"AnyCastSourceIpv4Address", self.AnyCastSourceIpv4Address}
    self.EntityData.Leafs["any-cast-source-state"] = types.YLeaf{"AnyCastSourceState", self.AnyCastSourceState}
    self.EntityData.Leafs["sync-mcast-ipv4-address"] = types.YLeaf{"SyncMcastIpv4Address", self.SyncMcastIpv4Address}
    self.EntityData.Leafs["sync-mcast-flags"] = types.YLeaf{"SyncMcastFlags", self.SyncMcastFlags}
    return &(self.EntityData)
}

