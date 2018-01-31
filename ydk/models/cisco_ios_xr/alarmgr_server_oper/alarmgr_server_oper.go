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
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of detail alarm commands.
    Detail Alarms_Detail

    // A set of brief alarm commands.
    Brief Alarms_Brief
}

func (alarms *Alarms) GetFilter() yfilter.YFilter { return alarms.YFilter }

func (alarms *Alarms) SetFilter(yf yfilter.YFilter) { alarms.YFilter = yf }

func (alarms *Alarms) GetGoName(yname string) string {
    if yname == "detail" { return "Detail" }
    if yname == "brief" { return "Brief" }
    return ""
}

func (alarms *Alarms) GetSegmentPath() string {
    return "Cisco-IOS-XR-alarmgr-server-oper:alarms"
}

func (alarms *Alarms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &alarms.Detail
    }
    if childYangName == "brief" {
        return &alarms.Brief
    }
    return nil
}

func (alarms *Alarms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &alarms.Detail
    children["brief"] = &alarms.Brief
    return children
}

func (alarms *Alarms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alarms *Alarms) GetBundleName() string { return "cisco_ios_xr" }

func (alarms *Alarms) GetYangName() string { return "alarms" }

func (alarms *Alarms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarms *Alarms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarms *Alarms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarms *Alarms) SetParent(parent types.Entity) { alarms.parent = parent }

func (alarms *Alarms) GetParent() types.Entity { return alarms.parent }

func (alarms *Alarms) GetParentYangName() string { return "Cisco-IOS-XR-alarmgr-server-oper" }

// Alarms_Detail
// A set of detail alarm commands.
type Alarms_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // show detail system scope alarm related data.
    DetailSystem Alarms_Detail_DetailSystem

    // Show detail card scope alarm related data.
    DetailCard Alarms_Detail_DetailCard
}

func (detail *Alarms_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Alarms_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Alarms_Detail) GetGoName(yname string) string {
    if yname == "detail-system" { return "DetailSystem" }
    if yname == "detail-card" { return "DetailCard" }
    return ""
}

func (detail *Alarms_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Alarms_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-system" {
        return &detail.DetailSystem
    }
    if childYangName == "detail-card" {
        return &detail.DetailCard
    }
    return nil
}

func (detail *Alarms_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail-system"] = &detail.DetailSystem
    children["detail-card"] = &detail.DetailCard
    return children
}

func (detail *Alarms_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *Alarms_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Alarms_Detail) GetYangName() string { return "detail" }

func (detail *Alarms_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Alarms_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Alarms_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Alarms_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Alarms_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Alarms_Detail) GetParentYangName() string { return "alarms" }

// Alarms_Detail_DetailSystem
// show detail system scope alarm related data.
type Alarms_Detail_DetailSystem struct {
    parent types.Entity
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

func (detailSystem *Alarms_Detail_DetailSystem) GetFilter() yfilter.YFilter { return detailSystem.YFilter }

func (detailSystem *Alarms_Detail_DetailSystem) SetFilter(yf yfilter.YFilter) { detailSystem.YFilter = yf }

func (detailSystem *Alarms_Detail_DetailSystem) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "history" { return "History" }
    if yname == "suppressed" { return "Suppressed" }
    if yname == "stats" { return "Stats" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (detailSystem *Alarms_Detail_DetailSystem) GetSegmentPath() string {
    return "detail-system"
}

func (detailSystem *Alarms_Detail_DetailSystem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active" {
        return &detailSystem.Active
    }
    if childYangName == "history" {
        return &detailSystem.History
    }
    if childYangName == "suppressed" {
        return &detailSystem.Suppressed
    }
    if childYangName == "stats" {
        return &detailSystem.Stats
    }
    if childYangName == "clients" {
        return &detailSystem.Clients
    }
    return nil
}

func (detailSystem *Alarms_Detail_DetailSystem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["active"] = &detailSystem.Active
    children["history"] = &detailSystem.History
    children["suppressed"] = &detailSystem.Suppressed
    children["stats"] = &detailSystem.Stats
    children["clients"] = &detailSystem.Clients
    return children
}

func (detailSystem *Alarms_Detail_DetailSystem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailSystem *Alarms_Detail_DetailSystem) GetBundleName() string { return "cisco_ios_xr" }

func (detailSystem *Alarms_Detail_DetailSystem) GetYangName() string { return "detail-system" }

func (detailSystem *Alarms_Detail_DetailSystem) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailSystem *Alarms_Detail_DetailSystem) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailSystem *Alarms_Detail_DetailSystem) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailSystem *Alarms_Detail_DetailSystem) SetParent(parent types.Entity) { detailSystem.parent = parent }

func (detailSystem *Alarms_Detail_DetailSystem) GetParent() types.Entity { return detailSystem.parent }

func (detailSystem *Alarms_Detail_DetailSystem) GetParentYangName() string { return "detail" }

// Alarms_Detail_DetailSystem_Active
// Show the active alarms at this scope.
type Alarms_Detail_DetailSystem_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_Active_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailSystem_Active_AlarmInfo
}

func (active *Alarms_Detail_DetailSystem_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *Alarms_Detail_DetailSystem_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *Alarms_Detail_DetailSystem_Active) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (active *Alarms_Detail_DetailSystem_Active) GetSegmentPath() string {
    return "active"
}

func (active *Alarms_Detail_DetailSystem_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range active.AlarmInfo {
            if active.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailSystem_Active_AlarmInfo{}
        active.AlarmInfo = append(active.AlarmInfo, child)
        return &active.AlarmInfo[len(active.AlarmInfo)-1]
    }
    return nil
}

func (active *Alarms_Detail_DetailSystem_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range active.AlarmInfo {
        children[active.AlarmInfo[i].GetSegmentPath()] = &active.AlarmInfo[i]
    }
    return children
}

func (active *Alarms_Detail_DetailSystem_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *Alarms_Detail_DetailSystem_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *Alarms_Detail_DetailSystem_Active) GetYangName() string { return "active" }

func (active *Alarms_Detail_DetailSystem_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *Alarms_Detail_DetailSystem_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *Alarms_Detail_DetailSystem_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *Alarms_Detail_DetailSystem_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *Alarms_Detail_DetailSystem_Active) GetParent() types.Entity { return active.parent }

func (active *Alarms_Detail_DetailSystem_Active) GetParentYangName() string { return "detail-system" }

// Alarms_Detail_DetailSystem_Active_AlarmInfo
// Alarm List
type Alarms_Detail_DetailSystem_Active_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "location" { return "Location" }
    if yname == "aid" { return "Aid" }
    if yname == "tag" { return "Tag" }
    if yname == "module" { return "Module" }
    if yname == "eid" { return "Eid" }
    if yname == "reporting-agent-id" { return "ReportingAgentId" }
    if yname == "pending-sync" { return "PendingSync" }
    if yname == "severity" { return "Severity" }
    if yname == "status" { return "Status" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "service-affecting" { return "ServiceAffecting" }
    if yname == "type" { return "Type" }
    if yname == "interface" { return "Interface" }
    if yname == "alarm-name" { return "AlarmName" }
    if yname == "otn" { return "Otn" }
    if yname == "tca" { return "Tca" }
    return ""
}

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn" {
        return &alarmInfo.Otn
    }
    if childYangName == "tca" {
        return &alarmInfo.Tca
    }
    return nil
}

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn"] = &alarmInfo.Otn
    children["tca"] = &alarmInfo.Tca
    return children
}

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = alarmInfo.Description
    leafs["location"] = alarmInfo.Location
    leafs["aid"] = alarmInfo.Aid
    leafs["tag"] = alarmInfo.Tag
    leafs["module"] = alarmInfo.Module
    leafs["eid"] = alarmInfo.Eid
    leafs["reporting-agent-id"] = alarmInfo.ReportingAgentId
    leafs["pending-sync"] = alarmInfo.PendingSync
    leafs["severity"] = alarmInfo.Severity
    leafs["status"] = alarmInfo.Status
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["service-affecting"] = alarmInfo.ServiceAffecting
    leafs["type"] = alarmInfo.Type
    leafs["interface"] = alarmInfo.Interface
    leafs["alarm-name"] = alarmInfo.AlarmName
    return leafs
}

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Detail_DetailSystem_Active_AlarmInfo) GetParentYangName() string { return "active" }

// Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "notification-source" { return "NotificationSource" }
    return ""
}

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = otn.Direction
    leafs["notification-source"] = otn.NotificationSource
    return leafs
}

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetYangName() string { return "otn" }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetParent() types.Entity { return otn.parent }

func (otn *Alarms_Detail_DetailSystem_Active_AlarmInfo_Otn) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetFilter() yfilter.YFilter { return tca.YFilter }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) SetFilter(yf yfilter.YFilter) { tca.YFilter = yf }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetGoName(yname string) string {
    if yname == "threshold-value" { return "ThresholdValue" }
    if yname == "current-value" { return "CurrentValue" }
    if yname == "bucket-type" { return "BucketType" }
    return ""
}

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetSegmentPath() string {
    return "tca"
}

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-value"] = tca.ThresholdValue
    leafs["current-value"] = tca.CurrentValue
    leafs["bucket-type"] = tca.BucketType
    return leafs
}

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetBundleName() string { return "cisco_ios_xr" }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetYangName() string { return "tca" }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) SetParent(parent types.Entity) { tca.parent = parent }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetParent() types.Entity { return tca.parent }

func (tca *Alarms_Detail_DetailSystem_Active_AlarmInfo_Tca) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailSystem_History
// Show the history alarms at this scope.
type Alarms_Detail_DetailSystem_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_History_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailSystem_History_AlarmInfo
}

func (history *Alarms_Detail_DetailSystem_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Alarms_Detail_DetailSystem_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Alarms_Detail_DetailSystem_History) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (history *Alarms_Detail_DetailSystem_History) GetSegmentPath() string {
    return "history"
}

func (history *Alarms_Detail_DetailSystem_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range history.AlarmInfo {
            if history.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailSystem_History_AlarmInfo{}
        history.AlarmInfo = append(history.AlarmInfo, child)
        return &history.AlarmInfo[len(history.AlarmInfo)-1]
    }
    return nil
}

func (history *Alarms_Detail_DetailSystem_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range history.AlarmInfo {
        children[history.AlarmInfo[i].GetSegmentPath()] = &history.AlarmInfo[i]
    }
    return children
}

func (history *Alarms_Detail_DetailSystem_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (history *Alarms_Detail_DetailSystem_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Alarms_Detail_DetailSystem_History) GetYangName() string { return "history" }

func (history *Alarms_Detail_DetailSystem_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Alarms_Detail_DetailSystem_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Alarms_Detail_DetailSystem_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Alarms_Detail_DetailSystem_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Alarms_Detail_DetailSystem_History) GetParent() types.Entity { return history.parent }

func (history *Alarms_Detail_DetailSystem_History) GetParentYangName() string { return "detail-system" }

// Alarms_Detail_DetailSystem_History_AlarmInfo
// Alarm List
type Alarms_Detail_DetailSystem_History_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "location" { return "Location" }
    if yname == "aid" { return "Aid" }
    if yname == "tag" { return "Tag" }
    if yname == "module" { return "Module" }
    if yname == "eid" { return "Eid" }
    if yname == "reporting-agent-id" { return "ReportingAgentId" }
    if yname == "pending-sync" { return "PendingSync" }
    if yname == "severity" { return "Severity" }
    if yname == "status" { return "Status" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "service-affecting" { return "ServiceAffecting" }
    if yname == "type" { return "Type" }
    if yname == "interface" { return "Interface" }
    if yname == "alarm-name" { return "AlarmName" }
    if yname == "otn" { return "Otn" }
    if yname == "tca" { return "Tca" }
    return ""
}

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn" {
        return &alarmInfo.Otn
    }
    if childYangName == "tca" {
        return &alarmInfo.Tca
    }
    return nil
}

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn"] = &alarmInfo.Otn
    children["tca"] = &alarmInfo.Tca
    return children
}

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = alarmInfo.Description
    leafs["location"] = alarmInfo.Location
    leafs["aid"] = alarmInfo.Aid
    leafs["tag"] = alarmInfo.Tag
    leafs["module"] = alarmInfo.Module
    leafs["eid"] = alarmInfo.Eid
    leafs["reporting-agent-id"] = alarmInfo.ReportingAgentId
    leafs["pending-sync"] = alarmInfo.PendingSync
    leafs["severity"] = alarmInfo.Severity
    leafs["status"] = alarmInfo.Status
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["service-affecting"] = alarmInfo.ServiceAffecting
    leafs["type"] = alarmInfo.Type
    leafs["interface"] = alarmInfo.Interface
    leafs["alarm-name"] = alarmInfo.AlarmName
    return leafs
}

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Detail_DetailSystem_History_AlarmInfo) GetParentYangName() string { return "history" }

// Alarms_Detail_DetailSystem_History_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailSystem_History_AlarmInfo_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "notification-source" { return "NotificationSource" }
    return ""
}

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = otn.Direction
    leafs["notification-source"] = otn.NotificationSource
    return leafs
}

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetYangName() string { return "otn" }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetParent() types.Entity { return otn.parent }

func (otn *Alarms_Detail_DetailSystem_History_AlarmInfo_Otn) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailSystem_History_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailSystem_History_AlarmInfo_Tca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetFilter() yfilter.YFilter { return tca.YFilter }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) SetFilter(yf yfilter.YFilter) { tca.YFilter = yf }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetGoName(yname string) string {
    if yname == "threshold-value" { return "ThresholdValue" }
    if yname == "current-value" { return "CurrentValue" }
    if yname == "bucket-type" { return "BucketType" }
    return ""
}

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetSegmentPath() string {
    return "tca"
}

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-value"] = tca.ThresholdValue
    leafs["current-value"] = tca.CurrentValue
    leafs["bucket-type"] = tca.BucketType
    return leafs
}

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetBundleName() string { return "cisco_ios_xr" }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetYangName() string { return "tca" }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) SetParent(parent types.Entity) { tca.parent = parent }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetParent() types.Entity { return tca.parent }

func (tca *Alarms_Detail_DetailSystem_History_AlarmInfo_Tca) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailSystem_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Detail_DetailSystem_Suppressed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetFilter() yfilter.YFilter { return suppressed.YFilter }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) SetFilter(yf yfilter.YFilter) { suppressed.YFilter = yf }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetGoName(yname string) string {
    if yname == "suppressed-info" { return "SuppressedInfo" }
    return ""
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetSegmentPath() string {
    return "suppressed"
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "suppressed-info" {
        for _, c := range suppressed.SuppressedInfo {
            if suppressed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo{}
        suppressed.SuppressedInfo = append(suppressed.SuppressedInfo, child)
        return &suppressed.SuppressedInfo[len(suppressed.SuppressedInfo)-1]
    }
    return nil
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range suppressed.SuppressedInfo {
        children[suppressed.SuppressedInfo[i].GetSegmentPath()] = &suppressed.SuppressedInfo[i]
    }
    return children
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetBundleName() string { return "cisco_ios_xr" }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetYangName() string { return "suppressed" }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) SetParent(parent types.Entity) { suppressed.parent = parent }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetParent() types.Entity { return suppressed.parent }

func (suppressed *Alarms_Detail_DetailSystem_Suppressed) GetParentYangName() string { return "detail-system" }

// Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo struct {
    parent types.Entity
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
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn
}

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetFilter() yfilter.YFilter { return suppressedInfo.YFilter }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) SetFilter(yf yfilter.YFilter) { suppressedInfo.YFilter = yf }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "location" { return "Location" }
    if yname == "aid" { return "Aid" }
    if yname == "tag" { return "Tag" }
    if yname == "module" { return "Module" }
    if yname == "eid" { return "Eid" }
    if yname == "reporting-agent-id" { return "ReportingAgentId" }
    if yname == "pending-sync" { return "PendingSync" }
    if yname == "severity" { return "Severity" }
    if yname == "status" { return "Status" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "suppressed-time" { return "SuppressedTime" }
    if yname == "suppressed-timestamp" { return "SuppressedTimestamp" }
    if yname == "service-affecting" { return "ServiceAffecting" }
    if yname == "interface" { return "Interface" }
    if yname == "alarm-name" { return "AlarmName" }
    if yname == "otn" { return "Otn" }
    return ""
}

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetSegmentPath() string {
    return "suppressed-info"
}

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn" {
        return &suppressedInfo.Otn
    }
    return nil
}

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn"] = &suppressedInfo.Otn
    return children
}

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = suppressedInfo.Description
    leafs["location"] = suppressedInfo.Location
    leafs["aid"] = suppressedInfo.Aid
    leafs["tag"] = suppressedInfo.Tag
    leafs["module"] = suppressedInfo.Module
    leafs["eid"] = suppressedInfo.Eid
    leafs["reporting-agent-id"] = suppressedInfo.ReportingAgentId
    leafs["pending-sync"] = suppressedInfo.PendingSync
    leafs["severity"] = suppressedInfo.Severity
    leafs["status"] = suppressedInfo.Status
    leafs["group"] = suppressedInfo.Group
    leafs["set-time"] = suppressedInfo.SetTime
    leafs["set-timestamp"] = suppressedInfo.SetTimestamp
    leafs["suppressed-time"] = suppressedInfo.SuppressedTime
    leafs["suppressed-timestamp"] = suppressedInfo.SuppressedTimestamp
    leafs["service-affecting"] = suppressedInfo.ServiceAffecting
    leafs["interface"] = suppressedInfo.Interface
    leafs["alarm-name"] = suppressedInfo.AlarmName
    return leafs
}

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetBundleName() string { return "cisco_ios_xr" }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetYangName() string { return "suppressed-info" }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) SetParent(parent types.Entity) { suppressedInfo.parent = parent }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetParent() types.Entity { return suppressedInfo.parent }

