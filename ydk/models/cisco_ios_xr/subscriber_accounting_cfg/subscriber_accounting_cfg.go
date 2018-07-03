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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber Prepaid Feature Configuration.
    PrepaidConfigurations SubscriberAccounting_PrepaidConfigurations
}

func (subscriberAccounting *SubscriberAccounting) GetEntityData() *types.CommonEntityData {
    subscriberAccounting.EntityData.YFilter = subscriberAccounting.YFilter
    subscriberAccounting.EntityData.YangName = "subscriber-accounting"
    subscriberAccounting.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccounting.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-accounting-cfg"
    subscriberAccounting.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-accounting-cfg:subscriber-accounting"
    subscriberAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccounting.EntityData.Children = types.NewOrderedMap()
    subscriberAccounting.EntityData.Children.Append("prepaid-configurations", types.YChild{"PrepaidConfigurations", &subscriberAccounting.PrepaidConfigurations})
    subscriberAccounting.EntityData.Leafs = types.NewOrderedMap()

    subscriberAccounting.EntityData.YListKeys = []string {}

    return &(subscriberAccounting.EntityData)
}

// SubscriberAccounting_PrepaidConfigurations
// Subscriber Prepaid Feature Configuration
type SubscriberAccounting_PrepaidConfigurations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prepaid configuration name or default. The type is slice of
    // SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration.
    PrepaidConfiguration []*SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration
}

func (prepaidConfigurations *SubscriberAccounting_PrepaidConfigurations) GetEntityData() *types.CommonEntityData {
    prepaidConfigurations.EntityData.YFilter = prepaidConfigurations.YFilter
    prepaidConfigurations.EntityData.YangName = "prepaid-configurations"
    prepaidConfigurations.EntityData.BundleName = "cisco_ios_xr"
    prepaidConfigurations.EntityData.ParentYangName = "subscriber-accounting"
    prepaidConfigurations.EntityData.SegmentPath = "prepaid-configurations"
    prepaidConfigurations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepaidConfigurations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepaidConfigurations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepaidConfigurations.EntityData.Children = types.NewOrderedMap()
    prepaidConfigurations.EntityData.Children.Append("prepaid-configuration", types.YChild{"PrepaidConfiguration", nil})
    for i := range prepaidConfigurations.PrepaidConfiguration {
        prepaidConfigurations.EntityData.Children.Append(types.GetSegmentPath(prepaidConfigurations.PrepaidConfiguration[i]), types.YChild{"PrepaidConfiguration", prepaidConfigurations.PrepaidConfiguration[i]})
    }
    prepaidConfigurations.EntityData.Leafs = types.NewOrderedMap()

    prepaidConfigurations.EntityData.YListKeys = []string {}

    return &(prepaidConfigurations.EntityData)
}

// SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration
// Prepaid configuration name or default
type SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration struct {
    EntityData types.CommonEntityData
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

func (prepaidConfiguration *SubscriberAccounting_PrepaidConfigurations_PrepaidConfiguration) GetEntityData() *types.CommonEntityData {
    prepaidConfiguration.EntityData.YFilter = prepaidConfiguration.YFilter
    prepaidConfiguration.EntityData.YangName = "prepaid-configuration"
    prepaidConfiguration.EntityData.BundleName = "cisco_ios_xr"
    prepaidConfiguration.EntityData.ParentYangName = "prepaid-configurations"
    prepaidConfiguration.EntityData.SegmentPath = "prepaid-configuration" + types.AddKeyToken(prepaidConfiguration.PrepaidConfigName, "prepaid-config-name")
    prepaidConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepaidConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepaidConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepaidConfiguration.EntityData.Children = types.NewOrderedMap()
    prepaidConfiguration.EntityData.Leafs = types.NewOrderedMap()
    prepaidConfiguration.EntityData.Leafs.Append("prepaid-config-name", types.YLeaf{"PrepaidConfigName", prepaidConfiguration.PrepaidConfigName})
    prepaidConfiguration.EntityData.Leafs.Append("password", types.YLeaf{"Password", prepaidConfiguration.Password})
    prepaidConfiguration.EntityData.Leafs.Append("volume-threshold", types.YLeaf{"VolumeThreshold", prepaidConfiguration.VolumeThreshold})
    prepaidConfiguration.EntityData.Leafs.Append("accounting-method-list", types.YLeaf{"AccountingMethodList", prepaidConfiguration.AccountingMethodList})
    prepaidConfiguration.EntityData.Leafs.Append("time-hold", types.YLeaf{"TimeHold", prepaidConfiguration.TimeHold})
    prepaidConfiguration.EntityData.Leafs.Append("author-method-list", types.YLeaf{"AuthorMethodList", prepaidConfiguration.AuthorMethodList})
    prepaidConfiguration.EntityData.Leafs.Append("traffic-direction", types.YLeaf{"TrafficDirection", prepaidConfiguration.TrafficDirection})
    prepaidConfiguration.EntityData.Leafs.Append("time-threshold", types.YLeaf{"TimeThreshold", prepaidConfiguration.TimeThreshold})
    prepaidConfiguration.EntityData.Leafs.Append("time-valid", types.YLeaf{"TimeValid", prepaidConfiguration.TimeValid})

    prepaidConfiguration.EntityData.YListKeys = []string {"PrepaidConfigName"}

    return &(prepaidConfiguration.EntityData)
}

