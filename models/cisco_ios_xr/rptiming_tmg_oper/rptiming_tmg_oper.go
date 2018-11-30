// This module contains a collection of YANG definitions
// for Cisco IOS-XR rptiming-tmg package operational data.
// 
// This module contains definitions
// for the following management objects:
//   timing-card: Timing PLL status and configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package rptiming_tmg_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rptiming_tmg_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-rptiming-tmg-oper timing-card}", reflect.TypeOf(TimingCard{}))
    ydk.RegisterEntity("Cisco-IOS-XR-rptiming-tmg-oper:timing-card", reflect.TypeOf(TimingCard{}))
}

// TimingCard
// Timing PLL status and configuration
type TimingCard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes applicable to timing.
    Nodes TimingCard_Nodes
}

func (timingCard *TimingCard) GetEntityData() *types.CommonEntityData {
    timingCard.EntityData.YFilter = timingCard.YFilter
    timingCard.EntityData.YangName = "timing-card"
    timingCard.EntityData.BundleName = "cisco_ios_xr"
    timingCard.EntityData.ParentYangName = "Cisco-IOS-XR-rptiming-tmg-oper"
    timingCard.EntityData.SegmentPath = "Cisco-IOS-XR-rptiming-tmg-oper:timing-card"
    timingCard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timingCard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timingCard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timingCard.EntityData.Children = types.NewOrderedMap()
    timingCard.EntityData.Children.Append("nodes", types.YChild{"Nodes", &timingCard.Nodes})
    timingCard.EntityData.Leafs = types.NewOrderedMap()

    timingCard.EntityData.YListKeys = []string {}

    return &(timingCard.EntityData)
}

// TimingCard_Nodes
// List of nodes applicable to timing
type TimingCard_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timing operational data for a single node. The type is slice of
    // TimingCard_Nodes_Node.
    Node []*TimingCard_Nodes_Node
}

func (nodes *TimingCard_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "timing-card"
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

// TimingCard_Nodes_Node
// Timing operational data for a single node
type TimingCard_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node Name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Display the timing card input clock status information.
    InputClock TimingCard_Nodes_Node_InputClock

    // Display the timing card PLL status information.
    Pll TimingCard_Nodes_Node_Pll
}

