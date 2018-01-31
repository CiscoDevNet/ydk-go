// This module contains a collection of YANG definitions
// for Cisco IOS-XR terminal-device package operational data.
// 
// This module contains definitions
// for the following management objects:
//   optical-interface: System-wide view of interface operational
//     data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
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

func (opticalInterface *OpticalInterface) GetFilter() yfilter.YFilter { return opticalInterface.YFilter }

func (opticalInterface *OpticalInterface) SetFilter(yf yfilter.YFilter) { opticalInterface.YFilter = yf }

func (opticalInterface *OpticalInterface) GetGoName(yname string) string {
    if yname == "config-status" { return "ConfigStatus" }
    if yname == "optical-channel-interfaces" { return "OpticalChannelInterfaces" }
    if yname == "graph" { return "Graph" }
    if yname == "operational-modes" { return "OperationalModes" }
    if yname == "optical-logical-interfaces" { return "OpticalLogicalInterfaces" }
    return ""
}

func (opticalInterface *OpticalInterface) GetSegmentPath() string {
    return "Cisco-IOS-XR-terminal-device-oper:optical-interface"
}

func (opticalInterface *OpticalInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config-status" {
        return &opticalInterface.ConfigStatus
    }
    if childYangName == "optical-channel-interfaces" {
        return &opticalInterface.OpticalChannelInterfaces
    }
    if childYangName == "graph" {
        return &opticalInterface.Graph
    }
    if childYangName == "operational-modes" {
        return &opticalInterface.OperationalModes
    }
    if childYangName == "optical-logical-interfaces" {
        return &opticalInterface.OpticalLogicalInterfaces
    }
    return nil
}

func (opticalInterface *OpticalInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config-status"] = &opticalInterface.ConfigStatus
    children["optical-channel-interfaces"] = &opticalInterface.OpticalChannelInterfaces
    children["graph"] = &opticalInterface.Graph
    children["operational-modes"] = &opticalInterface.OperationalModes
    children["optical-logical-interfaces"] = &opticalInterface.OpticalLogicalInterfaces
    return children
}

func (opticalInterface *OpticalInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalInterface *OpticalInterface) GetBundleName() string { return "cisco_ios_xr" }

func (opticalInterface *OpticalInterface) GetYangName() string { return "optical-interface" }

func (opticalInterface *OpticalInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalInterface *OpticalInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalInterface *OpticalInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalInterface *OpticalInterface) SetParent(parent types.Entity) { opticalInterface.parent = parent }

func (opticalInterface *OpticalInterface) GetParent() types.Entity { return opticalInterface.parent }

func (opticalInterface *OpticalInterface) GetParentYangName() string { return "Cisco-IOS-XR-terminal-device-oper" }

// OpticalInterface_ConfigStatus
// Table containing status information
type OpticalInterface_ConfigStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The bag containing partial config status.
    PartialConfig OpticalInterface_ConfigStatus_PartialConfig

    // The container containing slice status information.
    SliceTables OpticalInterface_ConfigStatus_SliceTables
}

func (configStatus *OpticalInterface_ConfigStatus) GetFilter() yfilter.YFilter { return configStatus.YFilter }

func (configStatus *OpticalInterface_ConfigStatus) SetFilter(yf yfilter.YFilter) { configStatus.YFilter = yf }

func (configStatus *OpticalInterface_ConfigStatus) GetGoName(yname string) string {
    if yname == "partial-config" { return "PartialConfig" }
    if yname == "slice-tables" { return "SliceTables" }
    return ""
}

func (configStatus *OpticalInterface_ConfigStatus) GetSegmentPath() string {
    return "config-status"
}

func (configStatus *OpticalInterface_ConfigStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "partial-config" {
        return &configStatus.PartialConfig
    }
    if childYangName == "slice-tables" {
        return &configStatus.SliceTables
    }
    return nil
}

func (configStatus *OpticalInterface_ConfigStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["partial-config"] = &configStatus.PartialConfig
    children["slice-tables"] = &configStatus.SliceTables
    return children
}

func (configStatus *OpticalInterface_ConfigStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configStatus *OpticalInterface_ConfigStatus) GetBundleName() string { return "cisco_ios_xr" }

func (configStatus *OpticalInterface_ConfigStatus) GetYangName() string { return "config-status" }

func (configStatus *OpticalInterface_ConfigStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configStatus *OpticalInterface_ConfigStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configStatus *OpticalInterface_ConfigStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configStatus *OpticalInterface_ConfigStatus) SetParent(parent types.Entity) { configStatus.parent = parent }

