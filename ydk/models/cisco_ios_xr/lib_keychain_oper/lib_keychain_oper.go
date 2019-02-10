// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-keychain package operational data.
// 
// This module contains definitions
// for the following management objects:
//   keychain: Keychain operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    // CMAC AES 12 bytes
    CrytoAlgo_aes_128_cmac_96 CrytoAlgo = "aes-128-cmac-96"

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

    // HMAC SHA1 12 bytes
    CrytoAlgo_hmac_sha1_96 CrytoAlgo = "hmac-sha1-96"

    // HMAC SHA256 32 bytes
    CrytoAlgo_hmac_sha_256 CrytoAlgo = "hmac-sha-256"
)

// Keychain
// Keychain operational data
type Keychain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of configured key names.
    Keys Keychain_Keys
}

func (keychain *Keychain) GetEntityData() *types.CommonEntityData {
    keychain.EntityData.YFilter = keychain.YFilter
    keychain.EntityData.YangName = "keychain"
    keychain.EntityData.BundleName = "cisco_ios_xr"
    keychain.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-oper"
    keychain.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-oper:keychain"
    keychain.EntityData.AbsolutePath = keychain.EntityData.SegmentPath
    keychain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychain.EntityData.Children = types.NewOrderedMap()
    keychain.EntityData.Children.Append("keys", types.YChild{"Keys", &keychain.Keys})
    keychain.EntityData.Leafs = types.NewOrderedMap()

    keychain.EntityData.YListKeys = []string {}

    return &(keychain.EntityData)
}

// Keychain_Keys
// List of configured key names
type Keychain_Keys struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured key name. The type is slice of Keychain_Keys_Key.
    Key []*Keychain_Keys_Key
}

func (keys *Keychain_Keys) GetEntityData() *types.CommonEntityData {
    keys.EntityData.YFilter = keys.YFilter
    keys.EntityData.YangName = "keys"
    keys.EntityData.BundleName = "cisco_ios_xr"
    keys.EntityData.ParentYangName = "keychain"
    keys.EntityData.SegmentPath = "keys"
    keys.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-oper:keychain/" + keys.EntityData.SegmentPath
    keys.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keys.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keys.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keys.EntityData.Children = types.NewOrderedMap()
    keys.EntityData.Children.Append("key", types.YChild{"Key", nil})
    for i := range keys.Key {
        keys.EntityData.Children.Append(types.GetSegmentPath(keys.Key[i]), types.YChild{"Key", keys.Key[i]})
    }
    keys.EntityData.Leafs = types.NewOrderedMap()

    keys.EntityData.YListKeys = []string {}

    return &(keys.EntityData)
}

// Keychain_Keys_Key
// Configured key name
type Keychain_Keys_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Key name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    KeyName interface{}

    // Accept tolerance is infinite if value is 0xffffffff. The type is string.
    AcceptTolerance interface{}

    // Key properties.
    Key Keychain_Keys_Key_Key
}

func (key *Keychain_Keys_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "keys"
    key.EntityData.SegmentPath = "key" + types.AddKeyToken(key.KeyName, "key-name")
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-oper:keychain/keys/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("key", types.YChild{"Key", &key.Key})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("key-name", types.YLeaf{"KeyName", key.KeyName})
    key.EntityData.Leafs.Append("accept-tolerance", types.YLeaf{"AcceptTolerance", key.AcceptTolerance})

    key.EntityData.YListKeys = []string {"KeyName"}

    return &(key.EntityData)
}

// Keychain_Keys_Key_Key
// Key properties
type Keychain_Keys_Key_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key id. The type is slice of Keychain_Keys_Key_Key_KeyId.
    KeyId []*Keychain_Keys_Key_Key_KeyId
}

func (key *Keychain_Keys_Key_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "key"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-oper:keychain/keys/key/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("key-id", types.YChild{"KeyId", nil})
    for i := range key.KeyId {
        types.SetYListKey(key.KeyId[i], i)
        key.EntityData.Children.Append(types.GetSegmentPath(key.KeyId[i]), types.YChild{"KeyId", key.KeyId[i]})
    }
    key.EntityData.Leafs = types.NewOrderedMap()

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Keychain_Keys_Key_Key_KeyId
// key id
type Keychain_Keys_Key_Key_KeyId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Key string. The type is string.
    KeyString interface{}

    // Type of key encryption. The type is Enc.
    Type interface{}

    // Key ID. The type is string.
    KeyId interface{}

    // Cryptographic algorithm. The type is CrytoAlgo.
    CryptographicAlgorithm interface{}

    // To check if it's a macsec key.
    Macsec Keychain_Keys_Key_Key_KeyId_Macsec

    // Lifetime of the key.
    SendLifetime Keychain_Keys_Key_Key_KeyId_SendLifetime

    // Accept Lifetime of the key.
    AcceptLifetime Keychain_Keys_Key_Key_KeyId_AcceptLifetime
}

