// This module contains a collection of YANG 
// definitions for operational data of streaming telemetry.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package mdt_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mdt_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mdt-oper mdt-oper-data}", reflect.TypeOf(MdtOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mdt-oper:mdt-oper-data", reflect.TypeOf(MdtOperData{}))
}

// MdtSubState represents Subscription states
type MdtSubState string

const (
    // The subscription is valid and may be sending updates.
    MdtSubState_sub_state_valid MdtSubState = "sub-state-valid"

    // The subscription has been suspended and is not
    // sending notifications even if there are updates.
    MdtSubState_sub_state_suspended MdtSubState = "sub-state-suspended"

    // The subscription is terminated. This state is valid
    // only for static subscriptions.
    MdtSubState_sub_state_terminated MdtSubState = "sub-state-terminated"

    // The subscription is invalid. This state is valid
    // only for static subscriptions.
    MdtSubState_sub_state_invalid MdtSubState = "sub-state-invalid"
)

// MdtConState represents Connection states.
type MdtConState string

const (
    // The connection is active and usable.
    MdtConState_con_state_active MdtConState = "con-state-active"

    // An attempt is being made to set the connection up.
    MdtConState_con_state_connecting MdtConState = "con-state-connecting"

    // The connection is down, but between connection
    // attempts. It is in this state, for example, during
    // the idle time between retries.
    MdtConState_con_state_pending MdtConState = "con-state-pending"

    // The connection is the process of being disconnected.
    MdtConState_con_state_disconnecting MdtConState = "con-state-disconnecting"
)

// MdtSubType represents Subscription types
type MdtSubType string

const (
    // Dynamic subscriptions
    // -do not survive reboot
    // -existence tied to connection they are created on
    // -send updates only to peer that creates them
    MdtSubType_sub_type_dynamic MdtSubType = "sub-type-dynamic"

    // Static subscriptions
    // -created, (modified), and deleted by management operations
    // -survive reboot
    // -receivers are configured 
    MdtSubType_sub_type_static MdtSubType = "sub-type-static"
)

// MdtReceiverState represents Receiver states.
type MdtReceiverState string

const (
    // The receiver configuration is invalid and
    // cannot be used.
    MdtReceiverState_rcvr_state_invalid MdtReceiverState = "rcvr-state-invalid"

    // The receiver is disconnected and there is no
    // attempt being made to connect to it.
    MdtReceiverState_rcvr_state_disconnected MdtReceiverState = "rcvr-state-disconnected"

    // An attempt is being made to connect to the receiver.
    MdtReceiverState_rcvr_state_connecting MdtReceiverState = "rcvr-state-connecting"

    // The receiver is connected, and update notifications
    // are being sent to the receiver when they occur
    MdtReceiverState_rcvr_state_connected MdtReceiverState = "rcvr-state-connected"
)

// MdtOperData
// MDT operational data.
type MdtOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MDT streams table. The list of supported streams.
    MdtStreams MdtOperData_MdtStreams

    // MDT subscription operational data. The type is slice of
    // MdtOperData_MdtSubscriptions.
    MdtSubscriptions []*MdtOperData_MdtSubscriptions

    // MDT subscription connection operational data. The type is slice of
    // MdtOperData_MdtConnections.
    MdtConnections []*MdtOperData_MdtConnections
}

