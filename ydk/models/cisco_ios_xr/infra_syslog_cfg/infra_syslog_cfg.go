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
    LogSeverity_error_ LogSeverity = "error"

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
    LoggingLevels_error_ LoggingLevels = "error"

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
    LogMessageSeverity_error_ LogMessageSeverity = "error"

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp debug/log messages configuration.
    Timestamps SyslogService_Timestamps
}

func (syslogService *SyslogService) GetEntityData() *types.CommonEntityData {
    syslogService.EntityData.YFilter = syslogService.YFilter
    syslogService.EntityData.YangName = "syslog-service"
    syslogService.EntityData.BundleName = "cisco_ios_xr"
    syslogService.EntityData.ParentYangName = "Cisco-IOS-XR-infra-syslog-cfg"
    syslogService.EntityData.SegmentPath = "Cisco-IOS-XR-infra-syslog-cfg:syslog-service"
    syslogService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslogService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslogService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslogService.EntityData.Children = types.NewOrderedMap()
    syslogService.EntityData.Children.Append("timestamps", types.YChild{"Timestamps", &syslogService.Timestamps})
    syslogService.EntityData.Leafs = types.NewOrderedMap()

    syslogService.EntityData.YListKeys = []string {}

    return &(syslogService.EntityData)
}

// SyslogService_Timestamps
// Timestamp debug/log messages configuration
type SyslogService_Timestamps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable timestamp debug/log messages. The type is interface{}.
    Enable interface{}

    // Timestamp log messages.
    Log SyslogService_Timestamps_Log

    // Timestamp debug messages.
    Debug SyslogService_Timestamps_Debug
}

func (timestamps *SyslogService_Timestamps) GetEntityData() *types.CommonEntityData {
    timestamps.EntityData.YFilter = timestamps.YFilter
    timestamps.EntityData.YangName = "timestamps"
    timestamps.EntityData.BundleName = "cisco_ios_xr"
    timestamps.EntityData.ParentYangName = "syslog-service"
    timestamps.EntityData.SegmentPath = "timestamps"
    timestamps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timestamps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timestamps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timestamps.EntityData.Children = types.NewOrderedMap()
    timestamps.EntityData.Children.Append("log", types.YChild{"Log", &timestamps.Log})
    timestamps.EntityData.Children.Append("debug", types.YChild{"Debug", &timestamps.Debug})
    timestamps.EntityData.Leafs = types.NewOrderedMap()
    timestamps.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", timestamps.Enable})

    timestamps.EntityData.YListKeys = []string {}

    return &(timestamps.EntityData)
}

// SyslogService_Timestamps_Log
// Timestamp log messages
type SyslogService_Timestamps_Log struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp with systime uptime. The type is interface{}.
    LogUptime interface{}

    // Disable timestamp log messages. The type is interface{}.
    LogTimestampDisable interface{}

    // Timestamp with date and time.
    LogDatetime SyslogService_Timestamps_Log_LogDatetime
}

func (log *SyslogService_Timestamps_Log) GetEntityData() *types.CommonEntityData {
    log.EntityData.YFilter = log.YFilter
    log.EntityData.YangName = "log"
    log.EntityData.BundleName = "cisco_ios_xr"
    log.EntityData.ParentYangName = "timestamps"
    log.EntityData.SegmentPath = "log"
    log.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    log.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    log.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    log.EntityData.Children = types.NewOrderedMap()
    log.EntityData.Children.Append("log-datetime", types.YChild{"LogDatetime", &log.LogDatetime})
    log.EntityData.Leafs = types.NewOrderedMap()
    log.EntityData.Leafs.Append("log-uptime", types.YLeaf{"LogUptime", log.LogUptime})
    log.EntityData.Leafs.Append("log-timestamp-disable", types.YLeaf{"LogTimestampDisable", log.LogTimestampDisable})

    log.EntityData.YListKeys = []string {}

    return &(log.EntityData)
}

// SyslogService_Timestamps_Log_LogDatetime
// Timestamp with date and time
type SyslogService_Timestamps_Log_LogDatetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set timestamp for log message.
    LogDatetimeValue SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue
}

func (logDatetime *SyslogService_Timestamps_Log_LogDatetime) GetEntityData() *types.CommonEntityData {
    logDatetime.EntityData.YFilter = logDatetime.YFilter
    logDatetime.EntityData.YangName = "log-datetime"
    logDatetime.EntityData.BundleName = "cisco_ios_xr"
    logDatetime.EntityData.ParentYangName = "log"
    logDatetime.EntityData.SegmentPath = "log-datetime"
    logDatetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logDatetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logDatetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logDatetime.EntityData.Children = types.NewOrderedMap()
    logDatetime.EntityData.Children.Append("log-datetime-value", types.YChild{"LogDatetimeValue", &logDatetime.LogDatetimeValue})
    logDatetime.EntityData.Leafs = types.NewOrderedMap()

    logDatetime.EntityData.YListKeys = []string {}

    return &(logDatetime.EntityData)
}

// SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue
// Set timestamp for log message
type SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue struct {
    EntityData types.CommonEntityData
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

func (logDatetimeValue *SyslogService_Timestamps_Log_LogDatetime_LogDatetimeValue) GetEntityData() *types.CommonEntityData {
    logDatetimeValue.EntityData.YFilter = logDatetimeValue.YFilter
    logDatetimeValue.EntityData.YangName = "log-datetime-value"
    logDatetimeValue.EntityData.BundleName = "cisco_ios_xr"
    logDatetimeValue.EntityData.ParentYangName = "log-datetime"
    logDatetimeValue.EntityData.SegmentPath = "log-datetime-value"
    logDatetimeValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logDatetimeValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logDatetimeValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logDatetimeValue.EntityData.Children = types.NewOrderedMap()
    logDatetimeValue.EntityData.Leafs = types.NewOrderedMap()
    logDatetimeValue.EntityData.Leafs.Append("time-stamp-value", types.YLeaf{"TimeStampValue", logDatetimeValue.TimeStampValue})
    logDatetimeValue.EntityData.Leafs.Append("msec", types.YLeaf{"Msec", logDatetimeValue.Msec})
    logDatetimeValue.EntityData.Leafs.Append("time-zone", types.YLeaf{"TimeZone", logDatetimeValue.TimeZone})
    logDatetimeValue.EntityData.Leafs.Append("year", types.YLeaf{"Year", logDatetimeValue.Year})

    logDatetimeValue.EntityData.YListKeys = []string {}

    return &(logDatetimeValue.EntityData)
}

// SyslogService_Timestamps_Debug
// Timestamp debug messages
type SyslogService_Timestamps_Debug struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable timestamp debug messages. The type is interface{}.
    DebugTimestampDisable interface{}

    // Timestamp with systime uptime. The type is interface{}.
    DebugUptime interface{}

    // Timestamp with date and time.
    DebugDatetime SyslogService_Timestamps_Debug_DebugDatetime
}

func (debug *SyslogService_Timestamps_Debug) GetEntityData() *types.CommonEntityData {
    debug.EntityData.YFilter = debug.YFilter
    debug.EntityData.YangName = "debug"
    debug.EntityData.BundleName = "cisco_ios_xr"
    debug.EntityData.ParentYangName = "timestamps"
    debug.EntityData.SegmentPath = "debug"
    debug.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    debug.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    debug.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    debug.EntityData.Children = types.NewOrderedMap()
    debug.EntityData.Children.Append("debug-datetime", types.YChild{"DebugDatetime", &debug.DebugDatetime})
    debug.EntityData.Leafs = types.NewOrderedMap()
    debug.EntityData.Leafs.Append("debug-timestamp-disable", types.YLeaf{"DebugTimestampDisable", debug.DebugTimestampDisable})
    debug.EntityData.Leafs.Append("debug-uptime", types.YLeaf{"DebugUptime", debug.DebugUptime})

    debug.EntityData.YListKeys = []string {}

    return &(debug.EntityData)
}

// SyslogService_Timestamps_Debug_DebugDatetime
// Timestamp with date and time
type SyslogService_Timestamps_Debug_DebugDatetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set time format for debug msg.
    DatetimeValue SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue
}

func (debugDatetime *SyslogService_Timestamps_Debug_DebugDatetime) GetEntityData() *types.CommonEntityData {
    debugDatetime.EntityData.YFilter = debugDatetime.YFilter
    debugDatetime.EntityData.YangName = "debug-datetime"
    debugDatetime.EntityData.BundleName = "cisco_ios_xr"
    debugDatetime.EntityData.ParentYangName = "debug"
    debugDatetime.EntityData.SegmentPath = "debug-datetime"
    debugDatetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    debugDatetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    debugDatetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    debugDatetime.EntityData.Children = types.NewOrderedMap()
    debugDatetime.EntityData.Children.Append("datetime-value", types.YChild{"DatetimeValue", &debugDatetime.DatetimeValue})
    debugDatetime.EntityData.Leafs = types.NewOrderedMap()

    debugDatetime.EntityData.YListKeys = []string {}

    return &(debugDatetime.EntityData)
}

// SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue
// Set time format for debug msg
type SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue struct {
    EntityData types.CommonEntityData
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

func (datetimeValue *SyslogService_Timestamps_Debug_DebugDatetime_DatetimeValue) GetEntityData() *types.CommonEntityData {
    datetimeValue.EntityData.YFilter = datetimeValue.YFilter
    datetimeValue.EntityData.YangName = "datetime-value"
    datetimeValue.EntityData.BundleName = "cisco_ios_xr"
    datetimeValue.EntityData.ParentYangName = "debug-datetime"
    datetimeValue.EntityData.SegmentPath = "datetime-value"
    datetimeValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    datetimeValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    datetimeValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    datetimeValue.EntityData.Children = types.NewOrderedMap()
    datetimeValue.EntityData.Leafs = types.NewOrderedMap()
    datetimeValue.EntityData.Leafs.Append("time-stamp-value", types.YLeaf{"TimeStampValue", datetimeValue.TimeStampValue})
    datetimeValue.EntityData.Leafs.Append("msec", types.YLeaf{"Msec", datetimeValue.Msec})
    datetimeValue.EntityData.Leafs.Append("time-zone", types.YLeaf{"TimeZone", datetimeValue.TimeZone})
    datetimeValue.EntityData.Leafs.Append("year", types.YLeaf{"Year", datetimeValue.Year})

    datetimeValue.EntityData.YListKeys = []string {}

    return &(datetimeValue.EntityData)
}

// Syslog
// syslog
type Syslog struct {
    EntityData types.CommonEntityData
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

func (syslog *Syslog) GetEntityData() *types.CommonEntityData {
    syslog.EntityData.YFilter = syslog.YFilter
    syslog.EntityData.YangName = "syslog"
    syslog.EntityData.BundleName = "cisco_ios_xr"
    syslog.EntityData.ParentYangName = "Cisco-IOS-XR-infra-syslog-cfg"
    syslog.EntityData.SegmentPath = "Cisco-IOS-XR-infra-syslog-cfg:syslog"
    syslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslog.EntityData.Children = types.NewOrderedMap()
    syslog.EntityData.Children.Append("monitor-logging", types.YChild{"MonitorLogging", &syslog.MonitorLogging})
    syslog.EntityData.Children.Append("history-logging", types.YChild{"HistoryLogging", &syslog.HistoryLogging})
    syslog.EntityData.Children.Append("logging-facilities", types.YChild{"LoggingFacilities", &syslog.LoggingFacilities})
    syslog.EntityData.Children.Append("trap-logging", types.YChild{"TrapLogging", &syslog.TrapLogging})
    syslog.EntityData.Children.Append("buffered-logging", types.YChild{"BufferedLogging", &syslog.BufferedLogging})
    syslog.EntityData.Children.Append("host-server", types.YChild{"HostServer", &syslog.HostServer})
    syslog.EntityData.Children.Append("console-logging", types.YChild{"ConsoleLogging", &syslog.ConsoleLogging})
    syslog.EntityData.Children.Append("files", types.YChild{"Files", &syslog.Files})
    syslog.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &syslog.Ipv4})
    syslog.EntityData.Children.Append("archive", types.YChild{"Archive", &syslog.Archive})
    syslog.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &syslog.Ipv6})
    syslog.EntityData.Children.Append("source-interface-table", types.YChild{"SourceInterfaceTable", &syslog.SourceInterfaceTable})
    syslog.EntityData.Children.Append("Cisco-IOS-XR-infra-alarm-logger-cfg:alarm-logger", types.YChild{"AlarmLogger", &syslog.AlarmLogger})
    syslog.EntityData.Children.Append("Cisco-IOS-XR-infra-correlator-cfg:correlator", types.YChild{"Correlator", &syslog.Correlator})
    syslog.EntityData.Children.Append("Cisco-IOS-XR-infra-correlator-cfg:suppression", types.YChild{"Suppression", &syslog.Suppression})
    syslog.EntityData.Leafs = types.NewOrderedMap()
    syslog.EntityData.Leafs.Append("host-name-prefix", types.YLeaf{"HostNamePrefix", syslog.HostNamePrefix})
    syslog.EntityData.Leafs.Append("local-log-file-size", types.YLeaf{"LocalLogFileSize", syslog.LocalLogFileSize})
    syslog.EntityData.Leafs.Append("enable-console-logging", types.YLeaf{"EnableConsoleLogging", syslog.EnableConsoleLogging})
    syslog.EntityData.Leafs.Append("suppress-duplicates", types.YLeaf{"SuppressDuplicates", syslog.SuppressDuplicates})

    syslog.EntityData.YListKeys = []string {}

    return &(syslog.EntityData)
}

// Syslog_MonitorLogging
// Set monitor logging
type Syslog_MonitorLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor Logging Level. The type is LoggingLevels. The default value is
    // debug.
    LoggingLevel interface{}

    // Set monitor logging discriminators.
    MonitorDiscriminator Syslog_MonitorLogging_MonitorDiscriminator
}

func (monitorLogging *Syslog_MonitorLogging) GetEntityData() *types.CommonEntityData {
    monitorLogging.EntityData.YFilter = monitorLogging.YFilter
    monitorLogging.EntityData.YangName = "monitor-logging"
    monitorLogging.EntityData.BundleName = "cisco_ios_xr"
    monitorLogging.EntityData.ParentYangName = "syslog"
    monitorLogging.EntityData.SegmentPath = "monitor-logging"
    monitorLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorLogging.EntityData.Children = types.NewOrderedMap()
    monitorLogging.EntityData.Children.Append("monitor-discriminator", types.YChild{"MonitorDiscriminator", &monitorLogging.MonitorDiscriminator})
    monitorLogging.EntityData.Leafs = types.NewOrderedMap()
    monitorLogging.EntityData.Leafs.Append("logging-level", types.YLeaf{"LoggingLevel", monitorLogging.LoggingLevel})

    monitorLogging.EntityData.YListKeys = []string {}

    return &(monitorLogging.EntityData)
}

