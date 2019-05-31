// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-syslog package operational data.
// 
// This module contains definitions
// for the following management objects:
//   logging: Logging History data
//   syslog: syslog
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_syslog_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_syslog_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-syslog-oper get-syslog}", reflect.TypeOf(GetSyslog{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-syslog-oper:get-syslog", reflect.TypeOf(GetSyslog{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-syslog-oper logging}", reflect.TypeOf(Logging{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-syslog-oper:logging", reflect.TypeOf(Logging{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-syslog-oper syslog}", reflect.TypeOf(Syslog{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-syslog-oper:syslog", reflect.TypeOf(Syslog{}))
}

// SystemMessageSeverity represents System message severity
type SystemMessageSeverity string

const (
    // Unknown
    SystemMessageSeverity_message_severity_unknown SystemMessageSeverity = "message-severity-unknown"

    // Emergency
    SystemMessageSeverity_message_severity_emergency SystemMessageSeverity = "message-severity-emergency"

    // Alert
    SystemMessageSeverity_message_severity_alert SystemMessageSeverity = "message-severity-alert"

    // Critical
    SystemMessageSeverity_message_severity_critical SystemMessageSeverity = "message-severity-critical"

    // Error
    SystemMessageSeverity_message_severity_error SystemMessageSeverity = "message-severity-error"

    // Warning
    SystemMessageSeverity_message_severity_warning SystemMessageSeverity = "message-severity-warning"

    // Notice
    SystemMessageSeverity_message_severity_notice SystemMessageSeverity = "message-severity-notice"

    // Informational
    SystemMessageSeverity_message_severity_informational SystemMessageSeverity = "message-severity-informational"

    // Debug
    SystemMessageSeverity_message_severity_debug SystemMessageSeverity = "message-severity-debug"
)

// GetSyslog
type GetSyslog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input GetSyslog_Input

    
    Output GetSyslog_Output
}

func (getSyslog *GetSyslog) GetEntityData() *types.CommonEntityData {
    getSyslog.EntityData.YFilter = getSyslog.YFilter
    getSyslog.EntityData.YangName = "get-syslog"
    getSyslog.EntityData.BundleName = "cisco_ios_xr"
    getSyslog.EntityData.ParentYangName = "Cisco-IOS-XR-infra-syslog-oper"
    getSyslog.EntityData.SegmentPath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog"
    getSyslog.EntityData.AbsolutePath = getSyslog.EntityData.SegmentPath
    getSyslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    getSyslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    getSyslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    getSyslog.EntityData.Children = types.NewOrderedMap()
    getSyslog.EntityData.Children.Append("input", types.YChild{"Input", &getSyslog.Input})
    getSyslog.EntityData.Children.Append("output", types.YChild{"Output", &getSyslog.Output})
    getSyslog.EntityData.Leafs = types.NewOrderedMap()

    getSyslog.EntityData.YListKeys = []string {}

    return &(getSyslog.EntityData)
}

// GetSyslog_Input
type GetSyslog_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of filters of syslog messages to be returned.
    Filters GetSyslog_Input_Filters
}

func (input *GetSyslog_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "get-syslog"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("filters", types.YChild{"Filters", &input.Filters})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// GetSyslog_Input_Filters
// A set of filters of syslog messages to be returned
type GetSyslog_Input_Filters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start timestamp of syslog messages to be returned. The type is string with
    // pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    StartTime interface{}

    // End timestamp of syslog messages to be returned. The type is string with
    // pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    EndTime interface{}
}

func (filters *GetSyslog_Input_Filters) GetEntityData() *types.CommonEntityData {
    filters.EntityData.YFilter = filters.YFilter
    filters.EntityData.YangName = "filters"
    filters.EntityData.BundleName = "cisco_ios_xr"
    filters.EntityData.ParentYangName = "input"
    filters.EntityData.SegmentPath = "filters"
    filters.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog/input/" + filters.EntityData.SegmentPath
    filters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filters.EntityData.Children = types.NewOrderedMap()
    filters.EntityData.Leafs = types.NewOrderedMap()
    filters.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", filters.StartTime})
    filters.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", filters.EndTime})

    filters.EntityData.YListKeys = []string {}

    return &(filters.EntityData)
}

// GetSyslog_Output
type GetSyslog_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output for RPC operation.
    Data GetSyslog_Output_Data
}

func (output *GetSyslog_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "get-syslog"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("data", types.YChild{"Data", &output.Data})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// GetSyslog_Output_Data
// Output for RPC operation
type GetSyslog_Output_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // syslog.
    Syslog GetSyslog_Output_Data_Syslog
}

func (data *GetSyslog_Output_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "output"
    data.EntityData.SegmentPath = "data"
    data.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog/output/" + data.EntityData.SegmentPath
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = types.NewOrderedMap()
    data.EntityData.Children.Append("syslog", types.YChild{"Syslog", &data.Syslog})
    data.EntityData.Leafs = types.NewOrderedMap()

    data.EntityData.YListKeys = []string {}

    return &(data.EntityData)
}

// GetSyslog_Output_Data_Syslog
// syslog
type GetSyslog_Output_Data_Syslog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Message table information.
    Messages GetSyslog_Output_Data_Syslog_Messages
}

