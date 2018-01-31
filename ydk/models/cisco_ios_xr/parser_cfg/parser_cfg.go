// This module contains a collection of YANG definitions
// for Cisco IOS-XR parser package configuration.
// 
// This module contains definitions
// for the following management objects:
//   parser: Parser configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package parser_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package parser_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-parser-cfg parser}", reflect.TypeOf(Parser{}))
    ydk.RegisterEntity("Cisco-IOS-XR-parser-cfg:parser", reflect.TypeOf(Parser{}))
}

// Parser
// Parser configuration
type Parser struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // indentation tracking.
    Indentation Parser_Indentation

    // Alias for command mapping.
    Alias Parser_Alias

    // cli commands history.
    History Parser_History

    // interactive mode.
    Interactive Parser_Interactive

    // Configure the Interface display order.
    InterfaceDisplay Parser_InterfaceDisplay

    // Ipv4 netmask-format to be configured.
    NetmaskFormat Parser_NetmaskFormat

    // cli configuration services.
    Configuration Parser_Configuration

    // Exit submode when only '!' seen in interactive mode.
    SubmodeExit Parser_SubmodeExit
}

func (parser *Parser) GetFilter() yfilter.YFilter { return parser.YFilter }

func (parser *Parser) SetFilter(yf yfilter.YFilter) { parser.YFilter = yf }

func (parser *Parser) GetGoName(yname string) string {
    if yname == "indentation" { return "Indentation" }
    if yname == "alias" { return "Alias" }
    if yname == "history" { return "History" }
    if yname == "interactive" { return "Interactive" }
    if yname == "interface-display" { return "InterfaceDisplay" }
    if yname == "netmask-format" { return "NetmaskFormat" }
    if yname == "configuration" { return "Configuration" }
    if yname == "submode-exit" { return "SubmodeExit" }
    return ""
}

func (parser *Parser) GetSegmentPath() string {
    return "Cisco-IOS-XR-parser-cfg:parser"
}

func (parser *Parser) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "indentation" {
        return &parser.Indentation
    }
    if childYangName == "alias" {
        return &parser.Alias
    }
    if childYangName == "history" {
        return &parser.History
    }
    if childYangName == "interactive" {
        return &parser.Interactive
    }
    if childYangName == "interface-display" {
        return &parser.InterfaceDisplay
    }
    if childYangName == "netmask-format" {
        return &parser.NetmaskFormat
    }
    if childYangName == "configuration" {
        return &parser.Configuration
    }
    if childYangName == "submode-exit" {
        return &parser.SubmodeExit
    }
    return nil
}

func (parser *Parser) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["indentation"] = &parser.Indentation
    children["alias"] = &parser.Alias
    children["history"] = &parser.History
    children["interactive"] = &parser.Interactive
    children["interface-display"] = &parser.InterfaceDisplay
    children["netmask-format"] = &parser.NetmaskFormat
    children["configuration"] = &parser.Configuration
    children["submode-exit"] = &parser.SubmodeExit
    return children
}

func (parser *Parser) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (parser *Parser) GetBundleName() string { return "cisco_ios_xr" }

func (parser *Parser) GetYangName() string { return "parser" }

func (parser *Parser) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (parser *Parser) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (parser *Parser) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (parser *Parser) SetParent(parent types.Entity) { parser.parent = parent }

func (parser *Parser) GetParent() types.Entity { return parser.parent }

func (parser *Parser) GetParentYangName() string { return "Cisco-IOS-XR-parser-cfg" }

// Parser_Indentation
// indentation tracking
type Parser_Indentation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable the indentation. The type is bool.
    IndentationDisable interface{}
}

func (indentation *Parser_Indentation) GetFilter() yfilter.YFilter { return indentation.YFilter }

func (indentation *Parser_Indentation) SetFilter(yf yfilter.YFilter) { indentation.YFilter = yf }

func (indentation *Parser_Indentation) GetGoName(yname string) string {
    if yname == "indentation-disable" { return "IndentationDisable" }
    return ""
}

func (indentation *Parser_Indentation) GetSegmentPath() string {
    return "indentation"
}

