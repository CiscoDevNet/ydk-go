// This module contains a collection of YANG definitions
// for Flexible NetFlow Operational data.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package flow_monitor_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package flow_monitor_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-flow-monitor-oper flow-monitors}", reflect.TypeOf(FlowMonitors{}))
    ydk.RegisterEntity("Cisco-IOS-XE-flow-monitor-oper:flow-monitors", reflect.TypeOf(FlowMonitors{}))
}

// FlowMonitors
// All of the flow monitors
type FlowMonitors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of Flow monitors. The type is slice of FlowMonitors_FlowMonitor.
    FlowMonitor []FlowMonitors_FlowMonitor
}

func (flowMonitors *FlowMonitors) GetFilter() yfilter.YFilter { return flowMonitors.YFilter }

func (flowMonitors *FlowMonitors) SetFilter(yf yfilter.YFilter) { flowMonitors.YFilter = yf }

func (flowMonitors *FlowMonitors) GetGoName(yname string) string {
    if yname == "flow-monitor" { return "FlowMonitor" }
    return ""
}

func (flowMonitors *FlowMonitors) GetSegmentPath() string {
    return "Cisco-IOS-XE-flow-monitor-oper:flow-monitors"
}

func (flowMonitors *FlowMonitors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-monitor" {
        for _, c := range flowMonitors.FlowMonitor {
            if flowMonitors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FlowMonitors_FlowMonitor{}
        flowMonitors.FlowMonitor = append(flowMonitors.FlowMonitor, child)
        return &flowMonitors.FlowMonitor[len(flowMonitors.FlowMonitor)-1]
    }
    return nil
}

func (flowMonitors *FlowMonitors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flowMonitors.FlowMonitor {
        children[flowMonitors.FlowMonitor[i].GetSegmentPath()] = &flowMonitors.FlowMonitor[i]
    }
    return children
}

func (flowMonitors *FlowMonitors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flowMonitors *FlowMonitors) GetBundleName() string { return "cisco_ios_xe" }

func (flowMonitors *FlowMonitors) GetYangName() string { return "flow-monitors" }

func (flowMonitors *FlowMonitors) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flowMonitors *FlowMonitors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flowMonitors *FlowMonitors) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flowMonitors *FlowMonitors) SetParent(parent types.Entity) { flowMonitors.parent = parent }

func (flowMonitors *FlowMonitors) GetParent() types.Entity { return flowMonitors.parent }

func (flowMonitors *FlowMonitors) GetParentYangName() string { return "Cisco-IOS-XE-flow-monitor-oper" }

// FlowMonitors_FlowMonitor
// List of Flow monitors
type FlowMonitors_FlowMonitor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the flow monitor. The type is string.
    Name interface{}

    // Time the flow monitor data was collected in seconds. The type is
    // interface{} with range: 0..18446744073709551615.
    TimeCollected interface{}

    // All the flows for this flow monitor.
    Flows FlowMonitors_FlowMonitor_Flows
}

func (flowMonitor *FlowMonitors_FlowMonitor) GetFilter() yfilter.YFilter { return flowMonitor.YFilter }

func (flowMonitor *FlowMonitors_FlowMonitor) SetFilter(yf yfilter.YFilter) { flowMonitor.YFilter = yf }

func (flowMonitor *FlowMonitors_FlowMonitor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "time-collected" { return "TimeCollected" }
    if yname == "flows" { return "Flows" }
    return ""
}

func (flowMonitor *FlowMonitors_FlowMonitor) GetSegmentPath() string {
    return "flow-monitor" + "[name='" + fmt.Sprintf("%v", flowMonitor.Name) + "']"
}

func (flowMonitor *FlowMonitors_FlowMonitor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flows" {
        return &flowMonitor.Flows
    }
    return nil
}

func (flowMonitor *FlowMonitors_FlowMonitor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flows"] = &flowMonitor.Flows
    return children
}

func (flowMonitor *FlowMonitors_FlowMonitor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = flowMonitor.Name
    leafs["time-collected"] = flowMonitor.TimeCollected
    return leafs
}

func (flowMonitor *FlowMonitors_FlowMonitor) GetBundleName() string { return "cisco_ios_xe" }

func (flowMonitor *FlowMonitors_FlowMonitor) GetYangName() string { return "flow-monitor" }

func (flowMonitor *FlowMonitors_FlowMonitor) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flowMonitor *FlowMonitors_FlowMonitor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flowMonitor *FlowMonitors_FlowMonitor) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flowMonitor *FlowMonitors_FlowMonitor) SetParent(parent types.Entity) { flowMonitor.parent = parent }

func (flowMonitor *FlowMonitors_FlowMonitor) GetParent() types.Entity { return flowMonitor.parent }

func (flowMonitor *FlowMonitors_FlowMonitor) GetParentYangName() string { return "flow-monitors" }

// FlowMonitors_FlowMonitor_Flows
// All the flows for this flow monitor
type FlowMonitors_FlowMonitor_Flows struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of flows. The type is slice of FlowMonitors_FlowMonitor_Flows_Flow.
    Flow []FlowMonitors_FlowMonitor_Flows_Flow
}

func (flows *FlowMonitors_FlowMonitor_Flows) GetFilter() yfilter.YFilter { return flows.YFilter }

