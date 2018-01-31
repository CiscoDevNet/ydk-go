// This module contains a collection of YANG definitions
// for Cisco IOS-XR ping action package configuration.
// 
// Copyright (c) 2016 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Ping_Input

    
    Output Ping_Output
}

func (ping *Ping) GetFilter() yfilter.YFilter { return ping.YFilter }

func (ping *Ping) SetFilter(yf yfilter.YFilter) { ping.YFilter = yf }

func (ping *Ping) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (ping *Ping) GetSegmentPath() string {
    return "Cisco-IOS-XR-ping-act:ping"
}

func (ping *Ping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &ping.Input
    }
    if childYangName == "output" {
        return &ping.Output
    }
    return nil
}

func (ping *Ping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &ping.Input
    children["output"] = &ping.Output
    return children
}

func (ping *Ping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ping *Ping) GetBundleName() string { return "cisco_ios_xr" }

func (ping *Ping) GetYangName() string { return "ping" }

func (ping *Ping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ping *Ping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ping *Ping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ping *Ping) SetParent(parent types.Entity) { ping.parent = parent }

func (ping *Ping) GetParent() types.Entity { return ping.parent }

func (ping *Ping) GetParentYangName() string { return "Cisco-IOS-XR-ping-act" }

// Ping_Input
type Ping_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Destination Ping_Input_Destination

    // The type is slice of Ping_Input_Ipv4.
    Ipv4 []Ping_Input_Ipv4

    
    Ipv6 Ping_Input_Ipv6
}

func (input *Ping_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Ping_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Ping_Input) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (input *Ping_Input) GetSegmentPath() string {
    return "input"
}

func (input *Ping_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        return &input.Destination
    }
    if childYangName == "ipv4" {
        for _, c := range input.Ipv4 {
            if input.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ping_Input_Ipv4{}
        input.Ipv4 = append(input.Ipv4, child)
        return &input.Ipv4[len(input.Ipv4)-1]
    }
    if childYangName == "ipv6" {
        return &input.Ipv6
    }
    return nil
}

func (input *Ping_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination"] = &input.Destination
    for i := range input.Ipv4 {
        children[input.Ipv4[i].GetSegmentPath()] = &input.Ipv4[i]
    }
    children["ipv6"] = &input.Ipv6
    return children
}

func (input *Ping_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *Ping_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Ping_Input) GetYangName() string { return "input" }

func (input *Ping_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Ping_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Ping_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Ping_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Ping_Input) GetParent() types.Entity { return input.parent }

func (input *Ping_Input) GetParentYangName() string { return "ping" }

// Ping_Input_Destination
type Ping_Input_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (destination *Ping_Input_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *Ping_Input_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *Ping_Input_Destination) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "repeat-count" { return "RepeatCount" }
    if yname == "data-size" { return "DataSize" }
    if yname == "timeout" { return "Timeout" }
    if yname == "interval" { return "Interval" }
    if yname == "pattern" { return "Pattern" }
    if yname == "sweep" { return "Sweep" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source" { return "Source" }
    if yname == "verbose" { return "Verbose" }
    if yname == "type-of-service" { return "TypeOfService" }
    if yname == "do-not-frag" { return "DoNotFrag" }
    if yname == "validate" { return "Validate" }
    if yname == "priority" { return "Priority" }
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    return ""
}

func (destination *Ping_Input_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *Ping_Input_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destination *Ping_Input_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destination *Ping_Input_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = destination.Destination
    leafs["repeat-count"] = destination.RepeatCount
    leafs["data-size"] = destination.DataSize
    leafs["timeout"] = destination.Timeout
    leafs["interval"] = destination.Interval
    leafs["pattern"] = destination.Pattern
    leafs["sweep"] = destination.Sweep
    leafs["vrf-name"] = destination.VrfName
    leafs["source"] = destination.Source
    leafs["verbose"] = destination.Verbose
    leafs["type-of-service"] = destination.TypeOfService
    leafs["do-not-frag"] = destination.DoNotFrag
    leafs["validate"] = destination.Validate
    leafs["priority"] = destination.Priority
    leafs["outgoing-interface"] = destination.OutgoingInterface
    return leafs
}

