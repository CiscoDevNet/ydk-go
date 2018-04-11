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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global Dot1x Profile Name. The type is slice of Dot1X_Dot1XProfile.
    Dot1XProfile []Dot1X_Dot1XProfile
}

func (dot1X *Dot1X) GetEntityData() *types.CommonEntityData {
    dot1X.EntityData.YFilter = dot1X.YFilter
    dot1X.EntityData.YangName = "dot1x"
    dot1X.EntityData.BundleName = "cisco_ios_xr"
    dot1X.EntityData.ParentYangName = "Cisco-IOS-XR-dot1x-cfg"
    dot1X.EntityData.SegmentPath = "Cisco-IOS-XR-dot1x-cfg:dot1x"
    dot1X.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1X.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1X.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1X.EntityData.Children = make(map[string]types.YChild)
    dot1X.EntityData.Children["dot1x-profile"] = types.YChild{"Dot1XProfile", nil}
    for i := range dot1X.Dot1XProfile {
        dot1X.EntityData.Children[types.GetSegmentPath(&dot1X.Dot1XProfile[i])] = types.YChild{"Dot1XProfile", &dot1X.Dot1XProfile[i]}
    }
    dot1X.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1X.EntityData)
}

// Dot1X_Dot1XProfile
// Global Dot1x Profile Name
type Dot1X_Dot1XProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Dot1x Profile. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProfileName interface{}

    // Dot1x PAE (Port Access Entity) Role. The type is string with pattern:
    // b'(supplicant)|(authenticator)|(both)'.
    Pae interface{}

    // Dot1x Supplicant Related Configuration.
    Supplicant Dot1X_Dot1XProfile_Supplicant

    // Dot1x Authenticator Related Configuration.
    Authenticator Dot1X_Dot1XProfile_Authenticator
}

func (dot1XProfile *Dot1X_Dot1XProfile) GetEntityData() *types.CommonEntityData {
    dot1XProfile.EntityData.YFilter = dot1XProfile.YFilter
    dot1XProfile.EntityData.YangName = "dot1x-profile"
    dot1XProfile.EntityData.BundleName = "cisco_ios_xr"
    dot1XProfile.EntityData.ParentYangName = "dot1x"
    dot1XProfile.EntityData.SegmentPath = "dot1x-profile" + "[profile-name='" + fmt.Sprintf("%v", dot1XProfile.ProfileName) + "']"
    dot1XProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1XProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1XProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1XProfile.EntityData.Children = make(map[string]types.YChild)
    dot1XProfile.EntityData.Children["supplicant"] = types.YChild{"Supplicant", &dot1XProfile.Supplicant}
    dot1XProfile.EntityData.Children["authenticator"] = types.YChild{"Authenticator", &dot1XProfile.Authenticator}
    dot1XProfile.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1XProfile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", dot1XProfile.ProfileName}
    dot1XProfile.EntityData.Leafs["pae"] = types.YLeaf{"Pae", dot1XProfile.Pae}
    return &(dot1XProfile.EntityData)
}

// Dot1X_Dot1XProfile_Supplicant
// Dot1x Supplicant Related Configuration
type Dot1X_Dot1XProfile_Supplicant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EAP Profile for Supplicant. The type is string.
    EapProfile interface{}
}

func (supplicant *Dot1X_Dot1XProfile_Supplicant) GetEntityData() *types.CommonEntityData {
    supplicant.EntityData.YFilter = supplicant.YFilter
    supplicant.EntityData.YangName = "supplicant"
    supplicant.EntityData.BundleName = "cisco_ios_xr"
    supplicant.EntityData.ParentYangName = "dot1x-profile"
    supplicant.EntityData.SegmentPath = "supplicant"
    supplicant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supplicant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supplicant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supplicant.EntityData.Children = make(map[string]types.YChild)
    supplicant.EntityData.Leafs = make(map[string]types.YLeaf)
    supplicant.EntityData.Leafs["eap-profile"] = types.YLeaf{"EapProfile", supplicant.EapProfile}
    return &(supplicant.EntityData)
}

// Dot1X_Dot1XProfile_Authenticator
// Dot1x Authenticator Related Configuration
type Dot1X_Dot1XProfile_Authenticator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timers for Authenticator.
    Timers Dot1X_Dot1XProfile_Authenticator_Timers
}

func (authenticator *Dot1X_Dot1XProfile_Authenticator) GetEntityData() *types.CommonEntityData {
    authenticator.EntityData.YFilter = authenticator.YFilter
    authenticator.EntityData.YangName = "authenticator"
    authenticator.EntityData.BundleName = "cisco_ios_xr"
    authenticator.EntityData.ParentYangName = "dot1x-profile"
    authenticator.EntityData.SegmentPath = "authenticator"
    authenticator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticator.EntityData.Children = make(map[string]types.YChild)
    authenticator.EntityData.Children["timers"] = types.YChild{"Timers", &authenticator.Timers}
    authenticator.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authenticator.EntityData)
}

