// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-mka package configuration.
// 
// This module contains definitions
// for the following management objects:
//   macsec: MACSec MKA
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package crypto_macsec_mka_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_macsec_mka_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-macsec-mka-cfg macsec}", reflect.TypeOf(Macsec{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-macsec-mka-cfg:macsec", reflect.TypeOf(Macsec{}))
}

// MacsecMkaConfOffset represents Macsec mka conf offset
type MacsecMkaConfOffset string

const (
    // CONF OFFSET 0
    MacsecMkaConfOffset_conf_off_set_0 MacsecMkaConfOffset = "conf-off-set-0"

    // CONF OFFSET 30
    MacsecMkaConfOffset_conf_off_set_30 MacsecMkaConfOffset = "conf-off-set-30"

    // CONF OFFSET 50
    MacsecMkaConfOffset_conf_off_set_50 MacsecMkaConfOffset = "conf-off-set-50"
)

// MacsecMkaSecurityPolicy represents Macsec mka security policy
type MacsecMkaSecurityPolicy string

const (
    // should secure
    MacsecMkaSecurityPolicy_should_secure MacsecMkaSecurityPolicy = "should-secure"

    // must secure
    MacsecMkaSecurityPolicy_must_secure MacsecMkaSecurityPolicy = "must-secure"
)

// MacsecMkaPolicyException represents Macsec mka policy exception
type MacsecMkaPolicyException string

const (
    // lacp in clear
    MacsecMkaPolicyException_lacp_in_clear MacsecMkaPolicyException = "lacp-in-clear"
)

// MacsecMkaCipherSuite represents Macsec mka cipher suite
type MacsecMkaCipherSuite string

const (
    // GCM AES 128
    MacsecMkaCipherSuite_gcm_aes_128 MacsecMkaCipherSuite = "gcm-aes-128"

    // GCM AES 256
    MacsecMkaCipherSuite_gcm_aes_256 MacsecMkaCipherSuite = "gcm-aes-256"

    // GCM AES XPN 128
    MacsecMkaCipherSuite_gcm_aes_xpn_128 MacsecMkaCipherSuite = "gcm-aes-xpn-128"

    // GCM AES XPN 256
    MacsecMkaCipherSuite_gcm_aes_xpn_256 MacsecMkaCipherSuite = "gcm-aes-xpn-256"
)

// Macsec
// MACSec MKA
type Macsec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable macsec on all data ports(system wide), has no impact on macsec
    // configs. The type is interface{}.
    Shutdown interface{}

    // MACSec Policy.
    PolicyNames Macsec_PolicyNames
}

func (macsec *Macsec) GetEntityData() *types.CommonEntityData {
    macsec.EntityData.YFilter = macsec.YFilter
    macsec.EntityData.YangName = "macsec"
    macsec.EntityData.BundleName = "cisco_ios_xr"
    macsec.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-macsec-mka-cfg"
    macsec.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-macsec-mka-cfg:macsec"
    macsec.EntityData.AbsolutePath = macsec.EntityData.SegmentPath
    macsec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsec.EntityData.Children = types.NewOrderedMap()
    macsec.EntityData.Children.Append("policy-names", types.YChild{"PolicyNames", &macsec.PolicyNames})
    macsec.EntityData.Leafs = types.NewOrderedMap()
    macsec.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", macsec.Shutdown})

    macsec.EntityData.YListKeys = []string {}

    return &(macsec.EntityData)
}

// Macsec_PolicyNames
// MACSec Policy
type Macsec_PolicyNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MACsec Policy Name. The type is slice of Macsec_PolicyNames_PolicyName.
    PolicyName []*Macsec_PolicyNames_PolicyName
}

func (policyNames *Macsec_PolicyNames) GetEntityData() *types.CommonEntityData {
    policyNames.EntityData.YFilter = policyNames.YFilter
    policyNames.EntityData.YangName = "policy-names"
    policyNames.EntityData.BundleName = "cisco_ios_xr"
    policyNames.EntityData.ParentYangName = "macsec"
    policyNames.EntityData.SegmentPath = "policy-names"
    policyNames.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-cfg:macsec/" + policyNames.EntityData.SegmentPath
    policyNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyNames.EntityData.Children = types.NewOrderedMap()
    policyNames.EntityData.Children.Append("policy-name", types.YChild{"PolicyName", nil})
    for i := range policyNames.PolicyName {
        policyNames.EntityData.Children.Append(types.GetSegmentPath(policyNames.PolicyName[i]), types.YChild{"PolicyName", policyNames.PolicyName[i]})
    }
    policyNames.EntityData.Leafs = types.NewOrderedMap()

    policyNames.EntityData.YListKeys = []string {}

    return &(policyNames.EntityData)
}