func (syslog *GetSyslog_Output_Data_Syslog) GetEntityData() *types.CommonEntityData {
    syslog.EntityData.YFilter = syslog.YFilter
    syslog.EntityData.YangName = "syslog"
    syslog.EntityData.BundleName = "cisco_ios_xr"
    syslog.EntityData.ParentYangName = "data"
    syslog.EntityData.SegmentPath = "syslog"
    syslog.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog/output/data/" + syslog.EntityData.SegmentPath
    syslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslog.EntityData.Children = types.NewOrderedMap()
    syslog.EntityData.Children.Append("messages", types.YChild{"Messages", &syslog.Messages})
    syslog.EntityData.Leafs = types.NewOrderedMap()

    syslog.EntityData.YListKeys = []string {}

    return &(syslog.EntityData)
}

// GetSyslog_Output_Data_Syslog_Messages
// Message table information
type GetSyslog_Output_Data_Syslog_Messages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A system message. The type is slice of
    // GetSyslog_Output_Data_Syslog_Messages_Message.
    Message []*GetSyslog_Output_Data_Syslog_Messages_Message
}

func (messages *GetSyslog_Output_Data_Syslog_Messages) GetEntityData() *types.CommonEntityData {
    messages.EntityData.YFilter = messages.YFilter
    messages.EntityData.YangName = "messages"
    messages.EntityData.BundleName = "cisco_ios_xr"
    messages.EntityData.ParentYangName = "syslog"
    messages.EntityData.SegmentPath = "messages"
    messages.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog/output/data/syslog/" + messages.EntityData.SegmentPath
    messages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    messages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    messages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    messages.EntityData.Children = types.NewOrderedMap()
    messages.EntityData.Children.Append("message", types.YChild{"Message", nil})
    for i := range messages.Message {
        messages.EntityData.Children.Append(types.GetSegmentPath(messages.Message[i]), types.YChild{"Message", messages.Message[i]})
    }
    messages.EntityData.Leafs = types.NewOrderedMap()

    messages.EntityData.YListKeys = []string {}

    return &(messages.EntityData)
}

// GetSyslog_Output_Data_Syslog_Messages_Message
// A system message
type GetSyslog_Output_Data_Syslog_Messages_Message struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Message ID of the system message. The type is
    // interface{} with range: 0..4294967295.
    MessageId interface{}

    // Message card location: 'RP', 'DRP', 'LC', 'SC', 'SP' or 'UNK' . The type is
    // string.
    CardType interface{}

    // Message source location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Time in milliseconds since 00:00:00 UTC, January 11970 of when message was
    // generated. The type is interface{} with range: 0..18446744073709551615.
    // Units are millisecond.
    TimeStamp interface{}

    // Time of day of event in DDD MMM DD  YYYY HH:MM :SS format, e.g Wed Apr 01
    // 2009  15:50:26. The type is string.
    TimeOfDay interface{}

    // Time Zone in UTC+/-HH:MM format,  e.g UTC+5:30, UTC-6. The type is string.
    TimeZone interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Message category. The type is string.
    Category interface{}

    // Message group. The type is string.
    Group interface{}

    // Message name. The type is string.
    MessageName interface{}

    // Message severity. The type is SystemMessageSeverity.
    Severity interface{}

    // Additional message text. The type is string.
    Text interface{}
}

func (message *GetSyslog_Output_Data_Syslog_Messages_Message) GetEntityData() *types.CommonEntityData {
    message.EntityData.YFilter = message.YFilter
    message.EntityData.YangName = "message"
    message.EntityData.BundleName = "cisco_ios_xr"
    message.EntityData.ParentYangName = "messages"
    message.EntityData.SegmentPath = "message" + types.AddKeyToken(message.MessageId, "message-id")
    message.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:get-syslog/output/data/syslog/messages/" + message.EntityData.SegmentPath
    message.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    message.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    message.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    message.EntityData.Children = types.NewOrderedMap()
    message.EntityData.Leafs = types.NewOrderedMap()
    message.EntityData.Leafs.Append("message-id", types.YLeaf{"MessageId", message.MessageId})
    message.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", message.CardType})
    message.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", message.NodeName})
    message.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", message.TimeStamp})
    message.EntityData.Leafs.Append("time-of-day", types.YLeaf{"TimeOfDay", message.TimeOfDay})
    message.EntityData.Leafs.Append("time-zone", types.YLeaf{"TimeZone", message.TimeZone})
    message.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", message.ProcessName})
    message.EntityData.Leafs.Append("category", types.YLeaf{"Category", message.Category})
    message.EntityData.Leafs.Append("group", types.YLeaf{"Group", message.Group})
    message.EntityData.Leafs.Append("message-name", types.YLeaf{"MessageName", message.MessageName})
    message.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", message.Severity})
    message.EntityData.Leafs.Append("text", types.YLeaf{"Text", message.Text})

    message.EntityData.YListKeys = []string {"MessageId"}

    return &(message.EntityData)
}

// Logging
// Logging History data
type Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Syslog Info .
    History Logging_History
}

