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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of components, keyed by component name. The type is slice of
    // Components_Component.
    Component []*Components_Component
}

func (components *Components) GetEntityData() *types.CommonEntityData {
    components.EntityData.YFilter = components.YFilter
    components.EntityData.YangName = "components"
    components.EntityData.BundleName = "openconfig"
    components.EntityData.ParentYangName = "openconfig-platform"
    components.EntityData.SegmentPath = "openconfig-platform:components"
    components.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    components.EntityData.NamespaceTable = openconfig.GetNamespaces()
    components.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    components.EntityData.Children = types.NewOrderedMap()
    components.EntityData.Children.Append("component", types.YChild{"Component", nil})
    for i := range components.Component {
        components.EntityData.Children.Append(types.GetSegmentPath(components.Component[i]), types.YChild{"Component", components.Component[i]})
    }
    components.EntityData.Leafs = types.NewOrderedMap()

    components.EntityData.YListKeys = []string {}

    return &(components.EntityData)
}

// Components_Component
// List of components, keyed by component name.
type Components_Component struct {
    EntityData types.CommonEntityData
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

func (component *Components_Component) GetEntityData() *types.CommonEntityData {
    component.EntityData.YFilter = component.YFilter
    component.EntityData.YangName = "component"
    component.EntityData.BundleName = "openconfig"
    component.EntityData.ParentYangName = "components"
    component.EntityData.SegmentPath = "component" + types.AddKeyToken(component.Name, "name")
    component.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    component.EntityData.NamespaceTable = openconfig.GetNamespaces()
    component.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    component.EntityData.Children = types.NewOrderedMap()
    component.EntityData.Children.Append("config", types.YChild{"Config", &component.Config})
    component.EntityData.Children.Append("state", types.YChild{"State", &component.State})
    component.EntityData.Children.Append("properties", types.YChild{"Properties", &component.Properties})
    component.EntityData.Children.Append("subcomponents", types.YChild{"Subcomponents", &component.Subcomponents})
    component.EntityData.Children.Append("openconfig-transport-line-common:optical-port", types.YChild{"OpticalPort", &component.OpticalPort})
    component.EntityData.Children.Append("openconfig-platform-transceiver:transceiver", types.YChild{"Transceiver", &component.Transceiver})
    component.EntityData.Children.Append("openconfig-terminal-device:optical-channel", types.YChild{"OpticalChannel", &component.OpticalChannel})
    component.EntityData.Leafs = types.NewOrderedMap()
    component.EntityData.Leafs.Append("name", types.YLeaf{"Name", component.Name})

    component.EntityData.YListKeys = []string {"Name"}

    return &(component.EntityData)
}

// Components_Component_Config
// Configuration data for each component
type Components_Component_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name for the component -- this will not be a configurable parameter
    // on many implementations. The type is string.
    Name interface{}
}

