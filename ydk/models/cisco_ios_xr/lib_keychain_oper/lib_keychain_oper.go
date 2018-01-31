// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-keychain package operational data.
// 
// This module contains definitions
// for the following management objects:
//   keychain: Keychain operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lib_keychain_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lib_keychain_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-keychain-oper keychain}", reflect.TypeOf(Keychain{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-keychain-oper:keychain", reflect.TypeOf(Keychain{}))
}

// Enc represents Type of password encryption
type Enc string

const (
    // Type 7 password type
    Enc_password_type7 Enc = "password-type7"

    // Type 6 Encryption
    Enc_password_type6 Enc = "password-type6"
)

// CrytoAlgo represents Cryptographic algorithm type
type CrytoAlgo string

const (
    // Not configured
    CrytoAlgo_not_configured CrytoAlgo = "not-configured"

    // HMAC SHA1 12 bytes
    CrytoAlgo_hmac_sha1_12 CrytoAlgo = "hmac-sha1-12"

    // MD5 16 bytes
    CrytoAlgo_md5 CrytoAlgo = "md5"

    // SHA1 20 bytes
    CrytoAlgo_sha1 CrytoAlgo = "sha1"

    // HMAC MD5 16 bytes
    CrytoAlgo_hmac_md5 CrytoAlgo = "hmac-md5"

    // HMAC SHA1 20 bytes
    CrytoAlgo_hmac_sha1_20 CrytoAlgo = "hmac-sha1-20"

    // CMAC AES 32 bytes
    CrytoAlgo_aes_128_cmac CrytoAlgo = "aes-128-cmac"

    // CMAC AES 64 bytes
    CrytoAlgo_aes_256_cmac CrytoAlgo = "aes-256-cmac"
)

// Keychain
// Keychain operational data
type Keychain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of configured key names.
    Keies Keychain_Keies
}

func (keychain *Keychain) GetFilter() yfilter.YFilter { return keychain.YFilter }

func (keychain *Keychain) SetFilter(yf yfilter.YFilter) { keychain.YFilter = yf }

func (keychain *Keychain) GetGoName(yname string) string {
    if yname == "keies" { return "Keies" }
    return ""
}

func (keychain *Keychain) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-keychain-oper:keychain"
}

func (keychain *Keychain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "keies" {
        return &keychain.Keies
    }
    return nil
}

func (keychain *Keychain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["keies"] = &keychain.Keies
    return children
}

func (keychain *Keychain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keychain *Keychain) GetBundleName() string { return "cisco_ios_xr" }

func (keychain *Keychain) GetYangName() string { return "keychain" }

func (keychain *Keychain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keychain *Keychain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keychain *Keychain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keychain *Keychain) SetParent(parent types.Entity) { keychain.parent = parent }

func (keychain *Keychain) GetParent() types.Entity { return keychain.parent }

func (keychain *Keychain) GetParentYangName() string { return "Cisco-IOS-XR-lib-keychain-oper" }

// Keychain_Keies
// List of configured key names
type Keychain_Keies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured key name. The type is slice of Keychain_Keies_Key.
    Key []Keychain_Keies_Key
}

func (keies *Keychain_Keies) GetFilter() yfilter.YFilter { return keies.YFilter }

func (keies *Keychain_Keies) SetFilter(yf yfilter.YFilter) { keies.YFilter = yf }

func (keies *Keychain_Keies) GetGoName(yname string) string {
    if yname == "key" { return "Key" }
    return ""
}

func (keies *Keychain_Keies) GetSegmentPath() string {
    return "keies"
}

