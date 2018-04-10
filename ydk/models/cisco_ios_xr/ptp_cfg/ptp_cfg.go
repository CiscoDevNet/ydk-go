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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

    // Clock class to be used while acquiring phase-lock to a parent clock. The
    // type is interface{} with range: 0..255.
    UncalibratedClockClass interface{}

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

    // Transparent clock configuration.
    TransparentClock Ptp_TransparentClock
}

func (ptp *Ptp) GetEntityData() *types.CommonEntityData {
    ptp.EntityData.YFilter = ptp.YFilter
    ptp.EntityData.YangName = "ptp"
    ptp.EntityData.BundleName = "cisco_ios_xr"
    ptp.EntityData.ParentYangName = "Cisco-IOS-XR-ptp-cfg"
    ptp.EntityData.SegmentPath = "Cisco-IOS-XR-ptp-cfg:ptp"
    ptp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptp.EntityData.Children = make(map[string]types.YChild)
    ptp.EntityData.Children["clock"] = types.YChild{"Clock", &ptp.Clock}
    ptp.EntityData.Children["profiles"] = types.YChild{"Profiles", &ptp.Profiles}
    ptp.EntityData.Children["utc-offset"] = types.YChild{"UtcOffset", &ptp.UtcOffset}
    ptp.EntityData.Children["logging"] = types.YChild{"Logging", &ptp.Logging}
    ptp.EntityData.Children["transparent-clock"] = types.YChild{"TransparentClock", &ptp.TransparentClock}
    ptp.EntityData.Leafs = make(map[string]types.YLeaf)
    ptp.EntityData.Leafs["time-of-day-priority"] = types.YLeaf{"TimeOfDayPriority", ptp.TimeOfDayPriority}
    ptp.EntityData.Leafs["frequency-priority"] = types.YLeaf{"FrequencyPriority", ptp.FrequencyPriority}
    ptp.EntityData.Leafs["startup-clock-class"] = types.YLeaf{"StartupClockClass", ptp.StartupClockClass}
    ptp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ptp.Enable}
    ptp.EntityData.Leafs["min-clock-class"] = types.YLeaf{"MinClockClass", ptp.MinClockClass}
    ptp.EntityData.Leafs["uncalibrated-clock-class"] = types.YLeaf{"UncalibratedClockClass", ptp.UncalibratedClockClass}
    ptp.EntityData.Leafs["freerun-clock-class"] = types.YLeaf{"FreerunClockClass", ptp.FreerunClockClass}
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
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = make(map[string]types.YChild)
    clock.EntityData.Children["profile"] = types.YChild{"Profile", &clock.Profile}
    clock.EntityData.Children["identity"] = types.YChild{"Identity", &clock.Identity}
    clock.EntityData.Leafs = make(map[string]types.YLeaf)
    clock.EntityData.Leafs["timescale"] = types.YLeaf{"Timescale", clock.Timescale}
    clock.EntityData.Leafs["domain"] = types.YLeaf{"Domain", clock.Domain}
    clock.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", clock.Priority2}
    clock.EntityData.Leafs["time-source"] = types.YLeaf{"TimeSource", clock.TimeSource}
    clock.EntityData.Leafs["priority1"] = types.YLeaf{"Priority1", clock.Priority1}
    clock.EntityData.Leafs["clock-class"] = types.YLeaf{"ClockClass", clock.ClockClass}
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
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["clock-profile"] = types.YLeaf{"ClockProfile", profile.ClockProfile}
    profile.EntityData.Leafs["telecom-clock-type"] = types.YLeaf{"TelecomClockType", profile.TelecomClockType}
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    identity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identity.EntityData.Children = make(map[string]types.YChild)
    identity.EntityData.Leafs = make(map[string]types.YLeaf)
    identity.EntityData.Leafs["clock-id-type"] = types.YLeaf{"ClockIdType", identity.ClockIdType}
    identity.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", identity.MacAddress}
    identity.EntityData.Leafs["eui"] = types.YLeaf{"Eui", identity.Eui}
    return &(identity.EntityData)
}

// Ptp_Profiles
// Table for profile configuration
type Ptp_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile configuration. The type is slice of Ptp_Profiles_Profile.
    Profile []Ptp_Profiles_Profile
}

