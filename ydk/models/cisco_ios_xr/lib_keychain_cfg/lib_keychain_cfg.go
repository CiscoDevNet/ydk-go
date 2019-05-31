// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-keychain package configuration.
// 
// This module contains definitions
// for the following management objects:
//   keychains: Configure a Key Chain
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package lib_keychain_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lib_keychain_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-keychain-cfg keychains}", reflect.TypeOf(Keychains{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-keychain-cfg:keychains", reflect.TypeOf(Keychains{}))
}

// MacsecEncryption represents Macsec encryption
type MacsecEncryption string

const (
    // Type7
    MacsecEncryption_type7 MacsecEncryption = "type7"

    // Type6
    MacsecEncryption_type6 MacsecEncryption = "type6"
)

// CryptoAlg represents Crypto alg
type CryptoAlg string

const (
    // AES 128 CMAC 96
    CryptoAlg_alg_aes_128_cmac_96 CryptoAlg = "alg-aes-128-cmac-96"

    // HMAC SHA 1 12
    CryptoAlg_alg_hmac_sha1_12 CryptoAlg = "alg-hmac-sha1-12"

    // MD5 16
    CryptoAlg_alg_md5_16 CryptoAlg = "alg-md5-16"

    // SHA 1 20
    CryptoAlg_alg_sha1_20 CryptoAlg = "alg-sha1-20"

    // HMAC MD5 16
    CryptoAlg_alg_hmac_md5_16 CryptoAlg = "alg-hmac-md5-16"

    // HMAC SHA 1 20
    CryptoAlg_alg_hmac_sha1_20 CryptoAlg = "alg-hmac-sha1-20"

    // HMAC SHA 1 96
    CryptoAlg_alg_hmac_sha1_96 CryptoAlg = "alg-hmac-sha1-96"

    // HMAC SHA 256
    CryptoAlg_alg_hmac_sha_256 CryptoAlg = "alg-hmac-sha-256"
)

// MacsecCryptoAlg represents Macsec crypto alg
type MacsecCryptoAlg string

const (
    // aes 128 cmac
    MacsecCryptoAlg_aes_128_cmac MacsecCryptoAlg = "aes-128-cmac"

    // aes 256 cmac
    MacsecCryptoAlg_aes_256_cmac MacsecCryptoAlg = "aes-256-cmac"
)

// KeyChainMonth represents Key chain month
type KeyChainMonth string

const (
    // January
    KeyChainMonth_jan KeyChainMonth = "jan"

    // February
    KeyChainMonth_feb KeyChainMonth = "feb"

    // March
    KeyChainMonth_mar KeyChainMonth = "mar"

    // April
    KeyChainMonth_apr KeyChainMonth = "apr"

    // May
    KeyChainMonth_may KeyChainMonth = "may"

    // June
    KeyChainMonth_jun KeyChainMonth = "jun"

    // July
    KeyChainMonth_jul KeyChainMonth = "jul"

    // August
    KeyChainMonth_aug KeyChainMonth = "aug"

    // September
    KeyChainMonth_sep KeyChainMonth = "sep"

    // October
    KeyChainMonth_oct KeyChainMonth = "oct"

    // November
    KeyChainMonth_nov KeyChainMonth = "nov"

    // December
    KeyChainMonth_dec KeyChainMonth = "dec"
)

// Keychains
// Configure a Key Chain
type Keychains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the key chain. The type is slice of Keychains_Keychain.
    Keychain []*Keychains_Keychain
}

func (keychains *Keychains) GetEntityData() *types.CommonEntityData {
    keychains.EntityData.YFilter = keychains.YFilter
    keychains.EntityData.YangName = "keychains"
    keychains.EntityData.BundleName = "cisco_ios_xr"
    keychains.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-cfg"
    keychains.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-cfg:keychains"
    keychains.EntityData.AbsolutePath = keychains.EntityData.SegmentPath
    keychains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychains.EntityData.Children = types.NewOrderedMap()
    keychains.EntityData.Children.Append("keychain", types.YChild{"Keychain", nil})
    for i := range keychains.Keychain {
        keychains.EntityData.Children.Append(types.GetSegmentPath(keychains.Keychain[i]), types.YChild{"Keychain", keychains.Keychain[i]})
    }
    keychains.EntityData.Leafs = types.NewOrderedMap()

    keychains.EntityData.YListKeys = []string {}

    return &(keychains.EntityData)
}

