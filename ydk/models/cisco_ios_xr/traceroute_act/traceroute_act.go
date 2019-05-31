// This module contains a collection of YANG definitions
// for Cisco IOS-XR ping action package configuration.
// 
// Copyright (c) 2016, 2018 by Cisco Systems, Inc.
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
    traceroute.EntityData.AbsolutePath = traceroute.EntityData.SegmentPath
    traceroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceroute.EntityData.Children = types.NewOrderedMap()
    traceroute.EntityData.Children.Append("input", types.YChild{"Input", &traceroute.Input})
    traceroute.EntityData.Children.Append("output", types.YChild{"Output", &traceroute.Output})
    traceroute.EntityData.Leafs = types.NewOrderedMap()

    traceroute.EntityData.YListKeys = []string {}

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
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("destination", types.YChild{"Destination", &input.Destination})
    input.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &input.Ipv4})
    input.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &input.Ipv6})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Traceroute_Input_Destination
// This type is a presence type.
type Traceroute_Input_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    // srv6 header. The type is interface{}.
    Srv6Header interface{}

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
    destination.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/input/" + destination.EntityData.SegmentPath
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = types.NewOrderedMap()
    destination.EntityData.Leafs = types.NewOrderedMap()
    destination.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", destination.Destination})
    destination.EntityData.Leafs.Append("source", types.YLeaf{"Source", destination.Source})
    destination.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", destination.Timeout})
    destination.EntityData.Leafs.Append("probe", types.YLeaf{"Probe", destination.Probe})
    destination.EntityData.Leafs.Append("numeric", types.YLeaf{"Numeric", destination.Numeric})
    destination.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", destination.VrfName})
    destination.EntityData.Leafs.Append("min-ttl", types.YLeaf{"MinTtl", destination.MinTtl})
    destination.EntityData.Leafs.Append("max-ttl", types.YLeaf{"MaxTtl", destination.MaxTtl})
    destination.EntityData.Leafs.Append("port", types.YLeaf{"Port", destination.Port})
    destination.EntityData.Leafs.Append("verbose", types.YLeaf{"Verbose", destination.Verbose})
    destination.EntityData.Leafs.Append("srv6-header", types.YLeaf{"Srv6Header", destination.Srv6Header})
    destination.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", destination.Priority})
    destination.EntityData.Leafs.Append("outgoing-interface", types.YLeaf{"OutgoingInterface", destination.OutgoingInterface})

    destination.EntityData.YListKeys = []string {}

    return &(destination.EntityData)
}

// Traceroute_Input_Ipv4
// This type is a presence type.
type Traceroute_Input_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    // srv6 header. The type is interface{}.
    Srv6Header interface{}
}

func (ipv4 *Traceroute_Input_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "input"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/input/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv4.Destination})
    ipv4.EntityData.Leafs.Append("source", types.YLeaf{"Source", ipv4.Source})
    ipv4.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ipv4.Timeout})
    ipv4.EntityData.Leafs.Append("probe", types.YLeaf{"Probe", ipv4.Probe})
    ipv4.EntityData.Leafs.Append("numeric", types.YLeaf{"Numeric", ipv4.Numeric})
    ipv4.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4.VrfName})
    ipv4.EntityData.Leafs.Append("min-ttl", types.YLeaf{"MinTtl", ipv4.MinTtl})
    ipv4.EntityData.Leafs.Append("max-ttl", types.YLeaf{"MaxTtl", ipv4.MaxTtl})
    ipv4.EntityData.Leafs.Append("port", types.YLeaf{"Port", ipv4.Port})
    ipv4.EntityData.Leafs.Append("verbose", types.YLeaf{"Verbose", ipv4.Verbose})
    ipv4.EntityData.Leafs.Append("srv6-header", types.YLeaf{"Srv6Header", ipv4.Srv6Header})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Traceroute_Input_Ipv6
