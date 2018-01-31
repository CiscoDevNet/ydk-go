// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-syslog package operational data.
// 
// This module contains definitions
// for the following management objects:
//   logging: Logging History data
//   syslog: syslog
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// Logging
// Logging History data
type Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Syslog Info .
    History Logging_History
}

func (logging *Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Logging) GetGoName(yname string) string {
    if yname == "history" { return "History" }
    return ""
}

func (logging *Logging) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-syslog-oper:logging"
}

func (logging *Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "history" {
        return &logging.History
    }
    return nil
}

func (logging *Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["history"] = &logging.History
    return children
}

func (logging *Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logging *Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Logging) GetYangName() string { return "logging" }

func (logging *Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Logging) GetParent() types.Entity { return logging.parent }

func (logging *Logging) GetParentYangName() string { return "Cisco-IOS-XR-infra-syslog-oper" }

// Logging_History
// Syslog Info 
type Logging_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Syslog Properties. The type is string.
    Properties interface{}

    // Syslog Message. The type is string.
    Message interface{}
}

func (history *Logging_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Logging_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Logging_History) GetGoName(yname string) string {
    if yname == "properties" { return "Properties" }
    if yname == "message" { return "Message" }
    return ""
}

func (history *Logging_History) GetSegmentPath() string {
    return "history"
}

func (history *Logging_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (history *Logging_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (history *Logging_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["properties"] = history.Properties
    leafs["message"] = history.Message
    return leafs
}

func (history *Logging_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Logging_History) GetYangName() string { return "history" }

func (history *Logging_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Logging_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Logging_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Logging_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Logging_History) GetParent() types.Entity { return history.parent }

func (history *Logging_History) GetParentYangName() string { return "logging" }

// Syslog
// syslog
type Syslog struct {
    parent types.Entity
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

func (syslog *Syslog) GetFilter() yfilter.YFilter { return syslog.YFilter }

func (syslog *Syslog) SetFilter(yf yfilter.YFilter) { syslog.YFilter = yf }

func (syslog *Syslog) GetGoName(yname string) string {
    if yname == "logging-files" { return "LoggingFiles" }
    if yname == "an-remote-servers" { return "AnRemoteServers" }
    if yname == "messages" { return "Messages" }
    if yname == "logging-statistics" { return "LoggingStatistics" }
    return ""
}

func (syslog *Syslog) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-syslog-oper:syslog"
}

func (syslog *Syslog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logging-files" {
        return &syslog.LoggingFiles
    }
    if childYangName == "an-remote-servers" {
        return &syslog.AnRemoteServers
    }
    if childYangName == "messages" {
        return &syslog.Messages
    }
    if childYangName == "logging-statistics" {
        return &syslog.LoggingStatistics
    }
    return nil
}

func (syslog *Syslog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["logging-files"] = &syslog.LoggingFiles
    children["an-remote-servers"] = &syslog.AnRemoteServers
    children["messages"] = &syslog.Messages
    children["logging-statistics"] = &syslog.LoggingStatistics
    return children
}

func (syslog *Syslog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
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

func (syslog *Syslog) GetParentYangName() string { return "Cisco-IOS-XR-infra-syslog-oper" }

// Syslog_LoggingFiles
// Logging Files information
type Syslog_LoggingFiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging Files. The type is slice of Syslog_LoggingFiles_FileLogDetail.
    FileLogDetail []Syslog_LoggingFiles_FileLogDetail
}

func (loggingFiles *Syslog_LoggingFiles) GetFilter() yfilter.YFilter { return loggingFiles.YFilter }

func (loggingFiles *Syslog_LoggingFiles) SetFilter(yf yfilter.YFilter) { loggingFiles.YFilter = yf }

func (loggingFiles *Syslog_LoggingFiles) GetGoName(yname string) string {
    if yname == "file-log-detail" { return "FileLogDetail" }
    return ""
}

func (loggingFiles *Syslog_LoggingFiles) GetSegmentPath() string {
    return "logging-files"
}

func (loggingFiles *Syslog_LoggingFiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "file-log-detail" {
        for _, c := range loggingFiles.FileLogDetail {
            if loggingFiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_LoggingFiles_FileLogDetail{}
        loggingFiles.FileLogDetail = append(loggingFiles.FileLogDetail, child)
        return &loggingFiles.FileLogDetail[len(loggingFiles.FileLogDetail)-1]
    }
    return nil
}

func (loggingFiles *Syslog_LoggingFiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range loggingFiles.FileLogDetail {
        children[loggingFiles.FileLogDetail[i].GetSegmentPath()] = &loggingFiles.FileLogDetail[i]
    }
    return children
}

func (loggingFiles *Syslog_LoggingFiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loggingFiles *Syslog_LoggingFiles) GetBundleName() string { return "cisco_ios_xr" }

func (loggingFiles *Syslog_LoggingFiles) GetYangName() string { return "logging-files" }

func (loggingFiles *Syslog_LoggingFiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loggingFiles *Syslog_LoggingFiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loggingFiles *Syslog_LoggingFiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loggingFiles *Syslog_LoggingFiles) SetParent(parent types.Entity) { loggingFiles.parent = parent }

func (loggingFiles *Syslog_LoggingFiles) GetParent() types.Entity { return loggingFiles.parent }

func (loggingFiles *Syslog_LoggingFiles) GetParentYangName() string { return "syslog" }

// Syslog_LoggingFiles_FileLogDetail
// Logging Files
type Syslog_LoggingFiles_FileLogDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // File path for logging messages. The type is string.
    FilePath interface{}

    // File name for logging messages. The type is string.
    FileName interface{}
}

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetFilter() yfilter.YFilter { return fileLogDetail.YFilter }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) SetFilter(yf yfilter.YFilter) { fileLogDetail.YFilter = yf }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetGoName(yname string) string {
    if yname == "file-path" { return "FilePath" }
    if yname == "file-name" { return "FileName" }
    return ""
}

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetSegmentPath() string {
    return "file-log-detail"
}

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["file-path"] = fileLogDetail.FilePath
    leafs["file-name"] = fileLogDetail.FileName
    return leafs
}

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetBundleName() string { return "cisco_ios_xr" }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetYangName() string { return "file-log-detail" }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) SetParent(parent types.Entity) { fileLogDetail.parent = parent }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetParent() types.Entity { return fileLogDetail.parent }

