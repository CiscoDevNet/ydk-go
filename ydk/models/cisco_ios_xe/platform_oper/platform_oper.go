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

    PlatformCompType_comp_container PlatformCompType = "comp-container"
)

// PlatformPropValueType represents Property value type
type PlatformPropValueType string

const (
    PlatformPropValueType_property_string PlatformPropValueType = "property-string"

    PlatformPropValueType_property_boolean PlatformPropValueType = "property-boolean"

    PlatformPropValueType_property_int64 PlatformPropValueType = "property-int64"

    PlatformPropValueType_property_uint64 PlatformPropValueType = "property-uint64"

    PlatformPropValueType_property_decimal64 PlatformPropValueType = "property-decimal64"
)

// Components
// Enclosing container for the components in the system
type Components struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of components, keyed by component name. The type is slice of
    // Components_Component.
    Component []Components_Component
}

func (components *Components) GetEntityData() *types.CommonEntityData {
    components.EntityData.YFilter = components.YFilter
    components.EntityData.YangName = "components"
    components.EntityData.BundleName = "cisco_ios_xe"
    components.EntityData.ParentYangName = "Cisco-IOS-XE-platform-oper"
    components.EntityData.SegmentPath = "Cisco-IOS-XE-platform-oper:components"
    components.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    components.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    components.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    components.EntityData.Children = make(map[string]types.YChild)
    components.EntityData.Children["component"] = types.YChild{"Component", nil}
    for i := range components.Component {
        components.EntityData.Children[types.GetSegmentPath(&components.Component[i])] = types.YChild{"Component", &components.Component[i]}
    }
    components.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(components.EntityData)
}

