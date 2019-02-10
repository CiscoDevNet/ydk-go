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

    
    MplsLdpLsrObjects MPLSLDPSTDMIB_MplsLdpLsrObjects

    
    MplsLdpEntityObjects MPLSLDPSTDMIB_MplsLdpEntityObjects

    
    MplsLdpSessionObjects MPLSLDPSTDMIB_MplsLdpSessionObjects

    
    MplsFecObjects MPLSLDPSTDMIB_MplsFecObjects

    // This table contains information about the MPLS Label Distribution Protocol
    // Entities which exist on this Label Switching Router (LSR) or Label Edge
    // Router (LER).
    MplsLdpEntityTable MPLSLDPSTDMIB_MplsLdpEntityTable

    // Information about LDP peers known by Entities in the mplsLdpEntityTable. 
    // The information in this table is based on information from the Entity-Peer
    // interaction during session initialization but is not appropriate for the
    // mplsLdpSessionTable, because objects in this table may or may not be used
    // in session establishment.
    MplsLdpPeerTable MPLSLDPSTDMIB_MplsLdpPeerTable

    // A table of Hello Adjacencies for Sessions.
    MplsLdpHelloAdjacencyTable MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable

    // A table of LDP LSP's which map to the mplsInSegmentTable in the
    // MPLS-LSR-STD-MIB module.
    MplsInSegmentLdpLspTable MPLSLDPSTDMIB_MplsInSegmentLdpLspTable

    // A table of LDP LSP's which map to the mplsOutSegmentTable in the
    // MPLS-LSR-STD-MIB.
    MplsOutSegmentLdpLspTable MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable

    // This table represents the FEC (Forwarding Equivalence Class) Information
    // associated with an LSP.
    MplsFecTable MPLSLDPSTDMIB_MplsFecTable

    // A table which shows the relationship between LDP LSPs and FECs.  Each row
    // represents a single LDP LSP to FEC association.
    MplsLdpLspFecTable MPLSLDPSTDMIB_MplsLdpLspFecTable

    // This table 'extends' the mplsLdpSessionTable. This table is used to store
    // Label Address Information from Label Address Messages received by this LSR
    // from Peers.  This table is read-only and should be updated   when Label
    // Withdraw Address Messages are received, i.e., Rows should be deleted as
    // appropriate.  NOTE:  since more than one address may be contained in a
    // Label Address Message, this table 'sparse augments', the
    // mplsLdpSessionTable's information.
    MplsLdpSessionPeerAddrTable MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable
}

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSLDPSTDMIB.EntityData.YFilter = mPLSLDPSTDMIB.YFilter
    mPLSLDPSTDMIB.EntityData.YangName = "MPLS-LDP-STD-MIB"
    mPLSLDPSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSLDPSTDMIB.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mPLSLDPSTDMIB.EntityData.SegmentPath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB"
    mPLSLDPSTDMIB.EntityData.AbsolutePath = mPLSLDPSTDMIB.EntityData.SegmentPath
    mPLSLDPSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSLDPSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSLDPSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSLDPSTDMIB.EntityData.Children = types.NewOrderedMap()
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpLsrObjects", types.YChild{"MplsLdpLsrObjects", &mPLSLDPSTDMIB.MplsLdpLsrObjects})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpEntityObjects", types.YChild{"MplsLdpEntityObjects", &mPLSLDPSTDMIB.MplsLdpEntityObjects})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpSessionObjects", types.YChild{"MplsLdpSessionObjects", &mPLSLDPSTDMIB.MplsLdpSessionObjects})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsFecObjects", types.YChild{"MplsFecObjects", &mPLSLDPSTDMIB.MplsFecObjects})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpEntityTable", types.YChild{"MplsLdpEntityTable", &mPLSLDPSTDMIB.MplsLdpEntityTable})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpPeerTable", types.YChild{"MplsLdpPeerTable", &mPLSLDPSTDMIB.MplsLdpPeerTable})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpHelloAdjacencyTable", types.YChild{"MplsLdpHelloAdjacencyTable", &mPLSLDPSTDMIB.MplsLdpHelloAdjacencyTable})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsInSegmentLdpLspTable", types.YChild{"MplsInSegmentLdpLspTable", &mPLSLDPSTDMIB.MplsInSegmentLdpLspTable})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsOutSegmentLdpLspTable", types.YChild{"MplsOutSegmentLdpLspTable", &mPLSLDPSTDMIB.MplsOutSegmentLdpLspTable})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsFecTable", types.YChild{"MplsFecTable", &mPLSLDPSTDMIB.MplsFecTable})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpLspFecTable", types.YChild{"MplsLdpLspFecTable", &mPLSLDPSTDMIB.MplsLdpLspFecTable})
    mPLSLDPSTDMIB.EntityData.Children.Append("mplsLdpSessionPeerAddrTable", types.YChild{"MplsLdpSessionPeerAddrTable", &mPLSLDPSTDMIB.MplsLdpSessionPeerAddrTable})
    mPLSLDPSTDMIB.EntityData.Leafs = types.NewOrderedMap()

    mPLSLDPSTDMIB.EntityData.YListKeys = []string {}

    return &(mPLSLDPSTDMIB.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpLsrObjects
type MPLSLDPSTDMIB_MplsLdpLsrObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Label Switching Router's Identifier. The type is string with length: 4.
    MplsLdpLsrId interface{}

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
    // Detection and which types. The type is MplsLdpLsrLoopDetectionCapable.
    MplsLdpLsrLoopDetectionCapable interface{}
}

