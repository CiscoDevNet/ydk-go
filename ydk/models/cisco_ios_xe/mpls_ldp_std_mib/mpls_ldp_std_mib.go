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
    parent types.Entity
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

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetFilter() yfilter.YFilter { return mPLSLDPSTDMIB.YFilter }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) SetFilter(yf yfilter.YFilter) { mPLSLDPSTDMIB.YFilter = yf }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetGoName(yname string) string {
    if yname == "mplsLdpLsrObjects" { return "Mplsldplsrobjects" }
    if yname == "mplsLdpEntityObjects" { return "Mplsldpentityobjects" }
    if yname == "mplsLdpSessionObjects" { return "Mplsldpsessionobjects" }
    if yname == "mplsFecObjects" { return "Mplsfecobjects" }
    if yname == "mplsLdpEntityTable" { return "Mplsldpentitytable" }
    if yname == "mplsLdpPeerTable" { return "Mplsldppeertable" }
    if yname == "mplsLdpHelloAdjacencyTable" { return "Mplsldphelloadjacencytable" }
    if yname == "mplsInSegmentLdpLspTable" { return "Mplsinsegmentldplsptable" }
    if yname == "mplsOutSegmentLdpLspTable" { return "Mplsoutsegmentldplsptable" }
    if yname == "mplsFecTable" { return "Mplsfectable" }
    if yname == "mplsLdpLspFecTable" { return "Mplsldplspfectable" }
    if yname == "mplsLdpSessionPeerAddrTable" { return "Mplsldpsessionpeeraddrtable" }
    return ""
}

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetSegmentPath() string {
    return "MPLS-LDP-STD-MIB:MPLS-LDP-STD-MIB"
}

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpLsrObjects" {
        return &mPLSLDPSTDMIB.Mplsldplsrobjects
    }
    if childYangName == "mplsLdpEntityObjects" {
        return &mPLSLDPSTDMIB.Mplsldpentityobjects
    }
    if childYangName == "mplsLdpSessionObjects" {
        return &mPLSLDPSTDMIB.Mplsldpsessionobjects
    }
    if childYangName == "mplsFecObjects" {
        return &mPLSLDPSTDMIB.Mplsfecobjects
    }
    if childYangName == "mplsLdpEntityTable" {
        return &mPLSLDPSTDMIB.Mplsldpentitytable
    }
    if childYangName == "mplsLdpPeerTable" {
        return &mPLSLDPSTDMIB.Mplsldppeertable
    }
    if childYangName == "mplsLdpHelloAdjacencyTable" {
        return &mPLSLDPSTDMIB.Mplsldphelloadjacencytable
    }
    if childYangName == "mplsInSegmentLdpLspTable" {
        return &mPLSLDPSTDMIB.Mplsinsegmentldplsptable
    }
    if childYangName == "mplsOutSegmentLdpLspTable" {
        return &mPLSLDPSTDMIB.Mplsoutsegmentldplsptable
    }
    if childYangName == "mplsFecTable" {
        return &mPLSLDPSTDMIB.Mplsfectable
    }
    if childYangName == "mplsLdpLspFecTable" {
        return &mPLSLDPSTDMIB.Mplsldplspfectable
    }
    if childYangName == "mplsLdpSessionPeerAddrTable" {
        return &mPLSLDPSTDMIB.Mplsldpsessionpeeraddrtable
    }
    return nil
}

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mplsLdpLsrObjects"] = &mPLSLDPSTDMIB.Mplsldplsrobjects
    children["mplsLdpEntityObjects"] = &mPLSLDPSTDMIB.Mplsldpentityobjects
    children["mplsLdpSessionObjects"] = &mPLSLDPSTDMIB.Mplsldpsessionobjects
    children["mplsFecObjects"] = &mPLSLDPSTDMIB.Mplsfecobjects
    children["mplsLdpEntityTable"] = &mPLSLDPSTDMIB.Mplsldpentitytable
    children["mplsLdpPeerTable"] = &mPLSLDPSTDMIB.Mplsldppeertable
    children["mplsLdpHelloAdjacencyTable"] = &mPLSLDPSTDMIB.Mplsldphelloadjacencytable
    children["mplsInSegmentLdpLspTable"] = &mPLSLDPSTDMIB.Mplsinsegmentldplsptable
    children["mplsOutSegmentLdpLspTable"] = &mPLSLDPSTDMIB.Mplsoutsegmentldplsptable
    children["mplsFecTable"] = &mPLSLDPSTDMIB.Mplsfectable
    children["mplsLdpLspFecTable"] = &mPLSLDPSTDMIB.Mplsldplspfectable
    children["mplsLdpSessionPeerAddrTable"] = &mPLSLDPSTDMIB.Mplsldpsessionpeeraddrtable
    return children
}

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetYangName() string { return "MPLS-LDP-STD-MIB" }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) SetParent(parent types.Entity) { mPLSLDPSTDMIB.parent = parent }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetParent() types.Entity { return mPLSLDPSTDMIB.parent }

func (mPLSLDPSTDMIB *MPLSLDPSTDMIB) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldplsrobjects
type MPLSLDPSTDMIB_Mplsldplsrobjects struct {
    parent types.Entity
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

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetFilter() yfilter.YFilter { return mplsldplsrobjects.YFilter }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) SetFilter(yf yfilter.YFilter) { mplsldplsrobjects.YFilter = yf }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetGoName(yname string) string {
    if yname == "mplsLdpLsrId" { return "Mplsldplsrid" }
    if yname == "mplsLdpLsrLoopDetectionCapable" { return "Mplsldplsrloopdetectioncapable" }
    return ""
}

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetSegmentPath() string {
    return "mplsLdpLsrObjects"
}

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpLsrId"] = mplsldplsrobjects.Mplsldplsrid
    leafs["mplsLdpLsrLoopDetectionCapable"] = mplsldplsrobjects.Mplsldplsrloopdetectioncapable
    return leafs
}

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetYangName() string { return "mplsLdpLsrObjects" }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) SetParent(parent types.Entity) { mplsldplsrobjects.parent = parent }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetParent() types.Entity { return mplsldplsrobjects.parent }

