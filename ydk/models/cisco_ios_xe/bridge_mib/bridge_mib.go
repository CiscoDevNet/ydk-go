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

    
    Dot1Dbase BRIDGEMIB_Dot1Dbase

    
    Dot1Dstp BRIDGEMIB_Dot1Dstp

    
    Dot1Dtp BRIDGEMIB_Dot1Dtp

    // A table that contains generic information about every port that is
    // associated with this bridge.  Transparent, source-route, and srt ports are
    // included.
    Dot1Dbaseporttable BRIDGEMIB_Dot1Dbaseporttable

    // A table that contains port-specific information for the Spanning Tree
    // Protocol.
    Dot1Dstpporttable BRIDGEMIB_Dot1Dstpporttable

    // A table that contains information about unicast entries for which the
    // bridge has forwarding and/or filtering information.  This information is
    // used by the transparent bridging function in determining how to propagate a
    // received frame.
    Dot1Dtpfdbtable BRIDGEMIB_Dot1Dtpfdbtable

    // A table that contains information about every port that is associated with
    // this transparent bridge.
    Dot1Dtpporttable BRIDGEMIB_Dot1Dtpporttable

    // A table containing filtering information configured into the bridge by
    // (local or network) management specifying the set of ports to which frames
    // received from specific ports and containing specific destination addresses
    // are allowed to be forwarded.  The value of zero in this table, as the port
    // number from which frames with a specific destination address are received,
    // is used to specify all ports for which there is no specific entry in this
    // table for that particular destination address.  Entries are valid for
    // unicast and for group/broadcast addresses.
    Dot1Dstatictable BRIDGEMIB_Dot1Dstatictable
}

func (bRIDGEMIB *BRIDGEMIB) GetEntityData() *types.CommonEntityData {
    bRIDGEMIB.EntityData.YFilter = bRIDGEMIB.YFilter
    bRIDGEMIB.EntityData.YangName = "BRIDGE-MIB"
    bRIDGEMIB.EntityData.BundleName = "cisco_ios_xe"
    bRIDGEMIB.EntityData.ParentYangName = "BRIDGE-MIB"
    bRIDGEMIB.EntityData.SegmentPath = "BRIDGE-MIB:BRIDGE-MIB"
    bRIDGEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bRIDGEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bRIDGEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bRIDGEMIB.EntityData.Children = make(map[string]types.YChild)
    bRIDGEMIB.EntityData.Children["dot1dBase"] = types.YChild{"Dot1Dbase", &bRIDGEMIB.Dot1Dbase}
    bRIDGEMIB.EntityData.Children["dot1dStp"] = types.YChild{"Dot1Dstp", &bRIDGEMIB.Dot1Dstp}
    bRIDGEMIB.EntityData.Children["dot1dTp"] = types.YChild{"Dot1Dtp", &bRIDGEMIB.Dot1Dtp}
    bRIDGEMIB.EntityData.Children["dot1dBasePortTable"] = types.YChild{"Dot1Dbaseporttable", &bRIDGEMIB.Dot1Dbaseporttable}
    bRIDGEMIB.EntityData.Children["dot1dStpPortTable"] = types.YChild{"Dot1Dstpporttable", &bRIDGEMIB.Dot1Dstpporttable}
    bRIDGEMIB.EntityData.Children["dot1dTpFdbTable"] = types.YChild{"Dot1Dtpfdbtable", &bRIDGEMIB.Dot1Dtpfdbtable}
    bRIDGEMIB.EntityData.Children["dot1dTpPortTable"] = types.YChild{"Dot1Dtpporttable", &bRIDGEMIB.Dot1Dtpporttable}
    bRIDGEMIB.EntityData.Children["dot1dStaticTable"] = types.YChild{"Dot1Dstatictable", &bRIDGEMIB.Dot1Dstatictable}
    bRIDGEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bRIDGEMIB.EntityData)
}

// BRIDGEMIB_Dot1Dbase
type BRIDGEMIB_Dot1Dbase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The MAC address used by this bridge when it must be referred to in a unique
    // fashion.  It is recommended that this be the numerically smallest MAC
    // address of all ports that belong to this bridge.  However, it is only 
    // required to be unique.  When concatenated with dot1dStpPriority, a unique
    // BridgeIdentifier is formed, which is used in the Spanning Tree Protocol.
    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Dbasebridgeaddress interface{}

    // The number of ports controlled by this bridging entity. The type is
    // interface{} with range: -2147483648..2147483647. Units are ports.
    Dot1Dbasenumports interface{}

    // Indicates what type of bridging this bridge can perform.  If a bridge is
    // actually performing a certain type of bridging, this will be indicated by
    // entries in the port table for the given type. The type is Dot1Dbasetype.
    Dot1Dbasetype interface{}
}

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetEntityData() *types.CommonEntityData {
    dot1Dbase.EntityData.YFilter = dot1Dbase.YFilter
    dot1Dbase.EntityData.YangName = "dot1dBase"
    dot1Dbase.EntityData.BundleName = "cisco_ios_xe"
    dot1Dbase.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dbase.EntityData.SegmentPath = "dot1dBase"
    dot1Dbase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dbase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dbase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dbase.EntityData.Children = make(map[string]types.YChild)
    dot1Dbase.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dbase.EntityData.Leafs["dot1dBaseBridgeAddress"] = types.YLeaf{"Dot1Dbasebridgeaddress", dot1Dbase.Dot1Dbasebridgeaddress}
    dot1Dbase.EntityData.Leafs["dot1dBaseNumPorts"] = types.YLeaf{"Dot1Dbasenumports", dot1Dbase.Dot1Dbasenumports}
    dot1Dbase.EntityData.Leafs["dot1dBaseType"] = types.YLeaf{"Dot1Dbasetype", dot1Dbase.Dot1Dbasetype}
    return &(dot1Dbase.EntityData)
}

// BRIDGEMIB_Dot1Dbase_Dot1Dbasetype represents entries in the port table for the given type.
type BRIDGEMIB_Dot1Dbase_Dot1Dbasetype string

