// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-cfgmgr-exec package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cfg-hist-gl: Configuration History Global path information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // History summary information for a specific type of history. The type is
    // slice of CfgHistGl_RecordType.
    RecordType []*CfgHistGl_RecordType
}

func (cfgHistGl *CfgHistGl) GetEntityData() *types.CommonEntityData {
    cfgHistGl.EntityData.YFilter = cfgHistGl.YFilter
    cfgHistGl.EntityData.YangName = "cfg-hist-gl"
    cfgHistGl.EntityData.BundleName = "cisco_ios_xr"
    cfgHistGl.EntityData.ParentYangName = "Cisco-IOS-XR-config-cfgmgr-exec-oper"
    cfgHistGl.EntityData.SegmentPath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:cfg-hist-gl"
    cfgHistGl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cfgHistGl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cfgHistGl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cfgHistGl.EntityData.Children = types.NewOrderedMap()
    cfgHistGl.EntityData.Children.Append("record-type", types.YChild{"RecordType", nil})
    for i := range cfgHistGl.RecordType {
        cfgHistGl.EntityData.Children.Append(types.GetSegmentPath(cfgHistGl.RecordType[i]), types.YChild{"RecordType", cfgHistGl.RecordType[i]})
    }
    cfgHistGl.EntityData.Leafs = types.NewOrderedMap()

    cfgHistGl.EntityData.YListKeys = []string {}

    return &(cfgHistGl.EntityData)
}

// CfgHistGl_RecordType
// History summary information for a specific type
// of history
type CfgHistGl_RecordType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Record type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    RecordType interface{}

    // History summary information for a specific type of history. The type is
    // slice of CfgHistGl_RecordType_Record.
    Record []*CfgHistGl_RecordType_Record
}

func (recordType *CfgHistGl_RecordType) GetEntityData() *types.CommonEntityData {
    recordType.EntityData.YFilter = recordType.YFilter
    recordType.EntityData.YangName = "record-type"
    recordType.EntityData.BundleName = "cisco_ios_xr"
    recordType.EntityData.ParentYangName = "cfg-hist-gl"
    recordType.EntityData.SegmentPath = "record-type" + types.AddKeyToken(recordType.RecordType, "record-type")
    recordType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    recordType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    recordType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    recordType.EntityData.Children = types.NewOrderedMap()
    recordType.EntityData.Children.Append("record", types.YChild{"Record", nil})
    for i := range recordType.Record {
        recordType.EntityData.Children.Append(types.GetSegmentPath(recordType.Record[i]), types.YChild{"Record", recordType.Record[i]})
    }
    recordType.EntityData.Leafs = types.NewOrderedMap()
    recordType.EntityData.Leafs.Append("record-type", types.YLeaf{"RecordType", recordType.RecordType})

    recordType.EntityData.YListKeys = []string {"RecordType"}

    return &(recordType.EntityData)
}

// CfgHistGl_RecordType_Record
// History summary information for a specific type
// of history
type CfgHistGl_RecordType_Record struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Record. The type is interface{} with range:
    // 0..4294967295.
    Record interface{}

    // Time stamp for the history. The type is interface{} with range:
    // 0..4294967295.
    Timestamp interface{}

    // Record type. The type is HistRecord.
    RecordType interface{}

    // Content of the history.
    Info CfgHistGl_RecordType_Record_Info
}

func (record *CfgHistGl_RecordType_Record) GetEntityData() *types.CommonEntityData {
    record.EntityData.YFilter = record.YFilter
    record.EntityData.YangName = "record"
    record.EntityData.BundleName = "cisco_ios_xr"
    record.EntityData.ParentYangName = "record-type"
    record.EntityData.SegmentPath = "record" + types.AddKeyToken(record.Record, "record")
    record.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    record.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    record.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    record.EntityData.Children = types.NewOrderedMap()
    record.EntityData.Children.Append("info", types.YChild{"Info", &record.Info})
    record.EntityData.Leafs = types.NewOrderedMap()
    record.EntityData.Leafs.Append("record", types.YLeaf{"Record", record.Record})
    record.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", record.Timestamp})
    record.EntityData.Leafs.Append("record-type", types.YLeaf{"RecordType", record.RecordType})

    record.EntityData.YListKeys = []string {"Record"}

    return &(record.EntityData)
}

