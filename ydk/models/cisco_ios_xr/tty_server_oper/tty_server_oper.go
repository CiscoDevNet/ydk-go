// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-server package operational data.
// 
// This module contains definitions
// for the following management objects:
//   tty: TTY Line Configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package tty_server_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tty_server_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tty-server-oper tty}", reflect.TypeOf(Tty{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tty-server-oper:tty", reflect.TypeOf(Tty{}))
}

// LineState represents Line state
type LineState string

const (
    // Line not connected
    LineState_none LineState = "none"

    // Line registered
    LineState_registered LineState = "registered"

    // Line active and in use
    LineState_in_use LineState = "in-use"
)

// SessionOperation represents Session operation
type SessionOperation string

const (
    // No sessions on the line
    SessionOperation_none SessionOperation = "none"

    // Session getting set up
    SessionOperation_setup SessionOperation = "setup"

    // Session active with a shell
    SessionOperation_shell SessionOperation = "shell"

    // Session in transitioning phase
    SessionOperation_transitioning SessionOperation = "transitioning"

    // Session ready to receive packets
    SessionOperation_packet SessionOperation = "packet"
)

// Tty
// TTY Line Configuration
type Tty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Nodes for console.
    ConsoleNodes Tty_ConsoleNodes

    // List of VTY lines.
    VtyLines Tty_VtyLines

    // List of Nodes attached with an auxiliary line.
    AuxiliaryNodes Tty_AuxiliaryNodes
}

func (tty *Tty) GetEntityData() *types.CommonEntityData {
    tty.EntityData.YFilter = tty.YFilter
    tty.EntityData.YangName = "tty"
    tty.EntityData.BundleName = "cisco_ios_xr"
    tty.EntityData.ParentYangName = "Cisco-IOS-XR-tty-server-oper"
    tty.EntityData.SegmentPath = "Cisco-IOS-XR-tty-server-oper:tty"
    tty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tty.EntityData.Children = types.NewOrderedMap()
    tty.EntityData.Children.Append("console-nodes", types.YChild{"ConsoleNodes", &tty.ConsoleNodes})
    tty.EntityData.Children.Append("vty-lines", types.YChild{"VtyLines", &tty.VtyLines})
    tty.EntityData.Children.Append("auxiliary-nodes", types.YChild{"AuxiliaryNodes", &tty.AuxiliaryNodes})
    tty.EntityData.Leafs = types.NewOrderedMap()

    tty.EntityData.YListKeys = []string {}

    return &(tty.EntityData)
}

// Tty_ConsoleNodes
// List of Nodes for console
type Tty_ConsoleNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Console line configuration on a node. The type is slice of
    // Tty_ConsoleNodes_ConsoleNode.
    ConsoleNode []*Tty_ConsoleNodes_ConsoleNode
}