func (indentation *Parser_Indentation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (indentation *Parser_Indentation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (indentation *Parser_Indentation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["indentation-disable"] = indentation.IndentationDisable
    return leafs
}

func (indentation *Parser_Indentation) GetBundleName() string { return "cisco_ios_xr" }

func (indentation *Parser_Indentation) GetYangName() string { return "indentation" }

func (indentation *Parser_Indentation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indentation *Parser_Indentation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indentation *Parser_Indentation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indentation *Parser_Indentation) SetParent(parent types.Entity) { indentation.parent = parent }

func (indentation *Parser_Indentation) GetParent() types.Entity { return indentation.parent }

func (indentation *Parser_Indentation) GetParentYangName() string { return "parser" }

// Parser_Alias
// Alias for command mapping
type Parser_Alias struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exec command alias.
    Execs Parser_Alias_Execs

    // Configuration command alias.
    Configurations Parser_Alias_Configurations

    // Table of all aliases configured.
    Alls Parser_Alias_Alls
}

func (alias *Parser_Alias) GetFilter() yfilter.YFilter { return alias.YFilter }

func (alias *Parser_Alias) SetFilter(yf yfilter.YFilter) { alias.YFilter = yf }

func (alias *Parser_Alias) GetGoName(yname string) string {
    if yname == "execs" { return "Execs" }
    if yname == "configurations" { return "Configurations" }
    if yname == "alls" { return "Alls" }
    return ""
}

func (alias *Parser_Alias) GetSegmentPath() string {
    return "alias"
}

func (alias *Parser_Alias) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "execs" {
        return &alias.Execs
    }
    if childYangName == "configurations" {
        return &alias.Configurations
    }
    if childYangName == "alls" {
        return &alias.Alls
    }
    return nil
}

func (alias *Parser_Alias) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["execs"] = &alias.Execs
    children["configurations"] = &alias.Configurations
    children["alls"] = &alias.Alls
    return children
}

func (alias *Parser_Alias) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alias *Parser_Alias) GetBundleName() string { return "cisco_ios_xr" }

func (alias *Parser_Alias) GetYangName() string { return "alias" }

func (alias *Parser_Alias) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alias *Parser_Alias) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alias *Parser_Alias) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alias *Parser_Alias) SetParent(parent types.Entity) { alias.parent = parent }

func (alias *Parser_Alias) GetParent() types.Entity { return alias.parent }

func (alias *Parser_Alias) GetParentYangName() string { return "parser" }

// Parser_Alias_Execs
// Exec command alias
type Parser_Alias_Execs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exec alias name. The type is slice of Parser_Alias_Execs_Exec.
    Exec []Parser_Alias_Execs_Exec
}

func (execs *Parser_Alias_Execs) GetFilter() yfilter.YFilter { return execs.YFilter }

func (execs *Parser_Alias_Execs) SetFilter(yf yfilter.YFilter) { execs.YFilter = yf }

func (execs *Parser_Alias_Execs) GetGoName(yname string) string {
    if yname == "exec" { return "Exec" }
    return ""
}

func (execs *Parser_Alias_Execs) GetSegmentPath() string {
    return "execs"
}

func (execs *Parser_Alias_Execs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exec" {
        for _, c := range execs.Exec {
            if execs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Parser_Alias_Execs_Exec{}
        execs.Exec = append(execs.Exec, child)
        return &execs.Exec[len(execs.Exec)-1]
    }
    return nil
}

func (execs *Parser_Alias_Execs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range execs.Exec {
        children[execs.Exec[i].GetSegmentPath()] = &execs.Exec[i]
    }
    return children
}

