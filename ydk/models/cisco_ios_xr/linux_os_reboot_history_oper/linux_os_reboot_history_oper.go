// This module contains a collection of YANG definitions
// for Cisco IOS-XR linux-os-reboot-history package operational data.
// 
// This module contains definitions
// for the following management objects:
//   reboot-history: Reboot History information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package linux_os_reboot_history_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package linux_os_reboot_history_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-linux-os-reboot-history-oper reboot-history}", reflect.TypeOf(RebootHistory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-linux-os-reboot-history-oper:reboot-history", reflect.TypeOf(RebootHistory{}))
}

// RebootHistory
// Reboot History information
type RebootHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ID. The type is slice of RebootHistory_Node.
    Node []RebootHistory_Node
}

func (rebootHistory *RebootHistory) GetEntityData() *types.CommonEntityData {
    rebootHistory.EntityData.YFilter = rebootHistory.YFilter
    rebootHistory.EntityData.YangName = "reboot-history"
    rebootHistory.EntityData.BundleName = "cisco_ios_xr"
    rebootHistory.EntityData.ParentYangName = "Cisco-IOS-XR-linux-os-reboot-history-oper"
    rebootHistory.EntityData.SegmentPath = "Cisco-IOS-XR-linux-os-reboot-history-oper:reboot-history"
    rebootHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHistory.EntityData.Children = make(map[string]types.YChild)
    rebootHistory.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range rebootHistory.Node {
        rebootHistory.EntityData.Children[types.GetSegmentPath(&rebootHistory.Node[i])] = types.YChild{"Node", &rebootHistory.Node[i]}
    }
    rebootHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rebootHistory.EntityData)
}

// RebootHistory_Node
// Node ID
type RebootHistory_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Last Reboots. The type is slice of RebootHistory_Node_RebootHistory.
    RebootHistory []RebootHistory_Node_RebootHistory_
}

func (node *RebootHistory_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "reboot-history"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["reboot-history"] = types.YChild{"RebootHistory", nil}
    for i := range node.RebootHistory {
        node.EntityData.Children[types.GetSegmentPath(&node.RebootHistory[i])] = types.YChild{"RebootHistory", &node.RebootHistory[i]}
    }
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// RebootHistory_Node_RebootHistory_
// Last Reboots
type RebootHistory_Node_RebootHistory_ struct {
    EntityData types.CommonEntityData
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

func (rebootHistory_ *RebootHistory_Node_RebootHistory_) GetEntityData() *types.CommonEntityData {
    rebootHistory_.EntityData.YFilter = rebootHistory_.YFilter
    rebootHistory_.EntityData.YangName = "reboot-history"
    rebootHistory_.EntityData.BundleName = "cisco_ios_xr"
    rebootHistory_.EntityData.ParentYangName = "node"
    rebootHistory_.EntityData.SegmentPath = "reboot-history"
    rebootHistory_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHistory_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHistory_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHistory_.EntityData.Children = make(map[string]types.YChild)
    rebootHistory_.EntityData.Leafs = make(map[string]types.YLeaf)
    rebootHistory_.EntityData.Leafs["no"] = types.YLeaf{"No", rebootHistory_.No}
    rebootHistory_.EntityData.Leafs["time"] = types.YLeaf{"Time", rebootHistory_.Time}
    rebootHistory_.EntityData.Leafs["cause-code"] = types.YLeaf{"CauseCode", rebootHistory_.CauseCode}
    rebootHistory_.EntityData.Leafs["reason"] = types.YLeaf{"Reason", rebootHistory_.Reason}
    return &(rebootHistory_.EntityData)
}

