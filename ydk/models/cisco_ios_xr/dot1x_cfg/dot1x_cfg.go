// This module contains a collection of YANG definitions
// for Cisco IOS-XR dot1x package configuration.
// 
// This module contains definitions
// for the following management objects:
//   dot1x: Global Dot1x Configuration
//   eap: eap
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dot1x-cfg dot1x}", reflect.TypeOf(Dot1x{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dot1x-cfg:dot1x", reflect.TypeOf(Dot1x{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dot1x-cfg eap}", reflect.TypeOf(Eap{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dot1x-cfg:eap", reflect.TypeOf(Eap{}))
}

// Dot1xServerDeadAction represents Dot1x server dead action
type Dot1xServerDeadAction string

const (
    // server dead action auth-fail
    Dot1xServerDeadAction_auth_fail Dot1xServerDeadAction = "auth-fail"

    // server dead action auth-retry
    Dot1xServerDeadAction_auth_retry Dot1xServerDeadAction = "auth-retry"
)

// Dot1x
// Global Dot1x Configuration
type Dot1x struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global Dot1x Profile Name. The type is slice of Dot1x_Dot1xProfile.
    Dot1xProfile []*Dot1x_Dot1xProfile
}

func (dot1x *Dot1x) GetEntityData() *types.CommonEntityData {
    dot1x.EntityData.YFilter = dot1x.YFilter
    dot1x.EntityData.YangName = "dot1x"
    dot1x.EntityData.BundleName = "cisco_ios_xr"
    dot1x.EntityData.ParentYangName = "Cisco-IOS-XR-dot1x-cfg"
    dot1x.EntityData.SegmentPath = "Cisco-IOS-XR-dot1x-cfg:dot1x"
    dot1x.EntityData.AbsolutePath = dot1x.EntityData.SegmentPath
    dot1x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1x.EntityData.Children = types.NewOrderedMap()
    dot1x.EntityData.Children.Append("dot1x-profile", types.YChild{"Dot1xProfile", nil})
    for i := range dot1x.Dot1xProfile {
        dot1x.EntityData.Children.Append(types.GetSegmentPath(dot1x.Dot1xProfile[i]), types.YChild{"Dot1xProfile", dot1x.Dot1xProfile[i]})
    }
    dot1x.EntityData.Leafs = types.NewOrderedMap()

    dot1x.EntityData.YListKeys = []string {}

    return &(dot1x.EntityData)
}

// Dot1x_Dot1xProfile
// Global Dot1x Profile Name
type Dot1x_Dot1xProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the Dot1x Profile. The type is string with
    // length: 1..63.
    ProfileName interface{}

    // Dot1x PAE (Port Access Entity) Role. The type is string with pattern:
    // b'(supplicant)|(authenticator)|(both)'.
    Pae interface{}

    // Dot1x Supplicant Related Configuration.
    Supplicant Dot1x_Dot1xProfile_Supplicant

    // Dot1x Authenticator Related Configuration.
    Authenticator Dot1x_Dot1xProfile_Authenticator
}

func (dot1xProfile *Dot1x_Dot1xProfile) GetEntityData() *types.CommonEntityData {
    dot1xProfile.EntityData.YFilter = dot1xProfile.YFilter
    dot1xProfile.EntityData.YangName = "dot1x-profile"
    dot1xProfile.EntityData.BundleName = "cisco_ios_xr"
    dot1xProfile.EntityData.ParentYangName = "dot1x"
    dot1xProfile.EntityData.SegmentPath = "dot1x-profile" + types.AddKeyToken(dot1xProfile.ProfileName, "profile-name")
    dot1xProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-cfg:dot1x/" + dot1xProfile.EntityData.SegmentPath
    dot1xProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1xProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1xProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1xProfile.EntityData.Children = types.NewOrderedMap()
    dot1xProfile.EntityData.Children.Append("supplicant", types.YChild{"Supplicant", &dot1xProfile.Supplicant})
    dot1xProfile.EntityData.Children.Append("authenticator", types.YChild{"Authenticator", &dot1xProfile.Authenticator})
    dot1xProfile.EntityData.Leafs = types.NewOrderedMap()
    dot1xProfile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", dot1xProfile.ProfileName})
    dot1xProfile.EntityData.Leafs.Append("pae", types.YLeaf{"Pae", dot1xProfile.Pae})

    dot1xProfile.EntityData.YListKeys = []string {"ProfileName"}

    return &(dot1xProfile.EntityData)
}

