// This module contains a collection of YANG definitions
// for Cisco IOS-XR mdrv-lib package configuration.
// 
// This module contains definitions
// for the following management objects:
//   fast-shutdown: Fast Shutdown configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mdrv_lib_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mdrv_lib_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mdrv-lib-cfg fast-shutdown}", reflect.TypeOf(FastShutdown{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mdrv-lib-cfg:fast-shutdown", reflect.TypeOf(FastShutdown{}))
}

// FastShutdown
// Fast Shutdown configuration
type FastShutdown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Fast Shutdown for all Ethernet interfaces. The type is interface{}.
    Ethernet interface{}
}

func (fastShutdown *FastShutdown) GetEntityData() *types.CommonEntityData {
    fastShutdown.EntityData.YFilter = fastShutdown.YFilter
    fastShutdown.EntityData.YangName = "fast-shutdown"
    fastShutdown.EntityData.BundleName = "cisco_ios_xr"
    fastShutdown.EntityData.ParentYangName = "Cisco-IOS-XR-mdrv-lib-cfg"
    fastShutdown.EntityData.SegmentPath = "Cisco-IOS-XR-mdrv-lib-cfg:fast-shutdown"
    fastShutdown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastShutdown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastShutdown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastShutdown.EntityData.Children = make(map[string]types.YChild)
    fastShutdown.EntityData.Leafs = make(map[string]types.YLeaf)
    fastShutdown.EntityData.Leafs["ethernet"] = types.YLeaf{"Ethernet", fastShutdown.Ethernet}
    return &(fastShutdown.EntityData)
}