func (suppressedInfo *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo) GetParentYangName() string { return "suppressed" }

// Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "notification-source" { return "NotificationSource" }
    return ""
}

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = otn.Direction
    leafs["notification-source"] = otn.NotificationSource
    return leafs
}

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetYangName() string { return "otn" }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetParent() types.Entity { return otn.parent }

func (otn *Alarms_Detail_DetailSystem_Suppressed_SuppressedInfo_Otn) GetParentYangName() string { return "suppressed-info" }

// Alarms_Detail_DetailSystem_Stats
// Show the service statistics.
type Alarms_Detail_DetailSystem_Stats struct {
    parent types.Entity
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

func (stats *Alarms_Detail_DetailSystem_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *Alarms_Detail_DetailSystem_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *Alarms_Detail_DetailSystem_Stats) GetGoName(yname string) string {
    if yname == "reported" { return "Reported" }
    if yname == "dropped" { return "Dropped" }
    if yname == "active" { return "Active" }
    if yname == "history" { return "History" }
    if yname == "suppressed" { return "Suppressed" }
    if yname == "sysadmin-active" { return "SysadminActive" }
    if yname == "sysadmin-history" { return "SysadminHistory" }
    if yname == "sysadmin-suppressed" { return "SysadminSuppressed" }
    if yname == "dropped-invalid-aid" { return "DroppedInvalidAid" }
    if yname == "dropped-insuff-mem" { return "DroppedInsuffMem" }
    if yname == "dropped-db-error" { return "DroppedDbError" }
    if yname == "dropped-clear-without-set" { return "DroppedClearWithoutSet" }
    if yname == "dropped-duplicate" { return "DroppedDuplicate" }
    if yname == "cache-hit" { return "CacheHit" }
    if yname == "cache-miss" { return "CacheMiss" }
    return ""
}

func (stats *Alarms_Detail_DetailSystem_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *Alarms_Detail_DetailSystem_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *Alarms_Detail_DetailSystem_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *Alarms_Detail_DetailSystem_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reported"] = stats.Reported
    leafs["dropped"] = stats.Dropped
    leafs["active"] = stats.Active
    leafs["history"] = stats.History
    leafs["suppressed"] = stats.Suppressed
    leafs["sysadmin-active"] = stats.SysadminActive
    leafs["sysadmin-history"] = stats.SysadminHistory
    leafs["sysadmin-suppressed"] = stats.SysadminSuppressed
    leafs["dropped-invalid-aid"] = stats.DroppedInvalidAid
    leafs["dropped-insuff-mem"] = stats.DroppedInsuffMem
    leafs["dropped-db-error"] = stats.DroppedDbError
    leafs["dropped-clear-without-set"] = stats.DroppedClearWithoutSet
    leafs["dropped-duplicate"] = stats.DroppedDuplicate
    leafs["cache-hit"] = stats.CacheHit
    leafs["cache-miss"] = stats.CacheMiss
    return leafs
}

func (stats *Alarms_Detail_DetailSystem_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *Alarms_Detail_DetailSystem_Stats) GetYangName() string { return "stats" }

func (stats *Alarms_Detail_DetailSystem_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *Alarms_Detail_DetailSystem_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *Alarms_Detail_DetailSystem_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *Alarms_Detail_DetailSystem_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *Alarms_Detail_DetailSystem_Stats) GetParent() types.Entity { return stats.parent }

func (stats *Alarms_Detail_DetailSystem_Stats) GetParentYangName() string { return "detail-system" }

// Alarms_Detail_DetailSystem_Clients
// Show the clients associated with this service.
type Alarms_Detail_DetailSystem_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Client List. The type is slice of
    // Alarms_Detail_DetailSystem_Clients_ClientInfo.
    ClientInfo []Alarms_Detail_DetailSystem_Clients_ClientInfo
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Alarms_Detail_DetailSystem_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Alarms_Detail_DetailSystem_Clients) GetGoName(yname string) string {
    if yname == "client-info" { return "ClientInfo" }
    return ""
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-info" {
        for _, c := range clients.ClientInfo {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailSystem_Clients_ClientInfo{}
        clients.ClientInfo = append(clients.ClientInfo, child)
        return &clients.ClientInfo[len(clients.ClientInfo)-1]
    }
    return nil
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.ClientInfo {
        children[clients.ClientInfo[i].GetSegmentPath()] = &clients.ClientInfo[i]
    }
    return children
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Alarms_Detail_DetailSystem_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Alarms_Detail_DetailSystem_Clients) GetYangName() string { return "clients" }

func (clients *Alarms_Detail_DetailSystem_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Alarms_Detail_DetailSystem_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Alarms_Detail_DetailSystem_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Alarms_Detail_DetailSystem_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Alarms_Detail_DetailSystem_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Alarms_Detail_DetailSystem_Clients) GetParentYangName() string { return "detail-system" }

// Alarms_Detail_DetailSystem_Clients_ClientInfo
// Client List
type Alarms_Detail_DetailSystem_Clients_ClientInfo struct {
    parent types.Entity
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

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetFilter() yfilter.YFilter { return clientInfo.YFilter }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) SetFilter(yf yfilter.YFilter) { clientInfo.YFilter = yf }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "id" { return "Id" }
    if yname == "location" { return "Location" }
    if yname == "handle" { return "Handle" }
    if yname == "state" { return "State" }
    if yname == "type" { return "Type" }
    if yname == "filter-disp" { return "FilterDisp" }
    if yname == "subscriber-id" { return "SubscriberId" }
    if yname == "filter-severity" { return "FilterSeverity" }
    if yname == "filter-state" { return "FilterState" }
    if yname == "filter-group" { return "FilterGroup" }
    if yname == "connect-count" { return "ConnectCount" }
    if yname == "connect-timestamp" { return "ConnectTimestamp" }
    if yname == "get-count" { return "GetCount" }
    if yname == "subscribe-count" { return "SubscribeCount" }
    if yname == "report-count" { return "ReportCount" }
    return ""
}

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetSegmentPath() string {
    return "client-info"
}

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = clientInfo.Name
    leafs["id"] = clientInfo.Id
    leafs["location"] = clientInfo.Location
    leafs["handle"] = clientInfo.Handle
    leafs["state"] = clientInfo.State
    leafs["type"] = clientInfo.Type
    leafs["filter-disp"] = clientInfo.FilterDisp
    leafs["subscriber-id"] = clientInfo.SubscriberId
    leafs["filter-severity"] = clientInfo.FilterSeverity
    leafs["filter-state"] = clientInfo.FilterState
    leafs["filter-group"] = clientInfo.FilterGroup
    leafs["connect-count"] = clientInfo.ConnectCount
    leafs["connect-timestamp"] = clientInfo.ConnectTimestamp
    leafs["get-count"] = clientInfo.GetCount
    leafs["subscribe-count"] = clientInfo.SubscribeCount
    leafs["report-count"] = clientInfo.ReportCount
    return leafs
}

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetBundleName() string { return "cisco_ios_xr" }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetYangName() string { return "client-info" }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) SetParent(parent types.Entity) { clientInfo.parent = parent }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetParent() types.Entity { return clientInfo.parent }

