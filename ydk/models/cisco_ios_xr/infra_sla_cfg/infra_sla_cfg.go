// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-sla package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sla: SLA prtocol and profile Configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of all SLA protocols.
    Protocols Sla_Protocols
}

func (sla *Sla) GetEntityData() *types.CommonEntityData {
    sla.EntityData.YFilter = sla.YFilter
    sla.EntityData.YangName = "sla"
    sla.EntityData.BundleName = "cisco_ios_xr"
    sla.EntityData.ParentYangName = "Cisco-IOS-XR-infra-sla-cfg"
    sla.EntityData.SegmentPath = "Cisco-IOS-XR-infra-sla-cfg:sla"
    sla.EntityData.AbsolutePath = sla.EntityData.SegmentPath
    sla.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sla.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sla.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sla.EntityData.Children = types.NewOrderedMap()
    sla.EntityData.Children.Append("protocols", types.YChild{"Protocols", &sla.Protocols})
    sla.EntityData.Leafs = types.NewOrderedMap()

    sla.EntityData.YListKeys = []string {}

    return &(sla.EntityData)
}

// Sla_Protocols
// Table of all SLA protocols
type Sla_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Ethernet SLA protocol.
    Ethernet Sla_Protocols_Ethernet
}

func (protocols *Sla_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "sla"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/" + protocols.EntityData.SegmentPath
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("Cisco-IOS-XR-ethernet-cfm-cfg:ethernet", types.YChild{"Ethernet", &protocols.Ethernet})
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// Sla_Protocols_Ethernet
// The Ethernet SLA protocol
type Sla_Protocols_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of SLA profiles on the protocol.
    Profiles Sla_Protocols_Ethernet_Profiles
}

func (ethernet *Sla_Protocols_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "protocols"
    ethernet.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-cfm-cfg:ethernet"
    ethernet.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/" + ethernet.EntityData.SegmentPath
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Children.Append("profiles", types.YChild{"Profiles", &ethernet.Profiles})
    ethernet.EntityData.Leafs = types.NewOrderedMap()

    ethernet.EntityData.YListKeys = []string {}

    return &(ethernet.EntityData)
}

// Sla_Protocols_Ethernet_Profiles
// Table of SLA profiles on the protocol
type Sla_Protocols_Ethernet_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile. The type is slice of
    // Sla_Protocols_Ethernet_Profiles_Profile.
    Profile []*Sla_Protocols_Ethernet_Profiles_Profile
}

func (profiles *Sla_Protocols_Ethernet_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "ethernet"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/" + profiles.EntityData.SegmentPath
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

// Sla_Protocols_Ethernet_Profiles_Profile
// Name of the profile
type Sla_Protocols_Ethernet_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (profile *Sla_Protocols_Ethernet_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("statistics", types.YChild{"Statistics", &profile.Statistics})
    profile.EntityData.Children.Append("schedule", types.YChild{"Schedule", &profile.Schedule})
    profile.EntityData.Children.Append("probe", types.YChild{"Probe", &profile.Probe})
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})
    profile.EntityData.Leafs.Append("packet-type", types.YLeaf{"PacketType", profile.PacketType})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics
// Statistics configuration for the SLA profile
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of statistic. The type is slice of
    // Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic.
    Statistic []*Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic
}

