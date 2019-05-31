// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-cfgmgr-exec package operational data.
// 
// This module contains definitions
// for the following management objects:
//   config-manager: Show Configuration <*> Global path information
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-config-cfgmgr-exec-oper config-manager}", reflect.TypeOf(ConfigManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager", reflect.TypeOf(ConfigManager{}))
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

// ConfigManager
// Show Configuration <*> Global path information
type ConfigManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Configuration <*> Global path information.
    Global ConfigManager_Global
}

func (configManager *ConfigManager) GetEntityData() *types.CommonEntityData {
    configManager.EntityData.YFilter = configManager.YFilter
    configManager.EntityData.YangName = "config-manager"
    configManager.EntityData.BundleName = "cisco_ios_xr"
    configManager.EntityData.ParentYangName = "Cisco-IOS-XR-config-cfgmgr-exec-oper"
    configManager.EntityData.SegmentPath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager"
    configManager.EntityData.AbsolutePath = configManager.EntityData.SegmentPath
    configManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configManager.EntityData.Children = types.NewOrderedMap()
    configManager.EntityData.Children.Append("global", types.YChild{"Global", &configManager.Global})
    configManager.EntityData.Leafs = types.NewOrderedMap()

    configManager.EntityData.YListKeys = []string {}

    return &(configManager.EntityData)
}

// ConfigManager_Global
// Show Configuration <*> Global path information
type ConfigManager_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Configuration Commit <*> Global path information.
    ConfigCommit ConfigManager_Global_ConfigCommit

    // Show Configuration History <*> Global path information.
    HistoryTables ConfigManager_Global_HistoryTables

    // Show Configuration Session <*> Global path information.
    ConfigSession ConfigManager_Global_ConfigSession
}

func (global *ConfigManager_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "config-manager"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("config-commit", types.YChild{"ConfigCommit", &global.ConfigCommit})
    global.EntityData.Children.Append("history-tables", types.YChild{"HistoryTables", &global.HistoryTables})
    global.EntityData.Children.Append("config-session", types.YChild{"ConfigSession", &global.ConfigSession})
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// ConfigManager_Global_ConfigCommit
// Show Configuration Commit <*> Global path
// information
type ConfigManager_Global_ConfigCommit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Configuration Commit List Detail.
    Commits ConfigManager_Global_ConfigCommit_Commits
}

func (configCommit *ConfigManager_Global_ConfigCommit) GetEntityData() *types.CommonEntityData {
    configCommit.EntityData.YFilter = configCommit.YFilter
    configCommit.EntityData.YangName = "config-commit"
    configCommit.EntityData.BundleName = "cisco_ios_xr"
    configCommit.EntityData.ParentYangName = "global"
    configCommit.EntityData.SegmentPath = "config-commit"
    configCommit.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/" + configCommit.EntityData.SegmentPath
    configCommit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configCommit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configCommit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configCommit.EntityData.Children = types.NewOrderedMap()
    configCommit.EntityData.Children.Append("commits", types.YChild{"Commits", &configCommit.Commits})
    configCommit.EntityData.Leafs = types.NewOrderedMap()

    configCommit.EntityData.YListKeys = []string {}

    return &(configCommit.EntityData)
}

// ConfigManager_Global_ConfigCommit_Commits
// Show Configuration Commit List Detail
type ConfigManager_Global_ConfigCommit_Commits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Commit information for a specific commit entry. The type is slice of
    // ConfigManager_Global_ConfigCommit_Commits_Commit.
    Commit []*ConfigManager_Global_ConfigCommit_Commits_Commit
}

