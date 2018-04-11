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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of racks.
    Racks EnvironmentalMonitoring_Racks
}

func (environmentalMonitoring *EnvironmentalMonitoring) GetEntityData() *types.CommonEntityData {
    environmentalMonitoring.EntityData.YFilter = environmentalMonitoring.YFilter
    environmentalMonitoring.EntityData.YangName = "environmental-monitoring"
    environmentalMonitoring.EntityData.BundleName = "cisco_ios_xr"
    environmentalMonitoring.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-sc-envmon-admin-oper"
    environmentalMonitoring.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-sc-envmon-admin-oper:environmental-monitoring"
    environmentalMonitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environmentalMonitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environmentalMonitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environmentalMonitoring.EntityData.Children = make(map[string]types.YChild)
    environmentalMonitoring.EntityData.Children["racks"] = types.YChild{"Racks", &environmentalMonitoring.Racks}
    environmentalMonitoring.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(environmentalMonitoring.EntityData)
}

// EnvironmentalMonitoring_Racks
// Table of racks
type EnvironmentalMonitoring_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of EnvironmentalMonitoring_Racks_Rack.
    Rack []EnvironmentalMonitoring_Racks_Rack
}

func (racks *EnvironmentalMonitoring_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "environmental-monitoring"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range racks.Rack {
        racks.EntityData.Children[types.GetSegmentPath(&racks.Rack[i])] = types.YChild{"Rack", &racks.Rack[i]}
    }
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(racks.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack
// Number
type EnvironmentalMonitoring_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots EnvironmentalMonitoring_Racks_Rack_Slots
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["rack"] = types.YLeaf{"Rack", rack.Rack}
    return &(rack.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots
// Table of slots
type EnvironmentalMonitoring_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of EnvironmentalMonitoring_Racks_Rack_Slots_Slot.
    Slot []EnvironmentalMonitoring_Racks_Rack_Slots_Slot
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot
// Name
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Slot interface{}

    // Table of modules.
    Modules EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules
}

func (slot *EnvironmentalMonitoring_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["modules"] = types.YChild{"Modules", &slot.Modules}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot"] = types.YLeaf{"Slot", slot.Slot}
    return &(slot.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules
// Table of modules
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module.
    Module []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetEntityData() *types.CommonEntityData {
    modules.EntityData.YFilter = modules.YFilter
    modules.EntityData.YangName = "modules"
    modules.EntityData.BundleName = "cisco_ios_xr"
    modules.EntityData.ParentYangName = "slot"
    modules.EntityData.SegmentPath = "modules"
    modules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modules.EntityData.Children = make(map[string]types.YChild)
    modules.EntityData.Children["module"] = types.YChild{"Module", nil}
    for i := range modules.Module {
        modules.EntityData.Children[types.GetSegmentPath(&modules.Module[i])] = types.YChild{"Module", &modules.Module[i]}
    }
    modules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(modules.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module
// Name
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Module name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Module interface{}

    // Table of sensor types.
    SensorTypes EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes

    // Module Power Draw.
    Power EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetEntityData() *types.CommonEntityData {
    module.EntityData.YFilter = module.YFilter
    module.EntityData.YangName = "module"
    module.EntityData.BundleName = "cisco_ios_xr"
    module.EntityData.ParentYangName = "modules"
    module.EntityData.SegmentPath = "module" + "[module='" + fmt.Sprintf("%v", module.Module) + "']"
    module.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    module.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    module.EntityData.Children = make(map[string]types.YChild)
    module.EntityData.Children["sensor-types"] = types.YChild{"SensorTypes", &module.SensorTypes}
    module.EntityData.Children["power"] = types.YChild{"Power", &module.Power}
    module.EntityData.Leafs = make(map[string]types.YLeaf)
    module.EntityData.Leafs["module"] = types.YLeaf{"Module", module.Module}
    return &(module.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes
// Table of sensor types
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of sensor. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType.
    SensorType []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetEntityData() *types.CommonEntityData {
    sensorTypes.EntityData.YFilter = sensorTypes.YFilter
    sensorTypes.EntityData.YangName = "sensor-types"
    sensorTypes.EntityData.BundleName = "cisco_ios_xr"
    sensorTypes.EntityData.ParentYangName = "module"
    sensorTypes.EntityData.SegmentPath = "sensor-types"
    sensorTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorTypes.EntityData.Children = make(map[string]types.YChild)
    sensorTypes.EntityData.Children["sensor-type"] = types.YChild{"SensorType", nil}
    for i := range sensorTypes.SensorType {
        sensorTypes.EntityData.Children[types.GetSegmentPath(&sensorTypes.SensorType[i])] = types.YChild{"SensorType", &sensorTypes.SensorType[i]}
    }
    sensorTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sensorTypes.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType
// Type of sensor
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor type. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Type_ interface{}

    // Table of sensors.
    SensorNames EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetEntityData() *types.CommonEntityData {
    sensorType.EntityData.YFilter = sensorType.YFilter
    sensorType.EntityData.YangName = "sensor-type"
    sensorType.EntityData.BundleName = "cisco_ios_xr"
    sensorType.EntityData.ParentYangName = "sensor-types"
    sensorType.EntityData.SegmentPath = "sensor-type" + "[type='" + fmt.Sprintf("%v", sensorType.Type_) + "']"
    sensorType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorType.EntityData.Children = make(map[string]types.YChild)
    sensorType.EntityData.Children["sensor-names"] = types.YChild{"SensorNames", &sensorType.SensorNames}
    sensorType.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorType.EntityData.Leafs["type"] = types.YLeaf{"Type_", sensorType.Type_}
    return &(sensorType.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames
// Table of sensors
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of sensor. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName.
    SensorName []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetEntityData() *types.CommonEntityData {
    sensorNames.EntityData.YFilter = sensorNames.YFilter
    sensorNames.EntityData.YangName = "sensor-names"
    sensorNames.EntityData.BundleName = "cisco_ios_xr"
    sensorNames.EntityData.ParentYangName = "sensor-type"
    sensorNames.EntityData.SegmentPath = "sensor-names"
    sensorNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorNames.EntityData.Children = make(map[string]types.YChild)
    sensorNames.EntityData.Children["sensor-name"] = types.YChild{"SensorName", nil}
    for i := range sensorNames.SensorName {
        sensorNames.EntityData.Children[types.GetSegmentPath(&sensorNames.SensorName[i])] = types.YChild{"SensorName", &sensorNames.SensorName[i]}
    }
    sensorNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sensorNames.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName
// Name of sensor
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // The sensor value. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    ValueBrief interface{}

    // The threshold information.
    Thresholds EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds

    // Detailed sensor information including the sensor value.
    ValueDetailed EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed
}

func (sensorName *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName) GetEntityData() *types.CommonEntityData {
    sensorName.EntityData.YFilter = sensorName.YFilter
    sensorName.EntityData.YangName = "sensor-name"
    sensorName.EntityData.BundleName = "cisco_ios_xr"
    sensorName.EntityData.ParentYangName = "sensor-names"
    sensorName.EntityData.SegmentPath = "sensor-name" + "[name='" + fmt.Sprintf("%v", sensorName.Name) + "']"
    sensorName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorName.EntityData.Children = make(map[string]types.YChild)
    sensorName.EntityData.Children["thresholds"] = types.YChild{"Thresholds", &sensorName.Thresholds}
    sensorName.EntityData.Children["value-detailed"] = types.YChild{"ValueDetailed", &sensorName.ValueDetailed}
    sensorName.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorName.EntityData.Leafs["name"] = types.YLeaf{"Name", sensorName.Name}
    sensorName.EntityData.Leafs["value-brief"] = types.YLeaf{"ValueBrief", sensorName.ValueBrief}
    return &(sensorName.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds
// The threshold information
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Types of thresholds. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold.
    Threshold []EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetEntityData() *types.CommonEntityData {
    thresholds.EntityData.YFilter = thresholds.YFilter
    thresholds.EntityData.YangName = "thresholds"
    thresholds.EntityData.BundleName = "cisco_ios_xr"
    thresholds.EntityData.ParentYangName = "sensor-name"
    thresholds.EntityData.SegmentPath = "thresholds"
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholds.EntityData.Children = make(map[string]types.YChild)
    thresholds.EntityData.Children["threshold"] = types.YChild{"Threshold", nil}
    for i := range thresholds.Threshold {
        thresholds.EntityData.Children[types.GetSegmentPath(&thresholds.Threshold[i])] = types.YChild{"Threshold", &thresholds.Threshold[i]}
    }
    thresholds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholds.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold
// Types of thresholds
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Threshold type. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Type_ interface{}

    // Threshold trap enable flag true-ENABLE, false-DISABLE. The type is bool.
    Trap interface{}

    // Threshold value for the sensor. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    ValueBrief interface{}

    // Detailed sensor threshold information.
    ValueDetailed EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed
}

func (threshold *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "thresholds"
    threshold.EntityData.SegmentPath = "threshold" + "[type='" + fmt.Sprintf("%v", threshold.Type_) + "']"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["value-detailed"] = types.YChild{"ValueDetailed", &threshold.ValueDetailed}
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["type"] = types.YLeaf{"Type_", threshold.Type_}
    threshold.EntityData.Leafs["trap"] = types.YLeaf{"Trap", threshold.Trap}
    threshold.EntityData.Leafs["value-brief"] = types.YLeaf{"ValueBrief", threshold.ValueBrief}
    return &(threshold.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed
// Detailed sensor threshold
// information
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed struct {
    EntityData types.CommonEntityData
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

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold_ValueDetailed) GetEntityData() *types.CommonEntityData {
    valueDetailed.EntityData.YFilter = valueDetailed.YFilter
    valueDetailed.EntityData.YangName = "value-detailed"
    valueDetailed.EntityData.BundleName = "cisco_ios_xr"
    valueDetailed.EntityData.ParentYangName = "threshold"
    valueDetailed.EntityData.SegmentPath = "value-detailed"
    valueDetailed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    valueDetailed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    valueDetailed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    valueDetailed.EntityData.Children = make(map[string]types.YChild)
    valueDetailed.EntityData.Leafs = make(map[string]types.YLeaf)
    valueDetailed.EntityData.Leafs["threshold-severity"] = types.YLeaf{"ThresholdSeverity", valueDetailed.ThresholdSeverity}
    valueDetailed.EntityData.Leafs["threshold-relation"] = types.YLeaf{"ThresholdRelation", valueDetailed.ThresholdRelation}
    valueDetailed.EntityData.Leafs["threshold-value"] = types.YLeaf{"ThresholdValue", valueDetailed.ThresholdValue}
    valueDetailed.EntityData.Leafs["threshold-evaluation"] = types.YLeaf{"ThresholdEvaluation", valueDetailed.ThresholdEvaluation}
    valueDetailed.EntityData.Leafs["threshold-notification-enabled"] = types.YLeaf{"ThresholdNotificationEnabled", valueDetailed.ThresholdNotificationEnabled}
    return &(valueDetailed.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed
// Detailed sensor information including
// the sensor value
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed struct {
    EntityData types.CommonEntityData
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

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetEntityData() *types.CommonEntityData {
    valueDetailed.EntityData.YFilter = valueDetailed.YFilter
    valueDetailed.EntityData.YangName = "value-detailed"
    valueDetailed.EntityData.BundleName = "cisco_ios_xr"
    valueDetailed.EntityData.ParentYangName = "sensor-name"
    valueDetailed.EntityData.SegmentPath = "value-detailed"
    valueDetailed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    valueDetailed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    valueDetailed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    valueDetailed.EntityData.Children = make(map[string]types.YChild)
    valueDetailed.EntityData.Leafs = make(map[string]types.YLeaf)
    valueDetailed.EntityData.Leafs["field-validity-bitmap"] = types.YLeaf{"FieldValidityBitmap", valueDetailed.FieldValidityBitmap}
    valueDetailed.EntityData.Leafs["device-description"] = types.YLeaf{"DeviceDescription", valueDetailed.DeviceDescription}
    valueDetailed.EntityData.Leafs["units"] = types.YLeaf{"Units", valueDetailed.Units}
    valueDetailed.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", valueDetailed.DeviceId}
    valueDetailed.EntityData.Leafs["value"] = types.YLeaf{"Value", valueDetailed.Value}
    valueDetailed.EntityData.Leafs["alarm-type"] = types.YLeaf{"AlarmType", valueDetailed.AlarmType}
    valueDetailed.EntityData.Leafs["data-type"] = types.YLeaf{"DataType", valueDetailed.DataType}
    valueDetailed.EntityData.Leafs["scale"] = types.YLeaf{"Scale", valueDetailed.Scale}
    valueDetailed.EntityData.Leafs["precision"] = types.YLeaf{"Precision", valueDetailed.Precision}
    valueDetailed.EntityData.Leafs["status"] = types.YLeaf{"Status", valueDetailed.Status}
    valueDetailed.EntityData.Leafs["age-time-stamp"] = types.YLeaf{"AgeTimeStamp", valueDetailed.AgeTimeStamp}
    valueDetailed.EntityData.Leafs["update-rate"] = types.YLeaf{"UpdateRate", valueDetailed.UpdateRate}
    return &(valueDetailed.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power
// Module Power Draw
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed power bag information.
    PowerBag EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag
}

func (power *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power) GetEntityData() *types.CommonEntityData {
    power.EntityData.YFilter = power.YFilter
    power.EntityData.YangName = "power"
    power.EntityData.BundleName = "cisco_ios_xr"
    power.EntityData.ParentYangName = "module"
    power.EntityData.SegmentPath = "power"
    power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    power.EntityData.Children = make(map[string]types.YChild)
    power.EntityData.Children["power-bag"] = types.YChild{"PowerBag", &power.PowerBag}
    power.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(power.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag
// Detailed power bag information
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag struct {
    EntityData types.CommonEntityData
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

func (powerBag *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power_PowerBag) GetEntityData() *types.CommonEntityData {
    powerBag.EntityData.YFilter = powerBag.YFilter
    powerBag.EntityData.YangName = "power-bag"
    powerBag.EntityData.BundleName = "cisco_ios_xr"
    powerBag.EntityData.ParentYangName = "power"
    powerBag.EntityData.SegmentPath = "power-bag"
    powerBag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerBag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerBag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerBag.EntityData.Children = make(map[string]types.YChild)
    powerBag.EntityData.Leafs = make(map[string]types.YLeaf)
    powerBag.EntityData.Leafs["power-value"] = types.YLeaf{"PowerValue", powerBag.PowerValue}
    powerBag.EntityData.Leafs["power-max-value"] = types.YLeaf{"PowerMaxValue", powerBag.PowerMaxValue}
    powerBag.EntityData.Leafs["power-unit-multiplier"] = types.YLeaf{"PowerUnitMultiplier", powerBag.PowerUnitMultiplier}
    powerBag.EntityData.Leafs["power-accuracy"] = types.YLeaf{"PowerAccuracy", powerBag.PowerAccuracy}
    powerBag.EntityData.Leafs["power-measure-caliber"] = types.YLeaf{"PowerMeasureCaliber", powerBag.PowerMeasureCaliber}
    powerBag.EntityData.Leafs["power-current-type"] = types.YLeaf{"PowerCurrentType", powerBag.PowerCurrentType}
    powerBag.EntityData.Leafs["power-origin"] = types.YLeaf{"PowerOrigin", powerBag.PowerOrigin}
    powerBag.EntityData.Leafs["power-admin-state"] = types.YLeaf{"PowerAdminState", powerBag.PowerAdminState}
    powerBag.EntityData.Leafs["power-oper-state"] = types.YLeaf{"PowerOperState", powerBag.PowerOperState}
    powerBag.EntityData.Leafs["power-state-enter-reason"] = types.YLeaf{"PowerStateEnterReason", powerBag.PowerStateEnterReason}
    return &(powerBag.EntityData)
}

