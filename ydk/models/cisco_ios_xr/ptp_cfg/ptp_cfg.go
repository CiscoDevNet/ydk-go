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
    parent types.Entity
    YFilter yfilter.YFilter

    // Time-of-day priority. The type is interface{} with range: 1..254. The
    // default value is 100.
    TimeOfDayPriority interface{}

    // Frequency priority. The type is interface{} with range: 1..254. The default
    // value is 254.
    FrequencyPriority interface{}

    // Enable the precision time protocol. The type is interface{}.
    Enable interface{}

    // Clocks with a clock-class higher than the minimum clock class will not be
    // considered for selection as a parent clock. The type is interface{} with
    // range: 0..255.
    MinClockClass interface{}

    // Clock class to be used while acquiring phase-lock to a parent clock. The
    // type is interface{} with range: 0..255.
    UncalibratedClockClass interface{}

    // PTP local clock configuration.
    Clock Ptp_Clock

    // Table for profile configuration.
    Profiles Ptp_Profiles

    // PTP logging configuration.
    Logging Ptp_Logging

    // Transparent clock configuration.
    TransparentClock Ptp_TransparentClock
}

func (ptp *Ptp) GetFilter() yfilter.YFilter { return ptp.YFilter }

func (ptp *Ptp) SetFilter(yf yfilter.YFilter) { ptp.YFilter = yf }

func (ptp *Ptp) GetGoName(yname string) string {
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "frequency-priority" { return "FrequencyPriority" }
    if yname == "enable" { return "Enable" }
    if yname == "min-clock-class" { return "MinClockClass" }
    if yname == "uncalibrated-clock-class" { return "UncalibratedClockClass" }
    if yname == "clock" { return "Clock" }
    if yname == "profiles" { return "Profiles" }
    if yname == "logging" { return "Logging" }
    if yname == "transparent-clock" { return "TransparentClock" }
    return ""
}

func (ptp *Ptp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ptp-cfg:ptp"
}

func (ptp *Ptp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock" {
        return &ptp.Clock
    }
    if childYangName == "profiles" {
        return &ptp.Profiles
    }
    if childYangName == "logging" {
        return &ptp.Logging
    }
    if childYangName == "transparent-clock" {
        return &ptp.TransparentClock
    }
    return nil
}

func (ptp *Ptp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock"] = &ptp.Clock
    children["profiles"] = &ptp.Profiles
    children["logging"] = &ptp.Logging
    children["transparent-clock"] = &ptp.TransparentClock
    return children
}

func (ptp *Ptp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-of-day-priority"] = ptp.TimeOfDayPriority
    leafs["frequency-priority"] = ptp.FrequencyPriority
    leafs["enable"] = ptp.Enable
    leafs["min-clock-class"] = ptp.MinClockClass
    leafs["uncalibrated-clock-class"] = ptp.UncalibratedClockClass
    return leafs
}

func (ptp *Ptp) GetBundleName() string { return "cisco_ios_xr" }

func (ptp *Ptp) GetYangName() string { return "ptp" }

func (ptp *Ptp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ptp *Ptp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ptp *Ptp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ptp *Ptp) SetParent(parent types.Entity) { ptp.parent = parent }

func (ptp *Ptp) GetParent() types.Entity { return ptp.parent }

func (ptp *Ptp) GetParentYangName() string { return "Cisco-IOS-XR-ptp-cfg" }

// Ptp_Clock
// PTP local clock configuration
type Ptp_Clock struct {
    parent types.Entity
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

func (clock *Ptp_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *Ptp_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *Ptp_Clock) GetGoName(yname string) string {
    if yname == "timescale" { return "Timescale" }
    if yname == "domain" { return "Domain" }
    if yname == "priority2" { return "Priority2" }
    if yname == "time-source" { return "TimeSource" }
    if yname == "priority1" { return "Priority1" }
    if yname == "clock-class" { return "ClockClass" }
    if yname == "profile" { return "Profile" }
    if yname == "identity" { return "Identity" }
    return ""
}

func (clock *Ptp_Clock) GetSegmentPath() string {
    return "clock"
}

func (clock *Ptp_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        return &clock.Profile
    }
    if childYangName == "identity" {
        return &clock.Identity
    }
    return nil
}

func (clock *Ptp_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profile"] = &clock.Profile
    children["identity"] = &clock.Identity
    return children
}

func (clock *Ptp_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timescale"] = clock.Timescale
    leafs["domain"] = clock.Domain
    leafs["priority2"] = clock.Priority2
    leafs["time-source"] = clock.TimeSource
    leafs["priority1"] = clock.Priority1
    leafs["clock-class"] = clock.ClockClass
    return leafs
}

func (clock *Ptp_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *Ptp_Clock) GetYangName() string { return "clock" }

func (clock *Ptp_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *Ptp_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *Ptp_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *Ptp_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *Ptp_Clock) GetParent() types.Entity { return clock.parent }

func (clock *Ptp_Clock) GetParentYangName() string { return "ptp" }

// Ptp_Clock_Profile
// Local clock PTP profile
type Ptp_Clock_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock profile. The type is PtpClockProfile. The default value is default.
    ClockProfile interface{}

    // Telecom clock type. The type is PtpTelecomClock.
    TelecomClockType interface{}
}

func (profile *Ptp_Clock_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ptp_Clock_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ptp_Clock_Profile) GetGoName(yname string) string {
    if yname == "clock-profile" { return "ClockProfile" }
    if yname == "telecom-clock-type" { return "TelecomClockType" }
    return ""
}

func (profile *Ptp_Clock_Profile) GetSegmentPath() string {
    return "profile"
}

func (profile *Ptp_Clock_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Ptp_Clock_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Ptp_Clock_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-profile"] = profile.ClockProfile
    leafs["telecom-clock-type"] = profile.TelecomClockType
    return leafs
}

