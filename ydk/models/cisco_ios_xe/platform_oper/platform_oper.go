// This module contains a collection of YANG definitions
// for monitoring of platform components.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package platform_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package platform_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-platform-oper components}", reflect.TypeOf(Components{}))
    ydk.RegisterEntity("Cisco-IOS-XE-platform-oper:components", reflect.TypeOf(Components{}))
}

// PlatformPropValueType represents Property value type
type PlatformPropValueType string

const (
    PlatformPropValueType_property_string PlatformPropValueType = "property-string"

    PlatformPropValueType_property_boolean PlatformPropValueType = "property-boolean"

    PlatformPropValueType_property_int64 PlatformPropValueType = "property-int64"

    PlatformPropValueType_property_uint64 PlatformPropValueType = "property-uint64"

    PlatformPropValueType_property_decimal64 PlatformPropValueType = "property-decimal64"
)

// PlatformCompType represents Component Type
type PlatformCompType string

const (
    PlatformCompType_comp_chassis PlatformCompType = "comp-chassis"

    PlatformCompType_comp_backplane PlatformCompType = "comp-backplane"

    PlatformCompType_comp_power_supply PlatformCompType = "comp-power-supply"

    PlatformCompType_comp_fan PlatformCompType = "comp-fan"

    PlatformCompType_comp_sensor PlatformCompType = "comp-sensor"

    PlatformCompType_comp_module PlatformCompType = "comp-module"

    PlatformCompType_comp_linecard PlatformCompType = "comp-linecard"

    PlatformCompType_comp_port PlatformCompType = "comp-port"

    PlatformCompType_comp_cpu PlatformCompType = "comp-cpu"

    PlatformCompType_comp_operating_system PlatformCompType = "comp-operating-system"

    PlatformCompType_comp_optical_channel PlatformCompType = "comp-optical-channel"

    PlatformCompType_CONTAINER PlatformCompType = "CONTAINER"
)

// Components
// Enclosing container for the components in the system
type Components struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of components, keyed by component name. The type is slice of
    // Components_Component.
    Component []Components_Component
}

func (components *Components) GetFilter() yfilter.YFilter { return components.YFilter }

func (components *Components) SetFilter(yf yfilter.YFilter) { components.YFilter = yf }

func (components *Components) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    return ""
}

func (components *Components) GetSegmentPath() string {
    return "Cisco-IOS-XE-platform-oper:components"
}

func (components *Components) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "component" {
        for _, c := range components.Component {
            if components.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Components_Component{}
        components.Component = append(components.Component, child)
        return &components.Component[len(components.Component)-1]
    }
    return nil
}

func (components *Components) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range components.Component {
        children[components.Component[i].GetSegmentPath()] = &components.Component[i]
    }
    return children
}

func (components *Components) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (components *Components) GetBundleName() string { return "cisco_ios_xe" }

func (components *Components) GetYangName() string { return "components" }

func (components *Components) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (components *Components) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (components *Components) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (components *Components) SetParent(parent types.Entity) { components.parent = parent }

func (components *Components) GetParent() types.Entity { return components.parent }

func (components *Components) GetParentYangName() string { return "Cisco-IOS-XE-platform-oper" }

// Components_Component
// List of components, keyed by component name
type Components_Component struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. References component name. The type is string.
    Cname interface{}

    // Operational state data for each component.
    State Components_Component_State

    // Platform component properties.
    PlatformProperties Components_Component_PlatformProperties

    // Platform subcomponents.
    PlatformSubcomponents Components_Component_PlatformSubcomponents
}

func (component *Components_Component) GetFilter() yfilter.YFilter { return component.YFilter }

func (component *Components_Component) SetFilter(yf yfilter.YFilter) { component.YFilter = yf }

func (component *Components_Component) GetGoName(yname string) string {
    if yname == "cname" { return "Cname" }
    if yname == "state" { return "State" }
    if yname == "platform-properties" { return "PlatformProperties" }
    if yname == "platform-subcomponents" { return "PlatformSubcomponents" }
    return ""
}

func (component *Components_Component) GetSegmentPath() string {
    return "component" + "[cname='" + fmt.Sprintf("%v", component.Cname) + "']"
}

func (component *Components_Component) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &component.State
    }
    if childYangName == "platform-properties" {
        return &component.PlatformProperties
    }
    if childYangName == "platform-subcomponents" {
        return &component.PlatformSubcomponents
    }
    return nil
}