func (clientInfo *Alarms_Detail_DetailSystem_Clients_ClientInfo) GetParentYangName() string { return "clients" }

// Alarms_Detail_DetailCard
// Show detail card scope alarm related data.
type Alarms_Detail_DetailCard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of DetailLocation.
    DetailLocations Alarms_Detail_DetailCard_DetailLocations
}

func (detailCard *Alarms_Detail_DetailCard) GetFilter() yfilter.YFilter { return detailCard.YFilter }

func (detailCard *Alarms_Detail_DetailCard) SetFilter(yf yfilter.YFilter) { detailCard.YFilter = yf }

func (detailCard *Alarms_Detail_DetailCard) GetGoName(yname string) string {
    if yname == "detail-locations" { return "DetailLocations" }
    return ""
}

func (detailCard *Alarms_Detail_DetailCard) GetSegmentPath() string {
    return "detail-card"
}

func (detailCard *Alarms_Detail_DetailCard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-locations" {
        return &detailCard.DetailLocations
    }
    return nil
}

func (detailCard *Alarms_Detail_DetailCard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail-locations"] = &detailCard.DetailLocations
    return children
}

func (detailCard *Alarms_Detail_DetailCard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailCard *Alarms_Detail_DetailCard) GetBundleName() string { return "cisco_ios_xr" }

func (detailCard *Alarms_Detail_DetailCard) GetYangName() string { return "detail-card" }

func (detailCard *Alarms_Detail_DetailCard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailCard *Alarms_Detail_DetailCard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailCard *Alarms_Detail_DetailCard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailCard *Alarms_Detail_DetailCard) SetParent(parent types.Entity) { detailCard.parent = parent }

func (detailCard *Alarms_Detail_DetailCard) GetParent() types.Entity { return detailCard.parent }

func (detailCard *Alarms_Detail_DetailCard) GetParentYangName() string { return "detail" }

// Alarms_Detail_DetailCard_DetailLocations
// Table of DetailLocation
type Alarms_Detail_DetailCard_DetailLocations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify a card location for alarms. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation.
    DetailLocation []Alarms_Detail_DetailCard_DetailLocations_DetailLocation
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetFilter() yfilter.YFilter { return detailLocations.YFilter }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) SetFilter(yf yfilter.YFilter) { detailLocations.YFilter = yf }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetGoName(yname string) string {
    if yname == "detail-location" { return "DetailLocation" }
    return ""
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetSegmentPath() string {
    return "detail-locations"
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-location" {
        for _, c := range detailLocations.DetailLocation {
            if detailLocations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailCard_DetailLocations_DetailLocation{}
        detailLocations.DetailLocation = append(detailLocations.DetailLocation, child)
        return &detailLocations.DetailLocation[len(detailLocations.DetailLocation)-1]
    }
    return nil
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detailLocations.DetailLocation {
        children[detailLocations.DetailLocation[i].GetSegmentPath()] = &detailLocations.DetailLocation[i]
    }
    return children
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetBundleName() string { return "cisco_ios_xr" }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetYangName() string { return "detail-locations" }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) SetParent(parent types.Entity) { detailLocations.parent = parent }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetParent() types.Entity { return detailLocations.parent }

func (detailLocations *Alarms_Detail_DetailCard_DetailLocations) GetParentYangName() string { return "detail-card" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation
// Specify a card location for alarms.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NodeID of the Location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetFilter() yfilter.YFilter { return detailLocation.YFilter }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) SetFilter(yf yfilter.YFilter) { detailLocation.YFilter = yf }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "active" { return "Active" }
    if yname == "history" { return "History" }
    if yname == "suppressed" { return "Suppressed" }
    if yname == "stats" { return "Stats" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetSegmentPath() string {
    return "detail-location" + "[node-id='" + fmt.Sprintf("%v", detailLocation.NodeId) + "']"
}

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active" {
        return &detailLocation.Active
    }
    if childYangName == "history" {
        return &detailLocation.History
    }
    if childYangName == "suppressed" {
        return &detailLocation.Suppressed
    }
    if childYangName == "stats" {
        return &detailLocation.Stats
    }
    if childYangName == "clients" {
        return &detailLocation.Clients
    }
    return nil
}

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["active"] = &detailLocation.Active
    children["history"] = &detailLocation.History
    children["suppressed"] = &detailLocation.Suppressed
    children["stats"] = &detailLocation.Stats
    children["clients"] = &detailLocation.Clients
    return children
}

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = detailLocation.NodeId
    return leafs
}

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetBundleName() string { return "cisco_ios_xr" }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetYangName() string { return "detail-location" }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) SetParent(parent types.Entity) { detailLocation.parent = parent }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetParent() types.Entity { return detailLocation.parent }