func (config *Components_Component_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "component"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Components_Component_State
// Operational state data for each component
type Components_Component_State struct {
    EntityData types.CommonEntityData
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

func (state *Components_Component_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "component"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("id", types.YLeaf{"Id", state.Id})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("mfg-name", types.YLeaf{"MfgName", state.MfgName})
    state.EntityData.Leafs.Append("version", types.YLeaf{"Version", state.Version})
    state.EntityData.Leafs.Append("serial-no", types.YLeaf{"SerialNo", state.SerialNo})
    state.EntityData.Leafs.Append("part-no", types.YLeaf{"PartNo", state.PartNo})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Components_Component_Properties
// Enclosing container 
type Components_Component_Properties struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of system properties for the component. The type is slice of
    // Components_Component_Properties_Property.
    Property []*Components_Component_Properties_Property
}

func (properties *Components_Component_Properties) GetEntityData() *types.CommonEntityData {
    properties.EntityData.YFilter = properties.YFilter
    properties.EntityData.YangName = "properties"
    properties.EntityData.BundleName = "openconfig"
    properties.EntityData.ParentYangName = "component"
    properties.EntityData.SegmentPath = "properties"
    properties.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    properties.EntityData.NamespaceTable = openconfig.GetNamespaces()
    properties.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    properties.EntityData.Children = types.NewOrderedMap()
    properties.EntityData.Children.Append("property", types.YChild{"Property", nil})
    for i := range properties.Property {
        properties.EntityData.Children.Append(types.GetSegmentPath(properties.Property[i]), types.YChild{"Property", properties.Property[i]})
    }
    properties.EntityData.Leafs = types.NewOrderedMap()

    properties.EntityData.YListKeys = []string {}

    return &(properties.EntityData)
}

// Components_Component_Properties_Property
// List of system properties for the component
type Components_Component_Properties_Property struct {
    EntityData types.CommonEntityData
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

func (property *Components_Component_Properties_Property) GetEntityData() *types.CommonEntityData {
    property.EntityData.YFilter = property.YFilter
    property.EntityData.YangName = "property"
    property.EntityData.BundleName = "openconfig"
    property.EntityData.ParentYangName = "properties"
    property.EntityData.SegmentPath = "property" + types.AddKeyToken(property.Name, "name")
    property.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    property.EntityData.NamespaceTable = openconfig.GetNamespaces()
    property.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    property.EntityData.Children = types.NewOrderedMap()
    property.EntityData.Children.Append("config", types.YChild{"Config", &property.Config})
    property.EntityData.Children.Append("state", types.YChild{"State", &property.State})
    property.EntityData.Leafs = types.NewOrderedMap()
    property.EntityData.Leafs.Append("name", types.YLeaf{"Name", property.Name})

    property.EntityData.YListKeys = []string {"Name"}

    return &(property.EntityData)
}

// Components_Component_Properties_Property_Config
// Configuration data for each property
type Components_Component_Properties_Property_Config struct {
    EntityData types.CommonEntityData
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

func (config *Components_Component_Properties_Property_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "property"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("value", types.YLeaf{"Value", config.Value})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Components_Component_Properties_Property_State
// Operational state data for each property
type Components_Component_Properties_Property_State struct {
    EntityData types.CommonEntityData
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

func (state *Components_Component_Properties_Property_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "property"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("value", types.YLeaf{"Value", state.Value})
    state.EntityData.Leafs.Append("configurable", types.YLeaf{"Configurable", state.Configurable})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Components_Component_Subcomponents
// Enclosing container for subcomponent references
type Components_Component_Subcomponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of subcomponent references. The type is slice of
    // Components_Component_Subcomponents_Subcomponent.
    Subcomponent []*Components_Component_Subcomponents_Subcomponent
}

func (subcomponents *Components_Component_Subcomponents) GetEntityData() *types.CommonEntityData {
    subcomponents.EntityData.YFilter = subcomponents.YFilter
    subcomponents.EntityData.YangName = "subcomponents"
    subcomponents.EntityData.BundleName = "openconfig"
    subcomponents.EntityData.ParentYangName = "component"
    subcomponents.EntityData.SegmentPath = "subcomponents"
    subcomponents.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subcomponents.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subcomponents.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subcomponents.EntityData.Children = types.NewOrderedMap()
    subcomponents.EntityData.Children.Append("subcomponent", types.YChild{"Subcomponent", nil})
    for i := range subcomponents.Subcomponent {
        subcomponents.EntityData.Children.Append(types.GetSegmentPath(subcomponents.Subcomponent[i]), types.YChild{"Subcomponent", subcomponents.Subcomponent[i]})
    }
    subcomponents.EntityData.Leafs = types.NewOrderedMap()

    subcomponents.EntityData.YListKeys = []string {}

    return &(subcomponents.EntityData)
}

// Components_Component_Subcomponents_Subcomponent
// List of subcomponent references
type Components_Component_Subcomponents_Subcomponent struct {
    EntityData types.CommonEntityData
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

func (subcomponent *Components_Component_Subcomponents_Subcomponent) GetEntityData() *types.CommonEntityData {
    subcomponent.EntityData.YFilter = subcomponent.YFilter
    subcomponent.EntityData.YangName = "subcomponent"
    subcomponent.EntityData.BundleName = "openconfig"
    subcomponent.EntityData.ParentYangName = "subcomponents"
    subcomponent.EntityData.SegmentPath = "subcomponent" + types.AddKeyToken(subcomponent.Name, "name")
    subcomponent.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subcomponent.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subcomponent.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subcomponent.EntityData.Children = types.NewOrderedMap()
    subcomponent.EntityData.Children.Append("config", types.YChild{"Config", &subcomponent.Config})
    subcomponent.EntityData.Children.Append("state", types.YChild{"State", &subcomponent.State})
    subcomponent.EntityData.Leafs = types.NewOrderedMap()
    subcomponent.EntityData.Leafs.Append("name", types.YLeaf{"Name", subcomponent.Name})

    subcomponent.EntityData.YListKeys = []string {"Name"}

    return &(subcomponent.EntityData)
}

// Components_Component_Subcomponents_Subcomponent_Config
// Configuration data 
type Components_Component_Subcomponents_Subcomponent_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the name of the subcomponent. The type is string. Refers to
    // platform.Components_Component_Config_Name
    Name interface{}
}

func (config *Components_Component_Subcomponents_Subcomponent_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "subcomponent"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Components_Component_Subcomponents_Subcomponent_State
// Operational state data 
type Components_Component_Subcomponents_Subcomponent_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the name of the subcomponent. The type is string. Refers to
    // platform.Components_Component_Config_Name
    Name interface{}
}

func (state *Components_Component_Subcomponents_Subcomponent_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "subcomponent"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Components_Component_OpticalPort
// Top-level container 
type Components_Component_OpticalPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational config data for optical line ports.
    Config Components_Component_OpticalPort_Config

    // Operational state data for optical line ports.
    State Components_Component_OpticalPort_State
}

func (opticalPort *Components_Component_OpticalPort) GetEntityData() *types.CommonEntityData {
    opticalPort.EntityData.YFilter = opticalPort.YFilter
    opticalPort.EntityData.YangName = "optical-port"
    opticalPort.EntityData.BundleName = "openconfig"
    opticalPort.EntityData.ParentYangName = "component"
    opticalPort.EntityData.SegmentPath = "openconfig-transport-line-common:optical-port"
    opticalPort.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalPort.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalPort.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalPort.EntityData.Children = types.NewOrderedMap()
    opticalPort.EntityData.Children.Append("config", types.YChild{"Config", &opticalPort.Config})
    opticalPort.EntityData.Children.Append("state", types.YChild{"State", &opticalPort.State})
    opticalPort.EntityData.Leafs = types.NewOrderedMap()

    opticalPort.EntityData.YListKeys = []string {}

    return &(opticalPort.EntityData)
}

// Components_Component_OpticalPort_Config
// Operational config data for optical line ports
type Components_Component_OpticalPort_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets the admin state of the optical-port. The type is AdminStateType.
    AdminState interface{}
}

func (config *Components_Component_OpticalPort_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "optical-port"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", config.AdminState})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Components_Component_OpticalPort_State
// Operational state data for optical line ports
type Components_Component_OpticalPort_State struct {
    EntityData types.CommonEntityData
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

func (state *Components_Component_OpticalPort_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "optical-port"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("input-power", types.YChild{"InputPower", &state.InputPower})
    state.EntityData.Children.Append("output-power", types.YChild{"OutputPower", &state.OutputPower})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", state.AdminState})
    state.EntityData.Leafs.Append("optical-port-type", types.YLeaf{"OpticalPortType", state.OpticalPortType})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Components_Component_OpticalPort_State_InputPower
// The total input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type Components_Component_OpticalPort_State_InputPower struct {
    EntityData types.CommonEntityData
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

func (inputPower *Components_Component_OpticalPort_State_InputPower) GetEntityData() *types.CommonEntityData {
    inputPower.EntityData.YFilter = inputPower.YFilter
    inputPower.EntityData.YangName = "input-power"
    inputPower.EntityData.BundleName = "openconfig"
    inputPower.EntityData.ParentYangName = "state"
    inputPower.EntityData.SegmentPath = "input-power"
    inputPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inputPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inputPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inputPower.EntityData.Children = types.NewOrderedMap()
    inputPower.EntityData.Leafs = types.NewOrderedMap()
    inputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPower.Instant})
    inputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPower.Avg})
    inputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPower.Min})
    inputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPower.Max})

    inputPower.EntityData.YListKeys = []string {}

    return &(inputPower.EntityData)
}

