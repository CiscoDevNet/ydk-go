// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-server package operational data.
// 
// This module contains definitions
// for the following management objects:
//   tty: TTY Line Configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of Nodes for console.
    ConsoleNodes Tty_ConsoleNodes

    // List of VTY lines.
    VtyLines Tty_VtyLines

    // List of Nodes attached with an auxiliary line.
    AuxiliaryNodes Tty_AuxiliaryNodes
}

func (tty *Tty) GetFilter() yfilter.YFilter { return tty.YFilter }

func (tty *Tty) SetFilter(yf yfilter.YFilter) { tty.YFilter = yf }

func (tty *Tty) GetGoName(yname string) string {
    if yname == "console-nodes" { return "ConsoleNodes" }
    if yname == "vty-lines" { return "VtyLines" }
    if yname == "auxiliary-nodes" { return "AuxiliaryNodes" }
    return ""
}

func (tty *Tty) GetSegmentPath() string {
    return "Cisco-IOS-XR-tty-server-oper:tty"
}

func (tty *Tty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "console-nodes" {
        return &tty.ConsoleNodes
    }
    if childYangName == "vty-lines" {
        return &tty.VtyLines
    }
    if childYangName == "auxiliary-nodes" {
        return &tty.AuxiliaryNodes
    }
    return nil
}

func (tty *Tty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["console-nodes"] = &tty.ConsoleNodes
    children["vty-lines"] = &tty.VtyLines
    children["auxiliary-nodes"] = &tty.AuxiliaryNodes
    return children
}

func (tty *Tty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tty *Tty) GetBundleName() string { return "cisco_ios_xr" }

func (tty *Tty) GetYangName() string { return "tty" }

func (tty *Tty) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tty *Tty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tty *Tty) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tty *Tty) SetParent(parent types.Entity) { tty.parent = parent }

func (tty *Tty) GetParent() types.Entity { return tty.parent }

func (tty *Tty) GetParentYangName() string { return "Cisco-IOS-XR-tty-server-oper" }

// Tty_ConsoleNodes
// List of Nodes for console
type Tty_ConsoleNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Console line configuration on a node. The type is slice of
    // Tty_ConsoleNodes_ConsoleNode.
    ConsoleNode []Tty_ConsoleNodes_ConsoleNode
}

func (consoleNodes *Tty_ConsoleNodes) GetFilter() yfilter.YFilter { return consoleNodes.YFilter }

func (consoleNodes *Tty_ConsoleNodes) SetFilter(yf yfilter.YFilter) { consoleNodes.YFilter = yf }

func (consoleNodes *Tty_ConsoleNodes) GetGoName(yname string) string {
    if yname == "console-node" { return "ConsoleNode" }
    return ""
}

func (consoleNodes *Tty_ConsoleNodes) GetSegmentPath() string {
    return "console-nodes"
}

func (consoleNodes *Tty_ConsoleNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "console-node" {
        for _, c := range consoleNodes.ConsoleNode {
            if consoleNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tty_ConsoleNodes_ConsoleNode{}
        consoleNodes.ConsoleNode = append(consoleNodes.ConsoleNode, child)
        return &consoleNodes.ConsoleNode[len(consoleNodes.ConsoleNode)-1]
    }
    return nil
}

func (consoleNodes *Tty_ConsoleNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range consoleNodes.ConsoleNode {
        children[consoleNodes.ConsoleNode[i].GetSegmentPath()] = &consoleNodes.ConsoleNode[i]
    }
    return children
}

func (consoleNodes *Tty_ConsoleNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (consoleNodes *Tty_ConsoleNodes) GetBundleName() string { return "cisco_ios_xr" }

func (consoleNodes *Tty_ConsoleNodes) GetYangName() string { return "console-nodes" }

func (consoleNodes *Tty_ConsoleNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (consoleNodes *Tty_ConsoleNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (consoleNodes *Tty_ConsoleNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (consoleNodes *Tty_ConsoleNodes) SetParent(parent types.Entity) { consoleNodes.parent = parent }

func (consoleNodes *Tty_ConsoleNodes) GetParent() types.Entity { return consoleNodes.parent }

func (consoleNodes *Tty_ConsoleNodes) GetParentYangName() string { return "tty" }

// Tty_ConsoleNodes_ConsoleNode
// Console line configuration on a node
type Tty_ConsoleNodes_ConsoleNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Id interface{}

    // Console line.
    ConsoleLine Tty_ConsoleNodes_ConsoleNode_ConsoleLine
}

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetFilter() yfilter.YFilter { return consoleNode.YFilter }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) SetFilter(yf yfilter.YFilter) { consoleNode.YFilter = yf }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "console-line" { return "ConsoleLine" }
    return ""
}

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetSegmentPath() string {
    return "console-node" + "[id='" + fmt.Sprintf("%v", consoleNode.Id) + "']"
}

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "console-line" {
        return &consoleNode.ConsoleLine
    }
    return nil
}

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["console-line"] = &consoleNode.ConsoleLine
    return children
}

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = consoleNode.Id
    return leafs
}

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetBundleName() string { return "cisco_ios_xr" }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetYangName() string { return "console-node" }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) SetParent(parent types.Entity) { consoleNode.parent = parent }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetParent() types.Entity { return consoleNode.parent }

func (consoleNode *Tty_ConsoleNodes_ConsoleNode) GetParentYangName() string { return "console-nodes" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine
// Console line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics of the console line.
    ConsoleStatistics Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics

    // Line state information.
    State Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State

    // Configuration information of the line.
    Configuration Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration
}

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetFilter() yfilter.YFilter { return consoleLine.YFilter }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) SetFilter(yf yfilter.YFilter) { consoleLine.YFilter = yf }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetGoName(yname string) string {
    if yname == "console-statistics" { return "ConsoleStatistics" }
    if yname == "state" { return "State" }
    if yname == "configuration" { return "Configuration" }
    return ""
}

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetSegmentPath() string {
    return "console-line"
}

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "console-statistics" {
        return &consoleLine.ConsoleStatistics
    }
    if childYangName == "state" {
        return &consoleLine.State
    }
    if childYangName == "configuration" {
        return &consoleLine.Configuration
    }
    return nil
}

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["console-statistics"] = &consoleLine.ConsoleStatistics
    children["state"] = &consoleLine.State
    children["configuration"] = &consoleLine.Configuration
    return children
}

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetBundleName() string { return "cisco_ios_xr" }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetYangName() string { return "console-line" }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) SetParent(parent types.Entity) { consoleLine.parent = parent }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetParent() types.Entity { return consoleLine.parent }

