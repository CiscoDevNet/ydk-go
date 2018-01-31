// This module contains a collection of YANG definitions
// for Cisco IOS-XR manageability-perfmgmt package operational data.
// 
// This module contains definitions
// for the following management objects:
//   perf-mgmt: Performance Management agent operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package manageability_perfmgmt_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package manageability_perfmgmt_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-manageability-perfmgmt-oper perf-mgmt}", reflect.TypeOf(PerfMgmt{}))
    ydk.RegisterEntity("Cisco-IOS-XR-manageability-perfmgmt-oper:perf-mgmt", reflect.TypeOf(PerfMgmt{}))
}

// PerfMgmt
// Performance Management agent operational data
type PerfMgmt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data from periodic requests.
    Periodic PerfMgmt_Periodic

    // Data from monitor (one history period) requests.
    Monitor PerfMgmt_Monitor
}

func (perfMgmt *PerfMgmt) GetFilter() yfilter.YFilter { return perfMgmt.YFilter }

func (perfMgmt *PerfMgmt) SetFilter(yf yfilter.YFilter) { perfMgmt.YFilter = yf }

func (perfMgmt *PerfMgmt) GetGoName(yname string) string {
    if yname == "periodic" { return "Periodic" }
    if yname == "monitor" { return "Monitor" }
    return ""
}

func (perfMgmt *PerfMgmt) GetSegmentPath() string {
    return "Cisco-IOS-XR-manageability-perfmgmt-oper:perf-mgmt"
}

func (perfMgmt *PerfMgmt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "periodic" {
        return &perfMgmt.Periodic
    }
    if childYangName == "monitor" {
        return &perfMgmt.Monitor
    }
    return nil
}

func (perfMgmt *PerfMgmt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["periodic"] = &perfMgmt.Periodic
    children["monitor"] = &perfMgmt.Monitor
    return children
}

func (perfMgmt *PerfMgmt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (perfMgmt *PerfMgmt) GetBundleName() string { return "cisco_ios_xr" }

func (perfMgmt *PerfMgmt) GetYangName() string { return "perf-mgmt" }

func (perfMgmt *PerfMgmt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (perfMgmt *PerfMgmt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (perfMgmt *PerfMgmt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (perfMgmt *PerfMgmt) SetParent(parent types.Entity) { perfMgmt.parent = parent }

func (perfMgmt *PerfMgmt) GetParent() types.Entity { return perfMgmt.parent }

func (perfMgmt *PerfMgmt) GetParentYangName() string { return "Cisco-IOS-XR-manageability-perfmgmt-oper" }

// PerfMgmt_Periodic
// Data from periodic requests
type PerfMgmt_Periodic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collected OSPF data.
    Ospf PerfMgmt_Periodic_Ospf

    // Collected MPLS data.
    Mpls PerfMgmt_Periodic_Mpls

    // Nodes for which data is collected.
    Nodes PerfMgmt_Periodic_Nodes

    // Collected BGP data.
    Bgp PerfMgmt_Periodic_Bgp

    // Collected Interface data.
    Interface PerfMgmt_Periodic_Interface
}

func (periodic *PerfMgmt_Periodic) GetFilter() yfilter.YFilter { return periodic.YFilter }

func (periodic *PerfMgmt_Periodic) SetFilter(yf yfilter.YFilter) { periodic.YFilter = yf }

func (periodic *PerfMgmt_Periodic) GetGoName(yname string) string {
    if yname == "ospf" { return "Ospf" }
    if yname == "mpls" { return "Mpls" }
    if yname == "nodes" { return "Nodes" }
    if yname == "bgp" { return "Bgp" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (periodic *PerfMgmt_Periodic) GetSegmentPath() string {
    return "periodic"
}

func (periodic *PerfMgmt_Periodic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf" {
        return &periodic.Ospf
    }
    if childYangName == "mpls" {
        return &periodic.Mpls
    }
    if childYangName == "nodes" {
        return &periodic.Nodes
    }
    if childYangName == "bgp" {
        return &periodic.Bgp
    }
    if childYangName == "interface" {
        return &periodic.Interface
    }
    return nil
}

func (periodic *PerfMgmt_Periodic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospf"] = &periodic.Ospf
    children["mpls"] = &periodic.Mpls
    children["nodes"] = &periodic.Nodes
    children["bgp"] = &periodic.Bgp
    children["interface"] = &periodic.Interface
    return children
}

func (periodic *PerfMgmt_Periodic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (periodic *PerfMgmt_Periodic) GetBundleName() string { return "cisco_ios_xr" }

func (periodic *PerfMgmt_Periodic) GetYangName() string { return "periodic" }

func (periodic *PerfMgmt_Periodic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (periodic *PerfMgmt_Periodic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (periodic *PerfMgmt_Periodic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (periodic *PerfMgmt_Periodic) SetParent(parent types.Entity) { periodic.parent = parent }

func (periodic *PerfMgmt_Periodic) GetParent() types.Entity { return periodic.parent }

func (periodic *PerfMgmt_Periodic) GetParentYangName() string { return "perf-mgmt" }

// PerfMgmt_Periodic_Ospf
// Collected OSPF data
type PerfMgmt_Periodic_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF v2 instances for which protocol statistics are collected.
    Ospfv2ProtocolInstances PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances

    // OSPF v3 instances for which protocol statistics are collected.
    Ospfv3ProtocolInstances PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances
}

func (ospf *PerfMgmt_Periodic_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PerfMgmt_Periodic_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PerfMgmt_Periodic_Ospf) GetGoName(yname string) string {
    if yname == "ospfv2-protocol-instances" { return "Ospfv2ProtocolInstances" }
    if yname == "ospfv3-protocol-instances" { return "Ospfv3ProtocolInstances" }
    return ""
}

func (ospf *PerfMgmt_Periodic_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PerfMgmt_Periodic_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2-protocol-instances" {
        return &ospf.Ospfv2ProtocolInstances
    }
    if childYangName == "ospfv3-protocol-instances" {
        return &ospf.Ospfv3ProtocolInstances
    }
    return nil
}

func (ospf *PerfMgmt_Periodic_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv2-protocol-instances"] = &ospf.Ospfv2ProtocolInstances
    children["ospfv3-protocol-instances"] = &ospf.Ospfv3ProtocolInstances
    return children
}

func (ospf *PerfMgmt_Periodic_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospf *PerfMgmt_Periodic_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PerfMgmt_Periodic_Ospf) GetYangName() string { return "ospf" }

func (ospf *PerfMgmt_Periodic_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PerfMgmt_Periodic_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PerfMgmt_Periodic_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PerfMgmt_Periodic_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PerfMgmt_Periodic_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PerfMgmt_Periodic_Ospf) GetParentYangName() string { return "periodic" }

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances
// OSPF v2 instances for which protocol statistics
// are collected
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v2 instance. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance.
    Ospfv2ProtocolInstance []PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
}

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetFilter() yfilter.YFilter { return ospfv2ProtocolInstances.YFilter }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) SetFilter(yf yfilter.YFilter) { ospfv2ProtocolInstances.YFilter = yf }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetGoName(yname string) string {
    if yname == "ospfv2-protocol-instance" { return "Ospfv2ProtocolInstance" }
    return ""
}

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetSegmentPath() string {
    return "ospfv2-protocol-instances"
}

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2-protocol-instance" {
        for _, c := range ospfv2ProtocolInstances.Ospfv2ProtocolInstance {
            if ospfv2ProtocolInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance{}
        ospfv2ProtocolInstances.Ospfv2ProtocolInstance = append(ospfv2ProtocolInstances.Ospfv2ProtocolInstance, child)
        return &ospfv2ProtocolInstances.Ospfv2ProtocolInstance[len(ospfv2ProtocolInstances.Ospfv2ProtocolInstance)-1]
    }
    return nil
}

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfv2ProtocolInstances.Ospfv2ProtocolInstance {
        children[ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i].GetSegmentPath()] = &ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i]
    }
    return children
}

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetYangName() string { return "ospfv2-protocol-instances" }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) SetParent(parent types.Entity) { ospfv2ProtocolInstances.parent = parent }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetParent() types.Entity { return ospfv2ProtocolInstances.parent }

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetParentYangName() string { return "ospf" }

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
// Protocol samples for a particular OSPF v2
// instance
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v2 instance.
    Samples PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
}

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetFilter() yfilter.YFilter { return ospfv2ProtocolInstance.YFilter }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) SetFilter(yf yfilter.YFilter) { ospfv2ProtocolInstance.YFilter = yf }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetSegmentPath() string {
    return "ospfv2-protocol-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfv2ProtocolInstance.InstanceName) + "']"
}

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &ospfv2ProtocolInstance.Samples
    }
    return nil
}

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &ospfv2ProtocolInstance.Samples
    return children
}

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfv2ProtocolInstance.InstanceName
    return leafs
}

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetYangName() string { return "ospfv2-protocol-instance" }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) SetParent(parent types.Entity) { ospfv2ProtocolInstance.parent = parent }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetParent() types.Entity { return ospfv2ProtocolInstance.parent }

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetParentYangName() string { return "ospfv2-protocol-instances" }

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
// Sample Table for an OSPV v2 instance
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample.
    Sample []PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetParentYangName() string { return "ospfv2-protocol-instance" }

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Total number of packets received. The type is interface{} with range:
    // 0..4294967295.
    InputPackets interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputPackets interface{}

    // Number of Hello packets received. The type is interface{} with range:
    // 0..4294967295.
    InputHelloPackets interface{}

    // Number of Hello packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputHelloPackets interface{}

    // Number of DBD packets received. The type is interface{} with range:
    // 0..4294967295.
    InputDbDs interface{}

    // Number of LSA received in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    InputDbDsLsa interface{}

    // Number of DBD packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDs interface{}

    // Number of LSA sent in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDsLsa interface{}

    // Number of LS Requests received. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequests interface{}

    // Number of LSA received in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequestsLsa interface{}

    // Number of LS Requests sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequests interface{}

    // Number of LSA sent in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequestsLsa interface{}

    // Number of LSA Updates received. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdates interface{}

    // Number of LSA received in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdatesLsa interface{}

    // Number of LSA Updates sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdates interface{}

    // Number of LSA sent in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdatesLsa interface{}

    // Number of LSA Acknowledgements received. The type is interface{} with
    // range: 0..4294967295.
    InputLsaAcks interface{}

    // Number of LSA received in LSA Acknowledgements. The type is interface{}
    // with range: 0..4294967295.
    InputLsaAcksLsa interface{}

    // Number of LSA Acknowledgements sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaAcks interface{}

    // Number of LSA sent in LSA Acknowledgements. The type is interface{} with
    // range: 0..4294967295.
    OutputLsaAcksLsa interface{}

    // Number of packets received with checksum errors. The type is interface{}
    // with range: 0..4294967295.
    ChecksumErrors interface{}
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-packets" { return "InputPackets" }
    if yname == "output-packets" { return "OutputPackets" }
    if yname == "input-hello-packets" { return "InputHelloPackets" }
    if yname == "output-hello-packets" { return "OutputHelloPackets" }
    if yname == "input-db-ds" { return "InputDbDs" }
    if yname == "input-db-ds-lsa" { return "InputDbDsLsa" }
    if yname == "output-db-ds" { return "OutputDbDs" }
    if yname == "output-db-ds-lsa" { return "OutputDbDsLsa" }
    if yname == "input-ls-requests" { return "InputLsRequests" }
    if yname == "input-ls-requests-lsa" { return "InputLsRequestsLsa" }
    if yname == "output-ls-requests" { return "OutputLsRequests" }
    if yname == "output-ls-requests-lsa" { return "OutputLsRequestsLsa" }
    if yname == "input-lsa-updates" { return "InputLsaUpdates" }
    if yname == "input-lsa-updates-lsa" { return "InputLsaUpdatesLsa" }
    if yname == "output-lsa-updates" { return "OutputLsaUpdates" }
    if yname == "output-lsa-updates-lsa" { return "OutputLsaUpdatesLsa" }
    if yname == "input-lsa-acks" { return "InputLsaAcks" }
    if yname == "input-lsa-acks-lsa" { return "InputLsaAcksLsa" }
    if yname == "output-lsa-acks" { return "OutputLsaAcks" }
    if yname == "output-lsa-acks-lsa" { return "OutputLsaAcksLsa" }
    if yname == "checksum-errors" { return "ChecksumErrors" }
    return ""
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-packets"] = sample.InputPackets
    leafs["output-packets"] = sample.OutputPackets
    leafs["input-hello-packets"] = sample.InputHelloPackets
    leafs["output-hello-packets"] = sample.OutputHelloPackets
    leafs["input-db-ds"] = sample.InputDbDs
    leafs["input-db-ds-lsa"] = sample.InputDbDsLsa
    leafs["output-db-ds"] = sample.OutputDbDs
    leafs["output-db-ds-lsa"] = sample.OutputDbDsLsa
    leafs["input-ls-requests"] = sample.InputLsRequests
    leafs["input-ls-requests-lsa"] = sample.InputLsRequestsLsa
    leafs["output-ls-requests"] = sample.OutputLsRequests
    leafs["output-ls-requests-lsa"] = sample.OutputLsRequestsLsa
    leafs["input-lsa-updates"] = sample.InputLsaUpdates
    leafs["input-lsa-updates-lsa"] = sample.InputLsaUpdatesLsa
    leafs["output-lsa-updates"] = sample.OutputLsaUpdates
    leafs["output-lsa-updates-lsa"] = sample.OutputLsaUpdatesLsa
    leafs["input-lsa-acks"] = sample.InputLsaAcks
    leafs["input-lsa-acks-lsa"] = sample.InputLsaAcksLsa
    leafs["output-lsa-acks"] = sample.OutputLsaAcks
    leafs["output-lsa-acks-lsa"] = sample.OutputLsaAcksLsa
    leafs["checksum-errors"] = sample.ChecksumErrors
    return leafs
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances
// OSPF v3 instances for which protocol statistics
// are collected
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v3 instance. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance.
    Ospfv3ProtocolInstance []PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
}

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetFilter() yfilter.YFilter { return ospfv3ProtocolInstances.YFilter }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) SetFilter(yf yfilter.YFilter) { ospfv3ProtocolInstances.YFilter = yf }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetGoName(yname string) string {
    if yname == "ospfv3-protocol-instance" { return "Ospfv3ProtocolInstance" }
    return ""
}

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetSegmentPath() string {
    return "ospfv3-protocol-instances"
}

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv3-protocol-instance" {
        for _, c := range ospfv3ProtocolInstances.Ospfv3ProtocolInstance {
            if ospfv3ProtocolInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance{}
        ospfv3ProtocolInstances.Ospfv3ProtocolInstance = append(ospfv3ProtocolInstances.Ospfv3ProtocolInstance, child)
        return &ospfv3ProtocolInstances.Ospfv3ProtocolInstance[len(ospfv3ProtocolInstances.Ospfv3ProtocolInstance)-1]
    }
    return nil
}

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfv3ProtocolInstances.Ospfv3ProtocolInstance {
        children[ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i].GetSegmentPath()] = &ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i]
    }
    return children
}

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetYangName() string { return "ospfv3-protocol-instances" }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) SetParent(parent types.Entity) { ospfv3ProtocolInstances.parent = parent }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetParent() types.Entity { return ospfv3ProtocolInstances.parent }

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetParentYangName() string { return "ospf" }

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
// Protocol samples for a particular OSPF v3
// instance
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v3 instance.
    Samples PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
}

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetFilter() yfilter.YFilter { return ospfv3ProtocolInstance.YFilter }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) SetFilter(yf yfilter.YFilter) { ospfv3ProtocolInstance.YFilter = yf }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetSegmentPath() string {
    return "ospfv3-protocol-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfv3ProtocolInstance.InstanceName) + "']"
}

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &ospfv3ProtocolInstance.Samples
    }
    return nil
}

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &ospfv3ProtocolInstance.Samples
    return children
}

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfv3ProtocolInstance.InstanceName
    return leafs
}

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetYangName() string { return "ospfv3-protocol-instance" }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) SetParent(parent types.Entity) { ospfv3ProtocolInstance.parent = parent }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetParent() types.Entity { return ospfv3ProtocolInstance.parent }

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetParentYangName() string { return "ospfv3-protocol-instances" }

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
// Sample Table for an OSPV v3 instance
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample.
    Sample []PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetParentYangName() string { return "ospfv3-protocol-instance" }

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Total number of packets received. The type is interface{} with range:
    // 0..4294967295.
    InputPackets interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputPackets interface{}

    // Number of Hello packets received. The type is interface{} with range:
    // 0..4294967295.
    InputHelloPackets interface{}

    // Number of Hello packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputHelloPackets interface{}

    // Number of DBD packets received. The type is interface{} with range:
    // 0..4294967295.
    InputDbDs interface{}

    // Number of LSA received in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    InputDbDsLsa interface{}

    // Number of DBD packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDs interface{}

    // Number of LSA sent in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDsLsa interface{}

    // Number of LS Requests received. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequests interface{}

    // Number of LSA received in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequestsLsa interface{}

    // Number of LS Requests sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequests interface{}

    // Number of LSA sent in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequestsLsa interface{}

    // Number of LSA Updates received. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdates interface{}

    // Number of LSA received in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdatesLsa interface{}

    // Number of LSA Updates sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdates interface{}

    // Number of LSA sent in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdatesLsa interface{}

    // Number of LSA Acknowledgements received. The type is interface{} with
    // range: 0..4294967295.
    InputLsaAcks interface{}

    // Number of LSA received in LSA Acknowledgements. The type is interface{}
    // with range: 0..4294967295.
    InputLsaAcksLsa interface{}

    // Number of LSA Acknowledgements sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaAcks interface{}

    // Number of LSA sent in LSA Acknowledgements. The type is interface{} with
    // range: 0..4294967295.
    OutputLsaAcksLsa interface{}
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-packets" { return "InputPackets" }
    if yname == "output-packets" { return "OutputPackets" }
    if yname == "input-hello-packets" { return "InputHelloPackets" }
    if yname == "output-hello-packets" { return "OutputHelloPackets" }
    if yname == "input-db-ds" { return "InputDbDs" }
    if yname == "input-db-ds-lsa" { return "InputDbDsLsa" }
    if yname == "output-db-ds" { return "OutputDbDs" }
    if yname == "output-db-ds-lsa" { return "OutputDbDsLsa" }
    if yname == "input-ls-requests" { return "InputLsRequests" }
    if yname == "input-ls-requests-lsa" { return "InputLsRequestsLsa" }
    if yname == "output-ls-requests" { return "OutputLsRequests" }
    if yname == "output-ls-requests-lsa" { return "OutputLsRequestsLsa" }
    if yname == "input-lsa-updates" { return "InputLsaUpdates" }
    if yname == "input-lsa-updates-lsa" { return "InputLsaUpdatesLsa" }
    if yname == "output-lsa-updates" { return "OutputLsaUpdates" }
    if yname == "output-lsa-updates-lsa" { return "OutputLsaUpdatesLsa" }
    if yname == "input-lsa-acks" { return "InputLsaAcks" }
    if yname == "input-lsa-acks-lsa" { return "InputLsaAcksLsa" }
    if yname == "output-lsa-acks" { return "OutputLsaAcks" }
    if yname == "output-lsa-acks-lsa" { return "OutputLsaAcksLsa" }
    return ""
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-packets"] = sample.InputPackets
    leafs["output-packets"] = sample.OutputPackets
    leafs["input-hello-packets"] = sample.InputHelloPackets
    leafs["output-hello-packets"] = sample.OutputHelloPackets
    leafs["input-db-ds"] = sample.InputDbDs
    leafs["input-db-ds-lsa"] = sample.InputDbDsLsa
    leafs["output-db-ds"] = sample.OutputDbDs
    leafs["output-db-ds-lsa"] = sample.OutputDbDsLsa
    leafs["input-ls-requests"] = sample.InputLsRequests
    leafs["input-ls-requests-lsa"] = sample.InputLsRequestsLsa
    leafs["output-ls-requests"] = sample.OutputLsRequests
    leafs["output-ls-requests-lsa"] = sample.OutputLsRequestsLsa
    leafs["input-lsa-updates"] = sample.InputLsaUpdates
    leafs["input-lsa-updates-lsa"] = sample.InputLsaUpdatesLsa
    leafs["output-lsa-updates"] = sample.OutputLsaUpdates
    leafs["output-lsa-updates-lsa"] = sample.OutputLsaUpdatesLsa
    leafs["input-lsa-acks"] = sample.InputLsaAcks
    leafs["input-lsa-acks-lsa"] = sample.InputLsaAcksLsa
    leafs["output-lsa-acks"] = sample.OutputLsaAcks
    leafs["output-lsa-acks-lsa"] = sample.OutputLsaAcksLsa
    return leafs
}

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Mpls
// Collected MPLS data
type PerfMgmt_Periodic_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP neighbors for which statistics are collected.
    LdpNeighbors PerfMgmt_Periodic_Mpls_LdpNeighbors
}

