// Data model which creates the configuration for the telemetry
// systems and functions on the device.
package telemetry

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package telemetry"))
    ydk.RegisterEntity("{http://openconfig.net/yang/telemetry telemetry-system}", reflect.TypeOf(TelemetrySystem{}))
    ydk.RegisterEntity("openconfig-telemetry:telemetry-system", reflect.TypeOf(TelemetrySystem{}))
}

// TelemetryStreamProtocol represents of the telemetry stream
type TelemetryStreamProtocol string

const (
    // TCP protocol transport for the stream
    TelemetryStreamProtocol_TCP TelemetryStreamProtocol = "TCP"

    // UDP protocol transport for the stream
    TelemetryStreamProtocol_UDP TelemetryStreamProtocol = "UDP"
)

// TelemetrySystem
// Top level configuration and state for the
// device's telemetry system.
type TelemetrySystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top level container for sensor-groups.
    SensorGroups TelemetrySystem_SensorGroups

    // Top level container for destination group configuration and state.
    DestinationGroups TelemetrySystem_DestinationGroups

    // This container holds information for both persistent and dynamic telemetry
    // subscriptions.
    Subscriptions TelemetrySystem_Subscriptions
}

func (telemetrySystem *TelemetrySystem) GetEntityData() *types.CommonEntityData {
    telemetrySystem.EntityData.YFilter = telemetrySystem.YFilter
    telemetrySystem.EntityData.YangName = "telemetry-system"
    telemetrySystem.EntityData.BundleName = "openconfig"
    telemetrySystem.EntityData.ParentYangName = "openconfig-telemetry"
    telemetrySystem.EntityData.SegmentPath = "openconfig-telemetry:telemetry-system"
    telemetrySystem.EntityData.AbsolutePath = telemetrySystem.EntityData.SegmentPath
    telemetrySystem.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    telemetrySystem.EntityData.NamespaceTable = openconfig.GetNamespaces()
    telemetrySystem.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    telemetrySystem.EntityData.Children = types.NewOrderedMap()
    telemetrySystem.EntityData.Children.Append("sensor-groups", types.YChild{"SensorGroups", &telemetrySystem.SensorGroups})
    telemetrySystem.EntityData.Children.Append("destination-groups", types.YChild{"DestinationGroups", &telemetrySystem.DestinationGroups})
    telemetrySystem.EntityData.Children.Append("subscriptions", types.YChild{"Subscriptions", &telemetrySystem.Subscriptions})
    telemetrySystem.EntityData.Leafs = types.NewOrderedMap()

    telemetrySystem.EntityData.YListKeys = []string {}

    return &(telemetrySystem.EntityData)
}

// TelemetrySystem_SensorGroups
// Top level container for sensor-groups.
type TelemetrySystem_SensorGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of telemetry sensory groups on the local system, where a sensor
    // grouping represents a resuable grouping of multiple paths and exclude
    // filters. The type is slice of TelemetrySystem_SensorGroups_SensorGroup.
    SensorGroup []*TelemetrySystem_SensorGroups_SensorGroup
}

func (sensorGroups *TelemetrySystem_SensorGroups) GetEntityData() *types.CommonEntityData {
    sensorGroups.EntityData.YFilter = sensorGroups.YFilter
    sensorGroups.EntityData.YangName = "sensor-groups"
    sensorGroups.EntityData.BundleName = "openconfig"
    sensorGroups.EntityData.ParentYangName = "telemetry-system"
    sensorGroups.EntityData.SegmentPath = "sensor-groups"
    sensorGroups.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/" + sensorGroups.EntityData.SegmentPath
    sensorGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorGroups.EntityData.Children = types.NewOrderedMap()
    sensorGroups.EntityData.Children.Append("sensor-group", types.YChild{"SensorGroup", nil})
    for i := range sensorGroups.SensorGroup {
        sensorGroups.EntityData.Children.Append(types.GetSegmentPath(sensorGroups.SensorGroup[i]), types.YChild{"SensorGroup", sensorGroups.SensorGroup[i]})
    }
    sensorGroups.EntityData.Leafs = types.NewOrderedMap()

    sensorGroups.EntityData.YListKeys = []string {}

    return &(sensorGroups.EntityData)
}

// TelemetrySystem_SensorGroups_SensorGroup
// List of telemetry sensory groups on the local
// system, where a sensor grouping represents a resuable
// grouping of multiple paths and exclude filters.
type TelemetrySystem_SensorGroups_SensorGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the name or identifier of the sensor
    // grouping. The type is string. Refers to
    // telemetry.TelemetrySystem_SensorGroups_SensorGroup_Config_SensorGroupId
    SensorGroupId interface{}

    // Configuration parameters relating to the telemetry sensor grouping.
    Config TelemetrySystem_SensorGroups_SensorGroup_Config

    // State information relating to the telemetry sensor group.
    State TelemetrySystem_SensorGroups_SensorGroup_State

    // Top level container to hold a set of sensor paths grouped together.
    SensorPaths TelemetrySystem_SensorGroups_SensorGroup_SensorPaths
}

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetEntityData() *types.CommonEntityData {
    sensorGroup.EntityData.YFilter = sensorGroup.YFilter
    sensorGroup.EntityData.YangName = "sensor-group"
    sensorGroup.EntityData.BundleName = "openconfig"
    sensorGroup.EntityData.ParentYangName = "sensor-groups"
    sensorGroup.EntityData.SegmentPath = "sensor-group" + types.AddKeyToken(sensorGroup.SensorGroupId, "sensor-group-id")
    sensorGroup.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/sensor-groups/" + sensorGroup.EntityData.SegmentPath
    sensorGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorGroup.EntityData.Children = types.NewOrderedMap()
    sensorGroup.EntityData.Children.Append("config", types.YChild{"Config", &sensorGroup.Config})
    sensorGroup.EntityData.Children.Append("state", types.YChild{"State", &sensorGroup.State})
    sensorGroup.EntityData.Children.Append("sensor-paths", types.YChild{"SensorPaths", &sensorGroup.SensorPaths})
    sensorGroup.EntityData.Leafs = types.NewOrderedMap()
    sensorGroup.EntityData.Leafs.Append("sensor-group-id", types.YLeaf{"SensorGroupId", sensorGroup.SensorGroupId})

    sensorGroup.EntityData.YListKeys = []string {"SensorGroupId"}

    return &(sensorGroup.EntityData)
}

