// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package crypto_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-generate-rsa-general-keys}", reflect.TypeOf(KeyGenerateRsaGeneralKeys{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-generate-rsa-general-keys", reflect.TypeOf(KeyGenerateRsaGeneralKeys{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-generate-rsa-usage-keys}", reflect.TypeOf(KeyGenerateRsaUsageKeys{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-generate-rsa-usage-keys", reflect.TypeOf(KeyGenerateRsaUsageKeys{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-generate-rsa}", reflect.TypeOf(KeyGenerateRsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-generate-rsa", reflect.TypeOf(KeyGenerateRsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-generate-dsa}", reflect.TypeOf(KeyGenerateDsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-generate-dsa", reflect.TypeOf(KeyGenerateDsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-zeroize-rsa}", reflect.TypeOf(KeyZeroizeRsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-zeroize-rsa", reflect.TypeOf(KeyZeroizeRsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-zeroize-dsa}", reflect.TypeOf(KeyZeroizeDsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-zeroize-dsa", reflect.TypeOf(KeyZeroizeDsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-zeroize-authentication-rsa}", reflect.TypeOf(KeyZeroizeAuthenticationRsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-zeroize-authentication-rsa", reflect.TypeOf(KeyZeroizeAuthenticationRsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-import-authentication-rsa}", reflect.TypeOf(KeyImportAuthenticationRsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-import-authentication-rsa", reflect.TypeOf(KeyImportAuthenticationRsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act ca-authenticate}", reflect.TypeOf(CaAuthenticate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:ca-authenticate", reflect.TypeOf(CaAuthenticate{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act ca-enroll}", reflect.TypeOf(CaEnroll{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:ca-enroll", reflect.TypeOf(CaEnroll{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act ca-import-certificate}", reflect.TypeOf(CaImportCertificate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:ca-import-certificate", reflect.TypeOf(CaImportCertificate{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act ca-cancel-enroll}", reflect.TypeOf(CaCancelEnroll{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:ca-cancel-enroll", reflect.TypeOf(CaCancelEnroll{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act ca-crl-request}", reflect.TypeOf(CaCrlRequest{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:ca-crl-request", reflect.TypeOf(CaCrlRequest{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act ca-trustpool-import-url}", reflect.TypeOf(CaTrustpoolImportUrl{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:ca-trustpool-import-url", reflect.TypeOf(CaTrustpoolImportUrl{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act ca-trustpool-import-url-clean}", reflect.TypeOf(CaTrustpoolImportUrlClean{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:ca-trustpool-import-url-clean", reflect.TypeOf(CaTrustpoolImportUrlClean{}))
}

// KeyGenerateRsaGeneralKeys
// Generate a general purpose RSA key pair for signing and encryption
type KeyGenerateRsaGeneralKeys struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input KeyGenerateRsaGeneralKeys_Input
}

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetFilter() yfilter.YFilter { return keyGenerateRsaGeneralKeys.YFilter }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) SetFilter(yf yfilter.YFilter) { keyGenerateRsaGeneralKeys.YFilter = yf }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-generate-rsa-general-keys"
}

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &keyGenerateRsaGeneralKeys.Input
    }
    return nil
}

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &keyGenerateRsaGeneralKeys.Input
    return children
}

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetBundleName() string { return "cisco_ios_xr" }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetYangName() string { return "key-generate-rsa-general-keys" }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) SetParent(parent types.Entity) { keyGenerateRsaGeneralKeys.parent = parent }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetParent() types.Entity { return keyGenerateRsaGeneralKeys.parent }

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyGenerateRsaGeneralKeys_Input
type KeyGenerateRsaGeneralKeys_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSA keypair label. The type is string.
    KeyLabel interface{}

    // Key modulus in the range of 512 to 4096 for your General Purpose Keypair.
    // Choosing a key modulus greater than 512 may take a few minutes. The type is
    // interface{} with range: 512..4096. This attribute is mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateRsaGeneralKeys_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *KeyGenerateRsaGeneralKeys_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *KeyGenerateRsaGeneralKeys_Input) GetGoName(yname string) string {
    if yname == "key-label" { return "KeyLabel" }
    if yname == "key-modulus" { return "KeyModulus" }
    return ""
}

func (input *KeyGenerateRsaGeneralKeys_Input) GetSegmentPath() string {
    return "input"
}

func (input *KeyGenerateRsaGeneralKeys_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *KeyGenerateRsaGeneralKeys_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *KeyGenerateRsaGeneralKeys_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-label"] = input.KeyLabel
    leafs["key-modulus"] = input.KeyModulus
    return leafs
}

