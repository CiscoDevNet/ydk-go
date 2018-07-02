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

// Action represents Action
type Action string

const (
    // Allow vendor specific DHCP Solicit
    Action_allow Action = "allow"

    // Drop vendor specific DHCP Solicit
    Action_drop Action = "drop"
)

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

// LinkLayerAddr represents Link layer addr
type LinkLayerAddr string

const (
    // Insert Received LinkLayerAddr Value from SADB
    LinkLayerAddr_set LinkLayerAddr = "set"
)

// SubscriberId represents Subscriber id
type SubscriberId string

const (
    // Insert Received Subscriber-ID Value from SADB
    SubscriberId_pppoe SubscriberId = "pppoe"
)

// Dhcpv6
// None
// This type is a presence type.
type Dhcpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Inner cos values for DHCPv6 packets to wards clients. The type is
    // interface{} with range: 0..7.
    InnerCos interface{}

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under DHCPv6. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // For BNG session, allow duid change for a client MAC. The type is
    // interface{}.
    AllowDuidChange interface{}

    // Configure outer cos values for DHCPv6 packet to wards client. The type is
    // interface{} with range: 0..7.
    OuterCos interface{}

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

    dhcpv6.EntityData.Children = types.NewOrderedMap()
    dhcpv6.EntityData.Children.Append("database", types.YChild{"Database", &dhcpv6.Database})
    dhcpv6.EntityData.Children.Append("profiles", types.YChild{"Profiles", &dhcpv6.Profiles})
    dhcpv6.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &dhcpv6.Interfaces})
    dhcpv6.EntityData.Leafs = types.NewOrderedMap()
    dhcpv6.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", dhcpv6.InnerCos})
    dhcpv6.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", dhcpv6.Enable})
    dhcpv6.EntityData.Leafs.Append("allow-duid-change", types.YLeaf{"AllowDuidChange", dhcpv6.AllowDuidChange})
    dhcpv6.EntityData.Leafs.Append("outer-cos", types.YLeaf{"OuterCos", dhcpv6.OuterCos})

    dhcpv6.EntityData.YListKeys = []string {}

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

    database.EntityData.Children = types.NewOrderedMap()
    database.EntityData.Leafs = types.NewOrderedMap()
    database.EntityData.Leafs.Append("proxy", types.YLeaf{"Proxy", database.Proxy})
    database.EntityData.Leafs.Append("server", types.YLeaf{"Server", database.Server})
    database.EntityData.Leafs.Append("relay", types.YLeaf{"Relay", database.Relay})
    database.EntityData.Leafs.Append("full-write-interval", types.YLeaf{"FullWriteInterval", database.FullWriteInterval})
    database.EntityData.Leafs.Append("incremental-write-interval", types.YLeaf{"IncrementalWriteInterval", database.IncrementalWriteInterval})

    database.EntityData.YListKeys = []string {}

    return &(database.EntityData)
}

// Dhcpv6_Profiles
// Table of Profile
type Dhcpv6_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile.
    Profile []*Dhcpv6_Profiles_Profile
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

    profiles.EntityData.Children = types.NewOrderedMap()
    profiles.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range profiles.Profile {
        profiles.EntityData.Children.Append(types.GetSegmentPath(profiles.Profile[i]), types.YChild{"Profile", profiles.Profile[i]})
    }
    profiles.EntityData.Leafs = types.NewOrderedMap()

    profiles.EntityData.YListKeys = []string {}

    return &(profiles.EntityData)
}

// Dhcpv6_Profiles_Profile
// None
type Dhcpv6_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("relay", types.YChild{"Relay", &profile.Relay})
    profile.EntityData.Children.Append("base", types.YChild{"Base", &profile.Base})
    profile.EntityData.Children.Append("proxy", types.YChild{"Proxy", &profile.Proxy})
    profile.EntityData.Children.Append("server", types.YChild{"Server", &profile.Server})
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
}

// Dhcpv6_Profiles_Profile_Relay
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Relay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Relay profile Source Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SrcIntfName interface{}

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Relay. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Enable route addition for IANA. The type is interface{}.
    IanaRouteAdd interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Relay_HelperAddresses

    // Specify relay option configuration.
    Option Dhcpv6_Profiles_Profile_Relay_Option
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

    relay.EntityData.Children = types.NewOrderedMap()
    relay.EntityData.Children.Append("helper-addresses", types.YChild{"HelperAddresses", &relay.HelperAddresses})
    relay.EntityData.Children.Append("option", types.YChild{"Option", &relay.Option})
    relay.EntityData.Leafs = types.NewOrderedMap()
    relay.EntityData.Leafs.Append("src-intf-name", types.YLeaf{"SrcIntfName", relay.SrcIntfName})
    relay.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", relay.Enable})
    relay.EntityData.Leafs.Append("iana-route-add", types.YLeaf{"IanaRouteAdd", relay.IanaRouteAdd})

    relay.EntityData.YListKeys = []string {}

    return &(relay.EntityData)
}