func (node *TimingCard_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("input-clock", types.YChild{"InputClock", &node.InputClock})
    node.EntityData.Children.Append("pll", types.YChild{"Pll", &node.Pll})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// TimingCard_Nodes_Node_InputClock
// Display the timing card input clock status
// information
type TimingCard_Nodes_Node_InputClock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IC Valid 1. The type is bool.
    Ic1Valid interface{}

    // IC Valid 2. The type is bool.
    Ic2Valid interface{}

    // IC Valid 3. The type is bool.
    Ic3Valid interface{}

    // IC Valid 4. The type is bool.
    Ic4Valid interface{}

    // IC Valid 5. The type is bool.
    Ic5Valid interface{}

    // IC Valid 6. The type is bool.
    Ic6Valid interface{}

    // IC Valid 7. The type is bool.
    Ic7Valid interface{}

    // IC Valid 8. The type is bool.
    Ic8Valid interface{}

    // IC Valid 9. The type is bool.
    Ic9Valid interface{}

    // IC Valid 10. The type is bool.
    Ic10Valid interface{}

    // IC Valid 11. The type is bool.
    Ic11Valid interface{}

    // IC Slot 1. The type is string with length: 0..50.
    Ic1Slot interface{}

    // IC Slot 2. The type is string with length: 0..50.
    Ic2Slot interface{}

    // IC Slot 3. The type is string with length: 0..50.
    Ic3Slot interface{}

    // IC Slot 4. The type is string with length: 0..50.
    Ic4Slot interface{}

    // IC Slot 5. The type is string with length: 0..50.
    Ic5Slot interface{}

    // IC Slot 6. The type is string with length: 0..50.
    Ic6Slot interface{}

    // IC Slot 7. The type is string with length: 0..50.
    Ic7Slot interface{}

    // IC Slot 8. The type is string with length: 0..50.
    Ic8Slot interface{}

    // IC Slot 9. The type is string with length: 0..50.
    Ic9Slot interface{}

    // IC Slot 10. The type is string with length: 0..50.
    Ic10Slot interface{}

    // IC Slot 11. The type is string with length: 0..50.
    Ic11Slot interface{}

    // IC1 Split XO Mode Status. The type is string with length: 0..50.
    Ic1SplitXom interface{}

    // IC2 Split XO Mode Status. The type is string with length: 0..50.
    Ic2SplitXom interface{}

    // IC3 Split XO Mode Status. The type is string with length: 0..50.
    Ic3SplitXom interface{}

    // IC4 Split XO Mode Status. The type is string with length: 0..50.
    Ic4SplitXom interface{}

    // IC5 Split XO Mode Status. The type is string with length: 0..50.
    Ic5SplitXom interface{}

    // IC6 Split XO Mode Status. The type is string with length: 0..50.
    Ic6SplitXom interface{}

    // IC7 Split XO Mode Status. The type is string with length: 0..50.
    Ic7SplitXom interface{}

    // IC8 Split XO Mode Status. The type is string with length: 0..50.
    Ic8SplitXom interface{}

    // IC9 Split XO Mode Status. The type is string with length: 0..50.
    Ic9SplitXom interface{}

    // IC10 Split XO Mode Status. The type is string with length: 0..50.
    Ic10SplitXom interface{}

    // IC11 Split XO Mode Status. The type is string with length: 0..50.
    Ic11SplitXom interface{}

    // IC1 ePPSM Status. The type is string with length: 0..50.
    Ic1Eppsm interface{}

    // IC2 ePPSM Status. The type is string with length: 0..50.
    Ic2Eppsm interface{}

    // IC3 ePPSM Status. The type is string with length: 0..50.
    Ic3Eppsm interface{}

    // IC4 ePPSM Status. The type is string with length: 0..50.
    Ic4Eppsm interface{}

    // IC5 ePPSM Status. The type is string with length: 0..50.
    Ic5Eppsm interface{}

    // IC6 ePPSM Status. The type is string with length: 0..50.
    Ic6Eppsm interface{}

    // IC7 ePPSM Status. The type is string with length: 0..50.
    Ic7Eppsm interface{}

    // IC8 ePPSM Status. The type is string with length: 0..50.
    Ic8Eppsm interface{}

    // IC9 ePPSM Status. The type is string with length: 0..50.
    Ic9Eppsm interface{}

    // IC10 ePPSM Status. The type is string with length: 0..50.
    Ic10Eppsm interface{}

    // IC11 ePPSM Status. The type is string with length: 0..50.
    Ic11Eppsm interface{}

    // IC1 PFM Status. The type is string with length: 0..50.
    Ic1Pfm interface{}

    // IC2 PFM Status. The type is string with length: 0..50.
    Ic2Pfm interface{}

    // IC3 PFM Status. The type is string with length: 0..50.
    Ic3Pfm interface{}

    // IC4 PFM Status. The type is string with length: 0..50.
    Ic4Pfm interface{}

    // IC5 PFM Status. The type is string with length: 0..50.
    Ic5Pfm interface{}

    // IC6 PFM Status. The type is string with length: 0..50.
    Ic6Pfm interface{}

    // IC7 PFM Status. The type is string with length: 0..50.
    Ic7Pfm interface{}

    // IC8 PFM Status. The type is string with length: 0..50.
    Ic8Pfm interface{}

    // IC9 PFM Status. The type is string with length: 0..50.
    Ic9Pfm interface{}

    // IC10 PFM Status. The type is string with length: 0..50.
    Ic10Pfm interface{}

    // IC11 PFM Status. The type is string with length: 0..50.
    Ic11Pfm interface{}

    // IC1 GST Status. The type is string with length: 0..50.
    Ic1Gst interface{}

    // IC2 GST Status. The type is string with length: 0..50.
    Ic2Gst interface{}

    // IC3 GST Status. The type is string with length: 0..50.
    Ic3Gst interface{}

    // IC4 GST Status. The type is string with length: 0..50.
    Ic4Gst interface{}

    // IC5 GST Status. The type is string with length: 0..50.
    Ic5Gst interface{}

    // IC6 GST Status. The type is string with length: 0..50.
    Ic6Gst interface{}

    // IC7 GST Status. The type is string with length: 0..50.
    Ic7Gst interface{}

    // IC8 GST Status. The type is string with length: 0..50.
    Ic8Gst interface{}

    // IC9 GST Status. The type is string with length: 0..50.
    Ic9Gst interface{}

    // IC10 GST Status. The type is string with length: 0..50.
    Ic10Gst interface{}

    // IC11 GST Status. The type is string with length: 0..50.
    Ic11Gst interface{}

    // IC1 CFM Status. The type is string with length: 0..50.
    Ic1Cfm interface{}

    // IC2 CFM Status. The type is string with length: 0..50.
    Ic2Cfm interface{}

    // IC3 CFM Status. The type is string with length: 0..50.
    Ic3Cfm interface{}

    // IC4 CFM Status. The type is string with length: 0..50.
    Ic4Cfm interface{}

    // IC5 CFM Status. The type is string with length: 0..50.
    Ic5Cfm interface{}

    // IC6 CFM Status. The type is string with length: 0..50.
    Ic6Cfm interface{}

    // IC7 CFM Status. The type is string with length: 0..50.
    Ic7Cfm interface{}

    // IC8 CFM Status. The type is string with length: 0..50.
    Ic8Cfm interface{}

    // IC9 CFM Status. The type is string with length: 0..50.
    Ic9Cfm interface{}

    // IC10 CFM Status. The type is string with length: 0..50.
    Ic10Cfm interface{}

    // IC11 CFM Status. The type is string with length: 0..50.
    Ic11Cfm interface{}

    // IC1 SCM Status. The type is string with length: 0..50.
    Ic1Scm interface{}

    // IC2 SCM Status. The type is string with length: 0..50.
    Ic2Scm interface{}

    // IC3 SCM Status. The type is string with length: 0..50.
    Ic3Scm interface{}

    // IC4 SCM Status. The type is string with length: 0..50.
    Ic4Scm interface{}

    // IC5 SCM Status. The type is string with length: 0..50.
    Ic5Scm interface{}

    // IC6 SCM Status. The type is string with length: 0..50.
    Ic6Scm interface{}

    // IC7 SCM Status. The type is string with length: 0..50.
    Ic7Scm interface{}

    // IC8 SCM Status. The type is string with length: 0..50.
    Ic8Scm interface{}

    // IC9 SCM Status. The type is string with length: 0..50.
    Ic9Scm interface{}

    // IC10 SCM Status. The type is string with length: 0..50.
    Ic10Scm interface{}

    // IC11 SCM Status. The type is string with length: 0..50.
    Ic11Scm interface{}

    // IC1 LOS Status. The type is string with length: 0..50.
    Ic1Los interface{}

    // IC2 LOS Status. The type is string with length: 0..50.
    Ic2Los interface{}

    // IC3 LOS Status. The type is string with length: 0..50.
    Ic3Los interface{}

    // IC4 LOS Status. The type is string with length: 0..50.
    Ic4Los interface{}

    // IC5 LOS Status. The type is string with length: 0..50.
    Ic5Los interface{}

    // IC6 LOS Status. The type is string with length: 0..50.
    Ic6Los interface{}

    // IC7 LOS Status. The type is string with length: 0..50.
    Ic7Los interface{}

    // IC8 LOS Status. The type is string with length: 0..50.
    Ic8Los interface{}

    // IC9 LOS Status. The type is string with length: 0..50.
    Ic9Los interface{}

    // IC10 LOS Status. The type is string with length: 0..50.
    Ic10Los interface{}

    // IC11 LOS Status. The type is string with length: 0..50.
    Ic11Los interface{}
}

