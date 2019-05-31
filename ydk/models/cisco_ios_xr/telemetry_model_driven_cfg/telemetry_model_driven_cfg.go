// This module contains a collection of YANG definitions
// for Cisco IOS-XR telemetry-model-driven package configuration.
// 
// This module contains definitions
// for the following management objects:
//   telemetry-model-driven: Model Driven Telemetry configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enable strict-timer for all subscriptions, default is relative timer. The
    // type is interface{}.
    StrictTimer interface{}

    // Enable Model Driven Telemetry. The type is interface{}.
    Enable interface{}

    // Maximum allowed sensor paths, default: 1000. The type is interface{} with
    // range: 0..4000.
    MaxSensorPaths interface{}

    // Maximum containers allowed per path, 0 disables the check. The type is
    // interface{} with range: 0..1024.
    MaxContainersPerPath interface{}

    // TCP send timeout value, default:30 sec,0 will disable the timeout. The type
    // is interface{} with range: 0..30.
    TcpSendTimeout interface{}

    // Sensor group configuration.
    SensorGroups TelemetryModelDriven_SensorGroups

    // Streaming Telemetry Subscription.
    Subscriptions TelemetryModelDriven_Subscriptions

    // Include fields with empty values in output.
    Include TelemetryModelDriven_Include

    // Destination Group configuration.
    DestinationGroups TelemetryModelDriven_DestinationGroups
}

func (telemetryModelDriven *TelemetryModelDriven) GetEntityData() *types.CommonEntityData {
    telemetryModelDriven.EntityData.YFilter = telemetryModelDriven.YFilter
    telemetryModelDriven.EntityData.YangName = "telemetry-model-driven"
    telemetryModelDriven.EntityData.BundleName = "cisco_ios_xr"
    telemetryModelDriven.EntityData.ParentYangName = "Cisco-IOS-XR-telemetry-model-driven-cfg"
    telemetryModelDriven.EntityData.SegmentPath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven"
    telemetryModelDriven.EntityData.AbsolutePath = telemetryModelDriven.EntityData.SegmentPath
    telemetryModelDriven.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telemetryModelDriven.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telemetryModelDriven.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telemetryModelDriven.EntityData.Children = types.NewOrderedMap()
    telemetryModelDriven.EntityData.Children.Append("sensor-groups", types.YChild{"SensorGroups", &telemetryModelDriven.SensorGroups})
    telemetryModelDriven.EntityData.Children.Append("subscriptions", types.YChild{"Subscriptions", &telemetryModelDriven.Subscriptions})
    telemetryModelDriven.EntityData.Children.Append("include", types.YChild{"Include", &telemetryModelDriven.Include})
    telemetryModelDriven.EntityData.Children.Append("destination-groups", types.YChild{"DestinationGroups", &telemetryModelDriven.DestinationGroups})
    telemetryModelDriven.EntityData.Leafs = types.NewOrderedMap()
    telemetryModelDriven.EntityData.Leafs.Append("strict-timer", types.YLeaf{"StrictTimer", telemetryModelDriven.StrictTimer})
    telemetryModelDriven.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", telemetryModelDriven.Enable})
    telemetryModelDriven.EntityData.Leafs.Append("max-sensor-paths", types.YLeaf{"MaxSensorPaths", telemetryModelDriven.MaxSensorPaths})
    telemetryModelDriven.EntityData.Leafs.Append("max-containers-per-path", types.YLeaf{"MaxContainersPerPath", telemetryModelDriven.MaxContainersPerPath})
    telemetryModelDriven.EntityData.Leafs.Append("tcp-send-timeout", types.YLeaf{"TcpSendTimeout", telemetryModelDriven.TcpSendTimeout})

    telemetryModelDriven.EntityData.YListKeys = []string {}

    return &(telemetryModelDriven.EntityData)
}

