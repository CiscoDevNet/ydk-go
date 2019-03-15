// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR Zapdisk module.
// 
// This module zapdisk is for factory card reset feature.
// 
// Copyright(c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_zapdisk

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_zapdisk"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-zapdisk zapdisk}", reflect.TypeOf(Zapdisk{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-zapdisk:zapdisk", reflect.TypeOf(Zapdisk{}))
}

// Zapdisk
type Zapdisk struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Zapdisk_Input

    
    Output Zapdisk_Output
}

func (zapdisk *Zapdisk) GetEntityData() *types.CommonEntityData {
    zapdisk.EntityData.YFilter = zapdisk.YFilter
    zapdisk.EntityData.YangName = "zapdisk"
    zapdisk.EntityData.BundleName = "cisco_ios_xr"
    zapdisk.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-zapdisk"
    zapdisk.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-zapdisk:zapdisk"
    zapdisk.EntityData.AbsolutePath = zapdisk.EntityData.SegmentPath
    zapdisk.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    zapdisk.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    zapdisk.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    zapdisk.EntityData.Children = types.NewOrderedMap()
    zapdisk.EntityData.Children.Append("input", types.YChild{"Input", &zapdisk.Input})
    zapdisk.EntityData.Children.Append("output", types.YChild{"Output", &zapdisk.Output})
    zapdisk.EntityData.Leafs = types.NewOrderedMap()

    zapdisk.EntityData.YListKeys = []string {}

    return &(zapdisk.EntityData)
}

// Zapdisk_Input
type Zapdisk_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Set interface{}

    
    CzapdiskUnset Zapdisk_Input_CzapdiskUnset
}

func (input *Zapdisk_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "zapdisk"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-zapdisk:zapdisk/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("czapdisk-unset", types.YChild{"CzapdiskUnset", &input.CzapdiskUnset})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("set", types.YLeaf{"Set", input.Set})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Zapdisk_Input_CzapdiskUnset
type Zapdisk_Input_CzapdiskUnset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Unset interface{}
}

func (czapdiskUnset *Zapdisk_Input_CzapdiskUnset) GetEntityData() *types.CommonEntityData {
    czapdiskUnset.EntityData.YFilter = czapdiskUnset.YFilter
    czapdiskUnset.EntityData.YangName = "czapdisk-unset"
    czapdiskUnset.EntityData.BundleName = "cisco_ios_xr"
    czapdiskUnset.EntityData.ParentYangName = "input"
    czapdiskUnset.EntityData.SegmentPath = "czapdisk-unset"
    czapdiskUnset.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-zapdisk:zapdisk/input/" + czapdiskUnset.EntityData.SegmentPath
    czapdiskUnset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    czapdiskUnset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    czapdiskUnset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    czapdiskUnset.EntityData.Children = types.NewOrderedMap()
    czapdiskUnset.EntityData.Leafs = types.NewOrderedMap()
    czapdiskUnset.EntityData.Leafs.Append("unset", types.YLeaf{"Unset", czapdiskUnset.Unset})

    czapdiskUnset.EntityData.YListKeys = []string {}

    return &(czapdiskUnset.EntityData)
}

// Zapdisk_Output
type Zapdisk_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string. This attribute is mandatory.
    Result interface{}
}

func (output *Zapdisk_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "zapdisk"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-zapdisk:zapdisk/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

