// This module contains a collection of YANG definitions
// for Cisco IOS-XR lmp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   lmp: Main common OLM/LMP configuration container
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package lmp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lmp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lmp-cfg lmp}", reflect.TypeOf(Lmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lmp-cfg:lmp", reflect.TypeOf(Lmp{}))
}

// OlmAddr represents Olm addr
type OlmAddr string

const (
    // IPv4 address
    OlmAddr_ipv4 OlmAddr = "ipv4"

    // IPv6 address
    OlmAddr_ipv6 OlmAddr = "ipv6"

    // Unnumbered address
    OlmAddr_unnumbered OlmAddr = "unnumbered"

    // NSAP address
    OlmAddr_nsap OlmAddr = "nsap"
)

// OlmSwitchingCap represents Olm switching cap
type OlmSwitchingCap string

const (
    // Lambda switch capable
    OlmSwitchingCap_lsc OlmSwitchingCap = "lsc"

    // Fiber switch capable
    OlmSwitchingCap_fsc OlmSwitchingCap = "fsc"
)

// Lmp
// Main common OLM/LMP configuration container
type Lmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable the OLM/LMP application. The type is interface{}.
    Enable interface{}

    // GMPLS UNI specific OLM/LMP configuration.
    GmplsUni Lmp_GmplsUni
}

func (lmp *Lmp) GetEntityData() *types.CommonEntityData {
    lmp.EntityData.YFilter = lmp.YFilter
    lmp.EntityData.YangName = "lmp"
    lmp.EntityData.BundleName = "cisco_ios_xr"
    lmp.EntityData.ParentYangName = "Cisco-IOS-XR-lmp-cfg"
    lmp.EntityData.SegmentPath = "Cisco-IOS-XR-lmp-cfg:lmp"
    lmp.EntityData.AbsolutePath = lmp.EntityData.SegmentPath
    lmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmp.EntityData.Children = types.NewOrderedMap()
    lmp.EntityData.Children.Append("gmpls-uni", types.YChild{"GmplsUni", &lmp.GmplsUni})
    lmp.EntityData.Leafs = types.NewOrderedMap()
    lmp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", lmp.Enable})

    lmp.EntityData.YListKeys = []string {}

    return &(lmp.EntityData)
}

// Lmp_GmplsUni
// GMPLS UNI specific OLM/LMP configuration
type Lmp_GmplsUni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor configuration.
    Neighbors Lmp_GmplsUni_Neighbors

    // Local GMPLS UNI router ID.
    RouterId Lmp_GmplsUni_RouterId

    // Configure GMPLS UNI OLM/LMP controllers.
    Controllers Lmp_GmplsUni_Controllers
}

func (gmplsUni *Lmp_GmplsUni) GetEntityData() *types.CommonEntityData {
    gmplsUni.EntityData.YFilter = gmplsUni.YFilter
    gmplsUni.EntityData.YangName = "gmpls-uni"
    gmplsUni.EntityData.BundleName = "cisco_ios_xr"
    gmplsUni.EntityData.ParentYangName = "lmp"
    gmplsUni.EntityData.SegmentPath = "gmpls-uni"
    gmplsUni.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/" + gmplsUni.EntityData.SegmentPath
    gmplsUni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gmplsUni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gmplsUni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gmplsUni.EntityData.Children = types.NewOrderedMap()
    gmplsUni.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &gmplsUni.Neighbors})
    gmplsUni.EntityData.Children.Append("router-id", types.YChild{"RouterId", &gmplsUni.RouterId})
    gmplsUni.EntityData.Children.Append("controllers", types.YChild{"Controllers", &gmplsUni.Controllers})
    gmplsUni.EntityData.Leafs = types.NewOrderedMap()

    gmplsUni.EntityData.YListKeys = []string {}

    return &(gmplsUni.EntityData)
}

// Lmp_GmplsUni_Neighbors
// Neighbor configuration
type Lmp_GmplsUni_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor configuration. The type is slice of
    // Lmp_GmplsUni_Neighbors_Neighbor.
    Neighbor []*Lmp_GmplsUni_Neighbors_Neighbor
}

