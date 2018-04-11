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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // indentation tracking.
    Indentation Parser_Indentation

    // Alias for command mapping.
    Alias Parser_Alias

    // cli commands history.
    History Parser_History

    // interactive mode.
    Interactive Parser_Interactive

    // Configuration to disable sysadmin login banner.
    SysadminLoginBanner Parser_SysadminLoginBanner

    // Configure the Interface display order.
    InterfaceDisplay Parser_InterfaceDisplay

    // Ipv4 netmask-format to be configured.
    NetmaskFormat Parser_NetmaskFormat

    // cli configuration services.
    Configuration Parser_Configuration

    // Exit submode when only '!' seen in interactive mode.
    SubmodeExit Parser_SubmodeExit
}

func (parser *Parser) GetEntityData() *types.CommonEntityData {
    parser.EntityData.YFilter = parser.YFilter
    parser.EntityData.YangName = "parser"
    parser.EntityData.BundleName = "cisco_ios_xr"
    parser.EntityData.ParentYangName = "Cisco-IOS-XR-parser-cfg"
    parser.EntityData.SegmentPath = "Cisco-IOS-XR-parser-cfg:parser"
    parser.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parser.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parser.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parser.EntityData.Children = make(map[string]types.YChild)
    parser.EntityData.Children["indentation"] = types.YChild{"Indentation", &parser.Indentation}
    parser.EntityData.Children["alias"] = types.YChild{"Alias", &parser.Alias}
    parser.EntityData.Children["history"] = types.YChild{"History", &parser.History}
    parser.EntityData.Children["interactive"] = types.YChild{"Interactive", &parser.Interactive}
    parser.EntityData.Children["sysadmin-login-banner"] = types.YChild{"SysadminLoginBanner", &parser.SysadminLoginBanner}
    parser.EntityData.Children["interface-display"] = types.YChild{"InterfaceDisplay", &parser.InterfaceDisplay}
    parser.EntityData.Children["netmask-format"] = types.YChild{"NetmaskFormat", &parser.NetmaskFormat}
    parser.EntityData.Children["configuration"] = types.YChild{"Configuration", &parser.Configuration}
    parser.EntityData.Children["submode-exit"] = types.YChild{"SubmodeExit", &parser.SubmodeExit}
    parser.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(parser.EntityData)
}

// Parser_Indentation
// indentation tracking
type Parser_Indentation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable the indentation. The type is bool.
    IndentationDisable interface{}
}

func (indentation *Parser_Indentation) GetEntityData() *types.CommonEntityData {
    indentation.EntityData.YFilter = indentation.YFilter
    indentation.EntityData.YangName = "indentation"
    indentation.EntityData.BundleName = "cisco_ios_xr"
    indentation.EntityData.ParentYangName = "parser"
    indentation.EntityData.SegmentPath = "indentation"
    indentation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indentation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indentation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indentation.EntityData.Children = make(map[string]types.YChild)
    indentation.EntityData.Leafs = make(map[string]types.YLeaf)
    indentation.EntityData.Leafs["indentation-disable"] = types.YLeaf{"IndentationDisable", indentation.IndentationDisable}
    return &(indentation.EntityData)
}

// Parser_Alias
// Alias for command mapping
type Parser_Alias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exec command alias.
    Execs Parser_Alias_Execs

    // Configuration command alias.
    Configurations Parser_Alias_Configurations

    // Table of all aliases configured.
    Alls Parser_Alias_Alls
}

func (alias *Parser_Alias) GetEntityData() *types.CommonEntityData {
    alias.EntityData.YFilter = alias.YFilter
    alias.EntityData.YangName = "alias"
    alias.EntityData.BundleName = "cisco_ios_xr"
    alias.EntityData.ParentYangName = "parser"
    alias.EntityData.SegmentPath = "alias"
    alias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alias.EntityData.Children = make(map[string]types.YChild)
    alias.EntityData.Children["execs"] = types.YChild{"Execs", &alias.Execs}
    alias.EntityData.Children["configurations"] = types.YChild{"Configurations", &alias.Configurations}
    alias.EntityData.Children["alls"] = types.YChild{"Alls", &alias.Alls}
    alias.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alias.EntityData)
}