func (profile *Ptp_Clock_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ptp_Clock_Profile) GetYangName() string { return "profile" }

func (profile *Ptp_Clock_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ptp_Clock_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ptp_Clock_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ptp_Clock_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ptp_Clock_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ptp_Clock_Profile) GetParentYangName() string { return "clock" }

// Ptp_Clock_Identity
// Local clock identity
type Ptp_Clock_Identity struct {
    parent types.Entity
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

func (identity *Ptp_Clock_Identity) GetFilter() yfilter.YFilter { return identity.YFilter }

func (identity *Ptp_Clock_Identity) SetFilter(yf yfilter.YFilter) { identity.YFilter = yf }

func (identity *Ptp_Clock_Identity) GetGoName(yname string) string {
    if yname == "clock-id-type" { return "ClockIdType" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "eui" { return "Eui" }
    return ""
}

func (identity *Ptp_Clock_Identity) GetSegmentPath() string {
    return "identity"
}

func (identity *Ptp_Clock_Identity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (identity *Ptp_Clock_Identity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (identity *Ptp_Clock_Identity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id-type"] = identity.ClockIdType
    leafs["mac-address"] = identity.MacAddress
    leafs["eui"] = identity.Eui
    return leafs
}

func (identity *Ptp_Clock_Identity) GetBundleName() string { return "cisco_ios_xr" }

func (identity *Ptp_Clock_Identity) GetYangName() string { return "identity" }

func (identity *Ptp_Clock_Identity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (identity *Ptp_Clock_Identity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (identity *Ptp_Clock_Identity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (identity *Ptp_Clock_Identity) SetParent(parent types.Entity) { identity.parent = parent }

func (identity *Ptp_Clock_Identity) GetParent() types.Entity { return identity.parent }

func (identity *Ptp_Clock_Identity) GetParentYangName() string { return "clock" }

// Ptp_Profiles
// Table for profile configuration
type Ptp_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Profile configuration. The type is slice of Ptp_Profiles_Profile.
    Profile []Ptp_Profiles_Profile
}

func (profiles *Ptp_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ptp_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ptp_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ptp_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ptp_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ptp_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ptp_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ptp_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ptp_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ptp_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ptp_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ptp_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ptp_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ptp_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ptp_Profiles) GetParentYangName() string { return "ptp" }

// Ptp_Profiles_Profile
// Profile configuration
type Ptp_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (profile *Ptp_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ptp_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ptp_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "sync-grant-duration" { return "SyncGrantDuration" }
    if yname == "general-cos" { return "GeneralCos" }
    if yname == "sync-timeout" { return "SyncTimeout" }
    if yname == "transport" { return "Transport" }
    if yname == "announce-timeout" { return "AnnounceTimeout" }
    if yname == "cos" { return "Cos" }
    if yname == "port-state" { return "PortState" }
    if yname == "delay-response-timeout" { return "DelayResponseTimeout" }
    if yname == "delay-response-grant-duration" { return "DelayResponseGrantDuration" }
    if yname == "event-cos" { return "EventCos" }
    if yname == "dscp" { return "Dscp" }
    if yname == "general-dscp" { return "GeneralDscp" }
    if yname == "clock-operation" { return "ClockOperation" }
    if yname == "announce-grant-duration" { return "AnnounceGrantDuration" }
    if yname == "unicast-grant-invalid-request" { return "UnicastGrantInvalidRequest" }
    if yname == "event-dscp" { return "EventDscp" }
    if yname == "announce-interval" { return "AnnounceInterval" }
    if yname == "source-ipv4-address" { return "SourceIpv4Address" }
    if yname == "slaves" { return "Slaves" }
    if yname == "sync-interval" { return "SyncInterval" }
    if yname == "masters" { return "Masters" }
    if yname == "communication" { return "Communication" }
    if yname == "delay-request-minimum-interval" { return "DelayRequestMinimumInterval" }
    if yname == "source-ipv6-address" { return "SourceIpv6Address" }
    return ""
}

func (profile *Ptp_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Ptp_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "announce-interval" {
        return &profile.AnnounceInterval
    }
    if childYangName == "source-ipv4-address" {
        return &profile.SourceIpv4Address
    }
    if childYangName == "slaves" {
        return &profile.Slaves
    }
    if childYangName == "sync-interval" {
        return &profile.SyncInterval
    }
    if childYangName == "masters" {
        return &profile.Masters
    }
    if childYangName == "communication" {
        return &profile.Communication
    }
    if childYangName == "delay-request-minimum-interval" {
        return &profile.DelayRequestMinimumInterval
    }
    if childYangName == "source-ipv6-address" {
        return &profile.SourceIpv6Address
    }
    return nil
}

func (profile *Ptp_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["announce-interval"] = &profile.AnnounceInterval
    children["source-ipv4-address"] = &profile.SourceIpv4Address
    children["slaves"] = &profile.Slaves
    children["sync-interval"] = &profile.SyncInterval
    children["masters"] = &profile.Masters
    children["communication"] = &profile.Communication
    children["delay-request-minimum-interval"] = &profile.DelayRequestMinimumInterval
    children["source-ipv6-address"] = &profile.SourceIpv6Address
    return children
}

func (profile *Ptp_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["sync-grant-duration"] = profile.SyncGrantDuration
    leafs["general-cos"] = profile.GeneralCos
    leafs["sync-timeout"] = profile.SyncTimeout
    leafs["transport"] = profile.Transport
    leafs["announce-timeout"] = profile.AnnounceTimeout
    leafs["cos"] = profile.Cos
    leafs["port-state"] = profile.PortState
    leafs["delay-response-timeout"] = profile.DelayResponseTimeout
    leafs["delay-response-grant-duration"] = profile.DelayResponseGrantDuration
    leafs["event-cos"] = profile.EventCos
    leafs["dscp"] = profile.Dscp
    leafs["general-dscp"] = profile.GeneralDscp
    leafs["clock-operation"] = profile.ClockOperation
    leafs["announce-grant-duration"] = profile.AnnounceGrantDuration
    leafs["unicast-grant-invalid-request"] = profile.UnicastGrantInvalidRequest
    leafs["event-dscp"] = profile.EventDscp
    return leafs
}

func (profile *Ptp_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ptp_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ptp_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ptp_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ptp_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ptp_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ptp_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ptp_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Ptp_Profiles_Profile_AnnounceInterval
// Announce interval
type Ptp_Profiles_Profile_AnnounceInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval or Frequency. The type is PtpTime. The default value is interval.
    TimeType interface{}

    // Time Period. The type is PtpTimePeriod. The default value is 2.
    TimePeriod interface{}
}

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetFilter() yfilter.YFilter { return announceInterval.YFilter }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) SetFilter(yf yfilter.YFilter) { announceInterval.YFilter = yf }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetGoName(yname string) string {
    if yname == "time-type" { return "TimeType" }
    if yname == "time-period" { return "TimePeriod" }
    return ""
}

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetSegmentPath() string {
    return "announce-interval"
}

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-type"] = announceInterval.TimeType
    leafs["time-period"] = announceInterval.TimePeriod
    return leafs
}

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetBundleName() string { return "cisco_ios_xr" }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetYangName() string { return "announce-interval" }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) SetParent(parent types.Entity) { announceInterval.parent = parent }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetParent() types.Entity { return announceInterval.parent }

func (announceInterval *Ptp_Profiles_Profile_AnnounceInterval) GetParentYangName() string { return "profile" }

// Ptp_Profiles_Profile_SourceIpv4Address
// Source IPv4 Address
type Ptp_Profiles_Profile_SourceIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable source IP address. The type is bool.
    Enable interface{}

    // Source IP address to use. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceIp interface{}
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetFilter() yfilter.YFilter { return sourceIpv4Address.YFilter }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) SetFilter(yf yfilter.YFilter) { sourceIpv4Address.YFilter = yf }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source-ip" { return "SourceIp" }
    return ""
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetSegmentPath() string {
    return "source-ipv4-address"
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = sourceIpv4Address.Enable
    leafs["source-ip"] = sourceIpv4Address.SourceIp
    return leafs
}

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetYangName() string { return "source-ipv4-address" }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) SetParent(parent types.Entity) { sourceIpv4Address.parent = parent }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetParent() types.Entity { return sourceIpv4Address.parent }

func (sourceIpv4Address *Ptp_Profiles_Profile_SourceIpv4Address) GetParentYangName() string { return "profile" }

// Ptp_Profiles_Profile_Slaves
// Table for slave configuration
type Ptp_Profiles_Profile_Slaves struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slave configuration. The type is slice of
    // Ptp_Profiles_Profile_Slaves_Slave.
    Slave []Ptp_Profiles_Profile_Slaves_Slave
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetFilter() yfilter.YFilter { return slaves.YFilter }

func (slaves *Ptp_Profiles_Profile_Slaves) SetFilter(yf yfilter.YFilter) { slaves.YFilter = yf }

func (slaves *Ptp_Profiles_Profile_Slaves) GetGoName(yname string) string {
    if yname == "slave" { return "Slave" }
    return ""
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetSegmentPath() string {
    return "slaves"
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave" {
        for _, c := range slaves.Slave {
            if slaves.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Profiles_Profile_Slaves_Slave{}
        slaves.Slave = append(slaves.Slave, child)
        return &slaves.Slave[len(slaves.Slave)-1]
    }
    return nil
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slaves.Slave {
        children[slaves.Slave[i].GetSegmentPath()] = &slaves.Slave[i]
    }
    return children
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slaves *Ptp_Profiles_Profile_Slaves) GetBundleName() string { return "cisco_ios_xr" }

func (slaves *Ptp_Profiles_Profile_Slaves) GetYangName() string { return "slaves" }

func (slaves *Ptp_Profiles_Profile_Slaves) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaves *Ptp_Profiles_Profile_Slaves) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaves *Ptp_Profiles_Profile_Slaves) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaves *Ptp_Profiles_Profile_Slaves) SetParent(parent types.Entity) { slaves.parent = parent }

