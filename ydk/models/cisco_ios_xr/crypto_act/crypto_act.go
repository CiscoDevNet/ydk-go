// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-generate-ecdsa}", reflect.TypeOf(KeyGenerateEcdsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-generate-ecdsa", reflect.TypeOf(KeyGenerateEcdsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-zeroize-rsa}", reflect.TypeOf(KeyZeroizeRsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-zeroize-rsa", reflect.TypeOf(KeyZeroizeRsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-zeroize-dsa}", reflect.TypeOf(KeyZeroizeDsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-zeroize-dsa", reflect.TypeOf(KeyZeroizeDsa{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-act key-zeroize-ecdsa}", reflect.TypeOf(KeyZeroizeEcdsa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-act:key-zeroize-ecdsa", reflect.TypeOf(KeyZeroizeEcdsa{}))
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyGenerateRsaGeneralKeys_Input
}

func (keyGenerateRsaGeneralKeys *KeyGenerateRsaGeneralKeys) GetEntityData() *types.CommonEntityData {
    keyGenerateRsaGeneralKeys.EntityData.YFilter = keyGenerateRsaGeneralKeys.YFilter
    keyGenerateRsaGeneralKeys.EntityData.YangName = "key-generate-rsa-general-keys"
    keyGenerateRsaGeneralKeys.EntityData.BundleName = "cisco_ios_xr"
    keyGenerateRsaGeneralKeys.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyGenerateRsaGeneralKeys.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-generate-rsa-general-keys"
    keyGenerateRsaGeneralKeys.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyGenerateRsaGeneralKeys.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyGenerateRsaGeneralKeys.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyGenerateRsaGeneralKeys.EntityData.Children = types.NewOrderedMap()
    keyGenerateRsaGeneralKeys.EntityData.Children.Append("input", types.YChild{"Input", &keyGenerateRsaGeneralKeys.Input})
    keyGenerateRsaGeneralKeys.EntityData.Leafs = types.NewOrderedMap()

    keyGenerateRsaGeneralKeys.EntityData.YListKeys = []string {}

    return &(keyGenerateRsaGeneralKeys.EntityData)
}

// KeyGenerateRsaGeneralKeys_Input
type KeyGenerateRsaGeneralKeys_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSA keypair label. The type is string.
    KeyLabel interface{}

    // Key modulus in the range of 512 to 4096 for your General Purpose Keypair.
    // Choosing a key modulus greater than 512 may take a few minutes. The type is
    // interface{} with range: 512..4096. This attribute is mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateRsaGeneralKeys_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-generate-rsa-general-keys"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("key-label", types.YLeaf{"KeyLabel", input.KeyLabel})
    input.EntityData.Leafs.Append("key-modulus", types.YLeaf{"KeyModulus", input.KeyModulus})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// KeyGenerateRsaUsageKeys
// Generate seperate RSA key pairs for signing and encryption
type KeyGenerateRsaUsageKeys struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyGenerateRsaUsageKeys_Input
}

func (keyGenerateRsaUsageKeys *KeyGenerateRsaUsageKeys) GetEntityData() *types.CommonEntityData {
    keyGenerateRsaUsageKeys.EntityData.YFilter = keyGenerateRsaUsageKeys.YFilter
    keyGenerateRsaUsageKeys.EntityData.YangName = "key-generate-rsa-usage-keys"
    keyGenerateRsaUsageKeys.EntityData.BundleName = "cisco_ios_xr"
    keyGenerateRsaUsageKeys.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyGenerateRsaUsageKeys.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-generate-rsa-usage-keys"
    keyGenerateRsaUsageKeys.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyGenerateRsaUsageKeys.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyGenerateRsaUsageKeys.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyGenerateRsaUsageKeys.EntityData.Children = types.NewOrderedMap()
    keyGenerateRsaUsageKeys.EntityData.Children.Append("input", types.YChild{"Input", &keyGenerateRsaUsageKeys.Input})
    keyGenerateRsaUsageKeys.EntityData.Leafs = types.NewOrderedMap()

    keyGenerateRsaUsageKeys.EntityData.YListKeys = []string {}

    return &(keyGenerateRsaUsageKeys.EntityData)
}