func (detailLocation *Alarms_Detail_DetailCard_DetailLocations_DetailLocation) GetParentYangName() string { return "detail-locations" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active
// Show the active alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetSegmentPath() string {
    return "active"
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range active.AlarmInfo {
            if active.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo{}
        active.AlarmInfo = append(active.AlarmInfo, child)
        return &active.AlarmInfo[len(active.AlarmInfo)-1]
    }
    return nil
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range active.AlarmInfo {
        children[active.AlarmInfo[i].GetSegmentPath()] = &active.AlarmInfo[i]
    }
    return children
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetYangName() string { return "active" }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetParent() types.Entity { return active.parent }

func (active *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active) GetParentYangName() string { return "detail-location" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo
// Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "location" { return "Location" }
    if yname == "aid" { return "Aid" }
    if yname == "tag" { return "Tag" }
    if yname == "module" { return "Module" }
    if yname == "eid" { return "Eid" }
    if yname == "reporting-agent-id" { return "ReportingAgentId" }
    if yname == "pending-sync" { return "PendingSync" }
    if yname == "severity" { return "Severity" }
    if yname == "status" { return "Status" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "service-affecting" { return "ServiceAffecting" }
    if yname == "type" { return "Type" }
    if yname == "interface" { return "Interface" }
    if yname == "alarm-name" { return "AlarmName" }
    if yname == "otn" { return "Otn" }
    if yname == "tca" { return "Tca" }
    return ""
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn" {
        return &alarmInfo.Otn
    }
    if childYangName == "tca" {
        return &alarmInfo.Tca
    }
    return nil
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn"] = &alarmInfo.Otn
    children["tca"] = &alarmInfo.Tca
    return children
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = alarmInfo.Description
    leafs["location"] = alarmInfo.Location
    leafs["aid"] = alarmInfo.Aid
    leafs["tag"] = alarmInfo.Tag
    leafs["module"] = alarmInfo.Module
    leafs["eid"] = alarmInfo.Eid
    leafs["reporting-agent-id"] = alarmInfo.ReportingAgentId
    leafs["pending-sync"] = alarmInfo.PendingSync
    leafs["severity"] = alarmInfo.Severity
    leafs["status"] = alarmInfo.Status
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["service-affecting"] = alarmInfo.ServiceAffecting
    leafs["type"] = alarmInfo.Type
    leafs["interface"] = alarmInfo.Interface
    leafs["alarm-name"] = alarmInfo.AlarmName
    return leafs
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo) GetParentYangName() string { return "active" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "notification-source" { return "NotificationSource" }
    return ""
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = otn.Direction
    leafs["notification-source"] = otn.NotificationSource
    return leafs
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetYangName() string { return "otn" }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetParent() types.Entity { return otn.parent }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Otn) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetFilter() yfilter.YFilter { return tca.YFilter }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) SetFilter(yf yfilter.YFilter) { tca.YFilter = yf }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetGoName(yname string) string {
    if yname == "threshold-value" { return "ThresholdValue" }
    if yname == "current-value" { return "CurrentValue" }
    if yname == "bucket-type" { return "BucketType" }
    return ""
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetSegmentPath() string {
    return "tca"
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-value"] = tca.ThresholdValue
    leafs["current-value"] = tca.CurrentValue
    leafs["bucket-type"] = tca.BucketType
    return leafs
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetBundleName() string { return "cisco_ios_xr" }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetYangName() string { return "tca" }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) SetParent(parent types.Entity) { tca.parent = parent }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetParent() types.Entity { return tca.parent }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Active_AlarmInfo_Tca) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History
// Show the history alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo.
    AlarmInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetSegmentPath() string {
    return "history"
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range history.AlarmInfo {
            if history.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo{}
        history.AlarmInfo = append(history.AlarmInfo, child)
        return &history.AlarmInfo[len(history.AlarmInfo)-1]
    }
    return nil
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range history.AlarmInfo {
        children[history.AlarmInfo[i].GetSegmentPath()] = &history.AlarmInfo[i]
    }
    return children
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetYangName() string { return "history" }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetParent() types.Entity { return history.parent }

func (history *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History) GetParentYangName() string { return "detail-location" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo
// Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "location" { return "Location" }
    if yname == "aid" { return "Aid" }
    if yname == "tag" { return "Tag" }
    if yname == "module" { return "Module" }
    if yname == "eid" { return "Eid" }
    if yname == "reporting-agent-id" { return "ReportingAgentId" }
    if yname == "pending-sync" { return "PendingSync" }
    if yname == "severity" { return "Severity" }
    if yname == "status" { return "Status" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "service-affecting" { return "ServiceAffecting" }
    if yname == "type" { return "Type" }
    if yname == "interface" { return "Interface" }
    if yname == "alarm-name" { return "AlarmName" }
    if yname == "otn" { return "Otn" }
    if yname == "tca" { return "Tca" }
    return ""
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn" {
        return &alarmInfo.Otn
    }
    if childYangName == "tca" {
        return &alarmInfo.Tca
    }
    return nil
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn"] = &alarmInfo.Otn
    children["tca"] = &alarmInfo.Tca
    return children
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = alarmInfo.Description
    leafs["location"] = alarmInfo.Location
    leafs["aid"] = alarmInfo.Aid
    leafs["tag"] = alarmInfo.Tag
    leafs["module"] = alarmInfo.Module
    leafs["eid"] = alarmInfo.Eid
    leafs["reporting-agent-id"] = alarmInfo.ReportingAgentId
    leafs["pending-sync"] = alarmInfo.PendingSync
    leafs["severity"] = alarmInfo.Severity
    leafs["status"] = alarmInfo.Status
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["service-affecting"] = alarmInfo.ServiceAffecting
    leafs["type"] = alarmInfo.Type
    leafs["interface"] = alarmInfo.Interface
    leafs["alarm-name"] = alarmInfo.AlarmName
    return leafs
}

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo) GetParentYangName() string { return "history" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "notification-source" { return "NotificationSource" }
    return ""
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = otn.Direction
    leafs["notification-source"] = otn.NotificationSource
    return leafs
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetYangName() string { return "otn" }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetParent() types.Entity { return otn.parent }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Otn) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca
// TCA feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm Threshold . The type is string with length: 0..20.
    ThresholdValue interface{}

    // Alarm Threshold. The type is string with length: 0..20.
    CurrentValue interface{}

    // Timing Bucket. The type is TimingBucket.
    BucketType interface{}
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetFilter() yfilter.YFilter { return tca.YFilter }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) SetFilter(yf yfilter.YFilter) { tca.YFilter = yf }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetGoName(yname string) string {
    if yname == "threshold-value" { return "ThresholdValue" }
    if yname == "current-value" { return "CurrentValue" }
    if yname == "bucket-type" { return "BucketType" }
    return ""
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetSegmentPath() string {
    return "tca"
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-value"] = tca.ThresholdValue
    leafs["current-value"] = tca.CurrentValue
    leafs["bucket-type"] = tca.BucketType
    return leafs
}

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetBundleName() string { return "cisco_ios_xr" }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetYangName() string { return "tca" }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) SetParent(parent types.Entity) { tca.parent = parent }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetParent() types.Entity { return tca.parent }

func (tca *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_History_AlarmInfo_Tca) GetParentYangName() string { return "alarm-info" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetFilter() yfilter.YFilter { return suppressed.YFilter }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) SetFilter(yf yfilter.YFilter) { suppressed.YFilter = yf }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetGoName(yname string) string {
    if yname == "suppressed-info" { return "SuppressedInfo" }
    return ""
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetSegmentPath() string {
    return "suppressed"
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "suppressed-info" {
        for _, c := range suppressed.SuppressedInfo {
            if suppressed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo{}
        suppressed.SuppressedInfo = append(suppressed.SuppressedInfo, child)
        return &suppressed.SuppressedInfo[len(suppressed.SuppressedInfo)-1]
    }
    return nil
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range suppressed.SuppressedInfo {
        children[suppressed.SuppressedInfo[i].GetSegmentPath()] = &suppressed.SuppressedInfo[i]
    }
    return children
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetBundleName() string { return "cisco_ios_xr" }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetYangName() string { return "suppressed" }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) SetParent(parent types.Entity) { suppressed.parent = parent }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetParent() types.Entity { return suppressed.parent }

func (suppressed *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed) GetParentYangName() string { return "detail-location" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo struct {
    parent types.Entity
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
    Interface interface{}

    // Alarm name. The type is string with length: 0..128.
    AlarmName interface{}

    // OTN feature specific alarm attributes.
    Otn Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn
}

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetFilter() yfilter.YFilter { return suppressedInfo.YFilter }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) SetFilter(yf yfilter.YFilter) { suppressedInfo.YFilter = yf }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "location" { return "Location" }
    if yname == "aid" { return "Aid" }
    if yname == "tag" { return "Tag" }
    if yname == "module" { return "Module" }
    if yname == "eid" { return "Eid" }
    if yname == "reporting-agent-id" { return "ReportingAgentId" }
    if yname == "pending-sync" { return "PendingSync" }
    if yname == "severity" { return "Severity" }
    if yname == "status" { return "Status" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "suppressed-time" { return "SuppressedTime" }
    if yname == "suppressed-timestamp" { return "SuppressedTimestamp" }
    if yname == "service-affecting" { return "ServiceAffecting" }
    if yname == "interface" { return "Interface" }
    if yname == "alarm-name" { return "AlarmName" }
    if yname == "otn" { return "Otn" }
    return ""
}

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetSegmentPath() string {
    return "suppressed-info"
}

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn" {
        return &suppressedInfo.Otn
    }
    return nil
}

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn"] = &suppressedInfo.Otn
    return children
}

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = suppressedInfo.Description
    leafs["location"] = suppressedInfo.Location
    leafs["aid"] = suppressedInfo.Aid
    leafs["tag"] = suppressedInfo.Tag
    leafs["module"] = suppressedInfo.Module
    leafs["eid"] = suppressedInfo.Eid
    leafs["reporting-agent-id"] = suppressedInfo.ReportingAgentId
    leafs["pending-sync"] = suppressedInfo.PendingSync
    leafs["severity"] = suppressedInfo.Severity
    leafs["status"] = suppressedInfo.Status
    leafs["group"] = suppressedInfo.Group
    leafs["set-time"] = suppressedInfo.SetTime
    leafs["set-timestamp"] = suppressedInfo.SetTimestamp
    leafs["suppressed-time"] = suppressedInfo.SuppressedTime
    leafs["suppressed-timestamp"] = suppressedInfo.SuppressedTimestamp
    leafs["service-affecting"] = suppressedInfo.ServiceAffecting
    leafs["interface"] = suppressedInfo.Interface
    leafs["alarm-name"] = suppressedInfo.AlarmName
    return leafs
}

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetBundleName() string { return "cisco_ios_xr" }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetYangName() string { return "suppressed-info" }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) SetParent(parent types.Entity) { suppressedInfo.parent = parent }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetParent() types.Entity { return suppressedInfo.parent }

func (suppressedInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo) GetParentYangName() string { return "suppressed" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn
// OTN feature specific alarm attributes
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm direction . The type is AlarmDirection.
    Direction interface{}

    // Source of Alarm. The type is AlarmNotificationSrc.
    NotificationSource interface{}
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "notification-source" { return "NotificationSource" }
    return ""
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = otn.Direction
    leafs["notification-source"] = otn.NotificationSource
    return leafs
}

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetYangName() string { return "otn" }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetParent() types.Entity { return otn.parent }

