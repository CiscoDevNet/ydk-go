// This module contains a collection of YANG definitions
// for bridge data.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package bridge_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bridge_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-bridge-oper bridge-instances}", reflect.TypeOf(BridgeInstances{}))
    ydk.RegisterEntity("Cisco-IOS-XE-bridge-oper:bridge-instances", reflect.TypeOf(BridgeInstances{}))
}

// BridgeMacType represents MAC address type
type BridgeMacType string

const (
    BridgeMacType_bridge_mac_type_static BridgeMacType = "bridge-mac-type-static"

    BridgeMacType_bridge_mac_type_dynamic BridgeMacType = "bridge-mac-type-dynamic"
)

// IntfStatusType represents Interface status type
type IntfStatusType string

const (
    IntfStatusType_up IntfStatusType = "up"

    IntfStatusType_down IntfStatusType = "down"
)

// BridgeInstances
// This is top level container for bridge table. It can have one
// or more bridge entry.
type BridgeInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The bridge lists is an ordered list of bridge entries. Each bridge entries
    // has a list of bridge interface members, and bridge attributes. The type is
    // slice of BridgeInstances_BridgeEntry.
    BridgeEntry []*BridgeInstances_BridgeEntry
}

func (bridgeInstances *BridgeInstances) GetEntityData() *types.CommonEntityData {
    bridgeInstances.EntityData.YFilter = bridgeInstances.YFilter
    bridgeInstances.EntityData.YangName = "bridge-instances"
    bridgeInstances.EntityData.BundleName = "cisco_ios_xe"
    bridgeInstances.EntityData.ParentYangName = "Cisco-IOS-XE-bridge-oper"
    bridgeInstances.EntityData.SegmentPath = "Cisco-IOS-XE-bridge-oper:bridge-instances"
    bridgeInstances.EntityData.AbsolutePath = bridgeInstances.EntityData.SegmentPath
    bridgeInstances.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeInstances.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeInstances.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeInstances.EntityData.Children = types.NewOrderedMap()
    bridgeInstances.EntityData.Children.Append("bridge-entry", types.YChild{"BridgeEntry", nil})
    for i := range bridgeInstances.BridgeEntry {
        bridgeInstances.EntityData.Children.Append(types.GetSegmentPath(bridgeInstances.BridgeEntry[i]), types.YChild{"BridgeEntry", bridgeInstances.BridgeEntry[i]})
    }
    bridgeInstances.EntityData.Leafs = types.NewOrderedMap()

    bridgeInstances.EntityData.YListKeys = []string {}

    return &(bridgeInstances.EntityData)
}

// BridgeInstances_BridgeEntry
// The bridge lists is an ordered list of bridge entries.
// Each bridge entries has a list of bridge interface members,
// and bridge attributes.
type BridgeInstances_BridgeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Bridge id <1..4094>. The type is interface{} with
    // range: 0..4294967295.
    BridgeId interface{}

    // VLAN name. The type is string.
    Name interface{}

    // VLAN identifier <1..4094>. The type is interface{} with range:
    // 0..4294967295.
    Vlan interface{}

    // The name of VLAN routing interface. The type is string.
    RoutingInterface interface{}

    // The maximum number of MAC learn limit for bridge. The type is interface{}
    // with range: 0..4294967295.
    MaxMacs interface{}

    // The number of MAC learned in bridge. The type is interface{} with range:
    // 0..4294967295.
    NumMacs interface{}

    // The aging time of MAC address,0 or <10..1000000> second. The type is
    // interface{} with range: 0..4294967295.
    AgeTime interface{}

    // The number of received packets in bridge. The type is interface{} with
    // range: 0..18446744073709551615.
    RxPackets interface{}

    // The number of received bytes in bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    RxOctets interface{}

    // The number of transmited packets out of bridge. The type is interface{}
    // with range: 0..18446744073709551615.
    TxPackets interface{}

    // The number of transmited bytes out of bridge. The type is interface{} with
    // range: 0..18446744073709551615.
    TxOctets interface{}

    // The number of flood packets in bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    FloodPackets interface{}

    // The number of flood bytes in bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    FloodOctets interface{}

    // L3 packets received from bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    RxRoutedPackets interface{}

    // L3 packets transmit from bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    TxRoutedPackets interface{}

    // MAC learned counter in bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    Learn interface{}

    // MAC aging counter in bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    Age interface{}

    // MAC move counter in bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    Move interface{}

    // Bridge interface member information.
    BridgeIntfEntries BridgeInstances_BridgeEntry_BridgeIntfEntries

    // Bridge matm member information.
    BridgeMatmEntries BridgeInstances_BridgeEntry_BridgeMatmEntries
}