func (profiles *Ptp_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "ptp"
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

// Ptp_Profiles_Profile
// Profile configuration
type Ptp_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["announce-interval"] = types.YChild{"AnnounceInterval", &profile.AnnounceInterval}
    profile.EntityData.Children["source-ipv4-address"] = types.YChild{"SourceIpv4Address", &profile.SourceIpv4Address}
    profile.EntityData.Children["slaves"] = types.YChild{"Slaves", &profile.Slaves}
    profile.EntityData.Children["sync-interval"] = types.YChild{"SyncInterval", &profile.SyncInterval}
    profile.EntityData.Children["masters"] = types.YChild{"Masters", &profile.Masters}
    profile.EntityData.Children["communication"] = types.YChild{"Communication", &profile.Communication}
    profile.EntityData.Children["delay-request-minimum-interval"] = types.YChild{"DelayRequestMinimumInterval", &profile.DelayRequestMinimumInterval}
    profile.EntityData.Children["source-ipv6-address"] = types.YChild{"SourceIpv6Address", &profile.SourceIpv6Address}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    profile.EntityData.Leafs["sync-grant-duration"] = types.YLeaf{"SyncGrantDuration", profile.SyncGrantDuration}
    profile.EntityData.Leafs["general-cos"] = types.YLeaf{"GeneralCos", profile.GeneralCos}
    profile.EntityData.Leafs["sync-timeout"] = types.YLeaf{"SyncTimeout", profile.SyncTimeout}
    profile.EntityData.Leafs["transport"] = types.YLeaf{"Transport", profile.Transport}
    profile.EntityData.Leafs["announce-timeout"] = types.YLeaf{"AnnounceTimeout", profile.AnnounceTimeout}
    profile.EntityData.Leafs["cos"] = types.YLeaf{"Cos", profile.Cos}
    profile.EntityData.Leafs["port-state"] = types.YLeaf{"PortState", profile.PortState}
    profile.EntityData.Leafs["delay-response-timeout"] = types.YLeaf{"DelayResponseTimeout", profile.DelayResponseTimeout}
    profile.EntityData.Leafs["delay-response-grant-duration"] = types.YLeaf{"DelayResponseGrantDuration", profile.DelayResponseGrantDuration}
    profile.EntityData.Leafs["event-cos"] = types.YLeaf{"EventCos", profile.EventCos}
    profile.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", profile.Dscp}
    profile.EntityData.Leafs["general-dscp"] = types.YLeaf{"GeneralDscp", profile.GeneralDscp}
    profile.EntityData.Leafs["clock-operation"] = types.YLeaf{"ClockOperation", profile.ClockOperation}
    profile.EntityData.Leafs["announce-grant-duration"] = types.YLeaf{"AnnounceGrantDuration", profile.AnnounceGrantDuration}
    profile.EntityData.Leafs["unicast-grant-invalid-request"] = types.YLeaf{"UnicastGrantInvalidRequest", profile.UnicastGrantInvalidRequest}
    profile.EntityData.Leafs["event-dscp"] = types.YLeaf{"EventDscp", profile.EventDscp}
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
    announceInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceInterval.EntityData.Children = make(map[string]types.YChild)
    announceInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    announceInterval.EntityData.Leafs["time-type"] = types.YLeaf{"TimeType", announceInterval.TimeType}
    announceInterval.EntityData.Leafs["time-period"] = types.YLeaf{"TimePeriod", announceInterval.TimePeriod}
    return &(announceInterval.EntityData)
}

// Ptp_Profiles_Profile_SourceIpv4Address
// Source IPv4 Address
type Ptp_Profiles_Profile_SourceIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable source IP address. The type is bool.
    Enable interface{}

    // Source IP address to use. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceIp interface{}
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetEntityData() *types.CommonEntityData {
    sourceIpv4Address.EntityData.YFilter = sourceIpv4Address.YFilter
    sourceIpv4Address.EntityData.YangName = "source-ipv4-address"
    sourceIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    sourceIpv4Address.EntityData.ParentYangName = "profile"
    sourceIpv4Address.EntityData.SegmentPath = "source-ipv4-address"
    sourceIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceIpv4Address.EntityData.Children = make(map[string]types.YChild)
    sourceIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    sourceIpv4Address.EntityData.Leafs["enable"] = types.YLeaf{"Enable", sourceIpv4Address.Enable}
    sourceIpv4Address.EntityData.Leafs["source-ip"] = types.YLeaf{"SourceIp", sourceIpv4Address.SourceIp}
    return &(sourceIpv4Address.EntityData)
}