func (consoleLine *Tty_ConsoleNodes_ConsoleNode_ConsoleLine) GetParentYangName() string { return "console-node" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics
// Statistics of the console line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics struct {
    parent types.Entity
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

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetFilter() yfilter.YFilter { return consoleStatistics.YFilter }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) SetFilter(yf yfilter.YFilter) { consoleStatistics.YFilter = yf }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetGoName(yname string) string {
    if yname == "rs232" { return "Rs232" }
    if yname == "general-statistics" { return "GeneralStatistics" }
    if yname == "exec" { return "Exec" }
    if yname == "aaa" { return "Aaa" }
    return ""
}

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetSegmentPath() string {
    return "console-statistics"
}

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rs232" {
        return &consoleStatistics.Rs232
    }
    if childYangName == "general-statistics" {
        return &consoleStatistics.GeneralStatistics
    }
    if childYangName == "exec" {
        return &consoleStatistics.Exec
    }
    if childYangName == "aaa" {
        return &consoleStatistics.Aaa
    }
    return nil
}

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rs232"] = &consoleStatistics.Rs232
    children["general-statistics"] = &consoleStatistics.GeneralStatistics
    children["exec"] = &consoleStatistics.Exec
    children["aaa"] = &consoleStatistics.Aaa
    return children
}

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetYangName() string { return "console-statistics" }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) SetParent(parent types.Entity) { consoleStatistics.parent = parent }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetParent() types.Entity { return consoleStatistics.parent }

func (consoleStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics) GetParentYangName() string { return "console-line" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232
// RS232 statistics of console line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232 struct {
    parent types.Entity
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

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetFilter() yfilter.YFilter { return rs232.YFilter }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) SetFilter(yf yfilter.YFilter) { rs232.YFilter = yf }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetGoName(yname string) string {
    if yname == "data-bits" { return "DataBits" }
    if yname == "exec-disabled" { return "ExecDisabled" }
    if yname == "hardware-flow-control-status" { return "HardwareFlowControlStatus" }
    if yname == "parity-status" { return "ParityStatus" }
    if yname == "baud-rate" { return "BaudRate" }
    if yname == "stop-bits" { return "StopBits" }
    if yname == "overrun-error-count" { return "OverrunErrorCount" }
    if yname == "framing-error-count" { return "FramingErrorCount" }
    if yname == "parity-error-count" { return "ParityErrorCount" }
    return ""
}

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetSegmentPath() string {
    return "rs232"
}

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-bits"] = rs232.DataBits
    leafs["exec-disabled"] = rs232.ExecDisabled
    leafs["hardware-flow-control-status"] = rs232.HardwareFlowControlStatus
    leafs["parity-status"] = rs232.ParityStatus
    leafs["baud-rate"] = rs232.BaudRate
    leafs["stop-bits"] = rs232.StopBits
    leafs["overrun-error-count"] = rs232.OverrunErrorCount
    leafs["framing-error-count"] = rs232.FramingErrorCount
    leafs["parity-error-count"] = rs232.ParityErrorCount
    return leafs
}

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetBundleName() string { return "cisco_ios_xr" }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetYangName() string { return "rs232" }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) SetParent(parent types.Entity) { rs232.parent = parent }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetParent() types.Entity { return rs232.parent }

func (rs232 *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Rs232) GetParentYangName() string { return "console-statistics" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics
// General statistics of line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics struct {
    parent types.Entity
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

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetFilter() yfilter.YFilter { return generalStatistics.YFilter }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) SetFilter(yf yfilter.YFilter) { generalStatistics.YFilter = yf }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetGoName(yname string) string {
    if yname == "terminal-length" { return "TerminalLength" }
    if yname == "terminal-width" { return "TerminalWidth" }
    if yname == "async-interface" { return "AsyncInterface" }
    if yname == "flow-control-start-character" { return "FlowControlStartCharacter" }
    if yname == "flow-control-stop-character" { return "FlowControlStopCharacter" }
    if yname == "domain-lookup-enabled" { return "DomainLookupEnabled" }
    if yname == "motd-banner-enabled" { return "MotdBannerEnabled" }
    if yname == "private-flag" { return "PrivateFlag" }
    if yname == "terminal-type" { return "TerminalType" }
    if yname == "absolute-timeout" { return "AbsoluteTimeout" }
    if yname == "idle-time" { return "IdleTime" }
    return ""
}

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetSegmentPath() string {
    return "general-statistics"
}

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminal-length"] = generalStatistics.TerminalLength
    leafs["terminal-width"] = generalStatistics.TerminalWidth
    leafs["async-interface"] = generalStatistics.AsyncInterface
    leafs["flow-control-start-character"] = generalStatistics.FlowControlStartCharacter
    leafs["flow-control-stop-character"] = generalStatistics.FlowControlStopCharacter
    leafs["domain-lookup-enabled"] = generalStatistics.DomainLookupEnabled
    leafs["motd-banner-enabled"] = generalStatistics.MotdBannerEnabled
    leafs["private-flag"] = generalStatistics.PrivateFlag
    leafs["terminal-type"] = generalStatistics.TerminalType
    leafs["absolute-timeout"] = generalStatistics.AbsoluteTimeout
    leafs["idle-time"] = generalStatistics.IdleTime
    return leafs
}

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetYangName() string { return "general-statistics" }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) SetParent(parent types.Entity) { generalStatistics.parent = parent }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetParent() types.Entity { return generalStatistics.parent }

func (generalStatistics *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_GeneralStatistics) GetParentYangName() string { return "console-statistics" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec
// Exec related statistics
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifies whether timestamp is enabled or not. The type is bool.
    TimeStampEnabled interface{}
}

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetFilter() yfilter.YFilter { return exec.YFilter }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) SetFilter(yf yfilter.YFilter) { exec.YFilter = yf }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetGoName(yname string) string {
    if yname == "time-stamp-enabled" { return "TimeStampEnabled" }
    return ""
}

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetSegmentPath() string {
    return "exec"
}

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-stamp-enabled"] = exec.TimeStampEnabled
    return leafs
}

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetBundleName() string { return "cisco_ios_xr" }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetYangName() string { return "exec" }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) SetParent(parent types.Entity) { exec.parent = parent }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetParent() types.Entity { return exec.parent }

func (exec *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Exec) GetParentYangName() string { return "console-statistics" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa
// AAA related statistics
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The authenticated username. The type is string.
    UserName interface{}
}

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetGoName(yname string) string {
    if yname == "user-name" { return "UserName" }
    return ""
}

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetSegmentPath() string {
    return "aaa"
}

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["user-name"] = aaa.UserName
    return leafs
}

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetBundleName() string { return "cisco_ios_xr" }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetYangName() string { return "aaa" }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) SetParent(parent types.Entity) { aaa.parent = parent }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetParent() types.Entity { return aaa.parent }

func (aaa *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_ConsoleStatistics_Aaa) GetParentYangName() string { return "console-statistics" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State
// Line state information
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information related to template applied to the line.
    Template Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template

    // General information.
    General Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General
}

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    if yname == "general" { return "General" }
    return ""
}

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetSegmentPath() string {
    return "state"
}

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        return &state.Template
    }
    if childYangName == "general" {
        return &state.General
    }
    return nil
}

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["template"] = &state.Template
    children["general"] = &state.General
    return children
}

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetYangName() string { return "state" }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetParent() types.Entity { return state.parent }

func (state *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State) GetParentYangName() string { return "console-line" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template
// Information related to template applied to the
// line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the template. The type is string.
    Name interface{}
}

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetSegmentPath() string {
    return "template"
}

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = template.Name
    return leafs
}

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetYangName() string { return "template" }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetParent() types.Entity { return template.parent }