func (bridgeEntry *BridgeInstances_BridgeEntry) GetEntityData() *types.CommonEntityData {
    bridgeEntry.EntityData.YFilter = bridgeEntry.YFilter
    bridgeEntry.EntityData.YangName = "bridge-entry"
    bridgeEntry.EntityData.BundleName = "cisco_ios_xe"
    bridgeEntry.EntityData.ParentYangName = "bridge-instances"
    bridgeEntry.EntityData.SegmentPath = "bridge-entry" + types.AddKeyToken(bridgeEntry.BridgeId, "bridge-id")
    bridgeEntry.EntityData.AbsolutePath = "Cisco-IOS-XE-bridge-oper:bridge-instances/" + bridgeEntry.EntityData.SegmentPath
    bridgeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeEntry.EntityData.Children = types.NewOrderedMap()
    bridgeEntry.EntityData.Children.Append("bridge-intf-entries", types.YChild{"BridgeIntfEntries", &bridgeEntry.BridgeIntfEntries})
    bridgeEntry.EntityData.Children.Append("bridge-matm-entries", types.YChild{"BridgeMatmEntries", &bridgeEntry.BridgeMatmEntries})
    bridgeEntry.EntityData.Leafs = types.NewOrderedMap()
    bridgeEntry.EntityData.Leafs.Append("bridge-id", types.YLeaf{"BridgeId", bridgeEntry.BridgeId})
    bridgeEntry.EntityData.Leafs.Append("name", types.YLeaf{"Name", bridgeEntry.Name})
    bridgeEntry.EntityData.Leafs.Append("vlan", types.YLeaf{"Vlan", bridgeEntry.Vlan})
    bridgeEntry.EntityData.Leafs.Append("routing-interface", types.YLeaf{"RoutingInterface", bridgeEntry.RoutingInterface})
    bridgeEntry.EntityData.Leafs.Append("max-macs", types.YLeaf{"MaxMacs", bridgeEntry.MaxMacs})
    bridgeEntry.EntityData.Leafs.Append("num-macs", types.YLeaf{"NumMacs", bridgeEntry.NumMacs})
    bridgeEntry.EntityData.Leafs.Append("age-time", types.YLeaf{"AgeTime", bridgeEntry.AgeTime})
    bridgeEntry.EntityData.Leafs.Append("rx-packets", types.YLeaf{"RxPackets", bridgeEntry.RxPackets})
    bridgeEntry.EntityData.Leafs.Append("rx-octets", types.YLeaf{"RxOctets", bridgeEntry.RxOctets})
    bridgeEntry.EntityData.Leafs.Append("tx-packets", types.YLeaf{"TxPackets", bridgeEntry.TxPackets})
    bridgeEntry.EntityData.Leafs.Append("tx-octets", types.YLeaf{"TxOctets", bridgeEntry.TxOctets})
    bridgeEntry.EntityData.Leafs.Append("flood-packets", types.YLeaf{"FloodPackets", bridgeEntry.FloodPackets})
    bridgeEntry.EntityData.Leafs.Append("flood-octets", types.YLeaf{"FloodOctets", bridgeEntry.FloodOctets})
    bridgeEntry.EntityData.Leafs.Append("rx-routed-packets", types.YLeaf{"RxRoutedPackets", bridgeEntry.RxRoutedPackets})
    bridgeEntry.EntityData.Leafs.Append("tx-routed-packets", types.YLeaf{"TxRoutedPackets", bridgeEntry.TxRoutedPackets})
    bridgeEntry.EntityData.Leafs.Append("learn", types.YLeaf{"Learn", bridgeEntry.Learn})
    bridgeEntry.EntityData.Leafs.Append("age", types.YLeaf{"Age", bridgeEntry.Age})
    bridgeEntry.EntityData.Leafs.Append("move", types.YLeaf{"Move", bridgeEntry.Move})

    bridgeEntry.EntityData.YListKeys = []string {"BridgeId"}

    return &(bridgeEntry.EntityData)
}

// BridgeInstances_BridgeEntry_BridgeIntfEntries
// Bridge interface member information
type BridgeInstances_BridgeEntry_BridgeIntfEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of bridge interface. The type is slice of
    // BridgeInstances_BridgeEntry_BridgeIntfEntries_BridgeIntfEntry.
    BridgeIntfEntry []*BridgeInstances_BridgeEntry_BridgeIntfEntries_BridgeIntfEntry
}

