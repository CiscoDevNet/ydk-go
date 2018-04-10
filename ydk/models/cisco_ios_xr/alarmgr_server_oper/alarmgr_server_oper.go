// This module contains a collection of YANG definitions
// for Cisco IOS-XR alarmgr-server package operational data.
// 
// This module contains definitions
// for the following management objects:
//   alarms: Show Alarms associated with XR
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// AlarmEvent represents Alarm event
type AlarmEvent string

const (
    // Default Alarm Event Type
    AlarmEvent_default_ AlarmEvent = "default"

    // Alarm Notifcation Event Type
    AlarmEvent_notification AlarmEvent = "notification"

    // Last Event Type
    AlarmEvent_last AlarmEvent = "last"
)

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
    alarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarms.EntityData.Children = make(map[string]types.YChild)
    alarms.EntityData.Children["detail"] = types.YChild{"Detail", &alarms.Detail}
    alarms.EntityData.Children["brief"] = types.YChild{"Brief", &alarms.Brief}
    alarms.EntityData.Leafs = make(map[string]types.YLeaf)
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
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["detail-system"] = types.YChild{"DetailSystem", &detail.DetailSystem}
    detail.EntityData.Children["detail-card"] = types.YChild{"DetailCard", &detail.DetailCard}
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detail.EntityData)
}

// Alarms_Detail_DetailSystem
// show detail system scope alarm related data.
type Alarms_Detail_DetailSystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    detailSystem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailSystem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailSystem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailSystem.EntityData.Children = make(map[string]types.YChild)
    detailSystem.EntityData.Children["active"] = types.YChild{"Active", &detailSystem.Active}
    detailSystem.EntityData.Children["history"] = types.YChild{"History", &detailSystem.History}
    detailSystem.EntityData.Children["suppressed"] = types.YChild{"Suppressed", &detailSystem.Suppressed}
    detailSystem.EntityData.Children["stats"] = types.YChild{"Stats", &detailSystem.Stats}
    detailSystem.EntityData.Children["clients"] = types.YChild{"Clients", &detailSystem.Clients}
    detailSystem.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detailSystem.EntityData)
}

// Alarms_Detail_DetailSystem_Active
// Show the active alarms at this scope.
type Alarms_Detail_DetailSystem_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_Active_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailSystem_Active_AlarmInfo
}

func (active *Alarms_Detail_DetailSystem_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "detail-system"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range active.AlarmInfo {
        active.EntityData.Children[types.GetSegmentPath(&active.AlarmInfo[i])] = types.YChild{"AlarmInfo", &active.AlarmInfo[i]}
    }
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(active.EntityData)
}

// Alarms_Detail_DetailSystem_Active_AlarmInfo
// Alarm List
type Alarms_Detail_DetailSystem_Active_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Type_ interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface_ interface{}

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
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Children["otn"] = types.YChild{"Otn", &alarmInfo.Otn}
    alarmInfo.EntityData.Children["tca"] = types.YChild{"Tca", &alarmInfo.Tca}
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["aid"] = types.YLeaf{"Aid", alarmInfo.Aid}
    alarmInfo.EntityData.Leafs["tag"] = types.YLeaf{"Tag", alarmInfo.Tag}
    alarmInfo.EntityData.Leafs["module"] = types.YLeaf{"Module", alarmInfo.Module}
    alarmInfo.EntityData.Leafs["eid"] = types.YLeaf{"Eid", alarmInfo.Eid}
    alarmInfo.EntityData.Leafs["reporting-agent-id"] = types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId}
    alarmInfo.EntityData.Leafs["pending-sync"] = types.YLeaf{"PendingSync", alarmInfo.PendingSync}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", alarmInfo.Status}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["service-affecting"] = types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting}
    alarmInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", alarmInfo.Type_}
    alarmInfo.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", alarmInfo.Interface_}
    alarmInfo.EntityData.Leafs["alarm-name"] = types.YLeaf{"AlarmName", alarmInfo.AlarmName}
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
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["direction"] = types.YLeaf{"Direction", otn.Direction}
    otn.EntityData.Leafs["notification-source"] = types.YLeaf{"NotificationSource", otn.NotificationSource}
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
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = make(map[string]types.YChild)
    tca.EntityData.Leafs = make(map[string]types.YLeaf)
    tca.EntityData.Leafs["threshold-value"] = types.YLeaf{"ThresholdValue", tca.ThresholdValue}
    tca.EntityData.Leafs["current-value"] = types.YLeaf{"CurrentValue", tca.CurrentValue}
    tca.EntityData.Leafs["bucket-type"] = types.YLeaf{"BucketType", tca.BucketType}
    return &(tca.EntityData)
}

