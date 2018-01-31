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
    parent types.Entity
    YFilter yfilter.YFilter

    // MDT streams table. The list of supported streams.
    MdtStreams MdtOperData_MdtStreams

    // MDT subscription operational data. The type is slice of
    // MdtOperData_MdtSubscriptions.
    MdtSubscriptions []MdtOperData_MdtSubscriptions

    // MDT subscription connection operational data. The type is slice of
    // MdtOperData_MdtConnections.
    MdtConnections []MdtOperData_MdtConnections
}

func (mdtOperData *MdtOperData) GetFilter() yfilter.YFilter { return mdtOperData.YFilter }

func (mdtOperData *MdtOperData) SetFilter(yf yfilter.YFilter) { mdtOperData.YFilter = yf }

func (mdtOperData *MdtOperData) GetGoName(yname string) string {
    if yname == "mdt-streams" { return "MdtStreams" }
    if yname == "mdt-subscriptions" { return "MdtSubscriptions" }
    if yname == "mdt-connections" { return "MdtConnections" }
    return ""
}

func (mdtOperData *MdtOperData) GetSegmentPath() string {
    return "Cisco-IOS-XE-mdt-oper:mdt-oper-data"
}

func (mdtOperData *MdtOperData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mdt-streams" {
        return &mdtOperData.MdtStreams
    }
    if childYangName == "mdt-subscriptions" {
        for _, c := range mdtOperData.MdtSubscriptions {
            if mdtOperData.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MdtOperData_MdtSubscriptions{}
        mdtOperData.MdtSubscriptions = append(mdtOperData.MdtSubscriptions, child)
        return &mdtOperData.MdtSubscriptions[len(mdtOperData.MdtSubscriptions)-1]
    }
    if childYangName == "mdt-connections" {
        for _, c := range mdtOperData.MdtConnections {
            if mdtOperData.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MdtOperData_MdtConnections{}
        mdtOperData.MdtConnections = append(mdtOperData.MdtConnections, child)
        return &mdtOperData.MdtConnections[len(mdtOperData.MdtConnections)-1]
    }
    return nil
}

func (mdtOperData *MdtOperData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mdt-streams"] = &mdtOperData.MdtStreams
    for i := range mdtOperData.MdtSubscriptions {
        children[mdtOperData.MdtSubscriptions[i].GetSegmentPath()] = &mdtOperData.MdtSubscriptions[i]
    }
    for i := range mdtOperData.MdtConnections {
        children[mdtOperData.MdtConnections[i].GetSegmentPath()] = &mdtOperData.MdtConnections[i]
    }
    return children
}

func (mdtOperData *MdtOperData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mdtOperData *MdtOperData) GetBundleName() string { return "cisco_ios_xe" }

func (mdtOperData *MdtOperData) GetYangName() string { return "mdt-oper-data" }

func (mdtOperData *MdtOperData) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtOperData *MdtOperData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtOperData *MdtOperData) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtOperData *MdtOperData) SetParent(parent types.Entity) { mdtOperData.parent = parent }

func (mdtOperData *MdtOperData) GetParent() types.Entity { return mdtOperData.parent }

func (mdtOperData *MdtOperData) GetParentYangName() string { return "Cisco-IOS-XE-mdt-oper" }

// MdtOperData_MdtStreams
// MDT streams table. The list of supported streams.
type MdtOperData_MdtStreams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of a supported stream. The type is slice of string.
    Stream []interface{}
}

func (mdtStreams *MdtOperData_MdtStreams) GetFilter() yfilter.YFilter { return mdtStreams.YFilter }

func (mdtStreams *MdtOperData_MdtStreams) SetFilter(yf yfilter.YFilter) { mdtStreams.YFilter = yf }

func (mdtStreams *MdtOperData_MdtStreams) GetGoName(yname string) string {
    if yname == "stream" { return "Stream" }
    return ""
}

func (mdtStreams *MdtOperData_MdtStreams) GetSegmentPath() string {
    return "mdt-streams"
}

