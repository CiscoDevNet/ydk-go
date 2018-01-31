// DMI self-management YANG module for IOS
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package cisco_ia

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ia"))
    ydk.RegisterEntity("{http://cisco.com/yang/cisco-ia sync-from}", reflect.TypeOf(SyncFrom{}))
    ydk.RegisterEntity("cisco-ia:sync-from", reflect.TypeOf(SyncFrom{}))
    ydk.RegisterEntity("{http://cisco.com/yang/cisco-ia save-config}", reflect.TypeOf(SaveConfig{}))
    ydk.RegisterEntity("cisco-ia:save-config", reflect.TypeOf(SaveConfig{}))
    ydk.RegisterEntity("{http://cisco.com/yang/cisco-ia is-syncing}", reflect.TypeOf(IsSyncing{}))
    ydk.RegisterEntity("cisco-ia:is-syncing", reflect.TypeOf(IsSyncing{}))
    ydk.RegisterEntity("{http://cisco.com/yang/cisco-ia checkpoint}", reflect.TypeOf(Checkpoint{}))
    ydk.RegisterEntity("cisco-ia:checkpoint", reflect.TypeOf(Checkpoint{}))
    ydk.RegisterEntity("{http://cisco.com/yang/cisco-ia revert}", reflect.TypeOf(Revert{}))
    ydk.RegisterEntity("cisco-ia:revert", reflect.TypeOf(Revert{}))
    ydk.RegisterEntity("{http://cisco.com/yang/cisco-ia rollback}", reflect.TypeOf(Rollback{}))
    ydk.RegisterEntity("cisco-ia:rollback", reflect.TypeOf(Rollback{}))
}

// CiaLogLevel represents Logging levels for DMI
type CiaLogLevel string

const (
    // No logging
    CiaLogLevel_none CiaLogLevel = "none"

    // Log errors only
    CiaLogLevel_error CiaLogLevel = "error"

    // Log errors and warnings only
    CiaLogLevel_warning CiaLogLevel = "warning"

    // Log errors, warnings, and information only
    CiaLogLevel_information CiaLogLevel = "information"

    // Log errors, warnings, information,
    // and debug messages
    CiaLogLevel_debug CiaLogLevel = "debug"
)

// CiaSyncType represents database from the Network Element
type CiaSyncType string

const (
    // Do no synchronize the DMI
    // database from the Network Element
    CiaSyncType_disabled CiaSyncType = "disabled"

    // Collect "show running" from 
    // the Network Element
    CiaSyncType_without_defaults CiaSyncType = "without-defaults"

    // Collect "show running all" from 
    // the Network Element
    CiaSyncType_include_defaults CiaSyncType = "include-defaults"
)

// SyslogSeverity represents Standard Syslog logging levels)
type SyslogSeverity string

const (
    // No logging
    SyslogSeverity_none SyslogSeverity = "none"

    // Emergency Level Msg
    SyslogSeverity_emergency SyslogSeverity = "emergency"

    // Alert Level Msg
    SyslogSeverity_alert SyslogSeverity = "alert"

    // Critical Level Msg
    SyslogSeverity_critical SyslogSeverity = "critical"

    // Error Level Msg
    SyslogSeverity_error SyslogSeverity = "error"

    // Warning Level Msg
    SyslogSeverity_warning SyslogSeverity = "warning"

    // Notification Level Msg
    SyslogSeverity_notice SyslogSeverity = "notice"

    // Informational Level Msg
    SyslogSeverity_info SyslogSeverity = "info"

    // Debugging Level Msg
    SyslogSeverity_debug SyslogSeverity = "debug"
)

// OnepLogLevel represents Logging levels for Onep
type OnepLogLevel string