// Alarms_Detail_DetailSystem_History
// Show the history alarms at this scope.
type Alarms_Detail_DetailSystem_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_History_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailSystem_History_AlarmInfo
}

func (history *Alarms_Detail_DetailSystem_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "detail-system"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range history.AlarmInfo {
        history.EntityData.Children[types.GetSegmentPath(&history.AlarmInfo[i])] = types.YChild{"AlarmInfo", &history.AlarmInfo[i]}
    }
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(history.EntityData)
}

// Alarms_Detail_DetailSystem_History_AlarmInfo
// Alarm List
type Alarms_Detail_DetailSystem_History_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Type_ interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface_ interface{}

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
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Children["otn"] = types.YChild{"Otn", &alarmInfo.Otn}
    alarmInfo.EntityData.Children["tca"] = types.YChild{"Tca", &alarmInfo.Tca}
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["aid"] = types.YLeaf{"Aid", alarmInfo.Aid}
    alarmInfo.EntityData.Leafs["tag"] = types.YLeaf{"Tag", alarmInfo.Tag}
    alarmInfo.EntityData.Leafs["module"] = types.YLeaf{"Module", alarmInfo.Module}
    alarmInfo.EntityData.Leafs["eid"] = types.YLeaf{"Eid", alarmInfo.Eid}
    alarmInfo.EntityData.Leafs["reporting-agent-id"] = types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId}
    alarmInfo.EntityData.Leafs["pending-sync"] = types.YLeaf{"PendingSync", alarmInfo.PendingSync}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", alarmInfo.Status}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["service-affecting"] = types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting}
    alarmInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", alarmInfo.Type_}
    alarmInfo.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", alarmInfo.Interface_}
    alarmInfo.EntityData.Leafs["alarm-name"] = types.YLeaf{"AlarmName", alarmInfo.AlarmName}
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
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["direction"] = types.YLeaf{"Direction", otn.Direction}
    otn.EntityData.Leafs["notification-source"] = types.YLeaf{"NotificationSource", otn.NotificationSource}
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
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = make(map[string]types.YChild)
    tca.EntityData.Leafs = make(map[string]types.YLeaf)
    tca.EntityData.Leafs["threshold-value"] = types.YLeaf{"ThresholdValue", tca.ThresholdValue}
    tca.EntityData.Leafs["current-value"] = types.YLeaf{"CurrentValue", tca.CurrentValue}
    tca.EntityData.Leafs["bucket-type"] = types.YLeaf{"BucketType", tca.BucketType}
    return &(tca.EntityData)
}

// Alarms_Detail_DetailSystem_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Detail_DetailSystem_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "detail-system"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = make(map[string]types.YChild)
    suppressed.EntityData.Children["suppressed-info"] = types.YChild{"SuppressedInfo", nil}
    for i := range suppressed.SuppressedInfo {
        suppressed.EntityData.Children[types.GetSegmentPath(&suppressed.SuppressedInfo[i])] = types.YChild{"SuppressedInfo", &suppressed.SuppressedInfo[i]}
    }
    suppressed.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(suppressed.EntityData)
}

// Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Interface_ interface{}

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
    suppressedInfo.EntityData.SegmentPath = "suppressed-info"
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = make(map[string]types.YChild)
    suppressedInfo.EntityData.Children["otn"] = types.YChild{"Otn", &suppressedInfo.Otn}
    suppressedInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    suppressedInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", suppressedInfo.Description}
    suppressedInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", suppressedInfo.Location}
    suppressedInfo.EntityData.Leafs["aid"] = types.YLeaf{"Aid", suppressedInfo.Aid}
    suppressedInfo.EntityData.Leafs["tag"] = types.YLeaf{"Tag", suppressedInfo.Tag}
    suppressedInfo.EntityData.Leafs["module"] = types.YLeaf{"Module", suppressedInfo.Module}
    suppressedInfo.EntityData.Leafs["eid"] = types.YLeaf{"Eid", suppressedInfo.Eid}
    suppressedInfo.EntityData.Leafs["reporting-agent-id"] = types.YLeaf{"ReportingAgentId", suppressedInfo.ReportingAgentId}
    suppressedInfo.EntityData.Leafs["pending-sync"] = types.YLeaf{"PendingSync", suppressedInfo.PendingSync}
    suppressedInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", suppressedInfo.Severity}
    suppressedInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", suppressedInfo.Status}
    suppressedInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", suppressedInfo.Group}
    suppressedInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", suppressedInfo.SetTime}
    suppressedInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp}
    suppressedInfo.EntityData.Leafs["suppressed-time"] = types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime}
    suppressedInfo.EntityData.Leafs["suppressed-timestamp"] = types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp}
    suppressedInfo.EntityData.Leafs["service-affecting"] = types.YLeaf{"ServiceAffecting", suppressedInfo.ServiceAffecting}
    suppressedInfo.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", suppressedInfo.Interface_}
    suppressedInfo.EntityData.Leafs["alarm-name"] = types.YLeaf{"AlarmName", suppressedInfo.AlarmName}
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
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["direction"] = types.YLeaf{"Direction", otn.Direction}
    otn.EntityData.Leafs["notification-source"] = types.YLeaf{"NotificationSource", otn.NotificationSource}
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
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["reported"] = types.YLeaf{"Reported", stats.Reported}
    stats.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", stats.Dropped}
    stats.EntityData.Leafs["active"] = types.YLeaf{"Active", stats.Active}
    stats.EntityData.Leafs["history"] = types.YLeaf{"History", stats.History}
    stats.EntityData.Leafs["suppressed"] = types.YLeaf{"Suppressed", stats.Suppressed}
    stats.EntityData.Leafs["sysadmin-active"] = types.YLeaf{"SysadminActive", stats.SysadminActive}
    stats.EntityData.Leafs["sysadmin-history"] = types.YLeaf{"SysadminHistory", stats.SysadminHistory}
    stats.EntityData.Leafs["sysadmin-suppressed"] = types.YLeaf{"SysadminSuppressed", stats.SysadminSuppressed}
    stats.EntityData.Leafs["dropped-invalid-aid"] = types.YLeaf{"DroppedInvalidAid", stats.DroppedInvalidAid}
    stats.EntityData.Leafs["dropped-insuff-mem"] = types.YLeaf{"DroppedInsuffMem", stats.DroppedInsuffMem}
    stats.EntityData.Leafs["dropped-db-error"] = types.YLeaf{"DroppedDbError", stats.DroppedDbError}
    stats.EntityData.Leafs["dropped-clear-without-set"] = types.YLeaf{"DroppedClearWithoutSet", stats.DroppedClearWithoutSet}
    stats.EntityData.Leafs["dropped-duplicate"] = types.YLeaf{"DroppedDuplicate", stats.DroppedDuplicate}
    stats.EntityData.Leafs["cache-hit"] = types.YLeaf{"CacheHit", stats.CacheHit}
    stats.EntityData.Leafs["cache-miss"] = types.YLeaf{"CacheMiss", stats.CacheMiss}
    return &(stats.EntityData)
}

