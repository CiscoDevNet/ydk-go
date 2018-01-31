// This module contains a collection of YANG definitions
// for Cisco IOS-XR rgmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   redundancy-group-manager: Redundancy Group Manager
//     Configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package rgmgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rgmgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-rgmgr-cfg redundancy-group-manager}", reflect.TypeOf(RedundancyGroupManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager", reflect.TypeOf(RedundancyGroupManager{}))
}

// IccpMode represents Iccp mode
type IccpMode string

const (
    // Run the ICCP group in Singleton mode
    IccpMode_singleton IccpMode = "singleton"
)

// RedundancyGroupManager
// Redundancy Group Manager Configuration
type RedundancyGroupManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable redundancy group manager. The type is interface{}.
    Enable interface{}

    // MR-APS groups.
    Aps RedundancyGroupManager_Aps

    // ICCP configuration.
    Iccp RedundancyGroupManager_Iccp
}

func (redundancyGroupManager *RedundancyGroupManager) GetFilter() yfilter.YFilter { return redundancyGroupManager.YFilter }

func (redundancyGroupManager *RedundancyGroupManager) SetFilter(yf yfilter.YFilter) { redundancyGroupManager.YFilter = yf }

func (redundancyGroupManager *RedundancyGroupManager) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "aps" { return "Aps" }
    if yname == "iccp" { return "Iccp" }
    return ""
}

func (redundancyGroupManager *RedundancyGroupManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-rgmgr-cfg:redundancy-group-manager"
}

func (redundancyGroupManager *RedundancyGroupManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aps" {
        return &redundancyGroupManager.Aps
    }
    if childYangName == "iccp" {
        return &redundancyGroupManager.Iccp
    }
    return nil
}

func (redundancyGroupManager *RedundancyGroupManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aps"] = &redundancyGroupManager.Aps
    children["iccp"] = &redundancyGroupManager.Iccp
    return children
}

func (redundancyGroupManager *RedundancyGroupManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = redundancyGroupManager.Enable
    return leafs
}

func (redundancyGroupManager *RedundancyGroupManager) GetBundleName() string { return "cisco_ios_xr" }

func (redundancyGroupManager *RedundancyGroupManager) GetYangName() string { return "redundancy-group-manager" }

func (redundancyGroupManager *RedundancyGroupManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redundancyGroupManager *RedundancyGroupManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redundancyGroupManager *RedundancyGroupManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redundancyGroupManager *RedundancyGroupManager) SetParent(parent types.Entity) { redundancyGroupManager.parent = parent }

func (redundancyGroupManager *RedundancyGroupManager) GetParent() types.Entity { return redundancyGroupManager.parent }

func (redundancyGroupManager *RedundancyGroupManager) GetParentYangName() string { return "Cisco-IOS-XR-rgmgr-cfg" }

// RedundancyGroupManager_Aps
// MR-APS groups
type RedundancyGroupManager_Aps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default SONET controller backup configuration.
    DefaultRedundancyGroup RedundancyGroupManager_Aps_DefaultRedundancyGroup

    // Redundancy Group Table.
    Groups RedundancyGroupManager_Aps_Groups
}

func (aps *RedundancyGroupManager_Aps) GetFilter() yfilter.YFilter { return aps.YFilter }

func (aps *RedundancyGroupManager_Aps) SetFilter(yf yfilter.YFilter) { aps.YFilter = yf }

func (aps *RedundancyGroupManager_Aps) GetGoName(yname string) string {
    if yname == "default-redundancy-group" { return "DefaultRedundancyGroup" }
    if yname == "groups" { return "Groups" }
    return ""
}

func (aps *RedundancyGroupManager_Aps) GetSegmentPath() string {
    return "aps"
}

func (aps *RedundancyGroupManager_Aps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-redundancy-group" {
        return &aps.DefaultRedundancyGroup
    }
    if childYangName == "groups" {
        return &aps.Groups
    }
    return nil
}