// Syslog_MonitorLogging_MonitorDiscriminator
// Set monitor logging discriminators
type Syslog_MonitorLogging_MonitorDiscriminator struct {
    EntityData types.CommonEntityData
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

func (monitorDiscriminator *Syslog_MonitorLogging_MonitorDiscriminator) GetEntityData() *types.CommonEntityData {
    monitorDiscriminator.EntityData.YFilter = monitorDiscriminator.YFilter
    monitorDiscriminator.EntityData.YangName = "monitor-discriminator"
    monitorDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    monitorDiscriminator.EntityData.ParentYangName = "monitor-logging"
    monitorDiscriminator.EntityData.SegmentPath = "monitor-discriminator"
    monitorDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorDiscriminator.EntityData.Children = types.NewOrderedMap()
    monitorDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    monitorDiscriminator.EntityData.Leafs.Append("match2", types.YLeaf{"Match2", monitorDiscriminator.Match2})
    monitorDiscriminator.EntityData.Leafs.Append("nomatch1", types.YLeaf{"Nomatch1", monitorDiscriminator.Nomatch1})
    monitorDiscriminator.EntityData.Leafs.Append("match1", types.YLeaf{"Match1", monitorDiscriminator.Match1})
    monitorDiscriminator.EntityData.Leafs.Append("nomatch3", types.YLeaf{"Nomatch3", monitorDiscriminator.Nomatch3})
    monitorDiscriminator.EntityData.Leafs.Append("match3", types.YLeaf{"Match3", monitorDiscriminator.Match3})
    monitorDiscriminator.EntityData.Leafs.Append("nomatch2", types.YLeaf{"Nomatch2", monitorDiscriminator.Nomatch2})

    monitorDiscriminator.EntityData.YListKeys = []string {}

    return &(monitorDiscriminator.EntityData)
}

// Syslog_HistoryLogging
// Set history logging
type Syslog_HistoryLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging history size. The type is interface{} with range: 1..500. The
    // default value is 1.
    HistorySize interface{}

    // History logging level. The type is LoggingLevels. The default value is
    // warning.
    LoggingLevel interface{}
}

func (historyLogging *Syslog_HistoryLogging) GetEntityData() *types.CommonEntityData {
    historyLogging.EntityData.YFilter = historyLogging.YFilter
    historyLogging.EntityData.YangName = "history-logging"
    historyLogging.EntityData.BundleName = "cisco_ios_xr"
    historyLogging.EntityData.ParentYangName = "syslog"
    historyLogging.EntityData.SegmentPath = "history-logging"
    historyLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyLogging.EntityData.Children = types.NewOrderedMap()
    historyLogging.EntityData.Leafs = types.NewOrderedMap()
    historyLogging.EntityData.Leafs.Append("history-size", types.YLeaf{"HistorySize", historyLogging.HistorySize})
    historyLogging.EntityData.Leafs.Append("logging-level", types.YLeaf{"LoggingLevel", historyLogging.LoggingLevel})

    historyLogging.EntityData.YListKeys = []string {}

    return &(historyLogging.EntityData)
}

// Syslog_LoggingFacilities
// Modify message logging facilities
type Syslog_LoggingFacilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Facility from which logging is done. The type is Facility. The default
    // value is local7.
    FacilityLevel interface{}
}

func (loggingFacilities *Syslog_LoggingFacilities) GetEntityData() *types.CommonEntityData {
    loggingFacilities.EntityData.YFilter = loggingFacilities.YFilter
    loggingFacilities.EntityData.YangName = "logging-facilities"
    loggingFacilities.EntityData.BundleName = "cisco_ios_xr"
    loggingFacilities.EntityData.ParentYangName = "syslog"
    loggingFacilities.EntityData.SegmentPath = "logging-facilities"
    loggingFacilities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loggingFacilities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loggingFacilities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loggingFacilities.EntityData.Children = types.NewOrderedMap()
    loggingFacilities.EntityData.Leafs = types.NewOrderedMap()
    loggingFacilities.EntityData.Leafs.Append("facility-level", types.YLeaf{"FacilityLevel", loggingFacilities.FacilityLevel})

    loggingFacilities.EntityData.YListKeys = []string {}

    return &(loggingFacilities.EntityData)
}

// Syslog_TrapLogging
// Set trap logging
type Syslog_TrapLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trap logging level. The type is LoggingLevels. The default value is info.
    LoggingLevel interface{}
}

func (trapLogging *Syslog_TrapLogging) GetEntityData() *types.CommonEntityData {
    trapLogging.EntityData.YFilter = trapLogging.YFilter
    trapLogging.EntityData.YangName = "trap-logging"
    trapLogging.EntityData.BundleName = "cisco_ios_xr"
    trapLogging.EntityData.ParentYangName = "syslog"
    trapLogging.EntityData.SegmentPath = "trap-logging"
    trapLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapLogging.EntityData.Children = types.NewOrderedMap()
    trapLogging.EntityData.Leafs = types.NewOrderedMap()
    trapLogging.EntityData.Leafs.Append("logging-level", types.YLeaf{"LoggingLevel", trapLogging.LoggingLevel})

    trapLogging.EntityData.YListKeys = []string {}

    return &(trapLogging.EntityData)
}

// Syslog_BufferedLogging
// Set buffered logging parameters
type Syslog_BufferedLogging struct {
    EntityData types.CommonEntityData
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

func (bufferedLogging *Syslog_BufferedLogging) GetEntityData() *types.CommonEntityData {
    bufferedLogging.EntityData.YFilter = bufferedLogging.YFilter
    bufferedLogging.EntityData.YangName = "buffered-logging"
    bufferedLogging.EntityData.BundleName = "cisco_ios_xr"
    bufferedLogging.EntityData.ParentYangName = "syslog"
    bufferedLogging.EntityData.SegmentPath = "buffered-logging"
    bufferedLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bufferedLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bufferedLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bufferedLogging.EntityData.Children = types.NewOrderedMap()
    bufferedLogging.EntityData.Children.Append("buffered-discriminator", types.YChild{"BufferedDiscriminator", &bufferedLogging.BufferedDiscriminator})
    bufferedLogging.EntityData.Leafs = types.NewOrderedMap()
    bufferedLogging.EntityData.Leafs.Append("logging-level", types.YLeaf{"LoggingLevel", bufferedLogging.LoggingLevel})
    bufferedLogging.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", bufferedLogging.BufferSize})

    bufferedLogging.EntityData.YListKeys = []string {}

    return &(bufferedLogging.EntityData)
}

// Syslog_BufferedLogging_BufferedDiscriminator
// Set buffered logging discriminators
type Syslog_BufferedLogging_BufferedDiscriminator struct {
    EntityData types.CommonEntityData
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

func (bufferedDiscriminator *Syslog_BufferedLogging_BufferedDiscriminator) GetEntityData() *types.CommonEntityData {
    bufferedDiscriminator.EntityData.YFilter = bufferedDiscriminator.YFilter
    bufferedDiscriminator.EntityData.YangName = "buffered-discriminator"
    bufferedDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    bufferedDiscriminator.EntityData.ParentYangName = "buffered-logging"
    bufferedDiscriminator.EntityData.SegmentPath = "buffered-discriminator"
    bufferedDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bufferedDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bufferedDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bufferedDiscriminator.EntityData.Children = types.NewOrderedMap()
    bufferedDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    bufferedDiscriminator.EntityData.Leafs.Append("match2", types.YLeaf{"Match2", bufferedDiscriminator.Match2})
    bufferedDiscriminator.EntityData.Leafs.Append("nomatch1", types.YLeaf{"Nomatch1", bufferedDiscriminator.Nomatch1})
    bufferedDiscriminator.EntityData.Leafs.Append("match1", types.YLeaf{"Match1", bufferedDiscriminator.Match1})
    bufferedDiscriminator.EntityData.Leafs.Append("nomatch3", types.YLeaf{"Nomatch3", bufferedDiscriminator.Nomatch3})
    bufferedDiscriminator.EntityData.Leafs.Append("match3", types.YLeaf{"Match3", bufferedDiscriminator.Match3})
    bufferedDiscriminator.EntityData.Leafs.Append("nomatch2", types.YLeaf{"Nomatch2", bufferedDiscriminator.Nomatch2})

    bufferedDiscriminator.EntityData.YListKeys = []string {}

    return &(bufferedDiscriminator.EntityData)
}

// Syslog_HostServer
// Configure logging host
type Syslog_HostServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs Syslog_HostServer_Vrfs
}

func (hostServer *Syslog_HostServer) GetEntityData() *types.CommonEntityData {
    hostServer.EntityData.YFilter = hostServer.YFilter
    hostServer.EntityData.YangName = "host-server"
    hostServer.EntityData.BundleName = "cisco_ios_xr"
    hostServer.EntityData.ParentYangName = "syslog"
    hostServer.EntityData.SegmentPath = "host-server"
    hostServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostServer.EntityData.Children = types.NewOrderedMap()
    hostServer.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &hostServer.Vrfs})
    hostServer.EntityData.Leafs = types.NewOrderedMap()

    hostServer.EntityData.YListKeys = []string {}

    return &(hostServer.EntityData)
}

// Syslog_HostServer_Vrfs
// VRF table
type Syslog_HostServer_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of Syslog_HostServer_Vrfs_Vrf.
    Vrf []*Syslog_HostServer_Vrfs_Vrf
}

func (vrfs *Syslog_HostServer_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "host-server"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf
// VRF specific data
type Syslog_HostServer_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // List of the IPv6 logging host.
    Ipv6s Syslog_HostServer_Vrfs_Vrf_Ipv6s

    // List of the logging host.
    Hosts Syslog_HostServer_Vrfs_Vrf_Hosts

    // List of the IPv4 logging host.
    Ipv4s Syslog_HostServer_Vrfs_Vrf_Ipv4s
}

func (vrf *Syslog_HostServer_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("ipv6s", types.YChild{"Ipv6s", &vrf.Ipv6s})
    vrf.EntityData.Children.Append("hosts", types.YChild{"Hosts", &vrf.Hosts})
    vrf.EntityData.Children.Append("ipv4s", types.YChild{"Ipv4s", &vrf.Ipv4s})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv6s
// List of the IPv6 logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6s struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address of the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6.
    Ipv6 []*Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6
}

func (ipv6s *Syslog_HostServer_Vrfs_Vrf_Ipv6s) GetEntityData() *types.CommonEntityData {
    ipv6s.EntityData.YFilter = ipv6s.YFilter
    ipv6s.EntityData.YangName = "ipv6s"
    ipv6s.EntityData.BundleName = "cisco_ios_xr"
    ipv6s.EntityData.ParentYangName = "vrf"
    ipv6s.EntityData.SegmentPath = "ipv6s"
    ipv6s.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6s.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6s.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6s.EntityData.Children = types.NewOrderedMap()
    ipv6s.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", nil})
    for i := range ipv6s.Ipv6 {
        ipv6s.EntityData.Children.Append(types.GetSegmentPath(ipv6s.Ipv6[i]), types.YChild{"Ipv6", ipv6s.Ipv6[i]})
    }
    ipv6s.EntityData.Leafs = types.NewOrderedMap()

    ipv6s.EntityData.YListKeys = []string {}

    return &(ipv6s.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6
// IPv6 address of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address of the logging host. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Severity/Port for the logging host.
    Ipv6SeverityPort Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityPort

    // Severity container of the logging host.
    Ipv6SeverityLevels Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels
}

func (ipv6 *Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "ipv6s"
    ipv6.EntityData.SegmentPath = "ipv6" + types.AddKeyToken(ipv6.Address, "address")
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("ipv6-severity-port", types.YChild{"Ipv6SeverityPort", &ipv6.Ipv6SeverityPort})
    ipv6.EntityData.Children.Append("ipv6-severity-levels", types.YChild{"Ipv6SeverityLevels", &ipv6.Ipv6SeverityLevels})
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv6.Address})

    ipv6.EntityData.YListKeys = []string {"Address"}

    return &(ipv6.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityPort
// Severity/Port for the logging host
// This type is a presence type.
type Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Severity for the logging host. The type is interface{} with range:
    // 0..4294967295. The default value is 6.
    Severity interface{}

    // Port for the logging host. The type is interface{} with range:
    // 0..4294967295. The default value is 514.
    Port interface{}
}

