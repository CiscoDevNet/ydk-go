// This module contains definitions
// for the Calvados model objects.
// 
// This module holds chassis, cards Enviroment data.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
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
    Trace []*Environment_Trace
}

func (environment *Environment) GetEntityData() *types.CommonEntityData {
    environment.EntityData.YFilter = environment.YFilter
    environment.EntityData.YangName = "environment"
    environment.EntityData.BundleName = "cisco_ios_xr"
    environment.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-envmon-ui"
    environment.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment"
    environment.EntityData.AbsolutePath = environment.EntityData.SegmentPath
    environment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environment.EntityData.Children = types.NewOrderedMap()
    environment.EntityData.Children.Append("oper", types.YChild{"Oper", &environment.Oper})
    environment.EntityData.Children.Append("all", types.YChild{"All", &environment.All})
    environment.EntityData.Children.Append("config", types.YChild{"Config", &environment.Config})
    environment.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range environment.Trace {
        environment.EntityData.Children.Append(types.GetSegmentPath(environment.Trace[i]), types.YChild{"Trace", environment.Trace[i]})
    }
    environment.EntityData.Leafs = types.NewOrderedMap()

    environment.EntityData.YListKeys = []string {}

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
    oper.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/" + oper.EntityData.SegmentPath
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = types.NewOrderedMap()
    oper.EntityData.Children.Append("temperatures", types.YChild{"Temperatures", &oper.Temperatures})
    oper.EntityData.Children.Append("voltages", types.YChild{"Voltages", &oper.Voltages})
    oper.EntityData.Children.Append("current", types.YChild{"Current", &oper.Current})
    oper.EntityData.Children.Append("fan", types.YChild{"Fan", &oper.Fan})
    oper.EntityData.Children.Append("power", types.YChild{"Power", &oper.Power})
    oper.EntityData.Children.Append("altitude", types.YChild{"Altitude", &oper.Altitude})
    oper.EntityData.Leafs = types.NewOrderedMap()

    oper.EntityData.YListKeys = []string {}

    return &(oper.EntityData)
}

// Environment_Oper_Temperatures
type Environment_Oper_Temperatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Temperatures_Location.
    Location []*Environment_Oper_Temperatures_Location
}

func (temperatures *Environment_Oper_Temperatures) GetEntityData() *types.CommonEntityData {
    temperatures.EntityData.YFilter = temperatures.YFilter
    temperatures.EntityData.YangName = "temperatures"
    temperatures.EntityData.BundleName = "cisco_ios_xr"
    temperatures.EntityData.ParentYangName = "oper"
    temperatures.EntityData.SegmentPath = "temperatures"
    temperatures.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/" + temperatures.EntityData.SegmentPath
    temperatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    temperatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    temperatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    temperatures.EntityData.Children = types.NewOrderedMap()
    temperatures.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range temperatures.Location {
        temperatures.EntityData.Children.Append(types.GetSegmentPath(temperatures.Location[i]), types.YChild{"Location", temperatures.Location[i]})
    }
    temperatures.EntityData.Leafs = types.NewOrderedMap()

    temperatures.EntityData.YListKeys = []string {}

    return &(temperatures.EntityData)
}

// Environment_Oper_Temperatures_Location
type Environment_Oper_Temperatures_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is interface{} with range: 0..4294967295.
    LocHeader interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is slice of
    // Environment_Oper_Temperatures_Location_SensorAttributes.
    SensorAttributes []*Environment_Oper_Temperatures_Location_SensorAttributes
}

func (location *Environment_Oper_Temperatures_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "temperatures"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/temperatures/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("sensor_attributes", types.YChild{"SensorAttributes", nil})
    for i := range location.SensorAttributes {
        location.EntityData.Children.Append(types.GetSegmentPath(location.SensorAttributes[i]), types.YChild{"SensorAttributes", location.SensorAttributes[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})
    location.EntityData.Leafs.Append("loc_header", types.YLeaf{"LocHeader", location.LocHeader})
    location.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", location.PrintHeader})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Environment_Oper_Temperatures_Location_SensorAttributes
type Environment_Oper_Temperatures_Location_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + types.AddKeyToken(sensorAttributes.Sensor, "sensor")
    sensorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/temperatures/location/" + sensorAttributes.EntityData.SegmentPath
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", sensorAttributes.Sensor})
    sensorAttributes.EntityData.Leafs.Append("sensor_id", types.YLeaf{"SensorId", sensorAttributes.SensorId})
    sensorAttributes.EntityData.Leafs.Append("alarm", types.YLeaf{"Alarm", sensorAttributes.Alarm})
    sensorAttributes.EntityData.Leafs.Append("temperature_value", types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue})
    sensorAttributes.EntityData.Leafs.Append("value", types.YLeaf{"Value", sensorAttributes.Value})
    sensorAttributes.EntityData.Leafs.Append("critical_lo", types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo})
    sensorAttributes.EntityData.Leafs.Append("major_lo", types.YLeaf{"MajorLo", sensorAttributes.MajorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_lo", types.YLeaf{"MinorLo", sensorAttributes.MinorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_hi", types.YLeaf{"MinorHi", sensorAttributes.MinorHi})
    sensorAttributes.EntityData.Leafs.Append("major_hi", types.YLeaf{"MajorHi", sensorAttributes.MajorHi})
    sensorAttributes.EntityData.Leafs.Append("critical_hi", types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi})

    sensorAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(sensorAttributes.EntityData)
}

// Environment_Oper_Voltages
type Environment_Oper_Voltages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Voltages_Location.
    Location []*Environment_Oper_Voltages_Location
}

func (voltages *Environment_Oper_Voltages) GetEntityData() *types.CommonEntityData {
    voltages.EntityData.YFilter = voltages.YFilter
    voltages.EntityData.YangName = "voltages"
    voltages.EntityData.BundleName = "cisco_ios_xr"
    voltages.EntityData.ParentYangName = "oper"
    voltages.EntityData.SegmentPath = "voltages"
    voltages.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/" + voltages.EntityData.SegmentPath
    voltages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voltages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voltages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voltages.EntityData.Children = types.NewOrderedMap()
    voltages.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range voltages.Location {
        voltages.EntityData.Children.Append(types.GetSegmentPath(voltages.Location[i]), types.YChild{"Location", voltages.Location[i]})
    }
    voltages.EntityData.Leafs = types.NewOrderedMap()

    voltages.EntityData.YListKeys = []string {}

    return &(voltages.EntityData)
}

// Environment_Oper_Voltages_Location
type Environment_Oper_Voltages_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is interface{} with range: 0..4294967295.
    LocHeader interface{}

    // The type is slice of Environment_Oper_Voltages_Location_SensorAttributes.
    SensorAttributes []*Environment_Oper_Voltages_Location_SensorAttributes
}

func (location *Environment_Oper_Voltages_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "voltages"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/voltages/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("sensor_attributes", types.YChild{"SensorAttributes", nil})
    for i := range location.SensorAttributes {
        location.EntityData.Children.Append(types.GetSegmentPath(location.SensorAttributes[i]), types.YChild{"SensorAttributes", location.SensorAttributes[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})
    location.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", location.PrintHeader})
    location.EntityData.Leafs.Append("loc_header", types.YLeaf{"LocHeader", location.LocHeader})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Environment_Oper_Voltages_Location_SensorAttributes
type Environment_Oper_Voltages_Location_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + types.AddKeyToken(sensorAttributes.Sensor, "sensor")
    sensorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/voltages/location/" + sensorAttributes.EntityData.SegmentPath
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", sensorAttributes.Sensor})
    sensorAttributes.EntityData.Leafs.Append("sensor_id", types.YLeaf{"SensorId", sensorAttributes.SensorId})
    sensorAttributes.EntityData.Leafs.Append("alarm", types.YLeaf{"Alarm", sensorAttributes.Alarm})
    sensorAttributes.EntityData.Leafs.Append("value", types.YLeaf{"Value", sensorAttributes.Value})
    sensorAttributes.EntityData.Leafs.Append("temperature_value", types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue})
    sensorAttributes.EntityData.Leafs.Append("critical_lo", types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo})
    sensorAttributes.EntityData.Leafs.Append("major_lo", types.YLeaf{"MajorLo", sensorAttributes.MajorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_lo", types.YLeaf{"MinorLo", sensorAttributes.MinorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_hi", types.YLeaf{"MinorHi", sensorAttributes.MinorHi})
    sensorAttributes.EntityData.Leafs.Append("major_hi", types.YLeaf{"MajorHi", sensorAttributes.MajorHi})
    sensorAttributes.EntityData.Leafs.Append("critical_hi", types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi})

    sensorAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(sensorAttributes.EntityData)
}

// Environment_Oper_Current
type Environment_Oper_Current struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Current_Location.
    Location []*Environment_Oper_Current_Location
}

func (current *Environment_Oper_Current) GetEntityData() *types.CommonEntityData {
    current.EntityData.YFilter = current.YFilter
    current.EntityData.YangName = "current"
    current.EntityData.BundleName = "cisco_ios_xr"
    current.EntityData.ParentYangName = "oper"
    current.EntityData.SegmentPath = "current"
    current.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/" + current.EntityData.SegmentPath
    current.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    current.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    current.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    current.EntityData.Children = types.NewOrderedMap()
    current.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range current.Location {
        current.EntityData.Children.Append(types.GetSegmentPath(current.Location[i]), types.YChild{"Location", current.Location[i]})
    }
    current.EntityData.Leafs = types.NewOrderedMap()

    current.EntityData.YListKeys = []string {}

    return &(current.EntityData)
}

