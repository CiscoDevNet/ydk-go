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
    parent types.Entity
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

func (bRIDGEMIB *BRIDGEMIB) GetFilter() yfilter.YFilter { return bRIDGEMIB.YFilter }

func (bRIDGEMIB *BRIDGEMIB) SetFilter(yf yfilter.YFilter) { bRIDGEMIB.YFilter = yf }

func (bRIDGEMIB *BRIDGEMIB) GetGoName(yname string) string {
    if yname == "dot1dBase" { return "Dot1Dbase" }
    if yname == "dot1dStp" { return "Dot1Dstp" }
    if yname == "dot1dTp" { return "Dot1Dtp" }
    if yname == "dot1dBasePortTable" { return "Dot1Dbaseporttable" }
    if yname == "dot1dStpPortTable" { return "Dot1Dstpporttable" }
    if yname == "dot1dTpFdbTable" { return "Dot1Dtpfdbtable" }
    if yname == "dot1dTpPortTable" { return "Dot1Dtpporttable" }
    if yname == "dot1dStaticTable" { return "Dot1Dstatictable" }
    return ""
}

func (bRIDGEMIB *BRIDGEMIB) GetSegmentPath() string {
    return "BRIDGE-MIB:BRIDGE-MIB"
}

func (bRIDGEMIB *BRIDGEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dBase" {
        return &bRIDGEMIB.Dot1Dbase
    }
    if childYangName == "dot1dStp" {
        return &bRIDGEMIB.Dot1Dstp
    }
    if childYangName == "dot1dTp" {
        return &bRIDGEMIB.Dot1Dtp
    }
    if childYangName == "dot1dBasePortTable" {
        return &bRIDGEMIB.Dot1Dbaseporttable
    }
    if childYangName == "dot1dStpPortTable" {
        return &bRIDGEMIB.Dot1Dstpporttable
    }
    if childYangName == "dot1dTpFdbTable" {
        return &bRIDGEMIB.Dot1Dtpfdbtable
    }
    if childYangName == "dot1dTpPortTable" {
        return &bRIDGEMIB.Dot1Dtpporttable
    }
    if childYangName == "dot1dStaticTable" {
        return &bRIDGEMIB.Dot1Dstatictable
    }
    return nil
}

func (bRIDGEMIB *BRIDGEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dot1dBase"] = &bRIDGEMIB.Dot1Dbase
    children["dot1dStp"] = &bRIDGEMIB.Dot1Dstp
    children["dot1dTp"] = &bRIDGEMIB.Dot1Dtp
    children["dot1dBasePortTable"] = &bRIDGEMIB.Dot1Dbaseporttable
    children["dot1dStpPortTable"] = &bRIDGEMIB.Dot1Dstpporttable
    children["dot1dTpFdbTable"] = &bRIDGEMIB.Dot1Dtpfdbtable
    children["dot1dTpPortTable"] = &bRIDGEMIB.Dot1Dtpporttable
    children["dot1dStaticTable"] = &bRIDGEMIB.Dot1Dstatictable
    return children
}

func (bRIDGEMIB *BRIDGEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bRIDGEMIB *BRIDGEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (bRIDGEMIB *BRIDGEMIB) GetYangName() string { return "BRIDGE-MIB" }

func (bRIDGEMIB *BRIDGEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bRIDGEMIB *BRIDGEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bRIDGEMIB *BRIDGEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bRIDGEMIB *BRIDGEMIB) SetParent(parent types.Entity) { bRIDGEMIB.parent = parent }

func (bRIDGEMIB *BRIDGEMIB) GetParent() types.Entity { return bRIDGEMIB.parent }

func (bRIDGEMIB *BRIDGEMIB) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dbase
type BRIDGEMIB_Dot1Dbase struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The MAC address used by this bridge when it must be referred to in a unique
    // fashion.  It is recommended that this be the numerically smallest MAC
    // address of all ports that belong to this bridge.  However, it is only 
    // required to be unique.  When concatenated with dot1dStpPriority, a unique
    // BridgeIdentifier is formed, which is used in the Spanning Tree Protocol.
    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1Dbasebridgeaddress interface{}

    // The number of ports controlled by this bridging entity. The type is
    // interface{} with range: -2147483648..2147483647. Units are ports.
    Dot1Dbasenumports interface{}

    // Indicates what type of bridging this bridge can perform.  If a bridge is
    // actually performing a certain type of bridging, this will be indicated by
    // entries in the port table for the given type. The type is Dot1Dbasetype.
    Dot1Dbasetype interface{}
}

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetFilter() yfilter.YFilter { return dot1Dbase.YFilter }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) SetFilter(yf yfilter.YFilter) { dot1Dbase.YFilter = yf }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetGoName(yname string) string {
    if yname == "dot1dBaseBridgeAddress" { return "Dot1Dbasebridgeaddress" }
    if yname == "dot1dBaseNumPorts" { return "Dot1Dbasenumports" }
    if yname == "dot1dBaseType" { return "Dot1Dbasetype" }
    return ""
}

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetSegmentPath() string {
    return "dot1dBase"
}

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBaseBridgeAddress"] = dot1Dbase.Dot1Dbasebridgeaddress
    leafs["dot1dBaseNumPorts"] = dot1Dbase.Dot1Dbasenumports
    leafs["dot1dBaseType"] = dot1Dbase.Dot1Dbasetype
    return leafs
}

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetYangName() string { return "dot1dBase" }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) SetParent(parent types.Entity) { dot1Dbase.parent = parent }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetParent() types.Entity { return dot1Dbase.parent }

