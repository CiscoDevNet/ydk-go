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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input SysmgrProcessRestart_Input
}

func (sysmgrProcessRestart *SysmgrProcessRestart) GetEntityData() *types.CommonEntityData {
    sysmgrProcessRestart.EntityData.YFilter = sysmgrProcessRestart.YFilter
    sysmgrProcessRestart.EntityData.YangName = "sysmgr-process-restart"
    sysmgrProcessRestart.EntityData.BundleName = "cisco_ios_xr"
    sysmgrProcessRestart.EntityData.ParentYangName = "Cisco-IOS-XR-sysmgr-act"
    sysmgrProcessRestart.EntityData.SegmentPath = "Cisco-IOS-XR-sysmgr-act:sysmgr-process-restart"
    sysmgrProcessRestart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysmgrProcessRestart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysmgrProcessRestart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysmgrProcessRestart.EntityData.Children = make(map[string]types.YChild)
    sysmgrProcessRestart.EntityData.Children["input"] = types.YChild{"Input", &sysmgrProcessRestart.Input}
    sysmgrProcessRestart.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sysmgrProcessRestart.EntityData)
}

// SysmgrProcessRestart_Input
type SysmgrProcessRestart_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XR process name or Job Id e.g. bgp, ospf. The type is string. This
    // attribute is mandatory.
    ProcessName interface{}

    // XR node identifier e.g. 0/RP0/CPU0, 0/0/CPU0. The type is string.
    Location interface{}
}

func (input *SysmgrProcessRestart_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "sysmgr-process-restart"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", input.ProcessName}
    input.EntityData.Leafs["location"] = types.YLeaf{"Location", input.Location}
    return &(input.EntityData)
}

