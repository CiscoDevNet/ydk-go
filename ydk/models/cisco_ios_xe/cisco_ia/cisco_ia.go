// DMI self-management YANG module for IOS
// Copyright (c) 2016, 2018 by Cisco Systems, Inc.
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
    CiaLogLevel_error_ CiaLogLevel = "error"

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
    SyslogSeverity_error_ SyslogSeverity = "error"

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
    OnepLogLevel_error_ OnepLogLevel = "error"

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input SyncFrom_Input

    
    Output SyncFrom_Output
}

func (syncFrom *SyncFrom) GetEntityData() *types.CommonEntityData {
    syncFrom.EntityData.YFilter = syncFrom.YFilter
    syncFrom.EntityData.YangName = "sync-from"
    syncFrom.EntityData.BundleName = "cisco_ios_xe"
    syncFrom.EntityData.ParentYangName = "cisco-ia"
    syncFrom.EntityData.SegmentPath = "cisco-ia:sync-from"
    syncFrom.EntityData.AbsolutePath = syncFrom.EntityData.SegmentPath
    syncFrom.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    syncFrom.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    syncFrom.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    syncFrom.EntityData.Children = types.NewOrderedMap()
    syncFrom.EntityData.Children.Append("input", types.YChild{"Input", &syncFrom.Input})
    syncFrom.EntityData.Children.Append("output", types.YChild{"Output", &syncFrom.Output})
    syncFrom.EntityData.Leafs = types.NewOrderedMap()

    syncFrom.EntityData.YListKeys = []string {}

    return &(syncFrom.EntityData)
}

// SyncFrom_Input
type SyncFrom_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sends the output of  "show running all" line by line to Confd. The type is
    // interface{}.
    SyncDefaults interface{}

    // Sync everything under /native. Ignore any preserve paths. The type is
    // interface{}.
    IgnorePresrvPaths interface{}
}

func (input *SyncFrom_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "sync-from"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "cisco-ia:sync-from/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("sync-defaults", types.YLeaf{"SyncDefaults", input.SyncDefaults})
    input.EntityData.Leafs.Append("ignore-presrv-paths", types.YLeaf{"IgnorePresrvPaths", input.IgnorePresrvPaths})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// SyncFrom_Output
type SyncFrom_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *SyncFrom_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "sync-from"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-ia:sync-from/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// SaveConfig
// Copy the running-config to 
// startup-config on the Network
// Element.
type SaveConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Output SaveConfig_Output
}

func (saveConfig *SaveConfig) GetEntityData() *types.CommonEntityData {
    saveConfig.EntityData.YFilter = saveConfig.YFilter
    saveConfig.EntityData.YangName = "save-config"
    saveConfig.EntityData.BundleName = "cisco_ios_xe"
    saveConfig.EntityData.ParentYangName = "cisco-ia"
    saveConfig.EntityData.SegmentPath = "cisco-ia:save-config"
    saveConfig.EntityData.AbsolutePath = saveConfig.EntityData.SegmentPath
    saveConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    saveConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    saveConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    saveConfig.EntityData.Children = types.NewOrderedMap()
    saveConfig.EntityData.Children.Append("output", types.YChild{"Output", &saveConfig.Output})
    saveConfig.EntityData.Leafs = types.NewOrderedMap()

    saveConfig.EntityData.YListKeys = []string {}

    return &(saveConfig.EntityData)
}

// SaveConfig_Output
type SaveConfig_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *SaveConfig_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "save-config"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-ia:save-config/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// IsSyncing
// Checks to see if sync from the
// network element to the running data store
// is in progress.
type IsSyncing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Output IsSyncing_Output
}

func (isSyncing *IsSyncing) GetEntityData() *types.CommonEntityData {
    isSyncing.EntityData.YFilter = isSyncing.YFilter
    isSyncing.EntityData.YangName = "is-syncing"
    isSyncing.EntityData.BundleName = "cisco_ios_xe"
    isSyncing.EntityData.ParentYangName = "cisco-ia"
    isSyncing.EntityData.SegmentPath = "cisco-ia:is-syncing"
    isSyncing.EntityData.AbsolutePath = isSyncing.EntityData.SegmentPath
    isSyncing.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    isSyncing.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    isSyncing.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    isSyncing.EntityData.Children = types.NewOrderedMap()
    isSyncing.EntityData.Children.Append("output", types.YChild{"Output", &isSyncing.Output})
    isSyncing.EntityData.Leafs = types.NewOrderedMap()

    isSyncing.EntityData.YListKeys = []string {}

    return &(isSyncing.EntityData)
}

