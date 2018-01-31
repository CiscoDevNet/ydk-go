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
    parent types.Entity
    YFilter yfilter.YFilter

    // Enclosing container for list of automatic protection switch modules.
    ApsModules Aps_ApsModules
}

func (aps *Aps) GetFilter() yfilter.YFilter { return aps.YFilter }

func (aps *Aps) SetFilter(yf yfilter.YFilter) { aps.YFilter = yf }

func (aps *Aps) GetGoName(yname string) string {
    if yname == "aps-modules" { return "ApsModules" }
    return ""
}

func (aps *Aps) GetSegmentPath() string {
    return "openconfig-transport-line-protection:aps"
}

func (aps *Aps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aps-modules" {
        return &aps.ApsModules
    }
    return nil
}

func (aps *Aps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aps-modules"] = &aps.ApsModules
    return children
}

func (aps *Aps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aps *Aps) GetBundleName() string { return "openconfig" }

func (aps *Aps) GetYangName() string { return "aps" }

func (aps *Aps) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aps *Aps) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aps *Aps) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aps *Aps) SetParent(parent types.Entity) { aps.parent = parent }

func (aps *Aps) GetParent() types.Entity { return aps.parent }

func (aps *Aps) GetParentYangName() string { return "openconfig-transport-line-protection" }

// Aps_ApsModules
// Enclosing container for list of automatic protection
// switch modules
type Aps_ApsModules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of automatic protection switch modules present in the device. The type
    // is slice of Aps_ApsModules_ApsModule.
    ApsModule []Aps_ApsModules_ApsModule
}

func (apsModules *Aps_ApsModules) GetFilter() yfilter.YFilter { return apsModules.YFilter }

func (apsModules *Aps_ApsModules) SetFilter(yf yfilter.YFilter) { apsModules.YFilter = yf }

func (apsModules *Aps_ApsModules) GetGoName(yname string) string {
    if yname == "aps-module" { return "ApsModule" }
    return ""
}

func (apsModules *Aps_ApsModules) GetSegmentPath() string {
    return "aps-modules"
}

func (apsModules *Aps_ApsModules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aps-module" {
        for _, c := range apsModules.ApsModule {
            if apsModules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aps_ApsModules_ApsModule{}
        apsModules.ApsModule = append(apsModules.ApsModule, child)
        return &apsModules.ApsModule[len(apsModules.ApsModule)-1]
    }
    return nil
}

func (apsModules *Aps_ApsModules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range apsModules.ApsModule {
        children[apsModules.ApsModule[i].GetSegmentPath()] = &apsModules.ApsModule[i]
    }
    return children
}

func (apsModules *Aps_ApsModules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (apsModules *Aps_ApsModules) GetBundleName() string { return "openconfig" }

func (apsModules *Aps_ApsModules) GetYangName() string { return "aps-modules" }

func (apsModules *Aps_ApsModules) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (apsModules *Aps_ApsModules) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (apsModules *Aps_ApsModules) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (apsModules *Aps_ApsModules) SetParent(parent types.Entity) { apsModules.parent = parent }

func (apsModules *Aps_ApsModules) GetParent() types.Entity { return apsModules.parent }

func (apsModules *Aps_ApsModules) GetParentYangName() string { return "aps" }

// Aps_ApsModules_ApsModule
// List of automatic protection switch modules present
// in the device
type Aps_ApsModules_ApsModule struct {
    parent types.Entity
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

func (apsModule *Aps_ApsModules_ApsModule) GetFilter() yfilter.YFilter { return apsModule.YFilter }

func (apsModule *Aps_ApsModules_ApsModule) SetFilter(yf yfilter.YFilter) { apsModule.YFilter = yf }

func (apsModule *Aps_ApsModules_ApsModule) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "ports" { return "Ports" }
    return ""
}

func (apsModule *Aps_ApsModules_ApsModule) GetSegmentPath() string {
    return "aps-module" + "[name='" + fmt.Sprintf("%v", apsModule.Name) + "']"
}

func (apsModule *Aps_ApsModules_ApsModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &apsModule.Config
    }
    if childYangName == "state" {
        return &apsModule.State
    }
    if childYangName == "ports" {
        return &apsModule.Ports
    }
    return nil
}