func (commits *ConfigManager_Global_ConfigCommit_Commits) GetEntityData() *types.CommonEntityData {
    commits.EntityData.YFilter = commits.YFilter
    commits.EntityData.YangName = "commits"
    commits.EntityData.BundleName = "cisco_ios_xr"
    commits.EntityData.ParentYangName = "config-commit"
    commits.EntityData.SegmentPath = "commits"
    commits.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/config-commit/" + commits.EntityData.SegmentPath
    commits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commits.EntityData.Children = types.NewOrderedMap()
    commits.EntityData.Children.Append("commit", types.YChild{"Commit", nil})
    for i := range commits.Commit {
        commits.EntityData.Children.Append(types.GetSegmentPath(commits.Commit[i]), types.YChild{"Commit", commits.Commit[i]})
    }
    commits.EntityData.Leafs = types.NewOrderedMap()

    commits.EntityData.YListKeys = []string {}

    return &(commits.EntityData)
}

// ConfigManager_Global_ConfigCommit_Commits_Commit
// Commit information for a specific commit
// entry
type ConfigManager_Global_ConfigCommit_Commits_Commit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Each Session Details. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Commit interface{}

    // Timestamp. The type is string.
    Timestamp interface{}

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

func (commit *ConfigManager_Global_ConfigCommit_Commits_Commit) GetEntityData() *types.CommonEntityData {
    commit.EntityData.YFilter = commit.YFilter
    commit.EntityData.YangName = "commit"
    commit.EntityData.BundleName = "cisco_ios_xr"
    commit.EntityData.ParentYangName = "commits"
    commit.EntityData.SegmentPath = "commit" + types.AddKeyToken(commit.Commit, "commit")
    commit.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/config-commit/commits/" + commit.EntityData.SegmentPath
    commit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commit.EntityData.Children = types.NewOrderedMap()
    commit.EntityData.Leafs = types.NewOrderedMap()
    commit.EntityData.Leafs.Append("commit", types.YLeaf{"Commit", commit.Commit})
    commit.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", commit.Timestamp})
    commit.EntityData.Leafs.Append("commit-id", types.YLeaf{"CommitId", commit.CommitId})
    commit.EntityData.Leafs.Append("user-id", types.YLeaf{"UserId", commit.UserId})
    commit.EntityData.Leafs.Append("line", types.YLeaf{"Line", commit.Line})
    commit.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", commit.ClientName})
    commit.EntityData.Leafs.Append("label", types.YLeaf{"Label", commit.Label})
    commit.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", commit.Comment})

    commit.EntityData.YListKeys = []string {"Commit"}

    return &(commit.EntityData)
}

// ConfigManager_Global_HistoryTables
// Show Configuration History <*> Global path
// information
type ConfigManager_Global_HistoryTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of possible type of History. The type is slice of
    // ConfigManager_Global_HistoryTables_HistoryTable.
    HistoryTable []*ConfigManager_Global_HistoryTables_HistoryTable
}

func (historyTables *ConfigManager_Global_HistoryTables) GetEntityData() *types.CommonEntityData {
    historyTables.EntityData.YFilter = historyTables.YFilter
    historyTables.EntityData.YangName = "history-tables"
    historyTables.EntityData.BundleName = "cisco_ios_xr"
    historyTables.EntityData.ParentYangName = "global"
    historyTables.EntityData.SegmentPath = "history-tables"
    historyTables.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/" + historyTables.EntityData.SegmentPath
    historyTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyTables.EntityData.Children = types.NewOrderedMap()
    historyTables.EntityData.Children.Append("history-table", types.YChild{"HistoryTable", nil})
    for i := range historyTables.HistoryTable {
        historyTables.EntityData.Children.Append(types.GetSegmentPath(historyTables.HistoryTable[i]), types.YChild{"HistoryTable", historyTables.HistoryTable[i]})
    }
    historyTables.EntityData.Leafs = types.NewOrderedMap()

    historyTables.EntityData.YListKeys = []string {}

    return &(historyTables.EntityData)
}

// ConfigManager_Global_HistoryTables_HistoryTable
// List of possible type of History
type ConfigManager_Global_HistoryTables_HistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Type of History. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HistoryType interface{}

    // History summary information for a specific type of history. The type is
    // slice of ConfigManager_Global_HistoryTables_HistoryTable_History.
    History []*ConfigManager_Global_HistoryTables_HistoryTable_History
}

