// This module contains a collection of YANG definitions for
// monitoring vlans in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package spanning_tree_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package spanning_tree_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree-oper stp-details}", reflect.TypeOf(StpDetails{}))
    ydk.RegisterEntity("Cisco-IOS-XE-spanning-tree-oper:stp-details", reflect.TypeOf(StpDetails{}))
}

// StpPortBpduguard represents Accept BPDUs on this interface
type StpPortBpduguard string

const (
    StpPortBpduguard_stp_port_bpduguard_disable StpPortBpduguard = "stp-port-bpduguard-disable"

    StpPortBpduguard_stp_port_bpduguard_enable StpPortBpduguard = "stp-port-bpduguard-enable"

    StpPortBpduguard_stp_port_bpduguard_default StpPortBpduguard = "stp-port-bpduguard-default"
)

// StpLinkRole represents Type definition for the different link types
type StpLinkRole string

const (
    StpLinkRole_stp_auto StpLinkRole = "stp-auto"

    StpLinkRole_stp_point_to_point StpLinkRole = "stp-point-to-point"

    StpLinkRole_stp_shared StpLinkRole = "stp-shared"
)

// StpMode represents Spanning tree operating mode
type StpMode string

const (
    StpMode_stp_mode_pvst StpMode = "stp-mode-pvst"

    StpMode_stp_mode_rapid_pvst StpMode = "stp-mode-rapid-pvst"

    StpMode_stp_mode_mst StpMode = "stp-mode-mst"
)

// StpPortRole represents Spanning Tree Protocol port roles
type StpPortRole string

const (
    StpPortRole_stp_master StpPortRole = "stp-master"

    StpPortRole_stp_alternate StpPortRole = "stp-alternate"

    StpPortRole_stp_root StpPortRole = "stp-root"

    StpPortRole_stp_designated StpPortRole = "stp-designated"

    StpPortRole_stp_backup StpPortRole = "stp-backup"
)

// StpPortState represents Spanning Tree Protocol port states
type StpPortState string

const (
    StpPortState_stp_disabled StpPortState = "stp-disabled"

    StpPortState_stp_blocking StpPortState = "stp-blocking"

    StpPortState_stp_listening StpPortState = "stp-listening"

    StpPortState_stp_learning StpPortState = "stp-learning"

    StpPortState_stp_forwarding StpPortState = "stp-forwarding"

    StpPortState_stp_broken StpPortState = "stp-broken"

    StpPortState_stp_invalid StpPortState = "stp-invalid"
)

// StpPortBpdufilter represents Send or receive BPDUs on this interface
type StpPortBpdufilter string

const (
    StpPortBpdufilter_stp_port_bpdufilter_disable StpPortBpdufilter = "stp-port-bpdufilter-disable"

    StpPortBpdufilter_stp_port_bpdufilter_enable StpPortBpdufilter = "stp-port-bpdufilter-enable"

    StpPortBpdufilter_stp_port_bpdufilter_default StpPortBpdufilter = "stp-port-bpdufilter-default"
)

// StpPortGuard represents Interface's spanning tree guard mode
type StpPortGuard string

const (
    StpPortGuard_stp_port_guard_default StpPortGuard = "stp-port-guard-default"

    StpPortGuard_stp_port_guard_root StpPortGuard = "stp-port-guard-root"

    StpPortGuard_stp_port_guard_loop StpPortGuard = "stp-port-guard-loop"

    StpPortGuard_stp_port_guard_none StpPortGuard = "stp-port-guard-none"
)

// StpDetails
// Top-level container for spanning tree operational data
type StpDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of mst/rapid-pvst spanning-tree, keyed by instance name. The type is
    // slice of StpDetails_StpDetail.
    StpDetail []*StpDetails_StpDetail

    // Global state data.
    StpGlobal StpDetails_StpGlobal
}