func (aps *RedundancyGroupManager_Aps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-redundancy-group"] = &aps.DefaultRedundancyGroup
    children["groups"] = &aps.Groups
    return children
}

func (aps *RedundancyGroupManager_Aps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aps *RedundancyGroupManager_Aps) GetBundleName() string { return "cisco_ios_xr" }

func (aps *RedundancyGroupManager_Aps) GetYangName() string { return "aps" }

func (aps *RedundancyGroupManager_Aps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aps *RedundancyGroupManager_Aps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aps *RedundancyGroupManager_Aps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aps *RedundancyGroupManager_Aps) SetParent(parent types.Entity) { aps.parent = parent }

func (aps *RedundancyGroupManager_Aps) GetParent() types.Entity { return aps.parent }

func (aps *RedundancyGroupManager_Aps) GetParentYangName() string { return "redundancy-group-manager" }

// RedundancyGroupManager_Aps_DefaultRedundancyGroup
// Default SONET controller backup configuration
type RedundancyGroupManager_Aps_DefaultRedundancyGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address of remote peer. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Backup interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    BackupInterfaceName interface{}
}

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetFilter() yfilter.YFilter { return defaultRedundancyGroup.YFilter }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) SetFilter(yf yfilter.YFilter) { defaultRedundancyGroup.YFilter = yf }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetGoName(yname string) string {
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "backup-interface-name" { return "BackupInterfaceName" }
    return ""
}

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetSegmentPath() string {
    return "default-redundancy-group"
}

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop-address"] = defaultRedundancyGroup.NextHopAddress
    leafs["backup-interface-name"] = defaultRedundancyGroup.BackupInterfaceName
    return leafs
}

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetBundleName() string { return "cisco_ios_xr" }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetYangName() string { return "default-redundancy-group" }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) SetParent(parent types.Entity) { defaultRedundancyGroup.parent = parent }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetParent() types.Entity { return defaultRedundancyGroup.parent }

func (defaultRedundancyGroup *RedundancyGroupManager_Aps_DefaultRedundancyGroup) GetParentYangName() string { return "aps" }

// RedundancyGroupManager_Aps_Groups
// Redundancy Group Table
type RedundancyGroupManager_Aps_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redundancy Group Configuration. The type is slice of
    // RedundancyGroupManager_Aps_Groups_Group.
    Group []RedundancyGroupManager_Aps_Groups_Group
}

func (groups *RedundancyGroupManager_Aps_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *RedundancyGroupManager_Aps_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *RedundancyGroupManager_Aps_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *RedundancyGroupManager_Aps_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *RedundancyGroupManager_Aps_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RedundancyGroupManager_Aps_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *RedundancyGroupManager_Aps_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *RedundancyGroupManager_Aps_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *RedundancyGroupManager_Aps_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *RedundancyGroupManager_Aps_Groups) GetYangName() string { return "groups" }

func (groups *RedundancyGroupManager_Aps_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *RedundancyGroupManager_Aps_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *RedundancyGroupManager_Aps_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *RedundancyGroupManager_Aps_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *RedundancyGroupManager_Aps_Groups) GetParent() types.Entity { return groups.parent }

func (groups *RedundancyGroupManager_Aps_Groups) GetParentYangName() string { return "aps" }

// RedundancyGroupManager_Aps_Groups_Group
// Redundancy Group Configuration
type RedundancyGroupManager_Aps_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The redundancy group ID. The type is interface{}
    // with range: 1..32.
    GroupId interface{}

    // Controller configuration.
    Controllers RedundancyGroupManager_Aps_Groups_Group_Controllers
}

func (group *RedundancyGroupManager_Aps_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *RedundancyGroupManager_Aps_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *RedundancyGroupManager_Aps_Groups_Group) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "controllers" { return "Controllers" }
    return ""
}

func (group *RedundancyGroupManager_Aps_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']"
}

