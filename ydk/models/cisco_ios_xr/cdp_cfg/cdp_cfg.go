// This module contains a collection of YANG definitions
// for Cisco IOS-XR cdp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   cdp: Global CDP configuration data
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package cdp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cdp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cdp-cfg cdp}", reflect.TypeOf(Cdp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cdp-cfg:cdp", reflect.TypeOf(Cdp{}))
}

// Cdp
// Global CDP configuration data
type Cdp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the rate at which CDP packets are sent. The type is interface{}
    // with range: 5..255. The default value is 60.
    Timer interface{}

    // Enable CDPv1 only advertisements. The type is interface{}.
    AdvertiseV1Only interface{}

    // Enable or disable CDP globally. The type is bool. The default value is
    // true.
    Enable interface{}

    // Length of time (in sec) that the receiver must keep a CDP packet. The type
    // is interface{} with range: 10..255. The default value is 180.
    HoldTime interface{}

    // Enable logging of adjacency changes. The type is interface{}.
    LogAdjacency interface{}
}

func (cdp *Cdp) GetFilter() yfilter.YFilter { return cdp.YFilter }

func (cdp *Cdp) SetFilter(yf yfilter.YFilter) { cdp.YFilter = yf }

func (cdp *Cdp) GetGoName(yname string) string {
    if yname == "timer" { return "Timer" }
    if yname == "advertise-v1-only" { return "AdvertiseV1Only" }
    if yname == "enable" { return "Enable" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "log-adjacency" { return "LogAdjacency" }
    return ""
}

func (cdp *Cdp) GetSegmentPath() string {
    return "Cisco-IOS-XR-cdp-cfg:cdp"
}

func (cdp *Cdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdp *Cdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdp *Cdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timer"] = cdp.Timer
    leafs["advertise-v1-only"] = cdp.AdvertiseV1Only
    leafs["enable"] = cdp.Enable
    leafs["hold-time"] = cdp.HoldTime
    leafs["log-adjacency"] = cdp.LogAdjacency
    return leafs
}

func (cdp *Cdp) GetBundleName() string { return "cisco_ios_xr" }

func (cdp *Cdp) GetYangName() string { return "cdp" }

func (cdp *Cdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdp *Cdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdp *Cdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdp *Cdp) SetParent(parent types.Entity) { cdp.parent = parent }

func (cdp *Cdp) GetParent() types.Entity { return cdp.parent }

func (cdp *Cdp) GetParentYangName() string { return "Cisco-IOS-XR-cdp-cfg" }

