// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-syslog package configuration.
// 
// This module contains definitions
// for the following management objects:
//   syslog-service: Syslog Timestamp Services
//   syslog: syslog
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_syslog_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_syslog_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-syslog-cfg syslog-service}", reflect.TypeOf(SyslogService{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-syslog-cfg:syslog-service", reflect.TypeOf(SyslogService{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-syslog-cfg syslog}", reflect.TypeOf(Syslog{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-syslog-cfg:syslog", reflect.TypeOf(Syslog{}))
}

// LogSeverity represents Log severity
type LogSeverity string

const (
    // System is unusable                (severity=0)
    LogSeverity_emergency LogSeverity = "emergency"

    // Immediate action needed           (severity=1)
    LogSeverity_alert LogSeverity = "alert"

    // Critical conditions               (severity=2)
    LogSeverity_critical LogSeverity = "critical"

    // Error conditions                  (severity=3)
    LogSeverity_error LogSeverity = "error"

    // Warning conditions                (severity=4)
    LogSeverity_warning LogSeverity = "warning"

    // Normal but significant conditions (severity=5)
    LogSeverity_notice LogSeverity = "notice"

    // Informational messages            (severity=6)
    LogSeverity_informational LogSeverity = "informational"

    // Debugging messages                (severity=7)
    LogSeverity_debug LogSeverity = "debug"
)

// Facility represents Facility
type Facility string

const (
    // Kernel Facility
    Facility_kern Facility = "kern"

    // User Facility
    Facility_user Facility = "user"

    // Mail Facility
    Facility_mail Facility = "mail"

    // Daemon Facility
    Facility_daemon Facility = "daemon"

    // Auth Facility
    Facility_auth Facility = "auth"

    // Syslog Facility
    Facility_syslog Facility = "syslog"

    // Lpr Facility
    Facility_lpr Facility = "lpr"

    // News Facility
    Facility_news Facility = "news"

    // Uucp Facility
    Facility_uucp Facility = "uucp"

    // Cron Facility
    Facility_cron Facility = "cron"

    // Authpriv Facility
    Facility_authpriv Facility = "authpriv"

    // Ftp Facility
    Facility_ftp Facility = "ftp"

    // Local0 Facility
    Facility_local0 Facility = "local0"

    // Local1 Facility
    Facility_local1 Facility = "local1"

    // Local2 Facility
    Facility_local2 Facility = "local2"

    // Local3 Facility
    Facility_local3 Facility = "local3"

    // Local4 Facility
    Facility_local4 Facility = "local4"

    // Local5 Facility
    Facility_local5 Facility = "local5"

    // Local6 Facility
    Facility_local6 Facility = "local6"

    // Local7 Facility
    Facility_local7 Facility = "local7"

    // System9 Facility
    Facility_sys9 Facility = "sys9"

    // System10 Facility
    Facility_sys10 Facility = "sys10"

    // System11 Facility
    Facility_sys11 Facility = "sys11"

    // System12 Facility
    Facility_sys12 Facility = "sys12"

    // System13 Facility
    Facility_sys13 Facility = "sys13"

    // System14 Facility
    Facility_sys14 Facility = "sys14"
)

// LogCollectFrequency represents Log collect frequency
type LogCollectFrequency string

const (
    // Collect log in files on a weekly basis
    LogCollectFrequency_weekly LogCollectFrequency = "weekly"

    // Collect log in files on a daily basis
    LogCollectFrequency_daily LogCollectFrequency = "daily"
)

// LoggingPrecedenceValue represents Logging precedence value
type LoggingPrecedenceValue string

const (
    // Applicable to precedence: value 0
    LoggingPrecedenceValue_routine LoggingPrecedenceValue = "routine"

    // Applicable to precedence: value 1
    LoggingPrecedenceValue_priority LoggingPrecedenceValue = "priority"

    // Applicable to precedence: value 2
    LoggingPrecedenceValue_immediate LoggingPrecedenceValue = "immediate"

    // Applicable to precedence: value 3
    LoggingPrecedenceValue_flash LoggingPrecedenceValue = "flash"

    // Applicable to precedence: value 4
    LoggingPrecedenceValue_flash_override LoggingPrecedenceValue = "flash-override"

    // Applicable to precedence: value 5
    LoggingPrecedenceValue_critical LoggingPrecedenceValue = "critical"

    // Applicable to precedence: value 6
    LoggingPrecedenceValue_internet LoggingPrecedenceValue = "internet"

    // Applicable to precedence: value 7
    LoggingPrecedenceValue_network LoggingPrecedenceValue = "network"
)

// LoggingTos represents Logging tos
type LoggingTos string

const (
    // Logging TOS type precedence
    LoggingTos_precedence LoggingTos = "precedence"

    // Logging TOS type DSCP
    LoggingTos_dscp LoggingTos = "dscp"
)

// LoggingLevels represents Logging levels
type LoggingLevels string

const (
    // Emergency Level Msg
    LoggingLevels_emergency LoggingLevels = "emergency"

    // Alert Level Msg
    LoggingLevels_alert LoggingLevels = "alert"

    // Critical Level Msg
    LoggingLevels_critical LoggingLevels = "critical"

    // Error Level Msg
    LoggingLevels_error LoggingLevels = "error"

    // Warning Level Msg
    LoggingLevels_warning LoggingLevels = "warning"

    // Notification Level Msg
    LoggingLevels_notice LoggingLevels = "notice"

    // Informational Level Msg
    LoggingLevels_info LoggingLevels = "info"

    // Debugging Level Msg
    LoggingLevels_debug LoggingLevels = "debug"

    // Disable logging
    LoggingLevels_disable LoggingLevels = "disable"
)

// LoggingPrecedence represents Logging precedence
type LoggingPrecedence string

const (
    // Logging TOS type precedence
    LoggingPrecedence_precedence LoggingPrecedence = "precedence"
)

// LoggingDscpValue represents Logging dscp value
type LoggingDscpValue string

const (
    // Applicable to DSCP: bits 000000
    LoggingDscpValue_default_ LoggingDscpValue = "default"

    // Applicable to DSCP: bits 001010
    LoggingDscpValue_af11 LoggingDscpValue = "af11"

    // Applicable to DSCP: bits 001100
    LoggingDscpValue_af12 LoggingDscpValue = "af12"

    // Applicable to DSCP: bits 001110
    LoggingDscpValue_af13 LoggingDscpValue = "af13"

    // Applicable to DSCP: bits 010010
    LoggingDscpValue_af21 LoggingDscpValue = "af21"

    // Applicable to DSCP: bits 010100
    LoggingDscpValue_af22 LoggingDscpValue = "af22"

    // Applicable to DSCP: bits 010110
    LoggingDscpValue_af23 LoggingDscpValue = "af23"

    // Applicable to DSCP: bits 011010
    LoggingDscpValue_af31 LoggingDscpValue = "af31"

    // Applicable to DSCP: bits 011100
    LoggingDscpValue_af32 LoggingDscpValue = "af32"

    // Applicable to DSCP: bits 011110
    LoggingDscpValue_af33 LoggingDscpValue = "af33"

    // Applicable to DSCP: bits 100010
    LoggingDscpValue_af41 LoggingDscpValue = "af41"

    // Applicable to DSCP: bits 100100
    LoggingDscpValue_af42 LoggingDscpValue = "af42"

    // Applicable to DSCP: bits 100110
    LoggingDscpValue_af43 LoggingDscpValue = "af43"

    // Applicable to DSCP: bits 101110
    LoggingDscpValue_ef LoggingDscpValue = "ef"

    // Applicable to DSCP: bits 001000
    LoggingDscpValue_cs1 LoggingDscpValue = "cs1"

    // Applicable to DSCP: bits 010000
    LoggingDscpValue_cs2 LoggingDscpValue = "cs2"

    // Applicable to DSCP: bits 011000
    LoggingDscpValue_cs3 LoggingDscpValue = "cs3"

    // Applicable to DSCP: bits 100000
    LoggingDscpValue_cs4 LoggingDscpValue = "cs4"

    // Applicable to DSCP: bits 101000
    LoggingDscpValue_cs5 LoggingDscpValue = "cs5"

    // Applicable to DSCP: bits 110000
    LoggingDscpValue_cs6 LoggingDscpValue = "cs6"

    // Applicable to DSCP: bits 111000
    LoggingDscpValue_cs7 LoggingDscpValue = "cs7"
)

// LogMessageSeverity represents Log message severity
type LogMessageSeverity string

const (
    // System is unusable                (severity=0)
    LogMessageSeverity_emergency LogMessageSeverity = "emergency"

    // Immediate action needed           (severity=1)
    LogMessageSeverity_alert LogMessageSeverity = "alert"

    // Critical conditions               (severity=2)
    LogMessageSeverity_critical LogMessageSeverity = "critical"

    // Error conditions                  (severity=3)
    LogMessageSeverity_error LogMessageSeverity = "error"

    // Warning conditions                (severity=4)
    LogMessageSeverity_warning LogMessageSeverity = "warning"

    // Normal but significant conditions (severity=5)
    LogMessageSeverity_notice LogMessageSeverity = "notice"

    // Informational messages            (severity=6)
    LogMessageSeverity_informational LogMessageSeverity = "informational"

    // Debugging messages                (severity=7)
    LogMessageSeverity_debug LogMessageSeverity = "debug"
)

// TimeInfo represents Time info
type TimeInfo string

const (
    // Exclude
    TimeInfo_disable TimeInfo = "disable"

    // Include
    TimeInfo_enable TimeInfo = "enable"
)

// LoggingDscp represents Logging dscp
type LoggingDscp string

const (
    // Logging TOS type DSCP
    LoggingDscp_dscp LoggingDscp = "dscp"
)

// SyslogService
// Syslog Timestamp Services
type SyslogService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timestamp debug/log messages configuration.
    Timestamps SyslogService_Timestamps
}

func (syslogService *SyslogService) GetFilter() yfilter.YFilter { return syslogService.YFilter }

func (syslogService *SyslogService) SetFilter(yf yfilter.YFilter) { syslogService.YFilter = yf }

func (syslogService *SyslogService) GetGoName(yname string) string {
    if yname == "timestamps" { return "Timestamps" }
    return ""
}

func (syslogService *SyslogService) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-syslog-cfg:syslog-service"
}

func (syslogService *SyslogService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timestamps" {
        return &syslogService.Timestamps
    }
    return nil
}

func (syslogService *SyslogService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["timestamps"] = &syslogService.Timestamps
    return children
}

func (syslogService *SyslogService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (syslogService *SyslogService) GetBundleName() string { return "cisco_ios_xr" }

func (syslogService *SyslogService) GetYangName() string { return "syslog-service" }

func (syslogService *SyslogService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syslogService *SyslogService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syslogService *SyslogService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syslogService *SyslogService) SetParent(parent types.Entity) { syslogService.parent = parent }

func (syslogService *SyslogService) GetParent() types.Entity { return syslogService.parent }

func (syslogService *SyslogService) GetParentYangName() string { return "Cisco-IOS-XR-infra-syslog-cfg" }

// SyslogService_Timestamps
// Timestamp debug/log messages configuration
type SyslogService_Timestamps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable timestamp debug/log messages. The type is interface{}.
    Enable interface{}

    // Timestamp log messages.
    Log SyslogService_Timestamps_Log

    // Timestamp debug messages.
    Debug SyslogService_Timestamps_Debug
}

func (timestamps *SyslogService_Timestamps) GetFilter() yfilter.YFilter { return timestamps.YFilter }

func (timestamps *SyslogService_Timestamps) SetFilter(yf yfilter.YFilter) { timestamps.YFilter = yf }

func (timestamps *SyslogService_Timestamps) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "log" { return "Log" }
    if yname == "debug" { return "Debug" }
    return ""
}

func (timestamps *SyslogService_Timestamps) GetSegmentPath() string {
    return "timestamps"
}

func (timestamps *SyslogService_Timestamps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log" {
        return &timestamps.Log
    }
    if childYangName == "debug" {
        return &timestamps.Debug
    }
    return nil
}

func (timestamps *SyslogService_Timestamps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log"] = &timestamps.Log
    children["debug"] = &timestamps.Debug
    return children
}

func (timestamps *SyslogService_Timestamps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = timestamps.Enable
    return leafs
}

func (timestamps *SyslogService_Timestamps) GetBundleName() string { return "cisco_ios_xr" }

func (timestamps *SyslogService_Timestamps) GetYangName() string { return "timestamps" }

func (timestamps *SyslogService_Timestamps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timestamps *SyslogService_Timestamps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timestamps *SyslogService_Timestamps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timestamps *SyslogService_Timestamps) SetParent(parent types.Entity) { timestamps.parent = parent }

func (timestamps *SyslogService_Timestamps) GetParent() types.Entity { return timestamps.parent }

func (timestamps *SyslogService_Timestamps) GetParentYangName() string { return "syslog-service" }

// SyslogService_Timestamps_Log
// Timestamp log messages
type SyslogService_Timestamps_Log struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timestamp with systime uptime. The type is interface{}.
    LogUptime interface{}

    // Disable timestamp log messages. The type is interface{}.
    LogTimestampDisable interface{}

    // Timestamp with date and time.
    LogDatetime SyslogService_Timestamps_Log_LogDatetime
}

func (log *SyslogService_Timestamps_Log) GetFilter() yfilter.YFilter { return log.YFilter }

func (log *SyslogService_Timestamps_Log) SetFilter(yf yfilter.YFilter) { log.YFilter = yf }

func (log *SyslogService_Timestamps_Log) GetGoName(yname string) string {
    if yname == "log-uptime" { return "LogUptime" }
    if yname == "log-timestamp-disable" { return "LogTimestampDisable" }
    if yname == "log-datetime" { return "LogDatetime" }
    return ""
}

func (log *SyslogService_Timestamps_Log) GetSegmentPath() string {
    return "log"
}

func (log *SyslogService_Timestamps_Log) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-datetime" {
        return &log.LogDatetime
    }
    return nil
}

func (log *SyslogService_Timestamps_Log) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-datetime"] = &log.LogDatetime
    return children
}

func (log *SyslogService_Timestamps_Log) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-uptime"] = log.LogUptime
    leafs["log-timestamp-disable"] = log.LogTimestampDisable
    return leafs
}

func (log *SyslogService_Timestamps_Log) GetBundleName() string { return "cisco_ios_xr" }

func (log *SyslogService_Timestamps_Log) GetYangName() string { return "log" }

func (log *SyslogService_Timestamps_Log) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (log *SyslogService_Timestamps_Log) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (log *SyslogService_Timestamps_Log) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (log *SyslogService_Timestamps_Log) SetParent(parent types.Entity) { log.parent = parent }

func (log *SyslogService_Timestamps_Log) GetParent() types.Entity { return log.parent }

func (log *SyslogService_Timestamps_Log) GetParentYangName() string { return "timestamps" }

// SyslogService_Timestamps_Log_LogDatetime
// Timestamp with date and time
type SyslogService_Timestamps_Log_LogDatetime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set timestamp for log message.
    LogDatetimeValue SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue
}

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetFilter() yfilter.YFilter { return logDatetime.YFilter }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) SetFilter(yf yfilter.YFilter) { logDatetime.YFilter = yf }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetGoName(yname string) string {
    if yname == "log-datetime-value" { return "LogDatetimeValue" }
    return ""
}

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetSegmentPath() string {
    return "log-datetime"
}

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-datetime-value" {
        return &logDatetime.LogDatetimeValue
    }
    return nil
}

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-datetime-value"] = &logDatetime.LogDatetimeValue
    return children
}

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetBundleName() string { return "cisco_ios_xr" }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetYangName() string { return "log-datetime" }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) SetParent(parent types.Entity) { logDatetime.parent = parent }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetParent() types.Entity { return logDatetime.parent }

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetParentYangName() string { return "log" }

// SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue
// Set timestamp for log message
type SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time. The type is TimeInfo. The default value is enable.
    TimeStampValue interface{}

    // Seconds. The type is TimeInfo. Units are second. The default value is
    // enable.
    Msec interface{}

    // Timezone. The type is TimeInfo. The default value is disable.
    TimeZone interface{}

    // Year. The type is TimeInfo. The default value is disable.
    Year interface{}
}

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetFilter() yfilter.YFilter { return logDatetimeValue.YFilter }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) SetFilter(yf yfilter.YFilter) { logDatetimeValue.YFilter = yf }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetGoName(yname string) string {
    if yname == "time-stamp-value" { return "TimeStampValue" }
    if yname == "msec" { return "Msec" }
    if yname == "time-zone" { return "TimeZone" }
    if yname == "year" { return "Year" }
    return ""
}

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetSegmentPath() string {
    return "log-datetime-value"
}

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-stamp-value"] = logDatetimeValue.TimeStampValue
    leafs["msec"] = logDatetimeValue.Msec
    leafs["time-zone"] = logDatetimeValue.TimeZone
    leafs["year"] = logDatetimeValue.Year
    return leafs
}

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetBundleName() string { return "cisco_ios_xr" }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetYangName() string { return "log-datetime-value" }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) SetParent(parent types.Entity) { logDatetimeValue.parent = parent }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetParent() types.Entity { return logDatetimeValue.parent }

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetParentYangName() string { return "log-datetime" }

// SyslogService_Timestamps_Debug
// Timestamp debug messages
type SyslogService_Timestamps_Debug struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable timestamp debug messages. The type is interface{}.
    DebugTimestampDisable interface{}

    // Timestamp with systime uptime. The type is interface{}.
    DebugUptime interface{}

    // Timestamp with date and time.
    DebugDatetime SyslogService_Timestamps_Debug_DebugDatetime
}

func (debug *SyslogService_Timestamps_Debug) GetFilter() yfilter.YFilter { return debug.YFilter }

func (debug *SyslogService_Timestamps_Debug) SetFilter(yf yfilter.YFilter) { debug.YFilter = yf }

func (debug *SyslogService_Timestamps_Debug) GetGoName(yname string) string {
    if yname == "debug-timestamp-disable" { return "DebugTimestampDisable" }
    if yname == "debug-uptime" { return "DebugUptime" }
    if yname == "debug-datetime" { return "DebugDatetime" }
    return ""
}

func (debug *SyslogService_Timestamps_Debug) GetSegmentPath() string {
    return "debug"
}

func (debug *SyslogService_Timestamps_Debug) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "debug-datetime" {
        return &debug.DebugDatetime
    }
    return nil
}

func (debug *SyslogService_Timestamps_Debug) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["debug-datetime"] = &debug.DebugDatetime
    return children
}

func (debug *SyslogService_Timestamps_Debug) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["debug-timestamp-disable"] = debug.DebugTimestampDisable
    leafs["debug-uptime"] = debug.DebugUptime
    return leafs
}

func (debug *SyslogService_Timestamps_Debug) GetBundleName() string { return "cisco_ios_xr" }

func (debug *SyslogService_Timestamps_Debug) GetYangName() string { return "debug" }

func (debug *SyslogService_Timestamps_Debug) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (debug *SyslogService_Timestamps_Debug) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (debug *SyslogService_Timestamps_Debug) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (debug *SyslogService_Timestamps_Debug) SetParent(parent types.Entity) { debug.parent = parent }

func (debug *SyslogService_Timestamps_Debug) GetParent() types.Entity { return debug.parent }

func (debug *SyslogService_Timestamps_Debug) GetParentYangName() string { return "timestamps" }

