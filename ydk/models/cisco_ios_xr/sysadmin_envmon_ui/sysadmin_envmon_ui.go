// This module contains definitions
// for the Calvados model objects.
// 
// This module holds chassis, cards Enviroment data.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_envmon_ui

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_envmon_ui"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-envmon-ui environment}", reflect.TypeOf(Environment{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-envmon-ui:environment", reflect.TypeOf(Environment{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-envmon-ui power-mgmt}", reflect.TypeOf(PowerMgmt{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt", reflect.TypeOf(PowerMgmt{}))
}

// Environment
type Environment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // environment operational (show) data.
    Oper Environment_Oper

    
    All Environment_All

    // environment configurational data.
    Config Environment_Config

    // show traceable processes. The type is slice of Environment_Trace.
    Trace []Environment_Trace
}

func (environment *Environment) GetEntityData() *types.CommonEntityData {
    environment.EntityData.YFilter = environment.YFilter
    environment.EntityData.YangName = "environment"
    environment.EntityData.BundleName = "cisco_ios_xr"
    environment.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-envmon-ui"
    environment.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment"
    environment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environment.EntityData.Children = make(map[string]types.YChild)
    environment.EntityData.Children["oper"] = types.YChild{"Oper", &environment.Oper}
    environment.EntityData.Children["all"] = types.YChild{"All", &environment.All}
    environment.EntityData.Children["config"] = types.YChild{"Config", &environment.Config}
    environment.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range environment.Trace {
        environment.EntityData.Children[types.GetSegmentPath(&environment.Trace[i])] = types.YChild{"Trace", &environment.Trace[i]}
    }
    environment.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(environment.EntityData)
}

// Environment_Oper
// environment operational (show) data
type Environment_Oper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Temperatures Environment_Oper_Temperatures

    
    Voltages Environment_Oper_Voltages

    
    Current Environment_Oper_Current

    
    Fan Environment_Oper_Fan

    
    Power Environment_Oper_Power

    
    Altitude Environment_Oper_Altitude
}

func (oper *Environment_Oper) GetEntityData() *types.CommonEntityData {
    oper.EntityData.YFilter = oper.YFilter
    oper.EntityData.YangName = "oper"
    oper.EntityData.BundleName = "cisco_ios_xr"
    oper.EntityData.ParentYangName = "environment"
    oper.EntityData.SegmentPath = "oper"
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = make(map[string]types.YChild)
    oper.EntityData.Children["temperatures"] = types.YChild{"Temperatures", &oper.Temperatures}
    oper.EntityData.Children["voltages"] = types.YChild{"Voltages", &oper.Voltages}
    oper.EntityData.Children["current"] = types.YChild{"Current", &oper.Current}
    oper.EntityData.Children["fan"] = types.YChild{"Fan", &oper.Fan}
    oper.EntityData.Children["power"] = types.YChild{"Power", &oper.Power}
    oper.EntityData.Children["altitude"] = types.YChild{"Altitude", &oper.Altitude}
    oper.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oper.EntityData)
}

// Environment_Oper_Temperatures
type Environment_Oper_Temperatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Temperatures_Location.
    Location []Environment_Oper_Temperatures_Location
}

func (temperatures *Environment_Oper_Temperatures) GetEntityData() *types.CommonEntityData {
    temperatures.EntityData.YFilter = temperatures.YFilter
    temperatures.EntityData.YangName = "temperatures"
    temperatures.EntityData.BundleName = "cisco_ios_xr"
    temperatures.EntityData.ParentYangName = "oper"
    temperatures.EntityData.SegmentPath = "temperatures"
    temperatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    temperatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    temperatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    temperatures.EntityData.Children = make(map[string]types.YChild)
    temperatures.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range temperatures.Location {
        temperatures.EntityData.Children[types.GetSegmentPath(&temperatures.Location[i])] = types.YChild{"Location", &temperatures.Location[i]}
    }
    temperatures.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(temperatures.EntityData)
}

// Environment_Oper_Temperatures_Location
type Environment_Oper_Temperatures_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of
    // Environment_Oper_Temperatures_Location_SensorAttributes.
    SensorAttributes []Environment_Oper_Temperatures_Location_SensorAttributes
}

func (location *Environment_Oper_Temperatures_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "temperatures"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["sensor_attributes"] = types.YChild{"SensorAttributes", nil}
    for i := range location.SensorAttributes {
        location.EntityData.Children[types.GetSegmentPath(&location.SensorAttributes[i])] = types.YChild{"SensorAttributes", &location.SensorAttributes[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Environment_Oper_Temperatures_Location_SensorAttributes
type Environment_Oper_Temperatures_Location_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is string.
    SensorId interface{}

    // The type is string.
    Alarm interface{}

    // The type is string.
    TemperatureValue interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    Value interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalHi interface{}
}

func (sensorAttributes *Environment_Oper_Temperatures_Location_SensorAttributes) GetEntityData() *types.CommonEntityData {
    sensorAttributes.EntityData.YFilter = sensorAttributes.YFilter
    sensorAttributes.EntityData.YangName = "sensor_attributes"
    sensorAttributes.EntityData.BundleName = "cisco_ios_xr"
    sensorAttributes.EntityData.ParentYangName = "location"
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + "[sensor='" + fmt.Sprintf("%v", sensorAttributes.Sensor) + "']"
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = make(map[string]types.YChild)
    sensorAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorAttributes.EntityData.Leafs["sensor"] = types.YLeaf{"Sensor", sensorAttributes.Sensor}
    sensorAttributes.EntityData.Leafs["sensor_id"] = types.YLeaf{"SensorId", sensorAttributes.SensorId}
    sensorAttributes.EntityData.Leafs["alarm"] = types.YLeaf{"Alarm", sensorAttributes.Alarm}
    sensorAttributes.EntityData.Leafs["temperature_value"] = types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue}
    sensorAttributes.EntityData.Leafs["value"] = types.YLeaf{"Value", sensorAttributes.Value}
    sensorAttributes.EntityData.Leafs["critical_lo"] = types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo}
    sensorAttributes.EntityData.Leafs["major_lo"] = types.YLeaf{"MajorLo", sensorAttributes.MajorLo}
    sensorAttributes.EntityData.Leafs["minor_lo"] = types.YLeaf{"MinorLo", sensorAttributes.MinorLo}
    sensorAttributes.EntityData.Leafs["minor_hi"] = types.YLeaf{"MinorHi", sensorAttributes.MinorHi}
    sensorAttributes.EntityData.Leafs["major_hi"] = types.YLeaf{"MajorHi", sensorAttributes.MajorHi}
    sensorAttributes.EntityData.Leafs["critical_hi"] = types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi}
    return &(sensorAttributes.EntityData)
}

// Environment_Oper_Voltages
type Environment_Oper_Voltages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Voltages_Location.
    Location []Environment_Oper_Voltages_Location
}

func (voltages *Environment_Oper_Voltages) GetEntityData() *types.CommonEntityData {
    voltages.EntityData.YFilter = voltages.YFilter
    voltages.EntityData.YangName = "voltages"
    voltages.EntityData.BundleName = "cisco_ios_xr"
    voltages.EntityData.ParentYangName = "oper"
    voltages.EntityData.SegmentPath = "voltages"
    voltages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voltages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voltages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voltages.EntityData.Children = make(map[string]types.YChild)
    voltages.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range voltages.Location {
        voltages.EntityData.Children[types.GetSegmentPath(&voltages.Location[i])] = types.YChild{"Location", &voltages.Location[i]}
    }
    voltages.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(voltages.EntityData)
}

// Environment_Oper_Voltages_Location
type Environment_Oper_Voltages_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_Oper_Voltages_Location_SensorAttributes.
    SensorAttributes []Environment_Oper_Voltages_Location_SensorAttributes
}

