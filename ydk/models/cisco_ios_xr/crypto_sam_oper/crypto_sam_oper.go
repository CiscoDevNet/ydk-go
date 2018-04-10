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

// CertificateIssuer represents Certificate issuers
type CertificateIssuer string

const (
    // Issuer is not known
    CertificateIssuer_unknown CertificateIssuer = "unknown"

    // Issuer is code signing server certificate
    // authority
    CertificateIssuer_code_signing_server_certificate_authority CertificateIssuer = "code-signing-server-certificate-authority"
)

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

// Sam
// Software authentication manager certificate
// information
type Sam struct {
    EntityData types.CommonEntityData
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

func (sam *Sam) GetEntityData() *types.CommonEntityData {
    sam.EntityData.YFilter = sam.YFilter
    sam.EntityData.YangName = "sam"
    sam.EntityData.BundleName = "cisco_ios_xr"
    sam.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-sam-oper"
    sam.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-sam-oper:sam"
    sam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sam.EntityData.Children = make(map[string]types.YChild)
    sam.EntityData.Children["system-information"] = types.YChild{"SystemInformation", &sam.SystemInformation}
    sam.EntityData.Children["log-contents"] = types.YChild{"LogContents", &sam.LogContents}
    sam.EntityData.Children["devices"] = types.YChild{"Devices", &sam.Devices}
    sam.EntityData.Children["packages"] = types.YChild{"Packages", &sam.Packages}
    sam.EntityData.Children["certificate-revocations"] = types.YChild{"CertificateRevocations", &sam.CertificateRevocations}
    sam.EntityData.Children["certificate-revocation-list-summary"] = types.YChild{"CertificateRevocationListSummary", &sam.CertificateRevocationListSummary}
    sam.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sam.EntityData)
}

// Sam_SystemInformation
// SAM system information
type Sam_SystemInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True if SAM status information runs. The type is bool.
    IsRunning interface{}

    // Prompt interval atreboot time in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PromptInterval interface{}

    // True if promptdefault response is true. The type is bool.
    IsDefaultResponse interface{}
}

func (systemInformation *Sam_SystemInformation) GetEntityData() *types.CommonEntityData {
    systemInformation.EntityData.YFilter = systemInformation.YFilter
    systemInformation.EntityData.YangName = "system-information"
    systemInformation.EntityData.BundleName = "cisco_ios_xr"
    systemInformation.EntityData.ParentYangName = "sam"
    systemInformation.EntityData.SegmentPath = "system-information"
    systemInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemInformation.EntityData.Children = make(map[string]types.YChild)
    systemInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    systemInformation.EntityData.Leafs["is-running"] = types.YLeaf{"IsRunning", systemInformation.IsRunning}
    systemInformation.EntityData.Leafs["prompt-interval"] = types.YLeaf{"PromptInterval", systemInformation.PromptInterval}
    systemInformation.EntityData.Leafs["is-default-response"] = types.YLeaf{"IsDefaultResponse", systemInformation.IsDefaultResponse}
    return &(systemInformation.EntityData)
}

// Sam_LogContents
// SAM log content table information
type Sam_LogContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of lines for SAM log message. The type is slice of
    // Sam_LogContents_LogContent.
    LogContent []Sam_LogContents_LogContent
}

func (logContents *Sam_LogContents) GetEntityData() *types.CommonEntityData {
    logContents.EntityData.YFilter = logContents.YFilter
    logContents.EntityData.YangName = "log-contents"
    logContents.EntityData.BundleName = "cisco_ios_xr"
    logContents.EntityData.ParentYangName = "sam"
    logContents.EntityData.SegmentPath = "log-contents"
    logContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContents.EntityData.Children = make(map[string]types.YChild)
    logContents.EntityData.Children["log-content"] = types.YChild{"LogContent", nil}
    for i := range logContents.LogContent {
        logContents.EntityData.Children[types.GetSegmentPath(&logContents.LogContent[i])] = types.YChild{"LogContent", &logContents.LogContent[i]}
    }
    logContents.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(logContents.EntityData)
}