func (configStatus *OpticalInterface_ConfigStatus) GetParent() types.Entity { return configStatus.parent }

func (configStatus *OpticalInterface_ConfigStatus) GetParentYangName() string { return "optical-interface" }

// OpticalInterface_ConfigStatus_PartialConfig
// The bag containing partial config status
type OpticalInterface_ConfigStatus_PartialConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PartialConfig. The type is interface{} with range: 0..255.
    PartialConfig interface{}
}

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetFilter() yfilter.YFilter { return partialConfig.YFilter }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) SetFilter(yf yfilter.YFilter) { partialConfig.YFilter = yf }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetGoName(yname string) string {
    if yname == "partial-config" { return "PartialConfig" }
    return ""
}

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetSegmentPath() string {
    return "partial-config"
}

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["partial-config"] = partialConfig.PartialConfig
    return leafs
}

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetBundleName() string { return "cisco_ios_xr" }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetYangName() string { return "partial-config" }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) SetParent(parent types.Entity) { partialConfig.parent = parent }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetParent() types.Entity { return partialConfig.parent }

func (partialConfig *OpticalInterface_ConfigStatus_PartialConfig) GetParentYangName() string { return "config-status" }

// OpticalInterface_ConfigStatus_SliceTables
// The container containing slice status
// information
type OpticalInterface_ConfigStatus_SliceTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The table contains list of slices present. The type is slice of
    // OpticalInterface_ConfigStatus_SliceTables_SliceTable.
    SliceTable []OpticalInterface_ConfigStatus_SliceTables_SliceTable
}

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetFilter() yfilter.YFilter { return sliceTables.YFilter }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) SetFilter(yf yfilter.YFilter) { sliceTables.YFilter = yf }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetGoName(yname string) string {
    if yname == "slice-table" { return "SliceTable" }
    return ""
}

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetSegmentPath() string {
    return "slice-tables"
}

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-table" {
        for _, c := range sliceTables.SliceTable {
            if sliceTables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalInterface_ConfigStatus_SliceTables_SliceTable{}
        sliceTables.SliceTable = append(sliceTables.SliceTable, child)
        return &sliceTables.SliceTable[len(sliceTables.SliceTable)-1]
    }
    return nil
}

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceTables.SliceTable {
        children[sliceTables.SliceTable[i].GetSegmentPath()] = &sliceTables.SliceTable[i]
    }
    return children
}

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetBundleName() string { return "cisco_ios_xr" }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetYangName() string { return "slice-tables" }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) SetParent(parent types.Entity) { sliceTables.parent = parent }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetParent() types.Entity { return sliceTables.parent }

func (sliceTables *OpticalInterface_ConfigStatus_SliceTables) GetParentYangName() string { return "config-status" }

// OpticalInterface_ConfigStatus_SliceTables_SliceTable
// The table contains list of slices present
type OpticalInterface_ConfigStatus_SliceTables_SliceTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index of slice. The type is interface{} with
    // range: -2147483648..2147483647.
    Index interface{}

    // The bag containing slice config status.
    SliceStatusAttr OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr
}

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetFilter() yfilter.YFilter { return sliceTable.YFilter }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) SetFilter(yf yfilter.YFilter) { sliceTable.YFilter = yf }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "slice-status-attr" { return "SliceStatusAttr" }
    return ""
}

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetSegmentPath() string {
    return "slice-table" + "[index='" + fmt.Sprintf("%v", sliceTable.Index) + "']"
}

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-status-attr" {
        return &sliceTable.SliceStatusAttr
    }
    return nil
}

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slice-status-attr"] = &sliceTable.SliceStatusAttr
    return children
}

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = sliceTable.Index
    return leafs
}

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetBundleName() string { return "cisco_ios_xr" }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetYangName() string { return "slice-table" }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) SetParent(parent types.Entity) { sliceTable.parent = parent }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetParent() types.Entity { return sliceTable.parent }

func (sliceTable *OpticalInterface_ConfigStatus_SliceTables_SliceTable) GetParentYangName() string { return "slice-tables" }

// OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr
// The bag containing slice config status
type OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr struct {
    parent types.Entity
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

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetFilter() yfilter.YFilter { return sliceStatusAttr.YFilter }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) SetFilter(yf yfilter.YFilter) { sliceStatusAttr.YFilter = yf }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetGoName(yname string) string {
    if yname == "slice" { return "Slice" }
    if yname == "prov-status" { return "ProvStatus" }
    if yname == "present-config" { return "PresentConfig" }
    if yname == "present-timestamp" { return "PresentTimestamp" }
    if yname == "past-config" { return "PastConfig" }
    if yname == "past-timestamp" { return "PastTimestamp" }
    if yname == "err-str" { return "ErrStr" }
    if yname == "err-timestamp" { return "ErrTimestamp" }
    return ""
}

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetSegmentPath() string {
    return "slice-status-attr"
}

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice"] = sliceStatusAttr.Slice
    leafs["prov-status"] = sliceStatusAttr.ProvStatus
    leafs["present-config"] = sliceStatusAttr.PresentConfig
    leafs["present-timestamp"] = sliceStatusAttr.PresentTimestamp
    leafs["past-config"] = sliceStatusAttr.PastConfig
    leafs["past-timestamp"] = sliceStatusAttr.PastTimestamp
    leafs["err-str"] = sliceStatusAttr.ErrStr
    leafs["err-timestamp"] = sliceStatusAttr.ErrTimestamp
    return leafs
}

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetBundleName() string { return "cisco_ios_xr" }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetYangName() string { return "slice-status-attr" }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) SetParent(parent types.Entity) { sliceStatusAttr.parent = parent }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetParent() types.Entity { return sliceStatusAttr.parent }

func (sliceStatusAttr *OpticalInterface_ConfigStatus_SliceTables_SliceTable_SliceStatusAttr) GetParentYangName() string { return "slice-table" }

// OpticalInterface_OpticalChannelInterfaces
// The operational attributes for a particular
// optical channel
type OpticalInterface_OpticalChannelInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The operational attributes for an optical channel. The type is slice of
    // OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface.
    OpticalChannelInterface []OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface
}

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetFilter() yfilter.YFilter { return opticalChannelInterfaces.YFilter }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) SetFilter(yf yfilter.YFilter) { opticalChannelInterfaces.YFilter = yf }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetGoName(yname string) string {
    if yname == "optical-channel-interface" { return "OpticalChannelInterface" }
    return ""
}

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetSegmentPath() string {
    return "optical-channel-interfaces"
}

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-channel-interface" {
        for _, c := range opticalChannelInterfaces.OpticalChannelInterface {
            if opticalChannelInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface{}
        opticalChannelInterfaces.OpticalChannelInterface = append(opticalChannelInterfaces.OpticalChannelInterface, child)
        return &opticalChannelInterfaces.OpticalChannelInterface[len(opticalChannelInterfaces.OpticalChannelInterface)-1]
    }
    return nil
}

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticalChannelInterfaces.OpticalChannelInterface {
        children[opticalChannelInterfaces.OpticalChannelInterface[i].GetSegmentPath()] = &opticalChannelInterfaces.OpticalChannelInterface[i]
    }
    return children
}

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetYangName() string { return "optical-channel-interfaces" }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) SetParent(parent types.Entity) { opticalChannelInterfaces.parent = parent }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetParent() types.Entity { return opticalChannelInterfaces.parent }

func (opticalChannelInterfaces *OpticalInterface_OpticalChannelInterfaces) GetParentYangName() string { return "optical-interface" }

// OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface
// The operational attributes for an optical
// channel
type OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the optical-channel. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Location interface{}

    // The operational attributes for an optical channel.
    OpticalChannelInterfaceAttr OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr
}

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetFilter() yfilter.YFilter { return opticalChannelInterface.YFilter }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) SetFilter(yf yfilter.YFilter) { opticalChannelInterface.YFilter = yf }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "optical-channel-interface-attr" { return "OpticalChannelInterfaceAttr" }
    return ""
}

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetSegmentPath() string {
    return "optical-channel-interface" + "[location='" + fmt.Sprintf("%v", opticalChannelInterface.Location) + "']"
}

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-channel-interface-attr" {
        return &opticalChannelInterface.OpticalChannelInterfaceAttr
    }
    return nil
}

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-channel-interface-attr"] = &opticalChannelInterface.OpticalChannelInterfaceAttr
    return children
}

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = opticalChannelInterface.Location
    return leafs
}

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetBundleName() string { return "cisco_ios_xr" }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetYangName() string { return "optical-channel-interface" }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) SetParent(parent types.Entity) { opticalChannelInterface.parent = parent }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetParent() types.Entity { return opticalChannelInterface.parent }

func (opticalChannelInterface *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface) GetParentYangName() string { return "optical-channel-interfaces" }

// OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr
// The operational attributes for an optical
// channel
type OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr struct {
    parent types.Entity
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

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetFilter() yfilter.YFilter { return opticalChannelInterfaceAttr.YFilter }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) SetFilter(yf yfilter.YFilter) { opticalChannelInterfaceAttr.YFilter = yf }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "index" { return "Index" }
    if yname == "frequency" { return "Frequency" }
    if yname == "power" { return "Power" }
    if yname == "oper-mode" { return "OperMode" }
    if yname == "line-port" { return "LinePort" }
    return ""
}

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetSegmentPath() string {
    return "optical-channel-interface-attr"
}

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = opticalChannelInterfaceAttr.Name
    leafs["index"] = opticalChannelInterfaceAttr.Index
    leafs["frequency"] = opticalChannelInterfaceAttr.Frequency
    leafs["power"] = opticalChannelInterfaceAttr.Power
    leafs["oper-mode"] = opticalChannelInterfaceAttr.OperMode
    leafs["line-port"] = opticalChannelInterfaceAttr.LinePort
    return leafs
}

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetBundleName() string { return "cisco_ios_xr" }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetYangName() string { return "optical-channel-interface-attr" }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) SetParent(parent types.Entity) { opticalChannelInterfaceAttr.parent = parent }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetParent() types.Entity { return opticalChannelInterfaceAttr.parent }

func (opticalChannelInterfaceAttr *OpticalInterface_OpticalChannelInterfaces_OpticalChannelInterface_OpticalChannelInterfaceAttr) GetParentYangName() string { return "optical-channel-interface" }

// OpticalInterface_Graph
// Table containing Graph Structure and related
// info
type OpticalInterface_Graph struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The path containg file which has adjacency list stored.
    AdjListPath OpticalInterface_Graph_AdjListPath

    // The path containg file which has graph structure stored.
    GraphStructurePath OpticalInterface_Graph_GraphStructurePath
}

func (graph *OpticalInterface_Graph) GetFilter() yfilter.YFilter { return graph.YFilter }

func (graph *OpticalInterface_Graph) SetFilter(yf yfilter.YFilter) { graph.YFilter = yf }

func (graph *OpticalInterface_Graph) GetGoName(yname string) string {
    if yname == "adj-list-path" { return "AdjListPath" }
    if yname == "graph-structure-path" { return "GraphStructurePath" }
    return ""
}

func (graph *OpticalInterface_Graph) GetSegmentPath() string {
    return "graph"
}

func (graph *OpticalInterface_Graph) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "adj-list-path" {
        return &graph.AdjListPath
    }
    if childYangName == "graph-structure-path" {
        return &graph.GraphStructurePath
    }
    return nil
}

func (graph *OpticalInterface_Graph) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["adj-list-path"] = &graph.AdjListPath
    children["graph-structure-path"] = &graph.GraphStructurePath
    return children
}

func (graph *OpticalInterface_Graph) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (graph *OpticalInterface_Graph) GetBundleName() string { return "cisco_ios_xr" }

func (graph *OpticalInterface_Graph) GetYangName() string { return "graph" }

func (graph *OpticalInterface_Graph) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (graph *OpticalInterface_Graph) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (graph *OpticalInterface_Graph) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (graph *OpticalInterface_Graph) SetParent(parent types.Entity) { graph.parent = parent }

func (graph *OpticalInterface_Graph) GetParent() types.Entity { return graph.parent }

func (graph *OpticalInterface_Graph) GetParentYangName() string { return "optical-interface" }

// OpticalInterface_Graph_AdjListPath
// The path containg file which has adjacency list
// stored
type OpticalInterface_Graph_AdjListPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path. The type is string with length: 0..256.
    Path interface{}
}

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetFilter() yfilter.YFilter { return adjListPath.YFilter }

func (adjListPath *OpticalInterface_Graph_AdjListPath) SetFilter(yf yfilter.YFilter) { adjListPath.YFilter = yf }

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetSegmentPath() string {
    return "adj-list-path"
}

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = adjListPath.Path
    return leafs
}

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetBundleName() string { return "cisco_ios_xr" }

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetYangName() string { return "adj-list-path" }

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjListPath *OpticalInterface_Graph_AdjListPath) SetParent(parent types.Entity) { adjListPath.parent = parent }

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetParent() types.Entity { return adjListPath.parent }

func (adjListPath *OpticalInterface_Graph_AdjListPath) GetParentYangName() string { return "graph" }

