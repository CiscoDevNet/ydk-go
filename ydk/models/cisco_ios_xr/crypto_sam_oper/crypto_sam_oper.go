// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-sam package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sam: Software authentication manager certificate information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_sam_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_sam_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-sam-oper sam}", reflect.TypeOf(Sam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-sam-oper:sam", reflect.TypeOf(Sam{}))
}

// LogError represents Log errors
type LogError string

const (
    // Log error is not known
    LogError_unknown LogError = "unknown"

    // Log error is message error
    LogError_log_message_error LogError = "log-message-error"

    // Log error is get issuer name failed
    LogError_get_issuer_name_failed LogError = "get-issuer-name-failed"
)

// LogCode represents Log code types
type LogCode string

const (
    // Log code is not known
    LogCode_unknown LogCode = "unknown"

    // Log code is SAM server restarted through router
    // reboot
    LogCode_sam_server_restared_router_reboot LogCode = "sam-server-restared-router-reboot"

    // Log code is SAM server restarted
    LogCode_sam_server_restared LogCode = "sam-server-restared"

    // Log code is Added certificate in table
    LogCode_added_certificate_in_table LogCode = "added-certificate-in-table"

    // Log code is Copied certificate in table
    LogCode_copied_certificate_in_table LogCode = "copied-certificate-in-table"

    // Log code is Certificate flag changed
    LogCode_certificate_flag_changed LogCode = "certificate-flag-changed"

    // Log code is validated ceritificate
    LogCode_validated_certificate LogCode = "validated-certificate"

    // Log code is Ceritificate expired detected
    LogCode_certificate_expired_detected LogCode = "certificate-expired-detected"

    // Log code is Ceritificate revoked detected
    LogCode_certificate_revoked_detected LogCode = "certificate-revoked-detected"

    // Log code is CA Ceritificate expired detected
    LogCode_ca_certificate_expired_detected LogCode = "ca-certificate-expired-detected"

    // Log code is CA Ceritificate revoked detected
    LogCode_ca_certificate_revoked_detected LogCode = "ca-certificate-revoked-detected"

    // Log code is Deleted certificate from table
    LogCode_deleted_certificate_from_table LogCode = "deleted-certificate-from-table"

    // Log code is CRL added/updated in table
    LogCode_crl_added_updated_in_table LogCode = "crl-added-updated-in-table"

    // Log code is Checked memory digest
    LogCode_checked_memory_digest LogCode = "checked-memory-digest"

    // Log code is NVRAM digest Mistmatch detected
    LogCode_nvram_digest_mismatch_detected LogCode = "nvram-digest-mismatch-detected"

    // Log code is Insecure backup file detected
    LogCode_insecure_backup_file_detected LogCode = "insecure-backup-file-detected"

    // Log code is Error during restore operation,
    // backup file might have not been intact
    LogCode_error_restore_operation LogCode = "error-restore-operation"

    // Log code is Found backup file on NVRAM for SAM
    // log had been deleted
    LogCode_backup_file_on_nvram_deleted LogCode = "backup-file-on-nvram-deleted"

    // Log code is SAM log backup file recovered from
    // system database
    LogCode_sam_log_file_recovered_from_system_database LogCode = "sam-log-file-recovered-from-system-database"

    // Log code is validated ELF
    LogCode_validated_elf LogCode = "validated-elf"

    // Log code is SAM system database name space
    // deleted/recovered by SAM
    LogCode_namespace_deleted_recovered_by_sam LogCode = "namespace-deleted-recovered-by-sam"
)

// CertificateIssuer represents Certificate issuers
type CertificateIssuer string

const (
    // Issuer is not known
    CertificateIssuer_unknown CertificateIssuer = "unknown"

    // Issuer is code signing server certificate
    // authority
    CertificateIssuer_code_signing_server_certificate_authority CertificateIssuer = "code-signing-server-certificate-authority"
)

// LogTables represents Log tables
type LogTables string

const (
    // Table is not known
    LogTables_unkown LogTables = "unkown"

    // Table is memory digest table
    LogTables_memory_digest_table LogTables = "memory-digest-table"

    // Table is system database digest table
    LogTables_system_database_digest LogTables = "system-database-digest"

    // Table is SAM table
    LogTables_sam_tables LogTables = "sam-tables"
)

// Sam
// Software authentication manager certificate
// information
type Sam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SAM system information.
    SystemInformation Sam_SystemInformation

    // SAM log content table information.
    LogContents Sam_LogContents

    // Certificate device table information.
    Devices Sam_Devices

    // SAM certificate information  package.
    Packages Sam_Packages

    // Certificate revocation list index table information.
    CertificateRevocations Sam_CertificateRevocations

    // Certificate revocation list summary information .
    CertificateRevocationListSummary Sam_CertificateRevocationListSummary
}

func (sam *Sam) GetFilter() yfilter.YFilter { return sam.YFilter }

func (sam *Sam) SetFilter(yf yfilter.YFilter) { sam.YFilter = yf }

func (sam *Sam) GetGoName(yname string) string {
    if yname == "system-information" { return "SystemInformation" }
    if yname == "log-contents" { return "LogContents" }
    if yname == "devices" { return "Devices" }
    if yname == "packages" { return "Packages" }
    if yname == "certificate-revocations" { return "CertificateRevocations" }
    if yname == "certificate-revocation-list-summary" { return "CertificateRevocationListSummary" }
    return ""
}

