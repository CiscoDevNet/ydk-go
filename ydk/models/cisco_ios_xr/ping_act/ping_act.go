// This module contains a collection of YANG definitions
// for Cisco IOS-XR ping action package configuration.
// 
// Copyright (c) 2016, 2018 by Cisco Systems, Inc.
// All rights reserved.
package ping_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ping_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ping-act ping}", reflect.TypeOf(Ping{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ping-act:ping", reflect.TypeOf(Ping{}))
}

// Ping
// Send echo messages
type Ping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Ping_Input

    
    Output Ping_Output
}

func (ping *Ping) GetEntityData() *types.CommonEntityData {
    ping.EntityData.YFilter = ping.YFilter
    ping.EntityData.YangName = "ping"
    ping.EntityData.BundleName = "cisco_ios_xr"
    ping.EntityData.ParentYangName = "Cisco-IOS-XR-ping-act"
    ping.EntityData.SegmentPath = "Cisco-IOS-XR-ping-act:ping"
    ping.EntityData.AbsolutePath = ping.EntityData.SegmentPath
    ping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ping.EntityData.Children = types.NewOrderedMap()
    ping.EntityData.Children.Append("input", types.YChild{"Input", &ping.Input})
    ping.EntityData.Children.Append("output", types.YChild{"Output", &ping.Output})
    ping.EntityData.Leafs = types.NewOrderedMap()

    ping.EntityData.YListKeys = []string {}

    return &(ping.EntityData)
}

// Ping_Input
type Ping_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Destination Ping_Input_Destination

    // The type is slice of Ping_Input_Ipv4.
    Ipv4 []*Ping_Input_Ipv4

    
    Ipv6 Ping_Input_Ipv6
}

func (input *Ping_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "ping"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("destination", types.YChild{"Destination", &input.Destination})
    input.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", nil})
    for i := range input.Ipv4 {
        input.EntityData.Children.Append(types.GetSegmentPath(input.Ipv4[i]), types.YChild{"Ipv4", input.Ipv4[i]})
    }
    input.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &input.Ipv6})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Ping_Input_Destination
// This type is a presence type.
type Ping_Input_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Ping destination address or hostname. The type is string. This attribute is
    // mandatory.
    Destination interface{}

    // Number of ping packets to be sent out. The type is interface{} with range:
    // 1..64. The default value is 5.
    RepeatCount interface{}

    // Size of ping packet. The type is interface{} with range: 36..18024. The
    // default value is 100.
    DataSize interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 2.
    Timeout interface{}

    // Ping interval in milli seconds. The type is interface{} with range:
    // 0..3600. The default value is 10.
    Interval interface{}

    // Pattern of payload data. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Pattern interface{}

    // Sweep is enabled. The type is bool.
    Sweep interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Source address or interface. The type is string.
    Source interface{}

    // Validate return packet. The type is bool.
    Verbose interface{}

    // Type of Service. The type is interface{} with range: 0..255.
    TypeOfService interface{}

    // Do Not Fragment. The type is bool.
    DoNotFrag interface{}

    // Validate return packet. The type is bool.
    Validate interface{}

    // Priority of the packet. The type is interface{} with range: 0..15.
    Priority interface{}

    // Outgoing interface, needed in case of ping to link local address. The type
    // is string.
    OutgoingInterface interface{}
}

func (destination *Ping_Input_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "input"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/input/" + destination.EntityData.SegmentPath
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = types.NewOrderedMap()
    destination.EntityData.Leafs = types.NewOrderedMap()
    destination.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", destination.Destination})
    destination.EntityData.Leafs.Append("repeat-count", types.YLeaf{"RepeatCount", destination.RepeatCount})
    destination.EntityData.Leafs.Append("data-size", types.YLeaf{"DataSize", destination.DataSize})
    destination.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", destination.Timeout})
    destination.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", destination.Interval})
    destination.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", destination.Pattern})
    destination.EntityData.Leafs.Append("sweep", types.YLeaf{"Sweep", destination.Sweep})
    destination.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", destination.VrfName})
    destination.EntityData.Leafs.Append("source", types.YLeaf{"Source", destination.Source})
    destination.EntityData.Leafs.Append("verbose", types.YLeaf{"Verbose", destination.Verbose})
    destination.EntityData.Leafs.Append("type-of-service", types.YLeaf{"TypeOfService", destination.TypeOfService})
    destination.EntityData.Leafs.Append("do-not-frag", types.YLeaf{"DoNotFrag", destination.DoNotFrag})
    destination.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", destination.Validate})
    destination.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", destination.Priority})
    destination.EntityData.Leafs.Append("outgoing-interface", types.YLeaf{"OutgoingInterface", destination.OutgoingInterface})

    destination.EntityData.YListKeys = []string {}

    return &(destination.EntityData)
}

