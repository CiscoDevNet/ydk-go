// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-keychain package configuration.
// 
// This module contains definitions
// for the following management objects:
//   keychains: Configure a Key Chain
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    Keychain []Keychains_Keychain
}

func (keychains *Keychains) GetEntityData() *types.CommonEntityData {
    keychains.EntityData.YFilter = keychains.YFilter
    keychains.EntityData.YangName = "keychains"
    keychains.EntityData.BundleName = "cisco_ios_xr"
    keychains.EntityData.ParentYangName = "Cisco-IOS-XR-lib-keychain-cfg"
    keychains.EntityData.SegmentPath = "Cisco-IOS-XR-lib-keychain-cfg:keychains"
    keychains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychains.EntityData.Children = make(map[string]types.YChild)
    keychains.EntityData.Children["keychain"] = types.YChild{"Keychain", nil}
    for i := range keychains.Keychain {
        keychains.EntityData.Children[types.GetSegmentPath(&keychains.Keychain[i])] = types.YChild{"Keychain", &keychains.Keychain[i]}
    }
    keychains.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(keychains.EntityData)
}

// Keychains_Keychain
// Name of the key chain
type Keychains_Keychain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the key chain. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ChainName interface{}

    // Accept Tolerance in seconds or infinite.
    AcceptTolerance Keychains_Keychain_AcceptTolerance

    // Name of the key chain for MACSec.
    MacsecKeychain Keychains_Keychain_MacsecKeychain

    // Configure a Key.
    Keies Keychains_Keychain_Keies
}

func (keychain *Keychains_Keychain) GetEntityData() *types.CommonEntityData {
    keychain.EntityData.YFilter = keychain.YFilter
    keychain.EntityData.YangName = "keychain"
    keychain.EntityData.BundleName = "cisco_ios_xr"
    keychain.EntityData.ParentYangName = "keychains"
    keychain.EntityData.SegmentPath = "keychain" + "[chain-name='" + fmt.Sprintf("%v", keychain.ChainName) + "']"
    keychain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychain.EntityData.Children = make(map[string]types.YChild)
    keychain.EntityData.Children["accept-tolerance"] = types.YChild{"AcceptTolerance", &keychain.AcceptTolerance}
    keychain.EntityData.Children["macsec-keychain"] = types.YChild{"MacsecKeychain", &keychain.MacsecKeychain}
    keychain.EntityData.Children["keies"] = types.YChild{"Keies", &keychain.Keies}
    keychain.EntityData.Leafs = make(map[string]types.YLeaf)
    keychain.EntityData.Leafs["chain-name"] = types.YLeaf{"ChainName", keychain.ChainName}
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
    acceptTolerance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptTolerance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptTolerance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptTolerance.EntityData.Children = make(map[string]types.YChild)
    acceptTolerance.EntityData.Leafs = make(map[string]types.YLeaf)
    acceptTolerance.EntityData.Leafs["value"] = types.YLeaf{"Value", acceptTolerance.Value}
    acceptTolerance.EntityData.Leafs["infinite"] = types.YLeaf{"Infinite", acceptTolerance.Infinite}
    return &(acceptTolerance.EntityData)
}

// Keychains_Keychain_MacsecKeychain
// Name of the key chain for MACSec
// This type is a presence type.
type Keychains_Keychain_MacsecKeychain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a Key.
    MacsecKeies Keychains_Keychain_MacsecKeychain_MacsecKeies
}

func (macsecKeychain *Keychains_Keychain_MacsecKeychain) GetEntityData() *types.CommonEntityData {
    macsecKeychain.EntityData.YFilter = macsecKeychain.YFilter
    macsecKeychain.EntityData.YangName = "macsec-keychain"
    macsecKeychain.EntityData.BundleName = "cisco_ios_xr"
    macsecKeychain.EntityData.ParentYangName = "keychain"
    macsecKeychain.EntityData.SegmentPath = "macsec-keychain"
    macsecKeychain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKeychain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKeychain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKeychain.EntityData.Children = make(map[string]types.YChild)
    macsecKeychain.EntityData.Children["macsec-keies"] = types.YChild{"MacsecKeies", &macsecKeychain.MacsecKeies}
    macsecKeychain.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macsecKeychain.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeies
// Configure a Key
type Keychains_Keychain_MacsecKeychain_MacsecKeies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key Identifier. The type is slice of
    // Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey.
    MacsecKey []Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey
}