func (mpls *PerfMgmt_Periodic_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *PerfMgmt_Periodic_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *PerfMgmt_Periodic_Mpls) GetGoName(yname string) string {
    if yname == "ldp-neighbors" { return "LdpNeighbors" }
    return ""
}

func (mpls *PerfMgmt_Periodic_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *PerfMgmt_Periodic_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-neighbors" {
        return &mpls.LdpNeighbors
    }
    return nil
}

func (mpls *PerfMgmt_Periodic_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ldp-neighbors"] = &mpls.LdpNeighbors
    return children
}

func (mpls *PerfMgmt_Periodic_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpls *PerfMgmt_Periodic_Mpls) GetBundleName() string { return "cisco_ios_xr" }

func (mpls *PerfMgmt_Periodic_Mpls) GetYangName() string { return "mpls" }

func (mpls *PerfMgmt_Periodic_Mpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpls *PerfMgmt_Periodic_Mpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpls *PerfMgmt_Periodic_Mpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpls *PerfMgmt_Periodic_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *PerfMgmt_Periodic_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *PerfMgmt_Periodic_Mpls) GetParentYangName() string { return "periodic" }

// PerfMgmt_Periodic_Mpls_LdpNeighbors
// LDP neighbors for which statistics are
// collected
type PerfMgmt_Periodic_Mpls_LdpNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular LDP neighbor. The type is slice of
    // PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor.
    LdpNeighbor []PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor
}

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetFilter() yfilter.YFilter { return ldpNeighbors.YFilter }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) SetFilter(yf yfilter.YFilter) { ldpNeighbors.YFilter = yf }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetGoName(yname string) string {
    if yname == "ldp-neighbor" { return "LdpNeighbor" }
    return ""
}

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetSegmentPath() string {
    return "ldp-neighbors"
}

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-neighbor" {
        for _, c := range ldpNeighbors.LdpNeighbor {
            if ldpNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor{}
        ldpNeighbors.LdpNeighbor = append(ldpNeighbors.LdpNeighbor, child)
        return &ldpNeighbors.LdpNeighbor[len(ldpNeighbors.LdpNeighbor)-1]
    }
    return nil
}

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ldpNeighbors.LdpNeighbor {
        children[ldpNeighbors.LdpNeighbor[i].GetSegmentPath()] = &ldpNeighbors.LdpNeighbor[i]
    }
    return children
}

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetBundleName() string { return "cisco_ios_xr" }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetYangName() string { return "ldp-neighbors" }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) SetParent(parent types.Entity) { ldpNeighbors.parent = parent }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetParent() types.Entity { return ldpNeighbors.parent }

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetParentYangName() string { return "mpls" }

// PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor
// Samples for a particular LDP neighbor
type PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Nbr interface{}

    // Samples for a particular LDP neighbor.
    Samples PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples
}

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetFilter() yfilter.YFilter { return ldpNeighbor.YFilter }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) SetFilter(yf yfilter.YFilter) { ldpNeighbor.YFilter = yf }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetGoName(yname string) string {
    if yname == "nbr" { return "Nbr" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetSegmentPath() string {
    return "ldp-neighbor" + "[nbr='" + fmt.Sprintf("%v", ldpNeighbor.Nbr) + "']"
}

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &ldpNeighbor.Samples
    }
    return nil
}

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &ldpNeighbor.Samples
    return children
}

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nbr"] = ldpNeighbor.Nbr
    return leafs
}

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetYangName() string { return "ldp-neighbor" }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) SetParent(parent types.Entity) { ldpNeighbor.parent = parent }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetParent() types.Entity { return ldpNeighbor.parent }

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetParentYangName() string { return "ldp-neighbors" }

// PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples
// Samples for a particular LDP neighbor
type PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP neighbor statistics sample. The type is slice of
    // PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample.
    Sample []PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetParentYangName() string { return "ldp-neighbor" }

// PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
// LDP neighbor statistics sample
type PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Total messages sent. The type is interface{} with range: 0..65535.
    TotalMsgsSent interface{}

    // Total messages received. The type is interface{} with range: 0..65535.
    TotalMsgsRcvd interface{}

    // Init messages sent. The type is interface{} with range: 0..65535.
    InitMsgsSent interface{}

    // Tnit messages received. The type is interface{} with range: 0..65535.
    InitMsgsRcvd interface{}

    // Address messages sent. The type is interface{} with range: 0..65535.
    AddressMsgsSent interface{}

    // Address messages received. The type is interface{} with range: 0..65535.
    AddressMsgsRcvd interface{}

    // Address withdraw messages sent. The type is interface{} with range:
    // 0..65535.
    AddressWithdrawMsgsSent interface{}

    // Address withdraw messages received. The type is interface{} with range:
    // 0..65535.
    AddressWithdrawMsgsRcvd interface{}

    // Label mapping messages sent. The type is interface{} with range: 0..65535.
    LabelMappingMsgsSent interface{}

    // Label mapping messages received. The type is interface{} with range:
    // 0..65535.
    LabelMappingMsgsRcvd interface{}

    // Label withdraw messages sent. The type is interface{} with range: 0..65535.
    LabelWithdrawMsgsSent interface{}

    // Label withdraw messages received. The type is interface{} with range:
    // 0..65535.
    LabelWithdrawMsgsRcvd interface{}

    // Label release messages sent. The type is interface{} with range: 0..65535.
    LabelReleaseMsgsSent interface{}

    // Label release messages received. The type is interface{} with range:
    // 0..65535.
    LabelReleaseMsgsRcvd interface{}

    // Notification messages sent. The type is interface{} with range: 0..65535.
    NotificationMsgsSent interface{}

    // Notification messages received. The type is interface{} with range:
    // 0..65535.
    NotificationMsgsRcvd interface{}

    // Keepalive messages sent. The type is interface{} with range: 0..65535.
    KeepaliveMsgsSent interface{}

    // Keepalive messages received. The type is interface{} with range: 0..65535.
    KeepaliveMsgsRcvd interface{}
}

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "total-msgs-sent" { return "TotalMsgsSent" }
    if yname == "total-msgs-rcvd" { return "TotalMsgsRcvd" }
    if yname == "init-msgs-sent" { return "InitMsgsSent" }
    if yname == "init-msgs-rcvd" { return "InitMsgsRcvd" }
    if yname == "address-msgs-sent" { return "AddressMsgsSent" }
    if yname == "address-msgs-rcvd" { return "AddressMsgsRcvd" }
    if yname == "address-withdraw-msgs-sent" { return "AddressWithdrawMsgsSent" }
    if yname == "address-withdraw-msgs-rcvd" { return "AddressWithdrawMsgsRcvd" }
    if yname == "label-mapping-msgs-sent" { return "LabelMappingMsgsSent" }
    if yname == "label-mapping-msgs-rcvd" { return "LabelMappingMsgsRcvd" }
    if yname == "label-withdraw-msgs-sent" { return "LabelWithdrawMsgsSent" }
    if yname == "label-withdraw-msgs-rcvd" { return "LabelWithdrawMsgsRcvd" }
    if yname == "label-release-msgs-sent" { return "LabelReleaseMsgsSent" }
    if yname == "label-release-msgs-rcvd" { return "LabelReleaseMsgsRcvd" }
    if yname == "notification-msgs-sent" { return "NotificationMsgsSent" }
    if yname == "notification-msgs-rcvd" { return "NotificationMsgsRcvd" }
    if yname == "keepalive-msgs-sent" { return "KeepaliveMsgsSent" }
    if yname == "keepalive-msgs-rcvd" { return "KeepaliveMsgsRcvd" }
    return ""
}

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["total-msgs-sent"] = sample.TotalMsgsSent
    leafs["total-msgs-rcvd"] = sample.TotalMsgsRcvd
    leafs["init-msgs-sent"] = sample.InitMsgsSent
    leafs["init-msgs-rcvd"] = sample.InitMsgsRcvd
    leafs["address-msgs-sent"] = sample.AddressMsgsSent
    leafs["address-msgs-rcvd"] = sample.AddressMsgsRcvd
    leafs["address-withdraw-msgs-sent"] = sample.AddressWithdrawMsgsSent
    leafs["address-withdraw-msgs-rcvd"] = sample.AddressWithdrawMsgsRcvd
    leafs["label-mapping-msgs-sent"] = sample.LabelMappingMsgsSent
    leafs["label-mapping-msgs-rcvd"] = sample.LabelMappingMsgsRcvd
    leafs["label-withdraw-msgs-sent"] = sample.LabelWithdrawMsgsSent
    leafs["label-withdraw-msgs-rcvd"] = sample.LabelWithdrawMsgsRcvd
    leafs["label-release-msgs-sent"] = sample.LabelReleaseMsgsSent
    leafs["label-release-msgs-rcvd"] = sample.LabelReleaseMsgsRcvd
    leafs["notification-msgs-sent"] = sample.NotificationMsgsSent
    leafs["notification-msgs-rcvd"] = sample.NotificationMsgsRcvd
    leafs["keepalive-msgs-sent"] = sample.KeepaliveMsgsSent
    leafs["keepalive-msgs-rcvd"] = sample.KeepaliveMsgsRcvd
    return leafs
}

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Nodes
// Nodes for which data is collected
type PerfMgmt_Periodic_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Instance. The type is slice of PerfMgmt_Periodic_Nodes_Node.
    Node []PerfMgmt_Periodic_Nodes_Node
}