func (dot1Dbase *BRIDGEMIB_Dot1Dbase) GetParentYangName() string { return "BRIDGE-MIB" }

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
    parent types.Entity
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

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetFilter() yfilter.YFilter { return dot1Dstp.YFilter }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) SetFilter(yf yfilter.YFilter) { dot1Dstp.YFilter = yf }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetGoName(yname string) string {
    if yname == "dot1dStpProtocolSpecification" { return "Dot1Dstpprotocolspecification" }
    if yname == "dot1dStpPriority" { return "Dot1Dstppriority" }
    if yname == "dot1dStpTimeSinceTopologyChange" { return "Dot1Dstptimesincetopologychange" }
    if yname == "dot1dStpTopChanges" { return "Dot1Dstptopchanges" }
    if yname == "dot1dStpDesignatedRoot" { return "Dot1Dstpdesignatedroot" }
    if yname == "dot1dStpRootCost" { return "Dot1Dstprootcost" }
    if yname == "dot1dStpRootPort" { return "Dot1Dstprootport" }
    if yname == "dot1dStpMaxAge" { return "Dot1Dstpmaxage" }
    if yname == "dot1dStpHelloTime" { return "Dot1Dstphellotime" }
    if yname == "dot1dStpHoldTime" { return "Dot1Dstpholdtime" }
    if yname == "dot1dStpForwardDelay" { return "Dot1Dstpforwarddelay" }
    if yname == "dot1dStpBridgeMaxAge" { return "Dot1Dstpbridgemaxage" }
    if yname == "dot1dStpBridgeHelloTime" { return "Dot1Dstpbridgehellotime" }
    if yname == "dot1dStpBridgeForwardDelay" { return "Dot1Dstpbridgeforwarddelay" }
    return ""
}

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetSegmentPath() string {
    return "dot1dStp"
}

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dStpProtocolSpecification"] = dot1Dstp.Dot1Dstpprotocolspecification
    leafs["dot1dStpPriority"] = dot1Dstp.Dot1Dstppriority
    leafs["dot1dStpTimeSinceTopologyChange"] = dot1Dstp.Dot1Dstptimesincetopologychange
    leafs["dot1dStpTopChanges"] = dot1Dstp.Dot1Dstptopchanges
    leafs["dot1dStpDesignatedRoot"] = dot1Dstp.Dot1Dstpdesignatedroot
    leafs["dot1dStpRootCost"] = dot1Dstp.Dot1Dstprootcost
    leafs["dot1dStpRootPort"] = dot1Dstp.Dot1Dstprootport
    leafs["dot1dStpMaxAge"] = dot1Dstp.Dot1Dstpmaxage
    leafs["dot1dStpHelloTime"] = dot1Dstp.Dot1Dstphellotime
    leafs["dot1dStpHoldTime"] = dot1Dstp.Dot1Dstpholdtime
    leafs["dot1dStpForwardDelay"] = dot1Dstp.Dot1Dstpforwarddelay
    leafs["dot1dStpBridgeMaxAge"] = dot1Dstp.Dot1Dstpbridgemaxage
    leafs["dot1dStpBridgeHelloTime"] = dot1Dstp.Dot1Dstpbridgehellotime
    leafs["dot1dStpBridgeForwardDelay"] = dot1Dstp.Dot1Dstpbridgeforwarddelay
    return leafs
}

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetYangName() string { return "dot1dStp" }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) SetParent(parent types.Entity) { dot1Dstp.parent = parent }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetParent() types.Entity { return dot1Dstp.parent }