func (apsModule *Aps_ApsModules_ApsModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &apsModule.Config
    children["state"] = &apsModule.State
    children["ports"] = &apsModule.Ports
    return children
}

func (apsModule *Aps_ApsModules_ApsModule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = apsModule.Name
    return leafs
}

func (apsModule *Aps_ApsModules_ApsModule) GetBundleName() string { return "openconfig" }

func (apsModule *Aps_ApsModules_ApsModule) GetYangName() string { return "aps-module" }

func (apsModule *Aps_ApsModules_ApsModule) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (apsModule *Aps_ApsModules_ApsModule) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (apsModule *Aps_ApsModules_ApsModule) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (apsModule *Aps_ApsModules_ApsModule) SetParent(parent types.Entity) { apsModule.parent = parent }

func (apsModule *Aps_ApsModules_ApsModule) GetParent() types.Entity { return apsModule.parent }

func (apsModule *Aps_ApsModules_ApsModule) GetParentYangName() string { return "aps-modules" }

// Aps_ApsModules_ApsModule_Config
// Configuration data for an automatic protection
// switch module
type Aps_ApsModules_ApsModule_Config struct {
    parent types.Entity
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

func (config *Aps_ApsModules_ApsModule_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Aps_ApsModules_ApsModule_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Aps_ApsModules_ApsModule_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "revertive" { return "Revertive" }
    if yname == "primary-switch-threshold" { return "PrimarySwitchThreshold" }
    if yname == "primary-switch-hysteresis" { return "PrimarySwitchHysteresis" }
    if yname == "secondary-switch-threshold" { return "SecondarySwitchThreshold" }
    if yname == "secondary-switch-hysteresis" { return "SecondarySwitchHysteresis" }
    return ""
}

func (config *Aps_ApsModules_ApsModule_Config) GetSegmentPath() string {
    return "config"
}

func (config *Aps_ApsModules_ApsModule_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Aps_ApsModules_ApsModule_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Aps_ApsModules_ApsModule_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["revertive"] = config.Revertive
    leafs["primary-switch-threshold"] = config.PrimarySwitchThreshold
    leafs["primary-switch-hysteresis"] = config.PrimarySwitchHysteresis
    leafs["secondary-switch-threshold"] = config.SecondarySwitchThreshold
    leafs["secondary-switch-hysteresis"] = config.SecondarySwitchHysteresis
    return leafs
}

func (config *Aps_ApsModules_ApsModule_Config) GetBundleName() string { return "openconfig" }

func (config *Aps_ApsModules_ApsModule_Config) GetYangName() string { return "config" }

func (config *Aps_ApsModules_ApsModule_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Aps_ApsModules_ApsModule_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Aps_ApsModules_ApsModule_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Aps_ApsModules_ApsModule_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Aps_ApsModules_ApsModule_Config) GetParent() types.Entity { return config.parent }

func (config *Aps_ApsModules_ApsModule_Config) GetParentYangName() string { return "aps-module" }

// Aps_ApsModules_ApsModule_State
// Operational state data for an automatic protection
// switch module
type Aps_ApsModules_ApsModule_State struct {
    parent types.Entity
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

func (state *Aps_ApsModules_ApsModule_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Aps_ApsModules_ApsModule_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Aps_ApsModules_ApsModule_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "revertive" { return "Revertive" }
    if yname == "primary-switch-threshold" { return "PrimarySwitchThreshold" }
    if yname == "primary-switch-hysteresis" { return "PrimarySwitchHysteresis" }
    if yname == "secondary-switch-threshold" { return "SecondarySwitchThreshold" }
    if yname == "secondary-switch-hysteresis" { return "SecondarySwitchHysteresis" }
    if yname == "active-path" { return "ActivePath" }
    return ""
}

func (state *Aps_ApsModules_ApsModule_State) GetSegmentPath() string {
    return "state"
}

func (state *Aps_ApsModules_ApsModule_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Aps_ApsModules_ApsModule_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Aps_ApsModules_ApsModule_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["revertive"] = state.Revertive
    leafs["primary-switch-threshold"] = state.PrimarySwitchThreshold
    leafs["primary-switch-hysteresis"] = state.PrimarySwitchHysteresis
    leafs["secondary-switch-threshold"] = state.SecondarySwitchThreshold
    leafs["secondary-switch-hysteresis"] = state.SecondarySwitchHysteresis
    leafs["active-path"] = state.ActivePath
    return leafs
}

func (state *Aps_ApsModules_ApsModule_State) GetBundleName() string { return "openconfig" }

func (state *Aps_ApsModules_ApsModule_State) GetYangName() string { return "state" }

func (state *Aps_ApsModules_ApsModule_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Aps_ApsModules_ApsModule_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Aps_ApsModules_ApsModule_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Aps_ApsModules_ApsModule_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Aps_ApsModules_ApsModule_State) GetParent() types.Entity { return state.parent }

func (state *Aps_ApsModules_ApsModule_State) GetParentYangName() string { return "aps-module" }

// Aps_ApsModules_ApsModule_Ports
// Top level grouping for automatic protection switch ports
type Aps_ApsModules_ApsModule_Ports struct {
    parent types.Entity
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

func (ports *Aps_ApsModules_ApsModule_Ports) GetFilter() yfilter.YFilter { return ports.YFilter }

func (ports *Aps_ApsModules_ApsModule_Ports) SetFilter(yf yfilter.YFilter) { ports.YFilter = yf }

func (ports *Aps_ApsModules_ApsModule_Ports) GetGoName(yname string) string {
    if yname == "line-primary-in" { return "LinePrimaryIn" }
    if yname == "line-primary-out" { return "LinePrimaryOut" }
    if yname == "line-secondary-in" { return "LineSecondaryIn" }
    if yname == "line-secondary-out" { return "LineSecondaryOut" }
    if yname == "common-in" { return "CommonIn" }
    if yname == "common-output" { return "CommonOutput" }
    return ""
}

func (ports *Aps_ApsModules_ApsModule_Ports) GetSegmentPath() string {
    return "ports"
}

func (ports *Aps_ApsModules_ApsModule_Ports) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "line-primary-in" {
        return &ports.LinePrimaryIn
    }
    if childYangName == "line-primary-out" {
        return &ports.LinePrimaryOut
    }
    if childYangName == "line-secondary-in" {
        return &ports.LineSecondaryIn
    }
    if childYangName == "line-secondary-out" {
        return &ports.LineSecondaryOut
    }
    if childYangName == "common-in" {
        return &ports.CommonIn
    }
    if childYangName == "common-output" {
        return &ports.CommonOutput
    }
    return nil
}

func (ports *Aps_ApsModules_ApsModule_Ports) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["line-primary-in"] = &ports.LinePrimaryIn
    children["line-primary-out"] = &ports.LinePrimaryOut
    children["line-secondary-in"] = &ports.LineSecondaryIn
    children["line-secondary-out"] = &ports.LineSecondaryOut
    children["common-in"] = &ports.CommonIn
    children["common-output"] = &ports.CommonOutput
    return children
}

func (ports *Aps_ApsModules_ApsModule_Ports) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ports *Aps_ApsModules_ApsModule_Ports) GetBundleName() string { return "openconfig" }