const (
    // No logging
    OnepLogLevel_none OnepLogLevel = "none"

    // Log fatal events only
    OnepLogLevel_fatal OnepLogLevel = "fatal"

    // Log fatal events and errors only
    OnepLogLevel_error OnepLogLevel = "error"

    // Log fatal events, errors, and warnings only
    OnepLogLevel_warning OnepLogLevel = "warning"

    // Log fatal events, errors, warnings, 
    // and information only
    OnepLogLevel_information OnepLogLevel = "information"

    // Log fatal events, errors, warnings, information,
    // and debug messages
    OnepLogLevel_debug OnepLogLevel = "debug"

    // Log all messages
    OnepLogLevel_trace OnepLogLevel = "trace"
)

// SyncFrom
// Synchronize the network element's 
// running-configuration to ConfD.
type SyncFrom struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input SyncFrom_Input

    
    Output SyncFrom_Output
}

func (syncFrom *SyncFrom) GetFilter() yfilter.YFilter { return syncFrom.YFilter }

func (syncFrom *SyncFrom) SetFilter(yf yfilter.YFilter) { syncFrom.YFilter = yf }

func (syncFrom *SyncFrom) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (syncFrom *SyncFrom) GetSegmentPath() string {
    return "cisco-ia:sync-from"
}

func (syncFrom *SyncFrom) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &syncFrom.Input
    }
    if childYangName == "output" {
        return &syncFrom.Output
    }
    return nil
}

func (syncFrom *SyncFrom) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &syncFrom.Input
    children["output"] = &syncFrom.Output
    return children
}

func (syncFrom *SyncFrom) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (syncFrom *SyncFrom) GetBundleName() string { return "cisco_ios_xe" }

func (syncFrom *SyncFrom) GetYangName() string { return "sync-from" }

func (syncFrom *SyncFrom) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (syncFrom *SyncFrom) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (syncFrom *SyncFrom) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (syncFrom *SyncFrom) SetParent(parent types.Entity) { syncFrom.parent = parent }

func (syncFrom *SyncFrom) GetParent() types.Entity { return syncFrom.parent }

func (syncFrom *SyncFrom) GetParentYangName() string { return "cisco-ia" }

// SyncFrom_Input
type SyncFrom_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sends the output of  "show running all" line by line to Confd. The type is
    // interface{}.
    SyncDefaults interface{}

    // Sync everything under /native. Ignore any preserve paths. The type is
    // interface{}.
    IgnorePresrvPaths interface{}
}

func (input *SyncFrom_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *SyncFrom_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *SyncFrom_Input) GetGoName(yname string) string {
    if yname == "sync-defaults" { return "SyncDefaults" }
    if yname == "ignore-presrv-paths" { return "IgnorePresrvPaths" }
    return ""
}

func (input *SyncFrom_Input) GetSegmentPath() string {
    return "input"
}

func (input *SyncFrom_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *SyncFrom_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *SyncFrom_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-defaults"] = input.SyncDefaults
    leafs["ignore-presrv-paths"] = input.IgnorePresrvPaths
    return leafs
}

func (input *SyncFrom_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *SyncFrom_Input) GetYangName() string { return "input" }

func (input *SyncFrom_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *SyncFrom_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *SyncFrom_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *SyncFrom_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *SyncFrom_Input) GetParent() types.Entity { return input.parent }

func (input *SyncFrom_Input) GetParentYangName() string { return "sync-from" }

// SyncFrom_Output
type SyncFrom_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *SyncFrom_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *SyncFrom_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *SyncFrom_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *SyncFrom_Output) GetSegmentPath() string {
    return "output"
}

func (output *SyncFrom_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *SyncFrom_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *SyncFrom_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *SyncFrom_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *SyncFrom_Output) GetYangName() string { return "output" }

func (output *SyncFrom_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *SyncFrom_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *SyncFrom_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *SyncFrom_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *SyncFrom_Output) GetParent() types.Entity { return output.parent }

func (output *SyncFrom_Output) GetParentYangName() string { return "sync-from" }

// SaveConfig
// Copy the running-config to 
// startup-config on the Network
// Element.
type SaveConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Output SaveConfig_Output
}

func (saveConfig *SaveConfig) GetFilter() yfilter.YFilter { return saveConfig.YFilter }

func (saveConfig *SaveConfig) SetFilter(yf yfilter.YFilter) { saveConfig.YFilter = yf }