func (location *Environment_Oper_Voltages_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "voltages"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["sensor_attributes"] = types.YChild{"SensorAttributes", nil}
    for i := range location.SensorAttributes {
        location.EntityData.Children[types.GetSegmentPath(&location.SensorAttributes[i])] = types.YChild{"SensorAttributes", &location.SensorAttributes[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Environment_Oper_Voltages_Location_SensorAttributes
type Environment_Oper_Voltages_Location_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is string.
    SensorId interface{}

    // The type is string.
    Alarm interface{}

    // The type is string.
    Value interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    TemperatureValue interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalHi interface{}
}

func (sensorAttributes *Environment_Oper_Voltages_Location_SensorAttributes) GetEntityData() *types.CommonEntityData {
    sensorAttributes.EntityData.YFilter = sensorAttributes.YFilter
    sensorAttributes.EntityData.YangName = "sensor_attributes"
    sensorAttributes.EntityData.BundleName = "cisco_ios_xr"
    sensorAttributes.EntityData.ParentYangName = "location"
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + "[sensor='" + fmt.Sprintf("%v", sensorAttributes.Sensor) + "']"
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = make(map[string]types.YChild)
    sensorAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorAttributes.EntityData.Leafs["sensor"] = types.YLeaf{"Sensor", sensorAttributes.Sensor}
    sensorAttributes.EntityData.Leafs["sensor_id"] = types.YLeaf{"SensorId", sensorAttributes.SensorId}
    sensorAttributes.EntityData.Leafs["alarm"] = types.YLeaf{"Alarm", sensorAttributes.Alarm}
    sensorAttributes.EntityData.Leafs["value"] = types.YLeaf{"Value", sensorAttributes.Value}
    sensorAttributes.EntityData.Leafs["temperature_value"] = types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue}
    sensorAttributes.EntityData.Leafs["critical_lo"] = types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo}
    sensorAttributes.EntityData.Leafs["major_lo"] = types.YLeaf{"MajorLo", sensorAttributes.MajorLo}
    sensorAttributes.EntityData.Leafs["minor_lo"] = types.YLeaf{"MinorLo", sensorAttributes.MinorLo}
    sensorAttributes.EntityData.Leafs["minor_hi"] = types.YLeaf{"MinorHi", sensorAttributes.MinorHi}
    sensorAttributes.EntityData.Leafs["major_hi"] = types.YLeaf{"MajorHi", sensorAttributes.MajorHi}
    sensorAttributes.EntityData.Leafs["critical_hi"] = types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi}
    return &(sensorAttributes.EntityData)
}

// Environment_Oper_Current
type Environment_Oper_Current struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Current_Location.
    Location []Environment_Oper_Current_Location
}

func (current *Environment_Oper_Current) GetEntityData() *types.CommonEntityData {
    current.EntityData.YFilter = current.YFilter
    current.EntityData.YangName = "current"
    current.EntityData.BundleName = "cisco_ios_xr"
    current.EntityData.ParentYangName = "oper"
    current.EntityData.SegmentPath = "current"
    current.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    current.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    current.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    current.EntityData.Children = make(map[string]types.YChild)
    current.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range current.Location {
        current.EntityData.Children[types.GetSegmentPath(&current.Location[i])] = types.YChild{"Location", &current.Location[i]}
    }
    current.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(current.EntityData)
}

// Environment_Oper_Current_Location
type Environment_Oper_Current_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_Oper_Current_Location_SensorAttributes.
    SensorAttributes []Environment_Oper_Current_Location_SensorAttributes
}

func (location *Environment_Oper_Current_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "current"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["sensor_attributes"] = types.YChild{"SensorAttributes", nil}
    for i := range location.SensorAttributes {
        location.EntityData.Children[types.GetSegmentPath(&location.SensorAttributes[i])] = types.YChild{"SensorAttributes", &location.SensorAttributes[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Environment_Oper_Current_Location_SensorAttributes
type Environment_Oper_Current_Location_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is string.
    SensorId interface{}

    // The type is string.
    Alarm interface{}

    // The type is string.
    Value interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    TemperatureValue interface{}
}

func (sensorAttributes *Environment_Oper_Current_Location_SensorAttributes) GetEntityData() *types.CommonEntityData {
    sensorAttributes.EntityData.YFilter = sensorAttributes.YFilter
    sensorAttributes.EntityData.YangName = "sensor_attributes"
    sensorAttributes.EntityData.BundleName = "cisco_ios_xr"
    sensorAttributes.EntityData.ParentYangName = "location"
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + "[sensor='" + fmt.Sprintf("%v", sensorAttributes.Sensor) + "']"
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = make(map[string]types.YChild)
    sensorAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorAttributes.EntityData.Leafs["sensor"] = types.YLeaf{"Sensor", sensorAttributes.Sensor}
    sensorAttributes.EntityData.Leafs["sensor_id"] = types.YLeaf{"SensorId", sensorAttributes.SensorId}
    sensorAttributes.EntityData.Leafs["alarm"] = types.YLeaf{"Alarm", sensorAttributes.Alarm}
    sensorAttributes.EntityData.Leafs["value"] = types.YLeaf{"Value", sensorAttributes.Value}
    sensorAttributes.EntityData.Leafs["temperature_value"] = types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue}
    return &(sensorAttributes.EntityData)
}

// Environment_Oper_Fan
type Environment_Oper_Fan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Fan_Location.
    Location []Environment_Oper_Fan_Location
}

func (fan *Environment_Oper_Fan) GetEntityData() *types.CommonEntityData {
    fan.EntityData.YFilter = fan.YFilter
    fan.EntityData.YangName = "fan"
    fan.EntityData.BundleName = "cisco_ios_xr"
    fan.EntityData.ParentYangName = "oper"
    fan.EntityData.SegmentPath = "fan"
    fan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fan.EntityData.Children = make(map[string]types.YChild)
    fan.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range fan.Location {
        fan.EntityData.Children[types.GetSegmentPath(&fan.Location[i])] = types.YChild{"Location", &fan.Location[i]}
    }
    fan.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fan.EntityData)
}

// Environment_Oper_Fan_Location
type Environment_Oper_Fan_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_Oper_Fan_Location_FanAttributes.
    FanAttributes []Environment_Oper_Fan_Location_FanAttributes
}

