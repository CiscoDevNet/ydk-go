// This module contains a collection of YANG definitions
// for Cisco IOS-XR qos package configuration.
// 
// This module contains definitions
// for the following management objects:
//   qos: Global QOS configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-l2vpn-cfg,
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package qos_ma_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package qos_ma_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-qos-ma-cfg qos}", reflect.TypeOf(Qos{}))
    ydk.RegisterEntity("Cisco-IOS-XR-qos-ma-cfg:qos", reflect.TypeOf(Qos{}))
}

// QosFieldNotSupported represents Qos field not supported
type QosFieldNotSupported string

const (
    // Dummy data type leave unspecified
    QosFieldNotSupported_not_supported QosFieldNotSupported = "not-supported"
)

// QosPolicyAccount represents Qos policy account
type QosPolicyAccount string

const (
    // Turn on Layer 1 accounting
    QosPolicyAccount_layer1 QosPolicyAccount = "layer1"

    // Turn on Layer 2 accounting
    QosPolicyAccount_layer2 QosPolicyAccount = "layer2"

    // Turn on Layer 2 accounting
    QosPolicyAccount_nolayer2 QosPolicyAccount = "nolayer2"

    // User defined accounting
    QosPolicyAccount_user_defined QosPolicyAccount = "user-defined"
)

// Qos
// Global QOS configuration.
type Qos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the fabric service policy. The type is string with length: 0..63.
    FabricServicePolicy interface{}
}

func (qos *Qos) GetEntityData() *types.CommonEntityData {
    qos.EntityData.YFilter = qos.YFilter
    qos.EntityData.YangName = "qos"
    qos.EntityData.BundleName = "cisco_ios_xr"
    qos.EntityData.ParentYangName = "Cisco-IOS-XR-qos-ma-cfg"
    qos.EntityData.SegmentPath = "Cisco-IOS-XR-qos-ma-cfg:qos"
    qos.EntityData.AbsolutePath = qos.EntityData.SegmentPath
    qos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qos.EntityData.Children = types.NewOrderedMap()
    qos.EntityData.Leafs = types.NewOrderedMap()
    qos.EntityData.Leafs.Append("fabric-service-policy", types.YLeaf{"FabricServicePolicy", qos.FabricServicePolicy})

    qos.EntityData.YListKeys = []string {}

    return &(qos.EntityData)
}