func (macsecKeies *Keychains_Keychain_MacsecKeychain_MacsecKeies) GetEntityData() *types.CommonEntityData {
    macsecKeies.EntityData.YFilter = macsecKeies.YFilter
    macsecKeies.EntityData.YangName = "macsec-keies"
    macsecKeies.EntityData.BundleName = "cisco_ios_xr"
    macsecKeies.EntityData.ParentYangName = "macsec-keychain"
    macsecKeies.EntityData.SegmentPath = "macsec-keies"
    macsecKeies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKeies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKeies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKeies.EntityData.Children = make(map[string]types.YChild)
    macsecKeies.EntityData.Children["macsec-key"] = types.YChild{"MacsecKey", nil}
    for i := range macsecKeies.MacsecKey {
        macsecKeies.EntityData.Children[types.GetSegmentPath(&macsecKeies.MacsecKey[i])] = types.YChild{"MacsecKey", &macsecKeies.MacsecKey[i]}
    }
    macsecKeies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macsecKeies.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey
// Key Identifier
type Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. 48-bit Key identifier. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    KeyId interface{}

    // Configure a key Lifetime.
    MacsecLifetime Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecLifetime

    // Configure a clear text/encrypted Key string along with cryptographic
    // algorithm.
    MacsecKeyString Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecKeyString
}

