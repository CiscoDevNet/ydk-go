// This module contains a collection of YANG definitions
// for Cisco IOS-XR reboot-history package operational data.
// 
// This module contains definitions
// for the following management objects:
//   reboot-history: Reboot History information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ID. The type is slice of RebootHistory_Node.
    Node []*RebootHistory_Node
}

func (rebootHistory *RebootHistory) GetEntityData() *types.CommonEntityData {
    rebootHistory.EntityData.YFilter = rebootHistory.YFilter
    rebootHistory.EntityData.YangName = "reboot-history"
    rebootHistory.EntityData.BundleName = "cisco_ios_xr"
    rebootHistory.EntityData.ParentYangName = "Cisco-IOS-XR-reboot-history-oper"
    rebootHistory.EntityData.SegmentPath = "Cisco-IOS-XR-reboot-history-oper:reboot-history"
    rebootHistory.EntityData.AbsolutePath = rebootHistory.EntityData.SegmentPath
    rebootHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHistory.EntityData.Children = types.NewOrderedMap()
    rebootHistory.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range rebootHistory.Node {
        rebootHistory.EntityData.Children.Append(types.GetSegmentPath(rebootHistory.Node[i]), types.YChild{"Node", rebootHistory.Node[i]})
    }
    rebootHistory.EntityData.Leafs = types.NewOrderedMap()

    rebootHistory.EntityData.YListKeys = []string {}

    return &(rebootHistory.EntityData)
}

// RebootHistory_Node
// Node ID
type RebootHistory_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Last Reboots. The type is slice of RebootHistory_Node_RebootHistory.
    RebootHistory []*RebootHistory_Node_RebootHistory
}

func (node *RebootHistory_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "reboot-history"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-reboot-history-oper:reboot-history/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("reboot-history", types.YChild{"RebootHistory", nil})
    for i := range node.RebootHistory {
        types.SetYListKey(node.RebootHistory[i], i)
        node.EntityData.Children.Append(types.GetSegmentPath(node.RebootHistory[i]), types.YChild{"RebootHistory", node.RebootHistory[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// RebootHistory_Node_RebootHistory
// Last Reboots
type RebootHistory_Node_RebootHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Number count. The type is interface{} with range: 0..4294967295.
    No interface{}

    // Time of reboot. The type is string.
    Time interface{}

    // Cause code for reboot. The type is interface{} with range: 0..4294967295.
    CauseCode interface{}

    // Reason for reboot. The type is string.
    Reason interface{}
}

func (rebootHistory *RebootHistory_Node_RebootHistory) GetEntityData() *types.CommonEntityData {
    rebootHistory.EntityData.YFilter = rebootHistory.YFilter
    rebootHistory.EntityData.YangName = "reboot-history"
    rebootHistory.EntityData.BundleName = "cisco_ios_xr"
    rebootHistory.EntityData.ParentYangName = "node"
    rebootHistory.EntityData.SegmentPath = "reboot-history" + types.AddNoKeyToken(rebootHistory)
    rebootHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-reboot-history-oper:reboot-history/node/" + rebootHistory.EntityData.SegmentPath
    rebootHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHistory.EntityData.Children = types.NewOrderedMap()
    rebootHistory.EntityData.Leafs = types.NewOrderedMap()
    rebootHistory.EntityData.Leafs.Append("no", types.YLeaf{"No", rebootHistory.No})
    rebootHistory.EntityData.Leafs.Append("time", types.YLeaf{"Time", rebootHistory.Time})
    rebootHistory.EntityData.Leafs.Append("cause-code", types.YLeaf{"CauseCode", rebootHistory.CauseCode})
    rebootHistory.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", rebootHistory.Reason})

    rebootHistory.EntityData.YListKeys = []string {}

    return &(rebootHistory.EntityData)
}