func (otn *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Suppressed_SuppressedInfo_Otn) GetParentYangName() string { return "suppressed-info" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats
// Show the service statistics.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats struct {
    parent types.Entity
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

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetGoName(yname string) string {
    if yname == "reported" { return "Reported" }
    if yname == "dropped" { return "Dropped" }
    if yname == "active" { return "Active" }
    if yname == "history" { return "History" }
    if yname == "suppressed" { return "Suppressed" }
    if yname == "sysadmin-active" { return "SysadminActive" }
    if yname == "sysadmin-history" { return "SysadminHistory" }
    if yname == "sysadmin-suppressed" { return "SysadminSuppressed" }
    if yname == "dropped-invalid-aid" { return "DroppedInvalidAid" }
    if yname == "dropped-insuff-mem" { return "DroppedInsuffMem" }
    if yname == "dropped-db-error" { return "DroppedDbError" }
    if yname == "dropped-clear-without-set" { return "DroppedClearWithoutSet" }
    if yname == "dropped-duplicate" { return "DroppedDuplicate" }
    if yname == "cache-hit" { return "CacheHit" }
    if yname == "cache-miss" { return "CacheMiss" }
    return ""
}

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reported"] = stats.Reported
    leafs["dropped"] = stats.Dropped
    leafs["active"] = stats.Active
    leafs["history"] = stats.History
    leafs["suppressed"] = stats.Suppressed
    leafs["sysadmin-active"] = stats.SysadminActive
    leafs["sysadmin-history"] = stats.SysadminHistory
    leafs["sysadmin-suppressed"] = stats.SysadminSuppressed
    leafs["dropped-invalid-aid"] = stats.DroppedInvalidAid
    leafs["dropped-insuff-mem"] = stats.DroppedInsuffMem
    leafs["dropped-db-error"] = stats.DroppedDbError
    leafs["dropped-clear-without-set"] = stats.DroppedClearWithoutSet
    leafs["dropped-duplicate"] = stats.DroppedDuplicate
    leafs["cache-hit"] = stats.CacheHit
    leafs["cache-miss"] = stats.CacheMiss
    return leafs
}

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetYangName() string { return "stats" }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetParent() types.Entity { return stats.parent }

func (stats *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Stats) GetParentYangName() string { return "detail-location" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients
// Show the clients associated with this
// service.
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Client List. The type is slice of
    // Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo.
    ClientInfo []Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetGoName(yname string) string {
    if yname == "client-info" { return "ClientInfo" }
    return ""
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-info" {
        for _, c := range clients.ClientInfo {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo{}
        clients.ClientInfo = append(clients.ClientInfo, child)
        return &clients.ClientInfo[len(clients.ClientInfo)-1]
    }
    return nil
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.ClientInfo {
        children[clients.ClientInfo[i].GetSegmentPath()] = &clients.ClientInfo[i]
    }
    return children
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetYangName() string { return "clients" }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients) GetParentYangName() string { return "detail-location" }

// Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo
// Client List
type Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo struct {
    parent types.Entity
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

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetFilter() yfilter.YFilter { return clientInfo.YFilter }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) SetFilter(yf yfilter.YFilter) { clientInfo.YFilter = yf }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "id" { return "Id" }
    if yname == "location" { return "Location" }
    if yname == "handle" { return "Handle" }
    if yname == "state" { return "State" }
    if yname == "type" { return "Type" }
    if yname == "filter-disp" { return "FilterDisp" }
    if yname == "subscriber-id" { return "SubscriberId" }
    if yname == "filter-severity" { return "FilterSeverity" }
    if yname == "filter-state" { return "FilterState" }
    if yname == "filter-group" { return "FilterGroup" }
    if yname == "connect-count" { return "ConnectCount" }
    if yname == "connect-timestamp" { return "ConnectTimestamp" }
    if yname == "get-count" { return "GetCount" }
    if yname == "subscribe-count" { return "SubscribeCount" }
    if yname == "report-count" { return "ReportCount" }
    return ""
}

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetSegmentPath() string {
    return "client-info"
}

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = clientInfo.Name
    leafs["id"] = clientInfo.Id
    leafs["location"] = clientInfo.Location
    leafs["handle"] = clientInfo.Handle
    leafs["state"] = clientInfo.State
    leafs["type"] = clientInfo.Type
    leafs["filter-disp"] = clientInfo.FilterDisp
    leafs["subscriber-id"] = clientInfo.SubscriberId
    leafs["filter-severity"] = clientInfo.FilterSeverity
    leafs["filter-state"] = clientInfo.FilterState
    leafs["filter-group"] = clientInfo.FilterGroup
    leafs["connect-count"] = clientInfo.ConnectCount
    leafs["connect-timestamp"] = clientInfo.ConnectTimestamp
    leafs["get-count"] = clientInfo.GetCount
    leafs["subscribe-count"] = clientInfo.SubscribeCount
    leafs["report-count"] = clientInfo.ReportCount
    return leafs
}

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetBundleName() string { return "cisco_ios_xr" }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetYangName() string { return "client-info" }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) SetParent(parent types.Entity) { clientInfo.parent = parent }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetParent() types.Entity { return clientInfo.parent }

func (clientInfo *Alarms_Detail_DetailCard_DetailLocations_DetailLocation_Clients_ClientInfo) GetParentYangName() string { return "clients" }

// Alarms_Brief
// A set of brief alarm commands.
type Alarms_Brief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Show brief card scope alarm related data.
    BriefCard Alarms_Brief_BriefCard

    // Show brief system scope alarm related data.
    BriefSystem Alarms_Brief_BriefSystem
}

func (brief *Alarms_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *Alarms_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *Alarms_Brief) GetGoName(yname string) string {
    if yname == "brief-card" { return "BriefCard" }
    if yname == "brief-system" { return "BriefSystem" }
    return ""
}

func (brief *Alarms_Brief) GetSegmentPath() string {
    return "brief"
}

func (brief *Alarms_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-card" {
        return &brief.BriefCard
    }
    if childYangName == "brief-system" {
        return &brief.BriefSystem
    }
    return nil
}

func (brief *Alarms_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-card"] = &brief.BriefCard
    children["brief-system"] = &brief.BriefSystem
    return children
}

func (brief *Alarms_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (brief *Alarms_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *Alarms_Brief) GetYangName() string { return "brief" }

func (brief *Alarms_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *Alarms_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *Alarms_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *Alarms_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *Alarms_Brief) GetParent() types.Entity { return brief.parent }

func (brief *Alarms_Brief) GetParentYangName() string { return "alarms" }

// Alarms_Brief_BriefCard
// Show brief card scope alarm related data.
type Alarms_Brief_BriefCard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of BriefLocation.
    BriefLocations Alarms_Brief_BriefCard_BriefLocations
}

func (briefCard *Alarms_Brief_BriefCard) GetFilter() yfilter.YFilter { return briefCard.YFilter }

func (briefCard *Alarms_Brief_BriefCard) SetFilter(yf yfilter.YFilter) { briefCard.YFilter = yf }

func (briefCard *Alarms_Brief_BriefCard) GetGoName(yname string) string {
    if yname == "brief-locations" { return "BriefLocations" }
    return ""
}

func (briefCard *Alarms_Brief_BriefCard) GetSegmentPath() string {
    return "brief-card"
}

func (briefCard *Alarms_Brief_BriefCard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-locations" {
        return &briefCard.BriefLocations
    }
    return nil
}

func (briefCard *Alarms_Brief_BriefCard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-locations"] = &briefCard.BriefLocations
    return children
}

func (briefCard *Alarms_Brief_BriefCard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefCard *Alarms_Brief_BriefCard) GetBundleName() string { return "cisco_ios_xr" }

func (briefCard *Alarms_Brief_BriefCard) GetYangName() string { return "brief-card" }

func (briefCard *Alarms_Brief_BriefCard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefCard *Alarms_Brief_BriefCard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefCard *Alarms_Brief_BriefCard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefCard *Alarms_Brief_BriefCard) SetParent(parent types.Entity) { briefCard.parent = parent }

func (briefCard *Alarms_Brief_BriefCard) GetParent() types.Entity { return briefCard.parent }

func (briefCard *Alarms_Brief_BriefCard) GetParentYangName() string { return "brief" }