// SyslogService_Timestamps_Debug_DebugDatetime
// Timestamp with date and time
type SyslogService_Timestamps_Debug_DebugDatetime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set time format for debug msg.
    DatetimeValue SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue
}

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetFilter() yfilter.YFilter { return debugDatetime.YFilter }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) SetFilter(yf yfilter.YFilter) { debugDatetime.YFilter = yf }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetGoName(yname string) string {
    if yname == "datetime-value" { return "DatetimeValue" }
    return ""
}

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetSegmentPath() string {
    return "debug-datetime"
}

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "datetime-value" {
        return &debugDatetime.DatetimeValue
    }
    return nil
}

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["datetime-value"] = &debugDatetime.DatetimeValue
    return children
}

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetBundleName() string { return "cisco_ios_xr" }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetYangName() string { return "debug-datetime" }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) SetParent(parent types.Entity) { debugDatetime.parent = parent }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetParent() types.Entity { return debugDatetime.parent }

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetParentYangName() string { return "debug" }

// SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue
// Set time format for debug msg
type SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time. The type is TimeInfo. The default value is enable.
    TimeStampValue interface{}

    // Seconds. The type is TimeInfo. Units are second. The default value is
    // enable.
    Msec interface{}

    // Timezone. The type is TimeInfo. The default value is disable.
    TimeZone interface{}

    // Year. The type is TimeInfo. The default value is disable.
    Year interface{}
}

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetFilter() yfilter.YFilter { return datetimeValue.YFilter }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) SetFilter(yf yfilter.YFilter) { datetimeValue.YFilter = yf }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetGoName(yname string) string {
    if yname == "time-stamp-value" { return "TimeStampValue" }
    if yname == "msec" { return "Msec" }
    if yname == "time-zone" { return "TimeZone" }
    if yname == "year" { return "Year" }
    return ""
}

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetSegmentPath() string {
    return "datetime-value"
}

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-stamp-value"] = datetimeValue.TimeStampValue
    leafs["msec"] = datetimeValue.Msec
    leafs["time-zone"] = datetimeValue.TimeZone
    leafs["year"] = datetimeValue.Year
    return leafs
}

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetBundleName() string { return "cisco_ios_xr" }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetYangName() string { return "datetime-value" }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) SetParent(parent types.Entity) { datetimeValue.parent = parent }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetParent() types.Entity { return datetimeValue.parent }

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetParentYangName() string { return "debug-datetime" }

// Syslog
// syslog
type Syslog struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hostname prefix to add on msgs to servers. The type is string.
    HostNamePrefix interface{}

    // Set size of the local log file. The type is interface{} with range:
    // 0..4294967295. The default value is 32768.
    LocalLogFileSize interface{}

    // Enabled or disabled. The type is bool.
    EnableConsoleLogging interface{}

    // Suppress consecutive duplicate messages. The type is interface{}.
    SuppressDuplicates interface{}

    // Set monitor logging.
    MonitorLogging Syslog_MonitorLogging

    // Set history logging.
    HistoryLogging Syslog_HistoryLogging

    // Modify message logging facilities.
    LoggingFacilities Syslog_LoggingFacilities

    // Set trap logging.
    TrapLogging Syslog_TrapLogging

    // Set buffered logging parameters.
    BufferedLogging Syslog_BufferedLogging

    // Configure logging host.
    HostServer Syslog_HostServer

    // Set console logging.
    ConsoleLogging Syslog_ConsoleLogging

    // Configure logging file destination.
    Files Syslog_Files

    // Syslog TOS bit for outgoing messages.
    Ipv4 Syslog_Ipv4

    // Archive attributes configuration.
    Archive Syslog_Archive

    // Syslog traffic class bit for outgoing messages.
    Ipv6 Syslog_Ipv6

    // Configure source interface.
    SourceInterfaceTable Syslog_SourceInterfaceTable

    // Alarm Logger Properties.
    AlarmLogger Syslog_AlarmLogger

    // Configure properties of the event correlator.
    Correlator Syslog_Correlator

    // Configure properties of the syslog/alarm suppression.
    Suppression Syslog_Suppression
}

func (syslog *Syslog) GetFilter() yfilter.YFilter { return syslog.YFilter }

func (syslog *Syslog) SetFilter(yf yfilter.YFilter) { syslog.YFilter = yf }

func (syslog *Syslog) GetGoName(yname string) string {
    if yname == "host-name-prefix" { return "HostNamePrefix" }
    if yname == "local-log-file-size" { return "LocalLogFileSize" }
    if yname == "enable-console-logging" { return "EnableConsoleLogging" }
    if yname == "suppress-duplicates" { return "SuppressDuplicates" }
    if yname == "monitor-logging" { return "MonitorLogging" }
    if yname == "history-logging" { return "HistoryLogging" }
    if yname == "logging-facilities" { return "LoggingFacilities" }
    if yname == "trap-logging" { return "TrapLogging" }
    if yname == "buffered-logging" { return "BufferedLogging" }
    if yname == "host-server" { return "HostServer" }
    if yname == "console-logging" { return "ConsoleLogging" }
    if yname == "files" { return "Files" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "archive" { return "Archive" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "source-interface-table" { return "SourceInterfaceTable" }
    if yname == "Cisco-IOS-XR-infra-alarm-logger-cfg:alarm-logger" { return "AlarmLogger" }
    if yname == "Cisco-IOS-XR-infra-correlator-cfg:correlator" { return "Correlator" }
    if yname == "Cisco-IOS-XR-infra-correlator-cfg:suppression" { return "Suppression" }
    return ""
}

func (syslog *Syslog) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-syslog-cfg:syslog"
}

func (syslog *Syslog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "monitor-logging" {
        return &syslog.MonitorLogging
    }
    if childYangName == "history-logging" {
        return &syslog.HistoryLogging
    }
    if childYangName == "logging-facilities" {
        return &syslog.LoggingFacilities
    }
    if childYangName == "trap-logging" {
        return &syslog.TrapLogging
    }
    if childYangName == "buffered-logging" {
        return &syslog.BufferedLogging
    }
    if childYangName == "host-server" {
        return &syslog.HostServer
    }
    if childYangName == "console-logging" {
        return &syslog.ConsoleLogging
    }
    if childYangName == "files" {
        return &syslog.Files
    }
    if childYangName == "ipv4" {
        return &syslog.Ipv4
    }
    if childYangName == "archive" {
        return &syslog.Archive
    }
    if childYangName == "ipv6" {
        return &syslog.Ipv6
    }
    if childYangName == "source-interface-table" {
        return &syslog.SourceInterfaceTable
    }
    if childYangName == "Cisco-IOS-XR-infra-alarm-logger-cfg:alarm-logger" {
        return &syslog.AlarmLogger
    }
    if childYangName == "Cisco-IOS-XR-infra-correlator-cfg:correlator" {
        return &syslog.Correlator
    }
    if childYangName == "Cisco-IOS-XR-infra-correlator-cfg:suppression" {
        return &syslog.Suppression
    }
    return nil
}

func (syslog *Syslog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["monitor-logging"] = &syslog.MonitorLogging
    children["history-logging"] = &syslog.HistoryLogging
    children["logging-facilities"] = &syslog.LoggingFacilities
    children["trap-logging"] = &syslog.TrapLogging
    children["buffered-logging"] = &syslog.BufferedLogging
    children["host-server"] = &syslog.HostServer
    children["console-logging"] = &syslog.ConsoleLogging
    children["files"] = &syslog.Files
    children["ipv4"] = &syslog.Ipv4
    children["archive"] = &syslog.Archive
    children["ipv6"] = &syslog.Ipv6
    children["source-interface-table"] = &syslog.SourceInterfaceTable
    children["Cisco-IOS-XR-infra-alarm-logger-cfg:alarm-logger"] = &syslog.AlarmLogger
    children["Cisco-IOS-XR-infra-correlator-cfg:correlator"] = &syslog.Correlator
    children["Cisco-IOS-XR-infra-correlator-cfg:suppression"] = &syslog.Suppression
    return children
}

func (syslog *Syslog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name-prefix"] = syslog.HostNamePrefix
    leafs["local-log-file-size"] = syslog.LocalLogFileSize
    leafs["enable-console-logging"] = syslog.EnableConsoleLogging
    leafs["suppress-duplicates"] = syslog.SuppressDuplicates
    return leafs
}

func (syslog *Syslog) GetBundleName() string { return "cisco_ios_xr" }

func (syslog *Syslog) GetYangName() string { return "syslog" }

func (syslog *Syslog) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syslog *Syslog) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syslog *Syslog) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syslog *Syslog) SetParent(parent types.Entity) { syslog.parent = parent }

func (syslog *Syslog) GetParent() types.Entity { return syslog.parent }

func (syslog *Syslog) GetParentYangName() string { return "Cisco-IOS-XR-infra-syslog-cfg" }

// Syslog_MonitorLogging
// Set monitor logging
type Syslog_MonitorLogging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor Logging Level. The type is LoggingLevels. The default value is
    // debug.
    LoggingLevel interface{}

    // Set monitor logging discriminators.
    MonitorDiscriminator Syslog_MonitorLogging_MonitorDiscriminator
}

func (monitorLogging *Syslog_MonitorLogging) GetFilter() yfilter.YFilter { return monitorLogging.YFilter }

func (monitorLogging *Syslog_MonitorLogging) SetFilter(yf yfilter.YFilter) { monitorLogging.YFilter = yf }

func (monitorLogging *Syslog_MonitorLogging) GetGoName(yname string) string {
    if yname == "logging-level" { return "LoggingLevel" }
    if yname == "monitor-discriminator" { return "MonitorDiscriminator" }
    return ""
}

func (monitorLogging *Syslog_MonitorLogging) GetSegmentPath() string {
    return "monitor-logging"
}

func (monitorLogging *Syslog_MonitorLogging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "monitor-discriminator" {
        return &monitorLogging.MonitorDiscriminator
    }
    return nil
}

func (monitorLogging *Syslog_MonitorLogging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["monitor-discriminator"] = &monitorLogging.MonitorDiscriminator
    return children
}

func (monitorLogging *Syslog_MonitorLogging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging-level"] = monitorLogging.LoggingLevel
    return leafs
}

func (monitorLogging *Syslog_MonitorLogging) GetBundleName() string { return "cisco_ios_xr" }

func (monitorLogging *Syslog_MonitorLogging) GetYangName() string { return "monitor-logging" }

func (monitorLogging *Syslog_MonitorLogging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (monitorLogging *Syslog_MonitorLogging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (monitorLogging *Syslog_MonitorLogging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (monitorLogging *Syslog_MonitorLogging) SetParent(parent types.Entity) { monitorLogging.parent = parent }

func (monitorLogging *Syslog_MonitorLogging) GetParent() types.Entity { return monitorLogging.parent }

func (monitorLogging *Syslog_MonitorLogging) GetParentYangName() string { return "syslog" }

// Syslog_MonitorLogging_MonitorDiscriminator
// Set monitor logging discriminators
type Syslog_MonitorLogging_MonitorDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set monitor logging match2 discriminator. The type is string.
    Match2 interface{}

    // Set monitor logging no-match1 discriminator. The type is string.
    Nomatch1 interface{}

    // Set monitor logging match1 discriminator. The type is string.
    Match1 interface{}

    // Set monitor logging no-match3 discriminator. The type is string.
    Nomatch3 interface{}

    // Set monitor logging match3 discriminator. The type is string.
    Match3 interface{}

    // Set monitor logging no-match2 discriminator. The type is string.
    Nomatch2 interface{}
}

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetFilter() yfilter.YFilter { return monitorDiscriminator.YFilter }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) SetFilter(yf yfilter.YFilter) { monitorDiscriminator.YFilter = yf }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetGoName(yname string) string {
    if yname == "match2" { return "Match2" }
    if yname == "nomatch1" { return "Nomatch1" }
    if yname == "match1" { return "Match1" }
    if yname == "nomatch3" { return "Nomatch3" }
    if yname == "match3" { return "Match3" }
    if yname == "nomatch2" { return "Nomatch2" }
    return ""
}

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetSegmentPath() string {
    return "monitor-discriminator"
}

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["match2"] = monitorDiscriminator.Match2
    leafs["nomatch1"] = monitorDiscriminator.Nomatch1
    leafs["match1"] = monitorDiscriminator.Match1
    leafs["nomatch3"] = monitorDiscriminator.Nomatch3
    leafs["match3"] = monitorDiscriminator.Match3
    leafs["nomatch2"] = monitorDiscriminator.Nomatch2
    return leafs
}

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetYangName() string { return "monitor-discriminator" }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) SetParent(parent types.Entity) { monitorDiscriminator.parent = parent }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetParent() types.Entity { return monitorDiscriminator.parent }

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetParentYangName() string { return "monitor-logging" }

// Syslog_HistoryLogging
// Set history logging
type Syslog_HistoryLogging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging history size. The type is interface{} with range: 1..500. The
    // default value is 1.
    HistorySize interface{}

    // History logging level. The type is LoggingLevels. The default value is
    // warning.
    LoggingLevel interface{}
}

func (historyLogging *Syslog_HistoryLogging) GetFilter() yfilter.YFilter { return historyLogging.YFilter }

func (historyLogging *Syslog_HistoryLogging) SetFilter(yf yfilter.YFilter) { historyLogging.YFilter = yf }

func (historyLogging *Syslog_HistoryLogging) GetGoName(yname string) string {
    if yname == "history-size" { return "HistorySize" }
    if yname == "logging-level" { return "LoggingLevel" }
    return ""
}

func (historyLogging *Syslog_HistoryLogging) GetSegmentPath() string {
    return "history-logging"
}

func (historyLogging *Syslog_HistoryLogging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (historyLogging *Syslog_HistoryLogging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (historyLogging *Syslog_HistoryLogging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["history-size"] = historyLogging.HistorySize
    leafs["logging-level"] = historyLogging.LoggingLevel
    return leafs
}

func (historyLogging *Syslog_HistoryLogging) GetBundleName() string { return "cisco_ios_xr" }

func (historyLogging *Syslog_HistoryLogging) GetYangName() string { return "history-logging" }

func (historyLogging *Syslog_HistoryLogging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (historyLogging *Syslog_HistoryLogging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (historyLogging *Syslog_HistoryLogging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (historyLogging *Syslog_HistoryLogging) SetParent(parent types.Entity) { historyLogging.parent = parent }

func (historyLogging *Syslog_HistoryLogging) GetParent() types.Entity { return historyLogging.parent }

func (historyLogging *Syslog_HistoryLogging) GetParentYangName() string { return "syslog" }

// Syslog_LoggingFacilities
// Modify message logging facilities
type Syslog_LoggingFacilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Facility from which logging is done. The type is Facility. The default
    // value is local7.
    FacilityLevel interface{}
}

func (loggingFacilities *Syslog_LoggingFacilities) GetFilter() yfilter.YFilter { return loggingFacilities.YFilter }

func (loggingFacilities *Syslog_LoggingFacilities) SetFilter(yf yfilter.YFilter) { loggingFacilities.YFilter = yf }

func (loggingFacilities *Syslog_LoggingFacilities) GetGoName(yname string) string {
    if yname == "facility-level" { return "FacilityLevel" }
    return ""
}

func (loggingFacilities *Syslog_LoggingFacilities) GetSegmentPath() string {
    return "logging-facilities"
}

func (loggingFacilities *Syslog_LoggingFacilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (loggingFacilities *Syslog_LoggingFacilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (loggingFacilities *Syslog_LoggingFacilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["facility-level"] = loggingFacilities.FacilityLevel
    return leafs
}

func (loggingFacilities *Syslog_LoggingFacilities) GetBundleName() string { return "cisco_ios_xr" }

func (loggingFacilities *Syslog_LoggingFacilities) GetYangName() string { return "logging-facilities" }

func (loggingFacilities *Syslog_LoggingFacilities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loggingFacilities *Syslog_LoggingFacilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loggingFacilities *Syslog_LoggingFacilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loggingFacilities *Syslog_LoggingFacilities) SetParent(parent types.Entity) { loggingFacilities.parent = parent }

func (loggingFacilities *Syslog_LoggingFacilities) GetParent() types.Entity { return loggingFacilities.parent }

func (loggingFacilities *Syslog_LoggingFacilities) GetParentYangName() string { return "syslog" }

// Syslog_TrapLogging
// Set trap logging
type Syslog_TrapLogging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trap logging level. The type is LoggingLevels. The default value is info.
    LoggingLevel interface{}
}

func (trapLogging *Syslog_TrapLogging) GetFilter() yfilter.YFilter { return trapLogging.YFilter }

func (trapLogging *Syslog_TrapLogging) SetFilter(yf yfilter.YFilter) { trapLogging.YFilter = yf }

func (trapLogging *Syslog_TrapLogging) GetGoName(yname string) string {
    if yname == "logging-level" { return "LoggingLevel" }
    return ""
}

func (trapLogging *Syslog_TrapLogging) GetSegmentPath() string {
    return "trap-logging"
}

func (trapLogging *Syslog_TrapLogging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapLogging *Syslog_TrapLogging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapLogging *Syslog_TrapLogging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging-level"] = trapLogging.LoggingLevel
    return leafs
}

func (trapLogging *Syslog_TrapLogging) GetBundleName() string { return "cisco_ios_xr" }

func (trapLogging *Syslog_TrapLogging) GetYangName() string { return "trap-logging" }

func (trapLogging *Syslog_TrapLogging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapLogging *Syslog_TrapLogging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapLogging *Syslog_TrapLogging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapLogging *Syslog_TrapLogging) SetParent(parent types.Entity) { trapLogging.parent = parent }

func (trapLogging *Syslog_TrapLogging) GetParent() types.Entity { return trapLogging.parent }

func (trapLogging *Syslog_TrapLogging) GetParentYangName() string { return "syslog" }

// Syslog_BufferedLogging
// Set buffered logging parameters
type Syslog_BufferedLogging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging level for Buffered logging. The type is LoggingLevels. The default
    // value is debug.
    LoggingLevel interface{}

    // Logging buffered size. The type is interface{} with range:
    // 4096..4294967295. The default value is 2097152.
    BufferSize interface{}

    // Set buffered logging discriminators.
    BufferedDiscriminator Syslog_BufferedLogging_BufferedDiscriminator
}

func (bufferedLogging *Syslog_BufferedLogging) GetFilter() yfilter.YFilter { return bufferedLogging.YFilter }

func (bufferedLogging *Syslog_BufferedLogging) SetFilter(yf yfilter.YFilter) { bufferedLogging.YFilter = yf }

func (bufferedLogging *Syslog_BufferedLogging) GetGoName(yname string) string {
    if yname == "logging-level" { return "LoggingLevel" }
    if yname == "buffer-size" { return "BufferSize" }
    if yname == "buffered-discriminator" { return "BufferedDiscriminator" }
    return ""
}

func (bufferedLogging *Syslog_BufferedLogging) GetSegmentPath() string {
    return "buffered-logging"
}

func (bufferedLogging *Syslog_BufferedLogging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "buffered-discriminator" {
        return &bufferedLogging.BufferedDiscriminator
    }
    return nil
}

func (bufferedLogging *Syslog_BufferedLogging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["buffered-discriminator"] = &bufferedLogging.BufferedDiscriminator
    return children
}

func (bufferedLogging *Syslog_BufferedLogging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging-level"] = bufferedLogging.LoggingLevel
    leafs["buffer-size"] = bufferedLogging.BufferSize
    return leafs
}

func (bufferedLogging *Syslog_BufferedLogging) GetBundleName() string { return "cisco_ios_xr" }

func (bufferedLogging *Syslog_BufferedLogging) GetYangName() string { return "buffered-logging" }

func (bufferedLogging *Syslog_BufferedLogging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bufferedLogging *Syslog_BufferedLogging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bufferedLogging *Syslog_BufferedLogging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bufferedLogging *Syslog_BufferedLogging) SetParent(parent types.Entity) { bufferedLogging.parent = parent }

func (bufferedLogging *Syslog_BufferedLogging) GetParent() types.Entity { return bufferedLogging.parent }

func (bufferedLogging *Syslog_BufferedLogging) GetParentYangName() string { return "syslog" }

// Syslog_BufferedLogging_BufferedDiscriminator
// Set buffered logging discriminators
type Syslog_BufferedLogging_BufferedDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set buffered logging match2 discriminator. The type is string.
    Match2 interface{}

    // Set buffered logging no-match1 discriminator. The type is string.
    Nomatch1 interface{}

    // Set buffered logging match1 discriminator. The type is string.
    Match1 interface{}

    // Set buffered logging no-match3 discriminator. The type is string.
    Nomatch3 interface{}

    // Set buffered logging match3 discriminator. The type is string.
    Match3 interface{}

    // Set buffered logging no-match2 discriminator. The type is string.
    Nomatch2 interface{}
}

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetFilter() yfilter.YFilter { return bufferedDiscriminator.YFilter }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) SetFilter(yf yfilter.YFilter) { bufferedDiscriminator.YFilter = yf }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetGoName(yname string) string {
    if yname == "match2" { return "Match2" }
    if yname == "nomatch1" { return "Nomatch1" }
    if yname == "match1" { return "Match1" }
    if yname == "nomatch3" { return "Nomatch3" }
    if yname == "match3" { return "Match3" }
    if yname == "nomatch2" { return "Nomatch2" }
    return ""
}

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetSegmentPath() string {
    return "buffered-discriminator"
}

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["match2"] = bufferedDiscriminator.Match2
    leafs["nomatch1"] = bufferedDiscriminator.Nomatch1
    leafs["match1"] = bufferedDiscriminator.Match1
    leafs["nomatch3"] = bufferedDiscriminator.Nomatch3
    leafs["match3"] = bufferedDiscriminator.Match3
    leafs["nomatch2"] = bufferedDiscriminator.Nomatch2
    return leafs
}

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetYangName() string { return "buffered-discriminator" }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) SetParent(parent types.Entity) { bufferedDiscriminator.parent = parent }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetParent() types.Entity { return bufferedDiscriminator.parent }

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetParentYangName() string { return "buffered-logging" }