func (fileLogDetail *Syslog_LoggingFiles_FileLogDetail) GetParentYangName() string { return "logging-files" }

// Syslog_AnRemoteServers
// Logging AN remote servers information
type Syslog_AnRemoteServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AN Remote Log Servers. The type is slice of
    // Syslog_AnRemoteServers_AnRemoteLogServer.
    AnRemoteLogServer []Syslog_AnRemoteServers_AnRemoteLogServer
}

func (anRemoteServers *Syslog_AnRemoteServers) GetFilter() yfilter.YFilter { return anRemoteServers.YFilter }

func (anRemoteServers *Syslog_AnRemoteServers) SetFilter(yf yfilter.YFilter) { anRemoteServers.YFilter = yf }

func (anRemoteServers *Syslog_AnRemoteServers) GetGoName(yname string) string {
    if yname == "an-remote-log-server" { return "AnRemoteLogServer" }
    return ""
}

func (anRemoteServers *Syslog_AnRemoteServers) GetSegmentPath() string {
    return "an-remote-servers"
}

func (anRemoteServers *Syslog_AnRemoteServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "an-remote-log-server" {
        for _, c := range anRemoteServers.AnRemoteLogServer {
            if anRemoteServers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_AnRemoteServers_AnRemoteLogServer{}
        anRemoteServers.AnRemoteLogServer = append(anRemoteServers.AnRemoteLogServer, child)
        return &anRemoteServers.AnRemoteLogServer[len(anRemoteServers.AnRemoteLogServer)-1]
    }
    return nil
}

func (anRemoteServers *Syslog_AnRemoteServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range anRemoteServers.AnRemoteLogServer {
        children[anRemoteServers.AnRemoteLogServer[i].GetSegmentPath()] = &anRemoteServers.AnRemoteLogServer[i]
    }
    return children
}

func (anRemoteServers *Syslog_AnRemoteServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (anRemoteServers *Syslog_AnRemoteServers) GetBundleName() string { return "cisco_ios_xr" }

func (anRemoteServers *Syslog_AnRemoteServers) GetYangName() string { return "an-remote-servers" }

func (anRemoteServers *Syslog_AnRemoteServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (anRemoteServers *Syslog_AnRemoteServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (anRemoteServers *Syslog_AnRemoteServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (anRemoteServers *Syslog_AnRemoteServers) SetParent(parent types.Entity) { anRemoteServers.parent = parent }

func (anRemoteServers *Syslog_AnRemoteServers) GetParent() types.Entity { return anRemoteServers.parent }

func (anRemoteServers *Syslog_AnRemoteServers) GetParentYangName() string { return "syslog" }

// Syslog_AnRemoteServers_AnRemoteLogServer
// AN Remote Log Servers
type Syslog_AnRemoteServers_AnRemoteLogServer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP Address. The type is string.
    IpAddress interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Severity. The type is string.
    VrfSeverity interface{}

    // Remote-Host Discriminator. The type is string.
    RhDiscriminator interface{}
}

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetFilter() yfilter.YFilter { return anRemoteLogServer.YFilter }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) SetFilter(yf yfilter.YFilter) { anRemoteLogServer.YFilter = yf }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-severity" { return "VrfSeverity" }
    if yname == "rh-discriminator" { return "RhDiscriminator" }
    return ""
}

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetSegmentPath() string {
    return "an-remote-log-server"
}

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = anRemoteLogServer.IpAddress
    leafs["vrf-name"] = anRemoteLogServer.VrfName
    leafs["vrf-severity"] = anRemoteLogServer.VrfSeverity
    leafs["rh-discriminator"] = anRemoteLogServer.RhDiscriminator
    return leafs
}

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetBundleName() string { return "cisco_ios_xr" }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetYangName() string { return "an-remote-log-server" }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) SetParent(parent types.Entity) { anRemoteLogServer.parent = parent }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetParent() types.Entity { return anRemoteLogServer.parent }

