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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for list of amplifiers.
    Amplifiers OpticalAmplifier_Amplifiers

    // Enclosing container for list of supervisory channels.
    SupervisoryChannels OpticalAmplifier_SupervisoryChannels
}

func (opticalAmplifier *OpticalAmplifier) GetEntityData() *types.CommonEntityData {
    opticalAmplifier.EntityData.YFilter = opticalAmplifier.YFilter
    opticalAmplifier.EntityData.YangName = "optical-amplifier"
    opticalAmplifier.EntityData.BundleName = "openconfig"
    opticalAmplifier.EntityData.ParentYangName = "openconfig-optical-amplifier"
    opticalAmplifier.EntityData.SegmentPath = "openconfig-optical-amplifier:optical-amplifier"
    opticalAmplifier.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalAmplifier.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalAmplifier.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalAmplifier.EntityData.Children = types.NewOrderedMap()
    opticalAmplifier.EntityData.Children.Append("amplifiers", types.YChild{"Amplifiers", &opticalAmplifier.Amplifiers})
    opticalAmplifier.EntityData.Children.Append("supervisory-channels", types.YChild{"SupervisoryChannels", &opticalAmplifier.SupervisoryChannels})
    opticalAmplifier.EntityData.Leafs = types.NewOrderedMap()

    opticalAmplifier.EntityData.YListKeys = []string {}

    return &(opticalAmplifier.EntityData)
}

// OpticalAmplifier_Amplifiers
// Enclosing container for list of amplifiers
type OpticalAmplifier_Amplifiers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of optical amplifiers present in the device. The type is slice of
    // OpticalAmplifier_Amplifiers_Amplifier.
    Amplifier []*OpticalAmplifier_Amplifiers_Amplifier
}

