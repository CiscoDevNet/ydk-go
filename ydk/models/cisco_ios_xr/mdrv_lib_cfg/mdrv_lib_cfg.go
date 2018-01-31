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
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Fast Shutdown for all Ethernet interfaces. The type is interface{}.
    Ethernet interface{}
}

func (fastShutdown *FastShutdown) GetFilter() yfilter.YFilter { return fastShutdown.YFilter }

func (fastShutdown *FastShutdown) SetFilter(yf yfilter.YFilter) { fastShutdown.YFilter = yf }

func (fastShutdown *FastShutdown) GetGoName(yname string) string {
    if yname == "ethernet" { return "Ethernet" }
    return ""
}

func (fastShutdown *FastShutdown) GetSegmentPath() string {
    return "Cisco-IOS-XR-mdrv-lib-cfg:fast-shutdown"
}

func (fastShutdown *FastShutdown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fastShutdown *FastShutdown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fastShutdown *FastShutdown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethernet"] = fastShutdown.Ethernet
    return leafs
}

func (fastShutdown *FastShutdown) GetBundleName() string { return "cisco_ios_xr" }

func (fastShutdown *FastShutdown) GetYangName() string { return "fast-shutdown" }

func (fastShutdown *FastShutdown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fastShutdown *FastShutdown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fastShutdown *FastShutdown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fastShutdown *FastShutdown) SetParent(parent types.Entity) { fastShutdown.parent = parent }

func (fastShutdown *FastShutdown) GetParent() types.Entity { return fastShutdown.parent }

func (fastShutdown *FastShutdown) GetParentYangName() string { return "Cisco-IOS-XR-mdrv-lib-cfg" }

