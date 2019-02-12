// The Bridge MIB module for managing devices that support
// IEEE 802.1D.
// 
// Copyright (C) The Internet Society (2005).  This version of
// this MIB module is part of RFC 4188; see the RFC itself for
// full legal notices.
package bridge_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bridge_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:BRIDGE-MIB BRIDGE-MIB}", reflect.TypeOf(BRIDGEMIB{}))
    ydk.RegisterEntity("BRIDGE-MIB:BRIDGE-MIB", reflect.TypeOf(BRIDGEMIB{}))
}

// BRIDGEMIB
type BRIDGEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Dot1dBase BRIDGEMIB_Dot1dBase

    
    Dot1dStp BRIDGEMIB_Dot1dStp

    
    Dot1dTp BRIDGEMIB_Dot1dTp

    // A table that contains generic information about every port that is
    // associated with this bridge.  Transparent, source-route, and srt ports are
    // included.
    Dot1dBasePortTable BRIDGEMIB_Dot1dBasePortTable

    // A table that contains port-specific information for the Spanning Tree
    // Protocol.
    Dot1dStpPortTable BRIDGEMIB_Dot1dStpPortTable

    // A table that contains information about unicast entries for which the
    // bridge has forwarding and/or filtering information.  This information is
    // used by the transparent bridging function in determining how to propagate a
    // received frame.
    Dot1dTpFdbTable BRIDGEMIB_Dot1dTpFdbTable

    // A table that contains information about every port that is associated with
    // this transparent bridge.
    Dot1dTpPortTable BRIDGEMIB_Dot1dTpPortTable

    // A table containing filtering information configured into the bridge by
    // (local or network) management specifying the set of ports to which frames
    // received from specific ports and containing specific destination addresses
    // are allowed to be forwarded.  The value of zero in this table, as the port
    // number from which frames with a specific destination address are received,
    // is used to specify all ports for which there is no specific entry in this
    // table for that particular destination address.  Entries are valid for
    // unicast and for group/broadcast addresses.
    Dot1dStaticTable BRIDGEMIB_Dot1dStaticTable
}

func (bRIDGEMIB *BRIDGEMIB) GetEntityData() *types.CommonEntityData {
    bRIDGEMIB.EntityData.YFilter = bRIDGEMIB.YFilter
    bRIDGEMIB.EntityData.YangName = "BRIDGE-MIB"
    bRIDGEMIB.EntityData.BundleName = "cisco_ios_xe"
    bRIDGEMIB.EntityData.ParentYangName = "BRIDGE-MIB"
    bRIDGEMIB.EntityData.SegmentPath = "BRIDGE-MIB:BRIDGE-MIB"
    bRIDGEMIB.EntityData.AbsolutePath = bRIDGEMIB.EntityData.SegmentPath
    bRIDGEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bRIDGEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bRIDGEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bRIDGEMIB.EntityData.Children = types.NewOrderedMap()
    bRIDGEMIB.EntityData.Children.Append("dot1dBase", types.YChild{"Dot1dBase", &bRIDGEMIB.Dot1dBase})
    bRIDGEMIB.EntityData.Children.Append("dot1dStp", types.YChild{"Dot1dStp", &bRIDGEMIB.Dot1dStp})
    bRIDGEMIB.EntityData.Children.Append("dot1dTp", types.YChild{"Dot1dTp", &bRIDGEMIB.Dot1dTp})
    bRIDGEMIB.EntityData.Children.Append("dot1dBasePortTable", types.YChild{"Dot1dBasePortTable", &bRIDGEMIB.Dot1dBasePortTable})
    bRIDGEMIB.EntityData.Children.Append("dot1dStpPortTable", types.YChild{"Dot1dStpPortTable", &bRIDGEMIB.Dot1dStpPortTable})
    bRIDGEMIB.EntityData.Children.Append("dot1dTpFdbTable", types.YChild{"Dot1dTpFdbTable", &bRIDGEMIB.Dot1dTpFdbTable})
    bRIDGEMIB.EntityData.Children.Append("dot1dTpPortTable", types.YChild{"Dot1dTpPortTable", &bRIDGEMIB.Dot1dTpPortTable})
    bRIDGEMIB.EntityData.Children.Append("dot1dStaticTable", types.YChild{"Dot1dStaticTable", &bRIDGEMIB.Dot1dStaticTable})
    bRIDGEMIB.EntityData.Leafs = types.NewOrderedMap()

    bRIDGEMIB.EntityData.YListKeys = []string {}

    return &(bRIDGEMIB.EntityData)
}

// BRIDGEMIB_Dot1dBase
type BRIDGEMIB_Dot1dBase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The MAC address used by this bridge when it must be referred to in a unique
    // fashion.  It is recommended that this be the numerically smallest MAC
    // address of all ports that belong to this bridge.  However, it is only 
    // required to be unique.  When concatenated with dot1dStpPriority, a unique
    // BridgeIdentifier is formed, which is used in the Spanning Tree Protocol.
    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1dBaseBridgeAddress interface{}

    // The number of ports controlled by this bridging entity. The type is
    // interface{} with range: -2147483648..2147483647. Units are ports.
    Dot1dBaseNumPorts interface{}

    // Indicates what type of bridging this bridge can perform.  If a bridge is
    // actually performing a certain type of bridging, this will be indicated by
    // entries in the port table for the given type. The type is Dot1dBaseType.
    Dot1dBaseType interface{}
}

func (dot1dBase *BRIDGEMIB_Dot1dBase) GetEntityData() *types.CommonEntityData {
    dot1dBase.EntityData.YFilter = dot1dBase.YFilter
    dot1dBase.EntityData.YangName = "dot1dBase"
    dot1dBase.EntityData.BundleName = "cisco_ios_xe"
    dot1dBase.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dBase.EntityData.SegmentPath = "dot1dBase"
    dot1dBase.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dBase.EntityData.SegmentPath
    dot1dBase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dBase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dBase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dBase.EntityData.Children = types.NewOrderedMap()
    dot1dBase.EntityData.Leafs = types.NewOrderedMap()
    dot1dBase.EntityData.Leafs.Append("dot1dBaseBridgeAddress", types.YLeaf{"Dot1dBaseBridgeAddress", dot1dBase.Dot1dBaseBridgeAddress})
    dot1dBase.EntityData.Leafs.Append("dot1dBaseNumPorts", types.YLeaf{"Dot1dBaseNumPorts", dot1dBase.Dot1dBaseNumPorts})
    dot1dBase.EntityData.Leafs.Append("dot1dBaseType", types.YLeaf{"Dot1dBaseType", dot1dBase.Dot1dBaseType})

    dot1dBase.EntityData.YListKeys = []string {}

    return &(dot1dBase.EntityData)
}

// BRIDGEMIB_Dot1dBase_Dot1dBaseType represents entries in the port table for the given type.
type BRIDGEMIB_Dot1dBase_Dot1dBaseType string

const (
    BRIDGEMIB_Dot1dBase_Dot1dBaseType_unknown BRIDGEMIB_Dot1dBase_Dot1dBaseType = "unknown"

    BRIDGEMIB_Dot1dBase_Dot1dBaseType_transparent_only BRIDGEMIB_Dot1dBase_Dot1dBaseType = "transparent-only"

    BRIDGEMIB_Dot1dBase_Dot1dBaseType_sourceroute_only BRIDGEMIB_Dot1dBase_Dot1dBaseType = "sourceroute-only"

    BRIDGEMIB_Dot1dBase_Dot1dBaseType_srt BRIDGEMIB_Dot1dBase_Dot1dBaseType = "srt"
)