func (template *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_Template) GetParentYangName() string { return "state" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General
// General information
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // application running of on the tty line. The type is SessionOperation.
    Operation interface{}

    // State of the line. The type is LineState.
    GeneralState interface{}
}

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetFilter() yfilter.YFilter { return general.YFilter }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) SetFilter(yf yfilter.YFilter) { general.YFilter = yf }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetGoName(yname string) string {
    if yname == "operation" { return "Operation" }
    if yname == "general-state" { return "GeneralState" }
    return ""
}

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetSegmentPath() string {
    return "general"
}

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation"] = general.Operation
    leafs["general-state"] = general.GeneralState
    return leafs
}

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetBundleName() string { return "cisco_ios_xr" }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetYangName() string { return "general" }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) SetParent(parent types.Entity) { general.parent = parent }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetParent() types.Entity { return general.parent }

func (general *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_State_General) GetParentYangName() string { return "state" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration
// Configuration information of the line
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Conection configuration information.
    ConnectionConfiguration Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration
}

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetGoName(yname string) string {
    if yname == "connection-configuration" { return "ConnectionConfiguration" }
    return ""
}

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connection-configuration" {
        return &configuration.ConnectionConfiguration
    }
    return nil
}

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connection-configuration"] = &configuration.ConnectionConfiguration
    return children
}

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetYangName() string { return "configuration" }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration) GetParentYangName() string { return "console-line" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration
// Conection configuration information
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL for outbound traffic. The type is string.
    AclOut interface{}

    // ACL for inbound traffic. The type is string.
    AclIn interface{}

    // Protocols to use when connecting to the terminal server.
    TransportInput Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput
}

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetFilter() yfilter.YFilter { return connectionConfiguration.YFilter }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) SetFilter(yf yfilter.YFilter) { connectionConfiguration.YFilter = yf }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetGoName(yname string) string {
    if yname == "acl-out" { return "AclOut" }
    if yname == "acl-in" { return "AclIn" }
    if yname == "transport-input" { return "TransportInput" }
    return ""
}

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetSegmentPath() string {
    return "connection-configuration"
}

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport-input" {
        return &connectionConfiguration.TransportInput
    }
    return nil
}

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transport-input"] = &connectionConfiguration.TransportInput
    return children
}

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl-out"] = connectionConfiguration.AclOut
    leafs["acl-in"] = connectionConfiguration.AclIn
    return leafs
}

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetYangName() string { return "connection-configuration" }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) SetParent(parent types.Entity) { connectionConfiguration.parent = parent }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetParent() types.Entity { return connectionConfiguration.parent }

func (connectionConfiguration *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration) GetParentYangName() string { return "configuration" }

// Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: -2147483648..2147483647.
    None interface{}
}

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetFilter() yfilter.YFilter { return transportInput.YFilter }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) SetFilter(yf yfilter.YFilter) { transportInput.YFilter = yf }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetGoName(yname string) string {
    if yname == "select" { return "Select" }
    if yname == "protocol1" { return "Protocol1" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "none" { return "None" }
    return ""
}

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetSegmentPath() string {
    return "transport-input"
}

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["select"] = transportInput.Select
    leafs["protocol1"] = transportInput.Protocol1
    leafs["protocol2"] = transportInput.Protocol2
    leafs["none"] = transportInput.None
    return leafs
}

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetBundleName() string { return "cisco_ios_xr" }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetYangName() string { return "transport-input" }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) SetParent(parent types.Entity) { transportInput.parent = parent }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetParent() types.Entity { return transportInput.parent }

func (transportInput *Tty_ConsoleNodes_ConsoleNode_ConsoleLine_Configuration_ConnectionConfiguration_TransportInput) GetParentYangName() string { return "connection-configuration" }

// Tty_VtyLines
// List of VTY lines
type Tty_VtyLines struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VTY Line. The type is slice of Tty_VtyLines_VtyLine.
    VtyLine []Tty_VtyLines_VtyLine
}

func (vtyLines *Tty_VtyLines) GetFilter() yfilter.YFilter { return vtyLines.YFilter }

func (vtyLines *Tty_VtyLines) SetFilter(yf yfilter.YFilter) { vtyLines.YFilter = yf }

func (vtyLines *Tty_VtyLines) GetGoName(yname string) string {
    if yname == "vty-line" { return "VtyLine" }
    return ""
}

func (vtyLines *Tty_VtyLines) GetSegmentPath() string {
    return "vty-lines"
}

func (vtyLines *Tty_VtyLines) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vty-line" {
        for _, c := range vtyLines.VtyLine {
            if vtyLines.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tty_VtyLines_VtyLine{}
        vtyLines.VtyLine = append(vtyLines.VtyLine, child)
        return &vtyLines.VtyLine[len(vtyLines.VtyLine)-1]
    }
    return nil
}

func (vtyLines *Tty_VtyLines) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtyLines.VtyLine {
        children[vtyLines.VtyLine[i].GetSegmentPath()] = &vtyLines.VtyLine[i]
    }
    return children
}

func (vtyLines *Tty_VtyLines) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtyLines *Tty_VtyLines) GetBundleName() string { return "cisco_ios_xr" }

func (vtyLines *Tty_VtyLines) GetYangName() string { return "vty-lines" }

func (vtyLines *Tty_VtyLines) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vtyLines *Tty_VtyLines) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vtyLines *Tty_VtyLines) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vtyLines *Tty_VtyLines) SetParent(parent types.Entity) { vtyLines.parent = parent }

func (vtyLines *Tty_VtyLines) GetParent() types.Entity { return vtyLines.parent }

func (vtyLines *Tty_VtyLines) GetParentYangName() string { return "tty" }

// Tty_VtyLines_VtyLine
// VTY Line
type Tty_VtyLines_VtyLine struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VTY Line number. The type is interface{} with
    // range: -2147483648..2147483647.
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

func (vtyLine *Tty_VtyLines_VtyLine) GetFilter() yfilter.YFilter { return vtyLine.YFilter }

func (vtyLine *Tty_VtyLines_VtyLine) SetFilter(yf yfilter.YFilter) { vtyLine.YFilter = yf }

func (vtyLine *Tty_VtyLines_VtyLine) GetGoName(yname string) string {
    if yname == "line-number" { return "LineNumber" }
    if yname == "vty-statistics" { return "VtyStatistics" }
    if yname == "state" { return "State" }
    if yname == "configuration" { return "Configuration" }
    if yname == "Cisco-IOS-XR-tty-management-oper:sessions" { return "Sessions" }
    return ""
}

func (vtyLine *Tty_VtyLines_VtyLine) GetSegmentPath() string {
    return "vty-line" + "[line-number='" + fmt.Sprintf("%v", vtyLine.LineNumber) + "']"
}

func (vtyLine *Tty_VtyLines_VtyLine) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vty-statistics" {
        return &vtyLine.VtyStatistics
    }
    if childYangName == "state" {
        return &vtyLine.State
    }
    if childYangName == "configuration" {
        return &vtyLine.Configuration
    }
    if childYangName == "Cisco-IOS-XR-tty-management-oper:sessions" {
        return &vtyLine.Sessions
    }
    return nil
}