func (logging *Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "Cisco-IOS-XR-infra-syslog-oper"
    logging.EntityData.SegmentPath = "Cisco-IOS-XR-infra-syslog-oper:logging"
    logging.EntityData.AbsolutePath = logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Children.Append("history", types.YChild{"History", &logging.History})
    logging.EntityData.Leafs = types.NewOrderedMap()

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Logging_History
// Syslog Info 
type Logging_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Syslog Properties. The type is string.
    Properties interface{}

    // Syslog Message. The type is string.
    Message interface{}
}

func (history *Logging_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "logging"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:logging/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("properties", types.YLeaf{"Properties", history.Properties})
    history.EntityData.Leafs.Append("message", types.YLeaf{"Message", history.Message})

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Syslog
// syslog
type Syslog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Files information.
    LoggingFiles Syslog_LoggingFiles

    // Logging AN remote servers information.
    AnRemoteServers Syslog_AnRemoteServers

    // Message table information.
    Messages Syslog_Messages

    // Logging statistics information.
    LoggingStatistics Syslog_LoggingStatistics
}

func (syslog *Syslog) GetEntityData() *types.CommonEntityData {
    syslog.EntityData.YFilter = syslog.YFilter
    syslog.EntityData.YangName = "syslog"
    syslog.EntityData.BundleName = "cisco_ios_xr"
    syslog.EntityData.ParentYangName = "Cisco-IOS-XR-infra-syslog-oper"
    syslog.EntityData.SegmentPath = "Cisco-IOS-XR-infra-syslog-oper:syslog"
    syslog.EntityData.AbsolutePath = syslog.EntityData.SegmentPath
    syslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslog.EntityData.Children = types.NewOrderedMap()
    syslog.EntityData.Children.Append("logging-files", types.YChild{"LoggingFiles", &syslog.LoggingFiles})
    syslog.EntityData.Children.Append("an-remote-servers", types.YChild{"AnRemoteServers", &syslog.AnRemoteServers})
    syslog.EntityData.Children.Append("messages", types.YChild{"Messages", &syslog.Messages})
    syslog.EntityData.Children.Append("logging-statistics", types.YChild{"LoggingStatistics", &syslog.LoggingStatistics})
    syslog.EntityData.Leafs = types.NewOrderedMap()

    syslog.EntityData.YListKeys = []string {}

    return &(syslog.EntityData)
}

// Syslog_LoggingFiles
// Logging Files information
type Syslog_LoggingFiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Files. The type is slice of Syslog_LoggingFiles_FileLogDetail.
    FileLogDetail []*Syslog_LoggingFiles_FileLogDetail
}

func (loggingFiles *Syslog_LoggingFiles) GetEntityData() *types.CommonEntityData {
    loggingFiles.EntityData.YFilter = loggingFiles.YFilter
    loggingFiles.EntityData.YangName = "logging-files"
    loggingFiles.EntityData.BundleName = "cisco_ios_xr"
    loggingFiles.EntityData.ParentYangName = "syslog"
    loggingFiles.EntityData.SegmentPath = "logging-files"
    loggingFiles.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/" + loggingFiles.EntityData.SegmentPath
    loggingFiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loggingFiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loggingFiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loggingFiles.EntityData.Children = types.NewOrderedMap()
    loggingFiles.EntityData.Children.Append("file-log-detail", types.YChild{"FileLogDetail", nil})
    for i := range loggingFiles.FileLogDetail {
        types.SetYListKey(loggingFiles.FileLogDetail[i], i)
        loggingFiles.EntityData.Children.Append(types.GetSegmentPath(loggingFiles.FileLogDetail[i]), types.YChild{"FileLogDetail", loggingFiles.FileLogDetail[i]})
    }
    loggingFiles.EntityData.Leafs = types.NewOrderedMap()

    loggingFiles.EntityData.YListKeys = []string {}

    return &(loggingFiles.EntityData)
}

// Syslog_LoggingFiles_FileLogDetail
// Logging Files
type Syslog_LoggingFiles_FileLogDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // File path for logging messages. The type is string.
    FilePath interface{}

    // File name for logging messages. The type is string.
    FileName interface{}
}

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetEntityData() *types.CommonEntityData {
    fileLogDetail.EntityData.YFilter = fileLogDetail.YFilter
    fileLogDetail.EntityData.YangName = "file-log-detail"
    fileLogDetail.EntityData.BundleName = "cisco_ios_xr"
    fileLogDetail.EntityData.ParentYangName = "logging-files"
    fileLogDetail.EntityData.SegmentPath = "file-log-detail" + types.AddNoKeyToken(fileLogDetail)
    fileLogDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-files/" + fileLogDetail.EntityData.SegmentPath
    fileLogDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileLogDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileLogDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileLogDetail.EntityData.Children = types.NewOrderedMap()
    fileLogDetail.EntityData.Leafs = types.NewOrderedMap()
    fileLogDetail.EntityData.Leafs.Append("file-path", types.YLeaf{"FilePath", fileLogDetail.FilePath})
    fileLogDetail.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", fileLogDetail.FileName})

    fileLogDetail.EntityData.YListKeys = []string {}

    return &(fileLogDetail.EntityData)
}

