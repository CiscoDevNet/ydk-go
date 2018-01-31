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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the fabric service policy. The type is string with length: 0..63.
    FabricServicePolicy interface{}
}

func (qos *Qos) GetFilter() yfilter.YFilter { return qos.YFilter }

func (qos *Qos) SetFilter(yf yfilter.YFilter) { qos.YFilter = yf }

func (qos *Qos) GetGoName(yname string) string {
    if yname == "fabric-service-policy" { return "FabricServicePolicy" }
    return ""
}

func (qos *Qos) GetSegmentPath() string {
    return "Cisco-IOS-XR-qos-ma-cfg:qos"
}

func (qos *Qos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qos *Qos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qos *Qos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fabric-service-policy"] = qos.FabricServicePolicy
    return leafs
}

func (qos *Qos) GetBundleName() string { return "cisco_ios_xr" }

func (qos *Qos) GetYangName() string { return "qos" }

func (qos *Qos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qos *Qos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qos *Qos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qos *Qos) SetParent(parent types.Entity) { qos.parent = parent }

func (qos *Qos) GetParent() types.Entity { return qos.parent }

func (qos *Qos) GetParentYangName() string { return "Cisco-IOS-XR-qos-ma-cfg" }