func (vtyLine *Tty_VtyLines_VtyLine) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vty-statistics"] = &vtyLine.VtyStatistics
    children["state"] = &vtyLine.State
    children["configuration"] = &vtyLine.Configuration
    children["Cisco-IOS-XR-tty-management-oper:sessions"] = &vtyLine.Sessions
    return children
}

func (vtyLine *Tty_VtyLines_VtyLine) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["line-number"] = vtyLine.LineNumber
    return leafs
}

func (vtyLine *Tty_VtyLines_VtyLine) GetBundleName() string { return "cisco_ios_xr" }

func (vtyLine *Tty_VtyLines_VtyLine) GetYangName() string { return "vty-line" }

func (vtyLine *Tty_VtyLines_VtyLine) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vtyLine *Tty_VtyLines_VtyLine) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vtyLine *Tty_VtyLines_VtyLine) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vtyLine *Tty_VtyLines_VtyLine) SetParent(parent types.Entity) { vtyLine.parent = parent }

func (vtyLine *Tty_VtyLines_VtyLine) GetParent() types.Entity { return vtyLine.parent }

func (vtyLine *Tty_VtyLines_VtyLine) GetParentYangName() string { return "vty-lines" }

// Tty_VtyLines_VtyLine_VtyStatistics
// Statistics of the VTY line
type Tty_VtyLines_VtyLine_VtyStatistics struct {
    parent types.Entity
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

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetFilter() yfilter.YFilter { return vtyStatistics.YFilter }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) SetFilter(yf yfilter.YFilter) { vtyStatistics.YFilter = yf }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetGoName(yname string) string {
    if yname == "connection" { return "Connection" }
    if yname == "general-statistics" { return "GeneralStatistics" }
    if yname == "exec" { return "Exec" }
    if yname == "aaa" { return "Aaa" }
    return ""
}

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetSegmentPath() string {
    return "vty-statistics"
}

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connection" {
        return &vtyStatistics.Connection
    }
    if childYangName == "general-statistics" {
        return &vtyStatistics.GeneralStatistics
    }
    if childYangName == "exec" {
        return &vtyStatistics.Exec
    }
    if childYangName == "aaa" {
        return &vtyStatistics.Aaa
    }
    return nil
}

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connection"] = &vtyStatistics.Connection
    children["general-statistics"] = &vtyStatistics.GeneralStatistics
    children["exec"] = &vtyStatistics.Exec
    children["aaa"] = &vtyStatistics.Aaa
    return children
}

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetYangName() string { return "vty-statistics" }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) SetParent(parent types.Entity) { vtyStatistics.parent = parent }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetParent() types.Entity { return vtyStatistics.parent }

func (vtyStatistics *Tty_VtyLines_VtyLine_VtyStatistics) GetParentYangName() string { return "vty-line" }

// Tty_VtyLines_VtyLine_VtyStatistics_Connection
// Connection related statistics
type Tty_VtyLines_VtyLine_VtyStatistics_Connection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Incoming host address(max). The type is string with length: 0..46.
    IncomingHostAddress interface{}

    // Incoming host address family. The type is interface{} with range:
    // 0..4294967295.
    HostAddressFamily interface{}

    // Input transport. The type is interface{} with range: 0..4294967295.
    Service interface{}
}

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetFilter() yfilter.YFilter { return connection.YFilter }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) SetFilter(yf yfilter.YFilter) { connection.YFilter = yf }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetGoName(yname string) string {
    if yname == "incoming-host-address" { return "IncomingHostAddress" }
    if yname == "host-address-family" { return "HostAddressFamily" }
    if yname == "service" { return "Service" }
    return ""
}

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetSegmentPath() string {
    return "connection"
}

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["incoming-host-address"] = connection.IncomingHostAddress
    leafs["host-address-family"] = connection.HostAddressFamily
    leafs["service"] = connection.Service
    return leafs
}

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetBundleName() string { return "cisco_ios_xr" }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetYangName() string { return "connection" }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) SetParent(parent types.Entity) { connection.parent = parent }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetParent() types.Entity { return connection.parent }

func (connection *Tty_VtyLines_VtyLine_VtyStatistics_Connection) GetParentYangName() string { return "vty-statistics" }

// Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics
// General statistics of line
type Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics struct {
    parent types.Entity
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

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetFilter() yfilter.YFilter { return generalStatistics.YFilter }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) SetFilter(yf yfilter.YFilter) { generalStatistics.YFilter = yf }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetGoName(yname string) string {
    if yname == "terminal-length" { return "TerminalLength" }
    if yname == "terminal-width" { return "TerminalWidth" }
    if yname == "async-interface" { return "AsyncInterface" }
    if yname == "flow-control-start-character" { return "FlowControlStartCharacter" }
    if yname == "flow-control-stop-character" { return "FlowControlStopCharacter" }
    if yname == "domain-lookup-enabled" { return "DomainLookupEnabled" }
    if yname == "motd-banner-enabled" { return "MotdBannerEnabled" }
    if yname == "private-flag" { return "PrivateFlag" }
    if yname == "terminal-type" { return "TerminalType" }
    if yname == "absolute-timeout" { return "AbsoluteTimeout" }
    if yname == "idle-time" { return "IdleTime" }
    return ""
}

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetSegmentPath() string {
    return "general-statistics"
}

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminal-length"] = generalStatistics.TerminalLength
    leafs["terminal-width"] = generalStatistics.TerminalWidth
    leafs["async-interface"] = generalStatistics.AsyncInterface
    leafs["flow-control-start-character"] = generalStatistics.FlowControlStartCharacter
    leafs["flow-control-stop-character"] = generalStatistics.FlowControlStopCharacter
    leafs["domain-lookup-enabled"] = generalStatistics.DomainLookupEnabled
    leafs["motd-banner-enabled"] = generalStatistics.MotdBannerEnabled
    leafs["private-flag"] = generalStatistics.PrivateFlag
    leafs["terminal-type"] = generalStatistics.TerminalType
    leafs["absolute-timeout"] = generalStatistics.AbsoluteTimeout
    leafs["idle-time"] = generalStatistics.IdleTime
    return leafs
}

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetYangName() string { return "general-statistics" }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) SetParent(parent types.Entity) { generalStatistics.parent = parent }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetParent() types.Entity { return generalStatistics.parent }

func (generalStatistics *Tty_VtyLines_VtyLine_VtyStatistics_GeneralStatistics) GetParentYangName() string { return "vty-statistics" }

// Tty_VtyLines_VtyLine_VtyStatistics_Exec
// Exec related statistics
type Tty_VtyLines_VtyLine_VtyStatistics_Exec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifies whether timestamp is enabled or not. The type is bool.
    TimeStampEnabled interface{}
}

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetFilter() yfilter.YFilter { return exec.YFilter }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) SetFilter(yf yfilter.YFilter) { exec.YFilter = yf }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetGoName(yname string) string {
    if yname == "time-stamp-enabled" { return "TimeStampEnabled" }
    return ""
}

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetSegmentPath() string {
    return "exec"
}

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-stamp-enabled"] = exec.TimeStampEnabled
    return leafs
}

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetBundleName() string { return "cisco_ios_xr" }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetYangName() string { return "exec" }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) SetParent(parent types.Entity) { exec.parent = parent }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetParent() types.Entity { return exec.parent }