// Environment_Oper_Current_Location
type Environment_Oper_Current_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is interface{} with range: 0..4294967295.
    LocHeader interface{}

    // The type is slice of Environment_Oper_Current_Location_SensorAttributes.
    SensorAttributes []*Environment_Oper_Current_Location_SensorAttributes
}

func (location *Environment_Oper_Current_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "current"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/current/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("sensor_attributes", types.YChild{"SensorAttributes", nil})
    for i := range location.SensorAttributes {
        location.EntityData.Children.Append(types.GetSegmentPath(location.SensorAttributes[i]), types.YChild{"SensorAttributes", location.SensorAttributes[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})
    location.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", location.PrintHeader})
    location.EntityData.Leafs.Append("loc_header", types.YLeaf{"LocHeader", location.LocHeader})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Environment_Oper_Current_Location_SensorAttributes
type Environment_Oper_Current_Location_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + types.AddKeyToken(sensorAttributes.Sensor, "sensor")
    sensorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/current/location/" + sensorAttributes.EntityData.SegmentPath
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", sensorAttributes.Sensor})
    sensorAttributes.EntityData.Leafs.Append("sensor_id", types.YLeaf{"SensorId", sensorAttributes.SensorId})
    sensorAttributes.EntityData.Leafs.Append("alarm", types.YLeaf{"Alarm", sensorAttributes.Alarm})
    sensorAttributes.EntityData.Leafs.Append("value", types.YLeaf{"Value", sensorAttributes.Value})
    sensorAttributes.EntityData.Leafs.Append("temperature_value", types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue})

    sensorAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(sensorAttributes.EntityData)
}

// Environment_Oper_Fan
type Environment_Oper_Fan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Fan_Location.
    Location []*Environment_Oper_Fan_Location
}

func (fan *Environment_Oper_Fan) GetEntityData() *types.CommonEntityData {
    fan.EntityData.YFilter = fan.YFilter
    fan.EntityData.YangName = "fan"
    fan.EntityData.BundleName = "cisco_ios_xr"
    fan.EntityData.ParentYangName = "oper"
    fan.EntityData.SegmentPath = "fan"
    fan.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/" + fan.EntityData.SegmentPath
    fan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fan.EntityData.Children = types.NewOrderedMap()
    fan.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range fan.Location {
        fan.EntityData.Children.Append(types.GetSegmentPath(fan.Location[i]), types.YChild{"Location", fan.Location[i]})
    }
    fan.EntityData.Leafs = types.NewOrderedMap()

    fan.EntityData.YListKeys = []string {}

    return &(fan.EntityData)
}

// Environment_Oper_Fan_Location
type Environment_Oper_Fan_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is interface{} with range: 0..4294967295.
    LocHeader interface{}

    // The type is slice of Environment_Oper_Fan_Location_FanAttributes.
    FanAttributes []*Environment_Oper_Fan_Location_FanAttributes
}

func (location *Environment_Oper_Fan_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "fan"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/fan/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("fan_attributes", types.YChild{"FanAttributes", nil})
    for i := range location.FanAttributes {
        location.EntityData.Children.Append(types.GetSegmentPath(location.FanAttributes[i]), types.YChild{"FanAttributes", location.FanAttributes[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})
    location.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", location.PrintHeader})
    location.EntityData.Leafs.Append("loc_header", types.YLeaf{"LocHeader", location.LocHeader})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Environment_Oper_Fan_Location_FanAttributes
type Environment_Oper_Fan_Location_FanAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LogicalSlot interface{}

    // The type is string.
    PrintFanHeader interface{}

    // The type is string.
    Location interface{}

    // The type is string.
    FruPID interface{}

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
    fanAttributes.EntityData.SegmentPath = "fan_attributes" + types.AddKeyToken(fanAttributes.LogicalSlot, "logical_slot")
    fanAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/fan/location/" + fanAttributes.EntityData.SegmentPath
    fanAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanAttributes.EntityData.Children = types.NewOrderedMap()
    fanAttributes.EntityData.Leafs = types.NewOrderedMap()
    fanAttributes.EntityData.Leafs.Append("logical_slot", types.YLeaf{"LogicalSlot", fanAttributes.LogicalSlot})
    fanAttributes.EntityData.Leafs.Append("print_fan_header", types.YLeaf{"PrintFanHeader", fanAttributes.PrintFanHeader})
    fanAttributes.EntityData.Leafs.Append("location", types.YLeaf{"Location", fanAttributes.Location})
    fanAttributes.EntityData.Leafs.Append("fru_PID", types.YLeaf{"FruPID", fanAttributes.FruPID})
    fanAttributes.EntityData.Leafs.Append("fans_speed", types.YLeaf{"FansSpeed", fanAttributes.FansSpeed})
    fanAttributes.EntityData.Leafs.Append("fan_header", types.YLeaf{"FanHeader", fanAttributes.FanHeader})
    fanAttributes.EntityData.Leafs.Append("speed_space", types.YLeaf{"SpeedSpace", fanAttributes.SpeedSpace})

    fanAttributes.EntityData.YListKeys = []string {"LogicalSlot"}

    return &(fanAttributes.EntityData)
}

// Environment_Oper_Power
type Environment_Oper_Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Power_Location.
    Location []*Environment_Oper_Power_Location
}

func (power *Environment_Oper_Power) GetEntityData() *types.CommonEntityData {
    power.EntityData.YFilter = power.YFilter
    power.EntityData.YangName = "power"
    power.EntityData.BundleName = "cisco_ios_xr"
    power.EntityData.ParentYangName = "oper"
    power.EntityData.SegmentPath = "power"
    power.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/" + power.EntityData.SegmentPath
    power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    power.EntityData.Children = types.NewOrderedMap()
    power.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range power.Location {
        power.EntityData.Children.Append(types.GetSegmentPath(power.Location[i]), types.YChild{"Location", power.Location[i]})
    }
    power.EntityData.Leafs = types.NewOrderedMap()

    power.EntityData.YListKeys = []string {}

    return &(power.EntityData)
}

// Environment_Oper_Power_Location
type Environment_Oper_Power_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_Oper_Power_Location_PemAttributes.
    PemAttributes []*Environment_Oper_Power_Location_PemAttributes
}

