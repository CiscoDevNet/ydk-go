// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-sc-envmon package
// admin-plane operational data.
// 
// This module contains definitions
// for the following management objects:
//   environmental-monitoring: Admin Environmental Monitoring
//     Operational data space
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_sc_envmon_admin_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_sc_envmon_admin_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-sc-envmon-admin-oper environmental-monitoring}", reflect.TypeOf(EnvironmentalMonitoring{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-sc-envmon-admin-oper:environmental-monitoring", reflect.TypeOf(EnvironmentalMonitoring{}))
}

// EnvironmentalMonitoring
// Admin Environmental Monitoring Operational data
// space
type EnvironmentalMonitoring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of racks.
    Racks EnvironmentalMonitoring_Racks
}

func (environmentalMonitoring *EnvironmentalMonitoring) GetFilter() yfilter.YFilter { return environmentalMonitoring.YFilter }

func (environmentalMonitoring *EnvironmentalMonitoring) SetFilter(yf yfilter.YFilter) { environmentalMonitoring.YFilter = yf }

func (environmentalMonitoring *EnvironmentalMonitoring) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (environmentalMonitoring *EnvironmentalMonitoring) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-sc-envmon-admin-oper:environmental-monitoring"
}

func (environmentalMonitoring *EnvironmentalMonitoring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &environmentalMonitoring.Racks
    }
    return nil
}

func (environmentalMonitoring *EnvironmentalMonitoring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &environmentalMonitoring.Racks
    return children
}