func (nodes *PerfMgmt_Periodic_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Periodic_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Periodic_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Periodic_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Periodic_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Periodic_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Periodic_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Periodic_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Periodic_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Periodic_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Periodic_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Periodic_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Periodic_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Periodic_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Periodic_Nodes) GetParentYangName() string { return "periodic" }

// PerfMgmt_Periodic_Nodes_Node
// Node Instance
type PerfMgmt_Periodic_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Node CPU data.
    SampleXr PerfMgmt_Periodic_Nodes_Node_SampleXr

    // Processes data.
    Processes PerfMgmt_Periodic_Nodes_Node_Processes

    // Node Memory data.
    Samples PerfMgmt_Periodic_Nodes_Node_Samples
}

func (node *PerfMgmt_Periodic_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Periodic_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Periodic_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "sample-xr" { return "SampleXr" }
    if yname == "processes" { return "Processes" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (node *PerfMgmt_Periodic_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Periodic_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample-xr" {
        return &node.SampleXr
    }
    if childYangName == "processes" {
        return &node.Processes
    }
    if childYangName == "samples" {
        return &node.Samples
    }
    return nil
}

func (node *PerfMgmt_Periodic_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sample-xr"] = &node.SampleXr
    children["processes"] = &node.Processes
    children["samples"] = &node.Samples
    return children
}

func (node *PerfMgmt_Periodic_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    return leafs
}

func (node *PerfMgmt_Periodic_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Periodic_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Periodic_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Periodic_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Periodic_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Periodic_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Periodic_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Periodic_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Periodic_Nodes_Node_SampleXr
// Node CPU data
type PerfMgmt_Periodic_Nodes_Node_SampleXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node CPU data sample. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample.
    Sample []PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample
}

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetFilter() yfilter.YFilter { return sampleXr.YFilter }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) SetFilter(yf yfilter.YFilter) { sampleXr.YFilter = yf }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetSegmentPath() string {
    return "sample-xr"
}

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range sampleXr.Sample {
            if sampleXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample{}
        sampleXr.Sample = append(sampleXr.Sample, child)
        return &sampleXr.Sample[len(sampleXr.Sample)-1]
    }
    return nil
}

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sampleXr.Sample {
        children[sampleXr.Sample[i].GetSegmentPath()] = &sampleXr.Sample[i]
    }
    return children
}

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetBundleName() string { return "cisco_ios_xr" }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetYangName() string { return "sample-xr" }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) SetParent(parent types.Entity) { sampleXr.parent = parent }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetParent() types.Entity { return sampleXr.parent }

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetParentYangName() string { return "node" }

// PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample
// Node CPU data sample
type PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Number of processes in the system. The type is interface{} with range:
    // 0..4294967295.
    NoProcesses interface{}

    // Average system %CPU utilization. The type is interface{} with range:
    // 0..4294967295.
    AverageCpuUsed interface{}
}

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "no-processes" { return "NoProcesses" }
    if yname == "average-cpu-used" { return "AverageCpuUsed" }
    return ""
}

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["no-processes"] = sample.NoProcesses
    leafs["average-cpu-used"] = sample.AverageCpuUsed
    return leafs
}

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetParentYangName() string { return "sample-xr" }

// PerfMgmt_Periodic_Nodes_Node_Processes
// Processes data
type PerfMgmt_Periodic_Nodes_Node_Processes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process data. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_Processes_Process.
    Process []PerfMgmt_Periodic_Nodes_Node_Processes_Process
}

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetFilter() yfilter.YFilter { return processes.YFilter }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) SetFilter(yf yfilter.YFilter) { processes.YFilter = yf }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    return ""
}

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetSegmentPath() string {
    return "processes"
}

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process" {
        for _, c := range processes.Process {
            if processes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Nodes_Node_Processes_Process{}
        processes.Process = append(processes.Process, child)
        return &processes.Process[len(processes.Process)-1]
    }
    return nil
}

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range processes.Process {
        children[processes.Process[i].GetSegmentPath()] = &processes.Process[i]
    }
    return children
}

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetBundleName() string { return "cisco_ios_xr" }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetYangName() string { return "processes" }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) SetParent(parent types.Entity) { processes.parent = parent }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetParent() types.Entity { return processes.parent }

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetParentYangName() string { return "node" }

// PerfMgmt_Periodic_Nodes_Node_Processes_Process
// Process data
type PerfMgmt_Periodic_Nodes_Node_Processes_Process struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Process data.
    Samples PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples
}

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetFilter() yfilter.YFilter { return process.YFilter }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) SetFilter(yf yfilter.YFilter) { process.YFilter = yf }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetGoName(yname string) string {
    if yname == "process-id" { return "ProcessId" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetSegmentPath() string {
    return "process" + "[process-id='" + fmt.Sprintf("%v", process.ProcessId) + "']"
}

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &process.Samples
    }
    return nil
}

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &process.Samples
    return children
}

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-id"] = process.ProcessId
    return leafs
}

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetBundleName() string { return "cisco_ios_xr" }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetYangName() string { return "process" }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) SetParent(parent types.Entity) { process.parent = parent }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetParent() types.Entity { return process.parent }

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetParentYangName() string { return "processes" }

// PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples
// Process data
type PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process data sample. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample.
    Sample []PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetParentYangName() string { return "process" }

// PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample
// Process data sample
type PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Max. dynamic memory (KBytes) used since startup time. The type is
    // interface{} with range: 0..4294967295. Units are kilobyte.
    PeakMemory interface{}

    // Average %CPU utilization. The type is interface{} with range:
    // 0..4294967295.
    AverageCpuUsed interface{}

    // Number of threads. The type is interface{} with range: 0..4294967295.
    NoThreads interface{}
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "peak-memory" { return "PeakMemory" }
    if yname == "average-cpu-used" { return "AverageCpuUsed" }
    if yname == "no-threads" { return "NoThreads" }
    return ""
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["peak-memory"] = sample.PeakMemory
    leafs["average-cpu-used"] = sample.AverageCpuUsed
    leafs["no-threads"] = sample.NoThreads
    return leafs
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Nodes_Node_Samples
// Node Memory data
type PerfMgmt_Periodic_Nodes_Node_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Memory data sample. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_Samples_Sample.
    Sample []PerfMgmt_Periodic_Nodes_Node_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Nodes_Node_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetParentYangName() string { return "node" }

// PerfMgmt_Periodic_Nodes_Node_Samples_Sample
// Node Memory data sample
type PerfMgmt_Periodic_Nodes_Node_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Current application memory (Bytes) in use. The type is interface{} with
    // range: 0..4294967295. Units are byte.
    CurrMemory interface{}

    // Max. system memory (MBytes) used since bootup. The type is interface{} with
    // range: 0..4294967295. Units are megabyte.
    PeakMemory interface{}
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "curr-memory" { return "CurrMemory" }
    if yname == "peak-memory" { return "PeakMemory" }
    return ""
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["curr-memory"] = sample.CurrMemory
    leafs["peak-memory"] = sample.PeakMemory
    return leafs
}

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Bgp
// Collected BGP data
type PerfMgmt_Periodic_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbors for which statistics are collected.
    BgpNeighbors PerfMgmt_Periodic_Bgp_BgpNeighbors
}

func (bgp *PerfMgmt_Periodic_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PerfMgmt_Periodic_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PerfMgmt_Periodic_Bgp) GetGoName(yname string) string {
    if yname == "bgp-neighbors" { return "BgpNeighbors" }
    return ""
}

func (bgp *PerfMgmt_Periodic_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PerfMgmt_Periodic_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-neighbors" {
        return &bgp.BgpNeighbors
    }
    return nil
}

func (bgp *PerfMgmt_Periodic_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-neighbors"] = &bgp.BgpNeighbors
    return children
}

func (bgp *PerfMgmt_Periodic_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *PerfMgmt_Periodic_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PerfMgmt_Periodic_Bgp) GetYangName() string { return "bgp" }

func (bgp *PerfMgmt_Periodic_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PerfMgmt_Periodic_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PerfMgmt_Periodic_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PerfMgmt_Periodic_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PerfMgmt_Periodic_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PerfMgmt_Periodic_Bgp) GetParentYangName() string { return "periodic" }

// PerfMgmt_Periodic_Bgp_BgpNeighbors
// Neighbors for which statistics are collected
type PerfMgmt_Periodic_Bgp_BgpNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for particular neighbor. The type is slice of
    // PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor.
    BgpNeighbor []PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor
}

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetFilter() yfilter.YFilter { return bgpNeighbors.YFilter }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) SetFilter(yf yfilter.YFilter) { bgpNeighbors.YFilter = yf }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetGoName(yname string) string {
    if yname == "bgp-neighbor" { return "BgpNeighbor" }
    return ""
}

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetSegmentPath() string {
    return "bgp-neighbors"
}

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-neighbor" {
        for _, c := range bgpNeighbors.BgpNeighbor {
            if bgpNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor{}
        bgpNeighbors.BgpNeighbor = append(bgpNeighbors.BgpNeighbor, child)
        return &bgpNeighbors.BgpNeighbor[len(bgpNeighbors.BgpNeighbor)-1]
    }
    return nil
}

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpNeighbors.BgpNeighbor {
        children[bgpNeighbors.BgpNeighbor[i].GetSegmentPath()] = &bgpNeighbors.BgpNeighbor[i]
    }
    return children
}

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetBundleName() string { return "cisco_ios_xr" }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetYangName() string { return "bgp-neighbors" }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) SetParent(parent types.Entity) { bgpNeighbors.parent = parent }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetParent() types.Entity { return bgpNeighbors.parent }

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetParentYangName() string { return "bgp" }

// PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor
// Samples for particular neighbor
type PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. BGP Neighbor Identifier. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Sample Table for a BGP neighbor.
    Samples PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples
}

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetFilter() yfilter.YFilter { return bgpNeighbor.YFilter }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) SetFilter(yf yfilter.YFilter) { bgpNeighbor.YFilter = yf }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetSegmentPath() string {
    return "bgp-neighbor" + "[ip-address='" + fmt.Sprintf("%v", bgpNeighbor.IpAddress) + "']"
}

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &bgpNeighbor.Samples
    }
    return nil
}

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &bgpNeighbor.Samples
    return children
}

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = bgpNeighbor.IpAddress
    return leafs
}

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetYangName() string { return "bgp-neighbor" }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) SetParent(parent types.Entity) { bgpNeighbor.parent = parent }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetParent() types.Entity { return bgpNeighbor.parent }

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetParentYangName() string { return "bgp-neighbors" }

// PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples
// Sample Table for a BGP neighbor
type PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor statistics sample. The type is slice of
    // PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample.
    Sample []PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetParentYangName() string { return "bgp-neighbor" }

// PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
// Neighbor statistics sample
type PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Number of messages received. The type is interface{} with range:
    // 0..4294967295.
    InputMessages interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    OutputMessages interface{}

    // Number of update messages received. The type is interface{} with range:
    // 0..4294967295.
    InputUpdateMessages interface{}

    // Number of update messages sent. The type is interface{} with range:
    // 0..4294967295.
    OutputUpdateMessages interface{}

    // Number of times the connection was established. The type is interface{}
    // with range: 0..4294967295.
    ConnEstablished interface{}

    // Number of times connection was dropped. The type is interface{} with range:
    // 0..4294967295.
    ConnDropped interface{}

    // Number of error notifications received on the connection. The type is
    // interface{} with range: 0..4294967295.
    ErrorsReceived interface{}

    // Number of error notifications sent on the connection. The type is
    // interface{} with range: 0..4294967295.
    ErrorsSent interface{}
}

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-messages" { return "InputMessages" }
    if yname == "output-messages" { return "OutputMessages" }
    if yname == "input-update-messages" { return "InputUpdateMessages" }
    if yname == "output-update-messages" { return "OutputUpdateMessages" }
    if yname == "conn-established" { return "ConnEstablished" }
    if yname == "conn-dropped" { return "ConnDropped" }
    if yname == "errors-received" { return "ErrorsReceived" }
    if yname == "errors-sent" { return "ErrorsSent" }
    return ""
}

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-messages"] = sample.InputMessages
    leafs["output-messages"] = sample.OutputMessages
    leafs["input-update-messages"] = sample.InputUpdateMessages
    leafs["output-update-messages"] = sample.OutputUpdateMessages
    leafs["conn-established"] = sample.ConnEstablished
    leafs["conn-dropped"] = sample.ConnDropped
    leafs["errors-received"] = sample.ErrorsReceived
    leafs["errors-sent"] = sample.ErrorsSent
    return leafs
}

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Interface
// Collected Interface data
type PerfMgmt_Periodic_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interfaces for which Generic Counters are collected.
    GenericCounterInterfaces PerfMgmt_Periodic_Interface_GenericCounterInterfaces

    // Interfaces for which Basic Counters are collected.
    BasicCounterInterfaces PerfMgmt_Periodic_Interface_BasicCounterInterfaces

    // Interfaces for which Data Rates are collected.
    DataRateInterfaces PerfMgmt_Periodic_Interface_DataRateInterfaces
}

func (self *PerfMgmt_Periodic_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PerfMgmt_Periodic_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PerfMgmt_Periodic_Interface) GetGoName(yname string) string {
    if yname == "generic-counter-interfaces" { return "GenericCounterInterfaces" }
    if yname == "basic-counter-interfaces" { return "BasicCounterInterfaces" }
    if yname == "data-rate-interfaces" { return "DataRateInterfaces" }
    return ""
}