// Ping_Input_Ipv4
type Ping_Input_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Ping destination address or hostname. The type is
    // string. This attribute is mandatory.
    Destination interface{}

    // Number of ping packets to be sent out. The type is interface{} with range:
    // 1..64. The default value is 5.
    RepeatCount interface{}

    // Size of ping packet. The type is interface{} with range: 36..18024. The
    // default value is 100.
    DataSize interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 2.
    Timeout interface{}

    // Ping interval in milli seconds. The type is interface{} with range:
    // 0..3600. The default value is 10.
    Interval interface{}

    // Pattern of payload data. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Pattern interface{}

    // Sweep is enabled. The type is bool.
    Sweep interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Source address or interface. The type is string.
    Source interface{}

    // Validate return packet. The type is bool.
    Verbose interface{}

    // Type of Service. The type is interface{} with range: 0..255.
    TypeOfService interface{}

    // Do Not Fragment. The type is bool.
    DoNotFrag interface{}

    // Validate return packet. The type is bool.
    Validate interface{}
}

func (ipv4 *Ping_Input_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "input"
    ipv4.EntityData.SegmentPath = "ipv4" + types.AddKeyToken(ipv4.Destination, "destination")
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/input/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv4.Destination})
    ipv4.EntityData.Leafs.Append("repeat-count", types.YLeaf{"RepeatCount", ipv4.RepeatCount})
    ipv4.EntityData.Leafs.Append("data-size", types.YLeaf{"DataSize", ipv4.DataSize})
    ipv4.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ipv4.Timeout})
    ipv4.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ipv4.Interval})
    ipv4.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", ipv4.Pattern})
    ipv4.EntityData.Leafs.Append("sweep", types.YLeaf{"Sweep", ipv4.Sweep})
    ipv4.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4.VrfName})
    ipv4.EntityData.Leafs.Append("source", types.YLeaf{"Source", ipv4.Source})
    ipv4.EntityData.Leafs.Append("verbose", types.YLeaf{"Verbose", ipv4.Verbose})
    ipv4.EntityData.Leafs.Append("type-of-service", types.YLeaf{"TypeOfService", ipv4.TypeOfService})
    ipv4.EntityData.Leafs.Append("do-not-frag", types.YLeaf{"DoNotFrag", ipv4.DoNotFrag})
    ipv4.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", ipv4.Validate})

    ipv4.EntityData.YListKeys = []string {"Destination"}

    return &(ipv4.EntityData)
}

// Ping_Input_Ipv6
// This type is a presence type.
type Ping_Input_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Ping destination address or hostname. The type is string. This attribute is
    // mandatory.
    Destination interface{}

    // Number of ping packets to be sent out. The type is interface{} with range:
    // 1..64. The default value is 5.
    RepeatCount interface{}

    // Size of ping packet. The type is interface{} with range: 36..18024. The
    // default value is 100.
    DataSize interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 2.
    Timeout interface{}

    // Ping interval in milli seconds. The type is interface{} with range:
    // 0..3600. The default value is 10.
    Interval interface{}

    // Pattern of payload data. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Pattern interface{}

    // Sweep is enabled. The type is bool.
    Sweep interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Source address or interface. The type is string.
    Source interface{}

    // Validate return packet. The type is bool.
    Verbose interface{}

    // Priority of the packet. The type is interface{} with range: 0..15.
    Priority interface{}

    // Outgoing interface, needed in case of ping to link local address. The type
    // is string.
    OutgoingInterface interface{}
}