// CfgHistGl_RecordType_Record_Info
// Content of the history
type CfgHistGl_RecordType_Record_Info struct {
    EntityData types.CommonEntityData
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

func (info *CfgHistGl_RecordType_Record_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "record"
    info.EntityData.SegmentPath = "info"
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", &info.AlarmInfo})
    info.EntityData.Children.Append("cfscheck-info", types.YChild{"CfscheckInfo", &info.CfscheckInfo})
    info.EntityData.Children.Append("commit-info", types.YChild{"CommitInfo", &info.CommitInfo})
    info.EntityData.Children.Append("oir-info", types.YChild{"OirInfo", &info.OirInfo})
    info.EntityData.Children.Append("shutdown-info", types.YChild{"ShutdownInfo", &info.ShutdownInfo})
    info.EntityData.Children.Append("startup-info", types.YChild{"StartupInfo", &info.StartupInfo})
    info.EntityData.Children.Append("backup-info", types.YChild{"BackupInfo", &info.BackupInfo})
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("type", types.YLeaf{"Type", info.Type})
    info.EntityData.Leafs.Append("a", types.YLeaf{"A", info.A})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// CfgHistGl_RecordType_Record_Info_AlarmInfo
// alarm info
type CfgHistGl_RecordType_Record_Info_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State. The type is string.
    State interface{}

    // Where. The type is string.
    Where interface{}
}

func (alarmInfo *CfgHistGl_RecordType_Record_Info_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "info"
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", alarmInfo.State})
    alarmInfo.EntityData.Leafs.Append("where", types.YLeaf{"Where", alarmInfo.Where})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// CfgHistGl_RecordType_Record_Info_CfscheckInfo
// cfscheck info
type CfgHistGl_RecordType_Record_Info_CfscheckInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UserId. The type is string.
    UserId interface{}

    // Line. The type is string.
    Line interface{}
}

func (cfscheckInfo *CfgHistGl_RecordType_Record_Info_CfscheckInfo) GetEntityData() *types.CommonEntityData {
    cfscheckInfo.EntityData.YFilter = cfscheckInfo.YFilter
    cfscheckInfo.EntityData.YangName = "cfscheck-info"
    cfscheckInfo.EntityData.BundleName = "cisco_ios_xr"
    cfscheckInfo.EntityData.ParentYangName = "info"
    cfscheckInfo.EntityData.SegmentPath = "cfscheck-info"
    cfscheckInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cfscheckInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cfscheckInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cfscheckInfo.EntityData.Children = types.NewOrderedMap()
    cfscheckInfo.EntityData.Leafs = types.NewOrderedMap()
    cfscheckInfo.EntityData.Leafs.Append("user-id", types.YLeaf{"UserId", cfscheckInfo.UserId})
    cfscheckInfo.EntityData.Leafs.Append("line", types.YLeaf{"Line", cfscheckInfo.Line})

    cfscheckInfo.EntityData.YListKeys = []string {}

    return &(cfscheckInfo.EntityData)
}

// CfgHistGl_RecordType_Record_Info_CommitInfo
// commit info
type CfgHistGl_RecordType_Record_Info_CommitInfo struct {
    EntityData types.CommonEntityData
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

func (commitInfo *CfgHistGl_RecordType_Record_Info_CommitInfo) GetEntityData() *types.CommonEntityData {
    commitInfo.EntityData.YFilter = commitInfo.YFilter
    commitInfo.EntityData.YangName = "commit-info"
    commitInfo.EntityData.BundleName = "cisco_ios_xr"
    commitInfo.EntityData.ParentYangName = "info"
    commitInfo.EntityData.SegmentPath = "commit-info"
    commitInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commitInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commitInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commitInfo.EntityData.Children = types.NewOrderedMap()
    commitInfo.EntityData.Leafs = types.NewOrderedMap()
    commitInfo.EntityData.Leafs.Append("commit-id", types.YLeaf{"CommitId", commitInfo.CommitId})
    commitInfo.EntityData.Leafs.Append("user-id", types.YLeaf{"UserId", commitInfo.UserId})
    commitInfo.EntityData.Leafs.Append("line", types.YLeaf{"Line", commitInfo.Line})
    commitInfo.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", commitInfo.ClientName})
    commitInfo.EntityData.Leafs.Append("label", types.YLeaf{"Label", commitInfo.Label})
    commitInfo.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", commitInfo.Comment})

    commitInfo.EntityData.YListKeys = []string {}

    return &(commitInfo.EntityData)
}

// CfgHistGl_RecordType_Record_Info_OirInfo
// oir info
type CfgHistGl_RecordType_Record_Info_OirInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config Type. The type is string.
    ConfigType interface{}

    // Operation. The type is string.
    Operation interface{}

    // Config Name. The type is string.
    ConfigName interface{}
}

func (oirInfo *CfgHistGl_RecordType_Record_Info_OirInfo) GetEntityData() *types.CommonEntityData {
    oirInfo.EntityData.YFilter = oirInfo.YFilter
    oirInfo.EntityData.YangName = "oir-info"
    oirInfo.EntityData.BundleName = "cisco_ios_xr"
    oirInfo.EntityData.ParentYangName = "info"
    oirInfo.EntityData.SegmentPath = "oir-info"
    oirInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirInfo.EntityData.Children = types.NewOrderedMap()
    oirInfo.EntityData.Leafs = types.NewOrderedMap()
    oirInfo.EntityData.Leafs.Append("config-type", types.YLeaf{"ConfigType", oirInfo.ConfigType})
    oirInfo.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", oirInfo.Operation})
    oirInfo.EntityData.Leafs.Append("config-name", types.YLeaf{"ConfigName", oirInfo.ConfigName})

    oirInfo.EntityData.YListKeys = []string {}

    return &(oirInfo.EntityData)
}

// CfgHistGl_RecordType_Record_Info_ShutdownInfo
// shutdown info
type CfgHistGl_RecordType_Record_Info_ShutdownInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Comment. The type is string.
    Comment interface{}
}