// Syslog_HostServer
// Configure logging host
type Syslog_HostServer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs Syslog_HostServer_Vrfs
}

func (hostServer *Syslog_HostServer) GetFilter() yfilter.YFilter { return hostServer.YFilter }

func (hostServer *Syslog_HostServer) SetFilter(yf yfilter.YFilter) { hostServer.YFilter = yf }

func (hostServer *Syslog_HostServer) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (hostServer *Syslog_HostServer) GetSegmentPath() string {
    return "host-server"
}

func (hostServer *Syslog_HostServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &hostServer.Vrfs
    }
    return nil
}

func (hostServer *Syslog_HostServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &hostServer.Vrfs
    return children
}

func (hostServer *Syslog_HostServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hostServer *Syslog_HostServer) GetBundleName() string { return "cisco_ios_xr" }

func (hostServer *Syslog_HostServer) GetYangName() string { return "host-server" }

func (hostServer *Syslog_HostServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostServer *Syslog_HostServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostServer *Syslog_HostServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostServer *Syslog_HostServer) SetParent(parent types.Entity) { hostServer.parent = parent }

func (hostServer *Syslog_HostServer) GetParent() types.Entity { return hostServer.parent }

func (hostServer *Syslog_HostServer) GetParentYangName() string { return "syslog" }

// Syslog_HostServer_Vrfs
// VRF table
type Syslog_HostServer_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of Syslog_HostServer_Vrfs_Vrf.
    Vrf []Syslog_HostServer_Vrfs_Vrf
}

func (vrfs *Syslog_HostServer_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Syslog_HostServer_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Syslog_HostServer_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Syslog_HostServer_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Syslog_HostServer_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_HostServer_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Syslog_HostServer_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Syslog_HostServer_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Syslog_HostServer_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Syslog_HostServer_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Syslog_HostServer_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Syslog_HostServer_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Syslog_HostServer_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Syslog_HostServer_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Syslog_HostServer_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Syslog_HostServer_Vrfs) GetParentYangName() string { return "host-server" }

// Syslog_HostServer_Vrfs_Vrf
// VRF specific data
type Syslog_HostServer_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // List of the IPv6 logging host.
    Ipv6S Syslog_HostServer_Vrfs_Vrf_Ipv6S

    // List of the logging host.
    Hosts Syslog_HostServer_Vrfs_Vrf_Hosts

    // List of the IPv4 logging host.
    Ipv4S Syslog_HostServer_Vrfs_Vrf_Ipv4S
}

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Syslog_HostServer_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv6s" { return "Ipv6S" }
    if yname == "hosts" { return "Hosts" }
    if yname == "ipv4s" { return "Ipv4S" }
    return ""
}

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6s" {
        return &vrf.Ipv6S
    }
    if childYangName == "hosts" {
        return &vrf.Hosts
    }
    if childYangName == "ipv4s" {
        return &vrf.Ipv4S
    }
    return nil
}

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6s"] = &vrf.Ipv6S
    children["hosts"] = &vrf.Hosts
    children["ipv4s"] = &vrf.Ipv4S
    return children
}

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Syslog_HostServer_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Syslog_HostServer_Vrfs_Vrf_Ipv6S
// List of the IPv6 logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address of the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6.
    Ipv6 []Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6
}

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetFilter() yfilter.YFilter { return ipv6S.YFilter }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) SetFilter(yf yfilter.YFilter) { ipv6S.YFilter = yf }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetSegmentPath() string {
    return "ipv6s"
}

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        for _, c := range ipv6S.Ipv6 {
            if ipv6S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6{}
        ipv6S.Ipv6 = append(ipv6S.Ipv6, child)
        return &ipv6S.Ipv6[len(ipv6S.Ipv6)-1]
    }
    return nil
}

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6S.Ipv6 {
        children[ipv6S.Ipv6[i].GetSegmentPath()] = &ipv6S.Ipv6[i]
    }
    return children
}

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetYangName() string { return "ipv6s" }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) SetParent(parent types.Entity) { ipv6S.parent = parent }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetParent() types.Entity { return ipv6S.parent }

func (ipv6S *Syslog_HostServer_Vrfs_Vrf_Ipv6S) GetParentYangName() string { return "vrf" }

// Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6
// IPv6 address of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address of the logging host. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Severity/Port for the logging host.
    Ipv6SeverityPort Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort

    // Set IPv6 logging discriminators.
    Ipv6Discriminator Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator

    // Severity container of the logging host.
    Ipv6SeverityLevels Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels
}

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "ipv6-severity-port" { return "Ipv6SeverityPort" }
    if yname == "ipv6-discriminator" { return "Ipv6Discriminator" }
    if yname == "ipv6-severity-levels" { return "Ipv6SeverityLevels" }
    return ""
}

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetSegmentPath() string {
    return "ipv6" + "[address='" + fmt.Sprintf("%v", ipv6.Address) + "']"
}

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-severity-port" {
        return &ipv6.Ipv6SeverityPort
    }
    if childYangName == "ipv6-discriminator" {
        return &ipv6.Ipv6Discriminator
    }
    if childYangName == "ipv6-severity-levels" {
        return &ipv6.Ipv6SeverityLevels
    }
    return nil
}

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6-severity-port"] = &ipv6.Ipv6SeverityPort
    children["ipv6-discriminator"] = &ipv6.Ipv6Discriminator
    children["ipv6-severity-levels"] = &ipv6.Ipv6SeverityLevels
    return children
}

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv6.Address
    return leafs
}

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6) GetParentYangName() string { return "ipv6s" }

// Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort
// Severity/Port for the logging host
// This type is a presence type.
type Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 6.
    Severity interface{}

    // Port for the logging host. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 514.
    Port interface{}
}

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetFilter() yfilter.YFilter { return ipv6SeverityPort.YFilter }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) SetFilter(yf yfilter.YFilter) { ipv6SeverityPort.YFilter = yf }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    if yname == "port" { return "Port" }
    return ""
}

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetSegmentPath() string {
    return "ipv6-severity-port"
}

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = ipv6SeverityPort.Severity
    leafs["port"] = ipv6SeverityPort.Port
    return leafs
}

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetYangName() string { return "ipv6-severity-port" }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) SetParent(parent types.Entity) { ipv6SeverityPort.parent = parent }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetParent() types.Entity { return ipv6SeverityPort.parent }

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityPort) GetParentYangName() string { return "ipv6" }

// Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator
// Set IPv6 logging discriminators
type Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set IPv6 logging no-match2 discriminator. The type is string.
    Nomatch2 interface{}

    // Set IPv6 logging match3 discriminator. The type is string.
    Match3 interface{}

    // Set IPv6 logging no-match3 discriminator. The type is string.
    Nomatch3 interface{}

    // Set IPv6 logging match1 discriminator. The type is string.
    Match1 interface{}

    // Set IPv6 logging no-match1 discriminator. The type is string.
    Nomatch1 interface{}

    // Set IPv6 logging match2 discriminator. The type is string.
    Match2 interface{}
}

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetFilter() yfilter.YFilter { return ipv6Discriminator.YFilter }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) SetFilter(yf yfilter.YFilter) { ipv6Discriminator.YFilter = yf }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetGoName(yname string) string {
    if yname == "nomatch2" { return "Nomatch2" }
    if yname == "match3" { return "Match3" }
    if yname == "nomatch3" { return "Nomatch3" }
    if yname == "match1" { return "Match1" }
    if yname == "nomatch1" { return "Nomatch1" }
    if yname == "match2" { return "Match2" }
    return ""
}

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetSegmentPath() string {
    return "ipv6-discriminator"
}

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nomatch2"] = ipv6Discriminator.Nomatch2
    leafs["match3"] = ipv6Discriminator.Match3
    leafs["nomatch3"] = ipv6Discriminator.Nomatch3
    leafs["match1"] = ipv6Discriminator.Match1
    leafs["nomatch1"] = ipv6Discriminator.Nomatch1
    leafs["match2"] = ipv6Discriminator.Match2
    return leafs
}

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetYangName() string { return "ipv6-discriminator" }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) SetParent(parent types.Entity) { ipv6Discriminator.parent = parent }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetParent() types.Entity { return ipv6Discriminator.parent }

func (ipv6Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6Discriminator) GetParentYangName() string { return "ipv6" }

// Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels
// Severity container of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel.
    Ipv6SeverityLevel []Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel
}

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetFilter() yfilter.YFilter { return ipv6SeverityLevels.YFilter }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) SetFilter(yf yfilter.YFilter) { ipv6SeverityLevels.YFilter = yf }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetGoName(yname string) string {
    if yname == "ipv6-severity-level" { return "Ipv6SeverityLevel" }
    return ""
}

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetSegmentPath() string {
    return "ipv6-severity-levels"
}

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-severity-level" {
        for _, c := range ipv6SeverityLevels.Ipv6SeverityLevel {
            if ipv6SeverityLevels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel{}
        ipv6SeverityLevels.Ipv6SeverityLevel = append(ipv6SeverityLevels.Ipv6SeverityLevel, child)
        return &ipv6SeverityLevels.Ipv6SeverityLevel[len(ipv6SeverityLevels.Ipv6SeverityLevel)-1]
    }
    return nil
}

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6SeverityLevels.Ipv6SeverityLevel {
        children[ipv6SeverityLevels.Ipv6SeverityLevel[i].GetSegmentPath()] = &ipv6SeverityLevels.Ipv6SeverityLevel[i]
    }
    return children
}

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetYangName() string { return "ipv6-severity-levels" }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) SetParent(parent types.Entity) { ipv6SeverityLevels.parent = parent }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetParent() types.Entity { return ipv6SeverityLevels.parent }

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels) GetParentYangName() string { return "ipv6" }

// Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel
// Severity for the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Severity for the logging host. The type is
    // LogSeverity.
    Severity interface{}
}

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetFilter() yfilter.YFilter { return ipv6SeverityLevel.YFilter }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) SetFilter(yf yfilter.YFilter) { ipv6SeverityLevel.YFilter = yf }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    return ""
}

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetSegmentPath() string {
    return "ipv6-severity-level" + "[severity='" + fmt.Sprintf("%v", ipv6SeverityLevel.Severity) + "']"
}

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = ipv6SeverityLevel.Severity
    return leafs
}

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetYangName() string { return "ipv6-severity-level" }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) SetParent(parent types.Entity) { ipv6SeverityLevel.parent = parent }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetParent() types.Entity { return ipv6SeverityLevel.parent }

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6S_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetParentYangName() string { return "ipv6-severity-levels" }

// Syslog_HostServer_Vrfs_Vrf_Hosts
// List of the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Hosts_Host.
    Host []Syslog_HostServer_Vrfs_Vrf_Hosts_Host
}

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_HostServer_Vrfs_Vrf_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetYangName() string { return "hosts" }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetParentYangName() string { return "vrf" }

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host
// Name of the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the logging host. The type is string.
    HostName interface{}

    // Severity container of the logging host.
    HostNameSeverities Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities

    // Set Hostname logging discriminators.
    HostNameDiscriminator Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator

    // Severity/Port for the logging host.
    HostSeverityPort Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort
}

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetGoName(yname string) string {
    if yname == "host-name" { return "HostName" }
    if yname == "host-name-severities" { return "HostNameSeverities" }
    if yname == "host-name-discriminator" { return "HostNameDiscriminator" }
    if yname == "host-severity-port" { return "HostSeverityPort" }
    return ""
}

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetSegmentPath() string {
    return "host" + "[host-name='" + fmt.Sprintf("%v", host.HostName) + "']"
}

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host-name-severities" {
        return &host.HostNameSeverities
    }
    if childYangName == "host-name-discriminator" {
        return &host.HostNameDiscriminator
    }
    if childYangName == "host-severity-port" {
        return &host.HostSeverityPort
    }
    return nil
}

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["host-name-severities"] = &host.HostNameSeverities
    children["host-name-discriminator"] = &host.HostNameDiscriminator
    children["host-severity-port"] = &host.HostSeverityPort
    return children
}

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name"] = host.HostName
    return leafs
}

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetYangName() string { return "host" }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetParentYangName() string { return "hosts" }

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities
// Severity container of the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity.
    HostNameSeverity []Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity
}

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetFilter() yfilter.YFilter { return hostNameSeverities.YFilter }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) SetFilter(yf yfilter.YFilter) { hostNameSeverities.YFilter = yf }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetGoName(yname string) string {
    if yname == "host-name-severity" { return "HostNameSeverity" }
    return ""
}

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetSegmentPath() string {
    return "host-name-severities"
}

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host-name-severity" {
        for _, c := range hostNameSeverities.HostNameSeverity {
            if hostNameSeverities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity{}
        hostNameSeverities.HostNameSeverity = append(hostNameSeverities.HostNameSeverity, child)
        return &hostNameSeverities.HostNameSeverity[len(hostNameSeverities.HostNameSeverity)-1]
    }
    return nil
}

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hostNameSeverities.HostNameSeverity {
        children[hostNameSeverities.HostNameSeverity[i].GetSegmentPath()] = &hostNameSeverities.HostNameSeverity[i]
    }
    return children
}

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetBundleName() string { return "cisco_ios_xr" }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetYangName() string { return "host-name-severities" }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) SetParent(parent types.Entity) { hostNameSeverities.parent = parent }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetParent() types.Entity { return hostNameSeverities.parent }

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetParentYangName() string { return "host" }

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity
// Severity for the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Severity for the logging host. The type is
    // LogSeverity.
    Severity interface{}
}

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetFilter() yfilter.YFilter { return hostNameSeverity.YFilter }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) SetFilter(yf yfilter.YFilter) { hostNameSeverity.YFilter = yf }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    return ""
}

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetSegmentPath() string {
    return "host-name-severity" + "[severity='" + fmt.Sprintf("%v", hostNameSeverity.Severity) + "']"
}

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = hostNameSeverity.Severity
    return leafs
}

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetBundleName() string { return "cisco_ios_xr" }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetYangName() string { return "host-name-severity" }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) SetParent(parent types.Entity) { hostNameSeverity.parent = parent }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetParent() types.Entity { return hostNameSeverity.parent }

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetParentYangName() string { return "host-name-severities" }

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator
// Set Hostname logging discriminators
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set hostname logging no-match2 discriminator. The type is string.
    Nomatch2 interface{}

    // Set hostname logging match3 discriminator. The type is string.
    Match3 interface{}

    // Set hostname logging no-match3 discriminator. The type is string.
    Nomatch3 interface{}

    // Set hostname logging match1 discriminator. The type is string.
    Match1 interface{}

    // Set hostname logging no-match1 discriminator. The type is string.
    Nomatch1 interface{}

    // Set hostname logging match2 discriminator. The type is string.
    Match2 interface{}
}

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetFilter() yfilter.YFilter { return hostNameDiscriminator.YFilter }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) SetFilter(yf yfilter.YFilter) { hostNameDiscriminator.YFilter = yf }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetGoName(yname string) string {
    if yname == "nomatch2" { return "Nomatch2" }
    if yname == "match3" { return "Match3" }
    if yname == "nomatch3" { return "Nomatch3" }
    if yname == "match1" { return "Match1" }
    if yname == "nomatch1" { return "Nomatch1" }
    if yname == "match2" { return "Match2" }
    return ""
}

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetSegmentPath() string {
    return "host-name-discriminator"
}

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nomatch2"] = hostNameDiscriminator.Nomatch2
    leafs["match3"] = hostNameDiscriminator.Match3
    leafs["nomatch3"] = hostNameDiscriminator.Nomatch3
    leafs["match1"] = hostNameDiscriminator.Match1
    leafs["nomatch1"] = hostNameDiscriminator.Nomatch1
    leafs["match2"] = hostNameDiscriminator.Match2
    return leafs
}

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetYangName() string { return "host-name-discriminator" }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) SetParent(parent types.Entity) { hostNameDiscriminator.parent = parent }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetParent() types.Entity { return hostNameDiscriminator.parent }

func (hostNameDiscriminator *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameDiscriminator) GetParentYangName() string { return "host" }

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort
// Severity/Port for the logging host
// This type is a presence type.
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 6.
    Severity interface{}

    // Port for the logging host. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 514.
    Port interface{}
}

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetFilter() yfilter.YFilter { return hostSeverityPort.YFilter }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) SetFilter(yf yfilter.YFilter) { hostSeverityPort.YFilter = yf }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    if yname == "port" { return "Port" }
    return ""
}

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetSegmentPath() string {
    return "host-severity-port"
}

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = hostSeverityPort.Severity
    leafs["port"] = hostSeverityPort.Port
    return leafs
}

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetBundleName() string { return "cisco_ios_xr" }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetYangName() string { return "host-severity-port" }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) SetParent(parent types.Entity) { hostSeverityPort.parent = parent }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetParent() types.Entity { return hostSeverityPort.parent }

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetParentYangName() string { return "host" }

// Syslog_HostServer_Vrfs_Vrf_Ipv4S
// List of the IPv4 logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address of the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4.
    Ipv4 []Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4
}

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetFilter() yfilter.YFilter { return ipv4S.YFilter }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) SetFilter(yf yfilter.YFilter) { ipv4S.YFilter = yf }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetSegmentPath() string {
    return "ipv4s"
}

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        for _, c := range ipv4S.Ipv4 {
            if ipv4S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4{}
        ipv4S.Ipv4 = append(ipv4S.Ipv4, child)
        return &ipv4S.Ipv4[len(ipv4S.Ipv4)-1]
    }
    return nil
}

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4S.Ipv4 {
        children[ipv4S.Ipv4[i].GetSegmentPath()] = &ipv4S.Ipv4[i]
    }
    return children
}

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetYangName() string { return "ipv4s" }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) SetParent(parent types.Entity) { ipv4S.parent = parent }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetParent() types.Entity { return ipv4S.parent }

func (ipv4S *Syslog_HostServer_Vrfs_Vrf_Ipv4S) GetParentYangName() string { return "vrf" }

// Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4
// IPv4 address of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address of the logging host. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Severity container of the logging host.
    Ipv4SeverityLevels Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels

    // Severity/Port for the logging host.
    Ipv4SeverityPort Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort

    // Set IPv4 logging discriminators.
    Ipv4Discriminator Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator
}

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "ipv4-severity-levels" { return "Ipv4SeverityLevels" }
    if yname == "ipv4-severity-port" { return "Ipv4SeverityPort" }
    if yname == "ipv4-discriminator" { return "Ipv4Discriminator" }
    return ""
}

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetSegmentPath() string {
    return "ipv4" + "[address='" + fmt.Sprintf("%v", ipv4.Address) + "']"
}

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-severity-levels" {
        return &ipv4.Ipv4SeverityLevels
    }
    if childYangName == "ipv4-severity-port" {
        return &ipv4.Ipv4SeverityPort
    }
    if childYangName == "ipv4-discriminator" {
        return &ipv4.Ipv4Discriminator
    }
    return nil
}

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-severity-levels"] = &ipv4.Ipv4SeverityLevels
    children["ipv4-severity-port"] = &ipv4.Ipv4SeverityPort
    children["ipv4-discriminator"] = &ipv4.Ipv4Discriminator
    return children
}

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv4.Address
    return leafs
}

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4) GetParentYangName() string { return "ipv4s" }

// Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels
// Severity container of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel.
    Ipv4SeverityLevel []Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel
}

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetFilter() yfilter.YFilter { return ipv4SeverityLevels.YFilter }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) SetFilter(yf yfilter.YFilter) { ipv4SeverityLevels.YFilter = yf }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetGoName(yname string) string {
    if yname == "ipv4-severity-level" { return "Ipv4SeverityLevel" }
    return ""
}

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetSegmentPath() string {
    return "ipv4-severity-levels"
}

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-severity-level" {
        for _, c := range ipv4SeverityLevels.Ipv4SeverityLevel {
            if ipv4SeverityLevels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel{}
        ipv4SeverityLevels.Ipv4SeverityLevel = append(ipv4SeverityLevels.Ipv4SeverityLevel, child)
        return &ipv4SeverityLevels.Ipv4SeverityLevel[len(ipv4SeverityLevels.Ipv4SeverityLevel)-1]
    }
    return nil
}

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4SeverityLevels.Ipv4SeverityLevel {
        children[ipv4SeverityLevels.Ipv4SeverityLevel[i].GetSegmentPath()] = &ipv4SeverityLevels.Ipv4SeverityLevel[i]
    }
    return children
}

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetYangName() string { return "ipv4-severity-levels" }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) SetParent(parent types.Entity) { ipv4SeverityLevels.parent = parent }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetParent() types.Entity { return ipv4SeverityLevels.parent }

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels) GetParentYangName() string { return "ipv4" }

// Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel
// Severity for the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Severity for the logging host. The type is
    // LogSeverity.
    Severity interface{}
}

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetFilter() yfilter.YFilter { return ipv4SeverityLevel.YFilter }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) SetFilter(yf yfilter.YFilter) { ipv4SeverityLevel.YFilter = yf }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    return ""
}

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetSegmentPath() string {
    return "ipv4-severity-level" + "[severity='" + fmt.Sprintf("%v", ipv4SeverityLevel.Severity) + "']"
}

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = ipv4SeverityLevel.Severity
    return leafs
}

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetYangName() string { return "ipv4-severity-level" }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) SetParent(parent types.Entity) { ipv4SeverityLevel.parent = parent }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetParent() types.Entity { return ipv4SeverityLevel.parent }

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetParentYangName() string { return "ipv4-severity-levels" }

// Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort
// Severity/Port for the logging host
// This type is a presence type.
type Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 6.
    Severity interface{}

    // Port for the logging host. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 514.
    Port interface{}
}

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetFilter() yfilter.YFilter { return ipv4SeverityPort.YFilter }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) SetFilter(yf yfilter.YFilter) { ipv4SeverityPort.YFilter = yf }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    if yname == "port" { return "Port" }
    return ""
}

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetSegmentPath() string {
    return "ipv4-severity-port"
}

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = ipv4SeverityPort.Severity
    leafs["port"] = ipv4SeverityPort.Port
    return leafs
}

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetYangName() string { return "ipv4-severity-port" }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) SetParent(parent types.Entity) { ipv4SeverityPort.parent = parent }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetParent() types.Entity { return ipv4SeverityPort.parent }

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4SeverityPort) GetParentYangName() string { return "ipv4" }

// Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator
// Set IPv4 logging discriminators
type Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set IPv4 logging no-match2 discriminator. The type is string.
    Nomatch2 interface{}

    // Set IPv4 logging match3 discriminator. The type is string.
    Match3 interface{}

    // Set IPv4 logging no-match3 discriminator. The type is string.
    Nomatch3 interface{}

    // Set IPv4 logging match1 discriminator. The type is string.
    Match1 interface{}

    // Set IPv4 logging no-match1 discriminator. The type is string.
    Nomatch1 interface{}

    // Set IPv4 logging match2 discriminator. The type is string.
    Match2 interface{}
}

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetFilter() yfilter.YFilter { return ipv4Discriminator.YFilter }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) SetFilter(yf yfilter.YFilter) { ipv4Discriminator.YFilter = yf }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetGoName(yname string) string {
    if yname == "nomatch2" { return "Nomatch2" }
    if yname == "match3" { return "Match3" }
    if yname == "nomatch3" { return "Nomatch3" }
    if yname == "match1" { return "Match1" }
    if yname == "nomatch1" { return "Nomatch1" }
    if yname == "match2" { return "Match2" }
    return ""
}

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetSegmentPath() string {
    return "ipv4-discriminator"
}

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nomatch2"] = ipv4Discriminator.Nomatch2
    leafs["match3"] = ipv4Discriminator.Match3
    leafs["nomatch3"] = ipv4Discriminator.Nomatch3
    leafs["match1"] = ipv4Discriminator.Match1
    leafs["nomatch1"] = ipv4Discriminator.Nomatch1
    leafs["match2"] = ipv4Discriminator.Match2
    return leafs
}

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetYangName() string { return "ipv4-discriminator" }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) SetParent(parent types.Entity) { ipv4Discriminator.parent = parent }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetParent() types.Entity { return ipv4Discriminator.parent }

func (ipv4Discriminator *Syslog_HostServer_Vrfs_Vrf_Ipv4S_Ipv4_Ipv4Discriminator) GetParentYangName() string { return "ipv4" }

// Syslog_ConsoleLogging
// Set console logging
type Syslog_ConsoleLogging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Console logging level. The type is LoggingLevels. The default value is
    // warning.
    LoggingLevel interface{}

    // Set console logging discriminators.
    ConsoleDiscriminator Syslog_ConsoleLogging_ConsoleDiscriminator
}

func (consoleLogging *Syslog_ConsoleLogging) GetFilter() yfilter.YFilter { return consoleLogging.YFilter }

func (consoleLogging *Syslog_ConsoleLogging) SetFilter(yf yfilter.YFilter) { consoleLogging.YFilter = yf }

func (consoleLogging *Syslog_ConsoleLogging) GetGoName(yname string) string {
    if yname == "logging-level" { return "LoggingLevel" }
    if yname == "console-discriminator" { return "ConsoleDiscriminator" }
    return ""
}

func (consoleLogging *Syslog_ConsoleLogging) GetSegmentPath() string {
    return "console-logging"
}

func (consoleLogging *Syslog_ConsoleLogging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "console-discriminator" {
        return &consoleLogging.ConsoleDiscriminator
    }
    return nil
}

func (consoleLogging *Syslog_ConsoleLogging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["console-discriminator"] = &consoleLogging.ConsoleDiscriminator
    return children
}

func (consoleLogging *Syslog_ConsoleLogging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging-level"] = consoleLogging.LoggingLevel
    return leafs
}

func (consoleLogging *Syslog_ConsoleLogging) GetBundleName() string { return "cisco_ios_xr" }

func (consoleLogging *Syslog_ConsoleLogging) GetYangName() string { return "console-logging" }

func (consoleLogging *Syslog_ConsoleLogging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (consoleLogging *Syslog_ConsoleLogging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (consoleLogging *Syslog_ConsoleLogging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (consoleLogging *Syslog_ConsoleLogging) SetParent(parent types.Entity) { consoleLogging.parent = parent }

func (consoleLogging *Syslog_ConsoleLogging) GetParent() types.Entity { return consoleLogging.parent }

func (consoleLogging *Syslog_ConsoleLogging) GetParentYangName() string { return "syslog" }

// Syslog_ConsoleLogging_ConsoleDiscriminator
// Set console logging discriminators
type Syslog_ConsoleLogging_ConsoleDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set console logging match2 discriminator. The type is string.
    Match2 interface{}

    // Set console logging no-match1 discriminator. The type is string.
    Nomatch1 interface{}

    // Set console logging match1 discriminator. The type is string.
    Match1 interface{}

    // Set console logging no-match3 discriminator. The type is string.
    Nomatch3 interface{}

    // Set console logging match3 discriminator. The type is string.
    Match3 interface{}

    // Set console logging no-match2 discriminator. The type is string.
    Nomatch2 interface{}
}

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetFilter() yfilter.YFilter { return consoleDiscriminator.YFilter }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) SetFilter(yf yfilter.YFilter) { consoleDiscriminator.YFilter = yf }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetGoName(yname string) string {
    if yname == "match2" { return "Match2" }
    if yname == "nomatch1" { return "Nomatch1" }
    if yname == "match1" { return "Match1" }
    if yname == "nomatch3" { return "Nomatch3" }
    if yname == "match3" { return "Match3" }
    if yname == "nomatch2" { return "Nomatch2" }
    return ""
}

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetSegmentPath() string {
    return "console-discriminator"
}

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["match2"] = consoleDiscriminator.Match2
    leafs["nomatch1"] = consoleDiscriminator.Nomatch1
    leafs["match1"] = consoleDiscriminator.Match1
    leafs["nomatch3"] = consoleDiscriminator.Nomatch3
    leafs["match3"] = consoleDiscriminator.Match3
    leafs["nomatch2"] = consoleDiscriminator.Nomatch2
    return leafs
}

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetYangName() string { return "console-discriminator" }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) SetParent(parent types.Entity) { consoleDiscriminator.parent = parent }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetParent() types.Entity { return consoleDiscriminator.parent }

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetParentYangName() string { return "console-logging" }

// Syslog_Files
// Configure logging file destination
type Syslog_Files struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify File Name. The type is slice of Syslog_Files_File.
    File []Syslog_Files_File
}

func (files *Syslog_Files) GetFilter() yfilter.YFilter { return files.YFilter }

func (files *Syslog_Files) SetFilter(yf yfilter.YFilter) { files.YFilter = yf }

func (files *Syslog_Files) GetGoName(yname string) string {
    if yname == "file" { return "File" }
    return ""
}

func (files *Syslog_Files) GetSegmentPath() string {
    return "files"
}

func (files *Syslog_Files) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "file" {
        for _, c := range files.File {
            if files.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Files_File{}
        files.File = append(files.File, child)
        return &files.File[len(files.File)-1]
    }
    return nil
}

func (files *Syslog_Files) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range files.File {
        children[files.File[i].GetSegmentPath()] = &files.File[i]
    }
    return children
}

func (files *Syslog_Files) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (files *Syslog_Files) GetBundleName() string { return "cisco_ios_xr" }

func (files *Syslog_Files) GetYangName() string { return "files" }

func (files *Syslog_Files) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (files *Syslog_Files) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (files *Syslog_Files) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (files *Syslog_Files) SetParent(parent types.Entity) { files.parent = parent }

func (files *Syslog_Files) GetParent() types.Entity { return files.parent }

func (files *Syslog_Files) GetParentYangName() string { return "syslog" }

// Syslog_Files_File
// Specify File Name
type Syslog_Files_File struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the file. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    FileName interface{}

    // Specifications of the logging file destination.
    FileSpecification Syslog_Files_File_FileSpecification

    // Attributes of the logging file destination.
    FileLogAttributes Syslog_Files_File_FileLogAttributes

    // Set File logging discriminators.
    FileLogDiscriminator Syslog_Files_File_FileLogDiscriminator
}

func (file *Syslog_Files_File) GetFilter() yfilter.YFilter { return file.YFilter }

func (file *Syslog_Files_File) SetFilter(yf yfilter.YFilter) { file.YFilter = yf }

func (file *Syslog_Files_File) GetGoName(yname string) string {
    if yname == "file-name" { return "FileName" }
    if yname == "file-specification" { return "FileSpecification" }
    if yname == "file-log-attributes" { return "FileLogAttributes" }
    if yname == "file-log-discriminator" { return "FileLogDiscriminator" }
    return ""
}

func (file *Syslog_Files_File) GetSegmentPath() string {
    return "file" + "[file-name='" + fmt.Sprintf("%v", file.FileName) + "']"
}

func (file *Syslog_Files_File) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "file-specification" {
        return &file.FileSpecification
    }
    if childYangName == "file-log-attributes" {
        return &file.FileLogAttributes
    }
    if childYangName == "file-log-discriminator" {
        return &file.FileLogDiscriminator
    }
    return nil
}

func (file *Syslog_Files_File) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["file-specification"] = &file.FileSpecification
    children["file-log-attributes"] = &file.FileLogAttributes
    children["file-log-discriminator"] = &file.FileLogDiscriminator
    return children
}

func (file *Syslog_Files_File) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["file-name"] = file.FileName
    return leafs
}

func (file *Syslog_Files_File) GetBundleName() string { return "cisco_ios_xr" }

func (file *Syslog_Files_File) GetYangName() string { return "file" }

func (file *Syslog_Files_File) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (file *Syslog_Files_File) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (file *Syslog_Files_File) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (file *Syslog_Files_File) SetParent(parent types.Entity) { file.parent = parent }

func (file *Syslog_Files_File) GetParent() types.Entity { return file.parent }

func (file *Syslog_Files_File) GetParentYangName() string { return "files" }

// Syslog_Files_File_FileSpecification
// Specifications of the logging file destination
type Syslog_Files_File_FileSpecification struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // File path. The type is string.
    Path interface{}

    // Maximum file size (in KB). The type is interface{} with range:
    // -2147483648..2147483647. The default value is 1024.
    MaxFileSize interface{}

    // Severity of messages. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 6.
    Severity interface{}
}

func (fileSpecification *Syslog_Files_File_FileSpecification) GetFilter() yfilter.YFilter { return fileSpecification.YFilter }

func (fileSpecification *Syslog_Files_File_FileSpecification) SetFilter(yf yfilter.YFilter) { fileSpecification.YFilter = yf }

func (fileSpecification *Syslog_Files_File_FileSpecification) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "max-file-size" { return "MaxFileSize" }
    if yname == "severity" { return "Severity" }
    return ""
}

func (fileSpecification *Syslog_Files_File_FileSpecification) GetSegmentPath() string {
    return "file-specification"
}

func (fileSpecification *Syslog_Files_File_FileSpecification) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fileSpecification *Syslog_Files_File_FileSpecification) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fileSpecification *Syslog_Files_File_FileSpecification) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = fileSpecification.Path
    leafs["max-file-size"] = fileSpecification.MaxFileSize
    leafs["severity"] = fileSpecification.Severity
    return leafs
}

func (fileSpecification *Syslog_Files_File_FileSpecification) GetBundleName() string { return "cisco_ios_xr" }

func (fileSpecification *Syslog_Files_File_FileSpecification) GetYangName() string { return "file-specification" }

func (fileSpecification *Syslog_Files_File_FileSpecification) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fileSpecification *Syslog_Files_File_FileSpecification) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fileSpecification *Syslog_Files_File_FileSpecification) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fileSpecification *Syslog_Files_File_FileSpecification) SetParent(parent types.Entity) { fileSpecification.parent = parent }

func (fileSpecification *Syslog_Files_File_FileSpecification) GetParent() types.Entity { return fileSpecification.parent }

func (fileSpecification *Syslog_Files_File_FileSpecification) GetParentYangName() string { return "file" }

// Syslog_Files_File_FileLogAttributes
// Attributes of the logging file destination
type Syslog_Files_File_FileLogAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum file size (in KB). The type is interface{} with range:
    // -2147483648..2147483647. The default value is 1024.
    MaxFileSize interface{}

    // Severity of messages. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 6.
    Severity interface{}
}

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetFilter() yfilter.YFilter { return fileLogAttributes.YFilter }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) SetFilter(yf yfilter.YFilter) { fileLogAttributes.YFilter = yf }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetGoName(yname string) string {
    if yname == "max-file-size" { return "MaxFileSize" }
    if yname == "severity" { return "Severity" }
    return ""
}

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetSegmentPath() string {
    return "file-log-attributes"
}

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-file-size"] = fileLogAttributes.MaxFileSize
    leafs["severity"] = fileLogAttributes.Severity
    return leafs
}

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetYangName() string { return "file-log-attributes" }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) SetParent(parent types.Entity) { fileLogAttributes.parent = parent }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetParent() types.Entity { return fileLogAttributes.parent }

func (fileLogAttributes *Syslog_Files_File_FileLogAttributes) GetParentYangName() string { return "file" }

// Syslog_Files_File_FileLogDiscriminator
// Set File logging discriminators
type Syslog_Files_File_FileLogDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set file logging no match discriminator 2. The type is string.
    Nomatch2 interface{}

    // Set file logging match discriminator 3. The type is string.
    Match3 interface{}

    // Set file logging no match discriminator 3. The type is string.
    Nomatch3 interface{}

    // Set file logging match discriminator 1. The type is string.
    Match1 interface{}

    // Set file logging no match discriminator 1. The type is string.
    Nomatch1 interface{}

    // Set file logging match discriminator 2. The type is string.
    Match2 interface{}
}

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetFilter() yfilter.YFilter { return fileLogDiscriminator.YFilter }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) SetFilter(yf yfilter.YFilter) { fileLogDiscriminator.YFilter = yf }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetGoName(yname string) string {
    if yname == "nomatch2" { return "Nomatch2" }
    if yname == "match3" { return "Match3" }
    if yname == "nomatch3" { return "Nomatch3" }
    if yname == "match1" { return "Match1" }
    if yname == "nomatch1" { return "Nomatch1" }
    if yname == "match2" { return "Match2" }
    return ""
}

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetSegmentPath() string {
    return "file-log-discriminator"
}

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nomatch2"] = fileLogDiscriminator.Nomatch2
    leafs["match3"] = fileLogDiscriminator.Match3
    leafs["nomatch3"] = fileLogDiscriminator.Nomatch3
    leafs["match1"] = fileLogDiscriminator.Match1
    leafs["nomatch1"] = fileLogDiscriminator.Nomatch1
    leafs["match2"] = fileLogDiscriminator.Match2
    return leafs
}

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetYangName() string { return "file-log-discriminator" }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) SetParent(parent types.Entity) { fileLogDiscriminator.parent = parent }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetParent() types.Entity { return fileLogDiscriminator.parent }

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetParentYangName() string { return "file" }

// Syslog_Ipv4
// Syslog TOS bit for outgoing messages
type Syslog_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DSCP value.
    Dscp Syslog_Ipv4_Dscp

    // Type of service.
    Tos Syslog_Ipv4_Tos

    // Precedence value.
    Precedence Syslog_Ipv4_Precedence
}

func (ipv4 *Syslog_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Syslog_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Syslog_Ipv4) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    if yname == "tos" { return "Tos" }
    if yname == "precedence" { return "Precedence" }
    return ""
}

func (ipv4 *Syslog_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Syslog_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dscp" {
        return &ipv4.Dscp
    }
    if childYangName == "tos" {
        return &ipv4.Tos
    }
    if childYangName == "precedence" {
        return &ipv4.Precedence
    }
    return nil
}