// Components_Component_OpticalPort_State_OutputPower
// The total output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type Components_Component_OpticalPort_State_OutputPower struct {
    EntityData types.CommonEntityData
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

func (outputPower *Components_Component_OpticalPort_State_OutputPower) GetEntityData() *types.CommonEntityData {
    outputPower.EntityData.YFilter = outputPower.YFilter
    outputPower.EntityData.YangName = "output-power"
    outputPower.EntityData.BundleName = "openconfig"
    outputPower.EntityData.ParentYangName = "state"
    outputPower.EntityData.SegmentPath = "output-power"
    outputPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    outputPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    outputPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    outputPower.EntityData.Children = types.NewOrderedMap()
    outputPower.EntityData.Leafs = types.NewOrderedMap()
    outputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPower.Instant})
    outputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPower.Avg})
    outputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPower.Min})
    outputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPower.Max})

    outputPower.EntityData.YListKeys = []string {}

    return &(outputPower.EntityData)
}

// Components_Component_Transceiver
// Top-level container for client port transceiver data
type Components_Component_Transceiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for client port transceivers.
    Config Components_Component_Transceiver_Config

    // Operational state data for client port transceivers.
    State Components_Component_Transceiver_State

    // Enclosing container for client channels.
    PhysicalChannels Components_Component_Transceiver_PhysicalChannels
}

func (transceiver *Components_Component_Transceiver) GetEntityData() *types.CommonEntityData {
    transceiver.EntityData.YFilter = transceiver.YFilter
    transceiver.EntityData.YangName = "transceiver"
    transceiver.EntityData.BundleName = "openconfig"
    transceiver.EntityData.ParentYangName = "component"
    transceiver.EntityData.SegmentPath = "openconfig-platform-transceiver:transceiver"
    transceiver.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transceiver.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transceiver.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transceiver.EntityData.Children = types.NewOrderedMap()
    transceiver.EntityData.Children.Append("config", types.YChild{"Config", &transceiver.Config})
    transceiver.EntityData.Children.Append("state", types.YChild{"State", &transceiver.State})
    transceiver.EntityData.Children.Append("physical-channels", types.YChild{"PhysicalChannels", &transceiver.PhysicalChannels})
    transceiver.EntityData.Leafs = types.NewOrderedMap()

    transceiver.EntityData.YListKeys = []string {}

    return &(transceiver.EntityData)
}

