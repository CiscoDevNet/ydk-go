// This module contains a collection of YANG definitions
// for Cisco IOS-XR ping action package configuration.
// 
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package traceroute_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package traceroute_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-traceroute-act traceroute}", reflect.TypeOf(Traceroute{}))
    ydk.RegisterEntity("Cisco-IOS-XR-traceroute-act:traceroute", reflect.TypeOf(Traceroute{}))
}

// Traceroute
// Trace route to destination
type Traceroute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Traceroute_Input

    
    Output Traceroute_Output
}

func (traceroute *Traceroute) GetFilter() yfilter.YFilter { return traceroute.YFilter }

func (traceroute *Traceroute) SetFilter(yf yfilter.YFilter) { traceroute.YFilter = yf }

func (traceroute *Traceroute) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (traceroute *Traceroute) GetSegmentPath() string {
    return "Cisco-IOS-XR-traceroute-act:traceroute"
}

func (traceroute *Traceroute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &traceroute.Input
    }
    if childYangName == "output" {
        return &traceroute.Output
    }
    return nil
}

func (traceroute *Traceroute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &traceroute.Input
    children["output"] = &traceroute.Output
    return children
}

func (traceroute *Traceroute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (traceroute *Traceroute) GetBundleName() string { return "cisco_ios_xr" }

func (traceroute *Traceroute) GetYangName() string { return "traceroute" }

func (traceroute *Traceroute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traceroute *Traceroute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traceroute *Traceroute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traceroute *Traceroute) SetParent(parent types.Entity) { traceroute.parent = parent }

func (traceroute *Traceroute) GetParent() types.Entity { return traceroute.parent }

func (traceroute *Traceroute) GetParentYangName() string { return "Cisco-IOS-XR-traceroute-act" }

// Traceroute_Input
type Traceroute_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Destination Traceroute_Input_Destination

    
    Ipv4 Traceroute_Input_Ipv4

    
    Ipv6 Traceroute_Input_Ipv6
}

func (input *Traceroute_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Traceroute_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Traceroute_Input) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (input *Traceroute_Input) GetSegmentPath() string {
    return "input"
}

func (input *Traceroute_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        return &input.Destination
    }
    if childYangName == "ipv4" {
        return &input.Ipv4
    }
    if childYangName == "ipv6" {
        return &input.Ipv6
    }
    return nil
}

func (input *Traceroute_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination"] = &input.Destination
    children["ipv4"] = &input.Ipv4
    children["ipv6"] = &input.Ipv6
    return children
}

func (input *Traceroute_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *Traceroute_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Traceroute_Input) GetYangName() string { return "input" }

func (input *Traceroute_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Traceroute_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Traceroute_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Traceroute_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Traceroute_Input) GetParent() types.Entity { return input.parent }

func (input *Traceroute_Input) GetParentYangName() string { return "traceroute" }

// Traceroute_Input_Destination
type Traceroute_Input_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination address or hostname. The type is string. This attribute is
    // mandatory.
    Destination interface{}

    // Source address or interface. The type is string.
    Source interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 3.
    Timeout interface{}

    // Probe count. The type is interface{} with range: 1..64. The default value
    // is 3.
    Probe interface{}

    // Numeric display only. The type is bool.
    Numeric interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // minimum time to live. The type is interface{} with range: 0..255. The
    // default value is 1.
    MinTtl interface{}

    // maximum time to live. The type is interface{} with range: 0..255. The
    // default value is 30.
    MaxTtl interface{}

    // Port numbe. The type is interface{} with range: 0..65535.
    Port interface{}

    // verbose output. The type is bool.
    Verbose interface{}

    // Priority of hte packet. The type is interface{} with range: 0..15.
    Priority interface{}

    // Outgoing interface, needed in case of traceroute to link local address. The
    // type is string.
    OutgoingInterface interface{}
}

func (destination *Traceroute_Input_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *Traceroute_Input_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *Traceroute_Input_Destination) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "source" { return "Source" }
    if yname == "timeout" { return "Timeout" }
    if yname == "probe" { return "Probe" }
    if yname == "numeric" { return "Numeric" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "min-ttl" { return "MinTtl" }
    if yname == "max-ttl" { return "MaxTtl" }
    if yname == "port" { return "Port" }
    if yname == "verbose" { return "Verbose" }
    if yname == "priority" { return "Priority" }
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    return ""
}