// KeyGenerateRsaUsageKeys_Input
type KeyGenerateRsaUsageKeys_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSA keypair label. The type is string.
    KeyLabel interface{}

    // Key modulus in the range of 512 to 4096 for your General Purpose Keypair.
    // Choosing a key modulus greater than 512 may take a few minutes. The type is
    // interface{} with range: 512..4096. This attribute is mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateRsaUsageKeys_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-generate-rsa-usage-keys"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("key-label", types.YLeaf{"KeyLabel", input.KeyLabel})
    input.EntityData.Leafs.Append("key-modulus", types.YLeaf{"KeyModulus", input.KeyModulus})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// KeyGenerateRsa
// Generate seperate RSA key pairs for signing and encryption
type KeyGenerateRsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyGenerateRsa_Input
}

func (keyGenerateRsa *KeyGenerateRsa) GetEntityData() *types.CommonEntityData {
    keyGenerateRsa.EntityData.YFilter = keyGenerateRsa.YFilter
    keyGenerateRsa.EntityData.YangName = "key-generate-rsa"
    keyGenerateRsa.EntityData.BundleName = "cisco_ios_xr"
    keyGenerateRsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyGenerateRsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-generate-rsa"
    keyGenerateRsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyGenerateRsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyGenerateRsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyGenerateRsa.EntityData.Children = types.NewOrderedMap()
    keyGenerateRsa.EntityData.Children.Append("input", types.YChild{"Input", &keyGenerateRsa.Input})
    keyGenerateRsa.EntityData.Leafs = types.NewOrderedMap()

    keyGenerateRsa.EntityData.YListKeys = []string {}

    return &(keyGenerateRsa.EntityData)
}

// KeyGenerateRsa_Input
type KeyGenerateRsa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSA keypair label. The type is string.
    KeyLabel interface{}

    // Key modulus in the range of 512 to 4096 for your General Purpose Keypair.
    // Choosing a key modulus greater than 512 may take a few minutes. The type is
    // interface{} with range: 512..4096. This attribute is mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateRsa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-generate-rsa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("key-label", types.YLeaf{"KeyLabel", input.KeyLabel})
    input.EntityData.Leafs.Append("key-modulus", types.YLeaf{"KeyModulus", input.KeyModulus})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// KeyGenerateDsa
// Generate DSA keys
type KeyGenerateDsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyGenerateDsa_Input
}

func (keyGenerateDsa *KeyGenerateDsa) GetEntityData() *types.CommonEntityData {
    keyGenerateDsa.EntityData.YFilter = keyGenerateDsa.YFilter
    keyGenerateDsa.EntityData.YangName = "key-generate-dsa"
    keyGenerateDsa.EntityData.BundleName = "cisco_ios_xr"
    keyGenerateDsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyGenerateDsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-generate-dsa"
    keyGenerateDsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyGenerateDsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyGenerateDsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyGenerateDsa.EntityData.Children = types.NewOrderedMap()
    keyGenerateDsa.EntityData.Children.Append("input", types.YChild{"Input", &keyGenerateDsa.Input})
    keyGenerateDsa.EntityData.Leafs = types.NewOrderedMap()

    keyGenerateDsa.EntityData.YListKeys = []string {}

    return &(keyGenerateDsa.EntityData)
}

// KeyGenerateDsa_Input
type KeyGenerateDsa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key modulus size can be 512, 768, 1024 or 2048 bits. The type is
    // interface{} with range: 512..None | 768..None | 1024..None | 2048..None.
    // This attribute is mandatory.
    KeyModulus interface{}
}

func (input *KeyGenerateDsa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-generate-dsa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("key-modulus", types.YLeaf{"KeyModulus", input.KeyModulus})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// KeyGenerateEcdsa
// Generate a ECDSA key of curve type nistp256 | nistp384 | nistp521
type KeyGenerateEcdsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyGenerateEcdsa_Input
}

func (keyGenerateEcdsa *KeyGenerateEcdsa) GetEntityData() *types.CommonEntityData {
    keyGenerateEcdsa.EntityData.YFilter = keyGenerateEcdsa.YFilter
    keyGenerateEcdsa.EntityData.YangName = "key-generate-ecdsa"
    keyGenerateEcdsa.EntityData.BundleName = "cisco_ios_xr"
    keyGenerateEcdsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyGenerateEcdsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-generate-ecdsa"
    keyGenerateEcdsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyGenerateEcdsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyGenerateEcdsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyGenerateEcdsa.EntityData.Children = types.NewOrderedMap()
    keyGenerateEcdsa.EntityData.Children.Append("input", types.YChild{"Input", &keyGenerateEcdsa.Input})
    keyGenerateEcdsa.EntityData.Leafs = types.NewOrderedMap()

    keyGenerateEcdsa.EntityData.YListKeys = []string {}

    return &(keyGenerateEcdsa.EntityData)
}