// Syslog_AnRemoteServers
// Logging AN remote servers information
type Syslog_AnRemoteServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AN Remote Log Servers. The type is slice of
    // Syslog_AnRemoteServers_AnRemoteLogServer.
    AnRemoteLogServer []*Syslog_AnRemoteServers_AnRemoteLogServer
}

func (anRemoteServers *Syslog_AnRemoteServers) GetEntityData() *types.CommonEntityData {
    anRemoteServers.EntityData.YFilter = anRemoteServers.YFilter
    anRemoteServers.EntityData.YangName = "an-remote-servers"
    anRemoteServers.EntityData.BundleName = "cisco_ios_xr"
    anRemoteServers.EntityData.ParentYangName = "syslog"
    anRemoteServers.EntityData.SegmentPath = "an-remote-servers"
    anRemoteServers.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/" + anRemoteServers.EntityData.SegmentPath
    anRemoteServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    anRemoteServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    anRemoteServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    anRemoteServers.EntityData.Children = types.NewOrderedMap()
    anRemoteServers.EntityData.Children.Append("an-remote-log-server", types.YChild{"AnRemoteLogServer", nil})
    for i := range anRemoteServers.AnRemoteLogServer {
        types.SetYListKey(anRemoteServers.AnRemoteLogServer[i], i)
        anRemoteServers.EntityData.Children.Append(types.GetSegmentPath(anRemoteServers.AnRemoteLogServer[i]), types.YChild{"AnRemoteLogServer", anRemoteServers.AnRemoteLogServer[i]})
    }
    anRemoteServers.EntityData.Leafs = types.NewOrderedMap()

    anRemoteServers.EntityData.YListKeys = []string {}

    return &(anRemoteServers.EntityData)
}

// Syslog_AnRemoteServers_AnRemoteLogServer
// AN Remote Log Servers
type Syslog_AnRemoteServers_AnRemoteLogServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IP Address. The type is string.
    IpAddress interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Severity. The type is string.
    VrfSeverity interface{}

    // Remote-Host Discriminator. The type is string.
    RhDiscriminator interface{}
}

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetEntityData() *types.CommonEntityData {
    anRemoteLogServer.EntityData.YFilter = anRemoteLogServer.YFilter
    anRemoteLogServer.EntityData.YangName = "an-remote-log-server"
    anRemoteLogServer.EntityData.BundleName = "cisco_ios_xr"
    anRemoteLogServer.EntityData.ParentYangName = "an-remote-servers"
    anRemoteLogServer.EntityData.SegmentPath = "an-remote-log-server" + types.AddNoKeyToken(anRemoteLogServer)
    anRemoteLogServer.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/an-remote-servers/" + anRemoteLogServer.EntityData.SegmentPath
    anRemoteLogServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    anRemoteLogServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    anRemoteLogServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    anRemoteLogServer.EntityData.Children = types.NewOrderedMap()
    anRemoteLogServer.EntityData.Leafs = types.NewOrderedMap()
    anRemoteLogServer.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", anRemoteLogServer.IpAddress})
    anRemoteLogServer.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", anRemoteLogServer.VrfName})
    anRemoteLogServer.EntityData.Leafs.Append("vrf-severity", types.YLeaf{"VrfSeverity", anRemoteLogServer.VrfSeverity})
    anRemoteLogServer.EntityData.Leafs.Append("rh-discriminator", types.YLeaf{"RhDiscriminator", anRemoteLogServer.RhDiscriminator})

    anRemoteLogServer.EntityData.YListKeys = []string {}

    return &(anRemoteLogServer.EntityData)
}

// Syslog_Messages
// Message table information
type Syslog_Messages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A system message. The type is slice of Syslog_Messages_Message.
    Message []*Syslog_Messages_Message
}

func (messages *Syslog_Messages) GetEntityData() *types.CommonEntityData {
    messages.EntityData.YFilter = messages.YFilter
    messages.EntityData.YangName = "messages"
    messages.EntityData.BundleName = "cisco_ios_xr"
    messages.EntityData.ParentYangName = "syslog"
    messages.EntityData.SegmentPath = "messages"
    messages.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/" + messages.EntityData.SegmentPath
    messages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    messages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    messages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    messages.EntityData.Children = types.NewOrderedMap()
    messages.EntityData.Children.Append("message", types.YChild{"Message", nil})
    for i := range messages.Message {
        messages.EntityData.Children.Append(types.GetSegmentPath(messages.Message[i]), types.YChild{"Message", messages.Message[i]})
    }
    messages.EntityData.Leafs = types.NewOrderedMap()

    messages.EntityData.YListKeys = []string {}

    return &(messages.EntityData)
}

