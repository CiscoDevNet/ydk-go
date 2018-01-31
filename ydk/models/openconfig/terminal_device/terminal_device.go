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
    parent types.Entity
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

func (terminalDevice *TerminalDevice) GetFilter() yfilter.YFilter { return terminalDevice.YFilter }

func (terminalDevice *TerminalDevice) SetFilter(yf yfilter.YFilter) { terminalDevice.YFilter = yf }

func (terminalDevice *TerminalDevice) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "logical-channels" { return "LogicalChannels" }
    if yname == "operational-modes" { return "OperationalModes" }
    return ""
}

func (terminalDevice *TerminalDevice) GetSegmentPath() string {
    return "openconfig-terminal-device:terminal-device"
}

func (terminalDevice *TerminalDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &terminalDevice.Config
    }
    if childYangName == "state" {
        return &terminalDevice.State
    }
    if childYangName == "logical-channels" {
        return &terminalDevice.LogicalChannels
    }
    if childYangName == "operational-modes" {
        return &terminalDevice.OperationalModes
    }
    return nil
}

func (terminalDevice *TerminalDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &terminalDevice.Config
    children["state"] = &terminalDevice.State
    children["logical-channels"] = &terminalDevice.LogicalChannels
    children["operational-modes"] = &terminalDevice.OperationalModes
    return children
}

