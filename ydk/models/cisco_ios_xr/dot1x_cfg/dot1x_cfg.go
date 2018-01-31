// This module contains a collection of YANG definitions
// for Cisco IOS-XR dot1x package configuration.
// 
// This module contains definitions
// for the following management objects:
//   dot1x: Global Dot1x Configuration
//   eap: eap
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dot1x_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dot1x_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dot1x-cfg dot1x}", reflect.TypeOf(Dot1X{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dot1x-cfg:dot1x", reflect.TypeOf(Dot1X{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dot1x-cfg eap}", reflect.TypeOf(Eap{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dot1x-cfg:eap", reflect.TypeOf(Eap{}))
}

// Dot1X
// Global Dot1x Configuration
type Dot1X struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global Dot1x Profile Name. The type is slice of Dot1X_Dot1XProfile.
    Dot1XProfile []Dot1X_Dot1XProfile
}

func (dot1X *Dot1X) GetFilter() yfilter.YFilter { return dot1X.YFilter }

func (dot1X *Dot1X) SetFilter(yf yfilter.YFilter) { dot1X.YFilter = yf }

func (dot1X *Dot1X) GetGoName(yname string) string {
    if yname == "dot1x-profile" { return "Dot1XProfile" }
    return ""
}

func (dot1X *Dot1X) GetSegmentPath() string {
    return "Cisco-IOS-XR-dot1x-cfg:dot1x"
}

func (dot1X *Dot1X) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1x-profile" {
        for _, c := range dot1X.Dot1XProfile {
            if dot1X.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dot1X_Dot1XProfile{}
        dot1X.Dot1XProfile = append(dot1X.Dot1XProfile, child)
        return &dot1X.Dot1XProfile[len(dot1X.Dot1XProfile)-1]
    }
    return nil
}

func (dot1X *Dot1X) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1X.Dot1XProfile {
        children[dot1X.Dot1XProfile[i].GetSegmentPath()] = &dot1X.Dot1XProfile[i]
    }
    return children
}

func (dot1X *Dot1X) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1X *Dot1X) GetBundleName() string { return "cisco_ios_xr" }

func (dot1X *Dot1X) GetYangName() string { return "dot1x" }

func (dot1X *Dot1X) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1X *Dot1X) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1X *Dot1X) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1X *Dot1X) SetParent(parent types.Entity) { dot1X.parent = parent }

func (dot1X *Dot1X) GetParent() types.Entity { return dot1X.parent }

func (dot1X *Dot1X) GetParentYangName() string { return "Cisco-IOS-XR-dot1x-cfg" }

// Dot1X_Dot1XProfile
// Global Dot1x Profile Name
type Dot1X_Dot1XProfile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Dot1x Profile. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Dot1x PAE (Port Access Entity) Role. The type is string with pattern:
    // (supplicant)|(authenticator)|(both).
    Pae interface{}

    // Dot1x Supplicant Related Configuration.
    Supplicant Dot1X_Dot1XProfile_Supplicant

    // Dot1x Authenticator Related Configuration.
    Authenticator Dot1X_Dot1XProfile_Authenticator
}

func (dot1XProfile *Dot1X_Dot1XProfile) GetFilter() yfilter.YFilter { return dot1XProfile.YFilter }

func (dot1XProfile *Dot1X_Dot1XProfile) SetFilter(yf yfilter.YFilter) { dot1XProfile.YFilter = yf }

func (dot1XProfile *Dot1X_Dot1XProfile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "pae" { return "Pae" }
    if yname == "supplicant" { return "Supplicant" }
    if yname == "authenticator" { return "Authenticator" }
    return ""
}

func (dot1XProfile *Dot1X_Dot1XProfile) GetSegmentPath() string {
    return "dot1x-profile" + "[profile-name='" + fmt.Sprintf("%v", dot1XProfile.ProfileName) + "']"
}

func (dot1XProfile *Dot1X_Dot1XProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "supplicant" {
        return &dot1XProfile.Supplicant
    }
    if childYangName == "authenticator" {
        return &dot1XProfile.Authenticator
    }
    return nil
}

func (dot1XProfile *Dot1X_Dot1XProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["supplicant"] = &dot1XProfile.Supplicant
    children["authenticator"] = &dot1XProfile.Authenticator
    return children
}

func (dot1XProfile *Dot1X_Dot1XProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = dot1XProfile.ProfileName
    leafs["pae"] = dot1XProfile.Pae
    return leafs
}

func (dot1XProfile *Dot1X_Dot1XProfile) GetBundleName() string { return "cisco_ios_xr" }

