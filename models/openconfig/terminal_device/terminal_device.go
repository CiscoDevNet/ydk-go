// This module describes a terminal optics device model for
// managing the terminal systems (client and line side) in a
// DWDM transport network.
// 
// Elements of the model:
// 
// physical port: corresponds to a physical, pluggable client
// port on the terminal device. Examples includes 10G, 40G, 100G
// (e.g., 10x10G, 4x25G or 1x100G) and 400G/1T in the future.
// Physical client ports will have associated operational state or
// PMs.
// 
// physical channel: a physical lane or channel in the
// physical client port.  Each physical client port has 1 or more
// channels. An example is 100GBASE-LR4 client physical port having
// 4x25G channels. Channels have their own optical PMs and can be
// monitored independently within a client physical port (e.g.,
// channel power).  Physical client channels are defined in the
// model as part of a physical client port, and are modeled
// primarily for reading their PMs.
// 
// logical channel: a logical grouping of logical grooming elements
// that may be assigned to subsequent grooming stages for
// multiplexing / de-multiplexing, or to an optical channel for
// line side transmission.  The logical channels can represent, for
// example, an ODU/OTU logical packing of the client
// data onto the line side.  Tributaries are similarly logical
// groupings of demand that can be represented in this structure and
// assigned to an optical channel.  Note that different types of
// logical channels may be present, each with their corresponding
// PMs.
// 
// optical channel:  corresponds to an optical carrier and is
// assigned a wavelength/frequency.  Optical channels have PMs
// such as power, BER, and operational mode.
// 
// Directionality:
// 
// To maintain simplicity in the model, the configuration is
// described from client-to-line direction.  The assumption is that
// equivalent reverse configuration is implicit, resulting in
// the same line-to-client configuration.
// 
// Physical layout:
// 
// The model does not assume a particular physical layout of client
// and line ports on the terminal device (e.g., such as number of
// ports per linecard, separate linecards for client and line ports,
// etc.).
package terminal_device

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package terminal_device"))
    ydk.RegisterEntity("{http://openconfig.net/yang/terminal-device terminal-device}", reflect.TypeOf(TerminalDevice{}))
    ydk.RegisterEntity("openconfig-terminal-device:terminal-device", reflect.TypeOf(TerminalDevice{}))
}

// TerminalDevice
// Top-level container for the terminal device
type TerminalDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for global terminal-device.
    Config TerminalDevice_Config

    // Operational state data for global terminal device.
    State TerminalDevice_State

    // Enclosing container the list of logical channels.
    LogicalChannels TerminalDevice_LogicalChannels

    // Enclosing container for list of operational modes.
    OperationalModes TerminalDevice_OperationalModes
}

func (terminalDevice *TerminalDevice) GetEntityData() *types.CommonEntityData {
    terminalDevice.EntityData.YFilter = terminalDevice.YFilter
    terminalDevice.EntityData.YangName = "terminal-device"
    terminalDevice.EntityData.BundleName = "openconfig"
    terminalDevice.EntityData.ParentYangName = "openconfig-terminal-device"
    terminalDevice.EntityData.SegmentPath = "openconfig-terminal-device:terminal-device"
    terminalDevice.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    terminalDevice.EntityData.NamespaceTable = openconfig.GetNamespaces()
    terminalDevice.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    terminalDevice.EntityData.Children = types.NewOrderedMap()
    terminalDevice.EntityData.Children.Append("config", types.YChild{"Config", &terminalDevice.Config})
    terminalDevice.EntityData.Children.Append("state", types.YChild{"State", &terminalDevice.State})
    terminalDevice.EntityData.Children.Append("logical-channels", types.YChild{"LogicalChannels", &terminalDevice.LogicalChannels})
    terminalDevice.EntityData.Children.Append("operational-modes", types.YChild{"OperationalModes", &terminalDevice.OperationalModes})
    terminalDevice.EntityData.Leafs = types.NewOrderedMap()

    terminalDevice.EntityData.YListKeys = []string {}

    return &(terminalDevice.EntityData)
}

// TerminalDevice_Config
// Configuration data for global terminal-device
type TerminalDevice_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *TerminalDevice_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "terminal-device"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TerminalDevice_State
// Operational state data for global terminal device
type TerminalDevice_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (state *TerminalDevice_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "terminal-device"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TerminalDevice_LogicalChannels
// Enclosing container the list of logical channels
type TerminalDevice_LogicalChannels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of logical channels. The type is slice of
    // TerminalDevice_LogicalChannels_Channel.
    Channel []*TerminalDevice_LogicalChannels_Channel
}

func (logicalChannels *TerminalDevice_LogicalChannels) GetEntityData() *types.CommonEntityData {
    logicalChannels.EntityData.YFilter = logicalChannels.YFilter
    logicalChannels.EntityData.YangName = "logical-channels"
    logicalChannels.EntityData.BundleName = "openconfig"
    logicalChannels.EntityData.ParentYangName = "terminal-device"
    logicalChannels.EntityData.SegmentPath = "logical-channels"
    logicalChannels.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    logicalChannels.EntityData.NamespaceTable = openconfig.GetNamespaces()
    logicalChannels.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    logicalChannels.EntityData.Children = types.NewOrderedMap()
    logicalChannels.EntityData.Children.Append("channel", types.YChild{"Channel", nil})
    for i := range logicalChannels.Channel {
        logicalChannels.EntityData.Children.Append(types.GetSegmentPath(logicalChannels.Channel[i]), types.YChild{"Channel", logicalChannels.Channel[i]})
    }
    logicalChannels.EntityData.Leafs = types.NewOrderedMap()

    logicalChannels.EntityData.YListKeys = []string {}

    return &(logicalChannels.EntityData)
}