func (terminalDevice *TerminalDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (terminalDevice *TerminalDevice) GetBundleName() string { return "openconfig" }

func (terminalDevice *TerminalDevice) GetYangName() string { return "terminal-device" }

func (terminalDevice *TerminalDevice) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (terminalDevice *TerminalDevice) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (terminalDevice *TerminalDevice) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (terminalDevice *TerminalDevice) SetParent(parent types.Entity) { terminalDevice.parent = parent }

func (terminalDevice *TerminalDevice) GetParent() types.Entity { return terminalDevice.parent }

func (terminalDevice *TerminalDevice) GetParentYangName() string { return "openconfig-terminal-device" }

// TerminalDevice_Config
// Configuration data for global terminal-device
type TerminalDevice_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (config *TerminalDevice_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TerminalDevice_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TerminalDevice_Config) GetGoName(yname string) string {
    return ""
}

func (config *TerminalDevice_Config) GetSegmentPath() string {
    return "config"
}

func (config *TerminalDevice_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TerminalDevice_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TerminalDevice_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (config *TerminalDevice_Config) GetBundleName() string { return "openconfig" }

func (config *TerminalDevice_Config) GetYangName() string { return "config" }

func (config *TerminalDevice_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TerminalDevice_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TerminalDevice_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TerminalDevice_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TerminalDevice_Config) GetParent() types.Entity { return config.parent }

func (config *TerminalDevice_Config) GetParentYangName() string { return "terminal-device" }

// TerminalDevice_State
// Operational state data for global terminal device
type TerminalDevice_State struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (state *TerminalDevice_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TerminalDevice_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TerminalDevice_State) GetGoName(yname string) string {
    return ""
}

func (state *TerminalDevice_State) GetSegmentPath() string {
    return "state"
}

func (state *TerminalDevice_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TerminalDevice_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TerminalDevice_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *TerminalDevice_State) GetBundleName() string { return "openconfig" }

func (state *TerminalDevice_State) GetYangName() string { return "state" }

func (state *TerminalDevice_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TerminalDevice_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TerminalDevice_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TerminalDevice_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TerminalDevice_State) GetParent() types.Entity { return state.parent }

func (state *TerminalDevice_State) GetParentYangName() string { return "terminal-device" }

// TerminalDevice_LogicalChannels
// Enclosing container the list of logical channels
type TerminalDevice_LogicalChannels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of logical channels. The type is slice of
    // TerminalDevice_LogicalChannels_Channel.
    Channel []TerminalDevice_LogicalChannels_Channel
}

func (logicalChannels *TerminalDevice_LogicalChannels) GetFilter() yfilter.YFilter { return logicalChannels.YFilter }

func (logicalChannels *TerminalDevice_LogicalChannels) SetFilter(yf yfilter.YFilter) { logicalChannels.YFilter = yf }

func (logicalChannels *TerminalDevice_LogicalChannels) GetGoName(yname string) string {
    if yname == "channel" { return "Channel" }
    return ""
}

func (logicalChannels *TerminalDevice_LogicalChannels) GetSegmentPath() string {
    return "logical-channels"
}

func (logicalChannels *TerminalDevice_LogicalChannels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "channel" {
        for _, c := range logicalChannels.Channel {
            if logicalChannels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TerminalDevice_LogicalChannels_Channel{}
        logicalChannels.Channel = append(logicalChannels.Channel, child)
        return &logicalChannels.Channel[len(logicalChannels.Channel)-1]
    }
    return nil
}

func (logicalChannels *TerminalDevice_LogicalChannels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logicalChannels.Channel {
        children[logicalChannels.Channel[i].GetSegmentPath()] = &logicalChannels.Channel[i]
    }
    return children
}

func (logicalChannels *TerminalDevice_LogicalChannels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logicalChannels *TerminalDevice_LogicalChannels) GetBundleName() string { return "openconfig" }

func (logicalChannels *TerminalDevice_LogicalChannels) GetYangName() string { return "logical-channels" }

func (logicalChannels *TerminalDevice_LogicalChannels) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (logicalChannels *TerminalDevice_LogicalChannels) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (logicalChannels *TerminalDevice_LogicalChannels) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (logicalChannels *TerminalDevice_LogicalChannels) SetParent(parent types.Entity) { logicalChannels.parent = parent }

func (logicalChannels *TerminalDevice_LogicalChannels) GetParent() types.Entity { return logicalChannels.parent }

func (logicalChannels *TerminalDevice_LogicalChannels) GetParentYangName() string { return "terminal-device" }

// TerminalDevice_LogicalChannels_Channel
// List of logical channels
type TerminalDevice_LogicalChannels_Channel struct {
    parent types.Entity
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

func (channel *TerminalDevice_LogicalChannels_Channel) GetFilter() yfilter.YFilter { return channel.YFilter }

func (channel *TerminalDevice_LogicalChannels_Channel) SetFilter(yf yfilter.YFilter) { channel.YFilter = yf }

func (channel *TerminalDevice_LogicalChannels_Channel) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "otn" { return "Otn" }
    if yname == "ethernet" { return "Ethernet" }
    if yname == "ingress" { return "Ingress" }
    if yname == "logical-channel-assignments" { return "LogicalChannelAssignments" }
    return ""
}

func (channel *TerminalDevice_LogicalChannels_Channel) GetSegmentPath() string {
    return "channel" + "[index='" + fmt.Sprintf("%v", channel.Index) + "']"
}

func (channel *TerminalDevice_LogicalChannels_Channel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &channel.Config
    }
    if childYangName == "state" {
        return &channel.State
    }
    if childYangName == "otn" {
        return &channel.Otn
    }
    if childYangName == "ethernet" {
        return &channel.Ethernet
    }
    if childYangName == "ingress" {
        return &channel.Ingress
    }
    if childYangName == "logical-channel-assignments" {
        return &channel.LogicalChannelAssignments
    }
    return nil
}

func (channel *TerminalDevice_LogicalChannels_Channel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &channel.Config
    children["state"] = &channel.State
    children["otn"] = &channel.Otn
    children["ethernet"] = &channel.Ethernet
    children["ingress"] = &channel.Ingress
    children["logical-channel-assignments"] = &channel.LogicalChannelAssignments
    return children
}

func (channel *TerminalDevice_LogicalChannels_Channel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = channel.Index
    return leafs
}

func (channel *TerminalDevice_LogicalChannels_Channel) GetBundleName() string { return "openconfig" }

func (channel *TerminalDevice_LogicalChannels_Channel) GetYangName() string { return "channel" }

func (channel *TerminalDevice_LogicalChannels_Channel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (channel *TerminalDevice_LogicalChannels_Channel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (channel *TerminalDevice_LogicalChannels_Channel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (channel *TerminalDevice_LogicalChannels_Channel) SetParent(parent types.Entity) { channel.parent = parent }

func (channel *TerminalDevice_LogicalChannels_Channel) GetParent() types.Entity { return channel.parent }

func (channel *TerminalDevice_LogicalChannels_Channel) GetParentYangName() string { return "logical-channels" }

// TerminalDevice_LogicalChannels_Channel_Config
// Configuration data for logical channels
type TerminalDevice_LogicalChannels_Channel_Config struct {
    parent types.Entity
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

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TerminalDevice_LogicalChannels_Channel_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "description" { return "Description" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "rate-class" { return "RateClass" }
    if yname == "trib-protocol" { return "TribProtocol" }
    if yname == "logical-channel-type" { return "LogicalChannelType" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    return ""
}

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetSegmentPath() string {
    return "config"
}

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = config.Index
    leafs["description"] = config.Description
    leafs["admin-state"] = config.AdminState
    leafs["rate-class"] = config.RateClass
    leafs["trib-protocol"] = config.TribProtocol
    leafs["logical-channel-type"] = config.LogicalChannelType
    leafs["loopback-mode"] = config.LoopbackMode
    return leafs
}

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetBundleName() string { return "openconfig" }

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetYangName() string { return "config" }

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TerminalDevice_LogicalChannels_Channel_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetParent() types.Entity { return config.parent }

func (config *TerminalDevice_LogicalChannels_Channel_Config) GetParentYangName() string { return "channel" }

// TerminalDevice_LogicalChannels_Channel_State
// Operational state data for logical channels
type TerminalDevice_LogicalChannels_Channel_State struct {
    parent types.Entity
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

func (state *TerminalDevice_LogicalChannels_Channel_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TerminalDevice_LogicalChannels_Channel_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TerminalDevice_LogicalChannels_Channel_State) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "description" { return "Description" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "rate-class" { return "RateClass" }
    if yname == "trib-protocol" { return "TribProtocol" }
    if yname == "logical-channel-type" { return "LogicalChannelType" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    if yname == "link-state" { return "LinkState" }
    return ""
}

func (state *TerminalDevice_LogicalChannels_Channel_State) GetSegmentPath() string {
    return "state"
}

func (state *TerminalDevice_LogicalChannels_Channel_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TerminalDevice_LogicalChannels_Channel_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TerminalDevice_LogicalChannels_Channel_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = state.Index
    leafs["description"] = state.Description
    leafs["admin-state"] = state.AdminState
    leafs["rate-class"] = state.RateClass
    leafs["trib-protocol"] = state.TribProtocol
    leafs["logical-channel-type"] = state.LogicalChannelType
    leafs["loopback-mode"] = state.LoopbackMode
    leafs["link-state"] = state.LinkState
    return leafs
}

func (state *TerminalDevice_LogicalChannels_Channel_State) GetBundleName() string { return "openconfig" }

func (state *TerminalDevice_LogicalChannels_Channel_State) GetYangName() string { return "state" }

func (state *TerminalDevice_LogicalChannels_Channel_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TerminalDevice_LogicalChannels_Channel_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TerminalDevice_LogicalChannels_Channel_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TerminalDevice_LogicalChannels_Channel_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TerminalDevice_LogicalChannels_Channel_State) GetParent() types.Entity { return state.parent }

func (state *TerminalDevice_LogicalChannels_Channel_State) GetParentYangName() string { return "channel" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for OTN protocol framing.
    Config TerminalDevice_LogicalChannels_Channel_Otn_Config

    // Operational state data for OTN protocol PMs, statistics, etc.
    State TerminalDevice_LogicalChannels_Channel_Otn_State
}

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &otn.Config
    }
    if childYangName == "state" {
        return &otn.State
    }
    return nil
}

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &otn.Config
    children["state"] = &otn.State
    return children
}

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetBundleName() string { return "openconfig" }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetYangName() string { return "otn" }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetParent() types.Entity { return otn.parent }

func (otn *TerminalDevice_LogicalChannels_Channel_Otn) GetParentYangName() string { return "channel" }

// TerminalDevice_LogicalChannels_Channel_Otn_Config
// Configuration data for OTN protocol framing
type TerminalDevice_LogicalChannels_Channel_Otn_Config struct {
    parent types.Entity
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

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetGoName(yname string) string {
    if yname == "tti-msg-transmit" { return "TtiMsgTransmit" }
    if yname == "tti-msg-expected" { return "TtiMsgExpected" }
    if yname == "tti-msg-auto" { return "TtiMsgAuto" }
    return ""
}

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetSegmentPath() string {
    return "config"
}

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tti-msg-transmit"] = config.TtiMsgTransmit
    leafs["tti-msg-expected"] = config.TtiMsgExpected
    leafs["tti-msg-auto"] = config.TtiMsgAuto
    return leafs
}

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetBundleName() string { return "openconfig" }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetYangName() string { return "config" }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetParent() types.Entity { return config.parent }

func (config *TerminalDevice_LogicalChannels_Channel_Otn_Config) GetParentYangName() string { return "otn" }

// TerminalDevice_LogicalChannels_Channel_Otn_State
// Operational state data for OTN protocol PMs, statistics,
// etc.
type TerminalDevice_LogicalChannels_Channel_Otn_State struct {
    parent types.Entity
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

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetGoName(yname string) string {
    if yname == "tti-msg-transmit" { return "TtiMsgTransmit" }
    if yname == "tti-msg-expected" { return "TtiMsgExpected" }
    if yname == "tti-msg-auto" { return "TtiMsgAuto" }
    if yname == "tti-msg-recv" { return "TtiMsgRecv" }
    if yname == "rdi-msg" { return "RdiMsg" }
    if yname == "errored-seconds" { return "ErroredSeconds" }
    if yname == "severely-errored-seconds" { return "SeverelyErroredSeconds" }
    if yname == "unavailable-seconds" { return "UnavailableSeconds" }
    if yname == "code-violations" { return "CodeViolations" }
    if yname == "fec-uncorrectable-words" { return "FecUncorrectableWords" }
    if yname == "fec-corrected-bytes" { return "FecCorrectedBytes" }
    if yname == "fec-corrected-bits" { return "FecCorrectedBits" }
    if yname == "background-block-errors" { return "BackgroundBlockErrors" }
    if yname == "pre-fec-ber" { return "PreFecBer" }
    if yname == "post-fec-ber" { return "PostFecBer" }
    if yname == "q-value" { return "QValue" }
    if yname == "esnr" { return "Esnr" }
    return ""
}

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetSegmentPath() string {
    return "state"
}

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pre-fec-ber" {
        return &state.PreFecBer
    }
    if childYangName == "post-fec-ber" {
        return &state.PostFecBer
    }
    if childYangName == "q-value" {
        return &state.QValue
    }
    if childYangName == "esnr" {
        return &state.Esnr
    }
    return nil
}

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pre-fec-ber"] = &state.PreFecBer
    children["post-fec-ber"] = &state.PostFecBer
    children["q-value"] = &state.QValue
    children["esnr"] = &state.Esnr
    return children
}

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tti-msg-transmit"] = state.TtiMsgTransmit
    leafs["tti-msg-expected"] = state.TtiMsgExpected
    leafs["tti-msg-auto"] = state.TtiMsgAuto
    leafs["tti-msg-recv"] = state.TtiMsgRecv
    leafs["rdi-msg"] = state.RdiMsg
    leafs["errored-seconds"] = state.ErroredSeconds
    leafs["severely-errored-seconds"] = state.SeverelyErroredSeconds
    leafs["unavailable-seconds"] = state.UnavailableSeconds
    leafs["code-violations"] = state.CodeViolations
    leafs["fec-uncorrectable-words"] = state.FecUncorrectableWords
    leafs["fec-corrected-bytes"] = state.FecCorrectedBytes
    leafs["fec-corrected-bits"] = state.FecCorrectedBits
    leafs["background-block-errors"] = state.BackgroundBlockErrors
    return leafs
}

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetBundleName() string { return "openconfig" }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetYangName() string { return "state" }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetParent() types.Entity { return state.parent }

func (state *TerminalDevice_LogicalChannels_Channel_Otn_State) GetParentYangName() string { return "otn" }

// TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer
// Bit error rate before forward error correction -- computed
// value
type TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer struct {
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

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetFilter() yfilter.YFilter { return preFecBer.YFilter }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) SetFilter(yf yfilter.YFilter) { preFecBer.YFilter = yf }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetSegmentPath() string {
    return "pre-fec-ber"
}

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = preFecBer.Instant
    leafs["avg"] = preFecBer.Avg
    leafs["min"] = preFecBer.Min
    leafs["max"] = preFecBer.Max
    return leafs
}

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetBundleName() string { return "openconfig" }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetYangName() string { return "pre-fec-ber" }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) SetParent(parent types.Entity) { preFecBer.parent = parent }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetParent() types.Entity { return preFecBer.parent }