func (location *Environment_Oper_Power_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "power"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/power/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("pem_attributes", types.YChild{"PemAttributes", nil})
    for i := range location.PemAttributes {
        location.EntityData.Children.Append(types.GetSegmentPath(location.PemAttributes[i]), types.YChild{"PemAttributes", location.PemAttributes[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Environment_Oper_Power_Location_PemAttributes
type Environment_Oper_Power_Location_PemAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    pemAttributes.EntityData.SegmentPath = "pem_attributes" + types.AddKeyToken(pemAttributes.Pem, "pem")
    pemAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/power/location/" + pemAttributes.EntityData.SegmentPath
    pemAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pemAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pemAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pemAttributes.EntityData.Children = types.NewOrderedMap()
    pemAttributes.EntityData.Leafs = types.NewOrderedMap()
    pemAttributes.EntityData.Leafs.Append("pem", types.YLeaf{"Pem", pemAttributes.Pem})
    pemAttributes.EntityData.Leafs.Append("pem_id", types.YLeaf{"PemId", pemAttributes.PemId})
    pemAttributes.EntityData.Leafs.Append("card_type", types.YLeaf{"CardType", pemAttributes.CardType})
    pemAttributes.EntityData.Leafs.Append("ps_type", types.YLeaf{"PsType", pemAttributes.PsType})
    pemAttributes.EntityData.Leafs.Append("shelf_num", types.YLeaf{"ShelfNum", pemAttributes.ShelfNum})
    pemAttributes.EntityData.Leafs.Append("supply_type", types.YLeaf{"SupplyType", pemAttributes.SupplyType})
    pemAttributes.EntityData.Leafs.Append("input_voltage", types.YLeaf{"InputVoltage", pemAttributes.InputVoltage})
    pemAttributes.EntityData.Leafs.Append("input_current", types.YLeaf{"InputCurrent", pemAttributes.InputCurrent})
    pemAttributes.EntityData.Leafs.Append("output_voltage", types.YLeaf{"OutputVoltage", pemAttributes.OutputVoltage})
    pemAttributes.EntityData.Leafs.Append("output_current", types.YLeaf{"OutputCurrent", pemAttributes.OutputCurrent})
    pemAttributes.EntityData.Leafs.Append("status", types.YLeaf{"Status", pemAttributes.Status})
    pemAttributes.EntityData.Leafs.Append("input_power_to_ps", types.YLeaf{"InputPowerToPs", pemAttributes.InputPowerToPs})
    pemAttributes.EntityData.Leafs.Append("input_current_to_ps", types.YLeaf{"InputCurrentToPs", pemAttributes.InputCurrentToPs})
    pemAttributes.EntityData.Leafs.Append("output_power_from_ps", types.YLeaf{"OutputPowerFromPs", pemAttributes.OutputPowerFromPs})
    pemAttributes.EntityData.Leafs.Append("output_current_from_ps", types.YLeaf{"OutputCurrentFromPs", pemAttributes.OutputCurrentFromPs})
    pemAttributes.EntityData.Leafs.Append("power_allocated", types.YLeaf{"PowerAllocated", pemAttributes.PowerAllocated})
    pemAttributes.EntityData.Leafs.Append("power_consumed", types.YLeaf{"PowerConsumed", pemAttributes.PowerConsumed})
    pemAttributes.EntityData.Leafs.Append("power_status", types.YLeaf{"PowerStatus", pemAttributes.PowerStatus})
    pemAttributes.EntityData.Leafs.Append("confgd_power_redundancy_mode", types.YLeaf{"ConfgdPowerRedundancyMode", pemAttributes.ConfgdPowerRedundancyMode})
    pemAttributes.EntityData.Leafs.Append("usable_power_capacity", types.YLeaf{"UsablePowerCapacity", pemAttributes.UsablePowerCapacity})
    pemAttributes.EntityData.Leafs.Append("protection_power_capacity", types.YLeaf{"ProtectionPowerCapacity", pemAttributes.ProtectionPowerCapacity})
    pemAttributes.EntityData.Leafs.Append("power_resrv_and_alloc", types.YLeaf{"PowerResrvAndAlloc", pemAttributes.PowerResrvAndAlloc})
    pemAttributes.EntityData.Leafs.Append("system_power_used", types.YLeaf{"SystemPowerUsed", pemAttributes.SystemPowerUsed})
    pemAttributes.EntityData.Leafs.Append("system_power_input", types.YLeaf{"SystemPowerInput", pemAttributes.SystemPowerInput})
    pemAttributes.EntityData.Leafs.Append("power_level", types.YLeaf{"PowerLevel", pemAttributes.PowerLevel})
    pemAttributes.EntityData.Leafs.Append("output_header", types.YLeaf{"OutputHeader", pemAttributes.OutputHeader})
    pemAttributes.EntityData.Leafs.Append("output_footer", types.YLeaf{"OutputFooter", pemAttributes.OutputFooter})
    pemAttributes.EntityData.Leafs.Append("ps_sum_footer", types.YLeaf{"PsSumFooter", pemAttributes.PsSumFooter})

    pemAttributes.EntityData.YListKeys = []string {"Pem"}

    return &(pemAttributes.EntityData)
}

// Environment_Oper_Altitude
type Environment_Oper_Altitude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Oper_Altitude_Location.
    Location []*Environment_Oper_Altitude_Location
}

func (altitude *Environment_Oper_Altitude) GetEntityData() *types.CommonEntityData {
    altitude.EntityData.YFilter = altitude.YFilter
    altitude.EntityData.YangName = "altitude"
    altitude.EntityData.BundleName = "cisco_ios_xr"
    altitude.EntityData.ParentYangName = "oper"
    altitude.EntityData.SegmentPath = "altitude"
    altitude.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/" + altitude.EntityData.SegmentPath
    altitude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altitude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altitude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altitude.EntityData.Children = types.NewOrderedMap()
    altitude.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range altitude.Location {
        altitude.EntityData.Children.Append(types.GetSegmentPath(altitude.Location[i]), types.YChild{"Location", altitude.Location[i]})
    }
    altitude.EntityData.Leafs = types.NewOrderedMap()

    altitude.EntityData.YListKeys = []string {}

    return &(altitude.EntityData)
}

// Environment_Oper_Altitude_Location
type Environment_Oper_Altitude_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_Oper_Altitude_Location_AltAttributes.
    AltAttributes []*Environment_Oper_Altitude_Location_AltAttributes
}

func (location *Environment_Oper_Altitude_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "altitude"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/altitude/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("alt_attributes", types.YChild{"AltAttributes", nil})
    for i := range location.AltAttributes {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AltAttributes[i]), types.YChild{"AltAttributes", location.AltAttributes[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Environment_Oper_Altitude_Location_AltAttributes
type Environment_Oper_Altitude_Location_AltAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is interface{} with range: 0..4294967295.
    Rack interface{}

    // The type is string.
    SensorValue interface{}

    // The type is string.
    Source interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}
}

func (altAttributes *Environment_Oper_Altitude_Location_AltAttributes) GetEntityData() *types.CommonEntityData {
    altAttributes.EntityData.YFilter = altAttributes.YFilter
    altAttributes.EntityData.YangName = "alt_attributes"
    altAttributes.EntityData.BundleName = "cisco_ios_xr"
    altAttributes.EntityData.ParentYangName = "location"
    altAttributes.EntityData.SegmentPath = "alt_attributes" + types.AddKeyToken(altAttributes.Sensor, "sensor")
    altAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/oper/altitude/location/" + altAttributes.EntityData.SegmentPath
    altAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altAttributes.EntityData.Children = types.NewOrderedMap()
    altAttributes.EntityData.Leafs = types.NewOrderedMap()
    altAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", altAttributes.Sensor})
    altAttributes.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", altAttributes.Rack})
    altAttributes.EntityData.Leafs.Append("sensor_value", types.YLeaf{"SensorValue", altAttributes.SensorValue})
    altAttributes.EntityData.Leafs.Append("source", types.YLeaf{"Source", altAttributes.Source})
    altAttributes.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", altAttributes.PrintHeader})

    altAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(altAttributes.EntityData)
}

// Environment_All
type Environment_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_All_Location.
    Location []*Environment_All_Location
}

func (all *Environment_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "environment"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Environment_All_Location
type Environment_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of Environment_All_Location_Temperatures.
    Temperatures []*Environment_All_Location_Temperatures

    // The type is slice of Environment_All_Location_Voltages.
    Voltages []*Environment_All_Location_Voltages

    // The type is slice of Environment_All_Location_Current.
    Current []*Environment_All_Location_Current

    // The type is slice of Environment_All_Location_Fan.
    Fan []*Environment_All_Location_Fan

    // The type is slice of Environment_All_Location_Power.
    Power []*Environment_All_Location_Power

    // The type is slice of Environment_All_Location_Altitude.
    Altitude []*Environment_All_Location_Altitude
}

func (location *Environment_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("temperatures", types.YChild{"Temperatures", nil})
    for i := range location.Temperatures {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Temperatures[i]), types.YChild{"Temperatures", location.Temperatures[i]})
    }
    location.EntityData.Children.Append("voltages", types.YChild{"Voltages", nil})
    for i := range location.Voltages {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Voltages[i]), types.YChild{"Voltages", location.Voltages[i]})
    }
    location.EntityData.Children.Append("current", types.YChild{"Current", nil})
    for i := range location.Current {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Current[i]), types.YChild{"Current", location.Current[i]})
    }
    location.EntityData.Children.Append("fan", types.YChild{"Fan", nil})
    for i := range location.Fan {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Fan[i]), types.YChild{"Fan", location.Fan[i]})
    }
    location.EntityData.Children.Append("power", types.YChild{"Power", nil})
    for i := range location.Power {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Power[i]), types.YChild{"Power", location.Power[i]})
    }
    location.EntityData.Children.Append("altitude", types.YChild{"Altitude", nil})
    for i := range location.Altitude {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Altitude[i]), types.YChild{"Altitude", location.Altitude[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Environment_All_Location_Temperatures
type Environment_All_Location_Temperatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is slice of
    // Environment_All_Location_Temperatures_SensorAttributes.
    SensorAttributes []*Environment_All_Location_Temperatures_SensorAttributes
}

func (temperatures *Environment_All_Location_Temperatures) GetEntityData() *types.CommonEntityData {
    temperatures.EntityData.YFilter = temperatures.YFilter
    temperatures.EntityData.YangName = "temperatures"
    temperatures.EntityData.BundleName = "cisco_ios_xr"
    temperatures.EntityData.ParentYangName = "location"
    temperatures.EntityData.SegmentPath = "temperatures" + types.AddKeyToken(temperatures.LocIden, "loc_iden")
    temperatures.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/" + temperatures.EntityData.SegmentPath
    temperatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    temperatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    temperatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    temperatures.EntityData.Children = types.NewOrderedMap()
    temperatures.EntityData.Children.Append("sensor_attributes", types.YChild{"SensorAttributes", nil})
    for i := range temperatures.SensorAttributes {
        temperatures.EntityData.Children.Append(types.GetSegmentPath(temperatures.SensorAttributes[i]), types.YChild{"SensorAttributes", temperatures.SensorAttributes[i]})
    }
    temperatures.EntityData.Leafs = types.NewOrderedMap()
    temperatures.EntityData.Leafs.Append("loc_iden", types.YLeaf{"LocIden", temperatures.LocIden})
    temperatures.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", temperatures.PrintHeader})

    temperatures.EntityData.YListKeys = []string {"LocIden"}

    return &(temperatures.EntityData)
}

// Environment_All_Location_Temperatures_SensorAttributes
type Environment_All_Location_Temperatures_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + types.AddKeyToken(sensorAttributes.Sensor, "sensor")
    sensorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/temperatures/" + sensorAttributes.EntityData.SegmentPath
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", sensorAttributes.Sensor})
    sensorAttributes.EntityData.Leafs.Append("sensor_id", types.YLeaf{"SensorId", sensorAttributes.SensorId})
    sensorAttributes.EntityData.Leafs.Append("alarm", types.YLeaf{"Alarm", sensorAttributes.Alarm})
    sensorAttributes.EntityData.Leafs.Append("temperature_value", types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue})
    sensorAttributes.EntityData.Leafs.Append("value", types.YLeaf{"Value", sensorAttributes.Value})
    sensorAttributes.EntityData.Leafs.Append("critical_lo", types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo})
    sensorAttributes.EntityData.Leafs.Append("major_lo", types.YLeaf{"MajorLo", sensorAttributes.MajorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_lo", types.YLeaf{"MinorLo", sensorAttributes.MinorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_hi", types.YLeaf{"MinorHi", sensorAttributes.MinorHi})
    sensorAttributes.EntityData.Leafs.Append("major_hi", types.YLeaf{"MajorHi", sensorAttributes.MajorHi})
    sensorAttributes.EntityData.Leafs.Append("critical_hi", types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi})

    sensorAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(sensorAttributes.EntityData)
}