func (bridgeIntfEntries *BridgeInstances_BridgeEntry_BridgeIntfEntries) GetEntityData() *types.CommonEntityData {
    bridgeIntfEntries.EntityData.YFilter = bridgeIntfEntries.YFilter
    bridgeIntfEntries.EntityData.YangName = "bridge-intf-entries"
    bridgeIntfEntries.EntityData.BundleName = "cisco_ios_xe"
    bridgeIntfEntries.EntityData.ParentYangName = "bridge-entry"
    bridgeIntfEntries.EntityData.SegmentPath = "bridge-intf-entries"
    bridgeIntfEntries.EntityData.AbsolutePath = "Cisco-IOS-XE-bridge-oper:bridge-instances/bridge-entry/" + bridgeIntfEntries.EntityData.SegmentPath
    bridgeIntfEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeIntfEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeIntfEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeIntfEntries.EntityData.Children = types.NewOrderedMap()
    bridgeIntfEntries.EntityData.Children.Append("bridge-intf-entry", types.YChild{"BridgeIntfEntry", nil})
    for i := range bridgeIntfEntries.BridgeIntfEntry {
        bridgeIntfEntries.EntityData.Children.Append(types.GetSegmentPath(bridgeIntfEntries.BridgeIntfEntry[i]), types.YChild{"BridgeIntfEntry", bridgeIntfEntries.BridgeIntfEntry[i]})
    }
    bridgeIntfEntries.EntityData.Leafs = types.NewOrderedMap()

    bridgeIntfEntries.EntityData.YListKeys = []string {}

    return &(bridgeIntfEntries.EntityData)
}

// BridgeInstances_BridgeEntry_BridgeIntfEntries_BridgeIntfEntry
// A list of bridge interface
type BridgeInstances_BridgeEntry_BridgeIntfEntries_BridgeIntfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Switch port name belong to the bridge. The type is
    // string.
    IfName interface{}

    // VLAN identifier. The type is interface{} with range: 0..4294967295.
    Vlan interface{}

    // If the VLAN is native VLAN. The type is bool.
    NativeVlan interface{}

    // Bridge interface administration status. The type is IntfStatusType.
    AdminStatus interface{}

    // Bridge interface operational status. The type is IntfStatusType.
    OperStatus interface{}

    // Bridge interface encapsulation type. The type is string.
    EncapType interface{}

    // Bridge interface index. The type is interface{} with range: 0..4294967295.
    Ifindex interface{}

    // The Maximum Transmission Unit(byte) of Bridge interface. The type is
    // interface{} with range: 0..4294967295.
    Mtu interface{}
}

func (bridgeIntfEntry *BridgeInstances_BridgeEntry_BridgeIntfEntries_BridgeIntfEntry) GetEntityData() *types.CommonEntityData {
    bridgeIntfEntry.EntityData.YFilter = bridgeIntfEntry.YFilter
    bridgeIntfEntry.EntityData.YangName = "bridge-intf-entry"
    bridgeIntfEntry.EntityData.BundleName = "cisco_ios_xe"
    bridgeIntfEntry.EntityData.ParentYangName = "bridge-intf-entries"
    bridgeIntfEntry.EntityData.SegmentPath = "bridge-intf-entry" + types.AddKeyToken(bridgeIntfEntry.IfName, "if-name")
    bridgeIntfEntry.EntityData.AbsolutePath = "Cisco-IOS-XE-bridge-oper:bridge-instances/bridge-entry/bridge-intf-entries/" + bridgeIntfEntry.EntityData.SegmentPath
    bridgeIntfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeIntfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeIntfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeIntfEntry.EntityData.Children = types.NewOrderedMap()
    bridgeIntfEntry.EntityData.Leafs = types.NewOrderedMap()
    bridgeIntfEntry.EntityData.Leafs.Append("if-name", types.YLeaf{"IfName", bridgeIntfEntry.IfName})
    bridgeIntfEntry.EntityData.Leafs.Append("vlan", types.YLeaf{"Vlan", bridgeIntfEntry.Vlan})
    bridgeIntfEntry.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", bridgeIntfEntry.NativeVlan})
    bridgeIntfEntry.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", bridgeIntfEntry.AdminStatus})
    bridgeIntfEntry.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", bridgeIntfEntry.OperStatus})
    bridgeIntfEntry.EntityData.Leafs.Append("encap-type", types.YLeaf{"EncapType", bridgeIntfEntry.EncapType})
    bridgeIntfEntry.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", bridgeIntfEntry.Ifindex})
    bridgeIntfEntry.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", bridgeIntfEntry.Mtu})

    bridgeIntfEntry.EntityData.YListKeys = []string {"IfName"}

    return &(bridgeIntfEntry.EntityData)
}