func (dot1XProfile *Dot1X_Dot1XProfile) GetYangName() string { return "dot1x-profile" }

func (dot1XProfile *Dot1X_Dot1XProfile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1XProfile *Dot1X_Dot1XProfile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1XProfile *Dot1X_Dot1XProfile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1XProfile *Dot1X_Dot1XProfile) SetParent(parent types.Entity) { dot1XProfile.parent = parent }

func (dot1XProfile *Dot1X_Dot1XProfile) GetParent() types.Entity { return dot1XProfile.parent }

func (dot1XProfile *Dot1X_Dot1XProfile) GetParentYangName() string { return "dot1x" }

// Dot1X_Dot1XProfile_Supplicant
// Dot1x Supplicant Related Configuration
type Dot1X_Dot1XProfile_Supplicant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EAP Profile for Supplicant. The type is string.
    EapProfile interface{}
}

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetFilter() yfilter.YFilter { return supplicant.YFilter }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) SetFilter(yf yfilter.YFilter) { supplicant.YFilter = yf }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetGoName(yname string) string {
    if yname == "eap-profile" { return "EapProfile" }
    return ""
}

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetSegmentPath() string {
    return "supplicant"
}

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["eap-profile"] = supplicant.EapProfile
    return leafs
}

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetBundleName() string { return "cisco_ios_xr" }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetYangName() string { return "supplicant" }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) SetParent(parent types.Entity) { supplicant.parent = parent }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetParent() types.Entity { return supplicant.parent }

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetParentYangName() string { return "dot1x-profile" }

// Dot1X_Dot1XProfile_Authenticator
// Dot1x Authenticator Related Configuration
type Dot1X_Dot1XProfile_Authenticator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timers for Authenticator.
    Timers Dot1X_Dot1XProfile_Authenticator_Timers
}

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetFilter() yfilter.YFilter { return authenticator.YFilter }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) SetFilter(yf yfilter.YFilter) { authenticator.YFilter = yf }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetGoName(yname string) string {
    if yname == "timers" { return "Timers" }
    return ""
}

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetSegmentPath() string {
    return "authenticator"
}

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timers" {
        return &authenticator.Timers
    }
    return nil
}

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["timers"] = &authenticator.Timers
    return children
}

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetBundleName() string { return "cisco_ios_xr" }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetYangName() string { return "authenticator" }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) SetParent(parent types.Entity) { authenticator.parent = parent }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetParent() types.Entity { return authenticator.parent }

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetParentYangName() string { return "dot1x-profile" }

// Dot1X_Dot1XProfile_Authenticator_Timers
// Timers for Authenticator
type Dot1X_Dot1XProfile_Authenticator_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // After this time ReAuthentication will be trigerred.
    ReauthTime Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime
}

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetGoName(yname string) string {
    if yname == "reauth-time" { return "ReauthTime" }
    return ""
}

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reauth-time" {
        return &timers.ReauthTime
    }
    return nil
}

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["reauth-time"] = &timers.ReauthTime
    return children
}

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetBundleName() string { return "cisco_ios_xr" }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetYangName() string { return "timers" }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetParentYangName() string { return "authenticator" }

// Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime
// After this time ReAuthentication will be
// trigerred
type Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reauth will be triggerred based on the EAP server configuration. The type
    // is bool.
    Server interface{}

    // Reauth will be triggerred based on the configuration in box. The type is
    // interface{} with range: 60..5184000. Units are second.
    Local interface{}
}

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetFilter() yfilter.YFilter { return reauthTime.YFilter }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) SetFilter(yf yfilter.YFilter) { reauthTime.YFilter = yf }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    if yname == "local" { return "Local" }
    return ""
}

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetSegmentPath() string {
    return "reauth-time"
}

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server"] = reauthTime.Server
    leafs["local"] = reauthTime.Local
    return leafs
}

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetBundleName() string { return "cisco_ios_xr" }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetYangName() string { return "reauth-time" }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) SetParent(parent types.Entity) { reauthTime.parent = parent }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetParent() types.Entity { return reauthTime.parent }

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetParentYangName() string { return "timers" }

// Eap
// eap
type Eap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global EAP Profile Configuration. The type is slice of Eap_EapProfile.
    EapProfile []Eap_EapProfile
}

func (eap *Eap) GetFilter() yfilter.YFilter { return eap.YFilter }

func (eap *Eap) SetFilter(yf yfilter.YFilter) { eap.YFilter = yf }

func (eap *Eap) GetGoName(yname string) string {
    if yname == "eap-profile" { return "EapProfile" }
    return ""
}

func (eap *Eap) GetSegmentPath() string {
    return "Cisco-IOS-XR-dot1x-cfg:eap"
}

