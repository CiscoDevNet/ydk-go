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

// MdtIp represents IP Type
type MdtIp string

const (
    // IPv4
    MdtIp_ipv4 MdtIp = "ipv4"

    // IPv6
    MdtIp_ipv6 MdtIp = "ipv6"
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

// TelemetryModelDriven
// Telemetry operational data
type TelemetryModelDriven struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Telemetry Destinations.
    Destinations TelemetryModelDriven_Destinations

    // Telemetry Subscriptions.
    Subscriptions TelemetryModelDriven_Subscriptions

    // Telemetry Sensor Groups.
    SensorGroups TelemetryModelDriven_SensorGroups
}

func (telemetryModelDriven *TelemetryModelDriven) GetFilter() yfilter.YFilter { return telemetryModelDriven.YFilter }

func (telemetryModelDriven *TelemetryModelDriven) SetFilter(yf yfilter.YFilter) { telemetryModelDriven.YFilter = yf }

func (telemetryModelDriven *TelemetryModelDriven) GetGoName(yname string) string {
    if yname == "destinations" { return "Destinations" }
    if yname == "subscriptions" { return "Subscriptions" }
    if yname == "sensor-groups" { return "SensorGroups" }
    return ""
}

func (telemetryModelDriven *TelemetryModelDriven) GetSegmentPath() string {
    return "Cisco-IOS-XR-telemetry-model-driven-oper:telemetry-model-driven"
}

func (telemetryModelDriven *TelemetryModelDriven) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destinations" {
        return &telemetryModelDriven.Destinations
    }
    if childYangName == "subscriptions" {
        return &telemetryModelDriven.Subscriptions
    }
    if childYangName == "sensor-groups" {
        return &telemetryModelDriven.SensorGroups
    }
    return nil
}

func (telemetryModelDriven *TelemetryModelDriven) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destinations"] = &telemetryModelDriven.Destinations
    children["subscriptions"] = &telemetryModelDriven.Subscriptions
    children["sensor-groups"] = &telemetryModelDriven.SensorGroups
    return children
}

func (telemetryModelDriven *TelemetryModelDriven) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
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

func (telemetryModelDriven *TelemetryModelDriven) GetParentYangName() string { return "Cisco-IOS-XR-telemetry-model-driven-oper" }

// TelemetryModelDriven_Destinations
// Telemetry Destinations
type TelemetryModelDriven_Destinations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Telemetry Destination. The type is slice of
    // TelemetryModelDriven_Destinations_Destination.
    Destination []TelemetryModelDriven_Destinations_Destination
}

func (destinations *TelemetryModelDriven_Destinations) GetFilter() yfilter.YFilter { return destinations.YFilter }

func (destinations *TelemetryModelDriven_Destinations) SetFilter(yf yfilter.YFilter) { destinations.YFilter = yf }

func (destinations *TelemetryModelDriven_Destinations) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    return ""
}

func (destinations *TelemetryModelDriven_Destinations) GetSegmentPath() string {
    return "destinations"
}

func (destinations *TelemetryModelDriven_Destinations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        for _, c := range destinations.Destination {
            if destinations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Destinations_Destination{}
        destinations.Destination = append(destinations.Destination, child)
        return &destinations.Destination[len(destinations.Destination)-1]
    }
    return nil
}

func (destinations *TelemetryModelDriven_Destinations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destinations.Destination {
        children[destinations.Destination[i].GetSegmentPath()] = &destinations.Destination[i]
    }
    return children
}

func (destinations *TelemetryModelDriven_Destinations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (destinations *TelemetryModelDriven_Destinations) GetBundleName() string { return "cisco_ios_xr" }

func (destinations *TelemetryModelDriven_Destinations) GetYangName() string { return "destinations" }

func (destinations *TelemetryModelDriven_Destinations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinations *TelemetryModelDriven_Destinations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinations *TelemetryModelDriven_Destinations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinations *TelemetryModelDriven_Destinations) SetParent(parent types.Entity) { destinations.parent = parent }

func (destinations *TelemetryModelDriven_Destinations) GetParent() types.Entity { return destinations.parent }

func (destinations *TelemetryModelDriven_Destinations) GetParentYangName() string { return "telemetry-model-driven" }

// TelemetryModelDriven_Destinations_Destination
// Telemetry Destination
type TelemetryModelDriven_Destinations_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Id of the destination. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    DestinationId interface{}

    // Destination Group name. The type is string.
    Id interface{}

    // Set if this is configured destination group. The type is interface{} with
    // range: 0..4294967295.
    Configured interface{}

    // list of destinations defined in this group. The type is slice of
    // TelemetryModelDriven_Destinations_Destination_Destination.
    Destination []TelemetryModelDriven_Destinations_Destination_Destination
}

