// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin dumper to
// configure file path options to copy the core files to.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_dumper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_dumper"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-dumper exception}", reflect.TypeOf(Exception{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-dumper:exception", reflect.TypeOf(Exception{}))
}

// Exception
type Exception struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Exception_Choice.
    Choice []*Exception_Choice
}

func (exception *Exception) GetEntityData() *types.CommonEntityData {
    exception.EntityData.YFilter = exception.YFilter
    exception.EntityData.YangName = "exception"
    exception.EntityData.BundleName = "cisco_ios_xr"
    exception.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-dumper"
    exception.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-dumper:exception"
    exception.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exception.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exception.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exception.EntityData.Children = types.NewOrderedMap()
    exception.EntityData.Children.Append("choice", types.YChild{"Choice", nil})
    for i := range exception.Choice {
        exception.EntityData.Children.Append(types.GetSegmentPath(exception.Choice[i]), types.YChild{"Choice", exception.Choice[i]})
    }
    exception.EntityData.Leafs = types.NewOrderedMap()

    exception.EntityData.YListKeys = []string {}

    return &(exception.EntityData)
}

// Exception_Choice
type Exception_Choice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..3.
    Order interface{}

    // The type is string. This attribute is mandatory.
    Filepath interface{}
}

func (choice *Exception_Choice) GetEntityData() *types.CommonEntityData {
    choice.EntityData.YFilter = choice.YFilter
    choice.EntityData.YangName = "choice"
    choice.EntityData.BundleName = "cisco_ios_xr"
    choice.EntityData.ParentYangName = "exception"
    choice.EntityData.SegmentPath = "choice" + types.AddKeyToken(choice.Order, "order")
    choice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    choice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    choice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    choice.EntityData.Children = types.NewOrderedMap()
    choice.EntityData.Leafs = types.NewOrderedMap()
    choice.EntityData.Leafs.Append("order", types.YLeaf{"Order", choice.Order})
    choice.EntityData.Leafs.Append("filepath", types.YLeaf{"Filepath", choice.Filepath})

    choice.EntityData.YListKeys = []string {"Order"}

    return &(choice.EntityData)
}

