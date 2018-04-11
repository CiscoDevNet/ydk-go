// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Fault management YANG model. 
// 
// Copyright(c) 2014-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_fm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_fm"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-fm fm}", reflect.TypeOf(Fm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-fm:fm", reflect.TypeOf(Fm{}))
}

// FmFaultStateT represents The status value for a given fault condition.
type FmFaultStateT string

const (
    FmFaultStateT_SET FmFaultStateT = "SET"

    FmFaultStateT_CLEAR FmFaultStateT = "CLEAR"

    FmFaultStateT_INFO FmFaultStateT = "INFO"

    FmFaultStateT_INVALID FmFaultStateT = "INVALID"

    FmFaultStateT_PARTIALLY_QUALIFIED FmFaultStateT = "PARTIALLY_QUALIFIED"

    FmFaultStateT_SOAKING_BEFORE_SET FmFaultStateT = "SOAKING_BEFORE_SET"

    FmFaultStateT_SOAKING_BEFORE_CLEAR FmFaultStateT = "SOAKING_BEFORE_CLEAR"

    FmFaultStateT_SUPPRESSED FmFaultStateT = "SUPPRESSED"

    FmFaultStateT_UPDATE FmFaultStateT = "UPDATE"
)

// FmFaultSeverityT
type FmFaultSeverityT string

const (
    FmFaultSeverityT_CRITICAL FmFaultSeverityT = "CRITICAL"

    FmFaultSeverityT_MAJOR FmFaultSeverityT = "MAJOR"

    FmFaultSeverityT_MINOR FmFaultSeverityT = "MINOR"

    FmFaultSeverityT_NR FmFaultSeverityT = "NR"
)

// FmActionT represents The List of supported Action Types
type FmActionT string

const (
    FmActionT_ISOLATION FmActionT = "ISOLATION"

    FmActionT_MITIGATION FmActionT = "MITIGATION"

    FmActionT_RECOVERY FmActionT = "RECOVERY"

    FmActionT_CORRELATION FmActionT = "CORRELATION"

    FmActionT_ALARM FmActionT = "ALARM"

    FmActionT_REPORT FmActionT = "REPORT"
)

// FmActionResultT represents The result of a certain fm action
type FmActionResultT string

const (
    FmActionResultT_SUCCESS FmActionResultT = "SUCCESS"

    FmActionResultT_FAILURE FmActionResultT = "FAILURE"

    FmActionResultT_NO_OP FmActionResultT = "NO-OP"
)

// FmRuleEvalResultT represents The result status of the evaluation of a FM rule.
type FmRuleEvalResultT string

const (
    FmRuleEvalResultT_SUCCESS FmRuleEvalResultT = "SUCCESS"

    FmRuleEvalResultT_FAILURE FmRuleEvalResultT = "FAILURE"
)

// GenericHaRole
type GenericHaRole string

const (
    GenericHaRole_no_ha_role GenericHaRole = "no-ha-role"

    GenericHaRole_Active GenericHaRole = "Active"

    GenericHaRole_Standby GenericHaRole = "Standby"
)

// FmCorrelationObjQualifierT
type FmCorrelationObjQualifierT string

const (
    FmCorrelationObjQualifierT_QUALIFIER_IGNORED FmCorrelationObjQualifierT = "QUALIFIER_IGNORED"

    FmCorrelationObjQualifierT_QUALIFIER_RACK FmCorrelationObjQualifierT = "QUALIFIER_RACK"

    FmCorrelationObjQualifierT_QUALIFIER_SLOT FmCorrelationObjQualifierT = "QUALIFIER_SLOT"

    FmCorrelationObjQualifierT_QUALIFIER_OBJECT FmCorrelationObjQualifierT = "QUALIFIER_OBJECT"
)

// FmHistoryStateT represents The fm history entry state.
type FmHistoryStateT string

const (
    FmHistoryStateT_FM_HISTORY_STATE_ACTIVE FmHistoryStateT = "FM_HISTORY_STATE_ACTIVE"

    FmHistoryStateT_FM_HISTORY_STATE_CLEARED FmHistoryStateT = "FM_HISTORY_STATE_CLEARED"

    FmHistoryStateT_FM_HISTORY_STATE_INVALID FmHistoryStateT = "FM_HISTORY_STATE_INVALID"
)

// FmServiceScopeT represents The fm service scope definting type.
type FmServiceScopeT string

const (
    FmServiceScopeT_FM_SERVICE_NODE_SCOPE FmServiceScopeT = "FM_SERVICE_NODE_SCOPE"

    FmServiceScopeT_FM_SERVICE_RACK_SCOPE FmServiceScopeT = "FM_SERVICE_RACK_SCOPE"

    FmServiceScopeT_FM_SERVICE_SYSTEM_SCOPE FmServiceScopeT = "FM_SERVICE_SYSTEM_SCOPE"
)

// Fm
// Sysadmin fault management operational data model
type Fm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents.
    Agents []Fm_Agents
}

func (fm *Fm) GetEntityData() *types.CommonEntityData {
    fm.EntityData.YFilter = fm.YFilter
    fm.EntityData.YangName = "fm"
    fm.EntityData.BundleName = "cisco_ios_xr"
    fm.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-fm"
    fm.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-fm:fm"
    fm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fm.EntityData.Children = make(map[string]types.YChild)
    fm.EntityData.Children["agents"] = types.YChild{"Agents", nil}
    for i := range fm.Agents {
        fm.EntityData.Children[types.GetSegmentPath(&fm.Agents[i])] = types.YChild{"Agents", &fm.Agents[i]}
    }
    fm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fm.EntityData)
}

// Fm_Agents
type Fm_Agents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // This attribute is a key. The type is string.
    Process interface{}

    // This attribute is a key. The type is string.
    Subsystem interface{}

    // This attribute is a key. The type is string.
    Agent interface{}

    
    FmInitials Fm_Agents_FmInitials

    
    FmTable Fm_Agents_FmTable

    
    FmInternals Fm_Agents_FmInternals

    
    FmAlarmMapping Fm_Agents_FmAlarmMapping

    
    FmStatistics Fm_Agents_FmStatistics
}