// Parser_Alias_Execs
// Exec command alias
type Parser_Alias_Execs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exec alias name. The type is slice of Parser_Alias_Execs_Exec.
    Exec []Parser_Alias_Execs_Exec
}

func (execs *Parser_Alias_Execs) GetEntityData() *types.CommonEntityData {
    execs.EntityData.YFilter = execs.YFilter
    execs.EntityData.YangName = "execs"
    execs.EntityData.BundleName = "cisco_ios_xr"
    execs.EntityData.ParentYangName = "alias"
    execs.EntityData.SegmentPath = "execs"
    execs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    execs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    execs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    execs.EntityData.Children = make(map[string]types.YChild)
    execs.EntityData.Children["exec"] = types.YChild{"Exec", nil}
    for i := range execs.Exec {
        execs.EntityData.Children[types.GetSegmentPath(&execs.Exec[i])] = types.YChild{"Exec", &execs.Exec[i]}
    }
    execs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(execs.EntityData)
}

// Parser_Alias_Execs_Exec
// Exec alias name
type Parser_Alias_Execs_Exec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Exec Alias name. The type is string with length:
    // 1..30.
    Identifier interface{}

    // Aliased exec command. The type is string. This attribute is mandatory.
    IdentifierXr interface{}
}

func (exec *Parser_Alias_Execs_Exec) GetEntityData() *types.CommonEntityData {
    exec.EntityData.YFilter = exec.YFilter
    exec.EntityData.YangName = "exec"
    exec.EntityData.BundleName = "cisco_ios_xr"
    exec.EntityData.ParentYangName = "execs"
    exec.EntityData.SegmentPath = "exec" + "[identifier='" + fmt.Sprintf("%v", exec.Identifier) + "']"
    exec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exec.EntityData.Children = make(map[string]types.YChild)
    exec.EntityData.Leafs = make(map[string]types.YLeaf)
    exec.EntityData.Leafs["identifier"] = types.YLeaf{"Identifier", exec.Identifier}
    exec.EntityData.Leafs["identifier-xr"] = types.YLeaf{"IdentifierXr", exec.IdentifierXr}
    return &(exec.EntityData)
}

// Parser_Alias_Configurations
// Configuration command alias
type Parser_Alias_Configurations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration Alias name. The type is slice of
    // Parser_Alias_Configurations_Configuration.
    Configuration []Parser_Alias_Configurations_Configuration
}

func (configurations *Parser_Alias_Configurations) GetEntityData() *types.CommonEntityData {
    configurations.EntityData.YFilter = configurations.YFilter
    configurations.EntityData.YangName = "configurations"
    configurations.EntityData.BundleName = "cisco_ios_xr"
    configurations.EntityData.ParentYangName = "alias"
    configurations.EntityData.SegmentPath = "configurations"
    configurations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurations.EntityData.Children = make(map[string]types.YChild)
    configurations.EntityData.Children["configuration"] = types.YChild{"Configuration", nil}
    for i := range configurations.Configuration {
        configurations.EntityData.Children[types.GetSegmentPath(&configurations.Configuration[i])] = types.YChild{"Configuration", &configurations.Configuration[i]}
    }
    configurations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configurations.EntityData)
}

// Parser_Alias_Configurations_Configuration
// Configuration Alias name
type Parser_Alias_Configurations_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Configuration alias name. The type is string with
    // length: 1..30.
    Identifier interface{}

    // Aliased config command. The type is string. This attribute is mandatory.
    IdentifierXr interface{}
}

func (configuration *Parser_Alias_Configurations_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "configurations"
    configuration.EntityData.SegmentPath = "configuration" + "[identifier='" + fmt.Sprintf("%v", configuration.Identifier) + "']"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    configuration.EntityData.Leafs["identifier"] = types.YLeaf{"Identifier", configuration.Identifier}
    configuration.EntityData.Leafs["identifier-xr"] = types.YLeaf{"IdentifierXr", configuration.IdentifierXr}
    return &(configuration.EntityData)
}