func (self *PerfMgmt_Periodic_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *PerfMgmt_Periodic_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interfaces" {
        return &self.GenericCounterInterfaces
    }
    if childYangName == "basic-counter-interfaces" {
        return &self.BasicCounterInterfaces
    }
    if childYangName == "data-rate-interfaces" {
        return &self.DataRateInterfaces
    }
    return nil
}

func (self *PerfMgmt_Periodic_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["generic-counter-interfaces"] = &self.GenericCounterInterfaces
    children["basic-counter-interfaces"] = &self.BasicCounterInterfaces
    children["data-rate-interfaces"] = &self.DataRateInterfaces
    return children
}

func (self *PerfMgmt_Periodic_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *PerfMgmt_Periodic_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *PerfMgmt_Periodic_Interface) GetYangName() string { return "interface" }

func (self *PerfMgmt_Periodic_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PerfMgmt_Periodic_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PerfMgmt_Periodic_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PerfMgmt_Periodic_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *PerfMgmt_Periodic_Interface) GetParent() types.Entity { return self.parent }

func (self *PerfMgmt_Periodic_Interface) GetParentYangName() string { return "periodic" }

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces
// Interfaces for which Generic Counters are
// collected
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface.
    GenericCounterInterface []PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface
}

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetFilter() yfilter.YFilter { return genericCounterInterfaces.YFilter }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) SetFilter(yf yfilter.YFilter) { genericCounterInterfaces.YFilter = yf }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetGoName(yname string) string {
    if yname == "generic-counter-interface" { return "GenericCounterInterface" }
    return ""
}

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetSegmentPath() string {
    return "generic-counter-interfaces"
}

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interface" {
        for _, c := range genericCounterInterfaces.GenericCounterInterface {
            if genericCounterInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface{}
        genericCounterInterfaces.GenericCounterInterface = append(genericCounterInterfaces.GenericCounterInterface, child)
        return &genericCounterInterfaces.GenericCounterInterface[len(genericCounterInterfaces.GenericCounterInterface)-1]
    }
    return nil
}

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range genericCounterInterfaces.GenericCounterInterface {
        children[genericCounterInterfaces.GenericCounterInterface[i].GetSegmentPath()] = &genericCounterInterfaces.GenericCounterInterface[i]
    }
    return children
}

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetYangName() string { return "generic-counter-interfaces" }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) SetParent(parent types.Entity) { genericCounterInterfaces.parent = parent }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetParent() types.Entity { return genericCounterInterfaces.parent }

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetParentYangName() string { return "interface" }

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface
// Samples for a particular interface
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Generic Counter samples for an interface.
    Samples PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
}

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetFilter() yfilter.YFilter { return genericCounterInterface.YFilter }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) SetFilter(yf yfilter.YFilter) { genericCounterInterface.YFilter = yf }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetSegmentPath() string {
    return "generic-counter-interface" + "[interface-name='" + fmt.Sprintf("%v", genericCounterInterface.InterfaceName) + "']"
}

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &genericCounterInterface.Samples
    }
    return nil
}

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &genericCounterInterface.Samples
    return children
}

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = genericCounterInterface.InterfaceName
    return leafs
}

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetYangName() string { return "generic-counter-interface" }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) SetParent(parent types.Entity) { genericCounterInterface.parent = parent }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetParent() types.Entity { return genericCounterInterface.parent }

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetParentYangName() string { return "generic-counter-interfaces" }

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
// Generic Counter samples for an interface
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample.
    Sample []PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetParentYangName() string { return "generic-counter-interface" }

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
// Generic Counters sample
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InPackets interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    InOctets interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    OutPackets interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    OutOctets interface{}

    // Unicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPkts interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InMulticastPkts interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InBroadcastPkts interface{}

    // Unicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPkts interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMulticastPkts interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBroadcastPkts interface{}

    // Outbound correct packets discarded. The type is interface{} with range:
    // 0..4294967295.
    OutputTotalDrops interface{}

    // Inbound correct packets discarded. The type is interface{} with range:
    // 0..4294967295.
    InputTotalDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Inbound packets discarded with unknown proto. The type is interface{} with
    // range: 0..4294967295.
    InputUnknownProto interface{}

    // Outbound incorrect packets discarded. The type is interface{} with range:
    // 0..4294967295.
    OutputTotalErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderrun interface{}

    // Inbound incorrect packets discarded. The type is interface{} with range:
    // 0..4294967295.
    InputTotalErrors interface{}

    // Inbound packets discarded with incorrect CRC. The type is interface{} with
    // range: 0..4294967295.
    InputCrc interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverrun interface{}

    // Inbound framing errors. The type is interface{} with range: 0..4294967295.
    InputFrame interface{}
}

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "in-packets" { return "InPackets" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "out-packets" { return "OutPackets" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "in-ucast-pkts" { return "InUcastPkts" }
    if yname == "in-multicast-pkts" { return "InMulticastPkts" }
    if yname == "in-broadcast-pkts" { return "InBroadcastPkts" }
    if yname == "out-ucast-pkts" { return "OutUcastPkts" }
    if yname == "out-multicast-pkts" { return "OutMulticastPkts" }
    if yname == "out-broadcast-pkts" { return "OutBroadcastPkts" }
    if yname == "output-total-drops" { return "OutputTotalDrops" }
    if yname == "input-total-drops" { return "InputTotalDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "input-unknown-proto" { return "InputUnknownProto" }
    if yname == "output-total-errors" { return "OutputTotalErrors" }
    if yname == "output-underrun" { return "OutputUnderrun" }
    if yname == "input-total-errors" { return "InputTotalErrors" }
    if yname == "input-crc" { return "InputCrc" }
    if yname == "input-overrun" { return "InputOverrun" }
    if yname == "input-frame" { return "InputFrame" }
    return ""
}

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["in-packets"] = sample.InPackets
    leafs["in-octets"] = sample.InOctets
    leafs["out-packets"] = sample.OutPackets
    leafs["out-octets"] = sample.OutOctets
    leafs["in-ucast-pkts"] = sample.InUcastPkts
    leafs["in-multicast-pkts"] = sample.InMulticastPkts
    leafs["in-broadcast-pkts"] = sample.InBroadcastPkts
    leafs["out-ucast-pkts"] = sample.OutUcastPkts
    leafs["out-multicast-pkts"] = sample.OutMulticastPkts
    leafs["out-broadcast-pkts"] = sample.OutBroadcastPkts
    leafs["output-total-drops"] = sample.OutputTotalDrops
    leafs["input-total-drops"] = sample.InputTotalDrops
    leafs["input-queue-drops"] = sample.InputQueueDrops
    leafs["input-unknown-proto"] = sample.InputUnknownProto
    leafs["output-total-errors"] = sample.OutputTotalErrors
    leafs["output-underrun"] = sample.OutputUnderrun
    leafs["input-total-errors"] = sample.InputTotalErrors
    leafs["input-crc"] = sample.InputCrc
    leafs["input-overrun"] = sample.InputOverrun
    leafs["input-frame"] = sample.InputFrame
    return leafs
}

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces
// Interfaces for which Basic Counters are
// collected
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface.
    BasicCounterInterface []PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface
}

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetFilter() yfilter.YFilter { return basicCounterInterfaces.YFilter }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) SetFilter(yf yfilter.YFilter) { basicCounterInterfaces.YFilter = yf }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetGoName(yname string) string {
    if yname == "basic-counter-interface" { return "BasicCounterInterface" }
    return ""
}

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetSegmentPath() string {
    return "basic-counter-interfaces"
}

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-counter-interface" {
        for _, c := range basicCounterInterfaces.BasicCounterInterface {
            if basicCounterInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface{}
        basicCounterInterfaces.BasicCounterInterface = append(basicCounterInterfaces.BasicCounterInterface, child)
        return &basicCounterInterfaces.BasicCounterInterface[len(basicCounterInterfaces.BasicCounterInterface)-1]
    }
    return nil
}

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range basicCounterInterfaces.BasicCounterInterface {
        children[basicCounterInterfaces.BasicCounterInterface[i].GetSegmentPath()] = &basicCounterInterfaces.BasicCounterInterface[i]
    }
    return children
}

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetYangName() string { return "basic-counter-interfaces" }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) SetParent(parent types.Entity) { basicCounterInterfaces.parent = parent }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetParent() types.Entity { return basicCounterInterfaces.parent }

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetParentYangName() string { return "interface" }

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface
// Samples for a particular interface
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Basic Counter samples for an interface.
    Samples PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
}

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetFilter() yfilter.YFilter { return basicCounterInterface.YFilter }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) SetFilter(yf yfilter.YFilter) { basicCounterInterface.YFilter = yf }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetSegmentPath() string {
    return "basic-counter-interface" + "[interface-name='" + fmt.Sprintf("%v", basicCounterInterface.InterfaceName) + "']"
}

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &basicCounterInterface.Samples
    }
    return nil
}

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &basicCounterInterface.Samples
    return children
}

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = basicCounterInterface.InterfaceName
    return leafs
}

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetYangName() string { return "basic-counter-interface" }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) SetParent(parent types.Entity) { basicCounterInterface.parent = parent }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetParent() types.Entity { return basicCounterInterface.parent }

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetParentYangName() string { return "basic-counter-interfaces" }

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
// Basic Counter samples for an interface
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Basic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample.
    Sample []PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetParentYangName() string { return "basic-counter-interface" }

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
// Basic Counters sample
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds from UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InPackets interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    InOctets interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    OutPackets interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    OutOctets interface{}

    // Inbound correct packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    InputTotalDrops interface{}

    // Input queue drops. The type is interface{} with range:
    // 0..18446744073709551615.
    InputQueueDrops interface{}

    // Inbound incorrect packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    InputTotalErrors interface{}

    // Outbound correct packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputTotalDrops interface{}

    // Output queue drops. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputQueueDrops interface{}

    // Outbound incorrect packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputTotalErrors interface{}
}

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "in-packets" { return "InPackets" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "out-packets" { return "OutPackets" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "input-total-drops" { return "InputTotalDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "input-total-errors" { return "InputTotalErrors" }
    if yname == "output-total-drops" { return "OutputTotalDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "output-total-errors" { return "OutputTotalErrors" }
    return ""
}

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["in-packets"] = sample.InPackets
    leafs["in-octets"] = sample.InOctets
    leafs["out-packets"] = sample.OutPackets
    leafs["out-octets"] = sample.OutOctets
    leafs["input-total-drops"] = sample.InputTotalDrops
    leafs["input-queue-drops"] = sample.InputQueueDrops
    leafs["input-total-errors"] = sample.InputTotalErrors
    leafs["output-total-drops"] = sample.OutputTotalDrops
    leafs["output-queue-drops"] = sample.OutputQueueDrops
    leafs["output-total-errors"] = sample.OutputTotalErrors
    return leafs
}

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Periodic_Interface_DataRateInterfaces
// Interfaces for which Data Rates are collected
type PerfMgmt_Periodic_Interface_DataRateInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface.
    DataRateInterface []PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface
}

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetFilter() yfilter.YFilter { return dataRateInterfaces.YFilter }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) SetFilter(yf yfilter.YFilter) { dataRateInterfaces.YFilter = yf }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetGoName(yname string) string {
    if yname == "data-rate-interface" { return "DataRateInterface" }
    return ""
}

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetSegmentPath() string {
    return "data-rate-interfaces"
}

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-rate-interface" {
        for _, c := range dataRateInterfaces.DataRateInterface {
            if dataRateInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface{}
        dataRateInterfaces.DataRateInterface = append(dataRateInterfaces.DataRateInterface, child)
        return &dataRateInterfaces.DataRateInterface[len(dataRateInterfaces.DataRateInterface)-1]
    }
    return nil
}

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dataRateInterfaces.DataRateInterface {
        children[dataRateInterfaces.DataRateInterface[i].GetSegmentPath()] = &dataRateInterfaces.DataRateInterface[i]
    }
    return children
}

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetYangName() string { return "data-rate-interfaces" }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) SetParent(parent types.Entity) { dataRateInterfaces.parent = parent }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetParent() types.Entity { return dataRateInterfaces.parent }

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetParentYangName() string { return "interface" }

// PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface
// Samples for a particular interface
type PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Data Rate samples for an interface.
    Samples PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples
}

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetFilter() yfilter.YFilter { return dataRateInterface.YFilter }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) SetFilter(yf yfilter.YFilter) { dataRateInterface.YFilter = yf }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetSegmentPath() string {
    return "data-rate-interface" + "[interface-name='" + fmt.Sprintf("%v", dataRateInterface.InterfaceName) + "']"
}

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &dataRateInterface.Samples
    }
    return nil
}

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &dataRateInterface.Samples
    return children
}

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = dataRateInterface.InterfaceName
    return leafs
}

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetYangName() string { return "data-rate-interface" }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) SetParent(parent types.Entity) { dataRateInterface.parent = parent }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetParent() types.Entity { return dataRateInterface.parent }

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetParentYangName() string { return "data-rate-interfaces" }

// PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples
// Data Rate samples for an interface
type PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data Rates sample. The type is slice of
    // PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample.
    Sample []PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetParentYangName() string { return "data-rate-interface" }

// PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
// Data Rates sample
type PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Input datarate in 1000's of bps. The type is interface{} with range:
    // 0..4294967295. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..4294967295. Units are packet/s.
    InputPacketRate interface{}

    // Output datarate in 1000's of bps. The type is interface{} with range:
    // 0..4294967295. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..4294967295. Units are packet/s.
    OutputPacketRate interface{}

    // Peak input datarate. The type is interface{} with range: 0..4294967295.
    InputPeakRate interface{}

    // Peak input packet rate. The type is interface{} with range: 0..4294967295.
    InputPeakPkts interface{}

    // Peak output datarate. The type is interface{} with range: 0..4294967295.
    OutputPeakRate interface{}

    // Peak output packet rate. The type is interface{} with range: 0..4294967295.
    OutputPeakPkts interface{}

    // Bandwidth (in kbps). The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    Bandwidth interface{}
}

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "input-peak-rate" { return "InputPeakRate" }
    if yname == "input-peak-pkts" { return "InputPeakPkts" }
    if yname == "output-peak-rate" { return "OutputPeakRate" }
    if yname == "output-peak-pkts" { return "OutputPeakPkts" }
    if yname == "bandwidth" { return "Bandwidth" }
    return ""
}

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-data-rate"] = sample.InputDataRate
    leafs["input-packet-rate"] = sample.InputPacketRate
    leafs["output-data-rate"] = sample.OutputDataRate
    leafs["output-packet-rate"] = sample.OutputPacketRate
    leafs["input-peak-rate"] = sample.InputPeakRate
    leafs["input-peak-pkts"] = sample.InputPeakPkts
    leafs["output-peak-rate"] = sample.OutputPeakRate
    leafs["output-peak-pkts"] = sample.OutputPeakPkts
    leafs["bandwidth"] = sample.Bandwidth
    return leafs
}

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor
// Data from monitor (one history period) requests
type PerfMgmt_Monitor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collected OSPF data.
    Ospf PerfMgmt_Monitor_Ospf

    // Collected MPLS data.
    Mpls PerfMgmt_Monitor_Mpls

    // Nodes for which data is collected.
    Nodes PerfMgmt_Monitor_Nodes

    // Collected BGP data.
    Bgp PerfMgmt_Monitor_Bgp

    // Collected Interface data.
    Interface PerfMgmt_Monitor_Interface
}