// TerminalDevice_LogicalChannels_Channel
// List of logical channels
type TerminalDevice_LogicalChannels_Channel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the index of the logical channel. The
    // type is string with range: 0..4294967295. Refers to
    // terminal_device.TerminalDevice_LogicalChannels_Channel_Config_Index
    Index interface{}

    // Configuration data for logical channels.
    Config TerminalDevice_LogicalChannels_Channel_Config

    // Operational state data for logical channels.
    State TerminalDevice_LogicalChannels_Channel_State

    // Top level container for OTU configuration when logical channel framing is
    // using an OTU protocol, e.g., OTU1, OTU3, etc.
    Otn TerminalDevice_LogicalChannels_Channel_Otn

    // Top level container for data related to Ethernet framing for the logical
    // channel.
    Ethernet TerminalDevice_LogicalChannels_Channel_Ethernet

    // Top-level container for specifying references to the source of signal for
    // the logical channel, either a transceiver or individual physical channels.
    Ingress TerminalDevice_LogicalChannels_Channel_Ingress

    // Enclosing container for tributary assignments.
    LogicalChannelAssignments TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments
}

func (channel *TerminalDevice_LogicalChannels_Channel) GetEntityData() *types.CommonEntityData {
    channel.EntityData.YFilter = channel.YFilter
    channel.EntityData.YangName = "channel"
    channel.EntityData.BundleName = "openconfig"
    channel.EntityData.ParentYangName = "logical-channels"
    channel.EntityData.SegmentPath = "channel" + types.AddKeyToken(channel.Index, "index")
    channel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    channel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    channel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    channel.EntityData.Children = types.NewOrderedMap()
    channel.EntityData.Children.Append("config", types.YChild{"Config", &channel.Config})
    channel.EntityData.Children.Append("state", types.YChild{"State", &channel.State})
    channel.EntityData.Children.Append("otn", types.YChild{"Otn", &channel.Otn})
    channel.EntityData.Children.Append("ethernet", types.YChild{"Ethernet", &channel.Ethernet})
    channel.EntityData.Children.Append("ingress", types.YChild{"Ingress", &channel.Ingress})
    channel.EntityData.Children.Append("logical-channel-assignments", types.YChild{"LogicalChannelAssignments", &channel.LogicalChannelAssignments})
    channel.EntityData.Leafs = types.NewOrderedMap()
    channel.EntityData.Leafs.Append("index", types.YLeaf{"Index", channel.Index})

    channel.EntityData.YListKeys = []string {"Index"}

    return &(channel.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Config
// Configuration data for logical channels
type TerminalDevice_LogicalChannels_Channel_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of the current logical channel. The type is interface{} with range:
    // 0..4294967295.
    Index interface{}

    // Description of the logical channel. The type is string.
    Description interface{}

    // Sets the admin state of the logical channel. The type is AdminStateType.
    AdminState interface{}

    // Rounded bit rate of the tributary signal. Exact bit rate will be refined by
    // protocol selection. The type is one of the following:
    // TRIBRATE10GTRIBRATE40GTRIBRATE100GTRIBRATE1GTRIBRATE2DOT5G.
    RateClass interface{}

    // Protocol framing of the tributary signal. If this LogicalChannel is
    // directly connected to a Client-Port or Optical-Channel, this is the
    // protocol of the associated port. If the LogicalChannel is connected to
    // other LogicalChannels, the TributaryProtocol of the LogicalChannels will
    // define a specific mapping/demapping or multiplexing/demultiplexing
    // function.  Not all protocols are valid, depending on the value of
    // trib-rate-class.  The expectation is that the NMS will validate that a
    // correct combination of rate class and protocol are specfied.  Basic
    // combinations are:  rate class: 1G protocols: 1GE  rate class: 2.5G
    // protocols: OC48, STM16  rate class: 10G protocols:  10GE LAN, 10GE WAN,
    // OC192, STM64, OTU2, OTU2e,            OTU1e, ODU2, ODU2e, ODU1e  rate
    // class: 40G protocols:  40GE, OC768, STM256, OTU3, ODU3  rate class: 100G
    // protocols:  100GE, 100G MLG, OTU4, OTUCn, ODU4. The type is one of the
    // following:
    // PROTOTU2EPROTODU2EPROTOC768PROT10GEWANPROTSTM16PROTOTUCNPROT1GEPROT100GEPROTOTU3PROTOTU2PROTOTU4PROTSTM256PROT10GELANPROTOC48PROTOC192PROT40GEPROT100GMLGPROTODU3PROTODU2PROTODU4PROTSTM64PROTOTU1E.
    TribProtocol interface{}

    // The type / stage of the logical element determines the configuration and
    // operational state parameters (PMs) available for the logical element. The
    // type is one of the following: PROTETHERNETPROTOTN.
    LogicalChannelType interface{}

    // Sets the loopback type on the logical channel. Setting the mode to
    // something besides NONE activates the loopback in the specified mode. The
    // type is LoopbackModeType.
    LoopbackMode interface{}
}

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetEntityData() *types.CommonEntityData {
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
    config.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", config.AdminState})
    config.EntityData.Leafs.Append("rate-class", types.YLeaf{"RateClass", config.RateClass})
    config.EntityData.Leafs.Append("trib-protocol", types.YLeaf{"TribProtocol", config.TribProtocol})
    config.EntityData.Leafs.Append("logical-channel-type", types.YLeaf{"LogicalChannelType", config.LogicalChannelType})
    config.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", config.LoopbackMode})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_State
// Operational state data for logical channels
type TerminalDevice_LogicalChannels_Channel_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of the current logical channel. The type is interface{} with range:
    // 0..4294967295.
    Index interface{}

    // Description of the logical channel. The type is string.
    Description interface{}

    // Sets the admin state of the logical channel. The type is AdminStateType.
    AdminState interface{}

    // Rounded bit rate of the tributary signal. Exact bit rate will be refined by
    // protocol selection. The type is one of the following:
    // TRIBRATE10GTRIBRATE40GTRIBRATE100GTRIBRATE1GTRIBRATE2DOT5G.
    RateClass interface{}

    // Protocol framing of the tributary signal. If this LogicalChannel is
    // directly connected to a Client-Port or Optical-Channel, this is the
    // protocol of the associated port. If the LogicalChannel is connected to
    // other LogicalChannels, the TributaryProtocol of the LogicalChannels will
    // define a specific mapping/demapping or multiplexing/demultiplexing
    // function.  Not all protocols are valid, depending on the value of
    // trib-rate-class.  The expectation is that the NMS will validate that a
    // correct combination of rate class and protocol are specfied.  Basic
    // combinations are:  rate class: 1G protocols: 1GE  rate class: 2.5G
    // protocols: OC48, STM16  rate class: 10G protocols:  10GE LAN, 10GE WAN,
    // OC192, STM64, OTU2, OTU2e,            OTU1e, ODU2, ODU2e, ODU1e  rate
    // class: 40G protocols:  40GE, OC768, STM256, OTU3, ODU3  rate class: 100G
    // protocols:  100GE, 100G MLG, OTU4, OTUCn, ODU4. The type is one of the
    // following:
    // PROTOTU2EPROTODU2EPROTOC768PROT10GEWANPROTSTM16PROTOTUCNPROT1GEPROT100GEPROTOTU3PROTOTU2PROTOTU4PROTSTM256PROT10GELANPROTOC48PROTOC192PROT40GEPROT100GMLGPROTODU3PROTODU2PROTODU4PROTSTM64PROTOTU1E.
    TribProtocol interface{}

    // The type / stage of the logical element determines the configuration and
    // operational state parameters (PMs) available for the logical element. The
    // type is one of the following: PROTETHERNETPROTOTN.
    LogicalChannelType interface{}

    // Sets the loopback type on the logical channel. Setting the mode to
    // something besides NONE activates the loopback in the specified mode. The
    // type is LoopbackModeType.
    LoopbackMode interface{}

    // Link-state of the Ethernet protocol on the logical channel, SONET / SDH
    // framed signal, etc. The type is LinkState.
    LinkState interface{}
}

func (state *TerminalDevice_LogicalChannels_Channel_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "channel"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", state.AdminState})
    state.EntityData.Leafs.Append("rate-class", types.YLeaf{"RateClass", state.RateClass})
    state.EntityData.Leafs.Append("trib-protocol", types.YLeaf{"TribProtocol", state.TribProtocol})
    state.EntityData.Leafs.Append("logical-channel-type", types.YLeaf{"LogicalChannelType", state.LogicalChannelType})
    state.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", state.LoopbackMode})
    state.EntityData.Leafs.Append("link-state", types.YLeaf{"LinkState", state.LinkState})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_State_LinkState represents SONET / SDH framed signal, etc.
type TerminalDevice_LogicalChannels_Channel_State_LinkState string

const (
    // Logical channel is operationally up
    TerminalDevice_LogicalChannels_Channel_State_LinkState_UP TerminalDevice_LogicalChannels_Channel_State_LinkState = "UP"

    // Logical channel is operationally down
    TerminalDevice_LogicalChannels_Channel_State_LinkState_DOWN TerminalDevice_LogicalChannels_Channel_State_LinkState = "DOWN"
)

// TerminalDevice_LogicalChannels_Channel_Otn
// Top level container for OTU configuration when logical
// channel framing is using an OTU protocol, e.g., OTU1, OTU3,
// etc.
type TerminalDevice_LogicalChannels_Channel_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for OTN protocol framing.
    Config TerminalDevice_LogicalChannels_Channel_Otn_Config

    // Operational state data for OTN protocol PMs, statistics, etc.
    State TerminalDevice_LogicalChannels_Channel_Otn_State
}

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "openconfig"
    otn.EntityData.ParentYangName = "channel"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    otn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Children.Append("config", types.YChild{"Config", &otn.Config})
    otn.EntityData.Children.Append("state", types.YChild{"State", &otn.State})
    otn.EntityData.Leafs = types.NewOrderedMap()

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Otn_Config
// Configuration data for OTN protocol framing
type TerminalDevice_LogicalChannels_Channel_Otn_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trail trace identifier (TTI) message transmitted. The type is string.
    TtiMsgTransmit interface{}

    // Trail trace identifier (TTI) message expected. The type is string.
    TtiMsgExpected interface{}

    // Trail trace identifier (TTI) transmit message automatically created. If
    // True, then setting a custom transmit message would be invalid. The type is
    // bool.
    TtiMsgAuto interface{}
}

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "otn"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("tti-msg-transmit", types.YLeaf{"TtiMsgTransmit", config.TtiMsgTransmit})
    config.EntityData.Leafs.Append("tti-msg-expected", types.YLeaf{"TtiMsgExpected", config.TtiMsgExpected})
    config.EntityData.Leafs.Append("tti-msg-auto", types.YLeaf{"TtiMsgAuto", config.TtiMsgAuto})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Otn_State