// TelemetrySystem_SensorGroups_SensorGroup_Config
// Configuration parameters relating to the
// telemetry sensor grouping
type TelemetrySystem_SensorGroups_SensorGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name or identifier for the sensor group itself. Will be referenced by other
    // configuration specifying a sensor group. The type is string.
    SensorGroupId interface{}
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "sensor-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/sensor-groups/sensor-group/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("sensor-group-id", types.YLeaf{"SensorGroupId", config.SensorGroupId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TelemetrySystem_SensorGroups_SensorGroup_State
// State information relating to the telemetry
// sensor group
type TelemetrySystem_SensorGroups_SensorGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name or identifier for the sensor group itself. Will be referenced by other
    // configuration specifying a sensor group. The type is string.
    SensorGroupId interface{}
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "sensor-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/sensor-groups/sensor-group/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("sensor-group-id", types.YLeaf{"SensorGroupId", state.SensorGroupId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths
// Top level container to hold a set of sensor
// paths grouped together
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of paths in the model which together comprise a sensor grouping.
    // Filters for each path to exclude items are also provided. The type is slice
    // of TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath.
    SensorPath []*TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath
}

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetEntityData() *types.CommonEntityData {
    sensorPaths.EntityData.YFilter = sensorPaths.YFilter
    sensorPaths.EntityData.YangName = "sensor-paths"
    sensorPaths.EntityData.BundleName = "openconfig"
    sensorPaths.EntityData.ParentYangName = "sensor-group"
    sensorPaths.EntityData.SegmentPath = "sensor-paths"
    sensorPaths.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/sensor-groups/sensor-group/" + sensorPaths.EntityData.SegmentPath
    sensorPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorPaths.EntityData.Children = types.NewOrderedMap()
    sensorPaths.EntityData.Children.Append("sensor-path", types.YChild{"SensorPath", nil})
    for i := range sensorPaths.SensorPath {
        sensorPaths.EntityData.Children.Append(types.GetSegmentPath(sensorPaths.SensorPath[i]), types.YChild{"SensorPath", sensorPaths.SensorPath[i]})
    }
    sensorPaths.EntityData.Leafs = types.NewOrderedMap()

    sensorPaths.EntityData.YListKeys = []string {}

    return &(sensorPaths.EntityData)
}

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath
// List of paths in the model which together
// comprise a sensor grouping. Filters for each path
// to exclude items are also provided.
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the path of interest. The type is
    // string. Refers to
    // telemetry.TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config_Path
    Path interface{}

    // Configuration parameters to configure a set of data model paths as a sensor
    // grouping.
    Config TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config

    // Configuration parameters to configure a set of data model paths as a sensor
    // grouping.
    State TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State
}

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetEntityData() *types.CommonEntityData {
    sensorPath.EntityData.YFilter = sensorPath.YFilter
    sensorPath.EntityData.YangName = "sensor-path"
    sensorPath.EntityData.BundleName = "openconfig"
    sensorPath.EntityData.ParentYangName = "sensor-paths"
    sensorPath.EntityData.SegmentPath = "sensor-path" + types.AddKeyToken(sensorPath.Path, "path")
    sensorPath.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/sensor-groups/sensor-group/sensor-paths/" + sensorPath.EntityData.SegmentPath
    sensorPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorPath.EntityData.Children = types.NewOrderedMap()
    sensorPath.EntityData.Children.Append("config", types.YChild{"Config", &sensorPath.Config})
    sensorPath.EntityData.Children.Append("state", types.YChild{"State", &sensorPath.State})
    sensorPath.EntityData.Leafs = types.NewOrderedMap()
    sensorPath.EntityData.Leafs.Append("path", types.YLeaf{"Path", sensorPath.Path})

    sensorPath.EntityData.YListKeys = []string {"Path"}

    return &(sensorPath.EntityData)
}

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config
// Configuration parameters to configure a set
// of data model paths as a sensor grouping
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path to a section of operational state of interest (the sensor). The type
    // is string.
    Path interface{}

    // Filter to exclude certain values out of the state values. The type is
    // string.
    ExcludeFilter interface{}
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "sensor-path"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/sensor-groups/sensor-group/sensor-paths/sensor-path/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("path", types.YLeaf{"Path", config.Path})
    config.EntityData.Leafs.Append("exclude-filter", types.YLeaf{"ExcludeFilter", config.ExcludeFilter})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State
// Configuration parameters to configure a set
// of data model paths as a sensor grouping
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path to a section of operational state of interest (the sensor). The type
    // is string.
    Path interface{}

    // Filter to exclude certain values out of the state values. The type is
    // string.
    ExcludeFilter interface{}
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "sensor-path"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/sensor-groups/sensor-group/sensor-paths/sensor-path/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("path", types.YLeaf{"Path", state.Path})
    state.EntityData.Leafs.Append("exclude-filter", types.YLeaf{"ExcludeFilter", state.ExcludeFilter})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_DestinationGroups
// Top level container for destination group configuration
// and state.
type TelemetrySystem_DestinationGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of destination-groups. Destination groups allow the reuse of common
    // telemetry destinations across the telemetry configuration. An operator
    // references a set of destinations via the configurable
    // destination-group-identifier.  A destination group may contain one or more
    // telemetry destinations. The type is slice of
    // TelemetrySystem_DestinationGroups_DestinationGroup.
    DestinationGroup []*TelemetrySystem_DestinationGroups_DestinationGroup
}