func (ipv6SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityPort) GetEntityData() *types.CommonEntityData {
    ipv6SeverityPort.EntityData.YFilter = ipv6SeverityPort.YFilter
    ipv6SeverityPort.EntityData.YangName = "ipv6-severity-port"
    ipv6SeverityPort.EntityData.BundleName = "cisco_ios_xr"
    ipv6SeverityPort.EntityData.ParentYangName = "ipv6"
    ipv6SeverityPort.EntityData.SegmentPath = "ipv6-severity-port"
    ipv6SeverityPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SeverityPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SeverityPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SeverityPort.EntityData.Children = types.NewOrderedMap()
    ipv6SeverityPort.EntityData.Leafs = types.NewOrderedMap()
    ipv6SeverityPort.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", ipv6SeverityPort.Severity})
    ipv6SeverityPort.EntityData.Leafs.Append("port", types.YLeaf{"Port", ipv6SeverityPort.Port})

    ipv6SeverityPort.EntityData.YListKeys = []string {}

    return &(ipv6SeverityPort.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels
// Severity container of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel.
    Ipv6SeverityLevel []*Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel
}

func (ipv6SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels) GetEntityData() *types.CommonEntityData {
    ipv6SeverityLevels.EntityData.YFilter = ipv6SeverityLevels.YFilter
    ipv6SeverityLevels.EntityData.YangName = "ipv6-severity-levels"
    ipv6SeverityLevels.EntityData.BundleName = "cisco_ios_xr"
    ipv6SeverityLevels.EntityData.ParentYangName = "ipv6"
    ipv6SeverityLevels.EntityData.SegmentPath = "ipv6-severity-levels"
    ipv6SeverityLevels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SeverityLevels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SeverityLevels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SeverityLevels.EntityData.Children = types.NewOrderedMap()
    ipv6SeverityLevels.EntityData.Children.Append("ipv6-severity-level", types.YChild{"Ipv6SeverityLevel", nil})
    for i := range ipv6SeverityLevels.Ipv6SeverityLevel {
        ipv6SeverityLevels.EntityData.Children.Append(types.GetSegmentPath(ipv6SeverityLevels.Ipv6SeverityLevel[i]), types.YChild{"Ipv6SeverityLevel", ipv6SeverityLevels.Ipv6SeverityLevel[i]})
    }
    ipv6SeverityLevels.EntityData.Leafs = types.NewOrderedMap()

    ipv6SeverityLevels.EntityData.YListKeys = []string {}

    return &(ipv6SeverityLevels.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel
// Severity for the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Severity for the logging host. The type is
    // LogSeverity.
    Severity interface{}
}

func (ipv6SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv6s_Ipv6_Ipv6SeverityLevels_Ipv6SeverityLevel) GetEntityData() *types.CommonEntityData {
    ipv6SeverityLevel.EntityData.YFilter = ipv6SeverityLevel.YFilter
    ipv6SeverityLevel.EntityData.YangName = "ipv6-severity-level"
    ipv6SeverityLevel.EntityData.BundleName = "cisco_ios_xr"
    ipv6SeverityLevel.EntityData.ParentYangName = "ipv6-severity-levels"
    ipv6SeverityLevel.EntityData.SegmentPath = "ipv6-severity-level" + types.AddKeyToken(ipv6SeverityLevel.Severity, "severity")
    ipv6SeverityLevel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SeverityLevel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SeverityLevel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SeverityLevel.EntityData.Children = types.NewOrderedMap()
    ipv6SeverityLevel.EntityData.Leafs = types.NewOrderedMap()
    ipv6SeverityLevel.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", ipv6SeverityLevel.Severity})

    ipv6SeverityLevel.EntityData.YListKeys = []string {"Severity"}

    return &(ipv6SeverityLevel.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Hosts
// List of the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Hosts_Host.
    Host []*Syslog_HostServer_Vrfs_Vrf_Hosts_Host
}

func (hosts *Syslog_HostServer_Vrfs_Vrf_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "vrf"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host
// Name of the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the logging host. The type is string.
    HostName interface{}

    // Severity container of the logging host.
    HostNameSeverities Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities

    // Severity/Port for the logging host.
    HostSeverityPort Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort
}

func (host *Syslog_HostServer_Vrfs_Vrf_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.HostName, "host-name")
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Children.Append("host-name-severities", types.YChild{"HostNameSeverities", &host.HostNameSeverities})
    host.EntityData.Children.Append("host-severity-port", types.YChild{"HostSeverityPort", &host.HostSeverityPort})
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", host.HostName})

    host.EntityData.YListKeys = []string {"HostName"}

    return &(host.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities
// Severity container of the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity.
    HostNameSeverity []*Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity
}

func (hostNameSeverities *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities) GetEntityData() *types.CommonEntityData {
    hostNameSeverities.EntityData.YFilter = hostNameSeverities.YFilter
    hostNameSeverities.EntityData.YangName = "host-name-severities"
    hostNameSeverities.EntityData.BundleName = "cisco_ios_xr"
    hostNameSeverities.EntityData.ParentYangName = "host"
    hostNameSeverities.EntityData.SegmentPath = "host-name-severities"
    hostNameSeverities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostNameSeverities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostNameSeverities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostNameSeverities.EntityData.Children = types.NewOrderedMap()
    hostNameSeverities.EntityData.Children.Append("host-name-severity", types.YChild{"HostNameSeverity", nil})
    for i := range hostNameSeverities.HostNameSeverity {
        hostNameSeverities.EntityData.Children.Append(types.GetSegmentPath(hostNameSeverities.HostNameSeverity[i]), types.YChild{"HostNameSeverity", hostNameSeverities.HostNameSeverity[i]})
    }
    hostNameSeverities.EntityData.Leafs = types.NewOrderedMap()

    hostNameSeverities.EntityData.YListKeys = []string {}

    return &(hostNameSeverities.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity
// Severity for the logging host
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Severity for the logging host. The type is
    // LogSeverity.
    Severity interface{}
}

func (hostNameSeverity *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostNameSeverities_HostNameSeverity) GetEntityData() *types.CommonEntityData {
    hostNameSeverity.EntityData.YFilter = hostNameSeverity.YFilter
    hostNameSeverity.EntityData.YangName = "host-name-severity"
    hostNameSeverity.EntityData.BundleName = "cisco_ios_xr"
    hostNameSeverity.EntityData.ParentYangName = "host-name-severities"
    hostNameSeverity.EntityData.SegmentPath = "host-name-severity" + types.AddKeyToken(hostNameSeverity.Severity, "severity")
    hostNameSeverity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostNameSeverity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostNameSeverity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostNameSeverity.EntityData.Children = types.NewOrderedMap()
    hostNameSeverity.EntityData.Leafs = types.NewOrderedMap()
    hostNameSeverity.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", hostNameSeverity.Severity})

    hostNameSeverity.EntityData.YListKeys = []string {"Severity"}

    return &(hostNameSeverity.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort
// Severity/Port for the logging host
// This type is a presence type.
type Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Severity for the logging host. The type is interface{} with range:
    // 0..4294967295. The default value is 6.
    Severity interface{}

    // Port for the logging host. The type is interface{} with range:
    // 0..4294967295. The default value is 514.
    Port interface{}
}

func (hostSeverityPort *Syslog_HostServer_Vrfs_Vrf_Hosts_Host_HostSeverityPort) GetEntityData() *types.CommonEntityData {
    hostSeverityPort.EntityData.YFilter = hostSeverityPort.YFilter
    hostSeverityPort.EntityData.YangName = "host-severity-port"
    hostSeverityPort.EntityData.BundleName = "cisco_ios_xr"
    hostSeverityPort.EntityData.ParentYangName = "host"
    hostSeverityPort.EntityData.SegmentPath = "host-severity-port"
    hostSeverityPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostSeverityPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostSeverityPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostSeverityPort.EntityData.Children = types.NewOrderedMap()
    hostSeverityPort.EntityData.Leafs = types.NewOrderedMap()
    hostSeverityPort.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", hostSeverityPort.Severity})
    hostSeverityPort.EntityData.Leafs.Append("port", types.YLeaf{"Port", hostSeverityPort.Port})

    hostSeverityPort.EntityData.YListKeys = []string {}

    return &(hostSeverityPort.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv4s
// List of the IPv4 logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4s struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address of the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4.
    Ipv4 []*Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4
}

func (ipv4s *Syslog_HostServer_Vrfs_Vrf_Ipv4s) GetEntityData() *types.CommonEntityData {
    ipv4s.EntityData.YFilter = ipv4s.YFilter
    ipv4s.EntityData.YangName = "ipv4s"
    ipv4s.EntityData.BundleName = "cisco_ios_xr"
    ipv4s.EntityData.ParentYangName = "vrf"
    ipv4s.EntityData.SegmentPath = "ipv4s"
    ipv4s.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4s.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4s.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4s.EntityData.Children = types.NewOrderedMap()
    ipv4s.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", nil})
    for i := range ipv4s.Ipv4 {
        ipv4s.EntityData.Children.Append(types.GetSegmentPath(ipv4s.Ipv4[i]), types.YChild{"Ipv4", ipv4s.Ipv4[i]})
    }
    ipv4s.EntityData.Leafs = types.NewOrderedMap()

    ipv4s.EntityData.YListKeys = []string {}

    return &(ipv4s.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4
// IPv4 address of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address of the logging host. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Severity container of the logging host.
    Ipv4SeverityLevels Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels

    // Severity/Port for the logging host.
    Ipv4SeverityPort Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityPort
}

func (ipv4 *Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "ipv4s"
    ipv4.EntityData.SegmentPath = "ipv4" + types.AddKeyToken(ipv4.Address, "address")
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("ipv4-severity-levels", types.YChild{"Ipv4SeverityLevels", &ipv4.Ipv4SeverityLevels})
    ipv4.EntityData.Children.Append("ipv4-severity-port", types.YChild{"Ipv4SeverityPort", &ipv4.Ipv4SeverityPort})
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4.Address})

    ipv4.EntityData.YListKeys = []string {"Address"}

    return &(ipv4.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels
// Severity container of the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Severity for the logging host. The type is slice of
    // Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel.
    Ipv4SeverityLevel []*Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel
}

func (ipv4SeverityLevels *Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels) GetEntityData() *types.CommonEntityData {
    ipv4SeverityLevels.EntityData.YFilter = ipv4SeverityLevels.YFilter
    ipv4SeverityLevels.EntityData.YangName = "ipv4-severity-levels"
    ipv4SeverityLevels.EntityData.BundleName = "cisco_ios_xr"
    ipv4SeverityLevels.EntityData.ParentYangName = "ipv4"
    ipv4SeverityLevels.EntityData.SegmentPath = "ipv4-severity-levels"
    ipv4SeverityLevels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SeverityLevels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SeverityLevels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SeverityLevels.EntityData.Children = types.NewOrderedMap()
    ipv4SeverityLevels.EntityData.Children.Append("ipv4-severity-level", types.YChild{"Ipv4SeverityLevel", nil})
    for i := range ipv4SeverityLevels.Ipv4SeverityLevel {
        ipv4SeverityLevels.EntityData.Children.Append(types.GetSegmentPath(ipv4SeverityLevels.Ipv4SeverityLevel[i]), types.YChild{"Ipv4SeverityLevel", ipv4SeverityLevels.Ipv4SeverityLevel[i]})
    }
    ipv4SeverityLevels.EntityData.Leafs = types.NewOrderedMap()

    ipv4SeverityLevels.EntityData.YListKeys = []string {}

    return &(ipv4SeverityLevels.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel
// Severity for the logging host
type Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Severity for the logging host. The type is
    // LogSeverity.
    Severity interface{}
}

func (ipv4SeverityLevel *Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityLevels_Ipv4SeverityLevel) GetEntityData() *types.CommonEntityData {
    ipv4SeverityLevel.EntityData.YFilter = ipv4SeverityLevel.YFilter
    ipv4SeverityLevel.EntityData.YangName = "ipv4-severity-level"
    ipv4SeverityLevel.EntityData.BundleName = "cisco_ios_xr"
    ipv4SeverityLevel.EntityData.ParentYangName = "ipv4-severity-levels"
    ipv4SeverityLevel.EntityData.SegmentPath = "ipv4-severity-level" + types.AddKeyToken(ipv4SeverityLevel.Severity, "severity")
    ipv4SeverityLevel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SeverityLevel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SeverityLevel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SeverityLevel.EntityData.Children = types.NewOrderedMap()
    ipv4SeverityLevel.EntityData.Leafs = types.NewOrderedMap()
    ipv4SeverityLevel.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", ipv4SeverityLevel.Severity})

    ipv4SeverityLevel.EntityData.YListKeys = []string {"Severity"}

    return &(ipv4SeverityLevel.EntityData)
}

// Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityPort
// Severity/Port for the logging host
// This type is a presence type.
type Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Severity for the logging host. The type is interface{} with range:
    // 0..4294967295. The default value is 6.
    Severity interface{}

    // Port for the logging host. The type is interface{} with range:
    // 0..4294967295. The default value is 514.
    Port interface{}
}

func (ipv4SeverityPort *Syslog_HostServer_Vrfs_Vrf_Ipv4s_Ipv4_Ipv4SeverityPort) GetEntityData() *types.CommonEntityData {
    ipv4SeverityPort.EntityData.YFilter = ipv4SeverityPort.YFilter
    ipv4SeverityPort.EntityData.YangName = "ipv4-severity-port"
    ipv4SeverityPort.EntityData.BundleName = "cisco_ios_xr"
    ipv4SeverityPort.EntityData.ParentYangName = "ipv4"
    ipv4SeverityPort.EntityData.SegmentPath = "ipv4-severity-port"
    ipv4SeverityPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SeverityPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SeverityPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SeverityPort.EntityData.Children = types.NewOrderedMap()
    ipv4SeverityPort.EntityData.Leafs = types.NewOrderedMap()
    ipv4SeverityPort.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", ipv4SeverityPort.Severity})
    ipv4SeverityPort.EntityData.Leafs.Append("port", types.YLeaf{"Port", ipv4SeverityPort.Port})

    ipv4SeverityPort.EntityData.YListKeys = []string {}

    return &(ipv4SeverityPort.EntityData)
}

// Syslog_ConsoleLogging
// Set console logging
type Syslog_ConsoleLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Console logging level. The type is LoggingLevels. The default value is
    // warning.
    LoggingLevel interface{}

    // Set console logging discriminators.
    ConsoleDiscriminator Syslog_ConsoleLogging_ConsoleDiscriminator
}

func (consoleLogging *Syslog_ConsoleLogging) GetEntityData() *types.CommonEntityData {
    consoleLogging.EntityData.YFilter = consoleLogging.YFilter
    consoleLogging.EntityData.YangName = "console-logging"
    consoleLogging.EntityData.BundleName = "cisco_ios_xr"
    consoleLogging.EntityData.ParentYangName = "syslog"
    consoleLogging.EntityData.SegmentPath = "console-logging"
    consoleLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    consoleLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    consoleLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    consoleLogging.EntityData.Children = types.NewOrderedMap()
    consoleLogging.EntityData.Children.Append("console-discriminator", types.YChild{"ConsoleDiscriminator", &consoleLogging.ConsoleDiscriminator})
    consoleLogging.EntityData.Leafs = types.NewOrderedMap()
    consoleLogging.EntityData.Leafs.Append("logging-level", types.YLeaf{"LoggingLevel", consoleLogging.LoggingLevel})

    consoleLogging.EntityData.YListKeys = []string {}

    return &(consoleLogging.EntityData)
}

// Syslog_ConsoleLogging_ConsoleDiscriminator
// Set console logging discriminators
type Syslog_ConsoleLogging_ConsoleDiscriminator struct {
    EntityData types.CommonEntityData
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

func (consoleDiscriminator *Syslog_ConsoleLogging_ConsoleDiscriminator) GetEntityData() *types.CommonEntityData {
    consoleDiscriminator.EntityData.YFilter = consoleDiscriminator.YFilter
    consoleDiscriminator.EntityData.YangName = "console-discriminator"
    consoleDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    consoleDiscriminator.EntityData.ParentYangName = "console-logging"
    consoleDiscriminator.EntityData.SegmentPath = "console-discriminator"
    consoleDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    consoleDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    consoleDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    consoleDiscriminator.EntityData.Children = types.NewOrderedMap()
    consoleDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    consoleDiscriminator.EntityData.Leafs.Append("match2", types.YLeaf{"Match2", consoleDiscriminator.Match2})
    consoleDiscriminator.EntityData.Leafs.Append("nomatch1", types.YLeaf{"Nomatch1", consoleDiscriminator.Nomatch1})
    consoleDiscriminator.EntityData.Leafs.Append("match1", types.YLeaf{"Match1", consoleDiscriminator.Match1})
    consoleDiscriminator.EntityData.Leafs.Append("nomatch3", types.YLeaf{"Nomatch3", consoleDiscriminator.Nomatch3})
    consoleDiscriminator.EntityData.Leafs.Append("match3", types.YLeaf{"Match3", consoleDiscriminator.Match3})
    consoleDiscriminator.EntityData.Leafs.Append("nomatch2", types.YLeaf{"Nomatch2", consoleDiscriminator.Nomatch2})

    consoleDiscriminator.EntityData.YListKeys = []string {}

    return &(consoleDiscriminator.EntityData)
}

// Syslog_Files
// Configure logging file destination
type Syslog_Files struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify File Name. The type is slice of Syslog_Files_File.
    File []*Syslog_Files_File
}

func (files *Syslog_Files) GetEntityData() *types.CommonEntityData {
    files.EntityData.YFilter = files.YFilter
    files.EntityData.YangName = "files"
    files.EntityData.BundleName = "cisco_ios_xr"
    files.EntityData.ParentYangName = "syslog"
    files.EntityData.SegmentPath = "files"
    files.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    files.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    files.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    files.EntityData.Children = types.NewOrderedMap()
    files.EntityData.Children.Append("file", types.YChild{"File", nil})
    for i := range files.File {
        files.EntityData.Children.Append(types.GetSegmentPath(files.File[i]), types.YChild{"File", files.File[i]})
    }
    files.EntityData.Leafs = types.NewOrderedMap()

    files.EntityData.YListKeys = []string {}

    return &(files.EntityData)
}

// Syslog_Files_File
// Specify File Name
type Syslog_Files_File struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the file. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    FileName interface{}

    // Specifications of the logging file destination.
    FileSpecification Syslog_Files_File_FileSpecification

    // Set File logging discriminators.
    FileLogDiscriminator Syslog_Files_File_FileLogDiscriminator
}

func (file *Syslog_Files_File) GetEntityData() *types.CommonEntityData {
    file.EntityData.YFilter = file.YFilter
    file.EntityData.YangName = "file"
    file.EntityData.BundleName = "cisco_ios_xr"
    file.EntityData.ParentYangName = "files"
    file.EntityData.SegmentPath = "file" + types.AddKeyToken(file.FileName, "file-name")
    file.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    file.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    file.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    file.EntityData.Children = types.NewOrderedMap()
    file.EntityData.Children.Append("file-specification", types.YChild{"FileSpecification", &file.FileSpecification})
    file.EntityData.Children.Append("file-log-discriminator", types.YChild{"FileLogDiscriminator", &file.FileLogDiscriminator})
    file.EntityData.Leafs = types.NewOrderedMap()
    file.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", file.FileName})

    file.EntityData.YListKeys = []string {"FileName"}

    return &(file.EntityData)
}

// Syslog_Files_File_FileSpecification
// Specifications of the logging file destination
type Syslog_Files_File_FileSpecification struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // File path. The type is string.
    Path interface{}

    // Maximum file size (in KB). The type is interface{} with range:
    // 0..4294967295. The default value is 1024.
    MaxFileSize interface{}

    // Severity of messages. The type is interface{} with range: 0..4294967295.
    // The default value is 6.
    Severity interface{}
}

func (fileSpecification *Syslog_Files_File_FileSpecification) GetEntityData() *types.CommonEntityData {
    fileSpecification.EntityData.YFilter = fileSpecification.YFilter
    fileSpecification.EntityData.YangName = "file-specification"
    fileSpecification.EntityData.BundleName = "cisco_ios_xr"
    fileSpecification.EntityData.ParentYangName = "file"
    fileSpecification.EntityData.SegmentPath = "file-specification"
    fileSpecification.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileSpecification.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileSpecification.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileSpecification.EntityData.Children = types.NewOrderedMap()
    fileSpecification.EntityData.Leafs = types.NewOrderedMap()
    fileSpecification.EntityData.Leafs.Append("path", types.YLeaf{"Path", fileSpecification.Path})
    fileSpecification.EntityData.Leafs.Append("max-file-size", types.YLeaf{"MaxFileSize", fileSpecification.MaxFileSize})
    fileSpecification.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", fileSpecification.Severity})

    fileSpecification.EntityData.YListKeys = []string {}

    return &(fileSpecification.EntityData)
}

// Syslog_Files_File_FileLogDiscriminator
// Set File logging discriminators
type Syslog_Files_File_FileLogDiscriminator struct {
    EntityData types.CommonEntityData
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

func (fileLogDiscriminator *Syslog_Files_File_FileLogDiscriminator) GetEntityData() *types.CommonEntityData {
    fileLogDiscriminator.EntityData.YFilter = fileLogDiscriminator.YFilter
    fileLogDiscriminator.EntityData.YangName = "file-log-discriminator"
    fileLogDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    fileLogDiscriminator.EntityData.ParentYangName = "file"
    fileLogDiscriminator.EntityData.SegmentPath = "file-log-discriminator"
    fileLogDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileLogDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileLogDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileLogDiscriminator.EntityData.Children = types.NewOrderedMap()
    fileLogDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    fileLogDiscriminator.EntityData.Leafs.Append("nomatch2", types.YLeaf{"Nomatch2", fileLogDiscriminator.Nomatch2})
    fileLogDiscriminator.EntityData.Leafs.Append("match3", types.YLeaf{"Match3", fileLogDiscriminator.Match3})
    fileLogDiscriminator.EntityData.Leafs.Append("nomatch3", types.YLeaf{"Nomatch3", fileLogDiscriminator.Nomatch3})
    fileLogDiscriminator.EntityData.Leafs.Append("match1", types.YLeaf{"Match1", fileLogDiscriminator.Match1})
    fileLogDiscriminator.EntityData.Leafs.Append("nomatch1", types.YLeaf{"Nomatch1", fileLogDiscriminator.Nomatch1})
    fileLogDiscriminator.EntityData.Leafs.Append("match2", types.YLeaf{"Match2", fileLogDiscriminator.Match2})

    fileLogDiscriminator.EntityData.YListKeys = []string {}

    return &(fileLogDiscriminator.EntityData)
}

// Syslog_Ipv4
// Syslog TOS bit for outgoing messages
type Syslog_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DSCP value.
    Dscp Syslog_Ipv4_Dscp

    // Type of service.
    Tos Syslog_Ipv4_Tos

    // Precedence value.
    Precedence Syslog_Ipv4_Precedence
}

func (ipv4 *Syslog_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "syslog"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("dscp", types.YChild{"Dscp", &ipv4.Dscp})
    ipv4.EntityData.Children.Append("tos", types.YChild{"Tos", &ipv4.Tos})
    ipv4.EntityData.Children.Append("precedence", types.YChild{"Precedence", &ipv4.Precedence})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Syslog_Ipv4_Dscp
// DSCP value
// This type is a presence type.
type Syslog_Ipv4_Dscp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (dscp *Syslog_Ipv4_Dscp) GetEntityData() *types.CommonEntityData {
    dscp.EntityData.YFilter = dscp.YFilter
    dscp.EntityData.YangName = "dscp"
    dscp.EntityData.BundleName = "cisco_ios_xr"
    dscp.EntityData.ParentYangName = "ipv4"
    dscp.EntityData.SegmentPath = "dscp"
    dscp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dscp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dscp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dscp.EntityData.Children = types.NewOrderedMap()
    dscp.EntityData.Leafs = types.NewOrderedMap()
    dscp.EntityData.Leafs.Append("type", types.YLeaf{"Type", dscp.Type})
    dscp.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", dscp.Unused})
    dscp.EntityData.Leafs.Append("value", types.YLeaf{"Value", dscp.Value})

    dscp.EntityData.YListKeys = []string {}

    return &(dscp.EntityData)
}

// Syslog_Ipv4_Tos
// Type of service
type Syslog_Ipv4_Tos struct {
    EntityData types.CommonEntityData
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

func (tos *Syslog_Ipv4_Tos) GetEntityData() *types.CommonEntityData {
    tos.EntityData.YFilter = tos.YFilter
    tos.EntityData.YangName = "tos"
    tos.EntityData.BundleName = "cisco_ios_xr"
    tos.EntityData.ParentYangName = "ipv4"
    tos.EntityData.SegmentPath = "tos"
    tos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tos.EntityData.Children = types.NewOrderedMap()
    tos.EntityData.Leafs = types.NewOrderedMap()
    tos.EntityData.Leafs.Append("type", types.YLeaf{"Type", tos.Type})
    tos.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", tos.Precedence})
    tos.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", tos.Dscp})

    tos.EntityData.YListKeys = []string {}

    return &(tos.EntityData)
}

// Syslog_Ipv4_Precedence
// Precedence value
// This type is a presence type.
type Syslog_Ipv4_Precedence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (precedence *Syslog_Ipv4_Precedence) GetEntityData() *types.CommonEntityData {
    precedence.EntityData.YFilter = precedence.YFilter
    precedence.EntityData.YangName = "precedence"
    precedence.EntityData.BundleName = "cisco_ios_xr"
    precedence.EntityData.ParentYangName = "ipv4"
    precedence.EntityData.SegmentPath = "precedence"
    precedence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    precedence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    precedence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    precedence.EntityData.Children = types.NewOrderedMap()
    precedence.EntityData.Leafs = types.NewOrderedMap()
    precedence.EntityData.Leafs.Append("type", types.YLeaf{"Type", precedence.Type})
    precedence.EntityData.Leafs.Append("value", types.YLeaf{"Value", precedence.Value})
    precedence.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", precedence.Unused})

    precedence.EntityData.YListKeys = []string {}

    return &(precedence.EntityData)
}

// Syslog_Archive
// Archive attributes configuration
type Syslog_Archive struct {
    EntityData types.CommonEntityData
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

func (archive *Syslog_Archive) GetEntityData() *types.CommonEntityData {
    archive.EntityData.YFilter = archive.YFilter
    archive.EntityData.YangName = "archive"
    archive.EntityData.BundleName = "cisco_ios_xr"
    archive.EntityData.ParentYangName = "syslog"
    archive.EntityData.SegmentPath = "archive"
    archive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    archive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    archive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    archive.EntityData.Children = types.NewOrderedMap()
    archive.EntityData.Leafs = types.NewOrderedMap()
    archive.EntityData.Leafs.Append("size", types.YLeaf{"Size", archive.Size})
    archive.EntityData.Leafs.Append("file-size", types.YLeaf{"FileSize", archive.FileSize})
    archive.EntityData.Leafs.Append("device", types.YLeaf{"Device", archive.Device})
    archive.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", archive.Threshold})
    archive.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", archive.Frequency})
    archive.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", archive.Severity})
    archive.EntityData.Leafs.Append("length", types.YLeaf{"Length", archive.Length})

    archive.EntityData.YListKeys = []string {}

    return &(archive.EntityData)
}

// Syslog_Ipv6
// Syslog traffic class bit for outgoing messages
type Syslog_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DSCP value.
    Dscp Syslog_Ipv6_Dscp

    // Type of traffic class.
    TrafficClass Syslog_Ipv6_TrafficClass

    // Precedence value.
    Precedence Syslog_Ipv6_Precedence
}

func (ipv6 *Syslog_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "syslog"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("dscp", types.YChild{"Dscp", &ipv6.Dscp})
    ipv6.EntityData.Children.Append("traffic-class", types.YChild{"TrafficClass", &ipv6.TrafficClass})
    ipv6.EntityData.Children.Append("precedence", types.YChild{"Precedence", &ipv6.Precedence})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Syslog_Ipv6_Dscp
// DSCP value
// This type is a presence type.
type Syslog_Ipv6_Dscp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (dscp *Syslog_Ipv6_Dscp) GetEntityData() *types.CommonEntityData {
    dscp.EntityData.YFilter = dscp.YFilter
    dscp.EntityData.YangName = "dscp"
    dscp.EntityData.BundleName = "cisco_ios_xr"
    dscp.EntityData.ParentYangName = "ipv6"
    dscp.EntityData.SegmentPath = "dscp"
    dscp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dscp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dscp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dscp.EntityData.Children = types.NewOrderedMap()
    dscp.EntityData.Leafs = types.NewOrderedMap()
    dscp.EntityData.Leafs.Append("type", types.YLeaf{"Type", dscp.Type})
    dscp.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", dscp.Unused})
    dscp.EntityData.Leafs.Append("value", types.YLeaf{"Value", dscp.Value})

    dscp.EntityData.YListKeys = []string {}

    return &(dscp.EntityData)
}

// Syslog_Ipv6_TrafficClass
// Type of traffic class
type Syslog_Ipv6_TrafficClass struct {
    EntityData types.CommonEntityData
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

func (trafficClass *Syslog_Ipv6_TrafficClass) GetEntityData() *types.CommonEntityData {
    trafficClass.EntityData.YFilter = trafficClass.YFilter
    trafficClass.EntityData.YangName = "traffic-class"
    trafficClass.EntityData.BundleName = "cisco_ios_xr"
    trafficClass.EntityData.ParentYangName = "ipv6"
    trafficClass.EntityData.SegmentPath = "traffic-class"
    trafficClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficClass.EntityData.Children = types.NewOrderedMap()
    trafficClass.EntityData.Leafs = types.NewOrderedMap()
    trafficClass.EntityData.Leafs.Append("type", types.YLeaf{"Type", trafficClass.Type})
    trafficClass.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", trafficClass.Precedence})
    trafficClass.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", trafficClass.Dscp})

    trafficClass.EntityData.YListKeys = []string {}

    return &(trafficClass.EntityData)
}

// Syslog_Ipv6_Precedence
// Precedence value
// This type is a presence type.
type Syslog_Ipv6_Precedence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (precedence *Syslog_Ipv6_Precedence) GetEntityData() *types.CommonEntityData {
    precedence.EntityData.YFilter = precedence.YFilter
    precedence.EntityData.YangName = "precedence"
    precedence.EntityData.BundleName = "cisco_ios_xr"
    precedence.EntityData.ParentYangName = "ipv6"
    precedence.EntityData.SegmentPath = "precedence"
    precedence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    precedence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    precedence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    precedence.EntityData.Children = types.NewOrderedMap()
    precedence.EntityData.Leafs = types.NewOrderedMap()
    precedence.EntityData.Leafs.Append("type", types.YLeaf{"Type", precedence.Type})
    precedence.EntityData.Leafs.Append("value", types.YLeaf{"Value", precedence.Value})
    precedence.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", precedence.Unused})

    precedence.EntityData.YListKeys = []string {}

    return &(precedence.EntityData)
}

// Syslog_SourceInterfaceTable
// Configure source interface
type Syslog_SourceInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify interface for source address in logging transactions.
    SourceInterfaceValues Syslog_SourceInterfaceTable_SourceInterfaceValues
}

func (sourceInterfaceTable *Syslog_SourceInterfaceTable) GetEntityData() *types.CommonEntityData {
    sourceInterfaceTable.EntityData.YFilter = sourceInterfaceTable.YFilter
    sourceInterfaceTable.EntityData.YangName = "source-interface-table"
    sourceInterfaceTable.EntityData.BundleName = "cisco_ios_xr"
    sourceInterfaceTable.EntityData.ParentYangName = "syslog"
    sourceInterfaceTable.EntityData.SegmentPath = "source-interface-table"
    sourceInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInterfaceTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInterfaceTable.EntityData.Children = types.NewOrderedMap()
    sourceInterfaceTable.EntityData.Children.Append("source-interface-values", types.YChild{"SourceInterfaceValues", &sourceInterfaceTable.SourceInterfaceValues})
    sourceInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    sourceInterfaceTable.EntityData.YListKeys = []string {}

    return &(sourceInterfaceTable.EntityData)
}

// Syslog_SourceInterfaceTable_SourceInterfaceValues
// Specify interface for source address in logging
// transactions
type Syslog_SourceInterfaceTable_SourceInterfaceValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source interface. The type is slice of
    // Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue.
    SourceInterfaceValue []*Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue
}

func (sourceInterfaceValues *Syslog_SourceInterfaceTable_SourceInterfaceValues) GetEntityData() *types.CommonEntityData {
    sourceInterfaceValues.EntityData.YFilter = sourceInterfaceValues.YFilter
    sourceInterfaceValues.EntityData.YangName = "source-interface-values"
    sourceInterfaceValues.EntityData.BundleName = "cisco_ios_xr"
    sourceInterfaceValues.EntityData.ParentYangName = "source-interface-table"
    sourceInterfaceValues.EntityData.SegmentPath = "source-interface-values"
    sourceInterfaceValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInterfaceValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInterfaceValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInterfaceValues.EntityData.Children = types.NewOrderedMap()
    sourceInterfaceValues.EntityData.Children.Append("source-interface-value", types.YChild{"SourceInterfaceValue", nil})
    for i := range sourceInterfaceValues.SourceInterfaceValue {
        sourceInterfaceValues.EntityData.Children.Append(types.GetSegmentPath(sourceInterfaceValues.SourceInterfaceValue[i]), types.YChild{"SourceInterfaceValue", sourceInterfaceValues.SourceInterfaceValue[i]})
    }
    sourceInterfaceValues.EntityData.Leafs = types.NewOrderedMap()

    sourceInterfaceValues.EntityData.YListKeys = []string {}

    return &(sourceInterfaceValues.EntityData)
}

// Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue
// Source interface
type Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Which Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SrcInterfaceNameValue interface{}

    // Configure source interface VRF.
    SourceInterfaceVrfs Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs
}

func (sourceInterfaceValue *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue) GetEntityData() *types.CommonEntityData {
    sourceInterfaceValue.EntityData.YFilter = sourceInterfaceValue.YFilter
    sourceInterfaceValue.EntityData.YangName = "source-interface-value"
    sourceInterfaceValue.EntityData.BundleName = "cisco_ios_xr"
    sourceInterfaceValue.EntityData.ParentYangName = "source-interface-values"
    sourceInterfaceValue.EntityData.SegmentPath = "source-interface-value" + types.AddKeyToken(sourceInterfaceValue.SrcInterfaceNameValue, "src-interface-name-value")
    sourceInterfaceValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInterfaceValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInterfaceValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInterfaceValue.EntityData.Children = types.NewOrderedMap()
    sourceInterfaceValue.EntityData.Children.Append("source-interface-vrfs", types.YChild{"SourceInterfaceVrfs", &sourceInterfaceValue.SourceInterfaceVrfs})
    sourceInterfaceValue.EntityData.Leafs = types.NewOrderedMap()
    sourceInterfaceValue.EntityData.Leafs.Append("src-interface-name-value", types.YLeaf{"SrcInterfaceNameValue", sourceInterfaceValue.SrcInterfaceNameValue})

    sourceInterfaceValue.EntityData.YListKeys = []string {"SrcInterfaceNameValue"}

    return &(sourceInterfaceValue.EntityData)
}

// Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs
// Configure source interface VRF
type Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify VRF for source interface. The type is slice of
    // Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf.
    SourceInterfaceVrf []*Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf
}

func (sourceInterfaceVrfs *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs) GetEntityData() *types.CommonEntityData {
    sourceInterfaceVrfs.EntityData.YFilter = sourceInterfaceVrfs.YFilter
    sourceInterfaceVrfs.EntityData.YangName = "source-interface-vrfs"
    sourceInterfaceVrfs.EntityData.BundleName = "cisco_ios_xr"
    sourceInterfaceVrfs.EntityData.ParentYangName = "source-interface-value"
    sourceInterfaceVrfs.EntityData.SegmentPath = "source-interface-vrfs"
    sourceInterfaceVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInterfaceVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInterfaceVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInterfaceVrfs.EntityData.Children = types.NewOrderedMap()
    sourceInterfaceVrfs.EntityData.Children.Append("source-interface-vrf", types.YChild{"SourceInterfaceVrf", nil})
    for i := range sourceInterfaceVrfs.SourceInterfaceVrf {
        sourceInterfaceVrfs.EntityData.Children.Append(types.GetSegmentPath(sourceInterfaceVrfs.SourceInterfaceVrf[i]), types.YChild{"SourceInterfaceVrf", sourceInterfaceVrfs.SourceInterfaceVrf[i]})
    }
    sourceInterfaceVrfs.EntityData.Leafs = types.NewOrderedMap()

    sourceInterfaceVrfs.EntityData.YListKeys = []string {}

    return &(sourceInterfaceVrfs.EntityData)
}

// Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf
// Specify VRF for source interface
type Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}
}

func (sourceInterfaceVrf *Syslog_SourceInterfaceTable_SourceInterfaceValues_SourceInterfaceValue_SourceInterfaceVrfs_SourceInterfaceVrf) GetEntityData() *types.CommonEntityData {
    sourceInterfaceVrf.EntityData.YFilter = sourceInterfaceVrf.YFilter
    sourceInterfaceVrf.EntityData.YangName = "source-interface-vrf"
    sourceInterfaceVrf.EntityData.BundleName = "cisco_ios_xr"
    sourceInterfaceVrf.EntityData.ParentYangName = "source-interface-vrfs"
    sourceInterfaceVrf.EntityData.SegmentPath = "source-interface-vrf" + types.AddKeyToken(sourceInterfaceVrf.VrfName, "vrf-name")
    sourceInterfaceVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInterfaceVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInterfaceVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInterfaceVrf.EntityData.Children = types.NewOrderedMap()
    sourceInterfaceVrf.EntityData.Leafs = types.NewOrderedMap()
    sourceInterfaceVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", sourceInterfaceVrf.VrfName})

    sourceInterfaceVrf.EntityData.YListKeys = []string {"VrfName"}

    return &(sourceInterfaceVrf.EntityData)
}

// Syslog_AlarmLogger
// Alarm Logger Properties
type Syslog_AlarmLogger struct {
    EntityData types.CommonEntityData
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

func (alarmLogger *Syslog_AlarmLogger) GetEntityData() *types.CommonEntityData {
    alarmLogger.EntityData.YFilter = alarmLogger.YFilter
    alarmLogger.EntityData.YangName = "alarm-logger"
    alarmLogger.EntityData.BundleName = "cisco_ios_xr"
    alarmLogger.EntityData.ParentYangName = "syslog"
    alarmLogger.EntityData.SegmentPath = "Cisco-IOS-XR-infra-alarm-logger-cfg:alarm-logger"
    alarmLogger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmLogger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmLogger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmLogger.EntityData.Children = types.NewOrderedMap()
    alarmLogger.EntityData.Children.Append("alarm-filter-strings", types.YChild{"AlarmFilterStrings", &alarmLogger.AlarmFilterStrings})
    alarmLogger.EntityData.Leafs = types.NewOrderedMap()
    alarmLogger.EntityData.Leafs.Append("pre-config-suppression", types.YLeaf{"PreConfigSuppression", alarmLogger.PreConfigSuppression})
    alarmLogger.EntityData.Leafs.Append("severity-level", types.YLeaf{"SeverityLevel", alarmLogger.SeverityLevel})
    alarmLogger.EntityData.Leafs.Append("pre-config-suppression-timeout", types.YLeaf{"PreConfigSuppressionTimeout", alarmLogger.PreConfigSuppressionTimeout})
    alarmLogger.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", alarmLogger.BufferSize})
    alarmLogger.EntityData.Leafs.Append("source-location", types.YLeaf{"SourceLocation", alarmLogger.SourceLocation})
    alarmLogger.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", alarmLogger.Threshold})

    alarmLogger.EntityData.YListKeys = []string {}

    return &(alarmLogger.EntityData)
}

// Syslog_AlarmLogger_AlarmFilterStrings
// List of filter strings
type Syslog_AlarmLogger_AlarmFilterStrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match string to filter alarms. The type is slice of
    // Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString.
    AlarmFilterString []*Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString
}

func (alarmFilterStrings *Syslog_AlarmLogger_AlarmFilterStrings) GetEntityData() *types.CommonEntityData {
    alarmFilterStrings.EntityData.YFilter = alarmFilterStrings.YFilter
    alarmFilterStrings.EntityData.YangName = "alarm-filter-strings"
    alarmFilterStrings.EntityData.BundleName = "cisco_ios_xr"
    alarmFilterStrings.EntityData.ParentYangName = "alarm-logger"
    alarmFilterStrings.EntityData.SegmentPath = "alarm-filter-strings"
    alarmFilterStrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmFilterStrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmFilterStrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmFilterStrings.EntityData.Children = types.NewOrderedMap()
    alarmFilterStrings.EntityData.Children.Append("alarm-filter-string", types.YChild{"AlarmFilterString", nil})
    for i := range alarmFilterStrings.AlarmFilterString {
        alarmFilterStrings.EntityData.Children.Append(types.GetSegmentPath(alarmFilterStrings.AlarmFilterString[i]), types.YChild{"AlarmFilterString", alarmFilterStrings.AlarmFilterString[i]})
    }
    alarmFilterStrings.EntityData.Leafs = types.NewOrderedMap()

    alarmFilterStrings.EntityData.YListKeys = []string {}

    return &(alarmFilterStrings.EntityData)
}

// Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString
// Match string to filter alarms
type Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Filter String. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    FilterString interface{}
}