func (ports *Aps_ApsModules_ApsModule_Ports) GetYangName() string { return "ports" }

func (ports *Aps_ApsModules_ApsModule_Ports) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ports *Aps_ApsModules_ApsModule_Ports) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ports *Aps_ApsModules_ApsModule_Ports) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ports *Aps_ApsModules_ApsModule_Ports) SetParent(parent types.Entity) { ports.parent = parent }

func (ports *Aps_ApsModules_ApsModule_Ports) GetParent() types.Entity { return ports.parent }

func (ports *Aps_ApsModules_ApsModule_Ports) GetParentYangName() string { return "aps-module" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn
// Container for information related to the line primary
// input port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the line primary input port.
    Config Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config

    // State data for the line primary input port.
    State Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State
}

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetFilter() yfilter.YFilter { return linePrimaryIn.YFilter }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) SetFilter(yf yfilter.YFilter) { linePrimaryIn.YFilter = yf }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetSegmentPath() string {
    return "line-primary-in"
}

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &linePrimaryIn.Config
    }
    if childYangName == "state" {
        return &linePrimaryIn.State
    }
    return nil
}

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &linePrimaryIn.Config
    children["state"] = &linePrimaryIn.State
    return children
}

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetBundleName() string { return "openconfig" }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetYangName() string { return "line-primary-in" }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) SetParent(parent types.Entity) { linePrimaryIn.parent = parent }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetParent() types.Entity { return linePrimaryIn.parent }

