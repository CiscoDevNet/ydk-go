// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-sc-envmon package operational data.
// 
// This module contains definitions
// for the following management objects:
//   environmental-monitoring-cli: Environmental Monitoring
//     Operational data space
//   environmental-monitoring: environmental monitoring
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_sc_envmon_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_sc_envmon_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-sc-envmon-oper environmental-monitoring-cli}", reflect.TypeOf(EnvironmentalMonitoringCli{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli", reflect.TypeOf(EnvironmentalMonitoringCli{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-sc-envmon-oper environmental-monitoring}", reflect.TypeOf(EnvironmentalMonitoring{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring", reflect.TypeOf(EnvironmentalMonitoring{}))
}

// EnvironmentalMonitoringCli
// Environmental Monitoring Operational data space
type EnvironmentalMonitoringCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of racks.
    RackClis EnvironmentalMonitoringCli_RackClis
}

func (environmentalMonitoringCli *EnvironmentalMonitoringCli) GetEntityData() *types.CommonEntityData {
    environmentalMonitoringCli.EntityData.YFilter = environmentalMonitoringCli.YFilter
    environmentalMonitoringCli.EntityData.YangName = "environmental-monitoring-cli"
    environmentalMonitoringCli.EntityData.BundleName = "cisco_ios_xr"
    environmentalMonitoringCli.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-sc-envmon-oper"
    environmentalMonitoringCli.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli"
    environmentalMonitoringCli.EntityData.AbsolutePath = environmentalMonitoringCli.EntityData.SegmentPath
    environmentalMonitoringCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environmentalMonitoringCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environmentalMonitoringCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environmentalMonitoringCli.EntityData.Children = types.NewOrderedMap()
    environmentalMonitoringCli.EntityData.Children.Append("rack-clis", types.YChild{"RackClis", &environmentalMonitoringCli.RackClis})
    environmentalMonitoringCli.EntityData.Leafs = types.NewOrderedMap()

    environmentalMonitoringCli.EntityData.YListKeys = []string {}

    return &(environmentalMonitoringCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis
// Table of racks
type EnvironmentalMonitoringCli_RackClis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of EnvironmentalMonitoringCli_RackClis_RackCli.
    RackCli []*EnvironmentalMonitoringCli_RackClis_RackCli
}

func (rackClis *EnvironmentalMonitoringCli_RackClis) GetEntityData() *types.CommonEntityData {
    rackClis.EntityData.YFilter = rackClis.YFilter
    rackClis.EntityData.YangName = "rack-clis"
    rackClis.EntityData.BundleName = "cisco_ios_xr"
    rackClis.EntityData.ParentYangName = "environmental-monitoring-cli"
    rackClis.EntityData.SegmentPath = "rack-clis"
    rackClis.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/" + rackClis.EntityData.SegmentPath
    rackClis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackClis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackClis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackClis.EntityData.Children = types.NewOrderedMap()
    rackClis.EntityData.Children.Append("rack-cli", types.YChild{"RackCli", nil})
    for i := range rackClis.RackCli {
        rackClis.EntityData.Children.Append(types.GetSegmentPath(rackClis.RackCli[i]), types.YChild{"RackCli", rackClis.RackCli[i]})
    }
    rackClis.EntityData.Leafs = types.NewOrderedMap()

    rackClis.EntityData.YListKeys = []string {}

    return &(rackClis.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli
// Number
type EnvironmentalMonitoringCli_RackClis_RackCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack number. The type is interface{} with range:
    // 0..4294967295.
    Rack interface{}

    // Table of slots.
    SlotClis EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis
}

func (rackCli *EnvironmentalMonitoringCli_RackClis_RackCli) GetEntityData() *types.CommonEntityData {
    rackCli.EntityData.YFilter = rackCli.YFilter
    rackCli.EntityData.YangName = "rack-cli"
    rackCli.EntityData.BundleName = "cisco_ios_xr"
    rackCli.EntityData.ParentYangName = "rack-clis"
    rackCli.EntityData.SegmentPath = "rack-cli" + types.AddKeyToken(rackCli.Rack, "rack")
    rackCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/" + rackCli.EntityData.SegmentPath
    rackCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackCli.EntityData.Children = types.NewOrderedMap()
    rackCli.EntityData.Children.Append("slot-clis", types.YChild{"SlotClis", &rackCli.SlotClis})
    rackCli.EntityData.Leafs = types.NewOrderedMap()
    rackCli.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", rackCli.Rack})

    rackCli.EntityData.YListKeys = []string {"Rack"}

    return &(rackCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis
// Table of slots
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli.
    SlotCli []*EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli
}

func (slotClis *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis) GetEntityData() *types.CommonEntityData {
    slotClis.EntityData.YFilter = slotClis.YFilter
    slotClis.EntityData.YangName = "slot-clis"
    slotClis.EntityData.BundleName = "cisco_ios_xr"
    slotClis.EntityData.ParentYangName = "rack-cli"
    slotClis.EntityData.SegmentPath = "slot-clis"
    slotClis.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/" + slotClis.EntityData.SegmentPath
    slotClis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slotClis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slotClis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slotClis.EntityData.Children = types.NewOrderedMap()
    slotClis.EntityData.Children.Append("slot-cli", types.YChild{"SlotCli", nil})
    for i := range slotClis.SlotCli {
        slotClis.EntityData.Children.Append(types.GetSegmentPath(slotClis.SlotCli[i]), types.YChild{"SlotCli", slotClis.SlotCli[i]})
    }
    slotClis.EntityData.Leafs = types.NewOrderedMap()

    slotClis.EntityData.YListKeys = []string {}

    return &(slotClis.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli
// Name
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Slot interface{}

    // Table of modules.
    ModuleClis EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis
}

func (slotCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli) GetEntityData() *types.CommonEntityData {
    slotCli.EntityData.YFilter = slotCli.YFilter
    slotCli.EntityData.YangName = "slot-cli"
    slotCli.EntityData.BundleName = "cisco_ios_xr"
    slotCli.EntityData.ParentYangName = "slot-clis"
    slotCli.EntityData.SegmentPath = "slot-cli" + types.AddKeyToken(slotCli.Slot, "slot")
    slotCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/" + slotCli.EntityData.SegmentPath
    slotCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slotCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slotCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slotCli.EntityData.Children = types.NewOrderedMap()
    slotCli.EntityData.Children.Append("module-clis", types.YChild{"ModuleClis", &slotCli.ModuleClis})
    slotCli.EntityData.Leafs = types.NewOrderedMap()
    slotCli.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", slotCli.Slot})

    slotCli.EntityData.YListKeys = []string {"Slot"}

    return &(slotCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis
// Table of modules
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli.
    ModuleCli []*EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli
}

func (moduleClis *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis) GetEntityData() *types.CommonEntityData {
    moduleClis.EntityData.YFilter = moduleClis.YFilter
    moduleClis.EntityData.YangName = "module-clis"
    moduleClis.EntityData.BundleName = "cisco_ios_xr"
    moduleClis.EntityData.ParentYangName = "slot-cli"
    moduleClis.EntityData.SegmentPath = "module-clis"
    moduleClis.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/" + moduleClis.EntityData.SegmentPath
    moduleClis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleClis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleClis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleClis.EntityData.Children = types.NewOrderedMap()
    moduleClis.EntityData.Children.Append("module-cli", types.YChild{"ModuleCli", nil})
    for i := range moduleClis.ModuleCli {
        moduleClis.EntityData.Children.Append(types.GetSegmentPath(moduleClis.ModuleCli[i]), types.YChild{"ModuleCli", moduleClis.ModuleCli[i]})
    }
    moduleClis.EntityData.Leafs = types.NewOrderedMap()

    moduleClis.EntityData.YListKeys = []string {}

    return &(moduleClis.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli
// Name
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Module name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Module interface{}

    // Table of sensor types.
    SensorTypeClis EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis

    // Module Power Draw.
    PowerCli EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli
}

func (moduleCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli) GetEntityData() *types.CommonEntityData {
    moduleCli.EntityData.YFilter = moduleCli.YFilter
    moduleCli.EntityData.YangName = "module-cli"
    moduleCli.EntityData.BundleName = "cisco_ios_xr"
    moduleCli.EntityData.ParentYangName = "module-clis"
    moduleCli.EntityData.SegmentPath = "module-cli" + types.AddKeyToken(moduleCli.Module, "module")
    moduleCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/" + moduleCli.EntityData.SegmentPath
    moduleCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleCli.EntityData.Children = types.NewOrderedMap()
    moduleCli.EntityData.Children.Append("sensor-type-clis", types.YChild{"SensorTypeClis", &moduleCli.SensorTypeClis})
    moduleCli.EntityData.Children.Append("power-cli", types.YChild{"PowerCli", &moduleCli.PowerCli})
    moduleCli.EntityData.Leafs = types.NewOrderedMap()
    moduleCli.EntityData.Leafs.Append("module", types.YLeaf{"Module", moduleCli.Module})

    moduleCli.EntityData.YListKeys = []string {"Module"}

    return &(moduleCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis
// Table of sensor types
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of sensor. The type is slice of
    // EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli.
    SensorTypeCli []*EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli
}

func (sensorTypeClis *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis) GetEntityData() *types.CommonEntityData {
    sensorTypeClis.EntityData.YFilter = sensorTypeClis.YFilter
    sensorTypeClis.EntityData.YangName = "sensor-type-clis"
    sensorTypeClis.EntityData.BundleName = "cisco_ios_xr"
    sensorTypeClis.EntityData.ParentYangName = "module-cli"
    sensorTypeClis.EntityData.SegmentPath = "sensor-type-clis"
    sensorTypeClis.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/" + sensorTypeClis.EntityData.SegmentPath
    sensorTypeClis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorTypeClis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorTypeClis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorTypeClis.EntityData.Children = types.NewOrderedMap()
    sensorTypeClis.EntityData.Children.Append("sensor-type-cli", types.YChild{"SensorTypeCli", nil})
    for i := range sensorTypeClis.SensorTypeCli {
        sensorTypeClis.EntityData.Children.Append(types.GetSegmentPath(sensorTypeClis.SensorTypeCli[i]), types.YChild{"SensorTypeCli", sensorTypeClis.SensorTypeCli[i]})
    }
    sensorTypeClis.EntityData.Leafs = types.NewOrderedMap()

    sensorTypeClis.EntityData.YListKeys = []string {}

    return &(sensorTypeClis.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli
// Type of sensor
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor type. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Type interface{}

    // Table of sensors.
    SensorNameClis EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis
}

func (sensorTypeCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli) GetEntityData() *types.CommonEntityData {
    sensorTypeCli.EntityData.YFilter = sensorTypeCli.YFilter
    sensorTypeCli.EntityData.YangName = "sensor-type-cli"
    sensorTypeCli.EntityData.BundleName = "cisco_ios_xr"
    sensorTypeCli.EntityData.ParentYangName = "sensor-type-clis"
    sensorTypeCli.EntityData.SegmentPath = "sensor-type-cli" + types.AddKeyToken(sensorTypeCli.Type, "type")
    sensorTypeCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/sensor-type-clis/" + sensorTypeCli.EntityData.SegmentPath
    sensorTypeCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorTypeCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorTypeCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorTypeCli.EntityData.Children = types.NewOrderedMap()
    sensorTypeCli.EntityData.Children.Append("sensor-name-clis", types.YChild{"SensorNameClis", &sensorTypeCli.SensorNameClis})
    sensorTypeCli.EntityData.Leafs = types.NewOrderedMap()
    sensorTypeCli.EntityData.Leafs.Append("type", types.YLeaf{"Type", sensorTypeCli.Type})

    sensorTypeCli.EntityData.YListKeys = []string {"Type"}

    return &(sensorTypeCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis
// Table of sensors
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of sensor. The type is slice of
    // EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli.
    SensorNameCli []*EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli
}

func (sensorNameClis *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis) GetEntityData() *types.CommonEntityData {
    sensorNameClis.EntityData.YFilter = sensorNameClis.YFilter
    sensorNameClis.EntityData.YangName = "sensor-name-clis"
    sensorNameClis.EntityData.BundleName = "cisco_ios_xr"
    sensorNameClis.EntityData.ParentYangName = "sensor-type-cli"
    sensorNameClis.EntityData.SegmentPath = "sensor-name-clis"
    sensorNameClis.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/sensor-type-clis/sensor-type-cli/" + sensorNameClis.EntityData.SegmentPath
    sensorNameClis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorNameClis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorNameClis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorNameClis.EntityData.Children = types.NewOrderedMap()
    sensorNameClis.EntityData.Children.Append("sensor-name-cli", types.YChild{"SensorNameCli", nil})
    for i := range sensorNameClis.SensorNameCli {
        sensorNameClis.EntityData.Children.Append(types.GetSegmentPath(sensorNameClis.SensorNameCli[i]), types.YChild{"SensorNameCli", sensorNameClis.SensorNameCli[i]})
    }
    sensorNameClis.EntityData.Leafs = types.NewOrderedMap()

    sensorNameClis.EntityData.YListKeys = []string {}

    return &(sensorNameClis.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli
// Name of sensor
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // The sensor value. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    ValueBriefCli interface{}

    // Detailed sensor information including the sensor value.
    ValueDetailedCli EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ValueDetailedCli

    // The threshold information.
    ThresholdClis EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis
}

func (sensorNameCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli) GetEntityData() *types.CommonEntityData {
    sensorNameCli.EntityData.YFilter = sensorNameCli.YFilter
    sensorNameCli.EntityData.YangName = "sensor-name-cli"
    sensorNameCli.EntityData.BundleName = "cisco_ios_xr"
    sensorNameCli.EntityData.ParentYangName = "sensor-name-clis"
    sensorNameCli.EntityData.SegmentPath = "sensor-name-cli" + types.AddKeyToken(sensorNameCli.Name, "name")
    sensorNameCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/sensor-type-clis/sensor-type-cli/sensor-name-clis/" + sensorNameCli.EntityData.SegmentPath
    sensorNameCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorNameCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorNameCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorNameCli.EntityData.Children = types.NewOrderedMap()
    sensorNameCli.EntityData.Children.Append("value-detailed-cli", types.YChild{"ValueDetailedCli", &sensorNameCli.ValueDetailedCli})
    sensorNameCli.EntityData.Children.Append("threshold-clis", types.YChild{"ThresholdClis", &sensorNameCli.ThresholdClis})
    sensorNameCli.EntityData.Leafs = types.NewOrderedMap()
    sensorNameCli.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensorNameCli.Name})
    sensorNameCli.EntityData.Leafs.Append("value-brief-cli", types.YLeaf{"ValueBriefCli", sensorNameCli.ValueBriefCli})

    sensorNameCli.EntityData.YListKeys = []string {"Name"}

    return &(sensorNameCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ValueDetailedCli
// Detailed sensor information including
// the sensor value
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ValueDetailedCli struct {
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

    // Average sensor value over time interval. The type is interface{} with
    // range: -2147483648..2147483647.
    Average interface{}

    // Minimum Sensor value over time interval. The type is interface{} with
    // range: -2147483648..2147483647.
    Minimum interface{}

    // Maximum Sensor value over time interval. The type is interface{} with
    // range: -2147483648..2147483647.
    Maximum interface{}

    // Time Interval over which sensor value is monitored. The type is interface{}
    // with range: -2147483648..2147483647.
    Interval interface{}
}

func (valueDetailedCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ValueDetailedCli) GetEntityData() *types.CommonEntityData {
    valueDetailedCli.EntityData.YFilter = valueDetailedCli.YFilter
    valueDetailedCli.EntityData.YangName = "value-detailed-cli"
    valueDetailedCli.EntityData.BundleName = "cisco_ios_xr"
    valueDetailedCli.EntityData.ParentYangName = "sensor-name-cli"
    valueDetailedCli.EntityData.SegmentPath = "value-detailed-cli"
    valueDetailedCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/sensor-type-clis/sensor-type-cli/sensor-name-clis/sensor-name-cli/" + valueDetailedCli.EntityData.SegmentPath
    valueDetailedCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    valueDetailedCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    valueDetailedCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    valueDetailedCli.EntityData.Children = types.NewOrderedMap()
    valueDetailedCli.EntityData.Leafs = types.NewOrderedMap()
    valueDetailedCli.EntityData.Leafs.Append("field-validity-bitmap", types.YLeaf{"FieldValidityBitmap", valueDetailedCli.FieldValidityBitmap})
    valueDetailedCli.EntityData.Leafs.Append("device-description", types.YLeaf{"DeviceDescription", valueDetailedCli.DeviceDescription})
    valueDetailedCli.EntityData.Leafs.Append("units", types.YLeaf{"Units", valueDetailedCli.Units})
    valueDetailedCli.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", valueDetailedCli.DeviceId})
    valueDetailedCli.EntityData.Leafs.Append("value", types.YLeaf{"Value", valueDetailedCli.Value})
    valueDetailedCli.EntityData.Leafs.Append("alarm-type", types.YLeaf{"AlarmType", valueDetailedCli.AlarmType})
    valueDetailedCli.EntityData.Leafs.Append("data-type", types.YLeaf{"DataType", valueDetailedCli.DataType})
    valueDetailedCli.EntityData.Leafs.Append("scale", types.YLeaf{"Scale", valueDetailedCli.Scale})
    valueDetailedCli.EntityData.Leafs.Append("precision", types.YLeaf{"Precision", valueDetailedCli.Precision})
    valueDetailedCli.EntityData.Leafs.Append("status", types.YLeaf{"Status", valueDetailedCli.Status})
    valueDetailedCli.EntityData.Leafs.Append("age-time-stamp", types.YLeaf{"AgeTimeStamp", valueDetailedCli.AgeTimeStamp})
    valueDetailedCli.EntityData.Leafs.Append("update-rate", types.YLeaf{"UpdateRate", valueDetailedCli.UpdateRate})
    valueDetailedCli.EntityData.Leafs.Append("average", types.YLeaf{"Average", valueDetailedCli.Average})
    valueDetailedCli.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", valueDetailedCli.Minimum})
    valueDetailedCli.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", valueDetailedCli.Maximum})
    valueDetailedCli.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", valueDetailedCli.Interval})

    valueDetailedCli.EntityData.YListKeys = []string {}

    return &(valueDetailedCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis
// The threshold information
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Types of thresholds. The type is slice of
    // EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli.
    ThresholdCli []*EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli
}

func (thresholdClis *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis) GetEntityData() *types.CommonEntityData {
    thresholdClis.EntityData.YFilter = thresholdClis.YFilter
    thresholdClis.EntityData.YangName = "threshold-clis"
    thresholdClis.EntityData.BundleName = "cisco_ios_xr"
    thresholdClis.EntityData.ParentYangName = "sensor-name-cli"
    thresholdClis.EntityData.SegmentPath = "threshold-clis"
    thresholdClis.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/sensor-type-clis/sensor-type-cli/sensor-name-clis/sensor-name-cli/" + thresholdClis.EntityData.SegmentPath
    thresholdClis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdClis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdClis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdClis.EntityData.Children = types.NewOrderedMap()
    thresholdClis.EntityData.Children.Append("threshold-cli", types.YChild{"ThresholdCli", nil})
    for i := range thresholdClis.ThresholdCli {
        thresholdClis.EntityData.Children.Append(types.GetSegmentPath(thresholdClis.ThresholdCli[i]), types.YChild{"ThresholdCli", thresholdClis.ThresholdCli[i]})
    }
    thresholdClis.EntityData.Leafs = types.NewOrderedMap()

    thresholdClis.EntityData.YListKeys = []string {}

    return &(thresholdClis.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli
// Types of thresholds
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Threshold type. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Type interface{}

    // Threshold trap enable flag true-ENABLE, false-DISABLE. The type is bool.
    TrapCli interface{}

    // Threshold value for the sensor. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    ValueBriefCli interface{}

    // Detailed sensor threshold information.
    ValueDetailedCli EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli_ValueDetailedCli
}

func (thresholdCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli) GetEntityData() *types.CommonEntityData {
    thresholdCli.EntityData.YFilter = thresholdCli.YFilter
    thresholdCli.EntityData.YangName = "threshold-cli"
    thresholdCli.EntityData.BundleName = "cisco_ios_xr"
    thresholdCli.EntityData.ParentYangName = "threshold-clis"
    thresholdCli.EntityData.SegmentPath = "threshold-cli" + types.AddKeyToken(thresholdCli.Type, "type")
    thresholdCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/sensor-type-clis/sensor-type-cli/sensor-name-clis/sensor-name-cli/threshold-clis/" + thresholdCli.EntityData.SegmentPath
    thresholdCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdCli.EntityData.Children = types.NewOrderedMap()
    thresholdCli.EntityData.Children.Append("value-detailed-cli", types.YChild{"ValueDetailedCli", &thresholdCli.ValueDetailedCli})
    thresholdCli.EntityData.Leafs = types.NewOrderedMap()
    thresholdCli.EntityData.Leafs.Append("type", types.YLeaf{"Type", thresholdCli.Type})
    thresholdCli.EntityData.Leafs.Append("trap-cli", types.YLeaf{"TrapCli", thresholdCli.TrapCli})
    thresholdCli.EntityData.Leafs.Append("value-brief-cli", types.YLeaf{"ValueBriefCli", thresholdCli.ValueBriefCli})

    thresholdCli.EntityData.YListKeys = []string {"Type"}

    return &(thresholdCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli_ValueDetailedCli
// Detailed sensor threshold
// information
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli_ValueDetailedCli struct {
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

func (valueDetailedCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_SensorTypeClis_SensorTypeCli_SensorNameClis_SensorNameCli_ThresholdClis_ThresholdCli_ValueDetailedCli) GetEntityData() *types.CommonEntityData {
    valueDetailedCli.EntityData.YFilter = valueDetailedCli.YFilter
    valueDetailedCli.EntityData.YangName = "value-detailed-cli"
    valueDetailedCli.EntityData.BundleName = "cisco_ios_xr"
    valueDetailedCli.EntityData.ParentYangName = "threshold-cli"
    valueDetailedCli.EntityData.SegmentPath = "value-detailed-cli"
    valueDetailedCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/sensor-type-clis/sensor-type-cli/sensor-name-clis/sensor-name-cli/threshold-clis/threshold-cli/" + valueDetailedCli.EntityData.SegmentPath
    valueDetailedCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    valueDetailedCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    valueDetailedCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    valueDetailedCli.EntityData.Children = types.NewOrderedMap()
    valueDetailedCli.EntityData.Leafs = types.NewOrderedMap()
    valueDetailedCli.EntityData.Leafs.Append("threshold-severity", types.YLeaf{"ThresholdSeverity", valueDetailedCli.ThresholdSeverity})
    valueDetailedCli.EntityData.Leafs.Append("threshold-relation", types.YLeaf{"ThresholdRelation", valueDetailedCli.ThresholdRelation})
    valueDetailedCli.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", valueDetailedCli.ThresholdValue})
    valueDetailedCli.EntityData.Leafs.Append("threshold-evaluation", types.YLeaf{"ThresholdEvaluation", valueDetailedCli.ThresholdEvaluation})
    valueDetailedCli.EntityData.Leafs.Append("threshold-notification-enabled", types.YLeaf{"ThresholdNotificationEnabled", valueDetailedCli.ThresholdNotificationEnabled})

    valueDetailedCli.EntityData.YListKeys = []string {}

    return &(valueDetailedCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli
// Module Power Draw
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed power bag information.
    PowerBagCli EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli_PowerBagCli
}

func (powerCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli) GetEntityData() *types.CommonEntityData {
    powerCli.EntityData.YFilter = powerCli.YFilter
    powerCli.EntityData.YangName = "power-cli"
    powerCli.EntityData.BundleName = "cisco_ios_xr"
    powerCli.EntityData.ParentYangName = "module-cli"
    powerCli.EntityData.SegmentPath = "power-cli"
    powerCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/" + powerCli.EntityData.SegmentPath
    powerCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerCli.EntityData.Children = types.NewOrderedMap()
    powerCli.EntityData.Children.Append("power-bag-cli", types.YChild{"PowerBagCli", &powerCli.PowerBagCli})
    powerCli.EntityData.Leafs = types.NewOrderedMap()

    powerCli.EntityData.YListKeys = []string {}

    return &(powerCli.EntityData)
}

// EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli_PowerBagCli
// Detailed power bag information
type EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli_PowerBagCli struct {
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

func (powerBagCli *EnvironmentalMonitoringCli_RackClis_RackCli_SlotClis_SlotCli_ModuleClis_ModuleCli_PowerCli_PowerBagCli) GetEntityData() *types.CommonEntityData {
    powerBagCli.EntityData.YFilter = powerBagCli.YFilter
    powerBagCli.EntityData.YangName = "power-bag-cli"
    powerBagCli.EntityData.BundleName = "cisco_ios_xr"
    powerBagCli.EntityData.ParentYangName = "power-cli"
    powerBagCli.EntityData.SegmentPath = "power-bag-cli"
    powerBagCli.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring-cli/rack-clis/rack-cli/slot-clis/slot-cli/module-clis/module-cli/power-cli/" + powerBagCli.EntityData.SegmentPath
    powerBagCli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerBagCli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerBagCli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerBagCli.EntityData.Children = types.NewOrderedMap()
    powerBagCli.EntityData.Leafs = types.NewOrderedMap()
    powerBagCli.EntityData.Leafs.Append("power-value", types.YLeaf{"PowerValue", powerBagCli.PowerValue})
    powerBagCli.EntityData.Leafs.Append("power-max-value", types.YLeaf{"PowerMaxValue", powerBagCli.PowerMaxValue})
    powerBagCli.EntityData.Leafs.Append("power-unit-multiplier", types.YLeaf{"PowerUnitMultiplier", powerBagCli.PowerUnitMultiplier})
    powerBagCli.EntityData.Leafs.Append("power-accuracy", types.YLeaf{"PowerAccuracy", powerBagCli.PowerAccuracy})
    powerBagCli.EntityData.Leafs.Append("power-measure-caliber", types.YLeaf{"PowerMeasureCaliber", powerBagCli.PowerMeasureCaliber})
    powerBagCli.EntityData.Leafs.Append("power-current-type", types.YLeaf{"PowerCurrentType", powerBagCli.PowerCurrentType})
    powerBagCli.EntityData.Leafs.Append("power-origin", types.YLeaf{"PowerOrigin", powerBagCli.PowerOrigin})
    powerBagCli.EntityData.Leafs.Append("power-admin-state", types.YLeaf{"PowerAdminState", powerBagCli.PowerAdminState})
    powerBagCli.EntityData.Leafs.Append("power-oper-state", types.YLeaf{"PowerOperState", powerBagCli.PowerOperState})
    powerBagCli.EntityData.Leafs.Append("power-state-enter-reason", types.YLeaf{"PowerStateEnterReason", powerBagCli.PowerStateEnterReason})

    powerBagCli.EntityData.YListKeys = []string {}

    return &(powerBagCli.EntityData)
}

// EnvironmentalMonitoring
// environmental monitoring
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
    environmentalMonitoring.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-sc-envmon-oper"
    environmentalMonitoring.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring"
    environmentalMonitoring.EntityData.AbsolutePath = environmentalMonitoring.EntityData.SegmentPath
    environmentalMonitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environmentalMonitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environmentalMonitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environmentalMonitoring.EntityData.Children = types.NewOrderedMap()
    environmentalMonitoring.EntityData.Children.Append("racks", types.YChild{"Racks", &environmentalMonitoring.Racks})
    environmentalMonitoring.EntityData.Leafs = types.NewOrderedMap()

    environmentalMonitoring.EntityData.YListKeys = []string {}

    return &(environmentalMonitoring.EntityData)
}

// EnvironmentalMonitoring_Racks
// Table of racks
type EnvironmentalMonitoring_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of EnvironmentalMonitoring_Racks_Rack.
    Rack []*EnvironmentalMonitoring_Racks_Rack
}

func (racks *EnvironmentalMonitoring_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "environmental-monitoring"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/" + racks.EntityData.SegmentPath
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range racks.Rack {
        racks.EntityData.Children.Append(types.GetSegmentPath(racks.Rack[i]), types.YChild{"Rack", racks.Rack[i]})
    }
    racks.EntityData.Leafs = types.NewOrderedMap()

    racks.EntityData.YListKeys = []string {}

    return &(racks.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack
// Number
type EnvironmentalMonitoring_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack number. The type is interface{} with range:
    // 0..4294967295.
    Rack interface{}

    // Table of slots.
    Slots EnvironmentalMonitoring_Racks_Rack_Slots
}

func (rack *EnvironmentalMonitoring_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Rack, "rack")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("slots", types.YChild{"Slots", &rack.Slots})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", rack.Rack})

    rack.EntityData.YListKeys = []string {"Rack"}

    return &(rack.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots
// Table of slots
type EnvironmentalMonitoring_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of EnvironmentalMonitoring_Racks_Rack_Slots_Slot.
    Slot []*EnvironmentalMonitoring_Racks_Rack_Slots_Slot
}

func (slots *EnvironmentalMonitoring_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/" + slots.EntityData.SegmentPath
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = types.NewOrderedMap()
    slots.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range slots.Slot {
        slots.EntityData.Children.Append(types.GetSegmentPath(slots.Slot[i]), types.YChild{"Slot", slots.Slot[i]})
    }
    slots.EntityData.Leafs = types.NewOrderedMap()

    slots.EntityData.YListKeys = []string {}

    return &(slots.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot
// Name
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.Slot, "slot")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("modules", types.YChild{"Modules", &slot.Modules})
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", slot.Slot})

    slot.EntityData.YListKeys = []string {"Slot"}

    return &(slot.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules
// Table of modules
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module.
    Module []*EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module
}

func (modules *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules) GetEntityData() *types.CommonEntityData {
    modules.EntityData.YFilter = modules.YFilter
    modules.EntityData.YangName = "modules"
    modules.EntityData.BundleName = "cisco_ios_xr"
    modules.EntityData.ParentYangName = "slot"
    modules.EntityData.SegmentPath = "modules"
    modules.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/" + modules.EntityData.SegmentPath
    modules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modules.EntityData.Children = types.NewOrderedMap()
    modules.EntityData.Children.Append("module", types.YChild{"Module", nil})
    for i := range modules.Module {
        modules.EntityData.Children.Append(types.GetSegmentPath(modules.Module[i]), types.YChild{"Module", modules.Module[i]})
    }
    modules.EntityData.Leafs = types.NewOrderedMap()

    modules.EntityData.YListKeys = []string {}

    return &(modules.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module
// Name
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Module name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Module interface{}

    // Module Power Draw.
    Power EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_Power

    // Table of sensor types.
    SensorTypes EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes
}

func (module *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module) GetEntityData() *types.CommonEntityData {
    module.EntityData.YFilter = module.YFilter
    module.EntityData.YangName = "module"
    module.EntityData.BundleName = "cisco_ios_xr"
    module.EntityData.ParentYangName = "modules"
    module.EntityData.SegmentPath = "module" + types.AddKeyToken(module.Module, "module")
    module.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/" + module.EntityData.SegmentPath
    module.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    module.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    module.EntityData.Children = types.NewOrderedMap()
    module.EntityData.Children.Append("power", types.YChild{"Power", &module.Power})
    module.EntityData.Children.Append("sensor-types", types.YChild{"SensorTypes", &module.SensorTypes})
    module.EntityData.Leafs = types.NewOrderedMap()
    module.EntityData.Leafs.Append("module", types.YLeaf{"Module", module.Module})

    module.EntityData.YListKeys = []string {"Module"}

    return &(module.EntityData)
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
    power.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/" + power.EntityData.SegmentPath
    power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    power.EntityData.Children = types.NewOrderedMap()
    power.EntityData.Children.Append("power-bag", types.YChild{"PowerBag", &power.PowerBag})
    power.EntityData.Leafs = types.NewOrderedMap()

    power.EntityData.YListKeys = []string {}

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
    powerBag.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/power/" + powerBag.EntityData.SegmentPath
    powerBag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerBag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerBag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerBag.EntityData.Children = types.NewOrderedMap()
    powerBag.EntityData.Leafs = types.NewOrderedMap()
    powerBag.EntityData.Leafs.Append("power-value", types.YLeaf{"PowerValue", powerBag.PowerValue})
    powerBag.EntityData.Leafs.Append("power-max-value", types.YLeaf{"PowerMaxValue", powerBag.PowerMaxValue})
    powerBag.EntityData.Leafs.Append("power-unit-multiplier", types.YLeaf{"PowerUnitMultiplier", powerBag.PowerUnitMultiplier})
    powerBag.EntityData.Leafs.Append("power-accuracy", types.YLeaf{"PowerAccuracy", powerBag.PowerAccuracy})
    powerBag.EntityData.Leafs.Append("power-measure-caliber", types.YLeaf{"PowerMeasureCaliber", powerBag.PowerMeasureCaliber})
    powerBag.EntityData.Leafs.Append("power-current-type", types.YLeaf{"PowerCurrentType", powerBag.PowerCurrentType})
    powerBag.EntityData.Leafs.Append("power-origin", types.YLeaf{"PowerOrigin", powerBag.PowerOrigin})
    powerBag.EntityData.Leafs.Append("power-admin-state", types.YLeaf{"PowerAdminState", powerBag.PowerAdminState})
    powerBag.EntityData.Leafs.Append("power-oper-state", types.YLeaf{"PowerOperState", powerBag.PowerOperState})
    powerBag.EntityData.Leafs.Append("power-state-enter-reason", types.YLeaf{"PowerStateEnterReason", powerBag.PowerStateEnterReason})

    powerBag.EntityData.YListKeys = []string {}

    return &(powerBag.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes
// Table of sensor types
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of sensor. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType.
    SensorType []*EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType
}

func (sensorTypes *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes) GetEntityData() *types.CommonEntityData {
    sensorTypes.EntityData.YFilter = sensorTypes.YFilter
    sensorTypes.EntityData.YangName = "sensor-types"
    sensorTypes.EntityData.BundleName = "cisco_ios_xr"
    sensorTypes.EntityData.ParentYangName = "module"
    sensorTypes.EntityData.SegmentPath = "sensor-types"
    sensorTypes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/" + sensorTypes.EntityData.SegmentPath
    sensorTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorTypes.EntityData.Children = types.NewOrderedMap()
    sensorTypes.EntityData.Children.Append("sensor-type", types.YChild{"SensorType", nil})
    for i := range sensorTypes.SensorType {
        sensorTypes.EntityData.Children.Append(types.GetSegmentPath(sensorTypes.SensorType[i]), types.YChild{"SensorType", sensorTypes.SensorType[i]})
    }
    sensorTypes.EntityData.Leafs = types.NewOrderedMap()

    sensorTypes.EntityData.YListKeys = []string {}

    return &(sensorTypes.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType
// Type of sensor
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor type. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Type interface{}

    // Table of sensors.
    SensorNames EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames
}

func (sensorType *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType) GetEntityData() *types.CommonEntityData {
    sensorType.EntityData.YFilter = sensorType.YFilter
    sensorType.EntityData.YangName = "sensor-type"
    sensorType.EntityData.BundleName = "cisco_ios_xr"
    sensorType.EntityData.ParentYangName = "sensor-types"
    sensorType.EntityData.SegmentPath = "sensor-type" + types.AddKeyToken(sensorType.Type, "type")
    sensorType.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/sensor-types/" + sensorType.EntityData.SegmentPath
    sensorType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorType.EntityData.Children = types.NewOrderedMap()
    sensorType.EntityData.Children.Append("sensor-names", types.YChild{"SensorNames", &sensorType.SensorNames})
    sensorType.EntityData.Leafs = types.NewOrderedMap()
    sensorType.EntityData.Leafs.Append("type", types.YLeaf{"Type", sensorType.Type})

    sensorType.EntityData.YListKeys = []string {"Type"}

    return &(sensorType.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames
// Table of sensors
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of sensor. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName.
    SensorName []*EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName
}

func (sensorNames *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames) GetEntityData() *types.CommonEntityData {
    sensorNames.EntityData.YFilter = sensorNames.YFilter
    sensorNames.EntityData.YangName = "sensor-names"
    sensorNames.EntityData.BundleName = "cisco_ios_xr"
    sensorNames.EntityData.ParentYangName = "sensor-type"
    sensorNames.EntityData.SegmentPath = "sensor-names"
    sensorNames.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/sensor-types/sensor-type/" + sensorNames.EntityData.SegmentPath
    sensorNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorNames.EntityData.Children = types.NewOrderedMap()
    sensorNames.EntityData.Children.Append("sensor-name", types.YChild{"SensorName", nil})
    for i := range sensorNames.SensorName {
        sensorNames.EntityData.Children.Append(types.GetSegmentPath(sensorNames.SensorName[i]), types.YChild{"SensorName", sensorNames.SensorName[i]})
    }
    sensorNames.EntityData.Leafs = types.NewOrderedMap()

    sensorNames.EntityData.YListKeys = []string {}

    return &(sensorNames.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName
// Name of sensor
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    sensorName.EntityData.SegmentPath = "sensor-name" + types.AddKeyToken(sensorName.Name, "name")
    sensorName.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/sensor-types/sensor-type/sensor-names/" + sensorName.EntityData.SegmentPath
    sensorName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorName.EntityData.Children = types.NewOrderedMap()
    sensorName.EntityData.Children.Append("thresholds", types.YChild{"Thresholds", &sensorName.Thresholds})
    sensorName.EntityData.Children.Append("value-detailed", types.YChild{"ValueDetailed", &sensorName.ValueDetailed})
    sensorName.EntityData.Leafs = types.NewOrderedMap()
    sensorName.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensorName.Name})
    sensorName.EntityData.Leafs.Append("value-brief", types.YLeaf{"ValueBrief", sensorName.ValueBrief})

    sensorName.EntityData.YListKeys = []string {"Name"}

    return &(sensorName.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds
// The threshold information
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Types of thresholds. The type is slice of
    // EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold.
    Threshold []*EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold
}

func (thresholds *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds) GetEntityData() *types.CommonEntityData {
    thresholds.EntityData.YFilter = thresholds.YFilter
    thresholds.EntityData.YangName = "thresholds"
    thresholds.EntityData.BundleName = "cisco_ios_xr"
    thresholds.EntityData.ParentYangName = "sensor-name"
    thresholds.EntityData.SegmentPath = "thresholds"
    thresholds.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/sensor-types/sensor-type/sensor-names/sensor-name/" + thresholds.EntityData.SegmentPath
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholds.EntityData.Children = types.NewOrderedMap()
    thresholds.EntityData.Children.Append("threshold", types.YChild{"Threshold", nil})
    for i := range thresholds.Threshold {
        thresholds.EntityData.Children.Append(types.GetSegmentPath(thresholds.Threshold[i]), types.YChild{"Threshold", thresholds.Threshold[i]})
    }
    thresholds.EntityData.Leafs = types.NewOrderedMap()

    thresholds.EntityData.YListKeys = []string {}

    return &(thresholds.EntityData)
}

// EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold
// Types of thresholds
type EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_Thresholds_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Threshold type. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Type interface{}

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
    threshold.EntityData.SegmentPath = "threshold" + types.AddKeyToken(threshold.Type, "type")
    threshold.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/sensor-types/sensor-type/sensor-names/sensor-name/thresholds/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("value-detailed", types.YChild{"ValueDetailed", &threshold.ValueDetailed})
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("type", types.YLeaf{"Type", threshold.Type})
    threshold.EntityData.Leafs.Append("trap", types.YLeaf{"Trap", threshold.Trap})
    threshold.EntityData.Leafs.Append("value-brief", types.YLeaf{"ValueBrief", threshold.ValueBrief})

    threshold.EntityData.YListKeys = []string {"Type"}

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
    valueDetailed.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/sensor-types/sensor-type/sensor-names/sensor-name/thresholds/threshold/" + valueDetailed.EntityData.SegmentPath
    valueDetailed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    valueDetailed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    valueDetailed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    valueDetailed.EntityData.Children = types.NewOrderedMap()
    valueDetailed.EntityData.Leafs = types.NewOrderedMap()
    valueDetailed.EntityData.Leafs.Append("threshold-severity", types.YLeaf{"ThresholdSeverity", valueDetailed.ThresholdSeverity})
    valueDetailed.EntityData.Leafs.Append("threshold-relation", types.YLeaf{"ThresholdRelation", valueDetailed.ThresholdRelation})
    valueDetailed.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", valueDetailed.ThresholdValue})
    valueDetailed.EntityData.Leafs.Append("threshold-evaluation", types.YLeaf{"ThresholdEvaluation", valueDetailed.ThresholdEvaluation})
    valueDetailed.EntityData.Leafs.Append("threshold-notification-enabled", types.YLeaf{"ThresholdNotificationEnabled", valueDetailed.ThresholdNotificationEnabled})

    valueDetailed.EntityData.YListKeys = []string {}

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

    // Average sensor value over time interval. The type is interface{} with
    // range: -2147483648..2147483647.
    Average interface{}

    // Minimum Sensor value over time interval. The type is interface{} with
    // range: -2147483648..2147483647.
    Minimum interface{}

    // Maximum Sensor value over time interval. The type is interface{} with
    // range: -2147483648..2147483647.
    Maximum interface{}

    // Time Interval over which sensor value is monitored. The type is interface{}
    // with range: -2147483648..2147483647.
    Interval interface{}
}

func (valueDetailed *EnvironmentalMonitoring_Racks_Rack_Slots_Slot_Modules_Module_SensorTypes_SensorType_SensorNames_SensorName_ValueDetailed) GetEntityData() *types.CommonEntityData {
    valueDetailed.EntityData.YFilter = valueDetailed.YFilter
    valueDetailed.EntityData.YangName = "value-detailed"
    valueDetailed.EntityData.BundleName = "cisco_ios_xr"
    valueDetailed.EntityData.ParentYangName = "sensor-name"
    valueDetailed.EntityData.SegmentPath = "value-detailed"
    valueDetailed.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-envmon-oper:environmental-monitoring/racks/rack/slots/slot/modules/module/sensor-types/sensor-type/sensor-names/sensor-name/" + valueDetailed.EntityData.SegmentPath
    valueDetailed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    valueDetailed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    valueDetailed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    valueDetailed.EntityData.Children = types.NewOrderedMap()
    valueDetailed.EntityData.Leafs = types.NewOrderedMap()
    valueDetailed.EntityData.Leafs.Append("field-validity-bitmap", types.YLeaf{"FieldValidityBitmap", valueDetailed.FieldValidityBitmap})
    valueDetailed.EntityData.Leafs.Append("device-description", types.YLeaf{"DeviceDescription", valueDetailed.DeviceDescription})
    valueDetailed.EntityData.Leafs.Append("units", types.YLeaf{"Units", valueDetailed.Units})
    valueDetailed.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", valueDetailed.DeviceId})
    valueDetailed.EntityData.Leafs.Append("value", types.YLeaf{"Value", valueDetailed.Value})
    valueDetailed.EntityData.Leafs.Append("alarm-type", types.YLeaf{"AlarmType", valueDetailed.AlarmType})
    valueDetailed.EntityData.Leafs.Append("data-type", types.YLeaf{"DataType", valueDetailed.DataType})
    valueDetailed.EntityData.Leafs.Append("scale", types.YLeaf{"Scale", valueDetailed.Scale})
    valueDetailed.EntityData.Leafs.Append("precision", types.YLeaf{"Precision", valueDetailed.Precision})
    valueDetailed.EntityData.Leafs.Append("status", types.YLeaf{"Status", valueDetailed.Status})
    valueDetailed.EntityData.Leafs.Append("age-time-stamp", types.YLeaf{"AgeTimeStamp", valueDetailed.AgeTimeStamp})
    valueDetailed.EntityData.Leafs.Append("update-rate", types.YLeaf{"UpdateRate", valueDetailed.UpdateRate})
    valueDetailed.EntityData.Leafs.Append("average", types.YLeaf{"Average", valueDetailed.Average})
    valueDetailed.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", valueDetailed.Minimum})
    valueDetailed.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", valueDetailed.Maximum})
    valueDetailed.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", valueDetailed.Interval})

    valueDetailed.EntityData.YListKeys = []string {}

    return &(valueDetailed.EntityData)
}