func (agents *Fm_Agents) GetEntityData() *types.CommonEntityData {
    agents.EntityData.YFilter = agents.YFilter
    agents.EntityData.YangName = "agents"
    agents.EntityData.BundleName = "cisco_ios_xr"
    agents.EntityData.ParentYangName = "fm"
    agents.EntityData.SegmentPath = "agents" + "[location='" + fmt.Sprintf("%v", agents.Location) + "']" + "[process='" + fmt.Sprintf("%v", agents.Process) + "']" + "[subsystem='" + fmt.Sprintf("%v", agents.Subsystem) + "']" + "[agent='" + fmt.Sprintf("%v", agents.Agent) + "']"
    agents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agents.EntityData.Children = make(map[string]types.YChild)
    agents.EntityData.Children["fm_initials"] = types.YChild{"FmInitials", &agents.FmInitials}
    agents.EntityData.Children["fm_table"] = types.YChild{"FmTable", &agents.FmTable}
    agents.EntityData.Children["fm_internals"] = types.YChild{"FmInternals", &agents.FmInternals}
    agents.EntityData.Children["fm_alarm_mapping"] = types.YChild{"FmAlarmMapping", &agents.FmAlarmMapping}
    agents.EntityData.Children["fm_statistics"] = types.YChild{"FmStatistics", &agents.FmStatistics}
    agents.EntityData.Leafs = make(map[string]types.YLeaf)
    agents.EntityData.Leafs["location"] = types.YLeaf{"Location", agents.Location}
    agents.EntityData.Leafs["process"] = types.YLeaf{"Process", agents.Process}
    agents.EntityData.Leafs["subsystem"] = types.YLeaf{"Subsystem", agents.Subsystem}
    agents.EntityData.Leafs["agent"] = types.YLeaf{"Agent", agents.Agent}
    return &(agents.EntityData)
}

// Fm_Agents_FmInitials
type Fm_Agents_FmInitials struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The levm pointer supplied to fm infra. The type is interface{} with range:
    // 0..18446744073709551615.
    Levm interface{}

    // The owner component Id. The type is interface{} with range: 0..4294967295.
    CompId interface{}

    // The name of the process in which this fm instance is active. The type is
    // string.
    Process interface{}

    // Default rule callback pointer. The type is interface{} with range:
    // 0..18446744073709551615.
    DefaultRuleCb interface{}

    // Default action callback pointer. The type is interface{} with range:
    // 0..18446744073709551615.
    DefaultActionCb interface{}

    // Default notification callback pointer. The type is interface{} with range:
    // 0..18446744073709551615.
    DefaultNotifCb interface{}

    // Default error callback pointer. The type is interface{} with range:
    // 0..18446744073709551615.
    DefaultErrorCb interface{}

    // Data Replica callback pointer. The type is interface{} with range:
    // 0..18446744073709551615.
    ReplicaCb interface{}
}

func (fmInitials *Fm_Agents_FmInitials) GetEntityData() *types.CommonEntityData {
    fmInitials.EntityData.YFilter = fmInitials.YFilter
    fmInitials.EntityData.YangName = "fm_initials"
    fmInitials.EntityData.BundleName = "cisco_ios_xr"
    fmInitials.EntityData.ParentYangName = "agents"
    fmInitials.EntityData.SegmentPath = "fm_initials"
    fmInitials.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmInitials.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmInitials.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmInitials.EntityData.Children = make(map[string]types.YChild)
    fmInitials.EntityData.Leafs = make(map[string]types.YLeaf)
    fmInitials.EntityData.Leafs["levm"] = types.YLeaf{"Levm", fmInitials.Levm}
    fmInitials.EntityData.Leafs["comp_id"] = types.YLeaf{"CompId", fmInitials.CompId}
    fmInitials.EntityData.Leafs["process"] = types.YLeaf{"Process", fmInitials.Process}
    fmInitials.EntityData.Leafs["default_rule_cb"] = types.YLeaf{"DefaultRuleCb", fmInitials.DefaultRuleCb}
    fmInitials.EntityData.Leafs["default_action_cb"] = types.YLeaf{"DefaultActionCb", fmInitials.DefaultActionCb}
    fmInitials.EntityData.Leafs["default_notif_cb"] = types.YLeaf{"DefaultNotifCb", fmInitials.DefaultNotifCb}
    fmInitials.EntityData.Leafs["default_error_cb"] = types.YLeaf{"DefaultErrorCb", fmInitials.DefaultErrorCb}
    fmInitials.EntityData.Leafs["replica_cb"] = types.YLeaf{"ReplicaCb", fmInitials.ReplicaCb}
    return &(fmInitials.EntityData)
}

// Fm_Agents_FmTable
type Fm_Agents_FmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents_FmTable_Brief.
    Brief []Fm_Agents_FmTable_Brief

    // The type is slice of Fm_Agents_FmTable_Entry.
    Entry []Fm_Agents_FmTable_Entry
}

func (fmTable *Fm_Agents_FmTable) GetEntityData() *types.CommonEntityData {
    fmTable.EntityData.YFilter = fmTable.YFilter
    fmTable.EntityData.YangName = "fm_table"
    fmTable.EntityData.BundleName = "cisco_ios_xr"
    fmTable.EntityData.ParentYangName = "agents"
    fmTable.EntityData.SegmentPath = "fm_table"
    fmTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmTable.EntityData.Children = make(map[string]types.YChild)
    fmTable.EntityData.Children["brief"] = types.YChild{"Brief", nil}
    for i := range fmTable.Brief {
        fmTable.EntityData.Children[types.GetSegmentPath(&fmTable.Brief[i])] = types.YChild{"Brief", &fmTable.Brief[i]}
    }
    fmTable.EntityData.Children["entry"] = types.YChild{"Entry", nil}
    for i := range fmTable.Entry {
        fmTable.EntityData.Children[types.GetSegmentPath(&fmTable.Entry[i])] = types.YChild{"Entry", &fmTable.Entry[i]}
    }
    fmTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fmTable.EntityData)
}

// Fm_Agents_FmTable_Brief
type Fm_Agents_FmTable_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // A descriptive name for the fault. The type is string.
    Name interface{}
}

func (brief *Fm_Agents_FmTable_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "fm_table"
    brief.EntityData.SegmentPath = "brief" + "[fm_subsystem_id='" + fmt.Sprintf("%v", brief.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", brief.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", brief.FmFaultTag) + "']"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", brief.FmSubsystemId}
    brief.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", brief.FmFaultType}
    brief.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", brief.FmFaultTag}
    brief.EntityData.Leafs["name"] = types.YLeaf{"Name", brief.Name}
    return &(brief.EntityData)
}