func (group *RedundancyGroupManager_Aps_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controllers" {
        return &group.Controllers
    }
    return nil
}

func (group *RedundancyGroupManager_Aps_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["controllers"] = &group.Controllers
    return children
}

func (group *RedundancyGroupManager_Aps_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = group.GroupId
    return leafs
}

func (group *RedundancyGroupManager_Aps_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *RedundancyGroupManager_Aps_Groups_Group) GetYangName() string { return "group" }

func (group *RedundancyGroupManager_Aps_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *RedundancyGroupManager_Aps_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *RedundancyGroupManager_Aps_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *RedundancyGroupManager_Aps_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *RedundancyGroupManager_Aps_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *RedundancyGroupManager_Aps_Groups_Group) GetParentYangName() string { return "groups" }

// RedundancyGroupManager_Aps_Groups_Group_Controllers
// Controller configuration
type RedundancyGroupManager_Aps_Groups_Group_Controllers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is slice of
    // RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller.
    Controller []RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller
}

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetFilter() yfilter.YFilter { return controllers.YFilter }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) SetFilter(yf yfilter.YFilter) { controllers.YFilter = yf }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetGoName(yname string) string {
    if yname == "controller" { return "Controller" }
    return ""
}

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetSegmentPath() string {
    return "controllers"
}

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controller" {
        for _, c := range controllers.Controller {
            if controllers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller{}
        controllers.Controller = append(controllers.Controller, child)
        return &controllers.Controller[len(controllers.Controller)-1]
    }
    return nil
}

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controllers.Controller {
        children[controllers.Controller[i].GetSegmentPath()] = &controllers.Controller[i]
    }
    return children
}

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetBundleName() string { return "cisco_ios_xr" }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetYangName() string { return "controllers" }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) SetParent(parent types.Entity) { controllers.parent = parent }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetParent() types.Entity { return controllers.parent }

func (controllers *RedundancyGroupManager_Aps_Groups_Group_Controllers) GetParentYangName() string { return "group" }

// RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller
// none
type RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Controller Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // IPv4 address of remote peer. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Backup interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    BackupInterfaceName interface{}
}

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetFilter() yfilter.YFilter { return controller.YFilter }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) SetFilter(yf yfilter.YFilter) { controller.YFilter = yf }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetGoName(yname string) string {
    if yname == "controller-name" { return "ControllerName" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "backup-interface-name" { return "BackupInterfaceName" }
    return ""
}

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetSegmentPath() string {
    return "controller" + "[controller-name='" + fmt.Sprintf("%v", controller.ControllerName) + "']"
}

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller-name"] = controller.ControllerName
    leafs["next-hop-address"] = controller.NextHopAddress
    leafs["backup-interface-name"] = controller.BackupInterfaceName
    return leafs
}

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetBundleName() string { return "cisco_ios_xr" }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetYangName() string { return "controller" }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) SetParent(parent types.Entity) { controller.parent = parent }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetParent() types.Entity { return controller.parent }

func (controller *RedundancyGroupManager_Aps_Groups_Group_Controllers_Controller) GetParentYangName() string { return "controllers" }

// RedundancyGroupManager_Iccp
// ICCP configuration
type RedundancyGroupManager_Iccp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redundancy Group Table Configuration.
    IccpGroups RedundancyGroupManager_Iccp_IccpGroups
}

func (iccp *RedundancyGroupManager_Iccp) GetFilter() yfilter.YFilter { return iccp.YFilter }

func (iccp *RedundancyGroupManager_Iccp) SetFilter(yf yfilter.YFilter) { iccp.YFilter = yf }

func (iccp *RedundancyGroupManager_Iccp) GetGoName(yname string) string {
    if yname == "iccp-groups" { return "IccpGroups" }
    return ""
}

func (iccp *RedundancyGroupManager_Iccp) GetSegmentPath() string {
    return "iccp"
}

