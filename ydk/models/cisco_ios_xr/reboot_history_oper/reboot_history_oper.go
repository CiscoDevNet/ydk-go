// This module contains a collection of YANG definitions
// for Cisco IOS-XR reboot-history package operational data.
// 
// This module contains definitions
// for the following management objects:
//   reboot-history: Reboot History information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package reboot_history_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package reboot_history_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-reboot-history-oper reboot-history}", reflect.TypeOf(RebootHistory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-reboot-history-oper:reboot-history", reflect.TypeOf(RebootHistory{}))
}

// RebootHistory
// Reboot History information
type RebootHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node ID. The type is slice of RebootHistory_Node.
    Node []RebootHistory_Node
}

func (rebootHistory *RebootHistory) GetFilter() yfilter.YFilter { return rebootHistory.YFilter }

func (rebootHistory *RebootHistory) SetFilter(yf yfilter.YFilter) { rebootHistory.YFilter = yf }

func (rebootHistory *RebootHistory) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (rebootHistory *RebootHistory) GetSegmentPath() string {
    return "Cisco-IOS-XR-reboot-history-oper:reboot-history"
}

func (rebootHistory *RebootHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range rebootHistory.Node {
            if rebootHistory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RebootHistory_Node{}
        rebootHistory.Node = append(rebootHistory.Node, child)
        return &rebootHistory.Node[len(rebootHistory.Node)-1]
    }
    return nil
}

func (rebootHistory *RebootHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rebootHistory.Node {
        children[rebootHistory.Node[i].GetSegmentPath()] = &rebootHistory.Node[i]
    }
    return children
}

func (rebootHistory *RebootHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rebootHistory *RebootHistory) GetBundleName() string { return "cisco_ios_xr" }

func (rebootHistory *RebootHistory) GetYangName() string { return "reboot-history" }

func (rebootHistory *RebootHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rebootHistory *RebootHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rebootHistory *RebootHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rebootHistory *RebootHistory) SetParent(parent types.Entity) { rebootHistory.parent = parent }

func (rebootHistory *RebootHistory) GetParent() types.Entity { return rebootHistory.parent }

func (rebootHistory *RebootHistory) GetParentYangName() string { return "Cisco-IOS-XR-reboot-history-oper" }

// RebootHistory_Node
// Node ID
type RebootHistory_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Last Reboots. The type is slice of RebootHistory_Node_RebootHistory.
    RebootHistory []RebootHistory_Node_RebootHistory
}

func (node *RebootHistory_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *RebootHistory_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *RebootHistory_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "reboot-history" { return "RebootHistory" }
    return ""
}

func (node *RebootHistory_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *RebootHistory_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reboot-history" {
        for _, c := range node.RebootHistory {
            if node.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RebootHistory_Node_RebootHistory{}
        node.RebootHistory = append(node.RebootHistory, child)
        return &node.RebootHistory[len(node.RebootHistory)-1]
    }
    return nil
}

func (node *RebootHistory_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range node.RebootHistory {
        children[node.RebootHistory[i].GetSegmentPath()] = &node.RebootHistory[i]
    }
    return children
}

func (node *RebootHistory_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *RebootHistory_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *RebootHistory_Node) GetYangName() string { return "node" }

func (node *RebootHistory_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *RebootHistory_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *RebootHistory_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *RebootHistory_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *RebootHistory_Node) GetParent() types.Entity { return node.parent }

func (node *RebootHistory_Node) GetParentYangName() string { return "reboot-history" }

// RebootHistory_Node_RebootHistory
// Last Reboots
type RebootHistory_Node_RebootHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number count. The type is interface{} with range: 0..4294967295.
    No interface{}

    // Time of reboot. The type is string.
    Time interface{}

    // Cause code for reboot. The type is interface{} with range: 0..4294967295.
    CauseCode interface{}

    // Reason for reboot. The type is string.
    Reason interface{}
}

func (rebootHistory *RebootHistory_Node_RebootHistory) GetFilter() yfilter.YFilter { return rebootHistory.YFilter }

func (rebootHistory *RebootHistory_Node_RebootHistory) SetFilter(yf yfilter.YFilter) { rebootHistory.YFilter = yf }

func (rebootHistory *RebootHistory_Node_RebootHistory) GetGoName(yname string) string {
    if yname == "no" { return "No" }
    if yname == "time" { return "Time" }
    if yname == "cause-code" { return "CauseCode" }
    if yname == "reason" { return "Reason" }
    return ""
}

func (rebootHistory *RebootHistory_Node_RebootHistory) GetSegmentPath() string {
    return "reboot-history"
}

func (rebootHistory *RebootHistory_Node_RebootHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rebootHistory *RebootHistory_Node_RebootHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rebootHistory *RebootHistory_Node_RebootHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["no"] = rebootHistory.No
    leafs["time"] = rebootHistory.Time
    leafs["cause-code"] = rebootHistory.CauseCode
    leafs["reason"] = rebootHistory.Reason
    return leafs
}

func (rebootHistory *RebootHistory_Node_RebootHistory) GetBundleName() string { return "cisco_ios_xr" }

func (rebootHistory *RebootHistory_Node_RebootHistory) GetYangName() string { return "reboot-history" }

func (rebootHistory *RebootHistory_Node_RebootHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rebootHistory *RebootHistory_Node_RebootHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rebootHistory *RebootHistory_Node_RebootHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rebootHistory *RebootHistory_Node_RebootHistory) SetParent(parent types.Entity) { rebootHistory.parent = parent }

func (rebootHistory *RebootHistory_Node_RebootHistory) GetParent() types.Entity { return rebootHistory.parent }

func (rebootHistory *RebootHistory_Node_RebootHistory) GetParentYangName() string { return "node" }