func (ipv6 *Ping_Input_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "input"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/input/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv6.Destination})
    ipv6.EntityData.Leafs.Append("repeat-count", types.YLeaf{"RepeatCount", ipv6.RepeatCount})
    ipv6.EntityData.Leafs.Append("data-size", types.YLeaf{"DataSize", ipv6.DataSize})
    ipv6.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ipv6.Timeout})
    ipv6.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ipv6.Interval})
    ipv6.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", ipv6.Pattern})
    ipv6.EntityData.Leafs.Append("sweep", types.YLeaf{"Sweep", ipv6.Sweep})
    ipv6.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6.VrfName})
    ipv6.EntityData.Leafs.Append("source", types.YLeaf{"Source", ipv6.Source})
    ipv6.EntityData.Leafs.Append("verbose", types.YLeaf{"Verbose", ipv6.Verbose})
    ipv6.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ipv6.Priority})
    ipv6.EntityData.Leafs.Append("outgoing-interface", types.YLeaf{"OutgoingInterface", ipv6.OutgoingInterface})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Ping_Output
type Ping_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    PingResponse Ping_Output_PingResponse
}

func (output *Ping_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "ping"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("ping-response", types.YChild{"PingResponse", &output.PingResponse})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Ping_Output_PingResponse
type Ping_Output_PingResponse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Ping_Output_PingResponse_Ipv4.
    Ipv4 []*Ping_Output_PingResponse_Ipv4

    
    Ipv6 Ping_Output_PingResponse_Ipv6
}

func (pingResponse *Ping_Output_PingResponse) GetEntityData() *types.CommonEntityData {
    pingResponse.EntityData.YFilter = pingResponse.YFilter
    pingResponse.EntityData.YangName = "ping-response"
    pingResponse.EntityData.BundleName = "cisco_ios_xr"
    pingResponse.EntityData.ParentYangName = "output"
    pingResponse.EntityData.SegmentPath = "ping-response"
    pingResponse.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/" + pingResponse.EntityData.SegmentPath
    pingResponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pingResponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pingResponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pingResponse.EntityData.Children = types.NewOrderedMap()
    pingResponse.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", nil})
    for i := range pingResponse.Ipv4 {
        pingResponse.EntityData.Children.Append(types.GetSegmentPath(pingResponse.Ipv4[i]), types.YChild{"Ipv4", pingResponse.Ipv4[i]})
    }
    pingResponse.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &pingResponse.Ipv6})
    pingResponse.EntityData.Leafs = types.NewOrderedMap()

    pingResponse.EntityData.YListKeys = []string {}

    return &(pingResponse.EntityData)
}

// Ping_Output_PingResponse_Ipv4
type Ping_Output_PingResponse_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Ping destination address or hostname. The type is
    // string. This attribute is mandatory.
    Destination interface{}

    // Number of ping packets to be sent out. The type is interface{} with range:
    // 1..64. The default value is 5.
    RepeatCount interface{}

    // Size of ping packet. The type is interface{} with range: 36..18024. The
    // default value is 100.
    DataSize interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 2.
    Timeout interface{}

    // Ping interval in milli seconds. The type is interface{} with range:
    // 0..3600. The default value is 10.
    Interval interface{}

    // Pattern of payload data. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Pattern interface{}

    // Sweep is enabled. The type is bool.
    Sweep interface{}

    // Number of packets reach to destination and get reply back. The type is
    // interface{} with range: 0..18446744073709551615.
    Hits interface{}

    // Total number of packets sent out. The type is interface{} with range:
    // 0..18446744073709551615.
    Total interface{}

    // Successful rate of ping. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessRate interface{}

    // Minimum value of Round Trip Time. The type is interface{} with range:
    // 0..18446744073709551615.
    RttMin interface{}

    // Average value of Round Trip Time. The type is interface{} with range:
    // 0..18446744073709551615.
    RttAvg interface{}

    // Maximum value of Round Trip Time. The type is interface{} with range:
    // 0..18446744073709551615.
    RttMax interface{}

    // Minimum value of sweep size. The type is interface{} with range:
    // 0..4294967295.
    SweepMin interface{}

    // Maximum value of sweep size. The type is interface{} with range:
    // 0..18446744073709551615.
    SweepMax interface{}

    // Rotate Pattern is enabled. The type is bool.
    RotatePattern interface{}

    // Error response for each ping, in case of bulk ping. The type is string.
    PingErrorResponse interface{}

    
    Replies Ping_Output_PingResponse_Ipv4_Replies
}

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "ping-response"
    ipv4.EntityData.SegmentPath = "ipv4" + types.AddKeyToken(ipv4.Destination, "destination")
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("replies", types.YChild{"Replies", &ipv4.Replies})
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv4.Destination})
    ipv4.EntityData.Leafs.Append("repeat-count", types.YLeaf{"RepeatCount", ipv4.RepeatCount})
    ipv4.EntityData.Leafs.Append("data-size", types.YLeaf{"DataSize", ipv4.DataSize})
    ipv4.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ipv4.Timeout})
    ipv4.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ipv4.Interval})
    ipv4.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", ipv4.Pattern})
    ipv4.EntityData.Leafs.Append("sweep", types.YLeaf{"Sweep", ipv4.Sweep})
    ipv4.EntityData.Leafs.Append("hits", types.YLeaf{"Hits", ipv4.Hits})
    ipv4.EntityData.Leafs.Append("total", types.YLeaf{"Total", ipv4.Total})
    ipv4.EntityData.Leafs.Append("success-rate", types.YLeaf{"SuccessRate", ipv4.SuccessRate})
    ipv4.EntityData.Leafs.Append("rtt-min", types.YLeaf{"RttMin", ipv4.RttMin})
    ipv4.EntityData.Leafs.Append("rtt-avg", types.YLeaf{"RttAvg", ipv4.RttAvg})
    ipv4.EntityData.Leafs.Append("rtt-max", types.YLeaf{"RttMax", ipv4.RttMax})
    ipv4.EntityData.Leafs.Append("sweep-min", types.YLeaf{"SweepMin", ipv4.SweepMin})
    ipv4.EntityData.Leafs.Append("sweep-max", types.YLeaf{"SweepMax", ipv4.SweepMax})
    ipv4.EntityData.Leafs.Append("rotate-pattern", types.YLeaf{"RotatePattern", ipv4.RotatePattern})
    ipv4.EntityData.Leafs.Append("ping-error-response", types.YLeaf{"PingErrorResponse", ipv4.PingErrorResponse})

    ipv4.EntityData.YListKeys = []string {"Destination"}

    return &(ipv4.EntityData)
}

