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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
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

func (rsvp *Rsvp) GetEntityData() *types.CommonEntityData {
    rsvp.EntityData.YFilter = rsvp.YFilter
    rsvp.EntityData.YangName = "rsvp"
    rsvp.EntityData.BundleName = "cisco_ios_xr"
    rsvp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-rsvp-cfg"
    rsvp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp"
    rsvp.EntityData.AbsolutePath = rsvp.EntityData.SegmentPath
    rsvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvp.EntityData.Children = types.NewOrderedMap()
    rsvp.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &rsvp.Neighbors})
    rsvp.EntityData.Children.Append("controllers", types.YChild{"Controllers", &rsvp.Controllers})
    rsvp.EntityData.Children.Append("global-logging", types.YChild{"GlobalLogging", &rsvp.GlobalLogging})
    rsvp.EntityData.Children.Append("global-bandwidth", types.YChild{"GlobalBandwidth", &rsvp.GlobalBandwidth})
    rsvp.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &rsvp.Interfaces})
    rsvp.EntityData.Children.Append("signalling", types.YChild{"Signalling", &rsvp.Signalling})
    rsvp.EntityData.Children.Append("authentication", types.YChild{"Authentication", &rsvp.Authentication})
    rsvp.EntityData.Leafs = types.NewOrderedMap()

    rsvp.EntityData.YListKeys = []string {}

    return &(rsvp.EntityData)
}

// Rsvp_Neighbors
// RSVP Neighbor Table
type Rsvp_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP neighbor configuration. The type is slice of Rsvp_Neighbors_Neighbor.
    Neighbor []*Rsvp_Neighbors_Neighbor
}

func (neighbors *Rsvp_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "rsvp"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Rsvp_Neighbors_Neighbor
// RSVP neighbor configuration
type Rsvp_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // Configure RSVP authentication.
    Authentication Rsvp_Neighbors_Neighbor_Authentication
}

func (neighbor *Rsvp_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Neighbor, "neighbor")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("authentication", types.YChild{"Authentication", &neighbor.Authentication})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", neighbor.Neighbor})

    neighbor.EntityData.YListKeys = []string {"Neighbor"}

    return &(neighbor.EntityData)
}

// Rsvp_Neighbors_Neighbor_Authentication
// Configure RSVP authentication
type Rsvp_Neighbors_Neighbor_Authentication struct {
    EntityData types.CommonEntityData
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

func (authentication *Rsvp_Neighbors_Neighbor_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "neighbor"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/neighbors/neighbor/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("life-time", types.YLeaf{"LifeTime", authentication.LifeTime})
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", authentication.WindowSize})
    authentication.EntityData.Leafs.Append("key-chain", types.YLeaf{"KeyChain", authentication.KeyChain})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Rsvp_Controllers
// Controller table
type Rsvp_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controller configuration. The type is slice of Rsvp_Controllers_Controller.
    Controller []*Rsvp_Controllers_Controller
}

func (controllers *Rsvp_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "rsvp"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/" + controllers.EntityData.SegmentPath
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

// Rsvp_Controllers_Controller
// Controller configuration
type Rsvp_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of controller. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    ControllerName interface{}

    // Enable RSVP on an interface. The type is interface{}.
    Enable interface{}

    // Configure RSVP signalling parameters.
    CntlSignalling Rsvp_Controllers_Controller_CntlSignalling
}

func (controller *Rsvp_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/controllers/" + controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("cntl-signalling", types.YChild{"CntlSignalling", &controller.CntlSignalling})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})
    controller.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", controller.Enable})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// Rsvp_Controllers_Controller_CntlSignalling
// Configure RSVP signalling parameters
type Rsvp_Controllers_Controller_CntlSignalling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure RSVP out-of-band signalling parameters.
    OutOfBand Rsvp_Controllers_Controller_CntlSignalling_OutOfBand
}

func (cntlSignalling *Rsvp_Controllers_Controller_CntlSignalling) GetEntityData() *types.CommonEntityData {
    cntlSignalling.EntityData.YFilter = cntlSignalling.YFilter
    cntlSignalling.EntityData.YangName = "cntl-signalling"
    cntlSignalling.EntityData.BundleName = "cisco_ios_xr"
    cntlSignalling.EntityData.ParentYangName = "controller"
    cntlSignalling.EntityData.SegmentPath = "cntl-signalling"
    cntlSignalling.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/controllers/controller/" + cntlSignalling.EntityData.SegmentPath
    cntlSignalling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cntlSignalling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cntlSignalling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cntlSignalling.EntityData.Children = types.NewOrderedMap()
    cntlSignalling.EntityData.Children.Append("out-of-band", types.YChild{"OutOfBand", &cntlSignalling.OutOfBand})
    cntlSignalling.EntityData.Leafs = types.NewOrderedMap()

    cntlSignalling.EntityData.YListKeys = []string {}

    return &(cntlSignalling.EntityData)
}

// Rsvp_Controllers_Controller_CntlSignalling_OutOfBand
// Configure RSVP out-of-band signalling parameters
type Rsvp_Controllers_Controller_CntlSignalling_OutOfBand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure max number of consecutive missed messages for state expiry for
    // out-of-band tunnels. The type is interface{} with range: 1..110000. The
    // default value is 38000.
    MissedMessages interface{}

    // Configure interval between successive refreshes for out-of-band tunnels.
    // The type is interface{} with range: 180..86400. Units are second.
    RefreshInterval interface{}
}

