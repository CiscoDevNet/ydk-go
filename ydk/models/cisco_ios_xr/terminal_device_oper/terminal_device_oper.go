// This module contains a collection of YANG definitions
// for Cisco IOS-XR terminal-device package operational data.
// 
// This module contains definitions
// for the following management objects:
//   optical-interface: System-wide view of interface operational
//     data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package terminal_device_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package terminal_device_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-terminal-device-oper optical-interface}", reflect.TypeOf(OpticalInterface{}))
    ydk.RegisterEntity("Cisco-IOS-XR-terminal-device-oper:optical-interface", reflect.TypeOf(OpticalInterface{}))
}

// LogicalProtocol represents Logical protocol
type LogicalProtocol string

const (
    // Unknown protocol framing
    LogicalProtocol_proto_type_unknown LogicalProtocol = "proto-type-unknown"

    // Ethernet protocol framing
    LogicalProtocol_proto_type_ethernet LogicalProtocol = "proto-type-ethernet"

    // OTN protocol framing
    LogicalProtocol_proto_type_otn LogicalProtocol = "proto-type-otn"
)

// TribProtocol represents Trib protocol
type TribProtocol string

const (
    // Unknown protocol
    TribProtocol_trib_proto_type_unknown TribProtocol = "trib-proto-type-unknown"

    // 1G Ethernet protocol
    TribProtocol_trib_proto_type1ge TribProtocol = "trib-proto-type1ge"

    // OC48 protocol
    TribProtocol_trib_proto_type_oc48 TribProtocol = "trib-proto-type-oc48"

    // STM 16 protocol
    TribProtocol_trib_proto_type_stm16 TribProtocol = "trib-proto-type-stm16"

    // 10G Ethernet LAN protocol
    TribProtocol_trib_proto_type10gelan TribProtocol = "trib-proto-type10gelan"

    // 10G Ethernet WAN protocol
    TribProtocol_trib_proto_type10gewan TribProtocol = "trib-proto-type10gewan"

    // OC 192 (9.6GB) port protocol
    TribProtocol_trib_proto_type_oc192 TribProtocol = "trib-proto-type-oc192"

    // STM 64 protocol
    TribProtocol_trib_proto_type_stm64 TribProtocol = "trib-proto-type-stm64"

    // OTU 2 protocol
    TribProtocol_trib_proto_type_otu2 TribProtocol = "trib-proto-type-otu2"

    // OTU 2e protocol
    TribProtocol_trib_proto_type_otu2e TribProtocol = "trib-proto-type-otu2e"

    // OTU 1e protocol
    TribProtocol_trib_proto_type_otu1e TribProtocol = "trib-proto-type-otu1e"

    // ODU 2 protocol
    TribProtocol_trib_proto_type_odu2 TribProtocol = "trib-proto-type-odu2"

    // ODU 2e protocol
    TribProtocol_trib_proto_type_odu2e TribProtocol = "trib-proto-type-odu2e"

    // 40G Ethernet port protocol
    TribProtocol_trib_proto_type40ge TribProtocol = "trib-proto-type40ge"

    // OC 768 protocol
    TribProtocol_trib_proto_type_oc768 TribProtocol = "trib-proto-type-oc768"

    // STM 256 protocol
    TribProtocol_trib_proto_type_stm256 TribProtocol = "trib-proto-type-stm256"

    // OTU 3 protocol
    TribProtocol_trib_proto_type_otu3 TribProtocol = "trib-proto-type-otu3"

    // ODU 3 protocol
    TribProtocol_trib_proto_type_odu3 TribProtocol = "trib-proto-type-odu3"

    // 100G Ethernet protocol
    TribProtocol_trib_proto_type100ge TribProtocol = "trib-proto-type100ge"

    // 100G MLG protocol
    TribProtocol_trib_proto_type100g_mlg TribProtocol = "trib-proto-type100g-mlg"

    // OTU4 signal protocol (112G) for transporting
    // 100GE signal
    TribProtocol_trib_proto_type_otu4 TribProtocol = "trib-proto-type-otu4"

    // OTU Cn protocol
    TribProtocol_trib_proto_type_otu_cn TribProtocol = "trib-proto-type-otu-cn"

    // ODU 4 protocol
    TribProtocol_trib_proto_type_odu4 TribProtocol = "trib-proto-type-odu4"
)

// TribRateClass represents Trib rate class
type TribRateClass string

const (
    // Unknown tributary signal rate
    TribRateClass_trib_rate_unknown TribRateClass = "trib-rate-unknown"

    // 1G tributary signal rate
    TribRateClass_trib_rate1g TribRateClass = "trib-rate1g"

    // 2.5G tributary signal rate
    TribRateClass_trib_rate25g TribRateClass = "trib-rate25g"

    // 10G tributary signal rate
    TribRateClass_trib_rate10g TribRateClass = "trib-rate10g"

    // 40G tributary signal rate
    TribRateClass_trib_rate40g TribRateClass = "trib-rate40g"

    // 100G tributary signal rate
    TribRateClass_trib_rate100g TribRateClass = "trib-rate100g"
)