func (exec *Tty_VtyLines_VtyLine_VtyStatistics_Exec) GetParentYangName() string { return "vty-statistics" }

// Tty_VtyLines_VtyLine_VtyStatistics_Aaa
// AAA related statistics
type Tty_VtyLines_VtyLine_VtyStatistics_Aaa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The authenticated username. The type is string.
    UserName interface{}
}

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetGoName(yname string) string {
    if yname == "user-name" { return "UserName" }
    return ""
}

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetSegmentPath() string {
    return "aaa"
}

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["user-name"] = aaa.UserName
    return leafs
}

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetBundleName() string { return "cisco_ios_xr" }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetYangName() string { return "aaa" }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) SetParent(parent types.Entity) { aaa.parent = parent }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetParent() types.Entity { return aaa.parent }

func (aaa *Tty_VtyLines_VtyLine_VtyStatistics_Aaa) GetParentYangName() string { return "vty-statistics" }

// Tty_VtyLines_VtyLine_State
// Line state information
type Tty_VtyLines_VtyLine_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information related to template applied to the line.
    Template Tty_VtyLines_VtyLine_State_Template

    // General information.
    General Tty_VtyLines_VtyLine_State_General
}

func (state *Tty_VtyLines_VtyLine_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Tty_VtyLines_VtyLine_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Tty_VtyLines_VtyLine_State) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    if yname == "general" { return "General" }
    return ""
}

func (state *Tty_VtyLines_VtyLine_State) GetSegmentPath() string {
    return "state"
}

func (state *Tty_VtyLines_VtyLine_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        return &state.Template
    }
    if childYangName == "general" {
        return &state.General
    }
    return nil
}

func (state *Tty_VtyLines_VtyLine_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["template"] = &state.Template
    children["general"] = &state.General
    return children
}

func (state *Tty_VtyLines_VtyLine_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *Tty_VtyLines_VtyLine_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *Tty_VtyLines_VtyLine_State) GetYangName() string { return "state" }

func (state *Tty_VtyLines_VtyLine_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *Tty_VtyLines_VtyLine_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *Tty_VtyLines_VtyLine_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *Tty_VtyLines_VtyLine_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Tty_VtyLines_VtyLine_State) GetParent() types.Entity { return state.parent }

func (state *Tty_VtyLines_VtyLine_State) GetParentYangName() string { return "vty-line" }

// Tty_VtyLines_VtyLine_State_Template
// Information related to template applied to the
// line
type Tty_VtyLines_VtyLine_State_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the template. The type is string.
    Name interface{}
}

func (template *Tty_VtyLines_VtyLine_State_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *Tty_VtyLines_VtyLine_State_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *Tty_VtyLines_VtyLine_State_Template) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (template *Tty_VtyLines_VtyLine_State_Template) GetSegmentPath() string {
    return "template"
}

func (template *Tty_VtyLines_VtyLine_State_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *Tty_VtyLines_VtyLine_State_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *Tty_VtyLines_VtyLine_State_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = template.Name
    return leafs
}

func (template *Tty_VtyLines_VtyLine_State_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *Tty_VtyLines_VtyLine_State_Template) GetYangName() string { return "template" }

func (template *Tty_VtyLines_VtyLine_State_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *Tty_VtyLines_VtyLine_State_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *Tty_VtyLines_VtyLine_State_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *Tty_VtyLines_VtyLine_State_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *Tty_VtyLines_VtyLine_State_Template) GetParent() types.Entity { return template.parent }

func (template *Tty_VtyLines_VtyLine_State_Template) GetParentYangName() string { return "state" }

// Tty_VtyLines_VtyLine_State_General
// General information
type Tty_VtyLines_VtyLine_State_General struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // application running of on the tty line. The type is SessionOperation.
    Operation interface{}

    // State of the line. The type is LineState.
    GeneralState interface{}
}

func (general *Tty_VtyLines_VtyLine_State_General) GetFilter() yfilter.YFilter { return general.YFilter }

func (general *Tty_VtyLines_VtyLine_State_General) SetFilter(yf yfilter.YFilter) { general.YFilter = yf }

func (general *Tty_VtyLines_VtyLine_State_General) GetGoName(yname string) string {
    if yname == "operation" { return "Operation" }
    if yname == "general-state" { return "GeneralState" }
    return ""
}

func (general *Tty_VtyLines_VtyLine_State_General) GetSegmentPath() string {
    return "general"
}

func (general *Tty_VtyLines_VtyLine_State_General) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (general *Tty_VtyLines_VtyLine_State_General) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (general *Tty_VtyLines_VtyLine_State_General) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation"] = general.Operation
    leafs["general-state"] = general.GeneralState
    return leafs
}

func (general *Tty_VtyLines_VtyLine_State_General) GetBundleName() string { return "cisco_ios_xr" }

func (general *Tty_VtyLines_VtyLine_State_General) GetYangName() string { return "general" }

func (general *Tty_VtyLines_VtyLine_State_General) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (general *Tty_VtyLines_VtyLine_State_General) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (general *Tty_VtyLines_VtyLine_State_General) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (general *Tty_VtyLines_VtyLine_State_General) SetParent(parent types.Entity) { general.parent = parent }

func (general *Tty_VtyLines_VtyLine_State_General) GetParent() types.Entity { return general.parent }

func (general *Tty_VtyLines_VtyLine_State_General) GetParentYangName() string { return "state" }

// Tty_VtyLines_VtyLine_Configuration
// Configuration information of the line
type Tty_VtyLines_VtyLine_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Conection configuration information.
    ConnectionConfiguration Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration
}

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Tty_VtyLines_VtyLine_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetGoName(yname string) string {
    if yname == "connection-configuration" { return "ConnectionConfiguration" }
    return ""
}

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connection-configuration" {
        return &configuration.ConnectionConfiguration
    }
    return nil
}

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connection-configuration"] = &configuration.ConnectionConfiguration
    return children
}

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetYangName() string { return "configuration" }

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Tty_VtyLines_VtyLine_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Tty_VtyLines_VtyLine_Configuration) GetParentYangName() string { return "vty-line" }

// Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration
// Conection configuration information
type Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL for outbound traffic. The type is string.
    AclOut interface{}

    // ACL for inbound traffic. The type is string.
    AclIn interface{}

    // Protocols to use when connecting to the terminal server.
    TransportInput Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput
}

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetFilter() yfilter.YFilter { return connectionConfiguration.YFilter }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) SetFilter(yf yfilter.YFilter) { connectionConfiguration.YFilter = yf }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetGoName(yname string) string {
    if yname == "acl-out" { return "AclOut" }
    if yname == "acl-in" { return "AclIn" }
    if yname == "transport-input" { return "TransportInput" }
    return ""
}

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetSegmentPath() string {
    return "connection-configuration"
}

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport-input" {
        return &connectionConfiguration.TransportInput
    }
    return nil
}

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transport-input"] = &connectionConfiguration.TransportInput
    return children
}

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl-out"] = connectionConfiguration.AclOut
    leafs["acl-in"] = connectionConfiguration.AclIn
    return leafs
}

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetYangName() string { return "connection-configuration" }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) SetParent(parent types.Entity) { connectionConfiguration.parent = parent }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetParent() types.Entity { return connectionConfiguration.parent }

