// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-rsvp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   rsvp: Global RSVP configuration commands
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_rsvp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_rsvp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-rsvp-cfg rsvp}", reflect.TypeOf(Rsvp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-rsvp-cfg:rsvp", reflect.TypeOf(Rsvp{}))
}

// RsvpBwCfg represents Rsvp bw cfg
type RsvpBwCfg string

const (
    // Configuration is in absolute bandwidth values
    RsvpBwCfg_absolute RsvpBwCfg = "absolute"

    // Configuration is in percentage of physical
    // bandwidth values
    RsvpBwCfg_percentage RsvpBwCfg = "percentage"
)

// RsvpBc0 represents Rsvp bc0
type RsvpBc0 string

const (
    // Keyword is bc0
    RsvpBc0_bc0 RsvpBc0 = "bc0"

    // Keyword is global-pool
    RsvpBc0_global_pool RsvpBc0 = "global-pool"

    // Keyword is not specified
    RsvpBc0_not_specified RsvpBc0 = "not-specified"
)

// RsvpBc1 represents Rsvp bc1
type RsvpBc1 string

const (
    // Keyword is bc1
    RsvpBc1_bc1 RsvpBc1 = "bc1"

    // Keyword is sub-pool
    RsvpBc1_sub_pool RsvpBc1 = "sub-pool"
)

// RsvpRdm represents Rsvp rdm
type RsvpRdm string

const (
    // RDM Keyword Specified
    RsvpRdm_rdm RsvpRdm = "rdm"

    // RDM Keyword Not Specified
    RsvpRdm_not_specified RsvpRdm = "not-specified"

    // Use Default Bandwidth - 75% Link Bandwidth
    RsvpRdm_use_default_bandwidth RsvpRdm = "use-default-bandwidth"
)

// Rsvp
// Global RSVP configuration commands
type Rsvp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP Neighbor Table.
    Neighbors Rsvp_Neighbors

    // Controller table.
    Controllers Rsvp_Controllers

    // Global Logging.
    GlobalLogging Rsvp_GlobalLogging

    // Configure Global Bandwidth Parameters.
    GlobalBandwidth Rsvp_GlobalBandwidth

    // Interface table.
    Interfaces Rsvp_Interfaces

    // Configure Global RSVP signalling parameters.
    Signalling Rsvp_Signalling

    // Configure RSVP authentication.
    Authentication Rsvp_Authentication
}

func (rsvp *Rsvp) GetFilter() yfilter.YFilter { return rsvp.YFilter }

func (rsvp *Rsvp) SetFilter(yf yfilter.YFilter) { rsvp.YFilter = yf }

func (rsvp *Rsvp) GetGoName(yname string) string {
    if yname == "neighbors" { return "Neighbors" }
    if yname == "controllers" { return "Controllers" }
    if yname == "global-logging" { return "GlobalLogging" }
    if yname == "global-bandwidth" { return "GlobalBandwidth" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "signalling" { return "Signalling" }
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (rsvp *Rsvp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-rsvp-cfg:rsvp"
}

func (rsvp *Rsvp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &rsvp.Neighbors
    }
    if childYangName == "controllers" {
        return &rsvp.Controllers
    }
    if childYangName == "global-logging" {
        return &rsvp.GlobalLogging
    }
    if childYangName == "global-bandwidth" {
        return &rsvp.GlobalBandwidth
    }
    if childYangName == "interfaces" {
        return &rsvp.Interfaces
    }
    if childYangName == "signalling" {
        return &rsvp.Signalling
    }
    if childYangName == "authentication" {
        return &rsvp.Authentication
    }
    return nil
}

func (rsvp *Rsvp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &rsvp.Neighbors
    children["controllers"] = &rsvp.Controllers
    children["global-logging"] = &rsvp.GlobalLogging
    children["global-bandwidth"] = &rsvp.GlobalBandwidth
    children["interfaces"] = &rsvp.Interfaces
    children["signalling"] = &rsvp.Signalling
    children["authentication"] = &rsvp.Authentication
    return children
}

func (rsvp *Rsvp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvp *Rsvp) GetBundleName() string { return "cisco_ios_xr" }

func (rsvp *Rsvp) GetYangName() string { return "rsvp" }

func (rsvp *Rsvp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rsvp *Rsvp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rsvp *Rsvp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rsvp *Rsvp) SetParent(parent types.Entity) { rsvp.parent = parent }

func (rsvp *Rsvp) GetParent() types.Entity { return rsvp.parent }

func (rsvp *Rsvp) GetParentYangName() string { return "Cisco-IOS-XR-ip-rsvp-cfg" }

// Rsvp_Neighbors
// RSVP Neighbor Table
type Rsvp_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP neighbor configuration. The type is slice of Rsvp_Neighbors_Neighbor.
    Neighbor []Rsvp_Neighbors_Neighbor
}

func (neighbors *Rsvp_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Rsvp_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Rsvp_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Rsvp_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Rsvp_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rsvp_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Rsvp_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Rsvp_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Rsvp_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Rsvp_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Rsvp_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Rsvp_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Rsvp_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Rsvp_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Rsvp_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Rsvp_Neighbors) GetParentYangName() string { return "rsvp" }

// Rsvp_Neighbors_Neighbor
// RSVP neighbor configuration
type Rsvp_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // Configure RSVP authentication.
    Authentication Rsvp_Neighbors_Neighbor_Authentication
}

func (neighbor *Rsvp_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Rsvp_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Rsvp_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (neighbor *Rsvp_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor='" + fmt.Sprintf("%v", neighbor.Neighbor) + "']"
}

func (neighbor *Rsvp_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        return &neighbor.Authentication
    }
    return nil
}

func (neighbor *Rsvp_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authentication"] = &neighbor.Authentication
    return children
}

func (neighbor *Rsvp_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor"] = neighbor.Neighbor
    return leafs
}

