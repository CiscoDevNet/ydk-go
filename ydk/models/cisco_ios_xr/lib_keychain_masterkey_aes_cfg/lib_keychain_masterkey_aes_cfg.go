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
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable password encryption.
    Encryption Password_Encryption
}

func (password *Password) GetFilter() yfilter.YFilter { return password.YFilter }

func (password *Password) SetFilter(yf yfilter.YFilter) { password.YFilter = yf }

func (password *Password) GetGoName(yname string) string {
    if yname == "encryption" { return "Encryption" }
    return ""
}

func (password *Password) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-keychain-masterkey-aes-cfg:password"
}

func (password *Password) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encryption" {
        return &password.Encryption
    }
    return nil
}

func (password *Password) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encryption"] = &password.Encryption
    return children
}

func (password *Password) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (password *Password) GetBundleName() string { return "cisco_ios_xr" }

func (password *Password) GetYangName() string { return "password" }

func (password *Password) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (password *Password) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (password *Password) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (password *Password) SetParent(parent types.Entity) { password.parent = parent }

func (password *Password) GetParent() types.Entity { return password.parent }

func (password *Password) GetParentYangName() string { return "Cisco-IOS-XR-lib-keychain-masterkey-aes-cfg" }

// Password_Encryption
// Enable password encryption
type Password_Encryption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // encryption type used to store key. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 0.
    Aes interface{}
}

func (encryption *Password_Encryption) GetFilter() yfilter.YFilter { return encryption.YFilter }

func (encryption *Password_Encryption) SetFilter(yf yfilter.YFilter) { encryption.YFilter = yf }

func (encryption *Password_Encryption) GetGoName(yname string) string {
    if yname == "aes" { return "Aes" }
    return ""
}

func (encryption *Password_Encryption) GetSegmentPath() string {
    return "encryption"
}

func (encryption *Password_Encryption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (encryption *Password_Encryption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (encryption *Password_Encryption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aes"] = encryption.Aes
    return leafs
}

func (encryption *Password_Encryption) GetBundleName() string { return "cisco_ios_xr" }

func (encryption *Password_Encryption) GetYangName() string { return "encryption" }

func (encryption *Password_Encryption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryption *Password_Encryption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryption *Password_Encryption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryption *Password_Encryption) SetParent(parent types.Entity) { encryption.parent = parent }

func (encryption *Password_Encryption) GetParent() types.Entity { return encryption.parent }

func (encryption *Password_Encryption) GetParentYangName() string { return "password" }