func (linePrimaryIn *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn) GetParentYangName() string { return "ports" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config
// Configuration data for the line primary input port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config struct {
    parent types.Entity
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

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "target-attenuation" { return "TargetAttenuation" }
    return ""
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetSegmentPath() string {
    return "config"
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["target-attenuation"] = config.TargetAttenuation
    return leafs
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetBundleName() string { return "openconfig" }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetYangName() string { return "config" }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetParent() types.Entity { return config.parent }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_Config) GetParentYangName() string { return "line-primary-in" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State
// State data for the line primary input port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State struct {
    parent types.Entity
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

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "target-attenuation" { return "TargetAttenuation" }
    if yname == "attenuation" { return "Attenuation" }
    if yname == "optical-power" { return "OpticalPower" }
    return ""
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetSegmentPath() string {
    return "state"
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-power" {
        return &state.OpticalPower
    }
    return nil
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-power"] = &state.OpticalPower
    return children
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["target-attenuation"] = state.TargetAttenuation
    leafs["attenuation"] = state.Attenuation
    return leafs
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetBundleName() string { return "openconfig" }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetYangName() string { return "state" }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetParent() types.Entity { return state.parent }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State) GetParentYangName() string { return "line-primary-in" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower
// The optical input power of this port in units of
// 0.01dBm. Optical input power represents the signal
// traversing from an external destination into the module.
// The power is measured before any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetFilter() yfilter.YFilter { return opticalPower.YFilter }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) SetFilter(yf yfilter.YFilter) { opticalPower.YFilter = yf }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetSegmentPath() string {
    return "optical-power"
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = opticalPower.Instant
    leafs["avg"] = opticalPower.Avg
    leafs["min"] = opticalPower.Min
    leafs["max"] = opticalPower.Max
    return leafs
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetBundleName() string { return "openconfig" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetYangName() string { return "optical-power" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) SetParent(parent types.Entity) { opticalPower.parent = parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetParent() types.Entity { return opticalPower.parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryIn_State_OpticalPower) GetParentYangName() string { return "state" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut
// Container for information related to the line primary
// output port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the line primary output port.
    Config Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config

    // State data for the line primary output port.
    State Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State
}

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetFilter() yfilter.YFilter { return linePrimaryOut.YFilter }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) SetFilter(yf yfilter.YFilter) { linePrimaryOut.YFilter = yf }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetSegmentPath() string {
    return "line-primary-out"
}

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &linePrimaryOut.Config
    }
    if childYangName == "state" {
        return &linePrimaryOut.State
    }
    return nil
}

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &linePrimaryOut.Config
    children["state"] = &linePrimaryOut.State
    return children
}

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetBundleName() string { return "openconfig" }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetYangName() string { return "line-primary-out" }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) SetParent(parent types.Entity) { linePrimaryOut.parent = parent }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetParent() types.Entity { return linePrimaryOut.parent }

