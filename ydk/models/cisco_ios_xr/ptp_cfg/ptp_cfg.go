// This module contains a collection of YANG definitions
// for Cisco IOS-XR ptp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ptp: PTP global configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ptp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ptp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ptp-cfg ptp}", reflect.TypeOf(Ptp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ptp-cfg:ptp", reflect.TypeOf(Ptp{}))
}

// Ptp
// PTP global configuration
type Ptp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock class to be used while acquiring phase-lock to a parent clock. Note
    // that this is deprecated and should not be used. The type is interface{}
    // with range: 0..255.
    UncalibratedClockClass interface{}

    // Time-of-day priority. The type is interface{} with range: 1..254. The
    // default value is 100.
    TimeOfDayPriority interface{}

    // Frequency priority. The type is interface{} with range: 1..254. The default
    // value is 254.
    FrequencyPriority interface{}

    // Startup clock class value. The type is interface{} with range: 0..255.
    StartupClockClass interface{}

    // Enable the precision time protocol. The type is interface{}.
    Enable interface{}

    // Clocks with a clock-class higher than the minimum clock class will not be
    // considered for selection as a parent clock. The type is interface{} with
    // range: 0..255.
    MinClockClass interface{}

    // Disable PTP as a source for frequency as only physical layer frequency
    // sources are used. The type is interface{}.
    PhysicalLayerFrequency interface{}

    // Freerun clock class value. The type is interface{} with range: 0..255.
    FreerunClockClass interface{}

    // PTP local clock configuration.
    Clock Ptp_Clock

    // Table for profile configuration.
    Profiles Ptp_Profiles

    // UTC offset configuration.
    UtcOffset Ptp_UtcOffset

    // PTP logging configuration.
    Logging Ptp_Logging

    // Clock class to be used while acquiring phase-lock to a parent clock.
    UncalibratedClockClass2 Ptp_UncalibratedClockClass2

    // Transparent clock configuration.
    TransparentClock Ptp_TransparentClock

    // PTP virtual port configuration.
    VirtualPort Ptp_VirtualPort
}

func (ptp *Ptp) GetEntityData() *types.CommonEntityData {
    ptp.EntityData.YFilter = ptp.YFilter
    ptp.EntityData.YangName = "ptp"
    ptp.EntityData.BundleName = "cisco_ios_xr"
    ptp.EntityData.ParentYangName = "Cisco-IOS-XR-ptp-cfg"
    ptp.EntityData.SegmentPath = "Cisco-IOS-XR-ptp-cfg:ptp"
    ptp.EntityData.AbsolutePath = ptp.EntityData.SegmentPath
    ptp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptp.EntityData.Children = types.NewOrderedMap()
    ptp.EntityData.Children.Append("clock", types.YChild{"Clock", &ptp.Clock})
    ptp.EntityData.Children.Append("profiles", types.YChild{"Profiles", &ptp.Profiles})
    ptp.EntityData.Children.Append("utc-offset", types.YChild{"UtcOffset", &ptp.UtcOffset})
    ptp.EntityData.Children.Append("logging", types.YChild{"Logging", &ptp.Logging})
    ptp.EntityData.Children.Append("uncalibrated-clock-class2", types.YChild{"UncalibratedClockClass2", &ptp.UncalibratedClockClass2})
    ptp.EntityData.Children.Append("transparent-clock", types.YChild{"TransparentClock", &ptp.TransparentClock})
    ptp.EntityData.Children.Append("virtual-port", types.YChild{"VirtualPort", &ptp.VirtualPort})
    ptp.EntityData.Leafs = types.NewOrderedMap()
    ptp.EntityData.Leafs.Append("uncalibrated-clock-class", types.YLeaf{"UncalibratedClockClass", ptp.UncalibratedClockClass})
    ptp.EntityData.Leafs.Append("time-of-day-priority", types.YLeaf{"TimeOfDayPriority", ptp.TimeOfDayPriority})
    ptp.EntityData.Leafs.Append("frequency-priority", types.YLeaf{"FrequencyPriority", ptp.FrequencyPriority})
    ptp.EntityData.Leafs.Append("startup-clock-class", types.YLeaf{"StartupClockClass", ptp.StartupClockClass})
    ptp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ptp.Enable})
    ptp.EntityData.Leafs.Append("min-clock-class", types.YLeaf{"MinClockClass", ptp.MinClockClass})
    ptp.EntityData.Leafs.Append("physical-layer-frequency", types.YLeaf{"PhysicalLayerFrequency", ptp.PhysicalLayerFrequency})
    ptp.EntityData.Leafs.Append("freerun-clock-class", types.YLeaf{"FreerunClockClass", ptp.FreerunClockClass})

    ptp.EntityData.YListKeys = []string {}

    return &(ptp.EntityData)
}

// Ptp_Clock
// PTP local clock configuration
type Ptp_Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local clock timescale. The type is PtpTimescale.
    Timescale interface{}

    // Local clock domain. The type is interface{} with range: 0..255. The default
    // value is 0.
    Domain interface{}

    // Local clock priority2. The type is interface{} with range: 0..255. The
    // default value is 128.
    Priority2 interface{}

    // Local clock time source. The type is PtpTimeSource.
    TimeSource interface{}

    // Local clock priority1. The type is interface{} with range: 0..255. The
    // default value is 128.
    Priority1 interface{}

    // Local clock class. The type is interface{} with range: 0..255. The default
    // value is 0.
    ClockClass interface{}

    // Local clock PTP profile.
    Profile Ptp_Clock_Profile

    // Local clock identity.
    Identity Ptp_Clock_Identity
}

func (clock *Ptp_Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "ptp"
    clock.EntityData.SegmentPath = "clock"
    clock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/" + clock.EntityData.SegmentPath
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("profile", types.YChild{"Profile", &clock.Profile})
    clock.EntityData.Children.Append("identity", types.YChild{"Identity", &clock.Identity})
    clock.EntityData.Leafs = types.NewOrderedMap()
    clock.EntityData.Leafs.Append("timescale", types.YLeaf{"Timescale", clock.Timescale})
    clock.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", clock.Domain})
    clock.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", clock.Priority2})
    clock.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", clock.TimeSource})
    clock.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", clock.Priority1})
    clock.EntityData.Leafs.Append("clock-class", types.YLeaf{"ClockClass", clock.ClockClass})

    clock.EntityData.YListKeys = []string {}

    return &(clock.EntityData)
}

// Ptp_Clock_Profile
// Local clock PTP profile
type Ptp_Clock_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock profile. The type is PtpClockProfile. The default value is default.
    ClockProfile interface{}

    // Telecom clock type. The type is PtpTelecomClock.
    TelecomClockType interface{}
}

func (profile *Ptp_Clock_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "clock"
    profile.EntityData.SegmentPath = "profile"
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/clock/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("clock-profile", types.YLeaf{"ClockProfile", profile.ClockProfile})
    profile.EntityData.Leafs.Append("telecom-clock-type", types.YLeaf{"TelecomClockType", profile.TelecomClockType})

    profile.EntityData.YListKeys = []string {}

    return &(profile.EntityData)
}

// Ptp_Clock_Identity
// Local clock identity
type Ptp_Clock_Identity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock identity type. The type is PtpClockId. The default value is
    // router-mac.
    ClockIdType interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // EUI-64 number. The type is string.
    Eui interface{}
}

func (identity *Ptp_Clock_Identity) GetEntityData() *types.CommonEntityData {
    identity.EntityData.YFilter = identity.YFilter
    identity.EntityData.YangName = "identity"
    identity.EntityData.BundleName = "cisco_ios_xr"
    identity.EntityData.ParentYangName = "clock"
    identity.EntityData.SegmentPath = "identity"
    identity.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/clock/" + identity.EntityData.SegmentPath
    identity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identity.EntityData.Children = types.NewOrderedMap()
    identity.EntityData.Leafs = types.NewOrderedMap()
    identity.EntityData.Leafs.Append("clock-id-type", types.YLeaf{"ClockIdType", identity.ClockIdType})
    identity.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", identity.MacAddress})
    identity.EntityData.Leafs.Append("eui", types.YLeaf{"Eui", identity.Eui})

    identity.EntityData.YListKeys = []string {}

    return &(identity.EntityData)
}

// Ptp_Profiles
// Table for profile configuration
type Ptp_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile configuration. The type is slice of Ptp_Profiles_Profile.
    Profile []*Ptp_Profiles_Profile
}