func (historyTable *ConfigManager_Global_HistoryTables_HistoryTable) GetEntityData() *types.CommonEntityData {
    historyTable.EntityData.YFilter = historyTable.YFilter
    historyTable.EntityData.YangName = "history-table"
    historyTable.EntityData.BundleName = "cisco_ios_xr"
    historyTable.EntityData.ParentYangName = "history-tables"
    historyTable.EntityData.SegmentPath = "history-table" + types.AddKeyToken(historyTable.HistoryType, "history-type")
    historyTable.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/" + historyTable.EntityData.SegmentPath
    historyTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyTable.EntityData.Children = types.NewOrderedMap()
    historyTable.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range historyTable.History {
        historyTable.EntityData.Children.Append(types.GetSegmentPath(historyTable.History[i]), types.YChild{"History", historyTable.History[i]})
    }
    historyTable.EntityData.Leafs = types.NewOrderedMap()
    historyTable.EntityData.Leafs.Append("history-type", types.YLeaf{"HistoryType", historyTable.HistoryType})

    historyTable.EntityData.YListKeys = []string {"HistoryType"}

    return &(historyTable.EntityData)
}

// ConfigManager_Global_HistoryTables_HistoryTable_History
// History summary information for a specific
// type of history
type ConfigManager_Global_HistoryTables_HistoryTable_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. History Record. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    History interface{}

    // Timestamp. The type is string.
    Timestamp interface{}

    // Content of the history.
    Info ConfigManager_Global_HistoryTables_HistoryTable_History_Info
}

func (history *ConfigManager_Global_HistoryTables_HistoryTable_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "history-table"
    history.EntityData.SegmentPath = "history" + types.AddKeyToken(history.History, "history")
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("info", types.YChild{"Info", &history.Info})
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("history", types.YLeaf{"History", history.History})
    history.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", history.Timestamp})

    history.EntityData.YListKeys = []string {"History"}

    return &(history.EntityData)
}

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info
// Content of the history
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is HistRecord.
    Type interface{}

    // B. The type is interface{} with range: 0..4294967295.
    A interface{}

    // alarm info.
    AlarmInfo ConfigManager_Global_HistoryTables_HistoryTable_History_Info_AlarmInfo

    // cfscheck info.
    CfscheckInfo ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CfscheckInfo

    // commit info.
    CommitInfo ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CommitInfo

    // oir info.
    OirInfo ConfigManager_Global_HistoryTables_HistoryTable_History_Info_OirInfo

    // shutdown info.
    ShutdownInfo ConfigManager_Global_HistoryTables_HistoryTable_History_Info_ShutdownInfo

    // startup info.
    StartupInfo ConfigManager_Global_HistoryTables_HistoryTable_History_Info_StartupInfo

    // backup info.
    BackupInfo ConfigManager_Global_HistoryTables_HistoryTable_History_Info_BackupInfo
}

func (info *ConfigManager_Global_HistoryTables_HistoryTable_History_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "history"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/" + info.EntityData.SegmentPath
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

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info_AlarmInfo
// alarm info
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State. The type is string.
    State interface{}

    // Where. The type is string.
    Where interface{}
}

func (alarmInfo *ConfigManager_Global_HistoryTables_HistoryTable_History_Info_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "info"
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/info/" + alarmInfo.EntityData.SegmentPath
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

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CfscheckInfo
// cfscheck info
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CfscheckInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UserId. The type is string.
    UserId interface{}

    // Line. The type is string.
    Line interface{}
}

