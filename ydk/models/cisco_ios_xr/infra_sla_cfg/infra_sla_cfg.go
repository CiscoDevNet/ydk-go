// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-sla package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sla: SLA prtocol and profile Configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_sla_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_sla_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-sla-cfg sla}", reflect.TypeOf(Sla{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-sla-cfg:sla", reflect.TypeOf(Sla{}))
}

// Sla
// SLA prtocol and profile Configuration
type Sla struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of all SLA protocols.
    Protocols Sla_Protocols
}

func (sla *Sla) GetFilter() yfilter.YFilter { return sla.YFilter }

func (sla *Sla) SetFilter(yf yfilter.YFilter) { sla.YFilter = yf }

func (sla *Sla) GetGoName(yname string) string {
    if yname == "protocols" { return "Protocols" }
    return ""
}

func (sla *Sla) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-sla-cfg:sla"
}

func (sla *Sla) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocols" {
        return &sla.Protocols
    }
    return nil
}

func (sla *Sla) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocols"] = &sla.Protocols
    return children
}

func (sla *Sla) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sla *Sla) GetBundleName() string { return "cisco_ios_xr" }

func (sla *Sla) GetYangName() string { return "sla" }

func (sla *Sla) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sla *Sla) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sla *Sla) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sla *Sla) SetParent(parent types.Entity) { sla.parent = parent }

func (sla *Sla) GetParent() types.Entity { return sla.parent }

func (sla *Sla) GetParentYangName() string { return "Cisco-IOS-XR-infra-sla-cfg" }

// Sla_Protocols
// Table of all SLA protocols
type Sla_Protocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Ethernet SLA protocol.
    Ethernet Sla_Protocols_Ethernet
}

func (protocols *Sla_Protocols) GetFilter() yfilter.YFilter { return protocols.YFilter }

func (protocols *Sla_Protocols) SetFilter(yf yfilter.YFilter) { protocols.YFilter = yf }

func (protocols *Sla_Protocols) GetGoName(yname string) string {
    if yname == "Cisco-IOS-XR-ethernet-cfm-cfg:ethernet" { return "Ethernet" }
    return ""
}

func (protocols *Sla_Protocols) GetSegmentPath() string {
    return "protocols"
}

func (protocols *Sla_Protocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-ethernet-cfm-cfg:ethernet" {
        return &protocols.Ethernet
    }
    return nil
}

func (protocols *Sla_Protocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-ethernet-cfm-cfg:ethernet"] = &protocols.Ethernet
    return children
}

func (protocols *Sla_Protocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocols *Sla_Protocols) GetBundleName() string { return "cisco_ios_xr" }

func (protocols *Sla_Protocols) GetYangName() string { return "protocols" }

func (protocols *Sla_Protocols) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocols *Sla_Protocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocols *Sla_Protocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocols *Sla_Protocols) SetParent(parent types.Entity) { protocols.parent = parent }

func (protocols *Sla_Protocols) GetParent() types.Entity { return protocols.parent }

func (protocols *Sla_Protocols) GetParentYangName() string { return "sla" }

// Sla_Protocols_Ethernet
// The Ethernet SLA protocol
type Sla_Protocols_Ethernet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of SLA profiles on the protocol.
    Profiles Sla_Protocols_Ethernet_Profiles
}

func (ethernet *Sla_Protocols_Ethernet) GetFilter() yfilter.YFilter { return ethernet.YFilter }

func (ethernet *Sla_Protocols_Ethernet) SetFilter(yf yfilter.YFilter) { ethernet.YFilter = yf }

func (ethernet *Sla_Protocols_Ethernet) GetGoName(yname string) string {
    if yname == "profiles" { return "Profiles" }
    return ""
}

func (ethernet *Sla_Protocols_Ethernet) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-cfm-cfg:ethernet"
}

func (ethernet *Sla_Protocols_Ethernet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profiles" {
        return &ethernet.Profiles
    }
    return nil
}

func (ethernet *Sla_Protocols_Ethernet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profiles"] = &ethernet.Profiles
    return children
}

func (ethernet *Sla_Protocols_Ethernet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernet *Sla_Protocols_Ethernet) GetBundleName() string { return "cisco_ios_xr" }