// OpticalInterface
// System-wide view of interface operational data
type OpticalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table containing status information.
    ConfigStatus OpticalInterface_ConfigStatus

    // The operational attributes for a particular optical channel.
    OpticalChannelInterfaces OpticalInterface_OpticalChannelInterfaces

    // Table containing Graph Structure and related info.
    Graph OpticalInterface_Graph

    // The Operational Mode Table.
    OperationalModes OpticalInterface_OperationalModes

    // The operational attributes for a logical channel.
    OpticalLogicalInterfaces OpticalInterface_OpticalLogicalInterfaces
}

func (opticalInterface *OpticalInterface) GetEntityData() *types.CommonEntityData {
    opticalInterface.EntityData.YFilter = opticalInterface.YFilter
    opticalInterface.EntityData.YangName = "optical-interface"
    opticalInterface.EntityData.BundleName = "cisco_ios_xr"
    opticalInterface.EntityData.ParentYangName = "Cisco-IOS-XR-terminal-device-oper"
    opticalInterface.EntityData.SegmentPath = "Cisco-IOS-XR-terminal-device-oper:optical-interface"
    opticalInterface.EntityData.AbsolutePath = opticalInterface.EntityData.SegmentPath
    opticalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalInterface.EntityData.Children = types.NewOrderedMap()
    opticalInterface.EntityData.Children.Append("config-status", types.YChild{"ConfigStatus", &opticalInterface.ConfigStatus})
    opticalInterface.EntityData.Children.Append("optical-channel-interfaces", types.YChild{"OpticalChannelInterfaces", &opticalInterface.OpticalChannelInterfaces})
    opticalInterface.EntityData.Children.Append("graph", types.YChild{"Graph", &opticalInterface.Graph})
    opticalInterface.EntityData.Children.Append("operational-modes", types.YChild{"OperationalModes", &opticalInterface.OperationalModes})
    opticalInterface.EntityData.Children.Append("optical-logical-interfaces", types.YChild{"OpticalLogicalInterfaces", &opticalInterface.OpticalLogicalInterfaces})
    opticalInterface.EntityData.Leafs = types.NewOrderedMap()

    opticalInterface.EntityData.YListKeys = []string {}

    return &(opticalInterface.EntityData)
}

// OpticalInterface_ConfigStatus
// Table containing status information
type OpticalInterface_ConfigStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The bag containing partial config status.
    PartialConfig OpticalInterface_ConfigStatus_PartialConfig

    // The container containing slice status information.
    SliceTables OpticalInterface_ConfigStatus_SliceTables
}

func (configStatus *OpticalInterface_ConfigStatus) GetEntityData() *types.CommonEntityData {
    configStatus.EntityData.YFilter = configStatus.YFilter
    configStatus.EntityData.YangName = "config-status"
    configStatus.EntityData.BundleName = "cisco_ios_xr"
    configStatus.EntityData.ParentYangName = "optical-interface"
    configStatus.EntityData.SegmentPath = "config-status"
    configStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/" + configStatus.EntityData.SegmentPath
    configStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configStatus.EntityData.Children = types.NewOrderedMap()
    configStatus.EntityData.Children.Append("partial-config", types.YChild{"PartialConfig", &configStatus.PartialConfig})
    configStatus.EntityData.Children.Append("slice-tables", types.YChild{"SliceTables", &configStatus.SliceTables})
    configStatus.EntityData.Leafs = types.NewOrderedMap()

    configStatus.EntityData.YListKeys = []string {}

    return &(configStatus.EntityData)
}

// OpticalInterface_ConfigStatus_PartialConfig
// The bag containing partial config status
type OpticalInterface_ConfigStatus_PartialConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PartialConfig. The type is interface{} with range: 0..255.
    PartialConfig interface{}
}

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetEntityData() *types.CommonEntityData {
    partialConfig.EntityData.YFilter = partialConfig.YFilter
    partialConfig.EntityData.YangName = "partial-config"
    partialConfig.EntityData.BundleName = "cisco_ios_xr"
    partialConfig.EntityData.ParentYangName = "config-status"
    partialConfig.EntityData.SegmentPath = "partial-config"
    partialConfig.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/config-status/" + partialConfig.EntityData.SegmentPath
    partialConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    partialConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    partialConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    partialConfig.EntityData.Children = types.NewOrderedMap()
    partialConfig.EntityData.Leafs = types.NewOrderedMap()
    partialConfig.EntityData.Leafs.Append("partial-config", types.YLeaf{"PartialConfig", partialConfig.PartialConfig})

    partialConfig.EntityData.YListKeys = []string {}

    return &(partialConfig.EntityData)
}

// OpticalInterface_ConfigStatus_SliceTables
// The container containing slice status
// information
type OpticalInterface_ConfigStatus_SliceTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The table contains list of slices present. The type is slice of
    // OpticalInterface_ConfigStatus_SliceTables_SliceTable.
    SliceTable []*OpticalInterface_ConfigStatus_SliceTables_SliceTable
}

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetEntityData() *types.CommonEntityData {
    sliceTables.EntityData.YFilter = sliceTables.YFilter
    sliceTables.EntityData.YangName = "slice-tables"
    sliceTables.EntityData.BundleName = "cisco_ios_xr"
    sliceTables.EntityData.ParentYangName = "config-status"
    sliceTables.EntityData.SegmentPath = "slice-tables"
    sliceTables.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/config-status/" + sliceTables.EntityData.SegmentPath
    sliceTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceTables.EntityData.Children = types.NewOrderedMap()
    sliceTables.EntityData.Children.Append("slice-table", types.YChild{"SliceTable", nil})
    for i := range sliceTables.SliceTable {
        sliceTables.EntityData.Children.Append(types.GetSegmentPath(sliceTables.SliceTable[i]), types.YChild{"SliceTable", sliceTables.SliceTable[i]})
    }
    sliceTables.EntityData.Leafs = types.NewOrderedMap()

    sliceTables.EntityData.YListKeys = []string {}

    return &(sliceTables.EntityData)
}