func (saveConfig *SaveConfig) GetGoName(yname string) string {
    if yname == "output" { return "Output" }
    return ""
}

func (saveConfig *SaveConfig) GetSegmentPath() string {
    return "cisco-ia:save-config"
}

func (saveConfig *SaveConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        return &saveConfig.Output
    }
    return nil
}

func (saveConfig *SaveConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output"] = &saveConfig.Output
    return children
}

func (saveConfig *SaveConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (saveConfig *SaveConfig) GetBundleName() string { return "cisco_ios_xe" }

func (saveConfig *SaveConfig) GetYangName() string { return "save-config" }

func (saveConfig *SaveConfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (saveConfig *SaveConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (saveConfig *SaveConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (saveConfig *SaveConfig) SetParent(parent types.Entity) { saveConfig.parent = parent }

func (saveConfig *SaveConfig) GetParent() types.Entity { return saveConfig.parent }

func (saveConfig *SaveConfig) GetParentYangName() string { return "cisco-ia" }

// SaveConfig_Output
type SaveConfig_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *SaveConfig_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *SaveConfig_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *SaveConfig_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *SaveConfig_Output) GetSegmentPath() string {
    return "output"
}

func (output *SaveConfig_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *SaveConfig_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *SaveConfig_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *SaveConfig_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *SaveConfig_Output) GetYangName() string { return "output" }

func (output *SaveConfig_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *SaveConfig_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *SaveConfig_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *SaveConfig_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *SaveConfig_Output) GetParent() types.Entity { return output.parent }

func (output *SaveConfig_Output) GetParentYangName() string { return "save-config" }

// IsSyncing
// Checks to see if sync from the
// network element to the running data store
// is in progress.
type IsSyncing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Output IsSyncing_Output
}

func (isSyncing *IsSyncing) GetFilter() yfilter.YFilter { return isSyncing.YFilter }

func (isSyncing *IsSyncing) SetFilter(yf yfilter.YFilter) { isSyncing.YFilter = yf }

func (isSyncing *IsSyncing) GetGoName(yname string) string {
    if yname == "output" { return "Output" }
    return ""
}

func (isSyncing *IsSyncing) GetSegmentPath() string {
    return "cisco-ia:is-syncing"
}

func (isSyncing *IsSyncing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        return &isSyncing.Output
    }
    return nil
}

func (isSyncing *IsSyncing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output"] = &isSyncing.Output
    return children
}

func (isSyncing *IsSyncing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (isSyncing *IsSyncing) GetBundleName() string { return "cisco_ios_xe" }

func (isSyncing *IsSyncing) GetYangName() string { return "is-syncing" }

func (isSyncing *IsSyncing) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (isSyncing *IsSyncing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (isSyncing *IsSyncing) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (isSyncing *IsSyncing) SetParent(parent types.Entity) { isSyncing.parent = parent }

func (isSyncing *IsSyncing) GetParent() types.Entity { return isSyncing.parent }

func (isSyncing *IsSyncing) GetParentYangName() string { return "cisco-ia" }

// IsSyncing_Output
type IsSyncing_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *IsSyncing_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *IsSyncing_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *IsSyncing_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *IsSyncing_Output) GetSegmentPath() string {
    return "output"
}

func (output *IsSyncing_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *IsSyncing_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *IsSyncing_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *IsSyncing_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *IsSyncing_Output) GetYangName() string { return "output" }

func (output *IsSyncing_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *IsSyncing_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *IsSyncing_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *IsSyncing_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *IsSyncing_Output) GetParent() types.Entity { return output.parent }

func (output *IsSyncing_Output) GetParentYangName() string { return "is-syncing" }

// Checkpoint
// Create a configuration rollback checkpoint.
// Equivalent to the "archive config" CLI
type Checkpoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Output Checkpoint_Output
}

func (checkpoint *Checkpoint) GetFilter() yfilter.YFilter { return checkpoint.YFilter }

func (checkpoint *Checkpoint) SetFilter(yf yfilter.YFilter) { checkpoint.YFilter = yf }

func (checkpoint *Checkpoint) GetGoName(yname string) string {
    if yname == "output" { return "Output" }
    return ""
}

func (checkpoint *Checkpoint) GetSegmentPath() string {
    return "cisco-ia:checkpoint"
}

func (checkpoint *Checkpoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        return &checkpoint.Output
    }
    return nil
}