// Fm_Agents_FmTable_Entry
type Fm_Agents_FmTable_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    
    Detail Fm_Agents_FmTable_Entry_Detail

    // Causal list of fault ids for the specified fault. The type is slice of
    // Fm_Agents_FmTable_Entry_CausalList.
    CausalList []Fm_Agents_FmTable_Entry_CausalList

    // Dependency list of fault ids. The type is slice of
    // Fm_Agents_FmTable_Entry_DependencyList.
    DependencyList []Fm_Agents_FmTable_Entry_DependencyList

    // Propagation list of fault agents. The type is slice of
    // Fm_Agents_FmTable_Entry_PropagationList.
    PropagationList []Fm_Agents_FmTable_Entry_PropagationList

    // Notification list of fault agents. The type is slice of
    // Fm_Agents_FmTable_Entry_NotificationList.
    NotificationList []Fm_Agents_FmTable_Entry_NotificationList

    // escalation list of fault agents. The type is slice of
    // Fm_Agents_FmTable_Entry_EscalationList.
    EscalationList []Fm_Agents_FmTable_Entry_EscalationList

    
    Faults Fm_Agents_FmTable_Entry_Faults

    
    WaitingList Fm_Agents_FmTable_Entry_WaitingList
}

func (entry *Fm_Agents_FmTable_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "fm_table"
    entry.EntityData.SegmentPath = "entry" + "[fm_subsystem_id='" + fmt.Sprintf("%v", entry.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", entry.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", entry.FmFaultTag) + "']"
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = make(map[string]types.YChild)
    entry.EntityData.Children["detail"] = types.YChild{"Detail", &entry.Detail}
    entry.EntityData.Children["causal_list"] = types.YChild{"CausalList", nil}
    for i := range entry.CausalList {
        entry.EntityData.Children[types.GetSegmentPath(&entry.CausalList[i])] = types.YChild{"CausalList", &entry.CausalList[i]}
    }
    entry.EntityData.Children["dependency_list"] = types.YChild{"DependencyList", nil}
    for i := range entry.DependencyList {
        entry.EntityData.Children[types.GetSegmentPath(&entry.DependencyList[i])] = types.YChild{"DependencyList", &entry.DependencyList[i]}
    }
    entry.EntityData.Children["propagation_list"] = types.YChild{"PropagationList", nil}
    for i := range entry.PropagationList {
        entry.EntityData.Children[types.GetSegmentPath(&entry.PropagationList[i])] = types.YChild{"PropagationList", &entry.PropagationList[i]}
    }
    entry.EntityData.Children["notification_list"] = types.YChild{"NotificationList", nil}
    for i := range entry.NotificationList {
        entry.EntityData.Children[types.GetSegmentPath(&entry.NotificationList[i])] = types.YChild{"NotificationList", &entry.NotificationList[i]}
    }
    entry.EntityData.Children["escalation_list"] = types.YChild{"EscalationList", nil}
    for i := range entry.EscalationList {
        entry.EntityData.Children[types.GetSegmentPath(&entry.EscalationList[i])] = types.YChild{"EscalationList", &entry.EscalationList[i]}
    }
    entry.EntityData.Children["faults"] = types.YChild{"Faults", &entry.Faults}
    entry.EntityData.Children["waiting_list"] = types.YChild{"WaitingList", &entry.WaitingList}
    entry.EntityData.Leafs = make(map[string]types.YLeaf)
    entry.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", entry.FmSubsystemId}
    entry.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", entry.FmFaultType}
    entry.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", entry.FmFaultTag}
    return &(entry.EntityData)
}

// Fm_Agents_FmTable_Entry_Detail
type Fm_Agents_FmTable_Entry_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // Fault type identifier. The type is string.
    FmFaultType interface{}

    // Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // A descriptive name for the fault. The type is string.
    Name interface{}

    // Description of the fault. The type is string.
    Description interface{}

    // Fault detection logic. The type is string.
    DetectionLogic interface{}

    // The qualifier for the object used for correlation. The type is
    // FmCorrelationObjQualifierT.
    CorrObjQualifier interface{}
}

func (detail *Fm_Agents_FmTable_Entry_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "entry"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", detail.FmSubsystemId}
    detail.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", detail.FmFaultType}
    detail.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", detail.FmFaultTag}
    detail.EntityData.Leafs["name"] = types.YLeaf{"Name", detail.Name}
    detail.EntityData.Leafs["description"] = types.YLeaf{"Description", detail.Description}
    detail.EntityData.Leafs["detection_logic"] = types.YLeaf{"DetectionLogic", detail.DetectionLogic}
    detail.EntityData.Leafs["corr_obj_qualifier"] = types.YLeaf{"CorrObjQualifier", detail.CorrObjQualifier}
    return &(detail.EntityData)
}

// Fm_Agents_FmTable_Entry_CausalList
// Causal list of fault ids for the specified fault.
type Fm_Agents_FmTable_Entry_CausalList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}
}

func (causalList *Fm_Agents_FmTable_Entry_CausalList) GetEntityData() *types.CommonEntityData {
    causalList.EntityData.YFilter = causalList.YFilter
    causalList.EntityData.YangName = "causal_list"
    causalList.EntityData.BundleName = "cisco_ios_xr"
    causalList.EntityData.ParentYangName = "entry"
    causalList.EntityData.SegmentPath = "causal_list" + "[fm_subsystem_id='" + fmt.Sprintf("%v", causalList.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", causalList.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", causalList.FmFaultTag) + "']"
    causalList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    causalList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    causalList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    causalList.EntityData.Children = make(map[string]types.YChild)
    causalList.EntityData.Leafs = make(map[string]types.YLeaf)
    causalList.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", causalList.FmSubsystemId}
    causalList.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", causalList.FmFaultType}
    causalList.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", causalList.FmFaultTag}
    return &(causalList.EntityData)
}

// Fm_Agents_FmTable_Entry_DependencyList
// Dependency list of fault ids.
type Fm_Agents_FmTable_Entry_DependencyList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}
}

func (dependencyList *Fm_Agents_FmTable_Entry_DependencyList) GetEntityData() *types.CommonEntityData {
    dependencyList.EntityData.YFilter = dependencyList.YFilter
    dependencyList.EntityData.YangName = "dependency_list"
    dependencyList.EntityData.BundleName = "cisco_ios_xr"
    dependencyList.EntityData.ParentYangName = "entry"
    dependencyList.EntityData.SegmentPath = "dependency_list" + "[fm_subsystem_id='" + fmt.Sprintf("%v", dependencyList.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", dependencyList.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", dependencyList.FmFaultTag) + "']"
    dependencyList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dependencyList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dependencyList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dependencyList.EntityData.Children = make(map[string]types.YChild)
    dependencyList.EntityData.Leafs = make(map[string]types.YLeaf)
    dependencyList.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", dependencyList.FmSubsystemId}
    dependencyList.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", dependencyList.FmFaultType}
    dependencyList.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", dependencyList.FmFaultTag}
    return &(dependencyList.EntityData)
}