// Syslog_Messages_Message
// A system message
type Syslog_Messages_Message struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Message ID of the system message. The type is
    // interface{} with range: 0..4294967295.
    MessageId interface{}

    // Message card location: 'RP', 'DRP', 'LC', 'SC', 'SP' or 'UNK' . The type is
    // string.
    CardType interface{}

    // Message source location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Time in milliseconds since 00:00:00 UTC, January 11970 of when message was
    // generated. The type is interface{} with range: 0..18446744073709551615.
    // Units are millisecond.
    TimeStamp interface{}

    // Time of day of event in DDD MMM DD  YYYY HH:MM :SS format, e.g Wed Apr 01
    // 2009  15:50:26. The type is string.
    TimeOfDay interface{}

    // Time Zone in UTC+/-HH:MM format,  e.g UTC+5:30, UTC-6. The type is string.
    TimeZone interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Message category. The type is string.
    Category interface{}

    // Message group. The type is string.
    Group interface{}

    // Message name. The type is string.
    MessageName interface{}

    // Message severity. The type is SystemMessageSeverity.
    Severity interface{}

    // Additional message text. The type is string.
    Text interface{}
}

func (message *Syslog_Messages_Message) GetEntityData() *types.CommonEntityData {
    message.EntityData.YFilter = message.YFilter
    message.EntityData.YangName = "message"
    message.EntityData.BundleName = "cisco_ios_xr"
    message.EntityData.ParentYangName = "messages"
    message.EntityData.SegmentPath = "message" + types.AddKeyToken(message.MessageId, "message-id")
    message.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/messages/" + message.EntityData.SegmentPath
    message.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    message.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    message.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    message.EntityData.Children = types.NewOrderedMap()
    message.EntityData.Leafs = types.NewOrderedMap()
    message.EntityData.Leafs.Append("message-id", types.YLeaf{"MessageId", message.MessageId})
    message.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", message.CardType})
    message.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", message.NodeName})
    message.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", message.TimeStamp})
    message.EntityData.Leafs.Append("time-of-day", types.YLeaf{"TimeOfDay", message.TimeOfDay})
    message.EntityData.Leafs.Append("time-zone", types.YLeaf{"TimeZone", message.TimeZone})
    message.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", message.ProcessName})
    message.EntityData.Leafs.Append("category", types.YLeaf{"Category", message.Category})
    message.EntityData.Leafs.Append("group", types.YLeaf{"Group", message.Group})
    message.EntityData.Leafs.Append("message-name", types.YLeaf{"MessageName", message.MessageName})
    message.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", message.Severity})
    message.EntityData.Leafs.Append("text", types.YLeaf{"Text", message.Text})

    message.EntityData.YListKeys = []string {"MessageId"}

    return &(message.EntityData)
}

// Syslog_LoggingStatistics
// Logging statistics information
type Syslog_LoggingStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging statistics.
    LoggingStats Syslog_LoggingStatistics_LoggingStats

    // Console logging statistics.
    ConsoleLoggingStats Syslog_LoggingStatistics_ConsoleLoggingStats

    // Monitor loggingstatistics.
    MonitorLoggingStats Syslog_LoggingStatistics_MonitorLoggingStats

    // Trap logging statistics.
    TrapLoggingStats Syslog_LoggingStatistics_TrapLoggingStats

    // Buffer logging statistics.
    BufferLoggingStats Syslog_LoggingStatistics_BufferLoggingStats

    // Remote logging statistics. The type is slice of
    // Syslog_LoggingStatistics_RemoteLoggingStat.
    RemoteLoggingStat []*Syslog_LoggingStatistics_RemoteLoggingStat

    // TLS Remote logging statistics. The type is slice of
    // Syslog_LoggingStatistics_TlsRemoteLoggingStat.
    TlsRemoteLoggingStat []*Syslog_LoggingStatistics_TlsRemoteLoggingStat

    // File logging statistics. The type is slice of
    // Syslog_LoggingStatistics_FileLoggingStat.
    FileLoggingStat []*Syslog_LoggingStatistics_FileLoggingStat
}