// BRIDGEMIB_Dot1dStp
type BRIDGEMIB_Dot1dStp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of what version of the Spanning Tree Protocol is being run. 
    // The value 'decLb100(2)' indicates the DEC LANbridge 100 Spanning Tree
    // protocol. IEEE 802.1D implementations will return 'ieee8021d(3)'. If future
    // versions of the IEEE Spanning Tree Protocol that are incompatible with the
    // current version are released a new value will be defined. The type is
    // Dot1dStpProtocolSpecification.
    Dot1dStpProtocolSpecification interface{}

    // The value of the write-able portion of the Bridge ID (i.e., the first two
    // octets of the (8 octet long) Bridge ID).  The other (last) 6 octets of the
    // Bridge ID are given by the value of dot1dBaseBridgeAddress. On bridges
    // supporting IEEE 802.1t or IEEE 802.1w, permissible values are 0-61440, in
    // steps of 4096. The type is interface{} with range: 0..65535.
    Dot1dStpPriority interface{}

    // The time (in hundredths of a second) since the last time a topology change
    // was detected by the bridge entity. For RSTP, this reports the time since
    // the tcWhile timer for any port on this Bridge was nonzero. The type is
    // interface{} with range: 0..4294967295. Units are centi-seconds.
    Dot1dStpTimeSinceTopologyChange interface{}

    // The total number of topology changes detected by this bridge since the
    // management entity was last reset or initialized. The type is interface{}
    // with range: 0..4294967295.
    Dot1dStpTopChanges interface{}

    // The bridge identifier of the root of the spanning tree, as determined by
    // the Spanning Tree Protocol, as executed by this node.  This value is used
    // as the Root Identifier parameter in all Configuration Bridge PDUs
    // originated by this node. The type is string with length: 8.
    Dot1dStpDesignatedRoot interface{}

    // The cost of the path to the root as seen from this bridge. The type is
    // interface{} with range: -2147483648..2147483647.
    Dot1dStpRootCost interface{}

    // The port number of the port that offers the lowest cost path from this
    // bridge to the root bridge. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot1dStpRootPort interface{}

    // The maximum age of Spanning Tree Protocol information learned from the
    // network on any port before it is discarded, in units of hundredths of a
    // second.  This is the actual value that this bridge is currently using. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // centi-seconds.
    Dot1dStpMaxAge interface{}

    // The amount of time between the transmission of Configuration bridge PDUs by
    // this node on any port when it is the root of the spanning tree, or trying
    // to become so, in units of hundredths of a second.  This is the actual value
    // that this bridge is currently using. The type is interface{} with range:
    // -2147483648..2147483647. Units are centi-seconds.
    Dot1dStpHelloTime interface{}

    // This time value determines the interval length during which no more than
    // two Configuration bridge PDUs shall be transmitted by this node, in units
    // of hundredths of a second. The type is interface{} with range:
    // -2147483648..2147483647. Units are centi-seconds.
    Dot1dStpHoldTime interface{}

    // This time value, measured in units of hundredths of a second, controls how
    // fast a port changes its spanning state when moving towards the Forwarding
    // state.  The value determines how long the port stays in each of the
    // Listening and Learning states, which precede the Forwarding state.  This
    // value is also used when a topology change has been detected and is
    // underway, to age all dynamic entries in the Forwarding Database. [Note that
    // this value is the one that this bridge is currently using, in contrast to
    // dot1dStpBridgeForwardDelay, which is the value that this bridge and all
    // others would start using if/when this bridge were to become the root.]. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // centi-seconds.
    Dot1dStpForwardDelay interface{}

    // The value that all bridges use for MaxAge when this bridge is acting as the
    // root.  Note that 802.1D-1998 specifies that the range for this parameter is
    // related to the value of dot1dStpBridgeHelloTime.  The granularity of this
    // timer is specified by 802.1D-1998 to be 1 second.  An agent may return a
    // badValue error if a set is attempted to a value that is not a whole number
    // of seconds. The type is interface{} with range: 600..4000. Units are
    // centi-seconds.
    Dot1dStpBridgeMaxAge interface{}

    // The value that all bridges use for HelloTime when this bridge is acting as
    // the root.  The granularity of this timer is specified by 802.1D-1998 to be
    // 1 second.  An agent may return a badValue error if a set is attempted  to a
    // value that is not a whole number of seconds. The type is interface{} with
    // range: 100..1000. Units are centi-seconds.
    Dot1dStpBridgeHelloTime interface{}

    // The value that all bridges use for ForwardDelay when this bridge is acting
    // as the root.  Note that 802.1D-1998 specifies that the range for this
    // parameter is related to the value of dot1dStpBridgeMaxAge.  The granularity
    // of this timer is specified by 802.1D-1998 to be 1 second.  An agent may
    // return a badValue error if a set is attempted to a value that is not a
    // whole number of seconds. The type is interface{} with range: 400..3000.
    // Units are centi-seconds.
    Dot1dStpBridgeForwardDelay interface{}
}

func (dot1dStp *BRIDGEMIB_Dot1dStp) GetEntityData() *types.CommonEntityData {
    dot1dStp.EntityData.YFilter = dot1dStp.YFilter
    dot1dStp.EntityData.YangName = "dot1dStp"
    dot1dStp.EntityData.BundleName = "cisco_ios_xe"
    dot1dStp.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dStp.EntityData.SegmentPath = "dot1dStp"
    dot1dStp.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dStp.EntityData.SegmentPath
    dot1dStp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dStp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dStp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dStp.EntityData.Children = types.NewOrderedMap()
    dot1dStp.EntityData.Leafs = types.NewOrderedMap()
    dot1dStp.EntityData.Leafs.Append("dot1dStpProtocolSpecification", types.YLeaf{"Dot1dStpProtocolSpecification", dot1dStp.Dot1dStpProtocolSpecification})
    dot1dStp.EntityData.Leafs.Append("dot1dStpPriority", types.YLeaf{"Dot1dStpPriority", dot1dStp.Dot1dStpPriority})
    dot1dStp.EntityData.Leafs.Append("dot1dStpTimeSinceTopologyChange", types.YLeaf{"Dot1dStpTimeSinceTopologyChange", dot1dStp.Dot1dStpTimeSinceTopologyChange})
    dot1dStp.EntityData.Leafs.Append("dot1dStpTopChanges", types.YLeaf{"Dot1dStpTopChanges", dot1dStp.Dot1dStpTopChanges})
    dot1dStp.EntityData.Leafs.Append("dot1dStpDesignatedRoot", types.YLeaf{"Dot1dStpDesignatedRoot", dot1dStp.Dot1dStpDesignatedRoot})
    dot1dStp.EntityData.Leafs.Append("dot1dStpRootCost", types.YLeaf{"Dot1dStpRootCost", dot1dStp.Dot1dStpRootCost})
    dot1dStp.EntityData.Leafs.Append("dot1dStpRootPort", types.YLeaf{"Dot1dStpRootPort", dot1dStp.Dot1dStpRootPort})
    dot1dStp.EntityData.Leafs.Append("dot1dStpMaxAge", types.YLeaf{"Dot1dStpMaxAge", dot1dStp.Dot1dStpMaxAge})
    dot1dStp.EntityData.Leafs.Append("dot1dStpHelloTime", types.YLeaf{"Dot1dStpHelloTime", dot1dStp.Dot1dStpHelloTime})
    dot1dStp.EntityData.Leafs.Append("dot1dStpHoldTime", types.YLeaf{"Dot1dStpHoldTime", dot1dStp.Dot1dStpHoldTime})
    dot1dStp.EntityData.Leafs.Append("dot1dStpForwardDelay", types.YLeaf{"Dot1dStpForwardDelay", dot1dStp.Dot1dStpForwardDelay})
    dot1dStp.EntityData.Leafs.Append("dot1dStpBridgeMaxAge", types.YLeaf{"Dot1dStpBridgeMaxAge", dot1dStp.Dot1dStpBridgeMaxAge})
    dot1dStp.EntityData.Leafs.Append("dot1dStpBridgeHelloTime", types.YLeaf{"Dot1dStpBridgeHelloTime", dot1dStp.Dot1dStpBridgeHelloTime})
    dot1dStp.EntityData.Leafs.Append("dot1dStpBridgeForwardDelay", types.YLeaf{"Dot1dStpBridgeForwardDelay", dot1dStp.Dot1dStpBridgeForwardDelay})

    dot1dStp.EntityData.YListKeys = []string {}

    return &(dot1dStp.EntityData)
}

// BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification represents are released a new value will be defined.
type BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification string

const (
    BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification_unknown BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification = "unknown"

    BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification_decLb100 BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification = "decLb100"

    BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification_ieee8021d BRIDGEMIB_Dot1dStp_Dot1dStpProtocolSpecification = "ieee8021d"
)

// BRIDGEMIB_Dot1dTp
type BRIDGEMIB_Dot1dTp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of Forwarding Database entries that have been or would
    // have been learned, but have been discarded due to a lack of storage space
    // in the Forwarding Database.  If this counter is increasing, it indicates
    // that the Forwarding Database is regularly becoming full (a condition that
    // has unpleasant performance effects on the subnetwork).  If this counter has
    // a significant value but is not presently increasing, it indicates that the
    // problem has been occurring but is not persistent. The type is interface{}
    // with range: 0..4294967295.
    Dot1dTpLearnedEntryDiscards interface{}

    // The timeout period in seconds for aging out dynamically-learned forwarding
    // information. 802.1D-1998 recommends a default of 300 seconds. The type is
    // interface{} with range: 10..1000000. Units are seconds.
    Dot1dTpAgingTime interface{}
}

func (dot1dTp *BRIDGEMIB_Dot1dTp) GetEntityData() *types.CommonEntityData {
    dot1dTp.EntityData.YFilter = dot1dTp.YFilter
    dot1dTp.EntityData.YangName = "dot1dTp"
    dot1dTp.EntityData.BundleName = "cisco_ios_xe"
    dot1dTp.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dTp.EntityData.SegmentPath = "dot1dTp"
    dot1dTp.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dTp.EntityData.SegmentPath
    dot1dTp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTp.EntityData.Children = types.NewOrderedMap()
    dot1dTp.EntityData.Leafs = types.NewOrderedMap()
    dot1dTp.EntityData.Leafs.Append("dot1dTpLearnedEntryDiscards", types.YLeaf{"Dot1dTpLearnedEntryDiscards", dot1dTp.Dot1dTpLearnedEntryDiscards})
    dot1dTp.EntityData.Leafs.Append("dot1dTpAgingTime", types.YLeaf{"Dot1dTpAgingTime", dot1dTp.Dot1dTpAgingTime})

    dot1dTp.EntityData.YListKeys = []string {}

    return &(dot1dTp.EntityData)
}

// BRIDGEMIB_Dot1dBasePortTable
// A table that contains generic information about every
// port that is associated with this bridge.  Transparent,
// source-route, and srt ports are included.
type BRIDGEMIB_Dot1dBasePortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of information for each port of the bridge. The type is slice of
    // BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry.
    Dot1dBasePortEntry []*BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry
}

func (dot1dBasePortTable *BRIDGEMIB_Dot1dBasePortTable) GetEntityData() *types.CommonEntityData {
    dot1dBasePortTable.EntityData.YFilter = dot1dBasePortTable.YFilter
    dot1dBasePortTable.EntityData.YangName = "dot1dBasePortTable"
    dot1dBasePortTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dBasePortTable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dBasePortTable.EntityData.SegmentPath = "dot1dBasePortTable"
    dot1dBasePortTable.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dBasePortTable.EntityData.SegmentPath
    dot1dBasePortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dBasePortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dBasePortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dBasePortTable.EntityData.Children = types.NewOrderedMap()
    dot1dBasePortTable.EntityData.Children.Append("dot1dBasePortEntry", types.YChild{"Dot1dBasePortEntry", nil})
    for i := range dot1dBasePortTable.Dot1dBasePortEntry {
        dot1dBasePortTable.EntityData.Children.Append(types.GetSegmentPath(dot1dBasePortTable.Dot1dBasePortEntry[i]), types.YChild{"Dot1dBasePortEntry", dot1dBasePortTable.Dot1dBasePortEntry[i]})
    }
    dot1dBasePortTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dBasePortTable.EntityData.YListKeys = []string {}

    return &(dot1dBasePortTable.EntityData)
}

// BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry
// A list of information for each port of the bridge.
type BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The port number of the port for which this entry
    // contains bridge management information. The type is interface{} with range:
    // 1..65535.
    Dot1dBasePort interface{}

    // The value of the instance of the ifIndex object, defined in IF-MIB, for the
    // interface corresponding to this port. The type is interface{} with range:
    // 1..2147483647.
    Dot1dBasePortIfIndex interface{}

    // For a port that (potentially) has the same value of dot1dBasePortIfIndex as
    // another port on the same bridge. This object contains the name of an object
    // instance unique to this port.  For example, in the case where multiple
    // ports correspond one-to-one with multiple X.25 virtual circuits, this value
    // might identify an (e.g., the first) object instance associated with the
    // X.25 virtual circuit corresponding to this port.  For a port which has a
    // unique value of dot1dBasePortIfIndex, this object can have the value { 0 0
    // }. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Dot1dBasePortCircuit interface{}

    // The number of frames discarded by this port due to excessive transit delay
    // through the bridge.  It is incremented by both transparent and source route
    // bridges. The type is interface{} with range: 0..4294967295.
    Dot1dBasePortDelayExceededDiscards interface{}

    // The number of frames discarded by this port due to an excessive size.  It
    // is incremented by both transparent and source route bridges. The type is
    // interface{} with range: 0..4294967295.
    Dot1dBasePortMtuExceededDiscards interface{}

    // Indicates the parts of IEEE 802.1D and 802.1Q that are optional on a
    // per-port basis, that are implemented by this device, and that are
    // manageable through this MIB.  dot1qDot1qTagging(0), -- supports 802.1Q VLAN
    // tagging of                       -- frames and GVRP.
    // dot1qConfigurableAcceptableFrameTypes(1),                       -- allows
    // modified values of                       -- dot1qPortAcceptableFrameTypes.
    // dot1qIngressFiltering(2)                       -- supports the discarding
    // of any                       -- frame received on a Port whose             
    // -- VLAN classification does not                       -- include that Port
    // in its Member                       -- set. The type is map[string]bool.
    Dot1dPortCapabilities interface{}

    // The default ingress User Priority for this port.  This only has effect on
    // media, such as Ethernet, that do not support native User Priority.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is interface{} with range: 0..7.
    Dot1dPortDefaultUserPriority interface{}

    // The number of egress traffic classes supported on this port.  This object
    // may optionally be read-only.  The value of this object MUST be retained
    // across reinitializations of the management system. The type is interface{}
    // with range: 1..8.
    Dot1dPortNumTrafficClasses interface{}

    // The GARP Join time, in centiseconds.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // interface{} with range: 0..2147483647.
    Dot1dPortGarpJoinTime interface{}

    // The GARP Leave time, in centiseconds.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // interface{} with range: 0..2147483647.
    Dot1dPortGarpLeaveTime interface{}

    // The GARP LeaveAll time, in centiseconds.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // interface{} with range: 0..2147483647.
    Dot1dPortGarpLeaveAllTime interface{}

    // The administrative state of GMRP operation on this port.  The value
    // enabled(1) indicates that GMRP is enabled on this port in all VLANs as long
    // as dot1dGmrpStatus is also enabled(1). A value of disabled(2) indicates
    // that GMRP is disabled on this port in all VLANs: any GMRP packets received
    // will be silently discarded, and no GMRP registrations will be propagated
    // from other ports.  Setting this to a value of enabled(1) will be stored by
    // the agent but will only take effect on the GMRP protocol operation if
    // dot1dGmrpStatus also indicates the value enabled(1).  This object affects
    // all GMRP Applicant and Registrar state machines on this port.  A transition
    // from disabled(2) to enabled(1) will cause a reset of all GMRP state
    // machines on this port.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is EnabledStatus.
    Dot1dPortGmrpStatus interface{}

    // The total number of failed GMRP registrations, for any reason, in all
    // VLANs, on this port. The type is interface{} with range: 0..4294967295.
    Dot1dPortGmrpFailedRegistrations interface{}

    // The Source MAC Address of the last GMRP message received on this port. The
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1dPortGmrpLastPduOrigin interface{}

    // The state of Restricted Group Registration on this port. If the value of
    // this control is true(1), then creation of a new dynamic entry is permitted
    // only if there is a Static Filtering Entry for the VLAN concerned, in which
    // the Registrar Administrative Control value is Normal Registration.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is bool.
    Dot1dPortRestrictedGroupRegistration interface{}

    // The PVID, the VLAN-ID assigned to untagged frames or Priority-Tagged frames
    // received on this port.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is interface{} with
    // range: 0..4294967295.
    Dot1qPvid interface{}

    // When this is admitOnlyVlanTagged(2), the device will discard untagged
    // frames or Priority-Tagged frames received on this port.  When admitAll(1),
    // untagged frames or Priority-Tagged frames received on this port will be
    // accepted and assigned to a VID based on the PVID and VID Set for this port.
    // This control does not affect VLAN-independent Bridge Protocol Data Unit
    // (BPDU) frames, such as GVRP and Spanning Tree Protocol (STP).  It does
    // affect VLAN- dependent BPDU frames, such as GMRP.  The value of this object
    // MUST be retained across reinitializations of the management system. The
    // type is Dot1qPortAcceptableFrameTypes.
    Dot1qPortAcceptableFrameTypes interface{}

    // When this is true(1), the device will discard incoming frames for VLANs
    // that do not include this Port in its  Member set.  When false(2), the port
    // will accept all incoming frames.  This control does not affect
    // VLAN-independent BPDU frames, such as GVRP and STP.  It does affect VLAN-
    // dependent BPDU frames, such as GMRP.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // bool.
    Dot1qPortIngressFiltering interface{}

    // The state of GVRP operation on this port.  The value enabled(1) indicates
    // that GVRP is enabled on this port, as long as dot1qGvrpStatus is also
    // enabled for this device.  When disabled(2) but dot1qGvrpStatus is still
    // enabled for the device, GVRP is disabled on this port: any GVRP packets
    // received will be silently discarded, and no GVRP registrations will be
    // propagated from other ports.  This object affects all GVRP Applicant and
    // Registrar state machines on this port.  A transition from disabled(2) to
    // enabled(1) will cause a reset of all GVRP state machines on this port.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is EnabledStatus.
    Dot1qPortGvrpStatus interface{}

    // The total number of failed GVRP registrations, for any reason, on this
    // port. The type is interface{} with range: 0..4294967295.
    Dot1qPortGvrpFailedRegistrations interface{}

    // The Source MAC Address of the last GVRP message received on this port. The
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1qPortGvrpLastPduOrigin interface{}

    // The state of Restricted VLAN Registration on this port. If the value of
    // this control is true(1), then creation of a new dynamic VLAN entry is
    // permitted only if there is a Static VLAN Registration Entry for the VLAN
    // concerned, in which the Registrar Administrative Control value for this
    // port is Normal Registration.  The value of this object MUST be retained
    // across reinitializations of the management system. The type is bool.
    Dot1qPortRestrictedVlanRegistration interface{}
}

func (dot1dBasePortEntry *BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry) GetEntityData() *types.CommonEntityData {
    dot1dBasePortEntry.EntityData.YFilter = dot1dBasePortEntry.YFilter
    dot1dBasePortEntry.EntityData.YangName = "dot1dBasePortEntry"
    dot1dBasePortEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dBasePortEntry.EntityData.ParentYangName = "dot1dBasePortTable"
    dot1dBasePortEntry.EntityData.SegmentPath = "dot1dBasePortEntry" + types.AddKeyToken(dot1dBasePortEntry.Dot1dBasePort, "dot1dBasePort")
    dot1dBasePortEntry.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/dot1dBasePortTable/" + dot1dBasePortEntry.EntityData.SegmentPath
    dot1dBasePortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dBasePortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dBasePortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dBasePortEntry.EntityData.Children = types.NewOrderedMap()
    dot1dBasePortEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", dot1dBasePortEntry.Dot1dBasePort})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dBasePortIfIndex", types.YLeaf{"Dot1dBasePortIfIndex", dot1dBasePortEntry.Dot1dBasePortIfIndex})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dBasePortCircuit", types.YLeaf{"Dot1dBasePortCircuit", dot1dBasePortEntry.Dot1dBasePortCircuit})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dBasePortDelayExceededDiscards", types.YLeaf{"Dot1dBasePortDelayExceededDiscards", dot1dBasePortEntry.Dot1dBasePortDelayExceededDiscards})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dBasePortMtuExceededDiscards", types.YLeaf{"Dot1dBasePortMtuExceededDiscards", dot1dBasePortEntry.Dot1dBasePortMtuExceededDiscards})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortCapabilities", types.YLeaf{"Dot1dPortCapabilities", dot1dBasePortEntry.Dot1dPortCapabilities})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortDefaultUserPriority", types.YLeaf{"Dot1dPortDefaultUserPriority", dot1dBasePortEntry.Dot1dPortDefaultUserPriority})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortNumTrafficClasses", types.YLeaf{"Dot1dPortNumTrafficClasses", dot1dBasePortEntry.Dot1dPortNumTrafficClasses})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortGarpJoinTime", types.YLeaf{"Dot1dPortGarpJoinTime", dot1dBasePortEntry.Dot1dPortGarpJoinTime})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortGarpLeaveTime", types.YLeaf{"Dot1dPortGarpLeaveTime", dot1dBasePortEntry.Dot1dPortGarpLeaveTime})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortGarpLeaveAllTime", types.YLeaf{"Dot1dPortGarpLeaveAllTime", dot1dBasePortEntry.Dot1dPortGarpLeaveAllTime})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortGmrpStatus", types.YLeaf{"Dot1dPortGmrpStatus", dot1dBasePortEntry.Dot1dPortGmrpStatus})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortGmrpFailedRegistrations", types.YLeaf{"Dot1dPortGmrpFailedRegistrations", dot1dBasePortEntry.Dot1dPortGmrpFailedRegistrations})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortGmrpLastPduOrigin", types.YLeaf{"Dot1dPortGmrpLastPduOrigin", dot1dBasePortEntry.Dot1dPortGmrpLastPduOrigin})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1dPortRestrictedGroupRegistration", types.YLeaf{"Dot1dPortRestrictedGroupRegistration", dot1dBasePortEntry.Dot1dPortRestrictedGroupRegistration})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1qPvid", types.YLeaf{"Dot1qPvid", dot1dBasePortEntry.Dot1qPvid})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1qPortAcceptableFrameTypes", types.YLeaf{"Dot1qPortAcceptableFrameTypes", dot1dBasePortEntry.Dot1qPortAcceptableFrameTypes})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1qPortIngressFiltering", types.YLeaf{"Dot1qPortIngressFiltering", dot1dBasePortEntry.Dot1qPortIngressFiltering})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1qPortGvrpStatus", types.YLeaf{"Dot1qPortGvrpStatus", dot1dBasePortEntry.Dot1qPortGvrpStatus})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1qPortGvrpFailedRegistrations", types.YLeaf{"Dot1qPortGvrpFailedRegistrations", dot1dBasePortEntry.Dot1qPortGvrpFailedRegistrations})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1qPortGvrpLastPduOrigin", types.YLeaf{"Dot1qPortGvrpLastPduOrigin", dot1dBasePortEntry.Dot1qPortGvrpLastPduOrigin})
    dot1dBasePortEntry.EntityData.Leafs.Append("dot1qPortRestrictedVlanRegistration", types.YLeaf{"Dot1qPortRestrictedVlanRegistration", dot1dBasePortEntry.Dot1qPortRestrictedVlanRegistration})

    dot1dBasePortEntry.EntityData.YListKeys = []string {"Dot1dBasePort"}

    return &(dot1dBasePortEntry.EntityData)
}

// BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1qPortAcceptableFrameTypes represents reinitializations of the management system.
type BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1qPortAcceptableFrameTypes string

const (
    BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1qPortAcceptableFrameTypes_admitAll BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1qPortAcceptableFrameTypes = "admitAll"

    BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1qPortAcceptableFrameTypes_admitOnlyVlanTagged BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1qPortAcceptableFrameTypes = "admitOnlyVlanTagged"
)

// BRIDGEMIB_Dot1dStpPortTable
// A table that contains port-specific information
// for the Spanning Tree Protocol.
type BRIDGEMIB_Dot1dStpPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of information maintained by every port about the Spanning Tree
    // Protocol state for that port. The type is slice of
    // BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry.
    Dot1dStpPortEntry []*BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry
}

func (dot1dStpPortTable *BRIDGEMIB_Dot1dStpPortTable) GetEntityData() *types.CommonEntityData {
    dot1dStpPortTable.EntityData.YFilter = dot1dStpPortTable.YFilter
    dot1dStpPortTable.EntityData.YangName = "dot1dStpPortTable"
    dot1dStpPortTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dStpPortTable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dStpPortTable.EntityData.SegmentPath = "dot1dStpPortTable"
    dot1dStpPortTable.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dStpPortTable.EntityData.SegmentPath
    dot1dStpPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dStpPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dStpPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dStpPortTable.EntityData.Children = types.NewOrderedMap()
    dot1dStpPortTable.EntityData.Children.Append("dot1dStpPortEntry", types.YChild{"Dot1dStpPortEntry", nil})
    for i := range dot1dStpPortTable.Dot1dStpPortEntry {
        dot1dStpPortTable.EntityData.Children.Append(types.GetSegmentPath(dot1dStpPortTable.Dot1dStpPortEntry[i]), types.YChild{"Dot1dStpPortEntry", dot1dStpPortTable.Dot1dStpPortEntry[i]})
    }
    dot1dStpPortTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dStpPortTable.EntityData.YListKeys = []string {}

    return &(dot1dStpPortTable.EntityData)
}

// BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry
// A list of information maintained by every port about
// the Spanning Tree Protocol state for that port.
type BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The port number of the port for which this entry
    // contains Spanning Tree Protocol management information. The type is
    // interface{} with range: 1..65535.
    Dot1dStpPort interface{}

    // The value of the priority field that is contained in the first (in network
    // byte order) octet of the (2 octet long) Port ID.  The other octet of the
    // Port ID is given by the value of dot1dStpPort. On bridges supporting IEEE
    // 802.1t or IEEE 802.1w, permissible values are 0-240, in steps of 16. The
    // type is interface{} with range: 0..255.
    Dot1dStpPortPriority interface{}

    // The port's current state, as defined by application of the Spanning Tree
    // Protocol.  This state controls what action a port takes on reception of a
    // frame.  If the bridge has detected a port that is malfunctioning, it will
    // place that port into the broken(6) state.  For ports that are disabled (see
    // dot1dStpPortEnable), this object will have a value of disabled(1). The type
    // is Dot1dStpPortState.
    Dot1dStpPortState interface{}

    // The enabled/disabled status of the port. The type is Dot1dStpPortEnable.
    Dot1dStpPortEnable interface{}

    // The contribution of this port to the path cost of paths towards the
    // spanning tree root which include this port.  802.1D-1998 recommends that
    // the default value of this parameter be in inverse proportion to  the speed
    // of the attached LAN.  New implementations should support
    // dot1dStpPortPathCost32. If the port path costs exceeds the maximum value of
    // this object then this object should report the maximum value, namely 65535.
    // Applications should try to read the dot1dStpPortPathCost32 object if this
    // object reports the maximum value. The type is interface{} with range:
    // 1..65535.
    Dot1dStpPortPathCost interface{}

    // The unique Bridge Identifier of the Bridge recorded as the Root in the
    // Configuration BPDUs transmitted by the Designated Bridge for the segment to
    // which the port is attached. The type is string with length: 8.
    Dot1dStpPortDesignatedRoot interface{}

    // The path cost of the Designated Port of the segment connected to this port.
    // This value is compared to the Root Path Cost field in received bridge PDUs.
    // The type is interface{} with range: -2147483648..2147483647.
    Dot1dStpPortDesignatedCost interface{}

    // The Bridge Identifier of the bridge that this port considers to be the
    // Designated Bridge for this port's segment. The type is string with length:
    // 8.
    Dot1dStpPortDesignatedBridge interface{}

    // The Port Identifier of the port on the Designated Bridge for this port's
    // segment. The type is string with length: 2.
    Dot1dStpPortDesignatedPort interface{}

    // The number of times this port has transitioned from the Learning state to
    // the Forwarding state. The type is interface{} with range: 0..4294967295.
    Dot1dStpPortForwardTransitions interface{}

    // The contribution of this port to the path cost of paths towards the
    // spanning tree root which include this port.  802.1D-1998 recommends that
    // the default value of this parameter be in inverse proportion to the speed
    // of the attached LAN.  This object replaces dot1dStpPortPathCost to support
    // IEEE 802.1t. The type is interface{} with range: 1..200000000.
    Dot1dStpPortPathCost32 interface{}

    // The contribution of this port to the path cost (in 32 bits value) of paths
    // towards the spanning tree root which include this port.  This object is
    // used to configure the spanning tree port path cost in 32 bits value range
    // when the stpxSpanningTreePathCostOperMode is long(2).  If the
    // stpxSpanningTreePathCostOperMode is short(1), this  MIB object is not
    // instantiated. The type is interface{} with range: 0..4294967295.
    StpxLongStpPortPathCost interface{}
}