func (input *KeyGenerateRsaGeneralKeys_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *KeyGenerateRsaGeneralKeys_Input) GetYangName() string { return "input" }

func (input *KeyGenerateRsaGeneralKeys_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *KeyGenerateRsaGeneralKeys_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *KeyGenerateRsaGeneralKeys_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *KeyGenerateRsaGeneralKeys_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *KeyGenerateRsaGeneralKeys_Input) GetParent() types.Entity { return input.parent }

func (input *KeyGenerateRsaGeneralKeys_Input) GetParentYangName() string { return "key-generate-rsa-general-keys" }

// KeyGenerateRsaUsageKeys
// Generate seperate RSA key pairs for signing and encryption
type KeyGenerateRsaUsageKeys struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input KeyGenerateRsaUsageKeys_Input
}

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetFilter() yfilter.YFilter { return keyGenerateRsaUsageKeys.YFilter }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) SetFilter(yf yfilter.YFilter) { keyGenerateRsaUsageKeys.YFilter = yf }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-generate-rsa-usage-keys"
}

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &keyGenerateRsaUsageKeys.Input
    }
    return nil
}

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &keyGenerateRsaUsageKeys.Input
    return children
}

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetBundleName() string { return "cisco_ios_xr" }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetYangName() string { return "key-generate-rsa-usage-keys" }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) SetParent(parent types.Entity) { keyGenerateRsaUsageKeys.parent = parent }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetParent() types.Entity { return keyGenerateRsaUsageKeys.parent }

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyGenerateRsaUsageKeys_Input
type KeyGenerateRsaUsageKeys_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSA keypair label. The type is string.
    KeyLabel interface{}

    // Key modulus in the range of 512 to 4096 for your General Purpose Keypair.
    // Choosing a key modulus greater than 512 may take a few minutes. The type is
    // interface{} with range: 512..4096. This attribute is mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateRsaUsageKeys_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *KeyGenerateRsaUsageKeys_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *KeyGenerateRsaUsageKeys_Input) GetGoName(yname string) string {
    if yname == "key-label" { return "KeyLabel" }
    if yname == "key-modulus" { return "KeyModulus" }
    return ""
}

func (input *KeyGenerateRsaUsageKeys_Input) GetSegmentPath() string {
    return "input"
}

func (input *KeyGenerateRsaUsageKeys_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *KeyGenerateRsaUsageKeys_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *KeyGenerateRsaUsageKeys_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-label"] = input.KeyLabel
    leafs["key-modulus"] = input.KeyModulus
    return leafs
}

func (input *KeyGenerateRsaUsageKeys_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *KeyGenerateRsaUsageKeys_Input) GetYangName() string { return "input" }

func (input *KeyGenerateRsaUsageKeys_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *KeyGenerateRsaUsageKeys_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *KeyGenerateRsaUsageKeys_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *KeyGenerateRsaUsageKeys_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *KeyGenerateRsaUsageKeys_Input) GetParent() types.Entity { return input.parent }

func (input *KeyGenerateRsaUsageKeys_Input) GetParentYangName() string { return "key-generate-rsa-usage-keys" }

// KeyGenerateRsa
// Generate seperate RSA key pairs for signing and encryption
type KeyGenerateRsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input KeyGenerateRsa_Input
}

func (keyGenerateRsa *KeyGenerateRsa) GetFilter() yfilter.YFilter { return keyGenerateRsa.YFilter }

func (keyGenerateRsa *KeyGenerateRsa) SetFilter(yf yfilter.YFilter) { keyGenerateRsa.YFilter = yf }

func (keyGenerateRsa *KeyGenerateRsa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (keyGenerateRsa *KeyGenerateRsa) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-generate-rsa"
}

func (keyGenerateRsa *KeyGenerateRsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &keyGenerateRsa.Input
    }
    return nil
}

func (keyGenerateRsa *KeyGenerateRsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &keyGenerateRsa.Input
    return children
}

func (keyGenerateRsa *KeyGenerateRsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyGenerateRsa *KeyGenerateRsa) GetBundleName() string { return "cisco_ios_xr" }

func (keyGenerateRsa *KeyGenerateRsa) GetYangName() string { return "key-generate-rsa" }

func (keyGenerateRsa *KeyGenerateRsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyGenerateRsa *KeyGenerateRsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyGenerateRsa *KeyGenerateRsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyGenerateRsa *KeyGenerateRsa) SetParent(parent types.Entity) { keyGenerateRsa.parent = parent }