func (mplsLdpLsrObjects *MPLSLDPSTDMIB_MplsLdpLsrObjects) GetEntityData() *types.CommonEntityData {
    mplsLdpLsrObjects.EntityData.YFilter = mplsLdpLsrObjects.YFilter
    mplsLdpLsrObjects.EntityData.YangName = "mplsLdpLsrObjects"
    mplsLdpLsrObjects.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpLsrObjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpLsrObjects.EntityData.SegmentPath = "mplsLdpLsrObjects"
    mplsLdpLsrObjects.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpLsrObjects.EntityData.SegmentPath
    mplsLdpLsrObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpLsrObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpLsrObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpLsrObjects.EntityData.Children = types.NewOrderedMap()
    mplsLdpLsrObjects.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpLsrObjects.EntityData.Leafs.Append("mplsLdpLsrId", types.YLeaf{"MplsLdpLsrId", mplsLdpLsrObjects.MplsLdpLsrId})
    mplsLdpLsrObjects.EntityData.Leafs.Append("mplsLdpLsrLoopDetectionCapable", types.YLeaf{"MplsLdpLsrLoopDetectionCapable", mplsLdpLsrObjects.MplsLdpLsrLoopDetectionCapable})

    mplsLdpLsrObjects.EntityData.YListKeys = []string {}

    return &(mplsLdpLsrObjects.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable represents which types.
type MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable string

const (
    MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable_none MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable = "none"

    MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable_other MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable = "other"

    MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable_hopCount MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable = "hopCount"

    MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable_pathVector MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable = "pathVector"

    MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable_hopCountAndPathVector MPLSLDPSTDMIB_MplsLdpLsrObjects_MplsLdpLsrLoopDetectionCapable = "hopCountAndPathVector"
)

// MPLSLDPSTDMIB_MplsLdpEntityObjects
type MPLSLDPSTDMIB_MplsLdpEntityObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the most recent addition or deletion
    // of an entry to/from the mplsLdpEntityTable/mplsLdpEntityStatsTable, or the
    // most recent change in value of any objects in the mplsLdpEntityTable.   If
    // no such changes have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    MplsLdpEntityLastChange interface{}

    // This object contains an appropriate value to be used for mplsLdpEntityIndex
    // when creating entries in the mplsLdpEntityTable. The value 0 indicates that
    // no unassigned entries are available. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityIndexNext interface{}
}

func (mplsLdpEntityObjects *MPLSLDPSTDMIB_MplsLdpEntityObjects) GetEntityData() *types.CommonEntityData {
    mplsLdpEntityObjects.EntityData.YFilter = mplsLdpEntityObjects.YFilter
    mplsLdpEntityObjects.EntityData.YangName = "mplsLdpEntityObjects"
    mplsLdpEntityObjects.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpEntityObjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpEntityObjects.EntityData.SegmentPath = "mplsLdpEntityObjects"
    mplsLdpEntityObjects.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpEntityObjects.EntityData.SegmentPath
    mplsLdpEntityObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpEntityObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpEntityObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpEntityObjects.EntityData.Children = types.NewOrderedMap()
    mplsLdpEntityObjects.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpEntityObjects.EntityData.Leafs.Append("mplsLdpEntityLastChange", types.YLeaf{"MplsLdpEntityLastChange", mplsLdpEntityObjects.MplsLdpEntityLastChange})
    mplsLdpEntityObjects.EntityData.Leafs.Append("mplsLdpEntityIndexNext", types.YLeaf{"MplsLdpEntityIndexNext", mplsLdpEntityObjects.MplsLdpEntityIndexNext})

    mplsLdpEntityObjects.EntityData.YListKeys = []string {}

    return &(mplsLdpEntityObjects.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpSessionObjects
type MPLSLDPSTDMIB_MplsLdpSessionObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the most recent addition or deletion
    // to/from the mplsLdpPeerTable/mplsLdpSessionTable. The type is interface{}
    // with range: 0..4294967295.
    MplsLdpPeerLastChange interface{}

    // The value of sysUpTime at the time of the most recent addition/deletion of
    // an entry to/from the mplsLdpLspFecTable or the most recent change in values
    // to any objects in the mplsLdpLspFecTable.  If no such changes have occurred
    // since the last re-initialization of the local management subsystem, then
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpLspFecLastChange interface{}
}

func (mplsLdpSessionObjects *MPLSLDPSTDMIB_MplsLdpSessionObjects) GetEntityData() *types.CommonEntityData {
    mplsLdpSessionObjects.EntityData.YFilter = mplsLdpSessionObjects.YFilter
    mplsLdpSessionObjects.EntityData.YangName = "mplsLdpSessionObjects"
    mplsLdpSessionObjects.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpSessionObjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpSessionObjects.EntityData.SegmentPath = "mplsLdpSessionObjects"
    mplsLdpSessionObjects.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpSessionObjects.EntityData.SegmentPath
    mplsLdpSessionObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpSessionObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpSessionObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpSessionObjects.EntityData.Children = types.NewOrderedMap()
    mplsLdpSessionObjects.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpSessionObjects.EntityData.Leafs.Append("mplsLdpPeerLastChange", types.YLeaf{"MplsLdpPeerLastChange", mplsLdpSessionObjects.MplsLdpPeerLastChange})
    mplsLdpSessionObjects.EntityData.Leafs.Append("mplsLdpLspFecLastChange", types.YLeaf{"MplsLdpLspFecLastChange", mplsLdpSessionObjects.MplsLdpLspFecLastChange})

    mplsLdpSessionObjects.EntityData.YListKeys = []string {}

    return &(mplsLdpSessionObjects.EntityData)
}

// MPLSLDPSTDMIB_MplsFecObjects
type MPLSLDPSTDMIB_MplsFecObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the most recent addition/deletion of
    // an entry to/from the mplsLdpFectTable or the most recent change in values
    // to any objects in the mplsLdpFecTable.  If no such changes have occurred
    // since the last re-initialization of the local management subsystem, then
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    MplsFecLastChange interface{}

    // This object contains an appropriate value to be used for mplsFecIndex when
    // creating entries in the mplsFecTable. The value 0 indicates that no
    // unassigned entries are available. The type is interface{} with range:
    // 0..4294967295.
    MplsFecIndexNext interface{}
}

func (mplsFecObjects *MPLSLDPSTDMIB_MplsFecObjects) GetEntityData() *types.CommonEntityData {
    mplsFecObjects.EntityData.YFilter = mplsFecObjects.YFilter
    mplsFecObjects.EntityData.YangName = "mplsFecObjects"
    mplsFecObjects.EntityData.BundleName = "cisco_ios_xe"
    mplsFecObjects.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsFecObjects.EntityData.SegmentPath = "mplsFecObjects"
    mplsFecObjects.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsFecObjects.EntityData.SegmentPath
    mplsFecObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsFecObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsFecObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsFecObjects.EntityData.Children = types.NewOrderedMap()
    mplsFecObjects.EntityData.Leafs = types.NewOrderedMap()
    mplsFecObjects.EntityData.Leafs.Append("mplsFecLastChange", types.YLeaf{"MplsFecLastChange", mplsFecObjects.MplsFecLastChange})
    mplsFecObjects.EntityData.Leafs.Append("mplsFecIndexNext", types.YLeaf{"MplsFecIndexNext", mplsFecObjects.MplsFecIndexNext})

    mplsFecObjects.EntityData.YListKeys = []string {}

    return &(mplsFecObjects.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpEntityTable
// This table contains information about the
// MPLS Label Distribution Protocol Entities which
// exist on this Label Switching Router (LSR)
// or Label Edge Router (LER).
type MPLSLDPSTDMIB_MplsLdpEntityTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents an LDP entity. An entry can be created by
    // a network administrator or by an SNMP agent as instructed by LDP. The type
    // is slice of MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry.
    MplsLdpEntityEntry []*MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry
}

func (mplsLdpEntityTable *MPLSLDPSTDMIB_MplsLdpEntityTable) GetEntityData() *types.CommonEntityData {
    mplsLdpEntityTable.EntityData.YFilter = mplsLdpEntityTable.YFilter
    mplsLdpEntityTable.EntityData.YangName = "mplsLdpEntityTable"
    mplsLdpEntityTable.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpEntityTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpEntityTable.EntityData.SegmentPath = "mplsLdpEntityTable"
    mplsLdpEntityTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpEntityTable.EntityData.SegmentPath
    mplsLdpEntityTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpEntityTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpEntityTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpEntityTable.EntityData.Children = types.NewOrderedMap()
    mplsLdpEntityTable.EntityData.Children.Append("mplsLdpEntityEntry", types.YChild{"MplsLdpEntityEntry", nil})
    for i := range mplsLdpEntityTable.MplsLdpEntityEntry {
        mplsLdpEntityTable.EntityData.Children.Append(types.GetSegmentPath(mplsLdpEntityTable.MplsLdpEntityEntry[i]), types.YChild{"MplsLdpEntityEntry", mplsLdpEntityTable.MplsLdpEntityEntry[i]})
    }
    mplsLdpEntityTable.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpEntityTable.EntityData.YListKeys = []string {}

    return &(mplsLdpEntityTable.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry
// An entry in this table represents an LDP entity.
// An entry can be created by a network administrator
// or by an SNMP agent as instructed by LDP.
type MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The LDP identifier. The type is string.
    MplsLdpEntityLdpId interface{}

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
    MplsLdpEntityIndex interface{}

    // The version number of the LDP protocol which will be used in the session
    // initialization message.  Section 3.5.3 in the LDP Specification specifies
    // that the version of the LDP protocol is negotiated during session
    // establishment. The value of this object represents the value that is sent
    // in the initialization message. The type is interface{} with range:
    // 1..65535.
    MplsLdpEntityProtocolVersion interface{}

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
    // the Peer. The type is MplsLdpEntityAdminStatus.
    MplsLdpEntityAdminStatus interface{}

    // The operational status of this LDP Entity.  The value of unknown(1)
    // indicates that the operational status cannot be determined at this time. 
    // The value of unknown should be a transient condition before changing to
    // enabled(2) or disabled(3). The type is MplsLdpEntityOperStatus.
    MplsLdpEntityOperStatus interface{}

    // The TCP Port for LDP.  The default value is the well-known value of this
    // port. The type is interface{} with range: 0..65535.
    MplsLdpEntityTcpPort interface{}

    // The UDP Discovery Port for LDP.  The default value is the well-known value
    // for this port. The type is interface{} with range: 0..65535.
    MplsLdpEntityUdpDscPort interface{}

    // The maximum PDU Length that is sent in the Common Session Parameters of an
    // Initialization Message. According to the LDP Specification [RFC3036] a
    // value of 255 or less specifies the default maximum length of 4096 octets,
    // this is why the value of this object starts at 256.  The operator should
    // explicitly choose the default value (i.e., 4096), or some other value.  The
    // receiving LSR MUST calculate the maximum PDU length for the session by
    // using the smaller of its and its peer's proposals for Max PDU Length. The
    // type is interface{} with range: 256..65535. Units are octets.
    MplsLdpEntityMaxPduLength interface{}

    // The 16-bit integer value which is the proposed keep alive hold timer for
    // this LDP Entity. The type is interface{} with range: 1..65535. Units are
    // seconds.
    MplsLdpEntityKeepAliveHoldTimer interface{}

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
    MplsLdpEntityHelloHoldTimer interface{}

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
    MplsLdpEntityInitSessionThreshold interface{}

    // For any given LDP session, the method of label distribution must be
    // specified. The type is MplsLabelDistributionMethod.
    MplsLdpEntityLabelDistMethod interface{}

    // The LDP Entity can be configured to use either conservative or liberal
    // label retention mode.  If the value of this object is conservative(1) then
    // advertized label mappings are retained only if they will be used to forward
    // packets, i.e., if label came from a valid next hop.  If the value of this
    // object is liberal(2) then all advertized label mappings are retained
    // whether they are from a valid next hop or not. The type is
    // MplsRetentionMode.
    MplsLdpEntityLabelRetentionMode interface{}

    // If the value of this object is 0 (zero) then Loop Detection for Path
    // Vectors is disabled.  Otherwise, if this object has a value greater than
    // zero, then Loop Dection for Path Vectors is enabled, and the Path Vector
    // Limit is this value. Also, the value of the object,
    // 'mplsLdpLsrLoopDetectionCapable', must be set to either 'pathVector(4)' or
    // 'hopCountAndPathVector(5)', if this object has a value greater than 0
    // (zero), otherwise it is ignored. The type is interface{} with range:
    // 0..255.
    MplsLdpEntityPathVectorLimit interface{}

    // If the value of this object is 0 (zero), then Loop Detection using Hop
    // Counters is disabled.  If the value of this object is greater than 0 (zero)
    // then Loop Detection using Hop Counters is enabled, and this object
    // specifies this Entity's maximum allowable value for the Hop Count. Also,
    // the value of the object mplsLdpLsrLoopDetectionCapable must be set to
    // either 'hopCount(3)' or 'hopCountAndPathVector(5)' if this object has a
    // value greater than 0 (zero), otherwise it is ignored. The type is
    // interface{} with range: 0..255.
    MplsLdpEntityHopCountLimit interface{}

    // This specifies whether the loopback or interface address is to be used as
    // the transport address in the transport address TLV of the hello message. 
    // If the value is interface(1), then the IP address of the interface from
    // which hello messages are sent is used as the transport address in the hello
    // message.  Otherwise, if the value is loopback(2), then the IP address of
    // the loopback interface is used as the transport address in the hello
    // message. The type is MplsLdpEntityTransportAddrKind.
    MplsLdpEntityTransportAddrKind interface{}

    // If this LDP entity uses targeted peer then set this to true. The type is
    // bool.
    MplsLdpEntityTargetPeer interface{}

    // The type of the internetwork layer address used for the Extended Discovery.
    // This object indicates how the value of mplsLdpEntityTargetPeerAddr is to be
    // interpreted. The type is InetAddressType.
    MplsLdpEntityTargetPeerAddrType interface{}

    // The value of the internetwork layer address used for the Extended
    // Discovery.  The value of mplsLdpEntityTargetPeerAddrType specifies how this
    // address is to be interpreted. The type is string with length: 0..255.
    MplsLdpEntityTargetPeerAddr interface{}

    // Specifies the optional parameters for the LDP Initialization Message.  If
    // the value is generic(1) then no optional parameters will be sent in the LDP
    // Initialization message associated with this Entity.  If the value is
    // atmParameters(2) then a row must be created in the mplsLdpEntityAtmTable,
    // which corresponds to this entry.  If the value is frameRelayParameters(3)
    // then a row must be created in the mplsLdpEntityFrameRelayTable, which
    // corresponds to this entry. The type is MplsLdpLabelType.
    MplsLdpEntityLabelType interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entity's counters suffered a discontinuity.  The relevant counters
    // are the specific instances associated with this entity of any Counter32
    // object contained in the 'mplsLdpEntityStatsTable'.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    MplsLdpEntityDiscontinuityTime interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    MplsLdpEntityStorageType interface{}

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
    MplsLdpEntityRowStatus interface{}

    // A count of the Session Initialization messages which were sent or received
    // by this LDP Entity and were NAK'd.   In other words, this counter counts
    // the number of session initializations that failed.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsSessionAttempts interface{}

    // A count of the Session Rejected/No Hello Error Notification Messages sent
    // or received by this LDP Entity.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of mplsLdpEntityDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    MplsLdpEntityStatsSessionRejectedNoHelloErrors interface{}

    // A count of the Session Rejected/Parameters Advertisement Mode Error
    // Notification Messages sent or received by this LDP Entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsSessionRejectedAdErrors interface{}

    // A count of the Session Rejected/Parameters  Max Pdu Length Error
    // Notification Messages sent or received by this LDP Entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsSessionRejectedMaxPduErrors interface{}

    // A count of the Session Rejected/Parameters Label Range Notification
    // Messages sent or received by this LDP Entity.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsSessionRejectedLRErrors interface{}

    // This object counts the number of Bad LDP Identifier Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsBadLdpIdentifierErrors interface{}

    // This object counts the number of Bad PDU Length Fatal Errors detected by
    // the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsBadPduLengthErrors interface{}

    // This object counts the number of Bad Message Length Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsBadMessageLengthErrors interface{}

    // This object counts the number of Bad TLV Length Fatal Errors detected by
    // the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsBadTlvLengthErrors interface{}

    // This object counts the number of Malformed TLV Value Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsMalformedTlvValueErrors interface{}

    // This object counts the number of Session Keep Alive Timer Expired Errors
    // detected by the session(s) (past and present) associated with this LDP
    // Entity.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of mplsLdpEntityDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    MplsLdpEntityStatsKeepAliveTimerExpErrors interface{}

    // This object counts the number of Shutdown Notifications received related to
    // session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsShutdownReceivedNotifications interface{}

    // This object counts the number of Shutdown Notfications sent related to
    // session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of  
    // mplsLdpEntityDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsLdpEntityStatsShutdownSentNotifications interface{}
}

func (mplsLdpEntityEntry *MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry) GetEntityData() *types.CommonEntityData {
    mplsLdpEntityEntry.EntityData.YFilter = mplsLdpEntityEntry.YFilter
    mplsLdpEntityEntry.EntityData.YangName = "mplsLdpEntityEntry"
    mplsLdpEntityEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpEntityEntry.EntityData.ParentYangName = "mplsLdpEntityTable"
    mplsLdpEntityEntry.EntityData.SegmentPath = "mplsLdpEntityEntry" + types.AddKeyToken(mplsLdpEntityEntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsLdpEntityEntry.MplsLdpEntityIndex, "mplsLdpEntityIndex")
    mplsLdpEntityEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsLdpEntityTable/" + mplsLdpEntityEntry.EntityData.SegmentPath
    mplsLdpEntityEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpEntityEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpEntityEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpEntityEntry.EntityData.Children = types.NewOrderedMap()
    mplsLdpEntityEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsLdpEntityEntry.MplsLdpEntityLdpId})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsLdpEntityEntry.MplsLdpEntityIndex})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityProtocolVersion", types.YLeaf{"MplsLdpEntityProtocolVersion", mplsLdpEntityEntry.MplsLdpEntityProtocolVersion})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityAdminStatus", types.YLeaf{"MplsLdpEntityAdminStatus", mplsLdpEntityEntry.MplsLdpEntityAdminStatus})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityOperStatus", types.YLeaf{"MplsLdpEntityOperStatus", mplsLdpEntityEntry.MplsLdpEntityOperStatus})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityTcpPort", types.YLeaf{"MplsLdpEntityTcpPort", mplsLdpEntityEntry.MplsLdpEntityTcpPort})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityUdpDscPort", types.YLeaf{"MplsLdpEntityUdpDscPort", mplsLdpEntityEntry.MplsLdpEntityUdpDscPort})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityMaxPduLength", types.YLeaf{"MplsLdpEntityMaxPduLength", mplsLdpEntityEntry.MplsLdpEntityMaxPduLength})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityKeepAliveHoldTimer", types.YLeaf{"MplsLdpEntityKeepAliveHoldTimer", mplsLdpEntityEntry.MplsLdpEntityKeepAliveHoldTimer})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityHelloHoldTimer", types.YLeaf{"MplsLdpEntityHelloHoldTimer", mplsLdpEntityEntry.MplsLdpEntityHelloHoldTimer})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityInitSessionThreshold", types.YLeaf{"MplsLdpEntityInitSessionThreshold", mplsLdpEntityEntry.MplsLdpEntityInitSessionThreshold})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityLabelDistMethod", types.YLeaf{"MplsLdpEntityLabelDistMethod", mplsLdpEntityEntry.MplsLdpEntityLabelDistMethod})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityLabelRetentionMode", types.YLeaf{"MplsLdpEntityLabelRetentionMode", mplsLdpEntityEntry.MplsLdpEntityLabelRetentionMode})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityPathVectorLimit", types.YLeaf{"MplsLdpEntityPathVectorLimit", mplsLdpEntityEntry.MplsLdpEntityPathVectorLimit})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityHopCountLimit", types.YLeaf{"MplsLdpEntityHopCountLimit", mplsLdpEntityEntry.MplsLdpEntityHopCountLimit})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityTransportAddrKind", types.YLeaf{"MplsLdpEntityTransportAddrKind", mplsLdpEntityEntry.MplsLdpEntityTransportAddrKind})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityTargetPeer", types.YLeaf{"MplsLdpEntityTargetPeer", mplsLdpEntityEntry.MplsLdpEntityTargetPeer})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityTargetPeerAddrType", types.YLeaf{"MplsLdpEntityTargetPeerAddrType", mplsLdpEntityEntry.MplsLdpEntityTargetPeerAddrType})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityTargetPeerAddr", types.YLeaf{"MplsLdpEntityTargetPeerAddr", mplsLdpEntityEntry.MplsLdpEntityTargetPeerAddr})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityLabelType", types.YLeaf{"MplsLdpEntityLabelType", mplsLdpEntityEntry.MplsLdpEntityLabelType})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityDiscontinuityTime", types.YLeaf{"MplsLdpEntityDiscontinuityTime", mplsLdpEntityEntry.MplsLdpEntityDiscontinuityTime})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStorageType", types.YLeaf{"MplsLdpEntityStorageType", mplsLdpEntityEntry.MplsLdpEntityStorageType})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityRowStatus", types.YLeaf{"MplsLdpEntityRowStatus", mplsLdpEntityEntry.MplsLdpEntityRowStatus})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsSessionAttempts", types.YLeaf{"MplsLdpEntityStatsSessionAttempts", mplsLdpEntityEntry.MplsLdpEntityStatsSessionAttempts})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsSessionRejectedNoHelloErrors", types.YLeaf{"MplsLdpEntityStatsSessionRejectedNoHelloErrors", mplsLdpEntityEntry.MplsLdpEntityStatsSessionRejectedNoHelloErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsSessionRejectedAdErrors", types.YLeaf{"MplsLdpEntityStatsSessionRejectedAdErrors", mplsLdpEntityEntry.MplsLdpEntityStatsSessionRejectedAdErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsSessionRejectedMaxPduErrors", types.YLeaf{"MplsLdpEntityStatsSessionRejectedMaxPduErrors", mplsLdpEntityEntry.MplsLdpEntityStatsSessionRejectedMaxPduErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsSessionRejectedLRErrors", types.YLeaf{"MplsLdpEntityStatsSessionRejectedLRErrors", mplsLdpEntityEntry.MplsLdpEntityStatsSessionRejectedLRErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsBadLdpIdentifierErrors", types.YLeaf{"MplsLdpEntityStatsBadLdpIdentifierErrors", mplsLdpEntityEntry.MplsLdpEntityStatsBadLdpIdentifierErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsBadPduLengthErrors", types.YLeaf{"MplsLdpEntityStatsBadPduLengthErrors", mplsLdpEntityEntry.MplsLdpEntityStatsBadPduLengthErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsBadMessageLengthErrors", types.YLeaf{"MplsLdpEntityStatsBadMessageLengthErrors", mplsLdpEntityEntry.MplsLdpEntityStatsBadMessageLengthErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsBadTlvLengthErrors", types.YLeaf{"MplsLdpEntityStatsBadTlvLengthErrors", mplsLdpEntityEntry.MplsLdpEntityStatsBadTlvLengthErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsMalformedTlvValueErrors", types.YLeaf{"MplsLdpEntityStatsMalformedTlvValueErrors", mplsLdpEntityEntry.MplsLdpEntityStatsMalformedTlvValueErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsKeepAliveTimerExpErrors", types.YLeaf{"MplsLdpEntityStatsKeepAliveTimerExpErrors", mplsLdpEntityEntry.MplsLdpEntityStatsKeepAliveTimerExpErrors})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsShutdownReceivedNotifications", types.YLeaf{"MplsLdpEntityStatsShutdownReceivedNotifications", mplsLdpEntityEntry.MplsLdpEntityStatsShutdownReceivedNotifications})
    mplsLdpEntityEntry.EntityData.Leafs.Append("mplsLdpEntityStatsShutdownSentNotifications", types.YLeaf{"MplsLdpEntityStatsShutdownSentNotifications", mplsLdpEntityEntry.MplsLdpEntityStatsShutdownSentNotifications})

    mplsLdpEntityEntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex"}

    return &(mplsLdpEntityEntry.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityAdminStatus represents with the Peer.
type MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityAdminStatus string

const (
    MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityAdminStatus_enable MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityAdminStatus = "enable"

    MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityAdminStatus_disable MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityAdminStatus = "disable"
)

// MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus represents to enabled(2) or disabled(3).
type MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus string

const (
    MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus_unknown MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus = "unknown"

    MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus_enabled MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus = "enabled"

    MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus_disabled MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityOperStatus = "disabled"
)

// MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityTransportAddrKind represents transport address in the hello message.
type MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityTransportAddrKind string

const (
    MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityTransportAddrKind_interface_ MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityTransportAddrKind = "interface"

    MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityTransportAddrKind_loopback MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityTransportAddrKind = "loopback"
)

// MPLSLDPSTDMIB_MplsLdpPeerTable
// Information about LDP peers known by Entities in
// the mplsLdpEntityTable.  The information in this table
// is based on information from the Entity-Peer interaction
// during session initialization but is not appropriate
// for the mplsLdpSessionTable, because objects in this
// table may or may not be used in session establishment.
type MPLSLDPSTDMIB_MplsLdpPeerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single Peer which is related to a Session.  This table
    // is augmented by the mplsLdpSessionTable. The type is slice of
    // MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry.
    MplsLdpPeerEntry []*MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry
}

func (mplsLdpPeerTable *MPLSLDPSTDMIB_MplsLdpPeerTable) GetEntityData() *types.CommonEntityData {
    mplsLdpPeerTable.EntityData.YFilter = mplsLdpPeerTable.YFilter
    mplsLdpPeerTable.EntityData.YangName = "mplsLdpPeerTable"
    mplsLdpPeerTable.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpPeerTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpPeerTable.EntityData.SegmentPath = "mplsLdpPeerTable"
    mplsLdpPeerTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpPeerTable.EntityData.SegmentPath
    mplsLdpPeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpPeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpPeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpPeerTable.EntityData.Children = types.NewOrderedMap()
    mplsLdpPeerTable.EntityData.Children.Append("mplsLdpPeerEntry", types.YChild{"MplsLdpPeerEntry", nil})
    for i := range mplsLdpPeerTable.MplsLdpPeerEntry {
        mplsLdpPeerTable.EntityData.Children.Append(types.GetSegmentPath(mplsLdpPeerTable.MplsLdpPeerEntry[i]), types.YChild{"MplsLdpPeerEntry", mplsLdpPeerTable.MplsLdpPeerEntry[i]})
    }
    mplsLdpPeerTable.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpPeerTable.EntityData.YListKeys = []string {}

    return &(mplsLdpPeerTable.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry
// Information about a single Peer which is related
// to a Session.  This table is augmented by
// the mplsLdpSessionTable.
type MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityLdpId
    MplsLdpEntityLdpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityIndex
    MplsLdpEntityIndex interface{}

    // This attribute is a key. The LDP identifier of this LDP Peer. The type is
    // string.
    MplsLdpPeerLdpId interface{}

    // For any given LDP session, the method of label distribution must be
    // specified. The type is MplsLabelDistributionMethod.
    MplsLdpPeerLabelDistMethod interface{}

    // If the value of this object is 0 (zero) then Loop Dection for Path Vectors
    // for this Peer is disabled.  Otherwise, if this object has a value greater
    // than zero, then Loop Dection for Path  Vectors for this Peer is enabled and
    // the Path Vector Limit is this value. The type is interface{} with range:
    // 0..255.
    MplsLdpPeerPathVectorLimit interface{}

    // The type of the Internet address for the mplsLdpPeerTransportAddr object. 
    // The LDP specification describes this as being either an IPv4 Transport
    // Address or IPv6 Transport   Address which is used in opening the LDP
    // session's TCP connection, or if the optional TLV is not present, then this
    // is the IPv4/IPv6 source address for the UPD packet carrying the Hellos. 
    // This object specifies how the value of the mplsLdpPeerTransportAddr object
    // should be interpreted. The type is InetAddressType.
    MplsLdpPeerTransportAddrType interface{}

    // The Internet address advertised by the peer in the Hello Message or the
    // Hello source address.  The type of this address is specified by the value
    // of the mplsLdpPeerTransportAddrType object. The type is string with length:
    // 0..255.
    MplsLdpPeerTransportAddr interface{}

    // The value of sysUpTime at the time this Session entered its current state
    // as denoted by the mplsLdpSessionState object. The type is interface{} with
    // range: 0..4294967295.
    MplsLdpSessionStateLastChange interface{}

    // The current state of the session, all of the states 1 to 5 are based on the
    // state machine for session negotiation behavior. The type is
    // MplsLdpSessionState.
    MplsLdpSessionState interface{}

    // During session establishment the LSR/LER takes either the active role or
    // the passive role based on address comparisons.  This object indicates
    // whether this LSR/LER was behaving in an active role or passive role during
    // this session's establishment.  The value of unknown(1), indicates that the
    // role is not able to be determined at the present time. The type is
    // MplsLdpSessionRole.
    MplsLdpSessionRole interface{}

    // The version of the LDP Protocol which this session is using.  This is the
    // version of   the LDP protocol which has been negotiated during session
    // initialization. The type is interface{} with range: 1..65535.
    MplsLdpSessionProtocolVersion interface{}

    // The keep alive hold time remaining for this session. The type is
    // interface{} with range: 0..2147483647.
    MplsLdpSessionKeepAliveHoldTimeRem interface{}

    // The negotiated KeepAlive Time which represents the amount of seconds
    // between keep alive messages.  The mplsLdpEntityKeepAliveHoldTimer related
    // to this Session is the value that was proposed as the KeepAlive Time for
    // this session.  This value is negotiated during session initialization
    // between the entity's proposed value (i.e., the value configured in
    // mplsLdpEntityKeepAliveHoldTimer) and the peer's proposed KeepAlive Hold
    // Timer value. This value is the smaller of the two proposed values. The type
    // is interface{} with range: 1..65535. Units are seconds.
    MplsLdpSessionKeepAliveTime interface{}

    // The value of maximum allowable length for LDP PDUs for this session.  This
    // value may have been negotiated during the Session Initialization.  This
    // object is related to the mplsLdpEntityMaxPduLength object.  The
    // mplsLdpEntityMaxPduLength object specifies the requested LDP PDU length,
    // and this object reflects the negotiated LDP PDU length between the Entity
    // and the Peer. The type is interface{} with range: 1..65535. Units are
    // octets.
    MplsLdpSessionMaxPduLength interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this session's counters suffered a discontinuity.  The relevant counters
    // are the specific instances associated with this session of any Counter32
    // object contained in the mplsLdpSessionStatsTable.  The initial value of
    // this object is the value of sysUpTime when the entry was created in this
    // table.  Also, a command generator can distinguish when a session between a
    // given Entity and Peer goes away and a new session is established.  This
    // value would change and thus indicate to the command generator that this is
    // a different session. The type is interface{} with range: 0..4294967295.
    MplsLdpSessionDiscontinuityTime interface{}

    // This object counts the number of Unknown Message Type Errors detected by
    // this LSR/LER during this session.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of mplsLdpSessionDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    MplsLdpSessionStatsUnknownMesTypeErrors interface{}

    // This object counts the number of Unknown TLV Errors detected by this
    // LSR/LER during this session.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of mplsLdpSessionDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    MplsLdpSessionStatsUnknownTlvErrors interface{}
}

func (mplsLdpPeerEntry *MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry) GetEntityData() *types.CommonEntityData {
    mplsLdpPeerEntry.EntityData.YFilter = mplsLdpPeerEntry.YFilter
    mplsLdpPeerEntry.EntityData.YangName = "mplsLdpPeerEntry"
    mplsLdpPeerEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpPeerEntry.EntityData.ParentYangName = "mplsLdpPeerTable"
    mplsLdpPeerEntry.EntityData.SegmentPath = "mplsLdpPeerEntry" + types.AddKeyToken(mplsLdpPeerEntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsLdpPeerEntry.MplsLdpEntityIndex, "mplsLdpEntityIndex") + types.AddKeyToken(mplsLdpPeerEntry.MplsLdpPeerLdpId, "mplsLdpPeerLdpId")
    mplsLdpPeerEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsLdpPeerTable/" + mplsLdpPeerEntry.EntityData.SegmentPath
    mplsLdpPeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpPeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpPeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpPeerEntry.EntityData.Children = types.NewOrderedMap()
    mplsLdpPeerEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsLdpPeerEntry.MplsLdpEntityLdpId})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsLdpPeerEntry.MplsLdpEntityIndex})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpPeerLdpId", types.YLeaf{"MplsLdpPeerLdpId", mplsLdpPeerEntry.MplsLdpPeerLdpId})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpPeerLabelDistMethod", types.YLeaf{"MplsLdpPeerLabelDistMethod", mplsLdpPeerEntry.MplsLdpPeerLabelDistMethod})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpPeerPathVectorLimit", types.YLeaf{"MplsLdpPeerPathVectorLimit", mplsLdpPeerEntry.MplsLdpPeerPathVectorLimit})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpPeerTransportAddrType", types.YLeaf{"MplsLdpPeerTransportAddrType", mplsLdpPeerEntry.MplsLdpPeerTransportAddrType})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpPeerTransportAddr", types.YLeaf{"MplsLdpPeerTransportAddr", mplsLdpPeerEntry.MplsLdpPeerTransportAddr})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionStateLastChange", types.YLeaf{"MplsLdpSessionStateLastChange", mplsLdpPeerEntry.MplsLdpSessionStateLastChange})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionState", types.YLeaf{"MplsLdpSessionState", mplsLdpPeerEntry.MplsLdpSessionState})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionRole", types.YLeaf{"MplsLdpSessionRole", mplsLdpPeerEntry.MplsLdpSessionRole})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionProtocolVersion", types.YLeaf{"MplsLdpSessionProtocolVersion", mplsLdpPeerEntry.MplsLdpSessionProtocolVersion})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionKeepAliveHoldTimeRem", types.YLeaf{"MplsLdpSessionKeepAliveHoldTimeRem", mplsLdpPeerEntry.MplsLdpSessionKeepAliveHoldTimeRem})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionKeepAliveTime", types.YLeaf{"MplsLdpSessionKeepAliveTime", mplsLdpPeerEntry.MplsLdpSessionKeepAliveTime})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionMaxPduLength", types.YLeaf{"MplsLdpSessionMaxPduLength", mplsLdpPeerEntry.MplsLdpSessionMaxPduLength})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionDiscontinuityTime", types.YLeaf{"MplsLdpSessionDiscontinuityTime", mplsLdpPeerEntry.MplsLdpSessionDiscontinuityTime})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionStatsUnknownMesTypeErrors", types.YLeaf{"MplsLdpSessionStatsUnknownMesTypeErrors", mplsLdpPeerEntry.MplsLdpSessionStatsUnknownMesTypeErrors})
    mplsLdpPeerEntry.EntityData.Leafs.Append("mplsLdpSessionStatsUnknownTlvErrors", types.YLeaf{"MplsLdpSessionStatsUnknownTlvErrors", mplsLdpPeerEntry.MplsLdpSessionStatsUnknownTlvErrors})

    mplsLdpPeerEntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex", "MplsLdpPeerLdpId"}

    return &(mplsLdpPeerEntry.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole represents able to be determined at the present time.
type MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole string

const (
    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole_unknown MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole = "unknown"

    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole_active MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole = "active"

    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole_passive MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionRole = "passive"
)

// MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState represents for session negotiation behavior.
type MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState string

const (
    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState_nonexistent MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState = "nonexistent"

    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState_initialized MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState = "initialized"

    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState_openrec MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState = "openrec"

    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState_opensent MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState = "opensent"

    MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState_operational MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpSessionState = "operational"
)

// MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable
// A table of Hello Adjacencies for Sessions.
type MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row represents a single LDP Hello Adjacency. An LDP Session can have
    // one or more Hello Adjacencies. The type is slice of
    // MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry.
    MplsLdpHelloAdjacencyEntry []*MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry
}

