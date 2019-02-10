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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input MasterKeyAdd_Input
}

func (masterKeyAdd *MasterKeyAdd) GetEntityData() *types.CommonEntityData {
    masterKeyAdd.EntityData.YFilter = masterKeyAdd.YFilter
    masterKeyAdd.EntityData.YangName = "master-key-add"
    masterKeyAdd.EntityData.BundleName = "cisco_ios_xr"
    masterKeyAdd.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-act"
    masterKeyAdd.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-act:master-key-add"
    masterKeyAdd.EntityData.AbsolutePath = masterKeyAdd.EntityData.SegmentPath
    masterKeyAdd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masterKeyAdd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masterKeyAdd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masterKeyAdd.EntityData.Children = types.NewOrderedMap()
    masterKeyAdd.EntityData.Children.Append("input", types.YChild{"Input", &masterKeyAdd.Input})
    masterKeyAdd.EntityData.Leafs = types.NewOrderedMap()

    masterKeyAdd.EntityData.YListKeys = []string {}

    return &(masterKeyAdd.EntityData)
}

// MasterKeyAdd_Input
type MasterKeyAdd_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // New master key to be added. The type is string.
    NewKey interface{}
}

func (input *MasterKeyAdd_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "master-key-add"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-act:master-key-add/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("new-key", types.YLeaf{"NewKey", input.NewKey})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// MasterKeyDelete
// Remove Master key
type MasterKeyDelete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (masterKeyDelete *MasterKeyDelete) GetEntityData() *types.CommonEntityData {
    masterKeyDelete.EntityData.YFilter = masterKeyDelete.YFilter
    masterKeyDelete.EntityData.YangName = "master-key-delete"
    masterKeyDelete.EntityData.BundleName = "cisco_ios_xr"
    masterKeyDelete.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-act"
    masterKeyDelete.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-act:master-key-delete"
    masterKeyDelete.EntityData.AbsolutePath = masterKeyDelete.EntityData.SegmentPath
    masterKeyDelete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masterKeyDelete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masterKeyDelete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masterKeyDelete.EntityData.Children = types.NewOrderedMap()
    masterKeyDelete.EntityData.Leafs = types.NewOrderedMap()

    masterKeyDelete.EntityData.YListKeys = []string {}

    return &(masterKeyDelete.EntityData)
}

// MasterKeyUpdate
// To update master key
type MasterKeyUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input MasterKeyUpdate_Input
}

func (masterKeyUpdate *MasterKeyUpdate) GetEntityData() *types.CommonEntityData {
    masterKeyUpdate.EntityData.YFilter = masterKeyUpdate.YFilter
    masterKeyUpdate.EntityData.YangName = "master-key-update"
    masterKeyUpdate.EntityData.BundleName = "cisco_ios_xr"
    masterKeyUpdate.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-act"
    masterKeyUpdate.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-act:master-key-update"
    masterKeyUpdate.EntityData.AbsolutePath = masterKeyUpdate.EntityData.SegmentPath
    masterKeyUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masterKeyUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masterKeyUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masterKeyUpdate.EntityData.Children = types.NewOrderedMap()
    masterKeyUpdate.EntityData.Children.Append("input", types.YChild{"Input", &masterKeyUpdate.Input})
    masterKeyUpdate.EntityData.Leafs = types.NewOrderedMap()

    masterKeyUpdate.EntityData.YListKeys = []string {}

    return &(masterKeyUpdate.EntityData)
}

// MasterKeyUpdate_Input
type MasterKeyUpdate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key already added/key to be replaced. The type is string. This attribute is
    // mandatory.
    OldKey interface{}

    // New master key to be added . The type is string. This attribute is
    // mandatory.
    NewKey interface{}
}

func (input *MasterKeyUpdate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "master-key-update"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-act:master-key-update/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("old-key", types.YLeaf{"OldKey", input.OldKey})
    input.EntityData.Leafs.Append("new-key", types.YLeaf{"NewKey", input.NewKey})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