func (mplsldplsrobjects *MPLSLDPSTDMIB_Mplsldplsrobjects) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

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
    parent types.Entity
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

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetFilter() yfilter.YFilter { return mplsldpentityobjects.YFilter }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) SetFilter(yf yfilter.YFilter) { mplsldpentityobjects.YFilter = yf }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLastChange" { return "Mplsldpentitylastchange" }
    if yname == "mplsLdpEntityIndexNext" { return "Mplsldpentityindexnext" }
    return ""
}

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetSegmentPath() string {
    return "mplsLdpEntityObjects"
}

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLastChange"] = mplsldpentityobjects.Mplsldpentitylastchange
    leafs["mplsLdpEntityIndexNext"] = mplsldpentityobjects.Mplsldpentityindexnext
    return leafs
}

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetYangName() string { return "mplsLdpEntityObjects" }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) SetParent(parent types.Entity) { mplsldpentityobjects.parent = parent }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetParent() types.Entity { return mplsldpentityobjects.parent }

func (mplsldpentityobjects *MPLSLDPSTDMIB_Mplsldpentityobjects) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldpsessionobjects
type MPLSLDPSTDMIB_Mplsldpsessionobjects struct {
    parent types.Entity
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

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetFilter() yfilter.YFilter { return mplsldpsessionobjects.YFilter }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) SetFilter(yf yfilter.YFilter) { mplsldpsessionobjects.YFilter = yf }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetGoName(yname string) string {
    if yname == "mplsLdpPeerLastChange" { return "Mplsldppeerlastchange" }
    if yname == "mplsLdpLspFecLastChange" { return "Mplsldplspfeclastchange" }
    return ""
}

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetSegmentPath() string {
    return "mplsLdpSessionObjects"
}

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpPeerLastChange"] = mplsldpsessionobjects.Mplsldppeerlastchange
    leafs["mplsLdpLspFecLastChange"] = mplsldpsessionobjects.Mplsldplspfeclastchange
    return leafs
}

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetYangName() string { return "mplsLdpSessionObjects" }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) SetParent(parent types.Entity) { mplsldpsessionobjects.parent = parent }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetParent() types.Entity { return mplsldpsessionobjects.parent }

func (mplsldpsessionobjects *MPLSLDPSTDMIB_Mplsldpsessionobjects) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsfecobjects
type MPLSLDPSTDMIB_Mplsfecobjects struct {
    parent types.Entity
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

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetFilter() yfilter.YFilter { return mplsfecobjects.YFilter }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) SetFilter(yf yfilter.YFilter) { mplsfecobjects.YFilter = yf }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetGoName(yname string) string {
    if yname == "mplsFecLastChange" { return "Mplsfeclastchange" }
    if yname == "mplsFecIndexNext" { return "Mplsfecindexnext" }
    return ""
}

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetSegmentPath() string {
    return "mplsFecObjects"
}

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsFecLastChange"] = mplsfecobjects.Mplsfeclastchange
    leafs["mplsFecIndexNext"] = mplsfecobjects.Mplsfecindexnext
    return leafs
}

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetBundleName() string { return "cisco_ios_xe" }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetYangName() string { return "mplsFecObjects" }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) SetParent(parent types.Entity) { mplsfecobjects.parent = parent }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetParent() types.Entity { return mplsfecobjects.parent }

func (mplsfecobjects *MPLSLDPSTDMIB_Mplsfecobjects) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldpentitytable
// This table contains information about the
// MPLS Label Distribution Protocol Entities which
// exist on this Label Switching Router (LSR)
// or Label Edge Router (LER).
type MPLSLDPSTDMIB_Mplsldpentitytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents an LDP entity. An entry can be created by
    // a network administrator or by an SNMP agent as instructed by LDP. The type
    // is slice of MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry.
    Mplsldpentityentry []MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry
}

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetFilter() yfilter.YFilter { return mplsldpentitytable.YFilter }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) SetFilter(yf yfilter.YFilter) { mplsldpentitytable.YFilter = yf }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetGoName(yname string) string {
    if yname == "mplsLdpEntityEntry" { return "Mplsldpentityentry" }
    return ""
}

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetSegmentPath() string {
    return "mplsLdpEntityTable"
}

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpEntityEntry" {
        for _, c := range mplsldpentitytable.Mplsldpentityentry {
            if mplsldpentitytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry{}
        mplsldpentitytable.Mplsldpentityentry = append(mplsldpentitytable.Mplsldpentityentry, child)
        return &mplsldpentitytable.Mplsldpentityentry[len(mplsldpentitytable.Mplsldpentityentry)-1]
    }
    return nil
}

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsldpentitytable.Mplsldpentityentry {
        children[mplsldpentitytable.Mplsldpentityentry[i].GetSegmentPath()] = &mplsldpentitytable.Mplsldpentityentry[i]
    }
    return children
}

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetYangName() string { return "mplsLdpEntityTable" }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) SetParent(parent types.Entity) { mplsldpentitytable.parent = parent }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetParent() types.Entity { return mplsldpentitytable.parent }