func (anRemoteLogServer *Syslog_AnRemoteServers_AnRemoteLogServer) GetParentYangName() string { return "an-remote-servers" }

// Syslog_Messages
// Message table information
type Syslog_Messages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A system message. The type is slice of Syslog_Messages_Message.
    Message []Syslog_Messages_Message
}

func (messages *Syslog_Messages) GetFilter() yfilter.YFilter { return messages.YFilter }

func (messages *Syslog_Messages) SetFilter(yf yfilter.YFilter) { messages.YFilter = yf }

func (messages *Syslog_Messages) GetGoName(yname string) string {
    if yname == "message" { return "Message" }
    return ""
}

func (messages *Syslog_Messages) GetSegmentPath() string {
    return "messages"
}

func (messages *Syslog_Messages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "message" {
        for _, c := range messages.Message {
            if messages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_Messages_Message{}
        messages.Message = append(messages.Message, child)
        return &messages.Message[len(messages.Message)-1]
    }
    return nil
}

func (messages *Syslog_Messages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range messages.Message {
        children[messages.Message[i].GetSegmentPath()] = &messages.Message[i]
    }
    return children
}

func (messages *Syslog_Messages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (messages *Syslog_Messages) GetBundleName() string { return "cisco_ios_xr" }

func (messages *Syslog_Messages) GetYangName() string { return "messages" }

func (messages *Syslog_Messages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (messages *Syslog_Messages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (messages *Syslog_Messages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (messages *Syslog_Messages) SetParent(parent types.Entity) { messages.parent = parent }

func (messages *Syslog_Messages) GetParent() types.Entity { return messages.parent }

func (messages *Syslog_Messages) GetParentYangName() string { return "syslog" }

// Syslog_Messages_Message
// A system message
type Syslog_Messages_Message struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Message ID of the system message. The type is
    // interface{} with range: -2147483648..2147483647.
    MessageId interface{}

    // Message card location: 'RP', 'DRP', 'LC', 'SC', 'SP' or 'UNK' . The type is
    // string.
    CardType interface{}

    // Message source location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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

func (message *Syslog_Messages_Message) GetFilter() yfilter.YFilter { return message.YFilter }

func (message *Syslog_Messages_Message) SetFilter(yf yfilter.YFilter) { message.YFilter = yf }

func (message *Syslog_Messages_Message) GetGoName(yname string) string {
    if yname == "message-id" { return "MessageId" }
    if yname == "card-type" { return "CardType" }
    if yname == "node-name" { return "NodeName" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "time-of-day" { return "TimeOfDay" }
    if yname == "time-zone" { return "TimeZone" }
    if yname == "process-name" { return "ProcessName" }
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "message-name" { return "MessageName" }
    if yname == "severity" { return "Severity" }
    if yname == "text" { return "Text" }
    return ""
}

func (message *Syslog_Messages_Message) GetSegmentPath() string {
    return "message" + "[message-id='" + fmt.Sprintf("%v", message.MessageId) + "']"
}

func (message *Syslog_Messages_Message) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (message *Syslog_Messages_Message) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (message *Syslog_Messages_Message) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["message-id"] = message.MessageId
    leafs["card-type"] = message.CardType
    leafs["node-name"] = message.NodeName
    leafs["time-stamp"] = message.TimeStamp
    leafs["time-of-day"] = message.TimeOfDay
    leafs["time-zone"] = message.TimeZone
    leafs["process-name"] = message.ProcessName
    leafs["category"] = message.Category
    leafs["group"] = message.Group
    leafs["message-name"] = message.MessageName
    leafs["severity"] = message.Severity
    leafs["text"] = message.Text
    return leafs
}

func (message *Syslog_Messages_Message) GetBundleName() string { return "cisco_ios_xr" }

func (message *Syslog_Messages_Message) GetYangName() string { return "message" }

func (message *Syslog_Messages_Message) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (message *Syslog_Messages_Message) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (message *Syslog_Messages_Message) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (message *Syslog_Messages_Message) SetParent(parent types.Entity) { message.parent = parent }

func (message *Syslog_Messages_Message) GetParent() types.Entity { return message.parent }

func (message *Syslog_Messages_Message) GetParentYangName() string { return "messages" }

// Syslog_LoggingStatistics
// Logging statistics information
type Syslog_LoggingStatistics struct {
    parent types.Entity
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
    RemoteLoggingStat []Syslog_LoggingStatistics_RemoteLoggingStat

    // TLS Remote logging statistics. The type is slice of
    // Syslog_LoggingStatistics_TlsRemoteLoggingStat.
    TlsRemoteLoggingStat []Syslog_LoggingStatistics_TlsRemoteLoggingStat

    // File logging statistics. The type is slice of
    // Syslog_LoggingStatistics_FileLoggingStat.
    FileLoggingStat []Syslog_LoggingStatistics_FileLoggingStat
}

func (loggingStatistics *Syslog_LoggingStatistics) GetFilter() yfilter.YFilter { return loggingStatistics.YFilter }

func (loggingStatistics *Syslog_LoggingStatistics) SetFilter(yf yfilter.YFilter) { loggingStatistics.YFilter = yf }

func (loggingStatistics *Syslog_LoggingStatistics) GetGoName(yname string) string {
    if yname == "logging-stats" { return "LoggingStats" }
    if yname == "console-logging-stats" { return "ConsoleLoggingStats" }
    if yname == "monitor-logging-stats" { return "MonitorLoggingStats" }
    if yname == "trap-logging-stats" { return "TrapLoggingStats" }
    if yname == "buffer-logging-stats" { return "BufferLoggingStats" }
    if yname == "remote-logging-stat" { return "RemoteLoggingStat" }
    if yname == "tls-remote-logging-stat" { return "TlsRemoteLoggingStat" }
    if yname == "file-logging-stat" { return "FileLoggingStat" }
    return ""
}

func (loggingStatistics *Syslog_LoggingStatistics) GetSegmentPath() string {
    return "logging-statistics"
}

func (loggingStatistics *Syslog_LoggingStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logging-stats" {
        return &loggingStatistics.LoggingStats
    }
    if childYangName == "console-logging-stats" {
        return &loggingStatistics.ConsoleLoggingStats
    }
    if childYangName == "monitor-logging-stats" {
        return &loggingStatistics.MonitorLoggingStats
    }
    if childYangName == "trap-logging-stats" {
        return &loggingStatistics.TrapLoggingStats
    }
    if childYangName == "buffer-logging-stats" {
        return &loggingStatistics.BufferLoggingStats
    }
    if childYangName == "remote-logging-stat" {
        for _, c := range loggingStatistics.RemoteLoggingStat {
            if loggingStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_LoggingStatistics_RemoteLoggingStat{}
        loggingStatistics.RemoteLoggingStat = append(loggingStatistics.RemoteLoggingStat, child)
        return &loggingStatistics.RemoteLoggingStat[len(loggingStatistics.RemoteLoggingStat)-1]
    }
    if childYangName == "tls-remote-logging-stat" {
        for _, c := range loggingStatistics.TlsRemoteLoggingStat {
            if loggingStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_LoggingStatistics_TlsRemoteLoggingStat{}
        loggingStatistics.TlsRemoteLoggingStat = append(loggingStatistics.TlsRemoteLoggingStat, child)
        return &loggingStatistics.TlsRemoteLoggingStat[len(loggingStatistics.TlsRemoteLoggingStat)-1]
    }
    if childYangName == "file-logging-stat" {
        for _, c := range loggingStatistics.FileLoggingStat {
            if loggingStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Syslog_LoggingStatistics_FileLoggingStat{}
        loggingStatistics.FileLoggingStat = append(loggingStatistics.FileLoggingStat, child)
        return &loggingStatistics.FileLoggingStat[len(loggingStatistics.FileLoggingStat)-1]
    }
    return nil
}

func (loggingStatistics *Syslog_LoggingStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["logging-stats"] = &loggingStatistics.LoggingStats
    children["console-logging-stats"] = &loggingStatistics.ConsoleLoggingStats
    children["monitor-logging-stats"] = &loggingStatistics.MonitorLoggingStats
    children["trap-logging-stats"] = &loggingStatistics.TrapLoggingStats
    children["buffer-logging-stats"] = &loggingStatistics.BufferLoggingStats
    for i := range loggingStatistics.RemoteLoggingStat {
        children[loggingStatistics.RemoteLoggingStat[i].GetSegmentPath()] = &loggingStatistics.RemoteLoggingStat[i]
    }
    for i := range loggingStatistics.TlsRemoteLoggingStat {
        children[loggingStatistics.TlsRemoteLoggingStat[i].GetSegmentPath()] = &loggingStatistics.TlsRemoteLoggingStat[i]
    }
    for i := range loggingStatistics.FileLoggingStat {
        children[loggingStatistics.FileLoggingStat[i].GetSegmentPath()] = &loggingStatistics.FileLoggingStat[i]
    }
    return children
}

func (loggingStatistics *Syslog_LoggingStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loggingStatistics *Syslog_LoggingStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (loggingStatistics *Syslog_LoggingStatistics) GetYangName() string { return "logging-statistics" }

func (loggingStatistics *Syslog_LoggingStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loggingStatistics *Syslog_LoggingStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loggingStatistics *Syslog_LoggingStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loggingStatistics *Syslog_LoggingStatistics) SetParent(parent types.Entity) { loggingStatistics.parent = parent }

func (loggingStatistics *Syslog_LoggingStatistics) GetParent() types.Entity { return loggingStatistics.parent }

func (loggingStatistics *Syslog_LoggingStatistics) GetParentYangName() string { return "syslog" }

// Syslog_LoggingStatistics_LoggingStats
// Logging statistics
type Syslog_LoggingStatistics_LoggingStats struct {
    parent types.Entity
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

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetFilter() yfilter.YFilter { return loggingStats.YFilter }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) SetFilter(yf yfilter.YFilter) { loggingStats.YFilter = yf }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetGoName(yname string) string {
    if yname == "is-log-enabled" { return "IsLogEnabled" }
    if yname == "drop-count" { return "DropCount" }
    if yname == "flush-count" { return "FlushCount" }
    if yname == "overrun-count" { return "OverrunCount" }
    return ""
}

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetSegmentPath() string {
    return "logging-stats"
}

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-log-enabled"] = loggingStats.IsLogEnabled
    leafs["drop-count"] = loggingStats.DropCount
    leafs["flush-count"] = loggingStats.FlushCount
    leafs["overrun-count"] = loggingStats.OverrunCount
    return leafs
}

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetBundleName() string { return "cisco_ios_xr" }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetYangName() string { return "logging-stats" }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) SetParent(parent types.Entity) { loggingStats.parent = parent }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetParent() types.Entity { return loggingStats.parent }

func (loggingStats *Syslog_LoggingStatistics_LoggingStats) GetParentYangName() string { return "logging-statistics" }

// Syslog_LoggingStatistics_ConsoleLoggingStats
// Console logging statistics
type Syslog_LoggingStatistics_ConsoleLoggingStats struct {
    parent types.Entity
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

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetFilter() yfilter.YFilter { return consoleLoggingStats.YFilter }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) SetFilter(yf yfilter.YFilter) { consoleLoggingStats.YFilter = yf }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetGoName(yname string) string {
    if yname == "is-log-enabled" { return "IsLogEnabled" }
    if yname == "severity" { return "Severity" }
    if yname == "message-count" { return "MessageCount" }
    if yname == "buffer-size" { return "BufferSize" }
    return ""
}

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetSegmentPath() string {
    return "console-logging-stats"
}

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-log-enabled"] = consoleLoggingStats.IsLogEnabled
    leafs["severity"] = consoleLoggingStats.Severity
    leafs["message-count"] = consoleLoggingStats.MessageCount
    leafs["buffer-size"] = consoleLoggingStats.BufferSize
    return leafs
}

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetBundleName() string { return "cisco_ios_xr" }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetYangName() string { return "console-logging-stats" }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) SetParent(parent types.Entity) { consoleLoggingStats.parent = parent }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetParent() types.Entity { return consoleLoggingStats.parent }

func (consoleLoggingStats *Syslog_LoggingStatistics_ConsoleLoggingStats) GetParentYangName() string { return "logging-statistics" }

// Syslog_LoggingStatistics_MonitorLoggingStats
// Monitor loggingstatistics
type Syslog_LoggingStatistics_MonitorLoggingStats struct {
    parent types.Entity
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

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetFilter() yfilter.YFilter { return monitorLoggingStats.YFilter }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) SetFilter(yf yfilter.YFilter) { monitorLoggingStats.YFilter = yf }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetGoName(yname string) string {
    if yname == "is-log-enabled" { return "IsLogEnabled" }
    if yname == "severity" { return "Severity" }
    if yname == "message-count" { return "MessageCount" }
    if yname == "buffer-size" { return "BufferSize" }
    return ""
}

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetSegmentPath() string {
    return "monitor-logging-stats"
}

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-log-enabled"] = monitorLoggingStats.IsLogEnabled
    leafs["severity"] = monitorLoggingStats.Severity
    leafs["message-count"] = monitorLoggingStats.MessageCount
    leafs["buffer-size"] = monitorLoggingStats.BufferSize
    return leafs
}

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetBundleName() string { return "cisco_ios_xr" }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetYangName() string { return "monitor-logging-stats" }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) SetParent(parent types.Entity) { monitorLoggingStats.parent = parent }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetParent() types.Entity { return monitorLoggingStats.parent }

func (monitorLoggingStats *Syslog_LoggingStatistics_MonitorLoggingStats) GetParentYangName() string { return "logging-statistics" }

// Syslog_LoggingStatistics_TrapLoggingStats
// Trap logging statistics
type Syslog_LoggingStatistics_TrapLoggingStats struct {
    parent types.Entity
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

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetFilter() yfilter.YFilter { return trapLoggingStats.YFilter }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) SetFilter(yf yfilter.YFilter) { trapLoggingStats.YFilter = yf }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetGoName(yname string) string {
    if yname == "is-log-enabled" { return "IsLogEnabled" }
    if yname == "severity" { return "Severity" }
    if yname == "message-count" { return "MessageCount" }
    if yname == "buffer-size" { return "BufferSize" }
    return ""
}

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetSegmentPath() string {
    return "trap-logging-stats"
}

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-log-enabled"] = trapLoggingStats.IsLogEnabled
    leafs["severity"] = trapLoggingStats.Severity
    leafs["message-count"] = trapLoggingStats.MessageCount
    leafs["buffer-size"] = trapLoggingStats.BufferSize
    return leafs
}

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetBundleName() string { return "cisco_ios_xr" }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetYangName() string { return "trap-logging-stats" }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) SetParent(parent types.Entity) { trapLoggingStats.parent = parent }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetParent() types.Entity { return trapLoggingStats.parent }