// Environment_All_Location_Voltages
type Environment_All_Location_Voltages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is slice of Environment_All_Location_Voltages_SensorAttributes.
    SensorAttributes []*Environment_All_Location_Voltages_SensorAttributes
}

func (voltages *Environment_All_Location_Voltages) GetEntityData() *types.CommonEntityData {
    voltages.EntityData.YFilter = voltages.YFilter
    voltages.EntityData.YangName = "voltages"
    voltages.EntityData.BundleName = "cisco_ios_xr"
    voltages.EntityData.ParentYangName = "location"
    voltages.EntityData.SegmentPath = "voltages" + types.AddKeyToken(voltages.LocIden, "loc_iden")
    voltages.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/" + voltages.EntityData.SegmentPath
    voltages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voltages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voltages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voltages.EntityData.Children = types.NewOrderedMap()
    voltages.EntityData.Children.Append("sensor_attributes", types.YChild{"SensorAttributes", nil})
    for i := range voltages.SensorAttributes {
        voltages.EntityData.Children.Append(types.GetSegmentPath(voltages.SensorAttributes[i]), types.YChild{"SensorAttributes", voltages.SensorAttributes[i]})
    }
    voltages.EntityData.Leafs = types.NewOrderedMap()
    voltages.EntityData.Leafs.Append("loc_iden", types.YLeaf{"LocIden", voltages.LocIden})
    voltages.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", voltages.PrintHeader})

    voltages.EntityData.YListKeys = []string {"LocIden"}

    return &(voltages.EntityData)
}

// Environment_All_Location_Voltages_SensorAttributes
type Environment_All_Location_Voltages_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + types.AddKeyToken(sensorAttributes.Sensor, "sensor")
    sensorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/voltages/" + sensorAttributes.EntityData.SegmentPath
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", sensorAttributes.Sensor})
    sensorAttributes.EntityData.Leafs.Append("sensor_id", types.YLeaf{"SensorId", sensorAttributes.SensorId})
    sensorAttributes.EntityData.Leafs.Append("alarm", types.YLeaf{"Alarm", sensorAttributes.Alarm})
    sensorAttributes.EntityData.Leafs.Append("value", types.YLeaf{"Value", sensorAttributes.Value})
    sensorAttributes.EntityData.Leafs.Append("temperature_value", types.YLeaf{"TemperatureValue", sensorAttributes.TemperatureValue})
    sensorAttributes.EntityData.Leafs.Append("critical_lo", types.YLeaf{"CriticalLo", sensorAttributes.CriticalLo})
    sensorAttributes.EntityData.Leafs.Append("major_lo", types.YLeaf{"MajorLo", sensorAttributes.MajorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_lo", types.YLeaf{"MinorLo", sensorAttributes.MinorLo})
    sensorAttributes.EntityData.Leafs.Append("minor_hi", types.YLeaf{"MinorHi", sensorAttributes.MinorHi})
    sensorAttributes.EntityData.Leafs.Append("major_hi", types.YLeaf{"MajorHi", sensorAttributes.MajorHi})
    sensorAttributes.EntityData.Leafs.Append("critical_hi", types.YLeaf{"CriticalHi", sensorAttributes.CriticalHi})

    sensorAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(sensorAttributes.EntityData)
}

// Environment_All_Location_Current
type Environment_All_Location_Current struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is slice of Environment_All_Location_Current_SensorAttributes.
    SensorAttributes []*Environment_All_Location_Current_SensorAttributes
}

func (current *Environment_All_Location_Current) GetEntityData() *types.CommonEntityData {
    current.EntityData.YFilter = current.YFilter
    current.EntityData.YangName = "current"
    current.EntityData.BundleName = "cisco_ios_xr"
    current.EntityData.ParentYangName = "location"
    current.EntityData.SegmentPath = "current" + types.AddKeyToken(current.LocIden, "loc_iden")
    current.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/" + current.EntityData.SegmentPath
    current.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    current.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    current.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    current.EntityData.Children = types.NewOrderedMap()
    current.EntityData.Children.Append("sensor_attributes", types.YChild{"SensorAttributes", nil})
    for i := range current.SensorAttributes {
        current.EntityData.Children.Append(types.GetSegmentPath(current.SensorAttributes[i]), types.YChild{"SensorAttributes", current.SensorAttributes[i]})
    }
    current.EntityData.Leafs = types.NewOrderedMap()
    current.EntityData.Leafs.Append("loc_iden", types.YLeaf{"LocIden", current.LocIden})
    current.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", current.PrintHeader})

    current.EntityData.YListKeys = []string {"LocIden"}

    return &(current.EntityData)
}

// Environment_All_Location_Current_SensorAttributes
type Environment_All_Location_Current_SensorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    sensorAttributes.EntityData.SegmentPath = "sensor_attributes" + types.AddKeyToken(sensorAttributes.Sensor, "sensor")
    sensorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/current/" + sensorAttributes.EntityData.SegmentPath
    sensorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorAttributes.EntityData.Children = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs = types.NewOrderedMap()
    sensorAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", sensorAttributes.Sensor})
    sensorAttributes.EntityData.Leafs.Append("sensor_id", types.YLeaf{"SensorId", sensorAttributes.SensorId})
    sensorAttributes.EntityData.Leafs.Append("value", types.YLeaf{"Value", sensorAttributes.Value})

    sensorAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(sensorAttributes.EntityData)
}

// Environment_All_Location_Fan
type Environment_All_Location_Fan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is slice of Environment_All_Location_Fan_FanAttributes.
    FanAttributes []*Environment_All_Location_Fan_FanAttributes
}

func (fan *Environment_All_Location_Fan) GetEntityData() *types.CommonEntityData {
    fan.EntityData.YFilter = fan.YFilter
    fan.EntityData.YangName = "fan"
    fan.EntityData.BundleName = "cisco_ios_xr"
    fan.EntityData.ParentYangName = "location"
    fan.EntityData.SegmentPath = "fan" + types.AddKeyToken(fan.LocIden, "loc_iden")
    fan.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/" + fan.EntityData.SegmentPath
    fan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fan.EntityData.Children = types.NewOrderedMap()
    fan.EntityData.Children.Append("fan_attributes", types.YChild{"FanAttributes", nil})
    for i := range fan.FanAttributes {
        fan.EntityData.Children.Append(types.GetSegmentPath(fan.FanAttributes[i]), types.YChild{"FanAttributes", fan.FanAttributes[i]})
    }
    fan.EntityData.Leafs = types.NewOrderedMap()
    fan.EntityData.Leafs.Append("loc_iden", types.YLeaf{"LocIden", fan.LocIden})

    fan.EntityData.YListKeys = []string {"LocIden"}

    return &(fan.EntityData)
}

// Environment_All_Location_Fan_FanAttributes
type Environment_All_Location_Fan_FanAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LogicalSlot interface{}

    // The type is string.
    PrintFanHeader interface{}

    // The type is string.
    Location interface{}

    // The type is string.
    FruPID interface{}

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
    fanAttributes.EntityData.SegmentPath = "fan_attributes" + types.AddKeyToken(fanAttributes.LogicalSlot, "logical_slot")
    fanAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/fan/" + fanAttributes.EntityData.SegmentPath
    fanAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanAttributes.EntityData.Children = types.NewOrderedMap()
    fanAttributes.EntityData.Leafs = types.NewOrderedMap()
    fanAttributes.EntityData.Leafs.Append("logical_slot", types.YLeaf{"LogicalSlot", fanAttributes.LogicalSlot})
    fanAttributes.EntityData.Leafs.Append("print_fan_header", types.YLeaf{"PrintFanHeader", fanAttributes.PrintFanHeader})
    fanAttributes.EntityData.Leafs.Append("location", types.YLeaf{"Location", fanAttributes.Location})
    fanAttributes.EntityData.Leafs.Append("fru_PID", types.YLeaf{"FruPID", fanAttributes.FruPID})
    fanAttributes.EntityData.Leafs.Append("fans_speed", types.YLeaf{"FansSpeed", fanAttributes.FansSpeed})
    fanAttributes.EntityData.Leafs.Append("fan_header", types.YLeaf{"FanHeader", fanAttributes.FanHeader})
    fanAttributes.EntityData.Leafs.Append("speed_space", types.YLeaf{"SpeedSpace", fanAttributes.SpeedSpace})

    fanAttributes.EntityData.YListKeys = []string {"LogicalSlot"}

    return &(fanAttributes.EntityData)
}

// Environment_All_Location_Power
type Environment_All_Location_Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is slice of Environment_All_Location_Power_PemAttributes.
    PemAttributes []*Environment_All_Location_Power_PemAttributes
}

