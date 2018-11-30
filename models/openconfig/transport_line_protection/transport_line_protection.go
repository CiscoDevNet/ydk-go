// This model describes configuration and operational state data
// for optical line protection elements, deployed as part of a
// transport line system. An Automatic Protection Switch (APS)
// is typically installed in the same device as the amplifiers
// and wave-router, however an APS can also be a standalone
// device. In both scenarios, it serves the same purpose of
// providing protection using two dark fiber pairs to ensure the
// amplifiers can still receive a signal if one of the two fiber
// pairs is broken.
package transport_line_protection

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package transport_line_protection"))
    ydk.RegisterEntity("{http://openconfig.net/yang/optical-transport-line-protection aps}", reflect.TypeOf(Aps{}))
    ydk.RegisterEntity("openconfig-transport-line-protection:aps", reflect.TypeOf(Aps{}))
}

type PRIMARY struct {
}

func (id PRIMARY) String() string {
	return "openconfig-transport-line-protection:PRIMARY"
}

type APSPATHS struct {
}

func (id APSPATHS) String() string {
	return "openconfig-transport-line-protection:APS_PATHS"
}

type SECONDARY struct {
}

func (id SECONDARY) String() string {
	return "openconfig-transport-line-protection:SECONDARY"
}

// Aps
// Top level grouping for automatic protection switch data
type Aps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for list of automatic protection switch modules.
    ApsModules Aps_ApsModules
}

func (aps *Aps) GetEntityData() *types.CommonEntityData {
    aps.EntityData.YFilter = aps.YFilter
    aps.EntityData.YangName = "aps"
    aps.EntityData.BundleName = "openconfig"
    aps.EntityData.ParentYangName = "openconfig-transport-line-protection"
    aps.EntityData.SegmentPath = "openconfig-transport-line-protection:aps"
    aps.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aps.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aps.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aps.EntityData.Children = types.NewOrderedMap()
    aps.EntityData.Children.Append("aps-modules", types.YChild{"ApsModules", &aps.ApsModules})
    aps.EntityData.Leafs = types.NewOrderedMap()

    aps.EntityData.YListKeys = []string {}

    return &(aps.EntityData)
}

// Aps_ApsModules
// Enclosing container for list of automatic protection
// switch modules
type Aps_ApsModules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of automatic protection switch modules present in the device. The type
    // is slice of Aps_ApsModules_ApsModule.
    ApsModule []*Aps_ApsModules_ApsModule
}

func (apsModules *Aps_ApsModules) GetEntityData() *types.CommonEntityData {
    apsModules.EntityData.YFilter = apsModules.YFilter
    apsModules.EntityData.YangName = "aps-modules"
    apsModules.EntityData.BundleName = "openconfig"
    apsModules.EntityData.ParentYangName = "aps"
    apsModules.EntityData.SegmentPath = "aps-modules"
    apsModules.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    apsModules.EntityData.NamespaceTable = openconfig.GetNamespaces()
    apsModules.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    apsModules.EntityData.Children = types.NewOrderedMap()
    apsModules.EntityData.Children.Append("aps-module", types.YChild{"ApsModule", nil})
    for i := range apsModules.ApsModule {
        apsModules.EntityData.Children.Append(types.GetSegmentPath(apsModules.ApsModule[i]), types.YChild{"ApsModule", apsModules.ApsModule[i]})
    }
    apsModules.EntityData.Leafs = types.NewOrderedMap()

    apsModules.EntityData.YListKeys = []string {}

    return &(apsModules.EntityData)
}

// Aps_ApsModules_ApsModule
// List of automatic protection switch modules present
// in the device
type Aps_ApsModules_ApsModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the config name list key. The type is
    // string. Refers to
    // transport_line_protection.Aps_ApsModules_ApsModule_Config_Name
    Name interface{}

    // Configuration data for an automatic protection switch module.
    Config Aps_ApsModules_ApsModule_Config

    // Operational state data for an automatic protection switch module.
    State Aps_ApsModules_ApsModule_State

    // Top level grouping for automatic protection switch ports.
    Ports Aps_ApsModules_ApsModule_Ports
}

