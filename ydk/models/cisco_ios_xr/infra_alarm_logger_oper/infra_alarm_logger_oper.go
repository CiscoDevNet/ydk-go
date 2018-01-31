// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-alarm-logger package operational data.
// 
// This module contains definitions
// for the following management objects:
//   alarm-logger: Alarm Logger operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_alarm_logger_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_alarm_logger_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-alarm-logger-oper alarm-logger}", reflect.TypeOf(AlarmLogger{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-alarm-logger-oper:alarm-logger", reflect.TypeOf(AlarmLogger{}))
}

// AlAlarmBistate represents Al alarm bistate
type AlAlarmBistate string

const (
    // not available
    AlAlarmBistate_not_available AlAlarmBistate = "not-available"

    // active
    AlAlarmBistate_active AlAlarmBistate = "active"

    // clear
    AlAlarmBistate_clear AlAlarmBistate = "clear"
)

// AlAlarmSeverity represents Al alarm severity
type AlAlarmSeverity string

const (
    // unknown
    AlAlarmSeverity_unknown AlAlarmSeverity = "unknown"

    // emergency
    AlAlarmSeverity_emergency AlAlarmSeverity = "emergency"

    // alert
    AlAlarmSeverity_alert AlAlarmSeverity = "alert"

    // critical
    AlAlarmSeverity_critical AlAlarmSeverity = "critical"

    // error
    AlAlarmSeverity_error AlAlarmSeverity = "error"

    // warning
    AlAlarmSeverity_warning AlAlarmSeverity = "warning"

    // notice
    AlAlarmSeverity_notice AlAlarmSeverity = "notice"

    // informational
    AlAlarmSeverity_informational AlAlarmSeverity = "informational"

    // debugging
    AlAlarmSeverity_debugging AlAlarmSeverity = "debugging"
)

// AlarmLogger
// Alarm Logger operational data
type AlarmLogger struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describes buffer utilization and parameters configured.
    BufferStatus AlarmLogger_BufferStatus

    // Table that contains the database of logged alarms.
    Alarms AlarmLogger_Alarms
}

func (alarmLogger *AlarmLogger) GetFilter() yfilter.YFilter { return alarmLogger.YFilter }

func (alarmLogger *AlarmLogger) SetFilter(yf yfilter.YFilter) { alarmLogger.YFilter = yf }

func (alarmLogger *AlarmLogger) GetGoName(yname string) string {
    if yname == "buffer-status" { return "BufferStatus" }
    if yname == "alarms" { return "Alarms" }
    return ""
}

func (alarmLogger *AlarmLogger) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-alarm-logger-oper:alarm-logger"
}

func (alarmLogger *AlarmLogger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "buffer-status" {
        return &alarmLogger.BufferStatus
    }
    if childYangName == "alarms" {
        return &alarmLogger.Alarms
    }
    return nil
}

func (alarmLogger *AlarmLogger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["buffer-status"] = &alarmLogger.BufferStatus
    children["alarms"] = &alarmLogger.Alarms
    return children
}

func (alarmLogger *AlarmLogger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alarmLogger *AlarmLogger) GetBundleName() string { return "cisco_ios_xr" }

func (alarmLogger *AlarmLogger) GetYangName() string { return "alarm-logger" }

func (alarmLogger *AlarmLogger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmLogger *AlarmLogger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmLogger *AlarmLogger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmLogger *AlarmLogger) SetParent(parent types.Entity) { alarmLogger.parent = parent }

func (alarmLogger *AlarmLogger) GetParent() types.Entity { return alarmLogger.parent }

func (alarmLogger *AlarmLogger) GetParentYangName() string { return "Cisco-IOS-XR-infra-alarm-logger-oper" }