func (environmentalMonitoring *EnvironmentalMonitoring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (environmentalMonitoring *EnvironmentalMonitoring) GetBundleName() string { return "cisco_ios_xr" }

func (environmentalMonitoring *EnvironmentalMonitoring) GetYangName() string { return "environmental-monitoring" }

func (environmentalMonitoring *EnvironmentalMonitoring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (environmentalMonitoring *EnvironmentalMonitoring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (environmentalMonitoring *EnvironmentalMonitoring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (environmentalMonitoring *EnvironmentalMonitoring) SetParent(parent types.Entity) { environmentalMonitoring.parent = parent }

func (environmentalMonitoring *EnvironmentalMonitoring) GetParent() types.Entity { return environmentalMonitoring.parent }

func (environmentalMonitoring *EnvironmentalMonitoring) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-sc-envmon-admin-oper" }

// EnvironmentalMonitoring_Racks
// Table of racks
type EnvironmentalMonitoring_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of EnvironmentalMonitoring_Racks_Rack.
    Rack []EnvironmentalMonitoring_Racks_Rack
}

func (racks *EnvironmentalMonitoring_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *EnvironmentalMonitoring_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *EnvironmentalMonitoring_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *EnvironmentalMonitoring_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *EnvironmentalMonitoring_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EnvironmentalMonitoring_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *EnvironmentalMonitoring_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *EnvironmentalMonitoring_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *EnvironmentalMonitoring_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *EnvironmentalMonitoring_Racks) GetYangName() string { return "racks" }

func (racks *EnvironmentalMonitoring_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *EnvironmentalMonitoring_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *EnvironmentalMonitoring_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *EnvironmentalMonitoring_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *EnvironmentalMonitoring_Racks) GetParent() types.Entity { return racks.parent }

func (racks *EnvironmentalMonitoring_Racks) GetParentYangName() string { return "environmental-monitoring" }

// EnvironmentalMonitoring_Racks_Rack
// Number
type EnvironmentalMonitoring_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots EnvironmentalMonitoring_Racks_Rack_Slots
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *EnvironmentalMonitoring_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *EnvironmentalMonitoring_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    if yname == "slots" { return "Slots" }
    return ""
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &rack.Slots
    }
    return nil
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &rack.Slots
    return children
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack"] = rack.Rack
    return leafs
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *EnvironmentalMonitoring_Racks_Rack) GetYangName() string { return "rack" }

func (rack *EnvironmentalMonitoring_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *EnvironmentalMonitoring_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *EnvironmentalMonitoring_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *EnvironmentalMonitoring_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *EnvironmentalMonitoring_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *EnvironmentalMonitoring_Racks_Rack) GetParentYangName() string { return "racks" }

// EnvironmentalMonitoring_Racks_Rack_Slots
// Table of slots
type EnvironmentalMonitoring_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is slice of EnvironmentalMonitoring_Racks_Rack_Slots_Slot.
    Slot []EnvironmentalMonitoring_Racks_Rack_Slots_Slot
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EnvironmentalMonitoring_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot
// Name
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Slot interface{}

    // Table of modules.
    Modules EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules
}

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "modules" { return "Modules" }
    return ""
}

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
}

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "modules" {
        return &slot.Modules
    }
    return nil
}

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["modules"] = &slot.Modules
    return children
}

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = slot.Slot
    return leafs
}

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules
// Table of modules
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module.
    Module []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetFilter() yfilter.YFilter { return modules.YFilter }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) SetFilter(yf yfilter.YFilter) { modules.YFilter = yf }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetGoName(yname string) string {
    if yname == "module" { return "Module" }
    return ""
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetSegmentPath() string {
    return "modules"
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "module" {
        for _, c := range modules.Module {
            if modules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module{}
        modules.Module = append(modules.Module, child)
        return &modules.Module[len(modules.Module)-1]
    }
    return nil
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range modules.Module {
        children[modules.Module[i].GetSegmentPath()] = &modules.Module[i]
    }
    return children
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetBundleName() string { return "cisco_ios_xr" }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetYangName() string { return "modules" }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) SetParent(parent types.Entity) { modules.parent = parent }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetParent() types.Entity { return modules.parent }

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetParentYangName() string { return "slot" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module
// Name
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Module name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Module interface{}

    // Table of sensor types.
    SensorTypes EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes

    // Module Power Draw.
    Power EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetFilter() yfilter.YFilter { return module.YFilter }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) SetFilter(yf yfilter.YFilter) { module.YFilter = yf }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetGoName(yname string) string {
    if yname == "module" { return "Module" }
    if yname == "sensor-types" { return "SensorTypes" }
    if yname == "power" { return "Power" }
    return ""
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetSegmentPath() string {
    return "module" + "[module='" + fmt.Sprintf("%v", module.Module) + "']"
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-types" {
        return &module.SensorTypes
    }
    if childYangName == "power" {
        return &module.Power
    }
    return nil
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensor-types"] = &module.SensorTypes
    children["power"] = &module.Power
    return children
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module"] = module.Module
    return leafs
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetBundleName() string { return "cisco_ios_xr" }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetYangName() string { return "module" }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) SetParent(parent types.Entity) { module.parent = parent }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetParent() types.Entity { return module.parent }

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetParentYangName() string { return "modules" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes
// Table of sensor types
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of sensor. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType.
    SensorType []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetFilter() yfilter.YFilter { return sensorTypes.YFilter }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) SetFilter(yf yfilter.YFilter) { sensorTypes.YFilter = yf }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetGoName(yname string) string {
    if yname == "sensor-type" { return "SensorType" }
    return ""
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetSegmentPath() string {
    return "sensor-types"
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-type" {
        for _, c := range sensorTypes.SensorType {
            if sensorTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType{}
        sensorTypes.SensorType = append(sensorTypes.SensorType, child)
        return &sensorTypes.SensorType[len(sensorTypes.SensorType)-1]
    }
    return nil
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorTypes.SensorType {
        children[sensorTypes.SensorType[i].GetSegmentPath()] = &sensorTypes.SensorType[i]
    }
    return children
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetBundleName() string { return "cisco_ios_xr" }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetYangName() string { return "sensor-types" }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) SetParent(parent types.Entity) { sensorTypes.parent = parent }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetParent() types.Entity { return sensorTypes.parent }

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetParentYangName() string { return "module" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType
// Type of sensor
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // Table of sensors.
    SensorNames EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetFilter() yfilter.YFilter { return sensorType.YFilter }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) SetFilter(yf yfilter.YFilter) { sensorType.YFilter = yf }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sensor-names" { return "SensorNames" }
    return ""
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetSegmentPath() string {
    return "sensor-type" + "[type='" + fmt.Sprintf("%v", sensorType.Type) + "']"
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-names" {
        return &sensorType.SensorNames
    }
    return nil
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensor-names"] = &sensorType.SensorNames
    return children
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = sensorType.Type
    return leafs
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetBundleName() string { return "cisco_ios_xr" }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetYangName() string { return "sensor-type" }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) SetParent(parent types.Entity) { sensorType.parent = parent }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetParent() types.Entity { return sensorType.parent }

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetParentYangName() string { return "sensor-types" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames
// Table of sensors
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of sensor. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName.
    SensorName []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetFilter() yfilter.YFilter { return sensorNames.YFilter }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) SetFilter(yf yfilter.YFilter) { sensorNames.YFilter = yf }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetGoName(yname string) string {
    if yname == "sensor-name" { return "SensorName" }
    return ""
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetSegmentPath() string {
    return "sensor-names"
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-name" {
        for _, c := range sensorNames.SensorName {
            if sensorNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName{}
        sensorNames.SensorName = append(sensorNames.SensorName, child)
        return &sensorNames.SensorName[len(sensorNames.SensorName)-1]
    }
    return nil
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorNames.SensorName {
        children[sensorNames.SensorName[i].GetSegmentPath()] = &sensorNames.SensorName[i]
    }
    return children
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetBundleName() string { return "cisco_ios_xr" }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetYangName() string { return "sensor-names" }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) SetParent(parent types.Entity) { sensorNames.parent = parent }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetParent() types.Entity { return sensorNames.parent }

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetParentYangName() string { return "sensor-type" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName
// Name of sensor
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // The sensor value. The type is interface{} with range:
    // -2147483648..2147483647.
    ValueBrief interface{}

    // The threshold information.
    Thresholds EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds

    // Detailed sensor information including the sensor value.
    ValueDetailed EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed
}

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetFilter() yfilter.YFilter { return sensorName.YFilter }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) SetFilter(yf yfilter.YFilter) { sensorName.YFilter = yf }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "value-brief" { return "ValueBrief" }
    if yname == "thresholds" { return "Thresholds" }
    if yname == "value-detailed" { return "ValueDetailed" }
    return ""
}

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetSegmentPath() string {
    return "sensor-name" + "[name='" + fmt.Sprintf("%v", sensorName.Name) + "']"
}

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "thresholds" {
        return &sensorName.Thresholds
    }
    if childYangName == "value-detailed" {
        return &sensorName.ValueDetailed
    }
    return nil
}

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["thresholds"] = &sensorName.Thresholds
    children["value-detailed"] = &sensorName.ValueDetailed
    return children
}

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensorName.Name
    leafs["value-brief"] = sensorName.ValueBrief
    return leafs
}

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetBundleName() string { return "cisco_ios_xr" }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetYangName() string { return "sensor-name" }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) SetParent(parent types.Entity) { sensorName.parent = parent }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetParent() types.Entity { return sensorName.parent }

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetParentYangName() string { return "sensor-names" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds
// The threshold information
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Types of thresholds. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold.
    Threshold []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetFilter() yfilter.YFilter { return thresholds.YFilter }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) SetFilter(yf yfilter.YFilter) { thresholds.YFilter = yf }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetSegmentPath() string {
    return "thresholds"
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        for _, c := range thresholds.Threshold {
            if thresholds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold{}
        thresholds.Threshold = append(thresholds.Threshold, child)
        return &thresholds.Threshold[len(thresholds.Threshold)-1]
    }
    return nil
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholds.Threshold {
        children[thresholds.Threshold[i].GetSegmentPath()] = &thresholds.Threshold[i]
    }
    return children
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetBundleName() string { return "cisco_ios_xr" }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetYangName() string { return "thresholds" }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) SetParent(parent types.Entity) { thresholds.parent = parent }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetParent() types.Entity { return thresholds.parent }

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetParentYangName() string { return "sensor-name" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold
// Types of thresholds
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Threshold type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // Threshold trap enable flag true-ENABLE, false-DISABLE. The type is bool.
    Trap interface{}

    // Threshold value for the sensor. The type is interface{} with range:
    // -2147483648..2147483647.
    ValueBrief interface{}

    // Detailed sensor threshold information.
    ValueDetailed EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed
}

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "trap" { return "Trap" }
    if yname == "value-brief" { return "ValueBrief" }
    if yname == "value-detailed" { return "ValueDetailed" }
    return ""
}

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetSegmentPath() string {
    return "threshold" + "[type='" + fmt.Sprintf("%v", threshold.Type) + "']"
}

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "value-detailed" {
        return &threshold.ValueDetailed
    }
    return nil
}

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["value-detailed"] = &threshold.ValueDetailed
    return children
}

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = threshold.Type
    leafs["trap"] = threshold.Trap
    leafs["value-brief"] = threshold.ValueBrief
    return leafs
}

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetYangName() string { return "threshold" }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetParentYangName() string { return "thresholds" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed
// Detailed sensor threshold
// information
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates minor, major, critical severities. The type is interface{} with
    // range: 0..4294967295.
    ThresholdSeverity interface{}

    // Indicates relation between sensor value and threshold. The type is
    // interface{} with range: 0..4294967295.
    ThresholdRelation interface{}

    // Value of the configured threshold. The type is interface{} with range:
    // 0..4294967295.
    ThresholdValue interface{}

    // Indicates the result of the most recent evaluation of the thresholD. The
    // type is bool.
    ThresholdEvaluation interface{}

    // Indicates whether or not a notification should result, in case of threshold
    // violation. The type is bool.
    ThresholdNotificationEnabled interface{}
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetFilter() yfilter.YFilter { return valueDetailed.YFilter }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) SetFilter(yf yfilter.YFilter) { valueDetailed.YFilter = yf }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetGoName(yname string) string {
    if yname == "threshold-severity" { return "ThresholdSeverity" }
    if yname == "threshold-relation" { return "ThresholdRelation" }
    if yname == "threshold-value" { return "ThresholdValue" }
    if yname == "threshold-evaluation" { return "ThresholdEvaluation" }
    if yname == "threshold-notification-enabled" { return "ThresholdNotificationEnabled" }
    return ""
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetSegmentPath() string {
    return "value-detailed"
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-severity"] = valueDetailed.ThresholdSeverity
    leafs["threshold-relation"] = valueDetailed.ThresholdRelation
    leafs["threshold-value"] = valueDetailed.ThresholdValue
    leafs["threshold-evaluation"] = valueDetailed.ThresholdEvaluation
    leafs["threshold-notification-enabled"] = valueDetailed.ThresholdNotificationEnabled
    return leafs
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetBundleName() string { return "cisco_ios_xr" }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetYangName() string { return "value-detailed" }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) SetParent(parent types.Entity) { valueDetailed.parent = parent }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetParent() types.Entity { return valueDetailed.parent }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetParentYangName() string { return "threshold" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed
// Detailed sensor information including
// the sensor value
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor valid bitmap. The type is interface{} with range: 0..4294967295.
    FieldValidityBitmap interface{}

    // Device Name. The type is string with length: 0..50.
    DeviceDescription interface{}

    // Units of variable being read. The type is string with length: 0..50.
    Units interface{}

    // Identifier for this device. The type is interface{} with range:
    // 0..4294967295.
    DeviceId interface{}

    // Current reading of sensor. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Indicates threshold violation. The type is interface{} with range:
    // 0..4294967295.
    AlarmType interface{}

    // Sensor data type enums. The type is interface{} with range: 0..4294967295.
    DataType interface{}

    // Sensor scale enums. The type is interface{} with range: 0..4294967295.
    Scale interface{}

    // Sensor precision range. The type is interface{} with range: 0..4294967295.
    Precision interface{}

    // Sensor operation state enums. The type is interface{} with range:
    // 0..4294967295.
    Status interface{}

    // Age of the sensor value; set to the current time if directly access the
    // value from sensor. The type is interface{} with range: 0..4294967295.
    AgeTimeStamp interface{}

    // Sensor value update rate;set to 0 if sensor value is updated and evaluated
    // immediately. The type is interface{} with range: 0..4294967295.
    UpdateRate interface{}
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetFilter() yfilter.YFilter { return valueDetailed.YFilter }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) SetFilter(yf yfilter.YFilter) { valueDetailed.YFilter = yf }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetGoName(yname string) string {
    if yname == "field-validity-bitmap" { return "FieldValidityBitmap" }
    if yname == "device-description" { return "DeviceDescription" }
    if yname == "units" { return "Units" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "value" { return "Value" }
    if yname == "alarm-type" { return "AlarmType" }
    if yname == "data-type" { return "DataType" }
    if yname == "scale" { return "Scale" }
    if yname == "precision" { return "Precision" }
    if yname == "status" { return "Status" }
    if yname == "age-time-stamp" { return "AgeTimeStamp" }
    if yname == "update-rate" { return "UpdateRate" }
    return ""
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetSegmentPath() string {
    return "value-detailed"
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["field-validity-bitmap"] = valueDetailed.FieldValidityBitmap
    leafs["device-description"] = valueDetailed.DeviceDescription
    leafs["units"] = valueDetailed.Units
    leafs["device-id"] = valueDetailed.DeviceId
    leafs["value"] = valueDetailed.Value
    leafs["alarm-type"] = valueDetailed.AlarmType
    leafs["data-type"] = valueDetailed.DataType
    leafs["scale"] = valueDetailed.Scale
    leafs["precision"] = valueDetailed.Precision
    leafs["status"] = valueDetailed.Status
    leafs["age-time-stamp"] = valueDetailed.AgeTimeStamp
    leafs["update-rate"] = valueDetailed.UpdateRate
    return leafs
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetBundleName() string { return "cisco_ios_xr" }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetYangName() string { return "value-detailed" }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) SetParent(parent types.Entity) { valueDetailed.parent = parent }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetParent() types.Entity { return valueDetailed.parent }

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetParentYangName() string { return "sensor-name" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power
// Module Power Draw
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed power bag information.
    PowerBag EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag
}

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetFilter() yfilter.YFilter { return power.YFilter }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) SetFilter(yf yfilter.YFilter) { power.YFilter = yf }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetGoName(yname string) string {
    if yname == "power-bag" { return "PowerBag" }
    return ""
}

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetSegmentPath() string {
    return "power"
}

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-bag" {
        return &power.PowerBag
    }
    return nil
}

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["power-bag"] = &power.PowerBag
    return children
}

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetBundleName() string { return "cisco_ios_xr" }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetYangName() string { return "power" }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) SetParent(parent types.Entity) { power.parent = parent }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetParent() types.Entity { return power.parent }

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetParentYangName() string { return "module" }

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag
// Detailed power bag information
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current Power Value of the Unit. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerValue interface{}

    // Max Power Value of the Unit. The type is interface{} with range:
    // -2147483648..2147483647.
    PowerMaxValue interface{}

    // Unit Multiplier of Power. The type is interface{} with range:
    // 0..4294967295.
    PowerUnitMultiplier interface{}

    // Accuracy of the Power Value. The type is interface{} with range:
    // 0..4294967295.
    PowerAccuracy interface{}

    // Measure Caliber. The type is interface{} with range: 0..4294967295.
    PowerMeasureCaliber interface{}

    // Current Type of the Unit. The type is interface{} with range:
    // 0..4294967295.
    PowerCurrentType interface{}

    // The Power Origin of the Unit. The type is interface{} with range:
    // 0..4294967295.
    PowerOrigin interface{}

    // Admin Status of the Unit. The type is interface{} with range:
    // 0..4294967295.
    PowerAdminState interface{}

    // Oper Status of the Unit. The type is interface{} with range: 0..4294967295.
    PowerOperState interface{}

    // Enter Reason for the State. The type is string with length: 0..50.
    PowerStateEnterReason interface{}
}

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetFilter() yfilter.YFilter { return powerBag.YFilter }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) SetFilter(yf yfilter.YFilter) { powerBag.YFilter = yf }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetGoName(yname string) string {
    if yname == "power-value" { return "PowerValue" }
    if yname == "power-max-value" { return "PowerMaxValue" }
    if yname == "power-unit-multiplier" { return "PowerUnitMultiplier" }
    if yname == "power-accuracy" { return "PowerAccuracy" }
    if yname == "power-measure-caliber" { return "PowerMeasureCaliber" }
    if yname == "power-current-type" { return "PowerCurrentType" }
    if yname == "power-origin" { return "PowerOrigin" }
    if yname == "power-admin-state" { return "PowerAdminState" }
    if yname == "power-oper-state" { return "PowerOperState" }
    if yname == "power-state-enter-reason" { return "PowerStateEnterReason" }
    return ""
}

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetSegmentPath() string {
    return "power-bag"
}

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["power-value"] = powerBag.PowerValue
    leafs["power-max-value"] = powerBag.PowerMaxValue
    leafs["power-unit-multiplier"] = powerBag.PowerUnitMultiplier
    leafs["power-accuracy"] = powerBag.PowerAccuracy
    leafs["power-measure-caliber"] = powerBag.PowerMeasureCaliber
    leafs["power-current-type"] = powerBag.PowerCurrentType
    leafs["power-origin"] = powerBag.PowerOrigin
    leafs["power-admin-state"] = powerBag.PowerAdminState
    leafs["power-oper-state"] = powerBag.PowerOperState
    leafs["power-state-enter-reason"] = powerBag.PowerStateEnterReason
    return leafs
}

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetBundleName() string { return "cisco_ios_xr" }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetYangName() string { return "power-bag" }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) SetParent(parent types.Entity) { powerBag.parent = parent }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetParent() types.Entity { return powerBag.parent }

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetParentYangName() string { return "power" }