// Fm_Agents_FmTable_Entry_PropagationList
// Propagation list of fault agents.
type Fm_Agents_FmTable_Entry_PropagationList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The remote agent id assocaited with this fault. The type is string.
    RemoteAgentId interface{}
}

func (propagationList *Fm_Agents_FmTable_Entry_PropagationList) GetEntityData() *types.CommonEntityData {
    propagationList.EntityData.YFilter = propagationList.YFilter
    propagationList.EntityData.YangName = "propagation_list"
    propagationList.EntityData.BundleName = "cisco_ios_xr"
    propagationList.EntityData.ParentYangName = "entry"
    propagationList.EntityData.SegmentPath = "propagation_list" + "[fm_subsystem_id='" + fmt.Sprintf("%v", propagationList.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", propagationList.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", propagationList.FmFaultTag) + "']"
    propagationList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    propagationList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    propagationList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    propagationList.EntityData.Children = make(map[string]types.YChild)
    propagationList.EntityData.Leafs = make(map[string]types.YLeaf)
    propagationList.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", propagationList.FmSubsystemId}
    propagationList.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", propagationList.FmFaultType}
    propagationList.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", propagationList.FmFaultTag}
    propagationList.EntityData.Leafs["remote_agent_id"] = types.YLeaf{"RemoteAgentId", propagationList.RemoteAgentId}
    return &(propagationList.EntityData)
}

// Fm_Agents_FmTable_Entry_NotificationList
// Notification list of fault agents.
type Fm_Agents_FmTable_Entry_NotificationList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The remote agent id assocaited with this fault. The type is string.
    RemoteAgentId interface{}
}

func (notificationList *Fm_Agents_FmTable_Entry_NotificationList) GetEntityData() *types.CommonEntityData {
    notificationList.EntityData.YFilter = notificationList.YFilter
    notificationList.EntityData.YangName = "notification_list"
    notificationList.EntityData.BundleName = "cisco_ios_xr"
    notificationList.EntityData.ParentYangName = "entry"
    notificationList.EntityData.SegmentPath = "notification_list" + "[fm_subsystem_id='" + fmt.Sprintf("%v", notificationList.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", notificationList.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", notificationList.FmFaultTag) + "']"
    notificationList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationList.EntityData.Children = make(map[string]types.YChild)
    notificationList.EntityData.Leafs = make(map[string]types.YLeaf)
    notificationList.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", notificationList.FmSubsystemId}
    notificationList.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", notificationList.FmFaultType}
    notificationList.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", notificationList.FmFaultTag}
    notificationList.EntityData.Leafs["remote_agent_id"] = types.YLeaf{"RemoteAgentId", notificationList.RemoteAgentId}
    return &(notificationList.EntityData)
}

// Fm_Agents_FmTable_Entry_EscalationList
// escalation list of fault agents.
type Fm_Agents_FmTable_Entry_EscalationList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The remote agent id assocaited with this fault. The type is string.
    RemoteAgentId interface{}
}

func (escalationList *Fm_Agents_FmTable_Entry_EscalationList) GetEntityData() *types.CommonEntityData {
    escalationList.EntityData.YFilter = escalationList.YFilter
    escalationList.EntityData.YangName = "escalation_list"
    escalationList.EntityData.BundleName = "cisco_ios_xr"
    escalationList.EntityData.ParentYangName = "entry"
    escalationList.EntityData.SegmentPath = "escalation_list" + "[fm_subsystem_id='" + fmt.Sprintf("%v", escalationList.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", escalationList.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", escalationList.FmFaultTag) + "']"
    escalationList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    escalationList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    escalationList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    escalationList.EntityData.Children = make(map[string]types.YChild)
    escalationList.EntityData.Leafs = make(map[string]types.YLeaf)
    escalationList.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", escalationList.FmSubsystemId}
    escalationList.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", escalationList.FmFaultType}
    escalationList.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", escalationList.FmFaultTag}
    escalationList.EntityData.Leafs["remote_agent_id"] = types.YLeaf{"RemoteAgentId", escalationList.RemoteAgentId}
    return &(escalationList.EntityData)
}

// Fm_Agents_FmTable_Entry_Faults
type Fm_Agents_FmTable_Entry_Faults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Active Fm_Agents_FmTable_Entry_Faults_Active

    
    History Fm_Agents_FmTable_Entry_Faults_History
}

func (faults *Fm_Agents_FmTable_Entry_Faults) GetEntityData() *types.CommonEntityData {
    faults.EntityData.YFilter = faults.YFilter
    faults.EntityData.YangName = "faults"
    faults.EntityData.BundleName = "cisco_ios_xr"
    faults.EntityData.ParentYangName = "entry"
    faults.EntityData.SegmentPath = "faults"
    faults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faults.EntityData.Children = make(map[string]types.YChild)
    faults.EntityData.Children["active"] = types.YChild{"Active", &faults.Active}
    faults.EntityData.Children["history"] = types.YChild{"History", &faults.History}
    faults.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(faults.EntityData)
}

// Fm_Agents_FmTable_Entry_Faults_Active
type Fm_Agents_FmTable_Entry_Faults_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents_FmTable_Entry_Faults_Active_Brief.
    Brief []Fm_Agents_FmTable_Entry_Faults_Active_Brief

    // The type is slice of Fm_Agents_FmTable_Entry_Faults_Active_Detail.
    Detail []Fm_Agents_FmTable_Entry_Faults_Active_Detail
}

func (active *Fm_Agents_FmTable_Entry_Faults_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "faults"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Children["brief"] = types.YChild{"Brief", nil}
    for i := range active.Brief {
        active.EntityData.Children[types.GetSegmentPath(&active.Brief[i])] = types.YChild{"Brief", &active.Brief[i]}
    }
    active.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range active.Detail {
        active.EntityData.Children[types.GetSegmentPath(&active.Detail[i])] = types.YChild{"Detail", &active.Detail[i]}
    }
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(active.EntityData)
}

// Fm_Agents_FmTable_Entry_Faults_Active_Brief
type Fm_Agents_FmTable_Entry_Faults_Active_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The fault object ID. The type is string.
    ObjectId interface{}

    // The fault occurence timestamp. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    FaultTimestamp interface{}
}