// Dot1X_Dot1XProfile_Authenticator_Timers
// Timers for Authenticator
type Dot1X_Dot1XProfile_Authenticator_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // After this time ReAuthentication will be trigerred.
    ReauthTime Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime
}

func (timers *Dot1X_Dot1XProfile_Authenticator_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "authenticator"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Children["reauth-time"] = types.YChild{"ReauthTime", &timers.ReauthTime}
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timers.EntityData)
}

// Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime
// After this time ReAuthentication will be
// trigerred
type Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reauth will be triggerred based on the EAP server configuration. The type
    // is bool.
    Server interface{}

    // Reauth will be triggerred based on the configuration in box. The type is
    // interface{} with range: 60..5184000. Units are second.
    Local interface{}
}

func (reauthTime *Dot1X_Dot1XProfile_Authenticator_Timers_ReauthTime) GetEntityData() *types.CommonEntityData {
    reauthTime.EntityData.YFilter = reauthTime.YFilter
    reauthTime.EntityData.YangName = "reauth-time"
    reauthTime.EntityData.BundleName = "cisco_ios_xr"
    reauthTime.EntityData.ParentYangName = "timers"
    reauthTime.EntityData.SegmentPath = "reauth-time"
    reauthTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reauthTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reauthTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reauthTime.EntityData.Children = make(map[string]types.YChild)
    reauthTime.EntityData.Leafs = make(map[string]types.YLeaf)
    reauthTime.EntityData.Leafs["server"] = types.YLeaf{"Server", reauthTime.Server}
    reauthTime.EntityData.Leafs["local"] = types.YLeaf{"Local", reauthTime.Local}
    return &(reauthTime.EntityData)
}

// Eap
// eap
type Eap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global EAP Profile Configuration. The type is slice of Eap_EapProfile.
    EapProfile []Eap_EapProfile
}

func (eap *Eap) GetEntityData() *types.CommonEntityData {
    eap.EntityData.YFilter = eap.YFilter
    eap.EntityData.YangName = "eap"
    eap.EntityData.BundleName = "cisco_ios_xr"
    eap.EntityData.ParentYangName = "Cisco-IOS-XR-dot1x-cfg"
    eap.EntityData.SegmentPath = "Cisco-IOS-XR-dot1x-cfg:eap"
    eap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eap.EntityData.Children = make(map[string]types.YChild)
    eap.EntityData.Children["eap-profile"] = types.YChild{"EapProfile", nil}
    for i := range eap.EapProfile {
        eap.EntityData.Children[types.GetSegmentPath(&eap.EapProfile[i])] = types.YChild{"EapProfile", &eap.EapProfile[i]}
    }
    eap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eap.EntityData)
}

// Eap_EapProfile
// Global EAP Profile Configuration
type Eap_EapProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the EAP Profile. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProfileName interface{}

    // Configure EAP Identity/UserName. The type is string.
    Identity interface{}

    // EAP TLS Configuration.
    Eaptls Eap_EapProfile_Eaptls
}

func (eapProfile *Eap_EapProfile) GetEntityData() *types.CommonEntityData {
    eapProfile.EntityData.YFilter = eapProfile.YFilter
    eapProfile.EntityData.YangName = "eap-profile"
    eapProfile.EntityData.BundleName = "cisco_ios_xr"
    eapProfile.EntityData.ParentYangName = "eap"
    eapProfile.EntityData.SegmentPath = "eap-profile" + "[profile-name='" + fmt.Sprintf("%v", eapProfile.ProfileName) + "']"
    eapProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eapProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eapProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eapProfile.EntityData.Children = make(map[string]types.YChild)
    eapProfile.EntityData.Children["eaptls"] = types.YChild{"Eaptls", &eapProfile.Eaptls}
    eapProfile.EntityData.Leafs = make(map[string]types.YLeaf)
    eapProfile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", eapProfile.ProfileName}
    eapProfile.EntityData.Leafs["identity"] = types.YLeaf{"Identity", eapProfile.Identity}
    return &(eapProfile.EntityData)
}

// Eap_EapProfile_Eaptls
// EAP TLS Configuration
type Eap_EapProfile_Eaptls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure PKI Trustpoint. The type is string.
    PkiTrustpoint interface{}
}

func (eaptls *Eap_EapProfile_Eaptls) GetEntityData() *types.CommonEntityData {
    eaptls.EntityData.YFilter = eaptls.YFilter
    eaptls.EntityData.YangName = "eaptls"
    eaptls.EntityData.BundleName = "cisco_ios_xr"
    eaptls.EntityData.ParentYangName = "eap-profile"
    eaptls.EntityData.SegmentPath = "eaptls"
    eaptls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eaptls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eaptls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eaptls.EntityData.Children = make(map[string]types.YChild)
    eaptls.EntityData.Leafs = make(map[string]types.YLeaf)
    eaptls.EntityData.Leafs["pki-trustpoint"] = types.YLeaf{"PkiTrustpoint", eaptls.PkiTrustpoint}
    return &(eaptls.EntityData)
}