func (linePrimaryOut *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut) GetParentYangName() string { return "ports" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config
// Configuration data for the line primary output port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetGoName(yname string) string {
    if yname == "target-attenuation" { return "TargetAttenuation" }
    return ""
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetSegmentPath() string {
    return "config"
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-attenuation"] = config.TargetAttenuation
    return leafs
}

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetBundleName() string { return "openconfig" }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetYangName() string { return "config" }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetParent() types.Entity { return config.parent }

func (config *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_Config) GetParentYangName() string { return "line-primary-out" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State
// State data for the line primary output port
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State struct {
    parent types.Entity
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

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetGoName(yname string) string {
    if yname == "target-attenuation" { return "TargetAttenuation" }
    if yname == "attenuation" { return "Attenuation" }
    if yname == "optical-power" { return "OpticalPower" }
    return ""
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetSegmentPath() string {
    return "state"
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-power" {
        return &state.OpticalPower
    }
    return nil
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-power"] = &state.OpticalPower
    return children
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-attenuation"] = state.TargetAttenuation
    leafs["attenuation"] = state.Attenuation
    return leafs
}

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetBundleName() string { return "openconfig" }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetYangName() string { return "state" }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetParent() types.Entity { return state.parent }

func (state *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State) GetParentYangName() string { return "line-primary-out" }

// Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower
// The optical output power of this port in units of
// 0.01dBm. Optical output power represents the signal
// traversing from the module to an external destination. The
// power is measured after any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetFilter() yfilter.YFilter { return opticalPower.YFilter }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) SetFilter(yf yfilter.YFilter) { opticalPower.YFilter = yf }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetSegmentPath() string {
    return "optical-power"
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = opticalPower.Instant
    leafs["avg"] = opticalPower.Avg
    leafs["min"] = opticalPower.Min
    leafs["max"] = opticalPower.Max
    return leafs
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetBundleName() string { return "openconfig" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetYangName() string { return "optical-power" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) SetParent(parent types.Entity) { opticalPower.parent = parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetParent() types.Entity { return opticalPower.parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LinePrimaryOut_State_OpticalPower) GetParentYangName() string { return "state" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn
// Container for information related to the line secondary
// input port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the line secondary input port.
    Config Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config

    // State data for the line secondary input port.
    State Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State
}

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetFilter() yfilter.YFilter { return lineSecondaryIn.YFilter }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) SetFilter(yf yfilter.YFilter) { lineSecondaryIn.YFilter = yf }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetSegmentPath() string {
    return "line-secondary-in"
}

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &lineSecondaryIn.Config
    }
    if childYangName == "state" {
        return &lineSecondaryIn.State
    }
    return nil
}

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &lineSecondaryIn.Config
    children["state"] = &lineSecondaryIn.State
    return children
}

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetBundleName() string { return "openconfig" }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetYangName() string { return "line-secondary-in" }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) SetParent(parent types.Entity) { lineSecondaryIn.parent = parent }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetParent() types.Entity { return lineSecondaryIn.parent }