func (inputClock *TimingCard_Nodes_Node_InputClock) GetEntityData() *types.CommonEntityData {
    inputClock.EntityData.YFilter = inputClock.YFilter
    inputClock.EntityData.YangName = "input-clock"
    inputClock.EntityData.BundleName = "cisco_ios_xr"
    inputClock.EntityData.ParentYangName = "node"
    inputClock.EntityData.SegmentPath = "input-clock"
    inputClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputClock.EntityData.Children = types.NewOrderedMap()
    inputClock.EntityData.Leafs = types.NewOrderedMap()
    inputClock.EntityData.Leafs.Append("ic1-valid", types.YLeaf{"Ic1Valid", inputClock.Ic1Valid})
    inputClock.EntityData.Leafs.Append("ic2-valid", types.YLeaf{"Ic2Valid", inputClock.Ic2Valid})
    inputClock.EntityData.Leafs.Append("ic3-valid", types.YLeaf{"Ic3Valid", inputClock.Ic3Valid})
    inputClock.EntityData.Leafs.Append("ic4-valid", types.YLeaf{"Ic4Valid", inputClock.Ic4Valid})
    inputClock.EntityData.Leafs.Append("ic5-valid", types.YLeaf{"Ic5Valid", inputClock.Ic5Valid})
    inputClock.EntityData.Leafs.Append("ic6-valid", types.YLeaf{"Ic6Valid", inputClock.Ic6Valid})
    inputClock.EntityData.Leafs.Append("ic7-valid", types.YLeaf{"Ic7Valid", inputClock.Ic7Valid})
    inputClock.EntityData.Leafs.Append("ic8-valid", types.YLeaf{"Ic8Valid", inputClock.Ic8Valid})
    inputClock.EntityData.Leafs.Append("ic9-valid", types.YLeaf{"Ic9Valid", inputClock.Ic9Valid})
    inputClock.EntityData.Leafs.Append("ic10-valid", types.YLeaf{"Ic10Valid", inputClock.Ic10Valid})
    inputClock.EntityData.Leafs.Append("ic11-valid", types.YLeaf{"Ic11Valid", inputClock.Ic11Valid})
    inputClock.EntityData.Leafs.Append("ic1-slot", types.YLeaf{"Ic1Slot", inputClock.Ic1Slot})
    inputClock.EntityData.Leafs.Append("ic2-slot", types.YLeaf{"Ic2Slot", inputClock.Ic2Slot})
    inputClock.EntityData.Leafs.Append("ic3-slot", types.YLeaf{"Ic3Slot", inputClock.Ic3Slot})
    inputClock.EntityData.Leafs.Append("ic4-slot", types.YLeaf{"Ic4Slot", inputClock.Ic4Slot})
    inputClock.EntityData.Leafs.Append("ic5-slot", types.YLeaf{"Ic5Slot", inputClock.Ic5Slot})
    inputClock.EntityData.Leafs.Append("ic6-slot", types.YLeaf{"Ic6Slot", inputClock.Ic6Slot})
    inputClock.EntityData.Leafs.Append("ic7-slot", types.YLeaf{"Ic7Slot", inputClock.Ic7Slot})
    inputClock.EntityData.Leafs.Append("ic8-slot", types.YLeaf{"Ic8Slot", inputClock.Ic8Slot})
    inputClock.EntityData.Leafs.Append("ic9-slot", types.YLeaf{"Ic9Slot", inputClock.Ic9Slot})
    inputClock.EntityData.Leafs.Append("ic10-slot", types.YLeaf{"Ic10Slot", inputClock.Ic10Slot})
    inputClock.EntityData.Leafs.Append("ic11-slot", types.YLeaf{"Ic11Slot", inputClock.Ic11Slot})
    inputClock.EntityData.Leafs.Append("ic1-split-xom", types.YLeaf{"Ic1SplitXom", inputClock.Ic1SplitXom})
    inputClock.EntityData.Leafs.Append("ic2-split-xom", types.YLeaf{"Ic2SplitXom", inputClock.Ic2SplitXom})
    inputClock.EntityData.Leafs.Append("ic3-split-xom", types.YLeaf{"Ic3SplitXom", inputClock.Ic3SplitXom})
    inputClock.EntityData.Leafs.Append("ic4-split-xom", types.YLeaf{"Ic4SplitXom", inputClock.Ic4SplitXom})
    inputClock.EntityData.Leafs.Append("ic5-split-xom", types.YLeaf{"Ic5SplitXom", inputClock.Ic5SplitXom})
    inputClock.EntityData.Leafs.Append("ic6-split-xom", types.YLeaf{"Ic6SplitXom", inputClock.Ic6SplitXom})
    inputClock.EntityData.Leafs.Append("ic7-split-xom", types.YLeaf{"Ic7SplitXom", inputClock.Ic7SplitXom})
    inputClock.EntityData.Leafs.Append("ic8-split-xom", types.YLeaf{"Ic8SplitXom", inputClock.Ic8SplitXom})
    inputClock.EntityData.Leafs.Append("ic9-split-xom", types.YLeaf{"Ic9SplitXom", inputClock.Ic9SplitXom})
    inputClock.EntityData.Leafs.Append("ic10-split-xom", types.YLeaf{"Ic10SplitXom", inputClock.Ic10SplitXom})
    inputClock.EntityData.Leafs.Append("ic11-split-xom", types.YLeaf{"Ic11SplitXom", inputClock.Ic11SplitXom})
    inputClock.EntityData.Leafs.Append("ic1-eppsm", types.YLeaf{"Ic1Eppsm", inputClock.Ic1Eppsm})
    inputClock.EntityData.Leafs.Append("ic2-eppsm", types.YLeaf{"Ic2Eppsm", inputClock.Ic2Eppsm})
    inputClock.EntityData.Leafs.Append("ic3-eppsm", types.YLeaf{"Ic3Eppsm", inputClock.Ic3Eppsm})
    inputClock.EntityData.Leafs.Append("ic4-eppsm", types.YLeaf{"Ic4Eppsm", inputClock.Ic4Eppsm})
    inputClock.EntityData.Leafs.Append("ic5-eppsm", types.YLeaf{"Ic5Eppsm", inputClock.Ic5Eppsm})
    inputClock.EntityData.Leafs.Append("ic6-eppsm", types.YLeaf{"Ic6Eppsm", inputClock.Ic6Eppsm})
    inputClock.EntityData.Leafs.Append("ic7-eppsm", types.YLeaf{"Ic7Eppsm", inputClock.Ic7Eppsm})
    inputClock.EntityData.Leafs.Append("ic8-eppsm", types.YLeaf{"Ic8Eppsm", inputClock.Ic8Eppsm})
    inputClock.EntityData.Leafs.Append("ic9-eppsm", types.YLeaf{"Ic9Eppsm", inputClock.Ic9Eppsm})
    inputClock.EntityData.Leafs.Append("ic10-eppsm", types.YLeaf{"Ic10Eppsm", inputClock.Ic10Eppsm})
    inputClock.EntityData.Leafs.Append("ic11-eppsm", types.YLeaf{"Ic11Eppsm", inputClock.Ic11Eppsm})
    inputClock.EntityData.Leafs.Append("ic1-pfm", types.YLeaf{"Ic1Pfm", inputClock.Ic1Pfm})
    inputClock.EntityData.Leafs.Append("ic2-pfm", types.YLeaf{"Ic2Pfm", inputClock.Ic2Pfm})
    inputClock.EntityData.Leafs.Append("ic3-pfm", types.YLeaf{"Ic3Pfm", inputClock.Ic3Pfm})
    inputClock.EntityData.Leafs.Append("ic4-pfm", types.YLeaf{"Ic4Pfm", inputClock.Ic4Pfm})
    inputClock.EntityData.Leafs.Append("ic5-pfm", types.YLeaf{"Ic5Pfm", inputClock.Ic5Pfm})
    inputClock.EntityData.Leafs.Append("ic6-pfm", types.YLeaf{"Ic6Pfm", inputClock.Ic6Pfm})
    inputClock.EntityData.Leafs.Append("ic7-pfm", types.YLeaf{"Ic7Pfm", inputClock.Ic7Pfm})
    inputClock.EntityData.Leafs.Append("ic8-pfm", types.YLeaf{"Ic8Pfm", inputClock.Ic8Pfm})
    inputClock.EntityData.Leafs.Append("ic9-pfm", types.YLeaf{"Ic9Pfm", inputClock.Ic9Pfm})
    inputClock.EntityData.Leafs.Append("ic10-pfm", types.YLeaf{"Ic10Pfm", inputClock.Ic10Pfm})
    inputClock.EntityData.Leafs.Append("ic11-pfm", types.YLeaf{"Ic11Pfm", inputClock.Ic11Pfm})
    inputClock.EntityData.Leafs.Append("ic1-gst", types.YLeaf{"Ic1Gst", inputClock.Ic1Gst})
    inputClock.EntityData.Leafs.Append("ic2-gst", types.YLeaf{"Ic2Gst", inputClock.Ic2Gst})
    inputClock.EntityData.Leafs.Append("ic3-gst", types.YLeaf{"Ic3Gst", inputClock.Ic3Gst})
    inputClock.EntityData.Leafs.Append("ic4-gst", types.YLeaf{"Ic4Gst", inputClock.Ic4Gst})
    inputClock.EntityData.Leafs.Append("ic5-gst", types.YLeaf{"Ic5Gst", inputClock.Ic5Gst})
    inputClock.EntityData.Leafs.Append("ic6-gst", types.YLeaf{"Ic6Gst", inputClock.Ic6Gst})
    inputClock.EntityData.Leafs.Append("ic7-gst", types.YLeaf{"Ic7Gst", inputClock.Ic7Gst})
    inputClock.EntityData.Leafs.Append("ic8-gst", types.YLeaf{"Ic8Gst", inputClock.Ic8Gst})
    inputClock.EntityData.Leafs.Append("ic9-gst", types.YLeaf{"Ic9Gst", inputClock.Ic9Gst})
    inputClock.EntityData.Leafs.Append("ic10-gst", types.YLeaf{"Ic10Gst", inputClock.Ic10Gst})
    inputClock.EntityData.Leafs.Append("ic11-gst", types.YLeaf{"Ic11Gst", inputClock.Ic11Gst})
    inputClock.EntityData.Leafs.Append("ic1-cfm", types.YLeaf{"Ic1Cfm", inputClock.Ic1Cfm})
    inputClock.EntityData.Leafs.Append("ic2-cfm", types.YLeaf{"Ic2Cfm", inputClock.Ic2Cfm})
    inputClock.EntityData.Leafs.Append("ic3-cfm", types.YLeaf{"Ic3Cfm", inputClock.Ic3Cfm})
    inputClock.EntityData.Leafs.Append("ic4-cfm", types.YLeaf{"Ic4Cfm", inputClock.Ic4Cfm})
    inputClock.EntityData.Leafs.Append("ic5-cfm", types.YLeaf{"Ic5Cfm", inputClock.Ic5Cfm})
    inputClock.EntityData.Leafs.Append("ic6-cfm", types.YLeaf{"Ic6Cfm", inputClock.Ic6Cfm})
    inputClock.EntityData.Leafs.Append("ic7-cfm", types.YLeaf{"Ic7Cfm", inputClock.Ic7Cfm})
    inputClock.EntityData.Leafs.Append("ic8-cfm", types.YLeaf{"Ic8Cfm", inputClock.Ic8Cfm})
    inputClock.EntityData.Leafs.Append("ic9-cfm", types.YLeaf{"Ic9Cfm", inputClock.Ic9Cfm})
    inputClock.EntityData.Leafs.Append("ic10-cfm", types.YLeaf{"Ic10Cfm", inputClock.Ic10Cfm})
    inputClock.EntityData.Leafs.Append("ic11-cfm", types.YLeaf{"Ic11Cfm", inputClock.Ic11Cfm})
    inputClock.EntityData.Leafs.Append("ic1-scm", types.YLeaf{"Ic1Scm", inputClock.Ic1Scm})
    inputClock.EntityData.Leafs.Append("ic2-scm", types.YLeaf{"Ic2Scm", inputClock.Ic2Scm})
    inputClock.EntityData.Leafs.Append("ic3-scm", types.YLeaf{"Ic3Scm", inputClock.Ic3Scm})
    inputClock.EntityData.Leafs.Append("ic4-scm", types.YLeaf{"Ic4Scm", inputClock.Ic4Scm})
    inputClock.EntityData.Leafs.Append("ic5-scm", types.YLeaf{"Ic5Scm", inputClock.Ic5Scm})
    inputClock.EntityData.Leafs.Append("ic6-scm", types.YLeaf{"Ic6Scm", inputClock.Ic6Scm})
    inputClock.EntityData.Leafs.Append("ic7-scm", types.YLeaf{"Ic7Scm", inputClock.Ic7Scm})
    inputClock.EntityData.Leafs.Append("ic8-scm", types.YLeaf{"Ic8Scm", inputClock.Ic8Scm})
    inputClock.EntityData.Leafs.Append("ic9-scm", types.YLeaf{"Ic9Scm", inputClock.Ic9Scm})
    inputClock.EntityData.Leafs.Append("ic10-scm", types.YLeaf{"Ic10Scm", inputClock.Ic10Scm})
    inputClock.EntityData.Leafs.Append("ic11-scm", types.YLeaf{"Ic11Scm", inputClock.Ic11Scm})
    inputClock.EntityData.Leafs.Append("ic1-los", types.YLeaf{"Ic1Los", inputClock.Ic1Los})
    inputClock.EntityData.Leafs.Append("ic2-los", types.YLeaf{"Ic2Los", inputClock.Ic2Los})
    inputClock.EntityData.Leafs.Append("ic3-los", types.YLeaf{"Ic3Los", inputClock.Ic3Los})
    inputClock.EntityData.Leafs.Append("ic4-los", types.YLeaf{"Ic4Los", inputClock.Ic4Los})
    inputClock.EntityData.Leafs.Append("ic5-los", types.YLeaf{"Ic5Los", inputClock.Ic5Los})
    inputClock.EntityData.Leafs.Append("ic6-los", types.YLeaf{"Ic6Los", inputClock.Ic6Los})
    inputClock.EntityData.Leafs.Append("ic7-los", types.YLeaf{"Ic7Los", inputClock.Ic7Los})
    inputClock.EntityData.Leafs.Append("ic8-los", types.YLeaf{"Ic8Los", inputClock.Ic8Los})
    inputClock.EntityData.Leafs.Append("ic9-los", types.YLeaf{"Ic9Los", inputClock.Ic9Los})
    inputClock.EntityData.Leafs.Append("ic10-los", types.YLeaf{"Ic10Los", inputClock.Ic10Los})
    inputClock.EntityData.Leafs.Append("ic11-los", types.YLeaf{"Ic11Los", inputClock.Ic11Los})

    inputClock.EntityData.YListKeys = []string {}

    return &(inputClock.EntityData)
}