func (destination *Ping_Input_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *Ping_Input_Destination) GetYangName() string { return "destination" }

func (destination *Ping_Input_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *Ping_Input_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *Ping_Input_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *Ping_Input_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *Ping_Input_Destination) GetParent() types.Entity { return destination.parent }

func (destination *Ping_Input_Destination) GetParentYangName() string { return "input" }

// Ping_Input_Ipv4
type Ping_Input_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ipv4 *Ping_Input_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Ping_Input_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Ping_Input_Ipv4) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "repeat-count" { return "RepeatCount" }
    if yname == "data-size" { return "DataSize" }
    if yname == "timeout" { return "Timeout" }
    if yname == "interval" { return "Interval" }
    if yname == "pattern" { return "Pattern" }
    if yname == "sweep" { return "Sweep" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source" { return "Source" }
    if yname == "verbose" { return "Verbose" }
    if yname == "type-of-service" { return "TypeOfService" }
    if yname == "do-not-frag" { return "DoNotFrag" }
    if yname == "validate" { return "Validate" }
    return ""
}

func (ipv4 *Ping_Input_Ipv4) GetSegmentPath() string {
    return "ipv4" + "[destination='" + fmt.Sprintf("%v", ipv4.Destination) + "']"
}

func (ipv4 *Ping_Input_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *Ping_Input_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *Ping_Input_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv4.Destination
    leafs["repeat-count"] = ipv4.RepeatCount
    leafs["data-size"] = ipv4.DataSize
    leafs["timeout"] = ipv4.Timeout
    leafs["interval"] = ipv4.Interval
    leafs["pattern"] = ipv4.Pattern
    leafs["sweep"] = ipv4.Sweep
    leafs["vrf-name"] = ipv4.VrfName
    leafs["source"] = ipv4.Source
    leafs["verbose"] = ipv4.Verbose
    leafs["type-of-service"] = ipv4.TypeOfService
    leafs["do-not-frag"] = ipv4.DoNotFrag
    leafs["validate"] = ipv4.Validate
    return leafs
}

func (ipv4 *Ping_Input_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Ping_Input_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Ping_Input_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Ping_Input_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Ping_Input_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Ping_Input_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Ping_Input_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Ping_Input_Ipv4) GetParentYangName() string { return "input" }

// Ping_Input_Ipv6
type Ping_Input_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ipv6 *Ping_Input_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Ping_Input_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Ping_Input_Ipv6) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "repeat-count" { return "RepeatCount" }
    if yname == "data-size" { return "DataSize" }
    if yname == "timeout" { return "Timeout" }
    if yname == "interval" { return "Interval" }
    if yname == "pattern" { return "Pattern" }
    if yname == "sweep" { return "Sweep" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source" { return "Source" }
    if yname == "verbose" { return "Verbose" }
    if yname == "priority" { return "Priority" }
    if yname == "outgoing-interface" { return "OutgoingInterface" }
    return ""
}

func (ipv6 *Ping_Input_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Ping_Input_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *Ping_Input_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *Ping_Input_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv6.Destination
    leafs["repeat-count"] = ipv6.RepeatCount
    leafs["data-size"] = ipv6.DataSize
    leafs["timeout"] = ipv6.Timeout
    leafs["interval"] = ipv6.Interval
    leafs["pattern"] = ipv6.Pattern
    leafs["sweep"] = ipv6.Sweep
    leafs["vrf-name"] = ipv6.VrfName
    leafs["source"] = ipv6.Source
    leafs["verbose"] = ipv6.Verbose
    leafs["priority"] = ipv6.Priority
    leafs["outgoing-interface"] = ipv6.OutgoingInterface
    return leafs
}