func (trapLoggingStats *Syslog_LoggingStatistics_TrapLoggingStats) GetParentYangName() string { return "logging-statistics" }

// Syslog_LoggingStatistics_BufferLoggingStats
// Buffer logging statistics
type Syslog_LoggingStatistics_BufferLoggingStats struct {
    parent types.Entity
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

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetFilter() yfilter.YFilter { return bufferLoggingStats.YFilter }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) SetFilter(yf yfilter.YFilter) { bufferLoggingStats.YFilter = yf }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetGoName(yname string) string {
    if yname == "is-log-enabled" { return "IsLogEnabled" }
    if yname == "severity" { return "Severity" }
    if yname == "message-count" { return "MessageCount" }
    if yname == "buffer-size" { return "BufferSize" }
    return ""
}

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetSegmentPath() string {
    return "buffer-logging-stats"
}

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-log-enabled"] = bufferLoggingStats.IsLogEnabled
    leafs["severity"] = bufferLoggingStats.Severity
    leafs["message-count"] = bufferLoggingStats.MessageCount
    leafs["buffer-size"] = bufferLoggingStats.BufferSize
    return leafs
}

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetBundleName() string { return "cisco_ios_xr" }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetYangName() string { return "buffer-logging-stats" }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) SetParent(parent types.Entity) { bufferLoggingStats.parent = parent }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetParent() types.Entity { return bufferLoggingStats.parent }