func (slaves *Ptp_Profiles_Profile_Slaves) GetParent() types.Entity { return slaves.parent }

func (slaves *Ptp_Profiles_Profile_Slaves) GetParentYangName() string { return "profile" }

// Ptp_Profiles_Profile_Slaves_Slave
// Slave configuration
type Ptp_Profiles_Profile_Slaves_Slave struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slave Transport Type. The type is PtpEncap.
    Transport interface{}

    // ethernet. The type is slice of Ptp_Profiles_Profile_Slaves_Slave_Ethernet.
    Ethernet []Ptp_Profiles_Profile_Slaves_Slave_Ethernet

    // ipv4 or ipv6. The type is slice of
    // Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6.
    Ipv4OrIpv6 []Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetFilter() yfilter.YFilter { return slave.YFilter }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) SetFilter(yf yfilter.YFilter) { slave.YFilter = yf }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetGoName(yname string) string {
    if yname == "transport" { return "Transport" }
    if yname == "ethernet" { return "Ethernet" }
    if yname == "ipv4-or-ipv6" { return "Ipv4OrIpv6" }
    return ""
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetSegmentPath() string {
    return "slave" + "[transport='" + fmt.Sprintf("%v", slave.Transport) + "']"
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet" {
        for _, c := range slave.Ethernet {
            if slave.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Profiles_Profile_Slaves_Slave_Ethernet{}
        slave.Ethernet = append(slave.Ethernet, child)
        return &slave.Ethernet[len(slave.Ethernet)-1]
    }
    if childYangName == "ipv4-or-ipv6" {
        for _, c := range slave.Ipv4OrIpv6 {
            if slave.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6{}
        slave.Ipv4OrIpv6 = append(slave.Ipv4OrIpv6, child)
        return &slave.Ipv4OrIpv6[len(slave.Ipv4OrIpv6)-1]
    }
    return nil
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slave.Ethernet {
        children[slave.Ethernet[i].GetSegmentPath()] = &slave.Ethernet[i]
    }
    for i := range slave.Ipv4OrIpv6 {
        children[slave.Ipv4OrIpv6[i].GetSegmentPath()] = &slave.Ipv4OrIpv6[i]
    }
    return children
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport"] = slave.Transport
    return leafs
}

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetBundleName() string { return "cisco_ios_xr" }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetYangName() string { return "slave" }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) SetParent(parent types.Entity) { slave.parent = parent }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetParent() types.Entity { return slave.parent }

func (slave *Ptp_Profiles_Profile_Slaves_Slave) GetParentYangName() string { return "slaves" }

// Ptp_Profiles_Profile_Slaves_Slave_Ethernet
// ethernet
type Ptp_Profiles_Profile_Slaves_Slave_Ethernet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slave MAC Address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SlaveMacAddress interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetFilter() yfilter.YFilter { return ethernet.YFilter }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) SetFilter(yf yfilter.YFilter) { ethernet.YFilter = yf }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetGoName(yname string) string {
    if yname == "slave-mac-address" { return "SlaveMacAddress" }
    if yname == "non-negotiated" { return "NonNegotiated" }
    return ""
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetSegmentPath() string {
    return "ethernet" + "[slave-mac-address='" + fmt.Sprintf("%v", ethernet.SlaveMacAddress) + "']"
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-mac-address"] = ethernet.SlaveMacAddress
    leafs["non-negotiated"] = ethernet.NonNegotiated
    return leafs
}

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetBundleName() string { return "cisco_ios_xr" }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetYangName() string { return "ethernet" }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) SetParent(parent types.Entity) { ethernet.parent = parent }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetParent() types.Entity { return ethernet.parent }