// KeyGenerateEcdsa_Input
type KeyGenerateEcdsa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is KeyModulus. This attribute is mandatory.
    KeyModulus interface{}

    // ECDSA key label. The type is string.
    KeyLabel interface{}
}

func (input *KeyGenerateEcdsa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-generate-ecdsa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("key-modulus", types.YLeaf{"KeyModulus", input.KeyModulus})
    input.EntityData.Leafs.Append("key-label", types.YLeaf{"KeyLabel", input.KeyLabel})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// KeyGenerateEcdsa_Input_KeyModulus
type KeyGenerateEcdsa_Input_KeyModulus string

const (
    KeyGenerateEcdsa_Input_KeyModulus_nistp256 KeyGenerateEcdsa_Input_KeyModulus = "nistp256"

    KeyGenerateEcdsa_Input_KeyModulus_nistp384 KeyGenerateEcdsa_Input_KeyModulus = "nistp384"

    KeyGenerateEcdsa_Input_KeyModulus_nistp521 KeyGenerateEcdsa_Input_KeyModulus = "nistp521"
)

// KeyZeroizeRsa
// Remove RSA keys
type KeyZeroizeRsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyZeroizeRsa_Input
}

func (keyZeroizeRsa *KeyZeroizeRsa) GetEntityData() *types.CommonEntityData {
    keyZeroizeRsa.EntityData.YFilter = keyZeroizeRsa.YFilter
    keyZeroizeRsa.EntityData.YangName = "key-zeroize-rsa"
    keyZeroizeRsa.EntityData.BundleName = "cisco_ios_xr"
    keyZeroizeRsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyZeroizeRsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-zeroize-rsa"
    keyZeroizeRsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyZeroizeRsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyZeroizeRsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyZeroizeRsa.EntityData.Children = types.NewOrderedMap()
    keyZeroizeRsa.EntityData.Children.Append("input", types.YChild{"Input", &keyZeroizeRsa.Input})
    keyZeroizeRsa.EntityData.Leafs = types.NewOrderedMap()

    keyZeroizeRsa.EntityData.YListKeys = []string {}

    return &(keyZeroizeRsa.EntityData)
}

// KeyZeroizeRsa_Input
type KeyZeroizeRsa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSA key label. The type is string.
    KeyLabel interface{}
}

func (input *KeyZeroizeRsa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-zeroize-rsa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("key-label", types.YLeaf{"KeyLabel", input.KeyLabel})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// KeyZeroizeDsa
// Remove DSA keys
type KeyZeroizeDsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (keyZeroizeDsa *KeyZeroizeDsa) GetEntityData() *types.CommonEntityData {
    keyZeroizeDsa.EntityData.YFilter = keyZeroizeDsa.YFilter
    keyZeroizeDsa.EntityData.YangName = "key-zeroize-dsa"
    keyZeroizeDsa.EntityData.BundleName = "cisco_ios_xr"
    keyZeroizeDsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyZeroizeDsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-zeroize-dsa"
    keyZeroizeDsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyZeroizeDsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyZeroizeDsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyZeroizeDsa.EntityData.Children = types.NewOrderedMap()
    keyZeroizeDsa.EntityData.Leafs = types.NewOrderedMap()

    keyZeroizeDsa.EntityData.YListKeys = []string {}

    return &(keyZeroizeDsa.EntityData)
}

// KeyZeroizeEcdsa
// Remove ECDSA key of curve type nistp256 | nistp384 | nistp521
type KeyZeroizeEcdsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyZeroizeEcdsa_Input
}

func (keyZeroizeEcdsa *KeyZeroizeEcdsa) GetEntityData() *types.CommonEntityData {
    keyZeroizeEcdsa.EntityData.YFilter = keyZeroizeEcdsa.YFilter
    keyZeroizeEcdsa.EntityData.YangName = "key-zeroize-ecdsa"
    keyZeroizeEcdsa.EntityData.BundleName = "cisco_ios_xr"
    keyZeroizeEcdsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyZeroizeEcdsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-zeroize-ecdsa"
    keyZeroizeEcdsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyZeroizeEcdsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyZeroizeEcdsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyZeroizeEcdsa.EntityData.Children = types.NewOrderedMap()
    keyZeroizeEcdsa.EntityData.Children.Append("input", types.YChild{"Input", &keyZeroizeEcdsa.Input})
    keyZeroizeEcdsa.EntityData.Leafs = types.NewOrderedMap()

    keyZeroizeEcdsa.EntityData.YListKeys = []string {}

    return &(keyZeroizeEcdsa.EntityData)
}

