// Copyright (C) The Internet Society (2004). The
// initial version of this MIB module was published
// 
// 
// in RFC 3815. For full legal notices see the RFC
// itself or see:
// http://www.ietf.org/copyrights/ianamib.html
// 
// This MIB contains managed object definitions for the
// 'Multiprotocol Label Switching, Label Distribution
// Protocol, LDP' document.
package mpls_ldp_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_ldp_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:MPLS-LDP-STD-MIB MPLS-LDP-STD-MIB}", reflect.TypeOf(MPLSLDPSTDMIB{}))
    ydk.RegisterEntity("MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB", reflect.TypeOf(MPLSLDPSTDMIB{}))
}

// MPLSLDPSTDMIB
type MPLSLDPSTDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Mplsldplsrobjects MPLSLDPSTDMIB_Mplsldplsrobjects

    
    Mplsldpentityobjects MPLSLDPSTDMIB_Mplsldpentityobjects

    
    Mplsldpsessionobjects MPLSLDPSTDMIB_Mplsldpsessionobjects

    
    Mplsfecobjects MPLSLDPSTDMIB_Mplsfecobjects

    // This table contains information about the MPLS Label Distribution Protocol
    // Entities which exist on this Label Switching Router (LSR) or Label Edge
    // Router (LER).
    Mplsldpentitytable MPLSLDPSTDMIB_Mplsldpentitytable

    // Information about LDP peers known by Entities in the mplsLdpEntityTable. 
    // The information in this table is based on information from the Entity-Peer
    // interaction during session initialization but is not appropriate for the
    // mplsLdpSessionTable, because objects in this table may or may not be used
    // in session establishment.
    Mplsldppeertable MPLSLDPSTDMIB_Mplsldppeertable

    // A table of Hello Adjacencies for Sessions.
    Mplsldphelloadjacencytable MPLSLDPSTDMIB_Mplsldphelloadjacencytable

    // A table of LDP LSP's which map to the mplsInSegmentTable in the
    // MPLS-LSR-STD-MIB module.
    Mplsinsegmentldplsptable MPLSLDPSTDMIB_Mplsinsegmentldplsptable

    // A table of LDP LSP's which map to the mplsOutSegmentTable in the
    // MPLS-LSR-STD-MIB.
    Mplsoutsegmentldplsptable MPLSLDPSTDMIB_Mplsoutsegmentldplsptable

    // This table represents the FEC (Forwarding Equivalence Class) Information
    // associated with an LSP.
    Mplsfectable MPLSLDPSTDMIB_Mplsfectable

    // A table which shows the relationship between LDP LSPs and FECs.  Each row
    // represents a single LDP LSP to FEC association.
    Mplsldplspfectable MPLSLDPSTDMIB_Mplsldplspfectable

    // This table 'extends' the mplsLdpSessionTable. This table is used to store
    // Label Address Information from Label Address Messages received by this LSR
    // from Peers.  This table is read-only and should be updated   when Label
    // Withdraw Address Messages are received, i.e., Rows should be deleted as
    // appropriate.  NOTE:  since more than one address may be contained in a
    // Label Address Message, this table 'sparse augments', the
    // mplsLdpSessionTable's information.
    Mplsldpsessionpeeraddrtable MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable
}

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSLDPSTDMIB.EntityData.YFilter = mPLSLDPSTDMIB.YFilter
    mPLSLDPSTDMIB.EntityData.YangName = "MPLS-LDP-STD-MIB"
    mPLSLDPSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSLDPSTDMIB.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mPLSLDPSTDMIB.EntityData.SegmentPath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB"
    mPLSLDPSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSLDPSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSLDPSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSLDPSTDMIB.EntityData.Children = make(map[string]types.YChild)
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpLsrObjects"] = types.YChild{"Mplsldplsrobjects", &mPLSLDPSTDMIB.Mplsldplsrobjects}
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpEntityObjects"] = types.YChild{"Mplsldpentityobjects", &mPLSLDPSTDMIB.Mplsldpentityobjects}
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpSessionObjects"] = types.YChild{"Mplsldpsessionobjects", &mPLSLDPSTDMIB.Mplsldpsessionobjects}
    mPLSLDPSTDMIB.EntityData.Children["mplsFecObjects"] = types.YChild{"Mplsfecobjects", &mPLSLDPSTDMIB.Mplsfecobjects}
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpEntityTable"] = types.YChild{"Mplsldpentitytable", &mPLSLDPSTDMIB.Mplsldpentitytable}
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpPeerTable"] = types.YChild{"Mplsldppeertable", &mPLSLDPSTDMIB.Mplsldppeertable}
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpHelloAdjacencyTable"] = types.YChild{"Mplsldphelloadjacencytable", &mPLSLDPSTDMIB.Mplsldphelloadjacencytable}
    mPLSLDPSTDMIB.EntityData.Children["mplsInSegmentLdpLspTable"] = types.YChild{"Mplsinsegmentldplsptable", &mPLSLDPSTDMIB.Mplsinsegmentldplsptable}
    mPLSLDPSTDMIB.EntityData.Children["mplsOutSegmentLdpLspTable"] = types.YChild{"Mplsoutsegmentldplsptable", &mPLSLDPSTDMIB.Mplsoutsegmentldplsptable}
    mPLSLDPSTDMIB.EntityData.Children["mplsFecTable"] = types.YChild{"Mplsfectable", &mPLSLDPSTDMIB.Mplsfectable}
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpLspFecTable"] = types.YChild{"Mplsldplspfectable", &mPLSLDPSTDMIB.Mplsldplspfectable}
    mPLSLDPSTDMIB.EntityData.Children["mplsLdpSessionPeerAddrTable"] = types.YChild{"Mplsldpsessionpeeraddrtable", &mPLSLDPSTDMIB.Mplsldpsessionpeeraddrtable}
    mPLSLDPSTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mPLSLDPSTDMIB.EntityData)
}

// MPLSLDPSTDMIB_Mplsldplsrobjects
type MPLSLDPSTDMIB_Mplsldplsrobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Label Switching Router's Identifier. The type is string with length: 4.
    Mplsldplsrid interface{}

    // A indication of whether this Label Switching Router supports loop
    // detection.  none(1) -- Loop Detection is not supported            on this
    // LSR.  other(2) -- Loop Detection is supported but             by a method
    // other than those             listed below.  hopCount(3) -- Loop Detection
    // is supported by                Hop Count only.  pathVector(4) -- Loop
    // Detection is supported by                  Path Vector only. 
    // hopCountAndPathVector(5) -- Loop Detection is                     
    // supported by both Hop Count                      And Path Vector.  Since
    // Loop Detection is determined during Session Initialization, an individual
    // session may not be running with loop detection.  This object simply gives
    // an indication of whether or not the LSR has the ability to support Loop
    // Detection and which types. The type is Mplsldplsrloopdetectioncapable.
    Mplsldplsrloopdetectioncapable interface{}
}

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetEntityData() *types.CommonEntityData {
    mplsldplsrobjects.EntityData.YFilter = mplsldplsrobjects.YFilter
    mplsldplsrobjects.EntityData.YangName = "mplsLdpLsrObjects"
    mplsldplsrobjects.EntityData.BundleName = "cisco_ios_xe"
    mplsldplsrobjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldplsrobjects.EntityData.SegmentPath = "mplsLdpLsrObjects"
    mplsldplsrobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldplsrobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldplsrobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldplsrobjects.EntityData.Children = make(map[string]types.YChild)
    mplsldplsrobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldplsrobjects.EntityData.Leafs["mplsLdpLsrId"] = types.YLeaf{"Mplsldplsrid", mplsldplsrobjects.Mplsldplsrid}
    mplsldplsrobjects.EntityData.Leafs["mplsLdpLsrLoopDetectionCapable"] = types.YLeaf{"Mplsldplsrloopdetectioncapable", mplsldplsrobjects.Mplsldplsrloopdetectioncapable}
    return &(mplsldplsrobjects.EntityData)
}

// MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable represents which types.
type MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable string

const (
    MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable_none MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable = "none"

    MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable_other MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable = "other"

    MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable_hopCount MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable = "hopCount"

    MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable_pathVector MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable = "pathVector"

    MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable_hopCountAndPathVector MPLSLDPSTDMIB_Mplsldplsrobjects_Mplsldplsrloopdetectioncapable = "hopCountAndPathVector"
)

// MPLSLDPSTDMIB_Mplsldpentityobjects
type MPLSLDPSTDMIB_Mplsldpentityobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the most recent addition or deletion
    // of an entry to/from the mplsLdpEntityTable/mplsLdpEntityStatsTable, or the
    // most recent change in value of any objects in the mplsLdpEntityTable.   If
    // no such changes have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Mplsldpentitylastchange interface{}

    // This object contains an appropriate value to be used for mplsLdpEntityIndex
    // when creating entries in the mplsLdpEntityTable. The value 0 indicates that
    // no unassigned entries are available. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentityindexnext interface{}
}

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetEntityData() *types.CommonEntityData {
    mplsldpentityobjects.EntityData.YFilter = mplsldpentityobjects.YFilter
    mplsldpentityobjects.EntityData.YangName = "mplsLdpEntityObjects"
    mplsldpentityobjects.EntityData.BundleName = "cisco_ios_xe"
    mplsldpentityobjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldpentityobjects.EntityData.SegmentPath = "mplsLdpEntityObjects"
    mplsldpentityobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpentityobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpentityobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpentityobjects.EntityData.Children = make(map[string]types.YChild)
    mplsldpentityobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldpentityobjects.EntityData.Leafs["mplsLdpEntityLastChange"] = types.YLeaf{"Mplsldpentitylastchange", mplsldpentityobjects.Mplsldpentitylastchange}
    mplsldpentityobjects.EntityData.Leafs["mplsLdpEntityIndexNext"] = types.YLeaf{"Mplsldpentityindexnext", mplsldpentityobjects.Mplsldpentityindexnext}
    return &(mplsldpentityobjects.EntityData)
}