func (iccp *RedundancyGroupManager_Iccp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "iccp-groups" {
        return &iccp.IccpGroups
    }
    return nil
}

func (iccp *RedundancyGroupManager_Iccp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["iccp-groups"] = &iccp.IccpGroups
    return children
}

func (iccp *RedundancyGroupManager_Iccp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iccp *RedundancyGroupManager_Iccp) GetBundleName() string { return "cisco_ios_xr" }

func (iccp *RedundancyGroupManager_Iccp) GetYangName() string { return "iccp" }

func (iccp *RedundancyGroupManager_Iccp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iccp *RedundancyGroupManager_Iccp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iccp *RedundancyGroupManager_Iccp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iccp *RedundancyGroupManager_Iccp) SetParent(parent types.Entity) { iccp.parent = parent }

func (iccp *RedundancyGroupManager_Iccp) GetParent() types.Entity { return iccp.parent }

func (iccp *RedundancyGroupManager_Iccp) GetParentYangName() string { return "redundancy-group-manager" }

// RedundancyGroupManager_Iccp_IccpGroups
// Redundancy Group Table Configuration
type RedundancyGroupManager_Iccp_IccpGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redundancy Group Configuration. The type is slice of
    // RedundancyGroupManager_Iccp_IccpGroups_IccpGroup.
    IccpGroup []RedundancyGroupManager_Iccp_IccpGroups_IccpGroup
}

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetFilter() yfilter.YFilter { return iccpGroups.YFilter }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) SetFilter(yf yfilter.YFilter) { iccpGroups.YFilter = yf }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetGoName(yname string) string {
    if yname == "iccp-group" { return "IccpGroup" }
    return ""
}

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetSegmentPath() string {
    return "iccp-groups"
}

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "iccp-group" {
        for _, c := range iccpGroups.IccpGroup {
            if iccpGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RedundancyGroupManager_Iccp_IccpGroups_IccpGroup{}
        iccpGroups.IccpGroup = append(iccpGroups.IccpGroup, child)
        return &iccpGroups.IccpGroup[len(iccpGroups.IccpGroup)-1]
    }
    return nil
}

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range iccpGroups.IccpGroup {
        children[iccpGroups.IccpGroup[i].GetSegmentPath()] = &iccpGroups.IccpGroup[i]
    }
    return children
}

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetBundleName() string { return "cisco_ios_xr" }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetYangName() string { return "iccp-groups" }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) SetParent(parent types.Entity) { iccpGroups.parent = parent }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetParent() types.Entity { return iccpGroups.parent }