func (statistics *Sla_Protocols_Ethernet_Profiles_Profile_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "profile"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("statistic", types.YChild{"Statistic", nil})
    for i := range statistics.Statistic {
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Statistic[i]), types.YChild{"Statistic", statistics.Statistic[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic
// Type of statistic
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Thresholds and associated actions for the given statistics type.
    Actions Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions

    // Aggregation to apply to results for the statistic.
    Aggregation Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation
}

func (statistic *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic) GetEntityData() *types.CommonEntityData {
    statistic.EntityData.YFilter = statistic.YFilter
    statistic.EntityData.YangName = "statistic"
    statistic.EntityData.BundleName = "cisco_ios_xr"
    statistic.EntityData.ParentYangName = "statistics"
    statistic.EntityData.SegmentPath = "statistic" + types.AddKeyToken(statistic.StatisticName, "statistic-name")
    statistic.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/statistics/" + statistic.EntityData.SegmentPath
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = types.NewOrderedMap()
    statistic.EntityData.Children.Append("buckets-size", types.YChild{"BucketsSize", &statistic.BucketsSize})
    statistic.EntityData.Children.Append("actions", types.YChild{"Actions", &statistic.Actions})
    statistic.EntityData.Children.Append("aggregation", types.YChild{"Aggregation", &statistic.Aggregation})
    statistic.EntityData.Leafs = types.NewOrderedMap()
    statistic.EntityData.Leafs.Append("statistic-name", types.YLeaf{"StatisticName", statistic.StatisticName})
    statistic.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", statistic.Enable})
    statistic.EntityData.Leafs.Append("buckets-archive", types.YLeaf{"BucketsArchive", statistic.BucketsArchive})

    statistic.EntityData.YListKeys = []string {"StatisticName"}

    return &(statistic.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize
// Size of the buckets into which statistics
// are collected
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Size of each bucket. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BucketsSize interface{}

    // Unit associated with the BucketsSize. The type is SlaBucketsSizeUnitsEnum.
    // This attribute is mandatory.
    BucketsSizeUnit interface{}
}

func (bucketsSize *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_BucketsSize) GetEntityData() *types.CommonEntityData {
    bucketsSize.EntityData.YFilter = bucketsSize.YFilter
    bucketsSize.EntityData.YangName = "buckets-size"
    bucketsSize.EntityData.BundleName = "cisco_ios_xr"
    bucketsSize.EntityData.ParentYangName = "statistic"
    bucketsSize.EntityData.SegmentPath = "buckets-size"
    bucketsSize.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/statistics/statistic/" + bucketsSize.EntityData.SegmentPath
    bucketsSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bucketsSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bucketsSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bucketsSize.EntityData.Children = types.NewOrderedMap()
    bucketsSize.EntityData.Leafs = types.NewOrderedMap()
    bucketsSize.EntityData.Leafs.Append("buckets-size", types.YLeaf{"BucketsSize", bucketsSize.BucketsSize})
    bucketsSize.EntityData.Leafs.Append("buckets-size-unit", types.YLeaf{"BucketsSizeUnit", bucketsSize.BucketsSizeUnit})

    bucketsSize.EntityData.YListKeys = []string {}

    return &(bucketsSize.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions
// Thresholds and associated actions for the
// given statistics type
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action to perform when the threshold is crossed. The type is slice of
    // Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions_Action.
    Action []*Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions_Action
}

func (actions *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions) GetEntityData() *types.CommonEntityData {
    actions.EntityData.YFilter = actions.YFilter
    actions.EntityData.YangName = "actions"
    actions.EntityData.BundleName = "cisco_ios_xr"
    actions.EntityData.ParentYangName = "statistic"
    actions.EntityData.SegmentPath = "actions"
    actions.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/statistics/statistic/" + actions.EntityData.SegmentPath
    actions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actions.EntityData.Children = types.NewOrderedMap()
    actions.EntityData.Children.Append("action", types.YChild{"Action", nil})
    for i := range actions.Action {
        actions.EntityData.Children.Append(types.GetSegmentPath(actions.Action[i]), types.YChild{"Action", actions.Action[i]})
    }
    actions.EntityData.Leafs = types.NewOrderedMap()

    actions.EntityData.YListKeys = []string {}

    return &(actions.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions_Action
// Action to perform when the threshold is
// crossed
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions_Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Threshold type. The type is SlaThresholdTypeEnum.
    ThresholdType interface{}

    // This attribute is a key. Action to take when the threshold is crossed. The
    // type is SlaActionTypeEnum.
    ActionType interface{}

    // This attribute is a key. Condition to be met to consider the threshold
    // crossed. The type is SlaThresholdConditionEnum.
    Condition interface{}

    // Threshold Value. The type is interface{} with range: 1..2147483647. This
    // attribute is mandatory.
    ThresholdValue interface{}

    // Bin number in-and-above which samples contribute towards a sample-count
    // threshold (required only when Condition is SampleCount). The type is
    // interface{} with range: 2..100.
    BinNumber interface{}
}

func (action *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Actions_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "actions"
    action.EntityData.SegmentPath = "action" + types.AddKeyToken(action.ThresholdType, "threshold-type") + types.AddKeyToken(action.ActionType, "action-type") + types.AddKeyToken(action.Condition, "condition")
    action.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/statistics/statistic/actions/" + action.EntityData.SegmentPath
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Leafs = types.NewOrderedMap()
    action.EntityData.Leafs.Append("threshold-type", types.YLeaf{"ThresholdType", action.ThresholdType})
    action.EntityData.Leafs.Append("action-type", types.YLeaf{"ActionType", action.ActionType})
    action.EntityData.Leafs.Append("condition", types.YLeaf{"Condition", action.Condition})
    action.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", action.ThresholdValue})
    action.EntityData.Leafs.Append("bin-number", types.YLeaf{"BinNumber", action.BinNumber})

    action.EntityData.YListKeys = []string {"ThresholdType", "ActionType", "Condition"}

    return &(action.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation
// Aggregation to apply to results for the
// statistic
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Number of bins to aggregate results into (0 for no aggregation). The type
    // is interface{} with range: 0..100. This attribute is mandatory.
    BinsCount interface{}

    // Width of each bin. The type is interface{} with range: 1..10000.
    BinsWidth interface{}

    // Tenths portion of the bin width. The type is interface{} with range: 0..9.
    BinsWidthTenths interface{}
}

func (aggregation *Sla_Protocols_Ethernet_Profiles_Profile_Statistics_Statistic_Aggregation) GetEntityData() *types.CommonEntityData {
    aggregation.EntityData.YFilter = aggregation.YFilter
    aggregation.EntityData.YangName = "aggregation"
    aggregation.EntityData.BundleName = "cisco_ios_xr"
    aggregation.EntityData.ParentYangName = "statistic"
    aggregation.EntityData.SegmentPath = "aggregation"
    aggregation.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/statistics/statistic/" + aggregation.EntityData.SegmentPath
    aggregation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregation.EntityData.Children = types.NewOrderedMap()
    aggregation.EntityData.Leafs = types.NewOrderedMap()
    aggregation.EntityData.Leafs.Append("bins-count", types.YLeaf{"BinsCount", aggregation.BinsCount})
    aggregation.EntityData.Leafs.Append("bins-width", types.YLeaf{"BinsWidth", aggregation.BinsWidth})
    aggregation.EntityData.Leafs.Append("bins-width-tenths", types.YLeaf{"BinsWidthTenths", aggregation.BinsWidthTenths})

    aggregation.EntityData.YListKeys = []string {}

    return &(aggregation.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Schedule
// Schedule to use for probes within an
// operation
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Schedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Interval between probes.  This must be specified if, and only if,
    // ProbeIntervalUnit is not 'Week' or 'Day'. The type is interface{} with
    // range: 1..90.
    ProbeInterval interface{}

    // Day of week on which to schedule probes. This must be specified if, and
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

func (schedule *Sla_Protocols_Ethernet_Profiles_Profile_Schedule) GetEntityData() *types.CommonEntityData {
    schedule.EntityData.YFilter = schedule.YFilter
    schedule.EntityData.YangName = "schedule"
    schedule.EntityData.BundleName = "cisco_ios_xr"
    schedule.EntityData.ParentYangName = "profile"
    schedule.EntityData.SegmentPath = "schedule"
    schedule.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/" + schedule.EntityData.SegmentPath
    schedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schedule.EntityData.Children = types.NewOrderedMap()
    schedule.EntityData.Leafs = types.NewOrderedMap()
    schedule.EntityData.Leafs.Append("probe-interval", types.YLeaf{"ProbeInterval", schedule.ProbeInterval})
    schedule.EntityData.Leafs.Append("probe-interval-day", types.YLeaf{"ProbeIntervalDay", schedule.ProbeIntervalDay})
    schedule.EntityData.Leafs.Append("probe-interval-unit", types.YLeaf{"ProbeIntervalUnit", schedule.ProbeIntervalUnit})
    schedule.EntityData.Leafs.Append("start-time-hour", types.YLeaf{"StartTimeHour", schedule.StartTimeHour})
    schedule.EntityData.Leafs.Append("start-time-minute", types.YLeaf{"StartTimeMinute", schedule.StartTimeMinute})
    schedule.EntityData.Leafs.Append("start-time-second", types.YLeaf{"StartTimeSecond", schedule.StartTimeSecond})
    schedule.EntityData.Leafs.Append("probe-duration", types.YLeaf{"ProbeDuration", schedule.ProbeDuration})
    schedule.EntityData.Leafs.Append("probe-duration-unit", types.YLeaf{"ProbeDurationUnit", schedule.ProbeDurationUnit})

    schedule.EntityData.YListKeys = []string {}

    return &(schedule.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Probe
// Probe configuration for the SLA profile
type Sla_Protocols_Ethernet_Profiles_Profile_Probe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority class to assign to outgoing SLA packets. The type is interface{}
    // with range: 0..7.
    Priority interface{}

    // Number of packets to use in each FLR calculation. The type is interface{}
    // with range: 10..12096000.
    SyntheticLossCalculationPackets interface{}

    // Schedule to use for packets within a burst. The default value is to send a
    // single packet once.
    Send Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send

    // Minimum size to pad outgoing packet to.
    PacketSizeAndPadding Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding
}

func (probe *Sla_Protocols_Ethernet_Profiles_Profile_Probe) GetEntityData() *types.CommonEntityData {
    probe.EntityData.YFilter = probe.YFilter
    probe.EntityData.YangName = "probe"
    probe.EntityData.BundleName = "cisco_ios_xr"
    probe.EntityData.ParentYangName = "profile"
    probe.EntityData.SegmentPath = "probe"
    probe.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/" + probe.EntityData.SegmentPath
    probe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probe.EntityData.Children = types.NewOrderedMap()
    probe.EntityData.Children.Append("send", types.YChild{"Send", &probe.Send})
    probe.EntityData.Children.Append("packet-size-and-padding", types.YChild{"PacketSizeAndPadding", &probe.PacketSizeAndPadding})
    probe.EntityData.Leafs = types.NewOrderedMap()
    probe.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", probe.Priority})
    probe.EntityData.Leafs.Append("synthetic-loss-calculation-packets", types.YLeaf{"SyntheticLossCalculationPackets", probe.SyntheticLossCalculationPackets})

    probe.EntityData.YListKeys = []string {}

    return &(probe.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send
// Schedule to use for packets within a burst.
// The default value is to send a single packet
// once.
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (send *Sla_Protocols_Ethernet_Profiles_Profile_Probe_Send) GetEntityData() *types.CommonEntityData {
    send.EntityData.YFilter = send.YFilter
    send.EntityData.YangName = "send"
    send.EntityData.BundleName = "cisco_ios_xr"
    send.EntityData.ParentYangName = "probe"
    send.EntityData.SegmentPath = "send"
    send.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/probe/" + send.EntityData.SegmentPath
    send.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    send.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    send.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    send.EntityData.Children = types.NewOrderedMap()
    send.EntityData.Leafs = types.NewOrderedMap()
    send.EntityData.Leafs.Append("burst-interval", types.YLeaf{"BurstInterval", send.BurstInterval})
    send.EntityData.Leafs.Append("burst-interval-unit", types.YLeaf{"BurstIntervalUnit", send.BurstIntervalUnit})
    send.EntityData.Leafs.Append("packet-interval", types.YLeaf{"PacketInterval", send.PacketInterval})
    send.EntityData.Leafs.Append("packet-interval-unit", types.YLeaf{"PacketIntervalUnit", send.PacketIntervalUnit})
    send.EntityData.Leafs.Append("packet-count", types.YLeaf{"PacketCount", send.PacketCount})
    send.EntityData.Leafs.Append("send-type", types.YLeaf{"SendType", send.SendType})

    send.EntityData.YListKeys = []string {}

    return &(send.EntityData)
}

// Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding
// Minimum size to pad outgoing packet to
// This type is a presence type.
type Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (packetSizeAndPadding *Sla_Protocols_Ethernet_Profiles_Profile_Probe_PacketSizeAndPadding) GetEntityData() *types.CommonEntityData {
    packetSizeAndPadding.EntityData.YFilter = packetSizeAndPadding.YFilter
    packetSizeAndPadding.EntityData.YangName = "packet-size-and-padding"
    packetSizeAndPadding.EntityData.BundleName = "cisco_ios_xr"
    packetSizeAndPadding.EntityData.ParentYangName = "probe"
    packetSizeAndPadding.EntityData.SegmentPath = "packet-size-and-padding"
    packetSizeAndPadding.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-sla-cfg:sla/protocols/Cisco-IOS-XR-ethernet-cfm-cfg:ethernet/profiles/profile/probe/" + packetSizeAndPadding.EntityData.SegmentPath
    packetSizeAndPadding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetSizeAndPadding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetSizeAndPadding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetSizeAndPadding.EntityData.Children = types.NewOrderedMap()
    packetSizeAndPadding.EntityData.Leafs = types.NewOrderedMap()
    packetSizeAndPadding.EntityData.Leafs.Append("size", types.YLeaf{"Size", packetSizeAndPadding.Size})
    packetSizeAndPadding.EntityData.Leafs.Append("padding-type", types.YLeaf{"PaddingType", packetSizeAndPadding.PaddingType})
    packetSizeAndPadding.EntityData.Leafs.Append("padding-value", types.YLeaf{"PaddingValue", packetSizeAndPadding.PaddingValue})

    packetSizeAndPadding.EntityData.YListKeys = []string {}

    return &(packetSizeAndPadding.EntityData)
}