func (outOfBand *Rsvp_Controllers_Controller_CntlSignalling_OutOfBand) GetEntityData() *types.CommonEntityData {
    outOfBand.EntityData.YFilter = outOfBand.YFilter
    outOfBand.EntityData.YangName = "out-of-band"
    outOfBand.EntityData.BundleName = "cisco_ios_xr"
    outOfBand.EntityData.ParentYangName = "cntl-signalling"
    outOfBand.EntityData.SegmentPath = "out-of-band"
    outOfBand.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/controllers/controller/cntl-signalling/" + outOfBand.EntityData.SegmentPath
    outOfBand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outOfBand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outOfBand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outOfBand.EntityData.Children = types.NewOrderedMap()
    outOfBand.EntityData.Leafs = types.NewOrderedMap()
    outOfBand.EntityData.Leafs.Append("missed-messages", types.YLeaf{"MissedMessages", outOfBand.MissedMessages})
    outOfBand.EntityData.Leafs.Append("refresh-interval", types.YLeaf{"RefreshInterval", outOfBand.RefreshInterval})

    outOfBand.EntityData.YListKeys = []string {}

    return &(outOfBand.EntityData)
}

// Rsvp_GlobalLogging
// Global Logging
type Rsvp_GlobalLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable NSR Status Logging. The type is interface{}.
    LogNsrStatus interface{}

    // Enable ISSU Status Logging. The type is interface{}.
    LogIssuStatus interface{}
}

func (globalLogging *Rsvp_GlobalLogging) GetEntityData() *types.CommonEntityData {
    globalLogging.EntityData.YFilter = globalLogging.YFilter
    globalLogging.EntityData.YangName = "global-logging"
    globalLogging.EntityData.BundleName = "cisco_ios_xr"
    globalLogging.EntityData.ParentYangName = "rsvp"
    globalLogging.EntityData.SegmentPath = "global-logging"
    globalLogging.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/" + globalLogging.EntityData.SegmentPath
    globalLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalLogging.EntityData.Children = types.NewOrderedMap()
    globalLogging.EntityData.Leafs = types.NewOrderedMap()
    globalLogging.EntityData.Leafs.Append("log-nsr-status", types.YLeaf{"LogNsrStatus", globalLogging.LogNsrStatus})
    globalLogging.EntityData.Leafs.Append("log-issu-status", types.YLeaf{"LogIssuStatus", globalLogging.LogIssuStatus})

    globalLogging.EntityData.YListKeys = []string {}

    return &(globalLogging.EntityData)
}

// Rsvp_GlobalBandwidth
// Configure Global Bandwidth Parameters
type Rsvp_GlobalBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Global RSVP signalling parameters.
    DefaultInterfacePercent Rsvp_GlobalBandwidth_DefaultInterfacePercent
}

func (globalBandwidth *Rsvp_GlobalBandwidth) GetEntityData() *types.CommonEntityData {
    globalBandwidth.EntityData.YFilter = globalBandwidth.YFilter
    globalBandwidth.EntityData.YangName = "global-bandwidth"
    globalBandwidth.EntityData.BundleName = "cisco_ios_xr"
    globalBandwidth.EntityData.ParentYangName = "rsvp"
    globalBandwidth.EntityData.SegmentPath = "global-bandwidth"
    globalBandwidth.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/" + globalBandwidth.EntityData.SegmentPath
    globalBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalBandwidth.EntityData.Children = types.NewOrderedMap()
    globalBandwidth.EntityData.Children.Append("default-interface-percent", types.YChild{"DefaultInterfacePercent", &globalBandwidth.DefaultInterfacePercent})
    globalBandwidth.EntityData.Leafs = types.NewOrderedMap()

    globalBandwidth.EntityData.YListKeys = []string {}

    return &(globalBandwidth.EntityData)
}