// Components_Component_Transceiver_Config
// Configuration data for client port transceivers
type Components_Component_Transceiver_Config struct {
    EntityData types.CommonEntityData
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

func (config *Components_Component_Transceiver_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "transceiver"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("form-factor", types.YLeaf{"FormFactor", config.FormFactor})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Components_Component_Transceiver_State
// Operational state data for client port transceivers
type Components_Component_Transceiver_State struct {
    EntityData types.CommonEntityData
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

func (state *Components_Component_Transceiver_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "transceiver"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("form-factor", types.YLeaf{"FormFactor", state.FormFactor})
    state.EntityData.Leafs.Append("present", types.YLeaf{"Present", state.Present})
    state.EntityData.Leafs.Append("connector-type", types.YLeaf{"ConnectorType", state.ConnectorType})
    state.EntityData.Leafs.Append("internal-temp", types.YLeaf{"InternalTemp", state.InternalTemp})
    state.EntityData.Leafs.Append("vendor", types.YLeaf{"Vendor", state.Vendor})
    state.EntityData.Leafs.Append("vendor-part", types.YLeaf{"VendorPart", state.VendorPart})
    state.EntityData.Leafs.Append("vendor-rev", types.YLeaf{"VendorRev", state.VendorRev})
    state.EntityData.Leafs.Append("ethernet-compliance-code", types.YLeaf{"EthernetComplianceCode", state.EthernetComplianceCode})
    state.EntityData.Leafs.Append("sonet-sdh-compliance-code", types.YLeaf{"SonetSdhComplianceCode", state.SonetSdhComplianceCode})
    state.EntityData.Leafs.Append("otn-compliance-code", types.YLeaf{"OtnComplianceCode", state.OtnComplianceCode})
    state.EntityData.Leafs.Append("serial-no", types.YLeaf{"SerialNo", state.SerialNo})
    state.EntityData.Leafs.Append("date-code", types.YLeaf{"DateCode", state.DateCode})
    state.EntityData.Leafs.Append("fault-condition", types.YLeaf{"FaultCondition", state.FaultCondition})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of client channels, keyed by index within a physical client port.  A
    // physical port with a single channel would have a single zero-indexed
    // element. The type is slice of
    // Components_Component_Transceiver_PhysicalChannels_Channel.
    Channel []*Components_Component_Transceiver_PhysicalChannels_Channel
}

func (physicalChannels *Components_Component_Transceiver_PhysicalChannels) GetEntityData() *types.CommonEntityData {
    physicalChannels.EntityData.YFilter = physicalChannels.YFilter
    physicalChannels.EntityData.YangName = "physical-channels"
    physicalChannels.EntityData.BundleName = "openconfig"
    physicalChannels.EntityData.ParentYangName = "transceiver"
    physicalChannels.EntityData.SegmentPath = "physical-channels"
    physicalChannels.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    physicalChannels.EntityData.NamespaceTable = openconfig.GetNamespaces()
    physicalChannels.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    physicalChannels.EntityData.Children = types.NewOrderedMap()
    physicalChannels.EntityData.Children.Append("channel", types.YChild{"Channel", nil})
    for i := range physicalChannels.Channel {
        physicalChannels.EntityData.Children.Append(types.GetSegmentPath(physicalChannels.Channel[i]), types.YChild{"Channel", physicalChannels.Channel[i]})
    }
    physicalChannels.EntityData.Leafs = types.NewOrderedMap()

    physicalChannels.EntityData.YListKeys = []string {}

    return &(physicalChannels.EntityData)
}

// Components_Component_Transceiver_PhysicalChannels_Channel
// List of client channels, keyed by index within a physical
// client port.  A physical port with a single channel would
// have a single zero-indexed element
type Components_Component_Transceiver_PhysicalChannels_Channel struct {
    EntityData types.CommonEntityData
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

func (channel *Components_Component_Transceiver_PhysicalChannels_Channel) GetEntityData() *types.CommonEntityData {
    channel.EntityData.YFilter = channel.YFilter
    channel.EntityData.YangName = "channel"
    channel.EntityData.BundleName = "openconfig"
    channel.EntityData.ParentYangName = "physical-channels"
    channel.EntityData.SegmentPath = "channel" + types.AddKeyToken(channel.Index, "index")
    channel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    channel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    channel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    channel.EntityData.Children = types.NewOrderedMap()
    channel.EntityData.Children.Append("config", types.YChild{"Config", &channel.Config})
    channel.EntityData.Children.Append("state", types.YChild{"State", &channel.State})
    channel.EntityData.Leafs = types.NewOrderedMap()
    channel.EntityData.Leafs.Append("index", types.YLeaf{"Index", channel.Index})

    channel.EntityData.YListKeys = []string {"Index"}

    return &(channel.EntityData)
}

// Components_Component_Transceiver_PhysicalChannels_Channel_Config
// Configuration data for physical channels
type Components_Component_Transceiver_PhysicalChannels_Channel_Config struct {
    EntityData types.CommonEntityData
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

func (config *Components_Component_Transceiver_PhysicalChannels_Channel_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "channel"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("index", types.YLeaf{"Index", config.Index})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})
    config.EntityData.Leafs.Append("tx-laser", types.YLeaf{"TxLaser", config.TxLaser})
    config.EntityData.Leafs.Append("target-output-power", types.YLeaf{"TargetOutputPower", config.TargetOutputPower})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Components_Component_Transceiver_PhysicalChannels_Channel_State
// Operational state data for channels
type Components_Component_Transceiver_PhysicalChannels_Channel_State struct {
    EntityData types.CommonEntityData
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

func (state *Components_Component_Transceiver_PhysicalChannels_Channel_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "channel"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("output-power", types.YChild{"OutputPower", &state.OutputPower})
    state.EntityData.Children.Append("input-power", types.YChild{"InputPower", &state.InputPower})
    state.EntityData.Children.Append("laser-bias-current", types.YChild{"LaserBiasCurrent", &state.LaserBiasCurrent})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("tx-laser", types.YLeaf{"TxLaser", state.TxLaser})
    state.EntityData.Leafs.Append("target-output-power", types.YLeaf{"TargetOutputPower", state.TargetOutputPower})
    state.EntityData.Leafs.Append("output-frequency", types.YLeaf{"OutputFrequency", state.OutputFrequency})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower
// The output optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels. If
// avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower struct {
    EntityData types.CommonEntityData
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

func (outputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_OutputPower) GetEntityData() *types.CommonEntityData {
    outputPower.EntityData.YFilter = outputPower.YFilter
    outputPower.EntityData.YangName = "output-power"
    outputPower.EntityData.BundleName = "openconfig"
    outputPower.EntityData.ParentYangName = "state"
    outputPower.EntityData.SegmentPath = "output-power"
    outputPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    outputPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    outputPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    outputPower.EntityData.Children = types.NewOrderedMap()
    outputPower.EntityData.Leafs = types.NewOrderedMap()
    outputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPower.Instant})
    outputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPower.Avg})
    outputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPower.Min})
    outputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPower.Max})

    outputPower.EntityData.YListKeys = []string {}

    return &(outputPower.EntityData)
}

// Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower
// The input optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels.
// If avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower struct {
    EntityData types.CommonEntityData
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

func (inputPower *Components_Component_Transceiver_PhysicalChannels_Channel_State_InputPower) GetEntityData() *types.CommonEntityData {
    inputPower.EntityData.YFilter = inputPower.YFilter
    inputPower.EntityData.YangName = "input-power"
    inputPower.EntityData.BundleName = "openconfig"
    inputPower.EntityData.ParentYangName = "state"
    inputPower.EntityData.SegmentPath = "input-power"
    inputPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inputPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inputPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inputPower.EntityData.Children = types.NewOrderedMap()
    inputPower.EntityData.Leafs = types.NewOrderedMap()
    inputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPower.Instant})
    inputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPower.Avg})
    inputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPower.Min})
    inputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPower.Max})

    inputPower.EntityData.YListKeys = []string {}

    return &(inputPower.EntityData)
}

// Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power.  The current is expressed in mA
// with up to one decimal precision. If avg/min/max statistics
// are not supported, the target is expected to just supply
// the instant value
type Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent struct {
    EntityData types.CommonEntityData
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

func (laserBiasCurrent *Components_Component_Transceiver_PhysicalChannels_Channel_State_LaserBiasCurrent) GetEntityData() *types.CommonEntityData {
    laserBiasCurrent.EntityData.YFilter = laserBiasCurrent.YFilter
    laserBiasCurrent.EntityData.YangName = "laser-bias-current"
    laserBiasCurrent.EntityData.BundleName = "openconfig"
    laserBiasCurrent.EntityData.ParentYangName = "state"
    laserBiasCurrent.EntityData.SegmentPath = "laser-bias-current"
    laserBiasCurrent.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    laserBiasCurrent.EntityData.NamespaceTable = openconfig.GetNamespaces()
    laserBiasCurrent.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    laserBiasCurrent.EntityData.Children = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", laserBiasCurrent.Instant})
    laserBiasCurrent.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", laserBiasCurrent.Avg})
    laserBiasCurrent.EntityData.Leafs.Append("min", types.YLeaf{"Min", laserBiasCurrent.Min})
    laserBiasCurrent.EntityData.Leafs.Append("max", types.YLeaf{"Max", laserBiasCurrent.Max})

    laserBiasCurrent.EntityData.YListKeys = []string {}

    return &(laserBiasCurrent.EntityData)
}

// Components_Component_OpticalChannel
// Enclosing container for the list of optical channels
type Components_Component_OpticalChannel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for optical channels.
    Config Components_Component_OpticalChannel_Config

    // Operational state data for optical channels.
    State Components_Component_OpticalChannel_State
}