func (neighbor *Rsvp_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *Rsvp_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Rsvp_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *Rsvp_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *Rsvp_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *Rsvp_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Rsvp_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Rsvp_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Rsvp_Neighbors_Neighbor_Authentication
// Configure RSVP authentication
type Rsvp_Neighbors_Neighbor_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Life time (in seconds) for each security association. The type is
    // interface{} with range: 30..86400. Units are second.
    LifeTime interface{}

    // Enable or disable RSVP authentication. The type is bool.
    Enable interface{}

    // Window-size to limit number of out-of-order messages. The type is
    // interface{} with range: 1..64.
    WindowSize interface{}

    // Key chain to authenticate RSVP signalling messages. The type is string with
    // length: 1..32.
    KeyChain interface{}
}

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetGoName(yname string) string {
    if yname == "life-time" { return "LifeTime" }
    if yname == "enable" { return "Enable" }
    if yname == "window-size" { return "WindowSize" }
    if yname == "key-chain" { return "KeyChain" }
    return ""
}

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["life-time"] = authentication.LifeTime
    leafs["enable"] = authentication.Enable
    leafs["window-size"] = authentication.WindowSize
    leafs["key-chain"] = authentication.KeyChain
    return leafs
}

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetYangName() string { return "authentication" }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetParentYangName() string { return "neighbor" }

// Rsvp_Controllers
// Controller table
type Rsvp_Controllers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Controller configuration. The type is slice of Rsvp_Controllers_Controller.
    Controller []Rsvp_Controllers_Controller
}

func (controllers *Rsvp_Controllers) GetFilter() yfilter.YFilter { return controllers.YFilter }

func (controllers *Rsvp_Controllers) SetFilter(yf yfilter.YFilter) { controllers.YFilter = yf }

func (controllers *Rsvp_Controllers) GetGoName(yname string) string {
    if yname == "controller" { return "Controller" }
    return ""
}

func (controllers *Rsvp_Controllers) GetSegmentPath() string {
    return "controllers"
}

func (controllers *Rsvp_Controllers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controller" {
        for _, c := range controllers.Controller {
            if controllers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rsvp_Controllers_Controller{}
        controllers.Controller = append(controllers.Controller, child)
        return &controllers.Controller[len(controllers.Controller)-1]
    }
    return nil
}

func (controllers *Rsvp_Controllers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controllers.Controller {
        children[controllers.Controller[i].GetSegmentPath()] = &controllers.Controller[i]
    }
    return children
}

func (controllers *Rsvp_Controllers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllers *Rsvp_Controllers) GetBundleName() string { return "cisco_ios_xr" }

func (controllers *Rsvp_Controllers) GetYangName() string { return "controllers" }

func (controllers *Rsvp_Controllers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllers *Rsvp_Controllers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllers *Rsvp_Controllers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllers *Rsvp_Controllers) SetParent(parent types.Entity) { controllers.parent = parent }

func (controllers *Rsvp_Controllers) GetParent() types.Entity { return controllers.parent }

func (controllers *Rsvp_Controllers) GetParentYangName() string { return "rsvp" }

// Rsvp_Controllers_Controller
// Controller configuration
type Rsvp_Controllers_Controller struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of controller. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // Enable RSVP on an interface. The type is interface{}.
    Enable interface{}

    // Configure RSVP signalling parameters.
    CntlSignalling Rsvp_Controllers_Controller_CntlSignalling
}

func (controller *Rsvp_Controllers_Controller) GetFilter() yfilter.YFilter { return controller.YFilter }

func (controller *Rsvp_Controllers_Controller) SetFilter(yf yfilter.YFilter) { controller.YFilter = yf }

func (controller *Rsvp_Controllers_Controller) GetGoName(yname string) string {
    if yname == "controller-name" { return "ControllerName" }
    if yname == "enable" { return "Enable" }
    if yname == "cntl-signalling" { return "CntlSignalling" }
    return ""
}

func (controller *Rsvp_Controllers_Controller) GetSegmentPath() string {
    return "controller" + "[controller-name='" + fmt.Sprintf("%v", controller.ControllerName) + "']"
}

func (controller *Rsvp_Controllers_Controller) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cntl-signalling" {
        return &controller.CntlSignalling
    }
    return nil
}

func (controller *Rsvp_Controllers_Controller) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cntl-signalling"] = &controller.CntlSignalling
    return children
}

func (controller *Rsvp_Controllers_Controller) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller-name"] = controller.ControllerName
    leafs["enable"] = controller.Enable
    return leafs
}

func (controller *Rsvp_Controllers_Controller) GetBundleName() string { return "cisco_ios_xr" }

func (controller *Rsvp_Controllers_Controller) GetYangName() string { return "controller" }

func (controller *Rsvp_Controllers_Controller) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controller *Rsvp_Controllers_Controller) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controller *Rsvp_Controllers_Controller) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controller *Rsvp_Controllers_Controller) SetParent(parent types.Entity) { controller.parent = parent }

func (controller *Rsvp_Controllers_Controller) GetParent() types.Entity { return controller.parent }

func (controller *Rsvp_Controllers_Controller) GetParentYangName() string { return "controllers" }

// Rsvp_Controllers_Controller_CntlSignalling
// Configure RSVP signalling parameters
type Rsvp_Controllers_Controller_CntlSignalling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure RSVP out-of-band signalling parameters.
    OutOfBand Rsvp_Controllers_Controller_CntlSignalling_OutOfBand
}

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetFilter() yfilter.YFilter { return cntlSignalling.YFilter }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) SetFilter(yf yfilter.YFilter) { cntlSignalling.YFilter = yf }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetGoName(yname string) string {
    if yname == "out-of-band" { return "OutOfBand" }
    return ""
}

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetSegmentPath() string {
    return "cntl-signalling"
}

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "out-of-band" {
        return &cntlSignalling.OutOfBand
    }
    return nil
}

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["out-of-band"] = &cntlSignalling.OutOfBand
    return children
}

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetBundleName() string { return "cisco_ios_xr" }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetYangName() string { return "cntl-signalling" }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) SetParent(parent types.Entity) { cntlSignalling.parent = parent }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetParent() types.Entity { return cntlSignalling.parent }

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetParentYangName() string { return "controller" }

// Rsvp_Controllers_Controller_CntlSignalling_OutOfBand
// Configure RSVP out-of-band signalling parameters
type Rsvp_Controllers_Controller_CntlSignalling_OutOfBand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure max number of consecutive missed messages for state expiry for
    // out-of-band tunnels. The type is interface{} with range: 1..110000. The
    // default value is 38000.
    MissedMessages interface{}

    // Configure interval between successive refreshes for out-of-band tunnels.
    // The type is interface{} with range: 180..86400. Units are second.
    RefreshInterval interface{}
}

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetFilter() yfilter.YFilter { return outOfBand.YFilter }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) SetFilter(yf yfilter.YFilter) { outOfBand.YFilter = yf }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetGoName(yname string) string {
    if yname == "missed-messages" { return "MissedMessages" }
    if yname == "refresh-interval" { return "RefreshInterval" }
    return ""
}

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetSegmentPath() string {
    return "out-of-band"
}

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["missed-messages"] = outOfBand.MissedMessages
    leafs["refresh-interval"] = outOfBand.RefreshInterval
    return leafs
}

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetBundleName() string { return "cisco_ios_xr" }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetYangName() string { return "out-of-band" }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) SetParent(parent types.Entity) { outOfBand.parent = parent }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetParent() types.Entity { return outOfBand.parent }

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetParentYangName() string { return "cntl-signalling" }