func (stpDetails *StpDetails) GetEntityData() *types.CommonEntityData {
    stpDetails.EntityData.YFilter = stpDetails.YFilter
    stpDetails.EntityData.YangName = "stp-details"
    stpDetails.EntityData.BundleName = "cisco_ios_xe"
    stpDetails.EntityData.ParentYangName = "Cisco-IOS-XE-spanning-tree-oper"
    stpDetails.EntityData.SegmentPath = "Cisco-IOS-XE-spanning-tree-oper:stp-details"
    stpDetails.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpDetails.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpDetails.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpDetails.EntityData.Children = types.NewOrderedMap()
    stpDetails.EntityData.Children.Append("stp-detail", types.YChild{"StpDetail", nil})
    for i := range stpDetails.StpDetail {
        stpDetails.EntityData.Children.Append(types.GetSegmentPath(stpDetails.StpDetail[i]), types.YChild{"StpDetail", stpDetails.StpDetail[i]})
    }
    stpDetails.EntityData.Children.Append("stp-global", types.YChild{"StpGlobal", &stpDetails.StpGlobal})
    stpDetails.EntityData.Leafs = types.NewOrderedMap()

    stpDetails.EntityData.YListKeys = []string {}

    return &(stpDetails.EntityData)
}

// StpDetails_StpDetail
// List of mst/rapid-pvst spanning-tree, keyed by instance name
type StpDetails_StpDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Spanning-tree enabled mode and id. The type is
    // string.
    Instance interface{}

    // The interval between periodic transmissions of configuration messages by
    // designated ports. The type is interface{} with range:
    // -2147483648..2147483647.
    HelloTime interface{}

    // The maximum age of the information transmitted by the bridge when it is the
    // root bridge. The type is interface{} with range: -2147483648..2147483647.
    MaxAge interface{}

    // The delay used by STP bridges to transition root and designated ports to
    // forwarding. The type is interface{} with range: -2147483648..2147483647.
    ForwardingDelay interface{}

    // The maximum number of BPDUs per second that the switch can send from an
    // interface. The type is interface{} with range: 0..4294967295.
    HoldCount interface{}

    // The manageable component of the Bridge Identifier. The type is interface{}
    // with range: 0..65535.
    BridgePriority interface{}

    // A unique 48-bit Universally Administered MAC Address assigned to the
    // bridge. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BridgeAddress interface{}

    // The bridge priority of the root of the spanning tree, as determined by the
    // Spanning Tree Protocol, as executed by this node. The type is interface{}
    // with range: 0..4294967295.
    DesignatedRootPriority interface{}

    // The bridge address of the root of the spanning tree, as determined by the
    // Spanning Tree Protocol, as executed by this node. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DesignatedRootAddress interface{}

    // The port number of the port which offers the lowest cost path from this
    // bridge to the root bridge. The type is interface{} with range: 0..65535.
    RootPort interface{}

    // The cost of the path to the root as seen from this bridge. The type is
    // interface{} with range: 0..18446744073709551615.
    RootCost interface{}

    // This time value determines the interval length during which no more than
    // two Configuration bridge PDUs shall be transmitted by this node. The type
    // is interface{} with range: 0..18446744073709551615.
    HoldTime interface{}

    // The total number of topology changes detected by this bridge since the
    // management entity was last reset or initialized. The type is interface{}
    // with range: 0..18446744073709551615.
    TopologyChanges interface{}

    // The time of the last topology change that was detected by the bridge
    // entity.The time is POSIX time UTC. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    TimeOfLastTopologyChange interface{}

    // List of interfaces on which STP is enable.
    Interfaces StpDetails_StpDetail_Interfaces
}

