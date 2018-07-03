// This module contains a collection of YANG definitions
// for Cisco IOS-XR fib-cef-lba package configuration.
// 
// This module contains definitions
// for the following management objects:
//   fiblb: FIB load-balancing
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fib_cef_lba_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fib_cef_lba_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fib-cef-lba-cfg fiblb}", reflect.TypeOf(Fiblb{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fib-cef-lba-cfg:fiblb", reflect.TypeOf(Fiblb{}))
}

// Fiblb
// FIB load-balancing
// This type is a presence type.
type Fiblb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Specify number of fields used for the load balancing. The type is
    // interface{} with range: 0..1. This attribute is mandatory.
    Fields interface{}

    // Payload Load-Balancing. The type is bool. This attribute is mandatory.
    Payload interface{}
}

func (fiblb *Fiblb) GetEntityData() *types.CommonEntityData {
    fiblb.EntityData.YFilter = fiblb.YFilter
    fiblb.EntityData.YangName = "fiblb"
    fiblb.EntityData.BundleName = "cisco_ios_xr"
    fiblb.EntityData.ParentYangName = "Cisco-IOS-XR-fib-cef-lba-cfg"
    fiblb.EntityData.SegmentPath = "Cisco-IOS-XR-fib-cef-lba-cfg:fiblb"
    fiblb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fiblb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fiblb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fiblb.EntityData.Children = types.NewOrderedMap()
    fiblb.EntityData.Leafs = types.NewOrderedMap()
    fiblb.EntityData.Leafs.Append("fields", types.YLeaf{"Fields", fiblb.Fields})
    fiblb.EntityData.Leafs.Append("payload", types.YLeaf{"Payload", fiblb.Payload})

    fiblb.EntityData.YListKeys = []string {}

    return &(fiblb.EntityData)
}

