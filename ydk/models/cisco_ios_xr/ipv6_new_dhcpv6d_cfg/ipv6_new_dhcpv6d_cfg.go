// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-new-dhcpv6d package configuration.
// 
// This module contains definitions
// for the following management objects:
//   dhcpv6: None
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_new_dhcpv6d_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_new_dhcpv6d_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg dhcpv6}", reflect.TypeOf(Dhcpv6{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg:dhcpv6", reflect.TypeOf(Dhcpv6{}))
}

// Insert represents Insert
type Insert string

const (
    // Insert locally generated/configured Interface
    // ID value
    Insert_local Insert = "local"

    // Insert received Interface ID value
    Insert_received Insert = "received"

    // Insert received Interface ID value from SADB
    Insert_pppoe Insert = "pppoe"
)

// SubscriberId represents Subscriber id
type SubscriberId string

const (
    // Insert Received Subscriber-ID Value from SADB
    SubscriberId_pppoe SubscriberId = "pppoe"
)

// LinkLayerAddr represents Link layer addr
type LinkLayerAddr string

const (
    // Insert Received LinkLayerAddr Value from SADB
    LinkLayerAddr_set LinkLayerAddr = "set"
)

// Dhcpv6
// None
// This type is a presence type.
type Dhcpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under DHCPv6. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // For BNG session, allow duid change for a client MAC. The type is
    // interface{}.
    AllowDuidChange interface{}

    // Enable DHCP binding database storage to file system.
    Database Dhcpv6_Database

    // Table of Profile.
    Profiles Dhcpv6_Profiles

    // Table of Interface.
    Interfaces Dhcpv6_Interfaces
}

func (dhcpv6 *Dhcpv6) GetEntityData() *types.CommonEntityData {
    dhcpv6.EntityData.YFilter = dhcpv6.YFilter
    dhcpv6.EntityData.YangName = "dhcpv6"
    dhcpv6.EntityData.BundleName = "cisco_ios_xr"
    dhcpv6.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg"
    dhcpv6.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg:dhcpv6"
    dhcpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6.EntityData.Children = make(map[string]types.YChild)
    dhcpv6.EntityData.Children["database"] = types.YChild{"Database", &dhcpv6.Database}
    dhcpv6.EntityData.Children["profiles"] = types.YChild{"Profiles", &dhcpv6.Profiles}
    dhcpv6.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &dhcpv6.Interfaces}
    dhcpv6.EntityData.Leafs = make(map[string]types.YLeaf)
    dhcpv6.EntityData.Leafs["enable"] = types.YLeaf{"Enable", dhcpv6.Enable}
    dhcpv6.EntityData.Leafs["allow-duid-change"] = types.YLeaf{"AllowDuidChange", dhcpv6.AllowDuidChange}
    return &(dhcpv6.EntityData)
}

// Dhcpv6_Database
// Enable DHCP binding database storage to file
// system
type Dhcpv6_Database struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable DHCP proxy binding database storage to file system. The type is
    // interface{}.
    Proxy interface{}

    // Enable DHCP server binding database storage to file system. The type is
    // interface{}.
    Server interface{}

    // Enable DHCP relay binding database storage to file system. The type is
    // interface{}.
    Relay interface{}

    // Full file write interval (default 10 minutes). The type is interface{} with
    // range: 1..1440. The default value is 10.
    FullWriteInterval interface{}

    // Incremental file write interval (default 1 minutes). The type is
    // interface{} with range: 1..1440. The default value is 1.
    IncrementalWriteInterval interface{}
}

func (database *Dhcpv6_Database) GetEntityData() *types.CommonEntityData {
    database.EntityData.YFilter = database.YFilter
    database.EntityData.YangName = "database"
    database.EntityData.BundleName = "cisco_ios_xr"
    database.EntityData.ParentYangName = "dhcpv6"
    database.EntityData.SegmentPath = "database"
    database.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    database.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    database.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    database.EntityData.Children = make(map[string]types.YChild)
    database.EntityData.Leafs = make(map[string]types.YLeaf)
    database.EntityData.Leafs["proxy"] = types.YLeaf{"Proxy", database.Proxy}
    database.EntityData.Leafs["server"] = types.YLeaf{"Server", database.Server}
    database.EntityData.Leafs["relay"] = types.YLeaf{"Relay", database.Relay}
    database.EntityData.Leafs["full-write-interval"] = types.YLeaf{"FullWriteInterval", database.FullWriteInterval}
    database.EntityData.Leafs["incremental-write-interval"] = types.YLeaf{"IncrementalWriteInterval", database.IncrementalWriteInterval}
    return &(database.EntityData)
}

// Dhcpv6_Profiles
// Table of Profile
type Dhcpv6_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile.
    Profile []Dhcpv6_Profiles_Profile
}