func (ipv6 *Ping_Input_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Ping_Input_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Ping_Input_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Ping_Input_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Ping_Input_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Ping_Input_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Ping_Input_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Ping_Input_Ipv6) GetParentYangName() string { return "input" }

// Ping_Output
type Ping_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    PingResponse Ping_Output_PingResponse
}

func (output *Ping_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Ping_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Ping_Output) GetGoName(yname string) string {
    if yname == "ping-response" { return "PingResponse" }
    return ""
}

func (output *Ping_Output) GetSegmentPath() string {
    return "output"
}

func (output *Ping_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ping-response" {
        return &output.PingResponse
    }
    return nil
}

func (output *Ping_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ping-response"] = &output.PingResponse
    return children
}

func (output *Ping_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (output *Ping_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *Ping_Output) GetYangName() string { return "output" }

func (output *Ping_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *Ping_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *Ping_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *Ping_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Ping_Output) GetParent() types.Entity { return output.parent }

func (output *Ping_Output) GetParentYangName() string { return "ping" }

// Ping_Output_PingResponse
type Ping_Output_PingResponse struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of Ping_Output_PingResponse_Ipv4.
    Ipv4 []Ping_Output_PingResponse_Ipv4

    
    Ipv6 Ping_Output_PingResponse_Ipv6
}

func (pingResponse *Ping_Output_PingResponse) GetFilter() yfilter.YFilter { return pingResponse.YFilter }

func (pingResponse *Ping_Output_PingResponse) SetFilter(yf yfilter.YFilter) { pingResponse.YFilter = yf }

func (pingResponse *Ping_Output_PingResponse) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (pingResponse *Ping_Output_PingResponse) GetSegmentPath() string {
    return "ping-response"
}

func (pingResponse *Ping_Output_PingResponse) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        for _, c := range pingResponse.Ipv4 {
            if pingResponse.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ping_Output_PingResponse_Ipv4{}
        pingResponse.Ipv4 = append(pingResponse.Ipv4, child)
        return &pingResponse.Ipv4[len(pingResponse.Ipv4)-1]
    }
    if childYangName == "ipv6" {
        return &pingResponse.Ipv6
    }
    return nil
}

func (pingResponse *Ping_Output_PingResponse) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pingResponse.Ipv4 {
        children[pingResponse.Ipv4[i].GetSegmentPath()] = &pingResponse.Ipv4[i]
    }
    children["ipv6"] = &pingResponse.Ipv6
    return children
}

func (pingResponse *Ping_Output_PingResponse) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pingResponse *Ping_Output_PingResponse) GetBundleName() string { return "cisco_ios_xr" }

func (pingResponse *Ping_Output_PingResponse) GetYangName() string { return "ping-response" }

func (pingResponse *Ping_Output_PingResponse) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pingResponse *Ping_Output_PingResponse) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pingResponse *Ping_Output_PingResponse) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pingResponse *Ping_Output_PingResponse) SetParent(parent types.Entity) { pingResponse.parent = parent }

func (pingResponse *Ping_Output_PingResponse) GetParent() types.Entity { return pingResponse.parent }

func (pingResponse *Ping_Output_PingResponse) GetParentYangName() string { return "output" }