func (preFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PreFecBer) GetParentYangName() string { return "state" }

// TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer
// Bit error rate after forward error correction -- computed
// value
type TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer struct {
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

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetFilter() yfilter.YFilter { return postFecBer.YFilter }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) SetFilter(yf yfilter.YFilter) { postFecBer.YFilter = yf }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetSegmentPath() string {
    return "post-fec-ber"
}

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = postFecBer.Instant
    leafs["avg"] = postFecBer.Avg
    leafs["min"] = postFecBer.Min
    leafs["max"] = postFecBer.Max
    return leafs
}

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetBundleName() string { return "openconfig" }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetYangName() string { return "post-fec-ber" }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) SetParent(parent types.Entity) { postFecBer.parent = parent }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetParent() types.Entity { return postFecBer.parent }

func (postFecBer *TerminalDevice_LogicalChannels_Channel_Otn_State_PostFecBer) GetParentYangName() string { return "state" }

// TerminalDevice_LogicalChannels_Channel_Otn_State_QValue
// Quality value (factor) of a channel
type TerminalDevice_LogicalChannels_Channel_Otn_State_QValue struct {
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

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetFilter() yfilter.YFilter { return qValue.YFilter }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) SetFilter(yf yfilter.YFilter) { qValue.YFilter = yf }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetSegmentPath() string {
    return "q-value"
}

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = qValue.Instant
    leafs["avg"] = qValue.Avg
    leafs["min"] = qValue.Min
    leafs["max"] = qValue.Max
    return leafs
}

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetBundleName() string { return "openconfig" }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetYangName() string { return "q-value" }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) SetParent(parent types.Entity) { qValue.parent = parent }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetParent() types.Entity { return qValue.parent }

