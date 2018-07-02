// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-mka package configuration.
// 
// This module contains definitions
// for the following management objects:
//   macsec: MACSec MKA
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

    // MACSec Policy. The type is slice of Macsec_Policy.
    Policy []*Macsec_Policy
}

func (macsec *Macsec) GetEntityData() *types.CommonEntityData {
    macsec.EntityData.YFilter = macsec.YFilter
    macsec.EntityData.YangName = "macsec"
    macsec.EntityData.BundleName = "cisco_ios_xr"
    macsec.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-macsec-mka-cfg"
    macsec.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-macsec-mka-cfg:macsec"
    macsec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsec.EntityData.Children = types.NewOrderedMap()
    macsec.EntityData.Children.Append("policy", types.YChild{"Policy", nil})
    for i := range macsec.Policy {
        macsec.EntityData.Children.Append(types.GetSegmentPath(macsec.Policy[i]), types.YChild{"Policy", macsec.Policy[i]})
    }
    macsec.EntityData.Leafs = types.NewOrderedMap()

    macsec.EntityData.YListKeys = []string {}

    return &(macsec.EntityData)
}

// Macsec_Policy
// MACSec Policy
type Macsec_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Policy of maximum length 16. The type
    // is string with length: 1..16.
    Name interface{}

    // TRUE enables data delay protection. The type is bool.
    DelayProtection interface{}

    // Security-Policy of Policy. The type is MacsecMkaSecurityPolicy.
    SecurityPolicy interface{}

    // Key-Server-Priority of Policy. The type is interface{} with range: 0..255.
    KeyServerPriority interface{}

    // Conf-Offset of Policy. The type is MacsecMkaConfOffset.
    ConfOffset interface{}

    // Interval after which key-server generates new SAK for a Secured Session.
    // The type is interface{} with range: 0..43200. Units are minute.
    SakRekeyInterval interface{}

    // Macsec policy exception for packets to be in clear. The type is
    // MacsecMkaPolicyException.
    PolicyException interface{}

    // Window-Size of Policy. The type is interface{} with range: 0..1024.
    WindowSize interface{}

    // Cipher-suite of Policy. The type is MacsecMkaCipherSuite.
    CipherSuite interface{}

    // TRUE enables Include ICV Indicator paramset in MKPDU. The type is bool.
    IncludeIcvIndicator interface{}

    // VLAN-Tags-In-Clear of Policy. The type is interface{} with range: 1..2.
    VlanTagsInClear interface{}
}

func (policy *Macsec_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "macsec"
    policy.EntityData.SegmentPath = "policy" + types.AddKeyToken(policy.Name, "name")
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = types.NewOrderedMap()
    policy.EntityData.Leafs = types.NewOrderedMap()
    policy.EntityData.Leafs.Append("name", types.YLeaf{"Name", policy.Name})
    policy.EntityData.Leafs.Append("delay-protection", types.YLeaf{"DelayProtection", policy.DelayProtection})
    policy.EntityData.Leafs.Append("security-policy", types.YLeaf{"SecurityPolicy", policy.SecurityPolicy})
    policy.EntityData.Leafs.Append("key-server-priority", types.YLeaf{"KeyServerPriority", policy.KeyServerPriority})
    policy.EntityData.Leafs.Append("conf-offset", types.YLeaf{"ConfOffset", policy.ConfOffset})
    policy.EntityData.Leafs.Append("sak-rekey-interval", types.YLeaf{"SakRekeyInterval", policy.SakRekeyInterval})
    policy.EntityData.Leafs.Append("policy-exception", types.YLeaf{"PolicyException", policy.PolicyException})
    policy.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", policy.WindowSize})
    policy.EntityData.Leafs.Append("cipher-suite", types.YLeaf{"CipherSuite", policy.CipherSuite})
    policy.EntityData.Leafs.Append("include-icv-indicator", types.YLeaf{"IncludeIcvIndicator", policy.IncludeIcvIndicator})
    policy.EntityData.Leafs.Append("vlan-tags-in-clear", types.YLeaf{"VlanTagsInClear", policy.VlanTagsInClear})

    policy.EntityData.YListKeys = []string {"Name"}

    return &(policy.EntityData)
}

