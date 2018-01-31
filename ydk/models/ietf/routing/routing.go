// This YANG module defines essential components for the management
// of a routing subsystem.
// 
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject to
// the license terms contained in, the Simplified BSD License set
// forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// The key words 'MUST', 'MUST NOT', 'REQUIRED', 'SHALL', 'SHALL
// NOT', 'SHOULD', 'SHOULD NOT', 'RECOMMENDED', 'MAY', and
// 'OPTIONAL' in the module text are to be interpreted as described
// in RFC 2119 (http://tools.ietf.org/html/rfc2119).
// 
// This version of this YANG module is part of RFC XXXX
// (http://tools.ietf.org/html/rfcXXXX); see the RFC itself for
// full legal notices.
package routing

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package routing"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-routing routing-state}", reflect.TypeOf(RoutingState{}))
    ydk.RegisterEntity("ietf-routing:routing-state", reflect.TypeOf(RoutingState{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-routing routing}", reflect.TypeOf(Routing{}))
    ydk.RegisterEntity("ietf-routing:routing", reflect.TypeOf(Routing{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-routing fib-route}", reflect.TypeOf(FibRoute{}))
    ydk.RegisterEntity("ietf-routing:fib-route", reflect.TypeOf(FibRoute{}))
}

type VrfRoutingInstance struct {
}

func (id VrfRoutingInstance) String() string {
	return "ietf-routing:vrf-routing-instance"
}

type Direct struct {
}

func (id Direct) String() string {
	return "ietf-routing:direct"
}

type DefaultRoutingInstance struct {
}

func (id DefaultRoutingInstance) String() string {
	return "ietf-routing:default-routing-instance"
}

type RoutingProtocol struct {
}

func (id RoutingProtocol) String() string {
	return "ietf-routing:routing-protocol"
}

type Static struct {
}

func (id Static) String() string {
	return "ietf-routing:static"
}

type Ipv4 struct {
}

func (id Ipv4) String() string {
	return "ietf-routing:ipv4"
}

type Ipv6 struct {
}

func (id Ipv6) String() string {
	return "ietf-routing:ipv6"
}

type RoutingInstance struct {
}

func (id RoutingInstance) String() string {
	return "ietf-routing:routing-instance"
}

type AddressFamily struct {
}

func (id AddressFamily) String() string {
	return "ietf-routing:address-family"
}

// RoutingState
// State data of the routing subsystem.
type RoutingState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each list entry is a container for state data of a routing instance.  An
    // implementation MUST support routing instance(s) of the type
    // 'rt:default-routing-instance', and MAY support other types. An
    // implementation MAY restrict the number of routing instances of each
    // supported type.  An implementation SHOULD create at least one
    // system-controlled instance, and MAY allow the clients to create
    // user-controlled routing instances in configuration. The type is slice of
    // RoutingState_RoutingInstance.
    RoutingInstance []RoutingState_RoutingInstance
}

func (routingState *RoutingState) GetFilter() yfilter.YFilter { return routingState.YFilter }

func (routingState *RoutingState) SetFilter(yf yfilter.YFilter) { routingState.YFilter = yf }

func (routingState *RoutingState) GetGoName(yname string) string {
    if yname == "routing-instance" { return "RoutingInstance" }
    return ""
}

func (routingState *RoutingState) GetSegmentPath() string {
    return "ietf-routing:routing-state"
}

func (routingState *RoutingState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routing-instance" {
        for _, c := range routingState.RoutingInstance {
            if routingState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance{}
        routingState.RoutingInstance = append(routingState.RoutingInstance, child)
        return &routingState.RoutingInstance[len(routingState.RoutingInstance)-1]
    }
    return nil
}

func (routingState *RoutingState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routingState.RoutingInstance {
        children[routingState.RoutingInstance[i].GetSegmentPath()] = &routingState.RoutingInstance[i]
    }
    return children
}

func (routingState *RoutingState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingState *RoutingState) GetBundleName() string { return "ietf" }

func (routingState *RoutingState) GetYangName() string { return "routing-state" }

func (routingState *RoutingState) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routingState *RoutingState) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routingState *RoutingState) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routingState *RoutingState) SetParent(parent types.Entity) { routingState.parent = parent }

func (routingState *RoutingState) GetParent() types.Entity { return routingState.parent }

func (routingState *RoutingState) GetParentYangName() string { return "ietf-routing" }

// RoutingState_RoutingInstance
// Each list entry is a container for state data of a routing
// instance.
// 
// An implementation MUST support routing instance(s) of the
// type 'rt:default-routing-instance', and MAY support other
// types. An implementation MAY restrict the number of routing
// instances of each supported type.
// 
// An implementation SHOULD create at least one
// system-controlled instance, and MAY allow the clients to
// create user-controlled routing instances in
// configuration.
type RoutingState_RoutingInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the routing instance.  For
    // system-controlled instances the name is persistent, i.e., it SHOULD NOT
    // change across reboots. The type is string.
    Name interface{}

    // The routing instance type. The type is one of the following:
    // VrfRoutingInstanceDefaultRoutingInstance.
    Type interface{}

    // A 32-bit number in the form of a dotted quad that is used by some routing
    // protocols identifying a router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    RouterId interface{}

    // Network layer interfaces belonging to the routing instance.
    Interfaces RoutingState_RoutingInstance_Interfaces

    // Container for the list of routing protocol instances.
    RoutingProtocols RoutingState_RoutingInstance_RoutingProtocols

    // Container for RIBs.
    Ribs RoutingState_RoutingInstance_Ribs
}

func (routingInstance *RoutingState_RoutingInstance) GetFilter() yfilter.YFilter { return routingInstance.YFilter }

func (routingInstance *RoutingState_RoutingInstance) SetFilter(yf yfilter.YFilter) { routingInstance.YFilter = yf }

func (routingInstance *RoutingState_RoutingInstance) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "router-id" { return "RouterId" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "routing-protocols" { return "RoutingProtocols" }
    if yname == "ribs" { return "Ribs" }
    return ""
}

func (routingInstance *RoutingState_RoutingInstance) GetSegmentPath() string {
    return "routing-instance" + "[name='" + fmt.Sprintf("%v", routingInstance.Name) + "']"
}

func (routingInstance *RoutingState_RoutingInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &routingInstance.Interfaces
    }
    if childYangName == "routing-protocols" {
        return &routingInstance.RoutingProtocols
    }
    if childYangName == "ribs" {
        return &routingInstance.Ribs
    }
    return nil
}

func (routingInstance *RoutingState_RoutingInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &routingInstance.Interfaces
    children["routing-protocols"] = &routingInstance.RoutingProtocols
    children["ribs"] = &routingInstance.Ribs
    return children
}

func (routingInstance *RoutingState_RoutingInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = routingInstance.Name
    leafs["type"] = routingInstance.Type
    leafs["router-id"] = routingInstance.RouterId
    return leafs
}

func (routingInstance *RoutingState_RoutingInstance) GetBundleName() string { return "ietf" }

func (routingInstance *RoutingState_RoutingInstance) GetYangName() string { return "routing-instance" }

func (routingInstance *RoutingState_RoutingInstance) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routingInstance *RoutingState_RoutingInstance) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routingInstance *RoutingState_RoutingInstance) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routingInstance *RoutingState_RoutingInstance) SetParent(parent types.Entity) { routingInstance.parent = parent }

func (routingInstance *RoutingState_RoutingInstance) GetParent() types.Entity { return routingInstance.parent }

func (routingInstance *RoutingState_RoutingInstance) GetParentYangName() string { return "routing-state" }

// RoutingState_RoutingInstance_Interfaces
// Network layer interfaces belonging to the routing
// instance.
type RoutingState_RoutingInstance_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry is a reference to the name of a configured network layer
    // interface. The type is slice of string. Refers to
    // interfaces.InterfacesState_Interface_Name
    Interface []interface{}
}

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *RoutingState_RoutingInstance_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = interfaces.Interface
    return leafs
}

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetBundleName() string { return "ietf" }

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interfaces *RoutingState_RoutingInstance_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetParentYangName() string { return "routing-instance" }

