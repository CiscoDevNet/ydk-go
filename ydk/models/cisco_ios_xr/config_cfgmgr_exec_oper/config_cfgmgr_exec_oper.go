// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-cfgmgr-exec package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cfg-hist-gl: Configuration History Global path information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package config_cfgmgr_exec_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package config_cfgmgr_exec_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-config-cfgmgr-exec-oper cfg-hist-gl}", reflect.TypeOf(CfgHistGl{}))
    ydk.RegisterEntity("Cisco-IOS-XR-config-cfgmgr-exec-oper:cfg-hist-gl", reflect.TypeOf(CfgHistGl{}))
}

// HistRecord represents Possible types of history
type HistRecord string

const (
    // All history
    HistRecord_cfghist_bag_record_all HistRecord = "cfghist-bag-record-all"

    // Alarm history
    HistRecord_cfghist_bag_record_alarm HistRecord = "cfghist-bag-record-alarm"

    // CfgCheck history
    HistRecord_cfghist_bag_record_cfs_check HistRecord = "cfghist-bag-record-cfs-check"

    // Commit history
    HistRecord_cfghist_bag_record_commit HistRecord = "cfghist-bag-record-commit"

    // OIR history
    HistRecord_cfghist_bag_record_oir HistRecord = "cfghist-bag-record-oir"

    // Shutdown history
    HistRecord_cfghist_bag_record_shutdown HistRecord = "cfghist-bag-record-shutdown"

    // Bootup history
    HistRecord_cfghist_bag_record_startup HistRecord = "cfghist-bag-record-startup"

    // Backup history
    HistRecord_cfghist_bag_record_backup HistRecord = "cfghist-bag-record-backup"

    // Rebase history
    HistRecord_cfghist_bag_record_rebase HistRecord = "cfghist-bag-record-rebase"

    // Last history
    HistRecord_cfghist_bag_record_last HistRecord = "cfghist-bag-record-last"
)

// CfgHistGl
// Configuration History Global path information
type CfgHistGl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // History summary information for a specific type of history. The type is
    // slice of CfgHistGl_RecordType.
    RecordType []CfgHistGl_RecordType
}

func (cfgHistGl *CfgHistGl) GetFilter() yfilter.YFilter { return cfgHistGl.YFilter }

func (cfgHistGl *CfgHistGl) SetFilter(yf yfilter.YFilter) { cfgHistGl.YFilter = yf }

func (cfgHistGl *CfgHistGl) GetGoName(yname string) string {
    if yname == "record-type" { return "RecordType" }
    return ""
}

func (cfgHistGl *CfgHistGl) GetSegmentPath() string {
    return "Cisco-IOS-XR-config-cfgmgr-exec-oper:cfg-hist-gl"
}

func (cfgHistGl *CfgHistGl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "record-type" {
        for _, c := range cfgHistGl.RecordType {
            if cfgHistGl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CfgHistGl_RecordType{}
        cfgHistGl.RecordType = append(cfgHistGl.RecordType, child)
        return &cfgHistGl.RecordType[len(cfgHistGl.RecordType)-1]
    }
    return nil
}

func (cfgHistGl *CfgHistGl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cfgHistGl.RecordType {
        children[cfgHistGl.RecordType[i].GetSegmentPath()] = &cfgHistGl.RecordType[i]
    }
    return children
}

func (cfgHistGl *CfgHistGl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cfgHistGl *CfgHistGl) GetBundleName() string { return "cisco_ios_xr" }

func (cfgHistGl *CfgHistGl) GetYangName() string { return "cfg-hist-gl" }

func (cfgHistGl *CfgHistGl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cfgHistGl *CfgHistGl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cfgHistGl *CfgHistGl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cfgHistGl *CfgHistGl) SetParent(parent types.Entity) { cfgHistGl.parent = parent }

func (cfgHistGl *CfgHistGl) GetParent() types.Entity { return cfgHistGl.parent }

func (cfgHistGl *CfgHistGl) GetParentYangName() string { return "Cisco-IOS-XR-config-cfgmgr-exec-oper" }

// CfgHistGl_RecordType
// History summary information for a specific type
// of history
type CfgHistGl_RecordType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Record type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    RecordType interface{}

    // History summary information for a specific type of history. The type is
    // slice of CfgHistGl_RecordType_Record.
    Record []CfgHistGl_RecordType_Record
}

func (recordType *CfgHistGl_RecordType) GetFilter() yfilter.YFilter { return recordType.YFilter }

func (recordType *CfgHistGl_RecordType) SetFilter(yf yfilter.YFilter) { recordType.YFilter = yf }

func (recordType *CfgHistGl_RecordType) GetGoName(yname string) string {
    if yname == "record-type" { return "RecordType" }
    if yname == "record" { return "Record" }
    return ""
}