func (destinationGroups *TelemetrySystem_DestinationGroups) GetEntityData() *types.CommonEntityData {
    destinationGroups.EntityData.YFilter = destinationGroups.YFilter
    destinationGroups.EntityData.YangName = "destination-groups"
    destinationGroups.EntityData.BundleName = "openconfig"
    destinationGroups.EntityData.ParentYangName = "telemetry-system"
    destinationGroups.EntityData.SegmentPath = "destination-groups"
    destinationGroups.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/" + destinationGroups.EntityData.SegmentPath
    destinationGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    destinationGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    destinationGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    destinationGroups.EntityData.Children = types.NewOrderedMap()
    destinationGroups.EntityData.Children.Append("destination-group", types.YChild{"DestinationGroup", nil})
    for i := range destinationGroups.DestinationGroup {
        destinationGroups.EntityData.Children.Append(types.GetSegmentPath(destinationGroups.DestinationGroup[i]), types.YChild{"DestinationGroup", destinationGroups.DestinationGroup[i]})
    }
    destinationGroups.EntityData.Leafs = types.NewOrderedMap()

    destinationGroups.EntityData.YListKeys = []string {}

    return &(destinationGroups.EntityData)
}

// TelemetrySystem_DestinationGroups_DestinationGroup
// List of destination-groups. Destination groups allow the
// reuse of common telemetry destinations across the
// telemetry configuration. An operator references a
// set of destinations via the configurable
// destination-group-identifier.
// 
// A destination group may contain one or more telemetry
// destinations
type TelemetrySystem_DestinationGroups_DestinationGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique identifier for the destination group. The
    // type is string. Refers to
    // telemetry.TelemetrySystem_DestinationGroups_DestinationGroup_Config_GroupId
    GroupId interface{}

    // Top level config container for destination groups.
    Config TelemetrySystem_DestinationGroups_DestinationGroup_Config

    // Top level state container for destination groups.
    State TelemetrySystem_DestinationGroups_DestinationGroup_State

    // The destination container lists the destination information such as IP
    // address and port of the telemetry messages from the network element.
    Destinations TelemetrySystem_DestinationGroups_DestinationGroup_Destinations
}

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetEntityData() *types.CommonEntityData {
    destinationGroup.EntityData.YFilter = destinationGroup.YFilter
    destinationGroup.EntityData.YangName = "destination-group"
    destinationGroup.EntityData.BundleName = "openconfig"
    destinationGroup.EntityData.ParentYangName = "destination-groups"
    destinationGroup.EntityData.SegmentPath = "destination-group" + types.AddKeyToken(destinationGroup.GroupId, "group-id")
    destinationGroup.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/destination-groups/" + destinationGroup.EntityData.SegmentPath
    destinationGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    destinationGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    destinationGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    destinationGroup.EntityData.Children = types.NewOrderedMap()
    destinationGroup.EntityData.Children.Append("config", types.YChild{"Config", &destinationGroup.Config})
    destinationGroup.EntityData.Children.Append("state", types.YChild{"State", &destinationGroup.State})
    destinationGroup.EntityData.Children.Append("destinations", types.YChild{"Destinations", &destinationGroup.Destinations})
    destinationGroup.EntityData.Leafs = types.NewOrderedMap()
    destinationGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", destinationGroup.GroupId})

    destinationGroup.EntityData.YListKeys = []string {"GroupId"}

    return &(destinationGroup.EntityData)
}

// TelemetrySystem_DestinationGroups_DestinationGroup_Config
// Top level config container for destination groups
type TelemetrySystem_DestinationGroups_DestinationGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique identifier for the destination group. The type is string.
    GroupId interface{}
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "destination-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/destination-groups/destination-group/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", config.GroupId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TelemetrySystem_DestinationGroups_DestinationGroup_State
// Top level state container for destination groups
type TelemetrySystem_DestinationGroups_DestinationGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique identifier for destination group. The type is string.
    GroupId interface{}
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "destination-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/destination-groups/destination-group/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", state.GroupId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations
// The destination container lists the destination
// information such as IP address and port of the
// telemetry messages from the network element.
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of telemetry stream destinations. The type is slice of
    // TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination.
    Destination []*TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination
}

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetEntityData() *types.CommonEntityData {
    destinations.EntityData.YFilter = destinations.YFilter
    destinations.EntityData.YangName = "destinations"
    destinations.EntityData.BundleName = "openconfig"
    destinations.EntityData.ParentYangName = "destination-group"
    destinations.EntityData.SegmentPath = "destinations"
    destinations.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/destination-groups/destination-group/" + destinations.EntityData.SegmentPath
    destinations.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    destinations.EntityData.NamespaceTable = openconfig.GetNamespaces()
    destinations.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    destinations.EntityData.Children = types.NewOrderedMap()
    destinations.EntityData.Children.Append("destination", types.YChild{"Destination", nil})
    for i := range destinations.Destination {
        destinations.EntityData.Children.Append(types.GetSegmentPath(destinations.Destination[i]), types.YChild{"Destination", destinations.Destination[i]})
    }
    destinations.EntityData.Leafs = types.NewOrderedMap()

    destinations.EntityData.YListKeys = []string {}

    return &(destinations.EntityData)
}

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination
// List of telemetry stream destinations
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the destination address of the
    // telemetry stream. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // This attribute is a key. Reference to the port number of the stream
    // destination. The type is string with range: 0..65535. Refers to
    // telemetry.TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config_DestinationPort
    DestinationPort interface{}

    // Configuration parameters relating to telemetry destinations.
    Config TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config

    // State information associated with telemetry destinations.
    State TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State
}

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "openconfig"
    destination.EntityData.ParentYangName = "destinations"
    destination.EntityData.SegmentPath = "destination" + types.AddKeyToken(destination.DestinationAddress, "destination-address") + types.AddKeyToken(destination.DestinationPort, "destination-port")
    destination.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/destination-groups/destination-group/destinations/" + destination.EntityData.SegmentPath
    destination.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    destination.EntityData.NamespaceTable = openconfig.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    destination.EntityData.Children = types.NewOrderedMap()
    destination.EntityData.Children.Append("config", types.YChild{"Config", &destination.Config})
    destination.EntityData.Children.Append("state", types.YChild{"State", &destination.State})
    destination.EntityData.Leafs = types.NewOrderedMap()
    destination.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", destination.DestinationAddress})
    destination.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", destination.DestinationPort})

    destination.EntityData.YListKeys = []string {"DestinationAddress", "DestinationPort"}

    return &(destination.EntityData)
}

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config
// Configuration parameters relating to
// telemetry destinations
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the telemetry stream destination. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Protocol (udp or tcp) port number for the telemetry stream destination. The
    // type is interface{} with range: 0..65535.
    DestinationPort interface{}

    // Protocol used to transmit telemetry data to the collector. The type is
    // TelemetryStreamProtocol.
    DestinationProtocol interface{}
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "destination"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/destination-groups/destination-group/destinations/destination/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", config.DestinationAddress})
    config.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", config.DestinationPort})
    config.EntityData.Leafs.Append("destination-protocol", types.YLeaf{"DestinationProtocol", config.DestinationProtocol})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State
