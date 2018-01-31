// This module contains a collection of YANG definitions
// for Cisco IOS-XR telemetry-model-driven package configuration.
// 
// This module contains definitions
// for the following management objects:
//   telemetry-model-driven: Model Driven Telemetry configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package telemetry_model_driven_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package telemetry_model_driven_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-telemetry-model-driven-cfg telemetry-model-driven}", reflect.TypeOf(TelemetryModelDriven{}))
    ydk.RegisterEntity("Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven", reflect.TypeOf(TelemetryModelDriven{}))
}

// ProtoType represents Proto type
type ProtoType string

const (
    // GRPC
    ProtoType_grpc ProtoType = "grpc"

    // tcp
    ProtoType_tcp ProtoType = "tcp"

    // udp
    ProtoType_udp ProtoType = "udp"
)

// EncodeType represents Encode type
type EncodeType string

const (
    // GPB
    EncodeType_gpb EncodeType = "gpb"

    // SELF DESCRIBING GPB
    EncodeType_self_describing_gpb EncodeType = "self-describing-gpb"

    // JSON
    EncodeType_json EncodeType = "json"
)

// MdtDscpValue represents Mdt dscp value
type MdtDscpValue string

const (
    // Applicable to DSCP: bits 000000
    MdtDscpValue_default_ MdtDscpValue = "default"

    // Applicable to DSCP: bits 001000
    MdtDscpValue_cs1 MdtDscpValue = "cs1"

    // Applicable to DSCP: bits 001010
    MdtDscpValue_af11 MdtDscpValue = "af11"

    // Applicable to DSCP: bits 001100
    MdtDscpValue_af12 MdtDscpValue = "af12"

    // Applicable to DSCP: bits 001110
    MdtDscpValue_af13 MdtDscpValue = "af13"

    // Applicable to DSCP: bits 010000
    MdtDscpValue_cs2 MdtDscpValue = "cs2"

    // Applicable to DSCP: bits 010010
    MdtDscpValue_af21 MdtDscpValue = "af21"

    // Applicable to DSCP: bits 010100
    MdtDscpValue_af22 MdtDscpValue = "af22"

    // Applicable to DSCP: bits 010110
    MdtDscpValue_af23 MdtDscpValue = "af23"

    // Applicable to DSCP: bits 011000
    MdtDscpValue_cs3 MdtDscpValue = "cs3"

    // Applicable to DSCP: bits 011010
    MdtDscpValue_af31 MdtDscpValue = "af31"

    // Applicable to DSCP: bits 011100
    MdtDscpValue_af32 MdtDscpValue = "af32"

    // Applicable to DSCP: bits 011110
    MdtDscpValue_af33 MdtDscpValue = "af33"

    // Applicable to DSCP: bits 100000
    MdtDscpValue_cs4 MdtDscpValue = "cs4"

    // Applicable to DSCP: bits 100010
    MdtDscpValue_af41 MdtDscpValue = "af41"

    // Applicable to DSCP: bits 100100
    MdtDscpValue_af42 MdtDscpValue = "af42"

    // Applicable to DSCP: bits 100110
    MdtDscpValue_af43 MdtDscpValue = "af43"

    // Applicable to DSCP: bits 101000
    MdtDscpValue_cs5 MdtDscpValue = "cs5"

    // Applicable to DSCP: bits 101110
    MdtDscpValue_ef MdtDscpValue = "ef"

    // Applicable to DSCP: bits 110000
    MdtDscpValue_cs6 MdtDscpValue = "cs6"

    // Applicable to DSCP: bits 111000
    MdtDscpValue_cs7 MdtDscpValue = "cs7"
)

// TelemetryModelDriven
// Model Driven Telemetry configuration
type TelemetryModelDriven struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Model Driven Telemetry. The type is interface{}.
    Enable interface{}

    // Sensor group configuration.
    SensorGroups TelemetryModelDriven_SensorGroups

    // Streaming Telemetry Subscription.
    Subscriptions TelemetryModelDriven_Subscriptions

    // Destination Group configuration.
    DestinationGroups TelemetryModelDriven_DestinationGroups
}

func (telemetryModelDriven *TelemetryModelDriven) GetFilter() yfilter.YFilter { return telemetryModelDriven.YFilter }

func (telemetryModelDriven *TelemetryModelDriven) SetFilter(yf yfilter.YFilter) { telemetryModelDriven.YFilter = yf }

