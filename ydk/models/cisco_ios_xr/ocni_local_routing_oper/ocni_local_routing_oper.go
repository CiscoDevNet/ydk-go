// This module contains a collection of YANG definitions
// for Cisco IOS-XR ocni-local-routing package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ocni: An OpenConfig description of a network-instance. This
//     may be a Layer 3 forwarding construct such as a virtual
//     routing and forwarding (VRF) instance, or a Layer 2 instance
//     such as a virtual switch instance (VSI). Mixed Layer 2 and
//     Layer 3 instances are also supported.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ocni_local_routing_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ocni_local_routing_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ocni-local-routing-oper ocni}", reflect.TypeOf(Ocni{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ocni-local-routing-oper:ocni", reflect.TypeOf(Ocni{}))
}

// Ocni
// An OpenConfig description of a network-instance.
// This may be a Layer 3 forwarding construct such
// as a virtual routing and forwarding (VRF)
// instance, or a Layer 2 instance such as a virtual
// switch instance (VSI). Mixed Layer 2 and Layer 3
// instances are also supported.
type Ocni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 static configuration.
    Vrfipv4 Ocni_Vrfipv4

    // IPv6 static configuration.
    Vrfipv6 Ocni_Vrfipv6
}

func (ocni *Ocni) GetEntityData() *types.CommonEntityData {
    ocni.EntityData.YFilter = ocni.YFilter
    ocni.EntityData.YangName = "ocni"
    ocni.EntityData.BundleName = "cisco_ios_xr"
    ocni.EntityData.ParentYangName = "Cisco-IOS-XR-ocni-local-routing-oper"
    ocni.EntityData.SegmentPath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni"
    ocni.EntityData.AbsolutePath = ocni.EntityData.SegmentPath
    ocni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ocni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ocni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ocni.EntityData.Children = types.NewOrderedMap()
    ocni.EntityData.Children.Append("vrfipv4", types.YChild{"Vrfipv4", &ocni.Vrfipv4})
    ocni.EntityData.Children.Append("vrfipv6", types.YChild{"Vrfipv6", &ocni.Vrfipv6})
    ocni.EntityData.Leafs = types.NewOrderedMap()

    ocni.EntityData.YListKeys = []string {}

    return &(ocni.EntityData)
}

// Ocni_Vrfipv4
// IPv4 static configuration
type Ocni_Vrfipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system.
    NetworkInstances Ocni_Vrfipv4_NetworkInstances
}

