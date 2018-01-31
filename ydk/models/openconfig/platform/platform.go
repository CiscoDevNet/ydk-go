// This module defines a data model for representing a system
// component inventory, which can include hardware or software
// elements arranged in an arbitrary structure. The primary
// relationship supported by the model is containment, e.g.,
// components containing subcomponents.
// It is expected that this model reflects every field replacable
// unit on the device at a minimum (i.e., additional information
// may be supplied about non-replacable components).
// Every element in the inventory is termed a 'component' with each
// component expected to have a unique name and type, and optionally
// a unique system-assigned identifier and FRU number.  The
// uniqueness is guaranteed by the system within the device.
// Components may have properties defined by the system that are
// modeled as a list of key-value pairs. These may or may not be
// user-configurable.  The model provides a flag for the system
// to optionally indicate which properties are user configurable.
// Each component also has a list of 'subcomponents' which are
// references to other components. Appearance in a list of
// subcomponents indicates a containment relationship as described
// above.  For example, a linecard component may have a list of
// references to port components that reside on the linecard.
// This schema is generic to allow devices to express their own
// platform-specific structure.  It may be augmented by additional
// component type-specific schemas that provide a common structure
// for well-known component types.  In these cases, the system is
// expected to populate the common component schema, and may
// optionally also represent the component and its properties in the
// generic structure.
// The properties for each component may include dynamic values,
// e.g., in the 'state' part of the schema.  For example, a CPU
// component may report its utilization, temperature, or other
// physical properties.  The intent is to capture all platform-
// specific physical data in one location, including inventory
// (presence or absence of a component) and state (physical
// attributes or status).
package platform

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package platform"))
    ydk.RegisterEntity("{http://openconfig.net/yang/platform components}", reflect.TypeOf(Components{}))
    ydk.RegisterEntity("openconfig-platform:components", reflect.TypeOf(Components{}))
}

// Components
// Enclosing container for the components in the system.
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
    return "openconfig-platform:components"
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

func (components *Components) GetBundleName() string { return "openconfig" }

func (components *Components) GetYangName() string { return "components" }

func (components *Components) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (components *Components) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (components *Components) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (components *Components) SetParent(parent types.Entity) { components.parent = parent }

func (components *Components) GetParent() types.Entity { return components.parent }

func (components *Components) GetParentYangName() string { return "openconfig-platform" }

// Components_Component
// List of components, keyed by component name.
type Components_Component struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. References the component name. The type is string.
    // Refers to platform.Components_Component_Config_Name
    Name interface{}

    // Configuration data for each component.
    Config Components_Component_Config

    // Operational state data for each component.
    State Components_Component_State

    // Enclosing container .
    Properties Components_Component_Properties

    // Enclosing container for subcomponent references.
    Subcomponents Components_Component_Subcomponents

    // Top-level container .
    OpticalPort Components_Component_OpticalPort

    // Top-level container for client port transceiver data.
    Transceiver Components_Component_Transceiver

    // Enclosing container for the list of optical channels.
    OpticalChannel Components_Component_OpticalChannel
}

func (component *Components_Component) GetFilter() yfilter.YFilter { return component.YFilter }

func (component *Components_Component) SetFilter(yf yfilter.YFilter) { component.YFilter = yf }

func (component *Components_Component) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "properties" { return "Properties" }
    if yname == "subcomponents" { return "Subcomponents" }
    if yname == "openconfig-transport-line-common:optical-port" { return "OpticalPort" }
    if yname == "openconfig-platform-transceiver:transceiver" { return "Transceiver" }
    if yname == "openconfig-terminal-device:optical-channel" { return "OpticalChannel" }
    return ""
}

func (component *Components_Component) GetSegmentPath() string {
    return "component" + "[name='" + fmt.Sprintf("%v", component.Name) + "']"
}

func (component *Components_Component) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &component.Config
    }
    if childYangName == "state" {
        return &component.State
    }
    if childYangName == "properties" {
        return &component.Properties
    }
    if childYangName == "subcomponents" {
        return &component.Subcomponents
    }
    if childYangName == "openconfig-transport-line-common:optical-port" {
        return &component.OpticalPort
    }
    if childYangName == "openconfig-platform-transceiver:transceiver" {
        return &component.Transceiver
    }
    if childYangName == "openconfig-terminal-device:optical-channel" {
        return &component.OpticalChannel
    }
    return nil
}

func (component *Components_Component) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &component.Config
    children["state"] = &component.State
    children["properties"] = &component.Properties
    children["subcomponents"] = &component.Subcomponents
    children["openconfig-transport-line-common:optical-port"] = &component.OpticalPort
    children["openconfig-platform-transceiver:transceiver"] = &component.Transceiver
    children["openconfig-terminal-device:optical-channel"] = &component.OpticalChannel
    return children
}

func (component *Components_Component) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = component.Name
    return leafs
}

func (component *Components_Component) GetBundleName() string { return "openconfig" }

func (component *Components_Component) GetYangName() string { return "component" }

func (component *Components_Component) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (component *Components_Component) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (component *Components_Component) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (component *Components_Component) SetParent(parent types.Entity) { component.parent = parent }

func (component *Components_Component) GetParent() types.Entity { return component.parent }

func (component *Components_Component) GetParentYangName() string { return "components" }

// Components_Component_Config
// Configuration data for each component
type Components_Component_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name for the component -- this will not be a configurable parameter
    // on many implementations. The type is string.
    Name interface{}
}

func (config *Components_Component_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Components_Component_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Components_Component_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (config *Components_Component_Config) GetSegmentPath() string {
    return "config"
}

func (config *Components_Component_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Components_Component_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Components_Component_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    return leafs
}

func (config *Components_Component_Config) GetBundleName() string { return "openconfig" }

func (config *Components_Component_Config) GetYangName() string { return "config" }

func (config *Components_Component_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Components_Component_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Components_Component_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Components_Component_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Components_Component_Config) GetParent() types.Entity { return config.parent }

func (config *Components_Component_Config) GetParentYangName() string { return "component" }

// Components_Component_State
// Operational state data for each component
type Components_Component_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name for the component -- this will not be a configurable parameter
    // on many implementations. The type is string.
    Name interface{}

    // Type of component as identified by the system. The type is one of the
    // following types: :go:struct:`OPENCONFIGHARDWARECOMPONENT
    // <ydk/models/platform_types/OPENCONFIGHARDWARECOMPONENT>`, or
    // :go:struct:`OPENCONFIGSOFTWARECOMPONENT
    // <ydk/models/platform_types/OPENCONFIGSOFTWARECOMPONENT>`.
    Type interface{}

    // Unique identifier assigned by the system for the component. The type is
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
}

func (state *Components_Component_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "id" { return "Id" }
    if yname == "description" { return "Description" }
    if yname == "mfg-name" { return "MfgName" }
    if yname == "version" { return "Version" }
    if yname == "serial-no" { return "SerialNo" }
    if yname == "part-no" { return "PartNo" }
    return ""
}

func (state *Components_Component_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Components_Component_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Components_Component_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["type"] = state.Type
    leafs["id"] = state.Id
    leafs["description"] = state.Description
    leafs["mfg-name"] = state.MfgName
    leafs["version"] = state.Version
    leafs["serial-no"] = state.SerialNo
    leafs["part-no"] = state.PartNo
    return leafs
}

func (state *Components_Component_State) GetBundleName() string { return "openconfig" }

func (state *Components_Component_State) GetYangName() string { return "state" }

func (state *Components_Component_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Components_Component_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Components_Component_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Components_Component_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_State) GetParentYangName() string { return "component" }

// Components_Component_Properties
// Enclosing container 
type Components_Component_Properties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of system properties for the component. The type is slice of
    // Components_Component_Properties_Property.
    Property []Components_Component_Properties_Property
}

