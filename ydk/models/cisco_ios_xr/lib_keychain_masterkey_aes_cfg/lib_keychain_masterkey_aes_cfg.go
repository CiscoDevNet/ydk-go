// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-keychain-masterkey-aes package configuration.
// 
// This module contains definitions
// for the following management objects:
//   password: Configure masterkey
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lib_keychain_masterkey_aes_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lib_keychain_masterkey_aes_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-keychain-masterkey-aes-cfg password}", reflect.TypeOf(Password{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-keychain-masterkey-aes-cfg:password", reflect.TypeOf(Password{}))
}

// Password
// Configure masterkey
type Password struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable password encryption.
    Encryption Password_Encryption
}

func (password *Password) GetEntityData() *types.CommonEntityData {
    password.EntityData.YFilter = password.YFilter
    password.EntityData.YangName = "password"
    password.EntityData.BundleName = "cisco_ios_xr"
    password.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-masterkey-aes-cfg"
    password.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-masterkey-aes-cfg:password"
    password.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    password.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    password.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    password.EntityData.Children = types.NewOrderedMap()
    password.EntityData.Children.Append("encryption", types.YChild{"Encryption", &password.Encryption})
    password.EntityData.Leafs = types.NewOrderedMap()

    password.EntityData.YListKeys = []string {}

    return &(password.EntityData)
}

// Password_Encryption
// Enable password encryption
type Password_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // encryption type used to store key. The type is interface{} with range:
    // 0..4294967295. The default value is 0.
    Aes interface{}
}

func (encryption *Password_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "password"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("aes", types.YLeaf{"Aes", encryption.Aes})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