func (loggingStatistics *Syslog_LoggingStatistics) GetEntityData() *types.CommonEntityData {
    loggingStatistics.EntityData.YFilter = loggingStatistics.YFilter
    loggingStatistics.EntityData.YangName = "logging-statistics"
    loggingStatistics.EntityData.BundleName = "cisco_ios_xr"
    loggingStatistics.EntityData.ParentYangName = "syslog"
    loggingStatistics.EntityData.SegmentPath = "logging-statistics"
    loggingStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/" + loggingStatistics.EntityData.SegmentPath
    loggingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loggingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loggingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loggingStatistics.EntityData.Children = types.NewOrderedMap()
    loggingStatistics.EntityData.Children.Append("logging-stats", types.YChild{"LoggingStats", &loggingStatistics.LoggingStats})
    loggingStatistics.EntityData.Children.Append("console-logging-stats", types.YChild{"ConsoleLoggingStats", &loggingStatistics.ConsoleLoggingStats})
    loggingStatistics.EntityData.Children.Append("monitor-logging-stats", types.YChild{"MonitorLoggingStats", &loggingStatistics.MonitorLoggingStats})
    loggingStatistics.EntityData.Children.Append("trap-logging-stats", types.YChild{"TrapLoggingStats", &loggingStatistics.TrapLoggingStats})
    loggingStatistics.EntityData.Children.Append("buffer-logging-stats", types.YChild{"BufferLoggingStats", &loggingStatistics.BufferLoggingStats})
    loggingStatistics.EntityData.Children.Append("remote-logging-stat", types.YChild{"RemoteLoggingStat", nil})
    for i := range loggingStatistics.RemoteLoggingStat {
        types.SetYListKey(loggingStatistics.RemoteLoggingStat[i], i)
        loggingStatistics.EntityData.Children.Append(types.GetSegmentPath(loggingStatistics.RemoteLoggingStat[i]), types.YChild{"RemoteLoggingStat", loggingStatistics.RemoteLoggingStat[i]})
    }
    loggingStatistics.EntityData.Children.Append("tls-remote-logging-stat", types.YChild{"TlsRemoteLoggingStat", nil})
    for i := range loggingStatistics.TlsRemoteLoggingStat {
        types.SetYListKey(loggingStatistics.TlsRemoteLoggingStat[i], i)
        loggingStatistics.EntityData.Children.Append(types.GetSegmentPath(loggingStatistics.TlsRemoteLoggingStat[i]), types.YChild{"TlsRemoteLoggingStat", loggingStatistics.TlsRemoteLoggingStat[i]})
    }
    loggingStatistics.EntityData.Children.Append("file-logging-stat", types.YChild{"FileLoggingStat", nil})
    for i := range loggingStatistics.FileLoggingStat {
        types.SetYListKey(loggingStatistics.FileLoggingStat[i], i)
        loggingStatistics.EntityData.Children.Append(types.GetSegmentPath(loggingStatistics.FileLoggingStat[i]), types.YChild{"FileLoggingStat", loggingStatistics.FileLoggingStat[i]})
    }
    loggingStatistics.EntityData.Leafs = types.NewOrderedMap()

    loggingStatistics.EntityData.YListKeys = []string {}

    return &(loggingStatistics.EntityData)
}

// Syslog_LoggingStatistics_LoggingStats
// Logging statistics
type Syslog_LoggingStatistics_LoggingStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is log enabled. The type is bool.
    IsLogEnabled interface{}

    // Number of messages dropped. The type is interface{} with range:
    // 0..4294967295.
    DropCount interface{}

    // Number of messages flushed. The type is interface{} with range:
    // 0..4294967295.
    FlushCount interface{}

    // Number of messages overrun. The type is interface{} with range:
    // 0..4294967295.
    OverrunCount interface{}
}

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetEntityData() *types.CommonEntityData {
    loggingStats.EntityData.YFilter = loggingStats.YFilter
    loggingStats.EntityData.YangName = "logging-stats"
    loggingStats.EntityData.BundleName = "cisco_ios_xr"
    loggingStats.EntityData.ParentYangName = "logging-statistics"
    loggingStats.EntityData.SegmentPath = "logging-stats"
    loggingStats.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + loggingStats.EntityData.SegmentPath
    loggingStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loggingStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loggingStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loggingStats.EntityData.Children = types.NewOrderedMap()
    loggingStats.EntityData.Leafs = types.NewOrderedMap()
    loggingStats.EntityData.Leafs.Append("is-log-enabled", types.YLeaf{"IsLogEnabled", loggingStats.IsLogEnabled})
    loggingStats.EntityData.Leafs.Append("drop-count", types.YLeaf{"DropCount", loggingStats.DropCount})
    loggingStats.EntityData.Leafs.Append("flush-count", types.YLeaf{"FlushCount", loggingStats.FlushCount})
    loggingStats.EntityData.Leafs.Append("overrun-count", types.YLeaf{"OverrunCount", loggingStats.OverrunCount})

    loggingStats.EntityData.YListKeys = []string {}

    return &(loggingStats.EntityData)
}

// Syslog_LoggingStatistics_ConsoleLoggingStats
// Console logging statistics
type Syslog_LoggingStatistics_ConsoleLoggingStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is log enabled. The type is bool.
    IsLogEnabled interface{}

    // Configured severity. The type is SystemMessageSeverity.
    Severity interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}

    // Buffer size in bytes if logging buffer isenabled. The type is interface{}
    // with range: 0..4294967295. Units are byte.
    BufferSize interface{}
}

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetEntityData() *types.CommonEntityData {
    consoleLoggingStats.EntityData.YFilter = consoleLoggingStats.YFilter
    consoleLoggingStats.EntityData.YangName = "console-logging-stats"
    consoleLoggingStats.EntityData.BundleName = "cisco_ios_xr"
    consoleLoggingStats.EntityData.ParentYangName = "logging-statistics"
    consoleLoggingStats.EntityData.SegmentPath = "console-logging-stats"
    consoleLoggingStats.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + consoleLoggingStats.EntityData.SegmentPath
    consoleLoggingStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    consoleLoggingStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    consoleLoggingStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    consoleLoggingStats.EntityData.Children = types.NewOrderedMap()
    consoleLoggingStats.EntityData.Leafs = types.NewOrderedMap()
    consoleLoggingStats.EntityData.Leafs.Append("is-log-enabled", types.YLeaf{"IsLogEnabled", consoleLoggingStats.IsLogEnabled})
    consoleLoggingStats.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", consoleLoggingStats.Severity})
    consoleLoggingStats.EntityData.Leafs.Append("message-count", types.YLeaf{"MessageCount", consoleLoggingStats.MessageCount})
    consoleLoggingStats.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", consoleLoggingStats.BufferSize})

    consoleLoggingStats.EntityData.YListKeys = []string {}

    return &(consoleLoggingStats.EntityData)
}