// OpticalInterface_Graph_GraphStructurePath
// The path containg file which has graph
// structure stored
type OpticalInterface_Graph_GraphStructurePath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path. The type is string with length: 0..256.
    Path interface{}
}

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetFilter() yfilter.YFilter { return graphStructurePath.YFilter }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) SetFilter(yf yfilter.YFilter) { graphStructurePath.YFilter = yf }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetSegmentPath() string {
    return "graph-structure-path"
}

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = graphStructurePath.Path
    return leafs
}

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetBundleName() string { return "cisco_ios_xr" }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetYangName() string { return "graph-structure-path" }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) SetParent(parent types.Entity) { graphStructurePath.parent = parent }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetParent() types.Entity { return graphStructurePath.parent }

func (graphStructurePath *OpticalInterface_Graph_GraphStructurePath) GetParentYangName() string { return "graph" }

// OpticalInterface_OperationalModes
// The Operational Mode Table
type OpticalInterface_OperationalModes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mode supported on Device. The type is slice of
    // OpticalInterface_OperationalModes_OperationalMode.
    OperationalMode []OpticalInterface_OperationalModes_OperationalMode
}

func (operationalModes *OpticalInterface_OperationalModes) GetFilter() yfilter.YFilter { return operationalModes.YFilter }

func (operationalModes *OpticalInterface_OperationalModes) SetFilter(yf yfilter.YFilter) { operationalModes.YFilter = yf }

func (operationalModes *OpticalInterface_OperationalModes) GetGoName(yname string) string {
    if yname == "operational-mode" { return "OperationalMode" }
    return ""
}

func (operationalModes *OpticalInterface_OperationalModes) GetSegmentPath() string {
    return "operational-modes"
}

func (operationalModes *OpticalInterface_OperationalModes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operational-mode" {
        for _, c := range operationalModes.OperationalMode {
            if operationalModes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalInterface_OperationalModes_OperationalMode{}
        operationalModes.OperationalMode = append(operationalModes.OperationalMode, child)
        return &operationalModes.OperationalMode[len(operationalModes.OperationalMode)-1]
    }
    return nil
}

func (operationalModes *OpticalInterface_OperationalModes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range operationalModes.OperationalMode {
        children[operationalModes.OperationalMode[i].GetSegmentPath()] = &operationalModes.OperationalMode[i]
    }
    return children
}

func (operationalModes *OpticalInterface_OperationalModes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operationalModes *OpticalInterface_OperationalModes) GetBundleName() string { return "cisco_ios_xr" }

func (operationalModes *OpticalInterface_OperationalModes) GetYangName() string { return "operational-modes" }

func (operationalModes *OpticalInterface_OperationalModes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationalModes *OpticalInterface_OperationalModes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationalModes *OpticalInterface_OperationalModes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationalModes *OpticalInterface_OperationalModes) SetParent(parent types.Entity) { operationalModes.parent = parent }

func (operationalModes *OpticalInterface_OperationalModes) GetParent() types.Entity { return operationalModes.parent }

func (operationalModes *OpticalInterface_OperationalModes) GetParentYangName() string { return "optical-interface" }

// OpticalInterface_OperationalModes_OperationalMode
// Mode supported on Device
type OpticalInterface_OperationalModes_OperationalMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Mode-id for supported mode on Device. The type is
    // interface{} with range: -2147483648..2147483647.
    ModeId interface{}

    // The operational attributes for mxp driver fec-mode.
    OperationalModeAttributes OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes
}

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetFilter() yfilter.YFilter { return operationalMode.YFilter }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) SetFilter(yf yfilter.YFilter) { operationalMode.YFilter = yf }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetGoName(yname string) string {
    if yname == "mode-id" { return "ModeId" }
    if yname == "operational-mode-attributes" { return "OperationalModeAttributes" }
    return ""
}

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetSegmentPath() string {
    return "operational-mode" + "[mode-id='" + fmt.Sprintf("%v", operationalMode.ModeId) + "']"
}

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operational-mode-attributes" {
        return &operationalMode.OperationalModeAttributes
    }
    return nil
}

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["operational-mode-attributes"] = &operationalMode.OperationalModeAttributes
    return children
}

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode-id"] = operationalMode.ModeId
    return leafs
}

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetBundleName() string { return "cisco_ios_xr" }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetYangName() string { return "operational-mode" }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) SetParent(parent types.Entity) { operationalMode.parent = parent }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetParent() types.Entity { return operationalMode.parent }

func (operationalMode *OpticalInterface_OperationalModes_OperationalMode) GetParentYangName() string { return "operational-modes" }

// OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes
// The operational attributes for mxp driver
// fec-mode
type OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Description. The type is string with length: 0..128.
    Description interface{}

    // VendorId. The type is string with length: 0..64.
    VendorId interface{}
}

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetFilter() yfilter.YFilter { return operationalModeAttributes.YFilter }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) SetFilter(yf yfilter.YFilter) { operationalModeAttributes.YFilter = yf }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "vendor-id" { return "VendorId" }
    return ""
}

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetSegmentPath() string {
    return "operational-mode-attributes"
}

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = operationalModeAttributes.Description
    leafs["vendor-id"] = operationalModeAttributes.VendorId
    return leafs
}

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetYangName() string { return "operational-mode-attributes" }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) SetParent(parent types.Entity) { operationalModeAttributes.parent = parent }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetParent() types.Entity { return operationalModeAttributes.parent }

func (operationalModeAttributes *OpticalInterface_OperationalModes_OperationalMode_OperationalModeAttributes) GetParentYangName() string { return "operational-mode" }

// OpticalInterface_OpticalLogicalInterfaces
// The operational attributes for a logical channel
type OpticalInterface_OpticalLogicalInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The operational attributes for a logical channel. The type is slice of
    // OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface.
    OpticalLogicalInterface []OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface
}

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetFilter() yfilter.YFilter { return opticalLogicalInterfaces.YFilter }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) SetFilter(yf yfilter.YFilter) { opticalLogicalInterfaces.YFilter = yf }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetGoName(yname string) string {
    if yname == "optical-logical-interface" { return "OpticalLogicalInterface" }
    return ""
}

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetSegmentPath() string {
    return "optical-logical-interfaces"
}

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-logical-interface" {
        for _, c := range opticalLogicalInterfaces.OpticalLogicalInterface {
            if opticalLogicalInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface{}
        opticalLogicalInterfaces.OpticalLogicalInterface = append(opticalLogicalInterfaces.OpticalLogicalInterface, child)
        return &opticalLogicalInterfaces.OpticalLogicalInterface[len(opticalLogicalInterfaces.OpticalLogicalInterface)-1]
    }
    return nil
}

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticalLogicalInterfaces.OpticalLogicalInterface {
        children[opticalLogicalInterfaces.OpticalLogicalInterface[i].GetSegmentPath()] = &opticalLogicalInterfaces.OpticalLogicalInterface[i]
    }
    return children
}

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetYangName() string { return "optical-logical-interfaces" }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) SetParent(parent types.Entity) { opticalLogicalInterfaces.parent = parent }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetParent() types.Entity { return opticalLogicalInterfaces.parent }

func (opticalLogicalInterfaces *OpticalInterface_OpticalLogicalInterfaces) GetParentYangName() string { return "optical-interface" }

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface
// The operational attributes for a logical
// channel
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the logical-channel. The type is
    // interface{} with range: -2147483648..2147483647.
    Index interface{}

    // The operational attributes for a particular logical channel.
    OpticalLogicalInterfaceAttr OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr

    // The operational attributes for a particular interface.
    OpticalLogicalInterfaceLogicalChannelAssignments OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments
}

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetFilter() yfilter.YFilter { return opticalLogicalInterface.YFilter }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) SetFilter(yf yfilter.YFilter) { opticalLogicalInterface.YFilter = yf }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "optical-logical-interface-attr" { return "OpticalLogicalInterfaceAttr" }
    if yname == "optical-logical-interface-logical-channel-assignments" { return "OpticalLogicalInterfaceLogicalChannelAssignments" }
    return ""
}

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetSegmentPath() string {
    return "optical-logical-interface" + "[index='" + fmt.Sprintf("%v", opticalLogicalInterface.Index) + "']"
}

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-logical-interface-attr" {
        return &opticalLogicalInterface.OpticalLogicalInterfaceAttr
    }
    if childYangName == "optical-logical-interface-logical-channel-assignments" {
        return &opticalLogicalInterface.OpticalLogicalInterfaceLogicalChannelAssignments
    }
    return nil
}

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-logical-interface-attr"] = &opticalLogicalInterface.OpticalLogicalInterfaceAttr
    children["optical-logical-interface-logical-channel-assignments"] = &opticalLogicalInterface.OpticalLogicalInterfaceLogicalChannelAssignments
    return children
}

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = opticalLogicalInterface.Index
    return leafs
}

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetBundleName() string { return "cisco_ios_xr" }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetYangName() string { return "optical-logical-interface" }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) SetParent(parent types.Entity) { opticalLogicalInterface.parent = parent }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetParent() types.Entity { return opticalLogicalInterface.parent }

