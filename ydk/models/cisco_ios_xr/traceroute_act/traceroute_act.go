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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Traceroute_Input

    
    Output Traceroute_Output
}

func (traceroute *Traceroute) GetEntityData() *types.CommonEntityData {
    traceroute.EntityData.YFilter = traceroute.YFilter
    traceroute.EntityData.YangName = "traceroute"
    traceroute.EntityData.BundleName = "cisco_ios_xr"
    traceroute.EntityData.ParentYangName = "Cisco-IOS-XR-traceroute-act"
    traceroute.EntityData.SegmentPath = "Cisco-IOS-XR-traceroute-act:traceroute"
    traceroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceroute.EntityData.Children = make(map[string]types.YChild)
    traceroute.EntityData.Children["input"] = types.YChild{"Input", &traceroute.Input}
    traceroute.EntityData.Children["output"] = types.YChild{"Output", &traceroute.Output}
    traceroute.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(traceroute.EntityData)
}

// Traceroute_Input
type Traceroute_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Destination Traceroute_Input_Destination

    
    Ipv4 Traceroute_Input_Ipv4

    
    Ipv6 Traceroute_Input_Ipv6
}

func (input *Traceroute_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "traceroute"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["destination"] = types.YChild{"Destination", &input.Destination}
    input.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &input.Ipv4}
    input.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &input.Ipv6}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// Traceroute_Input_Destination
type Traceroute_Input_Destination struct {
    EntityData types.CommonEntityData
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

func (destination *Traceroute_Input_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "input"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = make(map[string]types.YChild)
    destination.EntityData.Leafs = make(map[string]types.YLeaf)
    destination.EntityData.Leafs["destination"] = types.YLeaf{"Destination", destination.Destination}
    destination.EntityData.Leafs["source"] = types.YLeaf{"Source", destination.Source}
    destination.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", destination.Timeout}
    destination.EntityData.Leafs["probe"] = types.YLeaf{"Probe", destination.Probe}
    destination.EntityData.Leafs["numeric"] = types.YLeaf{"Numeric", destination.Numeric}
    destination.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", destination.VrfName}
    destination.EntityData.Leafs["min-ttl"] = types.YLeaf{"MinTtl", destination.MinTtl}
    destination.EntityData.Leafs["max-ttl"] = types.YLeaf{"MaxTtl", destination.MaxTtl}
    destination.EntityData.Leafs["port"] = types.YLeaf{"Port", destination.Port}
    destination.EntityData.Leafs["verbose"] = types.YLeaf{"Verbose", destination.Verbose}
    destination.EntityData.Leafs["priority"] = types.YLeaf{"Priority", destination.Priority}
    destination.EntityData.Leafs["outgoing-interface"] = types.YLeaf{"OutgoingInterface", destination.OutgoingInterface}
    return &(destination.EntityData)
}

// Traceroute_Input_Ipv4
type Traceroute_Input_Ipv4 struct {
    EntityData types.CommonEntityData
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

func (ipv4 *Traceroute_Input_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "input"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["destination"] = types.YLeaf{"Destination", ipv4.Destination}
    ipv4.EntityData.Leafs["source"] = types.YLeaf{"Source", ipv4.Source}
    ipv4.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", ipv4.Timeout}
    ipv4.EntityData.Leafs["probe"] = types.YLeaf{"Probe", ipv4.Probe}
    ipv4.EntityData.Leafs["numeric"] = types.YLeaf{"Numeric", ipv4.Numeric}
    ipv4.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4.VrfName}
    ipv4.EntityData.Leafs["min-ttl"] = types.YLeaf{"MinTtl", ipv4.MinTtl}
    ipv4.EntityData.Leafs["max-ttl"] = types.YLeaf{"MaxTtl", ipv4.MaxTtl}
    ipv4.EntityData.Leafs["port"] = types.YLeaf{"Port", ipv4.Port}
    ipv4.EntityData.Leafs["verbose"] = types.YLeaf{"Verbose", ipv4.Verbose}
    return &(ipv4.EntityData)
}

// Traceroute_Input_Ipv6
type Traceroute_Input_Ipv6 struct {
    EntityData types.CommonEntityData
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

func (ipv6 *Traceroute_Input_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "input"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["destination"] = types.YLeaf{"Destination", ipv6.Destination}
    ipv6.EntityData.Leafs["source"] = types.YLeaf{"Source", ipv6.Source}
    ipv6.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", ipv6.Timeout}
    ipv6.EntityData.Leafs["probe"] = types.YLeaf{"Probe", ipv6.Probe}
    ipv6.EntityData.Leafs["numeric"] = types.YLeaf{"Numeric", ipv6.Numeric}
    ipv6.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6.VrfName}
    ipv6.EntityData.Leafs["min-ttl"] = types.YLeaf{"MinTtl", ipv6.MinTtl}
    ipv6.EntityData.Leafs["max-ttl"] = types.YLeaf{"MaxTtl", ipv6.MaxTtl}
    ipv6.EntityData.Leafs["port"] = types.YLeaf{"Port", ipv6.Port}
    ipv6.EntityData.Leafs["verbose"] = types.YLeaf{"Verbose", ipv6.Verbose}
    ipv6.EntityData.Leafs["priority"] = types.YLeaf{"Priority", ipv6.Priority}
    ipv6.EntityData.Leafs["outgoing-interface"] = types.YLeaf{"OutgoingInterface", ipv6.OutgoingInterface}
    return &(ipv6.EntityData)
}

// Traceroute_Output
type Traceroute_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    TracerouteResponse Traceroute_Output_TracerouteResponse
}