func (sam *Sam) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-sam-oper:sam"
}

func (sam *Sam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "system-information" {
        return &sam.SystemInformation
    }
    if childYangName == "log-contents" {
        return &sam.LogContents
    }
    if childYangName == "devices" {
        return &sam.Devices
    }
    if childYangName == "packages" {
        return &sam.Packages
    }
    if childYangName == "certificate-revocations" {
        return &sam.CertificateRevocations
    }
    if childYangName == "certificate-revocation-list-summary" {
        return &sam.CertificateRevocationListSummary
    }
    return nil
}

func (sam *Sam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["system-information"] = &sam.SystemInformation
    children["log-contents"] = &sam.LogContents
    children["devices"] = &sam.Devices
    children["packages"] = &sam.Packages
    children["certificate-revocations"] = &sam.CertificateRevocations
    children["certificate-revocation-list-summary"] = &sam.CertificateRevocationListSummary
    return children
}

func (sam *Sam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sam *Sam) GetBundleName() string { return "cisco_ios_xr" }

func (sam *Sam) GetYangName() string { return "sam" }

func (sam *Sam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sam *Sam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sam *Sam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sam *Sam) SetParent(parent types.Entity) { sam.parent = parent }

func (sam *Sam) GetParent() types.Entity { return sam.parent }

func (sam *Sam) GetParentYangName() string { return "Cisco-IOS-XR-crypto-sam-oper" }

// Sam_SystemInformation
// SAM system information
type Sam_SystemInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // True if SAM status information runs. The type is bool.
    IsRunning interface{}

    // Prompt interval atreboot time in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PromptInterval interface{}

    // True if promptdefault response is true. The type is bool.
    IsDefaultResponse interface{}
}

func (systemInformation *Sam_SystemInformation) GetFilter() yfilter.YFilter { return systemInformation.YFilter }

func (systemInformation *Sam_SystemInformation) SetFilter(yf yfilter.YFilter) { systemInformation.YFilter = yf }

func (systemInformation *Sam_SystemInformation) GetGoName(yname string) string {
    if yname == "is-running" { return "IsRunning" }
    if yname == "prompt-interval" { return "PromptInterval" }
    if yname == "is-default-response" { return "IsDefaultResponse" }
    return ""
}

func (systemInformation *Sam_SystemInformation) GetSegmentPath() string {
    return "system-information"
}

func (systemInformation *Sam_SystemInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemInformation *Sam_SystemInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemInformation *Sam_SystemInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-running"] = systemInformation.IsRunning
    leafs["prompt-interval"] = systemInformation.PromptInterval
    leafs["is-default-response"] = systemInformation.IsDefaultResponse
    return leafs
}

func (systemInformation *Sam_SystemInformation) GetBundleName() string { return "cisco_ios_xr" }

func (systemInformation *Sam_SystemInformation) GetYangName() string { return "system-information" }

func (systemInformation *Sam_SystemInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemInformation *Sam_SystemInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemInformation *Sam_SystemInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemInformation *Sam_SystemInformation) SetParent(parent types.Entity) { systemInformation.parent = parent }

func (systemInformation *Sam_SystemInformation) GetParent() types.Entity { return systemInformation.parent }

func (systemInformation *Sam_SystemInformation) GetParentYangName() string { return "sam" }

// Sam_LogContents
// SAM log content table information
type Sam_LogContents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of lines for SAM log message. The type is slice of
    // Sam_LogContents_LogContent.
    LogContent []Sam_LogContents_LogContent
}

func (logContents *Sam_LogContents) GetFilter() yfilter.YFilter { return logContents.YFilter }

func (logContents *Sam_LogContents) SetFilter(yf yfilter.YFilter) { logContents.YFilter = yf }

func (logContents *Sam_LogContents) GetGoName(yname string) string {
    if yname == "log-content" { return "LogContent" }
    return ""
}

func (logContents *Sam_LogContents) GetSegmentPath() string {
    return "log-contents"
}

func (logContents *Sam_LogContents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-content" {
        for _, c := range logContents.LogContent {
            if logContents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sam_LogContents_LogContent{}
        logContents.LogContent = append(logContents.LogContent, child)
        return &logContents.LogContent[len(logContents.LogContent)-1]
    }
    return nil
}

func (logContents *Sam_LogContents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logContents.LogContent {
        children[logContents.LogContent[i].GetSegmentPath()] = &logContents.LogContent[i]
    }
    return children
}

func (logContents *Sam_LogContents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logContents *Sam_LogContents) GetBundleName() string { return "cisco_ios_xr" }

func (logContents *Sam_LogContents) GetYangName() string { return "log-contents" }

func (logContents *Sam_LogContents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContents *Sam_LogContents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContents *Sam_LogContents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContents *Sam_LogContents) SetParent(parent types.Entity) { logContents.parent = parent }

func (logContents *Sam_LogContents) GetParent() types.Entity { return logContents.parent }

func (logContents *Sam_LogContents) GetParentYangName() string { return "sam" }