func (macsecKey *Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey) GetEntityData() *types.CommonEntityData {
    macsecKey.EntityData.YFilter = macsecKey.YFilter
    macsecKey.EntityData.YangName = "macsec-key"
    macsecKey.EntityData.BundleName = "cisco_ios_xr"
    macsecKey.EntityData.ParentYangName = "macsec-keies"
    macsecKey.EntityData.SegmentPath = "macsec-key" + "[key-id='" + fmt.Sprintf("%v", macsecKey.KeyId) + "']"
    macsecKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKey.EntityData.Children = make(map[string]types.YChild)
    macsecKey.EntityData.Children["macsec-lifetime"] = types.YChild{"MacsecLifetime", &macsecKey.MacsecLifetime}
    macsecKey.EntityData.Children["macsec-key-string"] = types.YChild{"MacsecKeyString", &macsecKey.MacsecKeyString}
    macsecKey.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecKey.EntityData.Leafs["key-id"] = types.YLeaf{"KeyId", macsecKey.KeyId}
    return &(macsecKey.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecLifetime
// Configure a key Lifetime
type Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start Hour. The type is interface{} with range: 0..23.
    StartHour interface{}

    // Start Minutes. The type is interface{} with range: 0..59. Units are minute.
    StartMinutes interface{}

    // Start Seconds. The type is interface{} with range: 0..59. Units are second.
    StartSeconds interface{}

    // Start Date. The type is interface{} with range: 1..31.
    StartDate interface{}

    // Start Month. The type is KeyChainMonth.
    StartMonth interface{}

    // Start Year. The type is interface{} with range: 1993..2035.
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

func (macsecLifetime *Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecLifetime) GetEntityData() *types.CommonEntityData {
    macsecLifetime.EntityData.YFilter = macsecLifetime.YFilter
    macsecLifetime.EntityData.YangName = "macsec-lifetime"
    macsecLifetime.EntityData.BundleName = "cisco_ios_xr"
    macsecLifetime.EntityData.ParentYangName = "macsec-key"
    macsecLifetime.EntityData.SegmentPath = "macsec-lifetime"
    macsecLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecLifetime.EntityData.Children = make(map[string]types.YChild)
    macsecLifetime.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecLifetime.EntityData.Leafs["start-hour"] = types.YLeaf{"StartHour", macsecLifetime.StartHour}
    macsecLifetime.EntityData.Leafs["start-minutes"] = types.YLeaf{"StartMinutes", macsecLifetime.StartMinutes}
    macsecLifetime.EntityData.Leafs["start-seconds"] = types.YLeaf{"StartSeconds", macsecLifetime.StartSeconds}
    macsecLifetime.EntityData.Leafs["start-date"] = types.YLeaf{"StartDate", macsecLifetime.StartDate}
    macsecLifetime.EntityData.Leafs["start-month"] = types.YLeaf{"StartMonth", macsecLifetime.StartMonth}
    macsecLifetime.EntityData.Leafs["start-year"] = types.YLeaf{"StartYear", macsecLifetime.StartYear}
    macsecLifetime.EntityData.Leafs["life-time"] = types.YLeaf{"LifeTime", macsecLifetime.LifeTime}
    macsecLifetime.EntityData.Leafs["infinite-flag"] = types.YLeaf{"InfiniteFlag", macsecLifetime.InfiniteFlag}
    macsecLifetime.EntityData.Leafs["end-hour"] = types.YLeaf{"EndHour", macsecLifetime.EndHour}
    macsecLifetime.EntityData.Leafs["end-minutes"] = types.YLeaf{"EndMinutes", macsecLifetime.EndMinutes}
    macsecLifetime.EntityData.Leafs["end-seconds"] = types.YLeaf{"EndSeconds", macsecLifetime.EndSeconds}
    macsecLifetime.EntityData.Leafs["end-date"] = types.YLeaf{"EndDate", macsecLifetime.EndDate}
    macsecLifetime.EntityData.Leafs["end-month"] = types.YLeaf{"EndMonth", macsecLifetime.EndMonth}
    macsecLifetime.EntityData.Leafs["end-year"] = types.YLeaf{"EndYear", macsecLifetime.EndYear}
    return &(macsecLifetime.EntityData)
}

// Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecKeyString
// Configure a clear text/encrypted Key string
// along with cryptographic algorithm
// This type is a presence type.
type Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecKeyString struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key String. The type is string with pattern: b'(!.+)|([^!].+)'. This
    // attribute is mandatory.
    String_ interface{}

    // Cryptographic Algorithm. The type is MacsecCryptoAlg. This attribute is
    // mandatory.
    CryptographicAlgorithm interface{}

    // encryption type used to store key. The type is MacsecEncryption. The
    // default value is type7.
    EncryptionType interface{}
}

func (macsecKeyString *Keychains_Keychain_MacsecKeychain_MacsecKeies_MacsecKey_MacsecKeyString) GetEntityData() *types.CommonEntityData {
    macsecKeyString.EntityData.YFilter = macsecKeyString.YFilter
    macsecKeyString.EntityData.YangName = "macsec-key-string"
    macsecKeyString.EntityData.BundleName = "cisco_ios_xr"
    macsecKeyString.EntityData.ParentYangName = "macsec-key"
    macsecKeyString.EntityData.SegmentPath = "macsec-key-string"
    macsecKeyString.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecKeyString.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecKeyString.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecKeyString.EntityData.Children = make(map[string]types.YChild)
    macsecKeyString.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecKeyString.EntityData.Leafs["string"] = types.YLeaf{"String_", macsecKeyString.String_}
    macsecKeyString.EntityData.Leafs["cryptographic-algorithm"] = types.YLeaf{"CryptographicAlgorithm", macsecKeyString.CryptographicAlgorithm}
    macsecKeyString.EntityData.Leafs["encryption-type"] = types.YLeaf{"EncryptionType", macsecKeyString.EncryptionType}
    return &(macsecKeyString.EntityData)
}

// Keychains_Keychain_Keies
// Configure a Key
type Keychains_Keychain_Keies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key Identifier. The type is slice of Keychains_Keychain_Keies_Key.
    Key []Keychains_Keychain_Keies_Key
}

