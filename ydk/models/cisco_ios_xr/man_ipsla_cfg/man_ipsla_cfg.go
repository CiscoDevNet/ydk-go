// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-ipsla package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipsla: IPSLA configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package man_ipsla_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_ipsla_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-ipsla-cfg ipsla}", reflect.TypeOf(Ipsla{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-ipsla-cfg:ipsla", reflect.TypeOf(Ipsla{}))
}

// IpslaSecondaryFrequency represents Ipsla secondary frequency
type IpslaSecondaryFrequency string

const (
    // Enable secondary frequency for connection loss
    IpslaSecondaryFrequency_connection_loss IpslaSecondaryFrequency = "connection-loss"

    // Enable secondary frequency for timeout
    IpslaSecondaryFrequency_timeout IpslaSecondaryFrequency = "timeout"

    // Enable secondary frequency for timeout and
    // connection loss
    IpslaSecondaryFrequency_both IpslaSecondaryFrequency = "both"
)

// IpslaMonth represents Ipsla month
type IpslaMonth string

const (
    // January
    IpslaMonth_january IpslaMonth = "january"

    // February
    IpslaMonth_february IpslaMonth = "february"

    // March
    IpslaMonth_march IpslaMonth = "march"

    // April
    IpslaMonth_april IpslaMonth = "april"

    // May
    IpslaMonth_may IpslaMonth = "may"

    // June
    IpslaMonth_june IpslaMonth = "june"

    // July
    IpslaMonth_july IpslaMonth = "july"

    // August
    IpslaMonth_august IpslaMonth = "august"

    // September
    IpslaMonth_september IpslaMonth = "september"

    // October
    IpslaMonth_october IpslaMonth = "october"

    // November
    IpslaMonth_november IpslaMonth = "november"

    // December
    IpslaMonth_december IpslaMonth = "december"
)

// IpslaLspPingReplyMode represents Ipsla lsp ping reply mode
type IpslaLspPingReplyMode string

const (
    // Send replies via IPv4 UDP packets with Router
    // Alert option
    IpslaLspPingReplyMode_ipv4_udp_router_alert IpslaLspPingReplyMode = "ipv4-udp-router-alert"

    // Send replies via a Control Channel option
    IpslaLspPingReplyMode_control_channel IpslaLspPingReplyMode = "control-channel"
)

// IpslaLspTraceReplyMode represents Ipsla lsp trace reply mode
type IpslaLspTraceReplyMode string

const (
    // Send replies via IPv4 UDP packets with Router
    // Alert option
    IpslaLspTraceReplyMode_ipv4_udp_router_alert IpslaLspTraceReplyMode = "ipv4-udp-router-alert"
)

// IpslaLspMonitorReplyMode represents Ipsla lsp monitor reply mode
type IpslaLspMonitorReplyMode string

const (
    // Send replies via IPv4 UDP packets with Router
    // Alert option
    IpslaLspMonitorReplyMode_ipv4_udp_router_alert IpslaLspMonitorReplyMode = "ipv4-udp-router-alert"
)

// IpslaSched represents Ipsla sched
type IpslaSched string

const (
    // Schedule pending for later time
    IpslaSched_pending IpslaSched = "pending"

    // Schedule operation now
    IpslaSched_now IpslaSched = "now"

    // Schedule operation after specifed duration
    IpslaSched_after IpslaSched = "after"

    // Schedule operation at specified time
    IpslaSched_at IpslaSched = "at"
)

// IpslaLspReplyDscp represents Ipsla lsp reply dscp
type IpslaLspReplyDscp string

const (
    // Match packets with default dscp (000000)
    IpslaLspReplyDscp_default_ IpslaLspReplyDscp = "default"

    // Match packets with AF11 dscp (001010)
    IpslaLspReplyDscp_af11 IpslaLspReplyDscp = "af11"

    // Match packets with AF12 dscp (001100)
    IpslaLspReplyDscp_af12 IpslaLspReplyDscp = "af12"

    // Match packets with AF13 dscp (001110)
    IpslaLspReplyDscp_af13 IpslaLspReplyDscp = "af13"

    // Match packets with AF21 dscp (010010)
    IpslaLspReplyDscp_af21 IpslaLspReplyDscp = "af21"

    // Match packets with AF22 dscp (010100)
    IpslaLspReplyDscp_af22 IpslaLspReplyDscp = "af22"

    // Match packets with AF23 dscp (010110)
    IpslaLspReplyDscp_af23 IpslaLspReplyDscp = "af23"

    // Match packets with AF31 dscp (011010)
    IpslaLspReplyDscp_af31 IpslaLspReplyDscp = "af31"

    // Match packets with AF32 dscp (011100)
    IpslaLspReplyDscp_af32 IpslaLspReplyDscp = "af32"

    // Match packets with AF33 dscp (011110)
    IpslaLspReplyDscp_af33 IpslaLspReplyDscp = "af33"

    // Match packets with AF41 dscp (100010)
    IpslaLspReplyDscp_af41 IpslaLspReplyDscp = "af41"

    // Match packets with AF42 dscp (100100)
    IpslaLspReplyDscp_af42 IpslaLspReplyDscp = "af42"

    // Match packets with AF43 dscp (100110)
    IpslaLspReplyDscp_af43 IpslaLspReplyDscp = "af43"

    // Match packets with CS1(precedence 1) dscp
    // (001000)
    IpslaLspReplyDscp_cs1 IpslaLspReplyDscp = "cs1"

    // Match packets with CS2(precedence 2) dscp
    // (010000)
    IpslaLspReplyDscp_cs2 IpslaLspReplyDscp = "cs2"

    // Match packets with CS3(precedence 3) dscp
    // (011000)
    IpslaLspReplyDscp_cs3 IpslaLspReplyDscp = "cs3"

    // Match packets with CS4(precedence 4) dscp
    // (100000)
    IpslaLspReplyDscp_cs4 IpslaLspReplyDscp = "cs4"

    // Match packets with CS5(precedence 5) dscp
    // (101000)
    IpslaLspReplyDscp_cs5 IpslaLspReplyDscp = "cs5"

    // Match packets with CS6(precedence 6) dscp
    // (110000)
    IpslaLspReplyDscp_cs6 IpslaLspReplyDscp = "cs6"

    // Match packets with CS7(precedence 7) dscp
    // (111000)
    IpslaLspReplyDscp_cs7 IpslaLspReplyDscp = "cs7"

    // Match packets with EF dscp (101110)
    IpslaLspReplyDscp_ef IpslaLspReplyDscp = "ef"
)

// IpslaLife represents Ipsla life
type IpslaLife string

const (
    // Schedule operation to run forever
    IpslaLife_forever IpslaLife = "forever"
)

// IpslaThresholdTypes represents Ipsla threshold types
type IpslaThresholdTypes string

const (
    // Take action immediately after threshold is
    // violated
    IpslaThresholdTypes_immediate IpslaThresholdTypes = "immediate"

    // Take action after N consecutive threshold
    // violations
    IpslaThresholdTypes_consecutive IpslaThresholdTypes = "consecutive"

    // Take action after X violations in Y probes
    IpslaThresholdTypes_xof_y IpslaThresholdTypes = "xof-y"

    // Take action if average of N probes violates the
    // threshold
    IpslaThresholdTypes_average IpslaThresholdTypes = "average"
)

// IpslaLspMonitorThresholdTypes represents Ipsla lsp monitor threshold types
type IpslaLspMonitorThresholdTypes string

const (
    // Take action immediately after threshold is
    // violated
    IpslaLspMonitorThresholdTypes_immediate IpslaLspMonitorThresholdTypes = "immediate"

    // Take action after N consecutive threshold
    // violations
    IpslaLspMonitorThresholdTypes_consecutive IpslaLspMonitorThresholdTypes = "consecutive"
)

// IpslaHistoryFilter represents Ipsla history filter
type IpslaHistoryFilter string

const (
    // Store data for failed operations
    IpslaHistoryFilter_failed IpslaHistoryFilter = "failed"

    // Store data for all operations
    IpslaHistoryFilter_all IpslaHistoryFilter = "all"
)

// Ipsla
// IPSLA configuration
type Ipsla struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPSLA application common configuration.
    Common Ipsla_Common

    // MPLS LSP Monitor(MPLSLM) configuration.
    MplsLspMonitor Ipsla_MplsLspMonitor

    // IPSLA Operation configuration.
    Operation Ipsla_Operation

    // Responder configuration.
    Responder Ipsla_Responder

    // Provider Edge(PE) discovery configuration.
    MplsDiscovery Ipsla_MplsDiscovery
}

func (ipsla *Ipsla) GetEntityData() *types.CommonEntityData {
    ipsla.EntityData.YFilter = ipsla.YFilter
    ipsla.EntityData.YangName = "ipsla"
    ipsla.EntityData.BundleName = "cisco_ios_xr"
    ipsla.EntityData.ParentYangName = "Cisco-IOS-XR-man-ipsla-cfg"
    ipsla.EntityData.SegmentPath = "Cisco-IOS-XR-man-ipsla-cfg:ipsla"
    ipsla.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipsla.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipsla.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipsla.EntityData.Children = make(map[string]types.YChild)
    ipsla.EntityData.Children["common"] = types.YChild{"Common", &ipsla.Common}
    ipsla.EntityData.Children["mpls-lsp-monitor"] = types.YChild{"MplsLspMonitor", &ipsla.MplsLspMonitor}
    ipsla.EntityData.Children["operation"] = types.YChild{"Operation", &ipsla.Operation}
    ipsla.EntityData.Children["responder"] = types.YChild{"Responder", &ipsla.Responder}
    ipsla.EntityData.Children["mpls-discovery"] = types.YChild{"MplsDiscovery", &ipsla.MplsDiscovery}
    ipsla.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipsla.EntityData)
}

// Ipsla_Common
// IPSLA application common configuration
type Ipsla_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure low memory water mark (default 20M). The type is interface{} with
    // range: 0..4294967295. The default value is 20480.
    LowMemory interface{}

    // Authenticaion configuration.
    Authentication Ipsla_Common_Authentication
}

func (common *Ipsla_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "ipsla"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = make(map[string]types.YChild)
    common.EntityData.Children["authentication"] = types.YChild{"Authentication", &common.Authentication}
    common.EntityData.Leafs = make(map[string]types.YLeaf)
    common.EntityData.Leafs["low-memory"] = types.YLeaf{"LowMemory", common.LowMemory}
    return &(common.EntityData)
}

// Ipsla_Common_Authentication
// Authenticaion configuration
type Ipsla_Common_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use MD5 authentication for IPSLA control message. The type is string with
    // length: 1..32.
    KeyChain interface{}
}

func (authentication *Ipsla_Common_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "common"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["key-chain"] = types.YLeaf{"KeyChain", authentication.KeyChain}
    return &(authentication.EntityData)
}

// Ipsla_MplsLspMonitor
// MPLS LSP Monitor(MPLSLM) configuration
type Ipsla_MplsLspMonitor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLSLM Reaction configuration.
    Reactions Ipsla_MplsLspMonitor_Reactions

    // MPLSLM schedule configuration.
    Schedules Ipsla_MplsLspMonitor_Schedules

    // MPLS LSP Monitor definition table.
    Definitions Ipsla_MplsLspMonitor_Definitions
}

func (mplsLspMonitor *Ipsla_MplsLspMonitor) GetEntityData() *types.CommonEntityData {
    mplsLspMonitor.EntityData.YFilter = mplsLspMonitor.YFilter
    mplsLspMonitor.EntityData.YangName = "mpls-lsp-monitor"
    mplsLspMonitor.EntityData.BundleName = "cisco_ios_xr"
    mplsLspMonitor.EntityData.ParentYangName = "ipsla"
    mplsLspMonitor.EntityData.SegmentPath = "mpls-lsp-monitor"
    mplsLspMonitor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLspMonitor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLspMonitor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLspMonitor.EntityData.Children = make(map[string]types.YChild)
    mplsLspMonitor.EntityData.Children["reactions"] = types.YChild{"Reactions", &mplsLspMonitor.Reactions}
    mplsLspMonitor.EntityData.Children["schedules"] = types.YChild{"Schedules", &mplsLspMonitor.Schedules}
    mplsLspMonitor.EntityData.Children["definitions"] = types.YChild{"Definitions", &mplsLspMonitor.Definitions}
    mplsLspMonitor.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsLspMonitor.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions
// MPLSLM Reaction configuration
type Ipsla_MplsLspMonitor_Reactions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reaction configuration for an MPLSLM instance. The type is slice of
    // Ipsla_MplsLspMonitor_Reactions_Reaction.
    Reaction []Ipsla_MplsLspMonitor_Reactions_Reaction
}

func (reactions *Ipsla_MplsLspMonitor_Reactions) GetEntityData() *types.CommonEntityData {
    reactions.EntityData.YFilter = reactions.YFilter
    reactions.EntityData.YangName = "reactions"
    reactions.EntityData.BundleName = "cisco_ios_xr"
    reactions.EntityData.ParentYangName = "mpls-lsp-monitor"
    reactions.EntityData.SegmentPath = "reactions"
    reactions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reactions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reactions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reactions.EntityData.Children = make(map[string]types.YChild)
    reactions.EntityData.Children["reaction"] = types.YChild{"Reaction", nil}
    for i := range reactions.Reaction {
        reactions.EntityData.Children[types.GetSegmentPath(&reactions.Reaction[i])] = types.YChild{"Reaction", &reactions.Reaction[i]}
    }
    reactions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reactions.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction
// Reaction configuration for an MPLSLM instance
type Ipsla_MplsLspMonitor_Reactions_Reaction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor identifier. The type is interface{} with
    // range: 1..2048.
    MonitorId interface{}

    // Reaction condition specification.
    Condition Ipsla_MplsLspMonitor_Reactions_Reaction_Condition
}