func (cfscheckInfo *ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CfscheckInfo) GetEntityData() *types.CommonEntityData {
    cfscheckInfo.EntityData.YFilter = cfscheckInfo.YFilter
    cfscheckInfo.EntityData.YangName = "cfscheck-info"
    cfscheckInfo.EntityData.BundleName = "cisco_ios_xr"
    cfscheckInfo.EntityData.ParentYangName = "info"
    cfscheckInfo.EntityData.SegmentPath = "cfscheck-info"
    cfscheckInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/info/" + cfscheckInfo.EntityData.SegmentPath
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

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CommitInfo
// commit info
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CommitInfo struct {
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

func (commitInfo *ConfigManager_Global_HistoryTables_HistoryTable_History_Info_CommitInfo) GetEntityData() *types.CommonEntityData {
    commitInfo.EntityData.YFilter = commitInfo.YFilter
    commitInfo.EntityData.YangName = "commit-info"
    commitInfo.EntityData.BundleName = "cisco_ios_xr"
    commitInfo.EntityData.ParentYangName = "info"
    commitInfo.EntityData.SegmentPath = "commit-info"
    commitInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/info/" + commitInfo.EntityData.SegmentPath
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

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info_OirInfo
// oir info
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info_OirInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config Type. The type is string.
    ConfigType interface{}

    // Operation. The type is string.
    Operation interface{}

    // Config Name. The type is string.
    ConfigName interface{}
}

func (oirInfo *ConfigManager_Global_HistoryTables_HistoryTable_History_Info_OirInfo) GetEntityData() *types.CommonEntityData {
    oirInfo.EntityData.YFilter = oirInfo.YFilter
    oirInfo.EntityData.YangName = "oir-info"
    oirInfo.EntityData.BundleName = "cisco_ios_xr"
    oirInfo.EntityData.ParentYangName = "info"
    oirInfo.EntityData.SegmentPath = "oir-info"
    oirInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/info/" + oirInfo.EntityData.SegmentPath
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

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info_ShutdownInfo
// shutdown info
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info_ShutdownInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Comment. The type is string.
    Comment interface{}
}

func (shutdownInfo *ConfigManager_Global_HistoryTables_HistoryTable_History_Info_ShutdownInfo) GetEntityData() *types.CommonEntityData {
    shutdownInfo.EntityData.YFilter = shutdownInfo.YFilter
    shutdownInfo.EntityData.YangName = "shutdown-info"
    shutdownInfo.EntityData.BundleName = "cisco_ios_xr"
    shutdownInfo.EntityData.ParentYangName = "info"
    shutdownInfo.EntityData.SegmentPath = "shutdown-info"
    shutdownInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/info/" + shutdownInfo.EntityData.SegmentPath
    shutdownInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shutdownInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shutdownInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shutdownInfo.EntityData.Children = types.NewOrderedMap()
    shutdownInfo.EntityData.Leafs = types.NewOrderedMap()
    shutdownInfo.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", shutdownInfo.Comment})

    shutdownInfo.EntityData.YListKeys = []string {}

    return &(shutdownInfo.EntityData)
}

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info_StartupInfo
// startup info
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info_StartupInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // How Booted. The type is string.
    HowBooted interface{}

    // Boot Path. The type is string.
    BootPath interface{}
}

func (startupInfo *ConfigManager_Global_HistoryTables_HistoryTable_History_Info_StartupInfo) GetEntityData() *types.CommonEntityData {
    startupInfo.EntityData.YFilter = startupInfo.YFilter
    startupInfo.EntityData.YangName = "startup-info"
    startupInfo.EntityData.BundleName = "cisco_ios_xr"
    startupInfo.EntityData.ParentYangName = "info"
    startupInfo.EntityData.SegmentPath = "startup-info"
    startupInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/info/" + startupInfo.EntityData.SegmentPath
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

// ConfigManager_Global_HistoryTables_HistoryTable_History_Info_BackupInfo
// backup info
type ConfigManager_Global_HistoryTables_HistoryTable_History_Info_BackupInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Comment. The type is string.
    Comment interface{}
}

func (backupInfo *ConfigManager_Global_HistoryTables_HistoryTable_History_Info_BackupInfo) GetEntityData() *types.CommonEntityData {
    backupInfo.EntityData.YFilter = backupInfo.YFilter
    backupInfo.EntityData.YangName = "backup-info"
    backupInfo.EntityData.BundleName = "cisco_ios_xr"
    backupInfo.EntityData.ParentYangName = "info"
    backupInfo.EntityData.SegmentPath = "backup-info"
    backupInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/history-tables/history-table/history/info/" + backupInfo.EntityData.SegmentPath
    backupInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupInfo.EntityData.Children = types.NewOrderedMap()
    backupInfo.EntityData.Leafs = types.NewOrderedMap()
    backupInfo.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", backupInfo.Comment})

    backupInfo.EntityData.YListKeys = []string {}

    return &(backupInfo.EntityData)
}

// ConfigManager_Global_ConfigSession
// Show Configuration Session <*> Global path
// information
type ConfigManager_Global_ConfigSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Configuration Session Detail.
    Sessions ConfigManager_Global_ConfigSession_Sessions
}

