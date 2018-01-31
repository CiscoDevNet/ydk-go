// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: NCS1k HW module config
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_mxp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_mxp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-mxp-cfg hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-mxp-cfg:hardware-module", reflect.TypeOf(HardwareModule{}))
}

// ClientDataRate represents Client data rate
type ClientDataRate string

const (
    // TenGig
    ClientDataRate_ten_gig ClientDataRate = "ten-gig"

    // FortyGig
    ClientDataRate_forty_gig ClientDataRate = "forty-gig"

    // HundredGig
    ClientDataRate_hundred_gig ClientDataRate = "hundred-gig"

    // TenAndHundredGig
    ClientDataRate_ten_and_hundred_gig ClientDataRate = "ten-and-hundred-gig"
)

// TrunkDataRate represents Trunk data rate
type TrunkDataRate string

const (
    // HundredGig
    TrunkDataRate_hundred_gig TrunkDataRate = "hundred-gig"

    // TwoHundredGig
    TrunkDataRate_two_hundred_gig TrunkDataRate = "two-hundred-gig"

    // TwoHundredFiftyGig
    TrunkDataRate_two_hundred_fifty_gig TrunkDataRate = "two-hundred-fifty-gig"
)

// Fec represents Fec
type Fec string

const (
    // SoftDecision7
    Fec_sd7 Fec = "sd7"

    // SoftDecision20
    Fec_sd20 Fec = "sd20"
)

// HardwareModule
// NCS1k HW module config
type HardwareModule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is slice of HardwareModule_Node.
    Node []HardwareModule_Node
}

func (hardwareModule *HardwareModule) GetFilter() yfilter.YFilter { return hardwareModule.YFilter }

func (hardwareModule *HardwareModule) SetFilter(yf yfilter.YFilter) { hardwareModule.YFilter = yf }

func (hardwareModule *HardwareModule) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (hardwareModule *HardwareModule) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs1k-mxp-cfg:hardware-module"
}

func (hardwareModule *HardwareModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range hardwareModule.Node {
            if hardwareModule.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Node{}
        hardwareModule.Node = append(hardwareModule.Node, child)
        return &hardwareModule.Node[len(hardwareModule.Node)-1]
    }
    return nil
}

func (hardwareModule *HardwareModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareModule.Node {
        children[hardwareModule.Node[i].GetSegmentPath()] = &hardwareModule.Node[i]
    }
    return children
}

func (hardwareModule *HardwareModule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModule *HardwareModule) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModule *HardwareModule) GetYangName() string { return "hardware-module" }

func (hardwareModule *HardwareModule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModule *HardwareModule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModule *HardwareModule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModule *HardwareModule) SetParent(parent types.Entity) { hardwareModule.parent = parent }

func (hardwareModule *HardwareModule) GetParent() types.Entity { return hardwareModule.parent }

func (hardwareModule *HardwareModule) GetParentYangName() string { return "Cisco-IOS-XR-ncs1k-mxp-cfg" }

// HardwareModule_Node
// Node
type HardwareModule_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fully qualified line card specification. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Location interface{}

    // Slice to be Provisioned. The type is slice of HardwareModule_Node_Slice.
    Slice []HardwareModule_Node_Slice
}

func (node *HardwareModule_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModule_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModule_Node) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "slice" { return "Slice" }
    return ""
}

func (node *HardwareModule_Node) GetSegmentPath() string {
    return "node" + "[location='" + fmt.Sprintf("%v", node.Location) + "']"
}

func (node *HardwareModule_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice" {
        for _, c := range node.Slice {
            if node.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Node_Slice{}
        node.Slice = append(node.Slice, child)
        return &node.Slice[len(node.Slice)-1]
    }
    return nil
}

func (node *HardwareModule_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range node.Slice {
        children[node.Slice[i].GetSegmentPath()] = &node.Slice[i]
    }
    return children
}

func (node *HardwareModule_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = node.Location
    return leafs
}

func (node *HardwareModule_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModule_Node) GetYangName() string { return "node" }

func (node *HardwareModule_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModule_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModule_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModule_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModule_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModule_Node) GetParentYangName() string { return "hardware-module" }