// Parser_Alias_Alls
// Table of all aliases configured
type Parser_Alias_Alls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alias name to command mapping. The type is slice of Parser_Alias_Alls_All.
    All []Parser_Alias_Alls_All
}

func (alls *Parser_Alias_Alls) GetEntityData() *types.CommonEntityData {
    alls.EntityData.YFilter = alls.YFilter
    alls.EntityData.YangName = "alls"
    alls.EntityData.BundleName = "cisco_ios_xr"
    alls.EntityData.ParentYangName = "alias"
    alls.EntityData.SegmentPath = "alls"
    alls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alls.EntityData.Children = make(map[string]types.YChild)
    alls.EntityData.Children["all"] = types.YChild{"All", nil}
    for i := range alls.All {
        alls.EntityData.Children[types.GetSegmentPath(&alls.All[i])] = types.YChild{"All", &alls.All[i]}
    }
    alls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alls.EntityData)
}

// Parser_Alias_Alls_All
// Alias name to command mapping
type Parser_Alias_Alls_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Alias name. The type is string with length: 1..30.
    Identifier interface{}

    // The actual command. The type is string. This attribute is mandatory.
    IdentifierXr interface{}
}

func (all *Parser_Alias_Alls_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "alls"
    all.EntityData.SegmentPath = "all" + "[identifier='" + fmt.Sprintf("%v", all.Identifier) + "']"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    all.EntityData.Leafs["identifier"] = types.YLeaf{"Identifier", all.Identifier}
    all.EntityData.Leafs["identifier-xr"] = types.YLeaf{"IdentifierXr", all.IdentifierXr}
    return &(all.EntityData)
}

// Parser_History
// cli commands history
type Parser_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // maximum number of commands in history. The type is interface{} with range:
    // 1000..5000.
    Size interface{}
}

func (history *Parser_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "parser"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["size"] = types.YLeaf{"Size", history.Size}
    return &(history.EntityData)
}

// Parser_Interactive
// interactive mode
type Parser_Interactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable interactive mode. The type is bool.
    InteractiveDisable interface{}
}

func (interactive *Parser_Interactive) GetEntityData() *types.CommonEntityData {
    interactive.EntityData.YFilter = interactive.YFilter
    interactive.EntityData.YangName = "interactive"
    interactive.EntityData.BundleName = "cisco_ios_xr"
    interactive.EntityData.ParentYangName = "parser"
    interactive.EntityData.SegmentPath = "interactive"
    interactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interactive.EntityData.Children = make(map[string]types.YChild)
    interactive.EntityData.Leafs = make(map[string]types.YLeaf)
    interactive.EntityData.Leafs["interactive-disable"] = types.YLeaf{"InteractiveDisable", interactive.InteractiveDisable}
    return &(interactive.EntityData)
}

// Parser_SysadminLoginBanner
// Configuration to disable sysadmin login banner
type Parser_SysadminLoginBanner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable sysadmin login banner. The type is bool.
    SysadminLoginBannerDisable interface{}
}

func (sysadminLoginBanner *Parser_SysadminLoginBanner) GetEntityData() *types.CommonEntityData {
    sysadminLoginBanner.EntityData.YFilter = sysadminLoginBanner.YFilter
    sysadminLoginBanner.EntityData.YangName = "sysadmin-login-banner"
    sysadminLoginBanner.EntityData.BundleName = "cisco_ios_xr"
    sysadminLoginBanner.EntityData.ParentYangName = "parser"
    sysadminLoginBanner.EntityData.SegmentPath = "sysadmin-login-banner"
    sysadminLoginBanner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysadminLoginBanner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysadminLoginBanner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysadminLoginBanner.EntityData.Children = make(map[string]types.YChild)
    sysadminLoginBanner.EntityData.Leafs = make(map[string]types.YLeaf)
    sysadminLoginBanner.EntityData.Leafs["sysadmin-login-banner-disable"] = types.YLeaf{"SysadminLoginBannerDisable", sysadminLoginBanner.SysadminLoginBannerDisable}
    return &(sysadminLoginBanner.EntityData)
}

