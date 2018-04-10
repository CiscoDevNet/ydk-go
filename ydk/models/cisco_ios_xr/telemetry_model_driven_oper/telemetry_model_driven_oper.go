// This module contains a collection of YANG definitions
// for Cisco IOS-XR telemetry-model-driven package operational data.
// 
// This module contains definitions
// for the following management objects:
//   telemetry-model-driven: Telemetry operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package telemetry_model_driven_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package telemetry_model_driven_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-telemetry-model-driven-oper telemetry-model-driven}", reflect.TypeOf(TelemetryModelDriven{}))
    ydk.RegisterEntity("Cisco-IOS-XR-telemetry-model-driven-oper:telemetry-model-driven", reflect.TypeOf(TelemetryModelDriven{}))
}

// MdtSourceQosMarking represents DSCP source qos value for subscription
type MdtSourceQosMarking string

const (
    // 0
    MdtSourceQosMarking_dscp_default MdtSourceQosMarking = "dscp-default"

    // 8
    MdtSourceQosMarking_dscp_cs1 MdtSourceQosMarking = "dscp-cs1"

    // 10
    MdtSourceQosMarking_dscp_af11 MdtSourceQosMarking = "dscp-af11"

    // 12
    MdtSourceQosMarking_dscp_af12 MdtSourceQosMarking = "dscp-af12"

    // 14
    MdtSourceQosMarking_dscp_af13 MdtSourceQosMarking = "dscp-af13"

    // 16
    MdtSourceQosMarking_dscp_cs2 MdtSourceQosMarking = "dscp-cs2"

    // 18
    MdtSourceQosMarking_dscp_af21 MdtSourceQosMarking = "dscp-af21"

    // 20
    MdtSourceQosMarking_dscp_af22 MdtSourceQosMarking = "dscp-af22"

    // 22
    MdtSourceQosMarking_dscp_af23 MdtSourceQosMarking = "dscp-af23"

    // 24
    MdtSourceQosMarking_dscp_cs3 MdtSourceQosMarking = "dscp-cs3"

    // 26
    MdtSourceQosMarking_dscp_af31 MdtSourceQosMarking = "dscp-af31"

    // 28
    MdtSourceQosMarking_dscp_af32 MdtSourceQosMarking = "dscp-af32"

    // 30
    MdtSourceQosMarking_dscp_af33 MdtSourceQosMarking = "dscp-af33"

    // 32
    MdtSourceQosMarking_dscp_cs4 MdtSourceQosMarking = "dscp-cs4"

    // 34
    MdtSourceQosMarking_dscp_af41 MdtSourceQosMarking = "dscp-af41"

    // 36
    MdtSourceQosMarking_dscp_af42 MdtSourceQosMarking = "dscp-af42"

    // 38
    MdtSourceQosMarking_dscp_af43 MdtSourceQosMarking = "dscp-af43"

    // 40
    MdtSourceQosMarking_dscp_cs5 MdtSourceQosMarking = "dscp-cs5"

    // 46
    MdtSourceQosMarking_dscp_ef MdtSourceQosMarking = "dscp-ef"

    // 48
    MdtSourceQosMarking_dscp_cs6 MdtSourceQosMarking = "dscp-cs6"

    // 56
    MdtSourceQosMarking_dscp_cs7 MdtSourceQosMarking = "dscp-cs7"
)

// MdtSubsStateEnum represents Subscription State
type MdtSubsStateEnum string

const (
    // NA
    MdtSubsStateEnum_not_active MdtSubsStateEnum = "not-active"

    // Active
    MdtSubsStateEnum_active MdtSubsStateEnum = "active"

    // Paused
    MdtSubsStateEnum_paused MdtSubsStateEnum = "paused"
)

// MdtInternalPathStatus represents Internal Subscription Path Status
type MdtInternalPathStatus string

const (
    // Active
    MdtInternalPathStatus_active MdtInternalPathStatus = "active"

    // Internal Error
    MdtInternalPathStatus_internal_err MdtInternalPathStatus = "internal-err"

    // Plugin Active
    MdtInternalPathStatus_plugin_active MdtInternalPathStatus = "plugin-active"

    // Plugin Not Initialized
    MdtInternalPathStatus_plugin_not_initialized MdtInternalPathStatus = "plugin-not-initialized"

    // Plugin Unsupported Cadence
    MdtInternalPathStatus_plugin_invalid_cadence MdtInternalPathStatus = "plugin-invalid-cadence"

    // Plugin Subscription Error
    MdtInternalPathStatus_plugin_err MdtInternalPathStatus = "plugin-err"

    // Filter Error
    MdtInternalPathStatus_filter_err MdtInternalPathStatus = "filter-err"

    // Paused
    MdtInternalPathStatus_paused MdtInternalPathStatus = "paused"

    // Eventing Active
    MdtInternalPathStatus_event_ing_active MdtInternalPathStatus = "event-ing-active"

    // Eventing Not Active
    MdtInternalPathStatus_event_ing_not_active MdtInternalPathStatus = "event-ing-not-active"

    // Eventing Error
    MdtInternalPathStatus_event_ing_err MdtInternalPathStatus = "event-ing-err"
)

// MdtDestStateEnum represents Destination state
type MdtDestStateEnum string

const (
    // NA
    MdtDestStateEnum_dest_not_active MdtDestStateEnum = "dest-not-active"

    // Active
    MdtDestStateEnum_dest_active MdtDestStateEnum = "dest-active"

    // AskingPause
    MdtDestStateEnum_dest_asking_pause MdtDestStateEnum = "dest-asking-pause"

    // Paused
    MdtDestStateEnum_dest_paused MdtDestStateEnum = "dest-paused"

    // Resuming
    MdtDestStateEnum_dest_resuming MdtDestStateEnum = "dest-resuming"

    // ChannelNotFound
    MdtDestStateEnum_dest_channel_not_found MdtDestStateEnum = "dest-channel-not-found"
)

// MdtTransportEnum represents MDT Transport
type MdtTransportEnum string

const (
    // PROTOCOL NOT SET
    MdtTransportEnum_not_set MdtTransportEnum = "not-set"

    // GRPC
    MdtTransportEnum_grpc MdtTransportEnum = "grpc"

    // TCP
    MdtTransportEnum_tcp MdtTransportEnum = "tcp"

    // UDP
    MdtTransportEnum_udp MdtTransportEnum = "udp"

    // DIALIN
    MdtTransportEnum_dialin MdtTransportEnum = "dialin"
)

// MdtEncodingEnum represents MDT Encoding
type MdtEncodingEnum string

const (
    // ENCODING NOT SET
    MdtEncodingEnum_not_set MdtEncodingEnum = "not-set"

    // GPB
    MdtEncodingEnum_gpb MdtEncodingEnum = "gpb"

    // SELF DESCRIBING GPB
    MdtEncodingEnum_self_describing_gpb MdtEncodingEnum = "self-describing-gpb"

    // JSON
    MdtEncodingEnum_json MdtEncodingEnum = "json"
)

// MdtIp represents IP Type
type MdtIp string

const (
    // IPv4
    MdtIp_ipv4 MdtIp = "ipv4"

    // IPv6
    MdtIp_ipv6 MdtIp = "ipv6"
)

// TelemetryModelDriven
// Telemetry operational data
type TelemetryModelDriven struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Telemetry Destinations.
    Destinations TelemetryModelDriven_Destinations

    // Telemetry Subscriptions.
    Subscriptions TelemetryModelDriven_Subscriptions

    // Telemetry Sensor Groups.
    SensorGroups TelemetryModelDriven_SensorGroups
}

func (telemetryModelDriven *TelemetryModelDriven) GetEntityData() *types.CommonEntityData {
    telemetryModelDriven.EntityData.YFilter = telemetryModelDriven.YFilter
    telemetryModelDriven.EntityData.YangName = "telemetry-model-driven"
    telemetryModelDriven.EntityData.BundleName = "cisco_ios_xr"
    telemetryModelDriven.EntityData.ParentYangName = "Cisco-IOS-XR-telemetry-model-driven-oper"
    telemetryModelDriven.EntityData.SegmentPath = "Cisco-IOS-XR-telemetry-model-driven-oper:telemetry-model-driven"
    telemetryModelDriven.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telemetryModelDriven.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telemetryModelDriven.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telemetryModelDriven.EntityData.Children = make(map[string]types.YChild)
    telemetryModelDriven.EntityData.Children["destinations"] = types.YChild{"Destinations", &telemetryModelDriven.Destinations}
    telemetryModelDriven.EntityData.Children["subscriptions"] = types.YChild{"Subscriptions", &telemetryModelDriven.Subscriptions}
    telemetryModelDriven.EntityData.Children["sensor-groups"] = types.YChild{"SensorGroups", &telemetryModelDriven.SensorGroups}
    telemetryModelDriven.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(telemetryModelDriven.EntityData)
}