func (consoleNodes *Tty_ConsoleNodes) GetEntityData() *types.CommonEntityData {
    consoleNodes.EntityData.YFilter = consoleNodes.YFilter
    consoleNodes.EntityData.YangName = "console-nodes"
    consoleNodes.EntityData.BundleName = "cisco_ios_xr"
    consoleNodes.EntityData.ParentYangName = "tty"
    consoleNodes.EntityData.SegmentPath = "console-nodes"
    consoleNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    consoleNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    consoleNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    consoleNodes.EntityData.Children = types.NewOrderedMap()
    consoleNodes.EntityData.Children.Append("console-node", types.YChild{"ConsoleNode", nil})
    for i := range consoleNodes.ConsoleNode {
        consoleNodes.EntityData.Children.Append(types.GetSegmentPath(consoleNodes.ConsoleNode[i]), types.YChild{"ConsoleNode", consoleNodes.ConsoleNode[i]})
    }
    consoleNodes.EntityData.Leafs = types.NewOrderedMap()

    consoleNodes.EntityData.YListKeys = []string {}

    return &(consoleNodes.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode
// Console line configuration on a node
type Tty_ConsoleNodes_ConsoleNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Id interface{}

    // Console line.
    ConsoleLine Tty_ConsoleNodes_ConsoleNode_ConsoleLine
}

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetEntityData() *types.CommonEntityData {
    consoleNode.EntityData.YFilter = consoleNode.YFilter
    consoleNode.EntityData.YangName = "console-node"
    consoleNode.EntityData.BundleName = "cisco_ios_xr"
    consoleNode.EntityData.ParentYangName = "console-nodes"
    consoleNode.EntityData.SegmentPath = "console-node" + types.AddKeyToken(consoleNode.Id, "id")
    consoleNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    consoleNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    consoleNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    consoleNode.EntityData.Children = types.NewOrderedMap()
    consoleNode.EntityData.Children.Append("console-line", types.YChild{"ConsoleLine", &consoleNode.ConsoleLine})
    consoleNode.EntityData.Leafs = types.NewOrderedMap()
    consoleNode.EntityData.Leafs.Append("id", types.YLeaf{"Id", consoleNode.Id})

    consoleNode.EntityData.YListKeys = []string {"Id"}

    return &(consoleNode.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine
// Console line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of the console line.
    ConsoleStatistics Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics

    // Line state information.
    State Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State

    // Configuration information of the line.
    Configuration Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration
}

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetEntityData() *types.CommonEntityData {
    consoleLine.EntityData.YFilter = consoleLine.YFilter
    consoleLine.EntityData.YangName = "console-line"
    consoleLine.EntityData.BundleName = "cisco_ios_xr"
    consoleLine.EntityData.ParentYangName = "console-node"
    consoleLine.EntityData.SegmentPath = "console-line"
    consoleLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    consoleLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    consoleLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    consoleLine.EntityData.Children = types.NewOrderedMap()
    consoleLine.EntityData.Children.Append("console-statistics", types.YChild{"ConsoleStatistics", &consoleLine.ConsoleStatistics})
    consoleLine.EntityData.Children.Append("state", types.YChild{"State", &consoleLine.State})
    consoleLine.EntityData.Children.Append("configuration", types.YChild{"Configuration", &consoleLine.Configuration})
    consoleLine.EntityData.Leafs = types.NewOrderedMap()

    consoleLine.EntityData.YListKeys = []string {}

    return &(consoleLine.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics
// Statistics of the console line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RS232 statistics of console line.
    Rs232 Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232

    // General statistics of line.
    GeneralStatistics Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics

    // Exec related statistics.
    Exec Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec

    // AAA related statistics.
    Aaa Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa
}

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetEntityData() *types.CommonEntityData {
    consoleStatistics.EntityData.YFilter = consoleStatistics.YFilter
    consoleStatistics.EntityData.YangName = "console-statistics"
    consoleStatistics.EntityData.BundleName = "cisco_ios_xr"
    consoleStatistics.EntityData.ParentYangName = "console-line"
    consoleStatistics.EntityData.SegmentPath = "console-statistics"
    consoleStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    consoleStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    consoleStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    consoleStatistics.EntityData.Children = types.NewOrderedMap()
    consoleStatistics.EntityData.Children.Append("rs232", types.YChild{"Rs232", &consoleStatistics.Rs232})
    consoleStatistics.EntityData.Children.Append("general-statistics", types.YChild{"GeneralStatistics", &consoleStatistics.GeneralStatistics})
    consoleStatistics.EntityData.Children.Append("exec", types.YChild{"Exec", &consoleStatistics.Exec})
    consoleStatistics.EntityData.Children.Append("aaa", types.YChild{"Aaa", &consoleStatistics.Aaa})
    consoleStatistics.EntityData.Leafs = types.NewOrderedMap()

    consoleStatistics.EntityData.YListKeys = []string {}

    return &(consoleStatistics.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232
// RS232 statistics of console line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of databits. The type is interface{} with range: 0..4294967295.
    // Units are bit.
    DataBits interface{}

    // Exec disabled on TTY. The type is bool.
    ExecDisabled interface{}

    // Hardware flow control status. The type is interface{} with range:
    // 0..4294967295.
    HardwareFlowControlStatus interface{}

    // Parity status. The type is interface{} with range: 0..4294967295.
    ParityStatus interface{}

    // Inbound/Outbound baud rate in bps. The type is interface{} with range:
    // 0..4294967295. Units are bit/s.
    BaudRate interface{}

    // Number of stopbits. The type is interface{} with range: 0..4294967295.
    // Units are bit.
    StopBits interface{}

    // Overrun error count. The type is interface{} with range: 0..4294967295.
    OverrunErrorCount interface{}

    // Framing error count. The type is interface{} with range: 0..4294967295.
    FramingErrorCount interface{}

    // Parity error count. The type is interface{} with range: 0..4294967295.
    ParityErrorCount interface{}
}

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetEntityData() *types.CommonEntityData {
    rs232.EntityData.YFilter = rs232.YFilter
    rs232.EntityData.YangName = "rs232"
    rs232.EntityData.BundleName = "cisco_ios_xr"
    rs232.EntityData.ParentYangName = "console-statistics"
    rs232.EntityData.SegmentPath = "rs232"
    rs232.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rs232.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rs232.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rs232.EntityData.Children = types.NewOrderedMap()
    rs232.EntityData.Leafs = types.NewOrderedMap()
    rs232.EntityData.Leafs.Append("data-bits", types.YLeaf{"DataBits", rs232.DataBits})
    rs232.EntityData.Leafs.Append("exec-disabled", types.YLeaf{"ExecDisabled", rs232.ExecDisabled})
    rs232.EntityData.Leafs.Append("hardware-flow-control-status", types.YLeaf{"HardwareFlowControlStatus", rs232.HardwareFlowControlStatus})
    rs232.EntityData.Leafs.Append("parity-status", types.YLeaf{"ParityStatus", rs232.ParityStatus})
    rs232.EntityData.Leafs.Append("baud-rate", types.YLeaf{"BaudRate", rs232.BaudRate})
    rs232.EntityData.Leafs.Append("stop-bits", types.YLeaf{"StopBits", rs232.StopBits})
    rs232.EntityData.Leafs.Append("overrun-error-count", types.YLeaf{"OverrunErrorCount", rs232.OverrunErrorCount})
    rs232.EntityData.Leafs.Append("framing-error-count", types.YLeaf{"FramingErrorCount", rs232.FramingErrorCount})
    rs232.EntityData.Leafs.Append("parity-error-count", types.YLeaf{"ParityErrorCount", rs232.ParityErrorCount})

    rs232.EntityData.YListKeys = []string {}

    return &(rs232.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics
// General statistics of line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Terminal length. The type is interface{} with range: 0..4294967295.
    TerminalLength interface{}

    // Line width. The type is interface{} with range: 0..4294967295.
    TerminalWidth interface{}

    // Usable as async interface. The type is bool.
    AsyncInterface interface{}

    // Software flow control start char. The type is interface{} with range:
    // -128..127.
    FlowControlStartCharacter interface{}

    // Software flow control stop char. The type is interface{} with range:
    // -128..127.
    FlowControlStopCharacter interface{}

    // DNS resolution enabled. The type is bool.
    DomainLookupEnabled interface{}

    // MOTD banner enabled. The type is bool.
    MotdBannerEnabled interface{}

    // TTY private flag. The type is bool.
    PrivateFlag interface{}

    // Terminal type. The type is string.
    TerminalType interface{}

    // Absolute timeout period. The type is interface{} with range: 0..4294967295.
    AbsoluteTimeout interface{}

    // TTY idle time. The type is interface{} with range: 0..4294967295.
    IdleTime interface{}
}

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetEntityData() *types.CommonEntityData {
    generalStatistics.EntityData.YFilter = generalStatistics.YFilter
    generalStatistics.EntityData.YangName = "general-statistics"
    generalStatistics.EntityData.BundleName = "cisco_ios_xr"
    generalStatistics.EntityData.ParentYangName = "console-statistics"
    generalStatistics.EntityData.SegmentPath = "general-statistics"
    generalStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generalStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generalStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generalStatistics.EntityData.Children = types.NewOrderedMap()
    generalStatistics.EntityData.Leafs = types.NewOrderedMap()
    generalStatistics.EntityData.Leafs.Append("terminal-length", types.YLeaf{"TerminalLength", generalStatistics.TerminalLength})
    generalStatistics.EntityData.Leafs.Append("terminal-width", types.YLeaf{"TerminalWidth", generalStatistics.TerminalWidth})
    generalStatistics.EntityData.Leafs.Append("async-interface", types.YLeaf{"AsyncInterface", generalStatistics.AsyncInterface})
    generalStatistics.EntityData.Leafs.Append("flow-control-start-character", types.YLeaf{"FlowControlStartCharacter", generalStatistics.FlowControlStartCharacter})
    generalStatistics.EntityData.Leafs.Append("flow-control-stop-character", types.YLeaf{"FlowControlStopCharacter", generalStatistics.FlowControlStopCharacter})
    generalStatistics.EntityData.Leafs.Append("domain-lookup-enabled", types.YLeaf{"DomainLookupEnabled", generalStatistics.DomainLookupEnabled})
    generalStatistics.EntityData.Leafs.Append("motd-banner-enabled", types.YLeaf{"MotdBannerEnabled", generalStatistics.MotdBannerEnabled})
    generalStatistics.EntityData.Leafs.Append("private-flag", types.YLeaf{"PrivateFlag", generalStatistics.PrivateFlag})
    generalStatistics.EntityData.Leafs.Append("terminal-type", types.YLeaf{"TerminalType", generalStatistics.TerminalType})
    generalStatistics.EntityData.Leafs.Append("absolute-timeout", types.YLeaf{"AbsoluteTimeout", generalStatistics.AbsoluteTimeout})
    generalStatistics.EntityData.Leafs.Append("idle-time", types.YLeaf{"IdleTime", generalStatistics.IdleTime})

    generalStatistics.EntityData.YListKeys = []string {}

    return &(generalStatistics.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec
// Exec related statistics
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies whether timestamp is enabled or not. The type is bool.
    TimeStampEnabled interface{}
}

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetEntityData() *types.CommonEntityData {
    exec.EntityData.YFilter = exec.YFilter
    exec.EntityData.YangName = "exec"
    exec.EntityData.BundleName = "cisco_ios_xr"
    exec.EntityData.ParentYangName = "console-statistics"
    exec.EntityData.SegmentPath = "exec"
    exec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exec.EntityData.Children = types.NewOrderedMap()
    exec.EntityData.Leafs = types.NewOrderedMap()
    exec.EntityData.Leafs.Append("time-stamp-enabled", types.YLeaf{"TimeStampEnabled", exec.TimeStampEnabled})

    exec.EntityData.YListKeys = []string {}

    return &(exec.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa
// AAA related statistics
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The authenticated username. The type is string.
    UserName interface{}
}

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "console-statistics"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Leafs = types.NewOrderedMap()
    aaa.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", aaa.UserName})

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State
// Line state information
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information related to template applied to the line.
    Template Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template

    // General information.
    General Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General
}

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "console-line"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("template", types.YChild{"Template", &state.Template})
    state.EntityData.Children.Append("general", types.YChild{"General", &state.General})
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template
// Information related to template applied to the
// line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the template. The type is string.
    Name interface{}
}

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "state"
    template.EntityData.SegmentPath = "template"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("name", types.YLeaf{"Name", template.Name})

    template.EntityData.YListKeys = []string {}

    return &(template.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General
// General information
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // application running of on the tty line. The type is SessionOperation.
    Operation interface{}

    // State of the line. The type is LineState.
    GeneralState interface{}
}

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetEntityData() *types.CommonEntityData {
    general.EntityData.YFilter = general.YFilter
    general.EntityData.YangName = "general"
    general.EntityData.BundleName = "cisco_ios_xr"
    general.EntityData.ParentYangName = "state"
    general.EntityData.SegmentPath = "general"
    general.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    general.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    general.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    general.EntityData.Children = types.NewOrderedMap()
    general.EntityData.Leafs = types.NewOrderedMap()
    general.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", general.Operation})
    general.EntityData.Leafs.Append("general-state", types.YLeaf{"GeneralState", general.GeneralState})

    general.EntityData.YListKeys = []string {}

    return &(general.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration
// Configuration information of the line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Conection configuration information.
    ConnectionConfiguration Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration
}

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "console-line"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = types.NewOrderedMap()
    configuration.EntityData.Children.Append("connection-configuration", types.YChild{"ConnectionConfiguration", &configuration.ConnectionConfiguration})
    configuration.EntityData.Leafs = types.NewOrderedMap()

    configuration.EntityData.YListKeys = []string {}

    return &(configuration.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration
// Conection configuration information
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL for outbound traffic. The type is string.
    AclOut interface{}

    // ACL for inbound traffic. The type is string.
    AclIn interface{}

    // Protocols to use when connecting to the terminal server.
    TransportInput Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput
}

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetEntityData() *types.CommonEntityData {
    connectionConfiguration.EntityData.YFilter = connectionConfiguration.YFilter
    connectionConfiguration.EntityData.YangName = "connection-configuration"
    connectionConfiguration.EntityData.BundleName = "cisco_ios_xr"
    connectionConfiguration.EntityData.ParentYangName = "configuration"
    connectionConfiguration.EntityData.SegmentPath = "connection-configuration"
    connectionConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectionConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectionConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectionConfiguration.EntityData.Children = types.NewOrderedMap()
    connectionConfiguration.EntityData.Children.Append("transport-input", types.YChild{"TransportInput", &connectionConfiguration.TransportInput})
    connectionConfiguration.EntityData.Leafs = types.NewOrderedMap()
    connectionConfiguration.EntityData.Leafs.Append("acl-out", types.YLeaf{"AclOut", connectionConfiguration.AclOut})
    connectionConfiguration.EntityData.Leafs.Append("acl-in", types.YLeaf{"AclIn", connectionConfiguration.AclIn})

    connectionConfiguration.EntityData.YListKeys = []string {}

    return &(connectionConfiguration.EntityData)
}

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: 0..4294967295.
    None interface{}
}

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetEntityData() *types.CommonEntityData {
    transportInput.EntityData.YFilter = transportInput.YFilter
    transportInput.EntityData.YangName = "transport-input"
    transportInput.EntityData.BundleName = "cisco_ios_xr"
    transportInput.EntityData.ParentYangName = "connection-configuration"
    transportInput.EntityData.SegmentPath = "transport-input"
    transportInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportInput.EntityData.Children = types.NewOrderedMap()
    transportInput.EntityData.Leafs = types.NewOrderedMap()
    transportInput.EntityData.Leafs.Append("select", types.YLeaf{"Select", transportInput.Select})
    transportInput.EntityData.Leafs.Append("protocol1", types.YLeaf{"Protocol1", transportInput.Protocol1})
    transportInput.EntityData.Leafs.Append("protocol2", types.YLeaf{"Protocol2", transportInput.Protocol2})
    transportInput.EntityData.Leafs.Append("none", types.YLeaf{"None", transportInput.None})

    transportInput.EntityData.YListKeys = []string {}

    return &(transportInput.EntityData)
}

// Tty_VtyLines
// List of VTY lines
type Tty_VtyLines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VTY Line. The type is slice of Tty_VtyLines_VtyLine.
    VtyLine []*Tty_VtyLines_VtyLine
}

func (vtyLines *Tty_VtyLines) GetEntityData() *types.CommonEntityData {
    vtyLines.EntityData.YFilter = vtyLines.YFilter
    vtyLines.EntityData.YangName = "vty-lines"
    vtyLines.EntityData.BundleName = "cisco_ios_xr"
    vtyLines.EntityData.ParentYangName = "tty"
    vtyLines.EntityData.SegmentPath = "vty-lines"
    vtyLines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vtyLines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vtyLines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vtyLines.EntityData.Children = types.NewOrderedMap()
    vtyLines.EntityData.Children.Append("vty-line", types.YChild{"VtyLine", nil})
    for i := range vtyLines.VtyLine {
        vtyLines.EntityData.Children.Append(types.GetSegmentPath(vtyLines.VtyLine[i]), types.YChild{"VtyLine", vtyLines.VtyLine[i]})
    }
    vtyLines.EntityData.Leafs = types.NewOrderedMap()

    vtyLines.EntityData.YListKeys = []string {}

    return &(vtyLines.EntityData)
}

// Tty_VtyLines_VtyLine
// VTY Line
type Tty_VtyLines_VtyLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VTY Line number. The type is interface{} with
    // range: 0..4294967295.
    LineNumber interface{}

    // Statistics of the VTY line.
    VtyStatistics Tty_VtyLines_VtyLine_VtyStatistics

    // Line state information.
    State Tty_VtyLines_VtyLine_State

    // Configuration information of the line.
    Configuration Tty_VtyLines_VtyLine_Configuration

    // Outgoing sessions.
    Sessions Tty_VtyLines_VtyLine_Sessions
}

func (vtyLine *Tty_VtyLines_VtyLine) GetEntityData() *types.CommonEntityData {
    vtyLine.EntityData.YFilter = vtyLine.YFilter
    vtyLine.EntityData.YangName = "vty-line"
    vtyLine.EntityData.BundleName = "cisco_ios_xr"
    vtyLine.EntityData.ParentYangName = "vty-lines"
    vtyLine.EntityData.SegmentPath = "vty-line" + types.AddKeyToken(vtyLine.LineNumber, "line-number")
    vtyLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vtyLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vtyLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vtyLine.EntityData.Children = types.NewOrderedMap()
    vtyLine.EntityData.Children.Append("vty-statistics", types.YChild{"VtyStatistics", &vtyLine.VtyStatistics})
    vtyLine.EntityData.Children.Append("state", types.YChild{"State", &vtyLine.State})
    vtyLine.EntityData.Children.Append("configuration", types.YChild{"Configuration", &vtyLine.Configuration})
    vtyLine.EntityData.Children.Append("Cisco-IOS-XR-tty-management-oper:sessions", types.YChild{"Sessions", &vtyLine.Sessions})
    vtyLine.EntityData.Leafs = types.NewOrderedMap()
    vtyLine.EntityData.Leafs.Append("line-number", types.YLeaf{"LineNumber", vtyLine.LineNumber})

    vtyLine.EntityData.YListKeys = []string {"LineNumber"}

    return &(vtyLine.EntityData)
}

// Tty_VtyLines_VtyLine_VtyStatistics
// Statistics of the VTY line
type Tty_VtyLines_VtyLine_VtyStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Connection related statistics.
    Connection Tty_VtyLines_VtyLine_VtyStatistics_Connection

    // General statistics of line.
    GeneralStatistics Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics

    // Exec related statistics.
    Exec Tty_VtyLines_VtyLine_VtyStatistics_Exec

    // AAA related statistics.
    Aaa Tty_VtyLines_VtyLine_VtyStatistics_Aaa
}

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetEntityData() *types.CommonEntityData {
    vtyStatistics.EntityData.YFilter = vtyStatistics.YFilter
    vtyStatistics.EntityData.YangName = "vty-statistics"
    vtyStatistics.EntityData.BundleName = "cisco_ios_xr"
    vtyStatistics.EntityData.ParentYangName = "vty-line"
    vtyStatistics.EntityData.SegmentPath = "vty-statistics"
    vtyStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vtyStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vtyStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vtyStatistics.EntityData.Children = types.NewOrderedMap()
    vtyStatistics.EntityData.Children.Append("connection", types.YChild{"Connection", &vtyStatistics.Connection})
    vtyStatistics.EntityData.Children.Append("general-statistics", types.YChild{"GeneralStatistics", &vtyStatistics.GeneralStatistics})
    vtyStatistics.EntityData.Children.Append("exec", types.YChild{"Exec", &vtyStatistics.Exec})
    vtyStatistics.EntityData.Children.Append("aaa", types.YChild{"Aaa", &vtyStatistics.Aaa})
    vtyStatistics.EntityData.Leafs = types.NewOrderedMap()

    vtyStatistics.EntityData.YListKeys = []string {}

    return &(vtyStatistics.EntityData)
}

// Tty_VtyLines_VtyLine_VtyStatistics_Connection
// Connection related statistics
type Tty_VtyLines_VtyLine_VtyStatistics_Connection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Incoming host address(max). The type is string with length: 0..46.
    IncomingHostAddress interface{}

    // Incoming host address family. The type is interface{} with range:
    // 0..4294967295.
    HostAddressFamily interface{}

    // Input transport. The type is interface{} with range: 0..4294967295.
    Service interface{}
}

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetEntityData() *types.CommonEntityData {
    connection.EntityData.YFilter = connection.YFilter
    connection.EntityData.YangName = "connection"
    connection.EntityData.BundleName = "cisco_ios_xr"
    connection.EntityData.ParentYangName = "vty-statistics"
    connection.EntityData.SegmentPath = "connection"
    connection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connection.EntityData.Children = types.NewOrderedMap()
    connection.EntityData.Leafs = types.NewOrderedMap()
    connection.EntityData.Leafs.Append("incoming-host-address", types.YLeaf{"IncomingHostAddress", connection.IncomingHostAddress})
    connection.EntityData.Leafs.Append("host-address-family", types.YLeaf{"HostAddressFamily", connection.HostAddressFamily})
    connection.EntityData.Leafs.Append("service", types.YLeaf{"Service", connection.Service})

    connection.EntityData.YListKeys = []string {}

    return &(connection.EntityData)
}

// Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics
// General statistics of line
type Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Terminal length. The type is interface{} with range: 0..4294967295.
    TerminalLength interface{}

    // Line width. The type is interface{} with range: 0..4294967295.
    TerminalWidth interface{}

    // Usable as async interface. The type is bool.
    AsyncInterface interface{}

    // Software flow control start char. The type is interface{} with range:
    // -128..127.
    FlowControlStartCharacter interface{}

    // Software flow control stop char. The type is interface{} with range:
    // -128..127.
    FlowControlStopCharacter interface{}

    // DNS resolution enabled. The type is bool.
    DomainLookupEnabled interface{}

    // MOTD banner enabled. The type is bool.
    MotdBannerEnabled interface{}

    // TTY private flag. The type is bool.
    PrivateFlag interface{}

    // Terminal type. The type is string.
    TerminalType interface{}

    // Absolute timeout period. The type is interface{} with range: 0..4294967295.
    AbsoluteTimeout interface{}

    // TTY idle time. The type is interface{} with range: 0..4294967295.
    IdleTime interface{}
}

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetEntityData() *types.CommonEntityData {
    generalStatistics.EntityData.YFilter = generalStatistics.YFilter
    generalStatistics.EntityData.YangName = "general-statistics"
    generalStatistics.EntityData.BundleName = "cisco_ios_xr"
    generalStatistics.EntityData.ParentYangName = "vty-statistics"
    generalStatistics.EntityData.SegmentPath = "general-statistics"
    generalStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generalStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generalStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generalStatistics.EntityData.Children = types.NewOrderedMap()
    generalStatistics.EntityData.Leafs = types.NewOrderedMap()
    generalStatistics.EntityData.Leafs.Append("terminal-length", types.YLeaf{"TerminalLength", generalStatistics.TerminalLength})
    generalStatistics.EntityData.Leafs.Append("terminal-width", types.YLeaf{"TerminalWidth", generalStatistics.TerminalWidth})
    generalStatistics.EntityData.Leafs.Append("async-interface", types.YLeaf{"AsyncInterface", generalStatistics.AsyncInterface})
    generalStatistics.EntityData.Leafs.Append("flow-control-start-character", types.YLeaf{"FlowControlStartCharacter", generalStatistics.FlowControlStartCharacter})
    generalStatistics.EntityData.Leafs.Append("flow-control-stop-character", types.YLeaf{"FlowControlStopCharacter", generalStatistics.FlowControlStopCharacter})
    generalStatistics.EntityData.Leafs.Append("domain-lookup-enabled", types.YLeaf{"DomainLookupEnabled", generalStatistics.DomainLookupEnabled})
    generalStatistics.EntityData.Leafs.Append("motd-banner-enabled", types.YLeaf{"MotdBannerEnabled", generalStatistics.MotdBannerEnabled})
    generalStatistics.EntityData.Leafs.Append("private-flag", types.YLeaf{"PrivateFlag", generalStatistics.PrivateFlag})
    generalStatistics.EntityData.Leafs.Append("terminal-type", types.YLeaf{"TerminalType", generalStatistics.TerminalType})
    generalStatistics.EntityData.Leafs.Append("absolute-timeout", types.YLeaf{"AbsoluteTimeout", generalStatistics.AbsoluteTimeout})
    generalStatistics.EntityData.Leafs.Append("idle-time", types.YLeaf{"IdleTime", generalStatistics.IdleTime})

    generalStatistics.EntityData.YListKeys = []string {}

    return &(generalStatistics.EntityData)
}