func (mplsLdpHelloAdjacencyTable *MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable) GetEntityData() *types.CommonEntityData {
    mplsLdpHelloAdjacencyTable.EntityData.YFilter = mplsLdpHelloAdjacencyTable.YFilter
    mplsLdpHelloAdjacencyTable.EntityData.YangName = "mplsLdpHelloAdjacencyTable"
    mplsLdpHelloAdjacencyTable.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpHelloAdjacencyTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpHelloAdjacencyTable.EntityData.SegmentPath = "mplsLdpHelloAdjacencyTable"
    mplsLdpHelloAdjacencyTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpHelloAdjacencyTable.EntityData.SegmentPath
    mplsLdpHelloAdjacencyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpHelloAdjacencyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpHelloAdjacencyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpHelloAdjacencyTable.EntityData.Children = types.NewOrderedMap()
    mplsLdpHelloAdjacencyTable.EntityData.Children.Append("mplsLdpHelloAdjacencyEntry", types.YChild{"MplsLdpHelloAdjacencyEntry", nil})
    for i := range mplsLdpHelloAdjacencyTable.MplsLdpHelloAdjacencyEntry {
        mplsLdpHelloAdjacencyTable.EntityData.Children.Append(types.GetSegmentPath(mplsLdpHelloAdjacencyTable.MplsLdpHelloAdjacencyEntry[i]), types.YChild{"MplsLdpHelloAdjacencyEntry", mplsLdpHelloAdjacencyTable.MplsLdpHelloAdjacencyEntry[i]})
    }
    mplsLdpHelloAdjacencyTable.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpHelloAdjacencyTable.EntityData.YListKeys = []string {}

    return &(mplsLdpHelloAdjacencyTable.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry
// Each row represents a single LDP Hello Adjacency.
// An LDP Session can have one or more Hello
// Adjacencies.
type MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityLdpId
    MplsLdpEntityLdpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityIndex
    MplsLdpEntityIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpPeerLdpId
    MplsLdpPeerLdpId interface{}

    // This attribute is a key. An identifier for this specific adjacency. The
    // type is interface{} with range: 1..4294967295.
    MplsLdpHelloAdjacencyIndex interface{}

    // If the value of this object is 65535, this means that the hold time is
    // infinite (i.e., wait forever).  Otherwise, the time remaining for this
    // Hello Adjacency to receive its next Hello Message.  This interval will
    // change when the 'next' Hello Message which corresponds to this Hello
    // Adjacency is received unless it is infinite. The type is interface{} with
    // range: 0..2147483647. Units are seconds.
    MplsLdpHelloAdjacencyHoldTimeRem interface{}

    // The Hello hold time which is negotiated between the Entity and the Peer. 
    // The entity associated with this Hello Adjacency issues a proposed Hello
    // Hold Time value in the mplsLdpEntityHelloHoldTimer object.  The peer also
    // proposes a value and this object represents the negotiated value.  A value
    // of 0 means the default, which is 15 seconds for Link Hellos and 45 seconds
    // for Targeted Hellos. A value of 65535 indicates an infinite hold time. The
    // type is interface{} with range: 0..65535.
    MplsLdpHelloAdjacencyHoldTime interface{}

    // This adjacency is the result of a 'link' hello if the value of this object
    // is link(1).   Otherwise, it is a result of a 'targeted' hello, targeted(2).
    // The type is MplsLdpHelloAdjacencyType.
    MplsLdpHelloAdjacencyType interface{}
}

func (mplsLdpHelloAdjacencyEntry *MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry) GetEntityData() *types.CommonEntityData {
    mplsLdpHelloAdjacencyEntry.EntityData.YFilter = mplsLdpHelloAdjacencyEntry.YFilter
    mplsLdpHelloAdjacencyEntry.EntityData.YangName = "mplsLdpHelloAdjacencyEntry"
    mplsLdpHelloAdjacencyEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpHelloAdjacencyEntry.EntityData.ParentYangName = "mplsLdpHelloAdjacencyTable"
    mplsLdpHelloAdjacencyEntry.EntityData.SegmentPath = "mplsLdpHelloAdjacencyEntry" + types.AddKeyToken(mplsLdpHelloAdjacencyEntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsLdpHelloAdjacencyEntry.MplsLdpEntityIndex, "mplsLdpEntityIndex") + types.AddKeyToken(mplsLdpHelloAdjacencyEntry.MplsLdpPeerLdpId, "mplsLdpPeerLdpId") + types.AddKeyToken(mplsLdpHelloAdjacencyEntry.MplsLdpHelloAdjacencyIndex, "mplsLdpHelloAdjacencyIndex")
    mplsLdpHelloAdjacencyEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsLdpHelloAdjacencyTable/" + mplsLdpHelloAdjacencyEntry.EntityData.SegmentPath
    mplsLdpHelloAdjacencyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpHelloAdjacencyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpHelloAdjacencyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpHelloAdjacencyEntry.EntityData.Children = types.NewOrderedMap()
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsLdpHelloAdjacencyEntry.MplsLdpEntityLdpId})
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsLdpHelloAdjacencyEntry.MplsLdpEntityIndex})
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs.Append("mplsLdpPeerLdpId", types.YLeaf{"MplsLdpPeerLdpId", mplsLdpHelloAdjacencyEntry.MplsLdpPeerLdpId})
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs.Append("mplsLdpHelloAdjacencyIndex", types.YLeaf{"MplsLdpHelloAdjacencyIndex", mplsLdpHelloAdjacencyEntry.MplsLdpHelloAdjacencyIndex})
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs.Append("mplsLdpHelloAdjacencyHoldTimeRem", types.YLeaf{"MplsLdpHelloAdjacencyHoldTimeRem", mplsLdpHelloAdjacencyEntry.MplsLdpHelloAdjacencyHoldTimeRem})
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs.Append("mplsLdpHelloAdjacencyHoldTime", types.YLeaf{"MplsLdpHelloAdjacencyHoldTime", mplsLdpHelloAdjacencyEntry.MplsLdpHelloAdjacencyHoldTime})
    mplsLdpHelloAdjacencyEntry.EntityData.Leafs.Append("mplsLdpHelloAdjacencyType", types.YLeaf{"MplsLdpHelloAdjacencyType", mplsLdpHelloAdjacencyEntry.MplsLdpHelloAdjacencyType})

    mplsLdpHelloAdjacencyEntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex", "MplsLdpPeerLdpId", "MplsLdpHelloAdjacencyIndex"}

    return &(mplsLdpHelloAdjacencyEntry.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry_MplsLdpHelloAdjacencyType represents hello, targeted(2).
type MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry_MplsLdpHelloAdjacencyType string

const (
    MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry_MplsLdpHelloAdjacencyType_link MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry_MplsLdpHelloAdjacencyType = "link"

    MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry_MplsLdpHelloAdjacencyType_targeted MPLSLDPSTDMIB_MplsLdpHelloAdjacencyTable_MplsLdpHelloAdjacencyEntry_MplsLdpHelloAdjacencyType = "targeted"
)

// MPLSLDPSTDMIB_MplsInSegmentLdpLspTable
// A table of LDP LSP's which
// map to the mplsInSegmentTable in the
// MPLS-LSR-STD-MIB module.
type MPLSLDPSTDMIB_MplsInSegmentLdpLspTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents information on a single LDP LSP which is
    // represented by a session's index triple (mplsLdpEntityLdpId,
    // mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the index for the
    // mplsInSegmentTable (mplsInSegmentLdpLspLabelIndex) from the
    // MPLS-LSR-STD-MIB.  The information contained in a row is read-only. The
    // type is slice of
    // MPLSLDPSTDMIB_MplsInSegmentLdpLspTable_MplsInSegmentLdpLspEntry.
    MplsInSegmentLdpLspEntry []*MPLSLDPSTDMIB_MplsInSegmentLdpLspTable_MplsInSegmentLdpLspEntry
}

