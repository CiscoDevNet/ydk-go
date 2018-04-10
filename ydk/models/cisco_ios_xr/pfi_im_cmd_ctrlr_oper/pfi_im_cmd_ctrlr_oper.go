// This module contains a collection of YANG definitions
// for Cisco IOS-XR pfi-im-cmd-ctrlr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   controllers: Controller operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pfi_im_cmd_ctrlr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pfi_im_cmd_ctrlr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper controllers}", reflect.TypeOf(Controllers{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper:controllers", reflect.TypeOf(Controllers{}))
}

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
)

// Controllers
// Controller operational data
type Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Descriptions for controllers.
    Controllers Controllers_Controllers_
}

func (controllers *Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper"
    controllers.EntityData.SegmentPath = "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper:controllers"
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = make(map[string]types.YChild)
    controllers.EntityData.Children["controllers"] = types.YChild{"Controllers", &controllers.Controllers}
    controllers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(controllers.EntityData)
}

// Controllers_Controllers_
// Descriptions for controllers
type Controllers_Controllers_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Description for a particular controller. The type is slice of
    // Controllers_Controllers__Controller.
    Controller []Controllers_Controllers__Controller
}

func (controllers_ *Controllers_Controllers_) GetEntityData() *types.CommonEntityData {
    controllers_.EntityData.YFilter = controllers_.YFilter
    controllers_.EntityData.YangName = "controllers"
    controllers_.EntityData.BundleName = "cisco_ios_xr"
    controllers_.EntityData.ParentYangName = "controllers"
    controllers_.EntityData.SegmentPath = "controllers"
    controllers_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers_.EntityData.Children = make(map[string]types.YChild)
    controllers_.EntityData.Children["controller"] = types.YChild{"Controller", nil}
    for i := range controllers_.Controller {
        controllers_.EntityData.Children[types.GetSegmentPath(&controllers_.Controller[i])] = types.YChild{"Controller", &controllers_.Controller[i]}
    }
    controllers_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(controllers_.EntityData)
}

// Controllers_Controllers__Controller
// Description for a particular controller
type Controllers_Controllers__Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the controller. The type is string
    // with pattern: b'[a-zA-Z0-9./-]+'.
    InterafceName interface{}

    // Controller. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Controller interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    State interface{}

    // Controller description string. The type is string.
    Description interface{}
}

func (controller *Controllers_Controllers__Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + "[interafce-name='" + fmt.Sprintf("%v", controller.InterafceName) + "']"
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = make(map[string]types.YChild)
    controller.EntityData.Leafs = make(map[string]types.YLeaf)
    controller.EntityData.Leafs["interafce-name"] = types.YLeaf{"InterafceName", controller.InterafceName}
    controller.EntityData.Leafs["controller"] = types.YLeaf{"Controller", controller.Controller}
    controller.EntityData.Leafs["state"] = types.YLeaf{"State", controller.State}
    controller.EntityData.Leafs["description"] = types.YLeaf{"Description", controller.Description}
    return &(controller.EntityData)
}