func (mdtStreams *MdtOperData_MdtStreams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mdtStreams *MdtOperData_MdtStreams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mdtStreams *MdtOperData_MdtStreams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stream"] = mdtStreams.Stream
    return leafs
}

func (mdtStreams *MdtOperData_MdtStreams) GetBundleName() string { return "cisco_ios_xe" }

func (mdtStreams *MdtOperData_MdtStreams) GetYangName() string { return "mdt-streams" }

func (mdtStreams *MdtOperData_MdtStreams) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtStreams *MdtOperData_MdtStreams) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtStreams *MdtOperData_MdtStreams) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtStreams *MdtOperData_MdtStreams) SetParent(parent types.Entity) { mdtStreams.parent = parent }

func (mdtStreams *MdtOperData_MdtStreams) GetParent() types.Entity { return mdtStreams.parent }

func (mdtStreams *MdtOperData_MdtStreams) GetParentYangName() string { return "mdt-oper-data" }

// MdtOperData_MdtSubscriptions
// MDT subscription operational data.
type MdtOperData_MdtSubscriptions struct {
    parent types.Entity
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
    MdtReceivers []MdtOperData_MdtSubscriptions_MdtReceivers
}

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetFilter() yfilter.YFilter { return mdtSubscriptions.YFilter }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) SetFilter(yf yfilter.YFilter) { mdtSubscriptions.YFilter = yf }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "type" { return "Type" }
    if yname == "state" { return "State" }
    if yname == "comments" { return "Comments" }
    if yname == "updates-in" { return "UpdatesIn" }
    if yname == "updates-dampened" { return "UpdatesDampened" }
    if yname == "updates-dropped" { return "UpdatesDropped" }
    if yname == "base" { return "Base" }
    if yname == "mdt-receivers" { return "MdtReceivers" }
    return ""
}

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetSegmentPath() string {
    return "mdt-subscriptions" + "[subscription-id='" + fmt.Sprintf("%v", mdtSubscriptions.SubscriptionId) + "']"
}

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base" {
        return &mdtSubscriptions.Base
    }
    if childYangName == "mdt-receivers" {
        for _, c := range mdtSubscriptions.MdtReceivers {
            if mdtSubscriptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MdtOperData_MdtSubscriptions_MdtReceivers{}
        mdtSubscriptions.MdtReceivers = append(mdtSubscriptions.MdtReceivers, child)
        return &mdtSubscriptions.MdtReceivers[len(mdtSubscriptions.MdtReceivers)-1]
    }
    return nil
}

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base"] = &mdtSubscriptions.Base
    for i := range mdtSubscriptions.MdtReceivers {
        children[mdtSubscriptions.MdtReceivers[i].GetSegmentPath()] = &mdtSubscriptions.MdtReceivers[i]
    }
    return children
}

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = mdtSubscriptions.SubscriptionId
    leafs["type"] = mdtSubscriptions.Type
    leafs["state"] = mdtSubscriptions.State
    leafs["comments"] = mdtSubscriptions.Comments
    leafs["updates-in"] = mdtSubscriptions.UpdatesIn
    leafs["updates-dampened"] = mdtSubscriptions.UpdatesDampened
    leafs["updates-dropped"] = mdtSubscriptions.UpdatesDropped
    return leafs
}

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetBundleName() string { return "cisco_ios_xe" }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetYangName() string { return "mdt-subscriptions" }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) SetParent(parent types.Entity) { mdtSubscriptions.parent = parent }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetParent() types.Entity { return mdtSubscriptions.parent }

func (mdtSubscriptions *MdtOperData_MdtSubscriptions) GetParentYangName() string { return "mdt-oper-data" }

// MdtOperData_MdtSubscriptions_Base
// Common subscription information.
type MdtOperData_MdtSubscriptions_Base struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the event stream being subscribed to. The type is string. The
    // default value is NETCONF.
    Stream interface{}

    // Update notification encoding. The type is string. The default value is
    // encode-xml.
    Encoding interface{}

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
}

func (base *MdtOperData_MdtSubscriptions_Base) GetFilter() yfilter.YFilter { return base.YFilter }