func (keyGenerateRsa *KeyGenerateRsa) GetParent() types.Entity { return keyGenerateRsa.parent }

func (keyGenerateRsa *KeyGenerateRsa) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyGenerateRsa_Input
type KeyGenerateRsa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSA keypair label. The type is string.
    KeyLabel interface{}

    // Key modulus in the range of 512 to 4096 for your General Purpose Keypair.
    // Choosing a key modulus greater than 512 may take a few minutes. The type is
    // interface{} with range: 512..4096. This attribute is mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateRsa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *KeyGenerateRsa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *KeyGenerateRsa_Input) GetGoName(yname string) string {
    if yname == "key-label" { return "KeyLabel" }
    if yname == "key-modulus" { return "KeyModulus" }
    return ""
}

func (input *KeyGenerateRsa_Input) GetSegmentPath() string {
    return "input"
}

func (input *KeyGenerateRsa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *KeyGenerateRsa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *KeyGenerateRsa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-label"] = input.KeyLabel
    leafs["key-modulus"] = input.KeyModulus
    return leafs
}

func (input *KeyGenerateRsa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *KeyGenerateRsa_Input) GetYangName() string { return "input" }

func (input *KeyGenerateRsa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *KeyGenerateRsa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *KeyGenerateRsa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *KeyGenerateRsa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *KeyGenerateRsa_Input) GetParent() types.Entity { return input.parent }

func (input *KeyGenerateRsa_Input) GetParentYangName() string { return "key-generate-rsa" }

// KeyGenerateDsa
// Generate DSA keys
type KeyGenerateDsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input KeyGenerateDsa_Input
}

func (keyGenerateDsa *KeyGenerateDsa) GetFilter() yfilter.YFilter { return keyGenerateDsa.YFilter }

func (keyGenerateDsa *KeyGenerateDsa) SetFilter(yf yfilter.YFilter) { keyGenerateDsa.YFilter = yf }

func (keyGenerateDsa *KeyGenerateDsa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (keyGenerateDsa *KeyGenerateDsa) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-generate-dsa"
}

func (keyGenerateDsa *KeyGenerateDsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &keyGenerateDsa.Input
    }
    return nil
}

func (keyGenerateDsa *KeyGenerateDsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &keyGenerateDsa.Input
    return children
}

func (keyGenerateDsa *KeyGenerateDsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyGenerateDsa *KeyGenerateDsa) GetBundleName() string { return "cisco_ios_xr" }

func (keyGenerateDsa *KeyGenerateDsa) GetYangName() string { return "key-generate-dsa" }

func (keyGenerateDsa *KeyGenerateDsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyGenerateDsa *KeyGenerateDsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyGenerateDsa *KeyGenerateDsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyGenerateDsa *KeyGenerateDsa) SetParent(parent types.Entity) { keyGenerateDsa.parent = parent }

func (keyGenerateDsa *KeyGenerateDsa) GetParent() types.Entity { return keyGenerateDsa.parent }

func (keyGenerateDsa *KeyGenerateDsa) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyGenerateDsa_Input
type KeyGenerateDsa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key modulus size can be 512, 768, or 1024 bits. The type is interface{}
    // with range: 512..None | 768..None | 1024..None. This attribute is
    // mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateDsa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *KeyGenerateDsa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *KeyGenerateDsa_Input) GetGoName(yname string) string {
    if yname == "key-modulus" { return "KeyModulus" }
    return ""
}

func (input *KeyGenerateDsa_Input) GetSegmentPath() string {
    return "input"
}

func (input *KeyGenerateDsa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *KeyGenerateDsa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *KeyGenerateDsa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-modulus"] = input.KeyModulus
    return leafs
}

func (input *KeyGenerateDsa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *KeyGenerateDsa_Input) GetYangName() string { return "input" }

func (input *KeyGenerateDsa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *KeyGenerateDsa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *KeyGenerateDsa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *KeyGenerateDsa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *KeyGenerateDsa_Input) GetParent() types.Entity { return input.parent }

func (input *KeyGenerateDsa_Input) GetParentYangName() string { return "key-generate-dsa" }

// KeyZeroizeRsa
// Remove RSA keys
type KeyZeroizeRsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input KeyZeroizeRsa_Input
}

func (keyZeroizeRsa *KeyZeroizeRsa) GetFilter() yfilter.YFilter { return keyZeroizeRsa.YFilter }

func (keyZeroizeRsa *KeyZeroizeRsa) SetFilter(yf yfilter.YFilter) { keyZeroizeRsa.YFilter = yf }