// Rsvp_GlobalLogging
// Global Logging
type Rsvp_GlobalLogging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable NSR Status Logging. The type is interface{}.
    LogNsrStatus interface{}

    // Enable ISSU Status Logging. The type is interface{}.
    LogIssuStatus interface{}
}

func (globalLogging *Rsvp_GlobalLogging) GetFilter() yfilter.YFilter { return globalLogging.YFilter }

func (globalLogging *Rsvp_GlobalLogging) SetFilter(yf yfilter.YFilter) { globalLogging.YFilter = yf }

func (globalLogging *Rsvp_GlobalLogging) GetGoName(yname string) string {
    if yname == "log-nsr-status" { return "LogNsrStatus" }
    if yname == "log-issu-status" { return "LogIssuStatus" }
    return ""
}

func (globalLogging *Rsvp_GlobalLogging) GetSegmentPath() string {
    return "global-logging"
}

func (globalLogging *Rsvp_GlobalLogging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalLogging *Rsvp_GlobalLogging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalLogging *Rsvp_GlobalLogging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-nsr-status"] = globalLogging.LogNsrStatus
    leafs["log-issu-status"] = globalLogging.LogIssuStatus
    return leafs
}

func (globalLogging *Rsvp_GlobalLogging) GetBundleName() string { return "cisco_ios_xr" }

func (globalLogging *Rsvp_GlobalLogging) GetYangName() string { return "global-logging" }

func (globalLogging *Rsvp_GlobalLogging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalLogging *Rsvp_GlobalLogging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalLogging *Rsvp_GlobalLogging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalLogging *Rsvp_GlobalLogging) SetParent(parent types.Entity) { globalLogging.parent = parent }

func (globalLogging *Rsvp_GlobalLogging) GetParent() types.Entity { return globalLogging.parent }

func (globalLogging *Rsvp_GlobalLogging) GetParentYangName() string { return "rsvp" }

// Rsvp_GlobalBandwidth
// Configure Global Bandwidth Parameters
type Rsvp_GlobalBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Global RSVP signalling parameters.
    DefaultInterfacePercent Rsvp_GlobalBandwidth_DefaultInterfacePercent
}

func (globalBandwidth *Rsvp_GlobalBandwidth) GetFilter() yfilter.YFilter { return globalBandwidth.YFilter }

func (globalBandwidth *Rsvp_GlobalBandwidth) SetFilter(yf yfilter.YFilter) { globalBandwidth.YFilter = yf }

func (globalBandwidth *Rsvp_GlobalBandwidth) GetGoName(yname string) string {
    if yname == "default-interface-percent" { return "DefaultInterfacePercent" }
    return ""
}

func (globalBandwidth *Rsvp_GlobalBandwidth) GetSegmentPath() string {
    return "global-bandwidth"
}

func (globalBandwidth *Rsvp_GlobalBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-interface-percent" {
        return &globalBandwidth.DefaultInterfacePercent
    }
    return nil
}

func (globalBandwidth *Rsvp_GlobalBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-interface-percent"] = &globalBandwidth.DefaultInterfacePercent
    return children
}

func (globalBandwidth *Rsvp_GlobalBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalBandwidth *Rsvp_GlobalBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (globalBandwidth *Rsvp_GlobalBandwidth) GetYangName() string { return "global-bandwidth" }

func (globalBandwidth *Rsvp_GlobalBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalBandwidth *Rsvp_GlobalBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalBandwidth *Rsvp_GlobalBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalBandwidth *Rsvp_GlobalBandwidth) SetParent(parent types.Entity) { globalBandwidth.parent = parent }

func (globalBandwidth *Rsvp_GlobalBandwidth) GetParent() types.Entity { return globalBandwidth.parent }

func (globalBandwidth *Rsvp_GlobalBandwidth) GetParentYangName() string { return "rsvp" }

// Rsvp_GlobalBandwidth_DefaultInterfacePercent
// Configure Global RSVP signalling parameters
type Rsvp_GlobalBandwidth_DefaultInterfacePercent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure global default MAM I/F percent bandwidth parameters.
    Mam Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam

    // Configure global default RDM I/F percent bandwidth parameters.
    Rdm Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm
}

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetFilter() yfilter.YFilter { return defaultInterfacePercent.YFilter }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) SetFilter(yf yfilter.YFilter) { defaultInterfacePercent.YFilter = yf }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetGoName(yname string) string {
    if yname == "mam" { return "Mam" }
    if yname == "rdm" { return "Rdm" }
    return ""
}

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetSegmentPath() string {
    return "default-interface-percent"
}

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mam" {
        return &defaultInterfacePercent.Mam
    }
    if childYangName == "rdm" {
        return &defaultInterfacePercent.Rdm
    }
    return nil
}

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mam"] = &defaultInterfacePercent.Mam
    children["rdm"] = &defaultInterfacePercent.Rdm
    return children
}

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetBundleName() string { return "cisco_ios_xr" }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetYangName() string { return "default-interface-percent" }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) SetParent(parent types.Entity) { defaultInterfacePercent.parent = parent }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetParent() types.Entity { return defaultInterfacePercent.parent }

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetParentYangName() string { return "global-bandwidth" }

// Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam
// Configure global default MAM I/F percent
// bandwidth parameters
type Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default maximum reservable I/F % B/W . The type is interface{} with range:
    // 0..10000.
    MaxResPercent interface{}

    // Default BC0 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc0Percent interface{}

    // Default BC1 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc1Percent interface{}
}

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetFilter() yfilter.YFilter { return mam.YFilter }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) SetFilter(yf yfilter.YFilter) { mam.YFilter = yf }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetGoName(yname string) string {
    if yname == "max-res-percent" { return "MaxResPercent" }
    if yname == "bc0-percent" { return "Bc0Percent" }
    if yname == "bc1-percent" { return "Bc1Percent" }
    return ""
}

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetSegmentPath() string {
    return "mam"
}

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-res-percent"] = mam.MaxResPercent
    leafs["bc0-percent"] = mam.Bc0Percent
    leafs["bc1-percent"] = mam.Bc1Percent
    return leafs
}

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetBundleName() string { return "cisco_ios_xr" }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetYangName() string { return "mam" }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) SetParent(parent types.Entity) { mam.parent = parent }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetParent() types.Entity { return mam.parent }

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetParentYangName() string { return "default-interface-percent" }

// Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm
// Configure global default RDM I/F percent
// bandwidth parameters
type Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default BC0 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc0Percent interface{}

    // Default BC1 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc1Percent interface{}
}

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetFilter() yfilter.YFilter { return rdm.YFilter }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) SetFilter(yf yfilter.YFilter) { rdm.YFilter = yf }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetGoName(yname string) string {
    if yname == "bc0-percent" { return "Bc0Percent" }
    if yname == "bc1-percent" { return "Bc1Percent" }
    return ""
}

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetSegmentPath() string {
    return "rdm"
}

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bc0-percent"] = rdm.Bc0Percent
    leafs["bc1-percent"] = rdm.Bc1Percent
    return leafs
}

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetBundleName() string { return "cisco_ios_xr" }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetYangName() string { return "rdm" }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) SetParent(parent types.Entity) { rdm.parent = parent }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetParent() types.Entity { return rdm.parent }

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetParentYangName() string { return "default-interface-percent" }

// Rsvp_Interfaces
// Interface table
type Rsvp_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface configuration. The type is slice of Rsvp_Interfaces_Interface.
    Interface []Rsvp_Interfaces_Interface
}

func (interfaces *Rsvp_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Rsvp_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Rsvp_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Rsvp_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Rsvp_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rsvp_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Rsvp_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Rsvp_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Rsvp_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Rsvp_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Rsvp_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Rsvp_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Rsvp_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Rsvp_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Rsvp_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Rsvp_Interfaces) GetParentYangName() string { return "rsvp" }

// Rsvp_Interfaces_Interface
// Interface configuration
type Rsvp_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    Name interface{}

    // Enable RSVP on an interface. The type is interface{}.
    Enable interface{}

    // Configure RSVP signalling parameters.
    IfSignalling Rsvp_Interfaces_Interface_IfSignalling

    // Configure Bandwidth.
    Bandwidth Rsvp_Interfaces_Interface_Bandwidth

    // Configure RSVP authentication.
    Authentication Rsvp_Interfaces_Interface_Authentication
}

func (self *Rsvp_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Rsvp_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Rsvp_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "enable" { return "Enable" }
    if yname == "if-signalling" { return "IfSignalling" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (self *Rsvp_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Rsvp_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "if-signalling" {
        return &self.IfSignalling
    }
    if childYangName == "bandwidth" {
        return &self.Bandwidth
    }
    if childYangName == "authentication" {
        return &self.Authentication
    }
    return nil
}

func (self *Rsvp_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["if-signalling"] = &self.IfSignalling
    children["bandwidth"] = &self.Bandwidth
    children["authentication"] = &self.Authentication
    return children
}

func (self *Rsvp_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["enable"] = self.Enable
    return leafs
}

func (self *Rsvp_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Rsvp_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Rsvp_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Rsvp_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Rsvp_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Rsvp_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Rsvp_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Rsvp_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Rsvp_Interfaces_Interface_IfSignalling
// Configure RSVP signalling parameters
type Rsvp_Interfaces_Interface_IfSignalling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Differentiated Services Code Point (DSCP). The type is interface{} with
    // range: 0..63.
    Dscp interface{}

    // Configure max number of consecutive missed messages for state expiry. The
    // type is interface{} with range: 1..8. The default value is 4.
    MissedMessages interface{}

    // Enable IF-based Hello adjacency on a RSVP interface. The type is
    // interface{}.
    HelloGracefulRestartIfBased interface{}

    // Enable rate-limiting on the interface. The type is interface{}.
    Pacing interface{}

    // Configure interval between successive refreshes. The type is interface{}
    // with range: 10..180. Units are second. The default value is 45.
    RefreshInterval interface{}

    // Configure RSVP Refresh Reduction parameters.
    RefreshReduction Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction

    // Configure number of messages to be sent per interval.
    IntervalRate Rsvp_Interfaces_Interface_IfSignalling_IntervalRate

    // Configure RSVP out-of-band signalling parameters.
    OutOfBand Rsvp_Interfaces_Interface_IfSignalling_OutOfBand
}

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetFilter() yfilter.YFilter { return ifSignalling.YFilter }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) SetFilter(yf yfilter.YFilter) { ifSignalling.YFilter = yf }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    if yname == "missed-messages" { return "MissedMessages" }
    if yname == "hello-graceful-restart-if-based" { return "HelloGracefulRestartIfBased" }
    if yname == "pacing" { return "Pacing" }
    if yname == "refresh-interval" { return "RefreshInterval" }
    if yname == "refresh-reduction" { return "RefreshReduction" }
    if yname == "interval-rate" { return "IntervalRate" }
    if yname == "out-of-band" { return "OutOfBand" }
    return ""
}

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetSegmentPath() string {
    return "if-signalling"
}

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "refresh-reduction" {
        return &ifSignalling.RefreshReduction
    }
    if childYangName == "interval-rate" {
        return &ifSignalling.IntervalRate
    }
    if childYangName == "out-of-band" {
        return &ifSignalling.OutOfBand
    }
    return nil
}

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["refresh-reduction"] = &ifSignalling.RefreshReduction
    children["interval-rate"] = &ifSignalling.IntervalRate
    children["out-of-band"] = &ifSignalling.OutOfBand
    return children
}

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = ifSignalling.Dscp
    leafs["missed-messages"] = ifSignalling.MissedMessages
    leafs["hello-graceful-restart-if-based"] = ifSignalling.HelloGracefulRestartIfBased
    leafs["pacing"] = ifSignalling.Pacing
    leafs["refresh-interval"] = ifSignalling.RefreshInterval
    return leafs
}

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetBundleName() string { return "cisco_ios_xr" }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetYangName() string { return "if-signalling" }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) SetParent(parent types.Entity) { ifSignalling.parent = parent }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetParent() types.Entity { return ifSignalling.parent }

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetParentYangName() string { return "interface" }

// Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction
// Configure RSVP Refresh Reduction parameters
type Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable refresh reduction. The type is interface{}.
    Disable interface{}

    // Configure max size of a single RSVP ACK message. The type is interface{}
    // with range: 20..65000. Units are byte. The default value is 4096.
    ReliableAckMaxSize interface{}

    // Configure hold time for sending RSVP ACK message(s). The type is
    // interface{} with range: 100..5000. Units are millisecond. The default value
    // is 400.
    ReliableAckHoldTime interface{}

    // Configure min delay to wait for an ACK before a retransmit. The type is
    // interface{} with range: 100..10000. Units are millisecond. The default
    // value is 2100.
    ReliableRetransmitTime interface{}

    // Configure use of reliable messaging for summary refresh. The type is
    // interface{}.
    ReliableSRefresh interface{}

    // Configure max size of a single RSVP summary refresh message. The type is
    // interface{} with range: 20..65000. Units are byte. The default value is
    // 4096.
    SummaryMaxSize interface{}

    // Configure maximum size of a single RSVP Bundle message. The type is
    // interface{} with range: 512..65000. Units are byte. The default value is
    // 4096.
    BundleMessageMaxSize interface{}
}

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetFilter() yfilter.YFilter { return refreshReduction.YFilter }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) SetFilter(yf yfilter.YFilter) { refreshReduction.YFilter = yf }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "reliable-ack-max-size" { return "ReliableAckMaxSize" }
    if yname == "reliable-ack-hold-time" { return "ReliableAckHoldTime" }
    if yname == "reliable-retransmit-time" { return "ReliableRetransmitTime" }
    if yname == "reliable-s-refresh" { return "ReliableSRefresh" }
    if yname == "summary-max-size" { return "SummaryMaxSize" }
    if yname == "bundle-message-max-size" { return "BundleMessageMaxSize" }
    return ""
}

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetSegmentPath() string {
    return "refresh-reduction"
}

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = refreshReduction.Disable
    leafs["reliable-ack-max-size"] = refreshReduction.ReliableAckMaxSize
    leafs["reliable-ack-hold-time"] = refreshReduction.ReliableAckHoldTime
    leafs["reliable-retransmit-time"] = refreshReduction.ReliableRetransmitTime
    leafs["reliable-s-refresh"] = refreshReduction.ReliableSRefresh
    leafs["summary-max-size"] = refreshReduction.SummaryMaxSize
    leafs["bundle-message-max-size"] = refreshReduction.BundleMessageMaxSize
    return leafs
}

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetBundleName() string { return "cisco_ios_xr" }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetYangName() string { return "refresh-reduction" }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) SetParent(parent types.Entity) { refreshReduction.parent = parent }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetParent() types.Entity { return refreshReduction.parent }

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetParentYangName() string { return "if-signalling" }

// Rsvp_Interfaces_Interface_IfSignalling_IntervalRate
// Configure number of messages to be sent per
// interval
type Rsvp_Interfaces_Interface_IfSignalling_IntervalRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of messages to be sent per interval. The type is interface{} with
    // range: 1..500. The default value is 100.
    MessagesPerInterval interface{}

    // Size of an interval (milliseconds). The type is interface{} with range:
    // 250..2000. Units are millisecond. The default value is 1000.
    IntervalSize interface{}
}

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetFilter() yfilter.YFilter { return intervalRate.YFilter }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) SetFilter(yf yfilter.YFilter) { intervalRate.YFilter = yf }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetGoName(yname string) string {
    if yname == "messages-per-interval" { return "MessagesPerInterval" }
    if yname == "interval-size" { return "IntervalSize" }
    return ""
}

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetSegmentPath() string {
    return "interval-rate"
}

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["messages-per-interval"] = intervalRate.MessagesPerInterval
    leafs["interval-size"] = intervalRate.IntervalSize
    return leafs
}

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetBundleName() string { return "cisco_ios_xr" }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetYangName() string { return "interval-rate" }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) SetParent(parent types.Entity) { intervalRate.parent = parent }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetParent() types.Entity { return intervalRate.parent }

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetParentYangName() string { return "if-signalling" }

// Rsvp_Interfaces_Interface_IfSignalling_OutOfBand
// Configure RSVP out-of-band signalling parameters
type Rsvp_Interfaces_Interface_IfSignalling_OutOfBand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure max number of consecutive missed messages for state expiry for
    // out-of-band tunnels. The type is interface{} with range: 1..110000. The
    // default value is 38000.
    MissedMessages interface{}

    // Configure interval between successive refreshes for out-of-band tunnels.
    // The type is interface{} with range: 180..86400. Units are second.
    RefreshInterval interface{}
}

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetFilter() yfilter.YFilter { return outOfBand.YFilter }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) SetFilter(yf yfilter.YFilter) { outOfBand.YFilter = yf }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetGoName(yname string) string {
    if yname == "missed-messages" { return "MissedMessages" }
    if yname == "refresh-interval" { return "RefreshInterval" }
    return ""
}

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetSegmentPath() string {
    return "out-of-band"
}

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["missed-messages"] = outOfBand.MissedMessages
    leafs["refresh-interval"] = outOfBand.RefreshInterval
    return leafs
}

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetBundleName() string { return "cisco_ios_xr" }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetYangName() string { return "out-of-band" }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) SetParent(parent types.Entity) { outOfBand.parent = parent }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetParent() types.Entity { return outOfBand.parent }

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetParentYangName() string { return "if-signalling" }

// Rsvp_Interfaces_Interface_Bandwidth
// Configure Bandwidth
type Rsvp_Interfaces_Interface_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure MAM bandwidth parameters.
    Mam Rsvp_Interfaces_Interface_Bandwidth_Mam

    // Configure RDM bandwidth parameters.
    Rdm Rsvp_Interfaces_Interface_Bandwidth_Rdm
}

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetGoName(yname string) string {
    if yname == "mam" { return "Mam" }
    if yname == "rdm" { return "Rdm" }
    return ""
}

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mam" {
        return &bandwidth.Mam
    }
    if childYangName == "rdm" {
        return &bandwidth.Rdm
    }
    return nil
}

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mam"] = &bandwidth.Mam
    children["rdm"] = &bandwidth.Rdm
    return children
}

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetParentYangName() string { return "interface" }