func (recordType *CfgHistGl_RecordType) GetSegmentPath() string {
    return "record-type" + "[record-type='" + fmt.Sprintf("%v", recordType.RecordType) + "']"
}

func (recordType *CfgHistGl_RecordType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "record" {
        for _, c := range recordType.Record {
            if recordType.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CfgHistGl_RecordType_Record{}
        recordType.Record = append(recordType.Record, child)
        return &recordType.Record[len(recordType.Record)-1]
    }
    return nil
}

func (recordType *CfgHistGl_RecordType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range recordType.Record {
        children[recordType.Record[i].GetSegmentPath()] = &recordType.Record[i]
    }
    return children
}

func (recordType *CfgHistGl_RecordType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["record-type"] = recordType.RecordType
    return leafs
}

func (recordType *CfgHistGl_RecordType) GetBundleName() string { return "cisco_ios_xr" }

func (recordType *CfgHistGl_RecordType) GetYangName() string { return "record-type" }

func (recordType *CfgHistGl_RecordType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (recordType *CfgHistGl_RecordType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (recordType *CfgHistGl_RecordType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (recordType *CfgHistGl_RecordType) SetParent(parent types.Entity) { recordType.parent = parent }

func (recordType *CfgHistGl_RecordType) GetParent() types.Entity { return recordType.parent }

func (recordType *CfgHistGl_RecordType) GetParentYangName() string { return "cfg-hist-gl" }

// CfgHistGl_RecordType_Record
// History summary information for a specific type
// of history
type CfgHistGl_RecordType_Record struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Record. The type is interface{} with range:
    // -2147483648..2147483647.
    Record interface{}

    // Time stamp for the history. The type is interface{} with range:
    // 0..4294967295.
    Timestamp interface{}

    // Record type. The type is HistRecord.
    RecordType interface{}

    // Content of the history.
    Info CfgHistGl_RecordType_Record_Info
}

func (record *CfgHistGl_RecordType_Record) GetFilter() yfilter.YFilter { return record.YFilter }

func (record *CfgHistGl_RecordType_Record) SetFilter(yf yfilter.YFilter) { record.YFilter = yf }

func (record *CfgHistGl_RecordType_Record) GetGoName(yname string) string {
    if yname == "record" { return "Record" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "record-type" { return "RecordType" }
    if yname == "info" { return "Info" }
    return ""
}

func (record *CfgHistGl_RecordType_Record) GetSegmentPath() string {
    return "record" + "[record='" + fmt.Sprintf("%v", record.Record) + "']"
}

func (record *CfgHistGl_RecordType_Record) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "info" {
        return &record.Info
    }
    return nil
}

func (record *CfgHistGl_RecordType_Record) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["info"] = &record.Info
    return children
}

func (record *CfgHistGl_RecordType_Record) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["record"] = record.Record
    leafs["timestamp"] = record.Timestamp
    leafs["record-type"] = record.RecordType
    return leafs
}

func (record *CfgHistGl_RecordType_Record) GetBundleName() string { return "cisco_ios_xr" }

func (record *CfgHistGl_RecordType_Record) GetYangName() string { return "record" }

func (record *CfgHistGl_RecordType_Record) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (record *CfgHistGl_RecordType_Record) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (record *CfgHistGl_RecordType_Record) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (record *CfgHistGl_RecordType_Record) SetParent(parent types.Entity) { record.parent = parent }

func (record *CfgHistGl_RecordType_Record) GetParent() types.Entity { return record.parent }

func (record *CfgHistGl_RecordType_Record) GetParentYangName() string { return "record-type" }

// CfgHistGl_RecordType_Record_Info
// Content of the history
type CfgHistGl_RecordType_Record_Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is HistRecord.
    Type interface{}

    // B. The type is interface{} with range: 0..4294967295.
    A interface{}

    // alarm info.
    AlarmInfo CfgHistGl_RecordType_Record_Info_AlarmInfo

    // cfscheck info.
    CfscheckInfo CfgHistGl_RecordType_Record_Info_CfscheckInfo

    // commit info.
    CommitInfo CfgHistGl_RecordType_Record_Info_CommitInfo

    // oir info.
    OirInfo CfgHistGl_RecordType_Record_Info_OirInfo

    // shutdown info.
    ShutdownInfo CfgHistGl_RecordType_Record_Info_ShutdownInfo

    // startup info.
    StartupInfo CfgHistGl_RecordType_Record_Info_StartupInfo

    // backup info.
    BackupInfo CfgHistGl_RecordType_Record_Info_BackupInfo
}

func (info *CfgHistGl_RecordType_Record_Info) GetFilter() yfilter.YFilter { return info.YFilter }

func (info *CfgHistGl_RecordType_Record_Info) SetFilter(yf yfilter.YFilter) { info.YFilter = yf }

