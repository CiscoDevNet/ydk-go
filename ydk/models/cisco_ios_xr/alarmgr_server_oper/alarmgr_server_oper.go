// This module contains a collection of YANG definitions
// for Cisco IOS-XR alarmgr-server package operational data.
// 
// This module contains definitions
// for the following management objects:
//   alarms: Show Alarms associated with XR
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package alarmgr_server_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package alarmgr_server_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-alarmgr-server-oper alarms}", reflect.TypeOf(Alarms{}))
    ydk.RegisterEntity("Cisco-IOS-XR-alarmgr-server-oper:alarms", reflect.TypeOf(Alarms{}))
}

// TimingBucket represents Timing bucket
type TimingBucket string

const (
    // Bucket Type not applicable
    TimingBucket_not_specified TimingBucket = "not-specified"

    // Fifteen minute time bucket
    TimingBucket_fifteen_min TimingBucket = "fifteen-min"

    // One day time bucket
    TimingBucket_one_day TimingBucket = "one-day"
)

// AlarmSeverity represents Alarm severity
type AlarmSeverity string

const (
    // Unknown severity level
    AlarmSeverity_unknown AlarmSeverity = "unknown"

    // Severity level not reported will not raise an
    // alarm
    AlarmSeverity_not_reported AlarmSeverity = "not-reported"

    // Severity level of info to cater to events such
    // as Performance TCAS
    AlarmSeverity_not_alarmed AlarmSeverity = "not-alarmed"

    // Severity level of minor fault not traffic
    // affecting
    AlarmSeverity_minor AlarmSeverity = "minor"

    // Severity level of major fault leading to
    // service disruption
    AlarmSeverity_major AlarmSeverity = "major"

    // Severity level of critical leading to drops
    // ,route loss, loss of service etc.
    AlarmSeverity_critical AlarmSeverity = "critical"

    // Last severity level
    AlarmSeverity_severity_last AlarmSeverity = "severity-last"
)

// AlarmDirection represents Alarm direction
type AlarmDirection string

const (
    // Direction Not Specified
    AlarmDirection_not_specified AlarmDirection = "not-specified"

    // Direction Send
    AlarmDirection_send AlarmDirection = "send"

    // Direction Receive
    AlarmDirection_receive AlarmDirection = "receive"

    // Direction Send and Receive
    AlarmDirection_send_receive AlarmDirection = "send-receive"
)

// AlarmStatus represents Alarm status
type AlarmStatus string

const (
    // Unknown alarm status level
    AlarmStatus_unknown AlarmStatus = "unknown"

    // Status of active alarm that is SET by the
    // controller
    AlarmStatus_set AlarmStatus = "set"

    // Status of cleared alarm that is done by the
    // controller
    AlarmStatus_clear AlarmStatus = "clear"

    // Status of suppressed alarm that is done by the
    // controller
    AlarmStatus_suppress AlarmStatus = "suppress"

    // Last status level
    AlarmStatus_last AlarmStatus = "last"
)

// AlarmServiceAffecting represents Alarm service affecting
type AlarmServiceAffecting string

const (
    // Unknown whether alarm severity is service
    // affecting
    AlarmServiceAffecting_unknown AlarmServiceAffecting = "unknown"

    // Alarm severity is not service affecting
    AlarmServiceAffecting_not_service_affecting AlarmServiceAffecting = "not-service-affecting"

    // Alarm severity is service affecting
    AlarmServiceAffecting_service_affecting AlarmServiceAffecting = "service-affecting"
)

// AlarmNotificationSrc represents Alarm notification src
type AlarmNotificationSrc string

const (
    // Notification src not specified
    AlarmNotificationSrc_not_specified AlarmNotificationSrc = "not-specified"

    // Notification src near end
    AlarmNotificationSrc_near_end AlarmNotificationSrc = "near-end"

    // Notification src far end
    AlarmNotificationSrc_far_end AlarmNotificationSrc = "far-end"
)

// AlarmEvent represents Alarm event
type AlarmEvent string

const (
    // Default Alarm Event Type
    AlarmEvent_default_ AlarmEvent = "default"

    // Alarm Notifcation Event Type
    AlarmEvent_notification AlarmEvent = "notification"

    // Alarm Type Condition
    AlarmEvent_condition AlarmEvent = "condition"

    // Last Event Type
    AlarmEvent_last AlarmEvent = "last"
)

// AlarmClient represents Alarm client
type AlarmClient string

const (
    // Client type unknown
    AlarmClient_unknown AlarmClient = "unknown"

    // Client type producer
    AlarmClient_producer AlarmClient = "producer"

    // Client type consumer
    AlarmClient_consumer AlarmClient = "consumer"

    // Client type subscriber
    AlarmClient_subscriber AlarmClient = "subscriber"

    // Client type last
    AlarmClient_client_last AlarmClient = "client-last"
)

// AlarmClientState represents Alarm client state
type AlarmClientState string

const (
    // Starting state. Should be 0
    AlarmClientState_start AlarmClientState = "start"

    // Client initalized
    AlarmClientState_init AlarmClientState = "init"

    // Sent connect request
    AlarmClientState_connecting AlarmClientState = "connecting"

    // Initial connected
    AlarmClientState_connected AlarmClientState = "connected"

    // Has sent registration message
    AlarmClientState_registered AlarmClientState = "registered"

    // Has been disconnected due to request of error
    AlarmClientState_disconnected AlarmClientState = "disconnected"

    // The client is ready
    AlarmClientState_ready AlarmClientState = "ready"
)

// AlarmGroups represents Alarm groups
type AlarmGroups string

const (
    // An unknown alarm group
    AlarmGroups_unknown AlarmGroups = "unknown"

    // Environomental alarm group
    AlarmGroups_environ AlarmGroups = "environ"

    // Ethernet alarm group
    AlarmGroups_ethernet AlarmGroups = "ethernet"

    // Fabric related alarm group
    AlarmGroups_fabric AlarmGroups = "fabric"

    // Power and PEM group of alarms
    AlarmGroups_power AlarmGroups = "power"

    // Software group of alarms
    AlarmGroups_software AlarmGroups = "software"

    // Slice group of alarms
    AlarmGroups_slice AlarmGroups = "slice"

    // CPU group of alarms
    AlarmGroups_cpu AlarmGroups = "cpu"

    // Controller group of alarms
    AlarmGroups_controller AlarmGroups = "controller"

    // Sonet group of alarms
    AlarmGroups_sonet AlarmGroups = "sonet"

    // OTN group of alarms
    AlarmGroups_otn AlarmGroups = "otn"

    // SDH group of alarms
    AlarmGroups_sdh_controller AlarmGroups = "sdh-controller"

    // ASIC group of alarms
    AlarmGroups_asic AlarmGroups = "asic"

    // FPD group of alarms
    AlarmGroups_fpd_infra AlarmGroups = "fpd-infra"

    // Shelf group of alarms
    AlarmGroups_shelf AlarmGroups = "shelf"

    // MPA group of alarms
    AlarmGroups_mpa AlarmGroups = "mpa"

    // OTS group of alarms
    AlarmGroups_ots AlarmGroups = "ots"

    // Timing group of alarms
    AlarmGroups_timing AlarmGroups = "timing"

    // Last unused group
    AlarmGroups_last AlarmGroups = "last"
)

// Alarms
// Show Alarms associated with XR
type Alarms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of detail alarm commands.
    Detail Alarms_Detail

    // A set of brief alarm commands.
    Brief Alarms_Brief
}

func (alarms *Alarms) GetEntityData() *types.CommonEntityData {
    alarms.EntityData.YFilter = alarms.YFilter
    alarms.EntityData.YangName = "alarms"
    alarms.EntityData.BundleName = "cisco_ios_xr"
    alarms.EntityData.ParentYangName = "Cisco-IOS-XR-alarmgr-server-oper"
    alarms.EntityData.SegmentPath = "Cisco-IOS-XR-alarmgr-server-oper:alarms"
    alarms.EntityData.AbsolutePath = alarms.EntityData.SegmentPath
    alarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarms.EntityData.Children = types.NewOrderedMap()
    alarms.EntityData.Children.Append("detail", types.YChild{"Detail", &alarms.Detail})
    alarms.EntityData.Children.Append("brief", types.YChild{"Brief", &alarms.Brief})
    alarms.EntityData.Leafs = types.NewOrderedMap()

    alarms.EntityData.YListKeys = []string {}

    return &(alarms.EntityData)
}

// Alarms_Detail
// A set of detail alarm commands.
type Alarms_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show detail system scope alarm related data.
    DetailSystem Alarms_Detail_DetailSystem

    // Show detail card scope alarm related data.
    DetailCard Alarms_Detail_DetailCard
}

func (detail *Alarms_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "alarms"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("detail-system", types.YChild{"DetailSystem", &detail.DetailSystem})
    detail.EntityData.Children.Append("detail-card", types.YChild{"DetailCard", &detail.DetailCard})
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Alarms_Detail_DetailSystem
// show detail system scope alarm related data.
type Alarms_Detail_DetailSystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show the Conditions present at this scope.
    Conditions Alarms_Detail_DetailSystem_Conditions

    // Show the active alarms at this scope.
    Active Alarms_Detail_DetailSystem_Active

    // Show the history alarms at this scope.
    History Alarms_Detail_DetailSystem_History

    // Show the suppressed alarms at this scope.
    Suppressed Alarms_Detail_DetailSystem_Suppressed

    // Show the service statistics.
    Stats Alarms_Detail_DetailSystem_Stats

    // Show the clients associated with this service.
    Clients Alarms_Detail_DetailSystem_Clients
}

