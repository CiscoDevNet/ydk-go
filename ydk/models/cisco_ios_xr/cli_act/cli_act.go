// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package cli_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cli_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cli-act cli-command}", reflect.TypeOf(CliCommand{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cli-act:cli-command", reflect.TypeOf(CliCommand{}))
}

// CliCommand
// Actions on certificate revocation lists
type CliCommand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CliCommand_Input

    
    Output CliCommand_Output
}

func (cliCommand *CliCommand) GetEntityData() *types.CommonEntityData {
    cliCommand.EntityData.YFilter = cliCommand.YFilter
    cliCommand.EntityData.YangName = "cli-command"
    cliCommand.EntityData.BundleName = "cisco_ios_xr"
    cliCommand.EntityData.ParentYangName = "Cisco-IOS-XR-cli-act"
    cliCommand.EntityData.SegmentPath = "Cisco-IOS-XR-cli-act:cli-command"
    cliCommand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cliCommand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cliCommand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cliCommand.EntityData.Children = types.NewOrderedMap()
    cliCommand.EntityData.Children.Append("input", types.YChild{"Input", &cliCommand.Input})
    cliCommand.EntityData.Children.Append("output", types.YChild{"Output", &cliCommand.Output})
    cliCommand.EntityData.Leafs = types.NewOrderedMap()

    cliCommand.EntityData.YListKeys = []string {}

    return &(cliCommand.EntityData)
}

// CliCommand_Input
type CliCommand_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CLI command string. The type is string. This attribute is mandatory.
    Command interface{}
}

func (input *CliCommand_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "cli-command"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("command", types.YLeaf{"Command", input.Command})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CliCommand_Output
type CliCommand_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CLI command output string. The type is string.
    Response interface{}
}

func (output *CliCommand_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "cli-command"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("response", types.YLeaf{"Response", output.Response})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