// OpticalInterface_ConfigStatus_SliceTables_SliceTable
// The table contains list of slices present
type OpticalInterface_ConfigStatus_SliceTables_SliceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index of slice. The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // The bag containing slice config status.
    SliceStatusAttr OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr
}

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetEntityData() *types.CommonEntityData {
    sliceTable.EntityData.YFilter = sliceTable.YFilter
    sliceTable.EntityData.YangName = "slice-table"
    sliceTable.EntityData.BundleName = "cisco_ios_xr"
    sliceTable.EntityData.ParentYangName = "slice-tables"
    sliceTable.EntityData.SegmentPath = "slice-table" + types.AddKeyToken(sliceTable.Index, "index")
    sliceTable.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/config-status/slice-tables/" + sliceTable.EntityData.SegmentPath
    sliceTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceTable.EntityData.Children = types.NewOrderedMap()
    sliceTable.EntityData.Children.Append("slice-status-attr", types.YChild{"SliceStatusAttr", &sliceTable.SliceStatusAttr})
    sliceTable.EntityData.Leafs = types.NewOrderedMap()
    sliceTable.EntityData.Leafs.Append("index", types.YLeaf{"Index", sliceTable.Index})

    sliceTable.EntityData.YListKeys = []string {"Index"}

    return &(sliceTable.EntityData)
}

// OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr
// The bag containing slice config status
type OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slice. The type is interface{} with range: 0..255.
    Slice interface{}

    // ProvStatus. The type is string with length: 0..32.
    ProvStatus interface{}

    // PresentConfig. The type is string with length: 0..32.
    PresentConfig interface{}

    // PresentTimestamp. The type is string with length: 0..32.
    PresentTimestamp interface{}

    // PastConfig. The type is string with length: 0..32.
    PastConfig interface{}

    // PastTimestamp. The type is string with length: 0..32.
    PastTimestamp interface{}

    // ErrStr. The type is string with length: 0..1024.
    ErrStr interface{}

    // ErrTimestamp. The type is string with length: 0..32.
    ErrTimestamp interface{}
}

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetEntityData() *types.CommonEntityData {
    sliceStatusAttr.EntityData.YFilter = sliceStatusAttr.YFilter
    sliceStatusAttr.EntityData.YangName = "slice-status-attr"
    sliceStatusAttr.EntityData.BundleName = "cisco_ios_xr"
    sliceStatusAttr.EntityData.ParentYangName = "slice-table"
    sliceStatusAttr.EntityData.SegmentPath = "slice-status-attr"
    sliceStatusAttr.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/config-status/slice-tables/slice-table/" + sliceStatusAttr.EntityData.SegmentPath
    sliceStatusAttr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceStatusAttr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceStatusAttr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceStatusAttr.EntityData.Children = types.NewOrderedMap()
    sliceStatusAttr.EntityData.Leafs = types.NewOrderedMap()
    sliceStatusAttr.EntityData.Leafs.Append("slice", types.YLeaf{"Slice", sliceStatusAttr.Slice})
    sliceStatusAttr.EntityData.Leafs.Append("prov-status", types.YLeaf{"ProvStatus", sliceStatusAttr.ProvStatus})
    sliceStatusAttr.EntityData.Leafs.Append("present-config", types.YLeaf{"PresentConfig", sliceStatusAttr.PresentConfig})
    sliceStatusAttr.EntityData.Leafs.Append("present-timestamp", types.YLeaf{"PresentTimestamp", sliceStatusAttr.PresentTimestamp})
    sliceStatusAttr.EntityData.Leafs.Append("past-config", types.YLeaf{"PastConfig", sliceStatusAttr.PastConfig})
    sliceStatusAttr.EntityData.Leafs.Append("past-timestamp", types.YLeaf{"PastTimestamp", sliceStatusAttr.PastTimestamp})
    sliceStatusAttr.EntityData.Leafs.Append("err-str", types.YLeaf{"ErrStr", sliceStatusAttr.ErrStr})
    sliceStatusAttr.EntityData.Leafs.Append("err-timestamp", types.YLeaf{"ErrTimestamp", sliceStatusAttr.ErrTimestamp})

    sliceStatusAttr.EntityData.YListKeys = []string {}

    return &(sliceStatusAttr.EntityData)
}