// TimingCard_Nodes_Node_Pll
// Display the timing card PLL status information
type TimingCard_Nodes_Node_Pll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // T0 PLL state. The type is string with length: 0..50.
    T0PllState interface{}

    // T4 PLL state. The type is string with length: 0..50.
    T4PllState interface{}

    // 1588 PLL state. The type is string with length: 0..50.
    PtpPllState interface{}

    // T0 PLL selected. The type is string with length: 0..50.
    T0PllSelected interface{}

    // T4 PLL selected. The type is string with length: 0..50.
    T4PllSelected interface{}

    // 1588 PLL selected. The type is string with length: 0..50.
    PtpPllSelected interface{}

    // T0 PLL mode. The type is string with length: 0..50.
    T0PllMode interface{}

    // T4 PLL mode. The type is string with length: 0..50.
    T4PllMode interface{}

    // 1588 PLL mode. The type is string with length: 0..50.
    PtpPllMode interface{}

    // T0 PLL IC1 Priority. The type is interface{} with range: 0..255.
    T0PllIc1Prio interface{}

    // T0 PLL IC2 Priority. The type is interface{} with range: 0..255.
    T0PllIc2Prio interface{}

    // T0 PLL IC3 Priority. The type is interface{} with range: 0..255.
    T0PllIc3Prio interface{}

    // T0 PLL IC4 Priority. The type is interface{} with range: 0..255.
    T0PllIc4Prio interface{}

    // T0 PLL IC5 Priority. The type is interface{} with range: 0..255.
    T0PllIc5Prio interface{}

    // T0 PLL IC6 Priority. The type is interface{} with range: 0..255.
    T0PllIc6Prio interface{}

    // T0 PLL IC7 Priority. The type is interface{} with range: 0..255.
    T0PllIc7Prio interface{}

    // T0 PLL IC8 Priority. The type is interface{} with range: 0..255.
    T0PllIc8Prio interface{}

    // T0 PLL IC9 Priority. The type is interface{} with range: 0..255.
    T0PllIc9Prio interface{}

    // T0 PLL IC10 Priority. The type is interface{} with range: 0..255.
    T0PllIc10Prio interface{}

    // T0 PLL IC11 Priority. The type is interface{} with range: 0..255.
    T0PllIc11Prio interface{}

    // T4 PLL IC1 Priority. The type is interface{} with range: 0..255.
    T4PllIc1Prio interface{}

    // T4 PLL IC2 Priority. The type is interface{} with range: 0..255.
    T4PllIc2Prio interface{}

    // T4 PLL IC3 Priority. The type is interface{} with range: 0..255.
    T4PllIc3Prio interface{}

    // T4 PLL IC4 Priority. The type is interface{} with range: 0..255.
    T4PllIc4Prio interface{}

    // T4 PLL IC5 Priority. The type is interface{} with range: 0..255.
    T4PllIc5Prio interface{}

    // T4 PLL IC6 Priority. The type is interface{} with range: 0..255.
    T4PllIc6Prio interface{}

    // T4 PLL IC7 Priority. The type is interface{} with range: 0..255.
    T4PllIc7Prio interface{}

    // T4 PLL IC8 Priority. The type is interface{} with range: 0..255.
    T4PllIc8Prio interface{}

    // T4 PLL IC9 Priority. The type is interface{} with range: 0..255.
    T4PllIc9Prio interface{}

    // T4 PLL IC10 Priority. The type is interface{} with range: 0..255.
    T4PllIc10Prio interface{}

    // T4 PLL IC11 Priority. The type is interface{} with range: 0..255.
    T4PllIc11Prio interface{}

    // PTP PLL IC1 Priority. The type is interface{} with range: 0..255.
    PtpPllIc1Prio interface{}

    // PTP PLL IC2 Priority. The type is interface{} with range: 0..255.
    PtpPllIc2Prio interface{}

    // PTP PLL IC3 Priority. The type is interface{} with range: 0..255.
    PtpPllIc3Prio interface{}

    // PTP PLL IC4 Priority. The type is interface{} with range: 0..255.
    PtpPllIc4Prio interface{}

    // PTP PLL IC5 Priority. The type is interface{} with range: 0..255.
    PtpPllIc5Prio interface{}

    // PTP PLL IC6 Priority. The type is interface{} with range: 0..255.
    PtpPllIc6Prio interface{}

    // PTP PLL IC7 Priority. The type is interface{} with range: 0..255.
    PtpPllIc7Prio interface{}

    // PTP PLL IC8 Priority. The type is interface{} with range: 0..255.
    PtpPllIc8Prio interface{}

    // PTP PLL IC9 Priority. The type is interface{} with range: 0..255.
    PtpPllIc9Prio interface{}

    // PTP PLL IC10 Priority. The type is interface{} with range: 0..255.
    PtpPllIc10Prio interface{}

    // PTP PLL IC11 Priority. The type is interface{} with range: 0..255.
    PtpPllIc11Prio interface{}

    // IC Valid 1. The type is bool.
    Ic1Valid interface{}

    // IC Valid 2. The type is bool.
    Ic2Valid interface{}

    // IC Valid 3. The type is bool.
    Ic3Valid interface{}

    // IC Valid 4. The type is bool.
    Ic4Valid interface{}

    // IC Valid 5. The type is bool.
    Ic5Valid interface{}

    // IC Valid 6. The type is bool.
    Ic6Valid interface{}

    // IC Valid 7. The type is bool.
    Ic7Valid interface{}

    // IC Valid 8. The type is bool.
    Ic8Valid interface{}

    // IC Valid 9. The type is bool.
    Ic9Valid interface{}

    // IC Valid 10. The type is bool.
    Ic10Valid interface{}

    // IC Valid 11. The type is bool.
    Ic11Valid interface{}

    // T0 PLL Frequency Offset. The type is interface{} with range:
    // -2147483648..2147483647.
    T0PllFreqOffset interface{}

    // T4 PLL Frequency Offset. The type is interface{} with range:
    // -2147483648..2147483647.
    T4PllFreqOffset interface{}

    // PTP PLL Frequency Offset. The type is interface{} with range:
    // -2147483648..2147483647.
    PtpPllFreqOffset interface{}

    // T0 PLL Bandwidth. The type is string with length: 0..50.
    T0PllBandwidth interface{}

    // T4 PLL Bandwidth. The type is string with length: 0..50.
    T4PllBandwidth interface{}

    // PTP PLL Bandwidth. The type is string with length: 0..50.
    PtpPllBandwidth interface{}

    // T0 PLL PSL. The type is string with length: 0..50.
    T0PllPsl interface{}

    // T4 PLL PSL. The type is string with length: 0..50.
    T4PllPsl interface{}

    // PTP PLL PSL. The type is string with length: 0..50.
    PtpPllPsl interface{}

    // IC1 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic1QualMin interface{}

    // IC1 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic1QualMax interface{}

    // IC2 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic2QualMin interface{}

    // IC2 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic2QualMax interface{}

    // IC3 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic3QualMin interface{}

    // IC3 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic3QualMax interface{}

    // IC4 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic4QualMin interface{}

    // IC4 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic4QualMax interface{}

    // IC5 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic5QualMin interface{}

    // IC5 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic5QualMax interface{}

    // IC6 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic6QualMin interface{}

    // IC6 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic6QualMax interface{}

    // IC7 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic7QualMin interface{}

    // IC7 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic7QualMax interface{}

    // IC8 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic8QualMin interface{}

    // IC8 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic8QualMax interface{}

    // IC9 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic9QualMin interface{}

    // IC9 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic9QualMax interface{}

    // IC10 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic10QualMin interface{}

    // IC10 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic10QualMax interface{}

    // IC11 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic11QualMin interface{}

    // IC11 Pull-in/Hold-in Min. The type is string with length: 0..50.
    Ic11QualMax interface{}
}