func (ipv4 *Syslog_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dscp"] = &ipv4.Dscp
    children["tos"] = &ipv4.Tos
    children["precedence"] = &ipv4.Precedence
    return children
}

func (ipv4 *Syslog_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Syslog_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Syslog_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Syslog_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Syslog_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Syslog_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Syslog_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Syslog_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Syslog_Ipv4) GetParentYangName() string { return "syslog" }

// Syslog_Ipv4_Dscp
// DSCP value
// This type is a presence type.
type Syslog_Ipv4_Dscp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging TOS type DSCP. The type is LoggingDscp. This attribute is
    // mandatory.
    Type interface{}

    // Unused. The type is one of the following types: enumeration
    // LoggingPrecedenceValue, or int with range: 0..7.
    Unused interface{}

    // Logging DSCP value. The type is one of the following types: enumeration
    // LoggingDscpValue This attribute is mandatory., or int with range: 0..63
    // This attribute is mandatory..
    Value interface{}
}

func (dscp *Syslog_Ipv4_Dscp) GetFilter() yfilter.YFilter { return dscp.YFilter }

func (dscp *Syslog_Ipv4_Dscp) SetFilter(yf yfilter.YFilter) { dscp.YFilter = yf }

func (dscp *Syslog_Ipv4_Dscp) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "unused" { return "Unused" }
    if yname == "value" { return "Value" }
    return ""
}

func (dscp *Syslog_Ipv4_Dscp) GetSegmentPath() string {
    return "dscp"
}

func (dscp *Syslog_Ipv4_Dscp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscp *Syslog_Ipv4_Dscp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscp *Syslog_Ipv4_Dscp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = dscp.Type
    leafs["unused"] = dscp.Unused
    leafs["value"] = dscp.Value
    return leafs
}

func (dscp *Syslog_Ipv4_Dscp) GetBundleName() string { return "cisco_ios_xr" }

func (dscp *Syslog_Ipv4_Dscp) GetYangName() string { return "dscp" }

func (dscp *Syslog_Ipv4_Dscp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dscp *Syslog_Ipv4_Dscp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dscp *Syslog_Ipv4_Dscp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dscp *Syslog_Ipv4_Dscp) SetParent(parent types.Entity) { dscp.parent = parent }

func (dscp *Syslog_Ipv4_Dscp) GetParent() types.Entity { return dscp.parent }

func (dscp *Syslog_Ipv4_Dscp) GetParentYangName() string { return "ipv4" }

// Syslog_Ipv4_Tos
// Type of service
type Syslog_Ipv4_Tos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging TOS type DSCP or precedence. The type is LoggingTos.
    Type interface{}

    // Logging precedence value. The type is one of the following types:
    // enumeration LoggingPrecedenceValue, or int with range: 0..7.
    Precedence interface{}

    // Logging DSCP value. The type is one of the following types: enumeration
    // LoggingDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (tos *Syslog_Ipv4_Tos) GetFilter() yfilter.YFilter { return tos.YFilter }

func (tos *Syslog_Ipv4_Tos) SetFilter(yf yfilter.YFilter) { tos.YFilter = yf }

func (tos *Syslog_Ipv4_Tos) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "precedence" { return "Precedence" }
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (tos *Syslog_Ipv4_Tos) GetSegmentPath() string {
    return "tos"
}

func (tos *Syslog_Ipv4_Tos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tos *Syslog_Ipv4_Tos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tos *Syslog_Ipv4_Tos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = tos.Type
    leafs["precedence"] = tos.Precedence
    leafs["dscp"] = tos.Dscp
    return leafs
}

func (tos *Syslog_Ipv4_Tos) GetBundleName() string { return "cisco_ios_xr" }

func (tos *Syslog_Ipv4_Tos) GetYangName() string { return "tos" }

func (tos *Syslog_Ipv4_Tos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tos *Syslog_Ipv4_Tos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tos *Syslog_Ipv4_Tos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tos *Syslog_Ipv4_Tos) SetParent(parent types.Entity) { tos.parent = parent }

func (tos *Syslog_Ipv4_Tos) GetParent() types.Entity { return tos.parent }

func (tos *Syslog_Ipv4_Tos) GetParentYangName() string { return "ipv4" }

// Syslog_Ipv4_Precedence
// Precedence value
// This type is a presence type.
type Syslog_Ipv4_Precedence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging TOS type precedence. The type is LoggingPrecedence. This attribute
    // is mandatory.
    Type interface{}

    // Logging precedence value. The type is one of the following types:
    // enumeration LoggingPrecedenceValue This attribute is mandatory., or int
    // with range: 0..7 This attribute is mandatory..
    Value interface{}

    // Unused. The type is one of the following types: enumeration
    // LoggingDscpValue, or int with range: 0..63.
    Unused interface{}
}

func (precedence *Syslog_Ipv4_Precedence) GetFilter() yfilter.YFilter { return precedence.YFilter }

func (precedence *Syslog_Ipv4_Precedence) SetFilter(yf yfilter.YFilter) { precedence.YFilter = yf }

func (precedence *Syslog_Ipv4_Precedence) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "value" { return "Value" }
    if yname == "unused" { return "Unused" }
    return ""
}

func (precedence *Syslog_Ipv4_Precedence) GetSegmentPath() string {
    return "precedence"
}

func (precedence *Syslog_Ipv4_Precedence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (precedence *Syslog_Ipv4_Precedence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (precedence *Syslog_Ipv4_Precedence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = precedence.Type
    leafs["value"] = precedence.Value
    leafs["unused"] = precedence.Unused
    return leafs
}

func (precedence *Syslog_Ipv4_Precedence) GetBundleName() string { return "cisco_ios_xr" }

func (precedence *Syslog_Ipv4_Precedence) GetYangName() string { return "precedence" }

func (precedence *Syslog_Ipv4_Precedence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (precedence *Syslog_Ipv4_Precedence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (precedence *Syslog_Ipv4_Precedence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (precedence *Syslog_Ipv4_Precedence) SetParent(parent types.Entity) { precedence.parent = parent }

func (precedence *Syslog_Ipv4_Precedence) GetParent() types.Entity { return precedence.parent }

func (precedence *Syslog_Ipv4_Precedence) GetParentYangName() string { return "ipv4" }

// Syslog_Archive
// Archive attributes configuration
type Syslog_Archive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The total size of the archive. The type is interface{} with range: 1..2047.
    Size interface{}

    // The maximum file size for a single log file. The type is interface{} with
    // range: 1..2047.
    FileSize interface{}

    // '/disk0:' or '/disk1:' or '/harddisk:'. The type is string.
    Device interface{}

    // The size threshold at which a syslog is generated. The type is interface{}
    // with range: 1..99. Units are percentage.
    Threshold interface{}

    // The collection interval for logs. The type is LogCollectFrequency.
    Frequency interface{}

    // The minimum severity of log messages to archive. The type is
    // LogMessageSeverity.
    Severity interface{}

    // The maximum number of weeks of log to maintain. The type is interface{}
    // with range: 1..256.
    Length interface{}
}

func (archive *Syslog_Archive) GetFilter() yfilter.YFilter { return archive.YFilter }

func (archive *Syslog_Archive) SetFilter(yf yfilter.YFilter) { archive.YFilter = yf }

func (archive *Syslog_Archive) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "file-size" { return "FileSize" }
    if yname == "device" { return "Device" }
    if yname == "threshold" { return "Threshold" }
    if yname == "frequency" { return "Frequency" }
    if yname == "severity" { return "Severity" }
    if yname == "length" { return "Length" }
    return ""
}

func (archive *Syslog_Archive) GetSegmentPath() string {
    return "archive"
}

func (archive *Syslog_Archive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (archive *Syslog_Archive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (archive *Syslog_Archive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = archive.Size
    leafs["file-size"] = archive.FileSize
    leafs["device"] = archive.Device
    leafs["threshold"] = archive.Threshold
    leafs["frequency"] = archive.Frequency
    leafs["severity"] = archive.Severity
    leafs["length"] = archive.Length
    return leafs
}

func (archive *Syslog_Archive) GetBundleName() string { return "cisco_ios_xr" }

func (archive *Syslog_Archive) GetYangName() string { return "archive" }

func (archive *Syslog_Archive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (archive *Syslog_Archive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (archive *Syslog_Archive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (archive *Syslog_Archive) SetParent(parent types.Entity) { archive.parent = parent }

func (archive *Syslog_Archive) GetParent() types.Entity { return archive.parent }

func (archive *Syslog_Archive) GetParentYangName() string { return "syslog" }

// Syslog_Ipv6
// Syslog traffic class bit for outgoing messages
type Syslog_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DSCP value.
    Dscp Syslog_Ipv6_Dscp

    // Type of traffic class.
    TrafficClass Syslog_Ipv6_TrafficClass

    // Precedence value.
    Precedence Syslog_Ipv6_Precedence
}

func (ipv6 *Syslog_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Syslog_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Syslog_Ipv6) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "precedence" { return "Precedence" }
    return ""
}

func (ipv6 *Syslog_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Syslog_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dscp" {
        return &ipv6.Dscp
    }
    if childYangName == "traffic-class" {
        return &ipv6.TrafficClass
    }
    if childYangName == "precedence" {
        return &ipv6.Precedence
    }
    return nil
}

func (ipv6 *Syslog_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dscp"] = &ipv6.Dscp
    children["traffic-class"] = &ipv6.TrafficClass
    children["precedence"] = &ipv6.Precedence
    return children
}

func (ipv6 *Syslog_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Syslog_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Syslog_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Syslog_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Syslog_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Syslog_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Syslog_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Syslog_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Syslog_Ipv6) GetParentYangName() string { return "syslog" }

// Syslog_Ipv6_Dscp
// DSCP value
// This type is a presence type.
type Syslog_Ipv6_Dscp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging TOS type DSCP. The type is LoggingDscp. This attribute is
    // mandatory.
    Type interface{}

    // Unused. The type is one of the following types: enumeration
    // LoggingPrecedenceValue, or int with range: 0..7.
    Unused interface{}

    // Logging DSCP value. The type is one of the following types: enumeration
    // LoggingDscpValue This attribute is mandatory., or int with range: 0..63
    // This attribute is mandatory..
    Value interface{}
}

func (dscp *Syslog_Ipv6_Dscp) GetFilter() yfilter.YFilter { return dscp.YFilter }

func (dscp *Syslog_Ipv6_Dscp) SetFilter(yf yfilter.YFilter) { dscp.YFilter = yf }

func (dscp *Syslog_Ipv6_Dscp) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "unused" { return "Unused" }
    if yname == "value" { return "Value" }
    return ""
}

func (dscp *Syslog_Ipv6_Dscp) GetSegmentPath() string {
    return "dscp"
}

func (dscp *Syslog_Ipv6_Dscp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscp *Syslog_Ipv6_Dscp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscp *Syslog_Ipv6_Dscp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = dscp.Type
    leafs["unused"] = dscp.Unused
    leafs["value"] = dscp.Value
    return leafs
}

func (dscp *Syslog_Ipv6_Dscp) GetBundleName() string { return "cisco_ios_xr" }

func (dscp *Syslog_Ipv6_Dscp) GetYangName() string { return "dscp" }

func (dscp *Syslog_Ipv6_Dscp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dscp *Syslog_Ipv6_Dscp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dscp *Syslog_Ipv6_Dscp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dscp *Syslog_Ipv6_Dscp) SetParent(parent types.Entity) { dscp.parent = parent }

func (dscp *Syslog_Ipv6_Dscp) GetParent() types.Entity { return dscp.parent }

func (dscp *Syslog_Ipv6_Dscp) GetParentYangName() string { return "ipv6" }

// Syslog_Ipv6_TrafficClass
// Type of traffic class
type Syslog_Ipv6_TrafficClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging TOS type DSCP or precedence. The type is LoggingTos.
    Type interface{}

    // Logging precedence value. The type is one of the following types:
    // enumeration LoggingPrecedenceValue, or int with range: 0..7.
    Precedence interface{}

    // Logging DSCP value. The type is one of the following types: enumeration
    // LoggingDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (trafficClass *Syslog_Ipv6_TrafficClass) GetFilter() yfilter.YFilter { return trafficClass.YFilter }

func (trafficClass *Syslog_Ipv6_TrafficClass) SetFilter(yf yfilter.YFilter) { trafficClass.YFilter = yf }

func (trafficClass *Syslog_Ipv6_TrafficClass) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "precedence" { return "Precedence" }
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (trafficClass *Syslog_Ipv6_TrafficClass) GetSegmentPath() string {
    return "traffic-class"
}

func (trafficClass *Syslog_Ipv6_TrafficClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trafficClass *Syslog_Ipv6_TrafficClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trafficClass *Syslog_Ipv6_TrafficClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = trafficClass.Type
    leafs["precedence"] = trafficClass.Precedence
    leafs["dscp"] = trafficClass.Dscp
    return leafs
}

func (trafficClass *Syslog_Ipv6_TrafficClass) GetBundleName() string { return "cisco_ios_xr" }

func (trafficClass *Syslog_Ipv6_TrafficClass) GetYangName() string { return "traffic-class" }

func (trafficClass *Syslog_Ipv6_TrafficClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficClass *Syslog_Ipv6_TrafficClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficClass *Syslog_Ipv6_TrafficClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficClass *Syslog_Ipv6_TrafficClass) SetParent(parent types.Entity) { trafficClass.parent = parent }

func (trafficClass *Syslog_Ipv6_TrafficClass) GetParent() types.Entity { return trafficClass.parent }

func (trafficClass *Syslog_Ipv6_TrafficClass) GetParentYangName() string { return "ipv6" }

// Syslog_Ipv6_Precedence
// Precedence value
// This type is a presence type.
type Syslog_Ipv6_Precedence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging TOS type precedence. The type is LoggingPrecedence. This attribute
    // is mandatory.
    Type interface{}

    // Logging precedence value. The type is one of the following types:
    // enumeration LoggingPrecedenceValue This attribute is mandatory., or int
    // with range: 0..7 This attribute is mandatory..
    Value interface{}

    // Unused. The type is one of the following types: enumeration
    // LoggingDscpValue, or int with range: 0..63.
    Unused interface{}
}

func (precedence *Syslog_Ipv6_Precedence) GetFilter() yfilter.YFilter { return precedence.YFilter }

func (precedence *Syslog_Ipv6_Precedence) SetFilter(yf yfilter.YFilter) { precedence.YFilter = yf }

func (precedence *Syslog_Ipv6_Precedence) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "value" { return "Value" }
    if yname == "unused" { return "Unused" }
    return ""
}

func (precedence *Syslog_Ipv6_Precedence) GetSegmentPath() string {
    return "precedence"
}

func (precedence *Syslog_Ipv6_Precedence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (precedence *Syslog_Ipv6_Precedence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (precedence *Syslog_Ipv6_Precedence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = precedence.Type
    leafs["value"] = precedence.Value
    leafs["unused"] = precedence.Unused
    return leafs
}

func (precedence *Syslog_Ipv6_Precedence) GetBundleName() string { return "cisco_ios_xr" }

func (precedence *Syslog_Ipv6_Precedence) GetYangName() string { return "precedence" }

func (precedence *Syslog_Ipv6_Precedence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (precedence *Syslog_Ipv6_Precedence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (precedence *Syslog_Ipv6_Precedence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (precedence *Syslog_Ipv6_Precedence) SetParent(parent types.Entity) { precedence.parent = parent }

func (precedence *Syslog_Ipv6_Precedence) GetParent() types.Entity { return precedence.parent }

func (precedence *Syslog_Ipv6_Precedence) GetParentYangName() string { return "ipv6" }

// Syslog_SourceInterfaceTable
// Configure source interface
type Syslog_SourceInterfaceTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify interface for source address in logging transactions.
    SourceInterfaceValues Syslog_SourceInterfaceTable_SourceInterfaceValues
}

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetFilter() yfilter.YFilter { return sourceInterfaceTable.YFilter }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) SetFilter(yf yfilter.YFilter) { sourceInterfaceTable.YFilter = yf }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetGoName(yname string) string {
    if yname == "source-interface-values" { return "SourceInterfaceValues" }
    return ""
}

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetSegmentPath() string {
    return "source-interface-table"
}

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-interface-values" {
        return &sourceInterfaceTable.SourceInterfaceValues
    }
    return nil
}

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-interface-values"] = &sourceInterfaceTable.SourceInterfaceValues
    return children
}

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetBundleName() string { return "cisco_ios_xr" }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetYangName() string { return "source-interface-table" }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) SetParent(parent types.Entity) { sourceInterfaceTable.parent = parent }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetParent() types.Entity { return sourceInterfaceTable.parent }

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetParentYangName() string { return "syslog" }

// Syslog_SourceInterfaceTable_SourceInterfaceValues
// Specify interface for source address in logging
// transactions
type Syslog_SourceInterfaceTable_SourceInterfaceValues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source interface. The type is slice of
    // Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue.
    SourceInterfaceValue []Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue
}

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetFilter() yfilter.YFilter { return sourceInterfaceValues.YFilter }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) SetFilter(yf yfilter.YFilter) { sourceInterfaceValues.YFilter = yf }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetGoName(yname string) string {
    if yname == "source-interface-value" { return "SourceInterfaceValue" }
    return ""
}

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetSegmentPath() string {
    return "source-interface-values"
}

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-interface-value" {
        for _, c := range sourceInterfaceValues.SourceInterfaceValue {
            if sourceInterfaceValues.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue{}
        sourceInterfaceValues.SourceInterfaceValue = append(sourceInterfaceValues.SourceInterfaceValue, child)
        return &sourceInterfaceValues.SourceInterfaceValue[len(sourceInterfaceValues.SourceInterfaceValue)-1]
    }
    return nil
}

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sourceInterfaceValues.SourceInterfaceValue {
        children[sourceInterfaceValues.SourceInterfaceValue[i].GetSegmentPath()] = &sourceInterfaceValues.SourceInterfaceValue[i]
    }
    return children
}

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetBundleName() string { return "cisco_ios_xr" }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetYangName() string { return "source-interface-values" }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) SetParent(parent types.Entity) { sourceInterfaceValues.parent = parent }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetParent() types.Entity { return sourceInterfaceValues.parent }

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetParentYangName() string { return "source-interface-table" }

// Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue
// Source interface
type Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Which Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SrcInterfaceNameValue interface{}

    // Configure source interface VRF.
    SourceInterfaceVrfs Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs
}

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetFilter() yfilter.YFilter { return sourceInterfaceValue.YFilter }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) SetFilter(yf yfilter.YFilter) { sourceInterfaceValue.YFilter = yf }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetGoName(yname string) string {
    if yname == "src-interface-name-value" { return "SrcInterfaceNameValue" }
    if yname == "source-interface-vrfs" { return "SourceInterfaceVrfs" }
    return ""
}

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetSegmentPath() string {
    return "source-interface-value" + "[src-interface-name-value='" + fmt.Sprintf("%v", sourceInterfaceValue.SrcInterfaceNameValue) + "']"
}

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-interface-vrfs" {
        return &sourceInterfaceValue.SourceInterfaceVrfs
    }
    return nil
}

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-interface-vrfs"] = &sourceInterfaceValue.SourceInterfaceVrfs
    return children
}

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["src-interface-name-value"] = sourceInterfaceValue.SrcInterfaceNameValue
    return leafs
}

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetBundleName() string { return "cisco_ios_xr" }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetYangName() string { return "source-interface-value" }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) SetParent(parent types.Entity) { sourceInterfaceValue.parent = parent }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetParent() types.Entity { return sourceInterfaceValue.parent }

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetParentYangName() string { return "source-interface-values" }

// Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs
// Configure source interface VRF
type Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify VRF for source interface. The type is slice of
    // Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf.
    SourceInterfaceVrf []Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf
}

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetFilter() yfilter.YFilter { return sourceInterfaceVrfs.YFilter }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) SetFilter(yf yfilter.YFilter) { sourceInterfaceVrfs.YFilter = yf }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetGoName(yname string) string {
    if yname == "source-interface-vrf" { return "SourceInterfaceVrf" }
    return ""
}

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetSegmentPath() string {
    return "source-interface-vrfs"
}

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-interface-vrf" {
        for _, c := range sourceInterfaceVrfs.SourceInterfaceVrf {
            if sourceInterfaceVrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf{}
        sourceInterfaceVrfs.SourceInterfaceVrf = append(sourceInterfaceVrfs.SourceInterfaceVrf, child)
        return &sourceInterfaceVrfs.SourceInterfaceVrf[len(sourceInterfaceVrfs.SourceInterfaceVrf)-1]
    }
    return nil
}

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sourceInterfaceVrfs.SourceInterfaceVrf {
        children[sourceInterfaceVrfs.SourceInterfaceVrf[i].GetSegmentPath()] = &sourceInterfaceVrfs.SourceInterfaceVrf[i]
    }
    return children
}

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetBundleName() string { return "cisco_ios_xr" }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetYangName() string { return "source-interface-vrfs" }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) SetParent(parent types.Entity) { sourceInterfaceVrfs.parent = parent }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetParent() types.Entity { return sourceInterfaceVrfs.parent }

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetParentYangName() string { return "source-interface-value" }

// Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf
// Specify VRF for source interface
type Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}
}

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetFilter() yfilter.YFilter { return sourceInterfaceVrf.YFilter }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) SetFilter(yf yfilter.YFilter) { sourceInterfaceVrf.YFilter = yf }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetSegmentPath() string {
    return "source-interface-vrf" + "[vrf-name='" + fmt.Sprintf("%v", sourceInterfaceVrf.VrfName) + "']"
}

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = sourceInterfaceVrf.VrfName
    return leafs
}

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetBundleName() string { return "cisco_ios_xr" }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetYangName() string { return "source-interface-vrf" }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) SetParent(parent types.Entity) { sourceInterfaceVrf.parent = parent }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetParent() types.Entity { return sourceInterfaceVrf.parent }

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetParentYangName() string { return "source-interface-vrfs" }

// Syslog_AlarmLogger
// Alarm Logger Properties
type Syslog_AlarmLogger struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Suppress events from a card/VM till its configuration is complete. The type
    // is interface{}.
    PreConfigSuppression interface{}

    // Log all events with equal or higher (lower level) severity than this. The
    // type is AlarmLoggerSeverityLevel.
    SeverityLevel interface{}

    // Timeout (in minutes) for pre-config events suppression (default 15). The
    // type is interface{} with range: 1..60. Units are minute. The default value
    // is 15.
    PreConfigSuppressionTimeout interface{}

    // Set size of the local event buffer. The type is interface{} with range:
    // 1024..1024000.
    BufferSize interface{}

    // Enable alarm source location in message text. The type is interface{}.
    SourceLocation interface{}

    // Configure threshold (%) for capacity alarm. The type is interface{} with
    // range: 10..100. The default value is 90.
    Threshold interface{}

    // List of filter strings.
    AlarmFilterStrings Syslog_AlarmLogger_AlarmFilterStrings
}

func (alarmLogger *Syslog_AlarmLogger) GetFilter() yfilter.YFilter { return alarmLogger.YFilter }

func (alarmLogger *Syslog_AlarmLogger) SetFilter(yf yfilter.YFilter) { alarmLogger.YFilter = yf }

func (alarmLogger *Syslog_AlarmLogger) GetGoName(yname string) string {
    if yname == "pre-config-suppression" { return "PreConfigSuppression" }
    if yname == "severity-level" { return "SeverityLevel" }
    if yname == "pre-config-suppression-timeout" { return "PreConfigSuppressionTimeout" }
    if yname == "buffer-size" { return "BufferSize" }
    if yname == "source-location" { return "SourceLocation" }
    if yname == "threshold" { return "Threshold" }
    if yname == "alarm-filter-strings" { return "AlarmFilterStrings" }
    return ""
}

func (alarmLogger *Syslog_AlarmLogger) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-alarm-logger-cfg:alarm-logger"
}

func (alarmLogger *Syslog_AlarmLogger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-filter-strings" {
        return &alarmLogger.AlarmFilterStrings
    }
    return nil
}

func (alarmLogger *Syslog_AlarmLogger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["alarm-filter-strings"] = &alarmLogger.AlarmFilterStrings
    return children
}

func (alarmLogger *Syslog_AlarmLogger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pre-config-suppression"] = alarmLogger.PreConfigSuppression
    leafs["severity-level"] = alarmLogger.SeverityLevel
    leafs["pre-config-suppression-timeout"] = alarmLogger.PreConfigSuppressionTimeout
    leafs["buffer-size"] = alarmLogger.BufferSize
    leafs["source-location"] = alarmLogger.SourceLocation
    leafs["threshold"] = alarmLogger.Threshold
    return leafs
}

func (alarmLogger *Syslog_AlarmLogger) GetBundleName() string { return "cisco_ios_xr" }

func (alarmLogger *Syslog_AlarmLogger) GetYangName() string { return "alarm-logger" }

func (alarmLogger *Syslog_AlarmLogger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmLogger *Syslog_AlarmLogger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmLogger *Syslog_AlarmLogger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmLogger *Syslog_AlarmLogger) SetParent(parent types.Entity) { alarmLogger.parent = parent }

func (alarmLogger *Syslog_AlarmLogger) GetParent() types.Entity { return alarmLogger.parent }

func (alarmLogger *Syslog_AlarmLogger) GetParentYangName() string { return "syslog" }

// Syslog_AlarmLogger_AlarmFilterStrings
// List of filter strings
type Syslog_AlarmLogger_AlarmFilterStrings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match string to filter alarms. The type is slice of
    // Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString.
    AlarmFilterString []Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString
}

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetFilter() yfilter.YFilter { return alarmFilterStrings.YFilter }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) SetFilter(yf yfilter.YFilter) { alarmFilterStrings.YFilter = yf }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetGoName(yname string) string {
    if yname == "alarm-filter-string" { return "AlarmFilterString" }
    return ""
}

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetSegmentPath() string {
    return "alarm-filter-strings"
}

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-filter-string" {
        for _, c := range alarmFilterStrings.AlarmFilterString {
            if alarmFilterStrings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString{}
        alarmFilterStrings.AlarmFilterString = append(alarmFilterStrings.AlarmFilterString, child)
        return &alarmFilterStrings.AlarmFilterString[len(alarmFilterStrings.AlarmFilterString)-1]
    }
    return nil
}

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alarmFilterStrings.AlarmFilterString {
        children[alarmFilterStrings.AlarmFilterString[i].GetSegmentPath()] = &alarmFilterStrings.AlarmFilterString[i]
    }
    return children
}

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetBundleName() string { return "cisco_ios_xr" }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetYangName() string { return "alarm-filter-strings" }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) SetParent(parent types.Entity) { alarmFilterStrings.parent = parent }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetParent() types.Entity { return alarmFilterStrings.parent }

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetParentYangName() string { return "alarm-logger" }

// Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString
// Match string to filter alarms
type Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Filter String. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    FilterString interface{}
}

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetFilter() yfilter.YFilter { return alarmFilterString.YFilter }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) SetFilter(yf yfilter.YFilter) { alarmFilterString.YFilter = yf }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetGoName(yname string) string {
    if yname == "filter-string" { return "FilterString" }
    return ""
}

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetSegmentPath() string {
    return "alarm-filter-string" + "[filter-string='" + fmt.Sprintf("%v", alarmFilterString.FilterString) + "']"
}

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filter-string"] = alarmFilterString.FilterString
    return leafs
}

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetBundleName() string { return "cisco_ios_xr" }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetYangName() string { return "alarm-filter-string" }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) SetParent(parent types.Entity) { alarmFilterString.parent = parent }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetParent() types.Entity { return alarmFilterString.parent }

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetParentYangName() string { return "alarm-filter-strings" }

// Syslog_Correlator
// Configure properties of the event correlator
type Syslog_Correlator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure size of the correlator buffer. The type is interface{} with
    // range: 1024..52428800.
    BufferSize interface{}

    // Table of configured rules.
    Rules Syslog_Correlator_Rules

    // Table of configured rulesets.
    RuleSets Syslog_Correlator_RuleSets
}

func (correlator *Syslog_Correlator) GetFilter() yfilter.YFilter { return correlator.YFilter }

func (correlator *Syslog_Correlator) SetFilter(yf yfilter.YFilter) { correlator.YFilter = yf }

func (correlator *Syslog_Correlator) GetGoName(yname string) string {
    if yname == "buffer-size" { return "BufferSize" }
    if yname == "rules" { return "Rules" }
    if yname == "rule-sets" { return "RuleSets" }
    return ""
}

func (correlator *Syslog_Correlator) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-correlator-cfg:correlator"
}

func (correlator *Syslog_Correlator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rules" {
        return &correlator.Rules
    }
    if childYangName == "rule-sets" {
        return &correlator.RuleSets
    }
    return nil
}

func (correlator *Syslog_Correlator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rules"] = &correlator.Rules
    children["rule-sets"] = &correlator.RuleSets
    return children
}

func (correlator *Syslog_Correlator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["buffer-size"] = correlator.BufferSize
    return leafs
}

func (correlator *Syslog_Correlator) GetBundleName() string { return "cisco_ios_xr" }

func (correlator *Syslog_Correlator) GetYangName() string { return "correlator" }

func (correlator *Syslog_Correlator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (correlator *Syslog_Correlator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (correlator *Syslog_Correlator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (correlator *Syslog_Correlator) SetParent(parent types.Entity) { correlator.parent = parent }

func (correlator *Syslog_Correlator) GetParent() types.Entity { return correlator.parent }

func (correlator *Syslog_Correlator) GetParentYangName() string { return "syslog" }

// Syslog_Correlator_Rules
// Table of configured rules
type Syslog_Correlator_Rules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rule name. The type is slice of Syslog_Correlator_Rules_Rule.
    Rule []Syslog_Correlator_Rules_Rule
}

func (rules *Syslog_Correlator_Rules) GetFilter() yfilter.YFilter { return rules.YFilter }

func (rules *Syslog_Correlator_Rules) SetFilter(yf yfilter.YFilter) { rules.YFilter = yf }

func (rules *Syslog_Correlator_Rules) GetGoName(yname string) string {
    if yname == "rule" { return "Rule" }
    return ""
}

func (rules *Syslog_Correlator_Rules) GetSegmentPath() string {
    return "rules"
}

func (rules *Syslog_Correlator_Rules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule" {
        for _, c := range rules.Rule {
            if rules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_Rules_Rule{}
        rules.Rule = append(rules.Rule, child)
        return &rules.Rule[len(rules.Rule)-1]
    }
    return nil
}

func (rules *Syslog_Correlator_Rules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rules.Rule {
        children[rules.Rule[i].GetSegmentPath()] = &rules.Rule[i]
    }
    return children
}

func (rules *Syslog_Correlator_Rules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rules *Syslog_Correlator_Rules) GetBundleName() string { return "cisco_ios_xr" }

func (rules *Syslog_Correlator_Rules) GetYangName() string { return "rules" }

func (rules *Syslog_Correlator_Rules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rules *Syslog_Correlator_Rules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rules *Syslog_Correlator_Rules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rules *Syslog_Correlator_Rules) SetParent(parent types.Entity) { rules.parent = parent }

func (rules *Syslog_Correlator_Rules) GetParent() types.Entity { return rules.parent }

func (rules *Syslog_Correlator_Rules) GetParentYangName() string { return "correlator" }

// Syslog_Correlator_Rules_Rule
// Rule name
type Syslog_Correlator_Rules_Rule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Name interface{}

    // Configure a specified correlation rule.
    Definition Syslog_Correlator_Rules_Rule_Definition

    // The Non-Stateful Rule Type.
    NonStateful Syslog_Correlator_Rules_Rule_NonStateful

    // The Stateful Rule Type.
    Stateful Syslog_Correlator_Rules_Rule_Stateful

    // Apply the Rules.
    ApplyTo Syslog_Correlator_Rules_Rule_ApplyTo

    // Applied to the Rule or Ruleset.
    AppliedTo Syslog_Correlator_Rules_Rule_AppliedTo
}

func (rule *Syslog_Correlator_Rules_Rule) GetFilter() yfilter.YFilter { return rule.YFilter }

func (rule *Syslog_Correlator_Rules_Rule) SetFilter(yf yfilter.YFilter) { rule.YFilter = yf }

func (rule *Syslog_Correlator_Rules_Rule) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "definition" { return "Definition" }
    if yname == "non-stateful" { return "NonStateful" }
    if yname == "stateful" { return "Stateful" }
    if yname == "apply-to" { return "ApplyTo" }
    if yname == "applied-to" { return "AppliedTo" }
    return ""
}

func (rule *Syslog_Correlator_Rules_Rule) GetSegmentPath() string {
    return "rule" + "[name='" + fmt.Sprintf("%v", rule.Name) + "']"
}

func (rule *Syslog_Correlator_Rules_Rule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "definition" {
        return &rule.Definition
    }
    if childYangName == "non-stateful" {
        return &rule.NonStateful
    }
    if childYangName == "stateful" {
        return &rule.Stateful
    }
    if childYangName == "apply-to" {
        return &rule.ApplyTo
    }
    if childYangName == "applied-to" {
        return &rule.AppliedTo
    }
    return nil
}

func (rule *Syslog_Correlator_Rules_Rule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["definition"] = &rule.Definition
    children["non-stateful"] = &rule.NonStateful
    children["stateful"] = &rule.Stateful
    children["apply-to"] = &rule.ApplyTo
    children["applied-to"] = &rule.AppliedTo
    return children
}

func (rule *Syslog_Correlator_Rules_Rule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = rule.Name
    return leafs
}

func (rule *Syslog_Correlator_Rules_Rule) GetBundleName() string { return "cisco_ios_xr" }

func (rule *Syslog_Correlator_Rules_Rule) GetYangName() string { return "rule" }

func (rule *Syslog_Correlator_Rules_Rule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rule *Syslog_Correlator_Rules_Rule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rule *Syslog_Correlator_Rules_Rule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rule *Syslog_Correlator_Rules_Rule) SetParent(parent types.Entity) { rule.parent = parent }

func (rule *Syslog_Correlator_Rules_Rule) GetParent() types.Entity { return rule.parent }

func (rule *Syslog_Correlator_Rules_Rule) GetParentYangName() string { return "rules" }

// Syslog_Correlator_Rules_Rule_Definition
// Configure a specified correlation rule
type Syslog_Correlator_Rules_Rule_Definition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout (time the rule is to be active) in milliseconds. The type is
    // interface{} with range: 1..7200000.
    Timeout interface{}

    // Root message category name. The type is string.
    CategoryNameEntry1 interface{}

    // Root message group name. The type is string.
    GroupNameEntry1 interface{}

    // Root message code. The type is string.
    MessageCodeEntry1 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry2 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry2 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry2 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry3 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry3 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry3 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry4 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry4 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry4 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry5 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry5 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry5 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry6 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry6 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry6 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry7 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry7 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry7 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry8 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry8 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry8 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry9 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry9 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry9 interface{}

    // Correlated message category name. The type is string.
    CategoryNameEntry10 interface{}

    // Correlated message group name. The type is string.
    GroupNameEntry10 interface{}

    // Correlated message code. The type is string.
    MessageCodeEntry10 interface{}
}

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetFilter() yfilter.YFilter { return definition.YFilter }

func (definition *Syslog_Correlator_Rules_Rule_Definition) SetFilter(yf yfilter.YFilter) { definition.YFilter = yf }

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    if yname == "category-name-entry1" { return "CategoryNameEntry1" }
    if yname == "group-name-entry1" { return "GroupNameEntry1" }
    if yname == "message-code-entry1" { return "MessageCodeEntry1" }
    if yname == "category-name-entry2" { return "CategoryNameEntry2" }
    if yname == "group-name-entry2" { return "GroupNameEntry2" }
    if yname == "message-code-entry2" { return "MessageCodeEntry2" }
    if yname == "category-name-entry3" { return "CategoryNameEntry3" }
    if yname == "group-name-entry3" { return "GroupNameEntry3" }
    if yname == "message-code-entry3" { return "MessageCodeEntry3" }
    if yname == "category-name-entry4" { return "CategoryNameEntry4" }
    if yname == "group-name-entry4" { return "GroupNameEntry4" }
    if yname == "message-code-entry4" { return "MessageCodeEntry4" }
    if yname == "category-name-entry5" { return "CategoryNameEntry5" }
    if yname == "group-name-entry5" { return "GroupNameEntry5" }
    if yname == "message-code-entry5" { return "MessageCodeEntry5" }
    if yname == "category-name-entry6" { return "CategoryNameEntry6" }
    if yname == "group-name-entry6" { return "GroupNameEntry6" }
    if yname == "message-code-entry6" { return "MessageCodeEntry6" }
    if yname == "category-name-entry7" { return "CategoryNameEntry7" }
    if yname == "group-name-entry7" { return "GroupNameEntry7" }
    if yname == "message-code-entry7" { return "MessageCodeEntry7" }
    if yname == "category-name-entry8" { return "CategoryNameEntry8" }
    if yname == "group-name-entry8" { return "GroupNameEntry8" }
    if yname == "message-code-entry8" { return "MessageCodeEntry8" }
    if yname == "category-name-entry9" { return "CategoryNameEntry9" }
    if yname == "group-name-entry9" { return "GroupNameEntry9" }
    if yname == "message-code-entry9" { return "MessageCodeEntry9" }
    if yname == "category-name-entry10" { return "CategoryNameEntry10" }
    if yname == "group-name-entry10" { return "GroupNameEntry10" }
    if yname == "message-code-entry10" { return "MessageCodeEntry10" }
    return ""
}

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetSegmentPath() string {
    return "definition"
}

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = definition.Timeout
    leafs["category-name-entry1"] = definition.CategoryNameEntry1
    leafs["group-name-entry1"] = definition.GroupNameEntry1
    leafs["message-code-entry1"] = definition.MessageCodeEntry1
    leafs["category-name-entry2"] = definition.CategoryNameEntry2
    leafs["group-name-entry2"] = definition.GroupNameEntry2
    leafs["message-code-entry2"] = definition.MessageCodeEntry2
    leafs["category-name-entry3"] = definition.CategoryNameEntry3
    leafs["group-name-entry3"] = definition.GroupNameEntry3
    leafs["message-code-entry3"] = definition.MessageCodeEntry3
    leafs["category-name-entry4"] = definition.CategoryNameEntry4
    leafs["group-name-entry4"] = definition.GroupNameEntry4
    leafs["message-code-entry4"] = definition.MessageCodeEntry4
    leafs["category-name-entry5"] = definition.CategoryNameEntry5
    leafs["group-name-entry5"] = definition.GroupNameEntry5
    leafs["message-code-entry5"] = definition.MessageCodeEntry5
    leafs["category-name-entry6"] = definition.CategoryNameEntry6
    leafs["group-name-entry6"] = definition.GroupNameEntry6
    leafs["message-code-entry6"] = definition.MessageCodeEntry6
    leafs["category-name-entry7"] = definition.CategoryNameEntry7
    leafs["group-name-entry7"] = definition.GroupNameEntry7
    leafs["message-code-entry7"] = definition.MessageCodeEntry7
    leafs["category-name-entry8"] = definition.CategoryNameEntry8
    leafs["group-name-entry8"] = definition.GroupNameEntry8
    leafs["message-code-entry8"] = definition.MessageCodeEntry8
    leafs["category-name-entry9"] = definition.CategoryNameEntry9
    leafs["group-name-entry9"] = definition.GroupNameEntry9
    leafs["message-code-entry9"] = definition.MessageCodeEntry9
    leafs["category-name-entry10"] = definition.CategoryNameEntry10
    leafs["group-name-entry10"] = definition.GroupNameEntry10
    leafs["message-code-entry10"] = definition.MessageCodeEntry10
    return leafs
}

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetBundleName() string { return "cisco_ios_xr" }

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetYangName() string { return "definition" }

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (definition *Syslog_Correlator_Rules_Rule_Definition) SetParent(parent types.Entity) { definition.parent = parent }

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetParent() types.Entity { return definition.parent }

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetParentYangName() string { return "rule" }