func (opticalChannel *Components_Component_OpticalChannel) GetEntityData() *types.CommonEntityData {
    opticalChannel.EntityData.YFilter = opticalChannel.YFilter
    opticalChannel.EntityData.YangName = "optical-channel"
    opticalChannel.EntityData.BundleName = "openconfig"
    opticalChannel.EntityData.ParentYangName = "component"
    opticalChannel.EntityData.SegmentPath = "openconfig-terminal-device:optical-channel"
    opticalChannel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalChannel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalChannel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalChannel.EntityData.Children = types.NewOrderedMap()
    opticalChannel.EntityData.Children.Append("config", types.YChild{"Config", &opticalChannel.Config})
    opticalChannel.EntityData.Children.Append("state", types.YChild{"State", &opticalChannel.State})
    opticalChannel.EntityData.Leafs = types.NewOrderedMap()

    opticalChannel.EntityData.YListKeys = []string {}

    return &(opticalChannel.EntityData)
}

// Components_Component_OpticalChannel_Config
// Configuration data for optical channels
type Components_Component_OpticalChannel_Config struct {
    EntityData types.CommonEntityData
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

func (config *Components_Component_OpticalChannel_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "optical-channel"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", config.Frequency})
    config.EntityData.Leafs.Append("target-output-power", types.YLeaf{"TargetOutputPower", config.TargetOutputPower})
    config.EntityData.Leafs.Append("operational-mode", types.YLeaf{"OperationalMode", config.OperationalMode})
    config.EntityData.Leafs.Append("line-port", types.YLeaf{"LinePort", config.LinePort})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Components_Component_OpticalChannel_State
// Operational state data for optical channels
type Components_Component_OpticalChannel_State struct {
    EntityData types.CommonEntityData
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

func (state *Components_Component_OpticalChannel_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "optical-channel"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("output-power", types.YChild{"OutputPower", &state.OutputPower})
    state.EntityData.Children.Append("input-power", types.YChild{"InputPower", &state.InputPower})
    state.EntityData.Children.Append("laser-bias-current", types.YChild{"LaserBiasCurrent", &state.LaserBiasCurrent})
    state.EntityData.Children.Append("chromatic-dispersion", types.YChild{"ChromaticDispersion", &state.ChromaticDispersion})
    state.EntityData.Children.Append("polarization-mode-dispersion", types.YChild{"PolarizationModeDispersion", &state.PolarizationModeDispersion})
    state.EntityData.Children.Append("second-order-polarization-mode-dispersion", types.YChild{"SecondOrderPolarizationModeDispersion", &state.SecondOrderPolarizationModeDispersion})
    state.EntityData.Children.Append("polarization-dependent-loss", types.YChild{"PolarizationDependentLoss", &state.PolarizationDependentLoss})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", state.Frequency})
    state.EntityData.Leafs.Append("target-output-power", types.YLeaf{"TargetOutputPower", state.TargetOutputPower})
    state.EntityData.Leafs.Append("operational-mode", types.YLeaf{"OperationalMode", state.OperationalMode})
    state.EntityData.Leafs.Append("line-port", types.YLeaf{"LinePort", state.LinePort})
    state.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", state.GroupId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Components_Component_OpticalChannel_State_OutputPower
// The output optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels. If
// avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_OpticalChannel_State_OutputPower struct {
    EntityData types.CommonEntityData
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

func (outputPower *Components_Component_OpticalChannel_State_OutputPower) GetEntityData() *types.CommonEntityData {
    outputPower.EntityData.YFilter = outputPower.YFilter
    outputPower.EntityData.YangName = "output-power"
    outputPower.EntityData.BundleName = "openconfig"
    outputPower.EntityData.ParentYangName = "state"
    outputPower.EntityData.SegmentPath = "output-power"
    outputPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    outputPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    outputPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    outputPower.EntityData.Children = types.NewOrderedMap()
    outputPower.EntityData.Leafs = types.NewOrderedMap()
    outputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPower.Instant})
    outputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPower.Avg})
    outputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPower.Min})
    outputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPower.Max})

    outputPower.EntityData.YListKeys = []string {}

    return &(outputPower.EntityData)
}