// Ping_Output_PingResponse_Ipv4_Replies
type Ping_Output_PingResponse_Ipv4_Replies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Ping_Output_PingResponse_Ipv4_Replies_Reply.
    Reply []*Ping_Output_PingResponse_Ipv4_Replies_Reply
}

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetEntityData() *types.CommonEntityData {
    replies.EntityData.YFilter = replies.YFilter
    replies.EntityData.YangName = "replies"
    replies.EntityData.BundleName = "cisco_ios_xr"
    replies.EntityData.ParentYangName = "ipv4"
    replies.EntityData.SegmentPath = "replies"
    replies.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/ipv4/" + replies.EntityData.SegmentPath
    replies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replies.EntityData.Children = types.NewOrderedMap()
    replies.EntityData.Children.Append("reply", types.YChild{"Reply", nil})
    for i := range replies.Reply {
        replies.EntityData.Children.Append(types.GetSegmentPath(replies.Reply[i]), types.YChild{"Reply", replies.Reply[i]})
    }
    replies.EntityData.Leafs = types.NewOrderedMap()

    replies.EntityData.YListKeys = []string {}

    return &(replies.EntityData)
}

// Ping_Output_PingResponse_Ipv4_Replies_Reply
type Ping_Output_PingResponse_Ipv4_Replies_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the reply list. The type is interface{}
    // with range: 1..2147483647.
    ReplyIndex interface{}

    // Response for each packet. The type is string.
    Result interface{}

    
    BroadcastReplyAddresses Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses
}

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "replies"
    reply.EntityData.SegmentPath = "reply" + types.AddKeyToken(reply.ReplyIndex, "reply-index")
    reply.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/ipv4/replies/" + reply.EntityData.SegmentPath
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = types.NewOrderedMap()
    reply.EntityData.Children.Append("broadcast-reply-addresses", types.YChild{"BroadcastReplyAddresses", &reply.BroadcastReplyAddresses})
    reply.EntityData.Leafs = types.NewOrderedMap()
    reply.EntityData.Leafs.Append("reply-index", types.YLeaf{"ReplyIndex", reply.ReplyIndex})
    reply.EntityData.Leafs.Append("result", types.YLeaf{"Result", reply.Result})

    reply.EntityData.YListKeys = []string {"ReplyIndex"}

    return &(reply.EntityData)
}

// Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses
type Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress.
    BroadcastReplyAddress []*Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress
}

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetEntityData() *types.CommonEntityData {
    broadcastReplyAddresses.EntityData.YFilter = broadcastReplyAddresses.YFilter
    broadcastReplyAddresses.EntityData.YangName = "broadcast-reply-addresses"
    broadcastReplyAddresses.EntityData.BundleName = "cisco_ios_xr"
    broadcastReplyAddresses.EntityData.ParentYangName = "reply"
    broadcastReplyAddresses.EntityData.SegmentPath = "broadcast-reply-addresses"
    broadcastReplyAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/ipv4/replies/reply/" + broadcastReplyAddresses.EntityData.SegmentPath
    broadcastReplyAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    broadcastReplyAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    broadcastReplyAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    broadcastReplyAddresses.EntityData.Children = types.NewOrderedMap()
    broadcastReplyAddresses.EntityData.Children.Append("broadcast-reply-address", types.YChild{"BroadcastReplyAddress", nil})
    for i := range broadcastReplyAddresses.BroadcastReplyAddress {
        broadcastReplyAddresses.EntityData.Children.Append(types.GetSegmentPath(broadcastReplyAddresses.BroadcastReplyAddress[i]), types.YChild{"BroadcastReplyAddress", broadcastReplyAddresses.BroadcastReplyAddress[i]})
    }
    broadcastReplyAddresses.EntityData.Leafs = types.NewOrderedMap()

    broadcastReplyAddresses.EntityData.YListKeys = []string {}

    return &(broadcastReplyAddresses.EntityData)
}

// Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress
type Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Broadcast reply address. The type is string.
    ReplyAddress interface{}

    // Sign for each reply packet. The type is string.
    Result interface{}
}

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetEntityData() *types.CommonEntityData {
    broadcastReplyAddress.EntityData.YFilter = broadcastReplyAddress.YFilter
    broadcastReplyAddress.EntityData.YangName = "broadcast-reply-address"
    broadcastReplyAddress.EntityData.BundleName = "cisco_ios_xr"
    broadcastReplyAddress.EntityData.ParentYangName = "broadcast-reply-addresses"
    broadcastReplyAddress.EntityData.SegmentPath = "broadcast-reply-address" + types.AddKeyToken(broadcastReplyAddress.ReplyAddress, "reply-address")
    broadcastReplyAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/ipv4/replies/reply/broadcast-reply-addresses/" + broadcastReplyAddress.EntityData.SegmentPath
    broadcastReplyAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    broadcastReplyAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    broadcastReplyAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    broadcastReplyAddress.EntityData.Children = types.NewOrderedMap()
    broadcastReplyAddress.EntityData.Leafs = types.NewOrderedMap()
    broadcastReplyAddress.EntityData.Leafs.Append("reply-address", types.YLeaf{"ReplyAddress", broadcastReplyAddress.ReplyAddress})
    broadcastReplyAddress.EntityData.Leafs.Append("result", types.YLeaf{"Result", broadcastReplyAddress.Result})

    broadcastReplyAddress.EntityData.YListKeys = []string {"ReplyAddress"}

    return &(broadcastReplyAddress.EntityData)
}

