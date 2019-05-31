// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2016, 2018 by Cisco Systems, Inc.
// All rights reserved.
package shellutil_copy_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package shellutil_copy_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-shellutil-copy-act copy}", reflect.TypeOf(Copy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-shellutil-copy-act:copy", reflect.TypeOf(Copy{}))
}

// Copy
type Copy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Copy_Input

    
    Output Copy_Output
}

func (self *Copy) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "copy"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "Cisco-IOS-XR-shellutil-copy-act"
    self.EntityData.SegmentPath = "Cisco-IOS-XR-shellutil-copy-act:copy"
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

// Copy_Input
type Copy_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // source file name to copy. The type is string.
    Sourcename interface{}

    // destination file name. The type is string. This attribute is mandatory.
    Destinationname interface{}

    // source file system e.g disk0: tftp. The type is string. This attribute is
    // mandatory.
    Sourcefilesystem interface{}

    // destination file system e.g disk0:, tftp:. The type is string. This
    // attribute is mandatory.
    Destinationfilesystem interface{}

    // source location. The type is string.
    Sourcelocation interface{}

    // destination location. The type is string.
    Destinationlocation interface{}

    // vrf name. The type is string.
    Vrf interface{}

    // recurse files to copy. The type is bool.
    Recurse interface{}
}

func (input *Copy_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "copy"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-shellutil-copy-act:copy/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("sourcename", types.YLeaf{"Sourcename", input.Sourcename})
    input.EntityData.Leafs.Append("destinationname", types.YLeaf{"Destinationname", input.Destinationname})
    input.EntityData.Leafs.Append("sourcefilesystem", types.YLeaf{"Sourcefilesystem", input.Sourcefilesystem})
    input.EntityData.Leafs.Append("destinationfilesystem", types.YLeaf{"Destinationfilesystem", input.Destinationfilesystem})
    input.EntityData.Leafs.Append("sourcelocation", types.YLeaf{"Sourcelocation", input.Sourcelocation})
    input.EntityData.Leafs.Append("destinationlocation", types.YLeaf{"Destinationlocation", input.Destinationlocation})
    input.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", input.Vrf})
    input.EntityData.Leafs.Append("recurse", types.YLeaf{"Recurse", input.Recurse})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Copy_Output
type Copy_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status of copy operation. The type is string.
    Response interface{}
}

func (output *Copy_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "copy"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-shellutil-copy-act:copy/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("response", types.YLeaf{"Response", output.Response})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