func (opticalLogicalInterface *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface) GetParentYangName() string { return "optical-logical-interfaces" }

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr
// The operational attributes for a particular
// logical channel
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr struct {
    parent types.Entity
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

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetFilter() yfilter.YFilter { return opticalLogicalInterfaceAttr.YFilter }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) SetFilter(yf yfilter.YFilter) { opticalLogicalInterfaceAttr.YFilter = yf }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetGoName(yname string) string {
    if yname == "logical-channel-index" { return "LogicalChannelIndex" }
    if yname == "logical-channel-ifname" { return "LogicalChannelIfname" }
    if yname == "type" { return "Type" }
    if yname == "trib-rate-class" { return "TribRateClass" }
    if yname == "trib-protocol" { return "TribProtocol" }
    if yname == "protocol-type" { return "ProtocolType" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    if yname == "ingress-client-port" { return "IngressClientPort" }
    if yname == "ingress-physical-channel" { return "IngressPhysicalChannel" }
    if yname == "tti-transmit" { return "TtiTransmit" }
    if yname == "tti-expected" { return "TtiExpected" }
    return ""
}

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetSegmentPath() string {
    return "optical-logical-interface-attr"
}

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logical-channel-index"] = opticalLogicalInterfaceAttr.LogicalChannelIndex
    leafs["logical-channel-ifname"] = opticalLogicalInterfaceAttr.LogicalChannelIfname
    leafs["type"] = opticalLogicalInterfaceAttr.Type
    leafs["trib-rate-class"] = opticalLogicalInterfaceAttr.TribRateClass
    leafs["trib-protocol"] = opticalLogicalInterfaceAttr.TribProtocol
    leafs["protocol-type"] = opticalLogicalInterfaceAttr.ProtocolType
    leafs["admin-state"] = opticalLogicalInterfaceAttr.AdminState
    leafs["loopback-mode"] = opticalLogicalInterfaceAttr.LoopbackMode
    leafs["ingress-client-port"] = opticalLogicalInterfaceAttr.IngressClientPort
    leafs["ingress-physical-channel"] = opticalLogicalInterfaceAttr.IngressPhysicalChannel
    leafs["tti-transmit"] = opticalLogicalInterfaceAttr.TtiTransmit
    leafs["tti-expected"] = opticalLogicalInterfaceAttr.TtiExpected
    return leafs
}

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetBundleName() string { return "cisco_ios_xr" }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetYangName() string { return "optical-logical-interface-attr" }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) SetParent(parent types.Entity) { opticalLogicalInterfaceAttr.parent = parent }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetParent() types.Entity { return opticalLogicalInterfaceAttr.parent }

func (opticalLogicalInterfaceAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceAttr) GetParentYangName() string { return "optical-logical-interface" }

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments
// The operational attributes for a particular
// interface
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The operational attributes for a logical channel assignment. The type is
    // slice of
    // OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment.
    OpticalLogicalInterfaceLogicalChannelAssignment []OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment
}

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetFilter() yfilter.YFilter { return opticalLogicalInterfaceLogicalChannelAssignments.YFilter }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) SetFilter(yf yfilter.YFilter) { opticalLogicalInterfaceLogicalChannelAssignments.YFilter = yf }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetGoName(yname string) string {
    if yname == "optical-logical-interface-logical-channel-assignment" { return "OpticalLogicalInterfaceLogicalChannelAssignment" }
    return ""
}

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetSegmentPath() string {
    return "optical-logical-interface-logical-channel-assignments"
}

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-logical-interface-logical-channel-assignment" {
        for _, c := range opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment {
            if opticalLogicalInterfaceLogicalChannelAssignments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment{}
        opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment = append(opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment, child)
        return &opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment[len(opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment)-1]
    }
    return nil
}

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment {
        children[opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment[i].GetSegmentPath()] = &opticalLogicalInterfaceLogicalChannelAssignments.OpticalLogicalInterfaceLogicalChannelAssignment[i]
    }
    return children
}

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetBundleName() string { return "cisco_ios_xr" }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetYangName() string { return "optical-logical-interface-logical-channel-assignments" }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) SetParent(parent types.Entity) { opticalLogicalInterfaceLogicalChannelAssignments.parent = parent }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetParent() types.Entity { return opticalLogicalInterfaceLogicalChannelAssignments.parent }