// Dhcpv6_Profiles_Profile_Relay_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Relay_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the server helper address. The type is slice of
    // Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress.
    HelperAddress []*Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress
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

    helperAddresses.EntityData.Children = types.NewOrderedMap()
    helperAddresses.EntityData.Children.Append("helper-address", types.YChild{"HelperAddress", nil})
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children.Append(types.GetSegmentPath(helperAddresses.HelperAddress[i]), types.YChild{"HelperAddress", helperAddresses.HelperAddress[i]})
    }
    helperAddresses.EntityData.Leafs = types.NewOrderedMap()

    helperAddresses.EntityData.YListKeys = []string {}

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
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HelperAddress interface{}

    // Enable. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Helper-address Specific Source Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SrcIntfName interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + types.AddKeyToken(helperAddress.VrfName, "vrf-name") + types.AddKeyToken(helperAddress.HelperAddress, "helper-address")
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = types.NewOrderedMap()
    helperAddress.EntityData.Leafs = types.NewOrderedMap()
    helperAddress.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", helperAddress.VrfName})
    helperAddress.EntityData.Leafs.Append("helper-address", types.YLeaf{"HelperAddress", helperAddress.HelperAddress})
    helperAddress.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", helperAddress.Enable})
    helperAddress.EntityData.Leafs.Append("src-intf-name", types.YLeaf{"SrcIntfName", helperAddress.SrcIntfName})

    helperAddress.EntityData.YListKeys = []string {"VrfName", "HelperAddress"}

    return &(helperAddress.EntityData)
}

// Dhcpv6_Profiles_Profile_Relay_Option
// Specify relay option configuration
type Dhcpv6_Profiles_Profile_Relay_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter remote-id value. The type is string with length: 1..256.
    RemoteId interface{}
}

func (option *Dhcpv6_Profiles_Profile_Relay_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "relay"
    option.EntityData.SegmentPath = "option"
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Leafs = types.NewOrderedMap()
    option.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", option.RemoteId})

    option.EntityData.YListKeys = []string {}

    return &(option.EntityData)
}