func (component *Components_Component) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &component.State
    children["platform-properties"] = &component.PlatformProperties
    children["platform-subcomponents"] = &component.PlatformSubcomponents
    return children
}

func (component *Components_Component) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cname"] = component.Cname
    return leafs
}

func (component *Components_Component) GetBundleName() string { return "cisco_ios_xe" }

func (component *Components_Component) GetYangName() string { return "component" }

func (component *Components_Component) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (component *Components_Component) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (component *Components_Component) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (component *Components_Component) SetParent(parent types.Entity) { component.parent = parent }

func (component *Components_Component) GetParent() types.Entity { return component.parent }

func (component *Components_Component) GetParentYangName() string { return "components" }

// Components_Component_State
// Operational state data for each component
type Components_Component_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of component as identified by the system. The type is
    // PlatformCompType.
    Type interface{}

    // Unique identifier assigned to the component by the system. The type is
    // string.
    Id interface{}

    // System-supplied description of the component. The type is string.
    Description interface{}

    // System-supplied identifier for the manufacturer of the component.  This
    // data is particularly useful when a component manufacturer is different than
    // the overall device vendor. The type is string.
    MfgName interface{}

    // System-defined version string for a hardware, firmware, or software
    // component. The type is string.
    Version interface{}

    // System-assigned serial number of the component. The type is string.
    SerialNo interface{}

    // System-assigned part number for the component.  This should be present in
    // particular if the component is also an FRU (field replacable unit). The
    // type is string.
    PartNo interface{}

    // Temperature in degrees Celsius of the component. Values include the
    // instantaneous, average, minimum, and maximum statistics. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    Temp Components_Component_State_Temp
}

func (state *Components_Component_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_State) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "id" { return "Id" }
    if yname == "description" { return "Description" }
    if yname == "mfg-name" { return "MfgName" }
    if yname == "version" { return "Version" }
    if yname == "serial-no" { return "SerialNo" }
    if yname == "part-no" { return "PartNo" }
    if yname == "temp" { return "Temp" }
    return ""
}

func (state *Components_Component_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "temp" {
        return &state.Temp
    }
    return nil
}

func (state *Components_Component_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["temp"] = &state.Temp
    return children
}

func (state *Components_Component_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = state.Type
    leafs["id"] = state.Id
    leafs["description"] = state.Description
    leafs["mfg-name"] = state.MfgName
    leafs["version"] = state.Version
    leafs["serial-no"] = state.SerialNo
    leafs["part-no"] = state.PartNo
    return leafs
}

func (state *Components_Component_State) GetBundleName() string { return "cisco_ios_xe" }

func (state *Components_Component_State) GetYangName() string { return "state" }