// KeyZeroizeEcdsa_Input
type KeyZeroizeEcdsa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is KeyModulus. This attribute is mandatory.
    KeyModulus interface{}

    // ECDSA key label. The type is string.
    KeyLabel interface{}
}

func (input *KeyZeroizeEcdsa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-zeroize-ecdsa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("key-modulus", types.YLeaf{"KeyModulus", input.KeyModulus})
    input.EntityData.Leafs.Append("key-label", types.YLeaf{"KeyLabel", input.KeyLabel})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// KeyZeroizeEcdsa_Input_KeyModulus
type KeyZeroizeEcdsa_Input_KeyModulus string

const (
    KeyZeroizeEcdsa_Input_KeyModulus_nistp256 KeyZeroizeEcdsa_Input_KeyModulus = "nistp256"

    KeyZeroizeEcdsa_Input_KeyModulus_nistp384 KeyZeroizeEcdsa_Input_KeyModulus = "nistp384"

    KeyZeroizeEcdsa_Input_KeyModulus_nistp521 KeyZeroizeEcdsa_Input_KeyModulus = "nistp521"
)

// KeyZeroizeAuthenticationRsa
// Remove RSA authentication key
type KeyZeroizeAuthenticationRsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (keyZeroizeAuthenticationRsa *KeyZeroizeAuthenticationRsa) GetEntityData() *types.CommonEntityData {
    keyZeroizeAuthenticationRsa.EntityData.YFilter = keyZeroizeAuthenticationRsa.YFilter
    keyZeroizeAuthenticationRsa.EntityData.YangName = "key-zeroize-authentication-rsa"
    keyZeroizeAuthenticationRsa.EntityData.BundleName = "cisco_ios_xr"
    keyZeroizeAuthenticationRsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyZeroizeAuthenticationRsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-zeroize-authentication-rsa"
    keyZeroizeAuthenticationRsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyZeroizeAuthenticationRsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyZeroizeAuthenticationRsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyZeroizeAuthenticationRsa.EntityData.Children = types.NewOrderedMap()
    keyZeroizeAuthenticationRsa.EntityData.Leafs = types.NewOrderedMap()

    keyZeroizeAuthenticationRsa.EntityData.YListKeys = []string {}

    return &(keyZeroizeAuthenticationRsa.EntityData)
}

// KeyImportAuthenticationRsa
// Remove RSA authentication key
type KeyImportAuthenticationRsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KeyImportAuthenticationRsa_Input
}

func (keyImportAuthenticationRsa *KeyImportAuthenticationRsa) GetEntityData() *types.CommonEntityData {
    keyImportAuthenticationRsa.EntityData.YFilter = keyImportAuthenticationRsa.YFilter
    keyImportAuthenticationRsa.EntityData.YangName = "key-import-authentication-rsa"
    keyImportAuthenticationRsa.EntityData.BundleName = "cisco_ios_xr"
    keyImportAuthenticationRsa.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    keyImportAuthenticationRsa.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:key-import-authentication-rsa"
    keyImportAuthenticationRsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyImportAuthenticationRsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyImportAuthenticationRsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyImportAuthenticationRsa.EntityData.Children = types.NewOrderedMap()
    keyImportAuthenticationRsa.EntityData.Children.Append("input", types.YChild{"Input", &keyImportAuthenticationRsa.Input})
    keyImportAuthenticationRsa.EntityData.Leafs = types.NewOrderedMap()

    keyImportAuthenticationRsa.EntityData.YListKeys = []string {}

    return &(keyImportAuthenticationRsa.EntityData)
}

// KeyImportAuthenticationRsa_Input
type KeyImportAuthenticationRsa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path to RSA pubkey file. The type is string. This attribute is mandatory.
    Path interface{}
}

func (input *KeyImportAuthenticationRsa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "key-import-authentication-rsa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("path", types.YLeaf{"Path", input.Path})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CaAuthenticate
// Get the certification authority certificate
type CaAuthenticate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CaAuthenticate_Input
}