// Components_Component_OpticalChannel_State_InputPower
// The input optical power of this port in units of 0.01dBm.
// If the port is an aggregate of multiple physical channels,
// this attribute is the total power or sum of all channels.
// If avg/min/max statistics are not supported, the target is
// expected to just supply the instant value
type Components_Component_OpticalChannel_State_InputPower struct {
    EntityData types.CommonEntityData
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

func (inputPower *Components_Component_OpticalChannel_State_InputPower) GetEntityData() *types.CommonEntityData {
    inputPower.EntityData.YFilter = inputPower.YFilter
    inputPower.EntityData.YangName = "input-power"
    inputPower.EntityData.BundleName = "openconfig"
    inputPower.EntityData.ParentYangName = "state"
    inputPower.EntityData.SegmentPath = "input-power"
    inputPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inputPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inputPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inputPower.EntityData.Children = types.NewOrderedMap()
    inputPower.EntityData.Leafs = types.NewOrderedMap()
    inputPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPower.Instant})
    inputPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPower.Avg})
    inputPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPower.Min})
    inputPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPower.Max})

    inputPower.EntityData.YListKeys = []string {}

    return &(inputPower.EntityData)
}

// Components_Component_OpticalChannel_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power.  The current is expressed in mA
// with up to one decimal precision. If avg/min/max statistics
// are not supported, the target is expected to just supply
// the instant value
type Components_Component_OpticalChannel_State_LaserBiasCurrent struct {
    EntityData types.CommonEntityData
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

func (laserBiasCurrent *Components_Component_OpticalChannel_State_LaserBiasCurrent) GetEntityData() *types.CommonEntityData {
    laserBiasCurrent.EntityData.YFilter = laserBiasCurrent.YFilter
    laserBiasCurrent.EntityData.YangName = "laser-bias-current"
    laserBiasCurrent.EntityData.BundleName = "openconfig"
    laserBiasCurrent.EntityData.ParentYangName = "state"
    laserBiasCurrent.EntityData.SegmentPath = "laser-bias-current"
    laserBiasCurrent.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    laserBiasCurrent.EntityData.NamespaceTable = openconfig.GetNamespaces()
    laserBiasCurrent.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    laserBiasCurrent.EntityData.Children = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs = types.NewOrderedMap()
    laserBiasCurrent.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", laserBiasCurrent.Instant})
    laserBiasCurrent.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", laserBiasCurrent.Avg})
    laserBiasCurrent.EntityData.Leafs.Append("min", types.YLeaf{"Min", laserBiasCurrent.Min})
    laserBiasCurrent.EntityData.Leafs.Append("max", types.YLeaf{"Max", laserBiasCurrent.Max})

    laserBiasCurrent.EntityData.YListKeys = []string {}

    return &(laserBiasCurrent.EntityData)
}

// Components_Component_OpticalChannel_State_ChromaticDispersion
// Chromatic Dispersion of an optical channel
// in ps/nm as reported by receiver
type Components_Component_OpticalChannel_State_ChromaticDispersion struct {
    EntityData types.CommonEntityData
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

func (chromaticDispersion *Components_Component_OpticalChannel_State_ChromaticDispersion) GetEntityData() *types.CommonEntityData {
    chromaticDispersion.EntityData.YFilter = chromaticDispersion.YFilter
    chromaticDispersion.EntityData.YangName = "chromatic-dispersion"
    chromaticDispersion.EntityData.BundleName = "openconfig"
    chromaticDispersion.EntityData.ParentYangName = "state"
    chromaticDispersion.EntityData.SegmentPath = "chromatic-dispersion"
    chromaticDispersion.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    chromaticDispersion.EntityData.NamespaceTable = openconfig.GetNamespaces()
    chromaticDispersion.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    chromaticDispersion.EntityData.Children = types.NewOrderedMap()
    chromaticDispersion.EntityData.Leafs = types.NewOrderedMap()
    chromaticDispersion.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", chromaticDispersion.Instant})
    chromaticDispersion.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", chromaticDispersion.Avg})
    chromaticDispersion.EntityData.Leafs.Append("min", types.YLeaf{"Min", chromaticDispersion.Min})
    chromaticDispersion.EntityData.Leafs.Append("max", types.YLeaf{"Max", chromaticDispersion.Max})

    chromaticDispersion.EntityData.YListKeys = []string {}

    return &(chromaticDispersion.EntityData)
}

// Components_Component_OpticalChannel_State_PolarizationModeDispersion
// Polarization Mode Dispersion of an optical channel
// in ps as reported by receiver
type Components_Component_OpticalChannel_State_PolarizationModeDispersion struct {
    EntityData types.CommonEntityData
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

func (polarizationModeDispersion *Components_Component_OpticalChannel_State_PolarizationModeDispersion) GetEntityData() *types.CommonEntityData {
    polarizationModeDispersion.EntityData.YFilter = polarizationModeDispersion.YFilter
    polarizationModeDispersion.EntityData.YangName = "polarization-mode-dispersion"
    polarizationModeDispersion.EntityData.BundleName = "openconfig"
    polarizationModeDispersion.EntityData.ParentYangName = "state"
    polarizationModeDispersion.EntityData.SegmentPath = "polarization-mode-dispersion"
    polarizationModeDispersion.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    polarizationModeDispersion.EntityData.NamespaceTable = openconfig.GetNamespaces()
    polarizationModeDispersion.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    polarizationModeDispersion.EntityData.Children = types.NewOrderedMap()
    polarizationModeDispersion.EntityData.Leafs = types.NewOrderedMap()
    polarizationModeDispersion.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", polarizationModeDispersion.Instant})
    polarizationModeDispersion.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", polarizationModeDispersion.Avg})
    polarizationModeDispersion.EntityData.Leafs.Append("min", types.YLeaf{"Min", polarizationModeDispersion.Min})
    polarizationModeDispersion.EntityData.Leafs.Append("max", types.YLeaf{"Max", polarizationModeDispersion.Max})

    polarizationModeDispersion.EntityData.YListKeys = []string {}

    return &(polarizationModeDispersion.EntityData)
}

// Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion
// Second Order Polarization Mode Dispersion of an optical
// channel in ps^2 as reported by receiver
type Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion struct {
    EntityData types.CommonEntityData
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

func (secondOrderPolarizationModeDispersion *Components_Component_OpticalChannel_State_SecondOrderPolarizationModeDispersion) GetEntityData() *types.CommonEntityData {
    secondOrderPolarizationModeDispersion.EntityData.YFilter = secondOrderPolarizationModeDispersion.YFilter
    secondOrderPolarizationModeDispersion.EntityData.YangName = "second-order-polarization-mode-dispersion"
    secondOrderPolarizationModeDispersion.EntityData.BundleName = "openconfig"
    secondOrderPolarizationModeDispersion.EntityData.ParentYangName = "state"
    secondOrderPolarizationModeDispersion.EntityData.SegmentPath = "second-order-polarization-mode-dispersion"
    secondOrderPolarizationModeDispersion.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    secondOrderPolarizationModeDispersion.EntityData.NamespaceTable = openconfig.GetNamespaces()
    secondOrderPolarizationModeDispersion.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    secondOrderPolarizationModeDispersion.EntityData.Children = types.NewOrderedMap()
    secondOrderPolarizationModeDispersion.EntityData.Leafs = types.NewOrderedMap()
    secondOrderPolarizationModeDispersion.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", secondOrderPolarizationModeDispersion.Instant})
    secondOrderPolarizationModeDispersion.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", secondOrderPolarizationModeDispersion.Avg})
    secondOrderPolarizationModeDispersion.EntityData.Leafs.Append("min", types.YLeaf{"Min", secondOrderPolarizationModeDispersion.Min})
    secondOrderPolarizationModeDispersion.EntityData.Leafs.Append("max", types.YLeaf{"Max", secondOrderPolarizationModeDispersion.Max})

    secondOrderPolarizationModeDispersion.EntityData.YListKeys = []string {}

    return &(secondOrderPolarizationModeDispersion.EntityData)
}

// Components_Component_OpticalChannel_State_PolarizationDependentLoss
// Polarization Dependent Loss of an optical channel
// in dB as reported by receiver
type Components_Component_OpticalChannel_State_PolarizationDependentLoss struct {
    EntityData types.CommonEntityData
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

func (polarizationDependentLoss *Components_Component_OpticalChannel_State_PolarizationDependentLoss) GetEntityData() *types.CommonEntityData {
    polarizationDependentLoss.EntityData.YFilter = polarizationDependentLoss.YFilter
    polarizationDependentLoss.EntityData.YangName = "polarization-dependent-loss"
    polarizationDependentLoss.EntityData.BundleName = "openconfig"
    polarizationDependentLoss.EntityData.ParentYangName = "state"
    polarizationDependentLoss.EntityData.SegmentPath = "polarization-dependent-loss"
    polarizationDependentLoss.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    polarizationDependentLoss.EntityData.NamespaceTable = openconfig.GetNamespaces()
    polarizationDependentLoss.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    polarizationDependentLoss.EntityData.Children = types.NewOrderedMap()
    polarizationDependentLoss.EntityData.Leafs = types.NewOrderedMap()
    polarizationDependentLoss.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", polarizationDependentLoss.Instant})
    polarizationDependentLoss.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", polarizationDependentLoss.Avg})
    polarizationDependentLoss.EntityData.Leafs.Append("min", types.YLeaf{"Min", polarizationDependentLoss.Min})
    polarizationDependentLoss.EntityData.Leafs.Append("max", types.YLeaf{"Max", polarizationDependentLoss.Max})

    polarizationDependentLoss.EntityData.YListKeys = []string {}

    return &(polarizationDependentLoss.EntityData)
}