func (telemetryModelDriven *TelemetryModelDriven) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "sensor-groups" { return "SensorGroups" }
    if yname == "subscriptions" { return "Subscriptions" }
    if yname == "destination-groups" { return "DestinationGroups" }
    return ""
}

func (telemetryModelDriven *TelemetryModelDriven) GetSegmentPath() string {
    return "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven"
}

func (telemetryModelDriven *TelemetryModelDriven) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-groups" {
        return &telemetryModelDriven.SensorGroups
    }
    if childYangName == "subscriptions" {
        return &telemetryModelDriven.Subscriptions
    }
    if childYangName == "destination-groups" {
        return &telemetryModelDriven.DestinationGroups
    }
    return nil
}

func (telemetryModelDriven *TelemetryModelDriven) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensor-groups"] = &telemetryModelDriven.SensorGroups
    children["subscriptions"] = &telemetryModelDriven.Subscriptions
    children["destination-groups"] = &telemetryModelDriven.DestinationGroups
    return children
}

func (telemetryModelDriven *TelemetryModelDriven) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = telemetryModelDriven.Enable
    return leafs
}

func (telemetryModelDriven *TelemetryModelDriven) GetBundleName() string { return "cisco_ios_xr" }

func (telemetryModelDriven *TelemetryModelDriven) GetYangName() string { return "telemetry-model-driven" }

func (telemetryModelDriven *TelemetryModelDriven) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (telemetryModelDriven *TelemetryModelDriven) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (telemetryModelDriven *TelemetryModelDriven) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (telemetryModelDriven *TelemetryModelDriven) SetParent(parent types.Entity) { telemetryModelDriven.parent = parent }

func (telemetryModelDriven *TelemetryModelDriven) GetParent() types.Entity { return telemetryModelDriven.parent }

func (telemetryModelDriven *TelemetryModelDriven) GetParentYangName() string { return "Cisco-IOS-XR-telemetry-model-driven-cfg" }

// TelemetryModelDriven_SensorGroups
// Sensor group configuration
type TelemetryModelDriven_SensorGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor group configuration. The type is slice of
    // TelemetryModelDriven_SensorGroups_SensorGroup.
    SensorGroup []TelemetryModelDriven_SensorGroups_SensorGroup
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetFilter() yfilter.YFilter { return sensorGroups.YFilter }

func (sensorGroups *TelemetryModelDriven_SensorGroups) SetFilter(yf yfilter.YFilter) { sensorGroups.YFilter = yf }

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetGoName(yname string) string {
    if yname == "sensor-group" { return "SensorGroup" }
    return ""
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetSegmentPath() string {
    return "sensor-groups"
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-group" {
        for _, c := range sensorGroups.SensorGroup {
            if sensorGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_SensorGroups_SensorGroup{}
        sensorGroups.SensorGroup = append(sensorGroups.SensorGroup, child)
        return &sensorGroups.SensorGroup[len(sensorGroups.SensorGroup)-1]
    }
    return nil
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorGroups.SensorGroup {
        children[sensorGroups.SensorGroup[i].GetSegmentPath()] = &sensorGroups.SensorGroup[i]
    }
    return children
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetBundleName() string { return "cisco_ios_xr" }

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetYangName() string { return "sensor-groups" }

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorGroups *TelemetryModelDriven_SensorGroups) SetParent(parent types.Entity) { sensorGroups.parent = parent }

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetParent() types.Entity { return sensorGroups.parent }

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetParentYangName() string { return "telemetry-model-driven" }

// TelemetryModelDriven_SensorGroups_SensorGroup
// Sensor group configuration
type TelemetryModelDriven_SensorGroups_SensorGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this group. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    SensorGroupIdentifier interface{}

    // Sensor path configuration.
    SensorPaths TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetFilter() yfilter.YFilter { return sensorGroup.YFilter }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) SetFilter(yf yfilter.YFilter) { sensorGroup.YFilter = yf }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetGoName(yname string) string {
    if yname == "sensor-group-identifier" { return "SensorGroupIdentifier" }
    if yname == "sensor-paths" { return "SensorPaths" }
    return ""
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetSegmentPath() string {
    return "sensor-group" + "[sensor-group-identifier='" + fmt.Sprintf("%v", sensorGroup.SensorGroupIdentifier) + "']"
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-paths" {
        return &sensorGroup.SensorPaths
    }
    return nil
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensor-paths"] = &sensorGroup.SensorPaths
    return children
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group-identifier"] = sensorGroup.SensorGroupIdentifier
    return leafs
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetBundleName() string { return "cisco_ios_xr" }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetYangName() string { return "sensor-group" }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) SetParent(parent types.Entity) { sensorGroup.parent = parent }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetParent() types.Entity { return sensorGroup.parent }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetParentYangName() string { return "sensor-groups" }

// TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths
// Sensor path configuration
type TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor path configuration. The type is slice of
    // TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath.
    SensorPath []TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath
}

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetFilter() yfilter.YFilter { return sensorPaths.YFilter }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) SetFilter(yf yfilter.YFilter) { sensorPaths.YFilter = yf }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetGoName(yname string) string {
    if yname == "sensor-path" { return "SensorPath" }
    return ""
}

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetSegmentPath() string {
    return "sensor-paths"
}

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-path" {
        for _, c := range sensorPaths.SensorPath {
            if sensorPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath{}
        sensorPaths.SensorPath = append(sensorPaths.SensorPath, child)
        return &sensorPaths.SensorPath[len(sensorPaths.SensorPath)-1]
    }
    return nil
}

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorPaths.SensorPath {
        children[sensorPaths.SensorPath[i].GetSegmentPath()] = &sensorPaths.SensorPath[i]
    }
    return children
}

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetBundleName() string { return "cisco_ios_xr" }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetYangName() string { return "sensor-paths" }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) SetParent(parent types.Entity) { sensorPaths.parent = parent }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetParent() types.Entity { return sensorPaths.parent }

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetParentYangName() string { return "sensor-group" }

// TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath
// Sensor path configuration
type TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor Path. The type is string.
    TelemetrySensorPath interface{}
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetFilter() yfilter.YFilter { return sensorPath.YFilter }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) SetFilter(yf yfilter.YFilter) { sensorPath.YFilter = yf }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetGoName(yname string) string {
    if yname == "telemetry-sensor-path" { return "TelemetrySensorPath" }
    return ""
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetSegmentPath() string {
    return "sensor-path" + "[telemetry-sensor-path='" + fmt.Sprintf("%v", sensorPath.TelemetrySensorPath) + "']"
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["telemetry-sensor-path"] = sensorPath.TelemetrySensorPath
    return leafs
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetBundleName() string { return "cisco_ios_xr" }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetYangName() string { return "sensor-path" }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) SetParent(parent types.Entity) { sensorPath.parent = parent }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetParent() types.Entity { return sensorPath.parent }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetParentYangName() string { return "sensor-paths" }

// TelemetryModelDriven_Subscriptions
// Streaming Telemetry Subscription
type TelemetryModelDriven_Subscriptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Streaming Telemetry Subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription.
    Subscription []TelemetryModelDriven_Subscriptions_Subscription
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetFilter() yfilter.YFilter { return subscriptions.YFilter }

func (subscriptions *TelemetryModelDriven_Subscriptions) SetFilter(yf yfilter.YFilter) { subscriptions.YFilter = yf }

func (subscriptions *TelemetryModelDriven_Subscriptions) GetGoName(yname string) string {
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetSegmentPath() string {
    return "subscriptions"
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscription" {
        for _, c := range subscriptions.Subscription {
            if subscriptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription{}
        subscriptions.Subscription = append(subscriptions.Subscription, child)
        return &subscriptions.Subscription[len(subscriptions.Subscription)-1]
    }
    return nil
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subscriptions.Subscription {
        children[subscriptions.Subscription[i].GetSegmentPath()] = &subscriptions.Subscription[i]
    }
    return children
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetBundleName() string { return "cisco_ios_xr" }

func (subscriptions *TelemetryModelDriven_Subscriptions) GetYangName() string { return "subscriptions" }

func (subscriptions *TelemetryModelDriven_Subscriptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriptions *TelemetryModelDriven_Subscriptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriptions *TelemetryModelDriven_Subscriptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriptions *TelemetryModelDriven_Subscriptions) SetParent(parent types.Entity) { subscriptions.parent = parent }

func (subscriptions *TelemetryModelDriven_Subscriptions) GetParent() types.Entity { return subscriptions.parent }

func (subscriptions *TelemetryModelDriven_Subscriptions) GetParentYangName() string { return "telemetry-model-driven" }

// TelemetryModelDriven_Subscriptions_Subscription
// Streaming Telemetry Subscription
type TelemetryModelDriven_Subscriptions_Subscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Subscription identifier string. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    SubscriptionIdentifier interface{}

    // Outgoing DSCP value. The type is MdtDscpValue.
    SourceQosMarking interface{}

    // Source address to use for streaming telemetry information. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Associate Sensor Groups with Subscription.
    SensorProfiles TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles

    // Associate Destination Groups with Subscription.
    DestinationProfiles TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetGoName(yname string) string {
    if yname == "subscription-identifier" { return "SubscriptionIdentifier" }
    if yname == "source-qos-marking" { return "SourceQosMarking" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "sensor-profiles" { return "SensorProfiles" }
    if yname == "destination-profiles" { return "DestinationProfiles" }
    return ""
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetSegmentPath() string {
    return "subscription" + "[subscription-identifier='" + fmt.Sprintf("%v", subscription.SubscriptionIdentifier) + "']"
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-profiles" {
        return &subscription.SensorProfiles
    }
    if childYangName == "destination-profiles" {
        return &subscription.DestinationProfiles
    }
    return nil
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensor-profiles"] = &subscription.SensorProfiles
    children["destination-profiles"] = &subscription.DestinationProfiles
    return children
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-identifier"] = subscription.SubscriptionIdentifier
    leafs["source-qos-marking"] = subscription.SourceQosMarking
    leafs["source-interface"] = subscription.SourceInterface
    return leafs
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetBundleName() string { return "cisco_ios_xr" }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetYangName() string { return "subscription" }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetParentYangName() string { return "subscriptions" }

// TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles
// Associate Sensor Groups with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Associate Sensor Group with Subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile.
    SensorProfile []TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile
}

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetFilter() yfilter.YFilter { return sensorProfiles.YFilter }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) SetFilter(yf yfilter.YFilter) { sensorProfiles.YFilter = yf }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetGoName(yname string) string {
    if yname == "sensor-profile" { return "SensorProfile" }
    return ""
}

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetSegmentPath() string {
    return "sensor-profiles"
}

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-profile" {
        for _, c := range sensorProfiles.SensorProfile {
            if sensorProfiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile{}
        sensorProfiles.SensorProfile = append(sensorProfiles.SensorProfile, child)
        return &sensorProfiles.SensorProfile[len(sensorProfiles.SensorProfile)-1]
    }
    return nil
}

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorProfiles.SensorProfile {
        children[sensorProfiles.SensorProfile[i].GetSegmentPath()] = &sensorProfiles.SensorProfile[i]
    }
    return children
}

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetBundleName() string { return "cisco_ios_xr" }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetYangName() string { return "sensor-profiles" }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) SetParent(parent types.Entity) { sensorProfiles.parent = parent }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetParent() types.Entity { return sensorProfiles.parent }

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetParentYangName() string { return "subscription" }

// TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile
// Associate Sensor Group with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the telemetry sensor group name. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Sensorgroupid interface{}

    // use strict timer. The type is interface{}.
    StrictTimer interface{}

    // Sample interval in milliseconds. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    SampleInterval interface{}
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetFilter() yfilter.YFilter { return sensorProfile.YFilter }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) SetFilter(yf yfilter.YFilter) { sensorProfile.YFilter = yf }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetGoName(yname string) string {
    if yname == "sensorgroupid" { return "Sensorgroupid" }
    if yname == "strict-timer" { return "StrictTimer" }
    if yname == "sample-interval" { return "SampleInterval" }
    return ""
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetSegmentPath() string {
    return "sensor-profile" + "[sensorgroupid='" + fmt.Sprintf("%v", sensorProfile.Sensorgroupid) + "']"
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensorgroupid"] = sensorProfile.Sensorgroupid
    leafs["strict-timer"] = sensorProfile.StrictTimer
    leafs["sample-interval"] = sensorProfile.SampleInterval
    return leafs
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetBundleName() string { return "cisco_ios_xr" }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetYangName() string { return "sensor-profile" }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) SetParent(parent types.Entity) { sensorProfile.parent = parent }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetParent() types.Entity { return sensorProfile.parent }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetParentYangName() string { return "sensor-profiles" }

// TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles
// Associate Destination Groups with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Associate Destination Group with Subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile.
    DestinationProfile []TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile
}

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetFilter() yfilter.YFilter { return destinationProfiles.YFilter }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) SetFilter(yf yfilter.YFilter) { destinationProfiles.YFilter = yf }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetGoName(yname string) string {
    if yname == "destination-profile" { return "DestinationProfile" }
    return ""
}

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetSegmentPath() string {
    return "destination-profiles"
}

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-profile" {
        for _, c := range destinationProfiles.DestinationProfile {
            if destinationProfiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile{}
        destinationProfiles.DestinationProfile = append(destinationProfiles.DestinationProfile, child)
        return &destinationProfiles.DestinationProfile[len(destinationProfiles.DestinationProfile)-1]
    }
    return nil
}

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destinationProfiles.DestinationProfile {
        children[destinationProfiles.DestinationProfile[i].GetSegmentPath()] = &destinationProfiles.DestinationProfile[i]
    }
    return children
}

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetBundleName() string { return "cisco_ios_xr" }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetYangName() string { return "destination-profiles" }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) SetParent(parent types.Entity) { destinationProfiles.parent = parent }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetParent() types.Entity { return destinationProfiles.parent }

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetParentYangName() string { return "subscription" }

// TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile
// Associate Destination Group with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Destination Id to associate with Subscription. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    DestinationId interface{}
}

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetFilter() yfilter.YFilter { return destinationProfile.YFilter }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) SetFilter(yf yfilter.YFilter) { destinationProfile.YFilter = yf }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetGoName(yname string) string {
    if yname == "destination-id" { return "DestinationId" }
    return ""
}

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetSegmentPath() string {
    return "destination-profile" + "[destination-id='" + fmt.Sprintf("%v", destinationProfile.DestinationId) + "']"
}

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-id"] = destinationProfile.DestinationId
    return leafs
}

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetBundleName() string { return "cisco_ios_xr" }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetYangName() string { return "destination-profile" }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) SetParent(parent types.Entity) { destinationProfile.parent = parent }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetParent() types.Entity { return destinationProfile.parent }

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetParentYangName() string { return "destination-profiles" }

// TelemetryModelDriven_DestinationGroups
// Destination Group configuration
type TelemetryModelDriven_DestinationGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination Group. The type is slice of
    // TelemetryModelDriven_DestinationGroups_DestinationGroup.
    DestinationGroup []TelemetryModelDriven_DestinationGroups_DestinationGroup
}

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetFilter() yfilter.YFilter { return destinationGroups.YFilter }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) SetFilter(yf yfilter.YFilter) { destinationGroups.YFilter = yf }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetGoName(yname string) string {
    if yname == "destination-group" { return "DestinationGroup" }
    return ""
}

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetSegmentPath() string {
    return "destination-groups"
}

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-group" {
        for _, c := range destinationGroups.DestinationGroup {
            if destinationGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_DestinationGroups_DestinationGroup{}
        destinationGroups.DestinationGroup = append(destinationGroups.DestinationGroup, child)
        return &destinationGroups.DestinationGroup[len(destinationGroups.DestinationGroup)-1]
    }
    return nil
}

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destinationGroups.DestinationGroup {
        children[destinationGroups.DestinationGroup[i].GetSegmentPath()] = &destinationGroups.DestinationGroup[i]
    }
    return children
}

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetBundleName() string { return "cisco_ios_xr" }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetYangName() string { return "destination-groups" }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) SetParent(parent types.Entity) { destinationGroups.parent = parent }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetParent() types.Entity { return destinationGroups.parent }

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetParentYangName() string { return "telemetry-model-driven" }

