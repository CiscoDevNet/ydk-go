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
    parent types.Entity
    YFilter yfilter.YFilter

    // Top level container for sensor-groups.
    SensorGroups TelemetrySystem_SensorGroups

    // Top level container for destination group configuration and state.
    DestinationGroups TelemetrySystem_DestinationGroups

    // This container holds information for both persistent and dynamic telemetry
    // subscriptions.
    Subscriptions TelemetrySystem_Subscriptions
}

func (telemetrySystem *TelemetrySystem) GetFilter() yfilter.YFilter { return telemetrySystem.YFilter }

func (telemetrySystem *TelemetrySystem) SetFilter(yf yfilter.YFilter) { telemetrySystem.YFilter = yf }

func (telemetrySystem *TelemetrySystem) GetGoName(yname string) string {
    if yname == "sensor-groups" { return "SensorGroups" }
    if yname == "destination-groups" { return "DestinationGroups" }
    if yname == "subscriptions" { return "Subscriptions" }
    return ""
}

func (telemetrySystem *TelemetrySystem) GetSegmentPath() string {
    return "openconfig-telemetry:telemetry-system"
}

func (telemetrySystem *TelemetrySystem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-groups" {
        return &telemetrySystem.SensorGroups
    }
    if childYangName == "destination-groups" {
        return &telemetrySystem.DestinationGroups
    }
    if childYangName == "subscriptions" {
        return &telemetrySystem.Subscriptions
    }
    return nil
}

func (telemetrySystem *TelemetrySystem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensor-groups"] = &telemetrySystem.SensorGroups
    children["destination-groups"] = &telemetrySystem.DestinationGroups
    children["subscriptions"] = &telemetrySystem.Subscriptions
    return children
}

func (telemetrySystem *TelemetrySystem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (telemetrySystem *TelemetrySystem) GetBundleName() string { return "openconfig" }

func (telemetrySystem *TelemetrySystem) GetYangName() string { return "telemetry-system" }

func (telemetrySystem *TelemetrySystem) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (telemetrySystem *TelemetrySystem) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (telemetrySystem *TelemetrySystem) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (telemetrySystem *TelemetrySystem) SetParent(parent types.Entity) { telemetrySystem.parent = parent }

func (telemetrySystem *TelemetrySystem) GetParent() types.Entity { return telemetrySystem.parent }

func (telemetrySystem *TelemetrySystem) GetParentYangName() string { return "openconfig-telemetry" }

// TelemetrySystem_SensorGroups
// Top level container for sensor-groups.
type TelemetrySystem_SensorGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of telemetry sensory groups on the local system, where a sensor
    // grouping represents a resuable grouping of multiple paths and exclude
    // filters. The type is slice of TelemetrySystem_SensorGroups_SensorGroup.
    SensorGroup []TelemetrySystem_SensorGroups_SensorGroup
}

func (sensorGroups *TelemetrySystem_SensorGroups) GetFilter() yfilter.YFilter { return sensorGroups.YFilter }

func (sensorGroups *TelemetrySystem_SensorGroups) SetFilter(yf yfilter.YFilter) { sensorGroups.YFilter = yf }

func (sensorGroups *TelemetrySystem_SensorGroups) GetGoName(yname string) string {
    if yname == "sensor-group" { return "SensorGroup" }
    return ""
}

func (sensorGroups *TelemetrySystem_SensorGroups) GetSegmentPath() string {
    return "sensor-groups"
}

func (sensorGroups *TelemetrySystem_SensorGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-group" {
        for _, c := range sensorGroups.SensorGroup {
            if sensorGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_SensorGroups_SensorGroup{}
        sensorGroups.SensorGroup = append(sensorGroups.SensorGroup, child)
        return &sensorGroups.SensorGroup[len(sensorGroups.SensorGroup)-1]
    }
    return nil
}

func (sensorGroups *TelemetrySystem_SensorGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorGroups.SensorGroup {
        children[sensorGroups.SensorGroup[i].GetSegmentPath()] = &sensorGroups.SensorGroup[i]
    }
    return children
}

func (sensorGroups *TelemetrySystem_SensorGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorGroups *TelemetrySystem_SensorGroups) GetBundleName() string { return "openconfig" }

func (sensorGroups *TelemetrySystem_SensorGroups) GetYangName() string { return "sensor-groups" }

func (sensorGroups *TelemetrySystem_SensorGroups) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorGroups *TelemetrySystem_SensorGroups) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorGroups *TelemetrySystem_SensorGroups) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorGroups *TelemetrySystem_SensorGroups) SetParent(parent types.Entity) { sensorGroups.parent = parent }

func (sensorGroups *TelemetrySystem_SensorGroups) GetParent() types.Entity { return sensorGroups.parent }

func (sensorGroups *TelemetrySystem_SensorGroups) GetParentYangName() string { return "telemetry-system" }