func (properties *Components_Component_Properties) GetFilter() yfilter.YFilter { return properties.YFilter }

func (properties *Components_Component_Properties) SetFilter(yf yfilter.YFilter) { properties.YFilter = yf }

func (properties *Components_Component_Properties) GetGoName(yname string) string {
    if yname == "property" { return "Property" }
    return ""
}

func (properties *Components_Component_Properties) GetSegmentPath() string {
    return "properties"
}

func (properties *Components_Component_Properties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "property" {
        for _, c := range properties.Property {
            if properties.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Components_Component_Properties_Property{}
        properties.Property = append(properties.Property, child)
        return &properties.Property[len(properties.Property)-1]
    }
    return nil
}

func (properties *Components_Component_Properties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range properties.Property {
        children[properties.Property[i].GetSegmentPath()] = &properties.Property[i]
    }
    return children
}

func (properties *Components_Component_Properties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (properties *Components_Component_Properties) GetBundleName() string { return "openconfig" }

func (properties *Components_Component_Properties) GetYangName() string { return "properties" }

func (properties *Components_Component_Properties) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (properties *Components_Component_Properties) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (properties *Components_Component_Properties) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (properties *Components_Component_Properties) SetParent(parent types.Entity) { properties.parent = parent }

func (properties *Components_Component_Properties) GetParent() types.Entity { return properties.parent }

func (properties *Components_Component_Properties) GetParentYangName() string { return "component" }

// Components_Component_Properties_Property
// List of system properties for the component
type Components_Component_Properties_Property struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the property name. The type is
    // string. Refers to
    // platform.Components_Component_Properties_Property_Config_Name
    Name interface{}

    // Configuration data for each property.
    Config Components_Component_Properties_Property_Config

    // Operational state data for each property.
    State Components_Component_Properties_Property_State
}

func (property *Components_Component_Properties_Property) GetFilter() yfilter.YFilter { return property.YFilter }

func (property *Components_Component_Properties_Property) SetFilter(yf yfilter.YFilter) { property.YFilter = yf }

func (property *Components_Component_Properties_Property) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (property *Components_Component_Properties_Property) GetSegmentPath() string {
    return "property" + "[name='" + fmt.Sprintf("%v", property.Name) + "']"
}

func (property *Components_Component_Properties_Property) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &property.Config
    }
    if childYangName == "state" {
        return &property.State
    }
    return nil
}

func (property *Components_Component_Properties_Property) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &property.Config
    children["state"] = &property.State
    return children
}

func (property *Components_Component_Properties_Property) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = property.Name
    return leafs
}

func (property *Components_Component_Properties_Property) GetBundleName() string { return "openconfig" }

func (property *Components_Component_Properties_Property) GetYangName() string { return "property" }

func (property *Components_Component_Properties_Property) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (property *Components_Component_Properties_Property) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (property *Components_Component_Properties_Property) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (property *Components_Component_Properties_Property) SetParent(parent types.Entity) { property.parent = parent }

func (property *Components_Component_Properties_Property) GetParent() types.Entity { return property.parent }

func (property *Components_Component_Properties_Property) GetParentYangName() string { return "properties" }

// Components_Component_Properties_Property_Config
// Configuration data for each property
type Components_Component_Properties_Property_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // System-supplied name of the property -- this is typically non-configurable.
    // The type is string.
    Name interface{}

    // Property values can take on a variety of types.  Signed and unsigned
    // integer types may be provided in smaller sizes, e.g., int8, uint16, etc.
    // The type is one of the following types: string, or bool, or int with range:
    // -9223372036854775808..9223372036854775807, or int with range:
    // 0..18446744073709551615, or :go:struct:`Decimal64<Ydk_Decimal64>` with
    // range: -92233720368547758.08..92233720368547758.07.
    Value interface{}
}

func (config *Components_Component_Properties_Property_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Components_Component_Properties_Property_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Components_Component_Properties_Property_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "value" { return "Value" }
    return ""
}

func (config *Components_Component_Properties_Property_Config) GetSegmentPath() string {
    return "config"
}

func (config *Components_Component_Properties_Property_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Components_Component_Properties_Property_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Components_Component_Properties_Property_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["value"] = config.Value
    return leafs
}

func (config *Components_Component_Properties_Property_Config) GetBundleName() string { return "openconfig" }

func (config *Components_Component_Properties_Property_Config) GetYangName() string { return "config" }

func (config *Components_Component_Properties_Property_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Components_Component_Properties_Property_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Components_Component_Properties_Property_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Components_Component_Properties_Property_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Components_Component_Properties_Property_Config) GetParent() types.Entity { return config.parent }

func (config *Components_Component_Properties_Property_Config) GetParentYangName() string { return "property" }

// Components_Component_Properties_Property_State
// Operational state data for each property
type Components_Component_Properties_Property_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // System-supplied name of the property -- this is typically non-configurable.
    // The type is string.
    Name interface{}

    // Property values can take on a variety of types.  Signed and unsigned
    // integer types may be provided in smaller sizes, e.g., int8, uint16, etc.
    // The type is one of the following types: string, or bool, or int with range:
    // -9223372036854775808..9223372036854775807, or int with range:
    // 0..18446744073709551615, or :go:struct:`Decimal64<Ydk_Decimal64>` with
    // range: -92233720368547758.08..92233720368547758.07.
    Value interface{}

    // Indication whether the property is user-configurable. The type is bool.
    Configurable interface{}
}

func (state *Components_Component_Properties_Property_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_Properties_Property_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_Properties_Property_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "value" { return "Value" }
    if yname == "configurable" { return "Configurable" }
    return ""
}

func (state *Components_Component_Properties_Property_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_Properties_Property_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Components_Component_Properties_Property_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Components_Component_Properties_Property_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["value"] = state.Value
    leafs["configurable"] = state.Configurable
    return leafs
}

func (state *Components_Component_Properties_Property_State) GetBundleName() string { return "openconfig" }

func (state *Components_Component_Properties_Property_State) GetYangName() string { return "state" }

func (state *Components_Component_Properties_Property_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Components_Component_Properties_Property_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Components_Component_Properties_Property_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Components_Component_Properties_Property_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_Properties_Property_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_Properties_Property_State) GetParentYangName() string { return "property" }

// Components_Component_Subcomponents
// Enclosing container for subcomponent references
type Components_Component_Subcomponents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of subcomponent references. The type is slice of
    // Components_Component_Subcomponents_Subcomponent.
    Subcomponent []Components_Component_Subcomponents_Subcomponent
}

func (subcomponents *Components_Component_Subcomponents) GetFilter() yfilter.YFilter { return subcomponents.YFilter }

func (subcomponents *Components_Component_Subcomponents) SetFilter(yf yfilter.YFilter) { subcomponents.YFilter = yf }

func (subcomponents *Components_Component_Subcomponents) GetGoName(yname string) string {
    if yname == "subcomponent" { return "Subcomponent" }
    return ""
}

func (subcomponents *Components_Component_Subcomponents) GetSegmentPath() string {
    return "subcomponents"
}

func (subcomponents *Components_Component_Subcomponents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subcomponent" {
        for _, c := range subcomponents.Subcomponent {
            if subcomponents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Components_Component_Subcomponents_Subcomponent{}
        subcomponents.Subcomponent = append(subcomponents.Subcomponent, child)
        return &subcomponents.Subcomponent[len(subcomponents.Subcomponent)-1]
    }
    return nil
}

func (subcomponents *Components_Component_Subcomponents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subcomponents.Subcomponent {
        children[subcomponents.Subcomponent[i].GetSegmentPath()] = &subcomponents.Subcomponent[i]
    }
    return children
}

func (subcomponents *Components_Component_Subcomponents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subcomponents *Components_Component_Subcomponents) GetBundleName() string { return "openconfig" }

