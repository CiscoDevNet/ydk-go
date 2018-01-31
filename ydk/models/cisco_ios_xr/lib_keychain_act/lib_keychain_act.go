// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package lib_keychain_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lib_keychain_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-keychain-act master-key-add}", reflect.TypeOf(MasterKeyAdd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-keychain-act:master-key-add", reflect.TypeOf(MasterKeyAdd{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-keychain-act master-key-delete}", reflect.TypeOf(MasterKeyDelete{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-keychain-act:master-key-delete", reflect.TypeOf(MasterKeyDelete{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-keychain-act master-key-update}", reflect.TypeOf(MasterKeyUpdate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-keychain-act:master-key-update", reflect.TypeOf(MasterKeyUpdate{}))
}

// MasterKeyAdd
// To add a new master key
type MasterKeyAdd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input MasterKeyAdd_Input
}

func (masterKeyAdd *MasterKeyAdd) GetFilter() yfilter.YFilter { return masterKeyAdd.YFilter }

func (masterKeyAdd *MasterKeyAdd) SetFilter(yf yfilter.YFilter) { masterKeyAdd.YFilter = yf }

func (masterKeyAdd *MasterKeyAdd) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (masterKeyAdd *MasterKeyAdd) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-keychain-act:master-key-add"
}

func (masterKeyAdd *MasterKeyAdd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &masterKeyAdd.Input
    }
    return nil
}

func (masterKeyAdd *MasterKeyAdd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &masterKeyAdd.Input
    return children
}

func (masterKeyAdd *MasterKeyAdd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (masterKeyAdd *MasterKeyAdd) GetBundleName() string { return "cisco_ios_xr" }

func (masterKeyAdd *MasterKeyAdd) GetYangName() string { return "master-key-add" }

func (masterKeyAdd *MasterKeyAdd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (masterKeyAdd *MasterKeyAdd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (masterKeyAdd *MasterKeyAdd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (masterKeyAdd *MasterKeyAdd) SetParent(parent types.Entity) { masterKeyAdd.parent = parent }

func (masterKeyAdd *MasterKeyAdd) GetParent() types.Entity { return masterKeyAdd.parent }

func (masterKeyAdd *MasterKeyAdd) GetParentYangName() string { return "Cisco-IOS-XR-lib-keychain-act" }

// MasterKeyAdd_Input
type MasterKeyAdd_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // New master key to be added. The type is string.
    NewKey interface{}
}

func (input *MasterKeyAdd_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *MasterKeyAdd_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *MasterKeyAdd_Input) GetGoName(yname string) string {
    if yname == "new-key" { return "NewKey" }
    return ""
}

func (input *MasterKeyAdd_Input) GetSegmentPath() string {
    return "input"
}

func (input *MasterKeyAdd_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *MasterKeyAdd_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *MasterKeyAdd_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["new-key"] = input.NewKey
    return leafs
}

func (input *MasterKeyAdd_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *MasterKeyAdd_Input) GetYangName() string { return "input" }

func (input *MasterKeyAdd_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *MasterKeyAdd_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *MasterKeyAdd_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *MasterKeyAdd_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *MasterKeyAdd_Input) GetParent() types.Entity { return input.parent }

func (input *MasterKeyAdd_Input) GetParentYangName() string { return "master-key-add" }

// MasterKeyDelete
// Remove Master key
type MasterKeyDelete struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (masterKeyDelete *MasterKeyDelete) GetFilter() yfilter.YFilter { return masterKeyDelete.YFilter }

func (masterKeyDelete *MasterKeyDelete) SetFilter(yf yfilter.YFilter) { masterKeyDelete.YFilter = yf }

func (masterKeyDelete *MasterKeyDelete) GetGoName(yname string) string {
    return ""
}

func (masterKeyDelete *MasterKeyDelete) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-keychain-act:master-key-delete"
}

func (masterKeyDelete *MasterKeyDelete) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (masterKeyDelete *MasterKeyDelete) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (masterKeyDelete *MasterKeyDelete) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (masterKeyDelete *MasterKeyDelete) GetBundleName() string { return "cisco_ios_xr" }

func (masterKeyDelete *MasterKeyDelete) GetYangName() string { return "master-key-delete" }

func (masterKeyDelete *MasterKeyDelete) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (masterKeyDelete *MasterKeyDelete) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (masterKeyDelete *MasterKeyDelete) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (masterKeyDelete *MasterKeyDelete) SetParent(parent types.Entity) { masterKeyDelete.parent = parent }

func (masterKeyDelete *MasterKeyDelete) GetParent() types.Entity { return masterKeyDelete.parent }

func (masterKeyDelete *MasterKeyDelete) GetParentYangName() string { return "Cisco-IOS-XR-lib-keychain-act" }

// MasterKeyUpdate
// To update master key
type MasterKeyUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input MasterKeyUpdate_Input
}

func (masterKeyUpdate *MasterKeyUpdate) GetFilter() yfilter.YFilter { return masterKeyUpdate.YFilter }

func (masterKeyUpdate *MasterKeyUpdate) SetFilter(yf yfilter.YFilter) { masterKeyUpdate.YFilter = yf }

func (masterKeyUpdate *MasterKeyUpdate) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (masterKeyUpdate *MasterKeyUpdate) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-keychain-act:master-key-update"
}

func (masterKeyUpdate *MasterKeyUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &masterKeyUpdate.Input
    }
    return nil
}

func (masterKeyUpdate *MasterKeyUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &masterKeyUpdate.Input
    return children
}

func (masterKeyUpdate *MasterKeyUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (masterKeyUpdate *MasterKeyUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (masterKeyUpdate *MasterKeyUpdate) GetYangName() string { return "master-key-update" }

func (masterKeyUpdate *MasterKeyUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (masterKeyUpdate *MasterKeyUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (masterKeyUpdate *MasterKeyUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (masterKeyUpdate *MasterKeyUpdate) SetParent(parent types.Entity) { masterKeyUpdate.parent = parent }

func (masterKeyUpdate *MasterKeyUpdate) GetParent() types.Entity { return masterKeyUpdate.parent }

func (masterKeyUpdate *MasterKeyUpdate) GetParentYangName() string { return "Cisco-IOS-XR-lib-keychain-act" }

// MasterKeyUpdate_Input
type MasterKeyUpdate_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // key already added/key to be replaced. The type is string. This attribute is
    // mandatory.
    OldKey interface{}

    // New master key to be added . The type is string. This attribute is
    // mandatory.
    NewKey interface{}
}

func (input *MasterKeyUpdate_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *MasterKeyUpdate_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *MasterKeyUpdate_Input) GetGoName(yname string) string {
    if yname == "old-key" { return "OldKey" }
    if yname == "new-key" { return "NewKey" }
    return ""
}

func (input *MasterKeyUpdate_Input) GetSegmentPath() string {
    return "input"
}

func (input *MasterKeyUpdate_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *MasterKeyUpdate_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *MasterKeyUpdate_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["old-key"] = input.OldKey
    leafs["new-key"] = input.NewKey
    return leafs
}

func (input *MasterKeyUpdate_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *MasterKeyUpdate_Input) GetYangName() string { return "input" }

func (input *MasterKeyUpdate_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *MasterKeyUpdate_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *MasterKeyUpdate_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *MasterKeyUpdate_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *MasterKeyUpdate_Input) GetParent() types.Entity { return input.parent }

func (input *MasterKeyUpdate_Input) GetParentYangName() string { return "master-key-update" }