// Rsvp_Interfaces_Interface_Bandwidth_Mam
// Configure MAM bandwidth parameters
type Rsvp_Interfaces_Interface_Bandwidth_Mam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum reservable bandwidth (Kbps or percent of physical bandwidth). The
    // type is interface{} with range: 0..4294967295.
    MaxResvBandwidth interface{}

    // Largest reservable flow (Kbps or percent of physical bandwidth). The type
    // is interface{} with range: 0..4294967295.
    MaxResvFlow interface{}

    // Reservable bandwidth in BC0 (Kbps or percent of physical bandwidth). The
    // type is interface{} with range: 0..4294967295.
    Bc0Bandwidth interface{}

    // Reservable bandwidth in BC1 (Kbps or percent of physical bandwidth). The
    // type is interface{} with range: 0..4294967295.
    Bc1Bandwidth interface{}

    // Absolute or Percentage bandwidth mode. The type is RsvpBwCfg. Units are
    // percentage.
    BandwidthMode interface{}
}

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetFilter() yfilter.YFilter { return mam.YFilter }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) SetFilter(yf yfilter.YFilter) { mam.YFilter = yf }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetGoName(yname string) string {
    if yname == "max-resv-bandwidth" { return "MaxResvBandwidth" }
    if yname == "max-resv-flow" { return "MaxResvFlow" }
    if yname == "bc0-bandwidth" { return "Bc0Bandwidth" }
    if yname == "bc1-bandwidth" { return "Bc1Bandwidth" }
    if yname == "bandwidth-mode" { return "BandwidthMode" }
    return ""
}

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetSegmentPath() string {
    return "mam"
}

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-resv-bandwidth"] = mam.MaxResvBandwidth
    leafs["max-resv-flow"] = mam.MaxResvFlow
    leafs["bc0-bandwidth"] = mam.Bc0Bandwidth
    leafs["bc1-bandwidth"] = mam.Bc1Bandwidth
    leafs["bandwidth-mode"] = mam.BandwidthMode
    return leafs
}

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetBundleName() string { return "cisco_ios_xr" }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetYangName() string { return "mam" }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) SetParent(parent types.Entity) { mam.parent = parent }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetParent() types.Entity { return mam.parent }

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetParentYangName() string { return "bandwidth" }

// Rsvp_Interfaces_Interface_Bandwidth_Rdm
// Configure RDM bandwidth parameters
type Rsvp_Interfaces_Interface_Bandwidth_Rdm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Largest reservable flow (Kbps or percent of physical bandwidth). The type
    // is interface{} with range: 0..4294967295.
    MaxResvFlow interface{}

    // Reservable bandwidth in BC0 (Kbps or percent of physical bandwidth). The
    // type is interface{} with range: 0..4294967295.
    Bc0Bandwidth interface{}

    // Reservable bandwidth in BC1 (Kbps or percent of physical bandwidth). The
    // type is interface{} with range: 0..4294967295.
    Bc1Bandwidth interface{}

    // Set requests should always use RDM. The type is RsvpRdm.
    RdmKeyword interface{}

    // Set requests should always use BC0. The type is RsvpBc0.
    Bc0Keyword interface{}

    // Set requests should always use BC1. The type is RsvpBc1.
    Bc1Keyword interface{}

    // Absolute or Percentage bandwidth mode. The type is RsvpBwCfg. Units are
    // percentage.
    BandwidthMode interface{}
}

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetFilter() yfilter.YFilter { return rdm.YFilter }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) SetFilter(yf yfilter.YFilter) { rdm.YFilter = yf }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetGoName(yname string) string {
    if yname == "max-resv-flow" { return "MaxResvFlow" }
    if yname == "bc0-bandwidth" { return "Bc0Bandwidth" }
    if yname == "bc1-bandwidth" { return "Bc1Bandwidth" }
    if yname == "rdm-keyword" { return "RdmKeyword" }
    if yname == "bc0-keyword" { return "Bc0Keyword" }
    if yname == "bc1-keyword" { return "Bc1Keyword" }
    if yname == "bandwidth-mode" { return "BandwidthMode" }
    return ""
}

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetSegmentPath() string {
    return "rdm"
}

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-resv-flow"] = rdm.MaxResvFlow
    leafs["bc0-bandwidth"] = rdm.Bc0Bandwidth
    leafs["bc1-bandwidth"] = rdm.Bc1Bandwidth
    leafs["rdm-keyword"] = rdm.RdmKeyword
    leafs["bc0-keyword"] = rdm.Bc0Keyword
    leafs["bc1-keyword"] = rdm.Bc1Keyword
    leafs["bandwidth-mode"] = rdm.BandwidthMode
    return leafs
}

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetBundleName() string { return "cisco_ios_xr" }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetYangName() string { return "rdm" }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) SetParent(parent types.Entity) { rdm.parent = parent }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetParent() types.Entity { return rdm.parent }

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetParentYangName() string { return "bandwidth" }

// Rsvp_Interfaces_Interface_Authentication
// Configure RSVP authentication
type Rsvp_Interfaces_Interface_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Life time (in seconds) for each security association. The type is
    // interface{} with range: 30..86400. Units are second.
    LifeTime interface{}

    // Enable or disable RSVP authentication. The type is bool.
    Enable interface{}

    // Window-size to limit number of out-of-order messages. The type is
    // interface{} with range: 1..64.
    WindowSize interface{}

    // Key chain to authenticate RSVP signalling messages. The type is string with
    // length: 1..32.
    KeyChain interface{}
}

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Rsvp_Interfaces_Interface_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetGoName(yname string) string {
    if yname == "life-time" { return "LifeTime" }
    if yname == "enable" { return "Enable" }
    if yname == "window-size" { return "WindowSize" }
    if yname == "key-chain" { return "KeyChain" }
    return ""
}

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["life-time"] = authentication.LifeTime
    leafs["enable"] = authentication.Enable
    leafs["window-size"] = authentication.WindowSize
    leafs["key-chain"] = authentication.KeyChain
    return leafs
}

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetYangName() string { return "authentication" }

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Rsvp_Interfaces_Interface_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetParentYangName() string { return "interface" }