func (state *Components_Component_State) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (state *Components_Component_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (state *Components_Component_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (state *Components_Component_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_State) GetParentYangName() string { return "component" }

// Components_Component_State_Temp
// Temperature in degrees Celsius of the component. Values include
// the instantaneous, average, minimum, and maximum statistics. If
// avg/min/max statistics are not supported, the target is expected
// to just supply the instant value
type Components_Component_State_Temp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instantaneous temperature value of a component. The type is string with
    // range: -92233720368547758.08..92233720368547758.07.
    TempInstant interface{}

    // Arithmetic mean value of the statistic over a sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07.
    TempAvg interface{}

    // High water mark value of the statistic over a sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07.
    TempMax interface{}

    // Low water mark value of the statistic over a sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07.
    TempMin interface{}
}

func (temp *Components_Component_State_Temp) GetFilter() yfilter.YFilter { return temp.YFilter }

func (temp *Components_Component_State_Temp) SetFilter(yf yfilter.YFilter) { temp.YFilter = yf }

func (temp *Components_Component_State_Temp) GetGoName(yname string) string {
    if yname == "temp-instant" { return "TempInstant" }
    if yname == "temp-avg" { return "TempAvg" }
    if yname == "temp-max" { return "TempMax" }
    if yname == "temp-min" { return "TempMin" }
    return ""
}

func (temp *Components_Component_State_Temp) GetSegmentPath() string {
    return "temp"
}

func (temp *Components_Component_State_Temp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (temp *Components_Component_State_Temp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (temp *Components_Component_State_Temp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["temp-instant"] = temp.TempInstant
    leafs["temp-avg"] = temp.TempAvg
    leafs["temp-max"] = temp.TempMax
    leafs["temp-min"] = temp.TempMin
    return leafs
}

func (temp *Components_Component_State_Temp) GetBundleName() string { return "cisco_ios_xe" }

func (temp *Components_Component_State_Temp) GetYangName() string { return "temp" }

func (temp *Components_Component_State_Temp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (temp *Components_Component_State_Temp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (temp *Components_Component_State_Temp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (temp *Components_Component_State_Temp) SetParent(parent types.Entity) { temp.parent = parent }

func (temp *Components_Component_State_Temp) GetParent() types.Entity { return temp.parent }

func (temp *Components_Component_State_Temp) GetParentYangName() string { return "state" }

// Components_Component_PlatformProperties
// Platform component properties
type Components_Component_PlatformProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of platform component properties. The type is slice of
    // Components_Component_PlatformProperties_PlatformProperty.
    PlatformProperty []Components_Component_PlatformProperties_PlatformProperty
}

func (platformProperties *Components_Component_PlatformProperties) GetFilter() yfilter.YFilter { return platformProperties.YFilter }

func (platformProperties *Components_Component_PlatformProperties) SetFilter(yf yfilter.YFilter) { platformProperties.YFilter = yf }

func (platformProperties *Components_Component_PlatformProperties) GetGoName(yname string) string {
    if yname == "platform-property" { return "PlatformProperty" }
    return ""
}

func (platformProperties *Components_Component_PlatformProperties) GetSegmentPath() string {
    return "platform-properties"
}

func (platformProperties *Components_Component_PlatformProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "platform-property" {
        for _, c := range platformProperties.PlatformProperty {
            if platformProperties.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Components_Component_PlatformProperties_PlatformProperty{}
        platformProperties.PlatformProperty = append(platformProperties.PlatformProperty, child)
        return &platformProperties.PlatformProperty[len(platformProperties.PlatformProperty)-1]
    }
    return nil
}

func (platformProperties *Components_Component_PlatformProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range platformProperties.PlatformProperty {
        children[platformProperties.PlatformProperty[i].GetSegmentPath()] = &platformProperties.PlatformProperty[i]
    }
    return children
}

func (platformProperties *Components_Component_PlatformProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformProperties *Components_Component_PlatformProperties) GetBundleName() string { return "cisco_ios_xe" }

func (platformProperties *Components_Component_PlatformProperties) GetYangName() string { return "platform-properties" }

func (platformProperties *Components_Component_PlatformProperties) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (platformProperties *Components_Component_PlatformProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (platformProperties *Components_Component_PlatformProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (platformProperties *Components_Component_PlatformProperties) SetParent(parent types.Entity) { platformProperties.parent = parent }

func (platformProperties *Components_Component_PlatformProperties) GetParent() types.Entity { return platformProperties.parent }

func (platformProperties *Components_Component_PlatformProperties) GetParentYangName() string { return "component" }

// Components_Component_PlatformProperties_PlatformProperty
// List of platform component properties
type Components_Component_PlatformProperties_PlatformProperty struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Property name. The type is string.
    Name interface{}

    // Indication of whether the property is user-configurable. The type is bool.
    Configurable interface{}

    // Name of the parent component. The type is string.
    ParentPlatformComponentCnameKey interface{}

    // Property value.
    Value Components_Component_PlatformProperties_PlatformProperty_Value
}

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetFilter() yfilter.YFilter { return platformProperty.YFilter }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) SetFilter(yf yfilter.YFilter) { platformProperty.YFilter = yf }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "configurable" { return "Configurable" }
    if yname == "parent-platform-component-cname-key" { return "ParentPlatformComponentCnameKey" }
    if yname == "value" { return "Value" }
    return ""
}

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetSegmentPath() string {
    return "platform-property" + "[name='" + fmt.Sprintf("%v", platformProperty.Name) + "']"
}

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "value" {
        return &platformProperty.Value
    }
    return nil
}

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["value"] = &platformProperty.Value
    return children
}

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = platformProperty.Name
    leafs["configurable"] = platformProperty.Configurable
    leafs["parent-platform-component-cname-key"] = platformProperty.ParentPlatformComponentCnameKey
    return leafs
}

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetBundleName() string { return "cisco_ios_xe" }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetYangName() string { return "platform-property" }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) SetParent(parent types.Entity) { platformProperty.parent = parent }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetParent() types.Entity { return platformProperty.parent }

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetParentYangName() string { return "platform-properties" }

// Components_Component_PlatformProperties_PlatformProperty_Value
// Property value
type Components_Component_PlatformProperties_PlatformProperty_Value struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String property value. The type is string.
    String interface{}

    // Boolean property value. The type is bool.
    Boolean interface{}

    // Integer64 property value. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Intsixfour interface{}

    // Unsigned integer64 property value. The type is interface{} with range:
    // 0..18446744073709551615.
    Uintsixfour interface{}

    // Decimal64 property value. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Decimal interface{}
}

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetFilter() yfilter.YFilter { return value.YFilter }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) SetFilter(yf yfilter.YFilter) { value.YFilter = yf }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetGoName(yname string) string {
    if yname == "string" { return "String" }
    if yname == "boolean" { return "Boolean" }
    if yname == "intsixfour" { return "Intsixfour" }
    if yname == "uintsixfour" { return "Uintsixfour" }
    if yname == "decimal" { return "Decimal" }
    return ""
}

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetSegmentPath() string {
    return "value"
}

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["string"] = value.String
    leafs["boolean"] = value.Boolean
    leafs["intsixfour"] = value.Intsixfour
    leafs["uintsixfour"] = value.Uintsixfour
    leafs["decimal"] = value.Decimal
    return leafs
}

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetBundleName() string { return "cisco_ios_xe" }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetYangName() string { return "value" }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) SetParent(parent types.Entity) { value.parent = parent }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetParent() types.Entity { return value.parent }

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetParentYangName() string { return "platform-property" }