// Alarms_Detail_DetailSystem_Clients
// Show the clients associated with this service.
type Alarms_Detail_DetailSystem_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client List. The type is slice of
    // Alarms_Detail_DetailSystem_Clients_ClientInfo.
    ClientInfo []Alarms_Detail_DetailSystem_Clients_ClientInfo
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "detail-system"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client-info"] = types.YChild{"ClientInfo", nil}
    for i := range clients.ClientInfo {
        clients.EntityData.Children[types.GetSegmentPath(&clients.ClientInfo[i])] = types.YChild{"ClientInfo", &clients.ClientInfo[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clients.EntityData)
}

// Alarms_Detail_DetailSystem_Clients_ClientInfo
// Client List
type Alarms_Detail_DetailSystem_Clients_ClientInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Type_ interface{}

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
    clientInfo.EntityData.SegmentPath = "client-info"
    clientInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientInfo.EntityData.Children = make(map[string]types.YChild)
    clientInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    clientInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", clientInfo.Name}
    clientInfo.EntityData.Leafs["id"] = types.YLeaf{"Id", clientInfo.Id}
    clientInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", clientInfo.Location}
    clientInfo.EntityData.Leafs["handle"] = types.YLeaf{"Handle", clientInfo.Handle}
    clientInfo.EntityData.Leafs["state"] = types.YLeaf{"State", clientInfo.State}
    clientInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", clientInfo.Type_}
    clientInfo.EntityData.Leafs["filter-disp"] = types.YLeaf{"FilterDisp", clientInfo.FilterDisp}
    clientInfo.EntityData.Leafs["subscriber-id"] = types.YLeaf{"SubscriberId", clientInfo.SubscriberId}
    clientInfo.EntityData.Leafs["filter-severity"] = types.YLeaf{"FilterSeverity", clientInfo.FilterSeverity}
    clientInfo.EntityData.Leafs["filter-state"] = types.YLeaf{"FilterState", clientInfo.FilterState}
    clientInfo.EntityData.Leafs["filter-group"] = types.YLeaf{"FilterGroup", clientInfo.FilterGroup}
    clientInfo.EntityData.Leafs["connect-count"] = types.YLeaf{"ConnectCount", clientInfo.ConnectCount}
    clientInfo.EntityData.Leafs["connect-timestamp"] = types.YLeaf{"ConnectTimestamp", clientInfo.ConnectTimestamp}
    clientInfo.EntityData.Leafs["get-count"] = types.YLeaf{"GetCount", clientInfo.GetCount}
    clientInfo.EntityData.Leafs["subscribe-count"] = types.YLeaf{"SubscribeCount", clientInfo.SubscribeCount}
    clientInfo.EntityData.Leafs["report-count"] = types.YLeaf{"ReportCount", clientInfo.ReportCount}
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
    detailCard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailCard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailCard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailCard.EntityData.Children = make(map[string]types.YChild)
    detailCard.EntityData.Children["detail-locations"] = types.YChild{"DetailLocations", &detailCard.DetailLocations}
    detailCard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detailCard.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations
// Table of DetailLocation
type Alarms_Detail_DetailCard_DetailLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a card location for alarms. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation.
    DetailLocation []Alarms_Detail_DetailCard_DetailLocations_DetailLocation
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetEntityData() *types.CommonEntityData {
    detailLocations.EntityData.YFilter = detailLocations.YFilter
    detailLocations.EntityData.YangName = "detail-locations"
    detailLocations.EntityData.BundleName = "cisco_ios_xr"
    detailLocations.EntityData.ParentYangName = "detail-card"
    detailLocations.EntityData.SegmentPath = "detail-locations"
    detailLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailLocations.EntityData.Children = make(map[string]types.YChild)
    detailLocations.EntityData.Children["detail-location"] = types.YChild{"DetailLocation", nil}
    for i := range detailLocations.DetailLocation {
        detailLocations.EntityData.Children[types.GetSegmentPath(&detailLocations.DetailLocation[i])] = types.YChild{"DetailLocation", &detailLocations.DetailLocation[i]}
    }
    detailLocations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detailLocations.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation
// Specify a card location for alarms.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NodeID of the Location. The type is string with
    // pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

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
    detailLocation.EntityData.SegmentPath = "detail-location" + "[node-id='" + fmt.Sprintf("%v", detailLocation.NodeId) + "']"
    detailLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailLocation.EntityData.Children = make(map[string]types.YChild)
    detailLocation.EntityData.Children["active"] = types.YChild{"Active", &detailLocation.Active}
    detailLocation.EntityData.Children["history"] = types.YChild{"History", &detailLocation.History}
    detailLocation.EntityData.Children["suppressed"] = types.YChild{"Suppressed", &detailLocation.Suppressed}
    detailLocation.EntityData.Children["stats"] = types.YChild{"Stats", &detailLocation.Stats}
    detailLocation.EntityData.Children["clients"] = types.YChild{"Clients", &detailLocation.Clients}
    detailLocation.EntityData.Leafs = make(map[string]types.YLeaf)
    detailLocation.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", detailLocation.NodeId}
    return &(detailLocation.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active
// Show the active alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "detail-location"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range active.AlarmInfo {
        active.EntityData.Children[types.GetSegmentPath(&active.AlarmInfo[i])] = types.YChild{"AlarmInfo", &active.AlarmInfo[i]}
    }
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(active.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo
// Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Type_ interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface_ interface{}

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
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Children["otn"] = types.YChild{"Otn", &alarmInfo.Otn}
    alarmInfo.EntityData.Children["tca"] = types.YChild{"Tca", &alarmInfo.Tca}
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["aid"] = types.YLeaf{"Aid", alarmInfo.Aid}
    alarmInfo.EntityData.Leafs["tag"] = types.YLeaf{"Tag", alarmInfo.Tag}
    alarmInfo.EntityData.Leafs["module"] = types.YLeaf{"Module", alarmInfo.Module}
    alarmInfo.EntityData.Leafs["eid"] = types.YLeaf{"Eid", alarmInfo.Eid}
    alarmInfo.EntityData.Leafs["reporting-agent-id"] = types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId}
    alarmInfo.EntityData.Leafs["pending-sync"] = types.YLeaf{"PendingSync", alarmInfo.PendingSync}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", alarmInfo.Status}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["service-affecting"] = types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting}
    alarmInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", alarmInfo.Type_}
    alarmInfo.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", alarmInfo.Interface_}
    alarmInfo.EntityData.Leafs["alarm-name"] = types.YLeaf{"AlarmName", alarmInfo.AlarmName}
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
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["direction"] = types.YLeaf{"Direction", otn.Direction}
    otn.EntityData.Leafs["notification-source"] = types.YLeaf{"NotificationSource", otn.NotificationSource}
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
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = make(map[string]types.YChild)
    tca.EntityData.Leafs = make(map[string]types.YLeaf)
    tca.EntityData.Leafs["threshold-value"] = types.YLeaf{"ThresholdValue", tca.ThresholdValue}
    tca.EntityData.Leafs["current-value"] = types.YLeaf{"CurrentValue", tca.CurrentValue}
    tca.EntityData.Leafs["bucket-type"] = types.YLeaf{"BucketType", tca.BucketType}
    return &(tca.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History
// Show the history alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "detail-location"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range history.AlarmInfo {
        history.EntityData.Children[types.GetSegmentPath(&history.AlarmInfo[i])] = types.YChild{"AlarmInfo", &history.AlarmInfo[i]}
    }
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(history.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo
// Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Type_ interface{}

    // Alarm interface name. The type is string with length: 0..128.
    Interface_ interface{}

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
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Children["otn"] = types.YChild{"Otn", &alarmInfo.Otn}
    alarmInfo.EntityData.Children["tca"] = types.YChild{"Tca", &alarmInfo.Tca}
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["aid"] = types.YLeaf{"Aid", alarmInfo.Aid}
    alarmInfo.EntityData.Leafs["tag"] = types.YLeaf{"Tag", alarmInfo.Tag}
    alarmInfo.EntityData.Leafs["module"] = types.YLeaf{"Module", alarmInfo.Module}
    alarmInfo.EntityData.Leafs["eid"] = types.YLeaf{"Eid", alarmInfo.Eid}
    alarmInfo.EntityData.Leafs["reporting-agent-id"] = types.YLeaf{"ReportingAgentId", alarmInfo.ReportingAgentId}
    alarmInfo.EntityData.Leafs["pending-sync"] = types.YLeaf{"PendingSync", alarmInfo.PendingSync}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", alarmInfo.Status}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["service-affecting"] = types.YLeaf{"ServiceAffecting", alarmInfo.ServiceAffecting}
    alarmInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", alarmInfo.Type_}
    alarmInfo.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", alarmInfo.Interface_}
    alarmInfo.EntityData.Leafs["alarm-name"] = types.YLeaf{"AlarmName", alarmInfo.AlarmName}
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
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["direction"] = types.YLeaf{"Direction", otn.Direction}
    otn.EntityData.Leafs["notification-source"] = types.YLeaf{"NotificationSource", otn.NotificationSource}
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
    tca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tca.EntityData.Children = make(map[string]types.YChild)
    tca.EntityData.Leafs = make(map[string]types.YLeaf)
    tca.EntityData.Leafs["threshold-value"] = types.YLeaf{"ThresholdValue", tca.ThresholdValue}
    tca.EntityData.Leafs["current-value"] = types.YLeaf{"CurrentValue", tca.CurrentValue}
    tca.EntityData.Leafs["bucket-type"] = types.YLeaf{"BucketType", tca.BucketType}
    return &(tca.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "detail-location"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = make(map[string]types.YChild)
    suppressed.EntityData.Children["suppressed-info"] = types.YChild{"SuppressedInfo", nil}
    for i := range suppressed.SuppressedInfo {
        suppressed.EntityData.Children[types.GetSegmentPath(&suppressed.SuppressedInfo[i])] = types.YChild{"SuppressedInfo", &suppressed.SuppressedInfo[i]}
    }
    suppressed.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(suppressed.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Interface_ interface{}

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
    suppressedInfo.EntityData.SegmentPath = "suppressed-info"
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = make(map[string]types.YChild)
    suppressedInfo.EntityData.Children["otn"] = types.YChild{"Otn", &suppressedInfo.Otn}
    suppressedInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    suppressedInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", suppressedInfo.Description}
    suppressedInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", suppressedInfo.Location}
    suppressedInfo.EntityData.Leafs["aid"] = types.YLeaf{"Aid", suppressedInfo.Aid}
    suppressedInfo.EntityData.Leafs["tag"] = types.YLeaf{"Tag", suppressedInfo.Tag}
    suppressedInfo.EntityData.Leafs["module"] = types.YLeaf{"Module", suppressedInfo.Module}
    suppressedInfo.EntityData.Leafs["eid"] = types.YLeaf{"Eid", suppressedInfo.Eid}
    suppressedInfo.EntityData.Leafs["reporting-agent-id"] = types.YLeaf{"ReportingAgentId", suppressedInfo.ReportingAgentId}
    suppressedInfo.EntityData.Leafs["pending-sync"] = types.YLeaf{"PendingSync", suppressedInfo.PendingSync}
    suppressedInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", suppressedInfo.Severity}
    suppressedInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", suppressedInfo.Status}
    suppressedInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", suppressedInfo.Group}
    suppressedInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", suppressedInfo.SetTime}
    suppressedInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp}
    suppressedInfo.EntityData.Leafs["suppressed-time"] = types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime}
    suppressedInfo.EntityData.Leafs["suppressed-timestamp"] = types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp}
    suppressedInfo.EntityData.Leafs["service-affecting"] = types.YLeaf{"ServiceAffecting", suppressedInfo.ServiceAffecting}
    suppressedInfo.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", suppressedInfo.Interface_}
    suppressedInfo.EntityData.Leafs["alarm-name"] = types.YLeaf{"AlarmName", suppressedInfo.AlarmName}
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
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["direction"] = types.YLeaf{"Direction", otn.Direction}
    otn.EntityData.Leafs["notification-source"] = types.YLeaf{"NotificationSource", otn.NotificationSource}
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
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["reported"] = types.YLeaf{"Reported", stats.Reported}
    stats.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", stats.Dropped}
    stats.EntityData.Leafs["active"] = types.YLeaf{"Active", stats.Active}
    stats.EntityData.Leafs["history"] = types.YLeaf{"History", stats.History}
    stats.EntityData.Leafs["suppressed"] = types.YLeaf{"Suppressed", stats.Suppressed}
    stats.EntityData.Leafs["sysadmin-active"] = types.YLeaf{"SysadminActive", stats.SysadminActive}
    stats.EntityData.Leafs["sysadmin-history"] = types.YLeaf{"SysadminHistory", stats.SysadminHistory}
    stats.EntityData.Leafs["sysadmin-suppressed"] = types.YLeaf{"SysadminSuppressed", stats.SysadminSuppressed}
    stats.EntityData.Leafs["dropped-invalid-aid"] = types.YLeaf{"DroppedInvalidAid", stats.DroppedInvalidAid}
    stats.EntityData.Leafs["dropped-insuff-mem"] = types.YLeaf{"DroppedInsuffMem", stats.DroppedInsuffMem}
    stats.EntityData.Leafs["dropped-db-error"] = types.YLeaf{"DroppedDbError", stats.DroppedDbError}
    stats.EntityData.Leafs["dropped-clear-without-set"] = types.YLeaf{"DroppedClearWithoutSet", stats.DroppedClearWithoutSet}
    stats.EntityData.Leafs["dropped-duplicate"] = types.YLeaf{"DroppedDuplicate", stats.DroppedDuplicate}
    stats.EntityData.Leafs["cache-hit"] = types.YLeaf{"CacheHit", stats.CacheHit}
    stats.EntityData.Leafs["cache-miss"] = types.YLeaf{"CacheMiss", stats.CacheMiss}
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
    ClientInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "detail-location"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client-info"] = types.YChild{"ClientInfo", nil}
    for i := range clients.ClientInfo {
        clients.EntityData.Children[types.GetSegmentPath(&clients.ClientInfo[i])] = types.YChild{"ClientInfo", &clients.ClientInfo[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clients.EntityData)
}

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo
// Client List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Type_ interface{}

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
    clientInfo.EntityData.SegmentPath = "client-info"
    clientInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientInfo.EntityData.Children = make(map[string]types.YChild)
    clientInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    clientInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", clientInfo.Name}
    clientInfo.EntityData.Leafs["id"] = types.YLeaf{"Id", clientInfo.Id}
    clientInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", clientInfo.Location}
    clientInfo.EntityData.Leafs["handle"] = types.YLeaf{"Handle", clientInfo.Handle}
    clientInfo.EntityData.Leafs["state"] = types.YLeaf{"State", clientInfo.State}
    clientInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", clientInfo.Type_}
    clientInfo.EntityData.Leafs["filter-disp"] = types.YLeaf{"FilterDisp", clientInfo.FilterDisp}
    clientInfo.EntityData.Leafs["subscriber-id"] = types.YLeaf{"SubscriberId", clientInfo.SubscriberId}
    clientInfo.EntityData.Leafs["filter-severity"] = types.YLeaf{"FilterSeverity", clientInfo.FilterSeverity}
    clientInfo.EntityData.Leafs["filter-state"] = types.YLeaf{"FilterState", clientInfo.FilterState}
    clientInfo.EntityData.Leafs["filter-group"] = types.YLeaf{"FilterGroup", clientInfo.FilterGroup}
    clientInfo.EntityData.Leafs["connect-count"] = types.YLeaf{"ConnectCount", clientInfo.ConnectCount}
    clientInfo.EntityData.Leafs["connect-timestamp"] = types.YLeaf{"ConnectTimestamp", clientInfo.ConnectTimestamp}
    clientInfo.EntityData.Leafs["get-count"] = types.YLeaf{"GetCount", clientInfo.GetCount}
    clientInfo.EntityData.Leafs["subscribe-count"] = types.YLeaf{"SubscribeCount", clientInfo.SubscribeCount}
    clientInfo.EntityData.Leafs["report-count"] = types.YLeaf{"ReportCount", clientInfo.ReportCount}
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
}

func (brief *Alarms_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "alarms"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Children["brief-card"] = types.YChild{"BriefCard", &brief.BriefCard}
    brief.EntityData.Children["brief-system"] = types.YChild{"BriefSystem", &brief.BriefSystem}
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
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
    briefCard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefCard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefCard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefCard.EntityData.Children = make(map[string]types.YChild)
    briefCard.EntityData.Children["brief-locations"] = types.YChild{"BriefLocations", &briefCard.BriefLocations}
    briefCard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(briefCard.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations
// Table of BriefLocation
type Alarms_Brief_BriefCard_BriefLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a card location for alarms. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation.
    BriefLocation []Alarms_Brief_BriefCard_BriefLocations_BriefLocation
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetEntityData() *types.CommonEntityData {
    briefLocations.EntityData.YFilter = briefLocations.YFilter
    briefLocations.EntityData.YangName = "brief-locations"
    briefLocations.EntityData.BundleName = "cisco_ios_xr"
    briefLocations.EntityData.ParentYangName = "brief-card"
    briefLocations.EntityData.SegmentPath = "brief-locations"
    briefLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLocations.EntityData.Children = make(map[string]types.YChild)
    briefLocations.EntityData.Children["brief-location"] = types.YChild{"BriefLocation", nil}
    for i := range briefLocations.BriefLocation {
        briefLocations.EntityData.Children[types.GetSegmentPath(&briefLocations.BriefLocation[i])] = types.YChild{"BriefLocation", &briefLocations.BriefLocation[i]}
    }
    briefLocations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(briefLocations.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation
// Specify a card location for alarms.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NodeID of the Location. The type is string with
    // pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

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
    briefLocation.EntityData.SegmentPath = "brief-location" + "[node-id='" + fmt.Sprintf("%v", briefLocation.NodeId) + "']"
    briefLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLocation.EntityData.Children = make(map[string]types.YChild)
    briefLocation.EntityData.Children["active"] = types.YChild{"Active", &briefLocation.Active}
    briefLocation.EntityData.Children["history"] = types.YChild{"History", &briefLocation.History}
    briefLocation.EntityData.Children["suppressed"] = types.YChild{"Suppressed", &briefLocation.Suppressed}
    briefLocation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefLocation.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", briefLocation.NodeId}
    return &(briefLocation.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active
// Show the active alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "brief-location"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range active.AlarmInfo {
        active.EntityData.Children[types.GetSegmentPath(&active.AlarmInfo[i])] = types.YChild{"AlarmInfo", &active.AlarmInfo[i]}
    }
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(active.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo
// Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo struct {
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

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "active"
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History
// Show the history alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "brief-location"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range history.AlarmInfo {
        history.EntityData.Children[types.GetSegmentPath(&history.AlarmInfo[i])] = types.YChild{"AlarmInfo", &history.AlarmInfo[i]}
    }
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(history.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo
// Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo struct {
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

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "history"
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "brief-location"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = make(map[string]types.YChild)
    suppressed.EntityData.Children["suppressed-info"] = types.YChild{"SuppressedInfo", nil}
    for i := range suppressed.SuppressedInfo {
        suppressed.EntityData.Children[types.GetSegmentPath(&suppressed.SuppressedInfo[i])] = types.YChild{"SuppressedInfo", &suppressed.SuppressedInfo[i]}
    }
    suppressed.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(suppressed.EntityData)
}

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo struct {
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
    suppressedInfo.EntityData.SegmentPath = "suppressed-info"
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = make(map[string]types.YChild)
    suppressedInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    suppressedInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", suppressedInfo.Location}
    suppressedInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", suppressedInfo.Severity}
    suppressedInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", suppressedInfo.Group}
    suppressedInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", suppressedInfo.SetTime}
    suppressedInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp}
    suppressedInfo.EntityData.Leafs["suppressed-time"] = types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime}
    suppressedInfo.EntityData.Leafs["suppressed-timestamp"] = types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp}
    suppressedInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", suppressedInfo.Description}
    return &(suppressedInfo.EntityData)
}

// Alarms_Brief_BriefSystem
// Show brief system scope alarm related data.
type Alarms_Brief_BriefSystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    briefSystem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefSystem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefSystem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefSystem.EntityData.Children = make(map[string]types.YChild)
    briefSystem.EntityData.Children["active"] = types.YChild{"Active", &briefSystem.Active}
    briefSystem.EntityData.Children["history"] = types.YChild{"History", &briefSystem.History}
    briefSystem.EntityData.Children["suppressed"] = types.YChild{"Suppressed", &briefSystem.Suppressed}
    briefSystem.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(briefSystem.EntityData)
}

// Alarms_Brief_BriefSystem_Active
// Show the active alarms at this scope.
type Alarms_Brief_BriefSystem_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of Alarms_Brief_BriefSystem_Active_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefSystem_Active_AlarmInfo
}

func (active *Alarms_Brief_BriefSystem_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "brief-system"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range active.AlarmInfo {
        active.EntityData.Children[types.GetSegmentPath(&active.AlarmInfo[i])] = types.YChild{"AlarmInfo", &active.AlarmInfo[i]}
    }
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(active.EntityData)
}

// Alarms_Brief_BriefSystem_Active_AlarmInfo
// Alarm List
type Alarms_Brief_BriefSystem_Active_AlarmInfo struct {
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

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "active"
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefSystem_History
// Show the history alarms at this scope.
type Alarms_Brief_BriefSystem_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefSystem_History_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefSystem_History_AlarmInfo
}

func (history *Alarms_Brief_BriefSystem_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "brief-system"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["alarm-info"] = types.YChild{"AlarmInfo", nil}
    for i := range history.AlarmInfo {
        history.EntityData.Children[types.GetSegmentPath(&history.AlarmInfo[i])] = types.YChild{"AlarmInfo", &history.AlarmInfo[i]}
    }
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(history.EntityData)
}

// Alarms_Brief_BriefSystem_History_AlarmInfo
// Alarm List
type Alarms_Brief_BriefSystem_History_AlarmInfo struct {
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

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "history"
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = make(map[string]types.YChild)
    alarmInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    alarmInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", alarmInfo.Location}
    alarmInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarmInfo.Severity}
    alarmInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", alarmInfo.Group}
    alarmInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", alarmInfo.SetTime}
    alarmInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", alarmInfo.SetTimestamp}
    alarmInfo.EntityData.Leafs["clear-time"] = types.YLeaf{"ClearTime", alarmInfo.ClearTime}
    alarmInfo.EntityData.Leafs["clear-timestamp"] = types.YLeaf{"ClearTimestamp", alarmInfo.ClearTimestamp}
    alarmInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", alarmInfo.Description}
    return &(alarmInfo.EntityData)
}