// Rsvp_Signalling
// Configure Global RSVP signalling parameters
type Rsvp_Signalling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure max number of consecutive missed Hello messages. The type is
    // interface{} with range: 1..10. The default value is 3.
    HelloGracefulRestartMisses interface{}

    // Configure interval between successive Hello messages. The type is
    // interface{} with range: 3000..30000. Units are millisecond. The default
    // value is 5000.
    HelloGracefulRestartInterval interface{}

    // Configure out-of-band signalling parameters.
    GlobalOutOfBand Rsvp_Signalling_GlobalOutOfBand

    // Configure RSVP Graceful-Restart parameters.
    GracefulRestart Rsvp_Signalling_GracefulRestart

    // Configure prefix filtering parameters.
    PrefixFiltering Rsvp_Signalling_PrefixFiltering

    // Sending Path Error with State-Removal flag.
    Pesr Rsvp_Signalling_Pesr

    // RSVP message checksum computation.
    Checksum Rsvp_Signalling_Checksum
}

func (signalling *Rsvp_Signalling) GetFilter() yfilter.YFilter { return signalling.YFilter }

func (signalling *Rsvp_Signalling) SetFilter(yf yfilter.YFilter) { signalling.YFilter = yf }

func (signalling *Rsvp_Signalling) GetGoName(yname string) string {
    if yname == "hello-graceful-restart-misses" { return "HelloGracefulRestartMisses" }
    if yname == "hello-graceful-restart-interval" { return "HelloGracefulRestartInterval" }
    if yname == "global-out-of-band" { return "GlobalOutOfBand" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "prefix-filtering" { return "PrefixFiltering" }
    if yname == "pesr" { return "Pesr" }
    if yname == "checksum" { return "Checksum" }
    return ""
}

func (signalling *Rsvp_Signalling) GetSegmentPath() string {
    return "signalling"
}

func (signalling *Rsvp_Signalling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-out-of-band" {
        return &signalling.GlobalOutOfBand
    }
    if childYangName == "graceful-restart" {
        return &signalling.GracefulRestart
    }
    if childYangName == "prefix-filtering" {
        return &signalling.PrefixFiltering
    }
    if childYangName == "pesr" {
        return &signalling.Pesr
    }
    if childYangName == "checksum" {
        return &signalling.Checksum
    }
    return nil
}

func (signalling *Rsvp_Signalling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global-out-of-band"] = &signalling.GlobalOutOfBand
    children["graceful-restart"] = &signalling.GracefulRestart
    children["prefix-filtering"] = &signalling.PrefixFiltering
    children["pesr"] = &signalling.Pesr
    children["checksum"] = &signalling.Checksum
    return children
}

func (signalling *Rsvp_Signalling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-graceful-restart-misses"] = signalling.HelloGracefulRestartMisses
    leafs["hello-graceful-restart-interval"] = signalling.HelloGracefulRestartInterval
    return leafs
}

func (signalling *Rsvp_Signalling) GetBundleName() string { return "cisco_ios_xr" }

func (signalling *Rsvp_Signalling) GetYangName() string { return "signalling" }

func (signalling *Rsvp_Signalling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (signalling *Rsvp_Signalling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (signalling *Rsvp_Signalling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (signalling *Rsvp_Signalling) SetParent(parent types.Entity) { signalling.parent = parent }

func (signalling *Rsvp_Signalling) GetParent() types.Entity { return signalling.parent }

func (signalling *Rsvp_Signalling) GetParentYangName() string { return "rsvp" }

// Rsvp_Signalling_GlobalOutOfBand
// Configure out-of-band signalling parameters
type Rsvp_Signalling_GlobalOutOfBand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF used for out-of-band control signalling. The type is string with
    // length: 1..32.
    Vrf interface{}
}

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetFilter() yfilter.YFilter { return globalOutOfBand.YFilter }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) SetFilter(yf yfilter.YFilter) { globalOutOfBand.YFilter = yf }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetSegmentPath() string {
    return "global-out-of-band"
}

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf"] = globalOutOfBand.Vrf
    return leafs
}

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetBundleName() string { return "cisco_ios_xr" }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetYangName() string { return "global-out-of-band" }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) SetParent(parent types.Entity) { globalOutOfBand.parent = parent }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetParent() types.Entity { return globalOutOfBand.parent }

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetParentYangName() string { return "signalling" }

// Rsvp_Signalling_GracefulRestart
// Configure RSVP Graceful-Restart parameters
type Rsvp_Signalling_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable RSVP graceful restart. The type is bool.
    Enable interface{}

    // Graceful restart time (seconds). The type is interface{} with range:
    // 60..3600. Units are second. The default value is 120.
    RestartTime interface{}

    // Graceful restart recovery time (seconds). The type is interface{} with
    // range: 0..3600. Units are second. The default value is 120.
    RecoveryTime interface{}

    // Send LSP's ctype for recovery and suggested label.
    LspClassType Rsvp_Signalling_GracefulRestart_LspClassType
}

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "recovery-time" { return "RecoveryTime" }
    if yname == "lsp-class-type" { return "LspClassType" }
    return ""
}

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-class-type" {
        return &gracefulRestart.LspClassType
    }
    return nil
}

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lsp-class-type"] = &gracefulRestart.LspClassType
    return children
}

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = gracefulRestart.Enable
    leafs["restart-time"] = gracefulRestart.RestartTime
    leafs["recovery-time"] = gracefulRestart.RecoveryTime
    return leafs
}

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetBundleName() string { return "cisco_ios_xr" }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetParentYangName() string { return "signalling" }

// Rsvp_Signalling_GracefulRestart_LspClassType
// Send LSP's ctype for recovery and suggested
// label
type Rsvp_Signalling_GracefulRestart_LspClassType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Send LSP's ctype for recovery and suggested label. The type is bool.
    Enable interface{}
}

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetFilter() yfilter.YFilter { return lspClassType.YFilter }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) SetFilter(yf yfilter.YFilter) { lspClassType.YFilter = yf }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetSegmentPath() string {
    return "lsp-class-type"
}

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = lspClassType.Enable
    return leafs
}

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetBundleName() string { return "cisco_ios_xr" }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetYangName() string { return "lsp-class-type" }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) SetParent(parent types.Entity) { lspClassType.parent = parent }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetParent() types.Entity { return lspClassType.parent }

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetParentYangName() string { return "graceful-restart" }