// Components_Component
// List of components, keyed by component name
type Components_Component struct {
    EntityData types.CommonEntityData
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

func (component *Components_Component) GetEntityData() *types.CommonEntityData {
    component.EntityData.YFilter = component.YFilter
    component.EntityData.YangName = "component"
    component.EntityData.BundleName = "cisco_ios_xe"
    component.EntityData.ParentYangName = "components"
    component.EntityData.SegmentPath = "component" + "[cname='" + fmt.Sprintf("%v", component.Cname) + "']"
    component.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    component.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    component.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    component.EntityData.Children = make(map[string]types.YChild)
    component.EntityData.Children["state"] = types.YChild{"State", &component.State}
    component.EntityData.Children["platform-properties"] = types.YChild{"PlatformProperties", &component.PlatformProperties}
    component.EntityData.Children["platform-subcomponents"] = types.YChild{"PlatformSubcomponents", &component.PlatformSubcomponents}
    component.EntityData.Leafs = make(map[string]types.YLeaf)
    component.EntityData.Leafs["cname"] = types.YLeaf{"Cname", component.Cname}
    return &(component.EntityData)
}

// Components_Component_State
// Operational state data for each component
type Components_Component_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of component as identified by the system. The type is
    // PlatformCompType.
    Type_ interface{}

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

func (state *Components_Component_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xe"
    state.EntityData.ParentYangName = "component"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["temp"] = types.YChild{"Temp", &state.Temp}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["type"] = types.YLeaf{"Type_", state.Type_}
    state.EntityData.Leafs["id"] = types.YLeaf{"Id", state.Id}
    state.EntityData.Leafs["description"] = types.YLeaf{"Description", state.Description}
    state.EntityData.Leafs["mfg-name"] = types.YLeaf{"MfgName", state.MfgName}
    state.EntityData.Leafs["version"] = types.YLeaf{"Version", state.Version}
    state.EntityData.Leafs["serial-no"] = types.YLeaf{"SerialNo", state.SerialNo}
    state.EntityData.Leafs["part-no"] = types.YLeaf{"PartNo", state.PartNo}
    return &(state.EntityData)
}

// Components_Component_State_Temp
// Temperature in degrees Celsius of the component. Values include
// the instantaneous, average, minimum, and maximum statistics. If
// avg/min/max statistics are not supported, the target is expected
// to just supply the instant value
type Components_Component_State_Temp struct {
    EntityData types.CommonEntityData
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

func (temp *Components_Component_State_Temp) GetEntityData() *types.CommonEntityData {
    temp.EntityData.YFilter = temp.YFilter
    temp.EntityData.YangName = "temp"
    temp.EntityData.BundleName = "cisco_ios_xe"
    temp.EntityData.ParentYangName = "state"
    temp.EntityData.SegmentPath = "temp"
    temp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    temp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    temp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    temp.EntityData.Children = make(map[string]types.YChild)
    temp.EntityData.Leafs = make(map[string]types.YLeaf)
    temp.EntityData.Leafs["temp-instant"] = types.YLeaf{"TempInstant", temp.TempInstant}
    temp.EntityData.Leafs["temp-avg"] = types.YLeaf{"TempAvg", temp.TempAvg}
    temp.EntityData.Leafs["temp-max"] = types.YLeaf{"TempMax", temp.TempMax}
    temp.EntityData.Leafs["temp-min"] = types.YLeaf{"TempMin", temp.TempMin}
    return &(temp.EntityData)
}

// Components_Component_PlatformProperties
// Platform component properties
type Components_Component_PlatformProperties struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of platform component properties. The type is slice of
    // Components_Component_PlatformProperties_PlatformProperty.
    PlatformProperty []Components_Component_PlatformProperties_PlatformProperty
}

func (platformProperties *Components_Component_PlatformProperties) GetEntityData() *types.CommonEntityData {
    platformProperties.EntityData.YFilter = platformProperties.YFilter
    platformProperties.EntityData.YangName = "platform-properties"
    platformProperties.EntityData.BundleName = "cisco_ios_xe"
    platformProperties.EntityData.ParentYangName = "component"
    platformProperties.EntityData.SegmentPath = "platform-properties"
    platformProperties.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    platformProperties.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    platformProperties.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    platformProperties.EntityData.Children = make(map[string]types.YChild)
    platformProperties.EntityData.Children["platform-property"] = types.YChild{"PlatformProperty", nil}
    for i := range platformProperties.PlatformProperty {
        platformProperties.EntityData.Children[types.GetSegmentPath(&platformProperties.PlatformProperty[i])] = types.YChild{"PlatformProperty", &platformProperties.PlatformProperty[i]}
    }
    platformProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformProperties.EntityData)
}

// Components_Component_PlatformProperties_PlatformProperty
// List of platform component properties
type Components_Component_PlatformProperties_PlatformProperty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Property name. The type is string.
    Name interface{}

    // Indication of whether the property is user-configurable. The type is bool.
    Configurable interface{}

    // Property value.
    Value Components_Component_PlatformProperties_PlatformProperty_Value
}

func (platformProperty *Components_Component_PlatformProperties_PlatformProperty) GetEntityData() *types.CommonEntityData {
    platformProperty.EntityData.YFilter = platformProperty.YFilter
    platformProperty.EntityData.YangName = "platform-property"
    platformProperty.EntityData.BundleName = "cisco_ios_xe"
    platformProperty.EntityData.ParentYangName = "platform-properties"
    platformProperty.EntityData.SegmentPath = "platform-property" + "[name='" + fmt.Sprintf("%v", platformProperty.Name) + "']"
    platformProperty.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    platformProperty.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    platformProperty.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    platformProperty.EntityData.Children = make(map[string]types.YChild)
    platformProperty.EntityData.Children["value"] = types.YChild{"Value", &platformProperty.Value}
    platformProperty.EntityData.Leafs = make(map[string]types.YLeaf)
    platformProperty.EntityData.Leafs["name"] = types.YLeaf{"Name", platformProperty.Name}
    platformProperty.EntityData.Leafs["configurable"] = types.YLeaf{"Configurable", platformProperty.Configurable}
    return &(platformProperty.EntityData)
}

// Components_Component_PlatformProperties_PlatformProperty_Value
// Property value
type Components_Component_PlatformProperties_PlatformProperty_Value struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // String property value. The type is string.
    String_ interface{}

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