// Dot1x_Dot1xProfile_Supplicant
// Dot1x Supplicant Related Configuration
type Dot1x_Dot1xProfile_Supplicant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EAP Profile for Supplicant. The type is string with length: 1..63.
    EapProfile interface{}
}

func (supplicant *Dot1x_Dot1xProfile_Supplicant) GetEntityData() *types.CommonEntityData {
    supplicant.EntityData.YFilter = supplicant.YFilter
    supplicant.EntityData.YangName = "supplicant"
    supplicant.EntityData.BundleName = "cisco_ios_xr"
    supplicant.EntityData.ParentYangName = "dot1x-profile"
    supplicant.EntityData.SegmentPath = "supplicant"
    supplicant.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-cfg:dot1x/dot1x-profile/" + supplicant.EntityData.SegmentPath
    supplicant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supplicant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supplicant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supplicant.EntityData.Children = types.NewOrderedMap()
    supplicant.EntityData.Leafs = types.NewOrderedMap()
    supplicant.EntityData.Leafs.Append("eap-profile", types.YLeaf{"EapProfile", supplicant.EapProfile})

    supplicant.EntityData.YListKeys = []string {}

    return &(supplicant.EntityData)
}

// Dot1x_Dot1xProfile_Authenticator
// Dot1x Authenticator Related Configuration
type Dot1x_Dot1xProfile_Authenticator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EAP Profile for Local EAP Server. The type is string with length: 1..63.
    EapProfile interface{}

    // dot1x authenticator action on AAA server unreachability. The type is
    // Dot1xServerDeadAction.
    ServerDead interface{}

    // Timers for Authenticator.
    Timers Dot1x_Dot1xProfile_Authenticator_Timers
}

func (authenticator *Dot1x_Dot1xProfile_Authenticator) GetEntityData() *types.CommonEntityData {
    authenticator.EntityData.YFilter = authenticator.YFilter
    authenticator.EntityData.YangName = "authenticator"
    authenticator.EntityData.BundleName = "cisco_ios_xr"
    authenticator.EntityData.ParentYangName = "dot1x-profile"
    authenticator.EntityData.SegmentPath = "authenticator"
    authenticator.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-cfg:dot1x/dot1x-profile/" + authenticator.EntityData.SegmentPath
    authenticator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticator.EntityData.Children = types.NewOrderedMap()
    authenticator.EntityData.Children.Append("timers", types.YChild{"Timers", &authenticator.Timers})
    authenticator.EntityData.Leafs = types.NewOrderedMap()
    authenticator.EntityData.Leafs.Append("eap-profile", types.YLeaf{"EapProfile", authenticator.EapProfile})
    authenticator.EntityData.Leafs.Append("server-dead", types.YLeaf{"ServerDead", authenticator.ServerDead})

    authenticator.EntityData.YListKeys = []string {}

    return &(authenticator.EntityData)
}

// Dot1x_Dot1xProfile_Authenticator_Timers
// Timers for Authenticator
type Dot1x_Dot1xProfile_Authenticator_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // After this time ReAuthentication will be trigerred.
    ReauthTime Dot1x_Dot1xProfile_Authenticator_Timers_ReauthTime
}

func (timers *Dot1x_Dot1xProfile_Authenticator_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "authenticator"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-cfg:dot1x/dot1x-profile/authenticator/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("reauth-time", types.YChild{"ReauthTime", &timers.ReauthTime})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Dot1x_Dot1xProfile_Authenticator_Timers_ReauthTime
// After this time ReAuthentication will be
// trigerred
type Dot1x_Dot1xProfile_Authenticator_Timers_ReauthTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reauth will be triggerred based on the EAP server configuration. The type
    // is bool.
    Server interface{}

    // Reauth will be triggerred based on the configuration in box. The type is
    // interface{} with range: 60..5184000. Units are second.
    Local interface{}
}

func (reauthTime *Dot1x_Dot1xProfile_Authenticator_Timers_ReauthTime) GetEntityData() *types.CommonEntityData {
    reauthTime.EntityData.YFilter = reauthTime.YFilter
    reauthTime.EntityData.YangName = "reauth-time"
    reauthTime.EntityData.BundleName = "cisco_ios_xr"
    reauthTime.EntityData.ParentYangName = "timers"
    reauthTime.EntityData.SegmentPath = "reauth-time"
    reauthTime.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-cfg:dot1x/dot1x-profile/authenticator/timers/" + reauthTime.EntityData.SegmentPath
    reauthTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reauthTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reauthTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reauthTime.EntityData.Children = types.NewOrderedMap()
    reauthTime.EntityData.Leafs = types.NewOrderedMap()
    reauthTime.EntityData.Leafs.Append("server", types.YLeaf{"Server", reauthTime.Server})
    reauthTime.EntityData.Leafs.Append("local", types.YLeaf{"Local", reauthTime.Local})

    reauthTime.EntityData.YListKeys = []string {}

    return &(reauthTime.EntityData)
}