// Tty_VtyLines_VtyLine_VtyStatistics_Exec
// Exec related statistics
type Tty_VtyLines_VtyLine_VtyStatistics_Exec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies whether timestamp is enabled or not. The type is bool.
    TimeStampEnabled interface{}
}

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetEntityData() *types.CommonEntityData {
    exec.EntityData.YFilter = exec.YFilter
    exec.EntityData.YangName = "exec"
    exec.EntityData.BundleName = "cisco_ios_xr"
    exec.EntityData.ParentYangName = "vty-statistics"
    exec.EntityData.SegmentPath = "exec"
    exec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exec.EntityData.Children = types.NewOrderedMap()
    exec.EntityData.Leafs = types.NewOrderedMap()
    exec.EntityData.Leafs.Append("time-stamp-enabled", types.YLeaf{"TimeStampEnabled", exec.TimeStampEnabled})

    exec.EntityData.YListKeys = []string {}

    return &(exec.EntityData)
}

// Tty_VtyLines_VtyLine_VtyStatistics_Aaa
// AAA related statistics
type Tty_VtyLines_VtyLine_VtyStatistics_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The authenticated username. The type is string.
    UserName interface{}
}

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "vty-statistics"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Leafs = types.NewOrderedMap()
    aaa.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", aaa.UserName})

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Tty_VtyLines_VtyLine_State
// Line state information
type Tty_VtyLines_VtyLine_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information related to template applied to the line.
    Template Tty_VtyLines_VtyLine_State_Template

    // General information.
    General Tty_VtyLines_VtyLine_State_General
}

