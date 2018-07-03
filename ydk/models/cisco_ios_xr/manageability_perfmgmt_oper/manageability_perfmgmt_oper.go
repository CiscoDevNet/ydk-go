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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data from periodic requests.
    Periodic PerfMgmt_Periodic

    // Data from monitor (one history period) requests.
    Monitor PerfMgmt_Monitor
}

func (perfMgmt *PerfMgmt) GetEntityData() *types.CommonEntityData {
    perfMgmt.EntityData.YFilter = perfMgmt.YFilter
    perfMgmt.EntityData.YangName = "perf-mgmt"
    perfMgmt.EntityData.BundleName = "cisco_ios_xr"
    perfMgmt.EntityData.ParentYangName = "Cisco-IOS-XR-manageability-perfmgmt-oper"
    perfMgmt.EntityData.SegmentPath = "Cisco-IOS-XR-manageability-perfmgmt-oper:perf-mgmt"
    perfMgmt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perfMgmt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perfMgmt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perfMgmt.EntityData.Children = types.NewOrderedMap()
    perfMgmt.EntityData.Children.Append("periodic", types.YChild{"Periodic", &perfMgmt.Periodic})
    perfMgmt.EntityData.Children.Append("monitor", types.YChild{"Monitor", &perfMgmt.Monitor})
    perfMgmt.EntityData.Leafs = types.NewOrderedMap()

    perfMgmt.EntityData.YListKeys = []string {}

    return &(perfMgmt.EntityData)
}

// PerfMgmt_Periodic
// Data from periodic requests
type PerfMgmt_Periodic struct {
    EntityData types.CommonEntityData
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

func (periodic *PerfMgmt_Periodic) GetEntityData() *types.CommonEntityData {
    periodic.EntityData.YFilter = periodic.YFilter
    periodic.EntityData.YangName = "periodic"
    periodic.EntityData.BundleName = "cisco_ios_xr"
    periodic.EntityData.ParentYangName = "perf-mgmt"
    periodic.EntityData.SegmentPath = "periodic"
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = types.NewOrderedMap()
    periodic.EntityData.Children.Append("ospf", types.YChild{"Ospf", &periodic.Ospf})
    periodic.EntityData.Children.Append("mpls", types.YChild{"Mpls", &periodic.Mpls})
    periodic.EntityData.Children.Append("nodes", types.YChild{"Nodes", &periodic.Nodes})
    periodic.EntityData.Children.Append("bgp", types.YChild{"Bgp", &periodic.Bgp})
    periodic.EntityData.Children.Append("interface", types.YChild{"Interface", &periodic.Interface})
    periodic.EntityData.Leafs = types.NewOrderedMap()

    periodic.EntityData.YListKeys = []string {}

    return &(periodic.EntityData)
}

// PerfMgmt_Periodic_Ospf
// Collected OSPF data
type PerfMgmt_Periodic_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 instances for which protocol statistics are collected.
    Ospfv2ProtocolInstances PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances

    // OSPF v3 instances for which protocol statistics are collected.
    Ospfv3ProtocolInstances PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances
}

func (ospf *PerfMgmt_Periodic_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "periodic"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Children.Append("ospfv2-protocol-instances", types.YChild{"Ospfv2ProtocolInstances", &ospf.Ospfv2ProtocolInstances})
    ospf.EntityData.Children.Append("ospfv3-protocol-instances", types.YChild{"Ospfv3ProtocolInstances", &ospf.Ospfv3ProtocolInstances})
    ospf.EntityData.Leafs = types.NewOrderedMap()

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances
// OSPF v2 instances for which protocol statistics
// are collected
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v2 instance. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance.
    Ospfv2ProtocolInstance []*PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
}

func (ospfv2ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances) GetEntityData() *types.CommonEntityData {
    ospfv2ProtocolInstances.EntityData.YFilter = ospfv2ProtocolInstances.YFilter
    ospfv2ProtocolInstances.EntityData.YangName = "ospfv2-protocol-instances"
    ospfv2ProtocolInstances.EntityData.BundleName = "cisco_ios_xr"
    ospfv2ProtocolInstances.EntityData.ParentYangName = "ospf"
    ospfv2ProtocolInstances.EntityData.SegmentPath = "ospfv2-protocol-instances"
    ospfv2ProtocolInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2ProtocolInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2ProtocolInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2ProtocolInstances.EntityData.Children = types.NewOrderedMap()
    ospfv2ProtocolInstances.EntityData.Children.Append("ospfv2-protocol-instance", types.YChild{"Ospfv2ProtocolInstance", nil})
    for i := range ospfv2ProtocolInstances.Ospfv2ProtocolInstance {
        ospfv2ProtocolInstances.EntityData.Children.Append(types.GetSegmentPath(ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i]), types.YChild{"Ospfv2ProtocolInstance", ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i]})
    }
    ospfv2ProtocolInstances.EntityData.Leafs = types.NewOrderedMap()

    ospfv2ProtocolInstances.EntityData.YListKeys = []string {}

    return &(ospfv2ProtocolInstances.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
// Protocol samples for a particular OSPF v2
// instance
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v2 instance.
    Samples PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
}