func (qValue *TerminalDevice_LogicalChannels_Channel_Otn_State_QValue) GetParentYangName() string { return "state" }

// TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr
// Electrical signal to noise ratio. Baud rate
// normalized signal to noise ratio based on
// error vector magnitude
type TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr struct {
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

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetFilter() yfilter.YFilter { return esnr.YFilter }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) SetFilter(yf yfilter.YFilter) { esnr.YFilter = yf }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetGoName(yname string) string {
    if yname == "instant" { return "Instant" }
    if yname == "avg" { return "Avg" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetSegmentPath() string {
    return "esnr"
}

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instant"] = esnr.Instant
    leafs["avg"] = esnr.Avg
    leafs["min"] = esnr.Min
    leafs["max"] = esnr.Max
    return leafs
}

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetBundleName() string { return "openconfig" }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetYangName() string { return "esnr" }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) SetParent(parent types.Entity) { esnr.parent = parent }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetParent() types.Entity { return esnr.parent }

func (esnr *TerminalDevice_LogicalChannels_Channel_Otn_State_Esnr) GetParentYangName() string { return "state" }

// TerminalDevice_LogicalChannels_Channel_Ethernet
// Top level container for data related to Ethernet framing
// for the logical channel
type TerminalDevice_LogicalChannels_Channel_Ethernet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for Ethernet protocol framing on logical channels.
    Config TerminalDevice_LogicalChannels_Channel_Ethernet_Config

    // Operational state data for Ethernet protocol framing on logical channels.
    State TerminalDevice_LogicalChannels_Channel_Ethernet_State
}

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetFilter() yfilter.YFilter { return ethernet.YFilter }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) SetFilter(yf yfilter.YFilter) { ethernet.YFilter = yf }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetSegmentPath() string {
    return "ethernet"
}

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ethernet.Config
    }
    if childYangName == "state" {
        return &ethernet.State
    }
    return nil
}

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ethernet.Config
    children["state"] = &ethernet.State
    return children
}

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetBundleName() string { return "openconfig" }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetYangName() string { return "ethernet" }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) SetParent(parent types.Entity) { ethernet.parent = parent }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetParent() types.Entity { return ethernet.parent }

