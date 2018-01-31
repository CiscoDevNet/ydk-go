// This module contains a collection of YANG definitions
// for Cisco IOS-XR patch-panel package configuration.
// 
// This module contains definitions
// for the following management objects:
//   patch-panel: patch-panel service submode
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

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

func (patchPanel *PatchPanel) GetFilter() yfilter.YFilter { return patchPanel.YFilter }

func (patchPanel *PatchPanel) SetFilter(yf yfilter.YFilter) { patchPanel.YFilter = yf }

func (patchPanel *PatchPanel) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "user-name" { return "UserName" }
    if yname == "password" { return "Password" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (patchPanel *PatchPanel) GetSegmentPath() string {
    return "Cisco-IOS-XR-patch-panel-cfg:patch-panel"
}

func (patchPanel *PatchPanel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (patchPanel *PatchPanel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (patchPanel *PatchPanel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = patchPanel.Enable
    leafs["user-name"] = patchPanel.UserName
    leafs["password"] = patchPanel.Password
    leafs["ipv4"] = patchPanel.Ipv4
    return leafs
}

func (patchPanel *PatchPanel) GetBundleName() string { return "cisco_ios_xr" }

func (patchPanel *PatchPanel) GetYangName() string { return "patch-panel" }

func (patchPanel *PatchPanel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (patchPanel *PatchPanel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (patchPanel *PatchPanel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (patchPanel *PatchPanel) SetParent(parent types.Entity) { patchPanel.parent = parent }

func (patchPanel *PatchPanel) GetParent() types.Entity { return patchPanel.parent }

func (patchPanel *PatchPanel) GetParentYangName() string { return "Cisco-IOS-XR-patch-panel-cfg" }