func (location *Environment_Oper_Fan_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "fan"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["fan_attributes"] = types.YChild{"FanAttributes", nil}
    for i := range location.FanAttributes {
        location.EntityData.Children[types.GetSegmentPath(&location.FanAttributes[i])] = types.YChild{"FanAttributes", &location.FanAttributes[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Environment_Oper_Fan_Location_FanAttributes
type Environment_Oper_Fan_Location_FanAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LogicalSlot interface{}

    // The type is string.
    PrintFanHeader interface{}

    // The type is string.
    Location interface{}

    // The type is string.
    FruPid interface{}

    // The type is string.
    FansSpeed interface{}

    // The type is interface{} with range: 0..4294967295.
    FanHeader interface{}

    // The type is interface{} with range: 0..4294967295.
    SpeedSpace interface{}
}

func (fanAttributes *Environment_Oper_Fan_Location_FanAttributes) GetEntityData() *types.CommonEntityData {
    fanAttributes.EntityData.YFilter = fanAttributes.YFilter
    fanAttributes.EntityData.YangName = "fan_attributes"
    fanAttributes.EntityData.BundleName = "cisco_ios_xr"
    fanAttributes.EntityData.ParentYangName = "location"
    fanAttributes.EntityData.SegmentPath = "fan_attributes" + "[logical_slot='" + fmt.Sprintf("%v", fanAttributes.LogicalSlot) + "']"
    fanAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanAttributes.EntityData.Children = make(map[string]types.YChild)
    fanAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    fanAttributes.EntityData.Leafs["logical_slot"] = types.YLeaf{"LogicalSlot", fanAttributes.LogicalSlot}
    fanAttributes.EntityData.Leafs["print_fan_header"] = types.YLeaf{"PrintFanHeader", fanAttributes.PrintFanHeader}
    fanAttributes.EntityData.Leafs["location"] = types.YLeaf{"Location", fanAttributes.Location}
    fanAttributes.EntityData.Leafs["fru_PID"] = types.YLeaf{"FruPid", fanAttributes.FruPid}
    fanAttributes.EntityData.Leafs["fans_speed"] = types.YLeaf{"FansSpeed", fanAttributes.FansSpeed}
    fanAttributes.EntityData.Leafs["fan_header"] = types.YLeaf{"FanHeader", fanAttributes.FanHeader}
    fanAttributes.EntityData.Leafs["speed_space"] = types.YLeaf{"SpeedSpace", fanAttributes.SpeedSpace}
    return &(fanAttributes.EntityData)
}

// Environment_Oper_Power
type Environment_Oper_Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Power_Location.
    Location []Environment_Oper_Power_Location
}

func (power *Environment_Oper_Power) GetEntityData() *types.CommonEntityData {
    power.EntityData.YFilter = power.YFilter
    power.EntityData.YangName = "power"
    power.EntityData.BundleName = "cisco_ios_xr"
    power.EntityData.ParentYangName = "oper"
    power.EntityData.SegmentPath = "power"
    power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    power.EntityData.Children = make(map[string]types.YChild)
    power.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range power.Location {
        power.EntityData.Children[types.GetSegmentPath(&power.Location[i])] = types.YChild{"Location", &power.Location[i]}
    }
    power.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(power.EntityData)
}

// Environment_Oper_Power_Location
type Environment_Oper_Power_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_Oper_Power_Location_PemAttributes.
    PemAttributes []Environment_Oper_Power_Location_PemAttributes
}

func (location *Environment_Oper_Power_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "power"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["pem_attributes"] = types.YChild{"PemAttributes", nil}
    for i := range location.PemAttributes {
        location.EntityData.Children[types.GetSegmentPath(&location.PemAttributes[i])] = types.YChild{"PemAttributes", &location.PemAttributes[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Environment_Oper_Power_Location_PemAttributes
type Environment_Oper_Power_Location_PemAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Pem interface{}

    // The type is string.
    PemId interface{}

    // The type is string.
    CardType interface{}

    // The type is string.
    PsType interface{}

    // The type is interface{} with range: 0..65535.
    ShelfNum interface{}

    // The type is string.
    SupplyType interface{}

    // The type is string.
    InputVoltage interface{}

    // The type is string.
    InputCurrent interface{}

    // The type is string.
    OutputVoltage interface{}

    // The type is string.
    OutputCurrent interface{}

    // The type is string.
    Status interface{}

    // The type is interface{} with range: 0..4294967295.
    InputPowerToPs interface{}

    // The type is string.
    InputCurrentToPs interface{}

    // The type is interface{} with range: 0..4294967295.
    OutputPowerFromPs interface{}

    // The type is string.
    OutputCurrentFromPs interface{}

    // The type is interface{} with range: 0..4294967295.
    PowerAllocated interface{}

    // The type is string.
    PowerConsumed interface{}

    // The type is string.
    PowerStatus interface{}

    // The type is string.
    ConfgdPowerRedundancyMode interface{}

    // The type is interface{} with range: 0..4294967295.
    UsablePowerCapacity interface{}

    // The type is interface{} with range: 0..4294967295.
    ProtectionPowerCapacity interface{}

    // The type is interface{} with range: 0..4294967295.
    PowerResrvAndAlloc interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemPowerUsed interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemPowerInput interface{}

    // The type is interface{} with range: 0..65535.
    PowerLevel interface{}

    // The type is interface{} with range: 0..65535.
    OutputHeader interface{}

    // The type is interface{} with range: 0..65535.
    OutputFooter interface{}

    // The type is interface{} with range: 0..65535.
    PsSumFooter interface{}
}

func (pemAttributes *Environment_Oper_Power_Location_PemAttributes) GetEntityData() *types.CommonEntityData {
    pemAttributes.EntityData.YFilter = pemAttributes.YFilter
    pemAttributes.EntityData.YangName = "pem_attributes"
    pemAttributes.EntityData.BundleName = "cisco_ios_xr"
    pemAttributes.EntityData.ParentYangName = "location"
    pemAttributes.EntityData.SegmentPath = "pem_attributes" + "[pem='" + fmt.Sprintf("%v", pemAttributes.Pem) + "']"
    pemAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pemAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pemAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pemAttributes.EntityData.Children = make(map[string]types.YChild)
    pemAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    pemAttributes.EntityData.Leafs["pem"] = types.YLeaf{"Pem", pemAttributes.Pem}
    pemAttributes.EntityData.Leafs["pem_id"] = types.YLeaf{"PemId", pemAttributes.PemId}
    pemAttributes.EntityData.Leafs["card_type"] = types.YLeaf{"CardType", pemAttributes.CardType}
    pemAttributes.EntityData.Leafs["ps_type"] = types.YLeaf{"PsType", pemAttributes.PsType}
    pemAttributes.EntityData.Leafs["shelf_num"] = types.YLeaf{"ShelfNum", pemAttributes.ShelfNum}
    pemAttributes.EntityData.Leafs["supply_type"] = types.YLeaf{"SupplyType", pemAttributes.SupplyType}
    pemAttributes.EntityData.Leafs["input_voltage"] = types.YLeaf{"InputVoltage", pemAttributes.InputVoltage}
    pemAttributes.EntityData.Leafs["input_current"] = types.YLeaf{"InputCurrent", pemAttributes.InputCurrent}
    pemAttributes.EntityData.Leafs["output_voltage"] = types.YLeaf{"OutputVoltage", pemAttributes.OutputVoltage}
    pemAttributes.EntityData.Leafs["output_current"] = types.YLeaf{"OutputCurrent", pemAttributes.OutputCurrent}
    pemAttributes.EntityData.Leafs["status"] = types.YLeaf{"Status", pemAttributes.Status}
    pemAttributes.EntityData.Leafs["input_power_to_ps"] = types.YLeaf{"InputPowerToPs", pemAttributes.InputPowerToPs}
    pemAttributes.EntityData.Leafs["input_current_to_ps"] = types.YLeaf{"InputCurrentToPs", pemAttributes.InputCurrentToPs}
    pemAttributes.EntityData.Leafs["output_power_from_ps"] = types.YLeaf{"OutputPowerFromPs", pemAttributes.OutputPowerFromPs}
    pemAttributes.EntityData.Leafs["output_current_from_ps"] = types.YLeaf{"OutputCurrentFromPs", pemAttributes.OutputCurrentFromPs}
    pemAttributes.EntityData.Leafs["power_allocated"] = types.YLeaf{"PowerAllocated", pemAttributes.PowerAllocated}
    pemAttributes.EntityData.Leafs["power_consumed"] = types.YLeaf{"PowerConsumed", pemAttributes.PowerConsumed}
    pemAttributes.EntityData.Leafs["power_status"] = types.YLeaf{"PowerStatus", pemAttributes.PowerStatus}
    pemAttributes.EntityData.Leafs["confgd_power_redundancy_mode"] = types.YLeaf{"ConfgdPowerRedundancyMode", pemAttributes.ConfgdPowerRedundancyMode}
    pemAttributes.EntityData.Leafs["usable_power_capacity"] = types.YLeaf{"UsablePowerCapacity", pemAttributes.UsablePowerCapacity}
    pemAttributes.EntityData.Leafs["protection_power_capacity"] = types.YLeaf{"ProtectionPowerCapacity", pemAttributes.ProtectionPowerCapacity}
    pemAttributes.EntityData.Leafs["power_resrv_and_alloc"] = types.YLeaf{"PowerResrvAndAlloc", pemAttributes.PowerResrvAndAlloc}
    pemAttributes.EntityData.Leafs["system_power_used"] = types.YLeaf{"SystemPowerUsed", pemAttributes.SystemPowerUsed}
    pemAttributes.EntityData.Leafs["system_power_input"] = types.YLeaf{"SystemPowerInput", pemAttributes.SystemPowerInput}
    pemAttributes.EntityData.Leafs["power_level"] = types.YLeaf{"PowerLevel", pemAttributes.PowerLevel}
    pemAttributes.EntityData.Leafs["output_header"] = types.YLeaf{"OutputHeader", pemAttributes.OutputHeader}
    pemAttributes.EntityData.Leafs["output_footer"] = types.YLeaf{"OutputFooter", pemAttributes.OutputFooter}
    pemAttributes.EntityData.Leafs["ps_sum_footer"] = types.YLeaf{"PsSumFooter", pemAttributes.PsSumFooter}
    return &(pemAttributes.EntityData)
}

// Environment_Oper_Altitude
type Environment_Oper_Altitude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Altitude_Location.
    Location []Environment_Oper_Altitude_Location
}

func (altitude *Environment_Oper_Altitude) GetEntityData() *types.CommonEntityData {
    altitude.EntityData.YFilter = altitude.YFilter
    altitude.EntityData.YangName = "altitude"
    altitude.EntityData.BundleName = "cisco_ios_xr"
    altitude.EntityData.ParentYangName = "oper"
    altitude.EntityData.SegmentPath = "altitude"
    altitude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altitude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altitude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altitude.EntityData.Children = make(map[string]types.YChild)
    altitude.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range altitude.Location {
        altitude.EntityData.Children[types.GetSegmentPath(&altitude.Location[i])] = types.YChild{"Location", &altitude.Location[i]}
    }
    altitude.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(altitude.EntityData)
}

// Environment_Oper_Altitude_Location
type Environment_Oper_Altitude_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_Oper_Altitude_Location_AltAttributes.
    AltAttributes []Environment_Oper_Altitude_Location_AltAttributes
}

func (location *Environment_Oper_Altitude_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "altitude"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["alt_attributes"] = types.YChild{"AltAttributes", nil}
    for i := range location.AltAttributes {
        location.EntityData.Children[types.GetSegmentPath(&location.AltAttributes[i])] = types.YChild{"AltAttributes", &location.AltAttributes[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Environment_Oper_Altitude_Location_AltAttributes
type Environment_Oper_Altitude_Location_AltAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is interface{} with range: 0..4294967295.
    Rack interface{}

    // The type is string.
    SensorValue interface{}

    // The type is string.
    Source interface{}
}

func (altAttributes *Environment_Oper_Altitude_Location_AltAttributes) GetEntityData() *types.CommonEntityData {
    altAttributes.EntityData.YFilter = altAttributes.YFilter
    altAttributes.EntityData.YangName = "alt_attributes"
    altAttributes.EntityData.BundleName = "cisco_ios_xr"
    altAttributes.EntityData.ParentYangName = "location"
    altAttributes.EntityData.SegmentPath = "alt_attributes" + "[sensor='" + fmt.Sprintf("%v", altAttributes.Sensor) + "']"
    altAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altAttributes.EntityData.Children = make(map[string]types.YChild)
    altAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    altAttributes.EntityData.Leafs["sensor"] = types.YLeaf{"Sensor", altAttributes.Sensor}
    altAttributes.EntityData.Leafs["rack"] = types.YLeaf{"Rack", altAttributes.Rack}
    altAttributes.EntityData.Leafs["sensor_value"] = types.YLeaf{"SensorValue", altAttributes.SensorValue}
    altAttributes.EntityData.Leafs["source"] = types.YLeaf{"Source", altAttributes.Source}
    return &(altAttributes.EntityData)
}

// Environment_All
type Environment_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_All_Location.
    Location []Environment_All_Location
}

func (all *Environment_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "environment"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range all.Location {
        all.EntityData.Children[types.GetSegmentPath(&all.Location[i])] = types.YChild{"Location", &all.Location[i]}
    }
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(all.EntityData)
}

// Environment_All_Location
type Environment_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_All_Location_Temperatures.
    Temperatures []Environment_All_Location_Temperatures

    // The type is slice of Environment_All_Location_Voltages.
    Voltages []Environment_All_Location_Voltages

    // The type is slice of Environment_All_Location_Current.
    Current []Environment_All_Location_Current

    // The type is slice of Environment_All_Location_Fan.
    Fan []Environment_All_Location_Fan
}

func (location *Environment_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["temperatures"] = types.YChild{"Temperatures", nil}
    for i := range location.Temperatures {
        location.EntityData.Children[types.GetSegmentPath(&location.Temperatures[i])] = types.YChild{"Temperatures", &location.Temperatures[i]}
    }
    location.EntityData.Children["voltages"] = types.YChild{"Voltages", nil}
    for i := range location.Voltages {
        location.EntityData.Children[types.GetSegmentPath(&location.Voltages[i])] = types.YChild{"Voltages", &location.Voltages[i]}
    }
    location.EntityData.Children["current"] = types.YChild{"Current", nil}
    for i := range location.Current {
        location.EntityData.Children[types.GetSegmentPath(&location.Current[i])] = types.YChild{"Current", &location.Current[i]}
    }
    location.EntityData.Children["fan"] = types.YChild{"Fan", nil}
    for i := range location.Fan {
        location.EntityData.Children[types.GetSegmentPath(&location.Fan[i])] = types.YChild{"Fan", &location.Fan[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Environment_All_Location_Temperatures
type Environment_All_Location_Temperatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is slice of
    // Environment_All_Location_Temperatures_SensorAttributes.
    SensorAttributes []Environment_All_Location_Temperatures_SensorAttributes
}

func (temperatures *Environment_All_Location_Temperatures) GetEntityData() *types.CommonEntityData {
    temperatures.EntityData.YFilter = temperatures.YFilter
    temperatures.EntityData.YangName = "temperatures"
    temperatures.EntityData.BundleName = "cisco_ios_xr"
    temperatures.EntityData.ParentYangName = "location"
    temperatures.EntityData.SegmentPath = "temperatures" + "[loc_iden='" + fmt.Sprintf("%v", temperatures.LocIden) + "']"
    temperatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    temperatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    temperatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    temperatures.EntityData.Children = make(map[string]types.YChild)
    temperatures.EntityData.Children["sensor_attributes"] = types.YChild{"SensorAttributes", nil}
    for i := range temperatures.SensorAttributes {
        temperatures.EntityData.Children[types.GetSegmentPath(&temperatures.SensorAttributes[i])] = types.YChild{"SensorAttributes", &temperatures.SensorAttributes[i]}
    }
    temperatures.EntityData.Leafs = make(map[string]types.YLeaf)
    temperatures.EntityData.Leafs["loc_iden"] = types.YLeaf{"LocIden", temperatures.LocIden}
    temperatures.EntityData.Leafs["print_header"] = types.YLeaf{"PrintHeader", temperatures.PrintHeader}
    return &(temperatures.EntityData)
}

// Environment_All_Location_Temperatures_SensorAttributes
type Environment_All_Location_Temperatures_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is string.
    SensorId interface{}

    // The type is string.
    Alarm interface{}

    // The type is string.
    TemperatureValue interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    Value interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalHi interface{}
}

func (sensorAttributes *Environment_All_Location_Temperatures_SensorAttributes) GetEntityData() *types.CommonEntityData {
    sensorAttributes.EntityData.YFilter = sensorAttributes.YFilter
    sensorAttributes.EntityData.YangName = "sensor_attributes"
    sensorAttributes.EntityData.BundleName = "cisco_ios_xr"
    sensorAttributes.EntityData.ParentYangName = "temperatures"
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + "[sensor='" + fmt.Sprintf("%v", sensorAttributes.Sensor) + "']"
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = make(map[string]types.YChild)
    sensorAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorAttributes.EntityData.Leafs["sensor"] = types.YLeaf{"Sensor", sensorAttributes.Sensor}
    sensorAttributes.EntityData.Leafs["sensor_id"] = types.YLeaf{"SensorId", sensorAttributes.SensorId}
    sensorAttributes.EntityData.Leafs["alarm"] = types.YLeaf{"Alarm", sensorAttributes.Alarm}
    sensorAttributes.EntityData.Leafs["temperature_value"] = types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue}
    sensorAttributes.EntityData.Leafs["value"] = types.YLeaf{"Value", sensorAttributes.Value}
    sensorAttributes.EntityData.Leafs["critical_lo"] = types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo}
    sensorAttributes.EntityData.Leafs["major_lo"] = types.YLeaf{"MajorLo", sensorAttributes.MajorLo}
    sensorAttributes.EntityData.Leafs["minor_lo"] = types.YLeaf{"MinorLo", sensorAttributes.MinorLo}
    sensorAttributes.EntityData.Leafs["minor_hi"] = types.YLeaf{"MinorHi", sensorAttributes.MinorHi}
    sensorAttributes.EntityData.Leafs["major_hi"] = types.YLeaf{"MajorHi", sensorAttributes.MajorHi}
    sensorAttributes.EntityData.Leafs["critical_hi"] = types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi}
    return &(sensorAttributes.EntityData)
}

// Environment_All_Location_Voltages
type Environment_All_Location_Voltages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is slice of Environment_All_Location_Voltages_SensorAttributes.
    SensorAttributes []Environment_All_Location_Voltages_SensorAttributes
}

func (voltages *Environment_All_Location_Voltages) GetEntityData() *types.CommonEntityData {
    voltages.EntityData.YFilter = voltages.YFilter
    voltages.EntityData.YangName = "voltages"
    voltages.EntityData.BundleName = "cisco_ios_xr"
    voltages.EntityData.ParentYangName = "location"
    voltages.EntityData.SegmentPath = "voltages" + "[loc_iden='" + fmt.Sprintf("%v", voltages.LocIden) + "']"
    voltages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voltages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voltages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voltages.EntityData.Children = make(map[string]types.YChild)
    voltages.EntityData.Children["sensor_attributes"] = types.YChild{"SensorAttributes", nil}
    for i := range voltages.SensorAttributes {
        voltages.EntityData.Children[types.GetSegmentPath(&voltages.SensorAttributes[i])] = types.YChild{"SensorAttributes", &voltages.SensorAttributes[i]}
    }
    voltages.EntityData.Leafs = make(map[string]types.YLeaf)
    voltages.EntityData.Leafs["loc_iden"] = types.YLeaf{"LocIden", voltages.LocIden}
    voltages.EntityData.Leafs["print_header"] = types.YLeaf{"PrintHeader", voltages.PrintHeader}
    return &(voltages.EntityData)
}

// Environment_All_Location_Voltages_SensorAttributes
type Environment_All_Location_Voltages_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is string.
    SensorId interface{}

    // The type is string.
    Alarm interface{}

    // The type is string.
    Value interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    TemperatureValue interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorLo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MinorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    MajorHi interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    CriticalHi interface{}
}

