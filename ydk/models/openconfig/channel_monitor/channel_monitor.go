// This model describes operational state data for an optical
// channel monitor (OCM) for optical transport line system
// elements such as wavelength routers (ROADMs) and amplifiers.
package channel_monitor

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package channel_monitor"))
    ydk.RegisterEntity("{http://openconfig.net/yang/channel-monitor channel-monitors}", reflect.TypeOf(ChannelMonitors{}))
    ydk.RegisterEntity("openconfig-channel-monitor:channel-monitors", reflect.TypeOf(ChannelMonitors{}))
}

// ChannelMonitors
// Top-level container for optical channel monitors
type ChannelMonitors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of channel monitors, keyed by channel monitor name. The type is slice
    // of ChannelMonitors_ChannelMonitor.
    ChannelMonitor []ChannelMonitors_ChannelMonitor
}

func (channelMonitors *ChannelMonitors) GetFilter() yfilter.YFilter { return channelMonitors.YFilter }

func (channelMonitors *ChannelMonitors) SetFilter(yf yfilter.YFilter) { channelMonitors.YFilter = yf }

func (channelMonitors *ChannelMonitors) GetGoName(yname string) string {
    if yname == "channel-monitor" { return "ChannelMonitor" }
    return ""
}

func (channelMonitors *ChannelMonitors) GetSegmentPath() string {
    return "openconfig-channel-monitor:channel-monitors"
}

func (channelMonitors *ChannelMonitors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "channel-monitor" {
        for _, c := range channelMonitors.ChannelMonitor {
            if channelMonitors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ChannelMonitors_ChannelMonitor{}
        channelMonitors.ChannelMonitor = append(channelMonitors.ChannelMonitor, child)
        return &channelMonitors.ChannelMonitor[len(channelMonitors.ChannelMonitor)-1]
    }
    return nil
}

func (channelMonitors *ChannelMonitors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range channelMonitors.ChannelMonitor {
        children[channelMonitors.ChannelMonitor[i].GetSegmentPath()] = &channelMonitors.ChannelMonitor[i]
    }
    return children
}

func (channelMonitors *ChannelMonitors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (channelMonitors *ChannelMonitors) GetBundleName() string { return "openconfig" }

func (channelMonitors *ChannelMonitors) GetYangName() string { return "channel-monitors" }

func (channelMonitors *ChannelMonitors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (channelMonitors *ChannelMonitors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (channelMonitors *ChannelMonitors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (channelMonitors *ChannelMonitors) SetParent(parent types.Entity) { channelMonitors.parent = parent }

func (channelMonitors *ChannelMonitors) GetParent() types.Entity { return channelMonitors.parent }

func (channelMonitors *ChannelMonitors) GetParentYangName() string { return "openconfig-channel-monitor" }

// ChannelMonitors_ChannelMonitor
// List of channel monitors, keyed by channel monitor name.
type ChannelMonitors_ChannelMonitor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. References the optical channel monitor name. The
    // type is string. Refers to
    // channel_monitor.ChannelMonitors_ChannelMonitor_Config_Name
    Name interface{}

    // Configuration data .
    Config ChannelMonitors_ChannelMonitor_Config

    // Operational state data .
    State ChannelMonitors_ChannelMonitor_State

    // Enclosing container for the list of values describing the power spectral
    // density distribution.
    Channels ChannelMonitors_ChannelMonitor_Channels
}

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetFilter() yfilter.YFilter { return channelMonitor.YFilter }

func (channelMonitor *ChannelMonitors_ChannelMonitor) SetFilter(yf yfilter.YFilter) { channelMonitor.YFilter = yf }

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "channels" { return "Channels" }
    return ""
}

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetSegmentPath() string {
    return "channel-monitor" + "[name='" + fmt.Sprintf("%v", channelMonitor.Name) + "']"
}

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &channelMonitor.Config
    }
    if childYangName == "state" {
        return &channelMonitor.State
    }
    if childYangName == "channels" {
        return &channelMonitor.Channels
    }
    return nil
}

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &channelMonitor.Config
    children["state"] = &channelMonitor.State
    children["channels"] = &channelMonitor.Channels
    return children
}

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = channelMonitor.Name
    return leafs
}

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetBundleName() string { return "openconfig" }

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetYangName() string { return "channel-monitor" }

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (channelMonitor *ChannelMonitors_ChannelMonitor) SetParent(parent types.Entity) { channelMonitor.parent = parent }

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetParent() types.Entity { return channelMonitor.parent }

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetParentYangName() string { return "channel-monitors" }