// Alarms_Brief_BriefCard_BriefLocations
// Table of BriefLocation
type Alarms_Brief_BriefCard_BriefLocations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify a card location for alarms. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation.
    BriefLocation []Alarms_Brief_BriefCard_BriefLocations_BriefLocation
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetFilter() yfilter.YFilter { return briefLocations.YFilter }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) SetFilter(yf yfilter.YFilter) { briefLocations.YFilter = yf }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetGoName(yname string) string {
    if yname == "brief-location" { return "BriefLocation" }
    return ""
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetSegmentPath() string {
    return "brief-locations"
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-location" {
        for _, c := range briefLocations.BriefLocation {
            if briefLocations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Brief_BriefCard_BriefLocations_BriefLocation{}
        briefLocations.BriefLocation = append(briefLocations.BriefLocation, child)
        return &briefLocations.BriefLocation[len(briefLocations.BriefLocation)-1]
    }
    return nil
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefLocations.BriefLocation {
        children[briefLocations.BriefLocation[i].GetSegmentPath()] = &briefLocations.BriefLocation[i]
    }
    return children
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetBundleName() string { return "cisco_ios_xr" }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetYangName() string { return "brief-locations" }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) SetParent(parent types.Entity) { briefLocations.parent = parent }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetParent() types.Entity { return briefLocations.parent }

func (briefLocations *Alarms_Brief_BriefCard_BriefLocations) GetParentYangName() string { return "brief-card" }

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation
// Specify a card location for alarms.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NodeID of the Location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Show the active alarms at this scope.
    Active Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active

    // Show the history alarms at this scope.
    History Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History

    // Show the suppressed alarms at this scope.
    Suppressed Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed
}

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetFilter() yfilter.YFilter { return briefLocation.YFilter }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) SetFilter(yf yfilter.YFilter) { briefLocation.YFilter = yf }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "active" { return "Active" }
    if yname == "history" { return "History" }
    if yname == "suppressed" { return "Suppressed" }
    return ""
}

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetSegmentPath() string {
    return "brief-location" + "[node-id='" + fmt.Sprintf("%v", briefLocation.NodeId) + "']"
}

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active" {
        return &briefLocation.Active
    }
    if childYangName == "history" {
        return &briefLocation.History
    }
    if childYangName == "suppressed" {
        return &briefLocation.Suppressed
    }
    return nil
}

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["active"] = &briefLocation.Active
    children["history"] = &briefLocation.History
    children["suppressed"] = &briefLocation.Suppressed
    return children
}

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = briefLocation.NodeId
    return leafs
}

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetBundleName() string { return "cisco_ios_xr" }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetYangName() string { return "brief-location" }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) SetParent(parent types.Entity) { briefLocation.parent = parent }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetParent() types.Entity { return briefLocation.parent }

func (briefLocation *Alarms_Brief_BriefCard_BriefLocations_BriefLocation) GetParentYangName() string { return "brief-locations" }

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active
// Show the active alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetSegmentPath() string {
    return "active"
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range active.AlarmInfo {
            if active.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo{}
        active.AlarmInfo = append(active.AlarmInfo, child)
        return &active.AlarmInfo[len(active.AlarmInfo)-1]
    }
    return nil
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range active.AlarmInfo {
        children[active.AlarmInfo[i].GetSegmentPath()] = &active.AlarmInfo[i]
    }
    return children
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetYangName() string { return "active" }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetParent() types.Entity { return active.parent }

func (active *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active) GetParentYangName() string { return "brief-location" }

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo
// Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "severity" { return "Severity" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "description" { return "Description" }
    return ""
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = alarmInfo.Location
    leafs["severity"] = alarmInfo.Severity
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["description"] = alarmInfo.Description
    return leafs
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Active_AlarmInfo) GetParentYangName() string { return "active" }

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History
// Show the history alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetSegmentPath() string {
    return "history"
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range history.AlarmInfo {
            if history.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo{}
        history.AlarmInfo = append(history.AlarmInfo, child)
        return &history.AlarmInfo[len(history.AlarmInfo)-1]
    }
    return nil
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range history.AlarmInfo {
        children[history.AlarmInfo[i].GetSegmentPath()] = &history.AlarmInfo[i]
    }
    return children
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetYangName() string { return "history" }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetParent() types.Entity { return history.parent }

func (history *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History) GetParentYangName() string { return "brief-location" }

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo
// Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "severity" { return "Severity" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "description" { return "Description" }
    return ""
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = alarmInfo.Location
    leafs["severity"] = alarmInfo.Severity
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["description"] = alarmInfo.Description
    return leafs
}

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_History_AlarmInfo) GetParentYangName() string { return "history" }

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetFilter() yfilter.YFilter { return suppressed.YFilter }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) SetFilter(yf yfilter.YFilter) { suppressed.YFilter = yf }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetGoName(yname string) string {
    if yname == "suppressed-info" { return "SuppressedInfo" }
    return ""
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetSegmentPath() string {
    return "suppressed"
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "suppressed-info" {
        for _, c := range suppressed.SuppressedInfo {
            if suppressed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo{}
        suppressed.SuppressedInfo = append(suppressed.SuppressedInfo, child)
        return &suppressed.SuppressedInfo[len(suppressed.SuppressedInfo)-1]
    }
    return nil
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range suppressed.SuppressedInfo {
        children[suppressed.SuppressedInfo[i].GetSegmentPath()] = &suppressed.SuppressedInfo[i]
    }
    return children
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetBundleName() string { return "cisco_ios_xr" }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetYangName() string { return "suppressed" }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) SetParent(parent types.Entity) { suppressed.parent = parent }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetParent() types.Entity { return suppressed.parent }

func (suppressed *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed) GetParentYangName() string { return "brief-location" }

// Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo struct {
    parent types.Entity
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

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetFilter() yfilter.YFilter { return suppressedInfo.YFilter }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) SetFilter(yf yfilter.YFilter) { suppressedInfo.YFilter = yf }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "severity" { return "Severity" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "suppressed-time" { return "SuppressedTime" }
    if yname == "suppressed-timestamp" { return "SuppressedTimestamp" }
    if yname == "description" { return "Description" }
    return ""
}

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetSegmentPath() string {
    return "suppressed-info"
}

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = suppressedInfo.Location
    leafs["severity"] = suppressedInfo.Severity
    leafs["group"] = suppressedInfo.Group
    leafs["set-time"] = suppressedInfo.SetTime
    leafs["set-timestamp"] = suppressedInfo.SetTimestamp
    leafs["suppressed-time"] = suppressedInfo.SuppressedTime
    leafs["suppressed-timestamp"] = suppressedInfo.SuppressedTimestamp
    leafs["description"] = suppressedInfo.Description
    return leafs
}

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetBundleName() string { return "cisco_ios_xr" }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetYangName() string { return "suppressed-info" }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) SetParent(parent types.Entity) { suppressedInfo.parent = parent }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetParent() types.Entity { return suppressedInfo.parent }

func (suppressedInfo *Alarms_Brief_BriefCard_BriefLocations_BriefLocation_Suppressed_SuppressedInfo) GetParentYangName() string { return "suppressed" }

// Alarms_Brief_BriefSystem
// Show brief system scope alarm related data.
type Alarms_Brief_BriefSystem struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Show the active alarms at this scope.
    Active Alarms_Brief_BriefSystem_Active

    // Show the history alarms at this scope.
    History Alarms_Brief_BriefSystem_History

    // Show the suppressed alarms at this scope.
    Suppressed Alarms_Brief_BriefSystem_Suppressed
}

func (briefSystem *Alarms_Brief_BriefSystem) GetFilter() yfilter.YFilter { return briefSystem.YFilter }

func (briefSystem *Alarms_Brief_BriefSystem) SetFilter(yf yfilter.YFilter) { briefSystem.YFilter = yf }

func (briefSystem *Alarms_Brief_BriefSystem) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "history" { return "History" }
    if yname == "suppressed" { return "Suppressed" }
    return ""
}

func (briefSystem *Alarms_Brief_BriefSystem) GetSegmentPath() string {
    return "brief-system"
}

func (briefSystem *Alarms_Brief_BriefSystem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active" {
        return &briefSystem.Active
    }
    if childYangName == "history" {
        return &briefSystem.History
    }
    if childYangName == "suppressed" {
        return &briefSystem.Suppressed
    }
    return nil
}

func (briefSystem *Alarms_Brief_BriefSystem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["active"] = &briefSystem.Active
    children["history"] = &briefSystem.History
    children["suppressed"] = &briefSystem.Suppressed
    return children
}