func (mdtOperData *MdtOperData) GetEntityData() *types.CommonEntityData {
    mdtOperData.EntityData.YFilter = mdtOperData.YFilter
    mdtOperData.EntityData.YangName = "mdt-oper-data"
    mdtOperData.EntityData.BundleName = "cisco_ios_xe"
    mdtOperData.EntityData.ParentYangName = "Cisco-IOS-XE-mdt-oper"
    mdtOperData.EntityData.SegmentPath = "Cisco-IOS-XE-mdt-oper:mdt-oper-data"
    mdtOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtOperData.EntityData.Children = types.NewOrderedMap()
    mdtOperData.EntityData.Children.Append("mdt-streams", types.YChild{"MdtStreams", &mdtOperData.MdtStreams})
    mdtOperData.EntityData.Children.Append("mdt-subscriptions", types.YChild{"MdtSubscriptions", nil})
    for i := range mdtOperData.MdtSubscriptions {
        mdtOperData.EntityData.Children.Append(types.GetSegmentPath(mdtOperData.MdtSubscriptions[i]), types.YChild{"MdtSubscriptions", mdtOperData.MdtSubscriptions[i]})
    }
    mdtOperData.EntityData.Children.Append("mdt-connections", types.YChild{"MdtConnections", nil})
    for i := range mdtOperData.MdtConnections {
        mdtOperData.EntityData.Children.Append(types.GetSegmentPath(mdtOperData.MdtConnections[i]), types.YChild{"MdtConnections", mdtOperData.MdtConnections[i]})
    }
    mdtOperData.EntityData.Leafs = types.NewOrderedMap()

    mdtOperData.EntityData.YListKeys = []string {}

    return &(mdtOperData.EntityData)
}

// MdtOperData_MdtStreams
// MDT streams table. The list of supported streams.
type MdtOperData_MdtStreams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of a supported stream. The type is slice of string.
    Stream []interface{}
}

func (mdtStreams *MdtOperData_MdtStreams) GetEntityData() *types.CommonEntityData {
    mdtStreams.EntityData.YFilter = mdtStreams.YFilter
    mdtStreams.EntityData.YangName = "mdt-streams"
    mdtStreams.EntityData.BundleName = "cisco_ios_xe"
    mdtStreams.EntityData.ParentYangName = "mdt-oper-data"
    mdtStreams.EntityData.SegmentPath = "mdt-streams"
    mdtStreams.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtStreams.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtStreams.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtStreams.EntityData.Children = types.NewOrderedMap()
    mdtStreams.EntityData.Leafs = types.NewOrderedMap()
    mdtStreams.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", mdtStreams.Stream})

    mdtStreams.EntityData.YListKeys = []string {}

    return &(mdtStreams.EntityData)
}

// MdtOperData_MdtSubscriptions
// MDT subscription operational data.
type MdtOperData_MdtSubscriptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unique subscription identifier. The type is
    // interface{} with range: 0..4294967295.
    SubscriptionId interface{}

    // Subscription type. The type is MdtSubType.
    Type interface{}

    // Subscription state. The type is MdtSubState.
    State interface{}

    // Comments related to subcription state. The type is string.
    Comments interface{}

    // Number of updates received from sensors as candidates for notifications.
    // The type is interface{} with range: 0..18446744073709551615.
    UpdatesIn interface{}

    // Number of updates dropped due to dampening. The type is interface{} with
    // range: 0..18446744073709551615.
    UpdatesDampened interface{}

    // Number of updates dropped to other causes. The type is interface{} with
    // range: 0..18446744073709551615.
    UpdatesDropped interface{}

    // Common subscription information.
    Base MdtOperData_MdtSubscriptions_Base

    // List of MDT receivers. The type is slice of
    // MdtOperData_MdtSubscriptions_MdtReceivers.
    MdtReceivers []*MdtOperData_MdtSubscriptions_MdtReceivers
}

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetEntityData() *types.CommonEntityData {
    mdtSubscriptions.EntityData.YFilter = mdtSubscriptions.YFilter
    mdtSubscriptions.EntityData.YangName = "mdt-subscriptions"
    mdtSubscriptions.EntityData.BundleName = "cisco_ios_xe"
    mdtSubscriptions.EntityData.ParentYangName = "mdt-oper-data"
    mdtSubscriptions.EntityData.SegmentPath = "mdt-subscriptions" + types.AddKeyToken(mdtSubscriptions.SubscriptionId, "subscription-id")
    mdtSubscriptions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtSubscriptions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtSubscriptions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtSubscriptions.EntityData.Children = types.NewOrderedMap()
    mdtSubscriptions.EntityData.Children.Append("base", types.YChild{"Base", &mdtSubscriptions.Base})
    mdtSubscriptions.EntityData.Children.Append("mdt-receivers", types.YChild{"MdtReceivers", nil})
    for i := range mdtSubscriptions.MdtReceivers {
        mdtSubscriptions.EntityData.Children.Append(types.GetSegmentPath(mdtSubscriptions.MdtReceivers[i]), types.YChild{"MdtReceivers", mdtSubscriptions.MdtReceivers[i]})
    }
    mdtSubscriptions.EntityData.Leafs = types.NewOrderedMap()
    mdtSubscriptions.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", mdtSubscriptions.SubscriptionId})
    mdtSubscriptions.EntityData.Leafs.Append("type", types.YLeaf{"Type", mdtSubscriptions.Type})
    mdtSubscriptions.EntityData.Leafs.Append("state", types.YLeaf{"State", mdtSubscriptions.State})
    mdtSubscriptions.EntityData.Leafs.Append("comments", types.YLeaf{"Comments", mdtSubscriptions.Comments})
    mdtSubscriptions.EntityData.Leafs.Append("updates-in", types.YLeaf{"UpdatesIn", mdtSubscriptions.UpdatesIn})
    mdtSubscriptions.EntityData.Leafs.Append("updates-dampened", types.YLeaf{"UpdatesDampened", mdtSubscriptions.UpdatesDampened})
    mdtSubscriptions.EntityData.Leafs.Append("updates-dropped", types.YLeaf{"UpdatesDropped", mdtSubscriptions.UpdatesDropped})

    mdtSubscriptions.EntityData.YListKeys = []string {"SubscriptionId"}

    return &(mdtSubscriptions.EntityData)
}