func (execs *Parser_Alias_Execs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (execs *Parser_Alias_Execs) GetBundleName() string { return "cisco_ios_xr" }

func (execs *Parser_Alias_Execs) GetYangName() string { return "execs" }

func (execs *Parser_Alias_Execs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (execs *Parser_Alias_Execs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (execs *Parser_Alias_Execs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (execs *Parser_Alias_Execs) SetParent(parent types.Entity) { execs.parent = parent }

func (execs *Parser_Alias_Execs) GetParent() types.Entity { return execs.parent }

func (execs *Parser_Alias_Execs) GetParentYangName() string { return "alias" }

// Parser_Alias_Execs_Exec
// Exec alias name
type Parser_Alias_Execs_Exec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Exec Alias name. The type is string with length:
    // 1..30.
    Identifier interface{}

    // Aliased exec command. The type is string. This attribute is mandatory.
    IdentifierXr interface{}
}

func (exec *Parser_Alias_Execs_Exec) GetFilter() yfilter.YFilter { return exec.YFilter }

func (exec *Parser_Alias_Execs_Exec) SetFilter(yf yfilter.YFilter) { exec.YFilter = yf }

func (exec *Parser_Alias_Execs_Exec) GetGoName(yname string) string {
    if yname == "identifier" { return "Identifier" }
    if yname == "identifier-xr" { return "IdentifierXr" }
    return ""
}

func (exec *Parser_Alias_Execs_Exec) GetSegmentPath() string {
    return "exec" + "[identifier='" + fmt.Sprintf("%v", exec.Identifier) + "']"
}

func (exec *Parser_Alias_Execs_Exec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exec *Parser_Alias_Execs_Exec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exec *Parser_Alias_Execs_Exec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["identifier"] = exec.Identifier
    leafs["identifier-xr"] = exec.IdentifierXr
    return leafs
}

func (exec *Parser_Alias_Execs_Exec) GetBundleName() string { return "cisco_ios_xr" }

func (exec *Parser_Alias_Execs_Exec) GetYangName() string { return "exec" }

func (exec *Parser_Alias_Execs_Exec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exec *Parser_Alias_Execs_Exec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exec *Parser_Alias_Execs_Exec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exec *Parser_Alias_Execs_Exec) SetParent(parent types.Entity) { exec.parent = parent }

func (exec *Parser_Alias_Execs_Exec) GetParent() types.Entity { return exec.parent }

func (exec *Parser_Alias_Execs_Exec) GetParentYangName() string { return "execs" }

// Parser_Alias_Configurations
// Configuration command alias
type Parser_Alias_Configurations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration Alias name. The type is slice of
    // Parser_Alias_Configurations_Configuration.
    Configuration []Parser_Alias_Configurations_Configuration
}

func (configurations *Parser_Alias_Configurations) GetFilter() yfilter.YFilter { return configurations.YFilter }

func (configurations *Parser_Alias_Configurations) SetFilter(yf yfilter.YFilter) { configurations.YFilter = yf }

func (configurations *Parser_Alias_Configurations) GetGoName(yname string) string {
    if yname == "configuration" { return "Configuration" }
    return ""
}

func (configurations *Parser_Alias_Configurations) GetSegmentPath() string {
    return "configurations"
}

func (configurations *Parser_Alias_Configurations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configuration" {
        for _, c := range configurations.Configuration {
            if configurations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Parser_Alias_Configurations_Configuration{}
        configurations.Configuration = append(configurations.Configuration, child)
        return &configurations.Configuration[len(configurations.Configuration)-1]
    }
    return nil
}

func (configurations *Parser_Alias_Configurations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configurations.Configuration {
        children[configurations.Configuration[i].GetSegmentPath()] = &configurations.Configuration[i]
    }
    return children
}

func (configurations *Parser_Alias_Configurations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configurations *Parser_Alias_Configurations) GetBundleName() string { return "cisco_ios_xr" }

func (configurations *Parser_Alias_Configurations) GetYangName() string { return "configurations" }

func (configurations *Parser_Alias_Configurations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurations *Parser_Alias_Configurations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurations *Parser_Alias_Configurations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurations *Parser_Alias_Configurations) SetParent(parent types.Entity) { configurations.parent = parent }

func (configurations *Parser_Alias_Configurations) GetParent() types.Entity { return configurations.parent }

func (configurations *Parser_Alias_Configurations) GetParentYangName() string { return "alias" }

// Parser_Alias_Configurations_Configuration
// Configuration Alias name
type Parser_Alias_Configurations_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Configuration alias name. The type is string with
    // length: 1..30.
    Identifier interface{}

    // Aliased config command. The type is string. This attribute is mandatory.
    IdentifierXr interface{}
}

func (configuration *Parser_Alias_Configurations_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Parser_Alias_Configurations_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Parser_Alias_Configurations_Configuration) GetGoName(yname string) string {
    if yname == "identifier" { return "Identifier" }
    if yname == "identifier-xr" { return "IdentifierXr" }
    return ""
}

func (configuration *Parser_Alias_Configurations_Configuration) GetSegmentPath() string {
    return "configuration" + "[identifier='" + fmt.Sprintf("%v", configuration.Identifier) + "']"
}

func (configuration *Parser_Alias_Configurations_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuration *Parser_Alias_Configurations_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuration *Parser_Alias_Configurations_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["identifier"] = configuration.Identifier
    leafs["identifier-xr"] = configuration.IdentifierXr
    return leafs
}

func (configuration *Parser_Alias_Configurations_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Parser_Alias_Configurations_Configuration) GetYangName() string { return "configuration" }

func (configuration *Parser_Alias_Configurations_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Parser_Alias_Configurations_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Parser_Alias_Configurations_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Parser_Alias_Configurations_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Parser_Alias_Configurations_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Parser_Alias_Configurations_Configuration) GetParentYangName() string { return "configurations" }

// Parser_Alias_Alls
// Table of all aliases configured
type Parser_Alias_Alls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alias name to command mapping. The type is slice of Parser_Alias_Alls_All.
    All []Parser_Alias_Alls_All
}