func (info *CfgHistGl_RecordType_Record_Info) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "a" { return "A" }
    if yname == "alarm-info" { return "AlarmInfo" }
    if yname == "cfscheck-info" { return "CfscheckInfo" }
    if yname == "commit-info" { return "CommitInfo" }
    if yname == "oir-info" { return "OirInfo" }
    if yname == "shutdown-info" { return "ShutdownInfo" }
    if yname == "startup-info" { return "StartupInfo" }
    if yname == "backup-info" { return "BackupInfo" }
    return ""
}

func (info *CfgHistGl_RecordType_Record_Info) GetSegmentPath() string {
    return "info"
}

func (info *CfgHistGl_RecordType_Record_Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        return &info.AlarmInfo
    }
    if childYangName == "cfscheck-info" {
        return &info.CfscheckInfo
    }
    if childYangName == "commit-info" {
        return &info.CommitInfo
    }
    if childYangName == "oir-info" {
        return &info.OirInfo
    }
    if childYangName == "shutdown-info" {
        return &info.ShutdownInfo
    }
    if childYangName == "startup-info" {
        return &info.StartupInfo
    }
    if childYangName == "backup-info" {
        return &info.BackupInfo
    }
    return nil
}

func (info *CfgHistGl_RecordType_Record_Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["alarm-info"] = &info.AlarmInfo
    children["cfscheck-info"] = &info.CfscheckInfo
    children["commit-info"] = &info.CommitInfo
    children["oir-info"] = &info.OirInfo
    children["shutdown-info"] = &info.ShutdownInfo
    children["startup-info"] = &info.StartupInfo
    children["backup-info"] = &info.BackupInfo
    return children
}

func (info *CfgHistGl_RecordType_Record_Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = info.Type
    leafs["a"] = info.A
    return leafs
}

func (info *CfgHistGl_RecordType_Record_Info) GetBundleName() string { return "cisco_ios_xr" }

func (info *CfgHistGl_RecordType_Record_Info) GetYangName() string { return "info" }

func (info *CfgHistGl_RecordType_Record_Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (info *CfgHistGl_RecordType_Record_Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (info *CfgHistGl_RecordType_Record_Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (info *CfgHistGl_RecordType_Record_Info) SetParent(parent types.Entity) { info.parent = parent }

func (info *CfgHistGl_RecordType_Record_Info) GetParent() types.Entity { return info.parent }

func (info *CfgHistGl_RecordType_Record_Info) GetParentYangName() string { return "record" }

// CfgHistGl_RecordType_Record_Info_AlarmInfo
// alarm info
type CfgHistGl_RecordType_Record_Info_AlarmInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // State. The type is string.
    State interface{}

    // Where. The type is string.
    Where interface{}
}

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "where" { return "Where" }
    return ""
}

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = alarmInfo.State
    leafs["where"] = alarmInfo.Where
    return leafs
}

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetParentYangName() string { return "info" }

// CfgHistGl_RecordType_Record_Info_CfscheckInfo
// cfscheck info
type CfgHistGl_RecordType_Record_Info_CfscheckInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UserId. The type is string.
    UserId interface{}

    // Line. The type is string.
    Line interface{}
}

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetFilter() yfilter.YFilter { return cfscheckInfo.YFilter }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) SetFilter(yf yfilter.YFilter) { cfscheckInfo.YFilter = yf }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetGoName(yname string) string {
    if yname == "user-id" { return "UserId" }
    if yname == "line" { return "Line" }
    return ""
}

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetSegmentPath() string {
    return "cfscheck-info"
}

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["user-id"] = cfscheckInfo.UserId
    leafs["line"] = cfscheckInfo.Line
    return leafs
}

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetBundleName() string { return "cisco_ios_xr" }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetYangName() string { return "cfscheck-info" }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) SetParent(parent types.Entity) { cfscheckInfo.parent = parent }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetParent() types.Entity { return cfscheckInfo.parent }

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetParentYangName() string { return "info" }