// Ping_Output_PingResponse_Ipv6
// This type is a presence type.
type Ping_Output_PingResponse_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Ping destination address or hostname. The type is string. This attribute is
    // mandatory.
    Destination interface{}

    // Number of ping packets to be sent out. The type is interface{} with range:
    // 1..64. The default value is 5.
    RepeatCount interface{}

    // Size of ping packet. The type is interface{} with range: 36..18024. The
    // default value is 100.
    DataSize interface{}

    // Timeout in seconds. The type is interface{} with range: 0..36. The default
    // value is 2.
    Timeout interface{}

    // Ping interval in milli seconds. The type is interface{} with range:
    // 0..3600. The default value is 10.
    Interval interface{}

    // Pattern of payload data. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Pattern interface{}

    // Sweep is enabled. The type is bool.
    Sweep interface{}

    // Minimum value of sweep size. The type is interface{} with range:
    // 0..4294967295.
    SweepMin interface{}

    // Maximum value of sweep size. The type is interface{} with range:
    // 0..18446744073709551615.
    SweepMax interface{}

    // Rotate Pattern is enabled. The type is bool.
    RotatePattern interface{}

    // Number of packets reach to destination and get reply back. The type is
    // interface{} with range: 0..18446744073709551615.
    Hits interface{}

    // Total number of packets sent out. The type is interface{} with range:
    // 0..18446744073709551615.
    Total interface{}

    // Successful rate of ping. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessRate interface{}

    // Minimum value of Round Trip Time. The type is interface{} with range:
    // 0..18446744073709551615.
    RttMin interface{}

    // Average value of Round Trip Time. The type is interface{} with range:
    // 0..18446744073709551615.
    RttAvg interface{}

    // Maximum value of Round Trip Time. The type is interface{} with range:
    // 0..18446744073709551615.
    RttMax interface{}

    
    Replies Ping_Output_PingResponse_Ipv6_Replies
}

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "ping-response"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("replies", types.YChild{"Replies", &ipv6.Replies})
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", ipv6.Destination})
    ipv6.EntityData.Leafs.Append("repeat-count", types.YLeaf{"RepeatCount", ipv6.RepeatCount})
    ipv6.EntityData.Leafs.Append("data-size", types.YLeaf{"DataSize", ipv6.DataSize})
    ipv6.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ipv6.Timeout})
    ipv6.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ipv6.Interval})
    ipv6.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", ipv6.Pattern})
    ipv6.EntityData.Leafs.Append("sweep", types.YLeaf{"Sweep", ipv6.Sweep})
    ipv6.EntityData.Leafs.Append("sweep-min", types.YLeaf{"SweepMin", ipv6.SweepMin})
    ipv6.EntityData.Leafs.Append("sweep-max", types.YLeaf{"SweepMax", ipv6.SweepMax})
    ipv6.EntityData.Leafs.Append("rotate-pattern", types.YLeaf{"RotatePattern", ipv6.RotatePattern})
    ipv6.EntityData.Leafs.Append("hits", types.YLeaf{"Hits", ipv6.Hits})
    ipv6.EntityData.Leafs.Append("total", types.YLeaf{"Total", ipv6.Total})
    ipv6.EntityData.Leafs.Append("success-rate", types.YLeaf{"SuccessRate", ipv6.SuccessRate})
    ipv6.EntityData.Leafs.Append("rtt-min", types.YLeaf{"RttMin", ipv6.RttMin})
    ipv6.EntityData.Leafs.Append("rtt-avg", types.YLeaf{"RttAvg", ipv6.RttAvg})
    ipv6.EntityData.Leafs.Append("rtt-max", types.YLeaf{"RttMax", ipv6.RttMax})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Ping_Output_PingResponse_Ipv6_Replies
type Ping_Output_PingResponse_Ipv6_Replies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Ping_Output_PingResponse_Ipv6_Replies_Reply.
    Reply []*Ping_Output_PingResponse_Ipv6_Replies_Reply
}

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetEntityData() *types.CommonEntityData {
    replies.EntityData.YFilter = replies.YFilter
    replies.EntityData.YangName = "replies"
    replies.EntityData.BundleName = "cisco_ios_xr"
    replies.EntityData.ParentYangName = "ipv6"
    replies.EntityData.SegmentPath = "replies"
    replies.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/ipv6/" + replies.EntityData.SegmentPath
    replies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replies.EntityData.Children = types.NewOrderedMap()
    replies.EntityData.Children.Append("reply", types.YChild{"Reply", nil})
    for i := range replies.Reply {
        replies.EntityData.Children.Append(types.GetSegmentPath(replies.Reply[i]), types.YChild{"Reply", replies.Reply[i]})
    }
    replies.EntityData.Leafs = types.NewOrderedMap()

    replies.EntityData.YListKeys = []string {}

    return &(replies.EntityData)
}

// Ping_Output_PingResponse_Ipv6_Replies_Reply
type Ping_Output_PingResponse_Ipv6_Replies_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index of the reply list. The type is interface{}
    // with range: 1..2147483647.
    ReplyIndex interface{}

    // Response for each packet. The type is string.
    Result interface{}
}

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "replies"
    reply.EntityData.SegmentPath = "reply" + types.AddKeyToken(reply.ReplyIndex, "reply-index")
    reply.EntityData.AbsolutePath = "Cisco-IOS-XR-ping-act:ping/output/ping-response/ipv6/replies/" + reply.EntityData.SegmentPath
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = types.NewOrderedMap()
    reply.EntityData.Leafs = types.NewOrderedMap()
    reply.EntityData.Leafs.Append("reply-index", types.YLeaf{"ReplyIndex", reply.ReplyIndex})
    reply.EntityData.Leafs.Append("result", types.YLeaf{"Result", reply.Result})

    reply.EntityData.YListKeys = []string {"ReplyIndex"}

    return &(reply.EntityData)
}

