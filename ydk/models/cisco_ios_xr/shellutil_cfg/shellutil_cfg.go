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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure system's hostname. The type is string.
    HostName interface{}
}

func (hostNames *HostNames) GetFilter() yfilter.YFilter { return hostNames.YFilter }

func (hostNames *HostNames) SetFilter(yf yfilter.YFilter) { hostNames.YFilter = yf }

func (hostNames *HostNames) GetGoName(yname string) string {
    if yname == "host-name" { return "HostName" }
    return ""
}

func (hostNames *HostNames) GetSegmentPath() string {
    return "Cisco-IOS-XR-shellutil-cfg:host-names"
}

func (hostNames *HostNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostNames *HostNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostNames *HostNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name"] = hostNames.HostName
    return leafs
}

func (hostNames *HostNames) GetBundleName() string { return "cisco_ios_xr" }

func (hostNames *HostNames) GetYangName() string { return "host-names" }

func (hostNames *HostNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostNames *HostNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostNames *HostNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostNames *HostNames) SetParent(parent types.Entity) { hostNames.parent = parent }

func (hostNames *HostNames) GetParent() types.Entity { return hostNames.parent }

func (hostNames *HostNames) GetParentYangName() string { return "Cisco-IOS-XR-shellutil-cfg" }

