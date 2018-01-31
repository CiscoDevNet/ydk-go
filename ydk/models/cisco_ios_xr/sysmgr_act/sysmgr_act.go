// This module contains a collection of YANG definitions
// for Cisco IOS-XR sysmgr action package configuration.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysmgr_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysmgr_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sysmgr-act sysmgr-process-restart}", reflect.TypeOf(SysmgrProcessRestart{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysmgr-act:sysmgr-process-restart", reflect.TypeOf(SysmgrProcessRestart{}))
}

// SysmgrProcessRestart
// Restart an XR process
type SysmgrProcessRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input SysmgrProcessRestart_Input
}

func (sysmgrProcessRestart *SysmgrProcessRestart) GetFilter() yfilter.YFilter { return sysmgrProcessRestart.YFilter }

func (sysmgrProcessRestart *SysmgrProcessRestart) SetFilter(yf yfilter.YFilter) { sysmgrProcessRestart.YFilter = yf }

func (sysmgrProcessRestart *SysmgrProcessRestart) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (sysmgrProcessRestart *SysmgrProcessRestart) GetSegmentPath() string {
    return "Cisco-IOS-XR-sysmgr-act:sysmgr-process-restart"
}

func (sysmgrProcessRestart *SysmgrProcessRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &sysmgrProcessRestart.Input
    }
    return nil
}

func (sysmgrProcessRestart *SysmgrProcessRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &sysmgrProcessRestart.Input
    return children
}

func (sysmgrProcessRestart *SysmgrProcessRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sysmgrProcessRestart *SysmgrProcessRestart) GetBundleName() string { return "cisco_ios_xr" }

func (sysmgrProcessRestart *SysmgrProcessRestart) GetYangName() string { return "sysmgr-process-restart" }

func (sysmgrProcessRestart *SysmgrProcessRestart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sysmgrProcessRestart *SysmgrProcessRestart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sysmgrProcessRestart *SysmgrProcessRestart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sysmgrProcessRestart *SysmgrProcessRestart) SetParent(parent types.Entity) { sysmgrProcessRestart.parent = parent }

func (sysmgrProcessRestart *SysmgrProcessRestart) GetParent() types.Entity { return sysmgrProcessRestart.parent }

func (sysmgrProcessRestart *SysmgrProcessRestart) GetParentYangName() string { return "Cisco-IOS-XR-sysmgr-act" }

// SysmgrProcessRestart_Input
type SysmgrProcessRestart_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // XR process name or Job Id e.g. bgp, ospf. The type is string. This
    // attribute is mandatory.
    ProcessName interface{}

    // XR node identifier e.g. 0/RP0/CPU0, 0/0/CPU0. The type is string.
    Location interface{}
}

func (input *SysmgrProcessRestart_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *SysmgrProcessRestart_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *SysmgrProcessRestart_Input) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    if yname == "location" { return "Location" }
    return ""
}

func (input *SysmgrProcessRestart_Input) GetSegmentPath() string {
    return "input"
}

func (input *SysmgrProcessRestart_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *SysmgrProcessRestart_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *SysmgrProcessRestart_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-name"] = input.ProcessName
    leafs["location"] = input.Location
    return leafs
}

func (input *SysmgrProcessRestart_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *SysmgrProcessRestart_Input) GetYangName() string { return "input" }

func (input *SysmgrProcessRestart_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *SysmgrProcessRestart_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *SysmgrProcessRestart_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *SysmgrProcessRestart_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *SysmgrProcessRestart_Input) GetParent() types.Entity { return input.parent }

func (input *SysmgrProcessRestart_Input) GetParentYangName() string { return "sysmgr-process-restart" }