// This type is a presence type.
type Traceroute_Input_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    // srv6 header. The type is interface{}.
    Srv6Header interface{}

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
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/input/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv6.Destination})
    ipv6.EntityData.Leafs.Append("source", types.YLeaf{"Source", ipv6.Source})
    ipv6.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ipv6.Timeout})
    ipv6.EntityData.Leafs.Append("probe", types.YLeaf{"Probe", ipv6.Probe})
    ipv6.EntityData.Leafs.Append("numeric", types.YLeaf{"Numeric", ipv6.Numeric})
    ipv6.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6.VrfName})
    ipv6.EntityData.Leafs.Append("min-ttl", types.YLeaf{"MinTtl", ipv6.MinTtl})
    ipv6.EntityData.Leafs.Append("max-ttl", types.YLeaf{"MaxTtl", ipv6.MaxTtl})
    ipv6.EntityData.Leafs.Append("port", types.YLeaf{"Port", ipv6.Port})
    ipv6.EntityData.Leafs.Append("verbose", types.YLeaf{"Verbose", ipv6.Verbose})
    ipv6.EntityData.Leafs.Append("srv6-header", types.YLeaf{"Srv6Header", ipv6.Srv6Header})
    ipv6.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ipv6.Priority})
    ipv6.EntityData.Leafs.Append("outgoing-interface", types.YLeaf{"OutgoingInterface", ipv6.OutgoingInterface})

    ipv6.EntityData.YListKeys = []string {}

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
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("traceroute-response", types.YChild{"TracerouteResponse", &output.TracerouteResponse})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

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
    tracerouteResponse.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/" + tracerouteResponse.EntityData.SegmentPath
    tracerouteResponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracerouteResponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracerouteResponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracerouteResponse.EntityData.Children = types.NewOrderedMap()
    tracerouteResponse.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &tracerouteResponse.Ipv4})
    tracerouteResponse.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &tracerouteResponse.Ipv6})
    tracerouteResponse.EntityData.Leafs = types.NewOrderedMap()

    tracerouteResponse.EntityData.YListKeys = []string {}

    return &(tracerouteResponse.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4
// This type is a presence type.
type Traceroute_Output_TracerouteResponse_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("hops", types.YChild{"Hops", &ipv4.Hops})
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv4.Destination})
    ipv4.EntityData.Leafs.Append("verbose-output", types.YLeaf{"VerboseOutput", ipv4.VerboseOutput})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops
type Traceroute_Output_TracerouteResponse_Ipv4_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop.
    Hop []*Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv4_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "ipv4"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv4/" + hops.EntityData.SegmentPath
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = types.NewOrderedMap()
    hops.EntityData.Children.Append("hop", types.YChild{"Hop", nil})
    for i := range hops.Hop {
        hops.EntityData.Children.Append(types.GetSegmentPath(hops.Hop[i]), types.YChild{"Hop", hops.Hop[i]})
    }
    hops.EntityData.Leafs = types.NewOrderedMap()

    hops.EntityData.YListKeys = []string {}

    return &(hops.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    hop.EntityData.SegmentPath = "hop" + types.AddKeyToken(hop.HopIndex, "hop-index")
    hop.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv4/hops/" + hop.EntityData.SegmentPath
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = types.NewOrderedMap()
    hop.EntityData.Children.Append("probes", types.YChild{"Probes", &hop.Probes})
    hop.EntityData.Leafs = types.NewOrderedMap()
    hop.EntityData.Leafs.Append("hop-index", types.YLeaf{"HopIndex", hop.HopIndex})
    hop.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", hop.HopAddress})
    hop.EntityData.Leafs.Append("hop-hostname", types.YLeaf{"HopHostname", hop.HopHostname})

    hop.EntityData.YListKeys = []string {"HopIndex"}

    return &(hop.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe.
    Probe []*Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes) GetEntityData() *types.CommonEntityData {
    probes.EntityData.YFilter = probes.YFilter
    probes.EntityData.YangName = "probes"
    probes.EntityData.BundleName = "cisco_ios_xr"
    probes.EntityData.ParentYangName = "hop"
    probes.EntityData.SegmentPath = "probes"
    probes.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv4/hops/hop/" + probes.EntityData.SegmentPath
    probes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probes.EntityData.Children = types.NewOrderedMap()
    probes.EntityData.Children.Append("probe", types.YChild{"Probe", nil})
    for i := range probes.Probe {
        probes.EntityData.Children.Append(types.GetSegmentPath(probes.Probe[i]), types.YChild{"Probe", probes.Probe[i]})
    }
    probes.EntityData.Leafs = types.NewOrderedMap()

    probes.EntityData.YListKeys = []string {}

    return &(probes.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    
    Srv6Header Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header
}

func (probe *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe) GetEntityData() *types.CommonEntityData {
    probe.EntityData.YFilter = probe.YFilter
    probe.EntityData.YangName = "probe"
    probe.EntityData.BundleName = "cisco_ios_xr"
    probe.EntityData.ParentYangName = "probes"
    probe.EntityData.SegmentPath = "probe" + types.AddKeyToken(probe.ProbeIndex, "probe-index")
    probe.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv4/hops/hop/probes/" + probe.EntityData.SegmentPath
    probe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probe.EntityData.Children = types.NewOrderedMap()
    probe.EntityData.Children.Append("srv6-header", types.YChild{"Srv6Header", &probe.Srv6Header})
    probe.EntityData.Leafs = types.NewOrderedMap()
    probe.EntityData.Leafs.Append("probe-index", types.YLeaf{"ProbeIndex", probe.ProbeIndex})
    probe.EntityData.Leafs.Append("result", types.YLeaf{"Result", probe.Result})
    probe.EntityData.Leafs.Append("delta-time", types.YLeaf{"DeltaTime", probe.DeltaTime})
    probe.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", probe.HopAddress})
    probe.EntityData.Leafs.Append("hop-hostname", types.YLeaf{"HopHostname", probe.HopHostname})

    probe.EntityData.YListKeys = []string {"ProbeIndex"}

    return &(probe.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination address for srv6 header. The type is interface{} with range:
    // 0..4294967295.
    DestinationAddress interface{}

    // Number of segments left. The type is interface{} with range: 0..4294967295.
    SegmentsLeft interface{}

    
    Segments Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header_Segments
}

func (srv6Header *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header) GetEntityData() *types.CommonEntityData {
    srv6Header.EntityData.YFilter = srv6Header.YFilter
    srv6Header.EntityData.YangName = "srv6-header"
    srv6Header.EntityData.BundleName = "cisco_ios_xr"
    srv6Header.EntityData.ParentYangName = "probe"
    srv6Header.EntityData.SegmentPath = "srv6-header"
    srv6Header.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv4/hops/hop/probes/probe/" + srv6Header.EntityData.SegmentPath
    srv6Header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srv6Header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srv6Header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srv6Header.EntityData.Children = types.NewOrderedMap()
    srv6Header.EntityData.Children.Append("segments", types.YChild{"Segments", &srv6Header.Segments})
    srv6Header.EntityData.Leafs = types.NewOrderedMap()
    srv6Header.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", srv6Header.DestinationAddress})
    srv6Header.EntityData.Leafs.Append("segments-left", types.YLeaf{"SegmentsLeft", srv6Header.SegmentsLeft})

    srv6Header.EntityData.YListKeys = []string {}

    return &(srv6Header.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header_Segments
type Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sid in sidlist. The type is slice of string.
    Segment []interface{}
}

func (segments *Traceroute_Output_TracerouteResponse_Ipv4_Hops_Hop_Probes_Probe_Srv6Header_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "srv6-header"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv4/hops/hop/probes/probe/srv6-header/" + segments.EntityData.SegmentPath
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Leafs = types.NewOrderedMap()
    segments.EntityData.Leafs.Append("segment", types.YLeaf{"Segment", segments.Segment})

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6
// This type is a presence type.
type Traceroute_Output_TracerouteResponse_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("hops", types.YChild{"Hops", &ipv6.Hops})
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv6.Destination})
    ipv6.EntityData.Leafs.Append("verbose-output", types.YLeaf{"VerboseOutput", ipv6.VerboseOutput})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops
type Traceroute_Output_TracerouteResponse_Ipv6_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop.
    Hop []*Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop
}

func (hops *Traceroute_Output_TracerouteResponse_Ipv6_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "ipv6"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv6/" + hops.EntityData.SegmentPath
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = types.NewOrderedMap()
    hops.EntityData.Children.Append("hop", types.YChild{"Hop", nil})
    for i := range hops.Hop {
        hops.EntityData.Children.Append(types.GetSegmentPath(hops.Hop[i]), types.YChild{"Hop", hops.Hop[i]})
    }
    hops.EntityData.Leafs = types.NewOrderedMap()

    hops.EntityData.YListKeys = []string {}

    return &(hops.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    hop.EntityData.SegmentPath = "hop" + types.AddKeyToken(hop.HopIndex, "hop-index")
    hop.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv6/hops/" + hop.EntityData.SegmentPath
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = types.NewOrderedMap()
    hop.EntityData.Children.Append("probes", types.YChild{"Probes", &hop.Probes})
    hop.EntityData.Leafs = types.NewOrderedMap()
    hop.EntityData.Leafs.Append("hop-index", types.YLeaf{"HopIndex", hop.HopIndex})
    hop.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", hop.HopAddress})
    hop.EntityData.Leafs.Append("hop-hostname", types.YLeaf{"HopHostname", hop.HopHostname})

    hop.EntityData.YListKeys = []string {"HopIndex"}

    return &(hop.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe.
    Probe []*Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe
}

func (probes *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes) GetEntityData() *types.CommonEntityData {
    probes.EntityData.YFilter = probes.YFilter
    probes.EntityData.YangName = "probes"
    probes.EntityData.BundleName = "cisco_ios_xr"
    probes.EntityData.ParentYangName = "hop"
    probes.EntityData.SegmentPath = "probes"
    probes.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv6/hops/hop/" + probes.EntityData.SegmentPath
    probes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probes.EntityData.Children = types.NewOrderedMap()
    probes.EntityData.Children.Append("probe", types.YChild{"Probe", nil})
    for i := range probes.Probe {
        probes.EntityData.Children.Append(types.GetSegmentPath(probes.Probe[i]), types.YChild{"Probe", probes.Probe[i]})
    }
    probes.EntityData.Leafs = types.NewOrderedMap()

    probes.EntityData.YListKeys = []string {}

    return &(probes.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    
    Srv6Header Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header
}

func (probe *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe) GetEntityData() *types.CommonEntityData {
    probe.EntityData.YFilter = probe.YFilter
    probe.EntityData.YangName = "probe"
    probe.EntityData.BundleName = "cisco_ios_xr"
    probe.EntityData.ParentYangName = "probes"
    probe.EntityData.SegmentPath = "probe" + types.AddKeyToken(probe.ProbeIndex, "probe-index")
    probe.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv6/hops/hop/probes/" + probe.EntityData.SegmentPath
    probe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probe.EntityData.Children = types.NewOrderedMap()
    probe.EntityData.Children.Append("srv6-header", types.YChild{"Srv6Header", &probe.Srv6Header})
    probe.EntityData.Leafs = types.NewOrderedMap()
    probe.EntityData.Leafs.Append("probe-index", types.YLeaf{"ProbeIndex", probe.ProbeIndex})
    probe.EntityData.Leafs.Append("result", types.YLeaf{"Result", probe.Result})
    probe.EntityData.Leafs.Append("delta-time", types.YLeaf{"DeltaTime", probe.DeltaTime})
    probe.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", probe.HopAddress})
    probe.EntityData.Leafs.Append("hop-hostname", types.YLeaf{"HopHostname", probe.HopHostname})

    probe.EntityData.YListKeys = []string {"ProbeIndex"}

    return &(probe.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination address for srv6 header. The type is interface{} with range:
    // 0..4294967295.
    DestinationAddress interface{}

    // Number of segments left. The type is interface{} with range: 0..4294967295.
    SegmentsLeft interface{}

    
    Segments Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header_Segments
}

func (srv6Header *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header) GetEntityData() *types.CommonEntityData {
    srv6Header.EntityData.YFilter = srv6Header.YFilter
    srv6Header.EntityData.YangName = "srv6-header"
    srv6Header.EntityData.BundleName = "cisco_ios_xr"
    srv6Header.EntityData.ParentYangName = "probe"
    srv6Header.EntityData.SegmentPath = "srv6-header"
    srv6Header.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv6/hops/hop/probes/probe/" + srv6Header.EntityData.SegmentPath
    srv6Header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srv6Header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srv6Header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srv6Header.EntityData.Children = types.NewOrderedMap()
    srv6Header.EntityData.Children.Append("segments", types.YChild{"Segments", &srv6Header.Segments})
    srv6Header.EntityData.Leafs = types.NewOrderedMap()
    srv6Header.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", srv6Header.DestinationAddress})
    srv6Header.EntityData.Leafs.Append("segments-left", types.YLeaf{"SegmentsLeft", srv6Header.SegmentsLeft})

    srv6Header.EntityData.YListKeys = []string {}

    return &(srv6Header.EntityData)
}

// Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header_Segments
type Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sid in sidlist. The type is slice of string.
    Segment []interface{}
}

func (segments *Traceroute_Output_TracerouteResponse_Ipv6_Hops_Hop_Probes_Probe_Srv6Header_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "srv6-header"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.AbsolutePath = "Cisco-IOS-XR-traceroute-act:traceroute/output/traceroute-response/ipv6/hops/hop/probes/probe/srv6-header/" + segments.EntityData.SegmentPath
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Leafs = types.NewOrderedMap()
    segments.EntityData.Leafs.Append("segment", types.YLeaf{"Segment", segments.Segment})

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