func (neighbors *Lmp_GmplsUni_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "gmpls-uni"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/" + neighbors.EntityData.SegmentPath
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

// Lmp_GmplsUni_Neighbors_Neighbor
// Neighbor configuration
type Lmp_GmplsUni_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor name. The type is string with length:
    // 1..64.
    NeighborName interface{}

    // Neighbor creation. The type is interface{}.
    Enable interface{}

    // Neighbor router ID (IPv4 Address). The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborRouterId interface{}

    // IPCC configuration.
    Ipcc Lmp_GmplsUni_Neighbors_Neighbor_Ipcc
}

func (neighbor *Lmp_GmplsUni_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborName, "neighbor-name")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("ipcc", types.YChild{"Ipcc", &neighbor.Ipcc})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-name", types.YLeaf{"NeighborName", neighbor.NeighborName})
    neighbor.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", neighbor.Enable})
    neighbor.EntityData.Leafs.Append("neighbor-router-id", types.YLeaf{"NeighborRouterId", neighbor.NeighborRouterId})

    neighbor.EntityData.YListKeys = []string {"NeighborName"}

    return &(neighbor.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_Ipcc
// IPCC configuration
type Lmp_GmplsUni_Neighbors_Neighbor_Ipcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Routed IPCC configuration.
    Routed Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_Routed
}

func (ipcc *Lmp_GmplsUni_Neighbors_Neighbor_Ipcc) GetEntityData() *types.CommonEntityData {
    ipcc.EntityData.YFilter = ipcc.YFilter
    ipcc.EntityData.YangName = "ipcc"
    ipcc.EntityData.BundleName = "cisco_ios_xr"
    ipcc.EntityData.ParentYangName = "neighbor"
    ipcc.EntityData.SegmentPath = "ipcc"
    ipcc.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/neighbors/neighbor/" + ipcc.EntityData.SegmentPath
    ipcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcc.EntityData.Children = types.NewOrderedMap()
    ipcc.EntityData.Children.Append("routed", types.YChild{"Routed", &ipcc.Routed})
    ipcc.EntityData.Leafs = types.NewOrderedMap()

    ipcc.EntityData.YListKeys = []string {}

    return &(ipcc.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_Routed
// Routed IPCC configuration
type Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_Routed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Routed IPCC creation. The type is interface{}.
    Enable interface{}
}

func (routed *Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_Routed) GetEntityData() *types.CommonEntityData {
    routed.EntityData.YFilter = routed.YFilter
    routed.EntityData.YangName = "routed"
    routed.EntityData.BundleName = "cisco_ios_xr"
    routed.EntityData.ParentYangName = "ipcc"
    routed.EntityData.SegmentPath = "routed"
    routed.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/neighbors/neighbor/ipcc/" + routed.EntityData.SegmentPath
    routed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routed.EntityData.Children = types.NewOrderedMap()
    routed.EntityData.Leafs = types.NewOrderedMap()
    routed.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", routed.Enable})

    routed.EntityData.YListKeys = []string {}

    return &(routed.EntityData)
}

// Lmp_GmplsUni_RouterId
// Local GMPLS UNI router ID
// This type is a presence type.
type Lmp_GmplsUni_RouterId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Local router ID (IPv4 Address). The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (routerId *Lmp_GmplsUni_RouterId) GetEntityData() *types.CommonEntityData {
    routerId.EntityData.YFilter = routerId.YFilter
    routerId.EntityData.YangName = "router-id"
    routerId.EntityData.BundleName = "cisco_ios_xr"
    routerId.EntityData.ParentYangName = "gmpls-uni"
    routerId.EntityData.SegmentPath = "router-id"
    routerId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/" + routerId.EntityData.SegmentPath
    routerId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerId.EntityData.Children = types.NewOrderedMap()
    routerId.EntityData.Leafs = types.NewOrderedMap()
    routerId.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", routerId.InterfaceName})
    routerId.EntityData.Leafs.Append("address", types.YLeaf{"Address", routerId.Address})

    routerId.EntityData.YListKeys = []string {}

    return &(routerId.EntityData)
}

// Lmp_GmplsUni_Controllers
// Configure GMPLS UNI OLM/LMP controllers
type Lmp_GmplsUni_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure an GMPLS UNI OLM/LMP contoller. The type is slice of
    // Lmp_GmplsUni_Controllers_Controller.
    Controller []*Lmp_GmplsUni_Controllers_Controller
}