func (dot1Dstp *BRIDGEMIB_Dot1Dstp) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification represents are released a new value will be defined.
type BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification string

const (
    BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification_unknown BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification = "unknown"

    BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification_decLb100 BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification = "decLb100"

    BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification_ieee8021d BRIDGEMIB_Dot1Dstp_Dot1Dstpprotocolspecification = "ieee8021d"
)

// BRIDGEMIB_Dot1Dtp
type BRIDGEMIB_Dot1Dtp struct {
    parent types.Entity
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

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetFilter() yfilter.YFilter { return dot1Dtp.YFilter }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) SetFilter(yf yfilter.YFilter) { dot1Dtp.YFilter = yf }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetGoName(yname string) string {
    if yname == "dot1dTpLearnedEntryDiscards" { return "Dot1Dtplearnedentrydiscards" }
    if yname == "dot1dTpAgingTime" { return "Dot1Dtpagingtime" }
    return ""
}

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetSegmentPath() string {
    return "dot1dTp"
}

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dTpLearnedEntryDiscards"] = dot1Dtp.Dot1Dtplearnedentrydiscards
    leafs["dot1dTpAgingTime"] = dot1Dtp.Dot1Dtpagingtime
    return leafs
}

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetYangName() string { return "dot1dTp" }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) SetParent(parent types.Entity) { dot1Dtp.parent = parent }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetParent() types.Entity { return dot1Dtp.parent }

func (dot1Dtp *BRIDGEMIB_Dot1Dtp) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dbaseporttable
// A table that contains generic information about every
// port that is associated with this bridge.  Transparent,
// source-route, and srt ports are included.
type BRIDGEMIB_Dot1Dbaseporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of information for each port of the bridge. The type is slice of
    // BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry.
    Dot1Dbaseportentry []BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry
}

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetFilter() yfilter.YFilter { return dot1Dbaseporttable.YFilter }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) SetFilter(yf yfilter.YFilter) { dot1Dbaseporttable.YFilter = yf }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetGoName(yname string) string {
    if yname == "dot1dBasePortEntry" { return "Dot1Dbaseportentry" }
    return ""
}

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetSegmentPath() string {
    return "dot1dBasePortTable"
}

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dBasePortEntry" {
        for _, c := range dot1Dbaseporttable.Dot1Dbaseportentry {
            if dot1Dbaseporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry{}
        dot1Dbaseporttable.Dot1Dbaseportentry = append(dot1Dbaseporttable.Dot1Dbaseportentry, child)
        return &dot1Dbaseporttable.Dot1Dbaseportentry[len(dot1Dbaseporttable.Dot1Dbaseportentry)-1]
    }
    return nil
}

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dbaseporttable.Dot1Dbaseportentry {
        children[dot1Dbaseporttable.Dot1Dbaseportentry[i].GetSegmentPath()] = &dot1Dbaseporttable.Dot1Dbaseportentry[i]
    }
    return children
}

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetYangName() string { return "dot1dBasePortTable" }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) SetParent(parent types.Entity) { dot1Dbaseporttable.parent = parent }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetParent() types.Entity { return dot1Dbaseporttable.parent }