// Macsec_PolicyNames_PolicyName
// MACsec Policy Name
type Macsec_PolicyNames_PolicyName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the Policy of maximum length 16. The type
    // is string with length: 1..16.
    Name interface{}

    // Enables data delay protection. The type is interface{}.
    DelayProtection interface{}

    // Security-Policy of Policy. The type is MacsecMkaSecurityPolicy.
    SecurityPolicy interface{}

    // Key-Server-Priority of Policy. The type is interface{} with range: 0..255.
    KeyServerPriority interface{}

    // Conf-Offset of Policy. The type is MacsecMkaConfOffset.
    ConfOffset interface{}

    // DEPRECATED-Interval(in minutes) after which key-server generates new SAK
    // for a Secured Session, Default: OFF, recommended to use seconds option. The
    // type is interface{} with range: 1..43200. Units are minute.
    SakRekeyInterval interface{}

    // Macsec policy exception for packets to be in clear. The type is
    // MacsecMkaPolicyException.
    PolicyException interface{}

    // Window-Size of Policy. The type is interface{} with range: 0..1024.
    WindowSize interface{}

    // Cipher-suite of Policy. The type is MacsecMkaCipherSuite.
    CipherSuite interface{}

    // Enables Include ICV Indicator paramset in MKPDU. The type is interface{}.
    IncludeIcvIndicator interface{}

    // Interval(in seconds) after which key-server generates new SAK for a Secured
    // Session, Default: OFF. The type is interface{} with range: 60..2592000.
    // Units are second.
    SakRekeyIntervalSec interface{}

    // VLAN-Tags-In-Clear of Policy. The type is interface{} with range: 1..2.
    VlanTagsInClear interface{}
}

func (policyName *Macsec_PolicyNames_PolicyName) GetEntityData() *types.CommonEntityData {
    policyName.EntityData.YFilter = policyName.YFilter
    policyName.EntityData.YangName = "policy-name"
    policyName.EntityData.BundleName = "cisco_ios_xr"
    policyName.EntityData.ParentYangName = "policy-names"
    policyName.EntityData.SegmentPath = "policy-name" + types.AddKeyToken(policyName.Name, "name")
    policyName.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-cfg:macsec/policy-names/" + policyName.EntityData.SegmentPath
    policyName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyName.EntityData.Children = types.NewOrderedMap()
    policyName.EntityData.Leafs = types.NewOrderedMap()
    policyName.EntityData.Leafs.Append("name", types.YLeaf{"Name", policyName.Name})
    policyName.EntityData.Leafs.Append("delay-protection", types.YLeaf{"DelayProtection", policyName.DelayProtection})
    policyName.EntityData.Leafs.Append("security-policy", types.YLeaf{"SecurityPolicy", policyName.SecurityPolicy})
    policyName.EntityData.Leafs.Append("key-server-priority", types.YLeaf{"KeyServerPriority", policyName.KeyServerPriority})
    policyName.EntityData.Leafs.Append("conf-offset", types.YLeaf{"ConfOffset", policyName.ConfOffset})
    policyName.EntityData.Leafs.Append("sak-rekey-interval", types.YLeaf{"SakRekeyInterval", policyName.SakRekeyInterval})
    policyName.EntityData.Leafs.Append("policy-exception", types.YLeaf{"PolicyException", policyName.PolicyException})
    policyName.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", policyName.WindowSize})
    policyName.EntityData.Leafs.Append("cipher-suite", types.YLeaf{"CipherSuite", policyName.CipherSuite})
    policyName.EntityData.Leafs.Append("include-icv-indicator", types.YLeaf{"IncludeIcvIndicator", policyName.IncludeIcvIndicator})
    policyName.EntityData.Leafs.Append("sak-rekey-interval-sec", types.YLeaf{"SakRekeyIntervalSec", policyName.SakRekeyIntervalSec})
    policyName.EntityData.Leafs.Append("vlan-tags-in-clear", types.YLeaf{"VlanTagsInClear", policyName.VlanTagsInClear})

    policyName.EntityData.YListKeys = []string {"Name"}

    return &(policyName.EntityData)
}