func (controllers *Lmp_GmplsUni_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "gmpls-uni"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/" + controllers.EntityData.SegmentPath
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

// Lmp_GmplsUni_Controllers_Controller
// Configure an GMPLS UNI OLM/LMP contoller
type Lmp_GmplsUni_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Controller name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    ControllerName interface{}

    // Enable the OLM/LMP application on this controller. The type is interface{}.
    Enable interface{}

    // Local Link ID configuration.
    LocalLinkId Lmp_GmplsUni_Controllers_Controller_LocalLinkId

    // Neighbor controller adjacency configuration.
    Adjacency Lmp_GmplsUni_Controllers_Controller_Adjacency
}

func (controller *Lmp_GmplsUni_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/controllers/" + controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("local-link-id", types.YChild{"LocalLinkId", &controller.LocalLinkId})
    controller.EntityData.Children.Append("adjacency", types.YChild{"Adjacency", &controller.Adjacency})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})
    controller.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", controller.Enable})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// Lmp_GmplsUni_Controllers_Controller_LocalLinkId
// Local Link ID configuration
type Lmp_GmplsUni_Controllers_Controller_LocalLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local link ID address type. The type is OlmAddr.
    AddressType interface{}

    // Local address unnumbered . The type is interface{} with range:
    // 0..4294967295.
    Unnumbered interface{}

    // Local link ID address IPv4. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (localLinkId *Lmp_GmplsUni_Controllers_Controller_LocalLinkId) GetEntityData() *types.CommonEntityData {
    localLinkId.EntityData.YFilter = localLinkId.YFilter
    localLinkId.EntityData.YangName = "local-link-id"
    localLinkId.EntityData.BundleName = "cisco_ios_xr"
    localLinkId.EntityData.ParentYangName = "controller"
    localLinkId.EntityData.SegmentPath = "local-link-id"
    localLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/controllers/controller/" + localLinkId.EntityData.SegmentPath
    localLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLinkId.EntityData.Children = types.NewOrderedMap()
    localLinkId.EntityData.Leafs = types.NewOrderedMap()
    localLinkId.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", localLinkId.AddressType})
    localLinkId.EntityData.Leafs.Append("unnumbered", types.YLeaf{"Unnumbered", localLinkId.Unnumbered})
    localLinkId.EntityData.Leafs.Append("address", types.YLeaf{"Address", localLinkId.Address})

    localLinkId.EntityData.YListKeys = []string {}

    return &(localLinkId.EntityData)
}

// Lmp_GmplsUni_Controllers_Controller_Adjacency
// Neighbor controller adjacency configuration
type Lmp_GmplsUni_Controllers_Controller_Adjacency struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor data.
    RemoteNeighbor Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor
}

func (adjacency *Lmp_GmplsUni_Controllers_Controller_Adjacency) GetEntityData() *types.CommonEntityData {
    adjacency.EntityData.YFilter = adjacency.YFilter
    adjacency.EntityData.YangName = "adjacency"
    adjacency.EntityData.BundleName = "cisco_ios_xr"
    adjacency.EntityData.ParentYangName = "controller"
    adjacency.EntityData.SegmentPath = "adjacency"
    adjacency.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/controllers/controller/" + adjacency.EntityData.SegmentPath
    adjacency.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacency.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacency.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacency.EntityData.Children = types.NewOrderedMap()
    adjacency.EntityData.Children.Append("remote-neighbor", types.YChild{"RemoteNeighbor", &adjacency.RemoteNeighbor})
    adjacency.EntityData.Leafs = types.NewOrderedMap()

    adjacency.EntityData.YListKeys = []string {}

    return &(adjacency.EntityData)
}

// Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor
// Neighbor data
type Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create LMP controller to neighbor association. The type is string with
    // length: 1..64.
    NeighborAssociation interface{}

    // Neighbor link switching capability configuration. The type is
    // OlmSwitchingCap. The default value is lsc.
    LinkSwitchingCapability interface{}

    // Remote node flexi grid capability . The type is interface{} with range:
    // 0..4294967295.
    FlexiGridCapable interface{}

    // Neighbor Interface ID configuration.
    InterfaceId Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_InterfaceId

    // Neighbor Link ID configuration.
    LinkId Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_LinkId
}