// Sam_LogContents_LogContent
// Number of lines for SAM log message
type Sam_LogContents_LogContent struct {
    EntityData types.CommonEntityData
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

func (logContent *Sam_LogContents_LogContent) GetEntityData() *types.CommonEntityData {
    logContent.EntityData.YFilter = logContent.YFilter
    logContent.EntityData.YangName = "log-content"
    logContent.EntityData.BundleName = "cisco_ios_xr"
    logContent.EntityData.ParentYangName = "log-contents"
    logContent.EntityData.SegmentPath = "log-content" + "[number-of-lines='" + fmt.Sprintf("%v", logContent.NumberOfLines) + "']"
    logContent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContent.EntityData.Children = make(map[string]types.YChild)
    logContent.EntityData.Children["logs"] = types.YChild{"Logs", nil}
    for i := range logContent.Logs {
        logContent.EntityData.Children[types.GetSegmentPath(&logContent.Logs[i])] = types.YChild{"Logs", &logContent.Logs[i]}
    }
    logContent.EntityData.Leafs = make(map[string]types.YLeaf)
    logContent.EntityData.Leafs["number-of-lines"] = types.YLeaf{"NumberOfLines", logContent.NumberOfLines}
    logContent.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", logContent.TotalEntries}
    logContent.EntityData.Leafs["entries-shown"] = types.YLeaf{"EntriesShown", logContent.EntriesShown}
    return &(logContent.EntityData)
}

// Sam_LogContents_LogContent_Logs
// SAM logs
type Sam_LogContents_LogContent_Logs struct {
    EntityData types.CommonEntityData
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

func (logs *Sam_LogContents_LogContent_Logs) GetEntityData() *types.CommonEntityData {
    logs.EntityData.YFilter = logs.YFilter
    logs.EntityData.YangName = "logs"
    logs.EntityData.BundleName = "cisco_ios_xr"
    logs.EntityData.ParentYangName = "log-content"
    logs.EntityData.SegmentPath = "logs"
    logs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logs.EntityData.Children = make(map[string]types.YChild)
    logs.EntityData.Leafs = make(map[string]types.YLeaf)
    logs.EntityData.Leafs["time"] = types.YLeaf{"Time", logs.Time}
    logs.EntityData.Leafs["code"] = types.YLeaf{"Code", logs.Code}
    logs.EntityData.Leafs["target-device"] = types.YLeaf{"TargetDevice", logs.TargetDevice}
    logs.EntityData.Leafs["index"] = types.YLeaf{"Index", logs.Index}
    logs.EntityData.Leafs["error"] = types.YLeaf{"Error", logs.Error}
    logs.EntityData.Leafs["issuer"] = types.YLeaf{"Issuer", logs.Issuer}
    logs.EntityData.Leafs["serial-no"] = types.YLeaf{"SerialNo", logs.SerialNo}
    logs.EntityData.Leafs["sam-table-index"] = types.YLeaf{"SamTableIndex", logs.SamTableIndex}
    logs.EntityData.Leafs["update-time"] = types.YLeaf{"UpdateTime", logs.UpdateTime}
    logs.EntityData.Leafs["source-device"] = types.YLeaf{"SourceDevice", logs.SourceDevice}
    logs.EntityData.Leafs["table"] = types.YLeaf{"Table", logs.Table}
    return &(logs.EntityData)
}

// Sam_Devices
// Certificate device table information
type Sam_Devices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate table device information. The type is slice of
    // Sam_Devices_Device.
    Device []Sam_Devices_Device
}

func (devices *Sam_Devices) GetEntityData() *types.CommonEntityData {
    devices.EntityData.YFilter = devices.YFilter
    devices.EntityData.YangName = "devices"
    devices.EntityData.BundleName = "cisco_ios_xr"
    devices.EntityData.ParentYangName = "sam"
    devices.EntityData.SegmentPath = "devices"
    devices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    devices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    devices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    devices.EntityData.Children = make(map[string]types.YChild)
    devices.EntityData.Children["device"] = types.YChild{"Device", nil}
    for i := range devices.Device {
        devices.EntityData.Children[types.GetSegmentPath(&devices.Device[i])] = types.YChild{"Device", &devices.Device[i]}
    }
    devices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(devices.EntityData)
}

// Sam_Devices_Device
// Certificate table device information
type Sam_Devices_Device struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify device name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    DeviceName interface{}

    // Certificate table information.
    Certificate Sam_Devices_Device_Certificate
}

