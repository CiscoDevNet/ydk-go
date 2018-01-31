// This module contains a collection of YANG 
// definitions for configuration of streaming telemetry.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package mdt_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mdt_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mdt-cfg mdt-subscriptions}", reflect.TypeOf(MdtSubscriptions{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mdt-cfg:mdt-subscriptions", reflect.TypeOf(MdtSubscriptions{}))
}

// MdtSubscriptions
// Subscription configuration
type MdtSubscriptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of subscriptions. The type is slice of
    // MdtSubscriptions_MdtSubscription.
    MdtSubscription []MdtSubscriptions_MdtSubscription
}

func (mdtSubscriptions *MdtSubscriptions) GetFilter() yfilter.YFilter { return mdtSubscriptions.YFilter }

func (mdtSubscriptions *MdtSubscriptions) SetFilter(yf yfilter.YFilter) { mdtSubscriptions.YFilter = yf }

func (mdtSubscriptions *MdtSubscriptions) GetGoName(yname string) string {
    if yname == "mdt-subscription" { return "MdtSubscription" }
    return ""
}

func (mdtSubscriptions *MdtSubscriptions) GetSegmentPath() string {
    return "Cisco-IOS-XE-mdt-cfg:mdt-subscriptions"
}

func (mdtSubscriptions *MdtSubscriptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mdt-subscription" {
        for _, c := range mdtSubscriptions.MdtSubscription {
            if mdtSubscriptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MdtSubscriptions_MdtSubscription{}
        mdtSubscriptions.MdtSubscription = append(mdtSubscriptions.MdtSubscription, child)
        return &mdtSubscriptions.MdtSubscription[len(mdtSubscriptions.MdtSubscription)-1]
    }
    return nil
}

func (mdtSubscriptions *MdtSubscriptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mdtSubscriptions.MdtSubscription {
        children[mdtSubscriptions.MdtSubscription[i].GetSegmentPath()] = &mdtSubscriptions.MdtSubscription[i]
    }
    return children
}

func (mdtSubscriptions *MdtSubscriptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mdtSubscriptions *MdtSubscriptions) GetBundleName() string { return "cisco_ios_xe" }

func (mdtSubscriptions *MdtSubscriptions) GetYangName() string { return "mdt-subscriptions" }

func (mdtSubscriptions *MdtSubscriptions) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtSubscriptions *MdtSubscriptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtSubscriptions *MdtSubscriptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtSubscriptions *MdtSubscriptions) SetParent(parent types.Entity) { mdtSubscriptions.parent = parent }

func (mdtSubscriptions *MdtSubscriptions) GetParent() types.Entity { return mdtSubscriptions.parent }

func (mdtSubscriptions *MdtSubscriptions) GetParentYangName() string { return "Cisco-IOS-XE-mdt-cfg" }

// MdtSubscriptions_MdtSubscription
// List of subscriptions
type MdtSubscriptions_MdtSubscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Unique subscription identifier. The type is
    // interface{} with range: 0..2147483647.
    SubscriptionId interface{}

    // Common subscription information.
    Base MdtSubscriptions_MdtSubscription_Base

    // Configuration of receivers of configured subscriptions. The type is slice
    // of MdtSubscriptions_MdtSubscription_MdtReceivers.
    MdtReceivers []MdtSubscriptions_MdtSubscription_MdtReceivers
}

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetFilter() yfilter.YFilter { return mdtSubscription.YFilter }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) SetFilter(yf yfilter.YFilter) { mdtSubscription.YFilter = yf }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetGoName(yname string) string {
    if yname == "subscription-id" { return "SubscriptionId" }
    if yname == "base" { return "Base" }
    if yname == "mdt-receivers" { return "MdtReceivers" }
    return ""
}

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetSegmentPath() string {
    return "mdt-subscription" + "[subscription-id='" + fmt.Sprintf("%v", mdtSubscription.SubscriptionId) + "']"
}

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base" {
        return &mdtSubscription.Base
    }
    if childYangName == "mdt-receivers" {
        for _, c := range mdtSubscription.MdtReceivers {
            if mdtSubscription.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MdtSubscriptions_MdtSubscription_MdtReceivers{}
        mdtSubscription.MdtReceivers = append(mdtSubscription.MdtReceivers, child)
        return &mdtSubscription.MdtReceivers[len(mdtSubscription.MdtReceivers)-1]
    }
    return nil
}

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base"] = &mdtSubscription.Base
    for i := range mdtSubscription.MdtReceivers {
        children[mdtSubscription.MdtReceivers[i].GetSegmentPath()] = &mdtSubscription.MdtReceivers[i]
    }
    return children
}

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription-id"] = mdtSubscription.SubscriptionId
    return leafs
}

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetBundleName() string { return "cisco_ios_xe" }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetYangName() string { return "mdt-subscription" }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) SetParent(parent types.Entity) { mdtSubscription.parent = parent }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetParent() types.Entity { return mdtSubscription.parent }