// TelemetryModelDriven_Destinations
// Telemetry Destinations
type TelemetryModelDriven_Destinations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Telemetry Destination. The type is slice of
    // TelemetryModelDriven_Destinations_Destination.
    Destination []TelemetryModelDriven_Destinations_Destination
}

func (destinations *TelemetryModelDriven_Destinations) GetEntityData() *types.CommonEntityData {
    destinations.EntityData.YFilter = destinations.YFilter
    destinations.EntityData.YangName = "destinations"
    destinations.EntityData.BundleName = "cisco_ios_xr"
    destinations.EntityData.ParentYangName = "telemetry-model-driven"
    destinations.EntityData.SegmentPath = "destinations"
    destinations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinations.EntityData.Children = make(map[string]types.YChild)
    destinations.EntityData.Children["destination"] = types.YChild{"Destination", nil}
    for i := range destinations.Destination {
        destinations.EntityData.Children[types.GetSegmentPath(&destinations.Destination[i])] = types.YChild{"Destination", &destinations.Destination[i]}
    }
    destinations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(destinations.EntityData)
}

// TelemetryModelDriven_Destinations_Destination
// Telemetry Destination
type TelemetryModelDriven_Destinations_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Id of the destination. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    DestinationId interface{}

    // Destination Group name. The type is string.
    Id interface{}

    // Set if this is configured destination group. The type is interface{} with
    // range: 0..4294967295.
    Configured interface{}

    // list of destinations defined in this group. The type is slice of
    // TelemetryModelDriven_Destinations_Destination_Destination.
    Destination []TelemetryModelDriven_Destinations_Destination_Destination_
}

func (destination *TelemetryModelDriven_Destinations_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "destinations"
    destination.EntityData.SegmentPath = "destination" + "[destination-id='" + fmt.Sprintf("%v", destination.DestinationId) + "']"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = make(map[string]types.YChild)
    destination.EntityData.Children["destination"] = types.YChild{"Destination", nil}
    for i := range destination.Destination {
        destination.EntityData.Children[types.GetSegmentPath(&destination.Destination[i])] = types.YChild{"Destination", &destination.Destination[i]}
    }
    destination.EntityData.Leafs = make(map[string]types.YLeaf)
    destination.EntityData.Leafs["destination-id"] = types.YLeaf{"DestinationId", destination.DestinationId}
    destination.EntityData.Leafs["id"] = types.YLeaf{"Id", destination.Id}
    destination.EntityData.Leafs["configured"] = types.YLeaf{"Configured", destination.Configured}
    return &(destination.EntityData)
}

// TelemetryModelDriven_Destinations_Destination_Destination_
// list of destinations defined in this group
type TelemetryModelDriven_Destinations_Destination_Destination_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination.
    Destination TelemetryModelDriven_Destinations_Destination_Destination__Destination_

    // List of collection groups for this destination group. The type is slice of
    // TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup.
    CollectionGroup []TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup
}

func (destination_ *TelemetryModelDriven_Destinations_Destination_Destination_) GetEntityData() *types.CommonEntityData {
    destination_.EntityData.YFilter = destination_.YFilter
    destination_.EntityData.YangName = "destination"
    destination_.EntityData.BundleName = "cisco_ios_xr"
    destination_.EntityData.ParentYangName = "destination"
    destination_.EntityData.SegmentPath = "destination"
    destination_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination_.EntityData.Children = make(map[string]types.YChild)
    destination_.EntityData.Children["destination"] = types.YChild{"Destination", &destination_.Destination}
    destination_.EntityData.Children["collection-group"] = types.YChild{"CollectionGroup", nil}
    for i := range destination_.CollectionGroup {
        destination_.EntityData.Children[types.GetSegmentPath(&destination_.CollectionGroup[i])] = types.YChild{"CollectionGroup", &destination_.CollectionGroup[i]}
    }
    destination_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(destination_.EntityData)
}

// TelemetryModelDriven_Destinations_Destination_Destination__Destination_
// Destination
type TelemetryModelDriven_Destinations_Destination_Destination__Destination_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination Id. The type is string.
    Id interface{}

    // Sub Idstr. The type is string.
    SubIdStr interface{}

    // Destination Port number. The type is interface{} with range: 0..65535.
    DestPort interface{}

    // Destination group encoding. The type is MdtEncodingEnum.
    Encoding interface{}

    // Destination group transport. The type is MdtTransportEnum.
    Transport interface{}

    // Destination group vrf. The type is string.
    Vrf interface{}

    // Destination group vrf id. The type is interface{} with range:
    // 0..4294967295.
    VrfId interface{}

    // State of streaming on this destination. The type is MdtDestStateEnum.
    State interface{}

    // UDP MTU if this destination is UDP. The type is interface{} with range:
    // 0..4294967295.
    UdpMtu interface{}

    // TLS connection to this destination. The type is interface{} with range:
    // 0..4294967295.
    Tls interface{}

    // TLS Hostname of this destination. The type is string.
    TlsHost interface{}

    // Total number of packets sent for this destination. The type is interface{}
    // with range: 0..18446744073709551615.
    TotalNumOfPacketsSent interface{}

    // Total number of bytes sent for this destination. The type is interface{}
    // with range: 0..18446744073709551615. Units are byte.
    TotalNumOfBytesSent interface{}

    // Timestamp of the last collection. The type is interface{} with range:
    // 0..18446744073709551615.
    LastCollectionTime interface{}

    // DSCP setting for this destination. The type is interface{} with range:
    // 0..4294967295.
    Dscp interface{}

    // Sub Id. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    SubId []interface{}

    // Destination IP Address.
    DestIpAddress TelemetryModelDriven_Destinations_Destination_Destination__Destination__DestIpAddress
}

func (destination_ *TelemetryModelDriven_Destinations_Destination_Destination__Destination_) GetEntityData() *types.CommonEntityData {
    destination_.EntityData.YFilter = destination_.YFilter
    destination_.EntityData.YangName = "destination"
    destination_.EntityData.BundleName = "cisco_ios_xr"
    destination_.EntityData.ParentYangName = "destination"
    destination_.EntityData.SegmentPath = "destination"
    destination_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination_.EntityData.Children = make(map[string]types.YChild)
    destination_.EntityData.Children["dest-ip-address"] = types.YChild{"DestIpAddress", &destination_.DestIpAddress}
    destination_.EntityData.Leafs = make(map[string]types.YLeaf)
    destination_.EntityData.Leafs["id"] = types.YLeaf{"Id", destination_.Id}
    destination_.EntityData.Leafs["sub-id-str"] = types.YLeaf{"SubIdStr", destination_.SubIdStr}
    destination_.EntityData.Leafs["dest-port"] = types.YLeaf{"DestPort", destination_.DestPort}
    destination_.EntityData.Leafs["encoding"] = types.YLeaf{"Encoding", destination_.Encoding}
    destination_.EntityData.Leafs["transport"] = types.YLeaf{"Transport", destination_.Transport}
    destination_.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", destination_.Vrf}
    destination_.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", destination_.VrfId}
    destination_.EntityData.Leafs["state"] = types.YLeaf{"State", destination_.State}
    destination_.EntityData.Leafs["udp-mtu"] = types.YLeaf{"UdpMtu", destination_.UdpMtu}
    destination_.EntityData.Leafs["tls"] = types.YLeaf{"Tls", destination_.Tls}
    destination_.EntityData.Leafs["tls-host"] = types.YLeaf{"TlsHost", destination_.TlsHost}
    destination_.EntityData.Leafs["total-num-of-packets-sent"] = types.YLeaf{"TotalNumOfPacketsSent", destination_.TotalNumOfPacketsSent}
    destination_.EntityData.Leafs["total-num-of-bytes-sent"] = types.YLeaf{"TotalNumOfBytesSent", destination_.TotalNumOfBytesSent}
    destination_.EntityData.Leafs["last-collection-time"] = types.YLeaf{"LastCollectionTime", destination_.LastCollectionTime}
    destination_.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", destination_.Dscp}
    destination_.EntityData.Leafs["sub-id"] = types.YLeaf{"SubId", destination_.SubId}
    return &(destination_.EntityData)
}

