// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp-lldp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lldp-snoop-data: Information related to LLDP Snoop
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_mxp_lldp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_mxp_lldp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-mxp-lldp-oper lldp-snoop-data}", reflect.TypeOf(LldpSnoopData{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-mxp-lldp-oper:lldp-snoop-data", reflect.TypeOf(LldpSnoopData{}))
}

// LldpL3AddrProtocol represents Lldp l3 addr protocol
type LldpL3AddrProtocol string

const (
    // IPv4
    LldpL3AddrProtocol_ipv4 LldpL3AddrProtocol = "ipv4"

    // IPv6
    LldpL3AddrProtocol_ipv6 LldpL3AddrProtocol = "ipv6"
)

// LldpSnoopData
// Information related to LLDP Snoop
type LldpSnoopData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet controller snoop data.
    EthernetControllerNames LldpSnoopData_EthernetControllerNames
}

func (lldpSnoopData *LldpSnoopData) GetEntityData() *types.CommonEntityData {
    lldpSnoopData.EntityData.YFilter = lldpSnoopData.YFilter
    lldpSnoopData.EntityData.YangName = "lldp-snoop-data"
    lldpSnoopData.EntityData.BundleName = "cisco_ios_xr"
    lldpSnoopData.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1k-mxp-lldp-oper"
    lldpSnoopData.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1k-mxp-lldp-oper:lldp-snoop-data"
    lldpSnoopData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpSnoopData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpSnoopData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpSnoopData.EntityData.Children = make(map[string]types.YChild)
    lldpSnoopData.EntityData.Children["ethernet-controller-names"] = types.YChild{"EthernetControllerNames", &lldpSnoopData.EthernetControllerNames}
    lldpSnoopData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpSnoopData.EntityData)
}

// LldpSnoopData_EthernetControllerNames
// Ethernet controller snoop data
type LldpSnoopData_EthernetControllerNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // port Name. The type is slice of
    // LldpSnoopData_EthernetControllerNames_EthernetControllerName.
    EthernetControllerName []LldpSnoopData_EthernetControllerNames_EthernetControllerName
}

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetEntityData() *types.CommonEntityData {
    ethernetControllerNames.EntityData.YFilter = ethernetControllerNames.YFilter
    ethernetControllerNames.EntityData.YangName = "ethernet-controller-names"
    ethernetControllerNames.EntityData.BundleName = "cisco_ios_xr"
    ethernetControllerNames.EntityData.ParentYangName = "lldp-snoop-data"
    ethernetControllerNames.EntityData.SegmentPath = "ethernet-controller-names"
    ethernetControllerNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetControllerNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetControllerNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetControllerNames.EntityData.Children = make(map[string]types.YChild)
    ethernetControllerNames.EntityData.Children["ethernet-controller-name"] = types.YChild{"EthernetControllerName", nil}
    for i := range ethernetControllerNames.EthernetControllerName {
        ethernetControllerNames.EntityData.Children[types.GetSegmentPath(&ethernetControllerNames.EthernetControllerName[i])] = types.YChild{"EthernetControllerName", &ethernetControllerNames.EthernetControllerName[i]}
    }
    ethernetControllerNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetControllerNames.EntityData)
}