func (brief *Fm_Agents_FmTable_Entry_Faults_Active_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "active"
    brief.EntityData.SegmentPath = "brief" + "[object_id='" + fmt.Sprintf("%v", brief.ObjectId) + "']"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["object_id"] = types.YLeaf{"ObjectId", brief.ObjectId}
    brief.EntityData.Leafs["fault_timestamp"] = types.YLeaf{"FaultTimestamp", brief.FaultTimestamp}
    return &(brief.EntityData)
}

// Fm_Agents_FmTable_Entry_Faults_Active_Detail
type Fm_Agents_FmTable_Entry_Faults_Active_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The fault object ID. The type is string.
    ObjectId interface{}

    // Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // Fault type identifier. The type is string.
    FmFaultType interface{}

    // Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The severity of the fault reported out. The type is interface{} with range:
    // 0..65535.
    FaultSeverity interface{}

    // The state of the fault. The type is FmFaultStateT.
    FaultState interface{}

    // The agent id associated with the fault. The type is string.
    FaultAgentId interface{}

    // The fault occurence timestamp. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    FaultTimestamp interface{}

    // The state of the timer associated with this fault. The type is bool.
    FaultTimerState interface{}

    // The fault is acted on. The type is bool.
    FaultProcessed interface{}

    // The result of the mitigation action on the fault. The type is
    // FmActionResultT.
    MitigationResult interface{}

    // The result of the recovery action on the fault. The type is
    // FmActionResultT.
    RecoveryResult interface{}

    // The result of the correlation action on the fault. The type is
    // FmActionResultT.
    CorrelationResult interface{}

    // The result of the alarm action on the fault. The type is FmActionResultT.
    AlarmResult interface{}

    // The result of the default action on the fault. The type is FmActionResultT.
    DefaultResult interface{}

    // The length of opaque data bytes. The type is interface{} with range:
    // 0..65535.
    OpaqueDataLen interface{}

    // The occurrence count of the fault. The type is interface{} with range:
    // 0..18446744073709551615.
    OccurrenceCount interface{}

    // The history state of the fault. The type is FmHistoryStateT.
    HistoryState interface{}
}

func (detail *Fm_Agents_FmTable_Entry_Faults_Active_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "active"
    detail.EntityData.SegmentPath = "detail" + "[object_id='" + fmt.Sprintf("%v", detail.ObjectId) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["object_id"] = types.YLeaf{"ObjectId", detail.ObjectId}
    detail.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", detail.FmSubsystemId}
    detail.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", detail.FmFaultType}
    detail.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", detail.FmFaultTag}
    detail.EntityData.Leafs["fault_severity"] = types.YLeaf{"FaultSeverity", detail.FaultSeverity}
    detail.EntityData.Leafs["fault_state"] = types.YLeaf{"FaultState", detail.FaultState}
    detail.EntityData.Leafs["fault_agent_id"] = types.YLeaf{"FaultAgentId", detail.FaultAgentId}
    detail.EntityData.Leafs["fault_timestamp"] = types.YLeaf{"FaultTimestamp", detail.FaultTimestamp}
    detail.EntityData.Leafs["fault_timer_state"] = types.YLeaf{"FaultTimerState", detail.FaultTimerState}
    detail.EntityData.Leafs["fault_processed"] = types.YLeaf{"FaultProcessed", detail.FaultProcessed}
    detail.EntityData.Leafs["mitigation_result"] = types.YLeaf{"MitigationResult", detail.MitigationResult}
    detail.EntityData.Leafs["recovery_result"] = types.YLeaf{"RecoveryResult", detail.RecoveryResult}
    detail.EntityData.Leafs["correlation_result"] = types.YLeaf{"CorrelationResult", detail.CorrelationResult}
    detail.EntityData.Leafs["alarm_result"] = types.YLeaf{"AlarmResult", detail.AlarmResult}
    detail.EntityData.Leafs["default_result"] = types.YLeaf{"DefaultResult", detail.DefaultResult}
    detail.EntityData.Leafs["opaque_data_len"] = types.YLeaf{"OpaqueDataLen", detail.OpaqueDataLen}
    detail.EntityData.Leafs["occurrence_count"] = types.YLeaf{"OccurrenceCount", detail.OccurrenceCount}
    detail.EntityData.Leafs["history_state"] = types.YLeaf{"HistoryState", detail.HistoryState}
    return &(detail.EntityData)
}

// Fm_Agents_FmTable_Entry_Faults_History
type Fm_Agents_FmTable_Entry_Faults_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents_FmTable_Entry_Faults_History_Brief.
    Brief []Fm_Agents_FmTable_Entry_Faults_History_Brief

    // The type is slice of Fm_Agents_FmTable_Entry_Faults_History_Detail.
    Detail []Fm_Agents_FmTable_Entry_Faults_History_Detail
}

func (history *Fm_Agents_FmTable_Entry_Faults_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "faults"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["brief"] = types.YChild{"Brief", nil}
    for i := range history.Brief {
        history.EntityData.Children[types.GetSegmentPath(&history.Brief[i])] = types.YChild{"Brief", &history.Brief[i]}
    }
    history.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range history.Detail {
        history.EntityData.Children[types.GetSegmentPath(&history.Detail[i])] = types.YChild{"Detail", &history.Detail[i]}
    }
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(history.EntityData)
}

// Fm_Agents_FmTable_Entry_Faults_History_Brief
type Fm_Agents_FmTable_Entry_Faults_History_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The fault object ID. The type is string.
    ObjectId interface{}

    // The fault occurence timestamp. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    FaultTimestamp interface{}
}

func (brief *Fm_Agents_FmTable_Entry_Faults_History_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "history"
    brief.EntityData.SegmentPath = "brief" + "[object_id='" + fmt.Sprintf("%v", brief.ObjectId) + "']"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["object_id"] = types.YLeaf{"ObjectId", brief.ObjectId}
    brief.EntityData.Leafs["fault_timestamp"] = types.YLeaf{"FaultTimestamp", brief.FaultTimestamp}
    return &(brief.EntityData)
}