func (keies *Keychains_Keychain_Keies) GetEntityData() *types.CommonEntityData {
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

// Keychains_Keychain_Keies_Key
// Key Identifier
type Keychains_Keychain_Keies_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. 48-bit Key identifier. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    KeyId interface{}

    // Configure a clear text/encrypted Key string . The type is string with
    // pattern: b'(!.+)|([^!].+)'.
    KeyString interface{}

    // Configure the cryptographic algorithm. The type is CryptoAlg.
    CryptographicAlgorithm interface{}

    // Configure a key Acceptance Lifetime.
    AcceptLifetime Keychains_Keychain_Keies_Key_AcceptLifetime

    // Configure a Send Lifetime.
    SendLifetime Keychains_Keychain_Keies_Key_SendLifetime
}

func (key *Keychains_Keychain_Keies_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "keies"
    key.EntityData.SegmentPath = "key" + "[key-id='" + fmt.Sprintf("%v", key.KeyId) + "']"
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = make(map[string]types.YChild)
    key.EntityData.Children["accept-lifetime"] = types.YChild{"AcceptLifetime", &key.AcceptLifetime}
    key.EntityData.Children["send-lifetime"] = types.YChild{"SendLifetime", &key.SendLifetime}
    key.EntityData.Leafs = make(map[string]types.YLeaf)
    key.EntityData.Leafs["key-id"] = types.YLeaf{"KeyId", key.KeyId}
    key.EntityData.Leafs["key-string"] = types.YLeaf{"KeyString", key.KeyString}
    key.EntityData.Leafs["cryptographic-algorithm"] = types.YLeaf{"CryptographicAlgorithm", key.CryptographicAlgorithm}
    return &(key.EntityData)
}

// Keychains_Keychain_Keies_Key_AcceptLifetime
// Configure a key Acceptance Lifetime
type Keychains_Keychain_Keies_Key_AcceptLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start Hour. The type is interface{} with range: 0..23.
    StartHour interface{}

    // Start Minutes. The type is interface{} with range: 0..59. Units are minute.
    StartMinutes interface{}

    // Start Seconds. The type is interface{} with range: 0..59. Units are second.
    StartSeconds interface{}

    // Start Date. The type is interface{} with range: 1..31.
    StartDate interface{}

    // Start Month. The type is KeyChainMonth.
    StartMonth interface{}

    // Start Year. The type is interface{} with range: 1993..2035.
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

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetEntityData() *types.CommonEntityData {
    acceptLifetime.EntityData.YFilter = acceptLifetime.YFilter
    acceptLifetime.EntityData.YangName = "accept-lifetime"
    acceptLifetime.EntityData.BundleName = "cisco_ios_xr"
    acceptLifetime.EntityData.ParentYangName = "key"
    acceptLifetime.EntityData.SegmentPath = "accept-lifetime"
    acceptLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptLifetime.EntityData.Children = make(map[string]types.YChild)
    acceptLifetime.EntityData.Leafs = make(map[string]types.YLeaf)
    acceptLifetime.EntityData.Leafs["start-hour"] = types.YLeaf{"StartHour", acceptLifetime.StartHour}
    acceptLifetime.EntityData.Leafs["start-minutes"] = types.YLeaf{"StartMinutes", acceptLifetime.StartMinutes}
    acceptLifetime.EntityData.Leafs["start-seconds"] = types.YLeaf{"StartSeconds", acceptLifetime.StartSeconds}
    acceptLifetime.EntityData.Leafs["start-date"] = types.YLeaf{"StartDate", acceptLifetime.StartDate}
    acceptLifetime.EntityData.Leafs["start-month"] = types.YLeaf{"StartMonth", acceptLifetime.StartMonth}
    acceptLifetime.EntityData.Leafs["start-year"] = types.YLeaf{"StartYear", acceptLifetime.StartYear}
    acceptLifetime.EntityData.Leafs["life-time"] = types.YLeaf{"LifeTime", acceptLifetime.LifeTime}
    acceptLifetime.EntityData.Leafs["infinite-flag"] = types.YLeaf{"InfiniteFlag", acceptLifetime.InfiniteFlag}
    acceptLifetime.EntityData.Leafs["end-hour"] = types.YLeaf{"EndHour", acceptLifetime.EndHour}
    acceptLifetime.EntityData.Leafs["end-minutes"] = types.YLeaf{"EndMinutes", acceptLifetime.EndMinutes}
    acceptLifetime.EntityData.Leafs["end-seconds"] = types.YLeaf{"EndSeconds", acceptLifetime.EndSeconds}
    acceptLifetime.EntityData.Leafs["end-date"] = types.YLeaf{"EndDate", acceptLifetime.EndDate}
    acceptLifetime.EntityData.Leafs["end-month"] = types.YLeaf{"EndMonth", acceptLifetime.EndMonth}
    acceptLifetime.EntityData.Leafs["end-year"] = types.YLeaf{"EndYear", acceptLifetime.EndYear}
    return &(acceptLifetime.EntityData)
}