func (ethernet *Sla_Protocols_Ethernet) GetYangName() string { return "ethernet" }

func (ethernet *Sla_Protocols_Ethernet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernet *Sla_Protocols_Ethernet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernet *Sla_Protocols_Ethernet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernet *Sla_Protocols_Ethernet) SetParent(parent types.Entity) { ethernet.parent = parent }

func (ethernet *Sla_Protocols_Ethernet) GetParent() types.Entity { return ethernet.parent }

func (ethernet *Sla_Protocols_Ethernet) GetParentYangName() string { return "protocols" }

// Sla_Protocols_Ethernet_Profiles
// Table of SLA profiles on the protocol
type Sla_Protocols_Ethernet_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile. The type is slice of
    // Sla_Protocols_Ethernet_Profiles_Profile.
    Profile []Sla_Protocols_Ethernet_Profiles_Profile
}

func (profiles *Sla_Protocols_Ethernet_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Sla_Protocols_Ethernet_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Sla_Protocols_Ethernet_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Sla_Protocols_Ethernet_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Sla_Protocols_Ethernet_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Sla_Protocols_Ethernet_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Sla_Protocols_Ethernet_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Sla_Protocols_Ethernet_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Sla_Protocols_Ethernet_Profiles) GetYangName() string { return "profiles" }