func (connectionConfiguration *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration) GetParentYangName() string { return "configuration" }

// Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: -2147483648..2147483647.
    None interface{}
}

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetFilter() yfilter.YFilter { return transportInput.YFilter }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) SetFilter(yf yfilter.YFilter) { transportInput.YFilter = yf }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetGoName(yname string) string {
    if yname == "select" { return "Select" }
    if yname == "protocol1" { return "Protocol1" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "none" { return "None" }
    return ""
}

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetSegmentPath() string {
    return "transport-input"
}

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["select"] = transportInput.Select
    leafs["protocol1"] = transportInput.Protocol1
    leafs["protocol2"] = transportInput.Protocol2
    leafs["none"] = transportInput.None
    return leafs
}

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetBundleName() string { return "cisco_ios_xr" }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetYangName() string { return "transport-input" }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) SetParent(parent types.Entity) { transportInput.parent = parent }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetParent() types.Entity { return transportInput.parent }

func (transportInput *Tty_VtyLines_VtyLine_Configuration_ConnectionConfiguration_TransportInput) GetParentYangName() string { return "connection-configuration" }

// Tty_VtyLines_VtyLine_Sessions
// Outgoing sessions
type Tty_VtyLines_VtyLine_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of outgoing sessions. The type is slice of
    // Tty_VtyLines_VtyLine_Sessions_OutgoingConnection.
    OutgoingConnection []Tty_VtyLines_VtyLine_Sessions_OutgoingConnection
}

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *Tty_VtyLines_VtyLine_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetGoName(yname string) string {
    if yname == "outgoing-connection" { return "OutgoingConnection" }
    return ""
}

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetSegmentPath() string {
    return "Cisco-IOS-XR-tty-management-oper:sessions"
}

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "outgoing-connection" {
        for _, c := range sessions.OutgoingConnection {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tty_VtyLines_VtyLine_Sessions_OutgoingConnection{}
        sessions.OutgoingConnection = append(sessions.OutgoingConnection, child)
        return &sessions.OutgoingConnection[len(sessions.OutgoingConnection)-1]
    }
    return nil
}

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.OutgoingConnection {
        children[sessions.OutgoingConnection[i].GetSegmentPath()] = &sessions.OutgoingConnection[i]
    }
    return children
}

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetYangName() string { return "sessions" }

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *Tty_VtyLines_VtyLine_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *Tty_VtyLines_VtyLine_Sessions) GetParentYangName() string { return "vty-line" }

// Tty_VtyLines_VtyLine_Sessions_OutgoingConnection
// List of outgoing sessions
type Tty_VtyLines_VtyLine_Sessions_OutgoingConnection struct {
    parent types.Entity
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

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetFilter() yfilter.YFilter { return outgoingConnection.YFilter }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) SetFilter(yf yfilter.YFilter) { outgoingConnection.YFilter = yf }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetGoName(yname string) string {
    if yname == "connection-id" { return "ConnectionId" }
    if yname == "host-name" { return "HostName" }
    if yname == "transport-protocol" { return "TransportProtocol" }
    if yname == "is-last-active-session" { return "IsLastActiveSession" }
    if yname == "idle-time" { return "IdleTime" }
    if yname == "host-address" { return "HostAddress" }
    return ""
}

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetSegmentPath() string {
    return "outgoing-connection"
}

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host-address" {
        return &outgoingConnection.HostAddress
    }
    return nil
}

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["host-address"] = &outgoingConnection.HostAddress
    return children
}

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["connection-id"] = outgoingConnection.ConnectionId
    leafs["host-name"] = outgoingConnection.HostName
    leafs["transport-protocol"] = outgoingConnection.TransportProtocol
    leafs["is-last-active-session"] = outgoingConnection.IsLastActiveSession
    leafs["idle-time"] = outgoingConnection.IdleTime
    return leafs
}

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetBundleName() string { return "cisco_ios_xr" }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetYangName() string { return "outgoing-connection" }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) SetParent(parent types.Entity) { outgoingConnection.parent = parent }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetParent() types.Entity { return outgoingConnection.parent }

func (outgoingConnection *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection) GetParentYangName() string { return "sessions" }

// Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress
// Host address
type Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress struct {
    parent types.Entity
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

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetFilter() yfilter.YFilter { return hostAddress.YFilter }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) SetFilter(yf yfilter.YFilter) { hostAddress.YFilter = yf }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetSegmentPath() string {
    return "host-address"
}

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = hostAddress.AfName
    leafs["ipv4-address"] = hostAddress.Ipv4Address
    leafs["ipv6-address"] = hostAddress.Ipv6Address
    return leafs
}

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetBundleName() string { return "cisco_ios_xr" }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetYangName() string { return "host-address" }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) SetParent(parent types.Entity) { hostAddress.parent = parent }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetParent() types.Entity { return hostAddress.parent }

func (hostAddress *Tty_VtyLines_VtyLine_Sessions_OutgoingConnection_HostAddress) GetParentYangName() string { return "outgoing-connection" }

// Tty_AuxiliaryNodes
// List of Nodes attached with an auxiliary line
type Tty_AuxiliaryNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Line configuration on a node. The type is slice of
    // Tty_AuxiliaryNodes_AuxiliaryNode.
    AuxiliaryNode []Tty_AuxiliaryNodes_AuxiliaryNode
}

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetFilter() yfilter.YFilter { return auxiliaryNodes.YFilter }

func (auxiliaryNodes *Tty_AuxiliaryNodes) SetFilter(yf yfilter.YFilter) { auxiliaryNodes.YFilter = yf }

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetGoName(yname string) string {
    if yname == "auxiliary-node" { return "AuxiliaryNode" }
    return ""
}

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetSegmentPath() string {
    return "auxiliary-nodes"
}

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auxiliary-node" {
        for _, c := range auxiliaryNodes.AuxiliaryNode {
            if auxiliaryNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tty_AuxiliaryNodes_AuxiliaryNode{}
        auxiliaryNodes.AuxiliaryNode = append(auxiliaryNodes.AuxiliaryNode, child)
        return &auxiliaryNodes.AuxiliaryNode[len(auxiliaryNodes.AuxiliaryNode)-1]
    }
    return nil
}

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range auxiliaryNodes.AuxiliaryNode {
        children[auxiliaryNodes.AuxiliaryNode[i].GetSegmentPath()] = &auxiliaryNodes.AuxiliaryNode[i]
    }
    return children
}

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetBundleName() string { return "cisco_ios_xr" }

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetYangName() string { return "auxiliary-nodes" }

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auxiliaryNodes *Tty_AuxiliaryNodes) SetParent(parent types.Entity) { auxiliaryNodes.parent = parent }

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetParent() types.Entity { return auxiliaryNodes.parent }

func (auxiliaryNodes *Tty_AuxiliaryNodes) GetParentYangName() string { return "tty" }