// Ping_Output_PingResponse_Ipv4
type Ping_Output_PingResponse_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Ping_Output_PingResponse_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "repeat-count" { return "RepeatCount" }
    if yname == "data-size" { return "DataSize" }
    if yname == "timeout" { return "Timeout" }
    if yname == "interval" { return "Interval" }
    if yname == "pattern" { return "Pattern" }
    if yname == "sweep" { return "Sweep" }
    if yname == "hits" { return "Hits" }
    if yname == "total" { return "Total" }
    if yname == "success-rate" { return "SuccessRate" }
    if yname == "rtt-min" { return "RttMin" }
    if yname == "rtt-avg" { return "RttAvg" }
    if yname == "rtt-max" { return "RttMax" }
    if yname == "sweep-min" { return "SweepMin" }
    if yname == "sweep-max" { return "SweepMax" }
    if yname == "rotate-pattern" { return "RotatePattern" }
    if yname == "ping-error-response" { return "PingErrorResponse" }
    if yname == "replies" { return "Replies" }
    return ""
}

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetSegmentPath() string {
    return "ipv4" + "[destination='" + fmt.Sprintf("%v", ipv4.Destination) + "']"
}

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "replies" {
        return &ipv4.Replies
    }
    return nil
}

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["replies"] = &ipv4.Replies
    return children
}

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv4.Destination
    leafs["repeat-count"] = ipv4.RepeatCount
    leafs["data-size"] = ipv4.DataSize
    leafs["timeout"] = ipv4.Timeout
    leafs["interval"] = ipv4.Interval
    leafs["pattern"] = ipv4.Pattern
    leafs["sweep"] = ipv4.Sweep
    leafs["hits"] = ipv4.Hits
    leafs["total"] = ipv4.Total
    leafs["success-rate"] = ipv4.SuccessRate
    leafs["rtt-min"] = ipv4.RttMin
    leafs["rtt-avg"] = ipv4.RttAvg
    leafs["rtt-max"] = ipv4.RttMax
    leafs["sweep-min"] = ipv4.SweepMin
    leafs["sweep-max"] = ipv4.SweepMax
    leafs["rotate-pattern"] = ipv4.RotatePattern
    leafs["ping-error-response"] = ipv4.PingErrorResponse
    return leafs
}

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Ping_Output_PingResponse_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Ping_Output_PingResponse_Ipv4) GetParentYangName() string { return "ping-response" }

// Ping_Output_PingResponse_Ipv4_Replies
type Ping_Output_PingResponse_Ipv4_Replies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of Ping_Output_PingResponse_Ipv4_Replies_Reply.
    Reply []Ping_Output_PingResponse_Ipv4_Replies_Reply
}

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetFilter() yfilter.YFilter { return replies.YFilter }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) SetFilter(yf yfilter.YFilter) { replies.YFilter = yf }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetGoName(yname string) string {
    if yname == "reply" { return "Reply" }
    return ""
}

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetSegmentPath() string {
    return "replies"
}

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reply" {
        for _, c := range replies.Reply {
            if replies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ping_Output_PingResponse_Ipv4_Replies_Reply{}
        replies.Reply = append(replies.Reply, child)
        return &replies.Reply[len(replies.Reply)-1]
    }
    return nil
}

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range replies.Reply {
        children[replies.Reply[i].GetSegmentPath()] = &replies.Reply[i]
    }
    return children
}

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetBundleName() string { return "cisco_ios_xr" }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetYangName() string { return "replies" }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) SetParent(parent types.Entity) { replies.parent = parent }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetParent() types.Entity { return replies.parent }

func (replies *Ping_Output_PingResponse_Ipv4_Replies) GetParentYangName() string { return "ipv4" }

// Ping_Output_PingResponse_Ipv4_Replies_Reply
type Ping_Output_PingResponse_Ipv4_Replies_Reply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the reply list. The type is interface{}
    // with range: 1..2147483647.
    ReplyIndex interface{}

    // Response for each packet. The type is string.
    Result interface{}

    
    BroadcastReplyAddresses Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses
}

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetFilter() yfilter.YFilter { return reply.YFilter }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) SetFilter(yf yfilter.YFilter) { reply.YFilter = yf }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetGoName(yname string) string {
    if yname == "reply-index" { return "ReplyIndex" }
    if yname == "result" { return "Result" }
    if yname == "broadcast-reply-addresses" { return "BroadcastReplyAddresses" }
    return ""
}

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetSegmentPath() string {
    return "reply" + "[reply-index='" + fmt.Sprintf("%v", reply.ReplyIndex) + "']"
}

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "broadcast-reply-addresses" {
        return &reply.BroadcastReplyAddresses
    }
    return nil
}

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["broadcast-reply-addresses"] = &reply.BroadcastReplyAddresses
    return children
}

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reply-index"] = reply.ReplyIndex
    leafs["result"] = reply.Result
    return leafs
}

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetBundleName() string { return "cisco_ios_xr" }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetYangName() string { return "reply" }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) SetParent(parent types.Entity) { reply.parent = parent }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetParent() types.Entity { return reply.parent }

