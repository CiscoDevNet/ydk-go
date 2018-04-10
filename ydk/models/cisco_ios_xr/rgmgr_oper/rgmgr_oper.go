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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy group manager data.
    Controllers RedundancyGroupManager_Controllers
}

func (redundancyGroupManager *RedundancyGroupManager) GetEntityData() *types.CommonEntityData {
    redundancyGroupManager.EntityData.YFilter = redundancyGroupManager.YFilter
    redundancyGroupManager.EntityData.YangName = "redundancy-group-manager"
    redundancyGroupManager.EntityData.BundleName = "cisco_ios_xr"
    redundancyGroupManager.EntityData.ParentYangName = "Cisco-IOS-XR-rgmgr-oper"
    redundancyGroupManager.EntityData.SegmentPath = "Cisco-IOS-XR-rgmgr-oper:redundancy-group-manager"
    redundancyGroupManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancyGroupManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancyGroupManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancyGroupManager.EntityData.Children = make(map[string]types.YChild)
    redundancyGroupManager.EntityData.Children["controllers"] = types.YChild{"Controllers", &redundancyGroupManager.Controllers}
    redundancyGroupManager.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(redundancyGroupManager.EntityData)
}

// RedundancyGroupManager_Controllers
// Redundancy group manager data
type RedundancyGroupManager_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Display redundancy group by controller name. The type is slice of
    // RedundancyGroupManager_Controllers_Controller.
    Controller []RedundancyGroupManager_Controllers_Controller
}

func (controllers *RedundancyGroupManager_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "redundancy-group-manager"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = make(map[string]types.YChild)
    controllers.EntityData.Children["controller"] = types.YChild{"Controller", nil}
    for i := range controllers.Controller {
        controllers.EntityData.Children[types.GetSegmentPath(&controllers.Controller[i])] = types.YChild{"Controller", &controllers.Controller[i]}
    }
    controllers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(controllers.EntityData)
}

// RedundancyGroupManager_Controllers_Controller
// Display redundancy group by controller name
type RedundancyGroupManager_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Controller name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    ControllerName interface{}

    // Configured interchassis redundancy group number. The type is string with
    // length: 0..64.
    MultiRouterApsGroupNumber interface{}

    // Name of controller being backed up. The type is string with length: 0..64.
    ControllerNameXr interface{}

    // Handle of controller being backed up. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    ControllerHandle interface{}

    // Backup interface name. The type is string with length: 0..64.
    BackupInterfaceName interface{}

    // Backup interface handle. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    BackupInterfaceHandle interface{}

    // Backup interface next hop IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    BackupInterfaceNextHopIpAddress interface{}

    // Configured interchassis redundancy group state. The type is string with
    // length: 0..64.
    InterChassisGroupState interface{}
}

func (controller *RedundancyGroupManager_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + "[controller-name='" + fmt.Sprintf("%v", controller.ControllerName) + "']"
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = make(map[string]types.YChild)
    controller.EntityData.Leafs = make(map[string]types.YLeaf)
    controller.EntityData.Leafs["controller-name"] = types.YLeaf{"ControllerName", controller.ControllerName}
    controller.EntityData.Leafs["multi-router-aps-group-number"] = types.YLeaf{"MultiRouterApsGroupNumber", controller.MultiRouterApsGroupNumber}
    controller.EntityData.Leafs["controller-name-xr"] = types.YLeaf{"ControllerNameXr", controller.ControllerNameXr}
    controller.EntityData.Leafs["controller-handle"] = types.YLeaf{"ControllerHandle", controller.ControllerHandle}
    controller.EntityData.Leafs["backup-interface-name"] = types.YLeaf{"BackupInterfaceName", controller.BackupInterfaceName}
    controller.EntityData.Leafs["backup-interface-handle"] = types.YLeaf{"BackupInterfaceHandle", controller.BackupInterfaceHandle}
    controller.EntityData.Leafs["backup-interface-next-hop-ip-address"] = types.YLeaf{"BackupInterfaceNextHopIpAddress", controller.BackupInterfaceNextHopIpAddress}
    controller.EntityData.Leafs["inter-chassis-group-state"] = types.YLeaf{"InterChassisGroupState", controller.InterChassisGroupState}
    return &(controller.EntityData)
}

