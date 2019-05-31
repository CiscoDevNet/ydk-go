// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2016, 2018 by Cisco Systems, Inc.
// All rights reserved.
package shellutil_delete_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package shellutil_delete_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-shellutil-delete-act delete}", reflect.TypeOf(Delete{}))
    ydk.RegisterEntity("Cisco-IOS-XR-shellutil-delete-act:delete", reflect.TypeOf(Delete{}))
}

// Delete
type Delete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Delete_Input

    
    Output Delete_Output
}

func (self *Delete) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "delete"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "Cisco-IOS-XR-shellutil-delete-act"
    self.EntityData.SegmentPath = "Cisco-IOS-XR-shellutil-delete-act:delete"
    self.EntityData.AbsolutePath = self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("input", types.YChild{"Input", &self.Input})
    self.EntityData.Children.Append("output", types.YChild{"Output", &self.Output})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Delete_Input
type Delete_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // file name. The type is string. This attribute is mandatory.
    Name interface{}

    // location. The type is string.
    Location interface{}

    // Files in dir. The type is bool.
    Recurse interface{}
}

func (input *Delete_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "delete"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-shellutil-delete-act:delete/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("name", types.YLeaf{"Name", input.Name})
    input.EntityData.Leafs.Append("location", types.YLeaf{"Location", input.Location})
    input.EntityData.Leafs.Append("recurse", types.YLeaf{"Recurse", input.Recurse})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Delete_Output
type Delete_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status of delete operation. The type is string.
    Response interface{}
}

func (output *Delete_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "delete"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-shellutil-delete-act:delete/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("response", types.YLeaf{"Response", output.Response})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