func (checkpoint *Checkpoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output"] = &checkpoint.Output
    return children
}

func (checkpoint *Checkpoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (checkpoint *Checkpoint) GetBundleName() string { return "cisco_ios_xe" }

func (checkpoint *Checkpoint) GetYangName() string { return "checkpoint" }

func (checkpoint *Checkpoint) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (checkpoint *Checkpoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (checkpoint *Checkpoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (checkpoint *Checkpoint) SetParent(parent types.Entity) { checkpoint.parent = parent }

func (checkpoint *Checkpoint) GetParent() types.Entity { return checkpoint.parent }

func (checkpoint *Checkpoint) GetParentYangName() string { return "cisco-ia" }

// Checkpoint_Output
type Checkpoint_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Checkpoint_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Checkpoint_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Checkpoint_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *Checkpoint_Output) GetSegmentPath() string {
    return "output"
}

func (output *Checkpoint_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Checkpoint_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Checkpoint_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *Checkpoint_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *Checkpoint_Output) GetYangName() string { return "output" }

func (output *Checkpoint_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *Checkpoint_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *Checkpoint_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *Checkpoint_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Checkpoint_Output) GetParent() types.Entity { return output.parent }

func (output *Checkpoint_Output) GetParentYangName() string { return "checkpoint" }

// Revert
// Cancel the timed rollback and trigger the
// rollback immediately, or to reset parameters 
// for the timed rollback
type Revert struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Revert_Input

    
    Output Revert_Output
}

func (revert *Revert) GetFilter() yfilter.YFilter { return revert.YFilter }

func (revert *Revert) SetFilter(yf yfilter.YFilter) { revert.YFilter = yf }

func (revert *Revert) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (revert *Revert) GetSegmentPath() string {
    return "cisco-ia:revert"
}

func (revert *Revert) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &revert.Input
    }
    if childYangName == "output" {
        return &revert.Output
    }
    return nil
}

func (revert *Revert) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &revert.Input
    children["output"] = &revert.Output
    return children
}

func (revert *Revert) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (revert *Revert) GetBundleName() string { return "cisco_ios_xe" }

func (revert *Revert) GetYangName() string { return "revert" }

func (revert *Revert) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (revert *Revert) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (revert *Revert) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (revert *Revert) SetParent(parent types.Entity) { revert.parent = parent }

func (revert *Revert) GetParent() types.Entity { return revert.parent }

func (revert *Revert) GetParentYangName() string { return "cisco-ia" }

// Revert_Input
type Revert_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // To cancel the timed rollback and  trigger the rollback immediately. The
    // type is interface{}.
    Now interface{}

    // Reset parameters for the timed rollback. The type is interface{} with
    // range: 1..120.
    Timer interface{}

    // Maximum allowable time period of no activity before reverting to the saved
    // configuration. The type is interface{} with range: 1..120.
    Idle interface{}
}

func (input *Revert_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Revert_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Revert_Input) GetGoName(yname string) string {
    if yname == "now" { return "Now" }
    if yname == "timer" { return "Timer" }
    if yname == "idle" { return "Idle" }
    return ""
}

func (input *Revert_Input) GetSegmentPath() string {
    return "input"
}

func (input *Revert_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Revert_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Revert_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["now"] = input.Now
    leafs["timer"] = input.Timer
    leafs["idle"] = input.Idle
    return leafs
}

func (input *Revert_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *Revert_Input) GetYangName() string { return "input" }

func (input *Revert_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *Revert_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *Revert_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *Revert_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Revert_Input) GetParent() types.Entity { return input.parent }

func (input *Revert_Input) GetParentYangName() string { return "revert" }

// Revert_Output
type Revert_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Revert_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Revert_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Revert_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *Revert_Output) GetSegmentPath() string {
    return "output"
}