func (alls *Parser_Alias_Alls) GetFilter() yfilter.YFilter { return alls.YFilter }

func (alls *Parser_Alias_Alls) SetFilter(yf yfilter.YFilter) { alls.YFilter = yf }

func (alls *Parser_Alias_Alls) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    return ""
}

func (alls *Parser_Alias_Alls) GetSegmentPath() string {
    return "alls"
}

func (alls *Parser_Alias_Alls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all" {
        for _, c := range alls.All {
            if alls.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Parser_Alias_Alls_All{}
        alls.All = append(alls.All, child)
        return &alls.All[len(alls.All)-1]
    }
    return nil
}

func (alls *Parser_Alias_Alls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alls.All {
        children[alls.All[i].GetSegmentPath()] = &alls.All[i]
    }
    return children
}

func (alls *Parser_Alias_Alls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alls *Parser_Alias_Alls) GetBundleName() string { return "cisco_ios_xr" }

func (alls *Parser_Alias_Alls) GetYangName() string { return "alls" }

func (alls *Parser_Alias_Alls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alls *Parser_Alias_Alls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alls *Parser_Alias_Alls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alls *Parser_Alias_Alls) SetParent(parent types.Entity) { alls.parent = parent }

func (alls *Parser_Alias_Alls) GetParent() types.Entity { return alls.parent }

func (alls *Parser_Alias_Alls) GetParentYangName() string { return "alias" }

// Parser_Alias_Alls_All
// Alias name to command mapping
type Parser_Alias_Alls_All struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Alias name. The type is string with length: 1..30.
    Identifier interface{}

    // The actual command. The type is string. This attribute is mandatory.
    IdentifierXr interface{}
}

func (all *Parser_Alias_Alls_All) GetFilter() yfilter.YFilter { return all.YFilter }

func (all *Parser_Alias_Alls_All) SetFilter(yf yfilter.YFilter) { all.YFilter = yf }

func (all *Parser_Alias_Alls_All) GetGoName(yname string) string {
    if yname == "identifier" { return "Identifier" }
    if yname == "identifier-xr" { return "IdentifierXr" }
    return ""
}

func (all *Parser_Alias_Alls_All) GetSegmentPath() string {
    return "all" + "[identifier='" + fmt.Sprintf("%v", all.Identifier) + "']"
}

func (all *Parser_Alias_Alls_All) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (all *Parser_Alias_Alls_All) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (all *Parser_Alias_Alls_All) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["identifier"] = all.Identifier
    leafs["identifier-xr"] = all.IdentifierXr
    return leafs
}

func (all *Parser_Alias_Alls_All) GetBundleName() string { return "cisco_ios_xr" }

func (all *Parser_Alias_Alls_All) GetYangName() string { return "all" }

func (all *Parser_Alias_Alls_All) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (all *Parser_Alias_Alls_All) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (all *Parser_Alias_Alls_All) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (all *Parser_Alias_Alls_All) SetParent(parent types.Entity) { all.parent = parent }

func (all *Parser_Alias_Alls_All) GetParent() types.Entity { return all.parent }

func (all *Parser_Alias_Alls_All) GetParentYangName() string { return "alls" }

// Parser_History
// cli commands history
type Parser_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // maximum number of commands in history. The type is interface{} with range:
    // 1000..5000.
    Size interface{}
}

func (history *Parser_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Parser_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Parser_History) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    return ""
}

func (history *Parser_History) GetSegmentPath() string {
    return "history"
}