func (subcomponents *Components_Component_Subcomponents) GetYangName() string { return "subcomponents" }

func (subcomponents *Components_Component_Subcomponents) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subcomponents *Components_Component_Subcomponents) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subcomponents *Components_Component_Subcomponents) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subcomponents *Components_Component_Subcomponents) SetParent(parent types.Entity) { subcomponents.parent = parent }

func (subcomponents *Components_Component_Subcomponents) GetParent() types.Entity { return subcomponents.parent }

func (subcomponents *Components_Component_Subcomponents) GetParentYangName() string { return "component" }

// Components_Component_Subcomponents_Subcomponent
// List of subcomponent references
type Components_Component_Subcomponents_Subcomponent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the name list key. The type is
    // string. Refers to
    // platform.Components_Component_Subcomponents_Subcomponent_Config_Name
    Name interface{}

    // Configuration data .
    Config Components_Component_Subcomponents_Subcomponent_Config

    // Operational state data .
    State Components_Component_Subcomponents_Subcomponent_State
}

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetFilter() yfilter.YFilter { return subcomponent.YFilter }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) SetFilter(yf yfilter.YFilter) { subcomponent.YFilter = yf }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetSegmentPath() string {
    return "subcomponent" + "[name='" + fmt.Sprintf("%v", subcomponent.Name) + "']"
}

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &subcomponent.Config
    }
    if childYangName == "state" {
        return &subcomponent.State
    }
    return nil
}

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &subcomponent.Config
    children["state"] = &subcomponent.State
    return children
}

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = subcomponent.Name
    return leafs
}

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetBundleName() string { return "openconfig" }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetYangName() string { return "subcomponent" }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) SetParent(parent types.Entity) { subcomponent.parent = parent }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetParent() types.Entity { return subcomponent.parent }

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetParentYangName() string { return "subcomponents" }

// Components_Component_Subcomponents_Subcomponent_Config
// Configuration data 
type Components_Component_Subcomponents_Subcomponent_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to the name of the subcomponent. The type is string. Refers to
    // platform.Components_Component_Config_Name
    Name interface{}
}

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Components_Component_Subcomponents_Subcomponent_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetSegmentPath() string {
    return "config"
}

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    return leafs
}

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetBundleName() string { return "openconfig" }

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetYangName() string { return "config" }

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Components_Component_Subcomponents_Subcomponent_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetParent() types.Entity { return config.parent }

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetParentYangName() string { return "subcomponent" }

// Components_Component_Subcomponents_Subcomponent_State
// Operational state data 
type Components_Component_Subcomponents_Subcomponent_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to the name of the subcomponent. The type is string. Refers to
    // platform.Components_Component_Config_Name
    Name interface{}
}

func (state *Components_Component_Subcomponents_Subcomponent_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_Subcomponents_Subcomponent_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_Subcomponents_Subcomponent_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (state *Components_Component_Subcomponents_Subcomponent_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_Subcomponents_Subcomponent_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Components_Component_Subcomponents_Subcomponent_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Components_Component_Subcomponents_Subcomponent_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    return leafs
}

func (state *Components_Component_Subcomponents_Subcomponent_State) GetBundleName() string { return "openconfig" }

func (state *Components_Component_Subcomponents_Subcomponent_State) GetYangName() string { return "state" }

func (state *Components_Component_Subcomponents_Subcomponent_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Components_Component_Subcomponents_Subcomponent_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Components_Component_Subcomponents_Subcomponent_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Components_Component_Subcomponents_Subcomponent_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_Subcomponents_Subcomponent_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_Subcomponents_Subcomponent_State) GetParentYangName() string { return "subcomponent" }

// Components_Component_OpticalPort
// Top-level container 
type Components_Component_OpticalPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational config data for optical line ports.
    Config Components_Component_OpticalPort_Config

    // Operational state data for optical line ports.
    State Components_Component_OpticalPort_State
}

func (opticalPort *Components_Component_OpticalPort) GetFilter() yfilter.YFilter { return opticalPort.YFilter }

func (opticalPort *Components_Component_OpticalPort) SetFilter(yf yfilter.YFilter) { opticalPort.YFilter = yf }

func (opticalPort *Components_Component_OpticalPort) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (opticalPort *Components_Component_OpticalPort) GetSegmentPath() string {
    return "openconfig-transport-line-common:optical-port"
}

func (opticalPort *Components_Component_OpticalPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &opticalPort.Config
    }
    if childYangName == "state" {
        return &opticalPort.State
    }
    return nil
}

func (opticalPort *Components_Component_OpticalPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &opticalPort.Config
    children["state"] = &opticalPort.State
    return children
}

func (opticalPort *Components_Component_OpticalPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalPort *Components_Component_OpticalPort) GetBundleName() string { return "openconfig" }

func (opticalPort *Components_Component_OpticalPort) GetYangName() string { return "optical-port" }

func (opticalPort *Components_Component_OpticalPort) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalPort *Components_Component_OpticalPort) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalPort *Components_Component_OpticalPort) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalPort *Components_Component_OpticalPort) SetParent(parent types.Entity) { opticalPort.parent = parent }

func (opticalPort *Components_Component_OpticalPort) GetParent() types.Entity { return opticalPort.parent }

func (opticalPort *Components_Component_OpticalPort) GetParentYangName() string { return "component" }

// Components_Component_OpticalPort_Config
// Operational config data for optical line ports
type Components_Component_OpticalPort_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sets the admin state of the optical-port. The type is AdminStateType.
    AdminState interface{}
}

func (config *Components_Component_OpticalPort_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Components_Component_OpticalPort_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Components_Component_OpticalPort_Config) GetGoName(yname string) string {
    if yname == "admin-state" { return "AdminState" }
    return ""
}

func (config *Components_Component_OpticalPort_Config) GetSegmentPath() string {
    return "config"
}

func (config *Components_Component_OpticalPort_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Components_Component_OpticalPort_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Components_Component_OpticalPort_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-state"] = config.AdminState
    return leafs
}

func (config *Components_Component_OpticalPort_Config) GetBundleName() string { return "openconfig" }

func (config *Components_Component_OpticalPort_Config) GetYangName() string { return "config" }

func (config *Components_Component_OpticalPort_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Components_Component_OpticalPort_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Components_Component_OpticalPort_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Components_Component_OpticalPort_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Components_Component_OpticalPort_Config) GetParent() types.Entity { return config.parent }

func (config *Components_Component_OpticalPort_Config) GetParentYangName() string { return "optical-port" }

// Components_Component_OpticalPort_State
// Operational state data for optical line ports
type Components_Component_OpticalPort_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sets the admin state of the optical-port. The type is AdminStateType.
    AdminState interface{}

    // Indicates the type of transport line port.  This is an informational field
    // that should be made available by the device (e.g., in the
    // openconfig-platform model). The type is one of the following:
    // INGRESSMONITORDROPADDEGRESS.
    OpticalPortType interface{}

    // The total input optical power of this port in units of 0.01dBm. If
    // avg/min/max statistics are not supported, just supply the instant value.
    InputPower Components_Component_OpticalPort_State_InputPower

    // The total output optical power of this port in units of 0.01dBm. If
    // avg/min/max statistics are not supported, just supply the instant value.
    OutputPower Components_Component_OpticalPort_State_OutputPower
}

func (state *Components_Component_OpticalPort_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_OpticalPort_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_OpticalPort_State) GetGoName(yname string) string {
    if yname == "admin-state" { return "AdminState" }
    if yname == "optical-port-type" { return "OpticalPortType" }
    if yname == "input-power" { return "InputPower" }
    if yname == "output-power" { return "OutputPower" }
    return ""
}

func (state *Components_Component_OpticalPort_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_OpticalPort_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input-power" {
        return &state.InputPower
    }
    if childYangName == "output-power" {
        return &state.OutputPower
    }
    return nil
}