func (profiles *Ptp_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "ptp"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/" + profiles.EntityData.SegmentPath
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

// Ptp_Profiles_Profile
// Profile configuration
type Ptp_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Profile. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Sync unicast grant duration, in seconds. The type is interface{} with
    // range: 60..1000. Units are second. The default value is 300.
    SyncGrantDuration interface{}

    // General COS. The type is interface{} with range: 0..7. The default value is
    // 6.
    GeneralCos interface{}

    // Sync timeout, in milliseconds. The type is interface{} with range:
    // 100..100000. Units are millisecond. The default value is 5000.
    SyncTimeout interface{}

    // Transport. The type is PtpEncap. The default value is ipv4.
    Transport interface{}

    // Announce Timeout. The type is interface{} with range: 2..10. The default
    // value is 3.
    AnnounceTimeout interface{}

    // COS. The type is interface{} with range: 0..7. The default value is 6.
    Cos interface{}

    // IPv4 TTL. The type is interface{} with range: 1..255. The default value is
    // 255.
    Ipv4ttl interface{}

    // Port state restriction. The type is PtpPortState. The default value is any.
    PortState interface{}

    // Delay-Response timeout, in milliseconds. The type is interface{} with
    // range: 100..100000. Units are millisecond. The default value is 5000.
    DelayResponseTimeout interface{}

    // Delay-Response unicast grant duration, in seconds. The type is interface{}
    // with range: 60..1000. Units are second. The default value is 300.
    DelayResponseGrantDuration interface{}

    // Event COS. The type is interface{} with range: 0..7. The default value is
    // 6.
    EventCos interface{}

    // DSCP. The type is interface{} with range: 0..63. The default value is 46.
    Dscp interface{}

    // IPv6 Hop Limit. The type is interface{} with range: 1..255. The default
    // value is 255.
    Ipv6HopLimit interface{}

    // General DSCP. The type is interface{} with range: 0..63. The default value
    // is 46.
    GeneralDscp interface{}

    // Clock Operation. The type is PtpClockOperation. The default value is
    // two-step.
    ClockOperation interface{}

    // Announce unicast grant duration, in seconds. The type is interface{} with
    // range: 60..1000. Units are second. The default value is 300.
    AnnounceGrantDuration interface{}

    // Invalid unicast grant request response. The type is
    // PtpInvalidUnicastGrantRequestResponse. The default value is reduce.
    UnicastGrantInvalidRequest interface{}

    // Event DSCP. The type is interface{} with range: 0..63. The default value is
    // 46.
    EventDscp interface{}

    // Announce interval.
    AnnounceInterval Ptp_Profiles_Profile_AnnounceInterval

    // Table for interop configuration.
    Interop Ptp_Profiles_Profile_Interop

    // Source IPv4 Address.
    SourceIpv4Address Ptp_Profiles_Profile_SourceIpv4Address

    // Table for slave configuration.
    Slaves Ptp_Profiles_Profile_Slaves

    // Sync interval.
    SyncInterval Ptp_Profiles_Profile_SyncInterval

    // Table for master configuration.
    Masters Ptp_Profiles_Profile_Masters

    // Communication model.
    Communication Ptp_Profiles_Profile_Communication

    // Minimum delay request interval.
    DelayRequestMinimumInterval Ptp_Profiles_Profile_DelayRequestMinimumInterval

    // Source IPv6 Address.
    SourceIpv6Address Ptp_Profiles_Profile_SourceIpv6Address
}

func (profile *Ptp_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("announce-interval", types.YChild{"AnnounceInterval", &profile.AnnounceInterval})
    profile.EntityData.Children.Append("interop", types.YChild{"Interop", &profile.Interop})
    profile.EntityData.Children.Append("source-ipv4-address", types.YChild{"SourceIpv4Address", &profile.SourceIpv4Address})
    profile.EntityData.Children.Append("slaves", types.YChild{"Slaves", &profile.Slaves})
    profile.EntityData.Children.Append("sync-interval", types.YChild{"SyncInterval", &profile.SyncInterval})
    profile.EntityData.Children.Append("masters", types.YChild{"Masters", &profile.Masters})
    profile.EntityData.Children.Append("communication", types.YChild{"Communication", &profile.Communication})
    profile.EntityData.Children.Append("delay-request-minimum-interval", types.YChild{"DelayRequestMinimumInterval", &profile.DelayRequestMinimumInterval})
    profile.EntityData.Children.Append("source-ipv6-address", types.YChild{"SourceIpv6Address", &profile.SourceIpv6Address})
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})
    profile.EntityData.Leafs.Append("sync-grant-duration", types.YLeaf{"SyncGrantDuration", profile.SyncGrantDuration})
    profile.EntityData.Leafs.Append("general-cos", types.YLeaf{"GeneralCos", profile.GeneralCos})
    profile.EntityData.Leafs.Append("sync-timeout", types.YLeaf{"SyncTimeout", profile.SyncTimeout})
    profile.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", profile.Transport})
    profile.EntityData.Leafs.Append("announce-timeout", types.YLeaf{"AnnounceTimeout", profile.AnnounceTimeout})
    profile.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", profile.Cos})
    profile.EntityData.Leafs.Append("ipv4ttl", types.YLeaf{"Ipv4ttl", profile.Ipv4ttl})
    profile.EntityData.Leafs.Append("port-state", types.YLeaf{"PortState", profile.PortState})
    profile.EntityData.Leafs.Append("delay-response-timeout", types.YLeaf{"DelayResponseTimeout", profile.DelayResponseTimeout})
    profile.EntityData.Leafs.Append("delay-response-grant-duration", types.YLeaf{"DelayResponseGrantDuration", profile.DelayResponseGrantDuration})
    profile.EntityData.Leafs.Append("event-cos", types.YLeaf{"EventCos", profile.EventCos})
    profile.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", profile.Dscp})
    profile.EntityData.Leafs.Append("ipv6-hop-limit", types.YLeaf{"Ipv6HopLimit", profile.Ipv6HopLimit})
    profile.EntityData.Leafs.Append("general-dscp", types.YLeaf{"GeneralDscp", profile.GeneralDscp})
    profile.EntityData.Leafs.Append("clock-operation", types.YLeaf{"ClockOperation", profile.ClockOperation})
    profile.EntityData.Leafs.Append("announce-grant-duration", types.YLeaf{"AnnounceGrantDuration", profile.AnnounceGrantDuration})
    profile.EntityData.Leafs.Append("unicast-grant-invalid-request", types.YLeaf{"UnicastGrantInvalidRequest", profile.UnicastGrantInvalidRequest})
    profile.EntityData.Leafs.Append("event-dscp", types.YLeaf{"EventDscp", profile.EventDscp})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
}

// Ptp_Profiles_Profile_AnnounceInterval
// Announce interval
type Ptp_Profiles_Profile_AnnounceInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval or Frequency. The type is PtpTime. The default value is interval.
    TimeType interface{}

    // Time Period. The type is PtpTimePeriod. The default value is 2.
    TimePeriod interface{}
}

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetEntityData() *types.CommonEntityData {
    announceInterval.EntityData.YFilter = announceInterval.YFilter
    announceInterval.EntityData.YangName = "announce-interval"
    announceInterval.EntityData.BundleName = "cisco_ios_xr"
    announceInterval.EntityData.ParentYangName = "profile"
    announceInterval.EntityData.SegmentPath = "announce-interval"
    announceInterval.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + announceInterval.EntityData.SegmentPath
    announceInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceInterval.EntityData.Children = types.NewOrderedMap()
    announceInterval.EntityData.Leafs = types.NewOrderedMap()
    announceInterval.EntityData.Leafs.Append("time-type", types.YLeaf{"TimeType", announceInterval.TimeType})
    announceInterval.EntityData.Leafs.Append("time-period", types.YLeaf{"TimePeriod", announceInterval.TimePeriod})

    announceInterval.EntityData.YListKeys = []string {}

    return &(announceInterval.EntityData)
}

// Ptp_Profiles_Profile_Interop
// Table for interop configuration
type Ptp_Profiles_Profile_Interop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile to interoperate with. The type is PtpClockProfile.
    Profile interface{}

    // Domain number of the peer clock. The type is interface{} with range:
    // 0..255.
    Domain interface{}

    // Iteroperation configuration to be used on egress.
    EgressConversion Ptp_Profiles_Profile_Interop_EgressConversion

    // Iteroperation configuration to be used on ingress.
    IngressConversion Ptp_Profiles_Profile_Interop_IngressConversion
}