func (base *MdtOperData_MdtSubscriptions_Base) SetFilter(yf yfilter.YFilter) { base.YFilter = yf }

func (base *MdtOperData_MdtSubscriptions_Base) GetGoName(yname string) string {
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "no-trigger" { return "NoTrigger" }
    if yname == "period" { return "Period" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "no-filter" { return "NoFilter" }
    if yname == "xpath" { return "Xpath" }
    return ""
}

func (base *MdtOperData_MdtSubscriptions_Base) GetSegmentPath() string {
    return "base"
}

func (base *MdtOperData_MdtSubscriptions_Base) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (base *MdtOperData_MdtSubscriptions_Base) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (base *MdtOperData_MdtSubscriptions_Base) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stream"] = base.Stream
    leafs["encoding"] = base.Encoding
    leafs["no-trigger"] = base.NoTrigger
    leafs["period"] = base.Period
    leafs["no-synch-on-start"] = base.NoSynchOnStart
    leafs["no-filter"] = base.NoFilter
    leafs["xpath"] = base.Xpath
    return leafs
}

func (base *MdtOperData_MdtSubscriptions_Base) GetBundleName() string { return "cisco_ios_xe" }

func (base *MdtOperData_MdtSubscriptions_Base) GetYangName() string { return "base" }

func (base *MdtOperData_MdtSubscriptions_Base) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (base *MdtOperData_MdtSubscriptions_Base) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (base *MdtOperData_MdtSubscriptions_Base) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (base *MdtOperData_MdtSubscriptions_Base) SetParent(parent types.Entity) { base.parent = parent }

func (base *MdtOperData_MdtSubscriptions_Base) GetParent() types.Entity { return base.parent }

func (base *MdtOperData_MdtSubscriptions_Base) GetParentYangName() string { return "mdt-subscriptions" }

// MdtOperData_MdtSubscriptions_MdtReceivers
// List of MDT receivers.
type MdtOperData_MdtSubscriptions_MdtReceivers struct {
    parent types.Entity
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
}

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetFilter() yfilter.YFilter { return mdtReceivers.YFilter }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) SetFilter(yf yfilter.YFilter) { mdtReceivers.YFilter = yf }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "port" { return "Port" }
    if yname == "protocol" { return "Protocol" }
    if yname == "state" { return "State" }
    if yname == "comments" { return "Comments" }
    return ""
}

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetSegmentPath() string {
    return "mdt-receivers" + "[address='" + fmt.Sprintf("%v", mdtReceivers.Address) + "']" + "[port='" + fmt.Sprintf("%v", mdtReceivers.Port) + "']"
}

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = mdtReceivers.Address
    leafs["port"] = mdtReceivers.Port
    leafs["protocol"] = mdtReceivers.Protocol
    leafs["state"] = mdtReceivers.State
    leafs["comments"] = mdtReceivers.Comments
    return leafs
}

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetBundleName() string { return "cisco_ios_xe" }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetYangName() string { return "mdt-receivers" }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) SetParent(parent types.Entity) { mdtReceivers.parent = parent }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetParent() types.Entity { return mdtReceivers.parent }

func (mdtReceivers *MdtOperData_MdtSubscriptions_MdtReceivers) GetParentYangName() string { return "mdt-subscriptions" }

// MdtOperData_MdtConnections
// MDT subscription connection operational data.
type MdtOperData_MdtConnections struct {
    parent types.Entity
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

    // Transport protocol on this connection See transport-protocol from
    // subscribed-notifications for possible values. The type is string.
    Transport interface{}

    // Identity of the peer at the other end of the connection. May be empty,
    // depending on connection state. The type is string.
    PeerId interface{}

    // Connection state. The type is MdtConState.
    State interface{}

    // List of subscription specific statistics for this connection. The type is
    // slice of MdtOperData_MdtConnections_MdtSubConStats.
    MdtSubConStats []MdtOperData_MdtConnections_MdtSubConStats
}