func (output *Traceroute_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "traceroute"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Children["traceroute-response"] = types.YChild{"TracerouteResponse", &output.TracerouteResponse}
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(output.EntityData)
}

// Traceroute_Output_TracerouteResponse
type Traceroute_Output_TracerouteResponse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ipv4 Traceroute_Output_TracerouteResponse_Ipv4

    
    Ipv6 Traceroute_Output_TracerouteResponse_Ipv6
}

func (tracerouteResponse *Traceroute_Output_TracerouteResponse) GetEntityData() *types.CommonEntityData {
    tracerouteResponse.EntityData.YFilter = tracerouteResponse.YFilter
    tracerouteResponse.EntityData.YangName = "traceroute-response"
    tracerouteResponse.EntityData.BundleName = "cisco_ios_xr"
    tracerouteResponse.EntityData.ParentYangName = "output"
    tracerouteResponse.EntityData.SegmentPath = "traceroute-response"
    tracerouteResponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracerouteResponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracerouteResponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracerouteResponse.EntityData.Children = make(map[string]types.YChild)
    tracerouteResponse.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &tracerouteResponse.Ipv4}
    tracerouteResponse.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &tracerouteResponse.Ipv6}
    tracerouteResponse.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tracerouteResponse.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4
type Traceroute_Output_TracerouteResponse_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination address or hostname. The type is string.
    Destination interface{}

    // Verbose output. The type is string.
    VerboseOutput interface{}

    
    Hops Traceroute_Output_TracerouteResponse_Ipv4_Hops
}