func (keyZeroizeRsa *KeyZeroizeRsa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (keyZeroizeRsa *KeyZeroizeRsa) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-zeroize-rsa"
}

func (keyZeroizeRsa *KeyZeroizeRsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &keyZeroizeRsa.Input
    }
    return nil
}

func (keyZeroizeRsa *KeyZeroizeRsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &keyZeroizeRsa.Input
    return children
}

func (keyZeroizeRsa *KeyZeroizeRsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyZeroizeRsa *KeyZeroizeRsa) GetBundleName() string { return "cisco_ios_xr" }

func (keyZeroizeRsa *KeyZeroizeRsa) GetYangName() string { return "key-zeroize-rsa" }

func (keyZeroizeRsa *KeyZeroizeRsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyZeroizeRsa *KeyZeroizeRsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyZeroizeRsa *KeyZeroizeRsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyZeroizeRsa *KeyZeroizeRsa) SetParent(parent types.Entity) { keyZeroizeRsa.parent = parent }

func (keyZeroizeRsa *KeyZeroizeRsa) GetParent() types.Entity { return keyZeroizeRsa.parent }

func (keyZeroizeRsa *KeyZeroizeRsa) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyZeroizeRsa_Input
type KeyZeroizeRsa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSA key label. The type is string.
    KeyLabel interface{}
}

func (input *KeyZeroizeRsa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *KeyZeroizeRsa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *KeyZeroizeRsa_Input) GetGoName(yname string) string {
    if yname == "key-label" { return "KeyLabel" }
    return ""
}

func (input *KeyZeroizeRsa_Input) GetSegmentPath() string {
    return "input"
}

func (input *KeyZeroizeRsa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *KeyZeroizeRsa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *KeyZeroizeRsa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-label"] = input.KeyLabel
    return leafs
}

func (input *KeyZeroizeRsa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *KeyZeroizeRsa_Input) GetYangName() string { return "input" }

func (input *KeyZeroizeRsa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *KeyZeroizeRsa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *KeyZeroizeRsa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *KeyZeroizeRsa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *KeyZeroizeRsa_Input) GetParent() types.Entity { return input.parent }

func (input *KeyZeroizeRsa_Input) GetParentYangName() string { return "key-zeroize-rsa" }

// KeyZeroizeDsa
// Remove DSA keys
type KeyZeroizeDsa struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (keyZeroizeDsa *KeyZeroizeDsa) GetFilter() yfilter.YFilter { return keyZeroizeDsa.YFilter }

func (keyZeroizeDsa *KeyZeroizeDsa) SetFilter(yf yfilter.YFilter) { keyZeroizeDsa.YFilter = yf }

func (keyZeroizeDsa *KeyZeroizeDsa) GetGoName(yname string) string {
    return ""
}

func (keyZeroizeDsa *KeyZeroizeDsa) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-zeroize-dsa"
}

func (keyZeroizeDsa *KeyZeroizeDsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keyZeroizeDsa *KeyZeroizeDsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keyZeroizeDsa *KeyZeroizeDsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyZeroizeDsa *KeyZeroizeDsa) GetBundleName() string { return "cisco_ios_xr" }

func (keyZeroizeDsa *KeyZeroizeDsa) GetYangName() string { return "key-zeroize-dsa" }

func (keyZeroizeDsa *KeyZeroizeDsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyZeroizeDsa *KeyZeroizeDsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyZeroizeDsa *KeyZeroizeDsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyZeroizeDsa *KeyZeroizeDsa) SetParent(parent types.Entity) { keyZeroizeDsa.parent = parent }

func (keyZeroizeDsa *KeyZeroizeDsa) GetParent() types.Entity { return keyZeroizeDsa.parent }