func (remoteNeighbor *Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor) GetEntityData() *types.CommonEntityData {
    remoteNeighbor.EntityData.YFilter = remoteNeighbor.YFilter
    remoteNeighbor.EntityData.YangName = "remote-neighbor"
    remoteNeighbor.EntityData.BundleName = "cisco_ios_xr"
    remoteNeighbor.EntityData.ParentYangName = "adjacency"
    remoteNeighbor.EntityData.SegmentPath = "remote-neighbor"
    remoteNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/controllers/controller/adjacency/" + remoteNeighbor.EntityData.SegmentPath
    remoteNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNeighbor.EntityData.Children = types.NewOrderedMap()
    remoteNeighbor.EntityData.Children.Append("interface-id", types.YChild{"InterfaceId", &remoteNeighbor.InterfaceId})
    remoteNeighbor.EntityData.Children.Append("link-id", types.YChild{"LinkId", &remoteNeighbor.LinkId})
    remoteNeighbor.EntityData.Leafs = types.NewOrderedMap()
    remoteNeighbor.EntityData.Leafs.Append("neighbor-association", types.YLeaf{"NeighborAssociation", remoteNeighbor.NeighborAssociation})
    remoteNeighbor.EntityData.Leafs.Append("link-switching-capability", types.YLeaf{"LinkSwitchingCapability", remoteNeighbor.LinkSwitchingCapability})
    remoteNeighbor.EntityData.Leafs.Append("flexi-grid-capable", types.YLeaf{"FlexiGridCapable", remoteNeighbor.FlexiGridCapable})

    remoteNeighbor.EntityData.YListKeys = []string {}

    return &(remoteNeighbor.EntityData)
}

// Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_InterfaceId
// Neighbor Interface ID configuration
type Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_InterfaceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local link ID address type. The type is OlmAddr.
    AddressType interface{}

    // Local address unnumbered . The type is interface{} with range:
    // 0..4294967295.
    Unnumbered interface{}

    // Local link ID address IPv4. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (interfaceId *Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_InterfaceId) GetEntityData() *types.CommonEntityData {
    interfaceId.EntityData.YFilter = interfaceId.YFilter
    interfaceId.EntityData.YangName = "interface-id"
    interfaceId.EntityData.BundleName = "cisco_ios_xr"
    interfaceId.EntityData.ParentYangName = "remote-neighbor"
    interfaceId.EntityData.SegmentPath = "interface-id"
    interfaceId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/controllers/controller/adjacency/remote-neighbor/" + interfaceId.EntityData.SegmentPath
    interfaceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceId.EntityData.Children = types.NewOrderedMap()
    interfaceId.EntityData.Leafs = types.NewOrderedMap()
    interfaceId.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", interfaceId.AddressType})
    interfaceId.EntityData.Leafs.Append("unnumbered", types.YLeaf{"Unnumbered", interfaceId.Unnumbered})
    interfaceId.EntityData.Leafs.Append("address", types.YLeaf{"Address", interfaceId.Address})

    interfaceId.EntityData.YListKeys = []string {}

    return &(interfaceId.EntityData)
}

// Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_LinkId
// Neighbor Link ID configuration
type Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_LinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor link ID address type. The type is OlmAddr.
    AddressType interface{}

    // Neighbor address unnumbered [Not supported]. The type is interface{} with
    // range: 0..4294967295.
    Unnumbered interface{}

    // Neighbor ID address IPv4. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (linkId *Lmp_GmplsUni_Controllers_Controller_Adjacency_RemoteNeighbor_LinkId) GetEntityData() *types.CommonEntityData {
    linkId.EntityData.YFilter = linkId.YFilter
    linkId.EntityData.YangName = "link-id"
    linkId.EntityData.BundleName = "cisco_ios_xr"
    linkId.EntityData.ParentYangName = "remote-neighbor"
    linkId.EntityData.SegmentPath = "link-id"
    linkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-cfg:lmp/gmpls-uni/controllers/controller/adjacency/remote-neighbor/" + linkId.EntityData.SegmentPath
    linkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkId.EntityData.Children = types.NewOrderedMap()
    linkId.EntityData.Leafs = types.NewOrderedMap()
    linkId.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", linkId.AddressType})
    linkId.EntityData.Leafs.Append("unnumbered", types.YLeaf{"Unnumbered", linkId.Unnumbered})
    linkId.EntityData.Leafs.Append("address", types.YLeaf{"Address", linkId.Address})

    linkId.EntityData.YListKeys = []string {}

    return &(linkId.EntityData)
}