func (output *Revert_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Revert_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Revert_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *Revert_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *Revert_Output) GetYangName() string { return "output" }

func (output *Revert_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *Revert_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *Revert_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *Revert_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Revert_Output) GetParent() types.Entity { return output.parent }

func (output *Revert_Output) GetParentYangName() string { return "revert" }

// Rollback
// Replaces the current running configuration 
// file with a saved Cisco IOS XE configuration file.
type Rollback struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Rollback_Input

    
    Output Rollback_Output
}

func (rollback *Rollback) GetFilter() yfilter.YFilter { return rollback.YFilter }

func (rollback *Rollback) SetFilter(yf yfilter.YFilter) { rollback.YFilter = yf }

func (rollback *Rollback) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (rollback *Rollback) GetSegmentPath() string {
    return "cisco-ia:rollback"
}

func (rollback *Rollback) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &rollback.Input
    }
    if childYangName == "output" {
        return &rollback.Output
    }
    return nil
}

func (rollback *Rollback) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &rollback.Input
    children["output"] = &rollback.Output
    return children
}

func (rollback *Rollback) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rollback *Rollback) GetBundleName() string { return "cisco_ios_xe" }

func (rollback *Rollback) GetYangName() string { return "rollback" }

func (rollback *Rollback) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rollback *Rollback) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rollback *Rollback) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rollback *Rollback) SetParent(parent types.Entity) { rollback.parent = parent }

func (rollback *Rollback) GetParent() types.Entity { return rollback.parent }

func (rollback *Rollback) GetParentYangName() string { return "cisco-ia" }

// Rollback_Input
type Rollback_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Cisco IOS XE configuration file that is to  replace the current running
    // configuration. The type is string. This attribute is mandatory.
    TargetUrl interface{}

    // Display a list of the command lines applied by  the Cisco IOS XE software
    // parser during each pass  of the configuration replace operation. The total 
    // number of passes performed is also displayed. The type is bool. The default
    // value is false.
    Verbose interface{}

    // Disables the locking of the running  configuration file that prevents other
    // users from changing the running configuration  during a configuration
    // replace operation. The type is bool. The default value is false.
    Nolock interface{}

    // Reverts to the original configuration upon error. The type is interface{}.
    RevertOnError interface{}

    // Reverts to the original configuration if  specified time elapses. The type
    // is interface{} with range: 1..120.
    RevertTimer interface{}
}

func (input *Rollback_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Rollback_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Rollback_Input) GetGoName(yname string) string {
    if yname == "target-url" { return "TargetUrl" }
    if yname == "verbose" { return "Verbose" }
    if yname == "nolock" { return "Nolock" }
    if yname == "revert-on-error" { return "RevertOnError" }
    if yname == "revert-timer" { return "RevertTimer" }
    return ""
}

func (input *Rollback_Input) GetSegmentPath() string {
    return "input"
}

func (input *Rollback_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Rollback_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Rollback_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-url"] = input.TargetUrl
    leafs["verbose"] = input.Verbose
    leafs["nolock"] = input.Nolock
    leafs["revert-on-error"] = input.RevertOnError
    leafs["revert-timer"] = input.RevertTimer
    return leafs
}

func (input *Rollback_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *Rollback_Input) GetYangName() string { return "input" }

func (input *Rollback_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *Rollback_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *Rollback_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *Rollback_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Rollback_Input) GetParent() types.Entity { return input.parent }

func (input *Rollback_Input) GetParentYangName() string { return "rollback" }

// Rollback_Output
type Rollback_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Rollback_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Rollback_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Rollback_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *Rollback_Output) GetSegmentPath() string {
    return "output"
}

func (output *Rollback_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Rollback_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Rollback_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *Rollback_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *Rollback_Output) GetYangName() string { return "output" }

func (output *Rollback_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *Rollback_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *Rollback_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *Rollback_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Rollback_Output) GetParent() types.Entity { return output.parent }

func (output *Rollback_Output) GetParentYangName() string { return "rollback" }