func (sensorAttributes *Environment_All_Location_Voltages_SensorAttributes) GetEntityData() *types.CommonEntityData {
    sensorAttributes.EntityData.YFilter = sensorAttributes.YFilter
    sensorAttributes.EntityData.YangName = "sensor_attributes"
    sensorAttributes.EntityData.BundleName = "cisco_ios_xr"
    sensorAttributes.EntityData.ParentYangName = "voltages"
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + "[sensor='" + fmt.Sprintf("%v", sensorAttributes.Sensor) + "']"
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = make(map[string]types.YChild)
    sensorAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorAttributes.EntityData.Leafs["sensor"] = types.YLeaf{"Sensor", sensorAttributes.Sensor}
    sensorAttributes.EntityData.Leafs["sensor_id"] = types.YLeaf{"SensorId", sensorAttributes.SensorId}
    sensorAttributes.EntityData.Leafs["alarm"] = types.YLeaf{"Alarm", sensorAttributes.Alarm}
    sensorAttributes.EntityData.Leafs["value"] = types.YLeaf{"Value", sensorAttributes.Value}
    sensorAttributes.EntityData.Leafs["temperature_value"] = types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue}
    sensorAttributes.EntityData.Leafs["critical_lo"] = types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo}
    sensorAttributes.EntityData.Leafs["major_lo"] = types.YLeaf{"MajorLo", sensorAttributes.MajorLo}
    sensorAttributes.EntityData.Leafs["minor_lo"] = types.YLeaf{"MinorLo", sensorAttributes.MinorLo}
    sensorAttributes.EntityData.Leafs["minor_hi"] = types.YLeaf{"MinorHi", sensorAttributes.MinorHi}
    sensorAttributes.EntityData.Leafs["major_hi"] = types.YLeaf{"MajorHi", sensorAttributes.MajorHi}
    sensorAttributes.EntityData.Leafs["critical_hi"] = types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi}
    return &(sensorAttributes.EntityData)
}