// TelemetrySystem_SensorGroups_SensorGroup
// List of telemetry sensory groups on the local
// system, where a sensor grouping represents a resuable
// grouping of multiple paths and exclude filters.
type TelemetrySystem_SensorGroups_SensorGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetFilter() yfilter.YFilter { return sensorGroup.YFilter }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) SetFilter(yf yfilter.YFilter) { sensorGroup.YFilter = yf }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetGoName(yname string) string {
    if yname == "sensor-group-id" { return "SensorGroupId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "sensor-paths" { return "SensorPaths" }
    return ""
}

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetSegmentPath() string {
    return "sensor-group" + "[sensor-group-id='" + fmt.Sprintf("%v", sensorGroup.SensorGroupId) + "']"
}

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &sensorGroup.Config
    }
    if childYangName == "state" {
        return &sensorGroup.State
    }
    if childYangName == "sensor-paths" {
        return &sensorGroup.SensorPaths
    }
    return nil
}

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &sensorGroup.Config
    children["state"] = &sensorGroup.State
    children["sensor-paths"] = &sensorGroup.SensorPaths
    return children
}

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group-id"] = sensorGroup.SensorGroupId
    return leafs
}

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetBundleName() string { return "openconfig" }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetYangName() string { return "sensor-group" }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) SetParent(parent types.Entity) { sensorGroup.parent = parent }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetParent() types.Entity { return sensorGroup.parent }

func (sensorGroup *TelemetrySystem_SensorGroups_SensorGroup) GetParentYangName() string { return "sensor-groups" }

// TelemetrySystem_SensorGroups_SensorGroup_Config
// Configuration parameters relating to the
// telemetry sensor grouping
type TelemetrySystem_SensorGroups_SensorGroup_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name or identifier for the sensor group itself. Will be referenced by other
    // configuration specifying a sensor group. The type is string.
    SensorGroupId interface{}
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetGoName(yname string) string {
    if yname == "sensor-group-id" { return "SensorGroupId" }
    return ""
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group-id"] = config.SensorGroupId
    return leafs
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetBundleName() string { return "openconfig" }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetYangName() string { return "config" }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetParent() types.Entity { return config.parent }

func (config *TelemetrySystem_SensorGroups_SensorGroup_Config) GetParentYangName() string { return "sensor-group" }

// TelemetrySystem_SensorGroups_SensorGroup_State
// State information relating to the telemetry
// sensor group
type TelemetrySystem_SensorGroups_SensorGroup_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name or identifier for the sensor group itself. Will be referenced by other
    // configuration specifying a sensor group. The type is string.
    SensorGroupId interface{}
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetGoName(yname string) string {
    if yname == "sensor-group-id" { return "SensorGroupId" }
    return ""
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group-id"] = state.SensorGroupId
    return leafs
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_SensorGroups_SensorGroup_State) GetParentYangName() string { return "sensor-group" }

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths
// Top level container to hold a set of sensor
// paths grouped together
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of paths in the model which together comprise a sensor grouping.
    // Filters for each path to exclude items are also provided. The type is slice
    // of TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath.
    SensorPath []TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath
}

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetFilter() yfilter.YFilter { return sensorPaths.YFilter }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) SetFilter(yf yfilter.YFilter) { sensorPaths.YFilter = yf }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetGoName(yname string) string {
    if yname == "sensor-path" { return "SensorPath" }
    return ""
}

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetSegmentPath() string {
    return "sensor-paths"
}

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-path" {
        for _, c := range sensorPaths.SensorPath {
            if sensorPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath{}
        sensorPaths.SensorPath = append(sensorPaths.SensorPath, child)
        return &sensorPaths.SensorPath[len(sensorPaths.SensorPath)-1]
    }
    return nil
}

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorPaths.SensorPath {
        children[sensorPaths.SensorPath[i].GetSegmentPath()] = &sensorPaths.SensorPath[i]
    }
    return children
}

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetBundleName() string { return "openconfig" }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetYangName() string { return "sensor-paths" }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) SetParent(parent types.Entity) { sensorPaths.parent = parent }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetParent() types.Entity { return sensorPaths.parent }

func (sensorPaths *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths) GetParentYangName() string { return "sensor-group" }

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath
// List of paths in the model which together
// comprise a sensor grouping. Filters for each path
// to exclude items are also provided.
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetFilter() yfilter.YFilter { return sensorPath.YFilter }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) SetFilter(yf yfilter.YFilter) { sensorPath.YFilter = yf }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetSegmentPath() string {
    return "sensor-path" + "[path='" + fmt.Sprintf("%v", sensorPath.Path) + "']"
}

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &sensorPath.Config
    }
    if childYangName == "state" {
        return &sensorPath.State
    }
    return nil
}

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &sensorPath.Config
    children["state"] = &sensorPath.State
    return children
}

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = sensorPath.Path
    return leafs
}

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetBundleName() string { return "openconfig" }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetYangName() string { return "sensor-path" }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) SetParent(parent types.Entity) { sensorPath.parent = parent }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetParent() types.Entity { return sensorPath.parent }