// Fm_Agents_FmTable_Entry_Faults_History_Detail
type Fm_Agents_FmTable_Entry_Faults_History_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The fault object ID. The type is string.
    ObjectId interface{}

    // Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // Fault type identifier. The type is string.
    FmFaultType interface{}

    // Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The severity of the fault reported out. The type is interface{} with range:
    // 0..65535.
    FaultSeverity interface{}

    // The state of the fault. The type is FmFaultStateT.
    FaultState interface{}

    // The agent id associated with the fault. The type is string.
    FaultAgentId interface{}

    // The fault occurence timestamp. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    FaultTimestamp interface{}

    // The state of the timer associated with this fault. The type is bool.
    FaultTimerState interface{}

    // The fault is acted on. The type is bool.
    FaultProcessed interface{}

    // The result of the mitigation action on the fault. The type is
    // FmActionResultT.
    MitigationResult interface{}

    // The result of the recovery action on the fault. The type is
    // FmActionResultT.
    RecoveryResult interface{}

    // The result of the correlation action on the fault. The type is
    // FmActionResultT.
    CorrelationResult interface{}

    // The result of the alarm action on the fault. The type is FmActionResultT.
    AlarmResult interface{}

    // The result of the default action on the fault. The type is FmActionResultT.
    DefaultResult interface{}

    // The length of opaque data bytes. The type is interface{} with range:
    // 0..65535.
    OpaqueDataLen interface{}

    // The occurrence count of the fault. The type is interface{} with range:
    // 0..18446744073709551615.
    OccurrenceCount interface{}

    // The history state of the fault. The type is FmHistoryStateT.
    HistoryState interface{}
}

func (detail *Fm_Agents_FmTable_Entry_Faults_History_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "history"
    detail.EntityData.SegmentPath = "detail" + "[object_id='" + fmt.Sprintf("%v", detail.ObjectId) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["object_id"] = types.YLeaf{"ObjectId", detail.ObjectId}
    detail.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", detail.FmSubsystemId}
    detail.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", detail.FmFaultType}
    detail.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", detail.FmFaultTag}
    detail.EntityData.Leafs["fault_severity"] = types.YLeaf{"FaultSeverity", detail.FaultSeverity}
    detail.EntityData.Leafs["fault_state"] = types.YLeaf{"FaultState", detail.FaultState}
    detail.EntityData.Leafs["fault_agent_id"] = types.YLeaf{"FaultAgentId", detail.FaultAgentId}
    detail.EntityData.Leafs["fault_timestamp"] = types.YLeaf{"FaultTimestamp", detail.FaultTimestamp}
    detail.EntityData.Leafs["fault_timer_state"] = types.YLeaf{"FaultTimerState", detail.FaultTimerState}
    detail.EntityData.Leafs["fault_processed"] = types.YLeaf{"FaultProcessed", detail.FaultProcessed}
    detail.EntityData.Leafs["mitigation_result"] = types.YLeaf{"MitigationResult", detail.MitigationResult}
    detail.EntityData.Leafs["recovery_result"] = types.YLeaf{"RecoveryResult", detail.RecoveryResult}
    detail.EntityData.Leafs["correlation_result"] = types.YLeaf{"CorrelationResult", detail.CorrelationResult}
    detail.EntityData.Leafs["alarm_result"] = types.YLeaf{"AlarmResult", detail.AlarmResult}
    detail.EntityData.Leafs["default_result"] = types.YLeaf{"DefaultResult", detail.DefaultResult}
    detail.EntityData.Leafs["opaque_data_len"] = types.YLeaf{"OpaqueDataLen", detail.OpaqueDataLen}
    detail.EntityData.Leafs["occurrence_count"] = types.YLeaf{"OccurrenceCount", detail.OccurrenceCount}
    detail.EntityData.Leafs["history_state"] = types.YLeaf{"HistoryState", detail.HistoryState}
    return &(detail.EntityData)
}

// Fm_Agents_FmTable_Entry_WaitingList
type Fm_Agents_FmTable_Entry_WaitingList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents_FmTable_Entry_WaitingList_Brief.
    Brief []Fm_Agents_FmTable_Entry_WaitingList_Brief

    // The type is slice of Fm_Agents_FmTable_Entry_WaitingList_Entry.
    Entry []Fm_Agents_FmTable_Entry_WaitingList_Entry_
}

func (waitingList *Fm_Agents_FmTable_Entry_WaitingList) GetEntityData() *types.CommonEntityData {
    waitingList.EntityData.YFilter = waitingList.YFilter
    waitingList.EntityData.YangName = "waiting_list"
    waitingList.EntityData.BundleName = "cisco_ios_xr"
    waitingList.EntityData.ParentYangName = "entry"
    waitingList.EntityData.SegmentPath = "waiting_list"
    waitingList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    waitingList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    waitingList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    waitingList.EntityData.Children = make(map[string]types.YChild)
    waitingList.EntityData.Children["brief"] = types.YChild{"Brief", nil}
    for i := range waitingList.Brief {
        waitingList.EntityData.Children[types.GetSegmentPath(&waitingList.Brief[i])] = types.YChild{"Brief", &waitingList.Brief[i]}
    }
    waitingList.EntityData.Children["entry"] = types.YChild{"Entry", nil}
    for i := range waitingList.Entry {
        waitingList.EntityData.Children[types.GetSegmentPath(&waitingList.Entry[i])] = types.YChild{"Entry", &waitingList.Entry[i]}
    }
    waitingList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(waitingList.EntityData)
}

// Fm_Agents_FmTable_Entry_WaitingList_Brief
type Fm_Agents_FmTable_Entry_WaitingList_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The object Id of the entity that generated the fault. The type is string.
    ObjectId interface{}

    // The timestamp at which the fault occurred. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    FaultTimestamp interface{}

    // The state pf tje causal fault. The type is FmFaultStateT.
    FaultState interface{}

    // Count of occurrence of the fault event. The type is interface{} with range:
    // 0..18446744073709551615.
    FaultCount interface{}

    // FM correlation engine flag, internal. The type is interface{} with range:
    // 0..18446744073709551615.
    FaultFlag interface{}
}

func (brief *Fm_Agents_FmTable_Entry_WaitingList_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "waiting_list"
    brief.EntityData.SegmentPath = "brief" + "[fm_subsystem_id='" + fmt.Sprintf("%v", brief.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", brief.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", brief.FmFaultTag) + "']"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", brief.FmSubsystemId}
    brief.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", brief.FmFaultType}
    brief.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", brief.FmFaultTag}
    brief.EntityData.Leafs["object_id"] = types.YLeaf{"ObjectId", brief.ObjectId}
    brief.EntityData.Leafs["fault_timestamp"] = types.YLeaf{"FaultTimestamp", brief.FaultTimestamp}
    brief.EntityData.Leafs["fault_state"] = types.YLeaf{"FaultState", brief.FaultState}
    brief.EntityData.Leafs["fault_count"] = types.YLeaf{"FaultCount", brief.FaultCount}
    brief.EntityData.Leafs["fault_flag"] = types.YLeaf{"FaultFlag", brief.FaultFlag}
    return &(brief.EntityData)
}