// Dhcpv6_Profiles_Profile_Base
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Base. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Default match option.
    Default Dhcpv6_Profiles_Profile_Base_Default

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

    base.EntityData.Children = types.NewOrderedMap()
    base.EntityData.Children.Append("default", types.YChild{"Default", &base.Default})
    base.EntityData.Children.Append("match", types.YChild{"Match", &base.Match})
    base.EntityData.Leafs = types.NewOrderedMap()
    base.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", base.Enable})

    base.EntityData.YListKeys = []string {}

    return &(base.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Default
// Default match option
type Dhcpv6_Profiles_Profile_Base_Default struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter proxy or server profile. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Default_Profile.
    Profile []*Dhcpv6_Profiles_Profile_Base_Default_Profile
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

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range self.Profile {
        self.EntityData.Children.Append(types.GetSegmentPath(self.Profile[i]), types.YChild{"Profile", self.Profile[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Default_Profile
// Enter proxy or server profile
type Dhcpv6_Profiles_Profile_Base_Default_Profile struct {
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

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "default"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})
    profile.EntityData.Leafs.Append("server-mode", types.YLeaf{"ServerMode", profile.ServerMode})
    profile.EntityData.Leafs.Append("proxy-mode", types.YLeaf{"ProxyMode", profile.ProxyMode})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
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

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Children.Append("mode-classes", types.YChild{"ModeClasses", &match.ModeClasses})
    match.EntityData.Leafs = types.NewOrderedMap()

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses
// Table of ModeClass
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify PPP/IPoE class option. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass.
    ModeClass []*Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass
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

    modeClasses.EntityData.Children = types.NewOrderedMap()
    modeClasses.EntityData.Children.Append("mode-class", types.YChild{"ModeClass", nil})
    for i := range modeClasses.ModeClass {
        modeClasses.EntityData.Children.Append(types.GetSegmentPath(modeClasses.ModeClass[i]), types.YChild{"ModeClass", modeClasses.ModeClass[i]})
    }
    modeClasses.EntityData.Leafs = types.NewOrderedMap()

    modeClasses.EntityData.YListKeys = []string {}

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
    Profile []*Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetEntityData() *types.CommonEntityData {
    modeClass.EntityData.YFilter = modeClass.YFilter
    modeClass.EntityData.YangName = "mode-class"
    modeClass.EntityData.BundleName = "cisco_ios_xr"
    modeClass.EntityData.ParentYangName = "mode-classes"
    modeClass.EntityData.SegmentPath = "mode-class" + types.AddKeyToken(modeClass.ClassName, "class-name")
    modeClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modeClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modeClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modeClass.EntityData.Children = types.NewOrderedMap()
    modeClass.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range modeClass.Profile {
        modeClass.EntityData.Children.Append(types.GetSegmentPath(modeClass.Profile[i]), types.YChild{"Profile", modeClass.Profile[i]})
    }
    modeClass.EntityData.Leafs = types.NewOrderedMap()
    modeClass.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", modeClass.ClassName})

    modeClass.EntityData.YListKeys = []string {"ClassName"}

    return &(modeClass.EntityData)
}

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile
// Enter proxy or server profile
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile struct {
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

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "mode-class"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})
    profile.EntityData.Leafs.Append("server-mode", types.YLeaf{"ServerMode", profile.ServerMode})
    profile.EntityData.Leafs.Append("proxy-mode", types.YLeaf{"ProxyMode", profile.ProxyMode})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Proxy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Fill linkaddress in Relay fwd msg with Prefix from Router Advertisement for
    // PPPoE sessions. The type is interface{}.
    LinkaddressFromRaEnable interface{}

    // RouteDisable. The type is interface{}.
    RouteAddDisable interface{}

    // IPv6 address to be filled in link-address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkAddress interface{}

    // Create or enter proxy profile Source Interface Name. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
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

    proxy.EntityData.Children = types.NewOrderedMap()
    proxy.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &proxy.Interfaces})
    proxy.EntityData.Children.Append("relay", types.YChild{"Relay", &proxy.Relay})
    proxy.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &proxy.Vrfs})
    proxy.EntityData.Children.Append("authentication", types.YChild{"Authentication", &proxy.Authentication})
    proxy.EntityData.Children.Append("classes", types.YChild{"Classes", &proxy.Classes})
    proxy.EntityData.Children.Append("sessions", types.YChild{"Sessions", &proxy.Sessions})
    proxy.EntityData.Leafs = types.NewOrderedMap()
    proxy.EntityData.Leafs.Append("linkaddress-from-ra-enable", types.YLeaf{"LinkaddressFromRaEnable", proxy.LinkaddressFromRaEnable})
    proxy.EntityData.Leafs.Append("route-add-disable", types.YLeaf{"RouteAddDisable", proxy.RouteAddDisable})
    proxy.EntityData.Leafs.Append("link-address", types.YLeaf{"LinkAddress", proxy.LinkAddress})
    proxy.EntityData.Leafs.Append("src-intf-name", types.YLeaf{"SrcIntfName", proxy.SrcIntfName})
    proxy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", proxy.Enable})

    proxy.EntityData.YListKeys = []string {}

    return &(proxy.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Interfaces
// Table of Interface
type Dhcpv6_Profiles_Profile_Proxy_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface.
    Interface []*Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface
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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface
// None
type Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Physical interface ID. The type is string.
    InterfaceId interface{}
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", self.InterfaceId})

    self.EntityData.YListKeys = []string {"InterfaceName"}

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

    relay.EntityData.Children = types.NewOrderedMap()
    relay.EntityData.Children.Append("option", types.YChild{"Option", &relay.Option})
    relay.EntityData.Leafs = types.NewOrderedMap()

    relay.EntityData.YListKeys = []string {}

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
    // 0..4294967295.
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

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Children.Append("interface-id", types.YChild{"InterfaceId", &option.InterfaceId})
    option.EntityData.Leafs = types.NewOrderedMap()
    option.EntityData.Leafs.Append("subscriber-id", types.YLeaf{"SubscriberId", option.SubscriberId})
    option.EntityData.Leafs.Append("link-layer-addr", types.YLeaf{"LinkLayerAddr", option.LinkLayerAddr})
    option.EntityData.Leafs.Append("remote-i-dreceived", types.YLeaf{"RemoteIDreceived", option.RemoteIDreceived})
    option.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", option.RemoteId})

    option.EntityData.YListKeys = []string {}

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

    interfaceId.EntityData.Children = types.NewOrderedMap()
    interfaceId.EntityData.Leafs = types.NewOrderedMap()
    interfaceId.EntityData.Leafs.Append("insert", types.YLeaf{"Insert", interfaceId.Insert})

    interfaceId.EntityData.YListKeys = []string {}

    return &(interfaceId.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs
// VRF related configuration
type Dhcpv6_Profiles_Profile_Proxy_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy VRF name. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf.
    Vrf []*Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf
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

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf
// IPv6 DHCP proxy VRF name
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("helper-addresses", types.YChild{"HelperAddresses", &vrf.HelperAddresses})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 Helper Address. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []*Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
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

    helperAddresses.EntityData.Children = types.NewOrderedMap()
    helperAddresses.EntityData.Children.Append("helper-address", types.YChild{"HelperAddress", nil})
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children.Append(types.GetSegmentPath(helperAddresses.HelperAddress[i]), types.YChild{"HelperAddress", helperAddresses.HelperAddress[i]})
    }
    helperAddresses.EntityData.Leafs = types.NewOrderedMap()

    helperAddresses.EntityData.YListKeys = []string {}

    return &(helperAddresses.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
// DHCPv6 Helper Address
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. DHCPv6 Helper Address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HelperAddress interface{}

    // DHCPv6 HelperAddress Specific Output Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    OutInterface interface{}

    // DHCPv6 HelperAddress Output Interface. The type is interface{}.
    AnyOutInterface interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + types.AddKeyToken(helperAddress.HelperAddress, "helper-address")
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = types.NewOrderedMap()
    helperAddress.EntityData.Leafs = types.NewOrderedMap()
    helperAddress.EntityData.Leafs.Append("helper-address", types.YLeaf{"HelperAddress", helperAddress.HelperAddress})
    helperAddress.EntityData.Leafs.Append("out-interface", types.YLeaf{"OutInterface", helperAddress.OutInterface})
    helperAddress.EntityData.Leafs.Append("any-out-interface", types.YLeaf{"AnyOutInterface", helperAddress.AnyOutInterface})

    helperAddress.EntityData.YListKeys = []string {"HelperAddress"}

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

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("username", types.YLeaf{"Username", authentication.Username})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Classes
// Table of Class
type Dhcpv6_Profiles_Profile_Proxy_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile_Proxy_Classes_Class.
    Class []*Dhcpv6_Profiles_Profile_Proxy_Classes_Class
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

    classes.EntityData.Children = types.NewOrderedMap()
    classes.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range classes.Class {
        classes.EntityData.Children.Append(types.GetSegmentPath(classes.Class[i]), types.YChild{"Class", classes.Class[i]})
    }
    classes.EntityData.Leafs = types.NewOrderedMap()

    classes.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkAddress interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassName, "class-name")
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Children.Append("helper-addresses", types.YChild{"HelperAddresses", &class.HelperAddresses})
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", class.ClassName})
    class.EntityData.Leafs.Append("link-address", types.YLeaf{"LinkAddress", class.LinkAddress})

    class.EntityData.YListKeys = []string {"ClassName"}

    return &(class.EntityData)
}

// Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the server helper address. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress.
    HelperAddress []*Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress
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

    helperAddresses.EntityData.Children = types.NewOrderedMap()
    helperAddresses.EntityData.Children.Append("helper-address", types.YChild{"HelperAddress", nil})
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children.Append(types.GetSegmentPath(helperAddresses.HelperAddress[i]), types.YChild{"HelperAddress", helperAddresses.HelperAddress[i]})
    }
    helperAddresses.EntityData.Leafs = types.NewOrderedMap()

    helperAddresses.EntityData.YListKeys = []string {}

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
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HelperAddress interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + types.AddKeyToken(helperAddress.VrfName, "vrf-name") + types.AddKeyToken(helperAddress.HelperAddress, "helper-address")
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = types.NewOrderedMap()
    helperAddress.EntityData.Leafs = types.NewOrderedMap()
    helperAddress.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", helperAddress.VrfName})
    helperAddress.EntityData.Leafs.Append("helper-address", types.YLeaf{"HelperAddress", helperAddress.HelperAddress})

    helperAddress.EntityData.YListKeys = []string {"VrfName", "HelperAddress"}

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

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("mac", types.YChild{"Mac", &sessions.Mac})
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

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

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("throttle", types.YChild{"Throttle", &mac.Throttle})
    mac.EntityData.Leafs = types.NewOrderedMap()

    mac.EntityData.YListKeys = []string {}

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

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", throttle.Limit})
    throttle.EntityData.Leafs.Append("request", types.YLeaf{"Request", throttle.Request})
    throttle.EntityData.Leafs.Append("block", types.YLeaf{"Block", throttle.Block})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// Dhcpv6_Profiles_Profile_Server
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    // Client DUID.
    Dhcpv6duid Dhcpv6_Profiles_Profile_Server_Dhcpv6duid

    // Enable aaa dhcpv6 option force-insert.
    AaaServer Dhcpv6_Profiles_Profile_Server_AaaServer

    // DHCPv6 match.
    Options Dhcpv6_Profiles_Profile_Server_Options

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

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Children.Append("sessions", types.YChild{"Sessions", &server.Sessions})
    server.EntityData.Children.Append("dns-servers", types.YChild{"DnsServers", &server.DnsServers})
    server.EntityData.Children.Append("classes", types.YChild{"Classes", &server.Classes})
    server.EntityData.Children.Append("lease", types.YChild{"Lease", &server.Lease})
    server.EntityData.Children.Append("dhcpv6duid", types.YChild{"Dhcpv6duid", &server.Dhcpv6duid})
    server.EntityData.Children.Append("aaa-server", types.YChild{"AaaServer", &server.AaaServer})
    server.EntityData.Children.Append("options", types.YChild{"Options", &server.Options})
    server.EntityData.Children.Append("dhcpv6-options", types.YChild{"Dhcpv6Options", &server.Dhcpv6Options})
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("address-pool", types.YLeaf{"AddressPool", server.AddressPool})
    server.EntityData.Leafs.Append("aftr-name", types.YLeaf{"AftrName", server.AftrName})
    server.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", server.DomainName})
    server.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", server.Preference})
    server.EntityData.Leafs.Append("rapid-commit", types.YLeaf{"RapidCommit", server.RapidCommit})
    server.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", server.Enable})
    server.EntityData.Leafs.Append("prefix-pool", types.YLeaf{"PrefixPool", server.PrefixPool})

    server.EntityData.YListKeys = []string {}

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

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("mac", types.YChild{"Mac", &sessions.Mac})
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

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

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("throttle", types.YChild{"Throttle", &mac.Throttle})
    mac.EntityData.Leafs = types.NewOrderedMap()

    mac.EntityData.YListKeys = []string {}

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

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", throttle.Limit})
    throttle.EntityData.Leafs.Append("request", types.YLeaf{"Request", throttle.Request})
    throttle.EntityData.Leafs.Append("block", types.YLeaf{"Block", throttle.Block})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_DnsServers