func (sensorPath *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetParentYangName() string { return "sensor-paths" }

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config
// Configuration parameters to configure a set
// of data model paths as a sensor grouping
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path to a section of operational state of interest (the sensor). The type
    // is string.
    Path interface{}

    // Filter to exclude certain values out of the state values. The type is
    // string.
    ExcludeFilter interface{}
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "exclude-filter" { return "ExcludeFilter" }
    return ""
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetSegmentPath() string {
    return "config"
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = config.Path
    leafs["exclude-filter"] = config.ExcludeFilter
    return leafs
}

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetBundleName() string { return "openconfig" }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetYangName() string { return "config" }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetParent() types.Entity { return config.parent }

func (config *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_Config) GetParentYangName() string { return "sensor-path" }

// TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State
// Configuration parameters to configure a set
// of data model paths as a sensor grouping
type TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path to a section of operational state of interest (the sensor). The type
    // is string.
    Path interface{}

    // Filter to exclude certain values out of the state values. The type is
    // string.
    ExcludeFilter interface{}
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "exclude-filter" { return "ExcludeFilter" }
    return ""
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = state.Path
    leafs["exclude-filter"] = state.ExcludeFilter
    return leafs
}

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_SensorGroups_SensorGroup_SensorPaths_SensorPath_State) GetParentYangName() string { return "sensor-path" }

// TelemetrySystem_DestinationGroups
// Top level container for destination group configuration
// and state.
type TelemetrySystem_DestinationGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of destination-groups. Destination groups allow the reuse of common
    // telemetry destinations across the telemetry configuration. An operator
    // references a set of destinations via the configurable
    // destination-group-identifier.  A destination group may contain one or more
    // telemetry destinations. The type is slice of
    // TelemetrySystem_DestinationGroups_DestinationGroup.
    DestinationGroup []TelemetrySystem_DestinationGroups_DestinationGroup
}

func (destinationGroups *TelemetrySystem_DestinationGroups) GetFilter() yfilter.YFilter { return destinationGroups.YFilter }

func (destinationGroups *TelemetrySystem_DestinationGroups) SetFilter(yf yfilter.YFilter) { destinationGroups.YFilter = yf }

func (destinationGroups *TelemetrySystem_DestinationGroups) GetGoName(yname string) string {
    if yname == "destination-group" { return "DestinationGroup" }
    return ""
}

func (destinationGroups *TelemetrySystem_DestinationGroups) GetSegmentPath() string {
    return "destination-groups"
}

func (destinationGroups *TelemetrySystem_DestinationGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-group" {
        for _, c := range destinationGroups.DestinationGroup {
            if destinationGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_DestinationGroups_DestinationGroup{}
        destinationGroups.DestinationGroup = append(destinationGroups.DestinationGroup, child)
        return &destinationGroups.DestinationGroup[len(destinationGroups.DestinationGroup)-1]
    }
    return nil
}

func (destinationGroups *TelemetrySystem_DestinationGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destinationGroups.DestinationGroup {
        children[destinationGroups.DestinationGroup[i].GetSegmentPath()] = &destinationGroups.DestinationGroup[i]
    }
    return children
}

func (destinationGroups *TelemetrySystem_DestinationGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (destinationGroups *TelemetrySystem_DestinationGroups) GetBundleName() string { return "openconfig" }

func (destinationGroups *TelemetrySystem_DestinationGroups) GetYangName() string { return "destination-groups" }

func (destinationGroups *TelemetrySystem_DestinationGroups) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (destinationGroups *TelemetrySystem_DestinationGroups) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (destinationGroups *TelemetrySystem_DestinationGroups) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (destinationGroups *TelemetrySystem_DestinationGroups) SetParent(parent types.Entity) { destinationGroups.parent = parent }

func (destinationGroups *TelemetrySystem_DestinationGroups) GetParent() types.Entity { return destinationGroups.parent }

func (destinationGroups *TelemetrySystem_DestinationGroups) GetParentYangName() string { return "telemetry-system" }

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
    parent types.Entity
    YFilter yfilter.YFilter

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

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetFilter() yfilter.YFilter { return destinationGroup.YFilter }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) SetFilter(yf yfilter.YFilter) { destinationGroup.YFilter = yf }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "destinations" { return "Destinations" }
    return ""
}

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetSegmentPath() string {
    return "destination-group" + "[group-id='" + fmt.Sprintf("%v", destinationGroup.GroupId) + "']"
}

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &destinationGroup.Config
    }
    if childYangName == "state" {
        return &destinationGroup.State
    }
    if childYangName == "destinations" {
        return &destinationGroup.Destinations
    }
    return nil
}

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &destinationGroup.Config
    children["state"] = &destinationGroup.State
    children["destinations"] = &destinationGroup.Destinations
    return children
}

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = destinationGroup.GroupId
    return leafs
}

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetBundleName() string { return "openconfig" }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetYangName() string { return "destination-group" }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) SetParent(parent types.Entity) { destinationGroup.parent = parent }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetParent() types.Entity { return destinationGroup.parent }