func (destination *Traceroute_Input_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *Traceroute_Input_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destination *Traceroute_Input_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destination *Traceroute_Input_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = destination.Destination
    leafs["source"] = destination.Source
    leafs["timeout"] = destination.Timeout
    leafs["probe"] = destination.Probe
    leafs["numeric"] = destination.Numeric
    leafs["vrf-name"] = destination.VrfName
    leafs["min-ttl"] = destination.MinTtl
    leafs["max-ttl"] = destination.MaxTtl
    leafs["port"] = destination.Port
    leafs["verbose"] = destination.Verbose
    leafs["priority"] = destination.Priority
    leafs["outgoing-interface"] = destination.OutgoingInterface
    return leafs
}

func (destination *Traceroute_Input_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *Traceroute_Input_Destination) GetYangName() string { return "destination" }

func (destination *Traceroute_Input_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *Traceroute_Input_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *Traceroute_Input_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *Traceroute_Input_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *Traceroute_Input_Destination) GetParent() types.Entity { return destination.parent }

func (destination *Traceroute_Input_Destination) GetParentYangName() string { return "input" }

// Traceroute_Input_Ipv4
type Traceroute_Input_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination address or hostname. The type is string. This attribute is
    // mandatory.
    Destination interface{}

    // Source address or interface. The type is string.
    Source interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 3.
    Timeout interface{}

    // Probe count. The type is interface{} with range: 1..64. The default value
    // is 3.
    Probe interface{}

    // Numeric display only. The type is bool.
    Numeric interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // minimum time to live. The type is interface{} with range: 0..255. The
    // default value is 1.
    MinTtl interface{}

    // maximum time to live. The type is interface{} with range: 0..255. The
    // default value is 30.
    MaxTtl interface{}

    // Port numbe. The type is interface{} with range: 0..65535.
    Port interface{}

    // verbose output. The type is bool.
    Verbose interface{}
}

func (ipv4 *Traceroute_Input_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Traceroute_Input_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Traceroute_Input_Ipv4) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "source" { return "Source" }
    if yname == "timeout" { return "Timeout" }
    if yname == "probe" { return "Probe" }
    if yname == "numeric" { return "Numeric" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "min-ttl" { return "MinTtl" }
    if yname == "max-ttl" { return "MaxTtl" }
    if yname == "port" { return "Port" }
    if yname == "verbose" { return "Verbose" }
    return ""
}

func (ipv4 *Traceroute_Input_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Traceroute_Input_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *Traceroute_Input_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *Traceroute_Input_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv4.Destination
    leafs["source"] = ipv4.Source
    leafs["timeout"] = ipv4.Timeout
    leafs["probe"] = ipv4.Probe
    leafs["numeric"] = ipv4.Numeric
    leafs["vrf-name"] = ipv4.VrfName
    leafs["min-ttl"] = ipv4.MinTtl
    leafs["max-ttl"] = ipv4.MaxTtl
    leafs["port"] = ipv4.Port
    leafs["verbose"] = ipv4.Verbose
    return leafs
}

func (ipv4 *Traceroute_Input_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Traceroute_Input_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Traceroute_Input_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Traceroute_Input_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Traceroute_Input_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Traceroute_Input_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Traceroute_Input_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Traceroute_Input_Ipv4) GetParentYangName() string { return "input" }

// Traceroute_Input_Ipv6
type Traceroute_Input_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination address or hostname. The type is string. This attribute is
    // mandatory.
    Destination interface{}

    // Source address or interface. The type is string.
    Source interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 3.
    Timeout interface{}

    // Probe count. The type is interface{} with range: 1..64. The default value
    // is 3.
    Probe interface{}

    // Numeric display only. The type is bool.
    Numeric interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // minimum time to live. The type is interface{} with range: 0..255. The
    // default value is 1.
    MinTtl interface{}

    // maximum time to live. The type is interface{} with range: 0..255. The
    // default value is 30.
    MaxTtl interface{}

    // Port numbe. The type is interface{} with range: 0..65535.
    Port interface{}

    // verbose output. The type is bool.
    Verbose interface{}

    // Priority of hte packet. The type is interface{} with range: 0..15.
    Priority interface{}

    // Outgoing interface, needed in case of traceroute to link local address. The
    // type is string.
    OutgoingInterface interface{}
}

