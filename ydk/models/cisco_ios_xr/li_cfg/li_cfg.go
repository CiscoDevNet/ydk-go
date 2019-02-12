// This module contains a collection of YANG definitions
// for Cisco IOS-XR li package configuration.
// 
// This module contains definitions
// for the following management objects:
//   lawful-intercept: Lawful intercept configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package li_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package li_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-li-cfg lawful-intercept}", reflect.TypeOf(LawfulIntercept{}))
    ydk.RegisterEntity("Cisco-IOS-XR-li-cfg:lawful-intercept", reflect.TypeOf(LawfulIntercept{}))
}

// LawfulIntercept
// Lawful intercept configuration
type LawfulIntercept struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable lawful intercept feature. The type is interface{}.
    Disable interface{}
}

func (lawfulIntercept *LawfulIntercept) GetEntityData() *types.CommonEntityData {
    lawfulIntercept.EntityData.YFilter = lawfulIntercept.YFilter
    lawfulIntercept.EntityData.YangName = "lawful-intercept"
    lawfulIntercept.EntityData.BundleName = "cisco_ios_xr"
    lawfulIntercept.EntityData.ParentYangName = "Cisco-IOS-XR-li-cfg"
    lawfulIntercept.EntityData.SegmentPath = "Cisco-IOS-XR-li-cfg:lawful-intercept"
    lawfulIntercept.EntityData.AbsolutePath = lawfulIntercept.EntityData.SegmentPath
    lawfulIntercept.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lawfulIntercept.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lawfulIntercept.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lawfulIntercept.EntityData.Children = types.NewOrderedMap()
    lawfulIntercept.EntityData.Leafs = types.NewOrderedMap()
    lawfulIntercept.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", lawfulIntercept.Disable})

    lawfulIntercept.EntityData.YListKeys = []string {}

    return &(lawfulIntercept.EntityData)
}