func (alarmFilterString *Syslog_AlarmLogger_AlarmFilterStrings_AlarmFilterString) GetEntityData() *types.CommonEntityData {
    alarmFilterString.EntityData.YFilter = alarmFilterString.YFilter
    alarmFilterString.EntityData.YangName = "alarm-filter-string"
    alarmFilterString.EntityData.BundleName = "cisco_ios_xr"
    alarmFilterString.EntityData.ParentYangName = "alarm-filter-strings"
    alarmFilterString.EntityData.SegmentPath = "alarm-filter-string" + types.AddKeyToken(alarmFilterString.FilterString, "filter-string")
    alarmFilterString.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmFilterString.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmFilterString.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmFilterString.EntityData.Children = types.NewOrderedMap()
    alarmFilterString.EntityData.Leafs = types.NewOrderedMap()
    alarmFilterString.EntityData.Leafs.Append("filter-string", types.YLeaf{"FilterString", alarmFilterString.FilterString})

    alarmFilterString.EntityData.YListKeys = []string {"FilterString"}

    return &(alarmFilterString.EntityData)
}

// Syslog_Correlator
// Configure properties of the event correlator
type Syslog_Correlator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure size of the correlator buffer. The type is interface{} with
    // range: 1024..52428800.
    BufferSize interface{}

    // Table of configured rules.
    Rules Syslog_Correlator_Rules

    // Table of configured rulesets.
    RuleSets Syslog_Correlator_RuleSets
}