func (interop *Ptp_Profiles_Profile_Interop) GetEntityData() *types.CommonEntityData {
    interop.EntityData.YFilter = interop.YFilter
    interop.EntityData.YangName = "interop"
    interop.EntityData.BundleName = "cisco_ios_xr"
    interop.EntityData.ParentYangName = "profile"
    interop.EntityData.SegmentPath = "interop"
    interop.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + interop.EntityData.SegmentPath
    interop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interop.EntityData.Children = types.NewOrderedMap()
    interop.EntityData.Children.Append("egress-conversion", types.YChild{"EgressConversion", &interop.EgressConversion})
    interop.EntityData.Children.Append("ingress-conversion", types.YChild{"IngressConversion", &interop.IngressConversion})
    interop.EntityData.Leafs = types.NewOrderedMap()
    interop.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", interop.Profile})
    interop.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", interop.Domain})

    interop.EntityData.YListKeys = []string {}

    return &(interop.EntityData)
}

// Ptp_Profiles_Profile_Interop_EgressConversion
// Iteroperation configuration to be used on
// egress
type Ptp_Profiles_Profile_Interop_EgressConversion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock Accuracy value to use for the peer clock. The type is interface{}
    // with range: 0..254.
    ClockAccuracy interface{}

    // Priority2 value to use for the peer clock. The type is interface{} with
    // range: 0..255.
    Priority2 interface{}

    // Default clock class to use when a more specific mapping is not available.
    // The type is interface{} with range: 0..255.
    ClockClassDefault interface{}

    // OSLV value to use for the peer clock. The type is interface{} with range:
    // 0..65535.
    OffsetScaledLogVariance interface{}

    // Priority1 value to use for the peer clock. The type is interface{} with
    // range: 0..255.
    Priority1 interface{}

    // Table for specific mappings for given clock class values.
    ClockClassMappings Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings
}

func (egressConversion *Ptp_Profiles_Profile_Interop_EgressConversion) GetEntityData() *types.CommonEntityData {
    egressConversion.EntityData.YFilter = egressConversion.YFilter
    egressConversion.EntityData.YangName = "egress-conversion"
    egressConversion.EntityData.BundleName = "cisco_ios_xr"
    egressConversion.EntityData.ParentYangName = "interop"
    egressConversion.EntityData.SegmentPath = "egress-conversion"
    egressConversion.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/interop/" + egressConversion.EntityData.SegmentPath
    egressConversion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressConversion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressConversion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressConversion.EntityData.Children = types.NewOrderedMap()
    egressConversion.EntityData.Children.Append("clock-class-mappings", types.YChild{"ClockClassMappings", &egressConversion.ClockClassMappings})
    egressConversion.EntityData.Leafs = types.NewOrderedMap()
    egressConversion.EntityData.Leafs.Append("clock-accuracy", types.YLeaf{"ClockAccuracy", egressConversion.ClockAccuracy})
    egressConversion.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", egressConversion.Priority2})
    egressConversion.EntityData.Leafs.Append("clock-class-default", types.YLeaf{"ClockClassDefault", egressConversion.ClockClassDefault})
    egressConversion.EntityData.Leafs.Append("offset-scaled-log-variance", types.YLeaf{"OffsetScaledLogVariance", egressConversion.OffsetScaledLogVariance})
    egressConversion.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", egressConversion.Priority1})

    egressConversion.EntityData.YListKeys = []string {}

    return &(egressConversion.EntityData)
}

// Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings
// Table for specific mappings for given clock
// class values
type Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mapping for a given clock class value. The type is slice of
    // Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings_ClockClassMapping.
    ClockClassMapping []*Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings_ClockClassMapping
}

func (clockClassMappings *Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings) GetEntityData() *types.CommonEntityData {
    clockClassMappings.EntityData.YFilter = clockClassMappings.YFilter
    clockClassMappings.EntityData.YangName = "clock-class-mappings"
    clockClassMappings.EntityData.BundleName = "cisco_ios_xr"
    clockClassMappings.EntityData.ParentYangName = "egress-conversion"
    clockClassMappings.EntityData.SegmentPath = "clock-class-mappings"
    clockClassMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/interop/egress-conversion/" + clockClassMappings.EntityData.SegmentPath
    clockClassMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockClassMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockClassMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockClassMappings.EntityData.Children = types.NewOrderedMap()
    clockClassMappings.EntityData.Children.Append("clock-class-mapping", types.YChild{"ClockClassMapping", nil})
    for i := range clockClassMappings.ClockClassMapping {
        clockClassMappings.EntityData.Children.Append(types.GetSegmentPath(clockClassMappings.ClockClassMapping[i]), types.YChild{"ClockClassMapping", clockClassMappings.ClockClassMapping[i]})
    }
    clockClassMappings.EntityData.Leafs = types.NewOrderedMap()

    clockClassMappings.EntityData.YListKeys = []string {}

    return &(clockClassMappings.EntityData)
}

// Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings_ClockClassMapping
// Mapping for a given clock class value
type Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings_ClockClassMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Clock Class to map from. The type is interface{}
    // with range: 0..255.
    ClockClassFrom interface{}

    // Clock class to map to. The type is interface{} with range: 0..255. This
    // attribute is mandatory.
    ClockClassTo interface{}
}

func (clockClassMapping *Ptp_Profiles_Profile_Interop_EgressConversion_ClockClassMappings_ClockClassMapping) GetEntityData() *types.CommonEntityData {
    clockClassMapping.EntityData.YFilter = clockClassMapping.YFilter
    clockClassMapping.EntityData.YangName = "clock-class-mapping"
    clockClassMapping.EntityData.BundleName = "cisco_ios_xr"
    clockClassMapping.EntityData.ParentYangName = "clock-class-mappings"
    clockClassMapping.EntityData.SegmentPath = "clock-class-mapping" + types.AddKeyToken(clockClassMapping.ClockClassFrom, "clock-class-from")
    clockClassMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/interop/egress-conversion/clock-class-mappings/" + clockClassMapping.EntityData.SegmentPath
    clockClassMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockClassMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockClassMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockClassMapping.EntityData.Children = types.NewOrderedMap()
    clockClassMapping.EntityData.Leafs = types.NewOrderedMap()
    clockClassMapping.EntityData.Leafs.Append("clock-class-from", types.YLeaf{"ClockClassFrom", clockClassMapping.ClockClassFrom})
    clockClassMapping.EntityData.Leafs.Append("clock-class-to", types.YLeaf{"ClockClassTo", clockClassMapping.ClockClassTo})

    clockClassMapping.EntityData.YListKeys = []string {"ClockClassFrom"}

    return &(clockClassMapping.EntityData)
}

// Ptp_Profiles_Profile_Interop_IngressConversion
// Iteroperation configuration to be used on
// ingress
type Ptp_Profiles_Profile_Interop_IngressConversion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock Accuracy value to use for the peer clock. The type is interface{}
    // with range: 0..254.
    ClockAccuracy interface{}

    // Priority2 value to use for the peer clock. The type is interface{} with
    // range: 0..255.
    Priority2 interface{}

    // Default clock class to use when a more specific mapping is not available.
    // The type is interface{} with range: 0..255.
    ClockClassDefault interface{}

    // OSLV value to use for the peer clock. The type is interface{} with range:
    // 0..65535.
    OffsetScaledLogVariance interface{}

    // Priority1 value to use for the peer clock. The type is interface{} with
    // range: 0..255.
    Priority1 interface{}

    // Table for specific mappings for given clock class values.
    ClockClassMappings Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings
}