func (detailSystem *Alarms_Detail_DetailSystem) GetEntityData() *types.CommonEntityData {
    detailSystem.EntityData.YFilter = detailSystem.YFilter
    detailSystem.EntityData.YangName = "detail-system"
    detailSystem.EntityData.BundleName = "cisco_ios_xr"
    detailSystem.EntityData.ParentYangName = "detail"
    detailSystem.EntityData.SegmentPath = "detail-system"
    detailSystem.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/" + detailSystem.EntityData.SegmentPath
    detailSystem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailSystem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailSystem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailSystem.EntityData.Children = types.NewOrderedMap()
    detailSystem.EntityData.Children.Append("conditions", types.YChild{"Conditions", &detailSystem.Conditions})
    detailSystem.EntityData.Children.Append("active", types.YChild{"Active", &detailSystem.Active})
    detailSystem.EntityData.Children.Append("history", types.YChild{"History", &detailSystem.History})
    detailSystem.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", &detailSystem.Suppressed})
    detailSystem.EntityData.Children.Append("stats", types.YChild{"Stats", &detailSystem.Stats})
    detailSystem.EntityData.Children.Append("clients", types.YChild{"Clients", &detailSystem.Clients})
    detailSystem.EntityData.Leafs = types.NewOrderedMap()

    detailSystem.EntityData.YListKeys = []string {}

    return &(detailSystem.EntityData)
}

// Alarms_Detail_DetailSystem_Conditions
// Show the Conditions present at this scope.
type Alarms_Detail_DetailSystem_Conditions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_Conditions_AlarmInfo.
    AlarmInfo []*Alarms_Detail_DetailSystem_Conditions_AlarmInfo
}

func (conditions *Alarms_Detail_DetailSystem_Conditions) GetEntityData() *types.CommonEntityData {
    conditions.EntityData.YFilter = conditions.YFilter
    conditions.EntityData.YangName = "conditions"
    conditions.EntityData.BundleName = "cisco_ios_xr"
    conditions.EntityData.ParentYangName = "detail-system"
    conditions.EntityData.SegmentPath = "conditions"
    conditions.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/" + conditions.EntityData.SegmentPath
    conditions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conditions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conditions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conditions.EntityData.Children = types.NewOrderedMap()
    conditions.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range conditions.AlarmInfo {
        types.SetYListKey(conditions.AlarmInfo[i], i)
        conditions.EntityData.Children.Append(types.GetSegmentPath(conditions.AlarmInfo[i]), types.YChild{"AlarmInfo", conditions.AlarmInfo[i]})
    }
    conditions.EntityData.Leafs = types.NewOrderedMap()

    conditions.EntityData.YListKeys = []string {}

    return &(conditions.EntityData)
}

// Alarms_Detail_DetailSystem_Conditions_AlarmInfo
// Alarm List
type Alarms_Detail_DetailSystem_Conditions_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm service affecting. The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // alarm event type. The type is AlarmEvent.
    Type interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Otn

    // TCA feature specific alarm attributes.
    Tca Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Tca
}

func (alarmInfo *Alarms_Detail_DetailSystem_Conditions_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "conditions"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/conditions/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &alarmInfo.Otn})
    alarmInfo.EntityData.Children.Append("tca", types.YChild{"Tca", &alarmInfo.Tca})
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", alarmInfo.Aid})
    alarmInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", alarmInfo.Tag})
    alarmInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", alarmInfo.Module})
    alarmInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", alarmInfo.Eid})
    alarmInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId})
    alarmInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", alarmInfo.PendingSync})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", alarmInfo.Status})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting})
    alarmInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", alarmInfo.Type})
    alarmInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", alarmInfo.Interface})
    alarmInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", alarmInfo.AlarmName})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "alarm-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/conditions/alarm-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Tca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailSystem_Conditions_AlarmInfo_Tca) GetEntityData() *types.CommonEntityData {
    tca.EntityData.YFilter = tca.YFilter
    tca.EntityData.YangName = "tca"
    tca.EntityData.BundleName = "cisco_ios_xr"
    tca.EntityData.ParentYangName = "alarm-info"
    tca.EntityData.SegmentPath = "tca"
    tca.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/conditions/alarm-info/" + tca.EntityData.SegmentPath
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = types.NewOrderedMap()
    tca.EntityData.Leafs = types.NewOrderedMap()
    tca.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", tca.ThresholdValue})
    tca.EntityData.Leafs.Append("current-value", types.YLeaf{"CurrentValue", tca.CurrentValue})
    tca.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", tca.BucketType})

    tca.EntityData.YListKeys = []string {}

    return &(tca.EntityData)
}

// Alarms_Detail_DetailSystem_Active
// Show the active alarms at this scope.
type Alarms_Detail_DetailSystem_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_Active_AlarmInfo.
    AlarmInfo []*Alarms_Detail_DetailSystem_Active_AlarmInfo
}

func (active *Alarms_Detail_DetailSystem_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "detail-system"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range active.AlarmInfo {
        types.SetYListKey(active.AlarmInfo[i], i)
        active.EntityData.Children.Append(types.GetSegmentPath(active.AlarmInfo[i]), types.YChild{"AlarmInfo", active.AlarmInfo[i]})
    }
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Alarms_Detail_DetailSystem_Active_AlarmInfo
// Alarm List
type Alarms_Detail_DetailSystem_Active_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm service affecting. The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // alarm event type. The type is AlarmEvent.
    Type interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn

    // TCA feature specific alarm attributes.
    Tca Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca
}

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "active"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/active/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &alarmInfo.Otn})
    alarmInfo.EntityData.Children.Append("tca", types.YChild{"Tca", &alarmInfo.Tca})
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", alarmInfo.Aid})
    alarmInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", alarmInfo.Tag})
    alarmInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", alarmInfo.Module})
    alarmInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", alarmInfo.Eid})
    alarmInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId})
    alarmInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", alarmInfo.PendingSync})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", alarmInfo.Status})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting})
    alarmInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", alarmInfo.Type})
    alarmInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", alarmInfo.Interface})
    alarmInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", alarmInfo.AlarmName})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "alarm-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/active/alarm-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetEntityData() *types.CommonEntityData {
    tca.EntityData.YFilter = tca.YFilter
    tca.EntityData.YangName = "tca"
    tca.EntityData.BundleName = "cisco_ios_xr"
    tca.EntityData.ParentYangName = "alarm-info"
    tca.EntityData.SegmentPath = "tca"
    tca.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/active/alarm-info/" + tca.EntityData.SegmentPath
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = types.NewOrderedMap()
    tca.EntityData.Leafs = types.NewOrderedMap()
    tca.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", tca.ThresholdValue})
    tca.EntityData.Leafs.Append("current-value", types.YLeaf{"CurrentValue", tca.CurrentValue})
    tca.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", tca.BucketType})

    tca.EntityData.YListKeys = []string {}

    return &(tca.EntityData)
}

// Alarms_Detail_DetailSystem_History
// Show the history alarms at this scope.
type Alarms_Detail_DetailSystem_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_History_AlarmInfo.
    AlarmInfo []*Alarms_Detail_DetailSystem_History_AlarmInfo
}

func (history *Alarms_Detail_DetailSystem_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "detail-system"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range history.AlarmInfo {
        types.SetYListKey(history.AlarmInfo[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.AlarmInfo[i]), types.YChild{"AlarmInfo", history.AlarmInfo[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Alarms_Detail_DetailSystem_History_AlarmInfo
// Alarm List
type Alarms_Detail_DetailSystem_History_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm service affecting. The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // alarm event type. The type is AlarmEvent.
    Type interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailSystem_History_AlarmInfo_Otn

    // TCA feature specific alarm attributes.
    Tca Alarms_Detail_DetailSystem_History_AlarmInfo_Tca
}

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "history"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/history/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &alarmInfo.Otn})
    alarmInfo.EntityData.Children.Append("tca", types.YChild{"Tca", &alarmInfo.Tca})
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", alarmInfo.Aid})
    alarmInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", alarmInfo.Tag})
    alarmInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", alarmInfo.Module})
    alarmInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", alarmInfo.Eid})
    alarmInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId})
    alarmInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", alarmInfo.PendingSync})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", alarmInfo.Status})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting})
    alarmInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", alarmInfo.Type})
    alarmInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", alarmInfo.Interface})
    alarmInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", alarmInfo.AlarmName})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Detail_DetailSystem_History_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailSystem_History_AlarmInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "alarm-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/history/alarm-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailSystem_History_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailSystem_History_AlarmInfo_Tca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetEntityData() *types.CommonEntityData {
    tca.EntityData.YFilter = tca.YFilter
    tca.EntityData.YangName = "tca"
    tca.EntityData.BundleName = "cisco_ios_xr"
    tca.EntityData.ParentYangName = "alarm-info"
    tca.EntityData.SegmentPath = "tca"
    tca.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/history/alarm-info/" + tca.EntityData.SegmentPath
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = types.NewOrderedMap()
    tca.EntityData.Leafs = types.NewOrderedMap()
    tca.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", tca.ThresholdValue})
    tca.EntityData.Leafs.Append("current-value", types.YLeaf{"CurrentValue", tca.CurrentValue})
    tca.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", tca.BucketType})

    tca.EntityData.YListKeys = []string {}

    return &(tca.EntityData)
}