func (correlator *Syslog_Correlator) GetEntityData() *types.CommonEntityData {
    correlator.EntityData.YFilter = correlator.YFilter
    correlator.EntityData.YangName = "correlator"
    correlator.EntityData.BundleName = "cisco_ios_xr"
    correlator.EntityData.ParentYangName = "syslog"
    correlator.EntityData.SegmentPath = "Cisco-IOS-XR-infra-correlator-cfg:correlator"
    correlator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    correlator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    correlator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    correlator.EntityData.Children = types.NewOrderedMap()
    correlator.EntityData.Children.Append("rules", types.YChild{"Rules", &correlator.Rules})
    correlator.EntityData.Children.Append("rule-sets", types.YChild{"RuleSets", &correlator.RuleSets})
    correlator.EntityData.Leafs = types.NewOrderedMap()
    correlator.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", correlator.BufferSize})

    correlator.EntityData.YListKeys = []string {}

    return &(correlator.EntityData)
}

// Syslog_Correlator_Rules
// Table of configured rules
type Syslog_Correlator_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rule name. The type is slice of Syslog_Correlator_Rules_Rule.
    Rule []*Syslog_Correlator_Rules_Rule
}

func (rules *Syslog_Correlator_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "correlator"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = types.NewOrderedMap()
    rules.EntityData.Children.Append("rule", types.YChild{"Rule", nil})
    for i := range rules.Rule {
        rules.EntityData.Children.Append(types.GetSegmentPath(rules.Rule[i]), types.YChild{"Rule", rules.Rule[i]})
    }
    rules.EntityData.Leafs = types.NewOrderedMap()

    rules.EntityData.YListKeys = []string {}

    return &(rules.EntityData)
}