const (
    BRIDGEMIB_Dot1Dbase_Dot1Dbasetype_unknown BRIDGEMIB_Dot1Dbase_Dot1Dbasetype = "unknown"

    BRIDGEMIB_Dot1Dbase_Dot1Dbasetype_transparent_only BRIDGEMIB_Dot1Dbase_Dot1Dbasetype = "transparent-only"

    BRIDGEMIB_Dot1Dbase_Dot1Dbasetype_sourceroute_only BRIDGEMIB_Dot1Dbase_Dot1Dbasetype = "sourceroute-only"

    BRIDGEMIB_Dot1Dbase_Dot1Dbasetype_srt BRIDGEMIB_Dot1Dbase_Dot1Dbasetype = "srt"
)

// BRIDGEMIB_Dot1Dstp
type BRIDGEMIB_Dot1Dstp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of what version of the Spanning Tree Protocol is being run. 
    // The value 'decLb100(2)' indicates the DEC LANbridge 100 Spanning Tree
    // protocol. IEEE 802.1D implementations will return 'ieee8021d(3)'. If future
    // versions of the IEEE Spanning Tree Protocol that are incompatible with the
    // current version are released a new value will be defined. The type is
    // Dot1Dstpprotocolspecification.
    Dot1Dstpprotocolspecification interface{}

    // The value of the write-able portion of the Bridge ID (i.e., the first two
    // octets of the (8 octet long) Bridge ID).  The other (last) 6 octets of the
    // Bridge ID are given by the value of dot1dBaseBridgeAddress. On bridges
    // supporting IEEE 802.1t or IEEE 802.1w, permissible values are 0-61440, in
    // steps of 4096. The type is interface{} with range: 0..65535.
    Dot1Dstppriority interface{}

    // The time (in hundredths of a second) since the last time a topology change
    // was detected by the bridge entity. For RSTP, this reports the time since
    // the tcWhile timer for any port on this Bridge was nonzero. The type is
    // interface{} with range: 0..4294967295. Units are centi-seconds.
    Dot1Dstptimesincetopologychange interface{}

    // The total number of topology changes detected by this bridge since the
    // management entity was last reset or initialized. The type is interface{}
    // with range: 0..4294967295.
    Dot1Dstptopchanges interface{}

    // The bridge identifier of the root of the spanning tree, as determined by
    // the Spanning Tree Protocol, as executed by this node.  This value is used
    // as the Root Identifier parameter in all Configuration Bridge PDUs
    // originated by this node. The type is string with length: 8.
    Dot1Dstpdesignatedroot interface{}

    // The cost of the path to the root as seen from this bridge. The type is
    // interface{} with range: -2147483648..2147483647.
    Dot1Dstprootcost interface{}

    // The port number of the port that offers the lowest cost path from this
    // bridge to the root bridge. The type is interface{} with range:
    // -2147483648..2147483647.
    Dot1Dstprootport interface{}

    // The maximum age of Spanning Tree Protocol information learned from the
    // network on any port before it is discarded, in units of hundredths of a
    // second.  This is the actual value that this bridge is currently using. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // centi-seconds.
    Dot1Dstpmaxage interface{}

    // The amount of time between the transmission of Configuration bridge PDUs by
    // this node on any port when it is the root of the spanning tree, or trying
    // to become so, in units of hundredths of a second.  This is the actual value
    // that this bridge is currently using. The type is interface{} with range:
    // -2147483648..2147483647. Units are centi-seconds.
    Dot1Dstphellotime interface{}

    // This time value determines the interval length during which no more than
    // two Configuration bridge PDUs shall be transmitted by this node, in units
    // of hundredths of a second. The type is interface{} with range:
    // -2147483648..2147483647. Units are centi-seconds.
    Dot1Dstpholdtime interface{}

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
    Dot1Dstpforwarddelay interface{}

    // The value that all bridges use for MaxAge when this bridge is acting as the
    // root.  Note that 802.1D-1998 specifies that the range for this parameter is
    // related to the value of dot1dStpBridgeHelloTime.  The granularity of this
    // timer is specified by 802.1D-1998 to be 1 second.  An agent may return a
    // badValue error if a set is attempted to a value that is not a whole number
    // of seconds. The type is interface{} with range: 600..4000. Units are
    // centi-seconds.
    Dot1Dstpbridgemaxage interface{}

    // The value that all bridges use for HelloTime when this bridge is acting as
    // the root.  The granularity of this timer is specified by 802.1D-1998 to be
    // 1 second.  An agent may return a badValue error if a set is attempted  to a
    // value that is not a whole number of seconds. The type is interface{} with
    // range: 100..1000. Units are centi-seconds.
    Dot1Dstpbridgehellotime interface{}

    // The value that all bridges use for ForwardDelay when this bridge is acting
    // as the root.  Note that 802.1D-1998 specifies that the range for this
    // parameter is related to the value of dot1dStpBridgeMaxAge.  The granularity
    // of this timer is specified by 802.1D-1998 to be 1 second.  An agent may
    // return a badValue error if a set is attempted to a value that is not a
    // whole number of seconds. The type is interface{} with range: 400..3000.
    // Units are centi-seconds.
    Dot1Dstpbridgeforwarddelay interface{}
}

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetEntityData() *types.CommonEntityData {
    dot1Dstp.EntityData.YFilter = dot1Dstp.YFilter
    dot1Dstp.EntityData.YangName = "dot1dStp"
    dot1Dstp.EntityData.BundleName = "cisco_ios_xe"
    dot1Dstp.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dstp.EntityData.SegmentPath = "dot1dStp"
    dot1Dstp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dstp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dstp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dstp.EntityData.Children = make(map[string]types.YChild)
    dot1Dstp.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dstp.EntityData.Leafs["dot1dStpProtocolSpecification"] = types.YLeaf{"Dot1Dstpprotocolspecification", dot1Dstp.Dot1Dstpprotocolspecification}
    dot1Dstp.EntityData.Leafs["dot1dStpPriority"] = types.YLeaf{"Dot1Dstppriority", dot1Dstp.Dot1Dstppriority}
    dot1Dstp.EntityData.Leafs["dot1dStpTimeSinceTopologyChange"] = types.YLeaf{"Dot1Dstptimesincetopologychange", dot1Dstp.Dot1Dstptimesincetopologychange}
    dot1Dstp.EntityData.Leafs["dot1dStpTopChanges"] = types.YLeaf{"Dot1Dstptopchanges", dot1Dstp.Dot1Dstptopchanges}
    dot1Dstp.EntityData.Leafs["dot1dStpDesignatedRoot"] = types.YLeaf{"Dot1Dstpdesignatedroot", dot1Dstp.Dot1Dstpdesignatedroot}
    dot1Dstp.EntityData.Leafs["dot1dStpRootCost"] = types.YLeaf{"Dot1Dstprootcost", dot1Dstp.Dot1Dstprootcost}
    dot1Dstp.EntityData.Leafs["dot1dStpRootPort"] = types.YLeaf{"Dot1Dstprootport", dot1Dstp.Dot1Dstprootport}
    dot1Dstp.EntityData.Leafs["dot1dStpMaxAge"] = types.YLeaf{"Dot1Dstpmaxage", dot1Dstp.Dot1Dstpmaxage}
    dot1Dstp.EntityData.Leafs["dot1dStpHelloTime"] = types.YLeaf{"Dot1Dstphellotime", dot1Dstp.Dot1Dstphellotime}
    dot1Dstp.EntityData.Leafs["dot1dStpHoldTime"] = types.YLeaf{"Dot1Dstpholdtime", dot1Dstp.Dot1Dstpholdtime}
    dot1Dstp.EntityData.Leafs["dot1dStpForwardDelay"] = types.YLeaf{"Dot1Dstpforwarddelay", dot1Dstp.Dot1Dstpforwarddelay}
    dot1Dstp.EntityData.Leafs["dot1dStpBridgeMaxAge"] = types.YLeaf{"Dot1Dstpbridgemaxage", dot1Dstp.Dot1Dstpbridgemaxage}
    dot1Dstp.EntityData.Leafs["dot1dStpBridgeHelloTime"] = types.YLeaf{"Dot1Dstpbridgehellotime", dot1Dstp.Dot1Dstpbridgehellotime}
    dot1Dstp.EntityData.Leafs["dot1dStpBridgeForwardDelay"] = types.YLeaf{"Dot1Dstpbridgeforwarddelay", dot1Dstp.Dot1Dstpbridgeforwarddelay}
    return &(dot1Dstp.EntityData)
}

// BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification represents are released a new value will be defined.
type BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification string

const (
    BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification_unknown BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification = "unknown"

    BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification_decLb100 BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification = "decLb100"

    BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification_ieee8021d BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification = "ieee8021d"
)

// BRIDGEMIB_Dot1Dtp
type BRIDGEMIB_Dot1Dtp struct {
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
    Dot1Dtplearnedentrydiscards interface{}

    // The timeout period in seconds for aging out dynamically-learned forwarding
    // information. 802.1D-1998 recommends a default of 300 seconds. The type is
    // interface{} with range: 10..1000000. Units are seconds.
    Dot1Dtpagingtime interface{}
}

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetEntityData() *types.CommonEntityData {
    dot1Dtp.EntityData.YFilter = dot1Dtp.YFilter
    dot1Dtp.EntityData.YangName = "dot1dTp"
    dot1Dtp.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtp.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dtp.EntityData.SegmentPath = "dot1dTp"
    dot1Dtp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtp.EntityData.Children = make(map[string]types.YChild)
    dot1Dtp.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dtp.EntityData.Leafs["dot1dTpLearnedEntryDiscards"] = types.YLeaf{"Dot1Dtplearnedentrydiscards", dot1Dtp.Dot1Dtplearnedentrydiscards}
    dot1Dtp.EntityData.Leafs["dot1dTpAgingTime"] = types.YLeaf{"Dot1Dtpagingtime", dot1Dtp.Dot1Dtpagingtime}
    return &(dot1Dtp.EntityData)
}

// BRIDGEMIB_Dot1Dbaseporttable
// A table that contains generic information about every
// port that is associated with this bridge.  Transparent,
// source-route, and srt ports are included.
type BRIDGEMIB_Dot1Dbaseporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of information for each port of the bridge. The type is slice of
    // BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry.
    Dot1Dbaseportentry []BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry
}

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetEntityData() *types.CommonEntityData {
    dot1Dbaseporttable.EntityData.YFilter = dot1Dbaseporttable.YFilter
    dot1Dbaseporttable.EntityData.YangName = "dot1dBasePortTable"
    dot1Dbaseporttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dbaseporttable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dbaseporttable.EntityData.SegmentPath = "dot1dBasePortTable"
    dot1Dbaseporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dbaseporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dbaseporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dbaseporttable.EntityData.Children = make(map[string]types.YChild)
    dot1Dbaseporttable.EntityData.Children["dot1dBasePortEntry"] = types.YChild{"Dot1Dbaseportentry", nil}
    for i := range dot1Dbaseporttable.Dot1Dbaseportentry {
        dot1Dbaseporttable.EntityData.Children[types.GetSegmentPath(&dot1Dbaseporttable.Dot1Dbaseportentry[i])] = types.YChild{"Dot1Dbaseportentry", &dot1Dbaseporttable.Dot1Dbaseportentry[i]}
    }
    dot1Dbaseporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dbaseporttable.EntityData)
}

// BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry
// A list of information for each port of the bridge.
type BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The port number of the port for which this entry
    // contains bridge management information. The type is interface{} with range:
    // 1..65535.
    Dot1Dbaseport interface{}

    // The value of the instance of the ifIndex object, defined in IF-MIB, for the
    // interface corresponding to this port. The type is interface{} with range:
    // 1..2147483647.
    Dot1Dbaseportifindex interface{}

    // For a port that (potentially) has the same value of dot1dBasePortIfIndex as
    // another port on the same bridge. This object contains the name of an object
    // instance unique to this port.  For example, in the case where multiple
    // ports correspond one-to-one with multiple X.25 virtual circuits, this value
    // might identify an (e.g., the first) object instance associated with the
    // X.25 virtual circuit corresponding to this port.  For a port which has a
    // unique value of dot1dBasePortIfIndex, this object can have the value { 0 0
    // }. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Dot1Dbaseportcircuit interface{}

    // The number of frames discarded by this port due to excessive transit delay
    // through the bridge.  It is incremented by both transparent and source route
    // bridges. The type is interface{} with range: 0..4294967295.
    Dot1Dbaseportdelayexceededdiscards interface{}

    // The number of frames discarded by this port due to an excessive size.  It
    // is incremented by both transparent and source route bridges. The type is
    // interface{} with range: 0..4294967295.
    Dot1Dbaseportmtuexceededdiscards interface{}

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
    Dot1Dportcapabilities interface{}

    // The default ingress User Priority for this port.  This only has effect on
    // media, such as Ethernet, that do not support native User Priority.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is interface{} with range: 0..7.
    Dot1Dportdefaultuserpriority interface{}

    // The number of egress traffic classes supported on this port.  This object
    // may optionally be read-only.  The value of this object MUST be retained
    // across reinitializations of the management system. The type is interface{}
    // with range: 1..8.
    Dot1Dportnumtrafficclasses interface{}

    // The GARP Join time, in centiseconds.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // interface{} with range: 0..2147483647.
    Dot1Dportgarpjointime interface{}

    // The GARP Leave time, in centiseconds.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // interface{} with range: 0..2147483647.
    Dot1Dportgarpleavetime interface{}

    // The GARP LeaveAll time, in centiseconds.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // interface{} with range: 0..2147483647.
    Dot1Dportgarpleavealltime interface{}

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
    Dot1Dportgmrpstatus interface{}

    // The total number of failed GMRP registrations, for any reason, in all
    // VLANs, on this port. The type is interface{} with range: 0..4294967295.
    Dot1Dportgmrpfailedregistrations interface{}

    // The Source MAC Address of the last GMRP message received on this port. The
    // type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Dportgmrplastpduorigin interface{}

    // The state of Restricted Group Registration on this port. If the value of
    // this control is true(1), then creation of a new dynamic entry is permitted
    // only if there is a Static Filtering Entry for the VLAN concerned, in which
    // the Registrar Administrative Control value is Normal Registration.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is bool.
    Dot1Dportrestrictedgroupregistration interface{}

    // The PVID, the VLAN-ID assigned to untagged frames or Priority-Tagged frames
    // received on this port.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is interface{} with
    // range: 0..4294967295.
    Dot1Qpvid interface{}

    // When this is admitOnlyVlanTagged(2), the device will discard untagged
    // frames or Priority-Tagged frames received on this port.  When admitAll(1),
    // untagged frames or Priority-Tagged frames received on this port will be
    // accepted and assigned to a VID based on the PVID and VID Set for this port.
    // This control does not affect VLAN-independent Bridge Protocol Data Unit
    // (BPDU) frames, such as GVRP and Spanning Tree Protocol (STP).  It does
    // affect VLAN- dependent BPDU frames, such as GMRP.  The value of this object
    // MUST be retained across reinitializations of the management system. The
    // type is Dot1Qportacceptableframetypes.
    Dot1Qportacceptableframetypes interface{}

    // When this is true(1), the device will discard incoming frames for VLANs
    // that do not include this Port in its  Member set.  When false(2), the port
    // will accept all incoming frames.  This control does not affect
    // VLAN-independent BPDU frames, such as GVRP and STP.  It does affect VLAN-
    // dependent BPDU frames, such as GMRP.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // bool.
    Dot1Qportingressfiltering interface{}

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
    Dot1Qportgvrpstatus interface{}

    // The total number of failed GVRP registrations, for any reason, on this
    // port. The type is interface{} with range: 0..4294967295.
    Dot1Qportgvrpfailedregistrations interface{}

    // The Source MAC Address of the last GVRP message received on this port. The
    // type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Qportgvrplastpduorigin interface{}

    // The state of Restricted VLAN Registration on this port. If the value of
    // this control is true(1), then creation of a new dynamic VLAN entry is
    // permitted only if there is a Static VLAN Registration Entry for the VLAN
    // concerned, in which the Registrar Administrative Control value for this
    // port is Normal Registration.  The value of this object MUST be retained
    // across reinitializations of the management system. The type is bool.
    Dot1Qportrestrictedvlanregistration interface{}
}

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetEntityData() *types.CommonEntityData {
    dot1Dbaseportentry.EntityData.YFilter = dot1Dbaseportentry.YFilter
    dot1Dbaseportentry.EntityData.YangName = "dot1dBasePortEntry"
    dot1Dbaseportentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dbaseportentry.EntityData.ParentYangName = "dot1dBasePortTable"
    dot1Dbaseportentry.EntityData.SegmentPath = "dot1dBasePortEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Dbaseportentry.Dot1Dbaseport) + "']"
    dot1Dbaseportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dbaseportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dbaseportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dbaseportentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dbaseportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dbaseportentry.EntityData.Leafs["dot1dBasePort"] = types.YLeaf{"Dot1Dbaseport", dot1Dbaseportentry.Dot1Dbaseport}
    dot1Dbaseportentry.EntityData.Leafs["dot1dBasePortIfIndex"] = types.YLeaf{"Dot1Dbaseportifindex", dot1Dbaseportentry.Dot1Dbaseportifindex}
    dot1Dbaseportentry.EntityData.Leafs["dot1dBasePortCircuit"] = types.YLeaf{"Dot1Dbaseportcircuit", dot1Dbaseportentry.Dot1Dbaseportcircuit}
    dot1Dbaseportentry.EntityData.Leafs["dot1dBasePortDelayExceededDiscards"] = types.YLeaf{"Dot1Dbaseportdelayexceededdiscards", dot1Dbaseportentry.Dot1Dbaseportdelayexceededdiscards}
    dot1Dbaseportentry.EntityData.Leafs["dot1dBasePortMtuExceededDiscards"] = types.YLeaf{"Dot1Dbaseportmtuexceededdiscards", dot1Dbaseportentry.Dot1Dbaseportmtuexceededdiscards}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortCapabilities"] = types.YLeaf{"Dot1Dportcapabilities", dot1Dbaseportentry.Dot1Dportcapabilities}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortDefaultUserPriority"] = types.YLeaf{"Dot1Dportdefaultuserpriority", dot1Dbaseportentry.Dot1Dportdefaultuserpriority}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortNumTrafficClasses"] = types.YLeaf{"Dot1Dportnumtrafficclasses", dot1Dbaseportentry.Dot1Dportnumtrafficclasses}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortGarpJoinTime"] = types.YLeaf{"Dot1Dportgarpjointime", dot1Dbaseportentry.Dot1Dportgarpjointime}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortGarpLeaveTime"] = types.YLeaf{"Dot1Dportgarpleavetime", dot1Dbaseportentry.Dot1Dportgarpleavetime}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortGarpLeaveAllTime"] = types.YLeaf{"Dot1Dportgarpleavealltime", dot1Dbaseportentry.Dot1Dportgarpleavealltime}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortGmrpStatus"] = types.YLeaf{"Dot1Dportgmrpstatus", dot1Dbaseportentry.Dot1Dportgmrpstatus}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortGmrpFailedRegistrations"] = types.YLeaf{"Dot1Dportgmrpfailedregistrations", dot1Dbaseportentry.Dot1Dportgmrpfailedregistrations}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortGmrpLastPduOrigin"] = types.YLeaf{"Dot1Dportgmrplastpduorigin", dot1Dbaseportentry.Dot1Dportgmrplastpduorigin}
    dot1Dbaseportentry.EntityData.Leafs["dot1dPortRestrictedGroupRegistration"] = types.YLeaf{"Dot1Dportrestrictedgroupregistration", dot1Dbaseportentry.Dot1Dportrestrictedgroupregistration}
    dot1Dbaseportentry.EntityData.Leafs["dot1qPvid"] = types.YLeaf{"Dot1Qpvid", dot1Dbaseportentry.Dot1Qpvid}
    dot1Dbaseportentry.EntityData.Leafs["dot1qPortAcceptableFrameTypes"] = types.YLeaf{"Dot1Qportacceptableframetypes", dot1Dbaseportentry.Dot1Qportacceptableframetypes}
    dot1Dbaseportentry.EntityData.Leafs["dot1qPortIngressFiltering"] = types.YLeaf{"Dot1Qportingressfiltering", dot1Dbaseportentry.Dot1Qportingressfiltering}
    dot1Dbaseportentry.EntityData.Leafs["dot1qPortGvrpStatus"] = types.YLeaf{"Dot1Qportgvrpstatus", dot1Dbaseportentry.Dot1Qportgvrpstatus}
    dot1Dbaseportentry.EntityData.Leafs["dot1qPortGvrpFailedRegistrations"] = types.YLeaf{"Dot1Qportgvrpfailedregistrations", dot1Dbaseportentry.Dot1Qportgvrpfailedregistrations}
    dot1Dbaseportentry.EntityData.Leafs["dot1qPortGvrpLastPduOrigin"] = types.YLeaf{"Dot1Qportgvrplastpduorigin", dot1Dbaseportentry.Dot1Qportgvrplastpduorigin}
    dot1Dbaseportentry.EntityData.Leafs["dot1qPortRestrictedVlanRegistration"] = types.YLeaf{"Dot1Qportrestrictedvlanregistration", dot1Dbaseportentry.Dot1Qportrestrictedvlanregistration}
    return &(dot1Dbaseportentry.EntityData)
}

// BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Qportacceptableframetypes represents reinitializations of the management system.
type BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Qportacceptableframetypes string

const (
    BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Qportacceptableframetypes_admitAll BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Qportacceptableframetypes = "admitAll"

    BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Qportacceptableframetypes_admitOnlyVlanTagged BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Qportacceptableframetypes = "admitOnlyVlanTagged"
)

// BRIDGEMIB_Dot1Dstpporttable
// A table that contains port-specific information
// for the Spanning Tree Protocol.
type BRIDGEMIB_Dot1Dstpporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of information maintained by every port about the Spanning Tree
    // Protocol state for that port. The type is slice of
    // BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry.
    Dot1Dstpportentry []BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry
}

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetEntityData() *types.CommonEntityData {
    dot1Dstpporttable.EntityData.YFilter = dot1Dstpporttable.YFilter
    dot1Dstpporttable.EntityData.YangName = "dot1dStpPortTable"
    dot1Dstpporttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dstpporttable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dstpporttable.EntityData.SegmentPath = "dot1dStpPortTable"
    dot1Dstpporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dstpporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dstpporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dstpporttable.EntityData.Children = make(map[string]types.YChild)
    dot1Dstpporttable.EntityData.Children["dot1dStpPortEntry"] = types.YChild{"Dot1Dstpportentry", nil}
    for i := range dot1Dstpporttable.Dot1Dstpportentry {
        dot1Dstpporttable.EntityData.Children[types.GetSegmentPath(&dot1Dstpporttable.Dot1Dstpportentry[i])] = types.YChild{"Dot1Dstpportentry", &dot1Dstpporttable.Dot1Dstpportentry[i]}
    }
    dot1Dstpporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dstpporttable.EntityData)
}

// BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry
// A list of information maintained by every port about
// the Spanning Tree Protocol state for that port.
type BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The port number of the port for which this entry
    // contains Spanning Tree Protocol management information. The type is
    // interface{} with range: 1..65535.
    Dot1Dstpport interface{}

    // The value of the priority field that is contained in the first (in network
    // byte order) octet of the (2 octet long) Port ID.  The other octet of the
    // Port ID is given by the value of dot1dStpPort. On bridges supporting IEEE
    // 802.1t or IEEE 802.1w, permissible values are 0-240, in steps of 16. The
    // type is interface{} with range: 0..255.
    Dot1Dstpportpriority interface{}

    // The port's current state, as defined by application of the Spanning Tree
    // Protocol.  This state controls what action a port takes on reception of a
    // frame.  If the bridge has detected a port that is malfunctioning, it will
    // place that port into the broken(6) state.  For ports that are disabled (see
    // dot1dStpPortEnable), this object will have a value of disabled(1). The type
    // is Dot1Dstpportstate.
    Dot1Dstpportstate interface{}

    // The enabled/disabled status of the port. The type is Dot1Dstpportenable.
    Dot1Dstpportenable interface{}

    // The contribution of this port to the path cost of paths towards the
    // spanning tree root which include this port.  802.1D-1998 recommends that
    // the default value of this parameter be in inverse proportion to  the speed
    // of the attached LAN.  New implementations should support
    // dot1dStpPortPathCost32. If the port path costs exceeds the maximum value of
    // this object then this object should report the maximum value, namely 65535.
    // Applications should try to read the dot1dStpPortPathCost32 object if this
    // object reports the maximum value. The type is interface{} with range:
    // 1..65535.
    Dot1Dstpportpathcost interface{}

    // The unique Bridge Identifier of the Bridge recorded as the Root in the
    // Configuration BPDUs transmitted by the Designated Bridge for the segment to
    // which the port is attached. The type is string with length: 8.
    Dot1Dstpportdesignatedroot interface{}

    // The path cost of the Designated Port of the segment connected to this port.
    // This value is compared to the Root Path Cost field in received bridge PDUs.
    // The type is interface{} with range: -2147483648..2147483647.
    Dot1Dstpportdesignatedcost interface{}

    // The Bridge Identifier of the bridge that this port considers to be the
    // Designated Bridge for this port's segment. The type is string with length:
    // 8.
    Dot1Dstpportdesignatedbridge interface{}

    // The Port Identifier of the port on the Designated Bridge for this port's
    // segment. The type is string with length: 2.
    Dot1Dstpportdesignatedport interface{}

    // The number of times this port has transitioned from the Learning state to
    // the Forwarding state. The type is interface{} with range: 0..4294967295.
    Dot1Dstpportforwardtransitions interface{}

    // The contribution of this port to the path cost of paths towards the
    // spanning tree root which include this port.  802.1D-1998 recommends that
    // the default value of this parameter be in inverse proportion to the speed
    // of the attached LAN.  This object replaces dot1dStpPortPathCost to support
    // IEEE 802.1t. The type is interface{} with range: 1..200000000.
    Dot1Dstpportpathcost32 interface{}

    // The contribution of this port to the path cost (in 32 bits value) of paths
    // towards the spanning tree root which include this port.  This object is
    // used to configure the spanning tree port path cost in 32 bits value range
    // when the stpxSpanningTreePathCostOperMode is long(2).  If the
    // stpxSpanningTreePathCostOperMode is short(1), this  MIB object is not
    // instantiated. The type is interface{} with range: 0..4294967295.
    Stpxlongstpportpathcost interface{}
}

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetEntityData() *types.CommonEntityData {
    dot1Dstpportentry.EntityData.YFilter = dot1Dstpportentry.YFilter
    dot1Dstpportentry.EntityData.YangName = "dot1dStpPortEntry"
    dot1Dstpportentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dstpportentry.EntityData.ParentYangName = "dot1dStpPortTable"
    dot1Dstpportentry.EntityData.SegmentPath = "dot1dStpPortEntry" + "[dot1dStpPort='" + fmt.Sprintf("%v", dot1Dstpportentry.Dot1Dstpport) + "']"
    dot1Dstpportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dstpportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dstpportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dstpportentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dstpportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPort"] = types.YLeaf{"Dot1Dstpport", dot1Dstpportentry.Dot1Dstpport}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortPriority"] = types.YLeaf{"Dot1Dstpportpriority", dot1Dstpportentry.Dot1Dstpportpriority}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortState"] = types.YLeaf{"Dot1Dstpportstate", dot1Dstpportentry.Dot1Dstpportstate}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortEnable"] = types.YLeaf{"Dot1Dstpportenable", dot1Dstpportentry.Dot1Dstpportenable}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortPathCost"] = types.YLeaf{"Dot1Dstpportpathcost", dot1Dstpportentry.Dot1Dstpportpathcost}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortDesignatedRoot"] = types.YLeaf{"Dot1Dstpportdesignatedroot", dot1Dstpportentry.Dot1Dstpportdesignatedroot}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortDesignatedCost"] = types.YLeaf{"Dot1Dstpportdesignatedcost", dot1Dstpportentry.Dot1Dstpportdesignatedcost}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortDesignatedBridge"] = types.YLeaf{"Dot1Dstpportdesignatedbridge", dot1Dstpportentry.Dot1Dstpportdesignatedbridge}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortDesignatedPort"] = types.YLeaf{"Dot1Dstpportdesignatedport", dot1Dstpportentry.Dot1Dstpportdesignatedport}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortForwardTransitions"] = types.YLeaf{"Dot1Dstpportforwardtransitions", dot1Dstpportentry.Dot1Dstpportforwardtransitions}
    dot1Dstpportentry.EntityData.Leafs["dot1dStpPortPathCost32"] = types.YLeaf{"Dot1Dstpportpathcost32", dot1Dstpportentry.Dot1Dstpportpathcost32}
    dot1Dstpportentry.EntityData.Leafs["stpxLongStpPortPathCost"] = types.YLeaf{"Stpxlongstpportpathcost", dot1Dstpportentry.Stpxlongstpportpathcost}
    return &(dot1Dstpportentry.EntityData)
}

// BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportenable represents The enabled/disabled status of the port.
type BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportenable string

const (
    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportenable_enabled BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportenable = "enabled"

    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportenable_disabled BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportenable = "disabled"
)

// BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate represents object will have a value of disabled(1).
type BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate string

const (
    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate_disabled BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate = "disabled"

    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate_blocking BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate = "blocking"

    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate_listening BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate = "listening"

    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate_learning BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate = "learning"

    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate_forwarding BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate = "forwarding"

    BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate_broken BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry_Dot1Dstpportstate = "broken"
)

// BRIDGEMIB_Dot1Dtpfdbtable
// A table that contains information about unicast
// entries for which the bridge has forwarding and/or
// filtering information.  This information is used
// by the transparent bridging function in
// determining how to propagate a received frame.
type BRIDGEMIB_Dot1Dtpfdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a specific unicast MAC address for which the bridge has
    // some forwarding and/or filtering information. The type is slice of
    // BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry.
    Dot1Dtpfdbentry []BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry
}

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetEntityData() *types.CommonEntityData {
    dot1Dtpfdbtable.EntityData.YFilter = dot1Dtpfdbtable.YFilter
    dot1Dtpfdbtable.EntityData.YangName = "dot1dTpFdbTable"
    dot1Dtpfdbtable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtpfdbtable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dtpfdbtable.EntityData.SegmentPath = "dot1dTpFdbTable"
    dot1Dtpfdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtpfdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtpfdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtpfdbtable.EntityData.Children = make(map[string]types.YChild)
    dot1Dtpfdbtable.EntityData.Children["dot1dTpFdbEntry"] = types.YChild{"Dot1Dtpfdbentry", nil}
    for i := range dot1Dtpfdbtable.Dot1Dtpfdbentry {
        dot1Dtpfdbtable.EntityData.Children[types.GetSegmentPath(&dot1Dtpfdbtable.Dot1Dtpfdbentry[i])] = types.YChild{"Dot1Dtpfdbentry", &dot1Dtpfdbtable.Dot1Dtpfdbentry[i]}
    }
    dot1Dtpfdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dtpfdbtable.EntityData)
}

// BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry
// Information about a specific unicast MAC address
// for which the bridge has some forwarding and/or
// filtering information.
type BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unicast MAC address for which the bridge has
    // forwarding and/or filtering information. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Dtpfdbaddress interface{}

    // Either the value '0', or the port number of the port on which a frame
    // having a source address equal to the value of the corresponding instance of
    // dot1dTpFdbAddress has been seen.  A value of '0' indicates that the port
    // number has not been learned, but that the bridge does have some
    // forwarding/filtering information about this address (e.g., in the
    // dot1dStaticTable).  Implementors are encouraged to assign the port value to
    // this object whenever it is learned, even for addresses for which the
    // corresponding value of dot1dTpFdbStatus is not learned(3). The type is
    // interface{} with range: -2147483648..2147483647.
    Dot1Dtpfdbport interface{}

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
    // dot1dStaticAddress. The type is Dot1Dtpfdbstatus.
    Dot1Dtpfdbstatus interface{}
}

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetEntityData() *types.CommonEntityData {
    dot1Dtpfdbentry.EntityData.YFilter = dot1Dtpfdbentry.YFilter
    dot1Dtpfdbentry.EntityData.YangName = "dot1dTpFdbEntry"
    dot1Dtpfdbentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtpfdbentry.EntityData.ParentYangName = "dot1dTpFdbTable"
    dot1Dtpfdbentry.EntityData.SegmentPath = "dot1dTpFdbEntry" + "[dot1dTpFdbAddress='" + fmt.Sprintf("%v", dot1Dtpfdbentry.Dot1Dtpfdbaddress) + "']"
    dot1Dtpfdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtpfdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtpfdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtpfdbentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dtpfdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dtpfdbentry.EntityData.Leafs["dot1dTpFdbAddress"] = types.YLeaf{"Dot1Dtpfdbaddress", dot1Dtpfdbentry.Dot1Dtpfdbaddress}
    dot1Dtpfdbentry.EntityData.Leafs["dot1dTpFdbPort"] = types.YLeaf{"Dot1Dtpfdbport", dot1Dtpfdbentry.Dot1Dtpfdbport}
    dot1Dtpfdbentry.EntityData.Leafs["dot1dTpFdbStatus"] = types.YLeaf{"Dot1Dtpfdbstatus", dot1Dtpfdbentry.Dot1Dtpfdbstatus}
    return &(dot1Dtpfdbentry.EntityData)
}

// BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus represents         existing instance of dot1dStaticAddress.
type BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus string

const (
    BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus_other BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus = "other"

    BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus_invalid BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus = "invalid"

    BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus_learned BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus = "learned"

    BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus_self BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus = "self"

    BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus_mgmt BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry_Dot1Dtpfdbstatus = "mgmt"
)

// BRIDGEMIB_Dot1Dtpporttable
// A table that contains information about every port that
// is associated with this transparent bridge.
type BRIDGEMIB_Dot1Dtpporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of information for each port of a transparent bridge. The type is
    // slice of BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry.
    Dot1Dtpportentry []BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry
}

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetEntityData() *types.CommonEntityData {
    dot1Dtpporttable.EntityData.YFilter = dot1Dtpporttable.YFilter
    dot1Dtpporttable.EntityData.YangName = "dot1dTpPortTable"
    dot1Dtpporttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtpporttable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dtpporttable.EntityData.SegmentPath = "dot1dTpPortTable"
    dot1Dtpporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtpporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtpporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtpporttable.EntityData.Children = make(map[string]types.YChild)
    dot1Dtpporttable.EntityData.Children["dot1dTpPortEntry"] = types.YChild{"Dot1Dtpportentry", nil}
    for i := range dot1Dtpporttable.Dot1Dtpportentry {
        dot1Dtpporttable.EntityData.Children[types.GetSegmentPath(&dot1Dtpporttable.Dot1Dtpportentry[i])] = types.YChild{"Dot1Dtpportentry", &dot1Dtpporttable.Dot1Dtpportentry[i]}
    }
    dot1Dtpporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dtpporttable.EntityData)
}

// BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry
// A list of information for each port of a transparent
// bridge.
type BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The port number of the port for which this entry
    // contains Transparent bridging management information. The type is
    // interface{} with range: 1..65535.
    Dot1Dtpport interface{}

    // The maximum size of the INFO (non-MAC) field that  this port will receive
    // or transmit. The type is interface{} with range: -2147483648..2147483647.
    // Units are bytes.
    Dot1Dtpportmaxinfo interface{}

    // The number of frames that have been received by this port from its segment.
    // Note that a frame received on the interface corresponding to this port is
    // only counted by this object if and only if it is for a protocol being
    // processed by the local bridging function, including bridge management
    // frames. The type is interface{} with range: 0..4294967295. Units are
    // frames.
    Dot1Dtpportinframes interface{}

    // The number of frames that have been transmitted by this port to its
    // segment.  Note that a frame transmitted on the interface corresponding to
    // this port is only counted by this object if and only if it is for a
    // protocol being processed by the local bridging function, including bridge
    // management frames. The type is interface{} with range: 0..4294967295. Units
    // are frames.
    Dot1Dtpportoutframes interface{}

    // Count of received valid frames that were discarded (i.e., filtered) by the
    // Forwarding Process. The type is interface{} with range: 0..4294967295.
    // Units are frames.
    Dot1Dtpportindiscards interface{}
}

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetEntityData() *types.CommonEntityData {
    dot1Dtpportentry.EntityData.YFilter = dot1Dtpportentry.YFilter
    dot1Dtpportentry.EntityData.YangName = "dot1dTpPortEntry"
    dot1Dtpportentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtpportentry.EntityData.ParentYangName = "dot1dTpPortTable"
    dot1Dtpportentry.EntityData.SegmentPath = "dot1dTpPortEntry" + "[dot1dTpPort='" + fmt.Sprintf("%v", dot1Dtpportentry.Dot1Dtpport) + "']"
    dot1Dtpportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtpportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtpportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtpportentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dtpportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dtpportentry.EntityData.Leafs["dot1dTpPort"] = types.YLeaf{"Dot1Dtpport", dot1Dtpportentry.Dot1Dtpport}
    dot1Dtpportentry.EntityData.Leafs["dot1dTpPortMaxInfo"] = types.YLeaf{"Dot1Dtpportmaxinfo", dot1Dtpportentry.Dot1Dtpportmaxinfo}
    dot1Dtpportentry.EntityData.Leafs["dot1dTpPortInFrames"] = types.YLeaf{"Dot1Dtpportinframes", dot1Dtpportentry.Dot1Dtpportinframes}
    dot1Dtpportentry.EntityData.Leafs["dot1dTpPortOutFrames"] = types.YLeaf{"Dot1Dtpportoutframes", dot1Dtpportentry.Dot1Dtpportoutframes}
    dot1Dtpportentry.EntityData.Leafs["dot1dTpPortInDiscards"] = types.YLeaf{"Dot1Dtpportindiscards", dot1Dtpportentry.Dot1Dtpportindiscards}
    return &(dot1Dtpportentry.EntityData)
}