func (ethernet *TerminalDevice_LogicalChannels_Channel_Ethernet) GetParentYangName() string { return "channel" }

// TerminalDevice_LogicalChannels_Channel_Ethernet_Config
// Configuration data for Ethernet protocol framing on logical
// channels
type TerminalDevice_LogicalChannels_Channel_Ethernet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetGoName(yname string) string {
    return ""
}

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetSegmentPath() string {
    return "config"
}

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetBundleName() string { return "openconfig" }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetYangName() string { return "config" }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetParent() types.Entity { return config.parent }

func (config *TerminalDevice_LogicalChannels_Channel_Ethernet_Config) GetParentYangName() string { return "ethernet" }

// TerminalDevice_LogicalChannels_Channel_Ethernet_State
// Operational state data for Ethernet protocol framing on logical
// channels
type TerminalDevice_LogicalChannels_Channel_Ethernet_State struct {
    parent types.Entity
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
    In8021QFrames interface{}

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
    Out8021QFrames interface{}
}

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetGoName(yname string) string {
    if yname == "in-mac-control-frames" { return "InMacControlFrames" }
    if yname == "in-mac-pause-frames" { return "InMacPauseFrames" }
    if yname == "in-oversize-frames" { return "InOversizeFrames" }
    if yname == "in-jabber-frames" { return "InJabberFrames" }
    if yname == "in-fragment-frames" { return "InFragmentFrames" }
    if yname == "in-8021q-frames" { return "In8021QFrames" }
    if yname == "in-crc-errors" { return "InCrcErrors" }
    if yname == "out-mac-control-frames" { return "OutMacControlFrames" }
    if yname == "out-mac-pause-frames" { return "OutMacPauseFrames" }
    if yname == "out-8021q-frames" { return "Out8021QFrames" }
    return ""
}

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetSegmentPath() string {
    return "state"
}

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-mac-control-frames"] = state.InMacControlFrames
    leafs["in-mac-pause-frames"] = state.InMacPauseFrames
    leafs["in-oversize-frames"] = state.InOversizeFrames
    leafs["in-jabber-frames"] = state.InJabberFrames
    leafs["in-fragment-frames"] = state.InFragmentFrames
    leafs["in-8021q-frames"] = state.In8021QFrames
    leafs["in-crc-errors"] = state.InCrcErrors
    leafs["out-mac-control-frames"] = state.OutMacControlFrames
    leafs["out-mac-pause-frames"] = state.OutMacPauseFrames
    leafs["out-8021q-frames"] = state.Out8021QFrames
    return leafs
}

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetBundleName() string { return "openconfig" }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetYangName() string { return "state" }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetParent() types.Entity { return state.parent }