func (briefSystem *Alarms_Brief_BriefSystem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefSystem *Alarms_Brief_BriefSystem) GetBundleName() string { return "cisco_ios_xr" }

func (briefSystem *Alarms_Brief_BriefSystem) GetYangName() string { return "brief-system" }

func (briefSystem *Alarms_Brief_BriefSystem) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefSystem *Alarms_Brief_BriefSystem) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefSystem *Alarms_Brief_BriefSystem) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefSystem *Alarms_Brief_BriefSystem) SetParent(parent types.Entity) { briefSystem.parent = parent }

func (briefSystem *Alarms_Brief_BriefSystem) GetParent() types.Entity { return briefSystem.parent }

func (briefSystem *Alarms_Brief_BriefSystem) GetParentYangName() string { return "brief" }

// Alarms_Brief_BriefSystem_Active
// Show the active alarms at this scope.
type Alarms_Brief_BriefSystem_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of Alarms_Brief_BriefSystem_Active_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefSystem_Active_AlarmInfo
}

func (active *Alarms_Brief_BriefSystem_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *Alarms_Brief_BriefSystem_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *Alarms_Brief_BriefSystem_Active) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (active *Alarms_Brief_BriefSystem_Active) GetSegmentPath() string {
    return "active"
}

func (active *Alarms_Brief_BriefSystem_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range active.AlarmInfo {
            if active.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Brief_BriefSystem_Active_AlarmInfo{}
        active.AlarmInfo = append(active.AlarmInfo, child)
        return &active.AlarmInfo[len(active.AlarmInfo)-1]
    }
    return nil
}

func (active *Alarms_Brief_BriefSystem_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range active.AlarmInfo {
        children[active.AlarmInfo[i].GetSegmentPath()] = &active.AlarmInfo[i]
    }
    return children
}

func (active *Alarms_Brief_BriefSystem_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *Alarms_Brief_BriefSystem_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *Alarms_Brief_BriefSystem_Active) GetYangName() string { return "active" }

func (active *Alarms_Brief_BriefSystem_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *Alarms_Brief_BriefSystem_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *Alarms_Brief_BriefSystem_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *Alarms_Brief_BriefSystem_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *Alarms_Brief_BriefSystem_Active) GetParent() types.Entity { return active.parent }

func (active *Alarms_Brief_BriefSystem_Active) GetParentYangName() string { return "brief-system" }

// Alarms_Brief_BriefSystem_Active_AlarmInfo
// Alarm List
type Alarms_Brief_BriefSystem_Active_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "severity" { return "Severity" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "description" { return "Description" }
    return ""
}

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = alarmInfo.Location
    leafs["severity"] = alarmInfo.Severity
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["description"] = alarmInfo.Description
    return leafs
}

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Brief_BriefSystem_Active_AlarmInfo) GetParentYangName() string { return "active" }

// Alarms_Brief_BriefSystem_History
// Show the history alarms at this scope.
type Alarms_Brief_BriefSystem_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm List. The type is slice of
    // Alarms_Brief_BriefSystem_History_AlarmInfo.
    AlarmInfo []Alarms_Brief_BriefSystem_History_AlarmInfo
}

func (history *Alarms_Brief_BriefSystem_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Alarms_Brief_BriefSystem_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Alarms_Brief_BriefSystem_History) GetGoName(yname string) string {
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (history *Alarms_Brief_BriefSystem_History) GetSegmentPath() string {
    return "history"
}

func (history *Alarms_Brief_BriefSystem_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        for _, c := range history.AlarmInfo {
            if history.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Brief_BriefSystem_History_AlarmInfo{}
        history.AlarmInfo = append(history.AlarmInfo, child)
        return &history.AlarmInfo[len(history.AlarmInfo)-1]
    }
    return nil
}

func (history *Alarms_Brief_BriefSystem_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range history.AlarmInfo {
        children[history.AlarmInfo[i].GetSegmentPath()] = &history.AlarmInfo[i]
    }
    return children
}

func (history *Alarms_Brief_BriefSystem_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (history *Alarms_Brief_BriefSystem_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Alarms_Brief_BriefSystem_History) GetYangName() string { return "history" }

func (history *Alarms_Brief_BriefSystem_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Alarms_Brief_BriefSystem_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Alarms_Brief_BriefSystem_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Alarms_Brief_BriefSystem_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Alarms_Brief_BriefSystem_History) GetParent() types.Entity { return history.parent }

func (history *Alarms_Brief_BriefSystem_History) GetParentYangName() string { return "brief-system" }

// Alarms_Brief_BriefSystem_History_AlarmInfo
// Alarm List
type Alarms_Brief_BriefSystem_History_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "severity" { return "Severity" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "clear-time" { return "ClearTime" }
    if yname == "clear-timestamp" { return "ClearTimestamp" }
    if yname == "description" { return "Description" }
    return ""
}

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = alarmInfo.Location
    leafs["severity"] = alarmInfo.Severity
    leafs["group"] = alarmInfo.Group
    leafs["set-time"] = alarmInfo.SetTime
    leafs["set-timestamp"] = alarmInfo.SetTimestamp
    leafs["clear-time"] = alarmInfo.ClearTime
    leafs["clear-timestamp"] = alarmInfo.ClearTimestamp
    leafs["description"] = alarmInfo.Description
    return leafs
}

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Alarms_Brief_BriefSystem_History_AlarmInfo) GetParentYangName() string { return "history" }

// Alarms_Brief_BriefSystem_Suppressed
// Show the suppressed alarms at this scope.
type Alarms_Brief_BriefSystem_Suppressed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Suppressed Alarm List. The type is slice of
    // Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo.
    SuppressedInfo []Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetFilter() yfilter.YFilter { return suppressed.YFilter }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) SetFilter(yf yfilter.YFilter) { suppressed.YFilter = yf }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetGoName(yname string) string {
    if yname == "suppressed-info" { return "SuppressedInfo" }
    return ""
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetSegmentPath() string {
    return "suppressed"
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "suppressed-info" {
        for _, c := range suppressed.SuppressedInfo {
            if suppressed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo{}
        suppressed.SuppressedInfo = append(suppressed.SuppressedInfo, child)
        return &suppressed.SuppressedInfo[len(suppressed.SuppressedInfo)-1]
    }
    return nil
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range suppressed.SuppressedInfo {
        children[suppressed.SuppressedInfo[i].GetSegmentPath()] = &suppressed.SuppressedInfo[i]
    }
    return children
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetBundleName() string { return "cisco_ios_xr" }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetYangName() string { return "suppressed" }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) SetParent(parent types.Entity) { suppressed.parent = parent }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetParent() types.Entity { return suppressed.parent }

func (suppressed *Alarms_Brief_BriefSystem_Suppressed) GetParentYangName() string { return "brief-system" }

// Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo
// Suppressed Alarm List
type Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo struct {
    parent types.Entity
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

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetFilter() yfilter.YFilter { return suppressedInfo.YFilter }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) SetFilter(yf yfilter.YFilter) { suppressedInfo.YFilter = yf }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "severity" { return "Severity" }
    if yname == "group" { return "Group" }
    if yname == "set-time" { return "SetTime" }
    if yname == "set-timestamp" { return "SetTimestamp" }
    if yname == "suppressed-time" { return "SuppressedTime" }
    if yname == "suppressed-timestamp" { return "SuppressedTimestamp" }
    if yname == "description" { return "Description" }
    return ""
}

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetSegmentPath() string {
    return "suppressed-info"
}

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = suppressedInfo.Location
    leafs["severity"] = suppressedInfo.Severity
    leafs["group"] = suppressedInfo.Group
    leafs["set-time"] = suppressedInfo.SetTime
    leafs["set-timestamp"] = suppressedInfo.SetTimestamp
    leafs["suppressed-time"] = suppressedInfo.SuppressedTime
    leafs["suppressed-timestamp"] = suppressedInfo.SuppressedTimestamp
    leafs["description"] = suppressedInfo.Description
    return leafs
}

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetBundleName() string { return "cisco_ios_xr" }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetYangName() string { return "suppressed-info" }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) SetParent(parent types.Entity) { suppressedInfo.parent = parent }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetParent() types.Entity { return suppressedInfo.parent }

func (suppressedInfo *Alarms_Brief_BriefSystem_Suppressed_SuppressedInfo) GetParentYangName() string { return "suppressed" }

