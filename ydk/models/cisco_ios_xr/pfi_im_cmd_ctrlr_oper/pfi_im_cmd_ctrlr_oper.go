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
    parent types.Entity
    YFilter yfilter.YFilter

    // Descriptions for controllers.
    Controllers Controllers_Controllers
}

func (controllers *Controllers) GetFilter() yfilter.YFilter { return controllers.YFilter }

func (controllers *Controllers) SetFilter(yf yfilter.YFilter) { controllers.YFilter = yf }

func (controllers *Controllers) GetGoName(yname string) string {
    if yname == "controllers" { return "Controllers" }
    return ""
}

func (controllers *Controllers) GetSegmentPath() string {
    return "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper:controllers"
}

func (controllers *Controllers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controllers" {
        return &controllers.Controllers
    }
    return nil
}

func (controllers *Controllers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["controllers"] = &controllers.Controllers
    return children
}

func (controllers *Controllers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllers *Controllers) GetBundleName() string { return "cisco_ios_xr" }

func (controllers *Controllers) GetYangName() string { return "controllers" }

func (controllers *Controllers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllers *Controllers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllers *Controllers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllers *Controllers) SetParent(parent types.Entity) { controllers.parent = parent }

func (controllers *Controllers) GetParent() types.Entity { return controllers.parent }

func (controllers *Controllers) GetParentYangName() string { return "Cisco-IOS-XR-pfi-im-cmd-ctrlr-oper" }

// Controllers_Controllers
// Descriptions for controllers
type Controllers_Controllers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Description for a particular controller. The type is slice of
    // Controllers_Controllers_Controller.
    Controller []Controllers_Controllers_Controller
}

func (controllers *Controllers_Controllers) GetFilter() yfilter.YFilter { return controllers.YFilter }

func (controllers *Controllers_Controllers) SetFilter(yf yfilter.YFilter) { controllers.YFilter = yf }

func (controllers *Controllers_Controllers) GetGoName(yname string) string {
    if yname == "controller" { return "Controller" }
    return ""
}

func (controllers *Controllers_Controllers) GetSegmentPath() string {
    return "controllers"
}

func (controllers *Controllers_Controllers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controller" {
        for _, c := range controllers.Controller {
            if controllers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controllers_Controllers_Controller{}
        controllers.Controller = append(controllers.Controller, child)
        return &controllers.Controller[len(controllers.Controller)-1]
    }
    return nil
}

func (controllers *Controllers_Controllers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controllers.Controller {
        children[controllers.Controller[i].GetSegmentPath()] = &controllers.Controller[i]
    }
    return children
}

func (controllers *Controllers_Controllers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllers *Controllers_Controllers) GetBundleName() string { return "cisco_ios_xr" }

func (controllers *Controllers_Controllers) GetYangName() string { return "controllers" }

func (controllers *Controllers_Controllers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllers *Controllers_Controllers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllers *Controllers_Controllers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllers *Controllers_Controllers) SetParent(parent types.Entity) { controllers.parent = parent }

func (controllers *Controllers_Controllers) GetParent() types.Entity { return controllers.parent }

func (controllers *Controllers_Controllers) GetParentYangName() string { return "controllers" }

// Controllers_Controllers_Controller
// Description for a particular controller
type Controllers_Controllers_Controller struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the controller. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    InterafceName interface{}

    // Controller. The type is string with pattern: [a-zA-Z0-9./-]+.
    Controller interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    State interface{}

    // Controller description string. The type is string.
    Description interface{}
}

func (controller *Controllers_Controllers_Controller) GetFilter() yfilter.YFilter { return controller.YFilter }

func (controller *Controllers_Controllers_Controller) SetFilter(yf yfilter.YFilter) { controller.YFilter = yf }

func (controller *Controllers_Controllers_Controller) GetGoName(yname string) string {
    if yname == "interafce-name" { return "InterafceName" }
    if yname == "controller" { return "Controller" }
    if yname == "state" { return "State" }
    if yname == "description" { return "Description" }
    return ""
}

func (controller *Controllers_Controllers_Controller) GetSegmentPath() string {
    return "controller" + "[interafce-name='" + fmt.Sprintf("%v", controller.InterafceName) + "']"
}

func (controller *Controllers_Controllers_Controller) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (controller *Controllers_Controllers_Controller) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (controller *Controllers_Controllers_Controller) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interafce-name"] = controller.InterafceName
    leafs["controller"] = controller.Controller
    leafs["state"] = controller.State
    leafs["description"] = controller.Description
    return leafs
}

func (controller *Controllers_Controllers_Controller) GetBundleName() string { return "cisco_ios_xr" }

func (controller *Controllers_Controllers_Controller) GetYangName() string { return "controller" }

func (controller *Controllers_Controllers_Controller) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controller *Controllers_Controllers_Controller) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controller *Controllers_Controllers_Controller) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controller *Controllers_Controllers_Controller) SetParent(parent types.Entity) { controller.parent = parent }

func (controller *Controllers_Controllers_Controller) GetParent() types.Entity { return controller.parent }

func (controller *Controllers_Controllers_Controller) GetParentYangName() string { return "controllers" }

