// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-nsr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   nsr: NSR global configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_nsr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_nsr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-nsr-cfg nsr}", reflect.TypeOf(Nsr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-nsr-cfg:nsr", reflect.TypeOf(Nsr{}))
}

// Nsr
// NSR global configuration
type Nsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Recovery action for process failures on active RP/DRP.
    ProcessFailure Nsr_ProcessFailure
}

func (nsr *Nsr) GetEntityData() *types.CommonEntityData {
    nsr.EntityData.YFilter = nsr.YFilter
    nsr.EntityData.YangName = "nsr"
    nsr.EntityData.BundleName = "cisco_ios_xr"
    nsr.EntityData.ParentYangName = "Cisco-IOS-XR-infra-nsr-cfg"
    nsr.EntityData.SegmentPath = "Cisco-IOS-XR-infra-nsr-cfg:nsr"
    nsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nsr.EntityData.Children = types.NewOrderedMap()
    nsr.EntityData.Children.Append("process-failure", types.YChild{"ProcessFailure", &nsr.ProcessFailure})
    nsr.EntityData.Leafs = types.NewOrderedMap()

    nsr.EntityData.YListKeys = []string {}

    return &(nsr.EntityData)
}

// Nsr_ProcessFailure
// Recovery action for process failures on active
// RP/DRP
type Nsr_ProcessFailure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable RP/DRP switchover on process failures. The type is interface{}.
    Switchover interface{}
}

func (processFailure *Nsr_ProcessFailure) GetEntityData() *types.CommonEntityData {
    processFailure.EntityData.YFilter = processFailure.YFilter
    processFailure.EntityData.YangName = "process-failure"
    processFailure.EntityData.BundleName = "cisco_ios_xr"
    processFailure.EntityData.ParentYangName = "nsr"
    processFailure.EntityData.SegmentPath = "process-failure"
    processFailure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processFailure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processFailure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processFailure.EntityData.Children = types.NewOrderedMap()
    processFailure.EntityData.Leafs = types.NewOrderedMap()
    processFailure.EntityData.Leafs.Append("switchover", types.YLeaf{"Switchover", processFailure.Switchover})

    processFailure.EntityData.YListKeys = []string {}

    return &(processFailure.EntityData)
}