func (history *Parser_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (history *Parser_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (history *Parser_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = history.Size
    return leafs
}

func (history *Parser_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Parser_History) GetYangName() string { return "history" }

func (history *Parser_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Parser_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Parser_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Parser_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Parser_History) GetParent() types.Entity { return history.parent }

func (history *Parser_History) GetParentYangName() string { return "parser" }

// Parser_Interactive
// interactive mode
type Parser_Interactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable interactive mode. The type is bool.
    InteractiveDisable interface{}
}

func (interactive *Parser_Interactive) GetFilter() yfilter.YFilter { return interactive.YFilter }

func (interactive *Parser_Interactive) SetFilter(yf yfilter.YFilter) { interactive.YFilter = yf }

func (interactive *Parser_Interactive) GetGoName(yname string) string {
    if yname == "interactive-disable" { return "InteractiveDisable" }
    return ""
}

func (interactive *Parser_Interactive) GetSegmentPath() string {
    return "interactive"
}

func (interactive *Parser_Interactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interactive *Parser_Interactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interactive *Parser_Interactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interactive-disable"] = interactive.InteractiveDisable
    return leafs
}

func (interactive *Parser_Interactive) GetBundleName() string { return "cisco_ios_xr" }

func (interactive *Parser_Interactive) GetYangName() string { return "interactive" }

func (interactive *Parser_Interactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interactive *Parser_Interactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interactive *Parser_Interactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interactive *Parser_Interactive) SetParent(parent types.Entity) { interactive.parent = parent }

func (interactive *Parser_Interactive) GetParent() types.Entity { return interactive.parent }

func (interactive *Parser_Interactive) GetParentYangName() string { return "parser" }

// Parser_InterfaceDisplay
// Configure the Interface display order
type Parser_InterfaceDisplay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Interface display order as slot order. The type is bool.
    SlotOrder interface{}
}

func (interfaceDisplay *Parser_InterfaceDisplay) GetFilter() yfilter.YFilter { return interfaceDisplay.YFilter }

func (interfaceDisplay *Parser_InterfaceDisplay) SetFilter(yf yfilter.YFilter) { interfaceDisplay.YFilter = yf }

func (interfaceDisplay *Parser_InterfaceDisplay) GetGoName(yname string) string {
    if yname == "slot-order" { return "SlotOrder" }
    return ""
}

func (interfaceDisplay *Parser_InterfaceDisplay) GetSegmentPath() string {
    return "interface-display"
}

func (interfaceDisplay *Parser_InterfaceDisplay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceDisplay *Parser_InterfaceDisplay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceDisplay *Parser_InterfaceDisplay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot-order"] = interfaceDisplay.SlotOrder
    return leafs
}

func (interfaceDisplay *Parser_InterfaceDisplay) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceDisplay *Parser_InterfaceDisplay) GetYangName() string { return "interface-display" }

func (interfaceDisplay *Parser_InterfaceDisplay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceDisplay *Parser_InterfaceDisplay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceDisplay *Parser_InterfaceDisplay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceDisplay *Parser_InterfaceDisplay) SetParent(parent types.Entity) { interfaceDisplay.parent = parent }

func (interfaceDisplay *Parser_InterfaceDisplay) GetParent() types.Entity { return interfaceDisplay.parent }

func (interfaceDisplay *Parser_InterfaceDisplay) GetParentYangName() string { return "parser" }

// Parser_NetmaskFormat
// Ipv4 netmask-format to be configured
type Parser_NetmaskFormat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ipv4 netmask-format as bit-count. The type is bool.
    BitCount interface{}
}

func (netmaskFormat *Parser_NetmaskFormat) GetFilter() yfilter.YFilter { return netmaskFormat.YFilter }

func (netmaskFormat *Parser_NetmaskFormat) SetFilter(yf yfilter.YFilter) { netmaskFormat.YFilter = yf }

func (netmaskFormat *Parser_NetmaskFormat) GetGoName(yname string) string {
    if yname == "bit-count" { return "BitCount" }
    return ""
}

func (netmaskFormat *Parser_NetmaskFormat) GetSegmentPath() string {
    return "netmask-format"
}