// Syslog_LoggingStatistics_MonitorLoggingStats
// Monitor loggingstatistics
type Syslog_LoggingStatistics_MonitorLoggingStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is log enabled. The type is bool.
    IsLogEnabled interface{}

    // Configured severity. The type is SystemMessageSeverity.
    Severity interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}

    // Buffer size in bytes if logging buffer isenabled. The type is interface{}
    // with range: 0..4294967295. Units are byte.
    BufferSize interface{}
}

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetEntityData() *types.CommonEntityData {
    monitorLoggingStats.EntityData.YFilter = monitorLoggingStats.YFilter
    monitorLoggingStats.EntityData.YangName = "monitor-logging-stats"
    monitorLoggingStats.EntityData.BundleName = "cisco_ios_xr"
    monitorLoggingStats.EntityData.ParentYangName = "logging-statistics"
    monitorLoggingStats.EntityData.SegmentPath = "monitor-logging-stats"
    monitorLoggingStats.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + monitorLoggingStats.EntityData.SegmentPath
    monitorLoggingStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorLoggingStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorLoggingStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorLoggingStats.EntityData.Children = types.NewOrderedMap()
    monitorLoggingStats.EntityData.Leafs = types.NewOrderedMap()
    monitorLoggingStats.EntityData.Leafs.Append("is-log-enabled", types.YLeaf{"IsLogEnabled", monitorLoggingStats.IsLogEnabled})
    monitorLoggingStats.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", monitorLoggingStats.Severity})
    monitorLoggingStats.EntityData.Leafs.Append("message-count", types.YLeaf{"MessageCount", monitorLoggingStats.MessageCount})
    monitorLoggingStats.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", monitorLoggingStats.BufferSize})

    monitorLoggingStats.EntityData.YListKeys = []string {}

    return &(monitorLoggingStats.EntityData)
}

// Syslog_LoggingStatistics_TrapLoggingStats
// Trap logging statistics
type Syslog_LoggingStatistics_TrapLoggingStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is log enabled. The type is bool.
    IsLogEnabled interface{}

    // Configured severity. The type is SystemMessageSeverity.
    Severity interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}

    // Buffer size in bytes if logging buffer isenabled. The type is interface{}
    // with range: 0..4294967295. Units are byte.
    BufferSize interface{}
}

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetEntityData() *types.CommonEntityData {
    trapLoggingStats.EntityData.YFilter = trapLoggingStats.YFilter
    trapLoggingStats.EntityData.YangName = "trap-logging-stats"
    trapLoggingStats.EntityData.BundleName = "cisco_ios_xr"
    trapLoggingStats.EntityData.ParentYangName = "logging-statistics"
    trapLoggingStats.EntityData.SegmentPath = "trap-logging-stats"
    trapLoggingStats.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + trapLoggingStats.EntityData.SegmentPath
    trapLoggingStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapLoggingStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapLoggingStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapLoggingStats.EntityData.Children = types.NewOrderedMap()
    trapLoggingStats.EntityData.Leafs = types.NewOrderedMap()
    trapLoggingStats.EntityData.Leafs.Append("is-log-enabled", types.YLeaf{"IsLogEnabled", trapLoggingStats.IsLogEnabled})
    trapLoggingStats.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", trapLoggingStats.Severity})
    trapLoggingStats.EntityData.Leafs.Append("message-count", types.YLeaf{"MessageCount", trapLoggingStats.MessageCount})
    trapLoggingStats.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", trapLoggingStats.BufferSize})

    trapLoggingStats.EntityData.YListKeys = []string {}

    return &(trapLoggingStats.EntityData)
}

// Syslog_LoggingStatistics_BufferLoggingStats
// Buffer logging statistics
type Syslog_LoggingStatistics_BufferLoggingStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is log enabled. The type is bool.
    IsLogEnabled interface{}

    // Configured severity. The type is SystemMessageSeverity.
    Severity interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}

    // Buffer size in bytes if logging buffer isenabled. The type is interface{}
    // with range: 0..4294967295. Units are byte.
    BufferSize interface{}
}

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetEntityData() *types.CommonEntityData {
    bufferLoggingStats.EntityData.YFilter = bufferLoggingStats.YFilter
    bufferLoggingStats.EntityData.YangName = "buffer-logging-stats"
    bufferLoggingStats.EntityData.BundleName = "cisco_ios_xr"
    bufferLoggingStats.EntityData.ParentYangName = "logging-statistics"
    bufferLoggingStats.EntityData.SegmentPath = "buffer-logging-stats"
    bufferLoggingStats.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + bufferLoggingStats.EntityData.SegmentPath
    bufferLoggingStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bufferLoggingStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bufferLoggingStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bufferLoggingStats.EntityData.Children = types.NewOrderedMap()
    bufferLoggingStats.EntityData.Leafs = types.NewOrderedMap()
    bufferLoggingStats.EntityData.Leafs.Append("is-log-enabled", types.YLeaf{"IsLogEnabled", bufferLoggingStats.IsLogEnabled})
    bufferLoggingStats.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", bufferLoggingStats.Severity})
    bufferLoggingStats.EntityData.Leafs.Append("message-count", types.YLeaf{"MessageCount", bufferLoggingStats.MessageCount})
    bufferLoggingStats.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", bufferLoggingStats.BufferSize})

    bufferLoggingStats.EntityData.YListKeys = []string {}

    return &(bufferLoggingStats.EntityData)
}