func (bufferLoggingStats *Syslog_LoggingStatistics_BufferLoggingStats) GetParentYangName() string { return "logging-statistics" }

// Syslog_LoggingStatistics_RemoteLoggingStat
// Remote logging statistics
type Syslog_LoggingStatistics_RemoteLoggingStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Remote hostname. The type is string.
    RemoteHostName interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}
}

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetFilter() yfilter.YFilter { return remoteLoggingStat.YFilter }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) SetFilter(yf yfilter.YFilter) { remoteLoggingStat.YFilter = yf }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetGoName(yname string) string {
    if yname == "remote-host-name" { return "RemoteHostName" }
    if yname == "message-count" { return "MessageCount" }
    return ""
}

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetSegmentPath() string {
    return "remote-logging-stat"
}

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-host-name"] = remoteLoggingStat.RemoteHostName
    leafs["message-count"] = remoteLoggingStat.MessageCount
    return leafs
}

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetBundleName() string { return "cisco_ios_xr" }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetYangName() string { return "remote-logging-stat" }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) SetParent(parent types.Entity) { remoteLoggingStat.parent = parent }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetParent() types.Entity { return remoteLoggingStat.parent }

func (remoteLoggingStat *Syslog_LoggingStatistics_RemoteLoggingStat) GetParentYangName() string { return "logging-statistics" }