func (power *Environment_All_Location_Power) GetEntityData() *types.CommonEntityData {
    power.EntityData.YFilter = power.YFilter
    power.EntityData.YangName = "power"
    power.EntityData.BundleName = "cisco_ios_xr"
    power.EntityData.ParentYangName = "location"
    power.EntityData.SegmentPath = "power" + types.AddKeyToken(power.LocIden, "loc_iden")
    power.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/" + power.EntityData.SegmentPath
    power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    power.EntityData.Children = types.NewOrderedMap()
    power.EntityData.Children.Append("pem_attributes", types.YChild{"PemAttributes", nil})
    for i := range power.PemAttributes {
        power.EntityData.Children.Append(types.GetSegmentPath(power.PemAttributes[i]), types.YChild{"PemAttributes", power.PemAttributes[i]})
    }
    power.EntityData.Leafs = types.NewOrderedMap()
    power.EntityData.Leafs.Append("loc_iden", types.YLeaf{"LocIden", power.LocIden})

    power.EntityData.YListKeys = []string {"LocIden"}

    return &(power.EntityData)
}

// Environment_All_Location_Power_PemAttributes
type Environment_All_Location_Power_PemAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (pemAttributes *Environment_All_Location_Power_PemAttributes) GetEntityData() *types.CommonEntityData {
    pemAttributes.EntityData.YFilter = pemAttributes.YFilter
    pemAttributes.EntityData.YangName = "pem_attributes"
    pemAttributes.EntityData.BundleName = "cisco_ios_xr"
    pemAttributes.EntityData.ParentYangName = "power"
    pemAttributes.EntityData.SegmentPath = "pem_attributes" + types.AddKeyToken(pemAttributes.Pem, "pem")
    pemAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/power/" + pemAttributes.EntityData.SegmentPath
    pemAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pemAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pemAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pemAttributes.EntityData.Children = types.NewOrderedMap()
    pemAttributes.EntityData.Leafs = types.NewOrderedMap()
    pemAttributes.EntityData.Leafs.Append("pem", types.YLeaf{"Pem", pemAttributes.Pem})
    pemAttributes.EntityData.Leafs.Append("pem_id", types.YLeaf{"PemId", pemAttributes.PemId})
    pemAttributes.EntityData.Leafs.Append("card_type", types.YLeaf{"CardType", pemAttributes.CardType})
    pemAttributes.EntityData.Leafs.Append("ps_type", types.YLeaf{"PsType", pemAttributes.PsType})
    pemAttributes.EntityData.Leafs.Append("shelf_num", types.YLeaf{"ShelfNum", pemAttributes.ShelfNum})
    pemAttributes.EntityData.Leafs.Append("supply_type", types.YLeaf{"SupplyType", pemAttributes.SupplyType})
    pemAttributes.EntityData.Leafs.Append("input_voltage", types.YLeaf{"InputVoltage", pemAttributes.InputVoltage})
    pemAttributes.EntityData.Leafs.Append("input_current", types.YLeaf{"InputCurrent", pemAttributes.InputCurrent})
    pemAttributes.EntityData.Leafs.Append("output_voltage", types.YLeaf{"OutputVoltage", pemAttributes.OutputVoltage})
    pemAttributes.EntityData.Leafs.Append("output_current", types.YLeaf{"OutputCurrent", pemAttributes.OutputCurrent})
    pemAttributes.EntityData.Leafs.Append("status", types.YLeaf{"Status", pemAttributes.Status})
    pemAttributes.EntityData.Leafs.Append("input_power_to_ps", types.YLeaf{"InputPowerToPs", pemAttributes.InputPowerToPs})
    pemAttributes.EntityData.Leafs.Append("input_current_to_ps", types.YLeaf{"InputCurrentToPs", pemAttributes.InputCurrentToPs})
    pemAttributes.EntityData.Leafs.Append("output_power_from_ps", types.YLeaf{"OutputPowerFromPs", pemAttributes.OutputPowerFromPs})
    pemAttributes.EntityData.Leafs.Append("output_current_from_ps", types.YLeaf{"OutputCurrentFromPs", pemAttributes.OutputCurrentFromPs})
    pemAttributes.EntityData.Leafs.Append("power_allocated", types.YLeaf{"PowerAllocated", pemAttributes.PowerAllocated})
    pemAttributes.EntityData.Leafs.Append("power_consumed", types.YLeaf{"PowerConsumed", pemAttributes.PowerConsumed})
    pemAttributes.EntityData.Leafs.Append("power_status", types.YLeaf{"PowerStatus", pemAttributes.PowerStatus})
    pemAttributes.EntityData.Leafs.Append("confgd_power_redundancy_mode", types.YLeaf{"ConfgdPowerRedundancyMode", pemAttributes.ConfgdPowerRedundancyMode})
    pemAttributes.EntityData.Leafs.Append("usable_power_capacity", types.YLeaf{"UsablePowerCapacity", pemAttributes.UsablePowerCapacity})
    pemAttributes.EntityData.Leafs.Append("protection_power_capacity", types.YLeaf{"ProtectionPowerCapacity", pemAttributes.ProtectionPowerCapacity})
    pemAttributes.EntityData.Leafs.Append("power_resrv_and_alloc", types.YLeaf{"PowerResrvAndAlloc", pemAttributes.PowerResrvAndAlloc})
    pemAttributes.EntityData.Leafs.Append("system_power_used", types.YLeaf{"SystemPowerUsed", pemAttributes.SystemPowerUsed})
    pemAttributes.EntityData.Leafs.Append("system_power_input", types.YLeaf{"SystemPowerInput", pemAttributes.SystemPowerInput})
    pemAttributes.EntityData.Leafs.Append("power_level", types.YLeaf{"PowerLevel", pemAttributes.PowerLevel})
    pemAttributes.EntityData.Leafs.Append("output_header", types.YLeaf{"OutputHeader", pemAttributes.OutputHeader})
    pemAttributes.EntityData.Leafs.Append("output_footer", types.YLeaf{"OutputFooter", pemAttributes.OutputFooter})
    pemAttributes.EntityData.Leafs.Append("ps_sum_footer", types.YLeaf{"PsSumFooter", pemAttributes.PsSumFooter})

    pemAttributes.EntityData.YListKeys = []string {"Pem"}

    return &(pemAttributes.EntityData)
}

// Environment_All_Location_Altitude
type Environment_All_Location_Altitude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocIden interface{}

    // The type is slice of Environment_All_Location_Altitude_AltAttributes.
    AltAttributes []*Environment_All_Location_Altitude_AltAttributes
}

func (altitude *Environment_All_Location_Altitude) GetEntityData() *types.CommonEntityData {
    altitude.EntityData.YFilter = altitude.YFilter
    altitude.EntityData.YangName = "altitude"
    altitude.EntityData.BundleName = "cisco_ios_xr"
    altitude.EntityData.ParentYangName = "location"
    altitude.EntityData.SegmentPath = "altitude" + types.AddKeyToken(altitude.LocIden, "loc_iden")
    altitude.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/" + altitude.EntityData.SegmentPath
    altitude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altitude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altitude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altitude.EntityData.Children = types.NewOrderedMap()
    altitude.EntityData.Children.Append("alt_attributes", types.YChild{"AltAttributes", nil})
    for i := range altitude.AltAttributes {
        altitude.EntityData.Children.Append(types.GetSegmentPath(altitude.AltAttributes[i]), types.YChild{"AltAttributes", altitude.AltAttributes[i]})
    }
    altitude.EntityData.Leafs = types.NewOrderedMap()
    altitude.EntityData.Leafs.Append("loc_iden", types.YLeaf{"LocIden", altitude.LocIden})

    altitude.EntityData.YListKeys = []string {"LocIden"}

    return &(altitude.EntityData)
}

// Environment_All_Location_Altitude_AltAttributes
type Environment_All_Location_Altitude_AltAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Sensor interface{}

    // The type is bool. The default value is false.
    PrintHeader interface{}

    // The type is interface{} with range: 0..4294967295.
    Rack interface{}

    // The type is string.
    SensorValue interface{}

    // The type is string.
    Source interface{}
}

func (altAttributes *Environment_All_Location_Altitude_AltAttributes) GetEntityData() *types.CommonEntityData {
    altAttributes.EntityData.YFilter = altAttributes.YFilter
    altAttributes.EntityData.YangName = "alt_attributes"
    altAttributes.EntityData.BundleName = "cisco_ios_xr"
    altAttributes.EntityData.ParentYangName = "altitude"
    altAttributes.EntityData.SegmentPath = "alt_attributes" + types.AddKeyToken(altAttributes.Sensor, "sensor")
    altAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/all/location/altitude/" + altAttributes.EntityData.SegmentPath
    altAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altAttributes.EntityData.Children = types.NewOrderedMap()
    altAttributes.EntityData.Leafs = types.NewOrderedMap()
    altAttributes.EntityData.Leafs.Append("sensor", types.YLeaf{"Sensor", altAttributes.Sensor})
    altAttributes.EntityData.Leafs.Append("print_header", types.YLeaf{"PrintHeader", altAttributes.PrintHeader})
    altAttributes.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", altAttributes.Rack})
    altAttributes.EntityData.Leafs.Append("sensor_value", types.YLeaf{"SensorValue", altAttributes.SensorValue})
    altAttributes.EntityData.Leafs.Append("source", types.YLeaf{"Source", altAttributes.Source})

    altAttributes.EntityData.YListKeys = []string {"Sensor"}

    return &(altAttributes.EntityData)
}

