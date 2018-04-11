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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes buffer utilization and parameters configured.
    BufferStatus AlarmLogger_BufferStatus

    // Table that contains the database of logged alarms.
    Alarms AlarmLogger_Alarms
}

func (alarmLogger *AlarmLogger) GetEntityData() *types.CommonEntityData {
    alarmLogger.EntityData.YFilter = alarmLogger.YFilter
    alarmLogger.EntityData.YangName = "alarm-logger"
    alarmLogger.EntityData.BundleName = "cisco_ios_xr"
    alarmLogger.EntityData.ParentYangName = "Cisco-IOS-XR-infra-alarm-logger-oper"
    alarmLogger.EntityData.SegmentPath = "Cisco-IOS-XR-infra-alarm-logger-oper:alarm-logger"
    alarmLogger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmLogger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmLogger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmLogger.EntityData.Children = make(map[string]types.YChild)
    alarmLogger.EntityData.Children["buffer-status"] = types.YChild{"BufferStatus", &alarmLogger.BufferStatus}
    alarmLogger.EntityData.Children["alarms"] = types.YChild{"Alarms", &alarmLogger.Alarms}
    alarmLogger.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alarmLogger.EntityData)
}

// AlarmLogger_BufferStatus
// Describes buffer utilization and parameters
// configured
type AlarmLogger_BufferStatus struct {
    EntityData types.CommonEntityData
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

func (bufferStatus *AlarmLogger_BufferStatus) GetEntityData() *types.CommonEntityData {
    bufferStatus.EntityData.YFilter = bufferStatus.YFilter
    bufferStatus.EntityData.YangName = "buffer-status"
    bufferStatus.EntityData.BundleName = "cisco_ios_xr"
    bufferStatus.EntityData.ParentYangName = "alarm-logger"
    bufferStatus.EntityData.SegmentPath = "buffer-status"
    bufferStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bufferStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bufferStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bufferStatus.EntityData.Children = make(map[string]types.YChild)
    bufferStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    bufferStatus.EntityData.Leafs["log-buffer-size"] = types.YLeaf{"LogBufferSize", bufferStatus.LogBufferSize}
    bufferStatus.EntityData.Leafs["max-log-buffer-size"] = types.YLeaf{"MaxLogBufferSize", bufferStatus.MaxLogBufferSize}
    bufferStatus.EntityData.Leafs["record-count"] = types.YLeaf{"RecordCount", bufferStatus.RecordCount}
    bufferStatus.EntityData.Leafs["capacity-threshold"] = types.YLeaf{"CapacityThreshold", bufferStatus.CapacityThreshold}
    bufferStatus.EntityData.Leafs["severity-filter"] = types.YLeaf{"SeverityFilter", bufferStatus.SeverityFilter}
    return &(bufferStatus.EntityData)
}

// AlarmLogger_Alarms
// Table that contains the database of logged
// alarms
type AlarmLogger_Alarms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the logged alarms. The type is slice of AlarmLogger_Alarms_Alarm.
    Alarm []AlarmLogger_Alarms_Alarm
}

func (alarms *AlarmLogger_Alarms) GetEntityData() *types.CommonEntityData {
    alarms.EntityData.YFilter = alarms.YFilter
    alarms.EntityData.YangName = "alarms"
    alarms.EntityData.BundleName = "cisco_ios_xr"
    alarms.EntityData.ParentYangName = "alarm-logger"
    alarms.EntityData.SegmentPath = "alarms"
    alarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarms.EntityData.Children = make(map[string]types.YChild)
    alarms.EntityData.Children["alarm"] = types.YChild{"Alarm", nil}
    for i := range alarms.Alarm {
        alarms.EntityData.Children[types.GetSegmentPath(&alarms.Alarm[i])] = types.YChild{"Alarm", &alarms.Alarm[i]}
    }
    alarms.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alarms.EntityData)
}

// AlarmLogger_Alarms_Alarm
// One of the logged alarms
type AlarmLogger_Alarms_Alarm struct {
    EntityData types.CommonEntityData
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

func (alarm *AlarmLogger_Alarms_Alarm) GetEntityData() *types.CommonEntityData {
    alarm.EntityData.YFilter = alarm.YFilter
    alarm.EntityData.YangName = "alarm"
    alarm.EntityData.BundleName = "cisco_ios_xr"
    alarm.EntityData.ParentYangName = "alarms"
    alarm.EntityData.SegmentPath = "alarm" + "[event-id='" + fmt.Sprintf("%v", alarm.EventId) + "']"
    alarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarm.EntityData.Children = make(map[string]types.YChild)
    alarm.EntityData.Leafs = make(map[string]types.YLeaf)
    alarm.EntityData.Leafs["event-id"] = types.YLeaf{"EventId", alarm.EventId}
    alarm.EntityData.Leafs["source-id"] = types.YLeaf{"SourceId", alarm.SourceId}
    alarm.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", alarm.Timestamp}
    alarm.EntityData.Leafs["category"] = types.YLeaf{"Category", alarm.Category}
    alarm.EntityData.Leafs["group"] = types.YLeaf{"Group", alarm.Group}
    alarm.EntityData.Leafs["code"] = types.YLeaf{"Code", alarm.Code}
    alarm.EntityData.Leafs["severity"] = types.YLeaf{"Severity", alarm.Severity}
    alarm.EntityData.Leafs["state"] = types.YLeaf{"State", alarm.State}
    alarm.EntityData.Leafs["correlation-id"] = types.YLeaf{"CorrelationId", alarm.CorrelationId}
    alarm.EntityData.Leafs["is-admin"] = types.YLeaf{"IsAdmin", alarm.IsAdmin}
    alarm.EntityData.Leafs["additional-text"] = types.YLeaf{"AdditionalText", alarm.AdditionalText}
    return &(alarm.EntityData)
}