func (profiles *Dhcpv6_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "dhcpv6"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// Dhcpv6_Profiles_Profile
// None
type Dhcpv6_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProfileName interface{}

    // None.
    Relay Dhcpv6_Profiles_Profile_Relay

    // None.
    Base Dhcpv6_Profiles_Profile_Base

    // None.
    Proxy Dhcpv6_Profiles_Profile_Proxy

    // None.
    Server Dhcpv6_Profiles_Profile_Server
}

func (profile *Dhcpv6_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["relay"] = types.YChild{"Relay", &profile.Relay}
    profile.EntityData.Children["base"] = types.YChild{"Base", &profile.Base}
    profile.EntityData.Children["proxy"] = types.YChild{"Proxy", &profile.Proxy}
    profile.EntityData.Children["server"] = types.YChild{"Server", &profile.Server}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    return &(profile.EntityData)
}

// Dhcpv6_Profiles_Profile_Relay
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Relay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Relay. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Enable route addition for IANA. The type is interface{}.
    IanaRouteAdd interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Relay_HelperAddresses
}

func (relay *Dhcpv6_Profiles_Profile_Relay) GetEntityData() *types.CommonEntityData {
    relay.EntityData.YFilter = relay.YFilter
    relay.EntityData.YangName = "relay"
    relay.EntityData.BundleName = "cisco_ios_xr"
    relay.EntityData.ParentYangName = "profile"
    relay.EntityData.SegmentPath = "relay"
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = make(map[string]types.YChild)
    relay.EntityData.Children["helper-addresses"] = types.YChild{"HelperAddresses", &relay.HelperAddresses}
    relay.EntityData.Leafs = make(map[string]types.YLeaf)
    relay.EntityData.Leafs["enable"] = types.YLeaf{"Enable", relay.Enable}
    relay.EntityData.Leafs["iana-route-add"] = types.YLeaf{"IanaRouteAdd", relay.IanaRouteAdd}
    return &(relay.EntityData)
}

// Dhcpv6_Profiles_Profile_Relay_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Relay_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the server helper address. The type is slice of
    // Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress.
    HelperAddress []Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetEntityData() *types.CommonEntityData {
    helperAddresses.EntityData.YFilter = helperAddresses.YFilter
    helperAddresses.EntityData.YangName = "helper-addresses"
    helperAddresses.EntityData.BundleName = "cisco_ios_xr"
    helperAddresses.EntityData.ParentYangName = "relay"
    helperAddresses.EntityData.SegmentPath = "helper-addresses"
    helperAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddresses.EntityData.Children = make(map[string]types.YChild)
    helperAddresses.EntityData.Children["helper-address"] = types.YChild{"HelperAddress", nil}
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children[types.GetSegmentPath(&helperAddresses.HelperAddress[i])] = types.YChild{"HelperAddress", &helperAddresses.HelperAddress[i]}
    }
    helperAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helperAddresses.EntityData)
}

// Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress
// Specify the server helper address
type Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // This attribute is a key. Server Global unicast address. The type is string
    // with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    HelperAddress interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + "[vrf-name='" + fmt.Sprintf("%v", helperAddress.VrfName) + "']" + "[helper-address='" + fmt.Sprintf("%v", helperAddress.HelperAddress) + "']"
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = make(map[string]types.YChild)
    helperAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    helperAddress.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", helperAddress.VrfName}
    helperAddress.EntityData.Leafs["helper-address"] = types.YLeaf{"HelperAddress", helperAddress.HelperAddress}
    return &(helperAddress.EntityData)
}

// Dhcpv6_Profiles_Profile_Base
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Base. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Default match option.
    Default_ Dhcpv6_Profiles_Profile_Base_Default

    // Enter match option.
    Match Dhcpv6_Profiles_Profile_Base_Match
}