func (amplifiers *OpticalAmplifier_Amplifiers) GetEntityData() *types.CommonEntityData {
    amplifiers.EntityData.YFilter = amplifiers.YFilter
    amplifiers.EntityData.YangName = "amplifiers"
    amplifiers.EntityData.BundleName = "openconfig"
    amplifiers.EntityData.ParentYangName = "optical-amplifier"
    amplifiers.EntityData.SegmentPath = "amplifiers"
    amplifiers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    amplifiers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    amplifiers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    amplifiers.EntityData.Children = types.NewOrderedMap()
    amplifiers.EntityData.Children.Append("amplifier", types.YChild{"Amplifier", nil})
    for i := range amplifiers.Amplifier {
        amplifiers.EntityData.Children.Append(types.GetSegmentPath(amplifiers.Amplifier[i]), types.YChild{"Amplifier", amplifiers.Amplifier[i]})
    }
    amplifiers.EntityData.Leafs = types.NewOrderedMap()

    amplifiers.EntityData.YListKeys = []string {}

    return &(amplifiers.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier
// List of optical amplifiers present in the device
type OpticalAmplifier_Amplifiers_Amplifier struct {
    EntityData types.CommonEntityData
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

func (amplifier *OpticalAmplifier_Amplifiers_Amplifier) GetEntityData() *types.CommonEntityData {
    amplifier.EntityData.YFilter = amplifier.YFilter
    amplifier.EntityData.YangName = "amplifier"
    amplifier.EntityData.BundleName = "openconfig"
    amplifier.EntityData.ParentYangName = "amplifiers"
    amplifier.EntityData.SegmentPath = "amplifier" + types.AddKeyToken(amplifier.Name, "name")
    amplifier.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    amplifier.EntityData.NamespaceTable = openconfig.GetNamespaces()
    amplifier.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    amplifier.EntityData.Children = types.NewOrderedMap()
    amplifier.EntityData.Children.Append("config", types.YChild{"Config", &amplifier.Config})
    amplifier.EntityData.Children.Append("state", types.YChild{"State", &amplifier.State})
    amplifier.EntityData.Leafs = types.NewOrderedMap()
    amplifier.EntityData.Leafs.Append("name", types.YLeaf{"Name", amplifier.Name})

    amplifier.EntityData.YListKeys = []string {"Name"}

    return &(amplifier.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_Config
// Configuration data for the amplifier
type OpticalAmplifier_Amplifiers_Amplifier_Config struct {
    EntityData types.CommonEntityData
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

func (config *OpticalAmplifier_Amplifiers_Amplifier_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "amplifier"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("type", types.YLeaf{"Type", config.Type})
    config.EntityData.Leafs.Append("target-gain", types.YLeaf{"TargetGain", config.TargetGain})
    config.EntityData.Leafs.Append("target-gain-tilt", types.YLeaf{"TargetGainTilt", config.TargetGainTilt})
    config.EntityData.Leafs.Append("gain-range", types.YLeaf{"GainRange", config.GainRange})
    config.EntityData.Leafs.Append("amp-mode", types.YLeaf{"AmpMode", config.AmpMode})
    config.EntityData.Leafs.Append("target-output-power", types.YLeaf{"TargetOutputPower", config.TargetOutputPower})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State
// Operational state data for the amplifier
type OpticalAmplifier_Amplifiers_Amplifier_State struct {
    EntityData types.CommonEntityData
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

func (state *OpticalAmplifier_Amplifiers_Amplifier_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "amplifier"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("actual-gain", types.YChild{"ActualGain", &state.ActualGain})
    state.EntityData.Children.Append("actual-gain-tilt", types.YChild{"ActualGainTilt", &state.ActualGainTilt})
    state.EntityData.Children.Append("input-power-total", types.YChild{"InputPowerTotal", &state.InputPowerTotal})
    state.EntityData.Children.Append("input-power-c-band", types.YChild{"InputPowerCBand", &state.InputPowerCBand})
    state.EntityData.Children.Append("input-power-l-band", types.YChild{"InputPowerLBand", &state.InputPowerLBand})
    state.EntityData.Children.Append("output-power-total", types.YChild{"OutputPowerTotal", &state.OutputPowerTotal})
    state.EntityData.Children.Append("output-power-c-band", types.YChild{"OutputPowerCBand", &state.OutputPowerCBand})
    state.EntityData.Children.Append("output-power-l-band", types.YChild{"OutputPowerLBand", &state.OutputPowerLBand})
    state.EntityData.Children.Append("laser-bias-current", types.YChild{"LaserBiasCurrent", &state.LaserBiasCurrent})
    state.EntityData.Children.Append("optical-return-loss", types.YChild{"OpticalReturnLoss", &state.OpticalReturnLoss})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("target-gain", types.YLeaf{"TargetGain", state.TargetGain})
    state.EntityData.Leafs.Append("target-gain-tilt", types.YLeaf{"TargetGainTilt", state.TargetGainTilt})
    state.EntityData.Leafs.Append("gain-range", types.YLeaf{"GainRange", state.GainRange})
    state.EntityData.Leafs.Append("amp-mode", types.YLeaf{"AmpMode", state.AmpMode})
    state.EntityData.Leafs.Append("target-output-power", types.YLeaf{"TargetOutputPower", state.TargetOutputPower})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("ingress-port", types.YLeaf{"IngressPort", state.IngressPort})
    state.EntityData.Leafs.Append("egress-port", types.YLeaf{"EgressPort", state.EgressPort})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain
// The actual gain applied by the amplifier in units of
// 0.01dB. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain struct {
    EntityData types.CommonEntityData
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

func (actualGain *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGain) GetEntityData() *types.CommonEntityData {
    actualGain.EntityData.YFilter = actualGain.YFilter
    actualGain.EntityData.YangName = "actual-gain"
    actualGain.EntityData.BundleName = "openconfig"
    actualGain.EntityData.ParentYangName = "state"
    actualGain.EntityData.SegmentPath = "actual-gain"
    actualGain.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    actualGain.EntityData.NamespaceTable = openconfig.GetNamespaces()
    actualGain.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    actualGain.EntityData.Children = types.NewOrderedMap()
    actualGain.EntityData.Leafs = types.NewOrderedMap()
    actualGain.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", actualGain.Instant})
    actualGain.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", actualGain.Avg})
    actualGain.EntityData.Leafs.Append("min", types.YLeaf{"Min", actualGain.Min})
    actualGain.EntityData.Leafs.Append("max", types.YLeaf{"Max", actualGain.Max})

    actualGain.EntityData.YListKeys = []string {}

    return &(actualGain.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt
// The actual tilt applied by the amplifier in units of
// 0.01dB. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt struct {
    EntityData types.CommonEntityData
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

func (actualGainTilt *OpticalAmplifier_Amplifiers_Amplifier_State_ActualGainTilt) GetEntityData() *types.CommonEntityData {
    actualGainTilt.EntityData.YFilter = actualGainTilt.YFilter
    actualGainTilt.EntityData.YangName = "actual-gain-tilt"
    actualGainTilt.EntityData.BundleName = "openconfig"
    actualGainTilt.EntityData.ParentYangName = "state"
    actualGainTilt.EntityData.SegmentPath = "actual-gain-tilt"
    actualGainTilt.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    actualGainTilt.EntityData.NamespaceTable = openconfig.GetNamespaces()
    actualGainTilt.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    actualGainTilt.EntityData.Children = types.NewOrderedMap()
    actualGainTilt.EntityData.Leafs = types.NewOrderedMap()
    actualGainTilt.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", actualGainTilt.Instant})
    actualGainTilt.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", actualGainTilt.Avg})
    actualGainTilt.EntityData.Leafs.Append("min", types.YLeaf{"Min", actualGainTilt.Min})
    actualGainTilt.EntityData.Leafs.Append("max", types.YLeaf{"Max", actualGainTilt.Max})

    actualGainTilt.EntityData.YListKeys = []string {}

    return &(actualGainTilt.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal
// The total input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal struct {
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

func (inputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerTotal) GetEntityData() *types.CommonEntityData {
    inputPowerTotal.EntityData.YFilter = inputPowerTotal.YFilter
    inputPowerTotal.EntityData.YangName = "input-power-total"
    inputPowerTotal.EntityData.BundleName = "openconfig"
    inputPowerTotal.EntityData.ParentYangName = "state"
    inputPowerTotal.EntityData.SegmentPath = "input-power-total"
    inputPowerTotal.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inputPowerTotal.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inputPowerTotal.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inputPowerTotal.EntityData.Children = types.NewOrderedMap()
    inputPowerTotal.EntityData.Leafs = types.NewOrderedMap()
    inputPowerTotal.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPowerTotal.Instant})
    inputPowerTotal.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPowerTotal.Avg})
    inputPowerTotal.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPowerTotal.Min})
    inputPowerTotal.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPowerTotal.Max})

    inputPowerTotal.EntityData.YListKeys = []string {}

    return &(inputPowerTotal.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand
// The C band (consisting of approximately 191 to 195 THz or
// 1530nm to 1565 nm) input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand struct {
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

func (inputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerCBand) GetEntityData() *types.CommonEntityData {
    inputPowerCBand.EntityData.YFilter = inputPowerCBand.YFilter
    inputPowerCBand.EntityData.YangName = "input-power-c-band"
    inputPowerCBand.EntityData.BundleName = "openconfig"
    inputPowerCBand.EntityData.ParentYangName = "state"
    inputPowerCBand.EntityData.SegmentPath = "input-power-c-band"
    inputPowerCBand.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inputPowerCBand.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inputPowerCBand.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inputPowerCBand.EntityData.Children = types.NewOrderedMap()
    inputPowerCBand.EntityData.Leafs = types.NewOrderedMap()
    inputPowerCBand.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPowerCBand.Instant})
    inputPowerCBand.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPowerCBand.Avg})
    inputPowerCBand.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPowerCBand.Min})
    inputPowerCBand.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPowerCBand.Max})

    inputPowerCBand.EntityData.YListKeys = []string {}

    return &(inputPowerCBand.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand
// The L band (consisting of approximately 184 to 191 THz or
// 1565 to 1625 nm) input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand struct {
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

func (inputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_InputPowerLBand) GetEntityData() *types.CommonEntityData {
    inputPowerLBand.EntityData.YFilter = inputPowerLBand.YFilter
    inputPowerLBand.EntityData.YangName = "input-power-l-band"
    inputPowerLBand.EntityData.BundleName = "openconfig"
    inputPowerLBand.EntityData.ParentYangName = "state"
    inputPowerLBand.EntityData.SegmentPath = "input-power-l-band"
    inputPowerLBand.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inputPowerLBand.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inputPowerLBand.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inputPowerLBand.EntityData.Children = types.NewOrderedMap()
    inputPowerLBand.EntityData.Leafs = types.NewOrderedMap()
    inputPowerLBand.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", inputPowerLBand.Instant})
    inputPowerLBand.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", inputPowerLBand.Avg})
    inputPowerLBand.EntityData.Leafs.Append("min", types.YLeaf{"Min", inputPowerLBand.Min})
    inputPowerLBand.EntityData.Leafs.Append("max", types.YLeaf{"Max", inputPowerLBand.Max})

    inputPowerLBand.EntityData.YListKeys = []string {}

    return &(inputPowerLBand.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal
// The total output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal struct {
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

func (outputPowerTotal *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerTotal) GetEntityData() *types.CommonEntityData {
    outputPowerTotal.EntityData.YFilter = outputPowerTotal.YFilter
    outputPowerTotal.EntityData.YangName = "output-power-total"
    outputPowerTotal.EntityData.BundleName = "openconfig"
    outputPowerTotal.EntityData.ParentYangName = "state"
    outputPowerTotal.EntityData.SegmentPath = "output-power-total"
    outputPowerTotal.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    outputPowerTotal.EntityData.NamespaceTable = openconfig.GetNamespaces()
    outputPowerTotal.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    outputPowerTotal.EntityData.Children = types.NewOrderedMap()
    outputPowerTotal.EntityData.Leafs = types.NewOrderedMap()
    outputPowerTotal.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPowerTotal.Instant})
    outputPowerTotal.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPowerTotal.Avg})
    outputPowerTotal.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPowerTotal.Min})
    outputPowerTotal.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPowerTotal.Max})

    outputPowerTotal.EntityData.YListKeys = []string {}

    return &(outputPowerTotal.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand
// The C band (consisting of approximately 191 to 195 THz or
// 1530nm to 1565 nm)output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand struct {
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

func (outputPowerCBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerCBand) GetEntityData() *types.CommonEntityData {
    outputPowerCBand.EntityData.YFilter = outputPowerCBand.YFilter
    outputPowerCBand.EntityData.YangName = "output-power-c-band"
    outputPowerCBand.EntityData.BundleName = "openconfig"
    outputPowerCBand.EntityData.ParentYangName = "state"
    outputPowerCBand.EntityData.SegmentPath = "output-power-c-band"
    outputPowerCBand.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    outputPowerCBand.EntityData.NamespaceTable = openconfig.GetNamespaces()
    outputPowerCBand.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    outputPowerCBand.EntityData.Children = types.NewOrderedMap()
    outputPowerCBand.EntityData.Leafs = types.NewOrderedMap()
    outputPowerCBand.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPowerCBand.Instant})
    outputPowerCBand.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPowerCBand.Avg})
    outputPowerCBand.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPowerCBand.Min})
    outputPowerCBand.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPowerCBand.Max})

    outputPowerCBand.EntityData.YListKeys = []string {}

    return &(outputPowerCBand.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand
// The L band (consisting of approximately 184 to 191 THz or
// 1565 to 1625 nm)output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand struct {
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

func (outputPowerLBand *OpticalAmplifier_Amplifiers_Amplifier_State_OutputPowerLBand) GetEntityData() *types.CommonEntityData {
    outputPowerLBand.EntityData.YFilter = outputPowerLBand.YFilter
    outputPowerLBand.EntityData.YangName = "output-power-l-band"
    outputPowerLBand.EntityData.BundleName = "openconfig"
    outputPowerLBand.EntityData.ParentYangName = "state"
    outputPowerLBand.EntityData.SegmentPath = "output-power-l-band"
    outputPowerLBand.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    outputPowerLBand.EntityData.NamespaceTable = openconfig.GetNamespaces()
    outputPowerLBand.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    outputPowerLBand.EntityData.Children = types.NewOrderedMap()
    outputPowerLBand.EntityData.Leafs = types.NewOrderedMap()
    outputPowerLBand.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", outputPowerLBand.Instant})
    outputPowerLBand.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", outputPowerLBand.Avg})
    outputPowerLBand.EntityData.Leafs.Append("min", types.YLeaf{"Min", outputPowerLBand.Min})
    outputPowerLBand.EntityData.Leafs.Append("max", types.YLeaf{"Max", outputPowerLBand.Max})

    outputPowerLBand.EntityData.YListKeys = []string {}

    return &(outputPowerLBand.EntityData)
}

// OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power. The current is expressed in mA
// with up to two decimal precision. If avg/min/max statistics
// are not supported, just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent struct {
    EntityData types.CommonEntityData
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

func (laserBiasCurrent *OpticalAmplifier_Amplifiers_Amplifier_State_LaserBiasCurrent) GetEntityData() *types.CommonEntityData {
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

// OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss
// The optical return loss (ORL) is the ratio of the light
// reflected back into the port to the light launched out of
// the port. ORL is in units of 0.01dBm. If avg/min/max
// statistics are not supported, just supply the instant value
type OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss struct {
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

func (opticalReturnLoss *OpticalAmplifier_Amplifiers_Amplifier_State_OpticalReturnLoss) GetEntityData() *types.CommonEntityData {
    opticalReturnLoss.EntityData.YFilter = opticalReturnLoss.YFilter
    opticalReturnLoss.EntityData.YangName = "optical-return-loss"
    opticalReturnLoss.EntityData.BundleName = "openconfig"
    opticalReturnLoss.EntityData.ParentYangName = "state"
    opticalReturnLoss.EntityData.SegmentPath = "optical-return-loss"
    opticalReturnLoss.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    opticalReturnLoss.EntityData.NamespaceTable = openconfig.GetNamespaces()
    opticalReturnLoss.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    opticalReturnLoss.EntityData.Children = types.NewOrderedMap()
    opticalReturnLoss.EntityData.Leafs = types.NewOrderedMap()
    opticalReturnLoss.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", opticalReturnLoss.Instant})
    opticalReturnLoss.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", opticalReturnLoss.Avg})
    opticalReturnLoss.EntityData.Leafs.Append("min", types.YLeaf{"Min", opticalReturnLoss.Min})
    opticalReturnLoss.EntityData.Leafs.Append("max", types.YLeaf{"Max", opticalReturnLoss.Max})

    opticalReturnLoss.EntityData.YListKeys = []string {}

    return &(opticalReturnLoss.EntityData)
}

// OpticalAmplifier_SupervisoryChannels
// Enclosing container for list of supervisory channels
type OpticalAmplifier_SupervisoryChannels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of supervisory channels. The type is slice of
    // OpticalAmplifier_SupervisoryChannels_SupervisoryChannel.
    SupervisoryChannel []*OpticalAmplifier_SupervisoryChannels_SupervisoryChannel
}

func (supervisoryChannels *OpticalAmplifier_SupervisoryChannels) GetEntityData() *types.CommonEntityData {
    supervisoryChannels.EntityData.YFilter = supervisoryChannels.YFilter
    supervisoryChannels.EntityData.YangName = "supervisory-channels"
    supervisoryChannels.EntityData.BundleName = "openconfig"
    supervisoryChannels.EntityData.ParentYangName = "optical-amplifier"
    supervisoryChannels.EntityData.SegmentPath = "supervisory-channels"
    supervisoryChannels.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    supervisoryChannels.EntityData.NamespaceTable = openconfig.GetNamespaces()
    supervisoryChannels.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    supervisoryChannels.EntityData.Children = types.NewOrderedMap()
    supervisoryChannels.EntityData.Children.Append("supervisory-channel", types.YChild{"SupervisoryChannel", nil})
    for i := range supervisoryChannels.SupervisoryChannel {
        supervisoryChannels.EntityData.Children.Append(types.GetSegmentPath(supervisoryChannels.SupervisoryChannel[i]), types.YChild{"SupervisoryChannel", supervisoryChannels.SupervisoryChannel[i]})
    }
    supervisoryChannels.EntityData.Leafs = types.NewOrderedMap()

    supervisoryChannels.EntityData.YListKeys = []string {}

    return &(supervisoryChannels.EntityData)
}

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel
// List of supervisory channels
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel struct {
    EntityData types.CommonEntityData
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

func (supervisoryChannel *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel) GetEntityData() *types.CommonEntityData {
    supervisoryChannel.EntityData.YFilter = supervisoryChannel.YFilter
    supervisoryChannel.EntityData.YangName = "supervisory-channel"
    supervisoryChannel.EntityData.BundleName = "openconfig"
    supervisoryChannel.EntityData.ParentYangName = "supervisory-channels"
    supervisoryChannel.EntityData.SegmentPath = "supervisory-channel" + types.AddKeyToken(supervisoryChannel.Interface, "interface")
    supervisoryChannel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    supervisoryChannel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    supervisoryChannel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    supervisoryChannel.EntityData.Children = types.NewOrderedMap()
    supervisoryChannel.EntityData.Children.Append("config", types.YChild{"Config", &supervisoryChannel.Config})
    supervisoryChannel.EntityData.Children.Append("state", types.YChild{"State", &supervisoryChannel.State})
    supervisoryChannel.EntityData.Leafs = types.NewOrderedMap()
    supervisoryChannel.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", supervisoryChannel.Interface})

    supervisoryChannel.EntityData.YListKeys = []string {"Interface"}

    return &(supervisoryChannel.EntityData)
}

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config
// Configuration data for OSCs
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of references to OSC interfaces. The type is slice of string. Refers
    // to interfaces.Interfaces_Interface_Name
    Interface []interface{}
}

func (config *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "supervisory-channel"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State
// Operational state data for OSCs
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State struct {
    EntityData types.CommonEntityData
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

func (state *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "supervisory-channel"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("input-power", types.YChild{"InputPower", &state.InputPower})
    state.EntityData.Children.Append("output-power", types.YChild{"OutputPower", &state.OutputPower})
    state.EntityData.Children.Append("laser-bias-current", types.YChild{"LaserBiasCurrent", &state.LaserBiasCurrent})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower
// The input optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// the target is expected to just supply the instant value
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower struct {
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

func (inputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_InputPower) GetEntityData() *types.CommonEntityData {
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

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower
// The output optical power of this port in units
// of 0.01dBm. If avg/min/max statistics are not supported,
// the target is expected to just supply the instant value
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower struct {
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

func (outputPower *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_OutputPower) GetEntityData() *types.CommonEntityData {
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

// OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent
// The current applied by the system to the transmit laser to
// achieve the output power. The current is expressed in mA
// with up to one decimal precision. If avg/min/max statistics
// are not supported, the target is expected to just supply
// the instant value
type OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent struct {
    EntityData types.CommonEntityData
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

func (laserBiasCurrent *OpticalAmplifier_SupervisoryChannels_SupervisoryChannel_State_LaserBiasCurrent) GetEntityData() *types.CommonEntityData {
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

