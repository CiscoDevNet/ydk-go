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

// Enc represents Type of password encryption
type Enc string

const (
    // Type 7 password type
    Enc_password_type7 Enc = "password-type7"

    // Type 6 Encryption
    Enc_password_type6 Enc = "password-type6"
)

// Keychain
// Keychain operational data
type Keychain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of configured key names.
    Keies Keychain_Keies
}

func (keychain *Keychain) GetEntityData() *types.CommonEntityData {
    keychain.EntityData.YFilter = keychain.YFilter
    keychain.EntityData.YangName = "keychain"
    keychain.EntityData.BundleName = "cisco_ios_xr"
    keychain.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-oper"
    keychain.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-oper:keychain"
    keychain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychain.EntityData.Children = make(map[string]types.YChild)
    keychain.EntityData.Children["keies"] = types.YChild{"Keies", &keychain.Keies}
    keychain.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(keychain.EntityData)
}

// Keychain_Keies
// List of configured key names
type Keychain_Keies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured key name. The type is slice of Keychain_Keies_Key.
    Key []Keychain_Keies_Key
}

func (keies *Keychain_Keies) GetEntityData() *types.CommonEntityData {
    keies.EntityData.YFilter = keies.YFilter
    keies.EntityData.YangName = "keies"
    keies.EntityData.BundleName = "cisco_ios_xr"
    keies.EntityData.ParentYangName = "keychain"
    keies.EntityData.SegmentPath = "keies"
    keies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keies.EntityData.Children = make(map[string]types.YChild)
    keies.EntityData.Children["key"] = types.YChild{"Key", nil}
    for i := range keies.Key {
        keies.EntityData.Children[types.GetSegmentPath(&keies.Key[i])] = types.YChild{"Key", &keies.Key[i]}
    }
    keies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(keies.EntityData)
}

// Keychain_Keies_Key
// Configured key name
type Keychain_Keies_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Key name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    KeyName interface{}

    // Accept tolerance is infinite if value is 0xffffffff. The type is string.
    AcceptTolerance interface{}

    // Key properties.
    Key Keychain_Keies_Key_Key_
}

func (key *Keychain_Keies_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "keies"
    key.EntityData.SegmentPath = "key" + "[key-name='" + fmt.Sprintf("%v", key.KeyName) + "']"
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = make(map[string]types.YChild)
    key.EntityData.Children["key"] = types.YChild{"Key", &key.Key}
    key.EntityData.Leafs = make(map[string]types.YLeaf)
    key.EntityData.Leafs["key-name"] = types.YLeaf{"KeyName", key.KeyName}
    key.EntityData.Leafs["accept-tolerance"] = types.YLeaf{"AcceptTolerance", key.AcceptTolerance}
    return &(key.EntityData)
}

// Keychain_Keies_Key_Key_
// Key properties
type Keychain_Keies_Key_Key_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key id. The type is slice of Keychain_Keies_Key_Key__KeyId.
    KeyId []Keychain_Keies_Key_Key__KeyId
}

func (key_ *Keychain_Keies_Key_Key_) GetEntityData() *types.CommonEntityData {
    key_.EntityData.YFilter = key_.YFilter
    key_.EntityData.YangName = "key"
    key_.EntityData.BundleName = "cisco_ios_xr"
    key_.EntityData.ParentYangName = "key"
    key_.EntityData.SegmentPath = "key"
    key_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key_.EntityData.Children = make(map[string]types.YChild)
    key_.EntityData.Children["key-id"] = types.YChild{"KeyId", nil}
    for i := range key_.KeyId {
        key_.EntityData.Children[types.GetSegmentPath(&key_.KeyId[i])] = types.YChild{"KeyId", &key_.KeyId[i]}
    }
    key_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(key_.EntityData)
}

// Keychain_Keies_Key_Key__KeyId
// key id
type Keychain_Keies_Key_Key__KeyId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key string. The type is string.
    KeyString interface{}

    // Type of key encryption. The type is Enc.
    Type_ interface{}

    // Key ID. The type is string.
    KeyId interface{}

    // Cryptographic algorithm. The type is CrytoAlgo.
    CryptographicAlgorithm interface{}

    // To check if it's a macsec key.
    Macsec Keychain_Keies_Key_Key__KeyId_Macsec

    // Lifetime of the key.
    SendLifetime Keychain_Keies_Key_Key__KeyId_SendLifetime

    // Accept Lifetime of the key.
    AcceptLifetime Keychain_Keies_Key_Key__KeyId_AcceptLifetime
}