// Alarms_Detail_DetailSystem_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Detail_DetailSystem_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo.
    SuppressedInfo []*Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "detail-system"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Children.Append("suppressed-info", types.YChild{"SuppressedInfo", nil})
    for i := range suppressed.SuppressedInfo {
        types.SetYListKey(suppressed.SuppressedInfo[i], i)
        suppressed.EntityData.Children.Append(types.GetSegmentPath(suppressed.SuppressedInfo[i]), types.YChild{"SuppressedInfo", suppressed.SuppressedInfo[i]})
    }
    suppressed.EntityData.Leafs = types.NewOrderedMap()

    suppressed.EntityData.YListKeys = []string {}

    return &(suppressed.EntityData)
}

// Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm suppressed time. The type is string with length: 0..64.
    SuppressedTime interface{}

    // Alarm suppressed time(timestamp format). The type is interface{} with
    // range: 0..18446744073709551615.
    SuppressedTimestamp interface{}

    // Alarm service affecting . The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn
}

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetEntityData() *types.CommonEntityData {
    suppressedInfo.EntityData.YFilter = suppressedInfo.YFilter
    suppressedInfo.EntityData.YangName = "suppressed-info"
    suppressedInfo.EntityData.BundleName = "cisco_ios_xr"
    suppressedInfo.EntityData.ParentYangName = "suppressed"
    suppressedInfo.EntityData.SegmentPath = "suppressed-info" + types.AddNoKeyToken(suppressedInfo)
    suppressedInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/suppressed/" + suppressedInfo.EntityData.SegmentPath
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = types.NewOrderedMap()
    suppressedInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &suppressedInfo.Otn})
    suppressedInfo.EntityData.Leafs = types.NewOrderedMap()
    suppressedInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressedInfo.Description})
    suppressedInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", suppressedInfo.Location})
    suppressedInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressedInfo.Aid})
    suppressedInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressedInfo.Tag})
    suppressedInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressedInfo.Module})
    suppressedInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressedInfo.Eid})
    suppressedInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", suppressedInfo.ReportingAgentId})
    suppressedInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", suppressedInfo.PendingSync})
    suppressedInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressedInfo.Severity})
    suppressedInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", suppressedInfo.Status})
    suppressedInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressedInfo.Group})
    suppressedInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", suppressedInfo.SetTime})
    suppressedInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp})
    suppressedInfo.EntityData.Leafs.Append("suppressed-time", types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime})
    suppressedInfo.EntityData.Leafs.Append("suppressed-timestamp", types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp})
    suppressedInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", suppressedInfo.ServiceAffecting})
    suppressedInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", suppressedInfo.Interface})
    suppressedInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", suppressedInfo.AlarmName})

    suppressedInfo.EntityData.YListKeys = []string {}

    return &(suppressedInfo.EntityData)
}

// Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "suppressed-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/suppressed/suppressed-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailSystem_Stats
// Show the service statistics.
type Alarms_Detail_DetailSystem_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarms that were in all reported to this Alarm Mgr. The type is interface{}
    // with range: 0..18446744073709551615.
    Reported interface{}

    // Alarms that we couldn't keep track due to some error or other. The type is
    // interface{} with range: 0..18446744073709551615.
    Dropped interface{}

    // Alarms that are currently in the active state. The type is interface{} with
    // range: 0..18446744073709551615.
    Active interface{}

    // Alarms that are cleared. This one is counted over a long period of time.
    // The type is interface{} with range: 0..18446744073709551615.
    History interface{}

    // Alarms that are in suppressed state. The type is interface{} with range:
    // 0..18446744073709551615.
    Suppressed interface{}

    // Alarms that are currently in the active state(sysadmin plane). The type is
    // interface{} with range: 0..18446744073709551615.
    SysadminActive interface{}

    // Alarms that are cleared in sysadmin plane. This one is counted over a long
    // period of time. The type is interface{} with range:
    // 0..18446744073709551615.
    SysadminHistory interface{}

    // Alarms that are suppressed in sysadmin plane. The type is interface{} with
    // range: 0..18446744073709551615.
    SysadminSuppressed interface{}

    // Alarms dropped due to invalid aid. The type is interface{} with range:
    // 0..4294967295.
    DroppedInvalidAid interface{}

    // Alarms dropped due to insufficient memory. The type is interface{} with
    // range: 0..4294967295.
    DroppedInsuffMem interface{}

    // Alarms dropped due to db error. The type is interface{} with range:
    // 0..4294967295.
    DroppedDbError interface{}

    // Alarms dropped clear without set. The type is interface{} with range:
    // 0..4294967295.
    DroppedClearWithoutSet interface{}

    // Alarms dropped which were duplicate. The type is interface{} with range:
    // 0..4294967295.
    DroppedDuplicate interface{}

    // Total alarms which had the cache hit. The type is interface{} with range:
    // 0..4294967295.
    CacheHit interface{}

    // Total alarms which had the cache miss. The type is interface{} with range:
    // 0..4294967295.
    CacheMiss interface{}
}