func (state *Components_Component_OpticalPort_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input-power"] = &state.InputPower
    children["output-power"] = &state.OutputPower
    return children
}

func (state *Components_Component_OpticalPort_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-state"] = state.AdminState
    leafs["optical-port-type"] = state.OpticalPortType
    return leafs
}

func (state *Components_Component_OpticalPort_State) GetBundleName() string { return "openconfig" }

func (state *Components_Component_OpticalPort_State) GetYangName() string { return "state" }

func (state *Components_Component_OpticalPort_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Components_Component_OpticalPort_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Components_Component_OpticalPort_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Components_Component_OpticalPort_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_OpticalPort_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_OpticalPort_State) GetParentYangName() string { return "optical-port" }

// Components_Component_OpticalPort_State_InputPower
// The total input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type Components_Component_OpticalPort_State_InputPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -92233720368547758.08..92233720368547758.07.
    // Units are dBm.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dBm.
    Min interface{}

    // The maximum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dBm.
    Max interface{}
}

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetFilter() yfilter.YFilter { return inputPower.YFilter }

func (inputPower *Components_Component_OpticalPort_State_InputPower) SetFilter(yf yfilter.YFilter) { inputPower.YFilter = yf }

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetSegmentPath() string {
    return "input-power"
}

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = inputPower.Instant
    leafs["avg"] = inputPower.Avg
    leafs["min"] = inputPower.Min
    leafs["max"] = inputPower.Max
    return leafs
}

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetBundleName() string { return "openconfig" }

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetYangName() string { return "input-power" }

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inputPower *Components_Component_OpticalPort_State_InputPower) SetParent(parent types.Entity) { inputPower.parent = parent }

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetParent() types.Entity { return inputPower.parent }

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetParentYangName() string { return "state" }

// Components_Component_OpticalPort_State_OutputPower
// The total output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type Components_Component_OpticalPort_State_OutputPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -92233720368547758.08..92233720368547758.07.
    // Units are dBm.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dBm.
    Min interface{}

    // The maximum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dBm.
    Max interface{}
}

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetFilter() yfilter.YFilter { return outputPower.YFilter }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) SetFilter(yf yfilter.YFilter) { outputPower.YFilter = yf }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetSegmentPath() string {
    return "output-power"
}

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = outputPower.Instant
    leafs["avg"] = outputPower.Avg
    leafs["min"] = outputPower.Min
    leafs["max"] = outputPower.Max
    return leafs
}

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetBundleName() string { return "openconfig" }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetYangName() string { return "output-power" }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) SetParent(parent types.Entity) { outputPower.parent = parent }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetParent() types.Entity { return outputPower.parent }

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetParentYangName() string { return "state" }

// Components_Component_Transceiver
// Top-level container for client port transceiver data
type Components_Component_Transceiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for client port transceivers.
    Config Components_Component_Transceiver_Config

    // Operational state data for client port transceivers.
    State Components_Component_Transceiver_State

    // Enclosing container for client channels.
    PhysicalChannels Components_Component_Transceiver_PhysicalChannels
}

func (transceiver *Components_Component_Transceiver) GetFilter() yfilter.YFilter { return transceiver.YFilter }

func (transceiver *Components_Component_Transceiver) SetFilter(yf yfilter.YFilter) { transceiver.YFilter = yf }

func (transceiver *Components_Component_Transceiver) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "physical-channels" { return "PhysicalChannels" }
    return ""
}

func (transceiver *Components_Component_Transceiver) GetSegmentPath() string {
    return "openconfig-platform-transceiver:transceiver"
}

func (transceiver *Components_Component_Transceiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &transceiver.Config
    }
    if childYangName == "state" {
        return &transceiver.State
    }
    if childYangName == "physical-channels" {
        return &transceiver.PhysicalChannels
    }
    return nil
}

func (transceiver *Components_Component_Transceiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &transceiver.Config
    children["state"] = &transceiver.State
    children["physical-channels"] = &transceiver.PhysicalChannels
    return children
}

func (transceiver *Components_Component_Transceiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transceiver *Components_Component_Transceiver) GetBundleName() string { return "openconfig" }

func (transceiver *Components_Component_Transceiver) GetYangName() string { return "transceiver" }

func (transceiver *Components_Component_Transceiver) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (transceiver *Components_Component_Transceiver) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (transceiver *Components_Component_Transceiver) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (transceiver *Components_Component_Transceiver) SetParent(parent types.Entity) { transceiver.parent = parent }

func (transceiver *Components_Component_Transceiver) GetParent() types.Entity { return transceiver.parent }

func (transceiver *Components_Component_Transceiver) GetParentYangName() string { return "component" }

// Components_Component_Transceiver_Config
// Configuration data for client port transceivers
type Components_Component_Transceiver_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Turns power on / off to the transceiver -- provides a means to power on/off
    // the transceiver (in the case of SFP, SFP+, QSFP,...) or enable high-power
    // mode (in the case of CFP, CFP2, CFP4) and is optionally supported (device
    // can choose to always enable).  True = power on / high power, False =
    // powered off. The type is bool.
    Enabled interface{}

    // Indicates the type of optical transceiver used on this port.  If the client
    // port is built into the device and not plugable, then non-pluggable is the
    // corresponding state. If a device port supports multiple form factors (e.g.
    // QSFP28 and QSFP+, then the value of the transceiver installed shall be
    // reported. If no transceiver is present, then the value of the highest rate
    // form factor shall be reported (QSFP28, for example).  The form factor is
    // included in configuration data to allow pre-configuring a device with the
    // expected type of transceiver ahead of deployment.  The corresponding state
    // leaf should reflect the actual transceiver type plugged into the system.
    // The type is one of the following:
    // CFP2QSFP28CFP4CFP2ACOX2XFPSFPPLUSNONPLUGGABLEOTHERQSFPSFPCFP.
    FormFactor interface{}
}

func (config *Components_Component_Transceiver_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Components_Component_Transceiver_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Components_Component_Transceiver_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "form-factor" { return "FormFactor" }
    return ""
}

func (config *Components_Component_Transceiver_Config) GetSegmentPath() string {
    return "config"
}

func (config *Components_Component_Transceiver_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Components_Component_Transceiver_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Components_Component_Transceiver_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["form-factor"] = config.FormFactor
    return leafs
}

func (config *Components_Component_Transceiver_Config) GetBundleName() string { return "openconfig" }

func (config *Components_Component_Transceiver_Config) GetYangName() string { return "config" }

func (config *Components_Component_Transceiver_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Components_Component_Transceiver_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Components_Component_Transceiver_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Components_Component_Transceiver_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Components_Component_Transceiver_Config) GetParent() types.Entity { return config.parent }

func (config *Components_Component_Transceiver_Config) GetParentYangName() string { return "transceiver" }