func (mplsInSegmentLdpLspTable *MPLSLDPSTDMIB_MplsInSegmentLdpLspTable) GetEntityData() *types.CommonEntityData {
    mplsInSegmentLdpLspTable.EntityData.YFilter = mplsInSegmentLdpLspTable.YFilter
    mplsInSegmentLdpLspTable.EntityData.YangName = "mplsInSegmentLdpLspTable"
    mplsInSegmentLdpLspTable.EntityData.BundleName = "cisco_ios_xe"
    mplsInSegmentLdpLspTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsInSegmentLdpLspTable.EntityData.SegmentPath = "mplsInSegmentLdpLspTable"
    mplsInSegmentLdpLspTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsInSegmentLdpLspTable.EntityData.SegmentPath
    mplsInSegmentLdpLspTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInSegmentLdpLspTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInSegmentLdpLspTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInSegmentLdpLspTable.EntityData.Children = types.NewOrderedMap()
    mplsInSegmentLdpLspTable.EntityData.Children.Append("mplsInSegmentLdpLspEntry", types.YChild{"MplsInSegmentLdpLspEntry", nil})
    for i := range mplsInSegmentLdpLspTable.MplsInSegmentLdpLspEntry {
        mplsInSegmentLdpLspTable.EntityData.Children.Append(types.GetSegmentPath(mplsInSegmentLdpLspTable.MplsInSegmentLdpLspEntry[i]), types.YChild{"MplsInSegmentLdpLspEntry", mplsInSegmentLdpLspTable.MplsInSegmentLdpLspEntry[i]})
    }
    mplsInSegmentLdpLspTable.EntityData.Leafs = types.NewOrderedMap()

    mplsInSegmentLdpLspTable.EntityData.YListKeys = []string {}

    return &(mplsInSegmentLdpLspTable.EntityData)
}