// Tty_AuxiliaryNodes_AuxiliaryNode
// Line configuration on a node
type Tty_AuxiliaryNodes_AuxiliaryNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Id interface{}

    // Auxiliary line.
    AuxiliaryLine Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine
}

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetFilter() yfilter.YFilter { return auxiliaryNode.YFilter }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) SetFilter(yf yfilter.YFilter) { auxiliaryNode.YFilter = yf }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "auxiliary-line" { return "AuxiliaryLine" }
    return ""
}

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetSegmentPath() string {
    return "auxiliary-node" + "[id='" + fmt.Sprintf("%v", auxiliaryNode.Id) + "']"
}

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auxiliary-line" {
        return &auxiliaryNode.AuxiliaryLine
    }
    return nil
}

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auxiliary-line"] = &auxiliaryNode.AuxiliaryLine
    return children
}

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = auxiliaryNode.Id
    return leafs
}

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetBundleName() string { return "cisco_ios_xr" }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetYangName() string { return "auxiliary-node" }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) SetParent(parent types.Entity) { auxiliaryNode.parent = parent }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetParent() types.Entity { return auxiliaryNode.parent }

func (auxiliaryNode *Tty_AuxiliaryNodes_AuxiliaryNode) GetParentYangName() string { return "auxiliary-nodes" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine
// Auxiliary line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics of the auxiliary line.
    AuxiliaryStatistics Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics

    // Line state information.
    State Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State

    // Configuration information of the line.
    Configuration Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration
}

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetFilter() yfilter.YFilter { return auxiliaryLine.YFilter }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) SetFilter(yf yfilter.YFilter) { auxiliaryLine.YFilter = yf }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetGoName(yname string) string {
    if yname == "auxiliary-statistics" { return "AuxiliaryStatistics" }
    if yname == "state" { return "State" }
    if yname == "configuration" { return "Configuration" }
    return ""
}

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetSegmentPath() string {
    return "auxiliary-line"
}

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auxiliary-statistics" {
        return &auxiliaryLine.AuxiliaryStatistics
    }
    if childYangName == "state" {
        return &auxiliaryLine.State
    }
    if childYangName == "configuration" {
        return &auxiliaryLine.Configuration
    }
    return nil
}

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auxiliary-statistics"] = &auxiliaryLine.AuxiliaryStatistics
    children["state"] = &auxiliaryLine.State
    children["configuration"] = &auxiliaryLine.Configuration
    return children
}

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetBundleName() string { return "cisco_ios_xr" }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetYangName() string { return "auxiliary-line" }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) SetParent(parent types.Entity) { auxiliaryLine.parent = parent }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetParent() types.Entity { return auxiliaryLine.parent }

func (auxiliaryLine *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine) GetParentYangName() string { return "auxiliary-node" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics
// Statistics of the auxiliary line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics struct {
    parent types.Entity
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

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetFilter() yfilter.YFilter { return auxiliaryStatistics.YFilter }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) SetFilter(yf yfilter.YFilter) { auxiliaryStatistics.YFilter = yf }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetGoName(yname string) string {
    if yname == "rs232" { return "Rs232" }
    if yname == "general-statistics" { return "GeneralStatistics" }
    if yname == "exec" { return "Exec" }
    if yname == "aaa" { return "Aaa" }
    return ""
}

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetSegmentPath() string {
    return "auxiliary-statistics"
}

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rs232" {
        return &auxiliaryStatistics.Rs232
    }
    if childYangName == "general-statistics" {
        return &auxiliaryStatistics.GeneralStatistics
    }
    if childYangName == "exec" {
        return &auxiliaryStatistics.Exec
    }
    if childYangName == "aaa" {
        return &auxiliaryStatistics.Aaa
    }
    return nil
}

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rs232"] = &auxiliaryStatistics.Rs232
    children["general-statistics"] = &auxiliaryStatistics.GeneralStatistics
    children["exec"] = &auxiliaryStatistics.Exec
    children["aaa"] = &auxiliaryStatistics.Aaa
    return children
}

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetYangName() string { return "auxiliary-statistics" }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) SetParent(parent types.Entity) { auxiliaryStatistics.parent = parent }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetParent() types.Entity { return auxiliaryStatistics.parent }

func (auxiliaryStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics) GetParentYangName() string { return "auxiliary-line" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232
// RS232 statistics of console line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232 struct {
    parent types.Entity
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

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetFilter() yfilter.YFilter { return rs232.YFilter }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) SetFilter(yf yfilter.YFilter) { rs232.YFilter = yf }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetGoName(yname string) string {
    if yname == "data-bits" { return "DataBits" }
    if yname == "exec-disabled" { return "ExecDisabled" }
    if yname == "hardware-flow-control-status" { return "HardwareFlowControlStatus" }
    if yname == "parity-status" { return "ParityStatus" }
    if yname == "baud-rate" { return "BaudRate" }
    if yname == "stop-bits" { return "StopBits" }
    if yname == "overrun-error-count" { return "OverrunErrorCount" }
    if yname == "framing-error-count" { return "FramingErrorCount" }
    if yname == "parity-error-count" { return "ParityErrorCount" }
    return ""
}

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetSegmentPath() string {
    return "rs232"
}

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-bits"] = rs232.DataBits
    leafs["exec-disabled"] = rs232.ExecDisabled
    leafs["hardware-flow-control-status"] = rs232.HardwareFlowControlStatus
    leafs["parity-status"] = rs232.ParityStatus
    leafs["baud-rate"] = rs232.BaudRate
    leafs["stop-bits"] = rs232.StopBits
    leafs["overrun-error-count"] = rs232.OverrunErrorCount
    leafs["framing-error-count"] = rs232.FramingErrorCount
    leafs["parity-error-count"] = rs232.ParityErrorCount
    return leafs
}

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetBundleName() string { return "cisco_ios_xr" }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetYangName() string { return "rs232" }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) SetParent(parent types.Entity) { rs232.parent = parent }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetParent() types.Entity { return rs232.parent }

func (rs232 *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Rs232) GetParentYangName() string { return "auxiliary-statistics" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics
// General statistics of line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics struct {
    parent types.Entity
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

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetFilter() yfilter.YFilter { return generalStatistics.YFilter }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) SetFilter(yf yfilter.YFilter) { generalStatistics.YFilter = yf }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetGoName(yname string) string {
    if yname == "terminal-length" { return "TerminalLength" }
    if yname == "terminal-width" { return "TerminalWidth" }
    if yname == "async-interface" { return "AsyncInterface" }
    if yname == "flow-control-start-character" { return "FlowControlStartCharacter" }
    if yname == "flow-control-stop-character" { return "FlowControlStopCharacter" }
    if yname == "domain-lookup-enabled" { return "DomainLookupEnabled" }
    if yname == "motd-banner-enabled" { return "MotdBannerEnabled" }
    if yname == "private-flag" { return "PrivateFlag" }
    if yname == "terminal-type" { return "TerminalType" }
    if yname == "absolute-timeout" { return "AbsoluteTimeout" }
    if yname == "idle-time" { return "IdleTime" }
    return ""
}

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetSegmentPath() string {
    return "general-statistics"
}

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminal-length"] = generalStatistics.TerminalLength
    leafs["terminal-width"] = generalStatistics.TerminalWidth
    leafs["async-interface"] = generalStatistics.AsyncInterface
    leafs["flow-control-start-character"] = generalStatistics.FlowControlStartCharacter
    leafs["flow-control-stop-character"] = generalStatistics.FlowControlStopCharacter
    leafs["domain-lookup-enabled"] = generalStatistics.DomainLookupEnabled
    leafs["motd-banner-enabled"] = generalStatistics.MotdBannerEnabled
    leafs["private-flag"] = generalStatistics.PrivateFlag
    leafs["terminal-type"] = generalStatistics.TerminalType
    leafs["absolute-timeout"] = generalStatistics.AbsoluteTimeout
    leafs["idle-time"] = generalStatistics.IdleTime
    return leafs
}

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetYangName() string { return "general-statistics" }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) SetParent(parent types.Entity) { generalStatistics.parent = parent }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetParent() types.Entity { return generalStatistics.parent }