func (destinationGroup *TelemetrySystem_DestinationGroups_DestinationGroup) GetParentYangName() string { return "destination-groups" }

// TelemetrySystem_DestinationGroups_DestinationGroup_Config
// Top level config container for destination groups
type TelemetrySystem_DestinationGroups_DestinationGroup_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique identifier for the destination group. The type is string.
    GroupId interface{}
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = config.GroupId
    return leafs
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetBundleName() string { return "openconfig" }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetYangName() string { return "config" }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetParent() types.Entity { return config.parent }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Config) GetParentYangName() string { return "destination-group" }

// TelemetrySystem_DestinationGroups_DestinationGroup_State
// Top level state container for destination groups
type TelemetrySystem_DestinationGroups_DestinationGroup_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique identifier for destination group. The type is string.
    GroupId interface{}
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = state.GroupId
    return leafs
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_State) GetParentYangName() string { return "destination-group" }

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations
// The destination container lists the destination
// information such as IP address and port of the
// telemetry messages from the network element.
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of telemetry stream destinations. The type is slice of
    // TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination.
    Destination []TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination
}

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetFilter() yfilter.YFilter { return destinations.YFilter }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) SetFilter(yf yfilter.YFilter) { destinations.YFilter = yf }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    return ""
}

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetSegmentPath() string {
    return "destinations"
}

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        for _, c := range destinations.Destination {
            if destinations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination{}
        destinations.Destination = append(destinations.Destination, child)
        return &destinations.Destination[len(destinations.Destination)-1]
    }
    return nil
}

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destinations.Destination {
        children[destinations.Destination[i].GetSegmentPath()] = &destinations.Destination[i]
    }
    return children
}

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetBundleName() string { return "openconfig" }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetYangName() string { return "destinations" }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) SetParent(parent types.Entity) { destinations.parent = parent }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetParent() types.Entity { return destinations.parent }

func (destinations *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations) GetParentYangName() string { return "destination-group" }

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination
// List of telemetry stream destinations
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetGoName(yname string) string {
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetSegmentPath() string {
    return "destination" + "[destination-address='" + fmt.Sprintf("%v", destination.DestinationAddress) + "']" + "[destination-port='" + fmt.Sprintf("%v", destination.DestinationPort) + "']"
}

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &destination.Config
    }
    if childYangName == "state" {
        return &destination.State
    }
    return nil
}

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &destination.Config
    children["state"] = &destination.State
    return children
}

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-address"] = destination.DestinationAddress
    leafs["destination-port"] = destination.DestinationPort
    return leafs
}

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetBundleName() string { return "openconfig" }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetYangName() string { return "destination" }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetParent() types.Entity { return destination.parent }