// TelemetryModelDriven_Destinations_Destination_Destination__Destination__DestIpAddress
// Destination IP Address
type TelemetryModelDriven_Destinations_Destination_Destination__Destination__DestIpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPType. The type is MdtIp.
    IpType interface{}

    // IPV4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination__Destination__DestIpAddress) GetEntityData() *types.CommonEntityData {
    destIpAddress.EntityData.YFilter = destIpAddress.YFilter
    destIpAddress.EntityData.YangName = "dest-ip-address"
    destIpAddress.EntityData.BundleName = "cisco_ios_xr"
    destIpAddress.EntityData.ParentYangName = "destination"
    destIpAddress.EntityData.SegmentPath = "dest-ip-address"
    destIpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destIpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destIpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destIpAddress.EntityData.Children = make(map[string]types.YChild)
    destIpAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    destIpAddress.EntityData.Leafs["ip-type"] = types.YLeaf{"IpType", destIpAddress.IpType}
    destIpAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", destIpAddress.Ipv4Address}
    destIpAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", destIpAddress.Ipv6Address}
    return &(destIpAddress.EntityData)
}

// TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup
// List of collection groups for this destination
// group
type TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection Group id. The type is interface{} with range:
    // 0..18446744073709551615.
    Id interface{}

    // Period of the collections (ms). The type is interface{} with range:
    // 0..4294967295.
    Cadence interface{}

    // Completed collections count. The type is interface{} with range:
    // 0..4294967295.
    TotalCollections interface{}

    // Destination group encoding. The type is MdtEncodingEnum.
    Encoding interface{}

    // Timestamp of the start of last collection. The type is interface{} with
    // range: 0..18446744073709551615.
    LastCollectionStartTime interface{}

    // Timestamp of the end of last collection. The type is interface{} with
    // range: 0..18446744073709551615.
    LastCollectionEndTime interface{}

    // Maximum time for a collection (ms). The type is interface{} with range:
    // 0..4294967295.
    MaxCollectionTime interface{}

    // Minimum time for a collection (ms). The type is interface{} with range:
    // 0..4294967295.
    MinCollectionTime interface{}

    // Minimum time for all processing (ms). The type is interface{} with range:
    // 0..4294967295.
    MinTotalTime interface{}

    // Maximum time for all processing (ms). The type is interface{} with range:
    // 0..4294967295.
    MaxTotalTime interface{}

    // Average time for all processing (ms). The type is interface{} with range:
    // 0..4294967295.
    AvgTotalTime interface{}

    // Total number of errors. The type is interface{} with range: 0..4294967295.
    TotalOtherErrors interface{}

    // Total number of no data instances. The type is interface{} with range:
    // 0..4294967295.
    TotalOnDataInstances interface{}

    // Total number skipped (not ready). The type is interface{} with range:
    // 0..4294967295.
    TotalNotReady interface{}

    // Total number of send errors. The type is interface{} with range:
    // 0..4294967295.
    TotalSendErrors interface{}

    // Total number of send drops. The type is interface{} with range:
    // 0..4294967295.
    TotalSendDrops interface{}

    // Array of information for sensor paths within collection group. The type is
    // slice of
    // TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_CollectionPath.
    CollectionPath []TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_CollectionPath

    // Array of information for sysdb paths within collection group. The type is
    // slice of
    // TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_InternalCollectionGroup.
    InternalCollectionGroup []TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_InternalCollectionGroup
}

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup) GetEntityData() *types.CommonEntityData {
    collectionGroup.EntityData.YFilter = collectionGroup.YFilter
    collectionGroup.EntityData.YangName = "collection-group"
    collectionGroup.EntityData.BundleName = "cisco_ios_xr"
    collectionGroup.EntityData.ParentYangName = "destination"
    collectionGroup.EntityData.SegmentPath = "collection-group"
    collectionGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectionGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectionGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectionGroup.EntityData.Children = make(map[string]types.YChild)
    collectionGroup.EntityData.Children["collection-path"] = types.YChild{"CollectionPath", nil}
    for i := range collectionGroup.CollectionPath {
        collectionGroup.EntityData.Children[types.GetSegmentPath(&collectionGroup.CollectionPath[i])] = types.YChild{"CollectionPath", &collectionGroup.CollectionPath[i]}
    }
    collectionGroup.EntityData.Children["internal-collection-group"] = types.YChild{"InternalCollectionGroup", nil}
    for i := range collectionGroup.InternalCollectionGroup {
        collectionGroup.EntityData.Children[types.GetSegmentPath(&collectionGroup.InternalCollectionGroup[i])] = types.YChild{"InternalCollectionGroup", &collectionGroup.InternalCollectionGroup[i]}
    }
    collectionGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    collectionGroup.EntityData.Leafs["id"] = types.YLeaf{"Id", collectionGroup.Id}
    collectionGroup.EntityData.Leafs["cadence"] = types.YLeaf{"Cadence", collectionGroup.Cadence}
    collectionGroup.EntityData.Leafs["total-collections"] = types.YLeaf{"TotalCollections", collectionGroup.TotalCollections}
    collectionGroup.EntityData.Leafs["encoding"] = types.YLeaf{"Encoding", collectionGroup.Encoding}
    collectionGroup.EntityData.Leafs["last-collection-start-time"] = types.YLeaf{"LastCollectionStartTime", collectionGroup.LastCollectionStartTime}
    collectionGroup.EntityData.Leafs["last-collection-end-time"] = types.YLeaf{"LastCollectionEndTime", collectionGroup.LastCollectionEndTime}
    collectionGroup.EntityData.Leafs["max-collection-time"] = types.YLeaf{"MaxCollectionTime", collectionGroup.MaxCollectionTime}
    collectionGroup.EntityData.Leafs["min-collection-time"] = types.YLeaf{"MinCollectionTime", collectionGroup.MinCollectionTime}
    collectionGroup.EntityData.Leafs["min-total-time"] = types.YLeaf{"MinTotalTime", collectionGroup.MinTotalTime}
    collectionGroup.EntityData.Leafs["max-total-time"] = types.YLeaf{"MaxTotalTime", collectionGroup.MaxTotalTime}
    collectionGroup.EntityData.Leafs["avg-total-time"] = types.YLeaf{"AvgTotalTime", collectionGroup.AvgTotalTime}
    collectionGroup.EntityData.Leafs["total-other-errors"] = types.YLeaf{"TotalOtherErrors", collectionGroup.TotalOtherErrors}
    collectionGroup.EntityData.Leafs["total-on-data-instances"] = types.YLeaf{"TotalOnDataInstances", collectionGroup.TotalOnDataInstances}
    collectionGroup.EntityData.Leafs["total-not-ready"] = types.YLeaf{"TotalNotReady", collectionGroup.TotalNotReady}
    collectionGroup.EntityData.Leafs["total-send-errors"] = types.YLeaf{"TotalSendErrors", collectionGroup.TotalSendErrors}
    collectionGroup.EntityData.Leafs["total-send-drops"] = types.YLeaf{"TotalSendDrops", collectionGroup.TotalSendDrops}
    return &(collectionGroup.EntityData)
}

// TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_CollectionPath
// Array of information for sensor paths within
// collection group
type TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_CollectionPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_CollectionPath) GetEntityData() *types.CommonEntityData {
    collectionPath.EntityData.YFilter = collectionPath.YFilter
    collectionPath.EntityData.YangName = "collection-path"
    collectionPath.EntityData.BundleName = "cisco_ios_xr"
    collectionPath.EntityData.ParentYangName = "collection-group"
    collectionPath.EntityData.SegmentPath = "collection-path"
    collectionPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectionPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectionPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectionPath.EntityData.Children = make(map[string]types.YChild)
    collectionPath.EntityData.Leafs = make(map[string]types.YLeaf)
    collectionPath.EntityData.Leafs["path"] = types.YLeaf{"Path", collectionPath.Path}
    collectionPath.EntityData.Leafs["state"] = types.YLeaf{"State", collectionPath.State}
    collectionPath.EntityData.Leafs["status-str"] = types.YLeaf{"StatusStr", collectionPath.StatusStr}
    return &(collectionPath.EntityData)
}

// TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_InternalCollectionGroup
// Array of information for sysdb paths within
// collection group
type TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_InternalCollectionGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sysdb Path. The type is string.
    Path interface{}

    // Period of the collections (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    Cadence interface{}

    // Total number of gets. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetCount interface{}

    // Total number of lists. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalListCount interface{}

    // Total number of datalists. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalDatalistCount interface{}

    // Total number of finddata. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalFinddataCount interface{}

    // Total number of get bulk. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetBulkCount interface{}

    // Total number of items retrived from sysdb. The type is interface{} with
    // range: 0..18446744073709551615.
    TotalItemCount interface{}

    // Total number of get errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetErrors interface{}

    // Total number of list errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalListErrors interface{}

    // Total number of datalist errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalDatalistErrors interface{}

    // Total number of finddata errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalFinddataErrors interface{}

    // Total number of get bulk errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetBulkErrors interface{}

    // Total number of encode errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEncodeErrors interface{}

    // Total number of encode deferred. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEncodeNotready interface{}

    // Total number of send errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalSendErrors interface{}

    // Total number of send channel full. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalSendDrops interface{}

    // Total number of bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalSentBytes interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalSendPackets interface{}

    // Total number of send bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalSendBytesDropped interface{}

    // Completed collections count. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalCollections interface{}

    // Total number of collections missed. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalCollectionsMissed interface{}

    // Maximum time for a collection (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    MaxCollectionTime interface{}

    // Minimum time for a collection (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    MinCollectionTime interface{}

    // Average time for a collection (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    AvgCollectionTime interface{}

    // Collection method in use. The type is interface{} with range:
    // 0..18446744073709551615.
    CollectionMethod interface{}

    // Status of collection path. The type is MdtInternalPathStatus.
    Status interface{}
}

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination__CollectionGroup_InternalCollectionGroup) GetEntityData() *types.CommonEntityData {
    internalCollectionGroup.EntityData.YFilter = internalCollectionGroup.YFilter
    internalCollectionGroup.EntityData.YangName = "internal-collection-group"
    internalCollectionGroup.EntityData.BundleName = "cisco_ios_xr"
    internalCollectionGroup.EntityData.ParentYangName = "collection-group"
    internalCollectionGroup.EntityData.SegmentPath = "internal-collection-group"
    internalCollectionGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalCollectionGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalCollectionGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalCollectionGroup.EntityData.Children = make(map[string]types.YChild)
    internalCollectionGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    internalCollectionGroup.EntityData.Leafs["path"] = types.YLeaf{"Path", internalCollectionGroup.Path}
    internalCollectionGroup.EntityData.Leafs["cadence"] = types.YLeaf{"Cadence", internalCollectionGroup.Cadence}
    internalCollectionGroup.EntityData.Leafs["total-get-count"] = types.YLeaf{"TotalGetCount", internalCollectionGroup.TotalGetCount}
    internalCollectionGroup.EntityData.Leafs["total-list-count"] = types.YLeaf{"TotalListCount", internalCollectionGroup.TotalListCount}
    internalCollectionGroup.EntityData.Leafs["total-datalist-count"] = types.YLeaf{"TotalDatalistCount", internalCollectionGroup.TotalDatalistCount}
    internalCollectionGroup.EntityData.Leafs["total-finddata-count"] = types.YLeaf{"TotalFinddataCount", internalCollectionGroup.TotalFinddataCount}
    internalCollectionGroup.EntityData.Leafs["total-get-bulk-count"] = types.YLeaf{"TotalGetBulkCount", internalCollectionGroup.TotalGetBulkCount}
    internalCollectionGroup.EntityData.Leafs["total-item-count"] = types.YLeaf{"TotalItemCount", internalCollectionGroup.TotalItemCount}
    internalCollectionGroup.EntityData.Leafs["total-get-errors"] = types.YLeaf{"TotalGetErrors", internalCollectionGroup.TotalGetErrors}
    internalCollectionGroup.EntityData.Leafs["total-list-errors"] = types.YLeaf{"TotalListErrors", internalCollectionGroup.TotalListErrors}
    internalCollectionGroup.EntityData.Leafs["total-datalist-errors"] = types.YLeaf{"TotalDatalistErrors", internalCollectionGroup.TotalDatalistErrors}
    internalCollectionGroup.EntityData.Leafs["total-finddata-errors"] = types.YLeaf{"TotalFinddataErrors", internalCollectionGroup.TotalFinddataErrors}
    internalCollectionGroup.EntityData.Leafs["total-get-bulk-errors"] = types.YLeaf{"TotalGetBulkErrors", internalCollectionGroup.TotalGetBulkErrors}
    internalCollectionGroup.EntityData.Leafs["total-encode-errors"] = types.YLeaf{"TotalEncodeErrors", internalCollectionGroup.TotalEncodeErrors}
    internalCollectionGroup.EntityData.Leafs["total-encode-notready"] = types.YLeaf{"TotalEncodeNotready", internalCollectionGroup.TotalEncodeNotready}
    internalCollectionGroup.EntityData.Leafs["total-send-errors"] = types.YLeaf{"TotalSendErrors", internalCollectionGroup.TotalSendErrors}
    internalCollectionGroup.EntityData.Leafs["total-send-drops"] = types.YLeaf{"TotalSendDrops", internalCollectionGroup.TotalSendDrops}
    internalCollectionGroup.EntityData.Leafs["total-sent-bytes"] = types.YLeaf{"TotalSentBytes", internalCollectionGroup.TotalSentBytes}
    internalCollectionGroup.EntityData.Leafs["total-send-packets"] = types.YLeaf{"TotalSendPackets", internalCollectionGroup.TotalSendPackets}
    internalCollectionGroup.EntityData.Leafs["total-send-bytes-dropped"] = types.YLeaf{"TotalSendBytesDropped", internalCollectionGroup.TotalSendBytesDropped}
    internalCollectionGroup.EntityData.Leafs["total-collections"] = types.YLeaf{"TotalCollections", internalCollectionGroup.TotalCollections}
    internalCollectionGroup.EntityData.Leafs["total-collections-missed"] = types.YLeaf{"TotalCollectionsMissed", internalCollectionGroup.TotalCollectionsMissed}
    internalCollectionGroup.EntityData.Leafs["max-collection-time"] = types.YLeaf{"MaxCollectionTime", internalCollectionGroup.MaxCollectionTime}
    internalCollectionGroup.EntityData.Leafs["min-collection-time"] = types.YLeaf{"MinCollectionTime", internalCollectionGroup.MinCollectionTime}
    internalCollectionGroup.EntityData.Leafs["avg-collection-time"] = types.YLeaf{"AvgCollectionTime", internalCollectionGroup.AvgCollectionTime}
    internalCollectionGroup.EntityData.Leafs["collection-method"] = types.YLeaf{"CollectionMethod", internalCollectionGroup.CollectionMethod}
    internalCollectionGroup.EntityData.Leafs["status"] = types.YLeaf{"Status", internalCollectionGroup.Status}
    return &(internalCollectionGroup.EntityData)
}

// TelemetryModelDriven_Subscriptions
// Telemetry Subscriptions
type TelemetryModelDriven_Subscriptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Telemetry Subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription.
    Subscription []TelemetryModelDriven_Subscriptions_Subscription
}