func (state *Tty_VtyLines_VtyLine_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "vty-line"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("template", types.YChild{"Template", &state.Template})
    state.EntityData.Children.Append("general", types.YChild{"General", &state.General})
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Tty_VtyLines_VtyLine_State_Template
// Information related to template applied to the
// line
type Tty_VtyLines_VtyLine_State_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the template. The type is string.
    Name interface{}
}

func (template *Tty_VtyLines_VtyLine_State_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "state"
    template.EntityData.SegmentPath = "template"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("name", types.YLeaf{"Name", template.Name})

    template.EntityData.YListKeys = []string {}

    return &(template.EntityData)
}

// Tty_VtyLines_VtyLine_State_General
// General information
type Tty_VtyLines_VtyLine_State_General struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // application running of on the tty line. The type is SessionOperation.
    Operation interface{}

    // State of the line. The type is LineState.
    GeneralState interface{}
}

func (general *Tty_VtyLines_VtyLine_State_General) GetEntityData() *types.CommonEntityData {
    general.EntityData.YFilter = general.YFilter
    general.EntityData.YangName = "general"
    general.EntityData.BundleName = "cisco_ios_xr"
    general.EntityData.ParentYangName = "state"
    general.EntityData.SegmentPath = "general"
    general.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    general.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    general.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    general.EntityData.Children = types.NewOrderedMap()
    general.EntityData.Leafs = types.NewOrderedMap()
    general.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", general.Operation})
    general.EntityData.Leafs.Append("general-state", types.YLeaf{"GeneralState", general.GeneralState})

    general.EntityData.YListKeys = []string {}

    return &(general.EntityData)
}