func (apsModule *Aps_ApsModules_ApsModule) GetEntityData() *types.CommonEntityData {
    apsModule.EntityData.YFilter = apsModule.YFilter
    apsModule.EntityData.YangName = "aps-module"
    apsModule.EntityData.BundleName = "openconfig"
    apsModule.EntityData.ParentYangName = "aps-modules"
    apsModule.EntityData.SegmentPath = "aps-module" + types.AddKeyToken(apsModule.Name, "name")
    apsModule.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    apsModule.EntityData.NamespaceTable = openconfig.GetNamespaces()
    apsModule.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    apsModule.EntityData.Children = types.NewOrderedMap()
    apsModule.EntityData.Children.Append("config", types.YChild{"Config", &apsModule.Config})
    apsModule.EntityData.Children.Append("state", types.YChild{"State", &apsModule.State})
    apsModule.EntityData.Children.Append("ports", types.YChild{"Ports", &apsModule.Ports})
    apsModule.EntityData.Leafs = types.NewOrderedMap()
    apsModule.EntityData.Leafs.Append("name", types.YLeaf{"Name", apsModule.Name})

    apsModule.EntityData.YListKeys = []string {"Name"}

    return &(apsModule.EntityData)
}

// Aps_ApsModules_ApsModule_Config
// Configuration data for an automatic protection
// switch module
type Aps_ApsModules_ApsModule_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the component name (in the platform model) corresponding to
    // this automatic protection switch module in the device. The type is string.
    // Refers to platform.Components_Component_Name
    Name interface{}

    // Revertive behavior of the module. If True, then automatically revert after
    // protection switch once the fault is restored. The type is bool.
    Revertive interface{}

    // The threshold at which the primary line port will switch to the opposite
    // line port in increments of 0.01 dBm. If the hardware supports only one
    // switch threshold for primary and and secondary ports then it is recommended
    // to set both primary-switch-threshold and secondary-switch-threshold to the
    // same value to be explicit. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    PrimarySwitchThreshold interface{}

    // The delta in 0.01 dB between the primary-switch-threshold and the signal
    // received before initiating a reversion in order to prevent toggling between
    // ports when an input signal is very close to threshold. If the hardware
    // supports only one switch hysteresis for primary and secondary ports then it
    // is recommended to set both primary-switch-threshold and
    // secondary-switch-threshold to the same value to be explicit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    PrimarySwitchHysteresis interface{}

    // The threshold at which the secondary line port will switch to the opposite
    // line port in increments of 0.01 dBm. If the hardware supports only one
    // switch threshold for primary and and secondary ports then it is recommended
    // to set both primary-switch-threshold and secondary-switch-threshold to the
    // same value to be explicit. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    SecondarySwitchThreshold interface{}

    // The delta in 0.01 dB between the secondary-switch-threshold and the signal
    // received before initiating a reversion in order to prevent toggling between
    // ports when an input signal is very close to threshold. If the hardware
    // supports only one switch hysteresis for primary and secondary ports then it
    // is recommended to set both primary-switch-threshold and
    // secondary-switch-threshold to the same value to be explicit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    SecondarySwitchHysteresis interface{}
}

func (config *Aps_ApsModules_ApsModule_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "aps-module"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("revertive", types.YLeaf{"Revertive", config.Revertive})
    config.EntityData.Leafs.Append("primary-switch-threshold", types.YLeaf{"PrimarySwitchThreshold", config.PrimarySwitchThreshold})
    config.EntityData.Leafs.Append("primary-switch-hysteresis", types.YLeaf{"PrimarySwitchHysteresis", config.PrimarySwitchHysteresis})
    config.EntityData.Leafs.Append("secondary-switch-threshold", types.YLeaf{"SecondarySwitchThreshold", config.SecondarySwitchThreshold})
    config.EntityData.Leafs.Append("secondary-switch-hysteresis", types.YLeaf{"SecondarySwitchHysteresis", config.SecondarySwitchHysteresis})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Aps_ApsModules_ApsModule_State
// Operational state data for an automatic protection
// switch module
type Aps_ApsModules_ApsModule_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the component name (in the platform model) corresponding to
    // this automatic protection switch module in the device. The type is string.
    // Refers to platform.Components_Component_Name
    Name interface{}

    // Revertive behavior of the module. If True, then automatically revert after
    // protection switch once the fault is restored. The type is bool.
    Revertive interface{}

    // The threshold at which the primary line port will switch to the opposite
    // line port in increments of 0.01 dBm. If the hardware supports only one
    // switch threshold for primary and and secondary ports then it is recommended
    // to set both primary-switch-threshold and secondary-switch-threshold to the
    // same value to be explicit. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    PrimarySwitchThreshold interface{}

    // The delta in 0.01 dB between the primary-switch-threshold and the signal
    // received before initiating a reversion in order to prevent toggling between
    // ports when an input signal is very close to threshold. If the hardware
    // supports only one switch hysteresis for primary and secondary ports then it
    // is recommended to set both primary-switch-threshold and
    // secondary-switch-threshold to the same value to be explicit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    PrimarySwitchHysteresis interface{}

    // The threshold at which the secondary line port will switch to the opposite
    // line port in increments of 0.01 dBm. If the hardware supports only one
    // switch threshold for primary and and secondary ports then it is recommended
    // to set both primary-switch-threshold and secondary-switch-threshold to the
    // same value to be explicit. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    SecondarySwitchThreshold interface{}

    // The delta in 0.01 dB between the secondary-switch-threshold and the signal
    // received before initiating a reversion in order to prevent toggling between
    // ports when an input signal is very close to threshold. If the hardware
    // supports only one switch hysteresis for primary and secondary ports then it
    // is recommended to set both primary-switch-threshold and
    // secondary-switch-threshold to the same value to be explicit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    SecondarySwitchHysteresis interface{}

    // Indicates which line path on the automatic protection switch is currently
    // the active path connected to the common port. The type is one of the
    // following: PRIMARYSECONDARY.
    ActivePath interface{}
}