func (destination *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination) GetParentYangName() string { return "destinations" }

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config
// Configuration parameters relating to
// telemetry destinations
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config struct {
    parent types.Entity
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

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetGoName(yname string) string {
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "destination-protocol" { return "DestinationProtocol" }
    return ""
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetSegmentPath() string {
    return "config"
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-address"] = config.DestinationAddress
    leafs["destination-port"] = config.DestinationPort
    leafs["destination-protocol"] = config.DestinationProtocol
    return leafs
}

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetBundleName() string { return "openconfig" }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetYangName() string { return "config" }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetParent() types.Entity { return config.parent }

func (config *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_Config) GetParentYangName() string { return "destination" }

// TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State
// State information associated with
// telemetry destinations
type TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State struct {
    parent types.Entity
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

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetGoName(yname string) string {
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "destination-protocol" { return "DestinationProtocol" }
    return ""
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-address"] = state.DestinationAddress
    leafs["destination-port"] = state.DestinationPort
    leafs["destination-protocol"] = state.DestinationProtocol
    return leafs
}

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_DestinationGroups_DestinationGroup_Destinations_Destination_State) GetParentYangName() string { return "destination" }

// TelemetrySystem_Subscriptions
// This container holds information for both persistent
// and dynamic telemetry subscriptions.
type TelemetrySystem_Subscriptions struct {
    parent types.Entity
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

func (subscriptions *TelemetrySystem_Subscriptions) GetFilter() yfilter.YFilter { return subscriptions.YFilter }

func (subscriptions *TelemetrySystem_Subscriptions) SetFilter(yf yfilter.YFilter) { subscriptions.YFilter = yf }

func (subscriptions *TelemetrySystem_Subscriptions) GetGoName(yname string) string {
    if yname == "persistent" { return "Persistent" }
    if yname == "dynamic" { return "Dynamic" }
    return ""
}

func (subscriptions *TelemetrySystem_Subscriptions) GetSegmentPath() string {
    return "subscriptions"
}

func (subscriptions *TelemetrySystem_Subscriptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "persistent" {
        return &subscriptions.Persistent
    }
    if childYangName == "dynamic" {
        return &subscriptions.Dynamic
    }
    return nil
}

func (subscriptions *TelemetrySystem_Subscriptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["persistent"] = &subscriptions.Persistent
    children["dynamic"] = &subscriptions.Dynamic
    return children
}

func (subscriptions *TelemetrySystem_Subscriptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriptions *TelemetrySystem_Subscriptions) GetBundleName() string { return "openconfig" }

func (subscriptions *TelemetrySystem_Subscriptions) GetYangName() string { return "subscriptions" }

func (subscriptions *TelemetrySystem_Subscriptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subscriptions *TelemetrySystem_Subscriptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subscriptions *TelemetrySystem_Subscriptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subscriptions *TelemetrySystem_Subscriptions) SetParent(parent types.Entity) { subscriptions.parent = parent }

func (subscriptions *TelemetrySystem_Subscriptions) GetParent() types.Entity { return subscriptions.parent }

func (subscriptions *TelemetrySystem_Subscriptions) GetParentYangName() string { return "telemetry-system" }

// TelemetrySystem_Subscriptions_Persistent
// This container holds information relating to persistent
// telemetry subscriptions. A persistent telemetry
// subscription is configued locally on the device through
// configuration, and is persistent across device restarts or
// other redundancy changes.
type TelemetrySystem_Subscriptions_Persistent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of telemetry subscriptions. A telemetry subscription consists of a set
    // of collection destinations, stream attributes, and associated paths to
    // state information in the model (sensor data). The type is slice of
    // TelemetrySystem_Subscriptions_Persistent_Subscription.
    Subscription []TelemetrySystem_Subscriptions_Persistent_Subscription
}

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetFilter() yfilter.YFilter { return persistent.YFilter }

func (persistent *TelemetrySystem_Subscriptions_Persistent) SetFilter(yf yfilter.YFilter) { persistent.YFilter = yf }

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetGoName(yname string) string {
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetSegmentPath() string {
    return "persistent"
}

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscription" {
        for _, c := range persistent.Subscription {
            if persistent.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_Subscriptions_Persistent_Subscription{}
        persistent.Subscription = append(persistent.Subscription, child)
        return &persistent.Subscription[len(persistent.Subscription)-1]
    }
    return nil
}

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range persistent.Subscription {
        children[persistent.Subscription[i].GetSegmentPath()] = &persistent.Subscription[i]
    }
    return children
}

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetBundleName() string { return "openconfig" }

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetYangName() string { return "persistent" }

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (persistent *TelemetrySystem_Subscriptions_Persistent) SetParent(parent types.Entity) { persistent.parent = parent }

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetParent() types.Entity { return persistent.parent }

func (persistent *TelemetrySystem_Subscriptions_Persistent) GetParentYangName() string { return "subscriptions" }

// TelemetrySystem_Subscriptions_Persistent_Subscription
// List of telemetry subscriptions. A telemetry
// subscription consists of a set of collection
// destinations, stream attributes, and associated paths to
// state information in the model (sensor data)
type TelemetrySystem_Subscriptions_Persistent_Subscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "sensor-profiles" { return "SensorProfiles" }
    if yname == "destination-groups" { return "DestinationGroups" }
    return ""
}

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetSegmentPath() string {
    return "subscription" + "[subscription-id='" + fmt.Sprintf("%v", subscription.SubscriptionId) + "']"
}

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &subscription.Config
    }
    if childYangName == "state" {
        return &subscription.State
    }
    if childYangName == "sensor-profiles" {
        return &subscription.SensorProfiles
    }
    if childYangName == "destination-groups" {
        return &subscription.DestinationGroups
    }
    return nil
}

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &subscription.Config
    children["state"] = &subscription.State
    children["sensor-profiles"] = &subscription.SensorProfiles
    children["destination-groups"] = &subscription.DestinationGroups
    return children
}

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = subscription.SubscriptionId
    return leafs
}

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetBundleName() string { return "openconfig" }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetYangName() string { return "subscription" }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *TelemetrySystem_Subscriptions_Persistent_Subscription) GetParentYangName() string { return "persistent" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_Config
// Config parameters relating to the telemetry
// subscriptions on the local device
type TelemetrySystem_Subscriptions_Persistent_Subscription_Config struct {
    parent types.Entity
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

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "local-source-address" { return "LocalSourceAddress" }
    if yname == "originated-qos-marking" { return "OriginatedQosMarking" }
    return ""
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetSegmentPath() string {
    return "config"
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = config.SubscriptionId
    leafs["local-source-address"] = config.LocalSourceAddress
    leafs["originated-qos-marking"] = config.OriginatedQosMarking
    return leafs
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetBundleName() string { return "openconfig" }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetYangName() string { return "config" }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetParent() types.Entity { return config.parent }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_Config) GetParentYangName() string { return "subscription" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_State
// State parameters relating to the telemetry
// subscriptions on the local device
type TelemetrySystem_Subscriptions_Persistent_Subscription_State struct {
    parent types.Entity
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

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "local-source-address" { return "LocalSourceAddress" }
    if yname == "originated-qos-marking" { return "OriginatedQosMarking" }
    return ""
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = state.SubscriptionId
    leafs["local-source-address"] = state.LocalSourceAddress
    leafs["originated-qos-marking"] = state.OriginatedQosMarking
    return leafs
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_State) GetParentYangName() string { return "subscription" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles
// A sensor profile is a set of sensor groups or
// individual sensor paths which are associated with a
// telemetry subscription. This is the source of the
// telemetry data for the subscription to send to the
// defined collectors.
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of telemetry sensor groups used in the subscription. The type is slice
    // of
    // TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile.
    SensorProfile []TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile
}

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetFilter() yfilter.YFilter { return sensorProfiles.YFilter }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) SetFilter(yf yfilter.YFilter) { sensorProfiles.YFilter = yf }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetGoName(yname string) string {
    if yname == "sensor-profile" { return "SensorProfile" }
    return ""
}

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetSegmentPath() string {
    return "sensor-profiles"
}

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-profile" {
        for _, c := range sensorProfiles.SensorProfile {
            if sensorProfiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile{}
        sensorProfiles.SensorProfile = append(sensorProfiles.SensorProfile, child)
        return &sensorProfiles.SensorProfile[len(sensorProfiles.SensorProfile)-1]
    }
    return nil
}

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorProfiles.SensorProfile {
        children[sensorProfiles.SensorProfile[i].GetSegmentPath()] = &sensorProfiles.SensorProfile[i]
    }
    return children
}

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetBundleName() string { return "openconfig" }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetYangName() string { return "sensor-profiles" }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) SetParent(parent types.Entity) { sensorProfiles.parent = parent }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetParent() types.Entity { return sensorProfiles.parent }

