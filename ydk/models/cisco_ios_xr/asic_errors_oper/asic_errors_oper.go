// This module contains a collection of YANG definitions
// for Cisco IOS-XR asic-errors package operational data.
// 
// This module contains definitions
// for the following management objects:
//   asic-errors: Error summary of all asics
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asic_errors_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asic_errors_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asic-errors-oper asic-errors}", reflect.TypeOf(AsicErrors{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asic-errors-oper:asic-errors", reflect.TypeOf(AsicErrors{}))
}

// AsicErrors
// Error summary of all asics
type AsicErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic errors for each available nodes.
    Nodes AsicErrors_Nodes
}

func (asicErrors *AsicErrors) GetEntityData() *types.CommonEntityData {
    asicErrors.EntityData.YFilter = asicErrors.YFilter
    asicErrors.EntityData.YangName = "asic-errors"
    asicErrors.EntityData.BundleName = "cisco_ios_xr"
    asicErrors.EntityData.ParentYangName = "Cisco-IOS-XR-asic-errors-oper"
    asicErrors.EntityData.SegmentPath = "Cisco-IOS-XR-asic-errors-oper:asic-errors"
    asicErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrors.EntityData.Children = make(map[string]types.YChild)
    asicErrors.EntityData.Children["nodes"] = types.YChild{"Nodes", &asicErrors.Nodes}
    asicErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrors.EntityData)
}

// AsicErrors_Nodes
// Asic errors for each available nodes
type AsicErrors_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic error for a particular node. The type is slice of
    // AsicErrors_Nodes_Node.
    Node []AsicErrors_Nodes_Node
}

func (nodes *AsicErrors_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "asic-errors"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// AsicErrors_Nodes_Node
// Asic error for a particular node
type AsicErrors_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Asic on the node. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation.
    AsicInformation []AsicErrors_Nodes_Node_AsicInformation
}