// Syslog_Correlator_Rules_Rule_NonStateful
// The Non-Stateful Rule Type
type Syslog_Correlator_Rules_Rule_NonStateful struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable correlation on alarm context. The type is interface{}.
    ContextCorrelation interface{}

    // Rootcause Timeout (time to wait for rootcause) in milliseconds. The type is
    // interface{} with range: 1..7200000. Units are millisecond.
    TimeoutRootCause interface{}

    // Timeout (time to wait for active correlation) in milliseconds. The type is
    // interface{} with range: 1..7200000. Units are millisecond.
    Timeout interface{}

    // Table of configured non-rootcause.
    NonRootCauses Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses

    // The root cause.
    RootCause Syslog_Correlator_Rules_Rule_NonStateful_RootCause
}

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetFilter() yfilter.YFilter { return nonStateful.YFilter }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) SetFilter(yf yfilter.YFilter) { nonStateful.YFilter = yf }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetGoName(yname string) string {
    if yname == "context-correlation" { return "ContextCorrelation" }
    if yname == "timeout-root-cause" { return "TimeoutRootCause" }
    if yname == "timeout" { return "Timeout" }
    if yname == "non-root-causes" { return "NonRootCauses" }
    if yname == "root-cause" { return "RootCause" }
    return ""
}

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetSegmentPath() string {
    return "non-stateful"
}

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "non-root-causes" {
        return &nonStateful.NonRootCauses
    }
    if childYangName == "root-cause" {
        return &nonStateful.RootCause
    }
    return nil
}

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["non-root-causes"] = &nonStateful.NonRootCauses
    children["root-cause"] = &nonStateful.RootCause
    return children
}

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context-correlation"] = nonStateful.ContextCorrelation
    leafs["timeout-root-cause"] = nonStateful.TimeoutRootCause
    leafs["timeout"] = nonStateful.Timeout
    return leafs
}

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetBundleName() string { return "cisco_ios_xr" }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetYangName() string { return "non-stateful" }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) SetParent(parent types.Entity) { nonStateful.parent = parent }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetParent() types.Entity { return nonStateful.parent }

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetParentYangName() string { return "rule" }

// Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses
// Table of configured non-rootcause
type Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A non-rootcause. The type is slice of
    // Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause.
    NonRootCause []Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetFilter() yfilter.YFilter { return nonRootCauses.YFilter }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) SetFilter(yf yfilter.YFilter) { nonRootCauses.YFilter = yf }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetGoName(yname string) string {
    if yname == "non-root-cause" { return "NonRootCause" }
    return ""
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetSegmentPath() string {
    return "non-root-causes"
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "non-root-cause" {
        for _, c := range nonRootCauses.NonRootCause {
            if nonRootCauses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause{}
        nonRootCauses.NonRootCause = append(nonRootCauses.NonRootCause, child)
        return &nonRootCauses.NonRootCause[len(nonRootCauses.NonRootCause)-1]
    }
    return nil
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nonRootCauses.NonRootCause {
        children[nonRootCauses.NonRootCause[i].GetSegmentPath()] = &nonRootCauses.NonRootCause[i]
    }
    return children
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetBundleName() string { return "cisco_ios_xr" }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetYangName() string { return "non-root-causes" }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) SetParent(parent types.Entity) { nonRootCauses.parent = parent }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetParent() types.Entity { return nonRootCauses.parent }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetParentYangName() string { return "non-stateful" }

// Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause
// A non-rootcause
type Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Correlated message category. The type is string.
    Category interface{}

    // This attribute is a key. Correlated message group. The type is string.
    Group interface{}

    // This attribute is a key. Correlated message code. The type is string.
    MessageCode interface{}
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetFilter() yfilter.YFilter { return nonRootCause.YFilter }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) SetFilter(yf yfilter.YFilter) { nonRootCause.YFilter = yf }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "message-code" { return "MessageCode" }
    return ""
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetSegmentPath() string {
    return "non-root-cause" + "[category='" + fmt.Sprintf("%v", nonRootCause.Category) + "']" + "[group='" + fmt.Sprintf("%v", nonRootCause.Group) + "']" + "[message-code='" + fmt.Sprintf("%v", nonRootCause.MessageCode) + "']"
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = nonRootCause.Category
    leafs["group"] = nonRootCause.Group
    leafs["message-code"] = nonRootCause.MessageCode
    return leafs
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetBundleName() string { return "cisco_ios_xr" }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetYangName() string { return "non-root-cause" }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) SetParent(parent types.Entity) { nonRootCause.parent = parent }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetParent() types.Entity { return nonRootCause.parent }

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetParentYangName() string { return "non-root-causes" }

// Syslog_Correlator_Rules_Rule_NonStateful_RootCause
// The root cause
type Syslog_Correlator_Rules_Rule_NonStateful_RootCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Root message category. The type is string.
    Category interface{}

    // Root message group. The type is string.
    Group interface{}

    // Root message code. The type is string.
    MessageCode interface{}
}

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetFilter() yfilter.YFilter { return rootCause.YFilter }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) SetFilter(yf yfilter.YFilter) { rootCause.YFilter = yf }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "message-code" { return "MessageCode" }
    return ""
}

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetSegmentPath() string {
    return "root-cause"
}

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = rootCause.Category
    leafs["group"] = rootCause.Group
    leafs["message-code"] = rootCause.MessageCode
    return leafs
}

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetBundleName() string { return "cisco_ios_xr" }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetYangName() string { return "root-cause" }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) SetParent(parent types.Entity) { rootCause.parent = parent }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetParent() types.Entity { return rootCause.parent }

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetParentYangName() string { return "non-stateful" }

// Syslog_Correlator_Rules_Rule_Stateful
// The Stateful Rule Type
type Syslog_Correlator_Rules_Rule_Stateful struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable reparent of alarm on rootcause alarm clear. The type is interface{}.
    Reparent interface{}

    // Enable reissue of non-bistate alarms on rootcause alarm clear. The type is
    // interface{}.
    Reissue interface{}

    // Enable correlation on alarm context. The type is interface{}.
    ContextCorrelation interface{}

    // Rootcause Timeout (time to wait for rootcause) in milliseconds. The type is
    // interface{} with range: 1..7200000. Units are millisecond.
    TimeoutRootCause interface{}

    // Timeout (time to wait for active correlation) in milliseconds. The type is
    // interface{} with range: 1..7200000. Units are millisecond.
    Timeout interface{}

    // Table of configured non-rootcause.
    NonRootCauses Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses

    // The root cause.
    RootCause Syslog_Correlator_Rules_Rule_Stateful_RootCause
}

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetFilter() yfilter.YFilter { return stateful.YFilter }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) SetFilter(yf yfilter.YFilter) { stateful.YFilter = yf }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetGoName(yname string) string {
    if yname == "reparent" { return "Reparent" }
    if yname == "reissue" { return "Reissue" }
    if yname == "context-correlation" { return "ContextCorrelation" }
    if yname == "timeout-root-cause" { return "TimeoutRootCause" }
    if yname == "timeout" { return "Timeout" }
    if yname == "non-root-causes" { return "NonRootCauses" }
    if yname == "root-cause" { return "RootCause" }
    return ""
}

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetSegmentPath() string {
    return "stateful"
}

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "non-root-causes" {
        return &stateful.NonRootCauses
    }
    if childYangName == "root-cause" {
        return &stateful.RootCause
    }
    return nil
}

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["non-root-causes"] = &stateful.NonRootCauses
    children["root-cause"] = &stateful.RootCause
    return children
}

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reparent"] = stateful.Reparent
    leafs["reissue"] = stateful.Reissue
    leafs["context-correlation"] = stateful.ContextCorrelation
    leafs["timeout-root-cause"] = stateful.TimeoutRootCause
    leafs["timeout"] = stateful.Timeout
    return leafs
}

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetBundleName() string { return "cisco_ios_xr" }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetYangName() string { return "stateful" }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) SetParent(parent types.Entity) { stateful.parent = parent }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetParent() types.Entity { return stateful.parent }

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetParentYangName() string { return "rule" }

// Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses
// Table of configured non-rootcause
type Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A non-rootcause. The type is slice of
    // Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause.
    NonRootCause []Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetFilter() yfilter.YFilter { return nonRootCauses.YFilter }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) SetFilter(yf yfilter.YFilter) { nonRootCauses.YFilter = yf }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetGoName(yname string) string {
    if yname == "non-root-cause" { return "NonRootCause" }
    return ""
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetSegmentPath() string {
    return "non-root-causes"
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "non-root-cause" {
        for _, c := range nonRootCauses.NonRootCause {
            if nonRootCauses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause{}
        nonRootCauses.NonRootCause = append(nonRootCauses.NonRootCause, child)
        return &nonRootCauses.NonRootCause[len(nonRootCauses.NonRootCause)-1]
    }
    return nil
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nonRootCauses.NonRootCause {
        children[nonRootCauses.NonRootCause[i].GetSegmentPath()] = &nonRootCauses.NonRootCause[i]
    }
    return children
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetBundleName() string { return "cisco_ios_xr" }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetYangName() string { return "non-root-causes" }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) SetParent(parent types.Entity) { nonRootCauses.parent = parent }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetParent() types.Entity { return nonRootCauses.parent }

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetParentYangName() string { return "stateful" }

// Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause
// A non-rootcause
type Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Correlated message category. The type is string.
    Category interface{}

    // This attribute is a key. Correlated message group. The type is string.
    Group interface{}

    // This attribute is a key. Correlated message code. The type is string.
    MessageCode interface{}
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetFilter() yfilter.YFilter { return nonRootCause.YFilter }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) SetFilter(yf yfilter.YFilter) { nonRootCause.YFilter = yf }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "message-code" { return "MessageCode" }
    return ""
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetSegmentPath() string {
    return "non-root-cause" + "[category='" + fmt.Sprintf("%v", nonRootCause.Category) + "']" + "[group='" + fmt.Sprintf("%v", nonRootCause.Group) + "']" + "[message-code='" + fmt.Sprintf("%v", nonRootCause.MessageCode) + "']"
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = nonRootCause.Category
    leafs["group"] = nonRootCause.Group
    leafs["message-code"] = nonRootCause.MessageCode
    return leafs
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetBundleName() string { return "cisco_ios_xr" }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetYangName() string { return "non-root-cause" }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) SetParent(parent types.Entity) { nonRootCause.parent = parent }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetParent() types.Entity { return nonRootCause.parent }

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetParentYangName() string { return "non-root-causes" }

// Syslog_Correlator_Rules_Rule_Stateful_RootCause
// The root cause
type Syslog_Correlator_Rules_Rule_Stateful_RootCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Root message category. The type is string.
    Category interface{}

    // Root message group. The type is string.
    Group interface{}

    // Root message code. The type is string.
    MessageCode interface{}
}

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetFilter() yfilter.YFilter { return rootCause.YFilter }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) SetFilter(yf yfilter.YFilter) { rootCause.YFilter = yf }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "message-code" { return "MessageCode" }
    return ""
}

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetSegmentPath() string {
    return "root-cause"
}

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = rootCause.Category
    leafs["group"] = rootCause.Group
    leafs["message-code"] = rootCause.MessageCode
    return leafs
}

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetBundleName() string { return "cisco_ios_xr" }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetYangName() string { return "root-cause" }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) SetParent(parent types.Entity) { rootCause.parent = parent }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetParent() types.Entity { return rootCause.parent }

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetParentYangName() string { return "stateful" }

// Syslog_Correlator_Rules_Rule_ApplyTo
// Apply the Rules
type Syslog_Correlator_Rules_Rule_ApplyTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Apply the rule to all of the router. The type is interface{}.
    AllOfRouter interface{}

    // Apply rule to a specified list of contexts, e.g. interfaces.
    Contexts Syslog_Correlator_Rules_Rule_ApplyTo_Contexts

    // Apply rule to a specified list of Locations.
    Locations Syslog_Correlator_Rules_Rule_ApplyTo_Locations
}

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetFilter() yfilter.YFilter { return applyTo.YFilter }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) SetFilter(yf yfilter.YFilter) { applyTo.YFilter = yf }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetGoName(yname string) string {
    if yname == "all-of-router" { return "AllOfRouter" }
    if yname == "contexts" { return "Contexts" }
    if yname == "locations" { return "Locations" }
    return ""
}

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetSegmentPath() string {
    return "apply-to"
}

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contexts" {
        return &applyTo.Contexts
    }
    if childYangName == "locations" {
        return &applyTo.Locations
    }
    return nil
}

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["contexts"] = &applyTo.Contexts
    children["locations"] = &applyTo.Locations
    return children
}

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all-of-router"] = applyTo.AllOfRouter
    return leafs
}

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetBundleName() string { return "cisco_ios_xr" }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetYangName() string { return "apply-to" }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) SetParent(parent types.Entity) { applyTo.parent = parent }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetParent() types.Entity { return applyTo.parent }

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetParentYangName() string { return "rule" }

// Syslog_Correlator_Rules_Rule_ApplyTo_Contexts
// Apply rule to a specified list of contexts,
// e.g. interfaces
type Syslog_Correlator_Rules_Rule_ApplyTo_Contexts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One or more context names. The type is slice of string.
    Context []interface{}
}

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetFilter() yfilter.YFilter { return contexts.YFilter }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) SetFilter(yf yfilter.YFilter) { contexts.YFilter = yf }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    return ""
}

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetSegmentPath() string {
    return "contexts"
}

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context"] = contexts.Context
    return leafs
}

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetBundleName() string { return "cisco_ios_xr" }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetYangName() string { return "contexts" }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) SetParent(parent types.Entity) { contexts.parent = parent }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetParent() types.Entity { return contexts.parent }

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetParentYangName() string { return "apply-to" }

// Syslog_Correlator_Rules_Rule_ApplyTo_Locations
// Apply rule to a specified list of Locations
type Syslog_Correlator_Rules_Rule_ApplyTo_Locations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One or more Locations. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location []interface{}
}

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetFilter() yfilter.YFilter { return locations.YFilter }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) SetFilter(yf yfilter.YFilter) { locations.YFilter = yf }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetSegmentPath() string {
    return "locations"
}

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = locations.Location
    return leafs
}

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetBundleName() string { return "cisco_ios_xr" }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetYangName() string { return "locations" }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) SetParent(parent types.Entity) { locations.parent = parent }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetParent() types.Entity { return locations.parent }

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetParentYangName() string { return "apply-to" }

// Syslog_Correlator_Rules_Rule_AppliedTo
// Applied to the Rule or Ruleset
type Syslog_Correlator_Rules_Rule_AppliedTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Apply to all of the router. The type is interface{}.
    All interface{}

    // Table of configured contexts to apply.
    Contexts Syslog_Correlator_Rules_Rule_AppliedTo_Contexts

    // Table of configured locations to apply.
    Locations Syslog_Correlator_Rules_Rule_AppliedTo_Locations
}

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetFilter() yfilter.YFilter { return appliedTo.YFilter }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) SetFilter(yf yfilter.YFilter) { appliedTo.YFilter = yf }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    if yname == "contexts" { return "Contexts" }
    if yname == "locations" { return "Locations" }
    return ""
}

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetSegmentPath() string {
    return "applied-to"
}

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contexts" {
        return &appliedTo.Contexts
    }
    if childYangName == "locations" {
        return &appliedTo.Locations
    }
    return nil
}

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["contexts"] = &appliedTo.Contexts
    children["locations"] = &appliedTo.Locations
    return children
}

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all"] = appliedTo.All
    return leafs
}

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetBundleName() string { return "cisco_ios_xr" }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetYangName() string { return "applied-to" }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) SetParent(parent types.Entity) { appliedTo.parent = parent }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetParent() types.Entity { return appliedTo.parent }

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetParentYangName() string { return "rule" }

// Syslog_Correlator_Rules_Rule_AppliedTo_Contexts
// Table of configured contexts to apply
type Syslog_Correlator_Rules_Rule_AppliedTo_Contexts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A context. The type is slice of
    // Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context.
    Context []Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context
}

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetFilter() yfilter.YFilter { return contexts.YFilter }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) SetFilter(yf yfilter.YFilter) { contexts.YFilter = yf }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    return ""
}

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetSegmentPath() string {
    return "contexts"
}

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context" {
        for _, c := range contexts.Context {
            if contexts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context{}
        contexts.Context = append(contexts.Context, child)
        return &contexts.Context[len(contexts.Context)-1]
    }
    return nil
}

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contexts.Context {
        children[contexts.Context[i].GetSegmentPath()] = &contexts.Context[i]
    }
    return children
}

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetBundleName() string { return "cisco_ios_xr" }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetYangName() string { return "contexts" }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) SetParent(parent types.Entity) { contexts.parent = parent }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetParent() types.Entity { return contexts.parent }

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetParentYangName() string { return "applied-to" }

// Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context
// A context
type Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context. The type is string with length: 1..32.
    Context interface{}
}

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetFilter() yfilter.YFilter { return context.YFilter }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) SetFilter(yf yfilter.YFilter) { context.YFilter = yf }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    return ""
}

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetSegmentPath() string {
    return "context" + "[context='" + fmt.Sprintf("%v", context.Context) + "']"
}

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context"] = context.Context
    return leafs
}

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetBundleName() string { return "cisco_ios_xr" }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetYangName() string { return "context" }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) SetParent(parent types.Entity) { context.parent = parent }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetParent() types.Entity { return context.parent }

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetParentYangName() string { return "contexts" }

// Syslog_Correlator_Rules_Rule_AppliedTo_Locations
// Table of configured locations to apply
type Syslog_Correlator_Rules_Rule_AppliedTo_Locations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A location. The type is slice of
    // Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location.
    Location []Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location
}

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetFilter() yfilter.YFilter { return locations.YFilter }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) SetFilter(yf yfilter.YFilter) { locations.YFilter = yf }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetSegmentPath() string {
    return "locations"
}

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "location" {
        for _, c := range locations.Location {
            if locations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location{}
        locations.Location = append(locations.Location, child)
        return &locations.Location[len(locations.Location)-1]
    }
    return nil
}

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locations.Location {
        children[locations.Location[i].GetSegmentPath()] = &locations.Location[i]
    }
    return children
}

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetBundleName() string { return "cisco_ios_xr" }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetYangName() string { return "locations" }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) SetParent(parent types.Entity) { locations.parent = parent }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetParent() types.Entity { return locations.parent }

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetParentYangName() string { return "applied-to" }

// Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location
// A location
type Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}
}

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetFilter() yfilter.YFilter { return location.YFilter }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) SetFilter(yf yfilter.YFilter) { location.YFilter = yf }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetSegmentPath() string {
    return "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
}

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = location.Location
    return leafs
}

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetBundleName() string { return "cisco_ios_xr" }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetYangName() string { return "location" }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) SetParent(parent types.Entity) { location.parent = parent }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetParent() types.Entity { return location.parent }

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetParentYangName() string { return "locations" }