func (ethernet *Ptp_Profiles_Profile_Slaves_Slave_Ethernet) GetParentYangName() string { return "slave" }

// Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6
// ipv4 or ipv6
type Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slave IP Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SlaveIpAddress interface{}

    // Enable non-negotiated unicast on this interface. The type is bool.
    NonNegotiated interface{}
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetFilter() yfilter.YFilter { return ipv4OrIpv6.YFilter }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) SetFilter(yf yfilter.YFilter) { ipv4OrIpv6.YFilter = yf }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetGoName(yname string) string {
    if yname == "slave-ip-address" { return "SlaveIpAddress" }
    if yname == "non-negotiated" { return "NonNegotiated" }
    return ""
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetSegmentPath() string {
    return "ipv4-or-ipv6" + "[slave-ip-address='" + fmt.Sprintf("%v", ipv4OrIpv6.SlaveIpAddress) + "']"
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-ip-address"] = ipv4OrIpv6.SlaveIpAddress
    leafs["non-negotiated"] = ipv4OrIpv6.NonNegotiated
    return leafs
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetYangName() string { return "ipv4-or-ipv6" }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) SetParent(parent types.Entity) { ipv4OrIpv6.parent = parent }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetParent() types.Entity { return ipv4OrIpv6.parent }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Slaves_Slave_Ipv4OrIpv6) GetParentYangName() string { return "slave" }

// Ptp_Profiles_Profile_SyncInterval
// Sync interval
type Ptp_Profiles_Profile_SyncInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval or Frequency. The type is PtpTime. The default value is interval.
    TimeType interface{}

    // Time Period. The type is PtpTimePeriod. The default value is 1.
    TimePeriod interface{}
}

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetFilter() yfilter.YFilter { return syncInterval.YFilter }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) SetFilter(yf yfilter.YFilter) { syncInterval.YFilter = yf }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetGoName(yname string) string {
    if yname == "time-type" { return "TimeType" }
    if yname == "time-period" { return "TimePeriod" }
    return ""
}

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetSegmentPath() string {
    return "sync-interval"
}

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-type"] = syncInterval.TimeType
    leafs["time-period"] = syncInterval.TimePeriod
    return leafs
}

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetBundleName() string { return "cisco_ios_xr" }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetYangName() string { return "sync-interval" }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) SetParent(parent types.Entity) { syncInterval.parent = parent }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetParent() types.Entity { return syncInterval.parent }

func (syncInterval *Ptp_Profiles_Profile_SyncInterval) GetParentYangName() string { return "profile" }

// Ptp_Profiles_Profile_Masters
// Table for master configuration
type Ptp_Profiles_Profile_Masters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Master configuration. The type is slice of
    // Ptp_Profiles_Profile_Masters_Master.
    Master []Ptp_Profiles_Profile_Masters_Master
}

func (masters *Ptp_Profiles_Profile_Masters) GetFilter() yfilter.YFilter { return masters.YFilter }

func (masters *Ptp_Profiles_Profile_Masters) SetFilter(yf yfilter.YFilter) { masters.YFilter = yf }

func (masters *Ptp_Profiles_Profile_Masters) GetGoName(yname string) string {
    if yname == "master" { return "Master" }
    return ""
}

func (masters *Ptp_Profiles_Profile_Masters) GetSegmentPath() string {
    return "masters"
}