func (ipv6 *Traceroute_Input_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Traceroute_Input_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Traceroute_Input_Ipv6) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "source" { return "Source" }
    if yname == "timeout" { return "Timeout" }
    if yname == "probe" { return "Probe" }
    if yname == "numeric" { return "Numeric" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "min-ttl" { return "MinTtl" }
    if yname == "max-ttl" { return "MaxTtl" }
    if yname == "port" { return "Port" }
    if yname == "verbose" { return "Verbose" }
    if yname == "priority" { return "Priority" }
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    return ""
}

func (ipv6 *Traceroute_Input_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Traceroute_Input_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *Traceroute_Input_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *Traceroute_Input_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv6.Destination
    leafs["source"] = ipv6.Source
    leafs["timeout"] = ipv6.Timeout
    leafs["probe"] = ipv6.Probe
    leafs["numeric"] = ipv6.Numeric
    leafs["vrf-name"] = ipv6.VrfName
    leafs["min-ttl"] = ipv6.MinTtl
    leafs["max-ttl"] = ipv6.MaxTtl
    leafs["port"] = ipv6.Port
    leafs["verbose"] = ipv6.Verbose
    leafs["priority"] = ipv6.Priority
    leafs["outgoing-interface"] = ipv6.OutgoingInterface
    return leafs
}

func (ipv6 *Traceroute_Input_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Traceroute_Input_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Traceroute_Input_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Traceroute_Input_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Traceroute_Input_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Traceroute_Input_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Traceroute_Input_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Traceroute_Input_Ipv6) GetParentYangName() string { return "input" }

// Traceroute_Output
type Traceroute_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    TracerouteResponse Traceroute_Output_TracerouteResponse
}

func (output *Traceroute_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Traceroute_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Traceroute_Output) GetGoName(yname string) string {
    if yname == "traceroute-response" { return "TracerouteResponse" }
    return ""
}

func (output *Traceroute_Output) GetSegmentPath() string {
    return "output"
}

func (output *Traceroute_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traceroute-response" {
        return &output.TracerouteResponse
    }
    return nil
}

func (output *Traceroute_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traceroute-response"] = &output.TracerouteResponse
    return children
}

func (output *Traceroute_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (output *Traceroute_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *Traceroute_Output) GetYangName() string { return "output" }

func (output *Traceroute_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *Traceroute_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *Traceroute_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *Traceroute_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Traceroute_Output) GetParent() types.Entity { return output.parent }

func (output *Traceroute_Output) GetParentYangName() string { return "traceroute" }

// Traceroute_Output_TracerouteResponse
type Traceroute_Output_TracerouteResponse struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ipv4 Traceroute_Output_TracerouteResponse_Ipv4

    
    Ipv6 Traceroute_Output_TracerouteResponse_Ipv6
}

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetFilter() yfilter.YFilter { return tracerouteResponse.YFilter }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) SetFilter(yf yfilter.YFilter) { tracerouteResponse.YFilter = yf }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetSegmentPath() string {
    return "traceroute-response"
}

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &tracerouteResponse.Ipv4
    }
    if childYangName == "ipv6" {
        return &tracerouteResponse.Ipv6
    }
    return nil
}

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &tracerouteResponse.Ipv4
    children["ipv6"] = &tracerouteResponse.Ipv6
    return children
}

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetBundleName() string { return "cisco_ios_xr" }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetYangName() string { return "traceroute-response" }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) SetParent(parent types.Entity) { tracerouteResponse.parent = parent }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetParent() types.Entity { return tracerouteResponse.parent }

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetParentYangName() string { return "output" }