// MdtOperData_MdtSubscriptions_Base
// Common subscription information.
type MdtOperData_MdtSubscriptions_Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the event stream being subscribed to. The type is string. The
    // default value is NETCONF.
    Stream interface{}

    // Update notification encoding. The type is string. The default value is
    // encode-xml.
    Encoding interface{}

    // Network instance name for the VRF. The type is string.
    SourceVrf interface{}

    // The source address for the notifications. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Placeholder for unset value. The type is interface{} with range:
    // 0..4294967295. The default value is 0.
    NoTrigger interface{}

    // Period of update notifications in 100ths of a second. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory. Units
    // are centiseconds.
    Period interface{}

    // If true, there is no initial update notification with the current value of
    // all the data. NOT CURRENTLY SUPPORTED. If specified, must be false. The
    // type is bool.
    NoSynchOnStart interface{}

    // Placeholder for unset value. The type is interface{} with range:
    // 0..4294967295. The default value is 0.
    NoFilter interface{}

    // XPath expression describing the set of objects wanted as part of the
    // subscription. The type is string.
    Xpath interface{}

    // TDL-URI expression describing the set of objects wanted as part of the
    // subscription. The type is string.
    TdlUri interface{}

    // Transform name is the reference to  tdl transform scheme. The type is
    // string.
    TransformName interface{}
}

func (base *MdtOperData_MdtSubscriptions_Base) GetEntityData() *types.CommonEntityData {
    base.EntityData.YFilter = base.YFilter
    base.EntityData.YangName = "base"
    base.EntityData.BundleName = "cisco_ios_xe"
    base.EntityData.ParentYangName = "mdt-subscriptions"
    base.EntityData.SegmentPath = "base"
    base.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    base.EntityData.Children = types.NewOrderedMap()
    base.EntityData.Leafs = types.NewOrderedMap()
    base.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", base.Stream})
    base.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", base.Encoding})
    base.EntityData.Leafs.Append("source-vrf", types.YLeaf{"SourceVrf", base.SourceVrf})
    base.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", base.SourceAddress})
    base.EntityData.Leafs.Append("no-trigger", types.YLeaf{"NoTrigger", base.NoTrigger})
    base.EntityData.Leafs.Append("period", types.YLeaf{"Period", base.Period})
    base.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", base.NoSynchOnStart})
    base.EntityData.Leafs.Append("no-filter", types.YLeaf{"NoFilter", base.NoFilter})
    base.EntityData.Leafs.Append("xpath", types.YLeaf{"Xpath", base.Xpath})
    base.EntityData.Leafs.Append("tdl-uri", types.YLeaf{"TdlUri", base.TdlUri})
    base.EntityData.Leafs.Append("transform-name", types.YLeaf{"TransformName", base.TransformName})

    base.EntityData.YListKeys = []string {}

    return &(base.EntityData)
}