func (dot1Dbaseporttable *BRIDGEMIB_Dot1Dbaseporttable) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry
// A list of information for each port of the bridge.
type BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1Qportgvrplastpduorigin interface{}

    // The state of Restricted VLAN Registration on this port. If the value of
    // this control is true(1), then creation of a new dynamic VLAN entry is
    // permitted only if there is a Static VLAN Registration Entry for the VLAN
    // concerned, in which the Registrar Administrative Control value for this
    // port is Normal Registration.  The value of this object MUST be retained
    // across reinitializations of the management system. The type is bool.
    Dot1Qportrestrictedvlanregistration interface{}
}

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetFilter() yfilter.YFilter { return dot1Dbaseportentry.YFilter }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) SetFilter(yf yfilter.YFilter) { dot1Dbaseportentry.YFilter = yf }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetGoName(yname string) string {
    if yname == "dot1dBasePort" { return "Dot1Dbaseport" }
    if yname == "dot1dBasePortIfIndex" { return "Dot1Dbaseportifindex" }
    if yname == "dot1dBasePortCircuit" { return "Dot1Dbaseportcircuit" }
    if yname == "dot1dBasePortDelayExceededDiscards" { return "Dot1Dbaseportdelayexceededdiscards" }
    if yname == "dot1dBasePortMtuExceededDiscards" { return "Dot1Dbaseportmtuexceededdiscards" }
    if yname == "dot1dPortCapabilities" { return "Dot1Dportcapabilities" }
    if yname == "dot1dPortDefaultUserPriority" { return "Dot1Dportdefaultuserpriority" }
    if yname == "dot1dPortNumTrafficClasses" { return "Dot1Dportnumtrafficclasses" }
    if yname == "dot1dPortGarpJoinTime" { return "Dot1Dportgarpjointime" }
    if yname == "dot1dPortGarpLeaveTime" { return "Dot1Dportgarpleavetime" }
    if yname == "dot1dPortGarpLeaveAllTime" { return "Dot1Dportgarpleavealltime" }
    if yname == "dot1dPortGmrpStatus" { return "Dot1Dportgmrpstatus" }
    if yname == "dot1dPortGmrpFailedRegistrations" { return "Dot1Dportgmrpfailedregistrations" }
    if yname == "dot1dPortGmrpLastPduOrigin" { return "Dot1Dportgmrplastpduorigin" }
    if yname == "dot1dPortRestrictedGroupRegistration" { return "Dot1Dportrestrictedgroupregistration" }
    if yname == "dot1qPvid" { return "Dot1Qpvid" }
    if yname == "dot1qPortAcceptableFrameTypes" { return "Dot1Qportacceptableframetypes" }
    if yname == "dot1qPortIngressFiltering" { return "Dot1Qportingressfiltering" }
    if yname == "dot1qPortGvrpStatus" { return "Dot1Qportgvrpstatus" }
    if yname == "dot1qPortGvrpFailedRegistrations" { return "Dot1Qportgvrpfailedregistrations" }
    if yname == "dot1qPortGvrpLastPduOrigin" { return "Dot1Qportgvrplastpduorigin" }
    if yname == "dot1qPortRestrictedVlanRegistration" { return "Dot1Qportrestrictedvlanregistration" }
    return ""
}

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetSegmentPath() string {
    return "dot1dBasePortEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Dbaseportentry.Dot1Dbaseport) + "']"
}

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBasePort"] = dot1Dbaseportentry.Dot1Dbaseport
    leafs["dot1dBasePortIfIndex"] = dot1Dbaseportentry.Dot1Dbaseportifindex
    leafs["dot1dBasePortCircuit"] = dot1Dbaseportentry.Dot1Dbaseportcircuit
    leafs["dot1dBasePortDelayExceededDiscards"] = dot1Dbaseportentry.Dot1Dbaseportdelayexceededdiscards
    leafs["dot1dBasePortMtuExceededDiscards"] = dot1Dbaseportentry.Dot1Dbaseportmtuexceededdiscards
    leafs["dot1dPortCapabilities"] = dot1Dbaseportentry.Dot1Dportcapabilities
    leafs["dot1dPortDefaultUserPriority"] = dot1Dbaseportentry.Dot1Dportdefaultuserpriority
    leafs["dot1dPortNumTrafficClasses"] = dot1Dbaseportentry.Dot1Dportnumtrafficclasses
    leafs["dot1dPortGarpJoinTime"] = dot1Dbaseportentry.Dot1Dportgarpjointime
    leafs["dot1dPortGarpLeaveTime"] = dot1Dbaseportentry.Dot1Dportgarpleavetime
    leafs["dot1dPortGarpLeaveAllTime"] = dot1Dbaseportentry.Dot1Dportgarpleavealltime
    leafs["dot1dPortGmrpStatus"] = dot1Dbaseportentry.Dot1Dportgmrpstatus
    leafs["dot1dPortGmrpFailedRegistrations"] = dot1Dbaseportentry.Dot1Dportgmrpfailedregistrations
    leafs["dot1dPortGmrpLastPduOrigin"] = dot1Dbaseportentry.Dot1Dportgmrplastpduorigin
    leafs["dot1dPortRestrictedGroupRegistration"] = dot1Dbaseportentry.Dot1Dportrestrictedgroupregistration
    leafs["dot1qPvid"] = dot1Dbaseportentry.Dot1Qpvid
    leafs["dot1qPortAcceptableFrameTypes"] = dot1Dbaseportentry.Dot1Qportacceptableframetypes
    leafs["dot1qPortIngressFiltering"] = dot1Dbaseportentry.Dot1Qportingressfiltering
    leafs["dot1qPortGvrpStatus"] = dot1Dbaseportentry.Dot1Qportgvrpstatus
    leafs["dot1qPortGvrpFailedRegistrations"] = dot1Dbaseportentry.Dot1Qportgvrpfailedregistrations
    leafs["dot1qPortGvrpLastPduOrigin"] = dot1Dbaseportentry.Dot1Qportgvrplastpduorigin
    leafs["dot1qPortRestrictedVlanRegistration"] = dot1Dbaseportentry.Dot1Qportrestrictedvlanregistration
    return leafs
}

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetYangName() string { return "dot1dBasePortEntry" }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) SetParent(parent types.Entity) { dot1Dbaseportentry.parent = parent }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetParent() types.Entity { return dot1Dbaseportentry.parent }