// Tty_VtyLines_VtyLine_Configuration
// Configuration information of the line
type Tty_VtyLines_VtyLine_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Conection configuration information.
    ConnectionConfiguration Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration
}

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "vty-line"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = types.NewOrderedMap()
    configuration.EntityData.Children.Append("connection-configuration", types.YChild{"ConnectionConfiguration", &configuration.ConnectionConfiguration})
    configuration.EntityData.Leafs = types.NewOrderedMap()

    configuration.EntityData.YListKeys = []string {}

    return &(configuration.EntityData)
}

// Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration
// Conection configuration information
type Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL for outbound traffic. The type is string.
    AclOut interface{}

    // ACL for inbound traffic. The type is string.
    AclIn interface{}

    // Protocols to use when connecting to the terminal server.
    TransportInput Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput
}

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetEntityData() *types.CommonEntityData {
    connectionConfiguration.EntityData.YFilter = connectionConfiguration.YFilter
    connectionConfiguration.EntityData.YangName = "connection-configuration"
    connectionConfiguration.EntityData.BundleName = "cisco_ios_xr"
    connectionConfiguration.EntityData.ParentYangName = "configuration"
    connectionConfiguration.EntityData.SegmentPath = "connection-configuration"
    connectionConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectionConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectionConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectionConfiguration.EntityData.Children = types.NewOrderedMap()
    connectionConfiguration.EntityData.Children.Append("transport-input", types.YChild{"TransportInput", &connectionConfiguration.TransportInput})
    connectionConfiguration.EntityData.Leafs = types.NewOrderedMap()
    connectionConfiguration.EntityData.Leafs.Append("acl-out", types.YLeaf{"AclOut", connectionConfiguration.AclOut})
    connectionConfiguration.EntityData.Leafs.Append("acl-in", types.YLeaf{"AclIn", connectionConfiguration.AclIn})

    connectionConfiguration.EntityData.YListKeys = []string {}

    return &(connectionConfiguration.EntityData)
}

// Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: 0..4294967295.
    None interface{}
}

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetEntityData() *types.CommonEntityData {
    transportInput.EntityData.YFilter = transportInput.YFilter
    transportInput.EntityData.YangName = "transport-input"
    transportInput.EntityData.BundleName = "cisco_ios_xr"
    transportInput.EntityData.ParentYangName = "connection-configuration"
    transportInput.EntityData.SegmentPath = "transport-input"
    transportInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportInput.EntityData.Children = types.NewOrderedMap()
    transportInput.EntityData.Leafs = types.NewOrderedMap()
    transportInput.EntityData.Leafs.Append("select", types.YLeaf{"Select", transportInput.Select})
    transportInput.EntityData.Leafs.Append("protocol1", types.YLeaf{"Protocol1", transportInput.Protocol1})
    transportInput.EntityData.Leafs.Append("protocol2", types.YLeaf{"Protocol2", transportInput.Protocol2})
    transportInput.EntityData.Leafs.Append("none", types.YLeaf{"None", transportInput.None})

    transportInput.EntityData.YListKeys = []string {}

    return &(transportInput.EntityData)
}