func (ingressConversion *Ptp_Profiles_Profile_Interop_IngressConversion) GetEntityData() *types.CommonEntityData {
    ingressConversion.EntityData.YFilter = ingressConversion.YFilter
    ingressConversion.EntityData.YangName = "ingress-conversion"
    ingressConversion.EntityData.BundleName = "cisco_ios_xr"
    ingressConversion.EntityData.ParentYangName = "interop"
    ingressConversion.EntityData.SegmentPath = "ingress-conversion"
    ingressConversion.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/interop/" + ingressConversion.EntityData.SegmentPath
    ingressConversion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressConversion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressConversion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressConversion.EntityData.Children = types.NewOrderedMap()
    ingressConversion.EntityData.Children.Append("clock-class-mappings", types.YChild{"ClockClassMappings", &ingressConversion.ClockClassMappings})
    ingressConversion.EntityData.Leafs = types.NewOrderedMap()
    ingressConversion.EntityData.Leafs.Append("clock-accuracy", types.YLeaf{"ClockAccuracy", ingressConversion.ClockAccuracy})
    ingressConversion.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", ingressConversion.Priority2})
    ingressConversion.EntityData.Leafs.Append("clock-class-default", types.YLeaf{"ClockClassDefault", ingressConversion.ClockClassDefault})
    ingressConversion.EntityData.Leafs.Append("offset-scaled-log-variance", types.YLeaf{"OffsetScaledLogVariance", ingressConversion.OffsetScaledLogVariance})
    ingressConversion.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", ingressConversion.Priority1})

    ingressConversion.EntityData.YListKeys = []string {}

    return &(ingressConversion.EntityData)
}

// Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings
// Table for specific mappings for given clock
// class values
type Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mapping for a given clock class value. The type is slice of
    // Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings_ClockClassMapping.
    ClockClassMapping []*Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings_ClockClassMapping
}

func (clockClassMappings *Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings) GetEntityData() *types.CommonEntityData {
    clockClassMappings.EntityData.YFilter = clockClassMappings.YFilter
    clockClassMappings.EntityData.YangName = "clock-class-mappings"
    clockClassMappings.EntityData.BundleName = "cisco_ios_xr"
    clockClassMappings.EntityData.ParentYangName = "ingress-conversion"
    clockClassMappings.EntityData.SegmentPath = "clock-class-mappings"
    clockClassMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/interop/ingress-conversion/" + clockClassMappings.EntityData.SegmentPath
    clockClassMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockClassMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockClassMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockClassMappings.EntityData.Children = types.NewOrderedMap()
    clockClassMappings.EntityData.Children.Append("clock-class-mapping", types.YChild{"ClockClassMapping", nil})
    for i := range clockClassMappings.ClockClassMapping {
        clockClassMappings.EntityData.Children.Append(types.GetSegmentPath(clockClassMappings.ClockClassMapping[i]), types.YChild{"ClockClassMapping", clockClassMappings.ClockClassMapping[i]})
    }
    clockClassMappings.EntityData.Leafs = types.NewOrderedMap()

    clockClassMappings.EntityData.YListKeys = []string {}

    return &(clockClassMappings.EntityData)
}

// Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings_ClockClassMapping
// Mapping for a given clock class value
type Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings_ClockClassMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Clock Class to map from. The type is interface{}
    // with range: 0..255.
    ClockClassFrom interface{}

    // Clock class to map to. The type is interface{} with range: 0..255. This
    // attribute is mandatory.
    ClockClassTo interface{}
}

func (clockClassMapping *Ptp_Profiles_Profile_Interop_IngressConversion_ClockClassMappings_ClockClassMapping) GetEntityData() *types.CommonEntityData {
    clockClassMapping.EntityData.YFilter = clockClassMapping.YFilter
    clockClassMapping.EntityData.YangName = "clock-class-mapping"
    clockClassMapping.EntityData.BundleName = "cisco_ios_xr"
    clockClassMapping.EntityData.ParentYangName = "clock-class-mappings"
    clockClassMapping.EntityData.SegmentPath = "clock-class-mapping" + types.AddKeyToken(clockClassMapping.ClockClassFrom, "clock-class-from")
    clockClassMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/interop/ingress-conversion/clock-class-mappings/" + clockClassMapping.EntityData.SegmentPath
    clockClassMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockClassMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockClassMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockClassMapping.EntityData.Children = types.NewOrderedMap()
    clockClassMapping.EntityData.Leafs = types.NewOrderedMap()
    clockClassMapping.EntityData.Leafs.Append("clock-class-from", types.YLeaf{"ClockClassFrom", clockClassMapping.ClockClassFrom})
    clockClassMapping.EntityData.Leafs.Append("clock-class-to", types.YLeaf{"ClockClassTo", clockClassMapping.ClockClassTo})

    clockClassMapping.EntityData.YListKeys = []string {"ClockClassFrom"}

    return &(clockClassMapping.EntityData)
}

// Ptp_Profiles_Profile_SourceIpv4Address
// Source IPv4 Address
type Ptp_Profiles_Profile_SourceIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable source IP address. The type is bool.
    Enable interface{}

    // Source IP address to use. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceIp interface{}
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetEntityData() *types.CommonEntityData {
    sourceIpv4Address.EntityData.YFilter = sourceIpv4Address.YFilter
    sourceIpv4Address.EntityData.YangName = "source-ipv4-address"
    sourceIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    sourceIpv4Address.EntityData.ParentYangName = "profile"
    sourceIpv4Address.EntityData.SegmentPath = "source-ipv4-address"
    sourceIpv4Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + sourceIpv4Address.EntityData.SegmentPath
    sourceIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceIpv4Address.EntityData.Children = types.NewOrderedMap()
    sourceIpv4Address.EntityData.Leafs = types.NewOrderedMap()
    sourceIpv4Address.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", sourceIpv4Address.Enable})
    sourceIpv4Address.EntityData.Leafs.Append("source-ip", types.YLeaf{"SourceIp", sourceIpv4Address.SourceIp})

    sourceIpv4Address.EntityData.YListKeys = []string {}

    return &(sourceIpv4Address.EntityData)
}

// Ptp_Profiles_Profile_Slaves
// Table for slave configuration
type Ptp_Profiles_Profile_Slaves struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slave configuration. The type is slice of
    // Ptp_Profiles_Profile_Slaves_Slave.
    Slave []*Ptp_Profiles_Profile_Slaves_Slave
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetEntityData() *types.CommonEntityData {
    slaves.EntityData.YFilter = slaves.YFilter
    slaves.EntityData.YangName = "slaves"
    slaves.EntityData.BundleName = "cisco_ios_xr"
    slaves.EntityData.ParentYangName = "profile"
    slaves.EntityData.SegmentPath = "slaves"
    slaves.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + slaves.EntityData.SegmentPath
    slaves.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaves.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaves.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaves.EntityData.Children = types.NewOrderedMap()
    slaves.EntityData.Children.Append("slave", types.YChild{"Slave", nil})
    for i := range slaves.Slave {
        slaves.EntityData.Children.Append(types.GetSegmentPath(slaves.Slave[i]), types.YChild{"Slave", slaves.Slave[i]})
    }
    slaves.EntityData.Leafs = types.NewOrderedMap()

    slaves.EntityData.YListKeys = []string {}

    return &(slaves.EntityData)
}

// Ptp_Profiles_Profile_Slaves_Slave
// Slave configuration
type Ptp_Profiles_Profile_Slaves_Slave struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slave Transport Type. The type is PtpEncap.
    Transport interface{}

    // ethernet. The type is slice of Ptp_Profiles_Profile_Slaves_Slave_Ethernet.
    Ethernet []*Ptp_Profiles_Profile_Slaves_Slave_Ethernet

    // ipv4 or ipv6. The type is slice of
    // Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6.
    Ipv4OrIpv6 []*Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetEntityData() *types.CommonEntityData {
    slave.EntityData.YFilter = slave.YFilter
    slave.EntityData.YangName = "slave"
    slave.EntityData.BundleName = "cisco_ios_xr"
    slave.EntityData.ParentYangName = "slaves"
    slave.EntityData.SegmentPath = "slave" + types.AddKeyToken(slave.Transport, "transport")
    slave.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/slaves/" + slave.EntityData.SegmentPath
    slave.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slave.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slave.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slave.EntityData.Children = types.NewOrderedMap()
    slave.EntityData.Children.Append("ethernet", types.YChild{"Ethernet", nil})
    for i := range slave.Ethernet {
        slave.EntityData.Children.Append(types.GetSegmentPath(slave.Ethernet[i]), types.YChild{"Ethernet", slave.Ethernet[i]})
    }
    slave.EntityData.Children.Append("ipv4-or-ipv6", types.YChild{"Ipv4OrIpv6", nil})
    for i := range slave.Ipv4OrIpv6 {
        slave.EntityData.Children.Append(types.GetSegmentPath(slave.Ipv4OrIpv6[i]), types.YChild{"Ipv4OrIpv6", slave.Ipv4OrIpv6[i]})
    }
    slave.EntityData.Leafs = types.NewOrderedMap()
    slave.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", slave.Transport})

    slave.EntityData.YListKeys = []string {"Transport"}

    return &(slave.EntityData)
}