// OpticalInterface_OpticalChannelInterfaces
// The operational attributes for a particular
// optical channel
type OpticalInterface_OpticalChannelInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The operational attributes for an optical channel. The type is slice of
    // OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface.
    OpticalChannelInterface []*OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface
}

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetEntityData() *types.CommonEntityData {
    opticalChannelInterfaces.EntityData.YFilter = opticalChannelInterfaces.YFilter
    opticalChannelInterfaces.EntityData.YangName = "optical-channel-interfaces"
    opticalChannelInterfaces.EntityData.BundleName = "cisco_ios_xr"
    opticalChannelInterfaces.EntityData.ParentYangName = "optical-interface"
    opticalChannelInterfaces.EntityData.SegmentPath = "optical-channel-interfaces"
    opticalChannelInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/" + opticalChannelInterfaces.EntityData.SegmentPath
    opticalChannelInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalChannelInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalChannelInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalChannelInterfaces.EntityData.Children = types.NewOrderedMap()
    opticalChannelInterfaces.EntityData.Children.Append("optical-channel-interface", types.YChild{"OpticalChannelInterface", nil})
    for i := range opticalChannelInterfaces.OpticalChannelInterface {
        opticalChannelInterfaces.EntityData.Children.Append(types.GetSegmentPath(opticalChannelInterfaces.OpticalChannelInterface[i]), types.YChild{"OpticalChannelInterface", opticalChannelInterfaces.OpticalChannelInterface[i]})
    }
    opticalChannelInterfaces.EntityData.Leafs = types.NewOrderedMap()

    opticalChannelInterfaces.EntityData.YListKeys = []string {}

    return &(opticalChannelInterfaces.EntityData)
}

// OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface
// The operational attributes for an optical
// channel
type OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the optical-channel. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Location interface{}

    // The operational attributes for an optical channel.
    OpticalChannelInterfaceAttr OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr
}

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetEntityData() *types.CommonEntityData {
    opticalChannelInterface.EntityData.YFilter = opticalChannelInterface.YFilter
    opticalChannelInterface.EntityData.YangName = "optical-channel-interface"
    opticalChannelInterface.EntityData.BundleName = "cisco_ios_xr"
    opticalChannelInterface.EntityData.ParentYangName = "optical-channel-interfaces"
    opticalChannelInterface.EntityData.SegmentPath = "optical-channel-interface" + types.AddKeyToken(opticalChannelInterface.Location, "location")
    opticalChannelInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/optical-channel-interfaces/" + opticalChannelInterface.EntityData.SegmentPath
    opticalChannelInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalChannelInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalChannelInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalChannelInterface.EntityData.Children = types.NewOrderedMap()
    opticalChannelInterface.EntityData.Children.Append("optical-channel-interface-attr", types.YChild{"OpticalChannelInterfaceAttr", &opticalChannelInterface.OpticalChannelInterfaceAttr})
    opticalChannelInterface.EntityData.Leafs = types.NewOrderedMap()
    opticalChannelInterface.EntityData.Leafs.Append("location", types.YLeaf{"Location", opticalChannelInterface.Location})

    opticalChannelInterface.EntityData.YListKeys = []string {"Location"}

    return &(opticalChannelInterface.EntityData)
}

// OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr
// The operational attributes for an optical
// channel
type OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string with length: 0..128.
    Name interface{}

    // Index. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // Frequency. The type is interface{} with range: 0..18446744073709551615.
    Frequency interface{}

    // Power. The type is interface{} with range: 0..18446744073709551615.
    Power interface{}

    // OperMode. The type is interface{} with range: 0..4294967295.
    OperMode interface{}

    // LinePort. The type is string with length: 0..128.
    LinePort interface{}
}

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetEntityData() *types.CommonEntityData {
    opticalChannelInterfaceAttr.EntityData.YFilter = opticalChannelInterfaceAttr.YFilter
    opticalChannelInterfaceAttr.EntityData.YangName = "optical-channel-interface-attr"
    opticalChannelInterfaceAttr.EntityData.BundleName = "cisco_ios_xr"
    opticalChannelInterfaceAttr.EntityData.ParentYangName = "optical-channel-interface"
    opticalChannelInterfaceAttr.EntityData.SegmentPath = "optical-channel-interface-attr"
    opticalChannelInterfaceAttr.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/optical-channel-interfaces/optical-channel-interface/" + opticalChannelInterfaceAttr.EntityData.SegmentPath
    opticalChannelInterfaceAttr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalChannelInterfaceAttr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalChannelInterfaceAttr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalChannelInterfaceAttr.EntityData.Children = types.NewOrderedMap()
    opticalChannelInterfaceAttr.EntityData.Leafs = types.NewOrderedMap()
    opticalChannelInterfaceAttr.EntityData.Leafs.Append("name", types.YLeaf{"Name", opticalChannelInterfaceAttr.Name})
    opticalChannelInterfaceAttr.EntityData.Leafs.Append("index", types.YLeaf{"Index", opticalChannelInterfaceAttr.Index})
    opticalChannelInterfaceAttr.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", opticalChannelInterfaceAttr.Frequency})
    opticalChannelInterfaceAttr.EntityData.Leafs.Append("power", types.YLeaf{"Power", opticalChannelInterfaceAttr.Power})
    opticalChannelInterfaceAttr.EntityData.Leafs.Append("oper-mode", types.YLeaf{"OperMode", opticalChannelInterfaceAttr.OperMode})
    opticalChannelInterfaceAttr.EntityData.Leafs.Append("line-port", types.YLeaf{"LinePort", opticalChannelInterfaceAttr.LinePort})

    opticalChannelInterfaceAttr.EntityData.YListKeys = []string {}

    return &(opticalChannelInterfaceAttr.EntityData)
}