// DNS servers
type Dhcpv6_Profiles_Profile_Server_DnsServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server's IPv6 address. The type is one of the following types: slice of
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    dnsServers.EntityData.Children = types.NewOrderedMap()
    dnsServers.EntityData.Leafs = types.NewOrderedMap()
    dnsServers.EntityData.Leafs.Append("dns-server", types.YLeaf{"DnsServer", dnsServers.DnsServer})

    dnsServers.EntityData.YListKeys = []string {}

    return &(dnsServers.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Classes
// Table of Class
type Dhcpv6_Profiles_Profile_Server_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile_Server_Classes_Class.
    Class []*Dhcpv6_Profiles_Profile_Server_Classes_Class
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

    classes.EntityData.Children = types.NewOrderedMap()
    classes.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range classes.Class {
        classes.EntityData.Children.Append(types.GetSegmentPath(classes.Class[i]), types.YChild{"Class", classes.Class[i]})
    }
    classes.EntityData.Leafs = types.NewOrderedMap()

    classes.EntityData.YListKeys = []string {}

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

    // lease.
    Lease Dhcpv6_Profiles_Profile_Server_Classes_Class_Lease
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassName, "class-name")
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Children.Append("dns-servers", types.YChild{"DnsServers", &class.DnsServers})
    class.EntityData.Children.Append("lease", types.YChild{"Lease", &class.Lease})
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", class.ClassName})
    class.EntityData.Leafs.Append("address-pool", types.YLeaf{"AddressPool", class.AddressPool})
    class.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", class.DomainName})
    class.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", class.Preference})
    class.EntityData.Leafs.Append("prefix-pool", types.YLeaf{"PrefixPool", class.PrefixPool})

    class.EntityData.YListKeys = []string {"ClassName"}

    return &(class.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers
// DNS servers
type Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server's IPv6 address. The type is one of the following types: slice of
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    dnsServers.EntityData.Children = types.NewOrderedMap()
    dnsServers.EntityData.Leafs = types.NewOrderedMap()
    dnsServers.EntityData.Leafs.Append("dns-server", types.YLeaf{"DnsServer", dnsServers.DnsServer})

    dnsServers.EntityData.YListKeys = []string {}

    return &(dnsServers.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Classes_Class_Lease
// lease
type Dhcpv6_Profiles_Profile_Server_Classes_Class_Lease struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set string. The type is string.
    Infinite interface{}

    // Days. The type is interface{} with range: 0..365. Units are day.
    Days interface{}

    // Hours. The type is interface{} with range: 0..23. Units are hour.
    Hours interface{}

    // Minutes. The type is interface{} with range: 1..59. Units are minute.
    Minutes interface{}
}

func (lease *Dhcpv6_Profiles_Profile_Server_Classes_Class_Lease) GetEntityData() *types.CommonEntityData {
    lease.EntityData.YFilter = lease.YFilter
    lease.EntityData.YangName = "lease"
    lease.EntityData.BundleName = "cisco_ios_xr"
    lease.EntityData.ParentYangName = "class"
    lease.EntityData.SegmentPath = "lease"
    lease.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lease.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lease.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lease.EntityData.Children = types.NewOrderedMap()
    lease.EntityData.Leafs = types.NewOrderedMap()
    lease.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", lease.Infinite})
    lease.EntityData.Leafs.Append("days", types.YLeaf{"Days", lease.Days})
    lease.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lease.Hours})
    lease.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lease.Minutes})

    lease.EntityData.YListKeys = []string {}

    return &(lease.EntityData)
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

    lease.EntityData.Children = types.NewOrderedMap()
    lease.EntityData.Leafs = types.NewOrderedMap()
    lease.EntityData.Leafs.Append("days", types.YLeaf{"Days", lease.Days})
    lease.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lease.Hours})
    lease.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lease.Minutes})
    lease.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", lease.Infinite})

    lease.EntityData.YListKeys = []string {}

    return &(lease.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Dhcpv6duid
// Client DUID
type Dhcpv6_Profiles_Profile_Server_Dhcpv6duid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of DUID to be allowed. The type is interface{} with range: 1..4.
    AllowedType interface{}
}