func (stpDetail *StpDetails_StpDetail) GetEntityData() *types.CommonEntityData {
    stpDetail.EntityData.YFilter = stpDetail.YFilter
    stpDetail.EntityData.YangName = "stp-detail"
    stpDetail.EntityData.BundleName = "cisco_ios_xe"
    stpDetail.EntityData.ParentYangName = "stp-details"
    stpDetail.EntityData.SegmentPath = "stp-detail" + types.AddKeyToken(stpDetail.Instance, "instance")
    stpDetail.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpDetail.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpDetail.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpDetail.EntityData.Children = types.NewOrderedMap()
    stpDetail.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &stpDetail.Interfaces})
    stpDetail.EntityData.Leafs = types.NewOrderedMap()
    stpDetail.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", stpDetail.Instance})
    stpDetail.EntityData.Leafs.Append("hello-time", types.YLeaf{"HelloTime", stpDetail.HelloTime})
    stpDetail.EntityData.Leafs.Append("max-age", types.YLeaf{"MaxAge", stpDetail.MaxAge})
    stpDetail.EntityData.Leafs.Append("forwarding-delay", types.YLeaf{"ForwardingDelay", stpDetail.ForwardingDelay})
    stpDetail.EntityData.Leafs.Append("hold-count", types.YLeaf{"HoldCount", stpDetail.HoldCount})
    stpDetail.EntityData.Leafs.Append("bridge-priority", types.YLeaf{"BridgePriority", stpDetail.BridgePriority})
    stpDetail.EntityData.Leafs.Append("bridge-address", types.YLeaf{"BridgeAddress", stpDetail.BridgeAddress})
    stpDetail.EntityData.Leafs.Append("designated-root-priority", types.YLeaf{"DesignatedRootPriority", stpDetail.DesignatedRootPriority})
    stpDetail.EntityData.Leafs.Append("designated-root-address", types.YLeaf{"DesignatedRootAddress", stpDetail.DesignatedRootAddress})
    stpDetail.EntityData.Leafs.Append("root-port", types.YLeaf{"RootPort", stpDetail.RootPort})
    stpDetail.EntityData.Leafs.Append("root-cost", types.YLeaf{"RootCost", stpDetail.RootCost})
    stpDetail.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", stpDetail.HoldTime})
    stpDetail.EntityData.Leafs.Append("topology-changes", types.YLeaf{"TopologyChanges", stpDetail.TopologyChanges})
    stpDetail.EntityData.Leafs.Append("time-of-last-topology-change", types.YLeaf{"TimeOfLastTopologyChange", stpDetail.TimeOfLastTopologyChange})

    stpDetail.EntityData.YListKeys = []string {"Instance"}

    return &(stpDetail.EntityData)
}

// StpDetails_StpDetail_Interfaces
// List of interfaces on which STP is enable
type StpDetails_StpDetail_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interfaces on which STP is enable. The type is slice of
    // StpDetails_StpDetail_Interfaces_Interface.
    Interface []*StpDetails_StpDetail_Interfaces_Interface
}