func (mplsldpentitytable *MPLSLDPSTDMIB_Mplsldpentitytable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry
// An entry in this table represents an LDP entity.
// An entry can be created by a network administrator
// or by an SNMP agent as instructed by LDP.
type MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry struct {
    parent types.Entity
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

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetFilter() yfilter.YFilter { return mplsldpentityentry.YFilter }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) SetFilter(yf yfilter.YFilter) { mplsldpentityentry.YFilter = yf }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpEntityProtocolVersion" { return "Mplsldpentityprotocolversion" }
    if yname == "mplsLdpEntityAdminStatus" { return "Mplsldpentityadminstatus" }
    if yname == "mplsLdpEntityOperStatus" { return "Mplsldpentityoperstatus" }
    if yname == "mplsLdpEntityTcpPort" { return "Mplsldpentitytcpport" }
    if yname == "mplsLdpEntityUdpDscPort" { return "Mplsldpentityudpdscport" }
    if yname == "mplsLdpEntityMaxPduLength" { return "Mplsldpentitymaxpdulength" }
    if yname == "mplsLdpEntityKeepAliveHoldTimer" { return "Mplsldpentitykeepaliveholdtimer" }
    if yname == "mplsLdpEntityHelloHoldTimer" { return "Mplsldpentityhelloholdtimer" }
    if yname == "mplsLdpEntityInitSessionThreshold" { return "Mplsldpentityinitsessionthreshold" }
    if yname == "mplsLdpEntityLabelDistMethod" { return "Mplsldpentitylabeldistmethod" }
    if yname == "mplsLdpEntityLabelRetentionMode" { return "Mplsldpentitylabelretentionmode" }
    if yname == "mplsLdpEntityPathVectorLimit" { return "Mplsldpentitypathvectorlimit" }
    if yname == "mplsLdpEntityHopCountLimit" { return "Mplsldpentityhopcountlimit" }
    if yname == "mplsLdpEntityTransportAddrKind" { return "Mplsldpentitytransportaddrkind" }
    if yname == "mplsLdpEntityTargetPeer" { return "Mplsldpentitytargetpeer" }
    if yname == "mplsLdpEntityTargetPeerAddrType" { return "Mplsldpentitytargetpeeraddrtype" }
    if yname == "mplsLdpEntityTargetPeerAddr" { return "Mplsldpentitytargetpeeraddr" }
    if yname == "mplsLdpEntityLabelType" { return "Mplsldpentitylabeltype" }
    if yname == "mplsLdpEntityDiscontinuityTime" { return "Mplsldpentitydiscontinuitytime" }
    if yname == "mplsLdpEntityStorageType" { return "Mplsldpentitystoragetype" }
    if yname == "mplsLdpEntityRowStatus" { return "Mplsldpentityrowstatus" }
    if yname == "mplsLdpEntityStatsSessionAttempts" { return "Mplsldpentitystatssessionattempts" }
    if yname == "mplsLdpEntityStatsSessionRejectedNoHelloErrors" { return "Mplsldpentitystatssessionrejectednohelloerrors" }
    if yname == "mplsLdpEntityStatsSessionRejectedAdErrors" { return "Mplsldpentitystatssessionrejectedaderrors" }
    if yname == "mplsLdpEntityStatsSessionRejectedMaxPduErrors" { return "Mplsldpentitystatssessionrejectedmaxpduerrors" }
    if yname == "mplsLdpEntityStatsSessionRejectedLRErrors" { return "Mplsldpentitystatssessionrejectedlrerrors" }
    if yname == "mplsLdpEntityStatsBadLdpIdentifierErrors" { return "Mplsldpentitystatsbadldpidentifiererrors" }
    if yname == "mplsLdpEntityStatsBadPduLengthErrors" { return "Mplsldpentitystatsbadpdulengtherrors" }
    if yname == "mplsLdpEntityStatsBadMessageLengthErrors" { return "Mplsldpentitystatsbadmessagelengtherrors" }
    if yname == "mplsLdpEntityStatsBadTlvLengthErrors" { return "Mplsldpentitystatsbadtlvlengtherrors" }
    if yname == "mplsLdpEntityStatsMalformedTlvValueErrors" { return "Mplsldpentitystatsmalformedtlvvalueerrors" }
    if yname == "mplsLdpEntityStatsKeepAliveTimerExpErrors" { return "Mplsldpentitystatskeepalivetimerexperrors" }
    if yname == "mplsLdpEntityStatsShutdownReceivedNotifications" { return "Mplsldpentitystatsshutdownreceivednotifications" }
    if yname == "mplsLdpEntityStatsShutdownSentNotifications" { return "Mplsldpentitystatsshutdownsentnotifications" }
    return ""
}

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetSegmentPath() string {
    return "mplsLdpEntityEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldpentityentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldpentityentry.Mplsldpentityindex) + "']"
}

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsldpentityentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsldpentityentry.Mplsldpentityindex
    leafs["mplsLdpEntityProtocolVersion"] = mplsldpentityentry.Mplsldpentityprotocolversion
    leafs["mplsLdpEntityAdminStatus"] = mplsldpentityentry.Mplsldpentityadminstatus
    leafs["mplsLdpEntityOperStatus"] = mplsldpentityentry.Mplsldpentityoperstatus
    leafs["mplsLdpEntityTcpPort"] = mplsldpentityentry.Mplsldpentitytcpport
    leafs["mplsLdpEntityUdpDscPort"] = mplsldpentityentry.Mplsldpentityudpdscport
    leafs["mplsLdpEntityMaxPduLength"] = mplsldpentityentry.Mplsldpentitymaxpdulength
    leafs["mplsLdpEntityKeepAliveHoldTimer"] = mplsldpentityentry.Mplsldpentitykeepaliveholdtimer
    leafs["mplsLdpEntityHelloHoldTimer"] = mplsldpentityentry.Mplsldpentityhelloholdtimer
    leafs["mplsLdpEntityInitSessionThreshold"] = mplsldpentityentry.Mplsldpentityinitsessionthreshold
    leafs["mplsLdpEntityLabelDistMethod"] = mplsldpentityentry.Mplsldpentitylabeldistmethod
    leafs["mplsLdpEntityLabelRetentionMode"] = mplsldpentityentry.Mplsldpentitylabelretentionmode
    leafs["mplsLdpEntityPathVectorLimit"] = mplsldpentityentry.Mplsldpentitypathvectorlimit
    leafs["mplsLdpEntityHopCountLimit"] = mplsldpentityentry.Mplsldpentityhopcountlimit
    leafs["mplsLdpEntityTransportAddrKind"] = mplsldpentityentry.Mplsldpentitytransportaddrkind
    leafs["mplsLdpEntityTargetPeer"] = mplsldpentityentry.Mplsldpentitytargetpeer
    leafs["mplsLdpEntityTargetPeerAddrType"] = mplsldpentityentry.Mplsldpentitytargetpeeraddrtype
    leafs["mplsLdpEntityTargetPeerAddr"] = mplsldpentityentry.Mplsldpentitytargetpeeraddr
    leafs["mplsLdpEntityLabelType"] = mplsldpentityentry.Mplsldpentitylabeltype
    leafs["mplsLdpEntityDiscontinuityTime"] = mplsldpentityentry.Mplsldpentitydiscontinuitytime
    leafs["mplsLdpEntityStorageType"] = mplsldpentityentry.Mplsldpentitystoragetype
    leafs["mplsLdpEntityRowStatus"] = mplsldpentityentry.Mplsldpentityrowstatus
    leafs["mplsLdpEntityStatsSessionAttempts"] = mplsldpentityentry.Mplsldpentitystatssessionattempts
    leafs["mplsLdpEntityStatsSessionRejectedNoHelloErrors"] = mplsldpentityentry.Mplsldpentitystatssessionrejectednohelloerrors
    leafs["mplsLdpEntityStatsSessionRejectedAdErrors"] = mplsldpentityentry.Mplsldpentitystatssessionrejectedaderrors
    leafs["mplsLdpEntityStatsSessionRejectedMaxPduErrors"] = mplsldpentityentry.Mplsldpentitystatssessionrejectedmaxpduerrors
    leafs["mplsLdpEntityStatsSessionRejectedLRErrors"] = mplsldpentityentry.Mplsldpentitystatssessionrejectedlrerrors
    leafs["mplsLdpEntityStatsBadLdpIdentifierErrors"] = mplsldpentityentry.Mplsldpentitystatsbadldpidentifiererrors
    leafs["mplsLdpEntityStatsBadPduLengthErrors"] = mplsldpentityentry.Mplsldpentitystatsbadpdulengtherrors
    leafs["mplsLdpEntityStatsBadMessageLengthErrors"] = mplsldpentityentry.Mplsldpentitystatsbadmessagelengtherrors
    leafs["mplsLdpEntityStatsBadTlvLengthErrors"] = mplsldpentityentry.Mplsldpentitystatsbadtlvlengtherrors
    leafs["mplsLdpEntityStatsMalformedTlvValueErrors"] = mplsldpentityentry.Mplsldpentitystatsmalformedtlvvalueerrors
    leafs["mplsLdpEntityStatsKeepAliveTimerExpErrors"] = mplsldpentityentry.Mplsldpentitystatskeepalivetimerexperrors
    leafs["mplsLdpEntityStatsShutdownReceivedNotifications"] = mplsldpentityentry.Mplsldpentitystatsshutdownreceivednotifications
    leafs["mplsLdpEntityStatsShutdownSentNotifications"] = mplsldpentityentry.Mplsldpentitystatsshutdownsentnotifications
    return leafs
}

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetYangName() string { return "mplsLdpEntityEntry" }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) SetParent(parent types.Entity) { mplsldpentityentry.parent = parent }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetParent() types.Entity { return mplsldpentityentry.parent }

