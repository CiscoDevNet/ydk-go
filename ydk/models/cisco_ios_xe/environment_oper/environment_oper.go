// This module contains a collection of YANG definitions
// for monitoring Environment of a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package environment_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package environment_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-environment-oper environment-sensors}", reflect.TypeOf(EnvironmentSensors{}))
    ydk.RegisterEntity("Cisco-IOS-XE-environment-oper:environment-sensors", reflect.TypeOf(EnvironmentSensors{}))
}

// SensorUnitsType represents Units used by various sensors
type SensorUnitsType string

const (
    SensorUnitsType_watts SensorUnitsType = "watts"

    SensorUnitsType_celsius SensorUnitsType = "celsius"

    SensorUnitsType_millivolts SensorUnitsType = "millivolts"

    SensorUnitsType_amperes SensorUnitsType = "amperes"

    SensorUnitsType_volts_dc SensorUnitsType = "volts-dc"

    SensorUnitsType_volts_ac SensorUnitsType = "volts-ac"

    SensorUnitsType_milliamperes SensorUnitsType = "milliamperes"

    SensorUnitsType_unknown SensorUnitsType = "unknown"

    SensorUnitsType_revolutions_per_minute SensorUnitsType = "revolutions-per-minute"
)

// EnvironmentSensors
// Data nodes for Environmental Monitoring
type EnvironmentSensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of components on the device chasis. The type is slice of
    // EnvironmentSensors_EnvironmentSensor.
    EnvironmentSensor []*EnvironmentSensors_EnvironmentSensor
}

func (environmentSensors *EnvironmentSensors) GetEntityData() *types.CommonEntityData {
    environmentSensors.EntityData.YFilter = environmentSensors.YFilter
    environmentSensors.EntityData.YangName = "environment-sensors"
    environmentSensors.EntityData.BundleName = "cisco_ios_xe"
    environmentSensors.EntityData.ParentYangName = "Cisco-IOS-XE-environment-oper"
    environmentSensors.EntityData.SegmentPath = "Cisco-IOS-XE-environment-oper:environment-sensors"
    environmentSensors.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    environmentSensors.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    environmentSensors.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    environmentSensors.EntityData.Children = types.NewOrderedMap()
    environmentSensors.EntityData.Children.Append("environment-sensor", types.YChild{"EnvironmentSensor", nil})
    for i := range environmentSensors.EnvironmentSensor {
        environmentSensors.EntityData.Children.Append(types.GetSegmentPath(environmentSensors.EnvironmentSensor[i]), types.YChild{"EnvironmentSensor", environmentSensors.EnvironmentSensor[i]})
    }
    environmentSensors.EntityData.Leafs = types.NewOrderedMap()

    environmentSensors.EntityData.YListKeys = []string {}

    return &(environmentSensors.EntityData)
}

// EnvironmentSensors_EnvironmentSensor
// The list of components on the device chasis
type EnvironmentSensors_EnvironmentSensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the sensor component. This includes all
    // physical components of the chasis - both fixed and pluggable. The type is
    // string.
    Name interface{}

    // This attribute is a key. Sensor location. The type is string.
    Location interface{}

    // A description of the current state of the sensor. The type is string.
    State interface{}

    // Numerical value of the current sensor reading in sensor-units. The type is
    // interface{} with range: 0..4294967295.
    CurrentReading interface{}

    // Units corresponding to the current-reading value. The type is
    // SensorUnitsType.
    SensorUnits interface{}

    // Alarm threshold under which a critical alarm will be signaled. The type is
    // interface{} with range: -2147483648..2147483647.
    LowCriticalThreshold interface{}

    // No alarm above this threshold. The type is interface{} with range:
    // -2147483648..2147483647.
    LowNormalThreshold interface{}

    // No alarm below this threshold. The type is interface{} with range:
    // -2147483648..2147483647.
    HighNormalThreshold interface{}

    // Alarm threshold over which a critical  alarm will be signaled. The type is
    // interface{} with range: -2147483648..2147483647.
    HighCriticalThreshold interface{}
}

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetEntityData() *types.CommonEntityData {
    environmentSensor.EntityData.YFilter = environmentSensor.YFilter
    environmentSensor.EntityData.YangName = "environment-sensor"
    environmentSensor.EntityData.BundleName = "cisco_ios_xe"
    environmentSensor.EntityData.ParentYangName = "environment-sensors"
    environmentSensor.EntityData.SegmentPath = "environment-sensor" + types.AddKeyToken(environmentSensor.Name, "name") + types.AddKeyToken(environmentSensor.Location, "location")
    environmentSensor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    environmentSensor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    environmentSensor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    environmentSensor.EntityData.Children = types.NewOrderedMap()
    environmentSensor.EntityData.Leafs = types.NewOrderedMap()
    environmentSensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", environmentSensor.Name})
    environmentSensor.EntityData.Leafs.Append("location", types.YLeaf{"Location", environmentSensor.Location})
    environmentSensor.EntityData.Leafs.Append("state", types.YLeaf{"State", environmentSensor.State})
    environmentSensor.EntityData.Leafs.Append("current-reading", types.YLeaf{"CurrentReading", environmentSensor.CurrentReading})
    environmentSensor.EntityData.Leafs.Append("sensor-units", types.YLeaf{"SensorUnits", environmentSensor.SensorUnits})
    environmentSensor.EntityData.Leafs.Append("low-critical-threshold", types.YLeaf{"LowCriticalThreshold", environmentSensor.LowCriticalThreshold})
    environmentSensor.EntityData.Leafs.Append("low-normal-threshold", types.YLeaf{"LowNormalThreshold", environmentSensor.LowNormalThreshold})
    environmentSensor.EntityData.Leafs.Append("high-normal-threshold", types.YLeaf{"HighNormalThreshold", environmentSensor.HighNormalThreshold})
    environmentSensor.EntityData.Leafs.Append("high-critical-threshold", types.YLeaf{"HighCriticalThreshold", environmentSensor.HighCriticalThreshold})

    environmentSensor.EntityData.YListKeys = []string {"Name", "Location"}

    return &(environmentSensor.EntityData)
}