func (reply *Ping_Output_PingResponse_Ipv4_Replies_Reply) GetParentYangName() string { return "replies" }

// Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses
type Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of
    // Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress.
    BroadcastReplyAddress []Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress
}

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetFilter() yfilter.YFilter { return broadcastReplyAddresses.YFilter }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) SetFilter(yf yfilter.YFilter) { broadcastReplyAddresses.YFilter = yf }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetGoName(yname string) string {
    if yname == "broadcast-reply-address" { return "BroadcastReplyAddress" }
    return ""
}

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetSegmentPath() string {
    return "broadcast-reply-addresses"
}

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "broadcast-reply-address" {
        for _, c := range broadcastReplyAddresses.BroadcastReplyAddress {
            if broadcastReplyAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress{}
        broadcastReplyAddresses.BroadcastReplyAddress = append(broadcastReplyAddresses.BroadcastReplyAddress, child)
        return &broadcastReplyAddresses.BroadcastReplyAddress[len(broadcastReplyAddresses.BroadcastReplyAddress)-1]
    }
    return nil
}

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range broadcastReplyAddresses.BroadcastReplyAddress {
        children[broadcastReplyAddresses.BroadcastReplyAddress[i].GetSegmentPath()] = &broadcastReplyAddresses.BroadcastReplyAddress[i]
    }
    return children
}

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetYangName() string { return "broadcast-reply-addresses" }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) SetParent(parent types.Entity) { broadcastReplyAddresses.parent = parent }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetParent() types.Entity { return broadcastReplyAddresses.parent }

func (broadcastReplyAddresses *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses) GetParentYangName() string { return "reply" }

// Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress
type Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Broadcast reply address. The type is string.
    ReplyAddress interface{}

    // Sign for each reply packet. The type is string.
    Result interface{}
}

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetFilter() yfilter.YFilter { return broadcastReplyAddress.YFilter }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) SetFilter(yf yfilter.YFilter) { broadcastReplyAddress.YFilter = yf }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetGoName(yname string) string {
    if yname == "reply-address" { return "ReplyAddress" }
    if yname == "result" { return "Result" }
    return ""
}

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetSegmentPath() string {
    return "broadcast-reply-address" + "[reply-address='" + fmt.Sprintf("%v", broadcastReplyAddress.ReplyAddress) + "']"
}

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reply-address"] = broadcastReplyAddress.ReplyAddress
    leafs["result"] = broadcastReplyAddress.Result
    return leafs
}

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetBundleName() string { return "cisco_ios_xr" }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetYangName() string { return "broadcast-reply-address" }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) SetParent(parent types.Entity) { broadcastReplyAddress.parent = parent }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetParent() types.Entity { return broadcastReplyAddress.parent }

func (broadcastReplyAddress *Ping_Output_PingResponse_Ipv4_Replies_Reply_BroadcastReplyAddresses_BroadcastReplyAddress) GetParentYangName() string { return "broadcast-reply-addresses" }