func (eap *Eap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "eap-profile" {
        for _, c := range eap.EapProfile {
            if eap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Eap_EapProfile{}
        eap.EapProfile = append(eap.EapProfile, child)
        return &eap.EapProfile[len(eap.EapProfile)-1]
    }
    return nil
}

func (eap *Eap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eap.EapProfile {
        children[eap.EapProfile[i].GetSegmentPath()] = &eap.EapProfile[i]
    }
    return children
}

func (eap *Eap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eap *Eap) GetBundleName() string { return "cisco_ios_xr" }

func (eap *Eap) GetYangName() string { return "eap" }

func (eap *Eap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eap *Eap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eap *Eap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eap *Eap) SetParent(parent types.Entity) { eap.parent = parent }

func (eap *Eap) GetParent() types.Entity { return eap.parent }

func (eap *Eap) GetParentYangName() string { return "Cisco-IOS-XR-dot1x-cfg" }

// Eap_EapProfile
// Global EAP Profile Configuration
type Eap_EapProfile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the EAP Profile. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Configure EAP Identity/UserName. The type is string.
    Identity interface{}

    // EAP TLS Configuration.
    Eaptls Eap_EapProfile_Eaptls
}

func (eapProfile *Eap_EapProfile) GetFilter() yfilter.YFilter { return eapProfile.YFilter }

func (eapProfile *Eap_EapProfile) SetFilter(yf yfilter.YFilter) { eapProfile.YFilter = yf }

func (eapProfile *Eap_EapProfile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "identity" { return "Identity" }
    if yname == "eaptls" { return "Eaptls" }
    return ""
}

func (eapProfile *Eap_EapProfile) GetSegmentPath() string {
    return "eap-profile" + "[profile-name='" + fmt.Sprintf("%v", eapProfile.ProfileName) + "']"
}

func (eapProfile *Eap_EapProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "eaptls" {
        return &eapProfile.Eaptls
    }
    return nil
}

func (eapProfile *Eap_EapProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["eaptls"] = &eapProfile.Eaptls
    return children
}

func (eapProfile *Eap_EapProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = eapProfile.ProfileName
    leafs["identity"] = eapProfile.Identity
    return leafs
}

func (eapProfile *Eap_EapProfile) GetBundleName() string { return "cisco_ios_xr" }

func (eapProfile *Eap_EapProfile) GetYangName() string { return "eap-profile" }

func (eapProfile *Eap_EapProfile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eapProfile *Eap_EapProfile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eapProfile *Eap_EapProfile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eapProfile *Eap_EapProfile) SetParent(parent types.Entity) { eapProfile.parent = parent }

func (eapProfile *Eap_EapProfile) GetParent() types.Entity { return eapProfile.parent }

func (eapProfile *Eap_EapProfile) GetParentYangName() string { return "eap" }

// Eap_EapProfile_Eaptls
// EAP TLS Configuration
type Eap_EapProfile_Eaptls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure PKI Trustpoint. The type is string.
    PkiTrustpoint interface{}
}

func (eaptls *Eap_EapProfile_Eaptls) GetFilter() yfilter.YFilter { return eaptls.YFilter }

func (eaptls *Eap_EapProfile_Eaptls) SetFilter(yf yfilter.YFilter) { eaptls.YFilter = yf }

func (eaptls *Eap_EapProfile_Eaptls) GetGoName(yname string) string {
    if yname == "pki-trustpoint" { return "PkiTrustpoint" }
    return ""
}

func (eaptls *Eap_EapProfile_Eaptls) GetSegmentPath() string {
    return "eaptls"
}

func (eaptls *Eap_EapProfile_Eaptls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eaptls *Eap_EapProfile_Eaptls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eaptls *Eap_EapProfile_Eaptls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pki-trustpoint"] = eaptls.PkiTrustpoint
    return leafs
}

func (eaptls *Eap_EapProfile_Eaptls) GetBundleName() string { return "cisco_ios_xr" }

func (eaptls *Eap_EapProfile_Eaptls) GetYangName() string { return "eaptls" }

func (eaptls *Eap_EapProfile_Eaptls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eaptls *Eap_EapProfile_Eaptls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eaptls *Eap_EapProfile_Eaptls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eaptls *Eap_EapProfile_Eaptls) SetParent(parent types.Entity) { eaptls.parent = parent }

func (eaptls *Eap_EapProfile_Eaptls) GetParent() types.Entity { return eaptls.parent }

func (eaptls *Eap_EapProfile_Eaptls) GetParentYangName() string { return "eap-profile" }