// Fm_Agents_FmTable_Entry_WaitingList_Entry_
type Fm_Agents_FmTable_Entry_WaitingList_Entry_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The object Id of the entity that generated the fault. The type is string.
    ObjectId interface{}

    // The timestamp at which the fault occurred. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    FaultTimestamp interface{}

    // The state pf tje causal fault. The type is FmFaultStateT.
    FaultState interface{}

    // Count of occurrence of the fault event. The type is interface{} with range:
    // 0..18446744073709551615.
    FaultCount interface{}

    // FM correlation engine flag, internal. The type is interface{} with range:
    // 0..18446744073709551615.
    FaultFlag interface{}
}

func (entry_ *Fm_Agents_FmTable_Entry_WaitingList_Entry_) GetEntityData() *types.CommonEntityData {
    entry_.EntityData.YFilter = entry_.YFilter
    entry_.EntityData.YangName = "entry"
    entry_.EntityData.BundleName = "cisco_ios_xr"
    entry_.EntityData.ParentYangName = "waiting_list"
    entry_.EntityData.SegmentPath = "entry" + "[fm_subsystem_id='" + fmt.Sprintf("%v", entry_.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", entry_.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", entry_.FmFaultTag) + "']"
    entry_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry_.EntityData.Children = make(map[string]types.YChild)
    entry_.EntityData.Leafs = make(map[string]types.YLeaf)
    entry_.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", entry_.FmSubsystemId}
    entry_.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", entry_.FmFaultType}
    entry_.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", entry_.FmFaultTag}
    entry_.EntityData.Leafs["object_id"] = types.YLeaf{"ObjectId", entry_.ObjectId}
    entry_.EntityData.Leafs["fault_timestamp"] = types.YLeaf{"FaultTimestamp", entry_.FaultTimestamp}
    entry_.EntityData.Leafs["fault_state"] = types.YLeaf{"FaultState", entry_.FaultState}
    entry_.EntityData.Leafs["fault_count"] = types.YLeaf{"FaultCount", entry_.FaultCount}
    entry_.EntityData.Leafs["fault_flag"] = types.YLeaf{"FaultFlag", entry_.FaultFlag}
    return &(entry_.EntityData)
}

// Fm_Agents_FmInternals
type Fm_Agents_FmInternals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents_FmInternals_Detail.
    Detail []Fm_Agents_FmInternals_Detail
}

func (fmInternals *Fm_Agents_FmInternals) GetEntityData() *types.CommonEntityData {
    fmInternals.EntityData.YFilter = fmInternals.YFilter
    fmInternals.EntityData.YangName = "fm_internals"
    fmInternals.EntityData.BundleName = "cisco_ios_xr"
    fmInternals.EntityData.ParentYangName = "agents"
    fmInternals.EntityData.SegmentPath = "fm_internals"
    fmInternals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmInternals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmInternals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmInternals.EntityData.Children = make(map[string]types.YChild)
    fmInternals.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range fmInternals.Detail {
        fmInternals.EntityData.Children[types.GetSegmentPath(&fmInternals.Detail[i])] = types.YChild{"Detail", &fmInternals.Detail[i]}
    }
    fmInternals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fmInternals.EntityData)
}

// Fm_Agents_FmInternals_Detail
type Fm_Agents_FmInternals_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // Common action data. The type is interface{} with range: 0..65535.
    CommonAction interface{}

    // opaque action data. The type is interface{} with range: 0..65535.
    OpaqueAction interface{}

    // Pointer to the mitigation callback function. The type is interface{} with
    // range: 0..18446744073709551615.
    MitigationCb interface{}

    // Pointer to the recovery callback function. The type is interface{} with
    // range: 0..18446744073709551615.
    RecoveryCb interface{}

    // Flag indicates if alarm severity is dirty. The type is bool.
    AlarmSeverityDirty interface{}

    // Flag indicates all actions are disabled. The type is bool.
    DisableAction interface{}

    // Flag indicates all actions are repeated. The type is bool.
    RepeatAction interface{}

    // Flag indicates if causal list is present. The type is bool.
    HasCausalList interface{}

    // The parser tag of the XML parser. The type is interface{} with range:
    // 0..18446744073709551615.
    ParserTag interface{}

    // The parser tag string. The type is string.
    ParserTagString interface{}

    // list of fault rule declaring callbacks. The type is slice of
    // Fm_Agents_FmInternals_Detail_Rules.
    Rules []Fm_Agents_FmInternals_Detail_Rules
}

func (detail *Fm_Agents_FmInternals_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "fm_internals"
    detail.EntityData.SegmentPath = "detail" + "[fm_subsystem_id='" + fmt.Sprintf("%v", detail.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", detail.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", detail.FmFaultTag) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["rules"] = types.YChild{"Rules", nil}
    for i := range detail.Rules {
        detail.EntityData.Children[types.GetSegmentPath(&detail.Rules[i])] = types.YChild{"Rules", &detail.Rules[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", detail.FmSubsystemId}
    detail.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", detail.FmFaultType}
    detail.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", detail.FmFaultTag}
    detail.EntityData.Leafs["common_action"] = types.YLeaf{"CommonAction", detail.CommonAction}
    detail.EntityData.Leafs["opaque_action"] = types.YLeaf{"OpaqueAction", detail.OpaqueAction}
    detail.EntityData.Leafs["mitigation_cb"] = types.YLeaf{"MitigationCb", detail.MitigationCb}
    detail.EntityData.Leafs["recovery_cb"] = types.YLeaf{"RecoveryCb", detail.RecoveryCb}
    detail.EntityData.Leafs["alarm_severity_dirty"] = types.YLeaf{"AlarmSeverityDirty", detail.AlarmSeverityDirty}
    detail.EntityData.Leafs["disable_action"] = types.YLeaf{"DisableAction", detail.DisableAction}
    detail.EntityData.Leafs["repeat_action"] = types.YLeaf{"RepeatAction", detail.RepeatAction}
    detail.EntityData.Leafs["has_causal_list"] = types.YLeaf{"HasCausalList", detail.HasCausalList}
    detail.EntityData.Leafs["parser_tag"] = types.YLeaf{"ParserTag", detail.ParserTag}
    detail.EntityData.Leafs["parser_tag_string"] = types.YLeaf{"ParserTagString", detail.ParserTagString}
    return &(detail.EntityData)
}

// Fm_Agents_FmInternals_Detail_Rules
// list of fault rule declaring callbacks
type Fm_Agents_FmInternals_Detail_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The location associated with the fault. The type is string.
    FaultLocation interface{}

    // The callback function that declares the fault. The type is interface{} with
    // range: 0..18446744073709551615.
    RuleCb interface{}
}

func (rules *Fm_Agents_FmInternals_Detail_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "detail"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = make(map[string]types.YChild)
    rules.EntityData.Leafs = make(map[string]types.YLeaf)
    rules.EntityData.Leafs["fault_location"] = types.YLeaf{"FaultLocation", rules.FaultLocation}
    rules.EntityData.Leafs["rule_cb"] = types.YLeaf{"RuleCb", rules.RuleCb}
    return &(rules.EntityData)
}

// Fm_Agents_FmAlarmMapping
type Fm_Agents_FmAlarmMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents_FmAlarmMapping_Detail.
    Detail []Fm_Agents_FmAlarmMapping_Detail
}

func (fmAlarmMapping *Fm_Agents_FmAlarmMapping) GetEntityData() *types.CommonEntityData {
    fmAlarmMapping.EntityData.YFilter = fmAlarmMapping.YFilter
    fmAlarmMapping.EntityData.YangName = "fm_alarm_mapping"
    fmAlarmMapping.EntityData.BundleName = "cisco_ios_xr"
    fmAlarmMapping.EntityData.ParentYangName = "agents"
    fmAlarmMapping.EntityData.SegmentPath = "fm_alarm_mapping"
    fmAlarmMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmAlarmMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmAlarmMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmAlarmMapping.EntityData.Children = make(map[string]types.YChild)
    fmAlarmMapping.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range fmAlarmMapping.Detail {
        fmAlarmMapping.EntityData.Children[types.GetSegmentPath(&fmAlarmMapping.Detail[i])] = types.YChild{"Detail", &fmAlarmMapping.Detail[i]}
    }
    fmAlarmMapping.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fmAlarmMapping.EntityData)
}