func (interfaces *StpDetails_StpDetail_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "stp-detail"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// StpDetails_StpDetail_Interfaces_Interface
// List of interfaces on which STP is enable
type StpDetails_StpDetail_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the STP ethernet interface. The type
    // is string.
    Name interface{}

    // The port's contribution, when it is the Root Port, to the Root Path Cost
    // for the Bridge. The type is interface{} with range:
    // 0..18446744073709551615.
    Cost interface{}

    // The manageable component of the Port Identifier, also known as the Port
    // Priority. The type is interface{} with range: 0..65535.
    PortPriority interface{}

    // The port number of the bridge port. The type is interface{} with range:
    // 0..65535.
    PortNum interface{}

    // The current role of the bridge port. The type is StpPortRole.
    Role interface{}

    // The current state of the bridge port. The type is StpPortState.
    State interface{}

    // The bridge priority of the bridge recorded as the root in the configuration
    // BPDUs transmitted by the designated bridge for the segment to which the
    // port is attached. The type is interface{} with range: 0..4294967295.
    DesignatedRootPriority interface{}

    // The bridge address of the bridge recorded as the root in the configuration
    // BPDUs transmitted by the designated bridge for the segment to which the
    // port is attached. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DesignatedRootAddress interface{}

    // The path cost of the Designated Port of the segment connected to this port.
    // The type is interface{} with range: 0..4294967295.
    DesignatedCost interface{}

    // The bridge priority of the bridge that this port considers to be the
    // designated bridge for this port's segment. The type is interface{} with
    // range: 0..4294967295.
    DesignatedBridgePriority interface{}

    // The bridge address of the bridge that this port considers to be the
    // designated bridge for this port's segment. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DesignatedBridgeAddress interface{}

    // The Port priority of the port on the Designated Bridge for this port's
    // segment, two octet string. The type is interface{} with range: 0..65535.
    DesignatedPortPriority interface{}

    // The Port number of the port on the Designated Bridge for this port's
    // segment, two octet string. The type is interface{} with range: 0..65535.
    DesignatedPortNum interface{}

    // The number of times this port has transitioned from the Learning state to
    // the Forwarding state. The type is interface{} with range:
    // 0..18446744073709551615.
    ForwardTransitions interface{}

    // Interface's link type. The type is StpLinkRole.
    LinkType interface{}

    // Interface's spanning tree guard mode. The type is StpPortGuard.
    Guard interface{}

    // BPDU guard on this interface. The type is StpPortBpduguard.
    BpduGuard interface{}

    // BPDU filter on this interface. The type is StpPortBpdufilter.
    BpduFilter interface{}

    // The number of BPDU packet sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BpduSent interface{}

    // The number of BPDU packet received. The type is interface{} with range:
    // 0..18446744073709551615.
    BpduReceived interface{}
}

func (self *StpDetails_StpDetail_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Name, "name")
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})
    self.EntityData.Leafs.Append("port-priority", types.YLeaf{"PortPriority", self.PortPriority})
    self.EntityData.Leafs.Append("port-num", types.YLeaf{"PortNum", self.PortNum})
    self.EntityData.Leafs.Append("role", types.YLeaf{"Role", self.Role})
    self.EntityData.Leafs.Append("state", types.YLeaf{"State", self.State})
    self.EntityData.Leafs.Append("designated-root-priority", types.YLeaf{"DesignatedRootPriority", self.DesignatedRootPriority})
    self.EntityData.Leafs.Append("designated-root-address", types.YLeaf{"DesignatedRootAddress", self.DesignatedRootAddress})
    self.EntityData.Leafs.Append("designated-cost", types.YLeaf{"DesignatedCost", self.DesignatedCost})
    self.EntityData.Leafs.Append("designated-bridge-priority", types.YLeaf{"DesignatedBridgePriority", self.DesignatedBridgePriority})
    self.EntityData.Leafs.Append("designated-bridge-address", types.YLeaf{"DesignatedBridgeAddress", self.DesignatedBridgeAddress})
    self.EntityData.Leafs.Append("designated-port-priority", types.YLeaf{"DesignatedPortPriority", self.DesignatedPortPriority})
    self.EntityData.Leafs.Append("designated-port-num", types.YLeaf{"DesignatedPortNum", self.DesignatedPortNum})
    self.EntityData.Leafs.Append("forward-transitions", types.YLeaf{"ForwardTransitions", self.ForwardTransitions})
    self.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", self.LinkType})
    self.EntityData.Leafs.Append("guard", types.YLeaf{"Guard", self.Guard})
    self.EntityData.Leafs.Append("bpdu-guard", types.YLeaf{"BpduGuard", self.BpduGuard})
    self.EntityData.Leafs.Append("bpdu-filter", types.YLeaf{"BpduFilter", self.BpduFilter})
    self.EntityData.Leafs.Append("bpdu-sent", types.YLeaf{"BpduSent", self.BpduSent})
    self.EntityData.Leafs.Append("bpdu-received", types.YLeaf{"BpduReceived", self.BpduReceived})

    self.EntityData.YListKeys = []string {"Name"}

    return &(self.EntityData)
}