// TelemetryModelDriven_DestinationGroups_DestinationGroup
// Destination Group
type TelemetryModelDriven_DestinationGroups_DestinationGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. destination group id string. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    DestinationId interface{}

    // Vrf for the destination group. The type is string with length: 1..32.
    Vrf interface{}

    // Destination address configuration.
    Ipv6Destinations TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations

    // Destination address configuration.
    Ipv4Destinations TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations
}

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetFilter() yfilter.YFilter { return destinationGroup.YFilter }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) SetFilter(yf yfilter.YFilter) { destinationGroup.YFilter = yf }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetGoName(yname string) string {
    if yname == "destination-id" { return "DestinationId" }
    if yname == "vrf" { return "Vrf" }
    if yname == "ipv6-destinations" { return "Ipv6Destinations" }
    if yname == "ipv4-destinations" { return "Ipv4Destinations" }
    return ""
}

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetSegmentPath() string {
    return "destination-group" + "[destination-id='" + fmt.Sprintf("%v", destinationGroup.DestinationId) + "']"
}

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-destinations" {
        return &destinationGroup.Ipv6Destinations
    }
    if childYangName == "ipv4-destinations" {
        return &destinationGroup.Ipv4Destinations
    }
    return nil
}

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6-destinations"] = &destinationGroup.Ipv6Destinations
    children["ipv4-destinations"] = &destinationGroup.Ipv4Destinations
    return children
}

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-id"] = destinationGroup.DestinationId
    leafs["vrf"] = destinationGroup.Vrf
    return leafs
}

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetBundleName() string { return "cisco_ios_xr" }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetYangName() string { return "destination-group" }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) SetParent(parent types.Entity) { destinationGroup.parent = parent }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetParent() types.Entity { return destinationGroup.parent }

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetParentYangName() string { return "destination-groups" }

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations
// Destination address configuration
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // destination IP address. The type is slice of
    // TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination.
    Ipv6Destination []TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination
}

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetFilter() yfilter.YFilter { return ipv6Destinations.YFilter }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) SetFilter(yf yfilter.YFilter) { ipv6Destinations.YFilter = yf }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetGoName(yname string) string {
    if yname == "ipv6-destination" { return "Ipv6Destination" }
    return ""
}

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetSegmentPath() string {
    return "ipv6-destinations"
}

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-destination" {
        for _, c := range ipv6Destinations.Ipv6Destination {
            if ipv6Destinations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination{}
        ipv6Destinations.Ipv6Destination = append(ipv6Destinations.Ipv6Destination, child)
        return &ipv6Destinations.Ipv6Destination[len(ipv6Destinations.Ipv6Destination)-1]
    }
    return nil
}

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6Destinations.Ipv6Destination {
        children[ipv6Destinations.Ipv6Destination[i].GetSegmentPath()] = &ipv6Destinations.Ipv6Destination[i]
    }
    return children
}

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetYangName() string { return "ipv6-destinations" }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) SetParent(parent types.Entity) { ipv6Destinations.parent = parent }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetParent() types.Entity { return ipv6Destinations.parent }

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetParentYangName() string { return "destination-group" }

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination
// destination IP address
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Destination IPv6 address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // This attribute is a key. destination port. The type is interface{} with
    // range: 1..65535.
    DestinationPort interface{}

    // Encoding used to transmit telemetry data to the collector. The type is
    // EncodeType.
    Encoding interface{}

    // Transport Protocol used to transmit telemetry data to the collector.
    Protocol TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol
}

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetFilter() yfilter.YFilter { return ipv6Destination.YFilter }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) SetFilter(yf yfilter.YFilter) { ipv6Destination.YFilter = yf }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "encoding" { return "Encoding" }
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetSegmentPath() string {
    return "ipv6-destination" + "[ipv6-address='" + fmt.Sprintf("%v", ipv6Destination.Ipv6Address) + "']" + "[destination-port='" + fmt.Sprintf("%v", ipv6Destination.DestinationPort) + "']"
}

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        return &ipv6Destination.Protocol
    }
    return nil
}

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocol"] = &ipv6Destination.Protocol
    return children
}

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Destination.Ipv6Address
    leafs["destination-port"] = ipv6Destination.DestinationPort
    leafs["encoding"] = ipv6Destination.Encoding
    return leafs
}

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetYangName() string { return "ipv6-destination" }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) SetParent(parent types.Entity) { ipv6Destination.parent = parent }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetParent() types.Entity { return ipv6Destination.parent }

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetParentYangName() string { return "ipv6-destinations" }

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol
// Transport Protocol used to transmit telemetry
// data to the collector
// This type is a presence type.
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // protocol. The type is ProtoType. This attribute is mandatory.
    Protocol interface{}

    // tls hostname. The type is string.
    TlsHostname interface{}

    // no tls. The type is interface{} with range: -2147483648..2147483647. The
    // default value is 0.
    NoTls interface{}

    // udp packetsize. The type is interface{} with range: 484..65507. The default
    // value is 1472.
    Packetsize interface{}
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "tls-hostname" { return "TlsHostname" }
    if yname == "no-tls" { return "NoTls" }
    if yname == "packetsize" { return "Packetsize" }
    return ""
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetSegmentPath() string {
    return "protocol"
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = protocol.Protocol
    leafs["tls-hostname"] = protocol.TlsHostname
    leafs["no-tls"] = protocol.NoTls
    leafs["packetsize"] = protocol.Packetsize
    return leafs
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetYangName() string { return "protocol" }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetParentYangName() string { return "ipv6-destination" }

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations
// Destination address configuration
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // destination IP address. The type is slice of
    // TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination.
    Ipv4Destination []TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination
}

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetFilter() yfilter.YFilter { return ipv4Destinations.YFilter }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) SetFilter(yf yfilter.YFilter) { ipv4Destinations.YFilter = yf }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetGoName(yname string) string {
    if yname == "ipv4-destination" { return "Ipv4Destination" }
    return ""
}

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetSegmentPath() string {
    return "ipv4-destinations"
}

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-destination" {
        for _, c := range ipv4Destinations.Ipv4Destination {
            if ipv4Destinations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination{}
        ipv4Destinations.Ipv4Destination = append(ipv4Destinations.Ipv4Destination, child)
        return &ipv4Destinations.Ipv4Destination[len(ipv4Destinations.Ipv4Destination)-1]
    }
    return nil
}

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Destinations.Ipv4Destination {
        children[ipv4Destinations.Ipv4Destination[i].GetSegmentPath()] = &ipv4Destinations.Ipv4Destination[i]
    }
    return children
}

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetYangName() string { return "ipv4-destinations" }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) SetParent(parent types.Entity) { ipv4Destinations.parent = parent }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetParent() types.Entity { return ipv4Destinations.parent }

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetParentYangName() string { return "destination-group" }

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination
// destination IP address
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Destination IPv4 address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // This attribute is a key. destination port. The type is interface{} with
    // range: 1..65535.
    DestinationPort interface{}

    // Encoding used to transmit telemetry data to the collector. The type is
    // EncodeType.
    Encoding interface{}

    // Transport Protocol used to transmit telemetry data to the collector.
    Protocol TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol
}

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetFilter() yfilter.YFilter { return ipv4Destination.YFilter }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) SetFilter(yf yfilter.YFilter) { ipv4Destination.YFilter = yf }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "encoding" { return "Encoding" }
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetSegmentPath() string {
    return "ipv4-destination" + "[ipv4-address='" + fmt.Sprintf("%v", ipv4Destination.Ipv4Address) + "']" + "[destination-port='" + fmt.Sprintf("%v", ipv4Destination.DestinationPort) + "']"
}

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        return &ipv4Destination.Protocol
    }
    return nil
}

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocol"] = &ipv4Destination.Protocol
    return children
}

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = ipv4Destination.Ipv4Address
    leafs["destination-port"] = ipv4Destination.DestinationPort
    leafs["encoding"] = ipv4Destination.Encoding
    return leafs
}

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetYangName() string { return "ipv4-destination" }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) SetParent(parent types.Entity) { ipv4Destination.parent = parent }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetParent() types.Entity { return ipv4Destination.parent }

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetParentYangName() string { return "ipv4-destinations" }

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol
// Transport Protocol used to transmit telemetry
// data to the collector
// This type is a presence type.
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // protocol. The type is ProtoType. This attribute is mandatory.
    Protocol interface{}

    // tls hostname. The type is string.
    TlsHostname interface{}

    // no tls. The type is interface{} with range: -2147483648..2147483647. The
    // default value is 0.
    NoTls interface{}

    // udp packetsize. The type is interface{} with range: 484..65507. The default
    // value is 1472.
    Packetsize interface{}
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "tls-hostname" { return "TlsHostname" }
    if yname == "no-tls" { return "NoTls" }
    if yname == "packetsize" { return "Packetsize" }
    return ""
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetSegmentPath() string {
    return "protocol"
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = protocol.Protocol
    leafs["tls-hostname"] = protocol.TlsHostname
    leafs["no-tls"] = protocol.NoTls
    leafs["packetsize"] = protocol.Packetsize
    return leafs
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetYangName() string { return "protocol" }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetParentYangName() string { return "ipv4-destination" }