// Traceroute_Output_TracerouteResponse_Ipv4
type Traceroute_Output_TracerouteResponse_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination address or hostname. The type is string.
    Destination interface{}

    // Verbose output. The type is string.
    VerboseOutput interface{}

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv4_Hops.
    Hops []Traceroute_Output_TracerouteResponse_Ipv4_Hops
}

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "verbose-output" { return "VerboseOutput" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hops" {
        for _, c := range ipv4.Hops {
            if ipv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Traceroute_Output_TracerouteResponse_Ipv4_Hops{}
        ipv4.Hops = append(ipv4.Hops, child)
        return &ipv4.Hops[len(ipv4.Hops)-1]
    }
    return nil
}

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4.Hops {
        children[ipv4.Hops[i].GetSegmentPath()] = &ipv4.Hops[i]
    }
    return children
}

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv4.Destination
    leafs["verbose-output"] = ipv4.VerboseOutput
    return leafs
}

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetParentYangName() string { return "traceroute-response" }

// Traceroute_Output_TracerouteResponse_Ipv4_Hops
type Traceroute_Output_TracerouteResponse_Ipv4_Hops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the hop. The type is interface{} with
    // range: 0..4294967295.
    HopIndex interface{}

    // Address of the hop. The type is string.
    HopAddress interface{}

    // Hostname of the hop. The type is string.
    HopHostname interface{}

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes.
    Probes []Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetFilter() yfilter.YFilter { return hops.YFilter }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) SetFilter(yf yfilter.YFilter) { hops.YFilter = yf }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetGoName(yname string) string {
    if yname == "hop-index" { return "HopIndex" }
    if yname == "hop-address" { return "HopAddress" }
    if yname == "hop-hostname" { return "HopHostname" }
    if yname == "probes" { return "Probes" }
    return ""
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetSegmentPath() string {
    return "hops" + "[hop-index='" + fmt.Sprintf("%v", hops.HopIndex) + "']"
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "probes" {
        for _, c := range hops.Probes {
            if hops.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes{}
        hops.Probes = append(hops.Probes, child)
        return &hops.Probes[len(hops.Probes)-1]
    }
    return nil
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hops.Probes {
        children[hops.Probes[i].GetSegmentPath()] = &hops.Probes[i]
    }
    return children
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hop-index"] = hops.HopIndex
    leafs["hop-address"] = hops.HopAddress
    leafs["hop-hostname"] = hops.HopHostname
    return leafs
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetBundleName() string { return "cisco_ios_xr" }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetYangName() string { return "hops" }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) SetParent(parent types.Entity) { hops.parent = parent }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetParent() types.Entity { return hops.parent }

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetParentYangName() string { return "ipv4" }

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the probe. The type is interface{} with
    // range: 0..4294967295.
    ProbeIndex interface{}

    // Response for each probe. The type is string.
    Result interface{}

    // Delta time in seconds. The type is interface{} with range: 0..4294967295.
    DeltaTime interface{}

    // Address of the hop. The type is string.
    HopAddress interface{}

    // Hostname of the hop. The type is string.
    HopHostname interface{}
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetFilter() yfilter.YFilter { return probes.YFilter }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) SetFilter(yf yfilter.YFilter) { probes.YFilter = yf }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetGoName(yname string) string {
    if yname == "probe-index" { return "ProbeIndex" }
    if yname == "result" { return "Result" }
    if yname == "delta-time" { return "DeltaTime" }
    if yname == "hop-address" { return "HopAddress" }
    if yname == "hop-hostname" { return "HopHostname" }
    return ""
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetSegmentPath() string {
    return "probes" + "[probe-index='" + fmt.Sprintf("%v", probes.ProbeIndex) + "']"
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["probe-index"] = probes.ProbeIndex
    leafs["result"] = probes.Result
    leafs["delta-time"] = probes.DeltaTime
    leafs["hop-address"] = probes.HopAddress
    leafs["hop-hostname"] = probes.HopHostname
    return leafs
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetBundleName() string { return "cisco_ios_xr" }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetYangName() string { return "probes" }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) SetParent(parent types.Entity) { probes.parent = parent }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetParent() types.Entity { return probes.parent }

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Probes) GetParentYangName() string { return "hops" }

// Traceroute_Output_TracerouteResponse_Ipv6
type Traceroute_Output_TracerouteResponse_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination address or hostname. The type is string.
    Destination interface{}

    // Verbose output. The type is string.
    VerboseOutput interface{}

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv6_Hops.
    Hops []Traceroute_Output_TracerouteResponse_Ipv6_Hops
}

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "verbose-output" { return "VerboseOutput" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hops" {
        for _, c := range ipv6.Hops {
            if ipv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Traceroute_Output_TracerouteResponse_Ipv6_Hops{}
        ipv6.Hops = append(ipv6.Hops, child)
        return &ipv6.Hops[len(ipv6.Hops)-1]
    }
    return nil
}

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6.Hops {
        children[ipv6.Hops[i].GetSegmentPath()] = &ipv6.Hops[i]
    }
    return children
}

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv6.Destination
    leafs["verbose-output"] = ipv6.VerboseOutput
    return leafs
}

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetParentYangName() string { return "traceroute-response" }