func (keies *Keychain_Keies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "key" {
        for _, c := range keies.Key {
            if keies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Keychain_Keies_Key{}
        keies.Key = append(keies.Key, child)
        return &keies.Key[len(keies.Key)-1]
    }
    return nil
}

func (keies *Keychain_Keies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range keies.Key {
        children[keies.Key[i].GetSegmentPath()] = &keies.Key[i]
    }
    return children
}

func (keies *Keychain_Keies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keies *Keychain_Keies) GetBundleName() string { return "cisco_ios_xr" }

func (keies *Keychain_Keies) GetYangName() string { return "keies" }

func (keies *Keychain_Keies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keies *Keychain_Keies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keies *Keychain_Keies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keies *Keychain_Keies) SetParent(parent types.Entity) { keies.parent = parent }

func (keies *Keychain_Keies) GetParent() types.Entity { return keies.parent }

func (keies *Keychain_Keies) GetParentYangName() string { return "keychain" }

// Keychain_Keies_Key
// Configured key name
type Keychain_Keies_Key struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Key name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    KeyName interface{}

    // Accept tolerance is infinite if value is 0xffffffff. The type is string.
    AcceptTolerance interface{}

    // Key properties.
    Key Keychain_Keies_Key_Key
}

func (key *Keychain_Keies_Key) GetFilter() yfilter.YFilter { return key.YFilter }

func (key *Keychain_Keies_Key) SetFilter(yf yfilter.YFilter) { key.YFilter = yf }

func (key *Keychain_Keies_Key) GetGoName(yname string) string {
    if yname == "key-name" { return "KeyName" }
    if yname == "accept-tolerance" { return "AcceptTolerance" }
    if yname == "key" { return "Key" }
    return ""
}

func (key *Keychain_Keies_Key) GetSegmentPath() string {
    return "key" + "[key-name='" + fmt.Sprintf("%v", key.KeyName) + "']"
}

func (key *Keychain_Keies_Key) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "key" {
        return &key.Key
    }
    return nil
}

func (key *Keychain_Keies_Key) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["key"] = &key.Key
    return children
}

func (key *Keychain_Keies_Key) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-name"] = key.KeyName
    leafs["accept-tolerance"] = key.AcceptTolerance
    return leafs
}

func (key *Keychain_Keies_Key) GetBundleName() string { return "cisco_ios_xr" }

func (key *Keychain_Keies_Key) GetYangName() string { return "key" }

func (key *Keychain_Keies_Key) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (key *Keychain_Keies_Key) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (key *Keychain_Keies_Key) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (key *Keychain_Keies_Key) SetParent(parent types.Entity) { key.parent = parent }

func (key *Keychain_Keies_Key) GetParent() types.Entity { return key.parent }

func (key *Keychain_Keies_Key) GetParentYangName() string { return "keies" }

// Keychain_Keies_Key_Key
// Key properties
type Keychain_Keies_Key_Key struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // key id. The type is slice of Keychain_Keies_Key_Key_KeyId.
    KeyId []Keychain_Keies_Key_Key_KeyId
}

func (key *Keychain_Keies_Key_Key) GetFilter() yfilter.YFilter { return key.YFilter }

func (key *Keychain_Keies_Key_Key) SetFilter(yf yfilter.YFilter) { key.YFilter = yf }

func (key *Keychain_Keies_Key_Key) GetGoName(yname string) string {
    if yname == "key-id" { return "KeyId" }
    return ""
}

func (key *Keychain_Keies_Key_Key) GetSegmentPath() string {
    return "key"
}

func (key *Keychain_Keies_Key_Key) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "key-id" {
        for _, c := range key.KeyId {
            if key.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Keychain_Keies_Key_Key_KeyId{}
        key.KeyId = append(key.KeyId, child)
        return &key.KeyId[len(key.KeyId)-1]
    }
    return nil
}

func (key *Keychain_Keies_Key_Key) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range key.KeyId {
        children[key.KeyId[i].GetSegmentPath()] = &key.KeyId[i]
    }
    return children
}

func (key *Keychain_Keies_Key_Key) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (key *Keychain_Keies_Key_Key) GetBundleName() string { return "cisco_ios_xr" }

func (key *Keychain_Keies_Key_Key) GetYangName() string { return "key" }

func (key *Keychain_Keies_Key_Key) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (key *Keychain_Keies_Key_Key) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (key *Keychain_Keies_Key_Key) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (key *Keychain_Keies_Key_Key) SetParent(parent types.Entity) { key.parent = parent }

func (key *Keychain_Keies_Key_Key) GetParent() types.Entity { return key.parent }

func (key *Keychain_Keies_Key_Key) GetParentYangName() string { return "key" }