func (profiles *Sla_Protocols_Ethernet_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Sla_Protocols_Ethernet_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Sla_Protocols_Ethernet_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Sla_Protocols_Ethernet_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Sla_Protocols_Ethernet_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Sla_Protocols_Ethernet_Profiles) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_Profiles_Profile
// Name of the profile
type Sla_Protocols_Ethernet_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // The possible packet types are cfm-loopback, cfm-delay-measurement,
    // cfm-delay-measurement-version-0, cfm-loss-measurement and
    // cfm-synthetic-loss-measurement. The type is string.
    PacketType interface{}

    // Statistics configuration for the SLA profile.
    Statistics Sla_Protocols_Ethernet_Profiles_Profile_Statistics

    // Schedule to use for probes within an operation.
    Schedule Sla_Protocols_Ethernet_Profiles_Profile_Schedule

    // Probe configuration for the SLA profile.
    Probe Sla_Protocols_Ethernet_Profiles_Profile_Probe
}

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "packet-type" { return "PacketType" }
    if yname == "statistics" { return "Statistics" }
    if yname == "schedule" { return "Schedule" }
    if yname == "probe" { return "Probe" }
    return ""
}

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &profile.Statistics
    }
    if childYangName == "schedule" {
        return &profile.Schedule
    }
    if childYangName == "probe" {
        return &profile.Probe
    }
    return nil
}

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &profile.Statistics
    children["schedule"] = &profile.Schedule
    children["probe"] = &profile.Probe
    return children
}

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["packet-type"] = profile.PacketType
    return leafs
}

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics
// Statistics configuration for the SLA profile
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of statistic. The type is slice of
    // Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic.
    Statistic []Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic
}

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetGoName(yname string) string {
    if yname == "statistic" { return "Statistic" }
    return ""
}

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic" {
        for _, c := range statistics.Statistic {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic{}
        statistics.Statistic = append(statistics.Statistic, child)
        return &statistics.Statistic[len(statistics.Statistic)-1]
    }
    return nil
}

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Statistic {
        children[statistics.Statistic[i].GetSegmentPath()] = &statistics.Statistic[i]
    }
    return children
}

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetYangName() string { return "statistics" }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetParentYangName() string { return "profile" }

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic
// Type of statistic
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type of statistic to measure. The type is
    // SlaStatisticTypeEnum.
    StatisticName interface{}

    // Enable statistic gathering of the metric. The type is interface{}.
    Enable interface{}

    // Number of buckets to archive in memory. The type is interface{} with range:
    // 1..100.
    BucketsArchive interface{}

    // Size of the buckets into which statistics are collected.
    BucketsSize Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize

    // Aggregation to apply to results for the statistic.
    Aggregation Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation
}

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetFilter() yfilter.YFilter { return statistic.YFilter }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) SetFilter(yf yfilter.YFilter) { statistic.YFilter = yf }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetGoName(yname string) string {
    if yname == "statistic-name" { return "StatisticName" }
    if yname == "enable" { return "Enable" }
    if yname == "buckets-archive" { return "BucketsArchive" }
    if yname == "buckets-size" { return "BucketsSize" }
    if yname == "aggregation" { return "Aggregation" }
    return ""
}

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetSegmentPath() string {
    return "statistic" + "[statistic-name='" + fmt.Sprintf("%v", statistic.StatisticName) + "']"
}

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "buckets-size" {
        return &statistic.BucketsSize
    }
    if childYangName == "aggregation" {
        return &statistic.Aggregation
    }
    return nil
}

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["buckets-size"] = &statistic.BucketsSize
    children["aggregation"] = &statistic.Aggregation
    return children
}

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["statistic-name"] = statistic.StatisticName
    leafs["enable"] = statistic.Enable
    leafs["buckets-archive"] = statistic.BucketsArchive
    return leafs
}

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetBundleName() string { return "cisco_ios_xr" }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetYangName() string { return "statistic" }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) SetParent(parent types.Entity) { statistic.parent = parent }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetParent() types.Entity { return statistic.parent }

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetParentYangName() string { return "statistics" }

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize
// Size of the buckets into which statistics
// are collected
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size of each bucket. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BucketsSize interface{}

    // Unit associated with the BucketsSize. The type is SlaBucketsSizeUnitsEnum.
    // This attribute is mandatory.
    BucketsSizeUnit interface{}
}

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetFilter() yfilter.YFilter { return bucketsSize.YFilter }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) SetFilter(yf yfilter.YFilter) { bucketsSize.YFilter = yf }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetGoName(yname string) string {
    if yname == "buckets-size" { return "BucketsSize" }
    if yname == "buckets-size-unit" { return "BucketsSizeUnit" }
    return ""
}

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetSegmentPath() string {
    return "buckets-size"
}

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["buckets-size"] = bucketsSize.BucketsSize
    leafs["buckets-size-unit"] = bucketsSize.BucketsSizeUnit
    return leafs
}

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetBundleName() string { return "cisco_ios_xr" }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetYangName() string { return "buckets-size" }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) SetParent(parent types.Entity) { bucketsSize.parent = parent }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetParent() types.Entity { return bucketsSize.parent }

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetParentYangName() string { return "statistic" }

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation
// Aggregation to apply to results for the
// statistic
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of bins to aggregate results into (0 for no aggregation). The type
    // is interface{} with range: 0..100. This attribute is mandatory.
    BinsCount interface{}

    // Width of each bin. The type is interface{} with range: 1..10000.
    BinsWidth interface{}

    // Tenths portion of the bin width. The type is interface{} with range: 0..9.
    BinsWidthTenths interface{}
}

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetFilter() yfilter.YFilter { return aggregation.YFilter }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) SetFilter(yf yfilter.YFilter) { aggregation.YFilter = yf }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetGoName(yname string) string {
    if yname == "bins-count" { return "BinsCount" }
    if yname == "bins-width" { return "BinsWidth" }
    if yname == "bins-width-tenths" { return "BinsWidthTenths" }
    return ""
}

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetSegmentPath() string {
    return "aggregation"
}

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bins-count"] = aggregation.BinsCount
    leafs["bins-width"] = aggregation.BinsWidth
    leafs["bins-width-tenths"] = aggregation.BinsWidthTenths
    return leafs
}

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetBundleName() string { return "cisco_ios_xr" }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetYangName() string { return "aggregation" }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) SetParent(parent types.Entity) { aggregation.parent = parent }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetParent() types.Entity { return aggregation.parent }

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetParentYangName() string { return "statistic" }