func (monitor *PerfMgmt_Monitor) GetFilter() yfilter.YFilter { return monitor.YFilter }

func (monitor *PerfMgmt_Monitor) SetFilter(yf yfilter.YFilter) { monitor.YFilter = yf }

func (monitor *PerfMgmt_Monitor) GetGoName(yname string) string {
    if yname == "ospf" { return "Ospf" }
    if yname == "mpls" { return "Mpls" }
    if yname == "nodes" { return "Nodes" }
    if yname == "bgp" { return "Bgp" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (monitor *PerfMgmt_Monitor) GetSegmentPath() string {
    return "monitor"
}

func (monitor *PerfMgmt_Monitor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf" {
        return &monitor.Ospf
    }
    if childYangName == "mpls" {
        return &monitor.Mpls
    }
    if childYangName == "nodes" {
        return &monitor.Nodes
    }
    if childYangName == "bgp" {
        return &monitor.Bgp
    }
    if childYangName == "interface" {
        return &monitor.Interface
    }
    return nil
}

func (monitor *PerfMgmt_Monitor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospf"] = &monitor.Ospf
    children["mpls"] = &monitor.Mpls
    children["nodes"] = &monitor.Nodes
    children["bgp"] = &monitor.Bgp
    children["interface"] = &monitor.Interface
    return children
}

func (monitor *PerfMgmt_Monitor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (monitor *PerfMgmt_Monitor) GetBundleName() string { return "cisco_ios_xr" }

func (monitor *PerfMgmt_Monitor) GetYangName() string { return "monitor" }

func (monitor *PerfMgmt_Monitor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (monitor *PerfMgmt_Monitor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (monitor *PerfMgmt_Monitor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (monitor *PerfMgmt_Monitor) SetParent(parent types.Entity) { monitor.parent = parent }

func (monitor *PerfMgmt_Monitor) GetParent() types.Entity { return monitor.parent }

func (monitor *PerfMgmt_Monitor) GetParentYangName() string { return "perf-mgmt" }

// PerfMgmt_Monitor_Ospf
// Collected OSPF data
type PerfMgmt_Monitor_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF v2 instances for which protocol statistics are collected.
    Ospfv2ProtocolInstances PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances

    // OSPF v3 instances for which protocol statistics are collected.
    Ospfv3ProtocolInstances PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances
}

func (ospf *PerfMgmt_Monitor_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PerfMgmt_Monitor_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PerfMgmt_Monitor_Ospf) GetGoName(yname string) string {
    if yname == "ospfv2-protocol-instances" { return "Ospfv2ProtocolInstances" }
    if yname == "ospfv3-protocol-instances" { return "Ospfv3ProtocolInstances" }
    return ""
}

func (ospf *PerfMgmt_Monitor_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PerfMgmt_Monitor_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2-protocol-instances" {
        return &ospf.Ospfv2ProtocolInstances
    }
    if childYangName == "ospfv3-protocol-instances" {
        return &ospf.Ospfv3ProtocolInstances
    }
    return nil
}

func (ospf *PerfMgmt_Monitor_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv2-protocol-instances"] = &ospf.Ospfv2ProtocolInstances
    children["ospfv3-protocol-instances"] = &ospf.Ospfv3ProtocolInstances
    return children
}

func (ospf *PerfMgmt_Monitor_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospf *PerfMgmt_Monitor_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PerfMgmt_Monitor_Ospf) GetYangName() string { return "ospf" }

func (ospf *PerfMgmt_Monitor_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PerfMgmt_Monitor_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PerfMgmt_Monitor_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PerfMgmt_Monitor_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PerfMgmt_Monitor_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PerfMgmt_Monitor_Ospf) GetParentYangName() string { return "monitor" }

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances
// OSPF v2 instances for which protocol statistics
// are collected
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v2 instance. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance.
    Ospfv2ProtocolInstance []PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
}

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetFilter() yfilter.YFilter { return ospfv2ProtocolInstances.YFilter }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) SetFilter(yf yfilter.YFilter) { ospfv2ProtocolInstances.YFilter = yf }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetGoName(yname string) string {
    if yname == "ospfv2-protocol-instance" { return "Ospfv2ProtocolInstance" }
    return ""
}

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetSegmentPath() string {
    return "ospfv2-protocol-instances"
}

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2-protocol-instance" {
        for _, c := range ospfv2ProtocolInstances.Ospfv2ProtocolInstance {
            if ospfv2ProtocolInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance{}
        ospfv2ProtocolInstances.Ospfv2ProtocolInstance = append(ospfv2ProtocolInstances.Ospfv2ProtocolInstance, child)
        return &ospfv2ProtocolInstances.Ospfv2ProtocolInstance[len(ospfv2ProtocolInstances.Ospfv2ProtocolInstance)-1]
    }
    return nil
}

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfv2ProtocolInstances.Ospfv2ProtocolInstance {
        children[ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i].GetSegmentPath()] = &ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i]
    }
    return children
}

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetYangName() string { return "ospfv2-protocol-instances" }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) SetParent(parent types.Entity) { ospfv2ProtocolInstances.parent = parent }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetParent() types.Entity { return ospfv2ProtocolInstances.parent }

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetParentYangName() string { return "ospf" }

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
// Protocol samples for a particular OSPF v2
// instance
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v2 instance.
    Samples PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
}

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetFilter() yfilter.YFilter { return ospfv2ProtocolInstance.YFilter }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) SetFilter(yf yfilter.YFilter) { ospfv2ProtocolInstance.YFilter = yf }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetSegmentPath() string {
    return "ospfv2-protocol-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfv2ProtocolInstance.InstanceName) + "']"
}

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &ospfv2ProtocolInstance.Samples
    }
    return nil
}

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &ospfv2ProtocolInstance.Samples
    return children
}

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfv2ProtocolInstance.InstanceName
    return leafs
}

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetYangName() string { return "ospfv2-protocol-instance" }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) SetParent(parent types.Entity) { ospfv2ProtocolInstance.parent = parent }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetParent() types.Entity { return ospfv2ProtocolInstance.parent }

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetParentYangName() string { return "ospfv2-protocol-instances" }

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
// Sample Table for an OSPV v2 instance
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample.
    Sample []PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetParentYangName() string { return "ospfv2-protocol-instance" }

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Total number of packets received. The type is interface{} with range:
    // 0..4294967295.
    InputPackets interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputPackets interface{}

    // Number of Hello packets received. The type is interface{} with range:
    // 0..4294967295.
    InputHelloPackets interface{}

    // Number of Hello packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputHelloPackets interface{}

    // Number of DBD packets received. The type is interface{} with range:
    // 0..4294967295.
    InputDbDs interface{}

    // Number of LSA received in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    InputDbDsLsa interface{}

    // Number of DBD packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDs interface{}

    // Number of LSA sent in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDsLsa interface{}

    // Number of LS Requests received. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequests interface{}

    // Number of LSA received in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequestsLsa interface{}

    // Number of LS Requests sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequests interface{}

    // Number of LSA sent in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequestsLsa interface{}

    // Number of LSA Updates received. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdates interface{}

    // Number of LSA received in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdatesLsa interface{}

    // Number of LSA Updates sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdates interface{}

    // Number of LSA sent in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdatesLsa interface{}

    // Number of LSA Acknowledgements received. The type is interface{} with
    // range: 0..4294967295.
    InputLsaAcks interface{}

    // Number of LSA received in LSA Acknowledgements. The type is interface{}
    // with range: 0..4294967295.
    InputLsaAcksLsa interface{}

    // Number of LSA Acknowledgements sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaAcks interface{}

    // Number of LSA sent in LSA Acknowledgements. The type is interface{} with
    // range: 0..4294967295.
    OutputLsaAcksLsa interface{}

    // Number of packets received with checksum errors. The type is interface{}
    // with range: 0..4294967295.
    ChecksumErrors interface{}
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-packets" { return "InputPackets" }
    if yname == "output-packets" { return "OutputPackets" }
    if yname == "input-hello-packets" { return "InputHelloPackets" }
    if yname == "output-hello-packets" { return "OutputHelloPackets" }
    if yname == "input-db-ds" { return "InputDbDs" }
    if yname == "input-db-ds-lsa" { return "InputDbDsLsa" }
    if yname == "output-db-ds" { return "OutputDbDs" }
    if yname == "output-db-ds-lsa" { return "OutputDbDsLsa" }
    if yname == "input-ls-requests" { return "InputLsRequests" }
    if yname == "input-ls-requests-lsa" { return "InputLsRequestsLsa" }
    if yname == "output-ls-requests" { return "OutputLsRequests" }
    if yname == "output-ls-requests-lsa" { return "OutputLsRequestsLsa" }
    if yname == "input-lsa-updates" { return "InputLsaUpdates" }
    if yname == "input-lsa-updates-lsa" { return "InputLsaUpdatesLsa" }
    if yname == "output-lsa-updates" { return "OutputLsaUpdates" }
    if yname == "output-lsa-updates-lsa" { return "OutputLsaUpdatesLsa" }
    if yname == "input-lsa-acks" { return "InputLsaAcks" }
    if yname == "input-lsa-acks-lsa" { return "InputLsaAcksLsa" }
    if yname == "output-lsa-acks" { return "OutputLsaAcks" }
    if yname == "output-lsa-acks-lsa" { return "OutputLsaAcksLsa" }
    if yname == "checksum-errors" { return "ChecksumErrors" }
    return ""
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-packets"] = sample.InputPackets
    leafs["output-packets"] = sample.OutputPackets
    leafs["input-hello-packets"] = sample.InputHelloPackets
    leafs["output-hello-packets"] = sample.OutputHelloPackets
    leafs["input-db-ds"] = sample.InputDbDs
    leafs["input-db-ds-lsa"] = sample.InputDbDsLsa
    leafs["output-db-ds"] = sample.OutputDbDs
    leafs["output-db-ds-lsa"] = sample.OutputDbDsLsa
    leafs["input-ls-requests"] = sample.InputLsRequests
    leafs["input-ls-requests-lsa"] = sample.InputLsRequestsLsa
    leafs["output-ls-requests"] = sample.OutputLsRequests
    leafs["output-ls-requests-lsa"] = sample.OutputLsRequestsLsa
    leafs["input-lsa-updates"] = sample.InputLsaUpdates
    leafs["input-lsa-updates-lsa"] = sample.InputLsaUpdatesLsa
    leafs["output-lsa-updates"] = sample.OutputLsaUpdates
    leafs["output-lsa-updates-lsa"] = sample.OutputLsaUpdatesLsa
    leafs["input-lsa-acks"] = sample.InputLsaAcks
    leafs["input-lsa-acks-lsa"] = sample.InputLsaAcksLsa
    leafs["output-lsa-acks"] = sample.OutputLsaAcks
    leafs["output-lsa-acks-lsa"] = sample.OutputLsaAcksLsa
    leafs["checksum-errors"] = sample.ChecksumErrors
    return leafs
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances
// OSPF v3 instances for which protocol statistics
// are collected
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v3 instance. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance.
    Ospfv3ProtocolInstance []PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
}

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetFilter() yfilter.YFilter { return ospfv3ProtocolInstances.YFilter }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) SetFilter(yf yfilter.YFilter) { ospfv3ProtocolInstances.YFilter = yf }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetGoName(yname string) string {
    if yname == "ospfv3-protocol-instance" { return "Ospfv3ProtocolInstance" }
    return ""
}

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetSegmentPath() string {
    return "ospfv3-protocol-instances"
}

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv3-protocol-instance" {
        for _, c := range ospfv3ProtocolInstances.Ospfv3ProtocolInstance {
            if ospfv3ProtocolInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance{}
        ospfv3ProtocolInstances.Ospfv3ProtocolInstance = append(ospfv3ProtocolInstances.Ospfv3ProtocolInstance, child)
        return &ospfv3ProtocolInstances.Ospfv3ProtocolInstance[len(ospfv3ProtocolInstances.Ospfv3ProtocolInstance)-1]
    }
    return nil
}

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfv3ProtocolInstances.Ospfv3ProtocolInstance {
        children[ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i].GetSegmentPath()] = &ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i]
    }
    return children
}

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetYangName() string { return "ospfv3-protocol-instances" }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) SetParent(parent types.Entity) { ospfv3ProtocolInstances.parent = parent }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetParent() types.Entity { return ospfv3ProtocolInstances.parent }

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetParentYangName() string { return "ospf" }

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
// Protocol samples for a particular OSPF v3
// instance
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v3 instance.
    Samples PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
}

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetFilter() yfilter.YFilter { return ospfv3ProtocolInstance.YFilter }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) SetFilter(yf yfilter.YFilter) { ospfv3ProtocolInstance.YFilter = yf }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetSegmentPath() string {
    return "ospfv3-protocol-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfv3ProtocolInstance.InstanceName) + "']"
}

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &ospfv3ProtocolInstance.Samples
    }
    return nil
}

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &ospfv3ProtocolInstance.Samples
    return children
}

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfv3ProtocolInstance.InstanceName
    return leafs
}

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetYangName() string { return "ospfv3-protocol-instance" }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) SetParent(parent types.Entity) { ospfv3ProtocolInstance.parent = parent }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetParent() types.Entity { return ospfv3ProtocolInstance.parent }

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetParentYangName() string { return "ospfv3-protocol-instances" }

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
// Sample Table for an OSPV v3 instance
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample.
    Sample []PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetParentYangName() string { return "ospfv3-protocol-instance" }

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Total number of packets received. The type is interface{} with range:
    // 0..4294967295.
    InputPackets interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputPackets interface{}

    // Number of Hello packets received. The type is interface{} with range:
    // 0..4294967295.
    InputHelloPackets interface{}

    // Number of Hello packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputHelloPackets interface{}

    // Number of DBD packets received. The type is interface{} with range:
    // 0..4294967295.
    InputDbDs interface{}

    // Number of LSA received in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    InputDbDsLsa interface{}

    // Number of DBD packets sent. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDs interface{}

    // Number of LSA sent in DBD packets. The type is interface{} with range:
    // 0..4294967295.
    OutputDbDsLsa interface{}

    // Number of LS Requests received. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequests interface{}

    // Number of LSA received in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    InputLsRequestsLsa interface{}

    // Number of LS Requests sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequests interface{}

    // Number of LSA sent in LS Requests. The type is interface{} with range:
    // 0..4294967295.
    OutputLsRequestsLsa interface{}

    // Number of LSA Updates received. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdates interface{}

    // Number of LSA received in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    InputLsaUpdatesLsa interface{}

    // Number of LSA Updates sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdates interface{}

    // Number of LSA sent in LSA Updates. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaUpdatesLsa interface{}

    // Number of LSA Acknowledgements received. The type is interface{} with
    // range: 0..4294967295.
    InputLsaAcks interface{}

    // Number of LSA received in LSA Acknowledgements. The type is interface{}
    // with range: 0..4294967295.
    InputLsaAcksLsa interface{}

    // Number of LSA Acknowledgements sent. The type is interface{} with range:
    // 0..4294967295.
    OutputLsaAcks interface{}

    // Number of LSA sent in LSA Acknowledgements. The type is interface{} with
    // range: 0..4294967295.
    OutputLsaAcksLsa interface{}
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-packets" { return "InputPackets" }
    if yname == "output-packets" { return "OutputPackets" }
    if yname == "input-hello-packets" { return "InputHelloPackets" }
    if yname == "output-hello-packets" { return "OutputHelloPackets" }
    if yname == "input-db-ds" { return "InputDbDs" }
    if yname == "input-db-ds-lsa" { return "InputDbDsLsa" }
    if yname == "output-db-ds" { return "OutputDbDs" }
    if yname == "output-db-ds-lsa" { return "OutputDbDsLsa" }
    if yname == "input-ls-requests" { return "InputLsRequests" }
    if yname == "input-ls-requests-lsa" { return "InputLsRequestsLsa" }
    if yname == "output-ls-requests" { return "OutputLsRequests" }
    if yname == "output-ls-requests-lsa" { return "OutputLsRequestsLsa" }
    if yname == "input-lsa-updates" { return "InputLsaUpdates" }
    if yname == "input-lsa-updates-lsa" { return "InputLsaUpdatesLsa" }
    if yname == "output-lsa-updates" { return "OutputLsaUpdates" }
    if yname == "output-lsa-updates-lsa" { return "OutputLsaUpdatesLsa" }
    if yname == "input-lsa-acks" { return "InputLsaAcks" }
    if yname == "input-lsa-acks-lsa" { return "InputLsaAcksLsa" }
    if yname == "output-lsa-acks" { return "OutputLsaAcks" }
    if yname == "output-lsa-acks-lsa" { return "OutputLsaAcksLsa" }
    return ""
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-packets"] = sample.InputPackets
    leafs["output-packets"] = sample.OutputPackets
    leafs["input-hello-packets"] = sample.InputHelloPackets
    leafs["output-hello-packets"] = sample.OutputHelloPackets
    leafs["input-db-ds"] = sample.InputDbDs
    leafs["input-db-ds-lsa"] = sample.InputDbDsLsa
    leafs["output-db-ds"] = sample.OutputDbDs
    leafs["output-db-ds-lsa"] = sample.OutputDbDsLsa
    leafs["input-ls-requests"] = sample.InputLsRequests
    leafs["input-ls-requests-lsa"] = sample.InputLsRequestsLsa
    leafs["output-ls-requests"] = sample.OutputLsRequests
    leafs["output-ls-requests-lsa"] = sample.OutputLsRequestsLsa
    leafs["input-lsa-updates"] = sample.InputLsaUpdates
    leafs["input-lsa-updates-lsa"] = sample.InputLsaUpdatesLsa
    leafs["output-lsa-updates"] = sample.OutputLsaUpdates
    leafs["output-lsa-updates-lsa"] = sample.OutputLsaUpdatesLsa
    leafs["input-lsa-acks"] = sample.InputLsaAcks
    leafs["input-lsa-acks-lsa"] = sample.InputLsaAcksLsa
    leafs["output-lsa-acks"] = sample.OutputLsaAcks
    leafs["output-lsa-acks-lsa"] = sample.OutputLsaAcksLsa
    return leafs
}

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Mpls
// Collected MPLS data
type PerfMgmt_Monitor_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP neighbors for which statistics are collected.
    LdpNeighbors PerfMgmt_Monitor_Mpls_LdpNeighbors
}