func (mplsldpentityentry *MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry) GetParentYangName() string { return "mplsLdpEntityTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single Peer which is related to a Session.  This table
    // is augmented by the mplsLdpSessionTable. The type is slice of
    // MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry.
    Mplsldppeerentry []MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry
}

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetFilter() yfilter.YFilter { return mplsldppeertable.YFilter }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) SetFilter(yf yfilter.YFilter) { mplsldppeertable.YFilter = yf }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetGoName(yname string) string {
    if yname == "mplsLdpPeerEntry" { return "Mplsldppeerentry" }
    return ""
}

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetSegmentPath() string {
    return "mplsLdpPeerTable"
}

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpPeerEntry" {
        for _, c := range mplsldppeertable.Mplsldppeerentry {
            if mplsldppeertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry{}
        mplsldppeertable.Mplsldppeerentry = append(mplsldppeertable.Mplsldppeerentry, child)
        return &mplsldppeertable.Mplsldppeerentry[len(mplsldppeertable.Mplsldppeerentry)-1]
    }
    return nil
}

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsldppeertable.Mplsldppeerentry {
        children[mplsldppeertable.Mplsldppeerentry[i].GetSegmentPath()] = &mplsldppeertable.Mplsldppeerentry[i]
    }
    return children
}

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetYangName() string { return "mplsLdpPeerTable" }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) SetParent(parent types.Entity) { mplsldppeertable.parent = parent }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetParent() types.Entity { return mplsldppeertable.parent }

func (mplsldppeertable *MPLSLDPSTDMIB_Mplsldppeertable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry
// Information about a single Peer which is related
// to a Session.  This table is augmented by
// the mplsLdpSessionTable.
type MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry struct {
    parent types.Entity
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

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetFilter() yfilter.YFilter { return mplsldppeerentry.YFilter }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) SetFilter(yf yfilter.YFilter) { mplsldppeerentry.YFilter = yf }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpPeerLdpId" { return "Mplsldppeerldpid" }
    if yname == "mplsLdpPeerLabelDistMethod" { return "Mplsldppeerlabeldistmethod" }
    if yname == "mplsLdpPeerPathVectorLimit" { return "Mplsldppeerpathvectorlimit" }
    if yname == "mplsLdpPeerTransportAddrType" { return "Mplsldppeertransportaddrtype" }
    if yname == "mplsLdpPeerTransportAddr" { return "Mplsldppeertransportaddr" }
    if yname == "mplsLdpSessionStateLastChange" { return "Mplsldpsessionstatelastchange" }
    if yname == "mplsLdpSessionState" { return "Mplsldpsessionstate" }
    if yname == "mplsLdpSessionRole" { return "Mplsldpsessionrole" }
    if yname == "mplsLdpSessionProtocolVersion" { return "Mplsldpsessionprotocolversion" }
    if yname == "mplsLdpSessionKeepAliveHoldTimeRem" { return "Mplsldpsessionkeepaliveholdtimerem" }
    if yname == "mplsLdpSessionKeepAliveTime" { return "Mplsldpsessionkeepalivetime" }
    if yname == "mplsLdpSessionMaxPduLength" { return "Mplsldpsessionmaxpdulength" }
    if yname == "mplsLdpSessionDiscontinuityTime" { return "Mplsldpsessiondiscontinuitytime" }
    if yname == "mplsLdpSessionStatsUnknownMesTypeErrors" { return "Mplsldpsessionstatsunknownmestypeerrors" }
    if yname == "mplsLdpSessionStatsUnknownTlvErrors" { return "Mplsldpsessionstatsunknowntlverrors" }
    return ""
}

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetSegmentPath() string {
    return "mplsLdpPeerEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldppeerentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldppeerentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldppeerentry.Mplsldppeerldpid) + "']"
}

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsldppeerentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsldppeerentry.Mplsldpentityindex
    leafs["mplsLdpPeerLdpId"] = mplsldppeerentry.Mplsldppeerldpid
    leafs["mplsLdpPeerLabelDistMethod"] = mplsldppeerentry.Mplsldppeerlabeldistmethod
    leafs["mplsLdpPeerPathVectorLimit"] = mplsldppeerentry.Mplsldppeerpathvectorlimit
    leafs["mplsLdpPeerTransportAddrType"] = mplsldppeerentry.Mplsldppeertransportaddrtype
    leafs["mplsLdpPeerTransportAddr"] = mplsldppeerentry.Mplsldppeertransportaddr
    leafs["mplsLdpSessionStateLastChange"] = mplsldppeerentry.Mplsldpsessionstatelastchange
    leafs["mplsLdpSessionState"] = mplsldppeerentry.Mplsldpsessionstate
    leafs["mplsLdpSessionRole"] = mplsldppeerentry.Mplsldpsessionrole
    leafs["mplsLdpSessionProtocolVersion"] = mplsldppeerentry.Mplsldpsessionprotocolversion
    leafs["mplsLdpSessionKeepAliveHoldTimeRem"] = mplsldppeerentry.Mplsldpsessionkeepaliveholdtimerem
    leafs["mplsLdpSessionKeepAliveTime"] = mplsldppeerentry.Mplsldpsessionkeepalivetime
    leafs["mplsLdpSessionMaxPduLength"] = mplsldppeerentry.Mplsldpsessionmaxpdulength
    leafs["mplsLdpSessionDiscontinuityTime"] = mplsldppeerentry.Mplsldpsessiondiscontinuitytime
    leafs["mplsLdpSessionStatsUnknownMesTypeErrors"] = mplsldppeerentry.Mplsldpsessionstatsunknownmestypeerrors
    leafs["mplsLdpSessionStatsUnknownTlvErrors"] = mplsldppeerentry.Mplsldpsessionstatsunknowntlverrors
    return leafs
}

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetYangName() string { return "mplsLdpPeerEntry" }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) SetParent(parent types.Entity) { mplsldppeerentry.parent = parent }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetParent() types.Entity { return mplsldppeerentry.parent }