// Sla_Protocols_Ethernet_Profiles_Profile_Schedule
// Schedule to use for probes within an
// operation
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Schedule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between probes.  This must be specified if, and only if,
    // ProbeIntervalUnit is not 'Week' or 'Day'. The type is interface{} with
    // range: 1..90.
    ProbeInterval interface{}

    // Day of week on which to schedule probes.  This must be specified if, and
    // only if, ProbeIntervalUnit is 'Week'. The type is SlaProbeIntervalDayEnum.
    ProbeIntervalDay interface{}

    // Time unit associated with the ProbeInterval. The value must not be 'Once'. 
    // If 'Week' or 'Day' is specified, probes are scheduled weekly or daily
    // respectively. The type is SlaProbeIntervalUnitsEnum. This attribute is
    // mandatory.
    ProbeIntervalUnit interface{}

    // Time after midnight (in UTC) to send the first packet each day. The type is
    // interface{} with range: 0..23.
    StartTimeHour interface{}

    // Time after midnight (in UTC) to send the first packet each day. This must
    // be specified if, and only if, StartTimeHour is specified. The type is
    // interface{} with range: 0..59.
    StartTimeMinute interface{}

    // Time after midnight (in UTC) to send the first packet each day. This must
    // only be specified if StartTimeHour is specified, and must not be specified
    // if ProbeIntervalUnit is 'Week' or 'Day'. The type is interface{} with
    // range: 0..59.
    StartTimeSecond interface{}

    // Duration of each probe.  This must be specified if, and only if,
    // ProbeDurationUnit is specified. The type is interface{} with range:
    // 1..3600.
    ProbeDuration interface{}

    // Time unit associated with the ProbeDuration. The value must not be 'Once'.
    // The type is SlaProbeDurationUnitsEnum.
    ProbeDurationUnit interface{}
}

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetFilter() yfilter.YFilter { return schedule.YFilter }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) SetFilter(yf yfilter.YFilter) { schedule.YFilter = yf }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetGoName(yname string) string {
    if yname == "probe-interval" { return "ProbeInterval" }
    if yname == "probe-interval-day" { return "ProbeIntervalDay" }
    if yname == "probe-interval-unit" { return "ProbeIntervalUnit" }
    if yname == "start-time-hour" { return "StartTimeHour" }
    if yname == "start-time-minute" { return "StartTimeMinute" }
    if yname == "start-time-second" { return "StartTimeSecond" }
    if yname == "probe-duration" { return "ProbeDuration" }
    if yname == "probe-duration-unit" { return "ProbeDurationUnit" }
    return ""
}

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetSegmentPath() string {
    return "schedule"
}

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["probe-interval"] = schedule.ProbeInterval
    leafs["probe-interval-day"] = schedule.ProbeIntervalDay
    leafs["probe-interval-unit"] = schedule.ProbeIntervalUnit
    leafs["start-time-hour"] = schedule.StartTimeHour
    leafs["start-time-minute"] = schedule.StartTimeMinute
    leafs["start-time-second"] = schedule.StartTimeSecond
    leafs["probe-duration"] = schedule.ProbeDuration
    leafs["probe-duration-unit"] = schedule.ProbeDurationUnit
    return leafs
}

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetBundleName() string { return "cisco_ios_xr" }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetYangName() string { return "schedule" }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) SetParent(parent types.Entity) { schedule.parent = parent }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetParent() types.Entity { return schedule.parent }

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetParentYangName() string { return "profile" }

// Sla_Protocols_Ethernet_Profiles_Profile_Probe
// Probe configuration for the SLA profile
type Sla_Protocols_Ethernet_Profiles_Profile_Probe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Priority class to assign to outgoing SLA packets. The type is interface{}
    // with range: 0..7.
    Priority interface{}

    // Number of packets to use in each FLR calculation. The type is interface{}
    // with range: 10..12096000.
    SyntheticLossCalculationPackets interface{}

    // Schedule to use for packets within a burst.  The default value is to send a
    // single packet once.
    Send Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send

    // Minimum size to pad outgoing packet to.
    PacketSizeAndPadding Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding
}

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetFilter() yfilter.YFilter { return probe.YFilter }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) SetFilter(yf yfilter.YFilter) { probe.YFilter = yf }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetGoName(yname string) string {
    if yname == "priority" { return "Priority" }
    if yname == "synthetic-loss-calculation-packets" { return "SyntheticLossCalculationPackets" }
    if yname == "send" { return "Send" }
    if yname == "packet-size-and-padding" { return "PacketSizeAndPadding" }
    return ""
}

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetSegmentPath() string {
    return "probe"
}

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "send" {
        return &probe.Send
    }
    if childYangName == "packet-size-and-padding" {
        return &probe.PacketSizeAndPadding
    }
    return nil
}

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["send"] = &probe.Send
    children["packet-size-and-padding"] = &probe.PacketSizeAndPadding
    return children
}

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority"] = probe.Priority
    leafs["synthetic-loss-calculation-packets"] = probe.SyntheticLossCalculationPackets
    return leafs
}

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetBundleName() string { return "cisco_ios_xr" }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetYangName() string { return "probe" }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) SetParent(parent types.Entity) { probe.parent = parent }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetParent() types.Entity { return probe.parent }

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetParentYangName() string { return "profile" }

// Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send
// Schedule to use for packets within a burst. 
// The default value is to send a single packet
// once.
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between bursts.  This must be specified if, and only if, the
    // SendType is 'Burst' and the 'BurstIntervalUnit' is not 'Once'. The type is
    // interface{} with range: 1..3600.
    BurstInterval interface{}

    // Time unit associated with the BurstInterval .  This must be specified if,
    // and only if, SendType is 'Burst'. The type is SlaBurstIntervalUnitsEnum.
    BurstIntervalUnit interface{}

    // Interval between packets.  This must be specified if, and only if,
    // PacketIntervalUnit is not 'Once'. The type is interface{} with range:
    // 1..30000.
    PacketInterval interface{}

    // Time unit associated with the PacketInterval. The type is
    // SlaPacketIntervalUnitsEnum. This attribute is mandatory.
    PacketIntervalUnit interface{}

    // The number of packets in each burst.  This must be specified if, and only
    // if, the SendType is 'Burst'. The type is interface{} with range: 2..1200.
    PacketCount interface{}

    // The packet distribution: single packets or bursts of packets.  If 'Burst'
    // is specified , PacketCount and BurstInterval must be specified. The type is
    // SlaSend. This attribute is mandatory.
    SendType interface{}
}

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetFilter() yfilter.YFilter { return send.YFilter }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) SetFilter(yf yfilter.YFilter) { send.YFilter = yf }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetGoName(yname string) string {
    if yname == "burst-interval" { return "BurstInterval" }
    if yname == "burst-interval-unit" { return "BurstIntervalUnit" }
    if yname == "packet-interval" { return "PacketInterval" }
    if yname == "packet-interval-unit" { return "PacketIntervalUnit" }
    if yname == "packet-count" { return "PacketCount" }
    if yname == "send-type" { return "SendType" }
    return ""
}

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetSegmentPath() string {
    return "send"
}

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["burst-interval"] = send.BurstInterval
    leafs["burst-interval-unit"] = send.BurstIntervalUnit
    leafs["packet-interval"] = send.PacketInterval
    leafs["packet-interval-unit"] = send.PacketIntervalUnit
    leafs["packet-count"] = send.PacketCount
    leafs["send-type"] = send.SendType
    return leafs
}

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetBundleName() string { return "cisco_ios_xr" }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetYangName() string { return "send" }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) SetParent(parent types.Entity) { send.parent = parent }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetParent() types.Entity { return send.parent }

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetParentYangName() string { return "probe" }

// Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding
// Minimum size to pad outgoing packet to
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum size to pad outgoing packet to. The type is interface{} with range:
    // 1..9000. This attribute is mandatory.
    Size interface{}

    // Type of padding to be used for the packet. The type is SlaPaddingPattern.
    PaddingType interface{}

    // Pattern to be used for hex padding. This can be specified if, and only if,
    // the PaddingType is 'Hex'. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    PaddingValue interface{}
}

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetFilter() yfilter.YFilter { return packetSizeAndPadding.YFilter }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) SetFilter(yf yfilter.YFilter) { packetSizeAndPadding.YFilter = yf }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "padding-type" { return "PaddingType" }
    if yname == "padding-value" { return "PaddingValue" }
    return ""
}

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetSegmentPath() string {
    return "packet-size-and-padding"
}

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = packetSizeAndPadding.Size
    leafs["padding-type"] = packetSizeAndPadding.PaddingType
    leafs["padding-value"] = packetSizeAndPadding.PaddingValue
    return leafs
}

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetBundleName() string { return "cisco_ios_xr" }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetYangName() string { return "packet-size-and-padding" }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) SetParent(parent types.Entity) { packetSizeAndPadding.parent = parent }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetParent() types.Entity { return packetSizeAndPadding.parent }

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetParentYangName() string { return "probe" }