func (mpls *PerfMgmt_Monitor_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *PerfMgmt_Monitor_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *PerfMgmt_Monitor_Mpls) GetGoName(yname string) string {
    if yname == "ldp-neighbors" { return "LdpNeighbors" }
    return ""
}

func (mpls *PerfMgmt_Monitor_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *PerfMgmt_Monitor_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-neighbors" {
        return &mpls.LdpNeighbors
    }
    return nil
}

func (mpls *PerfMgmt_Monitor_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ldp-neighbors"] = &mpls.LdpNeighbors
    return children
}

func (mpls *PerfMgmt_Monitor_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpls *PerfMgmt_Monitor_Mpls) GetBundleName() string { return "cisco_ios_xr" }

func (mpls *PerfMgmt_Monitor_Mpls) GetYangName() string { return "mpls" }

func (mpls *PerfMgmt_Monitor_Mpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpls *PerfMgmt_Monitor_Mpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpls *PerfMgmt_Monitor_Mpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpls *PerfMgmt_Monitor_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *PerfMgmt_Monitor_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *PerfMgmt_Monitor_Mpls) GetParentYangName() string { return "monitor" }

// PerfMgmt_Monitor_Mpls_LdpNeighbors
// LDP neighbors for which statistics are
// collected
type PerfMgmt_Monitor_Mpls_LdpNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular LDP neighbor. The type is slice of
    // PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor.
    LdpNeighbor []PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor
}

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetFilter() yfilter.YFilter { return ldpNeighbors.YFilter }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) SetFilter(yf yfilter.YFilter) { ldpNeighbors.YFilter = yf }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetGoName(yname string) string {
    if yname == "ldp-neighbor" { return "LdpNeighbor" }
    return ""
}

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetSegmentPath() string {
    return "ldp-neighbors"
}

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-neighbor" {
        for _, c := range ldpNeighbors.LdpNeighbor {
            if ldpNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor{}
        ldpNeighbors.LdpNeighbor = append(ldpNeighbors.LdpNeighbor, child)
        return &ldpNeighbors.LdpNeighbor[len(ldpNeighbors.LdpNeighbor)-1]
    }
    return nil
}

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ldpNeighbors.LdpNeighbor {
        children[ldpNeighbors.LdpNeighbor[i].GetSegmentPath()] = &ldpNeighbors.LdpNeighbor[i]
    }
    return children
}

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetBundleName() string { return "cisco_ios_xr" }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetYangName() string { return "ldp-neighbors" }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) SetParent(parent types.Entity) { ldpNeighbors.parent = parent }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetParent() types.Entity { return ldpNeighbors.parent }

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetParentYangName() string { return "mpls" }

// PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor
// Samples for a particular LDP neighbor
type PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Nbr interface{}

    // Samples for a particular LDP neighbor.
    Samples PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples
}

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetFilter() yfilter.YFilter { return ldpNeighbor.YFilter }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) SetFilter(yf yfilter.YFilter) { ldpNeighbor.YFilter = yf }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetGoName(yname string) string {
    if yname == "nbr" { return "Nbr" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetSegmentPath() string {
    return "ldp-neighbor" + "[nbr='" + fmt.Sprintf("%v", ldpNeighbor.Nbr) + "']"
}

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &ldpNeighbor.Samples
    }
    return nil
}

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &ldpNeighbor.Samples
    return children
}

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nbr"] = ldpNeighbor.Nbr
    return leafs
}

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetYangName() string { return "ldp-neighbor" }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) SetParent(parent types.Entity) { ldpNeighbor.parent = parent }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetParent() types.Entity { return ldpNeighbor.parent }

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetParentYangName() string { return "ldp-neighbors" }

// PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples
// Samples for a particular LDP neighbor
type PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP neighbor statistics sample. The type is slice of
    // PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample.
    Sample []PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetParentYangName() string { return "ldp-neighbor" }

// PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
// LDP neighbor statistics sample
type PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Total messages sent. The type is interface{} with range: 0..65535.
    TotalMsgsSent interface{}

    // Total messages received. The type is interface{} with range: 0..65535.
    TotalMsgsRcvd interface{}

    // Init messages sent. The type is interface{} with range: 0..65535.
    InitMsgsSent interface{}

    // Tnit messages received. The type is interface{} with range: 0..65535.
    InitMsgsRcvd interface{}

    // Address messages sent. The type is interface{} with range: 0..65535.
    AddressMsgsSent interface{}

    // Address messages received. The type is interface{} with range: 0..65535.
    AddressMsgsRcvd interface{}

    // Address withdraw messages sent. The type is interface{} with range:
    // 0..65535.
    AddressWithdrawMsgsSent interface{}

    // Address withdraw messages received. The type is interface{} with range:
    // 0..65535.
    AddressWithdrawMsgsRcvd interface{}

    // Label mapping messages sent. The type is interface{} with range: 0..65535.
    LabelMappingMsgsSent interface{}

    // Label mapping messages received. The type is interface{} with range:
    // 0..65535.
    LabelMappingMsgsRcvd interface{}

    // Label withdraw messages sent. The type is interface{} with range: 0..65535.
    LabelWithdrawMsgsSent interface{}

    // Label withdraw messages received. The type is interface{} with range:
    // 0..65535.
    LabelWithdrawMsgsRcvd interface{}

    // Label release messages sent. The type is interface{} with range: 0..65535.
    LabelReleaseMsgsSent interface{}

    // Label release messages received. The type is interface{} with range:
    // 0..65535.
    LabelReleaseMsgsRcvd interface{}

    // Notification messages sent. The type is interface{} with range: 0..65535.
    NotificationMsgsSent interface{}

    // Notification messages received. The type is interface{} with range:
    // 0..65535.
    NotificationMsgsRcvd interface{}

    // Keepalive messages sent. The type is interface{} with range: 0..65535.
    KeepaliveMsgsSent interface{}

    // Keepalive messages received. The type is interface{} with range: 0..65535.
    KeepaliveMsgsRcvd interface{}
}

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "total-msgs-sent" { return "TotalMsgsSent" }
    if yname == "total-msgs-rcvd" { return "TotalMsgsRcvd" }
    if yname == "init-msgs-sent" { return "InitMsgsSent" }
    if yname == "init-msgs-rcvd" { return "InitMsgsRcvd" }
    if yname == "address-msgs-sent" { return "AddressMsgsSent" }
    if yname == "address-msgs-rcvd" { return "AddressMsgsRcvd" }
    if yname == "address-withdraw-msgs-sent" { return "AddressWithdrawMsgsSent" }
    if yname == "address-withdraw-msgs-rcvd" { return "AddressWithdrawMsgsRcvd" }
    if yname == "label-mapping-msgs-sent" { return "LabelMappingMsgsSent" }
    if yname == "label-mapping-msgs-rcvd" { return "LabelMappingMsgsRcvd" }
    if yname == "label-withdraw-msgs-sent" { return "LabelWithdrawMsgsSent" }
    if yname == "label-withdraw-msgs-rcvd" { return "LabelWithdrawMsgsRcvd" }
    if yname == "label-release-msgs-sent" { return "LabelReleaseMsgsSent" }
    if yname == "label-release-msgs-rcvd" { return "LabelReleaseMsgsRcvd" }
    if yname == "notification-msgs-sent" { return "NotificationMsgsSent" }
    if yname == "notification-msgs-rcvd" { return "NotificationMsgsRcvd" }
    if yname == "keepalive-msgs-sent" { return "KeepaliveMsgsSent" }
    if yname == "keepalive-msgs-rcvd" { return "KeepaliveMsgsRcvd" }
    return ""
}

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["total-msgs-sent"] = sample.TotalMsgsSent
    leafs["total-msgs-rcvd"] = sample.TotalMsgsRcvd
    leafs["init-msgs-sent"] = sample.InitMsgsSent
    leafs["init-msgs-rcvd"] = sample.InitMsgsRcvd
    leafs["address-msgs-sent"] = sample.AddressMsgsSent
    leafs["address-msgs-rcvd"] = sample.AddressMsgsRcvd
    leafs["address-withdraw-msgs-sent"] = sample.AddressWithdrawMsgsSent
    leafs["address-withdraw-msgs-rcvd"] = sample.AddressWithdrawMsgsRcvd
    leafs["label-mapping-msgs-sent"] = sample.LabelMappingMsgsSent
    leafs["label-mapping-msgs-rcvd"] = sample.LabelMappingMsgsRcvd
    leafs["label-withdraw-msgs-sent"] = sample.LabelWithdrawMsgsSent
    leafs["label-withdraw-msgs-rcvd"] = sample.LabelWithdrawMsgsRcvd
    leafs["label-release-msgs-sent"] = sample.LabelReleaseMsgsSent
    leafs["label-release-msgs-rcvd"] = sample.LabelReleaseMsgsRcvd
    leafs["notification-msgs-sent"] = sample.NotificationMsgsSent
    leafs["notification-msgs-rcvd"] = sample.NotificationMsgsRcvd
    leafs["keepalive-msgs-sent"] = sample.KeepaliveMsgsSent
    leafs["keepalive-msgs-rcvd"] = sample.KeepaliveMsgsRcvd
    return leafs
}

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Nodes
// Nodes for which data is collected
type PerfMgmt_Monitor_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Instance. The type is slice of PerfMgmt_Monitor_Nodes_Node.
    Node []PerfMgmt_Monitor_Nodes_Node
}

func (nodes *PerfMgmt_Monitor_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Monitor_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Monitor_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Monitor_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Monitor_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Monitor_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Monitor_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Monitor_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Monitor_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Monitor_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Monitor_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Monitor_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Monitor_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Monitor_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Monitor_Nodes) GetParentYangName() string { return "monitor" }

// PerfMgmt_Monitor_Nodes_Node
// Node Instance
type PerfMgmt_Monitor_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Node CPU data.
    SampleXr PerfMgmt_Monitor_Nodes_Node_SampleXr

    // Processes data.
    Processes PerfMgmt_Monitor_Nodes_Node_Processes

    // Node Memory data.
    Samples PerfMgmt_Monitor_Nodes_Node_Samples
}