// Components_Component_PlatformSubcomponents
// Platform subcomponents
type Components_Component_PlatformSubcomponents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of platform subcomponents. The type is slice of
    // Components_Component_PlatformSubcomponents_PlatformSubcomponent.
    PlatformSubcomponent []Components_Component_PlatformSubcomponents_PlatformSubcomponent
}

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetFilter() yfilter.YFilter { return platformSubcomponents.YFilter }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) SetFilter(yf yfilter.YFilter) { platformSubcomponents.YFilter = yf }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetGoName(yname string) string {
    if yname == "platform-subcomponent" { return "PlatformSubcomponent" }
    return ""
}

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetSegmentPath() string {
    return "platform-subcomponents"
}

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "platform-subcomponent" {
        for _, c := range platformSubcomponents.PlatformSubcomponent {
            if platformSubcomponents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Components_Component_PlatformSubcomponents_PlatformSubcomponent{}
        platformSubcomponents.PlatformSubcomponent = append(platformSubcomponents.PlatformSubcomponent, child)
        return &platformSubcomponents.PlatformSubcomponent[len(platformSubcomponents.PlatformSubcomponent)-1]
    }
    return nil
}

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range platformSubcomponents.PlatformSubcomponent {
        children[platformSubcomponents.PlatformSubcomponent[i].GetSegmentPath()] = &platformSubcomponents.PlatformSubcomponent[i]
    }
    return children
}

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetBundleName() string { return "cisco_ios_xe" }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetYangName() string { return "platform-subcomponents" }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) SetParent(parent types.Entity) { platformSubcomponents.parent = parent }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetParent() types.Entity { return platformSubcomponents.parent }

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetParentYangName() string { return "component" }

// Components_Component_PlatformSubcomponents_PlatformSubcomponent
// List of platform subcomponents
type Components_Component_PlatformSubcomponents_PlatformSubcomponent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Subcomponent name. The type is string.
    Name interface{}

    // Name of the parent component. The type is string.
    ParentPlatformComponentCnameKey interface{}
}

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetFilter() yfilter.YFilter { return platformSubcomponent.YFilter }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) SetFilter(yf yfilter.YFilter) { platformSubcomponent.YFilter = yf }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "parent-platform-component-cname-key" { return "ParentPlatformComponentCnameKey" }
    return ""
}

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetSegmentPath() string {
    return "platform-subcomponent" + "[name='" + fmt.Sprintf("%v", platformSubcomponent.Name) + "']"
}

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = platformSubcomponent.Name
    leafs["parent-platform-component-cname-key"] = platformSubcomponent.ParentPlatformComponentCnameKey
    return leafs
}

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetBundleName() string { return "cisco_ios_xe" }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetYangName() string { return "platform-subcomponent" }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) SetParent(parent types.Entity) { platformSubcomponent.parent = parent }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetParent() types.Entity { return platformSubcomponent.parent }

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetParentYangName() string { return "platform-subcomponents" }