func (value *Components_Component_PlatformProperties_PlatformProperty_Value) GetEntityData() *types.CommonEntityData {
    value.EntityData.YFilter = value.YFilter
    value.EntityData.YangName = "value"
    value.EntityData.BundleName = "cisco_ios_xe"
    value.EntityData.ParentYangName = "platform-property"
    value.EntityData.SegmentPath = "value"
    value.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    value.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    value.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    value.EntityData.Children = make(map[string]types.YChild)
    value.EntityData.Leafs = make(map[string]types.YLeaf)
    value.EntityData.Leafs["string"] = types.YLeaf{"String_", value.String_}
    value.EntityData.Leafs["boolean"] = types.YLeaf{"Boolean", value.Boolean}
    value.EntityData.Leafs["intsixfour"] = types.YLeaf{"Intsixfour", value.Intsixfour}
    value.EntityData.Leafs["uintsixfour"] = types.YLeaf{"Uintsixfour", value.Uintsixfour}
    value.EntityData.Leafs["decimal"] = types.YLeaf{"Decimal", value.Decimal}
    return &(value.EntityData)
}

// Components_Component_PlatformSubcomponents
// Platform subcomponents
type Components_Component_PlatformSubcomponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of platform subcomponents. The type is slice of
    // Components_Component_PlatformSubcomponents_PlatformSubcomponent.
    PlatformSubcomponent []Components_Component_PlatformSubcomponents_PlatformSubcomponent
}

func (platformSubcomponents *Components_Component_PlatformSubcomponents) GetEntityData() *types.CommonEntityData {
    platformSubcomponents.EntityData.YFilter = platformSubcomponents.YFilter
    platformSubcomponents.EntityData.YangName = "platform-subcomponents"
    platformSubcomponents.EntityData.BundleName = "cisco_ios_xe"
    platformSubcomponents.EntityData.ParentYangName = "component"
    platformSubcomponents.EntityData.SegmentPath = "platform-subcomponents"
    platformSubcomponents.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    platformSubcomponents.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    platformSubcomponents.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    platformSubcomponents.EntityData.Children = make(map[string]types.YChild)
    platformSubcomponents.EntityData.Children["platform-subcomponent"] = types.YChild{"PlatformSubcomponent", nil}
    for i := range platformSubcomponents.PlatformSubcomponent {
        platformSubcomponents.EntityData.Children[types.GetSegmentPath(&platformSubcomponents.PlatformSubcomponent[i])] = types.YChild{"PlatformSubcomponent", &platformSubcomponents.PlatformSubcomponent[i]}
    }
    platformSubcomponents.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformSubcomponents.EntityData)
}

// Components_Component_PlatformSubcomponents_PlatformSubcomponent
// List of platform subcomponents
type Components_Component_PlatformSubcomponents_PlatformSubcomponent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Subcomponent name. The type is string.
    Name interface{}
}

func (platformSubcomponent *Components_Component_PlatformSubcomponents_PlatformSubcomponent) GetEntityData() *types.CommonEntityData {
    platformSubcomponent.EntityData.YFilter = platformSubcomponent.YFilter
    platformSubcomponent.EntityData.YangName = "platform-subcomponent"
    platformSubcomponent.EntityData.BundleName = "cisco_ios_xe"
    platformSubcomponent.EntityData.ParentYangName = "platform-subcomponents"
    platformSubcomponent.EntityData.SegmentPath = "platform-subcomponent" + "[name='" + fmt.Sprintf("%v", platformSubcomponent.Name) + "']"
    platformSubcomponent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    platformSubcomponent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    platformSubcomponent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    platformSubcomponent.EntityData.Children = make(map[string]types.YChild)
    platformSubcomponent.EntityData.Leafs = make(map[string]types.YLeaf)
    platformSubcomponent.EntityData.Leafs["name"] = types.YLeaf{"Name", platformSubcomponent.Name}
    return &(platformSubcomponent.EntityData)
}

