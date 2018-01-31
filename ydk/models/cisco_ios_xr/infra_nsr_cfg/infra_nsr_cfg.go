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
    parent types.Entity
    YFilter yfilter.YFilter

    // Recovery action for process failures on active RP/DRP.
    ProcessFailure Nsr_ProcessFailure
}

func (nsr *Nsr) GetFilter() yfilter.YFilter { return nsr.YFilter }

func (nsr *Nsr) SetFilter(yf yfilter.YFilter) { nsr.YFilter = yf }

func (nsr *Nsr) GetGoName(yname string) string {
    if yname == "process-failure" { return "ProcessFailure" }
    return ""
}

func (nsr *Nsr) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-nsr-cfg:nsr"
}

func (nsr *Nsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-failure" {
        return &nsr.ProcessFailure
    }
    return nil
}

func (nsr *Nsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process-failure"] = &nsr.ProcessFailure
    return children
}

func (nsr *Nsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nsr *Nsr) GetBundleName() string { return "cisco_ios_xr" }

func (nsr *Nsr) GetYangName() string { return "nsr" }

func (nsr *Nsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nsr *Nsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nsr *Nsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nsr *Nsr) SetParent(parent types.Entity) { nsr.parent = parent }

func (nsr *Nsr) GetParent() types.Entity { return nsr.parent }

func (nsr *Nsr) GetParentYangName() string { return "Cisco-IOS-XR-infra-nsr-cfg" }

// Nsr_ProcessFailure
// Recovery action for process failures on active
// RP/DRP
type Nsr_ProcessFailure struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable RP/DRP switchover on process failures. The type is interface{}.
    Switchover interface{}
}

func (processFailure *Nsr_ProcessFailure) GetFilter() yfilter.YFilter { return processFailure.YFilter }

func (processFailure *Nsr_ProcessFailure) SetFilter(yf yfilter.YFilter) { processFailure.YFilter = yf }

func (processFailure *Nsr_ProcessFailure) GetGoName(yname string) string {
    if yname == "switchover" { return "Switchover" }
    return ""
}

func (processFailure *Nsr_ProcessFailure) GetSegmentPath() string {
    return "process-failure"
}

func (processFailure *Nsr_ProcessFailure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (processFailure *Nsr_ProcessFailure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (processFailure *Nsr_ProcessFailure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["switchover"] = processFailure.Switchover
    return leafs
}

func (processFailure *Nsr_ProcessFailure) GetBundleName() string { return "cisco_ios_xr" }

func (processFailure *Nsr_ProcessFailure) GetYangName() string { return "process-failure" }

func (processFailure *Nsr_ProcessFailure) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processFailure *Nsr_ProcessFailure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processFailure *Nsr_ProcessFailure) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processFailure *Nsr_ProcessFailure) SetParent(parent types.Entity) { processFailure.parent = parent }

func (processFailure *Nsr_ProcessFailure) GetParent() types.Entity { return processFailure.parent }

func (processFailure *Nsr_ProcessFailure) GetParentYangName() string { return "nsr" }