func (lineSecondaryIn *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn) GetParentYangName() string { return "ports" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config
// Configuration data for the line secondary input port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config struct {
    parent types.Entity
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

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "target-attenuation" { return "TargetAttenuation" }
    return ""
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetSegmentPath() string {
    return "config"
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["target-attenuation"] = config.TargetAttenuation
    return leafs
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetBundleName() string { return "openconfig" }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetYangName() string { return "config" }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetParent() types.Entity { return config.parent }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_Config) GetParentYangName() string { return "line-secondary-in" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State
// State data for the line secondary input port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State struct {
    parent types.Entity
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

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "target-attenuation" { return "TargetAttenuation" }
    if yname == "attenuation" { return "Attenuation" }
    if yname == "optical-power" { return "OpticalPower" }
    return ""
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetSegmentPath() string {
    return "state"
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-power" {
        return &state.OpticalPower
    }
    return nil
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-power"] = &state.OpticalPower
    return children
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["target-attenuation"] = state.TargetAttenuation
    leafs["attenuation"] = state.Attenuation
    return leafs
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetBundleName() string { return "openconfig" }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetYangName() string { return "state" }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetParent() types.Entity { return state.parent }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State) GetParentYangName() string { return "line-secondary-in" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower
// The optical input power of this port in units of
// 0.01dBm. Optical input power represents the signal
// traversing from an external destination into the module.
// The power is measured before any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetFilter() yfilter.YFilter { return opticalPower.YFilter }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) SetFilter(yf yfilter.YFilter) { opticalPower.YFilter = yf }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetSegmentPath() string {
    return "optical-power"
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = opticalPower.Instant
    leafs["avg"] = opticalPower.Avg
    leafs["min"] = opticalPower.Min
    leafs["max"] = opticalPower.Max
    return leafs
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetBundleName() string { return "openconfig" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetYangName() string { return "optical-power" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) SetParent(parent types.Entity) { opticalPower.parent = parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetParent() types.Entity { return opticalPower.parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryIn_State_OpticalPower) GetParentYangName() string { return "state" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut
// Container for information related to the line secondary
// output port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the line secondary output port.
    Config Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config

    // State data for the line secondary output port.
    State Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State
}

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetFilter() yfilter.YFilter { return lineSecondaryOut.YFilter }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) SetFilter(yf yfilter.YFilter) { lineSecondaryOut.YFilter = yf }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetSegmentPath() string {
    return "line-secondary-out"
}

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &lineSecondaryOut.Config
    }
    if childYangName == "state" {
        return &lineSecondaryOut.State
    }
    return nil
}

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &lineSecondaryOut.Config
    children["state"] = &lineSecondaryOut.State
    return children
}

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetBundleName() string { return "openconfig" }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetYangName() string { return "line-secondary-out" }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) SetParent(parent types.Entity) { lineSecondaryOut.parent = parent }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetParent() types.Entity { return lineSecondaryOut.parent }

func (lineSecondaryOut *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut) GetParentYangName() string { return "ports" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config
// Configuration data for the line secondary output port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetGoName(yname string) string {
    if yname == "target-attenuation" { return "TargetAttenuation" }
    return ""
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetSegmentPath() string {
    return "config"
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-attenuation"] = config.TargetAttenuation
    return leafs
}

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetBundleName() string { return "openconfig" }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetYangName() string { return "config" }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetParent() types.Entity { return config.parent }

func (config *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_Config) GetParentYangName() string { return "line-secondary-out" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State
// State data for the line secondary output port
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State struct {
    parent types.Entity
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

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetGoName(yname string) string {
    if yname == "target-attenuation" { return "TargetAttenuation" }
    if yname == "attenuation" { return "Attenuation" }
    if yname == "optical-power" { return "OpticalPower" }
    return ""
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetSegmentPath() string {
    return "state"
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-power" {
        return &state.OpticalPower
    }
    return nil
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-power"] = &state.OpticalPower
    return children
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-attenuation"] = state.TargetAttenuation
    leafs["attenuation"] = state.Attenuation
    return leafs
}

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetBundleName() string { return "openconfig" }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetYangName() string { return "state" }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetParent() types.Entity { return state.parent }

func (state *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State) GetParentYangName() string { return "line-secondary-out" }

// Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower
// The optical output power of this port in units of
// 0.01dBm. Optical output power represents the signal
// traversing from the module to an external destination. The
// power is measured after any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetFilter() yfilter.YFilter { return opticalPower.YFilter }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) SetFilter(yf yfilter.YFilter) { opticalPower.YFilter = yf }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetSegmentPath() string {
    return "optical-power"
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = opticalPower.Instant
    leafs["avg"] = opticalPower.Avg
    leafs["min"] = opticalPower.Min
    leafs["max"] = opticalPower.Max
    return leafs
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetBundleName() string { return "openconfig" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetYangName() string { return "optical-power" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) SetParent(parent types.Entity) { opticalPower.parent = parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetParent() types.Entity { return opticalPower.parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_LineSecondaryOut_State_OpticalPower) GetParentYangName() string { return "state" }

// Aps_ApsModules_ApsModule_Ports_CommonIn
// Container for information related to the line common
// input port
type Aps_ApsModules_ApsModule_Ports_CommonIn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the line common input port.
    Config Aps_ApsModules_ApsModule_Ports_CommonIn_Config

    // State data for the line common input port.
    State Aps_ApsModules_ApsModule_Ports_CommonIn_State
}

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetFilter() yfilter.YFilter { return commonIn.YFilter }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) SetFilter(yf yfilter.YFilter) { commonIn.YFilter = yf }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetSegmentPath() string {
    return "common-in"
}

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &commonIn.Config
    }
    if childYangName == "state" {
        return &commonIn.State
    }
    return nil
}

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &commonIn.Config
    children["state"] = &commonIn.State
    return children
}

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetBundleName() string { return "openconfig" }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetYangName() string { return "common-in" }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) SetParent(parent types.Entity) { commonIn.parent = parent }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetParent() types.Entity { return commonIn.parent }