func (reaction *Ipsla_MplsLspMonitor_Reactions_Reaction) GetEntityData() *types.CommonEntityData {
    reaction.EntityData.YFilter = reaction.YFilter
    reaction.EntityData.YangName = "reaction"
    reaction.EntityData.BundleName = "cisco_ios_xr"
    reaction.EntityData.ParentYangName = "reactions"
    reaction.EntityData.SegmentPath = "reaction" + "[monitor-id='" + fmt.Sprintf("%v", reaction.MonitorId) + "']"
    reaction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reaction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reaction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reaction.EntityData.Children = make(map[string]types.YChild)
    reaction.EntityData.Children["condition"] = types.YChild{"Condition", &reaction.Condition}
    reaction.EntityData.Leafs = make(map[string]types.YLeaf)
    reaction.EntityData.Leafs["monitor-id"] = types.YLeaf{"MonitorId", reaction.MonitorId}
    return &(reaction.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition
// Reaction condition specification
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // React on LPD Tree Trace violation for a monitored MPLSLM.
    LpdTreeTrace Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace

    // React on probe timeout.
    Timeout Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout

    // React on LPD Group violation for a monitored MPLSLM.
    LpdGroup Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup

    // React on connection loss for a monitored MPLSLM.
    ConnectionLoss Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss
}

func (condition *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition) GetEntityData() *types.CommonEntityData {
    condition.EntityData.YFilter = condition.YFilter
    condition.EntityData.YangName = "condition"
    condition.EntityData.BundleName = "cisco_ios_xr"
    condition.EntityData.ParentYangName = "reaction"
    condition.EntityData.SegmentPath = "condition"
    condition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    condition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    condition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    condition.EntityData.Children = make(map[string]types.YChild)
    condition.EntityData.Children["lpd-tree-trace"] = types.YChild{"LpdTreeTrace", &condition.LpdTreeTrace}
    condition.EntityData.Children["timeout"] = types.YChild{"Timeout", &condition.Timeout}
    condition.EntityData.Children["lpd-group"] = types.YChild{"LpdGroup", &condition.LpdGroup}
    condition.EntityData.Children["connection-loss"] = types.YChild{"ConnectionLoss", &condition.ConnectionLoss}
    condition.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(condition.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace
// React on LPD Tree Trace violation for a
// monitored MPLSLM
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular MPLSLM. The type is interface{}.
    Create interface{}

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace_ActionType
}

func (lpdTreeTrace *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace) GetEntityData() *types.CommonEntityData {
    lpdTreeTrace.EntityData.YFilter = lpdTreeTrace.YFilter
    lpdTreeTrace.EntityData.YangName = "lpd-tree-trace"
    lpdTreeTrace.EntityData.BundleName = "cisco_ios_xr"
    lpdTreeTrace.EntityData.ParentYangName = "condition"
    lpdTreeTrace.EntityData.SegmentPath = "lpd-tree-trace"
    lpdTreeTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdTreeTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdTreeTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdTreeTrace.EntityData.Children = make(map[string]types.YChild)
    lpdTreeTrace.EntityData.Children["action-type"] = types.YChild{"ActionType", &lpdTreeTrace.ActionType}
    lpdTreeTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    lpdTreeTrace.EntityData.Leafs["create"] = types.YLeaf{"Create", lpdTreeTrace.Create}
    return &(lpdTreeTrace.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}
}

func (actionType *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdTreeTrace_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "lpd-tree-trace"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    return &(actionType.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout
// React on probe timeout
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular MPLSLM. The type is interface{}.
    Create interface{}

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ThresholdType
}

func (timeout *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout) GetEntityData() *types.CommonEntityData {
    timeout.EntityData.YFilter = timeout.YFilter
    timeout.EntityData.YangName = "timeout"
    timeout.EntityData.BundleName = "cisco_ios_xr"
    timeout.EntityData.ParentYangName = "condition"
    timeout.EntityData.SegmentPath = "timeout"
    timeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeout.EntityData.Children = make(map[string]types.YChild)
    timeout.EntityData.Children["action-type"] = types.YChild{"ActionType", &timeout.ActionType}
    timeout.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &timeout.ThresholdType}
    timeout.EntityData.Leafs = make(map[string]types.YLeaf)
    timeout.EntityData.Leafs["create"] = types.YLeaf{"Create", timeout.Create}
    return &(timeout.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}
}

func (actionType *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "timeout"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    return &(actionType.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaLspMonitorThresholdTypes.
    ThreshType interface{}

    // Probe count for consecutive. The type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_Timeout_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "timeout"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup
// React on LPD Group violation for a monitored
// MPLSLM
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular MPLSLM. The type is interface{}.
    Create interface{}

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup_ActionType
}

func (lpdGroup *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup) GetEntityData() *types.CommonEntityData {
    lpdGroup.EntityData.YFilter = lpdGroup.YFilter
    lpdGroup.EntityData.YangName = "lpd-group"
    lpdGroup.EntityData.BundleName = "cisco_ios_xr"
    lpdGroup.EntityData.ParentYangName = "condition"
    lpdGroup.EntityData.SegmentPath = "lpd-group"
    lpdGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdGroup.EntityData.Children = make(map[string]types.YChild)
    lpdGroup.EntityData.Children["action-type"] = types.YChild{"ActionType", &lpdGroup.ActionType}
    lpdGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    lpdGroup.EntityData.Leafs["create"] = types.YLeaf{"Create", lpdGroup.Create}
    return &(lpdGroup.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}
}

func (actionType *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_LpdGroup_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "lpd-group"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    return &(actionType.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss
// React on connection loss for a monitored
// MPLSLM
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular MPLSLM. The type is interface{}.
    Create interface{}

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType
}

func (connectionLoss *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss) GetEntityData() *types.CommonEntityData {
    connectionLoss.EntityData.YFilter = connectionLoss.YFilter
    connectionLoss.EntityData.YangName = "connection-loss"
    connectionLoss.EntityData.BundleName = "cisco_ios_xr"
    connectionLoss.EntityData.ParentYangName = "condition"
    connectionLoss.EntityData.SegmentPath = "connection-loss"
    connectionLoss.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectionLoss.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectionLoss.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectionLoss.EntityData.Children = make(map[string]types.YChild)
    connectionLoss.EntityData.Children["action-type"] = types.YChild{"ActionType", &connectionLoss.ActionType}
    connectionLoss.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &connectionLoss.ThresholdType}
    connectionLoss.EntityData.Leafs = make(map[string]types.YLeaf)
    connectionLoss.EntityData.Leafs["create"] = types.YLeaf{"Create", connectionLoss.Create}
    return &(connectionLoss.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}
}

func (actionType *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "connection-loss"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    return &(actionType.EntityData)
}

// Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaLspMonitorThresholdTypes.
    ThreshType interface{}

    // Probe count for consecutive. The type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_MplsLspMonitor_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "connection-loss"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_MplsLspMonitor_Schedules
// MPLSLM schedule configuration
type Ipsla_MplsLspMonitor_Schedules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Schedule an MPLSLM instance. The type is slice of
    // Ipsla_MplsLspMonitor_Schedules_Schedule.
    Schedule []Ipsla_MplsLspMonitor_Schedules_Schedule
}

func (schedules *Ipsla_MplsLspMonitor_Schedules) GetEntityData() *types.CommonEntityData {
    schedules.EntityData.YFilter = schedules.YFilter
    schedules.EntityData.YangName = "schedules"
    schedules.EntityData.BundleName = "cisco_ios_xr"
    schedules.EntityData.ParentYangName = "mpls-lsp-monitor"
    schedules.EntityData.SegmentPath = "schedules"
    schedules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schedules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schedules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schedules.EntityData.Children = make(map[string]types.YChild)
    schedules.EntityData.Children["schedule"] = types.YChild{"Schedule", nil}
    for i := range schedules.Schedule {
        schedules.EntityData.Children[types.GetSegmentPath(&schedules.Schedule[i])] = types.YChild{"Schedule", &schedules.Schedule[i]}
    }
    schedules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(schedules.EntityData)
}

// Ipsla_MplsLspMonitor_Schedules_Schedule
// Schedule an MPLSLM instance
type Ipsla_MplsLspMonitor_Schedules_Schedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor indentifier. The type is interface{} with
    // range: 1..2048.
    MonitorId interface{}

    // Group schedule frequency of the probing. The type is interface{} with
    // range: 1..604800. Units are second.
    Frequency interface{}

    // Group schedule period range. The type is interface{} with range: 1..604800.
    // Units are second.
    Period interface{}

    // Start time of MPLSLM instance.
    StartTime Ipsla_MplsLspMonitor_Schedules_Schedule_StartTime
}

func (schedule *Ipsla_MplsLspMonitor_Schedules_Schedule) GetEntityData() *types.CommonEntityData {
    schedule.EntityData.YFilter = schedule.YFilter
    schedule.EntityData.YangName = "schedule"
    schedule.EntityData.BundleName = "cisco_ios_xr"
    schedule.EntityData.ParentYangName = "schedules"
    schedule.EntityData.SegmentPath = "schedule" + "[monitor-id='" + fmt.Sprintf("%v", schedule.MonitorId) + "']"
    schedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schedule.EntityData.Children = make(map[string]types.YChild)
    schedule.EntityData.Children["start-time"] = types.YChild{"StartTime", &schedule.StartTime}
    schedule.EntityData.Leafs = make(map[string]types.YLeaf)
    schedule.EntityData.Leafs["monitor-id"] = types.YLeaf{"MonitorId", schedule.MonitorId}
    schedule.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", schedule.Frequency}
    schedule.EntityData.Leafs["period"] = types.YLeaf{"Period", schedule.Period}
    return &(schedule.EntityData)
}

// Ipsla_MplsLspMonitor_Schedules_Schedule_StartTime
// Start time of MPLSLM instance
type Ipsla_MplsLspMonitor_Schedules_Schedule_StartTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of schedule. The type is IpslaSched.
    ScheduleType interface{}

    // Hour value(hh) in hh:mm:ss specification. The type is interface{} with
    // range: 0..23.
    Hour interface{}

    // Minute value(mm) in hh:mm:ss specification. The type is interface{} with
    // range: 0..59.
    Minute interface{}

    // Second value(ss) in hh:mm:ss specification. The type is interface{} with
    // range: 0..59.
    Second interface{}

    // Month of the year (optional. Default current month). The type is
    // IpslaMonth.
    Month interface{}

    // Day of the month(optional. Default today). The type is interface{} with
    // range: 1..31.
    Day interface{}

    // Year (optional. Default current year). The type is interface{} with range:
    // 1993..2035.
    Year interface{}
}