// Environment_Config
// environment configurational data
type Environment_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Router Environment_Config_Router

    
    AirFilter Environment_Config_AirFilter

    
    FanCtrl Environment_Config_FanCtrl

    
    Temperature Environment_Config_Temperature

    
    Monitoring Environment_Config_Monitoring

    
    RaiseFanSpeed Environment_Config_RaiseFanSpeed

    
    FanCtrlOptics Environment_Config_FanCtrlOptics

    
    GracefulShutdown Environment_Config_GracefulShutdown
}

func (config *Environment_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "environment"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("router", types.YChild{"Router", &config.Router})
    config.EntityData.Children.Append("air-filter", types.YChild{"AirFilter", &config.AirFilter})
    config.EntityData.Children.Append("fan-ctrl", types.YChild{"FanCtrl", &config.FanCtrl})
    config.EntityData.Children.Append("temperature", types.YChild{"Temperature", &config.Temperature})
    config.EntityData.Children.Append("monitoring", types.YChild{"Monitoring", &config.Monitoring})
    config.EntityData.Children.Append("raise-fan-speed", types.YChild{"RaiseFanSpeed", &config.RaiseFanSpeed})
    config.EntityData.Children.Append("fan-ctrl-optics", types.YChild{"FanCtrlOptics", &config.FanCtrlOptics})
    config.EntityData.Children.Append("graceful-shutdown", types.YChild{"GracefulShutdown", &config.GracefulShutdown})
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

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
    router.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + router.EntityData.SegmentPath
    router.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    router.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    router.EntityData.Children = types.NewOrderedMap()
    router.EntityData.Children.Append("altitude", types.YChild{"Altitude", &router.Altitude})
    router.EntityData.Leafs = types.NewOrderedMap()

    router.EntityData.YListKeys = []string {}

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
    altitude.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/router/" + altitude.EntityData.SegmentPath
    altitude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    altitude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    altitude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    altitude.EntityData.Children = types.NewOrderedMap()
    altitude.EntityData.Children.Append("all", types.YChild{"All", &altitude.All})
    altitude.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &altitude.RackLoc})
    altitude.EntityData.Leafs = types.NewOrderedMap()

    altitude.EntityData.YListKeys = []string {}

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
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/router/altitude/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("num_meters", types.YLeaf{"NumMeters", all.NumMeters})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Environment_Config_Router_Altitude_RackLoc
type Environment_Config_Router_Altitude_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Config_Router_Altitude_RackLoc_Location.
    Location []*Environment_Config_Router_Altitude_RackLoc_Location
}

func (rackLoc *Environment_Config_Router_Altitude_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "altitude"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/router/altitude/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_Router_Altitude_RackLoc_Location
type Environment_Config_Router_Altitude_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}

    // The type is interface{} with range: 1..4000.
    NumMeters interface{}
}

func (location *Environment_Config_Router_Altitude_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/router/altitude/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})
    location.EntityData.Leafs.Append("num_meters", types.YLeaf{"NumMeters", location.NumMeters})

    location.EntityData.YListKeys = []string {"RackId"}

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
    airFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + airFilter.EntityData.SegmentPath
    airFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    airFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    airFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    airFilter.EntityData.Children = types.NewOrderedMap()
    airFilter.EntityData.Children.Append("replaced", types.YChild{"Replaced", &airFilter.Replaced})
    airFilter.EntityData.Leafs = types.NewOrderedMap()

    airFilter.EntityData.YListKeys = []string {}

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
    replaced.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/air-filter/" + replaced.EntityData.SegmentPath
    replaced.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replaced.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replaced.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replaced.EntityData.Children = types.NewOrderedMap()
    replaced.EntityData.Children.Append("all", types.YChild{"All", &replaced.All})
    replaced.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &replaced.RackLoc})
    replaced.EntityData.Leafs = types.NewOrderedMap()

    replaced.EntityData.YListKeys = []string {}

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
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/air-filter/replaced/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("date", types.YLeaf{"Date", all.Date})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Environment_Config_AirFilter_Replaced_RackLoc
type Environment_Config_AirFilter_Replaced_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Environment_Config_AirFilter_Replaced_RackLoc_Location.
    Location []*Environment_Config_AirFilter_Replaced_RackLoc_Location
}

func (rackLoc *Environment_Config_AirFilter_Replaced_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "replaced"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/air-filter/replaced/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_AirFilter_Replaced_RackLoc_Location
type Environment_Config_AirFilter_Replaced_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}

    // The type is string.
    Date interface{}
}

func (location *Environment_Config_AirFilter_Replaced_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/air-filter/replaced/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})
    location.EntityData.Leafs.Append("date", types.YLeaf{"Date", location.Date})

    location.EntityData.YListKeys = []string {"RackId"}

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
    fanCtrl.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + fanCtrl.EntityData.SegmentPath
    fanCtrl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanCtrl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanCtrl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanCtrl.EntityData.Children = types.NewOrderedMap()
    fanCtrl.EntityData.Children.Append("disable", types.YChild{"Disable", &fanCtrl.Disable})
    fanCtrl.EntityData.Leafs = types.NewOrderedMap()

    fanCtrl.EntityData.YListKeys = []string {}

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
    disable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/fan-ctrl/" + disable.EntityData.SegmentPath
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = types.NewOrderedMap()
    disable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &disable.RackLoc})
    disable.EntityData.Leafs = types.NewOrderedMap()

    disable.EntityData.YListKeys = []string {}

    return &(disable.EntityData)
}

// Environment_Config_FanCtrl_Disable_RackLoc
type Environment_Config_FanCtrl_Disable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of Environment_Config_FanCtrl_Disable_RackLoc_Location.
    Location []*Environment_Config_FanCtrl_Disable_RackLoc_Location
}

func (rackLoc *Environment_Config_FanCtrl_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/fan-ctrl/disable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_FanCtrl_Disable_RackLoc_Location
type Environment_Config_FanCtrl_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *Environment_Config_FanCtrl_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/fan-ctrl/disable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

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
    temperature.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + temperature.EntityData.SegmentPath
    temperature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    temperature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    temperature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    temperature.EntityData.Children = types.NewOrderedMap()
    temperature.EntityData.Children.Append("disable", types.YChild{"Disable", &temperature.Disable})
    temperature.EntityData.Leafs = types.NewOrderedMap()

    temperature.EntityData.YListKeys = []string {}

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
    disable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/temperature/" + disable.EntityData.SegmentPath
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = types.NewOrderedMap()
    disable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &disable.RackLoc})
    disable.EntityData.Leafs = types.NewOrderedMap()

    disable.EntityData.YListKeys = []string {}

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
    Location []*Environment_Config_Temperature_Disable_RackLoc_Location
}

func (rackLoc *Environment_Config_Temperature_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/temperature/disable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_Temperature_Disable_RackLoc_Location
type Environment_Config_Temperature_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *Environment_Config_Temperature_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/temperature/disable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

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
    monitoring.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + monitoring.EntityData.SegmentPath
    monitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitoring.EntityData.Children = types.NewOrderedMap()
    monitoring.EntityData.Children.Append("disable", types.YChild{"Disable", &monitoring.Disable})
    monitoring.EntityData.Leafs = types.NewOrderedMap()

    monitoring.EntityData.YListKeys = []string {}

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
    disable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/monitoring/" + disable.EntityData.SegmentPath
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = types.NewOrderedMap()
    disable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &disable.RackLoc})
    disable.EntityData.Leafs = types.NewOrderedMap()

    disable.EntityData.YListKeys = []string {}

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
    Location []*Environment_Config_Monitoring_Disable_RackLoc_Location
}

func (rackLoc *Environment_Config_Monitoring_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/monitoring/disable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_Monitoring_Disable_RackLoc_Location
type Environment_Config_Monitoring_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *Environment_Config_Monitoring_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/monitoring/disable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

    return &(location.EntityData)
}

// Environment_Config_RaiseFanSpeed
type Environment_Config_RaiseFanSpeed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Environment_Config_RaiseFanSpeed_All

    
    RackLoc Environment_Config_RaiseFanSpeed_RackLoc
}

func (raiseFanSpeed *Environment_Config_RaiseFanSpeed) GetEntityData() *types.CommonEntityData {
    raiseFanSpeed.EntityData.YFilter = raiseFanSpeed.YFilter
    raiseFanSpeed.EntityData.YangName = "raise-fan-speed"
    raiseFanSpeed.EntityData.BundleName = "cisco_ios_xr"
    raiseFanSpeed.EntityData.ParentYangName = "config"
    raiseFanSpeed.EntityData.SegmentPath = "raise-fan-speed"
    raiseFanSpeed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + raiseFanSpeed.EntityData.SegmentPath
    raiseFanSpeed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raiseFanSpeed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raiseFanSpeed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raiseFanSpeed.EntityData.Children = types.NewOrderedMap()
    raiseFanSpeed.EntityData.Children.Append("all", types.YChild{"All", &raiseFanSpeed.All})
    raiseFanSpeed.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &raiseFanSpeed.RackLoc})
    raiseFanSpeed.EntityData.Leafs = types.NewOrderedMap()

    raiseFanSpeed.EntityData.YListKeys = []string {}

    return &(raiseFanSpeed.EntityData)
}