func (caAuthenticate *CaAuthenticate) GetEntityData() *types.CommonEntityData {
    caAuthenticate.EntityData.YFilter = caAuthenticate.YFilter
    caAuthenticate.EntityData.YangName = "ca-authenticate"
    caAuthenticate.EntityData.BundleName = "cisco_ios_xr"
    caAuthenticate.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    caAuthenticate.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:ca-authenticate"
    caAuthenticate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caAuthenticate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caAuthenticate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caAuthenticate.EntityData.Children = types.NewOrderedMap()
    caAuthenticate.EntityData.Children.Append("input", types.YChild{"Input", &caAuthenticate.Input})
    caAuthenticate.EntityData.Leafs = types.NewOrderedMap()

    caAuthenticate.EntityData.YListKeys = []string {}

    return &(caAuthenticate.EntityData)
}

// CaAuthenticate_Input
type CaAuthenticate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaAuthenticate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ca-authenticate"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("server-name", types.YLeaf{"ServerName", input.ServerName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CaEnroll
// Request a certificate from a CA
type CaEnroll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CaEnroll_Input
}

func (caEnroll *CaEnroll) GetEntityData() *types.CommonEntityData {
    caEnroll.EntityData.YFilter = caEnroll.YFilter
    caEnroll.EntityData.YangName = "ca-enroll"
    caEnroll.EntityData.BundleName = "cisco_ios_xr"
    caEnroll.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    caEnroll.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:ca-enroll"
    caEnroll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caEnroll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caEnroll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caEnroll.EntityData.Children = types.NewOrderedMap()
    caEnroll.EntityData.Children.Append("input", types.YChild{"Input", &caEnroll.Input})
    caEnroll.EntityData.Leafs = types.NewOrderedMap()

    caEnroll.EntityData.YListKeys = []string {}

    return &(caEnroll.EntityData)
}

// CaEnroll_Input
type CaEnroll_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaEnroll_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ca-enroll"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("server-name", types.YLeaf{"ServerName", input.ServerName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CaImportCertificate
// Import a certificate from a s/tftp server or the terminal
type CaImportCertificate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CaImportCertificate_Input
}

func (caImportCertificate *CaImportCertificate) GetEntityData() *types.CommonEntityData {
    caImportCertificate.EntityData.YFilter = caImportCertificate.YFilter
    caImportCertificate.EntityData.YangName = "ca-import-certificate"
    caImportCertificate.EntityData.BundleName = "cisco_ios_xr"
    caImportCertificate.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    caImportCertificate.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:ca-import-certificate"
    caImportCertificate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caImportCertificate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caImportCertificate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caImportCertificate.EntityData.Children = types.NewOrderedMap()
    caImportCertificate.EntityData.Children.Append("input", types.YChild{"Input", &caImportCertificate.Input})
    caImportCertificate.EntityData.Leafs = types.NewOrderedMap()

    caImportCertificate.EntityData.YListKeys = []string {}

    return &(caImportCertificate.EntityData)
}

// CaImportCertificate_Input
type CaImportCertificate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaImportCertificate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ca-import-certificate"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("server-name", types.YLeaf{"ServerName", input.ServerName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CaCancelEnroll
// Cancel enrollment from a CA
type CaCancelEnroll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CaCancelEnroll_Input
}

func (caCancelEnroll *CaCancelEnroll) GetEntityData() *types.CommonEntityData {
    caCancelEnroll.EntityData.YFilter = caCancelEnroll.YFilter
    caCancelEnroll.EntityData.YangName = "ca-cancel-enroll"
    caCancelEnroll.EntityData.BundleName = "cisco_ios_xr"
    caCancelEnroll.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    caCancelEnroll.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:ca-cancel-enroll"
    caCancelEnroll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caCancelEnroll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caCancelEnroll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caCancelEnroll.EntityData.Children = types.NewOrderedMap()
    caCancelEnroll.EntityData.Children.Append("input", types.YChild{"Input", &caCancelEnroll.Input})
    caCancelEnroll.EntityData.Leafs = types.NewOrderedMap()

    caCancelEnroll.EntityData.YListKeys = []string {}

    return &(caCancelEnroll.EntityData)
}

// CaCancelEnroll_Input
type CaCancelEnroll_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CA Server Name. The type is string. This attribute is mandatory.
    ServerName interface{}
}

func (input *CaCancelEnroll_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ca-cancel-enroll"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("server-name", types.YLeaf{"ServerName", input.ServerName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CaCrlRequest
// Actions on certificate revocation lists
type CaCrlRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CaCrlRequest_Input

    
    Output CaCrlRequest_Output
}