// StpDetails_StpGlobal
// Global state data
// This type is a presence type.
type StpDetails_StpGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Spanning tree mode enabled on the device. The type is StpMode.
    Mode interface{}

    // Enable STP Bridge Assurance feature. The type is interface{}.
    BridgeAssurance interface{}

    // Enable loopguard by default. The type is interface{}.
    LoopGuard interface{}

    // Enable portfast bpdu guard. The type is interface{}.
    BpduGuard interface{}

    // Enable portfast bpdu filter. The type is interface{}.
    BpduFilter interface{}

    // Enable guard to protect against etherchannel misconfiguration. The type is
    // interface{}.
    EtherchannelMisconfigGuard interface{}

    // Global state for MST ONLY.
    MstOnly StpDetails_StpGlobal_MstOnly
}

func (stpGlobal *StpDetails_StpGlobal) GetEntityData() *types.CommonEntityData {
    stpGlobal.EntityData.YFilter = stpGlobal.YFilter
    stpGlobal.EntityData.YangName = "stp-global"
    stpGlobal.EntityData.BundleName = "cisco_ios_xe"
    stpGlobal.EntityData.ParentYangName = "stp-details"
    stpGlobal.EntityData.SegmentPath = "stp-global"
    stpGlobal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpGlobal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpGlobal.EntityData.Children = types.NewOrderedMap()
    stpGlobal.EntityData.Children.Append("mst-only", types.YChild{"MstOnly", &stpGlobal.MstOnly})
    stpGlobal.EntityData.Leafs = types.NewOrderedMap()
    stpGlobal.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", stpGlobal.Mode})
    stpGlobal.EntityData.Leafs.Append("bridge-assurance", types.YLeaf{"BridgeAssurance", stpGlobal.BridgeAssurance})
    stpGlobal.EntityData.Leafs.Append("loop-guard", types.YLeaf{"LoopGuard", stpGlobal.LoopGuard})
    stpGlobal.EntityData.Leafs.Append("bpdu-guard", types.YLeaf{"BpduGuard", stpGlobal.BpduGuard})
    stpGlobal.EntityData.Leafs.Append("bpdu-filter", types.YLeaf{"BpduFilter", stpGlobal.BpduFilter})
    stpGlobal.EntityData.Leafs.Append("etherchannel-misconfig-guard", types.YLeaf{"EtherchannelMisconfigGuard", stpGlobal.EtherchannelMisconfigGuard})

    stpGlobal.EntityData.YListKeys = []string {}

    return &(stpGlobal.EntityData)
}

// StpDetails_StpGlobal_MstOnly
// Global state for MST ONLY
type StpDetails_StpGlobal_MstOnly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration revision number(used by MSTP only). The type is interface{}
    // with range: 0..65535.
    MstConfigRevision interface{}

    // Configuration name(used by MSTP only). The type is string.
    MstConfigName interface{}

    // The max hops value for the spanning tree(used by MSTP only). The type is
    // interface{} with range: 0..65535.
    MaxHops interface{}
}

func (mstOnly *StpDetails_StpGlobal_MstOnly) GetEntityData() *types.CommonEntityData {
    mstOnly.EntityData.YFilter = mstOnly.YFilter
    mstOnly.EntityData.YangName = "mst-only"
    mstOnly.EntityData.BundleName = "cisco_ios_xe"
    mstOnly.EntityData.ParentYangName = "stp-global"
    mstOnly.EntityData.SegmentPath = "mst-only"
    mstOnly.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mstOnly.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mstOnly.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mstOnly.EntityData.Children = types.NewOrderedMap()
    mstOnly.EntityData.Leafs = types.NewOrderedMap()
    mstOnly.EntityData.Leafs.Append("mst-config-revision", types.YLeaf{"MstConfigRevision", mstOnly.MstConfigRevision})
    mstOnly.EntityData.Leafs.Append("mst-config-name", types.YLeaf{"MstConfigName", mstOnly.MstConfigName})
    mstOnly.EntityData.Leafs.Append("max-hops", types.YLeaf{"MaxHops", mstOnly.MaxHops})

    mstOnly.EntityData.YListKeys = []string {}

    return &(mstOnly.EntityData)
}