// Traceroute_Output_TracerouteResponse_Ipv6_Hops
type Traceroute_Output_TracerouteResponse_Ipv6_Hops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the hop. The type is interface{} with
    // range: 0..4294967295.
    HopIndex interface{}

    // Address of the hop. The type is string.
    HopAddress interface{}

    // Hostname of the hop. The type is string.
    HopHostname interface{}

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes.
    Probes []Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetFilter() yfilter.YFilter { return hops.YFilter }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) SetFilter(yf yfilter.YFilter) { hops.YFilter = yf }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetGoName(yname string) string {
    if yname == "hop-index" { return "HopIndex" }
    if yname == "hop-address" { return "HopAddress" }
    if yname == "hop-hostname" { return "HopHostname" }
    if yname == "probes" { return "Probes" }
    return ""
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetSegmentPath() string {
    return "hops" + "[hop-index='" + fmt.Sprintf("%v", hops.HopIndex) + "']"
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "probes" {
        for _, c := range hops.Probes {
            if hops.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes{}
        hops.Probes = append(hops.Probes, child)
        return &hops.Probes[len(hops.Probes)-1]
    }
    return nil
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hops.Probes {
        children[hops.Probes[i].GetSegmentPath()] = &hops.Probes[i]
    }
    return children
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hop-index"] = hops.HopIndex
    leafs["hop-address"] = hops.HopAddress
    leafs["hop-hostname"] = hops.HopHostname
    return leafs
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetBundleName() string { return "cisco_ios_xr" }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetYangName() string { return "hops" }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) SetParent(parent types.Entity) { hops.parent = parent }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetParent() types.Entity { return hops.parent }

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetParentYangName() string { return "ipv6" }

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the probe. The type is interface{} with
    // range: 0..4294967295.
    ProbeIndex interface{}

    // Response for each probe. The type is string.
    Result interface{}

    // Delta time in seconds. The type is interface{} with range: 0..4294967295.
    DeltaTime interface{}

    // Address of the hop. The type is string.
    HopAddress interface{}

    // Hostname of the hop. The type is string.
    HopHostname interface{}
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetFilter() yfilter.YFilter { return probes.YFilter }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) SetFilter(yf yfilter.YFilter) { probes.YFilter = yf }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetGoName(yname string) string {
    if yname == "probe-index" { return "ProbeIndex" }
    if yname == "result" { return "Result" }
    if yname == "delta-time" { return "DeltaTime" }
    if yname == "hop-address" { return "HopAddress" }
    if yname == "hop-hostname" { return "HopHostname" }
    return ""
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetSegmentPath() string {
    return "probes" + "[probe-index='" + fmt.Sprintf("%v", probes.ProbeIndex) + "']"
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["probe-index"] = probes.ProbeIndex
    leafs["result"] = probes.Result
    leafs["delta-time"] = probes.DeltaTime
    leafs["hop-address"] = probes.HopAddress
    leafs["hop-hostname"] = probes.HopHostname
    return leafs
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetBundleName() string { return "cisco_ios_xr" }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetYangName() string { return "probes" }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) SetParent(parent types.Entity) { probes.parent = parent }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetParent() types.Entity { return probes.parent }

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Probes) GetParentYangName() string { return "hops" }