func (subscriptions *TelemetryModelDriven_Subscriptions) GetEntityData() *types.CommonEntityData {
    subscriptions.EntityData.YFilter = subscriptions.YFilter
    subscriptions.EntityData.YangName = "subscriptions"
    subscriptions.EntityData.BundleName = "cisco_ios_xr"
    subscriptions.EntityData.ParentYangName = "telemetry-model-driven"
    subscriptions.EntityData.SegmentPath = "subscriptions"
    subscriptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriptions.EntityData.Children = make(map[string]types.YChild)
    subscriptions.EntityData.Children["subscription"] = types.YChild{"Subscription", nil}
    for i := range subscriptions.Subscription {
        subscriptions.EntityData.Children[types.GetSegmentPath(&subscriptions.Subscription[i])] = types.YChild{"Subscription", &subscriptions.Subscription[i]}
    }
    subscriptions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(subscriptions.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription
// Telemetry Subscription
type TelemetryModelDriven_Subscriptions_Subscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Id of the subscription. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SubscriptionId interface{}

    // Subscription.
    Subscription TelemetryModelDriven_Subscriptions_Subscription_Subscription_

    // List of collection groups active for this subscription. The type is slice
    // of TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup.
    CollectionGroup []TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "cisco_ios_xr"
    subscription.EntityData.ParentYangName = "subscriptions"
    subscription.EntityData.SegmentPath = "subscription" + "[subscription-id='" + fmt.Sprintf("%v", subscription.SubscriptionId) + "']"
    subscription.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscription.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscription.EntityData.Children = make(map[string]types.YChild)
    subscription.EntityData.Children["subscription"] = types.YChild{"Subscription", &subscription.Subscription}
    subscription.EntityData.Children["collection-group"] = types.YChild{"CollectionGroup", nil}
    for i := range subscription.CollectionGroup {
        subscription.EntityData.Children[types.GetSegmentPath(&subscription.CollectionGroup[i])] = types.YChild{"CollectionGroup", &subscription.CollectionGroup[i]}
    }
    subscription.EntityData.Leafs = make(map[string]types.YLeaf)
    subscription.EntityData.Leafs["subscription-id"] = types.YLeaf{"SubscriptionId", subscription.SubscriptionId}
    return &(subscription.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_
// Subscription
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection Subscription name. The type is string.
    Id interface{}

    // Subscription state. The type is MdtSubsStateEnum.
    State interface{}

    // DSCP. The type is MdtSourceQosMarking.
    SourceQosMarking interface{}

    // configured source interface.
    SourceInterface TelemetryModelDriven_Subscriptions_Subscription_Subscription__SourceInterface

    // List of sensor groups within a subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile.
    SensorProfile []TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile

    // Array of destinations within a subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp.
    DestinationGrp []TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp
}

func (subscription_ *TelemetryModelDriven_Subscriptions_Subscription_Subscription_) GetEntityData() *types.CommonEntityData {
    subscription_.EntityData.YFilter = subscription_.YFilter
    subscription_.EntityData.YangName = "subscription"
    subscription_.EntityData.BundleName = "cisco_ios_xr"
    subscription_.EntityData.ParentYangName = "subscription"
    subscription_.EntityData.SegmentPath = "subscription"
    subscription_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscription_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscription_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscription_.EntityData.Children = make(map[string]types.YChild)
    subscription_.EntityData.Children["source-interface"] = types.YChild{"SourceInterface", &subscription_.SourceInterface}
    subscription_.EntityData.Children["sensor-profile"] = types.YChild{"SensorProfile", nil}
    for i := range subscription_.SensorProfile {
        subscription_.EntityData.Children[types.GetSegmentPath(&subscription_.SensorProfile[i])] = types.YChild{"SensorProfile", &subscription_.SensorProfile[i]}
    }
    subscription_.EntityData.Children["destination-grp"] = types.YChild{"DestinationGrp", nil}
    for i := range subscription_.DestinationGrp {
        subscription_.EntityData.Children[types.GetSegmentPath(&subscription_.DestinationGrp[i])] = types.YChild{"DestinationGrp", &subscription_.DestinationGrp[i]}
    }
    subscription_.EntityData.Leafs = make(map[string]types.YLeaf)
    subscription_.EntityData.Leafs["id"] = types.YLeaf{"Id", subscription_.Id}
    subscription_.EntityData.Leafs["state"] = types.YLeaf{"State", subscription_.State}
    subscription_.EntityData.Leafs["source-qos-marking"] = types.YLeaf{"SourceQosMarking", subscription_.SourceQosMarking}
    return &(subscription_.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription__SourceInterface
// configured source interface
type TelemetryModelDriven_Subscriptions_Subscription_Subscription__SourceInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Interface Name. The type is string.
    InterfaceName interface{}

    // interface state. The type is bool.
    State interface{}

    // IPV4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // Src Vrf Id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}
}

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription__SourceInterface) GetEntityData() *types.CommonEntityData {
    sourceInterface.EntityData.YFilter = sourceInterface.YFilter
    sourceInterface.EntityData.YangName = "source-interface"
    sourceInterface.EntityData.BundleName = "cisco_ios_xr"
    sourceInterface.EntityData.ParentYangName = "subscription"
    sourceInterface.EntityData.SegmentPath = "source-interface"
    sourceInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInterface.EntityData.Children = make(map[string]types.YChild)
    sourceInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    sourceInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", sourceInterface.InterfaceName}
    sourceInterface.EntityData.Leafs["state"] = types.YLeaf{"State", sourceInterface.State}
    sourceInterface.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", sourceInterface.Ipv4Address}
    sourceInterface.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", sourceInterface.Ipv6Address}
    sourceInterface.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", sourceInterface.VrfId}
    return &(sourceInterface.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile
// List of sensor groups within a subscription
type TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sample interval for the sensor group (ms). The type is interface{} with
    // range: 0..4294967295.
    SampleInterval interface{}

    // Heartbeat interval for the sensor group (s). The type is interface{} with
    // range: 0..4294967295.
    HeartbeatInterval interface{}

    // Suppress Redundant. The type is bool.
    SuppressRedundant interface{}

    // sensor group.
    SensorGroup TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile) GetEntityData() *types.CommonEntityData {
    sensorProfile.EntityData.YFilter = sensorProfile.YFilter
    sensorProfile.EntityData.YangName = "sensor-profile"
    sensorProfile.EntityData.BundleName = "cisco_ios_xr"
    sensorProfile.EntityData.ParentYangName = "subscription"
    sensorProfile.EntityData.SegmentPath = "sensor-profile"
    sensorProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorProfile.EntityData.Children = make(map[string]types.YChild)
    sensorProfile.EntityData.Children["sensor-group"] = types.YChild{"SensorGroup", &sensorProfile.SensorGroup}
    sensorProfile.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorProfile.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", sensorProfile.SampleInterval}
    sensorProfile.EntityData.Leafs["heartbeat-interval"] = types.YLeaf{"HeartbeatInterval", sensorProfile.HeartbeatInterval}
    sensorProfile.EntityData.Leafs["suppress-redundant"] = types.YLeaf{"SuppressRedundant", sensorProfile.SuppressRedundant}
    return &(sensorProfile.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup
// sensor group
type TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor Group name. The type is string.
    Id interface{}

    // Set if this is configured sensor group. The type is interface{} with range:
    // 0..4294967295.
    Configured interface{}

    // Array of information for sensor paths within sensor group. The type is
    // slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup_SensorPath.
    SensorPath []TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup_SensorPath
}

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup) GetEntityData() *types.CommonEntityData {
    sensorGroup.EntityData.YFilter = sensorGroup.YFilter
    sensorGroup.EntityData.YangName = "sensor-group"
    sensorGroup.EntityData.BundleName = "cisco_ios_xr"
    sensorGroup.EntityData.ParentYangName = "sensor-profile"
    sensorGroup.EntityData.SegmentPath = "sensor-group"
    sensorGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorGroup.EntityData.Children = make(map[string]types.YChild)
    sensorGroup.EntityData.Children["sensor-path"] = types.YChild{"SensorPath", nil}
    for i := range sensorGroup.SensorPath {
        sensorGroup.EntityData.Children[types.GetSegmentPath(&sensorGroup.SensorPath[i])] = types.YChild{"SensorPath", &sensorGroup.SensorPath[i]}
    }
    sensorGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorGroup.EntityData.Leafs["id"] = types.YLeaf{"Id", sensorGroup.Id}
    sensorGroup.EntityData.Leafs["configured"] = types.YLeaf{"Configured", sensorGroup.Configured}
    return &(sensorGroup.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup_SensorPath
// Array of information for sensor paths within
// sensor group
type TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup_SensorPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription__SensorProfile_SensorGroup_SensorPath) GetEntityData() *types.CommonEntityData {
    sensorPath.EntityData.YFilter = sensorPath.YFilter
    sensorPath.EntityData.YangName = "sensor-path"
    sensorPath.EntityData.BundleName = "cisco_ios_xr"
    sensorPath.EntityData.ParentYangName = "sensor-group"
    sensorPath.EntityData.SegmentPath = "sensor-path"
    sensorPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorPath.EntityData.Children = make(map[string]types.YChild)
    sensorPath.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorPath.EntityData.Leafs["path"] = types.YLeaf{"Path", sensorPath.Path}
    sensorPath.EntityData.Leafs["state"] = types.YLeaf{"State", sensorPath.State}
    sensorPath.EntityData.Leafs["status-str"] = types.YLeaf{"StatusStr", sensorPath.StatusStr}
    return &(sensorPath.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp
// Array of destinations within a subscription
type TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination Group name. The type is string.
    Id interface{}

    // Set if this is configured destination group. The type is interface{} with
    // range: 0..4294967295.
    Configured interface{}

    // list of destinations defined in this group. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination.
    Destination []TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination
}

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp) GetEntityData() *types.CommonEntityData {
    destinationGrp.EntityData.YFilter = destinationGrp.YFilter
    destinationGrp.EntityData.YangName = "destination-grp"
    destinationGrp.EntityData.BundleName = "cisco_ios_xr"
    destinationGrp.EntityData.ParentYangName = "subscription"
    destinationGrp.EntityData.SegmentPath = "destination-grp"
    destinationGrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationGrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationGrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationGrp.EntityData.Children = make(map[string]types.YChild)
    destinationGrp.EntityData.Children["destination"] = types.YChild{"Destination", nil}
    for i := range destinationGrp.Destination {
        destinationGrp.EntityData.Children[types.GetSegmentPath(&destinationGrp.Destination[i])] = types.YChild{"Destination", &destinationGrp.Destination[i]}
    }
    destinationGrp.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationGrp.EntityData.Leafs["id"] = types.YLeaf{"Id", destinationGrp.Id}
    destinationGrp.EntityData.Leafs["configured"] = types.YLeaf{"Configured", destinationGrp.Configured}
    return &(destinationGrp.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination
// list of destinations defined in this group
type TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination Id. The type is string.
    Id interface{}

    // Sub Idstr. The type is string.
    SubIdStr interface{}

    // Destination Port number. The type is interface{} with range: 0..65535.
    DestPort interface{}

    // Destination group encoding. The type is MdtEncodingEnum.
    Encoding interface{}

    // Destination group transport. The type is MdtTransportEnum.
    Transport interface{}

    // Destination group vrf. The type is string.
    Vrf interface{}

    // Destination group vrf id. The type is interface{} with range:
    // 0..4294967295.
    VrfId interface{}

    // State of streaming on this destination. The type is MdtDestStateEnum.
    State interface{}

    // UDP MTU if this destination is UDP. The type is interface{} with range:
    // 0..4294967295.
    UdpMtu interface{}

    // TLS connection to this destination. The type is interface{} with range:
    // 0..4294967295.
    Tls interface{}

    // TLS Hostname of this destination. The type is string.
    TlsHost interface{}

    // Total number of packets sent for this destination. The type is interface{}
    // with range: 0..18446744073709551615.
    TotalNumOfPacketsSent interface{}

    // Total number of bytes sent for this destination. The type is interface{}
    // with range: 0..18446744073709551615. Units are byte.
    TotalNumOfBytesSent interface{}

    // Timestamp of the last collection. The type is interface{} with range:
    // 0..18446744073709551615.
    LastCollectionTime interface{}

    // DSCP setting for this destination. The type is interface{} with range:
    // 0..4294967295.
    Dscp interface{}

    // Sub Id. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    SubId []interface{}

    // Destination IP Address.
    DestIpAddress TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination_DestIpAddress
}

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "destination-grp"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = make(map[string]types.YChild)
    destination.EntityData.Children["dest-ip-address"] = types.YChild{"DestIpAddress", &destination.DestIpAddress}
    destination.EntityData.Leafs = make(map[string]types.YLeaf)
    destination.EntityData.Leafs["id"] = types.YLeaf{"Id", destination.Id}
    destination.EntityData.Leafs["sub-id-str"] = types.YLeaf{"SubIdStr", destination.SubIdStr}
    destination.EntityData.Leafs["dest-port"] = types.YLeaf{"DestPort", destination.DestPort}
    destination.EntityData.Leafs["encoding"] = types.YLeaf{"Encoding", destination.Encoding}
    destination.EntityData.Leafs["transport"] = types.YLeaf{"Transport", destination.Transport}
    destination.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", destination.Vrf}
    destination.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", destination.VrfId}
    destination.EntityData.Leafs["state"] = types.YLeaf{"State", destination.State}
    destination.EntityData.Leafs["udp-mtu"] = types.YLeaf{"UdpMtu", destination.UdpMtu}
    destination.EntityData.Leafs["tls"] = types.YLeaf{"Tls", destination.Tls}
    destination.EntityData.Leafs["tls-host"] = types.YLeaf{"TlsHost", destination.TlsHost}
    destination.EntityData.Leafs["total-num-of-packets-sent"] = types.YLeaf{"TotalNumOfPacketsSent", destination.TotalNumOfPacketsSent}
    destination.EntityData.Leafs["total-num-of-bytes-sent"] = types.YLeaf{"TotalNumOfBytesSent", destination.TotalNumOfBytesSent}
    destination.EntityData.Leafs["last-collection-time"] = types.YLeaf{"LastCollectionTime", destination.LastCollectionTime}
    destination.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", destination.Dscp}
    destination.EntityData.Leafs["sub-id"] = types.YLeaf{"SubId", destination.SubId}
    return &(destination.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination_DestIpAddress
// Destination IP Address
type TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination_DestIpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPType. The type is MdtIp.
    IpType interface{}

    // IPV4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription__DestinationGrp_Destination_DestIpAddress) GetEntityData() *types.CommonEntityData {
    destIpAddress.EntityData.YFilter = destIpAddress.YFilter
    destIpAddress.EntityData.YangName = "dest-ip-address"
    destIpAddress.EntityData.BundleName = "cisco_ios_xr"
    destIpAddress.EntityData.ParentYangName = "destination"
    destIpAddress.EntityData.SegmentPath = "dest-ip-address"
    destIpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destIpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destIpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destIpAddress.EntityData.Children = make(map[string]types.YChild)
    destIpAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    destIpAddress.EntityData.Leafs["ip-type"] = types.YLeaf{"IpType", destIpAddress.IpType}
    destIpAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", destIpAddress.Ipv4Address}
    destIpAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", destIpAddress.Ipv6Address}
    return &(destIpAddress.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup
// List of collection groups active for this
// subscription
type TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection Group id. The type is interface{} with range:
    // 0..18446744073709551615.
    Id interface{}

    // Period of the collections (ms). The type is interface{} with range:
    // 0..4294967295.
    Cadence interface{}

    // Completed collections count. The type is interface{} with range:
    // 0..4294967295.
    TotalCollections interface{}

    // Destination group encoding. The type is MdtEncodingEnum.
    Encoding interface{}

    // Timestamp of the start of last collection. The type is interface{} with
    // range: 0..18446744073709551615.
    LastCollectionStartTime interface{}

    // Timestamp of the end of last collection. The type is interface{} with
    // range: 0..18446744073709551615.
    LastCollectionEndTime interface{}

    // Maximum time for a collection (ms). The type is interface{} with range:
    // 0..4294967295.
    MaxCollectionTime interface{}

    // Minimum time for a collection (ms). The type is interface{} with range:
    // 0..4294967295.
    MinCollectionTime interface{}

    // Minimum time for all processing (ms). The type is interface{} with range:
    // 0..4294967295.
    MinTotalTime interface{}

    // Maximum time for all processing (ms). The type is interface{} with range:
    // 0..4294967295.
    MaxTotalTime interface{}

    // Average time for all processing (ms). The type is interface{} with range:
    // 0..4294967295.
    AvgTotalTime interface{}

    // Total number of errors. The type is interface{} with range: 0..4294967295.
    TotalOtherErrors interface{}

    // Total number of no data instances. The type is interface{} with range:
    // 0..4294967295.
    TotalOnDataInstances interface{}

    // Total number skipped (not ready). The type is interface{} with range:
    // 0..4294967295.
    TotalNotReady interface{}

    // Total number of send errors. The type is interface{} with range:
    // 0..4294967295.
    TotalSendErrors interface{}

    // Total number of send drops. The type is interface{} with range:
    // 0..4294967295.
    TotalSendDrops interface{}

    // Array of information for sensor paths within collection group. The type is
    // slice of
    // TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath.
    CollectionPath []TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath

    // Array of information for sysdb paths within collection group. The type is
    // slice of
    // TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup.
    InternalCollectionGroup []TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup
}

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetEntityData() *types.CommonEntityData {
    collectionGroup.EntityData.YFilter = collectionGroup.YFilter
    collectionGroup.EntityData.YangName = "collection-group"
    collectionGroup.EntityData.BundleName = "cisco_ios_xr"
    collectionGroup.EntityData.ParentYangName = "subscription"
    collectionGroup.EntityData.SegmentPath = "collection-group"
    collectionGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectionGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectionGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectionGroup.EntityData.Children = make(map[string]types.YChild)
    collectionGroup.EntityData.Children["collection-path"] = types.YChild{"CollectionPath", nil}
    for i := range collectionGroup.CollectionPath {
        collectionGroup.EntityData.Children[types.GetSegmentPath(&collectionGroup.CollectionPath[i])] = types.YChild{"CollectionPath", &collectionGroup.CollectionPath[i]}
    }
    collectionGroup.EntityData.Children["internal-collection-group"] = types.YChild{"InternalCollectionGroup", nil}
    for i := range collectionGroup.InternalCollectionGroup {
        collectionGroup.EntityData.Children[types.GetSegmentPath(&collectionGroup.InternalCollectionGroup[i])] = types.YChild{"InternalCollectionGroup", &collectionGroup.InternalCollectionGroup[i]}
    }
    collectionGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    collectionGroup.EntityData.Leafs["id"] = types.YLeaf{"Id", collectionGroup.Id}
    collectionGroup.EntityData.Leafs["cadence"] = types.YLeaf{"Cadence", collectionGroup.Cadence}
    collectionGroup.EntityData.Leafs["total-collections"] = types.YLeaf{"TotalCollections", collectionGroup.TotalCollections}
    collectionGroup.EntityData.Leafs["encoding"] = types.YLeaf{"Encoding", collectionGroup.Encoding}
    collectionGroup.EntityData.Leafs["last-collection-start-time"] = types.YLeaf{"LastCollectionStartTime", collectionGroup.LastCollectionStartTime}
    collectionGroup.EntityData.Leafs["last-collection-end-time"] = types.YLeaf{"LastCollectionEndTime", collectionGroup.LastCollectionEndTime}
    collectionGroup.EntityData.Leafs["max-collection-time"] = types.YLeaf{"MaxCollectionTime", collectionGroup.MaxCollectionTime}
    collectionGroup.EntityData.Leafs["min-collection-time"] = types.YLeaf{"MinCollectionTime", collectionGroup.MinCollectionTime}
    collectionGroup.EntityData.Leafs["min-total-time"] = types.YLeaf{"MinTotalTime", collectionGroup.MinTotalTime}
    collectionGroup.EntityData.Leafs["max-total-time"] = types.YLeaf{"MaxTotalTime", collectionGroup.MaxTotalTime}
    collectionGroup.EntityData.Leafs["avg-total-time"] = types.YLeaf{"AvgTotalTime", collectionGroup.AvgTotalTime}
    collectionGroup.EntityData.Leafs["total-other-errors"] = types.YLeaf{"TotalOtherErrors", collectionGroup.TotalOtherErrors}
    collectionGroup.EntityData.Leafs["total-on-data-instances"] = types.YLeaf{"TotalOnDataInstances", collectionGroup.TotalOnDataInstances}
    collectionGroup.EntityData.Leafs["total-not-ready"] = types.YLeaf{"TotalNotReady", collectionGroup.TotalNotReady}
    collectionGroup.EntityData.Leafs["total-send-errors"] = types.YLeaf{"TotalSendErrors", collectionGroup.TotalSendErrors}
    collectionGroup.EntityData.Leafs["total-send-drops"] = types.YLeaf{"TotalSendDrops", collectionGroup.TotalSendDrops}
    return &(collectionGroup.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath
// Array of information for sensor paths within
// collection group
type TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetEntityData() *types.CommonEntityData {
    collectionPath.EntityData.YFilter = collectionPath.YFilter
    collectionPath.EntityData.YangName = "collection-path"
    collectionPath.EntityData.BundleName = "cisco_ios_xr"
    collectionPath.EntityData.ParentYangName = "collection-group"
    collectionPath.EntityData.SegmentPath = "collection-path"
    collectionPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectionPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectionPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectionPath.EntityData.Children = make(map[string]types.YChild)
    collectionPath.EntityData.Leafs = make(map[string]types.YLeaf)
    collectionPath.EntityData.Leafs["path"] = types.YLeaf{"Path", collectionPath.Path}
    collectionPath.EntityData.Leafs["state"] = types.YLeaf{"State", collectionPath.State}
    collectionPath.EntityData.Leafs["status-str"] = types.YLeaf{"StatusStr", collectionPath.StatusStr}
    return &(collectionPath.EntityData)
}

// TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup
// Array of information for sysdb paths within
// collection group
type TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sysdb Path. The type is string.
    Path interface{}

    // Period of the collections (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    Cadence interface{}

    // Total number of gets. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetCount interface{}

    // Total number of lists. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalListCount interface{}

    // Total number of datalists. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalDatalistCount interface{}

    // Total number of finddata. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalFinddataCount interface{}

    // Total number of get bulk. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetBulkCount interface{}

    // Total number of items retrived from sysdb. The type is interface{} with
    // range: 0..18446744073709551615.
    TotalItemCount interface{}

    // Total number of get errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetErrors interface{}

    // Total number of list errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalListErrors interface{}

    // Total number of datalist errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalDatalistErrors interface{}

    // Total number of finddata errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalFinddataErrors interface{}

    // Total number of get bulk errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGetBulkErrors interface{}

    // Total number of encode errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEncodeErrors interface{}

    // Total number of encode deferred. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEncodeNotready interface{}

    // Total number of send errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalSendErrors interface{}

    // Total number of send channel full. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalSendDrops interface{}

    // Total number of bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalSentBytes interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalSendPackets interface{}

    // Total number of send bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalSendBytesDropped interface{}

    // Completed collections count. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalCollections interface{}

    // Total number of collections missed. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalCollectionsMissed interface{}

    // Maximum time for a collection (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    MaxCollectionTime interface{}

    // Minimum time for a collection (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    MinCollectionTime interface{}

    // Average time for a collection (ms). The type is interface{} with range:
    // 0..18446744073709551615.
    AvgCollectionTime interface{}

    // Collection method in use. The type is interface{} with range:
    // 0..18446744073709551615.
    CollectionMethod interface{}

    // Status of collection path. The type is MdtInternalPathStatus.
    Status interface{}
}

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetEntityData() *types.CommonEntityData {
    internalCollectionGroup.EntityData.YFilter = internalCollectionGroup.YFilter
    internalCollectionGroup.EntityData.YangName = "internal-collection-group"
    internalCollectionGroup.EntityData.BundleName = "cisco_ios_xr"
    internalCollectionGroup.EntityData.ParentYangName = "collection-group"
    internalCollectionGroup.EntityData.SegmentPath = "internal-collection-group"
    internalCollectionGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalCollectionGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalCollectionGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalCollectionGroup.EntityData.Children = make(map[string]types.YChild)
    internalCollectionGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    internalCollectionGroup.EntityData.Leafs["path"] = types.YLeaf{"Path", internalCollectionGroup.Path}
    internalCollectionGroup.EntityData.Leafs["cadence"] = types.YLeaf{"Cadence", internalCollectionGroup.Cadence}
    internalCollectionGroup.EntityData.Leafs["total-get-count"] = types.YLeaf{"TotalGetCount", internalCollectionGroup.TotalGetCount}
    internalCollectionGroup.EntityData.Leafs["total-list-count"] = types.YLeaf{"TotalListCount", internalCollectionGroup.TotalListCount}
    internalCollectionGroup.EntityData.Leafs["total-datalist-count"] = types.YLeaf{"TotalDatalistCount", internalCollectionGroup.TotalDatalistCount}
    internalCollectionGroup.EntityData.Leafs["total-finddata-count"] = types.YLeaf{"TotalFinddataCount", internalCollectionGroup.TotalFinddataCount}
    internalCollectionGroup.EntityData.Leafs["total-get-bulk-count"] = types.YLeaf{"TotalGetBulkCount", internalCollectionGroup.TotalGetBulkCount}
    internalCollectionGroup.EntityData.Leafs["total-item-count"] = types.YLeaf{"TotalItemCount", internalCollectionGroup.TotalItemCount}
    internalCollectionGroup.EntityData.Leafs["total-get-errors"] = types.YLeaf{"TotalGetErrors", internalCollectionGroup.TotalGetErrors}
    internalCollectionGroup.EntityData.Leafs["total-list-errors"] = types.YLeaf{"TotalListErrors", internalCollectionGroup.TotalListErrors}
    internalCollectionGroup.EntityData.Leafs["total-datalist-errors"] = types.YLeaf{"TotalDatalistErrors", internalCollectionGroup.TotalDatalistErrors}
    internalCollectionGroup.EntityData.Leafs["total-finddata-errors"] = types.YLeaf{"TotalFinddataErrors", internalCollectionGroup.TotalFinddataErrors}
    internalCollectionGroup.EntityData.Leafs["total-get-bulk-errors"] = types.YLeaf{"TotalGetBulkErrors", internalCollectionGroup.TotalGetBulkErrors}
    internalCollectionGroup.EntityData.Leafs["total-encode-errors"] = types.YLeaf{"TotalEncodeErrors", internalCollectionGroup.TotalEncodeErrors}
    internalCollectionGroup.EntityData.Leafs["total-encode-notready"] = types.YLeaf{"TotalEncodeNotready", internalCollectionGroup.TotalEncodeNotready}
    internalCollectionGroup.EntityData.Leafs["total-send-errors"] = types.YLeaf{"TotalSendErrors", internalCollectionGroup.TotalSendErrors}
    internalCollectionGroup.EntityData.Leafs["total-send-drops"] = types.YLeaf{"TotalSendDrops", internalCollectionGroup.TotalSendDrops}
    internalCollectionGroup.EntityData.Leafs["total-sent-bytes"] = types.YLeaf{"TotalSentBytes", internalCollectionGroup.TotalSentBytes}
    internalCollectionGroup.EntityData.Leafs["total-send-packets"] = types.YLeaf{"TotalSendPackets", internalCollectionGroup.TotalSendPackets}
    internalCollectionGroup.EntityData.Leafs["total-send-bytes-dropped"] = types.YLeaf{"TotalSendBytesDropped", internalCollectionGroup.TotalSendBytesDropped}
    internalCollectionGroup.EntityData.Leafs["total-collections"] = types.YLeaf{"TotalCollections", internalCollectionGroup.TotalCollections}
    internalCollectionGroup.EntityData.Leafs["total-collections-missed"] = types.YLeaf{"TotalCollectionsMissed", internalCollectionGroup.TotalCollectionsMissed}
    internalCollectionGroup.EntityData.Leafs["max-collection-time"] = types.YLeaf{"MaxCollectionTime", internalCollectionGroup.MaxCollectionTime}
    internalCollectionGroup.EntityData.Leafs["min-collection-time"] = types.YLeaf{"MinCollectionTime", internalCollectionGroup.MinCollectionTime}
    internalCollectionGroup.EntityData.Leafs["avg-collection-time"] = types.YLeaf{"AvgCollectionTime", internalCollectionGroup.AvgCollectionTime}
    internalCollectionGroup.EntityData.Leafs["collection-method"] = types.YLeaf{"CollectionMethod", internalCollectionGroup.CollectionMethod}
    internalCollectionGroup.EntityData.Leafs["status"] = types.YLeaf{"Status", internalCollectionGroup.Status}
    return &(internalCollectionGroup.EntityData)
}

// TelemetryModelDriven_SensorGroups
// Telemetry Sensor Groups
type TelemetryModelDriven_SensorGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Telemetry Sensor Groups. The type is slice of
    // TelemetryModelDriven_SensorGroups_SensorGroup.
    SensorGroup []TelemetryModelDriven_SensorGroups_SensorGroup
}

func (sensorGroups *TelemetryModelDriven_SensorGroups) GetEntityData() *types.CommonEntityData {
    sensorGroups.EntityData.YFilter = sensorGroups.YFilter
    sensorGroups.EntityData.YangName = "sensor-groups"
    sensorGroups.EntityData.BundleName = "cisco_ios_xr"
    sensorGroups.EntityData.ParentYangName = "telemetry-model-driven"
    sensorGroups.EntityData.SegmentPath = "sensor-groups"
    sensorGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorGroups.EntityData.Children = make(map[string]types.YChild)
    sensorGroups.EntityData.Children["sensor-group"] = types.YChild{"SensorGroup", nil}
    for i := range sensorGroups.SensorGroup {
        sensorGroups.EntityData.Children[types.GetSegmentPath(&sensorGroups.SensorGroup[i])] = types.YChild{"SensorGroup", &sensorGroups.SensorGroup[i]}
    }
    sensorGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sensorGroups.EntityData)
}

// TelemetryModelDriven_SensorGroups_SensorGroup
// Telemetry Sensor Groups
type TelemetryModelDriven_SensorGroups_SensorGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Id of the sensor group. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SensorGroupId interface{}

    // Sensor Group name. The type is string.
    Id interface{}

    // Set if this is configured sensor group. The type is interface{} with range:
    // 0..4294967295.
    Configured interface{}

    // Array of information for sensor paths within sensor group. The type is
    // slice of TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath.
    SensorPath []TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetEntityData() *types.CommonEntityData {
    sensorGroup.EntityData.YFilter = sensorGroup.YFilter
    sensorGroup.EntityData.YangName = "sensor-group"
    sensorGroup.EntityData.BundleName = "cisco_ios_xr"
    sensorGroup.EntityData.ParentYangName = "sensor-groups"
    sensorGroup.EntityData.SegmentPath = "sensor-group" + "[sensor-group-id='" + fmt.Sprintf("%v", sensorGroup.SensorGroupId) + "']"
    sensorGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorGroup.EntityData.Children = make(map[string]types.YChild)
    sensorGroup.EntityData.Children["sensor-path"] = types.YChild{"SensorPath", nil}
    for i := range sensorGroup.SensorPath {
        sensorGroup.EntityData.Children[types.GetSegmentPath(&sensorGroup.SensorPath[i])] = types.YChild{"SensorPath", &sensorGroup.SensorPath[i]}
    }
    sensorGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorGroup.EntityData.Leafs["sensor-group-id"] = types.YLeaf{"SensorGroupId", sensorGroup.SensorGroupId}
    sensorGroup.EntityData.Leafs["id"] = types.YLeaf{"Id", sensorGroup.Id}
    sensorGroup.EntityData.Leafs["configured"] = types.YLeaf{"Configured", sensorGroup.Configured}
    return &(sensorGroup.EntityData)
}

// TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath
// Array of information for sensor paths within
// sensor group
type TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetEntityData() *types.CommonEntityData {
    sensorPath.EntityData.YFilter = sensorPath.YFilter
    sensorPath.EntityData.YangName = "sensor-path"
    sensorPath.EntityData.BundleName = "cisco_ios_xr"
    sensorPath.EntityData.ParentYangName = "sensor-group"
    sensorPath.EntityData.SegmentPath = "sensor-path"
    sensorPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorPath.EntityData.Children = make(map[string]types.YChild)
    sensorPath.EntityData.Leafs = make(map[string]types.YLeaf)
    sensorPath.EntityData.Leafs["path"] = types.YLeaf{"Path", sensorPath.Path}
    sensorPath.EntityData.Leafs["state"] = types.YLeaf{"State", sensorPath.State}
    sensorPath.EntityData.Leafs["status-str"] = types.YLeaf{"StatusStr", sensorPath.StatusStr}
    return &(sensorPath.EntityData)
}

