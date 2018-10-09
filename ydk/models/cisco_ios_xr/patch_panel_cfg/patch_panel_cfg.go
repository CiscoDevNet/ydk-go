// This module contains a collection of YANG definitions
// for Cisco IOS-XR patch-panel package configuration.
// 
// This module contains definitions
// for the following management objects:
//   patch-panel: patch-panel service submode
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package patch_panel_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package patch_panel_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-patch-panel-cfg patch-panel}", reflect.TypeOf(PatchPanel{}))
    ydk.RegisterEntity("Cisco-IOS-XR-patch-panel-cfg:patch-panel", reflect.TypeOf(PatchPanel{}))
}

// PatchPanel
// patch-panel service submode
// This type is a presence type.
type PatchPanel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable patch-panel service. The type is interface{}. This attribute is
    // mandatory.
    Enable interface{}

    // User name to be used for Authentication with Patch-Panel. The type is
    // string.
    UserName interface{}

    // Password name to be used for Authentication with Patch-Panel. The type is
    // string with pattern: (!.+)|([^!].+).
    Password interface{}

    // IP address for patch-panel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}
}

func (patchPanel *PatchPanel) GetEntityData() *types.CommonEntityData {
    patchPanel.EntityData.YFilter = patchPanel.YFilter
    patchPanel.EntityData.YangName = "patch-panel"
    patchPanel.EntityData.BundleName = "cisco_ios_xr"
    patchPanel.EntityData.ParentYangName = "Cisco-IOS-XR-patch-panel-cfg"
    patchPanel.EntityData.SegmentPath = "Cisco-IOS-XR-patch-panel-cfg:patch-panel"
    patchPanel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    patchPanel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    patchPanel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    patchPanel.EntityData.Children = types.NewOrderedMap()
    patchPanel.EntityData.Leafs = types.NewOrderedMap()
    patchPanel.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", patchPanel.Enable})
    patchPanel.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", patchPanel.UserName})
    patchPanel.EntityData.Leafs.Append("password", types.YLeaf{"Password", patchPanel.Password})
    patchPanel.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", patchPanel.Ipv4})

    patchPanel.EntityData.YListKeys = []string {}

    return &(patchPanel.EntityData)
}