func (masters *Ptp_Profiles_Profile_Masters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "master" {
        for _, c := range masters.Master {
            if masters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Profiles_Profile_Masters_Master{}
        masters.Master = append(masters.Master, child)
        return &masters.Master[len(masters.Master)-1]
    }
    return nil
}

func (masters *Ptp_Profiles_Profile_Masters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range masters.Master {
        children[masters.Master[i].GetSegmentPath()] = &masters.Master[i]
    }
    return children
}

func (masters *Ptp_Profiles_Profile_Masters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (masters *Ptp_Profiles_Profile_Masters) GetBundleName() string { return "cisco_ios_xr" }

func (masters *Ptp_Profiles_Profile_Masters) GetYangName() string { return "masters" }

func (masters *Ptp_Profiles_Profile_Masters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (masters *Ptp_Profiles_Profile_Masters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (masters *Ptp_Profiles_Profile_Masters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (masters *Ptp_Profiles_Profile_Masters) SetParent(parent types.Entity) { masters.parent = parent }

func (masters *Ptp_Profiles_Profile_Masters) GetParent() types.Entity { return masters.parent }

func (masters *Ptp_Profiles_Profile_Masters) GetParentYangName() string { return "profile" }

// Ptp_Profiles_Profile_Masters_Master
// Master configuration
type Ptp_Profiles_Profile_Masters_Master struct {
    parent types.Entity
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

func (master *Ptp_Profiles_Profile_Masters_Master) GetFilter() yfilter.YFilter { return master.YFilter }

func (master *Ptp_Profiles_Profile_Masters_Master) SetFilter(yf yfilter.YFilter) { master.YFilter = yf }

func (master *Ptp_Profiles_Profile_Masters_Master) GetGoName(yname string) string {
    if yname == "transport" { return "Transport" }
    if yname == "ethernet" { return "Ethernet" }
    if yname == "ipv4-or-ipv6" { return "Ipv4OrIpv6" }
    return ""
}

func (master *Ptp_Profiles_Profile_Masters_Master) GetSegmentPath() string {
    return "master" + "[transport='" + fmt.Sprintf("%v", master.Transport) + "']"
}

func (master *Ptp_Profiles_Profile_Masters_Master) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet" {
        for _, c := range master.Ethernet {
            if master.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Profiles_Profile_Masters_Master_Ethernet{}
        master.Ethernet = append(master.Ethernet, child)
        return &master.Ethernet[len(master.Ethernet)-1]
    }
    if childYangName == "ipv4-or-ipv6" {
        for _, c := range master.Ipv4OrIpv6 {
            if master.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6{}
        master.Ipv4OrIpv6 = append(master.Ipv4OrIpv6, child)
        return &master.Ipv4OrIpv6[len(master.Ipv4OrIpv6)-1]
    }
    return nil
}

func (master *Ptp_Profiles_Profile_Masters_Master) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range master.Ethernet {
        children[master.Ethernet[i].GetSegmentPath()] = &master.Ethernet[i]
    }
    for i := range master.Ipv4OrIpv6 {
        children[master.Ipv4OrIpv6[i].GetSegmentPath()] = &master.Ipv4OrIpv6[i]
    }
    return children
}

func (master *Ptp_Profiles_Profile_Masters_Master) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport"] = master.Transport
    return leafs
}

func (master *Ptp_Profiles_Profile_Masters_Master) GetBundleName() string { return "cisco_ios_xr" }

func (master *Ptp_Profiles_Profile_Masters_Master) GetYangName() string { return "master" }

func (master *Ptp_Profiles_Profile_Masters_Master) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (master *Ptp_Profiles_Profile_Masters_Master) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (master *Ptp_Profiles_Profile_Masters_Master) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (master *Ptp_Profiles_Profile_Masters_Master) SetParent(parent types.Entity) { master.parent = parent }

func (master *Ptp_Profiles_Profile_Masters_Master) GetParent() types.Entity { return master.parent }

func (master *Ptp_Profiles_Profile_Masters_Master) GetParentYangName() string { return "masters" }

// Ptp_Profiles_Profile_Masters_Master_Ethernet
// ethernet
type Ptp_Profiles_Profile_Masters_Master_Ethernet struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetFilter() yfilter.YFilter { return ethernet.YFilter }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) SetFilter(yf yfilter.YFilter) { ethernet.YFilter = yf }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetGoName(yname string) string {
    if yname == "master-mac-address" { return "MasterMacAddress" }
    if yname == "master-clock-class" { return "MasterClockClass" }
    if yname == "non-negotiated" { return "NonNegotiated" }
    if yname == "priority" { return "Priority" }
    if yname == "communication" { return "Communication" }
    if yname == "delay-asymmetry" { return "DelayAsymmetry" }
    return ""
}

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetSegmentPath() string {
    return "ethernet" + "[master-mac-address='" + fmt.Sprintf("%v", ethernet.MasterMacAddress) + "']"
}

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delay-asymmetry" {
        return &ethernet.DelayAsymmetry
    }
    return nil
}

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delay-asymmetry"] = &ethernet.DelayAsymmetry
    return children
}

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["master-mac-address"] = ethernet.MasterMacAddress
    leafs["master-clock-class"] = ethernet.MasterClockClass
    leafs["non-negotiated"] = ethernet.NonNegotiated
    leafs["priority"] = ethernet.Priority
    leafs["communication"] = ethernet.Communication
    return leafs
}

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetBundleName() string { return "cisco_ios_xr" }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetYangName() string { return "ethernet" }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) SetParent(parent types.Entity) { ethernet.parent = parent }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetParent() types.Entity { return ethernet.parent }

func (ethernet *Ptp_Profiles_Profile_Masters_Master_Ethernet) GetParentYangName() string { return "master" }

// Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry
// The delay asymmetry for this master
// This type is a presence type.
type Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // How much longer the master to slave path takes than the reverse. The type
    // is interface{} with range: -500000000..500000000. This attribute is
    // mandatory.
    Magnitude interface{}

    // The units to use for the delay asymmetry. The type is
    // PtpDelayAsymmetryUnits. This attribute is mandatory.
    Units interface{}
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetFilter() yfilter.YFilter { return delayAsymmetry.YFilter }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) SetFilter(yf yfilter.YFilter) { delayAsymmetry.YFilter = yf }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetGoName(yname string) string {
    if yname == "magnitude" { return "Magnitude" }
    if yname == "units" { return "Units" }
    return ""
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetSegmentPath() string {
    return "delay-asymmetry"
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["magnitude"] = delayAsymmetry.Magnitude
    leafs["units"] = delayAsymmetry.Units
    return leafs
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetBundleName() string { return "cisco_ios_xr" }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetYangName() string { return "delay-asymmetry" }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) SetParent(parent types.Entity) { delayAsymmetry.parent = parent }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetParent() types.Entity { return delayAsymmetry.parent }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ethernet_DelayAsymmetry) GetParentYangName() string { return "ethernet" }

// Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6
// ipv4 or ipv6
type Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetFilter() yfilter.YFilter { return ipv4OrIpv6.YFilter }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) SetFilter(yf yfilter.YFilter) { ipv4OrIpv6.YFilter = yf }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetGoName(yname string) string {
    if yname == "master-ip-address" { return "MasterIpAddress" }
    if yname == "master-clock-class" { return "MasterClockClass" }
    if yname == "non-negotiated" { return "NonNegotiated" }
    if yname == "priority" { return "Priority" }
    if yname == "communication" { return "Communication" }
    if yname == "delay-asymmetry" { return "DelayAsymmetry" }
    return ""
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetSegmentPath() string {
    return "ipv4-or-ipv6" + "[master-ip-address='" + fmt.Sprintf("%v", ipv4OrIpv6.MasterIpAddress) + "']"
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delay-asymmetry" {
        return &ipv4OrIpv6.DelayAsymmetry
    }
    return nil
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delay-asymmetry"] = &ipv4OrIpv6.DelayAsymmetry
    return children
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["master-ip-address"] = ipv4OrIpv6.MasterIpAddress
    leafs["master-clock-class"] = ipv4OrIpv6.MasterClockClass
    leafs["non-negotiated"] = ipv4OrIpv6.NonNegotiated
    leafs["priority"] = ipv4OrIpv6.Priority
    leafs["communication"] = ipv4OrIpv6.Communication
    return leafs
}

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetYangName() string { return "ipv4-or-ipv6" }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) SetParent(parent types.Entity) { ipv4OrIpv6.parent = parent }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetParent() types.Entity { return ipv4OrIpv6.parent }

func (ipv4OrIpv6 *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6) GetParentYangName() string { return "master" }

// Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry
// The delay asymmetry for this master
// This type is a presence type.
type Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // How much longer the master to slave path takes than the reverse. The type
    // is interface{} with range: -500000000..500000000. This attribute is
    // mandatory.
    Magnitude interface{}

    // The units to use for the delay asymmetry. The type is
    // PtpDelayAsymmetryUnits. This attribute is mandatory.
    Units interface{}
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetFilter() yfilter.YFilter { return delayAsymmetry.YFilter }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) SetFilter(yf yfilter.YFilter) { delayAsymmetry.YFilter = yf }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetGoName(yname string) string {
    if yname == "magnitude" { return "Magnitude" }
    if yname == "units" { return "Units" }
    return ""
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetSegmentPath() string {
    return "delay-asymmetry"
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["magnitude"] = delayAsymmetry.Magnitude
    leafs["units"] = delayAsymmetry.Units
    return leafs
}

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetBundleName() string { return "cisco_ios_xr" }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetYangName() string { return "delay-asymmetry" }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) SetParent(parent types.Entity) { delayAsymmetry.parent = parent }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetParent() types.Entity { return delayAsymmetry.parent }

func (delayAsymmetry *Ptp_Profiles_Profile_Masters_Master_Ipv4OrIpv6_DelayAsymmetry) GetParentYangName() string { return "ipv4-or-ipv6" }

// Ptp_Profiles_Profile_Communication
// Communication model
type Ptp_Profiles_Profile_Communication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Communication Model. The type is PtpTransport. The default value is
    // unicast.
    Model interface{}

    // Target address set. The type is bool. The default value is false.
    TargetAddressSet interface{}

    // Target address. The type is string.
    TargetAddress interface{}
}

func (communication *Ptp_Profiles_Profile_Communication) GetFilter() yfilter.YFilter { return communication.YFilter }

func (communication *Ptp_Profiles_Profile_Communication) SetFilter(yf yfilter.YFilter) { communication.YFilter = yf }

func (communication *Ptp_Profiles_Profile_Communication) GetGoName(yname string) string {
    if yname == "model" { return "Model" }
    if yname == "target-address-set" { return "TargetAddressSet" }
    if yname == "target-address" { return "TargetAddress" }
    return ""
}

func (communication *Ptp_Profiles_Profile_Communication) GetSegmentPath() string {
    return "communication"
}