// Operational state data for OTN protocol PMs, statistics,
// etc.
type TerminalDevice_LogicalChannels_Channel_Otn_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trail trace identifier (TTI) message transmitted. The type is string.
    TtiMsgTransmit interface{}

    // Trail trace identifier (TTI) message expected. The type is string.
    TtiMsgExpected interface{}

    // Trail trace identifier (TTI) transmit message automatically created. If
    // True, then setting a custom transmit message would be invalid. The type is
    // bool.
    TtiMsgAuto interface{}

    // Trail trace identifier (TTI) message received. The type is string.
    TtiMsgRecv interface{}

    // Remote defect indication (RDI) message received. The type is string.
    RdiMsg interface{}

    // The number of seconds that at least one errored blocks occurs, at least one
    // code violation occurs, loss of sync is detected or loss of signal is
    // detected. The type is interface{} with range: 0..18446744073709551615.
    ErroredSeconds interface{}

    // The number of seconds that loss of frame is detected OR the number of
    // errored blocks, code violations, loss of sync or loss of signal is detected
    // exceeds a predefined threshold. The type is interface{} with range:
    // 0..18446744073709551615.
    SeverelyErroredSeconds interface{}

    // The number of seconds during which the link is unavailable. The type is
    // interface{} with range: 0..18446744073709551615.
    UnavailableSeconds interface{}

    // For ethernet or fiberchannel links, the number of 8b/10b coding violations.
    // For SONET/SDH, the number of BIP (bit interleaved parity) errors. The type
    // is interface{} with range: 0..18446744073709551615.
    CodeViolations interface{}

    // The number words that were uncorrectable by the FEC. The type is
    // interface{} with range: 0..18446744073709551615.
    FecUncorrectableWords interface{}

    // The number of bytes that were corrected by the FEC. The type is interface{}
    // with range: 0..18446744073709551615.
    FecCorrectedBytes interface{}

    // The number of bits that were corrected by the FEC. The type is interface{}
    // with range: 0..18446744073709551615.
    FecCorrectedBits interface{}

    // The number of background block errors. The type is interface{} with range:
    // 0..18446744073709551615.
    BackgroundBlockErrors interface{}

    // Bit error rate before forward error correction -- computed value.
    PreFecBer TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer

    // Bit error rate after forward error correction -- computed value.
    PostFecBer TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer

    // Quality value (factor) of a channel.
    QValue TerminalDevice_LogicalChannels_Channel_Otn_State_QValue

    // Electrical signal to noise ratio. Baud rate normalized signal to noise
    // ratio based on error vector magnitude.
    Esnr TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr
}

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "otn"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("pre-fec-ber", types.YChild{"PreFecBer", &state.PreFecBer})
    state.EntityData.Children.Append("post-fec-ber", types.YChild{"PostFecBer", &state.PostFecBer})
    state.EntityData.Children.Append("q-value", types.YChild{"QValue", &state.QValue})
    state.EntityData.Children.Append("esnr", types.YChild{"Esnr", &state.Esnr})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("tti-msg-transmit", types.YLeaf{"TtiMsgTransmit", state.TtiMsgTransmit})
    state.EntityData.Leafs.Append("tti-msg-expected", types.YLeaf{"TtiMsgExpected", state.TtiMsgExpected})
    state.EntityData.Leafs.Append("tti-msg-auto", types.YLeaf{"TtiMsgAuto", state.TtiMsgAuto})
    state.EntityData.Leafs.Append("tti-msg-recv", types.YLeaf{"TtiMsgRecv", state.TtiMsgRecv})
    state.EntityData.Leafs.Append("rdi-msg", types.YLeaf{"RdiMsg", state.RdiMsg})
    state.EntityData.Leafs.Append("errored-seconds", types.YLeaf{"ErroredSeconds", state.ErroredSeconds})
    state.EntityData.Leafs.Append("severely-errored-seconds", types.YLeaf{"SeverelyErroredSeconds", state.SeverelyErroredSeconds})
    state.EntityData.Leafs.Append("unavailable-seconds", types.YLeaf{"UnavailableSeconds", state.UnavailableSeconds})
    state.EntityData.Leafs.Append("code-violations", types.YLeaf{"CodeViolations", state.CodeViolations})
    state.EntityData.Leafs.Append("fec-uncorrectable-words", types.YLeaf{"FecUncorrectableWords", state.FecUncorrectableWords})
    state.EntityData.Leafs.Append("fec-corrected-bytes", types.YLeaf{"FecCorrectedBytes", state.FecCorrectedBytes})
    state.EntityData.Leafs.Append("fec-corrected-bits", types.YLeaf{"FecCorrectedBits", state.FecCorrectedBits})
    state.EntityData.Leafs.Append("background-block-errors", types.YLeaf{"BackgroundBlockErrors", state.BackgroundBlockErrors})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer
// Bit error rate before forward error correction -- computed
// value
type TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer struct {
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

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetEntityData() *types.CommonEntityData {
    preFecBer.EntityData.YFilter = preFecBer.YFilter
    preFecBer.EntityData.YangName = "pre-fec-ber"
    preFecBer.EntityData.BundleName = "openconfig"
    preFecBer.EntityData.ParentYangName = "state"
    preFecBer.EntityData.SegmentPath = "pre-fec-ber"
    preFecBer.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    preFecBer.EntityData.NamespaceTable = openconfig.GetNamespaces()
    preFecBer.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    preFecBer.EntityData.Children = types.NewOrderedMap()
    preFecBer.EntityData.Leafs = types.NewOrderedMap()
    preFecBer.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", preFecBer.Instant})
    preFecBer.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", preFecBer.Avg})
    preFecBer.EntityData.Leafs.Append("min", types.YLeaf{"Min", preFecBer.Min})
    preFecBer.EntityData.Leafs.Append("max", types.YLeaf{"Max", preFecBer.Max})

    preFecBer.EntityData.YListKeys = []string {}

    return &(preFecBer.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer
// Bit error rate after forward error correction -- computed
// value
type TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer struct {
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

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetEntityData() *types.CommonEntityData {
    postFecBer.EntityData.YFilter = postFecBer.YFilter
    postFecBer.EntityData.YangName = "post-fec-ber"
    postFecBer.EntityData.BundleName = "openconfig"
    postFecBer.EntityData.ParentYangName = "state"
    postFecBer.EntityData.SegmentPath = "post-fec-ber"
    postFecBer.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    postFecBer.EntityData.NamespaceTable = openconfig.GetNamespaces()
    postFecBer.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    postFecBer.EntityData.Children = types.NewOrderedMap()
    postFecBer.EntityData.Leafs = types.NewOrderedMap()
    postFecBer.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", postFecBer.Instant})
    postFecBer.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", postFecBer.Avg})
    postFecBer.EntityData.Leafs.Append("min", types.YLeaf{"Min", postFecBer.Min})
    postFecBer.EntityData.Leafs.Append("max", types.YLeaf{"Max", postFecBer.Max})

    postFecBer.EntityData.YListKeys = []string {}

    return &(postFecBer.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Otn_State_QValue
// Quality value (factor) of a channel
type TerminalDevice_LogicalChannels_Channel_Otn_State_QValue struct {
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

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetEntityData() *types.CommonEntityData {
    qValue.EntityData.YFilter = qValue.YFilter
    qValue.EntityData.YangName = "q-value"
    qValue.EntityData.BundleName = "openconfig"
    qValue.EntityData.ParentYangName = "state"
    qValue.EntityData.SegmentPath = "q-value"
    qValue.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    qValue.EntityData.NamespaceTable = openconfig.GetNamespaces()
    qValue.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    qValue.EntityData.Children = types.NewOrderedMap()
    qValue.EntityData.Leafs = types.NewOrderedMap()
    qValue.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", qValue.Instant})
    qValue.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", qValue.Avg})
    qValue.EntityData.Leafs.Append("min", types.YLeaf{"Min", qValue.Min})
    qValue.EntityData.Leafs.Append("max", types.YLeaf{"Max", qValue.Max})

    qValue.EntityData.YListKeys = []string {}

    return &(qValue.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr
// Electrical signal to noise ratio. Baud rate
// normalized signal to noise ratio based on
// error vector magnitude
type TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr struct {
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

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetEntityData() *types.CommonEntityData {
    esnr.EntityData.YFilter = esnr.YFilter
    esnr.EntityData.YangName = "esnr"
    esnr.EntityData.BundleName = "openconfig"
    esnr.EntityData.ParentYangName = "state"
    esnr.EntityData.SegmentPath = "esnr"
    esnr.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    esnr.EntityData.NamespaceTable = openconfig.GetNamespaces()
    esnr.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    esnr.EntityData.Children = types.NewOrderedMap()
    esnr.EntityData.Leafs = types.NewOrderedMap()
    esnr.EntityData.Leafs.Append("instant", types.YLeaf{"Instant", esnr.Instant})
    esnr.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", esnr.Avg})
    esnr.EntityData.Leafs.Append("min", types.YLeaf{"Min", esnr.Min})
    esnr.EntityData.Leafs.Append("max", types.YLeaf{"Max", esnr.Max})

    esnr.EntityData.YListKeys = []string {}

    return &(esnr.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Ethernet
// Top level container for data related to Ethernet framing
// for the logical channel
type TerminalDevice_LogicalChannels_Channel_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for Ethernet protocol framing on logical channels.
    Config TerminalDevice_LogicalChannels_Channel_Ethernet_Config

    // Operational state data for Ethernet protocol framing on logical channels.
    State TerminalDevice_LogicalChannels_Channel_Ethernet_State
}

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "openconfig"
    ethernet.EntityData.ParentYangName = "channel"
    ethernet.EntityData.SegmentPath = "ethernet"
    ethernet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ethernet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Children.Append("config", types.YChild{"Config", &ethernet.Config})
    ethernet.EntityData.Children.Append("state", types.YChild{"State", &ethernet.State})
    ethernet.EntityData.Leafs = types.NewOrderedMap()

    ethernet.EntityData.YListKeys = []string {}

    return &(ethernet.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Ethernet_Config
// Configuration data for Ethernet protocol framing on logical
// channels
type TerminalDevice_LogicalChannels_Channel_Ethernet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ethernet"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Ethernet_State
// Operational state data for Ethernet protocol framing on logical
// channels
type TerminalDevice_LogicalChannels_Channel_Ethernet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC layer control frames received on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    InMacControlFrames interface{}

    // MAC layer PAUSE frames received on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    InMacPauseFrames interface{}

    // Number of oversize frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    InOversizeFrames interface{}

    // Number of jabber frames received on the interface.  Jabber frames are
    // typically defined as oversize frames which also have a bad CRC. 
    // Implementations may use slightly different definitions of what constitutes
    // a jabber frame.  Often indicative of a NIC hardware problem. The type is
    // interface{} with range: 0..18446744073709551615.
    InJabberFrames interface{}

    // Number of fragment frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    InFragmentFrames interface{}

    // Number of 802.1q tagged frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    In8021qFrames interface{}

    // Number of receive error events due to FCS/CRC check failure. The type is
    // interface{} with range: 0..18446744073709551615.
    InCrcErrors interface{}

    // MAC layer control frames sent on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    OutMacControlFrames interface{}

    // MAC layer PAUSE frames sent on the interface. The type is interface{} with
    // range: 0..18446744073709551615.
    OutMacPauseFrames interface{}

    // Number of 802.1q tagged frames sent on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    Out8021qFrames interface{}
}

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ethernet"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("in-mac-control-frames", types.YLeaf{"InMacControlFrames", state.InMacControlFrames})
    state.EntityData.Leafs.Append("in-mac-pause-frames", types.YLeaf{"InMacPauseFrames", state.InMacPauseFrames})
    state.EntityData.Leafs.Append("in-oversize-frames", types.YLeaf{"InOversizeFrames", state.InOversizeFrames})
    state.EntityData.Leafs.Append("in-jabber-frames", types.YLeaf{"InJabberFrames", state.InJabberFrames})
    state.EntityData.Leafs.Append("in-fragment-frames", types.YLeaf{"InFragmentFrames", state.InFragmentFrames})
    state.EntityData.Leafs.Append("in-8021q-frames", types.YLeaf{"In8021qFrames", state.In8021qFrames})
    state.EntityData.Leafs.Append("in-crc-errors", types.YLeaf{"InCrcErrors", state.InCrcErrors})
    state.EntityData.Leafs.Append("out-mac-control-frames", types.YLeaf{"OutMacControlFrames", state.OutMacControlFrames})
    state.EntityData.Leafs.Append("out-mac-pause-frames", types.YLeaf{"OutMacPauseFrames", state.OutMacPauseFrames})
    state.EntityData.Leafs.Append("out-8021q-frames", types.YLeaf{"Out8021qFrames", state.Out8021qFrames})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Ingress
// Top-level container for specifying references to the
// source of signal for the logical channel, either a
// transceiver or individual physical channels
type TerminalDevice_LogicalChannels_Channel_Ingress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the signal source for the logical channel.
    Config TerminalDevice_LogicalChannels_Channel_Ingress_Config

    // Operational state data for the signal source for the logical channel.
    State TerminalDevice_LogicalChannels_Channel_Ingress_State
}

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetEntityData() *types.CommonEntityData {
    ingress.EntityData.YFilter = ingress.YFilter
    ingress.EntityData.YangName = "ingress"
    ingress.EntityData.BundleName = "openconfig"
    ingress.EntityData.ParentYangName = "channel"
    ingress.EntityData.SegmentPath = "ingress"
    ingress.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ingress.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ingress.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ingress.EntityData.Children = types.NewOrderedMap()
    ingress.EntityData.Children.Append("config", types.YChild{"Config", &ingress.Config})
    ingress.EntityData.Children.Append("state", types.YChild{"State", &ingress.State})
    ingress.EntityData.Leafs = types.NewOrderedMap()

    ingress.EntityData.YListKeys = []string {}

    return &(ingress.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Ingress_Config
// Configuration data for the signal source for the
// logical channel
type TerminalDevice_LogicalChannels_Channel_Ingress_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the transceiver carrying the input signal for the logical
    // channel.  If specific physical channels are mapped to the logical channel
    // (as opposed to all physical channels carried by the transceiver), they can
    // be specified in the list of physical channel references. The type is
    // string. Refers to platform.Components_Component_Name
    Transceiver interface{}

    // This list should be populated with references to the client physical
    // channels that feed this logical channel from the transceiver specified in
    // the 'transceiver' leaf, which must be specified.  If this leaf-list is
    // empty, all physical channels in the transceiver are assumed to be mapped to
    // the logical channel. The type is slice of string with range: 0..65535.
    // Refers to
    // platform.Components_Component_Transceiver_PhysicalChannels_Channel_Index
    PhysicalChannel []interface{}
}

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ingress"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("transceiver", types.YLeaf{"Transceiver", config.Transceiver})
    config.EntityData.Leafs.Append("physical-channel", types.YLeaf{"PhysicalChannel", config.PhysicalChannel})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_Ingress_State
// Operational state data for the signal source for the
// logical channel
type TerminalDevice_LogicalChannels_Channel_Ingress_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the transceiver carrying the input signal for the logical
    // channel.  If specific physical channels are mapped to the logical channel
    // (as opposed to all physical channels carried by the transceiver), they can
    // be specified in the list of physical channel references. The type is
    // string. Refers to platform.Components_Component_Name
    Transceiver interface{}

    // This list should be populated with references to the client physical
    // channels that feed this logical channel from the transceiver specified in
    // the 'transceiver' leaf, which must be specified.  If this leaf-list is
    // empty, all physical channels in the transceiver are assumed to be mapped to
    // the logical channel. The type is slice of string with range: 0..65535.
    // Refers to
    // platform.Components_Component_Transceiver_PhysicalChannels_Channel_Index
    PhysicalChannel []interface{}
}

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ingress"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("transceiver", types.YLeaf{"Transceiver", state.Transceiver})
    state.EntityData.Leafs.Append("physical-channel", types.YLeaf{"PhysicalChannel", state.PhysicalChannel})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments
// Enclosing container for tributary assignments
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logical channel elements may be assigned directly to optical channels for
    // line-side transmission, or can be further groomed into additional stages of
    // logical channel elements.  The grooming can multiplex (i.e., split the
    // current element into multiple elements in the subsequent stage) or
    // de-multiplex (i.e., combine the current element with other elements into
    // the same element in the subsequent stage) logical elements in each stage. 
    // Note that to support the ability to groom the logical elements, the list of
    // logical channel elements should be populated with an entry for the logical
    // elements at each stage, starting with the initial assignment from the
    // respective client physical port.  Each logical element assignment consists
    // of a pointer to an element in the next stage, or to an optical channel,
    // along with a bandwidth allocation for the corresponding assignment (e.g.,
    // to split or combine signal). The type is slice of
    // TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment.
    Assignment []*TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment
}

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetEntityData() *types.CommonEntityData {
    logicalChannelAssignments.EntityData.YFilter = logicalChannelAssignments.YFilter
    logicalChannelAssignments.EntityData.YangName = "logical-channel-assignments"
    logicalChannelAssignments.EntityData.BundleName = "openconfig"
    logicalChannelAssignments.EntityData.ParentYangName = "channel"
    logicalChannelAssignments.EntityData.SegmentPath = "logical-channel-assignments"
    logicalChannelAssignments.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    logicalChannelAssignments.EntityData.NamespaceTable = openconfig.GetNamespaces()
    logicalChannelAssignments.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    logicalChannelAssignments.EntityData.Children = types.NewOrderedMap()
    logicalChannelAssignments.EntityData.Children.Append("assignment", types.YChild{"Assignment", nil})
    for i := range logicalChannelAssignments.Assignment {
        logicalChannelAssignments.EntityData.Children.Append(types.GetSegmentPath(logicalChannelAssignments.Assignment[i]), types.YChild{"Assignment", logicalChannelAssignments.Assignment[i]})
    }
    logicalChannelAssignments.EntityData.Leafs = types.NewOrderedMap()

    logicalChannelAssignments.EntityData.YListKeys = []string {}

    return &(logicalChannelAssignments.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment
// Logical channel elements may be assigned directly to
// optical channels for line-side transmission, or can be
// further groomed into additional stages of logical channel
// elements.  The grooming can multiplex (i.e., split the
// current element into multiple elements in the subsequent
// stage) or de-multiplex (i.e., combine the current element
// with other elements into the same element in the subsequent
// stage) logical elements in each stage.
// 
// Note that to support the ability to groom the logical
// elements, the list of logical channel elements should be
// populated with an entry for the logical elements at
// each stage, starting with the initial assignment from the
// respective client physical port.
// 
// Each logical element assignment consists of a pointer to
// an element in the next stage, or to an optical channel,
// along with a bandwidth allocation for the corresponding
// assignment (e.g., to split or combine signal).
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the index for the current tributary
    // assignment. The type is string with range: 0..4294967295. Refers to
    // terminal_device.TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config_Index
    Index interface{}

    // Configuration data for tributary assignments.
    Config TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config

    // Operational state data for tributary assignments.
    State TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State
}

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetEntityData() *types.CommonEntityData {
    assignment.EntityData.YFilter = assignment.YFilter
    assignment.EntityData.YangName = "assignment"
    assignment.EntityData.BundleName = "openconfig"
    assignment.EntityData.ParentYangName = "logical-channel-assignments"
    assignment.EntityData.SegmentPath = "assignment" + types.AddKeyToken(assignment.Index, "index")
    assignment.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    assignment.EntityData.NamespaceTable = openconfig.GetNamespaces()
    assignment.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    assignment.EntityData.Children = types.NewOrderedMap()
    assignment.EntityData.Children.Append("config", types.YChild{"Config", &assignment.Config})
    assignment.EntityData.Children.Append("state", types.YChild{"State", &assignment.State})
    assignment.EntityData.Leafs = types.NewOrderedMap()
    assignment.EntityData.Leafs.Append("index", types.YLeaf{"Index", assignment.Index})

    assignment.EntityData.YListKeys = []string {"Index"}

    return &(assignment.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config
// Configuration data for tributary assignments
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of the current logical client channel to tributary mapping. The type
    // is interface{} with range: 0..4294967295.
    Index interface{}

    // Name assigned to the logical client channel. The type is string.
    Description interface{}

    // Each logical channel element may be assigned to subsequent stages of
    // logical elements to implement further grooming, or can be assigned to a
    // line-side optical channel for transmission.  Each assignment also has an
    // associated bandwidth allocation. The type is AssignmentType.
    AssignmentType interface{}

    // Reference to another stage of logical channel elements. The type is string
    // with range: 0..4294967295. Refers to
    // terminal_device.TerminalDevice_LogicalChannels_Channel_Index
    LogicalChannel interface{}

    // Reference to the line-side optical channel that should carry the current
    // logical channel element.  Use this reference to exit the logical element
    // stage. The type is string. Refers to platform.Components_Component_Name
    OpticalChannel interface{}

    // Allocation of the logical client channel to the tributary or sub-channel,
    // expressed in Gbps. The type is string with range:
    // -9223372036854775.808..9223372036854775.807. Units are Gbps.
    Allocation interface{}
}

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "assignment"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("index", types.YLeaf{"Index", config.Index})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})
    config.EntityData.Leafs.Append("assignment-type", types.YLeaf{"AssignmentType", config.AssignmentType})
    config.EntityData.Leafs.Append("logical-channel", types.YLeaf{"LogicalChannel", config.LogicalChannel})
    config.EntityData.Leafs.Append("optical-channel", types.YLeaf{"OpticalChannel", config.OpticalChannel})
    config.EntityData.Leafs.Append("allocation", types.YLeaf{"Allocation", config.Allocation})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config_AssignmentType represents bandwidth allocation.
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config_AssignmentType string

const (
    // Subsequent channel is a logical channel
    TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config_AssignmentType_LOGICAL_CHANNEL TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config_AssignmentType = "LOGICAL_CHANNEL"

    // Subsequent channel is a optical channel / carrier
    TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config_AssignmentType_OPTICAL_CHANNEL TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config_AssignmentType = "OPTICAL_CHANNEL"
)

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State
// Operational state data for tributary assignments
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of the current logical client channel to tributary mapping. The type
    // is interface{} with range: 0..4294967295.
    Index interface{}

    // Name assigned to the logical client channel. The type is string.
    Description interface{}

    // Each logical channel element may be assigned to subsequent stages of
    // logical elements to implement further grooming, or can be assigned to a
    // line-side optical channel for transmission.  Each assignment also has an
    // associated bandwidth allocation. The type is AssignmentType.
    AssignmentType interface{}

    // Reference to another stage of logical channel elements. The type is string
    // with range: 0..4294967295. Refers to
    // terminal_device.TerminalDevice_LogicalChannels_Channel_Index
    LogicalChannel interface{}

    // Reference to the line-side optical channel that should carry the current
    // logical channel element.  Use this reference to exit the logical element
    // stage. The type is string. Refers to platform.Components_Component_Name
    OpticalChannel interface{}

    // Allocation of the logical client channel to the tributary or sub-channel,
    // expressed in Gbps. The type is string with range:
    // -9223372036854775.808..9223372036854775.807. Units are Gbps.
    Allocation interface{}
}

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "assignment"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("assignment-type", types.YLeaf{"AssignmentType", state.AssignmentType})
    state.EntityData.Leafs.Append("logical-channel", types.YLeaf{"LogicalChannel", state.LogicalChannel})
    state.EntityData.Leafs.Append("optical-channel", types.YLeaf{"OpticalChannel", state.OpticalChannel})
    state.EntityData.Leafs.Append("allocation", types.YLeaf{"Allocation", state.Allocation})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State_AssignmentType represents bandwidth allocation.
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State_AssignmentType string

const (
    // Subsequent channel is a logical channel
    TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State_AssignmentType_LOGICAL_CHANNEL TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State_AssignmentType = "LOGICAL_CHANNEL"

    // Subsequent channel is a optical channel / carrier
    TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State_AssignmentType_OPTICAL_CHANNEL TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State_AssignmentType = "OPTICAL_CHANNEL"
)

// TerminalDevice_OperationalModes
// Enclosing container for list of operational modes
type TerminalDevice_OperationalModes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of operational modes supported by the platform. The operational mode
    // provides a platform-defined summary of information such as symbol rate,
    // modulation, pulse shaping, etc. The type is slice of
    // TerminalDevice_OperationalModes_Mode.
    Mode []*TerminalDevice_OperationalModes_Mode
}

func (operationalModes *TerminalDevice_OperationalModes) GetEntityData() *types.CommonEntityData {
    operationalModes.EntityData.YFilter = operationalModes.YFilter
    operationalModes.EntityData.YangName = "operational-modes"
    operationalModes.EntityData.BundleName = "openconfig"
    operationalModes.EntityData.ParentYangName = "terminal-device"
    operationalModes.EntityData.SegmentPath = "operational-modes"
    operationalModes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    operationalModes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    operationalModes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    operationalModes.EntityData.Children = types.NewOrderedMap()
    operationalModes.EntityData.Children.Append("mode", types.YChild{"Mode", nil})
    for i := range operationalModes.Mode {
        operationalModes.EntityData.Children.Append(types.GetSegmentPath(operationalModes.Mode[i]), types.YChild{"Mode", operationalModes.Mode[i]})
    }
    operationalModes.EntityData.Leafs = types.NewOrderedMap()

    operationalModes.EntityData.YListKeys = []string {}

    return &(operationalModes.EntityData)
}

// TerminalDevice_OperationalModes_Mode
// List of operational modes supported by the platform.
// The operational mode provides a platform-defined summary
// of information such as symbol rate, modulation, pulse
// shaping, etc.
type TerminalDevice_OperationalModes_Mode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to mode-id. The type is string with
    // range: 0..65535. Refers to
    // terminal_device.TerminalDevice_OperationalModes_Mode_State_ModeId
    ModeId interface{}

    // Configuration data for operational mode.
    Config TerminalDevice_OperationalModes_Mode_Config

    // Operational state data for the platform-defined operational mode.
    State TerminalDevice_OperationalModes_Mode_State
}

func (mode *TerminalDevice_OperationalModes_Mode) GetEntityData() *types.CommonEntityData {
    mode.EntityData.YFilter = mode.YFilter
    mode.EntityData.YangName = "mode"
    mode.EntityData.BundleName = "openconfig"
    mode.EntityData.ParentYangName = "operational-modes"
    mode.EntityData.SegmentPath = "mode" + types.AddKeyToken(mode.ModeId, "mode-id")
    mode.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    mode.EntityData.NamespaceTable = openconfig.GetNamespaces()
    mode.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    mode.EntityData.Children = types.NewOrderedMap()
    mode.EntityData.Children.Append("config", types.YChild{"Config", &mode.Config})
    mode.EntityData.Children.Append("state", types.YChild{"State", &mode.State})
    mode.EntityData.Leafs = types.NewOrderedMap()
    mode.EntityData.Leafs.Append("mode-id", types.YLeaf{"ModeId", mode.ModeId})

    mode.EntityData.YListKeys = []string {"ModeId"}

    return &(mode.EntityData)
}

// TerminalDevice_OperationalModes_Mode_Config
// Configuration data for operational mode
type TerminalDevice_OperationalModes_Mode_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *TerminalDevice_OperationalModes_Mode_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "mode"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TerminalDevice_OperationalModes_Mode_State
// Operational state data for the platform-defined
// operational mode
type TerminalDevice_OperationalModes_Mode_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Two-octet encoding of the vendor-defined operational mode. The type is
    // interface{} with range: 0..65535.
    ModeId interface{}

    // Vendor-supplied textual description of the characteristics of this
    // operational mode to enable operators to select the appropriate mode for the
    // application. The type is string.
    Description interface{}

    // Identifier to represent the vendor / supplier of the platform and the
    // associated operational mode information. The type is string.
    VendorId interface{}
}

func (state *TerminalDevice_OperationalModes_Mode_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "mode"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("mode-id", types.YLeaf{"ModeId", state.ModeId})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("vendor-id", types.YLeaf{"VendorId", state.VendorId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

