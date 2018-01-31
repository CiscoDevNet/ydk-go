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
    SensorUnitsType_Watts SensorUnitsType = "Watts"

    SensorUnitsType_Celsius SensorUnitsType = "Celsius"

    SensorUnitsType_milliVolts SensorUnitsType = "milliVolts"

    SensorUnitsType_Amperes SensorUnitsType = "Amperes"

    SensorUnitsType_Volts_DC SensorUnitsType = "Volts-DC"

    SensorUnitsType_Volts_AC SensorUnitsType = "Volts-AC"

    SensorUnitsType_milliAmperes SensorUnitsType = "milliAmperes"
)

// EnvironmentSensors
// Data nodes for Environmental Monitoring
type EnvironmentSensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of components on the device chasis. The type is slice of
    // EnvironmentSensors_EnvironmentSensor.
    EnvironmentSensor []EnvironmentSensors_EnvironmentSensor
}

func (environmentSensors *EnvironmentSensors) GetFilter() yfilter.YFilter { return environmentSensors.YFilter }

func (environmentSensors *EnvironmentSensors) SetFilter(yf yfilter.YFilter) { environmentSensors.YFilter = yf }

func (environmentSensors *EnvironmentSensors) GetGoName(yname string) string {
    if yname == "environment-sensor" { return "EnvironmentSensor" }
    return ""
}

func (environmentSensors *EnvironmentSensors) GetSegmentPath() string {
    return "Cisco-IOS-XE-environment-oper:environment-sensors"
}

func (environmentSensors *EnvironmentSensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "environment-sensor" {
        for _, c := range environmentSensors.EnvironmentSensor {
            if environmentSensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EnvironmentSensors_EnvironmentSensor{}
        environmentSensors.EnvironmentSensor = append(environmentSensors.EnvironmentSensor, child)
        return &environmentSensors.EnvironmentSensor[len(environmentSensors.EnvironmentSensor)-1]
    }
    return nil
}

func (environmentSensors *EnvironmentSensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range environmentSensors.EnvironmentSensor {
        children[environmentSensors.EnvironmentSensor[i].GetSegmentPath()] = &environmentSensors.EnvironmentSensor[i]
    }
    return children
}

func (environmentSensors *EnvironmentSensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (environmentSensors *EnvironmentSensors) GetBundleName() string { return "cisco_ios_xe" }

func (environmentSensors *EnvironmentSensors) GetYangName() string { return "environment-sensors" }

func (environmentSensors *EnvironmentSensors) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (environmentSensors *EnvironmentSensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (environmentSensors *EnvironmentSensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (environmentSensors *EnvironmentSensors) SetParent(parent types.Entity) { environmentSensors.parent = parent }

func (environmentSensors *EnvironmentSensors) GetParent() types.Entity { return environmentSensors.parent }

func (environmentSensors *EnvironmentSensors) GetParentYangName() string { return "Cisco-IOS-XE-environment-oper" }

// EnvironmentSensors_EnvironmentSensor
// The list of components on the device chasis
type EnvironmentSensors_EnvironmentSensor struct {
    parent types.Entity
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
}

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetFilter() yfilter.YFilter { return environmentSensor.YFilter }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) SetFilter(yf yfilter.YFilter) { environmentSensor.YFilter = yf }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "location" { return "Location" }
    if yname == "state" { return "State" }
    if yname == "current-reading" { return "CurrentReading" }
    if yname == "sensor-units" { return "SensorUnits" }
    return ""
}

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetSegmentPath() string {
    return "environment-sensor" + "[name='" + fmt.Sprintf("%v", environmentSensor.Name) + "']" + "[location='" + fmt.Sprintf("%v", environmentSensor.Location) + "']"
}

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = environmentSensor.Name
    leafs["location"] = environmentSensor.Location
    leafs["state"] = environmentSensor.State
    leafs["current-reading"] = environmentSensor.CurrentReading
    leafs["sensor-units"] = environmentSensor.SensorUnits
    return leafs
}

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetBundleName() string { return "cisco_ios_xe" }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetYangName() string { return "environment-sensor" }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) SetParent(parent types.Entity) { environmentSensor.parent = parent }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetParent() types.Entity { return environmentSensor.parent }

func (environmentSensor *EnvironmentSensors_EnvironmentSensor) GetParentYangName() string { return "environment-sensors" }