func (ipv4 *Traceroute_Output_TracerouteResponse_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "traceroute-response"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["hops"] = types.YChild{"Hops", &ipv4.Hops}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["destination"] = types.YLeaf{"Destination", ipv4.Destination}
    ipv4.EntityData.Leafs["verbose-output"] = types.YLeaf{"VerboseOutput", ipv4.VerboseOutput}
    return &(ipv4.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops
type Traceroute_Output_TracerouteResponse_Ipv4_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop.
    Hop []Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "ipv4"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Children["hop"] = types.YChild{"Hop", nil}
    for i := range hops.Hop {
        hops.EntityData.Children[types.GetSegmentPath(&hops.Hop[i])] = types.YChild{"Hop", &hops.Hop[i]}
    }
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hops.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the hop. The type is interface{} with
    // range: 0..4294967295.
    HopIndex interface{}

    // Address of the hop. The type is string.
    HopAddress interface{}

    // Hostname of the hop. The type is string.
    HopHostname interface{}

    
    Probes Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes
}

func (hop *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop) GetEntityData() *types.CommonEntityData {
    hop.EntityData.YFilter = hop.YFilter
    hop.EntityData.YangName = "hop"
    hop.EntityData.BundleName = "cisco_ios_xr"
    hop.EntityData.ParentYangName = "hops"
    hop.EntityData.SegmentPath = "hop" + "[hop-index='" + fmt.Sprintf("%v", hop.HopIndex) + "']"
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = make(map[string]types.YChild)
    hop.EntityData.Children["probes"] = types.YChild{"Probes", &hop.Probes}
    hop.EntityData.Leafs = make(map[string]types.YLeaf)
    hop.EntityData.Leafs["hop-index"] = types.YLeaf{"HopIndex", hop.HopIndex}
    hop.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", hop.HopAddress}
    hop.EntityData.Leafs["hop-hostname"] = types.YLeaf{"HopHostname", hop.HopHostname}
    return &(hop.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe.
    Probe []Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes) GetEntityData() *types.CommonEntityData {
    probes.EntityData.YFilter = probes.YFilter
    probes.EntityData.YangName = "probes"
    probes.EntityData.BundleName = "cisco_ios_xr"
    probes.EntityData.ParentYangName = "hop"
    probes.EntityData.SegmentPath = "probes"
    probes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probes.EntityData.Children = make(map[string]types.YChild)
    probes.EntityData.Children["probe"] = types.YChild{"Probe", nil}
    for i := range probes.Probe {
        probes.EntityData.Children[types.GetSegmentPath(&probes.Probe[i])] = types.YChild{"Probe", &probes.Probe[i]}
    }
    probes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(probes.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe struct {
    EntityData types.CommonEntityData
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

func (probe *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe) GetEntityData() *types.CommonEntityData {
    probe.EntityData.YFilter = probe.YFilter
    probe.EntityData.YangName = "probe"
    probe.EntityData.BundleName = "cisco_ios_xr"
    probe.EntityData.ParentYangName = "probes"
    probe.EntityData.SegmentPath = "probe" + "[probe-index='" + fmt.Sprintf("%v", probe.ProbeIndex) + "']"
    probe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probe.EntityData.Children = make(map[string]types.YChild)
    probe.EntityData.Leafs = make(map[string]types.YLeaf)
    probe.EntityData.Leafs["probe-index"] = types.YLeaf{"ProbeIndex", probe.ProbeIndex}
    probe.EntityData.Leafs["result"] = types.YLeaf{"Result", probe.Result}
    probe.EntityData.Leafs["delta-time"] = types.YLeaf{"DeltaTime", probe.DeltaTime}
    probe.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", probe.HopAddress}
    probe.EntityData.Leafs["hop-hostname"] = types.YLeaf{"HopHostname", probe.HopHostname}
    return &(probe.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6
type Traceroute_Output_TracerouteResponse_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination address or hostname. The type is string.
    Destination interface{}

    // Verbose output. The type is string.
    VerboseOutput interface{}

    
    Hops Traceroute_Output_TracerouteResponse_Ipv6_Hops
}

func (ipv6 *Traceroute_Output_TracerouteResponse_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "traceroute-response"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["hops"] = types.YChild{"Hops", &ipv6.Hops}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["destination"] = types.YLeaf{"Destination", ipv6.Destination}
    ipv6.EntityData.Leafs["verbose-output"] = types.YLeaf{"VerboseOutput", ipv6.VerboseOutput}
    return &(ipv6.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops
type Traceroute_Output_TracerouteResponse_Ipv6_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop.
    Hop []Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "ipv6"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Children["hop"] = types.YChild{"Hop", nil}
    for i := range hops.Hop {
        hops.EntityData.Children[types.GetSegmentPath(&hops.Hop[i])] = types.YChild{"Hop", &hops.Hop[i]}
    }
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hops.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the hop. The type is interface{} with
    // range: 0..4294967295.
    HopIndex interface{}

    // Address of the hop. The type is string.
    HopAddress interface{}

    // Hostname of the hop. The type is string.
    HopHostname interface{}

    
    Probes Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes
}

func (hop *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop) GetEntityData() *types.CommonEntityData {
    hop.EntityData.YFilter = hop.YFilter
    hop.EntityData.YangName = "hop"
    hop.EntityData.BundleName = "cisco_ios_xr"
    hop.EntityData.ParentYangName = "hops"
    hop.EntityData.SegmentPath = "hop" + "[hop-index='" + fmt.Sprintf("%v", hop.HopIndex) + "']"
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = make(map[string]types.YChild)
    hop.EntityData.Children["probes"] = types.YChild{"Probes", &hop.Probes}
    hop.EntityData.Leafs = make(map[string]types.YLeaf)
    hop.EntityData.Leafs["hop-index"] = types.YLeaf{"HopIndex", hop.HopIndex}
    hop.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", hop.HopAddress}
    hop.EntityData.Leafs["hop-hostname"] = types.YLeaf{"HopHostname", hop.HopHostname}
    return &(hop.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe.
    Probe []Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes) GetEntityData() *types.CommonEntityData {
    probes.EntityData.YFilter = probes.YFilter
    probes.EntityData.YangName = "probes"
    probes.EntityData.BundleName = "cisco_ios_xr"
    probes.EntityData.ParentYangName = "hop"
    probes.EntityData.SegmentPath = "probes"
    probes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probes.EntityData.Children = make(map[string]types.YChild)
    probes.EntityData.Children["probe"] = types.YChild{"Probe", nil}
    for i := range probes.Probe {
        probes.EntityData.Children[types.GetSegmentPath(&probes.Probe[i])] = types.YChild{"Probe", &probes.Probe[i]}
    }
    probes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(probes.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe struct {
    EntityData types.CommonEntityData
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

func (probe *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe) GetEntityData() *types.CommonEntityData {
    probe.EntityData.YFilter = probe.YFilter
    probe.EntityData.YangName = "probe"
    probe.EntityData.BundleName = "cisco_ios_xr"
    probe.EntityData.ParentYangName = "probes"
    probe.EntityData.SegmentPath = "probe" + "[probe-index='" + fmt.Sprintf("%v", probe.ProbeIndex) + "']"
    probe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probe.EntityData.Children = make(map[string]types.YChild)
    probe.EntityData.Leafs = make(map[string]types.YLeaf)
    probe.EntityData.Leafs["probe-index"] = types.YLeaf{"ProbeIndex", probe.ProbeIndex}
    probe.EntityData.Leafs["result"] = types.YLeaf{"Result", probe.Result}
    probe.EntityData.Leafs["delta-time"] = types.YLeaf{"DeltaTime", probe.DeltaTime}
    probe.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", probe.HopAddress}
    probe.EntityData.Leafs["hop-hostname"] = types.YLeaf{"HopHostname", probe.HopHostname}
    return &(probe.EntityData)
}