// Syslog_Correlator_RuleSets
// Table of configured rulesets
type Syslog_Correlator_RuleSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ruleset name. The type is slice of Syslog_Correlator_RuleSets_RuleSet.
    RuleSet []Syslog_Correlator_RuleSets_RuleSet
}

func (ruleSets *Syslog_Correlator_RuleSets) GetFilter() yfilter.YFilter { return ruleSets.YFilter }

func (ruleSets *Syslog_Correlator_RuleSets) SetFilter(yf yfilter.YFilter) { ruleSets.YFilter = yf }

func (ruleSets *Syslog_Correlator_RuleSets) GetGoName(yname string) string {
    if yname == "rule-set" { return "RuleSet" }
    return ""
}

func (ruleSets *Syslog_Correlator_RuleSets) GetSegmentPath() string {
    return "rule-sets"
}

func (ruleSets *Syslog_Correlator_RuleSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-set" {
        for _, c := range ruleSets.RuleSet {
            if ruleSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_RuleSets_RuleSet{}
        ruleSets.RuleSet = append(ruleSets.RuleSet, child)
        return &ruleSets.RuleSet[len(ruleSets.RuleSet)-1]
    }
    return nil
}

func (ruleSets *Syslog_Correlator_RuleSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSets.RuleSet {
        children[ruleSets.RuleSet[i].GetSegmentPath()] = &ruleSets.RuleSet[i]
    }
    return children
}

func (ruleSets *Syslog_Correlator_RuleSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleSets *Syslog_Correlator_RuleSets) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSets *Syslog_Correlator_RuleSets) GetYangName() string { return "rule-sets" }

func (ruleSets *Syslog_Correlator_RuleSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSets *Syslog_Correlator_RuleSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSets *Syslog_Correlator_RuleSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSets *Syslog_Correlator_RuleSets) SetParent(parent types.Entity) { ruleSets.parent = parent }

func (ruleSets *Syslog_Correlator_RuleSets) GetParent() types.Entity { return ruleSets.parent }

func (ruleSets *Syslog_Correlator_RuleSets) GetParentYangName() string { return "correlator" }

// Syslog_Correlator_RuleSets_RuleSet
// Ruleset name
type Syslog_Correlator_RuleSets_RuleSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Ruleset name. The type is string with length:
    // 1..32.
    Name interface{}

    // Table of configured rulenames.
    Rulenames Syslog_Correlator_RuleSets_RuleSet_Rulenames

    // Applied to the Rule or Ruleset.
    AppliedTo Syslog_Correlator_RuleSets_RuleSet_AppliedTo
}

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetFilter() yfilter.YFilter { return ruleSet.YFilter }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) SetFilter(yf yfilter.YFilter) { ruleSet.YFilter = yf }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "rulenames" { return "Rulenames" }
    if yname == "applied-to" { return "AppliedTo" }
    return ""
}

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetSegmentPath() string {
    return "rule-set" + "[name='" + fmt.Sprintf("%v", ruleSet.Name) + "']"
}

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rulenames" {
        return &ruleSet.Rulenames
    }
    if childYangName == "applied-to" {
        return &ruleSet.AppliedTo
    }
    return nil
}

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rulenames"] = &ruleSet.Rulenames
    children["applied-to"] = &ruleSet.AppliedTo
    return children
}

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ruleSet.Name
    return leafs
}

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetYangName() string { return "rule-set" }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) SetParent(parent types.Entity) { ruleSet.parent = parent }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetParent() types.Entity { return ruleSet.parent }

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetParentYangName() string { return "rule-sets" }

// Syslog_Correlator_RuleSets_RuleSet_Rulenames
// Table of configured rulenames
type Syslog_Correlator_RuleSets_RuleSet_Rulenames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A rulename. The type is slice of
    // Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename.
    Rulename []Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename
}

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetFilter() yfilter.YFilter { return rulenames.YFilter }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) SetFilter(yf yfilter.YFilter) { rulenames.YFilter = yf }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetGoName(yname string) string {
    if yname == "rulename" { return "Rulename" }
    return ""
}

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetSegmentPath() string {
    return "rulenames"
}

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rulename" {
        for _, c := range rulenames.Rulename {
            if rulenames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename{}
        rulenames.Rulename = append(rulenames.Rulename, child)
        return &rulenames.Rulename[len(rulenames.Rulename)-1]
    }
    return nil
}

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rulenames.Rulename {
        children[rulenames.Rulename[i].GetSegmentPath()] = &rulenames.Rulename[i]
    }
    return children
}

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetBundleName() string { return "cisco_ios_xr" }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetYangName() string { return "rulenames" }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) SetParent(parent types.Entity) { rulenames.parent = parent }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetParent() types.Entity { return rulenames.parent }

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetParentYangName() string { return "rule-set" }

// Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename
// A rulename
type Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Rulename interface{}
}

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetFilter() yfilter.YFilter { return rulename.YFilter }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) SetFilter(yf yfilter.YFilter) { rulename.YFilter = yf }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetGoName(yname string) string {
    if yname == "rulename" { return "Rulename" }
    return ""
}

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetSegmentPath() string {
    return "rulename" + "[rulename='" + fmt.Sprintf("%v", rulename.Rulename) + "']"
}

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rulename"] = rulename.Rulename
    return leafs
}

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetBundleName() string { return "cisco_ios_xr" }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetYangName() string { return "rulename" }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) SetParent(parent types.Entity) { rulename.parent = parent }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetParent() types.Entity { return rulename.parent }

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetParentYangName() string { return "rulenames" }

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo
// Applied to the Rule or Ruleset
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Apply to all of the router. The type is interface{}.
    All interface{}

    // Table of configured contexts to apply.
    Contexts Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts

    // Table of configured locations to apply.
    Locations Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations
}

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetFilter() yfilter.YFilter { return appliedTo.YFilter }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) SetFilter(yf yfilter.YFilter) { appliedTo.YFilter = yf }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    if yname == "contexts" { return "Contexts" }
    if yname == "locations" { return "Locations" }
    return ""
}

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetSegmentPath() string {
    return "applied-to"
}

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contexts" {
        return &appliedTo.Contexts
    }
    if childYangName == "locations" {
        return &appliedTo.Locations
    }
    return nil
}

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["contexts"] = &appliedTo.Contexts
    children["locations"] = &appliedTo.Locations
    return children
}

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all"] = appliedTo.All
    return leafs
}

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetBundleName() string { return "cisco_ios_xr" }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetYangName() string { return "applied-to" }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) SetParent(parent types.Entity) { appliedTo.parent = parent }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetParent() types.Entity { return appliedTo.parent }

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetParentYangName() string { return "rule-set" }

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts
// Table of configured contexts to apply
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A context. The type is slice of
    // Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context.
    Context []Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context
}

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetFilter() yfilter.YFilter { return contexts.YFilter }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) SetFilter(yf yfilter.YFilter) { contexts.YFilter = yf }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    return ""
}

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetSegmentPath() string {
    return "contexts"
}

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context" {
        for _, c := range contexts.Context {
            if contexts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context{}
        contexts.Context = append(contexts.Context, child)
        return &contexts.Context[len(contexts.Context)-1]
    }
    return nil
}

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contexts.Context {
        children[contexts.Context[i].GetSegmentPath()] = &contexts.Context[i]
    }
    return children
}

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetBundleName() string { return "cisco_ios_xr" }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetYangName() string { return "contexts" }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) SetParent(parent types.Entity) { contexts.parent = parent }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetParent() types.Entity { return contexts.parent }

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetParentYangName() string { return "applied-to" }

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context
// A context
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context. The type is string with length: 1..32.
    Context interface{}
}

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetFilter() yfilter.YFilter { return context.YFilter }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) SetFilter(yf yfilter.YFilter) { context.YFilter = yf }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    return ""
}

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetSegmentPath() string {
    return "context" + "[context='" + fmt.Sprintf("%v", context.Context) + "']"
}

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context"] = context.Context
    return leafs
}

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetBundleName() string { return "cisco_ios_xr" }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetYangName() string { return "context" }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) SetParent(parent types.Entity) { context.parent = parent }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetParent() types.Entity { return context.parent }

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetParentYangName() string { return "contexts" }

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations
// Table of configured locations to apply
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A location. The type is slice of
    // Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location.
    Location []Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location
}

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetFilter() yfilter.YFilter { return locations.YFilter }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) SetFilter(yf yfilter.YFilter) { locations.YFilter = yf }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetSegmentPath() string {
    return "locations"
}

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "location" {
        for _, c := range locations.Location {
            if locations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location{}
        locations.Location = append(locations.Location, child)
        return &locations.Location[len(locations.Location)-1]
    }
    return nil
}

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locations.Location {
        children[locations.Location[i].GetSegmentPath()] = &locations.Location[i]
    }
    return children
}

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetBundleName() string { return "cisco_ios_xr" }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetYangName() string { return "locations" }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) SetParent(parent types.Entity) { locations.parent = parent }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetParent() types.Entity { return locations.parent }

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetParentYangName() string { return "applied-to" }

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location
// A location
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}
}

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetFilter() yfilter.YFilter { return location.YFilter }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) SetFilter(yf yfilter.YFilter) { location.YFilter = yf }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetSegmentPath() string {
    return "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
}

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = location.Location
    return leafs
}

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetBundleName() string { return "cisco_ios_xr" }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetYangName() string { return "location" }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) SetParent(parent types.Entity) { location.parent = parent }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetParent() types.Entity { return location.parent }

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetParentYangName() string { return "locations" }

// Syslog_Suppression
// Configure properties of the syslog/alarm
// suppression
type Syslog_Suppression struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of configured rules.
    Rules Syslog_Suppression_Rules
}

func (suppression *Syslog_Suppression) GetFilter() yfilter.YFilter { return suppression.YFilter }

func (suppression *Syslog_Suppression) SetFilter(yf yfilter.YFilter) { suppression.YFilter = yf }

func (suppression *Syslog_Suppression) GetGoName(yname string) string {
    if yname == "rules" { return "Rules" }
    return ""
}

func (suppression *Syslog_Suppression) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-correlator-cfg:suppression"
}

func (suppression *Syslog_Suppression) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rules" {
        return &suppression.Rules
    }
    return nil
}

func (suppression *Syslog_Suppression) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rules"] = &suppression.Rules
    return children
}

func (suppression *Syslog_Suppression) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (suppression *Syslog_Suppression) GetBundleName() string { return "cisco_ios_xr" }

func (suppression *Syslog_Suppression) GetYangName() string { return "suppression" }

func (suppression *Syslog_Suppression) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppression *Syslog_Suppression) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppression *Syslog_Suppression) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppression *Syslog_Suppression) SetParent(parent types.Entity) { suppression.parent = parent }

func (suppression *Syslog_Suppression) GetParent() types.Entity { return suppression.parent }

func (suppression *Syslog_Suppression) GetParentYangName() string { return "syslog" }

// Syslog_Suppression_Rules
// Table of configured rules
type Syslog_Suppression_Rules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rule name. The type is slice of Syslog_Suppression_Rules_Rule.
    Rule []Syslog_Suppression_Rules_Rule
}

func (rules *Syslog_Suppression_Rules) GetFilter() yfilter.YFilter { return rules.YFilter }

func (rules *Syslog_Suppression_Rules) SetFilter(yf yfilter.YFilter) { rules.YFilter = yf }

func (rules *Syslog_Suppression_Rules) GetGoName(yname string) string {
    if yname == "rule" { return "Rule" }
    return ""
}

func (rules *Syslog_Suppression_Rules) GetSegmentPath() string {
    return "rules"
}

func (rules *Syslog_Suppression_Rules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule" {
        for _, c := range rules.Rule {
            if rules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Suppression_Rules_Rule{}
        rules.Rule = append(rules.Rule, child)
        return &rules.Rule[len(rules.Rule)-1]
    }
    return nil
}

func (rules *Syslog_Suppression_Rules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rules.Rule {
        children[rules.Rule[i].GetSegmentPath()] = &rules.Rule[i]
    }
    return children
}

func (rules *Syslog_Suppression_Rules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rules *Syslog_Suppression_Rules) GetBundleName() string { return "cisco_ios_xr" }

func (rules *Syslog_Suppression_Rules) GetYangName() string { return "rules" }

func (rules *Syslog_Suppression_Rules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rules *Syslog_Suppression_Rules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rules *Syslog_Suppression_Rules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rules *Syslog_Suppression_Rules) SetParent(parent types.Entity) { rules.parent = parent }

func (rules *Syslog_Suppression_Rules) GetParent() types.Entity { return rules.parent }

func (rules *Syslog_Suppression_Rules) GetParentYangName() string { return "suppression" }

// Syslog_Suppression_Rules_Rule
// Rule name
type Syslog_Suppression_Rules_Rule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Name interface{}

    // Suppress all alarms. The type is interface{}.
    AllAlarms interface{}

    // Applied to the Rule.
    AppliedTo Syslog_Suppression_Rules_Rule_AppliedTo

    // Causes of alarms to be suppressed.
    AlarmCauses Syslog_Suppression_Rules_Rule_AlarmCauses
}

func (rule *Syslog_Suppression_Rules_Rule) GetFilter() yfilter.YFilter { return rule.YFilter }

func (rule *Syslog_Suppression_Rules_Rule) SetFilter(yf yfilter.YFilter) { rule.YFilter = yf }

func (rule *Syslog_Suppression_Rules_Rule) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "all-alarms" { return "AllAlarms" }
    if yname == "applied-to" { return "AppliedTo" }
    if yname == "alarm-causes" { return "AlarmCauses" }
    return ""
}

func (rule *Syslog_Suppression_Rules_Rule) GetSegmentPath() string {
    return "rule" + "[name='" + fmt.Sprintf("%v", rule.Name) + "']"
}

func (rule *Syslog_Suppression_Rules_Rule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "applied-to" {
        return &rule.AppliedTo
    }
    if childYangName == "alarm-causes" {
        return &rule.AlarmCauses
    }
    return nil
}

func (rule *Syslog_Suppression_Rules_Rule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["applied-to"] = &rule.AppliedTo
    children["alarm-causes"] = &rule.AlarmCauses
    return children
}

func (rule *Syslog_Suppression_Rules_Rule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = rule.Name
    leafs["all-alarms"] = rule.AllAlarms
    return leafs
}

func (rule *Syslog_Suppression_Rules_Rule) GetBundleName() string { return "cisco_ios_xr" }

func (rule *Syslog_Suppression_Rules_Rule) GetYangName() string { return "rule" }

func (rule *Syslog_Suppression_Rules_Rule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rule *Syslog_Suppression_Rules_Rule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rule *Syslog_Suppression_Rules_Rule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rule *Syslog_Suppression_Rules_Rule) SetParent(parent types.Entity) { rule.parent = parent }

func (rule *Syslog_Suppression_Rules_Rule) GetParent() types.Entity { return rule.parent }

func (rule *Syslog_Suppression_Rules_Rule) GetParentYangName() string { return "rules" }

// Syslog_Suppression_Rules_Rule_AppliedTo
// Applied to the Rule
type Syslog_Suppression_Rules_Rule_AppliedTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Apply to all of the router. The type is interface{}.
    All interface{}

    // Table of configured sources to apply.
    Sources Syslog_Suppression_Rules_Rule_AppliedTo_Sources
}

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetFilter() yfilter.YFilter { return appliedTo.YFilter }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) SetFilter(yf yfilter.YFilter) { appliedTo.YFilter = yf }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    if yname == "sources" { return "Sources" }
    return ""
}

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetSegmentPath() string {
    return "applied-to"
}

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sources" {
        return &appliedTo.Sources
    }
    return nil
}

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sources"] = &appliedTo.Sources
    return children
}

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all"] = appliedTo.All
    return leafs
}

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetBundleName() string { return "cisco_ios_xr" }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetYangName() string { return "applied-to" }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) SetParent(parent types.Entity) { appliedTo.parent = parent }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetParent() types.Entity { return appliedTo.parent }

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetParentYangName() string { return "rule" }

// Syslog_Suppression_Rules_Rule_AppliedTo_Sources
// Table of configured sources to apply
type Syslog_Suppression_Rules_Rule_AppliedTo_Sources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An alarm source. The type is slice of
    // Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source.
    Source []Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source
}

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetFilter() yfilter.YFilter { return sources.YFilter }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) SetFilter(yf yfilter.YFilter) { sources.YFilter = yf }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    return ""
}

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetSegmentPath() string {
    return "sources"
}

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        for _, c := range sources.Source {
            if sources.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source{}
        sources.Source = append(sources.Source, child)
        return &sources.Source[len(sources.Source)-1]
    }
    return nil
}

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sources.Source {
        children[sources.Source[i].GetSegmentPath()] = &sources.Source[i]
    }
    return children
}

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetBundleName() string { return "cisco_ios_xr" }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetYangName() string { return "sources" }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) SetParent(parent types.Entity) { sources.parent = parent }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetParent() types.Entity { return sources.parent }

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetParentYangName() string { return "applied-to" }

// Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source
// An alarm source
type Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Source interface{}
}

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    return ""
}

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetSegmentPath() string {
    return "source" + "[source='" + fmt.Sprintf("%v", source.Source) + "']"
}

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source"] = source.Source
    return leafs
}

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetYangName() string { return "source" }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetParent() types.Entity { return source.parent }

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetParentYangName() string { return "sources" }

// Syslog_Suppression_Rules_Rule_AlarmCauses
// Causes of alarms to be suppressed
type Syslog_Suppression_Rules_Rule_AlarmCauses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category, Group and Code of alarm/syslog to be suppressed. The type is
    // slice of Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause.
    AlarmCause []Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause
}

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetFilter() yfilter.YFilter { return alarmCauses.YFilter }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) SetFilter(yf yfilter.YFilter) { alarmCauses.YFilter = yf }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetGoName(yname string) string {
    if yname == "alarm-cause" { return "AlarmCause" }
    return ""
}

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetSegmentPath() string {
    return "alarm-causes"
}

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-cause" {
        for _, c := range alarmCauses.AlarmCause {
            if alarmCauses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause{}
        alarmCauses.AlarmCause = append(alarmCauses.AlarmCause, child)
        return &alarmCauses.AlarmCause[len(alarmCauses.AlarmCause)-1]
    }
    return nil
}

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alarmCauses.AlarmCause {
        children[alarmCauses.AlarmCause[i].GetSegmentPath()] = &alarmCauses.AlarmCause[i]
    }
    return children
}

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetBundleName() string { return "cisco_ios_xr" }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetYangName() string { return "alarm-causes" }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) SetParent(parent types.Entity) { alarmCauses.parent = parent }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetParent() types.Entity { return alarmCauses.parent }

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetParentYangName() string { return "rule" }

// Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause
// Category, Group and Code of alarm/syslog to
// be suppressed
type Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Category. The type is string with length: 1..32.
    Category interface{}

    // This attribute is a key. Group. The type is string with length: 1..32.
    Group interface{}

    // This attribute is a key. Code. The type is string with length: 1..32.
    Code interface{}
}

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetFilter() yfilter.YFilter { return alarmCause.YFilter }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) SetFilter(yf yfilter.YFilter) { alarmCause.YFilter = yf }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "code" { return "Code" }
    return ""
}

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetSegmentPath() string {
    return "alarm-cause" + "[category='" + fmt.Sprintf("%v", alarmCause.Category) + "']" + "[group='" + fmt.Sprintf("%v", alarmCause.Group) + "']" + "[code='" + fmt.Sprintf("%v", alarmCause.Code) + "']"
}

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = alarmCause.Category
    leafs["group"] = alarmCause.Group
    leafs["code"] = alarmCause.Code
    return leafs
}

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetBundleName() string { return "cisco_ios_xr" }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetYangName() string { return "alarm-cause" }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) SetParent(parent types.Entity) { alarmCause.parent = parent }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetParent() types.Entity { return alarmCause.parent }

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetParentYangName() string { return "alarm-causes" }