func (commonIn *Aps_ApsModules_ApsModule_Ports_CommonIn) GetParentYangName() string { return "ports" }

// Aps_ApsModules_ApsModule_Ports_CommonIn_Config
// Configuration data for the line common input port
type Aps_ApsModules_ApsModule_Ports_CommonIn_Config struct {
    parent types.Entity
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

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "target-attenuation" { return "TargetAttenuation" }
    return ""
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetSegmentPath() string {
    return "config"
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["target-attenuation"] = config.TargetAttenuation
    return leafs
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetBundleName() string { return "openconfig" }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetYangName() string { return "config" }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetParent() types.Entity { return config.parent }

func (config *Aps_ApsModules_ApsModule_Ports_CommonIn_Config) GetParentYangName() string { return "common-in" }

// Aps_ApsModules_ApsModule_Ports_CommonIn_State
// State data for the line common input port
type Aps_ApsModules_ApsModule_Ports_CommonIn_State struct {
    parent types.Entity
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

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "target-attenuation" { return "TargetAttenuation" }
    if yname == "attenuation" { return "Attenuation" }
    if yname == "optical-power" { return "OpticalPower" }
    return ""
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetSegmentPath() string {
    return "state"
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-power" {
        return &state.OpticalPower
    }
    return nil
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-power"] = &state.OpticalPower
    return children
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["target-attenuation"] = state.TargetAttenuation
    leafs["attenuation"] = state.Attenuation
    return leafs
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetBundleName() string { return "openconfig" }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetYangName() string { return "state" }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetParent() types.Entity { return state.parent }

func (state *Aps_ApsModules_ApsModule_Ports_CommonIn_State) GetParentYangName() string { return "common-in" }

// Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower
// The optical input power of this port in units of
// 0.01dBm. Optical input power represents the signal
// traversing from an external destination into the module.
// The power is measured before any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetFilter() yfilter.YFilter { return opticalPower.YFilter }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) SetFilter(yf yfilter.YFilter) { opticalPower.YFilter = yf }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetSegmentPath() string {
    return "optical-power"
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = opticalPower.Instant
    leafs["avg"] = opticalPower.Avg
    leafs["min"] = opticalPower.Min
    leafs["max"] = opticalPower.Max
    return leafs
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetBundleName() string { return "openconfig" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetYangName() string { return "optical-power" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) SetParent(parent types.Entity) { opticalPower.parent = parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetParent() types.Entity { return opticalPower.parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonIn_State_OpticalPower) GetParentYangName() string { return "state" }

// Aps_ApsModules_ApsModule_Ports_CommonOutput
// Container for information related to the line common
// output port
type Aps_ApsModules_ApsModule_Ports_CommonOutput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the line common output port.
    Config Aps_ApsModules_ApsModule_Ports_CommonOutput_Config

    // State data for the line common output port.
    State Aps_ApsModules_ApsModule_Ports_CommonOutput_State
}

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetFilter() yfilter.YFilter { return commonOutput.YFilter }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) SetFilter(yf yfilter.YFilter) { commonOutput.YFilter = yf }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetSegmentPath() string {
    return "common-output"
}

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &commonOutput.Config
    }
    if childYangName == "state" {
        return &commonOutput.State
    }
    return nil
}

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &commonOutput.Config
    children["state"] = &commonOutput.State
    return children
}

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetBundleName() string { return "openconfig" }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetYangName() string { return "common-output" }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) SetParent(parent types.Entity) { commonOutput.parent = parent }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetParent() types.Entity { return commonOutput.parent }