// Syslog_Correlator_Rules_Rule
// Rule name
type Syslog_Correlator_Rules_Rule struct {
    EntityData types.CommonEntityData
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

func (rule *Syslog_Correlator_Rules_Rule) GetEntityData() *types.CommonEntityData {
    rule.EntityData.YFilter = rule.YFilter
    rule.EntityData.YangName = "rule"
    rule.EntityData.BundleName = "cisco_ios_xr"
    rule.EntityData.ParentYangName = "rules"
    rule.EntityData.SegmentPath = "rule" + types.AddKeyToken(rule.Name, "name")
    rule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rule.EntityData.Children = types.NewOrderedMap()
    rule.EntityData.Children.Append("definition", types.YChild{"Definition", &rule.Definition})
    rule.EntityData.Children.Append("non-stateful", types.YChild{"NonStateful", &rule.NonStateful})
    rule.EntityData.Children.Append("stateful", types.YChild{"Stateful", &rule.Stateful})
    rule.EntityData.Children.Append("apply-to", types.YChild{"ApplyTo", &rule.ApplyTo})
    rule.EntityData.Children.Append("applied-to", types.YChild{"AppliedTo", &rule.AppliedTo})
    rule.EntityData.Leafs = types.NewOrderedMap()
    rule.EntityData.Leafs.Append("name", types.YLeaf{"Name", rule.Name})

    rule.EntityData.YListKeys = []string {"Name"}

    return &(rule.EntityData)
}

// Syslog_Correlator_Rules_Rule_Definition
// Configure a specified correlation rule
type Syslog_Correlator_Rules_Rule_Definition struct {
    EntityData types.CommonEntityData
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

func (definition *Syslog_Correlator_Rules_Rule_Definition) GetEntityData() *types.CommonEntityData {
    definition.EntityData.YFilter = definition.YFilter
    definition.EntityData.YangName = "definition"
    definition.EntityData.BundleName = "cisco_ios_xr"
    definition.EntityData.ParentYangName = "rule"
    definition.EntityData.SegmentPath = "definition"
    definition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    definition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    definition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    definition.EntityData.Children = types.NewOrderedMap()
    definition.EntityData.Leafs = types.NewOrderedMap()
    definition.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", definition.Timeout})
    definition.EntityData.Leafs.Append("category-name-entry1", types.YLeaf{"CategoryNameEntry1", definition.CategoryNameEntry1})
    definition.EntityData.Leafs.Append("group-name-entry1", types.YLeaf{"GroupNameEntry1", definition.GroupNameEntry1})
    definition.EntityData.Leafs.Append("message-code-entry1", types.YLeaf{"MessageCodeEntry1", definition.MessageCodeEntry1})
    definition.EntityData.Leafs.Append("category-name-entry2", types.YLeaf{"CategoryNameEntry2", definition.CategoryNameEntry2})
    definition.EntityData.Leafs.Append("group-name-entry2", types.YLeaf{"GroupNameEntry2", definition.GroupNameEntry2})
    definition.EntityData.Leafs.Append("message-code-entry2", types.YLeaf{"MessageCodeEntry2", definition.MessageCodeEntry2})
    definition.EntityData.Leafs.Append("category-name-entry3", types.YLeaf{"CategoryNameEntry3", definition.CategoryNameEntry3})
    definition.EntityData.Leafs.Append("group-name-entry3", types.YLeaf{"GroupNameEntry3", definition.GroupNameEntry3})
    definition.EntityData.Leafs.Append("message-code-entry3", types.YLeaf{"MessageCodeEntry3", definition.MessageCodeEntry3})
    definition.EntityData.Leafs.Append("category-name-entry4", types.YLeaf{"CategoryNameEntry4", definition.CategoryNameEntry4})
    definition.EntityData.Leafs.Append("group-name-entry4", types.YLeaf{"GroupNameEntry4", definition.GroupNameEntry4})
    definition.EntityData.Leafs.Append("message-code-entry4", types.YLeaf{"MessageCodeEntry4", definition.MessageCodeEntry4})
    definition.EntityData.Leafs.Append("category-name-entry5", types.YLeaf{"CategoryNameEntry5", definition.CategoryNameEntry5})
    definition.EntityData.Leafs.Append("group-name-entry5", types.YLeaf{"GroupNameEntry5", definition.GroupNameEntry5})
    definition.EntityData.Leafs.Append("message-code-entry5", types.YLeaf{"MessageCodeEntry5", definition.MessageCodeEntry5})
    definition.EntityData.Leafs.Append("category-name-entry6", types.YLeaf{"CategoryNameEntry6", definition.CategoryNameEntry6})
    definition.EntityData.Leafs.Append("group-name-entry6", types.YLeaf{"GroupNameEntry6", definition.GroupNameEntry6})
    definition.EntityData.Leafs.Append("message-code-entry6", types.YLeaf{"MessageCodeEntry6", definition.MessageCodeEntry6})
    definition.EntityData.Leafs.Append("category-name-entry7", types.YLeaf{"CategoryNameEntry7", definition.CategoryNameEntry7})
    definition.EntityData.Leafs.Append("group-name-entry7", types.YLeaf{"GroupNameEntry7", definition.GroupNameEntry7})
    definition.EntityData.Leafs.Append("message-code-entry7", types.YLeaf{"MessageCodeEntry7", definition.MessageCodeEntry7})
    definition.EntityData.Leafs.Append("category-name-entry8", types.YLeaf{"CategoryNameEntry8", definition.CategoryNameEntry8})
    definition.EntityData.Leafs.Append("group-name-entry8", types.YLeaf{"GroupNameEntry8", definition.GroupNameEntry8})
    definition.EntityData.Leafs.Append("message-code-entry8", types.YLeaf{"MessageCodeEntry8", definition.MessageCodeEntry8})
    definition.EntityData.Leafs.Append("category-name-entry9", types.YLeaf{"CategoryNameEntry9", definition.CategoryNameEntry9})
    definition.EntityData.Leafs.Append("group-name-entry9", types.YLeaf{"GroupNameEntry9", definition.GroupNameEntry9})
    definition.EntityData.Leafs.Append("message-code-entry9", types.YLeaf{"MessageCodeEntry9", definition.MessageCodeEntry9})
    definition.EntityData.Leafs.Append("category-name-entry10", types.YLeaf{"CategoryNameEntry10", definition.CategoryNameEntry10})
    definition.EntityData.Leafs.Append("group-name-entry10", types.YLeaf{"GroupNameEntry10", definition.GroupNameEntry10})
    definition.EntityData.Leafs.Append("message-code-entry10", types.YLeaf{"MessageCodeEntry10", definition.MessageCodeEntry10})

    definition.EntityData.YListKeys = []string {}

    return &(definition.EntityData)
}

// Syslog_Correlator_Rules_Rule_NonStateful
// The Non-Stateful Rule Type
type Syslog_Correlator_Rules_Rule_NonStateful struct {
    EntityData types.CommonEntityData
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

func (nonStateful *Syslog_Correlator_Rules_Rule_NonStateful) GetEntityData() *types.CommonEntityData {
    nonStateful.EntityData.YFilter = nonStateful.YFilter
    nonStateful.EntityData.YangName = "non-stateful"
    nonStateful.EntityData.BundleName = "cisco_ios_xr"
    nonStateful.EntityData.ParentYangName = "rule"
    nonStateful.EntityData.SegmentPath = "non-stateful"
    nonStateful.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonStateful.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonStateful.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonStateful.EntityData.Children = types.NewOrderedMap()
    nonStateful.EntityData.Children.Append("non-root-causes", types.YChild{"NonRootCauses", &nonStateful.NonRootCauses})
    nonStateful.EntityData.Children.Append("root-cause", types.YChild{"RootCause", &nonStateful.RootCause})
    nonStateful.EntityData.Leafs = types.NewOrderedMap()
    nonStateful.EntityData.Leafs.Append("context-correlation", types.YLeaf{"ContextCorrelation", nonStateful.ContextCorrelation})
    nonStateful.EntityData.Leafs.Append("timeout-root-cause", types.YLeaf{"TimeoutRootCause", nonStateful.TimeoutRootCause})
    nonStateful.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", nonStateful.Timeout})

    nonStateful.EntityData.YListKeys = []string {}

    return &(nonStateful.EntityData)
}

// Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses
// Table of configured non-rootcause
type Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A non-rootcause. The type is slice of
    // Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause.
    NonRootCause []*Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetEntityData() *types.CommonEntityData {
    nonRootCauses.EntityData.YFilter = nonRootCauses.YFilter
    nonRootCauses.EntityData.YangName = "non-root-causes"
    nonRootCauses.EntityData.BundleName = "cisco_ios_xr"
    nonRootCauses.EntityData.ParentYangName = "non-stateful"
    nonRootCauses.EntityData.SegmentPath = "non-root-causes"
    nonRootCauses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootCauses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootCauses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootCauses.EntityData.Children = types.NewOrderedMap()
    nonRootCauses.EntityData.Children.Append("non-root-cause", types.YChild{"NonRootCause", nil})
    for i := range nonRootCauses.NonRootCause {
        nonRootCauses.EntityData.Children.Append(types.GetSegmentPath(nonRootCauses.NonRootCause[i]), types.YChild{"NonRootCause", nonRootCauses.NonRootCause[i]})
    }
    nonRootCauses.EntityData.Leafs = types.NewOrderedMap()

    nonRootCauses.EntityData.YListKeys = []string {}

    return &(nonRootCauses.EntityData)
}

// Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause
// A non-rootcause
type Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Correlated message category. The type is string.
    Category interface{}

    // This attribute is a key. Correlated message group. The type is string.
    Group interface{}

    // This attribute is a key. Correlated message code. The type is string.
    MessageCode interface{}
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetEntityData() *types.CommonEntityData {
    nonRootCause.EntityData.YFilter = nonRootCause.YFilter
    nonRootCause.EntityData.YangName = "non-root-cause"
    nonRootCause.EntityData.BundleName = "cisco_ios_xr"
    nonRootCause.EntityData.ParentYangName = "non-root-causes"
    nonRootCause.EntityData.SegmentPath = "non-root-cause" + types.AddKeyToken(nonRootCause.Category, "category") + types.AddKeyToken(nonRootCause.Group, "group") + types.AddKeyToken(nonRootCause.MessageCode, "message-code")
    nonRootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootCause.EntityData.Children = types.NewOrderedMap()
    nonRootCause.EntityData.Leafs = types.NewOrderedMap()
    nonRootCause.EntityData.Leafs.Append("category", types.YLeaf{"Category", nonRootCause.Category})
    nonRootCause.EntityData.Leafs.Append("group", types.YLeaf{"Group", nonRootCause.Group})
    nonRootCause.EntityData.Leafs.Append("message-code", types.YLeaf{"MessageCode", nonRootCause.MessageCode})

    nonRootCause.EntityData.YListKeys = []string {"Category", "Group", "MessageCode"}

    return &(nonRootCause.EntityData)
}

// Syslog_Correlator_Rules_Rule_NonStateful_RootCause
// The root cause
type Syslog_Correlator_Rules_Rule_NonStateful_RootCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Root message category. The type is string.
    Category interface{}

    // Root message group. The type is string.
    Group interface{}

    // Root message code. The type is string.
    MessageCode interface{}
}

func (rootCause *Syslog_Correlator_Rules_Rule_NonStateful_RootCause) GetEntityData() *types.CommonEntityData {
    rootCause.EntityData.YFilter = rootCause.YFilter
    rootCause.EntityData.YangName = "root-cause"
    rootCause.EntityData.BundleName = "cisco_ios_xr"
    rootCause.EntityData.ParentYangName = "non-stateful"
    rootCause.EntityData.SegmentPath = "root-cause"
    rootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rootCause.EntityData.Children = types.NewOrderedMap()
    rootCause.EntityData.Leafs = types.NewOrderedMap()
    rootCause.EntityData.Leafs.Append("category", types.YLeaf{"Category", rootCause.Category})
    rootCause.EntityData.Leafs.Append("group", types.YLeaf{"Group", rootCause.Group})
    rootCause.EntityData.Leafs.Append("message-code", types.YLeaf{"MessageCode", rootCause.MessageCode})

    rootCause.EntityData.YListKeys = []string {}

    return &(rootCause.EntityData)
}

// Syslog_Correlator_Rules_Rule_Stateful
// The Stateful Rule Type
type Syslog_Correlator_Rules_Rule_Stateful struct {
    EntityData types.CommonEntityData
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

func (stateful *Syslog_Correlator_Rules_Rule_Stateful) GetEntityData() *types.CommonEntityData {
    stateful.EntityData.YFilter = stateful.YFilter
    stateful.EntityData.YangName = "stateful"
    stateful.EntityData.BundleName = "cisco_ios_xr"
    stateful.EntityData.ParentYangName = "rule"
    stateful.EntityData.SegmentPath = "stateful"
    stateful.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateful.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateful.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateful.EntityData.Children = types.NewOrderedMap()
    stateful.EntityData.Children.Append("non-root-causes", types.YChild{"NonRootCauses", &stateful.NonRootCauses})
    stateful.EntityData.Children.Append("root-cause", types.YChild{"RootCause", &stateful.RootCause})
    stateful.EntityData.Leafs = types.NewOrderedMap()
    stateful.EntityData.Leafs.Append("reparent", types.YLeaf{"Reparent", stateful.Reparent})
    stateful.EntityData.Leafs.Append("reissue", types.YLeaf{"Reissue", stateful.Reissue})
    stateful.EntityData.Leafs.Append("context-correlation", types.YLeaf{"ContextCorrelation", stateful.ContextCorrelation})
    stateful.EntityData.Leafs.Append("timeout-root-cause", types.YLeaf{"TimeoutRootCause", stateful.TimeoutRootCause})
    stateful.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", stateful.Timeout})

    stateful.EntityData.YListKeys = []string {}

    return &(stateful.EntityData)
}

// Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses
// Table of configured non-rootcause
type Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A non-rootcause. The type is slice of
    // Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause.
    NonRootCause []*Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause
}

func (nonRootCauses *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses) GetEntityData() *types.CommonEntityData {
    nonRootCauses.EntityData.YFilter = nonRootCauses.YFilter
    nonRootCauses.EntityData.YangName = "non-root-causes"
    nonRootCauses.EntityData.BundleName = "cisco_ios_xr"
    nonRootCauses.EntityData.ParentYangName = "stateful"
    nonRootCauses.EntityData.SegmentPath = "non-root-causes"
    nonRootCauses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootCauses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootCauses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootCauses.EntityData.Children = types.NewOrderedMap()
    nonRootCauses.EntityData.Children.Append("non-root-cause", types.YChild{"NonRootCause", nil})
    for i := range nonRootCauses.NonRootCause {
        nonRootCauses.EntityData.Children.Append(types.GetSegmentPath(nonRootCauses.NonRootCause[i]), types.YChild{"NonRootCause", nonRootCauses.NonRootCause[i]})
    }
    nonRootCauses.EntityData.Leafs = types.NewOrderedMap()

    nonRootCauses.EntityData.YListKeys = []string {}

    return &(nonRootCauses.EntityData)
}

// Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause
// A non-rootcause
type Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Correlated message category. The type is string.
    Category interface{}

    // This attribute is a key. Correlated message group. The type is string.
    Group interface{}

    // This attribute is a key. Correlated message code. The type is string.
    MessageCode interface{}
}

func (nonRootCause *Syslog_Correlator_Rules_Rule_Stateful_NonRootCauses_NonRootCause) GetEntityData() *types.CommonEntityData {
    nonRootCause.EntityData.YFilter = nonRootCause.YFilter
    nonRootCause.EntityData.YangName = "non-root-cause"
    nonRootCause.EntityData.BundleName = "cisco_ios_xr"
    nonRootCause.EntityData.ParentYangName = "non-root-causes"
    nonRootCause.EntityData.SegmentPath = "non-root-cause" + types.AddKeyToken(nonRootCause.Category, "category") + types.AddKeyToken(nonRootCause.Group, "group") + types.AddKeyToken(nonRootCause.MessageCode, "message-code")
    nonRootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootCause.EntityData.Children = types.NewOrderedMap()
    nonRootCause.EntityData.Leafs = types.NewOrderedMap()
    nonRootCause.EntityData.Leafs.Append("category", types.YLeaf{"Category", nonRootCause.Category})
    nonRootCause.EntityData.Leafs.Append("group", types.YLeaf{"Group", nonRootCause.Group})
    nonRootCause.EntityData.Leafs.Append("message-code", types.YLeaf{"MessageCode", nonRootCause.MessageCode})

    nonRootCause.EntityData.YListKeys = []string {"Category", "Group", "MessageCode"}

    return &(nonRootCause.EntityData)
}

// Syslog_Correlator_Rules_Rule_Stateful_RootCause
// The root cause
type Syslog_Correlator_Rules_Rule_Stateful_RootCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Root message category. The type is string.
    Category interface{}

    // Root message group. The type is string.
    Group interface{}

    // Root message code. The type is string.
    MessageCode interface{}
}