func (shutdownInfo *CfgHistGl_RecordType_Record_Info_ShutdownInfo) GetEntityData() *types.CommonEntityData {
    shutdownInfo.EntityData.YFilter = shutdownInfo.YFilter
    shutdownInfo.EntityData.YangName = "shutdown-info"
    shutdownInfo.EntityData.BundleName = "cisco_ios_xr"
    shutdownInfo.EntityData.ParentYangName = "info"
    shutdownInfo.EntityData.SegmentPath = "shutdown-info"
    shutdownInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shutdownInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shutdownInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shutdownInfo.EntityData.Children = types.NewOrderedMap()
    shutdownInfo.EntityData.Leafs = types.NewOrderedMap()
    shutdownInfo.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", shutdownInfo.Comment})

    shutdownInfo.EntityData.YListKeys = []string {}

    return &(shutdownInfo.EntityData)
}

// CfgHistGl_RecordType_Record_Info_StartupInfo
// startup info
type CfgHistGl_RecordType_Record_Info_StartupInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // How Booted. The type is string.
    HowBooted interface{}

    // Boot Path. The type is string.
    BootPath interface{}
}

func (startupInfo *CfgHistGl_RecordType_Record_Info_StartupInfo) GetEntityData() *types.CommonEntityData {
    startupInfo.EntityData.YFilter = startupInfo.YFilter
    startupInfo.EntityData.YangName = "startup-info"
    startupInfo.EntityData.BundleName = "cisco_ios_xr"
    startupInfo.EntityData.ParentYangName = "info"
    startupInfo.EntityData.SegmentPath = "startup-info"
    startupInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startupInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startupInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startupInfo.EntityData.Children = types.NewOrderedMap()
    startupInfo.EntityData.Leafs = types.NewOrderedMap()
    startupInfo.EntityData.Leafs.Append("how-booted", types.YLeaf{"HowBooted", startupInfo.HowBooted})
    startupInfo.EntityData.Leafs.Append("boot-path", types.YLeaf{"BootPath", startupInfo.BootPath})

    startupInfo.EntityData.YListKeys = []string {}

    return &(startupInfo.EntityData)
}

// CfgHistGl_RecordType_Record_Info_BackupInfo
// backup info
type CfgHistGl_RecordType_Record_Info_BackupInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Comment. The type is string.
    Comment interface{}
}

func (backupInfo *CfgHistGl_RecordType_Record_Info_BackupInfo) GetEntityData() *types.CommonEntityData {
    backupInfo.EntityData.YFilter = backupInfo.YFilter
    backupInfo.EntityData.YangName = "backup-info"
    backupInfo.EntityData.BundleName = "cisco_ios_xr"
    backupInfo.EntityData.ParentYangName = "info"
    backupInfo.EntityData.SegmentPath = "backup-info"
    backupInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupInfo.EntityData.Children = types.NewOrderedMap()
    backupInfo.EntityData.Leafs = types.NewOrderedMap()
    backupInfo.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", backupInfo.Comment})

    backupInfo.EntityData.YListKeys = []string {}

    return &(backupInfo.EntityData)
}

