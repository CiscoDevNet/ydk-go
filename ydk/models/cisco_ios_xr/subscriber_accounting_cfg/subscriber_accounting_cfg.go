// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-accounting package configuration.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-accounting: Subscriber Configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_accounting_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_accounting_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-accounting-cfg subscriber-accounting}", reflect.TypeOf(SubscriberAccounting{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-accounting-cfg:subscriber-accounting", reflect.TypeOf(SubscriberAccounting{}))
}

// SubscriberAccounting
// Subscriber Configuration
type SubscriberAccounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber Prepaid Feature Configuration.
    PrepaidConfigurations SubscriberAccounting_PrepaidConfigurations
}

func (subscriberAccounting *SubscriberAccounting) GetFilter() yfilter.YFilter { return subscriberAccounting.YFilter }

func (subscriberAccounting *SubscriberAccounting) SetFilter(yf yfilter.YFilter) { subscriberAccounting.YFilter = yf }

func (subscriberAccounting *SubscriberAccounting) GetGoName(yname string) string {
    if yname == "prepaid-configurations" { return "PrepaidConfigurations" }
    return ""
}

func (subscriberAccounting *SubscriberAccounting) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-accounting-cfg:subscriber-accounting"
}

func (subscriberAccounting *SubscriberAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepaid-configurations" {
        return &subscriberAccounting.PrepaidConfigurations
    }
    return nil
}

func (subscriberAccounting *SubscriberAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prepaid-configurations"] = &subscriberAccounting.PrepaidConfigurations
    return children
}

func (subscriberAccounting *SubscriberAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberAccounting *SubscriberAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberAccounting *SubscriberAccounting) GetYangName() string { return "subscriber-accounting" }

func (subscriberAccounting *SubscriberAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberAccounting *SubscriberAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberAccounting *SubscriberAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberAccounting *SubscriberAccounting) SetParent(parent types.Entity) { subscriberAccounting.parent = parent }

func (subscriberAccounting *SubscriberAccounting) GetParent() types.Entity { return subscriberAccounting.parent }

func (subscriberAccounting *SubscriberAccounting) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-accounting-cfg" }

// SubscriberAccounting_PrepaidConfigurations
// Subscriber Prepaid Feature Configuration
type SubscriberAccounting_PrepaidConfigurations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prepaid configuration name or default. The type is slice of
    // SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration.
    PrepaidConfiguration []SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration
}

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetFilter() yfilter.YFilter { return prepaidConfigurations.YFilter }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) SetFilter(yf yfilter.YFilter) { prepaidConfigurations.YFilter = yf }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetGoName(yname string) string {
    if yname == "prepaid-configuration" { return "PrepaidConfiguration" }
    return ""
}

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetSegmentPath() string {
    return "prepaid-configurations"
}

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepaid-configuration" {
        for _, c := range prepaidConfigurations.PrepaidConfiguration {
            if prepaidConfigurations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration{}
        prepaidConfigurations.PrepaidConfiguration = append(prepaidConfigurations.PrepaidConfiguration, child)
        return &prepaidConfigurations.PrepaidConfiguration[len(prepaidConfigurations.PrepaidConfiguration)-1]
    }
    return nil
}

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prepaidConfigurations.PrepaidConfiguration {
        children[prepaidConfigurations.PrepaidConfiguration[i].GetSegmentPath()] = &prepaidConfigurations.PrepaidConfiguration[i]
    }
    return children
}

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetBundleName() string { return "cisco_ios_xr" }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetYangName() string { return "prepaid-configurations" }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) SetParent(parent types.Entity) { prepaidConfigurations.parent = parent }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetParent() types.Entity { return prepaidConfigurations.parent }

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetParentYangName() string { return "subscriber-accounting" }

// SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration
// Prepaid configuration name or default
type SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prepaid configuration name or default. The type is
    // string with length: 1..64.
    PrepaidConfigName interface{}

    // Password to be used when placing prepaid (re)authorization requests. The
    // type is string with length: 1..64.
    Password interface{}

    // Threshold at which to send prepaid volume quota request. The type is
    // interface{} with range: -2147483648..2147483647.
    VolumeThreshold interface{}

    // Method list to be used when placing prepaid accounting requests. The type
    // is string with length: 1..64.
    AccountingMethodList interface{}

    // Idle Threshold for which prepaid quota is valid. The type is interface{}
    // with range: -2147483648..2147483647.
    TimeHold interface{}

    // Method list to be used when placing prepaid (re)authorization requests. The
    // type is string with length: 1..64.
    AuthorMethodList interface{}

    // Prepaid quota traffic direction. The type is string.
    TrafficDirection interface{}

    // Threshold at which to send prepaid time quota request. The type is
    // interface{} with range: -2147483648..2147483647.
    TimeThreshold interface{}

    // Threshold for which prepaid quota is valid. The type is interface{} with
    // range: -2147483648..2147483647.
    TimeValid interface{}
}

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetFilter() yfilter.YFilter { return prepaidConfiguration.YFilter }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) SetFilter(yf yfilter.YFilter) { prepaidConfiguration.YFilter = yf }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetGoName(yname string) string {
    if yname == "prepaid-config-name" { return "PrepaidConfigName" }
    if yname == "password" { return "Password" }
    if yname == "volume-threshold" { return "VolumeThreshold" }
    if yname == "accounting-method-list" { return "AccountingMethodList" }
    if yname == "time-hold" { return "TimeHold" }
    if yname == "author-method-list" { return "AuthorMethodList" }
    if yname == "traffic-direction" { return "TrafficDirection" }
    if yname == "time-threshold" { return "TimeThreshold" }
    if yname == "time-valid" { return "TimeValid" }
    return ""
}

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetSegmentPath() string {
    return "prepaid-configuration" + "[prepaid-config-name='" + fmt.Sprintf("%v", prepaidConfiguration.PrepaidConfigName) + "']"
}

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prepaid-config-name"] = prepaidConfiguration.PrepaidConfigName
    leafs["password"] = prepaidConfiguration.Password
    leafs["volume-threshold"] = prepaidConfiguration.VolumeThreshold
    leafs["accounting-method-list"] = prepaidConfiguration.AccountingMethodList
    leafs["time-hold"] = prepaidConfiguration.TimeHold
    leafs["author-method-list"] = prepaidConfiguration.AuthorMethodList
    leafs["traffic-direction"] = prepaidConfiguration.TrafficDirection
    leafs["time-threshold"] = prepaidConfiguration.TimeThreshold
    leafs["time-valid"] = prepaidConfiguration.TimeValid
    return leafs
}

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetYangName() string { return "prepaid-configuration" }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) SetParent(parent types.Entity) { prepaidConfiguration.parent = parent }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetParent() types.Entity { return prepaidConfiguration.parent }

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetParentYangName() string { return "prepaid-configurations" }