// Ptp_Profiles_Profile_Slaves_Slave_Ethernet
// ethernet
type Ptp_Profiles_Profile_Slaves_Slave_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slave MAC Address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SlaveMacAddress interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "slave"
    ethernet.EntityData.SegmentPath = "ethernet" + types.AddKeyToken(ethernet.SlaveMacAddress, "slave-mac-address")
    ethernet.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/slaves/slave/" + ethernet.EntityData.SegmentPath
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Leafs = types.NewOrderedMap()
    ethernet.EntityData.Leafs.Append("slave-mac-address", types.YLeaf{"SlaveMacAddress", ethernet.SlaveMacAddress})
    ethernet.EntityData.Leafs.Append("non-negotiated", types.YLeaf{"NonNegotiated", ethernet.NonNegotiated})

    ethernet.EntityData.YListKeys = []string {"SlaveMacAddress"}

    return &(ethernet.EntityData)
}

// Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6
// ipv4 or ipv6
type Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slave IP Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SlaveIpAddress interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetEntityData() *types.CommonEntityData {
    ipv4OrIpv6.EntityData.YFilter = ipv4OrIpv6.YFilter
    ipv4OrIpv6.EntityData.YangName = "ipv4-or-ipv6"
    ipv4OrIpv6.EntityData.BundleName = "cisco_ios_xr"
    ipv4OrIpv6.EntityData.ParentYangName = "slave"
    ipv4OrIpv6.EntityData.SegmentPath = "ipv4-or-ipv6" + types.AddKeyToken(ipv4OrIpv6.SlaveIpAddress, "slave-ip-address")
    ipv4OrIpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/slaves/slave/" + ipv4OrIpv6.EntityData.SegmentPath
    ipv4OrIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4OrIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4OrIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4OrIpv6.EntityData.Children = types.NewOrderedMap()
    ipv4OrIpv6.EntityData.Leafs = types.NewOrderedMap()
    ipv4OrIpv6.EntityData.Leafs.Append("slave-ip-address", types.YLeaf{"SlaveIpAddress", ipv4OrIpv6.SlaveIpAddress})
    ipv4OrIpv6.EntityData.Leafs.Append("non-negotiated", types.YLeaf{"NonNegotiated", ipv4OrIpv6.NonNegotiated})

    ipv4OrIpv6.EntityData.YListKeys = []string {"SlaveIpAddress"}

    return &(ipv4OrIpv6.EntityData)
}

// Ptp_Profiles_Profile_SyncInterval
// Sync interval
type Ptp_Profiles_Profile_SyncInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval or Frequency. The type is PtpTime. The default value is interval.
    TimeType interface{}

    // Time Period. The type is PtpTimePeriod. The default value is 1.
    TimePeriod interface{}
}

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetEntityData() *types.CommonEntityData {
    syncInterval.EntityData.YFilter = syncInterval.YFilter
    syncInterval.EntityData.YangName = "sync-interval"
    syncInterval.EntityData.BundleName = "cisco_ios_xr"
    syncInterval.EntityData.ParentYangName = "profile"
    syncInterval.EntityData.SegmentPath = "sync-interval"
    syncInterval.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + syncInterval.EntityData.SegmentPath
    syncInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncInterval.EntityData.Children = types.NewOrderedMap()
    syncInterval.EntityData.Leafs = types.NewOrderedMap()
    syncInterval.EntityData.Leafs.Append("time-type", types.YLeaf{"TimeType", syncInterval.TimeType})
    syncInterval.EntityData.Leafs.Append("time-period", types.YLeaf{"TimePeriod", syncInterval.TimePeriod})

    syncInterval.EntityData.YListKeys = []string {}

    return &(syncInterval.EntityData)
}

// Ptp_Profiles_Profile_Masters
// Table for master configuration
type Ptp_Profiles_Profile_Masters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Master configuration. The type is slice of
    // Ptp_Profiles_Profile_Masters_Master.
    Master []*Ptp_Profiles_Profile_Masters_Master
}

