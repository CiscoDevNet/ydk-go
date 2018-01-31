// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-pl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec-platform: MACSec operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_macsec_pl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_macsec_pl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-macsec-pl-oper macsec-platform}", reflect.TypeOf(MacsecPlatform{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-macsec-pl-oper:macsec-platform", reflect.TypeOf(MacsecPlatform{}))
}

// MacsecCard represents Macsec card
type MacsecCard string

const (
    // macsec none
    MacsecCard_macsec_none MacsecCard = "macsec-none"

    // macsec msfpga
    MacsecCard_macsec_msfpga MacsecCard = "macsec-msfpga"

    // macsec xlmsfpga
    MacsecCard_macsec_xlmsfpga MacsecCard = "macsec-xlmsfpga"

    // macsec apm
    MacsecCard_macsec_apm MacsecCard = "macsec-apm"
)

// MacsecPlatform
// MACSec operational data
type MacsecPlatform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NodeTable for all the nodes.
    Nodes MacsecPlatform_Nodes
}

func (macsecPlatform *MacsecPlatform) GetFilter() yfilter.YFilter { return macsecPlatform.YFilter }

func (macsecPlatform *MacsecPlatform) SetFilter(yf yfilter.YFilter) { macsecPlatform.YFilter = yf }

func (macsecPlatform *MacsecPlatform) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (macsecPlatform *MacsecPlatform) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-macsec-pl-oper:macsec-platform"
}

func (macsecPlatform *MacsecPlatform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &macsecPlatform.Nodes
    }
    return nil
}

func (macsecPlatform *MacsecPlatform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &macsecPlatform.Nodes
    return children
}

func (macsecPlatform *MacsecPlatform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macsecPlatform *MacsecPlatform) GetBundleName() string { return "cisco_ios_xr" }

func (macsecPlatform *MacsecPlatform) GetYangName() string { return "macsec-platform" }

func (macsecPlatform *MacsecPlatform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecPlatform *MacsecPlatform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecPlatform *MacsecPlatform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecPlatform *MacsecPlatform) SetParent(parent types.Entity) { macsecPlatform.parent = parent }

func (macsecPlatform *MacsecPlatform) GetParent() types.Entity { return macsecPlatform.parent }

func (macsecPlatform *MacsecPlatform) GetParentYangName() string { return "Cisco-IOS-XR-crypto-macsec-pl-oper" }

// MacsecPlatform_Nodes
// NodeTable for all the nodes
type MacsecPlatform_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node where macsec interfaces exist. The type is slice of
    // MacsecPlatform_Nodes_Node.
    Node []MacsecPlatform_Nodes_Node
}

func (nodes *MacsecPlatform_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *MacsecPlatform_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *MacsecPlatform_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *MacsecPlatform_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *MacsecPlatform_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecPlatform_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *MacsecPlatform_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *MacsecPlatform_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *MacsecPlatform_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *MacsecPlatform_Nodes) GetYangName() string { return "nodes" }

func (nodes *MacsecPlatform_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *MacsecPlatform_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *MacsecPlatform_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *MacsecPlatform_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *MacsecPlatform_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *MacsecPlatform_Nodes) GetParentYangName() string { return "macsec-platform" }

// MacsecPlatform_Nodes_Node
// Node where macsec interfaces exist
type MacsecPlatform_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table of Interfaces.
    Interfaces MacsecPlatform_Nodes_Node_Interfaces
}