func (flows *FlowMonitors_FlowMonitor_Flows) SetFilter(yf yfilter.YFilter) { flows.YFilter = yf }

func (flows *FlowMonitors_FlowMonitor_Flows) GetGoName(yname string) string {
    if yname == "flow" { return "Flow" }
    return ""
}

func (flows *FlowMonitors_FlowMonitor_Flows) GetSegmentPath() string {
    return "flows"
}

func (flows *FlowMonitors_FlowMonitor_Flows) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow" {
        for _, c := range flows.Flow {
            if flows.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FlowMonitors_FlowMonitor_Flows_Flow{}
        flows.Flow = append(flows.Flow, child)
        return &flows.Flow[len(flows.Flow)-1]
    }
    return nil
}

func (flows *FlowMonitors_FlowMonitor_Flows) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flows.Flow {
        children[flows.Flow[i].GetSegmentPath()] = &flows.Flow[i]
    }
    return children
}

func (flows *FlowMonitors_FlowMonitor_Flows) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flows *FlowMonitors_FlowMonitor_Flows) GetBundleName() string { return "cisco_ios_xe" }

func (flows *FlowMonitors_FlowMonitor_Flows) GetYangName() string { return "flows" }

func (flows *FlowMonitors_FlowMonitor_Flows) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flows *FlowMonitors_FlowMonitor_Flows) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flows *FlowMonitors_FlowMonitor_Flows) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flows *FlowMonitors_FlowMonitor_Flows) SetParent(parent types.Entity) { flows.parent = parent }

func (flows *FlowMonitors_FlowMonitor_Flows) GetParent() types.Entity { return flows.parent }

func (flows *FlowMonitors_FlowMonitor_Flows) GetParentYangName() string { return "flow-monitor" }

// FlowMonitors_FlowMonitor_Flows_Flow
// List of flows
type FlowMonitors_FlowMonitor_Flows_Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source address of the flow. The type is string.
    SourceAddress interface{}

    // This attribute is a key. Destination address of the flow. The type is
    // string.
    DestinationAddress interface{}

    // This attribute is a key. Input interface of the flow. The type is string.
    InterfaceInput interface{}

    // This attribute is a key. Multicast flow. The type is string.
    IsMulticast interface{}

    // This attribute is a key. VRF ID input. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    VrfIdInput interface{}

    // This attribute is a key. Source port number. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    SourcePort interface{}

    // This attribute is a key. Destination port number. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    DestinationPort interface{}

    // This attribute is a key. ip-tos value. The type is string.
    IpTos interface{}

    // This attribute is a key. IP protocol number. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    IpProtocol interface{}

    // Output interface of the flow. The type is string.
    InterfaceOutput interface{}

    // Number of bytes passed through. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Bytes interface{}

    // Number of packets passed through. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Packets interface{}
}

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetFilter() yfilter.YFilter { return flow.YFilter }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) SetFilter(yf yfilter.YFilter) { flow.YFilter = yf }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "interface-input" { return "InterfaceInput" }
    if yname == "is-multicast" { return "IsMulticast" }
    if yname == "vrf-id-input" { return "VrfIdInput" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "ip-tos" { return "IpTos" }
    if yname == "ip-protocol" { return "IpProtocol" }
    if yname == "interface-output" { return "InterfaceOutput" }
    if yname == "bytes" { return "Bytes" }
    if yname == "packets" { return "Packets" }
    return ""
}

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetSegmentPath() string {
    return "flow" + "[source-address='" + fmt.Sprintf("%v", flow.SourceAddress) + "']" + "[destination-address='" + fmt.Sprintf("%v", flow.DestinationAddress) + "']" + "[interface-input='" + fmt.Sprintf("%v", flow.InterfaceInput) + "']" + "[is-multicast='" + fmt.Sprintf("%v", flow.IsMulticast) + "']" + "[vrf-id-input='" + fmt.Sprintf("%v", flow.VrfIdInput) + "']" + "[source-port='" + fmt.Sprintf("%v", flow.SourcePort) + "']" + "[destination-port='" + fmt.Sprintf("%v", flow.DestinationPort) + "']" + "[ip-tos='" + fmt.Sprintf("%v", flow.IpTos) + "']" + "[ip-protocol='" + fmt.Sprintf("%v", flow.IpProtocol) + "']"
}

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = flow.SourceAddress
    leafs["destination-address"] = flow.DestinationAddress
    leafs["interface-input"] = flow.InterfaceInput
    leafs["is-multicast"] = flow.IsMulticast
    leafs["vrf-id-input"] = flow.VrfIdInput
    leafs["source-port"] = flow.SourcePort
    leafs["destination-port"] = flow.DestinationPort
    leafs["ip-tos"] = flow.IpTos
    leafs["ip-protocol"] = flow.IpProtocol
    leafs["interface-output"] = flow.InterfaceOutput
    leafs["bytes"] = flow.Bytes
    leafs["packets"] = flow.Packets
    return leafs
}

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetBundleName() string { return "cisco_ios_xe" }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetYangName() string { return "flow" }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) SetParent(parent types.Entity) { flow.parent = parent }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetParent() types.Entity { return flow.parent }

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetParentYangName() string { return "flows" }