func (configSession *ConfigManager_Global_ConfigSession) GetEntityData() *types.CommonEntityData {
    configSession.EntityData.YFilter = configSession.YFilter
    configSession.EntityData.YangName = "config-session"
    configSession.EntityData.BundleName = "cisco_ios_xr"
    configSession.EntityData.ParentYangName = "global"
    configSession.EntityData.SegmentPath = "config-session"
    configSession.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/" + configSession.EntityData.SegmentPath
    configSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configSession.EntityData.Children = types.NewOrderedMap()
    configSession.EntityData.Children.Append("sessions", types.YChild{"Sessions", &configSession.Sessions})
    configSession.EntityData.Leafs = types.NewOrderedMap()

    configSession.EntityData.YListKeys = []string {}

    return &(configSession.EntityData)
}

// ConfigManager_Global_ConfigSession_Sessions
// Show Configuration Session Detail
type ConfigManager_Global_ConfigSession_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session information for a specific session entry. The type is slice of
    // ConfigManager_Global_ConfigSession_Sessions_Session.
    Session []*ConfigManager_Global_ConfigSession_Sessions_Session
}

func (sessions *ConfigManager_Global_ConfigSession_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "config-session"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/config-session/" + sessions.EntityData.SegmentPath
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// ConfigManager_Global_ConfigSession_Sessions_Session
// Session information for a specific session
// entry
type ConfigManager_Global_ConfigSession_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Each Session Details. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Session interface{}

    // Session Id. The type is string.
    SessionId interface{}

    // Client Name. The type is string.
    ClientName interface{}

    // UserId. The type is string.
    UserId interface{}

    // TtyName. The type is string.
    TtyName interface{}

    // Timestamp. The type is string.
    Timestamp interface{}

    // Lock Flag. The type is string.
    LockFlag interface{}

    // Trial Flag. The type is string.
    TrialFlag interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Process Name. The type is string.
    ProcessName interface{}

    // Node Name. The type is string.
    NodeName interface{}

    // Elapsed Time. The type is string.
    ElapsedTime interface{}
}

func (session *ConfigManager_Global_ConfigSession_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.Session, "session")
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-exec-oper:config-manager/global/config-session/sessions/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session", types.YLeaf{"Session", session.Session})
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", session.ClientName})
    session.EntityData.Leafs.Append("user-id", types.YLeaf{"UserId", session.UserId})
    session.EntityData.Leafs.Append("tty-name", types.YLeaf{"TtyName", session.TtyName})
    session.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", session.Timestamp})
    session.EntityData.Leafs.Append("lock-flag", types.YLeaf{"LockFlag", session.LockFlag})
    session.EntityData.Leafs.Append("trial-flag", types.YLeaf{"TrialFlag", session.TrialFlag})
    session.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", session.Pid})
    session.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", session.ProcessName})
    session.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", session.NodeName})
    session.EntityData.Leafs.Append("elapsed-time", types.YLeaf{"ElapsedTime", session.ElapsedTime})

    session.EntityData.YListKeys = []string {"Session"}

    return &(session.EntityData)
}