func (state *TerminalDevice_LogicalChannels_Channel_Ethernet_State) GetParentYangName() string { return "ethernet" }

// TerminalDevice_LogicalChannels_Channel_Ingress
// Top-level container for specifying references to the
// source of signal for the logical channel, either a
// transceiver or individual physical channels
type TerminalDevice_LogicalChannels_Channel_Ingress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the signal source for the logical channel.
    Config TerminalDevice_LogicalChannels_Channel_Ingress_Config

    // Operational state data for the signal source for the logical channel.
    State TerminalDevice_LogicalChannels_Channel_Ingress_State
}

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetFilter() yfilter.YFilter { return ingress.YFilter }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) SetFilter(yf yfilter.YFilter) { ingress.YFilter = yf }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetSegmentPath() string {
    return "ingress"
}

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ingress.Config
    }
    if childYangName == "state" {
        return &ingress.State
    }
    return nil
}

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ingress.Config
    children["state"] = &ingress.State
    return children
}

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetBundleName() string { return "openconfig" }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetYangName() string { return "ingress" }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) SetParent(parent types.Entity) { ingress.parent = parent }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetParent() types.Entity { return ingress.parent }

func (ingress *TerminalDevice_LogicalChannels_Channel_Ingress) GetParentYangName() string { return "channel" }

// TerminalDevice_LogicalChannels_Channel_Ingress_Config
// Configuration data for the signal source for the
// logical channel
type TerminalDevice_LogicalChannels_Channel_Ingress_Config struct {
    parent types.Entity
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

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetGoName(yname string) string {
    if yname == "transceiver" { return "Transceiver" }
    if yname == "physical-channel" { return "PhysicalChannel" }
    return ""
}

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetSegmentPath() string {
    return "config"
}

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transceiver"] = config.Transceiver
    leafs["physical-channel"] = config.PhysicalChannel
    return leafs
}

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetBundleName() string { return "openconfig" }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetYangName() string { return "config" }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetParent() types.Entity { return config.parent }

func (config *TerminalDevice_LogicalChannels_Channel_Ingress_Config) GetParentYangName() string { return "ingress" }

// TerminalDevice_LogicalChannels_Channel_Ingress_State
// Operational state data for the signal source for the
// logical channel
type TerminalDevice_LogicalChannels_Channel_Ingress_State struct {
    parent types.Entity
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

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetGoName(yname string) string {
    if yname == "transceiver" { return "Transceiver" }
    if yname == "physical-channel" { return "PhysicalChannel" }
    return ""
}

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetSegmentPath() string {
    return "state"
}

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transceiver"] = state.Transceiver
    leafs["physical-channel"] = state.PhysicalChannel
    return leafs
}

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetBundleName() string { return "openconfig" }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetYangName() string { return "state" }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetParent() types.Entity { return state.parent }

func (state *TerminalDevice_LogicalChannels_Channel_Ingress_State) GetParentYangName() string { return "ingress" }

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments
// Enclosing container for tributary assignments
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments struct {
    parent types.Entity
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
    Assignment []TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment
}

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetFilter() yfilter.YFilter { return logicalChannelAssignments.YFilter }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) SetFilter(yf yfilter.YFilter) { logicalChannelAssignments.YFilter = yf }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetGoName(yname string) string {
    if yname == "assignment" { return "Assignment" }
    return ""
}

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetSegmentPath() string {
    return "logical-channel-assignments"
}

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "assignment" {
        for _, c := range logicalChannelAssignments.Assignment {
            if logicalChannelAssignments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment{}
        logicalChannelAssignments.Assignment = append(logicalChannelAssignments.Assignment, child)
        return &logicalChannelAssignments.Assignment[len(logicalChannelAssignments.Assignment)-1]
    }
    return nil
}

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logicalChannelAssignments.Assignment {
        children[logicalChannelAssignments.Assignment[i].GetSegmentPath()] = &logicalChannelAssignments.Assignment[i]
    }
    return children
}

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetBundleName() string { return "openconfig" }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetYangName() string { return "logical-channel-assignments" }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) SetParent(parent types.Entity) { logicalChannelAssignments.parent = parent }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetParent() types.Entity { return logicalChannelAssignments.parent }

