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

// Keychains
// Configure a Key Chain
type Keychains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the key chain. The type is slice of Keychains_Keychain.
    Keychain []Keychains_Keychain
}

func (keychains *Keychains) GetFilter() yfilter.YFilter { return keychains.YFilter }

func (keychains *Keychains) SetFilter(yf yfilter.YFilter) { keychains.YFilter = yf }

func (keychains *Keychains) GetGoName(yname string) string {
    if yname == "keychain" { return "Keychain" }
    return ""
}

func (keychains *Keychains) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-keychain-cfg:keychains"
}

func (keychains *Keychains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "keychain" {
        for _, c := range keychains.Keychain {
            if keychains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Keychains_Keychain{}
        keychains.Keychain = append(keychains.Keychain, child)
        return &keychains.Keychain[len(keychains.Keychain)-1]
    }
    return nil
}

func (keychains *Keychains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range keychains.Keychain {
        children[keychains.Keychain[i].GetSegmentPath()] = &keychains.Keychain[i]
    }
    return children
}

func (keychains *Keychains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keychains *Keychains) GetBundleName() string { return "cisco_ios_xr" }

func (keychains *Keychains) GetYangName() string { return "keychains" }

func (keychains *Keychains) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keychains *Keychains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keychains *Keychains) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keychains *Keychains) SetParent(parent types.Entity) { keychains.parent = parent }

func (keychains *Keychains) GetParent() types.Entity { return keychains.parent }

func (keychains *Keychains) GetParentYangName() string { return "Cisco-IOS-XR-lib-keychain-cfg" }

// Keychains_Keychain
// Name of the key chain
type Keychains_Keychain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the key chain. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ChainName interface{}

    // Accept Tolerance in seconds or infinite.
    AcceptTolerance Keychains_Keychain_AcceptTolerance

    // Configure a Key.
    Keies Keychains_Keychain_Keies
}

func (keychain *Keychains_Keychain) GetFilter() yfilter.YFilter { return keychain.YFilter }

func (keychain *Keychains_Keychain) SetFilter(yf yfilter.YFilter) { keychain.YFilter = yf }

func (keychain *Keychains_Keychain) GetGoName(yname string) string {
    if yname == "chain-name" { return "ChainName" }
    if yname == "accept-tolerance" { return "AcceptTolerance" }
    if yname == "keies" { return "Keies" }
    return ""
}

func (keychain *Keychains_Keychain) GetSegmentPath() string {
    return "keychain" + "[chain-name='" + fmt.Sprintf("%v", keychain.ChainName) + "']"
}

func (keychain *Keychains_Keychain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accept-tolerance" {
        return &keychain.AcceptTolerance
    }
    if childYangName == "keies" {
        return &keychain.Keies
    }
    return nil
}

func (keychain *Keychains_Keychain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accept-tolerance"] = &keychain.AcceptTolerance
    children["keies"] = &keychain.Keies
    return children
}

func (keychain *Keychains_Keychain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chain-name"] = keychain.ChainName
    return leafs
}

func (keychain *Keychains_Keychain) GetBundleName() string { return "cisco_ios_xr" }

func (keychain *Keychains_Keychain) GetYangName() string { return "keychain" }

func (keychain *Keychains_Keychain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keychain *Keychains_Keychain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keychain *Keychains_Keychain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keychain *Keychains_Keychain) SetParent(parent types.Entity) { keychain.parent = parent }

func (keychain *Keychains_Keychain) GetParent() types.Entity { return keychain.parent }

func (keychain *Keychains_Keychain) GetParentYangName() string { return "keychains" }

// Keychains_Keychain_AcceptTolerance
// Accept Tolerance in seconds or infinite
type Keychains_Keychain_AcceptTolerance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value in seconds. The type is interface{} with range: 1..8640000. Units are
    // second.
    Value interface{}

    // Infinite tolerance. The type is bool.
    Infinite interface{}
}

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetFilter() yfilter.YFilter { return acceptTolerance.YFilter }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) SetFilter(yf yfilter.YFilter) { acceptTolerance.YFilter = yf }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "infinite" { return "Infinite" }
    return ""
}

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetSegmentPath() string {
    return "accept-tolerance"
}

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = acceptTolerance.Value
    leafs["infinite"] = acceptTolerance.Infinite
    return leafs
}

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetBundleName() string { return "cisco_ios_xr" }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetYangName() string { return "accept-tolerance" }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) SetParent(parent types.Entity) { acceptTolerance.parent = parent }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetParent() types.Entity { return acceptTolerance.parent }