// Ptp_Profiles_Profile_Slaves
// Table for slave configuration
type Ptp_Profiles_Profile_Slaves struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slave configuration. The type is slice of
    // Ptp_Profiles_Profile_Slaves_Slave.
    Slave []Ptp_Profiles_Profile_Slaves_Slave
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetEntityData() *types.CommonEntityData {
    slaves.EntityData.YFilter = slaves.YFilter
    slaves.EntityData.YangName = "slaves"
    slaves.EntityData.BundleName = "cisco_ios_xr"
    slaves.EntityData.ParentYangName = "profile"
    slaves.EntityData.SegmentPath = "slaves"
    slaves.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaves.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaves.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaves.EntityData.Children = make(map[string]types.YChild)
    slaves.EntityData.Children["slave"] = types.YChild{"Slave", nil}
    for i := range slaves.Slave {
        slaves.EntityData.Children[types.GetSegmentPath(&slaves.Slave[i])] = types.YChild{"Slave", &slaves.Slave[i]}
    }
    slaves.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slaves.EntityData)
}

// Ptp_Profiles_Profile_Slaves_Slave
// Slave configuration
type Ptp_Profiles_Profile_Slaves_Slave struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slave Transport Type. The type is PtpEncap.
    Transport interface{}

    // ethernet. The type is slice of Ptp_Profiles_Profile_Slaves_Slave_Ethernet.
    Ethernet []Ptp_Profiles_Profile_Slaves_Slave_Ethernet

    // ipv4 or ipv6. The type is slice of
    // Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6.
    Ipv4OrIpv6 []Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetEntityData() *types.CommonEntityData {
    slave.EntityData.YFilter = slave.YFilter
    slave.EntityData.YangName = "slave"
    slave.EntityData.BundleName = "cisco_ios_xr"
    slave.EntityData.ParentYangName = "slaves"
    slave.EntityData.SegmentPath = "slave" + "[transport='" + fmt.Sprintf("%v", slave.Transport) + "']"
    slave.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slave.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slave.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slave.EntityData.Children = make(map[string]types.YChild)
    slave.EntityData.Children["ethernet"] = types.YChild{"Ethernet", nil}
    for i := range slave.Ethernet {
        slave.EntityData.Children[types.GetSegmentPath(&slave.Ethernet[i])] = types.YChild{"Ethernet", &slave.Ethernet[i]}
    }
    slave.EntityData.Children["ipv4-or-ipv6"] = types.YChild{"Ipv4OrIpv6", nil}
    for i := range slave.Ipv4OrIpv6 {
        slave.EntityData.Children[types.GetSegmentPath(&slave.Ipv4OrIpv6[i])] = types.YChild{"Ipv4OrIpv6", &slave.Ipv4OrIpv6[i]}
    }
    slave.EntityData.Leafs = make(map[string]types.YLeaf)
    slave.EntityData.Leafs["transport"] = types.YLeaf{"Transport", slave.Transport}
    return &(slave.EntityData)
}

// Ptp_Profiles_Profile_Slaves_Slave_Ethernet
// ethernet
type Ptp_Profiles_Profile_Slaves_Slave_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slave MAC Address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SlaveMacAddress interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "slave"
    ethernet.EntityData.SegmentPath = "ethernet" + "[slave-mac-address='" + fmt.Sprintf("%v", ethernet.SlaveMacAddress) + "']"
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = make(map[string]types.YChild)
    ethernet.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernet.EntityData.Leafs["slave-mac-address"] = types.YLeaf{"SlaveMacAddress", ethernet.SlaveMacAddress}
    ethernet.EntityData.Leafs["non-negotiated"] = types.YLeaf{"NonNegotiated", ethernet.NonNegotiated}
    return &(ethernet.EntityData)
}

// Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6
// ipv4 or ipv6
type Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slave IP Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SlaveIpAddress interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetEntityData() *types.CommonEntityData {
    ipv4OrIpv6.EntityData.YFilter = ipv4OrIpv6.YFilter
    ipv4OrIpv6.EntityData.YangName = "ipv4-or-ipv6"
    ipv4OrIpv6.EntityData.BundleName = "cisco_ios_xr"
    ipv4OrIpv6.EntityData.ParentYangName = "slave"
    ipv4OrIpv6.EntityData.SegmentPath = "ipv4-or-ipv6" + "[slave-ip-address='" + fmt.Sprintf("%v", ipv4OrIpv6.SlaveIpAddress) + "']"
    ipv4OrIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4OrIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4OrIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4OrIpv6.EntityData.Children = make(map[string]types.YChild)
    ipv4OrIpv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4OrIpv6.EntityData.Leafs["slave-ip-address"] = types.YLeaf{"SlaveIpAddress", ipv4OrIpv6.SlaveIpAddress}
    ipv4OrIpv6.EntityData.Leafs["non-negotiated"] = types.YLeaf{"NonNegotiated", ipv4OrIpv6.NonNegotiated}
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
    syncInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncInterval.EntityData.Children = make(map[string]types.YChild)
    syncInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    syncInterval.EntityData.Leafs["time-type"] = types.YLeaf{"TimeType", syncInterval.TimeType}
    syncInterval.EntityData.Leafs["time-period"] = types.YLeaf{"TimePeriod", syncInterval.TimePeriod}
    return &(syncInterval.EntityData)
}

// Ptp_Profiles_Profile_Masters
// Table for master configuration
type Ptp_Profiles_Profile_Masters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Master configuration. The type is slice of
    // Ptp_Profiles_Profile_Masters_Master.
    Master []Ptp_Profiles_Profile_Masters_Master
}