// State information associated with
// telemetry destinations
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the telemetry stream destination. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Protocol (udp or tcp) port number for the telemetry stream destination. The
    // type is interface{} with range: 0..65535.
    DestinationPort interface{}

    // Protocol used to transmit telemetry data to the collector. The type is
    // TelemetryStreamProtocol.
    DestinationProtocol interface{}
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "destination"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/destination-groups/destination-group/destinations/destination/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", state.DestinationAddress})
    state.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", state.DestinationPort})
    state.EntityData.Leafs.Append("destination-protocol", types.YLeaf{"DestinationProtocol", state.DestinationProtocol})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_Subscriptions
// This container holds information for both persistent
// and dynamic telemetry subscriptions.
type TelemetrySystem_Subscriptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This container holds information relating to persistent telemetry
    // subscriptions. A persistent telemetry subscription is configued locally on
    // the device through configuration, and is persistent across device restarts
    // or other redundancy changes.
    Persistent TelemetrySystem_Subscriptions_Persistent

    // This container holds information relating to dynamic telemetry
    // subscriptions. A dynamic subscription is typically configured through an
    // RPC channel, and does not persist across device restarts, or if the RPC
    // channel is reset or otherwise torn down.
    Dynamic TelemetrySystem_Subscriptions_Dynamic
}

