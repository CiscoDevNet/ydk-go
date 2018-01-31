// This model describes configuration and operational state data
// for optical amplifiers, deployed as part of a transport
// line system.
package optical_amplifier

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package optical_amplifier"))
    ydk.RegisterEntity("{http://openconfig.net/yang/optical-amplfier optical-amplifier}", reflect.TypeOf(OpticalAmplifier{}))
    ydk.RegisterEntity("openconfig-optical-amplifier:optical-amplifier", reflect.TypeOf(OpticalAmplifier{}))
}

type LOWGAINRANGE struct {
}

func (id LOWGAINRANGE) String() string {
	return "openconfig-optical-amplifier:LOW_GAIN_RANGE"
}

type BACKWARDRAMAN struct {
}

func (id BACKWARDRAMAN) String() string {
	return "openconfig-optical-amplifier:BACKWARD_RAMAN"
}

type CONSTANTGAIN struct {
}

func (id CONSTANTGAIN) String() string {
	return "openconfig-optical-amplifier:CONSTANT_GAIN"
}

type FIXEDGAINRANGE struct {
}

func (id FIXEDGAINRANGE) String() string {
	return "openconfig-optical-amplifier:FIXED_GAIN_RANGE"
}

type MIDGAINRANGE struct {
}

func (id MIDGAINRANGE) String() string {
	return "openconfig-optical-amplifier:MID_GAIN_RANGE"
}

type HIGHGAINRANGE struct {
}

func (id HIGHGAINRANGE) String() string {
	return "openconfig-optical-amplifier:HIGH_GAIN_RANGE"
}

type HYBRID struct {
}

func (id HYBRID) String() string {
	return "openconfig-optical-amplifier:HYBRID"
}

type FORWARDRAMAN struct {
}

func (id FORWARDRAMAN) String() string {
	return "openconfig-optical-amplifier:FORWARD_RAMAN"
}

type EDFA struct {
}

func (id EDFA) String() string {
	return "openconfig-optical-amplifier:EDFA"
}

type OPTICALAMPLIFIERTYPE struct {
}

func (id OPTICALAMPLIFIERTYPE) String() string {
	return "openconfig-optical-amplifier:OPTICAL_AMPLIFIER_TYPE"
}

type OPTICALAMPLIFIERMODE struct {
}

func (id OPTICALAMPLIFIERMODE) String() string {
	return "openconfig-optical-amplifier:OPTICAL_AMPLIFIER_MODE"
}

type GAINRANGE struct {
}

func (id GAINRANGE) String() string {
	return "openconfig-optical-amplifier:GAIN_RANGE"
}

type CONSTANTPOWER struct {
}

func (id CONSTANTPOWER) String() string {
	return "openconfig-optical-amplifier:CONSTANT_POWER"
}

// OpticalAmplifier
// Enclosing container for amplifiers and supervisory channels
type OpticalAmplifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enclosing container for list of amplifiers.
    Amplifiers OpticalAmplifier_Amplifiers

    // Enclosing container for list of supervisory channels.
    SupervisoryChannels OpticalAmplifier_SupervisoryChannels
}

func (opticalAmplifier *OpticalAmplifier) GetFilter() yfilter.YFilter { return opticalAmplifier.YFilter }

func (opticalAmplifier *OpticalAmplifier) SetFilter(yf yfilter.YFilter) { opticalAmplifier.YFilter = yf }

func (opticalAmplifier *OpticalAmplifier) GetGoName(yname string) string {
    if yname == "amplifiers" { return "Amplifiers" }
    if yname == "supervisory-channels" { return "SupervisoryChannels" }
    return ""
}

func (opticalAmplifier *OpticalAmplifier) GetSegmentPath() string {
    return "openconfig-optical-amplifier:optical-amplifier"
}

func (opticalAmplifier *OpticalAmplifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "amplifiers" {
        return &opticalAmplifier.Amplifiers
    }
    if childYangName == "supervisory-channels" {
        return &opticalAmplifier.SupervisoryChannels
    }
    return nil
}

func (opticalAmplifier *OpticalAmplifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["amplifiers"] = &opticalAmplifier.Amplifiers
    children["supervisory-channels"] = &opticalAmplifier.SupervisoryChannels
    return children
}