func (device *Sam_Devices_Device) GetEntityData() *types.CommonEntityData {
    device.EntityData.YFilter = device.YFilter
    device.EntityData.YangName = "device"
    device.EntityData.BundleName = "cisco_ios_xr"
    device.EntityData.ParentYangName = "devices"
    device.EntityData.SegmentPath = "device" + "[device-name='" + fmt.Sprintf("%v", device.DeviceName) + "']"
    device.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    device.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    device.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    device.EntityData.Children = make(map[string]types.YChild)
    device.EntityData.Children["certificate"] = types.YChild{"Certificate", &device.Certificate}
    device.EntityData.Leafs = make(map[string]types.YLeaf)
    device.EntityData.Leafs["device-name"] = types.YLeaf{"DeviceName", device.DeviceName}
    return &(device.EntityData)
}

// Sam_Devices_Device_Certificate
// Certificate table information
type Sam_Devices_Device_Certificate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate table brief information.
    Brief Sam_Devices_Device_Certificate_Brief

    // Certificate detail index table information.
    CertificateIndexes Sam_Devices_Device_Certificate_CertificateIndexes
}

func (certificate *Sam_Devices_Device_Certificate) GetEntityData() *types.CommonEntityData {
    certificate.EntityData.YFilter = certificate.YFilter
    certificate.EntityData.YangName = "certificate"
    certificate.EntityData.BundleName = "cisco_ios_xr"
    certificate.EntityData.ParentYangName = "device"
    certificate.EntityData.SegmentPath = "certificate"
    certificate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificate.EntityData.Children = make(map[string]types.YChild)
    certificate.EntityData.Children["brief"] = types.YChild{"Brief", &certificate.Brief}
    certificate.EntityData.Children["certificate-indexes"] = types.YChild{"CertificateIndexes", &certificate.CertificateIndexes}
    certificate.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(certificate.EntityData)
}

// Sam_Devices_Device_Certificate_Brief
// Certificate table brief information
type Sam_Devices_Device_Certificate_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate location. The type is string.
    Location interface{}

    // Certificate index. The type is interface{} with range: 0..65535.
    CertificateIndex interface{}

    // Certificate flags.
    CertificateFlags Sam_Devices_Device_Certificate_Brief_CertificateFlags
}

func (brief *Sam_Devices_Device_Certificate_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "certificate"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Children["certificate-flags"] = types.YChild{"CertificateFlags", &brief.CertificateFlags}
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["location"] = types.YLeaf{"Location", brief.Location}
    brief.EntityData.Leafs["certificate-index"] = types.YLeaf{"CertificateIndex", brief.CertificateIndex}
    return &(brief.EntityData)
}

// Sam_Devices_Device_Certificate_Brief_CertificateFlags
// Certificate flags
type Sam_Devices_Device_Certificate_Brief_CertificateFlags struct {
    EntityData types.CommonEntityData
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

func (certificateFlags *Sam_Devices_Device_Certificate_Brief_CertificateFlags) GetEntityData() *types.CommonEntityData {
    certificateFlags.EntityData.YFilter = certificateFlags.YFilter
    certificateFlags.EntityData.YangName = "certificate-flags"
    certificateFlags.EntityData.BundleName = "cisco_ios_xr"
    certificateFlags.EntityData.ParentYangName = "brief"
    certificateFlags.EntityData.SegmentPath = "certificate-flags"
    certificateFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateFlags.EntityData.Children = make(map[string]types.YChild)
    certificateFlags.EntityData.Leafs = make(map[string]types.YLeaf)
    certificateFlags.EntityData.Leafs["is-trusted"] = types.YLeaf{"IsTrusted", certificateFlags.IsTrusted}
    certificateFlags.EntityData.Leafs["is-revoked"] = types.YLeaf{"IsRevoked", certificateFlags.IsRevoked}
    certificateFlags.EntityData.Leafs["is-expired"] = types.YLeaf{"IsExpired", certificateFlags.IsExpired}
    certificateFlags.EntityData.Leafs["is-validated"] = types.YLeaf{"IsValidated", certificateFlags.IsValidated}
    return &(certificateFlags.EntityData)
}

// Sam_Devices_Device_Certificate_CertificateIndexes
// Certificate detail index table information
type Sam_Devices_Device_Certificate_CertificateIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate detail index information. The type is slice of
    // Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex.
    CertificateIndex []Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex
}