func (keyZeroizeDsa *KeyZeroizeDsa) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyZeroizeAuthenticationRsa
// Remove RSA authentication key
type KeyZeroizeAuthenticationRsa struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetFilter() yfilter.YFilter { return keyZeroizeAuthenticationRsa.YFilter }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) SetFilter(yf yfilter.YFilter) { keyZeroizeAuthenticationRsa.YFilter = yf }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetGoName(yname string) string {
    return ""
}

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-zeroize-authentication-rsa"
}

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetBundleName() string { return "cisco_ios_xr" }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetYangName() string { return "key-zeroize-authentication-rsa" }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) SetParent(parent types.Entity) { keyZeroizeAuthenticationRsa.parent = parent }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetParent() types.Entity { return keyZeroizeAuthenticationRsa.parent }

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyImportAuthenticationRsa
// Remove RSA authentication key
type KeyImportAuthenticationRsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input KeyImportAuthenticationRsa_Input
}

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetFilter() yfilter.YFilter { return keyImportAuthenticationRsa.YFilter }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) SetFilter(yf yfilter.YFilter) { keyImportAuthenticationRsa.YFilter = yf }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:key-import-authentication-rsa"
}

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &keyImportAuthenticationRsa.Input
    }
    return nil
}

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &keyImportAuthenticationRsa.Input
    return children
}

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetBundleName() string { return "cisco_ios_xr" }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetYangName() string { return "key-import-authentication-rsa" }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) SetParent(parent types.Entity) { keyImportAuthenticationRsa.parent = parent }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetParent() types.Entity { return keyImportAuthenticationRsa.parent }

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// KeyImportAuthenticationRsa_Input
type KeyImportAuthenticationRsa_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path to RSA pubkey file. The type is string. This attribute is mandatory.
    Path interface{}
}

func (input *KeyImportAuthenticationRsa_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *KeyImportAuthenticationRsa_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *KeyImportAuthenticationRsa_Input) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (input *KeyImportAuthenticationRsa_Input) GetSegmentPath() string {
    return "input"
}

func (input *KeyImportAuthenticationRsa_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *KeyImportAuthenticationRsa_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *KeyImportAuthenticationRsa_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = input.Path
    return leafs
}

func (input *KeyImportAuthenticationRsa_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *KeyImportAuthenticationRsa_Input) GetYangName() string { return "input" }

func (input *KeyImportAuthenticationRsa_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *KeyImportAuthenticationRsa_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *KeyImportAuthenticationRsa_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *KeyImportAuthenticationRsa_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *KeyImportAuthenticationRsa_Input) GetParent() types.Entity { return input.parent }

func (input *KeyImportAuthenticationRsa_Input) GetParentYangName() string { return "key-import-authentication-rsa" }

// CaAuthenticate
// Get the certification authority certificate
type CaAuthenticate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CaAuthenticate_Input
}

func (caAuthenticate *CaAuthenticate) GetFilter() yfilter.YFilter { return caAuthenticate.YFilter }

func (caAuthenticate *CaAuthenticate) SetFilter(yf yfilter.YFilter) { caAuthenticate.YFilter = yf }

func (caAuthenticate *CaAuthenticate) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (caAuthenticate *CaAuthenticate) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:ca-authenticate"
}

func (caAuthenticate *CaAuthenticate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &caAuthenticate.Input
    }
    return nil
}

func (caAuthenticate *CaAuthenticate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &caAuthenticate.Input
    return children
}

func (caAuthenticate *CaAuthenticate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caAuthenticate *CaAuthenticate) GetBundleName() string { return "cisco_ios_xr" }

func (caAuthenticate *CaAuthenticate) GetYangName() string { return "ca-authenticate" }

func (caAuthenticate *CaAuthenticate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caAuthenticate *CaAuthenticate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caAuthenticate *CaAuthenticate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caAuthenticate *CaAuthenticate) SetParent(parent types.Entity) { caAuthenticate.parent = parent }

func (caAuthenticate *CaAuthenticate) GetParent() types.Entity { return caAuthenticate.parent }

func (caAuthenticate *CaAuthenticate) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// CaAuthenticate_Input
type CaAuthenticate_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaAuthenticate_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CaAuthenticate_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CaAuthenticate_Input) GetGoName(yname string) string {
    if yname == "server-name" { return "ServerName" }
    return ""
}

func (input *CaAuthenticate_Input) GetSegmentPath() string {
    return "input"
}

func (input *CaAuthenticate_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CaAuthenticate_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CaAuthenticate_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-name"] = input.ServerName
    return leafs
}

func (input *CaAuthenticate_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *CaAuthenticate_Input) GetYangName() string { return "input" }

func (input *CaAuthenticate_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *CaAuthenticate_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *CaAuthenticate_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *CaAuthenticate_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CaAuthenticate_Input) GetParent() types.Entity { return input.parent }

func (input *CaAuthenticate_Input) GetParentYangName() string { return "ca-authenticate" }

// CaEnroll
// Request a certificate from a CA
type CaEnroll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CaEnroll_Input
}

func (caEnroll *CaEnroll) GetFilter() yfilter.YFilter { return caEnroll.YFilter }

func (caEnroll *CaEnroll) SetFilter(yf yfilter.YFilter) { caEnroll.YFilter = yf }

func (caEnroll *CaEnroll) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (caEnroll *CaEnroll) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:ca-enroll"
}