func (logicalChannelAssignments *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments) GetParentYangName() string { return "channel" }

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
    parent types.Entity
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

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetFilter() yfilter.YFilter { return assignment.YFilter }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) SetFilter(yf yfilter.YFilter) { assignment.YFilter = yf }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetSegmentPath() string {
    return "assignment" + "[index='" + fmt.Sprintf("%v", assignment.Index) + "']"
}

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &assignment.Config
    }
    if childYangName == "state" {
        return &assignment.State
    }
    return nil
}

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &assignment.Config
    children["state"] = &assignment.State
    return children
}

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = assignment.Index
    return leafs
}

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetBundleName() string { return "openconfig" }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetYangName() string { return "assignment" }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) SetParent(parent types.Entity) { assignment.parent = parent }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetParent() types.Entity { return assignment.parent }

func (assignment *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment) GetParentYangName() string { return "logical-channel-assignments" }

// TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config
// Configuration data for tributary assignments
type TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config struct {
    parent types.Entity
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

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "description" { return "Description" }
    if yname == "assignment-type" { return "AssignmentType" }
    if yname == "logical-channel" { return "LogicalChannel" }
    if yname == "optical-channel" { return "OpticalChannel" }
    if yname == "allocation" { return "Allocation" }
    return ""
}

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetSegmentPath() string {
    return "config"
}

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = config.Index
    leafs["description"] = config.Description
    leafs["assignment-type"] = config.AssignmentType
    leafs["logical-channel"] = config.LogicalChannel
    leafs["optical-channel"] = config.OpticalChannel
    leafs["allocation"] = config.Allocation
    return leafs
}

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetBundleName() string { return "openconfig" }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetYangName() string { return "config" }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetParent() types.Entity { return config.parent }

func (config *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_Config) GetParentYangName() string { return "assignment" }

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
    parent types.Entity
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

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "description" { return "Description" }
    if yname == "assignment-type" { return "AssignmentType" }
    if yname == "logical-channel" { return "LogicalChannel" }
    if yname == "optical-channel" { return "OpticalChannel" }
    if yname == "allocation" { return "Allocation" }
    return ""
}

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetSegmentPath() string {
    return "state"
}

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = state.Index
    leafs["description"] = state.Description
    leafs["assignment-type"] = state.AssignmentType
    leafs["logical-channel"] = state.LogicalChannel
    leafs["optical-channel"] = state.OpticalChannel
    leafs["allocation"] = state.Allocation
    return leafs
}

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetBundleName() string { return "openconfig" }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetYangName() string { return "state" }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetParent() types.Entity { return state.parent }

func (state *TerminalDevice_LogicalChannels_Channel_LogicalChannelAssignments_Assignment_State) GetParentYangName() string { return "assignment" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of operational modes supported by the platform. The operational mode
    // provides a platform-defined summary of information such as symbol rate,
    // modulation, pulse shaping, etc. The type is slice of
    // TerminalDevice_OperationalModes_Mode.
    Mode []TerminalDevice_OperationalModes_Mode
}

func (operationalModes *TerminalDevice_OperationalModes) GetFilter() yfilter.YFilter { return operationalModes.YFilter }

func (operationalModes *TerminalDevice_OperationalModes) SetFilter(yf yfilter.YFilter) { operationalModes.YFilter = yf }

func (operationalModes *TerminalDevice_OperationalModes) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    return ""
}

func (operationalModes *TerminalDevice_OperationalModes) GetSegmentPath() string {
    return "operational-modes"
}

func (operationalModes *TerminalDevice_OperationalModes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mode" {
        for _, c := range operationalModes.Mode {
            if operationalModes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TerminalDevice_OperationalModes_Mode{}
        operationalModes.Mode = append(operationalModes.Mode, child)
        return &operationalModes.Mode[len(operationalModes.Mode)-1]
    }
    return nil
}

func (operationalModes *TerminalDevice_OperationalModes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range operationalModes.Mode {
        children[operationalModes.Mode[i].GetSegmentPath()] = &operationalModes.Mode[i]
    }
    return children
}

func (operationalModes *TerminalDevice_OperationalModes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operationalModes *TerminalDevice_OperationalModes) GetBundleName() string { return "openconfig" }

func (operationalModes *TerminalDevice_OperationalModes) GetYangName() string { return "operational-modes" }