func (ospfv2ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetEntityData() *types.CommonEntityData {
    ospfv2ProtocolInstance.EntityData.YFilter = ospfv2ProtocolInstance.YFilter
    ospfv2ProtocolInstance.EntityData.YangName = "ospfv2-protocol-instance"
    ospfv2ProtocolInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfv2ProtocolInstance.EntityData.ParentYangName = "ospfv2-protocol-instances"
    ospfv2ProtocolInstance.EntityData.SegmentPath = "ospfv2-protocol-instance" + types.AddKeyToken(ospfv2ProtocolInstance.InstanceName, "instance-name")
    ospfv2ProtocolInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2ProtocolInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2ProtocolInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2ProtocolInstance.EntityData.Children = types.NewOrderedMap()
    ospfv2ProtocolInstance.EntityData.Children.Append("samples", types.YChild{"Samples", &ospfv2ProtocolInstance.Samples})
    ospfv2ProtocolInstance.EntityData.Leafs = types.NewOrderedMap()
    ospfv2ProtocolInstance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", ospfv2ProtocolInstance.InstanceName})

    ospfv2ProtocolInstance.EntityData.YListKeys = []string {"InstanceName"}

    return &(ospfv2ProtocolInstance.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
// Sample Table for an OSPV v2 instance
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "ospfv2-protocol-instance"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-packets", types.YLeaf{"InputPackets", sample.InputPackets})
    sample.EntityData.Leafs.Append("output-packets", types.YLeaf{"OutputPackets", sample.OutputPackets})
    sample.EntityData.Leafs.Append("input-hello-packets", types.YLeaf{"InputHelloPackets", sample.InputHelloPackets})
    sample.EntityData.Leafs.Append("output-hello-packets", types.YLeaf{"OutputHelloPackets", sample.OutputHelloPackets})
    sample.EntityData.Leafs.Append("input-db-ds", types.YLeaf{"InputDbDs", sample.InputDbDs})
    sample.EntityData.Leafs.Append("input-db-ds-lsa", types.YLeaf{"InputDbDsLsa", sample.InputDbDsLsa})
    sample.EntityData.Leafs.Append("output-db-ds", types.YLeaf{"OutputDbDs", sample.OutputDbDs})
    sample.EntityData.Leafs.Append("output-db-ds-lsa", types.YLeaf{"OutputDbDsLsa", sample.OutputDbDsLsa})
    sample.EntityData.Leafs.Append("input-ls-requests", types.YLeaf{"InputLsRequests", sample.InputLsRequests})
    sample.EntityData.Leafs.Append("input-ls-requests-lsa", types.YLeaf{"InputLsRequestsLsa", sample.InputLsRequestsLsa})
    sample.EntityData.Leafs.Append("output-ls-requests", types.YLeaf{"OutputLsRequests", sample.OutputLsRequests})
    sample.EntityData.Leafs.Append("output-ls-requests-lsa", types.YLeaf{"OutputLsRequestsLsa", sample.OutputLsRequestsLsa})
    sample.EntityData.Leafs.Append("input-lsa-updates", types.YLeaf{"InputLsaUpdates", sample.InputLsaUpdates})
    sample.EntityData.Leafs.Append("input-lsa-updates-lsa", types.YLeaf{"InputLsaUpdatesLsa", sample.InputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("output-lsa-updates", types.YLeaf{"OutputLsaUpdates", sample.OutputLsaUpdates})
    sample.EntityData.Leafs.Append("output-lsa-updates-lsa", types.YLeaf{"OutputLsaUpdatesLsa", sample.OutputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("input-lsa-acks", types.YLeaf{"InputLsaAcks", sample.InputLsaAcks})
    sample.EntityData.Leafs.Append("input-lsa-acks-lsa", types.YLeaf{"InputLsaAcksLsa", sample.InputLsaAcksLsa})
    sample.EntityData.Leafs.Append("output-lsa-acks", types.YLeaf{"OutputLsaAcks", sample.OutputLsaAcks})
    sample.EntityData.Leafs.Append("output-lsa-acks-lsa", types.YLeaf{"OutputLsaAcksLsa", sample.OutputLsaAcksLsa})
    sample.EntityData.Leafs.Append("checksum-errors", types.YLeaf{"ChecksumErrors", sample.ChecksumErrors})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances
// OSPF v3 instances for which protocol statistics
// are collected
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v3 instance. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance.
    Ospfv3ProtocolInstance []*PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
}

func (ospfv3ProtocolInstances *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances) GetEntityData() *types.CommonEntityData {
    ospfv3ProtocolInstances.EntityData.YFilter = ospfv3ProtocolInstances.YFilter
    ospfv3ProtocolInstances.EntityData.YangName = "ospfv3-protocol-instances"
    ospfv3ProtocolInstances.EntityData.BundleName = "cisco_ios_xr"
    ospfv3ProtocolInstances.EntityData.ParentYangName = "ospf"
    ospfv3ProtocolInstances.EntityData.SegmentPath = "ospfv3-protocol-instances"
    ospfv3ProtocolInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3ProtocolInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3ProtocolInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3ProtocolInstances.EntityData.Children = types.NewOrderedMap()
    ospfv3ProtocolInstances.EntityData.Children.Append("ospfv3-protocol-instance", types.YChild{"Ospfv3ProtocolInstance", nil})
    for i := range ospfv3ProtocolInstances.Ospfv3ProtocolInstance {
        ospfv3ProtocolInstances.EntityData.Children.Append(types.GetSegmentPath(ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i]), types.YChild{"Ospfv3ProtocolInstance", ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i]})
    }
    ospfv3ProtocolInstances.EntityData.Leafs = types.NewOrderedMap()

    ospfv3ProtocolInstances.EntityData.YListKeys = []string {}

    return &(ospfv3ProtocolInstances.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
// Protocol samples for a particular OSPF v3
// instance
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v3 instance.
    Samples PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
}

func (ospfv3ProtocolInstance *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetEntityData() *types.CommonEntityData {
    ospfv3ProtocolInstance.EntityData.YFilter = ospfv3ProtocolInstance.YFilter
    ospfv3ProtocolInstance.EntityData.YangName = "ospfv3-protocol-instance"
    ospfv3ProtocolInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfv3ProtocolInstance.EntityData.ParentYangName = "ospfv3-protocol-instances"
    ospfv3ProtocolInstance.EntityData.SegmentPath = "ospfv3-protocol-instance" + types.AddKeyToken(ospfv3ProtocolInstance.InstanceName, "instance-name")
    ospfv3ProtocolInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3ProtocolInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3ProtocolInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3ProtocolInstance.EntityData.Children = types.NewOrderedMap()
    ospfv3ProtocolInstance.EntityData.Children.Append("samples", types.YChild{"Samples", &ospfv3ProtocolInstance.Samples})
    ospfv3ProtocolInstance.EntityData.Leafs = types.NewOrderedMap()
    ospfv3ProtocolInstance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", ospfv3ProtocolInstance.InstanceName})

    ospfv3ProtocolInstance.EntityData.YListKeys = []string {"InstanceName"}

    return &(ospfv3ProtocolInstance.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
// Sample Table for an OSPV v3 instance
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "ospfv3-protocol-instance"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-packets", types.YLeaf{"InputPackets", sample.InputPackets})
    sample.EntityData.Leafs.Append("output-packets", types.YLeaf{"OutputPackets", sample.OutputPackets})
    sample.EntityData.Leafs.Append("input-hello-packets", types.YLeaf{"InputHelloPackets", sample.InputHelloPackets})
    sample.EntityData.Leafs.Append("output-hello-packets", types.YLeaf{"OutputHelloPackets", sample.OutputHelloPackets})
    sample.EntityData.Leafs.Append("input-db-ds", types.YLeaf{"InputDbDs", sample.InputDbDs})
    sample.EntityData.Leafs.Append("input-db-ds-lsa", types.YLeaf{"InputDbDsLsa", sample.InputDbDsLsa})
    sample.EntityData.Leafs.Append("output-db-ds", types.YLeaf{"OutputDbDs", sample.OutputDbDs})
    sample.EntityData.Leafs.Append("output-db-ds-lsa", types.YLeaf{"OutputDbDsLsa", sample.OutputDbDsLsa})
    sample.EntityData.Leafs.Append("input-ls-requests", types.YLeaf{"InputLsRequests", sample.InputLsRequests})
    sample.EntityData.Leafs.Append("input-ls-requests-lsa", types.YLeaf{"InputLsRequestsLsa", sample.InputLsRequestsLsa})
    sample.EntityData.Leafs.Append("output-ls-requests", types.YLeaf{"OutputLsRequests", sample.OutputLsRequests})
    sample.EntityData.Leafs.Append("output-ls-requests-lsa", types.YLeaf{"OutputLsRequestsLsa", sample.OutputLsRequestsLsa})
    sample.EntityData.Leafs.Append("input-lsa-updates", types.YLeaf{"InputLsaUpdates", sample.InputLsaUpdates})
    sample.EntityData.Leafs.Append("input-lsa-updates-lsa", types.YLeaf{"InputLsaUpdatesLsa", sample.InputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("output-lsa-updates", types.YLeaf{"OutputLsaUpdates", sample.OutputLsaUpdates})
    sample.EntityData.Leafs.Append("output-lsa-updates-lsa", types.YLeaf{"OutputLsaUpdatesLsa", sample.OutputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("input-lsa-acks", types.YLeaf{"InputLsaAcks", sample.InputLsaAcks})
    sample.EntityData.Leafs.Append("input-lsa-acks-lsa", types.YLeaf{"InputLsaAcksLsa", sample.InputLsaAcksLsa})
    sample.EntityData.Leafs.Append("output-lsa-acks", types.YLeaf{"OutputLsaAcks", sample.OutputLsaAcks})
    sample.EntityData.Leafs.Append("output-lsa-acks-lsa", types.YLeaf{"OutputLsaAcksLsa", sample.OutputLsaAcksLsa})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Mpls
// Collected MPLS data
type PerfMgmt_Periodic_Mpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP neighbors for which statistics are collected.
    LdpNeighbors PerfMgmt_Periodic_Mpls_LdpNeighbors
}

func (mpls *PerfMgmt_Periodic_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "cisco_ios_xr"
    mpls.EntityData.ParentYangName = "periodic"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpls.EntityData.Children = types.NewOrderedMap()
    mpls.EntityData.Children.Append("ldp-neighbors", types.YChild{"LdpNeighbors", &mpls.LdpNeighbors})
    mpls.EntityData.Leafs = types.NewOrderedMap()

    mpls.EntityData.YListKeys = []string {}

    return &(mpls.EntityData)
}

// PerfMgmt_Periodic_Mpls_LdpNeighbors
// LDP neighbors for which statistics are
// collected
type PerfMgmt_Periodic_Mpls_LdpNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular LDP neighbor. The type is slice of
    // PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor.
    LdpNeighbor []*PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor
}

func (ldpNeighbors *PerfMgmt_Periodic_Mpls_LdpNeighbors) GetEntityData() *types.CommonEntityData {
    ldpNeighbors.EntityData.YFilter = ldpNeighbors.YFilter
    ldpNeighbors.EntityData.YangName = "ldp-neighbors"
    ldpNeighbors.EntityData.BundleName = "cisco_ios_xr"
    ldpNeighbors.EntityData.ParentYangName = "mpls"
    ldpNeighbors.EntityData.SegmentPath = "ldp-neighbors"
    ldpNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpNeighbors.EntityData.Children = types.NewOrderedMap()
    ldpNeighbors.EntityData.Children.Append("ldp-neighbor", types.YChild{"LdpNeighbor", nil})
    for i := range ldpNeighbors.LdpNeighbor {
        ldpNeighbors.EntityData.Children.Append(types.GetSegmentPath(ldpNeighbors.LdpNeighbor[i]), types.YChild{"LdpNeighbor", ldpNeighbors.LdpNeighbor[i]})
    }
    ldpNeighbors.EntityData.Leafs = types.NewOrderedMap()

    ldpNeighbors.EntityData.YListKeys = []string {}

    return &(ldpNeighbors.EntityData)
}

// PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor
// Samples for a particular LDP neighbor
type PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Nbr interface{}

    // Samples for a particular LDP neighbor.
    Samples PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples
}

func (ldpNeighbor *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor) GetEntityData() *types.CommonEntityData {
    ldpNeighbor.EntityData.YFilter = ldpNeighbor.YFilter
    ldpNeighbor.EntityData.YangName = "ldp-neighbor"
    ldpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    ldpNeighbor.EntityData.ParentYangName = "ldp-neighbors"
    ldpNeighbor.EntityData.SegmentPath = "ldp-neighbor" + types.AddKeyToken(ldpNeighbor.Nbr, "nbr")
    ldpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpNeighbor.EntityData.Children = types.NewOrderedMap()
    ldpNeighbor.EntityData.Children.Append("samples", types.YChild{"Samples", &ldpNeighbor.Samples})
    ldpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    ldpNeighbor.EntityData.Leafs.Append("nbr", types.YLeaf{"Nbr", ldpNeighbor.Nbr})

    ldpNeighbor.EntityData.YListKeys = []string {"Nbr"}

    return &(ldpNeighbor.EntityData)
}

// PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples
// Samples for a particular LDP neighbor
type PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP neighbor statistics sample. The type is slice of
    // PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "ldp-neighbor"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
// LDP neighbor statistics sample
type PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("total-msgs-sent", types.YLeaf{"TotalMsgsSent", sample.TotalMsgsSent})
    sample.EntityData.Leafs.Append("total-msgs-rcvd", types.YLeaf{"TotalMsgsRcvd", sample.TotalMsgsRcvd})
    sample.EntityData.Leafs.Append("init-msgs-sent", types.YLeaf{"InitMsgsSent", sample.InitMsgsSent})
    sample.EntityData.Leafs.Append("init-msgs-rcvd", types.YLeaf{"InitMsgsRcvd", sample.InitMsgsRcvd})
    sample.EntityData.Leafs.Append("address-msgs-sent", types.YLeaf{"AddressMsgsSent", sample.AddressMsgsSent})
    sample.EntityData.Leafs.Append("address-msgs-rcvd", types.YLeaf{"AddressMsgsRcvd", sample.AddressMsgsRcvd})
    sample.EntityData.Leafs.Append("address-withdraw-msgs-sent", types.YLeaf{"AddressWithdrawMsgsSent", sample.AddressWithdrawMsgsSent})
    sample.EntityData.Leafs.Append("address-withdraw-msgs-rcvd", types.YLeaf{"AddressWithdrawMsgsRcvd", sample.AddressWithdrawMsgsRcvd})
    sample.EntityData.Leafs.Append("label-mapping-msgs-sent", types.YLeaf{"LabelMappingMsgsSent", sample.LabelMappingMsgsSent})
    sample.EntityData.Leafs.Append("label-mapping-msgs-rcvd", types.YLeaf{"LabelMappingMsgsRcvd", sample.LabelMappingMsgsRcvd})
    sample.EntityData.Leafs.Append("label-withdraw-msgs-sent", types.YLeaf{"LabelWithdrawMsgsSent", sample.LabelWithdrawMsgsSent})
    sample.EntityData.Leafs.Append("label-withdraw-msgs-rcvd", types.YLeaf{"LabelWithdrawMsgsRcvd", sample.LabelWithdrawMsgsRcvd})
    sample.EntityData.Leafs.Append("label-release-msgs-sent", types.YLeaf{"LabelReleaseMsgsSent", sample.LabelReleaseMsgsSent})
    sample.EntityData.Leafs.Append("label-release-msgs-rcvd", types.YLeaf{"LabelReleaseMsgsRcvd", sample.LabelReleaseMsgsRcvd})
    sample.EntityData.Leafs.Append("notification-msgs-sent", types.YLeaf{"NotificationMsgsSent", sample.NotificationMsgsSent})
    sample.EntityData.Leafs.Append("notification-msgs-rcvd", types.YLeaf{"NotificationMsgsRcvd", sample.NotificationMsgsRcvd})
    sample.EntityData.Leafs.Append("keepalive-msgs-sent", types.YLeaf{"KeepaliveMsgsSent", sample.KeepaliveMsgsSent})
    sample.EntityData.Leafs.Append("keepalive-msgs-rcvd", types.YLeaf{"KeepaliveMsgsRcvd", sample.KeepaliveMsgsRcvd})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Nodes
// Nodes for which data is collected
type PerfMgmt_Periodic_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Instance. The type is slice of PerfMgmt_Periodic_Nodes_Node.
    Node []*PerfMgmt_Periodic_Nodes_Node
}

func (nodes *PerfMgmt_Periodic_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "periodic"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node
// Node Instance
type PerfMgmt_Periodic_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *PerfMgmt_Periodic_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("sample-xr", types.YChild{"SampleXr", &node.SampleXr})
    node.EntityData.Children.Append("processes", types.YChild{"Processes", &node.Processes})
    node.EntityData.Children.Append("samples", types.YChild{"Samples", &node.Samples})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})

    node.EntityData.YListKeys = []string {"NodeId"}

    return &(node.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_SampleXr
// Node CPU data
type PerfMgmt_Periodic_Nodes_Node_SampleXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node CPU data sample. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample.
    Sample []*PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample
}

func (sampleXr *PerfMgmt_Periodic_Nodes_Node_SampleXr) GetEntityData() *types.CommonEntityData {
    sampleXr.EntityData.YFilter = sampleXr.YFilter
    sampleXr.EntityData.YangName = "sample-xr"
    sampleXr.EntityData.BundleName = "cisco_ios_xr"
    sampleXr.EntityData.ParentYangName = "node"
    sampleXr.EntityData.SegmentPath = "sample-xr"
    sampleXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sampleXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sampleXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sampleXr.EntityData.Children = types.NewOrderedMap()
    sampleXr.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range sampleXr.Sample {
        sampleXr.EntityData.Children.Append(types.GetSegmentPath(sampleXr.Sample[i]), types.YChild{"Sample", sampleXr.Sample[i]})
    }
    sampleXr.EntityData.Leafs = types.NewOrderedMap()

    sampleXr.EntityData.YListKeys = []string {}

    return &(sampleXr.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample
// Node CPU data sample
type PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Nodes_Node_SampleXr_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "sample-xr"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("no-processes", types.YLeaf{"NoProcesses", sample.NoProcesses})
    sample.EntityData.Leafs.Append("average-cpu-used", types.YLeaf{"AverageCpuUsed", sample.AverageCpuUsed})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_Processes
// Processes data
type PerfMgmt_Periodic_Nodes_Node_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process data. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_Processes_Process.
    Process []*PerfMgmt_Periodic_Nodes_Node_Processes_Process
}

func (processes *PerfMgmt_Periodic_Nodes_Node_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "node"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range processes.Process {
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.Process[i]), types.YChild{"Process", processes.Process[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_Processes_Process
// Process data
type PerfMgmt_Periodic_Nodes_Node_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Process data.
    Samples PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples
}

func (process *PerfMgmt_Periodic_Nodes_Node_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessId, "process-id")
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("samples", types.YChild{"Samples", &process.Samples})
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", process.ProcessId})

    process.EntityData.YListKeys = []string {"ProcessId"}

    return &(process.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples
// Process data
type PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process data sample. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "process"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample
// Process data sample
type PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Nodes_Node_Processes_Process_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("peak-memory", types.YLeaf{"PeakMemory", sample.PeakMemory})
    sample.EntityData.Leafs.Append("average-cpu-used", types.YLeaf{"AverageCpuUsed", sample.AverageCpuUsed})
    sample.EntityData.Leafs.Append("no-threads", types.YLeaf{"NoThreads", sample.NoThreads})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_Samples
// Node Memory data
type PerfMgmt_Periodic_Nodes_Node_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory data sample. The type is slice of
    // PerfMgmt_Periodic_Nodes_Node_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Nodes_Node_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Nodes_Node_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "node"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Nodes_Node_Samples_Sample
// Node Memory data sample
type PerfMgmt_Periodic_Nodes_Node_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Nodes_Node_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("curr-memory", types.YLeaf{"CurrMemory", sample.CurrMemory})
    sample.EntityData.Leafs.Append("peak-memory", types.YLeaf{"PeakMemory", sample.PeakMemory})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Bgp
// Collected BGP data
type PerfMgmt_Periodic_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbors for which statistics are collected.
    BgpNeighbors PerfMgmt_Periodic_Bgp_BgpNeighbors
}

func (bgp *PerfMgmt_Periodic_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "periodic"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("bgp-neighbors", types.YChild{"BgpNeighbors", &bgp.BgpNeighbors})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PerfMgmt_Periodic_Bgp_BgpNeighbors
// Neighbors for which statistics are collected
type PerfMgmt_Periodic_Bgp_BgpNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for particular neighbor. The type is slice of
    // PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor.
    BgpNeighbor []*PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor
}

func (bgpNeighbors *PerfMgmt_Periodic_Bgp_BgpNeighbors) GetEntityData() *types.CommonEntityData {
    bgpNeighbors.EntityData.YFilter = bgpNeighbors.YFilter
    bgpNeighbors.EntityData.YangName = "bgp-neighbors"
    bgpNeighbors.EntityData.BundleName = "cisco_ios_xr"
    bgpNeighbors.EntityData.ParentYangName = "bgp"
    bgpNeighbors.EntityData.SegmentPath = "bgp-neighbors"
    bgpNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpNeighbors.EntityData.Children = types.NewOrderedMap()
    bgpNeighbors.EntityData.Children.Append("bgp-neighbor", types.YChild{"BgpNeighbor", nil})
    for i := range bgpNeighbors.BgpNeighbor {
        bgpNeighbors.EntityData.Children.Append(types.GetSegmentPath(bgpNeighbors.BgpNeighbor[i]), types.YChild{"BgpNeighbor", bgpNeighbors.BgpNeighbor[i]})
    }
    bgpNeighbors.EntityData.Leafs = types.NewOrderedMap()

    bgpNeighbors.EntityData.YListKeys = []string {}

    return &(bgpNeighbors.EntityData)
}

// PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor
// Samples for particular neighbor
type PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP Neighbor Identifier. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Sample Table for a BGP neighbor.
    Samples PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples
}

func (bgpNeighbor *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor) GetEntityData() *types.CommonEntityData {
    bgpNeighbor.EntityData.YFilter = bgpNeighbor.YFilter
    bgpNeighbor.EntityData.YangName = "bgp-neighbor"
    bgpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    bgpNeighbor.EntityData.ParentYangName = "bgp-neighbors"
    bgpNeighbor.EntityData.SegmentPath = "bgp-neighbor" + types.AddKeyToken(bgpNeighbor.IpAddress, "ip-address")
    bgpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpNeighbor.EntityData.Children = types.NewOrderedMap()
    bgpNeighbor.EntityData.Children.Append("samples", types.YChild{"Samples", &bgpNeighbor.Samples})
    bgpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    bgpNeighbor.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", bgpNeighbor.IpAddress})

    bgpNeighbor.EntityData.YListKeys = []string {"IpAddress"}

    return &(bgpNeighbor.EntityData)
}

// PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples
// Sample Table for a BGP neighbor
type PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor statistics sample. The type is slice of
    // PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "bgp-neighbor"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
// Neighbor statistics sample
type PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-messages", types.YLeaf{"InputMessages", sample.InputMessages})
    sample.EntityData.Leafs.Append("output-messages", types.YLeaf{"OutputMessages", sample.OutputMessages})
    sample.EntityData.Leafs.Append("input-update-messages", types.YLeaf{"InputUpdateMessages", sample.InputUpdateMessages})
    sample.EntityData.Leafs.Append("output-update-messages", types.YLeaf{"OutputUpdateMessages", sample.OutputUpdateMessages})
    sample.EntityData.Leafs.Append("conn-established", types.YLeaf{"ConnEstablished", sample.ConnEstablished})
    sample.EntityData.Leafs.Append("conn-dropped", types.YLeaf{"ConnDropped", sample.ConnDropped})
    sample.EntityData.Leafs.Append("errors-received", types.YLeaf{"ErrorsReceived", sample.ErrorsReceived})
    sample.EntityData.Leafs.Append("errors-sent", types.YLeaf{"ErrorsSent", sample.ErrorsSent})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Interface
// Collected Interface data
type PerfMgmt_Periodic_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces for which Generic Counters are collected.
    GenericCounterInterfaces PerfMgmt_Periodic_Interface_GenericCounterInterfaces

    // Interfaces for which Basic Counters are collected.
    BasicCounterInterfaces PerfMgmt_Periodic_Interface_BasicCounterInterfaces

    // Interfaces for which Data Rates are collected.
    DataRateInterfaces PerfMgmt_Periodic_Interface_DataRateInterfaces
}

func (self *PerfMgmt_Periodic_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "periodic"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("generic-counter-interfaces", types.YChild{"GenericCounterInterfaces", &self.GenericCounterInterfaces})
    self.EntityData.Children.Append("basic-counter-interfaces", types.YChild{"BasicCounterInterfaces", &self.BasicCounterInterfaces})
    self.EntityData.Children.Append("data-rate-interfaces", types.YChild{"DataRateInterfaces", &self.DataRateInterfaces})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces
// Interfaces for which Generic Counters are
// collected
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface.
    GenericCounterInterface []*PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface
}

func (genericCounterInterfaces *PerfMgmt_Periodic_Interface_GenericCounterInterfaces) GetEntityData() *types.CommonEntityData {
    genericCounterInterfaces.EntityData.YFilter = genericCounterInterfaces.YFilter
    genericCounterInterfaces.EntityData.YangName = "generic-counter-interfaces"
    genericCounterInterfaces.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterfaces.EntityData.ParentYangName = "interface"
    genericCounterInterfaces.EntityData.SegmentPath = "generic-counter-interfaces"
    genericCounterInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterfaces.EntityData.Children = types.NewOrderedMap()
    genericCounterInterfaces.EntityData.Children.Append("generic-counter-interface", types.YChild{"GenericCounterInterface", nil})
    for i := range genericCounterInterfaces.GenericCounterInterface {
        genericCounterInterfaces.EntityData.Children.Append(types.GetSegmentPath(genericCounterInterfaces.GenericCounterInterface[i]), types.YChild{"GenericCounterInterface", genericCounterInterfaces.GenericCounterInterface[i]})
    }
    genericCounterInterfaces.EntityData.Leafs = types.NewOrderedMap()

    genericCounterInterfaces.EntityData.YListKeys = []string {}

    return &(genericCounterInterfaces.EntityData)
}

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface
// Samples for a particular interface
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Generic Counter samples for an interface.
    Samples PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
}

func (genericCounterInterface *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface) GetEntityData() *types.CommonEntityData {
    genericCounterInterface.EntityData.YFilter = genericCounterInterface.YFilter
    genericCounterInterface.EntityData.YangName = "generic-counter-interface"
    genericCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterface.EntityData.ParentYangName = "generic-counter-interfaces"
    genericCounterInterface.EntityData.SegmentPath = "generic-counter-interface" + types.AddKeyToken(genericCounterInterface.InterfaceName, "interface-name")
    genericCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterface.EntityData.Children = types.NewOrderedMap()
    genericCounterInterface.EntityData.Children.Append("samples", types.YChild{"Samples", &genericCounterInterface.Samples})
    genericCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    genericCounterInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", genericCounterInterface.InterfaceName})

    genericCounterInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
// Generic Counter samples for an interface
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "generic-counter-interface"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
// Generic Counters sample
type PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("in-packets", types.YLeaf{"InPackets", sample.InPackets})
    sample.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", sample.InOctets})
    sample.EntityData.Leafs.Append("out-packets", types.YLeaf{"OutPackets", sample.OutPackets})
    sample.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", sample.OutOctets})
    sample.EntityData.Leafs.Append("in-ucast-pkts", types.YLeaf{"InUcastPkts", sample.InUcastPkts})
    sample.EntityData.Leafs.Append("in-multicast-pkts", types.YLeaf{"InMulticastPkts", sample.InMulticastPkts})
    sample.EntityData.Leafs.Append("in-broadcast-pkts", types.YLeaf{"InBroadcastPkts", sample.InBroadcastPkts})
    sample.EntityData.Leafs.Append("out-ucast-pkts", types.YLeaf{"OutUcastPkts", sample.OutUcastPkts})
    sample.EntityData.Leafs.Append("out-multicast-pkts", types.YLeaf{"OutMulticastPkts", sample.OutMulticastPkts})
    sample.EntityData.Leafs.Append("out-broadcast-pkts", types.YLeaf{"OutBroadcastPkts", sample.OutBroadcastPkts})
    sample.EntityData.Leafs.Append("output-total-drops", types.YLeaf{"OutputTotalDrops", sample.OutputTotalDrops})
    sample.EntityData.Leafs.Append("input-total-drops", types.YLeaf{"InputTotalDrops", sample.InputTotalDrops})
    sample.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", sample.InputQueueDrops})
    sample.EntityData.Leafs.Append("input-unknown-proto", types.YLeaf{"InputUnknownProto", sample.InputUnknownProto})
    sample.EntityData.Leafs.Append("output-total-errors", types.YLeaf{"OutputTotalErrors", sample.OutputTotalErrors})
    sample.EntityData.Leafs.Append("output-underrun", types.YLeaf{"OutputUnderrun", sample.OutputUnderrun})
    sample.EntityData.Leafs.Append("input-total-errors", types.YLeaf{"InputTotalErrors", sample.InputTotalErrors})
    sample.EntityData.Leafs.Append("input-crc", types.YLeaf{"InputCrc", sample.InputCrc})
    sample.EntityData.Leafs.Append("input-overrun", types.YLeaf{"InputOverrun", sample.InputOverrun})
    sample.EntityData.Leafs.Append("input-frame", types.YLeaf{"InputFrame", sample.InputFrame})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces
// Interfaces for which Basic Counters are
// collected
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface.
    BasicCounterInterface []*PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface
}

func (basicCounterInterfaces *PerfMgmt_Periodic_Interface_BasicCounterInterfaces) GetEntityData() *types.CommonEntityData {
    basicCounterInterfaces.EntityData.YFilter = basicCounterInterfaces.YFilter
    basicCounterInterfaces.EntityData.YangName = "basic-counter-interfaces"
    basicCounterInterfaces.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterfaces.EntityData.ParentYangName = "interface"
    basicCounterInterfaces.EntityData.SegmentPath = "basic-counter-interfaces"
    basicCounterInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterfaces.EntityData.Children = types.NewOrderedMap()
    basicCounterInterfaces.EntityData.Children.Append("basic-counter-interface", types.YChild{"BasicCounterInterface", nil})
    for i := range basicCounterInterfaces.BasicCounterInterface {
        basicCounterInterfaces.EntityData.Children.Append(types.GetSegmentPath(basicCounterInterfaces.BasicCounterInterface[i]), types.YChild{"BasicCounterInterface", basicCounterInterfaces.BasicCounterInterface[i]})
    }
    basicCounterInterfaces.EntityData.Leafs = types.NewOrderedMap()

    basicCounterInterfaces.EntityData.YListKeys = []string {}

    return &(basicCounterInterfaces.EntityData)
}

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface
// Samples for a particular interface
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Basic Counter samples for an interface.
    Samples PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
}

func (basicCounterInterface *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface) GetEntityData() *types.CommonEntityData {
    basicCounterInterface.EntityData.YFilter = basicCounterInterface.YFilter
    basicCounterInterface.EntityData.YangName = "basic-counter-interface"
    basicCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterface.EntityData.ParentYangName = "basic-counter-interfaces"
    basicCounterInterface.EntityData.SegmentPath = "basic-counter-interface" + types.AddKeyToken(basicCounterInterface.InterfaceName, "interface-name")
    basicCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterface.EntityData.Children = types.NewOrderedMap()
    basicCounterInterface.EntityData.Children.Append("samples", types.YChild{"Samples", &basicCounterInterface.Samples})
    basicCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    basicCounterInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", basicCounterInterface.InterfaceName})

    basicCounterInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
// Basic Counter samples for an interface
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Basic Counters sample. The type is slice of
    // PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "basic-counter-interface"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
// Basic Counters sample
type PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("in-packets", types.YLeaf{"InPackets", sample.InPackets})
    sample.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", sample.InOctets})
    sample.EntityData.Leafs.Append("out-packets", types.YLeaf{"OutPackets", sample.OutPackets})
    sample.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", sample.OutOctets})
    sample.EntityData.Leafs.Append("input-total-drops", types.YLeaf{"InputTotalDrops", sample.InputTotalDrops})
    sample.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", sample.InputQueueDrops})
    sample.EntityData.Leafs.Append("input-total-errors", types.YLeaf{"InputTotalErrors", sample.InputTotalErrors})
    sample.EntityData.Leafs.Append("output-total-drops", types.YLeaf{"OutputTotalDrops", sample.OutputTotalDrops})
    sample.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", sample.OutputQueueDrops})
    sample.EntityData.Leafs.Append("output-total-errors", types.YLeaf{"OutputTotalErrors", sample.OutputTotalErrors})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Periodic_Interface_DataRateInterfaces
// Interfaces for which Data Rates are collected
type PerfMgmt_Periodic_Interface_DataRateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface.
    DataRateInterface []*PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface
}

func (dataRateInterfaces *PerfMgmt_Periodic_Interface_DataRateInterfaces) GetEntityData() *types.CommonEntityData {
    dataRateInterfaces.EntityData.YFilter = dataRateInterfaces.YFilter
    dataRateInterfaces.EntityData.YangName = "data-rate-interfaces"
    dataRateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterfaces.EntityData.ParentYangName = "interface"
    dataRateInterfaces.EntityData.SegmentPath = "data-rate-interfaces"
    dataRateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterfaces.EntityData.Children = types.NewOrderedMap()
    dataRateInterfaces.EntityData.Children.Append("data-rate-interface", types.YChild{"DataRateInterface", nil})
    for i := range dataRateInterfaces.DataRateInterface {
        dataRateInterfaces.EntityData.Children.Append(types.GetSegmentPath(dataRateInterfaces.DataRateInterface[i]), types.YChild{"DataRateInterface", dataRateInterfaces.DataRateInterface[i]})
    }
    dataRateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    dataRateInterfaces.EntityData.YListKeys = []string {}

    return &(dataRateInterfaces.EntityData)
}

// PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface
// Samples for a particular interface
type PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Data Rate samples for an interface.
    Samples PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples
}

func (dataRateInterface *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface) GetEntityData() *types.CommonEntityData {
    dataRateInterface.EntityData.YFilter = dataRateInterface.YFilter
    dataRateInterface.EntityData.YangName = "data-rate-interface"
    dataRateInterface.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterface.EntityData.ParentYangName = "data-rate-interfaces"
    dataRateInterface.EntityData.SegmentPath = "data-rate-interface" + types.AddKeyToken(dataRateInterface.InterfaceName, "interface-name")
    dataRateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterface.EntityData.Children = types.NewOrderedMap()
    dataRateInterface.EntityData.Children.Append("samples", types.YChild{"Samples", &dataRateInterface.Samples})
    dataRateInterface.EntityData.Leafs = types.NewOrderedMap()
    dataRateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", dataRateInterface.InterfaceName})

    dataRateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples
// Data Rate samples for an interface
type PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data Rates sample. The type is slice of
    // PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample.
    Sample []*PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
}

func (samples *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "data-rate-interface"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
// Data Rates sample
type PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Periodic_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", sample.InputDataRate})
    sample.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", sample.InputPacketRate})
    sample.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", sample.OutputDataRate})
    sample.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", sample.OutputPacketRate})
    sample.EntityData.Leafs.Append("input-peak-rate", types.YLeaf{"InputPeakRate", sample.InputPeakRate})
    sample.EntityData.Leafs.Append("input-peak-pkts", types.YLeaf{"InputPeakPkts", sample.InputPeakPkts})
    sample.EntityData.Leafs.Append("output-peak-rate", types.YLeaf{"OutputPeakRate", sample.OutputPeakRate})
    sample.EntityData.Leafs.Append("output-peak-pkts", types.YLeaf{"OutputPeakPkts", sample.OutputPeakPkts})
    sample.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", sample.Bandwidth})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor
// Data from monitor (one history period) requests
type PerfMgmt_Monitor struct {
    EntityData types.CommonEntityData
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

func (monitor *PerfMgmt_Monitor) GetEntityData() *types.CommonEntityData {
    monitor.EntityData.YFilter = monitor.YFilter
    monitor.EntityData.YangName = "monitor"
    monitor.EntityData.BundleName = "cisco_ios_xr"
    monitor.EntityData.ParentYangName = "perf-mgmt"
    monitor.EntityData.SegmentPath = "monitor"
    monitor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitor.EntityData.Children = types.NewOrderedMap()
    monitor.EntityData.Children.Append("ospf", types.YChild{"Ospf", &monitor.Ospf})
    monitor.EntityData.Children.Append("mpls", types.YChild{"Mpls", &monitor.Mpls})
    monitor.EntityData.Children.Append("nodes", types.YChild{"Nodes", &monitor.Nodes})
    monitor.EntityData.Children.Append("bgp", types.YChild{"Bgp", &monitor.Bgp})
    monitor.EntityData.Children.Append("interface", types.YChild{"Interface", &monitor.Interface})
    monitor.EntityData.Leafs = types.NewOrderedMap()

    monitor.EntityData.YListKeys = []string {}

    return &(monitor.EntityData)
}

// PerfMgmt_Monitor_Ospf
// Collected OSPF data
type PerfMgmt_Monitor_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 instances for which protocol statistics are collected.
    Ospfv2ProtocolInstances PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances

    // OSPF v3 instances for which protocol statistics are collected.
    Ospfv3ProtocolInstances PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances
}

func (ospf *PerfMgmt_Monitor_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "monitor"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Children.Append("ospfv2-protocol-instances", types.YChild{"Ospfv2ProtocolInstances", &ospf.Ospfv2ProtocolInstances})
    ospf.EntityData.Children.Append("ospfv3-protocol-instances", types.YChild{"Ospfv3ProtocolInstances", &ospf.Ospfv3ProtocolInstances})
    ospf.EntityData.Leafs = types.NewOrderedMap()

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances
// OSPF v2 instances for which protocol statistics
// are collected
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v2 instance. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance.
    Ospfv2ProtocolInstance []*PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
}

func (ospfv2ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances) GetEntityData() *types.CommonEntityData {
    ospfv2ProtocolInstances.EntityData.YFilter = ospfv2ProtocolInstances.YFilter
    ospfv2ProtocolInstances.EntityData.YangName = "ospfv2-protocol-instances"
    ospfv2ProtocolInstances.EntityData.BundleName = "cisco_ios_xr"
    ospfv2ProtocolInstances.EntityData.ParentYangName = "ospf"
    ospfv2ProtocolInstances.EntityData.SegmentPath = "ospfv2-protocol-instances"
    ospfv2ProtocolInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2ProtocolInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2ProtocolInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2ProtocolInstances.EntityData.Children = types.NewOrderedMap()
    ospfv2ProtocolInstances.EntityData.Children.Append("ospfv2-protocol-instance", types.YChild{"Ospfv2ProtocolInstance", nil})
    for i := range ospfv2ProtocolInstances.Ospfv2ProtocolInstance {
        ospfv2ProtocolInstances.EntityData.Children.Append(types.GetSegmentPath(ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i]), types.YChild{"Ospfv2ProtocolInstance", ospfv2ProtocolInstances.Ospfv2ProtocolInstance[i]})
    }
    ospfv2ProtocolInstances.EntityData.Leafs = types.NewOrderedMap()

    ospfv2ProtocolInstances.EntityData.YListKeys = []string {}

    return &(ospfv2ProtocolInstances.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance
// Protocol samples for a particular OSPF v2
// instance
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v2 instance.
    Samples PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
}

func (ospfv2ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance) GetEntityData() *types.CommonEntityData {
    ospfv2ProtocolInstance.EntityData.YFilter = ospfv2ProtocolInstance.YFilter
    ospfv2ProtocolInstance.EntityData.YangName = "ospfv2-protocol-instance"
    ospfv2ProtocolInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfv2ProtocolInstance.EntityData.ParentYangName = "ospfv2-protocol-instances"
    ospfv2ProtocolInstance.EntityData.SegmentPath = "ospfv2-protocol-instance" + types.AddKeyToken(ospfv2ProtocolInstance.InstanceName, "instance-name")
    ospfv2ProtocolInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2ProtocolInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2ProtocolInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2ProtocolInstance.EntityData.Children = types.NewOrderedMap()
    ospfv2ProtocolInstance.EntityData.Children.Append("samples", types.YChild{"Samples", &ospfv2ProtocolInstance.Samples})
    ospfv2ProtocolInstance.EntityData.Leafs = types.NewOrderedMap()
    ospfv2ProtocolInstance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", ospfv2ProtocolInstance.InstanceName})

    ospfv2ProtocolInstance.EntityData.YListKeys = []string {"InstanceName"}

    return &(ospfv2ProtocolInstance.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples
// Sample Table for an OSPV v2 instance
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "ospfv2-protocol-instance"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Ospf_Ospfv2ProtocolInstances_Ospfv2ProtocolInstance_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-packets", types.YLeaf{"InputPackets", sample.InputPackets})
    sample.EntityData.Leafs.Append("output-packets", types.YLeaf{"OutputPackets", sample.OutputPackets})
    sample.EntityData.Leafs.Append("input-hello-packets", types.YLeaf{"InputHelloPackets", sample.InputHelloPackets})
    sample.EntityData.Leafs.Append("output-hello-packets", types.YLeaf{"OutputHelloPackets", sample.OutputHelloPackets})
    sample.EntityData.Leafs.Append("input-db-ds", types.YLeaf{"InputDbDs", sample.InputDbDs})
    sample.EntityData.Leafs.Append("input-db-ds-lsa", types.YLeaf{"InputDbDsLsa", sample.InputDbDsLsa})
    sample.EntityData.Leafs.Append("output-db-ds", types.YLeaf{"OutputDbDs", sample.OutputDbDs})
    sample.EntityData.Leafs.Append("output-db-ds-lsa", types.YLeaf{"OutputDbDsLsa", sample.OutputDbDsLsa})
    sample.EntityData.Leafs.Append("input-ls-requests", types.YLeaf{"InputLsRequests", sample.InputLsRequests})
    sample.EntityData.Leafs.Append("input-ls-requests-lsa", types.YLeaf{"InputLsRequestsLsa", sample.InputLsRequestsLsa})
    sample.EntityData.Leafs.Append("output-ls-requests", types.YLeaf{"OutputLsRequests", sample.OutputLsRequests})
    sample.EntityData.Leafs.Append("output-ls-requests-lsa", types.YLeaf{"OutputLsRequestsLsa", sample.OutputLsRequestsLsa})
    sample.EntityData.Leafs.Append("input-lsa-updates", types.YLeaf{"InputLsaUpdates", sample.InputLsaUpdates})
    sample.EntityData.Leafs.Append("input-lsa-updates-lsa", types.YLeaf{"InputLsaUpdatesLsa", sample.InputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("output-lsa-updates", types.YLeaf{"OutputLsaUpdates", sample.OutputLsaUpdates})
    sample.EntityData.Leafs.Append("output-lsa-updates-lsa", types.YLeaf{"OutputLsaUpdatesLsa", sample.OutputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("input-lsa-acks", types.YLeaf{"InputLsaAcks", sample.InputLsaAcks})
    sample.EntityData.Leafs.Append("input-lsa-acks-lsa", types.YLeaf{"InputLsaAcksLsa", sample.InputLsaAcksLsa})
    sample.EntityData.Leafs.Append("output-lsa-acks", types.YLeaf{"OutputLsaAcks", sample.OutputLsaAcks})
    sample.EntityData.Leafs.Append("output-lsa-acks-lsa", types.YLeaf{"OutputLsaAcksLsa", sample.OutputLsaAcksLsa})
    sample.EntityData.Leafs.Append("checksum-errors", types.YLeaf{"ChecksumErrors", sample.ChecksumErrors})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances
// OSPF v3 instances for which protocol statistics
// are collected
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol samples for a particular OSPF v3 instance. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance.
    Ospfv3ProtocolInstance []*PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
}

func (ospfv3ProtocolInstances *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances) GetEntityData() *types.CommonEntityData {
    ospfv3ProtocolInstances.EntityData.YFilter = ospfv3ProtocolInstances.YFilter
    ospfv3ProtocolInstances.EntityData.YangName = "ospfv3-protocol-instances"
    ospfv3ProtocolInstances.EntityData.BundleName = "cisco_ios_xr"
    ospfv3ProtocolInstances.EntityData.ParentYangName = "ospf"
    ospfv3ProtocolInstances.EntityData.SegmentPath = "ospfv3-protocol-instances"
    ospfv3ProtocolInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3ProtocolInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3ProtocolInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3ProtocolInstances.EntityData.Children = types.NewOrderedMap()
    ospfv3ProtocolInstances.EntityData.Children.Append("ospfv3-protocol-instance", types.YChild{"Ospfv3ProtocolInstance", nil})
    for i := range ospfv3ProtocolInstances.Ospfv3ProtocolInstance {
        ospfv3ProtocolInstances.EntityData.Children.Append(types.GetSegmentPath(ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i]), types.YChild{"Ospfv3ProtocolInstance", ospfv3ProtocolInstances.Ospfv3ProtocolInstance[i]})
    }
    ospfv3ProtocolInstances.EntityData.Leafs = types.NewOrderedMap()

    ospfv3ProtocolInstances.EntityData.YListKeys = []string {}

    return &(ospfv3ProtocolInstances.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance
// Protocol samples for a particular OSPF v3
// instance
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Sample Table for an OSPV v3 instance.
    Samples PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
}

func (ospfv3ProtocolInstance *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance) GetEntityData() *types.CommonEntityData {
    ospfv3ProtocolInstance.EntityData.YFilter = ospfv3ProtocolInstance.YFilter
    ospfv3ProtocolInstance.EntityData.YangName = "ospfv3-protocol-instance"
    ospfv3ProtocolInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfv3ProtocolInstance.EntityData.ParentYangName = "ospfv3-protocol-instances"
    ospfv3ProtocolInstance.EntityData.SegmentPath = "ospfv3-protocol-instance" + types.AddKeyToken(ospfv3ProtocolInstance.InstanceName, "instance-name")
    ospfv3ProtocolInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3ProtocolInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3ProtocolInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3ProtocolInstance.EntityData.Children = types.NewOrderedMap()
    ospfv3ProtocolInstance.EntityData.Children.Append("samples", types.YChild{"Samples", &ospfv3ProtocolInstance.Samples})
    ospfv3ProtocolInstance.EntityData.Leafs = types.NewOrderedMap()
    ospfv3ProtocolInstance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", ospfv3ProtocolInstance.InstanceName})

    ospfv3ProtocolInstance.EntityData.YListKeys = []string {"InstanceName"}

    return &(ospfv3ProtocolInstance.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples
// Sample Table for an OSPV v3 instance
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "ospfv3-protocol-instance"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample
// Generic Counters sample
type PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Ospf_Ospfv3ProtocolInstances_Ospfv3ProtocolInstance_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-packets", types.YLeaf{"InputPackets", sample.InputPackets})
    sample.EntityData.Leafs.Append("output-packets", types.YLeaf{"OutputPackets", sample.OutputPackets})
    sample.EntityData.Leafs.Append("input-hello-packets", types.YLeaf{"InputHelloPackets", sample.InputHelloPackets})
    sample.EntityData.Leafs.Append("output-hello-packets", types.YLeaf{"OutputHelloPackets", sample.OutputHelloPackets})
    sample.EntityData.Leafs.Append("input-db-ds", types.YLeaf{"InputDbDs", sample.InputDbDs})
    sample.EntityData.Leafs.Append("input-db-ds-lsa", types.YLeaf{"InputDbDsLsa", sample.InputDbDsLsa})
    sample.EntityData.Leafs.Append("output-db-ds", types.YLeaf{"OutputDbDs", sample.OutputDbDs})
    sample.EntityData.Leafs.Append("output-db-ds-lsa", types.YLeaf{"OutputDbDsLsa", sample.OutputDbDsLsa})
    sample.EntityData.Leafs.Append("input-ls-requests", types.YLeaf{"InputLsRequests", sample.InputLsRequests})
    sample.EntityData.Leafs.Append("input-ls-requests-lsa", types.YLeaf{"InputLsRequestsLsa", sample.InputLsRequestsLsa})
    sample.EntityData.Leafs.Append("output-ls-requests", types.YLeaf{"OutputLsRequests", sample.OutputLsRequests})
    sample.EntityData.Leafs.Append("output-ls-requests-lsa", types.YLeaf{"OutputLsRequestsLsa", sample.OutputLsRequestsLsa})
    sample.EntityData.Leafs.Append("input-lsa-updates", types.YLeaf{"InputLsaUpdates", sample.InputLsaUpdates})
    sample.EntityData.Leafs.Append("input-lsa-updates-lsa", types.YLeaf{"InputLsaUpdatesLsa", sample.InputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("output-lsa-updates", types.YLeaf{"OutputLsaUpdates", sample.OutputLsaUpdates})
    sample.EntityData.Leafs.Append("output-lsa-updates-lsa", types.YLeaf{"OutputLsaUpdatesLsa", sample.OutputLsaUpdatesLsa})
    sample.EntityData.Leafs.Append("input-lsa-acks", types.YLeaf{"InputLsaAcks", sample.InputLsaAcks})
    sample.EntityData.Leafs.Append("input-lsa-acks-lsa", types.YLeaf{"InputLsaAcksLsa", sample.InputLsaAcksLsa})
    sample.EntityData.Leafs.Append("output-lsa-acks", types.YLeaf{"OutputLsaAcks", sample.OutputLsaAcks})
    sample.EntityData.Leafs.Append("output-lsa-acks-lsa", types.YLeaf{"OutputLsaAcksLsa", sample.OutputLsaAcksLsa})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Mpls
// Collected MPLS data
type PerfMgmt_Monitor_Mpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP neighbors for which statistics are collected.
    LdpNeighbors PerfMgmt_Monitor_Mpls_LdpNeighbors
}

func (mpls *PerfMgmt_Monitor_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "cisco_ios_xr"
    mpls.EntityData.ParentYangName = "monitor"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpls.EntityData.Children = types.NewOrderedMap()
    mpls.EntityData.Children.Append("ldp-neighbors", types.YChild{"LdpNeighbors", &mpls.LdpNeighbors})
    mpls.EntityData.Leafs = types.NewOrderedMap()

    mpls.EntityData.YListKeys = []string {}

    return &(mpls.EntityData)
}

// PerfMgmt_Monitor_Mpls_LdpNeighbors
// LDP neighbors for which statistics are
// collected
type PerfMgmt_Monitor_Mpls_LdpNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular LDP neighbor. The type is slice of
    // PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor.
    LdpNeighbor []*PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor
}

func (ldpNeighbors *PerfMgmt_Monitor_Mpls_LdpNeighbors) GetEntityData() *types.CommonEntityData {
    ldpNeighbors.EntityData.YFilter = ldpNeighbors.YFilter
    ldpNeighbors.EntityData.YangName = "ldp-neighbors"
    ldpNeighbors.EntityData.BundleName = "cisco_ios_xr"
    ldpNeighbors.EntityData.ParentYangName = "mpls"
    ldpNeighbors.EntityData.SegmentPath = "ldp-neighbors"
    ldpNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpNeighbors.EntityData.Children = types.NewOrderedMap()
    ldpNeighbors.EntityData.Children.Append("ldp-neighbor", types.YChild{"LdpNeighbor", nil})
    for i := range ldpNeighbors.LdpNeighbor {
        ldpNeighbors.EntityData.Children.Append(types.GetSegmentPath(ldpNeighbors.LdpNeighbor[i]), types.YChild{"LdpNeighbor", ldpNeighbors.LdpNeighbor[i]})
    }
    ldpNeighbors.EntityData.Leafs = types.NewOrderedMap()

    ldpNeighbors.EntityData.YListKeys = []string {}

    return &(ldpNeighbors.EntityData)
}

// PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor
// Samples for a particular LDP neighbor
type PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Nbr interface{}

    // Samples for a particular LDP neighbor.
    Samples PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples
}

func (ldpNeighbor *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor) GetEntityData() *types.CommonEntityData {
    ldpNeighbor.EntityData.YFilter = ldpNeighbor.YFilter
    ldpNeighbor.EntityData.YangName = "ldp-neighbor"
    ldpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    ldpNeighbor.EntityData.ParentYangName = "ldp-neighbors"
    ldpNeighbor.EntityData.SegmentPath = "ldp-neighbor" + types.AddKeyToken(ldpNeighbor.Nbr, "nbr")
    ldpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpNeighbor.EntityData.Children = types.NewOrderedMap()
    ldpNeighbor.EntityData.Children.Append("samples", types.YChild{"Samples", &ldpNeighbor.Samples})
    ldpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    ldpNeighbor.EntityData.Leafs.Append("nbr", types.YLeaf{"Nbr", ldpNeighbor.Nbr})

    ldpNeighbor.EntityData.YListKeys = []string {"Nbr"}

    return &(ldpNeighbor.EntityData)
}

// PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples
// Samples for a particular LDP neighbor
type PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP neighbor statistics sample. The type is slice of
    // PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "ldp-neighbor"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample
// LDP neighbor statistics sample
type PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Mpls_LdpNeighbors_LdpNeighbor_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("total-msgs-sent", types.YLeaf{"TotalMsgsSent", sample.TotalMsgsSent})
    sample.EntityData.Leafs.Append("total-msgs-rcvd", types.YLeaf{"TotalMsgsRcvd", sample.TotalMsgsRcvd})
    sample.EntityData.Leafs.Append("init-msgs-sent", types.YLeaf{"InitMsgsSent", sample.InitMsgsSent})
    sample.EntityData.Leafs.Append("init-msgs-rcvd", types.YLeaf{"InitMsgsRcvd", sample.InitMsgsRcvd})
    sample.EntityData.Leafs.Append("address-msgs-sent", types.YLeaf{"AddressMsgsSent", sample.AddressMsgsSent})
    sample.EntityData.Leafs.Append("address-msgs-rcvd", types.YLeaf{"AddressMsgsRcvd", sample.AddressMsgsRcvd})
    sample.EntityData.Leafs.Append("address-withdraw-msgs-sent", types.YLeaf{"AddressWithdrawMsgsSent", sample.AddressWithdrawMsgsSent})
    sample.EntityData.Leafs.Append("address-withdraw-msgs-rcvd", types.YLeaf{"AddressWithdrawMsgsRcvd", sample.AddressWithdrawMsgsRcvd})
    sample.EntityData.Leafs.Append("label-mapping-msgs-sent", types.YLeaf{"LabelMappingMsgsSent", sample.LabelMappingMsgsSent})
    sample.EntityData.Leafs.Append("label-mapping-msgs-rcvd", types.YLeaf{"LabelMappingMsgsRcvd", sample.LabelMappingMsgsRcvd})
    sample.EntityData.Leafs.Append("label-withdraw-msgs-sent", types.YLeaf{"LabelWithdrawMsgsSent", sample.LabelWithdrawMsgsSent})
    sample.EntityData.Leafs.Append("label-withdraw-msgs-rcvd", types.YLeaf{"LabelWithdrawMsgsRcvd", sample.LabelWithdrawMsgsRcvd})
    sample.EntityData.Leafs.Append("label-release-msgs-sent", types.YLeaf{"LabelReleaseMsgsSent", sample.LabelReleaseMsgsSent})
    sample.EntityData.Leafs.Append("label-release-msgs-rcvd", types.YLeaf{"LabelReleaseMsgsRcvd", sample.LabelReleaseMsgsRcvd})
    sample.EntityData.Leafs.Append("notification-msgs-sent", types.YLeaf{"NotificationMsgsSent", sample.NotificationMsgsSent})
    sample.EntityData.Leafs.Append("notification-msgs-rcvd", types.YLeaf{"NotificationMsgsRcvd", sample.NotificationMsgsRcvd})
    sample.EntityData.Leafs.Append("keepalive-msgs-sent", types.YLeaf{"KeepaliveMsgsSent", sample.KeepaliveMsgsSent})
    sample.EntityData.Leafs.Append("keepalive-msgs-rcvd", types.YLeaf{"KeepaliveMsgsRcvd", sample.KeepaliveMsgsRcvd})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Nodes
// Nodes for which data is collected
type PerfMgmt_Monitor_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Instance. The type is slice of PerfMgmt_Monitor_Nodes_Node.
    Node []*PerfMgmt_Monitor_Nodes_Node
}

func (nodes *PerfMgmt_Monitor_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "monitor"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node
// Node Instance
type PerfMgmt_Monitor_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *PerfMgmt_Monitor_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("sample-xr", types.YChild{"SampleXr", &node.SampleXr})
    node.EntityData.Children.Append("processes", types.YChild{"Processes", &node.Processes})
    node.EntityData.Children.Append("samples", types.YChild{"Samples", &node.Samples})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})

    node.EntityData.YListKeys = []string {"NodeId"}

    return &(node.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_SampleXr
// Node CPU data
type PerfMgmt_Monitor_Nodes_Node_SampleXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node CPU data sample. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample.
    Sample []*PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample
}

func (sampleXr *PerfMgmt_Monitor_Nodes_Node_SampleXr) GetEntityData() *types.CommonEntityData {
    sampleXr.EntityData.YFilter = sampleXr.YFilter
    sampleXr.EntityData.YangName = "sample-xr"
    sampleXr.EntityData.BundleName = "cisco_ios_xr"
    sampleXr.EntityData.ParentYangName = "node"
    sampleXr.EntityData.SegmentPath = "sample-xr"
    sampleXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sampleXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sampleXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sampleXr.EntityData.Children = types.NewOrderedMap()
    sampleXr.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range sampleXr.Sample {
        sampleXr.EntityData.Children.Append(types.GetSegmentPath(sampleXr.Sample[i]), types.YChild{"Sample", sampleXr.Sample[i]})
    }
    sampleXr.EntityData.Leafs = types.NewOrderedMap()

    sampleXr.EntityData.YListKeys = []string {}

    return &(sampleXr.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample
// Node CPU data sample
type PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Nodes_Node_SampleXr_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "sample-xr"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("no-processes", types.YLeaf{"NoProcesses", sample.NoProcesses})
    sample.EntityData.Leafs.Append("average-cpu-used", types.YLeaf{"AverageCpuUsed", sample.AverageCpuUsed})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_Processes
// Processes data
type PerfMgmt_Monitor_Nodes_Node_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process data. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_Processes_Process.
    Process []*PerfMgmt_Monitor_Nodes_Node_Processes_Process
}

func (processes *PerfMgmt_Monitor_Nodes_Node_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "node"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range processes.Process {
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.Process[i]), types.YChild{"Process", processes.Process[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_Processes_Process
// Process data
type PerfMgmt_Monitor_Nodes_Node_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Process data.
    Samples PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples
}

func (process *PerfMgmt_Monitor_Nodes_Node_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessId, "process-id")
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("samples", types.YChild{"Samples", &process.Samples})
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", process.ProcessId})

    process.EntityData.YListKeys = []string {"ProcessId"}

    return &(process.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples
// Process data
type PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process data sample. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "process"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample
// Process data sample
type PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Nodes_Node_Processes_Process_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("peak-memory", types.YLeaf{"PeakMemory", sample.PeakMemory})
    sample.EntityData.Leafs.Append("average-cpu-used", types.YLeaf{"AverageCpuUsed", sample.AverageCpuUsed})
    sample.EntityData.Leafs.Append("no-threads", types.YLeaf{"NoThreads", sample.NoThreads})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_Samples
// Node Memory data
type PerfMgmt_Monitor_Nodes_Node_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory data sample. The type is slice of
    // PerfMgmt_Monitor_Nodes_Node_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Nodes_Node_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Nodes_Node_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "node"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Nodes_Node_Samples_Sample
// Node Memory data sample
type PerfMgmt_Monitor_Nodes_Node_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Nodes_Node_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("curr-memory", types.YLeaf{"CurrMemory", sample.CurrMemory})
    sample.EntityData.Leafs.Append("peak-memory", types.YLeaf{"PeakMemory", sample.PeakMemory})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Bgp
// Collected BGP data
type PerfMgmt_Monitor_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbors for which statistics are collected.
    BgpNeighbors PerfMgmt_Monitor_Bgp_BgpNeighbors
}

func (bgp *PerfMgmt_Monitor_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "monitor"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("bgp-neighbors", types.YChild{"BgpNeighbors", &bgp.BgpNeighbors})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PerfMgmt_Monitor_Bgp_BgpNeighbors
// Neighbors for which statistics are collected
type PerfMgmt_Monitor_Bgp_BgpNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for particular neighbor. The type is slice of
    // PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor.
    BgpNeighbor []*PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor
}

func (bgpNeighbors *PerfMgmt_Monitor_Bgp_BgpNeighbors) GetEntityData() *types.CommonEntityData {
    bgpNeighbors.EntityData.YFilter = bgpNeighbors.YFilter
    bgpNeighbors.EntityData.YangName = "bgp-neighbors"
    bgpNeighbors.EntityData.BundleName = "cisco_ios_xr"
    bgpNeighbors.EntityData.ParentYangName = "bgp"
    bgpNeighbors.EntityData.SegmentPath = "bgp-neighbors"
    bgpNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpNeighbors.EntityData.Children = types.NewOrderedMap()
    bgpNeighbors.EntityData.Children.Append("bgp-neighbor", types.YChild{"BgpNeighbor", nil})
    for i := range bgpNeighbors.BgpNeighbor {
        bgpNeighbors.EntityData.Children.Append(types.GetSegmentPath(bgpNeighbors.BgpNeighbor[i]), types.YChild{"BgpNeighbor", bgpNeighbors.BgpNeighbor[i]})
    }
    bgpNeighbors.EntityData.Leafs = types.NewOrderedMap()

    bgpNeighbors.EntityData.YListKeys = []string {}

    return &(bgpNeighbors.EntityData)
}

// PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor
// Samples for particular neighbor
type PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP Neighbor Identifier. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Sample Table for a BGP neighbor.
    Samples PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples
}

func (bgpNeighbor *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor) GetEntityData() *types.CommonEntityData {
    bgpNeighbor.EntityData.YFilter = bgpNeighbor.YFilter
    bgpNeighbor.EntityData.YangName = "bgp-neighbor"
    bgpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    bgpNeighbor.EntityData.ParentYangName = "bgp-neighbors"
    bgpNeighbor.EntityData.SegmentPath = "bgp-neighbor" + types.AddKeyToken(bgpNeighbor.IpAddress, "ip-address")
    bgpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpNeighbor.EntityData.Children = types.NewOrderedMap()
    bgpNeighbor.EntityData.Children.Append("samples", types.YChild{"Samples", &bgpNeighbor.Samples})
    bgpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    bgpNeighbor.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", bgpNeighbor.IpAddress})

    bgpNeighbor.EntityData.YListKeys = []string {"IpAddress"}

    return &(bgpNeighbor.EntityData)
}

// PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples
// Sample Table for a BGP neighbor
type PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor statistics sample. The type is slice of
    // PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "bgp-neighbor"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample
// Neighbor statistics sample
type PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Bgp_BgpNeighbors_BgpNeighbor_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-messages", types.YLeaf{"InputMessages", sample.InputMessages})
    sample.EntityData.Leafs.Append("output-messages", types.YLeaf{"OutputMessages", sample.OutputMessages})
    sample.EntityData.Leafs.Append("input-update-messages", types.YLeaf{"InputUpdateMessages", sample.InputUpdateMessages})
    sample.EntityData.Leafs.Append("output-update-messages", types.YLeaf{"OutputUpdateMessages", sample.OutputUpdateMessages})
    sample.EntityData.Leafs.Append("conn-established", types.YLeaf{"ConnEstablished", sample.ConnEstablished})
    sample.EntityData.Leafs.Append("conn-dropped", types.YLeaf{"ConnDropped", sample.ConnDropped})
    sample.EntityData.Leafs.Append("errors-received", types.YLeaf{"ErrorsReceived", sample.ErrorsReceived})
    sample.EntityData.Leafs.Append("errors-sent", types.YLeaf{"ErrorsSent", sample.ErrorsSent})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Interface
// Collected Interface data
type PerfMgmt_Monitor_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces for which Generic Counters are collected.
    GenericCounterInterfaces PerfMgmt_Monitor_Interface_GenericCounterInterfaces

    // Interfaces for which Basic Counters are collected.
    BasicCounterInterfaces PerfMgmt_Monitor_Interface_BasicCounterInterfaces

    // Interfaces for which Data Rates are collected.
    DataRateInterfaces PerfMgmt_Monitor_Interface_DataRateInterfaces
}

func (self *PerfMgmt_Monitor_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "monitor"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("generic-counter-interfaces", types.YChild{"GenericCounterInterfaces", &self.GenericCounterInterfaces})
    self.EntityData.Children.Append("basic-counter-interfaces", types.YChild{"BasicCounterInterfaces", &self.BasicCounterInterfaces})
    self.EntityData.Children.Append("data-rate-interfaces", types.YChild{"DataRateInterfaces", &self.DataRateInterfaces})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces
// Interfaces for which Generic Counters are
// collected
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface.
    GenericCounterInterface []*PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface
}

func (genericCounterInterfaces *PerfMgmt_Monitor_Interface_GenericCounterInterfaces) GetEntityData() *types.CommonEntityData {
    genericCounterInterfaces.EntityData.YFilter = genericCounterInterfaces.YFilter
    genericCounterInterfaces.EntityData.YangName = "generic-counter-interfaces"
    genericCounterInterfaces.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterfaces.EntityData.ParentYangName = "interface"
    genericCounterInterfaces.EntityData.SegmentPath = "generic-counter-interfaces"
    genericCounterInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterfaces.EntityData.Children = types.NewOrderedMap()
    genericCounterInterfaces.EntityData.Children.Append("generic-counter-interface", types.YChild{"GenericCounterInterface", nil})
    for i := range genericCounterInterfaces.GenericCounterInterface {
        genericCounterInterfaces.EntityData.Children.Append(types.GetSegmentPath(genericCounterInterfaces.GenericCounterInterface[i]), types.YChild{"GenericCounterInterface", genericCounterInterfaces.GenericCounterInterface[i]})
    }
    genericCounterInterfaces.EntityData.Leafs = types.NewOrderedMap()

    genericCounterInterfaces.EntityData.YListKeys = []string {}

    return &(genericCounterInterfaces.EntityData)
}

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface
// Samples for a particular interface
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Generic Counter samples for an interface.
    Samples PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
}

func (genericCounterInterface *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface) GetEntityData() *types.CommonEntityData {
    genericCounterInterface.EntityData.YFilter = genericCounterInterface.YFilter
    genericCounterInterface.EntityData.YangName = "generic-counter-interface"
    genericCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterface.EntityData.ParentYangName = "generic-counter-interfaces"
    genericCounterInterface.EntityData.SegmentPath = "generic-counter-interface" + types.AddKeyToken(genericCounterInterface.InterfaceName, "interface-name")
    genericCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterface.EntityData.Children = types.NewOrderedMap()
    genericCounterInterface.EntityData.Children.Append("samples", types.YChild{"Samples", &genericCounterInterface.Samples})
    genericCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    genericCounterInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", genericCounterInterface.InterfaceName})

    genericCounterInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples
// Generic Counter samples for an interface
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "generic-counter-interface"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample
// Generic Counters sample
type PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Interface_GenericCounterInterfaces_GenericCounterInterface_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("in-packets", types.YLeaf{"InPackets", sample.InPackets})
    sample.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", sample.InOctets})
    sample.EntityData.Leafs.Append("out-packets", types.YLeaf{"OutPackets", sample.OutPackets})
    sample.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", sample.OutOctets})
    sample.EntityData.Leafs.Append("in-ucast-pkts", types.YLeaf{"InUcastPkts", sample.InUcastPkts})
    sample.EntityData.Leafs.Append("in-multicast-pkts", types.YLeaf{"InMulticastPkts", sample.InMulticastPkts})
    sample.EntityData.Leafs.Append("in-broadcast-pkts", types.YLeaf{"InBroadcastPkts", sample.InBroadcastPkts})
    sample.EntityData.Leafs.Append("out-ucast-pkts", types.YLeaf{"OutUcastPkts", sample.OutUcastPkts})
    sample.EntityData.Leafs.Append("out-multicast-pkts", types.YLeaf{"OutMulticastPkts", sample.OutMulticastPkts})
    sample.EntityData.Leafs.Append("out-broadcast-pkts", types.YLeaf{"OutBroadcastPkts", sample.OutBroadcastPkts})
    sample.EntityData.Leafs.Append("output-total-drops", types.YLeaf{"OutputTotalDrops", sample.OutputTotalDrops})
    sample.EntityData.Leafs.Append("input-total-drops", types.YLeaf{"InputTotalDrops", sample.InputTotalDrops})
    sample.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", sample.InputQueueDrops})
    sample.EntityData.Leafs.Append("input-unknown-proto", types.YLeaf{"InputUnknownProto", sample.InputUnknownProto})
    sample.EntityData.Leafs.Append("output-total-errors", types.YLeaf{"OutputTotalErrors", sample.OutputTotalErrors})
    sample.EntityData.Leafs.Append("output-underrun", types.YLeaf{"OutputUnderrun", sample.OutputUnderrun})
    sample.EntityData.Leafs.Append("input-total-errors", types.YLeaf{"InputTotalErrors", sample.InputTotalErrors})
    sample.EntityData.Leafs.Append("input-crc", types.YLeaf{"InputCrc", sample.InputCrc})
    sample.EntityData.Leafs.Append("input-overrun", types.YLeaf{"InputOverrun", sample.InputOverrun})
    sample.EntityData.Leafs.Append("input-frame", types.YLeaf{"InputFrame", sample.InputFrame})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces
// Interfaces for which Basic Counters are
// collected
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface.
    BasicCounterInterface []*PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface
}

func (basicCounterInterfaces *PerfMgmt_Monitor_Interface_BasicCounterInterfaces) GetEntityData() *types.CommonEntityData {
    basicCounterInterfaces.EntityData.YFilter = basicCounterInterfaces.YFilter
    basicCounterInterfaces.EntityData.YangName = "basic-counter-interfaces"
    basicCounterInterfaces.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterfaces.EntityData.ParentYangName = "interface"
    basicCounterInterfaces.EntityData.SegmentPath = "basic-counter-interfaces"
    basicCounterInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterfaces.EntityData.Children = types.NewOrderedMap()
    basicCounterInterfaces.EntityData.Children.Append("basic-counter-interface", types.YChild{"BasicCounterInterface", nil})
    for i := range basicCounterInterfaces.BasicCounterInterface {
        basicCounterInterfaces.EntityData.Children.Append(types.GetSegmentPath(basicCounterInterfaces.BasicCounterInterface[i]), types.YChild{"BasicCounterInterface", basicCounterInterfaces.BasicCounterInterface[i]})
    }
    basicCounterInterfaces.EntityData.Leafs = types.NewOrderedMap()

    basicCounterInterfaces.EntityData.YListKeys = []string {}

    return &(basicCounterInterfaces.EntityData)
}

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface
// Samples for a particular interface
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Basic Counter samples for an interface.
    Samples PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
}

func (basicCounterInterface *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface) GetEntityData() *types.CommonEntityData {
    basicCounterInterface.EntityData.YFilter = basicCounterInterface.YFilter
    basicCounterInterface.EntityData.YangName = "basic-counter-interface"
    basicCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterface.EntityData.ParentYangName = "basic-counter-interfaces"
    basicCounterInterface.EntityData.SegmentPath = "basic-counter-interface" + types.AddKeyToken(basicCounterInterface.InterfaceName, "interface-name")
    basicCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterface.EntityData.Children = types.NewOrderedMap()
    basicCounterInterface.EntityData.Children.Append("samples", types.YChild{"Samples", &basicCounterInterface.Samples})
    basicCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    basicCounterInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", basicCounterInterface.InterfaceName})

    basicCounterInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples
// Basic Counter samples for an interface
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Basic Counters sample. The type is slice of
    // PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "basic-counter-interface"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample
// Basic Counters sample
type PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Interface_BasicCounterInterfaces_BasicCounterInterface_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("in-packets", types.YLeaf{"InPackets", sample.InPackets})
    sample.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", sample.InOctets})
    sample.EntityData.Leafs.Append("out-packets", types.YLeaf{"OutPackets", sample.OutPackets})
    sample.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", sample.OutOctets})
    sample.EntityData.Leafs.Append("input-total-drops", types.YLeaf{"InputTotalDrops", sample.InputTotalDrops})
    sample.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", sample.InputQueueDrops})
    sample.EntityData.Leafs.Append("input-total-errors", types.YLeaf{"InputTotalErrors", sample.InputTotalErrors})
    sample.EntityData.Leafs.Append("output-total-drops", types.YLeaf{"OutputTotalDrops", sample.OutputTotalDrops})
    sample.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", sample.OutputQueueDrops})
    sample.EntityData.Leafs.Append("output-total-errors", types.YLeaf{"OutputTotalErrors", sample.OutputTotalErrors})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

// PerfMgmt_Monitor_Interface_DataRateInterfaces
// Interfaces for which Data Rates are collected
type PerfMgmt_Monitor_Interface_DataRateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Samples for a particular interface. The type is slice of
    // PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface.
    DataRateInterface []*PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface
}

func (dataRateInterfaces *PerfMgmt_Monitor_Interface_DataRateInterfaces) GetEntityData() *types.CommonEntityData {
    dataRateInterfaces.EntityData.YFilter = dataRateInterfaces.YFilter
    dataRateInterfaces.EntityData.YangName = "data-rate-interfaces"
    dataRateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterfaces.EntityData.ParentYangName = "interface"
    dataRateInterfaces.EntityData.SegmentPath = "data-rate-interfaces"
    dataRateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterfaces.EntityData.Children = types.NewOrderedMap()
    dataRateInterfaces.EntityData.Children.Append("data-rate-interface", types.YChild{"DataRateInterface", nil})
    for i := range dataRateInterfaces.DataRateInterface {
        dataRateInterfaces.EntityData.Children.Append(types.GetSegmentPath(dataRateInterfaces.DataRateInterface[i]), types.YChild{"DataRateInterface", dataRateInterfaces.DataRateInterface[i]})
    }
    dataRateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    dataRateInterfaces.EntityData.YListKeys = []string {}

    return &(dataRateInterfaces.EntityData)
}

// PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface
// Samples for a particular interface
type PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Data Rate samples for an interface.
    Samples PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples
}

func (dataRateInterface *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface) GetEntityData() *types.CommonEntityData {
    dataRateInterface.EntityData.YFilter = dataRateInterface.YFilter
    dataRateInterface.EntityData.YangName = "data-rate-interface"
    dataRateInterface.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterface.EntityData.ParentYangName = "data-rate-interfaces"
    dataRateInterface.EntityData.SegmentPath = "data-rate-interface" + types.AddKeyToken(dataRateInterface.InterfaceName, "interface-name")
    dataRateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterface.EntityData.Children = types.NewOrderedMap()
    dataRateInterface.EntityData.Children.Append("samples", types.YChild{"Samples", &dataRateInterface.Samples})
    dataRateInterface.EntityData.Leafs = types.NewOrderedMap()
    dataRateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", dataRateInterface.InterfaceName})

    dataRateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples
// Data Rate samples for an interface
type PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data Rates sample. The type is slice of
    // PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample.
    Sample []*PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
}

func (samples *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "data-rate-interface"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = types.NewOrderedMap()
    samples.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range samples.Sample {
        samples.EntityData.Children.Append(types.GetSegmentPath(samples.Sample[i]), types.YChild{"Sample", samples.Sample[i]})
    }
    samples.EntityData.Leafs = types.NewOrderedMap()

    samples.EntityData.YListKeys = []string {}

    return &(samples.EntityData)
}

// PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample
// Data Rates sample
type PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample struct {
    EntityData types.CommonEntityData
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

func (sample *PerfMgmt_Monitor_Interface_DataRateInterfaces_DataRateInterface_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + types.AddKeyToken(sample.SampleId, "sample-id")
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sample-id", types.YLeaf{"SampleId", sample.SampleId})
    sample.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", sample.TimeStamp})
    sample.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", sample.InputDataRate})
    sample.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", sample.InputPacketRate})
    sample.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", sample.OutputDataRate})
    sample.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", sample.OutputPacketRate})
    sample.EntityData.Leafs.Append("input-peak-rate", types.YLeaf{"InputPeakRate", sample.InputPeakRate})
    sample.EntityData.Leafs.Append("input-peak-pkts", types.YLeaf{"InputPeakPkts", sample.InputPeakPkts})
    sample.EntityData.Leafs.Append("output-peak-rate", types.YLeaf{"OutputPeakRate", sample.OutputPeakRate})
    sample.EntityData.Leafs.Append("output-peak-pkts", types.YLeaf{"OutputPeakPkts", sample.OutputPeakPkts})
    sample.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", sample.Bandwidth})

    sample.EntityData.YListKeys = []string {"SampleId"}

    return &(sample.EntityData)
}