func (caEnroll *CaEnroll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &caEnroll.Input
    }
    return nil
}

func (caEnroll *CaEnroll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &caEnroll.Input
    return children
}

func (caEnroll *CaEnroll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caEnroll *CaEnroll) GetBundleName() string { return "cisco_ios_xr" }

func (caEnroll *CaEnroll) GetYangName() string { return "ca-enroll" }

func (caEnroll *CaEnroll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caEnroll *CaEnroll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caEnroll *CaEnroll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caEnroll *CaEnroll) SetParent(parent types.Entity) { caEnroll.parent = parent }

func (caEnroll *CaEnroll) GetParent() types.Entity { return caEnroll.parent }

func (caEnroll *CaEnroll) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// CaEnroll_Input
type CaEnroll_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaEnroll_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CaEnroll_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CaEnroll_Input) GetGoName(yname string) string {
    if yname == "server-name" { return "ServerName" }
    return ""
}

func (input *CaEnroll_Input) GetSegmentPath() string {
    return "input"
}

func (input *CaEnroll_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CaEnroll_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CaEnroll_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-name"] = input.ServerName
    return leafs
}

func (input *CaEnroll_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *CaEnroll_Input) GetYangName() string { return "input" }

func (input *CaEnroll_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *CaEnroll_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *CaEnroll_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *CaEnroll_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CaEnroll_Input) GetParent() types.Entity { return input.parent }

func (input *CaEnroll_Input) GetParentYangName() string { return "ca-enroll" }

// CaImportCertificate
// Import a certificate from a s/tftp server or the terminal
type CaImportCertificate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CaImportCertificate_Input
}

func (caImportCertificate *CaImportCertificate) GetFilter() yfilter.YFilter { return caImportCertificate.YFilter }

func (caImportCertificate *CaImportCertificate) SetFilter(yf yfilter.YFilter) { caImportCertificate.YFilter = yf }

func (caImportCertificate *CaImportCertificate) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (caImportCertificate *CaImportCertificate) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:ca-import-certificate"
}

func (caImportCertificate *CaImportCertificate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &caImportCertificate.Input
    }
    return nil
}

func (caImportCertificate *CaImportCertificate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &caImportCertificate.Input
    return children
}

func (caImportCertificate *CaImportCertificate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caImportCertificate *CaImportCertificate) GetBundleName() string { return "cisco_ios_xr" }

func (caImportCertificate *CaImportCertificate) GetYangName() string { return "ca-import-certificate" }

func (caImportCertificate *CaImportCertificate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caImportCertificate *CaImportCertificate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caImportCertificate *CaImportCertificate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caImportCertificate *CaImportCertificate) SetParent(parent types.Entity) { caImportCertificate.parent = parent }

func (caImportCertificate *CaImportCertificate) GetParent() types.Entity { return caImportCertificate.parent }

func (caImportCertificate *CaImportCertificate) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// CaImportCertificate_Input
type CaImportCertificate_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaImportCertificate_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CaImportCertificate_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CaImportCertificate_Input) GetGoName(yname string) string {
    if yname == "server-name" { return "ServerName" }
    return ""
}

func (input *CaImportCertificate_Input) GetSegmentPath() string {
    return "input"
}

func (input *CaImportCertificate_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CaImportCertificate_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CaImportCertificate_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-name"] = input.ServerName
    return leafs
}

func (input *CaImportCertificate_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *CaImportCertificate_Input) GetYangName() string { return "input" }

func (input *CaImportCertificate_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *CaImportCertificate_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *CaImportCertificate_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *CaImportCertificate_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CaImportCertificate_Input) GetParent() types.Entity { return input.parent }

func (input *CaImportCertificate_Input) GetParentYangName() string { return "ca-import-certificate" }

// CaCancelEnroll
// Cancel enrollment from a CA
type CaCancelEnroll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CaCancelEnroll_Input
}

func (caCancelEnroll *CaCancelEnroll) GetFilter() yfilter.YFilter { return caCancelEnroll.YFilter }

func (caCancelEnroll *CaCancelEnroll) SetFilter(yf yfilter.YFilter) { caCancelEnroll.YFilter = yf }

func (caCancelEnroll *CaCancelEnroll) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (caCancelEnroll *CaCancelEnroll) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:ca-cancel-enroll"
}

func (caCancelEnroll *CaCancelEnroll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &caCancelEnroll.Input
    }
    return nil
}

func (caCancelEnroll *CaCancelEnroll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &caCancelEnroll.Input
    return children
}

func (caCancelEnroll *CaCancelEnroll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caCancelEnroll *CaCancelEnroll) GetBundleName() string { return "cisco_ios_xr" }