// OpticalInterface_Graph
// Table containing Graph Structure and related
// info
type OpticalInterface_Graph struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The path containg file which has adjacency list stored.
    AdjListPath OpticalInterface_Graph_AdjListPath

    // The path containg file which has graph structure stored.
    GraphStructurePath OpticalInterface_Graph_GraphStructurePath
}

func (graph *OpticalInterface_Graph) GetEntityData() *types.CommonEntityData {
    graph.EntityData.YFilter = graph.YFilter
    graph.EntityData.YangName = "graph"
    graph.EntityData.BundleName = "cisco_ios_xr"
    graph.EntityData.ParentYangName = "optical-interface"
    graph.EntityData.SegmentPath = "graph"
    graph.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/" + graph.EntityData.SegmentPath
    graph.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    graph.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    graph.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    graph.EntityData.Children = types.NewOrderedMap()
    graph.EntityData.Children.Append("adj-list-path", types.YChild{"AdjListPath", &graph.AdjListPath})
    graph.EntityData.Children.Append("graph-structure-path", types.YChild{"GraphStructurePath", &graph.GraphStructurePath})
    graph.EntityData.Leafs = types.NewOrderedMap()

    graph.EntityData.YListKeys = []string {}

    return &(graph.EntityData)
}

// OpticalInterface_Graph_AdjListPath
// The path containg file which has adjacency list
// stored
type OpticalInterface_Graph_AdjListPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path. The type is string with length: 0..256.
    Path interface{}
}

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetEntityData() *types.CommonEntityData {
    adjListPath.EntityData.YFilter = adjListPath.YFilter
    adjListPath.EntityData.YangName = "adj-list-path"
    adjListPath.EntityData.BundleName = "cisco_ios_xr"
    adjListPath.EntityData.ParentYangName = "graph"
    adjListPath.EntityData.SegmentPath = "adj-list-path"
    adjListPath.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/graph/" + adjListPath.EntityData.SegmentPath
    adjListPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjListPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjListPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjListPath.EntityData.Children = types.NewOrderedMap()
    adjListPath.EntityData.Leafs = types.NewOrderedMap()
    adjListPath.EntityData.Leafs.Append("path", types.YLeaf{"Path", adjListPath.Path})

    adjListPath.EntityData.YListKeys = []string {}

    return &(adjListPath.EntityData)
}

// OpticalInterface_Graph_GraphStructurePath
// The path containg file which has graph
// structure stored
type OpticalInterface_Graph_GraphStructurePath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path. The type is string with length: 0..256.
    Path interface{}
}

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetEntityData() *types.CommonEntityData {
    graphStructurePath.EntityData.YFilter = graphStructurePath.YFilter
    graphStructurePath.EntityData.YangName = "graph-structure-path"
    graphStructurePath.EntityData.BundleName = "cisco_ios_xr"
    graphStructurePath.EntityData.ParentYangName = "graph"
    graphStructurePath.EntityData.SegmentPath = "graph-structure-path"
    graphStructurePath.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/graph/" + graphStructurePath.EntityData.SegmentPath
    graphStructurePath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    graphStructurePath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    graphStructurePath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    graphStructurePath.EntityData.Children = types.NewOrderedMap()
    graphStructurePath.EntityData.Leafs = types.NewOrderedMap()
    graphStructurePath.EntityData.Leafs.Append("path", types.YLeaf{"Path", graphStructurePath.Path})

    graphStructurePath.EntityData.YListKeys = []string {}

    return &(graphStructurePath.EntityData)
}

// OpticalInterface_OperationalModes
// The Operational Mode Table
type OpticalInterface_OperationalModes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mode supported on Device. The type is slice of
    // OpticalInterface_OperationalModes_OperationalMode.
    OperationalMode []*OpticalInterface_OperationalModes_OperationalMode
}

func (operationalModes *OpticalInterface_OperationalModes) GetEntityData() *types.CommonEntityData {
    operationalModes.EntityData.YFilter = operationalModes.YFilter
    operationalModes.EntityData.YangName = "operational-modes"
    operationalModes.EntityData.BundleName = "cisco_ios_xr"
    operationalModes.EntityData.ParentYangName = "optical-interface"
    operationalModes.EntityData.SegmentPath = "operational-modes"
    operationalModes.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/" + operationalModes.EntityData.SegmentPath
    operationalModes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationalModes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationalModes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationalModes.EntityData.Children = types.NewOrderedMap()
    operationalModes.EntityData.Children.Append("operational-mode", types.YChild{"OperationalMode", nil})
    for i := range operationalModes.OperationalMode {
        operationalModes.EntityData.Children.Append(types.GetSegmentPath(operationalModes.OperationalMode[i]), types.YChild{"OperationalMode", operationalModes.OperationalMode[i]})
    }
    operationalModes.EntityData.Leafs = types.NewOrderedMap()

    operationalModes.EntityData.YListKeys = []string {}

    return &(operationalModes.EntityData)
}

