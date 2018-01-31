// This module contains a collection of YANG definitions
// for Cisco IOS-XR rgmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   redundancy-group-manager: Redundancy group manager operational
//     data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package rgmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rgmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-rgmgr-oper redundancy-group-manager}", reflect.TypeOf(RedundancyGroupManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-rgmgr-oper:redundancy-group-manager", reflect.TypeOf(RedundancyGroupManager{}))
}

// RedundancyGroupManager
// Redundancy group manager operational data
type RedundancyGroupManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redundancy group manager data.
    Controllers RedundancyGroupManager_Controllers
}

func (redundancyGroupManager *RedundancyGroupManager) GetFilter() yfilter.YFilter { return redundancyGroupManager.YFilter }

func (redundancyGroupManager *RedundancyGroupManager) SetFilter(yf yfilter.YFilter) { redundancyGroupManager.YFilter = yf }

func (redundancyGroupManager *RedundancyGroupManager) GetGoName(yname string) string {
    if yname == "controllers" { return "Controllers" }
    return ""
}

func (redundancyGroupManager *RedundancyGroupManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-rgmgr-oper:redundancy-group-manager"
}

func (redundancyGroupManager *RedundancyGroupManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controllers" {
        return &redundancyGroupManager.Controllers
    }
    return nil
}

func (redundancyGroupManager *RedundancyGroupManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["controllers"] = &redundancyGroupManager.Controllers
    return children
}

func (redundancyGroupManager *RedundancyGroupManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
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

func (redundancyGroupManager *RedundancyGroupManager) GetParentYangName() string { return "Cisco-IOS-XR-rgmgr-oper" }

// RedundancyGroupManager_Controllers
// Redundancy group manager data
type RedundancyGroupManager_Controllers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Display redundancy group by controller name. The type is slice of
    // RedundancyGroupManager_Controllers_Controller.
    Controller []RedundancyGroupManager_Controllers_Controller
}

func (controllers *RedundancyGroupManager_Controllers) GetFilter() yfilter.YFilter { return controllers.YFilter }

func (controllers *RedundancyGroupManager_Controllers) SetFilter(yf yfilter.YFilter) { controllers.YFilter = yf }

func (controllers *RedundancyGroupManager_Controllers) GetGoName(yname string) string {
    if yname == "controller" { return "Controller" }
    return ""
}

func (controllers *RedundancyGroupManager_Controllers) GetSegmentPath() string {
    return "controllers"
}

func (controllers *RedundancyGroupManager_Controllers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controller" {
        for _, c := range controllers.Controller {
            if controllers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RedundancyGroupManager_Controllers_Controller{}
        controllers.Controller = append(controllers.Controller, child)
        return &controllers.Controller[len(controllers.Controller)-1]
    }
    return nil
}

func (controllers *RedundancyGroupManager_Controllers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controllers.Controller {
        children[controllers.Controller[i].GetSegmentPath()] = &controllers.Controller[i]
    }
    return children
}

func (controllers *RedundancyGroupManager_Controllers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllers *RedundancyGroupManager_Controllers) GetBundleName() string { return "cisco_ios_xr" }

func (controllers *RedundancyGroupManager_Controllers) GetYangName() string { return "controllers" }

func (controllers *RedundancyGroupManager_Controllers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllers *RedundancyGroupManager_Controllers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllers *RedundancyGroupManager_Controllers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllers *RedundancyGroupManager_Controllers) SetParent(parent types.Entity) { controllers.parent = parent }

func (controllers *RedundancyGroupManager_Controllers) GetParent() types.Entity { return controllers.parent }

func (controllers *RedundancyGroupManager_Controllers) GetParentYangName() string { return "redundancy-group-manager" }

// RedundancyGroupManager_Controllers_Controller
// Display redundancy group by controller name
type RedundancyGroupManager_Controllers_Controller struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // Configured interchassis redundancy group number. The type is string with
    // length: 0..64.
    MultiRouterApsGroupNumber interface{}

    // Name of controller being backed up. The type is string with length: 0..64.
    ControllerNameXr interface{}

    // Handle of controller being backed up. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerHandle interface{}

    // Backup interface name. The type is string with length: 0..64.
    BackupInterfaceName interface{}

    // Backup interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    BackupInterfaceHandle interface{}

    // Backup interface next hop IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BackupInterfaceNextHopIpAddress interface{}

    // Configured interchassis redundancy group state. The type is string with
    // length: 0..64.
    InterChassisGroupState interface{}
}

func (controller *RedundancyGroupManager_Controllers_Controller) GetFilter() yfilter.YFilter { return controller.YFilter }

func (controller *RedundancyGroupManager_Controllers_Controller) SetFilter(yf yfilter.YFilter) { controller.YFilter = yf }

func (controller *RedundancyGroupManager_Controllers_Controller) GetGoName(yname string) string {
    if yname == "controller-name" { return "ControllerName" }
    if yname == "multi-router-aps-group-number" { return "MultiRouterApsGroupNumber" }
    if yname == "controller-name-xr" { return "ControllerNameXr" }
    if yname == "controller-handle" { return "ControllerHandle" }
    if yname == "backup-interface-name" { return "BackupInterfaceName" }
    if yname == "backup-interface-handle" { return "BackupInterfaceHandle" }
    if yname == "backup-interface-next-hop-ip-address" { return "BackupInterfaceNextHopIpAddress" }
    if yname == "inter-chassis-group-state" { return "InterChassisGroupState" }
    return ""
}

func (controller *RedundancyGroupManager_Controllers_Controller) GetSegmentPath() string {
    return "controller" + "[controller-name='" + fmt.Sprintf("%v", controller.ControllerName) + "']"
}

func (controller *RedundancyGroupManager_Controllers_Controller) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (controller *RedundancyGroupManager_Controllers_Controller) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (controller *RedundancyGroupManager_Controllers_Controller) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller-name"] = controller.ControllerName
    leafs["multi-router-aps-group-number"] = controller.MultiRouterApsGroupNumber
    leafs["controller-name-xr"] = controller.ControllerNameXr
    leafs["controller-handle"] = controller.ControllerHandle
    leafs["backup-interface-name"] = controller.BackupInterfaceName
    leafs["backup-interface-handle"] = controller.BackupInterfaceHandle
    leafs["backup-interface-next-hop-ip-address"] = controller.BackupInterfaceNextHopIpAddress
    leafs["inter-chassis-group-state"] = controller.InterChassisGroupState
    return leafs
}

func (controller *RedundancyGroupManager_Controllers_Controller) GetBundleName() string { return "cisco_ios_xr" }

func (controller *RedundancyGroupManager_Controllers_Controller) GetYangName() string { return "controller" }

func (controller *RedundancyGroupManager_Controllers_Controller) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controller *RedundancyGroupManager_Controllers_Controller) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controller *RedundancyGroupManager_Controllers_Controller) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controller *RedundancyGroupManager_Controllers_Controller) SetParent(parent types.Entity) { controller.parent = parent }

func (controller *RedundancyGroupManager_Controllers_Controller) GetParent() types.Entity { return controller.parent }

func (controller *RedundancyGroupManager_Controllers_Controller) GetParentYangName() string { return "controllers" }