// Environment_Config_RaiseFanSpeed_All
type Environment_Config_RaiseFanSpeed_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..100.
    SpeedPwm interface{}
}

func (all *Environment_Config_RaiseFanSpeed_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "raise-fan-speed"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/raise-fan-speed/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("speed_pwm", types.YLeaf{"SpeedPwm", all.SpeedPwm})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Environment_Config_RaiseFanSpeed_RackLoc
type Environment_Config_RaiseFanSpeed_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Environment_Config_RaiseFanSpeed_RackLoc_Location.
    Location []*Environment_Config_RaiseFanSpeed_RackLoc_Location
}

func (rackLoc *Environment_Config_RaiseFanSpeed_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "raise-fan-speed"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/raise-fan-speed/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_RaiseFanSpeed_RackLoc_Location
type Environment_Config_RaiseFanSpeed_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}

    // The type is interface{} with range: 0..100.
    SpeedPwm interface{}
}

func (location *Environment_Config_RaiseFanSpeed_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/raise-fan-speed/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})
    location.EntityData.Leafs.Append("speed_pwm", types.YLeaf{"SpeedPwm", location.SpeedPwm})

    location.EntityData.YListKeys = []string {"RackId"}

    return &(location.EntityData)
}

// Environment_Config_FanCtrlOptics
type Environment_Config_FanCtrlOptics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Enable Environment_Config_FanCtrlOptics_Enable
}

func (fanCtrlOptics *Environment_Config_FanCtrlOptics) GetEntityData() *types.CommonEntityData {
    fanCtrlOptics.EntityData.YFilter = fanCtrlOptics.YFilter
    fanCtrlOptics.EntityData.YangName = "fan-ctrl-optics"
    fanCtrlOptics.EntityData.BundleName = "cisco_ios_xr"
    fanCtrlOptics.EntityData.ParentYangName = "config"
    fanCtrlOptics.EntityData.SegmentPath = "fan-ctrl-optics"
    fanCtrlOptics.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + fanCtrlOptics.EntityData.SegmentPath
    fanCtrlOptics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fanCtrlOptics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fanCtrlOptics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fanCtrlOptics.EntityData.Children = types.NewOrderedMap()
    fanCtrlOptics.EntityData.Children.Append("enable", types.YChild{"Enable", &fanCtrlOptics.Enable})
    fanCtrlOptics.EntityData.Leafs = types.NewOrderedMap()

    fanCtrlOptics.EntityData.YListKeys = []string {}

    return &(fanCtrlOptics.EntityData)
}

// Environment_Config_FanCtrlOptics_Enable
type Environment_Config_FanCtrlOptics_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc Environment_Config_FanCtrlOptics_Enable_RackLoc
}

func (enable *Environment_Config_FanCtrlOptics_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "fan-ctrl-optics"
    enable.EntityData.SegmentPath = "enable"
    enable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/fan-ctrl-optics/" + enable.EntityData.SegmentPath
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = types.NewOrderedMap()
    enable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &enable.RackLoc})
    enable.EntityData.Leafs = types.NewOrderedMap()

    enable.EntityData.YListKeys = []string {}

    return &(enable.EntityData)
}

// Environment_Config_FanCtrlOptics_Enable_RackLoc
type Environment_Config_FanCtrlOptics_Enable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of
    // Environment_Config_FanCtrlOptics_Enable_RackLoc_Location.
    Location []*Environment_Config_FanCtrlOptics_Enable_RackLoc_Location
}

func (rackLoc *Environment_Config_FanCtrlOptics_Enable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "enable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/fan-ctrl-optics/enable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_FanCtrlOptics_Enable_RackLoc_Location
type Environment_Config_FanCtrlOptics_Enable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *Environment_Config_FanCtrlOptics_Enable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/fan-ctrl-optics/enable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

    return &(location.EntityData)
}

// Environment_Config_GracefulShutdown
type Environment_Config_GracefulShutdown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Disable Environment_Config_GracefulShutdown_Disable
}

func (gracefulShutdown *Environment_Config_GracefulShutdown) GetEntityData() *types.CommonEntityData {
    gracefulShutdown.EntityData.YFilter = gracefulShutdown.YFilter
    gracefulShutdown.EntityData.YangName = "graceful-shutdown"
    gracefulShutdown.EntityData.BundleName = "cisco_ios_xr"
    gracefulShutdown.EntityData.ParentYangName = "config"
    gracefulShutdown.EntityData.SegmentPath = "graceful-shutdown"
    gracefulShutdown.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/" + gracefulShutdown.EntityData.SegmentPath
    gracefulShutdown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gracefulShutdown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gracefulShutdown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gracefulShutdown.EntityData.Children = types.NewOrderedMap()
    gracefulShutdown.EntityData.Children.Append("disable", types.YChild{"Disable", &gracefulShutdown.Disable})
    gracefulShutdown.EntityData.Leafs = types.NewOrderedMap()

    gracefulShutdown.EntityData.YListKeys = []string {}

    return &(gracefulShutdown.EntityData)
}

// Environment_Config_GracefulShutdown_Disable
type Environment_Config_GracefulShutdown_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    RackLoc Environment_Config_GracefulShutdown_Disable_RackLoc
}

func (disable *Environment_Config_GracefulShutdown_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "graceful-shutdown"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/graceful-shutdown/" + disable.EntityData.SegmentPath
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = types.NewOrderedMap()
    disable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &disable.RackLoc})
    disable.EntityData.Leafs = types.NewOrderedMap()

    disable.EntityData.YListKeys = []string {}

    return &(disable.EntityData)
}

// Environment_Config_GracefulShutdown_Disable_RackLoc
type Environment_Config_GracefulShutdown_Disable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of
    // Environment_Config_GracefulShutdown_Disable_RackLoc_Location.
    Location []*Environment_Config_GracefulShutdown_Disable_RackLoc_Location
}

func (rackLoc *Environment_Config_GracefulShutdown_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/graceful-shutdown/disable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// Environment_Config_GracefulShutdown_Disable_RackLoc_Location
type Environment_Config_GracefulShutdown_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *Environment_Config_GracefulShutdown_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/config/graceful-shutdown/disable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

    return &(location.EntityData)
}

// Environment_Trace
// show traceable processes
type Environment_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Environment_Trace_Location.
    Location []*Environment_Trace_Location
}

func (trace *Environment_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "environment"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/" + trace.EntityData.SegmentPath
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range trace.Location {
        trace.EntityData.Children.Append(types.GetSegmentPath(trace.Location[i]), types.YChild{"Location", trace.Location[i]})
    }
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("buffer", types.YLeaf{"Buffer", trace.Buffer})

    trace.EntityData.YListKeys = []string {"Buffer"}

    return &(trace.EntityData)
}

// Environment_Trace_Location
type Environment_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Environment_Trace_Location_AllOptions.
    AllOptions []*Environment_Trace_Location_AllOptions
}

func (location *Environment_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/trace/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("all-options", types.YChild{"AllOptions", nil})
    for i := range location.AllOptions {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AllOptions[i]), types.YChild{"AllOptions", location.AllOptions[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location_name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Environment_Trace_Location_AllOptions
type Environment_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Environment_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Environment_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Environment_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/trace/location/" + allOptions.EntityData.SegmentPath
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        types.SetYListKey(allOptions.TraceBlocks[i], i)
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Environment_Trace_Location_AllOptions_TraceBlocks
type Environment_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Environment_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:environment/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

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
    powerMgmt.EntityData.AbsolutePath = powerMgmt.EntityData.SegmentPath
    powerMgmt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    powerMgmt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    powerMgmt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    powerMgmt.EntityData.Children = types.NewOrderedMap()
    powerMgmt.EntityData.Children.Append("config", types.YChild{"Config", &powerMgmt.Config})
    powerMgmt.EntityData.Leafs = types.NewOrderedMap()

    powerMgmt.EntityData.YListKeys = []string {}

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

    
    Progressive PowerMgmt_Config_Progressive
}

func (config *PowerMgmt_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "power-mgmt"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("action", types.YChild{"Action", &config.Action})
    config.EntityData.Children.Append("single-feed-mode", types.YChild{"SingleFeedMode", &config.SingleFeedMode})
    config.EntityData.Children.Append("extended-temp", types.YChild{"ExtendedTemp", &config.ExtendedTemp})
    config.EntityData.Children.Append("redundancy-num-pms", types.YChild{"RedundancyNumPms", &config.RedundancyNumPms})
    config.EntityData.Children.Append("progressive", types.YChild{"Progressive", &config.Progressive})
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

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
    action.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/" + action.EntityData.SegmentPath
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Children.Append("disable", types.YChild{"Disable", &action.Disable})
    action.EntityData.Leafs = types.NewOrderedMap()

    action.EntityData.YListKeys = []string {}

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
    disable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/action/" + disable.EntityData.SegmentPath
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = types.NewOrderedMap()
    disable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &disable.RackLoc})
    disable.EntityData.Leafs = types.NewOrderedMap()

    disable.EntityData.YListKeys = []string {}

    return &(disable.EntityData)
}

// PowerMgmt_Config_Action_Disable_RackLoc
type PowerMgmt_Config_Action_Disable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of PowerMgmt_Config_Action_Disable_RackLoc_Location.
    Location []*PowerMgmt_Config_Action_Disable_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_Action_Disable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "disable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/action/disable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_Action_Disable_RackLoc_Location