func (destination *TelemetryModelDriven_Destinations_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *TelemetryModelDriven_Destinations_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *TelemetryModelDriven_Destinations_Destination) GetGoName(yname string) string {
    if yname == "destination-id" { return "DestinationId" }
    if yname == "id" { return "Id" }
    if yname == "configured" { return "Configured" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (destination *TelemetryModelDriven_Destinations_Destination) GetSegmentPath() string {
    return "destination" + "[destination-id='" + fmt.Sprintf("%v", destination.DestinationId) + "']"
}

func (destination *TelemetryModelDriven_Destinations_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        for _, c := range destination.Destination {
            if destination.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Destinations_Destination_Destination{}
        destination.Destination = append(destination.Destination, child)
        return &destination.Destination[len(destination.Destination)-1]
    }
    return nil
}

func (destination *TelemetryModelDriven_Destinations_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destination.Destination {
        children[destination.Destination[i].GetSegmentPath()] = &destination.Destination[i]
    }
    return children
}

func (destination *TelemetryModelDriven_Destinations_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-id"] = destination.DestinationId
    leafs["id"] = destination.Id
    leafs["configured"] = destination.Configured
    return leafs
}

func (destination *TelemetryModelDriven_Destinations_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *TelemetryModelDriven_Destinations_Destination) GetYangName() string { return "destination" }

func (destination *TelemetryModelDriven_Destinations_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *TelemetryModelDriven_Destinations_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *TelemetryModelDriven_Destinations_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *TelemetryModelDriven_Destinations_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *TelemetryModelDriven_Destinations_Destination) GetParent() types.Entity { return destination.parent }

func (destination *TelemetryModelDriven_Destinations_Destination) GetParentYangName() string { return "destinations" }

// TelemetryModelDriven_Destinations_Destination_Destination
// list of destinations defined in this group
type TelemetryModelDriven_Destinations_Destination_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination.
    Destination TelemetryModelDriven_Destinations_Destination_Destination_Destination

    // List of collection groups for this destination group. The type is slice of
    // TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup.
    CollectionGroup []TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "collection-group" { return "CollectionGroup" }
    return ""
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        return &destination.Destination
    }
    if childYangName == "collection-group" {
        for _, c := range destination.CollectionGroup {
            if destination.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup{}
        destination.CollectionGroup = append(destination.CollectionGroup, child)
        return &destination.CollectionGroup[len(destination.CollectionGroup)-1]
    }
    return nil
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination"] = &destination.Destination
    for i := range destination.CollectionGroup {
        children[destination.CollectionGroup[i].GetSegmentPath()] = &destination.CollectionGroup[i]
    }
    return children
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetYangName() string { return "destination" }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetParent() types.Entity { return destination.parent }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination) GetParentYangName() string { return "destination" }

// TelemetryModelDriven_Destinations_Destination_Destination_Destination
// Destination
type TelemetryModelDriven_Destinations_Destination_Destination_Destination struct {
    parent types.Entity
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
    DestIpAddress TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "sub-id-str" { return "SubIdStr" }
    if yname == "dest-port" { return "DestPort" }
    if yname == "encoding" { return "Encoding" }
    if yname == "transport" { return "Transport" }
    if yname == "vrf" { return "Vrf" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "state" { return "State" }
    if yname == "udp-mtu" { return "UdpMtu" }
    if yname == "tls" { return "Tls" }
    if yname == "tls-host" { return "TlsHost" }
    if yname == "total-num-of-packets-sent" { return "TotalNumOfPacketsSent" }
    if yname == "total-num-of-bytes-sent" { return "TotalNumOfBytesSent" }
    if yname == "last-collection-time" { return "LastCollectionTime" }
    if yname == "dscp" { return "Dscp" }
    if yname == "sub-id" { return "SubId" }
    if yname == "dest-ip-address" { return "DestIpAddress" }
    return ""
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dest-ip-address" {
        return &destination.DestIpAddress
    }
    return nil
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dest-ip-address"] = &destination.DestIpAddress
    return children
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = destination.Id
    leafs["sub-id-str"] = destination.SubIdStr
    leafs["dest-port"] = destination.DestPort
    leafs["encoding"] = destination.Encoding
    leafs["transport"] = destination.Transport
    leafs["vrf"] = destination.Vrf
    leafs["vrf-id"] = destination.VrfId
    leafs["state"] = destination.State
    leafs["udp-mtu"] = destination.UdpMtu
    leafs["tls"] = destination.Tls
    leafs["tls-host"] = destination.TlsHost
    leafs["total-num-of-packets-sent"] = destination.TotalNumOfPacketsSent
    leafs["total-num-of-bytes-sent"] = destination.TotalNumOfBytesSent
    leafs["last-collection-time"] = destination.LastCollectionTime
    leafs["dscp"] = destination.Dscp
    leafs["sub-id"] = destination.SubId
    return leafs
}

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetYangName() string { return "destination" }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetParent() types.Entity { return destination.parent }

func (destination *TelemetryModelDriven_Destinations_Destination_Destination_Destination) GetParentYangName() string { return "destination" }

// TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress
// Destination IP Address
type TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPType. The type is MdtIp.
    IpType interface{}

    // IPV4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetFilter() yfilter.YFilter { return destIpAddress.YFilter }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) SetFilter(yf yfilter.YFilter) { destIpAddress.YFilter = yf }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetGoName(yname string) string {
    if yname == "ip-type" { return "IpType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetSegmentPath() string {
    return "dest-ip-address"
}

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-type"] = destIpAddress.IpType
    leafs["ipv4-address"] = destIpAddress.Ipv4Address
    leafs["ipv6-address"] = destIpAddress.Ipv6Address
    return leafs
}

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetYangName() string { return "dest-ip-address" }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) SetParent(parent types.Entity) { destIpAddress.parent = parent }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetParent() types.Entity { return destIpAddress.parent }

func (destIpAddress *TelemetryModelDriven_Destinations_Destination_Destination_Destination_DestIpAddress) GetParentYangName() string { return "destination" }

// TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup
// List of collection groups for this destination
// group
type TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup struct {
    parent types.Entity
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
    // TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath.
    CollectionPath []TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath

    // Array of information for sysdb paths within collection group. The type is
    // slice of
    // TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup.
    InternalCollectionGroup []TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup
}

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetFilter() yfilter.YFilter { return collectionGroup.YFilter }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) SetFilter(yf yfilter.YFilter) { collectionGroup.YFilter = yf }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "cadence" { return "Cadence" }
    if yname == "total-collections" { return "TotalCollections" }
    if yname == "encoding" { return "Encoding" }
    if yname == "last-collection-start-time" { return "LastCollectionStartTime" }
    if yname == "last-collection-end-time" { return "LastCollectionEndTime" }
    if yname == "max-collection-time" { return "MaxCollectionTime" }
    if yname == "min-collection-time" { return "MinCollectionTime" }
    if yname == "min-total-time" { return "MinTotalTime" }
    if yname == "max-total-time" { return "MaxTotalTime" }
    if yname == "avg-total-time" { return "AvgTotalTime" }
    if yname == "total-other-errors" { return "TotalOtherErrors" }
    if yname == "total-on-data-instances" { return "TotalOnDataInstances" }
    if yname == "total-not-ready" { return "TotalNotReady" }
    if yname == "total-send-errors" { return "TotalSendErrors" }
    if yname == "total-send-drops" { return "TotalSendDrops" }
    if yname == "collection-path" { return "CollectionPath" }
    if yname == "internal-collection-group" { return "InternalCollectionGroup" }
    return ""
}

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetSegmentPath() string {
    return "collection-group"
}

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "collection-path" {
        for _, c := range collectionGroup.CollectionPath {
            if collectionGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath{}
        collectionGroup.CollectionPath = append(collectionGroup.CollectionPath, child)
        return &collectionGroup.CollectionPath[len(collectionGroup.CollectionPath)-1]
    }
    if childYangName == "internal-collection-group" {
        for _, c := range collectionGroup.InternalCollectionGroup {
            if collectionGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup{}
        collectionGroup.InternalCollectionGroup = append(collectionGroup.InternalCollectionGroup, child)
        return &collectionGroup.InternalCollectionGroup[len(collectionGroup.InternalCollectionGroup)-1]
    }
    return nil
}

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range collectionGroup.CollectionPath {
        children[collectionGroup.CollectionPath[i].GetSegmentPath()] = &collectionGroup.CollectionPath[i]
    }
    for i := range collectionGroup.InternalCollectionGroup {
        children[collectionGroup.InternalCollectionGroup[i].GetSegmentPath()] = &collectionGroup.InternalCollectionGroup[i]
    }
    return children
}

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = collectionGroup.Id
    leafs["cadence"] = collectionGroup.Cadence
    leafs["total-collections"] = collectionGroup.TotalCollections
    leafs["encoding"] = collectionGroup.Encoding
    leafs["last-collection-start-time"] = collectionGroup.LastCollectionStartTime
    leafs["last-collection-end-time"] = collectionGroup.LastCollectionEndTime
    leafs["max-collection-time"] = collectionGroup.MaxCollectionTime
    leafs["min-collection-time"] = collectionGroup.MinCollectionTime
    leafs["min-total-time"] = collectionGroup.MinTotalTime
    leafs["max-total-time"] = collectionGroup.MaxTotalTime
    leafs["avg-total-time"] = collectionGroup.AvgTotalTime
    leafs["total-other-errors"] = collectionGroup.TotalOtherErrors
    leafs["total-on-data-instances"] = collectionGroup.TotalOnDataInstances
    leafs["total-not-ready"] = collectionGroup.TotalNotReady
    leafs["total-send-errors"] = collectionGroup.TotalSendErrors
    leafs["total-send-drops"] = collectionGroup.TotalSendDrops
    return leafs
}

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetBundleName() string { return "cisco_ios_xr" }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetYangName() string { return "collection-group" }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) SetParent(parent types.Entity) { collectionGroup.parent = parent }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetParent() types.Entity { return collectionGroup.parent }

func (collectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup) GetParentYangName() string { return "destination" }

// TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath
// Array of information for sensor paths within
// collection group
type TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetFilter() yfilter.YFilter { return collectionPath.YFilter }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) SetFilter(yf yfilter.YFilter) { collectionPath.YFilter = yf }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "state" { return "State" }
    if yname == "status-str" { return "StatusStr" }
    return ""
}

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetSegmentPath() string {
    return "collection-path"
}

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = collectionPath.Path
    leafs["state"] = collectionPath.State
    leafs["status-str"] = collectionPath.StatusStr
    return leafs
}

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetBundleName() string { return "cisco_ios_xr" }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetYangName() string { return "collection-path" }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) SetParent(parent types.Entity) { collectionPath.parent = parent }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetParent() types.Entity { return collectionPath.parent }

func (collectionPath *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_CollectionPath) GetParentYangName() string { return "collection-group" }

// TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup
// Array of information for sysdb paths within
// collection group
type TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup struct {
    parent types.Entity
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

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetFilter() yfilter.YFilter { return internalCollectionGroup.YFilter }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) SetFilter(yf yfilter.YFilter) { internalCollectionGroup.YFilter = yf }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "cadence" { return "Cadence" }
    if yname == "total-get-count" { return "TotalGetCount" }
    if yname == "total-list-count" { return "TotalListCount" }
    if yname == "total-datalist-count" { return "TotalDatalistCount" }
    if yname == "total-finddata-count" { return "TotalFinddataCount" }
    if yname == "total-get-bulk-count" { return "TotalGetBulkCount" }
    if yname == "total-item-count" { return "TotalItemCount" }
    if yname == "total-get-errors" { return "TotalGetErrors" }
    if yname == "total-list-errors" { return "TotalListErrors" }
    if yname == "total-datalist-errors" { return "TotalDatalistErrors" }
    if yname == "total-finddata-errors" { return "TotalFinddataErrors" }
    if yname == "total-get-bulk-errors" { return "TotalGetBulkErrors" }
    if yname == "total-encode-errors" { return "TotalEncodeErrors" }
    if yname == "total-encode-notready" { return "TotalEncodeNotready" }
    if yname == "total-send-errors" { return "TotalSendErrors" }
    if yname == "total-send-drops" { return "TotalSendDrops" }
    if yname == "total-sent-bytes" { return "TotalSentBytes" }
    if yname == "total-send-packets" { return "TotalSendPackets" }
    if yname == "total-send-bytes-dropped" { return "TotalSendBytesDropped" }
    if yname == "total-collections" { return "TotalCollections" }
    if yname == "total-collections-missed" { return "TotalCollectionsMissed" }
    if yname == "max-collection-time" { return "MaxCollectionTime" }
    if yname == "min-collection-time" { return "MinCollectionTime" }
    if yname == "avg-collection-time" { return "AvgCollectionTime" }
    if yname == "collection-method" { return "CollectionMethod" }
    if yname == "status" { return "Status" }
    return ""
}

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetSegmentPath() string {
    return "internal-collection-group"
}

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = internalCollectionGroup.Path
    leafs["cadence"] = internalCollectionGroup.Cadence
    leafs["total-get-count"] = internalCollectionGroup.TotalGetCount
    leafs["total-list-count"] = internalCollectionGroup.TotalListCount
    leafs["total-datalist-count"] = internalCollectionGroup.TotalDatalistCount
    leafs["total-finddata-count"] = internalCollectionGroup.TotalFinddataCount
    leafs["total-get-bulk-count"] = internalCollectionGroup.TotalGetBulkCount
    leafs["total-item-count"] = internalCollectionGroup.TotalItemCount
    leafs["total-get-errors"] = internalCollectionGroup.TotalGetErrors
    leafs["total-list-errors"] = internalCollectionGroup.TotalListErrors
    leafs["total-datalist-errors"] = internalCollectionGroup.TotalDatalistErrors
    leafs["total-finddata-errors"] = internalCollectionGroup.TotalFinddataErrors
    leafs["total-get-bulk-errors"] = internalCollectionGroup.TotalGetBulkErrors
    leafs["total-encode-errors"] = internalCollectionGroup.TotalEncodeErrors
    leafs["total-encode-notready"] = internalCollectionGroup.TotalEncodeNotready
    leafs["total-send-errors"] = internalCollectionGroup.TotalSendErrors
    leafs["total-send-drops"] = internalCollectionGroup.TotalSendDrops
    leafs["total-sent-bytes"] = internalCollectionGroup.TotalSentBytes
    leafs["total-send-packets"] = internalCollectionGroup.TotalSendPackets
    leafs["total-send-bytes-dropped"] = internalCollectionGroup.TotalSendBytesDropped
    leafs["total-collections"] = internalCollectionGroup.TotalCollections
    leafs["total-collections-missed"] = internalCollectionGroup.TotalCollectionsMissed
    leafs["max-collection-time"] = internalCollectionGroup.MaxCollectionTime
    leafs["min-collection-time"] = internalCollectionGroup.MinCollectionTime
    leafs["avg-collection-time"] = internalCollectionGroup.AvgCollectionTime
    leafs["collection-method"] = internalCollectionGroup.CollectionMethod
    leafs["status"] = internalCollectionGroup.Status
    return leafs
}

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetBundleName() string { return "cisco_ios_xr" }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetYangName() string { return "internal-collection-group" }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) SetParent(parent types.Entity) { internalCollectionGroup.parent = parent }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetParent() types.Entity { return internalCollectionGroup.parent }

func (internalCollectionGroup *TelemetryModelDriven_Destinations_Destination_Destination_CollectionGroup_InternalCollectionGroup) GetParentYangName() string { return "collection-group" }

// TelemetryModelDriven_Subscriptions
// Telemetry Subscriptions
type TelemetryModelDriven_Subscriptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Telemetry Subscription. The type is slice of
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
// Telemetry Subscription
type TelemetryModelDriven_Subscriptions_Subscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Id of the subscription. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    SubscriptionId interface{}

    // Subscription.
    Subscription TelemetryModelDriven_Subscriptions_Subscription_Subscription

    // List of collection groups active for this subscription. The type is slice
    // of TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup.
    CollectionGroup []TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "subscription" { return "Subscription" }
    if yname == "collection-group" { return "CollectionGroup" }
    return ""
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetSegmentPath() string {
    return "subscription" + "[subscription-id='" + fmt.Sprintf("%v", subscription.SubscriptionId) + "']"
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscription" {
        return &subscription.Subscription
    }
    if childYangName == "collection-group" {
        for _, c := range subscription.CollectionGroup {
            if subscription.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup{}
        subscription.CollectionGroup = append(subscription.CollectionGroup, child)
        return &subscription.CollectionGroup[len(subscription.CollectionGroup)-1]
    }
    return nil
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["subscription"] = &subscription.Subscription
    for i := range subscription.CollectionGroup {
        children[subscription.CollectionGroup[i].GetSegmentPath()] = &subscription.CollectionGroup[i]
    }
    return children
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = subscription.SubscriptionId
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

// TelemetryModelDriven_Subscriptions_Subscription_Subscription
// Subscription
type TelemetryModelDriven_Subscriptions_Subscription_Subscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection Subscription name. The type is string.
    Id interface{}

    // Subscription state. The type is MdtSubsStateEnum.
    State interface{}

    // DSCP. The type is MdtSourceQosMarking.
    SourceQosMarking interface{}

    // configured source interface.
    SourceInterface TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface

    // List of sensor groups within a subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile.
    SensorProfile []TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile

    // Array of destinations within a subscription. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp.
    DestinationGrp []TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "state" { return "State" }
    if yname == "source-qos-marking" { return "SourceQosMarking" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "sensor-profile" { return "SensorProfile" }
    if yname == "destination-grp" { return "DestinationGrp" }
    return ""
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetSegmentPath() string {
    return "subscription"
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-interface" {
        return &subscription.SourceInterface
    }
    if childYangName == "sensor-profile" {
        for _, c := range subscription.SensorProfile {
            if subscription.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile{}
        subscription.SensorProfile = append(subscription.SensorProfile, child)
        return &subscription.SensorProfile[len(subscription.SensorProfile)-1]
    }
    if childYangName == "destination-grp" {
        for _, c := range subscription.DestinationGrp {
            if subscription.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp{}
        subscription.DestinationGrp = append(subscription.DestinationGrp, child)
        return &subscription.DestinationGrp[len(subscription.DestinationGrp)-1]
    }
    return nil
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-interface"] = &subscription.SourceInterface
    for i := range subscription.SensorProfile {
        children[subscription.SensorProfile[i].GetSegmentPath()] = &subscription.SensorProfile[i]
    }
    for i := range subscription.DestinationGrp {
        children[subscription.DestinationGrp[i].GetSegmentPath()] = &subscription.DestinationGrp[i]
    }
    return children
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = subscription.Id
    leafs["state"] = subscription.State
    leafs["source-qos-marking"] = subscription.SourceQosMarking
    return leafs
}

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetBundleName() string { return "cisco_ios_xr" }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetYangName() string { return "subscription" }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *TelemetryModelDriven_Subscriptions_Subscription_Subscription) GetParentYangName() string { return "subscription" }

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface
// configured source interface
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source Interface Name. The type is string.
    InterfaceName interface{}

    // interface state. The type is bool.
    State interface{}

    // IPV4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Src Vrf Id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}
}

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetFilter() yfilter.YFilter { return sourceInterface.YFilter }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) SetFilter(yf yfilter.YFilter) { sourceInterface.YFilter = yf }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "state" { return "State" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-id" { return "VrfId" }
    return ""
}

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetSegmentPath() string {
    return "source-interface"
}

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = sourceInterface.InterfaceName
    leafs["state"] = sourceInterface.State
    leafs["ipv4-address"] = sourceInterface.Ipv4Address
    leafs["ipv6-address"] = sourceInterface.Ipv6Address
    leafs["vrf-id"] = sourceInterface.VrfId
    return leafs
}

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetBundleName() string { return "cisco_ios_xr" }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetYangName() string { return "source-interface" }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) SetParent(parent types.Entity) { sourceInterface.parent = parent }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetParent() types.Entity { return sourceInterface.parent }