// Components_Component_Transceiver_State
// Operational state data for client port transceivers
type Components_Component_Transceiver_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Turns power on / off to the transceiver -- provides a means to power on/off
    // the transceiver (in the case of SFP, SFP+, QSFP,...) or enable high-power
    // mode (in the case of CFP, CFP2, CFP4) and is optionally supported (device
    // can choose to always enable).  True = power on / high power, False =
    // powered off. The type is bool.
    Enabled interface{}

    // Indicates the type of optical transceiver used on this port.  If the client
    // port is built into the device and not plugable, then non-pluggable is the
    // corresponding state. If a device port supports multiple form factors (e.g.
    // QSFP28 and QSFP+, then the value of the transceiver installed shall be
    // reported. If no transceiver is present, then the value of the highest rate
    // form factor shall be reported (QSFP28, for example).  The form factor is
    // included in configuration data to allow pre-configuring a device with the
    // expected type of transceiver ahead of deployment.  The corresponding state
    // leaf should reflect the actual transceiver type plugged into the system.
    // The type is one of the following:
    // CFP2QSFP28CFP4CFP2ACOX2XFPSFPPLUSNONPLUGGABLEOTHERQSFPSFPCFP.
    FormFactor interface{}

    // Indicates whether a transceiver is present in the specified client port.
    // The type is Present.
    Present interface{}

    // Connector type used on this port. The type is one of the following:
    // SCCONNECTORMPOCONNECTORLCCONNECTOR.
    ConnectorType interface{}

    // Internally measured temperature in degrees Celsius. MSA valid range is
    // between -40 and +125C. Accuracy shall be better than +/- 3 degC over the
    // whole temperature range. The type is interface{} with range: -40..125.
    InternalTemp interface{}

    // Full name of transceiver vendor. 16-octet field that contains ASCII
    // characters, left-aligned and padded on the right with ASCII spaces (20h).
    // The type is string with length: 1..16.
    Vendor interface{}

    // Transceiver vendor's part number. 16-octet field that contains ASCII
    // characters, left-aligned and padded on the right with ASCII spaces (20h).
    // If part number is undefined, all 16 octets = 0h. The type is string with
    // length: 1..16.
    VendorPart interface{}

    // Transceiver vendor's revision number. 2-octet field that contains ASCII
    // characters, left-aligned and padded on the right with ASCII spaces (20h).
    // The type is string with length: 1..2.
    VendorRev interface{}

    // Ethernet PMD that the transceiver supports. The SFF/QSFP MSAs have
    // registers for this and CFP MSA has similar. The type is one of the
    // following:
    // ETH40GBASESR4ETH10GBASELRMETH4X10GBASESRETH100GAOCETH100GBASESR4ETH10GBASEZRETH100GBASEER4ETH40GBASEER4ETH100GACCETHUNDEFINEDETH40GBASELR4ETH40GBASEPSM4ETH10GBASELRETH100GBASESR10ETH4X10GBASELRETH100GBASELR4ETH100GBASECLR4ETH10GBASESRETH100GBASECWDM4ETH100GBASEPSM4ETH40GBASECR4ETH100GBASECR4ETH10GBASEER.
    EthernetComplianceCode interface{}

    // SONET/SDH application code supported by the port. The type is one of the
    // following: SONETUNDEFINEDVSR20003R3VSR20003R2VSR20003R5.
    SonetSdhComplianceCode interface{}

    // OTN application code supported by the port. The type is one of the
    // following: P1L12D2P1L12D1OTNUNDEFINEDP1S12D2.
    OtnComplianceCode interface{}

    // Transceiver serial number. 16-octet field that contains ASCII characters,
    // left-aligned and padded on the right with ASCII spaces (20h). If part
    // serial number is undefined, all 16 octets = 0h. The type is string with
    // length: 1..16.
    SerialNo interface{}

    // Representation of the transceiver date code, typically stored as YYMMDD. 
    // The time portion of the value is undefined and not intended to be read. The
    // type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    DateCode interface{}

    // Indicates if a fault condition exists in the transceiver. The type is bool.
    FaultCondition interface{}
}

func (state *Components_Component_Transceiver_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_Transceiver_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_Transceiver_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "form-factor" { return "FormFactor" }
    if yname == "present" { return "Present" }
    if yname == "connector-type" { return "ConnectorType" }
    if yname == "internal-temp" { return "InternalTemp" }
    if yname == "vendor" { return "Vendor" }
    if yname == "vendor-part" { return "VendorPart" }
    if yname == "vendor-rev" { return "VendorRev" }
    if yname == "ethernet-compliance-code" { return "EthernetComplianceCode" }
    if yname == "sonet-sdh-compliance-code" { return "SonetSdhComplianceCode" }
    if yname == "otn-compliance-code" { return "OtnComplianceCode" }
    if yname == "serial-no" { return "SerialNo" }
    if yname == "date-code" { return "DateCode" }
    if yname == "fault-condition" { return "FaultCondition" }
    return ""
}

func (state *Components_Component_Transceiver_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_Transceiver_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Components_Component_Transceiver_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Components_Component_Transceiver_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["form-factor"] = state.FormFactor
    leafs["present"] = state.Present
    leafs["connector-type"] = state.ConnectorType
    leafs["internal-temp"] = state.InternalTemp
    leafs["vendor"] = state.Vendor
    leafs["vendor-part"] = state.VendorPart
    leafs["vendor-rev"] = state.VendorRev
    leafs["ethernet-compliance-code"] = state.EthernetComplianceCode
    leafs["sonet-sdh-compliance-code"] = state.SonetSdhComplianceCode
    leafs["otn-compliance-code"] = state.OtnComplianceCode
    leafs["serial-no"] = state.SerialNo
    leafs["date-code"] = state.DateCode
    leafs["fault-condition"] = state.FaultCondition
    return leafs
}

func (state *Components_Component_Transceiver_State) GetBundleName() string { return "openconfig" }

func (state *Components_Component_Transceiver_State) GetYangName() string { return "state" }

func (state *Components_Component_Transceiver_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Components_Component_Transceiver_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Components_Component_Transceiver_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Components_Component_Transceiver_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_Transceiver_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_Transceiver_State) GetParentYangName() string { return "transceiver" }

// Components_Component_Transceiver_State_Present represents the specified client port.
type Components_Component_Transceiver_State_Present string

const (
    // Transceiver is present on the port
    Components_Component_Transceiver_State_Present_PRESENT Components_Component_Transceiver_State_Present = "PRESENT"

    // Transceiver is not present on the port
    Components_Component_Transceiver_State_Present_NOT_PRESENT Components_Component_Transceiver_State_Present = "NOT_PRESENT"
)

// Components_Component_Transceiver_PhysicalChannels
// Enclosing container for client channels
type Components_Component_Transceiver_PhysicalChannels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of client channels, keyed by index within a physical client port.  A
    // physical port with a single channel would have a single zero-indexed
    // element. The type is slice of
    // Components_Component_Transceiver_PhysicalChannels_Channel.
    Channel []Components_Component_Transceiver_PhysicalChannels_Channel
}

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetFilter() yfilter.YFilter { return physicalChannels.YFilter }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) SetFilter(yf yfilter.YFilter) { physicalChannels.YFilter = yf }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetGoName(yname string) string {
    if yname == "channel" { return "Channel" }
    return ""
}

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetSegmentPath() string {
    return "physical-channels"
}

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "channel" {
        for _, c := range physicalChannels.Channel {
            if physicalChannels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Components_Component_Transceiver_PhysicalChannels_Channel{}
        physicalChannels.Channel = append(physicalChannels.Channel, child)
        return &physicalChannels.Channel[len(physicalChannels.Channel)-1]
    }
    return nil
}

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range physicalChannels.Channel {
        children[physicalChannels.Channel[i].GetSegmentPath()] = &physicalChannels.Channel[i]
    }
    return children
}

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetBundleName() string { return "openconfig" }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetYangName() string { return "physical-channels" }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) SetParent(parent types.Entity) { physicalChannels.parent = parent }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetParent() types.Entity { return physicalChannels.parent }

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetParentYangName() string { return "transceiver" }