func (subscriptions *TelemetrySystem_Subscriptions) GetEntityData() *types.CommonEntityData {
    subscriptions.EntityData.YFilter = subscriptions.YFilter
    subscriptions.EntityData.YangName = "subscriptions"
    subscriptions.EntityData.BundleName = "openconfig"
    subscriptions.EntityData.ParentYangName = "telemetry-system"
    subscriptions.EntityData.SegmentPath = "subscriptions"
    subscriptions.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/" + subscriptions.EntityData.SegmentPath
    subscriptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subscriptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subscriptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subscriptions.EntityData.Children = types.NewOrderedMap()
    subscriptions.EntityData.Children.Append("persistent", types.YChild{"Persistent", &subscriptions.Persistent})
    subscriptions.EntityData.Children.Append("dynamic", types.YChild{"Dynamic", &subscriptions.Dynamic})
    subscriptions.EntityData.Leafs = types.NewOrderedMap()

    subscriptions.EntityData.YListKeys = []string {}

    return &(subscriptions.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent
// This container holds information relating to persistent
// telemetry subscriptions. A persistent telemetry
// subscription is configued locally on the device through
// configuration, and is persistent across device restarts or
// other redundancy changes.
type TelemetrySystem_Subscriptions_Persistent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of telemetry subscriptions. A telemetry subscription consists of a set
    // of collection destinations, stream attributes, and associated paths to
    // state information in the model (sensor data). The type is slice of
    // TelemetrySystem_Subscriptions_Persistent_Subscription.
    Subscription []*TelemetrySystem_Subscriptions_Persistent_Subscription
}

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetEntityData() *types.CommonEntityData {
    persistent.EntityData.YFilter = persistent.YFilter
    persistent.EntityData.YangName = "persistent"
    persistent.EntityData.BundleName = "openconfig"
    persistent.EntityData.ParentYangName = "subscriptions"
    persistent.EntityData.SegmentPath = "persistent"
    persistent.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/" + persistent.EntityData.SegmentPath
    persistent.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    persistent.EntityData.NamespaceTable = openconfig.GetNamespaces()
    persistent.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    persistent.EntityData.Children = types.NewOrderedMap()
    persistent.EntityData.Children.Append("subscription", types.YChild{"Subscription", nil})
    for i := range persistent.Subscription {
        persistent.EntityData.Children.Append(types.GetSegmentPath(persistent.Subscription[i]), types.YChild{"Subscription", persistent.Subscription[i]})
    }
    persistent.EntityData.Leafs = types.NewOrderedMap()

    persistent.EntityData.YListKeys = []string {}

    return &(persistent.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription
// List of telemetry subscriptions. A telemetry
// subscription consists of a set of collection
// destinations, stream attributes, and associated paths to
// state information in the model (sensor data)
type TelemetrySystem_Subscriptions_Persistent_Subscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the identifier of the subscription
    // itself. The id will be the handle to refer to the subscription once
    // created. The type is string with range: 0..18446744073709551615. Refers to
    // telemetry.TelemetrySystem_Subscriptions_Persistent_Subscription_Config_SubscriptionId
    SubscriptionId interface{}

    // Config parameters relating to the telemetry subscriptions on the local
    // device.
    Config TelemetrySystem_Subscriptions_Persistent_Subscription_Config

    // State parameters relating to the telemetry subscriptions on the local
    // device.
    State TelemetrySystem_Subscriptions_Persistent_Subscription_State

    // A sensor profile is a set of sensor groups or individual sensor paths which
    // are associated with a telemetry subscription. This is the source of the
    // telemetry data for the subscription to send to the defined collectors.
    SensorProfiles TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles

    // A subscription may specify destination addresses. If the subscription
    // supplies destination addresses, the network element will be the initiator
    // of the telemetry streaming, sending it to the destination(s) specified.  If
    // the destination set is omitted, the subscription preconfigures certain
    // elements such as paths and sample intervals under a specified subscription
    // ID. In this case, the network element will NOT initiate an outbound
    // connection for telemetry, but will wait for an inbound connection from a
    // network management system.  It is expected that the network management
    // system connecting to the network element will reference the preconfigured
    // subscription ID when initiating a subscription.
    DestinationGroups TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups
}

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "openconfig"
    subscription.EntityData.ParentYangName = "persistent"
    subscription.EntityData.SegmentPath = "subscription" + types.AddKeyToken(subscription.SubscriptionId, "subscription-id")
    subscription.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/" + subscription.EntityData.SegmentPath
    subscription.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subscription.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subscription.EntityData.Children = types.NewOrderedMap()
    subscription.EntityData.Children.Append("config", types.YChild{"Config", &subscription.Config})
    subscription.EntityData.Children.Append("state", types.YChild{"State", &subscription.State})
    subscription.EntityData.Children.Append("sensor-profiles", types.YChild{"SensorProfiles", &subscription.SensorProfiles})
    subscription.EntityData.Children.Append("destination-groups", types.YChild{"DestinationGroups", &subscription.DestinationGroups})
    subscription.EntityData.Leafs = types.NewOrderedMap()
    subscription.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", subscription.SubscriptionId})

    subscription.EntityData.YListKeys = []string {"SubscriptionId"}

    return &(subscription.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_Config
// Config parameters relating to the telemetry
// subscriptions on the local device
type TelemetrySystem_Subscriptions_Persistent_Subscription_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifer of the telemetry subscription. Will be used by configuration
    // operations needing to modify or delete the telemetry subscription. The type
    // is interface{} with range: 0..18446744073709551615.
    SubscriptionId interface{}

    // The IP address which will be the source of packets from the device to a
    // telemetry collector destination. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalSourceAddress interface{}

    // DSCP marking of packets generated by the telemetry subsystem on the network
    // device. The type is interface{} with range: 0..63.
    OriginatedQosMarking interface{}
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "subscription"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", config.SubscriptionId})
    config.EntityData.Leafs.Append("local-source-address", types.YLeaf{"LocalSourceAddress", config.LocalSourceAddress})
    config.EntityData.Leafs.Append("originated-qos-marking", types.YLeaf{"OriginatedQosMarking", config.OriginatedQosMarking})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_State
// State parameters relating to the telemetry
// subscriptions on the local device
type TelemetrySystem_Subscriptions_Persistent_Subscription_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifer of the telemetry subscription. Will be used by configuration
    // operations needing to modify or delete the telemetry subscription. The type
    // is interface{} with range: 0..18446744073709551615.
    SubscriptionId interface{}

    // The IP address which will be the source of packets from the device to a
    // telemetry collector destination. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalSourceAddress interface{}

    // DSCP marking of packets generated by the telemetry subsystem on the network
    // device. The type is interface{} with range: 0..63.
    OriginatedQosMarking interface{}
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "subscription"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", state.SubscriptionId})
    state.EntityData.Leafs.Append("local-source-address", types.YLeaf{"LocalSourceAddress", state.LocalSourceAddress})
    state.EntityData.Leafs.Append("originated-qos-marking", types.YLeaf{"OriginatedQosMarking", state.OriginatedQosMarking})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles
// A sensor profile is a set of sensor groups or
// individual sensor paths which are associated with a
// telemetry subscription. This is the source of the
// telemetry data for the subscription to send to the
// defined collectors.
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of telemetry sensor groups used in the subscription. The type is slice
    // of
    // TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile.
    SensorProfile []*TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile
}

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetEntityData() *types.CommonEntityData {
    sensorProfiles.EntityData.YFilter = sensorProfiles.YFilter
    sensorProfiles.EntityData.YangName = "sensor-profiles"
    sensorProfiles.EntityData.BundleName = "openconfig"
    sensorProfiles.EntityData.ParentYangName = "subscription"
    sensorProfiles.EntityData.SegmentPath = "sensor-profiles"
    sensorProfiles.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/" + sensorProfiles.EntityData.SegmentPath
    sensorProfiles.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorProfiles.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorProfiles.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorProfiles.EntityData.Children = types.NewOrderedMap()
    sensorProfiles.EntityData.Children.Append("sensor-profile", types.YChild{"SensorProfile", nil})
    for i := range sensorProfiles.SensorProfile {
        sensorProfiles.EntityData.Children.Append(types.GetSegmentPath(sensorProfiles.SensorProfile[i]), types.YChild{"SensorProfile", sensorProfiles.SensorProfile[i]})
    }
    sensorProfiles.EntityData.Leafs = types.NewOrderedMap()

    sensorProfiles.EntityData.YListKeys = []string {}

    return &(sensorProfiles.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile
// List of telemetry sensor groups used
// in the subscription
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the telemetry sensor group name. The
    // type is string. Refers to
    // telemetry.TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config_SensorGroup
    SensorGroup interface{}

    // Configuration parameters related to the sensor profile for a subscription.
    Config TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config

    // State information relating to the sensor profile for a subscription.
    State TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State
}

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetEntityData() *types.CommonEntityData {
    sensorProfile.EntityData.YFilter = sensorProfile.YFilter
    sensorProfile.EntityData.YangName = "sensor-profile"
    sensorProfile.EntityData.BundleName = "openconfig"
    sensorProfile.EntityData.ParentYangName = "sensor-profiles"
    sensorProfile.EntityData.SegmentPath = "sensor-profile" + types.AddKeyToken(sensorProfile.SensorGroup, "sensor-group")
    sensorProfile.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/sensor-profiles/" + sensorProfile.EntityData.SegmentPath
    sensorProfile.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorProfile.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorProfile.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorProfile.EntityData.Children = types.NewOrderedMap()
    sensorProfile.EntityData.Children.Append("config", types.YChild{"Config", &sensorProfile.Config})
    sensorProfile.EntityData.Children.Append("state", types.YChild{"State", &sensorProfile.State})
    sensorProfile.EntityData.Leafs = types.NewOrderedMap()
    sensorProfile.EntityData.Leafs.Append("sensor-group", types.YLeaf{"SensorGroup", sensorProfile.SensorGroup})

    sensorProfile.EntityData.YListKeys = []string {"SensorGroup"}

    return &(sensorProfile.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config
// Configuration parameters related to the sensor
// profile for a subscription
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the sensor group which is used in the profile. The type is
    // string. Refers to
    // telemetry.TelemetrySystem_SensorGroups_SensorGroup_Config_SensorGroupId
    SensorGroup interface{}

    // Time in milliseconds between the device's sample of a telemetry data
    // source. For example, setting this to 100 would require the local device to
    // collect the telemetry data every 100 milliseconds. There can be latency or
    // jitter in transmitting the data, but the sample must occur at the specified
    // interval.  The timestamp must reflect the actual time when the data was
    // sampled, not simply the previous sample timestamp + sample-interval.  If
    // sample-interval is set to 0, the telemetry sensor becomes event based. The
    // sensor must then emit data upon every change of the underlying data source.
    // The type is interface{} with range: 0..18446744073709551615.
    SampleInterval interface{}

    // Maximum time interval in seconds that may pass between updates from a
    // device to a telemetry collector. If this interval expires, but there is no
    // updated data to send (such as if suppress_updates has been configured), the
    // device must send a telemetry message to the collector. The type is
    // interface{} with range: 0..18446744073709551615.
    HeartbeatInterval interface{}

    // Boolean flag to control suppression of redundant telemetry updates to the
    // collector platform. If this flag is set to TRUE, then the collector will
    // only send an update at the configured interval if a subscribed data value
    // has changed. Otherwise, the device will not send an update to the collector
    // until expiration of the heartbeat interval. The type is bool.
    SuppressRedundant interface{}
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "sensor-profile"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/sensor-profiles/sensor-profile/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("sensor-group", types.YLeaf{"SensorGroup", config.SensorGroup})
    config.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", config.SampleInterval})
    config.EntityData.Leafs.Append("heartbeat-interval", types.YLeaf{"HeartbeatInterval", config.HeartbeatInterval})
    config.EntityData.Leafs.Append("suppress-redundant", types.YLeaf{"SuppressRedundant", config.SuppressRedundant})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State
// State information relating to the sensor profile
// for a subscription
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the sensor group which is used in the profile. The type is
    // string. Refers to
    // telemetry.TelemetrySystem_SensorGroups_SensorGroup_Config_SensorGroupId
    SensorGroup interface{}

    // Time in milliseconds between the device's sample of a telemetry data
    // source. For example, setting this to 100 would require the local device to
    // collect the telemetry data every 100 milliseconds. There can be latency or
    // jitter in transmitting the data, but the sample must occur at the specified
    // interval.  The timestamp must reflect the actual time when the data was
    // sampled, not simply the previous sample timestamp + sample-interval.  If
    // sample-interval is set to 0, the telemetry sensor becomes event based. The
    // sensor must then emit data upon every change of the underlying data source.
    // The type is interface{} with range: 0..18446744073709551615.
    SampleInterval interface{}

    // Maximum time interval in seconds that may pass between updates from a
    // device to a telemetry collector. If this interval expires, but there is no
    // updated data to send (such as if suppress_updates has been configured), the
    // device must send a telemetry message to the collector. The type is
    // interface{} with range: 0..18446744073709551615.
    HeartbeatInterval interface{}

    // Boolean flag to control suppression of redundant telemetry updates to the
    // collector platform. If this flag is set to TRUE, then the collector will
    // only send an update at the configured interval if a subscribed data value
    // has changed. Otherwise, the device will not send an update to the collector
    // until expiration of the heartbeat interval. The type is bool.
    SuppressRedundant interface{}
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "sensor-profile"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/sensor-profiles/sensor-profile/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("sensor-group", types.YLeaf{"SensorGroup", state.SensorGroup})
    state.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", state.SampleInterval})
    state.EntityData.Leafs.Append("heartbeat-interval", types.YLeaf{"HeartbeatInterval", state.HeartbeatInterval})
    state.EntityData.Leafs.Append("suppress-redundant", types.YLeaf{"SuppressRedundant", state.SuppressRedundant})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups
// A subscription may specify destination addresses.
// If the subscription supplies destination addresses,
// the network element will be the initiator of the
// telemetry streaming, sending it to the destination(s)
// specified.
// 
// If the destination set is omitted, the subscription
// preconfigures certain elements such as paths and
// sample intervals under a specified subscription ID.
// In this case, the network element will NOT initiate an
// outbound connection for telemetry, but will wait for
// an inbound connection from a network management
// system.
// 
// It is expected that the network management system
// connecting to the network element will reference
// the preconfigured subscription ID when initiating
// a subscription.
type TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifier of the previously defined destination group. The type is slice
    // of
    // TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup.
    DestinationGroup []*TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup
}

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetEntityData() *types.CommonEntityData {
    destinationGroups.EntityData.YFilter = destinationGroups.YFilter
    destinationGroups.EntityData.YangName = "destination-groups"
    destinationGroups.EntityData.BundleName = "openconfig"
    destinationGroups.EntityData.ParentYangName = "subscription"
    destinationGroups.EntityData.SegmentPath = "destination-groups"
    destinationGroups.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/" + destinationGroups.EntityData.SegmentPath
    destinationGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    destinationGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    destinationGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    destinationGroups.EntityData.Children = types.NewOrderedMap()
    destinationGroups.EntityData.Children.Append("destination-group", types.YChild{"DestinationGroup", nil})
    for i := range destinationGroups.DestinationGroup {
        destinationGroups.EntityData.Children.Append(types.GetSegmentPath(destinationGroups.DestinationGroup[i]), types.YChild{"DestinationGroup", destinationGroups.DestinationGroup[i]})
    }
    destinationGroups.EntityData.Leafs = types.NewOrderedMap()

    destinationGroups.EntityData.YListKeys = []string {}

    return &(destinationGroups.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup
// Identifier of the previously defined destination
// group
type TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The destination group id references a configured
    // group of destinations for the telemetry stream. The type is string. Refers
    // to
    // telemetry.TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config_GroupId
    GroupId interface{}

    // Configuration parameters related to telemetry destinations.
    Config TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config

    // State information related to telemetry destinations.
    State TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State
}

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetEntityData() *types.CommonEntityData {
    destinationGroup.EntityData.YFilter = destinationGroup.YFilter
    destinationGroup.EntityData.YangName = "destination-group"
    destinationGroup.EntityData.BundleName = "openconfig"
    destinationGroup.EntityData.ParentYangName = "destination-groups"
    destinationGroup.EntityData.SegmentPath = "destination-group" + types.AddKeyToken(destinationGroup.GroupId, "group-id")
    destinationGroup.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/destination-groups/" + destinationGroup.EntityData.SegmentPath
    destinationGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    destinationGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    destinationGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    destinationGroup.EntityData.Children = types.NewOrderedMap()
    destinationGroup.EntityData.Children.Append("config", types.YChild{"Config", &destinationGroup.Config})
    destinationGroup.EntityData.Children.Append("state", types.YChild{"State", &destinationGroup.State})
    destinationGroup.EntityData.Leafs = types.NewOrderedMap()
    destinationGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", destinationGroup.GroupId})

    destinationGroup.EntityData.YListKeys = []string {"GroupId"}

    return &(destinationGroup.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config
// Configuration parameters related to telemetry
// destinations.
type TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The destination group id references a reusable group of destination
    // addresses and ports for the telemetry stream. The type is string. Refers to
    // telemetry.TelemetrySystem_DestinationGroups_DestinationGroup_GroupId
    GroupId interface{}
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "destination-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/destination-groups/destination-group/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", config.GroupId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State
// State information related to telemetry
// destinations
type TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The destination group id references a reusable group of destination
    // addresses and ports for the telemetry stream. The type is string. Refers to
    // telemetry.TelemetrySystem_DestinationGroups_DestinationGroup_GroupId
    GroupId interface{}
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "destination-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/persistent/subscription/destination-groups/destination-group/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", state.GroupId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_Subscriptions_Dynamic
// This container holds information relating to dynamic
// telemetry subscriptions. A dynamic subscription is
// typically configured through an RPC channel, and does not
// persist across device restarts, or if the RPC channel is
// reset or otherwise torn down.
type TelemetrySystem_Subscriptions_Dynamic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List representation of telemetry subscriptions that are configured via an
    // inline RPC, otherwise known as dynamic telemetry subscriptions. The type is
    // slice of TelemetrySystem_Subscriptions_Dynamic_Subscription.
    Subscription []*TelemetrySystem_Subscriptions_Dynamic_Subscription
}

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetEntityData() *types.CommonEntityData {
    dynamic.EntityData.YFilter = dynamic.YFilter
    dynamic.EntityData.YangName = "dynamic"
    dynamic.EntityData.BundleName = "openconfig"
    dynamic.EntityData.ParentYangName = "subscriptions"
    dynamic.EntityData.SegmentPath = "dynamic"
    dynamic.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/" + dynamic.EntityData.SegmentPath
    dynamic.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    dynamic.EntityData.NamespaceTable = openconfig.GetNamespaces()
    dynamic.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    dynamic.EntityData.Children = types.NewOrderedMap()
    dynamic.EntityData.Children.Append("subscription", types.YChild{"Subscription", nil})
    for i := range dynamic.Subscription {
        dynamic.EntityData.Children.Append(types.GetSegmentPath(dynamic.Subscription[i]), types.YChild{"Subscription", dynamic.Subscription[i]})
    }
    dynamic.EntityData.Leafs = types.NewOrderedMap()

    dynamic.EntityData.YListKeys = []string {}

    return &(dynamic.EntityData)
}

// TelemetrySystem_Subscriptions_Dynamic_Subscription
// List representation of telemetry subscriptions that
// are configured via an inline RPC, otherwise known
// as dynamic telemetry subscriptions.
type TelemetrySystem_Subscriptions_Dynamic_Subscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the identifier of the subscription
    // itself. The id will be the handle to refer to the subscription once
    // created. The type is string with range: 0..18446744073709551615. Refers to
    // telemetry.TelemetrySystem_Subscriptions_Dynamic_Subscription_State_SubscriptionId
    SubscriptionId interface{}

    // State information relating to dynamic telemetry subscriptions.
    State TelemetrySystem_Subscriptions_Dynamic_Subscription_State

    // Top level container to hold a set of sensor paths grouped together.
    SensorPaths TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths
}

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "openconfig"
    subscription.EntityData.ParentYangName = "dynamic"
    subscription.EntityData.SegmentPath = "subscription" + types.AddKeyToken(subscription.SubscriptionId, "subscription-id")
    subscription.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/dynamic/" + subscription.EntityData.SegmentPath
    subscription.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subscription.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subscription.EntityData.Children = types.NewOrderedMap()
    subscription.EntityData.Children.Append("state", types.YChild{"State", &subscription.State})
    subscription.EntityData.Children.Append("sensor-paths", types.YChild{"SensorPaths", &subscription.SensorPaths})
    subscription.EntityData.Leafs = types.NewOrderedMap()
    subscription.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", subscription.SubscriptionId})

    subscription.EntityData.YListKeys = []string {"SubscriptionId"}

    return &(subscription.EntityData)
}

// TelemetrySystem_Subscriptions_Dynamic_Subscription_State
// State information relating to dynamic telemetry
// subscriptions.
type TelemetrySystem_Subscriptions_Dynamic_Subscription_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifer of the telemetry subscription. Will be used by configuration
    // operations needing to modify or delete the telemetry subscription. The type
    // is interface{} with range: 0..18446744073709551615.
    SubscriptionId interface{}

    // IP address of the telemetry stream destination. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Protocol (udp or tcp) port number for the telemetry stream destination. The
    // type is interface{} with range: 0..65535.
    DestinationPort interface{}

    // Protocol used to transmit telemetry data to the collector. The type is
    // TelemetryStreamProtocol.
    DestinationProtocol interface{}

    // Time in milliseconds between the device's sample of a telemetry data
    // source. For example, setting this to 100 would require the local device to
    // collect the telemetry data every 100 milliseconds. There can be latency or
    // jitter in transmitting the data, but the sample must occur at the specified
    // interval.  The timestamp must reflect the actual time when the data was
    // sampled, not simply the previous sample timestamp + sample-interval.  If
    // sample-interval is set to 0, the telemetry sensor becomes event based. The
    // sensor must then emit data upon every change of the underlying data source.
    // The type is interface{} with range: 0..18446744073709551615.
    SampleInterval interface{}

    // Maximum time interval in seconds that may pass between updates from a
    // device to a telemetry collector. If this interval expires, but there is no
    // updated data to send (such as if suppress_updates has been configured), the
    // device must send a telemetry message to the collector. The type is
    // interface{} with range: 0..18446744073709551615.
    HeartbeatInterval interface{}

    // Boolean flag to control suppression of redundant telemetry updates to the
    // collector platform. If this flag is set to TRUE, then the collector will
    // only send an update at the configured interval if a subscribed data value
    // has changed. Otherwise, the device will not send an update to the collector
    // until expiration of the heartbeat interval. The type is bool.
    SuppressRedundant interface{}

    // DSCP marking of packets generated by the telemetry subsystem on the network
    // device. The type is interface{} with range: 0..63.
    OriginatedQosMarking interface{}
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "subscription"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/dynamic/subscription/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", state.SubscriptionId})
    state.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", state.DestinationAddress})
    state.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", state.DestinationPort})
    state.EntityData.Leafs.Append("destination-protocol", types.YLeaf{"DestinationProtocol", state.DestinationProtocol})
    state.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", state.SampleInterval})
    state.EntityData.Leafs.Append("heartbeat-interval", types.YLeaf{"HeartbeatInterval", state.HeartbeatInterval})
    state.EntityData.Leafs.Append("suppress-redundant", types.YLeaf{"SuppressRedundant", state.SuppressRedundant})
    state.EntityData.Leafs.Append("originated-qos-marking", types.YLeaf{"OriginatedQosMarking", state.OriginatedQosMarking})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths
// Top level container to hold a set of sensor
// paths grouped together
type TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of paths in the model which together comprise a sensor grouping.
    // Filters for each path to exclude items are also provided. The type is slice
    // of
    // TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath.
    SensorPath []*TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath
}

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetEntityData() *types.CommonEntityData {
    sensorPaths.EntityData.YFilter = sensorPaths.YFilter
    sensorPaths.EntityData.YangName = "sensor-paths"
    sensorPaths.EntityData.BundleName = "openconfig"
    sensorPaths.EntityData.ParentYangName = "subscription"
    sensorPaths.EntityData.SegmentPath = "sensor-paths"
    sensorPaths.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/dynamic/subscription/" + sensorPaths.EntityData.SegmentPath
    sensorPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorPaths.EntityData.Children = types.NewOrderedMap()
    sensorPaths.EntityData.Children.Append("sensor-path", types.YChild{"SensorPath", nil})
    for i := range sensorPaths.SensorPath {
        sensorPaths.EntityData.Children.Append(types.GetSegmentPath(sensorPaths.SensorPath[i]), types.YChild{"SensorPath", sensorPaths.SensorPath[i]})
    }
    sensorPaths.EntityData.Leafs = types.NewOrderedMap()

    sensorPaths.EntityData.YListKeys = []string {}

    return &(sensorPaths.EntityData)
}

// TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath
// List of paths in the model which together
// comprise a sensor grouping. Filters for each path
// to exclude items are also provided.
type TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the path of interest. The type is
    // string. Refers to
    // telemetry.TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State_Path
    Path interface{}

    // State information for a dynamic subscription's paths of interest.
    State TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State
}

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetEntityData() *types.CommonEntityData {
    sensorPath.EntityData.YFilter = sensorPath.YFilter
    sensorPath.EntityData.YangName = "sensor-path"
    sensorPath.EntityData.BundleName = "openconfig"
    sensorPath.EntityData.ParentYangName = "sensor-paths"
    sensorPath.EntityData.SegmentPath = "sensor-path" + types.AddKeyToken(sensorPath.Path, "path")
    sensorPath.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/dynamic/subscription/sensor-paths/" + sensorPath.EntityData.SegmentPath
    sensorPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sensorPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sensorPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sensorPath.EntityData.Children = types.NewOrderedMap()
    sensorPath.EntityData.Children.Append("state", types.YChild{"State", &sensorPath.State})
    sensorPath.EntityData.Leafs = types.NewOrderedMap()
    sensorPath.EntityData.Leafs.Append("path", types.YLeaf{"Path", sensorPath.Path})

    sensorPath.EntityData.YListKeys = []string {"Path"}

    return &(sensorPath.EntityData)
}

// TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State
// State information for a dynamic subscription's
// paths of interest
type TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path to a section of operational state of interest (the sensor). The type
    // is string.
    Path interface{}

    // Filter to exclude certain values out of the state values. The type is
    // string.
    ExcludeFilter interface{}
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "sensor-path"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-telemetry:telemetry-system/subscriptions/dynamic/subscription/sensor-paths/sensor-path/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("path", types.YLeaf{"Path", state.Path})
    state.EntityData.Leafs.Append("exclude-filter", types.YLeaf{"ExcludeFilter", state.ExcludeFilter})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