func (node *PerfMgmt_Monitor_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Monitor_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Monitor_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "sample-xr" { return "SampleXr" }
    if yname == "processes" { return "Processes" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (node *PerfMgmt_Monitor_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Monitor_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample-xr" {
        return &node.SampleXr
    }
    if childYangName == "processes" {
        return &node.Processes
    }
    if childYangName == "samples" {
        return &node.Samples
    }
    return nil
}

func (node *PerfMgmt_Monitor_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sample-xr"] = &node.SampleXr
    children["processes"] = &node.Processes
    children["samples"] = &node.Samples
    return children
}

func (node *PerfMgmt_Monitor_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    return leafs
}

func (node *PerfMgmt_Monitor_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Monitor_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Monitor_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Monitor_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Monitor_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Monitor_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Monitor_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Monitor_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Monitor_Nodes_Node_SampleXr
// Node CPU data
type PerfMgmt_Monitor_Nodes_Node_SampleXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node CPU data sample. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample.
    Sample []PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample
}

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetFilter() yfilter.YFilter { return sampleXr.YFilter }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) SetFilter(yf yfilter.YFilter) { sampleXr.YFilter = yf }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetSegmentPath() string {
    return "sample-xr"
}

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range sampleXr.Sample {
            if sampleXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample{}
        sampleXr.Sample = append(sampleXr.Sample, child)
        return &sampleXr.Sample[len(sampleXr.Sample)-1]
    }
    return nil
}

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sampleXr.Sample {
        children[sampleXr.Sample[i].GetSegmentPath()] = &sampleXr.Sample[i]
    }
    return children
}

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetBundleName() string { return "cisco_ios_xr" }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetYangName() string { return "sample-xr" }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) SetParent(parent types.Entity) { sampleXr.parent = parent }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetParent() types.Entity { return sampleXr.parent }

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetParentYangName() string { return "node" }

// PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample
// Node CPU data sample
type PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Number of processes in the system. The type is interface{} with range:
    // 0..4294967295.
    NoProcesses interface{}

    // Average system %CPU utilization. The type is interface{} with range:
    // 0..4294967295.
    AverageCpuUsed interface{}
}

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "no-processes" { return "NoProcesses" }
    if yname == "average-cpu-used" { return "AverageCpuUsed" }
    return ""
}

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["no-processes"] = sample.NoProcesses
    leafs["average-cpu-used"] = sample.AverageCpuUsed
    return leafs
}

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetParentYangName() string { return "sample-xr" }

// PerfMgmt_Monitor_Nodes_Node_Processes
// Processes data
type PerfMgmt_Monitor_Nodes_Node_Processes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process data. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_Processes_Process.
    Process []PerfMgmt_Monitor_Nodes_Node_Processes_Process
}

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetFilter() yfilter.YFilter { return processes.YFilter }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) SetFilter(yf yfilter.YFilter) { processes.YFilter = yf }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    return ""
}

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetSegmentPath() string {
    return "processes"
}

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process" {
        for _, c := range processes.Process {
            if processes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Nodes_Node_Processes_Process{}
        processes.Process = append(processes.Process, child)
        return &processes.Process[len(processes.Process)-1]
    }
    return nil
}

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range processes.Process {
        children[processes.Process[i].GetSegmentPath()] = &processes.Process[i]
    }
    return children
}

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetBundleName() string { return "cisco_ios_xr" }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetYangName() string { return "processes" }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) SetParent(parent types.Entity) { processes.parent = parent }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetParent() types.Entity { return processes.parent }

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetParentYangName() string { return "node" }

// PerfMgmt_Monitor_Nodes_Node_Processes_Process
// Process data
type PerfMgmt_Monitor_Nodes_Node_Processes_Process struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Process data.
    Samples PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples
}

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetFilter() yfilter.YFilter { return process.YFilter }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) SetFilter(yf yfilter.YFilter) { process.YFilter = yf }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetGoName(yname string) string {
    if yname == "process-id" { return "ProcessId" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetSegmentPath() string {
    return "process" + "[process-id='" + fmt.Sprintf("%v", process.ProcessId) + "']"
}

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &process.Samples
    }
    return nil
}

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &process.Samples
    return children
}

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-id"] = process.ProcessId
    return leafs
}

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetBundleName() string { return "cisco_ios_xr" }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetYangName() string { return "process" }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) SetParent(parent types.Entity) { process.parent = parent }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetParent() types.Entity { return process.parent }

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetParentYangName() string { return "processes" }

// PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples
// Process data
type PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process data sample. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample.
    Sample []PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetParentYangName() string { return "process" }

// PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample
// Process data sample
type PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Max. dynamic memory (KBytes) used since startup time. The type is
    // interface{} with range: 0..4294967295. Units are kilobyte.
    PeakMemory interface{}

    // Average %CPU utilization. The type is interface{} with range:
    // 0..4294967295.
    AverageCpuUsed interface{}

    // Number of threads. The type is interface{} with range: 0..4294967295.
    NoThreads interface{}
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "peak-memory" { return "PeakMemory" }
    if yname == "average-cpu-used" { return "AverageCpuUsed" }
    if yname == "no-threads" { return "NoThreads" }
    return ""
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["peak-memory"] = sample.PeakMemory
    leafs["average-cpu-used"] = sample.AverageCpuUsed
    leafs["no-threads"] = sample.NoThreads
    return leafs
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Nodes_Node_Samples
// Node Memory data
type PerfMgmt_Monitor_Nodes_Node_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Memory data sample. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_Samples_Sample.
    Sample []PerfMgmt_Monitor_Nodes_Node_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Nodes_Node_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetParentYangName() string { return "node" }

// PerfMgmt_Monitor_Nodes_Node_Samples_Sample
// Node Memory data sample
type PerfMgmt_Monitor_Nodes_Node_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Current application memory (Bytes) in use. The type is interface{} with
    // range: 0..4294967295. Units are byte.
    CurrMemory interface{}

    // Max. system memory (MBytes) used since bootup. The type is interface{} with
    // range: 0..4294967295. Units are megabyte.
    PeakMemory interface{}
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "curr-memory" { return "CurrMemory" }
    if yname == "peak-memory" { return "PeakMemory" }
    return ""
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["curr-memory"] = sample.CurrMemory
    leafs["peak-memory"] = sample.PeakMemory
    return leafs
}

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Bgp
// Collected BGP data
type PerfMgmt_Monitor_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbors for which statistics are collected.
    BgpNeighbors PerfMgmt_Monitor_Bgp_BgpNeighbors
}

func (bgp *PerfMgmt_Monitor_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PerfMgmt_Monitor_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PerfMgmt_Monitor_Bgp) GetGoName(yname string) string {
    if yname == "bgp-neighbors" { return "BgpNeighbors" }
    return ""
}

func (bgp *PerfMgmt_Monitor_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PerfMgmt_Monitor_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-neighbors" {
        return &bgp.BgpNeighbors
    }
    return nil
}

func (bgp *PerfMgmt_Monitor_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-neighbors"] = &bgp.BgpNeighbors
    return children
}

func (bgp *PerfMgmt_Monitor_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *PerfMgmt_Monitor_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PerfMgmt_Monitor_Bgp) GetYangName() string { return "bgp" }

func (bgp *PerfMgmt_Monitor_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PerfMgmt_Monitor_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PerfMgmt_Monitor_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PerfMgmt_Monitor_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PerfMgmt_Monitor_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PerfMgmt_Monitor_Bgp) GetParentYangName() string { return "monitor" }

// PerfMgmt_Monitor_Bgp_BgpNeighbors
// Neighbors for which statistics are collected
type PerfMgmt_Monitor_Bgp_BgpNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for particular neighbor. The type is slice of
    // PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor.
    BgpNeighbor []PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor
}

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetFilter() yfilter.YFilter { return bgpNeighbors.YFilter }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) SetFilter(yf yfilter.YFilter) { bgpNeighbors.YFilter = yf }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetGoName(yname string) string {
    if yname == "bgp-neighbor" { return "BgpNeighbor" }
    return ""
}

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetSegmentPath() string {
    return "bgp-neighbors"
}

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-neighbor" {
        for _, c := range bgpNeighbors.BgpNeighbor {
            if bgpNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor{}
        bgpNeighbors.BgpNeighbor = append(bgpNeighbors.BgpNeighbor, child)
        return &bgpNeighbors.BgpNeighbor[len(bgpNeighbors.BgpNeighbor)-1]
    }
    return nil
}

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpNeighbors.BgpNeighbor {
        children[bgpNeighbors.BgpNeighbor[i].GetSegmentPath()] = &bgpNeighbors.BgpNeighbor[i]
    }
    return children
}

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetBundleName() string { return "cisco_ios_xr" }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetYangName() string { return "bgp-neighbors" }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) SetParent(parent types.Entity) { bgpNeighbors.parent = parent }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetParent() types.Entity { return bgpNeighbors.parent }

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetParentYangName() string { return "bgp" }

// PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor
// Samples for particular neighbor
type PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. BGP Neighbor Identifier. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Sample Table for a BGP neighbor.
    Samples PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples
}

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetFilter() yfilter.YFilter { return bgpNeighbor.YFilter }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) SetFilter(yf yfilter.YFilter) { bgpNeighbor.YFilter = yf }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetSegmentPath() string {
    return "bgp-neighbor" + "[ip-address='" + fmt.Sprintf("%v", bgpNeighbor.IpAddress) + "']"
}

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &bgpNeighbor.Samples
    }
    return nil
}

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &bgpNeighbor.Samples
    return children
}

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = bgpNeighbor.IpAddress
    return leafs
}

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetYangName() string { return "bgp-neighbor" }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) SetParent(parent types.Entity) { bgpNeighbor.parent = parent }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetParent() types.Entity { return bgpNeighbor.parent }

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetParentYangName() string { return "bgp-neighbors" }

// PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples
// Sample Table for a BGP neighbor
type PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor statistics sample. The type is slice of
    // PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample.
    Sample []PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetParentYangName() string { return "bgp-neighbor" }

// PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
// Neighbor statistics sample
type PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Number of messages received. The type is interface{} with range:
    // 0..4294967295.
    InputMessages interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    OutputMessages interface{}

    // Number of update messages received. The type is interface{} with range:
    // 0..4294967295.
    InputUpdateMessages interface{}

    // Number of update messages sent. The type is interface{} with range:
    // 0..4294967295.
    OutputUpdateMessages interface{}

    // Number of times the connection was established. The type is interface{}
    // with range: 0..4294967295.
    ConnEstablished interface{}

    // Number of times connection was dropped. The type is interface{} with range:
    // 0..4294967295.
    ConnDropped interface{}

    // Number of error notifications received on the connection. The type is
    // interface{} with range: 0..4294967295.
    ErrorsReceived interface{}

    // Number of error notifications sent on the connection. The type is
    // interface{} with range: 0..4294967295.
    ErrorsSent interface{}
}

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-messages" { return "InputMessages" }
    if yname == "output-messages" { return "OutputMessages" }
    if yname == "input-update-messages" { return "InputUpdateMessages" }
    if yname == "output-update-messages" { return "OutputUpdateMessages" }
    if yname == "conn-established" { return "ConnEstablished" }
    if yname == "conn-dropped" { return "ConnDropped" }
    if yname == "errors-received" { return "ErrorsReceived" }
    if yname == "errors-sent" { return "ErrorsSent" }
    return ""
}

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-messages"] = sample.InputMessages
    leafs["output-messages"] = sample.OutputMessages
    leafs["input-update-messages"] = sample.InputUpdateMessages
    leafs["output-update-messages"] = sample.OutputUpdateMessages
    leafs["conn-established"] = sample.ConnEstablished
    leafs["conn-dropped"] = sample.ConnDropped
    leafs["errors-received"] = sample.ErrorsReceived
    leafs["errors-sent"] = sample.ErrorsSent
    return leafs
}

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Interface
// Collected Interface data
type PerfMgmt_Monitor_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interfaces for which Generic Counters are collected.
    GenericCounterInterfaces PerfMgmt_Monitor_Interface_GenericCounterInterfaces

    // Interfaces for which Basic Counters are collected.
    BasicCounterInterfaces PerfMgmt_Monitor_Interface_BasicCounterInterfaces

    // Interfaces for which Data Rates are collected.
    DataRateInterfaces PerfMgmt_Monitor_Interface_DataRateInterfaces
}

func (self *PerfMgmt_Monitor_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PerfMgmt_Monitor_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PerfMgmt_Monitor_Interface) GetGoName(yname string) string {
    if yname == "generic-counter-interfaces" { return "GenericCounterInterfaces" }
    if yname == "basic-counter-interfaces" { return "BasicCounterInterfaces" }
    if yname == "data-rate-interfaces" { return "DataRateInterfaces" }
    return ""
}

func (self *PerfMgmt_Monitor_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *PerfMgmt_Monitor_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interfaces" {
        return &self.GenericCounterInterfaces
    }
    if childYangName == "basic-counter-interfaces" {
        return &self.BasicCounterInterfaces
    }
    if childYangName == "data-rate-interfaces" {
        return &self.DataRateInterfaces
    }
    return nil
}

func (self *PerfMgmt_Monitor_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["generic-counter-interfaces"] = &self.GenericCounterInterfaces
    children["basic-counter-interfaces"] = &self.BasicCounterInterfaces
    children["data-rate-interfaces"] = &self.DataRateInterfaces
    return children
}

func (self *PerfMgmt_Monitor_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *PerfMgmt_Monitor_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *PerfMgmt_Monitor_Interface) GetYangName() string { return "interface" }

func (self *PerfMgmt_Monitor_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PerfMgmt_Monitor_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PerfMgmt_Monitor_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PerfMgmt_Monitor_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *PerfMgmt_Monitor_Interface) GetParent() types.Entity { return self.parent }