func (stats *Alarms_Detail_DetailSystem_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "detail-system"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("reported", types.YLeaf{"Reported", stats.Reported})
    stats.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", stats.Dropped})
    stats.EntityData.Leafs.Append("active", types.YLeaf{"Active", stats.Active})
    stats.EntityData.Leafs.Append("history", types.YLeaf{"History", stats.History})
    stats.EntityData.Leafs.Append("suppressed", types.YLeaf{"Suppressed", stats.Suppressed})
    stats.EntityData.Leafs.Append("sysadmin-active", types.YLeaf{"SysadminActive", stats.SysadminActive})
    stats.EntityData.Leafs.Append("sysadmin-history", types.YLeaf{"SysadminHistory", stats.SysadminHistory})
    stats.EntityData.Leafs.Append("sysadmin-suppressed", types.YLeaf{"SysadminSuppressed", stats.SysadminSuppressed})
    stats.EntityData.Leafs.Append("dropped-invalid-aid", types.YLeaf{"DroppedInvalidAid", stats.DroppedInvalidAid})
    stats.EntityData.Leafs.Append("dropped-insuff-mem", types.YLeaf{"DroppedInsuffMem", stats.DroppedInsuffMem})
    stats.EntityData.Leafs.Append("dropped-db-error", types.YLeaf{"DroppedDbError", stats.DroppedDbError})
    stats.EntityData.Leafs.Append("dropped-clear-without-set", types.YLeaf{"DroppedClearWithoutSet", stats.DroppedClearWithoutSet})
    stats.EntityData.Leafs.Append("dropped-duplicate", types.YLeaf{"DroppedDuplicate", stats.DroppedDuplicate})
    stats.EntityData.Leafs.Append("cache-hit", types.YLeaf{"CacheHit", stats.CacheHit})
    stats.EntityData.Leafs.Append("cache-miss", types.YLeaf{"CacheMiss", stats.CacheMiss})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// Alarms_Detail_DetailSystem_Clients
// Show the clients associated with this service.
type Alarms_Detail_DetailSystem_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client List. The type is slice of
    // Alarms_Detail_DetailSystem_Clients_ClientInfo.
    ClientInfo []*Alarms_Detail_DetailSystem_Clients_ClientInfo
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "detail-system"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client-info", types.YChild{"ClientInfo", nil})
    for i := range clients.ClientInfo {
        types.SetYListKey(clients.ClientInfo[i], i)
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.ClientInfo[i]), types.YChild{"ClientInfo", clients.ClientInfo[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Alarms_Detail_DetailSystem_Clients_ClientInfo
// Client List
type Alarms_Detail_DetailSystem_Clients_ClientInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm client. The type is string with length: 0..128.
    Name interface{}

    // Alarms agent id of the client. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // The location of this client. The type is string with length: 0..128.
    Location interface{}

    // The client handle through which interface. The type is string with length:
    // 0..128.
    Handle interface{}

    // The current state of the client. The type is AlarmClientState.
    State interface{}

    // The type of the client. The type is AlarmClient.
    Type interface{}

    // The current subscription status of the client. The type is bool.
    FilterDisp interface{}

    // Alarms agent subscriber id of the client. The type is interface{} with
    // range: 0..4294967295.
    SubscriberId interface{}

    // The filter used for alarm severity. The type is AlarmSeverity.
    FilterSeverity interface{}

    // The filter used for alarm bi-state state+. The type is AlarmStatus.
    FilterState interface{}

    // The filter used for alarm group. The type is AlarmGroups.
    FilterGroup interface{}

    // Number of times the agent connected to the alarm mgr. The type is
    // interface{} with range: 0..4294967295.
    ConnectCount interface{}

    // Agent connect timestamp. The type is string with length: 0..64.
    ConnectTimestamp interface{}

    // Number of times the agent queried for alarms. The type is interface{} with
    // range: 0..4294967295.
    GetCount interface{}

    // Number of times the agent subscribed for alarms. The type is interface{}
    // with range: 0..4294967295.
    SubscribeCount interface{}

    // Number of times the agent reported alarms. The type is interface{} with
    // range: 0..4294967295.
    ReportCount interface{}
}

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetEntityData() *types.CommonEntityData {
    clientInfo.EntityData.YFilter = clientInfo.YFilter
    clientInfo.EntityData.YangName = "client-info"
    clientInfo.EntityData.BundleName = "cisco_ios_xr"
    clientInfo.EntityData.ParentYangName = "clients"
    clientInfo.EntityData.SegmentPath = "client-info" + types.AddNoKeyToken(clientInfo)
    clientInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-system/clients/" + clientInfo.EntityData.SegmentPath
    clientInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientInfo.EntityData.Children = types.NewOrderedMap()
    clientInfo.EntityData.Leafs = types.NewOrderedMap()
    clientInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", clientInfo.Name})
    clientInfo.EntityData.Leafs.Append("id", types.YLeaf{"Id", clientInfo.Id})
    clientInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", clientInfo.Location})
    clientInfo.EntityData.Leafs.Append("handle", types.YLeaf{"Handle", clientInfo.Handle})
    clientInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", clientInfo.State})
    clientInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", clientInfo.Type})
    clientInfo.EntityData.Leafs.Append("filter-disp", types.YLeaf{"FilterDisp", clientInfo.FilterDisp})
    clientInfo.EntityData.Leafs.Append("subscriber-id", types.YLeaf{"SubscriberId", clientInfo.SubscriberId})
    clientInfo.EntityData.Leafs.Append("filter-severity", types.YLeaf{"FilterSeverity", clientInfo.FilterSeverity})
    clientInfo.EntityData.Leafs.Append("filter-state", types.YLeaf{"FilterState", clientInfo.FilterState})
    clientInfo.EntityData.Leafs.Append("filter-group", types.YLeaf{"FilterGroup", clientInfo.FilterGroup})
    clientInfo.EntityData.Leafs.Append("connect-count", types.YLeaf{"ConnectCount", clientInfo.ConnectCount})
    clientInfo.EntityData.Leafs.Append("connect-timestamp", types.YLeaf{"ConnectTimestamp", clientInfo.ConnectTimestamp})
    clientInfo.EntityData.Leafs.Append("get-count", types.YLeaf{"GetCount", clientInfo.GetCount})
    clientInfo.EntityData.Leafs.Append("subscribe-count", types.YLeaf{"SubscribeCount", clientInfo.SubscribeCount})
    clientInfo.EntityData.Leafs.Append("report-count", types.YLeaf{"ReportCount", clientInfo.ReportCount})

    clientInfo.EntityData.YListKeys = []string {}

    return &(clientInfo.EntityData)
}

// Alarms_Detail_DetailCard
// Show detail card scope alarm related data.
type Alarms_Detail_DetailCard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of DetailLocation.
    DetailLocations Alarms_Detail_DetailCard_DetailLocations
}

func (detailCard *Alarms_Detail_DetailCard) GetEntityData() *types.CommonEntityData {
    detailCard.EntityData.YFilter = detailCard.YFilter
    detailCard.EntityData.YangName = "detail-card"
    detailCard.EntityData.BundleName = "cisco_ios_xr"
    detailCard.EntityData.ParentYangName = "detail"
    detailCard.EntityData.SegmentPath = "detail-card"
    detailCard.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/" + detailCard.EntityData.SegmentPath
    detailCard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailCard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailCard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailCard.EntityData.Children = types.NewOrderedMap()
    detailCard.EntityData.Children.Append("detail-locations", types.YChild{"DetailLocations", &detailCard.DetailLocations})
    detailCard.EntityData.Leafs = types.NewOrderedMap()

    detailCard.EntityData.YListKeys = []string {}

    return &(detailCard.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations
// Table of DetailLocation
type Alarms_Detail_DetailCard_DetailLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a card location for alarms. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation.
    DetailLocation []*Alarms_Detail_DetailCard_DetailLocations_DetailLocation
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetEntityData() *types.CommonEntityData {
    detailLocations.EntityData.YFilter = detailLocations.YFilter
    detailLocations.EntityData.YangName = "detail-locations"
    detailLocations.EntityData.BundleName = "cisco_ios_xr"
    detailLocations.EntityData.ParentYangName = "detail-card"
    detailLocations.EntityData.SegmentPath = "detail-locations"
    detailLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/" + detailLocations.EntityData.SegmentPath
    detailLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailLocations.EntityData.Children = types.NewOrderedMap()
    detailLocations.EntityData.Children.Append("detail-location", types.YChild{"DetailLocation", nil})
    for i := range detailLocations.DetailLocation {
        detailLocations.EntityData.Children.Append(types.GetSegmentPath(detailLocations.DetailLocation[i]), types.YChild{"DetailLocation", detailLocations.DetailLocation[i]})
    }
    detailLocations.EntityData.Leafs = types.NewOrderedMap()

    detailLocations.EntityData.YListKeys = []string {}

    return &(detailLocations.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation
// Specify a card location for alarms.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NodeID of the Location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Show the conditions present at this scope.
    Conditions Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions

    // Show the active alarms at this scope.
    Active Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active

    // Show the history alarms at this scope.
    History Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History

    // Show the suppressed alarms at this scope.
    Suppressed Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed

    // Show the service statistics.
    Stats Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats

    // Show the clients associated with this service.
    Clients Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients
}

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetEntityData() *types.CommonEntityData {
    detailLocation.EntityData.YFilter = detailLocation.YFilter
    detailLocation.EntityData.YangName = "detail-location"
    detailLocation.EntityData.BundleName = "cisco_ios_xr"
    detailLocation.EntityData.ParentYangName = "detail-locations"
    detailLocation.EntityData.SegmentPath = "detail-location" + types.AddKeyToken(detailLocation.NodeId, "node-id")
    detailLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/" + detailLocation.EntityData.SegmentPath
    detailLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailLocation.EntityData.Children = types.NewOrderedMap()
    detailLocation.EntityData.Children.Append("conditions", types.YChild{"Conditions", &detailLocation.Conditions})
    detailLocation.EntityData.Children.Append("active", types.YChild{"Active", &detailLocation.Active})
    detailLocation.EntityData.Children.Append("history", types.YChild{"History", &detailLocation.History})
    detailLocation.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", &detailLocation.Suppressed})
    detailLocation.EntityData.Children.Append("stats", types.YChild{"Stats", &detailLocation.Stats})
    detailLocation.EntityData.Children.Append("clients", types.YChild{"Clients", &detailLocation.Clients})
    detailLocation.EntityData.Leafs = types.NewOrderedMap()
    detailLocation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", detailLocation.NodeId})

    detailLocation.EntityData.YListKeys = []string {"NodeId"}

    return &(detailLocation.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions
// Show the conditions present at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo.
    AlarmInfo []*Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo
}

func (conditions *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions) GetEntityData() *types.CommonEntityData {
    conditions.EntityData.YFilter = conditions.YFilter
    conditions.EntityData.YangName = "conditions"
    conditions.EntityData.BundleName = "cisco_ios_xr"
    conditions.EntityData.ParentYangName = "detail-location"
    conditions.EntityData.SegmentPath = "conditions"
    conditions.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/" + conditions.EntityData.SegmentPath
    conditions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conditions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conditions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conditions.EntityData.Children = types.NewOrderedMap()
    conditions.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range conditions.AlarmInfo {
        types.SetYListKey(conditions.AlarmInfo[i], i)
        conditions.EntityData.Children.Append(types.GetSegmentPath(conditions.AlarmInfo[i]), types.YChild{"AlarmInfo", conditions.AlarmInfo[i]})
    }
    conditions.EntityData.Leafs = types.NewOrderedMap()

    conditions.EntityData.YListKeys = []string {}

    return &(conditions.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo
// Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm service affecting. The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // alarm event type. The type is AlarmEvent.
    Type interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Otn

    // TCA feature specific alarm attributes.
    Tca Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Tca
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "conditions"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/conditions/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &alarmInfo.Otn})
    alarmInfo.EntityData.Children.Append("tca", types.YChild{"Tca", &alarmInfo.Tca})
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", alarmInfo.Aid})
    alarmInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", alarmInfo.Tag})
    alarmInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", alarmInfo.Module})
    alarmInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", alarmInfo.Eid})
    alarmInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId})
    alarmInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", alarmInfo.PendingSync})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", alarmInfo.Status})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting})
    alarmInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", alarmInfo.Type})
    alarmInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", alarmInfo.Interface})
    alarmInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", alarmInfo.AlarmName})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "alarm-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/conditions/alarm-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Tca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Conditions_AlarmInfo_Tca) GetEntityData() *types.CommonEntityData {
    tca.EntityData.YFilter = tca.YFilter
    tca.EntityData.YangName = "tca"
    tca.EntityData.BundleName = "cisco_ios_xr"
    tca.EntityData.ParentYangName = "alarm-info"
    tca.EntityData.SegmentPath = "tca"
    tca.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/conditions/alarm-info/" + tca.EntityData.SegmentPath
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = types.NewOrderedMap()
    tca.EntityData.Leafs = types.NewOrderedMap()
    tca.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", tca.ThresholdValue})
    tca.EntityData.Leafs.Append("current-value", types.YLeaf{"CurrentValue", tca.CurrentValue})
    tca.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", tca.BucketType})

    tca.EntityData.YListKeys = []string {}

    return &(tca.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active
// Show the active alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo.
    AlarmInfo []*Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "detail-location"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range active.AlarmInfo {
        types.SetYListKey(active.AlarmInfo[i], i)
        active.EntityData.Children.Append(types.GetSegmentPath(active.AlarmInfo[i]), types.YChild{"AlarmInfo", active.AlarmInfo[i]})
    }
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo
// Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm service affecting. The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // alarm event type. The type is AlarmEvent.
    Type interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn

    // TCA feature specific alarm attributes.
    Tca Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "active"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/active/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &alarmInfo.Otn})
    alarmInfo.EntityData.Children.Append("tca", types.YChild{"Tca", &alarmInfo.Tca})
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", alarmInfo.Aid})
    alarmInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", alarmInfo.Tag})
    alarmInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", alarmInfo.Module})
    alarmInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", alarmInfo.Eid})
    alarmInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId})
    alarmInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", alarmInfo.PendingSync})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", alarmInfo.Status})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting})
    alarmInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", alarmInfo.Type})
    alarmInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", alarmInfo.Interface})
    alarmInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", alarmInfo.AlarmName})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "alarm-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/active/alarm-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetEntityData() *types.CommonEntityData {
    tca.EntityData.YFilter = tca.YFilter
    tca.EntityData.YangName = "tca"
    tca.EntityData.BundleName = "cisco_ios_xr"
    tca.EntityData.ParentYangName = "alarm-info"
    tca.EntityData.SegmentPath = "tca"
    tca.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/active/alarm-info/" + tca.EntityData.SegmentPath
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = types.NewOrderedMap()
    tca.EntityData.Leafs = types.NewOrderedMap()
    tca.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", tca.ThresholdValue})
    tca.EntityData.Leafs.Append("current-value", types.YLeaf{"CurrentValue", tca.CurrentValue})
    tca.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", tca.BucketType})

    tca.EntityData.YListKeys = []string {}

    return &(tca.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History
// Show the history alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo.
    AlarmInfo []*Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "detail-location"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range history.AlarmInfo {
        types.SetYListKey(history.AlarmInfo[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.AlarmInfo[i]), types.YChild{"AlarmInfo", history.AlarmInfo[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo
// Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm service affecting. The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // alarm event type. The type is AlarmEvent.
    Type interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn

    // TCA feature specific alarm attributes.
    Tca Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "history"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/history/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &alarmInfo.Otn})
    alarmInfo.EntityData.Children.Append("tca", types.YChild{"Tca", &alarmInfo.Tca})
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", alarmInfo.Aid})
    alarmInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", alarmInfo.Tag})
    alarmInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", alarmInfo.Module})
    alarmInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", alarmInfo.Eid})
    alarmInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId})
    alarmInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", alarmInfo.PendingSync})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", alarmInfo.Status})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting})
    alarmInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", alarmInfo.Type})
    alarmInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", alarmInfo.Interface})
    alarmInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", alarmInfo.AlarmName})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "alarm-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/history/alarm-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetEntityData() *types.CommonEntityData {
    tca.EntityData.YFilter = tca.YFilter
    tca.EntityData.YangName = "tca"
    tca.EntityData.BundleName = "cisco_ios_xr"
    tca.EntityData.ParentYangName = "alarm-info"
    tca.EntityData.SegmentPath = "tca"
    tca.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/history/alarm-info/" + tca.EntityData.SegmentPath
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = types.NewOrderedMap()
    tca.EntityData.Leafs = types.NewOrderedMap()
    tca.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", tca.ThresholdValue})
    tca.EntityData.Leafs.Append("current-value", types.YLeaf{"CurrentValue", tca.CurrentValue})
    tca.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", tca.BucketType})

    tca.EntityData.YListKeys = []string {}

    return &(tca.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo.
    SuppressedInfo []*Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "detail-location"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Children.Append("suppressed-info", types.YChild{"SuppressedInfo", nil})
    for i := range suppressed.SuppressedInfo {
        types.SetYListKey(suppressed.SuppressedInfo[i], i)
        suppressed.EntityData.Children.Append(types.GetSegmentPath(suppressed.SuppressedInfo[i]), types.YChild{"SuppressedInfo", suppressed.SuppressedInfo[i]})
    }
    suppressed.EntityData.Leafs = types.NewOrderedMap()

    suppressed.EntityData.YListKeys = []string {}

    return &(suppressed.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm description. The type is string with length: 0..256.
    Description interface{}

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm aid. The type is string with length: 0..128.
    Aid interface{}

    // Alarm tag description. The type is string with length: 0..128.
    Tag interface{}

    // Alarm module description. The type is string with length: 0..128.
    Module interface{}

    // Alarm eid. The type is string with length: 0..128.
    Eid interface{}

    // Reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // Pending async flag. The type is bool.
    PendingSync interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm status. The type is AlarmStatus.
    Status interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm suppressed time. The type is string with length: 0..64.
    SuppressedTime interface{}

    // Alarm suppressed time(timestamp format). The type is interface{} with
    // range: 0..18446744073709551615.
    SuppressedTimestamp interface{}

    // Alarm service affecting . The type is AlarmServiceAffecting.
    ServiceAffecting interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn
}

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetEntityData() *types.CommonEntityData {
    suppressedInfo.EntityData.YFilter = suppressedInfo.YFilter
    suppressedInfo.EntityData.YangName = "suppressed-info"
    suppressedInfo.EntityData.BundleName = "cisco_ios_xr"
    suppressedInfo.EntityData.ParentYangName = "suppressed"
    suppressedInfo.EntityData.SegmentPath = "suppressed-info" + types.AddNoKeyToken(suppressedInfo)
    suppressedInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/suppressed/" + suppressedInfo.EntityData.SegmentPath
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = types.NewOrderedMap()
    suppressedInfo.EntityData.Children.Append("otn", types.YChild{"Otn", &suppressedInfo.Otn})
    suppressedInfo.EntityData.Leafs = types.NewOrderedMap()
    suppressedInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressedInfo.Description})
    suppressedInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", suppressedInfo.Location})
    suppressedInfo.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressedInfo.Aid})
    suppressedInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressedInfo.Tag})
    suppressedInfo.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressedInfo.Module})
    suppressedInfo.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressedInfo.Eid})
    suppressedInfo.EntityData.Leafs.Append("reporting-agent-id", types.YLeaf{"ReportingAgentId", suppressedInfo.ReportingAgentId})
    suppressedInfo.EntityData.Leafs.Append("pending-sync", types.YLeaf{"PendingSync", suppressedInfo.PendingSync})
    suppressedInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressedInfo.Severity})
    suppressedInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", suppressedInfo.Status})
    suppressedInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressedInfo.Group})
    suppressedInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", suppressedInfo.SetTime})
    suppressedInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp})
    suppressedInfo.EntityData.Leafs.Append("suppressed-time", types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime})
    suppressedInfo.EntityData.Leafs.Append("suppressed-timestamp", types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp})
    suppressedInfo.EntityData.Leafs.Append("service-affecting", types.YLeaf{"ServiceAffecting", suppressedInfo.ServiceAffecting})
    suppressedInfo.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", suppressedInfo.Interface})
    suppressedInfo.EntityData.Leafs.Append("alarm-name", types.YLeaf{"AlarmName", suppressedInfo.AlarmName})

    suppressedInfo.EntityData.YListKeys = []string {}

    return &(suppressedInfo.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "suppressed-info"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/suppressed/suppressed-info/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otn.Direction})
    otn.EntityData.Leafs.Append("notification-source", types.YLeaf{"NotificationSource", otn.NotificationSource})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats
// Show the service statistics.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarms that were in all reported to this Alarm Mgr. The type is interface{}
    // with range: 0..18446744073709551615.
    Reported interface{}

    // Alarms that we couldn't keep track due to some error or other. The type is
    // interface{} with range: 0..18446744073709551615.
    Dropped interface{}

    // Alarms that are currently in the active state. The type is interface{} with
    // range: 0..18446744073709551615.
    Active interface{}

    // Alarms that are cleared. This one is counted over a long period of time.
    // The type is interface{} with range: 0..18446744073709551615.
    History interface{}

    // Alarms that are in suppressed state. The type is interface{} with range:
    // 0..18446744073709551615.
    Suppressed interface{}

    // Alarms that are currently in the active state(sysadmin plane). The type is
    // interface{} with range: 0..18446744073709551615.
    SysadminActive interface{}

    // Alarms that are cleared in sysadmin plane. This one is counted over a long
    // period of time. The type is interface{} with range:
    // 0..18446744073709551615.
    SysadminHistory interface{}

    // Alarms that are suppressed in sysadmin plane. The type is interface{} with
    // range: 0..18446744073709551615.
    SysadminSuppressed interface{}

    // Alarms dropped due to invalid aid. The type is interface{} with range:
    // 0..4294967295.
    DroppedInvalidAid interface{}

    // Alarms dropped due to insufficient memory. The type is interface{} with
    // range: 0..4294967295.
    DroppedInsuffMem interface{}

    // Alarms dropped due to db error. The type is interface{} with range:
    // 0..4294967295.
    DroppedDbError interface{}

    // Alarms dropped clear without set. The type is interface{} with range:
    // 0..4294967295.
    DroppedClearWithoutSet interface{}

    // Alarms dropped which were duplicate. The type is interface{} with range:
    // 0..4294967295.
    DroppedDuplicate interface{}

    // Total alarms which had the cache hit. The type is interface{} with range:
    // 0..4294967295.
    CacheHit interface{}

    // Total alarms which had the cache miss. The type is interface{} with range:
    // 0..4294967295.
    CacheMiss interface{}
}

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "detail-location"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("reported", types.YLeaf{"Reported", stats.Reported})
    stats.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", stats.Dropped})
    stats.EntityData.Leafs.Append("active", types.YLeaf{"Active", stats.Active})
    stats.EntityData.Leafs.Append("history", types.YLeaf{"History", stats.History})
    stats.EntityData.Leafs.Append("suppressed", types.YLeaf{"Suppressed", stats.Suppressed})
    stats.EntityData.Leafs.Append("sysadmin-active", types.YLeaf{"SysadminActive", stats.SysadminActive})
    stats.EntityData.Leafs.Append("sysadmin-history", types.YLeaf{"SysadminHistory", stats.SysadminHistory})
    stats.EntityData.Leafs.Append("sysadmin-suppressed", types.YLeaf{"SysadminSuppressed", stats.SysadminSuppressed})
    stats.EntityData.Leafs.Append("dropped-invalid-aid", types.YLeaf{"DroppedInvalidAid", stats.DroppedInvalidAid})
    stats.EntityData.Leafs.Append("dropped-insuff-mem", types.YLeaf{"DroppedInsuffMem", stats.DroppedInsuffMem})
    stats.EntityData.Leafs.Append("dropped-db-error", types.YLeaf{"DroppedDbError", stats.DroppedDbError})
    stats.EntityData.Leafs.Append("dropped-clear-without-set", types.YLeaf{"DroppedClearWithoutSet", stats.DroppedClearWithoutSet})
    stats.EntityData.Leafs.Append("dropped-duplicate", types.YLeaf{"DroppedDuplicate", stats.DroppedDuplicate})
    stats.EntityData.Leafs.Append("cache-hit", types.YLeaf{"CacheHit", stats.CacheHit})
    stats.EntityData.Leafs.Append("cache-miss", types.YLeaf{"CacheMiss", stats.CacheMiss})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients
// Show the clients associated with this
// service.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo.
    ClientInfo []*Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "detail-location"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client-info", types.YChild{"ClientInfo", nil})
    for i := range clients.ClientInfo {
        types.SetYListKey(clients.ClientInfo[i], i)
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.ClientInfo[i]), types.YChild{"ClientInfo", clients.ClientInfo[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo
// Client List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm client. The type is string with length: 0..128.
    Name interface{}

    // Alarms agent id of the client. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // The location of this client. The type is string with length: 0..128.
    Location interface{}

    // The client handle through which interface. The type is string with length:
    // 0..128.
    Handle interface{}

    // The current state of the client. The type is AlarmClientState.
    State interface{}

    // The type of the client. The type is AlarmClient.
    Type interface{}

    // The current subscription status of the client. The type is bool.
    FilterDisp interface{}

    // Alarms agent subscriber id of the client. The type is interface{} with
    // range: 0..4294967295.
    SubscriberId interface{}

    // The filter used for alarm severity. The type is AlarmSeverity.
    FilterSeverity interface{}

    // The filter used for alarm bi-state state+. The type is AlarmStatus.
    FilterState interface{}

    // The filter used for alarm group. The type is AlarmGroups.
    FilterGroup interface{}

    // Number of times the agent connected to the alarm mgr. The type is
    // interface{} with range: 0..4294967295.
    ConnectCount interface{}

    // Agent connect timestamp. The type is string with length: 0..64.
    ConnectTimestamp interface{}

    // Number of times the agent queried for alarms. The type is interface{} with
    // range: 0..4294967295.
    GetCount interface{}

    // Number of times the agent subscribed for alarms. The type is interface{}
    // with range: 0..4294967295.
    SubscribeCount interface{}

    // Number of times the agent reported alarms. The type is interface{} with
    // range: 0..4294967295.
    ReportCount interface{}
}

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetEntityData() *types.CommonEntityData {
    clientInfo.EntityData.YFilter = clientInfo.YFilter
    clientInfo.EntityData.YangName = "client-info"
    clientInfo.EntityData.BundleName = "cisco_ios_xr"
    clientInfo.EntityData.ParentYangName = "clients"
    clientInfo.EntityData.SegmentPath = "client-info" + types.AddNoKeyToken(clientInfo)
    clientInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/detail/detail-card/detail-locations/detail-location/clients/" + clientInfo.EntityData.SegmentPath
    clientInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientInfo.EntityData.Children = types.NewOrderedMap()
    clientInfo.EntityData.Leafs = types.NewOrderedMap()
    clientInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", clientInfo.Name})
    clientInfo.EntityData.Leafs.Append("id", types.YLeaf{"Id", clientInfo.Id})
    clientInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", clientInfo.Location})
    clientInfo.EntityData.Leafs.Append("handle", types.YLeaf{"Handle", clientInfo.Handle})
    clientInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", clientInfo.State})
    clientInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", clientInfo.Type})
    clientInfo.EntityData.Leafs.Append("filter-disp", types.YLeaf{"FilterDisp", clientInfo.FilterDisp})
    clientInfo.EntityData.Leafs.Append("subscriber-id", types.YLeaf{"SubscriberId", clientInfo.SubscriberId})
    clientInfo.EntityData.Leafs.Append("filter-severity", types.YLeaf{"FilterSeverity", clientInfo.FilterSeverity})
    clientInfo.EntityData.Leafs.Append("filter-state", types.YLeaf{"FilterState", clientInfo.FilterState})
    clientInfo.EntityData.Leafs.Append("filter-group", types.YLeaf{"FilterGroup", clientInfo.FilterGroup})
    clientInfo.EntityData.Leafs.Append("connect-count", types.YLeaf{"ConnectCount", clientInfo.ConnectCount})
    clientInfo.EntityData.Leafs.Append("connect-timestamp", types.YLeaf{"ConnectTimestamp", clientInfo.ConnectTimestamp})
    clientInfo.EntityData.Leafs.Append("get-count", types.YLeaf{"GetCount", clientInfo.GetCount})
    clientInfo.EntityData.Leafs.Append("subscribe-count", types.YLeaf{"SubscribeCount", clientInfo.SubscribeCount})
    clientInfo.EntityData.Leafs.Append("report-count", types.YLeaf{"ReportCount", clientInfo.ReportCount})

    clientInfo.EntityData.YListKeys = []string {}

    return &(clientInfo.EntityData)
}

// Alarms_Brief
// A set of brief alarm commands.
type Alarms_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show brief card scope alarm related data.
    BriefCard Alarms_Brief_BriefCard

    // Show brief system scope alarm related data.
    BriefSystem Alarms_Brief_BriefSystem

    // Show brief system scope alarm related data.
    AlarmId Alarms_Brief_AlarmId
}

func (brief *Alarms_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "alarms"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/" + brief.EntityData.SegmentPath
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Children.Append("brief-card", types.YChild{"BriefCard", &brief.BriefCard})
    brief.EntityData.Children.Append("brief-system", types.YChild{"BriefSystem", &brief.BriefSystem})
    brief.EntityData.Children.Append("alarm-id", types.YChild{"AlarmId", &brief.AlarmId})
    brief.EntityData.Leafs = types.NewOrderedMap()

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// Alarms_Brief_BriefCard
// Show brief card scope alarm related data.
type Alarms_Brief_BriefCard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of BriefLocation.
    BriefLocations Alarms_Brief_BriefCard_BriefLocations
}