// MdtOperData_MdtSubscriptions_MdtReceivers
// List of MDT receivers.
type MdtOperData_MdtSubscriptions_MdtReceivers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the receiver. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    Address interface{}

    // This attribute is a key. Network port of the receiver. The type is
    // interface{} with range: 0..65535. This attribute is mandatory.
    Port interface{}

    // Receiver transport protocol. The type is string.
    Protocol interface{}

    // Receiver state. The type is MdtReceiverState.
    State interface{}

    // Comments related to receiver state. The type is string.
    Comments interface{}

    // Receiver security profile. The type is string.
    SecurityProfile interface{}
}

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetEntityData() *types.CommonEntityData {
    mdtReceivers.EntityData.YFilter = mdtReceivers.YFilter
    mdtReceivers.EntityData.YangName = "mdt-receivers"
    mdtReceivers.EntityData.BundleName = "cisco_ios_xe"
    mdtReceivers.EntityData.ParentYangName = "mdt-subscriptions"
    mdtReceivers.EntityData.SegmentPath = "mdt-receivers" + types.AddKeyToken(mdtReceivers.Address, "address") + types.AddKeyToken(mdtReceivers.Port, "port")
    mdtReceivers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtReceivers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtReceivers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtReceivers.EntityData.Children = types.NewOrderedMap()
    mdtReceivers.EntityData.Leafs = types.NewOrderedMap()
    mdtReceivers.EntityData.Leafs.Append("address", types.YLeaf{"Address", mdtReceivers.Address})
    mdtReceivers.EntityData.Leafs.Append("port", types.YLeaf{"Port", mdtReceivers.Port})
    mdtReceivers.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", mdtReceivers.Protocol})
    mdtReceivers.EntityData.Leafs.Append("state", types.YLeaf{"State", mdtReceivers.State})
    mdtReceivers.EntityData.Leafs.Append("comments", types.YLeaf{"Comments", mdtReceivers.Comments})
    mdtReceivers.EntityData.Leafs.Append("security-profile", types.YLeaf{"SecurityProfile", mdtReceivers.SecurityProfile})

    mdtReceivers.EntityData.YListKeys = []string {"Address", "Port"}

    return &(mdtReceivers.EntityData)
}

// MdtOperData_MdtConnections
// MDT subscription connection operational data.
type MdtOperData_MdtConnections struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Network port. The type is interface{} with range:
    // 0..65535.
    Port interface{}

    // This attribute is a key. Network instance name for the VRF that the
    // connection originates from. The type is string.
    SourceVrf interface{}

    // This attribute is a key. The source address used for the connection. The
    // type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Transport protocol on this connection See transport-protocol from
    // subscribed-notifications for possible values. The type is string.
    Transport interface{}

    // Identity of the peer at the other end of the connection. May be empty,
    // depending on connection state. The type is string.
    PeerId interface{}

    // Connection state. The type is MdtConState.
    State interface{}

    // Security profile used with this connection. The type is string.
    SecurityProfile interface{}

    // List of subscription specific statistics for this connection. The type is
    // slice of MdtOperData_MdtConnections_MdtSubConStats.
    MdtSubConStats []*MdtOperData_MdtConnections_MdtSubConStats
}