// Sam_LogContents_LogContent
// Number of lines for SAM log message
type Sam_LogContents_LogContent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Number of lines. The type is interface{} with
    // range: -2147483648..2147483647.
    NumberOfLines interface{}

    // Total log entries available. The type is interface{} with range:
    // 0..4294967295.
    TotalEntries interface{}

    // Total entries shown. The type is interface{} with range: 0..4294967295.
    EntriesShown interface{}

    // SAM logs. The type is slice of Sam_LogContents_LogContent_Logs.
    Logs []Sam_LogContents_LogContent_Logs
}

func (logContent *Sam_LogContents_LogContent) GetFilter() yfilter.YFilter { return logContent.YFilter }

func (logContent *Sam_LogContents_LogContent) SetFilter(yf yfilter.YFilter) { logContent.YFilter = yf }

func (logContent *Sam_LogContents_LogContent) GetGoName(yname string) string {
    if yname == "number-of-lines" { return "NumberOfLines" }
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "entries-shown" { return "EntriesShown" }
    if yname == "logs" { return "Logs" }
    return ""
}

func (logContent *Sam_LogContents_LogContent) GetSegmentPath() string {
    return "log-content" + "[number-of-lines='" + fmt.Sprintf("%v", logContent.NumberOfLines) + "']"
}

func (logContent *Sam_LogContents_LogContent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logs" {
        for _, c := range logContent.Logs {
            if logContent.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sam_LogContents_LogContent_Logs{}
        logContent.Logs = append(logContent.Logs, child)
        return &logContent.Logs[len(logContent.Logs)-1]
    }
    return nil
}

func (logContent *Sam_LogContents_LogContent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logContent.Logs {
        children[logContent.Logs[i].GetSegmentPath()] = &logContent.Logs[i]
    }
    return children
}

func (logContent *Sam_LogContents_LogContent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number-of-lines"] = logContent.NumberOfLines
    leafs["total-entries"] = logContent.TotalEntries
    leafs["entries-shown"] = logContent.EntriesShown
    return leafs
}

func (logContent *Sam_LogContents_LogContent) GetBundleName() string { return "cisco_ios_xr" }

func (logContent *Sam_LogContents_LogContent) GetYangName() string { return "log-content" }

func (logContent *Sam_LogContents_LogContent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContent *Sam_LogContents_LogContent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContent *Sam_LogContents_LogContent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContent *Sam_LogContents_LogContent) SetParent(parent types.Entity) { logContent.parent = parent }

func (logContent *Sam_LogContents_LogContent) GetParent() types.Entity { return logContent.parent }

func (logContent *Sam_LogContents_LogContent) GetParentYangName() string { return "log-contents" }

// Sam_LogContents_LogContent_Logs
// SAM logs
type Sam_LogContents_LogContent_Logs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log time. The type is string.
    Time interface{}

    // Log code. The type is LogCode.
    Code interface{}

    // Target device. The type is string.
    TargetDevice interface{}

    // Device index. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // Log error message. The type is LogError.
    Error interface{}

    // Issuer of the certificate. The type is CertificateIssuer.
    Issuer interface{}

    // Serial number. The type is string.
    SerialNo interface{}

    // SAM table index. The type is interface{} with range: 0..4294967295.
    SamTableIndex interface{}

    // Last update time of the certificate. The type is string.
    UpdateTime interface{}

    // source device name. The type is string.
    SourceDevice interface{}

    // Log table information. The type is LogTables.
    Table interface{}
}

func (logs *Sam_LogContents_LogContent_Logs) GetFilter() yfilter.YFilter { return logs.YFilter }

func (logs *Sam_LogContents_LogContent_Logs) SetFilter(yf yfilter.YFilter) { logs.YFilter = yf }

func (logs *Sam_LogContents_LogContent_Logs) GetGoName(yname string) string {
    if yname == "time" { return "Time" }
    if yname == "code" { return "Code" }
    if yname == "target-device" { return "TargetDevice" }
    if yname == "index" { return "Index" }
    if yname == "error" { return "Error" }
    if yname == "issuer" { return "Issuer" }
    if yname == "serial-no" { return "SerialNo" }
    if yname == "sam-table-index" { return "SamTableIndex" }
    if yname == "update-time" { return "UpdateTime" }
    if yname == "source-device" { return "SourceDevice" }
    if yname == "table" { return "Table" }
    return ""
}

func (logs *Sam_LogContents_LogContent_Logs) GetSegmentPath() string {
    return "logs"
}

func (logs *Sam_LogContents_LogContent_Logs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logs *Sam_LogContents_LogContent_Logs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logs *Sam_LogContents_LogContent_Logs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time"] = logs.Time
    leafs["code"] = logs.Code
    leafs["target-device"] = logs.TargetDevice
    leafs["index"] = logs.Index
    leafs["error"] = logs.Error
    leafs["issuer"] = logs.Issuer
    leafs["serial-no"] = logs.SerialNo
    leafs["sam-table-index"] = logs.SamTableIndex
    leafs["update-time"] = logs.UpdateTime
    leafs["source-device"] = logs.SourceDevice
    leafs["table"] = logs.Table
    return leafs
}

func (logs *Sam_LogContents_LogContent_Logs) GetBundleName() string { return "cisco_ios_xr" }

func (logs *Sam_LogContents_LogContent_Logs) GetYangName() string { return "logs" }