func (acceptTolerance *Keychains_Keychain_AcceptTolerance) GetParentYangName() string { return "keychain" }

// Keychains_Keychain_Keies
// Configure a Key
type Keychains_Keychain_Keies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key Identifier. The type is slice of Keychains_Keychain_Keies_Key.
    Key []Keychains_Keychain_Keies_Key
}

func (keies *Keychains_Keychain_Keies) GetFilter() yfilter.YFilter { return keies.YFilter }

func (keies *Keychains_Keychain_Keies) SetFilter(yf yfilter.YFilter) { keies.YFilter = yf }

func (keies *Keychains_Keychain_Keies) GetGoName(yname string) string {
    if yname == "key" { return "Key" }
    return ""
}

func (keies *Keychains_Keychain_Keies) GetSegmentPath() string {
    return "keies"
}

func (keies *Keychains_Keychain_Keies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "key" {
        for _, c := range keies.Key {
            if keies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Keychains_Keychain_Keies_Key{}
        keies.Key = append(keies.Key, child)
        return &keies.Key[len(keies.Key)-1]
    }
    return nil
}

func (keies *Keychains_Keychain_Keies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range keies.Key {
        children[keies.Key[i].GetSegmentPath()] = &keies.Key[i]
    }
    return children
}

func (keies *Keychains_Keychain_Keies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keies *Keychains_Keychain_Keies) GetBundleName() string { return "cisco_ios_xr" }

func (keies *Keychains_Keychain_Keies) GetYangName() string { return "keies" }

func (keies *Keychains_Keychain_Keies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keies *Keychains_Keychain_Keies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keies *Keychains_Keychain_Keies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keies *Keychains_Keychain_Keies) SetParent(parent types.Entity) { keies.parent = parent }

func (keies *Keychains_Keychain_Keies) GetParent() types.Entity { return keies.parent }

func (keies *Keychains_Keychain_Keies) GetParentYangName() string { return "keychain" }

// Keychains_Keychain_Keies_Key
// Key Identifier
type Keychains_Keychain_Keies_Key struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 48-bit Key identifier. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    KeyId interface{}

    // Configure a clear text/encrypted Key string . The type is string with
    // pattern: (!.+)|([^!].+).
    KeyString interface{}

    // Configure the cryptographic algorithm. The type is CryptoAlg.
    CryptographicAlgorithm interface{}

    // Configure a key Acceptance Lifetime.
    AcceptLifetime Keychains_Keychain_Keies_Key_AcceptLifetime

    // Configure a Send Lifetime.
    SendLifetime Keychains_Keychain_Keies_Key_SendLifetime
}

func (key *Keychains_Keychain_Keies_Key) GetFilter() yfilter.YFilter { return key.YFilter }

func (key *Keychains_Keychain_Keies_Key) SetFilter(yf yfilter.YFilter) { key.YFilter = yf }

func (key *Keychains_Keychain_Keies_Key) GetGoName(yname string) string {
    if yname == "key-id" { return "KeyId" }
    if yname == "key-string" { return "KeyString" }
    if yname == "cryptographic-algorithm" { return "CryptographicAlgorithm" }
    if yname == "accept-lifetime" { return "AcceptLifetime" }
    if yname == "send-lifetime" { return "SendLifetime" }
    return ""
}

func (key *Keychains_Keychain_Keies_Key) GetSegmentPath() string {
    return "key" + "[key-id='" + fmt.Sprintf("%v", key.KeyId) + "']"
}

func (key *Keychains_Keychain_Keies_Key) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accept-lifetime" {
        return &key.AcceptLifetime
    }
    if childYangName == "send-lifetime" {
        return &key.SendLifetime
    }
    return nil
}

func (key *Keychains_Keychain_Keies_Key) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accept-lifetime"] = &key.AcceptLifetime
    children["send-lifetime"] = &key.SendLifetime
    return children
}

