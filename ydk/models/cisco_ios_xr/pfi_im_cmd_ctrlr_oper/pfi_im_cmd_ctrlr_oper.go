// This module contains a collection of YANG definitions
// for Cisco IOS-XR pfi-im-cmd-ctrlr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   controllers: Controller operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    Controllers Controllers_Controllers
}

func (controllers *Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper"
    controllers.EntityData.SegmentPath = "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper:controllers"
    controllers.EntityData.AbsolutePath = controllers.EntityData.SegmentPath
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controllers", types.YChild{"Controllers", &controllers.Controllers})
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// Controllers_Controllers
// Descriptions for controllers
type Controllers_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Description for a particular controller. The type is slice of
    // Controllers_Controllers_Controller.
    Controller []*Controllers_Controllers_Controller
}

func (controllers *Controllers_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "controllers"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.AbsolutePath = "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper:controllers/" + controllers.EntityData.SegmentPath
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controller", types.YChild{"Controller", nil})
    for i := range controllers.Controller {
        controllers.EntityData.Children.Append(types.GetSegmentPath(controllers.Controller[i]), types.YChild{"Controller", controllers.Controller[i]})
    }
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// Controllers_Controllers_Controller
// Description for a particular controller
type Controllers_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the controller. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    InterafceName interface{}

    // Controller. The type is string with pattern: [a-zA-Z0-9._/-]+.
    Controller interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    State interface{}

    // Controller description string. The type is string.
    Description interface{}
}

func (controller *Controllers_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.InterafceName, "interafce-name")
    controller.EntityData.AbsolutePath = "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper:controllers/controllers/" + controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("interafce-name", types.YLeaf{"InterafceName", controller.InterafceName})
    controller.EntityData.Leafs.Append("controller", types.YLeaf{"Controller", controller.Controller})
    controller.EntityData.Leafs.Append("state", types.YLeaf{"State", controller.State})
    controller.EntityData.Leafs.Append("description", types.YLeaf{"Description", controller.Description})

    controller.EntityData.YListKeys = []string {"InterafceName"}

    return &(controller.EntityData)
}

