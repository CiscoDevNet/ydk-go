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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the rate at which CDP packets are sent. The type is interface{}
    // with range: 5..254. The default value is 60.
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

func (cdp *Cdp) GetEntityData() *types.CommonEntityData {
    cdp.EntityData.YFilter = cdp.YFilter
    cdp.EntityData.YangName = "cdp"
    cdp.EntityData.BundleName = "cisco_ios_xr"
    cdp.EntityData.ParentYangName = "Cisco-IOS-XR-cdp-cfg"
    cdp.EntityData.SegmentPath = "Cisco-IOS-XR-cdp-cfg:cdp"
    cdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdp.EntityData.Children = types.NewOrderedMap()
    cdp.EntityData.Leafs = types.NewOrderedMap()
    cdp.EntityData.Leafs.Append("timer", types.YLeaf{"Timer", cdp.Timer})
    cdp.EntityData.Leafs.Append("advertise-v1-only", types.YLeaf{"AdvertiseV1Only", cdp.AdvertiseV1Only})
    cdp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", cdp.Enable})
    cdp.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", cdp.HoldTime})
    cdp.EntityData.Leafs.Append("log-adjacency", types.YLeaf{"LogAdjacency", cdp.LogAdjacency})

    cdp.EntityData.YListKeys = []string {}

    return &(cdp.EntityData)
}