func (dhcpv6duid *Dhcpv6_Profiles_Profile_Server_Dhcpv6duid) GetEntityData() *types.CommonEntityData {
    dhcpv6duid.EntityData.YFilter = dhcpv6duid.YFilter
    dhcpv6duid.EntityData.YangName = "dhcpv6duid"
    dhcpv6duid.EntityData.BundleName = "cisco_ios_xr"
    dhcpv6duid.EntityData.ParentYangName = "server"
    dhcpv6duid.EntityData.SegmentPath = "dhcpv6duid"
    dhcpv6duid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6duid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6duid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6duid.EntityData.Children = types.NewOrderedMap()
    dhcpv6duid.EntityData.Leafs = types.NewOrderedMap()
    dhcpv6duid.EntityData.Leafs.Append("allowed-type", types.YLeaf{"AllowedType", dhcpv6duid.AllowedType})

    dhcpv6duid.EntityData.YListKeys = []string {}

    return &(dhcpv6duid.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_AaaServer
// Enable aaa dhcpv6 option force-insert
type Dhcpv6_Profiles_Profile_Server_AaaServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable aaa dhcpv6 option force-insert.
    Dhcpv6Option Dhcpv6_Profiles_Profile_Server_AaaServer_Dhcpv6Option
}

func (aaaServer *Dhcpv6_Profiles_Profile_Server_AaaServer) GetEntityData() *types.CommonEntityData {
    aaaServer.EntityData.YFilter = aaaServer.YFilter
    aaaServer.EntityData.YangName = "aaa-server"
    aaaServer.EntityData.BundleName = "cisco_ios_xr"
    aaaServer.EntityData.ParentYangName = "server"
    aaaServer.EntityData.SegmentPath = "aaa-server"
    aaaServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaaServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaaServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaaServer.EntityData.Children = types.NewOrderedMap()
    aaaServer.EntityData.Children.Append("dhcpv6-option", types.YChild{"Dhcpv6Option", &aaaServer.Dhcpv6Option})
    aaaServer.EntityData.Leafs = types.NewOrderedMap()

    aaaServer.EntityData.YListKeys = []string {}

    return &(aaaServer.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_AaaServer_Dhcpv6Option
// Enable aaa dhcpv6 option force-insert
type Dhcpv6_Profiles_Profile_Server_AaaServer_Dhcpv6Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable aaa dhcpv6 option force-insert. The type is interface{}.
    ForceInsert interface{}
}

func (dhcpv6Option *Dhcpv6_Profiles_Profile_Server_AaaServer_Dhcpv6Option) GetEntityData() *types.CommonEntityData {
    dhcpv6Option.EntityData.YFilter = dhcpv6Option.YFilter
    dhcpv6Option.EntityData.YangName = "dhcpv6-option"
    dhcpv6Option.EntityData.BundleName = "cisco_ios_xr"
    dhcpv6Option.EntityData.ParentYangName = "aaa-server"
    dhcpv6Option.EntityData.SegmentPath = "dhcpv6-option"
    dhcpv6Option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6Option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6Option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6Option.EntityData.Children = types.NewOrderedMap()
    dhcpv6Option.EntityData.Leafs = types.NewOrderedMap()
    dhcpv6Option.EntityData.Leafs.Append("force-insert", types.YLeaf{"ForceInsert", dhcpv6Option.ForceInsert})

    dhcpv6Option.EntityData.YListKeys = []string {}

    return &(dhcpv6Option.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options
// DHCPv6 match
type Dhcpv6_Profiles_Profile_Server_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 match option. The type is slice of
    // Dhcpv6_Profiles_Profile_Server_Options_Option.
    Option []*Dhcpv6_Profiles_Profile_Server_Options_Option
}

func (options *Dhcpv6_Profiles_Profile_Server_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "server"
    options.EntityData.SegmentPath = "options"
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = types.NewOrderedMap()
    options.EntityData.Children.Append("option", types.YChild{"Option", nil})
    for i := range options.Option {
        options.EntityData.Children.Append(types.GetSegmentPath(options.Option[i]), types.YChild{"Option", options.Option[i]})
    }
    options.EntityData.Leafs = types.NewOrderedMap()

    options.EntityData.YListKeys = []string {}

    return &(options.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options_Option
// DHCPv6 match option
type Dhcpv6_Profiles_Profile_Server_Options_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set string. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: 0..4294967295.
    Format interface{}

    // This attribute is a key. Set string. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Value interface{}

    // match enterprise number.
    EnterpriseId Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId

    // match vendor class.
    VendorClass Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass
}

func (option *Dhcpv6_Profiles_Profile_Server_Options_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "options"
    option.EntityData.SegmentPath = "option" + types.AddKeyToken(option.Type, "type") + types.AddKeyToken(option.Format, "format") + types.AddKeyToken(option.Value, "value")
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Children.Append("enterprise-id", types.YChild{"EnterpriseId", &option.EnterpriseId})
    option.EntityData.Children.Append("vendor-class", types.YChild{"VendorClass", &option.VendorClass})
    option.EntityData.Leafs = types.NewOrderedMap()
    option.EntityData.Leafs.Append("type", types.YLeaf{"Type", option.Type})
    option.EntityData.Leafs.Append("format", types.YLeaf{"Format", option.Format})
    option.EntityData.Leafs.Append("value", types.YLeaf{"Value", option.Value})

    option.EntityData.YListKeys = []string {"Type", "Format", "Value"}

    return &(option.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId
// match enterprise number
type Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // defaut action for enterprise number.
    HexEnterpriseId Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_HexEnterpriseId

    // defaut action for enterprise number.
    DefaultEnterpriseId Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_DefaultEnterpriseId
}

func (enterpriseId *Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId) GetEntityData() *types.CommonEntityData {
    enterpriseId.EntityData.YFilter = enterpriseId.YFilter
    enterpriseId.EntityData.YangName = "enterprise-id"
    enterpriseId.EntityData.BundleName = "cisco_ios_xr"
    enterpriseId.EntityData.ParentYangName = "option"
    enterpriseId.EntityData.SegmentPath = "enterprise-id"
    enterpriseId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enterpriseId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enterpriseId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enterpriseId.EntityData.Children = types.NewOrderedMap()
    enterpriseId.EntityData.Children.Append("hex-enterprise-id", types.YChild{"HexEnterpriseId", &enterpriseId.HexEnterpriseId})
    enterpriseId.EntityData.Children.Append("default-enterprise-id", types.YChild{"DefaultEnterpriseId", &enterpriseId.DefaultEnterpriseId})
    enterpriseId.EntityData.Leafs = types.NewOrderedMap()

    enterpriseId.EntityData.YListKeys = []string {}

    return &(enterpriseId.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_HexEnterpriseId
// defaut action for enterprise number
type Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_HexEnterpriseId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Action to be take on match. The type is Action.
    Action interface{}
}

func (hexEnterpriseId *Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_HexEnterpriseId) GetEntityData() *types.CommonEntityData {
    hexEnterpriseId.EntityData.YFilter = hexEnterpriseId.YFilter
    hexEnterpriseId.EntityData.YangName = "hex-enterprise-id"
    hexEnterpriseId.EntityData.BundleName = "cisco_ios_xr"
    hexEnterpriseId.EntityData.ParentYangName = "enterprise-id"
    hexEnterpriseId.EntityData.SegmentPath = "hex-enterprise-id"
    hexEnterpriseId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hexEnterpriseId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hexEnterpriseId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hexEnterpriseId.EntityData.Children = types.NewOrderedMap()
    hexEnterpriseId.EntityData.Leafs = types.NewOrderedMap()
    hexEnterpriseId.EntityData.Leafs.Append("action", types.YLeaf{"Action", hexEnterpriseId.Action})

    hexEnterpriseId.EntityData.YListKeys = []string {}

    return &(hexEnterpriseId.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_DefaultEnterpriseId
// defaut action for enterprise number
type Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_DefaultEnterpriseId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Action to be take on match. The type is Action.
    Action interface{}
}

func (defaultEnterpriseId *Dhcpv6_Profiles_Profile_Server_Options_Option_EnterpriseId_DefaultEnterpriseId) GetEntityData() *types.CommonEntityData {
    defaultEnterpriseId.EntityData.YFilter = defaultEnterpriseId.YFilter
    defaultEnterpriseId.EntityData.YangName = "default-enterprise-id"
    defaultEnterpriseId.EntityData.BundleName = "cisco_ios_xr"
    defaultEnterpriseId.EntityData.ParentYangName = "enterprise-id"
    defaultEnterpriseId.EntityData.SegmentPath = "default-enterprise-id"
    defaultEnterpriseId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultEnterpriseId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultEnterpriseId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultEnterpriseId.EntityData.Children = types.NewOrderedMap()
    defaultEnterpriseId.EntityData.Leafs = types.NewOrderedMap()
    defaultEnterpriseId.EntityData.Leafs.Append("action", types.YLeaf{"Action", defaultEnterpriseId.Action})

    defaultEnterpriseId.EntityData.YListKeys = []string {}

    return &(defaultEnterpriseId.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass
// match vendor class
type Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // string action for vendor number.
    StrVendorClass Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_StrVendorClass

    // default action for enterprise number.
    DefaultVendorClass Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_DefaultVendorClass
}

func (vendorClass *Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass) GetEntityData() *types.CommonEntityData {
    vendorClass.EntityData.YFilter = vendorClass.YFilter
    vendorClass.EntityData.YangName = "vendor-class"
    vendorClass.EntityData.BundleName = "cisco_ios_xr"
    vendorClass.EntityData.ParentYangName = "option"
    vendorClass.EntityData.SegmentPath = "vendor-class"
    vendorClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vendorClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vendorClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vendorClass.EntityData.Children = types.NewOrderedMap()
    vendorClass.EntityData.Children.Append("str-vendor-class", types.YChild{"StrVendorClass", &vendorClass.StrVendorClass})
    vendorClass.EntityData.Children.Append("default-vendor-class", types.YChild{"DefaultVendorClass", &vendorClass.DefaultVendorClass})
    vendorClass.EntityData.Leafs = types.NewOrderedMap()

    vendorClass.EntityData.YListKeys = []string {}

    return &(vendorClass.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_StrVendorClass
// string action for vendor number
type Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_StrVendorClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Action to be take on match. The type is Action.
    Action interface{}
}

func (strVendorClass *Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_StrVendorClass) GetEntityData() *types.CommonEntityData {
    strVendorClass.EntityData.YFilter = strVendorClass.YFilter
    strVendorClass.EntityData.YangName = "str-vendor-class"
    strVendorClass.EntityData.BundleName = "cisco_ios_xr"
    strVendorClass.EntityData.ParentYangName = "vendor-class"
    strVendorClass.EntityData.SegmentPath = "str-vendor-class"
    strVendorClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    strVendorClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    strVendorClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    strVendorClass.EntityData.Children = types.NewOrderedMap()
    strVendorClass.EntityData.Leafs = types.NewOrderedMap()
    strVendorClass.EntityData.Leafs.Append("action", types.YLeaf{"Action", strVendorClass.Action})

    strVendorClass.EntityData.YListKeys = []string {}

    return &(strVendorClass.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_DefaultVendorClass
// default action for enterprise number
type Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_DefaultVendorClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Action to be take on match. The type is Action.
    Action interface{}
}

func (defaultVendorClass *Dhcpv6_Profiles_Profile_Server_Options_Option_VendorClass_DefaultVendorClass) GetEntityData() *types.CommonEntityData {
    defaultVendorClass.EntityData.YFilter = defaultVendorClass.YFilter
    defaultVendorClass.EntityData.YangName = "default-vendor-class"
    defaultVendorClass.EntityData.BundleName = "cisco_ios_xr"
    defaultVendorClass.EntityData.ParentYangName = "vendor-class"
    defaultVendorClass.EntityData.SegmentPath = "default-vendor-class"
    defaultVendorClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVendorClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVendorClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVendorClass.EntityData.Children = types.NewOrderedMap()
    defaultVendorClass.EntityData.Leafs = types.NewOrderedMap()
    defaultVendorClass.EntityData.Leafs.Append("action", types.YLeaf{"Action", defaultVendorClass.Action})

    defaultVendorClass.EntityData.YListKeys = []string {}

    return &(defaultVendorClass.EntityData)
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

    dhcpv6Options.EntityData.Children = types.NewOrderedMap()
    dhcpv6Options.EntityData.Children.Append("vendor-options", types.YChild{"VendorOptions", &dhcpv6Options.VendorOptions})
    dhcpv6Options.EntityData.Leafs = types.NewOrderedMap()

    dhcpv6Options.EntityData.YListKeys = []string {}

    return &(dhcpv6Options.EntityData)
}

// Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions
// Vendor options
type Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set string. The type is string.
    Type interface{}

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

    vendorOptions.EntityData.Children = types.NewOrderedMap()
    vendorOptions.EntityData.Leafs = types.NewOrderedMap()
    vendorOptions.EntityData.Leafs.Append("type", types.YLeaf{"Type", vendorOptions.Type})
    vendorOptions.EntityData.Leafs.Append("vendor-options", types.YLeaf{"VendorOptions", vendorOptions.VendorOptions})

    vendorOptions.EntityData.YListKeys = []string {}

    return &(vendorOptions.EntityData)
}

// Dhcpv6_Interfaces
// Table of Interface
type Dhcpv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Interfaces_Interface.
    Interface []*Dhcpv6_Interfaces_Interface
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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Dhcpv6_Interfaces_Interface
// None
type Dhcpv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &self.Pppoe})
    self.EntityData.Children.Append("proxy", types.YChild{"Proxy", &self.Proxy})
    self.EntityData.Children.Append("base", types.YChild{"Base", &self.Base})
    self.EntityData.Children.Append("server", types.YChild{"Server", &self.Server})
    self.EntityData.Children.Append("relay", types.YChild{"Relay", &self.Relay})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

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

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", pppoe.Profile})

    pppoe.EntityData.YListKeys = []string {}

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

    proxy.EntityData.Children = types.NewOrderedMap()
    proxy.EntityData.Leafs = types.NewOrderedMap()
    proxy.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", proxy.Profile})

    proxy.EntityData.YListKeys = []string {}

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

    base.EntityData.Children = types.NewOrderedMap()
    base.EntityData.Leafs = types.NewOrderedMap()
    base.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", base.Profile})

    base.EntityData.YListKeys = []string {}

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

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", server.Profile})

    server.EntityData.YListKeys = []string {}

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

    relay.EntityData.Children = types.NewOrderedMap()
    relay.EntityData.Leafs = types.NewOrderedMap()
    relay.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", relay.Profile})

    relay.EntityData.YListKeys = []string {}

    return &(relay.EntityData)
}