// Rsvp_GlobalBandwidth_DefaultInterfacePercent
// Configure Global RSVP signalling parameters
type Rsvp_GlobalBandwidth_DefaultInterfacePercent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure global default MAM I/F percent bandwidth parameters.
    Mam Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam

    // Configure global default RDM I/F percent bandwidth parameters.
    Rdm Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm
}

func (defaultInterfacePercent *Rsvp_GlobalBandwidth_DefaultInterfacePercent) GetEntityData() *types.CommonEntityData {
    defaultInterfacePercent.EntityData.YFilter = defaultInterfacePercent.YFilter
    defaultInterfacePercent.EntityData.YangName = "default-interface-percent"
    defaultInterfacePercent.EntityData.BundleName = "cisco_ios_xr"
    defaultInterfacePercent.EntityData.ParentYangName = "global-bandwidth"
    defaultInterfacePercent.EntityData.SegmentPath = "default-interface-percent"
    defaultInterfacePercent.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/global-bandwidth/" + defaultInterfacePercent.EntityData.SegmentPath
    defaultInterfacePercent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultInterfacePercent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultInterfacePercent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultInterfacePercent.EntityData.Children = types.NewOrderedMap()
    defaultInterfacePercent.EntityData.Children.Append("mam", types.YChild{"Mam", &defaultInterfacePercent.Mam})
    defaultInterfacePercent.EntityData.Children.Append("rdm", types.YChild{"Rdm", &defaultInterfacePercent.Rdm})
    defaultInterfacePercent.EntityData.Leafs = types.NewOrderedMap()

    defaultInterfacePercent.EntityData.YListKeys = []string {}

    return &(defaultInterfacePercent.EntityData)
}

// Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam
// Configure global default MAM I/F percent
// bandwidth parameters
type Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default maximum reservable I/F % B/W . The type is interface{} with range:
    // 0..10000.
    MaxResPercent interface{}

    // Default BC0 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc0Percent interface{}

    // Default BC1 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc1Percent interface{}
}

func (mam *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Mam) GetEntityData() *types.CommonEntityData {
    mam.EntityData.YFilter = mam.YFilter
    mam.EntityData.YangName = "mam"
    mam.EntityData.BundleName = "cisco_ios_xr"
    mam.EntityData.ParentYangName = "default-interface-percent"
    mam.EntityData.SegmentPath = "mam"
    mam.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/global-bandwidth/default-interface-percent/" + mam.EntityData.SegmentPath
    mam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mam.EntityData.Children = types.NewOrderedMap()
    mam.EntityData.Leafs = types.NewOrderedMap()
    mam.EntityData.Leafs.Append("max-res-percent", types.YLeaf{"MaxResPercent", mam.MaxResPercent})
    mam.EntityData.Leafs.Append("bc0-percent", types.YLeaf{"Bc0Percent", mam.Bc0Percent})
    mam.EntityData.Leafs.Append("bc1-percent", types.YLeaf{"Bc1Percent", mam.Bc1Percent})

    mam.EntityData.YListKeys = []string {}

    return &(mam.EntityData)
}

// Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm
// Configure global default RDM I/F percent
// bandwidth parameters
type Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default BC0 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc0Percent interface{}

    // Default BC1 pool I/F % B/W . The type is interface{} with range: 0..10000.
    Bc1Percent interface{}
}

func (rdm *Rsvp_GlobalBandwidth_DefaultInterfacePercent_Rdm) GetEntityData() *types.CommonEntityData {
    rdm.EntityData.YFilter = rdm.YFilter
    rdm.EntityData.YangName = "rdm"
    rdm.EntityData.BundleName = "cisco_ios_xr"
    rdm.EntityData.ParentYangName = "default-interface-percent"
    rdm.EntityData.SegmentPath = "rdm"
    rdm.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/global-bandwidth/default-interface-percent/" + rdm.EntityData.SegmentPath
    rdm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdm.EntityData.Children = types.NewOrderedMap()
    rdm.EntityData.Leafs = types.NewOrderedMap()
    rdm.EntityData.Leafs.Append("bc0-percent", types.YLeaf{"Bc0Percent", rdm.Bc0Percent})
    rdm.EntityData.Leafs.Append("bc1-percent", types.YLeaf{"Bc1Percent", rdm.Bc1Percent})

    rdm.EntityData.YListKeys = []string {}

    return &(rdm.EntityData)
}

// Rsvp_Interfaces
// Interface table
type Rsvp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface configuration. The type is slice of Rsvp_Interfaces_Interface.
    Interface []*Rsvp_Interfaces_Interface
}