// ChannelMonitors_ChannelMonitor_Config
// Configuration data 
type ChannelMonitors_ChannelMonitor_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to system-supplied name of the port on the optical channel
    // monitor (OCM). If this port is embedded in another card (i.e. an amplifier
    // card) the device should still define a port representing the OCM even if it
    // is internal and not physically present on the faceplate of the card. The
    // type is string. Refers to platform.Components_Component_Name
    Name interface{}

    // Reference to system-supplied name of the port that the channel monitor is
    // physically connected to. This port will be of type MONITOR. This port is a
    // tap off of the monitored-port and would be in the same card as the
    // monitored port. If this port is embedded in another card (i.e. an amplifier
    // card) the device should still define a port representing the monitor port
    // if it is internal and not physically present on the faceplate of the card.
    // The type is string. Refers to platform.Components_Component_Name
    MonitorPort interface{}
}

func (config *ChannelMonitors_ChannelMonitor_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *ChannelMonitors_ChannelMonitor_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *ChannelMonitors_ChannelMonitor_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "monitor-port" { return "MonitorPort" }
    return ""
}

func (config *ChannelMonitors_ChannelMonitor_Config) GetSegmentPath() string {
    return "config"
}

func (config *ChannelMonitors_ChannelMonitor_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *ChannelMonitors_ChannelMonitor_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *ChannelMonitors_ChannelMonitor_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["monitor-port"] = config.MonitorPort
    return leafs
}

func (config *ChannelMonitors_ChannelMonitor_Config) GetBundleName() string { return "openconfig" }

func (config *ChannelMonitors_ChannelMonitor_Config) GetYangName() string { return "config" }

func (config *ChannelMonitors_ChannelMonitor_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *ChannelMonitors_ChannelMonitor_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *ChannelMonitors_ChannelMonitor_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *ChannelMonitors_ChannelMonitor_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *ChannelMonitors_ChannelMonitor_Config) GetParent() types.Entity { return config.parent }

func (config *ChannelMonitors_ChannelMonitor_Config) GetParentYangName() string { return "channel-monitor" }

// ChannelMonitors_ChannelMonitor_State
// Operational state data 
type ChannelMonitors_ChannelMonitor_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to system-supplied name of the port on the optical channel
    // monitor (OCM). If this port is embedded in another card (i.e. an amplifier
    // card) the device should still define a port representing the OCM even if it
    // is internal and not physically present on the faceplate of the card. The
    // type is string. Refers to platform.Components_Component_Name
    Name interface{}

    // Reference to system-supplied name of the port that the channel monitor is
    // physically connected to. This port will be of type MONITOR. This port is a
    // tap off of the monitored-port and would be in the same card as the
    // monitored port. If this port is embedded in another card (i.e. an amplifier
    // card) the device should still define a port representing the monitor port
    // if it is internal and not physically present on the faceplate of the card.
    // The type is string. Refers to platform.Components_Component_Name
    MonitorPort interface{}
}

func (state *ChannelMonitors_ChannelMonitor_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *ChannelMonitors_ChannelMonitor_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *ChannelMonitors_ChannelMonitor_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "monitor-port" { return "MonitorPort" }
    return ""
}

func (state *ChannelMonitors_ChannelMonitor_State) GetSegmentPath() string {
    return "state"
}

func (state *ChannelMonitors_ChannelMonitor_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *ChannelMonitors_ChannelMonitor_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *ChannelMonitors_ChannelMonitor_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["monitor-port"] = state.MonitorPort
    return leafs
}

func (state *ChannelMonitors_ChannelMonitor_State) GetBundleName() string { return "openconfig" }

func (state *ChannelMonitors_ChannelMonitor_State) GetYangName() string { return "state" }

func (state *ChannelMonitors_ChannelMonitor_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *ChannelMonitors_ChannelMonitor_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *ChannelMonitors_ChannelMonitor_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *ChannelMonitors_ChannelMonitor_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *ChannelMonitors_ChannelMonitor_State) GetParent() types.Entity { return state.parent }

func (state *ChannelMonitors_ChannelMonitor_State) GetParentYangName() string { return "channel-monitor" }

// ChannelMonitors_ChannelMonitor_Channels
// Enclosing container for the list of values describing
// the power spectral density distribution
type ChannelMonitors_ChannelMonitor_Channels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of tuples describing the PSD distribution. The type is slice of
    // ChannelMonitors_ChannelMonitor_Channels_Channel.
    Channel []ChannelMonitors_ChannelMonitor_Channels_Channel
}

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetFilter() yfilter.YFilter { return channels.YFilter }

