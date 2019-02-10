// This module contains a collection of YANG definitions for
// monitoring linecards in a Network Element.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package linecard_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package linecard_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-linecard-oper linecard-oper-data}", reflect.TypeOf(LinecardOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-linecard-oper:linecard-oper-data", reflect.TypeOf(LinecardOperData{}))
}

// LinecardOperData
// Top-level container for linecard operational data
type LinecardOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of linecard instances, keyed by slot. The type is slice of
    // LinecardOperData_Linecard.
    Linecard []*LinecardOperData_Linecard
}

func (linecardOperData *LinecardOperData) GetEntityData() *types.CommonEntityData {
    linecardOperData.EntityData.YFilter = linecardOperData.YFilter
    linecardOperData.EntityData.YangName = "linecard-oper-data"
    linecardOperData.EntityData.BundleName = "cisco_ios_xe"
    linecardOperData.EntityData.ParentYangName = "Cisco-IOS-XE-linecard-oper"
    linecardOperData.EntityData.SegmentPath = "Cisco-IOS-XE-linecard-oper:linecard-oper-data"
    linecardOperData.EntityData.AbsolutePath = linecardOperData.EntityData.SegmentPath
    linecardOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linecardOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linecardOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linecardOperData.EntityData.Children = types.NewOrderedMap()
    linecardOperData.EntityData.Children.Append("linecard", types.YChild{"Linecard", nil})
    for i := range linecardOperData.Linecard {
        linecardOperData.EntityData.Children.Append(types.GetSegmentPath(linecardOperData.Linecard[i]), types.YChild{"Linecard", linecardOperData.Linecard[i]})
    }
    linecardOperData.EntityData.Leafs = types.NewOrderedMap()

    linecardOperData.EntityData.YListKeys = []string {}

    return &(linecardOperData.EntityData)
}

// LinecardOperData_Linecard
// List of linecard instances, keyed by slot
type LinecardOperData_Linecard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Physical location description of the linecard. The
    // type is string.
    Name interface{}

    // Power provided to the linecard: Enabled/Disabled. The type is bool.
    PowerAdminState interface{}
}

func (linecard *LinecardOperData_Linecard) GetEntityData() *types.CommonEntityData {
    linecard.EntityData.YFilter = linecard.YFilter
    linecard.EntityData.YangName = "linecard"
    linecard.EntityData.BundleName = "cisco_ios_xe"
    linecard.EntityData.ParentYangName = "linecard-oper-data"
    linecard.EntityData.SegmentPath = "linecard" + types.AddKeyToken(linecard.Name, "name")
    linecard.EntityData.AbsolutePath = "Cisco-IOS-XE-linecard-oper:linecard-oper-data/" + linecard.EntityData.SegmentPath
    linecard.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linecard.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linecard.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linecard.EntityData.Children = types.NewOrderedMap()
    linecard.EntityData.Leafs = types.NewOrderedMap()
    linecard.EntityData.Leafs.Append("name", types.YLeaf{"Name", linecard.Name})
    linecard.EntityData.Leafs.Append("power-admin-state", types.YLeaf{"PowerAdminState", linecard.PowerAdminState})

    linecard.EntityData.YListKeys = []string {"Name"}

    return &(linecard.EntityData)
}