// MPLSLDPSTDMIB_MplsInSegmentLdpLspTable_MplsInSegmentLdpLspEntry
// An entry in this table represents information
// on a single LDP LSP which is represented by
// a session's index triple (mplsLdpEntityLdpId,
// mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the
// index for the mplsInSegmentTable
// (mplsInSegmentLdpLspLabelIndex) from the
// MPLS-LSR-STD-MIB.
// 
// The information contained in a row is read-only.
type MPLSLDPSTDMIB_MplsInSegmentLdpLspTable_MplsInSegmentLdpLspEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityLdpId
    MplsLdpEntityLdpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityIndex
    MplsLdpEntityIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpPeerLdpId
    MplsLdpPeerLdpId interface{}

    // This attribute is a key. This contains the same value as the
    // mplsInSegmentIndex in the MPLS-LSR-STD-MIB's mplsInSegmentTable. The type
    // is string with length: 1..24.
    MplsInSegmentLdpLspIndex interface{}

    // The Layer 2 Label Type. The type is MplsLdpLabelType.
    MplsInSegmentLdpLspLabelType interface{}

    // The type of LSP connection. The type is MplsLspType.
    MplsInSegmentLdpLspType interface{}
}

func (mplsInSegmentLdpLspEntry *MPLSLDPSTDMIB_MplsInSegmentLdpLspTable_MplsInSegmentLdpLspEntry) GetEntityData() *types.CommonEntityData {
    mplsInSegmentLdpLspEntry.EntityData.YFilter = mplsInSegmentLdpLspEntry.YFilter
    mplsInSegmentLdpLspEntry.EntityData.YangName = "mplsInSegmentLdpLspEntry"
    mplsInSegmentLdpLspEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsInSegmentLdpLspEntry.EntityData.ParentYangName = "mplsInSegmentLdpLspTable"
    mplsInSegmentLdpLspEntry.EntityData.SegmentPath = "mplsInSegmentLdpLspEntry" + types.AddKeyToken(mplsInSegmentLdpLspEntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsInSegmentLdpLspEntry.MplsLdpEntityIndex, "mplsLdpEntityIndex") + types.AddKeyToken(mplsInSegmentLdpLspEntry.MplsLdpPeerLdpId, "mplsLdpPeerLdpId") + types.AddKeyToken(mplsInSegmentLdpLspEntry.MplsInSegmentLdpLspIndex, "mplsInSegmentLdpLspIndex")
    mplsInSegmentLdpLspEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsInSegmentLdpLspTable/" + mplsInSegmentLdpLspEntry.EntityData.SegmentPath
    mplsInSegmentLdpLspEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInSegmentLdpLspEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInSegmentLdpLspEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInSegmentLdpLspEntry.EntityData.Children = types.NewOrderedMap()
    mplsInSegmentLdpLspEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsInSegmentLdpLspEntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsInSegmentLdpLspEntry.MplsLdpEntityLdpId})
    mplsInSegmentLdpLspEntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsInSegmentLdpLspEntry.MplsLdpEntityIndex})
    mplsInSegmentLdpLspEntry.EntityData.Leafs.Append("mplsLdpPeerLdpId", types.YLeaf{"MplsLdpPeerLdpId", mplsInSegmentLdpLspEntry.MplsLdpPeerLdpId})
    mplsInSegmentLdpLspEntry.EntityData.Leafs.Append("mplsInSegmentLdpLspIndex", types.YLeaf{"MplsInSegmentLdpLspIndex", mplsInSegmentLdpLspEntry.MplsInSegmentLdpLspIndex})
    mplsInSegmentLdpLspEntry.EntityData.Leafs.Append("mplsInSegmentLdpLspLabelType", types.YLeaf{"MplsInSegmentLdpLspLabelType", mplsInSegmentLdpLspEntry.MplsInSegmentLdpLspLabelType})
    mplsInSegmentLdpLspEntry.EntityData.Leafs.Append("mplsInSegmentLdpLspType", types.YLeaf{"MplsInSegmentLdpLspType", mplsInSegmentLdpLspEntry.MplsInSegmentLdpLspType})

    mplsInSegmentLdpLspEntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex", "MplsLdpPeerLdpId", "MplsInSegmentLdpLspIndex"}

    return &(mplsInSegmentLdpLspEntry.EntityData)
}

// MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable
// A table of LDP LSP's which
// map to the mplsOutSegmentTable in the
// MPLS-LSR-STD-MIB.
type MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents information on a single LDP LSP which is
    // represented by a session's index triple (mplsLdpEntityLdpId,
    // mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the index
    // (mplsOutSegmentLdpLspIndex) for the mplsOutSegmentTable.  The information
    // contained in a row is read-only. The type is slice of
    // MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable_MplsOutSegmentLdpLspEntry.
    MplsOutSegmentLdpLspEntry []*MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable_MplsOutSegmentLdpLspEntry
}

func (mplsOutSegmentLdpLspTable *MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable) GetEntityData() *types.CommonEntityData {
    mplsOutSegmentLdpLspTable.EntityData.YFilter = mplsOutSegmentLdpLspTable.YFilter
    mplsOutSegmentLdpLspTable.EntityData.YangName = "mplsOutSegmentLdpLspTable"
    mplsOutSegmentLdpLspTable.EntityData.BundleName = "cisco_ios_xe"
    mplsOutSegmentLdpLspTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsOutSegmentLdpLspTable.EntityData.SegmentPath = "mplsOutSegmentLdpLspTable"
    mplsOutSegmentLdpLspTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsOutSegmentLdpLspTable.EntityData.SegmentPath
    mplsOutSegmentLdpLspTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsOutSegmentLdpLspTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsOutSegmentLdpLspTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsOutSegmentLdpLspTable.EntityData.Children = types.NewOrderedMap()
    mplsOutSegmentLdpLspTable.EntityData.Children.Append("mplsOutSegmentLdpLspEntry", types.YChild{"MplsOutSegmentLdpLspEntry", nil})
    for i := range mplsOutSegmentLdpLspTable.MplsOutSegmentLdpLspEntry {
        mplsOutSegmentLdpLspTable.EntityData.Children.Append(types.GetSegmentPath(mplsOutSegmentLdpLspTable.MplsOutSegmentLdpLspEntry[i]), types.YChild{"MplsOutSegmentLdpLspEntry", mplsOutSegmentLdpLspTable.MplsOutSegmentLdpLspEntry[i]})
    }
    mplsOutSegmentLdpLspTable.EntityData.Leafs = types.NewOrderedMap()

    mplsOutSegmentLdpLspTable.EntityData.YListKeys = []string {}

    return &(mplsOutSegmentLdpLspTable.EntityData)
}

// MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable_MplsOutSegmentLdpLspEntry
// An entry in this table represents information
// on a single LDP LSP which is represented by
// a session's index triple (mplsLdpEntityLdpId,
// mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the
// index (mplsOutSegmentLdpLspIndex)
// for the mplsOutSegmentTable.
// 
// The information contained in a row is read-only.
type MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable_MplsOutSegmentLdpLspEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityLdpId
    MplsLdpEntityLdpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityIndex
    MplsLdpEntityIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpPeerLdpId
    MplsLdpPeerLdpId interface{}

    // This attribute is a key. This contains the same value as the
    // mplsOutSegmentIndex in the MPLS-LSR-STD-MIB's mplsOutSegmentTable. The type
    // is string with length: 1..24.
    MplsOutSegmentLdpLspIndex interface{}

    // The Layer 2 Label Type. The type is MplsLdpLabelType.
    MplsOutSegmentLdpLspLabelType interface{}

    // The type of LSP connection. The type is MplsLspType.
    MplsOutSegmentLdpLspType interface{}
}

func (mplsOutSegmentLdpLspEntry *MPLSLDPSTDMIB_MplsOutSegmentLdpLspTable_MplsOutSegmentLdpLspEntry) GetEntityData() *types.CommonEntityData {
    mplsOutSegmentLdpLspEntry.EntityData.YFilter = mplsOutSegmentLdpLspEntry.YFilter
    mplsOutSegmentLdpLspEntry.EntityData.YangName = "mplsOutSegmentLdpLspEntry"
    mplsOutSegmentLdpLspEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsOutSegmentLdpLspEntry.EntityData.ParentYangName = "mplsOutSegmentLdpLspTable"
    mplsOutSegmentLdpLspEntry.EntityData.SegmentPath = "mplsOutSegmentLdpLspEntry" + types.AddKeyToken(mplsOutSegmentLdpLspEntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsOutSegmentLdpLspEntry.MplsLdpEntityIndex, "mplsLdpEntityIndex") + types.AddKeyToken(mplsOutSegmentLdpLspEntry.MplsLdpPeerLdpId, "mplsLdpPeerLdpId") + types.AddKeyToken(mplsOutSegmentLdpLspEntry.MplsOutSegmentLdpLspIndex, "mplsOutSegmentLdpLspIndex")
    mplsOutSegmentLdpLspEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsOutSegmentLdpLspTable/" + mplsOutSegmentLdpLspEntry.EntityData.SegmentPath
    mplsOutSegmentLdpLspEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsOutSegmentLdpLspEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsOutSegmentLdpLspEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsOutSegmentLdpLspEntry.EntityData.Children = types.NewOrderedMap()
    mplsOutSegmentLdpLspEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsOutSegmentLdpLspEntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsOutSegmentLdpLspEntry.MplsLdpEntityLdpId})
    mplsOutSegmentLdpLspEntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsOutSegmentLdpLspEntry.MplsLdpEntityIndex})
    mplsOutSegmentLdpLspEntry.EntityData.Leafs.Append("mplsLdpPeerLdpId", types.YLeaf{"MplsLdpPeerLdpId", mplsOutSegmentLdpLspEntry.MplsLdpPeerLdpId})
    mplsOutSegmentLdpLspEntry.EntityData.Leafs.Append("mplsOutSegmentLdpLspIndex", types.YLeaf{"MplsOutSegmentLdpLspIndex", mplsOutSegmentLdpLspEntry.MplsOutSegmentLdpLspIndex})
    mplsOutSegmentLdpLspEntry.EntityData.Leafs.Append("mplsOutSegmentLdpLspLabelType", types.YLeaf{"MplsOutSegmentLdpLspLabelType", mplsOutSegmentLdpLspEntry.MplsOutSegmentLdpLspLabelType})
    mplsOutSegmentLdpLspEntry.EntityData.Leafs.Append("mplsOutSegmentLdpLspType", types.YLeaf{"MplsOutSegmentLdpLspType", mplsOutSegmentLdpLspEntry.MplsOutSegmentLdpLspType})

    mplsOutSegmentLdpLspEntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex", "MplsLdpPeerLdpId", "MplsOutSegmentLdpLspIndex"}

    return &(mplsOutSegmentLdpLspEntry.EntityData)
}