// LldpSnoopData_EthernetControllerNames_EthernetControllerName
// port Name
type LldpSnoopData_EthernetControllerNames_EthernetControllerName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // Mac address of the neighbor. The type is string.
    SourceMac interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // LldpNeighbor. The type is string with length: 0..40.
    LldpNeighbor interface{}

    // LLDP Packet Drop Enabled. The type is bool.
    DropEnabled interface{}

    // Received LLDP Packets count. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLldpPkts interface{}

    // Management Addresses.
    NetworkAddresses LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses
}

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetEntityData() *types.CommonEntityData {
    ethernetControllerName.EntityData.YFilter = ethernetControllerName.YFilter
    ethernetControllerName.EntityData.YangName = "ethernet-controller-name"
    ethernetControllerName.EntityData.BundleName = "cisco_ios_xr"
    ethernetControllerName.EntityData.ParentYangName = "ethernet-controller-names"
    ethernetControllerName.EntityData.SegmentPath = "ethernet-controller-name" + "[name='" + fmt.Sprintf("%v", ethernetControllerName.Name) + "']"
    ethernetControllerName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetControllerName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetControllerName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetControllerName.EntityData.Children = make(map[string]types.YChild)
    ethernetControllerName.EntityData.Children["network-addresses"] = types.YChild{"NetworkAddresses", &ethernetControllerName.NetworkAddresses}
    ethernetControllerName.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetControllerName.EntityData.Leafs["name"] = types.YLeaf{"Name", ethernetControllerName.Name}
    ethernetControllerName.EntityData.Leafs["source-mac"] = types.YLeaf{"SourceMac", ethernetControllerName.SourceMac}
    ethernetControllerName.EntityData.Leafs["chassis-id"] = types.YLeaf{"ChassisId", ethernetControllerName.ChassisId}
    ethernetControllerName.EntityData.Leafs["port-id-detail"] = types.YLeaf{"PortIdDetail", ethernetControllerName.PortIdDetail}
    ethernetControllerName.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", ethernetControllerName.HoldTime}
    ethernetControllerName.EntityData.Leafs["port-description"] = types.YLeaf{"PortDescription", ethernetControllerName.PortDescription}
    ethernetControllerName.EntityData.Leafs["system-name"] = types.YLeaf{"SystemName", ethernetControllerName.SystemName}
    ethernetControllerName.EntityData.Leafs["system-description"] = types.YLeaf{"SystemDescription", ethernetControllerName.SystemDescription}
    ethernetControllerName.EntityData.Leafs["system-capabilities"] = types.YLeaf{"SystemCapabilities", ethernetControllerName.SystemCapabilities}
    ethernetControllerName.EntityData.Leafs["enabled-capabilities"] = types.YLeaf{"EnabledCapabilities", ethernetControllerName.EnabledCapabilities}
    ethernetControllerName.EntityData.Leafs["lldp-neighbor"] = types.YLeaf{"LldpNeighbor", ethernetControllerName.LldpNeighbor}
    ethernetControllerName.EntityData.Leafs["drop-enabled"] = types.YLeaf{"DropEnabled", ethernetControllerName.DropEnabled}
    ethernetControllerName.EntityData.Leafs["rx-lldp-pkts"] = types.YLeaf{"RxLldpPkts", ethernetControllerName.RxLldpPkts}
    return &(ethernetControllerName.EntityData)
}

// LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses
// Management Addresses
type LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "ethernet-controller-name"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = make(map[string]types.YChild)
    networkAddresses.EntityData.Children["lldp-addr-entry"] = types.YChild{"LldpAddrEntry", nil}
    for i := range networkAddresses.LldpAddrEntry {
        networkAddresses.EntityData.Children[types.GetSegmentPath(&networkAddresses.LldpAddrEntry[i])] = types.YChild{"LldpAddrEntry", &networkAddresses.LldpAddrEntry[i]}
    }
    networkAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networkAddresses.EntityData)
}

// LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetEntityData() *types.CommonEntityData {
    lldpAddrEntry.EntityData.YFilter = lldpAddrEntry.YFilter
    lldpAddrEntry.EntityData.YangName = "lldp-addr-entry"
    lldpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpAddrEntry.EntityData.ParentYangName = "network-addresses"
    lldpAddrEntry.EntityData.SegmentPath = "lldp-addr-entry"
    lldpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpAddrEntry.EntityData.Children = make(map[string]types.YChild)
    lldpAddrEntry.EntityData.Children["address"] = types.YChild{"Address", &lldpAddrEntry.Address}
    lldpAddrEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpAddrEntry.EntityData.Leafs["ma-subtype"] = types.YLeaf{"MaSubtype", lldpAddrEntry.MaSubtype}
    lldpAddrEntry.EntityData.Leafs["if-num"] = types.YLeaf{"IfNum", lldpAddrEntry.IfNum}
    return &(lldpAddrEntry.EntityData)
}

// LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "lldp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-type"] = types.YLeaf{"AddressType", address.AddressType}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