func (interfaces *Rsvp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "rsvp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Rsvp_Interfaces_Interface
// Interface configuration
type Rsvp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
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

func (self *Rsvp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Name, "name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("if-signalling", types.YChild{"IfSignalling", &self.IfSignalling})
    self.EntityData.Children.Append("bandwidth", types.YChild{"Bandwidth", &self.Bandwidth})
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})

    self.EntityData.YListKeys = []string {"Name"}

    return &(self.EntityData)
}

// Rsvp_Interfaces_Interface_IfSignalling
// Configure RSVP signalling parameters
type Rsvp_Interfaces_Interface_IfSignalling struct {
    EntityData types.CommonEntityData
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

func (ifSignalling *Rsvp_Interfaces_Interface_IfSignalling) GetEntityData() *types.CommonEntityData {
    ifSignalling.EntityData.YFilter = ifSignalling.YFilter
    ifSignalling.EntityData.YangName = "if-signalling"
    ifSignalling.EntityData.BundleName = "cisco_ios_xr"
    ifSignalling.EntityData.ParentYangName = "interface"
    ifSignalling.EntityData.SegmentPath = "if-signalling"
    ifSignalling.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/" + ifSignalling.EntityData.SegmentPath
    ifSignalling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifSignalling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifSignalling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifSignalling.EntityData.Children = types.NewOrderedMap()
    ifSignalling.EntityData.Children.Append("refresh-reduction", types.YChild{"RefreshReduction", &ifSignalling.RefreshReduction})
    ifSignalling.EntityData.Children.Append("interval-rate", types.YChild{"IntervalRate", &ifSignalling.IntervalRate})
    ifSignalling.EntityData.Children.Append("out-of-band", types.YChild{"OutOfBand", &ifSignalling.OutOfBand})
    ifSignalling.EntityData.Leafs = types.NewOrderedMap()
    ifSignalling.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", ifSignalling.Dscp})
    ifSignalling.EntityData.Leafs.Append("missed-messages", types.YLeaf{"MissedMessages", ifSignalling.MissedMessages})
    ifSignalling.EntityData.Leafs.Append("hello-graceful-restart-if-based", types.YLeaf{"HelloGracefulRestartIfBased", ifSignalling.HelloGracefulRestartIfBased})
    ifSignalling.EntityData.Leafs.Append("pacing", types.YLeaf{"Pacing", ifSignalling.Pacing})
    ifSignalling.EntityData.Leafs.Append("refresh-interval", types.YLeaf{"RefreshInterval", ifSignalling.RefreshInterval})

    ifSignalling.EntityData.YListKeys = []string {}

    return &(ifSignalling.EntityData)
}

// Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction
// Configure RSVP Refresh Reduction parameters
type Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction struct {
    EntityData types.CommonEntityData
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

func (refreshReduction *Rsvp_Interfaces_Interface_IfSignalling_RefreshReduction) GetEntityData() *types.CommonEntityData {
    refreshReduction.EntityData.YFilter = refreshReduction.YFilter
    refreshReduction.EntityData.YangName = "refresh-reduction"
    refreshReduction.EntityData.BundleName = "cisco_ios_xr"
    refreshReduction.EntityData.ParentYangName = "if-signalling"
    refreshReduction.EntityData.SegmentPath = "refresh-reduction"
    refreshReduction.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/if-signalling/" + refreshReduction.EntityData.SegmentPath
    refreshReduction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    refreshReduction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    refreshReduction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    refreshReduction.EntityData.Children = types.NewOrderedMap()
    refreshReduction.EntityData.Leafs = types.NewOrderedMap()
    refreshReduction.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", refreshReduction.Disable})
    refreshReduction.EntityData.Leafs.Append("reliable-ack-max-size", types.YLeaf{"ReliableAckMaxSize", refreshReduction.ReliableAckMaxSize})
    refreshReduction.EntityData.Leafs.Append("reliable-ack-hold-time", types.YLeaf{"ReliableAckHoldTime", refreshReduction.ReliableAckHoldTime})
    refreshReduction.EntityData.Leafs.Append("reliable-retransmit-time", types.YLeaf{"ReliableRetransmitTime", refreshReduction.ReliableRetransmitTime})
    refreshReduction.EntityData.Leafs.Append("reliable-s-refresh", types.YLeaf{"ReliableSRefresh", refreshReduction.ReliableSRefresh})
    refreshReduction.EntityData.Leafs.Append("summary-max-size", types.YLeaf{"SummaryMaxSize", refreshReduction.SummaryMaxSize})
    refreshReduction.EntityData.Leafs.Append("bundle-message-max-size", types.YLeaf{"BundleMessageMaxSize", refreshReduction.BundleMessageMaxSize})

    refreshReduction.EntityData.YListKeys = []string {}

    return &(refreshReduction.EntityData)
}

// Rsvp_Interfaces_Interface_IfSignalling_IntervalRate
// Configure number of messages to be sent per
// interval
type Rsvp_Interfaces_Interface_IfSignalling_IntervalRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of messages to be sent per interval. The type is interface{} with
    // range: 1..500. The default value is 100.
    MessagesPerInterval interface{}

    // Size of an interval (milliseconds). The type is interface{} with range:
    // 250..2000. Units are millisecond. The default value is 1000.
    IntervalSize interface{}
}

func (intervalRate *Rsvp_Interfaces_Interface_IfSignalling_IntervalRate) GetEntityData() *types.CommonEntityData {
    intervalRate.EntityData.YFilter = intervalRate.YFilter
    intervalRate.EntityData.YangName = "interval-rate"
    intervalRate.EntityData.BundleName = "cisco_ios_xr"
    intervalRate.EntityData.ParentYangName = "if-signalling"
    intervalRate.EntityData.SegmentPath = "interval-rate"
    intervalRate.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/if-signalling/" + intervalRate.EntityData.SegmentPath
    intervalRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intervalRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intervalRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intervalRate.EntityData.Children = types.NewOrderedMap()
    intervalRate.EntityData.Leafs = types.NewOrderedMap()
    intervalRate.EntityData.Leafs.Append("messages-per-interval", types.YLeaf{"MessagesPerInterval", intervalRate.MessagesPerInterval})
    intervalRate.EntityData.Leafs.Append("interval-size", types.YLeaf{"IntervalSize", intervalRate.IntervalSize})

    intervalRate.EntityData.YListKeys = []string {}

    return &(intervalRate.EntityData)
}

// Rsvp_Interfaces_Interface_IfSignalling_OutOfBand
// Configure RSVP out-of-band signalling parameters
type Rsvp_Interfaces_Interface_IfSignalling_OutOfBand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure max number of consecutive missed messages for state expiry for
    // out-of-band tunnels. The type is interface{} with range: 1..110000. The
    // default value is 38000.
    MissedMessages interface{}

    // Configure interval between successive refreshes for out-of-band tunnels.
    // The type is interface{} with range: 180..86400. Units are second.
    RefreshInterval interface{}
}