// MPLSLDPSTDMIB_MplsFecTable
// This table represents the FEC
// (Forwarding Equivalence Class)
// Information associated with an LSP.
type MPLSLDPSTDMIB_MplsFecTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row represents a single FEC Element. The type is slice of
    // MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry.
    MplsFecEntry []*MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry
}

func (mplsFecTable *MPLSLDPSTDMIB_MplsFecTable) GetEntityData() *types.CommonEntityData {
    mplsFecTable.EntityData.YFilter = mplsFecTable.YFilter
    mplsFecTable.EntityData.YangName = "mplsFecTable"
    mplsFecTable.EntityData.BundleName = "cisco_ios_xe"
    mplsFecTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsFecTable.EntityData.SegmentPath = "mplsFecTable"
    mplsFecTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsFecTable.EntityData.SegmentPath
    mplsFecTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsFecTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsFecTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsFecTable.EntityData.Children = types.NewOrderedMap()
    mplsFecTable.EntityData.Children.Append("mplsFecEntry", types.YChild{"MplsFecEntry", nil})
    for i := range mplsFecTable.MplsFecEntry {
        mplsFecTable.EntityData.Children.Append(types.GetSegmentPath(mplsFecTable.MplsFecEntry[i]), types.YChild{"MplsFecEntry", mplsFecTable.MplsFecEntry[i]})
    }
    mplsFecTable.EntityData.Leafs = types.NewOrderedMap()

    mplsFecTable.EntityData.YListKeys = []string {}

    return &(mplsFecTable.EntityData)
}

// MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry
// Each row represents a single FEC Element.
type MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index which uniquely identifies this entry.
    // The type is interface{} with range: 1..4294967295.
    MplsFecIndex interface{}

    // The type of the FEC.  If the value of this object is 'prefix(1)' then the
    // FEC type described by this row is an address prefix.  If the value of this
    // object is 'hostAddress(2)' then the FEC type described by this row is a
    // host address. The type is MplsFecType.
    MplsFecType interface{}

    // If the value of the 'mplsFecType' is 'hostAddress(2)' then this object is
    // undefined.  If the value of 'mplsFecType' is 'prefix(1)' then the value of
    // this object is the length in bits of the address prefix represented by
    // 'mplsFecAddr', or zero.  If the value of this object is zero, this
    // indicates that the prefix matches all addresses.  In this case the address
    // prefix MUST also be zero (i.e., 'mplsFecAddr' should have the value of
    // zero.). The type is interface{} with range: 0..2040.
    MplsFecAddrPrefixLength interface{}

    // The value of this object is the type of the Internet address.  The value of
    // this object, decides how the value of the mplsFecAddr object is
    // interpreted. The type is InetAddressType.
    MplsFecAddrType interface{}

    // The value of this object is interpreted based on the value of the
    // 'mplsFecAddrType' object.  This address is then further interpretted as an
    // being used with the address prefix, or as the host address.  This further
    // interpretation is indicated by the 'mplsFecType' object. In other words,
    // the FEC element is populated according to the Prefix FEC Element value
    // encoding, or the Host Address FEC Element encoding. The type is string with
    // length: 0..255.
    MplsFecAddr interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    MplsFecStorageType interface{}

    // The status of this conceptual row.  If the value of this object is
    // 'active(1)', then none of the writable objects of this entry can be
    // modified, except to set this object to 'destroy(6)'.  NOTE: if this row is
    // being referenced by any entry in the mplsLdpLspFecTable, then a request to
    // destroy this row, will result in an inconsistentValue error. The type is
    // RowStatus.
    MplsFecRowStatus interface{}
}

func (mplsFecEntry *MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry) GetEntityData() *types.CommonEntityData {
    mplsFecEntry.EntityData.YFilter = mplsFecEntry.YFilter
    mplsFecEntry.EntityData.YangName = "mplsFecEntry"
    mplsFecEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsFecEntry.EntityData.ParentYangName = "mplsFecTable"
    mplsFecEntry.EntityData.SegmentPath = "mplsFecEntry" + types.AddKeyToken(mplsFecEntry.MplsFecIndex, "mplsFecIndex")
    mplsFecEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsFecTable/" + mplsFecEntry.EntityData.SegmentPath
    mplsFecEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsFecEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsFecEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsFecEntry.EntityData.Children = types.NewOrderedMap()
    mplsFecEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsFecEntry.EntityData.Leafs.Append("mplsFecIndex", types.YLeaf{"MplsFecIndex", mplsFecEntry.MplsFecIndex})
    mplsFecEntry.EntityData.Leafs.Append("mplsFecType", types.YLeaf{"MplsFecType", mplsFecEntry.MplsFecType})
    mplsFecEntry.EntityData.Leafs.Append("mplsFecAddrPrefixLength", types.YLeaf{"MplsFecAddrPrefixLength", mplsFecEntry.MplsFecAddrPrefixLength})
    mplsFecEntry.EntityData.Leafs.Append("mplsFecAddrType", types.YLeaf{"MplsFecAddrType", mplsFecEntry.MplsFecAddrType})
    mplsFecEntry.EntityData.Leafs.Append("mplsFecAddr", types.YLeaf{"MplsFecAddr", mplsFecEntry.MplsFecAddr})
    mplsFecEntry.EntityData.Leafs.Append("mplsFecStorageType", types.YLeaf{"MplsFecStorageType", mplsFecEntry.MplsFecStorageType})
    mplsFecEntry.EntityData.Leafs.Append("mplsFecRowStatus", types.YLeaf{"MplsFecRowStatus", mplsFecEntry.MplsFecRowStatus})

    mplsFecEntry.EntityData.YListKeys = []string {"MplsFecIndex"}

    return &(mplsFecEntry.EntityData)
}

// MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry_MplsFecType represents the FEC type described by this row is a host address.
type MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry_MplsFecType string

const (
    MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry_MplsFecType_prefix MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry_MplsFecType = "prefix"

    MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry_MplsFecType_hostAddress MPLSLDPSTDMIB_MplsFecTable_MplsFecEntry_MplsFecType = "hostAddress"
)

// MPLSLDPSTDMIB_MplsLdpLspFecTable
// A table which shows the relationship between
// LDP LSPs and FECs.  Each row represents
// a single LDP LSP to FEC association.
type MPLSLDPSTDMIB_MplsLdpLspFecTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry represents a LDP LSP to FEC association. The type is slice of
    // MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry.
    MplsLdpLspFecEntry []*MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry
}

