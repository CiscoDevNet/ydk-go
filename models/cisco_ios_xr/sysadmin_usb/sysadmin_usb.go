// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all 'usb' commands for Sysadmin.
// 
// Copyright(c) 2018-2019 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_usb

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_usb"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-usb usb}", reflect.TypeOf(Usb{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-usb:usb", reflect.TypeOf(Usb{}))
}

// Usb
type Usb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (usb *Usb) GetEntityData() *types.CommonEntityData {
    usb.EntityData.YFilter = usb.YFilter
    usb.EntityData.YangName = "usb"
    usb.EntityData.BundleName = "cisco_ios_xr"
    usb.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-usb"
    usb.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-usb:usb"
    usb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usb.EntityData.Children = types.NewOrderedMap()
    usb.EntityData.Leafs = types.NewOrderedMap()

    usb.EntityData.YListKeys = []string {}

    return &(usb.EntityData)
}