// Syslog_LoggingStatistics_TlsRemoteLoggingStat
// TLS Remote logging statistics
type Syslog_LoggingStatistics_TlsRemoteLoggingStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TLS Remote hostname. The type is string.
    RemoteHostName interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}
}

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetFilter() yfilter.YFilter { return tlsRemoteLoggingStat.YFilter }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) SetFilter(yf yfilter.YFilter) { tlsRemoteLoggingStat.YFilter = yf }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetGoName(yname string) string {
    if yname == "remote-host-name" { return "RemoteHostName" }
    if yname == "message-count" { return "MessageCount" }
    return ""
}

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetSegmentPath() string {
    return "tls-remote-logging-stat"
}

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-host-name"] = tlsRemoteLoggingStat.RemoteHostName
    leafs["message-count"] = tlsRemoteLoggingStat.MessageCount
    return leafs
}

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetBundleName() string { return "cisco_ios_xr" }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetYangName() string { return "tls-remote-logging-stat" }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) SetParent(parent types.Entity) { tlsRemoteLoggingStat.parent = parent }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetParent() types.Entity { return tlsRemoteLoggingStat.parent }

func (tlsRemoteLoggingStat *Syslog_LoggingStatistics_TlsRemoteLoggingStat) GetParentYangName() string { return "logging-statistics" }

// Syslog_LoggingStatistics_FileLoggingStat
// File logging statistics
type Syslog_LoggingStatistics_FileLoggingStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // File name for logging messages. The type is string.
    FileName interface{}

    // Message count. The type is interface{} with range: 0..4294967295.
    MessageCount interface{}
}

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetFilter() yfilter.YFilter { return fileLoggingStat.YFilter }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) SetFilter(yf yfilter.YFilter) { fileLoggingStat.YFilter = yf }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetGoName(yname string) string {
    if yname == "file-name" { return "FileName" }
    if yname == "message-count" { return "MessageCount" }
    return ""
}

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetSegmentPath() string {
    return "file-logging-stat"
}

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["file-name"] = fileLoggingStat.FileName
    leafs["message-count"] = fileLoggingStat.MessageCount
    return leafs
}

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetBundleName() string { return "cisco_ios_xr" }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetYangName() string { return "file-logging-stat" }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) SetParent(parent types.Entity) { fileLoggingStat.parent = parent }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetParent() types.Entity { return fileLoggingStat.parent }

func (fileLoggingStat *Syslog_LoggingStatistics_FileLoggingStat) GetParentYangName() string { return "logging-statistics" }