// Keychains_Keychain
// Name of the key chain
type Keychains_Keychain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the key chain. The type is string with
    // length: 1..32.
    ChainName interface{}

    // Accept Tolerance in seconds or infinite.
    AcceptTolerance Keychains_Keychain_AcceptTolerance

    // Name of the key chain for MACSec.
    MacsecKeychain Keychains_Keychain_MacsecKeychain

    // Configure a Key.
    Keys Keychains_Keychain_Keys
}

func (keychain *Keychains_Keychain) GetEntityData() *types.CommonEntityData {
    keychain.EntityData.YFilter = keychain.YFilter
    keychain.EntityData.YangName = "keychain"
    keychain.EntityData.BundleName = "cisco_ios_xr"
    keychain.EntityData.ParentYangName = "keychains"
    keychain.EntityData.SegmentPath = "keychain" + types.AddKeyToken(keychain.ChainName, "chain-name")
    keychain.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/" + keychain.EntityData.SegmentPath
    keychain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychain.EntityData.Children = types.NewOrderedMap()
    keychain.EntityData.Children.Append("accept-tolerance", types.YChild{"AcceptTolerance", &keychain.AcceptTolerance})
    keychain.EntityData.Children.Append("macsec-keychain", types.YChild{"MacsecKeychain", &keychain.MacsecKeychain})
    keychain.EntityData.Children.Append("keys", types.YChild{"Keys", &keychain.Keys})
    keychain.EntityData.Leafs = types.NewOrderedMap()
    keychain.EntityData.Leafs.Append("chain-name", types.YLeaf{"ChainName", keychain.ChainName})

    keychain.EntityData.YListKeys = []string {"ChainName"}

    return &(keychain.EntityData)
}

// Keychains_Keychain_AcceptTolerance
// Accept Tolerance in seconds or infinite
type Keychains_Keychain_AcceptTolerance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value in seconds. The type is interface{} with range: 1..8640000. Units are
    // second.
    Value interface{}

    // Infinite tolerance. The type is bool.
    Infinite interface{}
}

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetEntityData() *types.CommonEntityData {
    acceptTolerance.EntityData.YFilter = acceptTolerance.YFilter
    acceptTolerance.EntityData.YangName = "accept-tolerance"
    acceptTolerance.EntityData.BundleName = "cisco_ios_xr"
    acceptTolerance.EntityData.ParentYangName = "keychain"
    acceptTolerance.EntityData.SegmentPath = "accept-tolerance"
    acceptTolerance.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/" + acceptTolerance.EntityData.SegmentPath
    acceptTolerance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptTolerance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptTolerance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptTolerance.EntityData.Children = types.NewOrderedMap()
    acceptTolerance.EntityData.Leafs = types.NewOrderedMap()
    acceptTolerance.EntityData.Leafs.Append("value", types.YLeaf{"Value", acceptTolerance.Value})
    acceptTolerance.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", acceptTolerance.Infinite})

    acceptTolerance.EntityData.YListKeys = []string {}

    return &(acceptTolerance.EntityData)
}

// Keychains_Keychain_MacsecKeychain
// Name of the key chain for MACSec
type Keychains_Keychain_MacsecKeychain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a Key.
    MacsecKeys Keychains_Keychain_MacsecKeychain_MacsecKeys
}