func (masters *Ptp_Profiles_Profile_Masters) GetEntityData() *types.CommonEntityData {
    masters.EntityData.YFilter = masters.YFilter
    masters.EntityData.YangName = "masters"
    masters.EntityData.BundleName = "cisco_ios_xr"
    masters.EntityData.ParentYangName = "profile"
    masters.EntityData.SegmentPath = "masters"
    masters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masters.EntityData.Children = make(map[string]types.YChild)
    masters.EntityData.Children["master"] = types.YChild{"Master", nil}
    for i := range masters.Master {
        masters.EntityData.Children[types.GetSegmentPath(&masters.Master[i])] = types.YChild{"Master", &masters.Master[i]}
    }
    masters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(masters.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master
// Master configuration
type Ptp_Profiles_Profile_Masters_Master struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Master Transport Type. The type is PtpEncap.
    Transport interface{}

    // ethernet. The type is slice of
    // Ptp_Profiles_Profile_Masters_Master_Ethernet.
    Ethernet []Ptp_Profiles_Profile_Masters_Master_Ethernet

    // ipv4 or ipv6. The type is slice of
    // Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6.
    Ipv4OrIpv6 []Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6
}

func (master *Ptp_Profiles_Profile_Masters_Master) GetEntityData() *types.CommonEntityData {
    master.EntityData.YFilter = master.YFilter
    master.EntityData.YangName = "master"
    master.EntityData.BundleName = "cisco_ios_xr"
    master.EntityData.ParentYangName = "masters"
    master.EntityData.SegmentPath = "master" + "[transport='" + fmt.Sprintf("%v", master.Transport) + "']"
    master.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    master.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    master.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    master.EntityData.Children = make(map[string]types.YChild)
    master.EntityData.Children["ethernet"] = types.YChild{"Ethernet", nil}
    for i := range master.Ethernet {
        master.EntityData.Children[types.GetSegmentPath(&master.Ethernet[i])] = types.YChild{"Ethernet", &master.Ethernet[i]}
    }
    master.EntityData.Children["ipv4-or-ipv6"] = types.YChild{"Ipv4OrIpv6", nil}
    for i := range master.Ipv4OrIpv6 {
        master.EntityData.Children[types.GetSegmentPath(&master.Ipv4OrIpv6[i])] = types.YChild{"Ipv4OrIpv6", &master.Ipv4OrIpv6[i]}
    }
    master.EntityData.Leafs = make(map[string]types.YLeaf)
    master.EntityData.Leafs["transport"] = types.YLeaf{"Transport", master.Transport}
    return &(master.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ethernet
// ethernet
type Ptp_Profiles_Profile_Masters_Master_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Master MAC Address - only used if Transport is
    // Ethernet. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    ethernet.EntityData.SegmentPath = "ethernet" + "[master-mac-address='" + fmt.Sprintf("%v", ethernet.MasterMacAddress) + "']"
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = make(map[string]types.YChild)
    ethernet.EntityData.Children["delay-asymmetry"] = types.YChild{"DelayAsymmetry", &ethernet.DelayAsymmetry}
    ethernet.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernet.EntityData.Leafs["master-mac-address"] = types.YLeaf{"MasterMacAddress", ethernet.MasterMacAddress}
    ethernet.EntityData.Leafs["master-clock-class"] = types.YLeaf{"MasterClockClass", ethernet.MasterClockClass}
    ethernet.EntityData.Leafs["non-negotiated"] = types.YLeaf{"NonNegotiated", ethernet.NonNegotiated}
    ethernet.EntityData.Leafs["priority"] = types.YLeaf{"Priority", ethernet.Priority}
    ethernet.EntityData.Leafs["communication"] = types.YLeaf{"Communication", ethernet.Communication}
    return &(ethernet.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry
// The delay asymmetry for this master
// This type is a presence type.
type Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    delayAsymmetry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayAsymmetry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayAsymmetry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayAsymmetry.EntityData.Children = make(map[string]types.YChild)
    delayAsymmetry.EntityData.Leafs = make(map[string]types.YLeaf)
    delayAsymmetry.EntityData.Leafs["magnitude"] = types.YLeaf{"Magnitude", delayAsymmetry.Magnitude}
    delayAsymmetry.EntityData.Leafs["units"] = types.YLeaf{"Units", delayAsymmetry.Units}
    return &(delayAsymmetry.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6
// ipv4 or ipv6
type Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Master IP Address - used if Transport is not
    // Ethernet. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    ipv4OrIpv6.EntityData.SegmentPath = "ipv4-or-ipv6" + "[master-ip-address='" + fmt.Sprintf("%v", ipv4OrIpv6.MasterIpAddress) + "']"
    ipv4OrIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4OrIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4OrIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4OrIpv6.EntityData.Children = make(map[string]types.YChild)
    ipv4OrIpv6.EntityData.Children["delay-asymmetry"] = types.YChild{"DelayAsymmetry", &ipv4OrIpv6.DelayAsymmetry}
    ipv4OrIpv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4OrIpv6.EntityData.Leafs["master-ip-address"] = types.YLeaf{"MasterIpAddress", ipv4OrIpv6.MasterIpAddress}
    ipv4OrIpv6.EntityData.Leafs["master-clock-class"] = types.YLeaf{"MasterClockClass", ipv4OrIpv6.MasterClockClass}
    ipv4OrIpv6.EntityData.Leafs["non-negotiated"] = types.YLeaf{"NonNegotiated", ipv4OrIpv6.NonNegotiated}
    ipv4OrIpv6.EntityData.Leafs["priority"] = types.YLeaf{"Priority", ipv4OrIpv6.Priority}
    ipv4OrIpv6.EntityData.Leafs["communication"] = types.YLeaf{"Communication", ipv4OrIpv6.Communication}
    return &(ipv4OrIpv6.EntityData)
}

// Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry
// The delay asymmetry for this master
// This type is a presence type.
type Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    delayAsymmetry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayAsymmetry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayAsymmetry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayAsymmetry.EntityData.Children = make(map[string]types.YChild)
    delayAsymmetry.EntityData.Leafs = make(map[string]types.YLeaf)
    delayAsymmetry.EntityData.Leafs["magnitude"] = types.YLeaf{"Magnitude", delayAsymmetry.Magnitude}
    delayAsymmetry.EntityData.Leafs["units"] = types.YLeaf{"Units", delayAsymmetry.Units}
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
    communication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    communication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    communication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    communication.EntityData.Children = make(map[string]types.YChild)
    communication.EntityData.Leafs = make(map[string]types.YLeaf)
    communication.EntityData.Leafs["model"] = types.YLeaf{"Model", communication.Model}
    communication.EntityData.Leafs["target-address-set"] = types.YLeaf{"TargetAddressSet", communication.TargetAddressSet}
    communication.EntityData.Leafs["target-address"] = types.YLeaf{"TargetAddress", communication.TargetAddress}
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
    delayRequestMinimumInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayRequestMinimumInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayRequestMinimumInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayRequestMinimumInterval.EntityData.Children = make(map[string]types.YChild)
    delayRequestMinimumInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    delayRequestMinimumInterval.EntityData.Leafs["time-type"] = types.YLeaf{"TimeType", delayRequestMinimumInterval.TimeType}
    delayRequestMinimumInterval.EntityData.Leafs["time-period"] = types.YLeaf{"TimePeriod", delayRequestMinimumInterval.TimePeriod}
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceIpv6 interface{}
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetEntityData() *types.CommonEntityData {
    sourceIpv6Address.EntityData.YFilter = sourceIpv6Address.YFilter
    sourceIpv6Address.EntityData.YangName = "source-ipv6-address"
    sourceIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    sourceIpv6Address.EntityData.ParentYangName = "profile"
    sourceIpv6Address.EntityData.SegmentPath = "source-ipv6-address"
    sourceIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceIpv6Address.EntityData.Children = make(map[string]types.YChild)
    sourceIpv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    sourceIpv6Address.EntityData.Leafs["enable"] = types.YLeaf{"Enable", sourceIpv6Address.Enable}
    sourceIpv6Address.EntityData.Leafs["source-ipv6"] = types.YLeaf{"SourceIpv6", sourceIpv6Address.SourceIpv6}
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
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = make(map[string]types.YChild)
    utcOffset.EntityData.Children["leap-second-file"] = types.YChild{"LeapSecondFile", &utcOffset.LeapSecondFile}
    utcOffset.EntityData.Children["scheduled-offsets"] = types.YChild{"ScheduledOffsets", &utcOffset.ScheduledOffsets}
    utcOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    utcOffset.EntityData.Leafs["base-offset"] = types.YLeaf{"BaseOffset", utcOffset.BaseOffset}
    return &(utcOffset.EntityData)
}

// Ptp_UtcOffset_LeapSecondFile
// Source file containing leap second information
// This type is a presence type.
type Ptp_UtcOffset_LeapSecondFile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    leapSecondFile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leapSecondFile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leapSecondFile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leapSecondFile.EntityData.Children = make(map[string]types.YChild)
    leapSecondFile.EntityData.Leafs = make(map[string]types.YLeaf)
    leapSecondFile.EntityData.Leafs["source-url"] = types.YLeaf{"SourceUrl", leapSecondFile.SourceUrl}
    leapSecondFile.EntityData.Leafs["polling-frequency"] = types.YLeaf{"PollingFrequency", leapSecondFile.PollingFrequency}
    return &(leapSecondFile.EntityData)
}

// Ptp_UtcOffset_ScheduledOffsets
// Table for scheduled UTC offset configuration
type Ptp_UtcOffset_ScheduledOffsets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Scheduled UTC offset configuration. The type is slice of
    // Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset.
    ScheduledOffset []Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset
}

func (scheduledOffsets *Ptp_UtcOffset_ScheduledOffsets) GetEntityData() *types.CommonEntityData {
    scheduledOffsets.EntityData.YFilter = scheduledOffsets.YFilter
    scheduledOffsets.EntityData.YangName = "scheduled-offsets"
    scheduledOffsets.EntityData.BundleName = "cisco_ios_xr"
    scheduledOffsets.EntityData.ParentYangName = "utc-offset"
    scheduledOffsets.EntityData.SegmentPath = "scheduled-offsets"
    scheduledOffsets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scheduledOffsets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scheduledOffsets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scheduledOffsets.EntityData.Children = make(map[string]types.YChild)
    scheduledOffsets.EntityData.Children["scheduled-offset"] = types.YChild{"ScheduledOffset", nil}
    for i := range scheduledOffsets.ScheduledOffset {
        scheduledOffsets.EntityData.Children[types.GetSegmentPath(&scheduledOffsets.ScheduledOffset[i])] = types.YChild{"ScheduledOffset", &scheduledOffsets.ScheduledOffset[i]}
    }
    scheduledOffsets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(scheduledOffsets.EntityData)
}

// Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset
// Scheduled UTC offset configuration
type Ptp_UtcOffset_ScheduledOffsets_ScheduledOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Offset application date, in ISO-8601 format
    // (YYYY-MM-DD). The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    scheduledOffset.EntityData.SegmentPath = "scheduled-offset" + "[date='" + fmt.Sprintf("%v", scheduledOffset.Date) + "']"
    scheduledOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scheduledOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scheduledOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scheduledOffset.EntityData.Children = make(map[string]types.YChild)
    scheduledOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    scheduledOffset.EntityData.Leafs["date"] = types.YLeaf{"Date", scheduledOffset.Date}
    scheduledOffset.EntityData.Leafs["offset"] = types.YLeaf{"Offset", scheduledOffset.Offset}
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
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Children["best-master-clock"] = types.YChild{"BestMasterClock", &logging.BestMasterClock}
    logging.EntityData.Children["Cisco-IOS-XR-asr9k-ptp-pd-cfg:servo"] = types.YChild{"Servo", &logging.Servo}
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
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
    bestMasterClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bestMasterClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bestMasterClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bestMasterClock.EntityData.Children = make(map[string]types.YChild)
    bestMasterClock.EntityData.Leafs = make(map[string]types.YLeaf)
    bestMasterClock.EntityData.Leafs["changes"] = types.YLeaf{"Changes", bestMasterClock.Changes}
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
    servo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servo.EntityData.Children = make(map[string]types.YChild)
    servo.EntityData.Leafs = make(map[string]types.YLeaf)
    servo.EntityData.Leafs["events"] = types.YLeaf{"Events", servo.Events}
    return &(servo.EntityData)
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
    transparentClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transparentClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transparentClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transparentClock.EntityData.Children = make(map[string]types.YChild)
    transparentClock.EntityData.Children["domains"] = types.YChild{"Domains", &transparentClock.Domains}
    transparentClock.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Domain []Ptp_TransparentClock_Domains_Domain
}

func (domains *Ptp_TransparentClock_Domains) GetEntityData() *types.CommonEntityData {
    domains.EntityData.YFilter = domains.YFilter
    domains.EntityData.YangName = "domains"
    domains.EntityData.BundleName = "cisco_ios_xr"
    domains.EntityData.ParentYangName = "transparent-clock"
    domains.EntityData.SegmentPath = "domains"
    domains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domains.EntityData.Children = make(map[string]types.YChild)
    domains.EntityData.Children["domain"] = types.YChild{"Domain", nil}
    for i := range domains.Domain {
        domains.EntityData.Children[types.GetSegmentPath(&domains.Domain[i])] = types.YChild{"Domain", &domains.Domain[i]}
    }
    domains.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(domains.EntityData)
}

// Ptp_TransparentClock_Domains_Domain
// Transparent clock domain configuration
type Ptp_TransparentClock_Domains_Domain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Domain. The type is string with pattern: b'(all)'.
    Domain interface{}
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetEntityData() *types.CommonEntityData {
    domain.EntityData.YFilter = domain.YFilter
    domain.EntityData.YangName = "domain"
    domain.EntityData.BundleName = "cisco_ios_xr"
    domain.EntityData.ParentYangName = "domains"
    domain.EntityData.SegmentPath = "domain" + "[domain='" + fmt.Sprintf("%v", domain.Domain) + "']"
    domain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domain.EntityData.Children = make(map[string]types.YChild)
    domain.EntityData.Leafs = make(map[string]types.YLeaf)
    domain.EntityData.Leafs["domain"] = types.YLeaf{"Domain", domain.Domain}
    return &(domain.EntityData)
}