// RoutingState_RoutingInstance_RoutingProtocols
// Container for the list of routing protocol instances.
type RoutingState_RoutingInstance_RoutingProtocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // State data of a routing protocol instance.  An implementation MUST provide
    // exactly one system-controlled instance of the type 'direct'. Other
    // instances MAY be created by configuration. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol.
    RoutingProtocol []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol
}

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetFilter() yfilter.YFilter { return routingProtocols.YFilter }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) SetFilter(yf yfilter.YFilter) { routingProtocols.YFilter = yf }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetGoName(yname string) string {
    if yname == "routing-protocol" { return "RoutingProtocol" }
    return ""
}

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetSegmentPath() string {
    return "routing-protocols"
}

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routing-protocol" {
        for _, c := range routingProtocols.RoutingProtocol {
            if routingProtocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol{}
        routingProtocols.RoutingProtocol = append(routingProtocols.RoutingProtocol, child)
        return &routingProtocols.RoutingProtocol[len(routingProtocols.RoutingProtocol)-1]
    }
    return nil
}

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routingProtocols.RoutingProtocol {
        children[routingProtocols.RoutingProtocol[i].GetSegmentPath()] = &routingProtocols.RoutingProtocol[i]
    }
    return children
}

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetBundleName() string { return "ietf" }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetYangName() string { return "routing-protocols" }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) SetParent(parent types.Entity) { routingProtocols.parent = parent }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetParent() types.Entity { return routingProtocols.parent }

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetParentYangName() string { return "routing-instance" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol
// State data of a routing protocol instance.
// 
// An implementation MUST provide exactly one
// system-controlled instance of the type 'direct'. Other
// instances MAY be created by configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type of the routing protocol. The type is one of
    // the following: Ospfv3Ospfv2OspfDirectStatic.
    Type interface{}

    // This attribute is a key. The name of the routing protocol instance.  For
    // system-controlled instances this name is persistent, i.e., it SHOULD NOT
    // change across reboots. The type is string.
    Name interface{}

    // OSPF.
    Ospf RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
}

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetFilter() yfilter.YFilter { return routingProtocol.YFilter }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) SetFilter(yf yfilter.YFilter) { routingProtocol.YFilter = yf }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "name" { return "Name" }
    if yname == "ietf-ospf:ospf" { return "Ospf" }
    return ""
}

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetSegmentPath() string {
    return "routing-protocol" + "[type='" + fmt.Sprintf("%v", routingProtocol.Type) + "']" + "[name='" + fmt.Sprintf("%v", routingProtocol.Name) + "']"
}

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ietf-ospf:ospf" {
        return &routingProtocol.Ospf
    }
    return nil
}

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ietf-ospf:ospf"] = &routingProtocol.Ospf
    return children
}

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = routingProtocol.Type
    leafs["name"] = routingProtocol.Name
    return leafs
}

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetBundleName() string { return "ietf" }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetYangName() string { return "routing-protocol" }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) SetParent(parent types.Entity) { routingProtocol.parent = parent }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetParent() types.Entity { return routingProtocol.parent }

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetParentYangName() string { return "routing-protocols" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
// OSPF
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF operation mode. The type is one of the following: ShipsInTheNight.
    OperationMode interface{}

    // An OSPF routing protocol instance. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance.
    Instance []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
}

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetGoName(yname string) string {
    if yname == "operation-mode" { return "OperationMode" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetSegmentPath() string {
    return "ietf-ospf:ospf"
}

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        for _, c := range ospf.Instance {
            if ospf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance{}
        ospf.Instance = append(ospf.Instance, child)
        return &ospf.Instance[len(ospf.Instance)-1]
    }
    return nil
}

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospf.Instance {
        children[ospf.Instance[i].GetSegmentPath()] = &ospf.Instance[i]
    }
    return children
}

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-mode"] = ospf.OperationMode
    return leafs
}

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetBundleName() string { return "ietf" }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetYangName() string { return "ospf" }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetParentYangName() string { return "routing-protocol" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
// An OSPF routing protocol instance.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address-family of the instance. The type is one of
    // the following: Ipv4Ipv4UnicastIpv6Ipv6Unicast.
    Af interface{}

    // Defined in RFC 2328. A 32-bit number that uniquely identifies the router.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    RouterId interface{}

    // List of OSPF areas. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area.
    Area []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area

    // List OSPF AS scope LSA databases. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas.
    AsScopeLsas []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas

    // OSPF topology. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
}

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    if yname == "as-scope-lsas" { return "AsScopeLsas" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetSegmentPath() string {
    return "instance" + "[af='" + fmt.Sprintf("%v", instance.Af) + "']"
}

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "area" {
        for _, c := range instance.Area {
            if instance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area{}
        instance.Area = append(instance.Area, child)
        return &instance.Area[len(instance.Area)-1]
    }
    if childYangName == "as-scope-lsas" {
        for _, c := range instance.AsScopeLsas {
            if instance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas{}
        instance.AsScopeLsas = append(instance.AsScopeLsas, child)
        return &instance.AsScopeLsas[len(instance.AsScopeLsas)-1]
    }
    if childYangName == "topology" {
        for _, c := range instance.Topology {
            if instance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology{}
        instance.Topology = append(instance.Topology, child)
        return &instance.Topology[len(instance.Topology)-1]
    }
    return nil
}

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range instance.Area {
        children[instance.Area[i].GetSegmentPath()] = &instance.Area[i]
    }
    for i := range instance.AsScopeLsas {
        children[instance.AsScopeLsas[i].GetSegmentPath()] = &instance.AsScopeLsas[i]
    }
    for i := range instance.Topology {
        children[instance.Topology[i].GetSegmentPath()] = &instance.Topology[i]
    }
    return children
}

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = instance.Af
    leafs["router-id"] = instance.RouterId
    return leafs
}

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetBundleName() string { return "ietf" }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetYangName() string { return "instance" }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetParent() types.Entity { return instance.parent }

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetParentYangName() string { return "ospf" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area
// List of OSPF areas
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AreaId interface{}

    // List of OSPF interfaces. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces.
    Interfaces []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces

    // List OSPF area scope LSA databases. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas.
    AreaScopeLsas []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetFilter() yfilter.YFilter { return area.YFilter }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) SetFilter(yf yfilter.YFilter) { area.YFilter = yf }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetGoName(yname string) string {
    if yname == "area-id" { return "AreaId" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "area-scope-lsas" { return "AreaScopeLsas" }
    return ""
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetSegmentPath() string {
    return "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        for _, c := range area.Interfaces {
            if area.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces{}
        area.Interfaces = append(area.Interfaces, child)
        return &area.Interfaces[len(area.Interfaces)-1]
    }
    if childYangName == "area-scope-lsas" {
        for _, c := range area.AreaScopeLsas {
            if area.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas{}
        area.AreaScopeLsas = append(area.AreaScopeLsas, child)
        return &area.AreaScopeLsas[len(area.AreaScopeLsas)-1]
    }
    return nil
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range area.Interfaces {
        children[area.Interfaces[i].GetSegmentPath()] = &area.Interfaces[i]
    }
    for i := range area.AreaScopeLsas {
        children[area.AreaScopeLsas[i].GetSegmentPath()] = &area.AreaScopeLsas[i]
    }
    return children
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["area-id"] = area.AreaId
    return leafs
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetBundleName() string { return "ietf" }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetYangName() string { return "area" }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) SetParent(parent types.Entity) { area.parent = parent }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetParent() types.Entity { return area.parent }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetParentYangName() string { return "instance" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces
// List of OSPF interfaces.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string.
    Interface interface{}

    // Network type. The type is NetworkType.
    NetworkType interface{}

    // Enable/Disable passive. The type is bool.
    Passive interface{}

    // Enable/Disable demand circuit. The type is bool.
    DemandCircuit interface{}

    // Set prefix as a node representative prefix. The type is bool. The default
    // value is false.
    NodeFlag interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // Interface state. The type is string.
    State interface{}

    // Hello timer. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    HelloTimer interface{}

    // Wait timer. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    WaitTimer interface{}

    // DR. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Dr interface{}

    // BDR. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bdr interface{}

    // Configure ospf multi-area.
    MultiArea RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea

    // Static configured neighbors.
    StaticNeighbors RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors

    // Fast-reroute configuration.
    FastReroute RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute

    // TTL security check.
    TtlSecurity RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity

    // Authentication configuration.
    Authentication RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication

    // List of OSPF neighbors. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor.
    Neighbor []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor

    // List OSPF link scope LSA databases. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas.
    LinkScopeLsas []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas

    // OSPF interface topology. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology
}

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "network-type" { return "NetworkType" }
    if yname == "passive" { return "Passive" }
    if yname == "demand-circuit" { return "DemandCircuit" }
    if yname == "node-flag" { return "NodeFlag" }
    if yname == "cost" { return "Cost" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "dead-interval" { return "DeadInterval" }
    if yname == "retransmit-interval" { return "RetransmitInterval" }
    if yname == "transmit-delay" { return "TransmitDelay" }
    if yname == "mtu-ignore" { return "MtuIgnore" }
    if yname == "lls" { return "Lls" }
    if yname == "prefix-suppression" { return "PrefixSuppression" }
    if yname == "bfd" { return "Bfd" }
    if yname == "enable" { return "Enable" }
    if yname == "state" { return "State" }
    if yname == "hello-timer" { return "HelloTimer" }
    if yname == "wait-timer" { return "WaitTimer" }
    if yname == "dr" { return "Dr" }
    if yname == "bdr" { return "Bdr" }
    if yname == "multi-area" { return "MultiArea" }
    if yname == "static-neighbors" { return "StaticNeighbors" }
    if yname == "fast-reroute" { return "FastReroute" }
    if yname == "ttl-security" { return "TtlSecurity" }
    if yname == "authentication" { return "Authentication" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "link-scope-lsas" { return "LinkScopeLsas" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetSegmentPath() string {
    return "interfaces" + "[interface='" + fmt.Sprintf("%v", interfaces.Interface) + "']"
}

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multi-area" {
        return &interfaces.MultiArea
    }
    if childYangName == "static-neighbors" {
        return &interfaces.StaticNeighbors
    }
    if childYangName == "fast-reroute" {
        return &interfaces.FastReroute
    }
    if childYangName == "ttl-security" {
        return &interfaces.TtlSecurity
    }
    if childYangName == "authentication" {
        return &interfaces.Authentication
    }
    if childYangName == "neighbor" {
        for _, c := range interfaces.Neighbor {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor{}
        interfaces.Neighbor = append(interfaces.Neighbor, child)
        return &interfaces.Neighbor[len(interfaces.Neighbor)-1]
    }
    if childYangName == "link-scope-lsas" {
        for _, c := range interfaces.LinkScopeLsas {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas{}
        interfaces.LinkScopeLsas = append(interfaces.LinkScopeLsas, child)
        return &interfaces.LinkScopeLsas[len(interfaces.LinkScopeLsas)-1]
    }
    if childYangName == "topology" {
        for _, c := range interfaces.Topology {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology{}
        interfaces.Topology = append(interfaces.Topology, child)
        return &interfaces.Topology[len(interfaces.Topology)-1]
    }
    return nil
}

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["multi-area"] = &interfaces.MultiArea
    children["static-neighbors"] = &interfaces.StaticNeighbors
    children["fast-reroute"] = &interfaces.FastReroute
    children["ttl-security"] = &interfaces.TtlSecurity
    children["authentication"] = &interfaces.Authentication
    for i := range interfaces.Neighbor {
        children[interfaces.Neighbor[i].GetSegmentPath()] = &interfaces.Neighbor[i]
    }
    for i := range interfaces.LinkScopeLsas {
        children[interfaces.LinkScopeLsas[i].GetSegmentPath()] = &interfaces.LinkScopeLsas[i]
    }
    for i := range interfaces.Topology {
        children[interfaces.Topology[i].GetSegmentPath()] = &interfaces.Topology[i]
    }
    return children
}

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = interfaces.Interface
    leafs["network-type"] = interfaces.NetworkType
    leafs["passive"] = interfaces.Passive
    leafs["demand-circuit"] = interfaces.DemandCircuit
    leafs["node-flag"] = interfaces.NodeFlag
    leafs["cost"] = interfaces.Cost
    leafs["hello-interval"] = interfaces.HelloInterval
    leafs["dead-interval"] = interfaces.DeadInterval
    leafs["retransmit-interval"] = interfaces.RetransmitInterval
    leafs["transmit-delay"] = interfaces.TransmitDelay
    leafs["mtu-ignore"] = interfaces.MtuIgnore
    leafs["lls"] = interfaces.Lls
    leafs["prefix-suppression"] = interfaces.PrefixSuppression
    leafs["bfd"] = interfaces.Bfd
    leafs["enable"] = interfaces.Enable
    leafs["state"] = interfaces.State
    leafs["hello-timer"] = interfaces.HelloTimer
    leafs["wait-timer"] = interfaces.WaitTimer
    leafs["dr"] = interfaces.Dr
    leafs["bdr"] = interfaces.Bdr
    return leafs
}

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetBundleName() string { return "ietf" }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetParentYangName() string { return "area" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea
// Configure ospf multi-area.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Multi-area ID. The type is one of the following types: int with range:
    // 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    MultiAreaId interface{}

    // Interface cost for multi-area. The type is interface{} with range:
    // 0..65535.
    Cost interface{}
}

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetFilter() yfilter.YFilter { return multiArea.YFilter }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) SetFilter(yf yfilter.YFilter) { multiArea.YFilter = yf }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetGoName(yname string) string {
    if yname == "multi-area-id" { return "MultiAreaId" }
    if yname == "cost" { return "Cost" }
    return ""
}

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetSegmentPath() string {
    return "multi-area"
}

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multi-area-id"] = multiArea.MultiAreaId
    leafs["cost"] = multiArea.Cost
    return leafs
}

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetBundleName() string { return "ietf" }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetYangName() string { return "multi-area" }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) SetParent(parent types.Entity) { multiArea.parent = parent }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetParent() types.Entity { return multiArea.parent }

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors
// Static configured neighbors.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify a neighbor router. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor.
    Neighbor []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor
}

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetFilter() yfilter.YFilter { return staticNeighbors.YFilter }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) SetFilter(yf yfilter.YFilter) { staticNeighbors.YFilter = yf }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetSegmentPath() string {
    return "static-neighbors"
}

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range staticNeighbors.Neighbor {
            if staticNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor{}
        staticNeighbors.Neighbor = append(staticNeighbors.Neighbor, child)
        return &staticNeighbors.Neighbor[len(staticNeighbors.Neighbor)-1]
    }
    return nil
}

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticNeighbors.Neighbor {
        children[staticNeighbors.Neighbor[i].GetSegmentPath()] = &staticNeighbors.Neighbor[i]
    }
    return children
}

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetBundleName() string { return "ietf" }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetYangName() string { return "static-neighbors" }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) SetParent(parent types.Entity) { staticNeighbors.parent = parent }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetParent() types.Entity { return staticNeighbors.parent }

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor
// Specify a neighbor router.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Neighbor cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Neighbor poll interval. The type is interface{} with range: 1..65535. Units
    // are seconds.
    PollInterval interface{}

    // Neighbor priority for DR election. The type is interface{} with range:
    // 1..255.
    Priority interface{}
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "cost" { return "Cost" }
    if yname == "poll-interval" { return "PollInterval" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[address='" + fmt.Sprintf("%v", neighbor.Address) + "']"
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = neighbor.Address
    leafs["cost"] = neighbor.Cost
    leafs["poll-interval"] = neighbor.PollInterval
    leafs["priority"] = neighbor.Priority
    return leafs
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetBundleName() string { return "ietf" }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetParentYangName() string { return "static-neighbors" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute
// Fast-reroute configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LFA configuration.
    Lfa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa
}

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetFilter() yfilter.YFilter { return fastReroute.YFilter }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) SetFilter(yf yfilter.YFilter) { fastReroute.YFilter = yf }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetGoName(yname string) string {
    if yname == "lfa" { return "Lfa" }
    return ""
}

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetSegmentPath() string {
    return "fast-reroute"
}

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lfa" {
        return &fastReroute.Lfa
    }
    return nil
}

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lfa"] = &fastReroute.Lfa
    return children
}

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetBundleName() string { return "ietf" }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetYangName() string { return "fast-reroute" }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) SetParent(parent types.Entity) { fastReroute.parent = parent }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetParent() types.Entity { return fastReroute.parent }

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa
// LFA configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prevent the interface to be used as backup. The type is bool.
    CandidateDisabled interface{}

    // Activates LFA. This model assumes activation of per-prefix LFA. The type is
    // bool.
    Enabled interface{}

    // Remote LFA configuration.
    RemoteLfa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa
}

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetFilter() yfilter.YFilter { return lfa.YFilter }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) SetFilter(yf yfilter.YFilter) { lfa.YFilter = yf }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetGoName(yname string) string {
    if yname == "candidate-disabled" { return "CandidateDisabled" }
    if yname == "enabled" { return "Enabled" }
    if yname == "remote-lfa" { return "RemoteLfa" }
    return ""
}

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetSegmentPath() string {
    return "lfa"
}

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-lfa" {
        return &lfa.RemoteLfa
    }
    return nil
}

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-lfa"] = &lfa.RemoteLfa
    return children
}

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate-disabled"] = lfa.CandidateDisabled
    leafs["enabled"] = lfa.Enabled
    return leafs
}

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetBundleName() string { return "ietf" }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetYangName() string { return "lfa" }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) SetParent(parent types.Entity) { lfa.parent = parent }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetParent() types.Entity { return lfa.parent }

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetParentYangName() string { return "fast-reroute" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa
// Remote LFA configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Activates remote LFA. The type is bool.
    Enabled interface{}
}

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetFilter() yfilter.YFilter { return remoteLfa.YFilter }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) SetFilter(yf yfilter.YFilter) { remoteLfa.YFilter = yf }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetSegmentPath() string {
    return "remote-lfa"
}

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = remoteLfa.Enabled
    return leafs
}

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetBundleName() string { return "ietf" }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetYangName() string { return "remote-lfa" }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) SetParent(parent types.Entity) { remoteLfa.parent = parent }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetParent() types.Entity { return remoteLfa.parent }

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetParentYangName() string { return "lfa" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity
// TTL security check.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetFilter() yfilter.YFilter { return ttlSecurity.YFilter }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) SetFilter(yf yfilter.YFilter) { ttlSecurity.YFilter = yf }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetSegmentPath() string {
    return "ttl-security"
}

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ttlSecurity.Enable
    leafs["hops"] = ttlSecurity.Hops
    return leafs
}

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetBundleName() string { return "ietf" }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetYangName() string { return "ttl-security" }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) SetParent(parent types.Entity) { ttlSecurity.parent = parent }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetParent() types.Entity { return ttlSecurity.parent }

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication
// Authentication configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm
}

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetGoName(yname string) string {
    if yname == "sa" { return "Sa" }
    if yname == "key-chain" { return "KeyChain" }
    if yname == "key" { return "Key" }
    if yname == "crypto-algorithm" { return "CryptoAlgorithm" }
    return ""
}

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crypto-algorithm" {
        return &authentication.CryptoAlgorithm
    }
    return nil
}

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["crypto-algorithm"] = &authentication.CryptoAlgorithm
    return children
}

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa"] = authentication.Sa
    leafs["key-chain"] = authentication.KeyChain
    leafs["key"] = authentication.Key
    return leafs
}

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetBundleName() string { return "ietf" }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetYangName() string { return "authentication" }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetFilter() yfilter.YFilter { return cryptoAlgorithm.YFilter }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) SetFilter(yf yfilter.YFilter) { cryptoAlgorithm.YFilter = yf }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetGoName(yname string) string {
    if yname == "hmac-sha1-12" { return "HmacSha112" }
    if yname == "hmac-sha1-20" { return "HmacSha120" }
    if yname == "md5" { return "Md5" }
    if yname == "sha-1" { return "Sha1" }
    if yname == "hmac-sha-1" { return "HmacSha1" }
    if yname == "hmac-sha-256" { return "HmacSha256" }
    if yname == "hmac-sha-384" { return "HmacSha384" }
    if yname == "hmac-sha-512" { return "HmacSha512" }
    return ""
}

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetSegmentPath() string {
    return "crypto-algorithm"
}

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hmac-sha1-12"] = cryptoAlgorithm.HmacSha112
    leafs["hmac-sha1-20"] = cryptoAlgorithm.HmacSha120
    leafs["md5"] = cryptoAlgorithm.Md5
    leafs["sha-1"] = cryptoAlgorithm.Sha1
    leafs["hmac-sha-1"] = cryptoAlgorithm.HmacSha1
    leafs["hmac-sha-256"] = cryptoAlgorithm.HmacSha256
    leafs["hmac-sha-384"] = cryptoAlgorithm.HmacSha384
    leafs["hmac-sha-512"] = cryptoAlgorithm.HmacSha512
    return leafs
}

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetBundleName() string { return "ietf" }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetYangName() string { return "crypto-algorithm" }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) SetParent(parent types.Entity) { cryptoAlgorithm.parent = parent }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetParent() types.Entity { return cryptoAlgorithm.parent }

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetParentYangName() string { return "authentication" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor
// List of OSPF neighbors.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Designated Router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Dr interface{}

    // Backup Designated Router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bdr interface{}

    // OSPF neighbor state. The type is NbrStateType.
    State interface{}
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "address" { return "Address" }
    if yname == "dr" { return "Dr" }
    if yname == "bdr" { return "Bdr" }
    if yname == "state" { return "State" }
    return ""
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor-id='" + fmt.Sprintf("%v", neighbor.NeighborId) + "']"
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    leafs["address"] = neighbor.Address
    leafs["dr"] = neighbor.Dr
    leafs["bdr"] = neighbor.Bdr
    leafs["state"] = neighbor.State
    return leafs
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetBundleName() string { return "ietf" }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas
// List OSPF link scope LSA databases
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..255.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa.
    LinkScopeLsa []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa
}

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetFilter() yfilter.YFilter { return linkScopeLsas.YFilter }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) SetFilter(yf yfilter.YFilter) { linkScopeLsas.YFilter = yf }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetGoName(yname string) string {
    if yname == "lsa-type" { return "LsaType" }
    if yname == "link-scope-lsa" { return "LinkScopeLsa" }
    return ""
}

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetSegmentPath() string {
    return "link-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", linkScopeLsas.LsaType) + "']"
}

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-scope-lsa" {
        for _, c := range linkScopeLsas.LinkScopeLsa {
            if linkScopeLsas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa{}
        linkScopeLsas.LinkScopeLsa = append(linkScopeLsas.LinkScopeLsa, child)
        return &linkScopeLsas.LinkScopeLsa[len(linkScopeLsas.LinkScopeLsa)-1]
    }
    return nil
}

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkScopeLsas.LinkScopeLsa {
        children[linkScopeLsas.LinkScopeLsa[i].GetSegmentPath()] = &linkScopeLsas.LinkScopeLsa[i]
    }
    return children
}

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-type"] = linkScopeLsas.LsaType
    return leafs
}

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetBundleName() string { return "ietf" }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetYangName() string { return "link-scope-lsas" }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) SetParent(parent types.Entity) { linkScopeLsas.parent = parent }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetParent() types.Entity { return linkScopeLsas.parent }

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa
// List of OSPF link scope LSAs
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or int with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RawData interface{}

    // OSPFv2 LSA.
    Ospfv2 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2

    // OSPFv3 LSA.
    Ospfv3 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3
}

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetFilter() yfilter.YFilter { return linkScopeLsa.YFilter }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) SetFilter(yf yfilter.YFilter) { linkScopeLsa.YFilter = yf }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetGoName(yname string) string {
    if yname == "lsa-id" { return "LsaId" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "decoded-completed" { return "DecodedCompleted" }
    if yname == "raw-data" { return "RawData" }
    if yname == "ospfv2" { return "Ospfv2" }
    if yname == "ospfv3" { return "Ospfv3" }
    return ""
}

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetSegmentPath() string {
    return "link-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", linkScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", linkScopeLsa.AdvRouter) + "']"
}

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2" {
        return &linkScopeLsa.Ospfv2
    }
    if childYangName == "ospfv3" {
        return &linkScopeLsa.Ospfv3
    }
    return nil
}

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv2"] = &linkScopeLsa.Ospfv2
    children["ospfv3"] = &linkScopeLsa.Ospfv3
    return children
}

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-id"] = linkScopeLsa.LsaId
    leafs["adv-router"] = linkScopeLsa.AdvRouter
    leafs["decoded-completed"] = linkScopeLsa.DecodedCompleted
    leafs["raw-data"] = linkScopeLsa.RawData
    return leafs
}

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetBundleName() string { return "ietf" }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetYangName() string { return "link-scope-lsa" }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) SetParent(parent types.Entity) { linkScopeLsa.parent = parent }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetParent() types.Entity { return linkScopeLsa.parent }

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetParentYangName() string { return "link-scope-lsas" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2
// OSPFv2 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header

    // Decoded OSPFv2 LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetFilter() yfilter.YFilter { return ospfv2.YFilter }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) SetFilter(yf yfilter.YFilter) { ospfv2.YFilter = yf }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "body" { return "Body" }
    return ""
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetSegmentPath() string {
    return "ospfv2"
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &ospfv2.Header
    }
    if childYangName == "body" {
        return &ospfv2.Body
    }
    return nil
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &ospfv2.Header
    children["body"] = &ospfv2.Body
    return children
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetBundleName() string { return "ietf" }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetYangName() string { return "ospfv2" }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) SetParent(parent types.Entity) { ospfv2.parent = parent }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetParent() types.Entity { return ospfv2.parent }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetParentYangName() string { return "link-scope-lsa" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header
// Decoded OSPFv2 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Options interface{}

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    OpaqueType interface{}

    // Opaque id. The type is interface{} with range: 0..16777215. This attribute
    // is mandatory.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type interface{}

    // LSA advertising router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "lsa-id" { return "LsaId" }
    if yname == "opaque-type" { return "OpaqueType" }
    if yname == "opaque-id" { return "OpaqueId" }
    if yname == "age" { return "Age" }
    if yname == "type" { return "Type" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "seq-num" { return "SeqNum" }
    if yname == "checksum" { return "Checksum" }
    if yname == "length" { return "Length" }
    return ""
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetSegmentPath() string {
    return "header"
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = header.Options
    leafs["lsa-id"] = header.LsaId
    leafs["opaque-type"] = header.OpaqueType
    leafs["opaque-id"] = header.OpaqueId
    leafs["age"] = header.Age
    leafs["type"] = header.Type
    leafs["adv-router"] = header.AdvRouter
    leafs["seq-num"] = header.SeqNum
    leafs["checksum"] = header.Checksum
    leafs["length"] = header.Length
    return leafs
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetBundleName() string { return "ietf" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetYangName() string { return "header" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetParent() types.Entity { return header.parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetParentYangName() string { return "ospfv2" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body
// Decoded OSPFv2 LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network

    // Summary LSA.
    Summary RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary

    // External LSA.
    External RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External

    // Opaque LSA.
    Opaque RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetFilter() yfilter.YFilter { return body.YFilter }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) SetFilter(yf yfilter.YFilter) { body.YFilter = yf }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetGoName(yname string) string {
    if yname == "router" { return "Router" }
    if yname == "network" { return "Network" }
    if yname == "summary" { return "Summary" }
    if yname == "external" { return "External" }
    if yname == "opaque" { return "Opaque" }
    return ""
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetSegmentPath() string {
    return "body"
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router" {
        return &body.Router
    }
    if childYangName == "network" {
        return &body.Network
    }
    if childYangName == "summary" {
        return &body.Summary
    }
    if childYangName == "external" {
        return &body.External
    }
    if childYangName == "opaque" {
        return &body.Opaque
    }
    return nil
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router"] = &body.Router
    children["network"] = &body.Network
    children["summary"] = &body.Summary
    children["external"] = &body.External
    children["opaque"] = &body.Opaque
    return children
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetBundleName() string { return "ietf" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetYangName() string { return "body" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) SetParent(parent types.Entity) { body.parent = parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetParent() types.Entity { return body.parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetParentYangName() string { return "ospfv2" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetFilter() yfilter.YFilter { return router.YFilter }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) SetFilter(yf yfilter.YFilter) { router.YFilter = yf }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetGoName(yname string) string {
    if yname == "flags" { return "Flags" }
    if yname == "num-of-links" { return "NumOfLinks" }
    if yname == "link" { return "Link" }
    return ""
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetSegmentPath() string {
    return "router"
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link" {
        for _, c := range router.Link {
            if router.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link{}
        router.Link = append(router.Link, child)
        return &router.Link[len(router.Link)-1]
    }
    return nil
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range router.Link {
        children[router.Link[i].GetSegmentPath()] = &router.Link[i]
    }
    return children
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flags"] = router.Flags
    leafs["num-of-links"] = router.NumOfLinks
    return leafs
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetBundleName() string { return "ietf" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetYangName() string { return "router" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) SetParent(parent types.Entity) { router.parent = parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetParent() types.Entity { return router.parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    LinkId interface{}

    // This attribute is a key. Link data. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or int with range: 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetGoName(yname string) string {
    if yname == "link-id" { return "LinkId" }
    if yname == "link-data" { return "LinkData" }
    if yname == "type" { return "Type" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetSegmentPath() string {
    return "link" + "[link-id='" + fmt.Sprintf("%v", link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", link.LinkData) + "']"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range link.Topology {
            if link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology{}
        link.Topology = append(link.Topology, child)
        return &link.Topology[len(link.Topology)-1]
    }
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range link.Topology {
        children[link.Topology[i].GetSegmentPath()] = &link.Topology[i]
    }
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-id"] = link.LinkId
    leafs["link-data"] = link.LinkData
    leafs["type"] = link.Type
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetParentYangName() string { return "router" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["metric"] = topology.Metric
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetParentYangName() string { return "link" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "attached-router" { return "AttachedRouter" }
    return ""
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetSegmentPath() string {
    return "network"
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = network.NetworkMask
    leafs["attached-router"] = network.AttachedRouter
    return leafs
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetBundleName() string { return "ietf" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetYangName() string { return "network" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetParent() types.Entity { return network.parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary
// Summary LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range summary.Topology {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology{}
        summary.Topology = append(summary.Topology, child)
        return &summary.Topology[len(summary.Topology)-1]
    }
    return nil
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.Topology {
        children[summary.Topology[i].GetSegmentPath()] = &summary.Topology[i]
    }
    return children
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = summary.NetworkMask
    return leafs
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetBundleName() string { return "ietf" }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetYangName() string { return "summary" }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetParent() types.Entity { return summary.parent }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["metric"] = topology.Metric
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetParentYangName() string { return "summary" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External
// External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetFilter() yfilter.YFilter { return external.YFilter }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) SetFilter(yf yfilter.YFilter) { external.YFilter = yf }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetSegmentPath() string {
    return "external"
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range external.Topology {
            if external.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology{}
        external.Topology = append(external.Topology, child)
        return &external.Topology[len(external.Topology)-1]
    }
    return nil
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range external.Topology {
        children[external.Topology[i].GetSegmentPath()] = &external.Topology[i]
    }
    return children
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = external.NetworkMask
    return leafs
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetBundleName() string { return "ietf" }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetYangName() string { return "external" }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) SetParent(parent types.Entity) { external.parent = parent }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetParent() types.Entity { return external.parent }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Forwarding address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "flags" { return "Flags" }
    if yname == "metric" { return "Metric" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["flags"] = topology.Flags
    leafs["metric"] = topology.Metric
    leafs["forwarding-address"] = topology.ForwardingAddress
    leafs["external-route-tag"] = topology.ExternalRouteTag
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetParentYangName() string { return "external" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque
// Opaque LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unknown TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv.
    UnknownTlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv

    // Router address TLV.
    RouterAddressTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv

    // Link TLV.
    LinkTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetFilter() yfilter.YFilter { return opaque.YFilter }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) SetFilter(yf yfilter.YFilter) { opaque.YFilter = yf }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetGoName(yname string) string {
    if yname == "unknown-tlv" { return "UnknownTlv" }
    if yname == "router-address-tlv" { return "RouterAddressTlv" }
    if yname == "link-tlv" { return "LinkTlv" }
    return ""
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetSegmentPath() string {
    return "opaque"
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-tlv" {
        for _, c := range opaque.UnknownTlv {
            if opaque.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv{}
        opaque.UnknownTlv = append(opaque.UnknownTlv, child)
        return &opaque.UnknownTlv[len(opaque.UnknownTlv)-1]
    }
    if childYangName == "router-address-tlv" {
        return &opaque.RouterAddressTlv
    }
    if childYangName == "link-tlv" {
        return &opaque.LinkTlv
    }
    return nil
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opaque.UnknownTlv {
        children[opaque.UnknownTlv[i].GetSegmentPath()] = &opaque.UnknownTlv[i]
    }
    children["router-address-tlv"] = &opaque.RouterAddressTlv
    children["link-tlv"] = &opaque.LinkTlv
    return children
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetBundleName() string { return "ietf" }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetYangName() string { return "opaque" }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) SetParent(parent types.Entity) { opaque.parent = parent }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetParent() types.Entity { return opaque.parent }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv
// Unknown TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetFilter() yfilter.YFilter { return unknownTlv.YFilter }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) SetFilter(yf yfilter.YFilter) { unknownTlv.YFilter = yf }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "length" { return "Length" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetSegmentPath() string {
    return "unknown-tlv" + "[type='" + fmt.Sprintf("%v", unknownTlv.Type) + "']"
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = unknownTlv.Type
    leafs["length"] = unknownTlv.Length
    leafs["value"] = unknownTlv.Value
    return leafs
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetBundleName() string { return "ietf" }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetYangName() string { return "unknown-tlv" }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) SetParent(parent types.Entity) { unknownTlv.parent = parent }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetParent() types.Entity { return unknownTlv.parent }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv
// Router address TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterAddress interface{}
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetFilter() yfilter.YFilter { return routerAddressTlv.YFilter }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) SetFilter(yf yfilter.YFilter) { routerAddressTlv.YFilter = yf }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetGoName(yname string) string {
    if yname == "router-address" { return "RouterAddress" }
    return ""
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetSegmentPath() string {
    return "router-address-tlv"
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-address"] = routerAddressTlv.RouterAddress
    return leafs
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetBundleName() string { return "ietf" }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetYangName() string { return "router-address-tlv" }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) SetParent(parent types.Entity) { routerAddressTlv.parent = parent }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetParent() types.Entity { return routerAddressTlv.parent }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv
// Link TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    LinkType interface{}

    // Link ID. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])
    // This attribute is mandatory..
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unreserved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}

    // Unknown sub-TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv.
    UnknownSubtlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetFilter() yfilter.YFilter { return linkTlv.YFilter }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) SetFilter(yf yfilter.YFilter) { linkTlv.YFilter = yf }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-id" { return "LinkId" }
    if yname == "local-if-ipv4-addr" { return "LocalIfIpv4Addr" }
    if yname == "local-remote-ipv4-addr" { return "LocalRemoteIpv4Addr" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "max-bandwidth" { return "MaxBandwidth" }
    if yname == "max-reservable-bandwidth" { return "MaxReservableBandwidth" }
    if yname == "unreserved-bandwidth" { return "UnreservedBandwidth" }
    if yname == "admin-group" { return "AdminGroup" }
    if yname == "unknown-subtlv" { return "UnknownSubtlv" }
    return ""
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetSegmentPath() string {
    return "link-tlv"
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-subtlv" {
        for _, c := range linkTlv.UnknownSubtlv {
            if linkTlv.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv{}
        linkTlv.UnknownSubtlv = append(linkTlv.UnknownSubtlv, child)
        return &linkTlv.UnknownSubtlv[len(linkTlv.UnknownSubtlv)-1]
    }
    return nil
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkTlv.UnknownSubtlv {
        children[linkTlv.UnknownSubtlv[i].GetSegmentPath()] = &linkTlv.UnknownSubtlv[i]
    }
    return children
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = linkTlv.LinkType
    leafs["link-id"] = linkTlv.LinkId
    leafs["local-if-ipv4-addr"] = linkTlv.LocalIfIpv4Addr
    leafs["local-remote-ipv4-addr"] = linkTlv.LocalRemoteIpv4Addr
    leafs["te-metric"] = linkTlv.TeMetric
    leafs["max-bandwidth"] = linkTlv.MaxBandwidth
    leafs["max-reservable-bandwidth"] = linkTlv.MaxReservableBandwidth
    leafs["unreserved-bandwidth"] = linkTlv.UnreservedBandwidth
    leafs["admin-group"] = linkTlv.AdminGroup
    return leafs
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetBundleName() string { return "ietf" }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetYangName() string { return "link-tlv" }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) SetParent(parent types.Entity) { linkTlv.parent = parent }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetParent() types.Entity { return linkTlv.parent }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
// Unknown sub-TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetFilter() yfilter.YFilter { return unknownSubtlv.YFilter }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) SetFilter(yf yfilter.YFilter) { unknownSubtlv.YFilter = yf }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "length" { return "Length" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetSegmentPath() string {
    return "unknown-subtlv" + "[type='" + fmt.Sprintf("%v", unknownSubtlv.Type) + "']"
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = unknownSubtlv.Type
    leafs["length"] = unknownSubtlv.Length
    leafs["value"] = unknownSubtlv.Value
    return leafs
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetBundleName() string { return "ietf" }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetYangName() string { return "unknown-subtlv" }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) SetParent(parent types.Entity) { unknownSubtlv.parent = parent }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetParent() types.Entity { return unknownSubtlv.parent }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetParentYangName() string { return "link-tlv" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3
// OSPFv3 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header

    // Decoded OSPF LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetFilter() yfilter.YFilter { return ospfv3.YFilter }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) SetFilter(yf yfilter.YFilter) { ospfv3.YFilter = yf }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "body" { return "Body" }
    return ""
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetSegmentPath() string {
    return "ospfv3"
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &ospfv3.Header
    }
    if childYangName == "body" {
        return &ospfv3.Body
    }
    return nil
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &ospfv3.Header
    children["body"] = &ospfv3.Body
    return children
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetBundleName() string { return "ietf" }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetYangName() string { return "ospfv3" }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) SetParent(parent types.Entity) { ospfv3.parent = parent }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetParent() types.Entity { return ospfv3.parent }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetParentYangName() string { return "link-scope-lsa" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header
// Decoded OSPFv3 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA ID. The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory.
    LsaId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type interface{}

    // LSA advertising router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetGoName(yname string) string {
    if yname == "lsa-id" { return "LsaId" }
    if yname == "age" { return "Age" }
    if yname == "type" { return "Type" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "seq-num" { return "SeqNum" }
    if yname == "checksum" { return "Checksum" }
    if yname == "length" { return "Length" }
    if yname == "options" { return "Options" }
    return ""
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetSegmentPath() string {
    return "header"
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-id"] = header.LsaId
    leafs["age"] = header.Age
    leafs["type"] = header.Type
    leafs["adv-router"] = header.AdvRouter
    leafs["seq-num"] = header.SeqNum
    leafs["checksum"] = header.Checksum
    leafs["length"] = header.Length
    leafs["options"] = header.Options
    return leafs
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetBundleName() string { return "ietf" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetYangName() string { return "header" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetParent() types.Entity { return header.parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetParentYangName() string { return "ospfv3" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body
// Decoded OSPF LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network

    // Inter-Area-Prefix LSA.
    InterAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix

    // Inter-Area-Router LSA.
    InterAreaRouter RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter

    // AS-External LSA.
    AsExternal RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal

    // NSSA LSA.
    Nssa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa

    // Link LSA.
    Link RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link

    // Intra-Area-Prefix LSA.
    IntraAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetFilter() yfilter.YFilter { return body.YFilter }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) SetFilter(yf yfilter.YFilter) { body.YFilter = yf }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetGoName(yname string) string {
    if yname == "router" { return "Router" }
    if yname == "network" { return "Network" }
    if yname == "inter-area-prefix" { return "InterAreaPrefix" }
    if yname == "inter-area-router" { return "InterAreaRouter" }
    if yname == "as-external" { return "AsExternal" }
    if yname == "nssa" { return "Nssa" }
    if yname == "link" { return "Link" }
    if yname == "intra-area-prefix" { return "IntraAreaPrefix" }
    return ""
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetSegmentPath() string {
    return "body"
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router" {
        return &body.Router
    }
    if childYangName == "network" {
        return &body.Network
    }
    if childYangName == "inter-area-prefix" {
        return &body.InterAreaPrefix
    }
    if childYangName == "inter-area-router" {
        return &body.InterAreaRouter
    }
    if childYangName == "as-external" {
        return &body.AsExternal
    }
    if childYangName == "nssa" {
        return &body.Nssa
    }
    if childYangName == "link" {
        return &body.Link
    }
    if childYangName == "intra-area-prefix" {
        return &body.IntraAreaPrefix
    }
    return nil
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router"] = &body.Router
    children["network"] = &body.Network
    children["inter-area-prefix"] = &body.InterAreaPrefix
    children["inter-area-router"] = &body.InterAreaRouter
    children["as-external"] = &body.AsExternal
    children["nssa"] = &body.Nssa
    children["link"] = &body.Link
    children["intra-area-prefix"] = &body.IntraAreaPrefix
    return children
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetBundleName() string { return "ietf" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetYangName() string { return "body" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) SetParent(parent types.Entity) { body.parent = parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetParent() types.Entity { return body.parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetParentYangName() string { return "ospfv3" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Flags interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetFilter() yfilter.YFilter { return router.YFilter }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) SetFilter(yf yfilter.YFilter) { router.YFilter = yf }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetGoName(yname string) string {
    if yname == "flags" { return "Flags" }
    if yname == "options" { return "Options" }
    if yname == "link" { return "Link" }
    return ""
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetSegmentPath() string {
    return "router"
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link" {
        for _, c := range router.Link {
            if router.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link{}
        router.Link = append(router.Link, child)
        return &router.Link[len(router.Link)-1]
    }
    return nil
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range router.Link {
        children[router.Link[i].GetSegmentPath()] = &router.Link[i]
    }
    return children
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flags"] = router.Flags
    leafs["options"] = router.Options
    return leafs
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetBundleName() string { return "ietf" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetYangName() string { return "router" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) SetParent(parent types.Entity) { router.parent = parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetParent() types.Entity { return router.parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor Interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor Router ID. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetGoName(yname string) string {
    if yname == "interface-id" { return "InterfaceId" }
    if yname == "neighbor-interface-id" { return "NeighborInterfaceId" }
    if yname == "neighbor-router-id" { return "NeighborRouterId" }
    if yname == "type" { return "Type" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetSegmentPath() string {
    return "link" + "[interface-id='" + fmt.Sprintf("%v", link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", link.NeighborRouterId) + "']"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-id"] = link.InterfaceId
    leafs["neighbor-interface-id"] = link.NeighborInterfaceId
    leafs["neighbor-router-id"] = link.NeighborRouterId
    leafs["type"] = link.Type
    leafs["metric"] = link.Metric
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetParentYangName() string { return "router" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "attached-router" { return "AttachedRouter" }
    return ""
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetSegmentPath() string {
    return "network"
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = network.Options
    leafs["attached-router"] = network.AttachedRouter
    return leafs
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetBundleName() string { return "ietf" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetYangName() string { return "network" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetParent() types.Entity { return network.parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix
// Inter-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetFilter() yfilter.YFilter { return interAreaPrefix.YFilter }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) SetFilter(yf yfilter.YFilter) { interAreaPrefix.YFilter = yf }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    return ""
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetSegmentPath() string {
    return "inter-area-prefix"
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = interAreaPrefix.Metric
    leafs["prefix"] = interAreaPrefix.Prefix
    leafs["prefix-options"] = interAreaPrefix.PrefixOptions
    return leafs
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetBundleName() string { return "ietf" }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetYangName() string { return "inter-area-prefix" }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) SetParent(parent types.Entity) { interAreaPrefix.parent = parent }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetParent() types.Entity { return interAreaPrefix.parent }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter
// Inter-Area-Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // The Router ID of the router being described by the LSA. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    DestinationRouterId interface{}
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetFilter() yfilter.YFilter { return interAreaRouter.YFilter }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) SetFilter(yf yfilter.YFilter) { interAreaRouter.YFilter = yf }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "metric" { return "Metric" }
    if yname == "destination-router-id" { return "DestinationRouterId" }
    return ""
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetSegmentPath() string {
    return "inter-area-router"
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = interAreaRouter.Options
    leafs["metric"] = interAreaRouter.Metric
    leafs["destination-router-id"] = interAreaRouter.DestinationRouterId
    return leafs
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetBundleName() string { return "ietf" }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetYangName() string { return "inter-area-router" }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) SetParent(parent types.Entity) { interAreaRouter.parent = parent }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetParent() types.Entity { return interAreaRouter.parent }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal
// AS-External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetFilter() yfilter.YFilter { return asExternal.YFilter }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) SetFilter(yf yfilter.YFilter) { asExternal.YFilter = yf }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "flags" { return "Flags" }
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    return ""
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetSegmentPath() string {
    return "as-external"
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = asExternal.Metric
    leafs["flags"] = asExternal.Flags
    leafs["referenced-ls-type"] = asExternal.ReferencedLsType
    leafs["prefix"] = asExternal.Prefix
    leafs["prefix-options"] = asExternal.PrefixOptions
    leafs["forwarding-address"] = asExternal.ForwardingAddress
    leafs["external-route-tag"] = asExternal.ExternalRouteTag
    leafs["referenced-link-state-id"] = asExternal.ReferencedLinkStateId
    return leafs
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetBundleName() string { return "ietf" }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetYangName() string { return "as-external" }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) SetParent(parent types.Entity) { asExternal.parent = parent }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetParent() types.Entity { return asExternal.parent }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa
// NSSA LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetFilter() yfilter.YFilter { return nssa.YFilter }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) SetFilter(yf yfilter.YFilter) { nssa.YFilter = yf }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "flags" { return "Flags" }
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    return ""
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetSegmentPath() string {
    return "nssa"
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = nssa.Metric
    leafs["flags"] = nssa.Flags
    leafs["referenced-ls-type"] = nssa.ReferencedLsType
    leafs["prefix"] = nssa.Prefix
    leafs["prefix-options"] = nssa.PrefixOptions
    leafs["forwarding-address"] = nssa.ForwardingAddress
    leafs["external-route-tag"] = nssa.ExternalRouteTag
    leafs["referenced-link-state-id"] = nssa.ReferencedLinkStateId
    return leafs
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetBundleName() string { return "ietf" }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetYangName() string { return "nssa" }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) SetParent(parent types.Entity) { nssa.parent = parent }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetParent() types.Entity { return nssa.parent }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link
// Link LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router Priority of the interface. The type is interface{} with range:
    // 0..255.
    RtrPriority interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetGoName(yname string) string {
    if yname == "rtr-priority" { return "RtrPriority" }
    if yname == "options" { return "Options" }
    if yname == "link-local-interface-address" { return "LinkLocalInterfaceAddress" }
    if yname == "num-of-prefixes" { return "NumOfPrefixes" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetSegmentPath() string {
    return "link"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list" {
        for _, c := range link.PrefixList {
            if link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList{}
        link.PrefixList = append(link.PrefixList, child)
        return &link.PrefixList[len(link.PrefixList)-1]
    }
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range link.PrefixList {
        children[link.PrefixList[i].GetSegmentPath()] = &link.PrefixList[i]
    }
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rtr-priority"] = link.RtrPriority
    leafs["options"] = link.Options
    leafs["link-local-interface-address"] = link.LinkLocalInterfaceAddress
    leafs["num-of-prefixes"] = link.NumOfPrefixes
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetFilter() yfilter.YFilter { return prefixList.YFilter }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) SetFilter(yf yfilter.YFilter) { prefixList.YFilter = yf }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    return ""
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetSegmentPath() string {
    return "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = prefixList.Prefix
    leafs["prefix-options"] = prefixList.PrefixOptions
    return leafs
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetBundleName() string { return "ietf" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetYangName() string { return "prefix-list" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) SetParent(parent types.Entity) { prefixList.parent = parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetParent() types.Entity { return prefixList.parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetParentYangName() string { return "link" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix
// Intra-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetFilter() yfilter.YFilter { return intraAreaPrefix.YFilter }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) SetFilter(yf yfilter.YFilter) { intraAreaPrefix.YFilter = yf }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetGoName(yname string) string {
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    if yname == "referenced-adv-router" { return "ReferencedAdvRouter" }
    if yname == "num-of-prefixes" { return "NumOfPrefixes" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetSegmentPath() string {
    return "intra-area-prefix"
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list" {
        for _, c := range intraAreaPrefix.PrefixList {
            if intraAreaPrefix.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList{}
        intraAreaPrefix.PrefixList = append(intraAreaPrefix.PrefixList, child)
        return &intraAreaPrefix.PrefixList[len(intraAreaPrefix.PrefixList)-1]
    }
    return nil
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intraAreaPrefix.PrefixList {
        children[intraAreaPrefix.PrefixList[i].GetSegmentPath()] = &intraAreaPrefix.PrefixList[i]
    }
    return children
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["referenced-ls-type"] = intraAreaPrefix.ReferencedLsType
    leafs["referenced-link-state-id"] = intraAreaPrefix.ReferencedLinkStateId
    leafs["referenced-adv-router"] = intraAreaPrefix.ReferencedAdvRouter
    leafs["num-of-prefixes"] = intraAreaPrefix.NumOfPrefixes
    return leafs
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetBundleName() string { return "ietf" }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetYangName() string { return "intra-area-prefix" }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) SetParent(parent types.Entity) { intraAreaPrefix.parent = parent }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetParent() types.Entity { return intraAreaPrefix.parent }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetFilter() yfilter.YFilter { return prefixList.YFilter }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) SetFilter(yf yfilter.YFilter) { prefixList.YFilter = yf }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetSegmentPath() string {
    return "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = prefixList.Prefix
    leafs["prefix-options"] = prefixList.PrefixOptions
    leafs["metric"] = prefixList.Metric
    return leafs
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetBundleName() string { return "ietf" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetYangName() string { return "prefix-list" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) SetParent(parent types.Entity) { prefixList.parent = parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetParent() types.Entity { return prefixList.parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetParentYangName() string { return "intra-area-prefix" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology
// OSPF interface topology.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string. Refers to routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetSegmentPath() string {
    return "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = topology.Name
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetParentYangName() string { return "interfaces" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType represents Network type.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType string

const (
    // Specify OSPF broadcast multi-access network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_broadcast RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "broadcast"

    // Specify OSPF Non-Broadcast Multi-Access
    // (NBMA) network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_non_broadcast RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "non-broadcast"

    // Specify OSPF point-to-multipoint network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_point_to_multipoint RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "point-to-multipoint"

    // Specify OSPF point-to-point network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_point_to_point RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "point-to-point"
)

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas
// List OSPF area scope LSA databases
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF area scope LSA type. The type is interface{}
    // with range: 0..255.
    LsaType interface{}

    // List of OSPF area scope LSAs. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa.
    AreaScopeLsa []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa
}

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetFilter() yfilter.YFilter { return areaScopeLsas.YFilter }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) SetFilter(yf yfilter.YFilter) { areaScopeLsas.YFilter = yf }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetGoName(yname string) string {
    if yname == "lsa-type" { return "LsaType" }
    if yname == "area-scope-lsa" { return "AreaScopeLsa" }
    return ""
}

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetSegmentPath() string {
    return "area-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", areaScopeLsas.LsaType) + "']"
}

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "area-scope-lsa" {
        for _, c := range areaScopeLsas.AreaScopeLsa {
            if areaScopeLsas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa{}
        areaScopeLsas.AreaScopeLsa = append(areaScopeLsas.AreaScopeLsa, child)
        return &areaScopeLsas.AreaScopeLsa[len(areaScopeLsas.AreaScopeLsa)-1]
    }
    return nil
}

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range areaScopeLsas.AreaScopeLsa {
        children[areaScopeLsas.AreaScopeLsa[i].GetSegmentPath()] = &areaScopeLsas.AreaScopeLsa[i]
    }
    return children
}

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-type"] = areaScopeLsas.LsaType
    return leafs
}

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetBundleName() string { return "ietf" }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetYangName() string { return "area-scope-lsas" }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) SetParent(parent types.Entity) { areaScopeLsas.parent = parent }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetParent() types.Entity { return areaScopeLsas.parent }

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetParentYangName() string { return "area" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa
// List of OSPF area scope LSAs
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or int with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RawData interface{}

    // OSPFv2 LSA.
    Ospfv2 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2

    // OSPFv3 LSA.
    Ospfv3 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3
}

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetFilter() yfilter.YFilter { return areaScopeLsa.YFilter }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) SetFilter(yf yfilter.YFilter) { areaScopeLsa.YFilter = yf }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetGoName(yname string) string {
    if yname == "lsa-id" { return "LsaId" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "decoded-completed" { return "DecodedCompleted" }
    if yname == "raw-data" { return "RawData" }
    if yname == "ospfv2" { return "Ospfv2" }
    if yname == "ospfv3" { return "Ospfv3" }
    return ""
}

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetSegmentPath() string {
    return "area-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", areaScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", areaScopeLsa.AdvRouter) + "']"
}

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2" {
        return &areaScopeLsa.Ospfv2
    }
    if childYangName == "ospfv3" {
        return &areaScopeLsa.Ospfv3
    }
    return nil
}

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv2"] = &areaScopeLsa.Ospfv2
    children["ospfv3"] = &areaScopeLsa.Ospfv3
    return children
}

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-id"] = areaScopeLsa.LsaId
    leafs["adv-router"] = areaScopeLsa.AdvRouter
    leafs["decoded-completed"] = areaScopeLsa.DecodedCompleted
    leafs["raw-data"] = areaScopeLsa.RawData
    return leafs
}

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetBundleName() string { return "ietf" }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetYangName() string { return "area-scope-lsa" }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) SetParent(parent types.Entity) { areaScopeLsa.parent = parent }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetParent() types.Entity { return areaScopeLsa.parent }

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetParentYangName() string { return "area-scope-lsas" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2
// OSPFv2 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header

    // Decoded OSPFv2 LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetFilter() yfilter.YFilter { return ospfv2.YFilter }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) SetFilter(yf yfilter.YFilter) { ospfv2.YFilter = yf }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "body" { return "Body" }
    return ""
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetSegmentPath() string {
    return "ospfv2"
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &ospfv2.Header
    }
    if childYangName == "body" {
        return &ospfv2.Body
    }
    return nil
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &ospfv2.Header
    children["body"] = &ospfv2.Body
    return children
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetBundleName() string { return "ietf" }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetYangName() string { return "ospfv2" }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) SetParent(parent types.Entity) { ospfv2.parent = parent }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetParent() types.Entity { return ospfv2.parent }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetParentYangName() string { return "area-scope-lsa" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header
// Decoded OSPFv2 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Options interface{}

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    OpaqueType interface{}

    // Opaque id. The type is interface{} with range: 0..16777215. This attribute
    // is mandatory.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type interface{}

    // LSA advertising router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "lsa-id" { return "LsaId" }
    if yname == "opaque-type" { return "OpaqueType" }
    if yname == "opaque-id" { return "OpaqueId" }
    if yname == "age" { return "Age" }
    if yname == "type" { return "Type" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "seq-num" { return "SeqNum" }
    if yname == "checksum" { return "Checksum" }
    if yname == "length" { return "Length" }
    return ""
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetSegmentPath() string {
    return "header"
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = header.Options
    leafs["lsa-id"] = header.LsaId
    leafs["opaque-type"] = header.OpaqueType
    leafs["opaque-id"] = header.OpaqueId
    leafs["age"] = header.Age
    leafs["type"] = header.Type
    leafs["adv-router"] = header.AdvRouter
    leafs["seq-num"] = header.SeqNum
    leafs["checksum"] = header.Checksum
    leafs["length"] = header.Length
    return leafs
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetBundleName() string { return "ietf" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetYangName() string { return "header" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetParent() types.Entity { return header.parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetParentYangName() string { return "ospfv2" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body
// Decoded OSPFv2 LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network

    // Summary LSA.
    Summary RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary

    // External LSA.
    External RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External

    // Opaque LSA.
    Opaque RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetFilter() yfilter.YFilter { return body.YFilter }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) SetFilter(yf yfilter.YFilter) { body.YFilter = yf }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetGoName(yname string) string {
    if yname == "router" { return "Router" }
    if yname == "network" { return "Network" }
    if yname == "summary" { return "Summary" }
    if yname == "external" { return "External" }
    if yname == "opaque" { return "Opaque" }
    return ""
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetSegmentPath() string {
    return "body"
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router" {
        return &body.Router
    }
    if childYangName == "network" {
        return &body.Network
    }
    if childYangName == "summary" {
        return &body.Summary
    }
    if childYangName == "external" {
        return &body.External
    }
    if childYangName == "opaque" {
        return &body.Opaque
    }
    return nil
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router"] = &body.Router
    children["network"] = &body.Network
    children["summary"] = &body.Summary
    children["external"] = &body.External
    children["opaque"] = &body.Opaque
    return children
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetBundleName() string { return "ietf" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetYangName() string { return "body" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) SetParent(parent types.Entity) { body.parent = parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetParent() types.Entity { return body.parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetParentYangName() string { return "ospfv2" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetFilter() yfilter.YFilter { return router.YFilter }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) SetFilter(yf yfilter.YFilter) { router.YFilter = yf }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetGoName(yname string) string {
    if yname == "flags" { return "Flags" }
    if yname == "num-of-links" { return "NumOfLinks" }
    if yname == "link" { return "Link" }
    return ""
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetSegmentPath() string {
    return "router"
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link" {
        for _, c := range router.Link {
            if router.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link{}
        router.Link = append(router.Link, child)
        return &router.Link[len(router.Link)-1]
    }
    return nil
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range router.Link {
        children[router.Link[i].GetSegmentPath()] = &router.Link[i]
    }
    return children
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flags"] = router.Flags
    leafs["num-of-links"] = router.NumOfLinks
    return leafs
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetBundleName() string { return "ietf" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetYangName() string { return "router" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) SetParent(parent types.Entity) { router.parent = parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetParent() types.Entity { return router.parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    LinkId interface{}

    // This attribute is a key. Link data. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or int with range: 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetGoName(yname string) string {
    if yname == "link-id" { return "LinkId" }
    if yname == "link-data" { return "LinkData" }
    if yname == "type" { return "Type" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetSegmentPath() string {
    return "link" + "[link-id='" + fmt.Sprintf("%v", link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", link.LinkData) + "']"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range link.Topology {
            if link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology{}
        link.Topology = append(link.Topology, child)
        return &link.Topology[len(link.Topology)-1]
    }
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range link.Topology {
        children[link.Topology[i].GetSegmentPath()] = &link.Topology[i]
    }
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-id"] = link.LinkId
    leafs["link-data"] = link.LinkData
    leafs["type"] = link.Type
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetParentYangName() string { return "router" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["metric"] = topology.Metric
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetParentYangName() string { return "link" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "attached-router" { return "AttachedRouter" }
    return ""
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetSegmentPath() string {
    return "network"
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = network.NetworkMask
    leafs["attached-router"] = network.AttachedRouter
    return leafs
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetBundleName() string { return "ietf" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetYangName() string { return "network" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetParent() types.Entity { return network.parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary
// Summary LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range summary.Topology {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology{}
        summary.Topology = append(summary.Topology, child)
        return &summary.Topology[len(summary.Topology)-1]
    }
    return nil
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.Topology {
        children[summary.Topology[i].GetSegmentPath()] = &summary.Topology[i]
    }
    return children
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = summary.NetworkMask
    return leafs
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetBundleName() string { return "ietf" }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetYangName() string { return "summary" }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetParent() types.Entity { return summary.parent }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["metric"] = topology.Metric
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetParentYangName() string { return "summary" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External
// External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetFilter() yfilter.YFilter { return external.YFilter }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) SetFilter(yf yfilter.YFilter) { external.YFilter = yf }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetSegmentPath() string {
    return "external"
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range external.Topology {
            if external.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology{}
        external.Topology = append(external.Topology, child)
        return &external.Topology[len(external.Topology)-1]
    }
    return nil
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range external.Topology {
        children[external.Topology[i].GetSegmentPath()] = &external.Topology[i]
    }
    return children
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = external.NetworkMask
    return leafs
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetBundleName() string { return "ietf" }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetYangName() string { return "external" }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) SetParent(parent types.Entity) { external.parent = parent }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetParent() types.Entity { return external.parent }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Forwarding address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "flags" { return "Flags" }
    if yname == "metric" { return "Metric" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["flags"] = topology.Flags
    leafs["metric"] = topology.Metric
    leafs["forwarding-address"] = topology.ForwardingAddress
    leafs["external-route-tag"] = topology.ExternalRouteTag
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetParentYangName() string { return "external" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque
// Opaque LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unknown TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv.
    UnknownTlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv

    // Router address TLV.
    RouterAddressTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv

    // Link TLV.
    LinkTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetFilter() yfilter.YFilter { return opaque.YFilter }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) SetFilter(yf yfilter.YFilter) { opaque.YFilter = yf }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetGoName(yname string) string {
    if yname == "unknown-tlv" { return "UnknownTlv" }
    if yname == "router-address-tlv" { return "RouterAddressTlv" }
    if yname == "link-tlv" { return "LinkTlv" }
    return ""
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetSegmentPath() string {
    return "opaque"
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-tlv" {
        for _, c := range opaque.UnknownTlv {
            if opaque.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv{}
        opaque.UnknownTlv = append(opaque.UnknownTlv, child)
        return &opaque.UnknownTlv[len(opaque.UnknownTlv)-1]
    }
    if childYangName == "router-address-tlv" {
        return &opaque.RouterAddressTlv
    }
    if childYangName == "link-tlv" {
        return &opaque.LinkTlv
    }
    return nil
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opaque.UnknownTlv {
        children[opaque.UnknownTlv[i].GetSegmentPath()] = &opaque.UnknownTlv[i]
    }
    children["router-address-tlv"] = &opaque.RouterAddressTlv
    children["link-tlv"] = &opaque.LinkTlv
    return children
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetBundleName() string { return "ietf" }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetYangName() string { return "opaque" }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) SetParent(parent types.Entity) { opaque.parent = parent }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetParent() types.Entity { return opaque.parent }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv
// Unknown TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetFilter() yfilter.YFilter { return unknownTlv.YFilter }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) SetFilter(yf yfilter.YFilter) { unknownTlv.YFilter = yf }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "length" { return "Length" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetSegmentPath() string {
    return "unknown-tlv" + "[type='" + fmt.Sprintf("%v", unknownTlv.Type) + "']"
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = unknownTlv.Type
    leafs["length"] = unknownTlv.Length
    leafs["value"] = unknownTlv.Value
    return leafs
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetBundleName() string { return "ietf" }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetYangName() string { return "unknown-tlv" }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) SetParent(parent types.Entity) { unknownTlv.parent = parent }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetParent() types.Entity { return unknownTlv.parent }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv
// Router address TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterAddress interface{}
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetFilter() yfilter.YFilter { return routerAddressTlv.YFilter }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) SetFilter(yf yfilter.YFilter) { routerAddressTlv.YFilter = yf }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetGoName(yname string) string {
    if yname == "router-address" { return "RouterAddress" }
    return ""
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetSegmentPath() string {
    return "router-address-tlv"
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-address"] = routerAddressTlv.RouterAddress
    return leafs
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetBundleName() string { return "ietf" }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetYangName() string { return "router-address-tlv" }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) SetParent(parent types.Entity) { routerAddressTlv.parent = parent }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetParent() types.Entity { return routerAddressTlv.parent }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv
// Link TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    LinkType interface{}

    // Link ID. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])
    // This attribute is mandatory..
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unreserved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}

    // Unknown sub-TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv.
    UnknownSubtlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetFilter() yfilter.YFilter { return linkTlv.YFilter }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) SetFilter(yf yfilter.YFilter) { linkTlv.YFilter = yf }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-id" { return "LinkId" }
    if yname == "local-if-ipv4-addr" { return "LocalIfIpv4Addr" }
    if yname == "local-remote-ipv4-addr" { return "LocalRemoteIpv4Addr" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "max-bandwidth" { return "MaxBandwidth" }
    if yname == "max-reservable-bandwidth" { return "MaxReservableBandwidth" }
    if yname == "unreserved-bandwidth" { return "UnreservedBandwidth" }
    if yname == "admin-group" { return "AdminGroup" }
    if yname == "unknown-subtlv" { return "UnknownSubtlv" }
    return ""
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetSegmentPath() string {
    return "link-tlv"
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-subtlv" {
        for _, c := range linkTlv.UnknownSubtlv {
            if linkTlv.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv{}
        linkTlv.UnknownSubtlv = append(linkTlv.UnknownSubtlv, child)
        return &linkTlv.UnknownSubtlv[len(linkTlv.UnknownSubtlv)-1]
    }
    return nil
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkTlv.UnknownSubtlv {
        children[linkTlv.UnknownSubtlv[i].GetSegmentPath()] = &linkTlv.UnknownSubtlv[i]
    }
    return children
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = linkTlv.LinkType
    leafs["link-id"] = linkTlv.LinkId
    leafs["local-if-ipv4-addr"] = linkTlv.LocalIfIpv4Addr
    leafs["local-remote-ipv4-addr"] = linkTlv.LocalRemoteIpv4Addr
    leafs["te-metric"] = linkTlv.TeMetric
    leafs["max-bandwidth"] = linkTlv.MaxBandwidth
    leafs["max-reservable-bandwidth"] = linkTlv.MaxReservableBandwidth
    leafs["unreserved-bandwidth"] = linkTlv.UnreservedBandwidth
    leafs["admin-group"] = linkTlv.AdminGroup
    return leafs
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetBundleName() string { return "ietf" }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetYangName() string { return "link-tlv" }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) SetParent(parent types.Entity) { linkTlv.parent = parent }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetParent() types.Entity { return linkTlv.parent }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
// Unknown sub-TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetFilter() yfilter.YFilter { return unknownSubtlv.YFilter }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) SetFilter(yf yfilter.YFilter) { unknownSubtlv.YFilter = yf }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "length" { return "Length" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetSegmentPath() string {
    return "unknown-subtlv" + "[type='" + fmt.Sprintf("%v", unknownSubtlv.Type) + "']"
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = unknownSubtlv.Type
    leafs["length"] = unknownSubtlv.Length
    leafs["value"] = unknownSubtlv.Value
    return leafs
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetBundleName() string { return "ietf" }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetYangName() string { return "unknown-subtlv" }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) SetParent(parent types.Entity) { unknownSubtlv.parent = parent }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetParent() types.Entity { return unknownSubtlv.parent }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetParentYangName() string { return "link-tlv" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3
// OSPFv3 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header

    // Decoded OSPF LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetFilter() yfilter.YFilter { return ospfv3.YFilter }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) SetFilter(yf yfilter.YFilter) { ospfv3.YFilter = yf }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "body" { return "Body" }
    return ""
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetSegmentPath() string {
    return "ospfv3"
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &ospfv3.Header
    }
    if childYangName == "body" {
        return &ospfv3.Body
    }
    return nil
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &ospfv3.Header
    children["body"] = &ospfv3.Body
    return children
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetBundleName() string { return "ietf" }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetYangName() string { return "ospfv3" }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) SetParent(parent types.Entity) { ospfv3.parent = parent }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetParent() types.Entity { return ospfv3.parent }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetParentYangName() string { return "area-scope-lsa" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header
// Decoded OSPFv3 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA ID. The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory.
    LsaId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type interface{}

    // LSA advertising router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetGoName(yname string) string {
    if yname == "lsa-id" { return "LsaId" }
    if yname == "age" { return "Age" }
    if yname == "type" { return "Type" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "seq-num" { return "SeqNum" }
    if yname == "checksum" { return "Checksum" }
    if yname == "length" { return "Length" }
    if yname == "options" { return "Options" }
    return ""
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetSegmentPath() string {
    return "header"
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-id"] = header.LsaId
    leafs["age"] = header.Age
    leafs["type"] = header.Type
    leafs["adv-router"] = header.AdvRouter
    leafs["seq-num"] = header.SeqNum
    leafs["checksum"] = header.Checksum
    leafs["length"] = header.Length
    leafs["options"] = header.Options
    return leafs
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetBundleName() string { return "ietf" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetYangName() string { return "header" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetParent() types.Entity { return header.parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetParentYangName() string { return "ospfv3" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body
// Decoded OSPF LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network

    // Inter-Area-Prefix LSA.
    InterAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix

    // Inter-Area-Router LSA.
    InterAreaRouter RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter

    // AS-External LSA.
    AsExternal RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal

    // NSSA LSA.
    Nssa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa

    // Link LSA.
    Link RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link

    // Intra-Area-Prefix LSA.
    IntraAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetFilter() yfilter.YFilter { return body.YFilter }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) SetFilter(yf yfilter.YFilter) { body.YFilter = yf }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetGoName(yname string) string {
    if yname == "router" { return "Router" }
    if yname == "network" { return "Network" }
    if yname == "inter-area-prefix" { return "InterAreaPrefix" }
    if yname == "inter-area-router" { return "InterAreaRouter" }
    if yname == "as-external" { return "AsExternal" }
    if yname == "nssa" { return "Nssa" }
    if yname == "link" { return "Link" }
    if yname == "intra-area-prefix" { return "IntraAreaPrefix" }
    return ""
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetSegmentPath() string {
    return "body"
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router" {
        return &body.Router
    }
    if childYangName == "network" {
        return &body.Network
    }
    if childYangName == "inter-area-prefix" {
        return &body.InterAreaPrefix
    }
    if childYangName == "inter-area-router" {
        return &body.InterAreaRouter
    }
    if childYangName == "as-external" {
        return &body.AsExternal
    }
    if childYangName == "nssa" {
        return &body.Nssa
    }
    if childYangName == "link" {
        return &body.Link
    }
    if childYangName == "intra-area-prefix" {
        return &body.IntraAreaPrefix
    }
    return nil
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router"] = &body.Router
    children["network"] = &body.Network
    children["inter-area-prefix"] = &body.InterAreaPrefix
    children["inter-area-router"] = &body.InterAreaRouter
    children["as-external"] = &body.AsExternal
    children["nssa"] = &body.Nssa
    children["link"] = &body.Link
    children["intra-area-prefix"] = &body.IntraAreaPrefix
    return children
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetBundleName() string { return "ietf" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetYangName() string { return "body" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) SetParent(parent types.Entity) { body.parent = parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetParent() types.Entity { return body.parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetParentYangName() string { return "ospfv3" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Flags interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetFilter() yfilter.YFilter { return router.YFilter }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) SetFilter(yf yfilter.YFilter) { router.YFilter = yf }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetGoName(yname string) string {
    if yname == "flags" { return "Flags" }
    if yname == "options" { return "Options" }
    if yname == "link" { return "Link" }
    return ""
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetSegmentPath() string {
    return "router"
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link" {
        for _, c := range router.Link {
            if router.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link{}
        router.Link = append(router.Link, child)
        return &router.Link[len(router.Link)-1]
    }
    return nil
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range router.Link {
        children[router.Link[i].GetSegmentPath()] = &router.Link[i]
    }
    return children
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flags"] = router.Flags
    leafs["options"] = router.Options
    return leafs
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetBundleName() string { return "ietf" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetYangName() string { return "router" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) SetParent(parent types.Entity) { router.parent = parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetParent() types.Entity { return router.parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor Interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor Router ID. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetGoName(yname string) string {
    if yname == "interface-id" { return "InterfaceId" }
    if yname == "neighbor-interface-id" { return "NeighborInterfaceId" }
    if yname == "neighbor-router-id" { return "NeighborRouterId" }
    if yname == "type" { return "Type" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetSegmentPath() string {
    return "link" + "[interface-id='" + fmt.Sprintf("%v", link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", link.NeighborRouterId) + "']"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-id"] = link.InterfaceId
    leafs["neighbor-interface-id"] = link.NeighborInterfaceId
    leafs["neighbor-router-id"] = link.NeighborRouterId
    leafs["type"] = link.Type
    leafs["metric"] = link.Metric
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetParentYangName() string { return "router" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "attached-router" { return "AttachedRouter" }
    return ""
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetSegmentPath() string {
    return "network"
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = network.Options
    leafs["attached-router"] = network.AttachedRouter
    return leafs
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetBundleName() string { return "ietf" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetYangName() string { return "network" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetParent() types.Entity { return network.parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix
// Inter-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetFilter() yfilter.YFilter { return interAreaPrefix.YFilter }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) SetFilter(yf yfilter.YFilter) { interAreaPrefix.YFilter = yf }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    return ""
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetSegmentPath() string {
    return "inter-area-prefix"
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = interAreaPrefix.Metric
    leafs["prefix"] = interAreaPrefix.Prefix
    leafs["prefix-options"] = interAreaPrefix.PrefixOptions
    return leafs
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetBundleName() string { return "ietf" }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetYangName() string { return "inter-area-prefix" }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) SetParent(parent types.Entity) { interAreaPrefix.parent = parent }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetParent() types.Entity { return interAreaPrefix.parent }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter
// Inter-Area-Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // The Router ID of the router being described by the LSA. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    DestinationRouterId interface{}
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetFilter() yfilter.YFilter { return interAreaRouter.YFilter }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) SetFilter(yf yfilter.YFilter) { interAreaRouter.YFilter = yf }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "metric" { return "Metric" }
    if yname == "destination-router-id" { return "DestinationRouterId" }
    return ""
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetSegmentPath() string {
    return "inter-area-router"
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = interAreaRouter.Options
    leafs["metric"] = interAreaRouter.Metric
    leafs["destination-router-id"] = interAreaRouter.DestinationRouterId
    return leafs
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetBundleName() string { return "ietf" }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetYangName() string { return "inter-area-router" }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) SetParent(parent types.Entity) { interAreaRouter.parent = parent }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetParent() types.Entity { return interAreaRouter.parent }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal
// AS-External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetFilter() yfilter.YFilter { return asExternal.YFilter }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) SetFilter(yf yfilter.YFilter) { asExternal.YFilter = yf }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "flags" { return "Flags" }
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    return ""
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetSegmentPath() string {
    return "as-external"
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = asExternal.Metric
    leafs["flags"] = asExternal.Flags
    leafs["referenced-ls-type"] = asExternal.ReferencedLsType
    leafs["prefix"] = asExternal.Prefix
    leafs["prefix-options"] = asExternal.PrefixOptions
    leafs["forwarding-address"] = asExternal.ForwardingAddress
    leafs["external-route-tag"] = asExternal.ExternalRouteTag
    leafs["referenced-link-state-id"] = asExternal.ReferencedLinkStateId
    return leafs
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetBundleName() string { return "ietf" }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetYangName() string { return "as-external" }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) SetParent(parent types.Entity) { asExternal.parent = parent }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetParent() types.Entity { return asExternal.parent }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa
// NSSA LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetFilter() yfilter.YFilter { return nssa.YFilter }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) SetFilter(yf yfilter.YFilter) { nssa.YFilter = yf }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "flags" { return "Flags" }
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    return ""
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetSegmentPath() string {
    return "nssa"
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = nssa.Metric
    leafs["flags"] = nssa.Flags
    leafs["referenced-ls-type"] = nssa.ReferencedLsType
    leafs["prefix"] = nssa.Prefix
    leafs["prefix-options"] = nssa.PrefixOptions
    leafs["forwarding-address"] = nssa.ForwardingAddress
    leafs["external-route-tag"] = nssa.ExternalRouteTag
    leafs["referenced-link-state-id"] = nssa.ReferencedLinkStateId
    return leafs
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetBundleName() string { return "ietf" }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetYangName() string { return "nssa" }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) SetParent(parent types.Entity) { nssa.parent = parent }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetParent() types.Entity { return nssa.parent }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link
// Link LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router Priority of the interface. The type is interface{} with range:
    // 0..255.
    RtrPriority interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetGoName(yname string) string {
    if yname == "rtr-priority" { return "RtrPriority" }
    if yname == "options" { return "Options" }
    if yname == "link-local-interface-address" { return "LinkLocalInterfaceAddress" }
    if yname == "num-of-prefixes" { return "NumOfPrefixes" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetSegmentPath() string {
    return "link"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list" {
        for _, c := range link.PrefixList {
            if link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList{}
        link.PrefixList = append(link.PrefixList, child)
        return &link.PrefixList[len(link.PrefixList)-1]
    }
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range link.PrefixList {
        children[link.PrefixList[i].GetSegmentPath()] = &link.PrefixList[i]
    }
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rtr-priority"] = link.RtrPriority
    leafs["options"] = link.Options
    leafs["link-local-interface-address"] = link.LinkLocalInterfaceAddress
    leafs["num-of-prefixes"] = link.NumOfPrefixes
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetFilter() yfilter.YFilter { return prefixList.YFilter }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) SetFilter(yf yfilter.YFilter) { prefixList.YFilter = yf }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    return ""
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetSegmentPath() string {
    return "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = prefixList.Prefix
    leafs["prefix-options"] = prefixList.PrefixOptions
    return leafs
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetBundleName() string { return "ietf" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetYangName() string { return "prefix-list" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) SetParent(parent types.Entity) { prefixList.parent = parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetParent() types.Entity { return prefixList.parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetParentYangName() string { return "link" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix
// Intra-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetFilter() yfilter.YFilter { return intraAreaPrefix.YFilter }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) SetFilter(yf yfilter.YFilter) { intraAreaPrefix.YFilter = yf }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetGoName(yname string) string {
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    if yname == "referenced-adv-router" { return "ReferencedAdvRouter" }
    if yname == "num-of-prefixes" { return "NumOfPrefixes" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetSegmentPath() string {
    return "intra-area-prefix"
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list" {
        for _, c := range intraAreaPrefix.PrefixList {
            if intraAreaPrefix.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList{}
        intraAreaPrefix.PrefixList = append(intraAreaPrefix.PrefixList, child)
        return &intraAreaPrefix.PrefixList[len(intraAreaPrefix.PrefixList)-1]
    }
    return nil
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intraAreaPrefix.PrefixList {
        children[intraAreaPrefix.PrefixList[i].GetSegmentPath()] = &intraAreaPrefix.PrefixList[i]
    }
    return children
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["referenced-ls-type"] = intraAreaPrefix.ReferencedLsType
    leafs["referenced-link-state-id"] = intraAreaPrefix.ReferencedLinkStateId
    leafs["referenced-adv-router"] = intraAreaPrefix.ReferencedAdvRouter
    leafs["num-of-prefixes"] = intraAreaPrefix.NumOfPrefixes
    return leafs
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetBundleName() string { return "ietf" }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetYangName() string { return "intra-area-prefix" }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) SetParent(parent types.Entity) { intraAreaPrefix.parent = parent }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetParent() types.Entity { return intraAreaPrefix.parent }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetFilter() yfilter.YFilter { return prefixList.YFilter }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) SetFilter(yf yfilter.YFilter) { prefixList.YFilter = yf }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetSegmentPath() string {
    return "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = prefixList.Prefix
    leafs["prefix-options"] = prefixList.PrefixOptions
    leafs["metric"] = prefixList.Metric
    return leafs
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetBundleName() string { return "ietf" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetYangName() string { return "prefix-list" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) SetParent(parent types.Entity) { prefixList.parent = parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetParent() types.Entity { return prefixList.parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetParentYangName() string { return "intra-area-prefix" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas
// List OSPF AS scope LSA databases
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF AS scope LSA type. The type is interface{}
    // with range: 0..255.
    LsaType interface{}

    // List of OSPF AS scope LSAs. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa.
    AsScopeLsa []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa
}

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetFilter() yfilter.YFilter { return asScopeLsas.YFilter }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) SetFilter(yf yfilter.YFilter) { asScopeLsas.YFilter = yf }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetGoName(yname string) string {
    if yname == "lsa-type" { return "LsaType" }
    if yname == "as-scope-lsa" { return "AsScopeLsa" }
    return ""
}

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetSegmentPath() string {
    return "as-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", asScopeLsas.LsaType) + "']"
}

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "as-scope-lsa" {
        for _, c := range asScopeLsas.AsScopeLsa {
            if asScopeLsas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa{}
        asScopeLsas.AsScopeLsa = append(asScopeLsas.AsScopeLsa, child)
        return &asScopeLsas.AsScopeLsa[len(asScopeLsas.AsScopeLsa)-1]
    }
    return nil
}

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asScopeLsas.AsScopeLsa {
        children[asScopeLsas.AsScopeLsa[i].GetSegmentPath()] = &asScopeLsas.AsScopeLsa[i]
    }
    return children
}

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-type"] = asScopeLsas.LsaType
    return leafs
}

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetBundleName() string { return "ietf" }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetYangName() string { return "as-scope-lsas" }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) SetParent(parent types.Entity) { asScopeLsas.parent = parent }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetParent() types.Entity { return asScopeLsas.parent }

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetParentYangName() string { return "instance" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa
// List of OSPF AS scope LSAs
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or int with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RawData interface{}

    // OSPFv2 LSA.
    Ospfv2 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2

    // OSPFv3 LSA.
    Ospfv3 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3
}

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetFilter() yfilter.YFilter { return asScopeLsa.YFilter }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) SetFilter(yf yfilter.YFilter) { asScopeLsa.YFilter = yf }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetGoName(yname string) string {
    if yname == "lsa-id" { return "LsaId" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "decoded-completed" { return "DecodedCompleted" }
    if yname == "raw-data" { return "RawData" }
    if yname == "ospfv2" { return "Ospfv2" }
    if yname == "ospfv3" { return "Ospfv3" }
    return ""
}

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetSegmentPath() string {
    return "as-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", asScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", asScopeLsa.AdvRouter) + "']"
}

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2" {
        return &asScopeLsa.Ospfv2
    }
    if childYangName == "ospfv3" {
        return &asScopeLsa.Ospfv3
    }
    return nil
}

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv2"] = &asScopeLsa.Ospfv2
    children["ospfv3"] = &asScopeLsa.Ospfv3
    return children
}

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-id"] = asScopeLsa.LsaId
    leafs["adv-router"] = asScopeLsa.AdvRouter
    leafs["decoded-completed"] = asScopeLsa.DecodedCompleted
    leafs["raw-data"] = asScopeLsa.RawData
    return leafs
}

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetBundleName() string { return "ietf" }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetYangName() string { return "as-scope-lsa" }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) SetParent(parent types.Entity) { asScopeLsa.parent = parent }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetParent() types.Entity { return asScopeLsa.parent }

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetParentYangName() string { return "as-scope-lsas" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2
// OSPFv2 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header

    // Decoded OSPFv2 LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetFilter() yfilter.YFilter { return ospfv2.YFilter }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) SetFilter(yf yfilter.YFilter) { ospfv2.YFilter = yf }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "body" { return "Body" }
    return ""
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetSegmentPath() string {
    return "ospfv2"
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &ospfv2.Header
    }
    if childYangName == "body" {
        return &ospfv2.Body
    }
    return nil
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &ospfv2.Header
    children["body"] = &ospfv2.Body
    return children
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetBundleName() string { return "ietf" }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetYangName() string { return "ospfv2" }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) SetParent(parent types.Entity) { ospfv2.parent = parent }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetParent() types.Entity { return ospfv2.parent }

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetParentYangName() string { return "as-scope-lsa" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header
// Decoded OSPFv2 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Options interface{}

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    OpaqueType interface{}

    // Opaque id. The type is interface{} with range: 0..16777215. This attribute
    // is mandatory.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type interface{}

    // LSA advertising router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "lsa-id" { return "LsaId" }
    if yname == "opaque-type" { return "OpaqueType" }
    if yname == "opaque-id" { return "OpaqueId" }
    if yname == "age" { return "Age" }
    if yname == "type" { return "Type" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "seq-num" { return "SeqNum" }
    if yname == "checksum" { return "Checksum" }
    if yname == "length" { return "Length" }
    return ""
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetSegmentPath() string {
    return "header"
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = header.Options
    leafs["lsa-id"] = header.LsaId
    leafs["opaque-type"] = header.OpaqueType
    leafs["opaque-id"] = header.OpaqueId
    leafs["age"] = header.Age
    leafs["type"] = header.Type
    leafs["adv-router"] = header.AdvRouter
    leafs["seq-num"] = header.SeqNum
    leafs["checksum"] = header.Checksum
    leafs["length"] = header.Length
    return leafs
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetBundleName() string { return "ietf" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetYangName() string { return "header" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetParent() types.Entity { return header.parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetParentYangName() string { return "ospfv2" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body
// Decoded OSPFv2 LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network

    // Summary LSA.
    Summary RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary

    // External LSA.
    External RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External

    // Opaque LSA.
    Opaque RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetFilter() yfilter.YFilter { return body.YFilter }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) SetFilter(yf yfilter.YFilter) { body.YFilter = yf }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetGoName(yname string) string {
    if yname == "router" { return "Router" }
    if yname == "network" { return "Network" }
    if yname == "summary" { return "Summary" }
    if yname == "external" { return "External" }
    if yname == "opaque" { return "Opaque" }
    return ""
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetSegmentPath() string {
    return "body"
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router" {
        return &body.Router
    }
    if childYangName == "network" {
        return &body.Network
    }
    if childYangName == "summary" {
        return &body.Summary
    }
    if childYangName == "external" {
        return &body.External
    }
    if childYangName == "opaque" {
        return &body.Opaque
    }
    return nil
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router"] = &body.Router
    children["network"] = &body.Network
    children["summary"] = &body.Summary
    children["external"] = &body.External
    children["opaque"] = &body.Opaque
    return children
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetBundleName() string { return "ietf" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetYangName() string { return "body" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) SetParent(parent types.Entity) { body.parent = parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetParent() types.Entity { return body.parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetParentYangName() string { return "ospfv2" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetFilter() yfilter.YFilter { return router.YFilter }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) SetFilter(yf yfilter.YFilter) { router.YFilter = yf }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetGoName(yname string) string {
    if yname == "flags" { return "Flags" }
    if yname == "num-of-links" { return "NumOfLinks" }
    if yname == "link" { return "Link" }
    return ""
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetSegmentPath() string {
    return "router"
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link" {
        for _, c := range router.Link {
            if router.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link{}
        router.Link = append(router.Link, child)
        return &router.Link[len(router.Link)-1]
    }
    return nil
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range router.Link {
        children[router.Link[i].GetSegmentPath()] = &router.Link[i]
    }
    return children
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flags"] = router.Flags
    leafs["num-of-links"] = router.NumOfLinks
    return leafs
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetBundleName() string { return "ietf" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetYangName() string { return "router" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) SetParent(parent types.Entity) { router.parent = parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetParent() types.Entity { return router.parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    LinkId interface{}

    // This attribute is a key. Link data. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or int with range: 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetGoName(yname string) string {
    if yname == "link-id" { return "LinkId" }
    if yname == "link-data" { return "LinkData" }
    if yname == "type" { return "Type" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetSegmentPath() string {
    return "link" + "[link-id='" + fmt.Sprintf("%v", link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", link.LinkData) + "']"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range link.Topology {
            if link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology{}
        link.Topology = append(link.Topology, child)
        return &link.Topology[len(link.Topology)-1]
    }
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range link.Topology {
        children[link.Topology[i].GetSegmentPath()] = &link.Topology[i]
    }
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-id"] = link.LinkId
    leafs["link-data"] = link.LinkData
    leafs["type"] = link.Type
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetParentYangName() string { return "router" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["metric"] = topology.Metric
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetParentYangName() string { return "link" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "attached-router" { return "AttachedRouter" }
    return ""
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetSegmentPath() string {
    return "network"
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = network.NetworkMask
    leafs["attached-router"] = network.AttachedRouter
    return leafs
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetBundleName() string { return "ietf" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetYangName() string { return "network" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetParent() types.Entity { return network.parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary
// Summary LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range summary.Topology {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology{}
        summary.Topology = append(summary.Topology, child)
        return &summary.Topology[len(summary.Topology)-1]
    }
    return nil
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.Topology {
        children[summary.Topology[i].GetSegmentPath()] = &summary.Topology[i]
    }
    return children
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = summary.NetworkMask
    return leafs
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetBundleName() string { return "ietf" }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetYangName() string { return "summary" }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetParent() types.Entity { return summary.parent }

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["metric"] = topology.Metric
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetParentYangName() string { return "summary" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External
// External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetFilter() yfilter.YFilter { return external.YFilter }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) SetFilter(yf yfilter.YFilter) { external.YFilter = yf }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetGoName(yname string) string {
    if yname == "network-mask" { return "NetworkMask" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetSegmentPath() string {
    return "external"
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology" {
        for _, c := range external.Topology {
            if external.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology{}
        external.Topology = append(external.Topology, child)
        return &external.Topology[len(external.Topology)-1]
    }
    return nil
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range external.Topology {
        children[external.Topology[i].GetSegmentPath()] = &external.Topology[i]
    }
    return children
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-mask"] = external.NetworkMask
    return leafs
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetBundleName() string { return "ietf" }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetYangName() string { return "external" }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) SetParent(parent types.Entity) { external.parent = parent }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetParent() types.Entity { return external.parent }

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Forwarding address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetGoName(yname string) string {
    if yname == "mt-id" { return "MtId" }
    if yname == "flags" { return "Flags" }
    if yname == "metric" { return "Metric" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetSegmentPath() string {
    return "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mt-id"] = topology.MtId
    leafs["flags"] = topology.Flags
    leafs["metric"] = topology.Metric
    leafs["forwarding-address"] = topology.ForwardingAddress
    leafs["external-route-tag"] = topology.ExternalRouteTag
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetParentYangName() string { return "external" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque
// Opaque LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unknown TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv.
    UnknownTlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv

    // Router address TLV.
    RouterAddressTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv

    // Link TLV.
    LinkTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetFilter() yfilter.YFilter { return opaque.YFilter }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) SetFilter(yf yfilter.YFilter) { opaque.YFilter = yf }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetGoName(yname string) string {
    if yname == "unknown-tlv" { return "UnknownTlv" }
    if yname == "router-address-tlv" { return "RouterAddressTlv" }
    if yname == "link-tlv" { return "LinkTlv" }
    return ""
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetSegmentPath() string {
    return "opaque"
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-tlv" {
        for _, c := range opaque.UnknownTlv {
            if opaque.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv{}
        opaque.UnknownTlv = append(opaque.UnknownTlv, child)
        return &opaque.UnknownTlv[len(opaque.UnknownTlv)-1]
    }
    if childYangName == "router-address-tlv" {
        return &opaque.RouterAddressTlv
    }
    if childYangName == "link-tlv" {
        return &opaque.LinkTlv
    }
    return nil
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opaque.UnknownTlv {
        children[opaque.UnknownTlv[i].GetSegmentPath()] = &opaque.UnknownTlv[i]
    }
    children["router-address-tlv"] = &opaque.RouterAddressTlv
    children["link-tlv"] = &opaque.LinkTlv
    return children
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetBundleName() string { return "ietf" }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetYangName() string { return "opaque" }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) SetParent(parent types.Entity) { opaque.parent = parent }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetParent() types.Entity { return opaque.parent }

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv
// Unknown TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetFilter() yfilter.YFilter { return unknownTlv.YFilter }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) SetFilter(yf yfilter.YFilter) { unknownTlv.YFilter = yf }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "length" { return "Length" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetSegmentPath() string {
    return "unknown-tlv" + "[type='" + fmt.Sprintf("%v", unknownTlv.Type) + "']"
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = unknownTlv.Type
    leafs["length"] = unknownTlv.Length
    leafs["value"] = unknownTlv.Value
    return leafs
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetBundleName() string { return "ietf" }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetYangName() string { return "unknown-tlv" }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) SetParent(parent types.Entity) { unknownTlv.parent = parent }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetParent() types.Entity { return unknownTlv.parent }

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv
// Router address TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterAddress interface{}
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetFilter() yfilter.YFilter { return routerAddressTlv.YFilter }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) SetFilter(yf yfilter.YFilter) { routerAddressTlv.YFilter = yf }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetGoName(yname string) string {
    if yname == "router-address" { return "RouterAddress" }
    return ""
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetSegmentPath() string {
    return "router-address-tlv"
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-address"] = routerAddressTlv.RouterAddress
    return leafs
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetBundleName() string { return "ietf" }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetYangName() string { return "router-address-tlv" }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) SetParent(parent types.Entity) { routerAddressTlv.parent = parent }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetParent() types.Entity { return routerAddressTlv.parent }

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv
// Link TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    LinkType interface{}

    // Link ID. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])
    // This attribute is mandatory..
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unreserved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}

    // Unknown sub-TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv.
    UnknownSubtlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetFilter() yfilter.YFilter { return linkTlv.YFilter }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) SetFilter(yf yfilter.YFilter) { linkTlv.YFilter = yf }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-id" { return "LinkId" }
    if yname == "local-if-ipv4-addr" { return "LocalIfIpv4Addr" }
    if yname == "local-remote-ipv4-addr" { return "LocalRemoteIpv4Addr" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "max-bandwidth" { return "MaxBandwidth" }
    if yname == "max-reservable-bandwidth" { return "MaxReservableBandwidth" }
    if yname == "unreserved-bandwidth" { return "UnreservedBandwidth" }
    if yname == "admin-group" { return "AdminGroup" }
    if yname == "unknown-subtlv" { return "UnknownSubtlv" }
    return ""
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetSegmentPath() string {
    return "link-tlv"
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-subtlv" {
        for _, c := range linkTlv.UnknownSubtlv {
            if linkTlv.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv{}
        linkTlv.UnknownSubtlv = append(linkTlv.UnknownSubtlv, child)
        return &linkTlv.UnknownSubtlv[len(linkTlv.UnknownSubtlv)-1]
    }
    return nil
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkTlv.UnknownSubtlv {
        children[linkTlv.UnknownSubtlv[i].GetSegmentPath()] = &linkTlv.UnknownSubtlv[i]
    }
    return children
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = linkTlv.LinkType
    leafs["link-id"] = linkTlv.LinkId
    leafs["local-if-ipv4-addr"] = linkTlv.LocalIfIpv4Addr
    leafs["local-remote-ipv4-addr"] = linkTlv.LocalRemoteIpv4Addr
    leafs["te-metric"] = linkTlv.TeMetric
    leafs["max-bandwidth"] = linkTlv.MaxBandwidth
    leafs["max-reservable-bandwidth"] = linkTlv.MaxReservableBandwidth
    leafs["unreserved-bandwidth"] = linkTlv.UnreservedBandwidth
    leafs["admin-group"] = linkTlv.AdminGroup
    return leafs
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetBundleName() string { return "ietf" }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetYangName() string { return "link-tlv" }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) SetParent(parent types.Entity) { linkTlv.parent = parent }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetParent() types.Entity { return linkTlv.parent }

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetParentYangName() string { return "opaque" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
// Unknown sub-TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetFilter() yfilter.YFilter { return unknownSubtlv.YFilter }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) SetFilter(yf yfilter.YFilter) { unknownSubtlv.YFilter = yf }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "length" { return "Length" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetSegmentPath() string {
    return "unknown-subtlv" + "[type='" + fmt.Sprintf("%v", unknownSubtlv.Type) + "']"
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = unknownSubtlv.Type
    leafs["length"] = unknownSubtlv.Length
    leafs["value"] = unknownSubtlv.Value
    return leafs
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetBundleName() string { return "ietf" }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetYangName() string { return "unknown-subtlv" }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) SetParent(parent types.Entity) { unknownSubtlv.parent = parent }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetParent() types.Entity { return unknownSubtlv.parent }

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetParentYangName() string { return "link-tlv" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3
// OSPFv3 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header

    // Decoded OSPF LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetFilter() yfilter.YFilter { return ospfv3.YFilter }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) SetFilter(yf yfilter.YFilter) { ospfv3.YFilter = yf }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "body" { return "Body" }
    return ""
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetSegmentPath() string {
    return "ospfv3"
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &ospfv3.Header
    }
    if childYangName == "body" {
        return &ospfv3.Body
    }
    return nil
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &ospfv3.Header
    children["body"] = &ospfv3.Body
    return children
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetBundleName() string { return "ietf" }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetYangName() string { return "ospfv3" }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) SetParent(parent types.Entity) { ospfv3.parent = parent }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetParent() types.Entity { return ospfv3.parent }

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetParentYangName() string { return "as-scope-lsa" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header
// Decoded OSPFv3 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA ID. The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory.
    LsaId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type interface{}

    // LSA advertising router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetGoName(yname string) string {
    if yname == "lsa-id" { return "LsaId" }
    if yname == "age" { return "Age" }
    if yname == "type" { return "Type" }
    if yname == "adv-router" { return "AdvRouter" }
    if yname == "seq-num" { return "SeqNum" }
    if yname == "checksum" { return "Checksum" }
    if yname == "length" { return "Length" }
    if yname == "options" { return "Options" }
    return ""
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetSegmentPath() string {
    return "header"
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsa-id"] = header.LsaId
    leafs["age"] = header.Age
    leafs["type"] = header.Type
    leafs["adv-router"] = header.AdvRouter
    leafs["seq-num"] = header.SeqNum
    leafs["checksum"] = header.Checksum
    leafs["length"] = header.Length
    leafs["options"] = header.Options
    return leafs
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetBundleName() string { return "ietf" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetYangName() string { return "header" }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetParent() types.Entity { return header.parent }

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetParentYangName() string { return "ospfv3" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body
// Decoded OSPF LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network

    // Inter-Area-Prefix LSA.
    InterAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix

    // Inter-Area-Router LSA.
    InterAreaRouter RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter

    // AS-External LSA.
    AsExternal RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal

    // NSSA LSA.
    Nssa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa

    // Link LSA.
    Link RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link

    // Intra-Area-Prefix LSA.
    IntraAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetFilter() yfilter.YFilter { return body.YFilter }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) SetFilter(yf yfilter.YFilter) { body.YFilter = yf }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetGoName(yname string) string {
    if yname == "router" { return "Router" }
    if yname == "network" { return "Network" }
    if yname == "inter-area-prefix" { return "InterAreaPrefix" }
    if yname == "inter-area-router" { return "InterAreaRouter" }
    if yname == "as-external" { return "AsExternal" }
    if yname == "nssa" { return "Nssa" }
    if yname == "link" { return "Link" }
    if yname == "intra-area-prefix" { return "IntraAreaPrefix" }
    return ""
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetSegmentPath() string {
    return "body"
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router" {
        return &body.Router
    }
    if childYangName == "network" {
        return &body.Network
    }
    if childYangName == "inter-area-prefix" {
        return &body.InterAreaPrefix
    }
    if childYangName == "inter-area-router" {
        return &body.InterAreaRouter
    }
    if childYangName == "as-external" {
        return &body.AsExternal
    }
    if childYangName == "nssa" {
        return &body.Nssa
    }
    if childYangName == "link" {
        return &body.Link
    }
    if childYangName == "intra-area-prefix" {
        return &body.IntraAreaPrefix
    }
    return nil
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router"] = &body.Router
    children["network"] = &body.Network
    children["inter-area-prefix"] = &body.InterAreaPrefix
    children["inter-area-router"] = &body.InterAreaRouter
    children["as-external"] = &body.AsExternal
    children["nssa"] = &body.Nssa
    children["link"] = &body.Link
    children["intra-area-prefix"] = &body.IntraAreaPrefix
    return children
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetBundleName() string { return "ietf" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetYangName() string { return "body" }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) SetParent(parent types.Entity) { body.parent = parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetParent() types.Entity { return body.parent }

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetParentYangName() string { return "ospfv3" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Flags interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetFilter() yfilter.YFilter { return router.YFilter }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) SetFilter(yf yfilter.YFilter) { router.YFilter = yf }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetGoName(yname string) string {
    if yname == "flags" { return "Flags" }
    if yname == "options" { return "Options" }
    if yname == "link" { return "Link" }
    return ""
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetSegmentPath() string {
    return "router"
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link" {
        for _, c := range router.Link {
            if router.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link{}
        router.Link = append(router.Link, child)
        return &router.Link[len(router.Link)-1]
    }
    return nil
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range router.Link {
        children[router.Link[i].GetSegmentPath()] = &router.Link[i]
    }
    return children
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flags"] = router.Flags
    leafs["options"] = router.Options
    return leafs
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetBundleName() string { return "ietf" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetYangName() string { return "router" }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) SetParent(parent types.Entity) { router.parent = parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetParent() types.Entity { return router.parent }

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor Interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor Router ID. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetGoName(yname string) string {
    if yname == "interface-id" { return "InterfaceId" }
    if yname == "neighbor-interface-id" { return "NeighborInterfaceId" }
    if yname == "neighbor-router-id" { return "NeighborRouterId" }
    if yname == "type" { return "Type" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetSegmentPath() string {
    return "link" + "[interface-id='" + fmt.Sprintf("%v", link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", link.NeighborRouterId) + "']"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-id"] = link.InterfaceId
    leafs["neighbor-interface-id"] = link.NeighborInterfaceId
    leafs["neighbor-router-id"] = link.NeighborRouterId
    leafs["type"] = link.Type
    leafs["metric"] = link.Metric
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetParentYangName() string { return "router" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "attached-router" { return "AttachedRouter" }
    return ""
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetSegmentPath() string {
    return "network"
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = network.Options
    leafs["attached-router"] = network.AttachedRouter
    return leafs
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetBundleName() string { return "ietf" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetYangName() string { return "network" }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetParent() types.Entity { return network.parent }

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix
// Inter-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetFilter() yfilter.YFilter { return interAreaPrefix.YFilter }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) SetFilter(yf yfilter.YFilter) { interAreaPrefix.YFilter = yf }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    return ""
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetSegmentPath() string {
    return "inter-area-prefix"
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = interAreaPrefix.Metric
    leafs["prefix"] = interAreaPrefix.Prefix
    leafs["prefix-options"] = interAreaPrefix.PrefixOptions
    return leafs
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetBundleName() string { return "ietf" }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetYangName() string { return "inter-area-prefix" }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) SetParent(parent types.Entity) { interAreaPrefix.parent = parent }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetParent() types.Entity { return interAreaPrefix.parent }

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter
// Inter-Area-Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // The Router ID of the router being described by the LSA. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    DestinationRouterId interface{}
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetFilter() yfilter.YFilter { return interAreaRouter.YFilter }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) SetFilter(yf yfilter.YFilter) { interAreaRouter.YFilter = yf }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    if yname == "metric" { return "Metric" }
    if yname == "destination-router-id" { return "DestinationRouterId" }
    return ""
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetSegmentPath() string {
    return "inter-area-router"
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["options"] = interAreaRouter.Options
    leafs["metric"] = interAreaRouter.Metric
    leafs["destination-router-id"] = interAreaRouter.DestinationRouterId
    return leafs
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetBundleName() string { return "ietf" }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetYangName() string { return "inter-area-router" }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) SetParent(parent types.Entity) { interAreaRouter.parent = parent }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetParent() types.Entity { return interAreaRouter.parent }

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal
// AS-External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetFilter() yfilter.YFilter { return asExternal.YFilter }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) SetFilter(yf yfilter.YFilter) { asExternal.YFilter = yf }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "flags" { return "Flags" }
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    return ""
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetSegmentPath() string {
    return "as-external"
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = asExternal.Metric
    leafs["flags"] = asExternal.Flags
    leafs["referenced-ls-type"] = asExternal.ReferencedLsType
    leafs["prefix"] = asExternal.Prefix
    leafs["prefix-options"] = asExternal.PrefixOptions
    leafs["forwarding-address"] = asExternal.ForwardingAddress
    leafs["external-route-tag"] = asExternal.ExternalRouteTag
    leafs["referenced-link-state-id"] = asExternal.ReferencedLinkStateId
    return leafs
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetBundleName() string { return "ietf" }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetYangName() string { return "as-external" }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) SetParent(parent types.Entity) { asExternal.parent = parent }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetParent() types.Entity { return asExternal.parent }

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa
// NSSA LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetFilter() yfilter.YFilter { return nssa.YFilter }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) SetFilter(yf yfilter.YFilter) { nssa.YFilter = yf }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "flags" { return "Flags" }
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "forwarding-address" { return "ForwardingAddress" }
    if yname == "external-route-tag" { return "ExternalRouteTag" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    return ""
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetSegmentPath() string {
    return "nssa"
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = nssa.Metric
    leafs["flags"] = nssa.Flags
    leafs["referenced-ls-type"] = nssa.ReferencedLsType
    leafs["prefix"] = nssa.Prefix
    leafs["prefix-options"] = nssa.PrefixOptions
    leafs["forwarding-address"] = nssa.ForwardingAddress
    leafs["external-route-tag"] = nssa.ExternalRouteTag
    leafs["referenced-link-state-id"] = nssa.ReferencedLinkStateId
    return leafs
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetBundleName() string { return "ietf" }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetYangName() string { return "nssa" }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) SetParent(parent types.Entity) { nssa.parent = parent }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetParent() types.Entity { return nssa.parent }

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link
// Link LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router Priority of the interface. The type is interface{} with range:
    // 0..255.
    RtrPriority interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetFilter() yfilter.YFilter { return link.YFilter }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) SetFilter(yf yfilter.YFilter) { link.YFilter = yf }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetGoName(yname string) string {
    if yname == "rtr-priority" { return "RtrPriority" }
    if yname == "options" { return "Options" }
    if yname == "link-local-interface-address" { return "LinkLocalInterfaceAddress" }
    if yname == "num-of-prefixes" { return "NumOfPrefixes" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetSegmentPath() string {
    return "link"
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list" {
        for _, c := range link.PrefixList {
            if link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList{}
        link.PrefixList = append(link.PrefixList, child)
        return &link.PrefixList[len(link.PrefixList)-1]
    }
    return nil
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range link.PrefixList {
        children[link.PrefixList[i].GetSegmentPath()] = &link.PrefixList[i]
    }
    return children
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rtr-priority"] = link.RtrPriority
    leafs["options"] = link.Options
    leafs["link-local-interface-address"] = link.LinkLocalInterfaceAddress
    leafs["num-of-prefixes"] = link.NumOfPrefixes
    return leafs
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetBundleName() string { return "ietf" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetYangName() string { return "link" }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) SetParent(parent types.Entity) { link.parent = parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetParent() types.Entity { return link.parent }

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetFilter() yfilter.YFilter { return prefixList.YFilter }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) SetFilter(yf yfilter.YFilter) { prefixList.YFilter = yf }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    return ""
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetSegmentPath() string {
    return "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = prefixList.Prefix
    leafs["prefix-options"] = prefixList.PrefixOptions
    return leafs
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetBundleName() string { return "ietf" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetYangName() string { return "prefix-list" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) SetParent(parent types.Entity) { prefixList.parent = parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetParent() types.Entity { return prefixList.parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetParentYangName() string { return "link" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix
// Intra-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetFilter() yfilter.YFilter { return intraAreaPrefix.YFilter }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) SetFilter(yf yfilter.YFilter) { intraAreaPrefix.YFilter = yf }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetGoName(yname string) string {
    if yname == "referenced-ls-type" { return "ReferencedLsType" }
    if yname == "referenced-link-state-id" { return "ReferencedLinkStateId" }
    if yname == "referenced-adv-router" { return "ReferencedAdvRouter" }
    if yname == "num-of-prefixes" { return "NumOfPrefixes" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetSegmentPath() string {
    return "intra-area-prefix"
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list" {
        for _, c := range intraAreaPrefix.PrefixList {
            if intraAreaPrefix.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList{}
        intraAreaPrefix.PrefixList = append(intraAreaPrefix.PrefixList, child)
        return &intraAreaPrefix.PrefixList[len(intraAreaPrefix.PrefixList)-1]
    }
    return nil
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intraAreaPrefix.PrefixList {
        children[intraAreaPrefix.PrefixList[i].GetSegmentPath()] = &intraAreaPrefix.PrefixList[i]
    }
    return children
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["referenced-ls-type"] = intraAreaPrefix.ReferencedLsType
    leafs["referenced-link-state-id"] = intraAreaPrefix.ReferencedLinkStateId
    leafs["referenced-adv-router"] = intraAreaPrefix.ReferencedAdvRouter
    leafs["num-of-prefixes"] = intraAreaPrefix.NumOfPrefixes
    return leafs
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetBundleName() string { return "ietf" }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetYangName() string { return "intra-area-prefix" }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) SetParent(parent types.Entity) { intraAreaPrefix.parent = parent }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetParent() types.Entity { return intraAreaPrefix.parent }

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetParentYangName() string { return "body" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetFilter() yfilter.YFilter { return prefixList.YFilter }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) SetFilter(yf yfilter.YFilter) { prefixList.YFilter = yf }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-options" { return "PrefixOptions" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetSegmentPath() string {
    return "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = prefixList.Prefix
    leafs["prefix-options"] = prefixList.PrefixOptions
    leafs["metric"] = prefixList.Metric
    return leafs
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetBundleName() string { return "ietf" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetYangName() string { return "prefix-list" }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) SetParent(parent types.Entity) { prefixList.parent = parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetParent() types.Entity { return prefixList.parent }

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetParentYangName() string { return "intra-area-prefix" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
// OSPF topology.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RIB. The type is string. Refers to
    // routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}

    // List of ospf areas. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area.
    Area []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "area" { return "Area" }
    return ""
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetSegmentPath() string {
    return "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "area" {
        for _, c := range topology.Area {
            if topology.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area{}
        topology.Area = append(topology.Area, child)
        return &topology.Area[len(topology.Area)-1]
    }
    return nil
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range topology.Area {
        children[topology.Area[i].GetSegmentPath()] = &topology.Area[i]
    }
    return children
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = topology.Name
    return leafs
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetBundleName() string { return "ietf" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetYangName() string { return "topology" }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetParent() types.Entity { return topology.parent }

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetParentYangName() string { return "instance" }

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
// List of ospf areas
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AreaId interface{}
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetFilter() yfilter.YFilter { return area.YFilter }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) SetFilter(yf yfilter.YFilter) { area.YFilter = yf }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetGoName(yname string) string {
    if yname == "area-id" { return "AreaId" }
    return ""
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetSegmentPath() string {
    return "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["area-id"] = area.AreaId
    return leafs
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetBundleName() string { return "ietf" }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetYangName() string { return "area" }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) SetParent(parent types.Entity) { area.parent = parent }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetParent() types.Entity { return area.parent }

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetParentYangName() string { return "topology" }

// RoutingState_RoutingInstance_Ribs
// Container for RIBs.
type RoutingState_RoutingInstance_Ribs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents a RIB identified by the 'name' key. All routes in a
    // RIB MUST belong to the same address family.  For each routing instance, an
    // implementation SHOULD provide one system-controlled default RIB for each
    // supported address family. The type is slice of
    // RoutingState_RoutingInstance_Ribs_Rib.
    Rib []RoutingState_RoutingInstance_Ribs_Rib
}

func (ribs *RoutingState_RoutingInstance_Ribs) GetFilter() yfilter.YFilter { return ribs.YFilter }

func (ribs *RoutingState_RoutingInstance_Ribs) SetFilter(yf yfilter.YFilter) { ribs.YFilter = yf }

func (ribs *RoutingState_RoutingInstance_Ribs) GetGoName(yname string) string {
    if yname == "rib" { return "Rib" }
    return ""
}

func (ribs *RoutingState_RoutingInstance_Ribs) GetSegmentPath() string {
    return "ribs"
}

func (ribs *RoutingState_RoutingInstance_Ribs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rib" {
        for _, c := range ribs.Rib {
            if ribs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_Ribs_Rib{}
        ribs.Rib = append(ribs.Rib, child)
        return &ribs.Rib[len(ribs.Rib)-1]
    }
    return nil
}

func (ribs *RoutingState_RoutingInstance_Ribs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ribs.Rib {
        children[ribs.Rib[i].GetSegmentPath()] = &ribs.Rib[i]
    }
    return children
}

func (ribs *RoutingState_RoutingInstance_Ribs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ribs *RoutingState_RoutingInstance_Ribs) GetBundleName() string { return "ietf" }

func (ribs *RoutingState_RoutingInstance_Ribs) GetYangName() string { return "ribs" }

func (ribs *RoutingState_RoutingInstance_Ribs) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ribs *RoutingState_RoutingInstance_Ribs) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ribs *RoutingState_RoutingInstance_Ribs) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ribs *RoutingState_RoutingInstance_Ribs) SetParent(parent types.Entity) { ribs.parent = parent }

func (ribs *RoutingState_RoutingInstance_Ribs) GetParent() types.Entity { return ribs.parent }

func (ribs *RoutingState_RoutingInstance_Ribs) GetParentYangName() string { return "routing-instance" }

// RoutingState_RoutingInstance_Ribs_Rib
// Each entry represents a RIB identified by the 'name'
// key. All routes in a RIB MUST belong to the same address
// family.
// 
// For each routing instance, an implementation SHOULD
// provide one system-controlled default RIB for each
// supported address family.
type RoutingState_RoutingInstance_Ribs_Rib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the RIB. The type is string.
    Name interface{}

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast. This attribute is mandatory.
    AddressFamily interface{}

    // This flag has the value of 'true' if and only if the RIB is the default RIB
    // for the given address family.  A default RIB always receives direct routes.
    // By default it also receives routes from all routing protocols. The type is
    // bool. The default value is true.
    DefaultRib interface{}

    // Current content of the RIB.
    Routes RoutingState_RoutingInstance_Ribs_Rib_Routes
}

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetFilter() yfilter.YFilter { return rib.YFilter }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) SetFilter(yf yfilter.YFilter) { rib.YFilter = yf }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "default-rib" { return "DefaultRib" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetSegmentPath() string {
    return "rib" + "[name='" + fmt.Sprintf("%v", rib.Name) + "']"
}

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &rib.Routes
    }
    return nil
}

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &rib.Routes
    return children
}

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = rib.Name
    leafs["address-family"] = rib.AddressFamily
    leafs["default-rib"] = rib.DefaultRib
    return leafs
}

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetBundleName() string { return "ietf" }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetYangName() string { return "rib" }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) SetParent(parent types.Entity) { rib.parent = parent }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetParent() types.Entity { return rib.parent }

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetParentYangName() string { return "ribs" }

// RoutingState_RoutingInstance_Ribs_Rib_Routes
// Current content of the RIB.
type RoutingState_RoutingInstance_Ribs_Rib_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A RIB route entry. This data node MUST be augmented with information
    // specific for routes of each address family. The type is slice of
    // RoutingState_RoutingInstance_Ribs_Rib_Routes_Route.
    Route []RoutingState_RoutingInstance_Ribs_Rib_Routes_Route
}

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingState_RoutingInstance_Ribs_Rib_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetBundleName() string { return "ietf" }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetYangName() string { return "routes" }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetParent() types.Entity { return routes.parent }

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetParentYangName() string { return "rib" }

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route
// A RIB route entry. This data node MUST be augmented
// with information specific for routes of each address
// family.
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Destination IP address with prefix. The type is
    // string.
    DestinationPrefix interface{}

    // This route attribute, also known as administrative distance, allows for
    // selecting the preferred route among routes with the same destination
    // prefix. A smaller value means a more preferred route. The type is
    // interface{} with range: 0..4294967295.
    RoutePreference interface{}

    // Route metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Type of the routing protocol from which the route originated. The type is
    // one of the following: Ospfv3Ospfv2OspfDirectStatic. This attribute is
    // mandatory.
    SourceProtocol interface{}

    // Presence of this leaf indicates that the route is preferred among all
    // routes in the same RIB that have the same destination prefix. The type is
    // interface{}.
    Active interface{}

    // Time stamp of the last modification of the route. If the route was never
    // modified, it is the time when the route was inserted into the RIB. The type
    // is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdated interface{}

    // Update source for the route. The type is string.
    UpdateSource interface{}

    // OSPF route tag. The type is interface{} with range: 0..4294967295. The
    // default value is 0.
    Tag interface{}

    // OSPF route type. The type is RouteType.
    RouteType interface{}

    // Route's next-hop attribute.
    NextHop RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop
}

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetGoName(yname string) string {
    if yname == "destination-prefix" { return "DestinationPrefix" }
    if yname == "route-preference" { return "RoutePreference" }
    if yname == "metric" { return "Metric" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "active" { return "Active" }
    if yname == "last-updated" { return "LastUpdated" }
    if yname == "update-source" { return "UpdateSource" }
    if yname == "tag" { return "Tag" }
    if yname == "route-type" { return "RouteType" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetSegmentPath() string {
    return "route" + "[destination-prefix='" + fmt.Sprintf("%v", route.DestinationPrefix) + "']"
}

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &route.NextHop
    }
    return nil
}

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &route.NextHop
    return children
}

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-prefix"] = route.DestinationPrefix
    leafs["route-preference"] = route.RoutePreference
    leafs["metric"] = route.Metric
    leafs["source-protocol"] = route.SourceProtocol
    leafs["active"] = route.Active
    leafs["last-updated"] = route.LastUpdated
    leafs["update-source"] = route.UpdateSource
    leafs["tag"] = route.Tag
    leafs["route-type"] = route.RouteType
    return leafs
}

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetBundleName() string { return "ietf" }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetYangName() string { return "route" }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetParentYangName() string { return "routes" }

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop
// Route's next-hop attribute.
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // IP address. The type is string.
    NextHopAddress interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}
}

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetGoName(yname string) string {
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "special-next-hop" { return "SpecialNextHop" }
    return ""
}

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outgoing-interface"] = nextHop.OutgoingInterface
    leafs["next-hop-address"] = nextHop.NextHopAddress
    leafs["special-next-hop"] = nextHop.SpecialNextHop
    return leafs
}

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetBundleName() string { return "ietf" }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetParentYangName() string { return "route" }

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop represents Special next-hop options.
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_blackhole RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_unreachable RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_prohibit RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_receive RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "receive"
)

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType represents OSPF route type
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType string

const (
    // OSPF intra-area route
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_intra_area RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "intra-area"

    // OSPF inter-area route
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_inter_area RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "inter-area"

    // OSPF external route type 1
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_external_1 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "external-1"

    // OSPF External route type 2
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_external_2 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "external-2"

    // OSPF NSSA external route type 1
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_nssa_1 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "nssa-1"

    // OSPF NSSA external route type 2
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_nssa_2 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "nssa-2"
)

// Routing
// Configuration parameters for the routing subsystem.
type Routing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of a routing instance. The type is slice of
    // Routing_RoutingInstance.
    RoutingInstance []Routing_RoutingInstance
}

func (routing *Routing) GetFilter() yfilter.YFilter { return routing.YFilter }

func (routing *Routing) SetFilter(yf yfilter.YFilter) { routing.YFilter = yf }

func (routing *Routing) GetGoName(yname string) string {
    if yname == "routing-instance" { return "RoutingInstance" }
    return ""
}

func (routing *Routing) GetSegmentPath() string {
    return "ietf-routing:routing"
}

func (routing *Routing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routing-instance" {
        for _, c := range routing.RoutingInstance {
            if routing.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance{}
        routing.RoutingInstance = append(routing.RoutingInstance, child)
        return &routing.RoutingInstance[len(routing.RoutingInstance)-1]
    }
    return nil
}

func (routing *Routing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routing.RoutingInstance {
        children[routing.RoutingInstance[i].GetSegmentPath()] = &routing.RoutingInstance[i]
    }
    return children
}

func (routing *Routing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routing *Routing) GetBundleName() string { return "ietf" }

func (routing *Routing) GetYangName() string { return "routing" }

func (routing *Routing) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routing *Routing) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routing *Routing) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routing *Routing) SetParent(parent types.Entity) { routing.parent = parent }

func (routing *Routing) GetParent() types.Entity { return routing.parent }

func (routing *Routing) GetParentYangName() string { return "ietf-routing" }

// Routing_RoutingInstance
// Configuration of a routing instance.
type Routing_RoutingInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the routing instance.  For
    // system-controlled entries, the value of this leaf must be the same as the
    // name of the corresponding entry in state data.  For user-controlled
    // entries, an arbitrary name can be used. The type is string.
    Name interface{}

    // The type of the routing instance. The type is one of the following:
    // VrfRoutingInstanceDefaultRoutingInstance. The default value is
    // rt:default-routing-instance.
    Type interface{}

    // Enable/disable the routing instance.  If this parameter is false, the
    // parent routing instance is disabled and does not appear in state data,
    // despite any other configuration that might be present. The type is bool.
    // The default value is true.
    Enabled interface{}

    // A 32-bit number in the form of a dotted quad that is used by some routing
    // protocols identifying a router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    RouterId interface{}

    // Textual description of the routing instance. The type is string.
    Description interface{}

    // Assignment of the routing instance's interfaces.
    Interfaces Routing_RoutingInstance_Interfaces

    // Configuration of routing protocol instances.
    RoutingProtocols Routing_RoutingInstance_RoutingProtocols

    // Configuration of RIBs.
    Ribs Routing_RoutingInstance_Ribs
}

func (routingInstance *Routing_RoutingInstance) GetFilter() yfilter.YFilter { return routingInstance.YFilter }

func (routingInstance *Routing_RoutingInstance) SetFilter(yf yfilter.YFilter) { routingInstance.YFilter = yf }

func (routingInstance *Routing_RoutingInstance) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "enabled" { return "Enabled" }
    if yname == "router-id" { return "RouterId" }
    if yname == "description" { return "Description" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "routing-protocols" { return "RoutingProtocols" }
    if yname == "ribs" { return "Ribs" }
    return ""
}

func (routingInstance *Routing_RoutingInstance) GetSegmentPath() string {
    return "routing-instance" + "[name='" + fmt.Sprintf("%v", routingInstance.Name) + "']"
}

func (routingInstance *Routing_RoutingInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &routingInstance.Interfaces
    }
    if childYangName == "routing-protocols" {
        return &routingInstance.RoutingProtocols
    }
    if childYangName == "ribs" {
        return &routingInstance.Ribs
    }
    return nil
}

func (routingInstance *Routing_RoutingInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &routingInstance.Interfaces
    children["routing-protocols"] = &routingInstance.RoutingProtocols
    children["ribs"] = &routingInstance.Ribs
    return children
}

func (routingInstance *Routing_RoutingInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = routingInstance.Name
    leafs["type"] = routingInstance.Type
    leafs["enabled"] = routingInstance.Enabled
    leafs["router-id"] = routingInstance.RouterId
    leafs["description"] = routingInstance.Description
    return leafs
}

func (routingInstance *Routing_RoutingInstance) GetBundleName() string { return "ietf" }

func (routingInstance *Routing_RoutingInstance) GetYangName() string { return "routing-instance" }

func (routingInstance *Routing_RoutingInstance) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routingInstance *Routing_RoutingInstance) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routingInstance *Routing_RoutingInstance) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routingInstance *Routing_RoutingInstance) SetParent(parent types.Entity) { routingInstance.parent = parent }

func (routingInstance *Routing_RoutingInstance) GetParent() types.Entity { return routingInstance.parent }

func (routingInstance *Routing_RoutingInstance) GetParentYangName() string { return "routing" }

// Routing_RoutingInstance_Interfaces
// Assignment of the routing instance's interfaces.
type Routing_RoutingInstance_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of a configured network layer interface to be assigned to the
    // routing-instance. The type is slice of string. Refers to
    // interfaces.Interfaces_Interface_Name
    Interface []interface{}
}

func (interfaces *Routing_RoutingInstance_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Routing_RoutingInstance_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Routing_RoutingInstance_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Routing_RoutingInstance_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Routing_RoutingInstance_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaces *Routing_RoutingInstance_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaces *Routing_RoutingInstance_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = interfaces.Interface
    return leafs
}

func (interfaces *Routing_RoutingInstance_Interfaces) GetBundleName() string { return "ietf" }

func (interfaces *Routing_RoutingInstance_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Routing_RoutingInstance_Interfaces) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (interfaces *Routing_RoutingInstance_Interfaces) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (interfaces *Routing_RoutingInstance_Interfaces) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (interfaces *Routing_RoutingInstance_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Routing_RoutingInstance_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Routing_RoutingInstance_Interfaces) GetParentYangName() string { return "routing-instance" }

// Routing_RoutingInstance_RoutingProtocols
// Configuration of routing protocol instances.
type Routing_RoutingInstance_RoutingProtocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains configuration of a routing protocol instance. The type
    // is slice of Routing_RoutingInstance_RoutingProtocols_RoutingProtocol.
    RoutingProtocol []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol
}

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetFilter() yfilter.YFilter { return routingProtocols.YFilter }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) SetFilter(yf yfilter.YFilter) { routingProtocols.YFilter = yf }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetGoName(yname string) string {
    if yname == "routing-protocol" { return "RoutingProtocol" }
    return ""
}

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetSegmentPath() string {
    return "routing-protocols"
}

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routing-protocol" {
        for _, c := range routingProtocols.RoutingProtocol {
            if routingProtocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol{}
        routingProtocols.RoutingProtocol = append(routingProtocols.RoutingProtocol, child)
        return &routingProtocols.RoutingProtocol[len(routingProtocols.RoutingProtocol)-1]
    }
    return nil
}

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routingProtocols.RoutingProtocol {
        children[routingProtocols.RoutingProtocol[i].GetSegmentPath()] = &routingProtocols.RoutingProtocol[i]
    }
    return children
}

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetBundleName() string { return "ietf" }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetYangName() string { return "routing-protocols" }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) SetParent(parent types.Entity) { routingProtocols.parent = parent }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetParent() types.Entity { return routingProtocols.parent }

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetParentYangName() string { return "routing-instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol
// Each entry contains configuration of a routing protocol
// instance.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type of the routing protocol - an identity derived
    // from the 'routing-protocol' base identity. The type is one of the
    // following: Ospfv3Ospfv2OspfDirectStatic.
    Type interface{}

    // This attribute is a key. An arbitrary name of the routing protocol
    // instance. The type is string.
    Name interface{}

    // Textual description of the routing protocol instance. The type is string.
    Description interface{}

    // Configuration of the 'static' pseudo-protocol.  Address-family-specific
    // modules augment this node with their lists of routes.
    StaticRoutes Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes

    // OSPF.
    Ospf Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
}

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetFilter() yfilter.YFilter { return routingProtocol.YFilter }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) SetFilter(yf yfilter.YFilter) { routingProtocol.YFilter = yf }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "static-routes" { return "StaticRoutes" }
    if yname == "ietf-ospf:ospf" { return "Ospf" }
    return ""
}

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetSegmentPath() string {
    return "routing-protocol" + "[type='" + fmt.Sprintf("%v", routingProtocol.Type) + "']" + "[name='" + fmt.Sprintf("%v", routingProtocol.Name) + "']"
}

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-routes" {
        return &routingProtocol.StaticRoutes
    }
    if childYangName == "ietf-ospf:ospf" {
        return &routingProtocol.Ospf
    }
    return nil
}

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["static-routes"] = &routingProtocol.StaticRoutes
    children["ietf-ospf:ospf"] = &routingProtocol.Ospf
    return children
}

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = routingProtocol.Type
    leafs["name"] = routingProtocol.Name
    leafs["description"] = routingProtocol.Description
    return leafs
}

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetBundleName() string { return "ietf" }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetYangName() string { return "routing-protocol" }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) SetParent(parent types.Entity) { routingProtocol.parent = parent }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetParent() types.Entity { return routingProtocol.parent }

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetParentYangName() string { return "routing-protocols" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes
// Configuration of the 'static' pseudo-protocol.
// 
// Address-family-specific modules augment this node with
// their lists of routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of a 'static' pseudo-protocol instance consists of a list of
    // routes.
    Ipv4 Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4

    // Configuration of a 'static' pseudo-protocol instance consists of a list of
    // routes.
    Ipv6 Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6
}

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetFilter() yfilter.YFilter { return staticRoutes.YFilter }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) SetFilter(yf yfilter.YFilter) { staticRoutes.YFilter = yf }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetGoName(yname string) string {
    if yname == "ietf-ipv4-unicast-routing:ipv4" { return "Ipv4" }
    if yname == "ietf-ipv6-unicast-routing:ipv6" { return "Ipv6" }
    return ""
}

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetSegmentPath() string {
    return "static-routes"
}

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ietf-ipv4-unicast-routing:ipv4" {
        return &staticRoutes.Ipv4
    }
    if childYangName == "ietf-ipv6-unicast-routing:ipv6" {
        return &staticRoutes.Ipv6
    }
    return nil
}

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ietf-ipv4-unicast-routing:ipv4"] = &staticRoutes.Ipv4
    children["ietf-ipv6-unicast-routing:ipv6"] = &staticRoutes.Ipv6
    return children
}

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetBundleName() string { return "ietf" }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetYangName() string { return "static-routes" }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) SetParent(parent types.Entity) { staticRoutes.parent = parent }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetParent() types.Entity { return staticRoutes.parent }

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetParentYangName() string { return "routing-protocol" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4
// Configuration of a 'static' pseudo-protocol instance
// consists of a list of routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A user-ordered list of static routes. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route.
    Route []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route
}

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetSegmentPath() string {
    return "ietf-ipv4-unicast-routing:ipv4"
}

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range ipv4.Route {
            if ipv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route{}
        ipv4.Route = append(ipv4.Route, child)
        return &ipv4.Route[len(ipv4.Route)-1]
    }
    return nil
}

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4.Route {
        children[ipv4.Route[i].GetSegmentPath()] = &ipv4.Route[i]
    }
    return children
}

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetBundleName() string { return "ietf" }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetParentYangName() string { return "static-routes" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route
// A user-ordered list of static routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 destination prefix. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    // This attribute is mandatory.
    DestinationPrefix interface{}

    // Textual description of the route. The type is string.
    Description interface{}

    // Configuration of next-hop.
    NextHop Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetGoName(yname string) string {
    if yname == "destination-prefix" { return "DestinationPrefix" }
    if yname == "description" { return "Description" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetSegmentPath() string {
    return "route" + "[destination-prefix='" + fmt.Sprintf("%v", route.DestinationPrefix) + "']"
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &route.NextHop
    }
    return nil
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &route.NextHop
    return children
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-prefix"] = route.DestinationPrefix
    leafs["description"] = route.Description
    return leafs
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetBundleName() string { return "ietf" }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetYangName() string { return "route" }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetParent() types.Entity { return route.parent }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetParentYangName() string { return "ipv4" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop
// Configuration of next-hop.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}

    // IPv4 address of the next-hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetGoName(yname string) string {
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    if yname == "special-next-hop" { return "SpecialNextHop" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    return ""
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outgoing-interface"] = nextHop.OutgoingInterface
    leafs["special-next-hop"] = nextHop.SpecialNextHop
    leafs["next-hop-address"] = nextHop.NextHopAddress
    return leafs
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetBundleName() string { return "ietf" }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetParentYangName() string { return "route" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop represents Special next-hop options.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_blackhole Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_unreachable Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_prohibit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_receive Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "receive"
)

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6
// Configuration of a 'static' pseudo-protocol instance
// consists of a list of routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A user-ordered list of static routes. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route.
    Route []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route
}

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetSegmentPath() string {
    return "ietf-ipv6-unicast-routing:ipv6"
}

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range ipv6.Route {
            if ipv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route{}
        ipv6.Route = append(ipv6.Route, child)
        return &ipv6.Route[len(ipv6.Route)-1]
    }
    return nil
}

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6.Route {
        children[ipv6.Route[i].GetSegmentPath()] = &ipv6.Route[i]
    }
    return children
}

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetBundleName() string { return "ietf" }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetParentYangName() string { return "static-routes" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route
// A user-ordered list of static routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 destination prefix. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    // This attribute is mandatory.
    DestinationPrefix interface{}

    // Textual description of the route. The type is string.
    Description interface{}

    // Configuration of next-hop.
    NextHop Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetGoName(yname string) string {
    if yname == "destination-prefix" { return "DestinationPrefix" }
    if yname == "description" { return "Description" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetSegmentPath() string {
    return "route" + "[destination-prefix='" + fmt.Sprintf("%v", route.DestinationPrefix) + "']"
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &route.NextHop
    }
    return nil
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &route.NextHop
    return children
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-prefix"] = route.DestinationPrefix
    leafs["description"] = route.Description
    return leafs
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetBundleName() string { return "ietf" }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetYangName() string { return "route" }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetParent() types.Entity { return route.parent }

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetParentYangName() string { return "ipv6" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop
// Configuration of next-hop.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}

    // IPv6 address of the next-hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetGoName(yname string) string {
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    if yname == "special-next-hop" { return "SpecialNextHop" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    return ""
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outgoing-interface"] = nextHop.OutgoingInterface
    leafs["special-next-hop"] = nextHop.SpecialNextHop
    leafs["next-hop-address"] = nextHop.NextHopAddress
    return leafs
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetBundleName() string { return "ietf" }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetParentYangName() string { return "route" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop represents Special next-hop options.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_blackhole Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_unreachable Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_prohibit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_receive Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "receive"
)

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
// OSPF.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF operation mode. The type is one of the following: ShipsInTheNight. The
    // default value is ospf:ships-in-the-night.
    OperationMode interface{}

    // Inheritance support to all instances.
    AllInstancesInherit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit

    // An OSPF routing protocol instance. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance.
    Instance []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
}

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetGoName(yname string) string {
    if yname == "operation-mode" { return "OperationMode" }
    if yname == "all-instances-inherit" { return "AllInstancesInherit" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetSegmentPath() string {
    return "ietf-ospf:ospf"
}

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all-instances-inherit" {
        return &ospf.AllInstancesInherit
    }
    if childYangName == "instance" {
        for _, c := range ospf.Instance {
            if ospf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance{}
        ospf.Instance = append(ospf.Instance, child)
        return &ospf.Instance[len(ospf.Instance)-1]
    }
    return nil
}

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["all-instances-inherit"] = &ospf.AllInstancesInherit
    for i := range ospf.Instance {
        children[ospf.Instance[i].GetSegmentPath()] = &ospf.Instance[i]
    }
    return children
}

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-mode"] = ospf.OperationMode
    return leafs
}

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetBundleName() string { return "ietf" }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetYangName() string { return "ospf" }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetParentYangName() string { return "routing-protocol" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit
// Inheritance support to all instances.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Area config to be inherited by all areas in all instances.
    Area Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area

    // Interface config to be inherited by all interfaces in all instances.
    Interface Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface
}

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetFilter() yfilter.YFilter { return allInstancesInherit.YFilter }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) SetFilter(yf yfilter.YFilter) { allInstancesInherit.YFilter = yf }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetGoName(yname string) string {
    if yname == "area" { return "Area" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetSegmentPath() string {
    return "all-instances-inherit"
}

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "area" {
        return &allInstancesInherit.Area
    }
    if childYangName == "interface" {
        return &allInstancesInherit.Interface
    }
    return nil
}

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["area"] = &allInstancesInherit.Area
    children["interface"] = &allInstancesInherit.Interface
    return children
}

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetBundleName() string { return "ietf" }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetYangName() string { return "all-instances-inherit" }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) SetParent(parent types.Entity) { allInstancesInherit.parent = parent }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetParent() types.Entity { return allInstancesInherit.parent }

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetParentYangName() string { return "ospf" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area
// Area config to be inherited by all areas in
// all instances.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetFilter() yfilter.YFilter { return area.YFilter }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) SetFilter(yf yfilter.YFilter) { area.YFilter = yf }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetGoName(yname string) string {
    return ""
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetSegmentPath() string {
    return "area"
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetBundleName() string { return "ietf" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetYangName() string { return "area" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) SetParent(parent types.Entity) { area.parent = parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetParent() types.Entity { return area.parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetParentYangName() string { return "all-instances-inherit" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface
// Interface config to be inherited by all interfaces
// in all instances.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetGoName(yname string) string {
    return ""
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetBundleName() string { return "ietf" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetYangName() string { return "interface" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetParent() types.Entity { return self.parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetParentYangName() string { return "all-instances-inherit" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
// An OSPF routing protocol instance.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address-family of the instance. The type is one of
    // the following: Ipv4Ipv4UnicastIpv6Ipv6Unicast.
    Af interface{}

    // Defined in RFC 2328. A 32-bit number that uniquely identifies the router.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    RouterId interface{}

    // Enable/Disable the protocol. The type is bool. The default value is true.
    Enable interface{}

    // Admin distance config state.
    AdminDistance Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance

    // NSR config state.
    Nsr Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr

    // Graceful restart config state.
    GracefulRestart Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart

    // Auto cost config state.
    AutoCost Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost

    // SPF calculation control.
    SpfControl Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl

    // Database maintenance control.
    DatabaseControl Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl

    // Protocol reload control.
    ReloadControl Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl

    // OSPF MPLS config state.
    Mpls Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls

    // This container may be augmented with global parameters for IPFRR.
    FastReroute Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute

    // Inheritance for all areas.
    AllAreasInherit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit

    // List of ospf areas. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area.
    Area []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area

    // OSPF topology. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology.
    Topology []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
}

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "router-id" { return "RouterId" }
    if yname == "enable" { return "Enable" }
    if yname == "admin-distance" { return "AdminDistance" }
    if yname == "nsr" { return "Nsr" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "auto-cost" { return "AutoCost" }
    if yname == "spf-control" { return "SpfControl" }
    if yname == "database-control" { return "DatabaseControl" }
    if yname == "reload-control" { return "ReloadControl" }
    if yname == "mpls" { return "Mpls" }
    if yname == "fast-reroute" { return "FastReroute" }
    if yname == "all-areas-inherit" { return "AllAreasInherit" }
    if yname == "area" { return "Area" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetSegmentPath() string {
    return "instance" + "[af='" + fmt.Sprintf("%v", instance.Af) + "']"
}

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "admin-distance" {
        return &instance.AdminDistance
    }
    if childYangName == "nsr" {
        return &instance.Nsr
    }
    if childYangName == "graceful-restart" {
        return &instance.GracefulRestart
    }
    if childYangName == "auto-cost" {
        return &instance.AutoCost
    }
    if childYangName == "spf-control" {
        return &instance.SpfControl
    }
    if childYangName == "database-control" {
        return &instance.DatabaseControl
    }
    if childYangName == "reload-control" {
        return &instance.ReloadControl
    }
    if childYangName == "mpls" {
        return &instance.Mpls
    }
    if childYangName == "fast-reroute" {
        return &instance.FastReroute
    }
    if childYangName == "all-areas-inherit" {
        return &instance.AllAreasInherit
    }
    if childYangName == "area" {
        for _, c := range instance.Area {
            if instance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area{}
        instance.Area = append(instance.Area, child)
        return &instance.Area[len(instance.Area)-1]
    }
    if childYangName == "topology" {
        for _, c := range instance.Topology {
            if instance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology{}
        instance.Topology = append(instance.Topology, child)
        return &instance.Topology[len(instance.Topology)-1]
    }
    return nil
}

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["admin-distance"] = &instance.AdminDistance
    children["nsr"] = &instance.Nsr
    children["graceful-restart"] = &instance.GracefulRestart
    children["auto-cost"] = &instance.AutoCost
    children["spf-control"] = &instance.SpfControl
    children["database-control"] = &instance.DatabaseControl
    children["reload-control"] = &instance.ReloadControl
    children["mpls"] = &instance.Mpls
    children["fast-reroute"] = &instance.FastReroute
    children["all-areas-inherit"] = &instance.AllAreasInherit
    for i := range instance.Area {
        children[instance.Area[i].GetSegmentPath()] = &instance.Area[i]
    }
    for i := range instance.Topology {
        children[instance.Topology[i].GetSegmentPath()] = &instance.Topology[i]
    }
    return children
}

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = instance.Af
    leafs["router-id"] = instance.RouterId
    leafs["enable"] = instance.Enable
    return leafs
}

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetBundleName() string { return "ietf" }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetYangName() string { return "instance" }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetParent() types.Entity { return instance.parent }

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetParentYangName() string { return "ospf" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance
// Admin distance config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Admin distance for intra-area route. The type is interface{} with range:
    // 0..255.
    IntraArea interface{}

    // Admin distance for inter-area route. The type is interface{} with range:
    // 0..255.
    InterArea interface{}

    // Admin distance for both intra-area and inter-area route. The type is
    // interface{} with range: 0..255.
    Internal interface{}

    // Admin distance for both external route. The type is interface{} with range:
    // 0..255.
    External interface{}
}

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetFilter() yfilter.YFilter { return adminDistance.YFilter }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) SetFilter(yf yfilter.YFilter) { adminDistance.YFilter = yf }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetGoName(yname string) string {
    if yname == "intra-area" { return "IntraArea" }
    if yname == "inter-area" { return "InterArea" }
    if yname == "internal" { return "Internal" }
    if yname == "external" { return "External" }
    return ""
}

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetSegmentPath() string {
    return "admin-distance"
}

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intra-area"] = adminDistance.IntraArea
    leafs["inter-area"] = adminDistance.InterArea
    leafs["internal"] = adminDistance.Internal
    leafs["external"] = adminDistance.External
    return leafs
}

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetBundleName() string { return "ietf" }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetYangName() string { return "admin-distance" }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) SetParent(parent types.Entity) { adminDistance.parent = parent }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetParent() types.Entity { return adminDistance.parent }

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr
// NSR config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable NSR. The type is bool.
    Enable interface{}
}

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetFilter() yfilter.YFilter { return nsr.YFilter }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) SetFilter(yf yfilter.YFilter) { nsr.YFilter = yf }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetSegmentPath() string {
    return "nsr"
}

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = nsr.Enable
    return leafs
}

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetBundleName() string { return "ietf" }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetYangName() string { return "nsr" }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) SetParent(parent types.Entity) { nsr.parent = parent }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetParent() types.Entity { return nsr.parent }

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart
// Graceful restart config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable graceful restart as defined in RFC 3623. The type is bool.
    Enable interface{}

    // Enable RestartHelperSupport in RFC 3623 Section B.2. The type is bool.
    HelperEnable interface{}

    // RestartInterval option in RFC 3623 Section B.1. The type is interface{}
    // with range: 1..1800. Units are seconds. The default value is 120.
    RestartInterval interface{}

    // RestartHelperStrictLSAChecking option in RFC 3623 Section B.2. The type is
    // bool.
    HelperStrictLsaChecking interface{}
}

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "helper-enable" { return "HelperEnable" }
    if yname == "restart-interval" { return "RestartInterval" }
    if yname == "helper-strict-lsa-checking" { return "HelperStrictLsaChecking" }
    return ""
}

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = gracefulRestart.Enable
    leafs["helper-enable"] = gracefulRestart.HelperEnable
    leafs["restart-interval"] = gracefulRestart.RestartInterval
    leafs["helper-strict-lsa-checking"] = gracefulRestart.HelperStrictLsaChecking
    return leafs
}

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetBundleName() string { return "ietf" }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost
// Auto cost config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable auto cost. The type is bool.
    Enable interface{}

    // Configure reference bandwidth in term of Mbits. The type is interface{}
    // with range: 1..4294967. Units are Mbits.
    ReferenceBandwidth interface{}
}

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetFilter() yfilter.YFilter { return autoCost.YFilter }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) SetFilter(yf yfilter.YFilter) { autoCost.YFilter = yf }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "reference-bandwidth" { return "ReferenceBandwidth" }
    return ""
}

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetSegmentPath() string {
    return "auto-cost"
}

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = autoCost.Enable
    leafs["reference-bandwidth"] = autoCost.ReferenceBandwidth
    return leafs
}

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetBundleName() string { return "ietf" }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetYangName() string { return "auto-cost" }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) SetParent(parent types.Entity) { autoCost.parent = parent }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetParent() types.Entity { return autoCost.parent }

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl
// SPF calculation control.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of ECMP paths. The type is interface{} with range: 1..32.
    Paths interface{}
}

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetFilter() yfilter.YFilter { return spfControl.YFilter }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) SetFilter(yf yfilter.YFilter) { spfControl.YFilter = yf }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetGoName(yname string) string {
    if yname == "paths" { return "Paths" }
    return ""
}

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetSegmentPath() string {
    return "spf-control"
}

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["paths"] = spfControl.Paths
    return leafs
}

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetBundleName() string { return "ietf" }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetYangName() string { return "spf-control" }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) SetParent(parent types.Entity) { spfControl.parent = parent }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetParent() types.Entity { return spfControl.parent }

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl
// Database maintenance control.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of LSAs OSPF will receive. The type is interface{} with
    // range: 1..4294967294.
    MaxLsa interface{}
}

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetFilter() yfilter.YFilter { return databaseControl.YFilter }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) SetFilter(yf yfilter.YFilter) { databaseControl.YFilter = yf }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetGoName(yname string) string {
    if yname == "max-lsa" { return "MaxLsa" }
    return ""
}

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetSegmentPath() string {
    return "database-control"
}

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-lsa"] = databaseControl.MaxLsa
    return leafs
}

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetBundleName() string { return "ietf" }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetYangName() string { return "database-control" }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) SetParent(parent types.Entity) { databaseControl.parent = parent }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetParent() types.Entity { return databaseControl.parent }

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl
// Protocol reload control.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetFilter() yfilter.YFilter { return reloadControl.YFilter }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) SetFilter(yf yfilter.YFilter) { reloadControl.YFilter = yf }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetGoName(yname string) string {
    return ""
}

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetSegmentPath() string {
    return "reload-control"
}

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetBundleName() string { return "ietf" }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetYangName() string { return "reload-control" }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) SetParent(parent types.Entity) { reloadControl.parent = parent }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetParent() types.Entity { return reloadControl.parent }

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls
// OSPF MPLS config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Traffic Engineering stable IP address for system.
    TeRid Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid

    // OSPF MPLS LDP config state.
    Ldp Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp
}

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetGoName(yname string) string {
    if yname == "te-rid" { return "TeRid" }
    if yname == "ldp" { return "Ldp" }
    return ""
}

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "te-rid" {
        return &mpls.TeRid
    }
    if childYangName == "ldp" {
        return &mpls.Ldp
    }
    return nil
}

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["te-rid"] = &mpls.TeRid
    children["ldp"] = &mpls.Ldp
    return children
}

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetBundleName() string { return "ietf" }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetYangName() string { return "mpls" }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid
// Traffic Engineering stable IP address for system.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Take the interface's IPv4 address as TE router ID. The type is string.
    // Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Explicitly configure the TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}
}

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetFilter() yfilter.YFilter { return teRid.YFilter }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) SetFilter(yf yfilter.YFilter) { teRid.YFilter = yf }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "router-id" { return "RouterId" }
    return ""
}

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetSegmentPath() string {
    return "te-rid"
}

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = teRid.Interface
    leafs["router-id"] = teRid.RouterId
    return leafs
}

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetBundleName() string { return "ietf" }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetYangName() string { return "te-rid" }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) SetParent(parent types.Entity) { teRid.parent = parent }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetParent() types.Entity { return teRid.parent }

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetParentYangName() string { return "mpls" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp
// OSPF MPLS LDP config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable LDP IGP synchronization. The type is bool.
    IgpSync interface{}

    // Enable LDP IGP interface auto-configuration. The type is bool.
    Autoconfig interface{}
}

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetFilter() yfilter.YFilter { return ldp.YFilter }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) SetFilter(yf yfilter.YFilter) { ldp.YFilter = yf }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetGoName(yname string) string {
    if yname == "igp-sync" { return "IgpSync" }
    if yname == "autoconfig" { return "Autoconfig" }
    return ""
}

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetSegmentPath() string {
    return "ldp"
}

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-sync"] = ldp.IgpSync
    leafs["autoconfig"] = ldp.Autoconfig
    return leafs
}

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetBundleName() string { return "ietf" }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetYangName() string { return "ldp" }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) SetParent(parent types.Entity) { ldp.parent = parent }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetParent() types.Entity { return ldp.parent }

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetParentYangName() string { return "mpls" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute
// This container may be augmented with global
// parameters for IPFRR.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This container may be augmented with global parameters for LFA. Creating
    // the container has no effect on LFA activation.
    Lfa Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetFilter() yfilter.YFilter { return fastReroute.YFilter }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) SetFilter(yf yfilter.YFilter) { fastReroute.YFilter = yf }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetGoName(yname string) string {
    if yname == "lfa" { return "Lfa" }
    return ""
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetSegmentPath() string {
    return "fast-reroute"
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lfa" {
        return &fastReroute.Lfa
    }
    return nil
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lfa"] = &fastReroute.Lfa
    return children
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetBundleName() string { return "ietf" }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetYangName() string { return "fast-reroute" }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) SetParent(parent types.Entity) { fastReroute.parent = parent }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetParent() types.Entity { return fastReroute.parent }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa
// This container may be augmented with
// global parameters for LFA.
// Creating the container has no effect on
// LFA activation.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetFilter() yfilter.YFilter { return lfa.YFilter }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) SetFilter(yf yfilter.YFilter) { lfa.YFilter = yf }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetGoName(yname string) string {
    return ""
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetSegmentPath() string {
    return "lfa"
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetBundleName() string { return "ietf" }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetYangName() string { return "lfa" }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) SetParent(parent types.Entity) { lfa.parent = parent }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetParent() types.Entity { return lfa.parent }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetParentYangName() string { return "fast-reroute" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit
// Inheritance for all areas.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Area config to be inherited by all areas.
    Area Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area

    // Interface config to be inherited by all interfaces in all areas.
    Interface Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface
}

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetFilter() yfilter.YFilter { return allAreasInherit.YFilter }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) SetFilter(yf yfilter.YFilter) { allAreasInherit.YFilter = yf }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetGoName(yname string) string {
    if yname == "area" { return "Area" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetSegmentPath() string {
    return "all-areas-inherit"
}

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "area" {
        return &allAreasInherit.Area
    }
    if childYangName == "interface" {
        return &allAreasInherit.Interface
    }
    return nil
}

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["area"] = &allAreasInherit.Area
    children["interface"] = &allAreasInherit.Interface
    return children
}

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetBundleName() string { return "ietf" }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetYangName() string { return "all-areas-inherit" }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) SetParent(parent types.Entity) { allAreasInherit.parent = parent }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetParent() types.Entity { return allAreasInherit.parent }

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area
// Area config to be inherited by all areas.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetFilter() yfilter.YFilter { return area.YFilter }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) SetFilter(yf yfilter.YFilter) { area.YFilter = yf }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetGoName(yname string) string {
    return ""
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetSegmentPath() string {
    return "area"
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetBundleName() string { return "ietf" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetYangName() string { return "area" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) SetParent(parent types.Entity) { area.parent = parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetParent() types.Entity { return area.parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetParentYangName() string { return "all-areas-inherit" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface
// Interface config to be inherited by all interfaces
// in all areas.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetGoName(yname string) string {
    return ""
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetBundleName() string { return "ietf" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetYangName() string { return "interface" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetParent() types.Entity { return self.parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetParentYangName() string { return "all-areas-inherit" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area
// List of ospf areas
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AreaId interface{}

    // Area type. The type is one of the following: NormalNssaStub. The default
    // value is normal.
    AreaType interface{}

    // Enable/Disable summary generation to the stub or NSSA area. The type is
    // bool.
    Summary interface{}

    // Set the summary default-cost for a stub or NSSA area. The type is
    // interface{} with range: 1..16777215.
    DefaultCost interface{}

    // Summarize routes matching address/mask (border routers only). The type is
    // slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range.
    Range []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range

    // Inheritance for all interfaces.
    AllInterfacesInherit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit

    // OSPF virtual link. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink.
    VirtualLink []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink

    // OSPF sham link. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink.
    ShamLink []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink

    // List of OSPF interfaces. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface.
    Interface []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetFilter() yfilter.YFilter { return area.YFilter }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) SetFilter(yf yfilter.YFilter) { area.YFilter = yf }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetGoName(yname string) string {
    if yname == "area-id" { return "AreaId" }
    if yname == "area-type" { return "AreaType" }
    if yname == "summary" { return "Summary" }
    if yname == "default-cost" { return "DefaultCost" }
    if yname == "range" { return "Range" }
    if yname == "all-interfaces-inherit" { return "AllInterfacesInherit" }
    if yname == "virtual-link" { return "VirtualLink" }
    if yname == "sham-link" { return "ShamLink" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetSegmentPath() string {
    return "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "range" {
        for _, c := range area.Range {
            if area.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range{}
        area.Range = append(area.Range, child)
        return &area.Range[len(area.Range)-1]
    }
    if childYangName == "all-interfaces-inherit" {
        return &area.AllInterfacesInherit
    }
    if childYangName == "virtual-link" {
        for _, c := range area.VirtualLink {
            if area.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink{}
        area.VirtualLink = append(area.VirtualLink, child)
        return &area.VirtualLink[len(area.VirtualLink)-1]
    }
    if childYangName == "sham-link" {
        for _, c := range area.ShamLink {
            if area.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink{}
        area.ShamLink = append(area.ShamLink, child)
        return &area.ShamLink[len(area.ShamLink)-1]
    }
    if childYangName == "interface" {
        for _, c := range area.Interface {
            if area.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface{}
        area.Interface = append(area.Interface, child)
        return &area.Interface[len(area.Interface)-1]
    }
    return nil
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range area.Range {
        children[area.Range[i].GetSegmentPath()] = &area.Range[i]
    }
    children["all-interfaces-inherit"] = &area.AllInterfacesInherit
    for i := range area.VirtualLink {
        children[area.VirtualLink[i].GetSegmentPath()] = &area.VirtualLink[i]
    }
    for i := range area.ShamLink {
        children[area.ShamLink[i].GetSegmentPath()] = &area.ShamLink[i]
    }
    for i := range area.Interface {
        children[area.Interface[i].GetSegmentPath()] = &area.Interface[i]
    }
    return children
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["area-id"] = area.AreaId
    leafs["area-type"] = area.AreaType
    leafs["summary"] = area.Summary
    leafs["default-cost"] = area.DefaultCost
    return leafs
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetBundleName() string { return "ietf" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetYangName() string { return "area" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) SetParent(parent types.Entity) { area.parent = parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetParent() types.Entity { return area.parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range
// Summarize routes matching address/mask (border
// routers only)
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 or IPv6 prefix. The type is string.
    Prefix interface{}

    // Advertise or hide. The type is bool.
    Advertise interface{}

    // Cost of summary route. The type is interface{} with range: 0..16777214.
    Cost interface{}
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "advertise" { return "Advertise" }
    if yname == "cost" { return "Cost" }
    return ""
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetSegmentPath() string {
    return "range" + "[prefix='" + fmt.Sprintf("%v", self.Prefix) + "']"
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = self.Prefix
    leafs["advertise"] = self.Advertise
    leafs["cost"] = self.Cost
    return leafs
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetBundleName() string { return "ietf" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetYangName() string { return "range" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetParent() types.Entity { return self.parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetParentYangName() string { return "area" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit
// Inheritance for all interfaces
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface config to be inherited by all interfaces.
    Interface Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface
}

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetFilter() yfilter.YFilter { return allInterfacesInherit.YFilter }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) SetFilter(yf yfilter.YFilter) { allInterfacesInherit.YFilter = yf }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetSegmentPath() string {
    return "all-interfaces-inherit"
}

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &allInterfacesInherit.Interface
    }
    return nil
}

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &allInterfacesInherit.Interface
    return children
}

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetBundleName() string { return "ietf" }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetYangName() string { return "all-interfaces-inherit" }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) SetParent(parent types.Entity) { allInterfacesInherit.parent = parent }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetParent() types.Entity { return allInterfacesInherit.parent }

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetParentYangName() string { return "area" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface
// Interface config to be inherited by all
// interfaces.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetGoName(yname string) string {
    return ""
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetBundleName() string { return "ietf" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetYangName() string { return "interface" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetParent() types.Entity { return self.parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetParentYangName() string { return "all-interfaces-inherit" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink
// OSPF virtual link
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Virtual link router ID. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    RouterId interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // TTL security check.
    TtlSecurity Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity

    // Authentication configuration.
    Authentication Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication
}

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetFilter() yfilter.YFilter { return virtualLink.YFilter }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) SetFilter(yf yfilter.YFilter) { virtualLink.YFilter = yf }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "cost" { return "Cost" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "dead-interval" { return "DeadInterval" }
    if yname == "retransmit-interval" { return "RetransmitInterval" }
    if yname == "transmit-delay" { return "TransmitDelay" }
    if yname == "mtu-ignore" { return "MtuIgnore" }
    if yname == "lls" { return "Lls" }
    if yname == "prefix-suppression" { return "PrefixSuppression" }
    if yname == "bfd" { return "Bfd" }
    if yname == "enable" { return "Enable" }
    if yname == "ttl-security" { return "TtlSecurity" }
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetSegmentPath() string {
    return "virtual-link" + "[router-id='" + fmt.Sprintf("%v", virtualLink.RouterId) + "']"
}

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ttl-security" {
        return &virtualLink.TtlSecurity
    }
    if childYangName == "authentication" {
        return &virtualLink.Authentication
    }
    return nil
}

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ttl-security"] = &virtualLink.TtlSecurity
    children["authentication"] = &virtualLink.Authentication
    return children
}

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = virtualLink.RouterId
    leafs["cost"] = virtualLink.Cost
    leafs["hello-interval"] = virtualLink.HelloInterval
    leafs["dead-interval"] = virtualLink.DeadInterval
    leafs["retransmit-interval"] = virtualLink.RetransmitInterval
    leafs["transmit-delay"] = virtualLink.TransmitDelay
    leafs["mtu-ignore"] = virtualLink.MtuIgnore
    leafs["lls"] = virtualLink.Lls
    leafs["prefix-suppression"] = virtualLink.PrefixSuppression
    leafs["bfd"] = virtualLink.Bfd
    leafs["enable"] = virtualLink.Enable
    return leafs
}

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetBundleName() string { return "ietf" }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetYangName() string { return "virtual-link" }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) SetParent(parent types.Entity) { virtualLink.parent = parent }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetParent() types.Entity { return virtualLink.parent }

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetParentYangName() string { return "area" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity
// TTL security check.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetFilter() yfilter.YFilter { return ttlSecurity.YFilter }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) SetFilter(yf yfilter.YFilter) { ttlSecurity.YFilter = yf }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetSegmentPath() string {
    return "ttl-security"
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ttlSecurity.Enable
    leafs["hops"] = ttlSecurity.Hops
    return leafs
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetBundleName() string { return "ietf" }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetYangName() string { return "ttl-security" }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) SetParent(parent types.Entity) { ttlSecurity.parent = parent }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetParent() types.Entity { return ttlSecurity.parent }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetParentYangName() string { return "virtual-link" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication
// Authentication configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetGoName(yname string) string {
    if yname == "sa" { return "Sa" }
    if yname == "key-chain" { return "KeyChain" }
    if yname == "key" { return "Key" }
    if yname == "crypto-algorithm" { return "CryptoAlgorithm" }
    return ""
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crypto-algorithm" {
        return &authentication.CryptoAlgorithm
    }
    return nil
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["crypto-algorithm"] = &authentication.CryptoAlgorithm
    return children
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa"] = authentication.Sa
    leafs["key-chain"] = authentication.KeyChain
    leafs["key"] = authentication.Key
    return leafs
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetBundleName() string { return "ietf" }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetYangName() string { return "authentication" }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetParentYangName() string { return "virtual-link" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetFilter() yfilter.YFilter { return cryptoAlgorithm.YFilter }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) SetFilter(yf yfilter.YFilter) { cryptoAlgorithm.YFilter = yf }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetGoName(yname string) string {
    if yname == "hmac-sha1-12" { return "HmacSha112" }
    if yname == "hmac-sha1-20" { return "HmacSha120" }
    if yname == "md5" { return "Md5" }
    if yname == "sha-1" { return "Sha1" }
    if yname == "hmac-sha-1" { return "HmacSha1" }
    if yname == "hmac-sha-256" { return "HmacSha256" }
    if yname == "hmac-sha-384" { return "HmacSha384" }
    if yname == "hmac-sha-512" { return "HmacSha512" }
    return ""
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetSegmentPath() string {
    return "crypto-algorithm"
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hmac-sha1-12"] = cryptoAlgorithm.HmacSha112
    leafs["hmac-sha1-20"] = cryptoAlgorithm.HmacSha120
    leafs["md5"] = cryptoAlgorithm.Md5
    leafs["sha-1"] = cryptoAlgorithm.Sha1
    leafs["hmac-sha-1"] = cryptoAlgorithm.HmacSha1
    leafs["hmac-sha-256"] = cryptoAlgorithm.HmacSha256
    leafs["hmac-sha-384"] = cryptoAlgorithm.HmacSha384
    leafs["hmac-sha-512"] = cryptoAlgorithm.HmacSha512
    return leafs
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetBundleName() string { return "ietf" }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetYangName() string { return "crypto-algorithm" }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) SetParent(parent types.Entity) { cryptoAlgorithm.parent = parent }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetParent() types.Entity { return cryptoAlgorithm.parent }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetParentYangName() string { return "authentication" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink
// OSPF sham link
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address of the local end-point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalId interface{}

    // This attribute is a key. Address of the remote end-point. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteId interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // TTL security check.
    TtlSecurity Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity

    // Authentication configuration.
    Authentication Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication
}

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetFilter() yfilter.YFilter { return shamLink.YFilter }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) SetFilter(yf yfilter.YFilter) { shamLink.YFilter = yf }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetGoName(yname string) string {
    if yname == "local-id" { return "LocalId" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "cost" { return "Cost" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "dead-interval" { return "DeadInterval" }
    if yname == "retransmit-interval" { return "RetransmitInterval" }
    if yname == "transmit-delay" { return "TransmitDelay" }
    if yname == "mtu-ignore" { return "MtuIgnore" }
    if yname == "lls" { return "Lls" }
    if yname == "prefix-suppression" { return "PrefixSuppression" }
    if yname == "bfd" { return "Bfd" }
    if yname == "enable" { return "Enable" }
    if yname == "ttl-security" { return "TtlSecurity" }
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetSegmentPath() string {
    return "sham-link" + "[local-id='" + fmt.Sprintf("%v", shamLink.LocalId) + "']" + "[remote-id='" + fmt.Sprintf("%v", shamLink.RemoteId) + "']"
}

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ttl-security" {
        return &shamLink.TtlSecurity
    }
    if childYangName == "authentication" {
        return &shamLink.Authentication
    }
    return nil
}

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ttl-security"] = &shamLink.TtlSecurity
    children["authentication"] = &shamLink.Authentication
    return children
}

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-id"] = shamLink.LocalId
    leafs["remote-id"] = shamLink.RemoteId
    leafs["cost"] = shamLink.Cost
    leafs["hello-interval"] = shamLink.HelloInterval
    leafs["dead-interval"] = shamLink.DeadInterval
    leafs["retransmit-interval"] = shamLink.RetransmitInterval
    leafs["transmit-delay"] = shamLink.TransmitDelay
    leafs["mtu-ignore"] = shamLink.MtuIgnore
    leafs["lls"] = shamLink.Lls
    leafs["prefix-suppression"] = shamLink.PrefixSuppression
    leafs["bfd"] = shamLink.Bfd
    leafs["enable"] = shamLink.Enable
    return leafs
}

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetBundleName() string { return "ietf" }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetYangName() string { return "sham-link" }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) SetParent(parent types.Entity) { shamLink.parent = parent }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetParent() types.Entity { return shamLink.parent }

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetParentYangName() string { return "area" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity
// TTL security check.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetFilter() yfilter.YFilter { return ttlSecurity.YFilter }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) SetFilter(yf yfilter.YFilter) { ttlSecurity.YFilter = yf }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetSegmentPath() string {
    return "ttl-security"
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ttlSecurity.Enable
    leafs["hops"] = ttlSecurity.Hops
    return leafs
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetBundleName() string { return "ietf" }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetYangName() string { return "ttl-security" }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) SetParent(parent types.Entity) { ttlSecurity.parent = parent }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetParent() types.Entity { return ttlSecurity.parent }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetParentYangName() string { return "sham-link" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication
// Authentication configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetGoName(yname string) string {
    if yname == "sa" { return "Sa" }
    if yname == "key-chain" { return "KeyChain" }
    if yname == "key" { return "Key" }
    if yname == "crypto-algorithm" { return "CryptoAlgorithm" }
    return ""
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crypto-algorithm" {
        return &authentication.CryptoAlgorithm
    }
    return nil
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["crypto-algorithm"] = &authentication.CryptoAlgorithm
    return children
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa"] = authentication.Sa
    leafs["key-chain"] = authentication.KeyChain
    leafs["key"] = authentication.Key
    return leafs
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetBundleName() string { return "ietf" }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetYangName() string { return "authentication" }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetParentYangName() string { return "sham-link" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetFilter() yfilter.YFilter { return cryptoAlgorithm.YFilter }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) SetFilter(yf yfilter.YFilter) { cryptoAlgorithm.YFilter = yf }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetGoName(yname string) string {
    if yname == "hmac-sha1-12" { return "HmacSha112" }
    if yname == "hmac-sha1-20" { return "HmacSha120" }
    if yname == "md5" { return "Md5" }
    if yname == "sha-1" { return "Sha1" }
    if yname == "hmac-sha-1" { return "HmacSha1" }
    if yname == "hmac-sha-256" { return "HmacSha256" }
    if yname == "hmac-sha-384" { return "HmacSha384" }
    if yname == "hmac-sha-512" { return "HmacSha512" }
    return ""
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetSegmentPath() string {
    return "crypto-algorithm"
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hmac-sha1-12"] = cryptoAlgorithm.HmacSha112
    leafs["hmac-sha1-20"] = cryptoAlgorithm.HmacSha120
    leafs["md5"] = cryptoAlgorithm.Md5
    leafs["sha-1"] = cryptoAlgorithm.Sha1
    leafs["hmac-sha-1"] = cryptoAlgorithm.HmacSha1
    leafs["hmac-sha-256"] = cryptoAlgorithm.HmacSha256
    leafs["hmac-sha-384"] = cryptoAlgorithm.HmacSha384
    leafs["hmac-sha-512"] = cryptoAlgorithm.HmacSha512
    return leafs
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetBundleName() string { return "ietf" }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetYangName() string { return "crypto-algorithm" }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) SetParent(parent types.Entity) { cryptoAlgorithm.parent = parent }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetParent() types.Entity { return cryptoAlgorithm.parent }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetParentYangName() string { return "authentication" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface
// List of OSPF interfaces.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Network type. The type is NetworkType.
    NetworkType interface{}

    // Enable/Disable passive. The type is bool.
    Passive interface{}

    // Enable/Disable demand circuit. The type is bool.
    DemandCircuit interface{}

    // Set prefix as a node representative prefix. The type is bool. The default
    // value is false.
    NodeFlag interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // Configure ospf multi-area.
    MultiArea Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea

    // Static configured neighbors.
    StaticNeighbors Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors

    // Fast-reroute configuration.
    FastReroute Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute

    // TTL security check.
    TtlSecurity Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity

    // Authentication configuration.
    Authentication Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication

    // OSPF interface topology. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology.
    Topology []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "network-type" { return "NetworkType" }
    if yname == "passive" { return "Passive" }
    if yname == "demand-circuit" { return "DemandCircuit" }
    if yname == "node-flag" { return "NodeFlag" }
    if yname == "cost" { return "Cost" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "dead-interval" { return "DeadInterval" }
    if yname == "retransmit-interval" { return "RetransmitInterval" }
    if yname == "transmit-delay" { return "TransmitDelay" }
    if yname == "mtu-ignore" { return "MtuIgnore" }
    if yname == "lls" { return "Lls" }
    if yname == "prefix-suppression" { return "PrefixSuppression" }
    if yname == "bfd" { return "Bfd" }
    if yname == "enable" { return "Enable" }
    if yname == "multi-area" { return "MultiArea" }
    if yname == "static-neighbors" { return "StaticNeighbors" }
    if yname == "fast-reroute" { return "FastReroute" }
    if yname == "ttl-security" { return "TtlSecurity" }
    if yname == "authentication" { return "Authentication" }
    if yname == "topology" { return "Topology" }
    return ""
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multi-area" {
        return &self.MultiArea
    }
    if childYangName == "static-neighbors" {
        return &self.StaticNeighbors
    }
    if childYangName == "fast-reroute" {
        return &self.FastReroute
    }
    if childYangName == "ttl-security" {
        return &self.TtlSecurity
    }
    if childYangName == "authentication" {
        return &self.Authentication
    }
    if childYangName == "topology" {
        for _, c := range self.Topology {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology{}
        self.Topology = append(self.Topology, child)
        return &self.Topology[len(self.Topology)-1]
    }
    return nil
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["multi-area"] = &self.MultiArea
    children["static-neighbors"] = &self.StaticNeighbors
    children["fast-reroute"] = &self.FastReroute
    children["ttl-security"] = &self.TtlSecurity
    children["authentication"] = &self.Authentication
    for i := range self.Topology {
        children[self.Topology[i].GetSegmentPath()] = &self.Topology[i]
    }
    return children
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["network-type"] = self.NetworkType
    leafs["passive"] = self.Passive
    leafs["demand-circuit"] = self.DemandCircuit
    leafs["node-flag"] = self.NodeFlag
    leafs["cost"] = self.Cost
    leafs["hello-interval"] = self.HelloInterval
    leafs["dead-interval"] = self.DeadInterval
    leafs["retransmit-interval"] = self.RetransmitInterval
    leafs["transmit-delay"] = self.TransmitDelay
    leafs["mtu-ignore"] = self.MtuIgnore
    leafs["lls"] = self.Lls
    leafs["prefix-suppression"] = self.PrefixSuppression
    leafs["bfd"] = self.Bfd
    leafs["enable"] = self.Enable
    return leafs
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetBundleName() string { return "ietf" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetYangName() string { return "interface" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetParent() types.Entity { return self.parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetParentYangName() string { return "area" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea
// Configure ospf multi-area.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Multi-area ID. The type is one of the following types: int with range:
    // 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    MultiAreaId interface{}

    // Interface cost for multi-area. The type is interface{} with range:
    // 0..65535.
    Cost interface{}
}

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetFilter() yfilter.YFilter { return multiArea.YFilter }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) SetFilter(yf yfilter.YFilter) { multiArea.YFilter = yf }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetGoName(yname string) string {
    if yname == "multi-area-id" { return "MultiAreaId" }
    if yname == "cost" { return "Cost" }
    return ""
}

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetSegmentPath() string {
    return "multi-area"
}

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multi-area-id"] = multiArea.MultiAreaId
    leafs["cost"] = multiArea.Cost
    return leafs
}

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetBundleName() string { return "ietf" }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetYangName() string { return "multi-area" }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) SetParent(parent types.Entity) { multiArea.parent = parent }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetParent() types.Entity { return multiArea.parent }

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetParentYangName() string { return "interface" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors
// Static configured neighbors.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify a neighbor router. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor.
    Neighbor []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor
}

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetFilter() yfilter.YFilter { return staticNeighbors.YFilter }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) SetFilter(yf yfilter.YFilter) { staticNeighbors.YFilter = yf }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetSegmentPath() string {
    return "static-neighbors"
}

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range staticNeighbors.Neighbor {
            if staticNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor{}
        staticNeighbors.Neighbor = append(staticNeighbors.Neighbor, child)
        return &staticNeighbors.Neighbor[len(staticNeighbors.Neighbor)-1]
    }
    return nil
}

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticNeighbors.Neighbor {
        children[staticNeighbors.Neighbor[i].GetSegmentPath()] = &staticNeighbors.Neighbor[i]
    }
    return children
}

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetBundleName() string { return "ietf" }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetYangName() string { return "static-neighbors" }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) SetParent(parent types.Entity) { staticNeighbors.parent = parent }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetParent() types.Entity { return staticNeighbors.parent }

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetParentYangName() string { return "interface" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor
// Specify a neighbor router.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Neighbor cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Neighbor poll interval. The type is interface{} with range: 1..65535. Units
    // are seconds.
    PollInterval interface{}

    // Neighbor priority for DR election. The type is interface{} with range:
    // 1..255.
    Priority interface{}
}

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "cost" { return "Cost" }
    if yname == "poll-interval" { return "PollInterval" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[address='" + fmt.Sprintf("%v", neighbor.Address) + "']"
}

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = neighbor.Address
    leafs["cost"] = neighbor.Cost
    leafs["poll-interval"] = neighbor.PollInterval
    leafs["priority"] = neighbor.Priority
    return leafs
}

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetBundleName() string { return "ietf" }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetParentYangName() string { return "static-neighbors" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute
// Fast-reroute configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LFA configuration.
    Lfa Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetFilter() yfilter.YFilter { return fastReroute.YFilter }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) SetFilter(yf yfilter.YFilter) { fastReroute.YFilter = yf }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetGoName(yname string) string {
    if yname == "lfa" { return "Lfa" }
    return ""
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetSegmentPath() string {
    return "fast-reroute"
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lfa" {
        return &fastReroute.Lfa
    }
    return nil
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lfa"] = &fastReroute.Lfa
    return children
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetBundleName() string { return "ietf" }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetYangName() string { return "fast-reroute" }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) SetParent(parent types.Entity) { fastReroute.parent = parent }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetParent() types.Entity { return fastReroute.parent }

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetParentYangName() string { return "interface" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa
// LFA configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prevent the interface to be used as backup. The type is bool.
    CandidateDisabled interface{}

    // Activates LFA. This model assumes activation of per-prefix LFA. The type is
    // bool.
    Enabled interface{}

    // Remote LFA configuration.
    RemoteLfa Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetFilter() yfilter.YFilter { return lfa.YFilter }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) SetFilter(yf yfilter.YFilter) { lfa.YFilter = yf }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetGoName(yname string) string {
    if yname == "candidate-disabled" { return "CandidateDisabled" }
    if yname == "enabled" { return "Enabled" }
    if yname == "remote-lfa" { return "RemoteLfa" }
    return ""
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetSegmentPath() string {
    return "lfa"
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-lfa" {
        return &lfa.RemoteLfa
    }
    return nil
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-lfa"] = &lfa.RemoteLfa
    return children
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate-disabled"] = lfa.CandidateDisabled
    leafs["enabled"] = lfa.Enabled
    return leafs
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetBundleName() string { return "ietf" }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetYangName() string { return "lfa" }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) SetParent(parent types.Entity) { lfa.parent = parent }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetParent() types.Entity { return lfa.parent }

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetParentYangName() string { return "fast-reroute" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa
// Remote LFA configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Activates remote LFA. The type is bool.
    Enabled interface{}
}

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetFilter() yfilter.YFilter { return remoteLfa.YFilter }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) SetFilter(yf yfilter.YFilter) { remoteLfa.YFilter = yf }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetSegmentPath() string {
    return "remote-lfa"
}

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = remoteLfa.Enabled
    return leafs
}

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetBundleName() string { return "ietf" }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetYangName() string { return "remote-lfa" }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) SetParent(parent types.Entity) { remoteLfa.parent = parent }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetParent() types.Entity { return remoteLfa.parent }

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetParentYangName() string { return "lfa" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity
// TTL security check.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetFilter() yfilter.YFilter { return ttlSecurity.YFilter }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) SetFilter(yf yfilter.YFilter) { ttlSecurity.YFilter = yf }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetSegmentPath() string {
    return "ttl-security"
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ttlSecurity.Enable
    leafs["hops"] = ttlSecurity.Hops
    return leafs
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetBundleName() string { return "ietf" }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetYangName() string { return "ttl-security" }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) SetParent(parent types.Entity) { ttlSecurity.parent = parent }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetParent() types.Entity { return ttlSecurity.parent }

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetParentYangName() string { return "interface" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication
// Authentication configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetGoName(yname string) string {
    if yname == "sa" { return "Sa" }
    if yname == "key-chain" { return "KeyChain" }
    if yname == "key" { return "Key" }
    if yname == "crypto-algorithm" { return "CryptoAlgorithm" }
    return ""
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crypto-algorithm" {
        return &authentication.CryptoAlgorithm
    }
    return nil
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["crypto-algorithm"] = &authentication.CryptoAlgorithm
    return children
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa"] = authentication.Sa
    leafs["key-chain"] = authentication.KeyChain
    leafs["key"] = authentication.Key
    return leafs
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetBundleName() string { return "ietf" }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetYangName() string { return "authentication" }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetParentYangName() string { return "interface" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetFilter() yfilter.YFilter { return cryptoAlgorithm.YFilter }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) SetFilter(yf yfilter.YFilter) { cryptoAlgorithm.YFilter = yf }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetGoName(yname string) string {
    if yname == "hmac-sha1-12" { return "HmacSha112" }
    if yname == "hmac-sha1-20" { return "HmacSha120" }
    if yname == "md5" { return "Md5" }
    if yname == "sha-1" { return "Sha1" }
    if yname == "hmac-sha-1" { return "HmacSha1" }
    if yname == "hmac-sha-256" { return "HmacSha256" }
    if yname == "hmac-sha-384" { return "HmacSha384" }
    if yname == "hmac-sha-512" { return "HmacSha512" }
    return ""
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetSegmentPath() string {
    return "crypto-algorithm"
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hmac-sha1-12"] = cryptoAlgorithm.HmacSha112
    leafs["hmac-sha1-20"] = cryptoAlgorithm.HmacSha120
    leafs["md5"] = cryptoAlgorithm.Md5
    leafs["sha-1"] = cryptoAlgorithm.Sha1
    leafs["hmac-sha-1"] = cryptoAlgorithm.HmacSha1
    leafs["hmac-sha-256"] = cryptoAlgorithm.HmacSha256
    leafs["hmac-sha-384"] = cryptoAlgorithm.HmacSha384
    leafs["hmac-sha-512"] = cryptoAlgorithm.HmacSha512
    return leafs
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetBundleName() string { return "ietf" }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetYangName() string { return "crypto-algorithm" }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) SetParent(parent types.Entity) { cryptoAlgorithm.parent = parent }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetParent() types.Entity { return cryptoAlgorithm.parent }

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetParentYangName() string { return "authentication" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology
// OSPF interface topology.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string. Refers to routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}

    // Interface cost for this topology. The type is interface{} with range:
    // 0..4294967295.
    Cost interface{}
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "cost" { return "Cost" }
    return ""
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetSegmentPath() string {
    return "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = topology.Name
    leafs["cost"] = topology.Cost
    return leafs
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetBundleName() string { return "ietf" }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetYangName() string { return "topology" }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetParent() types.Entity { return topology.parent }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetParentYangName() string { return "interface" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType represents Network type.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType string

const (
    // Specify OSPF broadcast multi-access network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_broadcast Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "broadcast"

    // Specify OSPF Non-Broadcast Multi-Access
    // (NBMA) network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_non_broadcast Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "non-broadcast"

    // Specify OSPF point-to-multipoint network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_point_to_multipoint Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "point-to-multipoint"

    // Specify OSPF point-to-point network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_point_to_point Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "point-to-point"
)

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
// OSPF topology.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RIB. The type is string. Refers to
    // routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}

    // List of ospf areas. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area.
    Area []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetFilter() yfilter.YFilter { return topology.YFilter }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) SetFilter(yf yfilter.YFilter) { topology.YFilter = yf }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "area" { return "Area" }
    return ""
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetSegmentPath() string {
    return "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "area" {
        for _, c := range topology.Area {
            if topology.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area{}
        topology.Area = append(topology.Area, child)
        return &topology.Area[len(topology.Area)-1]
    }
    return nil
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range topology.Area {
        children[topology.Area[i].GetSegmentPath()] = &topology.Area[i]
    }
    return children
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = topology.Name
    return leafs
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetBundleName() string { return "ietf" }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetYangName() string { return "topology" }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) SetParent(parent types.Entity) { topology.parent = parent }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetParent() types.Entity { return topology.parent }

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetParentYangName() string { return "instance" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
// List of ospf areas
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    AreaId interface{}

    // Area type. The type is one of the following: NormalNssaStub. The default
    // value is normal.
    AreaType interface{}

    // Enable/Disable summary generation to the stub or NSSA area. The type is
    // bool.
    Summary interface{}

    // Set the summary default-cost for a stub or NSSA area. The type is
    // interface{} with range: 1..16777215.
    DefaultCost interface{}

    // Summarize routes matching address/mask (border routers only). The type is
    // slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range.
    Range []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetFilter() yfilter.YFilter { return area.YFilter }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) SetFilter(yf yfilter.YFilter) { area.YFilter = yf }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetGoName(yname string) string {
    if yname == "area-id" { return "AreaId" }
    if yname == "area-type" { return "AreaType" }
    if yname == "summary" { return "Summary" }
    if yname == "default-cost" { return "DefaultCost" }
    if yname == "range" { return "Range" }
    return ""
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetSegmentPath() string {
    return "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "range" {
        for _, c := range area.Range {
            if area.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range{}
        area.Range = append(area.Range, child)
        return &area.Range[len(area.Range)-1]
    }
    return nil
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range area.Range {
        children[area.Range[i].GetSegmentPath()] = &area.Range[i]
    }
    return children
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["area-id"] = area.AreaId
    leafs["area-type"] = area.AreaType
    leafs["summary"] = area.Summary
    leafs["default-cost"] = area.DefaultCost
    return leafs
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetBundleName() string { return "ietf" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetYangName() string { return "area" }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) SetParent(parent types.Entity) { area.parent = parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetParent() types.Entity { return area.parent }

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetParentYangName() string { return "topology" }

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range
// Summarize routes matching address/mask (border
// routers only)
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 or IPv6 prefix. The type is string.
    Prefix interface{}

    // Advertise or hide. The type is bool.
    Advertise interface{}

    // Cost of summary route. The type is interface{} with range: 0..16777214.
    Cost interface{}
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "advertise" { return "Advertise" }
    if yname == "cost" { return "Cost" }
    return ""
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetSegmentPath() string {
    return "range" + "[prefix='" + fmt.Sprintf("%v", self.Prefix) + "']"
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = self.Prefix
    leafs["advertise"] = self.Advertise
    leafs["cost"] = self.Cost
    return leafs
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetBundleName() string { return "ietf" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetYangName() string { return "range" }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetParent() types.Entity { return self.parent }

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetParentYangName() string { return "area" }

// Routing_RoutingInstance_Ribs
// Configuration of RIBs.
type Routing_RoutingInstance_Ribs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains configuration for a RIB identified by the 'name' key. 
    // Entries having the same key as a system-controlled entry of the list
    // /routing-state/routing-instance/ribs/rib are used for configuring
    // parameters of that entry. Other entries define additional user-controlled
    // RIBs. The type is slice of Routing_RoutingInstance_Ribs_Rib.
    Rib []Routing_RoutingInstance_Ribs_Rib
}

func (ribs *Routing_RoutingInstance_Ribs) GetFilter() yfilter.YFilter { return ribs.YFilter }

func (ribs *Routing_RoutingInstance_Ribs) SetFilter(yf yfilter.YFilter) { ribs.YFilter = yf }

func (ribs *Routing_RoutingInstance_Ribs) GetGoName(yname string) string {
    if yname == "rib" { return "Rib" }
    return ""
}

func (ribs *Routing_RoutingInstance_Ribs) GetSegmentPath() string {
    return "ribs"
}

func (ribs *Routing_RoutingInstance_Ribs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rib" {
        for _, c := range ribs.Rib {
            if ribs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Routing_RoutingInstance_Ribs_Rib{}
        ribs.Rib = append(ribs.Rib, child)
        return &ribs.Rib[len(ribs.Rib)-1]
    }
    return nil
}

func (ribs *Routing_RoutingInstance_Ribs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ribs.Rib {
        children[ribs.Rib[i].GetSegmentPath()] = &ribs.Rib[i]
    }
    return children
}

func (ribs *Routing_RoutingInstance_Ribs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ribs *Routing_RoutingInstance_Ribs) GetBundleName() string { return "ietf" }

func (ribs *Routing_RoutingInstance_Ribs) GetYangName() string { return "ribs" }

func (ribs *Routing_RoutingInstance_Ribs) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (ribs *Routing_RoutingInstance_Ribs) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (ribs *Routing_RoutingInstance_Ribs) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (ribs *Routing_RoutingInstance_Ribs) SetParent(parent types.Entity) { ribs.parent = parent }

func (ribs *Routing_RoutingInstance_Ribs) GetParent() types.Entity { return ribs.parent }

func (ribs *Routing_RoutingInstance_Ribs) GetParentYangName() string { return "routing-instance" }

// Routing_RoutingInstance_Ribs_Rib
// Each entry contains configuration for a RIB identified
// by the 'name' key.
// 
// Entries having the same key as a system-controlled entry
// of the list /routing-state/routing-instance/ribs/rib are
// used for configuring parameters of that entry. Other
// entries define additional user-controlled RIBs.
type Routing_RoutingInstance_Ribs_Rib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the RIB.  For system-controlled
    // entries, the value of this leaf must be the same as the name of the
    // corresponding entry in state data.  For user-controlled entries, an
    // arbitrary name can be used. The type is string.
    Name interface{}

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast.
    AddressFamily interface{}

    // Textual description of the RIB. The type is string.
    Description interface{}
}

func (rib *Routing_RoutingInstance_Ribs_Rib) GetFilter() yfilter.YFilter { return rib.YFilter }

func (rib *Routing_RoutingInstance_Ribs_Rib) SetFilter(yf yfilter.YFilter) { rib.YFilter = yf }

func (rib *Routing_RoutingInstance_Ribs_Rib) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "description" { return "Description" }
    return ""
}

func (rib *Routing_RoutingInstance_Ribs_Rib) GetSegmentPath() string {
    return "rib" + "[name='" + fmt.Sprintf("%v", rib.Name) + "']"
}

func (rib *Routing_RoutingInstance_Ribs_Rib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rib *Routing_RoutingInstance_Ribs_Rib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rib *Routing_RoutingInstance_Ribs_Rib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = rib.Name
    leafs["address-family"] = rib.AddressFamily
    leafs["description"] = rib.Description
    return leafs
}

func (rib *Routing_RoutingInstance_Ribs_Rib) GetBundleName() string { return "ietf" }

func (rib *Routing_RoutingInstance_Ribs_Rib) GetYangName() string { return "rib" }

func (rib *Routing_RoutingInstance_Ribs_Rib) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (rib *Routing_RoutingInstance_Ribs_Rib) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (rib *Routing_RoutingInstance_Ribs_Rib) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (rib *Routing_RoutingInstance_Ribs_Rib) SetParent(parent types.Entity) { rib.parent = parent }

func (rib *Routing_RoutingInstance_Ribs_Rib) GetParent() types.Entity { return rib.parent }

func (rib *Routing_RoutingInstance_Ribs_Rib) GetParentYangName() string { return "ribs" }

// FibRoute
// Return the active FIB route that a routing-instance uses for
// sending packets to a destination address.
type FibRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input FibRoute_Input

    
    Output FibRoute_Output
}

func (fibRoute *FibRoute) GetFilter() yfilter.YFilter { return fibRoute.YFilter }

func (fibRoute *FibRoute) SetFilter(yf yfilter.YFilter) { fibRoute.YFilter = yf }

func (fibRoute *FibRoute) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (fibRoute *FibRoute) GetSegmentPath() string {
    return "ietf-routing:fib-route"
}

func (fibRoute *FibRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &fibRoute.Input
    }
    if childYangName == "output" {
        return &fibRoute.Output
    }
    return nil
}

func (fibRoute *FibRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &fibRoute.Input
    children["output"] = &fibRoute.Output
    return children
}

func (fibRoute *FibRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fibRoute *FibRoute) GetBundleName() string { return "ietf" }

func (fibRoute *FibRoute) GetYangName() string { return "fib-route" }

func (fibRoute *FibRoute) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (fibRoute *FibRoute) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (fibRoute *FibRoute) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (fibRoute *FibRoute) SetParent(parent types.Entity) { fibRoute.parent = parent }

func (fibRoute *FibRoute) GetParent() types.Entity { return fibRoute.parent }

func (fibRoute *FibRoute) GetParentYangName() string { return "ietf-routing" }

// FibRoute_Input
type FibRoute_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the routing instance whose forwarding information base is being
    // queried.  If the routing instance with name equal to the value of this
    // parameter doesn't exist, then this operation SHALL fail with error-tag
    // 'data-missing' and error-app-tag 'routing-instance-not-found'. The type is
    // string. This attribute is mandatory.
    RoutingInstanceName interface{}

    // Network layer destination address.  Address family specific modules MUST
    // augment this container with a leaf named 'address'.
    DestinationAddress FibRoute_Input_DestinationAddress
}

func (input *FibRoute_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *FibRoute_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *FibRoute_Input) GetGoName(yname string) string {
    if yname == "routing-instance-name" { return "RoutingInstanceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    return ""
}

func (input *FibRoute_Input) GetSegmentPath() string {
    return "input"
}

func (input *FibRoute_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-address" {
        return &input.DestinationAddress
    }
    return nil
}

func (input *FibRoute_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination-address"] = &input.DestinationAddress
    return children
}

func (input *FibRoute_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["routing-instance-name"] = input.RoutingInstanceName
    return leafs
}

func (input *FibRoute_Input) GetBundleName() string { return "ietf" }

func (input *FibRoute_Input) GetYangName() string { return "input" }

func (input *FibRoute_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *FibRoute_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *FibRoute_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *FibRoute_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *FibRoute_Input) GetParent() types.Entity { return input.parent }

func (input *FibRoute_Input) GetParentYangName() string { return "fib-route" }

// FibRoute_Input_DestinationAddress
// Network layer destination address.
// 
// Address family specific modules MUST augment this
// container with a leaf named 'address'.
type FibRoute_Input_DestinationAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast. This attribute is mandatory.
    AddressFamily interface{}

    // IPv4 destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IetfIpv4UnicastRoutingAddress interface{}

    // IPv6 destination address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IetfIpv6UnicastRoutingAddress interface{}
}

func (destinationAddress *FibRoute_Input_DestinationAddress) GetFilter() yfilter.YFilter { return destinationAddress.YFilter }

func (destinationAddress *FibRoute_Input_DestinationAddress) SetFilter(yf yfilter.YFilter) { destinationAddress.YFilter = yf }

func (destinationAddress *FibRoute_Input_DestinationAddress) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "address" { return "IetfIpv4UnicastRoutingAddress" }
    if yname == "address" { return "IetfIpv6UnicastRoutingAddress" }
    return ""
}

func (destinationAddress *FibRoute_Input_DestinationAddress) GetSegmentPath() string {
    return "destination-address"
}

func (destinationAddress *FibRoute_Input_DestinationAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationAddress *FibRoute_Input_DestinationAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationAddress *FibRoute_Input_DestinationAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = destinationAddress.AddressFamily
    leafs["address"] = destinationAddress.IetfIpv4UnicastRoutingAddress
    leafs["address"] = destinationAddress.IetfIpv6UnicastRoutingAddress
    return leafs
}

func (destinationAddress *FibRoute_Input_DestinationAddress) GetBundleName() string { return "ietf" }

func (destinationAddress *FibRoute_Input_DestinationAddress) GetYangName() string { return "destination-address" }

func (destinationAddress *FibRoute_Input_DestinationAddress) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (destinationAddress *FibRoute_Input_DestinationAddress) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (destinationAddress *FibRoute_Input_DestinationAddress) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (destinationAddress *FibRoute_Input_DestinationAddress) SetParent(parent types.Entity) { destinationAddress.parent = parent }

func (destinationAddress *FibRoute_Input_DestinationAddress) GetParent() types.Entity { return destinationAddress.parent }

func (destinationAddress *FibRoute_Input_DestinationAddress) GetParentYangName() string { return "input" }

// FibRoute_Output
type FibRoute_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The active FIB route for the specified destination.  If the routing
    // instance has no active FIB route for the destination address, no output is
    // returned - the server SHALL send an <rpc-reply> containing a single element
    // <ok>.  Address family specific modules MUST augment this list with
    // appropriate route contents.
    Route FibRoute_Output_Route
}

func (output *FibRoute_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *FibRoute_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *FibRoute_Output) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (output *FibRoute_Output) GetSegmentPath() string {
    return "output"
}

func (output *FibRoute_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        return &output.Route
    }
    return nil
}

func (output *FibRoute_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route"] = &output.Route
    return children
}

func (output *FibRoute_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (output *FibRoute_Output) GetBundleName() string { return "ietf" }

func (output *FibRoute_Output) GetYangName() string { return "output" }

func (output *FibRoute_Output) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (output *FibRoute_Output) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (output *FibRoute_Output) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (output *FibRoute_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *FibRoute_Output) GetParent() types.Entity { return output.parent }

func (output *FibRoute_Output) GetParentYangName() string { return "fib-route" }

// FibRoute_Output_Route
// The active FIB route for the specified destination.
// 
// If the routing instance has no active FIB route for the
// destination address, no output is returned - the server
// SHALL send an <rpc-reply> containing a single element
// <ok>.
// 
// Address family specific modules MUST augment this list
// with appropriate route contents.
type FibRoute_Output_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast. This attribute is mandatory.
    AddressFamily interface{}

    // Type of the routing protocol from which the route originated. The type is
    // one of the following: Ospfv3Ospfv2OspfDirectStatic. This attribute is
    // mandatory.
    SourceProtocol interface{}

    // Presence of this leaf indicates that the route is preferred among all
    // routes in the same RIB that have the same destination prefix. The type is
    // interface{}.
    Active interface{}

    // Time stamp of the last modification of the route. If the route was never
    // modified, it is the time when the route was inserted into the RIB. The type
    // is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdated interface{}

    // IPv4 destination prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    IetfIpv4UnicastRoutingDestinationPrefix interface{}

    // IPv6 destination prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    IetfIpv6UnicastRoutingDestinationPrefix interface{}

    // Route's next-hop attribute.
    NextHop FibRoute_Output_Route_NextHop
}

func (route *FibRoute_Output_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *FibRoute_Output_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *FibRoute_Output_Route) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "active" { return "Active" }
    if yname == "last-updated" { return "LastUpdated" }
    if yname == "destination-prefix" { return "IetfIpv4UnicastRoutingDestinationPrefix" }
    if yname == "destination-prefix" { return "IetfIpv6UnicastRoutingDestinationPrefix" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (route *FibRoute_Output_Route) GetSegmentPath() string {
    return "route"
}

func (route *FibRoute_Output_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &route.NextHop
    }
    return nil
}

func (route *FibRoute_Output_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &route.NextHop
    return children
}

func (route *FibRoute_Output_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = route.AddressFamily
    leafs["source-protocol"] = route.SourceProtocol
    leafs["active"] = route.Active
    leafs["last-updated"] = route.LastUpdated
    leafs["destination-prefix"] = route.IetfIpv4UnicastRoutingDestinationPrefix
    leafs["destination-prefix"] = route.IetfIpv6UnicastRoutingDestinationPrefix
    return leafs
}

func (route *FibRoute_Output_Route) GetBundleName() string { return "ietf" }

func (route *FibRoute_Output_Route) GetYangName() string { return "route" }

func (route *FibRoute_Output_Route) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (route *FibRoute_Output_Route) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (route *FibRoute_Output_Route) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (route *FibRoute_Output_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *FibRoute_Output_Route) GetParent() types.Entity { return route.parent }

func (route *FibRoute_Output_Route) GetParentYangName() string { return "output" }

// FibRoute_Output_Route_NextHop
// Route's next-hop attribute.
type FibRoute_Output_Route_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // IP address. The type is string.
    IetfRoutingNextHopAddress interface{}

    // IPv4 address of the next-hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IetfIpv4UnicastRoutingNextHopAddress interface{}

    // IPv6 address of the next-hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IetfIpv6UnicastRoutingNextHopAddress interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}
}

func (nextHop *FibRoute_Output_Route_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *FibRoute_Output_Route_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *FibRoute_Output_Route_NextHop) GetGoName(yname string) string {
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    if yname == "next-hop-address" { return "IetfRoutingNextHopAddress" }
    if yname == "next-hop-address" { return "IetfIpv4UnicastRoutingNextHopAddress" }
    if yname == "next-hop-address" { return "IetfIpv6UnicastRoutingNextHopAddress" }
    if yname == "special-next-hop" { return "SpecialNextHop" }
    return ""
}

func (nextHop *FibRoute_Output_Route_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *FibRoute_Output_Route_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *FibRoute_Output_Route_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *FibRoute_Output_Route_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outgoing-interface"] = nextHop.OutgoingInterface
    leafs["next-hop-address"] = nextHop.IetfRoutingNextHopAddress
    leafs["next-hop-address"] = nextHop.IetfIpv4UnicastRoutingNextHopAddress
    leafs["next-hop-address"] = nextHop.IetfIpv6UnicastRoutingNextHopAddress
    leafs["special-next-hop"] = nextHop.SpecialNextHop
    return leafs
}

func (nextHop *FibRoute_Output_Route_NextHop) GetBundleName() string { return "ietf" }

func (nextHop *FibRoute_Output_Route_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *FibRoute_Output_Route_NextHop) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (nextHop *FibRoute_Output_Route_NextHop) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (nextHop *FibRoute_Output_Route_NextHop) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (nextHop *FibRoute_Output_Route_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *FibRoute_Output_Route_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *FibRoute_Output_Route_NextHop) GetParentYangName() string { return "route" }

// FibRoute_Output_Route_NextHop_SpecialNextHop represents Special next-hop options.
type FibRoute_Output_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    FibRoute_Output_Route_NextHop_SpecialNextHop_blackhole FibRoute_Output_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    FibRoute_Output_Route_NextHop_SpecialNextHop_unreachable FibRoute_Output_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    FibRoute_Output_Route_NextHop_SpecialNextHop_prohibit FibRoute_Output_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    FibRoute_Output_Route_NextHop_SpecialNextHop_receive FibRoute_Output_Route_NextHop_SpecialNextHop = "receive"
)