func (mplsldppeerentry *MPLSLDPSTDMIB_Mplsldppeertable_Mplsldppeerentry) GetParentYangName() string { return "mplsLdpPeerTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each row represents a single LDP Hello Adjacency. An LDP Session can have
    // one or more Hello Adjacencies. The type is slice of
    // MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry.
    Mplsldphelloadjacencyentry []MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry
}

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetFilter() yfilter.YFilter { return mplsldphelloadjacencytable.YFilter }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) SetFilter(yf yfilter.YFilter) { mplsldphelloadjacencytable.YFilter = yf }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetGoName(yname string) string {
    if yname == "mplsLdpHelloAdjacencyEntry" { return "Mplsldphelloadjacencyentry" }
    return ""
}

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetSegmentPath() string {
    return "mplsLdpHelloAdjacencyTable"
}

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpHelloAdjacencyEntry" {
        for _, c := range mplsldphelloadjacencytable.Mplsldphelloadjacencyentry {
            if mplsldphelloadjacencytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry{}
        mplsldphelloadjacencytable.Mplsldphelloadjacencyentry = append(mplsldphelloadjacencytable.Mplsldphelloadjacencyentry, child)
        return &mplsldphelloadjacencytable.Mplsldphelloadjacencyentry[len(mplsldphelloadjacencytable.Mplsldphelloadjacencyentry)-1]
    }
    return nil
}

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsldphelloadjacencytable.Mplsldphelloadjacencyentry {
        children[mplsldphelloadjacencytable.Mplsldphelloadjacencyentry[i].GetSegmentPath()] = &mplsldphelloadjacencytable.Mplsldphelloadjacencyentry[i]
    }
    return children
}

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetYangName() string { return "mplsLdpHelloAdjacencyTable" }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) SetParent(parent types.Entity) { mplsldphelloadjacencytable.parent = parent }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetParent() types.Entity { return mplsldphelloadjacencytable.parent }

func (mplsldphelloadjacencytable *MPLSLDPSTDMIB_Mplsldphelloadjacencytable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry
// Each row represents a single LDP Hello Adjacency.
// An LDP Session can have one or more Hello
// Adjacencies.
type MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry struct {
    parent types.Entity
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

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetFilter() yfilter.YFilter { return mplsldphelloadjacencyentry.YFilter }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) SetFilter(yf yfilter.YFilter) { mplsldphelloadjacencyentry.YFilter = yf }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpPeerLdpId" { return "Mplsldppeerldpid" }
    if yname == "mplsLdpHelloAdjacencyIndex" { return "Mplsldphelloadjacencyindex" }
    if yname == "mplsLdpHelloAdjacencyHoldTimeRem" { return "Mplsldphelloadjacencyholdtimerem" }
    if yname == "mplsLdpHelloAdjacencyHoldTime" { return "Mplsldphelloadjacencyholdtime" }
    if yname == "mplsLdpHelloAdjacencyType" { return "Mplsldphelloadjacencytype" }
    return ""
}

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetSegmentPath() string {
    return "mplsLdpHelloAdjacencyEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldppeerldpid) + "']" + "[mplsLdpHelloAdjacencyIndex='" + fmt.Sprintf("%v", mplsldphelloadjacencyentry.Mplsldphelloadjacencyindex) + "']"
}

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsldphelloadjacencyentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsldphelloadjacencyentry.Mplsldpentityindex
    leafs["mplsLdpPeerLdpId"] = mplsldphelloadjacencyentry.Mplsldppeerldpid
    leafs["mplsLdpHelloAdjacencyIndex"] = mplsldphelloadjacencyentry.Mplsldphelloadjacencyindex
    leafs["mplsLdpHelloAdjacencyHoldTimeRem"] = mplsldphelloadjacencyentry.Mplsldphelloadjacencyholdtimerem
    leafs["mplsLdpHelloAdjacencyHoldTime"] = mplsldphelloadjacencyentry.Mplsldphelloadjacencyholdtime
    leafs["mplsLdpHelloAdjacencyType"] = mplsldphelloadjacencyentry.Mplsldphelloadjacencytype
    return leafs
}

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetYangName() string { return "mplsLdpHelloAdjacencyEntry" }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) SetParent(parent types.Entity) { mplsldphelloadjacencyentry.parent = parent }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetParent() types.Entity { return mplsldphelloadjacencyentry.parent }

