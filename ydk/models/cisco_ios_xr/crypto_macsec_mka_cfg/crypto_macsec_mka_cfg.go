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
    parent types.Entity
    YFilter yfilter.YFilter

    // MACSec Policy. The type is slice of Macsec_Policy.
    Policy []Macsec_Policy
}

func (macsec *Macsec) GetFilter() yfilter.YFilter { return macsec.YFilter }

func (macsec *Macsec) SetFilter(yf yfilter.YFilter) { macsec.YFilter = yf }

func (macsec *Macsec) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    return ""
}

func (macsec *Macsec) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-macsec-mka-cfg:macsec"
}

func (macsec *Macsec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy" {
        for _, c := range macsec.Policy {
            if macsec.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Policy{}
        macsec.Policy = append(macsec.Policy, child)
        return &macsec.Policy[len(macsec.Policy)-1]
    }
    return nil
}

func (macsec *Macsec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macsec.Policy {
        children[macsec.Policy[i].GetSegmentPath()] = &macsec.Policy[i]
    }
    return children
}

func (macsec *Macsec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macsec *Macsec) GetBundleName() string { return "cisco_ios_xr" }

func (macsec *Macsec) GetYangName() string { return "macsec" }

func (macsec *Macsec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsec *Macsec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsec *Macsec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsec *Macsec) SetParent(parent types.Entity) { macsec.parent = parent }

func (macsec *Macsec) GetParent() types.Entity { return macsec.parent }

func (macsec *Macsec) GetParentYangName() string { return "Cisco-IOS-XR-crypto-macsec-mka-cfg" }

// Macsec_Policy
// MACSec Policy
type Macsec_Policy struct {
    parent types.Entity
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

func (policy *Macsec_Policy) GetFilter() yfilter.YFilter { return policy.YFilter }

func (policy *Macsec_Policy) SetFilter(yf yfilter.YFilter) { policy.YFilter = yf }

func (policy *Macsec_Policy) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "delay-protection" { return "DelayProtection" }
    if yname == "security-policy" { return "SecurityPolicy" }
    if yname == "key-server-priority" { return "KeyServerPriority" }
    if yname == "conf-offset" { return "ConfOffset" }
    if yname == "sak-rekey-interval" { return "SakRekeyInterval" }
    if yname == "policy-exception" { return "PolicyException" }
    if yname == "window-size" { return "WindowSize" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "include-icv-indicator" { return "IncludeIcvIndicator" }
    if yname == "vlan-tags-in-clear" { return "VlanTagsInClear" }
    return ""
}

func (policy *Macsec_Policy) GetSegmentPath() string {
    return "policy" + "[name='" + fmt.Sprintf("%v", policy.Name) + "']"
}

func (policy *Macsec_Policy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policy *Macsec_Policy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policy *Macsec_Policy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = policy.Name
    leafs["delay-protection"] = policy.DelayProtection
    leafs["security-policy"] = policy.SecurityPolicy
    leafs["key-server-priority"] = policy.KeyServerPriority
    leafs["conf-offset"] = policy.ConfOffset
    leafs["sak-rekey-interval"] = policy.SakRekeyInterval
    leafs["policy-exception"] = policy.PolicyException
    leafs["window-size"] = policy.WindowSize
    leafs["cipher-suite"] = policy.CipherSuite
    leafs["include-icv-indicator"] = policy.IncludeIcvIndicator
    leafs["vlan-tags-in-clear"] = policy.VlanTagsInClear
    return leafs
}

func (policy *Macsec_Policy) GetBundleName() string { return "cisco_ios_xr" }

func (policy *Macsec_Policy) GetYangName() string { return "policy" }

func (policy *Macsec_Policy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policy *Macsec_Policy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policy *Macsec_Policy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policy *Macsec_Policy) SetParent(parent types.Entity) { policy.parent = parent }

func (policy *Macsec_Policy) GetParent() types.Entity { return policy.parent }

func (policy *Macsec_Policy) GetParentYangName() string { return "macsec" }