// HardwareModule_Node_Slice
// Slice to be Provisioned
type HardwareModule_Node_Slice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set Slice. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SliceId interface{}

    // Drop LLDP Packets. The type is bool.
    Lldp interface{}

    // Data rates & FEC.
    Values HardwareModule_Node_Slice_Values
}

func (slice *HardwareModule_Node_Slice) GetFilter() yfilter.YFilter { return slice.YFilter }

func (slice *HardwareModule_Node_Slice) SetFilter(yf yfilter.YFilter) { slice.YFilter = yf }

func (slice *HardwareModule_Node_Slice) GetGoName(yname string) string {
    if yname == "slice-id" { return "SliceId" }
    if yname == "lldp" { return "Lldp" }
    if yname == "values" { return "Values" }
    return ""
}

func (slice *HardwareModule_Node_Slice) GetSegmentPath() string {
    return "slice" + "[slice-id='" + fmt.Sprintf("%v", slice.SliceId) + "']"
}

func (slice *HardwareModule_Node_Slice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "values" {
        return &slice.Values
    }
    return nil
}

func (slice *HardwareModule_Node_Slice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["values"] = &slice.Values
    return children
}

func (slice *HardwareModule_Node_Slice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice-id"] = slice.SliceId
    leafs["lldp"] = slice.Lldp
    return leafs
}

func (slice *HardwareModule_Node_Slice) GetBundleName() string { return "cisco_ios_xr" }

func (slice *HardwareModule_Node_Slice) GetYangName() string { return "slice" }

func (slice *HardwareModule_Node_Slice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slice *HardwareModule_Node_Slice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slice *HardwareModule_Node_Slice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slice *HardwareModule_Node_Slice) SetParent(parent types.Entity) { slice.parent = parent }

func (slice *HardwareModule_Node_Slice) GetParent() types.Entity { return slice.parent }

func (slice *HardwareModule_Node_Slice) GetParentYangName() string { return "node" }

// HardwareModule_Node_Slice_Values
// Data rates & FEC
type HardwareModule_Node_Slice_Values struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Client Rate. The type is ClientDataRate.
    ClientRate interface{}

    // TrunkRate. The type is TrunkDataRate.
    TrunkRate interface{}

    // FEC. The type is Fec.
    Fec interface{}

    // Encrypted. The type is bool. The default value is false.
    Encrypted interface{}
}

func (values *HardwareModule_Node_Slice_Values) GetFilter() yfilter.YFilter { return values.YFilter }

func (values *HardwareModule_Node_Slice_Values) SetFilter(yf yfilter.YFilter) { values.YFilter = yf }

func (values *HardwareModule_Node_Slice_Values) GetGoName(yname string) string {
    if yname == "client-rate" { return "ClientRate" }
    if yname == "trunk-rate" { return "TrunkRate" }
    if yname == "fec" { return "Fec" }
    if yname == "encrypted" { return "Encrypted" }
    return ""
}

func (values *HardwareModule_Node_Slice_Values) GetSegmentPath() string {
    return "values"
}

func (values *HardwareModule_Node_Slice_Values) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (values *HardwareModule_Node_Slice_Values) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (values *HardwareModule_Node_Slice_Values) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-rate"] = values.ClientRate
    leafs["trunk-rate"] = values.TrunkRate
    leafs["fec"] = values.Fec
    leafs["encrypted"] = values.Encrypted
    return leafs
}

func (values *HardwareModule_Node_Slice_Values) GetBundleName() string { return "cisco_ios_xr" }

func (values *HardwareModule_Node_Slice_Values) GetYangName() string { return "values" }

func (values *HardwareModule_Node_Slice_Values) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (values *HardwareModule_Node_Slice_Values) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (values *HardwareModule_Node_Slice_Values) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (values *HardwareModule_Node_Slice_Values) SetParent(parent types.Entity) { values.parent = parent }

func (values *HardwareModule_Node_Slice_Values) GetParent() types.Entity { return values.parent }

func (values *HardwareModule_Node_Slice_Values) GetParentYangName() string { return "slice" }