func (startTime *Ipsla_MplsLspMonitor_Schedules_Schedule_StartTime) GetEntityData() *types.CommonEntityData {
    startTime.EntityData.YFilter = startTime.YFilter
    startTime.EntityData.YangName = "start-time"
    startTime.EntityData.BundleName = "cisco_ios_xr"
    startTime.EntityData.ParentYangName = "schedule"
    startTime.EntityData.SegmentPath = "start-time"
    startTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startTime.EntityData.Children = make(map[string]types.YChild)
    startTime.EntityData.Leafs = make(map[string]types.YLeaf)
    startTime.EntityData.Leafs["schedule-type"] = types.YLeaf{"ScheduleType", startTime.ScheduleType}
    startTime.EntityData.Leafs["hour"] = types.YLeaf{"Hour", startTime.Hour}
    startTime.EntityData.Leafs["minute"] = types.YLeaf{"Minute", startTime.Minute}
    startTime.EntityData.Leafs["second"] = types.YLeaf{"Second", startTime.Second}
    startTime.EntityData.Leafs["month"] = types.YLeaf{"Month", startTime.Month}
    startTime.EntityData.Leafs["day"] = types.YLeaf{"Day", startTime.Day}
    startTime.EntityData.Leafs["year"] = types.YLeaf{"Year", startTime.Year}
    return &(startTime.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions
// MPLS LSP Monitor definition table
type Ipsla_MplsLspMonitor_Definitions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LSP Monitor definition. The type is slice of
    // Ipsla_MplsLspMonitor_Definitions_Definition.
    Definition []Ipsla_MplsLspMonitor_Definitions_Definition
}

func (definitions *Ipsla_MplsLspMonitor_Definitions) GetEntityData() *types.CommonEntityData {
    definitions.EntityData.YFilter = definitions.YFilter
    definitions.EntityData.YangName = "definitions"
    definitions.EntityData.BundleName = "cisco_ios_xr"
    definitions.EntityData.ParentYangName = "mpls-lsp-monitor"
    definitions.EntityData.SegmentPath = "definitions"
    definitions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    definitions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    definitions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    definitions.EntityData.Children = make(map[string]types.YChild)
    definitions.EntityData.Children["definition"] = types.YChild{"Definition", nil}
    for i := range definitions.Definition {
        definitions.EntityData.Children[types.GetSegmentPath(&definitions.Definition[i])] = types.YChild{"Definition", &definitions.Definition[i]}
    }
    definitions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(definitions.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition
// MPLS LSP Monitor definition
type Ipsla_MplsLspMonitor_Definitions_Definition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor identifier. The type is interface{} with
    // range: 1..2048.
    MonitorId interface{}

    // Operation type specification.
    OperationType Ipsla_MplsLspMonitor_Definitions_Definition_OperationType
}

func (definition *Ipsla_MplsLspMonitor_Definitions_Definition) GetEntityData() *types.CommonEntityData {
    definition.EntityData.YFilter = definition.YFilter
    definition.EntityData.YangName = "definition"
    definition.EntityData.BundleName = "cisco_ios_xr"
    definition.EntityData.ParentYangName = "definitions"
    definition.EntityData.SegmentPath = "definition" + "[monitor-id='" + fmt.Sprintf("%v", definition.MonitorId) + "']"
    definition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    definition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    definition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    definition.EntityData.Children = make(map[string]types.YChild)
    definition.EntityData.Children["operation-type"] = types.YChild{"OperationType", &definition.OperationType}
    definition.EntityData.Leafs = make(map[string]types.YLeaf)
    definition.EntityData.Leafs["monitor-id"] = types.YLeaf{"MonitorId", definition.MonitorId}
    return &(definition.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType
// Operation type specification
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Perform MPLS LSP Trace operation.
    MplsLspTrace Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace

    // Perform MPLS LSP Ping operation.
    MplsLspPing Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing
}

func (operationType *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType) GetEntityData() *types.CommonEntityData {
    operationType.EntityData.YFilter = operationType.YFilter
    operationType.EntityData.YangName = "operation-type"
    operationType.EntityData.BundleName = "cisco_ios_xr"
    operationType.EntityData.ParentYangName = "definition"
    operationType.EntityData.SegmentPath = "operation-type"
    operationType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationType.EntityData.Children = make(map[string]types.YChild)
    operationType.EntityData.Children["mpls-lsp-trace"] = types.YChild{"MplsLspTrace", &operationType.MplsLspTrace}
    operationType.EntityData.Children["mpls-lsp-ping"] = types.YChild{"MplsLspPing", &operationType.MplsLspPing}
    operationType.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(operationType.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace
// Perform MPLS LSP Trace operation
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time to live value. The type is interface{} with range: 1..255. The default
    // value is 30.
    Ttl interface{}

    // EXP bits in MPLS LSP echo request header. The type is interface{} with
    // range: 0..7. The default value is 0.
    ExpBits interface{}

    // Add a tag for this MPLSLM instance. The type is string with length: 1..100.
    Tag interface{}

    // Attributes used for path selection during LSP load balancing. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // The default value is 1.0.0.127.
    LspSelector interface{}

    // Echo request output interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Apply access list to filter PE addresses. The type is string with length:
    // 1..32.
    Accesslist interface{}

    // Create MPLSLM instance with specified type. The type is interface{}.
    Create interface{}

    // Echo request output nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    OutputNexthop interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Forced option for the MPLS LSP operation. The type is interface{}.
    ForceExplicitNull interface{}

    // Specify a VRF instance to be monitored. The type is string with length:
    // 1..32.
    Vrf interface{}

    // Echo reply options for the MPLS LSP operation.
    Reply Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Reply

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Statistics

    // Scanning parameters configuration.
    Scan Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Scan
}

func (mplsLspTrace *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace) GetEntityData() *types.CommonEntityData {
    mplsLspTrace.EntityData.YFilter = mplsLspTrace.YFilter
    mplsLspTrace.EntityData.YangName = "mpls-lsp-trace"
    mplsLspTrace.EntityData.BundleName = "cisco_ios_xr"
    mplsLspTrace.EntityData.ParentYangName = "operation-type"
    mplsLspTrace.EntityData.SegmentPath = "mpls-lsp-trace"
    mplsLspTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLspTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLspTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLspTrace.EntityData.Children = make(map[string]types.YChild)
    mplsLspTrace.EntityData.Children["reply"] = types.YChild{"Reply", &mplsLspTrace.Reply}
    mplsLspTrace.EntityData.Children["statistics"] = types.YChild{"Statistics", &mplsLspTrace.Statistics}
    mplsLspTrace.EntityData.Children["scan"] = types.YChild{"Scan", &mplsLspTrace.Scan}
    mplsLspTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLspTrace.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", mplsLspTrace.Ttl}
    mplsLspTrace.EntityData.Leafs["exp-bits"] = types.YLeaf{"ExpBits", mplsLspTrace.ExpBits}
    mplsLspTrace.EntityData.Leafs["tag"] = types.YLeaf{"Tag", mplsLspTrace.Tag}
    mplsLspTrace.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", mplsLspTrace.LspSelector}
    mplsLspTrace.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", mplsLspTrace.OutputInterface}
    mplsLspTrace.EntityData.Leafs["accesslist"] = types.YLeaf{"Accesslist", mplsLspTrace.Accesslist}
    mplsLspTrace.EntityData.Leafs["create"] = types.YLeaf{"Create", mplsLspTrace.Create}
    mplsLspTrace.EntityData.Leafs["output-nexthop"] = types.YLeaf{"OutputNexthop", mplsLspTrace.OutputNexthop}
    mplsLspTrace.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", mplsLspTrace.Timeout}
    mplsLspTrace.EntityData.Leafs["force-explicit-null"] = types.YLeaf{"ForceExplicitNull", mplsLspTrace.ForceExplicitNull}
    mplsLspTrace.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", mplsLspTrace.Vrf}
    return &(mplsLspTrace.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Reply
// Echo reply options for the MPLS LSP operation
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DSCP bits in the reply IP header. The type is one of the following types:
    // enumeration IpslaLspReplyDscp, or int with range: 0..63.
    DscpBits interface{}

    // Enables use of router alert in echo reply packets. The type is
    // IpslaLspMonitorReplyMode.
    Mode interface{}
}

func (reply *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "mpls-lsp-trace"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = make(map[string]types.YChild)
    reply.EntityData.Leafs = make(map[string]types.YLeaf)
    reply.EntityData.Leafs["dscp-bits"] = types.YLeaf{"DscpBits", reply.DscpBits}
    reply.EntityData.Leafs["mode"] = types.YLeaf{"Mode", reply.Mode}
    return &(reply.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Statistics
// Statistics collection aggregated over an hour
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..2. Units are hour. The default value is 2.
    Hours interface{}
}

func (statistics *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "mpls-lsp-trace"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    return &(statistics.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Scan
// Scanning parameters configuration
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Scan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time interval for automatic discovery. The type is interface{} with range:
    // 1..70560. Units are minute. The default value is 240.
    Interval interface{}

    // Number of times for automatic deletion. The type is interface{} with range:
    // 0..2147483647. The default value is 1.
    DeleteFactor interface{}
}

func (scan *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspTrace_Scan) GetEntityData() *types.CommonEntityData {
    scan.EntityData.YFilter = scan.YFilter
    scan.EntityData.YangName = "scan"
    scan.EntityData.BundleName = "cisco_ios_xr"
    scan.EntityData.ParentYangName = "mpls-lsp-trace"
    scan.EntityData.SegmentPath = "scan"
    scan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scan.EntityData.Children = make(map[string]types.YChild)
    scan.EntityData.Leafs = make(map[string]types.YLeaf)
    scan.EntityData.Leafs["interval"] = types.YLeaf{"Interval", scan.Interval}
    scan.EntityData.Leafs["delete-factor"] = types.YLeaf{"DeleteFactor", scan.DeleteFactor}
    return &(scan.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing
// Perform MPLS LSP Ping operation
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time to live value. The type is interface{} with range: 1..255. The default
    // value is 255.
    Ttl interface{}

    // EXP bits in MPLS LSP echo request header. The type is interface{} with
    // range: 0..7. The default value is 0.
    ExpBits interface{}

    // Add a tag for this MPLSLM instance. The type is string with length: 1..100.
    Tag interface{}

    // Attributes used for path selection during LSP load balancing. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // The default value is 1.0.0.127.
    LspSelector interface{}

    // Echo request output interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Apply access list to filter PE addresses. The type is string with length:
    // 1..32.
    Accesslist interface{}

    // Create MPLSLM instance with specified type. The type is interface{}.
    Create interface{}

    // Echo request output nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    OutputNexthop interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Forced option for the MPLS LSP operation. The type is interface{}.
    ForceExplicitNull interface{}

    // Specify a VRF instance to be monitored. The type is string with length:
    // 1..32.
    Vrf interface{}

    // Protocol data size in payload of probe packets.
    DataSize Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_DataSize

    // Path discover configuration.
    PathDiscover Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover

    // Echo reply options for the MPLS LSP operation.
    Reply Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Reply

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Statistics

    // Scanning parameters configuration.
    Scan Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Scan
}

func (mplsLspPing *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing) GetEntityData() *types.CommonEntityData {
    mplsLspPing.EntityData.YFilter = mplsLspPing.YFilter
    mplsLspPing.EntityData.YangName = "mpls-lsp-ping"
    mplsLspPing.EntityData.BundleName = "cisco_ios_xr"
    mplsLspPing.EntityData.ParentYangName = "operation-type"
    mplsLspPing.EntityData.SegmentPath = "mpls-lsp-ping"
    mplsLspPing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLspPing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLspPing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLspPing.EntityData.Children = make(map[string]types.YChild)
    mplsLspPing.EntityData.Children["data-size"] = types.YChild{"DataSize", &mplsLspPing.DataSize}
    mplsLspPing.EntityData.Children["path-discover"] = types.YChild{"PathDiscover", &mplsLspPing.PathDiscover}
    mplsLspPing.EntityData.Children["reply"] = types.YChild{"Reply", &mplsLspPing.Reply}
    mplsLspPing.EntityData.Children["statistics"] = types.YChild{"Statistics", &mplsLspPing.Statistics}
    mplsLspPing.EntityData.Children["scan"] = types.YChild{"Scan", &mplsLspPing.Scan}
    mplsLspPing.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLspPing.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", mplsLspPing.Ttl}
    mplsLspPing.EntityData.Leafs["exp-bits"] = types.YLeaf{"ExpBits", mplsLspPing.ExpBits}
    mplsLspPing.EntityData.Leafs["tag"] = types.YLeaf{"Tag", mplsLspPing.Tag}
    mplsLspPing.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", mplsLspPing.LspSelector}
    mplsLspPing.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", mplsLspPing.OutputInterface}
    mplsLspPing.EntityData.Leafs["accesslist"] = types.YLeaf{"Accesslist", mplsLspPing.Accesslist}
    mplsLspPing.EntityData.Leafs["create"] = types.YLeaf{"Create", mplsLspPing.Create}
    mplsLspPing.EntityData.Leafs["output-nexthop"] = types.YLeaf{"OutputNexthop", mplsLspPing.OutputNexthop}
    mplsLspPing.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", mplsLspPing.Timeout}
    mplsLspPing.EntityData.Leafs["force-explicit-null"] = types.YLeaf{"ForceExplicitNull", mplsLspPing.ForceExplicitNull}
    mplsLspPing.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", mplsLspPing.Vrf}
    return &(mplsLspPing.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_DataSize
// Protocol data size in payload of probe
// packets
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_DataSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload size in request probe packet. The type is interface{} with range:
    // 100..17986. Units are byte. The default value is 100.
    Request interface{}
}

func (dataSize *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_DataSize) GetEntityData() *types.CommonEntityData {
    dataSize.EntityData.YFilter = dataSize.YFilter
    dataSize.EntityData.YangName = "data-size"
    dataSize.EntityData.BundleName = "cisco_ios_xr"
    dataSize.EntityData.ParentYangName = "mpls-lsp-ping"
    dataSize.EntityData.SegmentPath = "data-size"
    dataSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataSize.EntityData.Children = make(map[string]types.YChild)
    dataSize.EntityData.Leafs = make(map[string]types.YLeaf)
    dataSize.EntityData.Leafs["request"] = types.YLeaf{"Request", dataSize.Request}
    return &(dataSize.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover
// Path discover configuration
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time period for finishing path discovery. The type is interface{} with
    // range: 0..7200. Units are minute. The default value is 0.
    ScanPeriod interface{}

    // Create LPD instance. The type is interface{}.
    Create interface{}

    // Session parameters configuration.
    Session Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Session

    // Path parameters configuration.
    Path Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path

    // Echo parameters configuration.
    Echo Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo
}

func (pathDiscover *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover) GetEntityData() *types.CommonEntityData {
    pathDiscover.EntityData.YFilter = pathDiscover.YFilter
    pathDiscover.EntityData.YangName = "path-discover"
    pathDiscover.EntityData.BundleName = "cisco_ios_xr"
    pathDiscover.EntityData.ParentYangName = "mpls-lsp-ping"
    pathDiscover.EntityData.SegmentPath = "path-discover"
    pathDiscover.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathDiscover.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathDiscover.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathDiscover.EntityData.Children = make(map[string]types.YChild)
    pathDiscover.EntityData.Children["session"] = types.YChild{"Session", &pathDiscover.Session}
    pathDiscover.EntityData.Children["path"] = types.YChild{"Path", &pathDiscover.Path}
    pathDiscover.EntityData.Children["echo"] = types.YChild{"Echo", &pathDiscover.Echo}
    pathDiscover.EntityData.Leafs = make(map[string]types.YLeaf)
    pathDiscover.EntityData.Leafs["scan-period"] = types.YLeaf{"ScanPeriod", pathDiscover.ScanPeriod}
    pathDiscover.EntityData.Leafs["create"] = types.YLeaf{"Create", pathDiscover.Create}
    return &(pathDiscover.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Session
// Session parameters configuration
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeout value for path discovery request. The type is interface{} with
    // range: 1..900. Units are second. The default value is 120.
    Timeout interface{}

    // Number of concurrent active path discovery requests at one time. The type
    // is interface{} with range: 1..15. The default value is 1.
    Limit interface{}
}

func (session *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "path-discover"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", session.Timeout}
    session.EntityData.Leafs["limit"] = types.YLeaf{"Limit", session.Limit}
    return &(session.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path
// Path parameters configuration
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of attempts before declaring the path as down. The type is
    // interface{} with range: 1..16. The default value is 1.
    Retry interface{}

    // Frequency to be used if path failure condition is detected.
    SecondaryFrequency Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path_SecondaryFrequency
}

func (path *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "path-discover"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["secondary-frequency"] = types.YChild{"SecondaryFrequency", &path.SecondaryFrequency}
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["retry"] = types.YLeaf{"Retry", path.Retry}
    return &(path.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path_SecondaryFrequency
// Frequency to be used if path failure
// condition is detected
// This type is a presence type.
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path_SecondaryFrequency struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Condition type of path failure. The type is IpslaSecondaryFrequency. This
    // attribute is mandatory.
    Type_ interface{}

    // Frequency value in seconds. The type is interface{} with range: 1..604800.
    // This attribute is mandatory. Units are second.
    Frequency interface{}
}

func (secondaryFrequency *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Path_SecondaryFrequency) GetEntityData() *types.CommonEntityData {
    secondaryFrequency.EntityData.YFilter = secondaryFrequency.YFilter
    secondaryFrequency.EntityData.YangName = "secondary-frequency"
    secondaryFrequency.EntityData.BundleName = "cisco_ios_xr"
    secondaryFrequency.EntityData.ParentYangName = "path"
    secondaryFrequency.EntityData.SegmentPath = "secondary-frequency"
    secondaryFrequency.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryFrequency.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryFrequency.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryFrequency.EntityData.Children = make(map[string]types.YChild)
    secondaryFrequency.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryFrequency.EntityData.Leafs["type"] = types.YLeaf{"Type_", secondaryFrequency.Type_}
    secondaryFrequency.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", secondaryFrequency.Frequency}
    return &(secondaryFrequency.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo
// Echo parameters configuration
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Send interval between echo requests during path discovery. The type is
    // interface{} with range: 0..3600000. Units are millisecond. The default
    // value is 0.
    Interval interface{}

    // Timeout value for echo requests during path discovery. The type is
    // interface{} with range: 1..3600. Units are second. The default value is 5.
    Timeout interface{}

    // Number of timeout retry attempts during path discovery. The type is
    // interface{} with range: 0..10. The default value is 3.
    Retry interface{}

    // Maximum IPv4 address used as destination in echo request. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // The default value is 127.255.255.255.
    MaximumLspSelector interface{}

    // Downstream map multipath settings.
    Multipath Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo_Multipath
}

func (echo *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo) GetEntityData() *types.CommonEntityData {
    echo.EntityData.YFilter = echo.YFilter
    echo.EntityData.YangName = "echo"
    echo.EntityData.BundleName = "cisco_ios_xr"
    echo.EntityData.ParentYangName = "path-discover"
    echo.EntityData.SegmentPath = "echo"
    echo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echo.EntityData.Children = make(map[string]types.YChild)
    echo.EntityData.Children["multipath"] = types.YChild{"Multipath", &echo.Multipath}
    echo.EntityData.Leafs = make(map[string]types.YLeaf)
    echo.EntityData.Leafs["interval"] = types.YLeaf{"Interval", echo.Interval}
    echo.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", echo.Timeout}
    echo.EntityData.Leafs["retry"] = types.YLeaf{"Retry", echo.Retry}
    echo.EntityData.Leafs["maximum-lsp-selector"] = types.YLeaf{"MaximumLspSelector", echo.MaximumLspSelector}
    return &(echo.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo_Multipath
// Downstream map multipath settings
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo_Multipath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multipath bit size. The type is interface{} with range: 1..256. The default
    // value is 32.
    BitmapSize interface{}
}

func (multipath *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_PathDiscover_Echo_Multipath) GetEntityData() *types.CommonEntityData {
    multipath.EntityData.YFilter = multipath.YFilter
    multipath.EntityData.YangName = "multipath"
    multipath.EntityData.BundleName = "cisco_ios_xr"
    multipath.EntityData.ParentYangName = "echo"
    multipath.EntityData.SegmentPath = "multipath"
    multipath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multipath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multipath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multipath.EntityData.Children = make(map[string]types.YChild)
    multipath.EntityData.Leafs = make(map[string]types.YLeaf)
    multipath.EntityData.Leafs["bitmap-size"] = types.YLeaf{"BitmapSize", multipath.BitmapSize}
    return &(multipath.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Reply
// Echo reply options for the MPLS LSP operation
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DSCP bits in the reply IP header. The type is one of the following types:
    // enumeration IpslaLspReplyDscp, or int with range: 0..63.
    DscpBits interface{}

    // Enables use of router alert in echo reply packets. The type is
    // IpslaLspMonitorReplyMode.
    Mode interface{}
}

func (reply *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "mpls-lsp-ping"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = make(map[string]types.YChild)
    reply.EntityData.Leafs = make(map[string]types.YLeaf)
    reply.EntityData.Leafs["dscp-bits"] = types.YLeaf{"DscpBits", reply.DscpBits}
    reply.EntityData.Leafs["mode"] = types.YLeaf{"Mode", reply.Mode}
    return &(reply.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Statistics
// Statistics collection aggregated over an hour
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..2. Units are hour. The default value is 2.
    Hours interface{}
}

func (statistics *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "mpls-lsp-ping"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    return &(statistics.EntityData)
}

// Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Scan
// Scanning parameters configuration
type Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Scan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time interval for automatic discovery. The type is interface{} with range:
    // 1..70560. Units are minute. The default value is 240.
    Interval interface{}

    // Number of times for automatic deletion. The type is interface{} with range:
    // 0..2147483647. The default value is 1.
    DeleteFactor interface{}
}

func (scan *Ipsla_MplsLspMonitor_Definitions_Definition_OperationType_MplsLspPing_Scan) GetEntityData() *types.CommonEntityData {
    scan.EntityData.YFilter = scan.YFilter
    scan.EntityData.YangName = "scan"
    scan.EntityData.BundleName = "cisco_ios_xr"
    scan.EntityData.ParentYangName = "mpls-lsp-ping"
    scan.EntityData.SegmentPath = "scan"
    scan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scan.EntityData.Children = make(map[string]types.YChild)
    scan.EntityData.Leafs = make(map[string]types.YLeaf)
    scan.EntityData.Leafs["interval"] = types.YLeaf{"Interval", scan.Interval}
    scan.EntityData.Leafs["delete-factor"] = types.YLeaf{"DeleteFactor", scan.DeleteFactor}
    return &(scan.EntityData)
}

// Ipsla_Operation
// IPSLA Operation configuration
type Ipsla_Operation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Schedule an operation.
    Schedules Ipsla_Operation_Schedules

    // Reaction configuration.
    Reactions Ipsla_Operation_Reactions

    // Reaction trigger configuration.
    ReactionTriggers Ipsla_Operation_ReactionTriggers

    // Operation definition table.
    Definitions Ipsla_Operation_Definitions
}

func (operation *Ipsla_Operation) GetEntityData() *types.CommonEntityData {
    operation.EntityData.YFilter = operation.YFilter
    operation.EntityData.YangName = "operation"
    operation.EntityData.BundleName = "cisco_ios_xr"
    operation.EntityData.ParentYangName = "ipsla"
    operation.EntityData.SegmentPath = "operation"
    operation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operation.EntityData.Children = make(map[string]types.YChild)
    operation.EntityData.Children["schedules"] = types.YChild{"Schedules", &operation.Schedules}
    operation.EntityData.Children["reactions"] = types.YChild{"Reactions", &operation.Reactions}
    operation.EntityData.Children["reaction-triggers"] = types.YChild{"ReactionTriggers", &operation.ReactionTriggers}
    operation.EntityData.Children["definitions"] = types.YChild{"Definitions", &operation.Definitions}
    operation.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(operation.EntityData)
}

// Ipsla_Operation_Schedules
// Schedule an operation
type Ipsla_Operation_Schedules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation schedule configuration. The type is slice of
    // Ipsla_Operation_Schedules_Schedule.
    Schedule []Ipsla_Operation_Schedules_Schedule
}

func (schedules *Ipsla_Operation_Schedules) GetEntityData() *types.CommonEntityData {
    schedules.EntityData.YFilter = schedules.YFilter
    schedules.EntityData.YangName = "schedules"
    schedules.EntityData.BundleName = "cisco_ios_xr"
    schedules.EntityData.ParentYangName = "operation"
    schedules.EntityData.SegmentPath = "schedules"
    schedules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schedules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schedules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schedules.EntityData.Children = make(map[string]types.YChild)
    schedules.EntityData.Children["schedule"] = types.YChild{"Schedule", nil}
    for i := range schedules.Schedule {
        schedules.EntityData.Children[types.GetSegmentPath(&schedules.Schedule[i])] = types.YChild{"Schedule", &schedules.Schedule[i]}
    }
    schedules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(schedules.EntityData)
}

// Ipsla_Operation_Schedules_Schedule
// Operation schedule configuration
type Ipsla_Operation_Schedules_Schedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operation number. The type is interface{} with
    // range: 1..2048.
    OperationId interface{}

    // Length of the time to execute (default 3600 seconds). The type is one of
    // the following types: enumeration IpslaLife Units are second., or int with
    // range: 0..2147483647 Units are second..
    Life interface{}

    // How long to keep this entry after it becomes inactive. The type is
    // interface{} with range: 0..2073600. Units are second.
    Ageout interface{}

    // probe to be scheduled automatically every day. The type is interface{}.
    Recurring interface{}

    // Start time of the operation.
    StartTime Ipsla_Operation_Schedules_Schedule_StartTime
}

func (schedule *Ipsla_Operation_Schedules_Schedule) GetEntityData() *types.CommonEntityData {
    schedule.EntityData.YFilter = schedule.YFilter
    schedule.EntityData.YangName = "schedule"
    schedule.EntityData.BundleName = "cisco_ios_xr"
    schedule.EntityData.ParentYangName = "schedules"
    schedule.EntityData.SegmentPath = "schedule" + "[operation-id='" + fmt.Sprintf("%v", schedule.OperationId) + "']"
    schedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schedule.EntityData.Children = make(map[string]types.YChild)
    schedule.EntityData.Children["start-time"] = types.YChild{"StartTime", &schedule.StartTime}
    schedule.EntityData.Leafs = make(map[string]types.YLeaf)
    schedule.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", schedule.OperationId}
    schedule.EntityData.Leafs["life"] = types.YLeaf{"Life", schedule.Life}
    schedule.EntityData.Leafs["ageout"] = types.YLeaf{"Ageout", schedule.Ageout}
    schedule.EntityData.Leafs["recurring"] = types.YLeaf{"Recurring", schedule.Recurring}
    return &(schedule.EntityData)
}

// Ipsla_Operation_Schedules_Schedule_StartTime
// Start time of the operation
type Ipsla_Operation_Schedules_Schedule_StartTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of schedule. The type is IpslaSched.
    ScheduleType interface{}

    // Hour value(hh) in hh:mm:ss specification. The type is interface{} with
    // range: 0..23.
    Hour interface{}

    // Minute value(mm) in hh:mm:ss specification. The type is interface{} with
    // range: 0..59.
    Minute interface{}

    // Second value(ss) in hh:mm:ss specification. The type is interface{} with
    // range: 0..59.
    Second interface{}

    // Month of the year (optional. Default current month). The type is
    // IpslaMonth.
    Month interface{}

    // Day of the month(optional. Default today). The type is interface{} with
    // range: 1..31.
    Day interface{}

    // Year(optional. Default current year). The type is interface{} with range:
    // 1993..2035.
    Year interface{}
}

func (startTime *Ipsla_Operation_Schedules_Schedule_StartTime) GetEntityData() *types.CommonEntityData {
    startTime.EntityData.YFilter = startTime.YFilter
    startTime.EntityData.YangName = "start-time"
    startTime.EntityData.BundleName = "cisco_ios_xr"
    startTime.EntityData.ParentYangName = "schedule"
    startTime.EntityData.SegmentPath = "start-time"
    startTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startTime.EntityData.Children = make(map[string]types.YChild)
    startTime.EntityData.Leafs = make(map[string]types.YLeaf)
    startTime.EntityData.Leafs["schedule-type"] = types.YLeaf{"ScheduleType", startTime.ScheduleType}
    startTime.EntityData.Leafs["hour"] = types.YLeaf{"Hour", startTime.Hour}
    startTime.EntityData.Leafs["minute"] = types.YLeaf{"Minute", startTime.Minute}
    startTime.EntityData.Leafs["second"] = types.YLeaf{"Second", startTime.Second}
    startTime.EntityData.Leafs["month"] = types.YLeaf{"Month", startTime.Month}
    startTime.EntityData.Leafs["day"] = types.YLeaf{"Day", startTime.Day}
    startTime.EntityData.Leafs["year"] = types.YLeaf{"Year", startTime.Year}
    return &(startTime.EntityData)
}

// Ipsla_Operation_Reactions
// Reaction configuration
type Ipsla_Operation_Reactions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reaction configuration for an operation. The type is slice of
    // Ipsla_Operation_Reactions_Reaction.
    Reaction []Ipsla_Operation_Reactions_Reaction
}

func (reactions *Ipsla_Operation_Reactions) GetEntityData() *types.CommonEntityData {
    reactions.EntityData.YFilter = reactions.YFilter
    reactions.EntityData.YangName = "reactions"
    reactions.EntityData.BundleName = "cisco_ios_xr"
    reactions.EntityData.ParentYangName = "operation"
    reactions.EntityData.SegmentPath = "reactions"
    reactions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reactions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reactions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reactions.EntityData.Children = make(map[string]types.YChild)
    reactions.EntityData.Children["reaction"] = types.YChild{"Reaction", nil}
    for i := range reactions.Reaction {
        reactions.EntityData.Children[types.GetSegmentPath(&reactions.Reaction[i])] = types.YChild{"Reaction", &reactions.Reaction[i]}
    }
    reactions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reactions.EntityData)
}

// Ipsla_Operation_Reactions_Reaction
// Reaction configuration for an operation
type Ipsla_Operation_Reactions_Reaction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operation number. The type is interface{} with
    // range: 1..2048.
    OperationId interface{}

    // Reaction condition specification.
    Condition Ipsla_Operation_Reactions_Reaction_Condition
}

func (reaction *Ipsla_Operation_Reactions_Reaction) GetEntityData() *types.CommonEntityData {
    reaction.EntityData.YFilter = reaction.YFilter
    reaction.EntityData.YangName = "reaction"
    reaction.EntityData.BundleName = "cisco_ios_xr"
    reaction.EntityData.ParentYangName = "reactions"
    reaction.EntityData.SegmentPath = "reaction" + "[operation-id='" + fmt.Sprintf("%v", reaction.OperationId) + "']"
    reaction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reaction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reaction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reaction.EntityData.Children = make(map[string]types.YChild)
    reaction.EntityData.Children["condition"] = types.YChild{"Condition", &reaction.Condition}
    reaction.EntityData.Leafs = make(map[string]types.YLeaf)
    reaction.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", reaction.OperationId}
    return &(reaction.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition
// Reaction condition specification
type Ipsla_Operation_Reactions_Reaction_Condition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // React on destination to source jitter threshold violation.
    JitterAverageDs Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs

    // React on probe timeout.
    Timeout Ipsla_Operation_Reactions_Reaction_Condition_Timeout

    // React on average round trip jitter threshold violation.
    JitterAverage Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage

    // React on error verfication violation.
    VerifyError Ipsla_Operation_Reactions_Reaction_Condition_VerifyError

    // React on round trip time threshold violation.
    Rtt Ipsla_Operation_Reactions_Reaction_Condition_Rtt

    // React on destination to source packet loss threshold violation.
    PacketLossSd Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd

    // React on average source to destination jitter threshold violation.
    JitterAverageSd Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd

    // React on connection loss for a monitored operation.
    ConnectionLoss Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss

    // React on source to destination packet loss threshold violation.
    PacketLossDs Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs
}

func (condition *Ipsla_Operation_Reactions_Reaction_Condition) GetEntityData() *types.CommonEntityData {
    condition.EntityData.YFilter = condition.YFilter
    condition.EntityData.YangName = "condition"
    condition.EntityData.BundleName = "cisco_ios_xr"
    condition.EntityData.ParentYangName = "reaction"
    condition.EntityData.SegmentPath = "condition"
    condition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    condition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    condition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    condition.EntityData.Children = make(map[string]types.YChild)
    condition.EntityData.Children["jitter-average-ds"] = types.YChild{"JitterAverageDs", &condition.JitterAverageDs}
    condition.EntityData.Children["timeout"] = types.YChild{"Timeout", &condition.Timeout}
    condition.EntityData.Children["jitter-average"] = types.YChild{"JitterAverage", &condition.JitterAverage}
    condition.EntityData.Children["verify-error"] = types.YChild{"VerifyError", &condition.VerifyError}
    condition.EntityData.Children["rtt"] = types.YChild{"Rtt", &condition.Rtt}
    condition.EntityData.Children["packet-loss-sd"] = types.YChild{"PacketLossSd", &condition.PacketLossSd}
    condition.EntityData.Children["jitter-average-sd"] = types.YChild{"JitterAverageSd", &condition.JitterAverageSd}
    condition.EntityData.Children["connection-loss"] = types.YChild{"ConnectionLoss", &condition.ConnectionLoss}
    condition.EntityData.Children["packet-loss-ds"] = types.YChild{"PacketLossDs", &condition.PacketLossDs}
    condition.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(condition.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs
// React on destination to source jitter
// threshold violation
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Specify threshold limits for the monitored element.
    ThresholdLimits Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdLimits

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdType
}

func (jitterAverageDs *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs) GetEntityData() *types.CommonEntityData {
    jitterAverageDs.EntityData.YFilter = jitterAverageDs.YFilter
    jitterAverageDs.EntityData.YangName = "jitter-average-ds"
    jitterAverageDs.EntityData.BundleName = "cisco_ios_xr"
    jitterAverageDs.EntityData.ParentYangName = "condition"
    jitterAverageDs.EntityData.SegmentPath = "jitter-average-ds"
    jitterAverageDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    jitterAverageDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    jitterAverageDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    jitterAverageDs.EntityData.Children = make(map[string]types.YChild)
    jitterAverageDs.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &jitterAverageDs.ThresholdLimits}
    jitterAverageDs.EntityData.Children["action-type"] = types.YChild{"ActionType", &jitterAverageDs.ActionType}
    jitterAverageDs.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &jitterAverageDs.ThresholdType}
    jitterAverageDs.EntityData.Leafs = make(map[string]types.YLeaf)
    jitterAverageDs.EntityData.Leafs["create"] = types.YLeaf{"Create", jitterAverageDs.Create}
    return &(jitterAverageDs.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdLimits
// Specify threshold limits for the monitored
// element
// This type is a presence type.
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold lower limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    LowerLimit interface{}

    // Threshold upper limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    UpperLimit interface{}
}

func (thresholdLimits *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "jitter-average-ds"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdLimits.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", thresholdLimits.LowerLimit}
    thresholdLimits.EntityData.Leafs["upper-limit"] = types.YLeaf{"UpperLimit", thresholdLimits.UpperLimit}
    return &(thresholdLimits.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "jitter-average-ds"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageDs_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "jitter-average-ds"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_Timeout
// React on probe timeout
type Ipsla_Operation_Reactions_Reaction_Condition_Timeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ThresholdType
}

func (timeout *Ipsla_Operation_Reactions_Reaction_Condition_Timeout) GetEntityData() *types.CommonEntityData {
    timeout.EntityData.YFilter = timeout.YFilter
    timeout.EntityData.YangName = "timeout"
    timeout.EntityData.BundleName = "cisco_ios_xr"
    timeout.EntityData.ParentYangName = "condition"
    timeout.EntityData.SegmentPath = "timeout"
    timeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeout.EntityData.Children = make(map[string]types.YChild)
    timeout.EntityData.Children["action-type"] = types.YChild{"ActionType", &timeout.ActionType}
    timeout.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &timeout.ThresholdType}
    timeout.EntityData.Leafs = make(map[string]types.YLeaf)
    timeout.EntityData.Leafs["create"] = types.YLeaf{"Create", timeout.Create}
    return &(timeout.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "timeout"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_Timeout_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "timeout"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage
// React on average round trip jitter threshold
// violation
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Specify threshold limits for the monitored element.
    ThresholdLimits Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdLimits

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdType
}

func (jitterAverage *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage) GetEntityData() *types.CommonEntityData {
    jitterAverage.EntityData.YFilter = jitterAverage.YFilter
    jitterAverage.EntityData.YangName = "jitter-average"
    jitterAverage.EntityData.BundleName = "cisco_ios_xr"
    jitterAverage.EntityData.ParentYangName = "condition"
    jitterAverage.EntityData.SegmentPath = "jitter-average"
    jitterAverage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    jitterAverage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    jitterAverage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    jitterAverage.EntityData.Children = make(map[string]types.YChild)
    jitterAverage.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &jitterAverage.ThresholdLimits}
    jitterAverage.EntityData.Children["action-type"] = types.YChild{"ActionType", &jitterAverage.ActionType}
    jitterAverage.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &jitterAverage.ThresholdType}
    jitterAverage.EntityData.Leafs = make(map[string]types.YLeaf)
    jitterAverage.EntityData.Leafs["create"] = types.YLeaf{"Create", jitterAverage.Create}
    return &(jitterAverage.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdLimits
// Specify threshold limits for the monitored
// element
// This type is a presence type.
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold lower limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    LowerLimit interface{}

    // Threshold upper limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    UpperLimit interface{}
}

func (thresholdLimits *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "jitter-average"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdLimits.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", thresholdLimits.LowerLimit}
    thresholdLimits.EntityData.Leafs["upper-limit"] = types.YLeaf{"UpperLimit", thresholdLimits.UpperLimit}
    return &(thresholdLimits.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "jitter-average"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverage_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "jitter-average"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_VerifyError
// React on error verfication violation
type Ipsla_Operation_Reactions_Reaction_Condition_VerifyError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ThresholdType
}

func (verifyError *Ipsla_Operation_Reactions_Reaction_Condition_VerifyError) GetEntityData() *types.CommonEntityData {
    verifyError.EntityData.YFilter = verifyError.YFilter
    verifyError.EntityData.YangName = "verify-error"
    verifyError.EntityData.BundleName = "cisco_ios_xr"
    verifyError.EntityData.ParentYangName = "condition"
    verifyError.EntityData.SegmentPath = "verify-error"
    verifyError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    verifyError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    verifyError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    verifyError.EntityData.Children = make(map[string]types.YChild)
    verifyError.EntityData.Children["action-type"] = types.YChild{"ActionType", &verifyError.ActionType}
    verifyError.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &verifyError.ThresholdType}
    verifyError.EntityData.Leafs = make(map[string]types.YLeaf)
    verifyError.EntityData.Leafs["create"] = types.YLeaf{"Create", verifyError.Create}
    return &(verifyError.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "verify-error"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_VerifyError_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "verify-error"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_Rtt
// React on round trip time threshold violation
type Ipsla_Operation_Reactions_Reaction_Condition_Rtt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Specify threshold limits for the monitored element.
    ThresholdLimits Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdLimits

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdType
}

func (rtt *Ipsla_Operation_Reactions_Reaction_Condition_Rtt) GetEntityData() *types.CommonEntityData {
    rtt.EntityData.YFilter = rtt.YFilter
    rtt.EntityData.YangName = "rtt"
    rtt.EntityData.BundleName = "cisco_ios_xr"
    rtt.EntityData.ParentYangName = "condition"
    rtt.EntityData.SegmentPath = "rtt"
    rtt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtt.EntityData.Children = make(map[string]types.YChild)
    rtt.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &rtt.ThresholdLimits}
    rtt.EntityData.Children["action-type"] = types.YChild{"ActionType", &rtt.ActionType}
    rtt.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &rtt.ThresholdType}
    rtt.EntityData.Leafs = make(map[string]types.YLeaf)
    rtt.EntityData.Leafs["create"] = types.YLeaf{"Create", rtt.Create}
    return &(rtt.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdLimits
// Specify threshold limits for the monitored
// element
// This type is a presence type.
type Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold lower limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    LowerLimit interface{}

    // Threshold upper limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    UpperLimit interface{}
}

func (thresholdLimits *Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "rtt"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdLimits.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", thresholdLimits.LowerLimit}
    thresholdLimits.EntityData.Leafs["upper-limit"] = types.YLeaf{"UpperLimit", thresholdLimits.UpperLimit}
    return &(thresholdLimits.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "rtt"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_Rtt_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "rtt"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd
// React on destination to source packet loss
// threshold violation
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Specify threshold limits for the monitored element.
    ThresholdLimits Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdLimits

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdType
}

func (packetLossSd *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd) GetEntityData() *types.CommonEntityData {
    packetLossSd.EntityData.YFilter = packetLossSd.YFilter
    packetLossSd.EntityData.YangName = "packet-loss-sd"
    packetLossSd.EntityData.BundleName = "cisco_ios_xr"
    packetLossSd.EntityData.ParentYangName = "condition"
    packetLossSd.EntityData.SegmentPath = "packet-loss-sd"
    packetLossSd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetLossSd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetLossSd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetLossSd.EntityData.Children = make(map[string]types.YChild)
    packetLossSd.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &packetLossSd.ThresholdLimits}
    packetLossSd.EntityData.Children["action-type"] = types.YChild{"ActionType", &packetLossSd.ActionType}
    packetLossSd.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &packetLossSd.ThresholdType}
    packetLossSd.EntityData.Leafs = make(map[string]types.YLeaf)
    packetLossSd.EntityData.Leafs["create"] = types.YLeaf{"Create", packetLossSd.Create}
    return &(packetLossSd.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdLimits
// Specify threshold limits for the monitored
// element
// This type is a presence type.
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold lower limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    LowerLimit interface{}

    // Threshold upper limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    UpperLimit interface{}
}

func (thresholdLimits *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "packet-loss-sd"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdLimits.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", thresholdLimits.LowerLimit}
    thresholdLimits.EntityData.Leafs["upper-limit"] = types.YLeaf{"UpperLimit", thresholdLimits.UpperLimit}
    return &(thresholdLimits.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "packet-loss-sd"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossSd_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "packet-loss-sd"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd
// React on average source to destination
// jitter threshold violation
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Specify threshold limits for the monitored element.
    ThresholdLimits Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdLimits

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdType
}

func (jitterAverageSd *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd) GetEntityData() *types.CommonEntityData {
    jitterAverageSd.EntityData.YFilter = jitterAverageSd.YFilter
    jitterAverageSd.EntityData.YangName = "jitter-average-sd"
    jitterAverageSd.EntityData.BundleName = "cisco_ios_xr"
    jitterAverageSd.EntityData.ParentYangName = "condition"
    jitterAverageSd.EntityData.SegmentPath = "jitter-average-sd"
    jitterAverageSd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    jitterAverageSd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    jitterAverageSd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    jitterAverageSd.EntityData.Children = make(map[string]types.YChild)
    jitterAverageSd.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &jitterAverageSd.ThresholdLimits}
    jitterAverageSd.EntityData.Children["action-type"] = types.YChild{"ActionType", &jitterAverageSd.ActionType}
    jitterAverageSd.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &jitterAverageSd.ThresholdType}
    jitterAverageSd.EntityData.Leafs = make(map[string]types.YLeaf)
    jitterAverageSd.EntityData.Leafs["create"] = types.YLeaf{"Create", jitterAverageSd.Create}
    return &(jitterAverageSd.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdLimits
// Specify threshold limits for the monitored
// element
// This type is a presence type.
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold lower limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    LowerLimit interface{}

    // Threshold upper limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    UpperLimit interface{}
}

func (thresholdLimits *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "jitter-average-sd"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdLimits.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", thresholdLimits.LowerLimit}
    thresholdLimits.EntityData.Leafs["upper-limit"] = types.YLeaf{"UpperLimit", thresholdLimits.UpperLimit}
    return &(thresholdLimits.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "jitter-average-sd"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_JitterAverageSd_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "jitter-average-sd"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss
// React on connection loss for a monitored
// operation
type Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType
}

func (connectionLoss *Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss) GetEntityData() *types.CommonEntityData {
    connectionLoss.EntityData.YFilter = connectionLoss.YFilter
    connectionLoss.EntityData.YangName = "connection-loss"
    connectionLoss.EntityData.BundleName = "cisco_ios_xr"
    connectionLoss.EntityData.ParentYangName = "condition"
    connectionLoss.EntityData.SegmentPath = "connection-loss"
    connectionLoss.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectionLoss.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectionLoss.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectionLoss.EntityData.Children = make(map[string]types.YChild)
    connectionLoss.EntityData.Children["action-type"] = types.YChild{"ActionType", &connectionLoss.ActionType}
    connectionLoss.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &connectionLoss.ThresholdType}
    connectionLoss.EntityData.Leafs = make(map[string]types.YLeaf)
    connectionLoss.EntityData.Leafs["create"] = types.YLeaf{"Create", connectionLoss.Create}
    return &(connectionLoss.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "connection-loss"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_ConnectionLoss_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "connection-loss"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs
// React on source to destination packet loss
// threshold violation
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create reaction condition for a particular operation. The type is
    // interface{}.
    Create interface{}

    // Specify threshold limits for the monitored element.
    ThresholdLimits Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdLimits

    // Type of action to be taken on threshold violation(s).
    ActionType Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ActionType

    // Type of thresholding to perform on the monitored element.
    ThresholdType Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdType
}

func (packetLossDs *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs) GetEntityData() *types.CommonEntityData {
    packetLossDs.EntityData.YFilter = packetLossDs.YFilter
    packetLossDs.EntityData.YangName = "packet-loss-ds"
    packetLossDs.EntityData.BundleName = "cisco_ios_xr"
    packetLossDs.EntityData.ParentYangName = "condition"
    packetLossDs.EntityData.SegmentPath = "packet-loss-ds"
    packetLossDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetLossDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetLossDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetLossDs.EntityData.Children = make(map[string]types.YChild)
    packetLossDs.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &packetLossDs.ThresholdLimits}
    packetLossDs.EntityData.Children["action-type"] = types.YChild{"ActionType", &packetLossDs.ActionType}
    packetLossDs.EntityData.Children["threshold-type"] = types.YChild{"ThresholdType", &packetLossDs.ThresholdType}
    packetLossDs.EntityData.Leafs = make(map[string]types.YLeaf)
    packetLossDs.EntityData.Leafs["create"] = types.YLeaf{"Create", packetLossDs.Create}
    return &(packetLossDs.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdLimits
// Specify threshold limits for the monitored
// element
// This type is a presence type.
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold lower limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    LowerLimit interface{}

    // Threshold upper limit value. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    UpperLimit interface{}
}

func (thresholdLimits *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "packet-loss-ds"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdLimits.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", thresholdLimits.LowerLimit}
    thresholdLimits.EntityData.Leafs["upper-limit"] = types.YLeaf{"UpperLimit", thresholdLimits.UpperLimit}
    return &(thresholdLimits.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ActionType
// Type of action to be taken on threshold
// violation(s)
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ActionType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generate a syslog alarm on threshold violation. The type is interface{}.
    Logging interface{}

    // Generate trigger to active reaction triggered operation(s). The type is
    // interface{}.
    Trigger interface{}
}

func (actionType *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ActionType) GetEntityData() *types.CommonEntityData {
    actionType.EntityData.YFilter = actionType.YFilter
    actionType.EntityData.YangName = "action-type"
    actionType.EntityData.BundleName = "cisco_ios_xr"
    actionType.EntityData.ParentYangName = "packet-loss-ds"
    actionType.EntityData.SegmentPath = "action-type"
    actionType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionType.EntityData.Children = make(map[string]types.YChild)
    actionType.EntityData.Leafs = make(map[string]types.YLeaf)
    actionType.EntityData.Leafs["logging"] = types.YLeaf{"Logging", actionType.Logging}
    actionType.EntityData.Leafs["trigger"] = types.YLeaf{"Trigger", actionType.Trigger}
    return &(actionType.EntityData)
}

// Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdType
// Type of thresholding to perform on the monitored
// element
type Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of thresholding to perform. The type is IpslaThresholdTypes.
    ThreshType interface{}

    // Probe count for avarage, consecutive case or X value for XofY case. The
    // type is interface{} with range: 1..16.
    Count1 interface{}

    // Y value, when threshold type is XofY. The type is interface{} with range:
    // 1..16.
    Count2 interface{}
}

func (thresholdType *Ipsla_Operation_Reactions_Reaction_Condition_PacketLossDs_ThresholdType) GetEntityData() *types.CommonEntityData {
    thresholdType.EntityData.YFilter = thresholdType.YFilter
    thresholdType.EntityData.YangName = "threshold-type"
    thresholdType.EntityData.BundleName = "cisco_ios_xr"
    thresholdType.EntityData.ParentYangName = "packet-loss-ds"
    thresholdType.EntityData.SegmentPath = "threshold-type"
    thresholdType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdType.EntityData.Children = make(map[string]types.YChild)
    thresholdType.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdType.EntityData.Leafs["thresh-type"] = types.YLeaf{"ThreshType", thresholdType.ThreshType}
    thresholdType.EntityData.Leafs["count1"] = types.YLeaf{"Count1", thresholdType.Count1}
    thresholdType.EntityData.Leafs["count2"] = types.YLeaf{"Count2", thresholdType.Count2}
    return &(thresholdType.EntityData)
}

// Ipsla_Operation_ReactionTriggers
// Reaction trigger configuration
type Ipsla_Operation_ReactionTriggers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reaction trigger for an operation. The type is slice of
    // Ipsla_Operation_ReactionTriggers_ReactionTrigger.
    ReactionTrigger []Ipsla_Operation_ReactionTriggers_ReactionTrigger
}

func (reactionTriggers *Ipsla_Operation_ReactionTriggers) GetEntityData() *types.CommonEntityData {
    reactionTriggers.EntityData.YFilter = reactionTriggers.YFilter
    reactionTriggers.EntityData.YangName = "reaction-triggers"
    reactionTriggers.EntityData.BundleName = "cisco_ios_xr"
    reactionTriggers.EntityData.ParentYangName = "operation"
    reactionTriggers.EntityData.SegmentPath = "reaction-triggers"
    reactionTriggers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reactionTriggers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reactionTriggers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reactionTriggers.EntityData.Children = make(map[string]types.YChild)
    reactionTriggers.EntityData.Children["reaction-trigger"] = types.YChild{"ReactionTrigger", nil}
    for i := range reactionTriggers.ReactionTrigger {
        reactionTriggers.EntityData.Children[types.GetSegmentPath(&reactionTriggers.ReactionTrigger[i])] = types.YChild{"ReactionTrigger", &reactionTriggers.ReactionTrigger[i]}
    }
    reactionTriggers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reactionTriggers.EntityData)
}

// Ipsla_Operation_ReactionTriggers_ReactionTrigger
// Reaction trigger for an operation
type Ipsla_Operation_ReactionTriggers_ReactionTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operation number of the operation generating a
    // trigger. The type is interface{} with range: 1..2048.
    OperationId interface{}

    // Operation number of the operation to be triggered. The type is interface{}
    // with range: 1..2048.
    TriggeredOpId interface{}
}

func (reactionTrigger *Ipsla_Operation_ReactionTriggers_ReactionTrigger) GetEntityData() *types.CommonEntityData {
    reactionTrigger.EntityData.YFilter = reactionTrigger.YFilter
    reactionTrigger.EntityData.YangName = "reaction-trigger"
    reactionTrigger.EntityData.BundleName = "cisco_ios_xr"
    reactionTrigger.EntityData.ParentYangName = "reaction-triggers"
    reactionTrigger.EntityData.SegmentPath = "reaction-trigger" + "[operation-id='" + fmt.Sprintf("%v", reactionTrigger.OperationId) + "']"
    reactionTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reactionTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reactionTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reactionTrigger.EntityData.Children = make(map[string]types.YChild)
    reactionTrigger.EntityData.Leafs = make(map[string]types.YLeaf)
    reactionTrigger.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", reactionTrigger.OperationId}
    reactionTrigger.EntityData.Leafs["triggered-op-id"] = types.YLeaf{"TriggeredOpId", reactionTrigger.TriggeredOpId}
    return &(reactionTrigger.EntityData)
}

// Ipsla_Operation_Definitions
// Operation definition table
type Ipsla_Operation_Definitions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation definition. The type is slice of
    // Ipsla_Operation_Definitions_Definition.
    Definition []Ipsla_Operation_Definitions_Definition
}

func (definitions *Ipsla_Operation_Definitions) GetEntityData() *types.CommonEntityData {
    definitions.EntityData.YFilter = definitions.YFilter
    definitions.EntityData.YangName = "definitions"
    definitions.EntityData.BundleName = "cisco_ios_xr"
    definitions.EntityData.ParentYangName = "operation"
    definitions.EntityData.SegmentPath = "definitions"
    definitions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    definitions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    definitions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    definitions.EntityData.Children = make(map[string]types.YChild)
    definitions.EntityData.Children["definition"] = types.YChild{"Definition", nil}
    for i := range definitions.Definition {
        definitions.EntityData.Children[types.GetSegmentPath(&definitions.Definition[i])] = types.YChild{"Definition", &definitions.Definition[i]}
    }
    definitions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(definitions.EntityData)
}

// Ipsla_Operation_Definitions_Definition
// Operation definition
type Ipsla_Operation_Definitions_Definition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operation number. The type is interface{} with
    // range: 1..2048.
    OperationId interface{}

    // Operation type specification.
    OperationType Ipsla_Operation_Definitions_Definition_OperationType
}

func (definition *Ipsla_Operation_Definitions_Definition) GetEntityData() *types.CommonEntityData {
    definition.EntityData.YFilter = definition.YFilter
    definition.EntityData.YangName = "definition"
    definition.EntityData.BundleName = "cisco_ios_xr"
    definition.EntityData.ParentYangName = "definitions"
    definition.EntityData.SegmentPath = "definition" + "[operation-id='" + fmt.Sprintf("%v", definition.OperationId) + "']"
    definition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    definition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    definition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    definition.EntityData.Children = make(map[string]types.YChild)
    definition.EntityData.Children["operation-type"] = types.YChild{"OperationType", &definition.OperationType}
    definition.EntityData.Leafs = make(map[string]types.YLeaf)
    definition.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", definition.OperationId}
    return &(definition.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType
// Operation type specification
type Ipsla_Operation_Definitions_Definition_OperationType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICMPEcho Operation type.
    IcmpEcho Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho

    // MPLS LSP Ping Operation type.
    MplsLspPing Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing

    // UDPEcho Operation type.
    UdpEcho Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho

    // MPLS LSP Trace Operation type.
    MplsLspTrace Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace

    // UDPJitter Operation type.
    UdpJitter Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter

    // ICMPPathEcho Operation type.
    IcmpPathEcho Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho

    // ICMPPathJitter Operation type.
    IcmpPathJitter Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter
}

func (operationType *Ipsla_Operation_Definitions_Definition_OperationType) GetEntityData() *types.CommonEntityData {
    operationType.EntityData.YFilter = operationType.YFilter
    operationType.EntityData.YangName = "operation-type"
    operationType.EntityData.BundleName = "cisco_ios_xr"
    operationType.EntityData.ParentYangName = "definition"
    operationType.EntityData.SegmentPath = "operation-type"
    operationType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationType.EntityData.Children = make(map[string]types.YChild)
    operationType.EntityData.Children["icmp-echo"] = types.YChild{"IcmpEcho", &operationType.IcmpEcho}
    operationType.EntityData.Children["mpls-lsp-ping"] = types.YChild{"MplsLspPing", &operationType.MplsLspPing}
    operationType.EntityData.Children["udp-echo"] = types.YChild{"UdpEcho", &operationType.UdpEcho}
    operationType.EntityData.Children["mpls-lsp-trace"] = types.YChild{"MplsLspTrace", &operationType.MplsLspTrace}
    operationType.EntityData.Children["udp-jitter"] = types.YChild{"UdpJitter", &operationType.UdpJitter}
    operationType.EntityData.Children["icmp-path-echo"] = types.YChild{"IcmpPathEcho", &operationType.IcmpPathEcho}
    operationType.EntityData.Children["icmp-path-jitter"] = types.YChild{"IcmpPathJitter", &operationType.IcmpPathJitter}
    operationType.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(operationType.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho
// ICMPEcho Operation type
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter IPv6 address of the source device. The type is string.
    SourceAddressV6 interface{}

    // Enter IPv6 address of the destination device. The type is string.
    DestAddressV6 interface{}

    // Enter IPv4 address of the source device. The type is string.
    SourceAddress interface{}

    // Type of service setting in probe packet. The type is interface{} with
    // range: 0..255. The default value is 0.
    Tos interface{}

    // Create operation with specified type. The type is interface{}.
    Create interface{}

    // Configure IPSLA for a VPN Routing/Forwarding instance). The type is string
    // with length: 1..32.
    Vrf interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Probe interval in seconds. The type is interface{} with range: 1..604800.
    // Units are second. The default value is 60.
    Frequency interface{}

    // Enter IPv4 address of the destination device. The type is string.
    DestAddress interface{}

    // Add a tag for this operation. The type is string with length: 1..100.
    Tag interface{}

    // Protocol data size in payload of probe packets.
    DataSize Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_DataSize

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_Statistics

    // Configure the history parameters for this operation.
    History Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_History

    // Table of statistics collection intervals.
    EnhancedStats Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats
}

func (icmpEcho *Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho) GetEntityData() *types.CommonEntityData {
    icmpEcho.EntityData.YFilter = icmpEcho.YFilter
    icmpEcho.EntityData.YangName = "icmp-echo"
    icmpEcho.EntityData.BundleName = "cisco_ios_xr"
    icmpEcho.EntityData.ParentYangName = "operation-type"
    icmpEcho.EntityData.SegmentPath = "icmp-echo"
    icmpEcho.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpEcho.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpEcho.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpEcho.EntityData.Children = make(map[string]types.YChild)
    icmpEcho.EntityData.Children["data-size"] = types.YChild{"DataSize", &icmpEcho.DataSize}
    icmpEcho.EntityData.Children["statistics"] = types.YChild{"Statistics", &icmpEcho.Statistics}
    icmpEcho.EntityData.Children["history"] = types.YChild{"History", &icmpEcho.History}
    icmpEcho.EntityData.Children["enhanced-stats"] = types.YChild{"EnhancedStats", &icmpEcho.EnhancedStats}
    icmpEcho.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpEcho.EntityData.Leafs["source-address-v6"] = types.YLeaf{"SourceAddressV6", icmpEcho.SourceAddressV6}
    icmpEcho.EntityData.Leafs["dest-address-v6"] = types.YLeaf{"DestAddressV6", icmpEcho.DestAddressV6}
    icmpEcho.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpEcho.SourceAddress}
    icmpEcho.EntityData.Leafs["tos"] = types.YLeaf{"Tos", icmpEcho.Tos}
    icmpEcho.EntityData.Leafs["create"] = types.YLeaf{"Create", icmpEcho.Create}
    icmpEcho.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", icmpEcho.Vrf}
    icmpEcho.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", icmpEcho.Timeout}
    icmpEcho.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", icmpEcho.Frequency}
    icmpEcho.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpEcho.DestAddress}
    icmpEcho.EntityData.Leafs["tag"] = types.YLeaf{"Tag", icmpEcho.Tag}
    return &(icmpEcho.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_DataSize
// Protocol data size in payload of probe
// packets
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_DataSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload size in request probe packet. The type is interface{} with range:
    // 0..16384. Units are byte. The default value is 36.
    Request interface{}
}

func (dataSize *Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_DataSize) GetEntityData() *types.CommonEntityData {
    dataSize.EntityData.YFilter = dataSize.YFilter
    dataSize.EntityData.YangName = "data-size"
    dataSize.EntityData.BundleName = "cisco_ios_xr"
    dataSize.EntityData.ParentYangName = "icmp-echo"
    dataSize.EntityData.SegmentPath = "data-size"
    dataSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataSize.EntityData.Children = make(map[string]types.YChild)
    dataSize.EntityData.Leafs = make(map[string]types.YLeaf)
    dataSize.EntityData.Leafs["request"] = types.YLeaf{"Request", dataSize.Request}
    return &(dataSize.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_Statistics
// Statistics collection aggregated over an hour
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..25. Units are hour. The default value is 2.
    Hours interface{}

    // Specify distribution interval in milliseconds. The type is interface{} with
    // range: 1..100. Units are millisecond. The default value is 20.
    DistInterval interface{}

    // Count of distribution intervals maintained. The type is interface{} with
    // range: 1..20. The default value is 1.
    DistCount interface{}
}

func (statistics *Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "icmp-echo"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    statistics.EntityData.Leafs["dist-interval"] = types.YLeaf{"DistInterval", statistics.DistInterval}
    statistics.EntityData.Leafs["dist-count"] = types.YLeaf{"DistCount", statistics.DistCount}
    return &(statistics.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_History
// Configure the history parameters for this
// operation
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify number of lives to be kept. The type is interface{} with range:
    // 0..2. The default value is 0.
    Lives interface{}

    // Choose type of data to be stored in history buffer. The type is
    // IpslaHistoryFilter.
    HistoryFilter interface{}

    // Specify number of history buckets. The type is interface{} with range:
    // 1..60. The default value is 15.
    Buckets interface{}
}

func (history *Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "icmp-echo"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["lives"] = types.YLeaf{"Lives", history.Lives}
    history.EntityData.Leafs["history-filter"] = types.YLeaf{"HistoryFilter", history.HistoryFilter}
    history.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", history.Buckets}
    return &(history.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats
// Table of statistics collection intervals
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a specified time interval. The type is slice of
    // Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats_EnhancedStat.
    EnhancedStat []Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats_EnhancedStat
}

func (enhancedStats *Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats) GetEntityData() *types.CommonEntityData {
    enhancedStats.EntityData.YFilter = enhancedStats.YFilter
    enhancedStats.EntityData.YangName = "enhanced-stats"
    enhancedStats.EntityData.BundleName = "cisco_ios_xr"
    enhancedStats.EntityData.ParentYangName = "icmp-echo"
    enhancedStats.EntityData.SegmentPath = "enhanced-stats"
    enhancedStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStats.EntityData.Children = make(map[string]types.YChild)
    enhancedStats.EntityData.Children["enhanced-stat"] = types.YChild{"EnhancedStat", nil}
    for i := range enhancedStats.EnhancedStat {
        enhancedStats.EntityData.Children[types.GetSegmentPath(&enhancedStats.EnhancedStat[i])] = types.YChild{"EnhancedStat", &enhancedStats.EnhancedStat[i]}
    }
    enhancedStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enhancedStats.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats_EnhancedStat
// Statistics for a specified time interval
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats_EnhancedStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interval in seconds. The type is interface{} with
    // range: 1..3600. Units are second.
    Interval interface{}

    // Buckets of enhanced statistics kept. The type is interface{} with range:
    // 1..100. The default value is 15.
    Buckets interface{}
}

func (enhancedStat *Ipsla_Operation_Definitions_Definition_OperationType_IcmpEcho_EnhancedStats_EnhancedStat) GetEntityData() *types.CommonEntityData {
    enhancedStat.EntityData.YFilter = enhancedStat.YFilter
    enhancedStat.EntityData.YangName = "enhanced-stat"
    enhancedStat.EntityData.BundleName = "cisco_ios_xr"
    enhancedStat.EntityData.ParentYangName = "enhanced-stats"
    enhancedStat.EntityData.SegmentPath = "enhanced-stat" + "[interval='" + fmt.Sprintf("%v", enhancedStat.Interval) + "']"
    enhancedStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStat.EntityData.Children = make(map[string]types.YChild)
    enhancedStat.EntityData.Leafs = make(map[string]types.YLeaf)
    enhancedStat.EntityData.Leafs["interval"] = types.YLeaf{"Interval", enhancedStat.Interval}
    enhancedStat.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", enhancedStat.Buckets}
    return &(enhancedStat.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing
// MPLS LSP Ping Operation type
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time to live value. The type is interface{} with range: 1..255. The default
    // value is 255.
    Ttl interface{}

    // Enter IPv4 address of the source device. The type is string.
    SourceAddress interface{}

    // Echo request output nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    OutputNexthop interface{}

    // Create operation with specified type. The type is interface{}.
    Create interface{}

    // Attributes used for path selection during LSP load balancing. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // The default value is 1.0.0.127.
    LspSelector interface{}

    // EXP bits in MPLS LSP echo reply header. The type is interface{} with range:
    // 0..7. The default value is 0.
    ExpBits interface{}

    // Forced option for the MPLS LSP operation. The type is interface{}.
    ForceExplicitNull interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Echo request output interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Probe interval in seconds. The type is interface{} with range: 1..604800.
    // Units are second. The default value is 60.
    Frequency interface{}

    // Add a tag for this operation. The type is string with length: 1..100.
    Tag interface{}

    // Protocol data size in payload of probe packets.
    DataSize Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_DataSize

    // Echo reply options for the MPLS LSP operation.
    Reply Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Reply

    // Target for the MPLS LSP operation.
    Target Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Statistics

    // Configure the history parameters for this operation.
    History Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_History

    // Table of statistics collection intervals.
    EnhancedStats Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats
}

func (mplsLspPing *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing) GetEntityData() *types.CommonEntityData {
    mplsLspPing.EntityData.YFilter = mplsLspPing.YFilter
    mplsLspPing.EntityData.YangName = "mpls-lsp-ping"
    mplsLspPing.EntityData.BundleName = "cisco_ios_xr"
    mplsLspPing.EntityData.ParentYangName = "operation-type"
    mplsLspPing.EntityData.SegmentPath = "mpls-lsp-ping"
    mplsLspPing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLspPing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLspPing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLspPing.EntityData.Children = make(map[string]types.YChild)
    mplsLspPing.EntityData.Children["data-size"] = types.YChild{"DataSize", &mplsLspPing.DataSize}
    mplsLspPing.EntityData.Children["reply"] = types.YChild{"Reply", &mplsLspPing.Reply}
    mplsLspPing.EntityData.Children["target"] = types.YChild{"Target", &mplsLspPing.Target}
    mplsLspPing.EntityData.Children["statistics"] = types.YChild{"Statistics", &mplsLspPing.Statistics}
    mplsLspPing.EntityData.Children["history"] = types.YChild{"History", &mplsLspPing.History}
    mplsLspPing.EntityData.Children["enhanced-stats"] = types.YChild{"EnhancedStats", &mplsLspPing.EnhancedStats}
    mplsLspPing.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLspPing.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", mplsLspPing.Ttl}
    mplsLspPing.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", mplsLspPing.SourceAddress}
    mplsLspPing.EntityData.Leafs["output-nexthop"] = types.YLeaf{"OutputNexthop", mplsLspPing.OutputNexthop}
    mplsLspPing.EntityData.Leafs["create"] = types.YLeaf{"Create", mplsLspPing.Create}
    mplsLspPing.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", mplsLspPing.LspSelector}
    mplsLspPing.EntityData.Leafs["exp-bits"] = types.YLeaf{"ExpBits", mplsLspPing.ExpBits}
    mplsLspPing.EntityData.Leafs["force-explicit-null"] = types.YLeaf{"ForceExplicitNull", mplsLspPing.ForceExplicitNull}
    mplsLspPing.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", mplsLspPing.Timeout}
    mplsLspPing.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", mplsLspPing.OutputInterface}
    mplsLspPing.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", mplsLspPing.Frequency}
    mplsLspPing.EntityData.Leafs["tag"] = types.YLeaf{"Tag", mplsLspPing.Tag}
    return &(mplsLspPing.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_DataSize
// Protocol data size in payload of probe
// packets
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_DataSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload size in request probe packet. The type is interface{} with range:
    // 100..17986. The default value is 100.
    Request interface{}
}

func (dataSize *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_DataSize) GetEntityData() *types.CommonEntityData {
    dataSize.EntityData.YFilter = dataSize.YFilter
    dataSize.EntityData.YangName = "data-size"
    dataSize.EntityData.BundleName = "cisco_ios_xr"
    dataSize.EntityData.ParentYangName = "mpls-lsp-ping"
    dataSize.EntityData.SegmentPath = "data-size"
    dataSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataSize.EntityData.Children = make(map[string]types.YChild)
    dataSize.EntityData.Leafs = make(map[string]types.YLeaf)
    dataSize.EntityData.Leafs["request"] = types.YLeaf{"Request", dataSize.Request}
    return &(dataSize.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Reply
// Echo reply options for the MPLS LSP
// operation
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables use of router alert in echo reply packets. The type is
    // IpslaLspPingReplyMode.
    Mode interface{}

    // DSCP bits in the reply IP header. The type is one of the following types:
    // enumeration IpslaLspReplyDscp, or int with range: 0..63.
    DscpBits interface{}
}

func (reply *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "mpls-lsp-ping"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = make(map[string]types.YChild)
    reply.EntityData.Leafs = make(map[string]types.YLeaf)
    reply.EntityData.Leafs["mode"] = types.YLeaf{"Mode", reply.Mode}
    reply.EntityData.Leafs["dscp-bits"] = types.YLeaf{"DscpBits", reply.DscpBits}
    return &(reply.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target
// Target for the MPLS LSP operation
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic engineering target.
    TrafficEngineering Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_TrafficEngineering

    // Target specified as an IPv4 address.
    Ipv4 Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4

    // Pseudowire target.
    Pseudowire Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire
}

func (target *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "mpls-lsp-ping"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["traffic-engineering"] = types.YChild{"TrafficEngineering", &target.TrafficEngineering}
    target.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &target.Ipv4}
    target.EntityData.Children["pseudowire"] = types.YChild{"Pseudowire", &target.Pseudowire}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(target.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_TrafficEngineering
// Traffic engineering target
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_TrafficEngineering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel interface number. The type is interface{} with range: 0..65535.
    Tunnel interface{}
}

func (trafficEngineering *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_TrafficEngineering) GetEntityData() *types.CommonEntityData {
    trafficEngineering.EntityData.YFilter = trafficEngineering.YFilter
    trafficEngineering.EntityData.YangName = "traffic-engineering"
    trafficEngineering.EntityData.BundleName = "cisco_ios_xr"
    trafficEngineering.EntityData.ParentYangName = "target"
    trafficEngineering.EntityData.SegmentPath = "traffic-engineering"
    trafficEngineering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficEngineering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficEngineering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficEngineering.EntityData.Children = make(map[string]types.YChild)
    trafficEngineering.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficEngineering.EntityData.Leafs["tunnel"] = types.YLeaf{"Tunnel", trafficEngineering.Tunnel}
    return &(trafficEngineering.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4
// Target specified as an IPv4 address
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target FEC address with mask.
    FecAddress Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4_FecAddress
}

func (ipv4 *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "target"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["fec-address"] = types.YChild{"FecAddress", &ipv4.FecAddress}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4_FecAddress
// Target FEC address with mask
// This type is a presence type.
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4_FecAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address for target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Address interface{}

    // IP netmask for target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Mask interface{}
}

func (fecAddress *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Ipv4_FecAddress) GetEntityData() *types.CommonEntityData {
    fecAddress.EntityData.YFilter = fecAddress.YFilter
    fecAddress.EntityData.YangName = "fec-address"
    fecAddress.EntityData.BundleName = "cisco_ios_xr"
    fecAddress.EntityData.ParentYangName = "ipv4"
    fecAddress.EntityData.SegmentPath = "fec-address"
    fecAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fecAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fecAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fecAddress.EntityData.Children = make(map[string]types.YChild)
    fecAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    fecAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", fecAddress.Address}
    fecAddress.EntityData.Leafs["mask"] = types.YLeaf{"Mask", fecAddress.Mask}
    return &(fecAddress.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire
// Pseudowire target
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target address.
    TargetAddress Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire_TargetAddress
}

func (pseudowire *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire) GetEntityData() *types.CommonEntityData {
    pseudowire.EntityData.YFilter = pseudowire.YFilter
    pseudowire.EntityData.YangName = "pseudowire"
    pseudowire.EntityData.BundleName = "cisco_ios_xr"
    pseudowire.EntityData.ParentYangName = "target"
    pseudowire.EntityData.SegmentPath = "pseudowire"
    pseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowire.EntityData.Children = make(map[string]types.YChild)
    pseudowire.EntityData.Children["target-address"] = types.YChild{"TargetAddress", &pseudowire.TargetAddress}
    pseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pseudowire.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire_TargetAddress
// Target address
// This type is a presence type.
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Address interface{}

    // Virtual circuit ID. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    VcId interface{}
}

func (targetAddress *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Target_Pseudowire_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "pseudowire"
    targetAddress.EntityData.SegmentPath = "target-address"
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = make(map[string]types.YChild)
    targetAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    targetAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", targetAddress.Address}
    targetAddress.EntityData.Leafs["vc-id"] = types.YLeaf{"VcId", targetAddress.VcId}
    return &(targetAddress.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Statistics
// Statistics collection aggregated over an hour
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..25. Units are hour. The default value is 2.
    Hours interface{}

    // Specify distribution interval in milliseconds. The type is interface{} with
    // range: 1..100. Units are millisecond. The default value is 20.
    DistInterval interface{}

    // Count of distribution intervals maintained. The type is interface{} with
    // range: 1..20. The default value is 1.
    DistCount interface{}
}

func (statistics *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "mpls-lsp-ping"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    statistics.EntityData.Leafs["dist-interval"] = types.YLeaf{"DistInterval", statistics.DistInterval}
    statistics.EntityData.Leafs["dist-count"] = types.YLeaf{"DistCount", statistics.DistCount}
    return &(statistics.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_History
// Configure the history parameters for this
// operation
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify number of lives to be kept. The type is interface{} with range:
    // 0..2. The default value is 0.
    Lives interface{}

    // Choose type of data to be stored in history buffer. The type is
    // IpslaHistoryFilter.
    HistoryFilter interface{}

    // Specify number of history buckets. The type is interface{} with range:
    // 1..60. The default value is 15.
    Buckets interface{}
}

func (history *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "mpls-lsp-ping"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["lives"] = types.YLeaf{"Lives", history.Lives}
    history.EntityData.Leafs["history-filter"] = types.YLeaf{"HistoryFilter", history.HistoryFilter}
    history.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", history.Buckets}
    return &(history.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats
// Table of statistics collection intervals
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a specified time interval. The type is slice of
    // Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats_EnhancedStat.
    EnhancedStat []Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats_EnhancedStat
}

func (enhancedStats *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats) GetEntityData() *types.CommonEntityData {
    enhancedStats.EntityData.YFilter = enhancedStats.YFilter
    enhancedStats.EntityData.YangName = "enhanced-stats"
    enhancedStats.EntityData.BundleName = "cisco_ios_xr"
    enhancedStats.EntityData.ParentYangName = "mpls-lsp-ping"
    enhancedStats.EntityData.SegmentPath = "enhanced-stats"
    enhancedStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStats.EntityData.Children = make(map[string]types.YChild)
    enhancedStats.EntityData.Children["enhanced-stat"] = types.YChild{"EnhancedStat", nil}
    for i := range enhancedStats.EnhancedStat {
        enhancedStats.EntityData.Children[types.GetSegmentPath(&enhancedStats.EnhancedStat[i])] = types.YChild{"EnhancedStat", &enhancedStats.EnhancedStat[i]}
    }
    enhancedStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enhancedStats.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats_EnhancedStat
// Statistics for a specified time interval
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats_EnhancedStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interval in seconds. The type is interface{} with
    // range: 1..3600. Units are second.
    Interval interface{}

    // Buckets of enhanced statistics kept. The type is interface{} with range:
    // 1..100. The default value is 15.
    Buckets interface{}
}

func (enhancedStat *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspPing_EnhancedStats_EnhancedStat) GetEntityData() *types.CommonEntityData {
    enhancedStat.EntityData.YFilter = enhancedStat.YFilter
    enhancedStat.EntityData.YangName = "enhanced-stat"
    enhancedStat.EntityData.BundleName = "cisco_ios_xr"
    enhancedStat.EntityData.ParentYangName = "enhanced-stats"
    enhancedStat.EntityData.SegmentPath = "enhanced-stat" + "[interval='" + fmt.Sprintf("%v", enhancedStat.Interval) + "']"
    enhancedStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStat.EntityData.Children = make(map[string]types.YChild)
    enhancedStat.EntityData.Leafs = make(map[string]types.YLeaf)
    enhancedStat.EntityData.Leafs["interval"] = types.YLeaf{"Interval", enhancedStat.Interval}
    enhancedStat.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", enhancedStat.Buckets}
    return &(enhancedStat.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho
// UDPEcho Operation type
type Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter IPv4 address of the source device. The type is string.
    SourceAddress interface{}

    // Type of service setting in probe packet. The type is interface{} with
    // range: 0..255. The default value is 0.
    Tos interface{}

    // Disable control packets. The type is interface{}.
    ControlDisable interface{}

    // Port number on source device. The type is interface{} with range: 0..65535.
    SourcePort interface{}

    // Create operation with specified type. The type is interface{}.
    Create interface{}

    // Configure IPSLA for a VPN Routing/Forwarding instance). The type is string
    // with length: 1..32.
    Vrf interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Probe interval in seconds. The type is interface{} with range: 1..604800.
    // Units are second. The default value is 60.
    Frequency interface{}

    // Port number on target device. The type is interface{} with range: 0..65535.
    DestPort interface{}

    // Check each IPSLA response for corruption. The type is interface{}.
    VerifyData interface{}

    // Enter IPv4 address of the destination device. The type is string.
    DestAddress interface{}

    // Add a tag for this operation. The type is string with length: 1..100.
    Tag interface{}

    // Protocol data size in payload of probe packets.
    DataSize Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_DataSize

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_Statistics

    // Configure the history parameters for this operation.
    History Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_History

    // Table of statistics collection intervals.
    EnhancedStats Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats
}

func (udpEcho *Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho) GetEntityData() *types.CommonEntityData {
    udpEcho.EntityData.YFilter = udpEcho.YFilter
    udpEcho.EntityData.YangName = "udp-echo"
    udpEcho.EntityData.BundleName = "cisco_ios_xr"
    udpEcho.EntityData.ParentYangName = "operation-type"
    udpEcho.EntityData.SegmentPath = "udp-echo"
    udpEcho.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpEcho.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpEcho.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpEcho.EntityData.Children = make(map[string]types.YChild)
    udpEcho.EntityData.Children["data-size"] = types.YChild{"DataSize", &udpEcho.DataSize}
    udpEcho.EntityData.Children["statistics"] = types.YChild{"Statistics", &udpEcho.Statistics}
    udpEcho.EntityData.Children["history"] = types.YChild{"History", &udpEcho.History}
    udpEcho.EntityData.Children["enhanced-stats"] = types.YChild{"EnhancedStats", &udpEcho.EnhancedStats}
    udpEcho.EntityData.Leafs = make(map[string]types.YLeaf)
    udpEcho.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", udpEcho.SourceAddress}
    udpEcho.EntityData.Leafs["tos"] = types.YLeaf{"Tos", udpEcho.Tos}
    udpEcho.EntityData.Leafs["control-disable"] = types.YLeaf{"ControlDisable", udpEcho.ControlDisable}
    udpEcho.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", udpEcho.SourcePort}
    udpEcho.EntityData.Leafs["create"] = types.YLeaf{"Create", udpEcho.Create}
    udpEcho.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", udpEcho.Vrf}
    udpEcho.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", udpEcho.Timeout}
    udpEcho.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", udpEcho.Frequency}
    udpEcho.EntityData.Leafs["dest-port"] = types.YLeaf{"DestPort", udpEcho.DestPort}
    udpEcho.EntityData.Leafs["verify-data"] = types.YLeaf{"VerifyData", udpEcho.VerifyData}
    udpEcho.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", udpEcho.DestAddress}
    udpEcho.EntityData.Leafs["tag"] = types.YLeaf{"Tag", udpEcho.Tag}
    return &(udpEcho.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_DataSize
// Protocol data size in payload of probe
// packets
type Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_DataSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload size in request probe packet. The type is interface{} with range:
    // 16..1500. Units are byte. The default value is 16.
    Request interface{}
}

func (dataSize *Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_DataSize) GetEntityData() *types.CommonEntityData {
    dataSize.EntityData.YFilter = dataSize.YFilter
    dataSize.EntityData.YangName = "data-size"
    dataSize.EntityData.BundleName = "cisco_ios_xr"
    dataSize.EntityData.ParentYangName = "udp-echo"
    dataSize.EntityData.SegmentPath = "data-size"
    dataSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataSize.EntityData.Children = make(map[string]types.YChild)
    dataSize.EntityData.Leafs = make(map[string]types.YLeaf)
    dataSize.EntityData.Leafs["request"] = types.YLeaf{"Request", dataSize.Request}
    return &(dataSize.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_Statistics
// Statistics collection aggregated over an hour
type Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..25. Units are hour. The default value is 2.
    Hours interface{}

    // Specify distribution interval in milliseconds. The type is interface{} with
    // range: 1..100. Units are millisecond. The default value is 20.
    DistInterval interface{}

    // Count of distribution intervals maintained. The type is interface{} with
    // range: 1..20. The default value is 1.
    DistCount interface{}
}

func (statistics *Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "udp-echo"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    statistics.EntityData.Leafs["dist-interval"] = types.YLeaf{"DistInterval", statistics.DistInterval}
    statistics.EntityData.Leafs["dist-count"] = types.YLeaf{"DistCount", statistics.DistCount}
    return &(statistics.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_History
// Configure the history parameters for this
// operation
type Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify number of lives to be kept. The type is interface{} with range:
    // 0..2. The default value is 0.
    Lives interface{}

    // Choose type of data to be stored in history buffer. The type is
    // IpslaHistoryFilter.
    HistoryFilter interface{}

    // Specify number of history buckets. The type is interface{} with range:
    // 1..60. The default value is 15.
    Buckets interface{}
}

func (history *Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "udp-echo"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["lives"] = types.YLeaf{"Lives", history.Lives}
    history.EntityData.Leafs["history-filter"] = types.YLeaf{"HistoryFilter", history.HistoryFilter}
    history.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", history.Buckets}
    return &(history.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats
// Table of statistics collection intervals
type Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a specified time interval. The type is slice of
    // Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats_EnhancedStat.
    EnhancedStat []Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats_EnhancedStat
}

func (enhancedStats *Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats) GetEntityData() *types.CommonEntityData {
    enhancedStats.EntityData.YFilter = enhancedStats.YFilter
    enhancedStats.EntityData.YangName = "enhanced-stats"
    enhancedStats.EntityData.BundleName = "cisco_ios_xr"
    enhancedStats.EntityData.ParentYangName = "udp-echo"
    enhancedStats.EntityData.SegmentPath = "enhanced-stats"
    enhancedStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStats.EntityData.Children = make(map[string]types.YChild)
    enhancedStats.EntityData.Children["enhanced-stat"] = types.YChild{"EnhancedStat", nil}
    for i := range enhancedStats.EnhancedStat {
        enhancedStats.EntityData.Children[types.GetSegmentPath(&enhancedStats.EnhancedStat[i])] = types.YChild{"EnhancedStat", &enhancedStats.EnhancedStat[i]}
    }
    enhancedStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enhancedStats.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats_EnhancedStat
// Statistics for a specified time interval
type Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats_EnhancedStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interval in seconds. The type is interface{} with
    // range: 1..3600. Units are second.
    Interval interface{}

    // Buckets of enhanced statistics kept. The type is interface{} with range:
    // 1..100. The default value is 15.
    Buckets interface{}
}

func (enhancedStat *Ipsla_Operation_Definitions_Definition_OperationType_UdpEcho_EnhancedStats_EnhancedStat) GetEntityData() *types.CommonEntityData {
    enhancedStat.EntityData.YFilter = enhancedStat.YFilter
    enhancedStat.EntityData.YangName = "enhanced-stat"
    enhancedStat.EntityData.BundleName = "cisco_ios_xr"
    enhancedStat.EntityData.ParentYangName = "enhanced-stats"
    enhancedStat.EntityData.SegmentPath = "enhanced-stat" + "[interval='" + fmt.Sprintf("%v", enhancedStat.Interval) + "']"
    enhancedStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStat.EntityData.Children = make(map[string]types.YChild)
    enhancedStat.EntityData.Leafs = make(map[string]types.YLeaf)
    enhancedStat.EntityData.Leafs["interval"] = types.YLeaf{"Interval", enhancedStat.Interval}
    enhancedStat.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", enhancedStat.Buckets}
    return &(enhancedStat.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace
// MPLS LSP Trace Operation type
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time to live value. The type is interface{} with range: 1..255. The default
    // value is 30.
    Ttl interface{}

    // Enter IPv4 address of the source device. The type is string.
    SourceAddress interface{}

    // Echo request output nexthop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    OutputNexthop interface{}

    // Create operation with specified type. The type is interface{}.
    Create interface{}

    // Attributes used for path selection during LSP load balancing. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // The default value is 1.0.0.127.
    LspSelector interface{}

    // EXP bits in MPLS LSP echo reply header. The type is interface{} with range:
    // 0..7. The default value is 0.
    ExpBits interface{}

    // Forced option for the MPLS LSP operation. The type is interface{}.
    ForceExplicitNull interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Echo request output interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Probe interval in seconds. The type is interface{} with range: 1..604800.
    // Units are second. The default value is 60.
    Frequency interface{}

    // Add a tag for this operation. The type is string with length: 1..100.
    Tag interface{}

    // Target for the MPLS LSP operation.
    Target Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target

    // Echo reply options for the MPLS LSP operation.
    Reply Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Reply

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Statistics

    // Configure the history parameters for this operation.
    History Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_History
}

func (mplsLspTrace *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace) GetEntityData() *types.CommonEntityData {
    mplsLspTrace.EntityData.YFilter = mplsLspTrace.YFilter
    mplsLspTrace.EntityData.YangName = "mpls-lsp-trace"
    mplsLspTrace.EntityData.BundleName = "cisco_ios_xr"
    mplsLspTrace.EntityData.ParentYangName = "operation-type"
    mplsLspTrace.EntityData.SegmentPath = "mpls-lsp-trace"
    mplsLspTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLspTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLspTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLspTrace.EntityData.Children = make(map[string]types.YChild)
    mplsLspTrace.EntityData.Children["target"] = types.YChild{"Target", &mplsLspTrace.Target}
    mplsLspTrace.EntityData.Children["reply"] = types.YChild{"Reply", &mplsLspTrace.Reply}
    mplsLspTrace.EntityData.Children["statistics"] = types.YChild{"Statistics", &mplsLspTrace.Statistics}
    mplsLspTrace.EntityData.Children["history"] = types.YChild{"History", &mplsLspTrace.History}
    mplsLspTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLspTrace.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", mplsLspTrace.Ttl}
    mplsLspTrace.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", mplsLspTrace.SourceAddress}
    mplsLspTrace.EntityData.Leafs["output-nexthop"] = types.YLeaf{"OutputNexthop", mplsLspTrace.OutputNexthop}
    mplsLspTrace.EntityData.Leafs["create"] = types.YLeaf{"Create", mplsLspTrace.Create}
    mplsLspTrace.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", mplsLspTrace.LspSelector}
    mplsLspTrace.EntityData.Leafs["exp-bits"] = types.YLeaf{"ExpBits", mplsLspTrace.ExpBits}
    mplsLspTrace.EntityData.Leafs["force-explicit-null"] = types.YLeaf{"ForceExplicitNull", mplsLspTrace.ForceExplicitNull}
    mplsLspTrace.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", mplsLspTrace.Timeout}
    mplsLspTrace.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", mplsLspTrace.OutputInterface}
    mplsLspTrace.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", mplsLspTrace.Frequency}
    mplsLspTrace.EntityData.Leafs["tag"] = types.YLeaf{"Tag", mplsLspTrace.Tag}
    return &(mplsLspTrace.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target
// Target for the MPLS LSP operation
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic engineering target.
    TrafficEngineering Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_TrafficEngineering

    // Target specified as an IPv4 address.
    Ipv4 Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4
}

func (target *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "mpls-lsp-trace"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["traffic-engineering"] = types.YChild{"TrafficEngineering", &target.TrafficEngineering}
    target.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &target.Ipv4}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(target.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_TrafficEngineering
// Traffic engineering target
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_TrafficEngineering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel interface number. The type is interface{} with range: 0..65535.
    Tunnel interface{}
}

func (trafficEngineering *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_TrafficEngineering) GetEntityData() *types.CommonEntityData {
    trafficEngineering.EntityData.YFilter = trafficEngineering.YFilter
    trafficEngineering.EntityData.YangName = "traffic-engineering"
    trafficEngineering.EntityData.BundleName = "cisco_ios_xr"
    trafficEngineering.EntityData.ParentYangName = "target"
    trafficEngineering.EntityData.SegmentPath = "traffic-engineering"
    trafficEngineering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficEngineering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficEngineering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficEngineering.EntityData.Children = make(map[string]types.YChild)
    trafficEngineering.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficEngineering.EntityData.Leafs["tunnel"] = types.YLeaf{"Tunnel", trafficEngineering.Tunnel}
    return &(trafficEngineering.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4
// Target specified as an IPv4 address
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target FEC address with mask.
    FecAddress Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4_FecAddress
}

func (ipv4 *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "target"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["fec-address"] = types.YChild{"FecAddress", &ipv4.FecAddress}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4_FecAddress
// Target FEC address with mask
// This type is a presence type.
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4_FecAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address for target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Address interface{}

    // IP netmask for target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Mask interface{}
}

func (fecAddress *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Target_Ipv4_FecAddress) GetEntityData() *types.CommonEntityData {
    fecAddress.EntityData.YFilter = fecAddress.YFilter
    fecAddress.EntityData.YangName = "fec-address"
    fecAddress.EntityData.BundleName = "cisco_ios_xr"
    fecAddress.EntityData.ParentYangName = "ipv4"
    fecAddress.EntityData.SegmentPath = "fec-address"
    fecAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fecAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fecAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fecAddress.EntityData.Children = make(map[string]types.YChild)
    fecAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    fecAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", fecAddress.Address}
    fecAddress.EntityData.Leafs["mask"] = types.YLeaf{"Mask", fecAddress.Mask}
    return &(fecAddress.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Reply
// Echo reply options for the MPLS LSP
// operation
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables use of router alert in echo reply packets. The type is
    // IpslaLspTraceReplyMode.
    Mode interface{}

    // DSCP bits in the reply IP header. The type is one of the following types:
    // enumeration IpslaLspReplyDscp, or int with range: 0..63.
    DscpBits interface{}
}

func (reply *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "mpls-lsp-trace"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = make(map[string]types.YChild)
    reply.EntityData.Leafs = make(map[string]types.YLeaf)
    reply.EntityData.Leafs["mode"] = types.YLeaf{"Mode", reply.Mode}
    reply.EntityData.Leafs["dscp-bits"] = types.YLeaf{"DscpBits", reply.DscpBits}
    return &(reply.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Statistics
// Statistics collection aggregated over an hour
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..25. Units are hour. The default value is 2.
    Hours interface{}

    // Specify distribution interval in milliseconds. The type is interface{} with
    // range: 1..100. Units are millisecond. The default value is 20.
    DistInterval interface{}

    // Count of distribution intervals maintained. The type is interface{} with
    // range: 1..20. The default value is 1.
    DistCount interface{}
}

func (statistics *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "mpls-lsp-trace"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    statistics.EntityData.Leafs["dist-interval"] = types.YLeaf{"DistInterval", statistics.DistInterval}
    statistics.EntityData.Leafs["dist-count"] = types.YLeaf{"DistCount", statistics.DistCount}
    return &(statistics.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_History
// Configure the history parameters for this
// operation
type Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify number of lives to be kept. The type is interface{} with range:
    // 0..2. The default value is 0.
    Lives interface{}

    // Choose type of data to be stored in history buffer. The type is
    // IpslaHistoryFilter.
    HistoryFilter interface{}

    // Specify number of history buckets. The type is interface{} with range:
    // 1..60. The default value is 15.
    Buckets interface{}
}

func (history *Ipsla_Operation_Definitions_Definition_OperationType_MplsLspTrace_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "mpls-lsp-trace"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["lives"] = types.YLeaf{"Lives", history.Lives}
    history.EntityData.Leafs["history-filter"] = types.YLeaf{"HistoryFilter", history.HistoryFilter}
    history.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", history.Buckets}
    return &(history.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter
// UDPJitter Operation type
type Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter IPv4 address of the source device. The type is string.
    SourceAddress interface{}

    // Type of service setting in probe packet. The type is interface{} with
    // range: 0..255. The default value is 0.
    Tos interface{}

    // Disable control packets. The type is interface{}.
    ControlDisable interface{}

    // Port number on source device. The type is interface{} with range: 0..65535.
    SourcePort interface{}

    // Create operation with specified type. The type is interface{}.
    Create interface{}

    // Configure IPSLA for a VPN Routing/Forwarding instance). The type is string
    // with length: 1..32.
    Vrf interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Probe interval in seconds. The type is interface{} with range: 1..604800.
    // Units are second. The default value is 60.
    Frequency interface{}

    // Port number on target device. The type is interface{} with range: 0..65535.
    DestPort interface{}

    // Check each IPSLA response for corruption. The type is interface{}.
    VerifyData interface{}

    // Enter IPv4 address of the destination device. The type is string.
    DestAddress interface{}

    // Add a tag for this operation. The type is string with length: 1..100.
    Tag interface{}

    // Protocol data size in payload of probe packets.
    DataSize Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_DataSize

    // Probe packet stream configuration parameters.
    Packet Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Packet

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Statistics

    // Table of statistics collection intervals.
    EnhancedStats Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats
}

func (udpJitter *Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter) GetEntityData() *types.CommonEntityData {
    udpJitter.EntityData.YFilter = udpJitter.YFilter
    udpJitter.EntityData.YangName = "udp-jitter"
    udpJitter.EntityData.BundleName = "cisco_ios_xr"
    udpJitter.EntityData.ParentYangName = "operation-type"
    udpJitter.EntityData.SegmentPath = "udp-jitter"
    udpJitter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitter.EntityData.Children = make(map[string]types.YChild)
    udpJitter.EntityData.Children["data-size"] = types.YChild{"DataSize", &udpJitter.DataSize}
    udpJitter.EntityData.Children["packet"] = types.YChild{"Packet", &udpJitter.Packet}
    udpJitter.EntityData.Children["statistics"] = types.YChild{"Statistics", &udpJitter.Statistics}
    udpJitter.EntityData.Children["enhanced-stats"] = types.YChild{"EnhancedStats", &udpJitter.EnhancedStats}
    udpJitter.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitter.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", udpJitter.SourceAddress}
    udpJitter.EntityData.Leafs["tos"] = types.YLeaf{"Tos", udpJitter.Tos}
    udpJitter.EntityData.Leafs["control-disable"] = types.YLeaf{"ControlDisable", udpJitter.ControlDisable}
    udpJitter.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", udpJitter.SourcePort}
    udpJitter.EntityData.Leafs["create"] = types.YLeaf{"Create", udpJitter.Create}
    udpJitter.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", udpJitter.Vrf}
    udpJitter.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", udpJitter.Timeout}
    udpJitter.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", udpJitter.Frequency}
    udpJitter.EntityData.Leafs["dest-port"] = types.YLeaf{"DestPort", udpJitter.DestPort}
    udpJitter.EntityData.Leafs["verify-data"] = types.YLeaf{"VerifyData", udpJitter.VerifyData}
    udpJitter.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", udpJitter.DestAddress}
    udpJitter.EntityData.Leafs["tag"] = types.YLeaf{"Tag", udpJitter.Tag}
    return &(udpJitter.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_DataSize
// Protocol data size in payload of probe
// packets
type Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_DataSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload size in request probe packet. The type is interface{} with range:
    // 28..1500. Units are byte. The default value is 32.
    Request interface{}
}

func (dataSize *Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_DataSize) GetEntityData() *types.CommonEntityData {
    dataSize.EntityData.YFilter = dataSize.YFilter
    dataSize.EntityData.YangName = "data-size"
    dataSize.EntityData.BundleName = "cisco_ios_xr"
    dataSize.EntityData.ParentYangName = "udp-jitter"
    dataSize.EntityData.SegmentPath = "data-size"
    dataSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataSize.EntityData.Children = make(map[string]types.YChild)
    dataSize.EntityData.Leafs = make(map[string]types.YLeaf)
    dataSize.EntityData.Leafs["request"] = types.YLeaf{"Request", dataSize.Request}
    return &(dataSize.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Packet
// Probe packet stream configuration
// parameters
type Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Packet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets to be transmitted during a probe. The type is interface{}
    // with range: 1..60000. The default value is 10.
    Count interface{}

    // Packet interval in milliseconds. The type is interface{} with range:
    // 1..60000. Units are millisecond. The default value is 20.
    Interval interface{}
}

func (packet *Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Packet) GetEntityData() *types.CommonEntityData {
    packet.EntityData.YFilter = packet.YFilter
    packet.EntityData.YangName = "packet"
    packet.EntityData.BundleName = "cisco_ios_xr"
    packet.EntityData.ParentYangName = "udp-jitter"
    packet.EntityData.SegmentPath = "packet"
    packet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packet.EntityData.Children = make(map[string]types.YChild)
    packet.EntityData.Leafs = make(map[string]types.YLeaf)
    packet.EntityData.Leafs["count"] = types.YLeaf{"Count", packet.Count}
    packet.EntityData.Leafs["interval"] = types.YLeaf{"Interval", packet.Interval}
    return &(packet.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Statistics
// Statistics collection aggregated over an hour
type Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..25. Units are hour. The default value is 2.
    Hours interface{}

    // Specify distribution interval in milliseconds. The type is interface{} with
    // range: 1..100. Units are millisecond. The default value is 20.
    DistInterval interface{}

    // Count of distribution intervals maintained. The type is interface{} with
    // range: 1..20. The default value is 1.
    DistCount interface{}
}

func (statistics *Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "udp-jitter"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    statistics.EntityData.Leafs["dist-interval"] = types.YLeaf{"DistInterval", statistics.DistInterval}
    statistics.EntityData.Leafs["dist-count"] = types.YLeaf{"DistCount", statistics.DistCount}
    return &(statistics.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats
// Table of statistics collection intervals
type Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a specified time interval. The type is slice of
    // Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats_EnhancedStat.
    EnhancedStat []Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats_EnhancedStat
}

func (enhancedStats *Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats) GetEntityData() *types.CommonEntityData {
    enhancedStats.EntityData.YFilter = enhancedStats.YFilter
    enhancedStats.EntityData.YangName = "enhanced-stats"
    enhancedStats.EntityData.BundleName = "cisco_ios_xr"
    enhancedStats.EntityData.ParentYangName = "udp-jitter"
    enhancedStats.EntityData.SegmentPath = "enhanced-stats"
    enhancedStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStats.EntityData.Children = make(map[string]types.YChild)
    enhancedStats.EntityData.Children["enhanced-stat"] = types.YChild{"EnhancedStat", nil}
    for i := range enhancedStats.EnhancedStat {
        enhancedStats.EntityData.Children[types.GetSegmentPath(&enhancedStats.EnhancedStat[i])] = types.YChild{"EnhancedStat", &enhancedStats.EnhancedStat[i]}
    }
    enhancedStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enhancedStats.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats_EnhancedStat
// Statistics for a specified time interval
type Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats_EnhancedStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interval in seconds. The type is interface{} with
    // range: 1..3600. Units are second.
    Interval interface{}

    // Buckets of enhanced statistics kept. The type is interface{} with range:
    // 1..100. The default value is 15.
    Buckets interface{}
}

func (enhancedStat *Ipsla_Operation_Definitions_Definition_OperationType_UdpJitter_EnhancedStats_EnhancedStat) GetEntityData() *types.CommonEntityData {
    enhancedStat.EntityData.YFilter = enhancedStat.YFilter
    enhancedStat.EntityData.YangName = "enhanced-stat"
    enhancedStat.EntityData.BundleName = "cisco_ios_xr"
    enhancedStat.EntityData.ParentYangName = "enhanced-stats"
    enhancedStat.EntityData.SegmentPath = "enhanced-stat" + "[interval='" + fmt.Sprintf("%v", enhancedStat.Interval) + "']"
    enhancedStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedStat.EntityData.Children = make(map[string]types.YChild)
    enhancedStat.EntityData.Leafs = make(map[string]types.YLeaf)
    enhancedStat.EntityData.Leafs["interval"] = types.YLeaf{"Interval", enhancedStat.Interval}
    enhancedStat.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", enhancedStat.Buckets}
    return &(enhancedStat.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho
// ICMPPathEcho Operation type
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter IPv4 address of the source device. The type is string.
    SourceAddress interface{}

    // Type of service setting in probe packet. The type is interface{} with
    // range: 0..255. The default value is 0.
    Tos interface{}

    // Create operation with specified type. The type is interface{}.
    Create interface{}

    // Configure IPSLA for a VPN Routing/Forwarding instance). The type is string
    // with length: 1..32.
    Vrf interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Probe interval in seconds. The type is interface{} with range: 1..604800.
    // Units are second. The default value is 60.
    Frequency interface{}

    // Enter IPv4 address of the destination device. The type is string.
    DestAddress interface{}

    // Add a tag for this operation. The type is string with length: 1..100.
    Tag interface{}

    // Configure the history parameters for this operation.
    History Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_History

    // Protocol data size in payload of probe packets.
    DataSize Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_DataSize

    // Statistics collection aggregated over an hour.
    Statistics Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_Statistics

    // Loose source routing path (up to 8 intermediate nodes).
    LsrPath Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_LsrPath
}

func (icmpPathEcho *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho) GetEntityData() *types.CommonEntityData {
    icmpPathEcho.EntityData.YFilter = icmpPathEcho.YFilter
    icmpPathEcho.EntityData.YangName = "icmp-path-echo"
    icmpPathEcho.EntityData.BundleName = "cisco_ios_xr"
    icmpPathEcho.EntityData.ParentYangName = "operation-type"
    icmpPathEcho.EntityData.SegmentPath = "icmp-path-echo"
    icmpPathEcho.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathEcho.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathEcho.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathEcho.EntityData.Children = make(map[string]types.YChild)
    icmpPathEcho.EntityData.Children["history"] = types.YChild{"History", &icmpPathEcho.History}
    icmpPathEcho.EntityData.Children["data-size"] = types.YChild{"DataSize", &icmpPathEcho.DataSize}
    icmpPathEcho.EntityData.Children["statistics"] = types.YChild{"Statistics", &icmpPathEcho.Statistics}
    icmpPathEcho.EntityData.Children["lsr-path"] = types.YChild{"LsrPath", &icmpPathEcho.LsrPath}
    icmpPathEcho.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathEcho.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathEcho.SourceAddress}
    icmpPathEcho.EntityData.Leafs["tos"] = types.YLeaf{"Tos", icmpPathEcho.Tos}
    icmpPathEcho.EntityData.Leafs["create"] = types.YLeaf{"Create", icmpPathEcho.Create}
    icmpPathEcho.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", icmpPathEcho.Vrf}
    icmpPathEcho.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", icmpPathEcho.Timeout}
    icmpPathEcho.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", icmpPathEcho.Frequency}
    icmpPathEcho.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathEcho.DestAddress}
    icmpPathEcho.EntityData.Leafs["tag"] = types.YLeaf{"Tag", icmpPathEcho.Tag}
    return &(icmpPathEcho.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_History
// Configure the history parameters for this
// operation
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify number of samples to keep. The type is interface{} with range:
    // 1..30. The default value is 16.
    Samples interface{}

    // Specify number of history buckets. The type is interface{} with range:
    // 1..60. The default value is 15.
    Buckets interface{}

    // Choose type of data to be stored in history buffer. The type is
    // IpslaHistoryFilter.
    HistoryFilter interface{}

    // Specify number of lives to be kept. The type is interface{} with range:
    // 0..2. The default value is 0.
    Lives interface{}
}

func (history *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "icmp-path-echo"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["samples"] = types.YLeaf{"Samples", history.Samples}
    history.EntityData.Leafs["buckets"] = types.YLeaf{"Buckets", history.Buckets}
    history.EntityData.Leafs["history-filter"] = types.YLeaf{"HistoryFilter", history.HistoryFilter}
    history.EntityData.Leafs["lives"] = types.YLeaf{"Lives", history.Lives}
    return &(history.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_DataSize
// Protocol data size in payload of probe
// packets
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_DataSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload size in request probe packet. The type is interface{} with range:
    // 0..16384. Units are byte. The default value is 36.
    Request interface{}
}

func (dataSize *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_DataSize) GetEntityData() *types.CommonEntityData {
    dataSize.EntityData.YFilter = dataSize.YFilter
    dataSize.EntityData.YangName = "data-size"
    dataSize.EntityData.BundleName = "cisco_ios_xr"
    dataSize.EntityData.ParentYangName = "icmp-path-echo"
    dataSize.EntityData.SegmentPath = "data-size"
    dataSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataSize.EntityData.Children = make(map[string]types.YChild)
    dataSize.EntityData.Leafs = make(map[string]types.YLeaf)
    dataSize.EntityData.Leafs["request"] = types.YLeaf{"Request", dataSize.Request}
    return &(dataSize.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_Statistics
// Statistics collection aggregated over an
// hour
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of paths for which statistics are kept. The type is
    // interface{} with range: 1..128. The default value is 5.
    Paths interface{}

    // Specify distribution interval in milliseconds. The type is interface{} with
    // range: 1..100. Units are millisecond. The default value is 20.
    DistInterval interface{}

    // Count of distribution intervals maintained. The type is interface{} with
    // range: 1..20. The default value is 1.
    DistCount interface{}

    // Number of hours for which hourly statistics are kept. The type is
    // interface{} with range: 0..25. Units are hour. The default value is 2.
    Hours interface{}

    // Maximum hops per path for which statistics are kept. The type is
    // interface{} with range: 1..30. The default value is 16.
    Hops interface{}
}

func (statistics *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "icmp-path-echo"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["paths"] = types.YLeaf{"Paths", statistics.Paths}
    statistics.EntityData.Leafs["dist-interval"] = types.YLeaf{"DistInterval", statistics.DistInterval}
    statistics.EntityData.Leafs["dist-count"] = types.YLeaf{"DistCount", statistics.DistCount}
    statistics.EntityData.Leafs["hours"] = types.YLeaf{"Hours", statistics.Hours}
    statistics.EntityData.Leafs["hops"] = types.YLeaf{"Hops", statistics.Hops}
    return &(statistics.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_LsrPath
// Loose source routing path (up to 8 intermediate
// nodes)
// This type is a presence type.
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_LsrPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Node1 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node2 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node3 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node4 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node5 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node6 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node7 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node8 interface{}
}

func (lsrPath *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathEcho_LsrPath) GetEntityData() *types.CommonEntityData {
    lsrPath.EntityData.YFilter = lsrPath.YFilter
    lsrPath.EntityData.YangName = "lsr-path"
    lsrPath.EntityData.BundleName = "cisco_ios_xr"
    lsrPath.EntityData.ParentYangName = "icmp-path-echo"
    lsrPath.EntityData.SegmentPath = "lsr-path"
    lsrPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsrPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsrPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsrPath.EntityData.Children = make(map[string]types.YChild)
    lsrPath.EntityData.Leafs = make(map[string]types.YLeaf)
    lsrPath.EntityData.Leafs["node1"] = types.YLeaf{"Node1", lsrPath.Node1}
    lsrPath.EntityData.Leafs["node2"] = types.YLeaf{"Node2", lsrPath.Node2}
    lsrPath.EntityData.Leafs["node3"] = types.YLeaf{"Node3", lsrPath.Node3}
    lsrPath.EntityData.Leafs["node4"] = types.YLeaf{"Node4", lsrPath.Node4}
    lsrPath.EntityData.Leafs["node5"] = types.YLeaf{"Node5", lsrPath.Node5}
    lsrPath.EntityData.Leafs["node6"] = types.YLeaf{"Node6", lsrPath.Node6}
    lsrPath.EntityData.Leafs["node7"] = types.YLeaf{"Node7", lsrPath.Node7}
    lsrPath.EntityData.Leafs["node8"] = types.YLeaf{"Node8", lsrPath.Node8}
    return &(lsrPath.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter
// ICMPPathJitter Operation type
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter IPv4 address of the source device. The type is string.
    SourceAddress interface{}

    // Type of service setting in probe packet. The type is interface{} with
    // range: 0..255. The default value is 0.
    Tos interface{}

    // Create operation with specified type. The type is interface{}.
    Create interface{}

    // Configure IPSLA for a VPN Routing/Forwarding instance). The type is string
    // with length: 1..32.
    Vrf interface{}

    // Probe/Control timeout in milliseconds. The type is interface{} with range:
    // 1..604800000. Units are millisecond. The default value is 5000.
    Timeout interface{}

    // Probe interval in seconds. The type is interface{} with range: 1..604800.
    // Units are second. The default value is 60.
    Frequency interface{}

    // Enter IPv4 address of the destination device. The type is string.
    DestAddress interface{}

    // Add a tag for this operation. The type is string with length: 1..100.
    Tag interface{}

    // Protocol data size in payload of probe packets.
    DataSize Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_DataSize

    // Probe packet stream configuration parameters.
    Packet Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_Packet

    // Loose source routing path (up to 8 intermediate nodes).
    LsrPath Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_LsrPath
}

func (icmpPathJitter *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter) GetEntityData() *types.CommonEntityData {
    icmpPathJitter.EntityData.YFilter = icmpPathJitter.YFilter
    icmpPathJitter.EntityData.YangName = "icmp-path-jitter"
    icmpPathJitter.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitter.EntityData.ParentYangName = "operation-type"
    icmpPathJitter.EntityData.SegmentPath = "icmp-path-jitter"
    icmpPathJitter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitter.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitter.EntityData.Children["data-size"] = types.YChild{"DataSize", &icmpPathJitter.DataSize}
    icmpPathJitter.EntityData.Children["packet"] = types.YChild{"Packet", &icmpPathJitter.Packet}
    icmpPathJitter.EntityData.Children["lsr-path"] = types.YChild{"LsrPath", &icmpPathJitter.LsrPath}
    icmpPathJitter.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitter.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitter.SourceAddress}
    icmpPathJitter.EntityData.Leafs["tos"] = types.YLeaf{"Tos", icmpPathJitter.Tos}
    icmpPathJitter.EntityData.Leafs["create"] = types.YLeaf{"Create", icmpPathJitter.Create}
    icmpPathJitter.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", icmpPathJitter.Vrf}
    icmpPathJitter.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", icmpPathJitter.Timeout}
    icmpPathJitter.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", icmpPathJitter.Frequency}
    icmpPathJitter.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitter.DestAddress}
    icmpPathJitter.EntityData.Leafs["tag"] = types.YLeaf{"Tag", icmpPathJitter.Tag}
    return &(icmpPathJitter.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_DataSize
// Protocol data size in payload of probe
// packets
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_DataSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload size in request probe packet. The type is interface{} with range:
    // 0..16384. Units are byte. The default value is 36.
    Request interface{}
}

func (dataSize *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_DataSize) GetEntityData() *types.CommonEntityData {
    dataSize.EntityData.YFilter = dataSize.YFilter
    dataSize.EntityData.YangName = "data-size"
    dataSize.EntityData.BundleName = "cisco_ios_xr"
    dataSize.EntityData.ParentYangName = "icmp-path-jitter"
    dataSize.EntityData.SegmentPath = "data-size"
    dataSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataSize.EntityData.Children = make(map[string]types.YChild)
    dataSize.EntityData.Leafs = make(map[string]types.YLeaf)
    dataSize.EntityData.Leafs["request"] = types.YLeaf{"Request", dataSize.Request}
    return &(dataSize.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_Packet
// Probe packet stream configuration
// parameters
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_Packet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets to be transmitted during a probe. The type is interface{}
    // with range: 1..100. The default value is 10.
    Count interface{}

    // Packet interval in milliseconds. The type is interface{} with range:
    // 1..60000. Units are millisecond. The default value is 20.
    Interval interface{}
}

func (packet *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_Packet) GetEntityData() *types.CommonEntityData {
    packet.EntityData.YFilter = packet.YFilter
    packet.EntityData.YangName = "packet"
    packet.EntityData.BundleName = "cisco_ios_xr"
    packet.EntityData.ParentYangName = "icmp-path-jitter"
    packet.EntityData.SegmentPath = "packet"
    packet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packet.EntityData.Children = make(map[string]types.YChild)
    packet.EntityData.Leafs = make(map[string]types.YLeaf)
    packet.EntityData.Leafs["count"] = types.YLeaf{"Count", packet.Count}
    packet.EntityData.Leafs["interval"] = types.YLeaf{"Interval", packet.Interval}
    return &(packet.EntityData)
}

// Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_LsrPath
// Loose source routing path (up to 8 intermediate
// nodes)
// This type is a presence type.
type Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_LsrPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Node1 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node2 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node3 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node4 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node5 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node6 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node7 interface{}

    // IPv4 address of the intermediate node. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Node8 interface{}
}

func (lsrPath *Ipsla_Operation_Definitions_Definition_OperationType_IcmpPathJitter_LsrPath) GetEntityData() *types.CommonEntityData {
    lsrPath.EntityData.YFilter = lsrPath.YFilter
    lsrPath.EntityData.YangName = "lsr-path"
    lsrPath.EntityData.BundleName = "cisco_ios_xr"
    lsrPath.EntityData.ParentYangName = "icmp-path-jitter"
    lsrPath.EntityData.SegmentPath = "lsr-path"
    lsrPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsrPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsrPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsrPath.EntityData.Children = make(map[string]types.YChild)
    lsrPath.EntityData.Leafs = make(map[string]types.YLeaf)
    lsrPath.EntityData.Leafs["node1"] = types.YLeaf{"Node1", lsrPath.Node1}
    lsrPath.EntityData.Leafs["node2"] = types.YLeaf{"Node2", lsrPath.Node2}
    lsrPath.EntityData.Leafs["node3"] = types.YLeaf{"Node3", lsrPath.Node3}
    lsrPath.EntityData.Leafs["node4"] = types.YLeaf{"Node4", lsrPath.Node4}
    lsrPath.EntityData.Leafs["node5"] = types.YLeaf{"Node5", lsrPath.Node5}
    lsrPath.EntityData.Leafs["node6"] = types.YLeaf{"Node6", lsrPath.Node6}
    lsrPath.EntityData.Leafs["node7"] = types.YLeaf{"Node7", lsrPath.Node7}
    lsrPath.EntityData.Leafs["node8"] = types.YLeaf{"Node8", lsrPath.Node8}
    return &(lsrPath.EntityData)
}

// Ipsla_Responder
// Responder configuration
type Ipsla_Responder struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Starts the responder process. The type is interface{}.
    Enable interface{}

    // Configure IPSLA Responder port type.
    Type_ Ipsla_Responder_Type
}

func (responder *Ipsla_Responder) GetEntityData() *types.CommonEntityData {
    responder.EntityData.YFilter = responder.YFilter
    responder.EntityData.YangName = "responder"
    responder.EntityData.BundleName = "cisco_ios_xr"
    responder.EntityData.ParentYangName = "ipsla"
    responder.EntityData.SegmentPath = "responder"
    responder.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    responder.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    responder.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    responder.EntityData.Children = make(map[string]types.YChild)
    responder.EntityData.Children["type"] = types.YChild{"Type_", &responder.Type_}
    responder.EntityData.Leafs = make(map[string]types.YLeaf)
    responder.EntityData.Leafs["enable"] = types.YLeaf{"Enable", responder.Enable}
    return &(responder.EntityData)
}

// Ipsla_Responder_Type
// Configure IPSLA Responder port type
type Ipsla_Responder_Type struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IPSLA Responder UDP address and port.
    Udp Ipsla_Responder_Type_Udp
}

func (self *Ipsla_Responder_Type) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "type"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "responder"
    self.EntityData.SegmentPath = "type"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["udp"] = types.YChild{"Udp", &self.Udp}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Ipsla_Responder_Type_Udp
// Configure IPSLA Responder UDP address and port
type Ipsla_Responder_Type_Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IP address.
    Addresses Ipsla_Responder_Type_Udp_Addresses
}

func (udp *Ipsla_Responder_Type_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xr"
    udp.EntityData.ParentYangName = "type"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udp.EntityData.Children = make(map[string]types.YChild)
    udp.EntityData.Children["addresses"] = types.YChild{"Addresses", &udp.Addresses}
    udp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(udp.EntityData)
}

// Ipsla_Responder_Type_Udp_Addresses
// Configure IP address
type Ipsla_Responder_Type_Udp_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IP address for the permanent port. The type is slice of
    // Ipsla_Responder_Type_Udp_Addresses_Address.
    Address []Ipsla_Responder_Type_Udp_Addresses_Address
}

func (addresses *Ipsla_Responder_Type_Udp_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "udp"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = make(map[string]types.YChild)
    addresses.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range addresses.Address {
        addresses.EntityData.Children[types.GetSegmentPath(&addresses.Address[i])] = types.YChild{"Address", &addresses.Address[i]}
    }
    addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addresses.EntityData)
}

// Ipsla_Responder_Type_Udp_Addresses_Address
// Configure IP address for the permanent port
type Ipsla_Responder_Type_Udp_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the Responder. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddress interface{}

    // Configure port.
    Ports Ipsla_Responder_Type_Udp_Addresses_Address_Ports
}

func (address *Ipsla_Responder_Type_Udp_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + "[local-address='" + fmt.Sprintf("%v", address.LocalAddress) + "']"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["ports"] = types.YChild{"Ports", &address.Ports}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", address.LocalAddress}
    return &(address.EntityData)
}

// Ipsla_Responder_Type_Udp_Addresses_Address_Ports
// Configure port
type Ipsla_Responder_Type_Udp_Addresses_Address_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure port number for the permanent port. The type is slice of
    // Ipsla_Responder_Type_Udp_Addresses_Address_Ports_Port.
    Port []Ipsla_Responder_Type_Udp_Addresses_Address_Ports_Port
}

func (ports *Ipsla_Responder_Type_Udp_Addresses_Address_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "address"
    ports.EntityData.SegmentPath = "ports"
    ports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ports.EntityData.Children = make(map[string]types.YChild)
    ports.EntityData.Children["port"] = types.YChild{"Port", nil}
    for i := range ports.Port {
        ports.EntityData.Children[types.GetSegmentPath(&ports.Port[i])] = types.YChild{"Port", &ports.Port[i]}
    }
    ports.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ports.EntityData)
}

// Ipsla_Responder_Type_Udp_Addresses_Address_Ports_Port
// Configure port number for the permanent
// port
type Ipsla_Responder_Type_Udp_Addresses_Address_Ports_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port number to be enabled. The type is interface{}
    // with range: 0..65535.
    Port interface{}
}

func (port *Ipsla_Responder_Type_Udp_Addresses_Address_Ports_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "ports"
    port.EntityData.SegmentPath = "port" + "[port='" + fmt.Sprintf("%v", port.Port) + "']"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    port.EntityData.Leafs["port"] = types.YLeaf{"Port", port.Port}
    return &(port.EntityData)
}

// Ipsla_MplsDiscovery
// Provider Edge(PE) discovery configuration
type Ipsla_MplsDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Layer 3 VPN PE discovery configuration.
    Vpn Ipsla_MplsDiscovery_Vpn
}

func (mplsDiscovery *Ipsla_MplsDiscovery) GetEntityData() *types.CommonEntityData {
    mplsDiscovery.EntityData.YFilter = mplsDiscovery.YFilter
    mplsDiscovery.EntityData.YangName = "mpls-discovery"
    mplsDiscovery.EntityData.BundleName = "cisco_ios_xr"
    mplsDiscovery.EntityData.ParentYangName = "ipsla"
    mplsDiscovery.EntityData.SegmentPath = "mpls-discovery"
    mplsDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsDiscovery.EntityData.Children = make(map[string]types.YChild)
    mplsDiscovery.EntityData.Children["vpn"] = types.YChild{"Vpn", &mplsDiscovery.Vpn}
    mplsDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsDiscovery.EntityData)
}

// Ipsla_MplsDiscovery_Vpn
// Layer 3 VPN PE discovery configuration
type Ipsla_MplsDiscovery_Vpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a discovery refresh interval. The type is interface{} with range:
    // 30..70560. Units are minute. The default value is 300.
    Interval interface{}
}

func (vpn *Ipsla_MplsDiscovery_Vpn) GetEntityData() *types.CommonEntityData {
    vpn.EntityData.YFilter = vpn.YFilter
    vpn.EntityData.YangName = "vpn"
    vpn.EntityData.BundleName = "cisco_ios_xr"
    vpn.EntityData.ParentYangName = "mpls-discovery"
    vpn.EntityData.SegmentPath = "vpn"
    vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpn.EntityData.Children = make(map[string]types.YChild)
    vpn.EntityData.Leafs = make(map[string]types.YLeaf)
    vpn.EntityData.Leafs["interval"] = types.YLeaf{"Interval", vpn.Interval}
    return &(vpn.EntityData)
}