func (vrfipv4 *Ocni_Vrfipv4) GetEntityData() *types.CommonEntityData {
    vrfipv4.EntityData.YFilter = vrfipv4.YFilter
    vrfipv4.EntityData.YangName = "vrfipv4"
    vrfipv4.EntityData.BundleName = "cisco_ios_xr"
    vrfipv4.EntityData.ParentYangName = "ocni"
    vrfipv4.EntityData.SegmentPath = "vrfipv4"
    vrfipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/" + vrfipv4.EntityData.SegmentPath
    vrfipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfipv4.EntityData.Children = types.NewOrderedMap()
    vrfipv4.EntityData.Children.Append("network-instances", types.YChild{"NetworkInstances", &vrfipv4.NetworkInstances})
    vrfipv4.EntityData.Leafs = types.NewOrderedMap()

    vrfipv4.EntityData.YListKeys = []string {}

    return &(vrfipv4.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances
// Network instances configured on the local system
type Ocni_Vrfipv4_NetworkInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system. The type is slice of
    // Ocni_Vrfipv4_NetworkInstances_NetworkInstance.
    NetworkInstance []*Ocni_Vrfipv4_NetworkInstances_NetworkInstance
}

func (networkInstances *Ocni_Vrfipv4_NetworkInstances) GetEntityData() *types.CommonEntityData {
    networkInstances.EntityData.YFilter = networkInstances.YFilter
    networkInstances.EntityData.YangName = "network-instances"
    networkInstances.EntityData.BundleName = "cisco_ios_xr"
    networkInstances.EntityData.ParentYangName = "vrfipv4"
    networkInstances.EntityData.SegmentPath = "network-instances"
    networkInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/" + networkInstances.EntityData.SegmentPath
    networkInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkInstances.EntityData.Children = types.NewOrderedMap()
    networkInstances.EntityData.Children.Append("network-instance", types.YChild{"NetworkInstance", nil})
    for i := range networkInstances.NetworkInstance {
        networkInstances.EntityData.Children.Append(types.GetSegmentPath(networkInstances.NetworkInstance[i]), types.YChild{"NetworkInstance", networkInstances.NetworkInstance[i]})
    }
    networkInstances.EntityData.Leafs = types.NewOrderedMap()

    networkInstances.EntityData.YListKeys = []string {}

    return &(networkInstances.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance
// Network instances configured on the local
// system
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique name identifying the network instance.
    // The type is string.
    Name interface{}

    // A process (instance) of a routing protocol. Some systems may not support
    // more than one instance of a particular routing protocol.
    Protocols Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols
}

func (networkInstance *Ocni_Vrfipv4_NetworkInstances_NetworkInstance) GetEntityData() *types.CommonEntityData {
    networkInstance.EntityData.YFilter = networkInstance.YFilter
    networkInstance.EntityData.YangName = "network-instance"
    networkInstance.EntityData.BundleName = "cisco_ios_xr"
    networkInstance.EntityData.ParentYangName = "network-instances"
    networkInstance.EntityData.SegmentPath = "network-instance" + types.AddKeyToken(networkInstance.Name, "name")
    networkInstance.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/" + networkInstance.EntityData.SegmentPath
    networkInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkInstance.EntityData.Children = types.NewOrderedMap()
    networkInstance.EntityData.Children.Append("protocols", types.YChild{"Protocols", &networkInstance.Protocols})
    networkInstance.EntityData.Leafs = types.NewOrderedMap()
    networkInstance.EntityData.Leafs.Append("name", types.YLeaf{"Name", networkInstance.Name})

    networkInstance.EntityData.YListKeys = []string {"Name"}

    return &(networkInstance.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols
// A process (instance) of a routing protocol.
// Some systems may not support more than one
// instance of a particular routing protocol
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A process (instance) of a routing protocol. Some systems may not support
    // more than one instance of a particular routing protocol. The type is slice
    // of Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol.
    Protocol []*Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol
}

func (protocols *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "network-instance"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/" + protocols.EntityData.SegmentPath
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range protocols.Protocol {
        types.SetYListKey(protocols.Protocol[i], i)
        protocols.EntityData.Children.Append(types.GetSegmentPath(protocols.Protocol[i]), types.YChild{"Protocol", protocols.Protocol[i]})
    }
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol
// A process (instance) of a routing protocol.
// Some systems may not support more than one
// instance of a particular routing protocol
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The protocol name for the routing or forwarding protocol to be
    // instantiated. The type is string.
    Identifier interface{}

    // An operator-assigned identifier for the routing or forwarding protocol. For
    // some processes this leaf may be system defined. The type is string.
    Name interface{}

    // State parameters relating to the routing protocol instance.
    State Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_State

    // List of locally configured static routes.
    StaticRoutes Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes
}

func (protocol *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "protocols"
    protocol.EntityData.SegmentPath = "protocol" + types.AddNoKeyToken(protocol)
    protocol.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/" + protocol.EntityData.SegmentPath
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Children.Append("state", types.YChild{"State", &protocol.State})
    protocol.EntityData.Children.Append("static-routes", types.YChild{"StaticRoutes", &protocol.StaticRoutes})
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", protocol.Identifier})
    protocol.EntityData.Leafs.Append("name", types.YLeaf{"Name", protocol.Name})

    protocol.EntityData.YListKeys = []string {}

    return &(protocol.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_State
// State parameters relating to the routing
// protocol instance
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The protocol identifier for the instance. The type is string.
    Identifier interface{}

    // A unique name for the protocol instance. The type is string.
    Name interface{}

    // A boolean value indicating whether the local protocol instance is enabled.
    // The type is bool.
    Enabled interface{}

    // The default metric within the RIB for entries that are installed by this
    // protocol instance. This value may be overridden by protocol specific
    // configuration options. The lower the metric specified the more preferable
    // the RIB entry is to be selected for use within the network instance. Where
    // multiple entries have the same metric value then these equal cost paths
    // should be treated according to the specified ECMP path selection behaviour
    // for the instance. The type is interface{} with range: 0..4294967295.
    DefaultMetric interface{}
}

func (state *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "protocol"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", state.Identifier})
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", state.DefaultMetric})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes
// List of locally configured static routes
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of locally configured static routes. The type is slice of
    // Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute.
    StaticRoute []*Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute
}

func (staticRoutes *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes) GetEntityData() *types.CommonEntityData {
    staticRoutes.EntityData.YFilter = staticRoutes.YFilter
    staticRoutes.EntityData.YangName = "static-routes"
    staticRoutes.EntityData.BundleName = "cisco_ios_xr"
    staticRoutes.EntityData.ParentYangName = "protocol"
    staticRoutes.EntityData.SegmentPath = "static-routes"
    staticRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/" + staticRoutes.EntityData.SegmentPath
    staticRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRoutes.EntityData.Children = types.NewOrderedMap()
    staticRoutes.EntityData.Children.Append("static-route", types.YChild{"StaticRoute", nil})
    for i := range staticRoutes.StaticRoute {
        staticRoutes.EntityData.Children.Append(types.GetSegmentPath(staticRoutes.StaticRoute[i]), types.YChild{"StaticRoute", staticRoutes.StaticRoute[i]})
    }
    staticRoutes.EntityData.Leafs = types.NewOrderedMap()

    staticRoutes.EntityData.YListKeys = []string {}

    return &(staticRoutes.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute
// List of locally configured static routes
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the destination prefix list key. The
    // type is string.
    Prefix interface{}

    // A list of next-hops to be utilised for the static route being specified.
    NextHops Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops

    // Operational state data for static routes.
    StaticRoutesState Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState
}

func (staticRoute *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute) GetEntityData() *types.CommonEntityData {
    staticRoute.EntityData.YFilter = staticRoute.YFilter
    staticRoute.EntityData.YangName = "static-route"
    staticRoute.EntityData.BundleName = "cisco_ios_xr"
    staticRoute.EntityData.ParentYangName = "static-routes"
    staticRoute.EntityData.SegmentPath = "static-route" + types.AddKeyToken(staticRoute.Prefix, "prefix")
    staticRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/static-routes/" + staticRoute.EntityData.SegmentPath
    staticRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRoute.EntityData.Children = types.NewOrderedMap()
    staticRoute.EntityData.Children.Append("next-hops", types.YChild{"NextHops", &staticRoute.NextHops})
    staticRoute.EntityData.Children.Append("static-routes-state", types.YChild{"StaticRoutesState", &staticRoute.StaticRoutesState})
    staticRoute.EntityData.Leafs = types.NewOrderedMap()
    staticRoute.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", staticRoute.Prefix})

    staticRoute.EntityData.YListKeys = []string {"Prefix"}

    return &(staticRoute.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops
// A list of next-hops to be utilised for the
// static route being specified.
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of next-hops to be utilised for the static route being specified.
    // The type is slice of
    // Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop.
    NextHop []*Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop
}

func (nextHops *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops) GetEntityData() *types.CommonEntityData {
    nextHops.EntityData.YFilter = nextHops.YFilter
    nextHops.EntityData.YangName = "next-hops"
    nextHops.EntityData.BundleName = "cisco_ios_xr"
    nextHops.EntityData.ParentYangName = "static-route"
    nextHops.EntityData.SegmentPath = "next-hops"
    nextHops.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/static-routes/static-route/" + nextHops.EntityData.SegmentPath
    nextHops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHops.EntityData.Children = types.NewOrderedMap()
    nextHops.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range nextHops.NextHop {
        nextHops.EntityData.Children.Append(types.GetSegmentPath(nextHops.NextHop[i]), types.YChild{"NextHop", nextHops.NextHop[i]})
    }
    nextHops.EntityData.Leafs = types.NewOrderedMap()

    nextHops.EntityData.YListKeys = []string {}

    return &(nextHops.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop
// A list of next-hops to be utilised for
// the static route being specified.
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A reference to the index of the current next-hop.
    // The index is intended to be a user-specified value which can be used to
    // reference the next-hop in question, without any other semantics being
    // assigned to it. The type is string.
    Index interface{}

    // Operational state parameters relating to the next-hop entry.
    State Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State

    // Reference to an interface or subinterface.
    InterfaceRef Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef
}

func (nextHop *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hops"
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("state", types.YChild{"State", &nextHop.State})
    nextHop.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &nextHop.InterfaceRef})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})

    nextHop.EntityData.YListKeys = []string {"Index"}

    return &(nextHop.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State
// Operational state parameters relating to the
// next-hop entry
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An user-specified identifier utilised to uniquely reference the next-hop
    // entry in the next-hop list. The value of this index has no semantic meaning
    // other than for referencing the entry. The type is string.
    Index interface{}

    // The next-hop that is to be used for the static route - this may be
    // specified as an IP address, an interface or a pre-defined next-hop type -
    // for instance, DROP or LOCAL_LINK. When this leaf is not set, and the
    // interface-ref value is specified for the next-hop, then the system should
    // treat the prefix as though it is directly connected to the interface. The
    // type is string.
    NextHop interface{}

    // A metric which is utilised to specify the preference of the next-hop entry
    // when it is injected into the RIB. The lower the metric, the more preferable
    // the prefix is. When this value is not specified the metric is inherited
    // from the default metric utilised for static routes within the network
    // instance that the static routes are being instantiated. When multiple
    // next-hops are specified for a static route, the metric is utilised to
    // determine which of the next-hops is to be installed in the RIB. When
    // multiple next-hops have the same metric (be it specified, or simply the
    // default) then these next-hops should all be installed in the RIB. The type
    // is interface{} with range: 0..4294967295.
    Metric interface{}
}

func (state *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "next-hop"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/next-hop/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", state.NextHop})
    state.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", state.Metric})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef
// Reference to an interface or subinterface
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational state for interface-ref.
    State Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State
}

func (interfaceRef *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "cisco_ios_xr"
    interfaceRef.EntityData.ParentYangName = "next-hop"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/next-hop/" + interfaceRef.EntityData.SegmentPath
    interfaceRef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State
// Operational state for interface-ref
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string.
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // interface{} with range: 0..4294967295.
    Subinterface interface{}
}

func (state *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/next-hop/interface-ref/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState
// Operational state data for static routes
type Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination prefix for the static route, either IPv4 or IPv6. The type is
    // string.
    Prefix interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is string.
    SetTag interface{}
}

func (staticRoutesState *Ocni_Vrfipv4_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState) GetEntityData() *types.CommonEntityData {
    staticRoutesState.EntityData.YFilter = staticRoutesState.YFilter
    staticRoutesState.EntityData.YangName = "static-routes-state"
    staticRoutesState.EntityData.BundleName = "cisco_ios_xr"
    staticRoutesState.EntityData.ParentYangName = "static-route"
    staticRoutesState.EntityData.SegmentPath = "static-routes-state"
    staticRoutesState.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv4/network-instances/network-instance/protocols/protocol/static-routes/static-route/" + staticRoutesState.EntityData.SegmentPath
    staticRoutesState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRoutesState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRoutesState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRoutesState.EntityData.Children = types.NewOrderedMap()
    staticRoutesState.EntityData.Leafs = types.NewOrderedMap()
    staticRoutesState.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", staticRoutesState.Prefix})
    staticRoutesState.EntityData.Leafs.Append("set-tag", types.YLeaf{"SetTag", staticRoutesState.SetTag})

    staticRoutesState.EntityData.YListKeys = []string {}

    return &(staticRoutesState.EntityData)
}

// Ocni_Vrfipv6
// IPv6 static configuration
type Ocni_Vrfipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system.
    NetworkInstances Ocni_Vrfipv6_NetworkInstances
}

func (vrfipv6 *Ocni_Vrfipv6) GetEntityData() *types.CommonEntityData {
    vrfipv6.EntityData.YFilter = vrfipv6.YFilter
    vrfipv6.EntityData.YangName = "vrfipv6"
    vrfipv6.EntityData.BundleName = "cisco_ios_xr"
    vrfipv6.EntityData.ParentYangName = "ocni"
    vrfipv6.EntityData.SegmentPath = "vrfipv6"
    vrfipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/" + vrfipv6.EntityData.SegmentPath
    vrfipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfipv6.EntityData.Children = types.NewOrderedMap()
    vrfipv6.EntityData.Children.Append("network-instances", types.YChild{"NetworkInstances", &vrfipv6.NetworkInstances})
    vrfipv6.EntityData.Leafs = types.NewOrderedMap()

    vrfipv6.EntityData.YListKeys = []string {}

    return &(vrfipv6.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances
// Network instances configured on the local system
type Ocni_Vrfipv6_NetworkInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system. The type is slice of
    // Ocni_Vrfipv6_NetworkInstances_NetworkInstance.
    NetworkInstance []*Ocni_Vrfipv6_NetworkInstances_NetworkInstance
}

func (networkInstances *Ocni_Vrfipv6_NetworkInstances) GetEntityData() *types.CommonEntityData {
    networkInstances.EntityData.YFilter = networkInstances.YFilter
    networkInstances.EntityData.YangName = "network-instances"
    networkInstances.EntityData.BundleName = "cisco_ios_xr"
    networkInstances.EntityData.ParentYangName = "vrfipv6"
    networkInstances.EntityData.SegmentPath = "network-instances"
    networkInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/" + networkInstances.EntityData.SegmentPath
    networkInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkInstances.EntityData.Children = types.NewOrderedMap()
    networkInstances.EntityData.Children.Append("network-instance", types.YChild{"NetworkInstance", nil})
    for i := range networkInstances.NetworkInstance {
        networkInstances.EntityData.Children.Append(types.GetSegmentPath(networkInstances.NetworkInstance[i]), types.YChild{"NetworkInstance", networkInstances.NetworkInstance[i]})
    }
    networkInstances.EntityData.Leafs = types.NewOrderedMap()

    networkInstances.EntityData.YListKeys = []string {}

    return &(networkInstances.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance
// Network instances configured on the local
// system
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique name identifying the network instance.
    // The type is string.
    Name interface{}

    // A process (instance) of a routing protocol. Some systems may not support
    // more than one instance of a particular routing protocol.
    Protocols Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols
}

func (networkInstance *Ocni_Vrfipv6_NetworkInstances_NetworkInstance) GetEntityData() *types.CommonEntityData {
    networkInstance.EntityData.YFilter = networkInstance.YFilter
    networkInstance.EntityData.YangName = "network-instance"
    networkInstance.EntityData.BundleName = "cisco_ios_xr"
    networkInstance.EntityData.ParentYangName = "network-instances"
    networkInstance.EntityData.SegmentPath = "network-instance" + types.AddKeyToken(networkInstance.Name, "name")
    networkInstance.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/" + networkInstance.EntityData.SegmentPath
    networkInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkInstance.EntityData.Children = types.NewOrderedMap()
    networkInstance.EntityData.Children.Append("protocols", types.YChild{"Protocols", &networkInstance.Protocols})
    networkInstance.EntityData.Leafs = types.NewOrderedMap()
    networkInstance.EntityData.Leafs.Append("name", types.YLeaf{"Name", networkInstance.Name})

    networkInstance.EntityData.YListKeys = []string {"Name"}

    return &(networkInstance.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols
// A process (instance) of a routing protocol.
// Some systems may not support more than one
// instance of a particular routing protocol
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A process (instance) of a routing protocol. Some systems may not support
    // more than one instance of a particular routing protocol. The type is slice
    // of Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol.
    Protocol []*Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol
}

func (protocols *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "network-instance"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/" + protocols.EntityData.SegmentPath
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range protocols.Protocol {
        types.SetYListKey(protocols.Protocol[i], i)
        protocols.EntityData.Children.Append(types.GetSegmentPath(protocols.Protocol[i]), types.YChild{"Protocol", protocols.Protocol[i]})
    }
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol
// A process (instance) of a routing protocol.
// Some systems may not support more than one
// instance of a particular routing protocol
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The protocol name for the routing or forwarding protocol to be
    // instantiated. The type is string.
    Identifier interface{}

    // An operator-assigned identifier for the routing or forwarding protocol. For
    // some processes this leaf may be system defined. The type is string.
    Name interface{}

    // State parameters relating to the routing protocol instance.
    State Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_State

    // List of locally configured static routes.
    StaticRoutes Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes
}

func (protocol *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "protocols"
    protocol.EntityData.SegmentPath = "protocol" + types.AddNoKeyToken(protocol)
    protocol.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/" + protocol.EntityData.SegmentPath
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Children.Append("state", types.YChild{"State", &protocol.State})
    protocol.EntityData.Children.Append("static-routes", types.YChild{"StaticRoutes", &protocol.StaticRoutes})
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", protocol.Identifier})
    protocol.EntityData.Leafs.Append("name", types.YLeaf{"Name", protocol.Name})

    protocol.EntityData.YListKeys = []string {}

    return &(protocol.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_State
// State parameters relating to the routing
// protocol instance
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The protocol identifier for the instance. The type is string.
    Identifier interface{}

    // A unique name for the protocol instance. The type is string.
    Name interface{}

    // A boolean value indicating whether the local protocol instance is enabled.
    // The type is bool.
    Enabled interface{}

    // The default metric within the RIB for entries that are installed by this
    // protocol instance. This value may be overridden by protocol specific
    // configuration options. The lower the metric specified the more preferable
    // the RIB entry is to be selected for use within the network instance. Where
    // multiple entries have the same metric value then these equal cost paths
    // should be treated according to the specified ECMP path selection behaviour
    // for the instance. The type is interface{} with range: 0..4294967295.
    DefaultMetric interface{}
}

func (state *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "protocol"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", state.Identifier})
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", state.DefaultMetric})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes
// List of locally configured static routes
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of locally configured static routes. The type is slice of
    // Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute.
    StaticRoute []*Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute
}

func (staticRoutes *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes) GetEntityData() *types.CommonEntityData {
    staticRoutes.EntityData.YFilter = staticRoutes.YFilter
    staticRoutes.EntityData.YangName = "static-routes"
    staticRoutes.EntityData.BundleName = "cisco_ios_xr"
    staticRoutes.EntityData.ParentYangName = "protocol"
    staticRoutes.EntityData.SegmentPath = "static-routes"
    staticRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/" + staticRoutes.EntityData.SegmentPath
    staticRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRoutes.EntityData.Children = types.NewOrderedMap()
    staticRoutes.EntityData.Children.Append("static-route", types.YChild{"StaticRoute", nil})
    for i := range staticRoutes.StaticRoute {
        staticRoutes.EntityData.Children.Append(types.GetSegmentPath(staticRoutes.StaticRoute[i]), types.YChild{"StaticRoute", staticRoutes.StaticRoute[i]})
    }
    staticRoutes.EntityData.Leafs = types.NewOrderedMap()

    staticRoutes.EntityData.YListKeys = []string {}

    return &(staticRoutes.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute
// List of locally configured static routes
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the destination prefix list key. The
    // type is string.
    Prefix interface{}

    // A list of next-hops to be utilised for the static route being specified.
    NextHops Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops

    // Operational state data for static routes.
    StaticRoutesState Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState
}

func (staticRoute *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute) GetEntityData() *types.CommonEntityData {
    staticRoute.EntityData.YFilter = staticRoute.YFilter
    staticRoute.EntityData.YangName = "static-route"
    staticRoute.EntityData.BundleName = "cisco_ios_xr"
    staticRoute.EntityData.ParentYangName = "static-routes"
    staticRoute.EntityData.SegmentPath = "static-route" + types.AddKeyToken(staticRoute.Prefix, "prefix")
    staticRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/static-routes/" + staticRoute.EntityData.SegmentPath
    staticRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRoute.EntityData.Children = types.NewOrderedMap()
    staticRoute.EntityData.Children.Append("next-hops", types.YChild{"NextHops", &staticRoute.NextHops})
    staticRoute.EntityData.Children.Append("static-routes-state", types.YChild{"StaticRoutesState", &staticRoute.StaticRoutesState})
    staticRoute.EntityData.Leafs = types.NewOrderedMap()
    staticRoute.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", staticRoute.Prefix})

    staticRoute.EntityData.YListKeys = []string {"Prefix"}

    return &(staticRoute.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops
// A list of next-hops to be utilised for the
// static route being specified.
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of next-hops to be utilised for the static route being specified.
    // The type is slice of
    // Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop.
    NextHop []*Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop
}

func (nextHops *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops) GetEntityData() *types.CommonEntityData {
    nextHops.EntityData.YFilter = nextHops.YFilter
    nextHops.EntityData.YangName = "next-hops"
    nextHops.EntityData.BundleName = "cisco_ios_xr"
    nextHops.EntityData.ParentYangName = "static-route"
    nextHops.EntityData.SegmentPath = "next-hops"
    nextHops.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/static-routes/static-route/" + nextHops.EntityData.SegmentPath
    nextHops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHops.EntityData.Children = types.NewOrderedMap()
    nextHops.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range nextHops.NextHop {
        nextHops.EntityData.Children.Append(types.GetSegmentPath(nextHops.NextHop[i]), types.YChild{"NextHop", nextHops.NextHop[i]})
    }
    nextHops.EntityData.Leafs = types.NewOrderedMap()

    nextHops.EntityData.YListKeys = []string {}

    return &(nextHops.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop
// A list of next-hops to be utilised for
// the static route being specified.
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A reference to the index of the current next-hop.
    // The index is intended to be a user-specified value which can be used to
    // reference the next-hop in question, without any other semantics being
    // assigned to it. The type is string.
    Index interface{}

    // Operational state parameters relating to the next-hop entry.
    State Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State

    // Reference to an interface or subinterface.
    InterfaceRef Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef
}

func (nextHop *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hops"
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("state", types.YChild{"State", &nextHop.State})
    nextHop.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &nextHop.InterfaceRef})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})

    nextHop.EntityData.YListKeys = []string {"Index"}

    return &(nextHop.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State
// Operational state parameters relating to the
// next-hop entry
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An user-specified identifier utilised to uniquely reference the next-hop
    // entry in the next-hop list. The value of this index has no semantic meaning
    // other than for referencing the entry. The type is string.
    Index interface{}

    // The next-hop that is to be used for the static route - this may be
    // specified as an IP address, an interface or a pre-defined next-hop type -
    // for instance, DROP or LOCAL_LINK. When this leaf is not set, and the
    // interface-ref value is specified for the next-hop, then the system should
    // treat the prefix as though it is directly connected to the interface. The
    // type is string.
    NextHop interface{}

    // A metric which is utilised to specify the preference of the next-hop entry
    // when it is injected into the RIB. The lower the metric, the more preferable
    // the prefix is. When this value is not specified the metric is inherited
    // from the default metric utilised for static routes within the network
    // instance that the static routes are being instantiated. When multiple
    // next-hops are specified for a static route, the metric is utilised to
    // determine which of the next-hops is to be installed in the RIB. When
    // multiple next-hops have the same metric (be it specified, or simply the
    // default) then these next-hops should all be installed in the RIB. The type
    // is interface{} with range: 0..4294967295.
    Metric interface{}
}

func (state *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "next-hop"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/next-hop/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", state.NextHop})
    state.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", state.Metric})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef
// Reference to an interface or subinterface
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational state for interface-ref.
    State Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State
}

func (interfaceRef *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "cisco_ios_xr"
    interfaceRef.EntityData.ParentYangName = "next-hop"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/next-hop/" + interfaceRef.EntityData.SegmentPath
    interfaceRef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State
// Operational state for interface-ref
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string.
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // interface{} with range: 0..4294967295.
    Subinterface interface{}
}

func (state *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_NextHops_NextHop_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/static-routes/static-route/next-hops/next-hop/interface-ref/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState
// Operational state data for static routes
type Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination prefix for the static route, either IPv4 or IPv6. The type is
    // string.
    Prefix interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is string.
    SetTag interface{}
}

func (staticRoutesState *Ocni_Vrfipv6_NetworkInstances_NetworkInstance_Protocols_Protocol_StaticRoutes_StaticRoute_StaticRoutesState) GetEntityData() *types.CommonEntityData {
    staticRoutesState.EntityData.YFilter = staticRoutesState.YFilter
    staticRoutesState.EntityData.YangName = "static-routes-state"
    staticRoutesState.EntityData.BundleName = "cisco_ios_xr"
    staticRoutesState.EntityData.ParentYangName = "static-route"
    staticRoutesState.EntityData.SegmentPath = "static-routes-state"
    staticRoutesState.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-local-routing-oper:ocni/vrfipv6/network-instances/network-instance/protocols/protocol/static-routes/static-route/" + staticRoutesState.EntityData.SegmentPath
    staticRoutesState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRoutesState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRoutesState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRoutesState.EntityData.Children = types.NewOrderedMap()
    staticRoutesState.EntityData.Leafs = types.NewOrderedMap()
    staticRoutesState.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", staticRoutesState.Prefix})
    staticRoutesState.EntityData.Leafs.Append("set-tag", types.YLeaf{"SetTag", staticRoutesState.SetTag})

    staticRoutesState.EntityData.YListKeys = []string {}

    return &(staticRoutesState.EntityData)
}