func (logs *Sam_LogContents_LogContent_Logs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logs *Sam_LogContents_LogContent_Logs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logs *Sam_LogContents_LogContent_Logs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logs *Sam_LogContents_LogContent_Logs) SetParent(parent types.Entity) { logs.parent = parent }

func (logs *Sam_LogContents_LogContent_Logs) GetParent() types.Entity { return logs.parent }

func (logs *Sam_LogContents_LogContent_Logs) GetParentYangName() string { return "log-content" }

// Sam_Devices
// Certificate device table information
type Sam_Devices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Certificate table device information. The type is slice of
    // Sam_Devices_Device.
    Device []Sam_Devices_Device
}

func (devices *Sam_Devices) GetFilter() yfilter.YFilter { return devices.YFilter }

func (devices *Sam_Devices) SetFilter(yf yfilter.YFilter) { devices.YFilter = yf }

func (devices *Sam_Devices) GetGoName(yname string) string {
    if yname == "device" { return "Device" }
    return ""
}

func (devices *Sam_Devices) GetSegmentPath() string {
    return "devices"
}

func (devices *Sam_Devices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "device" {
        for _, c := range devices.Device {
            if devices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sam_Devices_Device{}
        devices.Device = append(devices.Device, child)
        return &devices.Device[len(devices.Device)-1]
    }
    return nil
}

func (devices *Sam_Devices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range devices.Device {
        children[devices.Device[i].GetSegmentPath()] = &devices.Device[i]
    }
    return children
}

func (devices *Sam_Devices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (devices *Sam_Devices) GetBundleName() string { return "cisco_ios_xr" }

func (devices *Sam_Devices) GetYangName() string { return "devices" }

func (devices *Sam_Devices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (devices *Sam_Devices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (devices *Sam_Devices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (devices *Sam_Devices) SetParent(parent types.Entity) { devices.parent = parent }

func (devices *Sam_Devices) GetParent() types.Entity { return devices.parent }

func (devices *Sam_Devices) GetParentYangName() string { return "sam" }

// Sam_Devices_Device
// Certificate table device information
type Sam_Devices_Device struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify device name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    DeviceName interface{}

    // Certificate table information.
    Certificate Sam_Devices_Device_Certificate
}

func (device *Sam_Devices_Device) GetFilter() yfilter.YFilter { return device.YFilter }

func (device *Sam_Devices_Device) SetFilter(yf yfilter.YFilter) { device.YFilter = yf }

func (device *Sam_Devices_Device) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "certificate" { return "Certificate" }
    return ""
}

func (device *Sam_Devices_Device) GetSegmentPath() string {
    return "device" + "[device-name='" + fmt.Sprintf("%v", device.DeviceName) + "']"
}

func (device *Sam_Devices_Device) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "certificate" {
        return &device.Certificate
    }
    return nil
}

func (device *Sam_Devices_Device) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["certificate"] = &device.Certificate
    return children
}

func (device *Sam_Devices_Device) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = device.DeviceName
    return leafs
}

func (device *Sam_Devices_Device) GetBundleName() string { return "cisco_ios_xr" }

func (device *Sam_Devices_Device) GetYangName() string { return "device" }

func (device *Sam_Devices_Device) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (device *Sam_Devices_Device) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (device *Sam_Devices_Device) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (device *Sam_Devices_Device) SetParent(parent types.Entity) { device.parent = parent }

func (device *Sam_Devices_Device) GetParent() types.Entity { return device.parent }

func (device *Sam_Devices_Device) GetParentYangName() string { return "devices" }

// Sam_Devices_Device_Certificate
// Certificate table information
type Sam_Devices_Device_Certificate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Certificate table brief information.
    Brief Sam_Devices_Device_Certificate_Brief

    // Certificate detail index table information.
    CertificateIndexes Sam_Devices_Device_Certificate_CertificateIndexes
}

func (certificate *Sam_Devices_Device_Certificate) GetFilter() yfilter.YFilter { return certificate.YFilter }

func (certificate *Sam_Devices_Device_Certificate) SetFilter(yf yfilter.YFilter) { certificate.YFilter = yf }

func (certificate *Sam_Devices_Device_Certificate) GetGoName(yname string) string {
    if yname == "brief" { return "Brief" }
    if yname == "certificate-indexes" { return "CertificateIndexes" }
    return ""
}

func (certificate *Sam_Devices_Device_Certificate) GetSegmentPath() string {
    return "certificate"
}

func (certificate *Sam_Devices_Device_Certificate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief" {
        return &certificate.Brief
    }
    if childYangName == "certificate-indexes" {
        return &certificate.CertificateIndexes
    }
    return nil
}

func (certificate *Sam_Devices_Device_Certificate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief"] = &certificate.Brief
    children["certificate-indexes"] = &certificate.CertificateIndexes
    return children
}

func (certificate *Sam_Devices_Device_Certificate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (certificate *Sam_Devices_Device_Certificate) GetBundleName() string { return "cisco_ios_xr" }

func (certificate *Sam_Devices_Device_Certificate) GetYangName() string { return "certificate" }

func (certificate *Sam_Devices_Device_Certificate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificate *Sam_Devices_Device_Certificate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificate *Sam_Devices_Device_Certificate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificate *Sam_Devices_Device_Certificate) SetParent(parent types.Entity) { certificate.parent = parent }

func (certificate *Sam_Devices_Device_Certificate) GetParent() types.Entity { return certificate.parent }

func (certificate *Sam_Devices_Device_Certificate) GetParentYangName() string { return "device" }

// Sam_Devices_Device_Certificate_Brief
// Certificate table brief information
type Sam_Devices_Device_Certificate_Brief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Certificate location. The type is string.
    Location interface{}

    // Certificate index. The type is interface{} with range: 0..65535.
    CertificateIndex interface{}

    // Certificate flags.
    CertificateFlags Sam_Devices_Device_Certificate_Brief_CertificateFlags
}