func (sensorProfiles *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles) GetParentYangName() string { return "subscription" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile
// List of telemetry sensor groups used
// in the subscription
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the telemetry sensor group name. The
    // type is string. Refers to
    // telemetry.TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config_SensorGroup
    SensorGroup interface{}

    // Configuration parameters related to the sensor profile for a subscription.
    Config TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config

    // State information relating to the sensor profile for a subscription.
    State TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State
}

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetFilter() yfilter.YFilter { return sensorProfile.YFilter }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) SetFilter(yf yfilter.YFilter) { sensorProfile.YFilter = yf }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetGoName(yname string) string {
    if yname == "sensor-group" { return "SensorGroup" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetSegmentPath() string {
    return "sensor-profile" + "[sensor-group='" + fmt.Sprintf("%v", sensorProfile.SensorGroup) + "']"
}

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &sensorProfile.Config
    }
    if childYangName == "state" {
        return &sensorProfile.State
    }
    return nil
}

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &sensorProfile.Config
    children["state"] = &sensorProfile.State
    return children
}

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group"] = sensorProfile.SensorGroup
    return leafs
}

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetBundleName() string { return "openconfig" }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetYangName() string { return "sensor-profile" }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) SetParent(parent types.Entity) { sensorProfile.parent = parent }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetParent() types.Entity { return sensorProfile.parent }