// Syslog_LoggingStatistics_RemoteLoggingStat
// Remote logging statistics
type Syslog_LoggingStatistics_RemoteLoggingStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Remote hostname. The type is string.
    RemoteHostName interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}
}

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetEntityData() *types.CommonEntityData {
    remoteLoggingStat.EntityData.YFilter = remoteLoggingStat.YFilter
    remoteLoggingStat.EntityData.YangName = "remote-logging-stat"
    remoteLoggingStat.EntityData.BundleName = "cisco_ios_xr"
    remoteLoggingStat.EntityData.ParentYangName = "logging-statistics"
    remoteLoggingStat.EntityData.SegmentPath = "remote-logging-stat" + types.AddNoKeyToken(remoteLoggingStat)
    remoteLoggingStat.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + remoteLoggingStat.EntityData.SegmentPath
    remoteLoggingStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLoggingStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLoggingStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLoggingStat.EntityData.Children = types.NewOrderedMap()
    remoteLoggingStat.EntityData.Leafs = types.NewOrderedMap()
    remoteLoggingStat.EntityData.Leafs.Append("remote-host-name", types.YLeaf{"RemoteHostName", remoteLoggingStat.RemoteHostName})
    remoteLoggingStat.EntityData.Leafs.Append("message-count", types.YLeaf{"MessageCount", remoteLoggingStat.MessageCount})

    remoteLoggingStat.EntityData.YListKeys = []string {}

    return &(remoteLoggingStat.EntityData)
}

// Syslog_LoggingStatistics_TlsRemoteLoggingStat
// TLS Remote logging statistics
type Syslog_LoggingStatistics_TlsRemoteLoggingStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // TLS Remote hostname. The type is string.
    RemoteHostName interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}
}

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetEntityData() *types.CommonEntityData {
    tlsRemoteLoggingStat.EntityData.YFilter = tlsRemoteLoggingStat.YFilter
    tlsRemoteLoggingStat.EntityData.YangName = "tls-remote-logging-stat"
    tlsRemoteLoggingStat.EntityData.BundleName = "cisco_ios_xr"
    tlsRemoteLoggingStat.EntityData.ParentYangName = "logging-statistics"
    tlsRemoteLoggingStat.EntityData.SegmentPath = "tls-remote-logging-stat" + types.AddNoKeyToken(tlsRemoteLoggingStat)
    tlsRemoteLoggingStat.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + tlsRemoteLoggingStat.EntityData.SegmentPath
    tlsRemoteLoggingStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tlsRemoteLoggingStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tlsRemoteLoggingStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tlsRemoteLoggingStat.EntityData.Children = types.NewOrderedMap()
    tlsRemoteLoggingStat.EntityData.Leafs = types.NewOrderedMap()
    tlsRemoteLoggingStat.EntityData.Leafs.Append("remote-host-name", types.YLeaf{"RemoteHostName", tlsRemoteLoggingStat.RemoteHostName})
    tlsRemoteLoggingStat.EntityData.Leafs.Append("message-count", types.YLeaf{"MessageCount", tlsRemoteLoggingStat.MessageCount})

    tlsRemoteLoggingStat.EntityData.YListKeys = []string {}

    return &(tlsRemoteLoggingStat.EntityData)
}

// Syslog_LoggingStatistics_FileLoggingStat
// File logging statistics
type Syslog_LoggingStatistics_FileLoggingStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // File name for logging messages. The type is string.
    FileName interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}
}

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetEntityData() *types.CommonEntityData {
    fileLoggingStat.EntityData.YFilter = fileLoggingStat.YFilter
    fileLoggingStat.EntityData.YangName = "file-logging-stat"
    fileLoggingStat.EntityData.BundleName = "cisco_ios_xr"
    fileLoggingStat.EntityData.ParentYangName = "logging-statistics"
    fileLoggingStat.EntityData.SegmentPath = "file-logging-stat" + types.AddNoKeyToken(fileLoggingStat)
    fileLoggingStat.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-syslog-oper:syslog/logging-statistics/" + fileLoggingStat.EntityData.SegmentPath
    fileLoggingStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileLoggingStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileLoggingStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileLoggingStat.EntityData.Children = types.NewOrderedMap()
    fileLoggingStat.EntityData.Leafs = types.NewOrderedMap()
    fileLoggingStat.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", fileLoggingStat.FileName})
    fileLoggingStat.EntityData.Leafs.Append("message-count", types.YLeaf{"MessageCount", fileLoggingStat.MessageCount})

    fileLoggingStat.EntityData.YListKeys = []string {}

    return &(fileLoggingStat.EntityData)
}