// Components_Component_Transceiver_PhysicalChannels_Channel
// List of client channels, keyed by index within a physical
// client port.  A physical port with a single channel would
// have a single zero-indexed element
type Components_Component_Transceiver_PhysicalChannels_Channel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the index number of the channel. The
    // type is string with range: 0..65535. Refers to
    // platform.Components_Component_Transceiver_PhysicalChannels_Channel_Config_Index
    Index interface{}

    // Configuration data for physical channels.
    Config Components_Component_Transceiver_PhysicalChannels_Channel_Config

    // Operational state data for channels.
    State Components_Component_Transceiver_PhysicalChannels_Channel_State
}

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetFilter() yfilter.YFilter { return channel.YFilter }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) SetFilter(yf yfilter.YFilter) { channel.YFilter = yf }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetSegmentPath() string {
    return "channel" + "[index='" + fmt.Sprintf("%v", channel.Index) + "']"
}

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &channel.Config
    }
    if childYangName == "state" {
        return &channel.State
    }
    return nil
}

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &channel.Config
    children["state"] = &channel.State
    return children
}

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = channel.Index
    return leafs
}

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetBundleName() string { return "openconfig" }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetYangName() string { return "channel" }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) SetParent(parent types.Entity) { channel.parent = parent }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetParent() types.Entity { return channel.parent }

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetParentYangName() string { return "physical-channels" }

// Components_Component_Transceiver_PhysicalChannels_Channel_Config
// Configuration data for physical channels
type Components_Component_Transceiver_PhysicalChannels_Channel_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Index of the physical channnel or lane within a physical client port. The
    // type is interface{} with range: 0..65535.
    Index interface{}

    // Text description for the client physical channel. The type is string.
    Description interface{}

    // Enable (true) or disable (false) the transmit label for the channel. The
    // type is bool.
    TxLaser interface{}

    // Target output optical power level of the optical channel, expressed in
    // increments of 0.01 dBm (decibel-milliwats). The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    TargetOutputPower interface{}
}

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "description" { return "Description" }
    if yname == "tx-laser" { return "TxLaser" }
    if yname == "target-output-power" { return "TargetOutputPower" }
    return ""
}

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetSegmentPath() string {
    return "config"
}

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = config.Index
    leafs["description"] = config.Description
    leafs["tx-laser"] = config.TxLaser
    leafs["target-output-power"] = config.TargetOutputPower
    return leafs
}

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetBundleName() string { return "openconfig" }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetYangName() string { return "config" }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetParent() types.Entity { return config.parent }

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetParentYangName() string { return "channel" }

// Components_Component_Transceiver_PhysicalChannels_Channel_State
// Operational state data for channels
type Components_Component_Transceiver_PhysicalChannels_Channel_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Index of the physical channnel or lane within a physical client port. The
    // type is interface{} with range: 0..65535.
    Index interface{}

    // Text description for the client physical channel. The type is string.
    Description interface{}

    // Enable (true) or disable (false) the transmit label for the channel. The
    // type is bool.
    TxLaser interface{}

    // Target output optical power level of the optical channel, expressed in
    // increments of 0.01 dBm (decibel-milliwats). The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    TargetOutputPower interface{}

    // The frequency in MHz of the individual physical channel (e.g. ITU C50 -
    // 195.0THz and would be reported as 195,000,000 MHz in this model). This
    // attribute is not configurable on most client ports. The type is interface{}
    // with range: 0..18446744073709551615.
    OutputFrequency interface{}

    // The output optical power of this port in units of 0.01dBm. If the port is
    // an aggregate of multiple physical channels, this attribute is the total
    // power or sum of all channels. If avg/min/max statistics are not supported,
    // the target is expected to just supply the instant value.
    OutputPower Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower

    // The input optical power of this port in units of 0.01dBm. If the port is an
    // aggregate of multiple physical channels, this attribute is the total power
    // or sum of all channels. If avg/min/max statistics are not supported, the
    // target is expected to just supply the instant value.
    InputPower Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower

    // The current applied by the system to the transmit laser to achieve the
    // output power.  The current is expressed in mA with up to one decimal
    // precision. If avg/min/max statistics are not supported, the target is
    // expected to just supply the instant value.
    LaserBiasCurrent Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent
}

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "description" { return "Description" }
    if yname == "tx-laser" { return "TxLaser" }
    if yname == "target-output-power" { return "TargetOutputPower" }
    if yname == "output-frequency" { return "OutputFrequency" }
    if yname == "output-power" { return "OutputPower" }
    if yname == "input-power" { return "InputPower" }
    if yname == "laser-bias-current" { return "LaserBiasCurrent" }
    return ""
}

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output-power" {
        return &state.OutputPower
    }
    if childYangName == "input-power" {
        return &state.InputPower
    }
    if childYangName == "laser-bias-current" {
        return &state.LaserBiasCurrent
    }
    return nil
}

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output-power"] = &state.OutputPower
    children["input-power"] = &state.InputPower
    children["laser-bias-current"] = &state.LaserBiasCurrent
    return children
}

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = state.Index
    leafs["description"] = state.Description
    leafs["tx-laser"] = state.TxLaser
    leafs["target-output-power"] = state.TargetOutputPower
    leafs["output-frequency"] = state.OutputFrequency
    return leafs
}

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetBundleName() string { return "openconfig" }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetYangName() string { return "state" }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetParentYangName() string { return "channel" }

// Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower
// The output optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels. If
// avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetFilter() yfilter.YFilter { return outputPower.YFilter }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) SetFilter(yf yfilter.YFilter) { outputPower.YFilter = yf }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetSegmentPath() string {
    return "output-power"
}

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = outputPower.Instant
    leafs["avg"] = outputPower.Avg
    leafs["min"] = outputPower.Min
    leafs["max"] = outputPower.Max
    return leafs
}

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetBundleName() string { return "openconfig" }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetYangName() string { return "output-power" }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) SetParent(parent types.Entity) { outputPower.parent = parent }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetParent() types.Entity { return outputPower.parent }

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetParentYangName() string { return "state" }

// Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower
// The input optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels.
// If avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetFilter() yfilter.YFilter { return inputPower.YFilter }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) SetFilter(yf yfilter.YFilter) { inputPower.YFilter = yf }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetSegmentPath() string {
    return "input-power"
}

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = inputPower.Instant
    leafs["avg"] = inputPower.Avg
    leafs["min"] = inputPower.Min
    leafs["max"] = inputPower.Max
    return leafs
}

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetBundleName() string { return "openconfig" }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetYangName() string { return "input-power" }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) SetParent(parent types.Entity) { inputPower.parent = parent }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetParent() types.Entity { return inputPower.parent }

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetParentYangName() string { return "state" }

// Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power.  The current is expressed in mA
// with up to one decimal precision. If avg/min/max statistics
// are not supported, the target is expected to just supply
// the instant value
type Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetFilter() yfilter.YFilter { return laserBiasCurrent.YFilter }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) SetFilter(yf yfilter.YFilter) { laserBiasCurrent.YFilter = yf }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetSegmentPath() string {
    return "laser-bias-current"
}

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = laserBiasCurrent.Instant
    leafs["avg"] = laserBiasCurrent.Avg
    leafs["min"] = laserBiasCurrent.Min
    leafs["max"] = laserBiasCurrent.Max
    return leafs
}

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetBundleName() string { return "openconfig" }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetYangName() string { return "laser-bias-current" }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) SetParent(parent types.Entity) { laserBiasCurrent.parent = parent }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetParent() types.Entity { return laserBiasCurrent.parent }

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetParentYangName() string { return "state" }

// Components_Component_OpticalChannel
// Enclosing container for the list of optical channels
type Components_Component_OpticalChannel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for optical channels.
    Config Components_Component_OpticalChannel_Config

    // Operational state data for optical channels.
    State Components_Component_OpticalChannel_State
}

func (opticalChannel *Components_Component_OpticalChannel) GetFilter() yfilter.YFilter { return opticalChannel.YFilter }

func (opticalChannel *Components_Component_OpticalChannel) SetFilter(yf yfilter.YFilter) { opticalChannel.YFilter = yf }

func (opticalChannel *Components_Component_OpticalChannel) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (opticalChannel *Components_Component_OpticalChannel) GetSegmentPath() string {
    return "openconfig-terminal-device:optical-channel"
}