func (mplsldphelloadjacencyentry *MPLSLDPSTDMIB_Mplsldphelloadjacencytable_Mplsldphelloadjacencyentry) GetParentYangName() string { return "mplsLdpHelloAdjacencyTable" }

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
    parent types.Entity
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

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetFilter() yfilter.YFilter { return mplsinsegmentldplsptable.YFilter }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) SetFilter(yf yfilter.YFilter) { mplsinsegmentldplsptable.YFilter = yf }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetGoName(yname string) string {
    if yname == "mplsInSegmentLdpLspEntry" { return "Mplsinsegmentldplspentry" }
    return ""
}

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetSegmentPath() string {
    return "mplsInSegmentLdpLspTable"
}

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsInSegmentLdpLspEntry" {
        for _, c := range mplsinsegmentldplsptable.Mplsinsegmentldplspentry {
            if mplsinsegmentldplsptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry{}
        mplsinsegmentldplsptable.Mplsinsegmentldplspentry = append(mplsinsegmentldplsptable.Mplsinsegmentldplspentry, child)
        return &mplsinsegmentldplsptable.Mplsinsegmentldplspentry[len(mplsinsegmentldplsptable.Mplsinsegmentldplspentry)-1]
    }
    return nil
}

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsinsegmentldplsptable.Mplsinsegmentldplspentry {
        children[mplsinsegmentldplsptable.Mplsinsegmentldplspentry[i].GetSegmentPath()] = &mplsinsegmentldplsptable.Mplsinsegmentldplspentry[i]
    }
    return children
}

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetYangName() string { return "mplsInSegmentLdpLspTable" }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) SetParent(parent types.Entity) { mplsinsegmentldplsptable.parent = parent }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetParent() types.Entity { return mplsinsegmentldplsptable.parent }

func (mplsinsegmentldplsptable *MPLSLDPSTDMIB_Mplsinsegmentldplsptable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

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
    parent types.Entity
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

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetFilter() yfilter.YFilter { return mplsinsegmentldplspentry.YFilter }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) SetFilter(yf yfilter.YFilter) { mplsinsegmentldplspentry.YFilter = yf }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpPeerLdpId" { return "Mplsldppeerldpid" }
    if yname == "mplsInSegmentLdpLspIndex" { return "Mplsinsegmentldplspindex" }
    if yname == "mplsInSegmentLdpLspLabelType" { return "Mplsinsegmentldplsplabeltype" }
    if yname == "mplsInSegmentLdpLspType" { return "Mplsinsegmentldplsptype" }
    return ""
}

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetSegmentPath() string {
    return "mplsInSegmentLdpLspEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsldppeerldpid) + "']" + "[mplsInSegmentLdpLspIndex='" + fmt.Sprintf("%v", mplsinsegmentldplspentry.Mplsinsegmentldplspindex) + "']"
}

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsinsegmentldplspentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsinsegmentldplspentry.Mplsldpentityindex
    leafs["mplsLdpPeerLdpId"] = mplsinsegmentldplspentry.Mplsldppeerldpid
    leafs["mplsInSegmentLdpLspIndex"] = mplsinsegmentldplspentry.Mplsinsegmentldplspindex
    leafs["mplsInSegmentLdpLspLabelType"] = mplsinsegmentldplspentry.Mplsinsegmentldplsplabeltype
    leafs["mplsInSegmentLdpLspType"] = mplsinsegmentldplspentry.Mplsinsegmentldplsptype
    return leafs
}

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetYangName() string { return "mplsInSegmentLdpLspEntry" }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) SetParent(parent types.Entity) { mplsinsegmentldplspentry.parent = parent }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetParent() types.Entity { return mplsinsegmentldplspentry.parent }

func (mplsinsegmentldplspentry *MPLSLDPSTDMIB_Mplsinsegmentldplsptable_Mplsinsegmentldplspentry) GetParentYangName() string { return "mplsInSegmentLdpLspTable" }

// MPLSLDPSTDMIB_Mplsoutsegmentldplsptable
// A table of LDP LSP's which
// map to the mplsOutSegmentTable in the
// MPLS-LSR-STD-MIB.
type MPLSLDPSTDMIB_Mplsoutsegmentldplsptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents information on a single LDP LSP which is
    // represented by a session's index triple (mplsLdpEntityLdpId,
    // mplsLdpEntityIndex, mplsLdpPeerLdpId) AND the index
    // (mplsOutSegmentLdpLspIndex) for the mplsOutSegmentTable.  The information
    // contained in a row is read-only. The type is slice of
    // MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry.
    Mplsoutsegmentldplspentry []MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry
}

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetFilter() yfilter.YFilter { return mplsoutsegmentldplsptable.YFilter }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) SetFilter(yf yfilter.YFilter) { mplsoutsegmentldplsptable.YFilter = yf }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetGoName(yname string) string {
    if yname == "mplsOutSegmentLdpLspEntry" { return "Mplsoutsegmentldplspentry" }
    return ""
}

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetSegmentPath() string {
    return "mplsOutSegmentLdpLspTable"
}

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsOutSegmentLdpLspEntry" {
        for _, c := range mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry {
            if mplsoutsegmentldplsptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry{}
        mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry = append(mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry, child)
        return &mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry[len(mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry)-1]
    }
    return nil
}

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry {
        children[mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry[i].GetSegmentPath()] = &mplsoutsegmentldplsptable.Mplsoutsegmentldplspentry[i]
    }
    return children
}

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetYangName() string { return "mplsOutSegmentLdpLspTable" }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) SetParent(parent types.Entity) { mplsoutsegmentldplsptable.parent = parent }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetParent() types.Entity { return mplsoutsegmentldplsptable.parent }

func (mplsoutsegmentldplsptable *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

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
    parent types.Entity
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

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetFilter() yfilter.YFilter { return mplsoutsegmentldplspentry.YFilter }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) SetFilter(yf yfilter.YFilter) { mplsoutsegmentldplspentry.YFilter = yf }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpPeerLdpId" { return "Mplsldppeerldpid" }
    if yname == "mplsOutSegmentLdpLspIndex" { return "Mplsoutsegmentldplspindex" }
    if yname == "mplsOutSegmentLdpLspLabelType" { return "Mplsoutsegmentldplsplabeltype" }
    if yname == "mplsOutSegmentLdpLspType" { return "Mplsoutsegmentldplsptype" }
    return ""
}

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetSegmentPath() string {
    return "mplsOutSegmentLdpLspEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsldppeerldpid) + "']" + "[mplsOutSegmentLdpLspIndex='" + fmt.Sprintf("%v", mplsoutsegmentldplspentry.Mplsoutsegmentldplspindex) + "']"
}

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsoutsegmentldplspentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsoutsegmentldplspentry.Mplsldpentityindex
    leafs["mplsLdpPeerLdpId"] = mplsoutsegmentldplspentry.Mplsldppeerldpid
    leafs["mplsOutSegmentLdpLspIndex"] = mplsoutsegmentldplspentry.Mplsoutsegmentldplspindex
    leafs["mplsOutSegmentLdpLspLabelType"] = mplsoutsegmentldplspentry.Mplsoutsegmentldplsplabeltype
    leafs["mplsOutSegmentLdpLspType"] = mplsoutsegmentldplspentry.Mplsoutsegmentldplsptype
    return leafs
}

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetYangName() string { return "mplsOutSegmentLdpLspEntry" }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) SetParent(parent types.Entity) { mplsoutsegmentldplspentry.parent = parent }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetParent() types.Entity { return mplsoutsegmentldplspentry.parent }