// Alarms_Brief_BriefSystem_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Brief_BriefSystem_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "brief-system"
    suppressed.EntityData.SegmentPath = "suppressed"
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = make(map[string]types.YChild)
    suppressed.EntityData.Children["suppressed-info"] = types.YChild{"SuppressedInfo", nil}
    for i := range suppressed.SuppressedInfo {
        suppressed.EntityData.Children[types.GetSegmentPath(&suppressed.SuppressedInfo[i])] = types.YChild{"SuppressedInfo", &suppressed.SuppressedInfo[i]}
    }
    suppressed.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(suppressed.EntityData)
}

// Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo struct {
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
    suppressedInfo.EntityData.SegmentPath = "suppressed-info"
    suppressedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressedInfo.EntityData.Children = make(map[string]types.YChild)
    suppressedInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    suppressedInfo.EntityData.Leafs["location"] = types.YLeaf{"Location", suppressedInfo.Location}
    suppressedInfo.EntityData.Leafs["severity"] = types.YLeaf{"Severity", suppressedInfo.Severity}
    suppressedInfo.EntityData.Leafs["group"] = types.YLeaf{"Group", suppressedInfo.Group}
    suppressedInfo.EntityData.Leafs["set-time"] = types.YLeaf{"SetTime", suppressedInfo.SetTime}
    suppressedInfo.EntityData.Leafs["set-timestamp"] = types.YLeaf{"SetTimestamp", suppressedInfo.SetTimestamp}
    suppressedInfo.EntityData.Leafs["suppressed-time"] = types.YLeaf{"SuppressedTime", suppressedInfo.SuppressedTime}
    suppressedInfo.EntityData.Leafs["suppressed-timestamp"] = types.YLeaf{"SuppressedTimestamp", suppressedInfo.SuppressedTimestamp}
    suppressedInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", suppressedInfo.Description}
    return &(suppressedInfo.EntityData)
}