func (opticalChannel *Components_Component_OpticalChannel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &opticalChannel.Config
    }
    if childYangName == "state" {
        return &opticalChannel.State
    }
    return nil
}

func (opticalChannel *Components_Component_OpticalChannel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &opticalChannel.Config
    children["state"] = &opticalChannel.State
    return children
}

func (opticalChannel *Components_Component_OpticalChannel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalChannel *Components_Component_OpticalChannel) GetBundleName() string { return "openconfig" }

func (opticalChannel *Components_Component_OpticalChannel) GetYangName() string { return "optical-channel" }

func (opticalChannel *Components_Component_OpticalChannel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalChannel *Components_Component_OpticalChannel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalChannel *Components_Component_OpticalChannel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalChannel *Components_Component_OpticalChannel) SetParent(parent types.Entity) { opticalChannel.parent = parent }

func (opticalChannel *Components_Component_OpticalChannel) GetParent() types.Entity { return opticalChannel.parent }

func (opticalChannel *Components_Component_OpticalChannel) GetParentYangName() string { return "component" }

// Components_Component_OpticalChannel_Config
// Configuration data for optical channels
type Components_Component_OpticalChannel_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Frequency of the optical channel, expressed in MHz. The type is interface{}
    // with range: 0..18446744073709551615.
    Frequency interface{}

    // Target output optical power level of the optical channel, expressed in
    // increments of 0.01 dBm (decibel-milliwats). The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    TargetOutputPower interface{}

    // Vendor-specific mode identifier -- sets the operational mode for the
    // channel. The type is interface{} with range: 0..65535.
    OperationalMode interface{}

    // Reference to the line-side physical port that carries this optical channel.
    // The target port should be a component in the physical inventory data model.
    // The type is string. Refers to platform.Components_Component_Name
    LinePort interface{}
}

func (config *Components_Component_OpticalChannel_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Components_Component_OpticalChannel_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Components_Component_OpticalChannel_Config) GetGoName(yname string) string {
    if yname == "frequency" { return "Frequency" }
    if yname == "target-output-power" { return "TargetOutputPower" }
    if yname == "operational-mode" { return "OperationalMode" }
    if yname == "line-port" { return "LinePort" }
    return ""
}

func (config *Components_Component_OpticalChannel_Config) GetSegmentPath() string {
    return "config"
}

func (config *Components_Component_OpticalChannel_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Components_Component_OpticalChannel_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Components_Component_OpticalChannel_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frequency"] = config.Frequency
    leafs["target-output-power"] = config.TargetOutputPower
    leafs["operational-mode"] = config.OperationalMode
    leafs["line-port"] = config.LinePort
    return leafs
}

func (config *Components_Component_OpticalChannel_Config) GetBundleName() string { return "openconfig" }

func (config *Components_Component_OpticalChannel_Config) GetYangName() string { return "config" }

func (config *Components_Component_OpticalChannel_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Components_Component_OpticalChannel_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Components_Component_OpticalChannel_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Components_Component_OpticalChannel_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Components_Component_OpticalChannel_Config) GetParent() types.Entity { return config.parent }

func (config *Components_Component_OpticalChannel_Config) GetParentYangName() string { return "optical-channel" }

// Components_Component_OpticalChannel_State
// Operational state data for optical channels
type Components_Component_OpticalChannel_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Frequency of the optical channel, expressed in MHz. The type is interface{}
    // with range: 0..18446744073709551615.
    Frequency interface{}

    // Target output optical power level of the optical channel, expressed in
    // increments of 0.01 dBm (decibel-milliwats). The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    TargetOutputPower interface{}

    // Vendor-specific mode identifier -- sets the operational mode for the
    // channel. The type is interface{} with range: 0..65535.
    OperationalMode interface{}

    // Reference to the line-side physical port that carries this optical channel.
    // The target port should be a component in the physical inventory data model.
    // The type is string. Refers to platform.Components_Component_Name
    LinePort interface{}

    // If the device places constraints on which optical channels must be managed
    // together (e.g., transmitted on the same line port), it can indicate that by
    // setting the group-id to the same value across related optical channels. The
    // type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // The output optical power of this port in units of 0.01dBm. If the port is
    // an aggregate of multiple physical channels, this attribute is the total
    // power or sum of all channels. If avg/min/max statistics are not supported,
    // the target is expected to just supply the instant value.
    OutputPower Components_Component_OpticalChannel_State_OutputPower

    // The input optical power of this port in units of 0.01dBm. If the port is an
    // aggregate of multiple physical channels, this attribute is the total power
    // or sum of all channels. If avg/min/max statistics are not supported, the
    // target is expected to just supply the instant value.
    InputPower Components_Component_OpticalChannel_State_InputPower

    // The current applied by the system to the transmit laser to achieve the
    // output power.  The current is expressed in mA with up to one decimal
    // precision. If avg/min/max statistics are not supported, the target is
    // expected to just supply the instant value.
    LaserBiasCurrent Components_Component_OpticalChannel_State_LaserBiasCurrent

    // Chromatic Dispersion of an optical channel in ps/nm as reported by
    // receiver.
    ChromaticDispersion Components_Component_OpticalChannel_State_ChromaticDispersion

    // Polarization Mode Dispersion of an optical channel in ps as reported by
    // receiver.
    PolarizationModeDispersion Components_Component_OpticalChannel_State_PolarizationModeDispersion

    // Second Order Polarization Mode Dispersion of an optical channel in ps^2 as
    // reported by receiver.
    SecondOrderPolarizationModeDispersion Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion

    // Polarization Dependent Loss of an optical channel in dB as reported by
    // receiver.
    PolarizationDependentLoss Components_Component_OpticalChannel_State_PolarizationDependentLoss
}

func (state *Components_Component_OpticalChannel_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Components_Component_OpticalChannel_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Components_Component_OpticalChannel_State) GetGoName(yname string) string {
    if yname == "frequency" { return "Frequency" }
    if yname == "target-output-power" { return "TargetOutputPower" }
    if yname == "operational-mode" { return "OperationalMode" }
    if yname == "line-port" { return "LinePort" }
    if yname == "group-id" { return "GroupId" }
    if yname == "output-power" { return "OutputPower" }
    if yname == "input-power" { return "InputPower" }
    if yname == "laser-bias-current" { return "LaserBiasCurrent" }
    if yname == "chromatic-dispersion" { return "ChromaticDispersion" }
    if yname == "polarization-mode-dispersion" { return "PolarizationModeDispersion" }
    if yname == "second-order-polarization-mode-dispersion" { return "SecondOrderPolarizationModeDispersion" }
    if yname == "polarization-dependent-loss" { return "PolarizationDependentLoss" }
    return ""
}

func (state *Components_Component_OpticalChannel_State) GetSegmentPath() string {
    return "state"
}

func (state *Components_Component_OpticalChannel_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output-power" {
        return &state.OutputPower
    }
    if childYangName == "input-power" {
        return &state.InputPower
    }
    if childYangName == "laser-bias-current" {
        return &state.LaserBiasCurrent
    }
    if childYangName == "chromatic-dispersion" {
        return &state.ChromaticDispersion
    }
    if childYangName == "polarization-mode-dispersion" {
        return &state.PolarizationModeDispersion
    }
    if childYangName == "second-order-polarization-mode-dispersion" {
        return &state.SecondOrderPolarizationModeDispersion
    }
    if childYangName == "polarization-dependent-loss" {
        return &state.PolarizationDependentLoss
    }
    return nil
}

func (state *Components_Component_OpticalChannel_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output-power"] = &state.OutputPower
    children["input-power"] = &state.InputPower
    children["laser-bias-current"] = &state.LaserBiasCurrent
    children["chromatic-dispersion"] = &state.ChromaticDispersion
    children["polarization-mode-dispersion"] = &state.PolarizationModeDispersion
    children["second-order-polarization-mode-dispersion"] = &state.SecondOrderPolarizationModeDispersion
    children["polarization-dependent-loss"] = &state.PolarizationDependentLoss
    return children
}