func (caCancelEnroll *CaCancelEnroll) GetYangName() string { return "ca-cancel-enroll" }

func (caCancelEnroll *CaCancelEnroll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caCancelEnroll *CaCancelEnroll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caCancelEnroll *CaCancelEnroll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caCancelEnroll *CaCancelEnroll) SetParent(parent types.Entity) { caCancelEnroll.parent = parent }

func (caCancelEnroll *CaCancelEnroll) GetParent() types.Entity { return caCancelEnroll.parent }

func (caCancelEnroll *CaCancelEnroll) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// CaCancelEnroll_Input
type CaCancelEnroll_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaCancelEnroll_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CaCancelEnroll_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CaCancelEnroll_Input) GetGoName(yname string) string {
    if yname == "server-name" { return "ServerName" }
    return ""
}

func (input *CaCancelEnroll_Input) GetSegmentPath() string {
    return "input"
}

func (input *CaCancelEnroll_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CaCancelEnroll_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CaCancelEnroll_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-name"] = input.ServerName
    return leafs
}

func (input *CaCancelEnroll_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *CaCancelEnroll_Input) GetYangName() string { return "input" }

func (input *CaCancelEnroll_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *CaCancelEnroll_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *CaCancelEnroll_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *CaCancelEnroll_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CaCancelEnroll_Input) GetParent() types.Entity { return input.parent }

func (input *CaCancelEnroll_Input) GetParentYangName() string { return "ca-cancel-enroll" }

// CaCrlRequest
// Actions on certificate revocation lists
type CaCrlRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CaCrlRequest_Input

    
    Output CaCrlRequest_Output
}

func (caCrlRequest *CaCrlRequest) GetFilter() yfilter.YFilter { return caCrlRequest.YFilter }

func (caCrlRequest *CaCrlRequest) SetFilter(yf yfilter.YFilter) { caCrlRequest.YFilter = yf }

func (caCrlRequest *CaCrlRequest) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (caCrlRequest *CaCrlRequest) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:ca-crl-request"
}

func (caCrlRequest *CaCrlRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &caCrlRequest.Input
    }
    if childYangName == "output" {
        return &caCrlRequest.Output
    }
    return nil
}

func (caCrlRequest *CaCrlRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &caCrlRequest.Input
    children["output"] = &caCrlRequest.Output
    return children
}

func (caCrlRequest *CaCrlRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caCrlRequest *CaCrlRequest) GetBundleName() string { return "cisco_ios_xr" }

func (caCrlRequest *CaCrlRequest) GetYangName() string { return "ca-crl-request" }

func (caCrlRequest *CaCrlRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caCrlRequest *CaCrlRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caCrlRequest *CaCrlRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caCrlRequest *CaCrlRequest) SetParent(parent types.Entity) { caCrlRequest.parent = parent }

func (caCrlRequest *CaCrlRequest) GetParent() types.Entity { return caCrlRequest.parent }

func (caCrlRequest *CaCrlRequest) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// CaCrlRequest_Input
type CaCrlRequest_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CRL Distribution Point in URI format. The type is string. This attribute is
    // mandatory.
    Uri interface{}
}

func (input *CaCrlRequest_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CaCrlRequest_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CaCrlRequest_Input) GetGoName(yname string) string {
    if yname == "uri" { return "Uri" }
    return ""
}

func (input *CaCrlRequest_Input) GetSegmentPath() string {
    return "input"
}

func (input *CaCrlRequest_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CaCrlRequest_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CaCrlRequest_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["uri"] = input.Uri
    return leafs
}

func (input *CaCrlRequest_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *CaCrlRequest_Input) GetYangName() string { return "input" }

func (input *CaCrlRequest_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *CaCrlRequest_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *CaCrlRequest_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *CaCrlRequest_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CaCrlRequest_Input) GetParent() types.Entity { return input.parent }

func (input *CaCrlRequest_Input) GetParentYangName() string { return "ca-crl-request" }

// CaCrlRequest_Output
type CaCrlRequest_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Certificate returned. The type is string. This attribute is mandatory.
    Certificate interface{}
}

func (output *CaCrlRequest_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *CaCrlRequest_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *CaCrlRequest_Output) GetGoName(yname string) string {
    if yname == "certificate" { return "Certificate" }
    return ""
}

func (output *CaCrlRequest_Output) GetSegmentPath() string {
    return "output"
}

func (output *CaCrlRequest_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *CaCrlRequest_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *CaCrlRequest_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["certificate"] = output.Certificate
    return leafs
}