func (dot1dStpPortEntry *BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry) GetEntityData() *types.CommonEntityData {
    dot1dStpPortEntry.EntityData.YFilter = dot1dStpPortEntry.YFilter
    dot1dStpPortEntry.EntityData.YangName = "dot1dStpPortEntry"
    dot1dStpPortEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dStpPortEntry.EntityData.ParentYangName = "dot1dStpPortTable"
    dot1dStpPortEntry.EntityData.SegmentPath = "dot1dStpPortEntry" + types.AddKeyToken(dot1dStpPortEntry.Dot1dStpPort, "dot1dStpPort")
    dot1dStpPortEntry.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/dot1dStpPortTable/" + dot1dStpPortEntry.EntityData.SegmentPath
    dot1dStpPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dStpPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dStpPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dStpPortEntry.EntityData.Children = types.NewOrderedMap()
    dot1dStpPortEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPort", types.YLeaf{"Dot1dStpPort", dot1dStpPortEntry.Dot1dStpPort})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortPriority", types.YLeaf{"Dot1dStpPortPriority", dot1dStpPortEntry.Dot1dStpPortPriority})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortState", types.YLeaf{"Dot1dStpPortState", dot1dStpPortEntry.Dot1dStpPortState})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortEnable", types.YLeaf{"Dot1dStpPortEnable", dot1dStpPortEntry.Dot1dStpPortEnable})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortPathCost", types.YLeaf{"Dot1dStpPortPathCost", dot1dStpPortEntry.Dot1dStpPortPathCost})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortDesignatedRoot", types.YLeaf{"Dot1dStpPortDesignatedRoot", dot1dStpPortEntry.Dot1dStpPortDesignatedRoot})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortDesignatedCost", types.YLeaf{"Dot1dStpPortDesignatedCost", dot1dStpPortEntry.Dot1dStpPortDesignatedCost})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortDesignatedBridge", types.YLeaf{"Dot1dStpPortDesignatedBridge", dot1dStpPortEntry.Dot1dStpPortDesignatedBridge})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortDesignatedPort", types.YLeaf{"Dot1dStpPortDesignatedPort", dot1dStpPortEntry.Dot1dStpPortDesignatedPort})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortForwardTransitions", types.YLeaf{"Dot1dStpPortForwardTransitions", dot1dStpPortEntry.Dot1dStpPortForwardTransitions})
    dot1dStpPortEntry.EntityData.Leafs.Append("dot1dStpPortPathCost32", types.YLeaf{"Dot1dStpPortPathCost32", dot1dStpPortEntry.Dot1dStpPortPathCost32})
    dot1dStpPortEntry.EntityData.Leafs.Append("stpxLongStpPortPathCost", types.YLeaf{"StpxLongStpPortPathCost", dot1dStpPortEntry.StpxLongStpPortPathCost})

    dot1dStpPortEntry.EntityData.YListKeys = []string {"Dot1dStpPort"}

    return &(dot1dStpPortEntry.EntityData)
}

// BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortEnable represents The enabled/disabled status of the port.
type BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortEnable string

const (
    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortEnable_enabled BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortEnable = "enabled"

    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortEnable_disabled BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortEnable = "disabled"
)

// BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState represents object will have a value of disabled(1).
type BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState string

const (
    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState_disabled BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState = "disabled"

    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState_blocking BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState = "blocking"

    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState_listening BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState = "listening"

    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState_learning BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState = "learning"

    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState_forwarding BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState = "forwarding"

    BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState_broken BRIDGEMIB_Dot1dStpPortTable_Dot1dStpPortEntry_Dot1dStpPortState = "broken"
)

// BRIDGEMIB_Dot1dTpFdbTable
// A table that contains information about unicast
// entries for which the bridge has forwarding and/or
// filtering information.  This information is used
// by the transparent bridging function in
// determining how to propagate a received frame.
type BRIDGEMIB_Dot1dTpFdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a specific unicast MAC address for which the bridge has
    // some forwarding and/or filtering information. The type is slice of
    // BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry.
    Dot1dTpFdbEntry []*BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry
}

func (dot1dTpFdbTable *BRIDGEMIB_Dot1dTpFdbTable) GetEntityData() *types.CommonEntityData {
    dot1dTpFdbTable.EntityData.YFilter = dot1dTpFdbTable.YFilter
    dot1dTpFdbTable.EntityData.YangName = "dot1dTpFdbTable"
    dot1dTpFdbTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpFdbTable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dTpFdbTable.EntityData.SegmentPath = "dot1dTpFdbTable"
    dot1dTpFdbTable.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dTpFdbTable.EntityData.SegmentPath
    dot1dTpFdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpFdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpFdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpFdbTable.EntityData.Children = types.NewOrderedMap()
    dot1dTpFdbTable.EntityData.Children.Append("dot1dTpFdbEntry", types.YChild{"Dot1dTpFdbEntry", nil})
    for i := range dot1dTpFdbTable.Dot1dTpFdbEntry {
        dot1dTpFdbTable.EntityData.Children.Append(types.GetSegmentPath(dot1dTpFdbTable.Dot1dTpFdbEntry[i]), types.YChild{"Dot1dTpFdbEntry", dot1dTpFdbTable.Dot1dTpFdbEntry[i]})
    }
    dot1dTpFdbTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dTpFdbTable.EntityData.YListKeys = []string {}

    return &(dot1dTpFdbTable.EntityData)
}

// BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry
// Information about a specific unicast MAC address
// for which the bridge has some forwarding and/or
// filtering information.
type BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unicast MAC address for which the bridge has
    // forwarding and/or filtering information. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1dTpFdbAddress interface{}

    // Either the value '0', or the port number of the port on which a frame
    // having a source address equal to the value of the corresponding instance of
    // dot1dTpFdbAddress has been seen.  A value of '0' indicates that the port
    // number has not been learned, but that the bridge does have some
    // forwarding/filtering information about this address (e.g., in the
    // dot1dStaticTable).  Implementors are encouraged to assign the port value to
    // this object whenever it is learned, even for addresses for which the
    // corresponding value of dot1dTpFdbStatus is not learned(3). The type is
    // interface{} with range: -2147483648..2147483647.
    Dot1dTpFdbPort interface{}

    // The status of this entry.  The meanings of the values are:     other(1) -
    // none of the following.  This would         include the case where some
    // other MIB object         (not the corresponding instance of        
    // dot1dTpFdbPort, nor an entry in the         dot1dStaticTable) is being used
    // to determine if         and how frames addressed to the value of the       
    // corresponding instance of dot1dTpFdbAddress are         being forwarded.   
    // invalid(2) - this entry is no longer valid (e.g.,         it was learned
    // but has since aged out), but has         not yet been flushed from the
    // table.     learned(3) - the value of the corresponding instance         of
    // dot1dTpFdbPort was learned, and is being         used.     self(4) - the
    // value of the corresponding instance of         dot1dTpFdbAddress represents
    // one of the bridge's         addresses.  The corresponding instance of      
    // dot1dTpFdbPort indicates which of the bridge's         ports has this
    // address.     mgmt(5) - the value of the corresponding instance of        
    // dot1dTpFdbAddress is also the value of an         existing instance of
    // dot1dStaticAddress. The type is Dot1dTpFdbStatus.
    Dot1dTpFdbStatus interface{}
}