func (base *Dhcpv6_Profiles_Profile_Base) GetEntityData() *types.CommonEntityData {
    base.EntityData.YFilter = base.YFilter
    base.EntityData.YangName = "base"
    base.EntityData.BundleName = "cisco_ios_xr"
    base.EntityData.ParentYangName = "profile"
    base.EntityData.SegmentPath = "base"
    base.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    base.EntityData.Children = make(map[string]types.YChild)
    base.EntityData.Children["default"] = types.YChild{"Default_", &base.Default_}
    base.EntityData.Children["match"] = types.YChild{"Match", &base.Match}
    base.EntityData.Leafs = make(map[string]types.YLeaf)
    base.EntityData.Leafs["enable"] = types.YLeaf{"Enable", base.Enable}
    return &(base.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Default
// Default match option
type Dhcpv6_Profiles_Profile_Base_Default struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter proxy or server profile. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Default_Profile.
    Profile []Dhcpv6_Profiles_Profile_Base_Default_Profile_
}

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "default"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "base"
    self.EntityData.SegmentPath = "default"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range self.Profile {
        self.EntityData.Children[types.GetSegmentPath(&self.Profile[i])] = types.YChild{"Profile", &self.Profile[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Default_Profile_
// Enter proxy or server profile
type Dhcpv6_Profiles_Profile_Base_Default_Profile_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with length:
    // 1..64.
    ProfileName interface{}

    // Specify mode-class based Server option. The type is interface{}.
    ServerMode interface{}

    // Specify mode-class based Proxy Option. The type is interface{}.
    ProxyMode interface{}
}

func (profile_ *Dhcpv6_Profiles_Profile_Base_Default_Profile_) GetEntityData() *types.CommonEntityData {
    profile_.EntityData.YFilter = profile_.YFilter
    profile_.EntityData.YangName = "profile"
    profile_.EntityData.BundleName = "cisco_ios_xr"
    profile_.EntityData.ParentYangName = "default"
    profile_.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile_.ProfileName) + "']"
    profile_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile_.EntityData.Children = make(map[string]types.YChild)
    profile_.EntityData.Leafs = make(map[string]types.YLeaf)
    profile_.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile_.ProfileName}
    profile_.EntityData.Leafs["server-mode"] = types.YLeaf{"ServerMode", profile_.ServerMode}
    profile_.EntityData.Leafs["proxy-mode"] = types.YLeaf{"ProxyMode", profile_.ProxyMode}
    return &(profile_.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Match
// Enter match option
type Dhcpv6_Profiles_Profile_Base_Match struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of ModeClass.
    ModeClasses Dhcpv6_Profiles_Profile_Base_Match_ModeClasses
}

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "base"
    match.EntityData.SegmentPath = "match"
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = make(map[string]types.YChild)
    match.EntityData.Children["mode-classes"] = types.YChild{"ModeClasses", &match.ModeClasses}
    match.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(match.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses
// Table of ModeClass
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify PPP/IPoE class option. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass.
    ModeClass []Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass
}

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetEntityData() *types.CommonEntityData {
    modeClasses.EntityData.YFilter = modeClasses.YFilter
    modeClasses.EntityData.YangName = "mode-classes"
    modeClasses.EntityData.BundleName = "cisco_ios_xr"
    modeClasses.EntityData.ParentYangName = "match"
    modeClasses.EntityData.SegmentPath = "mode-classes"
    modeClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modeClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modeClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modeClasses.EntityData.Children = make(map[string]types.YChild)
    modeClasses.EntityData.Children["mode-class"] = types.YChild{"ModeClass", nil}
    for i := range modeClasses.ModeClass {
        modeClasses.EntityData.Children[types.GetSegmentPath(&modeClasses.ModeClass[i])] = types.YChild{"ModeClass", &modeClasses.ModeClass[i]}
    }
    modeClasses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(modeClasses.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass
// Specify PPP/IPoE class option
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Class name. The type is string with length:
    // 1..128.
    ClassName interface{}

    // Enter proxy or server profile. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile.
    Profile []Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile_
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetEntityData() *types.CommonEntityData {
    modeClass.EntityData.YFilter = modeClass.YFilter
    modeClass.EntityData.YangName = "mode-class"
    modeClass.EntityData.BundleName = "cisco_ios_xr"
    modeClass.EntityData.ParentYangName = "mode-classes"
    modeClass.EntityData.SegmentPath = "mode-class" + "[class-name='" + fmt.Sprintf("%v", modeClass.ClassName) + "']"
    modeClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modeClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modeClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modeClass.EntityData.Children = make(map[string]types.YChild)
    modeClass.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range modeClass.Profile {
        modeClass.EntityData.Children[types.GetSegmentPath(&modeClass.Profile[i])] = types.YChild{"Profile", &modeClass.Profile[i]}
    }
    modeClass.EntityData.Leafs = make(map[string]types.YLeaf)
    modeClass.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", modeClass.ClassName}
    return &(modeClass.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile_
// Enter proxy or server profile
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with length:
    // 1..64.
    ProfileName interface{}

    // Specify mode-class based Server option. The type is interface{}.
    ServerMode interface{}

    // Specify mode-class based Proxy Option. The type is interface{}.
    ProxyMode interface{}
}

func (profile_ *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile_) GetEntityData() *types.CommonEntityData {
    profile_.EntityData.YFilter = profile_.YFilter
    profile_.EntityData.YangName = "profile"
    profile_.EntityData.BundleName = "cisco_ios_xr"
    profile_.EntityData.ParentYangName = "mode-class"
    profile_.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile_.ProfileName) + "']"
    profile_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile_.EntityData.Children = make(map[string]types.YChild)
    profile_.EntityData.Leafs = make(map[string]types.YLeaf)
    profile_.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile_.ProfileName}
    profile_.EntityData.Leafs["server-mode"] = types.YLeaf{"ServerMode", profile_.ServerMode}
    profile_.EntityData.Leafs["proxy-mode"] = types.YLeaf{"ProxyMode", profile_.ProxyMode}
    return &(profile_.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Proxy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fill linkaddress in Relay fwd msg with Prefix from Router Advertisement for
    // PPPoE sessions. The type is interface{}.
    LinkaddressFromRaEnable interface{}

    // RouteDisable. The type is interface{}.
    RouteAddDisable interface{}

    // IPv6 address to be filled in link-address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkAddress interface{}

    // Create or enter proxy profile Source Interface Name. The type is string
    // with pattern: b'[a-zA-Z0-9./-]+'.
    SrcIntfName interface{}

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Proxy. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table of Interface.
    Interfaces Dhcpv6_Profiles_Profile_Proxy_Interfaces

    // Specify relay configuration.
    Relay Dhcpv6_Profiles_Profile_Proxy_Relay

    // VRF related configuration.
    Vrfs Dhcpv6_Profiles_Profile_Proxy_Vrfs

    // Authentication username format.
    Authentication Dhcpv6_Profiles_Profile_Proxy_Authentication

    // Table of Class.
    Classes Dhcpv6_Profiles_Profile_Proxy_Classes

    // Change sessions configuration.
    Sessions Dhcpv6_Profiles_Profile_Proxy_Sessions
}

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetEntityData() *types.CommonEntityData {
    proxy.EntityData.YFilter = proxy.YFilter
    proxy.EntityData.YangName = "proxy"
    proxy.EntityData.BundleName = "cisco_ios_xr"
    proxy.EntityData.ParentYangName = "profile"
    proxy.EntityData.SegmentPath = "proxy"
    proxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxy.EntityData.Children = make(map[string]types.YChild)
    proxy.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &proxy.Interfaces}
    proxy.EntityData.Children["relay"] = types.YChild{"Relay", &proxy.Relay}
    proxy.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &proxy.Vrfs}
    proxy.EntityData.Children["authentication"] = types.YChild{"Authentication", &proxy.Authentication}
    proxy.EntityData.Children["classes"] = types.YChild{"Classes", &proxy.Classes}
    proxy.EntityData.Children["sessions"] = types.YChild{"Sessions", &proxy.Sessions}
    proxy.EntityData.Leafs = make(map[string]types.YLeaf)
    proxy.EntityData.Leafs["linkaddress-from-ra-enable"] = types.YLeaf{"LinkaddressFromRaEnable", proxy.LinkaddressFromRaEnable}
    proxy.EntityData.Leafs["route-add-disable"] = types.YLeaf{"RouteAddDisable", proxy.RouteAddDisable}
    proxy.EntityData.Leafs["link-address"] = types.YLeaf{"LinkAddress", proxy.LinkAddress}
    proxy.EntityData.Leafs["src-intf-name"] = types.YLeaf{"SrcIntfName", proxy.SrcIntfName}
    proxy.EntityData.Leafs["enable"] = types.YLeaf{"Enable", proxy.Enable}
    return &(proxy.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Interfaces
// Table of Interface
type Dhcpv6_Profiles_Profile_Proxy_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface_.
    Interface_ []Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface
}

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "proxy"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface
// None
type Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Physical interface ID. The type is string.
    InterfaceId interface{}
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", self.InterfaceId}
    return &(self.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Relay
// Specify relay configuration
type Dhcpv6_Profiles_Profile_Proxy_Relay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify relay option configuration.
    Option Dhcpv6_Profiles_Profile_Proxy_Relay_Option
}

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetEntityData() *types.CommonEntityData {
    relay.EntityData.YFilter = relay.YFilter
    relay.EntityData.YangName = "relay"
    relay.EntityData.BundleName = "cisco_ios_xr"
    relay.EntityData.ParentYangName = "proxy"
    relay.EntityData.SegmentPath = "relay"
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = make(map[string]types.YChild)
    relay.EntityData.Children["option"] = types.YChild{"Option", &relay.Option}
    relay.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(relay.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Relay_Option
// Specify relay option configuration
type Dhcpv6_Profiles_Profile_Proxy_Relay_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Received SubscriberID. The type is SubscriberId.
    SubscriberId interface{}

    // Configure Received link-layer-Addr. The type is LinkLayerAddr.
    LinkLayerAddr interface{}

    // Set remote-id value from SADB. The type is interface{} with range:
    // -2147483648..2147483647.
    RemoteIDreceived interface{}

    // Enter remote-id value. The type is string with length: 1..256.
    RemoteId interface{}

    // Interface Id option.
    InterfaceId Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId
}

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "relay"
    option.EntityData.SegmentPath = "option"
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = make(map[string]types.YChild)
    option.EntityData.Children["interface-id"] = types.YChild{"InterfaceId", &option.InterfaceId}
    option.EntityData.Leafs = make(map[string]types.YLeaf)
    option.EntityData.Leafs["subscriber-id"] = types.YLeaf{"SubscriberId", option.SubscriberId}
    option.EntityData.Leafs["link-layer-addr"] = types.YLeaf{"LinkLayerAddr", option.LinkLayerAddr}
    option.EntityData.Leafs["remote-i-dreceived"] = types.YLeaf{"RemoteIDreceived", option.RemoteIDreceived}
    option.EntityData.Leafs["remote-id"] = types.YLeaf{"RemoteId", option.RemoteId}
    return &(option.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId
// Interface Id option
type Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure InterfaceID insert type. The type is Insert.
    Insert interface{}
}

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetEntityData() *types.CommonEntityData {
    interfaceId.EntityData.YFilter = interfaceId.YFilter
    interfaceId.EntityData.YangName = "interface-id"
    interfaceId.EntityData.BundleName = "cisco_ios_xr"
    interfaceId.EntityData.ParentYangName = "option"
    interfaceId.EntityData.SegmentPath = "interface-id"
    interfaceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceId.EntityData.Children = make(map[string]types.YChild)
    interfaceId.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceId.EntityData.Leafs["insert"] = types.YLeaf{"Insert", interfaceId.Insert}
    return &(interfaceId.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs
// VRF related configuration
type Dhcpv6_Profiles_Profile_Proxy_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy VRF name. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf.
    Vrf []Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "proxy"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf
// IPv6 DHCP proxy VRF name
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["helper-addresses"] = types.YChild{"HelperAddresses", &vrf.HelperAddresses}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 Helper Address. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetEntityData() *types.CommonEntityData {
    helperAddresses.EntityData.YFilter = helperAddresses.YFilter
    helperAddresses.EntityData.YangName = "helper-addresses"
    helperAddresses.EntityData.BundleName = "cisco_ios_xr"
    helperAddresses.EntityData.ParentYangName = "vrf"
    helperAddresses.EntityData.SegmentPath = "helper-addresses"
    helperAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddresses.EntityData.Children = make(map[string]types.YChild)
    helperAddresses.EntityData.Children["helper-address"] = types.YChild{"HelperAddress", nil}
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children[types.GetSegmentPath(&helperAddresses.HelperAddress[i])] = types.YChild{"HelperAddress", &helperAddresses.HelperAddress[i]}
    }
    helperAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helperAddresses.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
// DHCPv6 Helper Address
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. DHCPv6 Helper Address. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    HelperAddress interface{}

    // DHCPv6 HelperAddress Specific Output Interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    OutInterface interface{}

    // DHCPv6 HelperAddress Output Interface. The type is interface{}.
    AnyOutInterface interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + "[helper-address='" + fmt.Sprintf("%v", helperAddress.HelperAddress) + "']"
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = make(map[string]types.YChild)
    helperAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    helperAddress.EntityData.Leafs["helper-address"] = types.YLeaf{"HelperAddress", helperAddress.HelperAddress}
    helperAddress.EntityData.Leafs["out-interface"] = types.YLeaf{"OutInterface", helperAddress.OutInterface}
    helperAddress.EntityData.Leafs["any-out-interface"] = types.YLeaf{"AnyOutInterface", helperAddress.AnyOutInterface}
    return &(helperAddress.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Authentication
// Authentication username format
type Dhcpv6_Profiles_Profile_Proxy_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set username as DUID. The type is interface{}.
    Username interface{}
}

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "proxy"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["username"] = types.YLeaf{"Username", authentication.Username}
    return &(authentication.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Classes
// Table of Class
type Dhcpv6_Profiles_Profile_Proxy_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile_Proxy_Classes_Class.
    Class []Dhcpv6_Profiles_Profile_Proxy_Classes_Class
}

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "proxy"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Classes_Class
// None
type Dhcpv6_Profiles_Profile_Proxy_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Class name. The type is string with length:
    // 1..128.
    ClassName interface{}

    // IPv6 address to be filled in link-address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkAddress interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["helper-addresses"] = types.YChild{"HelperAddresses", &class.HelperAddresses}
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", class.ClassName}
    class.EntityData.Leafs["link-address"] = types.YLeaf{"LinkAddress", class.LinkAddress}
    return &(class.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the server helper address. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress.
    HelperAddress []Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetEntityData() *types.CommonEntityData {
    helperAddresses.EntityData.YFilter = helperAddresses.YFilter
    helperAddresses.EntityData.YangName = "helper-addresses"
    helperAddresses.EntityData.BundleName = "cisco_ios_xr"
    helperAddresses.EntityData.ParentYangName = "class"
    helperAddresses.EntityData.SegmentPath = "helper-addresses"
    helperAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddresses.EntityData.Children = make(map[string]types.YChild)
    helperAddresses.EntityData.Children["helper-address"] = types.YChild{"HelperAddress", nil}
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children[types.GetSegmentPath(&helperAddresses.HelperAddress[i])] = types.YChild{"HelperAddress", &helperAddresses.HelperAddress[i]}
    }
    helperAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helperAddresses.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress
// Specify the server helper address
type Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // This attribute is a key. Server address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    HelperAddress interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + "[vrf-name='" + fmt.Sprintf("%v", helperAddress.VrfName) + "']" + "[helper-address='" + fmt.Sprintf("%v", helperAddress.HelperAddress) + "']"
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = make(map[string]types.YChild)
    helperAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    helperAddress.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", helperAddress.VrfName}
    helperAddress.EntityData.Leafs["helper-address"] = types.YLeaf{"HelperAddress", helperAddress.HelperAddress}
    return &(helperAddress.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Sessions
// Change sessions configuration
type Dhcpv6_Profiles_Profile_Proxy_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    Mac Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac
}

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "proxy"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["mac"] = types.YChild{"Mac", &sessions.Mac}
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac
// Throttle DHCP sessions based on MAC address
type Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    Throttle Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle
}

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "sessions"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["throttle"] = types.YChild{"Throttle", &mac.Throttle}
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mac.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle
// Throttle DHCP sessions from any one MAC
// address
type Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of solicits at which to throttle. The type is interface{} with
    // range: 1..65535.
    Limit interface{}

    // Throttle request period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Request interface{}

    // Throttle blocking period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Block interface{}
}

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "mac"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = make(map[string]types.YChild)
    throttle.EntityData.Leafs = make(map[string]types.YLeaf)
    throttle.EntityData.Leafs["limit"] = types.YLeaf{"Limit", throttle.Limit}
    throttle.EntityData.Leafs["request"] = types.YLeaf{"Request", throttle.Request}
    throttle.EntityData.Leafs["block"] = types.YLeaf{"Block", throttle.Block}
    return &(throttle.EntityData)
}

// Dhcpv6_Profiles_Profile_Server
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address pool name. The type is string with length: 1..64.
    AddressPool interface{}

    // AFTR name. The type is string with length: 1..64.
    AftrName interface{}

    // Domain name. The type is string with length: 1..64.
    DomainName interface{}

    // DHCP Server Preference. The type is interface{} with range: 0..255.
    Preference interface{}

    // Allow RAPID Commit. The type is interface{}.
    RapidCommit interface{}

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Server. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Prefix pool name. The type is string with length: 1..64.
    PrefixPool interface{}

    // Change sessions configuration.
    Sessions Dhcpv6_Profiles_Profile_Server_Sessions

    // DNS servers.
    DnsServers Dhcpv6_Profiles_Profile_Server_DnsServers

    // Table of Class.
    Classes Dhcpv6_Profiles_Profile_Server_Classes

    // lease.
    Lease Dhcpv6_Profiles_Profile_Server_Lease

    // DHCPv6 options.
    Dhcpv6Options Dhcpv6_Profiles_Profile_Server_Dhcpv6Options
}

func (server *Dhcpv6_Profiles_Profile_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "profile"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Children["sessions"] = types.YChild{"Sessions", &server.Sessions}
    server.EntityData.Children["dns-servers"] = types.YChild{"DnsServers", &server.DnsServers}
    server.EntityData.Children["classes"] = types.YChild{"Classes", &server.Classes}
    server.EntityData.Children["lease"] = types.YChild{"Lease", &server.Lease}
    server.EntityData.Children["dhcpv6-options"] = types.YChild{"Dhcpv6Options", &server.Dhcpv6Options}
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    server.EntityData.Leafs["address-pool"] = types.YLeaf{"AddressPool", server.AddressPool}
    server.EntityData.Leafs["aftr-name"] = types.YLeaf{"AftrName", server.AftrName}
    server.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", server.DomainName}
    server.EntityData.Leafs["preference"] = types.YLeaf{"Preference", server.Preference}
    server.EntityData.Leafs["rapid-commit"] = types.YLeaf{"RapidCommit", server.RapidCommit}
    server.EntityData.Leafs["enable"] = types.YLeaf{"Enable", server.Enable}
    server.EntityData.Leafs["prefix-pool"] = types.YLeaf{"PrefixPool", server.PrefixPool}
    return &(server.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Sessions
// Change sessions configuration
type Dhcpv6_Profiles_Profile_Server_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    Mac Dhcpv6_Profiles_Profile_Server_Sessions_Mac
}

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "server"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["mac"] = types.YChild{"Mac", &sessions.Mac}
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Sessions_Mac
// Throttle DHCP sessions based on MAC address
type Dhcpv6_Profiles_Profile_Server_Sessions_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    Throttle Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle
}

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "sessions"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["throttle"] = types.YChild{"Throttle", &mac.Throttle}
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mac.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle
// Throttle DHCP sessions from any one MAC
// address
type Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of solicits at which to throttle. The type is interface{} with
    // range: 1..65535.
    Limit interface{}

    // Throttle request period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Request interface{}

    // Throttle blocking period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Block interface{}
}

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "mac"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = make(map[string]types.YChild)
    throttle.EntityData.Leafs = make(map[string]types.YLeaf)
    throttle.EntityData.Leafs["limit"] = types.YLeaf{"Limit", throttle.Limit}
    throttle.EntityData.Leafs["request"] = types.YLeaf{"Request", throttle.Request}
    throttle.EntityData.Leafs["block"] = types.YLeaf{"Block", throttle.Block}
    return &(throttle.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_DnsServers
// DNS servers
type Dhcpv6_Profiles_Profile_Server_DnsServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server's IPv6 address. The type is one of the following types: slice of
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DnsServer []interface{}
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetEntityData() *types.CommonEntityData {
    dnsServers.EntityData.YFilter = dnsServers.YFilter
    dnsServers.EntityData.YangName = "dns-servers"
    dnsServers.EntityData.BundleName = "cisco_ios_xr"
    dnsServers.EntityData.ParentYangName = "server"
    dnsServers.EntityData.SegmentPath = "dns-servers"
    dnsServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnsServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnsServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnsServers.EntityData.Children = make(map[string]types.YChild)
    dnsServers.EntityData.Leafs = make(map[string]types.YLeaf)
    dnsServers.EntityData.Leafs["dns-server"] = types.YLeaf{"DnsServer", dnsServers.DnsServer}
    return &(dnsServers.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Classes
// Table of Class
type Dhcpv6_Profiles_Profile_Server_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile_Server_Classes_Class.
    Class []Dhcpv6_Profiles_Profile_Server_Classes_Class
}

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "server"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Classes_Class
// None
type Dhcpv6_Profiles_Profile_Server_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. class name. The type is string with length:
    // 1..128.
    ClassName interface{}

    // Address pool name. The type is string with length: 1..64.
    AddressPool interface{}

    // Domain name. The type is string with length: 1..64.
    DomainName interface{}

    // DHCP Server Preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // Prefix pool name. The type is string with length: 1..64.
    PrefixPool interface{}

    // DNS servers.
    DnsServers Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["dns-servers"] = types.YChild{"DnsServers", &class.DnsServers}
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", class.ClassName}
    class.EntityData.Leafs["address-pool"] = types.YLeaf{"AddressPool", class.AddressPool}
    class.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", class.DomainName}
    class.EntityData.Leafs["preference"] = types.YLeaf{"Preference", class.Preference}
    class.EntityData.Leafs["prefix-pool"] = types.YLeaf{"PrefixPool", class.PrefixPool}
    return &(class.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers
// DNS servers
type Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server's IPv6 address. The type is one of the following types: slice of
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DnsServer []interface{}
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetEntityData() *types.CommonEntityData {
    dnsServers.EntityData.YFilter = dnsServers.YFilter
    dnsServers.EntityData.YangName = "dns-servers"
    dnsServers.EntityData.BundleName = "cisco_ios_xr"
    dnsServers.EntityData.ParentYangName = "class"
    dnsServers.EntityData.SegmentPath = "dns-servers"
    dnsServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnsServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnsServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnsServers.EntityData.Children = make(map[string]types.YChild)
    dnsServers.EntityData.Leafs = make(map[string]types.YLeaf)
    dnsServers.EntityData.Leafs["dns-server"] = types.YLeaf{"DnsServer", dnsServers.DnsServer}
    return &(dnsServers.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Lease
// lease
type Dhcpv6_Profiles_Profile_Server_Lease struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Days. The type is interface{} with range: 0..365. Units are day.
    Days interface{}

    // Hours. The type is interface{} with range: 0..23. Units are hour.
    Hours interface{}

    // Minutes. The type is interface{} with range: 1..59. Units are minute.
    Minutes interface{}

    // Set string. The type is string.
    Infinite interface{}
}

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetEntityData() *types.CommonEntityData {
    lease.EntityData.YFilter = lease.YFilter
    lease.EntityData.YangName = "lease"
    lease.EntityData.BundleName = "cisco_ios_xr"
    lease.EntityData.ParentYangName = "server"
    lease.EntityData.SegmentPath = "lease"
    lease.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lease.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lease.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lease.EntityData.Children = make(map[string]types.YChild)
    lease.EntityData.Leafs = make(map[string]types.YLeaf)
    lease.EntityData.Leafs["days"] = types.YLeaf{"Days", lease.Days}
    lease.EntityData.Leafs["hours"] = types.YLeaf{"Hours", lease.Hours}
    lease.EntityData.Leafs["minutes"] = types.YLeaf{"Minutes", lease.Minutes}
    lease.EntityData.Leafs["infinite"] = types.YLeaf{"Infinite", lease.Infinite}
    return &(lease.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Dhcpv6Options
// DHCPv6 options
type Dhcpv6_Profiles_Profile_Server_Dhcpv6Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vendor options.
    VendorOptions Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions
}

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetEntityData() *types.CommonEntityData {
    dhcpv6Options.EntityData.YFilter = dhcpv6Options.YFilter
    dhcpv6Options.EntityData.YangName = "dhcpv6-options"
    dhcpv6Options.EntityData.BundleName = "cisco_ios_xr"
    dhcpv6Options.EntityData.ParentYangName = "server"
    dhcpv6Options.EntityData.SegmentPath = "dhcpv6-options"
    dhcpv6Options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6Options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6Options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6Options.EntityData.Children = make(map[string]types.YChild)
    dhcpv6Options.EntityData.Children["vendor-options"] = types.YChild{"VendorOptions", &dhcpv6Options.VendorOptions}
    dhcpv6Options.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dhcpv6Options.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions
// Vendor options
type Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set string. The type is string.
    Type_ interface{}

    // Vendor options. The type is string with length: 1..512.
    VendorOptions interface{}
}

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetEntityData() *types.CommonEntityData {
    vendorOptions.EntityData.YFilter = vendorOptions.YFilter
    vendorOptions.EntityData.YangName = "vendor-options"
    vendorOptions.EntityData.BundleName = "cisco_ios_xr"
    vendorOptions.EntityData.ParentYangName = "dhcpv6-options"
    vendorOptions.EntityData.SegmentPath = "vendor-options"
    vendorOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vendorOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vendorOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vendorOptions.EntityData.Children = make(map[string]types.YChild)
    vendorOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    vendorOptions.EntityData.Leafs["type"] = types.YLeaf{"Type_", vendorOptions.Type_}
    vendorOptions.EntityData.Leafs["vendor-options"] = types.YLeaf{"VendorOptions", vendorOptions.VendorOptions}
    return &(vendorOptions.EntityData)
}

// Dhcpv6_Interfaces
// Table of Interface
type Dhcpv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Interfaces_Interface_.
    Interface_ []Dhcpv6_Interfaces_Interface
}

func (interfaces *Dhcpv6_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "dhcpv6"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Dhcpv6_Interfaces_Interface
// None
type Dhcpv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // PPPoE subscriber interface.
    Pppoe Dhcpv6_Interfaces_Interface_Pppoe

    // Assign a proxy profile to interface.
    Proxy Dhcpv6_Interfaces_Interface_Proxy

    // Assign a base profile to interface.
    Base Dhcpv6_Interfaces_Interface_Base

    // Assign a server profile to interface.
    Server Dhcpv6_Interfaces_Interface_Server

    // Assign a relay profile to interface.
    Relay Dhcpv6_Interfaces_Interface_Relay
}

func (self *Dhcpv6_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &self.Pppoe}
    self.EntityData.Children["proxy"] = types.YChild{"Proxy", &self.Proxy}
    self.EntityData.Children["base"] = types.YChild{"Base", &self.Base}
    self.EntityData.Children["server"] = types.YChild{"Server", &self.Server}
    self.EntityData.Children["relay"] = types.YChild{"Relay", &self.Relay}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// Dhcpv6_Interfaces_Interface_Pppoe
// PPPoE subscriber interface
type Dhcpv6_Interfaces_Interface_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "interface"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["profile"] = types.YLeaf{"Profile", pppoe.Profile}
    return &(pppoe.EntityData)
}

// Dhcpv6_Interfaces_Interface_Proxy
// Assign a proxy profile to interface
type Dhcpv6_Interfaces_Interface_Proxy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetEntityData() *types.CommonEntityData {
    proxy.EntityData.YFilter = proxy.YFilter
    proxy.EntityData.YangName = "proxy"
    proxy.EntityData.BundleName = "cisco_ios_xr"
    proxy.EntityData.ParentYangName = "interface"
    proxy.EntityData.SegmentPath = "proxy"
    proxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxy.EntityData.Children = make(map[string]types.YChild)
    proxy.EntityData.Leafs = make(map[string]types.YLeaf)
    proxy.EntityData.Leafs["profile"] = types.YLeaf{"Profile", proxy.Profile}
    return &(proxy.EntityData)
}

// Dhcpv6_Interfaces_Interface_Base
// Assign a base profile to interface
type Dhcpv6_Interfaces_Interface_Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (base *Dhcpv6_Interfaces_Interface_Base) GetEntityData() *types.CommonEntityData {
    base.EntityData.YFilter = base.YFilter
    base.EntityData.YangName = "base"
    base.EntityData.BundleName = "cisco_ios_xr"
    base.EntityData.ParentYangName = "interface"
    base.EntityData.SegmentPath = "base"
    base.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    base.EntityData.Children = make(map[string]types.YChild)
    base.EntityData.Leafs = make(map[string]types.YLeaf)
    base.EntityData.Leafs["profile"] = types.YLeaf{"Profile", base.Profile}
    return &(base.EntityData)
}

// Dhcpv6_Interfaces_Interface_Server
// Assign a server profile to interface
type Dhcpv6_Interfaces_Interface_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (server *Dhcpv6_Interfaces_Interface_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "interface"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    server.EntityData.Leafs["profile"] = types.YLeaf{"Profile", server.Profile}
    return &(server.EntityData)
}

// Dhcpv6_Interfaces_Interface_Relay
// Assign a relay profile to interface
type Dhcpv6_Interfaces_Interface_Relay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetEntityData() *types.CommonEntityData {
    relay.EntityData.YFilter = relay.YFilter
    relay.EntityData.YangName = "relay"
    relay.EntityData.BundleName = "cisco_ios_xr"
    relay.EntityData.ParentYangName = "interface"
    relay.EntityData.SegmentPath = "relay"
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = make(map[string]types.YChild)
    relay.EntityData.Leafs = make(map[string]types.YLeaf)
    relay.EntityData.Leafs["profile"] = types.YLeaf{"Profile", relay.Profile}
    return &(relay.EntityData)
}

