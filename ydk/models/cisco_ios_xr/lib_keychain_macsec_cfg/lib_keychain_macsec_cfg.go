// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-keychain-macsec package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mac-sec-keychains: Configure a Key Chain
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lib_keychain_macsec_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lib_keychain_macsec_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-keychain-macsec-cfg mac-sec-keychains}", reflect.TypeOf(MacSecKeychains{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-keychain-macsec-cfg:mac-sec-keychains", reflect.TypeOf(MacSecKeychains{}))
}

// MacSecKeyChainMonth represents Mac sec key chain month
type MacSecKeyChainMonth string

const (
    // January
    MacSecKeyChainMonth_jan MacSecKeyChainMonth = "jan"

    // February
    MacSecKeyChainMonth_feb MacSecKeyChainMonth = "feb"

    // March
    MacSecKeyChainMonth_mar MacSecKeyChainMonth = "mar"

    // April
    MacSecKeyChainMonth_apr MacSecKeyChainMonth = "apr"

    // May
    MacSecKeyChainMonth_may MacSecKeyChainMonth = "may"

    // June
    MacSecKeyChainMonth_jun MacSecKeyChainMonth = "jun"

    // July
    MacSecKeyChainMonth_jul MacSecKeyChainMonth = "jul"

    // August
    MacSecKeyChainMonth_aug MacSecKeyChainMonth = "aug"

    // September
    MacSecKeyChainMonth_sep MacSecKeyChainMonth = "sep"

    // October
    MacSecKeyChainMonth_oct MacSecKeyChainMonth = "oct"

    // November
    MacSecKeyChainMonth_nov MacSecKeyChainMonth = "nov"

    // December
    MacSecKeyChainMonth_dec MacSecKeyChainMonth = "dec"
)

// MacSecCryptoAlg represents Mac sec crypto alg
type MacSecCryptoAlg string

const (
    // aes 128 cmac
    MacSecCryptoAlg_aes_128_cmac MacSecCryptoAlg = "aes-128-cmac"

    // aes 256 cmac
    MacSecCryptoAlg_aes_256_cmac MacSecCryptoAlg = "aes-256-cmac"
)

// MacSecEncryption represents Mac sec encryption
type MacSecEncryption string

const (
    // Type7
    MacSecEncryption_type7 MacSecEncryption = "type7"

    // Type6
    MacSecEncryption_type6 MacSecEncryption = "type6"
)

// MacSecKeychains
// Configure a Key Chain
type MacSecKeychains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the key chain for MACSec. The type is slice of
    // MacSecKeychains_MacSecKeychain.
    MacSecKeychain []MacSecKeychains_MacSecKeychain
}

func (macSecKeychains *MacSecKeychains) GetFilter() yfilter.YFilter { return macSecKeychains.YFilter }

func (macSecKeychains *MacSecKeychains) SetFilter(yf yfilter.YFilter) { macSecKeychains.YFilter = yf }

func (macSecKeychains *MacSecKeychains) GetGoName(yname string) string {
    if yname == "mac-sec-keychain" { return "MacSecKeychain" }
    return ""
}

func (macSecKeychains *MacSecKeychains) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-keychain-macsec-cfg:mac-sec-keychains"
}

func (macSecKeychains *MacSecKeychains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-sec-keychain" {
        for _, c := range macSecKeychains.MacSecKeychain {
            if macSecKeychains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacSecKeychains_MacSecKeychain{}
        macSecKeychains.MacSecKeychain = append(macSecKeychains.MacSecKeychain, child)
        return &macSecKeychains.MacSecKeychain[len(macSecKeychains.MacSecKeychain)-1]
    }
    return nil
}

func (macSecKeychains *MacSecKeychains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macSecKeychains.MacSecKeychain {
        children[macSecKeychains.MacSecKeychain[i].GetSegmentPath()] = &macSecKeychains.MacSecKeychain[i]
    }
    return children
}

func (macSecKeychains *MacSecKeychains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macSecKeychains *MacSecKeychains) GetBundleName() string { return "cisco_ios_xr" }

func (macSecKeychains *MacSecKeychains) GetYangName() string { return "mac-sec-keychains" }

func (macSecKeychains *MacSecKeychains) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macSecKeychains *MacSecKeychains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macSecKeychains *MacSecKeychains) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macSecKeychains *MacSecKeychains) SetParent(parent types.Entity) { macSecKeychains.parent = parent }

func (macSecKeychains *MacSecKeychains) GetParent() types.Entity { return macSecKeychains.parent }

