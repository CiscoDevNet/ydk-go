// This module contains a collection of YANG definitions
// for IOS-XE ARP operational data.
package arp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package arp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-arp-oper arp-data}", reflect.TypeOf(ArpData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-arp-oper:arp-data", reflect.TypeOf(ArpData{}))
}

// IosArpMode represents The mode that this entry is running in
type IosArpMode string

const (
    // Undefined - error
    IosArpMode_ios_arp_mode_null IosArpMode = "ios-arp-mode-null"

    // Entry has been learned
    IosArpMode_ios_arp_mode_dynamic IosArpMode = "ios-arp-mode-dynamic"

    // We've requested, but have no reply yet
    IosArpMode_ios_arp_mode_incomplete IosArpMode = "ios-arp-mode-incomplete"

    // Interface entry
    IosArpMode_ios_arp_mode_interface IosArpMode = "ios-arp-mode-interface"

    // Static Entry
    IosArpMode_ios_arp_mode_static IosArpMode = "ios-arp-mode-static"

    // Static - We're fronting this host
    IosArpMode_ios_arp_mode_alias IosArpMode = "ios-arp-mode-alias"

    // Simple Application ARP
    IosArpMode_ios_arp_mode_app_simple IosArpMode = "ios-arp-mode-app-simple"

    // Application Alias
    IosArpMode_ios_arp_mode_app_alias IosArpMode = "ios-arp-mode-app-alias"

    // Application Timer
    IosArpMode_ios_arp_mode_app_timer IosArpMode = "ios-arp-mode-app-timer"
)

// ArpData
// This module contains a collection of YANG definitions for
// monitoring the operation of IOS-XE ARP.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
type ArpData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of current VRFs. The type is slice of ArpData_ArpVrf.
    ArpVrf []*ArpData_ArpVrf
}

func (arpData *ArpData) GetEntityData() *types.CommonEntityData {
    arpData.EntityData.YFilter = arpData.YFilter
    arpData.EntityData.YangName = "arp-data"
    arpData.EntityData.BundleName = "cisco_ios_xe"
    arpData.EntityData.ParentYangName = "Cisco-IOS-XE-arp-oper"
    arpData.EntityData.SegmentPath = "Cisco-IOS-XE-arp-oper:arp-data"
    arpData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    arpData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    arpData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    arpData.EntityData.Children = types.NewOrderedMap()
    arpData.EntityData.Children.Append("arp-vrf", types.YChild{"ArpVrf", nil})
    for i := range arpData.ArpVrf {
        arpData.EntityData.Children.Append(types.GetSegmentPath(arpData.ArpVrf[i]), types.YChild{"ArpVrf", arpData.ArpVrf[i]})
    }
    arpData.EntityData.Leafs = types.NewOrderedMap()

    arpData.EntityData.YListKeys = []string {}

    return &(arpData.EntityData)
}

// ArpData_ArpVrf
// List of current VRFs
type ArpData_ArpVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name that this entry is tied to. The type is
    // string.
    Vrf interface{}

    // ARP entries associated with this VRF. The type is slice of
    // ArpData_ArpVrf_ArpOper.
    ArpOper []*ArpData_ArpVrf_ArpOper
}

func (arpVrf *ArpData_ArpVrf) GetEntityData() *types.CommonEntityData {
    arpVrf.EntityData.YFilter = arpVrf.YFilter
    arpVrf.EntityData.YangName = "arp-vrf"
    arpVrf.EntityData.BundleName = "cisco_ios_xe"
    arpVrf.EntityData.ParentYangName = "arp-data"
    arpVrf.EntityData.SegmentPath = "arp-vrf" + types.AddKeyToken(arpVrf.Vrf, "vrf")
    arpVrf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    arpVrf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    arpVrf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    arpVrf.EntityData.Children = types.NewOrderedMap()
    arpVrf.EntityData.Children.Append("arp-oper", types.YChild{"ArpOper", nil})
    for i := range arpVrf.ArpOper {
        arpVrf.EntityData.Children.Append(types.GetSegmentPath(arpVrf.ArpOper[i]), types.YChild{"ArpOper", arpVrf.ArpOper[i]})
    }
    arpVrf.EntityData.Leafs = types.NewOrderedMap()
    arpVrf.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", arpVrf.Vrf})

    arpVrf.EntityData.YListKeys = []string {"Vrf"}

    return &(arpVrf.EntityData)
}

// ArpData_ArpVrf_ArpOper
// ARP entries associated with this VRF
type ArpData_ArpVrf_ArpOper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. High level protocol address. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Protocol that produced the entry. The type is IosEncapsType.
    Enctype interface{}

    // Interface associated with this ARP entry. The type is string.
    Interface interface{}

    // Protocol that this ARP entry belongs to. The type is IosLinktype.
    Type interface{}

    // The mode that this entry is running in. The type is IosArpMode.
    Mode interface{}

    // Type of HW address. The type is IosSnpaType.
    Hwtype interface{}

    // hardware address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Hardware interface{}

    // Time of the last update. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Time interface{}
}

func (arpOper *ArpData_ArpVrf_ArpOper) GetEntityData() *types.CommonEntityData {
    arpOper.EntityData.YFilter = arpOper.YFilter
    arpOper.EntityData.YangName = "arp-oper"
    arpOper.EntityData.BundleName = "cisco_ios_xe"
    arpOper.EntityData.ParentYangName = "arp-vrf"
    arpOper.EntityData.SegmentPath = "arp-oper" + types.AddKeyToken(arpOper.Address, "address")
    arpOper.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    arpOper.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    arpOper.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    arpOper.EntityData.Children = types.NewOrderedMap()
    arpOper.EntityData.Leafs = types.NewOrderedMap()
    arpOper.EntityData.Leafs.Append("address", types.YLeaf{"Address", arpOper.Address})
    arpOper.EntityData.Leafs.Append("enctype", types.YLeaf{"Enctype", arpOper.Enctype})
    arpOper.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", arpOper.Interface})
    arpOper.EntityData.Leafs.Append("type", types.YLeaf{"Type", arpOper.Type})
    arpOper.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", arpOper.Mode})
    arpOper.EntityData.Leafs.Append("hwtype", types.YLeaf{"Hwtype", arpOper.Hwtype})
    arpOper.EntityData.Leafs.Append("hardware", types.YLeaf{"Hardware", arpOper.Hardware})
    arpOper.EntityData.Leafs.Append("time", types.YLeaf{"Time", arpOper.Time})

    arpOper.EntityData.YListKeys = []string {"Address"}

    return &(arpOper.EntityData)
}