func (masters *Ptp_Profiles_Profile_Masters) GetEntityData() *types.CommonEntityData {
    masters.EntityData.YFilter = masters.YFilter
    masters.EntityData.YangName = "masters"
    masters.EntityData.BundleName = "cisco_ios_xr"
    masters.EntityData.ParentYangName = "profile"
    masters.EntityData.SegmentPath = "masters"
    masters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + masters.EntityData.SegmentPath
    masters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masters.EntityData.Children = types.NewOrderedMap()
    masters.EntityData.Children.Append("master", types.YChild{"Master", nil})
    for i := range masters.Master {
        masters.EntityData.Children.Append(types.GetSegmentPath(masters.Master[i]), types.YChild{"Master", masters.Master[i]})
    }
    masters.EntityData.Leafs = types.NewOrderedMap()

    masters.EntityData.YListKeys = []string {}

    return &(masters.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master
// Master configuration
type Ptp_Profiles_Profile_Masters_Master struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Master Transport Type. The type is PtpEncap.
    Transport interface{}

    // ethernet. The type is slice of
    // Ptp_Profiles_Profile_Masters_Master_Ethernet.
    Ethernet []*Ptp_Profiles_Profile_Masters_Master_Ethernet

    // ipv4 or ipv6. The type is slice of
    // Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6.
    Ipv4OrIpv6 []*Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6
}

func (master *Ptp_Profiles_Profile_Masters_Master) GetEntityData() *types.CommonEntityData {
    master.EntityData.YFilter = master.YFilter
    master.EntityData.YangName = "master"
    master.EntityData.BundleName = "cisco_ios_xr"
    master.EntityData.ParentYangName = "masters"
    master.EntityData.SegmentPath = "master" + types.AddKeyToken(master.Transport, "transport")
    master.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/masters/" + master.EntityData.SegmentPath
    master.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    master.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    master.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    master.EntityData.Children = types.NewOrderedMap()
    master.EntityData.Children.Append("ethernet", types.YChild{"Ethernet", nil})
    for i := range master.Ethernet {
        master.EntityData.Children.Append(types.GetSegmentPath(master.Ethernet[i]), types.YChild{"Ethernet", master.Ethernet[i]})
    }
    master.EntityData.Children.Append("ipv4-or-ipv6", types.YChild{"Ipv4OrIpv6", nil})
    for i := range master.Ipv4OrIpv6 {
        master.EntityData.Children.Append(types.GetSegmentPath(master.Ipv4OrIpv6[i]), types.YChild{"Ipv4OrIpv6", master.Ipv4OrIpv6[i]})
    }
    master.EntityData.Leafs = types.NewOrderedMap()
    master.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", master.Transport})

    master.EntityData.YListKeys = []string {"Transport"}

    return &(master.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ethernet
// ethernet
type Ptp_Profiles_Profile_Masters_Master_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Master MAC Address - only used if Transport is
    // Ethernet. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MasterMacAddress interface{}

    // Master clock class. The type is interface{} with range: 0..255.
    MasterClockClass interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}

    // Master priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Communication Model. The type is PtpTransport. The default value is
    // unicast.
    Communication interface{}

    // The delay asymmetry for this master.
    DelayAsymmetry Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry
}

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "master"
    ethernet.EntityData.SegmentPath = "ethernet" + types.AddKeyToken(ethernet.MasterMacAddress, "master-mac-address")
    ethernet.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/masters/master/" + ethernet.EntityData.SegmentPath
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Children.Append("delay-asymmetry", types.YChild{"DelayAsymmetry", &ethernet.DelayAsymmetry})
    ethernet.EntityData.Leafs = types.NewOrderedMap()
    ethernet.EntityData.Leafs.Append("master-mac-address", types.YLeaf{"MasterMacAddress", ethernet.MasterMacAddress})
    ethernet.EntityData.Leafs.Append("master-clock-class", types.YLeaf{"MasterClockClass", ethernet.MasterClockClass})
    ethernet.EntityData.Leafs.Append("non-negotiated", types.YLeaf{"NonNegotiated", ethernet.NonNegotiated})
    ethernet.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ethernet.Priority})
    ethernet.EntityData.Leafs.Append("communication", types.YLeaf{"Communication", ethernet.Communication})

    ethernet.EntityData.YListKeys = []string {"MasterMacAddress"}

    return &(ethernet.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry
// The delay asymmetry for this master
// This type is a presence type.
type Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // How much longer the master to slave path takes than the reverse. The type
    // is interface{} with range: -500000000..500000000. This attribute is
    // mandatory.
    Magnitude interface{}

    // The units to use for the delay asymmetry. The type is
    // PtpDelayAsymmetryUnits. This attribute is mandatory.
    Units interface{}
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetEntityData() *types.CommonEntityData {
    delayAsymmetry.EntityData.YFilter = delayAsymmetry.YFilter
    delayAsymmetry.EntityData.YangName = "delay-asymmetry"
    delayAsymmetry.EntityData.BundleName = "cisco_ios_xr"
    delayAsymmetry.EntityData.ParentYangName = "ethernet"
    delayAsymmetry.EntityData.SegmentPath = "delay-asymmetry"
    delayAsymmetry.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/masters/master/ethernet/" + delayAsymmetry.EntityData.SegmentPath
    delayAsymmetry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayAsymmetry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayAsymmetry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayAsymmetry.EntityData.Children = types.NewOrderedMap()
    delayAsymmetry.EntityData.Leafs = types.NewOrderedMap()
    delayAsymmetry.EntityData.Leafs.Append("magnitude", types.YLeaf{"Magnitude", delayAsymmetry.Magnitude})
    delayAsymmetry.EntityData.Leafs.Append("units", types.YLeaf{"Units", delayAsymmetry.Units})

    delayAsymmetry.EntityData.YListKeys = []string {}

    return &(delayAsymmetry.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6
// ipv4 or ipv6
type Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Master IP Address - used if Transport is not
    // Ethernet. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    MasterIpAddress interface{}

    // Master clock class. The type is interface{} with range: 0..255.
    MasterClockClass interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}

    // Master priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Communication Model. The type is PtpTransport. The default value is
    // unicast.
    Communication interface{}

    // The delay asymmetry for this master.
    DelayAsymmetry Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetEntityData() *types.CommonEntityData {
    ipv4OrIpv6.EntityData.YFilter = ipv4OrIpv6.YFilter
    ipv4OrIpv6.EntityData.YangName = "ipv4-or-ipv6"
    ipv4OrIpv6.EntityData.BundleName = "cisco_ios_xr"
    ipv4OrIpv6.EntityData.ParentYangName = "master"
    ipv4OrIpv6.EntityData.SegmentPath = "ipv4-or-ipv6" + types.AddKeyToken(ipv4OrIpv6.MasterIpAddress, "master-ip-address")
    ipv4OrIpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/masters/master/" + ipv4OrIpv6.EntityData.SegmentPath
    ipv4OrIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4OrIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4OrIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4OrIpv6.EntityData.Children = types.NewOrderedMap()
    ipv4OrIpv6.EntityData.Children.Append("delay-asymmetry", types.YChild{"DelayAsymmetry", &ipv4OrIpv6.DelayAsymmetry})
    ipv4OrIpv6.EntityData.Leafs = types.NewOrderedMap()
    ipv4OrIpv6.EntityData.Leafs.Append("master-ip-address", types.YLeaf{"MasterIpAddress", ipv4OrIpv6.MasterIpAddress})
    ipv4OrIpv6.EntityData.Leafs.Append("master-clock-class", types.YLeaf{"MasterClockClass", ipv4OrIpv6.MasterClockClass})
    ipv4OrIpv6.EntityData.Leafs.Append("non-negotiated", types.YLeaf{"NonNegotiated", ipv4OrIpv6.NonNegotiated})
    ipv4OrIpv6.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ipv4OrIpv6.Priority})
    ipv4OrIpv6.EntityData.Leafs.Append("communication", types.YLeaf{"Communication", ipv4OrIpv6.Communication})

    ipv4OrIpv6.EntityData.YListKeys = []string {"MasterIpAddress"}

    return &(ipv4OrIpv6.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry
// The delay asymmetry for this master
// This type is a presence type.
type Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // How much longer the master to slave path takes than the reverse. The type
    // is interface{} with range: -500000000..500000000. This attribute is
    // mandatory.
    Magnitude interface{}

    // The units to use for the delay asymmetry. The type is
    // PtpDelayAsymmetryUnits. This attribute is mandatory.
    Units interface{}
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetEntityData() *types.CommonEntityData {
    delayAsymmetry.EntityData.YFilter = delayAsymmetry.YFilter
    delayAsymmetry.EntityData.YangName = "delay-asymmetry"
    delayAsymmetry.EntityData.BundleName = "cisco_ios_xr"
    delayAsymmetry.EntityData.ParentYangName = "ipv4-or-ipv6"
    delayAsymmetry.EntityData.SegmentPath = "delay-asymmetry"
    delayAsymmetry.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/masters/master/ipv4-or-ipv6/" + delayAsymmetry.EntityData.SegmentPath
    delayAsymmetry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayAsymmetry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayAsymmetry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayAsymmetry.EntityData.Children = types.NewOrderedMap()
    delayAsymmetry.EntityData.Leafs = types.NewOrderedMap()
    delayAsymmetry.EntityData.Leafs.Append("magnitude", types.YLeaf{"Magnitude", delayAsymmetry.Magnitude})
    delayAsymmetry.EntityData.Leafs.Append("units", types.YLeaf{"Units", delayAsymmetry.Units})

    delayAsymmetry.EntityData.YListKeys = []string {}

    return &(delayAsymmetry.EntityData)
}

// Ptp_Profiles_Profile_Communication
// Communication model
type Ptp_Profiles_Profile_Communication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Communication Model. The type is PtpTransport. The default value is
    // unicast.
    Model interface{}

    // Target address set. The type is bool. The default value is false.
    TargetAddressSet interface{}

    // Target address. The type is string.
    TargetAddress interface{}
}

func (communication *Ptp_Profiles_Profile_Communication) GetEntityData() *types.CommonEntityData {
    communication.EntityData.YFilter = communication.YFilter
    communication.EntityData.YangName = "communication"
    communication.EntityData.BundleName = "cisco_ios_xr"
    communication.EntityData.ParentYangName = "profile"
    communication.EntityData.SegmentPath = "communication"
    communication.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + communication.EntityData.SegmentPath
    communication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    communication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    communication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    communication.EntityData.Children = types.NewOrderedMap()
    communication.EntityData.Leafs = types.NewOrderedMap()
    communication.EntityData.Leafs.Append("model", types.YLeaf{"Model", communication.Model})
    communication.EntityData.Leafs.Append("target-address-set", types.YLeaf{"TargetAddressSet", communication.TargetAddressSet})
    communication.EntityData.Leafs.Append("target-address", types.YLeaf{"TargetAddress", communication.TargetAddress})

    communication.EntityData.YListKeys = []string {}

    return &(communication.EntityData)
}

// Ptp_Profiles_Profile_DelayRequestMinimumInterval
// Minimum delay request interval
type Ptp_Profiles_Profile_DelayRequestMinimumInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval or Frequency. The type is PtpTime. The default value is interval.
    TimeType interface{}

    // Time Period. The type is PtpTimePeriod. The default value is 1.
    TimePeriod interface{}
}

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetEntityData() *types.CommonEntityData {
    delayRequestMinimumInterval.EntityData.YFilter = delayRequestMinimumInterval.YFilter
    delayRequestMinimumInterval.EntityData.YangName = "delay-request-minimum-interval"
    delayRequestMinimumInterval.EntityData.BundleName = "cisco_ios_xr"
    delayRequestMinimumInterval.EntityData.ParentYangName = "profile"
    delayRequestMinimumInterval.EntityData.SegmentPath = "delay-request-minimum-interval"
    delayRequestMinimumInterval.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + delayRequestMinimumInterval.EntityData.SegmentPath
    delayRequestMinimumInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayRequestMinimumInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayRequestMinimumInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayRequestMinimumInterval.EntityData.Children = types.NewOrderedMap()
    delayRequestMinimumInterval.EntityData.Leafs = types.NewOrderedMap()
    delayRequestMinimumInterval.EntityData.Leafs.Append("time-type", types.YLeaf{"TimeType", delayRequestMinimumInterval.TimeType})
    delayRequestMinimumInterval.EntityData.Leafs.Append("time-period", types.YLeaf{"TimePeriod", delayRequestMinimumInterval.TimePeriod})

    delayRequestMinimumInterval.EntityData.YListKeys = []string {}

    return &(delayRequestMinimumInterval.EntityData)
}