func (self *PerfMgmt_Monitor_Interface) GetParentYangName() string { return "monitor" }

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces
// Interfaces for which Generic Counters are
// collected
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface.
    GenericCounterInterface []PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface
}

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetFilter() yfilter.YFilter { return genericCounterInterfaces.YFilter }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) SetFilter(yf yfilter.YFilter) { genericCounterInterfaces.YFilter = yf }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetGoName(yname string) string {
    if yname == "generic-counter-interface" { return "GenericCounterInterface" }
    return ""
}

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetSegmentPath() string {
    return "generic-counter-interfaces"
}

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interface" {
        for _, c := range genericCounterInterfaces.GenericCounterInterface {
            if genericCounterInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface{}
        genericCounterInterfaces.GenericCounterInterface = append(genericCounterInterfaces.GenericCounterInterface, child)
        return &genericCounterInterfaces.GenericCounterInterface[len(genericCounterInterfaces.GenericCounterInterface)-1]
    }
    return nil
}

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range genericCounterInterfaces.GenericCounterInterface {
        children[genericCounterInterfaces.GenericCounterInterface[i].GetSegmentPath()] = &genericCounterInterfaces.GenericCounterInterface[i]
    }
    return children
}

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetYangName() string { return "generic-counter-interfaces" }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) SetParent(parent types.Entity) { genericCounterInterfaces.parent = parent }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetParent() types.Entity { return genericCounterInterfaces.parent }

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetParentYangName() string { return "interface" }

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface
// Samples for a particular interface
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Generic Counter samples for an interface.
    Samples PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
}

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetFilter() yfilter.YFilter { return genericCounterInterface.YFilter }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) SetFilter(yf yfilter.YFilter) { genericCounterInterface.YFilter = yf }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetSegmentPath() string {
    return "generic-counter-interface" + "[interface-name='" + fmt.Sprintf("%v", genericCounterInterface.InterfaceName) + "']"
}

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &genericCounterInterface.Samples
    }
    return nil
}

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &genericCounterInterface.Samples
    return children
}

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = genericCounterInterface.InterfaceName
    return leafs
}

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetYangName() string { return "generic-counter-interface" }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) SetParent(parent types.Entity) { genericCounterInterface.parent = parent }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetParent() types.Entity { return genericCounterInterface.parent }

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetParentYangName() string { return "generic-counter-interfaces" }

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
// Generic Counter samples for an interface
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample.
    Sample []PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetParentYangName() string { return "generic-counter-interface" }

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
// Generic Counters sample
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InPackets interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    InOctets interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    OutPackets interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    OutOctets interface{}

    // Unicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPkts interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InMulticastPkts interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InBroadcastPkts interface{}

    // Unicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPkts interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMulticastPkts interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBroadcastPkts interface{}

    // Outbound correct packets discarded. The type is interface{} with range:
    // 0..4294967295.
    OutputTotalDrops interface{}

    // Inbound correct packets discarded. The type is interface{} with range:
    // 0..4294967295.
    InputTotalDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Inbound packets discarded with unknown proto. The type is interface{} with
    // range: 0..4294967295.
    InputUnknownProto interface{}

    // Outbound incorrect packets discarded. The type is interface{} with range:
    // 0..4294967295.
    OutputTotalErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderrun interface{}

    // Inbound incorrect packets discarded. The type is interface{} with range:
    // 0..4294967295.
    InputTotalErrors interface{}

    // Inbound packets discarded with incorrect CRC. The type is interface{} with
    // range: 0..4294967295.
    InputCrc interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverrun interface{}

    // Inbound framing errors. The type is interface{} with range: 0..4294967295.
    InputFrame interface{}
}

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "in-packets" { return "InPackets" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "out-packets" { return "OutPackets" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "in-ucast-pkts" { return "InUcastPkts" }
    if yname == "in-multicast-pkts" { return "InMulticastPkts" }
    if yname == "in-broadcast-pkts" { return "InBroadcastPkts" }
    if yname == "out-ucast-pkts" { return "OutUcastPkts" }
    if yname == "out-multicast-pkts" { return "OutMulticastPkts" }
    if yname == "out-broadcast-pkts" { return "OutBroadcastPkts" }
    if yname == "output-total-drops" { return "OutputTotalDrops" }
    if yname == "input-total-drops" { return "InputTotalDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "input-unknown-proto" { return "InputUnknownProto" }
    if yname == "output-total-errors" { return "OutputTotalErrors" }
    if yname == "output-underrun" { return "OutputUnderrun" }
    if yname == "input-total-errors" { return "InputTotalErrors" }
    if yname == "input-crc" { return "InputCrc" }
    if yname == "input-overrun" { return "InputOverrun" }
    if yname == "input-frame" { return "InputFrame" }
    return ""
}

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["in-packets"] = sample.InPackets
    leafs["in-octets"] = sample.InOctets
    leafs["out-packets"] = sample.OutPackets
    leafs["out-octets"] = sample.OutOctets
    leafs["in-ucast-pkts"] = sample.InUcastPkts
    leafs["in-multicast-pkts"] = sample.InMulticastPkts
    leafs["in-broadcast-pkts"] = sample.InBroadcastPkts
    leafs["out-ucast-pkts"] = sample.OutUcastPkts
    leafs["out-multicast-pkts"] = sample.OutMulticastPkts
    leafs["out-broadcast-pkts"] = sample.OutBroadcastPkts
    leafs["output-total-drops"] = sample.OutputTotalDrops
    leafs["input-total-drops"] = sample.InputTotalDrops
    leafs["input-queue-drops"] = sample.InputQueueDrops
    leafs["input-unknown-proto"] = sample.InputUnknownProto
    leafs["output-total-errors"] = sample.OutputTotalErrors
    leafs["output-underrun"] = sample.OutputUnderrun
    leafs["input-total-errors"] = sample.InputTotalErrors
    leafs["input-crc"] = sample.InputCrc
    leafs["input-overrun"] = sample.InputOverrun
    leafs["input-frame"] = sample.InputFrame
    return leafs
}

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces
// Interfaces for which Basic Counters are
// collected
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface.
    BasicCounterInterface []PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface
}

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetFilter() yfilter.YFilter { return basicCounterInterfaces.YFilter }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) SetFilter(yf yfilter.YFilter) { basicCounterInterfaces.YFilter = yf }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetGoName(yname string) string {
    if yname == "basic-counter-interface" { return "BasicCounterInterface" }
    return ""
}

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetSegmentPath() string {
    return "basic-counter-interfaces"
}

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-counter-interface" {
        for _, c := range basicCounterInterfaces.BasicCounterInterface {
            if basicCounterInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface{}
        basicCounterInterfaces.BasicCounterInterface = append(basicCounterInterfaces.BasicCounterInterface, child)
        return &basicCounterInterfaces.BasicCounterInterface[len(basicCounterInterfaces.BasicCounterInterface)-1]
    }
    return nil
}

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range basicCounterInterfaces.BasicCounterInterface {
        children[basicCounterInterfaces.BasicCounterInterface[i].GetSegmentPath()] = &basicCounterInterfaces.BasicCounterInterface[i]
    }
    return children
}

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetYangName() string { return "basic-counter-interfaces" }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) SetParent(parent types.Entity) { basicCounterInterfaces.parent = parent }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetParent() types.Entity { return basicCounterInterfaces.parent }

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetParentYangName() string { return "interface" }

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface
// Samples for a particular interface
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Basic Counter samples for an interface.
    Samples PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
}

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetFilter() yfilter.YFilter { return basicCounterInterface.YFilter }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) SetFilter(yf yfilter.YFilter) { basicCounterInterface.YFilter = yf }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetSegmentPath() string {
    return "basic-counter-interface" + "[interface-name='" + fmt.Sprintf("%v", basicCounterInterface.InterfaceName) + "']"
}

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &basicCounterInterface.Samples
    }
    return nil
}

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &basicCounterInterface.Samples
    return children
}

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = basicCounterInterface.InterfaceName
    return leafs
}

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetYangName() string { return "basic-counter-interface" }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) SetParent(parent types.Entity) { basicCounterInterface.parent = parent }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetParent() types.Entity { return basicCounterInterface.parent }

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetParentYangName() string { return "basic-counter-interfaces" }

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
// Basic Counter samples for an interface
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Basic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample.
    Sample []PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetParentYangName() string { return "basic-counter-interface" }

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
// Basic Counters sample
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds from UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InPackets interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    InOctets interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    OutPackets interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    OutOctets interface{}

    // Inbound correct packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    InputTotalDrops interface{}

    // Input queue drops. The type is interface{} with range:
    // 0..18446744073709551615.
    InputQueueDrops interface{}

    // Inbound incorrect packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    InputTotalErrors interface{}

    // Outbound correct packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputTotalDrops interface{}

    // Output queue drops. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputQueueDrops interface{}

    // Outbound incorrect packets discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputTotalErrors interface{}
}

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "in-packets" { return "InPackets" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "out-packets" { return "OutPackets" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "input-total-drops" { return "InputTotalDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "input-total-errors" { return "InputTotalErrors" }
    if yname == "output-total-drops" { return "OutputTotalDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "output-total-errors" { return "OutputTotalErrors" }
    return ""
}

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["in-packets"] = sample.InPackets
    leafs["in-octets"] = sample.InOctets
    leafs["out-packets"] = sample.OutPackets
    leafs["out-octets"] = sample.OutOctets
    leafs["input-total-drops"] = sample.InputTotalDrops
    leafs["input-queue-drops"] = sample.InputQueueDrops
    leafs["input-total-errors"] = sample.InputTotalErrors
    leafs["output-total-drops"] = sample.OutputTotalDrops
    leafs["output-queue-drops"] = sample.OutputQueueDrops
    leafs["output-total-errors"] = sample.OutputTotalErrors
    return leafs
}

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetParentYangName() string { return "samples" }

// PerfMgmt_Monitor_Interface_DataRateInterfaces
// Interfaces for which Data Rates are collected
type PerfMgmt_Monitor_Interface_DataRateInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface.
    DataRateInterface []PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface
}

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetFilter() yfilter.YFilter { return dataRateInterfaces.YFilter }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) SetFilter(yf yfilter.YFilter) { dataRateInterfaces.YFilter = yf }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetGoName(yname string) string {
    if yname == "data-rate-interface" { return "DataRateInterface" }
    return ""
}

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetSegmentPath() string {
    return "data-rate-interfaces"
}

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-rate-interface" {
        for _, c := range dataRateInterfaces.DataRateInterface {
            if dataRateInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface{}
        dataRateInterfaces.DataRateInterface = append(dataRateInterfaces.DataRateInterface, child)
        return &dataRateInterfaces.DataRateInterface[len(dataRateInterfaces.DataRateInterface)-1]
    }
    return nil
}

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dataRateInterfaces.DataRateInterface {
        children[dataRateInterfaces.DataRateInterface[i].GetSegmentPath()] = &dataRateInterfaces.DataRateInterface[i]
    }
    return children
}

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetYangName() string { return "data-rate-interfaces" }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) SetParent(parent types.Entity) { dataRateInterfaces.parent = parent }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetParent() types.Entity { return dataRateInterfaces.parent }

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetParentYangName() string { return "interface" }

// PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface
// Samples for a particular interface
type PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Data Rate samples for an interface.
    Samples PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples
}

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetFilter() yfilter.YFilter { return dataRateInterface.YFilter }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) SetFilter(yf yfilter.YFilter) { dataRateInterface.YFilter = yf }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "samples" { return "Samples" }
    return ""
}

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetSegmentPath() string {
    return "data-rate-interface" + "[interface-name='" + fmt.Sprintf("%v", dataRateInterface.InterfaceName) + "']"
}

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "samples" {
        return &dataRateInterface.Samples
    }
    return nil
}

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["samples"] = &dataRateInterface.Samples
    return children
}

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = dataRateInterface.InterfaceName
    return leafs
}

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetYangName() string { return "data-rate-interface" }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) SetParent(parent types.Entity) { dataRateInterface.parent = parent }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetParent() types.Entity { return dataRateInterface.parent }

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetParentYangName() string { return "data-rate-interfaces" }

// PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples
// Data Rate samples for an interface
type PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data Rates sample. The type is slice of
    // PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample.
    Sample []PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetFilter() yfilter.YFilter { return samples.YFilter }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) SetFilter(yf yfilter.YFilter) { samples.YFilter = yf }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetSegmentPath() string {
    return "samples"
}

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range samples.Sample {
            if samples.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample{}
        samples.Sample = append(samples.Sample, child)
        return &samples.Sample[len(samples.Sample)-1]
    }
    return nil
}

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samples.Sample {
        children[samples.Sample[i].GetSegmentPath()] = &samples.Sample[i]
    }
    return children
}

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetBundleName() string { return "cisco_ios_xr" }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetYangName() string { return "samples" }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) SetParent(parent types.Entity) { samples.parent = parent }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetParent() types.Entity { return samples.parent }

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetParentYangName() string { return "data-rate-interface" }

// PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
// Data Rates sample
type PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sample ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleId interface{}

    // Timestamp of sample in seconds drom UCT. The type is interface{} with
    // range: 0..18446744073709551615. Units are second.
    TimeStamp interface{}

    // Input datarate in 1000's of bps. The type is interface{} with range:
    // 0..4294967295. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..4294967295. Units are packet/s.
    InputPacketRate interface{}

    // Output datarate in 1000's of bps. The type is interface{} with range:
    // 0..4294967295. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..4294967295. Units are packet/s.
    OutputPacketRate interface{}

    // Peak input datarate. The type is interface{} with range: 0..4294967295.
    InputPeakRate interface{}

    // Peak input packet rate. The type is interface{} with range: 0..4294967295.
    InputPeakPkts interface{}

    // Peak output datarate. The type is interface{} with range: 0..4294967295.
    OutputPeakRate interface{}

    // Peak output packet rate. The type is interface{} with range: 0..4294967295.
    OutputPeakPkts interface{}

    // Bandwidth (in kbps). The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    Bandwidth interface{}
}

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetGoName(yname string) string {
    if yname == "sample-id" { return "SampleId" }
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "input-peak-rate" { return "InputPeakRate" }
    if yname == "input-peak-pkts" { return "InputPeakPkts" }
    if yname == "output-peak-rate" { return "OutputPeakRate" }
    if yname == "output-peak-pkts" { return "OutputPeakPkts" }
    if yname == "bandwidth" { return "Bandwidth" }
    return ""
}

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetSegmentPath() string {
    return "sample" + "[sample-id='" + fmt.Sprintf("%v", sample.SampleId) + "']"
}

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-id"] = sample.SampleId
    leafs["time-stamp"] = sample.TimeStamp
    leafs["input-data-rate"] = sample.InputDataRate
    leafs["input-packet-rate"] = sample.InputPacketRate
    leafs["output-data-rate"] = sample.OutputDataRate
    leafs["output-packet-rate"] = sample.OutputPacketRate
    leafs["input-peak-rate"] = sample.InputPeakRate
    leafs["input-peak-pkts"] = sample.InputPeakPkts
    leafs["output-peak-rate"] = sample.OutputPeakRate
    leafs["output-peak-pkts"] = sample.OutputPeakPkts
    leafs["bandwidth"] = sample.Bandwidth
    return leafs
}

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetYangName() string { return "sample" }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetParent() types.Entity { return sample.parent }

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetParentYangName() string { return "samples" }