func (mplsoutsegmentldplspentry *MPLSLDPSTDMIB_Mplsoutsegmentldplsptable_Mplsoutsegmentldplspentry) GetParentYangName() string { return "mplsOutSegmentLdpLspTable" }

// MPLSLDPSTDMIB_Mplsfectable
// This table represents the FEC
// (Forwarding Equivalence Class)
// Information associated with an LSP.
type MPLSLDPSTDMIB_Mplsfectable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each row represents a single FEC Element. The type is slice of
    // MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry.
    Mplsfecentry []MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry
}

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetFilter() yfilter.YFilter { return mplsfectable.YFilter }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) SetFilter(yf yfilter.YFilter) { mplsfectable.YFilter = yf }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetGoName(yname string) string {
    if yname == "mplsFecEntry" { return "Mplsfecentry" }
    return ""
}

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetSegmentPath() string {
    return "mplsFecTable"
}

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsFecEntry" {
        for _, c := range mplsfectable.Mplsfecentry {
            if mplsfectable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry{}
        mplsfectable.Mplsfecentry = append(mplsfectable.Mplsfecentry, child)
        return &mplsfectable.Mplsfecentry[len(mplsfectable.Mplsfecentry)-1]
    }
    return nil
}

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsfectable.Mplsfecentry {
        children[mplsfectable.Mplsfecentry[i].GetSegmentPath()] = &mplsfectable.Mplsfecentry[i]
    }
    return children
}

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetYangName() string { return "mplsFecTable" }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) SetParent(parent types.Entity) { mplsfectable.parent = parent }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetParent() types.Entity { return mplsfectable.parent }

func (mplsfectable *MPLSLDPSTDMIB_Mplsfectable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry
// Each row represents a single FEC Element.
type MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry struct {
    parent types.Entity
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

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetFilter() yfilter.YFilter { return mplsfecentry.YFilter }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) SetFilter(yf yfilter.YFilter) { mplsfecentry.YFilter = yf }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetGoName(yname string) string {
    if yname == "mplsFecIndex" { return "Mplsfecindex" }
    if yname == "mplsFecType" { return "Mplsfectype" }
    if yname == "mplsFecAddrPrefixLength" { return "Mplsfecaddrprefixlength" }
    if yname == "mplsFecAddrType" { return "Mplsfecaddrtype" }
    if yname == "mplsFecAddr" { return "Mplsfecaddr" }
    if yname == "mplsFecStorageType" { return "Mplsfecstoragetype" }
    if yname == "mplsFecRowStatus" { return "Mplsfecrowstatus" }
    return ""
}

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetSegmentPath() string {
    return "mplsFecEntry" + "[mplsFecIndex='" + fmt.Sprintf("%v", mplsfecentry.Mplsfecindex) + "']"
}

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsFecIndex"] = mplsfecentry.Mplsfecindex
    leafs["mplsFecType"] = mplsfecentry.Mplsfectype
    leafs["mplsFecAddrPrefixLength"] = mplsfecentry.Mplsfecaddrprefixlength
    leafs["mplsFecAddrType"] = mplsfecentry.Mplsfecaddrtype
    leafs["mplsFecAddr"] = mplsfecentry.Mplsfecaddr
    leafs["mplsFecStorageType"] = mplsfecentry.Mplsfecstoragetype
    leafs["mplsFecRowStatus"] = mplsfecentry.Mplsfecrowstatus
    return leafs
}

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetYangName() string { return "mplsFecEntry" }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) SetParent(parent types.Entity) { mplsfecentry.parent = parent }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetParent() types.Entity { return mplsfecentry.parent }

func (mplsfecentry *MPLSLDPSTDMIB_Mplsfectable_Mplsfecentry) GetParentYangName() string { return "mplsFecTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry represents a LDP LSP to FEC association. The type is slice of
    // MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry.
    Mplsldplspfecentry []MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry
}

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetFilter() yfilter.YFilter { return mplsldplspfectable.YFilter }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) SetFilter(yf yfilter.YFilter) { mplsldplspfectable.YFilter = yf }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetGoName(yname string) string {
    if yname == "mplsLdpLspFecEntry" { return "Mplsldplspfecentry" }
    return ""
}

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetSegmentPath() string {
    return "mplsLdpLspFecTable"
}

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpLspFecEntry" {
        for _, c := range mplsldplspfectable.Mplsldplspfecentry {
            if mplsldplspfectable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry{}
        mplsldplspfectable.Mplsldplspfecentry = append(mplsldplspfectable.Mplsldplspfecentry, child)
        return &mplsldplspfectable.Mplsldplspfecentry[len(mplsldplspfectable.Mplsldplspfecentry)-1]
    }
    return nil
}

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsldplspfectable.Mplsldplspfecentry {
        children[mplsldplspfectable.Mplsldplspfecentry[i].GetSegmentPath()] = &mplsldplspfectable.Mplsldplspfecentry[i]
    }
    return children
}

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetYangName() string { return "mplsLdpLspFecTable" }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) SetParent(parent types.Entity) { mplsldplspfectable.parent = parent }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetParent() types.Entity { return mplsldplspfectable.parent }