// TelemetryModelDriven_SensorGroups
// Sensor group configuration
type TelemetryModelDriven_SensorGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor group configuration. The type is slice of
    // TelemetryModelDriven_SensorGroups_SensorGroup.
    SensorGroup []*TelemetryModelDriven_SensorGroups_SensorGroup
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetEntityData() *types.CommonEntityData {
    sensorGroups.EntityData.YFilter = sensorGroups.YFilter
    sensorGroups.EntityData.YangName = "sensor-groups"
    sensorGroups.EntityData.BundleName = "cisco_ios_xr"
    sensorGroups.EntityData.ParentYangName = "telemetry-model-driven"
    sensorGroups.EntityData.SegmentPath = "sensor-groups"
    sensorGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/" + sensorGroups.EntityData.SegmentPath
    sensorGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorGroups.EntityData.Children = types.NewOrderedMap()
    sensorGroups.EntityData.Children.Append("sensor-group", types.YChild{"SensorGroup", nil})
    for i := range sensorGroups.SensorGroup {
        sensorGroups.EntityData.Children.Append(types.GetSegmentPath(sensorGroups.SensorGroup[i]), types.YChild{"SensorGroup", sensorGroups.SensorGroup[i]})
    }
    sensorGroups.EntityData.Leafs = types.NewOrderedMap()

    sensorGroups.EntityData.YListKeys = []string {}

    return &(sensorGroups.EntityData)
}

// TelemetryModelDriven_SensorGroups_SensorGroup
// Sensor group configuration
type TelemetryModelDriven_SensorGroups_SensorGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The identifier for this group. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SensorGroupIdentifier interface{}

    // Sensor path configuration.
    SensorPaths TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetEntityData() *types.CommonEntityData {
    sensorGroup.EntityData.YFilter = sensorGroup.YFilter
    sensorGroup.EntityData.YangName = "sensor-group"
    sensorGroup.EntityData.BundleName = "cisco_ios_xr"
    sensorGroup.EntityData.ParentYangName = "sensor-groups"
    sensorGroup.EntityData.SegmentPath = "sensor-group" + types.AddKeyToken(sensorGroup.SensorGroupIdentifier, "sensor-group-identifier")
    sensorGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/sensor-groups/" + sensorGroup.EntityData.SegmentPath
    sensorGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorGroup.EntityData.Children = types.NewOrderedMap()
    sensorGroup.EntityData.Children.Append("sensor-paths", types.YChild{"SensorPaths", &sensorGroup.SensorPaths})
    sensorGroup.EntityData.Leafs = types.NewOrderedMap()
    sensorGroup.EntityData.Leafs.Append("sensor-group-identifier", types.YLeaf{"SensorGroupIdentifier", sensorGroup.SensorGroupIdentifier})

    sensorGroup.EntityData.YListKeys = []string {"SensorGroupIdentifier"}

    return &(sensorGroup.EntityData)
}

// TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths
// Sensor path configuration
type TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor path configuration. The type is slice of
    // TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath.
    SensorPath []*TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath
}

func (sensorPaths *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths) GetEntityData() *types.CommonEntityData {
    sensorPaths.EntityData.YFilter = sensorPaths.YFilter
    sensorPaths.EntityData.YangName = "sensor-paths"
    sensorPaths.EntityData.BundleName = "cisco_ios_xr"
    sensorPaths.EntityData.ParentYangName = "sensor-group"
    sensorPaths.EntityData.SegmentPath = "sensor-paths"
    sensorPaths.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/sensor-groups/sensor-group/" + sensorPaths.EntityData.SegmentPath
    sensorPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorPaths.EntityData.Children = types.NewOrderedMap()
    sensorPaths.EntityData.Children.Append("sensor-path", types.YChild{"SensorPath", nil})
    for i := range sensorPaths.SensorPath {
        sensorPaths.EntityData.Children.Append(types.GetSegmentPath(sensorPaths.SensorPath[i]), types.YChild{"SensorPath", sensorPaths.SensorPath[i]})
    }
    sensorPaths.EntityData.Leafs = types.NewOrderedMap()

    sensorPaths.EntityData.YListKeys = []string {}

    return &(sensorPaths.EntityData)
}

// TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath
// Sensor path configuration
type TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor Path. The type is string.
    TelemetrySensorPath interface{}
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPaths_SensorPath) GetEntityData() *types.CommonEntityData {
    sensorPath.EntityData.YFilter = sensorPath.YFilter
    sensorPath.EntityData.YangName = "sensor-path"
    sensorPath.EntityData.BundleName = "cisco_ios_xr"
    sensorPath.EntityData.ParentYangName = "sensor-paths"
    sensorPath.EntityData.SegmentPath = "sensor-path" + types.AddKeyToken(sensorPath.TelemetrySensorPath, "telemetry-sensor-path")
    sensorPath.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/sensor-groups/sensor-group/sensor-paths/" + sensorPath.EntityData.SegmentPath
    sensorPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorPath.EntityData.Children = types.NewOrderedMap()
    sensorPath.EntityData.Leafs = types.NewOrderedMap()
    sensorPath.EntityData.Leafs.Append("telemetry-sensor-path", types.YLeaf{"TelemetrySensorPath", sensorPath.TelemetrySensorPath})

    sensorPath.EntityData.YListKeys = []string {"TelemetrySensorPath"}

    return &(sensorPath.EntityData)
}

// TelemetryModelDriven_Subscriptions
// Streaming Telemetry Subscription
type TelemetryModelDriven_Subscriptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Streaming Telemetry Subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription.
    Subscription []*TelemetryModelDriven_Subscriptions_Subscription
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetEntityData() *types.CommonEntityData {
    subscriptions.EntityData.YFilter = subscriptions.YFilter
    subscriptions.EntityData.YangName = "subscriptions"
    subscriptions.EntityData.BundleName = "cisco_ios_xr"
    subscriptions.EntityData.ParentYangName = "telemetry-model-driven"
    subscriptions.EntityData.SegmentPath = "subscriptions"
    subscriptions.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/" + subscriptions.EntityData.SegmentPath
    subscriptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriptions.EntityData.Children = types.NewOrderedMap()
    subscriptions.EntityData.Children.Append("subscription", types.YChild{"Subscription", nil})
    for i := range subscriptions.Subscription {
        subscriptions.EntityData.Children.Append(types.GetSegmentPath(subscriptions.Subscription[i]), types.YChild{"Subscription", subscriptions.Subscription[i]})
    }
    subscriptions.EntityData.Leafs = types.NewOrderedMap()

    subscriptions.EntityData.YListKeys = []string {}

    return &(subscriptions.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription
// Streaming Telemetry Subscription
type TelemetryModelDriven_Subscriptions_Subscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscription identifier string. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SubscriptionIdentifier interface{}

    // Outgoing DSCP value. The type is MdtDscpValue.
    SourceQosMarking interface{}

    // Source address to use for streaming telemetry information. The type is
    // string with pattern: b'[a-zA-Z0-9._/-]+'.
    SourceInterface interface{}

    // Associate Sensor Groups with Subscription.
    SensorProfiles TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles

    // Associate Destination Groups with Subscription.
    DestinationProfiles TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "cisco_ios_xr"
    subscription.EntityData.ParentYangName = "subscriptions"
    subscription.EntityData.SegmentPath = "subscription" + types.AddKeyToken(subscription.SubscriptionIdentifier, "subscription-identifier")
    subscription.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/subscriptions/" + subscription.EntityData.SegmentPath
    subscription.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscription.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscription.EntityData.Children = types.NewOrderedMap()
    subscription.EntityData.Children.Append("sensor-profiles", types.YChild{"SensorProfiles", &subscription.SensorProfiles})
    subscription.EntityData.Children.Append("destination-profiles", types.YChild{"DestinationProfiles", &subscription.DestinationProfiles})
    subscription.EntityData.Leafs = types.NewOrderedMap()
    subscription.EntityData.Leafs.Append("subscription-identifier", types.YLeaf{"SubscriptionIdentifier", subscription.SubscriptionIdentifier})
    subscription.EntityData.Leafs.Append("source-qos-marking", types.YLeaf{"SourceQosMarking", subscription.SourceQosMarking})
    subscription.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", subscription.SourceInterface})

    subscription.EntityData.YListKeys = []string {"SubscriptionIdentifier"}

    return &(subscription.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles
// Associate Sensor Groups with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Associate Sensor Group with Subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile.
    SensorProfile []*TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile
}

func (sensorProfiles *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles) GetEntityData() *types.CommonEntityData {
    sensorProfiles.EntityData.YFilter = sensorProfiles.YFilter
    sensorProfiles.EntityData.YangName = "sensor-profiles"
    sensorProfiles.EntityData.BundleName = "cisco_ios_xr"
    sensorProfiles.EntityData.ParentYangName = "subscription"
    sensorProfiles.EntityData.SegmentPath = "sensor-profiles"
    sensorProfiles.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/subscriptions/subscription/" + sensorProfiles.EntityData.SegmentPath
    sensorProfiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorProfiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorProfiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorProfiles.EntityData.Children = types.NewOrderedMap()
    sensorProfiles.EntityData.Children.Append("sensor-profile", types.YChild{"SensorProfile", nil})
    for i := range sensorProfiles.SensorProfile {
        sensorProfiles.EntityData.Children.Append(types.GetSegmentPath(sensorProfiles.SensorProfile[i]), types.YChild{"SensorProfile", sensorProfiles.SensorProfile[i]})
    }
    sensorProfiles.EntityData.Leafs = types.NewOrderedMap()

    sensorProfiles.EntityData.YListKeys = []string {}

    return &(sensorProfiles.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile
// Associate Sensor Group with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the telemetry sensor group name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Sensorgroupid interface{}

    // use strict timer. The type is interface{}.
    StrictTimer interface{}

    // Sample interval in milliseconds. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory. Units are millisecond.
    SampleInterval interface{}
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_SensorProfiles_SensorProfile) GetEntityData() *types.CommonEntityData {
    sensorProfile.EntityData.YFilter = sensorProfile.YFilter
    sensorProfile.EntityData.YangName = "sensor-profile"
    sensorProfile.EntityData.BundleName = "cisco_ios_xr"
    sensorProfile.EntityData.ParentYangName = "sensor-profiles"
    sensorProfile.EntityData.SegmentPath = "sensor-profile" + types.AddKeyToken(sensorProfile.Sensorgroupid, "sensorgroupid")
    sensorProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/subscriptions/subscription/sensor-profiles/" + sensorProfile.EntityData.SegmentPath
    sensorProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorProfile.EntityData.Children = types.NewOrderedMap()
    sensorProfile.EntityData.Leafs = types.NewOrderedMap()
    sensorProfile.EntityData.Leafs.Append("sensorgroupid", types.YLeaf{"Sensorgroupid", sensorProfile.Sensorgroupid})
    sensorProfile.EntityData.Leafs.Append("strict-timer", types.YLeaf{"StrictTimer", sensorProfile.StrictTimer})
    sensorProfile.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", sensorProfile.SampleInterval})

    sensorProfile.EntityData.YListKeys = []string {"Sensorgroupid"}

    return &(sensorProfile.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles
// Associate Destination Groups with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Associate Destination Group with Subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile.
    DestinationProfile []*TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile
}

func (destinationProfiles *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles) GetEntityData() *types.CommonEntityData {
    destinationProfiles.EntityData.YFilter = destinationProfiles.YFilter
    destinationProfiles.EntityData.YangName = "destination-profiles"
    destinationProfiles.EntityData.BundleName = "cisco_ios_xr"
    destinationProfiles.EntityData.ParentYangName = "subscription"
    destinationProfiles.EntityData.SegmentPath = "destination-profiles"
    destinationProfiles.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/subscriptions/subscription/" + destinationProfiles.EntityData.SegmentPath
    destinationProfiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationProfiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationProfiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationProfiles.EntityData.Children = types.NewOrderedMap()
    destinationProfiles.EntityData.Children.Append("destination-profile", types.YChild{"DestinationProfile", nil})
    for i := range destinationProfiles.DestinationProfile {
        destinationProfiles.EntityData.Children.Append(types.GetSegmentPath(destinationProfiles.DestinationProfile[i]), types.YChild{"DestinationProfile", destinationProfiles.DestinationProfile[i]})
    }
    destinationProfiles.EntityData.Leafs = types.NewOrderedMap()

    destinationProfiles.EntityData.YListKeys = []string {}

    return &(destinationProfiles.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile
// Associate Destination Group with Subscription
type TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Destination Id to associate with Subscription. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    DestinationId interface{}
}

func (destinationProfile *TelemetryModelDriven_Subscriptions_Subscription_DestinationProfiles_DestinationProfile) GetEntityData() *types.CommonEntityData {
    destinationProfile.EntityData.YFilter = destinationProfile.YFilter
    destinationProfile.EntityData.YangName = "destination-profile"
    destinationProfile.EntityData.BundleName = "cisco_ios_xr"
    destinationProfile.EntityData.ParentYangName = "destination-profiles"
    destinationProfile.EntityData.SegmentPath = "destination-profile" + types.AddKeyToken(destinationProfile.DestinationId, "destination-id")
    destinationProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/subscriptions/subscription/destination-profiles/" + destinationProfile.EntityData.SegmentPath
    destinationProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationProfile.EntityData.Children = types.NewOrderedMap()
    destinationProfile.EntityData.Leafs = types.NewOrderedMap()
    destinationProfile.EntityData.Leafs.Append("destination-id", types.YLeaf{"DestinationId", destinationProfile.DestinationId})

    destinationProfile.EntityData.YListKeys = []string {"DestinationId"}

    return &(destinationProfile.EntityData)
}

// TelemetryModelDriven_Include
// Include fields with empty values in output.
type TelemetryModelDriven_Include struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Include fields with empty values in output.
    Empty TelemetryModelDriven_Include_Empty
}

func (include *TelemetryModelDriven_Include) GetEntityData() *types.CommonEntityData {
    include.EntityData.YFilter = include.YFilter
    include.EntityData.YangName = "include"
    include.EntityData.BundleName = "cisco_ios_xr"
    include.EntityData.ParentYangName = "telemetry-model-driven"
    include.EntityData.SegmentPath = "include"
    include.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/" + include.EntityData.SegmentPath
    include.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    include.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    include.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    include.EntityData.Children = types.NewOrderedMap()
    include.EntityData.Children.Append("empty", types.YChild{"Empty", &include.Empty})
    include.EntityData.Leafs = types.NewOrderedMap()

    include.EntityData.YListKeys = []string {}

    return &(include.EntityData)
}

// TelemetryModelDriven_Include_Empty
// Include fields with empty values in output.
type TelemetryModelDriven_Include_Empty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // String type fields with empty string value, for example, are omitted by
    // default. This provides an option to override this behavior and include them
    // in the output. The type is interface{}.
    Values interface{}
}

func (empty *TelemetryModelDriven_Include_Empty) GetEntityData() *types.CommonEntityData {
    empty.EntityData.YFilter = empty.YFilter
    empty.EntityData.YangName = "empty"
    empty.EntityData.BundleName = "cisco_ios_xr"
    empty.EntityData.ParentYangName = "include"
    empty.EntityData.SegmentPath = "empty"
    empty.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/include/" + empty.EntityData.SegmentPath
    empty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    empty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    empty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    empty.EntityData.Children = types.NewOrderedMap()
    empty.EntityData.Leafs = types.NewOrderedMap()
    empty.EntityData.Leafs.Append("values", types.YLeaf{"Values", empty.Values})

    empty.EntityData.YListKeys = []string {}

    return &(empty.EntityData)
}

// TelemetryModelDriven_DestinationGroups
// Destination Group configuration
type TelemetryModelDriven_DestinationGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination Group. The type is slice of
    // TelemetryModelDriven_DestinationGroups_DestinationGroup.
    DestinationGroup []*TelemetryModelDriven_DestinationGroups_DestinationGroup
}

func (destinationGroups *TelemetryModelDriven_DestinationGroups) GetEntityData() *types.CommonEntityData {
    destinationGroups.EntityData.YFilter = destinationGroups.YFilter
    destinationGroups.EntityData.YangName = "destination-groups"
    destinationGroups.EntityData.BundleName = "cisco_ios_xr"
    destinationGroups.EntityData.ParentYangName = "telemetry-model-driven"
    destinationGroups.EntityData.SegmentPath = "destination-groups"
    destinationGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/" + destinationGroups.EntityData.SegmentPath
    destinationGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationGroups.EntityData.Children = types.NewOrderedMap()
    destinationGroups.EntityData.Children.Append("destination-group", types.YChild{"DestinationGroup", nil})
    for i := range destinationGroups.DestinationGroup {
        destinationGroups.EntityData.Children.Append(types.GetSegmentPath(destinationGroups.DestinationGroup[i]), types.YChild{"DestinationGroup", destinationGroups.DestinationGroup[i]})
    }
    destinationGroups.EntityData.Leafs = types.NewOrderedMap()

    destinationGroups.EntityData.YListKeys = []string {}

    return &(destinationGroups.EntityData)
}

// TelemetryModelDriven_DestinationGroups_DestinationGroup
// Destination Group
type TelemetryModelDriven_DestinationGroups_DestinationGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. destination group id string. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    DestinationId interface{}

    // Vrf for the destination group. The type is string with length: 1..32.
    Vrf interface{}

    // Destination address configuration.
    Ipv6Destinations TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations

    // Destination address configuration.
    Ipv4Destinations TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations
}

func (destinationGroup *TelemetryModelDriven_DestinationGroups_DestinationGroup) GetEntityData() *types.CommonEntityData {
    destinationGroup.EntityData.YFilter = destinationGroup.YFilter
    destinationGroup.EntityData.YangName = "destination-group"
    destinationGroup.EntityData.BundleName = "cisco_ios_xr"
    destinationGroup.EntityData.ParentYangName = "destination-groups"
    destinationGroup.EntityData.SegmentPath = "destination-group" + types.AddKeyToken(destinationGroup.DestinationId, "destination-id")
    destinationGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/destination-groups/" + destinationGroup.EntityData.SegmentPath
    destinationGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationGroup.EntityData.Children = types.NewOrderedMap()
    destinationGroup.EntityData.Children.Append("ipv6-destinations", types.YChild{"Ipv6Destinations", &destinationGroup.Ipv6Destinations})
    destinationGroup.EntityData.Children.Append("ipv4-destinations", types.YChild{"Ipv4Destinations", &destinationGroup.Ipv4Destinations})
    destinationGroup.EntityData.Leafs = types.NewOrderedMap()
    destinationGroup.EntityData.Leafs.Append("destination-id", types.YLeaf{"DestinationId", destinationGroup.DestinationId})
    destinationGroup.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", destinationGroup.Vrf})

    destinationGroup.EntityData.YListKeys = []string {"DestinationId"}

    return &(destinationGroup.EntityData)
}

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations
// Destination address configuration
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // destination IP address. The type is slice of
    // TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination.
    Ipv6Destination []*TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination
}