func (opticalAmplifier *OpticalAmplifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalAmplifier *OpticalAmplifier) GetBundleName() string { return "openconfig" }

func (opticalAmplifier *OpticalAmplifier) GetYangName() string { return "optical-amplifier" }

func (opticalAmplifier *OpticalAmplifier) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalAmplifier *OpticalAmplifier) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalAmplifier *OpticalAmplifier) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalAmplifier *OpticalAmplifier) SetParent(parent types.Entity) { opticalAmplifier.parent = parent }

func (opticalAmplifier *OpticalAmplifier) GetParent() types.Entity { return opticalAmplifier.parent }

func (opticalAmplifier *OpticalAmplifier) GetParentYangName() string { return "openconfig-optical-amplifier" }

// OpticalAmplifier_Amplifiers
// Enclosing container for list of amplifiers
type OpticalAmplifier_Amplifiers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of optical amplifiers present in the device. The type is slice of
    // OpticalAmplifier_Amplifiers_Amplifier.
    Amplifier []OpticalAmplifier_Amplifiers_Amplifier
}

func (amplifiers *OpticalAmplifier_Amplifiers) GetFilter() yfilter.YFilter { return amplifiers.YFilter }

func (amplifiers *OpticalAmplifier_Amplifiers) SetFilter(yf yfilter.YFilter) { amplifiers.YFilter = yf }

func (amplifiers *OpticalAmplifier_Amplifiers) GetGoName(yname string) string {
    if yname == "amplifier" { return "Amplifier" }
    return ""
}

func (amplifiers *OpticalAmplifier_Amplifiers) GetSegmentPath() string {
    return "amplifiers"
}

func (amplifiers *OpticalAmplifier_Amplifiers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "amplifier" {
        for _, c := range amplifiers.Amplifier {
            if amplifiers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalAmplifier_Amplifiers_Amplifier{}
        amplifiers.Amplifier = append(amplifiers.Amplifier, child)
        return &amplifiers.Amplifier[len(amplifiers.Amplifier)-1]
    }
    return nil
}

func (amplifiers *OpticalAmplifier_Amplifiers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range amplifiers.Amplifier {
        children[amplifiers.Amplifier[i].GetSegmentPath()] = &amplifiers.Amplifier[i]
    }
    return children
}

func (amplifiers *OpticalAmplifier_Amplifiers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (amplifiers *OpticalAmplifier_Amplifiers) GetBundleName() string { return "openconfig" }

func (amplifiers *OpticalAmplifier_Amplifiers) GetYangName() string { return "amplifiers" }

func (amplifiers *OpticalAmplifier_Amplifiers) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (amplifiers *OpticalAmplifier_Amplifiers) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (amplifiers *OpticalAmplifier_Amplifiers) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (amplifiers *OpticalAmplifier_Amplifiers) SetParent(parent types.Entity) { amplifiers.parent = parent }

func (amplifiers *OpticalAmplifier_Amplifiers) GetParent() types.Entity { return amplifiers.parent }

func (amplifiers *OpticalAmplifier_Amplifiers) GetParentYangName() string { return "optical-amplifier" }

// OpticalAmplifier_Amplifiers_Amplifier
// List of optical amplifiers present in the device
type OpticalAmplifier_Amplifiers_Amplifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the name of the amplifier. The type
    // is string. Refers to
    // optical_amplifier.OpticalAmplifier_Amplifiers_Amplifier_Config_Name
    Name interface{}

    // Configuration data for the amplifier.
    Config OpticalAmplifier_Amplifiers_Amplifier_Config

    // Operational state data for the amplifier.
    State OpticalAmplifier_Amplifiers_Amplifier_State
}

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetFilter() yfilter.YFilter { return amplifier.YFilter }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) SetFilter(yf yfilter.YFilter) { amplifier.YFilter = yf }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetSegmentPath() string {
    return "amplifier" + "[name='" + fmt.Sprintf("%v", amplifier.Name) + "']"
}

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &amplifier.Config
    }
    if childYangName == "state" {
        return &amplifier.State
    }
    return nil
}

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &amplifier.Config
    children["state"] = &amplifier.State
    return children
}

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = amplifier.Name
    return leafs
}

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetBundleName() string { return "openconfig" }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetYangName() string { return "amplifier" }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) SetParent(parent types.Entity) { amplifier.parent = parent }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetParent() types.Entity { return amplifier.parent }

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetParentYangName() string { return "amplifiers" }