func (sourceInterface *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SourceInterface) GetParentYangName() string { return "subscription" }

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile
// List of sensor groups within a subscription
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile struct {
    parent types.Entity
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
    SensorGroup TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetFilter() yfilter.YFilter { return sensorProfile.YFilter }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) SetFilter(yf yfilter.YFilter) { sensorProfile.YFilter = yf }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetGoName(yname string) string {
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "heartbeat-interval" { return "HeartbeatInterval" }
    if yname == "suppress-redundant" { return "SuppressRedundant" }
    if yname == "sensor-group" { return "SensorGroup" }
    return ""
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetSegmentPath() string {
    return "sensor-profile"
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-group" {
        return &sensorProfile.SensorGroup
    }
    return nil
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensor-group"] = &sensorProfile.SensorGroup
    return children
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-interval"] = sensorProfile.SampleInterval
    leafs["heartbeat-interval"] = sensorProfile.HeartbeatInterval
    leafs["suppress-redundant"] = sensorProfile.SuppressRedundant
    return leafs
}

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetBundleName() string { return "cisco_ios_xr" }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetYangName() string { return "sensor-profile" }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) SetParent(parent types.Entity) { sensorProfile.parent = parent }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetParent() types.Entity { return sensorProfile.parent }

func (sensorProfile *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile) GetParentYangName() string { return "subscription" }

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup
// sensor group
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor Group name. The type is string.
    Id interface{}

    // Set if this is configured sensor group. The type is interface{} with range:
    // 0..4294967295.
    Configured interface{}

    // Array of information for sensor paths within sensor group. The type is
    // slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath.
    SensorPath []TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath
}

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetFilter() yfilter.YFilter { return sensorGroup.YFilter }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) SetFilter(yf yfilter.YFilter) { sensorGroup.YFilter = yf }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "configured" { return "Configured" }
    if yname == "sensor-path" { return "SensorPath" }
    return ""
}

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetSegmentPath() string {
    return "sensor-group"
}

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-path" {
        for _, c := range sensorGroup.SensorPath {
            if sensorGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath{}
        sensorGroup.SensorPath = append(sensorGroup.SensorPath, child)
        return &sensorGroup.SensorPath[len(sensorGroup.SensorPath)-1]
    }
    return nil
}

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorGroup.SensorPath {
        children[sensorGroup.SensorPath[i].GetSegmentPath()] = &sensorGroup.SensorPath[i]
    }
    return children
}

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = sensorGroup.Id
    leafs["configured"] = sensorGroup.Configured
    return leafs
}

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetBundleName() string { return "cisco_ios_xr" }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetYangName() string { return "sensor-group" }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) SetParent(parent types.Entity) { sensorGroup.parent = parent }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetParent() types.Entity { return sensorGroup.parent }

func (sensorGroup *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup) GetParentYangName() string { return "sensor-profile" }

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath
// Array of information for sensor paths within
// sensor group
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetFilter() yfilter.YFilter { return sensorPath.YFilter }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) SetFilter(yf yfilter.YFilter) { sensorPath.YFilter = yf }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "state" { return "State" }
    if yname == "status-str" { return "StatusStr" }
    return ""
}

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetSegmentPath() string {
    return "sensor-path"
}

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = sensorPath.Path
    leafs["state"] = sensorPath.State
    leafs["status-str"] = sensorPath.StatusStr
    return leafs
}

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetBundleName() string { return "cisco_ios_xr" }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetYangName() string { return "sensor-path" }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) SetParent(parent types.Entity) { sensorPath.parent = parent }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetParent() types.Entity { return sensorPath.parent }