// BridgeInstances_BridgeEntry_BridgeMatmEntries
// Bridge matm member information
type BridgeInstances_BridgeEntry_BridgeMatmEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of mac table. The type is slice of
    // BridgeInstances_BridgeEntry_BridgeMatmEntries_BridgeMatmEntry.
    BridgeMatmEntry []*BridgeInstances_BridgeEntry_BridgeMatmEntries_BridgeMatmEntry
}

func (bridgeMatmEntries *BridgeInstances_BridgeEntry_BridgeMatmEntries) GetEntityData() *types.CommonEntityData {
    bridgeMatmEntries.EntityData.YFilter = bridgeMatmEntries.YFilter
    bridgeMatmEntries.EntityData.YangName = "bridge-matm-entries"
    bridgeMatmEntries.EntityData.BundleName = "cisco_ios_xe"
    bridgeMatmEntries.EntityData.ParentYangName = "bridge-entry"
    bridgeMatmEntries.EntityData.SegmentPath = "bridge-matm-entries"
    bridgeMatmEntries.EntityData.AbsolutePath = "Cisco-IOS-XE-bridge-oper:bridge-instances/bridge-entry/" + bridgeMatmEntries.EntityData.SegmentPath
    bridgeMatmEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeMatmEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeMatmEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeMatmEntries.EntityData.Children = types.NewOrderedMap()
    bridgeMatmEntries.EntityData.Children.Append("bridge-matm-entry", types.YChild{"BridgeMatmEntry", nil})
    for i := range bridgeMatmEntries.BridgeMatmEntry {
        bridgeMatmEntries.EntityData.Children.Append(types.GetSegmentPath(bridgeMatmEntries.BridgeMatmEntry[i]), types.YChild{"BridgeMatmEntry", bridgeMatmEntries.BridgeMatmEntry[i]})
    }
    bridgeMatmEntries.EntityData.Leafs = types.NewOrderedMap()

    bridgeMatmEntries.EntityData.YListKeys = []string {}

    return &(bridgeMatmEntries.EntityData)
}

// BridgeInstances_BridgeEntry_BridgeMatmEntries_BridgeMatmEntry
// A list of mac table
type BridgeInstances_BridgeEntry_BridgeMatmEntries_BridgeMatmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Interface name which MAC learnt from. The type is slice of string.
    Interface []interface{}

    // VLAN identifier. The type is interface{} with range: 0..4294967295.
    Vlan interface{}

    // MAC type. The type is BridgeMacType.
    Type interface{}
}

func (bridgeMatmEntry *BridgeInstances_BridgeEntry_BridgeMatmEntries_BridgeMatmEntry) GetEntityData() *types.CommonEntityData {
    bridgeMatmEntry.EntityData.YFilter = bridgeMatmEntry.YFilter
    bridgeMatmEntry.EntityData.YangName = "bridge-matm-entry"
    bridgeMatmEntry.EntityData.BundleName = "cisco_ios_xe"
    bridgeMatmEntry.EntityData.ParentYangName = "bridge-matm-entries"
    bridgeMatmEntry.EntityData.SegmentPath = "bridge-matm-entry" + types.AddKeyToken(bridgeMatmEntry.MacAddress, "mac-address")
    bridgeMatmEntry.EntityData.AbsolutePath = "Cisco-IOS-XE-bridge-oper:bridge-instances/bridge-entry/bridge-matm-entries/" + bridgeMatmEntry.EntityData.SegmentPath
    bridgeMatmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeMatmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeMatmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeMatmEntry.EntityData.Children = types.NewOrderedMap()
    bridgeMatmEntry.EntityData.Leafs = types.NewOrderedMap()
    bridgeMatmEntry.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", bridgeMatmEntry.MacAddress})
    bridgeMatmEntry.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", bridgeMatmEntry.Interface})
    bridgeMatmEntry.EntityData.Leafs.Append("vlan", types.YLeaf{"Vlan", bridgeMatmEntry.Vlan})
    bridgeMatmEntry.EntityData.Leafs.Append("type", types.YLeaf{"Type", bridgeMatmEntry.Type})

    bridgeMatmEntry.EntityData.YListKeys = []string {"MacAddress"}

    return &(bridgeMatmEntry.EntityData)
}