// Rsvp_Signalling_PrefixFiltering
// Configure prefix filtering parameters
type Rsvp_Signalling_PrefixFiltering struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure an ACL to perform prefix filtering of RSVP Router Alert messages.
    // The type is string with length: 1..65.
    Acl interface{}

    // Configure RSVP behaviour for scenarios where ACL match yields a default
    // (implicit) deny.
    DefaultDenyAction Rsvp_Signalling_PrefixFiltering_DefaultDenyAction
}

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetFilter() yfilter.YFilter { return prefixFiltering.YFilter }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) SetFilter(yf yfilter.YFilter) { prefixFiltering.YFilter = yf }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetGoName(yname string) string {
    if yname == "acl" { return "Acl" }
    if yname == "default-deny-action" { return "DefaultDenyAction" }
    return ""
}

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetSegmentPath() string {
    return "prefix-filtering"
}

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-deny-action" {
        return &prefixFiltering.DefaultDenyAction
    }
    return nil
}

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-deny-action"] = &prefixFiltering.DefaultDenyAction
    return children
}

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl"] = prefixFiltering.Acl
    return leafs
}

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetBundleName() string { return "cisco_ios_xr" }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetYangName() string { return "prefix-filtering" }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) SetParent(parent types.Entity) { prefixFiltering.parent = parent }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetParent() types.Entity { return prefixFiltering.parent }

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetParentYangName() string { return "signalling" }

// Rsvp_Signalling_PrefixFiltering_DefaultDenyAction
// Configure RSVP behaviour for scenarios where
// ACL match yields a default (implicit) deny
type Rsvp_Signalling_PrefixFiltering_DefaultDenyAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure RSVP to drop packets when ACL match yields a default (implicit)
    // deny. The type is interface{}.
    Drop interface{}
}

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetFilter() yfilter.YFilter { return defaultDenyAction.YFilter }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) SetFilter(yf yfilter.YFilter) { defaultDenyAction.YFilter = yf }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetGoName(yname string) string {
    if yname == "drop" { return "Drop" }
    return ""
}

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetSegmentPath() string {
    return "default-deny-action"
}

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["drop"] = defaultDenyAction.Drop
    return leafs
}

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetBundleName() string { return "cisco_ios_xr" }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetYangName() string { return "default-deny-action" }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) SetParent(parent types.Entity) { defaultDenyAction.parent = parent }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetParent() types.Entity { return defaultDenyAction.parent }

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetParentYangName() string { return "prefix-filtering" }

// Rsvp_Signalling_Pesr
// Sending Path Error with State-Removal flag
type Rsvp_Signalling_Pesr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable RSVP PESR. The type is interface{}.
    Disable interface{}
}

func (pesr *Rsvp_Signalling_Pesr) GetFilter() yfilter.YFilter { return pesr.YFilter }

func (pesr *Rsvp_Signalling_Pesr) SetFilter(yf yfilter.YFilter) { pesr.YFilter = yf }

func (pesr *Rsvp_Signalling_Pesr) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (pesr *Rsvp_Signalling_Pesr) GetSegmentPath() string {
    return "pesr"
}

func (pesr *Rsvp_Signalling_Pesr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pesr *Rsvp_Signalling_Pesr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pesr *Rsvp_Signalling_Pesr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = pesr.Disable
    return leafs
}

func (pesr *Rsvp_Signalling_Pesr) GetBundleName() string { return "cisco_ios_xr" }

func (pesr *Rsvp_Signalling_Pesr) GetYangName() string { return "pesr" }

func (pesr *Rsvp_Signalling_Pesr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pesr *Rsvp_Signalling_Pesr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pesr *Rsvp_Signalling_Pesr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pesr *Rsvp_Signalling_Pesr) SetParent(parent types.Entity) { pesr.parent = parent }

func (pesr *Rsvp_Signalling_Pesr) GetParent() types.Entity { return pesr.parent }

func (pesr *Rsvp_Signalling_Pesr) GetParentYangName() string { return "signalling" }

// Rsvp_Signalling_Checksum
// RSVP message checksum computation
type Rsvp_Signalling_Checksum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable RSVP message checksum computation. The type is interface{}.
    Disable interface{}
}

func (checksum *Rsvp_Signalling_Checksum) GetFilter() yfilter.YFilter { return checksum.YFilter }

func (checksum *Rsvp_Signalling_Checksum) SetFilter(yf yfilter.YFilter) { checksum.YFilter = yf }

func (checksum *Rsvp_Signalling_Checksum) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (checksum *Rsvp_Signalling_Checksum) GetSegmentPath() string {
    return "checksum"
}

func (checksum *Rsvp_Signalling_Checksum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (checksum *Rsvp_Signalling_Checksum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (checksum *Rsvp_Signalling_Checksum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = checksum.Disable
    return leafs
}

func (checksum *Rsvp_Signalling_Checksum) GetBundleName() string { return "cisco_ios_xr" }

func (checksum *Rsvp_Signalling_Checksum) GetYangName() string { return "checksum" }

func (checksum *Rsvp_Signalling_Checksum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (checksum *Rsvp_Signalling_Checksum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (checksum *Rsvp_Signalling_Checksum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (checksum *Rsvp_Signalling_Checksum) SetParent(parent types.Entity) { checksum.parent = parent }

func (checksum *Rsvp_Signalling_Checksum) GetParent() types.Entity { return checksum.parent }

func (checksum *Rsvp_Signalling_Checksum) GetParentYangName() string { return "signalling" }

// Rsvp_Authentication
// Configure RSVP authentication
type Rsvp_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Life time (in seconds) for each security association. The type is
    // interface{} with range: 30..86400. Units are second.
    LifeTime interface{}

    // Enable or disable RSVP authentication. The type is bool.
    Enable interface{}

    // Window-size to limit number of out-of-order messages. The type is
    // interface{} with range: 1..64.
    WindowSize interface{}

    // Key chain to authenticate RSVP signalling messages. The type is string with
    // length: 1..32.
    KeyChain interface{}
}

func (authentication *Rsvp_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Rsvp_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Rsvp_Authentication) GetGoName(yname string) string {
    if yname == "life-time" { return "LifeTime" }
    if yname == "enable" { return "Enable" }
    if yname == "window-size" { return "WindowSize" }
    if yname == "key-chain" { return "KeyChain" }
    return ""
}

func (authentication *Rsvp_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Rsvp_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Rsvp_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Rsvp_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["life-time"] = authentication.LifeTime
    leafs["enable"] = authentication.Enable
    leafs["window-size"] = authentication.WindowSize
    leafs["key-chain"] = authentication.KeyChain
    return leafs
}

func (authentication *Rsvp_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Rsvp_Authentication) GetYangName() string { return "authentication" }

func (authentication *Rsvp_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Rsvp_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Rsvp_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Rsvp_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Rsvp_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Rsvp_Authentication) GetParentYangName() string { return "rsvp" }