func (communication *Ptp_Profiles_Profile_Communication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (communication *Ptp_Profiles_Profile_Communication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (communication *Ptp_Profiles_Profile_Communication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["model"] = communication.Model
    leafs["target-address-set"] = communication.TargetAddressSet
    leafs["target-address"] = communication.TargetAddress
    return leafs
}

func (communication *Ptp_Profiles_Profile_Communication) GetBundleName() string { return "cisco_ios_xr" }

func (communication *Ptp_Profiles_Profile_Communication) GetYangName() string { return "communication" }

func (communication *Ptp_Profiles_Profile_Communication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (communication *Ptp_Profiles_Profile_Communication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (communication *Ptp_Profiles_Profile_Communication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (communication *Ptp_Profiles_Profile_Communication) SetParent(parent types.Entity) { communication.parent = parent }

func (communication *Ptp_Profiles_Profile_Communication) GetParent() types.Entity { return communication.parent }

func (communication *Ptp_Profiles_Profile_Communication) GetParentYangName() string { return "profile" }

// Ptp_Profiles_Profile_DelayRequestMinimumInterval
// Minimum delay request interval
type Ptp_Profiles_Profile_DelayRequestMinimumInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval or Frequency. The type is PtpTime. The default value is interval.
    TimeType interface{}

    // Time Period. The type is PtpTimePeriod. The default value is 1.
    TimePeriod interface{}
}

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetFilter() yfilter.YFilter { return delayRequestMinimumInterval.YFilter }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) SetFilter(yf yfilter.YFilter) { delayRequestMinimumInterval.YFilter = yf }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetGoName(yname string) string {
    if yname == "time-type" { return "TimeType" }
    if yname == "time-period" { return "TimePeriod" }
    return ""
}

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetSegmentPath() string {
    return "delay-request-minimum-interval"
}

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-type"] = delayRequestMinimumInterval.TimeType
    leafs["time-period"] = delayRequestMinimumInterval.TimePeriod
    return leafs
}

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetBundleName() string { return "cisco_ios_xr" }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetYangName() string { return "delay-request-minimum-interval" }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) SetParent(parent types.Entity) { delayRequestMinimumInterval.parent = parent }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetParent() types.Entity { return delayRequestMinimumInterval.parent }

func (delayRequestMinimumInterval *Ptp_Profiles_Profile_DelayRequestMinimumInterval) GetParentYangName() string { return "profile" }

// Ptp_Profiles_Profile_SourceIpv6Address
// Source IPv6 Address
type Ptp_Profiles_Profile_SourceIpv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable source IPv6 address. The type is bool.
    Enable interface{}

    // Source IPv6 address to use. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceIpv6 interface{}
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetFilter() yfilter.YFilter { return sourceIpv6Address.YFilter }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) SetFilter(yf yfilter.YFilter) { sourceIpv6Address.YFilter = yf }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "source-ipv6" { return "SourceIpv6" }
    return ""
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetSegmentPath() string {
    return "source-ipv6-address"
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = sourceIpv6Address.Enable
    leafs["source-ipv6"] = sourceIpv6Address.SourceIpv6
    return leafs
}

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetYangName() string { return "source-ipv6-address" }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) SetParent(parent types.Entity) { sourceIpv6Address.parent = parent }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetParent() types.Entity { return sourceIpv6Address.parent }

func (sourceIpv6Address *Ptp_Profiles_Profile_SourceIpv6Address) GetParentYangName() string { return "profile" }

// Ptp_Logging
// PTP logging configuration
type Ptp_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PTP best master clock logging configuration.
    BestMasterClock Ptp_Logging_BestMasterClock

    // PTP PD Servo logging configuration.
    Servo Ptp_Logging_Servo
}

func (logging *Ptp_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Ptp_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Ptp_Logging) GetGoName(yname string) string {
    if yname == "best-master-clock" { return "BestMasterClock" }
    if yname == "Cisco-IOS-XR-asr9k-ptp-pd-cfg:servo" { return "Servo" }
    return ""
}

func (logging *Ptp_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *Ptp_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "best-master-clock" {
        return &logging.BestMasterClock
    }
    if childYangName == "Cisco-IOS-XR-asr9k-ptp-pd-cfg:servo" {
        return &logging.Servo
    }
    return nil
}

func (logging *Ptp_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["best-master-clock"] = &logging.BestMasterClock
    children["Cisco-IOS-XR-asr9k-ptp-pd-cfg:servo"] = &logging.Servo
    return children
}

func (logging *Ptp_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logging *Ptp_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Ptp_Logging) GetYangName() string { return "logging" }

func (logging *Ptp_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Ptp_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Ptp_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Ptp_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Ptp_Logging) GetParent() types.Entity { return logging.parent }

func (logging *Ptp_Logging) GetParentYangName() string { return "ptp" }

// Ptp_Logging_BestMasterClock
// PTP best master clock logging configuration
type Ptp_Logging_BestMasterClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable best master clock changes logging. The type is interface{}.
    Changes interface{}
}

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetFilter() yfilter.YFilter { return bestMasterClock.YFilter }

func (bestMasterClock *Ptp_Logging_BestMasterClock) SetFilter(yf yfilter.YFilter) { bestMasterClock.YFilter = yf }

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetGoName(yname string) string {
    if yname == "changes" { return "Changes" }
    return ""
}

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetSegmentPath() string {
    return "best-master-clock"
}

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["changes"] = bestMasterClock.Changes
    return leafs
}

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetBundleName() string { return "cisco_ios_xr" }

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetYangName() string { return "best-master-clock" }

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bestMasterClock *Ptp_Logging_BestMasterClock) SetParent(parent types.Entity) { bestMasterClock.parent = parent }

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetParent() types.Entity { return bestMasterClock.parent }

func (bestMasterClock *Ptp_Logging_BestMasterClock) GetParentYangName() string { return "logging" }

// Ptp_Logging_Servo
// PTP PD Servo logging configuration
type Ptp_Logging_Servo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable servo events logging. The type is interface{}.
    Events interface{}
}

func (servo *Ptp_Logging_Servo) GetFilter() yfilter.YFilter { return servo.YFilter }

func (servo *Ptp_Logging_Servo) SetFilter(yf yfilter.YFilter) { servo.YFilter = yf }

func (servo *Ptp_Logging_Servo) GetGoName(yname string) string {
    if yname == "events" { return "Events" }
    return ""
}