func (outOfBand *Rsvp_Interfaces_Interface_IfSignalling_OutOfBand) GetEntityData() *types.CommonEntityData {
    outOfBand.EntityData.YFilter = outOfBand.YFilter
    outOfBand.EntityData.YangName = "out-of-band"
    outOfBand.EntityData.BundleName = "cisco_ios_xr"
    outOfBand.EntityData.ParentYangName = "if-signalling"
    outOfBand.EntityData.SegmentPath = "out-of-band"
    outOfBand.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/if-signalling/" + outOfBand.EntityData.SegmentPath
    outOfBand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outOfBand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outOfBand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outOfBand.EntityData.Children = types.NewOrderedMap()
    outOfBand.EntityData.Leafs = types.NewOrderedMap()
    outOfBand.EntityData.Leafs.Append("missed-messages", types.YLeaf{"MissedMessages", outOfBand.MissedMessages})
    outOfBand.EntityData.Leafs.Append("refresh-interval", types.YLeaf{"RefreshInterval", outOfBand.RefreshInterval})

    outOfBand.EntityData.YListKeys = []string {}

    return &(outOfBand.EntityData)
}

// Rsvp_Interfaces_Interface_Bandwidth
// Configure Bandwidth
type Rsvp_Interfaces_Interface_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure MAM bandwidth parameters.
    Mam Rsvp_Interfaces_Interface_Bandwidth_Mam

    // Configure RDM bandwidth parameters.
    Rdm Rsvp_Interfaces_Interface_Bandwidth_Rdm
}

func (bandwidth *Rsvp_Interfaces_Interface_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "interface"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/" + bandwidth.EntityData.SegmentPath
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = types.NewOrderedMap()
    bandwidth.EntityData.Children.Append("mam", types.YChild{"Mam", &bandwidth.Mam})
    bandwidth.EntityData.Children.Append("rdm", types.YChild{"Rdm", &bandwidth.Rdm})
    bandwidth.EntityData.Leafs = types.NewOrderedMap()

    bandwidth.EntityData.YListKeys = []string {}

    return &(bandwidth.EntityData)
}