func (mplsLdpLspFecTable *MPLSLDPSTDMIB_MplsLdpLspFecTable) GetEntityData() *types.CommonEntityData {
    mplsLdpLspFecTable.EntityData.YFilter = mplsLdpLspFecTable.YFilter
    mplsLdpLspFecTable.EntityData.YangName = "mplsLdpLspFecTable"
    mplsLdpLspFecTable.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpLspFecTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpLspFecTable.EntityData.SegmentPath = "mplsLdpLspFecTable"
    mplsLdpLspFecTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpLspFecTable.EntityData.SegmentPath
    mplsLdpLspFecTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpLspFecTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpLspFecTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpLspFecTable.EntityData.Children = types.NewOrderedMap()
    mplsLdpLspFecTable.EntityData.Children.Append("mplsLdpLspFecEntry", types.YChild{"MplsLdpLspFecEntry", nil})
    for i := range mplsLdpLspFecTable.MplsLdpLspFecEntry {
        mplsLdpLspFecTable.EntityData.Children.Append(types.GetSegmentPath(mplsLdpLspFecTable.MplsLdpLspFecEntry[i]), types.YChild{"MplsLdpLspFecEntry", mplsLdpLspFecTable.MplsLdpLspFecEntry[i]})
    }
    mplsLdpLspFecTable.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpLspFecTable.EntityData.YListKeys = []string {}

    return &(mplsLdpLspFecTable.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry
// An entry represents a LDP LSP
// to FEC association.
type MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityLdpId
    MplsLdpEntityLdpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityIndex
    MplsLdpEntityIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpPeerLdpId
    MplsLdpPeerLdpId interface{}

    // This attribute is a key. If the value is inSegment(1), then this indicates
    // that the following index, mplsLdpLspFecSegmentIndex, contains the same
    // value as the mplsInSegmentLdpLspIndex.  Otherwise, if the value of this
    // object is   outSegment(2),  then this indicates that following index,
    // mplsLdpLspFecSegmentIndex, contains the same value as the
    // mplsOutSegmentLdpLspIndex. The type is MplsLdpLspFecSegment.
    MplsLdpLspFecSegment interface{}

    // This attribute is a key. This index is interpretted by using the value of
    // the mplsLdpLspFecSegment.  If the mplsLdpLspFecSegment is inSegment(1),
    // then this index has the same value as mplsInSegmentLdpLspIndex.  If the
    // mplsLdpLspFecSegment is outSegment(2), then this index has the same value
    // as mplsOutSegmentLdpLspIndex. The type is string with length: 1..24.
    MplsLdpLspFecSegmentIndex interface{}

    // This attribute is a key. This index identifies the FEC entry in the
    // mplsFecTable associated with this session. In other words, the value of
    // this index is the same as the value of the mplsFecIndex that denotes the
    // FEC associated with this Session. The type is interface{} with range:
    // 1..4294967295.
    MplsLdpLspFecIndex interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    MplsLdpLspFecStorageType interface{}

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
    MplsLdpLspFecRowStatus interface{}
}

func (mplsLdpLspFecEntry *MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry) GetEntityData() *types.CommonEntityData {
    mplsLdpLspFecEntry.EntityData.YFilter = mplsLdpLspFecEntry.YFilter
    mplsLdpLspFecEntry.EntityData.YangName = "mplsLdpLspFecEntry"
    mplsLdpLspFecEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpLspFecEntry.EntityData.ParentYangName = "mplsLdpLspFecTable"
    mplsLdpLspFecEntry.EntityData.SegmentPath = "mplsLdpLspFecEntry" + types.AddKeyToken(mplsLdpLspFecEntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsLdpLspFecEntry.MplsLdpEntityIndex, "mplsLdpEntityIndex") + types.AddKeyToken(mplsLdpLspFecEntry.MplsLdpPeerLdpId, "mplsLdpPeerLdpId") + types.AddKeyToken(mplsLdpLspFecEntry.MplsLdpLspFecSegment, "mplsLdpLspFecSegment") + types.AddKeyToken(mplsLdpLspFecEntry.MplsLdpLspFecSegmentIndex, "mplsLdpLspFecSegmentIndex") + types.AddKeyToken(mplsLdpLspFecEntry.MplsLdpLspFecIndex, "mplsLdpLspFecIndex")
    mplsLdpLspFecEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsLdpLspFecTable/" + mplsLdpLspFecEntry.EntityData.SegmentPath
    mplsLdpLspFecEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpLspFecEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpLspFecEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpLspFecEntry.EntityData.Children = types.NewOrderedMap()
    mplsLdpLspFecEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsLdpLspFecEntry.MplsLdpEntityLdpId})
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsLdpLspFecEntry.MplsLdpEntityIndex})
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpPeerLdpId", types.YLeaf{"MplsLdpPeerLdpId", mplsLdpLspFecEntry.MplsLdpPeerLdpId})
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpLspFecSegment", types.YLeaf{"MplsLdpLspFecSegment", mplsLdpLspFecEntry.MplsLdpLspFecSegment})
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpLspFecSegmentIndex", types.YLeaf{"MplsLdpLspFecSegmentIndex", mplsLdpLspFecEntry.MplsLdpLspFecSegmentIndex})
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpLspFecIndex", types.YLeaf{"MplsLdpLspFecIndex", mplsLdpLspFecEntry.MplsLdpLspFecIndex})
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpLspFecStorageType", types.YLeaf{"MplsLdpLspFecStorageType", mplsLdpLspFecEntry.MplsLdpLspFecStorageType})
    mplsLdpLspFecEntry.EntityData.Leafs.Append("mplsLdpLspFecRowStatus", types.YLeaf{"MplsLdpLspFecRowStatus", mplsLdpLspFecEntry.MplsLdpLspFecRowStatus})

    mplsLdpLspFecEntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex", "MplsLdpPeerLdpId", "MplsLdpLspFecSegment", "MplsLdpLspFecSegmentIndex", "MplsLdpLspFecIndex"}

    return &(mplsLdpLspFecEntry.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry_MplsLdpLspFecSegment represents value as the mplsOutSegmentLdpLspIndex.
type MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry_MplsLdpLspFecSegment string

const (
    MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry_MplsLdpLspFecSegment_inSegment MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry_MplsLdpLspFecSegment = "inSegment"

    MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry_MplsLdpLspFecSegment_outSegment MPLSLDPSTDMIB_MplsLdpLspFecTable_MplsLdpLspFecEntry_MplsLdpLspFecSegment = "outSegment"
)

// MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable
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
type MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents information on a session's single next
    // hop address which was advertised in an Address Message from the LDP peer.
    // The information contained in a row is read-only. The type is slice of
    // MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable_MplsLdpSessionPeerAddrEntry.
    MplsLdpSessionPeerAddrEntry []*MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable_MplsLdpSessionPeerAddrEntry
}

func (mplsLdpSessionPeerAddrTable *MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable) GetEntityData() *types.CommonEntityData {
    mplsLdpSessionPeerAddrTable.EntityData.YFilter = mplsLdpSessionPeerAddrTable.YFilter
    mplsLdpSessionPeerAddrTable.EntityData.YangName = "mplsLdpSessionPeerAddrTable"
    mplsLdpSessionPeerAddrTable.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpSessionPeerAddrTable.EntityData.ParentYangName = "MPLS-LDP-STD-MIB"
    mplsLdpSessionPeerAddrTable.EntityData.SegmentPath = "mplsLdpSessionPeerAddrTable"
    mplsLdpSessionPeerAddrTable.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/" + mplsLdpSessionPeerAddrTable.EntityData.SegmentPath
    mplsLdpSessionPeerAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpSessionPeerAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpSessionPeerAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpSessionPeerAddrTable.EntityData.Children = types.NewOrderedMap()
    mplsLdpSessionPeerAddrTable.EntityData.Children.Append("mplsLdpSessionPeerAddrEntry", types.YChild{"MplsLdpSessionPeerAddrEntry", nil})
    for i := range mplsLdpSessionPeerAddrTable.MplsLdpSessionPeerAddrEntry {
        mplsLdpSessionPeerAddrTable.EntityData.Children.Append(types.GetSegmentPath(mplsLdpSessionPeerAddrTable.MplsLdpSessionPeerAddrEntry[i]), types.YChild{"MplsLdpSessionPeerAddrEntry", mplsLdpSessionPeerAddrTable.MplsLdpSessionPeerAddrEntry[i]})
    }
    mplsLdpSessionPeerAddrTable.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpSessionPeerAddrTable.EntityData.YListKeys = []string {}

    return &(mplsLdpSessionPeerAddrTable.EntityData)
}

// MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable_MplsLdpSessionPeerAddrEntry
// An entry in this table represents information on
// a session's single next hop address which was
// advertised in an Address Message from the LDP peer.
// The information contained in a row is read-only.
type MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable_MplsLdpSessionPeerAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityLdpId
    MplsLdpEntityLdpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityIndex
    MplsLdpEntityIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpPeerTable_MplsLdpPeerEntry_MplsLdpPeerLdpId
    MplsLdpPeerLdpId interface{}

    // This attribute is a key. An index which uniquely identifies this entry
    // within a given session. The type is interface{} with range: 1..4294967295.
    MplsLdpSessionPeerAddrIndex interface{}

    // The internetwork layer address type of this Next Hop Address as specified
    // in the Label Address Message associated with this Session. The value of
    // this object indicates how to interpret the value of  
    // mplsLdpSessionPeerNextHopAddr. The type is InetAddressType.
    MplsLdpSessionPeerNextHopAddrType interface{}

    // The next hop address.  The type of this address is specified by the value
    // of the mplsLdpSessionPeerNextHopAddrType. The type is string with length:
    // 0..255.
    MplsLdpSessionPeerNextHopAddr interface{}
}

func (mplsLdpSessionPeerAddrEntry *MPLSLDPSTDMIB_MplsLdpSessionPeerAddrTable_MplsLdpSessionPeerAddrEntry) GetEntityData() *types.CommonEntityData {
    mplsLdpSessionPeerAddrEntry.EntityData.YFilter = mplsLdpSessionPeerAddrEntry.YFilter
    mplsLdpSessionPeerAddrEntry.EntityData.YangName = "mplsLdpSessionPeerAddrEntry"
    mplsLdpSessionPeerAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpSessionPeerAddrEntry.EntityData.ParentYangName = "mplsLdpSessionPeerAddrTable"
    mplsLdpSessionPeerAddrEntry.EntityData.SegmentPath = "mplsLdpSessionPeerAddrEntry" + types.AddKeyToken(mplsLdpSessionPeerAddrEntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsLdpSessionPeerAddrEntry.MplsLdpEntityIndex, "mplsLdpEntityIndex") + types.AddKeyToken(mplsLdpSessionPeerAddrEntry.MplsLdpPeerLdpId, "mplsLdpPeerLdpId") + types.AddKeyToken(mplsLdpSessionPeerAddrEntry.MplsLdpSessionPeerAddrIndex, "mplsLdpSessionPeerAddrIndex")
    mplsLdpSessionPeerAddrEntry.EntityData.AbsolutePath = "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB/mplsLdpSessionPeerAddrTable/" + mplsLdpSessionPeerAddrEntry.EntityData.SegmentPath
    mplsLdpSessionPeerAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpSessionPeerAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpSessionPeerAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpSessionPeerAddrEntry.EntityData.Children = types.NewOrderedMap()
    mplsLdpSessionPeerAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpSessionPeerAddrEntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsLdpSessionPeerAddrEntry.MplsLdpEntityLdpId})
    mplsLdpSessionPeerAddrEntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsLdpSessionPeerAddrEntry.MplsLdpEntityIndex})
    mplsLdpSessionPeerAddrEntry.EntityData.Leafs.Append("mplsLdpPeerLdpId", types.YLeaf{"MplsLdpPeerLdpId", mplsLdpSessionPeerAddrEntry.MplsLdpPeerLdpId})
    mplsLdpSessionPeerAddrEntry.EntityData.Leafs.Append("mplsLdpSessionPeerAddrIndex", types.YLeaf{"MplsLdpSessionPeerAddrIndex", mplsLdpSessionPeerAddrEntry.MplsLdpSessionPeerAddrIndex})
    mplsLdpSessionPeerAddrEntry.EntityData.Leafs.Append("mplsLdpSessionPeerNextHopAddrType", types.YLeaf{"MplsLdpSessionPeerNextHopAddrType", mplsLdpSessionPeerAddrEntry.MplsLdpSessionPeerNextHopAddrType})
    mplsLdpSessionPeerAddrEntry.EntityData.Leafs.Append("mplsLdpSessionPeerNextHopAddr", types.YLeaf{"MplsLdpSessionPeerNextHopAddr", mplsLdpSessionPeerAddrEntry.MplsLdpSessionPeerNextHopAddr})

    mplsLdpSessionPeerAddrEntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex", "MplsLdpPeerLdpId", "MplsLdpSessionPeerAddrIndex"}

    return &(mplsLdpSessionPeerAddrEntry.EntityData)
}