// OpticalInterface_OperationalModes_OperationalMode
// Mode supported on Device
type OpticalInterface_OperationalModes_OperationalMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Mode-id for supported mode on Device. The type is
    // interface{} with range: 0..4294967295.
    ModeId interface{}

    // The operational attributes for mxp driver fec-mode.
    OperationalModeAttributes OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes
}

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetEntityData() *types.CommonEntityData {
    operationalMode.EntityData.YFilter = operationalMode.YFilter
    operationalMode.EntityData.YangName = "operational-mode"
    operationalMode.EntityData.BundleName = "cisco_ios_xr"
    operationalMode.EntityData.ParentYangName = "operational-modes"
    operationalMode.EntityData.SegmentPath = "operational-mode" + types.AddKeyToken(operationalMode.ModeId, "mode-id")
    operationalMode.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/operational-modes/" + operationalMode.EntityData.SegmentPath
    operationalMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationalMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationalMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationalMode.EntityData.Children = types.NewOrderedMap()
    operationalMode.EntityData.Children.Append("operational-mode-attributes", types.YChild{"OperationalModeAttributes", &operationalMode.OperationalModeAttributes})
    operationalMode.EntityData.Leafs = types.NewOrderedMap()
    operationalMode.EntityData.Leafs.Append("mode-id", types.YLeaf{"ModeId", operationalMode.ModeId})

    operationalMode.EntityData.YListKeys = []string {"ModeId"}

    return &(operationalMode.EntityData)
}

// OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes
// The operational attributes for mxp driver
// fec-mode
type OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Description. The type is string with length: 0..128.
    Description interface{}

    // VendorId. The type is string with length: 0..64.
    VendorId interface{}
}

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetEntityData() *types.CommonEntityData {
    operationalModeAttributes.EntityData.YFilter = operationalModeAttributes.YFilter
    operationalModeAttributes.EntityData.YangName = "operational-mode-attributes"
    operationalModeAttributes.EntityData.BundleName = "cisco_ios_xr"
    operationalModeAttributes.EntityData.ParentYangName = "operational-mode"
    operationalModeAttributes.EntityData.SegmentPath = "operational-mode-attributes"
    operationalModeAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/operational-modes/operational-mode/" + operationalModeAttributes.EntityData.SegmentPath
    operationalModeAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationalModeAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationalModeAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationalModeAttributes.EntityData.Children = types.NewOrderedMap()
    operationalModeAttributes.EntityData.Leafs = types.NewOrderedMap()
    operationalModeAttributes.EntityData.Leafs.Append("description", types.YLeaf{"Description", operationalModeAttributes.Description})
    operationalModeAttributes.EntityData.Leafs.Append("vendor-id", types.YLeaf{"VendorId", operationalModeAttributes.VendorId})

    operationalModeAttributes.EntityData.YListKeys = []string {}

    return &(operationalModeAttributes.EntityData)
}