func (node *MacsecPlatform_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *MacsecPlatform_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *MacsecPlatform_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (node *MacsecPlatform_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *MacsecPlatform_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    return nil
}

func (node *MacsecPlatform_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &node.Interfaces
    return children
}

func (node *MacsecPlatform_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *MacsecPlatform_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *MacsecPlatform_Nodes_Node) GetYangName() string { return "node" }

func (node *MacsecPlatform_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *MacsecPlatform_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *MacsecPlatform_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *MacsecPlatform_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *MacsecPlatform_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *MacsecPlatform_Nodes_Node) GetParentYangName() string { return "nodes" }

// MacsecPlatform_Nodes_Node_Interfaces
// Table of Interfaces
type MacsecPlatform_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Where Macsec is configured. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface.
    Interface []MacsecPlatform_Nodes_Node_Interfaces_Interface
}

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecPlatform_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface
// Interface Where Macsec is configured
type MacsecPlatform_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Value. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // The Hardware Statistics.
    HwStatistics MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics

    // Table of Hardware SAs.
    HwSas MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas

    // Table of Hardware Flows.
    HwFlowS MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS

    // The Software Statistics.
    SwStatistics MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics
}

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "hw-statistics" { return "HwStatistics" }
    if yname == "hw-sas" { return "HwSas" }
    if yname == "hw-flow-s" { return "HwFlowS" }
    if yname == "sw-statistics" { return "SwStatistics" }
    return ""
}

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-statistics" {
        return &self.HwStatistics
    }
    if childYangName == "hw-sas" {
        return &self.HwSas
    }
    if childYangName == "hw-flow-s" {
        return &self.HwFlowS
    }
    if childYangName == "sw-statistics" {
        return &self.SwStatistics
    }
    return nil
}

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-statistics"] = &self.HwStatistics
    children["hw-sas"] = &self.HwSas
    children["hw-flow-s"] = &self.HwFlowS
    children["sw-statistics"] = &self.SwStatistics
    return children
}

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    return leafs
}

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics
// The Hardware Statistics
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext
}

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetFilter() yfilter.YFilter { return hwStatistics.YFilter }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) SetFilter(yf yfilter.YFilter) { hwStatistics.YFilter = yf }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetGoName(yname string) string {
    if yname == "ext" { return "Ext" }
    return ""
}

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetSegmentPath() string {
    return "hw-statistics"
}

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext" {
        return &hwStatistics.Ext
    }
    return nil
}

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ext"] = &hwStatistics.Ext
    return children
}

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetYangName() string { return "hw-statistics" }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) SetParent(parent types.Entity) { hwStatistics.parent = parent }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetParent() types.Entity { return hwStatistics.parent }

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetParentYangName() string { return "interface" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type interface{}

    // MSFPGA Stats.
    MsfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats

    // XLFPGA Stats.
    XlfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats

    // ES200 Stats.
    Es200Stats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetFilter() yfilter.YFilter { return ext.YFilter }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) SetFilter(yf yfilter.YFilter) { ext.YFilter = yf }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "msfpga-stats" { return "MsfpgaStats" }
    if yname == "xlfpga-stats" { return "XlfpgaStats" }
    if yname == "es200-stats" { return "Es200Stats" }
    return ""
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetSegmentPath() string {
    return "ext"
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msfpga-stats" {
        return &ext.MsfpgaStats
    }
    if childYangName == "xlfpga-stats" {
        return &ext.XlfpgaStats
    }
    if childYangName == "es200-stats" {
        return &ext.Es200Stats
    }
    return nil
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["msfpga-stats"] = &ext.MsfpgaStats
    children["xlfpga-stats"] = &ext.XlfpgaStats
    children["es200-stats"] = &ext.Es200Stats
    return children
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = ext.Type
    return leafs
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetBundleName() string { return "cisco_ios_xr" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetYangName() string { return "ext" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) SetParent(parent types.Entity) { ext.parent = parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetParent() types.Entity { return ext.parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetParentYangName() string { return "hw-statistics" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats
// MSFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetFilter() yfilter.YFilter { return msfpgaStats.YFilter }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) SetFilter(yf yfilter.YFilter) { msfpgaStats.YFilter = yf }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetGoName(yname string) string {
    if yname == "tx-sa-stats" { return "TxSaStats" }
    if yname == "rx-sa-stats" { return "RxSaStats" }
    if yname == "tx-interface-macsec-stats" { return "TxInterfaceMacsecStats" }
    if yname == "rx-interface-macsec-stats" { return "RxInterfaceMacsecStats" }
    return ""
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetSegmentPath() string {
    return "msfpga-stats"
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-sa-stats" {
        return &msfpgaStats.TxSaStats
    }
    if childYangName == "rx-sa-stats" {
        return &msfpgaStats.RxSaStats
    }
    if childYangName == "tx-interface-macsec-stats" {
        return &msfpgaStats.TxInterfaceMacsecStats
    }
    if childYangName == "rx-interface-macsec-stats" {
        return &msfpgaStats.RxInterfaceMacsecStats
    }
    return nil
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-sa-stats"] = &msfpgaStats.TxSaStats
    children["rx-sa-stats"] = &msfpgaStats.RxSaStats
    children["tx-interface-macsec-stats"] = &msfpgaStats.TxInterfaceMacsecStats
    children["rx-interface-macsec-stats"] = &msfpgaStats.RxInterfaceMacsecStats
    return children
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetBundleName() string { return "cisco_ios_xr" }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetYangName() string { return "msfpga-stats" }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) SetParent(parent types.Entity) { msfpgaStats.parent = parent }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetParent() types.Entity { return msfpgaStats.parent }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Pkts Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsProtected interface{}

    // Tx Pkts Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncrypted interface{}

    // Tx Octets Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsProtected interface{}

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncrypted interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetFilter() yfilter.YFilter { return txSaStats.YFilter }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) SetFilter(yf yfilter.YFilter) { txSaStats.YFilter = yf }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetGoName(yname string) string {
    if yname == "out-pkts-protected" { return "OutPktsProtected" }
    if yname == "out-pkts-encrypted" { return "OutPktsEncrypted" }
    if yname == "out-octets-protected" { return "OutOctetsProtected" }
    if yname == "out-octets-encrypted" { return "OutOctetsEncrypted" }
    return ""
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetSegmentPath() string {
    return "tx-sa-stats"
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkts-protected"] = txSaStats.OutPktsProtected
    leafs["out-pkts-encrypted"] = txSaStats.OutPktsEncrypted
    leafs["out-octets-protected"] = txSaStats.OutOctetsProtected
    leafs["out-octets-encrypted"] = txSaStats.OutOctetsEncrypted
    return leafs
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetYangName() string { return "tx-sa-stats" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) SetParent(parent types.Entity) { txSaStats.parent = parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetParent() types.Entity { return txSaStats.parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // Rx Pkts Not Valid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotValid interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsInvalid interface{}

    // Rx Pkts OK. The type is interface{} with range: 0..18446744073709551615.
    InPktsOk interface{}

    // Rx Pkts Delayed. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsDelayed interface{}

    // Rx Pkts Late. The type is interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // Rx Pkts Unchecked. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnchecked interface{}

    // Rx Octets Validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecrypted interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetFilter() yfilter.YFilter { return rxSaStats.YFilter }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) SetFilter(yf yfilter.YFilter) { rxSaStats.YFilter = yf }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetGoName(yname string) string {
    if yname == "in-pkts-unused-sa" { return "InPktsUnusedSa" }
    if yname == "in-pkts-not-using-sa" { return "InPktsNotUsingSa" }
    if yname == "in-pkts-not-valid" { return "InPktsNotValid" }
    if yname == "in-pkts-invalid" { return "InPktsInvalid" }
    if yname == "in-pkts-ok" { return "InPktsOk" }
    if yname == "in-pkts-delayed" { return "InPktsDelayed" }
    if yname == "in-pkts-late" { return "InPktsLate" }
    if yname == "in-pkts-unchecked" { return "InPktsUnchecked" }
    if yname == "in-octets-validated" { return "InOctetsValidated" }
    if yname == "in-octets-decrypted" { return "InOctetsDecrypted" }
    return ""
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetSegmentPath() string {
    return "rx-sa-stats"
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-unused-sa"] = rxSaStats.InPktsUnusedSa
    leafs["in-pkts-not-using-sa"] = rxSaStats.InPktsNotUsingSa
    leafs["in-pkts-not-valid"] = rxSaStats.InPktsNotValid
    leafs["in-pkts-invalid"] = rxSaStats.InPktsInvalid
    leafs["in-pkts-ok"] = rxSaStats.InPktsOk
    leafs["in-pkts-delayed"] = rxSaStats.InPktsDelayed
    leafs["in-pkts-late"] = rxSaStats.InPktsLate
    leafs["in-pkts-unchecked"] = rxSaStats.InPktsUnchecked
    leafs["in-octets-validated"] = rxSaStats.InOctetsValidated
    leafs["in-octets-decrypted"] = rxSaStats.InOctetsDecrypted
    return leafs
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetYangName() string { return "rx-sa-stats" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) SetParent(parent types.Entity) { rxSaStats.parent = parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetParent() types.Entity { return rxSaStats.parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUncontrolled interface{}

    // Tx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUntagged interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktTooLong interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return txInterfaceMacsecStats.YFilter }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { txInterfaceMacsecStats.YFilter = yf }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "out-pkt-uncontrolled" { return "OutPktUncontrolled" }
    if yname == "out-pkt-untagged" { return "OutPktUntagged" }
    if yname == "out-pkt-too-long" { return "OutPktTooLong" }
    return ""
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetSegmentPath() string {
    return "tx-interface-macsec-stats"
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkt-uncontrolled"] = txInterfaceMacsecStats.OutPktUncontrolled
    leafs["out-pkt-untagged"] = txInterfaceMacsecStats.OutPktUntagged
    leafs["out-pkt-too-long"] = txInterfaceMacsecStats.OutPktTooLong
    return leafs
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetYangName() string { return "tx-interface-macsec-stats" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) SetParent(parent types.Entity) { txInterfaceMacsecStats.parent = parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetParent() types.Entity { return txInterfaceMacsecStats.parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUntagged interface{}

    // Rx Pkts Notag. The type is interface{} with range: 0..18446744073709551615.
    InPktNotag interface{}

    // Rx Pkts Bad tag. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktBadTag interface{}

    // Rx Pkts No Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // Rx Pkts Unknown Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUnknownSci interface{}

    // Rx Pkts Tagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktTagged interface{}

    // Rx Pkts Over Run. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktOverrun interface{}

    // Rx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUncontrolled interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return rxInterfaceMacsecStats.YFilter }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { rxInterfaceMacsecStats.YFilter = yf }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "in-pkt-untagged" { return "InPktUntagged" }
    if yname == "in-pkt-notag" { return "InPktNotag" }
    if yname == "in-pkt-bad-tag" { return "InPktBadTag" }
    if yname == "in-pkt-no-sci" { return "InPktNoSci" }
    if yname == "in-pkt-unknown-sci" { return "InPktUnknownSci" }
    if yname == "in-pkt-tagged" { return "InPktTagged" }
    if yname == "in-pkt-overrun" { return "InPktOverrun" }
    if yname == "in-pkt-uncontrolled" { return "InPktUncontrolled" }
    return ""
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetSegmentPath() string {
    return "rx-interface-macsec-stats"
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkt-untagged"] = rxInterfaceMacsecStats.InPktUntagged
    leafs["in-pkt-notag"] = rxInterfaceMacsecStats.InPktNotag
    leafs["in-pkt-bad-tag"] = rxInterfaceMacsecStats.InPktBadTag
    leafs["in-pkt-no-sci"] = rxInterfaceMacsecStats.InPktNoSci
    leafs["in-pkt-unknown-sci"] = rxInterfaceMacsecStats.InPktUnknownSci
    leafs["in-pkt-tagged"] = rxInterfaceMacsecStats.InPktTagged
    leafs["in-pkt-overrun"] = rxInterfaceMacsecStats.InPktOverrun
    leafs["in-pkt-uncontrolled"] = rxInterfaceMacsecStats.InPktUncontrolled
    return leafs
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetYangName() string { return "rx-interface-macsec-stats" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) SetParent(parent types.Entity) { rxInterfaceMacsecStats.parent = parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetParent() types.Entity { return rxInterfaceMacsecStats.parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats
// XLFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SC and SA Level Stats.
    MacsecTxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats

    // Rx SC and SA Level Stats.
    MacsecRxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetFilter() yfilter.YFilter { return xlfpgaStats.YFilter }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) SetFilter(yf yfilter.YFilter) { xlfpgaStats.YFilter = yf }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetGoName(yname string) string {
    if yname == "macsec-tx-stats" { return "MacsecTxStats" }
    if yname == "macsec-rx-stats" { return "MacsecRxStats" }
    return ""
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetSegmentPath() string {
    return "xlfpga-stats"
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macsec-tx-stats" {
        return &xlfpgaStats.MacsecTxStats
    }
    if childYangName == "macsec-rx-stats" {
        return &xlfpgaStats.MacsecRxStats
    }
    return nil
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["macsec-tx-stats"] = &xlfpgaStats.MacsecTxStats
    children["macsec-rx-stats"] = &xlfpgaStats.MacsecRxStats
    return children
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetBundleName() string { return "cisco_ios_xr" }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetYangName() string { return "xlfpga-stats" }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) SetParent(parent types.Entity) { xlfpgaStats.parent = parent }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetParent() types.Entity { return xlfpgaStats.parent }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats
// Tx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedOctets interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    ScToolongPkts interface{}

    // Tx packets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedPkts interface{}

    // Tx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Tx Overrun Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Tx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Tx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Tx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Current Tx AN. The type is interface{} with range: 0..18446744073709551615.
    CurrentAn interface{}

    // Current Tx SA Encrypted Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    SaEncryptedPkts interface{}
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetFilter() yfilter.YFilter { return macsecTxStats.YFilter }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) SetFilter(yf yfilter.YFilter) { macsecTxStats.YFilter = yf }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetGoName(yname string) string {
    if yname == "sc-encrypted-octets" { return "ScEncryptedOctets" }
    if yname == "sc-toolong-pkts" { return "ScToolongPkts" }
    if yname == "sc-encrypted-pkts" { return "ScEncryptedPkts" }
    if yname == "sc-untagged-pkts" { return "ScUntaggedPkts" }
    if yname == "sc-overrun-pkts" { return "ScOverrunPkts" }
    if yname == "sc-bypass-pkts" { return "ScBypassPkts" }
    if yname == "sc-eapol-pkts" { return "ScEapolPkts" }
    if yname == "sc-dropped-pkts" { return "ScDroppedPkts" }
    if yname == "current-an" { return "CurrentAn" }
    if yname == "sa-encrypted-pkts" { return "SaEncryptedPkts" }
    return ""
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetSegmentPath() string {
    return "macsec-tx-stats"
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sc-encrypted-octets"] = macsecTxStats.ScEncryptedOctets
    leafs["sc-toolong-pkts"] = macsecTxStats.ScToolongPkts
    leafs["sc-encrypted-pkts"] = macsecTxStats.ScEncryptedPkts
    leafs["sc-untagged-pkts"] = macsecTxStats.ScUntaggedPkts
    leafs["sc-overrun-pkts"] = macsecTxStats.ScOverrunPkts
    leafs["sc-bypass-pkts"] = macsecTxStats.ScBypassPkts
    leafs["sc-eapol-pkts"] = macsecTxStats.ScEapolPkts
    leafs["sc-dropped-pkts"] = macsecTxStats.ScDroppedPkts
    leafs["current-an"] = macsecTxStats.CurrentAn
    leafs["sa-encrypted-pkts"] = macsecTxStats.SaEncryptedPkts
    return leafs
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetBundleName() string { return "cisco_ios_xr" }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetYangName() string { return "macsec-tx-stats" }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) SetParent(parent types.Entity) { macsecTxStats.parent = parent }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetParent() types.Entity { return macsecTxStats.parent }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetParentYangName() string { return "xlfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats
// Rx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDecryptedOctets interface{}

    // Rx No Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoTagPkts interface{}

    // Rx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Rx Bad Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBadTagPkts interface{}

    // Rx Late Pkts. The type is interface{} with range: 0..18446744073709551615.
    ScLatePkts interface{}

    // Rx Delayed Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDelayedPkts interface{}

    // Rx Unchecked Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUncheckedPkts interface{}

    // Rx No SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoSciPkts interface{}

    // Rx Unknown SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnknownSciPkts interface{}

    // Rx Pkts Ok. The type is interface{} with range: 0..18446744073709551615.
    ScOkPkts interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotUsingPkts interface{}

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnusedPkts interface{}

    // Rx Not Valid Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotValidPkts interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    ScInvalidPkts interface{}

    // Rx Overrun Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Rx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Rx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Rx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Rx SA Level Stats. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat.
    RxSaStat []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetFilter() yfilter.YFilter { return macsecRxStats.YFilter }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) SetFilter(yf yfilter.YFilter) { macsecRxStats.YFilter = yf }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetGoName(yname string) string {
    if yname == "sc-decrypted-octets" { return "ScDecryptedOctets" }
    if yname == "sc-no-tag-pkts" { return "ScNoTagPkts" }
    if yname == "sc-untagged-pkts" { return "ScUntaggedPkts" }
    if yname == "sc-bad-tag-pkts" { return "ScBadTagPkts" }
    if yname == "sc-late-pkts" { return "ScLatePkts" }
    if yname == "sc-delayed-pkts" { return "ScDelayedPkts" }
    if yname == "sc-unchecked-pkts" { return "ScUncheckedPkts" }
    if yname == "sc-no-sci-pkts" { return "ScNoSciPkts" }
    if yname == "sc-unknown-sci-pkts" { return "ScUnknownSciPkts" }
    if yname == "sc-ok-pkts" { return "ScOkPkts" }
    if yname == "sc-not-using-pkts" { return "ScNotUsingPkts" }
    if yname == "sc-unused-pkts" { return "ScUnusedPkts" }
    if yname == "sc-not-valid-pkts" { return "ScNotValidPkts" }
    if yname == "sc-invalid-pkts" { return "ScInvalidPkts" }
    if yname == "sc-overrun-pkts" { return "ScOverrunPkts" }
    if yname == "sc-bypass-pkts" { return "ScBypassPkts" }
    if yname == "sc-eapol-pkts" { return "ScEapolPkts" }
    if yname == "sc-dropped-pkts" { return "ScDroppedPkts" }
    if yname == "rx-sa-stat" { return "RxSaStat" }
    return ""
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetSegmentPath() string {
    return "macsec-rx-stats"
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-sa-stat" {
        for _, c := range macsecRxStats.RxSaStat {
            if macsecRxStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat{}
        macsecRxStats.RxSaStat = append(macsecRxStats.RxSaStat, child)
        return &macsecRxStats.RxSaStat[len(macsecRxStats.RxSaStat)-1]
    }
    return nil
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macsecRxStats.RxSaStat {
        children[macsecRxStats.RxSaStat[i].GetSegmentPath()] = &macsecRxStats.RxSaStat[i]
    }
    return children
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sc-decrypted-octets"] = macsecRxStats.ScDecryptedOctets
    leafs["sc-no-tag-pkts"] = macsecRxStats.ScNoTagPkts
    leafs["sc-untagged-pkts"] = macsecRxStats.ScUntaggedPkts
    leafs["sc-bad-tag-pkts"] = macsecRxStats.ScBadTagPkts
    leafs["sc-late-pkts"] = macsecRxStats.ScLatePkts
    leafs["sc-delayed-pkts"] = macsecRxStats.ScDelayedPkts
    leafs["sc-unchecked-pkts"] = macsecRxStats.ScUncheckedPkts
    leafs["sc-no-sci-pkts"] = macsecRxStats.ScNoSciPkts
    leafs["sc-unknown-sci-pkts"] = macsecRxStats.ScUnknownSciPkts
    leafs["sc-ok-pkts"] = macsecRxStats.ScOkPkts
    leafs["sc-not-using-pkts"] = macsecRxStats.ScNotUsingPkts
    leafs["sc-unused-pkts"] = macsecRxStats.ScUnusedPkts
    leafs["sc-not-valid-pkts"] = macsecRxStats.ScNotValidPkts
    leafs["sc-invalid-pkts"] = macsecRxStats.ScInvalidPkts
    leafs["sc-overrun-pkts"] = macsecRxStats.ScOverrunPkts
    leafs["sc-bypass-pkts"] = macsecRxStats.ScBypassPkts
    leafs["sc-eapol-pkts"] = macsecRxStats.ScEapolPkts
    leafs["sc-dropped-pkts"] = macsecRxStats.ScDroppedPkts
    return leafs
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetBundleName() string { return "cisco_ios_xr" }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetYangName() string { return "macsec-rx-stats" }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) SetParent(parent types.Entity) { macsecRxStats.parent = parent }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetParent() types.Entity { return macsecRxStats.parent }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetParentYangName() string { return "xlfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
// Rx SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current Rx AN. The type is interface{} with range: 0..18446744073709551615.
    An interface{}

    // Rx Ok Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaOkPkts interface{}

    // Rx Pkts not using SA for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotUsingPkts interface{}

    // Rx Pkts Unused Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaUnusedPkts interface{}

    // Rx Not Valid Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotValidPkts interface{}

    // Rx Invalid Pkts for current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaInvalidPkts interface{}
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetFilter() yfilter.YFilter { return rxSaStat.YFilter }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) SetFilter(yf yfilter.YFilter) { rxSaStat.YFilter = yf }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetGoName(yname string) string {
    if yname == "an" { return "An" }
    if yname == "sa-ok-pkts" { return "SaOkPkts" }
    if yname == "sa-not-using-pkts" { return "SaNotUsingPkts" }
    if yname == "sa-unused-pkts" { return "SaUnusedPkts" }
    if yname == "sa-not-valid-pkts" { return "SaNotValidPkts" }
    if yname == "sa-invalid-pkts" { return "SaInvalidPkts" }
    return ""
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetSegmentPath() string {
    return "rx-sa-stat"
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["an"] = rxSaStat.An
    leafs["sa-ok-pkts"] = rxSaStat.SaOkPkts
    leafs["sa-not-using-pkts"] = rxSaStat.SaNotUsingPkts
    leafs["sa-unused-pkts"] = rxSaStat.SaUnusedPkts
    leafs["sa-not-valid-pkts"] = rxSaStat.SaNotValidPkts
    leafs["sa-invalid-pkts"] = rxSaStat.SaInvalidPkts
    return leafs
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetBundleName() string { return "cisco_ios_xr" }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetYangName() string { return "rx-sa-stat" }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) SetParent(parent types.Entity) { rxSaStat.parent = parent }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetParent() types.Entity { return rxSaStat.parent }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetParentYangName() string { return "macsec-rx-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats
// ES200 Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats

    // Tx SC Macsec Stats.
    TxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats

    // Rx SC Macsec Stats.
    RxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats

    // Port level TX Stats.
    TxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats

    // Port level RX Stats.
    RxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetFilter() yfilter.YFilter { return es200Stats.YFilter }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) SetFilter(yf yfilter.YFilter) { es200Stats.YFilter = yf }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetGoName(yname string) string {
    if yname == "tx-sa-stats" { return "TxSaStats" }
    if yname == "rx-sa-stats" { return "RxSaStats" }
    if yname == "tx-sc-macsec-stats" { return "TxScMacsecStats" }
    if yname == "rx-sc-macsec-stats" { return "RxScMacsecStats" }
    if yname == "tx-interface-macsec-stats" { return "TxInterfaceMacsecStats" }
    if yname == "rx-interface-macsec-stats" { return "RxInterfaceMacsecStats" }
    if yname == "tx-port-stats" { return "TxPortStats" }
    if yname == "rx-port-stats" { return "RxPortStats" }
    return ""
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetSegmentPath() string {
    return "es200-stats"
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-sa-stats" {
        return &es200Stats.TxSaStats
    }
    if childYangName == "rx-sa-stats" {
        return &es200Stats.RxSaStats
    }
    if childYangName == "tx-sc-macsec-stats" {
        return &es200Stats.TxScMacsecStats
    }
    if childYangName == "rx-sc-macsec-stats" {
        return &es200Stats.RxScMacsecStats
    }
    if childYangName == "tx-interface-macsec-stats" {
        return &es200Stats.TxInterfaceMacsecStats
    }
    if childYangName == "rx-interface-macsec-stats" {
        return &es200Stats.RxInterfaceMacsecStats
    }
    if childYangName == "tx-port-stats" {
        return &es200Stats.TxPortStats
    }
    if childYangName == "rx-port-stats" {
        return &es200Stats.RxPortStats
    }
    return nil
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-sa-stats"] = &es200Stats.TxSaStats
    children["rx-sa-stats"] = &es200Stats.RxSaStats
    children["tx-sc-macsec-stats"] = &es200Stats.TxScMacsecStats
    children["rx-sc-macsec-stats"] = &es200Stats.RxScMacsecStats
    children["tx-interface-macsec-stats"] = &es200Stats.TxInterfaceMacsecStats
    children["rx-interface-macsec-stats"] = &es200Stats.RxInterfaceMacsecStats
    children["tx-port-stats"] = &es200Stats.TxPortStats
    children["rx-port-stats"] = &es200Stats.RxPortStats
    return children
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetBundleName() string { return "cisco_ios_xr" }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetYangName() string { return "es200-stats" }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) SetParent(parent types.Entity) { es200Stats.parent = parent }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetParent() types.Entity { return es200Stats.parent }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // packets exceeding egress MTU. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsTooLong interface{}

    // packets encrypted/protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncryptedProtected interface{}

    // octets1 encrypted/protected ?. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncryptedProtected1 interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetFilter() yfilter.YFilter { return txSaStats.YFilter }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) SetFilter(yf yfilter.YFilter) { txSaStats.YFilter = yf }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetGoName(yname string) string {
    if yname == "out-pkts-too-long" { return "OutPktsTooLong" }
    if yname == "out-pkts-encrypted-protected" { return "OutPktsEncryptedProtected" }
    if yname == "out-octets-encrypted-protected1" { return "OutOctetsEncryptedProtected1" }
    return ""
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetSegmentPath() string {
    return "tx-sa-stats"
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkts-too-long"] = txSaStats.OutPktsTooLong
    leafs["out-pkts-encrypted-protected"] = txSaStats.OutPktsEncryptedProtected
    leafs["out-octets-encrypted-protected1"] = txSaStats.OutOctetsEncryptedProtected1
    return leafs
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetYangName() string { return "tx-sa-stats" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) SetParent(parent types.Entity) { txSaStats.parent = parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetParent() types.Entity { return txSaStats.parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // frame not valid & validateFrames disabled. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsUnchecked interface{}

    // PN of packet outside replay window & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsDelayed interface{}

    // PN of packet outside replay window & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // packets with no error. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsOk interface{}

    // packet not valid & validateFrames !strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsInvalid interface{}

    // packet not valid & validateFrames strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsNotValid interface{}

    // packet assigned to SA not in use & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // packet assigned to SA not in use & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // octets1 decrypted/validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecryptedValidated1 interface{}

    // octets validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetFilter() yfilter.YFilter { return rxSaStats.YFilter }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) SetFilter(yf yfilter.YFilter) { rxSaStats.YFilter = yf }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetGoName(yname string) string {
    if yname == "in-pkts-unchecked" { return "InPktsUnchecked" }
    if yname == "in-pkts-delayed" { return "InPktsDelayed" }
    if yname == "in-pkts-late" { return "InPktsLate" }
    if yname == "in-pkts-ok" { return "InPktsOk" }
    if yname == "in-pkts-invalid" { return "InPktsInvalid" }
    if yname == "in-pkts-not-valid" { return "InPktsNotValid" }
    if yname == "in-pkts-not-using-sa" { return "InPktsNotUsingSa" }
    if yname == "in-pkts-unused-sa" { return "InPktsUnusedSa" }
    if yname == "in-octets-decrypted-validated1" { return "InOctetsDecryptedValidated1" }
    if yname == "in-octets-validated" { return "InOctetsValidated" }
    return ""
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetSegmentPath() string {
    return "rx-sa-stats"
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-unchecked"] = rxSaStats.InPktsUnchecked
    leafs["in-pkts-delayed"] = rxSaStats.InPktsDelayed
    leafs["in-pkts-late"] = rxSaStats.InPktsLate
    leafs["in-pkts-ok"] = rxSaStats.InPktsOk
    leafs["in-pkts-invalid"] = rxSaStats.InPktsInvalid
    leafs["in-pkts-not-valid"] = rxSaStats.InPktsNotValid
    leafs["in-pkts-not-using-sa"] = rxSaStats.InPktsNotUsingSa
    leafs["in-pkts-unused-sa"] = rxSaStats.InPktsUnusedSa
    leafs["in-octets-decrypted-validated1"] = rxSaStats.InOctetsDecryptedValidated1
    leafs["in-octets-validated"] = rxSaStats.InOctetsValidated
    return leafs
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetYangName() string { return "rx-sa-stats" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) SetParent(parent types.Entity) { rxSaStats.parent = parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetParent() types.Entity { return rxSaStats.parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats
// Tx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsSaNotInUse interface{}
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetFilter() yfilter.YFilter { return txScMacsecStats.YFilter }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) SetFilter(yf yfilter.YFilter) { txScMacsecStats.YFilter = yf }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetGoName(yname string) string {
    if yname == "out-pkts-sa-not-in-use" { return "OutPktsSaNotInUse" }
    return ""
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetSegmentPath() string {
    return "tx-sc-macsec-stats"
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkts-sa-not-in-use"] = txScMacsecStats.OutPktsSaNotInUse
    return leafs
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetYangName() string { return "tx-sc-macsec-stats" }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) SetParent(parent types.Entity) { txScMacsecStats.parent = parent }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetParent() types.Entity { return txScMacsecStats.parent }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats
// Rx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsSaNotInUse interface{}
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetFilter() yfilter.YFilter { return rxScMacsecStats.YFilter }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) SetFilter(yf yfilter.YFilter) { rxScMacsecStats.YFilter = yf }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetGoName(yname string) string {
    if yname == "in-pkts-sa-not-in-use" { return "InPktsSaNotInUse" }
    return ""
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetSegmentPath() string {
    return "rx-sc-macsec-stats"
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-sa-not-in-use"] = rxScMacsecStats.InPktsSaNotInUse
    return leafs
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetYangName() string { return "rx-sc-macsec-stats" }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) SetParent(parent types.Entity) { rxScMacsecStats.parent = parent }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetParent() types.Entity { return rxScMacsecStats.parent }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // egress packet that is classified as control packet. The type is interface{}
    // with range: 0..18446744073709551615.
    OutPktCtrl interface{}

    // egress packet to go out untagged when protectFrames not set. The type is
    // interface{} with range: 0..18446744073709551615.
    OutPktsUntagged interface{}

    // Octets tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsUnctrl interface{}

    // Octets tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCtrl interface{}

    // Octets tx on common port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCommon interface{}

    // Unicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsUnctrl interface{}

    // Unicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsCtrl interface{}

    // Multicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsUnctrl interface{}

    // Multicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsCtrl interface{}

    // Broadcast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsUnctrl interface{}

    // Broadcast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsClass interface{}

    // Packets dropped due to overflow in  processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsData interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return txInterfaceMacsecStats.YFilter }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { txInterfaceMacsecStats.YFilter = yf }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "transform-error-pkts" { return "TransformErrorPkts" }
    if yname == "out-pkt-ctrl" { return "OutPktCtrl" }
    if yname == "out-pkts-untagged" { return "OutPktsUntagged" }
    if yname == "out-octets-unctrl" { return "OutOctetsUnctrl" }
    if yname == "out-octets-ctrl" { return "OutOctetsCtrl" }
    if yname == "out-octets-common" { return "OutOctetsCommon" }
    if yname == "out-ucast-pkts-unctrl" { return "OutUcastPktsUnctrl" }
    if yname == "out-ucast-pkts-ctrl" { return "OutUcastPktsCtrl" }
    if yname == "out-mcast-pkts-unctrl" { return "OutMcastPktsUnctrl" }
    if yname == "out-mcast-pkts-ctrl" { return "OutMcastPktsCtrl" }
    if yname == "out-bcast-pkts-unctrl" { return "OutBcastPktsUnctrl" }
    if yname == "out-bcast-pkts-ctrl" { return "OutBcastPktsCtrl" }
    if yname == "out-rx-drop-pkts-unctrl" { return "OutRxDropPktsUnctrl" }
    if yname == "out-rx-drop-pkts-ctrl" { return "OutRxDropPktsCtrl" }
    if yname == "out-rx-err-pkts-unctrl" { return "OutRxErrPktsUnctrl" }
    if yname == "out-rx-err-pkts-ctrl" { return "OutRxErrPktsCtrl" }
    if yname == "out-drop-pkts-class" { return "OutDropPktsClass" }
    if yname == "out-drop-pkts-data" { return "OutDropPktsData" }
    return ""
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetSegmentPath() string {
    return "tx-interface-macsec-stats"
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transform-error-pkts"] = txInterfaceMacsecStats.TransformErrorPkts
    leafs["out-pkt-ctrl"] = txInterfaceMacsecStats.OutPktCtrl
    leafs["out-pkts-untagged"] = txInterfaceMacsecStats.OutPktsUntagged
    leafs["out-octets-unctrl"] = txInterfaceMacsecStats.OutOctetsUnctrl
    leafs["out-octets-ctrl"] = txInterfaceMacsecStats.OutOctetsCtrl
    leafs["out-octets-common"] = txInterfaceMacsecStats.OutOctetsCommon
    leafs["out-ucast-pkts-unctrl"] = txInterfaceMacsecStats.OutUcastPktsUnctrl
    leafs["out-ucast-pkts-ctrl"] = txInterfaceMacsecStats.OutUcastPktsCtrl
    leafs["out-mcast-pkts-unctrl"] = txInterfaceMacsecStats.OutMcastPktsUnctrl
    leafs["out-mcast-pkts-ctrl"] = txInterfaceMacsecStats.OutMcastPktsCtrl
    leafs["out-bcast-pkts-unctrl"] = txInterfaceMacsecStats.OutBcastPktsUnctrl
    leafs["out-bcast-pkts-ctrl"] = txInterfaceMacsecStats.OutBcastPktsCtrl
    leafs["out-rx-drop-pkts-unctrl"] = txInterfaceMacsecStats.OutRxDropPktsUnctrl
    leafs["out-rx-drop-pkts-ctrl"] = txInterfaceMacsecStats.OutRxDropPktsCtrl
    leafs["out-rx-err-pkts-unctrl"] = txInterfaceMacsecStats.OutRxErrPktsUnctrl
    leafs["out-rx-err-pkts-ctrl"] = txInterfaceMacsecStats.OutRxErrPktsCtrl
    leafs["out-drop-pkts-class"] = txInterfaceMacsecStats.OutDropPktsClass
    leafs["out-drop-pkts-data"] = txInterfaceMacsecStats.OutDropPktsData
    return leafs
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetYangName() string { return "tx-interface-macsec-stats" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) SetParent(parent types.Entity) { txInterfaceMacsecStats.parent = parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetParent() types.Entity { return txInterfaceMacsecStats.parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // ingress packet that is classified as control packet. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktCtrl interface{}

    // ingress packet untagged & validateFrames is strict. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktNoTag interface{}

    // ingress packet untagged & validateFrames is  !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUntagged interface{}

    // ingress frames received with an invalid MACSec tag or ICV                  
    // added with next one gives InPktsSCIMiss. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktBadTag interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is !strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnknownSci interface{}

    // ingress packets that are control or KaY packets. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktsTaggedCtrl interface{}

    // Octets rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsUnctrl interface{}

    // Octets rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsCtrl interface{}

    // Unicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsUnctrl interface{}

    // Unicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsCtrl interface{}

    // Multicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsUnctrl interface{}

    // Multicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsCtrl interface{}

    // Broadcast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsUnctrl interface{}

    // Broadcast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsClass interface{}

    // Packets dropped due to overflow in processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsData interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return rxInterfaceMacsecStats.YFilter }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { rxInterfaceMacsecStats.YFilter = yf }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "transform-error-pkts" { return "TransformErrorPkts" }
    if yname == "in-pkt-ctrl" { return "InPktCtrl" }
    if yname == "in-pkt-no-tag" { return "InPktNoTag" }
    if yname == "in-pkts-untagged" { return "InPktsUntagged" }
    if yname == "in-pkt-bad-tag" { return "InPktBadTag" }
    if yname == "in-pkt-no-sci" { return "InPktNoSci" }
    if yname == "in-pkts-unknown-sci" { return "InPktsUnknownSci" }
    if yname == "in-pkts-tagged-ctrl" { return "InPktsTaggedCtrl" }
    if yname == "in-octets-unctrl" { return "InOctetsUnctrl" }
    if yname == "in-octets-ctrl" { return "InOctetsCtrl" }
    if yname == "in-ucast-pkts-unctrl" { return "InUcastPktsUnctrl" }
    if yname == "in-ucast-pkts-ctrl" { return "InUcastPktsCtrl" }
    if yname == "in-mcast-pkts-unctrl" { return "InMcastPktsUnctrl" }
    if yname == "in-mcast-pkts-ctrl" { return "InMcastPktsCtrl" }
    if yname == "in-bcast-pkts-unctrl" { return "InBcastPktsUnctrl" }
    if yname == "in-bcast-pkts-ctrl" { return "InBcastPktsCtrl" }
    if yname == "in-rx-drop-pkts-unctrl" { return "InRxDropPktsUnctrl" }
    if yname == "in-rx-drop-pkts-ctrl" { return "InRxDropPktsCtrl" }
    if yname == "in-rx-error-pkts-unctrl" { return "InRxErrorPktsUnctrl" }
    if yname == "in-rx-error-pkts-ctrl" { return "InRxErrorPktsCtrl" }
    if yname == "in-drop-pkts-class" { return "InDropPktsClass" }
    if yname == "in-drop-pkts-data" { return "InDropPktsData" }
    return ""
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetSegmentPath() string {
    return "rx-interface-macsec-stats"
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transform-error-pkts"] = rxInterfaceMacsecStats.TransformErrorPkts
    leafs["in-pkt-ctrl"] = rxInterfaceMacsecStats.InPktCtrl
    leafs["in-pkt-no-tag"] = rxInterfaceMacsecStats.InPktNoTag
    leafs["in-pkts-untagged"] = rxInterfaceMacsecStats.InPktsUntagged
    leafs["in-pkt-bad-tag"] = rxInterfaceMacsecStats.InPktBadTag
    leafs["in-pkt-no-sci"] = rxInterfaceMacsecStats.InPktNoSci
    leafs["in-pkts-unknown-sci"] = rxInterfaceMacsecStats.InPktsUnknownSci
    leafs["in-pkts-tagged-ctrl"] = rxInterfaceMacsecStats.InPktsTaggedCtrl
    leafs["in-octets-unctrl"] = rxInterfaceMacsecStats.InOctetsUnctrl
    leafs["in-octets-ctrl"] = rxInterfaceMacsecStats.InOctetsCtrl
    leafs["in-ucast-pkts-unctrl"] = rxInterfaceMacsecStats.InUcastPktsUnctrl
    leafs["in-ucast-pkts-ctrl"] = rxInterfaceMacsecStats.InUcastPktsCtrl
    leafs["in-mcast-pkts-unctrl"] = rxInterfaceMacsecStats.InMcastPktsUnctrl
    leafs["in-mcast-pkts-ctrl"] = rxInterfaceMacsecStats.InMcastPktsCtrl
    leafs["in-bcast-pkts-unctrl"] = rxInterfaceMacsecStats.InBcastPktsUnctrl
    leafs["in-bcast-pkts-ctrl"] = rxInterfaceMacsecStats.InBcastPktsCtrl
    leafs["in-rx-drop-pkts-unctrl"] = rxInterfaceMacsecStats.InRxDropPktsUnctrl
    leafs["in-rx-drop-pkts-ctrl"] = rxInterfaceMacsecStats.InRxDropPktsCtrl
    leafs["in-rx-error-pkts-unctrl"] = rxInterfaceMacsecStats.InRxErrorPktsUnctrl
    leafs["in-rx-error-pkts-ctrl"] = rxInterfaceMacsecStats.InRxErrorPktsCtrl
    leafs["in-drop-pkts-class"] = rxInterfaceMacsecStats.InDropPktsClass
    leafs["in-drop-pkts-data"] = rxInterfaceMacsecStats.InDropPktsData
    return leafs
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetYangName() string { return "rx-interface-macsec-stats" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) SetParent(parent types.Entity) { rxInterfaceMacsecStats.parent = parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetParent() types.Entity { return rxInterfaceMacsecStats.parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats
// Port level TX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetFilter() yfilter.YFilter { return txPortStats.YFilter }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) SetFilter(yf yfilter.YFilter) { txPortStats.YFilter = yf }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetGoName(yname string) string {
    if yname == "multi-flow-match" { return "MultiFlowMatch" }
    if yname == "parser-dropped" { return "ParserDropped" }
    if yname == "flow-miss" { return "FlowMiss" }
    if yname == "pkts-ctrl" { return "PktsCtrl" }
    if yname == "pkts-data" { return "PktsData" }
    if yname == "pkts-dropped" { return "PktsDropped" }
    if yname == "pkts-err-in" { return "PktsErrIn" }
    return ""
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetSegmentPath() string {
    return "tx-port-stats"
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multi-flow-match"] = txPortStats.MultiFlowMatch
    leafs["parser-dropped"] = txPortStats.ParserDropped
    leafs["flow-miss"] = txPortStats.FlowMiss
    leafs["pkts-ctrl"] = txPortStats.PktsCtrl
    leafs["pkts-data"] = txPortStats.PktsData
    leafs["pkts-dropped"] = txPortStats.PktsDropped
    leafs["pkts-err-in"] = txPortStats.PktsErrIn
    return leafs
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetBundleName() string { return "cisco_ios_xr" }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetYangName() string { return "tx-port-stats" }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) SetParent(parent types.Entity) { txPortStats.parent = parent }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetParent() types.Entity { return txPortStats.parent }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats
// Port level RX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetFilter() yfilter.YFilter { return rxPortStats.YFilter }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) SetFilter(yf yfilter.YFilter) { rxPortStats.YFilter = yf }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetGoName(yname string) string {
    if yname == "multi-flow-match" { return "MultiFlowMatch" }
    if yname == "parser-dropped" { return "ParserDropped" }
    if yname == "flow-miss" { return "FlowMiss" }
    if yname == "pkts-ctrl" { return "PktsCtrl" }
    if yname == "pkts-data" { return "PktsData" }
    if yname == "pkts-dropped" { return "PktsDropped" }
    if yname == "pkts-err-in" { return "PktsErrIn" }
    return ""
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetSegmentPath() string {
    return "rx-port-stats"
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multi-flow-match"] = rxPortStats.MultiFlowMatch
    leafs["parser-dropped"] = rxPortStats.ParserDropped
    leafs["flow-miss"] = rxPortStats.FlowMiss
    leafs["pkts-ctrl"] = rxPortStats.PktsCtrl
    leafs["pkts-data"] = rxPortStats.PktsData
    leafs["pkts-dropped"] = rxPortStats.PktsDropped
    leafs["pkts-err-in"] = rxPortStats.PktsErrIn
    return leafs
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetYangName() string { return "rx-port-stats" }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) SetParent(parent types.Entity) { rxPortStats.parent = parent }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetParent() types.Entity { return rxPortStats.parent }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas
// Table of Hardware SAs
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hardware Security Association. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa.
    HwSa []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa
}

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetFilter() yfilter.YFilter { return hwSas.YFilter }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) SetFilter(yf yfilter.YFilter) { hwSas.YFilter = yf }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetGoName(yname string) string {
    if yname == "hw-sa" { return "HwSa" }
    return ""
}

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetSegmentPath() string {
    return "hw-sas"
}

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-sa" {
        for _, c := range hwSas.HwSa {
            if hwSas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa{}
        hwSas.HwSa = append(hwSas.HwSa, child)
        return &hwSas.HwSa[len(hwSas.HwSa)-1]
    }
    return nil
}

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwSas.HwSa {
        children[hwSas.HwSa[i].GetSegmentPath()] = &hwSas.HwSa[i]
    }
    return children
}

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetBundleName() string { return "cisco_ios_xr" }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetYangName() string { return "hw-sas" }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) SetParent(parent types.Entity) { hwSas.parent = parent }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetParent() types.Entity { return hwSas.parent }

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetParentYangName() string { return "interface" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa
// Hardware Security Association
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SA ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SaId interface{}

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext
}

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetFilter() yfilter.YFilter { return hwSa.YFilter }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) SetFilter(yf yfilter.YFilter) { hwSa.YFilter = yf }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetGoName(yname string) string {
    if yname == "sa-id" { return "SaId" }
    if yname == "ext" { return "Ext" }
    return ""
}

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetSegmentPath() string {
    return "hw-sa" + "[sa-id='" + fmt.Sprintf("%v", hwSa.SaId) + "']"
}

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext" {
        return &hwSa.Ext
    }
    return nil
}

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ext"] = &hwSa.Ext
    return children
}

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa-id"] = hwSa.SaId
    return leafs
}

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetBundleName() string { return "cisco_ios_xr" }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetYangName() string { return "hw-sa" }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) SetParent(parent types.Entity) { hwSa.parent = parent }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetParent() types.Entity { return hwSa.parent }

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetParentYangName() string { return "hw-sas" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type interface{}

    // MSFPGA SA Information.
    MsfpgaSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa

    // XLFPGA SA Information.
    XlfpgaSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa

    // ES200 SA Information.
    Es200Sa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetFilter() yfilter.YFilter { return ext.YFilter }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) SetFilter(yf yfilter.YFilter) { ext.YFilter = yf }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "msfpga-sa" { return "MsfpgaSa" }
    if yname == "xlfpga-sa" { return "XlfpgaSa" }
    if yname == "es200-sa" { return "Es200Sa" }
    return ""
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetSegmentPath() string {
    return "ext"
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msfpga-sa" {
        return &ext.MsfpgaSa
    }
    if childYangName == "xlfpga-sa" {
        return &ext.XlfpgaSa
    }
    if childYangName == "es200-sa" {
        return &ext.Es200Sa
    }
    return nil
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["msfpga-sa"] = &ext.MsfpgaSa
    children["xlfpga-sa"] = &ext.XlfpgaSa
    children["es200-sa"] = &ext.Es200Sa
    return children
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = ext.Type
    return leafs
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetBundleName() string { return "cisco_ios_xr" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetYangName() string { return "ext" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) SetParent(parent types.Entity) { ext.parent = parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetParent() types.Entity { return ext.parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetParentYangName() string { return "hw-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa
// MSFPGA SA Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SA Details.
    TxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa

    // Rx SA Details.
    RxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa
}

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetFilter() yfilter.YFilter { return msfpgaSa.YFilter }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) SetFilter(yf yfilter.YFilter) { msfpgaSa.YFilter = yf }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetGoName(yname string) string {
    if yname == "tx-sa" { return "TxSa" }
    if yname == "rx-sa" { return "RxSa" }
    return ""
}

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetSegmentPath() string {
    return "msfpga-sa"
}

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-sa" {
        return &msfpgaSa.TxSa
    }
    if childYangName == "rx-sa" {
        return &msfpgaSa.RxSa
    }
    return nil
}

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-sa"] = &msfpgaSa.TxSa
    children["rx-sa"] = &msfpgaSa.RxSa
    return children
}

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetBundleName() string { return "cisco_ios_xr" }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetYangName() string { return "msfpga-sa" }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) SetParent(parent types.Entity) { msfpgaSa.parent = parent }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetParent() types.Entity { return msfpgaSa.parent }

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa
// Tx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SA Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // Crypto Algorithm. The type is interface{} with range: 0..255.
    CryptoAlgo interface{}

    // Key Length. The type is interface{} with range: 0..255.
    KeyLen interface{}

    // Association Number. The type is interface{} with range: 0..255.
    An interface{}

    // XPN EN. The type is interface{} with range: 0..255.
    Xpn interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Next Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    NextPn interface{}

    // Conf offset. The type is interface{} with range: 0..255.
    COffset interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // Q bit. The type is bool.
    QBit interface{}

    // QQ bit. The type is bool.
    QqBit interface{}
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetFilter() yfilter.YFilter { return txSa.YFilter }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) SetFilter(yf yfilter.YFilter) { txSa.YFilter = yf }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetGoName(yname string) string {
    if yname == "sa-id" { return "SaId" }
    if yname == "valid" { return "Valid" }
    if yname == "is-egress" { return "IsEgress" }
    if yname == "crypto-algo" { return "CryptoAlgo" }
    if yname == "key-len" { return "KeyLen" }
    if yname == "an" { return "An" }
    if yname == "xpn" { return "Xpn" }
    if yname == "sci" { return "Sci" }
    if yname == "in-use" { return "InUse" }
    if yname == "next-pn" { return "NextPn" }
    if yname == "c-offset" { return "COffset" }
    if yname == "action" { return "Action" }
    if yname == "q-bit" { return "QBit" }
    if yname == "qq-bit" { return "QqBit" }
    return ""
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetSegmentPath() string {
    return "tx-sa"
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa-id"] = txSa.SaId
    leafs["valid"] = txSa.Valid
    leafs["is-egress"] = txSa.IsEgress
    leafs["crypto-algo"] = txSa.CryptoAlgo
    leafs["key-len"] = txSa.KeyLen
    leafs["an"] = txSa.An
    leafs["xpn"] = txSa.Xpn
    leafs["sci"] = txSa.Sci
    leafs["in-use"] = txSa.InUse
    leafs["next-pn"] = txSa.NextPn
    leafs["c-offset"] = txSa.COffset
    leafs["action"] = txSa.Action
    leafs["q-bit"] = txSa.QBit
    leafs["qq-bit"] = txSa.QqBit
    return leafs
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetBundleName() string { return "cisco_ios_xr" }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetYangName() string { return "tx-sa" }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) SetParent(parent types.Entity) { txSa.parent = parent }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetParent() types.Entity { return txSa.parent }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetParentYangName() string { return "msfpga-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa
// Rx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SA Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // Crypto Algorithm. The type is interface{} with range: 0..255.
    CryptoAlgo interface{}

    // Key Length. The type is interface{} with range: 0..255.
    KeyLen interface{}

    // Association Number. The type is interface{} with range: 0..255.
    An interface{}

    // XPN EN. The type is interface{} with range: 0..255.
    Xpn interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Next Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    NextPn interface{}

    // Conf offset. The type is interface{} with range: 0..255.
    COffset interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // Q bit. The type is bool.
    QBit interface{}

    // QQ bit. The type is bool.
    QqBit interface{}
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetFilter() yfilter.YFilter { return rxSa.YFilter }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) SetFilter(yf yfilter.YFilter) { rxSa.YFilter = yf }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetGoName(yname string) string {
    if yname == "sa-id" { return "SaId" }
    if yname == "valid" { return "Valid" }
    if yname == "is-egress" { return "IsEgress" }
    if yname == "crypto-algo" { return "CryptoAlgo" }
    if yname == "key-len" { return "KeyLen" }
    if yname == "an" { return "An" }
    if yname == "xpn" { return "Xpn" }
    if yname == "sci" { return "Sci" }
    if yname == "in-use" { return "InUse" }
    if yname == "next-pn" { return "NextPn" }
    if yname == "c-offset" { return "COffset" }
    if yname == "action" { return "Action" }
    if yname == "q-bit" { return "QBit" }
    if yname == "qq-bit" { return "QqBit" }
    return ""
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetSegmentPath() string {
    return "rx-sa"
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sa-id"] = rxSa.SaId
    leafs["valid"] = rxSa.Valid
    leafs["is-egress"] = rxSa.IsEgress
    leafs["crypto-algo"] = rxSa.CryptoAlgo
    leafs["key-len"] = rxSa.KeyLen
    leafs["an"] = rxSa.An
    leafs["xpn"] = rxSa.Xpn
    leafs["sci"] = rxSa.Sci
    leafs["in-use"] = rxSa.InUse
    leafs["next-pn"] = rxSa.NextPn
    leafs["c-offset"] = rxSa.COffset
    leafs["action"] = rxSa.Action
    leafs["q-bit"] = rxSa.QBit
    leafs["qq-bit"] = rxSa.QqBit
    return leafs
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetBundleName() string { return "cisco_ios_xr" }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetYangName() string { return "rx-sa" }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) SetParent(parent types.Entity) { rxSa.parent = parent }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetParent() types.Entity { return rxSa.parent }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetParentYangName() string { return "msfpga-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa
// XLFPGA SA Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SA Details.
    TxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa

    // Rx SA Details.
    RxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa
}

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetFilter() yfilter.YFilter { return xlfpgaSa.YFilter }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) SetFilter(yf yfilter.YFilter) { xlfpgaSa.YFilter = yf }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetGoName(yname string) string {
    if yname == "tx-sa" { return "TxSa" }
    if yname == "rx-sa" { return "RxSa" }
    return ""
}

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetSegmentPath() string {
    return "xlfpga-sa"
}

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-sa" {
        return &xlfpgaSa.TxSa
    }
    if childYangName == "rx-sa" {
        return &xlfpgaSa.RxSa
    }
    return nil
}

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-sa"] = &xlfpgaSa.TxSa
    children["rx-sa"] = &xlfpgaSa.RxSa
    return children
}

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetBundleName() string { return "cisco_ios_xr" }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetYangName() string { return "xlfpga-sa" }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) SetParent(parent types.Entity) { xlfpgaSa.parent = parent }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetParent() types.Entity { return xlfpgaSa.parent }

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa
// Tx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnable interface{}

    // Secure Mode - Must/Should. The type is interface{} with range: 0..255.
    SecureMode interface{}

    // Secure Channel ID. The type is interface{} with range:
    // 0..18446744073709551615.
    SecureChannelId interface{}

    // Sec Tag Length(bytes) . The type is interface{} with range: 0..4294967295.
    // Units are byte.
    SectagLength interface{}

    // Cipher Suite Used. The type is interface{} with range: 0..4294967295.
    CipherSuite interface{}

    // Confidentiality Offset. The type is interface{} with range: 0..255.
    ConfidentialityOffset interface{}

    // FCS Error Config. The type is interface{} with range: 0..255.
    FcsErrCfg interface{}

    // Max Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNum interface{}

    // Association Number. The type is interface{} with range: 0..255.
    An interface{}

    // Initial Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    InitialPacketNumber interface{}

    // Short Secure Channel ID. The type is interface{} with range: 0..4294967295.
    Ssci interface{}

    // Current Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    CurrentPacketNum interface{}

    // CRC Value. The type is interface{} with range: 0..4294967295.
    CrcValue interface{}
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetFilter() yfilter.YFilter { return txSa.YFilter }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) SetFilter(yf yfilter.YFilter) { txSa.YFilter = yf }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetGoName(yname string) string {
    if yname == "protection-enable" { return "ProtectionEnable" }
    if yname == "secure-mode" { return "SecureMode" }
    if yname == "secure-channel-id" { return "SecureChannelId" }
    if yname == "sectag-length" { return "SectagLength" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "confidentiality-offset" { return "ConfidentialityOffset" }
    if yname == "fcs-err-cfg" { return "FcsErrCfg" }
    if yname == "max-packet-num" { return "MaxPacketNum" }
    if yname == "an" { return "An" }
    if yname == "initial-packet-number" { return "InitialPacketNumber" }
    if yname == "ssci" { return "Ssci" }
    if yname == "current-packet-num" { return "CurrentPacketNum" }
    if yname == "crc-value" { return "CrcValue" }
    return ""
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetSegmentPath() string {
    return "tx-sa"
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protection-enable"] = txSa.ProtectionEnable
    leafs["secure-mode"] = txSa.SecureMode
    leafs["secure-channel-id"] = txSa.SecureChannelId
    leafs["sectag-length"] = txSa.SectagLength
    leafs["cipher-suite"] = txSa.CipherSuite
    leafs["confidentiality-offset"] = txSa.ConfidentialityOffset
    leafs["fcs-err-cfg"] = txSa.FcsErrCfg
    leafs["max-packet-num"] = txSa.MaxPacketNum
    leafs["an"] = txSa.An
    leafs["initial-packet-number"] = txSa.InitialPacketNumber
    leafs["ssci"] = txSa.Ssci
    leafs["current-packet-num"] = txSa.CurrentPacketNum
    leafs["crc-value"] = txSa.CrcValue
    return leafs
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetBundleName() string { return "cisco_ios_xr" }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetYangName() string { return "tx-sa" }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) SetParent(parent types.Entity) { txSa.parent = parent }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetParent() types.Entity { return txSa.parent }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetParentYangName() string { return "xlfpga-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa
// Rx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnable interface{}

    // Secure Mode - Must/Should. The type is interface{} with range:
    // 0..4294967295.
    SecureMode interface{}

    // Replay Protect Mode. The type is bool.
    ReplayProtectMode interface{}

    // Validation Mode. The type is interface{} with range: 0..4294967295.
    ValidationMode interface{}

    // Replay Window . The type is interface{} with range: 0..4294967295.
    ReplayWindow interface{}

    // Secure Channel ID. The type is interface{} with range:
    // 0..18446744073709551615.
    SecureChannelId interface{}

    // Cipher Suite Used. The type is interface{} with range: 0..4294967295.
    CipherSuite interface{}

    // Confidentiality Offset. The type is interface{} with range: 0..255.
    ConfidentialityOffset interface{}

    // FCS Error Config. The type is interface{} with range: 0..4294967295.
    FcsErrCfg interface{}

    // Auth  Error Config. The type is interface{} with range: 0..4294967295.
    AuthErrCfg interface{}

    // Max Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNum interface{}

    // Num of AN's in Use. The type is interface{} with range: 0..4294967295.
    NumAnInUse interface{}

    // Association Number. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    An interface{}

    // Recent Association Num. The type is interface{} with range: 0..255.
    RecentAn interface{}

    // Untagged Pkts Detected. The type is bool.
    PktUntaggedDetected interface{}

    // Tagged Pkts Detected. The type is bool.
    PktTaggedDetected interface{}

    // Tagged Pkts Validated. The type is bool.
    PktTaggedValidated interface{}

    // Current Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    CurrentPacketNum interface{}

    // Short Secure Channel ID. The type is slice of interface{} with range:
    // 0..4294967295.
    Ssci []interface{}

    // Lowest Acceptable Packet Number. The type is slice of interface{} with
    // range: 0..18446744073709551615.
    LowestAcceptablePacketNum []interface{}

    // Next expected Packet Number. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    NextExpectedPacketNum []interface{}

    // CRC Value. The type is slice of interface{} with range: 0..4294967295.
    CrcValue []interface{}
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetFilter() yfilter.YFilter { return rxSa.YFilter }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) SetFilter(yf yfilter.YFilter) { rxSa.YFilter = yf }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetGoName(yname string) string {
    if yname == "protection-enable" { return "ProtectionEnable" }
    if yname == "secure-mode" { return "SecureMode" }
    if yname == "replay-protect-mode" { return "ReplayProtectMode" }
    if yname == "validation-mode" { return "ValidationMode" }
    if yname == "replay-window" { return "ReplayWindow" }
    if yname == "secure-channel-id" { return "SecureChannelId" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "confidentiality-offset" { return "ConfidentialityOffset" }
    if yname == "fcs-err-cfg" { return "FcsErrCfg" }
    if yname == "auth-err-cfg" { return "AuthErrCfg" }
    if yname == "max-packet-num" { return "MaxPacketNum" }
    if yname == "num-an-in-use" { return "NumAnInUse" }
    if yname == "an" { return "An" }
    if yname == "recent-an" { return "RecentAn" }
    if yname == "pkt-untagged-detected" { return "PktUntaggedDetected" }
    if yname == "pkt-tagged-detected" { return "PktTaggedDetected" }
    if yname == "pkt-tagged-validated" { return "PktTaggedValidated" }
    if yname == "current-packet-num" { return "CurrentPacketNum" }
    if yname == "ssci" { return "Ssci" }
    if yname == "lowest-acceptable-packet-num" { return "LowestAcceptablePacketNum" }
    if yname == "next-expected-packet-num" { return "NextExpectedPacketNum" }
    if yname == "crc-value" { return "CrcValue" }
    return ""
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetSegmentPath() string {
    return "rx-sa"
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protection-enable"] = rxSa.ProtectionEnable
    leafs["secure-mode"] = rxSa.SecureMode
    leafs["replay-protect-mode"] = rxSa.ReplayProtectMode
    leafs["validation-mode"] = rxSa.ValidationMode
    leafs["replay-window"] = rxSa.ReplayWindow
    leafs["secure-channel-id"] = rxSa.SecureChannelId
    leafs["cipher-suite"] = rxSa.CipherSuite
    leafs["confidentiality-offset"] = rxSa.ConfidentialityOffset
    leafs["fcs-err-cfg"] = rxSa.FcsErrCfg
    leafs["auth-err-cfg"] = rxSa.AuthErrCfg
    leafs["max-packet-num"] = rxSa.MaxPacketNum
    leafs["num-an-in-use"] = rxSa.NumAnInUse
    leafs["an"] = rxSa.An
    leafs["recent-an"] = rxSa.RecentAn
    leafs["pkt-untagged-detected"] = rxSa.PktUntaggedDetected
    leafs["pkt-tagged-detected"] = rxSa.PktTaggedDetected
    leafs["pkt-tagged-validated"] = rxSa.PktTaggedValidated
    leafs["current-packet-num"] = rxSa.CurrentPacketNum
    leafs["ssci"] = rxSa.Ssci
    leafs["lowest-acceptable-packet-num"] = rxSa.LowestAcceptablePacketNum
    leafs["next-expected-packet-num"] = rxSa.NextExpectedPacketNum
    leafs["crc-value"] = rxSa.CrcValue
    return leafs
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetBundleName() string { return "cisco_ios_xr" }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetYangName() string { return "rx-sa" }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) SetParent(parent types.Entity) { rxSa.parent = parent }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetParent() types.Entity { return rxSa.parent }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetParentYangName() string { return "xlfpga-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa
// ES200 SA Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SA Details.
    TxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa

    // Rx SA Details. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa.
    RxSa []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa
}

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetFilter() yfilter.YFilter { return es200Sa.YFilter }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) SetFilter(yf yfilter.YFilter) { es200Sa.YFilter = yf }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetGoName(yname string) string {
    if yname == "tx-sa" { return "TxSa" }
    if yname == "rx-sa" { return "RxSa" }
    return ""
}

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetSegmentPath() string {
    return "es200-sa"
}

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-sa" {
        return &es200Sa.TxSa
    }
    if childYangName == "rx-sa" {
        for _, c := range es200Sa.RxSa {
            if es200Sa.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa{}
        es200Sa.RxSa = append(es200Sa.RxSa, child)
        return &es200Sa.RxSa[len(es200Sa.RxSa)-1]
    }
    return nil
}

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-sa"] = &es200Sa.TxSa
    for i := range es200Sa.RxSa {
        children[es200Sa.RxSa[i].GetSegmentPath()] = &es200Sa.RxSa[i]
    }
    return children
}

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetBundleName() string { return "cisco_ios_xr" }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetYangName() string { return "es200-sa" }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) SetParent(parent types.Entity) { es200Sa.parent = parent }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetParent() types.Entity { return es200Sa.parent }

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa
// Tx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is structure valid. The type is bool.
    IsValid interface{}

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SC Number. The type is interface{} with range: 0..4294967295.
    ScNo interface{}

    // packets exceeding egress MTU. The type is interface{} with range: 0..255.
    OutPktsTooLong interface{}

    // packets encrypted/protected. The type is interface{} with range: 0..255.
    OutPktsEncryptedProtected interface{}

    // octets1 encrypted/protected. The type is interface{} with range: 0..255.
    OutOctetsEncryptedProtected1 interface{}

    // Initial Packet Number. The type is interface{} with range: 0..255.
    InitialPktNumber interface{}

    // Current packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    CurrentPktNumber interface{}

    // Maximum packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPktNumber interface{}

    // Xform Params.
    XformParams MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetFilter() yfilter.YFilter { return txSa.YFilter }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) SetFilter(yf yfilter.YFilter) { txSa.YFilter = yf }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetGoName(yname string) string {
    if yname == "is-valid" { return "IsValid" }
    if yname == "sa-id" { return "SaId" }
    if yname == "sc-no" { return "ScNo" }
    if yname == "out-pkts-too-long" { return "OutPktsTooLong" }
    if yname == "out-pkts-encrypted-protected" { return "OutPktsEncryptedProtected" }
    if yname == "out-octets-encrypted-protected1" { return "OutOctetsEncryptedProtected1" }
    if yname == "initial-pkt-number" { return "InitialPktNumber" }
    if yname == "current-pkt-number" { return "CurrentPktNumber" }
    if yname == "max-pkt-number" { return "MaxPktNumber" }
    if yname == "xform-params" { return "XformParams" }
    return ""
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetSegmentPath() string {
    return "tx-sa"
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "xform-params" {
        return &txSa.XformParams
    }
    return nil
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["xform-params"] = &txSa.XformParams
    return children
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-valid"] = txSa.IsValid
    leafs["sa-id"] = txSa.SaId
    leafs["sc-no"] = txSa.ScNo
    leafs["out-pkts-too-long"] = txSa.OutPktsTooLong
    leafs["out-pkts-encrypted-protected"] = txSa.OutPktsEncryptedProtected
    leafs["out-octets-encrypted-protected1"] = txSa.OutOctetsEncryptedProtected1
    leafs["initial-pkt-number"] = txSa.InitialPktNumber
    leafs["current-pkt-number"] = txSa.CurrentPktNumber
    leafs["max-pkt-number"] = txSa.MaxPktNumber
    return leafs
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetBundleName() string { return "cisco_ios_xr" }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetYangName() string { return "tx-sa" }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) SetParent(parent types.Entity) { txSa.parent = parent }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetParent() types.Entity { return txSa.parent }

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetParentYangName() string { return "es200-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams
//  Xform Params
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // range of pkt nos considered valid. The type is interface{} with range:
    // 0..4294967295.
    ReplayWinSize interface{}

    // Cryptographic algo used. The type is string.
    CryptAlgo interface{}

    // APM_TRUE if this is Egress Transform record, APM_FALSE otherwise. The type
    // is bool.
    IsEgressTr interface{}

    // AES Key length. The type is string.
    AesKeyLen interface{}

    // Association Number for egress. The type is interface{} with range: 0..255.
    AssocNum interface{}

    // TRUE if Seq Num is 64-bit, FALSE if it is 32-bit. The type is bool.
    IsSeqNum64Bit interface{}

    // TRUE to generate the authKey, so authKey in this struct not used           
    // APM_FALSE to use provided authKey. The type is bool.
    BgenAuthKey interface{}
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetFilter() yfilter.YFilter { return xformParams.YFilter }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) SetFilter(yf yfilter.YFilter) { xformParams.YFilter = yf }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetGoName(yname string) string {
    if yname == "replay-win-size" { return "ReplayWinSize" }
    if yname == "crypt-algo" { return "CryptAlgo" }
    if yname == "is-egress-tr" { return "IsEgressTr" }
    if yname == "aes-key-len" { return "AesKeyLen" }
    if yname == "assoc-num" { return "AssocNum" }
    if yname == "is-seq-num64-bit" { return "IsSeqNum64Bit" }
    if yname == "bgen-auth-key" { return "BgenAuthKey" }
    return ""
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetSegmentPath() string {
    return "xform-params"
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["replay-win-size"] = xformParams.ReplayWinSize
    leafs["crypt-algo"] = xformParams.CryptAlgo
    leafs["is-egress-tr"] = xformParams.IsEgressTr
    leafs["aes-key-len"] = xformParams.AesKeyLen
    leafs["assoc-num"] = xformParams.AssocNum
    leafs["is-seq-num64-bit"] = xformParams.IsSeqNum64Bit
    leafs["bgen-auth-key"] = xformParams.BgenAuthKey
    return leafs
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetBundleName() string { return "cisco_ios_xr" }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetYangName() string { return "xform-params" }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) SetParent(parent types.Entity) { xformParams.parent = parent }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetParent() types.Entity { return xformParams.parent }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetParentYangName() string { return "tx-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa
// Rx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is structure valid. The type is bool.
    IsValid interface{}

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SC Number. The type is interface{} with range: 0..4294967295.
    ScNo interface{}

    // frame not valid & validateFrames disabled. The type is interface{} with
    // range: 0..255.
    InPktsUnchecked interface{}

    // PN of packet outside replay window & validateFrames !strict. The type is
    // interface{} with range: 0..255.
    InPktsDelayed interface{}

    // PN of packet outside replay window & validateFrames strict. The type is
    // interface{} with range: 0..255.
    InPktsLate interface{}

    // packets with no error. The type is interface{} with range: 0..255.
    InPktsOk interface{}

    // packet not valid & validateFrames !strict. The type is interface{} with
    // range: 0..255.
    InPktsInvalid interface{}

    // packet not valid & validateFrames strict. The type is interface{} with
    // range: 0..255.
    InPktsNotValid interface{}

    // packet assigned to SA not in use & validateFrames strict. The type is
    // interface{} with range: 0..255.
    InPktsNotUsingSa interface{}

    // packet assigned to SA not in use& validateFrames !strict. The type is
    // interface{} with range: 0..255.
    InPktsUnusedSa interface{}

    // octets1 decrypted/validated. The type is interface{} with range: 0..255.
    InOctetsDecryptedValidated1 interface{}

    // octets validated. The type is interface{} with range: 0..255.
    InOctetsValidated interface{}

    // Xform Params.
    XformParams MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetFilter() yfilter.YFilter { return rxSa.YFilter }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) SetFilter(yf yfilter.YFilter) { rxSa.YFilter = yf }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetGoName(yname string) string {
    if yname == "is-valid" { return "IsValid" }
    if yname == "sa-id" { return "SaId" }
    if yname == "sc-no" { return "ScNo" }
    if yname == "in-pkts-unchecked" { return "InPktsUnchecked" }
    if yname == "in-pkts-delayed" { return "InPktsDelayed" }
    if yname == "in-pkts-late" { return "InPktsLate" }
    if yname == "in-pkts-ok" { return "InPktsOk" }
    if yname == "in-pkts-invalid" { return "InPktsInvalid" }
    if yname == "in-pkts-not-valid" { return "InPktsNotValid" }
    if yname == "in-pkts-not-using-sa" { return "InPktsNotUsingSa" }
    if yname == "in-pkts-unused-sa" { return "InPktsUnusedSa" }
    if yname == "in-octets-decrypted-validated1" { return "InOctetsDecryptedValidated1" }
    if yname == "in-octets-validated" { return "InOctetsValidated" }
    if yname == "xform-params" { return "XformParams" }
    return ""
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetSegmentPath() string {
    return "rx-sa"
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "xform-params" {
        return &rxSa.XformParams
    }
    return nil
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["xform-params"] = &rxSa.XformParams
    return children
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-valid"] = rxSa.IsValid
    leafs["sa-id"] = rxSa.SaId
    leafs["sc-no"] = rxSa.ScNo
    leafs["in-pkts-unchecked"] = rxSa.InPktsUnchecked
    leafs["in-pkts-delayed"] = rxSa.InPktsDelayed
    leafs["in-pkts-late"] = rxSa.InPktsLate
    leafs["in-pkts-ok"] = rxSa.InPktsOk
    leafs["in-pkts-invalid"] = rxSa.InPktsInvalid
    leafs["in-pkts-not-valid"] = rxSa.InPktsNotValid
    leafs["in-pkts-not-using-sa"] = rxSa.InPktsNotUsingSa
    leafs["in-pkts-unused-sa"] = rxSa.InPktsUnusedSa
    leafs["in-octets-decrypted-validated1"] = rxSa.InOctetsDecryptedValidated1
    leafs["in-octets-validated"] = rxSa.InOctetsValidated
    return leafs
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetBundleName() string { return "cisco_ios_xr" }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetYangName() string { return "rx-sa" }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) SetParent(parent types.Entity) { rxSa.parent = parent }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetParent() types.Entity { return rxSa.parent }

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetParentYangName() string { return "es200-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams
//  Xform Params
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // range of pkt nos considered valid. The type is interface{} with range:
    // 0..4294967295.
    ReplayWinSize interface{}

    // Cryptographic algo used. The type is string.
    CryptAlgo interface{}

    // APM_TRUE if this is Egress Transform record, APM_FALSE otherwise. The type
    // is bool.
    IsEgressTr interface{}

    // AES Key length. The type is string.
    AesKeyLen interface{}

    // Association Number for egress. The type is interface{} with range: 0..255.
    AssocNum interface{}

    // TRUE if Seq Num is 64-bit, FALSE if it is 32-bit. The type is bool.
    IsSeqNum64Bit interface{}

    // TRUE to generate the authKey, so authKey in this struct not used           
    // APM_FALSE to use provided authKey. The type is bool.
    BgenAuthKey interface{}
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetFilter() yfilter.YFilter { return xformParams.YFilter }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) SetFilter(yf yfilter.YFilter) { xformParams.YFilter = yf }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetGoName(yname string) string {
    if yname == "replay-win-size" { return "ReplayWinSize" }
    if yname == "crypt-algo" { return "CryptAlgo" }
    if yname == "is-egress-tr" { return "IsEgressTr" }
    if yname == "aes-key-len" { return "AesKeyLen" }
    if yname == "assoc-num" { return "AssocNum" }
    if yname == "is-seq-num64-bit" { return "IsSeqNum64Bit" }
    if yname == "bgen-auth-key" { return "BgenAuthKey" }
    return ""
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetSegmentPath() string {
    return "xform-params"
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["replay-win-size"] = xformParams.ReplayWinSize
    leafs["crypt-algo"] = xformParams.CryptAlgo
    leafs["is-egress-tr"] = xformParams.IsEgressTr
    leafs["aes-key-len"] = xformParams.AesKeyLen
    leafs["assoc-num"] = xformParams.AssocNum
    leafs["is-seq-num64-bit"] = xformParams.IsSeqNum64Bit
    leafs["bgen-auth-key"] = xformParams.BgenAuthKey
    return leafs
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetBundleName() string { return "cisco_ios_xr" }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetYangName() string { return "xform-params" }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) SetParent(parent types.Entity) { xformParams.parent = parent }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetParent() types.Entity { return xformParams.parent }

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetParentYangName() string { return "rx-sa" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS
// Table of Hardware Flows
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hardware Flow. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow.
    HwFlow []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow
}

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetFilter() yfilter.YFilter { return hwFlowS.YFilter }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) SetFilter(yf yfilter.YFilter) { hwFlowS.YFilter = yf }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetGoName(yname string) string {
    if yname == "hw-flow" { return "HwFlow" }
    return ""
}

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetSegmentPath() string {
    return "hw-flow-s"
}

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-flow" {
        for _, c := range hwFlowS.HwFlow {
            if hwFlowS.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow{}
        hwFlowS.HwFlow = append(hwFlowS.HwFlow, child)
        return &hwFlowS.HwFlow[len(hwFlowS.HwFlow)-1]
    }
    return nil
}

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwFlowS.HwFlow {
        children[hwFlowS.HwFlow[i].GetSegmentPath()] = &hwFlowS.HwFlow[i]
    }
    return children
}

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetBundleName() string { return "cisco_ios_xr" }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetYangName() string { return "hw-flow-s" }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) SetParent(parent types.Entity) { hwFlowS.parent = parent }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetParent() types.Entity { return hwFlowS.parent }

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetParentYangName() string { return "interface" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow
// Hardware Flow
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. FLOW ID. The type is interface{} with range:
    // -2147483648..2147483647.
    FlowId interface{}

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext
}

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetFilter() yfilter.YFilter { return hwFlow.YFilter }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) SetFilter(yf yfilter.YFilter) { hwFlow.YFilter = yf }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetGoName(yname string) string {
    if yname == "flow-id" { return "FlowId" }
    if yname == "ext" { return "Ext" }
    return ""
}

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetSegmentPath() string {
    return "hw-flow" + "[flow-id='" + fmt.Sprintf("%v", hwFlow.FlowId) + "']"
}

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext" {
        return &hwFlow.Ext
    }
    return nil
}

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ext"] = &hwFlow.Ext
    return children
}

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-id"] = hwFlow.FlowId
    return leafs
}

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetBundleName() string { return "cisco_ios_xr" }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetYangName() string { return "hw-flow" }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) SetParent(parent types.Entity) { hwFlow.parent = parent }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetParent() types.Entity { return hwFlow.parent }

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetParentYangName() string { return "hw-flow-s" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type interface{}

    // MSFPGA Flow Information.
    MsfpgaFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow

    // ES200 Flow Information.
    Es200Flow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetFilter() yfilter.YFilter { return ext.YFilter }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) SetFilter(yf yfilter.YFilter) { ext.YFilter = yf }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "msfpga-flow" { return "MsfpgaFlow" }
    if yname == "es200-flow" { return "Es200Flow" }
    return ""
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetSegmentPath() string {
    return "ext"
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msfpga-flow" {
        return &ext.MsfpgaFlow
    }
    if childYangName == "es200-flow" {
        return &ext.Es200Flow
    }
    return nil
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["msfpga-flow"] = &ext.MsfpgaFlow
    children["es200-flow"] = &ext.Es200Flow
    return children
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = ext.Type
    return leafs
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetBundleName() string { return "cisco_ios_xr" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetYangName() string { return "ext" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) SetParent(parent types.Entity) { ext.parent = parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetParent() types.Entity { return ext.parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetParentYangName() string { return "hw-flow" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow
// MSFPGA Flow Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Flow Details.
    TxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow

    // Rx Flow Details.
    RxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow
}

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetFilter() yfilter.YFilter { return msfpgaFlow.YFilter }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) SetFilter(yf yfilter.YFilter) { msfpgaFlow.YFilter = yf }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetGoName(yname string) string {
    if yname == "tx-flow" { return "TxFlow" }
    if yname == "rx-flow" { return "RxFlow" }
    return ""
}

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetSegmentPath() string {
    return "msfpga-flow"
}

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-flow" {
        return &msfpgaFlow.TxFlow
    }
    if childYangName == "rx-flow" {
        return &msfpgaFlow.RxFlow
    }
    return nil
}

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-flow"] = &msfpgaFlow.TxFlow
    children["rx-flow"] = &msfpgaFlow.RxFlow
    return children
}

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetBundleName() string { return "cisco_ios_xr" }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetYangName() string { return "msfpga-flow" }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) SetParent(parent types.Entity) { msfpgaFlow.parent = parent }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetParent() types.Entity { return msfpgaFlow.parent }

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow
// Tx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow Index. The type is interface{} with range: 0..255.
    FlowId interface{}

    // Flow Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // If MAC SA in Use. The type is bool.
    SmacInuse interface{}

    // If MAC DA in Use. The type is bool.
    DmacInuse interface{}

    // Ether Type. The type is interface{} with range: 0..65535.
    Ethertype interface{}

    // Outer VLAN ID. The type is interface{} with range: 0..65535.
    OuterVlan interface{}

    // Outer Vlan UserPri. The type is interface{} with range: 0..255.
    OuterVlanUp interface{}

    // Outer Vlan TPID. The type is interface{} with range: 0..65535.
    OuterVlanTpid interface{}

    // Inner VLAN ID. The type is interface{} with range: 0..65535.
    InnerVlan interface{}

    // Inner Vlan UserPri. The type is interface{} with range: 0..255.
    InnerVlanUp interface{}

    // Inner Vlan TPID. The type is interface{} with range: 0..65535.
    InnerVlanTpid interface{}

    // Source Port. The type is interface{} with range: 0..4294967295.
    SourcePort interface{}

    // Source Port ChkEn. The type is bool.
    SourcePortChk interface{}

    // If SCI in use. The type is bool.
    SciInuse interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // Match Priority. The type is interface{} with range: 0..255.
    MatchPri interface{}

    // Is Control Pkt. The type is bool.
    IsCtrlPkt interface{}

    // Ctrl Pkt ChkEn. The type is bool.
    CtrlCheck interface{}

    // MatchUntagged. The type is bool.
    MatchUntagged interface{}

    // MatchTagged. The type is bool.
    MatchTagged interface{}

    // Match Bad Tag. The type is bool.
    MatchBadTag interface{}

    // MatchKaYTag. The type is bool.
    MatchKayTag interface{}

    // TCI V. The type is interface{} with range: 0..255.
    TciV interface{}

    // TCI ES. The type is interface{} with range: 0..255.
    TciEXr interface{}

    // TCI SC. The type is interface{} with range: 0..255.
    TciSc interface{}

    // TCI SCB. The type is interface{} with range: 0..255.
    TciScb interface{}

    // TCI E. The type is interface{} with range: 0..255.
    Tci interface{}

    // TCI C. The type is interface{} with range: 0..255.
    TciC interface{}

    // TCI AN. The type is interface{} with range: 0..255.
    TciAn interface{}

    // TciAnChkEn. The type is bool.
    TciAnChk interface{}

    // TciChkEn. The type is bool.
    TciChk interface{}

    // SAI. The type is interface{} with range: 0..4294967295.
    Sai interface{}

    // MAC SA. The type is slice of interface{} with range: 0..255.
    Macsa []interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetFilter() yfilter.YFilter { return txFlow.YFilter }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) SetFilter(yf yfilter.YFilter) { txFlow.YFilter = yf }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetGoName(yname string) string {
    if yname == "flow-id" { return "FlowId" }
    if yname == "valid" { return "Valid" }
    if yname == "is-egress" { return "IsEgress" }
    if yname == "in-use" { return "InUse" }
    if yname == "action" { return "Action" }
    if yname == "smac-inuse" { return "SmacInuse" }
    if yname == "dmac-inuse" { return "DmacInuse" }
    if yname == "ethertype" { return "Ethertype" }
    if yname == "outer-vlan" { return "OuterVlan" }
    if yname == "outer-vlan-up" { return "OuterVlanUp" }
    if yname == "outer-vlan-tpid" { return "OuterVlanTpid" }
    if yname == "inner-vlan" { return "InnerVlan" }
    if yname == "inner-vlan-up" { return "InnerVlanUp" }
    if yname == "inner-vlan-tpid" { return "InnerVlanTpid" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "source-port-chk" { return "SourcePortChk" }
    if yname == "sci-inuse" { return "SciInuse" }
    if yname == "sci" { return "Sci" }
    if yname == "match-pri" { return "MatchPri" }
    if yname == "is-ctrl-pkt" { return "IsCtrlPkt" }
    if yname == "ctrl-check" { return "CtrlCheck" }
    if yname == "match-untagged" { return "MatchUntagged" }
    if yname == "match-tagged" { return "MatchTagged" }
    if yname == "match-bad-tag" { return "MatchBadTag" }
    if yname == "match-kay-tag" { return "MatchKayTag" }
    if yname == "tci-v" { return "TciV" }
    if yname == "tci-e-xr" { return "TciEXr" }
    if yname == "tci-sc" { return "TciSc" }
    if yname == "tci-scb" { return "TciScb" }
    if yname == "tci" { return "Tci" }
    if yname == "tci-c" { return "TciC" }
    if yname == "tci-an" { return "TciAn" }
    if yname == "tci-an-chk" { return "TciAnChk" }
    if yname == "tci-chk" { return "TciChk" }
    if yname == "sai" { return "Sai" }
    if yname == "macsa" { return "Macsa" }
    if yname == "macda" { return "Macda" }
    return ""
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetSegmentPath() string {
    return "tx-flow"
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-id"] = txFlow.FlowId
    leafs["valid"] = txFlow.Valid
    leafs["is-egress"] = txFlow.IsEgress
    leafs["in-use"] = txFlow.InUse
    leafs["action"] = txFlow.Action
    leafs["smac-inuse"] = txFlow.SmacInuse
    leafs["dmac-inuse"] = txFlow.DmacInuse
    leafs["ethertype"] = txFlow.Ethertype
    leafs["outer-vlan"] = txFlow.OuterVlan
    leafs["outer-vlan-up"] = txFlow.OuterVlanUp
    leafs["outer-vlan-tpid"] = txFlow.OuterVlanTpid
    leafs["inner-vlan"] = txFlow.InnerVlan
    leafs["inner-vlan-up"] = txFlow.InnerVlanUp
    leafs["inner-vlan-tpid"] = txFlow.InnerVlanTpid
    leafs["source-port"] = txFlow.SourcePort
    leafs["source-port-chk"] = txFlow.SourcePortChk
    leafs["sci-inuse"] = txFlow.SciInuse
    leafs["sci"] = txFlow.Sci
    leafs["match-pri"] = txFlow.MatchPri
    leafs["is-ctrl-pkt"] = txFlow.IsCtrlPkt
    leafs["ctrl-check"] = txFlow.CtrlCheck
    leafs["match-untagged"] = txFlow.MatchUntagged
    leafs["match-tagged"] = txFlow.MatchTagged
    leafs["match-bad-tag"] = txFlow.MatchBadTag
    leafs["match-kay-tag"] = txFlow.MatchKayTag
    leafs["tci-v"] = txFlow.TciV
    leafs["tci-e-xr"] = txFlow.TciEXr
    leafs["tci-sc"] = txFlow.TciSc
    leafs["tci-scb"] = txFlow.TciScb
    leafs["tci"] = txFlow.Tci
    leafs["tci-c"] = txFlow.TciC
    leafs["tci-an"] = txFlow.TciAn
    leafs["tci-an-chk"] = txFlow.TciAnChk
    leafs["tci-chk"] = txFlow.TciChk
    leafs["sai"] = txFlow.Sai
    leafs["macsa"] = txFlow.Macsa
    leafs["macda"] = txFlow.Macda
    return leafs
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetBundleName() string { return "cisco_ios_xr" }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetYangName() string { return "tx-flow" }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) SetParent(parent types.Entity) { txFlow.parent = parent }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetParent() types.Entity { return txFlow.parent }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetParentYangName() string { return "msfpga-flow" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow
// Rx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow Index. The type is interface{} with range: 0..255.
    FlowId interface{}

    // Flow Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // If MAC SA in Use. The type is bool.
    SmacInuse interface{}

    // If MAC DA in Use. The type is bool.
    DmacInuse interface{}

    // Ether Type. The type is interface{} with range: 0..65535.
    Ethertype interface{}

    // Outer VLAN ID. The type is interface{} with range: 0..65535.
    OuterVlan interface{}

    // Outer Vlan UserPri. The type is interface{} with range: 0..255.
    OuterVlanUp interface{}

    // Outer Vlan TPID. The type is interface{} with range: 0..65535.
    OuterVlanTpid interface{}

    // Inner VLAN ID. The type is interface{} with range: 0..65535.
    InnerVlan interface{}

    // Inner Vlan UserPri. The type is interface{} with range: 0..255.
    InnerVlanUp interface{}

    // Inner Vlan TPID. The type is interface{} with range: 0..65535.
    InnerVlanTpid interface{}

    // Source Port. The type is interface{} with range: 0..4294967295.
    SourcePort interface{}

    // Source Port ChkEn. The type is bool.
    SourcePortChk interface{}

    // If SCI in use. The type is bool.
    SciInuse interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // Match Priority. The type is interface{} with range: 0..255.
    MatchPri interface{}

    // Is Control Pkt. The type is bool.
    IsCtrlPkt interface{}

    // Ctrl Pkt ChkEn. The type is bool.
    CtrlCheck interface{}

    // MatchUntagged. The type is bool.
    MatchUntagged interface{}

    // MatchTagged. The type is bool.
    MatchTagged interface{}

    // Match Bad Tag. The type is bool.
    MatchBadTag interface{}

    // MatchKaYTag. The type is bool.
    MatchKayTag interface{}

    // TCI V. The type is interface{} with range: 0..255.
    TciV interface{}

    // TCI ES. The type is interface{} with range: 0..255.
    TciEXr interface{}

    // TCI SC. The type is interface{} with range: 0..255.
    TciSc interface{}

    // TCI SCB. The type is interface{} with range: 0..255.
    TciScb interface{}

    // TCI E. The type is interface{} with range: 0..255.
    Tci interface{}

    // TCI C. The type is interface{} with range: 0..255.
    TciC interface{}

    // TCI AN. The type is interface{} with range: 0..255.
    TciAn interface{}

    // TciAnChkEn. The type is bool.
    TciAnChk interface{}

    // TciChkEn. The type is bool.
    TciChk interface{}

    // SAI. The type is interface{} with range: 0..4294967295.
    Sai interface{}

    // MAC SA. The type is slice of interface{} with range: 0..255.
    Macsa []interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetFilter() yfilter.YFilter { return rxFlow.YFilter }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) SetFilter(yf yfilter.YFilter) { rxFlow.YFilter = yf }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetGoName(yname string) string {
    if yname == "flow-id" { return "FlowId" }
    if yname == "valid" { return "Valid" }
    if yname == "is-egress" { return "IsEgress" }
    if yname == "in-use" { return "InUse" }
    if yname == "action" { return "Action" }
    if yname == "smac-inuse" { return "SmacInuse" }
    if yname == "dmac-inuse" { return "DmacInuse" }
    if yname == "ethertype" { return "Ethertype" }
    if yname == "outer-vlan" { return "OuterVlan" }
    if yname == "outer-vlan-up" { return "OuterVlanUp" }
    if yname == "outer-vlan-tpid" { return "OuterVlanTpid" }
    if yname == "inner-vlan" { return "InnerVlan" }
    if yname == "inner-vlan-up" { return "InnerVlanUp" }
    if yname == "inner-vlan-tpid" { return "InnerVlanTpid" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "source-port-chk" { return "SourcePortChk" }
    if yname == "sci-inuse" { return "SciInuse" }
    if yname == "sci" { return "Sci" }
    if yname == "match-pri" { return "MatchPri" }
    if yname == "is-ctrl-pkt" { return "IsCtrlPkt" }
    if yname == "ctrl-check" { return "CtrlCheck" }
    if yname == "match-untagged" { return "MatchUntagged" }
    if yname == "match-tagged" { return "MatchTagged" }
    if yname == "match-bad-tag" { return "MatchBadTag" }
    if yname == "match-kay-tag" { return "MatchKayTag" }
    if yname == "tci-v" { return "TciV" }
    if yname == "tci-e-xr" { return "TciEXr" }
    if yname == "tci-sc" { return "TciSc" }
    if yname == "tci-scb" { return "TciScb" }
    if yname == "tci" { return "Tci" }
    if yname == "tci-c" { return "TciC" }
    if yname == "tci-an" { return "TciAn" }
    if yname == "tci-an-chk" { return "TciAnChk" }
    if yname == "tci-chk" { return "TciChk" }
    if yname == "sai" { return "Sai" }
    if yname == "macsa" { return "Macsa" }
    if yname == "macda" { return "Macda" }
    return ""
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetSegmentPath() string {
    return "rx-flow"
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-id"] = rxFlow.FlowId
    leafs["valid"] = rxFlow.Valid
    leafs["is-egress"] = rxFlow.IsEgress
    leafs["in-use"] = rxFlow.InUse
    leafs["action"] = rxFlow.Action
    leafs["smac-inuse"] = rxFlow.SmacInuse
    leafs["dmac-inuse"] = rxFlow.DmacInuse
    leafs["ethertype"] = rxFlow.Ethertype
    leafs["outer-vlan"] = rxFlow.OuterVlan
    leafs["outer-vlan-up"] = rxFlow.OuterVlanUp
    leafs["outer-vlan-tpid"] = rxFlow.OuterVlanTpid
    leafs["inner-vlan"] = rxFlow.InnerVlan
    leafs["inner-vlan-up"] = rxFlow.InnerVlanUp
    leafs["inner-vlan-tpid"] = rxFlow.InnerVlanTpid
    leafs["source-port"] = rxFlow.SourcePort
    leafs["source-port-chk"] = rxFlow.SourcePortChk
    leafs["sci-inuse"] = rxFlow.SciInuse
    leafs["sci"] = rxFlow.Sci
    leafs["match-pri"] = rxFlow.MatchPri
    leafs["is-ctrl-pkt"] = rxFlow.IsCtrlPkt
    leafs["ctrl-check"] = rxFlow.CtrlCheck
    leafs["match-untagged"] = rxFlow.MatchUntagged
    leafs["match-tagged"] = rxFlow.MatchTagged
    leafs["match-bad-tag"] = rxFlow.MatchBadTag
    leafs["match-kay-tag"] = rxFlow.MatchKayTag
    leafs["tci-v"] = rxFlow.TciV
    leafs["tci-e-xr"] = rxFlow.TciEXr
    leafs["tci-sc"] = rxFlow.TciSc
    leafs["tci-scb"] = rxFlow.TciScb
    leafs["tci"] = rxFlow.Tci
    leafs["tci-c"] = rxFlow.TciC
    leafs["tci-an"] = rxFlow.TciAn
    leafs["tci-an-chk"] = rxFlow.TciAnChk
    leafs["tci-chk"] = rxFlow.TciChk
    leafs["sai"] = rxFlow.Sai
    leafs["macsa"] = rxFlow.Macsa
    leafs["macda"] = rxFlow.Macda
    return leafs
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetBundleName() string { return "cisco_ios_xr" }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetYangName() string { return "rx-flow" }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) SetParent(parent types.Entity) { rxFlow.parent = parent }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetParent() types.Entity { return rxFlow.parent }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetParentYangName() string { return "msfpga-flow" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow
// ES200 Flow Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Flow Details.
    TxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow

    // Rx Flow Details.
    RxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow
}

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetFilter() yfilter.YFilter { return es200Flow.YFilter }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) SetFilter(yf yfilter.YFilter) { es200Flow.YFilter = yf }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetGoName(yname string) string {
    if yname == "tx-flow" { return "TxFlow" }
    if yname == "rx-flow" { return "RxFlow" }
    return ""
}

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetSegmentPath() string {
    return "es200-flow"
}

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-flow" {
        return &es200Flow.TxFlow
    }
    if childYangName == "rx-flow" {
        return &es200Flow.RxFlow
    }
    return nil
}

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-flow"] = &es200Flow.TxFlow
    children["rx-flow"] = &es200Flow.RxFlow
    return children
}

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetBundleName() string { return "cisco_ios_xr" }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetYangName() string { return "es200-flow" }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) SetParent(parent types.Entity) { es200Flow.parent = parent }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetParent() types.Entity { return es200Flow.parent }

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow
// Tx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow Number. The type is interface{} with range: 0..4294967295.
    FlowNo interface{}

    // Is Flow Enabled. The type is bool.
    IsFlowEnabled interface{}

    // Parsed EtherType to match could be 0 if Ethertype should'nt                
    // be matched can be 0x88E5 for MACSec tag. The type is interface{} with
    // range: 0..65535.
    Ethertype interface{}

    // VLAN ID for outer tag use this when             only one tag should be
    // matched. The type is interface{} with range: 0..65535.
    OuterVlanId interface{}

    // VLAN User Priority for outer tag  use            this when only one tag
    // should be matched. The type is interface{} with range: 0..255.
    OuterVlanUserPri interface{}

    // VLAN ID for inner tag used when two              VLAN Tags should be
    // matched. The type is interface{} with range: 0..65535.
    InnerVlanId interface{}

    // VLAN User priority for inner tag use            when matching two VLAN
    // tags. The type is interface{} with range: 0..255.
    InnerVlanUserPri interface{}

    // SCI to be matched value required for            ingress only, pass NULL for
    // egress. The type is interface{} with range: 0..18446744073709551615.
    Psci interface{}

    // priority for match 0-15(highest) . The type is interface{} with range:
    // 0..255.
    MatchPriority interface{}

    // value of 'v' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciV interface{}

    // value of 'es' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciEXr interface{}

    // value of 'sc' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciSc interface{}

    // value of 'scb' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciScb interface{}

    // value of 'e' in TCI to match (1bit ). The type is interface{} with range:
    // 0..255.
    Tci interface{}

    // value of 'c' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciC interface{}

    // TCI bits will be checked only when this          bit is enabled. All the
    // values of TCI bits       are mandatory when TCI check is used. The type is
    // bool.
    TciChk interface{}

    // Type of packet. See ethMscCfyEPktType_e. The type is string.
    PktType interface{}

    // No. of MPLS or VLAN tags See ethMscCfyETagNum_e . The type is string.
    TagNum interface{}

    // Dei to match for innner Vlan tag. The type is bool.
    InnerVlanDei interface{}

    // Dei to match for outer Vlan tag. The type is bool.
    OuterVlanDei interface{}

    // Service Instance id . The type is interface{} with range: 0..4294967295.
    PbbSid interface{}

    // Backbone Vlan id . The type is interface{} with range: 0..4294967295.
    PbbBvid interface{}

    // pcp . The type is interface{} with range: 0..255.
    PbbPcp interface{}

    // dei . The type is interface{} with range: 0..255.
    PbbDei interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls1Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls1Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls1Bos interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls2Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls2Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls2Bos interface{}

    // Plain bits to compare. Max values:               untagged pkt - 40 bits
    // after EthType             1 VLAN tag - 24 bits after parsed EthType       
    // 2 VLAN tags- 8 bits after parsed EthType         1 MPLS tag - 32 bits after
    // 1st tag               2 MPLS tags- 8 bits following after 2nd          or
    // atmost 5th MPLS tag                           PBB - 16 bits after C-SA     
    // PBB with VLAN tag - 16 bits of VLAN tag . The type is interface{} with
    // range: 0..18446744073709551615. Units are bit.
    PlainBits interface{}

    // No. of bits used in plainBits. The type is interface{} with range: 0..255.
    PlainBitsSize interface{}

    // Force the pkt as control pkt irrepective         of the results of control
    // packet detector. The type is bool.
    ForceCtrl interface{}

    // Drop the packet. The type is bool.
    Drop interface{}

    // DA mask. The type is interface{} with range: 0..18446744073709551615.
    MaskDa interface{}

    // Parsed EtherType mask. The type is interface{} with range: 0..4294967295.
    MaskEthertype interface{}

    // Plain Bits mask. The type is interface{} with range:
    // 0..18446744073709551615.
    MaskPlainBits interface{}

    // Pkts matching the Flow. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowHits interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetFilter() yfilter.YFilter { return txFlow.YFilter }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) SetFilter(yf yfilter.YFilter) { txFlow.YFilter = yf }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetGoName(yname string) string {
    if yname == "flow-no" { return "FlowNo" }
    if yname == "is-flow-enabled" { return "IsFlowEnabled" }
    if yname == "ethertype" { return "Ethertype" }
    if yname == "outer-vlan-id" { return "OuterVlanId" }
    if yname == "outer-vlan-user-pri" { return "OuterVlanUserPri" }
    if yname == "inner-vlan-id" { return "InnerVlanId" }
    if yname == "inner-vlan-user-pri" { return "InnerVlanUserPri" }
    if yname == "psci" { return "Psci" }
    if yname == "match-priority" { return "MatchPriority" }
    if yname == "tci-v" { return "TciV" }
    if yname == "tci-e-xr" { return "TciEXr" }
    if yname == "tci-sc" { return "TciSc" }
    if yname == "tci-scb" { return "TciScb" }
    if yname == "tci" { return "Tci" }
    if yname == "tci-c" { return "TciC" }
    if yname == "tci-chk" { return "TciChk" }
    if yname == "pkt-type" { return "PktType" }
    if yname == "tag-num" { return "TagNum" }
    if yname == "inner-vlan-dei" { return "InnerVlanDei" }
    if yname == "outer-vlan-dei" { return "OuterVlanDei" }
    if yname == "pbb-sid" { return "PbbSid" }
    if yname == "pbb-bvid" { return "PbbBvid" }
    if yname == "pbb-pcp" { return "PbbPcp" }
    if yname == "pbb-dei" { return "PbbDei" }
    if yname == "mpls1-label" { return "Mpls1Label" }
    if yname == "mpls1-exp" { return "Mpls1Exp" }
    if yname == "mpls1-bos" { return "Mpls1Bos" }
    if yname == "mpls2-label" { return "Mpls2Label" }
    if yname == "mpls2-exp" { return "Mpls2Exp" }
    if yname == "mpls2-bos" { return "Mpls2Bos" }
    if yname == "plain-bits" { return "PlainBits" }
    if yname == "plain-bits-size" { return "PlainBitsSize" }
    if yname == "force-ctrl" { return "ForceCtrl" }
    if yname == "drop" { return "Drop" }
    if yname == "mask-da" { return "MaskDa" }
    if yname == "mask-ethertype" { return "MaskEthertype" }
    if yname == "mask-plain-bits" { return "MaskPlainBits" }
    if yname == "flow-hits" { return "FlowHits" }
    if yname == "macda" { return "Macda" }
    return ""
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetSegmentPath() string {
    return "tx-flow"
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-no"] = txFlow.FlowNo
    leafs["is-flow-enabled"] = txFlow.IsFlowEnabled
    leafs["ethertype"] = txFlow.Ethertype
    leafs["outer-vlan-id"] = txFlow.OuterVlanId
    leafs["outer-vlan-user-pri"] = txFlow.OuterVlanUserPri
    leafs["inner-vlan-id"] = txFlow.InnerVlanId
    leafs["inner-vlan-user-pri"] = txFlow.InnerVlanUserPri
    leafs["psci"] = txFlow.Psci
    leafs["match-priority"] = txFlow.MatchPriority
    leafs["tci-v"] = txFlow.TciV
    leafs["tci-e-xr"] = txFlow.TciEXr
    leafs["tci-sc"] = txFlow.TciSc
    leafs["tci-scb"] = txFlow.TciScb
    leafs["tci"] = txFlow.Tci
    leafs["tci-c"] = txFlow.TciC
    leafs["tci-chk"] = txFlow.TciChk
    leafs["pkt-type"] = txFlow.PktType
    leafs["tag-num"] = txFlow.TagNum
    leafs["inner-vlan-dei"] = txFlow.InnerVlanDei
    leafs["outer-vlan-dei"] = txFlow.OuterVlanDei
    leafs["pbb-sid"] = txFlow.PbbSid
    leafs["pbb-bvid"] = txFlow.PbbBvid
    leafs["pbb-pcp"] = txFlow.PbbPcp
    leafs["pbb-dei"] = txFlow.PbbDei
    leafs["mpls1-label"] = txFlow.Mpls1Label
    leafs["mpls1-exp"] = txFlow.Mpls1Exp
    leafs["mpls1-bos"] = txFlow.Mpls1Bos
    leafs["mpls2-label"] = txFlow.Mpls2Label
    leafs["mpls2-exp"] = txFlow.Mpls2Exp
    leafs["mpls2-bos"] = txFlow.Mpls2Bos
    leafs["plain-bits"] = txFlow.PlainBits
    leafs["plain-bits-size"] = txFlow.PlainBitsSize
    leafs["force-ctrl"] = txFlow.ForceCtrl
    leafs["drop"] = txFlow.Drop
    leafs["mask-da"] = txFlow.MaskDa
    leafs["mask-ethertype"] = txFlow.MaskEthertype
    leafs["mask-plain-bits"] = txFlow.MaskPlainBits
    leafs["flow-hits"] = txFlow.FlowHits
    leafs["macda"] = txFlow.Macda
    return leafs
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetBundleName() string { return "cisco_ios_xr" }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetYangName() string { return "tx-flow" }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) SetParent(parent types.Entity) { txFlow.parent = parent }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetParent() types.Entity { return txFlow.parent }

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetParentYangName() string { return "es200-flow" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow
// Rx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow Number. The type is interface{} with range: 0..4294967295.
    FlowNo interface{}

    // Is Flow Enabled. The type is bool.
    IsFlowEnabled interface{}

    // Parsed EtherType to match could be 0 if Ethertype should'nt                
    // be matched can be 0x88E5 for MACSec tag. The type is interface{} with
    // range: 0..65535.
    Ethertype interface{}

    // VLAN ID for outer tag use this when             only one tag should be
    // matched. The type is interface{} with range: 0..65535.
    OuterVlanId interface{}

    // VLAN User Priority for outer tag  use            this when only one tag
    // should be matched. The type is interface{} with range: 0..255.
    OuterVlanUserPri interface{}

    // VLAN ID for inner tag used when two              VLAN Tags should be
    // matched. The type is interface{} with range: 0..65535.
    InnerVlanId interface{}

    // VLAN User priority for inner tag use            when matching two VLAN
    // tags. The type is interface{} with range: 0..255.
    InnerVlanUserPri interface{}

    // SCI to be matched value required for            ingress only, pass NULL for
    // egress. The type is interface{} with range: 0..18446744073709551615.
    Psci interface{}

    // priority for match 0-15(highest) . The type is interface{} with range:
    // 0..255.
    MatchPriority interface{}

    // value of 'v' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciV interface{}

    // value of 'es' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciEXr interface{}

    // value of 'sc' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciSc interface{}

    // value of 'scb' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciScb interface{}

    // value of 'e' in TCI to match (1bit ). The type is interface{} with range:
    // 0..255.
    Tci interface{}

    // value of 'c' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciC interface{}

    // TCI bits will be checked only when this          bit is enabled. All the
    // values of TCI bits       are mandatory when TCI check is used. The type is
    // bool.
    TciChk interface{}

    // Type of packet. See ethMscCfyEPktType_e. The type is string.
    PktType interface{}

    // No. of MPLS or VLAN tags See ethMscCfyETagNum_e . The type is string.
    TagNum interface{}

    // Dei to match for innner Vlan tag. The type is bool.
    InnerVlanDei interface{}

    // Dei to match for outer Vlan tag. The type is bool.
    OuterVlanDei interface{}

    // Service Instance id . The type is interface{} with range: 0..4294967295.
    PbbSid interface{}

    // Backbone Vlan id . The type is interface{} with range: 0..4294967295.
    PbbBvid interface{}

    // pcp . The type is interface{} with range: 0..255.
    PbbPcp interface{}

    // dei . The type is interface{} with range: 0..255.
    PbbDei interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls1Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls1Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls1Bos interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls2Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls2Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls2Bos interface{}

    // Plain bits to compare. Max values:               untagged pkt - 40 bits
    // after EthType             1 VLAN tag - 24 bits after parsed EthType       
    // 2 VLAN tags- 8 bits after parsed EthType         1 MPLS tag - 32 bits after
    // 1st tag               2 MPLS tags- 8 bits following after 2nd          or
    // atmost 5th MPLS tag                           PBB - 16 bits after C-SA     
    // PBB with VLAN tag - 16 bits of VLAN tag . The type is interface{} with
    // range: 0..18446744073709551615. Units are bit.
    PlainBits interface{}

    // No. of bits used in plainBits. The type is interface{} with range: 0..255.
    PlainBitsSize interface{}

    // Force the pkt as control pkt irrepective         of the results of control
    // packet detector. The type is bool.
    ForceCtrl interface{}

    // Drop the packet. The type is bool.
    Drop interface{}

    // DA mask. The type is interface{} with range: 0..18446744073709551615.
    MaskDa interface{}

    // Parsed EtherType mask. The type is interface{} with range: 0..4294967295.
    MaskEthertype interface{}

    // Plain Bits mask. The type is interface{} with range:
    // 0..18446744073709551615.
    MaskPlainBits interface{}

    // Pkts matching the Flow. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowHits interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetFilter() yfilter.YFilter { return rxFlow.YFilter }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) SetFilter(yf yfilter.YFilter) { rxFlow.YFilter = yf }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetGoName(yname string) string {
    if yname == "flow-no" { return "FlowNo" }
    if yname == "is-flow-enabled" { return "IsFlowEnabled" }
    if yname == "ethertype" { return "Ethertype" }
    if yname == "outer-vlan-id" { return "OuterVlanId" }
    if yname == "outer-vlan-user-pri" { return "OuterVlanUserPri" }
    if yname == "inner-vlan-id" { return "InnerVlanId" }
    if yname == "inner-vlan-user-pri" { return "InnerVlanUserPri" }
    if yname == "psci" { return "Psci" }
    if yname == "match-priority" { return "MatchPriority" }
    if yname == "tci-v" { return "TciV" }
    if yname == "tci-e-xr" { return "TciEXr" }
    if yname == "tci-sc" { return "TciSc" }
    if yname == "tci-scb" { return "TciScb" }
    if yname == "tci" { return "Tci" }
    if yname == "tci-c" { return "TciC" }
    if yname == "tci-chk" { return "TciChk" }
    if yname == "pkt-type" { return "PktType" }
    if yname == "tag-num" { return "TagNum" }
    if yname == "inner-vlan-dei" { return "InnerVlanDei" }
    if yname == "outer-vlan-dei" { return "OuterVlanDei" }
    if yname == "pbb-sid" { return "PbbSid" }
    if yname == "pbb-bvid" { return "PbbBvid" }
    if yname == "pbb-pcp" { return "PbbPcp" }
    if yname == "pbb-dei" { return "PbbDei" }
    if yname == "mpls1-label" { return "Mpls1Label" }
    if yname == "mpls1-exp" { return "Mpls1Exp" }
    if yname == "mpls1-bos" { return "Mpls1Bos" }
    if yname == "mpls2-label" { return "Mpls2Label" }
    if yname == "mpls2-exp" { return "Mpls2Exp" }
    if yname == "mpls2-bos" { return "Mpls2Bos" }
    if yname == "plain-bits" { return "PlainBits" }
    if yname == "plain-bits-size" { return "PlainBitsSize" }
    if yname == "force-ctrl" { return "ForceCtrl" }
    if yname == "drop" { return "Drop" }
    if yname == "mask-da" { return "MaskDa" }
    if yname == "mask-ethertype" { return "MaskEthertype" }
    if yname == "mask-plain-bits" { return "MaskPlainBits" }
    if yname == "flow-hits" { return "FlowHits" }
    if yname == "macda" { return "Macda" }
    return ""
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetSegmentPath() string {
    return "rx-flow"
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-no"] = rxFlow.FlowNo
    leafs["is-flow-enabled"] = rxFlow.IsFlowEnabled
    leafs["ethertype"] = rxFlow.Ethertype
    leafs["outer-vlan-id"] = rxFlow.OuterVlanId
    leafs["outer-vlan-user-pri"] = rxFlow.OuterVlanUserPri
    leafs["inner-vlan-id"] = rxFlow.InnerVlanId
    leafs["inner-vlan-user-pri"] = rxFlow.InnerVlanUserPri
    leafs["psci"] = rxFlow.Psci
    leafs["match-priority"] = rxFlow.MatchPriority
    leafs["tci-v"] = rxFlow.TciV
    leafs["tci-e-xr"] = rxFlow.TciEXr
    leafs["tci-sc"] = rxFlow.TciSc
    leafs["tci-scb"] = rxFlow.TciScb
    leafs["tci"] = rxFlow.Tci
    leafs["tci-c"] = rxFlow.TciC
    leafs["tci-chk"] = rxFlow.TciChk
    leafs["pkt-type"] = rxFlow.PktType
    leafs["tag-num"] = rxFlow.TagNum
    leafs["inner-vlan-dei"] = rxFlow.InnerVlanDei
    leafs["outer-vlan-dei"] = rxFlow.OuterVlanDei
    leafs["pbb-sid"] = rxFlow.PbbSid
    leafs["pbb-bvid"] = rxFlow.PbbBvid
    leafs["pbb-pcp"] = rxFlow.PbbPcp
    leafs["pbb-dei"] = rxFlow.PbbDei
    leafs["mpls1-label"] = rxFlow.Mpls1Label
    leafs["mpls1-exp"] = rxFlow.Mpls1Exp
    leafs["mpls1-bos"] = rxFlow.Mpls1Bos
    leafs["mpls2-label"] = rxFlow.Mpls2Label
    leafs["mpls2-exp"] = rxFlow.Mpls2Exp
    leafs["mpls2-bos"] = rxFlow.Mpls2Bos
    leafs["plain-bits"] = rxFlow.PlainBits
    leafs["plain-bits-size"] = rxFlow.PlainBitsSize
    leafs["force-ctrl"] = rxFlow.ForceCtrl
    leafs["drop"] = rxFlow.Drop
    leafs["mask-da"] = rxFlow.MaskDa
    leafs["mask-ethertype"] = rxFlow.MaskEthertype
    leafs["mask-plain-bits"] = rxFlow.MaskPlainBits
    leafs["flow-hits"] = rxFlow.FlowHits
    leafs["macda"] = rxFlow.Macda
    return leafs
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetBundleName() string { return "cisco_ios_xr" }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetYangName() string { return "rx-flow" }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) SetParent(parent types.Entity) { rxFlow.parent = parent }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetParent() types.Entity { return rxFlow.parent }

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetParentYangName() string { return "es200-flow" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics
// The Software Statistics
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext
}

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetFilter() yfilter.YFilter { return swStatistics.YFilter }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) SetFilter(yf yfilter.YFilter) { swStatistics.YFilter = yf }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetGoName(yname string) string {
    if yname == "ext" { return "Ext" }
    return ""
}

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetSegmentPath() string {
    return "sw-statistics"
}

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext" {
        return &swStatistics.Ext
    }
    return nil
}

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ext"] = &swStatistics.Ext
    return children
}

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetYangName() string { return "sw-statistics" }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) SetParent(parent types.Entity) { swStatistics.parent = parent }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetParent() types.Entity { return swStatistics.parent }

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetParentYangName() string { return "interface" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type interface{}

    // MSFPGA Stats.
    MsfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats

    // XLFPGA Stats.
    XlfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats

    // ES200 Stats.
    Es200Stats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetFilter() yfilter.YFilter { return ext.YFilter }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) SetFilter(yf yfilter.YFilter) { ext.YFilter = yf }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "msfpga-stats" { return "MsfpgaStats" }
    if yname == "xlfpga-stats" { return "XlfpgaStats" }
    if yname == "es200-stats" { return "Es200Stats" }
    return ""
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetSegmentPath() string {
    return "ext"
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "msfpga-stats" {
        return &ext.MsfpgaStats
    }
    if childYangName == "xlfpga-stats" {
        return &ext.XlfpgaStats
    }
    if childYangName == "es200-stats" {
        return &ext.Es200Stats
    }
    return nil
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["msfpga-stats"] = &ext.MsfpgaStats
    children["xlfpga-stats"] = &ext.XlfpgaStats
    children["es200-stats"] = &ext.Es200Stats
    return children
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = ext.Type
    return leafs
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetBundleName() string { return "cisco_ios_xr" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetYangName() string { return "ext" }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) SetParent(parent types.Entity) { ext.parent = parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetParent() types.Entity { return ext.parent }

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetParentYangName() string { return "sw-statistics" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats
// MSFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetFilter() yfilter.YFilter { return msfpgaStats.YFilter }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) SetFilter(yf yfilter.YFilter) { msfpgaStats.YFilter = yf }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetGoName(yname string) string {
    if yname == "tx-sa-stats" { return "TxSaStats" }
    if yname == "rx-sa-stats" { return "RxSaStats" }
    if yname == "tx-interface-macsec-stats" { return "TxInterfaceMacsecStats" }
    if yname == "rx-interface-macsec-stats" { return "RxInterfaceMacsecStats" }
    return ""
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetSegmentPath() string {
    return "msfpga-stats"
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-sa-stats" {
        return &msfpgaStats.TxSaStats
    }
    if childYangName == "rx-sa-stats" {
        return &msfpgaStats.RxSaStats
    }
    if childYangName == "tx-interface-macsec-stats" {
        return &msfpgaStats.TxInterfaceMacsecStats
    }
    if childYangName == "rx-interface-macsec-stats" {
        return &msfpgaStats.RxInterfaceMacsecStats
    }
    return nil
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-sa-stats"] = &msfpgaStats.TxSaStats
    children["rx-sa-stats"] = &msfpgaStats.RxSaStats
    children["tx-interface-macsec-stats"] = &msfpgaStats.TxInterfaceMacsecStats
    children["rx-interface-macsec-stats"] = &msfpgaStats.RxInterfaceMacsecStats
    return children
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetBundleName() string { return "cisco_ios_xr" }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetYangName() string { return "msfpga-stats" }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) SetParent(parent types.Entity) { msfpgaStats.parent = parent }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetParent() types.Entity { return msfpgaStats.parent }

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Pkts Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsProtected interface{}

    // Tx Pkts Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncrypted interface{}

    // Tx Octets Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsProtected interface{}

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncrypted interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetFilter() yfilter.YFilter { return txSaStats.YFilter }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) SetFilter(yf yfilter.YFilter) { txSaStats.YFilter = yf }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetGoName(yname string) string {
    if yname == "out-pkts-protected" { return "OutPktsProtected" }
    if yname == "out-pkts-encrypted" { return "OutPktsEncrypted" }
    if yname == "out-octets-protected" { return "OutOctetsProtected" }
    if yname == "out-octets-encrypted" { return "OutOctetsEncrypted" }
    return ""
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetSegmentPath() string {
    return "tx-sa-stats"
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkts-protected"] = txSaStats.OutPktsProtected
    leafs["out-pkts-encrypted"] = txSaStats.OutPktsEncrypted
    leafs["out-octets-protected"] = txSaStats.OutOctetsProtected
    leafs["out-octets-encrypted"] = txSaStats.OutOctetsEncrypted
    return leafs
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetYangName() string { return "tx-sa-stats" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) SetParent(parent types.Entity) { txSaStats.parent = parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetParent() types.Entity { return txSaStats.parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // Rx Pkts Not Valid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotValid interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsInvalid interface{}

    // Rx Pkts OK. The type is interface{} with range: 0..18446744073709551615.
    InPktsOk interface{}

    // Rx Pkts Delayed. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsDelayed interface{}

    // Rx Pkts Late. The type is interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // Rx Pkts Unchecked. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnchecked interface{}

    // Rx Octets Validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecrypted interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetFilter() yfilter.YFilter { return rxSaStats.YFilter }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) SetFilter(yf yfilter.YFilter) { rxSaStats.YFilter = yf }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetGoName(yname string) string {
    if yname == "in-pkts-unused-sa" { return "InPktsUnusedSa" }
    if yname == "in-pkts-not-using-sa" { return "InPktsNotUsingSa" }
    if yname == "in-pkts-not-valid" { return "InPktsNotValid" }
    if yname == "in-pkts-invalid" { return "InPktsInvalid" }
    if yname == "in-pkts-ok" { return "InPktsOk" }
    if yname == "in-pkts-delayed" { return "InPktsDelayed" }
    if yname == "in-pkts-late" { return "InPktsLate" }
    if yname == "in-pkts-unchecked" { return "InPktsUnchecked" }
    if yname == "in-octets-validated" { return "InOctetsValidated" }
    if yname == "in-octets-decrypted" { return "InOctetsDecrypted" }
    return ""
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetSegmentPath() string {
    return "rx-sa-stats"
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-unused-sa"] = rxSaStats.InPktsUnusedSa
    leafs["in-pkts-not-using-sa"] = rxSaStats.InPktsNotUsingSa
    leafs["in-pkts-not-valid"] = rxSaStats.InPktsNotValid
    leafs["in-pkts-invalid"] = rxSaStats.InPktsInvalid
    leafs["in-pkts-ok"] = rxSaStats.InPktsOk
    leafs["in-pkts-delayed"] = rxSaStats.InPktsDelayed
    leafs["in-pkts-late"] = rxSaStats.InPktsLate
    leafs["in-pkts-unchecked"] = rxSaStats.InPktsUnchecked
    leafs["in-octets-validated"] = rxSaStats.InOctetsValidated
    leafs["in-octets-decrypted"] = rxSaStats.InOctetsDecrypted
    return leafs
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetYangName() string { return "rx-sa-stats" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) SetParent(parent types.Entity) { rxSaStats.parent = parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetParent() types.Entity { return rxSaStats.parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUncontrolled interface{}

    // Tx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUntagged interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktTooLong interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return txInterfaceMacsecStats.YFilter }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { txInterfaceMacsecStats.YFilter = yf }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "out-pkt-uncontrolled" { return "OutPktUncontrolled" }
    if yname == "out-pkt-untagged" { return "OutPktUntagged" }
    if yname == "out-pkt-too-long" { return "OutPktTooLong" }
    return ""
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetSegmentPath() string {
    return "tx-interface-macsec-stats"
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkt-uncontrolled"] = txInterfaceMacsecStats.OutPktUncontrolled
    leafs["out-pkt-untagged"] = txInterfaceMacsecStats.OutPktUntagged
    leafs["out-pkt-too-long"] = txInterfaceMacsecStats.OutPktTooLong
    return leafs
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetYangName() string { return "tx-interface-macsec-stats" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) SetParent(parent types.Entity) { txInterfaceMacsecStats.parent = parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetParent() types.Entity { return txInterfaceMacsecStats.parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUntagged interface{}

    // Rx Pkts Notag. The type is interface{} with range: 0..18446744073709551615.
    InPktNotag interface{}

    // Rx Pkts Bad tag. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktBadTag interface{}

    // Rx Pkts No Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // Rx Pkts Unknown Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUnknownSci interface{}

    // Rx Pkts Tagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktTagged interface{}

    // Rx Pkts Over Run. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktOverrun interface{}

    // Rx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUncontrolled interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return rxInterfaceMacsecStats.YFilter }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { rxInterfaceMacsecStats.YFilter = yf }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "in-pkt-untagged" { return "InPktUntagged" }
    if yname == "in-pkt-notag" { return "InPktNotag" }
    if yname == "in-pkt-bad-tag" { return "InPktBadTag" }
    if yname == "in-pkt-no-sci" { return "InPktNoSci" }
    if yname == "in-pkt-unknown-sci" { return "InPktUnknownSci" }
    if yname == "in-pkt-tagged" { return "InPktTagged" }
    if yname == "in-pkt-overrun" { return "InPktOverrun" }
    if yname == "in-pkt-uncontrolled" { return "InPktUncontrolled" }
    return ""
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetSegmentPath() string {
    return "rx-interface-macsec-stats"
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkt-untagged"] = rxInterfaceMacsecStats.InPktUntagged
    leafs["in-pkt-notag"] = rxInterfaceMacsecStats.InPktNotag
    leafs["in-pkt-bad-tag"] = rxInterfaceMacsecStats.InPktBadTag
    leafs["in-pkt-no-sci"] = rxInterfaceMacsecStats.InPktNoSci
    leafs["in-pkt-unknown-sci"] = rxInterfaceMacsecStats.InPktUnknownSci
    leafs["in-pkt-tagged"] = rxInterfaceMacsecStats.InPktTagged
    leafs["in-pkt-overrun"] = rxInterfaceMacsecStats.InPktOverrun
    leafs["in-pkt-uncontrolled"] = rxInterfaceMacsecStats.InPktUncontrolled
    return leafs
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetYangName() string { return "rx-interface-macsec-stats" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) SetParent(parent types.Entity) { rxInterfaceMacsecStats.parent = parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetParent() types.Entity { return rxInterfaceMacsecStats.parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetParentYangName() string { return "msfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats
// XLFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SC and SA Level Stats.
    MacsecTxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats

    // Rx SC and SA Level Stats.
    MacsecRxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetFilter() yfilter.YFilter { return xlfpgaStats.YFilter }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) SetFilter(yf yfilter.YFilter) { xlfpgaStats.YFilter = yf }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetGoName(yname string) string {
    if yname == "macsec-tx-stats" { return "MacsecTxStats" }
    if yname == "macsec-rx-stats" { return "MacsecRxStats" }
    return ""
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetSegmentPath() string {
    return "xlfpga-stats"
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macsec-tx-stats" {
        return &xlfpgaStats.MacsecTxStats
    }
    if childYangName == "macsec-rx-stats" {
        return &xlfpgaStats.MacsecRxStats
    }
    return nil
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["macsec-tx-stats"] = &xlfpgaStats.MacsecTxStats
    children["macsec-rx-stats"] = &xlfpgaStats.MacsecRxStats
    return children
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetBundleName() string { return "cisco_ios_xr" }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetYangName() string { return "xlfpga-stats" }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) SetParent(parent types.Entity) { xlfpgaStats.parent = parent }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetParent() types.Entity { return xlfpgaStats.parent }

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats
// Tx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedOctets interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    ScToolongPkts interface{}

    // Tx packets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedPkts interface{}

    // Tx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Tx Overrun Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Tx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Tx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Tx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Current Tx AN. The type is interface{} with range: 0..18446744073709551615.
    CurrentAn interface{}

    // Current Tx SA Encrypted Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    SaEncryptedPkts interface{}
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetFilter() yfilter.YFilter { return macsecTxStats.YFilter }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) SetFilter(yf yfilter.YFilter) { macsecTxStats.YFilter = yf }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetGoName(yname string) string {
    if yname == "sc-encrypted-octets" { return "ScEncryptedOctets" }
    if yname == "sc-toolong-pkts" { return "ScToolongPkts" }
    if yname == "sc-encrypted-pkts" { return "ScEncryptedPkts" }
    if yname == "sc-untagged-pkts" { return "ScUntaggedPkts" }
    if yname == "sc-overrun-pkts" { return "ScOverrunPkts" }
    if yname == "sc-bypass-pkts" { return "ScBypassPkts" }
    if yname == "sc-eapol-pkts" { return "ScEapolPkts" }
    if yname == "sc-dropped-pkts" { return "ScDroppedPkts" }
    if yname == "current-an" { return "CurrentAn" }
    if yname == "sa-encrypted-pkts" { return "SaEncryptedPkts" }
    return ""
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetSegmentPath() string {
    return "macsec-tx-stats"
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sc-encrypted-octets"] = macsecTxStats.ScEncryptedOctets
    leafs["sc-toolong-pkts"] = macsecTxStats.ScToolongPkts
    leafs["sc-encrypted-pkts"] = macsecTxStats.ScEncryptedPkts
    leafs["sc-untagged-pkts"] = macsecTxStats.ScUntaggedPkts
    leafs["sc-overrun-pkts"] = macsecTxStats.ScOverrunPkts
    leafs["sc-bypass-pkts"] = macsecTxStats.ScBypassPkts
    leafs["sc-eapol-pkts"] = macsecTxStats.ScEapolPkts
    leafs["sc-dropped-pkts"] = macsecTxStats.ScDroppedPkts
    leafs["current-an"] = macsecTxStats.CurrentAn
    leafs["sa-encrypted-pkts"] = macsecTxStats.SaEncryptedPkts
    return leafs
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetBundleName() string { return "cisco_ios_xr" }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetYangName() string { return "macsec-tx-stats" }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) SetParent(parent types.Entity) { macsecTxStats.parent = parent }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetParent() types.Entity { return macsecTxStats.parent }

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetParentYangName() string { return "xlfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats
// Rx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDecryptedOctets interface{}

    // Rx No Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoTagPkts interface{}

    // Rx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Rx Bad Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBadTagPkts interface{}

    // Rx Late Pkts. The type is interface{} with range: 0..18446744073709551615.
    ScLatePkts interface{}

    // Rx Delayed Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDelayedPkts interface{}

    // Rx Unchecked Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUncheckedPkts interface{}

    // Rx No SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoSciPkts interface{}

    // Rx Unknown SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnknownSciPkts interface{}

    // Rx Pkts Ok. The type is interface{} with range: 0..18446744073709551615.
    ScOkPkts interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotUsingPkts interface{}

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnusedPkts interface{}

    // Rx Not Valid Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotValidPkts interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    ScInvalidPkts interface{}

    // Rx Overrun Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Rx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Rx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Rx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Rx SA Level Stats. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat.
    RxSaStat []MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetFilter() yfilter.YFilter { return macsecRxStats.YFilter }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) SetFilter(yf yfilter.YFilter) { macsecRxStats.YFilter = yf }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetGoName(yname string) string {
    if yname == "sc-decrypted-octets" { return "ScDecryptedOctets" }
    if yname == "sc-no-tag-pkts" { return "ScNoTagPkts" }
    if yname == "sc-untagged-pkts" { return "ScUntaggedPkts" }
    if yname == "sc-bad-tag-pkts" { return "ScBadTagPkts" }
    if yname == "sc-late-pkts" { return "ScLatePkts" }
    if yname == "sc-delayed-pkts" { return "ScDelayedPkts" }
    if yname == "sc-unchecked-pkts" { return "ScUncheckedPkts" }
    if yname == "sc-no-sci-pkts" { return "ScNoSciPkts" }
    if yname == "sc-unknown-sci-pkts" { return "ScUnknownSciPkts" }
    if yname == "sc-ok-pkts" { return "ScOkPkts" }
    if yname == "sc-not-using-pkts" { return "ScNotUsingPkts" }
    if yname == "sc-unused-pkts" { return "ScUnusedPkts" }
    if yname == "sc-not-valid-pkts" { return "ScNotValidPkts" }
    if yname == "sc-invalid-pkts" { return "ScInvalidPkts" }
    if yname == "sc-overrun-pkts" { return "ScOverrunPkts" }
    if yname == "sc-bypass-pkts" { return "ScBypassPkts" }
    if yname == "sc-eapol-pkts" { return "ScEapolPkts" }
    if yname == "sc-dropped-pkts" { return "ScDroppedPkts" }
    if yname == "rx-sa-stat" { return "RxSaStat" }
    return ""
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetSegmentPath() string {
    return "macsec-rx-stats"
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-sa-stat" {
        for _, c := range macsecRxStats.RxSaStat {
            if macsecRxStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat{}
        macsecRxStats.RxSaStat = append(macsecRxStats.RxSaStat, child)
        return &macsecRxStats.RxSaStat[len(macsecRxStats.RxSaStat)-1]
    }
    return nil
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macsecRxStats.RxSaStat {
        children[macsecRxStats.RxSaStat[i].GetSegmentPath()] = &macsecRxStats.RxSaStat[i]
    }
    return children
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sc-decrypted-octets"] = macsecRxStats.ScDecryptedOctets
    leafs["sc-no-tag-pkts"] = macsecRxStats.ScNoTagPkts
    leafs["sc-untagged-pkts"] = macsecRxStats.ScUntaggedPkts
    leafs["sc-bad-tag-pkts"] = macsecRxStats.ScBadTagPkts
    leafs["sc-late-pkts"] = macsecRxStats.ScLatePkts
    leafs["sc-delayed-pkts"] = macsecRxStats.ScDelayedPkts
    leafs["sc-unchecked-pkts"] = macsecRxStats.ScUncheckedPkts
    leafs["sc-no-sci-pkts"] = macsecRxStats.ScNoSciPkts
    leafs["sc-unknown-sci-pkts"] = macsecRxStats.ScUnknownSciPkts
    leafs["sc-ok-pkts"] = macsecRxStats.ScOkPkts
    leafs["sc-not-using-pkts"] = macsecRxStats.ScNotUsingPkts
    leafs["sc-unused-pkts"] = macsecRxStats.ScUnusedPkts
    leafs["sc-not-valid-pkts"] = macsecRxStats.ScNotValidPkts
    leafs["sc-invalid-pkts"] = macsecRxStats.ScInvalidPkts
    leafs["sc-overrun-pkts"] = macsecRxStats.ScOverrunPkts
    leafs["sc-bypass-pkts"] = macsecRxStats.ScBypassPkts
    leafs["sc-eapol-pkts"] = macsecRxStats.ScEapolPkts
    leafs["sc-dropped-pkts"] = macsecRxStats.ScDroppedPkts
    return leafs
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetBundleName() string { return "cisco_ios_xr" }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetYangName() string { return "macsec-rx-stats" }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) SetParent(parent types.Entity) { macsecRxStats.parent = parent }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetParent() types.Entity { return macsecRxStats.parent }

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetParentYangName() string { return "xlfpga-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
// Rx SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current Rx AN. The type is interface{} with range: 0..18446744073709551615.
    An interface{}

    // Rx Ok Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaOkPkts interface{}

    // Rx Pkts not using SA for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotUsingPkts interface{}

    // Rx Pkts Unused Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaUnusedPkts interface{}

    // Rx Not Valid Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotValidPkts interface{}

    // Rx Invalid Pkts for current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaInvalidPkts interface{}
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetFilter() yfilter.YFilter { return rxSaStat.YFilter }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) SetFilter(yf yfilter.YFilter) { rxSaStat.YFilter = yf }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetGoName(yname string) string {
    if yname == "an" { return "An" }
    if yname == "sa-ok-pkts" { return "SaOkPkts" }
    if yname == "sa-not-using-pkts" { return "SaNotUsingPkts" }
    if yname == "sa-unused-pkts" { return "SaUnusedPkts" }
    if yname == "sa-not-valid-pkts" { return "SaNotValidPkts" }
    if yname == "sa-invalid-pkts" { return "SaInvalidPkts" }
    return ""
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetSegmentPath() string {
    return "rx-sa-stat"
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["an"] = rxSaStat.An
    leafs["sa-ok-pkts"] = rxSaStat.SaOkPkts
    leafs["sa-not-using-pkts"] = rxSaStat.SaNotUsingPkts
    leafs["sa-unused-pkts"] = rxSaStat.SaUnusedPkts
    leafs["sa-not-valid-pkts"] = rxSaStat.SaNotValidPkts
    leafs["sa-invalid-pkts"] = rxSaStat.SaInvalidPkts
    return leafs
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetBundleName() string { return "cisco_ios_xr" }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetYangName() string { return "rx-sa-stat" }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) SetParent(parent types.Entity) { rxSaStat.parent = parent }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetParent() types.Entity { return rxSaStat.parent }

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetParentYangName() string { return "macsec-rx-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats
// ES200 Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats

    // Tx SC Macsec Stats.
    TxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats

    // Rx SC Macsec Stats.
    RxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats

    // Port level TX Stats.
    TxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats

    // Port level RX Stats.
    RxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetFilter() yfilter.YFilter { return es200Stats.YFilter }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) SetFilter(yf yfilter.YFilter) { es200Stats.YFilter = yf }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetGoName(yname string) string {
    if yname == "tx-sa-stats" { return "TxSaStats" }
    if yname == "rx-sa-stats" { return "RxSaStats" }
    if yname == "tx-sc-macsec-stats" { return "TxScMacsecStats" }
    if yname == "rx-sc-macsec-stats" { return "RxScMacsecStats" }
    if yname == "tx-interface-macsec-stats" { return "TxInterfaceMacsecStats" }
    if yname == "rx-interface-macsec-stats" { return "RxInterfaceMacsecStats" }
    if yname == "tx-port-stats" { return "TxPortStats" }
    if yname == "rx-port-stats" { return "RxPortStats" }
    return ""
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetSegmentPath() string {
    return "es200-stats"
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-sa-stats" {
        return &es200Stats.TxSaStats
    }
    if childYangName == "rx-sa-stats" {
        return &es200Stats.RxSaStats
    }
    if childYangName == "tx-sc-macsec-stats" {
        return &es200Stats.TxScMacsecStats
    }
    if childYangName == "rx-sc-macsec-stats" {
        return &es200Stats.RxScMacsecStats
    }
    if childYangName == "tx-interface-macsec-stats" {
        return &es200Stats.TxInterfaceMacsecStats
    }
    if childYangName == "rx-interface-macsec-stats" {
        return &es200Stats.RxInterfaceMacsecStats
    }
    if childYangName == "tx-port-stats" {
        return &es200Stats.TxPortStats
    }
    if childYangName == "rx-port-stats" {
        return &es200Stats.RxPortStats
    }
    return nil
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-sa-stats"] = &es200Stats.TxSaStats
    children["rx-sa-stats"] = &es200Stats.RxSaStats
    children["tx-sc-macsec-stats"] = &es200Stats.TxScMacsecStats
    children["rx-sc-macsec-stats"] = &es200Stats.RxScMacsecStats
    children["tx-interface-macsec-stats"] = &es200Stats.TxInterfaceMacsecStats
    children["rx-interface-macsec-stats"] = &es200Stats.RxInterfaceMacsecStats
    children["tx-port-stats"] = &es200Stats.TxPortStats
    children["rx-port-stats"] = &es200Stats.RxPortStats
    return children
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetBundleName() string { return "cisco_ios_xr" }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetYangName() string { return "es200-stats" }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) SetParent(parent types.Entity) { es200Stats.parent = parent }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetParent() types.Entity { return es200Stats.parent }

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetParentYangName() string { return "ext" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // packets exceeding egress MTU. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsTooLong interface{}

    // packets encrypted/protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncryptedProtected interface{}

    // octets1 encrypted/protected ?. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncryptedProtected1 interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetFilter() yfilter.YFilter { return txSaStats.YFilter }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) SetFilter(yf yfilter.YFilter) { txSaStats.YFilter = yf }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetGoName(yname string) string {
    if yname == "out-pkts-too-long" { return "OutPktsTooLong" }
    if yname == "out-pkts-encrypted-protected" { return "OutPktsEncryptedProtected" }
    if yname == "out-octets-encrypted-protected1" { return "OutOctetsEncryptedProtected1" }
    return ""
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetSegmentPath() string {
    return "tx-sa-stats"
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkts-too-long"] = txSaStats.OutPktsTooLong
    leafs["out-pkts-encrypted-protected"] = txSaStats.OutPktsEncryptedProtected
    leafs["out-octets-encrypted-protected1"] = txSaStats.OutOctetsEncryptedProtected1
    return leafs
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetYangName() string { return "tx-sa-stats" }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) SetParent(parent types.Entity) { txSaStats.parent = parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetParent() types.Entity { return txSaStats.parent }

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // frame not valid & validateFrames disabled. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsUnchecked interface{}

    // PN of packet outside replay window & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsDelayed interface{}

    // PN of packet outside replay window & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // packets with no error. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsOk interface{}

    // packet not valid & validateFrames !strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsInvalid interface{}

    // packet not valid & validateFrames strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsNotValid interface{}

    // packet assigned to SA not in use & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // packet assigned to SA not in use & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // octets1 decrypted/validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecryptedValidated1 interface{}

    // octets validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetFilter() yfilter.YFilter { return rxSaStats.YFilter }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) SetFilter(yf yfilter.YFilter) { rxSaStats.YFilter = yf }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetGoName(yname string) string {
    if yname == "in-pkts-unchecked" { return "InPktsUnchecked" }
    if yname == "in-pkts-delayed" { return "InPktsDelayed" }
    if yname == "in-pkts-late" { return "InPktsLate" }
    if yname == "in-pkts-ok" { return "InPktsOk" }
    if yname == "in-pkts-invalid" { return "InPktsInvalid" }
    if yname == "in-pkts-not-valid" { return "InPktsNotValid" }
    if yname == "in-pkts-not-using-sa" { return "InPktsNotUsingSa" }
    if yname == "in-pkts-unused-sa" { return "InPktsUnusedSa" }
    if yname == "in-octets-decrypted-validated1" { return "InOctetsDecryptedValidated1" }
    if yname == "in-octets-validated" { return "InOctetsValidated" }
    return ""
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetSegmentPath() string {
    return "rx-sa-stats"
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-unchecked"] = rxSaStats.InPktsUnchecked
    leafs["in-pkts-delayed"] = rxSaStats.InPktsDelayed
    leafs["in-pkts-late"] = rxSaStats.InPktsLate
    leafs["in-pkts-ok"] = rxSaStats.InPktsOk
    leafs["in-pkts-invalid"] = rxSaStats.InPktsInvalid
    leafs["in-pkts-not-valid"] = rxSaStats.InPktsNotValid
    leafs["in-pkts-not-using-sa"] = rxSaStats.InPktsNotUsingSa
    leafs["in-pkts-unused-sa"] = rxSaStats.InPktsUnusedSa
    leafs["in-octets-decrypted-validated1"] = rxSaStats.InOctetsDecryptedValidated1
    leafs["in-octets-validated"] = rxSaStats.InOctetsValidated
    return leafs
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetYangName() string { return "rx-sa-stats" }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) SetParent(parent types.Entity) { rxSaStats.parent = parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetParent() types.Entity { return rxSaStats.parent }

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats
// Tx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsSaNotInUse interface{}
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetFilter() yfilter.YFilter { return txScMacsecStats.YFilter }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) SetFilter(yf yfilter.YFilter) { txScMacsecStats.YFilter = yf }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetGoName(yname string) string {
    if yname == "out-pkts-sa-not-in-use" { return "OutPktsSaNotInUse" }
    return ""
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetSegmentPath() string {
    return "tx-sc-macsec-stats"
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkts-sa-not-in-use"] = txScMacsecStats.OutPktsSaNotInUse
    return leafs
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetYangName() string { return "tx-sc-macsec-stats" }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) SetParent(parent types.Entity) { txScMacsecStats.parent = parent }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetParent() types.Entity { return txScMacsecStats.parent }

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats
// Rx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsSaNotInUse interface{}
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetFilter() yfilter.YFilter { return rxScMacsecStats.YFilter }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) SetFilter(yf yfilter.YFilter) { rxScMacsecStats.YFilter = yf }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetGoName(yname string) string {
    if yname == "in-pkts-sa-not-in-use" { return "InPktsSaNotInUse" }
    return ""
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetSegmentPath() string {
    return "rx-sc-macsec-stats"
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-sa-not-in-use"] = rxScMacsecStats.InPktsSaNotInUse
    return leafs
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetYangName() string { return "rx-sc-macsec-stats" }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) SetParent(parent types.Entity) { rxScMacsecStats.parent = parent }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetParent() types.Entity { return rxScMacsecStats.parent }

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // egress packet that is classified as control packet. The type is interface{}
    // with range: 0..18446744073709551615.
    OutPktCtrl interface{}

    // egress packet to go out untagged when protectFrames not set. The type is
    // interface{} with range: 0..18446744073709551615.
    OutPktsUntagged interface{}

    // Octets tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsUnctrl interface{}

    // Octets tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCtrl interface{}

    // Octets tx on common port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCommon interface{}

    // Unicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsUnctrl interface{}

    // Unicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsCtrl interface{}

    // Multicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsUnctrl interface{}

    // Multicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsCtrl interface{}

    // Broadcast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsUnctrl interface{}

    // Broadcast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsClass interface{}

    // Packets dropped due to overflow in  processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsData interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return txInterfaceMacsecStats.YFilter }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { txInterfaceMacsecStats.YFilter = yf }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "transform-error-pkts" { return "TransformErrorPkts" }
    if yname == "out-pkt-ctrl" { return "OutPktCtrl" }
    if yname == "out-pkts-untagged" { return "OutPktsUntagged" }
    if yname == "out-octets-unctrl" { return "OutOctetsUnctrl" }
    if yname == "out-octets-ctrl" { return "OutOctetsCtrl" }
    if yname == "out-octets-common" { return "OutOctetsCommon" }
    if yname == "out-ucast-pkts-unctrl" { return "OutUcastPktsUnctrl" }
    if yname == "out-ucast-pkts-ctrl" { return "OutUcastPktsCtrl" }
    if yname == "out-mcast-pkts-unctrl" { return "OutMcastPktsUnctrl" }
    if yname == "out-mcast-pkts-ctrl" { return "OutMcastPktsCtrl" }
    if yname == "out-bcast-pkts-unctrl" { return "OutBcastPktsUnctrl" }
    if yname == "out-bcast-pkts-ctrl" { return "OutBcastPktsCtrl" }
    if yname == "out-rx-drop-pkts-unctrl" { return "OutRxDropPktsUnctrl" }
    if yname == "out-rx-drop-pkts-ctrl" { return "OutRxDropPktsCtrl" }
    if yname == "out-rx-err-pkts-unctrl" { return "OutRxErrPktsUnctrl" }
    if yname == "out-rx-err-pkts-ctrl" { return "OutRxErrPktsCtrl" }
    if yname == "out-drop-pkts-class" { return "OutDropPktsClass" }
    if yname == "out-drop-pkts-data" { return "OutDropPktsData" }
    return ""
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetSegmentPath() string {
    return "tx-interface-macsec-stats"
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transform-error-pkts"] = txInterfaceMacsecStats.TransformErrorPkts
    leafs["out-pkt-ctrl"] = txInterfaceMacsecStats.OutPktCtrl
    leafs["out-pkts-untagged"] = txInterfaceMacsecStats.OutPktsUntagged
    leafs["out-octets-unctrl"] = txInterfaceMacsecStats.OutOctetsUnctrl
    leafs["out-octets-ctrl"] = txInterfaceMacsecStats.OutOctetsCtrl
    leafs["out-octets-common"] = txInterfaceMacsecStats.OutOctetsCommon
    leafs["out-ucast-pkts-unctrl"] = txInterfaceMacsecStats.OutUcastPktsUnctrl
    leafs["out-ucast-pkts-ctrl"] = txInterfaceMacsecStats.OutUcastPktsCtrl
    leafs["out-mcast-pkts-unctrl"] = txInterfaceMacsecStats.OutMcastPktsUnctrl
    leafs["out-mcast-pkts-ctrl"] = txInterfaceMacsecStats.OutMcastPktsCtrl
    leafs["out-bcast-pkts-unctrl"] = txInterfaceMacsecStats.OutBcastPktsUnctrl
    leafs["out-bcast-pkts-ctrl"] = txInterfaceMacsecStats.OutBcastPktsCtrl
    leafs["out-rx-drop-pkts-unctrl"] = txInterfaceMacsecStats.OutRxDropPktsUnctrl
    leafs["out-rx-drop-pkts-ctrl"] = txInterfaceMacsecStats.OutRxDropPktsCtrl
    leafs["out-rx-err-pkts-unctrl"] = txInterfaceMacsecStats.OutRxErrPktsUnctrl
    leafs["out-rx-err-pkts-ctrl"] = txInterfaceMacsecStats.OutRxErrPktsCtrl
    leafs["out-drop-pkts-class"] = txInterfaceMacsecStats.OutDropPktsClass
    leafs["out-drop-pkts-data"] = txInterfaceMacsecStats.OutDropPktsData
    return leafs
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetYangName() string { return "tx-interface-macsec-stats" }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) SetParent(parent types.Entity) { txInterfaceMacsecStats.parent = parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetParent() types.Entity { return txInterfaceMacsecStats.parent }

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // ingress packet that is classified as control packet. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktCtrl interface{}

    // ingress packet untagged & validateFrames is strict. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktNoTag interface{}

    // ingress packet untagged & validateFrames is  !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUntagged interface{}

    // ingress frames received with an invalid MACSec tag or ICV                  
    // added with next one gives InPktsSCIMiss. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktBadTag interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is !strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnknownSci interface{}

    // ingress packets that are control or KaY packets. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktsTaggedCtrl interface{}

    // Octets rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsUnctrl interface{}

    // Octets rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsCtrl interface{}

    // Unicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsUnctrl interface{}

    // Unicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsCtrl interface{}

    // Multicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsUnctrl interface{}

    // Multicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsCtrl interface{}

    // Broadcast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsUnctrl interface{}

    // Broadcast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsClass interface{}

    // Packets dropped due to overflow in processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsData interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetFilter() yfilter.YFilter { return rxInterfaceMacsecStats.YFilter }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) SetFilter(yf yfilter.YFilter) { rxInterfaceMacsecStats.YFilter = yf }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetGoName(yname string) string {
    if yname == "transform-error-pkts" { return "TransformErrorPkts" }
    if yname == "in-pkt-ctrl" { return "InPktCtrl" }
    if yname == "in-pkt-no-tag" { return "InPktNoTag" }
    if yname == "in-pkts-untagged" { return "InPktsUntagged" }
    if yname == "in-pkt-bad-tag" { return "InPktBadTag" }
    if yname == "in-pkt-no-sci" { return "InPktNoSci" }
    if yname == "in-pkts-unknown-sci" { return "InPktsUnknownSci" }
    if yname == "in-pkts-tagged-ctrl" { return "InPktsTaggedCtrl" }
    if yname == "in-octets-unctrl" { return "InOctetsUnctrl" }
    if yname == "in-octets-ctrl" { return "InOctetsCtrl" }
    if yname == "in-ucast-pkts-unctrl" { return "InUcastPktsUnctrl" }
    if yname == "in-ucast-pkts-ctrl" { return "InUcastPktsCtrl" }
    if yname == "in-mcast-pkts-unctrl" { return "InMcastPktsUnctrl" }
    if yname == "in-mcast-pkts-ctrl" { return "InMcastPktsCtrl" }
    if yname == "in-bcast-pkts-unctrl" { return "InBcastPktsUnctrl" }
    if yname == "in-bcast-pkts-ctrl" { return "InBcastPktsCtrl" }
    if yname == "in-rx-drop-pkts-unctrl" { return "InRxDropPktsUnctrl" }
    if yname == "in-rx-drop-pkts-ctrl" { return "InRxDropPktsCtrl" }
    if yname == "in-rx-error-pkts-unctrl" { return "InRxErrorPktsUnctrl" }
    if yname == "in-rx-error-pkts-ctrl" { return "InRxErrorPktsCtrl" }
    if yname == "in-drop-pkts-class" { return "InDropPktsClass" }
    if yname == "in-drop-pkts-data" { return "InDropPktsData" }
    return ""
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetSegmentPath() string {
    return "rx-interface-macsec-stats"
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transform-error-pkts"] = rxInterfaceMacsecStats.TransformErrorPkts
    leafs["in-pkt-ctrl"] = rxInterfaceMacsecStats.InPktCtrl
    leafs["in-pkt-no-tag"] = rxInterfaceMacsecStats.InPktNoTag
    leafs["in-pkts-untagged"] = rxInterfaceMacsecStats.InPktsUntagged
    leafs["in-pkt-bad-tag"] = rxInterfaceMacsecStats.InPktBadTag
    leafs["in-pkt-no-sci"] = rxInterfaceMacsecStats.InPktNoSci
    leafs["in-pkts-unknown-sci"] = rxInterfaceMacsecStats.InPktsUnknownSci
    leafs["in-pkts-tagged-ctrl"] = rxInterfaceMacsecStats.InPktsTaggedCtrl
    leafs["in-octets-unctrl"] = rxInterfaceMacsecStats.InOctetsUnctrl
    leafs["in-octets-ctrl"] = rxInterfaceMacsecStats.InOctetsCtrl
    leafs["in-ucast-pkts-unctrl"] = rxInterfaceMacsecStats.InUcastPktsUnctrl
    leafs["in-ucast-pkts-ctrl"] = rxInterfaceMacsecStats.InUcastPktsCtrl
    leafs["in-mcast-pkts-unctrl"] = rxInterfaceMacsecStats.InMcastPktsUnctrl
    leafs["in-mcast-pkts-ctrl"] = rxInterfaceMacsecStats.InMcastPktsCtrl
    leafs["in-bcast-pkts-unctrl"] = rxInterfaceMacsecStats.InBcastPktsUnctrl
    leafs["in-bcast-pkts-ctrl"] = rxInterfaceMacsecStats.InBcastPktsCtrl
    leafs["in-rx-drop-pkts-unctrl"] = rxInterfaceMacsecStats.InRxDropPktsUnctrl
    leafs["in-rx-drop-pkts-ctrl"] = rxInterfaceMacsecStats.InRxDropPktsCtrl
    leafs["in-rx-error-pkts-unctrl"] = rxInterfaceMacsecStats.InRxErrorPktsUnctrl
    leafs["in-rx-error-pkts-ctrl"] = rxInterfaceMacsecStats.InRxErrorPktsCtrl
    leafs["in-drop-pkts-class"] = rxInterfaceMacsecStats.InDropPktsClass
    leafs["in-drop-pkts-data"] = rxInterfaceMacsecStats.InDropPktsData
    return leafs
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetYangName() string { return "rx-interface-macsec-stats" }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) SetParent(parent types.Entity) { rxInterfaceMacsecStats.parent = parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetParent() types.Entity { return rxInterfaceMacsecStats.parent }

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats
// Port level TX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetFilter() yfilter.YFilter { return txPortStats.YFilter }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) SetFilter(yf yfilter.YFilter) { txPortStats.YFilter = yf }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetGoName(yname string) string {
    if yname == "multi-flow-match" { return "MultiFlowMatch" }
    if yname == "parser-dropped" { return "ParserDropped" }
    if yname == "flow-miss" { return "FlowMiss" }
    if yname == "pkts-ctrl" { return "PktsCtrl" }
    if yname == "pkts-data" { return "PktsData" }
    if yname == "pkts-dropped" { return "PktsDropped" }
    if yname == "pkts-err-in" { return "PktsErrIn" }
    return ""
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetSegmentPath() string {
    return "tx-port-stats"
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multi-flow-match"] = txPortStats.MultiFlowMatch
    leafs["parser-dropped"] = txPortStats.ParserDropped
    leafs["flow-miss"] = txPortStats.FlowMiss
    leafs["pkts-ctrl"] = txPortStats.PktsCtrl
    leafs["pkts-data"] = txPortStats.PktsData
    leafs["pkts-dropped"] = txPortStats.PktsDropped
    leafs["pkts-err-in"] = txPortStats.PktsErrIn
    return leafs
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetBundleName() string { return "cisco_ios_xr" }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetYangName() string { return "tx-port-stats" }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) SetParent(parent types.Entity) { txPortStats.parent = parent }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetParent() types.Entity { return txPortStats.parent }

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetParentYangName() string { return "es200-stats" }

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats
// Port level RX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetFilter() yfilter.YFilter { return rxPortStats.YFilter }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) SetFilter(yf yfilter.YFilter) { rxPortStats.YFilter = yf }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetGoName(yname string) string {
    if yname == "multi-flow-match" { return "MultiFlowMatch" }
    if yname == "parser-dropped" { return "ParserDropped" }
    if yname == "flow-miss" { return "FlowMiss" }
    if yname == "pkts-ctrl" { return "PktsCtrl" }
    if yname == "pkts-data" { return "PktsData" }
    if yname == "pkts-dropped" { return "PktsDropped" }
    if yname == "pkts-err-in" { return "PktsErrIn" }
    return ""
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetSegmentPath() string {
    return "rx-port-stats"
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multi-flow-match"] = rxPortStats.MultiFlowMatch
    leafs["parser-dropped"] = rxPortStats.ParserDropped
    leafs["flow-miss"] = rxPortStats.FlowMiss
    leafs["pkts-ctrl"] = rxPortStats.PktsCtrl
    leafs["pkts-data"] = rxPortStats.PktsData
    leafs["pkts-dropped"] = rxPortStats.PktsDropped
    leafs["pkts-err-in"] = rxPortStats.PktsErrIn
    return leafs
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetYangName() string { return "rx-port-stats" }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) SetParent(parent types.Entity) { rxPortStats.parent = parent }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetParent() types.Entity { return rxPortStats.parent }

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetParentYangName() string { return "es200-stats" }