func (pll *TimingCard_Nodes_Node_Pll) GetEntityData() *types.CommonEntityData {
    pll.EntityData.YFilter = pll.YFilter
    pll.EntityData.YangName = "pll"
    pll.EntityData.BundleName = "cisco_ios_xr"
    pll.EntityData.ParentYangName = "node"
    pll.EntityData.SegmentPath = "pll"
    pll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pll.EntityData.Children = types.NewOrderedMap()
    pll.EntityData.Leafs = types.NewOrderedMap()
    pll.EntityData.Leafs.Append("t0-pll-state", types.YLeaf{"T0PllState", pll.T0PllState})
    pll.EntityData.Leafs.Append("t4-pll-state", types.YLeaf{"T4PllState", pll.T4PllState})
    pll.EntityData.Leafs.Append("ptp-pll-state", types.YLeaf{"PtpPllState", pll.PtpPllState})
    pll.EntityData.Leafs.Append("t0-pll-selected", types.YLeaf{"T0PllSelected", pll.T0PllSelected})
    pll.EntityData.Leafs.Append("t4-pll-selected", types.YLeaf{"T4PllSelected", pll.T4PllSelected})
    pll.EntityData.Leafs.Append("ptp-pll-selected", types.YLeaf{"PtpPllSelected", pll.PtpPllSelected})
    pll.EntityData.Leafs.Append("t0-pll-mode", types.YLeaf{"T0PllMode", pll.T0PllMode})
    pll.EntityData.Leafs.Append("t4-pll-mode", types.YLeaf{"T4PllMode", pll.T4PllMode})
    pll.EntityData.Leafs.Append("ptp-pll-mode", types.YLeaf{"PtpPllMode", pll.PtpPllMode})
    pll.EntityData.Leafs.Append("t0-pll-ic1-prio", types.YLeaf{"T0PllIc1Prio", pll.T0PllIc1Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic2-prio", types.YLeaf{"T0PllIc2Prio", pll.T0PllIc2Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic3-prio", types.YLeaf{"T0PllIc3Prio", pll.T0PllIc3Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic4-prio", types.YLeaf{"T0PllIc4Prio", pll.T0PllIc4Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic5-prio", types.YLeaf{"T0PllIc5Prio", pll.T0PllIc5Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic6-prio", types.YLeaf{"T0PllIc6Prio", pll.T0PllIc6Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic7-prio", types.YLeaf{"T0PllIc7Prio", pll.T0PllIc7Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic8-prio", types.YLeaf{"T0PllIc8Prio", pll.T0PllIc8Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic9-prio", types.YLeaf{"T0PllIc9Prio", pll.T0PllIc9Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic10-prio", types.YLeaf{"T0PllIc10Prio", pll.T0PllIc10Prio})
    pll.EntityData.Leafs.Append("t0-pll-ic11-prio", types.YLeaf{"T0PllIc11Prio", pll.T0PllIc11Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic1-prio", types.YLeaf{"T4PllIc1Prio", pll.T4PllIc1Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic2-prio", types.YLeaf{"T4PllIc2Prio", pll.T4PllIc2Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic3-prio", types.YLeaf{"T4PllIc3Prio", pll.T4PllIc3Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic4-prio", types.YLeaf{"T4PllIc4Prio", pll.T4PllIc4Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic5-prio", types.YLeaf{"T4PllIc5Prio", pll.T4PllIc5Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic6-prio", types.YLeaf{"T4PllIc6Prio", pll.T4PllIc6Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic7-prio", types.YLeaf{"T4PllIc7Prio", pll.T4PllIc7Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic8-prio", types.YLeaf{"T4PllIc8Prio", pll.T4PllIc8Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic9-prio", types.YLeaf{"T4PllIc9Prio", pll.T4PllIc9Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic10-prio", types.YLeaf{"T4PllIc10Prio", pll.T4PllIc10Prio})
    pll.EntityData.Leafs.Append("t4-pll-ic11-prio", types.YLeaf{"T4PllIc11Prio", pll.T4PllIc11Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic1-prio", types.YLeaf{"PtpPllIc1Prio", pll.PtpPllIc1Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic2-prio", types.YLeaf{"PtpPllIc2Prio", pll.PtpPllIc2Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic3-prio", types.YLeaf{"PtpPllIc3Prio", pll.PtpPllIc3Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic4-prio", types.YLeaf{"PtpPllIc4Prio", pll.PtpPllIc4Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic5-prio", types.YLeaf{"PtpPllIc5Prio", pll.PtpPllIc5Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic6-prio", types.YLeaf{"PtpPllIc6Prio", pll.PtpPllIc6Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic7-prio", types.YLeaf{"PtpPllIc7Prio", pll.PtpPllIc7Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic8-prio", types.YLeaf{"PtpPllIc8Prio", pll.PtpPllIc8Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic9-prio", types.YLeaf{"PtpPllIc9Prio", pll.PtpPllIc9Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic10-prio", types.YLeaf{"PtpPllIc10Prio", pll.PtpPllIc10Prio})
    pll.EntityData.Leafs.Append("ptp-pll-ic11-prio", types.YLeaf{"PtpPllIc11Prio", pll.PtpPllIc11Prio})
    pll.EntityData.Leafs.Append("ic1-valid", types.YLeaf{"Ic1Valid", pll.Ic1Valid})
    pll.EntityData.Leafs.Append("ic2-valid", types.YLeaf{"Ic2Valid", pll.Ic2Valid})
    pll.EntityData.Leafs.Append("ic3-valid", types.YLeaf{"Ic3Valid", pll.Ic3Valid})
    pll.EntityData.Leafs.Append("ic4-valid", types.YLeaf{"Ic4Valid", pll.Ic4Valid})
    pll.EntityData.Leafs.Append("ic5-valid", types.YLeaf{"Ic5Valid", pll.Ic5Valid})
    pll.EntityData.Leafs.Append("ic6-valid", types.YLeaf{"Ic6Valid", pll.Ic6Valid})
    pll.EntityData.Leafs.Append("ic7-valid", types.YLeaf{"Ic7Valid", pll.Ic7Valid})
    pll.EntityData.Leafs.Append("ic8-valid", types.YLeaf{"Ic8Valid", pll.Ic8Valid})
    pll.EntityData.Leafs.Append("ic9-valid", types.YLeaf{"Ic9Valid", pll.Ic9Valid})
    pll.EntityData.Leafs.Append("ic10-valid", types.YLeaf{"Ic10Valid", pll.Ic10Valid})
    pll.EntityData.Leafs.Append("ic11-valid", types.YLeaf{"Ic11Valid", pll.Ic11Valid})
    pll.EntityData.Leafs.Append("t0-pll-freq-offset", types.YLeaf{"T0PllFreqOffset", pll.T0PllFreqOffset})
    pll.EntityData.Leafs.Append("t4-pll-freq-offset", types.YLeaf{"T4PllFreqOffset", pll.T4PllFreqOffset})
    pll.EntityData.Leafs.Append("ptp-pll-freq-offset", types.YLeaf{"PtpPllFreqOffset", pll.PtpPllFreqOffset})
    pll.EntityData.Leafs.Append("t0-pll-bandwidth", types.YLeaf{"T0PllBandwidth", pll.T0PllBandwidth})
    pll.EntityData.Leafs.Append("t4-pll-bandwidth", types.YLeaf{"T4PllBandwidth", pll.T4PllBandwidth})
    pll.EntityData.Leafs.Append("ptp-pll-bandwidth", types.YLeaf{"PtpPllBandwidth", pll.PtpPllBandwidth})
    pll.EntityData.Leafs.Append("t0-pll-psl", types.YLeaf{"T0PllPsl", pll.T0PllPsl})
    pll.EntityData.Leafs.Append("t4-pll-psl", types.YLeaf{"T4PllPsl", pll.T4PllPsl})
    pll.EntityData.Leafs.Append("ptp-pll-psl", types.YLeaf{"PtpPllPsl", pll.PtpPllPsl})
    pll.EntityData.Leafs.Append("ic1-qual-min", types.YLeaf{"Ic1QualMin", pll.Ic1QualMin})
    pll.EntityData.Leafs.Append("ic1-qual-max", types.YLeaf{"Ic1QualMax", pll.Ic1QualMax})
    pll.EntityData.Leafs.Append("ic2-qual-min", types.YLeaf{"Ic2QualMin", pll.Ic2QualMin})
    pll.EntityData.Leafs.Append("ic2-qual-max", types.YLeaf{"Ic2QualMax", pll.Ic2QualMax})
    pll.EntityData.Leafs.Append("ic3-qual-min", types.YLeaf{"Ic3QualMin", pll.Ic3QualMin})
    pll.EntityData.Leafs.Append("ic3-qual-max", types.YLeaf{"Ic3QualMax", pll.Ic3QualMax})
    pll.EntityData.Leafs.Append("ic4-qual-min", types.YLeaf{"Ic4QualMin", pll.Ic4QualMin})
    pll.EntityData.Leafs.Append("ic4-qual-max", types.YLeaf{"Ic4QualMax", pll.Ic4QualMax})
    pll.EntityData.Leafs.Append("ic5-qual-min", types.YLeaf{"Ic5QualMin", pll.Ic5QualMin})
    pll.EntityData.Leafs.Append("ic5-qual-max", types.YLeaf{"Ic5QualMax", pll.Ic5QualMax})
    pll.EntityData.Leafs.Append("ic6-qual-min", types.YLeaf{"Ic6QualMin", pll.Ic6QualMin})
    pll.EntityData.Leafs.Append("ic6-qual-max", types.YLeaf{"Ic6QualMax", pll.Ic6QualMax})
    pll.EntityData.Leafs.Append("ic7-qual-min", types.YLeaf{"Ic7QualMin", pll.Ic7QualMin})
    pll.EntityData.Leafs.Append("ic7-qual-max", types.YLeaf{"Ic7QualMax", pll.Ic7QualMax})
    pll.EntityData.Leafs.Append("ic8-qual-min", types.YLeaf{"Ic8QualMin", pll.Ic8QualMin})
    pll.EntityData.Leafs.Append("ic8-qual-max", types.YLeaf{"Ic8QualMax", pll.Ic8QualMax})
    pll.EntityData.Leafs.Append("ic9-qual-min", types.YLeaf{"Ic9QualMin", pll.Ic9QualMin})
    pll.EntityData.Leafs.Append("ic9-qual-max", types.YLeaf{"Ic9QualMax", pll.Ic9QualMax})
    pll.EntityData.Leafs.Append("ic10-qual-min", types.YLeaf{"Ic10QualMin", pll.Ic10QualMin})
    pll.EntityData.Leafs.Append("ic10-qual-max", types.YLeaf{"Ic10QualMax", pll.Ic10QualMax})
    pll.EntityData.Leafs.Append("ic11-qual-min", types.YLeaf{"Ic11QualMin", pll.Ic11QualMin})
    pll.EntityData.Leafs.Append("ic11-qual-max", types.YLeaf{"Ic11QualMax", pll.Ic11QualMax})

    pll.EntityData.YListKeys = []string {}

    return &(pll.EntityData)
}