// Environment_All_Location_Current
type Environment_All_Location_Current struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is slice of Environment_All_Location_Current_SensorAttributes.
    SensorAttributes []Environment_All_Location_Current_SensorAttributes
}

func (current *Environment_All_Location_Current) GetEntityData() *types.CommonEntityData {
    current.EntityData.YFilter = current.YFilter
    current.EntityData.YangName = "current"
    current.EntityData.BundleName = "cisco_ios_xr"
    current.EntityData.ParentYangName = "location"
    current.EntityData.SegmentPath = "current" + "[loc_iden='" + fmt.Sprintf("%v", current.LocIden) + "']"
    current.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    current.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    current.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    current.EntityData.Children = make(map[string]types.YChild)
    current.EntityData.Children["sensor_attributes"] = types.YChild{"SensorAttributes", nil}
    for i := range current.SensorAttributes {
        current.EntityData.Children[types.GetSegmentPath(&current.SensorAttributes[i])] = types.YChild{"SensorAttributes", &current.SensorAttributes[i]}
    }
    current.EntityData.Leafs = make(map[string]types.YLeaf)
    current.EntityData.Leafs["loc_iden"] = types.YLeaf{"LocIden", current.LocIden}
    current.EntityData.Leafs["print_header"] = types.YLeaf{"PrintHeader", current.PrintHeader}
    return &(current.EntityData)
}

// Environment_All_Location_Current_SensorAttributes
type Environment_All_Location_Current_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is string.
    SensorId interface{}

    // The type is string.
    Value interface{}
}

func (sensorAttributes *Environment_All_Location_Current_SensorAttributes) GetEntityData() *types.CommonEntityData {
    sensorAttributes.EntityData.YFilter = sensorAttributes.YFilter
    sensorAttributes.EntityData.YangName = "sensor_attributes"
    sensorAttributes.EntityData.BundleName = "cisco_ios_xr"
    sensorAttributes.EntityData.ParentYangName = "current"
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + "[sensor='" + fmt.Sprintf("%v", sensorAttributes.Sensor) + "']"
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = make(map[string]types.YChild)
    sensorAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorAttributes.EntityData.Leafs["sensor"] = types.YLeaf{"Sensor", sensorAttributes.Sensor}
    sensorAttributes.EntityData.Leafs["sensor_id"] = types.YLeaf{"SensorId", sensorAttributes.SensorId}
    sensorAttributes.EntityData.Leafs["value"] = types.YLeaf{"Value", sensorAttributes.Value}
    return &(sensorAttributes.EntityData)
}

// Environment_All_Location_Fan
type Environment_All_Location_Fan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is slice of Environment_All_Location_Fan_FanAttributes.
    FanAttributes []Environment_All_Location_Fan_FanAttributes
}

func (fan *Environment_All_Location_Fan) GetEntityData() *types.CommonEntityData {
    fan.EntityData.YFilter = fan.YFilter
    fan.EntityData.YangName = "fan"
    fan.EntityData.BundleName = "cisco_ios_xr"
    fan.EntityData.ParentYangName = "location"
    fan.EntityData.SegmentPath = "fan" + "[loc_iden='" + fmt.Sprintf("%v", fan.LocIden) + "']"
    fan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fan.EntityData.Children = make(map[string]types.YChild)
    fan.EntityData.Children["fan_attributes"] = types.YChild{"FanAttributes", nil}
    for i := range fan.FanAttributes {
        fan.EntityData.Children[types.GetSegmentPath(&fan.FanAttributes[i])] = types.YChild{"FanAttributes", &fan.FanAttributes[i]}
    }
    fan.EntityData.Leafs = make(map[string]types.YLeaf)
    fan.EntityData.Leafs["loc_iden"] = types.YLeaf{"LocIden", fan.LocIden}
    return &(fan.EntityData)
}

// Environment_All_Location_Fan_FanAttributes
type Environment_All_Location_Fan_FanAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LogicalSlot interface{}

    // The type is string.
    PrintFanHeader interface{}

    // The type is string.
    Location interface{}

    // The type is string.
    FruPid interface{}

    // The type is string.
    FansSpeed interface{}

    // The type is interface{} with range: 0..4294967295.
    FanHeader interface{}

    // The type is interface{} with range: 0..4294967295.
    SpeedSpace interface{}
}