func (sensorProfile *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile) GetParentYangName() string { return "sensor-profiles" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config
// Configuration parameters related to the sensor
// profile for a subscription
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config struct {
    parent types.Entity
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

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetGoName(yname string) string {
    if yname == "sensor-group" { return "SensorGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "heartbeat-interval" { return "HeartbeatInterval" }
    if yname == "suppress-redundant" { return "SuppressRedundant" }
    return ""
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetSegmentPath() string {
    return "config"
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group"] = config.SensorGroup
    leafs["sample-interval"] = config.SampleInterval
    leafs["heartbeat-interval"] = config.HeartbeatInterval
    leafs["suppress-redundant"] = config.SuppressRedundant
    return leafs
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetBundleName() string { return "openconfig" }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetYangName() string { return "config" }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetParent() types.Entity { return config.parent }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_Config) GetParentYangName() string { return "sensor-profile" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State
// State information relating to the sensor profile
// for a subscription
type TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State struct {
    parent types.Entity
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

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetGoName(yname string) string {
    if yname == "sensor-group" { return "SensorGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "heartbeat-interval" { return "HeartbeatInterval" }
    if yname == "suppress-redundant" { return "SuppressRedundant" }
    return ""
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group"] = state.SensorGroup
    leafs["sample-interval"] = state.SampleInterval
    leafs["heartbeat-interval"] = state.HeartbeatInterval
    leafs["suppress-redundant"] = state.SuppressRedundant
    return leafs
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_SensorProfiles_SensorProfile_State) GetParentYangName() string { return "sensor-profile" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Identifier of the previously defined destination group. The type is slice
    // of
    // TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup.
    DestinationGroup []TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup
}

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetFilter() yfilter.YFilter { return destinationGroups.YFilter }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) SetFilter(yf yfilter.YFilter) { destinationGroups.YFilter = yf }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetGoName(yname string) string {
    if yname == "destination-group" { return "DestinationGroup" }
    return ""
}

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetSegmentPath() string {
    return "destination-groups"
}

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-group" {
        for _, c := range destinationGroups.DestinationGroup {
            if destinationGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup{}
        destinationGroups.DestinationGroup = append(destinationGroups.DestinationGroup, child)
        return &destinationGroups.DestinationGroup[len(destinationGroups.DestinationGroup)-1]
    }
    return nil
}

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destinationGroups.DestinationGroup {
        children[destinationGroups.DestinationGroup[i].GetSegmentPath()] = &destinationGroups.DestinationGroup[i]
    }
    return children
}

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetBundleName() string { return "openconfig" }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetYangName() string { return "destination-groups" }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) SetParent(parent types.Entity) { destinationGroups.parent = parent }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetParent() types.Entity { return destinationGroups.parent }

func (destinationGroups *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups) GetParentYangName() string { return "subscription" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup
// Identifier of the previously defined destination
// group
type TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetFilter() yfilter.YFilter { return destinationGroup.YFilter }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) SetFilter(yf yfilter.YFilter) { destinationGroup.YFilter = yf }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetSegmentPath() string {
    return "destination-group" + "[group-id='" + fmt.Sprintf("%v", destinationGroup.GroupId) + "']"
}

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &destinationGroup.Config
    }
    if childYangName == "state" {
        return &destinationGroup.State
    }
    return nil
}

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &destinationGroup.Config
    children["state"] = &destinationGroup.State
    return children
}

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = destinationGroup.GroupId
    return leafs
}

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetBundleName() string { return "openconfig" }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetYangName() string { return "destination-group" }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) SetParent(parent types.Entity) { destinationGroup.parent = parent }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetParent() types.Entity { return destinationGroup.parent }

func (destinationGroup *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup) GetParentYangName() string { return "destination-groups" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config
// Configuration parameters related to telemetry
// destinations.
type TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The destination group id references a reusable group of destination
    // addresses and ports for the telemetry stream. The type is string. Refers to
    // telemetry.TelemetrySystem_DestinationGroups_DestinationGroup_GroupId
    GroupId interface{}
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = config.GroupId
    return leafs
}

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetBundleName() string { return "openconfig" }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetYangName() string { return "config" }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetParent() types.Entity { return config.parent }

func (config *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_Config) GetParentYangName() string { return "destination-group" }

// TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State
// State information related to telemetry
// destinations
type TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The destination group id references a reusable group of destination
    // addresses and ports for the telemetry stream. The type is string. Refers to
    // telemetry.TelemetrySystem_DestinationGroups_DestinationGroup_GroupId
    GroupId interface{}
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = state.GroupId
    return leafs
}

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_Subscriptions_Persistent_Subscription_DestinationGroups_DestinationGroup_State) GetParentYangName() string { return "destination-group" }

// TelemetrySystem_Subscriptions_Dynamic
// This container holds information relating to dynamic
// telemetry subscriptions. A dynamic subscription is
// typically configured through an RPC channel, and does not
// persist across device restarts, or if the RPC channel is
// reset or otherwise torn down.
type TelemetrySystem_Subscriptions_Dynamic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List representation of telemetry subscriptions that are configured via an
    // inline RPC, otherwise known as dynamic telemetry subscriptions. The type is
    // slice of TelemetrySystem_Subscriptions_Dynamic_Subscription.
    Subscription []TelemetrySystem_Subscriptions_Dynamic_Subscription
}

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetFilter() yfilter.YFilter { return dynamic.YFilter }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) SetFilter(yf yfilter.YFilter) { dynamic.YFilter = yf }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetGoName(yname string) string {
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetSegmentPath() string {
    return "dynamic"
}

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscription" {
        for _, c := range dynamic.Subscription {
            if dynamic.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_Subscriptions_Dynamic_Subscription{}
        dynamic.Subscription = append(dynamic.Subscription, child)
        return &dynamic.Subscription[len(dynamic.Subscription)-1]
    }
    return nil
}

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dynamic.Subscription {
        children[dynamic.Subscription[i].GetSegmentPath()] = &dynamic.Subscription[i]
    }
    return children
}

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetBundleName() string { return "openconfig" }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetYangName() string { return "dynamic" }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) SetParent(parent types.Entity) { dynamic.parent = parent }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetParent() types.Entity { return dynamic.parent }

func (dynamic *TelemetrySystem_Subscriptions_Dynamic) GetParentYangName() string { return "subscriptions" }