func (brief *Sam_Devices_Device_Certificate_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *Sam_Devices_Device_Certificate_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *Sam_Devices_Device_Certificate_Brief) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "certificate-index" { return "CertificateIndex" }
    if yname == "certificate-flags" { return "CertificateFlags" }
    return ""
}

func (brief *Sam_Devices_Device_Certificate_Brief) GetSegmentPath() string {
    return "brief"
}

func (brief *Sam_Devices_Device_Certificate_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "certificate-flags" {
        return &brief.CertificateFlags
    }
    return nil
}

func (brief *Sam_Devices_Device_Certificate_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["certificate-flags"] = &brief.CertificateFlags
    return children
}

func (brief *Sam_Devices_Device_Certificate_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = brief.Location
    leafs["certificate-index"] = brief.CertificateIndex
    return leafs
}

func (brief *Sam_Devices_Device_Certificate_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *Sam_Devices_Device_Certificate_Brief) GetYangName() string { return "brief" }

func (brief *Sam_Devices_Device_Certificate_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *Sam_Devices_Device_Certificate_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *Sam_Devices_Device_Certificate_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *Sam_Devices_Device_Certificate_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *Sam_Devices_Device_Certificate_Brief) GetParent() types.Entity { return brief.parent }

func (brief *Sam_Devices_Device_Certificate_Brief) GetParentYangName() string { return "certificate" }

// Sam_Devices_Device_Certificate_Brief_CertificateFlags
// Certificate flags
type Sam_Devices_Device_Certificate_Brief_CertificateFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trusted flag. The type is bool.
    IsTrusted interface{}

    // Revoked flag. The type is bool.
    IsRevoked interface{}

    // Expired flag. The type is bool.
    IsExpired interface{}

    // Validated flag. The type is bool.
    IsValidated interface{}
}

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetFilter() yfilter.YFilter { return certificateFlags.YFilter }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) SetFilter(yf yfilter.YFilter) { certificateFlags.YFilter = yf }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetGoName(yname string) string {
    if yname == "is-trusted" { return "IsTrusted" }
    if yname == "is-revoked" { return "IsRevoked" }
    if yname == "is-expired" { return "IsExpired" }
    if yname == "is-validated" { return "IsValidated" }
    return ""
}

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetSegmentPath() string {
    return "certificate-flags"
}

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-trusted"] = certificateFlags.IsTrusted
    leafs["is-revoked"] = certificateFlags.IsRevoked
    leafs["is-expired"] = certificateFlags.IsExpired
    leafs["is-validated"] = certificateFlags.IsValidated
    return leafs
}

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetBundleName() string { return "cisco_ios_xr" }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetYangName() string { return "certificate-flags" }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) SetParent(parent types.Entity) { certificateFlags.parent = parent }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetParent() types.Entity { return certificateFlags.parent }

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetParentYangName() string { return "brief" }

// Sam_Devices_Device_Certificate_CertificateIndexes
// Certificate detail index table information
type Sam_Devices_Device_Certificate_CertificateIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Certificate detail index information. The type is slice of
    // Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex.
    CertificateIndex []Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex
}

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetFilter() yfilter.YFilter { return certificateIndexes.YFilter }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) SetFilter(yf yfilter.YFilter) { certificateIndexes.YFilter = yf }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetGoName(yname string) string {
    if yname == "certificate-index" { return "CertificateIndex" }
    return ""
}

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetSegmentPath() string {
    return "certificate-indexes"
}

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "certificate-index" {
        for _, c := range certificateIndexes.CertificateIndex {
            if certificateIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex{}
        certificateIndexes.CertificateIndex = append(certificateIndexes.CertificateIndex, child)
        return &certificateIndexes.CertificateIndex[len(certificateIndexes.CertificateIndex)-1]
    }
    return nil
}

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range certificateIndexes.CertificateIndex {
        children[certificateIndexes.CertificateIndex[i].GetSegmentPath()] = &certificateIndexes.CertificateIndex[i]
    }
    return children
}

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetYangName() string { return "certificate-indexes" }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) SetParent(parent types.Entity) { certificateIndexes.parent = parent }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetParent() types.Entity { return certificateIndexes.parent }

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetParentYangName() string { return "certificate" }

// Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex
// Certificate detail index information
type Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify certificate index. The type is interface{}
    // with range: -2147483648..2147483647.
    Index interface{}

    // Certificate table detail information.
    Detail Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail
}

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetFilter() yfilter.YFilter { return certificateIndex.YFilter }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) SetFilter(yf yfilter.YFilter) { certificateIndex.YFilter = yf }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetSegmentPath() string {
    return "certificate-index" + "[index='" + fmt.Sprintf("%v", certificateIndex.Index) + "']"
}

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &certificateIndex.Detail
    }
    return nil
}

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &certificateIndex.Detail
    return children
}

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = certificateIndex.Index
    return leafs
}

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetBundleName() string { return "cisco_ios_xr" }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetYangName() string { return "certificate-index" }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) SetParent(parent types.Entity) { certificateIndex.parent = parent }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetParent() types.Entity { return certificateIndex.parent }

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetParentYangName() string { return "certificate-indexes" }

// Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail
// Certificate table detail information
type Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Certificate location. The type is string.
    Location interface{}

    // Certificate index. The type is interface{} with range: 0..65535.
    CertificateIndex interface{}

    // Certificate flags.
    CertificateFlags Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags
}

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "certificate-index" { return "CertificateIndex" }
    if yname == "certificate-flags" { return "CertificateFlags" }
    return ""
}

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "certificate-flags" {
        return &detail.CertificateFlags
    }
    return nil
}

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["certificate-flags"] = &detail.CertificateFlags
    return children
}

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = detail.Location
    leafs["certificate-index"] = detail.CertificateIndex
    return leafs
}

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetYangName() string { return "detail" }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetParentYangName() string { return "certificate-index" }

// Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags
// Certificate flags
type Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trusted flag. The type is bool.
    IsTrusted interface{}

    // Revoked flag. The type is bool.
    IsRevoked interface{}

    // Expired flag. The type is bool.
    IsExpired interface{}

    // Validated flag. The type is bool.
    IsValidated interface{}
}

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetFilter() yfilter.YFilter { return certificateFlags.YFilter }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) SetFilter(yf yfilter.YFilter) { certificateFlags.YFilter = yf }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetGoName(yname string) string {
    if yname == "is-trusted" { return "IsTrusted" }
    if yname == "is-revoked" { return "IsRevoked" }
    if yname == "is-expired" { return "IsExpired" }
    if yname == "is-validated" { return "IsValidated" }
    return ""
}

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetSegmentPath() string {
    return "certificate-flags"
}

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-trusted"] = certificateFlags.IsTrusted
    leafs["is-revoked"] = certificateFlags.IsRevoked
    leafs["is-expired"] = certificateFlags.IsExpired
    leafs["is-validated"] = certificateFlags.IsValidated
    return leafs
}

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetBundleName() string { return "cisco_ios_xr" }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetYangName() string { return "certificate-flags" }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) SetParent(parent types.Entity) { certificateFlags.parent = parent }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetParent() types.Entity { return certificateFlags.parent }

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetParentYangName() string { return "detail" }

// Sam_Packages
// SAM certificate information  package
type Sam_Packages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SAM certificate information for a specific package. The type is slice of
    // Sam_Packages_Package.
    Package []Sam_Packages_Package
}

func (packages *Sam_Packages) GetFilter() yfilter.YFilter { return packages.YFilter }

func (packages *Sam_Packages) SetFilter(yf yfilter.YFilter) { packages.YFilter = yf }

func (packages *Sam_Packages) GetGoName(yname string) string {
    if yname == "package" { return "Package" }
    return ""
}

func (packages *Sam_Packages) GetSegmentPath() string {
    return "packages"
}

func (packages *Sam_Packages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        for _, c := range packages.Package {
            if packages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sam_Packages_Package{}
        packages.Package = append(packages.Package, child)
        return &packages.Package[len(packages.Package)-1]
    }
    return nil
}

func (packages *Sam_Packages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range packages.Package {
        children[packages.Package[i].GetSegmentPath()] = &packages.Package[i]
    }
    return children
}

func (packages *Sam_Packages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packages *Sam_Packages) GetBundleName() string { return "cisco_ios_xr" }

func (packages *Sam_Packages) GetYangName() string { return "packages" }

func (packages *Sam_Packages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packages *Sam_Packages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packages *Sam_Packages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packages *Sam_Packages) SetParent(parent types.Entity) { packages.parent = parent }

func (packages *Sam_Packages) GetParent() types.Entity { return packages.parent }

func (packages *Sam_Packages) GetParentYangName() string { return "sam" }

// Sam_Packages_Package
// SAM certificate information for a specific
// package
type Sam_Packages_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify package name. The type is string.
    PackageName interface{}

    // Certificate location. The type is string.
    Location interface{}

    // Certificate index. The type is interface{} with range: 0..65535.
    CertificateIndex interface{}

    // Certificate flags.
    CertificateFlags Sam_Packages_Package_CertificateFlags
}

func (self *Sam_Packages_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Sam_Packages_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Sam_Packages_Package) GetGoName(yname string) string {
    if yname == "package-name" { return "PackageName" }
    if yname == "location" { return "Location" }
    if yname == "certificate-index" { return "CertificateIndex" }
    if yname == "certificate-flags" { return "CertificateFlags" }
    return ""
}

func (self *Sam_Packages_Package) GetSegmentPath() string {
    return "package" + "[package-name='" + fmt.Sprintf("%v", self.PackageName) + "']"
}

func (self *Sam_Packages_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "certificate-flags" {
        return &self.CertificateFlags
    }
    return nil
}

func (self *Sam_Packages_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["certificate-flags"] = &self.CertificateFlags
    return children
}

func (self *Sam_Packages_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["package-name"] = self.PackageName
    leafs["location"] = self.Location
    leafs["certificate-index"] = self.CertificateIndex
    return leafs
}