// OpticalInterface_OpticalLogicalInterfaces
// The operational attributes for a logical channel
type OpticalInterface_OpticalLogicalInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The operational attributes for a logical channel. The type is slice of
    // OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface.
    OpticalLogicalInterface []*OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface
}

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetEntityData() *types.CommonEntityData {
    opticalLogicalInterfaces.EntityData.YFilter = opticalLogicalInterfaces.YFilter
    opticalLogicalInterfaces.EntityData.YangName = "optical-logical-interfaces"
    opticalLogicalInterfaces.EntityData.BundleName = "cisco_ios_xr"
    opticalLogicalInterfaces.EntityData.ParentYangName = "optical-interface"
    opticalLogicalInterfaces.EntityData.SegmentPath = "optical-logical-interfaces"
    opticalLogicalInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/" + opticalLogicalInterfaces.EntityData.SegmentPath
    opticalLogicalInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalLogicalInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalLogicalInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalLogicalInterfaces.EntityData.Children = types.NewOrderedMap()
    opticalLogicalInterfaces.EntityData.Children.Append("optical-logical-interface", types.YChild{"OpticalLogicalInterface", nil})
    for i := range opticalLogicalInterfaces.OpticalLogicalInterface {
        opticalLogicalInterfaces.EntityData.Children.Append(types.GetSegmentPath(opticalLogicalInterfaces.OpticalLogicalInterface[i]), types.YChild{"OpticalLogicalInterface", opticalLogicalInterfaces.OpticalLogicalInterface[i]})
    }
    opticalLogicalInterfaces.EntityData.Leafs = types.NewOrderedMap()

    opticalLogicalInterfaces.EntityData.YListKeys = []string {}

    return &(opticalLogicalInterfaces.EntityData)
}

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface
// The operational attributes for a logical
// channel
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index of the logical-channel. The type is
    // interface{} with range: 0..4294967295.
    Index interface{}

    // The operational attributes for a particular logical channel.
    OpticalLogicalInterfaceAttr OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr

    // The operational attributes for a particular interface.
    OpticalLogicalInterfaceLogicalChannelAssignments OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments
}

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetEntityData() *types.CommonEntityData {
    opticalLogicalInterface.EntityData.YFilter = opticalLogicalInterface.YFilter
    opticalLogicalInterface.EntityData.YangName = "optical-logical-interface"
    opticalLogicalInterface.EntityData.BundleName = "cisco_ios_xr"
    opticalLogicalInterface.EntityData.ParentYangName = "optical-logical-interfaces"
    opticalLogicalInterface.EntityData.SegmentPath = "optical-logical-interface" + types.AddKeyToken(opticalLogicalInterface.Index, "index")
    opticalLogicalInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/optical-logical-interfaces/" + opticalLogicalInterface.EntityData.SegmentPath
    opticalLogicalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalLogicalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalLogicalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalLogicalInterface.EntityData.Children = types.NewOrderedMap()
    opticalLogicalInterface.EntityData.Children.Append("optical-logical-interface-attr", types.YChild{"OpticalLogicalInterfaceAttr", &opticalLogicalInterface.OpticalLogicalInterfaceAttr})
    opticalLogicalInterface.EntityData.Children.Append("optical-logical-interface-logical-channel-assignments", types.YChild{"OpticalLogicalInterfaceLogicalChannelAssignments", &opticalLogicalInterface.OpticalLogicalInterfaceLogicalChannelAssignments})
    opticalLogicalInterface.EntityData.Leafs = types.NewOrderedMap()
    opticalLogicalInterface.EntityData.Leafs.Append("index", types.YLeaf{"Index", opticalLogicalInterface.Index})

    opticalLogicalInterface.EntityData.YListKeys = []string {"Index"}

    return &(opticalLogicalInterface.EntityData)
}

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr
// The operational attributes for a particular
// logical channel
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LogicalChannelIndex. The type is interface{} with range: 0..4294967295.
    LogicalChannelIndex interface{}

    // LogicalChannelIfname. The type is string with length: 0..128.
    LogicalChannelIfname interface{}

    // Type. The type is string with length: 0..32.
    Type interface{}

    // TribRateClass. The type is TribRateClass.
    TribRateClass interface{}

    // TribProtocol. The type is TribProtocol.
    TribProtocol interface{}

    // ProtocolType. The type is LogicalProtocol.
    ProtocolType interface{}

    // AdminState. The type is interface{} with range: 0..4294967295.
    AdminState interface{}

    // LoopbackMode. The type is interface{} with range: 0..4294967295.
    LoopbackMode interface{}

    // IngressClientPort. The type is string with length: 0..128.
    IngressClientPort interface{}

    // IngressPhysicalChannel. The type is interface{} with range: 0..4294967295.
    IngressPhysicalChannel interface{}

    // TtiTransmit. The type is string with length: 0..256.
    TtiTransmit interface{}

    // TtiExpected. The type is string with length: 0..256.
    TtiExpected interface{}
}

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetEntityData() *types.CommonEntityData {
    opticalLogicalInterfaceAttr.EntityData.YFilter = opticalLogicalInterfaceAttr.YFilter
    opticalLogicalInterfaceAttr.EntityData.YangName = "optical-logical-interface-attr"
    opticalLogicalInterfaceAttr.EntityData.BundleName = "cisco_ios_xr"
    opticalLogicalInterfaceAttr.EntityData.ParentYangName = "optical-logical-interface"
    opticalLogicalInterfaceAttr.EntityData.SegmentPath = "optical-logical-interface-attr"
    opticalLogicalInterfaceAttr.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/optical-logical-interfaces/optical-logical-interface/" + opticalLogicalInterfaceAttr.EntityData.SegmentPath
    opticalLogicalInterfaceAttr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalLogicalInterfaceAttr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalLogicalInterfaceAttr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalLogicalInterfaceAttr.EntityData.Children = types.NewOrderedMap()
    opticalLogicalInterfaceAttr.EntityData.Leafs = types.NewOrderedMap()
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("logical-channel-index", types.YLeaf{"LogicalChannelIndex", opticalLogicalInterfaceAttr.LogicalChannelIndex})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("logical-channel-ifname", types.YLeaf{"LogicalChannelIfname", opticalLogicalInterfaceAttr.LogicalChannelIfname})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("type", types.YLeaf{"Type", opticalLogicalInterfaceAttr.Type})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("trib-rate-class", types.YLeaf{"TribRateClass", opticalLogicalInterfaceAttr.TribRateClass})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("trib-protocol", types.YLeaf{"TribProtocol", opticalLogicalInterfaceAttr.TribProtocol})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("protocol-type", types.YLeaf{"ProtocolType", opticalLogicalInterfaceAttr.ProtocolType})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", opticalLogicalInterfaceAttr.AdminState})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", opticalLogicalInterfaceAttr.LoopbackMode})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("ingress-client-port", types.YLeaf{"IngressClientPort", opticalLogicalInterfaceAttr.IngressClientPort})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("ingress-physical-channel", types.YLeaf{"IngressPhysicalChannel", opticalLogicalInterfaceAttr.IngressPhysicalChannel})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("tti-transmit", types.YLeaf{"TtiTransmit", opticalLogicalInterfaceAttr.TtiTransmit})
    opticalLogicalInterfaceAttr.EntityData.Leafs.Append("tti-expected", types.YLeaf{"TtiExpected", opticalLogicalInterfaceAttr.TtiExpected})

    opticalLogicalInterfaceAttr.EntityData.YListKeys = []string {}

    return &(opticalLogicalInterfaceAttr.EntityData)
}

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments
// The operational attributes for a particular
// interface
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The operational attributes for a logical channel assignment. The type is
    // slice of
    // OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment.
    OpticalLogicalInterfaceLogicalChannelAssignment []*OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment
}

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetEntityData() *types.CommonEntityData {
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.YFilter = opticalLogicalInterfaceLogicalChannelAssignments.YFilter
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.YangName = "optical-logical-interface-logical-channel-assignments"
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.BundleName = "cisco_ios_xr"
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.ParentYangName = "optical-logical-interface"
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.SegmentPath = "optical-logical-interface-logical-channel-assignments"
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/optical-logical-interfaces/optical-logical-interface/" + opticalLogicalInterfaceLogicalChannelAssignments.EntityData.SegmentPath
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.Children = types.NewOrderedMap()
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.Children.Append("optical-logical-interface-logical-channel-assignment", types.YChild{"OpticalLogicalInterfaceLogicalChannelAssignment", nil})
    for i := range opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment {
        opticalLogicalInterfaceLogicalChannelAssignments.EntityData.Children.Append(types.GetSegmentPath(opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment[i]), types.YChild{"OpticalLogicalInterfaceLogicalChannelAssignment", opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment[i]})
    }
    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.Leafs = types.NewOrderedMap()

    opticalLogicalInterfaceLogicalChannelAssignments.EntityData.YListKeys = []string {}

    return &(opticalLogicalInterfaceLogicalChannelAssignments.EntityData)
}

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment
// The operational attributes for a logical
// channel assignment
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index of the logical-channel. The type is
    // interface{} with range: 0..4294967295.
    Index interface{}

    // The operational attributes for a logical channel assignment.
    OpticalLogicalInterfaceLogicalChannelAssignmentAttr OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr
}

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetEntityData() *types.CommonEntityData {
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.YFilter = opticalLogicalInterfaceLogicalChannelAssignment.YFilter
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.YangName = "optical-logical-interface-logical-channel-assignment"
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.BundleName = "cisco_ios_xr"
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.ParentYangName = "optical-logical-interface-logical-channel-assignments"
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.SegmentPath = "optical-logical-interface-logical-channel-assignment" + types.AddKeyToken(opticalLogicalInterfaceLogicalChannelAssignment.Index, "index")
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/optical-logical-interfaces/optical-logical-interface/optical-logical-interface-logical-channel-assignments/" + opticalLogicalInterfaceLogicalChannelAssignment.EntityData.SegmentPath
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.Children = types.NewOrderedMap()
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.Children.Append("optical-logical-interface-logical-channel-assignment-attr", types.YChild{"OpticalLogicalInterfaceLogicalChannelAssignmentAttr", &opticalLogicalInterfaceLogicalChannelAssignment.OpticalLogicalInterfaceLogicalChannelAssignmentAttr})
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.Leafs = types.NewOrderedMap()
    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.Leafs.Append("index", types.YLeaf{"Index", opticalLogicalInterfaceLogicalChannelAssignment.Index})

    opticalLogicalInterfaceLogicalChannelAssignment.EntityData.YListKeys = []string {"Index"}

    return &(opticalLogicalInterfaceLogicalChannelAssignment.EntityData)
}

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr
// The operational attributes for a logical
// channel assignment
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // Name. The type is string with length: 0..128.
    Name interface{}

    // IsLogicalLink. The type is bool.
    IsLogicalLink interface{}

    // LogicalChannel. The type is interface{} with range: 0..4294967295.
    LogicalChannel interface{}

    // OpticalChannel. The type is string with length: 0..128.
    OpticalChannel interface{}

    // Allocation. The type is interface{} with range: 0..4294967295.
    Allocation interface{}

    // AssignmentType. The type is interface{} with range: 0..4294967295.
    AssignmentType interface{}
}

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetEntityData() *types.CommonEntityData {
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.YFilter = opticalLogicalInterfaceLogicalChannelAssignmentAttr.YFilter
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.YangName = "optical-logical-interface-logical-channel-assignment-attr"
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.BundleName = "cisco_ios_xr"
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.ParentYangName = "optical-logical-interface-logical-channel-assignment"
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.SegmentPath = "optical-logical-interface-logical-channel-assignment-attr"
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-oper:optical-interface/optical-logical-interfaces/optical-logical-interface/optical-logical-interface-logical-channel-assignments/optical-logical-interface-logical-channel-assignment/" + opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.SegmentPath
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Children = types.NewOrderedMap()
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs = types.NewOrderedMap()
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs.Append("index", types.YLeaf{"Index", opticalLogicalInterfaceLogicalChannelAssignmentAttr.Index})
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs.Append("name", types.YLeaf{"Name", opticalLogicalInterfaceLogicalChannelAssignmentAttr.Name})
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs.Append("is-logical-link", types.YLeaf{"IsLogicalLink", opticalLogicalInterfaceLogicalChannelAssignmentAttr.IsLogicalLink})
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs.Append("logical-channel", types.YLeaf{"LogicalChannel", opticalLogicalInterfaceLogicalChannelAssignmentAttr.LogicalChannel})
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs.Append("optical-channel", types.YLeaf{"OpticalChannel", opticalLogicalInterfaceLogicalChannelAssignmentAttr.OpticalChannel})
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs.Append("allocation", types.YLeaf{"Allocation", opticalLogicalInterfaceLogicalChannelAssignmentAttr.Allocation})
    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.Leafs.Append("assignment-type", types.YLeaf{"AssignmentType", opticalLogicalInterfaceLogicalChannelAssignmentAttr.AssignmentType})

    opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData.YListKeys = []string {}

    return &(opticalLogicalInterfaceLogicalChannelAssignmentAttr.EntityData)
}