func (macsecKeychain *Keychains_Keychain_MacsecKeychain) GetEntityData() *types.CommonEntityData {
    macsecKeychain.EntityData.YFilter = macsecKeychain.YFilter
    macsecKeychain.EntityData.YangName = "macsec-keychain"
    macsecKeychain.EntityData.BundleName = "cisco_ios_xr"
    macsecKeychain.EntityData.ParentYangName = "keychain"
    macsecKeychain.EntityData.SegmentPath = "macsec-keychain"
    macsecKeychain.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/" + macsecKeychain.EntityData.SegmentPath
    macsecKeychain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKeychain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKeychain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKeychain.EntityData.Children = types.NewOrderedMap()
    macsecKeychain.EntityData.Children.Append("macsec-keys", types.YChild{"MacsecKeys", &macsecKeychain.MacsecKeys})
    macsecKeychain.EntityData.Leafs = types.NewOrderedMap()

    macsecKeychain.EntityData.YListKeys = []string {}

    return &(macsecKeychain.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeys
// Configure a Key
type Keychains_Keychain_MacsecKeychain_MacsecKeys struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key Identifier. The type is slice of
    // Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey.
    MacsecKey []*Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey
}

func (macsecKeys *Keychains_Keychain_MacsecKeychain_MacsecKeys) GetEntityData() *types.CommonEntityData {
    macsecKeys.EntityData.YFilter = macsecKeys.YFilter
    macsecKeys.EntityData.YangName = "macsec-keys"
    macsecKeys.EntityData.BundleName = "cisco_ios_xr"
    macsecKeys.EntityData.ParentYangName = "macsec-keychain"
    macsecKeys.EntityData.SegmentPath = "macsec-keys"
    macsecKeys.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/macsec-keychain/" + macsecKeys.EntityData.SegmentPath
    macsecKeys.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKeys.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKeys.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKeys.EntityData.Children = types.NewOrderedMap()
    macsecKeys.EntityData.Children.Append("macsec-key", types.YChild{"MacsecKey", nil})
    for i := range macsecKeys.MacsecKey {
        macsecKeys.EntityData.Children.Append(types.GetSegmentPath(macsecKeys.MacsecKey[i]), types.YChild{"MacsecKey", macsecKeys.MacsecKey[i]})
    }
    macsecKeys.EntityData.Leafs = types.NewOrderedMap()

    macsecKeys.EntityData.YListKeys = []string {}

    return &(macsecKeys.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey
// Key Identifier
type Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Enter CKN as non-zero hex string of even length,
    // length range: <02-64>, i.e 32 bytes of MACsec CKN. The type is string with
    // length: 2..64.
    KeyId interface{}

    // Configure a key Lifetime.
    MacsecLifetime Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecLifetime

    // Configure a clear text/encrypted Key string along with cryptographic
    // algorithm.
    MacsecKeyString Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecKeyString
}

func (macsecKey *Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey) GetEntityData() *types.CommonEntityData {
    macsecKey.EntityData.YFilter = macsecKey.YFilter
    macsecKey.EntityData.YangName = "macsec-key"
    macsecKey.EntityData.BundleName = "cisco_ios_xr"
    macsecKey.EntityData.ParentYangName = "macsec-keys"
    macsecKey.EntityData.SegmentPath = "macsec-key" + types.AddKeyToken(macsecKey.KeyId, "key-id")
    macsecKey.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/macsec-keychain/macsec-keys/" + macsecKey.EntityData.SegmentPath
    macsecKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKey.EntityData.Children = types.NewOrderedMap()
    macsecKey.EntityData.Children.Append("macsec-lifetime", types.YChild{"MacsecLifetime", &macsecKey.MacsecLifetime})
    macsecKey.EntityData.Children.Append("macsec-key-string", types.YChild{"MacsecKeyString", &macsecKey.MacsecKeyString})
    macsecKey.EntityData.Leafs = types.NewOrderedMap()
    macsecKey.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", macsecKey.KeyId})

    macsecKey.EntityData.YListKeys = []string {"KeyId"}

    return &(macsecKey.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecLifetime
// Configure a key Lifetime
// This type is a presence type.
type Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Start Hour. The type is interface{} with range: 0..23. This attribute is
    // mandatory.
    StartHour interface{}

    // Start Minutes. The type is interface{} with range: 0..59. This attribute is
    // mandatory. Units are minute.
    StartMinutes interface{}

    // Start Seconds. The type is interface{} with range: 0..59. This attribute is
    // mandatory. Units are second.
    StartSeconds interface{}

    // Start Date. The type is interface{} with range: 1..31. This attribute is
    // mandatory.
    StartDate interface{}

    // Start Month. The type is KeyChainMonth. This attribute is mandatory.
    StartMonth interface{}

    // Start Year. The type is interface{} with range: 1993..2035. This attribute
    // is mandatory.
    StartYear interface{}

    // Lifetime duration in seconds. The type is interface{} with range:
    // 1..2147483647. Units are second.
    LifeTime interface{}

    // Infinite Lifetime flag. The type is bool.
    InfiniteFlag interface{}

    // End Hour. The type is interface{} with range: 0..23.
    EndHour interface{}

    // End Minutes. The type is interface{} with range: 0..59. Units are minute.
    EndMinutes interface{}

    // End Seconds. The type is interface{} with range: 0..59. Units are second.
    EndSeconds interface{}

    // End Date. The type is interface{} with range: 1..31.
    EndDate interface{}

    // End Month. The type is KeyChainMonth.
    EndMonth interface{}

    // End Year. The type is interface{} with range: 1993..2035.
    EndYear interface{}
}

func (macsecLifetime *Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecLifetime) GetEntityData() *types.CommonEntityData {
    macsecLifetime.EntityData.YFilter = macsecLifetime.YFilter
    macsecLifetime.EntityData.YangName = "macsec-lifetime"
    macsecLifetime.EntityData.BundleName = "cisco_ios_xr"
    macsecLifetime.EntityData.ParentYangName = "macsec-key"
    macsecLifetime.EntityData.SegmentPath = "macsec-lifetime"
    macsecLifetime.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/macsec-keychain/macsec-keys/macsec-key/" + macsecLifetime.EntityData.SegmentPath
    macsecLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecLifetime.EntityData.Children = types.NewOrderedMap()
    macsecLifetime.EntityData.Leafs = types.NewOrderedMap()
    macsecLifetime.EntityData.Leafs.Append("start-hour", types.YLeaf{"StartHour", macsecLifetime.StartHour})
    macsecLifetime.EntityData.Leafs.Append("start-minutes", types.YLeaf{"StartMinutes", macsecLifetime.StartMinutes})
    macsecLifetime.EntityData.Leafs.Append("start-seconds", types.YLeaf{"StartSeconds", macsecLifetime.StartSeconds})
    macsecLifetime.EntityData.Leafs.Append("start-date", types.YLeaf{"StartDate", macsecLifetime.StartDate})
    macsecLifetime.EntityData.Leafs.Append("start-month", types.YLeaf{"StartMonth", macsecLifetime.StartMonth})
    macsecLifetime.EntityData.Leafs.Append("start-year", types.YLeaf{"StartYear", macsecLifetime.StartYear})
    macsecLifetime.EntityData.Leafs.Append("life-time", types.YLeaf{"LifeTime", macsecLifetime.LifeTime})
    macsecLifetime.EntityData.Leafs.Append("infinite-flag", types.YLeaf{"InfiniteFlag", macsecLifetime.InfiniteFlag})
    macsecLifetime.EntityData.Leafs.Append("end-hour", types.YLeaf{"EndHour", macsecLifetime.EndHour})
    macsecLifetime.EntityData.Leafs.Append("end-minutes", types.YLeaf{"EndMinutes", macsecLifetime.EndMinutes})
    macsecLifetime.EntityData.Leafs.Append("end-seconds", types.YLeaf{"EndSeconds", macsecLifetime.EndSeconds})
    macsecLifetime.EntityData.Leafs.Append("end-date", types.YLeaf{"EndDate", macsecLifetime.EndDate})
    macsecLifetime.EntityData.Leafs.Append("end-month", types.YLeaf{"EndMonth", macsecLifetime.EndMonth})
    macsecLifetime.EntityData.Leafs.Append("end-year", types.YLeaf{"EndYear", macsecLifetime.EndYear})

    macsecLifetime.EntityData.YListKeys = []string {}

    return &(macsecLifetime.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecKeyString
// Configure a clear text/encrypted Key string
// along with cryptographic algorithm
// This type is a presence type.
type Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecKeyString struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Key String. The type is string with pattern: b'(!.+)|([^!].+)'. This
    // attribute is mandatory.
    String interface{}

    // Cryptographic Algorithm. The type is MacsecCryptoAlg. This attribute is
    // mandatory.
    CryptographicAlgorithm interface{}

    // encryption type used to store key. The type is MacsecEncryption. The
    // default value is type7.
    EncryptionType interface{}
}

func (macsecKeyString *Keychains_Keychain_MacsecKeychain_MacsecKeys_MacsecKey_MacsecKeyString) GetEntityData() *types.CommonEntityData {
    macsecKeyString.EntityData.YFilter = macsecKeyString.YFilter
    macsecKeyString.EntityData.YangName = "macsec-key-string"
    macsecKeyString.EntityData.BundleName = "cisco_ios_xr"
    macsecKeyString.EntityData.ParentYangName = "macsec-key"
    macsecKeyString.EntityData.SegmentPath = "macsec-key-string"
    macsecKeyString.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/macsec-keychain/macsec-keys/macsec-key/" + macsecKeyString.EntityData.SegmentPath
    macsecKeyString.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKeyString.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKeyString.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKeyString.EntityData.Children = types.NewOrderedMap()
    macsecKeyString.EntityData.Leafs = types.NewOrderedMap()
    macsecKeyString.EntityData.Leafs.Append("string", types.YLeaf{"String", macsecKeyString.String})
    macsecKeyString.EntityData.Leafs.Append("cryptographic-algorithm", types.YLeaf{"CryptographicAlgorithm", macsecKeyString.CryptographicAlgorithm})
    macsecKeyString.EntityData.Leafs.Append("encryption-type", types.YLeaf{"EncryptionType", macsecKeyString.EncryptionType})

    macsecKeyString.EntityData.YListKeys = []string {}

    return &(macsecKeyString.EntityData)
}

// Keychains_Keychain_Keys
// Configure a Key
type Keychains_Keychain_Keys struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key Identifier. The type is slice of Keychains_Keychain_Keys_Key.
    Key []*Keychains_Keychain_Keys_Key
}

func (keys *Keychains_Keychain_Keys) GetEntityData() *types.CommonEntityData {
    keys.EntityData.YFilter = keys.YFilter
    keys.EntityData.YangName = "keys"
    keys.EntityData.BundleName = "cisco_ios_xr"
    keys.EntityData.ParentYangName = "keychain"
    keys.EntityData.SegmentPath = "keys"
    keys.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/" + keys.EntityData.SegmentPath
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

// Keychains_Keychain_Keys_Key
// Key Identifier
type Keychains_Keychain_Keys_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 48-bit Key identifier. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    KeyId interface{}

    // Configure a clear text/encrypted Key string . The type is string with
    // pattern: b'(!.+)|([^!].+)'.
    KeyString interface{}

    // Configure the cryptographic algorithm. The type is CryptoAlg.
    CryptographicAlgorithm interface{}

    // Configure a key Acceptance Lifetime.
    AcceptLifetime Keychains_Keychain_Keys_Key_AcceptLifetime

    // Configure a Send Lifetime.
    SendLifetime Keychains_Keychain_Keys_Key_SendLifetime
}

func (key *Keychains_Keychain_Keys_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "keys"
    key.EntityData.SegmentPath = "key" + types.AddKeyToken(key.KeyId, "key-id")
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/keys/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("accept-lifetime", types.YChild{"AcceptLifetime", &key.AcceptLifetime})
    key.EntityData.Children.Append("send-lifetime", types.YChild{"SendLifetime", &key.SendLifetime})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", key.KeyId})
    key.EntityData.Leafs.Append("key-string", types.YLeaf{"KeyString", key.KeyString})
    key.EntityData.Leafs.Append("cryptographic-algorithm", types.YLeaf{"CryptographicAlgorithm", key.CryptographicAlgorithm})

    key.EntityData.YListKeys = []string {"KeyId"}

    return &(key.EntityData)
}

// Keychains_Keychain_Keys_Key_AcceptLifetime
// Configure a key Acceptance Lifetime
// This type is a presence type.
type Keychains_Keychain_Keys_Key_AcceptLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Start Hour. The type is interface{} with range: 0..23. This attribute is
    // mandatory.
    StartHour interface{}

    // Start Minutes. The type is interface{} with range: 0..59. This attribute is
    // mandatory. Units are minute.
    StartMinutes interface{}

    // Start Seconds. The type is interface{} with range: 0..59. This attribute is
    // mandatory. Units are second.
    StartSeconds interface{}

    // Start Date. The type is interface{} with range: 1..31. This attribute is
    // mandatory.
    StartDate interface{}

    // Start Month. The type is KeyChainMonth. This attribute is mandatory.
    StartMonth interface{}

    // Start Year. The type is interface{} with range: 1993..2035. This attribute
    // is mandatory.
    StartYear interface{}

    // Lifetime duration in seconds. The type is interface{} with range:
    // 1..2147483647. Units are second.
    LifeTime interface{}

    // Infinite Lifetime flag. The type is bool.
    InfiniteFlag interface{}

    // End Hour. The type is interface{} with range: 0..23.
    EndHour interface{}

    // End Minutes. The type is interface{} with range: 0..59. Units are minute.
    EndMinutes interface{}

    // End Seconds. The type is interface{} with range: 0..59. Units are second.
    EndSeconds interface{}

    // End Date. The type is interface{} with range: 1..31.
    EndDate interface{}

    // End Month. The type is KeyChainMonth.
    EndMonth interface{}

    // End Year. The type is interface{} with range: 1993..2035.
    EndYear interface{}
}

func (acceptLifetime *Keychains_Keychain_Keys_Key_AcceptLifetime) GetEntityData() *types.CommonEntityData {
    acceptLifetime.EntityData.YFilter = acceptLifetime.YFilter
    acceptLifetime.EntityData.YangName = "accept-lifetime"
    acceptLifetime.EntityData.BundleName = "cisco_ios_xr"
    acceptLifetime.EntityData.ParentYangName = "key"
    acceptLifetime.EntityData.SegmentPath = "accept-lifetime"
    acceptLifetime.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/keys/key/" + acceptLifetime.EntityData.SegmentPath
    acceptLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptLifetime.EntityData.Children = types.NewOrderedMap()
    acceptLifetime.EntityData.Leafs = types.NewOrderedMap()
    acceptLifetime.EntityData.Leafs.Append("start-hour", types.YLeaf{"StartHour", acceptLifetime.StartHour})
    acceptLifetime.EntityData.Leafs.Append("start-minutes", types.YLeaf{"StartMinutes", acceptLifetime.StartMinutes})
    acceptLifetime.EntityData.Leafs.Append("start-seconds", types.YLeaf{"StartSeconds", acceptLifetime.StartSeconds})
    acceptLifetime.EntityData.Leafs.Append("start-date", types.YLeaf{"StartDate", acceptLifetime.StartDate})
    acceptLifetime.EntityData.Leafs.Append("start-month", types.YLeaf{"StartMonth", acceptLifetime.StartMonth})
    acceptLifetime.EntityData.Leafs.Append("start-year", types.YLeaf{"StartYear", acceptLifetime.StartYear})
    acceptLifetime.EntityData.Leafs.Append("life-time", types.YLeaf{"LifeTime", acceptLifetime.LifeTime})
    acceptLifetime.EntityData.Leafs.Append("infinite-flag", types.YLeaf{"InfiniteFlag", acceptLifetime.InfiniteFlag})
    acceptLifetime.EntityData.Leafs.Append("end-hour", types.YLeaf{"EndHour", acceptLifetime.EndHour})
    acceptLifetime.EntityData.Leafs.Append("end-minutes", types.YLeaf{"EndMinutes", acceptLifetime.EndMinutes})
    acceptLifetime.EntityData.Leafs.Append("end-seconds", types.YLeaf{"EndSeconds", acceptLifetime.EndSeconds})
    acceptLifetime.EntityData.Leafs.Append("end-date", types.YLeaf{"EndDate", acceptLifetime.EndDate})
    acceptLifetime.EntityData.Leafs.Append("end-month", types.YLeaf{"EndMonth", acceptLifetime.EndMonth})
    acceptLifetime.EntityData.Leafs.Append("end-year", types.YLeaf{"EndYear", acceptLifetime.EndYear})

    acceptLifetime.EntityData.YListKeys = []string {}

    return &(acceptLifetime.EntityData)
}

// Keychains_Keychain_Keys_Key_SendLifetime
// Configure a Send Lifetime
// This type is a presence type.
type Keychains_Keychain_Keys_Key_SendLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Start Hour. The type is interface{} with range: 0..23. This attribute is
    // mandatory.
    StartHour interface{}

    // Start Minutes. The type is interface{} with range: 0..59. This attribute is
    // mandatory. Units are minute.
    StartMinutes interface{}

    // Start Seconds. The type is interface{} with range: 0..59. This attribute is
    // mandatory. Units are second.
    StartSeconds interface{}

    // Start Date. The type is interface{} with range: 1..31. This attribute is
    // mandatory.
    StartDate interface{}

    // Start Month. The type is KeyChainMonth. This attribute is mandatory.
    StartMonth interface{}

    // Start Year. The type is interface{} with range: 1993..2035. This attribute
    // is mandatory.
    StartYear interface{}

    // Lifetime duration in seconds. The type is interface{} with range:
    // 1..2147483647. Units are second.
    LifeTime interface{}

    // Infinite Lifetime flag. The type is bool.
    InfiniteFlag interface{}

    // End Hour. The type is interface{} with range: 0..23.
    EndHour interface{}

    // End Minutes. The type is interface{} with range: 0..59. Units are minute.
    EndMinutes interface{}

    // End Seconds. The type is interface{} with range: 0..59. Units are second.
    EndSeconds interface{}

    // End Date. The type is interface{} with range: 1..31.
    EndDate interface{}

    // End Month. The type is KeyChainMonth.
    EndMonth interface{}

    // End Year. The type is interface{} with range: 1993..2035.
    EndYear interface{}
}

func (sendLifetime *Keychains_Keychain_Keys_Key_SendLifetime) GetEntityData() *types.CommonEntityData {
    sendLifetime.EntityData.YFilter = sendLifetime.YFilter
    sendLifetime.EntityData.YangName = "send-lifetime"
    sendLifetime.EntityData.BundleName = "cisco_ios_xr"
    sendLifetime.EntityData.ParentYangName = "key"
    sendLifetime.EntityData.SegmentPath = "send-lifetime"
    sendLifetime.EntityData.AbsolutePath = "Cisco-IOS-XR-lib-keychain-cfg:keychains/keychain/keys/key/" + sendLifetime.EntityData.SegmentPath
    sendLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendLifetime.EntityData.Children = types.NewOrderedMap()
    sendLifetime.EntityData.Leafs = types.NewOrderedMap()
    sendLifetime.EntityData.Leafs.Append("start-hour", types.YLeaf{"StartHour", sendLifetime.StartHour})
    sendLifetime.EntityData.Leafs.Append("start-minutes", types.YLeaf{"StartMinutes", sendLifetime.StartMinutes})
    sendLifetime.EntityData.Leafs.Append("start-seconds", types.YLeaf{"StartSeconds", sendLifetime.StartSeconds})
    sendLifetime.EntityData.Leafs.Append("start-date", types.YLeaf{"StartDate", sendLifetime.StartDate})
    sendLifetime.EntityData.Leafs.Append("start-month", types.YLeaf{"StartMonth", sendLifetime.StartMonth})
    sendLifetime.EntityData.Leafs.Append("start-year", types.YLeaf{"StartYear", sendLifetime.StartYear})
    sendLifetime.EntityData.Leafs.Append("life-time", types.YLeaf{"LifeTime", sendLifetime.LifeTime})
    sendLifetime.EntityData.Leafs.Append("infinite-flag", types.YLeaf{"InfiniteFlag", sendLifetime.InfiniteFlag})
    sendLifetime.EntityData.Leafs.Append("end-hour", types.YLeaf{"EndHour", sendLifetime.EndHour})
    sendLifetime.EntityData.Leafs.Append("end-minutes", types.YLeaf{"EndMinutes", sendLifetime.EndMinutes})
    sendLifetime.EntityData.Leafs.Append("end-seconds", types.YLeaf{"EndSeconds", sendLifetime.EndSeconds})
    sendLifetime.EntityData.Leafs.Append("end-date", types.YLeaf{"EndDate", sendLifetime.EndDate})
    sendLifetime.EntityData.Leafs.Append("end-month", types.YLeaf{"EndMonth", sendLifetime.EndMonth})
    sendLifetime.EntityData.Leafs.Append("end-year", types.YLeaf{"EndYear", sendLifetime.EndYear})

    sendLifetime.EntityData.YListKeys = []string {}

    return &(sendLifetime.EntityData)
}