func (servo *Ptp_Logging_Servo) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-ptp-pd-cfg:servo"
}

func (servo *Ptp_Logging_Servo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (servo *Ptp_Logging_Servo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (servo *Ptp_Logging_Servo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["events"] = servo.Events
    return leafs
}

func (servo *Ptp_Logging_Servo) GetBundleName() string { return "cisco_ios_xr" }

func (servo *Ptp_Logging_Servo) GetYangName() string { return "servo" }

func (servo *Ptp_Logging_Servo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servo *Ptp_Logging_Servo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servo *Ptp_Logging_Servo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servo *Ptp_Logging_Servo) SetParent(parent types.Entity) { servo.parent = parent }

func (servo *Ptp_Logging_Servo) GetParent() types.Entity { return servo.parent }

func (servo *Ptp_Logging_Servo) GetParentYangName() string { return "logging" }

// Ptp_TransparentClock
// Transparent clock configuration
type Ptp_TransparentClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of domains containing transparent clock configuration.
    Domains Ptp_TransparentClock_Domains
}

func (transparentClock *Ptp_TransparentClock) GetFilter() yfilter.YFilter { return transparentClock.YFilter }

func (transparentClock *Ptp_TransparentClock) SetFilter(yf yfilter.YFilter) { transparentClock.YFilter = yf }

func (transparentClock *Ptp_TransparentClock) GetGoName(yname string) string {
    if yname == "domains" { return "Domains" }
    return ""
}

func (transparentClock *Ptp_TransparentClock) GetSegmentPath() string {
    return "transparent-clock"
}

func (transparentClock *Ptp_TransparentClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "domains" {
        return &transparentClock.Domains
    }
    return nil
}

func (transparentClock *Ptp_TransparentClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["domains"] = &transparentClock.Domains
    return children
}

func (transparentClock *Ptp_TransparentClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transparentClock *Ptp_TransparentClock) GetBundleName() string { return "cisco_ios_xr" }

func (transparentClock *Ptp_TransparentClock) GetYangName() string { return "transparent-clock" }

func (transparentClock *Ptp_TransparentClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transparentClock *Ptp_TransparentClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transparentClock *Ptp_TransparentClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transparentClock *Ptp_TransparentClock) SetParent(parent types.Entity) { transparentClock.parent = parent }

func (transparentClock *Ptp_TransparentClock) GetParent() types.Entity { return transparentClock.parent }

func (transparentClock *Ptp_TransparentClock) GetParentYangName() string { return "ptp" }

// Ptp_TransparentClock_Domains
// Table of domains containing transparent clock
// configuration
type Ptp_TransparentClock_Domains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transparent clock domain configuration. The type is slice of
    // Ptp_TransparentClock_Domains_Domain.
    Domain []Ptp_TransparentClock_Domains_Domain
}

func (domains *Ptp_TransparentClock_Domains) GetFilter() yfilter.YFilter { return domains.YFilter }

func (domains *Ptp_TransparentClock_Domains) SetFilter(yf yfilter.YFilter) { domains.YFilter = yf }

func (domains *Ptp_TransparentClock_Domains) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    return ""
}

func (domains *Ptp_TransparentClock_Domains) GetSegmentPath() string {
    return "domains"
}

func (domains *Ptp_TransparentClock_Domains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "domain" {
        for _, c := range domains.Domain {
            if domains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_TransparentClock_Domains_Domain{}
        domains.Domain = append(domains.Domain, child)
        return &domains.Domain[len(domains.Domain)-1]
    }
    return nil
}

func (domains *Ptp_TransparentClock_Domains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range domains.Domain {
        children[domains.Domain[i].GetSegmentPath()] = &domains.Domain[i]
    }
    return children
}

func (domains *Ptp_TransparentClock_Domains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (domains *Ptp_TransparentClock_Domains) GetBundleName() string { return "cisco_ios_xr" }

func (domains *Ptp_TransparentClock_Domains) GetYangName() string { return "domains" }

func (domains *Ptp_TransparentClock_Domains) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domains *Ptp_TransparentClock_Domains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domains *Ptp_TransparentClock_Domains) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domains *Ptp_TransparentClock_Domains) SetParent(parent types.Entity) { domains.parent = parent }

func (domains *Ptp_TransparentClock_Domains) GetParent() types.Entity { return domains.parent }

func (domains *Ptp_TransparentClock_Domains) GetParentYangName() string { return "transparent-clock" }

// Ptp_TransparentClock_Domains_Domain
// Transparent clock domain configuration
type Ptp_TransparentClock_Domains_Domain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Domain. The type is string with pattern: (all).
    Domain interface{}
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetFilter() yfilter.YFilter { return domain.YFilter }

func (domain *Ptp_TransparentClock_Domains_Domain) SetFilter(yf yfilter.YFilter) { domain.YFilter = yf }

func (domain *Ptp_TransparentClock_Domains_Domain) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    return ""
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetSegmentPath() string {
    return "domain" + "[domain='" + fmt.Sprintf("%v", domain.Domain) + "']"
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = domain.Domain
    return leafs
}

func (domain *Ptp_TransparentClock_Domains_Domain) GetBundleName() string { return "cisco_ios_xr" }

func (domain *Ptp_TransparentClock_Domains_Domain) GetYangName() string { return "domain" }

func (domain *Ptp_TransparentClock_Domains_Domain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domain *Ptp_TransparentClock_Domains_Domain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domain *Ptp_TransparentClock_Domains_Domain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domain *Ptp_TransparentClock_Domains_Domain) SetParent(parent types.Entity) { domain.parent = parent }

func (domain *Ptp_TransparentClock_Domains_Domain) GetParent() types.Entity { return domain.parent }

func (domain *Ptp_TransparentClock_Domains_Domain) GetParentYangName() string { return "domains" }