// Rsvp_Interfaces_Interface_Bandwidth_Mam
// Configure MAM bandwidth parameters
type Rsvp_Interfaces_Interface_Bandwidth_Mam struct {
    EntityData types.CommonEntityData
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

func (mam *Rsvp_Interfaces_Interface_Bandwidth_Mam) GetEntityData() *types.CommonEntityData {
    mam.EntityData.YFilter = mam.YFilter
    mam.EntityData.YangName = "mam"
    mam.EntityData.BundleName = "cisco_ios_xr"
    mam.EntityData.ParentYangName = "bandwidth"
    mam.EntityData.SegmentPath = "mam"
    mam.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/bandwidth/" + mam.EntityData.SegmentPath
    mam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mam.EntityData.Children = types.NewOrderedMap()
    mam.EntityData.Leafs = types.NewOrderedMap()
    mam.EntityData.Leafs.Append("max-resv-bandwidth", types.YLeaf{"MaxResvBandwidth", mam.MaxResvBandwidth})
    mam.EntityData.Leafs.Append("max-resv-flow", types.YLeaf{"MaxResvFlow", mam.MaxResvFlow})
    mam.EntityData.Leafs.Append("bc0-bandwidth", types.YLeaf{"Bc0Bandwidth", mam.Bc0Bandwidth})
    mam.EntityData.Leafs.Append("bc1-bandwidth", types.YLeaf{"Bc1Bandwidth", mam.Bc1Bandwidth})
    mam.EntityData.Leafs.Append("bandwidth-mode", types.YLeaf{"BandwidthMode", mam.BandwidthMode})

    mam.EntityData.YListKeys = []string {}

    return &(mam.EntityData)
}

// Rsvp_Interfaces_Interface_Bandwidth_Rdm
// Configure RDM bandwidth parameters
type Rsvp_Interfaces_Interface_Bandwidth_Rdm struct {
    EntityData types.CommonEntityData
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

func (rdm *Rsvp_Interfaces_Interface_Bandwidth_Rdm) GetEntityData() *types.CommonEntityData {
    rdm.EntityData.YFilter = rdm.YFilter
    rdm.EntityData.YangName = "rdm"
    rdm.EntityData.BundleName = "cisco_ios_xr"
    rdm.EntityData.ParentYangName = "bandwidth"
    rdm.EntityData.SegmentPath = "rdm"
    rdm.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/bandwidth/" + rdm.EntityData.SegmentPath
    rdm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdm.EntityData.Children = types.NewOrderedMap()
    rdm.EntityData.Leafs = types.NewOrderedMap()
    rdm.EntityData.Leafs.Append("max-resv-flow", types.YLeaf{"MaxResvFlow", rdm.MaxResvFlow})
    rdm.EntityData.Leafs.Append("bc0-bandwidth", types.YLeaf{"Bc0Bandwidth", rdm.Bc0Bandwidth})
    rdm.EntityData.Leafs.Append("bc1-bandwidth", types.YLeaf{"Bc1Bandwidth", rdm.Bc1Bandwidth})
    rdm.EntityData.Leafs.Append("rdm-keyword", types.YLeaf{"RdmKeyword", rdm.RdmKeyword})
    rdm.EntityData.Leafs.Append("bc0-keyword", types.YLeaf{"Bc0Keyword", rdm.Bc0Keyword})
    rdm.EntityData.Leafs.Append("bc1-keyword", types.YLeaf{"Bc1Keyword", rdm.Bc1Keyword})
    rdm.EntityData.Leafs.Append("bandwidth-mode", types.YLeaf{"BandwidthMode", rdm.BandwidthMode})

    rdm.EntityData.YListKeys = []string {}

    return &(rdm.EntityData)
}

// Rsvp_Interfaces_Interface_Authentication
// Configure RSVP authentication
type Rsvp_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
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

func (authentication *Rsvp_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/interfaces/interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("life-time", types.YLeaf{"LifeTime", authentication.LifeTime})
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", authentication.WindowSize})
    authentication.EntityData.Leafs.Append("key-chain", types.YLeaf{"KeyChain", authentication.KeyChain})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Rsvp_Signalling
// Configure Global RSVP signalling parameters
type Rsvp_Signalling struct {
    EntityData types.CommonEntityData
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

func (signalling *Rsvp_Signalling) GetEntityData() *types.CommonEntityData {
    signalling.EntityData.YFilter = signalling.YFilter
    signalling.EntityData.YangName = "signalling"
    signalling.EntityData.BundleName = "cisco_ios_xr"
    signalling.EntityData.ParentYangName = "rsvp"
    signalling.EntityData.SegmentPath = "signalling"
    signalling.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/" + signalling.EntityData.SegmentPath
    signalling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalling.EntityData.Children = types.NewOrderedMap()
    signalling.EntityData.Children.Append("global-out-of-band", types.YChild{"GlobalOutOfBand", &signalling.GlobalOutOfBand})
    signalling.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &signalling.GracefulRestart})
    signalling.EntityData.Children.Append("prefix-filtering", types.YChild{"PrefixFiltering", &signalling.PrefixFiltering})
    signalling.EntityData.Children.Append("pesr", types.YChild{"Pesr", &signalling.Pesr})
    signalling.EntityData.Children.Append("checksum", types.YChild{"Checksum", &signalling.Checksum})
    signalling.EntityData.Leafs = types.NewOrderedMap()
    signalling.EntityData.Leafs.Append("hello-graceful-restart-misses", types.YLeaf{"HelloGracefulRestartMisses", signalling.HelloGracefulRestartMisses})
    signalling.EntityData.Leafs.Append("hello-graceful-restart-interval", types.YLeaf{"HelloGracefulRestartInterval", signalling.HelloGracefulRestartInterval})

    signalling.EntityData.YListKeys = []string {}

    return &(signalling.EntityData)
}

// Rsvp_Signalling_GlobalOutOfBand
// Configure out-of-band signalling parameters
type Rsvp_Signalling_GlobalOutOfBand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF used for out-of-band control signalling. The type is string with
    // length: 1..32.
    Vrf interface{}
}

func (globalOutOfBand *Rsvp_Signalling_GlobalOutOfBand) GetEntityData() *types.CommonEntityData {
    globalOutOfBand.EntityData.YFilter = globalOutOfBand.YFilter
    globalOutOfBand.EntityData.YangName = "global-out-of-band"
    globalOutOfBand.EntityData.BundleName = "cisco_ios_xr"
    globalOutOfBand.EntityData.ParentYangName = "signalling"
    globalOutOfBand.EntityData.SegmentPath = "global-out-of-band"
    globalOutOfBand.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/signalling/" + globalOutOfBand.EntityData.SegmentPath
    globalOutOfBand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalOutOfBand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalOutOfBand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalOutOfBand.EntityData.Children = types.NewOrderedMap()
    globalOutOfBand.EntityData.Leafs = types.NewOrderedMap()
    globalOutOfBand.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", globalOutOfBand.Vrf})

    globalOutOfBand.EntityData.YListKeys = []string {}

    return &(globalOutOfBand.EntityData)
}

// Rsvp_Signalling_GracefulRestart
// Configure RSVP Graceful-Restart parameters
type Rsvp_Signalling_GracefulRestart struct {
    EntityData types.CommonEntityData
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

func (gracefulRestart *Rsvp_Signalling_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xr"
    gracefulRestart.EntityData.ParentYangName = "signalling"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/signalling/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("lsp-class-type", types.YChild{"LspClassType", &gracefulRestart.LspClassType})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", gracefulRestart.Enable})
    gracefulRestart.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", gracefulRestart.RestartTime})
    gracefulRestart.EntityData.Leafs.Append("recovery-time", types.YLeaf{"RecoveryTime", gracefulRestart.RecoveryTime})

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Rsvp_Signalling_GracefulRestart_LspClassType
// Send LSP's ctype for recovery and suggested
// label
type Rsvp_Signalling_GracefulRestart_LspClassType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Send LSP's ctype for recovery and suggested label. The type is bool.
    Enable interface{}
}

func (lspClassType *Rsvp_Signalling_GracefulRestart_LspClassType) GetEntityData() *types.CommonEntityData {
    lspClassType.EntityData.YFilter = lspClassType.YFilter
    lspClassType.EntityData.YangName = "lsp-class-type"
    lspClassType.EntityData.BundleName = "cisco_ios_xr"
    lspClassType.EntityData.ParentYangName = "graceful-restart"
    lspClassType.EntityData.SegmentPath = "lsp-class-type"
    lspClassType.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/signalling/graceful-restart/" + lspClassType.EntityData.SegmentPath
    lspClassType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspClassType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspClassType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspClassType.EntityData.Children = types.NewOrderedMap()
    lspClassType.EntityData.Leafs = types.NewOrderedMap()
    lspClassType.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", lspClassType.Enable})

    lspClassType.EntityData.YListKeys = []string {}

    return &(lspClassType.EntityData)
}

// Rsvp_Signalling_PrefixFiltering
// Configure prefix filtering parameters
type Rsvp_Signalling_PrefixFiltering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure an ACL to perform prefix filtering of RSVP Router Alert messages.
    // The type is string with length: 1..65.
    Acl interface{}

    // Configure RSVP behaviour for scenarios where ACL match yields a default
    // (implicit) deny.
    DefaultDenyAction Rsvp_Signalling_PrefixFiltering_DefaultDenyAction
}

func (prefixFiltering *Rsvp_Signalling_PrefixFiltering) GetEntityData() *types.CommonEntityData {
    prefixFiltering.EntityData.YFilter = prefixFiltering.YFilter
    prefixFiltering.EntityData.YangName = "prefix-filtering"
    prefixFiltering.EntityData.BundleName = "cisco_ios_xr"
    prefixFiltering.EntityData.ParentYangName = "signalling"
    prefixFiltering.EntityData.SegmentPath = "prefix-filtering"
    prefixFiltering.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/signalling/" + prefixFiltering.EntityData.SegmentPath
    prefixFiltering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixFiltering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixFiltering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixFiltering.EntityData.Children = types.NewOrderedMap()
    prefixFiltering.EntityData.Children.Append("default-deny-action", types.YChild{"DefaultDenyAction", &prefixFiltering.DefaultDenyAction})
    prefixFiltering.EntityData.Leafs = types.NewOrderedMap()
    prefixFiltering.EntityData.Leafs.Append("acl", types.YLeaf{"Acl", prefixFiltering.Acl})

    prefixFiltering.EntityData.YListKeys = []string {}

    return &(prefixFiltering.EntityData)
}

// Rsvp_Signalling_PrefixFiltering_DefaultDenyAction
// Configure RSVP behaviour for scenarios where
// ACL match yields a default (implicit) deny
type Rsvp_Signalling_PrefixFiltering_DefaultDenyAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure RSVP to drop packets when ACL match yields a default (implicit)
    // deny. The type is interface{}.
    Drop interface{}
}