// Keychains_Keychain_Keies_Key_SendLifetime
// Configure a Send Lifetime
type Keychains_Keychain_Keies_Key_SendLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start Hour. The type is interface{} with range: 0..23.
    StartHour interface{}

    // Start Minutes. The type is interface{} with range: 0..59. Units are minute.
    StartMinutes interface{}

    // Start Seconds. The type is interface{} with range: 0..59. Units are second.
    StartSeconds interface{}

    // Start Date. The type is interface{} with range: 1..31.
    StartDate interface{}

    // Start Month. The type is KeyChainMonth.
    StartMonth interface{}

    // Start Year. The type is interface{} with range: 1993..2035.
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

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetEntityData() *types.CommonEntityData {
    sendLifetime.EntityData.YFilter = sendLifetime.YFilter
    sendLifetime.EntityData.YangName = "send-lifetime"
    sendLifetime.EntityData.BundleName = "cisco_ios_xr"
    sendLifetime.EntityData.ParentYangName = "key"
    sendLifetime.EntityData.SegmentPath = "send-lifetime"
    sendLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendLifetime.EntityData.Children = make(map[string]types.YChild)
    sendLifetime.EntityData.Leafs = make(map[string]types.YLeaf)
    sendLifetime.EntityData.Leafs["start-hour"] = types.YLeaf{"StartHour", sendLifetime.StartHour}
    sendLifetime.EntityData.Leafs["start-minutes"] = types.YLeaf{"StartMinutes", sendLifetime.StartMinutes}
    sendLifetime.EntityData.Leafs["start-seconds"] = types.YLeaf{"StartSeconds", sendLifetime.StartSeconds}
    sendLifetime.EntityData.Leafs["start-date"] = types.YLeaf{"StartDate", sendLifetime.StartDate}
    sendLifetime.EntityData.Leafs["start-month"] = types.YLeaf{"StartMonth", sendLifetime.StartMonth}
    sendLifetime.EntityData.Leafs["start-year"] = types.YLeaf{"StartYear", sendLifetime.StartYear}
    sendLifetime.EntityData.Leafs["life-time"] = types.YLeaf{"LifeTime", sendLifetime.LifeTime}
    sendLifetime.EntityData.Leafs["infinite-flag"] = types.YLeaf{"InfiniteFlag", sendLifetime.InfiniteFlag}
    sendLifetime.EntityData.Leafs["end-hour"] = types.YLeaf{"EndHour", sendLifetime.EndHour}
    sendLifetime.EntityData.Leafs["end-minutes"] = types.YLeaf{"EndMinutes", sendLifetime.EndMinutes}
    sendLifetime.EntityData.Leafs["end-seconds"] = types.YLeaf{"EndSeconds", sendLifetime.EndSeconds}
    sendLifetime.EntityData.Leafs["end-date"] = types.YLeaf{"EndDate", sendLifetime.EndDate}
    sendLifetime.EntityData.Leafs["end-month"] = types.YLeaf{"EndMonth", sendLifetime.EndMonth}
    sendLifetime.EntityData.Leafs["end-year"] = types.YLeaf{"EndYear", sendLifetime.EndYear}
    return &(sendLifetime.EntityData)
}