func (dot1dTpFdbEntry *BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry) GetEntityData() *types.CommonEntityData {
    dot1dTpFdbEntry.EntityData.YFilter = dot1dTpFdbEntry.YFilter
    dot1dTpFdbEntry.EntityData.YangName = "dot1dTpFdbEntry"
    dot1dTpFdbEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpFdbEntry.EntityData.ParentYangName = "dot1dTpFdbTable"
    dot1dTpFdbEntry.EntityData.SegmentPath = "dot1dTpFdbEntry" + types.AddKeyToken(dot1dTpFdbEntry.Dot1dTpFdbAddress, "dot1dTpFdbAddress")
    dot1dTpFdbEntry.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/dot1dTpFdbTable/" + dot1dTpFdbEntry.EntityData.SegmentPath
    dot1dTpFdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpFdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpFdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpFdbEntry.EntityData.Children = types.NewOrderedMap()
    dot1dTpFdbEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dTpFdbEntry.EntityData.Leafs.Append("dot1dTpFdbAddress", types.YLeaf{"Dot1dTpFdbAddress", dot1dTpFdbEntry.Dot1dTpFdbAddress})
    dot1dTpFdbEntry.EntityData.Leafs.Append("dot1dTpFdbPort", types.YLeaf{"Dot1dTpFdbPort", dot1dTpFdbEntry.Dot1dTpFdbPort})
    dot1dTpFdbEntry.EntityData.Leafs.Append("dot1dTpFdbStatus", types.YLeaf{"Dot1dTpFdbStatus", dot1dTpFdbEntry.Dot1dTpFdbStatus})

    dot1dTpFdbEntry.EntityData.YListKeys = []string {"Dot1dTpFdbAddress"}

    return &(dot1dTpFdbEntry.EntityData)
}

// BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus represents         existing instance of dot1dStaticAddress.
type BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus string

const (
    BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus_other BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus = "other"

    BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus_invalid BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus = "invalid"

    BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus_learned BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus = "learned"

    BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus_self BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus = "self"

    BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus_mgmt BRIDGEMIB_Dot1dTpFdbTable_Dot1dTpFdbEntry_Dot1dTpFdbStatus = "mgmt"
)

// BRIDGEMIB_Dot1dTpPortTable
// A table that contains information about every port that
// is associated with this transparent bridge.
type BRIDGEMIB_Dot1dTpPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of information for each port of a transparent bridge. The type is
    // slice of BRIDGEMIB_Dot1dTpPortTable_Dot1dTpPortEntry.
    Dot1dTpPortEntry []*BRIDGEMIB_Dot1dTpPortTable_Dot1dTpPortEntry
}

func (dot1dTpPortTable *BRIDGEMIB_Dot1dTpPortTable) GetEntityData() *types.CommonEntityData {
    dot1dTpPortTable.EntityData.YFilter = dot1dTpPortTable.YFilter
    dot1dTpPortTable.EntityData.YangName = "dot1dTpPortTable"
    dot1dTpPortTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpPortTable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dTpPortTable.EntityData.SegmentPath = "dot1dTpPortTable"
    dot1dTpPortTable.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dTpPortTable.EntityData.SegmentPath
    dot1dTpPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpPortTable.EntityData.Children = types.NewOrderedMap()
    dot1dTpPortTable.EntityData.Children.Append("dot1dTpPortEntry", types.YChild{"Dot1dTpPortEntry", nil})
    for i := range dot1dTpPortTable.Dot1dTpPortEntry {
        dot1dTpPortTable.EntityData.Children.Append(types.GetSegmentPath(dot1dTpPortTable.Dot1dTpPortEntry[i]), types.YChild{"Dot1dTpPortEntry", dot1dTpPortTable.Dot1dTpPortEntry[i]})
    }
    dot1dTpPortTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dTpPortTable.EntityData.YListKeys = []string {}

    return &(dot1dTpPortTable.EntityData)
}

// BRIDGEMIB_Dot1dTpPortTable_Dot1dTpPortEntry
// A list of information for each port of a transparent
// bridge.
type BRIDGEMIB_Dot1dTpPortTable_Dot1dTpPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The port number of the port for which this entry
    // contains Transparent bridging management information. The type is
    // interface{} with range: 1..65535.
    Dot1dTpPort interface{}

    // The maximum size of the INFO (non-MAC) field that  this port will receive
    // or transmit. The type is interface{} with range: -2147483648..2147483647.
    // Units are bytes.
    Dot1dTpPortMaxInfo interface{}

    // The number of frames that have been received by this port from its segment.
    // Note that a frame received on the interface corresponding to this port is
    // only counted by this object if and only if it is for a protocol being
    // processed by the local bridging function, including bridge management
    // frames. The type is interface{} with range: 0..4294967295. Units are
    // frames.
    Dot1dTpPortInFrames interface{}

    // The number of frames that have been transmitted by this port to its
    // segment.  Note that a frame transmitted on the interface corresponding to
    // this port is only counted by this object if and only if it is for a
    // protocol being processed by the local bridging function, including bridge
    // management frames. The type is interface{} with range: 0..4294967295. Units
    // are frames.
    Dot1dTpPortOutFrames interface{}

    // Count of received valid frames that were discarded (i.e., filtered) by the
    // Forwarding Process. The type is interface{} with range: 0..4294967295.
    // Units are frames.
    Dot1dTpPortInDiscards interface{}
}

func (dot1dTpPortEntry *BRIDGEMIB_Dot1dTpPortTable_Dot1dTpPortEntry) GetEntityData() *types.CommonEntityData {
    dot1dTpPortEntry.EntityData.YFilter = dot1dTpPortEntry.YFilter
    dot1dTpPortEntry.EntityData.YangName = "dot1dTpPortEntry"
    dot1dTpPortEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpPortEntry.EntityData.ParentYangName = "dot1dTpPortTable"
    dot1dTpPortEntry.EntityData.SegmentPath = "dot1dTpPortEntry" + types.AddKeyToken(dot1dTpPortEntry.Dot1dTpPort, "dot1dTpPort")
    dot1dTpPortEntry.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/dot1dTpPortTable/" + dot1dTpPortEntry.EntityData.SegmentPath
    dot1dTpPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpPortEntry.EntityData.Children = types.NewOrderedMap()
    dot1dTpPortEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dTpPortEntry.EntityData.Leafs.Append("dot1dTpPort", types.YLeaf{"Dot1dTpPort", dot1dTpPortEntry.Dot1dTpPort})
    dot1dTpPortEntry.EntityData.Leafs.Append("dot1dTpPortMaxInfo", types.YLeaf{"Dot1dTpPortMaxInfo", dot1dTpPortEntry.Dot1dTpPortMaxInfo})
    dot1dTpPortEntry.EntityData.Leafs.Append("dot1dTpPortInFrames", types.YLeaf{"Dot1dTpPortInFrames", dot1dTpPortEntry.Dot1dTpPortInFrames})
    dot1dTpPortEntry.EntityData.Leafs.Append("dot1dTpPortOutFrames", types.YLeaf{"Dot1dTpPortOutFrames", dot1dTpPortEntry.Dot1dTpPortOutFrames})
    dot1dTpPortEntry.EntityData.Leafs.Append("dot1dTpPortInDiscards", types.YLeaf{"Dot1dTpPortInDiscards", dot1dTpPortEntry.Dot1dTpPortInDiscards})

    dot1dTpPortEntry.EntityData.YListKeys = []string {"Dot1dTpPort"}

    return &(dot1dTpPortEntry.EntityData)
}