// Parser_InterfaceDisplay
// Configure the Interface display order
type Parser_InterfaceDisplay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Interface display order as slot order. The type is bool.
    SlotOrder interface{}
}

func (interfaceDisplay *Parser_InterfaceDisplay) GetEntityData() *types.CommonEntityData {
    interfaceDisplay.EntityData.YFilter = interfaceDisplay.YFilter
    interfaceDisplay.EntityData.YangName = "interface-display"
    interfaceDisplay.EntityData.BundleName = "cisco_ios_xr"
    interfaceDisplay.EntityData.ParentYangName = "parser"
    interfaceDisplay.EntityData.SegmentPath = "interface-display"
    interfaceDisplay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDisplay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDisplay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDisplay.EntityData.Children = make(map[string]types.YChild)
    interfaceDisplay.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceDisplay.EntityData.Leafs["slot-order"] = types.YLeaf{"SlotOrder", interfaceDisplay.SlotOrder}
    return &(interfaceDisplay.EntityData)
}

// Parser_NetmaskFormat
// Ipv4 netmask-format to be configured
type Parser_NetmaskFormat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ipv4 netmask-format as bit-count. The type is bool.
    BitCount interface{}
}

func (netmaskFormat *Parser_NetmaskFormat) GetEntityData() *types.CommonEntityData {
    netmaskFormat.EntityData.YFilter = netmaskFormat.YFilter
    netmaskFormat.EntityData.YangName = "netmask-format"
    netmaskFormat.EntityData.BundleName = "cisco_ios_xr"
    netmaskFormat.EntityData.ParentYangName = "parser"
    netmaskFormat.EntityData.SegmentPath = "netmask-format"
    netmaskFormat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netmaskFormat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netmaskFormat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netmaskFormat.EntityData.Children = make(map[string]types.YChild)
    netmaskFormat.EntityData.Leafs = make(map[string]types.YLeaf)
    netmaskFormat.EntityData.Leafs["bit-count"] = types.YLeaf{"BitCount", netmaskFormat.BitCount}
    return &(netmaskFormat.EntityData)
}

// Parser_Configuration
// cli configuration services
type Parser_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable for read-only access users.
    Disable Parser_Configuration_Disable
}

func (configuration *Parser_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "parser"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Children["disable"] = types.YChild{"Disable", &configuration.Disable}
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configuration.EntityData)
}

// Parser_Configuration_Disable
// disable for read-only access users
type Parser_Configuration_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable config mode for usergroup. The type is string.
    Usergroup interface{}
}

func (disable *Parser_Configuration_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "configuration"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = make(map[string]types.YChild)
    disable.EntityData.Leafs = make(map[string]types.YLeaf)
    disable.EntityData.Leafs["usergroup"] = types.YLeaf{"Usergroup", disable.Usergroup}
    return &(disable.EntityData)
}

// Parser_SubmodeExit
// Exit submode when only '!' seen in interactive
// mode
type Parser_SubmodeExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable the feature. The type is bool.
    Enable interface{}
}

func (submodeExit *Parser_SubmodeExit) GetEntityData() *types.CommonEntityData {
    submodeExit.EntityData.YFilter = submodeExit.YFilter
    submodeExit.EntityData.YangName = "submode-exit"
    submodeExit.EntityData.BundleName = "cisco_ios_xr"
    submodeExit.EntityData.ParentYangName = "parser"
    submodeExit.EntityData.SegmentPath = "submode-exit"
    submodeExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    submodeExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    submodeExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    submodeExit.EntityData.Children = make(map[string]types.YChild)
    submodeExit.EntityData.Leafs = make(map[string]types.YLeaf)
    submodeExit.EntityData.Leafs["enable"] = types.YLeaf{"Enable", submodeExit.Enable}
    return &(submodeExit.EntityData)
}