// TelemetrySystem_Subscriptions_Dynamic_Subscription
// List representation of telemetry subscriptions that
// are configured via an inline RPC, otherwise known
// as dynamic telemetry subscriptions.
type TelemetrySystem_Subscriptions_Dynamic_Subscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "state" { return "State" }
    if yname == "sensor-paths" { return "SensorPaths" }
    return ""
}

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetSegmentPath() string {
    return "subscription" + "[subscription-id='" + fmt.Sprintf("%v", subscription.SubscriptionId) + "']"
}

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &subscription.State
    }
    if childYangName == "sensor-paths" {
        return &subscription.SensorPaths
    }
    return nil
}

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &subscription.State
    children["sensor-paths"] = &subscription.SensorPaths
    return children
}

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = subscription.SubscriptionId
    return leafs
}

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetBundleName() string { return "openconfig" }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetYangName() string { return "subscription" }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *TelemetrySystem_Subscriptions_Dynamic_Subscription) GetParentYangName() string { return "dynamic" }

// TelemetrySystem_Subscriptions_Dynamic_Subscription_State
// State information relating to dynamic telemetry
// subscriptions.
type TelemetrySystem_Subscriptions_Dynamic_Subscription_State struct {
    parent types.Entity
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

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "destination-protocol" { return "DestinationProtocol" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "heartbeat-interval" { return "HeartbeatInterval" }
    if yname == "suppress-redundant" { return "SuppressRedundant" }
    if yname == "originated-qos-marking" { return "OriginatedQosMarking" }
    return ""
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = state.SubscriptionId
    leafs["destination-address"] = state.DestinationAddress
    leafs["destination-port"] = state.DestinationPort
    leafs["destination-protocol"] = state.DestinationProtocol
    leafs["sample-interval"] = state.SampleInterval
    leafs["heartbeat-interval"] = state.HeartbeatInterval
    leafs["suppress-redundant"] = state.SuppressRedundant
    leafs["originated-qos-marking"] = state.OriginatedQosMarking
    return leafs
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_State) GetParentYangName() string { return "subscription" }

// TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths
// Top level container to hold a set of sensor
// paths grouped together
type TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of paths in the model which together comprise a sensor grouping.
    // Filters for each path to exclude items are also provided. The type is slice
    // of
    // TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath.
    SensorPath []TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath
}

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetFilter() yfilter.YFilter { return sensorPaths.YFilter }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) SetFilter(yf yfilter.YFilter) { sensorPaths.YFilter = yf }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetGoName(yname string) string {
    if yname == "sensor-path" { return "SensorPath" }
    return ""
}

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetSegmentPath() string {
    return "sensor-paths"
}

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-path" {
        for _, c := range sensorPaths.SensorPath {
            if sensorPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath{}
        sensorPaths.SensorPath = append(sensorPaths.SensorPath, child)
        return &sensorPaths.SensorPath[len(sensorPaths.SensorPath)-1]
    }
    return nil
}

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorPaths.SensorPath {
        children[sensorPaths.SensorPath[i].GetSegmentPath()] = &sensorPaths.SensorPath[i]
    }
    return children
}

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetBundleName() string { return "openconfig" }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetYangName() string { return "sensor-paths" }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) SetParent(parent types.Entity) { sensorPaths.parent = parent }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetParent() types.Entity { return sensorPaths.parent }

func (sensorPaths *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths) GetParentYangName() string { return "subscription" }

// TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath
// List of paths in the model which together
// comprise a sensor grouping. Filters for each path
// to exclude items are also provided.
type TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the path of interest. The type is
    // string. Refers to
    // telemetry.TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State_Path
    Path interface{}

    // State information for a dynamic subscription's paths of interest.
    State TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State
}

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetFilter() yfilter.YFilter { return sensorPath.YFilter }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) SetFilter(yf yfilter.YFilter) { sensorPath.YFilter = yf }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "state" { return "State" }
    return ""
}

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetSegmentPath() string {
    return "sensor-path" + "[path='" + fmt.Sprintf("%v", sensorPath.Path) + "']"
}

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &sensorPath.State
    }
    return nil
}

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &sensorPath.State
    return children
}

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = sensorPath.Path
    return leafs
}

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetBundleName() string { return "openconfig" }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetYangName() string { return "sensor-path" }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) SetParent(parent types.Entity) { sensorPath.parent = parent }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetParent() types.Entity { return sensorPath.parent }

func (sensorPath *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath) GetParentYangName() string { return "sensor-paths" }

// TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State
// State information for a dynamic subscription's
// paths of interest
type TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path to a section of operational state of interest (the sensor). The type
    // is string.
    Path interface{}

    // Filter to exclude certain values out of the state values. The type is
    // string.
    ExcludeFilter interface{}
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "exclude-filter" { return "ExcludeFilter" }
    return ""
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetSegmentPath() string {
    return "state"
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = state.Path
    leafs["exclude-filter"] = state.ExcludeFilter
    return leafs
}

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetBundleName() string { return "openconfig" }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetYangName() string { return "state" }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetParent() types.Entity { return state.parent }

func (state *TelemetrySystem_Subscriptions_Dynamic_Subscription_SensorPaths_SensorPath_State) GetParentYangName() string { return "sensor-path" }