// MPLSLDPSTDMIB_Mplsldpsessionobjects
type MPLSLDPSTDMIB_Mplsldpsessionobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the most recent addition or deletion
    // to/from the mplsLdpPeerTable/mplsLdpSessionTable. The type is interface{}
    // with range: 0..4294967295.
    Mplsldppeerlastchange interface{}

    // The value of sysUpTime at the time of the most recent addition/deletion of
    // an entry to/from the mplsLdpLspFecTable or the most recent change in values
    // to any objects in the mplsLdpLspFecTable.  If no such changes have occurred
    // since the last re-initialization of the local management subsystem, then
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    Mplsldplspfeclastchange interface{}
}

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetEntityData() *types.CommonEntityData {
    mplsldpsessionobjects.EntityData.YFilter = mplsldpsessionobjects.YFilter
    mplsldpsessionobjects.EntityData.YangName = "mplsLdpSessionObjects"
    mplsldpsessionobjects.EntityData.BundleName = "cisco_ios_xe"
    mplsldpsessionobjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldpsessionobjects.EntityData.SegmentPath = "mplsLdpSessionObjects"
    mplsldpsessionobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpsessionobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpsessionobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpsessionobjects.EntityData.Children = make(map[string]types.YChild)
    mplsldpsessionobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldpsessionobjects.EntityData.Leafs["mplsLdpPeerLastChange"] = types.YLeaf{"Mplsldppeerlastchange", mplsldpsessionobjects.Mplsldppeerlastchange}
    mplsldpsessionobjects.EntityData.Leafs["mplsLdpLspFecLastChange"] = types.YLeaf{"Mplsldplspfeclastchange", mplsldpsessionobjects.Mplsldplspfeclastchange}
    return &(mplsldpsessionobjects.EntityData)
}

// MPLSLDPSTDMIB_Mplsfecobjects
type MPLSLDPSTDMIB_Mplsfecobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the most recent addition/deletion of
    // an entry to/from the mplsLdpFectTable or the most recent change in values
    // to any objects in the mplsLdpFecTable.  If no such changes have occurred
    // since the last re-initialization of the local management subsystem, then
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    Mplsfeclastchange interface{}

    // This object contains an appropriate value to be used for mplsFecIndex when
    // creating entries in the mplsFecTable. The value 0 indicates that no
    // unassigned entries are available. The type is interface{} with range:
    // 0..4294967295.
    Mplsfecindexnext interface{}
}

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetEntityData() *types.CommonEntityData {
    mplsfecobjects.EntityData.YFilter = mplsfecobjects.YFilter
    mplsfecobjects.EntityData.YangName = "mplsFecObjects"
    mplsfecobjects.EntityData.BundleName = "cisco_ios_xe"
    mplsfecobjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsfecobjects.EntityData.SegmentPath = "mplsFecObjects"
    mplsfecobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsfecobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsfecobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsfecobjects.EntityData.Children = make(map[string]types.YChild)
    mplsfecobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsfecobjects.EntityData.Leafs["mplsFecLastChange"] = types.YLeaf{"Mplsfeclastchange", mplsfecobjects.Mplsfeclastchange}
    mplsfecobjects.EntityData.Leafs["mplsFecIndexNext"] = types.YLeaf{"Mplsfecindexnext", mplsfecobjects.Mplsfecindexnext}
    return &(mplsfecobjects.EntityData)
}

// MPLSLDPSTDMIB_Mplsldpentitytable
// This table contains information about the
// MPLS Label Distribution Protocol Entities which
// exist on this Label Switching Router (LSR)
// or Label Edge Router (LER).
type MPLSLDPSTDMIB_Mplsldpentitytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents an LDP entity. An entry can be created by
    // a network administrator or by an SNMP agent as instructed by LDP. The type
    // is slice of MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry.
    Mplsldpentityentry []MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry
}

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetEntityData() *types.CommonEntityData {
    mplsldpentitytable.EntityData.YFilter = mplsldpentitytable.YFilter
    mplsldpentitytable.EntityData.YangName = "mplsLdpEntityTable"
    mplsldpentitytable.EntityData.BundleName = "cisco_ios_xe"
    mplsldpentitytable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldpentitytable.EntityData.SegmentPath = "mplsLdpEntityTable"
    mplsldpentitytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpentitytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpentitytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpentitytable.EntityData.Children = make(map[string]types.YChild)
    mplsldpentitytable.EntityData.Children["mplsLdpEntityEntry"] = types.YChild{"Mplsldpentityentry", nil}
    for i := range mplsldpentitytable.Mplsldpentityentry {
        mplsldpentitytable.EntityData.Children[types.GetSegmentPath(&mplsldpentitytable.Mplsldpentityentry[i])] = types.YChild{"Mplsldpentityentry", &mplsldpentitytable.Mplsldpentityentry[i]}
    }
    mplsldpentitytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsldpentitytable.EntityData)
}

// MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry
// An entry in this table represents an LDP entity.
// An entry can be created by a network administrator
// or by an SNMP agent as instructed by LDP.
type MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The LDP identifier. The type is string.
    Mplsldpentityldpid interface{}

    // This attribute is a key. This index is used as a secondary index to
    // uniquely identify this row.  Before creating a row in this table, the
    // 'mplsLdpEntityIndexNext' object should be retrieved. That value should be
    // used for the value of this index when creating a row in this table.  NOTE: 
    // if a value of zero (0) is retrieved, that indicates that no rows can be
    // created in this table at this time.  A secondary index (this object) is
    // meaningful to some but not all, LDP implementations.  For example an LDP
    // implementation which uses PPP would use this index to differentiate PPP
    // sub-links.  Another way to use this index is to give this the value of
    // ifIndex.  However, this is dependant   on the implementation. The type is
    // interface{} with range: 1..4294967295.
    Mplsldpentityindex interface{}

    // The version number of the LDP protocol which will be used in the session
    // initialization message.  Section 3.5.3 in the LDP Specification specifies
    // that the version of the LDP protocol is negotiated during session
    // establishment. The value of this object represents the value that is sent
    // in the initialization message. The type is interface{} with range:
    // 1..65535.
    Mplsldpentityprotocolversion interface{}

    // The administrative status of this LDP Entity. If this object is changed
    // from 'enable' to 'disable' and this entity has already attempted to
    // establish contact with a Peer, then all contact with that Peer is lost and
    // all information from that Peer needs to be removed from the MIB. (This
    // implies that the network management subsystem should clean up any related
    // entry in the mplsLdpPeerTable.  This further implies that a 'tear-down' for
    // that session is issued and the session and all information related to that
    // session cease to exist).  At this point the operator is able to change
    // values which are related to this entity.  When the admin status is set back
    // to 'enable', then this Entity will attempt to establish a new session with
    // the Peer. The type is Mplsldpentityadminstatus.
    Mplsldpentityadminstatus interface{}

    // The operational status of this LDP Entity.  The value of unknown(1)
    // indicates that the operational status cannot be determined at this time. 
    // The value of unknown should be a transient condition before changing to
    // enabled(2) or disabled(3). The type is Mplsldpentityoperstatus.
    Mplsldpentityoperstatus interface{}

    // The TCP Port for LDP.  The default value is the well-known value of this
    // port. The type is interface{} with range: 0..65535.
    Mplsldpentitytcpport interface{}

    // The UDP Discovery Port for LDP.  The default value is the well-known value
    // for this port. The type is interface{} with range: 0..65535.
    Mplsldpentityudpdscport interface{}

    // The maximum PDU Length that is sent in the Common Session Parameters of an
    // Initialization Message. According to the LDP Specification [RFC3036] a
    // value of 255 or less specifies the default maximum length of 4096 octets,
    // this is why the value of this object starts at 256.  The operator should
    // explicitly choose the default value (i.e., 4096), or some other value.  The
    // receiving LSR MUST calculate the maximum PDU length for the session by
    // using the smaller of its and its peer's proposals for Max PDU Length. The
    // type is interface{} with range: 256..65535. Units are octets.
    Mplsldpentitymaxpdulength interface{}

    // The 16-bit integer value which is the proposed keep alive hold timer for
    // this LDP Entity. The type is interface{} with range: 1..65535. Units are
    // seconds.
    Mplsldpentitykeepaliveholdtimer interface{}

    // The 16-bit integer value which is the proposed Hello hold timer for this
    // LDP Entity. The Hello Hold time in seconds.   An LSR maintains a record of
    // Hellos received from potential peers.  This object represents the Hold Time
    // in the Common Hello Parameters TLV of the Hello Message.  A value of 0 is a
    // default value and should be interpretted in conjunction with the
    // mplsLdpEntityTargetPeer object.  If the value of this object is 0: if the
    // value of the mplsLdpEntityTargetPeer object is false(2), then this
    // specifies that the Hold Time's actual default value is 15 seconds (i.e.,
    // the default Hold time for Link Hellos is 15 seconds).  Otherwise if the
    // value of the mplsLdpEntityTargetPeer object is true(1), then this specifies
    // that the Hold Time's actual default value is 45 seconds (i.e., the default
    // Hold time for Targeted Hellos is 45 seconds).  A value of 65535 means
    // infinite (i.e., wait forever).  All other values represent the amount of
    // time in seconds to wait for a Hello Message.  Setting the hold time to a
    // value smaller than 15 is not recommended, although not forbidden according
    // to RFC3036. The type is interface{} with range: 0..65535. Units are
    // seconds.
    Mplsldpentityhelloholdtimer interface{}

    // When attempting to establish a session with a given Peer, the given LDP
    // Entity should send out the SNMP notification,
    // 'mplsLdpInitSessionThresholdExceeded', when the number of Session
    // Initialization messages sent exceeds this threshold.  The notification is
    // used to notify an operator when this Entity and its Peer are possibly
    // engaged in an endless sequence of messages as each NAKs the other's  
    // Initialization messages with Error Notification messages.  Setting this
    // threshold which triggers the notification is one way to notify the
    // operator.  The notification should be generated each time this threshold is
    // exceeded and for every subsequent Initialization message which is NAK'd
    // with an Error Notification message after this threshold is exceeded.  A
    // value of 0 (zero) for this object indicates that the threshold is infinity,
    // thus the SNMP notification will never be generated. The type is interface{}
    // with range: 0..100.
    Mplsldpentityinitsessionthreshold interface{}

    // For any given LDP session, the method of label distribution must be
    // specified. The type is MplsLabelDistributionMethod.
    Mplsldpentitylabeldistmethod interface{}

    // The LDP Entity can be configured to use either conservative or liberal
    // label retention mode.  If the value of this object is conservative(1) then
    // advertized label mappings are retained only if they will be used to forward
    // packets, i.e., if label came from a valid next hop.  If the value of this
    // object is liberal(2) then all advertized label mappings are retained
    // whether they are from a valid next hop or not. The type is
    // MplsRetentionMode.
    Mplsldpentitylabelretentionmode interface{}

    // If the value of this object is 0 (zero) then Loop Detection for Path
    // Vectors is disabled.  Otherwise, if this object has a value greater than
    // zero, then Loop Dection for Path Vectors is enabled, and the Path Vector
    // Limit is this value. Also, the value of the object,
    // 'mplsLdpLsrLoopDetectionCapable', must be set to either 'pathVector(4)' or
    // 'hopCountAndPathVector(5)', if this object has a value greater than 0
    // (zero), otherwise it is ignored. The type is interface{} with range:
    // 0..255.
    Mplsldpentitypathvectorlimit interface{}

    // If the value of this object is 0 (zero), then Loop Detection using Hop
    // Counters is disabled.  If the value of this object is greater than 0 (zero)
    // then Loop Detection using Hop Counters is enabled, and this object
    // specifies this Entity's maximum allowable value for the Hop Count. Also,
    // the value of the object mplsLdpLsrLoopDetectionCapable must be set to
    // either 'hopCount(3)' or 'hopCountAndPathVector(5)' if this object has a
    // value greater than 0 (zero), otherwise it is ignored. The type is
    // interface{} with range: 0..255.
    Mplsldpentityhopcountlimit interface{}

    // This specifies whether the loopback or interface address is to be used as
    // the transport address in the transport address TLV of the hello message. 
    // If the value is interface(1), then the IP address of the interface from
    // which hello messages are sent is used as the transport address in the hello
    // message.  Otherwise, if the value is loopback(2), then the IP address of
    // the loopback interface is used as the transport address in the hello
    // message. The type is Mplsldpentitytransportaddrkind.
    Mplsldpentitytransportaddrkind interface{}

    // If this LDP entity uses targeted peer then set this to true. The type is
    // bool.
    Mplsldpentitytargetpeer interface{}

    // The type of the internetwork layer address used for the Extended Discovery.
    // This object indicates how the value of mplsLdpEntityTargetPeerAddr is to be
    // interpreted. The type is InetAddressType.
    Mplsldpentitytargetpeeraddrtype interface{}

    // The value of the internetwork layer address used for the Extended
    // Discovery.  The value of mplsLdpEntityTargetPeerAddrType specifies how this
    // address is to be interpreted. The type is string with length: 0..255.
    Mplsldpentitytargetpeeraddr interface{}

    // Specifies the optional parameters for the LDP Initialization Message.  If
    // the value is generic(1) then no optional parameters will be sent in the LDP
    // Initialization message associated with this Entity.  If the value is
    // atmParameters(2) then a row must be created in the mplsLdpEntityAtmTable,
    // which corresponds to this entry.  If the value is frameRelayParameters(3)
    // then a row must be created in the mplsLdpEntityFrameRelayTable, which
    // corresponds to this entry. The type is MplsLdpLabelType.
    Mplsldpentitylabeltype interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entity's counters suffered a discontinuity.  The relevant counters
    // are the specific instances associated with this entity of any Counter32
    // object contained in the 'mplsLdpEntityStatsTable'.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Mplsldpentitydiscontinuitytime interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    Mplsldpentitystoragetype interface{}

    // The status of this conceptual row.  All writable objects in this row may be
    // modified at any time, however, as described in detail in the section
    // entitled, 'Changing Values After Session Establishment', and again
    // described in the DESCRIPTION clause of the mplsLdpEntityAdminStatus object,
    // if a session has been initiated with a Peer, changing objects in this table
    // will wreak havoc with the session and interrupt traffic.  To repeat again:
    // the recommended procedure is to set the mplsLdpEntityAdminStatus to down,
    // thereby explicitly causing a session to be torn down. Then, change objects
    // in this entry, then set the mplsLdpEntityAdminStatus to enable, which
    // enables a new session to be initiated. The type is RowStatus.
    Mplsldpentityrowstatus interface{}

    // A count of the Session Initialization messages which were sent or received
    // by this LDP Entity and were NAK'd.   In other words, this counter counts
    // the number of session initializations that failed.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatssessionattempts interface{}

    // A count of the Session Rejected/No Hello Error Notification Messages sent
    // or received by this LDP Entity.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of mplsLdpEntityDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Mplsldpentitystatssessionrejectednohelloerrors interface{}

    // A count of the Session Rejected/Parameters Advertisement Mode Error
    // Notification Messages sent or received by this LDP Entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatssessionrejectedaderrors interface{}

    // A count of the Session Rejected/Parameters  Max Pdu Length Error
    // Notification Messages sent or received by this LDP Entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatssessionrejectedmaxpduerrors interface{}

    // A count of the Session Rejected/Parameters Label Range Notification
    // Messages sent or received by this LDP Entity.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatssessionrejectedlrerrors interface{}

    // This object counts the number of Bad LDP Identifier Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatsbadldpidentifiererrors interface{}

    // This object counts the number of Bad PDU Length Fatal Errors detected by
    // the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatsbadpdulengtherrors interface{}

    // This object counts the number of Bad Message Length Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatsbadmessagelengtherrors interface{}

    // This object counts the number of Bad TLV Length Fatal Errors detected by
    // the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatsbadtlvlengtherrors interface{}

    // This object counts the number of Malformed TLV Value Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatsmalformedtlvvalueerrors interface{}

    // This object counts the number of Session Keep Alive Timer Expired Errors
    // detected by the session(s) (past and present) associated with this LDP
    // Entity.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of mplsLdpEntityDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Mplsldpentitystatskeepalivetimerexperrors interface{}

    // This object counts the number of Shutdown Notifications received related to
    // session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatsshutdownreceivednotifications interface{}

    // This object counts the number of Shutdown Notfications sent related to
    // session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of  
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsldpentitystatsshutdownsentnotifications interface{}
}

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetEntityData() *types.CommonEntityData {
    mplsldpentityentry.EntityData.YFilter = mplsldpentityentry.YFilter
    mplsldpentityentry.EntityData.YangName = "mplsLdpEntityEntry"
    mplsldpentityentry.EntityData.BundleName = "cisco_ios_xe"
    mplsldpentityentry.EntityData.ParentYangName = "mplsLdpEntityTable"
    mplsldpentityentry.EntityData.SegmentPath = "mplsLdpEntityEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldpentityentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldpentityentry.Mplsldpentityindex) + "']"
    mplsldpentityentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpentityentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpentityentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpentityentry.EntityData.Children = make(map[string]types.YChild)
    mplsldpentityentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsldpentityentry.Mplsldpentityldpid}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsldpentityentry.Mplsldpentityindex}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityProtocolVersion"] = types.YLeaf{"Mplsldpentityprotocolversion", mplsldpentityentry.Mplsldpentityprotocolversion}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityAdminStatus"] = types.YLeaf{"Mplsldpentityadminstatus", mplsldpentityentry.Mplsldpentityadminstatus}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityOperStatus"] = types.YLeaf{"Mplsldpentityoperstatus", mplsldpentityentry.Mplsldpentityoperstatus}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityTcpPort"] = types.YLeaf{"Mplsldpentitytcpport", mplsldpentityentry.Mplsldpentitytcpport}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityUdpDscPort"] = types.YLeaf{"Mplsldpentityudpdscport", mplsldpentityentry.Mplsldpentityudpdscport}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityMaxPduLength"] = types.YLeaf{"Mplsldpentitymaxpdulength", mplsldpentityentry.Mplsldpentitymaxpdulength}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityKeepAliveHoldTimer"] = types.YLeaf{"Mplsldpentitykeepaliveholdtimer", mplsldpentityentry.Mplsldpentitykeepaliveholdtimer}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityHelloHoldTimer"] = types.YLeaf{"Mplsldpentityhelloholdtimer", mplsldpentityentry.Mplsldpentityhelloholdtimer}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityInitSessionThreshold"] = types.YLeaf{"Mplsldpentityinitsessionthreshold", mplsldpentityentry.Mplsldpentityinitsessionthreshold}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityLabelDistMethod"] = types.YLeaf{"Mplsldpentitylabeldistmethod", mplsldpentityentry.Mplsldpentitylabeldistmethod}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityLabelRetentionMode"] = types.YLeaf{"Mplsldpentitylabelretentionmode", mplsldpentityentry.Mplsldpentitylabelretentionmode}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityPathVectorLimit"] = types.YLeaf{"Mplsldpentitypathvectorlimit", mplsldpentityentry.Mplsldpentitypathvectorlimit}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityHopCountLimit"] = types.YLeaf{"Mplsldpentityhopcountlimit", mplsldpentityentry.Mplsldpentityhopcountlimit}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityTransportAddrKind"] = types.YLeaf{"Mplsldpentitytransportaddrkind", mplsldpentityentry.Mplsldpentitytransportaddrkind}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityTargetPeer"] = types.YLeaf{"Mplsldpentitytargetpeer", mplsldpentityentry.Mplsldpentitytargetpeer}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityTargetPeerAddrType"] = types.YLeaf{"Mplsldpentitytargetpeeraddrtype", mplsldpentityentry.Mplsldpentitytargetpeeraddrtype}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityTargetPeerAddr"] = types.YLeaf{"Mplsldpentitytargetpeeraddr", mplsldpentityentry.Mplsldpentitytargetpeeraddr}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityLabelType"] = types.YLeaf{"Mplsldpentitylabeltype", mplsldpentityentry.Mplsldpentitylabeltype}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityDiscontinuityTime"] = types.YLeaf{"Mplsldpentitydiscontinuitytime", mplsldpentityentry.Mplsldpentitydiscontinuitytime}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStorageType"] = types.YLeaf{"Mplsldpentitystoragetype", mplsldpentityentry.Mplsldpentitystoragetype}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityRowStatus"] = types.YLeaf{"Mplsldpentityrowstatus", mplsldpentityentry.Mplsldpentityrowstatus}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsSessionAttempts"] = types.YLeaf{"Mplsldpentitystatssessionattempts", mplsldpentityentry.Mplsldpentitystatssessionattempts}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsSessionRejectedNoHelloErrors"] = types.YLeaf{"Mplsldpentitystatssessionrejectednohelloerrors", mplsldpentityentry.Mplsldpentitystatssessionrejectednohelloerrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsSessionRejectedAdErrors"] = types.YLeaf{"Mplsldpentitystatssessionrejectedaderrors", mplsldpentityentry.Mplsldpentitystatssessionrejectedaderrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsSessionRejectedMaxPduErrors"] = types.YLeaf{"Mplsldpentitystatssessionrejectedmaxpduerrors", mplsldpentityentry.Mplsldpentitystatssessionrejectedmaxpduerrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsSessionRejectedLRErrors"] = types.YLeaf{"Mplsldpentitystatssessionrejectedlrerrors", mplsldpentityentry.Mplsldpentitystatssessionrejectedlrerrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsBadLdpIdentifierErrors"] = types.YLeaf{"Mplsldpentitystatsbadldpidentifiererrors", mplsldpentityentry.Mplsldpentitystatsbadldpidentifiererrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsBadPduLengthErrors"] = types.YLeaf{"Mplsldpentitystatsbadpdulengtherrors", mplsldpentityentry.Mplsldpentitystatsbadpdulengtherrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsBadMessageLengthErrors"] = types.YLeaf{"Mplsldpentitystatsbadmessagelengtherrors", mplsldpentityentry.Mplsldpentitystatsbadmessagelengtherrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsBadTlvLengthErrors"] = types.YLeaf{"Mplsldpentitystatsbadtlvlengtherrors", mplsldpentityentry.Mplsldpentitystatsbadtlvlengtherrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsMalformedTlvValueErrors"] = types.YLeaf{"Mplsldpentitystatsmalformedtlvvalueerrors", mplsldpentityentry.Mplsldpentitystatsmalformedtlvvalueerrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsKeepAliveTimerExpErrors"] = types.YLeaf{"Mplsldpentitystatskeepalivetimerexperrors", mplsldpentityentry.Mplsldpentitystatskeepalivetimerexperrors}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsShutdownReceivedNotifications"] = types.YLeaf{"Mplsldpentitystatsshutdownreceivednotifications", mplsldpentityentry.Mplsldpentitystatsshutdownreceivednotifications}
    mplsldpentityentry.EntityData.Leafs["mplsLdpEntityStatsShutdownSentNotifications"] = types.YLeaf{"Mplsldpentitystatsshutdownsentnotifications", mplsldpentityentry.Mplsldpentitystatsshutdownsentnotifications}
    return &(mplsldpentityentry.EntityData)
}

// MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityadminstatus represents with the Peer.
type MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityadminstatus string

const (
    MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityadminstatus_enable MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityadminstatus = "enable"

    MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityadminstatus_disable MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityadminstatus = "disable"
)

// MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus represents to enabled(2) or disabled(3).
type MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus string

const (
    MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus_unknown MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus = "unknown"

    MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus_enabled MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus = "enabled"

    MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus_disabled MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityoperstatus = "disabled"
)

// MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentitytransportaddrkind represents transport address in the hello message.
type MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentitytransportaddrkind string

const (
    MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentitytransportaddrkind_interface_ MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentitytransportaddrkind = "interface"

    MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentitytransportaddrkind_loopback MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentitytransportaddrkind = "loopback"
)

// MPLSLDPSTDMIB_Mplsldppeertable
// Information about LDP peers known by Entities in
// the mplsLdpEntityTable.  The information in this table
// is based on information from the Entity-Peer interaction
// during session initialization but is not appropriate
// for the mplsLdpSessionTable, because objects in this
// table may or may not be used in session establishment.
type MPLSLDPSTDMIB_Mplsldppeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single Peer which is related to a Session.  This table
    // is augmented by the mplsLdpSessionTable. The type is slice of
    // MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry.
    Mplsldppeerentry []MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry
}

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetEntityData() *types.CommonEntityData {
    mplsldppeertable.EntityData.YFilter = mplsldppeertable.YFilter
    mplsldppeertable.EntityData.YangName = "mplsLdpPeerTable"
    mplsldppeertable.EntityData.BundleName = "cisco_ios_xe"
    mplsldppeertable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldppeertable.EntityData.SegmentPath = "mplsLdpPeerTable"
    mplsldppeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldppeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldppeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldppeertable.EntityData.Children = make(map[string]types.YChild)
    mplsldppeertable.EntityData.Children["mplsLdpPeerEntry"] = types.YChild{"Mplsldppeerentry", nil}
    for i := range mplsldppeertable.Mplsldppeerentry {
        mplsldppeertable.EntityData.Children[types.GetSegmentPath(&mplsldppeertable.Mplsldppeerentry[i])] = types.YChild{"Mplsldppeerentry", &mplsldppeertable.Mplsldppeerentry[i]}
    }
    mplsldppeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsldppeertable.EntityData)
}

// MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry
// Information about a single Peer which is related
// to a Session.  This table is augmented by
// the mplsLdpSessionTable.
type MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityldpid
    Mplsldpentityldpid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityindex
    Mplsldpentityindex interface{}

    // This attribute is a key. The LDP identifier of this LDP Peer. The type is
    // string.
    Mplsldppeerldpid interface{}

    // For any given LDP session, the method of label distribution must be
    // specified. The type is MplsLabelDistributionMethod.
    Mplsldppeerlabeldistmethod interface{}

    // If the value of this object is 0 (zero) then Loop Dection for Path Vectors
    // for this Peer is disabled.  Otherwise, if this object has a value greater
    // than zero, then Loop Dection for Path  Vectors for this Peer is enabled and
    // the Path Vector Limit is this value. The type is interface{} with range:
    // 0..255.
    Mplsldppeerpathvectorlimit interface{}

    // The type of the Internet address for the mplsLdpPeerTransportAddr object. 
    // The LDP specification describes this as being either an IPv4 Transport
    // Address or IPv6 Transport   Address which is used in opening the LDP
    // session's TCP connection, or if the optional TLV is not present, then this
    // is the IPv4/IPv6 source address for the UPD packet carrying the Hellos. 
    // This object specifies how the value of the mplsLdpPeerTransportAddr object
    // should be interpreted. The type is InetAddressType.
    Mplsldppeertransportaddrtype interface{}

    // The Internet address advertised by the peer in the Hello Message or the
    // Hello source address.  The type of this address is specified by the value
    // of the mplsLdpPeerTransportAddrType object. The type is string with length:
    // 0..255.
    Mplsldppeertransportaddr interface{}

    // The value of sysUpTime at the time this Session entered its current state
    // as denoted by the mplsLdpSessionState object. The type is interface{} with
    // range: 0..4294967295.
    Mplsldpsessionstatelastchange interface{}

    // The current state of the session, all of the states 1 to 5 are based on the
    // state machine for session negotiation behavior. The type is
    // Mplsldpsessionstate.
    Mplsldpsessionstate interface{}

    // During session establishment the LSR/LER takes either the active role or
    // the passive role based on address comparisons.  This object indicates
    // whether this LSR/LER was behaving in an active role or passive role during
    // this session's establishment.  The value of unknown(1), indicates that the
    // role is not able to be determined at the present time. The type is
    // Mplsldpsessionrole.
    Mplsldpsessionrole interface{}

    // The version of the LDP Protocol which this session is using.  This is the
    // version of   the LDP protocol which has been negotiated during session
    // initialization. The type is interface{} with range: 1..65535.
    Mplsldpsessionprotocolversion interface{}

    // The keep alive hold time remaining for this session. The type is
    // interface{} with range: 0..2147483647.
    Mplsldpsessionkeepaliveholdtimerem interface{}

    // The negotiated KeepAlive Time which represents the amount of seconds
    // between keep alive messages.  The mplsLdpEntityKeepAliveHoldTimer related
    // to this Session is the value that was proposed as the KeepAlive Time for
    // this session.  This value is negotiated during session initialization
    // between the entity's proposed value (i.e., the value configured in
    // mplsLdpEntityKeepAliveHoldTimer) and the peer's proposed KeepAlive Hold
    // Timer value. This value is the smaller of the two proposed values. The type
    // is interface{} with range: 1..65535. Units are seconds.
    Mplsldpsessionkeepalivetime interface{}

    // The value of maximum allowable length for LDP PDUs for this session.  This
    // value may have been negotiated during the Session Initialization.  This
    // object is related to the mplsLdpEntityMaxPduLength object.  The
    // mplsLdpEntityMaxPduLength object specifies the requested LDP PDU length,
    // and this object reflects the negotiated LDP PDU length between the Entity
    // and the Peer. The type is interface{} with range: 1..65535. Units are
    // octets.
    Mplsldpsessionmaxpdulength interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this session's counters suffered a discontinuity.  The relevant counters
    // are the specific instances associated with this session of any Counter32
    // object contained in the mplsLdpSessionStatsTable.  The initial value of
    // this object is the value of sysUpTime when the entry was created in this
    // table.  Also, a command generator can distinguish when a session between a
    // given Entity and Peer goes away and a new session is established.  This
    // value would change and thus indicate to the command generator that this is
    // a different session. The type is interface{} with range: 0..4294967295.
    Mplsldpsessiondiscontinuitytime interface{}

    // This object counts the number of Unknown Message Type Errors detected by
    // this LSR/LER during this session.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of mplsLdpSessionDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Mplsldpsessionstatsunknownmestypeerrors interface{}

    // This object counts the number of Unknown TLV Errors detected by this
    // LSR/LER during this session.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of mplsLdpSessionDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Mplsldpsessionstatsunknowntlverrors interface{}
}

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetEntityData() *types.CommonEntityData {
    mplsldppeerentry.EntityData.YFilter = mplsldppeerentry.YFilter
    mplsldppeerentry.EntityData.YangName = "mplsLdpPeerEntry"
    mplsldppeerentry.EntityData.BundleName = "cisco_ios_xe"
    mplsldppeerentry.EntityData.ParentYangName = "mplsLdpPeerTable"
    mplsldppeerentry.EntityData.SegmentPath = "mplsLdpPeerEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldppeerentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldppeerentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldppeerentry.Mplsldppeerldpid) + "']"
    mplsldppeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldppeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldppeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldppeerentry.EntityData.Children = make(map[string]types.YChild)
    mplsldppeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldppeerentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsldppeerentry.Mplsldpentityldpid}
    mplsldppeerentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsldppeerentry.Mplsldpentityindex}
    mplsldppeerentry.EntityData.Leafs["mplsLdpPeerLdpId"] = types.YLeaf{"Mplsldppeerldpid", mplsldppeerentry.Mplsldppeerldpid}
    mplsldppeerentry.EntityData.Leafs["mplsLdpPeerLabelDistMethod"] = types.YLeaf{"Mplsldppeerlabeldistmethod", mplsldppeerentry.Mplsldppeerlabeldistmethod}
    mplsldppeerentry.EntityData.Leafs["mplsLdpPeerPathVectorLimit"] = types.YLeaf{"Mplsldppeerpathvectorlimit", mplsldppeerentry.Mplsldppeerpathvectorlimit}
    mplsldppeerentry.EntityData.Leafs["mplsLdpPeerTransportAddrType"] = types.YLeaf{"Mplsldppeertransportaddrtype", mplsldppeerentry.Mplsldppeertransportaddrtype}
    mplsldppeerentry.EntityData.Leafs["mplsLdpPeerTransportAddr"] = types.YLeaf{"Mplsldppeertransportaddr", mplsldppeerentry.Mplsldppeertransportaddr}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionStateLastChange"] = types.YLeaf{"Mplsldpsessionstatelastchange", mplsldppeerentry.Mplsldpsessionstatelastchange}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionState"] = types.YLeaf{"Mplsldpsessionstate", mplsldppeerentry.Mplsldpsessionstate}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionRole"] = types.YLeaf{"Mplsldpsessionrole", mplsldppeerentry.Mplsldpsessionrole}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionProtocolVersion"] = types.YLeaf{"Mplsldpsessionprotocolversion", mplsldppeerentry.Mplsldpsessionprotocolversion}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionKeepAliveHoldTimeRem"] = types.YLeaf{"Mplsldpsessionkeepaliveholdtimerem", mplsldppeerentry.Mplsldpsessionkeepaliveholdtimerem}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionKeepAliveTime"] = types.YLeaf{"Mplsldpsessionkeepalivetime", mplsldppeerentry.Mplsldpsessionkeepalivetime}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionMaxPduLength"] = types.YLeaf{"Mplsldpsessionmaxpdulength", mplsldppeerentry.Mplsldpsessionmaxpdulength}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionDiscontinuityTime"] = types.YLeaf{"Mplsldpsessiondiscontinuitytime", mplsldppeerentry.Mplsldpsessiondiscontinuitytime}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionStatsUnknownMesTypeErrors"] = types.YLeaf{"Mplsldpsessionstatsunknownmestypeerrors", mplsldppeerentry.Mplsldpsessionstatsunknownmestypeerrors}
    mplsldppeerentry.EntityData.Leafs["mplsLdpSessionStatsUnknownTlvErrors"] = types.YLeaf{"Mplsldpsessionstatsunknowntlverrors", mplsldppeerentry.Mplsldpsessionstatsunknowntlverrors}
    return &(mplsldppeerentry.EntityData)
}

// MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole represents able to be determined at the present time.
type MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole string

const (
    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole_unknown MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole = "unknown"

    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole_active MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole = "active"

    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole_passive MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionrole = "passive"
)

// MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate represents for session negotiation behavior.
type MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate string

const (
    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate_nonexistent MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate = "nonexistent"

    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate_initialized MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate = "initialized"

    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate_openrec MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate = "openrec"

    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate_opensent MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate = "opensent"

    MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate_operational MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldpsessionstate = "operational"
)

// MPLSLDPSTDMIB_Mplsldphelloadjacencytable
// A table of Hello Adjacencies for Sessions.
type MPLSLDPSTDMIB_Mplsldphelloadjacencytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row represents a single LDP Hello Adjacency. An LDP Session can have
    // one or more Hello Adjacencies. The type is slice of
    // MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry.
    Mplsldphelloadjacencyentry []MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry
}

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetEntityData() *types.CommonEntityData {
    mplsldphelloadjacencytable.EntityData.YFilter = mplsldphelloadjacencytable.YFilter
    mplsldphelloadjacencytable.EntityData.YangName = "mplsLdpHelloAdjacencyTable"
    mplsldphelloadjacencytable.EntityData.BundleName = "cisco_ios_xe"
    mplsldphelloadjacencytable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldphelloadjacencytable.EntityData.SegmentPath = "mplsLdpHelloAdjacencyTable"
    mplsldphelloadjacencytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldphelloadjacencytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldphelloadjacencytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldphelloadjacencytable.EntityData.Children = make(map[string]types.YChild)
    mplsldphelloadjacencytable.EntityData.Children["mplsLdpHelloAdjacencyEntry"] = types.YChild{"Mplsldphelloadjacencyentry", nil}
    for i := range mplsldphelloadjacencytable.Mplsldphelloadjacencyentry {
        mplsldphelloadjacencytable.EntityData.Children[types.GetSegmentPath(&mplsldphelloadjacencytable.Mplsldphelloadjacencyentry[i])] = types.YChild{"Mplsldphelloadjacencyentry", &mplsldphelloadjacencytable.Mplsldphelloadjacencyentry[i]}
    }
    mplsldphelloadjacencytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsldphelloadjacencytable.EntityData)
}

// MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry
// Each row represents a single LDP Hello Adjacency.
// An LDP Session can have one or more Hello
// Adjacencies.
type MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityldpid
    Mplsldpentityldpid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityindex
    Mplsldpentityindex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldppeerldpid
    Mplsldppeerldpid interface{}

    // This attribute is a key. An identifier for this specific adjacency. The
    // type is interface{} with range: 1..4294967295.
    Mplsldphelloadjacencyindex interface{}

    // If the value of this object is 65535, this means that the hold time is
    // infinite (i.e., wait forever).  Otherwise, the time remaining for this
    // Hello Adjacency to receive its next Hello Message.  This interval will
    // change when the 'next' Hello Message which corresponds to this Hello
    // Adjacency is received unless it is infinite. The type is interface{} with
    // range: 0..2147483647. Units are seconds.
    Mplsldphelloadjacencyholdtimerem interface{}

    // The Hello hold time which is negotiated between the Entity and the Peer. 
    // The entity associated with this Hello Adjacency issues a proposed Hello
    // Hold Time value in the mplsLdpEntityHelloHoldTimer object.  The peer also
    // proposes a value and this object represents the negotiated value.  A value
    // of 0 means the default, which is 15 seconds for Link Hellos and 45 seconds
    // for Targeted Hellos. A value of 65535 indicates an infinite hold time. The
    // type is interface{} with range: 0..65535.
    Mplsldphelloadjacencyholdtime interface{}

    // This adjacency is the result of a 'link' hello if the value of this object
    // is link(1).   Otherwise, it is a result of a 'targeted' hello, targeted(2).
    // The type is Mplsldphelloadjacencytype.
    Mplsldphelloadjacencytype interface{}
}

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetEntityData() *types.CommonEntityData {
    mplsldphelloadjacencyentry.EntityData.YFilter = mplsldphelloadjacencyentry.YFilter
    mplsldphelloadjacencyentry.EntityData.YangName = "mplsLdpHelloAdjacencyEntry"
    mplsldphelloadjacencyentry.EntityData.BundleName = "cisco_ios_xe"
    mplsldphelloadjacencyentry.EntityData.ParentYangName = "mplsLdpHelloAdjacencyTable"
    mplsldphelloadjacencyentry.EntityData.SegmentPath = "mplsLdpHelloAdjacencyEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldppeerldpid) + "']" + "[mplsLdpHelloAdjacencyIndex='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldphelloadjacencyindex) + "']"
    mplsldphelloadjacencyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldphelloadjacencyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldphelloadjacencyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldphelloadjacencyentry.EntityData.Children = make(map[string]types.YChild)
    mplsldphelloadjacencyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldphelloadjacencyentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsldphelloadjacencyentry.Mplsldpentityldpid}
    mplsldphelloadjacencyentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsldphelloadjacencyentry.Mplsldpentityindex}
    mplsldphelloadjacencyentry.EntityData.Leafs["mplsLdpPeerLdpId"] = types.YLeaf{"Mplsldppeerldpid", mplsldphelloadjacencyentry.Mplsldppeerldpid}
    mplsldphelloadjacencyentry.EntityData.Leafs["mplsLdpHelloAdjacencyIndex"] = types.YLeaf{"Mplsldphelloadjacencyindex", mplsldphelloadjacencyentry.Mplsldphelloadjacencyindex}
    mplsldphelloadjacencyentry.EntityData.Leafs["mplsLdpHelloAdjacencyHoldTimeRem"] = types.YLeaf{"Mplsldphelloadjacencyholdtimerem", mplsldphelloadjacencyentry.Mplsldphelloadjacencyholdtimerem}
    mplsldphelloadjacencyentry.EntityData.Leafs["mplsLdpHelloAdjacencyHoldTime"] = types.YLeaf{"Mplsldphelloadjacencyholdtime", mplsldphelloadjacencyentry.Mplsldphelloadjacencyholdtime}
    mplsldphelloadjacencyentry.EntityData.Leafs["mplsLdpHelloAdjacencyType"] = types.YLeaf{"Mplsldphelloadjacencytype", mplsldphelloadjacencyentry.Mplsldphelloadjacencytype}
    return &(mplsldphelloadjacencyentry.EntityData)
}

// MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry_Mplsldphelloadjacencytype represents hello, targeted(2).
type MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry_Mplsldphelloadjacencytype string

const (
    MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry_Mplsldphelloadjacencytype_link MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry_Mplsldphelloadjacencytype = "link"

    MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry_Mplsldphelloadjacencytype_targeted MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry_Mplsldphelloadjacencytype = "targeted"
)

// MPLSLDPSTDMIB_Mplsinsegmentldplsptable
// A table of LDP LSP's which
// map to the mplsInSegmentTable in the
// MPLS-LSR-STD-MIB module.
type MPLSLDPSTDMIB_Mplsinsegmentldplsptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents information on a single LDP LSP which is
    // represented by a session's index triple (mplsLdpEntityLdpId,
    // mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the index for the
    // mplsInSegmentTable (mplsInSegmentLdpLspLabelIndex) from the
    // MPLS-LSR-STD-MIB.  The information contained in a row is read-only. The
    // type is slice of
    // MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry.
    Mplsinsegmentldplspentry []MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry
}

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetEntityData() *types.CommonEntityData {
    mplsinsegmentldplsptable.EntityData.YFilter = mplsinsegmentldplsptable.YFilter
    mplsinsegmentldplsptable.EntityData.YangName = "mplsInSegmentLdpLspTable"
    mplsinsegmentldplsptable.EntityData.BundleName = "cisco_ios_xe"
    mplsinsegmentldplsptable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsinsegmentldplsptable.EntityData.SegmentPath = "mplsInSegmentLdpLspTable"
    mplsinsegmentldplsptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinsegmentldplsptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinsegmentldplsptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinsegmentldplsptable.EntityData.Children = make(map[string]types.YChild)
    mplsinsegmentldplsptable.EntityData.Children["mplsInSegmentLdpLspEntry"] = types.YChild{"Mplsinsegmentldplspentry", nil}
    for i := range mplsinsegmentldplsptable.Mplsinsegmentldplspentry {
        mplsinsegmentldplsptable.EntityData.Children[types.GetSegmentPath(&mplsinsegmentldplsptable.Mplsinsegmentldplspentry[i])] = types.YChild{"Mplsinsegmentldplspentry", &mplsinsegmentldplsptable.Mplsinsegmentldplspentry[i]}
    }
    mplsinsegmentldplsptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsinsegmentldplsptable.EntityData)
}

// MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry
// An entry in this table represents information
// on a single LDP LSP which is represented by
// a session's index triple (mplsLdpEntityLdpId,
// mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the
// index for the mplsInSegmentTable
// (mplsInSegmentLdpLspLabelIndex) from the
// MPLS-LSR-STD-MIB.
// 
// The information contained in a row is read-only.
type MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityldpid
    Mplsldpentityldpid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityindex
    Mplsldpentityindex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldppeerldpid
    Mplsldppeerldpid interface{}

    // This attribute is a key. This contains the same value as the
    // mplsInSegmentIndex in the MPLS-LSR-STD-MIB's mplsInSegmentTable. The type
    // is string with length: 1..24.
    Mplsinsegmentldplspindex interface{}

    // The Layer 2 Label Type. The type is MplsLdpLabelType.
    Mplsinsegmentldplsplabeltype interface{}

    // The type of LSP connection. The type is MplsLspType.
    Mplsinsegmentldplsptype interface{}
}

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetEntityData() *types.CommonEntityData {
    mplsinsegmentldplspentry.EntityData.YFilter = mplsinsegmentldplspentry.YFilter
    mplsinsegmentldplspentry.EntityData.YangName = "mplsInSegmentLdpLspEntry"
    mplsinsegmentldplspentry.EntityData.BundleName = "cisco_ios_xe"
    mplsinsegmentldplspentry.EntityData.ParentYangName = "mplsInSegmentLdpLspTable"
    mplsinsegmentldplspentry.EntityData.SegmentPath = "mplsInSegmentLdpLspEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsldppeerldpid) + "']" + "[mplsInSegmentLdpLspIndex='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsinsegmentldplspindex) + "']"
    mplsinsegmentldplspentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinsegmentldplspentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinsegmentldplspentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinsegmentldplspentry.EntityData.Children = make(map[string]types.YChild)
    mplsinsegmentldplspentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsinsegmentldplspentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsinsegmentldplspentry.Mplsldpentityldpid}
    mplsinsegmentldplspentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsinsegmentldplspentry.Mplsldpentityindex}
    mplsinsegmentldplspentry.EntityData.Leafs["mplsLdpPeerLdpId"] = types.YLeaf{"Mplsldppeerldpid", mplsinsegmentldplspentry.Mplsldppeerldpid}
    mplsinsegmentldplspentry.EntityData.Leafs["mplsInSegmentLdpLspIndex"] = types.YLeaf{"Mplsinsegmentldplspindex", mplsinsegmentldplspentry.Mplsinsegmentldplspindex}
    mplsinsegmentldplspentry.EntityData.Leafs["mplsInSegmentLdpLspLabelType"] = types.YLeaf{"Mplsinsegmentldplsplabeltype", mplsinsegmentldplspentry.Mplsinsegmentldplsplabeltype}
    mplsinsegmentldplspentry.EntityData.Leafs["mplsInSegmentLdpLspType"] = types.YLeaf{"Mplsinsegmentldplsptype", mplsinsegmentldplspentry.Mplsinsegmentldplsptype}
    return &(mplsinsegmentldplspentry.EntityData)
}

// MPLSLDPSTDMIB_Mplsoutsegmentldplsptable
// A table of LDP LSP's which
// map to the mplsOutSegmentTable in the
// MPLS-LSR-STD-MIB.
type MPLSLDPSTDMIB_Mplsoutsegmentldplsptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents information on a single LDP LSP which is
    // represented by a session's index triple (mplsLdpEntityLdpId,
    // mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the index
    // (mplsOutSegmentLdpLspIndex) for the mplsOutSegmentTable.  The information
    // contained in a row is read-only. The type is slice of
    // MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry.
    Mplsoutsegmentldplspentry []MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry
}

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetEntityData() *types.CommonEntityData {
    mplsoutsegmentldplsptable.EntityData.YFilter = mplsoutsegmentldplsptable.YFilter
    mplsoutsegmentldplsptable.EntityData.YangName = "mplsOutSegmentLdpLspTable"
    mplsoutsegmentldplsptable.EntityData.BundleName = "cisco_ios_xe"
    mplsoutsegmentldplsptable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsoutsegmentldplsptable.EntityData.SegmentPath = "mplsOutSegmentLdpLspTable"
    mplsoutsegmentldplsptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsoutsegmentldplsptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsoutsegmentldplsptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsoutsegmentldplsptable.EntityData.Children = make(map[string]types.YChild)
    mplsoutsegmentldplsptable.EntityData.Children["mplsOutSegmentLdpLspEntry"] = types.YChild{"Mplsoutsegmentldplspentry", nil}
    for i := range mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry {
        mplsoutsegmentldplsptable.EntityData.Children[types.GetSegmentPath(&mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry[i])] = types.YChild{"Mplsoutsegmentldplspentry", &mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry[i]}
    }
    mplsoutsegmentldplsptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsoutsegmentldplsptable.EntityData)
}

// MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry
// An entry in this table represents information
// on a single LDP LSP which is represented by
// a session's index triple (mplsLdpEntityLdpId,
// mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the
// index (mplsOutSegmentLdpLspIndex)
// for the mplsOutSegmentTable.
// 
// The information contained in a row is read-only.
type MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityldpid
    Mplsldpentityldpid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityindex
    Mplsldpentityindex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldppeerldpid
    Mplsldppeerldpid interface{}

    // This attribute is a key. This contains the same value as the
    // mplsOutSegmentIndex in the MPLS-LSR-STD-MIB's mplsOutSegmentTable. The type
    // is string with length: 1..24.
    Mplsoutsegmentldplspindex interface{}

    // The Layer 2 Label Type. The type is MplsLdpLabelType.
    Mplsoutsegmentldplsplabeltype interface{}

    // The type of LSP connection. The type is MplsLspType.
    Mplsoutsegmentldplsptype interface{}
}

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetEntityData() *types.CommonEntityData {
    mplsoutsegmentldplspentry.EntityData.YFilter = mplsoutsegmentldplspentry.YFilter
    mplsoutsegmentldplspentry.EntityData.YangName = "mplsOutSegmentLdpLspEntry"
    mplsoutsegmentldplspentry.EntityData.BundleName = "cisco_ios_xe"
    mplsoutsegmentldplspentry.EntityData.ParentYangName = "mplsOutSegmentLdpLspTable"
    mplsoutsegmentldplspentry.EntityData.SegmentPath = "mplsOutSegmentLdpLspEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsldppeerldpid) + "']" + "[mplsOutSegmentLdpLspIndex='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsoutsegmentldplspindex) + "']"
    mplsoutsegmentldplspentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsoutsegmentldplspentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsoutsegmentldplspentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsoutsegmentldplspentry.EntityData.Children = make(map[string]types.YChild)
    mplsoutsegmentldplspentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsoutsegmentldplspentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsoutsegmentldplspentry.Mplsldpentityldpid}
    mplsoutsegmentldplspentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsoutsegmentldplspentry.Mplsldpentityindex}
    mplsoutsegmentldplspentry.EntityData.Leafs["mplsLdpPeerLdpId"] = types.YLeaf{"Mplsldppeerldpid", mplsoutsegmentldplspentry.Mplsldppeerldpid}
    mplsoutsegmentldplspentry.EntityData.Leafs["mplsOutSegmentLdpLspIndex"] = types.YLeaf{"Mplsoutsegmentldplspindex", mplsoutsegmentldplspentry.Mplsoutsegmentldplspindex}
    mplsoutsegmentldplspentry.EntityData.Leafs["mplsOutSegmentLdpLspLabelType"] = types.YLeaf{"Mplsoutsegmentldplsplabeltype", mplsoutsegmentldplspentry.Mplsoutsegmentldplsplabeltype}
    mplsoutsegmentldplspentry.EntityData.Leafs["mplsOutSegmentLdpLspType"] = types.YLeaf{"Mplsoutsegmentldplsptype", mplsoutsegmentldplspentry.Mplsoutsegmentldplsptype}
    return &(mplsoutsegmentldplspentry.EntityData)
}

// MPLSLDPSTDMIB_Mplsfectable
// This table represents the FEC
// (Forwarding Equivalence Class)
// Information associated with an LSP.
type MPLSLDPSTDMIB_Mplsfectable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row represents a single FEC Element. The type is slice of
    // MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry.
    Mplsfecentry []MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry
}

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetEntityData() *types.CommonEntityData {
    mplsfectable.EntityData.YFilter = mplsfectable.YFilter
    mplsfectable.EntityData.YangName = "mplsFecTable"
    mplsfectable.EntityData.BundleName = "cisco_ios_xe"
    mplsfectable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsfectable.EntityData.SegmentPath = "mplsFecTable"
    mplsfectable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsfectable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsfectable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsfectable.EntityData.Children = make(map[string]types.YChild)
    mplsfectable.EntityData.Children["mplsFecEntry"] = types.YChild{"Mplsfecentry", nil}
    for i := range mplsfectable.Mplsfecentry {
        mplsfectable.EntityData.Children[types.GetSegmentPath(&mplsfectable.Mplsfecentry[i])] = types.YChild{"Mplsfecentry", &mplsfectable.Mplsfecentry[i]}
    }
    mplsfectable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsfectable.EntityData)
}

// MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry
// Each row represents a single FEC Element.
type MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index which uniquely identifies this entry.
    // The type is interface{} with range: 1..4294967295.
    Mplsfecindex interface{}

    // The type of the FEC.  If the value of this object is 'prefix(1)' then the
    // FEC type described by this row is an address prefix.  If the value of this
    // object is 'hostAddress(2)' then the FEC type described by this row is a
    // host address. The type is Mplsfectype.
    Mplsfectype interface{}

    // If the value of the 'mplsFecType' is 'hostAddress(2)' then this object is
    // undefined.  If the value of 'mplsFecType' is 'prefix(1)' then the value of
    // this object is the length in bits of the address prefix represented by
    // 'mplsFecAddr', or zero.  If the value of this object is zero, this
    // indicates that the prefix matches all addresses.  In this case the address
    // prefix MUST also be zero (i.e., 'mplsFecAddr' should have the value of
    // zero.). The type is interface{} with range: 0..2040.
    Mplsfecaddrprefixlength interface{}

    // The value of this object is the type of the Internet address.  The value of
    // this object, decides how the value of the mplsFecAddr object is
    // interpreted. The type is InetAddressType.
    Mplsfecaddrtype interface{}

    // The value of this object is interpreted based on the value of the
    // 'mplsFecAddrType' object.  This address is then further interpretted as an
    // being used with the address prefix, or as the host address.  This further
    // interpretation is indicated by the 'mplsFecType' object. In other words,
    // the FEC element is populated according to the Prefix FEC Element value
    // encoding, or the Host Address FEC Element encoding. The type is string with
    // length: 0..255.
    Mplsfecaddr interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    Mplsfecstoragetype interface{}

    // The status of this conceptual row.  If the value of this object is
    // 'active(1)', then none of the writable objects of this entry can be
    // modified, except to set this object to 'destroy(6)'.  NOTE: if this row is
    // being referenced by any entry in the mplsLdpLspFecTable, then a request to
    // destroy this row, will result in an inconsistentValue error. The type is
    // RowStatus.
    Mplsfecrowstatus interface{}
}

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetEntityData() *types.CommonEntityData {
    mplsfecentry.EntityData.YFilter = mplsfecentry.YFilter
    mplsfecentry.EntityData.YangName = "mplsFecEntry"
    mplsfecentry.EntityData.BundleName = "cisco_ios_xe"
    mplsfecentry.EntityData.ParentYangName = "mplsFecTable"
    mplsfecentry.EntityData.SegmentPath = "mplsFecEntry" + "[mplsFecIndex='" + fmt.Sprintf("%v", mplsfecentry.Mplsfecindex) + "']"
    mplsfecentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsfecentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsfecentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsfecentry.EntityData.Children = make(map[string]types.YChild)
    mplsfecentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsfecentry.EntityData.Leafs["mplsFecIndex"] = types.YLeaf{"Mplsfecindex", mplsfecentry.Mplsfecindex}
    mplsfecentry.EntityData.Leafs["mplsFecType"] = types.YLeaf{"Mplsfectype", mplsfecentry.Mplsfectype}
    mplsfecentry.EntityData.Leafs["mplsFecAddrPrefixLength"] = types.YLeaf{"Mplsfecaddrprefixlength", mplsfecentry.Mplsfecaddrprefixlength}
    mplsfecentry.EntityData.Leafs["mplsFecAddrType"] = types.YLeaf{"Mplsfecaddrtype", mplsfecentry.Mplsfecaddrtype}
    mplsfecentry.EntityData.Leafs["mplsFecAddr"] = types.YLeaf{"Mplsfecaddr", mplsfecentry.Mplsfecaddr}
    mplsfecentry.EntityData.Leafs["mplsFecStorageType"] = types.YLeaf{"Mplsfecstoragetype", mplsfecentry.Mplsfecstoragetype}
    mplsfecentry.EntityData.Leafs["mplsFecRowStatus"] = types.YLeaf{"Mplsfecrowstatus", mplsfecentry.Mplsfecrowstatus}
    return &(mplsfecentry.EntityData)
}

// MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry_Mplsfectype represents the FEC type described by this row is a host address.
type MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry_Mplsfectype string

const (
    MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry_Mplsfectype_prefix MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry_Mplsfectype = "prefix"

    MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry_Mplsfectype_hostAddress MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry_Mplsfectype = "hostAddress"
)

// MPLSLDPSTDMIB_Mplsldplspfectable
// A table which shows the relationship between
// LDP LSPs and FECs.  Each row represents
// a single LDP LSP to FEC association.
type MPLSLDPSTDMIB_Mplsldplspfectable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry represents a LDP LSP to FEC association. The type is slice of
    // MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry.
    Mplsldplspfecentry []MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry
}

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetEntityData() *types.CommonEntityData {
    mplsldplspfectable.EntityData.YFilter = mplsldplspfectable.YFilter
    mplsldplspfectable.EntityData.YangName = "mplsLdpLspFecTable"
    mplsldplspfectable.EntityData.BundleName = "cisco_ios_xe"
    mplsldplspfectable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldplspfectable.EntityData.SegmentPath = "mplsLdpLspFecTable"
    mplsldplspfectable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldplspfectable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldplspfectable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldplspfectable.EntityData.Children = make(map[string]types.YChild)
    mplsldplspfectable.EntityData.Children["mplsLdpLspFecEntry"] = types.YChild{"Mplsldplspfecentry", nil}
    for i := range mplsldplspfectable.Mplsldplspfecentry {
        mplsldplspfectable.EntityData.Children[types.GetSegmentPath(&mplsldplspfectable.Mplsldplspfecentry[i])] = types.YChild{"Mplsldplspfecentry", &mplsldplspfectable.Mplsldplspfecentry[i]}
    }
    mplsldplspfectable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsldplspfectable.EntityData)
}

// MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry
// An entry represents a LDP LSP
// to FEC association.
type MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityldpid
    Mplsldpentityldpid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityindex
    Mplsldpentityindex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldppeerldpid
    Mplsldppeerldpid interface{}

    // This attribute is a key. If the value is inSegment(1), then this indicates
    // that the following index, mplsLdpLspFecSegmentIndex, contains the same
    // value as the mplsInSegmentLdpLspIndex.  Otherwise, if the value of this
    // object is   outSegment(2),  then this indicates that following index,
    // mplsLdpLspFecSegmentIndex, contains the same value as the
    // mplsOutSegmentLdpLspIndex. The type is Mplsldplspfecsegment.
    Mplsldplspfecsegment interface{}

    // This attribute is a key. This index is interpretted by using the value of
    // the mplsLdpLspFecSegment.  If the mplsLdpLspFecSegment is inSegment(1),
    // then this index has the same value as mplsInSegmentLdpLspIndex.  If the
    // mplsLdpLspFecSegment is outSegment(2), then this index has the same value
    // as mplsOutSegmentLdpLspIndex. The type is string with length: 1..24.
    Mplsldplspfecsegmentindex interface{}

    // This attribute is a key. This index identifies the FEC entry in the
    // mplsFecTable associated with this session. In other words, the value of
    // this index is the same as the value of the mplsFecIndex that denotes the
    // FEC associated with this Session. The type is interface{} with range:
    // 1..4294967295.
    Mplsldplspfecindex interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    Mplsldplspfecstoragetype interface{}

    // The status of this conceptual row.  If the value of this object is
    // 'active(1)', then none of the writable objects of this entry can be
    // modified.  The Agent should delete this row when the session ceases to
    // exist.  If an operator wants to associate the session with a different FEC,
    // the recommended procedure is (as described in detail in the section
    // entitled, 'Changing Values After Session Establishment', and again
    // described in the DESCRIPTION clause of the mplsLdpEntityAdminStatus object)
    // is to set the mplsLdpEntityAdminStatus to down, thereby explicitly causing
    // a session to be torn down. This will also cause this entry to be deleted. 
    // Then, set the mplsLdpEntityAdminStatus to enable which enables a new
    // session to be initiated. Once the session is initiated, an entry may be
    // added to this table to associate the new session with a FEC. The type is
    // RowStatus.
    Mplsldplspfecrowstatus interface{}
}

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetEntityData() *types.CommonEntityData {
    mplsldplspfecentry.EntityData.YFilter = mplsldplspfecentry.YFilter
    mplsldplspfecentry.EntityData.YangName = "mplsLdpLspFecEntry"
    mplsldplspfecentry.EntityData.BundleName = "cisco_ios_xe"
    mplsldplspfecentry.EntityData.ParentYangName = "mplsLdpLspFecTable"
    mplsldplspfecentry.EntityData.SegmentPath = "mplsLdpLspFecEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldppeerldpid) + "']" + "[mplsLdpLspFecSegment='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldplspfecsegment) + "']" + "[mplsLdpLspFecSegmentIndex='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldplspfecsegmentindex) + "']" + "[mplsLdpLspFecIndex='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldplspfecindex) + "']"
    mplsldplspfecentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldplspfecentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldplspfecentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldplspfecentry.EntityData.Children = make(map[string]types.YChild)
    mplsldplspfecentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldplspfecentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsldplspfecentry.Mplsldpentityldpid}
    mplsldplspfecentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsldplspfecentry.Mplsldpentityindex}
    mplsldplspfecentry.EntityData.Leafs["mplsLdpPeerLdpId"] = types.YLeaf{"Mplsldppeerldpid", mplsldplspfecentry.Mplsldppeerldpid}
    mplsldplspfecentry.EntityData.Leafs["mplsLdpLspFecSegment"] = types.YLeaf{"Mplsldplspfecsegment", mplsldplspfecentry.Mplsldplspfecsegment}
    mplsldplspfecentry.EntityData.Leafs["mplsLdpLspFecSegmentIndex"] = types.YLeaf{"Mplsldplspfecsegmentindex", mplsldplspfecentry.Mplsldplspfecsegmentindex}
    mplsldplspfecentry.EntityData.Leafs["mplsLdpLspFecIndex"] = types.YLeaf{"Mplsldplspfecindex", mplsldplspfecentry.Mplsldplspfecindex}
    mplsldplspfecentry.EntityData.Leafs["mplsLdpLspFecStorageType"] = types.YLeaf{"Mplsldplspfecstoragetype", mplsldplspfecentry.Mplsldplspfecstoragetype}
    mplsldplspfecentry.EntityData.Leafs["mplsLdpLspFecRowStatus"] = types.YLeaf{"Mplsldplspfecrowstatus", mplsldplspfecentry.Mplsldplspfecrowstatus}
    return &(mplsldplspfecentry.EntityData)
}

// MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry_Mplsldplspfecsegment represents value as the mplsOutSegmentLdpLspIndex.
type MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry_Mplsldplspfecsegment string

const (
    MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry_Mplsldplspfecsegment_inSegment MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry_Mplsldplspfecsegment = "inSegment"

    MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry_Mplsldplspfecsegment_outSegment MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry_Mplsldplspfecsegment = "outSegment"
)

// MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable
// This table 'extends' the mplsLdpSessionTable.
// This table is used to store Label Address Information
// from Label Address Messages received by this LSR from
// Peers.  This table is read-only and should be updated
// 
// 
// when Label Withdraw Address Messages are received, i.e.,
// Rows should be deleted as appropriate.
// 
// NOTE:  since more than one address may be contained
// in a Label Address Message, this table 'sparse augments',
// the mplsLdpSessionTable's information.
type MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents information on a session's single next
    // hop address which was advertised in an Address Message from the LDP peer.
    // The information contained in a row is read-only. The type is slice of
    // MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry.
    Mplsldpsessionpeeraddrentry []MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry
}

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetEntityData() *types.CommonEntityData {
    mplsldpsessionpeeraddrtable.EntityData.YFilter = mplsldpsessionpeeraddrtable.YFilter
    mplsldpsessionpeeraddrtable.EntityData.YangName = "mplsLdpSessionPeerAddrTable"
    mplsldpsessionpeeraddrtable.EntityData.BundleName = "cisco_ios_xe"
    mplsldpsessionpeeraddrtable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsldpsessionpeeraddrtable.EntityData.SegmentPath = "mplsLdpSessionPeerAddrTable"
    mplsldpsessionpeeraddrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpsessionpeeraddrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpsessionpeeraddrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpsessionpeeraddrtable.EntityData.Children = make(map[string]types.YChild)
    mplsldpsessionpeeraddrtable.EntityData.Children["mplsLdpSessionPeerAddrEntry"] = types.YChild{"Mplsldpsessionpeeraddrentry", nil}
    for i := range mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry {
        mplsldpsessionpeeraddrtable.EntityData.Children[types.GetSegmentPath(&mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry[i])] = types.YChild{"Mplsldpsessionpeeraddrentry", &mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry[i]}
    }
    mplsldpsessionpeeraddrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsldpsessionpeeraddrtable.EntityData)
}

// MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry
// An entry in this table represents information on
// a session's single next hop address which was
// advertised in an Address Message from the LDP peer.
// The information contained in a row is read-only.
type MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityldpid
    Mplsldpentityldpid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityindex
    Mplsldpentityindex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry_Mplsldppeerldpid
    Mplsldppeerldpid interface{}

    // This attribute is a key. An index which uniquely identifies this entry
    // within a given session. The type is interface{} with range: 1..4294967295.
    Mplsldpsessionpeeraddrindex interface{}

    // The internetwork layer address type of this Next Hop Address as specified
    // in the Label Address Message associated with this Session. The value of
    // this object indicates how to interpret the value of  
    // mplsLdpSessionPeerNextHopAddr. The type is InetAddressType.
    Mplsldpsessionpeernexthopaddrtype interface{}

    // The next hop address.  The type of this address is specified by the value
    // of the mplsLdpSessionPeerNextHopAddrType. The type is string with length:
    // 0..255.
    Mplsldpsessionpeernexthopaddr interface{}
}

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetEntityData() *types.CommonEntityData {
    mplsldpsessionpeeraddrentry.EntityData.YFilter = mplsldpsessionpeeraddrentry.YFilter
    mplsldpsessionpeeraddrentry.EntityData.YangName = "mplsLdpSessionPeerAddrEntry"
    mplsldpsessionpeeraddrentry.EntityData.BundleName = "cisco_ios_xe"
    mplsldpsessionpeeraddrentry.EntityData.ParentYangName = "mplsLdpSessionPeerAddrTable"
    mplsldpsessionpeeraddrentry.EntityData.SegmentPath = "mplsLdpSessionPeerAddrEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldppeerldpid) + "']" + "[mplsLdpSessionPeerAddrIndex='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldpsessionpeeraddrindex) + "']"
    mplsldpsessionpeeraddrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpsessionpeeraddrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpsessionpeeraddrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpsessionpeeraddrentry.EntityData.Children = make(map[string]types.YChild)
    mplsldpsessionpeeraddrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldpsessionpeeraddrentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsldpsessionpeeraddrentry.Mplsldpentityldpid}
    mplsldpsessionpeeraddrentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsldpsessionpeeraddrentry.Mplsldpentityindex}
    mplsldpsessionpeeraddrentry.EntityData.Leafs["mplsLdpPeerLdpId"] = types.YLeaf{"Mplsldppeerldpid", mplsldpsessionpeeraddrentry.Mplsldppeerldpid}
    mplsldpsessionpeeraddrentry.EntityData.Leafs["mplsLdpSessionPeerAddrIndex"] = types.YLeaf{"Mplsldpsessionpeeraddrindex", mplsldpsessionpeeraddrentry.Mplsldpsessionpeeraddrindex}
    mplsldpsessionpeeraddrentry.EntityData.Leafs["mplsLdpSessionPeerNextHopAddrType"] = types.YLeaf{"Mplsldpsessionpeernexthopaddrtype", mplsldpsessionpeeraddrentry.Mplsldpsessionpeernexthopaddrtype}
    mplsldpsessionpeeraddrentry.EntityData.Leafs["mplsLdpSessionPeerNextHopAddr"] = types.YLeaf{"Mplsldpsessionpeernexthopaddr", mplsldpsessionpeeraddrentry.Mplsldpsessionpeernexthopaddr}
    return &(mplsldpsessionpeeraddrentry.EntityData)
}