func (generalStatistics *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_GeneralStatistics) GetParentYangName() string { return "auxiliary-statistics" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec
// Exec related statistics
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifies whether timestamp is enabled or not. The type is bool.
    TimeStampEnabled interface{}
}

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetFilter() yfilter.YFilter { return exec.YFilter }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) SetFilter(yf yfilter.YFilter) { exec.YFilter = yf }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetGoName(yname string) string {
    if yname == "time-stamp-enabled" { return "TimeStampEnabled" }
    return ""
}

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetSegmentPath() string {
    return "exec"
}

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-stamp-enabled"] = exec.TimeStampEnabled
    return leafs
}

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetBundleName() string { return "cisco_ios_xr" }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetYangName() string { return "exec" }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) SetParent(parent types.Entity) { exec.parent = parent }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetParent() types.Entity { return exec.parent }

func (exec *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Exec) GetParentYangName() string { return "auxiliary-statistics" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa
// AAA related statistics
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The authenticated username. The type is string.
    UserName interface{}
}

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetGoName(yname string) string {
    if yname == "user-name" { return "UserName" }
    return ""
}

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetSegmentPath() string {
    return "aaa"
}

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["user-name"] = aaa.UserName
    return leafs
}

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetBundleName() string { return "cisco_ios_xr" }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetYangName() string { return "aaa" }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) SetParent(parent types.Entity) { aaa.parent = parent }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetParent() types.Entity { return aaa.parent }

func (aaa *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_AuxiliaryStatistics_Aaa) GetParentYangName() string { return "auxiliary-statistics" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State
// Line state information
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information related to template applied to the line.
    Template Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template

    // General information.
    General Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General
}

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    if yname == "general" { return "General" }
    return ""
}

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetSegmentPath() string {
    return "state"
}

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        return &state.Template
    }
    if childYangName == "general" {
        return &state.General
    }
    return nil
}

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["template"] = &state.Template
    children["general"] = &state.General
    return children
}

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetYangName() string { return "state" }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetParent() types.Entity { return state.parent }

func (state *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State) GetParentYangName() string { return "auxiliary-line" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template
// Information related to template applied to the
// line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the template. The type is string.
    Name interface{}
}

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetSegmentPath() string {
    return "template"
}

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = template.Name
    return leafs
}

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetYangName() string { return "template" }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetParent() types.Entity { return template.parent }

func (template *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_Template) GetParentYangName() string { return "state" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General
// General information
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // application running of on the tty line. The type is SessionOperation.
    Operation interface{}

    // State of the line. The type is LineState.
    GeneralState interface{}
}

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetFilter() yfilter.YFilter { return general.YFilter }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) SetFilter(yf yfilter.YFilter) { general.YFilter = yf }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetGoName(yname string) string {
    if yname == "operation" { return "Operation" }
    if yname == "general-state" { return "GeneralState" }
    return ""
}

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetSegmentPath() string {
    return "general"
}

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation"] = general.Operation
    leafs["general-state"] = general.GeneralState
    return leafs
}

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetBundleName() string { return "cisco_ios_xr" }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetYangName() string { return "general" }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) SetParent(parent types.Entity) { general.parent = parent }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetParent() types.Entity { return general.parent }

func (general *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_State_General) GetParentYangName() string { return "state" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration
// Configuration information of the line
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Conection configuration information.
    ConnectionConfiguration Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration
}

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetGoName(yname string) string {
    if yname == "connection-configuration" { return "ConnectionConfiguration" }
    return ""
}

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connection-configuration" {
        return &configuration.ConnectionConfiguration
    }
    return nil
}

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connection-configuration"] = &configuration.ConnectionConfiguration
    return children
}

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetYangName() string { return "configuration" }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration) GetParentYangName() string { return "auxiliary-line" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration
// Conection configuration information
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL for outbound traffic. The type is string.
    AclOut interface{}

    // ACL for inbound traffic. The type is string.
    AclIn interface{}

    // Protocols to use when connecting to the terminal server.
    TransportInput Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput
}

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetFilter() yfilter.YFilter { return connectionConfiguration.YFilter }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) SetFilter(yf yfilter.YFilter) { connectionConfiguration.YFilter = yf }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetGoName(yname string) string {
    if yname == "acl-out" { return "AclOut" }
    if yname == "acl-in" { return "AclIn" }
    if yname == "transport-input" { return "TransportInput" }
    return ""
}

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetSegmentPath() string {
    return "connection-configuration"
}

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport-input" {
        return &connectionConfiguration.TransportInput
    }
    return nil
}

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transport-input"] = &connectionConfiguration.TransportInput
    return children
}

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl-out"] = connectionConfiguration.AclOut
    leafs["acl-in"] = connectionConfiguration.AclIn
    return leafs
}

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetYangName() string { return "connection-configuration" }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) SetParent(parent types.Entity) { connectionConfiguration.parent = parent }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetParent() types.Entity { return connectionConfiguration.parent }

func (connectionConfiguration *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration) GetParentYangName() string { return "configuration" }

// Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: -2147483648..2147483647.
    None interface{}
}

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetFilter() yfilter.YFilter { return transportInput.YFilter }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) SetFilter(yf yfilter.YFilter) { transportInput.YFilter = yf }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetGoName(yname string) string {
    if yname == "select" { return "Select" }
    if yname == "protocol1" { return "Protocol1" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "none" { return "None" }
    return ""
}

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetSegmentPath() string {
    return "transport-input"
}

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["select"] = transportInput.Select
    leafs["protocol1"] = transportInput.Protocol1
    leafs["protocol2"] = transportInput.Protocol2
    leafs["none"] = transportInput.None
    return leafs
}

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetBundleName() string { return "cisco_ios_xr" }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetYangName() string { return "transport-input" }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) SetParent(parent types.Entity) { transportInput.parent = parent }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetParent() types.Entity { return transportInput.parent }

func (transportInput *Tty_AuxiliaryNodes_AuxiliaryNode_AuxiliaryLine_Configuration_ConnectionConfiguration_TransportInput) GetParentYangName() string { return "connection-configuration" }

