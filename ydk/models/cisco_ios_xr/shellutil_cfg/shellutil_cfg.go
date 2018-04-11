// This module contains a collection of YANG definitions
// for Cisco IOS-XR shellutil package configuration.
// 
// This module contains definitions
// for the following management objects:
//   host-names: Container Schema for hostname configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package shellutil_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package shellutil_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-shellutil-cfg host-names}", reflect.TypeOf(HostNames{}))
    ydk.RegisterEntity("Cisco-IOS-XR-shellutil-cfg:host-names", reflect.TypeOf(HostNames{}))
}

// HostNames
// Container Schema for hostname configuration
type HostNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure system's hostname. The type is string.
    HostName interface{}
}

func (hostNames *HostNames) GetEntityData() *types.CommonEntityData {
    hostNames.EntityData.YFilter = hostNames.YFilter
    hostNames.EntityData.YangName = "host-names"
    hostNames.EntityData.BundleName = "cisco_ios_xr"
    hostNames.EntityData.ParentYangName = "Cisco-IOS-XR-shellutil-cfg"
    hostNames.EntityData.SegmentPath = "Cisco-IOS-XR-shellutil-cfg:host-names"
    hostNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostNames.EntityData.Children = make(map[string]types.YChild)
    hostNames.EntityData.Leafs = make(map[string]types.YLeaf)
    hostNames.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", hostNames.HostName}
    return &(hostNames.EntityData)
}