// OpticalAmplifier_Amplifiers_Amplifier_Config
// Configuration data for the amplifier
type OpticalAmplifier_Amplifiers_Amplifier_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // User-defined name assigned to identify a specific amplifier in the device.
    // The type is string.
    Name interface{}

    // Type of the amplifier. The type is one of the following:
    // BACKWARDRAMANHYBRIDFORWARDRAMANEDFA.
    Type interface{}

    // Positive gain applied by the amplifier. The type is string with range:
    // 0..92233720368547758.07. Units are dB.
    TargetGain interface{}

    // Gain tilt control. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetGainTilt interface{}

    // Selected gain range.  The gain range is a platform-defined value indicating
    // the switched gain amplifier setting. The type is one of the following:
    // LOWGAINRANGEFIXEDGAINRANGEMIDGAINRANGEHIGHGAINRANGE.
    GainRange interface{}

    // The operating mode of the amplifier. The type is one of the following:
    // CONSTANTGAINCONSTANTPOWER.
    AmpMode interface{}

    // Output optical power of the amplifier. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    TargetOutputPower interface{}

    // Turns power on / off to the amplifiers gain module. The type is bool.
    Enabled interface{}
}

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "target-gain" { return "TargetGain" }
    if yname == "target-gain-tilt" { return "TargetGainTilt" }
    if yname == "gain-range" { return "GainRange" }
    if yname == "amp-mode" { return "AmpMode" }
    if yname == "target-output-power" { return "TargetOutputPower" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetSegmentPath() string {
    return "config"
}

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["type"] = config.Type
    leafs["target-gain"] = config.TargetGain
    leafs["target-gain-tilt"] = config.TargetGainTilt
    leafs["gain-range"] = config.GainRange
    leafs["amp-mode"] = config.AmpMode
    leafs["target-output-power"] = config.TargetOutputPower
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetBundleName() string { return "openconfig" }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetYangName() string { return "config" }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetParent() types.Entity { return config.parent }

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetParentYangName() string { return "amplifier" }

// OpticalAmplifier_Amplifiers_Amplifier_State
// Operational state data for the amplifier
type OpticalAmplifier_Amplifiers_Amplifier_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // User-defined name assigned to identify a specific amplifier in the device.
    // The type is string.
    Name interface{}

    // Type of the amplifier. The type is one of the following:
    // BACKWARDRAMANHYBRIDFORWARDRAMANEDFA.
    Type interface{}

    // Positive gain applied by the amplifier. The type is string with range:
    // 0..92233720368547758.07. Units are dB.
    TargetGain interface{}

    // Gain tilt control. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    TargetGainTilt interface{}

    // Selected gain range.  The gain range is a platform-defined value indicating
    // the switched gain amplifier setting. The type is one of the following:
    // LOWGAINRANGEFIXEDGAINRANGEMIDGAINRANGEHIGHGAINRANGE.
    GainRange interface{}

    // The operating mode of the amplifier. The type is one of the following:
    // CONSTANTGAINCONSTANTPOWER.
    AmpMode interface{}

    // Output optical power of the amplifier. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dBm.
    TargetOutputPower interface{}

    // Turns power on / off to the amplifiers gain module. The type is bool.
    Enabled interface{}

    // Reference to system-supplied name of the amplifier ingress port. This leaf
    // is only valid for ports of type INGRESS. The type is string. Refers to
    // platform.Components_Component_Name
    IngressPort interface{}

    // Reference to system-supplied name of the amplifier egress port. This leaf
    // is only valid for ports of type EGRESS. The type is string. Refers to
    // platform.Components_Component_Name
    EgressPort interface{}

    // The actual gain applied by the amplifier in units of 0.01dB. If avg/min/max
    // statistics are not supported, just supply the instant value.
    ActualGain OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain

    // The actual tilt applied by the amplifier in units of 0.01dB. If avg/min/max
    // statistics are not supported, just supply the instant value.
    ActualGainTilt OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt

    // The total input optical power of this port in units of 0.01dBm. If
    // avg/min/max statistics are not supported, just supply the instant value.
    InputPowerTotal OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal

    // The C band (consisting of approximately 191 to 195 THz or 1530nm to 1565
    // nm) input optical power of this port in units of 0.01dBm. If avg/min/max
    // statistics are not supported, just supply the instant value.
    InputPowerCBand OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand

    // The L band (consisting of approximately 184 to 191 THz or 1565 to 1625 nm)
    // input optical power of this port in units of 0.01dBm. If avg/min/max
    // statistics are not supported, just supply the instant value.
    InputPowerLBand OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand

    // The total output optical power of this port in units of 0.01dBm. If
    // avg/min/max statistics are not supported, just supply the instant value.
    OutputPowerTotal OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal

    // The C band (consisting of approximately 191 to 195 THz or 1530nm to 1565
    // nm)output optical power of this port in units of 0.01dBm. If avg/min/max
    // statistics are not supported, just supply the instant value.
    OutputPowerCBand OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand

    // The L band (consisting of approximately 184 to 191 THz or 1565 to 1625
    // nm)output optical power of this port in units of 0.01dBm. If avg/min/max
    // statistics are not supported, just supply the instant value.
    OutputPowerLBand OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand

    // The current applied by the system to the transmit laser to achieve the
    // output power. The current is expressed in mA with up to two decimal
    // precision. If avg/min/max statistics are not supported, just supply the
    // instant value.
    LaserBiasCurrent OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent

    // The optical return loss (ORL) is the ratio of the light reflected back into
    // the port to the light launched out of the port. ORL is in units of 0.01dBm.
    // If avg/min/max statistics are not supported, just supply the instant value.
    OpticalReturnLoss OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss
}

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "target-gain" { return "TargetGain" }
    if yname == "target-gain-tilt" { return "TargetGainTilt" }
    if yname == "gain-range" { return "GainRange" }
    if yname == "amp-mode" { return "AmpMode" }
    if yname == "target-output-power" { return "TargetOutputPower" }
    if yname == "enabled" { return "Enabled" }
    if yname == "ingress-port" { return "IngressPort" }
    if yname == "egress-port" { return "EgressPort" }
    if yname == "actual-gain" { return "ActualGain" }
    if yname == "actual-gain-tilt" { return "ActualGainTilt" }
    if yname == "input-power-total" { return "InputPowerTotal" }
    if yname == "input-power-c-band" { return "InputPowerCBand" }
    if yname == "input-power-l-band" { return "InputPowerLBand" }
    if yname == "output-power-total" { return "OutputPowerTotal" }
    if yname == "output-power-c-band" { return "OutputPowerCBand" }
    if yname == "output-power-l-band" { return "OutputPowerLBand" }
    if yname == "laser-bias-current" { return "LaserBiasCurrent" }
    if yname == "optical-return-loss" { return "OpticalReturnLoss" }
    return ""
}

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetSegmentPath() string {
    return "state"
}

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "actual-gain" {
        return &state.ActualGain
    }
    if childYangName == "actual-gain-tilt" {
        return &state.ActualGainTilt
    }
    if childYangName == "input-power-total" {
        return &state.InputPowerTotal
    }
    if childYangName == "input-power-c-band" {
        return &state.InputPowerCBand
    }
    if childYangName == "input-power-l-band" {
        return &state.InputPowerLBand
    }
    if childYangName == "output-power-total" {
        return &state.OutputPowerTotal
    }
    if childYangName == "output-power-c-band" {
        return &state.OutputPowerCBand
    }
    if childYangName == "output-power-l-band" {
        return &state.OutputPowerLBand
    }
    if childYangName == "laser-bias-current" {
        return &state.LaserBiasCurrent
    }
    if childYangName == "optical-return-loss" {
        return &state.OpticalReturnLoss
    }
    return nil
}

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["actual-gain"] = &state.ActualGain
    children["actual-gain-tilt"] = &state.ActualGainTilt
    children["input-power-total"] = &state.InputPowerTotal
    children["input-power-c-band"] = &state.InputPowerCBand
    children["input-power-l-band"] = &state.InputPowerLBand
    children["output-power-total"] = &state.OutputPowerTotal
    children["output-power-c-band"] = &state.OutputPowerCBand
    children["output-power-l-band"] = &state.OutputPowerLBand
    children["laser-bias-current"] = &state.LaserBiasCurrent
    children["optical-return-loss"] = &state.OpticalReturnLoss
    return children
}

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["type"] = state.Type
    leafs["target-gain"] = state.TargetGain
    leafs["target-gain-tilt"] = state.TargetGainTilt
    leafs["gain-range"] = state.GainRange
    leafs["amp-mode"] = state.AmpMode
    leafs["target-output-power"] = state.TargetOutputPower
    leafs["enabled"] = state.Enabled
    leafs["ingress-port"] = state.IngressPort
    leafs["egress-port"] = state.EgressPort
    return leafs
}

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetBundleName() string { return "openconfig" }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetYangName() string { return "state" }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetParent() types.Entity { return state.parent }

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetParentYangName() string { return "amplifier" }

// OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain
// The actual gain applied by the amplifier in units of
// 0.01dB. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -92233720368547758.08..92233720368547758.07.
    // Units are dB.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    Min interface{}

    // The maximum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    Max interface{}
}

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetFilter() yfilter.YFilter { return actualGain.YFilter }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) SetFilter(yf yfilter.YFilter) { actualGain.YFilter = yf }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetSegmentPath() string {
    return "actual-gain"
}

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = actualGain.Instant
    leafs["avg"] = actualGain.Avg
    leafs["min"] = actualGain.Min
    leafs["max"] = actualGain.Max
    return leafs
}

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetBundleName() string { return "openconfig" }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetYangName() string { return "actual-gain" }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) SetParent(parent types.Entity) { actualGain.parent = parent }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetParent() types.Entity { return actualGain.parent }

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt
// The actual tilt applied by the amplifier in units of
// 0.01dB. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are dB.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -92233720368547758.08..92233720368547758.07.
    // Units are dB.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    Min interface{}

    // The maximum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // dB.
    Max interface{}
}

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetFilter() yfilter.YFilter { return actualGainTilt.YFilter }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) SetFilter(yf yfilter.YFilter) { actualGainTilt.YFilter = yf }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetSegmentPath() string {
    return "actual-gain-tilt"
}

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = actualGainTilt.Instant
    leafs["avg"] = actualGainTilt.Avg
    leafs["min"] = actualGainTilt.Min
    leafs["max"] = actualGainTilt.Max
    return leafs
}

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetBundleName() string { return "openconfig" }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetYangName() string { return "actual-gain-tilt" }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) SetParent(parent types.Entity) { actualGainTilt.parent = parent }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetParent() types.Entity { return actualGainTilt.parent }

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal
// The total input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal struct {
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

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetFilter() yfilter.YFilter { return inputPowerTotal.YFilter }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) SetFilter(yf yfilter.YFilter) { inputPowerTotal.YFilter = yf }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetSegmentPath() string {
    return "input-power-total"
}

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = inputPowerTotal.Instant
    leafs["avg"] = inputPowerTotal.Avg
    leafs["min"] = inputPowerTotal.Min
    leafs["max"] = inputPowerTotal.Max
    return leafs
}

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetBundleName() string { return "openconfig" }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetYangName() string { return "input-power-total" }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) SetParent(parent types.Entity) { inputPowerTotal.parent = parent }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetParent() types.Entity { return inputPowerTotal.parent }

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand
// The C band (consisting of approximately 191 to 195 THz or
// 1530nm to 1565 nm) input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand struct {
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

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetFilter() yfilter.YFilter { return inputPowerCBand.YFilter }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) SetFilter(yf yfilter.YFilter) { inputPowerCBand.YFilter = yf }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetSegmentPath() string {
    return "input-power-c-band"
}

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = inputPowerCBand.Instant
    leafs["avg"] = inputPowerCBand.Avg
    leafs["min"] = inputPowerCBand.Min
    leafs["max"] = inputPowerCBand.Max
    return leafs
}

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetBundleName() string { return "openconfig" }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetYangName() string { return "input-power-c-band" }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) SetParent(parent types.Entity) { inputPowerCBand.parent = parent }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetParent() types.Entity { return inputPowerCBand.parent }

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand
// The L band (consisting of approximately 184 to 191 THz or
// 1565 to 1625 nm) input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand struct {
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

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetFilter() yfilter.YFilter { return inputPowerLBand.YFilter }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) SetFilter(yf yfilter.YFilter) { inputPowerLBand.YFilter = yf }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetSegmentPath() string {
    return "input-power-l-band"
}

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = inputPowerLBand.Instant
    leafs["avg"] = inputPowerLBand.Avg
    leafs["min"] = inputPowerLBand.Min
    leafs["max"] = inputPowerLBand.Max
    return leafs
}

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetBundleName() string { return "openconfig" }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetYangName() string { return "input-power-l-band" }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) SetParent(parent types.Entity) { inputPowerLBand.parent = parent }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetParent() types.Entity { return inputPowerLBand.parent }

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal
// The total output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal struct {
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

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetFilter() yfilter.YFilter { return outputPowerTotal.YFilter }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) SetFilter(yf yfilter.YFilter) { outputPowerTotal.YFilter = yf }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetSegmentPath() string {
    return "output-power-total"
}

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = outputPowerTotal.Instant
    leafs["avg"] = outputPowerTotal.Avg
    leafs["min"] = outputPowerTotal.Min
    leafs["max"] = outputPowerTotal.Max
    return leafs
}

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetBundleName() string { return "openconfig" }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetYangName() string { return "output-power-total" }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) SetParent(parent types.Entity) { outputPowerTotal.parent = parent }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetParent() types.Entity { return outputPowerTotal.parent }

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand
// The C band (consisting of approximately 191 to 195 THz or
// 1530nm to 1565 nm)output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand struct {
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

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetFilter() yfilter.YFilter { return outputPowerCBand.YFilter }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) SetFilter(yf yfilter.YFilter) { outputPowerCBand.YFilter = yf }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetSegmentPath() string {
    return "output-power-c-band"
}

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = outputPowerCBand.Instant
    leafs["avg"] = outputPowerCBand.Avg
    leafs["min"] = outputPowerCBand.Min
    leafs["max"] = outputPowerCBand.Max
    return leafs
}

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetBundleName() string { return "openconfig" }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetYangName() string { return "output-power-c-band" }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) SetParent(parent types.Entity) { outputPowerCBand.parent = parent }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetParent() types.Entity { return outputPowerCBand.parent }

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand
// The L band (consisting of approximately 184 to 191 THz or
// 1565 to 1625 nm)output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand struct {
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

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetFilter() yfilter.YFilter { return outputPowerLBand.YFilter }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) SetFilter(yf yfilter.YFilter) { outputPowerLBand.YFilter = yf }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetSegmentPath() string {
    return "output-power-l-band"
}

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = outputPowerLBand.Instant
    leafs["avg"] = outputPowerLBand.Avg
    leafs["min"] = outputPowerLBand.Min
    leafs["max"] = outputPowerLBand.Max
    return leafs
}

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetBundleName() string { return "openconfig" }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetYangName() string { return "output-power-l-band" }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) SetParent(parent types.Entity) { outputPowerLBand.parent = parent }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetParent() types.Entity { return outputPowerLBand.parent }

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power. The current is expressed in mA
// with up to two decimal precision. If avg/min/max statistics
// are not supported, just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are mA.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -92233720368547758.08..92233720368547758.07.
    // Units are mA.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // mA.
    Min interface{}

    // The maximum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // mA.
    Max interface{}
}

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetFilter() yfilter.YFilter { return laserBiasCurrent.YFilter }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) SetFilter(yf yfilter.YFilter) { laserBiasCurrent.YFilter = yf }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetSegmentPath() string {
    return "laser-bias-current"
}

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = laserBiasCurrent.Instant
    leafs["avg"] = laserBiasCurrent.Avg
    leafs["min"] = laserBiasCurrent.Min
    leafs["max"] = laserBiasCurrent.Max
    return leafs
}

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetBundleName() string { return "openconfig" }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetYangName() string { return "laser-bias-current" }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) SetParent(parent types.Entity) { laserBiasCurrent.parent = parent }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetParent() types.Entity { return laserBiasCurrent.parent }

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetParentYangName() string { return "state" }

// OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss
// The optical return loss (ORL) is the ratio of the light
// reflected back into the port to the light launched out of
// the port. ORL is in units of 0.01dBm. If avg/min/max
// statistics are not supported, just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss struct {
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

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetFilter() yfilter.YFilter { return opticalReturnLoss.YFilter }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) SetFilter(yf yfilter.YFilter) { opticalReturnLoss.YFilter = yf }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetSegmentPath() string {
    return "optical-return-loss"
}

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = opticalReturnLoss.Instant
    leafs["avg"] = opticalReturnLoss.Avg
    leafs["min"] = opticalReturnLoss.Min
    leafs["max"] = opticalReturnLoss.Max
    return leafs
}

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetBundleName() string { return "openconfig" }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetYangName() string { return "optical-return-loss" }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) SetParent(parent types.Entity) { opticalReturnLoss.parent = parent }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetParent() types.Entity { return opticalReturnLoss.parent }

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetParentYangName() string { return "state" }

// OpticalAmplifier_SupervisoryChannels
// Enclosing container for list of supervisory channels
type OpticalAmplifier_SupervisoryChannels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of supervisory channels. The type is slice of
    // OpticalAmplifier_SupervisoryChannels_SupervisoryChannel.
    SupervisoryChannel []OpticalAmplifier_SupervisoryChannels_SupervisoryChannel
}

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetFilter() yfilter.YFilter { return supervisoryChannels.YFilter }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) SetFilter(yf yfilter.YFilter) { supervisoryChannels.YFilter = yf }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetGoName(yname string) string {
    if yname == "supervisory-channel" { return "SupervisoryChannel" }
    return ""
}

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetSegmentPath() string {
    return "supervisory-channels"
}

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "supervisory-channel" {
        for _, c := range supervisoryChannels.SupervisoryChannel {
            if supervisoryChannels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalAmplifier_SupervisoryChannels_SupervisoryChannel{}
        supervisoryChannels.SupervisoryChannel = append(supervisoryChannels.SupervisoryChannel, child)
        return &supervisoryChannels.SupervisoryChannel[len(supervisoryChannels.SupervisoryChannel)-1]
    }
    return nil
}

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range supervisoryChannels.SupervisoryChannel {
        children[supervisoryChannels.SupervisoryChannel[i].GetSegmentPath()] = &supervisoryChannels.SupervisoryChannel[i]
    }
    return children
}

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetBundleName() string { return "openconfig" }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetYangName() string { return "supervisory-channels" }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) SetParent(parent types.Entity) { supervisoryChannels.parent = parent }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetParent() types.Entity { return supervisoryChannels.parent }

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetParentYangName() string { return "optical-amplifier" }

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel
// List of supervisory channels
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the interface of the supervisory
    // channel. The type is string. Refers to
    // optical_amplifier.OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config_Interface
    Interface interface{}

    // Configuration data for OSCs.
    Config OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config

    // Operational state data for OSCs.
    State OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State
}

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetFilter() yfilter.YFilter { return supervisoryChannel.YFilter }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) SetFilter(yf yfilter.YFilter) { supervisoryChannel.YFilter = yf }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetSegmentPath() string {
    return "supervisory-channel" + "[interface='" + fmt.Sprintf("%v", supervisoryChannel.Interface) + "']"
}

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &supervisoryChannel.Config
    }
    if childYangName == "state" {
        return &supervisoryChannel.State
    }
    return nil
}

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &supervisoryChannel.Config
    children["state"] = &supervisoryChannel.State
    return children
}

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = supervisoryChannel.Interface
    return leafs
}

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetBundleName() string { return "openconfig" }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetYangName() string { return "supervisory-channel" }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) SetParent(parent types.Entity) { supervisoryChannel.parent = parent }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetParent() types.Entity { return supervisoryChannel.parent }

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetParentYangName() string { return "supervisory-channels" }

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config
// Configuration data for OSCs
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of references to OSC interfaces. The type is slice of string. Refers
    // to interfaces.Interfaces_Interface_Name
    Interface []interface{}
}

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetSegmentPath() string {
    return "config"
}

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    return leafs
}

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetBundleName() string { return "openconfig" }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetYangName() string { return "config" }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetParent() types.Entity { return config.parent }

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetParentYangName() string { return "supervisory-channel" }

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State
// Operational state data for OSCs
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of references to OSC interfaces. The type is slice of string. Refers
    // to interfaces.Interfaces_Interface_Name
    Interface []interface{}

    // The input optical power of this port in units of 0.01dBm. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    InputPower OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower

    // The output optical power of this port in units of 0.01dBm. If avg/min/max
    // statistics are not supported, the target is expected to just supply the
    // instant value.
    OutputPower OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower

    // The current applied by the system to the transmit laser to achieve the
    // output power. The current is expressed in mA with up to one decimal
    // precision. If avg/min/max statistics are not supported, the target is
    // expected to just supply the instant value.
    LaserBiasCurrent OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent
}

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "input-power" { return "InputPower" }
    if yname == "output-power" { return "OutputPower" }
    if yname == "laser-bias-current" { return "LaserBiasCurrent" }
    return ""
}

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetSegmentPath() string {
    return "state"
}

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input-power" {
        return &state.InputPower
    }
    if childYangName == "output-power" {
        return &state.OutputPower
    }
    if childYangName == "laser-bias-current" {
        return &state.LaserBiasCurrent
    }
    return nil
}

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input-power"] = &state.InputPower
    children["output-power"] = &state.OutputPower
    children["laser-bias-current"] = &state.LaserBiasCurrent
    return children
}

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    return leafs
}

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetBundleName() string { return "openconfig" }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetYangName() string { return "state" }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetParent() types.Entity { return state.parent }

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetParentYangName() string { return "supervisory-channel" }

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower
// The input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// the target is expected to just supply the instant value
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower struct {
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

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetFilter() yfilter.YFilter { return inputPower.YFilter }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) SetFilter(yf yfilter.YFilter) { inputPower.YFilter = yf }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetSegmentPath() string {
    return "input-power"
}

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = inputPower.Instant
    leafs["avg"] = inputPower.Avg
    leafs["min"] = inputPower.Min
    leafs["max"] = inputPower.Max
    return leafs
}

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetBundleName() string { return "openconfig" }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetYangName() string { return "input-power" }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) SetParent(parent types.Entity) { inputPower.parent = parent }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetParent() types.Entity { return inputPower.parent }

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetParentYangName() string { return "state" }

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower
// The output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// the target is expected to just supply the instant value
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower struct {
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

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetFilter() yfilter.YFilter { return outputPower.YFilter }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) SetFilter(yf yfilter.YFilter) { outputPower.YFilter = yf }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetSegmentPath() string {
    return "output-power"
}

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = outputPower.Instant
    leafs["avg"] = outputPower.Avg
    leafs["min"] = outputPower.Min
    leafs["max"] = outputPower.Max
    return leafs
}

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetBundleName() string { return "openconfig" }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetYangName() string { return "output-power" }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) SetParent(parent types.Entity) { outputPower.parent = parent }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetParent() types.Entity { return outputPower.parent }

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetParentYangName() string { return "state" }

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power. The current is expressed in mA
// with up to one decimal precision. If avg/min/max statistics
// are not supported, the target is expected to just supply
// the instant value
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The instantaneous value of the statistic. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are mA.
    Instant interface{}

    // The arithmetic mean value of the statistic over the sampling period. The
    // type is string with range: -92233720368547758.08..92233720368547758.07.
    // Units are mA.
    Avg interface{}

    // The minimum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // mA.
    Min interface{}

    // The maximum value of the statistic over the sampling period. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // mA.
    Max interface{}
}

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetFilter() yfilter.YFilter { return laserBiasCurrent.YFilter }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) SetFilter(yf yfilter.YFilter) { laserBiasCurrent.YFilter = yf }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetSegmentPath() string {
    return "laser-bias-current"
}

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = laserBiasCurrent.Instant
    leafs["avg"] = laserBiasCurrent.Avg
    leafs["min"] = laserBiasCurrent.Min
    leafs["max"] = laserBiasCurrent.Max
    return leafs
}

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetBundleName() string { return "openconfig" }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetYangName() string { return "laser-bias-current" }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) SetParent(parent types.Entity) { laserBiasCurrent.parent = parent }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetParent() types.Entity { return laserBiasCurrent.parent }

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetParentYangName() string { return "state" }