func (node *AsicErrors_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["asic-information"] = types.YChild{"AsicInformation", nil}
    for i := range node.AsicInformation {
        node.EntityData.Children[types.GetSegmentPath(&node.AsicInformation[i])] = types.YChild{"AsicInformation", &node.AsicInformation[i]}
    }
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation
// Asic on the node
type AsicErrors_Nodes_Node_AsicInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Asic string. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Asic interface{}

    // All asic instance on the node.
    AllInstances AsicErrors_Nodes_Node_AsicInformation_AllInstances

    // All asic errors  on the node.
    Instances AsicErrors_Nodes_Node_AsicInformation_Instances
}

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetEntityData() *types.CommonEntityData {
    asicInformation.EntityData.YFilter = asicInformation.YFilter
    asicInformation.EntityData.YangName = "asic-information"
    asicInformation.EntityData.BundleName = "cisco_ios_xr"
    asicInformation.EntityData.ParentYangName = "node"
    asicInformation.EntityData.SegmentPath = "asic-information" + "[asic='" + fmt.Sprintf("%v", asicInformation.Asic) + "']"
    asicInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicInformation.EntityData.Children = make(map[string]types.YChild)
    asicInformation.EntityData.Children["all-instances"] = types.YChild{"AllInstances", &asicInformation.AllInstances}
    asicInformation.EntityData.Children["instances"] = types.YChild{"Instances", &asicInformation.Instances}
    asicInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    asicInformation.EntityData.Leafs["asic"] = types.YLeaf{"Asic", asicInformation.Asic}
    return &(asicInformation.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_AllInstances
// All asic instance on the node
type AsicErrors_Nodes_Node_AsicInformation_AllInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error path of all instances.
    AllErrorPath AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath
}

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetEntityData() *types.CommonEntityData {
    allInstances.EntityData.YFilter = allInstances.YFilter
    allInstances.EntityData.YangName = "all-instances"
    allInstances.EntityData.BundleName = "cisco_ios_xr"
    allInstances.EntityData.ParentYangName = "asic-information"
    allInstances.EntityData.SegmentPath = "all-instances"
    allInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allInstances.EntityData.Children = make(map[string]types.YChild)
    allInstances.EntityData.Children["all-error-path"] = types.YChild{"AllErrorPath", &allInstances.AllErrorPath}
    allInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(allInstances.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath
// Error path of all instances
type AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of all instances errors.
    Summary AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary
}

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetEntityData() *types.CommonEntityData {
    allErrorPath.EntityData.YFilter = allErrorPath.YFilter
    allErrorPath.EntityData.YangName = "all-error-path"
    allErrorPath.EntityData.BundleName = "cisco_ios_xr"
    allErrorPath.EntityData.ParentYangName = "all-instances"
    allErrorPath.EntityData.SegmentPath = "all-error-path"
    allErrorPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allErrorPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allErrorPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allErrorPath.EntityData.Children = make(map[string]types.YChild)
    allErrorPath.EntityData.Children["summary"] = types.YChild{"Summary", &allErrorPath.Summary}
    allErrorPath.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(allErrorPath.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary
// Summary of all instances errors
type AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // legacy client. The type is bool.
    LegacyClient interface{}

    // cih client. The type is bool.
    CihClient interface{}

    // sum data. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData.
    SumData []AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData
}

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "all-error-path"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["sum-data"] = types.YChild{"SumData", nil}
    for i := range summary.SumData {
        summary.EntityData.Children[types.GetSegmentPath(&summary.SumData[i])] = types.YChild{"SumData", &summary.SumData[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["legacy-client"] = types.YLeaf{"LegacyClient", summary.LegacyClient}
    summary.EntityData.Leafs["cih-client"] = types.YLeaf{"CihClient", summary.CihClient}
    return &(summary.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData
// sum data
type AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // num nodes. The type is interface{} with range: 0..4294967295.
    NumNodes interface{}

    // crc err count. The type is interface{} with range: 0..4294967295.
    CrcErrCount interface{}

    // sbe err count. The type is interface{} with range: 0..4294967295.
    SbeErrCount interface{}

    // mbe err count. The type is interface{} with range: 0..4294967295.
    MbeErrCount interface{}

    // par err count. The type is interface{} with range: 0..4294967295.
    ParErrCount interface{}

    // gen err count. The type is interface{} with range: 0..4294967295.
    GenErrCount interface{}

    // reset err count. The type is interface{} with range: 0..4294967295.
    ResetErrCount interface{}

    // err count. The type is slice of interface{} with range: 0..4294967295.
    ErrCount []interface{}

    // pcie err count. The type is slice of interface{} with range: 0..4294967295.
    PcieErrCount []interface{}

    // node key. The type is slice of interface{} with range: 0..4294967295.
    NodeKey []interface{}
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetEntityData() *types.CommonEntityData {
    sumData.EntityData.YFilter = sumData.YFilter
    sumData.EntityData.YangName = "sum-data"
    sumData.EntityData.BundleName = "cisco_ios_xr"
    sumData.EntityData.ParentYangName = "summary"
    sumData.EntityData.SegmentPath = "sum-data"
    sumData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumData.EntityData.Children = make(map[string]types.YChild)
    sumData.EntityData.Leafs = make(map[string]types.YLeaf)
    sumData.EntityData.Leafs["num-nodes"] = types.YLeaf{"NumNodes", sumData.NumNodes}
    sumData.EntityData.Leafs["crc-err-count"] = types.YLeaf{"CrcErrCount", sumData.CrcErrCount}
    sumData.EntityData.Leafs["sbe-err-count"] = types.YLeaf{"SbeErrCount", sumData.SbeErrCount}
    sumData.EntityData.Leafs["mbe-err-count"] = types.YLeaf{"MbeErrCount", sumData.MbeErrCount}
    sumData.EntityData.Leafs["par-err-count"] = types.YLeaf{"ParErrCount", sumData.ParErrCount}
    sumData.EntityData.Leafs["gen-err-count"] = types.YLeaf{"GenErrCount", sumData.GenErrCount}
    sumData.EntityData.Leafs["reset-err-count"] = types.YLeaf{"ResetErrCount", sumData.ResetErrCount}
    sumData.EntityData.Leafs["err-count"] = types.YLeaf{"ErrCount", sumData.ErrCount}
    sumData.EntityData.Leafs["pcie-err-count"] = types.YLeaf{"PcieErrCount", sumData.PcieErrCount}
    sumData.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", sumData.NodeKey}
    return &(sumData.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances
// All asic errors  on the node
type AsicErrors_Nodes_Node_AsicInformation_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Particular asic instance on the node. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance.
    Instance []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance
}

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "asic-information"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = make(map[string]types.YChild)
    instances.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range instances.Instance {
        instances.EntityData.Children[types.GetSegmentPath(&instances.Instance[i])] = types.YChild{"Instance", &instances.Instance[i]}
    }
    instances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(instances.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance
// Particular asic instance on the node
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. asic instance. The type is interface{} with range:
    // -2147483648..2147483647.
    AsicInstance interface{}

    // Error path of the instances.
    ErrorPath AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath
}

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + "[asic-instance='" + fmt.Sprintf("%v", instance.AsicInstance) + "']"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["error-path"] = types.YChild{"ErrorPath", &instance.ErrorPath}
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", instance.AsicInstance}
    return &(instance.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath
// Error path of the instances
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multiple bit soft error information.
    MultipleBitSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors

    // Indirect hard error information.
    AsicErrorGenericSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft

    // CRC hard error information.
    CrcHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors

    // Indirect hard error information.
    AsicErrorSbeSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft

    // Hardware soft error information.
    HardwareSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors

    // Indirect hard error information.
    AsicErrorCrcSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft

    // Indirect hard error information.
    AsicErrorParitySoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft

    // IO soft error information.
    IoSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors

    // Reset soft error information.
    ResetSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors

    // Barrier hard error information.
    BarrierHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors

    // Ucode soft error information.
    UcodeSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors

    // Indirect hard error information.
    AsicErrorResetHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard

    // Single bit hard error information.
    SingleBitHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors

    // Indirect hard error information.
    IndirectHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors

    // OOR thresh information.
    OutofResourceSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft

    // CRC soft error information.
    CrcSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors

    // Time out hard error information.
    TimeOutHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors

    // Barrier soft error information.
    BarrierSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors

    // Indirect hard error information.
    AsicErrorMbeSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft

    // BP hard error information.
    BackPressureHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors

    // Single bit soft error information.
    SingleBitSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors

    // Indirect soft error information.
    IndirectSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors

    // Generic hard error information.
    GenericHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors

    // Link hard error information.
    LinkHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors

    // Configuration hard error information.
    ConfigurationHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors

    // Summary for a specific instance.
    InstanceSummary AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary

    // Unexpected hard error information.
    UnexpectedHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors

    // Time out soft error information.
    TimeOutSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors

    // Indirect hard error information.
    AsicErrorGenericHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard

    // Parity hard error information.
    ParityHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors

    // Descriptor hard error information.
    DescriptorHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors

    // Interface hard error information.
    InterfaceHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors

    // Indirect hard error information.
    AsicErrorSbeHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard

    // Indirect hard error information.
    AsicErrorCrcHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard

    // Indirect hard error information.
    AsicErrorParityHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard

    // Indirect hard error information.
    AsicErrorResetSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft

    // BP soft error information.
    BackPressureSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors

    // Generic soft error information.
    GenericSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors

    // Link soft error information.
    LinkSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors

    // Configuration soft error information.
    ConfigurationSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors

    // Multiple bit hard error information.
    MultipleBitHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors

    // Unexpected soft error information.
    UnexpectedSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors

    // OOR thresh information.
    OutofResourceHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard

    // Hardware hard error information.
    HardwareHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors

    // Parity soft error information.
    ParitySoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors

    // Descriptor soft error information.
    DescriptorSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors

    // Interface soft error information.
    InterfaceSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors

    // IO hard error information.
    IoHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors

    // Reset hard error information.
    ResetHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors

    // UCode hard error information.
    UcodeHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors

    // Indirect hard error information.
    AsicErrorMbeHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard
}

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetEntityData() *types.CommonEntityData {
    errorPath.EntityData.YFilter = errorPath.YFilter
    errorPath.EntityData.YangName = "error-path"
    errorPath.EntityData.BundleName = "cisco_ios_xr"
    errorPath.EntityData.ParentYangName = "instance"
    errorPath.EntityData.SegmentPath = "error-path"
    errorPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorPath.EntityData.Children = make(map[string]types.YChild)
    errorPath.EntityData.Children["multiple-bit-soft-errors"] = types.YChild{"MultipleBitSoftErrors", &errorPath.MultipleBitSoftErrors}
    errorPath.EntityData.Children["asic-error-generic-soft"] = types.YChild{"AsicErrorGenericSoft", &errorPath.AsicErrorGenericSoft}
    errorPath.EntityData.Children["crc-hard-errors"] = types.YChild{"CrcHardErrors", &errorPath.CrcHardErrors}
    errorPath.EntityData.Children["asic-error-sbe-soft"] = types.YChild{"AsicErrorSbeSoft", &errorPath.AsicErrorSbeSoft}
    errorPath.EntityData.Children["hardware-soft-errors"] = types.YChild{"HardwareSoftErrors", &errorPath.HardwareSoftErrors}
    errorPath.EntityData.Children["asic-error-crc-soft"] = types.YChild{"AsicErrorCrcSoft", &errorPath.AsicErrorCrcSoft}
    errorPath.EntityData.Children["asic-error-parity-soft"] = types.YChild{"AsicErrorParitySoft", &errorPath.AsicErrorParitySoft}
    errorPath.EntityData.Children["io-soft-errors"] = types.YChild{"IoSoftErrors", &errorPath.IoSoftErrors}
    errorPath.EntityData.Children["reset-soft-errors"] = types.YChild{"ResetSoftErrors", &errorPath.ResetSoftErrors}
    errorPath.EntityData.Children["barrier-hard-errors"] = types.YChild{"BarrierHardErrors", &errorPath.BarrierHardErrors}
    errorPath.EntityData.Children["ucode-soft-errors"] = types.YChild{"UcodeSoftErrors", &errorPath.UcodeSoftErrors}
    errorPath.EntityData.Children["asic-error-reset-hard"] = types.YChild{"AsicErrorResetHard", &errorPath.AsicErrorResetHard}
    errorPath.EntityData.Children["single-bit-hard-errors"] = types.YChild{"SingleBitHardErrors", &errorPath.SingleBitHardErrors}
    errorPath.EntityData.Children["indirect-hard-errors"] = types.YChild{"IndirectHardErrors", &errorPath.IndirectHardErrors}
    errorPath.EntityData.Children["outof-resource-soft"] = types.YChild{"OutofResourceSoft", &errorPath.OutofResourceSoft}
    errorPath.EntityData.Children["crc-soft-errors"] = types.YChild{"CrcSoftErrors", &errorPath.CrcSoftErrors}
    errorPath.EntityData.Children["time-out-hard-errors"] = types.YChild{"TimeOutHardErrors", &errorPath.TimeOutHardErrors}
    errorPath.EntityData.Children["barrier-soft-errors"] = types.YChild{"BarrierSoftErrors", &errorPath.BarrierSoftErrors}
    errorPath.EntityData.Children["asic-error-mbe-soft"] = types.YChild{"AsicErrorMbeSoft", &errorPath.AsicErrorMbeSoft}
    errorPath.EntityData.Children["back-pressure-hard-errors"] = types.YChild{"BackPressureHardErrors", &errorPath.BackPressureHardErrors}
    errorPath.EntityData.Children["single-bit-soft-errors"] = types.YChild{"SingleBitSoftErrors", &errorPath.SingleBitSoftErrors}
    errorPath.EntityData.Children["indirect-soft-errors"] = types.YChild{"IndirectSoftErrors", &errorPath.IndirectSoftErrors}
    errorPath.EntityData.Children["generic-hard-errors"] = types.YChild{"GenericHardErrors", &errorPath.GenericHardErrors}
    errorPath.EntityData.Children["link-hard-errors"] = types.YChild{"LinkHardErrors", &errorPath.LinkHardErrors}
    errorPath.EntityData.Children["configuration-hard-errors"] = types.YChild{"ConfigurationHardErrors", &errorPath.ConfigurationHardErrors}
    errorPath.EntityData.Children["instance-summary"] = types.YChild{"InstanceSummary", &errorPath.InstanceSummary}
    errorPath.EntityData.Children["unexpected-hard-errors"] = types.YChild{"UnexpectedHardErrors", &errorPath.UnexpectedHardErrors}
    errorPath.EntityData.Children["time-out-soft-errors"] = types.YChild{"TimeOutSoftErrors", &errorPath.TimeOutSoftErrors}
    errorPath.EntityData.Children["asic-error-generic-hard"] = types.YChild{"AsicErrorGenericHard", &errorPath.AsicErrorGenericHard}
    errorPath.EntityData.Children["parity-hard-errors"] = types.YChild{"ParityHardErrors", &errorPath.ParityHardErrors}
    errorPath.EntityData.Children["descriptor-hard-errors"] = types.YChild{"DescriptorHardErrors", &errorPath.DescriptorHardErrors}
    errorPath.EntityData.Children["interface-hard-errors"] = types.YChild{"InterfaceHardErrors", &errorPath.InterfaceHardErrors}
    errorPath.EntityData.Children["asic-error-sbe-hard"] = types.YChild{"AsicErrorSbeHard", &errorPath.AsicErrorSbeHard}
    errorPath.EntityData.Children["asic-error-crc-hard"] = types.YChild{"AsicErrorCrcHard", &errorPath.AsicErrorCrcHard}
    errorPath.EntityData.Children["asic-error-parity-hard"] = types.YChild{"AsicErrorParityHard", &errorPath.AsicErrorParityHard}
    errorPath.EntityData.Children["asic-error-reset-soft"] = types.YChild{"AsicErrorResetSoft", &errorPath.AsicErrorResetSoft}
    errorPath.EntityData.Children["back-pressure-soft-errors"] = types.YChild{"BackPressureSoftErrors", &errorPath.BackPressureSoftErrors}
    errorPath.EntityData.Children["generic-soft-errors"] = types.YChild{"GenericSoftErrors", &errorPath.GenericSoftErrors}
    errorPath.EntityData.Children["link-soft-errors"] = types.YChild{"LinkSoftErrors", &errorPath.LinkSoftErrors}
    errorPath.EntityData.Children["configuration-soft-errors"] = types.YChild{"ConfigurationSoftErrors", &errorPath.ConfigurationSoftErrors}
    errorPath.EntityData.Children["multiple-bit-hard-errors"] = types.YChild{"MultipleBitHardErrors", &errorPath.MultipleBitHardErrors}
    errorPath.EntityData.Children["unexpected-soft-errors"] = types.YChild{"UnexpectedSoftErrors", &errorPath.UnexpectedSoftErrors}
    errorPath.EntityData.Children["outof-resource-hard"] = types.YChild{"OutofResourceHard", &errorPath.OutofResourceHard}
    errorPath.EntityData.Children["hardware-hard-errors"] = types.YChild{"HardwareHardErrors", &errorPath.HardwareHardErrors}
    errorPath.EntityData.Children["parity-soft-errors"] = types.YChild{"ParitySoftErrors", &errorPath.ParitySoftErrors}
    errorPath.EntityData.Children["descriptor-soft-errors"] = types.YChild{"DescriptorSoftErrors", &errorPath.DescriptorSoftErrors}
    errorPath.EntityData.Children["interface-soft-errors"] = types.YChild{"InterfaceSoftErrors", &errorPath.InterfaceSoftErrors}
    errorPath.EntityData.Children["io-hard-errors"] = types.YChild{"IoHardErrors", &errorPath.IoHardErrors}
    errorPath.EntityData.Children["reset-hard-errors"] = types.YChild{"ResetHardErrors", &errorPath.ResetHardErrors}
    errorPath.EntityData.Children["ucode-hard-errors"] = types.YChild{"UcodeHardErrors", &errorPath.UcodeHardErrors}
    errorPath.EntityData.Children["asic-error-mbe-hard"] = types.YChild{"AsicErrorMbeHard", &errorPath.AsicErrorMbeHard}
    errorPath.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(errorPath.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors
// Multiple bit soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error
}

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetEntityData() *types.CommonEntityData {
    multipleBitSoftErrors.EntityData.YFilter = multipleBitSoftErrors.YFilter
    multipleBitSoftErrors.EntityData.YangName = "multiple-bit-soft-errors"
    multipleBitSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    multipleBitSoftErrors.EntityData.ParentYangName = "error-path"
    multipleBitSoftErrors.EntityData.SegmentPath = "multiple-bit-soft-errors"
    multipleBitSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multipleBitSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multipleBitSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multipleBitSoftErrors.EntityData.Children = make(map[string]types.YChild)
    multipleBitSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range multipleBitSoftErrors.Error {
        multipleBitSoftErrors.EntityData.Children[types.GetSegmentPath(&multipleBitSoftErrors.Error[i])] = types.YChild{"Error", &multipleBitSoftErrors.Error[i]}
    }
    multipleBitSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multipleBitSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "multiple-bit-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error
}

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetEntityData() *types.CommonEntityData {
    asicErrorGenericSoft.EntityData.YFilter = asicErrorGenericSoft.YFilter
    asicErrorGenericSoft.EntityData.YangName = "asic-error-generic-soft"
    asicErrorGenericSoft.EntityData.BundleName = "cisco_ios_xr"
    asicErrorGenericSoft.EntityData.ParentYangName = "error-path"
    asicErrorGenericSoft.EntityData.SegmentPath = "asic-error-generic-soft"
    asicErrorGenericSoft.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorGenericSoft.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorGenericSoft.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorGenericSoft.EntityData.Children = make(map[string]types.YChild)
    asicErrorGenericSoft.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorGenericSoft.Error {
        asicErrorGenericSoft.EntityData.Children[types.GetSegmentPath(&asicErrorGenericSoft.Error[i])] = types.YChild{"Error", &asicErrorGenericSoft.Error[i]}
    }
    asicErrorGenericSoft.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorGenericSoft.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-generic-soft"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors
// CRC hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error
}

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetEntityData() *types.CommonEntityData {
    crcHardErrors.EntityData.YFilter = crcHardErrors.YFilter
    crcHardErrors.EntityData.YangName = "crc-hard-errors"
    crcHardErrors.EntityData.BundleName = "cisco_ios_xr"
    crcHardErrors.EntityData.ParentYangName = "error-path"
    crcHardErrors.EntityData.SegmentPath = "crc-hard-errors"
    crcHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crcHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crcHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crcHardErrors.EntityData.Children = make(map[string]types.YChild)
    crcHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range crcHardErrors.Error {
        crcHardErrors.EntityData.Children[types.GetSegmentPath(&crcHardErrors.Error[i])] = types.YChild{"Error", &crcHardErrors.Error[i]}
    }
    crcHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(crcHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "crc-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error
}

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetEntityData() *types.CommonEntityData {
    asicErrorSbeSoft.EntityData.YFilter = asicErrorSbeSoft.YFilter
    asicErrorSbeSoft.EntityData.YangName = "asic-error-sbe-soft"
    asicErrorSbeSoft.EntityData.BundleName = "cisco_ios_xr"
    asicErrorSbeSoft.EntityData.ParentYangName = "error-path"
    asicErrorSbeSoft.EntityData.SegmentPath = "asic-error-sbe-soft"
    asicErrorSbeSoft.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorSbeSoft.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorSbeSoft.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorSbeSoft.EntityData.Children = make(map[string]types.YChild)
    asicErrorSbeSoft.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorSbeSoft.Error {
        asicErrorSbeSoft.EntityData.Children[types.GetSegmentPath(&asicErrorSbeSoft.Error[i])] = types.YChild{"Error", &asicErrorSbeSoft.Error[i]}
    }
    asicErrorSbeSoft.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorSbeSoft.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-sbe-soft"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors
// Hardware soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error
}

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetEntityData() *types.CommonEntityData {
    hardwareSoftErrors.EntityData.YFilter = hardwareSoftErrors.YFilter
    hardwareSoftErrors.EntityData.YangName = "hardware-soft-errors"
    hardwareSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    hardwareSoftErrors.EntityData.ParentYangName = "error-path"
    hardwareSoftErrors.EntityData.SegmentPath = "hardware-soft-errors"
    hardwareSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareSoftErrors.EntityData.Children = make(map[string]types.YChild)
    hardwareSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range hardwareSoftErrors.Error {
        hardwareSoftErrors.EntityData.Children[types.GetSegmentPath(&hardwareSoftErrors.Error[i])] = types.YChild{"Error", &hardwareSoftErrors.Error[i]}
    }
    hardwareSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "hardware-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error
}

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetEntityData() *types.CommonEntityData {
    asicErrorCrcSoft.EntityData.YFilter = asicErrorCrcSoft.YFilter
    asicErrorCrcSoft.EntityData.YangName = "asic-error-crc-soft"
    asicErrorCrcSoft.EntityData.BundleName = "cisco_ios_xr"
    asicErrorCrcSoft.EntityData.ParentYangName = "error-path"
    asicErrorCrcSoft.EntityData.SegmentPath = "asic-error-crc-soft"
    asicErrorCrcSoft.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorCrcSoft.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorCrcSoft.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorCrcSoft.EntityData.Children = make(map[string]types.YChild)
    asicErrorCrcSoft.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorCrcSoft.Error {
        asicErrorCrcSoft.EntityData.Children[types.GetSegmentPath(&asicErrorCrcSoft.Error[i])] = types.YChild{"Error", &asicErrorCrcSoft.Error[i]}
    }
    asicErrorCrcSoft.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorCrcSoft.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-crc-soft"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error
}

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetEntityData() *types.CommonEntityData {
    asicErrorParitySoft.EntityData.YFilter = asicErrorParitySoft.YFilter
    asicErrorParitySoft.EntityData.YangName = "asic-error-parity-soft"
    asicErrorParitySoft.EntityData.BundleName = "cisco_ios_xr"
    asicErrorParitySoft.EntityData.ParentYangName = "error-path"
    asicErrorParitySoft.EntityData.SegmentPath = "asic-error-parity-soft"
    asicErrorParitySoft.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorParitySoft.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorParitySoft.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorParitySoft.EntityData.Children = make(map[string]types.YChild)
    asicErrorParitySoft.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorParitySoft.Error {
        asicErrorParitySoft.EntityData.Children[types.GetSegmentPath(&asicErrorParitySoft.Error[i])] = types.YChild{"Error", &asicErrorParitySoft.Error[i]}
    }
    asicErrorParitySoft.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorParitySoft.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-parity-soft"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors
// IO soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error
}

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetEntityData() *types.CommonEntityData {
    ioSoftErrors.EntityData.YFilter = ioSoftErrors.YFilter
    ioSoftErrors.EntityData.YangName = "io-soft-errors"
    ioSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    ioSoftErrors.EntityData.ParentYangName = "error-path"
    ioSoftErrors.EntityData.SegmentPath = "io-soft-errors"
    ioSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ioSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ioSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ioSoftErrors.EntityData.Children = make(map[string]types.YChild)
    ioSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range ioSoftErrors.Error {
        ioSoftErrors.EntityData.Children[types.GetSegmentPath(&ioSoftErrors.Error[i])] = types.YChild{"Error", &ioSoftErrors.Error[i]}
    }
    ioSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ioSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "io-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors
// Reset soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error
}

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetEntityData() *types.CommonEntityData {
    resetSoftErrors.EntityData.YFilter = resetSoftErrors.YFilter
    resetSoftErrors.EntityData.YangName = "reset-soft-errors"
    resetSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    resetSoftErrors.EntityData.ParentYangName = "error-path"
    resetSoftErrors.EntityData.SegmentPath = "reset-soft-errors"
    resetSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resetSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resetSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resetSoftErrors.EntityData.Children = make(map[string]types.YChild)
    resetSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range resetSoftErrors.Error {
        resetSoftErrors.EntityData.Children[types.GetSegmentPath(&resetSoftErrors.Error[i])] = types.YChild{"Error", &resetSoftErrors.Error[i]}
    }
    resetSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(resetSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "reset-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors
// Barrier hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error
}

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetEntityData() *types.CommonEntityData {
    barrierHardErrors.EntityData.YFilter = barrierHardErrors.YFilter
    barrierHardErrors.EntityData.YangName = "barrier-hard-errors"
    barrierHardErrors.EntityData.BundleName = "cisco_ios_xr"
    barrierHardErrors.EntityData.ParentYangName = "error-path"
    barrierHardErrors.EntityData.SegmentPath = "barrier-hard-errors"
    barrierHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    barrierHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    barrierHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    barrierHardErrors.EntityData.Children = make(map[string]types.YChild)
    barrierHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range barrierHardErrors.Error {
        barrierHardErrors.EntityData.Children[types.GetSegmentPath(&barrierHardErrors.Error[i])] = types.YChild{"Error", &barrierHardErrors.Error[i]}
    }
    barrierHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(barrierHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "barrier-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors
// Ucode soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error
}

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetEntityData() *types.CommonEntityData {
    ucodeSoftErrors.EntityData.YFilter = ucodeSoftErrors.YFilter
    ucodeSoftErrors.EntityData.YangName = "ucode-soft-errors"
    ucodeSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    ucodeSoftErrors.EntityData.ParentYangName = "error-path"
    ucodeSoftErrors.EntityData.SegmentPath = "ucode-soft-errors"
    ucodeSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucodeSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucodeSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucodeSoftErrors.EntityData.Children = make(map[string]types.YChild)
    ucodeSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range ucodeSoftErrors.Error {
        ucodeSoftErrors.EntityData.Children[types.GetSegmentPath(&ucodeSoftErrors.Error[i])] = types.YChild{"Error", &ucodeSoftErrors.Error[i]}
    }
    ucodeSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ucodeSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "ucode-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error
}

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetEntityData() *types.CommonEntityData {
    asicErrorResetHard.EntityData.YFilter = asicErrorResetHard.YFilter
    asicErrorResetHard.EntityData.YangName = "asic-error-reset-hard"
    asicErrorResetHard.EntityData.BundleName = "cisco_ios_xr"
    asicErrorResetHard.EntityData.ParentYangName = "error-path"
    asicErrorResetHard.EntityData.SegmentPath = "asic-error-reset-hard"
    asicErrorResetHard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorResetHard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorResetHard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorResetHard.EntityData.Children = make(map[string]types.YChild)
    asicErrorResetHard.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorResetHard.Error {
        asicErrorResetHard.EntityData.Children[types.GetSegmentPath(&asicErrorResetHard.Error[i])] = types.YChild{"Error", &asicErrorResetHard.Error[i]}
    }
    asicErrorResetHard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorResetHard.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-reset-hard"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors
// Single bit hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error
}

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetEntityData() *types.CommonEntityData {
    singleBitHardErrors.EntityData.YFilter = singleBitHardErrors.YFilter
    singleBitHardErrors.EntityData.YangName = "single-bit-hard-errors"
    singleBitHardErrors.EntityData.BundleName = "cisco_ios_xr"
    singleBitHardErrors.EntityData.ParentYangName = "error-path"
    singleBitHardErrors.EntityData.SegmentPath = "single-bit-hard-errors"
    singleBitHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleBitHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleBitHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleBitHardErrors.EntityData.Children = make(map[string]types.YChild)
    singleBitHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range singleBitHardErrors.Error {
        singleBitHardErrors.EntityData.Children[types.GetSegmentPath(&singleBitHardErrors.Error[i])] = types.YChild{"Error", &singleBitHardErrors.Error[i]}
    }
    singleBitHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(singleBitHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "single-bit-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error
}

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetEntityData() *types.CommonEntityData {
    indirectHardErrors.EntityData.YFilter = indirectHardErrors.YFilter
    indirectHardErrors.EntityData.YangName = "indirect-hard-errors"
    indirectHardErrors.EntityData.BundleName = "cisco_ios_xr"
    indirectHardErrors.EntityData.ParentYangName = "error-path"
    indirectHardErrors.EntityData.SegmentPath = "indirect-hard-errors"
    indirectHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indirectHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indirectHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indirectHardErrors.EntityData.Children = make(map[string]types.YChild)
    indirectHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range indirectHardErrors.Error {
        indirectHardErrors.EntityData.Children[types.GetSegmentPath(&indirectHardErrors.Error[i])] = types.YChild{"Error", &indirectHardErrors.Error[i]}
    }
    indirectHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(indirectHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "indirect-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft
// OOR thresh information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error
}

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetEntityData() *types.CommonEntityData {
    outofResourceSoft.EntityData.YFilter = outofResourceSoft.YFilter
    outofResourceSoft.EntityData.YangName = "outof-resource-soft"
    outofResourceSoft.EntityData.BundleName = "cisco_ios_xr"
    outofResourceSoft.EntityData.ParentYangName = "error-path"
    outofResourceSoft.EntityData.SegmentPath = "outof-resource-soft"
    outofResourceSoft.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outofResourceSoft.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outofResourceSoft.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outofResourceSoft.EntityData.Children = make(map[string]types.YChild)
    outofResourceSoft.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range outofResourceSoft.Error {
        outofResourceSoft.EntityData.Children[types.GetSegmentPath(&outofResourceSoft.Error[i])] = types.YChild{"Error", &outofResourceSoft.Error[i]}
    }
    outofResourceSoft.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(outofResourceSoft.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "outof-resource-soft"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors
// CRC soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error
}

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetEntityData() *types.CommonEntityData {
    crcSoftErrors.EntityData.YFilter = crcSoftErrors.YFilter
    crcSoftErrors.EntityData.YangName = "crc-soft-errors"
    crcSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    crcSoftErrors.EntityData.ParentYangName = "error-path"
    crcSoftErrors.EntityData.SegmentPath = "crc-soft-errors"
    crcSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crcSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crcSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crcSoftErrors.EntityData.Children = make(map[string]types.YChild)
    crcSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range crcSoftErrors.Error {
        crcSoftErrors.EntityData.Children[types.GetSegmentPath(&crcSoftErrors.Error[i])] = types.YChild{"Error", &crcSoftErrors.Error[i]}
    }
    crcSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(crcSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "crc-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors
// Time out hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error
}

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetEntityData() *types.CommonEntityData {
    timeOutHardErrors.EntityData.YFilter = timeOutHardErrors.YFilter
    timeOutHardErrors.EntityData.YangName = "time-out-hard-errors"
    timeOutHardErrors.EntityData.BundleName = "cisco_ios_xr"
    timeOutHardErrors.EntityData.ParentYangName = "error-path"
    timeOutHardErrors.EntityData.SegmentPath = "time-out-hard-errors"
    timeOutHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeOutHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeOutHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeOutHardErrors.EntityData.Children = make(map[string]types.YChild)
    timeOutHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range timeOutHardErrors.Error {
        timeOutHardErrors.EntityData.Children[types.GetSegmentPath(&timeOutHardErrors.Error[i])] = types.YChild{"Error", &timeOutHardErrors.Error[i]}
    }
    timeOutHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timeOutHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "time-out-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors
// Barrier soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error
}

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetEntityData() *types.CommonEntityData {
    barrierSoftErrors.EntityData.YFilter = barrierSoftErrors.YFilter
    barrierSoftErrors.EntityData.YangName = "barrier-soft-errors"
    barrierSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    barrierSoftErrors.EntityData.ParentYangName = "error-path"
    barrierSoftErrors.EntityData.SegmentPath = "barrier-soft-errors"
    barrierSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    barrierSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    barrierSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    barrierSoftErrors.EntityData.Children = make(map[string]types.YChild)
    barrierSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range barrierSoftErrors.Error {
        barrierSoftErrors.EntityData.Children[types.GetSegmentPath(&barrierSoftErrors.Error[i])] = types.YChild{"Error", &barrierSoftErrors.Error[i]}
    }
    barrierSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(barrierSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "barrier-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error
}

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetEntityData() *types.CommonEntityData {
    asicErrorMbeSoft.EntityData.YFilter = asicErrorMbeSoft.YFilter
    asicErrorMbeSoft.EntityData.YangName = "asic-error-mbe-soft"
    asicErrorMbeSoft.EntityData.BundleName = "cisco_ios_xr"
    asicErrorMbeSoft.EntityData.ParentYangName = "error-path"
    asicErrorMbeSoft.EntityData.SegmentPath = "asic-error-mbe-soft"
    asicErrorMbeSoft.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorMbeSoft.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorMbeSoft.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorMbeSoft.EntityData.Children = make(map[string]types.YChild)
    asicErrorMbeSoft.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorMbeSoft.Error {
        asicErrorMbeSoft.EntityData.Children[types.GetSegmentPath(&asicErrorMbeSoft.Error[i])] = types.YChild{"Error", &asicErrorMbeSoft.Error[i]}
    }
    asicErrorMbeSoft.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorMbeSoft.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-mbe-soft"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors
// BP hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error
}

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetEntityData() *types.CommonEntityData {
    backPressureHardErrors.EntityData.YFilter = backPressureHardErrors.YFilter
    backPressureHardErrors.EntityData.YangName = "back-pressure-hard-errors"
    backPressureHardErrors.EntityData.BundleName = "cisco_ios_xr"
    backPressureHardErrors.EntityData.ParentYangName = "error-path"
    backPressureHardErrors.EntityData.SegmentPath = "back-pressure-hard-errors"
    backPressureHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backPressureHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backPressureHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backPressureHardErrors.EntityData.Children = make(map[string]types.YChild)
    backPressureHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range backPressureHardErrors.Error {
        backPressureHardErrors.EntityData.Children[types.GetSegmentPath(&backPressureHardErrors.Error[i])] = types.YChild{"Error", &backPressureHardErrors.Error[i]}
    }
    backPressureHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(backPressureHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "back-pressure-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors
// Single bit soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error
}

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetEntityData() *types.CommonEntityData {
    singleBitSoftErrors.EntityData.YFilter = singleBitSoftErrors.YFilter
    singleBitSoftErrors.EntityData.YangName = "single-bit-soft-errors"
    singleBitSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    singleBitSoftErrors.EntityData.ParentYangName = "error-path"
    singleBitSoftErrors.EntityData.SegmentPath = "single-bit-soft-errors"
    singleBitSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleBitSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleBitSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleBitSoftErrors.EntityData.Children = make(map[string]types.YChild)
    singleBitSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range singleBitSoftErrors.Error {
        singleBitSoftErrors.EntityData.Children[types.GetSegmentPath(&singleBitSoftErrors.Error[i])] = types.YChild{"Error", &singleBitSoftErrors.Error[i]}
    }
    singleBitSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(singleBitSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "single-bit-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors
// Indirect soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error
}

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetEntityData() *types.CommonEntityData {
    indirectSoftErrors.EntityData.YFilter = indirectSoftErrors.YFilter
    indirectSoftErrors.EntityData.YangName = "indirect-soft-errors"
    indirectSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    indirectSoftErrors.EntityData.ParentYangName = "error-path"
    indirectSoftErrors.EntityData.SegmentPath = "indirect-soft-errors"
    indirectSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indirectSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indirectSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indirectSoftErrors.EntityData.Children = make(map[string]types.YChild)
    indirectSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range indirectSoftErrors.Error {
        indirectSoftErrors.EntityData.Children[types.GetSegmentPath(&indirectSoftErrors.Error[i])] = types.YChild{"Error", &indirectSoftErrors.Error[i]}
    }
    indirectSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(indirectSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "indirect-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors
// Generic hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error
}

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetEntityData() *types.CommonEntityData {
    genericHardErrors.EntityData.YFilter = genericHardErrors.YFilter
    genericHardErrors.EntityData.YangName = "generic-hard-errors"
    genericHardErrors.EntityData.BundleName = "cisco_ios_xr"
    genericHardErrors.EntityData.ParentYangName = "error-path"
    genericHardErrors.EntityData.SegmentPath = "generic-hard-errors"
    genericHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericHardErrors.EntityData.Children = make(map[string]types.YChild)
    genericHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range genericHardErrors.Error {
        genericHardErrors.EntityData.Children[types.GetSegmentPath(&genericHardErrors.Error[i])] = types.YChild{"Error", &genericHardErrors.Error[i]}
    }
    genericHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(genericHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "generic-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors
// Link hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error
}

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetEntityData() *types.CommonEntityData {
    linkHardErrors.EntityData.YFilter = linkHardErrors.YFilter
    linkHardErrors.EntityData.YangName = "link-hard-errors"
    linkHardErrors.EntityData.BundleName = "cisco_ios_xr"
    linkHardErrors.EntityData.ParentYangName = "error-path"
    linkHardErrors.EntityData.SegmentPath = "link-hard-errors"
    linkHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkHardErrors.EntityData.Children = make(map[string]types.YChild)
    linkHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range linkHardErrors.Error {
        linkHardErrors.EntityData.Children[types.GetSegmentPath(&linkHardErrors.Error[i])] = types.YChild{"Error", &linkHardErrors.Error[i]}
    }
    linkHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(linkHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "link-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors
// Configuration hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error
}

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetEntityData() *types.CommonEntityData {
    configurationHardErrors.EntityData.YFilter = configurationHardErrors.YFilter
    configurationHardErrors.EntityData.YangName = "configuration-hard-errors"
    configurationHardErrors.EntityData.BundleName = "cisco_ios_xr"
    configurationHardErrors.EntityData.ParentYangName = "error-path"
    configurationHardErrors.EntityData.SegmentPath = "configuration-hard-errors"
    configurationHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationHardErrors.EntityData.Children = make(map[string]types.YChild)
    configurationHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range configurationHardErrors.Error {
        configurationHardErrors.EntityData.Children[types.GetSegmentPath(&configurationHardErrors.Error[i])] = types.YChild{"Error", &configurationHardErrors.Error[i]}
    }
    configurationHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configurationHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "configuration-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary
// Summary for a specific instance
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // legacy client. The type is bool.
    LegacyClient interface{}

    // cih client. The type is bool.
    CihClient interface{}

    // sum data. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData.
    SumData []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData
}

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetEntityData() *types.CommonEntityData {
    instanceSummary.EntityData.YFilter = instanceSummary.YFilter
    instanceSummary.EntityData.YangName = "instance-summary"
    instanceSummary.EntityData.BundleName = "cisco_ios_xr"
    instanceSummary.EntityData.ParentYangName = "error-path"
    instanceSummary.EntityData.SegmentPath = "instance-summary"
    instanceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instanceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instanceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instanceSummary.EntityData.Children = make(map[string]types.YChild)
    instanceSummary.EntityData.Children["sum-data"] = types.YChild{"SumData", nil}
    for i := range instanceSummary.SumData {
        instanceSummary.EntityData.Children[types.GetSegmentPath(&instanceSummary.SumData[i])] = types.YChild{"SumData", &instanceSummary.SumData[i]}
    }
    instanceSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    instanceSummary.EntityData.Leafs["legacy-client"] = types.YLeaf{"LegacyClient", instanceSummary.LegacyClient}
    instanceSummary.EntityData.Leafs["cih-client"] = types.YLeaf{"CihClient", instanceSummary.CihClient}
    return &(instanceSummary.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData
// sum data
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // num nodes. The type is interface{} with range: 0..4294967295.
    NumNodes interface{}

    // crc err count. The type is interface{} with range: 0..4294967295.
    CrcErrCount interface{}

    // sbe err count. The type is interface{} with range: 0..4294967295.
    SbeErrCount interface{}

    // mbe err count. The type is interface{} with range: 0..4294967295.
    MbeErrCount interface{}

    // par err count. The type is interface{} with range: 0..4294967295.
    ParErrCount interface{}

    // gen err count. The type is interface{} with range: 0..4294967295.
    GenErrCount interface{}

    // reset err count. The type is interface{} with range: 0..4294967295.
    ResetErrCount interface{}

    // err count. The type is slice of interface{} with range: 0..4294967295.
    ErrCount []interface{}

    // pcie err count. The type is slice of interface{} with range: 0..4294967295.
    PcieErrCount []interface{}

    // node key. The type is slice of interface{} with range: 0..4294967295.
    NodeKey []interface{}
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetEntityData() *types.CommonEntityData {
    sumData.EntityData.YFilter = sumData.YFilter
    sumData.EntityData.YangName = "sum-data"
    sumData.EntityData.BundleName = "cisco_ios_xr"
    sumData.EntityData.ParentYangName = "instance-summary"
    sumData.EntityData.SegmentPath = "sum-data"
    sumData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumData.EntityData.Children = make(map[string]types.YChild)
    sumData.EntityData.Leafs = make(map[string]types.YLeaf)
    sumData.EntityData.Leafs["num-nodes"] = types.YLeaf{"NumNodes", sumData.NumNodes}
    sumData.EntityData.Leafs["crc-err-count"] = types.YLeaf{"CrcErrCount", sumData.CrcErrCount}
    sumData.EntityData.Leafs["sbe-err-count"] = types.YLeaf{"SbeErrCount", sumData.SbeErrCount}
    sumData.EntityData.Leafs["mbe-err-count"] = types.YLeaf{"MbeErrCount", sumData.MbeErrCount}
    sumData.EntityData.Leafs["par-err-count"] = types.YLeaf{"ParErrCount", sumData.ParErrCount}
    sumData.EntityData.Leafs["gen-err-count"] = types.YLeaf{"GenErrCount", sumData.GenErrCount}
    sumData.EntityData.Leafs["reset-err-count"] = types.YLeaf{"ResetErrCount", sumData.ResetErrCount}
    sumData.EntityData.Leafs["err-count"] = types.YLeaf{"ErrCount", sumData.ErrCount}
    sumData.EntityData.Leafs["pcie-err-count"] = types.YLeaf{"PcieErrCount", sumData.PcieErrCount}
    sumData.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", sumData.NodeKey}
    return &(sumData.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors
// Unexpected hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error
}

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetEntityData() *types.CommonEntityData {
    unexpectedHardErrors.EntityData.YFilter = unexpectedHardErrors.YFilter
    unexpectedHardErrors.EntityData.YangName = "unexpected-hard-errors"
    unexpectedHardErrors.EntityData.BundleName = "cisco_ios_xr"
    unexpectedHardErrors.EntityData.ParentYangName = "error-path"
    unexpectedHardErrors.EntityData.SegmentPath = "unexpected-hard-errors"
    unexpectedHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unexpectedHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unexpectedHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unexpectedHardErrors.EntityData.Children = make(map[string]types.YChild)
    unexpectedHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range unexpectedHardErrors.Error {
        unexpectedHardErrors.EntityData.Children[types.GetSegmentPath(&unexpectedHardErrors.Error[i])] = types.YChild{"Error", &unexpectedHardErrors.Error[i]}
    }
    unexpectedHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(unexpectedHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "unexpected-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors
// Time out soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error
}

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetEntityData() *types.CommonEntityData {
    timeOutSoftErrors.EntityData.YFilter = timeOutSoftErrors.YFilter
    timeOutSoftErrors.EntityData.YangName = "time-out-soft-errors"
    timeOutSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    timeOutSoftErrors.EntityData.ParentYangName = "error-path"
    timeOutSoftErrors.EntityData.SegmentPath = "time-out-soft-errors"
    timeOutSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeOutSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeOutSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeOutSoftErrors.EntityData.Children = make(map[string]types.YChild)
    timeOutSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range timeOutSoftErrors.Error {
        timeOutSoftErrors.EntityData.Children[types.GetSegmentPath(&timeOutSoftErrors.Error[i])] = types.YChild{"Error", &timeOutSoftErrors.Error[i]}
    }
    timeOutSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timeOutSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "time-out-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error
}

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetEntityData() *types.CommonEntityData {
    asicErrorGenericHard.EntityData.YFilter = asicErrorGenericHard.YFilter
    asicErrorGenericHard.EntityData.YangName = "asic-error-generic-hard"
    asicErrorGenericHard.EntityData.BundleName = "cisco_ios_xr"
    asicErrorGenericHard.EntityData.ParentYangName = "error-path"
    asicErrorGenericHard.EntityData.SegmentPath = "asic-error-generic-hard"
    asicErrorGenericHard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorGenericHard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorGenericHard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorGenericHard.EntityData.Children = make(map[string]types.YChild)
    asicErrorGenericHard.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorGenericHard.Error {
        asicErrorGenericHard.EntityData.Children[types.GetSegmentPath(&asicErrorGenericHard.Error[i])] = types.YChild{"Error", &asicErrorGenericHard.Error[i]}
    }
    asicErrorGenericHard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorGenericHard.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-generic-hard"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors
// Parity hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error
}

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetEntityData() *types.CommonEntityData {
    parityHardErrors.EntityData.YFilter = parityHardErrors.YFilter
    parityHardErrors.EntityData.YangName = "parity-hard-errors"
    parityHardErrors.EntityData.BundleName = "cisco_ios_xr"
    parityHardErrors.EntityData.ParentYangName = "error-path"
    parityHardErrors.EntityData.SegmentPath = "parity-hard-errors"
    parityHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parityHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parityHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parityHardErrors.EntityData.Children = make(map[string]types.YChild)
    parityHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range parityHardErrors.Error {
        parityHardErrors.EntityData.Children[types.GetSegmentPath(&parityHardErrors.Error[i])] = types.YChild{"Error", &parityHardErrors.Error[i]}
    }
    parityHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(parityHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "parity-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors
// Descriptor hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error
}

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetEntityData() *types.CommonEntityData {
    descriptorHardErrors.EntityData.YFilter = descriptorHardErrors.YFilter
    descriptorHardErrors.EntityData.YangName = "descriptor-hard-errors"
    descriptorHardErrors.EntityData.BundleName = "cisco_ios_xr"
    descriptorHardErrors.EntityData.ParentYangName = "error-path"
    descriptorHardErrors.EntityData.SegmentPath = "descriptor-hard-errors"
    descriptorHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    descriptorHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    descriptorHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    descriptorHardErrors.EntityData.Children = make(map[string]types.YChild)
    descriptorHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range descriptorHardErrors.Error {
        descriptorHardErrors.EntityData.Children[types.GetSegmentPath(&descriptorHardErrors.Error[i])] = types.YChild{"Error", &descriptorHardErrors.Error[i]}
    }
    descriptorHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(descriptorHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "descriptor-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors
// Interface hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error
}

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetEntityData() *types.CommonEntityData {
    interfaceHardErrors.EntityData.YFilter = interfaceHardErrors.YFilter
    interfaceHardErrors.EntityData.YangName = "interface-hard-errors"
    interfaceHardErrors.EntityData.BundleName = "cisco_ios_xr"
    interfaceHardErrors.EntityData.ParentYangName = "error-path"
    interfaceHardErrors.EntityData.SegmentPath = "interface-hard-errors"
    interfaceHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceHardErrors.EntityData.Children = make(map[string]types.YChild)
    interfaceHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range interfaceHardErrors.Error {
        interfaceHardErrors.EntityData.Children[types.GetSegmentPath(&interfaceHardErrors.Error[i])] = types.YChild{"Error", &interfaceHardErrors.Error[i]}
    }
    interfaceHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "interface-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error
}

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetEntityData() *types.CommonEntityData {
    asicErrorSbeHard.EntityData.YFilter = asicErrorSbeHard.YFilter
    asicErrorSbeHard.EntityData.YangName = "asic-error-sbe-hard"
    asicErrorSbeHard.EntityData.BundleName = "cisco_ios_xr"
    asicErrorSbeHard.EntityData.ParentYangName = "error-path"
    asicErrorSbeHard.EntityData.SegmentPath = "asic-error-sbe-hard"
    asicErrorSbeHard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorSbeHard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorSbeHard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorSbeHard.EntityData.Children = make(map[string]types.YChild)
    asicErrorSbeHard.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorSbeHard.Error {
        asicErrorSbeHard.EntityData.Children[types.GetSegmentPath(&asicErrorSbeHard.Error[i])] = types.YChild{"Error", &asicErrorSbeHard.Error[i]}
    }
    asicErrorSbeHard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorSbeHard.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-sbe-hard"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error
}

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetEntityData() *types.CommonEntityData {
    asicErrorCrcHard.EntityData.YFilter = asicErrorCrcHard.YFilter
    asicErrorCrcHard.EntityData.YangName = "asic-error-crc-hard"
    asicErrorCrcHard.EntityData.BundleName = "cisco_ios_xr"
    asicErrorCrcHard.EntityData.ParentYangName = "error-path"
    asicErrorCrcHard.EntityData.SegmentPath = "asic-error-crc-hard"
    asicErrorCrcHard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorCrcHard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorCrcHard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorCrcHard.EntityData.Children = make(map[string]types.YChild)
    asicErrorCrcHard.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorCrcHard.Error {
        asicErrorCrcHard.EntityData.Children[types.GetSegmentPath(&asicErrorCrcHard.Error[i])] = types.YChild{"Error", &asicErrorCrcHard.Error[i]}
    }
    asicErrorCrcHard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorCrcHard.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-crc-hard"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error
}

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetEntityData() *types.CommonEntityData {
    asicErrorParityHard.EntityData.YFilter = asicErrorParityHard.YFilter
    asicErrorParityHard.EntityData.YangName = "asic-error-parity-hard"
    asicErrorParityHard.EntityData.BundleName = "cisco_ios_xr"
    asicErrorParityHard.EntityData.ParentYangName = "error-path"
    asicErrorParityHard.EntityData.SegmentPath = "asic-error-parity-hard"
    asicErrorParityHard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorParityHard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorParityHard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorParityHard.EntityData.Children = make(map[string]types.YChild)
    asicErrorParityHard.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorParityHard.Error {
        asicErrorParityHard.EntityData.Children[types.GetSegmentPath(&asicErrorParityHard.Error[i])] = types.YChild{"Error", &asicErrorParityHard.Error[i]}
    }
    asicErrorParityHard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorParityHard.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-parity-hard"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error
}

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetEntityData() *types.CommonEntityData {
    asicErrorResetSoft.EntityData.YFilter = asicErrorResetSoft.YFilter
    asicErrorResetSoft.EntityData.YangName = "asic-error-reset-soft"
    asicErrorResetSoft.EntityData.BundleName = "cisco_ios_xr"
    asicErrorResetSoft.EntityData.ParentYangName = "error-path"
    asicErrorResetSoft.EntityData.SegmentPath = "asic-error-reset-soft"
    asicErrorResetSoft.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorResetSoft.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorResetSoft.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorResetSoft.EntityData.Children = make(map[string]types.YChild)
    asicErrorResetSoft.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorResetSoft.Error {
        asicErrorResetSoft.EntityData.Children[types.GetSegmentPath(&asicErrorResetSoft.Error[i])] = types.YChild{"Error", &asicErrorResetSoft.Error[i]}
    }
    asicErrorResetSoft.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorResetSoft.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-reset-soft"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors
// BP soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error
}

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetEntityData() *types.CommonEntityData {
    backPressureSoftErrors.EntityData.YFilter = backPressureSoftErrors.YFilter
    backPressureSoftErrors.EntityData.YangName = "back-pressure-soft-errors"
    backPressureSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    backPressureSoftErrors.EntityData.ParentYangName = "error-path"
    backPressureSoftErrors.EntityData.SegmentPath = "back-pressure-soft-errors"
    backPressureSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backPressureSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backPressureSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backPressureSoftErrors.EntityData.Children = make(map[string]types.YChild)
    backPressureSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range backPressureSoftErrors.Error {
        backPressureSoftErrors.EntityData.Children[types.GetSegmentPath(&backPressureSoftErrors.Error[i])] = types.YChild{"Error", &backPressureSoftErrors.Error[i]}
    }
    backPressureSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(backPressureSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "back-pressure-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors
// Generic soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error
}

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetEntityData() *types.CommonEntityData {
    genericSoftErrors.EntityData.YFilter = genericSoftErrors.YFilter
    genericSoftErrors.EntityData.YangName = "generic-soft-errors"
    genericSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    genericSoftErrors.EntityData.ParentYangName = "error-path"
    genericSoftErrors.EntityData.SegmentPath = "generic-soft-errors"
    genericSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericSoftErrors.EntityData.Children = make(map[string]types.YChild)
    genericSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range genericSoftErrors.Error {
        genericSoftErrors.EntityData.Children[types.GetSegmentPath(&genericSoftErrors.Error[i])] = types.YChild{"Error", &genericSoftErrors.Error[i]}
    }
    genericSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(genericSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "generic-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors
// Link soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error
}

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetEntityData() *types.CommonEntityData {
    linkSoftErrors.EntityData.YFilter = linkSoftErrors.YFilter
    linkSoftErrors.EntityData.YangName = "link-soft-errors"
    linkSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    linkSoftErrors.EntityData.ParentYangName = "error-path"
    linkSoftErrors.EntityData.SegmentPath = "link-soft-errors"
    linkSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkSoftErrors.EntityData.Children = make(map[string]types.YChild)
    linkSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range linkSoftErrors.Error {
        linkSoftErrors.EntityData.Children[types.GetSegmentPath(&linkSoftErrors.Error[i])] = types.YChild{"Error", &linkSoftErrors.Error[i]}
    }
    linkSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(linkSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "link-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors
// Configuration soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error
}

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetEntityData() *types.CommonEntityData {
    configurationSoftErrors.EntityData.YFilter = configurationSoftErrors.YFilter
    configurationSoftErrors.EntityData.YangName = "configuration-soft-errors"
    configurationSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    configurationSoftErrors.EntityData.ParentYangName = "error-path"
    configurationSoftErrors.EntityData.SegmentPath = "configuration-soft-errors"
    configurationSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationSoftErrors.EntityData.Children = make(map[string]types.YChild)
    configurationSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range configurationSoftErrors.Error {
        configurationSoftErrors.EntityData.Children[types.GetSegmentPath(&configurationSoftErrors.Error[i])] = types.YChild{"Error", &configurationSoftErrors.Error[i]}
    }
    configurationSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configurationSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "configuration-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors
// Multiple bit hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error
}

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetEntityData() *types.CommonEntityData {
    multipleBitHardErrors.EntityData.YFilter = multipleBitHardErrors.YFilter
    multipleBitHardErrors.EntityData.YangName = "multiple-bit-hard-errors"
    multipleBitHardErrors.EntityData.BundleName = "cisco_ios_xr"
    multipleBitHardErrors.EntityData.ParentYangName = "error-path"
    multipleBitHardErrors.EntityData.SegmentPath = "multiple-bit-hard-errors"
    multipleBitHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multipleBitHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multipleBitHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multipleBitHardErrors.EntityData.Children = make(map[string]types.YChild)
    multipleBitHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range multipleBitHardErrors.Error {
        multipleBitHardErrors.EntityData.Children[types.GetSegmentPath(&multipleBitHardErrors.Error[i])] = types.YChild{"Error", &multipleBitHardErrors.Error[i]}
    }
    multipleBitHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multipleBitHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "multiple-bit-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors
// Unexpected soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error
}

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetEntityData() *types.CommonEntityData {
    unexpectedSoftErrors.EntityData.YFilter = unexpectedSoftErrors.YFilter
    unexpectedSoftErrors.EntityData.YangName = "unexpected-soft-errors"
    unexpectedSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    unexpectedSoftErrors.EntityData.ParentYangName = "error-path"
    unexpectedSoftErrors.EntityData.SegmentPath = "unexpected-soft-errors"
    unexpectedSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unexpectedSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unexpectedSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unexpectedSoftErrors.EntityData.Children = make(map[string]types.YChild)
    unexpectedSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range unexpectedSoftErrors.Error {
        unexpectedSoftErrors.EntityData.Children[types.GetSegmentPath(&unexpectedSoftErrors.Error[i])] = types.YChild{"Error", &unexpectedSoftErrors.Error[i]}
    }
    unexpectedSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(unexpectedSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "unexpected-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard
// OOR thresh information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error
}

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetEntityData() *types.CommonEntityData {
    outofResourceHard.EntityData.YFilter = outofResourceHard.YFilter
    outofResourceHard.EntityData.YangName = "outof-resource-hard"
    outofResourceHard.EntityData.BundleName = "cisco_ios_xr"
    outofResourceHard.EntityData.ParentYangName = "error-path"
    outofResourceHard.EntityData.SegmentPath = "outof-resource-hard"
    outofResourceHard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outofResourceHard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outofResourceHard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outofResourceHard.EntityData.Children = make(map[string]types.YChild)
    outofResourceHard.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range outofResourceHard.Error {
        outofResourceHard.EntityData.Children[types.GetSegmentPath(&outofResourceHard.Error[i])] = types.YChild{"Error", &outofResourceHard.Error[i]}
    }
    outofResourceHard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(outofResourceHard.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "outof-resource-hard"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors
// Hardware hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error
}

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetEntityData() *types.CommonEntityData {
    hardwareHardErrors.EntityData.YFilter = hardwareHardErrors.YFilter
    hardwareHardErrors.EntityData.YangName = "hardware-hard-errors"
    hardwareHardErrors.EntityData.BundleName = "cisco_ios_xr"
    hardwareHardErrors.EntityData.ParentYangName = "error-path"
    hardwareHardErrors.EntityData.SegmentPath = "hardware-hard-errors"
    hardwareHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareHardErrors.EntityData.Children = make(map[string]types.YChild)
    hardwareHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range hardwareHardErrors.Error {
        hardwareHardErrors.EntityData.Children[types.GetSegmentPath(&hardwareHardErrors.Error[i])] = types.YChild{"Error", &hardwareHardErrors.Error[i]}
    }
    hardwareHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "hardware-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors
// Parity soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error
}

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetEntityData() *types.CommonEntityData {
    paritySoftErrors.EntityData.YFilter = paritySoftErrors.YFilter
    paritySoftErrors.EntityData.YangName = "parity-soft-errors"
    paritySoftErrors.EntityData.BundleName = "cisco_ios_xr"
    paritySoftErrors.EntityData.ParentYangName = "error-path"
    paritySoftErrors.EntityData.SegmentPath = "parity-soft-errors"
    paritySoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paritySoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paritySoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paritySoftErrors.EntityData.Children = make(map[string]types.YChild)
    paritySoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range paritySoftErrors.Error {
        paritySoftErrors.EntityData.Children[types.GetSegmentPath(&paritySoftErrors.Error[i])] = types.YChild{"Error", &paritySoftErrors.Error[i]}
    }
    paritySoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paritySoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "parity-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors
// Descriptor soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error
}

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetEntityData() *types.CommonEntityData {
    descriptorSoftErrors.EntityData.YFilter = descriptorSoftErrors.YFilter
    descriptorSoftErrors.EntityData.YangName = "descriptor-soft-errors"
    descriptorSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    descriptorSoftErrors.EntityData.ParentYangName = "error-path"
    descriptorSoftErrors.EntityData.SegmentPath = "descriptor-soft-errors"
    descriptorSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    descriptorSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    descriptorSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    descriptorSoftErrors.EntityData.Children = make(map[string]types.YChild)
    descriptorSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range descriptorSoftErrors.Error {
        descriptorSoftErrors.EntityData.Children[types.GetSegmentPath(&descriptorSoftErrors.Error[i])] = types.YChild{"Error", &descriptorSoftErrors.Error[i]}
    }
    descriptorSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(descriptorSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "descriptor-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors
// Interface soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error
}

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetEntityData() *types.CommonEntityData {
    interfaceSoftErrors.EntityData.YFilter = interfaceSoftErrors.YFilter
    interfaceSoftErrors.EntityData.YangName = "interface-soft-errors"
    interfaceSoftErrors.EntityData.BundleName = "cisco_ios_xr"
    interfaceSoftErrors.EntityData.ParentYangName = "error-path"
    interfaceSoftErrors.EntityData.SegmentPath = "interface-soft-errors"
    interfaceSoftErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSoftErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSoftErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSoftErrors.EntityData.Children = make(map[string]types.YChild)
    interfaceSoftErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range interfaceSoftErrors.Error {
        interfaceSoftErrors.EntityData.Children[types.GetSegmentPath(&interfaceSoftErrors.Error[i])] = types.YChild{"Error", &interfaceSoftErrors.Error[i]}
    }
    interfaceSoftErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceSoftErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "interface-soft-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors
// IO hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error
}

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetEntityData() *types.CommonEntityData {
    ioHardErrors.EntityData.YFilter = ioHardErrors.YFilter
    ioHardErrors.EntityData.YangName = "io-hard-errors"
    ioHardErrors.EntityData.BundleName = "cisco_ios_xr"
    ioHardErrors.EntityData.ParentYangName = "error-path"
    ioHardErrors.EntityData.SegmentPath = "io-hard-errors"
    ioHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ioHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ioHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ioHardErrors.EntityData.Children = make(map[string]types.YChild)
    ioHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range ioHardErrors.Error {
        ioHardErrors.EntityData.Children[types.GetSegmentPath(&ioHardErrors.Error[i])] = types.YChild{"Error", &ioHardErrors.Error[i]}
    }
    ioHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ioHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "io-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors
// Reset hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error
}

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetEntityData() *types.CommonEntityData {
    resetHardErrors.EntityData.YFilter = resetHardErrors.YFilter
    resetHardErrors.EntityData.YangName = "reset-hard-errors"
    resetHardErrors.EntityData.BundleName = "cisco_ios_xr"
    resetHardErrors.EntityData.ParentYangName = "error-path"
    resetHardErrors.EntityData.SegmentPath = "reset-hard-errors"
    resetHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resetHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resetHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resetHardErrors.EntityData.Children = make(map[string]types.YChild)
    resetHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range resetHardErrors.Error {
        resetHardErrors.EntityData.Children[types.GetSegmentPath(&resetHardErrors.Error[i])] = types.YChild{"Error", &resetHardErrors.Error[i]}
    }
    resetHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(resetHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "reset-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors
// UCode hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error
}

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetEntityData() *types.CommonEntityData {
    ucodeHardErrors.EntityData.YFilter = ucodeHardErrors.YFilter
    ucodeHardErrors.EntityData.YangName = "ucode-hard-errors"
    ucodeHardErrors.EntityData.BundleName = "cisco_ios_xr"
    ucodeHardErrors.EntityData.ParentYangName = "error-path"
    ucodeHardErrors.EntityData.SegmentPath = "ucode-hard-errors"
    ucodeHardErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucodeHardErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucodeHardErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucodeHardErrors.EntityData.Children = make(map[string]types.YChild)
    ucodeHardErrors.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range ucodeHardErrors.Error {
        ucodeHardErrors.EntityData.Children[types.GetSegmentPath(&ucodeHardErrors.Error[i])] = types.YChild{"Error", &ucodeHardErrors.Error[i]}
    }
    ucodeHardErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ucodeHardErrors.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "ucode-hard-errors"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error
}

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetEntityData() *types.CommonEntityData {
    asicErrorMbeHard.EntityData.YFilter = asicErrorMbeHard.YFilter
    asicErrorMbeHard.EntityData.YangName = "asic-error-mbe-hard"
    asicErrorMbeHard.EntityData.BundleName = "cisco_ios_xr"
    asicErrorMbeHard.EntityData.ParentYangName = "error-path"
    asicErrorMbeHard.EntityData.SegmentPath = "asic-error-mbe-hard"
    asicErrorMbeHard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorMbeHard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorMbeHard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorMbeHard.EntityData.Children = make(map[string]types.YChild)
    asicErrorMbeHard.EntityData.Children["error"] = types.YChild{"Error", nil}
    for i := range asicErrorMbeHard.Error {
        asicErrorMbeHard.EntityData.Children[types.GetSegmentPath(&asicErrorMbeHard.Error[i])] = types.YChild{"Error", &asicErrorMbeHard.Error[i]}
    }
    asicErrorMbeHard.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicErrorMbeHard.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "asic-error-mbe-hard"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["csrs-info"] = types.YChild{"CsrsInfo", nil}
    for i := range error.CsrsInfo {
        error.EntityData.Children[types.GetSegmentPath(&error.CsrsInfo[i])] = types.YChild{"CsrsInfo", &error.CsrsInfo[i]}
    }
    error.EntityData.Children["last-err"] = types.YChild{"LastErr", nil}
    for i := range error.LastErr {
        error.EntityData.Children[types.GetSegmentPath(&error.LastErr[i])] = types.YChild{"LastErr", &error.LastErr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["name"] = types.YLeaf{"Name", error.Name}
    error.EntityData.Leafs["asic-info"] = types.YLeaf{"AsicInfo", error.AsicInfo}
    error.EntityData.Leafs["node-key"] = types.YLeaf{"NodeKey", error.NodeKey}
    error.EntityData.Leafs["alarm-on"] = types.YLeaf{"AlarmOn", error.AlarmOn}
    error.EntityData.Leafs["thresh-hi"] = types.YLeaf{"ThreshHi", error.ThreshHi}
    error.EntityData.Leafs["period-hi"] = types.YLeaf{"PeriodHi", error.PeriodHi}
    error.EntityData.Leafs["thresh-lo"] = types.YLeaf{"ThreshLo", error.ThreshLo}
    error.EntityData.Leafs["period-lo"] = types.YLeaf{"PeriodLo", error.PeriodLo}
    error.EntityData.Leafs["count"] = types.YLeaf{"Count", error.Count}
    error.EntityData.Leafs["intr-type"] = types.YLeaf{"IntrType", error.IntrType}
    error.EntityData.Leafs["leaf-id"] = types.YLeaf{"LeafId", error.LeafId}
    error.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", error.LastCleared}
    return &(error.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetEntityData() *types.CommonEntityData {
    csrsInfo.EntityData.YFilter = csrsInfo.YFilter
    csrsInfo.EntityData.YangName = "csrs-info"
    csrsInfo.EntityData.BundleName = "cisco_ios_xr"
    csrsInfo.EntityData.ParentYangName = "error"
    csrsInfo.EntityData.SegmentPath = "csrs-info"
    csrsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csrsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csrsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csrsInfo.EntityData.Children = make(map[string]types.YChild)
    csrsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    csrsInfo.EntityData.Leafs["name"] = types.YLeaf{"Name", csrsInfo.Name}
    csrsInfo.EntityData.Leafs["address"] = types.YLeaf{"Address", csrsInfo.Address}
    csrsInfo.EntityData.Leafs["width"] = types.YLeaf{"Width", csrsInfo.Width}
    return &(csrsInfo.EntityData)
}

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetEntityData() *types.CommonEntityData {
    lastErr.EntityData.YFilter = lastErr.YFilter
    lastErr.EntityData.YangName = "last-err"
    lastErr.EntityData.BundleName = "cisco_ios_xr"
    lastErr.EntityData.ParentYangName = "error"
    lastErr.EntityData.SegmentPath = "last-err"
    lastErr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErr.EntityData.Children = make(map[string]types.YChild)
    lastErr.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErr.EntityData.Leafs["at-time"] = types.YLeaf{"AtTime", lastErr.AtTime}
    lastErr.EntityData.Leafs["at-time-nsec"] = types.YLeaf{"AtTimeNsec", lastErr.AtTimeNsec}
    lastErr.EntityData.Leafs["counter-val"] = types.YLeaf{"CounterVal", lastErr.CounterVal}
    lastErr.EntityData.Leafs["error-desc"] = types.YLeaf{"ErrorDesc", lastErr.ErrorDesc}
    lastErr.EntityData.Leafs["error-regval"] = types.YLeaf{"ErrorRegval", lastErr.ErrorRegval}
    return &(lastErr.EntityData)
}