func (dot1Dbaseportentry *BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry) GetParentYangName() string { return "dot1dBasePortTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of information maintained by every port about the Spanning Tree
    // Protocol state for that port. The type is slice of
    // BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry.
    Dot1Dstpportentry []BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry
}

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetFilter() yfilter.YFilter { return dot1Dstpporttable.YFilter }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) SetFilter(yf yfilter.YFilter) { dot1Dstpporttable.YFilter = yf }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetGoName(yname string) string {
    if yname == "dot1dStpPortEntry" { return "Dot1Dstpportentry" }
    return ""
}

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetSegmentPath() string {
    return "dot1dStpPortTable"
}

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dStpPortEntry" {
        for _, c := range dot1Dstpporttable.Dot1Dstpportentry {
            if dot1Dstpporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry{}
        dot1Dstpporttable.Dot1Dstpportentry = append(dot1Dstpporttable.Dot1Dstpportentry, child)
        return &dot1Dstpporttable.Dot1Dstpportentry[len(dot1Dstpporttable.Dot1Dstpportentry)-1]
    }
    return nil
}

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dstpporttable.Dot1Dstpportentry {
        children[dot1Dstpporttable.Dot1Dstpportentry[i].GetSegmentPath()] = &dot1Dstpporttable.Dot1Dstpportentry[i]
    }
    return children
}

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetYangName() string { return "dot1dStpPortTable" }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) SetParent(parent types.Entity) { dot1Dstpporttable.parent = parent }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetParent() types.Entity { return dot1Dstpporttable.parent }