func (mdtConnections *MdtOperData_MdtConnections) GetEntityData() *types.CommonEntityData {
    mdtConnections.EntityData.YFilter = mdtConnections.YFilter
    mdtConnections.EntityData.YangName = "mdt-connections"
    mdtConnections.EntityData.BundleName = "cisco_ios_xe"
    mdtConnections.EntityData.ParentYangName = "mdt-oper-data"
    mdtConnections.EntityData.SegmentPath = "mdt-connections" + types.AddKeyToken(mdtConnections.Address, "address") + types.AddKeyToken(mdtConnections.Port, "port") + types.AddKeyToken(mdtConnections.SourceVrf, "source-vrf") + types.AddKeyToken(mdtConnections.SourceAddress, "source-address")
    mdtConnections.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtConnections.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtConnections.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtConnections.EntityData.Children = types.NewOrderedMap()
    mdtConnections.EntityData.Children.Append("mdt-sub-con-stats", types.YChild{"MdtSubConStats", nil})
    for i := range mdtConnections.MdtSubConStats {
        mdtConnections.EntityData.Children.Append(types.GetSegmentPath(mdtConnections.MdtSubConStats[i]), types.YChild{"MdtSubConStats", mdtConnections.MdtSubConStats[i]})
    }
    mdtConnections.EntityData.Leafs = types.NewOrderedMap()
    mdtConnections.EntityData.Leafs.Append("address", types.YLeaf{"Address", mdtConnections.Address})
    mdtConnections.EntityData.Leafs.Append("port", types.YLeaf{"Port", mdtConnections.Port})
    mdtConnections.EntityData.Leafs.Append("source-vrf", types.YLeaf{"SourceVrf", mdtConnections.SourceVrf})
    mdtConnections.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", mdtConnections.SourceAddress})
    mdtConnections.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", mdtConnections.Transport})
    mdtConnections.EntityData.Leafs.Append("peer-id", types.YLeaf{"PeerId", mdtConnections.PeerId})
    mdtConnections.EntityData.Leafs.Append("state", types.YLeaf{"State", mdtConnections.State})
    mdtConnections.EntityData.Leafs.Append("security-profile", types.YLeaf{"SecurityProfile", mdtConnections.SecurityProfile})

    mdtConnections.EntityData.YListKeys = []string {"Address", "Port", "SourceVrf", "SourceAddress"}

    return &(mdtConnections.EntityData)
}

// MdtOperData_MdtConnections_MdtSubConStats
// List of subscription specific statistics for this
// connection.
type MdtOperData_MdtConnections_MdtSubConStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Subscription identifier. The type is interface{}
    // with range: 0..4294967295.
    SubId interface{}

    // Number of update notifications sent to the receiver using this
    // subscription. The type is interface{} with range: 0..18446744073709551615.
    UpdatesSent interface{}

    // Number of dropped update notifications due to error or events not in other
    // counters using this subscription. The type is interface{} with range:
    // 0..18446744073709551615.
    UpdatesDropped interface{}
}

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetEntityData() *types.CommonEntityData {
    mdtSubConStats.EntityData.YFilter = mdtSubConStats.YFilter
    mdtSubConStats.EntityData.YangName = "mdt-sub-con-stats"
    mdtSubConStats.EntityData.BundleName = "cisco_ios_xe"
    mdtSubConStats.EntityData.ParentYangName = "mdt-connections"
    mdtSubConStats.EntityData.SegmentPath = "mdt-sub-con-stats" + types.AddKeyToken(mdtSubConStats.SubId, "sub-id")
    mdtSubConStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtSubConStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtSubConStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtSubConStats.EntityData.Children = types.NewOrderedMap()
    mdtSubConStats.EntityData.Leafs = types.NewOrderedMap()
    mdtSubConStats.EntityData.Leafs.Append("sub-id", types.YLeaf{"SubId", mdtSubConStats.SubId})
    mdtSubConStats.EntityData.Leafs.Append("updates-sent", types.YLeaf{"UpdatesSent", mdtSubConStats.UpdatesSent})
    mdtSubConStats.EntityData.Leafs.Append("updates-dropped", types.YLeaf{"UpdatesDropped", mdtSubConStats.UpdatesDropped})

    mdtSubConStats.EntityData.YListKeys = []string {"SubId"}

    return &(mdtSubConStats.EntityData)
}