// AlarmLogger_BufferStatus
// Describes buffer utilization and parameters
// configured
type AlarmLogger_BufferStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current Logging Buffer Size (Bytes). The type is interface{} with range:
    // 0..4294967295. Units are byte.
    LogBufferSize interface{}

    // Maximum Logging Buffer Size (Bytes) . The type is interface{} with range:
    // 0..4294967295. Units are byte.
    MaxLogBufferSize interface{}

    // Number of Records in the Buffer. The type is interface{} with range:
    // 0..4294967295.
    RecordCount interface{}

    // Percentage of the buffer utilization which, when exceeded, triggers the 
    // generation of a notification for the clients of the XML agent. The type is
    // interface{} with range: 0..4294967295. Units are percentage.
    CapacityThreshold interface{}

    // Severity Filter. The type is AlAlarmSeverity.
    SeverityFilter interface{}
}

func (bufferStatus *AlarmLogger_BufferStatus) GetFilter() yfilter.YFilter { return bufferStatus.YFilter }

func (bufferStatus *AlarmLogger_BufferStatus) SetFilter(yf yfilter.YFilter) { bufferStatus.YFilter = yf }

func (bufferStatus *AlarmLogger_BufferStatus) GetGoName(yname string) string {
    if yname == "log-buffer-size" { return "LogBufferSize" }
    if yname == "max-log-buffer-size" { return "MaxLogBufferSize" }
    if yname == "record-count" { return "RecordCount" }
    if yname == "capacity-threshold" { return "CapacityThreshold" }
    if yname == "severity-filter" { return "SeverityFilter" }
    return ""
}

func (bufferStatus *AlarmLogger_BufferStatus) GetSegmentPath() string {
    return "buffer-status"
}

func (bufferStatus *AlarmLogger_BufferStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bufferStatus *AlarmLogger_BufferStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bufferStatus *AlarmLogger_BufferStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-buffer-size"] = bufferStatus.LogBufferSize
    leafs["max-log-buffer-size"] = bufferStatus.MaxLogBufferSize
    leafs["record-count"] = bufferStatus.RecordCount
    leafs["capacity-threshold"] = bufferStatus.CapacityThreshold
    leafs["severity-filter"] = bufferStatus.SeverityFilter
    return leafs
}

func (bufferStatus *AlarmLogger_BufferStatus) GetBundleName() string { return "cisco_ios_xr" }

func (bufferStatus *AlarmLogger_BufferStatus) GetYangName() string { return "buffer-status" }

func (bufferStatus *AlarmLogger_BufferStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bufferStatus *AlarmLogger_BufferStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bufferStatus *AlarmLogger_BufferStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bufferStatus *AlarmLogger_BufferStatus) SetParent(parent types.Entity) { bufferStatus.parent = parent }

func (bufferStatus *AlarmLogger_BufferStatus) GetParent() types.Entity { return bufferStatus.parent }

func (bufferStatus *AlarmLogger_BufferStatus) GetParentYangName() string { return "alarm-logger" }

// AlarmLogger_Alarms
// Table that contains the database of logged
// alarms
type AlarmLogger_Alarms struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One of the logged alarms. The type is slice of AlarmLogger_Alarms_Alarm.
    Alarm []AlarmLogger_Alarms_Alarm
}

func (alarms *AlarmLogger_Alarms) GetFilter() yfilter.YFilter { return alarms.YFilter }

func (alarms *AlarmLogger_Alarms) SetFilter(yf yfilter.YFilter) { alarms.YFilter = yf }

func (alarms *AlarmLogger_Alarms) GetGoName(yname string) string {
    if yname == "alarm" { return "Alarm" }
    return ""
}

func (alarms *AlarmLogger_Alarms) GetSegmentPath() string {
    return "alarms"
}

func (alarms *AlarmLogger_Alarms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm" {
        for _, c := range alarms.Alarm {
            if alarms.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AlarmLogger_Alarms_Alarm{}
        alarms.Alarm = append(alarms.Alarm, child)
        return &alarms.Alarm[len(alarms.Alarm)-1]
    }
    return nil
}

func (alarms *AlarmLogger_Alarms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alarms.Alarm {
        children[alarms.Alarm[i].GetSegmentPath()] = &alarms.Alarm[i]
    }
    return children
}