func (defaultDenyAction *Rsvp_Signalling_PrefixFiltering_DefaultDenyAction) GetEntityData() *types.CommonEntityData {
    defaultDenyAction.EntityData.YFilter = defaultDenyAction.YFilter
    defaultDenyAction.EntityData.YangName = "default-deny-action"
    defaultDenyAction.EntityData.BundleName = "cisco_ios_xr"
    defaultDenyAction.EntityData.ParentYangName = "prefix-filtering"
    defaultDenyAction.EntityData.SegmentPath = "default-deny-action"
    defaultDenyAction.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/signalling/prefix-filtering/" + defaultDenyAction.EntityData.SegmentPath
    defaultDenyAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultDenyAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultDenyAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultDenyAction.EntityData.Children = types.NewOrderedMap()
    defaultDenyAction.EntityData.Leafs = types.NewOrderedMap()
    defaultDenyAction.EntityData.Leafs.Append("drop", types.YLeaf{"Drop", defaultDenyAction.Drop})

    defaultDenyAction.EntityData.YListKeys = []string {}

    return &(defaultDenyAction.EntityData)
}

// Rsvp_Signalling_Pesr
// Sending Path Error with State-Removal flag
type Rsvp_Signalling_Pesr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable RSVP PESR. The type is interface{}.
    Disable interface{}
}