func (fanAttributes *Environment_All_Location_Fan_FanAttributes) GetEntityData() *types.CommonEntityData {
    fanAttributes.EntityData.YFilter = fanAttributes.YFilter
    fanAttributes.EntityData.YangName = "fan_attributes"
    fanAttributes.EntityData.BundleName = "cisco_ios_xr"
    fanAttributes.EntityData.ParentYangName = "fan"
    fanAttributes.EntityData.SegmentPath = "fan_attributes" + "[logical_slot='" + fmt.Sprintf("%v", fanAttributes.LogicalSlot) + "']"
    fanAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanAttributes.EntityData.Children = make(map[string]types.YChild)
    fanAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    fanAttributes.EntityData.Leafs["logical_slot"] = types.YLeaf{"LogicalSlot", fanAttributes.LogicalSlot}
    fanAttributes.EntityData.Leafs["print_fan_header"] = types.YLeaf{"PrintFanHeader", fanAttributes.PrintFanHeader}
    fanAttributes.EntityData.Leafs["location"] = types.YLeaf{"Location", fanAttributes.Location}
    fanAttributes.EntityData.Leafs["fru_PID"] = types.YLeaf{"FruPid", fanAttributes.FruPid}
    fanAttributes.EntityData.Leafs["fans_speed"] = types.YLeaf{"FansSpeed", fanAttributes.FansSpeed}
    fanAttributes.EntityData.Leafs["fan_header"] = types.YLeaf{"FanHeader", fanAttributes.FanHeader}
    fanAttributes.EntityData.Leafs["speed_space"] = types.YLeaf{"SpeedSpace", fanAttributes.SpeedSpace}
    return &(fanAttributes.EntityData)
}

// Environment_Config
// environment configurational data
type Environment_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..100.
    RaiseFanSpeed interface{}

    // The type is interface{} with range: 0..1.
    FanCtrlOptics interface{}

    // The type is interface{} with range: 0..1.
    GracefulShutdown interface{}

    
    Router Environment_Config_Router

    
    AirFilter Environment_Config_AirFilter

    
    FanCtrl Environment_Config_FanCtrl

    
    Temperature Environment_Config_Temperature

    
    Monitoring Environment_Config_Monitoring
}

func (config *Environment_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "environment"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["router"] = types.YChild{"Router", &config.Router}
    config.EntityData.Children["air-filter"] = types.YChild{"AirFilter", &config.AirFilter}
    config.EntityData.Children["fan-ctrl"] = types.YChild{"FanCtrl", &config.FanCtrl}
    config.EntityData.Children["temperature"] = types.YChild{"Temperature", &config.Temperature}
    config.EntityData.Children["monitoring"] = types.YChild{"Monitoring", &config.Monitoring}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["raise-fan-speed"] = types.YLeaf{"RaiseFanSpeed", config.RaiseFanSpeed}
    config.EntityData.Leafs["fan-ctrl-optics"] = types.YLeaf{"FanCtrlOptics", config.FanCtrlOptics}
    config.EntityData.Leafs["graceful-shutdown"] = types.YLeaf{"GracefulShutdown", config.GracefulShutdown}
    return &(config.EntityData)
}

// Environment_Config_Router
type Environment_Config_Router struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Altitude Environment_Config_Router_Altitude
}

func (router *Environment_Config_Router) GetEntityData() *types.CommonEntityData {
    router.EntityData.YFilter = router.YFilter
    router.EntityData.YangName = "router"
    router.EntityData.BundleName = "cisco_ios_xr"
    router.EntityData.ParentYangName = "config"
    router.EntityData.SegmentPath = "router"
    router.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    router.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    router.EntityData.Children = make(map[string]types.YChild)
    router.EntityData.Children["altitude"] = types.YChild{"Altitude", &router.Altitude}
    router.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(router.EntityData)
}

// Environment_Config_Router_Altitude
type Environment_Config_Router_Altitude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Environment_Config_Router_Altitude_All

    
    RackLoc Environment_Config_Router_Altitude_RackLoc
}

func (altitude *Environment_Config_Router_Altitude) GetEntityData() *types.CommonEntityData {
    altitude.EntityData.YFilter = altitude.YFilter
    altitude.EntityData.YangName = "altitude"
    altitude.EntityData.BundleName = "cisco_ios_xr"
    altitude.EntityData.ParentYangName = "router"
    altitude.EntityData.SegmentPath = "altitude"
    altitude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altitude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altitude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altitude.EntityData.Children = make(map[string]types.YChild)
    altitude.EntityData.Children["all"] = types.YChild{"All", &altitude.All}
    altitude.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &altitude.RackLoc}
    altitude.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(altitude.EntityData)
}

// Environment_Config_Router_Altitude_All
type Environment_Config_Router_Altitude_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 1..4000.
    NumMeters interface{}
}

func (all *Environment_Config_Router_Altitude_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "altitude"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    all.EntityData.Leafs["num_meters"] = types.YLeaf{"NumMeters", all.NumMeters}
    return &(all.EntityData)
}

// Environment_Config_Router_Altitude_RackLoc
type Environment_Config_Router_Altitude_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Config_Router_Altitude_RackLoc_Location.
    Location []Environment_Config_Router_Altitude_RackLoc_Location
}

func (rackLoc *Environment_Config_Router_Altitude_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "altitude"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rackLoc.EntityData)
}

// Environment_Config_Router_Altitude_RackLoc_Location
type Environment_Config_Router_Altitude_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}

    // The type is interface{} with range: 1..4000.
    NumMeters interface{}
}

func (location *Environment_Config_Router_Altitude_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    location.EntityData.Leafs["num_meters"] = types.YLeaf{"NumMeters", location.NumMeters}
    return &(location.EntityData)
}

// Environment_Config_AirFilter
type Environment_Config_AirFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Replaced Environment_Config_AirFilter_Replaced
}

func (airFilter *Environment_Config_AirFilter) GetEntityData() *types.CommonEntityData {
    airFilter.EntityData.YFilter = airFilter.YFilter
    airFilter.EntityData.YangName = "air-filter"
    airFilter.EntityData.BundleName = "cisco_ios_xr"
    airFilter.EntityData.ParentYangName = "config"
    airFilter.EntityData.SegmentPath = "air-filter"
    airFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    airFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    airFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    airFilter.EntityData.Children = make(map[string]types.YChild)
    airFilter.EntityData.Children["replaced"] = types.YChild{"Replaced", &airFilter.Replaced}
    airFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(airFilter.EntityData)
}

// Environment_Config_AirFilter_Replaced
type Environment_Config_AirFilter_Replaced struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Environment_Config_AirFilter_Replaced_All

    
    RackLoc Environment_Config_AirFilter_Replaced_RackLoc
}

func (replaced *Environment_Config_AirFilter_Replaced) GetEntityData() *types.CommonEntityData {
    replaced.EntityData.YFilter = replaced.YFilter
    replaced.EntityData.YangName = "replaced"
    replaced.EntityData.BundleName = "cisco_ios_xr"
    replaced.EntityData.ParentYangName = "air-filter"
    replaced.EntityData.SegmentPath = "replaced"
    replaced.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replaced.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replaced.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replaced.EntityData.Children = make(map[string]types.YChild)
    replaced.EntityData.Children["all"] = types.YChild{"All", &replaced.All}
    replaced.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &replaced.RackLoc}
    replaced.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(replaced.EntityData)
}

// Environment_Config_AirFilter_Replaced_All
type Environment_Config_AirFilter_Replaced_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Date interface{}
}

func (all *Environment_Config_AirFilter_Replaced_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "replaced"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    all.EntityData.Leafs["date"] = types.YLeaf{"Date", all.Date}
    return &(all.EntityData)
}

// Environment_Config_AirFilter_Replaced_RackLoc
type Environment_Config_AirFilter_Replaced_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Environment_Config_AirFilter_Replaced_RackLoc_Location.
    Location []Environment_Config_AirFilter_Replaced_RackLoc_Location
}

func (rackLoc *Environment_Config_AirFilter_Replaced_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "replaced"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rackLoc.EntityData)
}

// Environment_Config_AirFilter_Replaced_RackLoc_Location
type Environment_Config_AirFilter_Replaced_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}

    // The type is string.
    Date interface{}
}

func (location *Environment_Config_AirFilter_Replaced_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    location.EntityData.Leafs["date"] = types.YLeaf{"Date", location.Date}
    return &(location.EntityData)
}

// Environment_Config_FanCtrl
type Environment_Config_FanCtrl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Disable Environment_Config_FanCtrl_Disable
}