func (self *Sam_Packages_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Sam_Packages_Package) GetYangName() string { return "package" }

func (self *Sam_Packages_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Sam_Packages_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Sam_Packages_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Sam_Packages_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Sam_Packages_Package) GetParent() types.Entity { return self.parent }

func (self *Sam_Packages_Package) GetParentYangName() string { return "packages" }

// Sam_Packages_Package_CertificateFlags
// Certificate flags
type Sam_Packages_Package_CertificateFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trusted flag. The type is bool.
    IsTrusted interface{}

    // Revoked flag. The type is bool.
    IsRevoked interface{}

    // Expired flag. The type is bool.
    IsExpired interface{}

    // Validated flag. The type is bool.
    IsValidated interface{}
}

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetFilter() yfilter.YFilter { return certificateFlags.YFilter }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) SetFilter(yf yfilter.YFilter) { certificateFlags.YFilter = yf }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetGoName(yname string) string {
    if yname == "is-trusted" { return "IsTrusted" }
    if yname == "is-revoked" { return "IsRevoked" }
    if yname == "is-expired" { return "IsExpired" }
    if yname == "is-validated" { return "IsValidated" }
    return ""
}

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetSegmentPath() string {
    return "certificate-flags"
}

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-trusted"] = certificateFlags.IsTrusted
    leafs["is-revoked"] = certificateFlags.IsRevoked
    leafs["is-expired"] = certificateFlags.IsExpired
    leafs["is-validated"] = certificateFlags.IsValidated
    return leafs
}

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetBundleName() string { return "cisco_ios_xr" }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetYangName() string { return "certificate-flags" }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) SetParent(parent types.Entity) { certificateFlags.parent = parent }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetParent() types.Entity { return certificateFlags.parent }

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetParentYangName() string { return "package" }

// Sam_CertificateRevocations
// Certificate revocation list index table
// information
type Sam_CertificateRevocations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Certificate revocation list index information. The type is slice of
    // Sam_CertificateRevocations_CertificateRevocation.
    CertificateRevocation []Sam_CertificateRevocations_CertificateRevocation
}

func (certificateRevocations *Sam_CertificateRevocations) GetFilter() yfilter.YFilter { return certificateRevocations.YFilter }

func (certificateRevocations *Sam_CertificateRevocations) SetFilter(yf yfilter.YFilter) { certificateRevocations.YFilter = yf }

func (certificateRevocations *Sam_CertificateRevocations) GetGoName(yname string) string {
    if yname == "certificate-revocation" { return "CertificateRevocation" }
    return ""
}

func (certificateRevocations *Sam_CertificateRevocations) GetSegmentPath() string {
    return "certificate-revocations"
}

func (certificateRevocations *Sam_CertificateRevocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "certificate-revocation" {
        for _, c := range certificateRevocations.CertificateRevocation {
            if certificateRevocations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sam_CertificateRevocations_CertificateRevocation{}
        certificateRevocations.CertificateRevocation = append(certificateRevocations.CertificateRevocation, child)
        return &certificateRevocations.CertificateRevocation[len(certificateRevocations.CertificateRevocation)-1]
    }
    return nil
}

func (certificateRevocations *Sam_CertificateRevocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range certificateRevocations.CertificateRevocation {
        children[certificateRevocations.CertificateRevocation[i].GetSegmentPath()] = &certificateRevocations.CertificateRevocation[i]
    }
    return children
}

func (certificateRevocations *Sam_CertificateRevocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (certificateRevocations *Sam_CertificateRevocations) GetBundleName() string { return "cisco_ios_xr" }

func (certificateRevocations *Sam_CertificateRevocations) GetYangName() string { return "certificate-revocations" }

func (certificateRevocations *Sam_CertificateRevocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateRevocations *Sam_CertificateRevocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateRevocations *Sam_CertificateRevocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateRevocations *Sam_CertificateRevocations) SetParent(parent types.Entity) { certificateRevocations.parent = parent }

func (certificateRevocations *Sam_CertificateRevocations) GetParent() types.Entity { return certificateRevocations.parent }

func (certificateRevocations *Sam_CertificateRevocations) GetParentYangName() string { return "sam" }

// Sam_CertificateRevocations_CertificateRevocation
// Certificate revocation list index information
type Sam_CertificateRevocations_CertificateRevocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. CRL index. The type is interface{} with range:
    // -2147483648..2147483647.
    CrlIndex interface{}

    // Certificate revocation list detail information.
    CertificateRevocationListDetail Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail
}

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetFilter() yfilter.YFilter { return certificateRevocation.YFilter }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) SetFilter(yf yfilter.YFilter) { certificateRevocation.YFilter = yf }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetGoName(yname string) string {
    if yname == "crl-index" { return "CrlIndex" }
    if yname == "certificate-revocation-list-detail" { return "CertificateRevocationListDetail" }
    return ""
}

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetSegmentPath() string {
    return "certificate-revocation" + "[crl-index='" + fmt.Sprintf("%v", certificateRevocation.CrlIndex) + "']"
}

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "certificate-revocation-list-detail" {
        return &certificateRevocation.CertificateRevocationListDetail
    }
    return nil
}

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["certificate-revocation-list-detail"] = &certificateRevocation.CertificateRevocationListDetail
    return children
}

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["crl-index"] = certificateRevocation.CrlIndex
    return leafs
}

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetBundleName() string { return "cisco_ios_xr" }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetYangName() string { return "certificate-revocation" }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) SetParent(parent types.Entity) { certificateRevocation.parent = parent }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetParent() types.Entity { return certificateRevocation.parent }

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetParentYangName() string { return "certificate-revocations" }

// Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail
// Certificate revocation list detail information
type Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CRL index. The type is interface{} with range: 0..65535.
    CrlIndex interface{}

    // Updated time of CRL is displayed. The type is string.
    Updates interface{}

    // Issuer name.
    Issuer Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer
}

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetFilter() yfilter.YFilter { return certificateRevocationListDetail.YFilter }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) SetFilter(yf yfilter.YFilter) { certificateRevocationListDetail.YFilter = yf }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetGoName(yname string) string {
    if yname == "crl-index" { return "CrlIndex" }
    if yname == "updates" { return "Updates" }
    if yname == "issuer" { return "Issuer" }
    return ""
}

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetSegmentPath() string {
    return "certificate-revocation-list-detail"
}

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "issuer" {
        return &certificateRevocationListDetail.Issuer
    }
    return nil
}

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["issuer"] = &certificateRevocationListDetail.Issuer
    return children
}

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["crl-index"] = certificateRevocationListDetail.CrlIndex
    leafs["updates"] = certificateRevocationListDetail.Updates
    return leafs
}

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetBundleName() string { return "cisco_ios_xr" }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetYangName() string { return "certificate-revocation-list-detail" }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) SetParent(parent types.Entity) { certificateRevocationListDetail.parent = parent }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetParent() types.Entity { return certificateRevocationListDetail.parent }

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetParentYangName() string { return "certificate-revocation" }

// Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer
// Issuer name
type Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Common name. The type is string.
    CommonName interface{}

    // Organization. The type is string.
    Organization interface{}

    // Country. The type is string.
    Country interface{}
}

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetFilter() yfilter.YFilter { return issuer.YFilter }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) SetFilter(yf yfilter.YFilter) { issuer.YFilter = yf }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetGoName(yname string) string {
    if yname == "common-name" { return "CommonName" }
    if yname == "organization" { return "Organization" }
    if yname == "country" { return "Country" }
    return ""
}

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetSegmentPath() string {
    return "issuer"
}

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-name"] = issuer.CommonName
    leafs["organization"] = issuer.Organization
    leafs["country"] = issuer.Country
    return leafs
}

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetBundleName() string { return "cisco_ios_xr" }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetYangName() string { return "issuer" }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) SetParent(parent types.Entity) { issuer.parent = parent }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetParent() types.Entity { return issuer.parent }

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetParentYangName() string { return "certificate-revocation-list-detail" }

// Sam_CertificateRevocationListSummary
// Certificate revocation list summary information 
type Sam_CertificateRevocationListSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CRL index. The type is interface{} with range: 0..65535.
    CrlIndex interface{}

    // Updated time of CRL is displayed. The type is string.
    Updates interface{}

    // Issuer name.
    Issuer Sam_CertificateRevocationListSummary_Issuer
}

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetFilter() yfilter.YFilter { return certificateRevocationListSummary.YFilter }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) SetFilter(yf yfilter.YFilter) { certificateRevocationListSummary.YFilter = yf }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetGoName(yname string) string {
    if yname == "crl-index" { return "CrlIndex" }
    if yname == "updates" { return "Updates" }
    if yname == "issuer" { return "Issuer" }
    return ""
}

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetSegmentPath() string {
    return "certificate-revocation-list-summary"
}

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "issuer" {
        return &certificateRevocationListSummary.Issuer
    }
    return nil
}

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["issuer"] = &certificateRevocationListSummary.Issuer
    return children
}

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["crl-index"] = certificateRevocationListSummary.CrlIndex
    leafs["updates"] = certificateRevocationListSummary.Updates
    return leafs
}

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetBundleName() string { return "cisco_ios_xr" }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetYangName() string { return "certificate-revocation-list-summary" }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) SetParent(parent types.Entity) { certificateRevocationListSummary.parent = parent }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetParent() types.Entity { return certificateRevocationListSummary.parent }

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetParentYangName() string { return "sam" }

// Sam_CertificateRevocationListSummary_Issuer
// Issuer name
type Sam_CertificateRevocationListSummary_Issuer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Common name. The type is string.
    CommonName interface{}

    // Organization. The type is string.
    Organization interface{}

    // Country. The type is string.
    Country interface{}
}

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetFilter() yfilter.YFilter { return issuer.YFilter }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) SetFilter(yf yfilter.YFilter) { issuer.YFilter = yf }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetGoName(yname string) string {
    if yname == "common-name" { return "CommonName" }
    if yname == "organization" { return "Organization" }
    if yname == "country" { return "Country" }
    return ""
}

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetSegmentPath() string {
    return "issuer"
}

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-name"] = issuer.CommonName
    leafs["organization"] = issuer.Organization
    leafs["country"] = issuer.Country
    return leafs
}

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetBundleName() string { return "cisco_ios_xr" }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetYangName() string { return "issuer" }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) SetParent(parent types.Entity) { issuer.parent = parent }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetParent() types.Entity { return issuer.parent }

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetParentYangName() string { return "certificate-revocation-list-summary" }

