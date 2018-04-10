// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1001-ots-act package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-action: NCS1k1 HW module config
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1001_ots_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1001_ots_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1001-ots-act psm-manual-switch-to}", reflect.TypeOf(PsmManualSwitchTo{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1001-ots-act:psm-manual-switch-to", reflect.TypeOf(PsmManualSwitchTo{}))
}

// OtsPsmManualSwitch represents Ots psm manual switch
type OtsPsmManualSwitch string

const (
    // Working port
    OtsPsmManualSwitch_working OtsPsmManualSwitch = "working"

    // Protected port
    OtsPsmManualSwitch_protected OtsPsmManualSwitch = "protected"
)

// PsmManualSwitchTo
// Psm manual switch to port
type PsmManualSwitchTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input PsmManualSwitchTo_Input
}

func (psmManualSwitchTo *PsmManualSwitchTo) GetEntityData() *types.CommonEntityData {
    psmManualSwitchTo.EntityData.YFilter = psmManualSwitchTo.YFilter
    psmManualSwitchTo.EntityData.YangName = "psm-manual-switch-to"
    psmManualSwitchTo.EntityData.BundleName = "cisco_ios_xr"
    psmManualSwitchTo.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1001-ots-act"
    psmManualSwitchTo.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1001-ots-act:psm-manual-switch-to"
    psmManualSwitchTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    psmManualSwitchTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    psmManualSwitchTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    psmManualSwitchTo.EntityData.Children = make(map[string]types.YChild)
    psmManualSwitchTo.EntityData.Children["input"] = types.YChild{"Input", &psmManualSwitchTo.Input}
    psmManualSwitchTo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(psmManualSwitchTo.EntityData)
}

// PsmManualSwitchTo_Input
type PsmManualSwitchTo_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set Slot. The type is interface{} with range: 1..3. This attribute is
    // mandatory.
    SlotId interface{}

    // Switch active path to selected port. The type is OtsPsmManualSwitch. This
    // attribute is mandatory.
    ManualSwitchTo interface{}
}

func (input *PsmManualSwitchTo_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "psm-manual-switch-to"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["slot-id"] = types.YLeaf{"SlotId", input.SlotId}
    input.EntityData.Leafs["manual-switch-to"] = types.YLeaf{"ManualSwitchTo", input.ManualSwitchTo}
    return &(input.EntityData)
}