func (rootCause *Syslog_Correlator_Rules_Rule_Stateful_RootCause) GetEntityData() *types.CommonEntityData {
    rootCause.EntityData.YFilter = rootCause.YFilter
    rootCause.EntityData.YangName = "root-cause"
    rootCause.EntityData.BundleName = "cisco_ios_xr"
    rootCause.EntityData.ParentYangName = "stateful"
    rootCause.EntityData.SegmentPath = "root-cause"
    rootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rootCause.EntityData.Children = types.NewOrderedMap()
    rootCause.EntityData.Leafs = types.NewOrderedMap()
    rootCause.EntityData.Leafs.Append("category", types.YLeaf{"Category", rootCause.Category})
    rootCause.EntityData.Leafs.Append("group", types.YLeaf{"Group", rootCause.Group})
    rootCause.EntityData.Leafs.Append("message-code", types.YLeaf{"MessageCode", rootCause.MessageCode})

    rootCause.EntityData.YListKeys = []string {}

    return &(rootCause.EntityData)
}

// Syslog_Correlator_Rules_Rule_ApplyTo
// Apply the Rules
type Syslog_Correlator_Rules_Rule_ApplyTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Apply the rule to all of the router. The type is interface{}.
    AllOfRouter interface{}

    // Apply rule to a specified list of contexts, e.g. interfaces.
    Contexts Syslog_Correlator_Rules_Rule_ApplyTo_Contexts

    // Apply rule to a specified list of Locations.
    Locations Syslog_Correlator_Rules_Rule_ApplyTo_Locations
}

func (applyTo *Syslog_Correlator_Rules_Rule_ApplyTo) GetEntityData() *types.CommonEntityData {
    applyTo.EntityData.YFilter = applyTo.YFilter
    applyTo.EntityData.YangName = "apply-to"
    applyTo.EntityData.BundleName = "cisco_ios_xr"
    applyTo.EntityData.ParentYangName = "rule"
    applyTo.EntityData.SegmentPath = "apply-to"
    applyTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applyTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applyTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applyTo.EntityData.Children = types.NewOrderedMap()
    applyTo.EntityData.Children.Append("contexts", types.YChild{"Contexts", &applyTo.Contexts})
    applyTo.EntityData.Children.Append("locations", types.YChild{"Locations", &applyTo.Locations})
    applyTo.EntityData.Leafs = types.NewOrderedMap()
    applyTo.EntityData.Leafs.Append("all-of-router", types.YLeaf{"AllOfRouter", applyTo.AllOfRouter})

    applyTo.EntityData.YListKeys = []string {}

    return &(applyTo.EntityData)
}

// Syslog_Correlator_Rules_Rule_ApplyTo_Contexts
// Apply rule to a specified list of contexts,
// e.g. interfaces
type Syslog_Correlator_Rules_Rule_ApplyTo_Contexts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One or more context names. The type is slice of string.
    Context []interface{}
}

func (contexts *Syslog_Correlator_Rules_Rule_ApplyTo_Contexts) GetEntityData() *types.CommonEntityData {
    contexts.EntityData.YFilter = contexts.YFilter
    contexts.EntityData.YangName = "contexts"
    contexts.EntityData.BundleName = "cisco_ios_xr"
    contexts.EntityData.ParentYangName = "apply-to"
    contexts.EntityData.SegmentPath = "contexts"
    contexts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contexts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contexts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contexts.EntityData.Children = types.NewOrderedMap()
    contexts.EntityData.Leafs = types.NewOrderedMap()
    contexts.EntityData.Leafs.Append("context", types.YLeaf{"Context", contexts.Context})

    contexts.EntityData.YListKeys = []string {}

    return &(contexts.EntityData)
}

// Syslog_Correlator_Rules_Rule_ApplyTo_Locations
// Apply rule to a specified list of Locations
type Syslog_Correlator_Rules_Rule_ApplyTo_Locations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One or more Locations. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location []interface{}
}

func (locations *Syslog_Correlator_Rules_Rule_ApplyTo_Locations) GetEntityData() *types.CommonEntityData {
    locations.EntityData.YFilter = locations.YFilter
    locations.EntityData.YangName = "locations"
    locations.EntityData.BundleName = "cisco_ios_xr"
    locations.EntityData.ParentYangName = "apply-to"
    locations.EntityData.SegmentPath = "locations"
    locations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locations.EntityData.Children = types.NewOrderedMap()
    locations.EntityData.Leafs = types.NewOrderedMap()
    locations.EntityData.Leafs.Append("location", types.YLeaf{"Location", locations.Location})

    locations.EntityData.YListKeys = []string {}

    return &(locations.EntityData)
}

// Syslog_Correlator_Rules_Rule_AppliedTo
// Applied to the Rule or Ruleset
type Syslog_Correlator_Rules_Rule_AppliedTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Apply to all of the router. The type is interface{}.
    All interface{}

    // Table of configured contexts to apply.
    Contexts Syslog_Correlator_Rules_Rule_AppliedTo_Contexts

    // Table of configured locations to apply.
    Locations Syslog_Correlator_Rules_Rule_AppliedTo_Locations
}

func (appliedTo *Syslog_Correlator_Rules_Rule_AppliedTo) GetEntityData() *types.CommonEntityData {
    appliedTo.EntityData.YFilter = appliedTo.YFilter
    appliedTo.EntityData.YangName = "applied-to"
    appliedTo.EntityData.BundleName = "cisco_ios_xr"
    appliedTo.EntityData.ParentYangName = "rule"
    appliedTo.EntityData.SegmentPath = "applied-to"
    appliedTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appliedTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appliedTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appliedTo.EntityData.Children = types.NewOrderedMap()
    appliedTo.EntityData.Children.Append("contexts", types.YChild{"Contexts", &appliedTo.Contexts})
    appliedTo.EntityData.Children.Append("locations", types.YChild{"Locations", &appliedTo.Locations})
    appliedTo.EntityData.Leafs = types.NewOrderedMap()
    appliedTo.EntityData.Leafs.Append("all", types.YLeaf{"All", appliedTo.All})

    appliedTo.EntityData.YListKeys = []string {}

    return &(appliedTo.EntityData)
}

// Syslog_Correlator_Rules_Rule_AppliedTo_Contexts
// Table of configured contexts to apply
type Syslog_Correlator_Rules_Rule_AppliedTo_Contexts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A context. The type is slice of
    // Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context.
    Context []*Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context
}

func (contexts *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts) GetEntityData() *types.CommonEntityData {
    contexts.EntityData.YFilter = contexts.YFilter
    contexts.EntityData.YangName = "contexts"
    contexts.EntityData.BundleName = "cisco_ios_xr"
    contexts.EntityData.ParentYangName = "applied-to"
    contexts.EntityData.SegmentPath = "contexts"
    contexts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contexts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contexts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contexts.EntityData.Children = types.NewOrderedMap()
    contexts.EntityData.Children.Append("context", types.YChild{"Context", nil})
    for i := range contexts.Context {
        contexts.EntityData.Children.Append(types.GetSegmentPath(contexts.Context[i]), types.YChild{"Context", contexts.Context[i]})
    }
    contexts.EntityData.Leafs = types.NewOrderedMap()

    contexts.EntityData.YListKeys = []string {}

    return &(contexts.EntityData)
}

// Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context
// A context
type Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context. The type is string with length: 1..32.
    Context interface{}
}

func (context *Syslog_Correlator_Rules_Rule_AppliedTo_Contexts_Context) GetEntityData() *types.CommonEntityData {
    context.EntityData.YFilter = context.YFilter
    context.EntityData.YangName = "context"
    context.EntityData.BundleName = "cisco_ios_xr"
    context.EntityData.ParentYangName = "contexts"
    context.EntityData.SegmentPath = "context" + types.AddKeyToken(context.Context, "context")
    context.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    context.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    context.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    context.EntityData.Children = types.NewOrderedMap()
    context.EntityData.Leafs = types.NewOrderedMap()
    context.EntityData.Leafs.Append("context", types.YLeaf{"Context", context.Context})

    context.EntityData.YListKeys = []string {"Context"}

    return &(context.EntityData)
}

// Syslog_Correlator_Rules_Rule_AppliedTo_Locations
// Table of configured locations to apply
type Syslog_Correlator_Rules_Rule_AppliedTo_Locations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A location. The type is slice of
    // Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location.
    Location []*Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location
}

func (locations *Syslog_Correlator_Rules_Rule_AppliedTo_Locations) GetEntityData() *types.CommonEntityData {
    locations.EntityData.YFilter = locations.YFilter
    locations.EntityData.YangName = "locations"
    locations.EntityData.BundleName = "cisco_ios_xr"
    locations.EntityData.ParentYangName = "applied-to"
    locations.EntityData.SegmentPath = "locations"
    locations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locations.EntityData.Children = types.NewOrderedMap()
    locations.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range locations.Location {
        locations.EntityData.Children.Append(types.GetSegmentPath(locations.Location[i]), types.YChild{"Location", locations.Location[i]})
    }
    locations.EntityData.Leafs = types.NewOrderedMap()

    locations.EntityData.YListKeys = []string {}

    return &(locations.EntityData)
}

// Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location
// A location
type Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}
}

func (location *Syslog_Correlator_Rules_Rule_AppliedTo_Locations_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "locations"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Syslog_Correlator_RuleSets
// Table of configured rulesets
type Syslog_Correlator_RuleSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ruleset name. The type is slice of Syslog_Correlator_RuleSets_RuleSet.
    RuleSet []*Syslog_Correlator_RuleSets_RuleSet
}