// Ptp_Profiles_Profile_SourceIpv6Address
// Source IPv6 Address
type Ptp_Profiles_Profile_SourceIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable source IPv6 address. The type is bool.
    Enable interface{}

    // Source IPv6 address to use. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceIpv6 interface{}
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetEntityData() *types.CommonEntityData {
    sourceIpv6Address.EntityData.YFilter = sourceIpv6Address.YFilter
    sourceIpv6Address.EntityData.YangName = "source-ipv6-address"
    sourceIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    sourceIpv6Address.EntityData.ParentYangName = "profile"
    sourceIpv6Address.EntityData.SegmentPath = "source-ipv6-address"
    sourceIpv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/profiles/profile/" + sourceIpv6Address.EntityData.SegmentPath
    sourceIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceIpv6Address.EntityData.Children = types.NewOrderedMap()
    sourceIpv6Address.EntityData.Leafs = types.NewOrderedMap()
    sourceIpv6Address.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", sourceIpv6Address.Enable})
    sourceIpv6Address.EntityData.Leafs.Append("source-ipv6", types.YLeaf{"SourceIpv6", sourceIpv6Address.SourceIpv6})

    sourceIpv6Address.EntityData.YListKeys = []string {}

    return &(sourceIpv6Address.EntityData)
}

// Ptp_UtcOffset
// UTC offset configuration
type Ptp_UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Base UTC offset configuration. The type is interface{} with range:
    // 0..32767. Units are second.
    BaseOffset interface{}

    // Source file containing leap second information.
    LeapSecondFile Ptp_UtcOffset_LeapSecondFile

    // Table for scheduled UTC offset configuration.
    ScheduledOffsets Ptp_UtcOffset_ScheduledOffsets
}

func (utcOffset *Ptp_UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "ptp"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/" + utcOffset.EntityData.SegmentPath
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = types.NewOrderedMap()
    utcOffset.EntityData.Children.Append("leap-second-file", types.YChild{"LeapSecondFile", &utcOffset.LeapSecondFile})
    utcOffset.EntityData.Children.Append("scheduled-offsets", types.YChild{"ScheduledOffsets", &utcOffset.ScheduledOffsets})
    utcOffset.EntityData.Leafs = types.NewOrderedMap()
    utcOffset.EntityData.Leafs.Append("base-offset", types.YLeaf{"BaseOffset", utcOffset.BaseOffset})

    utcOffset.EntityData.YListKeys = []string {}

    return &(utcOffset.EntityData)
}

// Ptp_UtcOffset_LeapSecondFile
// Source file containing leap second information
// This type is a presence type.
type Ptp_UtcOffset_LeapSecondFile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // URL of source file. The type is string. This attribute is mandatory.
    SourceUrl interface{}

    // Polling frequency, in days. The type is interface{} with range: 1..365.
    // Units are day.
    PollingFrequency interface{}
}

func (leapSecondFile *Ptp_UtcOffset_LeapSecondFile) GetEntityData() *types.CommonEntityData {
    leapSecondFile.EntityData.YFilter = leapSecondFile.YFilter
    leapSecondFile.EntityData.YangName = "leap-second-file"
    leapSecondFile.EntityData.BundleName = "cisco_ios_xr"
    leapSecondFile.EntityData.ParentYangName = "utc-offset"
    leapSecondFile.EntityData.SegmentPath = "leap-second-file"
    leapSecondFile.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/utc-offset/" + leapSecondFile.EntityData.SegmentPath
    leapSecondFile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leapSecondFile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leapSecondFile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leapSecondFile.EntityData.Children = types.NewOrderedMap()
    leapSecondFile.EntityData.Leafs = types.NewOrderedMap()
    leapSecondFile.EntityData.Leafs.Append("source-url", types.YLeaf{"SourceUrl", leapSecondFile.SourceUrl})
    leapSecondFile.EntityData.Leafs.Append("polling-frequency", types.YLeaf{"PollingFrequency", leapSecondFile.PollingFrequency})

    leapSecondFile.EntityData.YListKeys = []string {}

    return &(leapSecondFile.EntityData)
}

// Ptp_UtcOffset_ScheduledOffsets
// Table for scheduled UTC offset configuration
type Ptp_UtcOffset_ScheduledOffsets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Scheduled UTC offset configuration. The type is slice of
    // Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset.
    ScheduledOffset []*Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset
}

func (scheduledOffsets *Ptp_UtcOffset_ScheduledOffsets) GetEntityData() *types.CommonEntityData {
    scheduledOffsets.EntityData.YFilter = scheduledOffsets.YFilter
    scheduledOffsets.EntityData.YangName = "scheduled-offsets"
    scheduledOffsets.EntityData.BundleName = "cisco_ios_xr"
    scheduledOffsets.EntityData.ParentYangName = "utc-offset"
    scheduledOffsets.EntityData.SegmentPath = "scheduled-offsets"
    scheduledOffsets.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/utc-offset/" + scheduledOffsets.EntityData.SegmentPath
    scheduledOffsets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scheduledOffsets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scheduledOffsets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scheduledOffsets.EntityData.Children = types.NewOrderedMap()
    scheduledOffsets.EntityData.Children.Append("scheduled-offset", types.YChild{"ScheduledOffset", nil})
    for i := range scheduledOffsets.ScheduledOffset {
        scheduledOffsets.EntityData.Children.Append(types.GetSegmentPath(scheduledOffsets.ScheduledOffset[i]), types.YChild{"ScheduledOffset", scheduledOffsets.ScheduledOffset[i]})
    }
    scheduledOffsets.EntityData.Leafs = types.NewOrderedMap()

    scheduledOffsets.EntityData.YListKeys = []string {}

    return &(scheduledOffsets.EntityData)
}

// Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset
// Scheduled UTC offset configuration
type Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Offset application date, in ISO-8601 format
    // (YYYY-MM-DD). The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Date interface{}

    // UTC offset, in seconds. The type is interface{} with range: 0..32767. This
    // attribute is mandatory. Units are second.
    Offset interface{}
}

func (scheduledOffset *Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset) GetEntityData() *types.CommonEntityData {
    scheduledOffset.EntityData.YFilter = scheduledOffset.YFilter
    scheduledOffset.EntityData.YangName = "scheduled-offset"
    scheduledOffset.EntityData.BundleName = "cisco_ios_xr"
    scheduledOffset.EntityData.ParentYangName = "scheduled-offsets"
    scheduledOffset.EntityData.SegmentPath = "scheduled-offset" + types.AddKeyToken(scheduledOffset.Date, "date")
    scheduledOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/utc-offset/scheduled-offsets/" + scheduledOffset.EntityData.SegmentPath
    scheduledOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scheduledOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scheduledOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scheduledOffset.EntityData.Children = types.NewOrderedMap()
    scheduledOffset.EntityData.Leafs = types.NewOrderedMap()
    scheduledOffset.EntityData.Leafs.Append("date", types.YLeaf{"Date", scheduledOffset.Date})
    scheduledOffset.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", scheduledOffset.Offset})

    scheduledOffset.EntityData.YListKeys = []string {"Date"}

    return &(scheduledOffset.EntityData)
}

// Ptp_Logging
// PTP logging configuration
type Ptp_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PTP best master clock logging configuration.
    BestMasterClock Ptp_Logging_BestMasterClock

    // PTP PD Servo logging configuration.
    Servo Ptp_Logging_Servo
}

func (logging *Ptp_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "ptp"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/" + logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Children.Append("best-master-clock", types.YChild{"BestMasterClock", &logging.BestMasterClock})
    logging.EntityData.Children.Append("Cisco-IOS-XR-asr9k-ptp-pd-cfg:servo", types.YChild{"Servo", &logging.Servo})
    logging.EntityData.Leafs = types.NewOrderedMap()

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Ptp_Logging_BestMasterClock
// PTP best master clock logging configuration
type Ptp_Logging_BestMasterClock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable best master clock changes logging. The type is interface{}.
    Changes interface{}
}

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetEntityData() *types.CommonEntityData {
    bestMasterClock.EntityData.YFilter = bestMasterClock.YFilter
    bestMasterClock.EntityData.YangName = "best-master-clock"
    bestMasterClock.EntityData.BundleName = "cisco_ios_xr"
    bestMasterClock.EntityData.ParentYangName = "logging"
    bestMasterClock.EntityData.SegmentPath = "best-master-clock"
    bestMasterClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/logging/" + bestMasterClock.EntityData.SegmentPath
    bestMasterClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bestMasterClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bestMasterClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bestMasterClock.EntityData.Children = types.NewOrderedMap()
    bestMasterClock.EntityData.Leafs = types.NewOrderedMap()
    bestMasterClock.EntityData.Leafs.Append("changes", types.YLeaf{"Changes", bestMasterClock.Changes})

    bestMasterClock.EntityData.YListKeys = []string {}

    return &(bestMasterClock.EntityData)
}