func (alarms *AlarmLogger_Alarms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alarms *AlarmLogger_Alarms) GetBundleName() string { return "cisco_ios_xr" }

func (alarms *AlarmLogger_Alarms) GetYangName() string { return "alarms" }

func (alarms *AlarmLogger_Alarms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarms *AlarmLogger_Alarms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarms *AlarmLogger_Alarms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarms *AlarmLogger_Alarms) SetParent(parent types.Entity) { alarms.parent = parent }

func (alarms *AlarmLogger_Alarms) GetParent() types.Entity { return alarms.parent }

func (alarms *AlarmLogger_Alarms) GetParentYangName() string { return "alarm-logger" }

// AlarmLogger_Alarms_Alarm
// One of the logged alarms
type AlarmLogger_Alarms_Alarm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Event ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EventId interface{}

    // Source Identifier(Location).Indicates the node in which the alarm was
    // generated. The type is string.
    SourceId interface{}

    // Time when the alarm was generated. It is expressed in number of
    // milliseconds since 00:00 :00 UTC, January 1, 1970. The type is interface{}
    // with range: 0..18446744073709551615. Units are millisecond.
    Timestamp interface{}

    // Category of the alarm. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs to. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}

    // Severity of the alarm. The type is AlAlarmSeverity.
    Severity interface{}

    // State of the alarm (bistate alarms only). The type is AlAlarmBistate.
    State interface{}

    // Correlation Identifier. The type is interface{} with range: 0..4294967295.
    CorrelationId interface{}

    // Indicates the event id admin-level. The type is bool.
    IsAdmin interface{}

    // Full text of the Alarm. The type is string.
    AdditionalText interface{}
}

func (alarm *AlarmLogger_Alarms_Alarm) GetFilter() yfilter.YFilter { return alarm.YFilter }

func (alarm *AlarmLogger_Alarms_Alarm) SetFilter(yf yfilter.YFilter) { alarm.YFilter = yf }

func (alarm *AlarmLogger_Alarms_Alarm) GetGoName(yname string) string {
    if yname == "event-id" { return "EventId" }
    if yname == "source-id" { return "SourceId" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "code" { return "Code" }
    if yname == "severity" { return "Severity" }
    if yname == "state" { return "State" }
    if yname == "correlation-id" { return "CorrelationId" }
    if yname == "is-admin" { return "IsAdmin" }
    if yname == "additional-text" { return "AdditionalText" }
    return ""
}

func (alarm *AlarmLogger_Alarms_Alarm) GetSegmentPath() string {
    return "alarm" + "[event-id='" + fmt.Sprintf("%v", alarm.EventId) + "']"
}

func (alarm *AlarmLogger_Alarms_Alarm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarm *AlarmLogger_Alarms_Alarm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarm *AlarmLogger_Alarms_Alarm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-id"] = alarm.EventId
    leafs["source-id"] = alarm.SourceId
    leafs["timestamp"] = alarm.Timestamp
    leafs["category"] = alarm.Category
    leafs["group"] = alarm.Group
    leafs["code"] = alarm.Code
    leafs["severity"] = alarm.Severity
    leafs["state"] = alarm.State
    leafs["correlation-id"] = alarm.CorrelationId
    leafs["is-admin"] = alarm.IsAdmin
    leafs["additional-text"] = alarm.AdditionalText
    return leafs
}

func (alarm *AlarmLogger_Alarms_Alarm) GetBundleName() string { return "cisco_ios_xr" }

func (alarm *AlarmLogger_Alarms_Alarm) GetYangName() string { return "alarm" }

func (alarm *AlarmLogger_Alarms_Alarm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarm *AlarmLogger_Alarms_Alarm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarm *AlarmLogger_Alarms_Alarm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarm *AlarmLogger_Alarms_Alarm) SetParent(parent types.Entity) { alarm.parent = parent }

func (alarm *AlarmLogger_Alarms_Alarm) GetParent() types.Entity { return alarm.parent }

func (alarm *AlarmLogger_Alarms_Alarm) GetParentYangName() string { return "alarms" }