func (opticalLogicalInterfaceLogicalChannelAssignments *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments) GetParentYangName() string { return "optical-logical-interface" }

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment
// The operational attributes for a logical
// channel assignment
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the logical-channel. The type is
    // interface{} with range: -2147483648..2147483647.
    Index interface{}

    // The operational attributes for a logical channel assignment.
    OpticalLogicalInterfaceLogicalChannelAssignmentAttr OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr
}

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetFilter() yfilter.YFilter { return opticalLogicalInterfaceLogicalChannelAssignment.YFilter }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) SetFilter(yf yfilter.YFilter) { opticalLogicalInterfaceLogicalChannelAssignment.YFilter = yf }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "optical-logical-interface-logical-channel-assignment-attr" { return "OpticalLogicalInterfaceLogicalChannelAssignmentAttr" }
    return ""
}

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetSegmentPath() string {
    return "optical-logical-interface-logical-channel-assignment" + "[index='" + fmt.Sprintf("%v", opticalLogicalInterfaceLogicalChannelAssignment.Index) + "']"
}

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-logical-interface-logical-channel-assignment-attr" {
        return &opticalLogicalInterfaceLogicalChannelAssignment.OpticalLogicalInterfaceLogicalChannelAssignmentAttr
    }
    return nil
}

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optical-logical-interface-logical-channel-assignment-attr"] = &opticalLogicalInterfaceLogicalChannelAssignment.OpticalLogicalInterfaceLogicalChannelAssignmentAttr
    return children
}

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = opticalLogicalInterfaceLogicalChannelAssignment.Index
    return leafs
}

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetBundleName() string { return "cisco_ios_xr" }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetYangName() string { return "optical-logical-interface-logical-channel-assignment" }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) SetParent(parent types.Entity) { opticalLogicalInterfaceLogicalChannelAssignment.parent = parent }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetParent() types.Entity { return opticalLogicalInterfaceLogicalChannelAssignment.parent }

func (opticalLogicalInterfaceLogicalChannelAssignment *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment) GetParentYangName() string { return "optical-logical-interface-logical-channel-assignments" }

// OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr
// The operational attributes for a logical
// channel assignment
type OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr struct {
    parent types.Entity
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

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetFilter() yfilter.YFilter { return opticalLogicalInterfaceLogicalChannelAssignmentAttr.YFilter }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) SetFilter(yf yfilter.YFilter) { opticalLogicalInterfaceLogicalChannelAssignmentAttr.YFilter = yf }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "name" { return "Name" }
    if yname == "is-logical-link" { return "IsLogicalLink" }
    if yname == "logical-channel" { return "LogicalChannel" }
    if yname == "optical-channel" { return "OpticalChannel" }
    if yname == "allocation" { return "Allocation" }
    if yname == "assignment-type" { return "AssignmentType" }
    return ""
}

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetSegmentPath() string {
    return "optical-logical-interface-logical-channel-assignment-attr"
}

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = opticalLogicalInterfaceLogicalChannelAssignmentAttr.Index
    leafs["name"] = opticalLogicalInterfaceLogicalChannelAssignmentAttr.Name
    leafs["is-logical-link"] = opticalLogicalInterfaceLogicalChannelAssignmentAttr.IsLogicalLink
    leafs["logical-channel"] = opticalLogicalInterfaceLogicalChannelAssignmentAttr.LogicalChannel
    leafs["optical-channel"] = opticalLogicalInterfaceLogicalChannelAssignmentAttr.OpticalChannel
    leafs["allocation"] = opticalLogicalInterfaceLogicalChannelAssignmentAttr.Allocation
    leafs["assignment-type"] = opticalLogicalInterfaceLogicalChannelAssignmentAttr.AssignmentType
    return leafs
}

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetBundleName() string { return "cisco_ios_xr" }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetYangName() string { return "optical-logical-interface-logical-channel-assignment-attr" }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) SetParent(parent types.Entity) { opticalLogicalInterfaceLogicalChannelAssignmentAttr.parent = parent }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetParent() types.Entity { return opticalLogicalInterfaceLogicalChannelAssignmentAttr.parent }

func (opticalLogicalInterfaceLogicalChannelAssignmentAttr *OpticalInterface_OpticalLogicalInterfaces_OpticalLogicalInterface_OpticalLogicalInterfaceLogicalChannelAssignments_OpticalLogicalInterfaceLogicalChannelAssignment_OpticalLogicalInterfaceLogicalChannelAssignmentAttr) GetParentYangName() string { return "optical-logical-interface-logical-channel-assignment" }