func (dot1Dstpporttable *BRIDGEMIB_Dot1Dstpporttable) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry
// A list of information maintained by every port about
// the Spanning Tree Protocol state for that port.
type BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry struct {
    parent types.Entity
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

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetFilter() yfilter.YFilter { return dot1Dstpportentry.YFilter }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) SetFilter(yf yfilter.YFilter) { dot1Dstpportentry.YFilter = yf }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetGoName(yname string) string {
    if yname == "dot1dStpPort" { return "Dot1Dstpport" }
    if yname == "dot1dStpPortPriority" { return "Dot1Dstpportpriority" }
    if yname == "dot1dStpPortState" { return "Dot1Dstpportstate" }
    if yname == "dot1dStpPortEnable" { return "Dot1Dstpportenable" }
    if yname == "dot1dStpPortPathCost" { return "Dot1Dstpportpathcost" }
    if yname == "dot1dStpPortDesignatedRoot" { return "Dot1Dstpportdesignatedroot" }
    if yname == "dot1dStpPortDesignatedCost" { return "Dot1Dstpportdesignatedcost" }
    if yname == "dot1dStpPortDesignatedBridge" { return "Dot1Dstpportdesignatedbridge" }
    if yname == "dot1dStpPortDesignatedPort" { return "Dot1Dstpportdesignatedport" }
    if yname == "dot1dStpPortForwardTransitions" { return "Dot1Dstpportforwardtransitions" }
    if yname == "dot1dStpPortPathCost32" { return "Dot1Dstpportpathcost32" }
    if yname == "stpxLongStpPortPathCost" { return "Stpxlongstpportpathcost" }
    return ""
}

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetSegmentPath() string {
    return "dot1dStpPortEntry" + "[dot1dStpPort='" + fmt.Sprintf("%v", dot1Dstpportentry.Dot1Dstpport) + "']"
}

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dStpPort"] = dot1Dstpportentry.Dot1Dstpport
    leafs["dot1dStpPortPriority"] = dot1Dstpportentry.Dot1Dstpportpriority
    leafs["dot1dStpPortState"] = dot1Dstpportentry.Dot1Dstpportstate
    leafs["dot1dStpPortEnable"] = dot1Dstpportentry.Dot1Dstpportenable
    leafs["dot1dStpPortPathCost"] = dot1Dstpportentry.Dot1Dstpportpathcost
    leafs["dot1dStpPortDesignatedRoot"] = dot1Dstpportentry.Dot1Dstpportdesignatedroot
    leafs["dot1dStpPortDesignatedCost"] = dot1Dstpportentry.Dot1Dstpportdesignatedcost
    leafs["dot1dStpPortDesignatedBridge"] = dot1Dstpportentry.Dot1Dstpportdesignatedbridge
    leafs["dot1dStpPortDesignatedPort"] = dot1Dstpportentry.Dot1Dstpportdesignatedport
    leafs["dot1dStpPortForwardTransitions"] = dot1Dstpportentry.Dot1Dstpportforwardtransitions
    leafs["dot1dStpPortPathCost32"] = dot1Dstpportentry.Dot1Dstpportpathcost32
    leafs["stpxLongStpPortPathCost"] = dot1Dstpportentry.Stpxlongstpportpathcost
    return leafs
}

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetYangName() string { return "dot1dStpPortEntry" }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) SetParent(parent types.Entity) { dot1Dstpportentry.parent = parent }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetParent() types.Entity { return dot1Dstpportentry.parent }

func (dot1Dstpportentry *BRIDGEMIB_Dot1Dstpporttable_Dot1Dstpportentry) GetParentYangName() string { return "dot1dStpPortTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a specific unicast MAC address for which the bridge has
    // some forwarding and/or filtering information. The type is slice of
    // BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry.
    Dot1Dtpfdbentry []BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry
}

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetFilter() yfilter.YFilter { return dot1Dtpfdbtable.YFilter }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) SetFilter(yf yfilter.YFilter) { dot1Dtpfdbtable.YFilter = yf }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetGoName(yname string) string {
    if yname == "dot1dTpFdbEntry" { return "Dot1Dtpfdbentry" }
    return ""
}

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetSegmentPath() string {
    return "dot1dTpFdbTable"
}

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dTpFdbEntry" {
        for _, c := range dot1Dtpfdbtable.Dot1Dtpfdbentry {
            if dot1Dtpfdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry{}
        dot1Dtpfdbtable.Dot1Dtpfdbentry = append(dot1Dtpfdbtable.Dot1Dtpfdbentry, child)
        return &dot1Dtpfdbtable.Dot1Dtpfdbentry[len(dot1Dtpfdbtable.Dot1Dtpfdbentry)-1]
    }
    return nil
}

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dtpfdbtable.Dot1Dtpfdbentry {
        children[dot1Dtpfdbtable.Dot1Dtpfdbentry[i].GetSegmentPath()] = &dot1Dtpfdbtable.Dot1Dtpfdbentry[i]
    }
    return children
}

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetYangName() string { return "dot1dTpFdbTable" }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) SetParent(parent types.Entity) { dot1Dtpfdbtable.parent = parent }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetParent() types.Entity { return dot1Dtpfdbtable.parent }