func (briefCard *Alarms_Brief_BriefCard) GetEntityData() *types.CommonEntityData {
    briefCard.EntityData.YFilter = briefCard.YFilter
    briefCard.EntityData.YangName = "brief-card"
    briefCard.EntityData.BundleName = "cisco_ios_xr"
    briefCard.EntityData.ParentYangName = "brief"
    briefCard.EntityData.SegmentPath = "brief-card"
    briefCard.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/" + briefCard.EntityData.SegmentPath
    briefCard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefCard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefCard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefCard.EntityData.Children = types.NewOrderedMap()
    briefCard.EntityData.Children.Append("brief-locations", types.YChild{"BriefLocations", &briefCard.BriefLocations})
    briefCard.EntityData.Leafs = types.NewOrderedMap()

    briefCard.EntityData.YListKeys = []string {}

    return &(briefCard.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations
// Table of BriefLocation
type Alarms_Brief_BriefCard_BriefLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a card location for alarms. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation.
    BriefLocation []*Alarms_Brief_BriefCard_BriefLocations_BriefLocation
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetEntityData() *types.CommonEntityData {
    briefLocations.EntityData.YFilter = briefLocations.YFilter
    briefLocations.EntityData.YangName = "brief-locations"
    briefLocations.EntityData.BundleName = "cisco_ios_xr"
    briefLocations.EntityData.ParentYangName = "brief-card"
    briefLocations.EntityData.SegmentPath = "brief-locations"
    briefLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/" + briefLocations.EntityData.SegmentPath
    briefLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLocations.EntityData.Children = types.NewOrderedMap()
    briefLocations.EntityData.Children.Append("brief-location", types.YChild{"BriefLocation", nil})
    for i := range briefLocations.BriefLocation {
        briefLocations.EntityData.Children.Append(types.GetSegmentPath(briefLocations.BriefLocation[i]), types.YChild{"BriefLocation", briefLocations.BriefLocation[i]})
    }
    briefLocations.EntityData.Leafs = types.NewOrderedMap()

    briefLocations.EntityData.YListKeys = []string {}

    return &(briefLocations.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation
// Specify a card location for alarms.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NodeID of the Location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Show the conditions present at this scope.
    Conditions Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions

    // Show the active alarms at this scope.
    Active Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active

    // Show the history alarms at this scope.
    History Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History

    // Show the suppressed alarms at this scope.
    Suppressed Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed
}

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetEntityData() *types.CommonEntityData {
    briefLocation.EntityData.YFilter = briefLocation.YFilter
    briefLocation.EntityData.YangName = "brief-location"
    briefLocation.EntityData.BundleName = "cisco_ios_xr"
    briefLocation.EntityData.ParentYangName = "brief-locations"
    briefLocation.EntityData.SegmentPath = "brief-location" + types.AddKeyToken(briefLocation.NodeId, "node-id")
    briefLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/" + briefLocation.EntityData.SegmentPath
    briefLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLocation.EntityData.Children = types.NewOrderedMap()
    briefLocation.EntityData.Children.Append("conditions", types.YChild{"Conditions", &briefLocation.Conditions})
    briefLocation.EntityData.Children.Append("active", types.YChild{"Active", &briefLocation.Active})
    briefLocation.EntityData.Children.Append("history", types.YChild{"History", &briefLocation.History})
    briefLocation.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", &briefLocation.Suppressed})
    briefLocation.EntityData.Leafs = types.NewOrderedMap()
    briefLocation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", briefLocation.NodeId})

    briefLocation.EntityData.YListKeys = []string {"NodeId"}

    return &(briefLocation.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions
// Show the conditions present at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions_AlarmInfo.
    AlarmInfo []*Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions_AlarmInfo
}

func (conditions *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions) GetEntityData() *types.CommonEntityData {
    conditions.EntityData.YFilter = conditions.YFilter
    conditions.EntityData.YangName = "conditions"
    conditions.EntityData.BundleName = "cisco_ios_xr"
    conditions.EntityData.ParentYangName = "brief-location"
    conditions.EntityData.SegmentPath = "conditions"
    conditions.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/" + conditions.EntityData.SegmentPath
    conditions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conditions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conditions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conditions.EntityData.Children = types.NewOrderedMap()
    conditions.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range conditions.AlarmInfo {
        types.SetYListKey(conditions.AlarmInfo[i], i)
        conditions.EntityData.Children.Append(types.GetSegmentPath(conditions.AlarmInfo[i]), types.YChild{"AlarmInfo", conditions.AlarmInfo[i]})
    }
    conditions.EntityData.Leafs = types.NewOrderedMap()

    conditions.EntityData.YListKeys = []string {}

    return &(conditions.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions_AlarmInfo
// Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Conditions_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "conditions"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/conditions/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active
// Show the active alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo.
    AlarmInfo []*Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "brief-location"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range active.AlarmInfo {
        types.SetYListKey(active.AlarmInfo[i], i)
        active.EntityData.Children.Append(types.GetSegmentPath(active.AlarmInfo[i]), types.YChild{"AlarmInfo", active.AlarmInfo[i]})
    }
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo
// Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "active"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/active/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History
// Show the history alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo.
    AlarmInfo []*Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "brief-location"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range history.AlarmInfo {
        types.SetYListKey(history.AlarmInfo[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.AlarmInfo[i]), types.YChild{"AlarmInfo", history.AlarmInfo[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo
// Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "history"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/history/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo.
    SuppressedInfo []*Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "brief-location"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Children.Append("suppressed-info", types.YChild{"SuppressedInfo", nil})
    for i := range suppressed.SuppressedInfo {
        types.SetYListKey(suppressed.SuppressedInfo[i], i)
        suppressed.EntityData.Children.Append(types.GetSegmentPath(suppressed.SuppressedInfo[i]), types.YChild{"SuppressedInfo", suppressed.SuppressedInfo[i]})
    }
    suppressed.EntityData.Leafs = types.NewOrderedMap()

    suppressed.EntityData.YListKeys = []string {}

    return &(suppressed.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm suppressed time. The type is string with length: 0..64.
    SuppressedTime interface{}

    // Alarm suppressed time(timestamp format). The type is interface{} with
    // range: 0..18446744073709551615.
    SuppressedTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetEntityData() *types.CommonEntityData {
    suppressedInfo.EntityData.YFilter = suppressedInfo.YFilter
    suppressedInfo.EntityData.YangName = "suppressed-info"
    suppressedInfo.EntityData.BundleName = "cisco_ios_xr"
    suppressedInfo.EntityData.ParentYangName = "suppressed"
    suppressedInfo.EntityData.SegmentPath = "suppressed-info" + types.AddNoKeyToken(suppressedInfo)
    suppressedInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-card/brief-locations/brief-location/suppressed/" + suppressedInfo.EntityData.SegmentPath
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = types.NewOrderedMap()
    suppressedInfo.EntityData.Leafs = types.NewOrderedMap()
    suppressedInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", suppressedInfo.Location})
    suppressedInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressedInfo.Severity})
    suppressedInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressedInfo.Group})
    suppressedInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", suppressedInfo.SetTime})
    suppressedInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp})
    suppressedInfo.EntityData.Leafs.Append("suppressed-time", types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime})
    suppressedInfo.EntityData.Leafs.Append("suppressed-timestamp", types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp})
    suppressedInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressedInfo.Description})

    suppressedInfo.EntityData.YListKeys = []string {}

    return &(suppressedInfo.EntityData)
}

// Alarms_Brief_BriefSystem
// Show brief system scope alarm related data.
type Alarms_Brief_BriefSystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show the conditions present at this scope.
    Conditions Alarms_Brief_BriefSystem_Conditions

    // Show the active alarms at this scope.
    Active Alarms_Brief_BriefSystem_Active

    // Show the history alarms at this scope.
    History Alarms_Brief_BriefSystem_History

    // Show the suppressed alarms at this scope.
    Suppressed Alarms_Brief_BriefSystem_Suppressed
}

func (briefSystem *Alarms_Brief_BriefSystem) GetEntityData() *types.CommonEntityData {
    briefSystem.EntityData.YFilter = briefSystem.YFilter
    briefSystem.EntityData.YangName = "brief-system"
    briefSystem.EntityData.BundleName = "cisco_ios_xr"
    briefSystem.EntityData.ParentYangName = "brief"
    briefSystem.EntityData.SegmentPath = "brief-system"
    briefSystem.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/" + briefSystem.EntityData.SegmentPath
    briefSystem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefSystem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefSystem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefSystem.EntityData.Children = types.NewOrderedMap()
    briefSystem.EntityData.Children.Append("conditions", types.YChild{"Conditions", &briefSystem.Conditions})
    briefSystem.EntityData.Children.Append("active", types.YChild{"Active", &briefSystem.Active})
    briefSystem.EntityData.Children.Append("history", types.YChild{"History", &briefSystem.History})
    briefSystem.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", &briefSystem.Suppressed})
    briefSystem.EntityData.Leafs = types.NewOrderedMap()

    briefSystem.EntityData.YListKeys = []string {}

    return &(briefSystem.EntityData)
}

// Alarms_Brief_BriefSystem_Conditions
// Show the conditions present at this scope.
type Alarms_Brief_BriefSystem_Conditions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefSystem_Conditions_AlarmInfo.
    AlarmInfo []*Alarms_Brief_BriefSystem_Conditions_AlarmInfo
}

func (conditions *Alarms_Brief_BriefSystem_Conditions) GetEntityData() *types.CommonEntityData {
    conditions.EntityData.YFilter = conditions.YFilter
    conditions.EntityData.YangName = "conditions"
    conditions.EntityData.BundleName = "cisco_ios_xr"
    conditions.EntityData.ParentYangName = "brief-system"
    conditions.EntityData.SegmentPath = "conditions"
    conditions.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/" + conditions.EntityData.SegmentPath
    conditions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conditions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conditions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conditions.EntityData.Children = types.NewOrderedMap()
    conditions.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range conditions.AlarmInfo {
        types.SetYListKey(conditions.AlarmInfo[i], i)
        conditions.EntityData.Children.Append(types.GetSegmentPath(conditions.AlarmInfo[i]), types.YChild{"AlarmInfo", conditions.AlarmInfo[i]})
    }
    conditions.EntityData.Leafs = types.NewOrderedMap()

    conditions.EntityData.YListKeys = []string {}

    return &(conditions.EntityData)
}

// Alarms_Brief_BriefSystem_Conditions_AlarmInfo
// Alarm List
type Alarms_Brief_BriefSystem_Conditions_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (alarmInfo *Alarms_Brief_BriefSystem_Conditions_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "conditions"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/conditions/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefSystem_Active
// Show the active alarms at this scope.
type Alarms_Brief_BriefSystem_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of Alarms_Brief_BriefSystem_Active_AlarmInfo.
    AlarmInfo []*Alarms_Brief_BriefSystem_Active_AlarmInfo
}

