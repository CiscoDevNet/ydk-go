// This module contains a collection of YANG definitions
// for Cisco IOS-XR accounting package configuration.
// 
// This module contains definitions
// for the following management objects:
//   accounting: Global Accounting configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package accounting_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package accounting_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-accounting-cfg accounting}", reflect.TypeOf(Accounting{}))
    ydk.RegisterEntity("Cisco-IOS-XR-accounting-cfg:accounting", reflect.TypeOf(Accounting{}))
}

// Accounting
// Global Accounting configuration commands
type Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Accounting. The type is interface{}.
    Enable interface{}

    // Interfaces configuration.
    Interfaces Accounting_Interfaces
}

func (accounting *Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "Cisco-IOS-XR-accounting-cfg"
    accounting.EntityData.SegmentPath = "Cisco-IOS-XR-accounting-cfg:accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &accounting.Interfaces})
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", accounting.Enable})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Accounting_Interfaces
// Interfaces configuration
type Accounting_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable accounting on Interfaces. The type is interface{}.
    Enable interface{}

    // Interfaces MPLS configuration.
    Mpls Accounting_Interfaces_Mpls

    // Interfaces Segment Routing configuration.
    SegmentRouting Accounting_Interfaces_SegmentRouting
}

func (interfaces *Accounting_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "accounting"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("mpls", types.YChild{"Mpls", &interfaces.Mpls})
    interfaces.EntityData.Children.Append("segment-routing", types.YChild{"SegmentRouting", &interfaces.SegmentRouting})
    interfaces.EntityData.Leafs = types.NewOrderedMap()
    interfaces.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaces.Enable})

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Accounting_Interfaces_Mpls
// Interfaces MPLS configuration
type Accounting_Interfaces_Mpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable accounting on MPLS. The type is interface{}.
    Enable interface{}

    // Enable accounting on MPLS IPv4 RSVP TE. The type is interface{}.
    EnableV4rsvpte interface{}
}

func (mpls *Accounting_Interfaces_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "cisco_ios_xr"
    mpls.EntityData.ParentYangName = "interfaces"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpls.EntityData.Children = types.NewOrderedMap()
    mpls.EntityData.Leafs = types.NewOrderedMap()
    mpls.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mpls.Enable})
    mpls.EntityData.Leafs.Append("enable-v4rsvpte", types.YLeaf{"EnableV4rsvpte", mpls.EnableV4rsvpte})

    mpls.EntityData.YListKeys = []string {}

    return &(mpls.EntityData)
}

// Accounting_Interfaces_SegmentRouting
// Interfaces Segment Routing configuration
type Accounting_Interfaces_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable accounting on Segment Routing. The type is interface{}.
    Enable interface{}

    // Enable accounting on Segment Routing MPLS IPv4. The type is interface{}.
    EnableMplsv4 interface{}

    // Enable accounting on Segment Routing MPLS IPv6. The type is interface{}.
    EnableMplsv6 interface{}
}

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "cisco_ios_xr"
    segmentRouting.EntityData.ParentYangName = "interfaces"
    segmentRouting.EntityData.SegmentPath = "segment-routing"
    segmentRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRouting.EntityData.Children = types.NewOrderedMap()
    segmentRouting.EntityData.Leafs = types.NewOrderedMap()
    segmentRouting.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", segmentRouting.Enable})
    segmentRouting.EntityData.Leafs.Append("enable-mplsv4", types.YLeaf{"EnableMplsv4", segmentRouting.EnableMplsv4})
    segmentRouting.EntityData.Leafs.Append("enable-mplsv6", types.YLeaf{"EnableMplsv6", segmentRouting.EnableMplsv6})

    segmentRouting.EntityData.YListKeys = []string {}

    return &(segmentRouting.EntityData)
}