func (keyId *Keychain_Keies_Key_Key__KeyId) GetEntityData() *types.CommonEntityData {
    keyId.EntityData.YFilter = keyId.YFilter
    keyId.EntityData.YangName = "key-id"
    keyId.EntityData.BundleName = "cisco_ios_xr"
    keyId.EntityData.ParentYangName = "key"
    keyId.EntityData.SegmentPath = "key-id"
    keyId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyId.EntityData.Children = make(map[string]types.YChild)
    keyId.EntityData.Children["macsec"] = types.YChild{"Macsec", &keyId.Macsec}
    keyId.EntityData.Children["send-lifetime"] = types.YChild{"SendLifetime", &keyId.SendLifetime}
    keyId.EntityData.Children["accept-lifetime"] = types.YChild{"AcceptLifetime", &keyId.AcceptLifetime}
    keyId.EntityData.Leafs = make(map[string]types.YLeaf)
    keyId.EntityData.Leafs["key-string"] = types.YLeaf{"KeyString", keyId.KeyString}
    keyId.EntityData.Leafs["type"] = types.YLeaf{"Type_", keyId.Type_}
    keyId.EntityData.Leafs["key-id"] = types.YLeaf{"KeyId", keyId.KeyId}
    keyId.EntityData.Leafs["cryptographic-algorithm"] = types.YLeaf{"CryptographicAlgorithm", keyId.CryptographicAlgorithm}
    return &(keyId.EntityData)
}

// Keychain_Keies_Key_Key__KeyId_Macsec
// To check if it's a macsec key
type Keychain_Keies_Key_Key__KeyId_Macsec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // To check if it's a macsec key. The type is bool.
    IsMacsecKey interface{}
}

func (macsec *Keychain_Keies_Key_Key__KeyId_Macsec) GetEntityData() *types.CommonEntityData {
    macsec.EntityData.YFilter = macsec.YFilter
    macsec.EntityData.YangName = "macsec"
    macsec.EntityData.BundleName = "cisco_ios_xr"
    macsec.EntityData.ParentYangName = "key-id"
    macsec.EntityData.SegmentPath = "macsec"
    macsec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsec.EntityData.Children = make(map[string]types.YChild)
    macsec.EntityData.Leafs = make(map[string]types.YLeaf)
    macsec.EntityData.Leafs["is-macsec-key"] = types.YLeaf{"IsMacsecKey", macsec.IsMacsecKey}
    return &(macsec.EntityData)
}

// Keychain_Keies_Key_Key__KeyId_SendLifetime
// Lifetime of the key
type Keychain_Keies_Key_Key__KeyId_SendLifetime struct {
    EntityData types.CommonEntityData
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

func (sendLifetime *Keychain_Keies_Key_Key__KeyId_SendLifetime) GetEntityData() *types.CommonEntityData {
    sendLifetime.EntityData.YFilter = sendLifetime.YFilter
    sendLifetime.EntityData.YangName = "send-lifetime"
    sendLifetime.EntityData.BundleName = "cisco_ios_xr"
    sendLifetime.EntityData.ParentYangName = "key-id"
    sendLifetime.EntityData.SegmentPath = "send-lifetime"
    sendLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendLifetime.EntityData.Children = make(map[string]types.YChild)
    sendLifetime.EntityData.Leafs = make(map[string]types.YLeaf)
    sendLifetime.EntityData.Leafs["start"] = types.YLeaf{"Start", sendLifetime.Start}
    sendLifetime.EntityData.Leafs["end"] = types.YLeaf{"End", sendLifetime.End}
    sendLifetime.EntityData.Leafs["duration"] = types.YLeaf{"Duration", sendLifetime.Duration}
    sendLifetime.EntityData.Leafs["is-always-valid"] = types.YLeaf{"IsAlwaysValid", sendLifetime.IsAlwaysValid}
    sendLifetime.EntityData.Leafs["is-valid-now"] = types.YLeaf{"IsValidNow", sendLifetime.IsValidNow}
    return &(sendLifetime.EntityData)
}

// Keychain_Keies_Key_Key__KeyId_AcceptLifetime
// Accept Lifetime of the key
type Keychain_Keies_Key_Key__KeyId_AcceptLifetime struct {
    EntityData types.CommonEntityData
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

func (acceptLifetime *Keychain_Keies_Key_Key__KeyId_AcceptLifetime) GetEntityData() *types.CommonEntityData {
    acceptLifetime.EntityData.YFilter = acceptLifetime.YFilter
    acceptLifetime.EntityData.YangName = "accept-lifetime"
    acceptLifetime.EntityData.BundleName = "cisco_ios_xr"
    acceptLifetime.EntityData.ParentYangName = "key-id"
    acceptLifetime.EntityData.SegmentPath = "accept-lifetime"
    acceptLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptLifetime.EntityData.Children = make(map[string]types.YChild)
    acceptLifetime.EntityData.Leafs = make(map[string]types.YLeaf)
    acceptLifetime.EntityData.Leafs["start"] = types.YLeaf{"Start", acceptLifetime.Start}
    acceptLifetime.EntityData.Leafs["end"] = types.YLeaf{"End", acceptLifetime.End}
    acceptLifetime.EntityData.Leafs["duration"] = types.YLeaf{"Duration", acceptLifetime.Duration}
    acceptLifetime.EntityData.Leafs["is-always-valid"] = types.YLeaf{"IsAlwaysValid", acceptLifetime.IsAlwaysValid}
    acceptLifetime.EntityData.Leafs["is-valid-now"] = types.YLeaf{"IsValidNow", acceptLifetime.IsValidNow}
    return &(acceptLifetime.EntityData)
}