func (ipv6Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations) GetEntityData() *types.CommonEntityData {
    ipv6Destinations.EntityData.YFilter = ipv6Destinations.YFilter
    ipv6Destinations.EntityData.YangName = "ipv6-destinations"
    ipv6Destinations.EntityData.BundleName = "cisco_ios_xr"
    ipv6Destinations.EntityData.ParentYangName = "destination-group"
    ipv6Destinations.EntityData.SegmentPath = "ipv6-destinations"
    ipv6Destinations.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/destination-groups/destination-group/" + ipv6Destinations.EntityData.SegmentPath
    ipv6Destinations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Destinations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Destinations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Destinations.EntityData.Children = types.NewOrderedMap()
    ipv6Destinations.EntityData.Children.Append("ipv6-destination", types.YChild{"Ipv6Destination", nil})
    for i := range ipv6Destinations.Ipv6Destination {
        ipv6Destinations.EntityData.Children.Append(types.GetSegmentPath(ipv6Destinations.Ipv6Destination[i]), types.YChild{"Ipv6Destination", ipv6Destinations.Ipv6Destination[i]})
    }
    ipv6Destinations.EntityData.Leafs = types.NewOrderedMap()

    ipv6Destinations.EntityData.YListKeys = []string {}

    return &(ipv6Destinations.EntityData)
}

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination
// destination IP address
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Destination IPv6 address. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (ipv6Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination) GetEntityData() *types.CommonEntityData {
    ipv6Destination.EntityData.YFilter = ipv6Destination.YFilter
    ipv6Destination.EntityData.YangName = "ipv6-destination"
    ipv6Destination.EntityData.BundleName = "cisco_ios_xr"
    ipv6Destination.EntityData.ParentYangName = "ipv6-destinations"
    ipv6Destination.EntityData.SegmentPath = "ipv6-destination" + types.AddKeyToken(ipv6Destination.Ipv6Address, "ipv6-address") + types.AddKeyToken(ipv6Destination.DestinationPort, "destination-port")
    ipv6Destination.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/destination-groups/destination-group/ipv6-destinations/" + ipv6Destination.EntityData.SegmentPath
    ipv6Destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Destination.EntityData.Children = types.NewOrderedMap()
    ipv6Destination.EntityData.Children.Append("protocol", types.YChild{"Protocol", &ipv6Destination.Protocol})
    ipv6Destination.EntityData.Leafs = types.NewOrderedMap()
    ipv6Destination.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Destination.Ipv6Address})
    ipv6Destination.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", ipv6Destination.DestinationPort})
    ipv6Destination.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", ipv6Destination.Encoding})

    ipv6Destination.EntityData.YListKeys = []string {"Ipv6Address", "DestinationPort"}

    return &(ipv6Destination.EntityData)
}

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol
// Transport Protocol used to transmit telemetry
// data to the collector
// This type is a presence type.
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // protocol. The type is ProtoType. This attribute is mandatory.
    Protocol interface{}

    // tls hostname. The type is string.
    TlsHostname interface{}

    // no tls. The type is interface{}.
    NoTls interface{}

    // udp packetsize. The type is interface{} with range: 484..65507. The default
    // value is 1472.
    Packetsize interface{}
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv6Destinations_Ipv6Destination_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "ipv6-destination"
    protocol.EntityData.SegmentPath = "protocol"
    protocol.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/destination-groups/destination-group/ipv6-destinations/ipv6-destination/" + protocol.EntityData.SegmentPath
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", protocol.Protocol})
    protocol.EntityData.Leafs.Append("tls-hostname", types.YLeaf{"TlsHostname", protocol.TlsHostname})
    protocol.EntityData.Leafs.Append("no-tls", types.YLeaf{"NoTls", protocol.NoTls})
    protocol.EntityData.Leafs.Append("packetsize", types.YLeaf{"Packetsize", protocol.Packetsize})

    protocol.EntityData.YListKeys = []string {}

    return &(protocol.EntityData)
}

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations
// Destination address configuration
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // destination IP address. The type is slice of
    // TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination.
    Ipv4Destination []*TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination
}