func (dot1Dtpfdbtable *BRIDGEMIB_Dot1Dtpfdbtable) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry
// Information about a specific unicast MAC address
// for which the bridge has some forwarding and/or
// filtering information.
type BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A unicast MAC address for which the bridge has
    // forwarding and/or filtering information. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetFilter() yfilter.YFilter { return dot1Dtpfdbentry.YFilter }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) SetFilter(yf yfilter.YFilter) { dot1Dtpfdbentry.YFilter = yf }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetGoName(yname string) string {
    if yname == "dot1dTpFdbAddress" { return "Dot1Dtpfdbaddress" }
    if yname == "dot1dTpFdbPort" { return "Dot1Dtpfdbport" }
    if yname == "dot1dTpFdbStatus" { return "Dot1Dtpfdbstatus" }
    return ""
}

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetSegmentPath() string {
    return "dot1dTpFdbEntry" + "[dot1dTpFdbAddress='" + fmt.Sprintf("%v", dot1Dtpfdbentry.Dot1Dtpfdbaddress) + "']"
}

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dTpFdbAddress"] = dot1Dtpfdbentry.Dot1Dtpfdbaddress
    leafs["dot1dTpFdbPort"] = dot1Dtpfdbentry.Dot1Dtpfdbport
    leafs["dot1dTpFdbStatus"] = dot1Dtpfdbentry.Dot1Dtpfdbstatus
    return leafs
}

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetYangName() string { return "dot1dTpFdbEntry" }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) SetParent(parent types.Entity) { dot1Dtpfdbentry.parent = parent }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetParent() types.Entity { return dot1Dtpfdbentry.parent }

func (dot1Dtpfdbentry *BRIDGEMIB_Dot1Dtpfdbtable_Dot1Dtpfdbentry) GetParentYangName() string { return "dot1dTpFdbTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of information for each port of a transparent bridge. The type is
    // slice of BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry.
    Dot1Dtpportentry []BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry
}

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetFilter() yfilter.YFilter { return dot1Dtpporttable.YFilter }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) SetFilter(yf yfilter.YFilter) { dot1Dtpporttable.YFilter = yf }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetGoName(yname string) string {
    if yname == "dot1dTpPortEntry" { return "Dot1Dtpportentry" }
    return ""
}

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetSegmentPath() string {
    return "dot1dTpPortTable"
}

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dTpPortEntry" {
        for _, c := range dot1Dtpporttable.Dot1Dtpportentry {
            if dot1Dtpporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry{}
        dot1Dtpporttable.Dot1Dtpportentry = append(dot1Dtpporttable.Dot1Dtpportentry, child)
        return &dot1Dtpporttable.Dot1Dtpportentry[len(dot1Dtpporttable.Dot1Dtpportentry)-1]
    }
    return nil
}

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dtpporttable.Dot1Dtpportentry {
        children[dot1Dtpporttable.Dot1Dtpportentry[i].GetSegmentPath()] = &dot1Dtpporttable.Dot1Dtpportentry[i]
    }
    return children
}

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetYangName() string { return "dot1dTpPortTable" }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) SetParent(parent types.Entity) { dot1Dtpporttable.parent = parent }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetParent() types.Entity { return dot1Dtpporttable.parent }

func (dot1Dtpporttable *BRIDGEMIB_Dot1Dtpporttable) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry
// A list of information for each port of a transparent
// bridge.
type BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry struct {
    parent types.Entity
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

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetFilter() yfilter.YFilter { return dot1Dtpportentry.YFilter }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) SetFilter(yf yfilter.YFilter) { dot1Dtpportentry.YFilter = yf }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetGoName(yname string) string {
    if yname == "dot1dTpPort" { return "Dot1Dtpport" }
    if yname == "dot1dTpPortMaxInfo" { return "Dot1Dtpportmaxinfo" }
    if yname == "dot1dTpPortInFrames" { return "Dot1Dtpportinframes" }
    if yname == "dot1dTpPortOutFrames" { return "Dot1Dtpportoutframes" }
    if yname == "dot1dTpPortInDiscards" { return "Dot1Dtpportindiscards" }
    return ""
}

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetSegmentPath() string {
    return "dot1dTpPortEntry" + "[dot1dTpPort='" + fmt.Sprintf("%v", dot1Dtpportentry.Dot1Dtpport) + "']"
}

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dTpPort"] = dot1Dtpportentry.Dot1Dtpport
    leafs["dot1dTpPortMaxInfo"] = dot1Dtpportentry.Dot1Dtpportmaxinfo
    leafs["dot1dTpPortInFrames"] = dot1Dtpportentry.Dot1Dtpportinframes
    leafs["dot1dTpPortOutFrames"] = dot1Dtpportentry.Dot1Dtpportoutframes
    leafs["dot1dTpPortInDiscards"] = dot1Dtpportentry.Dot1Dtpportindiscards
    return leafs
}

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetYangName() string { return "dot1dTpPortEntry" }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) SetParent(parent types.Entity) { dot1Dtpportentry.parent = parent }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetParent() types.Entity { return dot1Dtpportentry.parent }