func (iccpGroups *RedundancyGroupManager_Iccp_IccpGroups) GetParentYangName() string { return "iccp" }

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup
// Redundancy Group Configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The redundancy icc group number. The type is
    // interface{} with range: 1..4294967295.
    GroupNumber interface{}

    // ICCP isolation recovery delay. The type is interface{} with range: 30..600.
    // Units are second.
    IsolationRecoveryDelay interface{}

    // ICCP mode. The type is IccpMode.
    Mode interface{}

    // ICCP backbone configuration.
    Backbones RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones

    // ICCP member configuration.
    Members RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members

    // Multi-chassis Link Aggregation Control Protocol commands.
    Mlacp RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp

    // nV Satellite configuration.
    NvSatellite RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite
}

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetFilter() yfilter.YFilter { return iccpGroup.YFilter }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) SetFilter(yf yfilter.YFilter) { iccpGroup.YFilter = yf }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetGoName(yname string) string {
    if yname == "group-number" { return "GroupNumber" }
    if yname == "isolation-recovery-delay" { return "IsolationRecoveryDelay" }
    if yname == "mode" { return "Mode" }
    if yname == "backbones" { return "Backbones" }
    if yname == "members" { return "Members" }
    if yname == "Cisco-IOS-XR-bundlemgr-cfg:mlacp" { return "Mlacp" }
    if yname == "Cisco-IOS-XR-icpe-infra-cfg:nv-satellite" { return "NvSatellite" }
    return ""
}

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetSegmentPath() string {
    return "iccp-group" + "[group-number='" + fmt.Sprintf("%v", iccpGroup.GroupNumber) + "']"
}

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "backbones" {
        return &iccpGroup.Backbones
    }
    if childYangName == "members" {
        return &iccpGroup.Members
    }
    if childYangName == "Cisco-IOS-XR-bundlemgr-cfg:mlacp" {
        return &iccpGroup.Mlacp
    }
    if childYangName == "Cisco-IOS-XR-icpe-infra-cfg:nv-satellite" {
        return &iccpGroup.NvSatellite
    }
    return nil
}

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["backbones"] = &iccpGroup.Backbones
    children["members"] = &iccpGroup.Members
    children["Cisco-IOS-XR-bundlemgr-cfg:mlacp"] = &iccpGroup.Mlacp
    children["Cisco-IOS-XR-icpe-infra-cfg:nv-satellite"] = &iccpGroup.NvSatellite
    return children
}

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-number"] = iccpGroup.GroupNumber
    leafs["isolation-recovery-delay"] = iccpGroup.IsolationRecoveryDelay
    leafs["mode"] = iccpGroup.Mode
    return leafs
}

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetBundleName() string { return "cisco_ios_xr" }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetYangName() string { return "iccp-group" }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) SetParent(parent types.Entity) { iccpGroup.parent = parent }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetParent() types.Entity { return iccpGroup.parent }

func (iccpGroup *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup) GetParentYangName() string { return "iccp-groups" }

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones
// ICCP backbone configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICCP backbone interface configuration. The type is slice of
    // RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone.
    Backbone []RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone
}

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetFilter() yfilter.YFilter { return backbones.YFilter }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) SetFilter(yf yfilter.YFilter) { backbones.YFilter = yf }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetGoName(yname string) string {
    if yname == "backbone" { return "Backbone" }
    return ""
}

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetSegmentPath() string {
    return "backbones"
}

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "backbone" {
        for _, c := range backbones.Backbone {
            if backbones.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone{}
        backbones.Backbone = append(backbones.Backbone, child)
        return &backbones.Backbone[len(backbones.Backbone)-1]
    }
    return nil
}

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range backbones.Backbone {
        children[backbones.Backbone[i].GetSegmentPath()] = &backbones.Backbone[i]
    }
    return children
}

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetBundleName() string { return "cisco_ios_xr" }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetYangName() string { return "backbones" }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) SetParent(parent types.Entity) { backbones.parent = parent }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetParent() types.Entity { return backbones.parent }

func (backbones *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones) GetParentYangName() string { return "iccp-group" }

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone
// ICCP backbone interface configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    BackboneName interface{}
}

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetFilter() yfilter.YFilter { return backbone.YFilter }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) SetFilter(yf yfilter.YFilter) { backbone.YFilter = yf }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetGoName(yname string) string {
    if yname == "backbone-name" { return "BackboneName" }
    return ""
}

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetSegmentPath() string {
    return "backbone" + "[backbone-name='" + fmt.Sprintf("%v", backbone.BackboneName) + "']"
}

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["backbone-name"] = backbone.BackboneName
    return leafs
}

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetBundleName() string { return "cisco_ios_xr" }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetYangName() string { return "backbone" }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) SetParent(parent types.Entity) { backbone.parent = parent }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetParent() types.Entity { return backbone.parent }