func (certificateIndexes *Sam_Devices_Device_Certificate_CertificateIndexes) GetEntityData() *types.CommonEntityData {
    certificateIndexes.EntityData.YFilter = certificateIndexes.YFilter
    certificateIndexes.EntityData.YangName = "certificate-indexes"
    certificateIndexes.EntityData.BundleName = "cisco_ios_xr"
    certificateIndexes.EntityData.ParentYangName = "certificate"
    certificateIndexes.EntityData.SegmentPath = "certificate-indexes"
    certificateIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateIndexes.EntityData.Children = make(map[string]types.YChild)
    certificateIndexes.EntityData.Children["certificate-index"] = types.YChild{"CertificateIndex", nil}
    for i := range certificateIndexes.CertificateIndex {
        certificateIndexes.EntityData.Children[types.GetSegmentPath(&certificateIndexes.CertificateIndex[i])] = types.YChild{"CertificateIndex", &certificateIndexes.CertificateIndex[i]}
    }
    certificateIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(certificateIndexes.EntityData)
}

// Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex
// Certificate detail index information
type Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify certificate index. The type is interface{}
    // with range: -2147483648..2147483647.
    Index interface{}

    // Certificate table detail information.
    Detail Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail
}

func (certificateIndex *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex) GetEntityData() *types.CommonEntityData {
    certificateIndex.EntityData.YFilter = certificateIndex.YFilter
    certificateIndex.EntityData.YangName = "certificate-index"
    certificateIndex.EntityData.BundleName = "cisco_ios_xr"
    certificateIndex.EntityData.ParentYangName = "certificate-indexes"
    certificateIndex.EntityData.SegmentPath = "certificate-index" + "[index='" + fmt.Sprintf("%v", certificateIndex.Index) + "']"
    certificateIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateIndex.EntityData.Children = make(map[string]types.YChild)
    certificateIndex.EntityData.Children["detail"] = types.YChild{"Detail", &certificateIndex.Detail}
    certificateIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    certificateIndex.EntityData.Leafs["index"] = types.YLeaf{"Index", certificateIndex.Index}
    return &(certificateIndex.EntityData)
}

// Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail
// Certificate table detail information
type Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate location. The type is string.
    Location interface{}

    // Certificate index. The type is interface{} with range: 0..65535.
    CertificateIndex interface{}

    // Certificate flags.
    CertificateFlags Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags
}

func (detail *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "certificate-index"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["certificate-flags"] = types.YChild{"CertificateFlags", &detail.CertificateFlags}
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["location"] = types.YLeaf{"Location", detail.Location}
    detail.EntityData.Leafs["certificate-index"] = types.YLeaf{"CertificateIndex", detail.CertificateIndex}
    return &(detail.EntityData)
}

// Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags
// Certificate flags
type Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags struct {
    EntityData types.CommonEntityData
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

func (certificateFlags *Sam_Devices_Device_Certificate_CertificateIndexes_CertificateIndex_Detail_CertificateFlags) GetEntityData() *types.CommonEntityData {
    certificateFlags.EntityData.YFilter = certificateFlags.YFilter
    certificateFlags.EntityData.YangName = "certificate-flags"
    certificateFlags.EntityData.BundleName = "cisco_ios_xr"
    certificateFlags.EntityData.ParentYangName = "detail"
    certificateFlags.EntityData.SegmentPath = "certificate-flags"
    certificateFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateFlags.EntityData.Children = make(map[string]types.YChild)
    certificateFlags.EntityData.Leafs = make(map[string]types.YLeaf)
    certificateFlags.EntityData.Leafs["is-trusted"] = types.YLeaf{"IsTrusted", certificateFlags.IsTrusted}
    certificateFlags.EntityData.Leafs["is-revoked"] = types.YLeaf{"IsRevoked", certificateFlags.IsRevoked}
    certificateFlags.EntityData.Leafs["is-expired"] = types.YLeaf{"IsExpired", certificateFlags.IsExpired}
    certificateFlags.EntityData.Leafs["is-validated"] = types.YLeaf{"IsValidated", certificateFlags.IsValidated}
    return &(certificateFlags.EntityData)
}

// Sam_Packages
// SAM certificate information  package
type Sam_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SAM certificate information for a specific package. The type is slice of
    // Sam_Packages_Package_.
    Package_ []Sam_Packages_Package
}

func (packages *Sam_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "sam"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = make(map[string]types.YChild)
    packages.EntityData.Children["package"] = types.YChild{"Package_", nil}
    for i := range packages.Package_ {
        packages.EntityData.Children[types.GetSegmentPath(&packages.Package_[i])] = types.YChild{"Package_", &packages.Package_[i]}
    }
    packages.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(packages.EntityData)
}