func (ipv4Destinations *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations) GetEntityData() *types.CommonEntityData {
    ipv4Destinations.EntityData.YFilter = ipv4Destinations.YFilter
    ipv4Destinations.EntityData.YangName = "ipv4-destinations"
    ipv4Destinations.EntityData.BundleName = "cisco_ios_xr"
    ipv4Destinations.EntityData.ParentYangName = "destination-group"
    ipv4Destinations.EntityData.SegmentPath = "ipv4-destinations"
    ipv4Destinations.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/destination-groups/destination-group/" + ipv4Destinations.EntityData.SegmentPath
    ipv4Destinations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Destinations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Destinations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Destinations.EntityData.Children = types.NewOrderedMap()
    ipv4Destinations.EntityData.Children.Append("ipv4-destination", types.YChild{"Ipv4Destination", nil})
    for i := range ipv4Destinations.Ipv4Destination {
        ipv4Destinations.EntityData.Children.Append(types.GetSegmentPath(ipv4Destinations.Ipv4Destination[i]), types.YChild{"Ipv4Destination", ipv4Destinations.Ipv4Destination[i]})
    }
    ipv4Destinations.EntityData.Leafs = types.NewOrderedMap()

    ipv4Destinations.EntityData.YListKeys = []string {}

    return &(ipv4Destinations.EntityData)
}

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination
// destination IP address
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Destination IPv4 address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipv4Destination *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination) GetEntityData() *types.CommonEntityData {
    ipv4Destination.EntityData.YFilter = ipv4Destination.YFilter
    ipv4Destination.EntityData.YangName = "ipv4-destination"
    ipv4Destination.EntityData.BundleName = "cisco_ios_xr"
    ipv4Destination.EntityData.ParentYangName = "ipv4-destinations"
    ipv4Destination.EntityData.SegmentPath = "ipv4-destination" + types.AddKeyToken(ipv4Destination.Ipv4Address, "ipv4-address") + types.AddKeyToken(ipv4Destination.DestinationPort, "destination-port")
    ipv4Destination.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/destination-groups/destination-group/ipv4-destinations/" + ipv4Destination.EntityData.SegmentPath
    ipv4Destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Destination.EntityData.Children = types.NewOrderedMap()
    ipv4Destination.EntityData.Children.Append("protocol", types.YChild{"Protocol", &ipv4Destination.Protocol})
    ipv4Destination.EntityData.Leafs = types.NewOrderedMap()
    ipv4Destination.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", ipv4Destination.Ipv4Address})
    ipv4Destination.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", ipv4Destination.DestinationPort})
    ipv4Destination.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", ipv4Destination.Encoding})

    ipv4Destination.EntityData.YListKeys = []string {"Ipv4Address", "DestinationPort"}

    return &(ipv4Destination.EntityData)
}

// TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol
// Transport Protocol used to transmit telemetry
// data to the collector
// This type is a presence type.
type TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // protocol. The type is ProtoType. This attribute is mandatory.
    Protocol interface{}

    // tls hostname. The type is string.
    TlsHostname interface{}

    // no tls. The type is interface{}.
    NoTls interface{}

    // udp packetsize. The type is interface{} with range: 484..65507. The default
    // value is 1472.
    Packetsize interface{}
}

func (protocol *TelemetryModelDriven_DestinationGroups_DestinationGroup_Ipv4Destinations_Ipv4Destination_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "ipv4-destination"
    protocol.EntityData.SegmentPath = "protocol"
    protocol.EntityData.AbsolutePath = "Cisco-IOS-XR-telemetry-model-driven-cfg:telemetry-model-driven/destination-groups/destination-group/ipv4-destinations/ipv4-destination/" + protocol.EntityData.SegmentPath
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", protocol.Protocol})
    protocol.EntityData.Leafs.Append("tls-hostname", types.YLeaf{"TlsHostname", protocol.TlsHostname})
    protocol.EntityData.Leafs.Append("no-tls", types.YLeaf{"NoTls", protocol.NoTls})
    protocol.EntityData.Leafs.Append("packetsize", types.YLeaf{"Packetsize", protocol.Packetsize})

    protocol.EntityData.YListKeys = []string {}

    return &(protocol.EntityData)
}