func (fanCtrl *Environment_Config_FanCtrl) GetEntityData() *types.CommonEntityData {
    fanCtrl.EntityData.YFilter = fanCtrl.YFilter
    fanCtrl.EntityData.YangName = "fan-ctrl"
    fanCtrl.EntityData.BundleName = "cisco_ios_xr"
    fanCtrl.EntityData.ParentYangName = "config"
    fanCtrl.EntityData.SegmentPath = "fan-ctrl"
    fanCtrl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanCtrl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanCtrl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanCtrl.EntityData.Children = make(map[string]types.YChild)
    fanCtrl.EntityData.Children["disable"] = types.YChild{"Disable", &fanCtrl.Disable}
    fanCtrl.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fanCtrl.EntityData)
}

// Environment_Config_FanCtrl_Disable
type Environment_Config_FanCtrl_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc Environment_Config_FanCtrl_Disable_RackLoc
}

func (disable *Environment_Config_FanCtrl_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "fan-ctrl"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = make(map[string]types.YChild)
    disable.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &disable.RackLoc}
    disable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(disable.EntityData)
}

// Environment_Config_FanCtrl_Disable_RackLoc
type Environment_Config_FanCtrl_Disable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of Environment_Config_FanCtrl_Disable_RackLoc_Location.
    Location []Environment_Config_FanCtrl_Disable_RackLoc_Location
}

func (rackLoc *Environment_Config_FanCtrl_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    rackLoc.EntityData.Leafs["all"] = types.YLeaf{"All", rackLoc.All}
    return &(rackLoc.EntityData)
}

// Environment_Config_FanCtrl_Disable_RackLoc_Location
type Environment_Config_FanCtrl_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}
}

func (location *Environment_Config_FanCtrl_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    return &(location.EntityData)
}

// Environment_Config_Temperature
type Environment_Config_Temperature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Disable Environment_Config_Temperature_Disable
}

func (temperature *Environment_Config_Temperature) GetEntityData() *types.CommonEntityData {
    temperature.EntityData.YFilter = temperature.YFilter
    temperature.EntityData.YangName = "temperature"
    temperature.EntityData.BundleName = "cisco_ios_xr"
    temperature.EntityData.ParentYangName = "config"
    temperature.EntityData.SegmentPath = "temperature"
    temperature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    temperature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    temperature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    temperature.EntityData.Children = make(map[string]types.YChild)
    temperature.EntityData.Children["disable"] = types.YChild{"Disable", &temperature.Disable}
    temperature.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(temperature.EntityData)
}

// Environment_Config_Temperature_Disable
type Environment_Config_Temperature_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc Environment_Config_Temperature_Disable_RackLoc
}

func (disable *Environment_Config_Temperature_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "temperature"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = make(map[string]types.YChild)
    disable.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &disable.RackLoc}
    disable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(disable.EntityData)
}

// Environment_Config_Temperature_Disable_RackLoc
type Environment_Config_Temperature_Disable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of
    // Environment_Config_Temperature_Disable_RackLoc_Location.
    Location []Environment_Config_Temperature_Disable_RackLoc_Location
}

func (rackLoc *Environment_Config_Temperature_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    rackLoc.EntityData.Leafs["all"] = types.YLeaf{"All", rackLoc.All}
    return &(rackLoc.EntityData)
}

// Environment_Config_Temperature_Disable_RackLoc_Location
type Environment_Config_Temperature_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}
}

func (location *Environment_Config_Temperature_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    return &(location.EntityData)
}

// Environment_Config_Monitoring
type Environment_Config_Monitoring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Disable Environment_Config_Monitoring_Disable
}

func (monitoring *Environment_Config_Monitoring) GetEntityData() *types.CommonEntityData {
    monitoring.EntityData.YFilter = monitoring.YFilter
    monitoring.EntityData.YangName = "monitoring"
    monitoring.EntityData.BundleName = "cisco_ios_xr"
    monitoring.EntityData.ParentYangName = "config"
    monitoring.EntityData.SegmentPath = "monitoring"
    monitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitoring.EntityData.Children = make(map[string]types.YChild)
    monitoring.EntityData.Children["disable"] = types.YChild{"Disable", &monitoring.Disable}
    monitoring.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(monitoring.EntityData)
}

// Environment_Config_Monitoring_Disable
type Environment_Config_Monitoring_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc Environment_Config_Monitoring_Disable_RackLoc
}

func (disable *Environment_Config_Monitoring_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "monitoring"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = make(map[string]types.YChild)
    disable.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &disable.RackLoc}
    disable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(disable.EntityData)
}

// Environment_Config_Monitoring_Disable_RackLoc
type Environment_Config_Monitoring_Disable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of
    // Environment_Config_Monitoring_Disable_RackLoc_Location.
    Location []Environment_Config_Monitoring_Disable_RackLoc_Location
}

func (rackLoc *Environment_Config_Monitoring_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    rackLoc.EntityData.Leafs["all"] = types.YLeaf{"All", rackLoc.All}
    return &(rackLoc.EntityData)
}

// Environment_Config_Monitoring_Disable_RackLoc_Location
type Environment_Config_Monitoring_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}
}

func (location *Environment_Config_Monitoring_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    return &(location.EntityData)
}

// Environment_Trace
// show traceable processes
type Environment_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Environment_Trace_Location.
    Location []Environment_Trace_Location
}

func (trace *Environment_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "environment"
    trace.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace.Buffer) + "']"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace.Location {
        trace.EntityData.Children[types.GetSegmentPath(&trace.Location[i])] = types.YChild{"Location", &trace.Location[i]}
    }
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    trace.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace.Buffer}
    return &(trace.EntityData)
}

// Environment_Trace_Location
type Environment_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Environment_Trace_Location_AllOptions.
    AllOptions []Environment_Trace_Location_AllOptions
}

func (location *Environment_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + "[location_name='" + fmt.Sprintf("%v", location.LocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["all-options"] = types.YChild{"AllOptions", nil}
    for i := range location.AllOptions {
        location.EntityData.Children[types.GetSegmentPath(&location.AllOptions[i])] = types.YChild{"AllOptions", &location.AllOptions[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location_name"] = types.YLeaf{"LocationName", location.LocationName}
    return &(location.EntityData)
}

// Environment_Trace_Location_AllOptions
type Environment_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Environment_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Environment_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Environment_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + "[option='" + fmt.Sprintf("%v", allOptions.Option) + "']"
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = make(map[string]types.YChild)
    allOptions.EntityData.Children["trace-blocks"] = types.YChild{"TraceBlocks", nil}
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children[types.GetSegmentPath(&allOptions.TraceBlocks[i])] = types.YChild{"TraceBlocks", &allOptions.TraceBlocks[i]}
    }
    allOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    allOptions.EntityData.Leafs["option"] = types.YLeaf{"Option", allOptions.Option}
    return &(allOptions.EntityData)
}

// Environment_Trace_Location_AllOptions_TraceBlocks
type Environment_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Environment_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = make(map[string]types.YChild)
    traceBlocks.EntityData.Leafs = make(map[string]types.YLeaf)
    traceBlocks.EntityData.Leafs["data"] = types.YLeaf{"Data", traceBlocks.Data}
    return &(traceBlocks.EntityData)
}

// PowerMgmt
type PowerMgmt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Power Trays and PEMs configurational data.
    Config PowerMgmt_Config
}

func (powerMgmt *PowerMgmt) GetEntityData() *types.CommonEntityData {
    powerMgmt.EntityData.YFilter = powerMgmt.YFilter
    powerMgmt.EntityData.YangName = "power-mgmt"
    powerMgmt.EntityData.BundleName = "cisco_ios_xr"
    powerMgmt.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-envmon-ui"
    powerMgmt.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt"
    powerMgmt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerMgmt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerMgmt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerMgmt.EntityData.Children = make(map[string]types.YChild)
    powerMgmt.EntityData.Children["config"] = types.YChild{"Config", &powerMgmt.Config}
    powerMgmt.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(powerMgmt.EntityData)
}

// PowerMgmt_Config
// Power Trays and PEMs configurational data
type PowerMgmt_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Action PowerMgmt_Config_Action

    
    SingleFeedMode PowerMgmt_Config_SingleFeedMode

    
    ExtendedTemp PowerMgmt_Config_ExtendedTemp

    
    RedundancyNumPms PowerMgmt_Config_RedundancyNumPms
}