func (macSecKeychains *MacSecKeychains) GetParentYangName() string { return "Cisco-IOS-XR-lib-keychain-macsec-cfg" }

// MacSecKeychains_MacSecKeychain
// Name of the key chain for MACSec
type MacSecKeychains_MacSecKeychain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the key chain. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ChainName interface{}

    // Configure a Key.
    Keies MacSecKeychains_MacSecKeychain_Keies
}

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetFilter() yfilter.YFilter { return macSecKeychain.YFilter }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) SetFilter(yf yfilter.YFilter) { macSecKeychain.YFilter = yf }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetGoName(yname string) string {
    if yname == "chain-name" { return "ChainName" }
    if yname == "keies" { return "Keies" }
    return ""
}

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetSegmentPath() string {
    return "mac-sec-keychain" + "[chain-name='" + fmt.Sprintf("%v", macSecKeychain.ChainName) + "']"
}

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "keies" {
        return &macSecKeychain.Keies
    }
    return nil
}

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["keies"] = &macSecKeychain.Keies
    return children
}

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chain-name"] = macSecKeychain.ChainName
    return leafs
}

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetBundleName() string { return "cisco_ios_xr" }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetYangName() string { return "mac-sec-keychain" }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) SetParent(parent types.Entity) { macSecKeychain.parent = parent }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetParent() types.Entity { return macSecKeychain.parent }

func (macSecKeychain *MacSecKeychains_MacSecKeychain) GetParentYangName() string { return "mac-sec-keychains" }

// MacSecKeychains_MacSecKeychain_Keies
// Configure a Key
type MacSecKeychains_MacSecKeychain_Keies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key Identifier. The type is slice of
    // MacSecKeychains_MacSecKeychain_Keies_Key.
    Key []MacSecKeychains_MacSecKeychain_Keies_Key
}

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetFilter() yfilter.YFilter { return keies.YFilter }

func (keies *MacSecKeychains_MacSecKeychain_Keies) SetFilter(yf yfilter.YFilter) { keies.YFilter = yf }

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetGoName(yname string) string {
    if yname == "key" { return "Key" }
    return ""
}

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetSegmentPath() string {
    return "keies"
}

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "key" {
        for _, c := range keies.Key {
            if keies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacSecKeychains_MacSecKeychain_Keies_Key{}
        keies.Key = append(keies.Key, child)
        return &keies.Key[len(keies.Key)-1]
    }
    return nil
}

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range keies.Key {
        children[keies.Key[i].GetSegmentPath()] = &keies.Key[i]
    }
    return children
}

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetBundleName() string { return "cisco_ios_xr" }

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetYangName() string { return "keies" }

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keies *MacSecKeychains_MacSecKeychain_Keies) SetParent(parent types.Entity) { keies.parent = parent }

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetParent() types.Entity { return keies.parent }

func (keies *MacSecKeychains_MacSecKeychain_Keies) GetParentYangName() string { return "mac-sec-keychain" }

// MacSecKeychains_MacSecKeychain_Keies_Key
// Key Identifier
type MacSecKeychains_MacSecKeychain_Keies_Key struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 48-bit Key identifier. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    KeyId interface{}

    // Configure a key Lifetime.
    Lifetime MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime

    // Configure a clear text/encrypted Key string along with cryptographic
    // algorithm.
    KeyString MacSecKeychains_MacSecKeychain_Keies_Key_KeyString
}

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetFilter() yfilter.YFilter { return key.YFilter }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) SetFilter(yf yfilter.YFilter) { key.YFilter = yf }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetGoName(yname string) string {
    if yname == "key-id" { return "KeyId" }
    if yname == "lifetime" { return "Lifetime" }
    if yname == "key-string" { return "KeyString" }
    return ""
}

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetSegmentPath() string {
    return "key" + "[key-id='" + fmt.Sprintf("%v", key.KeyId) + "']"
}

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lifetime" {
        return &key.Lifetime
    }
    if childYangName == "key-string" {
        return &key.KeyString
    }
    return nil
}

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lifetime"] = &key.Lifetime
    children["key-string"] = &key.KeyString
    return children
}

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-id"] = key.KeyId
    return leafs
}

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetBundleName() string { return "cisco_ios_xr" }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetYangName() string { return "key" }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) SetParent(parent types.Entity) { key.parent = parent }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetParent() types.Entity { return key.parent }

func (key *MacSecKeychains_MacSecKeychain_Keies_Key) GetParentYangName() string { return "keies" }

// MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime
// Configure a key Lifetime
type MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start Hour. The type is interface{} with range: 0..23.
    StartHour interface{}

    // Start Minutes. The type is interface{} with range: 0..59. Units are minute.
    StartMinutes interface{}

    // Start Seconds. The type is interface{} with range: 0..59. Units are second.
    StartSeconds interface{}

    // Start Date. The type is interface{} with range: 1..31.
    StartDate interface{}

    // Start Month. The type is MacSecKeyChainMonth.
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

    // End Month. The type is MacSecKeyChainMonth.
    EndMonth interface{}

    // End Year. The type is interface{} with range: 1993..2035.
    EndYear interface{}
}

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetFilter() yfilter.YFilter { return lifetime.YFilter }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) SetFilter(yf yfilter.YFilter) { lifetime.YFilter = yf }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetGoName(yname string) string {
    if yname == "start-hour" { return "StartHour" }
    if yname == "start-minutes" { return "StartMinutes" }
    if yname == "start-seconds" { return "StartSeconds" }
    if yname == "start-date" { return "StartDate" }
    if yname == "start-month" { return "StartMonth" }
    if yname == "start-year" { return "StartYear" }
    if yname == "life-time" { return "LifeTime" }
    if yname == "infinite-flag" { return "InfiniteFlag" }
    if yname == "end-hour" { return "EndHour" }
    if yname == "end-minutes" { return "EndMinutes" }
    if yname == "end-seconds" { return "EndSeconds" }
    if yname == "end-date" { return "EndDate" }
    if yname == "end-month" { return "EndMonth" }
    if yname == "end-year" { return "EndYear" }
    return ""
}

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetSegmentPath() string {
    return "lifetime"
}

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-hour"] = lifetime.StartHour
    leafs["start-minutes"] = lifetime.StartMinutes
    leafs["start-seconds"] = lifetime.StartSeconds
    leafs["start-date"] = lifetime.StartDate
    leafs["start-month"] = lifetime.StartMonth
    leafs["start-year"] = lifetime.StartYear
    leafs["life-time"] = lifetime.LifeTime
    leafs["infinite-flag"] = lifetime.InfiniteFlag
    leafs["end-hour"] = lifetime.EndHour
    leafs["end-minutes"] = lifetime.EndMinutes
    leafs["end-seconds"] = lifetime.EndSeconds
    leafs["end-date"] = lifetime.EndDate
    leafs["end-month"] = lifetime.EndMonth
    leafs["end-year"] = lifetime.EndYear
    return leafs
}

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetBundleName() string { return "cisco_ios_xr" }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetYangName() string { return "lifetime" }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) SetParent(parent types.Entity) { lifetime.parent = parent }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetParent() types.Entity { return lifetime.parent }

func (lifetime *MacSecKeychains_MacSecKeychain_Keies_Key_Lifetime) GetParentYangName() string { return "key" }

// MacSecKeychains_MacSecKeychain_Keies_Key_KeyString
// Configure a clear text/encrypted Key string
// along with cryptographic algorithm
// This type is a presence type.
type MacSecKeychains_MacSecKeychain_Keies_Key_KeyString struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key String. The type is string with pattern: (!.+)|([^!].+). This attribute
    // is mandatory.
    String interface{}

    // Cryptographic Algorithm. The type is MacSecCryptoAlg. This attribute is
    // mandatory.
    CryptographicAlgorithm interface{}

    // encryption type used to store key. The type is MacSecEncryption. The
    // default value is type7.
    EncryptionType interface{}
}

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetFilter() yfilter.YFilter { return keyString.YFilter }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) SetFilter(yf yfilter.YFilter) { keyString.YFilter = yf }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetGoName(yname string) string {
    if yname == "string" { return "String" }
    if yname == "cryptographic-algorithm" { return "CryptographicAlgorithm" }
    if yname == "encryption-type" { return "EncryptionType" }
    return ""
}

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetSegmentPath() string {
    return "key-string"
}

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["string"] = keyString.String
    leafs["cryptographic-algorithm"] = keyString.CryptographicAlgorithm
    leafs["encryption-type"] = keyString.EncryptionType
    return leafs
}

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetBundleName() string { return "cisco_ios_xr" }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetYangName() string { return "key-string" }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) SetParent(parent types.Entity) { keyString.parent = parent }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetParent() types.Entity { return keyString.parent }

func (keyString *MacSecKeychains_MacSecKeychain_Keies_Key_KeyString) GetParentYangName() string { return "key" }