func (pesr *Rsvp_Signalling_Pesr) GetEntityData() *types.CommonEntityData {
    pesr.EntityData.YFilter = pesr.YFilter
    pesr.EntityData.YangName = "pesr"
    pesr.EntityData.BundleName = "cisco_ios_xr"
    pesr.EntityData.ParentYangName = "signalling"
    pesr.EntityData.SegmentPath = "pesr"
    pesr.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/signalling/" + pesr.EntityData.SegmentPath
    pesr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pesr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pesr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pesr.EntityData.Children = types.NewOrderedMap()
    pesr.EntityData.Leafs = types.NewOrderedMap()
    pesr.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pesr.Disable})

    pesr.EntityData.YListKeys = []string {}

    return &(pesr.EntityData)
}

// Rsvp_Signalling_Checksum
// RSVP message checksum computation
type Rsvp_Signalling_Checksum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable RSVP message checksum computation. The type is interface{}.
    Disable interface{}
}

func (checksum *Rsvp_Signalling_Checksum) GetEntityData() *types.CommonEntityData {
    checksum.EntityData.YFilter = checksum.YFilter
    checksum.EntityData.YangName = "checksum"
    checksum.EntityData.BundleName = "cisco_ios_xr"
    checksum.EntityData.ParentYangName = "signalling"
    checksum.EntityData.SegmentPath = "checksum"
    checksum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/signalling/" + checksum.EntityData.SegmentPath
    checksum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    checksum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    checksum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    checksum.EntityData.Children = types.NewOrderedMap()
    checksum.EntityData.Leafs = types.NewOrderedMap()
    checksum.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", checksum.Disable})

    checksum.EntityData.YListKeys = []string {}

    return &(checksum.EntityData)
}

// Rsvp_Authentication
// Configure RSVP authentication
type Rsvp_Authentication struct {
    EntityData types.CommonEntityData
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

func (authentication *Rsvp_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "rsvp"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("life-time", types.YLeaf{"LifeTime", authentication.LifeTime})
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", authentication.WindowSize})
    authentication.EntityData.Leafs.Append("key-chain", types.YLeaf{"KeyChain", authentication.KeyChain})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