func (channels *ChannelMonitors_ChannelMonitor_Channels) SetFilter(yf yfilter.YFilter) { channels.YFilter = yf }

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetGoName(yname string) string {
    if yname == "channel" { return "Channel" }
    return ""
}

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetSegmentPath() string {
    return "channels"
}

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "channel" {
        for _, c := range channels.Channel {
            if channels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ChannelMonitors_ChannelMonitor_Channels_Channel{}
        channels.Channel = append(channels.Channel, child)
        return &channels.Channel[len(channels.Channel)-1]
    }
    return nil
}

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range channels.Channel {
        children[channels.Channel[i].GetSegmentPath()] = &channels.Channel[i]
    }
    return children
}

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetBundleName() string { return "openconfig" }

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetYangName() string { return "channels" }

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (channels *ChannelMonitors_ChannelMonitor_Channels) SetParent(parent types.Entity) { channels.parent = parent }

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetParent() types.Entity { return channels.parent }

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetParentYangName() string { return "channel-monitor" }

// ChannelMonitors_ChannelMonitor_Channels_Channel
// List of tuples describing the PSD distribution
type ChannelMonitors_ChannelMonitor_Channels_Channel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the list key. The type is string with
    // range: 0..18446744073709551615. Refers to
    // channel_monitor.ChannelMonitors_ChannelMonitor_Channels_Channel_State_LowerFrequency
    LowerFrequency interface{}

    // This attribute is a key. Reference to the list key. The type is string with
    // range: 0..18446744073709551615. Refers to
    // channel_monitor.ChannelMonitors_ChannelMonitor_Channels_Channel_State_UpperFrequency
    UpperFrequency interface{}

    // Operational state data for PSD.
    State ChannelMonitors_ChannelMonitor_Channels_Channel_State
}

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetFilter() yfilter.YFilter { return channel.YFilter }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) SetFilter(yf yfilter.YFilter) { channel.YFilter = yf }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetGoName(yname string) string {
    if yname == "lower-frequency" { return "LowerFrequency" }
    if yname == "upper-frequency" { return "UpperFrequency" }
    if yname == "state" { return "State" }
    return ""
}

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetSegmentPath() string {
    return "channel" + "[lower-frequency='" + fmt.Sprintf("%v", channel.LowerFrequency) + "']" + "[upper-frequency='" + fmt.Sprintf("%v", channel.UpperFrequency) + "']"
}

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &channel.State
    }
    return nil
}

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &channel.State
    return children
}

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-frequency"] = channel.LowerFrequency
    leafs["upper-frequency"] = channel.UpperFrequency
    return leafs
}

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetBundleName() string { return "openconfig" }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetYangName() string { return "channel" }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) SetParent(parent types.Entity) { channel.parent = parent }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetParent() types.Entity { return channel.parent }

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetParentYangName() string { return "channels" }

// ChannelMonitors_ChannelMonitor_Channels_Channel_State
// Operational state data for PSD
type ChannelMonitors_ChannelMonitor_Channels_Channel_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Lower frequency of the specified PSD. The type is interface{} with range:
    // 0..18446744073709551615.
    LowerFrequency interface{}

    // Upper frequency of the specified PSD. The type is interface{} with range:
    // 0..18446744073709551615.
    UpperFrequency interface{}

    // Power spectral density expressed in nanowatts per megahertz, nW/MHz.  These
    // units allow the value to often be greater than 1.0.  It also avoids dealing
    // with zero values for 0dBm.  For example, a 40GHz wide channel with 0dBm
    // power would be:  0dBm = 1mW = 10^6nW  40GHz = 40,000MHz  0dBm/40GHz =
    // 10^6nW/40,000MHz = 1000/40 = 25. The type is string with length: 32. Units
    // are nW/MHz.
    Psd interface{}
}

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetGoName(yname string) string {
    if yname == "lower-frequency" { return "LowerFrequency" }
    if yname == "upper-frequency" { return "UpperFrequency" }
    if yname == "psd" { return "Psd" }
    return ""
}

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetSegmentPath() string {
    return "state"
}

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-frequency"] = state.LowerFrequency
    leafs["upper-frequency"] = state.UpperFrequency
    leafs["psd"] = state.Psd
    return leafs
}

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetBundleName() string { return "openconfig" }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetYangName() string { return "state" }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetParent() types.Entity { return state.parent }

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetParentYangName() string { return "channel" }