func (state *Components_Component_OpticalChannel_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frequency"] = state.Frequency
    leafs["target-output-power"] = state.TargetOutputPower
    leafs["operational-mode"] = state.OperationalMode
    leafs["line-port"] = state.LinePort
    leafs["group-id"] = state.GroupId
    return leafs
}

func (state *Components_Component_OpticalChannel_State) GetBundleName() string { return "openconfig" }

func (state *Components_Component_OpticalChannel_State) GetYangName() string { return "state" }

func (state *Components_Component_OpticalChannel_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Components_Component_OpticalChannel_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Components_Component_OpticalChannel_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Components_Component_OpticalChannel_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Components_Component_OpticalChannel_State) GetParent() types.Entity { return state.parent }

func (state *Components_Component_OpticalChannel_State) GetParentYangName() string { return "optical-channel" }

// Components_Component_OpticalChannel_State_OutputPower
// The output optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels. If
// avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_OpticalChannel_State_OutputPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetFilter() yfilter.YFilter { return outputPower.YFilter }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) SetFilter(yf yfilter.YFilter) { outputPower.YFilter = yf }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetSegmentPath() string {
    return "output-power"
}

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = outputPower.Instant
    leafs["avg"] = outputPower.Avg
    leafs["min"] = outputPower.Min
    leafs["max"] = outputPower.Max
    return leafs
}

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetBundleName() string { return "openconfig" }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetYangName() string { return "output-power" }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) SetParent(parent types.Entity) { outputPower.parent = parent }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetParent() types.Entity { return outputPower.parent }

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetParentYangName() string { return "state" }

// Components_Component_OpticalChannel_State_InputPower
// The input optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels.
// If avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_OpticalChannel_State_InputPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetFilter() yfilter.YFilter { return inputPower.YFilter }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) SetFilter(yf yfilter.YFilter) { inputPower.YFilter = yf }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetSegmentPath() string {
    return "input-power"
}

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = inputPower.Instant
    leafs["avg"] = inputPower.Avg
    leafs["min"] = inputPower.Min
    leafs["max"] = inputPower.Max
    return leafs
}

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetBundleName() string { return "openconfig" }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetYangName() string { return "input-power" }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) SetParent(parent types.Entity) { inputPower.parent = parent }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetParent() types.Entity { return inputPower.parent }

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetParentYangName() string { return "state" }

// Components_Component_OpticalChannel_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power.  The current is expressed in mA
// with up to one decimal precision. If avg/min/max statistics
// are not supported, the target is expected to just supply
// the instant value
type Components_Component_OpticalChannel_State_LaserBiasCurrent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetFilter() yfilter.YFilter { return laserBiasCurrent.YFilter }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) SetFilter(yf yfilter.YFilter) { laserBiasCurrent.YFilter = yf }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetSegmentPath() string {
    return "laser-bias-current"
}

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = laserBiasCurrent.Instant
    leafs["avg"] = laserBiasCurrent.Avg
    leafs["min"] = laserBiasCurrent.Min
    leafs["max"] = laserBiasCurrent.Max
    return leafs
}

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetBundleName() string { return "openconfig" }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetYangName() string { return "laser-bias-current" }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) SetParent(parent types.Entity) { laserBiasCurrent.parent = parent }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetParent() types.Entity { return laserBiasCurrent.parent }

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetParentYangName() string { return "state" }

// Components_Component_OpticalChannel_State_ChromaticDispersion
// Chromatic Dispersion of an optical channel
// in ps/nm as reported by receiver
type Components_Component_OpticalChannel_State_ChromaticDispersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetFilter() yfilter.YFilter { return chromaticDispersion.YFilter }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) SetFilter(yf yfilter.YFilter) { chromaticDispersion.YFilter = yf }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetSegmentPath() string {
    return "chromatic-dispersion"
}

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = chromaticDispersion.Instant
    leafs["avg"] = chromaticDispersion.Avg
    leafs["min"] = chromaticDispersion.Min
    leafs["max"] = chromaticDispersion.Max
    return leafs
}

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetBundleName() string { return "openconfig" }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetYangName() string { return "chromatic-dispersion" }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) SetParent(parent types.Entity) { chromaticDispersion.parent = parent }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetParent() types.Entity { return chromaticDispersion.parent }

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetParentYangName() string { return "state" }

// Components_Component_OpticalChannel_State_PolarizationModeDispersion
// Polarization Mode Dispersion of an optical channel
// in ps as reported by receiver
type Components_Component_OpticalChannel_State_PolarizationModeDispersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetFilter() yfilter.YFilter { return polarizationModeDispersion.YFilter }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) SetFilter(yf yfilter.YFilter) { polarizationModeDispersion.YFilter = yf }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetSegmentPath() string {
    return "polarization-mode-dispersion"
}

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = polarizationModeDispersion.Instant
    leafs["avg"] = polarizationModeDispersion.Avg
    leafs["min"] = polarizationModeDispersion.Min
    leafs["max"] = polarizationModeDispersion.Max
    return leafs
}

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetBundleName() string { return "openconfig" }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetYangName() string { return "polarization-mode-dispersion" }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) SetParent(parent types.Entity) { polarizationModeDispersion.parent = parent }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetParent() types.Entity { return polarizationModeDispersion.parent }

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetParentYangName() string { return "state" }

// Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion
// Second Order Polarization Mode Dispersion of an optical
// channel in ps^2 as reported by receiver
type Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetFilter() yfilter.YFilter { return secondOrderPolarizationModeDispersion.YFilter }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) SetFilter(yf yfilter.YFilter) { secondOrderPolarizationModeDispersion.YFilter = yf }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetSegmentPath() string {
    return "second-order-polarization-mode-dispersion"
}

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = secondOrderPolarizationModeDispersion.Instant
    leafs["avg"] = secondOrderPolarizationModeDispersion.Avg
    leafs["min"] = secondOrderPolarizationModeDispersion.Min
    leafs["max"] = secondOrderPolarizationModeDispersion.Max
    return leafs
}

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetBundleName() string { return "openconfig" }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetYangName() string { return "second-order-polarization-mode-dispersion" }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) SetParent(parent types.Entity) { secondOrderPolarizationModeDispersion.parent = parent }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetParent() types.Entity { return secondOrderPolarizationModeDispersion.parent }

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetParentYangName() string { return "state" }

// Components_Component_OpticalChannel_State_PolarizationDependentLoss
// Polarization Dependent Loss of an optical channel
// in dB as reported by receiver
type Components_Component_OpticalChannel_State_PolarizationDependentLoss struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -922337203685477580.8..922337203685477580.7.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -922337203685477580.8..922337203685477580.7.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Min interface{}

    // The maximum value of the statitic over the sampling period. The type is
    // string with range: -922337203685477580.8..922337203685477580.7.
    Max interface{}
}

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetFilter() yfilter.YFilter { return polarizationDependentLoss.YFilter }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) SetFilter(yf yfilter.YFilter) { polarizationDependentLoss.YFilter = yf }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetSegmentPath() string {
    return "polarization-dependent-loss"
}

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = polarizationDependentLoss.Instant
    leafs["avg"] = polarizationDependentLoss.Avg
    leafs["min"] = polarizationDependentLoss.Min
    leafs["max"] = polarizationDependentLoss.Max
    return leafs
}

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetBundleName() string { return "openconfig" }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetYangName() string { return "polarization-dependent-loss" }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) SetParent(parent types.Entity) { polarizationDependentLoss.parent = parent }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetParent() types.Entity { return polarizationDependentLoss.parent }

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetParentYangName() string { return "state" }