// Keychain_Keies_Key_Key_KeyId
// key id
type Keychain_Keies_Key_Key_KeyId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key string. The type is string.
    KeyString interface{}

    // Type of key encryption. The type is Enc.
    Type interface{}

    // Key ID. The type is interface{} with range: 0..18446744073709551615.
    KeyId interface{}

    // Cryptographic algorithm. The type is CrytoAlgo.
    CryptographicAlgorithm interface{}

    // To check if it's a macsec key.
    Macsec Keychain_Keies_Key_Key_KeyId_Macsec

    // Lifetime of the key.
    SendLifetime Keychain_Keies_Key_Key_KeyId_SendLifetime

    // Accept Lifetime of the key.
    AcceptLifetime Keychain_Keies_Key_Key_KeyId_AcceptLifetime
}

func (keyId *Keychain_Keies_Key_Key_KeyId) GetFilter() yfilter.YFilter { return keyId.YFilter }

func (keyId *Keychain_Keies_Key_Key_KeyId) SetFilter(yf yfilter.YFilter) { keyId.YFilter = yf }

func (keyId *Keychain_Keies_Key_Key_KeyId) GetGoName(yname string) string {
    if yname == "key-string" { return "KeyString" }
    if yname == "type" { return "Type" }
    if yname == "key-id" { return "KeyId" }
    if yname == "cryptographic-algorithm" { return "CryptographicAlgorithm" }
    if yname == "macsec" { return "Macsec" }
    if yname == "send-lifetime" { return "SendLifetime" }
    if yname == "accept-lifetime" { return "AcceptLifetime" }
    return ""
}

func (keyId *Keychain_Keies_Key_Key_KeyId) GetSegmentPath() string {
    return "key-id"
}

func (keyId *Keychain_Keies_Key_Key_KeyId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macsec" {
        return &keyId.Macsec
    }
    if childYangName == "send-lifetime" {
        return &keyId.SendLifetime
    }
    if childYangName == "accept-lifetime" {
        return &keyId.AcceptLifetime
    }
    return nil
}

func (keyId *Keychain_Keies_Key_Key_KeyId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["macsec"] = &keyId.Macsec
    children["send-lifetime"] = &keyId.SendLifetime
    children["accept-lifetime"] = &keyId.AcceptLifetime
    return children
}

func (keyId *Keychain_Keies_Key_Key_KeyId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-string"] = keyId.KeyString
    leafs["type"] = keyId.Type
    leafs["key-id"] = keyId.KeyId
    leafs["cryptographic-algorithm"] = keyId.CryptographicAlgorithm
    return leafs
}

func (keyId *Keychain_Keies_Key_Key_KeyId) GetBundleName() string { return "cisco_ios_xr" }

func (keyId *Keychain_Keies_Key_Key_KeyId) GetYangName() string { return "key-id" }

func (keyId *Keychain_Keies_Key_Key_KeyId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyId *Keychain_Keies_Key_Key_KeyId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyId *Keychain_Keies_Key_Key_KeyId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyId *Keychain_Keies_Key_Key_KeyId) SetParent(parent types.Entity) { keyId.parent = parent }

func (keyId *Keychain_Keies_Key_Key_KeyId) GetParent() types.Entity { return keyId.parent }

func (keyId *Keychain_Keies_Key_Key_KeyId) GetParentYangName() string { return "key" }

// Keychain_Keies_Key_Key_KeyId_Macsec
// To check if it's a macsec key
type Keychain_Keies_Key_Key_KeyId_Macsec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // To check if it's a macsec key. The type is bool.
    IsMacsecKey interface{}
}

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetFilter() yfilter.YFilter { return macsec.YFilter }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) SetFilter(yf yfilter.YFilter) { macsec.YFilter = yf }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetGoName(yname string) string {
    if yname == "is-macsec-key" { return "IsMacsecKey" }
    return ""
}

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetSegmentPath() string {
    return "macsec"
}

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-macsec-key"] = macsec.IsMacsecKey
    return leafs
}

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetBundleName() string { return "cisco_ios_xr" }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetYangName() string { return "macsec" }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) SetParent(parent types.Entity) { macsec.parent = parent }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetParent() types.Entity { return macsec.parent }