// Ptp_Logging_Servo
// PTP PD Servo logging configuration
type Ptp_Logging_Servo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable servo events logging. The type is interface{}.
    Events interface{}
}

func (servo *Ptp_Logging_Servo) GetEntityData() *types.CommonEntityData {
    servo.EntityData.YFilter = servo.YFilter
    servo.EntityData.YangName = "servo"
    servo.EntityData.BundleName = "cisco_ios_xr"
    servo.EntityData.ParentYangName = "logging"
    servo.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-ptp-pd-cfg:servo"
    servo.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/logging/" + servo.EntityData.SegmentPath
    servo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servo.EntityData.Children = types.NewOrderedMap()
    servo.EntityData.Leafs = types.NewOrderedMap()
    servo.EntityData.Leafs.Append("events", types.YLeaf{"Events", servo.Events})

    servo.EntityData.YListKeys = []string {}

    return &(servo.EntityData)
}

// Ptp_UncalibratedClockClass2
// Clock class to be used while acquiring
// phase-lock to a parent clock.
// This type is a presence type.
type Ptp_UncalibratedClockClass2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Clock Class. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    ClockClass interface{}

    // Unless from holdover flag. The type is bool.
    UnlessFromHoldover interface{}
}

func (uncalibratedClockClass2 *Ptp_UncalibratedClockClass2) GetEntityData() *types.CommonEntityData {
    uncalibratedClockClass2.EntityData.YFilter = uncalibratedClockClass2.YFilter
    uncalibratedClockClass2.EntityData.YangName = "uncalibrated-clock-class2"
    uncalibratedClockClass2.EntityData.BundleName = "cisco_ios_xr"
    uncalibratedClockClass2.EntityData.ParentYangName = "ptp"
    uncalibratedClockClass2.EntityData.SegmentPath = "uncalibrated-clock-class2"
    uncalibratedClockClass2.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/" + uncalibratedClockClass2.EntityData.SegmentPath
    uncalibratedClockClass2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    uncalibratedClockClass2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    uncalibratedClockClass2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    uncalibratedClockClass2.EntityData.Children = types.NewOrderedMap()
    uncalibratedClockClass2.EntityData.Leafs = types.NewOrderedMap()
    uncalibratedClockClass2.EntityData.Leafs.Append("clock-class", types.YLeaf{"ClockClass", uncalibratedClockClass2.ClockClass})
    uncalibratedClockClass2.EntityData.Leafs.Append("unless-from-holdover", types.YLeaf{"UnlessFromHoldover", uncalibratedClockClass2.UnlessFromHoldover})

    uncalibratedClockClass2.EntityData.YListKeys = []string {}

    return &(uncalibratedClockClass2.EntityData)
}

// Ptp_TransparentClock
// Transparent clock configuration
type Ptp_TransparentClock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of domains containing transparent clock configuration.
    Domains Ptp_TransparentClock_Domains
}

func (transparentClock *Ptp_TransparentClock) GetEntityData() *types.CommonEntityData {
    transparentClock.EntityData.YFilter = transparentClock.YFilter
    transparentClock.EntityData.YangName = "transparent-clock"
    transparentClock.EntityData.BundleName = "cisco_ios_xr"
    transparentClock.EntityData.ParentYangName = "ptp"
    transparentClock.EntityData.SegmentPath = "transparent-clock"
    transparentClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/" + transparentClock.EntityData.SegmentPath
    transparentClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transparentClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transparentClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transparentClock.EntityData.Children = types.NewOrderedMap()
    transparentClock.EntityData.Children.Append("domains", types.YChild{"Domains", &transparentClock.Domains})
    transparentClock.EntityData.Leafs = types.NewOrderedMap()

    transparentClock.EntityData.YListKeys = []string {}

    return &(transparentClock.EntityData)
}

// Ptp_TransparentClock_Domains
// Table of domains containing transparent clock
// configuration
type Ptp_TransparentClock_Domains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transparent clock domain configuration. The type is slice of
    // Ptp_TransparentClock_Domains_Domain.
    Domain []*Ptp_TransparentClock_Domains_Domain
}

func (domains *Ptp_TransparentClock_Domains) GetEntityData() *types.CommonEntityData {
    domains.EntityData.YFilter = domains.YFilter
    domains.EntityData.YangName = "domains"
    domains.EntityData.BundleName = "cisco_ios_xr"
    domains.EntityData.ParentYangName = "transparent-clock"
    domains.EntityData.SegmentPath = "domains"
    domains.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/transparent-clock/" + domains.EntityData.SegmentPath
    domains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domains.EntityData.Children = types.NewOrderedMap()
    domains.EntityData.Children.Append("domain", types.YChild{"Domain", nil})
    for i := range domains.Domain {
        domains.EntityData.Children.Append(types.GetSegmentPath(domains.Domain[i]), types.YChild{"Domain", domains.Domain[i]})
    }
    domains.EntityData.Leafs = types.NewOrderedMap()

    domains.EntityData.YListKeys = []string {}

    return &(domains.EntityData)
}

// Ptp_TransparentClock_Domains_Domain
// Transparent clock domain configuration
type Ptp_TransparentClock_Domains_Domain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Domain. The type is string with pattern: (all).
    Domain interface{}
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetEntityData() *types.CommonEntityData {
    domain.EntityData.YFilter = domain.YFilter
    domain.EntityData.YangName = "domain"
    domain.EntityData.BundleName = "cisco_ios_xr"
    domain.EntityData.ParentYangName = "domains"
    domain.EntityData.SegmentPath = "domain" + types.AddKeyToken(domain.Domain, "domain")
    domain.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/transparent-clock/domains/" + domain.EntityData.SegmentPath
    domain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domain.EntityData.Children = types.NewOrderedMap()
    domain.EntityData.Leafs = types.NewOrderedMap()
    domain.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", domain.Domain})

    domain.EntityData.YListKeys = []string {"Domain"}

    return &(domain.EntityData)
}

// Ptp_VirtualPort
// PTP virtual port configuration
type Ptp_VirtualPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual port clock accuracy. The type is interface{} with range: 0..254.
    ClockAccuracy interface{}

    // Enable the PTP Virtual Port. The type is interface{}.
    Enable interface{}

    // Virtual port priority2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Virtual port local priority. The type is interface{} with range: 1..255.
    LocalPriority interface{}

    // Virtual port OSLV. The type is interface{} with range: 0..65535.
    OffsetScaledLogVariance interface{}

    // Virtual port priority1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Virtual port clock class. The type is interface{} with range: 0..255.
    ClockClass interface{}
}

func (virtualPort *Ptp_VirtualPort) GetEntityData() *types.CommonEntityData {
    virtualPort.EntityData.YFilter = virtualPort.YFilter
    virtualPort.EntityData.YangName = "virtual-port"
    virtualPort.EntityData.BundleName = "cisco_ios_xr"
    virtualPort.EntityData.ParentYangName = "ptp"
    virtualPort.EntityData.SegmentPath = "virtual-port"
    virtualPort.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-cfg:ptp/" + virtualPort.EntityData.SegmentPath
    virtualPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualPort.EntityData.Children = types.NewOrderedMap()
    virtualPort.EntityData.Leafs = types.NewOrderedMap()
    virtualPort.EntityData.Leafs.Append("clock-accuracy", types.YLeaf{"ClockAccuracy", virtualPort.ClockAccuracy})
    virtualPort.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", virtualPort.Enable})
    virtualPort.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", virtualPort.Priority2})
    virtualPort.EntityData.Leafs.Append("local-priority", types.YLeaf{"LocalPriority", virtualPort.LocalPriority})
    virtualPort.EntityData.Leafs.Append("offset-scaled-log-variance", types.YLeaf{"OffsetScaledLogVariance", virtualPort.OffsetScaledLogVariance})
    virtualPort.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", virtualPort.Priority1})
    virtualPort.EntityData.Leafs.Append("clock-class", types.YLeaf{"ClockClass", virtualPort.ClockClass})

    virtualPort.EntityData.YListKeys = []string {}

    return &(virtualPort.EntityData)
}