func (keyId *Keychain_Keys_Key_Key_KeyId) GetEntityData() *types.CommonEntityData {
    keyId.EntityData.YFilter = keyId.YFilter
    keyId.EntityData.YangName = "key-id"
    keyId.EntityData.BundleName = "cisco_ios_xr"
    keyId.EntityData.ParentYangName = "key"
    keyId.EntityData.SegmentPath = "key-id" + types.AddNoKeyToken(keyId)
    keyId.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-oper:keychain/keys/key/key/" + keyId.EntityData.SegmentPath
    keyId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyId.EntityData.Children = types.NewOrderedMap()
    keyId.EntityData.Children.Append("macsec", types.YChild{"Macsec", &keyId.Macsec})
    keyId.EntityData.Children.Append("send-lifetime", types.YChild{"SendLifetime", &keyId.SendLifetime})
    keyId.EntityData.Children.Append("accept-lifetime", types.YChild{"AcceptLifetime", &keyId.AcceptLifetime})
    keyId.EntityData.Leafs = types.NewOrderedMap()
    keyId.EntityData.Leafs.Append("key-string", types.YLeaf{"KeyString", keyId.KeyString})
    keyId.EntityData.Leafs.Append("type", types.YLeaf{"Type", keyId.Type})
    keyId.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", keyId.KeyId})
    keyId.EntityData.Leafs.Append("cryptographic-algorithm", types.YLeaf{"CryptographicAlgorithm", keyId.CryptographicAlgorithm})

    keyId.EntityData.YListKeys = []string {}

    return &(keyId.EntityData)
}

// Keychain_Keys_Key_Key_KeyId_Macsec
// To check if it's a macsec key
type Keychain_Keys_Key_Key_KeyId_Macsec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // To check if it's a macsec key. The type is bool.
    IsMacsecKey interface{}
}

func (macsec *Keychain_Keys_Key_Key_KeyId_Macsec) GetEntityData() *types.CommonEntityData {
    macsec.EntityData.YFilter = macsec.YFilter
    macsec.EntityData.YangName = "macsec"
    macsec.EntityData.BundleName = "cisco_ios_xr"
    macsec.EntityData.ParentYangName = "key-id"
    macsec.EntityData.SegmentPath = "macsec"
    macsec.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-oper:keychain/keys/key/key/key-id/" + macsec.EntityData.SegmentPath
    macsec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsec.EntityData.Children = types.NewOrderedMap()
    macsec.EntityData.Leafs = types.NewOrderedMap()
    macsec.EntityData.Leafs.Append("is-macsec-key", types.YLeaf{"IsMacsecKey", macsec.IsMacsecKey})

    macsec.EntityData.YListKeys = []string {}

    return &(macsec.EntityData)
}

// Keychain_Keys_Key_Key_KeyId_SendLifetime
// Lifetime of the key
type Keychain_Keys_Key_Key_KeyId_SendLifetime struct {
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

func (sendLifetime *Keychain_Keys_Key_Key_KeyId_SendLifetime) GetEntityData() *types.CommonEntityData {
    sendLifetime.EntityData.YFilter = sendLifetime.YFilter
    sendLifetime.EntityData.YangName = "send-lifetime"
    sendLifetime.EntityData.BundleName = "cisco_ios_xr"
    sendLifetime.EntityData.ParentYangName = "key-id"
    sendLifetime.EntityData.SegmentPath = "send-lifetime"
    sendLifetime.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-oper:keychain/keys/key/key/key-id/" + sendLifetime.EntityData.SegmentPath
    sendLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendLifetime.EntityData.Children = types.NewOrderedMap()
    sendLifetime.EntityData.Leafs = types.NewOrderedMap()
    sendLifetime.EntityData.Leafs.Append("start", types.YLeaf{"Start", sendLifetime.Start})
    sendLifetime.EntityData.Leafs.Append("end", types.YLeaf{"End", sendLifetime.End})
    sendLifetime.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", sendLifetime.Duration})
    sendLifetime.EntityData.Leafs.Append("is-always-valid", types.YLeaf{"IsAlwaysValid", sendLifetime.IsAlwaysValid})
    sendLifetime.EntityData.Leafs.Append("is-valid-now", types.YLeaf{"IsValidNow", sendLifetime.IsValidNow})

    sendLifetime.EntityData.YListKeys = []string {}

    return &(sendLifetime.EntityData)
}

// Keychain_Keys_Key_Key_KeyId_AcceptLifetime
// Accept Lifetime of the key
type Keychain_Keys_Key_Key_KeyId_AcceptLifetime struct {
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

func (acceptLifetime *Keychain_Keys_Key_Key_KeyId_AcceptLifetime) GetEntityData() *types.CommonEntityData {
    acceptLifetime.EntityData.YFilter = acceptLifetime.YFilter
    acceptLifetime.EntityData.YangName = "accept-lifetime"
    acceptLifetime.EntityData.BundleName = "cisco_ios_xr"
    acceptLifetime.EntityData.ParentYangName = "key-id"
    acceptLifetime.EntityData.SegmentPath = "accept-lifetime"
    acceptLifetime.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-oper:keychain/keys/key/key/key-id/" + acceptLifetime.EntityData.SegmentPath
    acceptLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptLifetime.EntityData.Children = types.NewOrderedMap()
    acceptLifetime.EntityData.Leafs = types.NewOrderedMap()
    acceptLifetime.EntityData.Leafs.Append("start", types.YLeaf{"Start", acceptLifetime.Start})
    acceptLifetime.EntityData.Leafs.Append("end", types.YLeaf{"End", acceptLifetime.End})
    acceptLifetime.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", acceptLifetime.Duration})
    acceptLifetime.EntityData.Leafs.Append("is-always-valid", types.YLeaf{"IsAlwaysValid", acceptLifetime.IsAlwaysValid})
    acceptLifetime.EntityData.Leafs.Append("is-valid-now", types.YLeaf{"IsValidNow", acceptLifetime.IsValidNow})

    acceptLifetime.EntityData.YListKeys = []string {}

    return &(acceptLifetime.EntityData)
}