func (macsec *Keychain_Keies_Key_Key_KeyId_Macsec) GetParentYangName() string { return "key-id" }

// Keychain_Keies_Key_Key_KeyId_SendLifetime
// Lifetime of the key
type Keychain_Keies_Key_Key_KeyId_SendLifetime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key life start time in format : day-of-week month date-of-month HH:MM:SS
    // year eg: Thu Feb 1 18:32:14 2011. The type is string.
    Start interface{}

    // Key life end time in format : day-of-week month date-of-month HH:MM:SS year
    // eg: Thu Feb 1 18:32 :14 2011. The type is string.
    End interface{}

    // Duration of the key in seconds. value 0xffffffff reflects infinite, never
    // expires, is configured . The type is string. Units are second.
    Duration interface{}

    // Is TRUE if duration is 0xffffffff . The type is bool.
    IsAlwaysValid interface{}

    // Is TRUE if current time is betweenstart and end lifetime , else FALSE. The
    // type is bool.
    IsValidNow interface{}
}

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetFilter() yfilter.YFilter { return sendLifetime.YFilter }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) SetFilter(yf yfilter.YFilter) { sendLifetime.YFilter = yf }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "duration" { return "Duration" }
    if yname == "is-always-valid" { return "IsAlwaysValid" }
    if yname == "is-valid-now" { return "IsValidNow" }
    return ""
}

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetSegmentPath() string {
    return "send-lifetime"
}

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = sendLifetime.Start
    leafs["end"] = sendLifetime.End
    leafs["duration"] = sendLifetime.Duration
    leafs["is-always-valid"] = sendLifetime.IsAlwaysValid
    leafs["is-valid-now"] = sendLifetime.IsValidNow
    return leafs
}

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetBundleName() string { return "cisco_ios_xr" }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetYangName() string { return "send-lifetime" }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) SetParent(parent types.Entity) { sendLifetime.parent = parent }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetParent() types.Entity { return sendLifetime.parent }

func (sendLifetime *Keychain_Keies_Key_Key_KeyId_SendLifetime) GetParentYangName() string { return "key-id" }

// Keychain_Keies_Key_Key_KeyId_AcceptLifetime
// Accept Lifetime of the key
type Keychain_Keies_Key_Key_KeyId_AcceptLifetime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key life start time in format : day-of-week month date-of-month HH:MM:SS
    // year eg: Thu Feb 1 18:32:14 2011. The type is string.
    Start interface{}

    // Key life end time in format : day-of-week month date-of-month HH:MM:SS year
    // eg: Thu Feb 1 18:32 :14 2011. The type is string.
    End interface{}

    // Duration of the key in seconds. value 0xffffffff reflects infinite, never
    // expires, is configured . The type is string. Units are second.
    Duration interface{}

    // Is TRUE if duration is 0xffffffff . The type is bool.
    IsAlwaysValid interface{}

    // Is TRUE if current time is betweenstart and end lifetime , else FALSE. The
    // type is bool.
    IsValidNow interface{}
}

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetFilter() yfilter.YFilter { return acceptLifetime.YFilter }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) SetFilter(yf yfilter.YFilter) { acceptLifetime.YFilter = yf }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "duration" { return "Duration" }
    if yname == "is-always-valid" { return "IsAlwaysValid" }
    if yname == "is-valid-now" { return "IsValidNow" }
    return ""
}

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetSegmentPath() string {
    return "accept-lifetime"
}

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = acceptLifetime.Start
    leafs["end"] = acceptLifetime.End
    leafs["duration"] = acceptLifetime.Duration
    leafs["is-always-valid"] = acceptLifetime.IsAlwaysValid
    leafs["is-valid-now"] = acceptLifetime.IsValidNow
    return leafs
}

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetBundleName() string { return "cisco_ios_xr" }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetYangName() string { return "accept-lifetime" }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) SetParent(parent types.Entity) { acceptLifetime.parent = parent }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetParent() types.Entity { return acceptLifetime.parent }

func (acceptLifetime *Keychain_Keies_Key_Key_KeyId_AcceptLifetime) GetParentYangName() string { return "key-id" }