func (output *CaCrlRequest_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *CaCrlRequest_Output) GetYangName() string { return "output" }

func (output *CaCrlRequest_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *CaCrlRequest_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *CaCrlRequest_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *CaCrlRequest_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *CaCrlRequest_Output) GetParent() types.Entity { return output.parent }

func (output *CaCrlRequest_Output) GetParentYangName() string { return "ca-crl-request" }

// CaTrustpoolImportUrl
// Manual import trustpool certificates from URL
type CaTrustpoolImportUrl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CaTrustpoolImportUrl_Input
}

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetFilter() yfilter.YFilter { return caTrustpoolImportUrl.YFilter }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) SetFilter(yf yfilter.YFilter) { caTrustpoolImportUrl.YFilter = yf }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:ca-trustpool-import-url"
}

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &caTrustpoolImportUrl.Input
    }
    return nil
}

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &caTrustpoolImportUrl.Input
    return children
}

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetBundleName() string { return "cisco_ios_xr" }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetYangName() string { return "ca-trustpool-import-url" }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) SetParent(parent types.Entity) { caTrustpoolImportUrl.parent = parent }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetParent() types.Entity { return caTrustpoolImportUrl.parent }

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// CaTrustpoolImportUrl_Input
type CaTrustpoolImportUrl_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // in URL format. The type is string.
    Url interface{}
}

func (input *CaTrustpoolImportUrl_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CaTrustpoolImportUrl_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CaTrustpoolImportUrl_Input) GetGoName(yname string) string {
    if yname == "url" { return "Url" }
    return ""
}

func (input *CaTrustpoolImportUrl_Input) GetSegmentPath() string {
    return "input"
}

func (input *CaTrustpoolImportUrl_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CaTrustpoolImportUrl_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CaTrustpoolImportUrl_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["url"] = input.Url
    return leafs
}

func (input *CaTrustpoolImportUrl_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *CaTrustpoolImportUrl_Input) GetYangName() string { return "input" }

func (input *CaTrustpoolImportUrl_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *CaTrustpoolImportUrl_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *CaTrustpoolImportUrl_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *CaTrustpoolImportUrl_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CaTrustpoolImportUrl_Input) GetParent() types.Entity { return input.parent }

func (input *CaTrustpoolImportUrl_Input) GetParentYangName() string { return "ca-trustpool-import-url" }

// CaTrustpoolImportUrlClean
// Remove downloaded certificates in trustpool
type CaTrustpoolImportUrlClean struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CaTrustpoolImportUrlClean_Input
}

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetFilter() yfilter.YFilter { return caTrustpoolImportUrlClean.YFilter }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) SetFilter(yf yfilter.YFilter) { caTrustpoolImportUrlClean.YFilter = yf }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-act:ca-trustpool-import-url-clean"
}

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &caTrustpoolImportUrlClean.Input
    }
    return nil
}

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &caTrustpoolImportUrlClean.Input
    return children
}

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetBundleName() string { return "cisco_ios_xr" }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetYangName() string { return "ca-trustpool-import-url-clean" }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) SetParent(parent types.Entity) { caTrustpoolImportUrlClean.parent = parent }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetParent() types.Entity { return caTrustpoolImportUrlClean.parent }

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetParentYangName() string { return "Cisco-IOS-XR-crypto-act" }

// CaTrustpoolImportUrlClean_Input
type CaTrustpoolImportUrlClean_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // in URL format. The type is string.
    Url interface{}
}

func (input *CaTrustpoolImportUrlClean_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CaTrustpoolImportUrlClean_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CaTrustpoolImportUrlClean_Input) GetGoName(yname string) string {
    if yname == "url" { return "Url" }
    return ""
}

func (input *CaTrustpoolImportUrlClean_Input) GetSegmentPath() string {
    return "input"
}

func (input *CaTrustpoolImportUrlClean_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CaTrustpoolImportUrlClean_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CaTrustpoolImportUrlClean_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["url"] = input.Url
    return leafs
}

func (input *CaTrustpoolImportUrlClean_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *CaTrustpoolImportUrlClean_Input) GetYangName() string { return "input" }

func (input *CaTrustpoolImportUrlClean_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *CaTrustpoolImportUrlClean_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *CaTrustpoolImportUrlClean_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *CaTrustpoolImportUrlClean_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CaTrustpoolImportUrlClean_Input) GetParent() types.Entity { return input.parent }

func (input *CaTrustpoolImportUrlClean_Input) GetParentYangName() string { return "ca-trustpool-import-url-clean" }