func (mplsldplspfectable *MPLSLDPSTDMIB_Mplsldplspfectable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry
// An entry represents a LDP LSP
// to FEC association.
type MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry struct {
    parent types.Entity
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

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetFilter() yfilter.YFilter { return mplsldplspfecentry.YFilter }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) SetFilter(yf yfilter.YFilter) { mplsldplspfecentry.YFilter = yf }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpPeerLdpId" { return "Mplsldppeerldpid" }
    if yname == "mplsLdpLspFecSegment" { return "Mplsldplspfecsegment" }
    if yname == "mplsLdpLspFecSegmentIndex" { return "Mplsldplspfecsegmentindex" }
    if yname == "mplsLdpLspFecIndex" { return "Mplsldplspfecindex" }
    if yname == "mplsLdpLspFecStorageType" { return "Mplsldplspfecstoragetype" }
    if yname == "mplsLdpLspFecRowStatus" { return "Mplsldplspfecrowstatus" }
    return ""
}

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetSegmentPath() string {
    return "mplsLdpLspFecEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldppeerldpid) + "']" + "[mplsLdpLspFecSegment='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldplspfecsegment) + "']" + "[mplsLdpLspFecSegmentIndex='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldplspfecsegmentindex) + "']" + "[mplsLdpLspFecIndex='" + fmt.Sprintf("%v", mplsldplspfecentry.Mplsldplspfecindex) + "']"
}

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsldplspfecentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsldplspfecentry.Mplsldpentityindex
    leafs["mplsLdpPeerLdpId"] = mplsldplspfecentry.Mplsldppeerldpid
    leafs["mplsLdpLspFecSegment"] = mplsldplspfecentry.Mplsldplspfecsegment
    leafs["mplsLdpLspFecSegmentIndex"] = mplsldplspfecentry.Mplsldplspfecsegmentindex
    leafs["mplsLdpLspFecIndex"] = mplsldplspfecentry.Mplsldplspfecindex
    leafs["mplsLdpLspFecStorageType"] = mplsldplspfecentry.Mplsldplspfecstoragetype
    leafs["mplsLdpLspFecRowStatus"] = mplsldplspfecentry.Mplsldplspfecrowstatus
    return leafs
}

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetYangName() string { return "mplsLdpLspFecEntry" }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) SetParent(parent types.Entity) { mplsldplspfecentry.parent = parent }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetParent() types.Entity { return mplsldplspfecentry.parent }

func (mplsldplspfecentry *MPLSLDPSTDMIB_Mplsldplspfectable_Mplsldplspfecentry) GetParentYangName() string { return "mplsLdpLspFecTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents information on a session's single next
    // hop address which was advertised in an Address Message from the LDP peer.
    // The information contained in a row is read-only. The type is slice of
    // MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry.
    Mplsldpsessionpeeraddrentry []MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry
}

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetFilter() yfilter.YFilter { return mplsldpsessionpeeraddrtable.YFilter }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) SetFilter(yf yfilter.YFilter) { mplsldpsessionpeeraddrtable.YFilter = yf }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetGoName(yname string) string {
    if yname == "mplsLdpSessionPeerAddrEntry" { return "Mplsldpsessionpeeraddrentry" }
    return ""
}

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetSegmentPath() string {
    return "mplsLdpSessionPeerAddrTable"
}

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpSessionPeerAddrEntry" {
        for _, c := range mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry {
            if mplsldpsessionpeeraddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry{}
        mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry = append(mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry, child)
        return &mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry[len(mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry)-1]
    }
    return nil
}

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry {
        children[mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry[i].GetSegmentPath()] = &mplsldpsessionpeeraddrtable.Mplsldpsessionpeeraddrentry[i]
    }
    return children
}

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetYangName() string { return "mplsLdpSessionPeerAddrTable" }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) SetParent(parent types.Entity) { mplsldpsessionpeeraddrtable.parent = parent }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetParent() types.Entity { return mplsldpsessionpeeraddrtable.parent }

func (mplsldpsessionpeeraddrtable *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable) GetParentYangName() string { return "MPLS-LDP-STD-MIB" }

// MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry
// An entry in this table represents information on
// a session's single next hop address which was
// advertised in an Address Message from the LDP peer.
// The information contained in a row is read-only.
type MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry struct {
    parent types.Entity
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

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetFilter() yfilter.YFilter { return mplsldpsessionpeeraddrentry.YFilter }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) SetFilter(yf yfilter.YFilter) { mplsldpsessionpeeraddrentry.YFilter = yf }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpPeerLdpId" { return "Mplsldppeerldpid" }
    if yname == "mplsLdpSessionPeerAddrIndex" { return "Mplsldpsessionpeeraddrindex" }
    if yname == "mplsLdpSessionPeerNextHopAddrType" { return "Mplsldpsessionpeernexthopaddrtype" }
    if yname == "mplsLdpSessionPeerNextHopAddr" { return "Mplsldpsessionpeernexthopaddr" }
    return ""
}

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetSegmentPath() string {
    return "mplsLdpSessionPeerAddrEntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldpentityindex) + "']" + "[mplsLdpPeerLdpId='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldppeerldpid) + "']" + "[mplsLdpSessionPeerAddrIndex='" + fmt.Sprintf("%v", mplsldpsessionpeeraddrentry.Mplsldpsessionpeeraddrindex) + "']"
}

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsldpsessionpeeraddrentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsldpsessionpeeraddrentry.Mplsldpentityindex
    leafs["mplsLdpPeerLdpId"] = mplsldpsessionpeeraddrentry.Mplsldppeerldpid
    leafs["mplsLdpSessionPeerAddrIndex"] = mplsldpsessionpeeraddrentry.Mplsldpsessionpeeraddrindex
    leafs["mplsLdpSessionPeerNextHopAddrType"] = mplsldpsessionpeeraddrentry.Mplsldpsessionpeernexthopaddrtype
    leafs["mplsLdpSessionPeerNextHopAddr"] = mplsldpsessionpeeraddrentry.Mplsldpsessionpeernexthopaddr
    return leafs
}

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetYangName() string { return "mplsLdpSessionPeerAddrEntry" }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) SetParent(parent types.Entity) { mplsldpsessionpeeraddrentry.parent = parent }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetParent() types.Entity { return mplsldpsessionpeeraddrentry.parent }

func (mplsldpsessionpeeraddrentry *MPLSLDPSTDMIB_Mplsldpsessionpeeraddrtable_Mplsldpsessionpeeraddrentry) GetParentYangName() string { return "mplsLdpSessionPeerAddrTable" }