func (state *Aps_ApsModules_ApsModule_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "aps-module"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("revertive", types.YLeaf{"Revertive", state.Revertive})
    state.EntityData.Leafs.Append("primary-switch-threshold", types.YLeaf{"PrimarySwitchThreshold", state.PrimarySwitchThreshold})
    state.EntityData.Leafs.Append("primary-switch-hysteresis", types.YLeaf{"PrimarySwitchHysteresis", state.PrimarySwitchHysteresis})
    state.EntityData.Leafs.Append("secondary-switch-threshold", types.YLeaf{"SecondarySwitchThreshold", state.SecondarySwitchThreshold})
    state.EntityData.Leafs.Append("secondary-switch-hysteresis", types.YLeaf{"SecondarySwitchHysteresis", state.SecondarySwitchHysteresis})
    state.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", state.ActivePath})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Aps_ApsModules_ApsModule_Ports
// Top level grouping for automatic protection switch ports
type Aps_ApsModules_ApsModule_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Container for information related to the line primary input port.
    LinePrimaryIn Aps_ApsModules_ApsModule_Ports_LinePrimaryIn

    // Container for information related to the line primary output port.
    LinePrimaryOut Aps_ApsModules_ApsModule_Ports_LinePrimaryOut

    // Container for information related to the line secondary input port.
    LineSecondaryIn Aps_ApsModules_ApsModule_Ports_LineSecondaryIn

    // Container for information related to the line secondary output port.
    LineSecondaryOut Aps_ApsModules_ApsModule_Ports_LineSecondaryOut

    // Container for information related to the line common input port.
    CommonIn Aps_ApsModules_ApsModule_Ports_CommonIn

    // Container for information related to the line common output port.
    CommonOutput Aps_ApsModules_ApsModule_Ports_CommonOutput
}