// Eap
// eap
type Eap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global EAP Profile Configuration. The type is slice of Eap_EapProfile.
    EapProfile []*Eap_EapProfile
}

func (eap *Eap) GetEntityData() *types.CommonEntityData {
    eap.EntityData.YFilter = eap.YFilter
    eap.EntityData.YangName = "eap"
    eap.EntityData.BundleName = "cisco_ios_xr"
    eap.EntityData.ParentYangName = "Cisco-IOS-XR-dot1x-cfg"
    eap.EntityData.SegmentPath = "Cisco-IOS-XR-dot1x-cfg:eap"
    eap.EntityData.AbsolutePath = eap.EntityData.SegmentPath
    eap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eap.EntityData.Children = types.NewOrderedMap()
    eap.EntityData.Children.Append("eap-profile", types.YChild{"EapProfile", nil})
    for i := range eap.EapProfile {
        eap.EntityData.Children.Append(types.GetSegmentPath(eap.EapProfile[i]), types.YChild{"EapProfile", eap.EapProfile[i]})
    }
    eap.EntityData.Leafs = types.NewOrderedMap()

    eap.EntityData.YListKeys = []string {}

    return &(eap.EntityData)
}

// Eap_EapProfile
// Global EAP Profile Configuration
type Eap_EapProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the EAP Profile. The type is string with
    // length: 1..63.
    ProfileName interface{}

    // Configure backward compatibility for TLS 1.0. The type is interface{}.
    AllowEapTls10 interface{}

    // Configure EAP Identity/UserName. The type is string with length: 1..63.
    Identity interface{}

    // EAP TLS Configuration.
    Eaptls Eap_EapProfile_Eaptls
}

func (eapProfile *Eap_EapProfile) GetEntityData() *types.CommonEntityData {
    eapProfile.EntityData.YFilter = eapProfile.YFilter
    eapProfile.EntityData.YangName = "eap-profile"
    eapProfile.EntityData.BundleName = "cisco_ios_xr"
    eapProfile.EntityData.ParentYangName = "eap"
    eapProfile.EntityData.SegmentPath = "eap-profile" + types.AddKeyToken(eapProfile.ProfileName, "profile-name")
    eapProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-cfg:eap/" + eapProfile.EntityData.SegmentPath
    eapProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eapProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eapProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eapProfile.EntityData.Children = types.NewOrderedMap()
    eapProfile.EntityData.Children.Append("eaptls", types.YChild{"Eaptls", &eapProfile.Eaptls})
    eapProfile.EntityData.Leafs = types.NewOrderedMap()
    eapProfile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", eapProfile.ProfileName})
    eapProfile.EntityData.Leafs.Append("allow-eap-tls1-0", types.YLeaf{"AllowEapTls10", eapProfile.AllowEapTls10})
    eapProfile.EntityData.Leafs.Append("identity", types.YLeaf{"Identity", eapProfile.Identity})

    eapProfile.EntityData.YListKeys = []string {"ProfileName"}

    return &(eapProfile.EntityData)
}

// Eap_EapProfile_Eaptls
// EAP TLS Configuration
type Eap_EapProfile_Eaptls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure PKI Trustpoint. The type is string with length: 1..63.
    PkiTrustpoint interface{}
}

func (eaptls *Eap_EapProfile_Eaptls) GetEntityData() *types.CommonEntityData {
    eaptls.EntityData.YFilter = eaptls.YFilter
    eaptls.EntityData.YangName = "eaptls"
    eaptls.EntityData.BundleName = "cisco_ios_xr"
    eaptls.EntityData.ParentYangName = "eap-profile"
    eaptls.EntityData.SegmentPath = "eaptls"
    eaptls.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-cfg:eap/eap-profile/" + eaptls.EntityData.SegmentPath
    eaptls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eaptls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eaptls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eaptls.EntityData.Children = types.NewOrderedMap()
    eaptls.EntityData.Leafs = types.NewOrderedMap()
    eaptls.EntityData.Leafs.Append("pki-trustpoint", types.YLeaf{"PkiTrustpoint", eaptls.PkiTrustpoint})

    eaptls.EntityData.YListKeys = []string {}

    return &(eaptls.EntityData)
}