func (mdtSubscription *MdtSubscriptions_MdtSubscription) GetParentYangName() string { return "mdt-subscriptions" }

// MdtSubscriptions_MdtSubscription_Base
// Common subscription information.
type MdtSubscriptions_MdtSubscription_Base struct {
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

func (base *MdtSubscriptions_MdtSubscription_Base) GetFilter() yfilter.YFilter { return base.YFilter }

func (base *MdtSubscriptions_MdtSubscription_Base) SetFilter(yf yfilter.YFilter) { base.YFilter = yf }

func (base *MdtSubscriptions_MdtSubscription_Base) GetGoName(yname string) string {
    if yname == "stream" { return "Stream" }
    if yname == "encoding" { return "Encoding" }
    if yname == "no-trigger" { return "NoTrigger" }
    if yname == "period" { return "Period" }
    if yname == "no-synch-on-start" { return "NoSynchOnStart" }
    if yname == "no-filter" { return "NoFilter" }
    if yname == "xpath" { return "Xpath" }
    return ""
}

func (base *MdtSubscriptions_MdtSubscription_Base) GetSegmentPath() string {
    return "base"
}

func (base *MdtSubscriptions_MdtSubscription_Base) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (base *MdtSubscriptions_MdtSubscription_Base) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (base *MdtSubscriptions_MdtSubscription_Base) GetLeafs() map[string]interface{} {
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

func (base *MdtSubscriptions_MdtSubscription_Base) GetBundleName() string { return "cisco_ios_xe" }

func (base *MdtSubscriptions_MdtSubscription_Base) GetYangName() string { return "base" }

func (base *MdtSubscriptions_MdtSubscription_Base) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (base *MdtSubscriptions_MdtSubscription_Base) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (base *MdtSubscriptions_MdtSubscription_Base) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (base *MdtSubscriptions_MdtSubscription_Base) SetParent(parent types.Entity) { base.parent = parent }

func (base *MdtSubscriptions_MdtSubscription_Base) GetParent() types.Entity { return base.parent }

func (base *MdtSubscriptions_MdtSubscription_Base) GetParentYangName() string { return "mdt-subscription" }

// MdtSubscriptions_MdtSubscription_MdtReceivers
// Configuration of receivers of configured subscriptions.
type MdtSubscriptions_MdtSubscription_MdtReceivers struct {
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

    // Receiver transport protocol. The type is string. The default value is
    // netconf.
    Protocol interface{}
}

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetFilter() yfilter.YFilter { return mdtReceivers.YFilter }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) SetFilter(yf yfilter.YFilter) { mdtReceivers.YFilter = yf }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "port" { return "Port" }
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetSegmentPath() string {
    return "mdt-receivers" + "[address='" + fmt.Sprintf("%v", mdtReceivers.Address) + "']" + "[port='" + fmt.Sprintf("%v", mdtReceivers.Port) + "']"
}

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = mdtReceivers.Address
    leafs["port"] = mdtReceivers.Port
    leafs["protocol"] = mdtReceivers.Protocol
    return leafs
}

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetBundleName() string { return "cisco_ios_xe" }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetYangName() string { return "mdt-receivers" }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) SetParent(parent types.Entity) { mdtReceivers.parent = parent }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetParent() types.Entity { return mdtReceivers.parent }

func (mdtReceivers *MdtSubscriptions_MdtSubscription_MdtReceivers) GetParentYangName() string { return "mdt-subscription" }