func (netmaskFormat *Parser_NetmaskFormat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (netmaskFormat *Parser_NetmaskFormat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (netmaskFormat *Parser_NetmaskFormat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bit-count"] = netmaskFormat.BitCount
    return leafs
}

func (netmaskFormat *Parser_NetmaskFormat) GetBundleName() string { return "cisco_ios_xr" }

func (netmaskFormat *Parser_NetmaskFormat) GetYangName() string { return "netmask-format" }

func (netmaskFormat *Parser_NetmaskFormat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netmaskFormat *Parser_NetmaskFormat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netmaskFormat *Parser_NetmaskFormat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netmaskFormat *Parser_NetmaskFormat) SetParent(parent types.Entity) { netmaskFormat.parent = parent }

func (netmaskFormat *Parser_NetmaskFormat) GetParent() types.Entity { return netmaskFormat.parent }

func (netmaskFormat *Parser_NetmaskFormat) GetParentYangName() string { return "parser" }

// Parser_Configuration
// cli configuration services
type Parser_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable for read-only access users.
    Disable Parser_Configuration_Disable
}

func (configuration *Parser_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Parser_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Parser_Configuration) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (configuration *Parser_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Parser_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "disable" {
        return &configuration.Disable
    }
    return nil
}

func (configuration *Parser_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["disable"] = &configuration.Disable
    return children
}

func (configuration *Parser_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configuration *Parser_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Parser_Configuration) GetYangName() string { return "configuration" }

func (configuration *Parser_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Parser_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Parser_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Parser_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Parser_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Parser_Configuration) GetParentYangName() string { return "parser" }

// Parser_Configuration_Disable
// disable for read-only access users
type Parser_Configuration_Disable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable config mode for usergroup. The type is string.
    Usergroup interface{}
}

func (disable *Parser_Configuration_Disable) GetFilter() yfilter.YFilter { return disable.YFilter }

func (disable *Parser_Configuration_Disable) SetFilter(yf yfilter.YFilter) { disable.YFilter = yf }

func (disable *Parser_Configuration_Disable) GetGoName(yname string) string {
    if yname == "usergroup" { return "Usergroup" }
    return ""
}

func (disable *Parser_Configuration_Disable) GetSegmentPath() string {
    return "disable"
}

func (disable *Parser_Configuration_Disable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (disable *Parser_Configuration_Disable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (disable *Parser_Configuration_Disable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["usergroup"] = disable.Usergroup
    return leafs
}

func (disable *Parser_Configuration_Disable) GetBundleName() string { return "cisco_ios_xr" }

func (disable *Parser_Configuration_Disable) GetYangName() string { return "disable" }

func (disable *Parser_Configuration_Disable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (disable *Parser_Configuration_Disable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (disable *Parser_Configuration_Disable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (disable *Parser_Configuration_Disable) SetParent(parent types.Entity) { disable.parent = parent }

func (disable *Parser_Configuration_Disable) GetParent() types.Entity { return disable.parent }

func (disable *Parser_Configuration_Disable) GetParentYangName() string { return "configuration" }

// Parser_SubmodeExit
// Exit submode when only '!' seen in interactive
// mode
type Parser_SubmodeExit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable the feature. The type is bool.
    Enable interface{}
}

func (submodeExit *Parser_SubmodeExit) GetFilter() yfilter.YFilter { return submodeExit.YFilter }

func (submodeExit *Parser_SubmodeExit) SetFilter(yf yfilter.YFilter) { submodeExit.YFilter = yf }

func (submodeExit *Parser_SubmodeExit) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (submodeExit *Parser_SubmodeExit) GetSegmentPath() string {
    return "submode-exit"
}

func (submodeExit *Parser_SubmodeExit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (submodeExit *Parser_SubmodeExit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (submodeExit *Parser_SubmodeExit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = submodeExit.Enable
    return leafs
}

func (submodeExit *Parser_SubmodeExit) GetBundleName() string { return "cisco_ios_xr" }

func (submodeExit *Parser_SubmodeExit) GetYangName() string { return "submode-exit" }

func (submodeExit *Parser_SubmodeExit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (submodeExit *Parser_SubmodeExit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (submodeExit *Parser_SubmodeExit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (submodeExit *Parser_SubmodeExit) SetParent(parent types.Entity) { submodeExit.parent = parent }

func (submodeExit *Parser_SubmodeExit) GetParent() types.Entity { return submodeExit.parent }

func (submodeExit *Parser_SubmodeExit) GetParentYangName() string { return "parser" }