// BRIDGEMIB_Dot1Dstatictable
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
type BRIDGEMIB_Dot1Dstatictable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the bridge by (local or network)
    // management specifying the set of ports to which frames received from a
    // specific port and containing a specific destination address are allowed to
    // be forwarded. The type is slice of
    // BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry.
    Dot1Dstaticentry []BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry
}

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetEntityData() *types.CommonEntityData {
    dot1Dstatictable.EntityData.YFilter = dot1Dstatictable.YFilter
    dot1Dstatictable.EntityData.YangName = "dot1dStaticTable"
    dot1Dstatictable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dstatictable.EntityData.ParentYangName = "BRIDGE-MIB"
    dot1Dstatictable.EntityData.SegmentPath = "dot1dStaticTable"
    dot1Dstatictable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dstatictable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dstatictable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dstatictable.EntityData.Children = make(map[string]types.YChild)
    dot1Dstatictable.EntityData.Children["dot1dStaticEntry"] = types.YChild{"Dot1Dstaticentry", nil}
    for i := range dot1Dstatictable.Dot1Dstaticentry {
        dot1Dstatictable.EntityData.Children[types.GetSegmentPath(&dot1Dstatictable.Dot1Dstaticentry[i])] = types.YChild{"Dot1Dstaticentry", &dot1Dstatictable.Dot1Dstaticentry[i]}
    }
    dot1Dstatictable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dstatictable.EntityData)
}

// BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry
// Filtering information configured into the bridge by
// (local or network) management specifying the set of
// ports to which frames received from a specific port and
// containing a specific destination address are allowed to
// be forwarded.
type BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object can take the value
    // of a unicast address, a group address, or the broadcast address. The type
    // is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Dstaticaddress interface{}

    // This attribute is a key. Either the value '0', or the port number of the
    // port from which a frame must be received in order for this entry's
    // filtering information to apply.  A value of zero indicates that this entry
    // applies on all ports of the bridge for which there is no other applicable
    // entry. The type is interface{} with range: 0..65535.
    Dot1Dstaticreceiveport interface{}

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
    Dot1Dstaticallowedtogoto interface{}

    // This object indicates the status of this entry. The default value is
    // permanent(3).     other(1) - this entry is currently in use but the        
    // conditions under which it will remain so are         different from each of
    // the following values.     invalid(2) - writing this value to the object    
    // removes the corresponding entry.     permanent(3) - this entry is currently
    // in use and         will remain so after the next reset of the        
    // bridge.     deleteOnReset(4) - this entry is currently in use         and
    // will remain so until the next reset of the         bridge.    
    // deleteOnTimeout(5) - this entry is currently in use         and will remain
    // so until it is aged out. The type is Dot1Dstaticstatus.
    Dot1Dstaticstatus interface{}
}

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetEntityData() *types.CommonEntityData {
    dot1Dstaticentry.EntityData.YFilter = dot1Dstaticentry.YFilter
    dot1Dstaticentry.EntityData.YangName = "dot1dStaticEntry"
    dot1Dstaticentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dstaticentry.EntityData.ParentYangName = "dot1dStaticTable"
    dot1Dstaticentry.EntityData.SegmentPath = "dot1dStaticEntry" + "[dot1dStaticAddress='" + fmt.Sprintf("%v", dot1Dstaticentry.Dot1Dstaticaddress) + "']" + "[dot1dStaticReceivePort='" + fmt.Sprintf("%v", dot1Dstaticentry.Dot1Dstaticreceiveport) + "']"
    dot1Dstaticentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dstaticentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dstaticentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dstaticentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dstaticentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dstaticentry.EntityData.Leafs["dot1dStaticAddress"] = types.YLeaf{"Dot1Dstaticaddress", dot1Dstaticentry.Dot1Dstaticaddress}
    dot1Dstaticentry.EntityData.Leafs["dot1dStaticReceivePort"] = types.YLeaf{"Dot1Dstaticreceiveport", dot1Dstaticentry.Dot1Dstaticreceiveport}
    dot1Dstaticentry.EntityData.Leafs["dot1dStaticAllowedToGoTo"] = types.YLeaf{"Dot1Dstaticallowedtogoto", dot1Dstaticentry.Dot1Dstaticallowedtogoto}
    dot1Dstaticentry.EntityData.Leafs["dot1dStaticStatus"] = types.YLeaf{"Dot1Dstaticstatus", dot1Dstaticentry.Dot1Dstaticstatus}
    return &(dot1Dstaticentry.EntityData)
}

// BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus represents         and will remain so until it is aged out.
type BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus string

const (
    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_other BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "other"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_invalid BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "invalid"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_permanent BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "permanent"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_deleteOnReset BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "deleteOnReset"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_deleteOnTimeout BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "deleteOnTimeout"
)

