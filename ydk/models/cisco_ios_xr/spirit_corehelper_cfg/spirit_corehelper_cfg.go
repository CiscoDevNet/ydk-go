// This module contains a collection of YANG definitions
// for Cisco IOS-XR spirit-corehelper package configuration.
// 
// This module contains definitions
// for the following management objects:
//   exception: Core dump configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package spirit_corehelper_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package spirit_corehelper_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-corehelper-cfg exception}", reflect.TypeOf(Exception{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-corehelper-cfg:exception", reflect.TypeOf(Exception{}))
}

// Exception
// Core dump configuration commands
type Exception struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Container for the order of preference.
    File Exception_File
}

func (exception *Exception) GetFilter() yfilter.YFilter { return exception.YFilter }

func (exception *Exception) SetFilter(yf yfilter.YFilter) { exception.YFilter = yf }

func (exception *Exception) GetGoName(yname string) string {
    if yname == "file" { return "File" }
    return ""
}

func (exception *Exception) GetSegmentPath() string {
    return "Cisco-IOS-XR-spirit-corehelper-cfg:exception"
}

func (exception *Exception) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "file" {
        return &exception.File
    }
    return nil
}

func (exception *Exception) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["file"] = &exception.File
    return children
}

func (exception *Exception) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exception *Exception) GetBundleName() string { return "cisco_ios_xr" }

func (exception *Exception) GetYangName() string { return "exception" }

func (exception *Exception) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exception *Exception) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exception *Exception) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exception *Exception) SetParent(parent types.Entity) { exception.parent = parent }

func (exception *Exception) GetParent() types.Entity { return exception.parent }

func (exception *Exception) GetParentYangName() string { return "Cisco-IOS-XR-spirit-corehelper-cfg" }

// Exception_File
// Container for the order of preference
type Exception_File struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Preference of the dump location. The type is string.
    Choice2 interface{}

    // Preference of the dump location. The type is string.
    Choice1 interface{}

    // Preference of the dump location. The type is string.
    Choice3 interface{}
}

func (file *Exception_File) GetFilter() yfilter.YFilter { return file.YFilter }

func (file *Exception_File) SetFilter(yf yfilter.YFilter) { file.YFilter = yf }

func (file *Exception_File) GetGoName(yname string) string {
    if yname == "choice2" { return "Choice2" }
    if yname == "choice1" { return "Choice1" }
    if yname == "choice3" { return "Choice3" }
    return ""
}

func (file *Exception_File) GetSegmentPath() string {
    return "file"
}

func (file *Exception_File) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (file *Exception_File) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (file *Exception_File) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice2"] = file.Choice2
    leafs["choice1"] = file.Choice1
    leafs["choice3"] = file.Choice3
    return leafs
}

func (file *Exception_File) GetBundleName() string { return "cisco_ios_xr" }

func (file *Exception_File) GetYangName() string { return "file" }

func (file *Exception_File) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (file *Exception_File) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (file *Exception_File) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (file *Exception_File) SetParent(parent types.Entity) { file.parent = parent }

func (file *Exception_File) GetParent() types.Entity { return file.parent }

func (file *Exception_File) GetParentYangName() string { return "exception" }