// CfgHistGl_RecordType_Record_Info_CommitInfo
// commit info
type CfgHistGl_RecordType_Record_Info_CommitInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CommitId. The type is string.
    CommitId interface{}

    // UserId. The type is string.
    UserId interface{}

    // Line. The type is string.
    Line interface{}

    // Client name. The type is string.
    ClientName interface{}

    // Label. The type is string.
    Label interface{}

    // Comment. The type is string.
    Comment interface{}
}

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetFilter() yfilter.YFilter { return commitInfo.YFilter }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) SetFilter(yf yfilter.YFilter) { commitInfo.YFilter = yf }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetGoName(yname string) string {
    if yname == "commit-id" { return "CommitId" }
    if yname == "user-id" { return "UserId" }
    if yname == "line" { return "Line" }
    if yname == "client-name" { return "ClientName" }
    if yname == "label" { return "Label" }
    if yname == "comment" { return "Comment" }
    return ""
}

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetSegmentPath() string {
    return "commit-info"
}

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["commit-id"] = commitInfo.CommitId
    leafs["user-id"] = commitInfo.UserId
    leafs["line"] = commitInfo.Line
    leafs["client-name"] = commitInfo.ClientName
    leafs["label"] = commitInfo.Label
    leafs["comment"] = commitInfo.Comment
    return leafs
}

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetBundleName() string { return "cisco_ios_xr" }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetYangName() string { return "commit-info" }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) SetParent(parent types.Entity) { commitInfo.parent = parent }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetParent() types.Entity { return commitInfo.parent }

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetParentYangName() string { return "info" }

// CfgHistGl_RecordType_Record_Info_OirInfo
// oir info
type CfgHistGl_RecordType_Record_Info_OirInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config Type. The type is string.
    ConfigType interface{}

    // Operation. The type is string.
    Operation interface{}

    // Config Name. The type is string.
    ConfigName interface{}
}

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetFilter() yfilter.YFilter { return oirInfo.YFilter }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) SetFilter(yf yfilter.YFilter) { oirInfo.YFilter = yf }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetGoName(yname string) string {
    if yname == "config-type" { return "ConfigType" }
    if yname == "operation" { return "Operation" }
    if yname == "config-name" { return "ConfigName" }
    return ""
}

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetSegmentPath() string {
    return "oir-info"
}

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["config-type"] = oirInfo.ConfigType
    leafs["operation"] = oirInfo.Operation
    leafs["config-name"] = oirInfo.ConfigName
    return leafs
}

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetBundleName() string { return "cisco_ios_xr" }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetYangName() string { return "oir-info" }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) SetParent(parent types.Entity) { oirInfo.parent = parent }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetParent() types.Entity { return oirInfo.parent }

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetParentYangName() string { return "info" }

// CfgHistGl_RecordType_Record_Info_ShutdownInfo
// shutdown info
type CfgHistGl_RecordType_Record_Info_ShutdownInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Comment. The type is string.
    Comment interface{}
}

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetFilter() yfilter.YFilter { return shutdownInfo.YFilter }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) SetFilter(yf yfilter.YFilter) { shutdownInfo.YFilter = yf }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetGoName(yname string) string {
    if yname == "comment" { return "Comment" }
    return ""
}

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetSegmentPath() string {
    return "shutdown-info"
}

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["comment"] = shutdownInfo.Comment
    return leafs
}

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetBundleName() string { return "cisco_ios_xr" }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetYangName() string { return "shutdown-info" }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) SetParent(parent types.Entity) { shutdownInfo.parent = parent }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetParent() types.Entity { return shutdownInfo.parent }

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetParentYangName() string { return "info" }

// CfgHistGl_RecordType_Record_Info_StartupInfo
// startup info
type CfgHistGl_RecordType_Record_Info_StartupInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // How Booted. The type is string.
    HowBooted interface{}

    // Boot Path. The type is string.
    BootPath interface{}
}

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetFilter() yfilter.YFilter { return startupInfo.YFilter }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) SetFilter(yf yfilter.YFilter) { startupInfo.YFilter = yf }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetGoName(yname string) string {
    if yname == "how-booted" { return "HowBooted" }
    if yname == "boot-path" { return "BootPath" }
    return ""
}

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetSegmentPath() string {
    return "startup-info"
}

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["how-booted"] = startupInfo.HowBooted
    leafs["boot-path"] = startupInfo.BootPath
    return leafs
}

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetBundleName() string { return "cisco_ios_xr" }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetYangName() string { return "startup-info" }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) SetParent(parent types.Entity) { startupInfo.parent = parent }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetParent() types.Entity { return startupInfo.parent }

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetParentYangName() string { return "info" }

// CfgHistGl_RecordType_Record_Info_BackupInfo
// backup info
type CfgHistGl_RecordType_Record_Info_BackupInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Comment. The type is string.
    Comment interface{}
}

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetFilter() yfilter.YFilter { return backupInfo.YFilter }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) SetFilter(yf yfilter.YFilter) { backupInfo.YFilter = yf }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetGoName(yname string) string {
    if yname == "comment" { return "Comment" }
    return ""
}

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetSegmentPath() string {
    return "backup-info"
}

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["comment"] = backupInfo.Comment
    return leafs
}

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetBundleName() string { return "cisco_ios_xr" }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetYangName() string { return "backup-info" }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) SetParent(parent types.Entity) { backupInfo.parent = parent }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetParent() types.Entity { return backupInfo.parent }

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetParentYangName() string { return "info" }