// Ping_Output_PingResponse_Ipv6
type Ping_Output_PingResponse_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Ping_Output_PingResponse_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    if yname == "repeat-count" { return "RepeatCount" }
    if yname == "data-size" { return "DataSize" }
    if yname == "timeout" { return "Timeout" }
    if yname == "interval" { return "Interval" }
    if yname == "pattern" { return "Pattern" }
    if yname == "sweep" { return "Sweep" }
    if yname == "sweep-min" { return "SweepMin" }
    if yname == "sweep-max" { return "SweepMax" }
    if yname == "rotate-pattern" { return "RotatePattern" }
    if yname == "hits" { return "Hits" }
    if yname == "total" { return "Total" }
    if yname == "success-rate" { return "SuccessRate" }
    if yname == "rtt-min" { return "RttMin" }
    if yname == "rtt-avg" { return "RttAvg" }
    if yname == "rtt-max" { return "RttMax" }
    if yname == "replies" { return "Replies" }
    return ""
}

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "replies" {
        return &ipv6.Replies
    }
    return nil
}

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["replies"] = &ipv6.Replies
    return children
}

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = ipv6.Destination
    leafs["repeat-count"] = ipv6.RepeatCount
    leafs["data-size"] = ipv6.DataSize
    leafs["timeout"] = ipv6.Timeout
    leafs["interval"] = ipv6.Interval
    leafs["pattern"] = ipv6.Pattern
    leafs["sweep"] = ipv6.Sweep
    leafs["sweep-min"] = ipv6.SweepMin
    leafs["sweep-max"] = ipv6.SweepMax
    leafs["rotate-pattern"] = ipv6.RotatePattern
    leafs["hits"] = ipv6.Hits
    leafs["total"] = ipv6.Total
    leafs["success-rate"] = ipv6.SuccessRate
    leafs["rtt-min"] = ipv6.RttMin
    leafs["rtt-avg"] = ipv6.RttAvg
    leafs["rtt-max"] = ipv6.RttMax
    return leafs
}

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Ping_Output_PingResponse_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Ping_Output_PingResponse_Ipv6) GetParentYangName() string { return "ping-response" }

// Ping_Output_PingResponse_Ipv6_Replies
type Ping_Output_PingResponse_Ipv6_Replies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of Ping_Output_PingResponse_Ipv6_Replies_Reply.
    Reply []Ping_Output_PingResponse_Ipv6_Replies_Reply
}

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetFilter() yfilter.YFilter { return replies.YFilter }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) SetFilter(yf yfilter.YFilter) { replies.YFilter = yf }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetGoName(yname string) string {
    if yname == "reply" { return "Reply" }
    return ""
}

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetSegmentPath() string {
    return "replies"
}

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reply" {
        for _, c := range replies.Reply {
            if replies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ping_Output_PingResponse_Ipv6_Replies_Reply{}
        replies.Reply = append(replies.Reply, child)
        return &replies.Reply[len(replies.Reply)-1]
    }
    return nil
}

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range replies.Reply {
        children[replies.Reply[i].GetSegmentPath()] = &replies.Reply[i]
    }
    return children
}

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetBundleName() string { return "cisco_ios_xr" }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetYangName() string { return "replies" }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) SetParent(parent types.Entity) { replies.parent = parent }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetParent() types.Entity { return replies.parent }

func (replies *Ping_Output_PingResponse_Ipv6_Replies) GetParentYangName() string { return "ipv6" }

// Ping_Output_PingResponse_Ipv6_Replies_Reply
type Ping_Output_PingResponse_Ipv6_Replies_Reply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index of the reply list. The type is interface{}
    // with range: 1..2147483647.
    ReplyIndex interface{}

    // Response for each packet. The type is string.
    Result interface{}
}

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetFilter() yfilter.YFilter { return reply.YFilter }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) SetFilter(yf yfilter.YFilter) { reply.YFilter = yf }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetGoName(yname string) string {
    if yname == "reply-index" { return "ReplyIndex" }
    if yname == "result" { return "Result" }
    return ""
}

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetSegmentPath() string {
    return "reply" + "[reply-index='" + fmt.Sprintf("%v", reply.ReplyIndex) + "']"
}

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reply-index"] = reply.ReplyIndex
    leafs["result"] = reply.Result
    return leafs
}

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetBundleName() string { return "cisco_ios_xr" }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetYangName() string { return "reply" }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) SetParent(parent types.Entity) { reply.parent = parent }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetParent() types.Entity { return reply.parent }

func (reply *Ping_Output_PingResponse_Ipv6_Replies_Reply) GetParentYangName() string { return "replies" }