func (operationalModes *TerminalDevice_OperationalModes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (operationalModes *TerminalDevice_OperationalModes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (operationalModes *TerminalDevice_OperationalModes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (operationalModes *TerminalDevice_OperationalModes) SetParent(parent types.Entity) { operationalModes.parent = parent }

func (operationalModes *TerminalDevice_OperationalModes) GetParent() types.Entity { return operationalModes.parent }

func (operationalModes *TerminalDevice_OperationalModes) GetParentYangName() string { return "terminal-device" }

// TerminalDevice_OperationalModes_Mode
// List of operational modes supported by the platform.
// The operational mode provides a platform-defined summary
// of information such as symbol rate, modulation, pulse
// shaping, etc.
type TerminalDevice_OperationalModes_Mode struct {
    parent types.Entity
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

func (mode *TerminalDevice_OperationalModes_Mode) GetFilter() yfilter.YFilter { return mode.YFilter }

func (mode *TerminalDevice_OperationalModes_Mode) SetFilter(yf yfilter.YFilter) { mode.YFilter = yf }

func (mode *TerminalDevice_OperationalModes_Mode) GetGoName(yname string) string {
    if yname == "mode-id" { return "ModeId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (mode *TerminalDevice_OperationalModes_Mode) GetSegmentPath() string {
    return "mode" + "[mode-id='" + fmt.Sprintf("%v", mode.ModeId) + "']"
}

func (mode *TerminalDevice_OperationalModes_Mode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &mode.Config
    }
    if childYangName == "state" {
        return &mode.State
    }
    return nil
}

func (mode *TerminalDevice_OperationalModes_Mode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &mode.Config
    children["state"] = &mode.State
    return children
}

func (mode *TerminalDevice_OperationalModes_Mode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode-id"] = mode.ModeId
    return leafs
}

func (mode *TerminalDevice_OperationalModes_Mode) GetBundleName() string { return "openconfig" }

func (mode *TerminalDevice_OperationalModes_Mode) GetYangName() string { return "mode" }

func (mode *TerminalDevice_OperationalModes_Mode) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (mode *TerminalDevice_OperationalModes_Mode) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (mode *TerminalDevice_OperationalModes_Mode) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (mode *TerminalDevice_OperationalModes_Mode) SetParent(parent types.Entity) { mode.parent = parent }

func (mode *TerminalDevice_OperationalModes_Mode) GetParent() types.Entity { return mode.parent }

func (mode *TerminalDevice_OperationalModes_Mode) GetParentYangName() string { return "operational-modes" }

// TerminalDevice_OperationalModes_Mode_Config
// Configuration data for operational mode
type TerminalDevice_OperationalModes_Mode_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (config *TerminalDevice_OperationalModes_Mode_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TerminalDevice_OperationalModes_Mode_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TerminalDevice_OperationalModes_Mode_Config) GetGoName(yname string) string {
    return ""
}

func (config *TerminalDevice_OperationalModes_Mode_Config) GetSegmentPath() string {
    return "config"
}

func (config *TerminalDevice_OperationalModes_Mode_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TerminalDevice_OperationalModes_Mode_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TerminalDevice_OperationalModes_Mode_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (config *TerminalDevice_OperationalModes_Mode_Config) GetBundleName() string { return "openconfig" }

func (config *TerminalDevice_OperationalModes_Mode_Config) GetYangName() string { return "config" }

func (config *TerminalDevice_OperationalModes_Mode_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TerminalDevice_OperationalModes_Mode_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TerminalDevice_OperationalModes_Mode_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TerminalDevice_OperationalModes_Mode_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TerminalDevice_OperationalModes_Mode_Config) GetParent() types.Entity { return config.parent }

func (config *TerminalDevice_OperationalModes_Mode_Config) GetParentYangName() string { return "mode" }

// TerminalDevice_OperationalModes_Mode_State
// Operational state data for the platform-defined
// operational mode
type TerminalDevice_OperationalModes_Mode_State struct {
    parent types.Entity
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

func (state *TerminalDevice_OperationalModes_Mode_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TerminalDevice_OperationalModes_Mode_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TerminalDevice_OperationalModes_Mode_State) GetGoName(yname string) string {
    if yname == "mode-id" { return "ModeId" }
    if yname == "description" { return "Description" }
    if yname == "vendor-id" { return "VendorId" }
    return ""
}

func (state *TerminalDevice_OperationalModes_Mode_State) GetSegmentPath() string {
    return "state"
}

func (state *TerminalDevice_OperationalModes_Mode_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TerminalDevice_OperationalModes_Mode_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TerminalDevice_OperationalModes_Mode_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode-id"] = state.ModeId
    leafs["description"] = state.Description
    leafs["vendor-id"] = state.VendorId
    return leafs
}

func (state *TerminalDevice_OperationalModes_Mode_State) GetBundleName() string { return "openconfig" }

func (state *TerminalDevice_OperationalModes_Mode_State) GetYangName() string { return "state" }

func (state *TerminalDevice_OperationalModes_Mode_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TerminalDevice_OperationalModes_Mode_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TerminalDevice_OperationalModes_Mode_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TerminalDevice_OperationalModes_Mode_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TerminalDevice_OperationalModes_Mode_State) GetParent() types.Entity { return state.parent }

func (state *TerminalDevice_OperationalModes_Mode_State) GetParentYangName() string { return "mode" }