func (caCrlRequest *CaCrlRequest) GetEntityData() *types.CommonEntityData {
    caCrlRequest.EntityData.YFilter = caCrlRequest.YFilter
    caCrlRequest.EntityData.YangName = "ca-crl-request"
    caCrlRequest.EntityData.BundleName = "cisco_ios_xr"
    caCrlRequest.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    caCrlRequest.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:ca-crl-request"
    caCrlRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caCrlRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caCrlRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caCrlRequest.EntityData.Children = types.NewOrderedMap()
    caCrlRequest.EntityData.Children.Append("input", types.YChild{"Input", &caCrlRequest.Input})
    caCrlRequest.EntityData.Children.Append("output", types.YChild{"Output", &caCrlRequest.Output})
    caCrlRequest.EntityData.Leafs = types.NewOrderedMap()

    caCrlRequest.EntityData.YListKeys = []string {}

    return &(caCrlRequest.EntityData)
}

// CaCrlRequest_Input
type CaCrlRequest_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CRL Distribution Point in URI format. The type is string. This attribute is
    // mandatory.
    Uri interface{}
}

func (input *CaCrlRequest_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ca-crl-request"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("uri", types.YLeaf{"Uri", input.Uri})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CaCrlRequest_Output
type CaCrlRequest_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate returned. The type is string. This attribute is mandatory.
    Certificate interface{}
}

func (output *CaCrlRequest_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "ca-crl-request"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("certificate", types.YLeaf{"Certificate", output.Certificate})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// CaTrustpoolImportUrl
// Manual import trustpool certificates from URL
type CaTrustpoolImportUrl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CaTrustpoolImportUrl_Input
}

func (caTrustpoolImportUrl *CaTrustpoolImportUrl) GetEntityData() *types.CommonEntityData {
    caTrustpoolImportUrl.EntityData.YFilter = caTrustpoolImportUrl.YFilter
    caTrustpoolImportUrl.EntityData.YangName = "ca-trustpool-import-url"
    caTrustpoolImportUrl.EntityData.BundleName = "cisco_ios_xr"
    caTrustpoolImportUrl.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    caTrustpoolImportUrl.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:ca-trustpool-import-url"
    caTrustpoolImportUrl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caTrustpoolImportUrl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caTrustpoolImportUrl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caTrustpoolImportUrl.EntityData.Children = types.NewOrderedMap()
    caTrustpoolImportUrl.EntityData.Children.Append("input", types.YChild{"Input", &caTrustpoolImportUrl.Input})
    caTrustpoolImportUrl.EntityData.Leafs = types.NewOrderedMap()

    caTrustpoolImportUrl.EntityData.YListKeys = []string {}

    return &(caTrustpoolImportUrl.EntityData)
}

// CaTrustpoolImportUrl_Input
type CaTrustpoolImportUrl_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // in URL format. The type is string.
    Url interface{}
}

func (input *CaTrustpoolImportUrl_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ca-trustpool-import-url"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("url", types.YLeaf{"Url", input.Url})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CaTrustpoolImportUrlClean
// Remove downloaded certificates in trustpool
type CaTrustpoolImportUrlClean struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CaTrustpoolImportUrlClean_Input
}

func (caTrustpoolImportUrlClean *CaTrustpoolImportUrlClean) GetEntityData() *types.CommonEntityData {
    caTrustpoolImportUrlClean.EntityData.YFilter = caTrustpoolImportUrlClean.YFilter
    caTrustpoolImportUrlClean.EntityData.YangName = "ca-trustpool-import-url-clean"
    caTrustpoolImportUrlClean.EntityData.BundleName = "cisco_ios_xr"
    caTrustpoolImportUrlClean.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-act"
    caTrustpoolImportUrlClean.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-act:ca-trustpool-import-url-clean"
    caTrustpoolImportUrlClean.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caTrustpoolImportUrlClean.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caTrustpoolImportUrlClean.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caTrustpoolImportUrlClean.EntityData.Children = types.NewOrderedMap()
    caTrustpoolImportUrlClean.EntityData.Children.Append("input", types.YChild{"Input", &caTrustpoolImportUrlClean.Input})
    caTrustpoolImportUrlClean.EntityData.Leafs = types.NewOrderedMap()

    caTrustpoolImportUrlClean.EntityData.YListKeys = []string {}

    return &(caTrustpoolImportUrlClean.EntityData)
}

// CaTrustpoolImportUrlClean_Input
type CaTrustpoolImportUrlClean_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // in URL format. The type is string.
    Url interface{}
}

func (input *CaTrustpoolImportUrlClean_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ca-trustpool-import-url-clean"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("url", types.YLeaf{"Url", input.Url})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