func (key *Keychains_Keychain_Keies_Key) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-id"] = key.KeyId
    leafs["key-string"] = key.KeyString
    leafs["cryptographic-algorithm"] = key.CryptographicAlgorithm
    return leafs
}

func (key *Keychains_Keychain_Keies_Key) GetBundleName() string { return "cisco_ios_xr" }

func (key *Keychains_Keychain_Keies_Key) GetYangName() string { return "key" }

func (key *Keychains_Keychain_Keies_Key) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (key *Keychains_Keychain_Keies_Key) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (key *Keychains_Keychain_Keies_Key) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (key *Keychains_Keychain_Keies_Key) SetParent(parent types.Entity) { key.parent = parent }

func (key *Keychains_Keychain_Keies_Key) GetParent() types.Entity { return key.parent }

func (key *Keychains_Keychain_Keies_Key) GetParentYangName() string { return "keies" }

// Keychains_Keychain_Keies_Key_AcceptLifetime
// Configure a key Acceptance Lifetime
type Keychains_Keychain_Keies_Key_AcceptLifetime struct {
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

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetFilter() yfilter.YFilter { return acceptLifetime.YFilter }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) SetFilter(yf yfilter.YFilter) { acceptLifetime.YFilter = yf }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetGoName(yname string) string {
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

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetSegmentPath() string {
    return "accept-lifetime"
}

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-hour"] = acceptLifetime.StartHour
    leafs["start-minutes"] = acceptLifetime.StartMinutes
    leafs["start-seconds"] = acceptLifetime.StartSeconds
    leafs["start-date"] = acceptLifetime.StartDate
    leafs["start-month"] = acceptLifetime.StartMonth
    leafs["start-year"] = acceptLifetime.StartYear
    leafs["life-time"] = acceptLifetime.LifeTime
    leafs["infinite-flag"] = acceptLifetime.InfiniteFlag
    leafs["end-hour"] = acceptLifetime.EndHour
    leafs["end-minutes"] = acceptLifetime.EndMinutes
    leafs["end-seconds"] = acceptLifetime.EndSeconds
    leafs["end-date"] = acceptLifetime.EndDate
    leafs["end-month"] = acceptLifetime.EndMonth
    leafs["end-year"] = acceptLifetime.EndYear
    return leafs
}

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetBundleName() string { return "cisco_ios_xr" }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetYangName() string { return "accept-lifetime" }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) SetParent(parent types.Entity) { acceptLifetime.parent = parent }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetParent() types.Entity { return acceptLifetime.parent }

func (acceptLifetime *Keychains_Keychain_Keies_Key_AcceptLifetime) GetParentYangName() string { return "key" }

// Keychains_Keychain_Keies_Key_SendLifetime
// Configure a Send Lifetime
type Keychains_Keychain_Keies_Key_SendLifetime struct {
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

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetFilter() yfilter.YFilter { return sendLifetime.YFilter }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) SetFilter(yf yfilter.YFilter) { sendLifetime.YFilter = yf }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetGoName(yname string) string {
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

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetSegmentPath() string {
    return "send-lifetime"
}

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-hour"] = sendLifetime.StartHour
    leafs["start-minutes"] = sendLifetime.StartMinutes
    leafs["start-seconds"] = sendLifetime.StartSeconds
    leafs["start-date"] = sendLifetime.StartDate
    leafs["start-month"] = sendLifetime.StartMonth
    leafs["start-year"] = sendLifetime.StartYear
    leafs["life-time"] = sendLifetime.LifeTime
    leafs["infinite-flag"] = sendLifetime.InfiniteFlag
    leafs["end-hour"] = sendLifetime.EndHour
    leafs["end-minutes"] = sendLifetime.EndMinutes
    leafs["end-seconds"] = sendLifetime.EndSeconds
    leafs["end-date"] = sendLifetime.EndDate
    leafs["end-month"] = sendLifetime.EndMonth
    leafs["end-year"] = sendLifetime.EndYear
    return leafs
}

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetBundleName() string { return "cisco_ios_xr" }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetYangName() string { return "send-lifetime" }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) SetParent(parent types.Entity) { sendLifetime.parent = parent }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetParent() types.Entity { return sendLifetime.parent }

func (sendLifetime *Keychains_Keychain_Keies_Key_SendLifetime) GetParentYangName() string { return "key" }