func (dot1Dtpportentry *BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry) GetParentYangName() string { return "dot1dTpPortTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Filtering information configured into the bridge by (local or network)
    // management specifying the set of ports to which frames received from a
    // specific port and containing a specific destination address are allowed to
    // be forwarded. The type is slice of
    // BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry.
    Dot1Dstaticentry []BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry
}

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetFilter() yfilter.YFilter { return dot1Dstatictable.YFilter }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) SetFilter(yf yfilter.YFilter) { dot1Dstatictable.YFilter = yf }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetGoName(yname string) string {
    if yname == "dot1dStaticEntry" { return "Dot1Dstaticentry" }
    return ""
}

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetSegmentPath() string {
    return "dot1dStaticTable"
}

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dStaticEntry" {
        for _, c := range dot1Dstatictable.Dot1Dstaticentry {
            if dot1Dstatictable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry{}
        dot1Dstatictable.Dot1Dstaticentry = append(dot1Dstatictable.Dot1Dstaticentry, child)
        return &dot1Dstatictable.Dot1Dstaticentry[len(dot1Dstatictable.Dot1Dstaticentry)-1]
    }
    return nil
}

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dstatictable.Dot1Dstaticentry {
        children[dot1Dstatictable.Dot1Dstaticentry[i].GetSegmentPath()] = &dot1Dstatictable.Dot1Dstaticentry[i]
    }
    return children
}

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetYangName() string { return "dot1dStaticTable" }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) SetParent(parent types.Entity) { dot1Dstatictable.parent = parent }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetParent() types.Entity { return dot1Dstatictable.parent }

func (dot1Dstatictable *BRIDGEMIB_Dot1Dstatictable) GetParentYangName() string { return "BRIDGE-MIB" }

// BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry
// Filtering information configured into the bridge by
// (local or network) management specifying the set of
// ports to which frames received from a specific port and
// containing a specific destination address are allowed to
// be forwarded.
type BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object can take the value
    // of a unicast address, a group address, or the broadcast address. The type
    // is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetFilter() yfilter.YFilter { return dot1Dstaticentry.YFilter }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) SetFilter(yf yfilter.YFilter) { dot1Dstaticentry.YFilter = yf }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetGoName(yname string) string {
    if yname == "dot1dStaticAddress" { return "Dot1Dstaticaddress" }
    if yname == "dot1dStaticReceivePort" { return "Dot1Dstaticreceiveport" }
    if yname == "dot1dStaticAllowedToGoTo" { return "Dot1Dstaticallowedtogoto" }
    if yname == "dot1dStaticStatus" { return "Dot1Dstaticstatus" }
    return ""
}

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetSegmentPath() string {
    return "dot1dStaticEntry" + "[dot1dStaticAddress='" + fmt.Sprintf("%v", dot1Dstaticentry.Dot1Dstaticaddress) + "']" + "[dot1dStaticReceivePort='" + fmt.Sprintf("%v", dot1Dstaticentry.Dot1Dstaticreceiveport) + "']"
}

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dStaticAddress"] = dot1Dstaticentry.Dot1Dstaticaddress
    leafs["dot1dStaticReceivePort"] = dot1Dstaticentry.Dot1Dstaticreceiveport
    leafs["dot1dStaticAllowedToGoTo"] = dot1Dstaticentry.Dot1Dstaticallowedtogoto
    leafs["dot1dStaticStatus"] = dot1Dstaticentry.Dot1Dstaticstatus
    return leafs
}

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetYangName() string { return "dot1dStaticEntry" }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) SetParent(parent types.Entity) { dot1Dstaticentry.parent = parent }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetParent() types.Entity { return dot1Dstaticentry.parent }

func (dot1Dstaticentry *BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry) GetParentYangName() string { return "dot1dStaticTable" }

// BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus represents         and will remain so until it is aged out.
type BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus string

const (
    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_other BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "other"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_invalid BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "invalid"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_permanent BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "permanent"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_deleteOnReset BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "deleteOnReset"

    BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus_deleteOnTimeout BRIDGEMIB_Dot1Dstatictable_Dot1Dstaticentry_Dot1Dstaticstatus = "deleteOnTimeout"
)

