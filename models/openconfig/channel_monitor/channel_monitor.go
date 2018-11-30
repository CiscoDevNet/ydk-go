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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of channel monitors, keyed by channel monitor name. The type is slice
    // of ChannelMonitors_ChannelMonitor.
    ChannelMonitor []*ChannelMonitors_ChannelMonitor
}

func (channelMonitors *ChannelMonitors) GetEntityData() *types.CommonEntityData {
    channelMonitors.EntityData.YFilter = channelMonitors.YFilter
    channelMonitors.EntityData.YangName = "channel-monitors"
    channelMonitors.EntityData.BundleName = "openconfig"
    channelMonitors.EntityData.ParentYangName = "openconfig-channel-monitor"
    channelMonitors.EntityData.SegmentPath = "openconfig-channel-monitor:channel-monitors"
    channelMonitors.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    channelMonitors.EntityData.NamespaceTable = openconfig.GetNamespaces()
    channelMonitors.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    channelMonitors.EntityData.Children = types.NewOrderedMap()
    channelMonitors.EntityData.Children.Append("channel-monitor", types.YChild{"ChannelMonitor", nil})
    for i := range channelMonitors.ChannelMonitor {
        channelMonitors.EntityData.Children.Append(types.GetSegmentPath(channelMonitors.ChannelMonitor[i]), types.YChild{"ChannelMonitor", channelMonitors.ChannelMonitor[i]})
    }
    channelMonitors.EntityData.Leafs = types.NewOrderedMap()

    channelMonitors.EntityData.YListKeys = []string {}

    return &(channelMonitors.EntityData)
}

// ChannelMonitors_ChannelMonitor
// List of channel monitors, keyed by channel monitor name.
type ChannelMonitors_ChannelMonitor struct {
    EntityData types.CommonEntityData
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

func (channelMonitor *ChannelMonitors_ChannelMonitor) GetEntityData() *types.CommonEntityData {
    channelMonitor.EntityData.YFilter = channelMonitor.YFilter
    channelMonitor.EntityData.YangName = "channel-monitor"
    channelMonitor.EntityData.BundleName = "openconfig"
    channelMonitor.EntityData.ParentYangName = "channel-monitors"
    channelMonitor.EntityData.SegmentPath = "channel-monitor" + types.AddKeyToken(channelMonitor.Name, "name")
    channelMonitor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    channelMonitor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    channelMonitor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    channelMonitor.EntityData.Children = types.NewOrderedMap()
    channelMonitor.EntityData.Children.Append("config", types.YChild{"Config", &channelMonitor.Config})
    channelMonitor.EntityData.Children.Append("state", types.YChild{"State", &channelMonitor.State})
    channelMonitor.EntityData.Children.Append("channels", types.YChild{"Channels", &channelMonitor.Channels})
    channelMonitor.EntityData.Leafs = types.NewOrderedMap()
    channelMonitor.EntityData.Leafs.Append("name", types.YLeaf{"Name", channelMonitor.Name})

    channelMonitor.EntityData.YListKeys = []string {"Name"}

    return &(channelMonitor.EntityData)
}

// ChannelMonitors_ChannelMonitor_Config
// Configuration data 
type ChannelMonitors_ChannelMonitor_Config struct {
    EntityData types.CommonEntityData
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

func (config *ChannelMonitors_ChannelMonitor_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "channel-monitor"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("monitor-port", types.YLeaf{"MonitorPort", config.MonitorPort})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// ChannelMonitors_ChannelMonitor_State
// Operational state data 
type ChannelMonitors_ChannelMonitor_State struct {
    EntityData types.CommonEntityData
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

func (state *ChannelMonitors_ChannelMonitor_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "channel-monitor"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("monitor-port", types.YLeaf{"MonitorPort", state.MonitorPort})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// ChannelMonitors_ChannelMonitor_Channels
// Enclosing container for the list of values describing
// the power spectral density distribution
type ChannelMonitors_ChannelMonitor_Channels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of tuples describing the PSD distribution. The type is slice of
    // ChannelMonitors_ChannelMonitor_Channels_Channel.
    Channel []*ChannelMonitors_ChannelMonitor_Channels_Channel
}

func (channels *ChannelMonitors_ChannelMonitor_Channels) GetEntityData() *types.CommonEntityData {
    channels.EntityData.YFilter = channels.YFilter
    channels.EntityData.YangName = "channels"
    channels.EntityData.BundleName = "openconfig"
    channels.EntityData.ParentYangName = "channel-monitor"
    channels.EntityData.SegmentPath = "channels"
    channels.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    channels.EntityData.NamespaceTable = openconfig.GetNamespaces()
    channels.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    channels.EntityData.Children = types.NewOrderedMap()
    channels.EntityData.Children.Append("channel", types.YChild{"Channel", nil})
    for i := range channels.Channel {
        channels.EntityData.Children.Append(types.GetSegmentPath(channels.Channel[i]), types.YChild{"Channel", channels.Channel[i]})
    }
    channels.EntityData.Leafs = types.NewOrderedMap()

    channels.EntityData.YListKeys = []string {}

    return &(channels.EntityData)
}

// ChannelMonitors_ChannelMonitor_Channels_Channel
// List of tuples describing the PSD distribution
type ChannelMonitors_ChannelMonitor_Channels_Channel struct {
    EntityData types.CommonEntityData
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

func (channel *ChannelMonitors_ChannelMonitor_Channels_Channel) GetEntityData() *types.CommonEntityData {
    channel.EntityData.YFilter = channel.YFilter
    channel.EntityData.YangName = "channel"
    channel.EntityData.BundleName = "openconfig"
    channel.EntityData.ParentYangName = "channels"
    channel.EntityData.SegmentPath = "channel" + types.AddKeyToken(channel.LowerFrequency, "lower-frequency") + types.AddKeyToken(channel.UpperFrequency, "upper-frequency")
    channel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    channel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    channel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    channel.EntityData.Children = types.NewOrderedMap()
    channel.EntityData.Children.Append("state", types.YChild{"State", &channel.State})
    channel.EntityData.Leafs = types.NewOrderedMap()
    channel.EntityData.Leafs.Append("lower-frequency", types.YLeaf{"LowerFrequency", channel.LowerFrequency})
    channel.EntityData.Leafs.Append("upper-frequency", types.YLeaf{"UpperFrequency", channel.UpperFrequency})

    channel.EntityData.YListKeys = []string {"LowerFrequency", "UpperFrequency"}

    return &(channel.EntityData)
}

// ChannelMonitors_ChannelMonitor_Channels_Channel_State
// Operational state data for PSD
type ChannelMonitors_ChannelMonitor_Channels_Channel_State struct {
    EntityData types.CommonEntityData
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

func (state *ChannelMonitors_ChannelMonitor_Channels_Channel_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("lower-frequency", types.YLeaf{"LowerFrequency", state.LowerFrequency})
    state.EntityData.Leafs.Append("upper-frequency", types.YLeaf{"UpperFrequency", state.UpperFrequency})
    state.EntityData.Leafs.Append("psd", types.YLeaf{"Psd", state.Psd})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