func (active *Alarms_Brief_BriefSystem_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "brief-system"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range active.AlarmInfo {
        types.SetYListKey(active.AlarmInfo[i], i)
        active.EntityData.Children.Append(types.GetSegmentPath(active.AlarmInfo[i]), types.YChild{"AlarmInfo", active.AlarmInfo[i]})
    }
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Alarms_Brief_BriefSystem_Active_AlarmInfo
// Alarm List
type Alarms_Brief_BriefSystem_Active_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "active"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/active/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefSystem_History
// Show the history alarms at this scope.
type Alarms_Brief_BriefSystem_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefSystem_History_AlarmInfo.
    AlarmInfo []*Alarms_Brief_BriefSystem_History_AlarmInfo
}

func (history *Alarms_Brief_BriefSystem_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "brief-system"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", nil})
    for i := range history.AlarmInfo {
        types.SetYListKey(history.AlarmInfo[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.AlarmInfo[i]), types.YChild{"AlarmInfo", history.AlarmInfo[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Alarms_Brief_BriefSystem_History_AlarmInfo
// Alarm List
type Alarms_Brief_BriefSystem_History_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "history"
    alarmInfo.EntityData.SegmentPath = "alarm-info" + types.AddNoKeyToken(alarmInfo)
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/history/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarmInfo.Location})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarmInfo.SetTime})
    alarmInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp})
    alarmInfo.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarmInfo.ClearTime})
    alarmInfo.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp})
    alarmInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarmInfo.Description})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefSystem_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Brief_BriefSystem_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo.
    SuppressedInfo []*Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "brief-system"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Children.Append("suppressed-info", types.YChild{"SuppressedInfo", nil})
    for i := range suppressed.SuppressedInfo {
        types.SetYListKey(suppressed.SuppressedInfo[i], i)
        suppressed.EntityData.Children.Append(types.GetSegmentPath(suppressed.SuppressedInfo[i]), types.YChild{"SuppressedInfo", suppressed.SuppressedInfo[i]})
    }
    suppressed.EntityData.Leafs = types.NewOrderedMap()

    suppressed.EntityData.YListKeys = []string {}

    return &(suppressed.EntityData)
}

// Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm suppressed time. The type is string with length: 0..64.
    SuppressedTime interface{}

    // Alarm suppressed time(timestamp format). The type is interface{} with
    // range: 0..18446744073709551615.
    SuppressedTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetEntityData() *types.CommonEntityData {
    suppressedInfo.EntityData.YFilter = suppressedInfo.YFilter
    suppressedInfo.EntityData.YangName = "suppressed-info"
    suppressedInfo.EntityData.BundleName = "cisco_ios_xr"
    suppressedInfo.EntityData.ParentYangName = "suppressed"
    suppressedInfo.EntityData.SegmentPath = "suppressed-info" + types.AddNoKeyToken(suppressedInfo)
    suppressedInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/brief-system/suppressed/" + suppressedInfo.EntityData.SegmentPath
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = types.NewOrderedMap()
    suppressedInfo.EntityData.Leafs = types.NewOrderedMap()
    suppressedInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", suppressedInfo.Location})
    suppressedInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressedInfo.Severity})
    suppressedInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressedInfo.Group})
    suppressedInfo.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", suppressedInfo.SetTime})
    suppressedInfo.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp})
    suppressedInfo.EntityData.Leafs.Append("suppressed-time", types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime})
    suppressedInfo.EntityData.Leafs.Append("suppressed-timestamp", types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp})
    suppressedInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressedInfo.Description})

    suppressedInfo.EntityData.YListKeys = []string {}

    return &(suppressedInfo.EntityData)
}

// Alarms_Brief_AlarmId
// Show brief system scope alarm related data.
type Alarms_Brief_AlarmId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of ActiveAlarm.
    ActiveAlarms Alarms_Brief_AlarmId_ActiveAlarms
}

func (alarmId *Alarms_Brief_AlarmId) GetEntityData() *types.CommonEntityData {
    alarmId.EntityData.YFilter = alarmId.YFilter
    alarmId.EntityData.YangName = "alarm-id"
    alarmId.EntityData.BundleName = "cisco_ios_xr"
    alarmId.EntityData.ParentYangName = "brief"
    alarmId.EntityData.SegmentPath = "alarm-id"
    alarmId.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/" + alarmId.EntityData.SegmentPath
    alarmId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmId.EntityData.Children = types.NewOrderedMap()
    alarmId.EntityData.Children.Append("active-alarms", types.YChild{"ActiveAlarms", &alarmId.ActiveAlarms})
    alarmId.EntityData.Leafs = types.NewOrderedMap()

    alarmId.EntityData.YListKeys = []string {}

    return &(alarmId.EntityData)
}

// Alarms_Brief_AlarmId_ActiveAlarms
// Table of ActiveAlarm
type Alarms_Brief_AlarmId_ActiveAlarms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show the active alarms at this scope. The type is slice of
    // Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm.
    ActiveAlarm []*Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm
}

func (activeAlarms *Alarms_Brief_AlarmId_ActiveAlarms) GetEntityData() *types.CommonEntityData {
    activeAlarms.EntityData.YFilter = activeAlarms.YFilter
    activeAlarms.EntityData.YangName = "active-alarms"
    activeAlarms.EntityData.BundleName = "cisco_ios_xr"
    activeAlarms.EntityData.ParentYangName = "alarm-id"
    activeAlarms.EntityData.SegmentPath = "active-alarms"
    activeAlarms.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/alarm-id/" + activeAlarms.EntityData.SegmentPath
    activeAlarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAlarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAlarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAlarms.EntityData.Children = types.NewOrderedMap()
    activeAlarms.EntityData.Children.Append("active-alarm", types.YChild{"ActiveAlarm", nil})
    for i := range activeAlarms.ActiveAlarm {
        activeAlarms.EntityData.Children.Append(types.GetSegmentPath(activeAlarms.ActiveAlarm[i]), types.YChild{"ActiveAlarm", activeAlarms.ActiveAlarm[i]})
    }
    activeAlarms.EntityData.Leafs = types.NewOrderedMap()

    activeAlarms.EntityData.YListKeys = []string {}

    return &(activeAlarms.EntityData)
}

// Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm
// Show the active alarms at this scope.
type Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Alarm ID. The type is string.
    Aid interface{}

    // Alarm Brief.
    Alarm Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm_Alarm
}

func (activeAlarm *Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm) GetEntityData() *types.CommonEntityData {
    activeAlarm.EntityData.YFilter = activeAlarm.YFilter
    activeAlarm.EntityData.YangName = "active-alarm"
    activeAlarm.EntityData.BundleName = "cisco_ios_xr"
    activeAlarm.EntityData.ParentYangName = "active-alarms"
    activeAlarm.EntityData.SegmentPath = "active-alarm" + types.AddKeyToken(activeAlarm.Aid, "aid")
    activeAlarm.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/alarm-id/active-alarms/" + activeAlarm.EntityData.SegmentPath
    activeAlarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAlarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAlarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAlarm.EntityData.Children = types.NewOrderedMap()
    activeAlarm.EntityData.Children.Append("alarm", types.YChild{"Alarm", &activeAlarm.Alarm})
    activeAlarm.EntityData.Leafs = types.NewOrderedMap()
    activeAlarm.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", activeAlarm.Aid})

    activeAlarm.EntityData.YListKeys = []string {"Aid"}

    return &(activeAlarm.EntityData)
}

// Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm_Alarm
// Alarm Brief
type Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm_Alarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm location. The type is string with length: 0..128.
    Location interface{}

    // Alarm severity. The type is AlarmSeverity.
    Severity interface{}

    // Alarm group. The type is AlarmGroups.
    Group interface{}

    // Alarm set time. The type is string with length: 0..64.
    SetTime interface{}

    // Alarm set time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    SetTimestamp interface{}

    // Alarm clear time. The type is string with length: 0..64.
    ClearTime interface{}

    // Alarm clear time(timestamp format). The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTimestamp interface{}

    // Alarm description. The type is string with length: 0..256.
    Description interface{}
}

func (alarm *Alarms_Brief_AlarmId_ActiveAlarms_ActiveAlarm_Alarm) GetEntityData() *types.CommonEntityData {
    alarm.EntityData.YFilter = alarm.YFilter
    alarm.EntityData.YangName = "alarm"
    alarm.EntityData.BundleName = "cisco_ios_xr"
    alarm.EntityData.ParentYangName = "active-alarm"
    alarm.EntityData.SegmentPath = "alarm"
    alarm.EntityData.AbsolutePath = "Cisco-IOS-XR-alarmgr-server-oper:alarms/brief/alarm-id/active-alarms/active-alarm/" + alarm.EntityData.SegmentPath
    alarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarm.EntityData.Children = types.NewOrderedMap()
    alarm.EntityData.Leafs = types.NewOrderedMap()
    alarm.EntityData.Leafs.Append("location", types.YLeaf{"Location", alarm.Location})
    alarm.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarm.Severity})
    alarm.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarm.Group})
    alarm.EntityData.Leafs.Append("set-time", types.YLeaf{"SetTime", alarm.SetTime})
    alarm.EntityData.Leafs.Append("set-timestamp", types.YLeaf{"SetTimestamp", alarm.SetTimestamp})
    alarm.EntityData.Leafs.Append("clear-time", types.YLeaf{"ClearTime", alarm.ClearTime})
    alarm.EntityData.Leafs.Append("clear-timestamp", types.YLeaf{"ClearTimestamp", alarm.ClearTimestamp})
    alarm.EntityData.Leafs.Append("description", types.YLeaf{"Description", alarm.Description})

    alarm.EntityData.YListKeys = []string {}

    return &(alarm.EntityData)
}