type PowerMgmt_Config_Action_Disable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *PowerMgmt_Config_Action_Disable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/action/disable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

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
    singleFeedMode.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/" + singleFeedMode.EntityData.SegmentPath
    singleFeedMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleFeedMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleFeedMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleFeedMode.EntityData.Children = types.NewOrderedMap()
    singleFeedMode.EntityData.Children.Append("enable", types.YChild{"Enable", &singleFeedMode.Enable})
    singleFeedMode.EntityData.Leafs = types.NewOrderedMap()

    singleFeedMode.EntityData.YListKeys = []string {}

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
    enable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/single-feed-mode/" + enable.EntityData.SegmentPath
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = types.NewOrderedMap()
    enable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &enable.RackLoc})
    enable.EntityData.Leafs = types.NewOrderedMap()

    enable.EntityData.YListKeys = []string {}

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
    Location []*PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_SingleFeedMode_Enable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "enable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/single-feed-mode/enable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location
type PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *PowerMgmt_Config_SingleFeedMode_Enable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/single-feed-mode/enable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

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
    extendedTemp.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/" + extendedTemp.EntityData.SegmentPath
    extendedTemp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedTemp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedTemp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedTemp.EntityData.Children = types.NewOrderedMap()
    extendedTemp.EntityData.Children.Append("enable", types.YChild{"Enable", &extendedTemp.Enable})
    extendedTemp.EntityData.Leafs = types.NewOrderedMap()

    extendedTemp.EntityData.YListKeys = []string {}

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
    enable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/extended-temp/" + enable.EntityData.SegmentPath
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = types.NewOrderedMap()
    enable.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &enable.RackLoc})
    enable.EntityData.Leafs = types.NewOrderedMap()

    enable.EntityData.YListKeys = []string {}

    return &(enable.EntityData)
}

// PowerMgmt_Config_ExtendedTemp_Enable_RackLoc
type PowerMgmt_Config_ExtendedTemp_Enable_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    All interface{}

    // The type is slice of PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location.
    Location []*PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_ExtendedTemp_Enable_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "enable"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/extended-temp/enable/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()
    rackLoc.EntityData.Leafs.Append("all", types.YLeaf{"All", rackLoc.All})

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location
type PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}
}

func (location *PowerMgmt_Config_ExtendedTemp_Enable_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/extended-temp/enable/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})

    location.EntityData.YListKeys = []string {"RackId"}

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
    redundancyNumPms.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/" + redundancyNumPms.EntityData.SegmentPath
    redundancyNumPms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancyNumPms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancyNumPms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancyNumPms.EntityData.Children = types.NewOrderedMap()
    redundancyNumPms.EntityData.Children.Append("all", types.YChild{"All", &redundancyNumPms.All})
    redundancyNumPms.EntityData.Children.Append("rack_loc", types.YChild{"RackLoc", &redundancyNumPms.RackLoc})
    redundancyNumPms.EntityData.Leafs = types.NewOrderedMap()

    redundancyNumPms.EntityData.YListKeys = []string {}

    return &(redundancyNumPms.EntityData)
}

// PowerMgmt_Config_RedundancyNumPms_All
type PowerMgmt_Config_RedundancyNumPms_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..5.
    NumPm interface{}
}

func (all *PowerMgmt_Config_RedundancyNumPms_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "redundancy-num-pms"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/redundancy-num-pms/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("num_pm", types.YLeaf{"NumPm", all.NumPm})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// PowerMgmt_Config_RedundancyNumPms_RackLoc
type PowerMgmt_Config_RedundancyNumPms_RackLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of PowerMgmt_Config_RedundancyNumPms_RackLoc_Location.
    Location []*PowerMgmt_Config_RedundancyNumPms_RackLoc_Location
}

func (rackLoc *PowerMgmt_Config_RedundancyNumPms_RackLoc) GetEntityData() *types.CommonEntityData {
    rackLoc.EntityData.YFilter = rackLoc.YFilter
    rackLoc.EntityData.YangName = "rack_loc"
    rackLoc.EntityData.BundleName = "cisco_ios_xr"
    rackLoc.EntityData.ParentYangName = "redundancy-num-pms"
    rackLoc.EntityData.SegmentPath = "rack_loc"
    rackLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/redundancy-num-pms/" + rackLoc.EntityData.SegmentPath
    rackLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLoc.EntityData.Children = types.NewOrderedMap()
    rackLoc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range rackLoc.Location {
        rackLoc.EntityData.Children.Append(types.GetSegmentPath(rackLoc.Location[i]), types.YChild{"Location", rackLoc.Location[i]})
    }
    rackLoc.EntityData.Leafs = types.NewOrderedMap()

    rackLoc.EntityData.YListKeys = []string {}

    return &(rackLoc.EntityData)
}

// PowerMgmt_Config_RedundancyNumPms_RackLoc_Location
type PowerMgmt_Config_RedundancyNumPms_RackLoc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is RackId.
    RackId interface{}

    // The type is interface{} with range: 0..5.
    NumPm interface{}
}

func (location *PowerMgmt_Config_RedundancyNumPms_RackLoc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rack_loc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.RackId, "rackId")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/redundancy-num-pms/rack_loc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rackId", types.YLeaf{"RackId", location.RackId})
    location.EntityData.Leafs.Append("num_pm", types.YLeaf{"NumPm", location.NumPm})

    location.EntityData.YListKeys = []string {"RackId"}

    return &(location.EntityData)
}

// PowerMgmt_Config_Progressive
type PowerMgmt_Config_Progressive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of PowerMgmt_Config_Progressive_Enable.
    Enable []*PowerMgmt_Config_Progressive_Enable
}

func (progressive *PowerMgmt_Config_Progressive) GetEntityData() *types.CommonEntityData {
    progressive.EntityData.YFilter = progressive.YFilter
    progressive.EntityData.YangName = "progressive"
    progressive.EntityData.BundleName = "cisco_ios_xr"
    progressive.EntityData.ParentYangName = "config"
    progressive.EntityData.SegmentPath = "progressive"
    progressive.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/" + progressive.EntityData.SegmentPath
    progressive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    progressive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    progressive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    progressive.EntityData.Children = types.NewOrderedMap()
    progressive.EntityData.Children.Append("enable", types.YChild{"Enable", nil})
    for i := range progressive.Enable {
        progressive.EntityData.Children.Append(types.GetSegmentPath(progressive.Enable[i]), types.YChild{"Enable", progressive.Enable[i]})
    }
    progressive.EntityData.Leafs = types.NewOrderedMap()

    progressive.EntityData.YListKeys = []string {}

    return &(progressive.EntityData)
}

// PowerMgmt_Config_Progressive_Enable
type PowerMgmt_Config_Progressive_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is Enabled.
    Enabled interface{}

    // The type is interface{} with range: 3000..72000.
    SyslogThreshold interface{}

    // The type is interface{} with range: 3000..72000.
    ShutdownThreshold interface{}

    
    Priority PowerMgmt_Config_Progressive_Enable_Priority
}

func (enable *PowerMgmt_Config_Progressive_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "progressive"
    enable.EntityData.SegmentPath = "enable" + types.AddKeyToken(enable.Enabled, "enabled")
    enable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/progressive/" + enable.EntityData.SegmentPath
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = types.NewOrderedMap()
    enable.EntityData.Children.Append("priority", types.YChild{"Priority", &enable.Priority})
    enable.EntityData.Leafs = types.NewOrderedMap()
    enable.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", enable.Enabled})
    enable.EntityData.Leafs.Append("syslog-threshold", types.YLeaf{"SyslogThreshold", enable.SyslogThreshold})
    enable.EntityData.Leafs.Append("shutdown-threshold", types.YLeaf{"ShutdownThreshold", enable.ShutdownThreshold})

    enable.EntityData.YListKeys = []string {"Enabled"}

    return &(enable.EntityData)
}

// PowerMgmt_Config_Progressive_Enable_Priority
type PowerMgmt_Config_Progressive_Enable_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of PowerMgmt_Config_Progressive_Enable_Priority_Location.
    Location []*PowerMgmt_Config_Progressive_Enable_Priority_Location
}

func (priority *PowerMgmt_Config_Progressive_Enable_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "enable"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/progressive/enable/" + priority.EntityData.SegmentPath
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range priority.Location {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.Location[i]), types.YChild{"Location", priority.Location[i]})
    }
    priority.EntityData.Leafs = types.NewOrderedMap()

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// PowerMgmt_Config_Progressive_Enable_Priority_Location
type PowerMgmt_Config_Progressive_Enable_Priority_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern: 0/([0-9]|1[0-9]).
    Loc interface{}

    // The type is interface{} with range: 0..19.
    Prior interface{}
}

func (location *PowerMgmt_Config_Progressive_Enable_Priority_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "priority"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Loc, "loc")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-envmon-ui:power-mgmt/config/progressive/enable/priority/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("loc", types.YLeaf{"Loc", location.Loc})
    location.EntityData.Leafs.Append("prior", types.YLeaf{"Prior", location.Prior})

    location.EntityData.YListKeys = []string {"Loc"}

    return &(location.EntityData)
}

// PowerMgmt_Config_Progressive_Enable_Enabled
type PowerMgmt_Config_Progressive_Enable_Enabled string

const (
    PowerMgmt_Config_Progressive_Enable_Enabled_enable PowerMgmt_Config_Progressive_Enable_Enabled = "enable"
)