func (commonOutput *Aps_ApsModules_ApsModule_Ports_CommonOutput) GetParentYangName() string { return "ports" }

// Aps_ApsModules_ApsModule_Ports_CommonOutput_Config
// Configuration data for the line common output port
type Aps_ApsModules_ApsModule_Ports_CommonOutput_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Target attenuation of the variable optical attenuator associated with the
    // port in increments of 0.01 dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetAttenuation interface{}
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetGoName(yname string) string {
    if yname == "target-attenuation" { return "TargetAttenuation" }
    return ""
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetSegmentPath() string {
    return "config"
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-attenuation"] = config.TargetAttenuation
    return leafs
}

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetBundleName() string { return "openconfig" }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetYangName() string { return "config" }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetParent() types.Entity { return config.parent }

func (config *Aps_ApsModules_ApsModule_Ports_CommonOutput_Config) GetParentYangName() string { return "common-output" }

// Aps_ApsModules_ApsModule_Ports_CommonOutput_State
// State data for the line common output port
type Aps_ApsModules_ApsModule_Ports_CommonOutput_State struct {
    parent types.Entity
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

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetGoName(yname string) string {
    if yname == "target-attenuation" { return "TargetAttenuation" }
    if yname == "attenuation" { return "Attenuation" }
    if yname == "optical-power" { return "OpticalPower" }
    return ""
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetSegmentPath() string {
    return "state"
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-power" {
        return &state.OpticalPower
    }
    return nil
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-power"] = &state.OpticalPower
    return children
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-attenuation"] = state.TargetAttenuation
    leafs["attenuation"] = state.Attenuation
    return leafs
}

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetBundleName() string { return "openconfig" }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetYangName() string { return "state" }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetParent() types.Entity { return state.parent }

func (state *Aps_ApsModules_ApsModule_Ports_CommonOutput_State) GetParentYangName() string { return "common-output" }

// Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower
// The optical output power of this port in units of
// 0.01dBm. Optical output power represents the signal
// traversing from the module to an external destination. The
// power is measured after any attenuation. If avg/min/max
// statistics are not supported, the target is expected to
// just supply the instant value
type Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower struct {
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

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetFilter() yfilter.YFilter { return opticalPower.YFilter }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) SetFilter(yf yfilter.YFilter) { opticalPower.YFilter = yf }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetSegmentPath() string {
    return "optical-power"
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = opticalPower.Instant
    leafs["avg"] = opticalPower.Avg
    leafs["min"] = opticalPower.Min
    leafs["max"] = opticalPower.Max
    return leafs
}

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetBundleName() string { return "openconfig" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetYangName() string { return "optical-power" }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) SetParent(parent types.Entity) { opticalPower.parent = parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetParent() types.Entity { return opticalPower.parent }

func (opticalPower *Aps_ApsModules_ApsModule_Ports_CommonOutput_State_OpticalPower) GetParentYangName() string { return "state" }