func (ports *Aps_ApsModules_ApsModule_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "openconfig"
    ports.EntityData.ParentYangName = "aps-module"
    ports.EntityData.SegmentPath = "ports"
    ports.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ports.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ports.EntityData.Children = types.NewOrderedMap()
    ports.EntityData.Children.Append("line-primary-in", types.YChild{"LinePrimaryIn", &ports.LinePrimaryIn})
    ports.EntityData.Children.Append("line-primary-out", types.YChild{"LinePrimaryOut", &ports.LinePrimaryOut})
    ports.EntityData.Children.Append("line-secondary-in", types.YChild{"LineSecondaryIn", &ports.LineSecondaryIn})
    ports.EntityData.Children.Append("line-secondary-out", types.YChild{"LineSecondaryOut", &ports.LineSecondaryOut})
    ports.EntityData.Children.Append("common-in", types.YChild{"CommonIn", &ports.CommonIn})
    ports.EntityData.Children.Append("common-output", types.YChild{"CommonOutput", &ports.CommonOutput})
    ports.EntityData.Leafs = types.NewOrderedMap()

    ports.EntityData.YListKeys = []string {}

    return &(ports.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn
// Container for information related to the line primary
// input port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the line primary input port.
    Config Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config

    // State data for the line primary input port.
    State Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State
}

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetEntityData() *types.CommonEntityData {
    linePrimaryIn.EntityData.YFilter = linePrimaryIn.YFilter
    linePrimaryIn.EntityData.YangName = "line-primary-in"
    linePrimaryIn.EntityData.BundleName = "openconfig"
    linePrimaryIn.EntityData.ParentYangName = "ports"
    linePrimaryIn.EntityData.SegmentPath = "line-primary-in"
    linePrimaryIn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    linePrimaryIn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    linePrimaryIn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    linePrimaryIn.EntityData.Children = types.NewOrderedMap()
    linePrimaryIn.EntityData.Children.Append("config", types.YChild{"Config", &linePrimaryIn.Config})
    linePrimaryIn.EntityData.Children.Append("state", types.YChild{"State", &linePrimaryIn.State})
    linePrimaryIn.EntityData.Leafs = types.NewOrderedMap()

    linePrimaryIn.EntityData.YListKeys = []string {}

    return &(linePrimaryIn.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config
// Configuration data for the line primary input port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf contains the configured, desired state of the port. Disabling the
    // port turns off alarm reporting for the port. The type is bool. The default
    // value is true.
    Enabled interface{}

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "line-primary-in"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", config.TargetAttenuation})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State
// State data for the line primary input port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf contains the configured, desired state of the port. Disabling the
    // port turns off alarm reporting for the port. The type is bool. The default
    // value is true.
    Enabled interface{}

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}

    // The attenuation of the variable optical attenuator associated with the port
    // in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Attenuation interface{}

    // The optical input power of this port in units of 0.01dBm. Optical input
    // power represents the signal traversing from an external destination into
    // the module. The power is measured before any attenuation. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    OpticalPower Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "line-primary-in"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("optical-power", types.YChild{"OpticalPower", &state.OpticalPower})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", state.TargetAttenuation})
    state.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", state.Attenuation})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower
// The optical input power of this port in units of
// 0.01dBm. Optical input power represents the signal
// traversing from an external destination into the module.
// The power is measured before any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetEntityData() *types.CommonEntityData {
    opticalPower.EntityData.YFilter = opticalPower.YFilter
    opticalPower.EntityData.YangName = "optical-power"
    opticalPower.EntityData.BundleName = "openconfig"
    opticalPower.EntityData.ParentYangName = "state"
    opticalPower.EntityData.SegmentPath = "optical-power"
    opticalPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalPower.EntityData.Children = types.NewOrderedMap()
    opticalPower.EntityData.Leafs = types.NewOrderedMap()
    opticalPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", opticalPower.Instant})
    opticalPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", opticalPower.Avg})
    opticalPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", opticalPower.Min})
    opticalPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", opticalPower.Max})

    opticalPower.EntityData.YListKeys = []string {}

    return &(opticalPower.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut
// Container for information related to the line primary
// output port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the line primary output port.
    Config Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config

    // State data for the line primary output port.
    State Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State
}

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetEntityData() *types.CommonEntityData {
    linePrimaryOut.EntityData.YFilter = linePrimaryOut.YFilter
    linePrimaryOut.EntityData.YangName = "line-primary-out"
    linePrimaryOut.EntityData.BundleName = "openconfig"
    linePrimaryOut.EntityData.ParentYangName = "ports"
    linePrimaryOut.EntityData.SegmentPath = "line-primary-out"
    linePrimaryOut.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    linePrimaryOut.EntityData.NamespaceTable = openconfig.GetNamespaces()
    linePrimaryOut.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    linePrimaryOut.EntityData.Children = types.NewOrderedMap()
    linePrimaryOut.EntityData.Children.Append("config", types.YChild{"Config", &linePrimaryOut.Config})
    linePrimaryOut.EntityData.Children.Append("state", types.YChild{"State", &linePrimaryOut.State})
    linePrimaryOut.EntityData.Leafs = types.NewOrderedMap()

    linePrimaryOut.EntityData.YListKeys = []string {}

    return &(linePrimaryOut.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config
// Configuration data for the line primary output port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "line-primary-out"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", config.TargetAttenuation})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State
// State data for the line primary output port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}

    // The attenuation of the variable optical attenuator associated with the port
    // in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Attenuation interface{}

    // The optical output power of this port in units of 0.01dBm. Optical output
    // power represents the signal traversing from the module to an external
    // destination. The power is measured after any attenuation. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    OpticalPower Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "line-primary-out"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("optical-power", types.YChild{"OpticalPower", &state.OpticalPower})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", state.TargetAttenuation})
    state.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", state.Attenuation})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower
// The optical output power of this port in units of
// 0.01dBm. Optical output power represents the signal
// traversing from the module to an external destination. The
// power is measured after any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetEntityData() *types.CommonEntityData {
    opticalPower.EntityData.YFilter = opticalPower.YFilter
    opticalPower.EntityData.YangName = "optical-power"
    opticalPower.EntityData.BundleName = "openconfig"
    opticalPower.EntityData.ParentYangName = "state"
    opticalPower.EntityData.SegmentPath = "optical-power"
    opticalPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalPower.EntityData.Children = types.NewOrderedMap()
    opticalPower.EntityData.Leafs = types.NewOrderedMap()
    opticalPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", opticalPower.Instant})
    opticalPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", opticalPower.Avg})
    opticalPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", opticalPower.Min})
    opticalPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", opticalPower.Max})

    opticalPower.EntityData.YListKeys = []string {}

    return &(opticalPower.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn
// Container for information related to the line secondary
// input port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the line secondary input port.
    Config Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config

    // State data for the line secondary input port.
    State Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State
}

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetEntityData() *types.CommonEntityData {
    lineSecondaryIn.EntityData.YFilter = lineSecondaryIn.YFilter
    lineSecondaryIn.EntityData.YangName = "line-secondary-in"
    lineSecondaryIn.EntityData.BundleName = "openconfig"
    lineSecondaryIn.EntityData.ParentYangName = "ports"
    lineSecondaryIn.EntityData.SegmentPath = "line-secondary-in"
    lineSecondaryIn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    lineSecondaryIn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    lineSecondaryIn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    lineSecondaryIn.EntityData.Children = types.NewOrderedMap()
    lineSecondaryIn.EntityData.Children.Append("config", types.YChild{"Config", &lineSecondaryIn.Config})
    lineSecondaryIn.EntityData.Children.Append("state", types.YChild{"State", &lineSecondaryIn.State})
    lineSecondaryIn.EntityData.Leafs = types.NewOrderedMap()

    lineSecondaryIn.EntityData.YListKeys = []string {}

    return &(lineSecondaryIn.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config
// Configuration data for the line secondary input port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf contains the configured, desired state of the port. Disabling the
    // port turns off alarm reporting for the port. The type is bool. The default
    // value is true.
    Enabled interface{}

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "line-secondary-in"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", config.TargetAttenuation})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State
// State data for the line secondary input port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf contains the configured, desired state of the port. Disabling the
    // port turns off alarm reporting for the port. The type is bool. The default
    // value is true.
    Enabled interface{}

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}

    // The attenuation of the variable optical attenuator associated with the port
    // in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Attenuation interface{}

    // The optical input power of this port in units of 0.01dBm. Optical input
    // power represents the signal traversing from an external destination into
    // the module. The power is measured before any attenuation. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    OpticalPower Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "line-secondary-in"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("optical-power", types.YChild{"OpticalPower", &state.OpticalPower})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", state.TargetAttenuation})
    state.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", state.Attenuation})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower
// The optical input power of this port in units of
// 0.01dBm. Optical input power represents the signal
// traversing from an external destination into the module.
// The power is measured before any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetEntityData() *types.CommonEntityData {
    opticalPower.EntityData.YFilter = opticalPower.YFilter
    opticalPower.EntityData.YangName = "optical-power"
    opticalPower.EntityData.BundleName = "openconfig"
    opticalPower.EntityData.ParentYangName = "state"
    opticalPower.EntityData.SegmentPath = "optical-power"
    opticalPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalPower.EntityData.Children = types.NewOrderedMap()
    opticalPower.EntityData.Leafs = types.NewOrderedMap()
    opticalPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", opticalPower.Instant})
    opticalPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", opticalPower.Avg})
    opticalPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", opticalPower.Min})
    opticalPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", opticalPower.Max})

    opticalPower.EntityData.YListKeys = []string {}

    return &(opticalPower.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut
// Container for information related to the line secondary
// output port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the line secondary output port.
    Config Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config

    // State data for the line secondary output port.
    State Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State
}

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetEntityData() *types.CommonEntityData {
    lineSecondaryOut.EntityData.YFilter = lineSecondaryOut.YFilter
    lineSecondaryOut.EntityData.YangName = "line-secondary-out"
    lineSecondaryOut.EntityData.BundleName = "openconfig"
    lineSecondaryOut.EntityData.ParentYangName = "ports"
    lineSecondaryOut.EntityData.SegmentPath = "line-secondary-out"
    lineSecondaryOut.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    lineSecondaryOut.EntityData.NamespaceTable = openconfig.GetNamespaces()
    lineSecondaryOut.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    lineSecondaryOut.EntityData.Children = types.NewOrderedMap()
    lineSecondaryOut.EntityData.Children.Append("config", types.YChild{"Config", &lineSecondaryOut.Config})
    lineSecondaryOut.EntityData.Children.Append("state", types.YChild{"State", &lineSecondaryOut.State})
    lineSecondaryOut.EntityData.Leafs = types.NewOrderedMap()

    lineSecondaryOut.EntityData.YListKeys = []string {}

    return &(lineSecondaryOut.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config
// Configuration data for the line secondary output port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "line-secondary-out"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", config.TargetAttenuation})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State
// State data for the line secondary output port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}

    // The attenuation of the variable optical attenuator associated with the port
    // in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Attenuation interface{}

    // The optical output power of this port in units of 0.01dBm. Optical output
    // power represents the signal traversing from the module to an external
    // destination. The power is measured after any attenuation. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    OpticalPower Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "line-secondary-out"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("optical-power", types.YChild{"OpticalPower", &state.OpticalPower})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", state.TargetAttenuation})
    state.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", state.Attenuation})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower
// The optical output power of this port in units of
// 0.01dBm. Optical output power represents the signal
// traversing from the module to an external destination. The
// power is measured after any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetEntityData() *types.CommonEntityData {
    opticalPower.EntityData.YFilter = opticalPower.YFilter
    opticalPower.EntityData.YangName = "optical-power"
    opticalPower.EntityData.BundleName = "openconfig"
    opticalPower.EntityData.ParentYangName = "state"
    opticalPower.EntityData.SegmentPath = "optical-power"
    opticalPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalPower.EntityData.Children = types.NewOrderedMap()
    opticalPower.EntityData.Leafs = types.NewOrderedMap()
    opticalPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", opticalPower.Instant})
    opticalPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", opticalPower.Avg})
    opticalPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", opticalPower.Min})
    opticalPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", opticalPower.Max})

    opticalPower.EntityData.YListKeys = []string {}

    return &(opticalPower.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonIn
// Container for information related to the line common
// input port
type Aps_ApsModules_ApsModule_Ports_CommonIn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the line common input port.
    Config Aps_ApsModules_ApsModule_Ports_CommonIn_Config

    // State data for the line common input port.
    State Aps_ApsModules_ApsModule_Ports_CommonIn_State
}

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetEntityData() *types.CommonEntityData {
    commonIn.EntityData.YFilter = commonIn.YFilter
    commonIn.EntityData.YangName = "common-in"
    commonIn.EntityData.BundleName = "openconfig"
    commonIn.EntityData.ParentYangName = "ports"
    commonIn.EntityData.SegmentPath = "common-in"
    commonIn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    commonIn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    commonIn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    commonIn.EntityData.Children = types.NewOrderedMap()
    commonIn.EntityData.Children.Append("config", types.YChild{"Config", &commonIn.Config})
    commonIn.EntityData.Children.Append("state", types.YChild{"State", &commonIn.State})
    commonIn.EntityData.Leafs = types.NewOrderedMap()

    commonIn.EntityData.YListKeys = []string {}

    return &(commonIn.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonIn_Config
// Configuration data for the line common input port
type Aps_ApsModules_ApsModule_Ports_CommonIn_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf contains the configured, desired state of the port. Disabling the
    // port turns off alarm reporting for the port. The type is bool. The default
    // value is true.
    Enabled interface{}

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "common-in"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", config.TargetAttenuation})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonIn_State
// State data for the line common input port
type Aps_ApsModules_ApsModule_Ports_CommonIn_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf contains the configured, desired state of the port. Disabling the
    // port turns off alarm reporting for the port. The type is bool. The default
    // value is true.
    Enabled interface{}

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}

    // The attenuation of the variable optical attenuator associated with the port
    // in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Attenuation interface{}

    // The optical input power of this port in units of 0.01dBm. Optical input
    // power represents the signal traversing from an external destination into
    // the module. The power is measured before any attenuation. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    OpticalPower Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "common-in"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("optical-power", types.YChild{"OpticalPower", &state.OpticalPower})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", state.TargetAttenuation})
    state.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", state.Attenuation})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower
// The optical input power of this port in units of
// 0.01dBm. Optical input power represents the signal
// traversing from an external destination into the module.
// The power is measured before any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetEntityData() *types.CommonEntityData {
    opticalPower.EntityData.YFilter = opticalPower.YFilter
    opticalPower.EntityData.YangName = "optical-power"
    opticalPower.EntityData.BundleName = "openconfig"
    opticalPower.EntityData.ParentYangName = "state"
    opticalPower.EntityData.SegmentPath = "optical-power"
    opticalPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalPower.EntityData.Children = types.NewOrderedMap()
    opticalPower.EntityData.Leafs = types.NewOrderedMap()
    opticalPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", opticalPower.Instant})
    opticalPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", opticalPower.Avg})
    opticalPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", opticalPower.Min})
    opticalPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", opticalPower.Max})

    opticalPower.EntityData.YListKeys = []string {}

    return &(opticalPower.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonOutput
// Container for information related to the line common
// output port
type Aps_ApsModules_ApsModule_Ports_CommonOutput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the line common output port.
    Config Aps_ApsModules_ApsModule_Ports_CommonOutput_Config

    // State data for the line common output port.
    State Aps_ApsModules_ApsModule_Ports_CommonOutput_State
}

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetEntityData() *types.CommonEntityData {
    commonOutput.EntityData.YFilter = commonOutput.YFilter
    commonOutput.EntityData.YangName = "common-output"
    commonOutput.EntityData.BundleName = "openconfig"
    commonOutput.EntityData.ParentYangName = "ports"
    commonOutput.EntityData.SegmentPath = "common-output"
    commonOutput.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    commonOutput.EntityData.NamespaceTable = openconfig.GetNamespaces()
    commonOutput.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    commonOutput.EntityData.Children = types.NewOrderedMap()
    commonOutput.EntityData.Children.Append("config", types.YChild{"Config", &commonOutput.Config})
    commonOutput.EntityData.Children.Append("state", types.YChild{"State", &commonOutput.State})
    commonOutput.EntityData.Leafs = types.NewOrderedMap()

    commonOutput.EntityData.YListKeys = []string {}

    return &(commonOutput.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonOutput_Config
// Configuration data for the line common output port
type Aps_ApsModules_ApsModule_Ports_CommonOutput_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "common-output"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", config.TargetAttenuation})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonOutput_State
// State data for the line common output port
type Aps_ApsModules_ApsModule_Ports_CommonOutput_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}

    // The attenuation of the variable optical attenuator associated with the port
    // in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Attenuation interface{}

    // The optical output power of this port in units of 0.01dBm. Optical output
    // power represents the signal traversing from the module to an external
    // destination. The power is measured after any attenuation. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    OpticalPower Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "common-output"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("optical-power", types.YChild{"OpticalPower", &state.OpticalPower})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("target-attenuation", types.YLeaf{"TargetAttenuation", state.TargetAttenuation})
    state.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", state.Attenuation})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower
// The optical output power of this port in units of
// 0.01dBm. Optical output power represents the signal
// traversing from the module to an external destination. The
// power is measured after any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetEntityData() *types.CommonEntityData {
    opticalPower.EntityData.YFilter = opticalPower.YFilter
    opticalPower.EntityData.YangName = "optical-power"
    opticalPower.EntityData.BundleName = "openconfig"
    opticalPower.EntityData.ParentYangName = "state"
    opticalPower.EntityData.SegmentPath = "optical-power"
    opticalPower.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalPower.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalPower.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalPower.EntityData.Children = types.NewOrderedMap()
    opticalPower.EntityData.Leafs = types.NewOrderedMap()
    opticalPower.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", opticalPower.Instant})
    opticalPower.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", opticalPower.Avg})
    opticalPower.EntityData.Leafs.Append("min", types.YLeaf{"Min", opticalPower.Min})
    opticalPower.EntityData.Leafs.Append("max", types.YLeaf{"Max", opticalPower.Max})

    opticalPower.EntityData.YListKeys = []string {}

    return &(opticalPower.EntityData)
}