// BRIDGEMIB_Dot1dStaticTable
// A table containing filtering information configured
// into the bridge by (local or network) management
// specifying the set of ports to which frames received
// from specific ports and containing specific destination
// addresses are allowed to be forwarded.  The value of
// zero in this table, as the port number from which frames
// with a specific destination address are received, is
// used to specify all ports for which there is no specific
// entry in this table for that particular destination
// address.  Entries are valid for unicast and for
// group/broadcast addresses.
type BRIDGEMIB_Dot1dStaticTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the bridge by (local or network)
    // management specifying the set of ports to which frames received from a
    // specific port and containing a specific destination address are allowed to
    // be forwarded. The type is slice of
    // BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry.
    Dot1dStaticEntry []*BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry
}

func (dot1dStaticTable *BRIDGEMIB_Dot1dStaticTable) GetEntityData() *types.CommonEntityData {
    dot1dStaticTable.EntityData.YFilter = dot1dStaticTable.YFilter
    dot1dStaticTable.EntityData.YangName = "dot1dStaticTable"
    dot1dStaticTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dStaticTable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1dStaticTable.EntityData.SegmentPath = "dot1dStaticTable"
    dot1dStaticTable.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/" + dot1dStaticTable.EntityData.SegmentPath
    dot1dStaticTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dStaticTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dStaticTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dStaticTable.EntityData.Children = types.NewOrderedMap()
    dot1dStaticTable.EntityData.Children.Append("dot1dStaticEntry", types.YChild{"Dot1dStaticEntry", nil})
    for i := range dot1dStaticTable.Dot1dStaticEntry {
        dot1dStaticTable.EntityData.Children.Append(types.GetSegmentPath(dot1dStaticTable.Dot1dStaticEntry[i]), types.YChild{"Dot1dStaticEntry", dot1dStaticTable.Dot1dStaticEntry[i]})
    }
    dot1dStaticTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dStaticTable.EntityData.YListKeys = []string {}

    return &(dot1dStaticTable.EntityData)
}

// BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry
// Filtering information configured into the bridge by
// (local or network) management specifying the set of
// ports to which frames received from a specific port and
// containing a specific destination address are allowed to
// be forwarded.
type BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object can take the value
    // of a unicast address, a group address, or the broadcast address. The type
    // is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1dStaticAddress interface{}

    // This attribute is a key. Either the value '0', or the port number of the
    // port from which a frame must be received in order for this entry's
    // filtering information to apply.  A value of zero indicates that this entry
    // applies on all ports of the bridge for which there is no other applicable
    // entry. The type is interface{} with range: 0..65535.
    Dot1dStaticReceivePort interface{}

    // The set of ports to which frames received from a specific port and destined
    // for a specific MAC address, are allowed to be forwarded.  Each octet within
    // the value of this object specifies a set of eight ports, with the first
    // octet specifying ports 1 through 8, the second octet specifying ports 9
    // through 16, etc.  Within each octet, the most significant bit represents
    // the lowest numbered port, and the least significant bit represents the
    // highest numbered port.  Thus, each port of the bridge is represented by a
    // single bit within the value of this object.  If that bit has a value of
    // '1', then that port is included in the set of ports; the port is not
    // included if its bit has a value of '0'.  (Note that the setting of the bit
    // corresponding to the port from which a frame is received is irrelevant.) 
    // The default value of this object is a string of ones of appropriate length.
    // The value of this object may exceed the required minimum maximum message
    // size of some SNMP transport (484 bytes, in the case of SNMP over UDP, see
    // RFC 3417, section 3.2). SNMP engines on bridges supporting a large number
    // of ports must support appropriate maximum message sizes. The type is string
    // with length: 0..512.
    Dot1dStaticAllowedToGoTo interface{}

    // This object indicates the status of this entry. The default value is
    // permanent(3).     other(1) - this entry is currently in use but the        
    // conditions under which it will remain so are         different from each of
    // the following values.     invalid(2) - writing this value to the object    
    // removes the corresponding entry.     permanent(3) - this entry is currently
    // in use and         will remain so after the next reset of the        
    // bridge.     deleteOnReset(4) - this entry is currently in use         and
    // will remain so until the next reset of the         bridge.    
    // deleteOnTimeout(5) - this entry is currently in use         and will remain
    // so until it is aged out. The type is Dot1dStaticStatus.
    Dot1dStaticStatus interface{}
}

func (dot1dStaticEntry *BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry) GetEntityData() *types.CommonEntityData {
    dot1dStaticEntry.EntityData.YFilter = dot1dStaticEntry.YFilter
    dot1dStaticEntry.EntityData.YangName = "dot1dStaticEntry"
    dot1dStaticEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dStaticEntry.EntityData.ParentYangName = "dot1dStaticTable"
    dot1dStaticEntry.EntityData.SegmentPath = "dot1dStaticEntry" + types.AddKeyToken(dot1dStaticEntry.Dot1dStaticAddress, "dot1dStaticAddress") + types.AddKeyToken(dot1dStaticEntry.Dot1dStaticReceivePort, "dot1dStaticReceivePort")
    dot1dStaticEntry.EntityData.AbsolutePath = "BRIDGE-MIB:BRIDGE-MIB/dot1dStaticTable/" + dot1dStaticEntry.EntityData.SegmentPath
    dot1dStaticEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dStaticEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dStaticEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dStaticEntry.EntityData.Children = types.NewOrderedMap()
    dot1dStaticEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dStaticEntry.EntityData.Leafs.Append("dot1dStaticAddress", types.YLeaf{"Dot1dStaticAddress", dot1dStaticEntry.Dot1dStaticAddress})
    dot1dStaticEntry.EntityData.Leafs.Append("dot1dStaticReceivePort", types.YLeaf{"Dot1dStaticReceivePort", dot1dStaticEntry.Dot1dStaticReceivePort})
    dot1dStaticEntry.EntityData.Leafs.Append("dot1dStaticAllowedToGoTo", types.YLeaf{"Dot1dStaticAllowedToGoTo", dot1dStaticEntry.Dot1dStaticAllowedToGoTo})
    dot1dStaticEntry.EntityData.Leafs.Append("dot1dStaticStatus", types.YLeaf{"Dot1dStaticStatus", dot1dStaticEntry.Dot1dStaticStatus})

    dot1dStaticEntry.EntityData.YListKeys = []string {"Dot1dStaticAddress", "Dot1dStaticReceivePort"}

    return &(dot1dStaticEntry.EntityData)
}

// BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus represents         and will remain so until it is aged out.
type BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus string

const (
    BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus_other BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus = "other"

    BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus_invalid BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus = "invalid"

    BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus_permanent BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus = "permanent"

    BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus_deleteOnReset BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus = "deleteOnReset"

    BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus_deleteOnTimeout BRIDGEMIB_Dot1dStaticTable_Dot1dStaticEntry_Dot1dStaticStatus = "deleteOnTimeout"
)