func (backbone *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Backbones_Backbone) GetParentYangName() string { return "backbones" }

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members
// ICCP member configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICCP member configuration. The type is slice of
    // RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member.
    Member []RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member
}

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetFilter() yfilter.YFilter { return members.YFilter }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) SetFilter(yf yfilter.YFilter) { members.YFilter = yf }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetGoName(yname string) string {
    if yname == "member" { return "Member" }
    return ""
}

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetSegmentPath() string {
    return "members"
}

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member" {
        for _, c := range members.Member {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member{}
        members.Member = append(members.Member, child)
        return &members.Member[len(members.Member)-1]
    }
    return nil
}

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range members.Member {
        children[members.Member[i].GetSegmentPath()] = &members.Member[i]
    }
    return children
}

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetBundleName() string { return "cisco_ios_xr" }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetYangName() string { return "members" }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) SetParent(parent types.Entity) { members.parent = parent }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetParent() types.Entity { return members.parent }

func (members *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members) GetParentYangName() string { return "iccp-group" }

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member
// ICCP member configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}
}

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetFilter() yfilter.YFilter { return member.YFilter }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) SetFilter(yf yfilter.YFilter) { member.YFilter = yf }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    return ""
}

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetSegmentPath() string {
    return "member" + "[neighbor-address='" + fmt.Sprintf("%v", member.NeighborAddress) + "']"
}

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = member.NeighborAddress
    return leafs
}

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetBundleName() string { return "cisco_ios_xr" }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetYangName() string { return "member" }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) SetParent(parent types.Entity) { member.parent = parent }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetParent() types.Entity { return member.parent }

func (member *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Members_Member) GetParentYangName() string { return "members" }

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp
// Multi-chassis Link Aggregation Control Protocol
// commands
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of seconds to wait before assuming mLACP peer is down. The type is
    // interface{} with range: 0..65534.
    ConnectTimeout interface{}

    // Unique LACP identifier for this system. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemMac interface{}

    // Unique identifier for this system in the ICCP Group. The type is
    // interface{} with range: 0..7.
    Node interface{}

    // Priority for this system. Lower value is higher priority. The type is
    // interface{} with range: 1..65535.
    SystemPriority interface{}
}

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetFilter() yfilter.YFilter { return mlacp.YFilter }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) SetFilter(yf yfilter.YFilter) { mlacp.YFilter = yf }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetGoName(yname string) string {
    if yname == "connect-timeout" { return "ConnectTimeout" }
    if yname == "system-mac" { return "SystemMac" }
    if yname == "node" { return "Node" }
    if yname == "system-priority" { return "SystemPriority" }
    return ""
}

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetSegmentPath() string {
    return "Cisco-IOS-XR-bundlemgr-cfg:mlacp"
}

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["connect-timeout"] = mlacp.ConnectTimeout
    leafs["system-mac"] = mlacp.SystemMac
    leafs["node"] = mlacp.Node
    leafs["system-priority"] = mlacp.SystemPriority
    return leafs
}

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetBundleName() string { return "cisco_ios_xr" }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetYangName() string { return "mlacp" }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) SetParent(parent types.Entity) { mlacp.parent = parent }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetParent() types.Entity { return mlacp.parent }

func (mlacp *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_Mlacp) GetParentYangName() string { return "iccp-group" }

// RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite
// nV Satellite configuration
type RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Optional identifier for this system. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemMac interface{}
}

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetFilter() yfilter.YFilter { return nvSatellite.YFilter }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) SetFilter(yf yfilter.YFilter) { nvSatellite.YFilter = yf }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetGoName(yname string) string {
    if yname == "system-mac" { return "SystemMac" }
    return ""
}

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetSegmentPath() string {
    return "Cisco-IOS-XR-icpe-infra-cfg:nv-satellite"
}

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-mac"] = nvSatellite.SystemMac
    return leafs
}

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetBundleName() string { return "cisco_ios_xr" }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetYangName() string { return "nv-satellite" }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) SetParent(parent types.Entity) { nvSatellite.parent = parent }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetParent() types.Entity { return nvSatellite.parent }

func (nvSatellite *RedundancyGroupManager_Iccp_IccpGroups_IccpGroup_NvSatellite) GetParentYangName() string { return "iccp-group" }