func (sensorPath *TelemetryModelDriven_Subscriptions_Subscription_Subscription_SensorProfile_SensorGroup_SensorPath) GetParentYangName() string { return "sensor-group" }

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp
// Array of destinations within a subscription
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination Group name. The type is string.
    Id interface{}

    // Set if this is configured destination group. The type is interface{} with
    // range: 0..4294967295.
    Configured interface{}

    // list of destinations defined in this group. The type is slice of
    // TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination.
    Destination []TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination
}

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetFilter() yfilter.YFilter { return destinationGrp.YFilter }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) SetFilter(yf yfilter.YFilter) { destinationGrp.YFilter = yf }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "configured" { return "Configured" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetSegmentPath() string {
    return "destination-grp"
}

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        for _, c := range destinationGrp.Destination {
            if destinationGrp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination{}
        destinationGrp.Destination = append(destinationGrp.Destination, child)
        return &destinationGrp.Destination[len(destinationGrp.Destination)-1]
    }
    return nil
}

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range destinationGrp.Destination {
        children[destinationGrp.Destination[i].GetSegmentPath()] = &destinationGrp.Destination[i]
    }
    return children
}

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = destinationGrp.Id
    leafs["configured"] = destinationGrp.Configured
    return leafs
}

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetBundleName() string { return "cisco_ios_xr" }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetYangName() string { return "destination-grp" }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) SetParent(parent types.Entity) { destinationGrp.parent = parent }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetParent() types.Entity { return destinationGrp.parent }

func (destinationGrp *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp) GetParentYangName() string { return "subscription" }

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination
// list of destinations defined in this group
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination struct {
    parent types.Entity
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
    DestIpAddress TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress
}

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "sub-id-str" { return "SubIdStr" }
    if yname == "dest-port" { return "DestPort" }
    if yname == "encoding" { return "Encoding" }
    if yname == "transport" { return "Transport" }
    if yname == "vrf" { return "Vrf" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "state" { return "State" }
    if yname == "udp-mtu" { return "UdpMtu" }
    if yname == "tls" { return "Tls" }
    if yname == "tls-host" { return "TlsHost" }
    if yname == "total-num-of-packets-sent" { return "TotalNumOfPacketsSent" }
    if yname == "total-num-of-bytes-sent" { return "TotalNumOfBytesSent" }
    if yname == "last-collection-time" { return "LastCollectionTime" }
    if yname == "dscp" { return "Dscp" }
    if yname == "sub-id" { return "SubId" }
    if yname == "dest-ip-address" { return "DestIpAddress" }
    return ""
}

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dest-ip-address" {
        return &destination.DestIpAddress
    }
    return nil
}

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dest-ip-address"] = &destination.DestIpAddress
    return children
}

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = destination.Id
    leafs["sub-id-str"] = destination.SubIdStr
    leafs["dest-port"] = destination.DestPort
    leafs["encoding"] = destination.Encoding
    leafs["transport"] = destination.Transport
    leafs["vrf"] = destination.Vrf
    leafs["vrf-id"] = destination.VrfId
    leafs["state"] = destination.State
    leafs["udp-mtu"] = destination.UdpMtu
    leafs["tls"] = destination.Tls
    leafs["tls-host"] = destination.TlsHost
    leafs["total-num-of-packets-sent"] = destination.TotalNumOfPacketsSent
    leafs["total-num-of-bytes-sent"] = destination.TotalNumOfBytesSent
    leafs["last-collection-time"] = destination.LastCollectionTime
    leafs["dscp"] = destination.Dscp
    leafs["sub-id"] = destination.SubId
    return leafs
}

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetYangName() string { return "destination" }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetParent() types.Entity { return destination.parent }

func (destination *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination) GetParentYangName() string { return "destination-grp" }

// TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress
// Destination IP Address
type TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPType. The type is MdtIp.
    IpType interface{}

    // IPV4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetFilter() yfilter.YFilter { return destIpAddress.YFilter }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) SetFilter(yf yfilter.YFilter) { destIpAddress.YFilter = yf }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetGoName(yname string) string {
    if yname == "ip-type" { return "IpType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetSegmentPath() string {
    return "dest-ip-address"
}

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-type"] = destIpAddress.IpType
    leafs["ipv4-address"] = destIpAddress.Ipv4Address
    leafs["ipv6-address"] = destIpAddress.Ipv6Address
    return leafs
}

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetYangName() string { return "dest-ip-address" }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) SetParent(parent types.Entity) { destIpAddress.parent = parent }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetParent() types.Entity { return destIpAddress.parent }

func (destIpAddress *TelemetryModelDriven_Subscriptions_Subscription_Subscription_DestinationGrp_Destination_DestIpAddress) GetParentYangName() string { return "destination" }

// TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup
// List of collection groups active for this
// subscription
type TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup struct {
    parent types.Entity
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

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetFilter() yfilter.YFilter { return collectionGroup.YFilter }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) SetFilter(yf yfilter.YFilter) { collectionGroup.YFilter = yf }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "cadence" { return "Cadence" }
    if yname == "total-collections" { return "TotalCollections" }
    if yname == "encoding" { return "Encoding" }
    if yname == "last-collection-start-time" { return "LastCollectionStartTime" }
    if yname == "last-collection-end-time" { return "LastCollectionEndTime" }
    if yname == "max-collection-time" { return "MaxCollectionTime" }
    if yname == "min-collection-time" { return "MinCollectionTime" }
    if yname == "min-total-time" { return "MinTotalTime" }
    if yname == "max-total-time" { return "MaxTotalTime" }
    if yname == "avg-total-time" { return "AvgTotalTime" }
    if yname == "total-other-errors" { return "TotalOtherErrors" }
    if yname == "total-on-data-instances" { return "TotalOnDataInstances" }
    if yname == "total-not-ready" { return "TotalNotReady" }
    if yname == "total-send-errors" { return "TotalSendErrors" }
    if yname == "total-send-drops" { return "TotalSendDrops" }
    if yname == "collection-path" { return "CollectionPath" }
    if yname == "internal-collection-group" { return "InternalCollectionGroup" }
    return ""
}

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetSegmentPath() string {
    return "collection-group"
}

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "collection-path" {
        for _, c := range collectionGroup.CollectionPath {
            if collectionGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath{}
        collectionGroup.CollectionPath = append(collectionGroup.CollectionPath, child)
        return &collectionGroup.CollectionPath[len(collectionGroup.CollectionPath)-1]
    }
    if childYangName == "internal-collection-group" {
        for _, c := range collectionGroup.InternalCollectionGroup {
            if collectionGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup{}
        collectionGroup.InternalCollectionGroup = append(collectionGroup.InternalCollectionGroup, child)
        return &collectionGroup.InternalCollectionGroup[len(collectionGroup.InternalCollectionGroup)-1]
    }
    return nil
}

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range collectionGroup.CollectionPath {
        children[collectionGroup.CollectionPath[i].GetSegmentPath()] = &collectionGroup.CollectionPath[i]
    }
    for i := range collectionGroup.InternalCollectionGroup {
        children[collectionGroup.InternalCollectionGroup[i].GetSegmentPath()] = &collectionGroup.InternalCollectionGroup[i]
    }
    return children
}

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = collectionGroup.Id
    leafs["cadence"] = collectionGroup.Cadence
    leafs["total-collections"] = collectionGroup.TotalCollections
    leafs["encoding"] = collectionGroup.Encoding
    leafs["last-collection-start-time"] = collectionGroup.LastCollectionStartTime
    leafs["last-collection-end-time"] = collectionGroup.LastCollectionEndTime
    leafs["max-collection-time"] = collectionGroup.MaxCollectionTime
    leafs["min-collection-time"] = collectionGroup.MinCollectionTime
    leafs["min-total-time"] = collectionGroup.MinTotalTime
    leafs["max-total-time"] = collectionGroup.MaxTotalTime
    leafs["avg-total-time"] = collectionGroup.AvgTotalTime
    leafs["total-other-errors"] = collectionGroup.TotalOtherErrors
    leafs["total-on-data-instances"] = collectionGroup.TotalOnDataInstances
    leafs["total-not-ready"] = collectionGroup.TotalNotReady
    leafs["total-send-errors"] = collectionGroup.TotalSendErrors
    leafs["total-send-drops"] = collectionGroup.TotalSendDrops
    return leafs
}

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetBundleName() string { return "cisco_ios_xr" }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetYangName() string { return "collection-group" }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) SetParent(parent types.Entity) { collectionGroup.parent = parent }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetParent() types.Entity { return collectionGroup.parent }

func (collectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup) GetParentYangName() string { return "subscription" }

// TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath
// Array of information for sensor paths within
// collection group
type TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetFilter() yfilter.YFilter { return collectionPath.YFilter }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) SetFilter(yf yfilter.YFilter) { collectionPath.YFilter = yf }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "state" { return "State" }
    if yname == "status-str" { return "StatusStr" }
    return ""
}

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetSegmentPath() string {
    return "collection-path"
}

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = collectionPath.Path
    leafs["state"] = collectionPath.State
    leafs["status-str"] = collectionPath.StatusStr
    return leafs
}

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetBundleName() string { return "cisco_ios_xr" }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetYangName() string { return "collection-path" }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) SetParent(parent types.Entity) { collectionPath.parent = parent }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetParent() types.Entity { return collectionPath.parent }

func (collectionPath *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_CollectionPath) GetParentYangName() string { return "collection-group" }

// TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup
// Array of information for sysdb paths within
// collection group
type TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup struct {
    parent types.Entity
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

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetFilter() yfilter.YFilter { return internalCollectionGroup.YFilter }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) SetFilter(yf yfilter.YFilter) { internalCollectionGroup.YFilter = yf }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "cadence" { return "Cadence" }
    if yname == "total-get-count" { return "TotalGetCount" }
    if yname == "total-list-count" { return "TotalListCount" }
    if yname == "total-datalist-count" { return "TotalDatalistCount" }
    if yname == "total-finddata-count" { return "TotalFinddataCount" }
    if yname == "total-get-bulk-count" { return "TotalGetBulkCount" }
    if yname == "total-item-count" { return "TotalItemCount" }
    if yname == "total-get-errors" { return "TotalGetErrors" }
    if yname == "total-list-errors" { return "TotalListErrors" }
    if yname == "total-datalist-errors" { return "TotalDatalistErrors" }
    if yname == "total-finddata-errors" { return "TotalFinddataErrors" }
    if yname == "total-get-bulk-errors" { return "TotalGetBulkErrors" }
    if yname == "total-encode-errors" { return "TotalEncodeErrors" }
    if yname == "total-encode-notready" { return "TotalEncodeNotready" }
    if yname == "total-send-errors" { return "TotalSendErrors" }
    if yname == "total-send-drops" { return "TotalSendDrops" }
    if yname == "total-sent-bytes" { return "TotalSentBytes" }
    if yname == "total-send-packets" { return "TotalSendPackets" }
    if yname == "total-send-bytes-dropped" { return "TotalSendBytesDropped" }
    if yname == "total-collections" { return "TotalCollections" }
    if yname == "total-collections-missed" { return "TotalCollectionsMissed" }
    if yname == "max-collection-time" { return "MaxCollectionTime" }
    if yname == "min-collection-time" { return "MinCollectionTime" }
    if yname == "avg-collection-time" { return "AvgCollectionTime" }
    if yname == "collection-method" { return "CollectionMethod" }
    if yname == "status" { return "Status" }
    return ""
}

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetSegmentPath() string {
    return "internal-collection-group"
}

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = internalCollectionGroup.Path
    leafs["cadence"] = internalCollectionGroup.Cadence
    leafs["total-get-count"] = internalCollectionGroup.TotalGetCount
    leafs["total-list-count"] = internalCollectionGroup.TotalListCount
    leafs["total-datalist-count"] = internalCollectionGroup.TotalDatalistCount
    leafs["total-finddata-count"] = internalCollectionGroup.TotalFinddataCount
    leafs["total-get-bulk-count"] = internalCollectionGroup.TotalGetBulkCount
    leafs["total-item-count"] = internalCollectionGroup.TotalItemCount
    leafs["total-get-errors"] = internalCollectionGroup.TotalGetErrors
    leafs["total-list-errors"] = internalCollectionGroup.TotalListErrors
    leafs["total-datalist-errors"] = internalCollectionGroup.TotalDatalistErrors
    leafs["total-finddata-errors"] = internalCollectionGroup.TotalFinddataErrors
    leafs["total-get-bulk-errors"] = internalCollectionGroup.TotalGetBulkErrors
    leafs["total-encode-errors"] = internalCollectionGroup.TotalEncodeErrors
    leafs["total-encode-notready"] = internalCollectionGroup.TotalEncodeNotready
    leafs["total-send-errors"] = internalCollectionGroup.TotalSendErrors
    leafs["total-send-drops"] = internalCollectionGroup.TotalSendDrops
    leafs["total-sent-bytes"] = internalCollectionGroup.TotalSentBytes
    leafs["total-send-packets"] = internalCollectionGroup.TotalSendPackets
    leafs["total-send-bytes-dropped"] = internalCollectionGroup.TotalSendBytesDropped
    leafs["total-collections"] = internalCollectionGroup.TotalCollections
    leafs["total-collections-missed"] = internalCollectionGroup.TotalCollectionsMissed
    leafs["max-collection-time"] = internalCollectionGroup.MaxCollectionTime
    leafs["min-collection-time"] = internalCollectionGroup.MinCollectionTime
    leafs["avg-collection-time"] = internalCollectionGroup.AvgCollectionTime
    leafs["collection-method"] = internalCollectionGroup.CollectionMethod
    leafs["status"] = internalCollectionGroup.Status
    return leafs
}

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetBundleName() string { return "cisco_ios_xr" }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetYangName() string { return "internal-collection-group" }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) SetParent(parent types.Entity) { internalCollectionGroup.parent = parent }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetParent() types.Entity { return internalCollectionGroup.parent }

func (internalCollectionGroup *TelemetryModelDriven_Subscriptions_Subscription_CollectionGroup_InternalCollectionGroup) GetParentYangName() string { return "collection-group" }

// TelemetryModelDriven_SensorGroups
// Telemetry Sensor Groups
type TelemetryModelDriven_SensorGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Telemetry Sensor Groups. The type is slice of
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
// Telemetry Sensor Groups
type TelemetryModelDriven_SensorGroups_SensorGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Id of the sensor group. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetFilter() yfilter.YFilter { return sensorGroup.YFilter }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) SetFilter(yf yfilter.YFilter) { sensorGroup.YFilter = yf }

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetGoName(yname string) string {
    if yname == "sensor-group-id" { return "SensorGroupId" }
    if yname == "id" { return "Id" }
    if yname == "configured" { return "Configured" }
    if yname == "sensor-path" { return "SensorPath" }
    return ""
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetSegmentPath() string {
    return "sensor-group" + "[sensor-group-id='" + fmt.Sprintf("%v", sensorGroup.SensorGroupId) + "']"
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor-path" {
        for _, c := range sensorGroup.SensorPath {
            if sensorGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath{}
        sensorGroup.SensorPath = append(sensorGroup.SensorPath, child)
        return &sensorGroup.SensorPath[len(sensorGroup.SensorPath)-1]
    }
    return nil
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensorGroup.SensorPath {
        children[sensorGroup.SensorPath[i].GetSegmentPath()] = &sensorGroup.SensorPath[i]
    }
    return children
}

func (sensorGroup *TelemetryModelDriven_SensorGroups_SensorGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-group-id"] = sensorGroup.SensorGroupId
    leafs["id"] = sensorGroup.Id
    leafs["configured"] = sensorGroup.Configured
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

// TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath
// Array of information for sensor paths within
// sensor group
type TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor Path. The type is string.
    Path interface{}

    // State, if sensor path is resolved or not. The type is bool.
    State interface{}

    // Error str, if there are any errors resolving the sensor path. The type is
    // string.
    StatusStr interface{}
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetFilter() yfilter.YFilter { return sensorPath.YFilter }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) SetFilter(yf yfilter.YFilter) { sensorPath.YFilter = yf }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "state" { return "State" }
    if yname == "status-str" { return "StatusStr" }
    return ""
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetSegmentPath() string {
    return "sensor-path"
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = sensorPath.Path
    leafs["state"] = sensorPath.State
    leafs["status-str"] = sensorPath.StatusStr
    return leafs
}

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetBundleName() string { return "cisco_ios_xr" }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetYangName() string { return "sensor-path" }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) SetParent(parent types.Entity) { sensorPath.parent = parent }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetParent() types.Entity { return sensorPath.parent }

func (sensorPath *TelemetryModelDriven_SensorGroups_SensorGroup_SensorPath) GetParentYangName() string { return "sensor-group" }