func (ruleSets *Syslog_Correlator_RuleSets) GetEntityData() *types.CommonEntityData {
    ruleSets.EntityData.YFilter = ruleSets.YFilter
    ruleSets.EntityData.YangName = "rule-sets"
    ruleSets.EntityData.BundleName = "cisco_ios_xr"
    ruleSets.EntityData.ParentYangName = "correlator"
    ruleSets.EntityData.SegmentPath = "rule-sets"
    ruleSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSets.EntityData.Children = types.NewOrderedMap()
    ruleSets.EntityData.Children.Append("rule-set", types.YChild{"RuleSet", nil})
    for i := range ruleSets.RuleSet {
        ruleSets.EntityData.Children.Append(types.GetSegmentPath(ruleSets.RuleSet[i]), types.YChild{"RuleSet", ruleSets.RuleSet[i]})
    }
    ruleSets.EntityData.Leafs = types.NewOrderedMap()

    ruleSets.EntityData.YListKeys = []string {}

    return &(ruleSets.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet
// Ruleset name
type Syslog_Correlator_RuleSets_RuleSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ruleset name. The type is string with length:
    // 1..32.
    Name interface{}

    // Table of configured rulenames.
    Rulenames Syslog_Correlator_RuleSets_RuleSet_Rulenames

    // Applied to the Rule or Ruleset.
    AppliedTo Syslog_Correlator_RuleSets_RuleSet_AppliedTo
}

func (ruleSet *Syslog_Correlator_RuleSets_RuleSet) GetEntityData() *types.CommonEntityData {
    ruleSet.EntityData.YFilter = ruleSet.YFilter
    ruleSet.EntityData.YangName = "rule-set"
    ruleSet.EntityData.BundleName = "cisco_ios_xr"
    ruleSet.EntityData.ParentYangName = "rule-sets"
    ruleSet.EntityData.SegmentPath = "rule-set" + types.AddKeyToken(ruleSet.Name, "name")
    ruleSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSet.EntityData.Children = types.NewOrderedMap()
    ruleSet.EntityData.Children.Append("rulenames", types.YChild{"Rulenames", &ruleSet.Rulenames})
    ruleSet.EntityData.Children.Append("applied-to", types.YChild{"AppliedTo", &ruleSet.AppliedTo})
    ruleSet.EntityData.Leafs = types.NewOrderedMap()
    ruleSet.EntityData.Leafs.Append("name", types.YLeaf{"Name", ruleSet.Name})

    ruleSet.EntityData.YListKeys = []string {"Name"}

    return &(ruleSet.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet_Rulenames
// Table of configured rulenames
type Syslog_Correlator_RuleSets_RuleSet_Rulenames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A rulename. The type is slice of
    // Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename.
    Rulename []*Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename
}

func (rulenames *Syslog_Correlator_RuleSets_RuleSet_Rulenames) GetEntityData() *types.CommonEntityData {
    rulenames.EntityData.YFilter = rulenames.YFilter
    rulenames.EntityData.YangName = "rulenames"
    rulenames.EntityData.BundleName = "cisco_ios_xr"
    rulenames.EntityData.ParentYangName = "rule-set"
    rulenames.EntityData.SegmentPath = "rulenames"
    rulenames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulenames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulenames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulenames.EntityData.Children = types.NewOrderedMap()
    rulenames.EntityData.Children.Append("rulename", types.YChild{"Rulename", nil})
    for i := range rulenames.Rulename {
        rulenames.EntityData.Children.Append(types.GetSegmentPath(rulenames.Rulename[i]), types.YChild{"Rulename", rulenames.Rulename[i]})
    }
    rulenames.EntityData.Leafs = types.NewOrderedMap()

    rulenames.EntityData.YListKeys = []string {}

    return &(rulenames.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename
// A rulename
type Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Rulename interface{}
}

func (rulename *Syslog_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetEntityData() *types.CommonEntityData {
    rulename.EntityData.YFilter = rulename.YFilter
    rulename.EntityData.YangName = "rulename"
    rulename.EntityData.BundleName = "cisco_ios_xr"
    rulename.EntityData.ParentYangName = "rulenames"
    rulename.EntityData.SegmentPath = "rulename" + types.AddKeyToken(rulename.Rulename, "rulename")
    rulename.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulename.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulename.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulename.EntityData.Children = types.NewOrderedMap()
    rulename.EntityData.Leafs = types.NewOrderedMap()
    rulename.EntityData.Leafs.Append("rulename", types.YLeaf{"Rulename", rulename.Rulename})

    rulename.EntityData.YListKeys = []string {"Rulename"}

    return &(rulename.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo
// Applied to the Rule or Ruleset
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Apply to all of the router. The type is interface{}.
    All interface{}

    // Table of configured contexts to apply.
    Contexts Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts

    // Table of configured locations to apply.
    Locations Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations
}

func (appliedTo *Syslog_Correlator_RuleSets_RuleSet_AppliedTo) GetEntityData() *types.CommonEntityData {
    appliedTo.EntityData.YFilter = appliedTo.YFilter
    appliedTo.EntityData.YangName = "applied-to"
    appliedTo.EntityData.BundleName = "cisco_ios_xr"
    appliedTo.EntityData.ParentYangName = "rule-set"
    appliedTo.EntityData.SegmentPath = "applied-to"
    appliedTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appliedTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appliedTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appliedTo.EntityData.Children = types.NewOrderedMap()
    appliedTo.EntityData.Children.Append("contexts", types.YChild{"Contexts", &appliedTo.Contexts})
    appliedTo.EntityData.Children.Append("locations", types.YChild{"Locations", &appliedTo.Locations})
    appliedTo.EntityData.Leafs = types.NewOrderedMap()
    appliedTo.EntityData.Leafs.Append("all", types.YLeaf{"All", appliedTo.All})

    appliedTo.EntityData.YListKeys = []string {}

    return &(appliedTo.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts
// Table of configured contexts to apply
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A context. The type is slice of
    // Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context.
    Context []*Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context
}

func (contexts *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts) GetEntityData() *types.CommonEntityData {
    contexts.EntityData.YFilter = contexts.YFilter
    contexts.EntityData.YangName = "contexts"
    contexts.EntityData.BundleName = "cisco_ios_xr"
    contexts.EntityData.ParentYangName = "applied-to"
    contexts.EntityData.SegmentPath = "contexts"
    contexts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contexts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contexts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contexts.EntityData.Children = types.NewOrderedMap()
    contexts.EntityData.Children.Append("context", types.YChild{"Context", nil})
    for i := range contexts.Context {
        contexts.EntityData.Children.Append(types.GetSegmentPath(contexts.Context[i]), types.YChild{"Context", contexts.Context[i]})
    }
    contexts.EntityData.Leafs = types.NewOrderedMap()

    contexts.EntityData.YListKeys = []string {}

    return &(contexts.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context
// A context
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context. The type is string with length: 1..32.
    Context interface{}
}

func (context *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Contexts_Context) GetEntityData() *types.CommonEntityData {
    context.EntityData.YFilter = context.YFilter
    context.EntityData.YangName = "context"
    context.EntityData.BundleName = "cisco_ios_xr"
    context.EntityData.ParentYangName = "contexts"
    context.EntityData.SegmentPath = "context" + types.AddKeyToken(context.Context, "context")
    context.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    context.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    context.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    context.EntityData.Children = types.NewOrderedMap()
    context.EntityData.Leafs = types.NewOrderedMap()
    context.EntityData.Leafs.Append("context", types.YLeaf{"Context", context.Context})

    context.EntityData.YListKeys = []string {"Context"}

    return &(context.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations
// Table of configured locations to apply
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A location. The type is slice of
    // Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location.
    Location []*Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location
}

func (locations *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations) GetEntityData() *types.CommonEntityData {
    locations.EntityData.YFilter = locations.YFilter
    locations.EntityData.YangName = "locations"
    locations.EntityData.BundleName = "cisco_ios_xr"
    locations.EntityData.ParentYangName = "applied-to"
    locations.EntityData.SegmentPath = "locations"
    locations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locations.EntityData.Children = types.NewOrderedMap()
    locations.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range locations.Location {
        locations.EntityData.Children.Append(types.GetSegmentPath(locations.Location[i]), types.YChild{"Location", locations.Location[i]})
    }
    locations.EntityData.Leafs = types.NewOrderedMap()

    locations.EntityData.YListKeys = []string {}

    return &(locations.EntityData)
}

// Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location
// A location
type Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}
}

func (location *Syslog_Correlator_RuleSets_RuleSet_AppliedTo_Locations_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "locations"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Syslog_Suppression
// Configure properties of the syslog/alarm
// suppression
type Syslog_Suppression struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of configured rules.
    Rules Syslog_Suppression_Rules
}

func (suppression *Syslog_Suppression) GetEntityData() *types.CommonEntityData {
    suppression.EntityData.YFilter = suppression.YFilter
    suppression.EntityData.YangName = "suppression"
    suppression.EntityData.BundleName = "cisco_ios_xr"
    suppression.EntityData.ParentYangName = "syslog"
    suppression.EntityData.SegmentPath = "Cisco-IOS-XR-infra-correlator-cfg:suppression"
    suppression.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppression.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppression.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppression.EntityData.Children = types.NewOrderedMap()
    suppression.EntityData.Children.Append("rules", types.YChild{"Rules", &suppression.Rules})
    suppression.EntityData.Leafs = types.NewOrderedMap()

    suppression.EntityData.YListKeys = []string {}

    return &(suppression.EntityData)
}

// Syslog_Suppression_Rules
// Table of configured rules
type Syslog_Suppression_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rule name. The type is slice of Syslog_Suppression_Rules_Rule.
    Rule []*Syslog_Suppression_Rules_Rule
}

func (rules *Syslog_Suppression_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "suppression"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = types.NewOrderedMap()
    rules.EntityData.Children.Append("rule", types.YChild{"Rule", nil})
    for i := range rules.Rule {
        rules.EntityData.Children.Append(types.GetSegmentPath(rules.Rule[i]), types.YChild{"Rule", rules.Rule[i]})
    }
    rules.EntityData.Leafs = types.NewOrderedMap()

    rules.EntityData.YListKeys = []string {}

    return &(rules.EntityData)
}

// Syslog_Suppression_Rules_Rule
// Rule name
type Syslog_Suppression_Rules_Rule struct {
    EntityData types.CommonEntityData
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

func (rule *Syslog_Suppression_Rules_Rule) GetEntityData() *types.CommonEntityData {
    rule.EntityData.YFilter = rule.YFilter
    rule.EntityData.YangName = "rule"
    rule.EntityData.BundleName = "cisco_ios_xr"
    rule.EntityData.ParentYangName = "rules"
    rule.EntityData.SegmentPath = "rule" + types.AddKeyToken(rule.Name, "name")
    rule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rule.EntityData.Children = types.NewOrderedMap()
    rule.EntityData.Children.Append("applied-to", types.YChild{"AppliedTo", &rule.AppliedTo})
    rule.EntityData.Children.Append("alarm-causes", types.YChild{"AlarmCauses", &rule.AlarmCauses})
    rule.EntityData.Leafs = types.NewOrderedMap()
    rule.EntityData.Leafs.Append("name", types.YLeaf{"Name", rule.Name})
    rule.EntityData.Leafs.Append("all-alarms", types.YLeaf{"AllAlarms", rule.AllAlarms})

    rule.EntityData.YListKeys = []string {"Name"}

    return &(rule.EntityData)
}

// Syslog_Suppression_Rules_Rule_AppliedTo
// Applied to the Rule
type Syslog_Suppression_Rules_Rule_AppliedTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Apply to all of the router. The type is interface{}.
    All interface{}

    // Table of configured sources to apply.
    Sources Syslog_Suppression_Rules_Rule_AppliedTo_Sources
}

func (appliedTo *Syslog_Suppression_Rules_Rule_AppliedTo) GetEntityData() *types.CommonEntityData {
    appliedTo.EntityData.YFilter = appliedTo.YFilter
    appliedTo.EntityData.YangName = "applied-to"
    appliedTo.EntityData.BundleName = "cisco_ios_xr"
    appliedTo.EntityData.ParentYangName = "rule"
    appliedTo.EntityData.SegmentPath = "applied-to"
    appliedTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appliedTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appliedTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appliedTo.EntityData.Children = types.NewOrderedMap()
    appliedTo.EntityData.Children.Append("sources", types.YChild{"Sources", &appliedTo.Sources})
    appliedTo.EntityData.Leafs = types.NewOrderedMap()
    appliedTo.EntityData.Leafs.Append("all", types.YLeaf{"All", appliedTo.All})

    appliedTo.EntityData.YListKeys = []string {}

    return &(appliedTo.EntityData)
}

// Syslog_Suppression_Rules_Rule_AppliedTo_Sources
// Table of configured sources to apply
type Syslog_Suppression_Rules_Rule_AppliedTo_Sources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An alarm source. The type is slice of
    // Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source.
    Source []*Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source
}

func (sources *Syslog_Suppression_Rules_Rule_AppliedTo_Sources) GetEntityData() *types.CommonEntityData {
    sources.EntityData.YFilter = sources.YFilter
    sources.EntityData.YangName = "sources"
    sources.EntityData.BundleName = "cisco_ios_xr"
    sources.EntityData.ParentYangName = "applied-to"
    sources.EntityData.SegmentPath = "sources"
    sources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sources.EntityData.Children = types.NewOrderedMap()
    sources.EntityData.Children.Append("source", types.YChild{"Source", nil})
    for i := range sources.Source {
        sources.EntityData.Children.Append(types.GetSegmentPath(sources.Source[i]), types.YChild{"Source", sources.Source[i]})
    }
    sources.EntityData.Leafs = types.NewOrderedMap()

    sources.EntityData.YListKeys = []string {}

    return &(sources.EntityData)
}

// Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source
// An alarm source
type Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Source. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Source interface{}
}

func (source *Syslog_Suppression_Rules_Rule_AppliedTo_Sources_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "sources"
    source.EntityData.SegmentPath = "source" + types.AddKeyToken(source.Source, "source")
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("source", types.YLeaf{"Source", source.Source})

    source.EntityData.YListKeys = []string {"Source"}

    return &(source.EntityData)
}

// Syslog_Suppression_Rules_Rule_AlarmCauses
// Causes of alarms to be suppressed
type Syslog_Suppression_Rules_Rule_AlarmCauses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Category, Group and Code of alarm/syslog to be suppressed. The type is
    // slice of Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause.
    AlarmCause []*Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause
}

func (alarmCauses *Syslog_Suppression_Rules_Rule_AlarmCauses) GetEntityData() *types.CommonEntityData {
    alarmCauses.EntityData.YFilter = alarmCauses.YFilter
    alarmCauses.EntityData.YangName = "alarm-causes"
    alarmCauses.EntityData.BundleName = "cisco_ios_xr"
    alarmCauses.EntityData.ParentYangName = "rule"
    alarmCauses.EntityData.SegmentPath = "alarm-causes"
    alarmCauses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmCauses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmCauses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmCauses.EntityData.Children = types.NewOrderedMap()
    alarmCauses.EntityData.Children.Append("alarm-cause", types.YChild{"AlarmCause", nil})
    for i := range alarmCauses.AlarmCause {
        alarmCauses.EntityData.Children.Append(types.GetSegmentPath(alarmCauses.AlarmCause[i]), types.YChild{"AlarmCause", alarmCauses.AlarmCause[i]})
    }
    alarmCauses.EntityData.Leafs = types.NewOrderedMap()

    alarmCauses.EntityData.YListKeys = []string {}

    return &(alarmCauses.EntityData)
}

// Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause
// Category, Group and Code of alarm/syslog to
// be suppressed
type Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Category. The type is string with length: 1..32.
    Category interface{}

    // This attribute is a key. Group. The type is string with length: 1..32.
    Group interface{}

    // This attribute is a key. Code. The type is string with length: 1..32.
    Code interface{}
}

func (alarmCause *Syslog_Suppression_Rules_Rule_AlarmCauses_AlarmCause) GetEntityData() *types.CommonEntityData {
    alarmCause.EntityData.YFilter = alarmCause.YFilter
    alarmCause.EntityData.YangName = "alarm-cause"
    alarmCause.EntityData.BundleName = "cisco_ios_xr"
    alarmCause.EntityData.ParentYangName = "alarm-causes"
    alarmCause.EntityData.SegmentPath = "alarm-cause" + types.AddKeyToken(alarmCause.Category, "category") + types.AddKeyToken(alarmCause.Group, "group") + types.AddKeyToken(alarmCause.Code, "code")
    alarmCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmCause.EntityData.Children = types.NewOrderedMap()
    alarmCause.EntityData.Leafs = types.NewOrderedMap()
    alarmCause.EntityData.Leafs.Append("category", types.YLeaf{"Category", alarmCause.Category})
    alarmCause.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmCause.Group})
    alarmCause.EntityData.Leafs.Append("code", types.YLeaf{"Code", alarmCause.Code})

    alarmCause.EntityData.YListKeys = []string {"Category", "Group", "Code"}

    return &(alarmCause.EntityData)
}