func (config *PowerMgmt_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "power-mgmt"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["action"] = types.YChild{"Action", &config.Action}
    config.EntityData.Children["single-feed-mode"] = types.YChild{"SingleFeedMode", &config.SingleFeedMode}
    config.EntityData.Children["extended-temp"] = types.YChild{"ExtendedTemp", &config.ExtendedTemp}
    config.EntityData.Children["redundancy-num-pms"] = types.YChild{"RedundancyNumPms", &config.RedundancyNumPms}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(config.EntityData)
}

// PowerMgmt_Config_Action
type PowerMgmt_Config_Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Disable PowerMgmt_Config_Action_Disable
}

func (action *PowerMgmt_Config_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "config"
    action.EntityData.SegmentPath = "action"
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = make(map[string]types.YChild)
    action.EntityData.Children["disable"] = types.YChild{"Disable", &action.Disable}
    action.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(action.EntityData)
}

// PowerMgmt_Config_Action_Disable
type PowerMgmt_Config_Action_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc PowerMgmt_Config_Action_Disable_RackLoc
}

func (disable *PowerMgmt_Config_Action_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "action"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = make(map[string]types.YChild)
    disable.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &disable.RackLoc}
    disable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(disable.EntityData)
}

// PowerMgmt_Config_Action_Disable_RackLoc
type PowerMgmt_Config_Action_Disable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of PowerMgmt_Config_Action_Disable_RackLoc_Location.
    Location []PowerMgmt_Config_Action_Disable_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_Action_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    rackLoc.EntityData.Leafs["all"] = types.YLeaf{"All", rackLoc.All}
    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_Action_Disable_RackLoc_Location
type PowerMgmt_Config_Action_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}
}

func (location *PowerMgmt_Config_Action_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    return &(location.EntityData)
}

// PowerMgmt_Config_SingleFeedMode
type PowerMgmt_Config_SingleFeedMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Enable PowerMgmt_Config_SingleFeedMode_Enable
}

func (singleFeedMode *PowerMgmt_Config_SingleFeedMode) GetEntityData() *types.CommonEntityData {
    singleFeedMode.EntityData.YFilter = singleFeedMode.YFilter
    singleFeedMode.EntityData.YangName = "single-feed-mode"
    singleFeedMode.EntityData.BundleName = "cisco_ios_xr"
    singleFeedMode.EntityData.ParentYangName = "config"
    singleFeedMode.EntityData.SegmentPath = "single-feed-mode"
    singleFeedMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleFeedMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleFeedMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleFeedMode.EntityData.Children = make(map[string]types.YChild)
    singleFeedMode.EntityData.Children["enable"] = types.YChild{"Enable", &singleFeedMode.Enable}
    singleFeedMode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(singleFeedMode.EntityData)
}

// PowerMgmt_Config_SingleFeedMode_Enable
type PowerMgmt_Config_SingleFeedMode_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc PowerMgmt_Config_SingleFeedMode_Enable_RackLoc
}

func (enable *PowerMgmt_Config_SingleFeedMode_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "single-feed-mode"
    enable.EntityData.SegmentPath = "enable"
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = make(map[string]types.YChild)
    enable.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &enable.RackLoc}
    enable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enable.EntityData)
}

// PowerMgmt_Config_SingleFeedMode_Enable_RackLoc
type PowerMgmt_Config_SingleFeedMode_Enable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of
    // PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location.
    Location []PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_SingleFeedMode_Enable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "enable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    rackLoc.EntityData.Leafs["all"] = types.YLeaf{"All", rackLoc.All}
    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location
type PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}
}

func (location *PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    return &(location.EntityData)
}

// PowerMgmt_Config_ExtendedTemp
type PowerMgmt_Config_ExtendedTemp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Enable PowerMgmt_Config_ExtendedTemp_Enable
}

func (extendedTemp *PowerMgmt_Config_ExtendedTemp) GetEntityData() *types.CommonEntityData {
    extendedTemp.EntityData.YFilter = extendedTemp.YFilter
    extendedTemp.EntityData.YangName = "extended-temp"
    extendedTemp.EntityData.BundleName = "cisco_ios_xr"
    extendedTemp.EntityData.ParentYangName = "config"
    extendedTemp.EntityData.SegmentPath = "extended-temp"
    extendedTemp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedTemp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedTemp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedTemp.EntityData.Children = make(map[string]types.YChild)
    extendedTemp.EntityData.Children["enable"] = types.YChild{"Enable", &extendedTemp.Enable}
    extendedTemp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extendedTemp.EntityData)
}

// PowerMgmt_Config_ExtendedTemp_Enable
type PowerMgmt_Config_ExtendedTemp_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc PowerMgmt_Config_ExtendedTemp_Enable_RackLoc
}

func (enable *PowerMgmt_Config_ExtendedTemp_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "extended-temp"
    enable.EntityData.SegmentPath = "enable"
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = make(map[string]types.YChild)
    enable.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &enable.RackLoc}
    enable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enable.EntityData)
}

// PowerMgmt_Config_ExtendedTemp_Enable_RackLoc
type PowerMgmt_Config_ExtendedTemp_Enable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location.
    Location []PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_ExtendedTemp_Enable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "enable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    rackLoc.EntityData.Leafs["all"] = types.YLeaf{"All", rackLoc.All}
    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location
type PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}
}

func (location *PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    return &(location.EntityData)
}

// PowerMgmt_Config_RedundancyNumPms
type PowerMgmt_Config_RedundancyNumPms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All PowerMgmt_Config_RedundancyNumPms_All

    
    RackLoc PowerMgmt_Config_RedundancyNumPms_RackLoc
}

func (redundancyNumPms *PowerMgmt_Config_RedundancyNumPms) GetEntityData() *types.CommonEntityData {
    redundancyNumPms.EntityData.YFilter = redundancyNumPms.YFilter
    redundancyNumPms.EntityData.YangName = "redundancy-num-pms"
    redundancyNumPms.EntityData.BundleName = "cisco_ios_xr"
    redundancyNumPms.EntityData.ParentYangName = "config"
    redundancyNumPms.EntityData.SegmentPath = "redundancy-num-pms"
    redundancyNumPms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancyNumPms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancyNumPms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancyNumPms.EntityData.Children = make(map[string]types.YChild)
    redundancyNumPms.EntityData.Children["all"] = types.YChild{"All", &redundancyNumPms.All}
    redundancyNumPms.EntityData.Children["rack_loc"] = types.YChild{"RackLoc", &redundancyNumPms.RackLoc}
    redundancyNumPms.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(redundancyNumPms.EntityData)
}

// PowerMgmt_Config_RedundancyNumPms_All
type PowerMgmt_Config_RedundancyNumPms_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..12.
    NumPm interface{}
}

func (all *PowerMgmt_Config_RedundancyNumPms_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "redundancy-num-pms"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    all.EntityData.Leafs["num_pm"] = types.YLeaf{"NumPm", all.NumPm}
    return &(all.EntityData)
}

// PowerMgmt_Config_RedundancyNumPms_RackLoc
type PowerMgmt_Config_RedundancyNumPms_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of PowerMgmt_Config_RedundancyNumPms_RackLoc_Location.
    Location []PowerMgmt_Config_RedundancyNumPms_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_RedundancyNumPms_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "redundancy-num-pms"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = make(map[string]types.YChild)
    rackLoc.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children[types.GetSegmentPath(&rackLoc.Location[i])] = types.YChild{"Location", &rackLoc.Location[i]}
    }
    rackLoc.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_RedundancyNumPms_RackLoc_Location
type PowerMgmt_Config_RedundancyNumPms_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is RackId.
    Rackid interface{}

    // The type is interface{} with range: 0..12.
    NumPm interface{}
}

func (location *PowerMgmt_Config_RedundancyNumPms_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + "[rackId='" + fmt.Sprintf("%v", location.Rackid) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rackId"] = types.YLeaf{"Rackid", location.Rackid}
    location.EntityData.Leafs["num_pm"] = types.YLeaf{"NumPm", location.NumPm}
    return &(location.EntityData)
}