// Fm_Agents_FmAlarmMapping_Detail
type Fm_Agents_FmAlarmMapping_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // The alarm grouping for this fault. The type is interface{} with range:
    // 0..65535.
    AlarmGroup interface{}

    // The alarm severity for this fault. The type is interface{} with range:
    // 0..65535.
    AlarmSeverity interface{}
}

func (detail *Fm_Agents_FmAlarmMapping_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "fm_alarm_mapping"
    detail.EntityData.SegmentPath = "detail" + "[fm_subsystem_id='" + fmt.Sprintf("%v", detail.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", detail.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", detail.FmFaultTag) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", detail.FmSubsystemId}
    detail.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", detail.FmFaultType}
    detail.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", detail.FmFaultTag}
    detail.EntityData.Leafs["alarm_group"] = types.YLeaf{"AlarmGroup", detail.AlarmGroup}
    detail.EntityData.Leafs["alarm_severity"] = types.YLeaf{"AlarmSeverity", detail.AlarmSeverity}
    return &(detail.EntityData)
}

// Fm_Agents_FmStatistics
type Fm_Agents_FmStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Fm_Agents_FmStatistics_Detail.
    Detail []Fm_Agents_FmStatistics_Detail
}

func (fmStatistics *Fm_Agents_FmStatistics) GetEntityData() *types.CommonEntityData {
    fmStatistics.EntityData.YFilter = fmStatistics.YFilter
    fmStatistics.EntityData.YangName = "fm_statistics"
    fmStatistics.EntityData.BundleName = "cisco_ios_xr"
    fmStatistics.EntityData.ParentYangName = "agents"
    fmStatistics.EntityData.SegmentPath = "fm_statistics"
    fmStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmStatistics.EntityData.Children = make(map[string]types.YChild)
    fmStatistics.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range fmStatistics.Detail {
        fmStatistics.EntityData.Children[types.GetSegmentPath(&fmStatistics.Detail[i])] = types.YChild{"Detail", &fmStatistics.Detail[i]}
    }
    fmStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fmStatistics.EntityData)
}

// Fm_Agents_FmStatistics_Detail
type Fm_Agents_FmStatistics_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fault sub-system identifier. The type is string.
    FmSubsystemId interface{}

    // This attribute is a key. Fault type identifier. The type is string.
    FmFaultType interface{}

    // This attribute is a key. Fault tag identifier. The type is string.
    FmFaultTag interface{}

    // Threshold count for the fault. The type is interface{} with range:
    // 0..4294967295.
    ThresholdCount interface{}

    // Object occurrence count. The type is interface{} with range: 0..4294967295.
    AllObjectOccurCount interface{}

    // Number of times the fault is declared. The type is interface{} with range:
    // 0..4294967295.
    DeclaredCount interface{}

    // Number of times the fault is cleared. The type is interface{} with range:
    // 0..4294967295.
    ClearedCount interface{}

    // Number of times the fault is info. The type is interface{} with range:
    // 0..4294967295.
    InfoCount interface{}

    // The hold time in ms for soaking the fault. The type is interface{} with
    // range: 0..4294967295.
    HoldTime interface{}
}

func (detail *Fm_Agents_FmStatistics_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "fm_statistics"
    detail.EntityData.SegmentPath = "detail" + "[fm_subsystem_id='" + fmt.Sprintf("%v", detail.FmSubsystemId) + "']" + "[fm_fault_type='" + fmt.Sprintf("%v", detail.FmFaultType) + "']" + "[fm_fault_tag='" + fmt.Sprintf("%v", detail.FmFaultTag) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["fm_subsystem_id"] = types.YLeaf{"FmSubsystemId", detail.FmSubsystemId}
    detail.EntityData.Leafs["fm_fault_type"] = types.YLeaf{"FmFaultType", detail.FmFaultType}
    detail.EntityData.Leafs["fm_fault_tag"] = types.YLeaf{"FmFaultTag", detail.FmFaultTag}
    detail.EntityData.Leafs["threshold_count"] = types.YLeaf{"ThresholdCount", detail.ThresholdCount}
    detail.EntityData.Leafs["all_object_occur_count"] = types.YLeaf{"AllObjectOccurCount", detail.AllObjectOccurCount}
    detail.EntityData.Leafs["declared_count"] = types.YLeaf{"DeclaredCount", detail.DeclaredCount}
    detail.EntityData.Leafs["cleared_count"] = types.YLeaf{"ClearedCount", detail.ClearedCount}
    detail.EntityData.Leafs["info_count"] = types.YLeaf{"InfoCount", detail.InfoCount}
    detail.EntityData.Leafs["hold_time"] = types.YLeaf{"HoldTime", detail.HoldTime}
    return &(detail.EntityData)
}