// Tty_VtyLines_VtyLine_Sessions
// Outgoing sessions
type Tty_VtyLines_VtyLine_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of outgoing sessions. The type is slice of
    // Tty_VtyLines_VtyLine_Sessions_OutgoingConnection.
    OutgoingConnection []*Tty_VtyLines_VtyLine_Sessions_OutgoingConnection
}

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "vty-line"
    sessions.EntityData.SegmentPath = "Cisco-IOS-XR-tty-management-oper:sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("outgoing-connection", types.YChild{"OutgoingConnection", nil})
    for i := range sessions.OutgoingConnection {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.OutgoingConnection[i]), types.YChild{"OutgoingConnection", sessions.OutgoingConnection[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// Tty_VtyLines_VtyLine_Sessions_OutgoingConnection
// List of outgoing sessions
type Tty_VtyLines_VtyLine_Sessions_OutgoingConnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Connection ID [1-20]. The type is interface{} with range: 0..255.
    ConnectionId interface{}

    // Host name. The type is string.
    HostName interface{}

    // Session transport protocol. The type is TransportService.
    TransportProtocol interface{}

    // True indicates last active session. The type is bool.
    IsLastActiveSession interface{}

    // Elapsed time since session was suspended (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    IdleTime interface{}

    // Host address.
    HostAddress Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress
}

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetEntityData() *types.CommonEntityData {
    outgoingConnection.EntityData.YFilter = outgoingConnection.YFilter
    outgoingConnection.EntityData.YangName = "outgoing-connection"
    outgoingConnection.EntityData.BundleName = "cisco_ios_xr"
    outgoingConnection.EntityData.ParentYangName = "sessions"
    outgoingConnection.EntityData.SegmentPath = "outgoing-connection"
    outgoingConnection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outgoingConnection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outgoingConnection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outgoingConnection.EntityData.Children = types.NewOrderedMap()
    outgoingConnection.EntityData.Children.Append("host-address", types.YChild{"HostAddress", &outgoingConnection.HostAddress})
    outgoingConnection.EntityData.Leafs = types.NewOrderedMap()
    outgoingConnection.EntityData.Leafs.Append("connection-id", types.YLeaf{"ConnectionId", outgoingConnection.ConnectionId})
    outgoingConnection.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", outgoingConnection.HostName})
    outgoingConnection.EntityData.Leafs.Append("transport-protocol", types.YLeaf{"TransportProtocol", outgoingConnection.TransportProtocol})
    outgoingConnection.EntityData.Leafs.Append("is-last-active-session", types.YLeaf{"IsLastActiveSession", outgoingConnection.IsLastActiveSession})
    outgoingConnection.EntityData.Leafs.Append("idle-time", types.YLeaf{"IdleTime", outgoingConnection.IdleTime})

    outgoingConnection.EntityData.YListKeys = []string {}

    return &(outgoingConnection.EntityData)
}

// Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress
// Host address
type Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetEntityData() *types.CommonEntityData {
    hostAddress.EntityData.YFilter = hostAddress.YFilter
    hostAddress.EntityData.YangName = "host-address"
    hostAddress.EntityData.BundleName = "cisco_ios_xr"
    hostAddress.EntityData.ParentYangName = "outgoing-connection"
    hostAddress.EntityData.SegmentPath = "host-address"
    hostAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostAddress.EntityData.Children = types.NewOrderedMap()
    hostAddress.EntityData.Leafs = types.NewOrderedMap()
    hostAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", hostAddress.AfName})
    hostAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", hostAddress.Ipv4Address})
    hostAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", hostAddress.Ipv6Address})

    hostAddress.EntityData.YListKeys = []string {}

    return &(hostAddress.EntityData)
}

// Tty_AuxiliaryNodes
// List of Nodes attached with an auxiliary line
type Tty_AuxiliaryNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Line configuration on a node. The type is slice of
    // Tty_AuxiliaryNodes_AuxiliaryNode.
    AuxiliaryNode []*Tty_AuxiliaryNodes_AuxiliaryNode
}

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetEntityData() *types.CommonEntityData {
    auxiliaryNodes.EntityData.YFilter = auxiliaryNodes.YFilter
    auxiliaryNodes.EntityData.YangName = "auxiliary-nodes"
    auxiliaryNodes.EntityData.BundleName = "cisco_ios_xr"
    auxiliaryNodes.EntityData.ParentYangName = "tty"
    auxiliaryNodes.EntityData.SegmentPath = "auxiliary-nodes"
    auxiliaryNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auxiliaryNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auxiliaryNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auxiliaryNodes.EntityData.Children = types.NewOrderedMap()
    auxiliaryNodes.EntityData.Children.Append("auxiliary-node", types.YChild{"AuxiliaryNode", nil})
    for i := range auxiliaryNodes.AuxiliaryNode {
        auxiliaryNodes.EntityData.Children.Append(types.GetSegmentPath(auxiliaryNodes.AuxiliaryNode[i]), types.YChild{"AuxiliaryNode", auxiliaryNodes.AuxiliaryNode[i]})
    }
    auxiliaryNodes.EntityData.Leafs = types.NewOrderedMap()

    auxiliaryNodes.EntityData.YListKeys = []string {}

    return &(auxiliaryNodes.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode
// Line configuration on a node
type Tty_AuxiliaryNodes_AuxiliaryNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Id interface{}

    // Auxiliary line.
    AuxiliaryLine Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine
}

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetEntityData() *types.CommonEntityData {
    auxiliaryNode.EntityData.YFilter = auxiliaryNode.YFilter
    auxiliaryNode.EntityData.YangName = "auxiliary-node"
    auxiliaryNode.EntityData.BundleName = "cisco_ios_xr"
    auxiliaryNode.EntityData.ParentYangName = "auxiliary-nodes"
    auxiliaryNode.EntityData.SegmentPath = "auxiliary-node" + types.AddKeyToken(auxiliaryNode.Id, "id")
    auxiliaryNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auxiliaryNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auxiliaryNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auxiliaryNode.EntityData.Children = types.NewOrderedMap()
    auxiliaryNode.EntityData.Children.Append("auxiliary-line", types.YChild{"AuxiliaryLine", &auxiliaryNode.AuxiliaryLine})
    auxiliaryNode.EntityData.Leafs = types.NewOrderedMap()
    auxiliaryNode.EntityData.Leafs.Append("id", types.YLeaf{"Id", auxiliaryNode.Id})

    auxiliaryNode.EntityData.YListKeys = []string {"Id"}

    return &(auxiliaryNode.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine
// Auxiliary line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of the auxiliary line.
    AuxiliaryStatistics Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics

    // Line state information.
    State Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State

    // Configuration information of the line.
    Configuration Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration
}

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetEntityData() *types.CommonEntityData {
    auxiliaryLine.EntityData.YFilter = auxiliaryLine.YFilter
    auxiliaryLine.EntityData.YangName = "auxiliary-line"
    auxiliaryLine.EntityData.BundleName = "cisco_ios_xr"
    auxiliaryLine.EntityData.ParentYangName = "auxiliary-node"
    auxiliaryLine.EntityData.SegmentPath = "auxiliary-line"
    auxiliaryLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auxiliaryLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auxiliaryLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auxiliaryLine.EntityData.Children = types.NewOrderedMap()
    auxiliaryLine.EntityData.Children.Append("auxiliary-statistics", types.YChild{"AuxiliaryStatistics", &auxiliaryLine.AuxiliaryStatistics})
    auxiliaryLine.EntityData.Children.Append("state", types.YChild{"State", &auxiliaryLine.State})
    auxiliaryLine.EntityData.Children.Append("configuration", types.YChild{"Configuration", &auxiliaryLine.Configuration})
    auxiliaryLine.EntityData.Leafs = types.NewOrderedMap()

    auxiliaryLine.EntityData.YListKeys = []string {}

    return &(auxiliaryLine.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics
// Statistics of the auxiliary line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RS232 statistics of console line.
    Rs232 Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232

    // General statistics of line.
    GeneralStatistics Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics

    // Exec related statistics.
    Exec Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec

    // AAA related statistics.
    Aaa Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa
}

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetEntityData() *types.CommonEntityData {
    auxiliaryStatistics.EntityData.YFilter = auxiliaryStatistics.YFilter
    auxiliaryStatistics.EntityData.YangName = "auxiliary-statistics"
    auxiliaryStatistics.EntityData.BundleName = "cisco_ios_xr"
    auxiliaryStatistics.EntityData.ParentYangName = "auxiliary-line"
    auxiliaryStatistics.EntityData.SegmentPath = "auxiliary-statistics"
    auxiliaryStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auxiliaryStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auxiliaryStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auxiliaryStatistics.EntityData.Children = types.NewOrderedMap()
    auxiliaryStatistics.EntityData.Children.Append("rs232", types.YChild{"Rs232", &auxiliaryStatistics.Rs232})
    auxiliaryStatistics.EntityData.Children.Append("general-statistics", types.YChild{"GeneralStatistics", &auxiliaryStatistics.GeneralStatistics})
    auxiliaryStatistics.EntityData.Children.Append("exec", types.YChild{"Exec", &auxiliaryStatistics.Exec})
    auxiliaryStatistics.EntityData.Children.Append("aaa", types.YChild{"Aaa", &auxiliaryStatistics.Aaa})
    auxiliaryStatistics.EntityData.Leafs = types.NewOrderedMap()

    auxiliaryStatistics.EntityData.YListKeys = []string {}

    return &(auxiliaryStatistics.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232
// RS232 statistics of console line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of databits. The type is interface{} with range: 0..4294967295.
    // Units are bit.
    DataBits interface{}

    // Exec disabled on TTY. The type is bool.
    ExecDisabled interface{}

    // Hardware flow control status. The type is interface{} with range:
    // 0..4294967295.
    HardwareFlowControlStatus interface{}

    // Parity status. The type is interface{} with range: 0..4294967295.
    ParityStatus interface{}

    // Inbound/Outbound baud rate in bps. The type is interface{} with range:
    // 0..4294967295. Units are bit/s.
    BaudRate interface{}

    // Number of stopbits. The type is interface{} with range: 0..4294967295.
    // Units are bit.
    StopBits interface{}

    // Overrun error count. The type is interface{} with range: 0..4294967295.
    OverrunErrorCount interface{}

    // Framing error count. The type is interface{} with range: 0..4294967295.
    FramingErrorCount interface{}

    // Parity error count. The type is interface{} with range: 0..4294967295.
    ParityErrorCount interface{}
}

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetEntityData() *types.CommonEntityData {
    rs232.EntityData.YFilter = rs232.YFilter
    rs232.EntityData.YangName = "rs232"
    rs232.EntityData.BundleName = "cisco_ios_xr"
    rs232.EntityData.ParentYangName = "auxiliary-statistics"
    rs232.EntityData.SegmentPath = "rs232"
    rs232.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rs232.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rs232.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rs232.EntityData.Children = types.NewOrderedMap()
    rs232.EntityData.Leafs = types.NewOrderedMap()
    rs232.EntityData.Leafs.Append("data-bits", types.YLeaf{"DataBits", rs232.DataBits})
    rs232.EntityData.Leafs.Append("exec-disabled", types.YLeaf{"ExecDisabled", rs232.ExecDisabled})
    rs232.EntityData.Leafs.Append("hardware-flow-control-status", types.YLeaf{"HardwareFlowControlStatus", rs232.HardwareFlowControlStatus})
    rs232.EntityData.Leafs.Append("parity-status", types.YLeaf{"ParityStatus", rs232.ParityStatus})
    rs232.EntityData.Leafs.Append("baud-rate", types.YLeaf{"BaudRate", rs232.BaudRate})
    rs232.EntityData.Leafs.Append("stop-bits", types.YLeaf{"StopBits", rs232.StopBits})
    rs232.EntityData.Leafs.Append("overrun-error-count", types.YLeaf{"OverrunErrorCount", rs232.OverrunErrorCount})
    rs232.EntityData.Leafs.Append("framing-error-count", types.YLeaf{"FramingErrorCount", rs232.FramingErrorCount})
    rs232.EntityData.Leafs.Append("parity-error-count", types.YLeaf{"ParityErrorCount", rs232.ParityErrorCount})

    rs232.EntityData.YListKeys = []string {}

    return &(rs232.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics
// General statistics of line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Terminal length. The type is interface{} with range: 0..4294967295.
    TerminalLength interface{}

    // Line width. The type is interface{} with range: 0..4294967295.
    TerminalWidth interface{}

    // Usable as async interface. The type is bool.
    AsyncInterface interface{}

    // Software flow control start char. The type is interface{} with range:
    // -128..127.
    FlowControlStartCharacter interface{}

    // Software flow control stop char. The type is interface{} with range:
    // -128..127.
    FlowControlStopCharacter interface{}

    // DNS resolution enabled. The type is bool.
    DomainLookupEnabled interface{}

    // MOTD banner enabled. The type is bool.
    MotdBannerEnabled interface{}

    // TTY private flag. The type is bool.
    PrivateFlag interface{}

    // Terminal type. The type is string.
    TerminalType interface{}

    // Absolute timeout period. The type is interface{} with range: 0..4294967295.
    AbsoluteTimeout interface{}

    // TTY idle time. The type is interface{} with range: 0..4294967295.
    IdleTime interface{}
}

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetEntityData() *types.CommonEntityData {
    generalStatistics.EntityData.YFilter = generalStatistics.YFilter
    generalStatistics.EntityData.YangName = "general-statistics"
    generalStatistics.EntityData.BundleName = "cisco_ios_xr"
    generalStatistics.EntityData.ParentYangName = "auxiliary-statistics"
    generalStatistics.EntityData.SegmentPath = "general-statistics"
    generalStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generalStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generalStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generalStatistics.EntityData.Children = types.NewOrderedMap()
    generalStatistics.EntityData.Leafs = types.NewOrderedMap()
    generalStatistics.EntityData.Leafs.Append("terminal-length", types.YLeaf{"TerminalLength", generalStatistics.TerminalLength})
    generalStatistics.EntityData.Leafs.Append("terminal-width", types.YLeaf{"TerminalWidth", generalStatistics.TerminalWidth})
    generalStatistics.EntityData.Leafs.Append("async-interface", types.YLeaf{"AsyncInterface", generalStatistics.AsyncInterface})
    generalStatistics.EntityData.Leafs.Append("flow-control-start-character", types.YLeaf{"FlowControlStartCharacter", generalStatistics.FlowControlStartCharacter})
    generalStatistics.EntityData.Leafs.Append("flow-control-stop-character", types.YLeaf{"FlowControlStopCharacter", generalStatistics.FlowControlStopCharacter})
    generalStatistics.EntityData.Leafs.Append("domain-lookup-enabled", types.YLeaf{"DomainLookupEnabled", generalStatistics.DomainLookupEnabled})
    generalStatistics.EntityData.Leafs.Append("motd-banner-enabled", types.YLeaf{"MotdBannerEnabled", generalStatistics.MotdBannerEnabled})
    generalStatistics.EntityData.Leafs.Append("private-flag", types.YLeaf{"PrivateFlag", generalStatistics.PrivateFlag})
    generalStatistics.EntityData.Leafs.Append("terminal-type", types.YLeaf{"TerminalType", generalStatistics.TerminalType})
    generalStatistics.EntityData.Leafs.Append("absolute-timeout", types.YLeaf{"AbsoluteTimeout", generalStatistics.AbsoluteTimeout})
    generalStatistics.EntityData.Leafs.Append("idle-time", types.YLeaf{"IdleTime", generalStatistics.IdleTime})

    generalStatistics.EntityData.YListKeys = []string {}

    return &(generalStatistics.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec
// Exec related statistics
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies whether timestamp is enabled or not. The type is bool.
    TimeStampEnabled interface{}
}

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetEntityData() *types.CommonEntityData {
    exec.EntityData.YFilter = exec.YFilter
    exec.EntityData.YangName = "exec"
    exec.EntityData.BundleName = "cisco_ios_xr"
    exec.EntityData.ParentYangName = "auxiliary-statistics"
    exec.EntityData.SegmentPath = "exec"
    exec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exec.EntityData.Children = types.NewOrderedMap()
    exec.EntityData.Leafs = types.NewOrderedMap()
    exec.EntityData.Leafs.Append("time-stamp-enabled", types.YLeaf{"TimeStampEnabled", exec.TimeStampEnabled})

    exec.EntityData.YListKeys = []string {}

    return &(exec.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa
// AAA related statistics
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The authenticated username. The type is string.
    UserName interface{}
}

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "auxiliary-statistics"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Leafs = types.NewOrderedMap()
    aaa.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", aaa.UserName})

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State
// Line state information
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information related to template applied to the line.
    Template Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template

    // General information.
    General Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General
}

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "auxiliary-line"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("template", types.YChild{"Template", &state.Template})
    state.EntityData.Children.Append("general", types.YChild{"General", &state.General})
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template
// Information related to template applied to the
// line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the template. The type is string.
    Name interface{}
}

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "state"
    template.EntityData.SegmentPath = "template"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("name", types.YLeaf{"Name", template.Name})

    template.EntityData.YListKeys = []string {}

    return &(template.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General
// General information
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // application running of on the tty line. The type is SessionOperation.
    Operation interface{}

    // State of the line. The type is LineState.
    GeneralState interface{}
}

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetEntityData() *types.CommonEntityData {
    general.EntityData.YFilter = general.YFilter
    general.EntityData.YangName = "general"
    general.EntityData.BundleName = "cisco_ios_xr"
    general.EntityData.ParentYangName = "state"
    general.EntityData.SegmentPath = "general"
    general.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    general.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    general.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    general.EntityData.Children = types.NewOrderedMap()
    general.EntityData.Leafs = types.NewOrderedMap()
    general.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", general.Operation})
    general.EntityData.Leafs.Append("general-state", types.YLeaf{"GeneralState", general.GeneralState})

    general.EntityData.YListKeys = []string {}

    return &(general.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration
// Configuration information of the line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Conection configuration information.
    ConnectionConfiguration Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration
}

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "auxiliary-line"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = types.NewOrderedMap()
    configuration.EntityData.Children.Append("connection-configuration", types.YChild{"ConnectionConfiguration", &configuration.ConnectionConfiguration})
    configuration.EntityData.Leafs = types.NewOrderedMap()

    configuration.EntityData.YListKeys = []string {}

    return &(configuration.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration
// Conection configuration information
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL for outbound traffic. The type is string.
    AclOut interface{}

    // ACL for inbound traffic. The type is string.
    AclIn interface{}

    // Protocols to use when connecting to the terminal server.
    TransportInput Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput
}

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetEntityData() *types.CommonEntityData {
    connectionConfiguration.EntityData.YFilter = connectionConfiguration.YFilter
    connectionConfiguration.EntityData.YangName = "connection-configuration"
    connectionConfiguration.EntityData.BundleName = "cisco_ios_xr"
    connectionConfiguration.EntityData.ParentYangName = "configuration"
    connectionConfiguration.EntityData.SegmentPath = "connection-configuration"
    connectionConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectionConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectionConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectionConfiguration.EntityData.Children = types.NewOrderedMap()
    connectionConfiguration.EntityData.Children.Append("transport-input", types.YChild{"TransportInput", &connectionConfiguration.TransportInput})
    connectionConfiguration.EntityData.Leafs = types.NewOrderedMap()
    connectionConfiguration.EntityData.Leafs.Append("acl-out", types.YLeaf{"AclOut", connectionConfiguration.AclOut})
    connectionConfiguration.EntityData.Leafs.Append("acl-in", types.YLeaf{"AclIn", connectionConfiguration.AclIn})

    connectionConfiguration.EntityData.YListKeys = []string {}

    return &(connectionConfiguration.EntityData)
}

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: 0..4294967295.
    None interface{}
}

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetEntityData() *types.CommonEntityData {
    transportInput.EntityData.YFilter = transportInput.YFilter
    transportInput.EntityData.YangName = "transport-input"
    transportInput.EntityData.BundleName = "cisco_ios_xr"
    transportInput.EntityData.ParentYangName = "connection-configuration"
    transportInput.EntityData.SegmentPath = "transport-input"
    transportInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportInput.EntityData.Children = types.NewOrderedMap()
    transportInput.EntityData.Leafs = types.NewOrderedMap()
    transportInput.EntityData.Leafs.Append("select", types.YLeaf{"Select", transportInput.Select})
    transportInput.EntityData.Leafs.Append("protocol1", types.YLeaf{"Protocol1", transportInput.Protocol1})
    transportInput.EntityData.Leafs.Append("protocol2", types.YLeaf{"Protocol2", transportInput.Protocol2})
    transportInput.EntityData.Leafs.Append("none", types.YLeaf{"None", transportInput.None})

    transportInput.EntityData.YListKeys = []string {}

    return &(transportInput.EntityData)
}