// Sam_Packages_Package
// SAM certificate information for a specific
// package
type Sam_Packages_Package struct {
    EntityData types.CommonEntityData
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

func (self *Sam_Packages_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "packages"
    self.EntityData.SegmentPath = "package" + "[package-name='" + fmt.Sprintf("%v", self.PackageName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["certificate-flags"] = types.YChild{"CertificateFlags", &self.CertificateFlags}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["package-name"] = types.YLeaf{"PackageName", self.PackageName}
    self.EntityData.Leafs["location"] = types.YLeaf{"Location", self.Location}
    self.EntityData.Leafs["certificate-index"] = types.YLeaf{"CertificateIndex", self.CertificateIndex}
    return &(self.EntityData)
}

// Sam_Packages_Package_CertificateFlags
// Certificate flags
type Sam_Packages_Package_CertificateFlags struct {
    EntityData types.CommonEntityData
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

func (certificateFlags *Sam_Packages_Package_CertificateFlags) GetEntityData() *types.CommonEntityData {
    certificateFlags.EntityData.YFilter = certificateFlags.YFilter
    certificateFlags.EntityData.YangName = "certificate-flags"
    certificateFlags.EntityData.BundleName = "cisco_ios_xr"
    certificateFlags.EntityData.ParentYangName = "package"
    certificateFlags.EntityData.SegmentPath = "certificate-flags"
    certificateFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateFlags.EntityData.Children = make(map[string]types.YChild)
    certificateFlags.EntityData.Leafs = make(map[string]types.YLeaf)
    certificateFlags.EntityData.Leafs["is-trusted"] = types.YLeaf{"IsTrusted", certificateFlags.IsTrusted}
    certificateFlags.EntityData.Leafs["is-revoked"] = types.YLeaf{"IsRevoked", certificateFlags.IsRevoked}
    certificateFlags.EntityData.Leafs["is-expired"] = types.YLeaf{"IsExpired", certificateFlags.IsExpired}
    certificateFlags.EntityData.Leafs["is-validated"] = types.YLeaf{"IsValidated", certificateFlags.IsValidated}
    return &(certificateFlags.EntityData)
}

// Sam_CertificateRevocations
// Certificate revocation list index table
// information
type Sam_CertificateRevocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate revocation list index information. The type is slice of
    // Sam_CertificateRevocations_CertificateRevocation.
    CertificateRevocation []Sam_CertificateRevocations_CertificateRevocation
}

func (certificateRevocations *Sam_CertificateRevocations) GetEntityData() *types.CommonEntityData {
    certificateRevocations.EntityData.YFilter = certificateRevocations.YFilter
    certificateRevocations.EntityData.YangName = "certificate-revocations"
    certificateRevocations.EntityData.BundleName = "cisco_ios_xr"
    certificateRevocations.EntityData.ParentYangName = "sam"
    certificateRevocations.EntityData.SegmentPath = "certificate-revocations"
    certificateRevocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateRevocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateRevocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateRevocations.EntityData.Children = make(map[string]types.YChild)
    certificateRevocations.EntityData.Children["certificate-revocation"] = types.YChild{"CertificateRevocation", nil}
    for i := range certificateRevocations.CertificateRevocation {
        certificateRevocations.EntityData.Children[types.GetSegmentPath(&certificateRevocations.CertificateRevocation[i])] = types.YChild{"CertificateRevocation", &certificateRevocations.CertificateRevocation[i]}
    }
    certificateRevocations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(certificateRevocations.EntityData)
}

// Sam_CertificateRevocations_CertificateRevocation
// Certificate revocation list index information
type Sam_CertificateRevocations_CertificateRevocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. CRL index. The type is interface{} with range:
    // -2147483648..2147483647.
    CrlIndex interface{}

    // Certificate revocation list detail information.
    CertificateRevocationListDetail Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail
}

func (certificateRevocation *Sam_CertificateRevocations_CertificateRevocation) GetEntityData() *types.CommonEntityData {
    certificateRevocation.EntityData.YFilter = certificateRevocation.YFilter
    certificateRevocation.EntityData.YangName = "certificate-revocation"
    certificateRevocation.EntityData.BundleName = "cisco_ios_xr"
    certificateRevocation.EntityData.ParentYangName = "certificate-revocations"
    certificateRevocation.EntityData.SegmentPath = "certificate-revocation" + "[crl-index='" + fmt.Sprintf("%v", certificateRevocation.CrlIndex) + "']"
    certificateRevocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateRevocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateRevocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateRevocation.EntityData.Children = make(map[string]types.YChild)
    certificateRevocation.EntityData.Children["certificate-revocation-list-detail"] = types.YChild{"CertificateRevocationListDetail", &certificateRevocation.CertificateRevocationListDetail}
    certificateRevocation.EntityData.Leafs = make(map[string]types.YLeaf)
    certificateRevocation.EntityData.Leafs["crl-index"] = types.YLeaf{"CrlIndex", certificateRevocation.CrlIndex}
    return &(certificateRevocation.EntityData)
}

// Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail
// Certificate revocation list detail information
type Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CRL index. The type is interface{} with range: 0..65535.
    CrlIndex interface{}

    // Updated time of CRL is displayed. The type is string.
    Updates interface{}

    // Issuer name.
    Issuer Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer
}

func (certificateRevocationListDetail *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail) GetEntityData() *types.CommonEntityData {
    certificateRevocationListDetail.EntityData.YFilter = certificateRevocationListDetail.YFilter
    certificateRevocationListDetail.EntityData.YangName = "certificate-revocation-list-detail"
    certificateRevocationListDetail.EntityData.BundleName = "cisco_ios_xr"
    certificateRevocationListDetail.EntityData.ParentYangName = "certificate-revocation"
    certificateRevocationListDetail.EntityData.SegmentPath = "certificate-revocation-list-detail"
    certificateRevocationListDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateRevocationListDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateRevocationListDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateRevocationListDetail.EntityData.Children = make(map[string]types.YChild)
    certificateRevocationListDetail.EntityData.Children["issuer"] = types.YChild{"Issuer", &certificateRevocationListDetail.Issuer}
    certificateRevocationListDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    certificateRevocationListDetail.EntityData.Leafs["crl-index"] = types.YLeaf{"CrlIndex", certificateRevocationListDetail.CrlIndex}
    certificateRevocationListDetail.EntityData.Leafs["updates"] = types.YLeaf{"Updates", certificateRevocationListDetail.Updates}
    return &(certificateRevocationListDetail.EntityData)
}

// Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer
// Issuer name
type Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Common name. The type is string.
    CommonName interface{}

    // Organization. The type is string.
    Organization interface{}

    // Country. The type is string.
    Country interface{}
}

func (issuer *Sam_CertificateRevocations_CertificateRevocation_CertificateRevocationListDetail_Issuer) GetEntityData() *types.CommonEntityData {
    issuer.EntityData.YFilter = issuer.YFilter
    issuer.EntityData.YangName = "issuer"
    issuer.EntityData.BundleName = "cisco_ios_xr"
    issuer.EntityData.ParentYangName = "certificate-revocation-list-detail"
    issuer.EntityData.SegmentPath = "issuer"
    issuer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issuer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issuer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issuer.EntityData.Children = make(map[string]types.YChild)
    issuer.EntityData.Leafs = make(map[string]types.YLeaf)
    issuer.EntityData.Leafs["common-name"] = types.YLeaf{"CommonName", issuer.CommonName}
    issuer.EntityData.Leafs["organization"] = types.YLeaf{"Organization", issuer.Organization}
    issuer.EntityData.Leafs["country"] = types.YLeaf{"Country", issuer.Country}
    return &(issuer.EntityData)
}

// Sam_CertificateRevocationListSummary
// Certificate revocation list summary information 
type Sam_CertificateRevocationListSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CRL index. The type is interface{} with range: 0..65535.
    CrlIndex interface{}

    // Updated time of CRL is displayed. The type is string.
    Updates interface{}

    // Issuer name.
    Issuer Sam_CertificateRevocationListSummary_Issuer
}

func (certificateRevocationListSummary *Sam_CertificateRevocationListSummary) GetEntityData() *types.CommonEntityData {
    certificateRevocationListSummary.EntityData.YFilter = certificateRevocationListSummary.YFilter
    certificateRevocationListSummary.EntityData.YangName = "certificate-revocation-list-summary"
    certificateRevocationListSummary.EntityData.BundleName = "cisco_ios_xr"
    certificateRevocationListSummary.EntityData.ParentYangName = "sam"
    certificateRevocationListSummary.EntityData.SegmentPath = "certificate-revocation-list-summary"
    certificateRevocationListSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificateRevocationListSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificateRevocationListSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificateRevocationListSummary.EntityData.Children = make(map[string]types.YChild)
    certificateRevocationListSummary.EntityData.Children["issuer"] = types.YChild{"Issuer", &certificateRevocationListSummary.Issuer}
    certificateRevocationListSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    certificateRevocationListSummary.EntityData.Leafs["crl-index"] = types.YLeaf{"CrlIndex", certificateRevocationListSummary.CrlIndex}
    certificateRevocationListSummary.EntityData.Leafs["updates"] = types.YLeaf{"Updates", certificateRevocationListSummary.Updates}
    return &(certificateRevocationListSummary.EntityData)
}

// Sam_CertificateRevocationListSummary_Issuer
// Issuer name
type Sam_CertificateRevocationListSummary_Issuer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Common name. The type is string.
    CommonName interface{}

    // Organization. The type is string.
    Organization interface{}

    // Country. The type is string.
    Country interface{}
}

func (issuer *Sam_CertificateRevocationListSummary_Issuer) GetEntityData() *types.CommonEntityData {
    issuer.EntityData.YFilter = issuer.YFilter
    issuer.EntityData.YangName = "issuer"
    issuer.EntityData.BundleName = "cisco_ios_xr"
    issuer.EntityData.ParentYangName = "certificate-revocation-list-summary"
    issuer.EntityData.SegmentPath = "issuer"
    issuer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issuer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issuer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issuer.EntityData.Children = make(map[string]types.YChild)
    issuer.EntityData.Leafs = make(map[string]types.YLeaf)
    issuer.EntityData.Leafs["common-name"] = types.YLeaf{"CommonName", issuer.CommonName}
    issuer.EntityData.Leafs["organization"] = types.YLeaf{"Organization", issuer.Organization}
    issuer.EntityData.Leafs["country"] = types.YLeaf{"Country", issuer.Country}
    return &(issuer.EntityData)
}