func (mdtConnections *MdtOperData_MdtConnections) GetFilter() yfilter.YFilter { return mdtConnections.YFilter }

func (mdtConnections *MdtOperData_MdtConnections) SetFilter(yf yfilter.YFilter) { mdtConnections.YFilter = yf }

func (mdtConnections *MdtOperData_MdtConnections) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "port" { return "Port" }
    if yname == "transport" { return "Transport" }
    if yname == "peer-id" { return "PeerId" }
    if yname == "state" { return "State" }
    if yname == "mdt-sub-con-stats" { return "MdtSubConStats" }
    return ""
}

func (mdtConnections *MdtOperData_MdtConnections) GetSegmentPath() string {
    return "mdt-connections" + "[address='" + fmt.Sprintf("%v", mdtConnections.Address) + "']" + "[port='" + fmt.Sprintf("%v", mdtConnections.Port) + "']"
}

func (mdtConnections *MdtOperData_MdtConnections) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mdt-sub-con-stats" {
        for _, c := range mdtConnections.MdtSubConStats {
            if mdtConnections.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MdtOperData_MdtConnections_MdtSubConStats{}
        mdtConnections.MdtSubConStats = append(mdtConnections.MdtSubConStats, child)
        return &mdtConnections.MdtSubConStats[len(mdtConnections.MdtSubConStats)-1]
    }
    return nil
}

func (mdtConnections *MdtOperData_MdtConnections) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mdtConnections.MdtSubConStats {
        children[mdtConnections.MdtSubConStats[i].GetSegmentPath()] = &mdtConnections.MdtSubConStats[i]
    }
    return children
}

func (mdtConnections *MdtOperData_MdtConnections) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = mdtConnections.Address
    leafs["port"] = mdtConnections.Port
    leafs["transport"] = mdtConnections.Transport
    leafs["peer-id"] = mdtConnections.PeerId
    leafs["state"] = mdtConnections.State
    return leafs
}

func (mdtConnections *MdtOperData_MdtConnections) GetBundleName() string { return "cisco_ios_xe" }

func (mdtConnections *MdtOperData_MdtConnections) GetYangName() string { return "mdt-connections" }

func (mdtConnections *MdtOperData_MdtConnections) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtConnections *MdtOperData_MdtConnections) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtConnections *MdtOperData_MdtConnections) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtConnections *MdtOperData_MdtConnections) SetParent(parent types.Entity) { mdtConnections.parent = parent }

func (mdtConnections *MdtOperData_MdtConnections) GetParent() types.Entity { return mdtConnections.parent }

func (mdtConnections *MdtOperData_MdtConnections) GetParentYangName() string { return "mdt-oper-data" }

// MdtOperData_MdtConnections_MdtSubConStats
// List of subscription specific statistics for this
// connection.
type MdtOperData_MdtConnections_MdtSubConStats struct {
    parent types.Entity
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

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetFilter() yfilter.YFilter { return mdtSubConStats.YFilter }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) SetFilter(yf yfilter.YFilter) { mdtSubConStats.YFilter = yf }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetGoName(yname string) string {
    if yname == "sub-id" { return "SubId" }
    if yname == "updates-sent" { return "UpdatesSent" }
    if yname == "updates-dropped" { return "UpdatesDropped" }
    return ""
}

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetSegmentPath() string {
    return "mdt-sub-con-stats" + "[sub-id='" + fmt.Sprintf("%v", mdtSubConStats.SubId) + "']"
}

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sub-id"] = mdtSubConStats.SubId
    leafs["updates-sent"] = mdtSubConStats.UpdatesSent
    leafs["updates-dropped"] = mdtSubConStats.UpdatesDropped
    return leafs
}

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetBundleName() string { return "cisco_ios_xe" }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetYangName() string { return "mdt-sub-con-stats" }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) SetParent(parent types.Entity) { mdtSubConStats.parent = parent }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetParent() types.Entity { return mdtSubConStats.parent }

func (mdtSubConStats *MdtOperData_MdtConnections_MdtSubConStats) GetParentYangName() string { return "mdt-connections" }