// IsSyncing_Output
type IsSyncing_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *IsSyncing_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "is-syncing"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-ia:is-syncing/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Checkpoint
// Create a configuration rollback checkpoint.
// Equivalent to the "archive config" CLI
type Checkpoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Output Checkpoint_Output
}

func (checkpoint *Checkpoint) GetEntityData() *types.CommonEntityData {
    checkpoint.EntityData.YFilter = checkpoint.YFilter
    checkpoint.EntityData.YangName = "checkpoint"
    checkpoint.EntityData.BundleName = "cisco_ios_xe"
    checkpoint.EntityData.ParentYangName = "cisco-ia"
    checkpoint.EntityData.SegmentPath = "cisco-ia:checkpoint"
    checkpoint.EntityData.AbsolutePath = checkpoint.EntityData.SegmentPath
    checkpoint.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    checkpoint.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    checkpoint.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    checkpoint.EntityData.Children = types.NewOrderedMap()
    checkpoint.EntityData.Children.Append("output", types.YChild{"Output", &checkpoint.Output})
    checkpoint.EntityData.Leafs = types.NewOrderedMap()

    checkpoint.EntityData.YListKeys = []string {}

    return &(checkpoint.EntityData)
}

// Checkpoint_Output
type Checkpoint_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Checkpoint_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "checkpoint"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-ia:checkpoint/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Revert
// Cancel the timed rollback and trigger the
// rollback immediately, or to reset parameters 
// for the timed rollback
type Revert struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Revert_Input

    
    Output Revert_Output
}

func (revert *Revert) GetEntityData() *types.CommonEntityData {
    revert.EntityData.YFilter = revert.YFilter
    revert.EntityData.YangName = "revert"
    revert.EntityData.BundleName = "cisco_ios_xe"
    revert.EntityData.ParentYangName = "cisco-ia"
    revert.EntityData.SegmentPath = "cisco-ia:revert"
    revert.EntityData.AbsolutePath = revert.EntityData.SegmentPath
    revert.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    revert.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    revert.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    revert.EntityData.Children = types.NewOrderedMap()
    revert.EntityData.Children.Append("input", types.YChild{"Input", &revert.Input})
    revert.EntityData.Children.Append("output", types.YChild{"Output", &revert.Output})
    revert.EntityData.Leafs = types.NewOrderedMap()

    revert.EntityData.YListKeys = []string {}

    return &(revert.EntityData)
}

// Revert_Input
type Revert_Input struct {
    EntityData types.CommonEntityData
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

func (input *Revert_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "revert"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "cisco-ia:revert/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("now", types.YLeaf{"Now", input.Now})
    input.EntityData.Leafs.Append("timer", types.YLeaf{"Timer", input.Timer})
    input.EntityData.Leafs.Append("idle", types.YLeaf{"Idle", input.Idle})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Revert_Output
type Revert_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Revert_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "revert"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-ia:revert/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Rollback
// Replaces the current running configuration 
// file with a saved Cisco IOS XE configuration file.
type Rollback struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Rollback_Input

    
    Output Rollback_Output
}

func (rollback *Rollback) GetEntityData() *types.CommonEntityData {
    rollback.EntityData.YFilter = rollback.YFilter
    rollback.EntityData.YangName = "rollback"
    rollback.EntityData.BundleName = "cisco_ios_xe"
    rollback.EntityData.ParentYangName = "cisco-ia"
    rollback.EntityData.SegmentPath = "cisco-ia:rollback"
    rollback.EntityData.AbsolutePath = rollback.EntityData.SegmentPath
    rollback.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rollback.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rollback.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rollback.EntityData.Children = types.NewOrderedMap()
    rollback.EntityData.Children.Append("input", types.YChild{"Input", &rollback.Input})
    rollback.EntityData.Children.Append("output", types.YChild{"Output", &rollback.Output})
    rollback.EntityData.Leafs = types.NewOrderedMap()

    rollback.EntityData.YListKeys = []string {}

    return &(rollback.EntityData)
}

// Rollback_Input
type Rollback_Input struct {
    EntityData types.CommonEntityData
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

func (input *Rollback_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "rollback"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "cisco-ia:rollback/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("target-url", types.YLeaf{"TargetUrl", input.TargetUrl})
    input.EntityData.Leafs.Append("verbose", types.YLeaf{"Verbose", input.Verbose})
    input.EntityData.Leafs.Append("nolock", types.YLeaf{"Nolock", input.Nolock})
    input.EntityData.Leafs.Append("revert-on-error", types.YLeaf{"RevertOnError", input.RevertOnError})
    input.EntityData.Leafs.Append("revert-timer", types.YLeaf{"RevertTimer", input.RevertTimer})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Rollback_Output
type Rollback_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Rollback_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "rollback"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-ia:rollback/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

