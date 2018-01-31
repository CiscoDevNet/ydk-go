// Copyright (C) The Internet Society (year). The
// initial version of this MIB module was published
// in RFC 3815. For full legal notices see the RFC
// itself or see:
// http://www.ietf.org/copyrights/ianamib.html
// 
// This MIB contains managed object definitions for
// configuring and monitoring the Multiprotocol Label
// Switching (MPLS), Label Distribution Protocol (LDP),
// utilizing ethernet as the Layer 2 media.
package mpls_ldp_generic_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_ldp_generic_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:MPLS-LDP-GENERIC-STD-MIB MPLS-LDP-GENERIC-STD-MIB}", reflect.TypeOf(MPLSLDPGENERICSTDMIB{}))
    ydk.RegisterEntity("MPLS-LDP-GENERIC-STD-MIB:MPLS-LDP-GENERIC-STD-MIB", reflect.TypeOf(MPLSLDPGENERICSTDMIB{}))
}

// MPLSLDPGENERICSTDMIB
type MPLSLDPGENERICSTDMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The MPLS LDP Entity Generic Label Range (LR) Table.  The purpose of this
    // table is to provide a mechanism for configurating a contiguous range of
    // generic labels, or a 'label range' for LDP Entities.  LDP Entities which
    // use Generic Labels must have at least one entry in this table.  In other
    // words, this table 'extends' the mpldLdpEntityTable for Generic Labels.
    Mplsldpentitygenericlrtable MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable
}

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetFilter() yfilter.YFilter { return mPLSLDPGENERICSTDMIB.YFilter }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) SetFilter(yf yfilter.YFilter) { mPLSLDPGENERICSTDMIB.YFilter = yf }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetGoName(yname string) string {
    if yname == "mplsLdpEntityGenericLRTable" { return "Mplsldpentitygenericlrtable" }
    return ""
}

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetSegmentPath() string {
    return "MPLS-LDP-GENERIC-STD-MIB:MPLS-LDP-GENERIC-STD-MIB"
}

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpEntityGenericLRTable" {
        return &mPLSLDPGENERICSTDMIB.Mplsldpentitygenericlrtable
    }
    return nil
}

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mplsLdpEntityGenericLRTable"] = &mPLSLDPGENERICSTDMIB.Mplsldpentitygenericlrtable
    return children
}

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetYangName() string { return "MPLS-LDP-GENERIC-STD-MIB" }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) SetParent(parent types.Entity) { mPLSLDPGENERICSTDMIB.parent = parent }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetParent() types.Entity { return mPLSLDPGENERICSTDMIB.parent }

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetParentYangName() string { return "MPLS-LDP-GENERIC-STD-MIB" }

// MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable
// The MPLS LDP Entity Generic Label Range (LR)
// Table.
// 
// The purpose of this table is to provide a mechanism
// for configurating a contiguous range of generic labels,
// or a 'label range' for LDP Entities.
// 
// LDP Entities which use Generic Labels must have at least
// one entry in this table.  In other words, this table
// 'extends' the mpldLdpEntityTable for Generic Labels.
type MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in the LDP Entity Generic Label Range (LR) Table.  One entry in this
    // table contains information on a single range of labels represented by the
    // configured Upper and Lower Bounds pairs.  NOTE: there is NO corresponding
    // LDP message which relates to the information in this table, however, this
    // table does provide a way for a user to 'reserve' a generic label range. 
    // NOTE:  The ranges for a specific LDP Entity are UNIQUE and non-overlapping.
    // A row will not be created unless a unique and non-overlapping range is
    // specified. The type is slice of
    // MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry.
    Mplsldpentitygenericlrentry []MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry
}

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetFilter() yfilter.YFilter { return mplsldpentitygenericlrtable.YFilter }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) SetFilter(yf yfilter.YFilter) { mplsldpentitygenericlrtable.YFilter = yf }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetGoName(yname string) string {
    if yname == "mplsLdpEntityGenericLREntry" { return "Mplsldpentitygenericlrentry" }
    return ""
}

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetSegmentPath() string {
    return "mplsLdpEntityGenericLRTable"
}

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLdpEntityGenericLREntry" {
        for _, c := range mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry {
            if mplsldpentitygenericlrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry{}
        mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry = append(mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry, child)
        return &mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry[len(mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry)-1]
    }
    return nil
}

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry {
        children[mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry[i].GetSegmentPath()] = &mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry[i]
    }
    return children
}

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetYangName() string { return "mplsLdpEntityGenericLRTable" }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) SetParent(parent types.Entity) { mplsldpentitygenericlrtable.parent = parent }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetParent() types.Entity { return mplsldpentitygenericlrtable.parent }

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetParentYangName() string { return "MPLS-LDP-GENERIC-STD-MIB" }

// MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry
// A row in the LDP Entity Generic Label
// Range (LR) Table.  One entry in this table contains
// information on a single range of labels
// represented by the configured Upper and Lower
// Bounds pairs.  NOTE: there is NO corresponding
// LDP message which relates to the information
// in this table, however, this table does provide
// a way for a user to 'reserve' a generic label
// range.
// 
// NOTE:  The ranges for a specific LDP Entity
// are UNIQUE and non-overlapping.
// 
// A row will not be created unless a unique and
// non-overlapping range is specified.
type MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityldpid
    Mplsldpentityldpid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_Mplsldpentitytable_Mplsldpentityentry_Mplsldpentityindex
    Mplsldpentityindex interface{}

    // This attribute is a key. The minimum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    Mplsldpentitygenericlrmin interface{}

    // This attribute is a key. The maximum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    Mplsldpentitygenericlrmax interface{}

    // This value of this object is perPlatform(1), then this means that the label
    // space type is per platform.  If this object is perInterface(2), then this
    // means that the label space type is per Interface. The type is
    // Mplsldpentitygenericlabelspace.
    Mplsldpentitygenericlabelspace interface{}

    // This value represents either the InterfaceIndex of the 'ifLayer' where
    // these Generic Label would be created,   or 0 (zero).  The value of zero
    // means that the InterfaceIndex is not known.  However, if the InterfaceIndex
    // is known, then it must be represented by this value.  If an InterfaceIndex
    // becomes known, then the network management entity (e.g., SNMP agent)
    // responsible for this object MUST change the value from 0 (zero) to the
    // value of the InterfaceIndex. The type is interface{} with range:
    // 0..2147483647.
    Mplsldpentitygenericifindexorzero interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    Mplsldpentitygenericlrstoragetype interface{}

    // The status of this conceptual row.  All writable objects in this row may be
    // modified at any time, however, as described in  detail in the section
    // entitled, 'Changing Values After Session Establishment', and again
    // described in the DESCRIPTION clause of the mplsLdpEntityAdminStatus object,
    // if a session has been initiated with a Peer, changing objects in this table
    // will wreak havoc with the session and interrupt traffic. To repeat again: 
    // the recommended procedure is to set the mplsLdpEntityAdminStatus to down,
    // thereby explicitly causing a session to be torn down. Then, change objects
    // in this entry, then set the mplsLdpEntityAdminStatus to enable which
    // enables a new session to be initiated.  There must exist at least one entry
    // in this table for every LDP Entity that has a generic label configured. The
    // type is RowStatus.
    Mplsldpentitygenericlrrowstatus interface{}
}

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetFilter() yfilter.YFilter { return mplsldpentitygenericlrentry.YFilter }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) SetFilter(yf yfilter.YFilter) { mplsldpentitygenericlrentry.YFilter = yf }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetGoName(yname string) string {
    if yname == "mplsLdpEntityLdpId" { return "Mplsldpentityldpid" }
    if yname == "mplsLdpEntityIndex" { return "Mplsldpentityindex" }
    if yname == "mplsLdpEntityGenericLRMin" { return "Mplsldpentitygenericlrmin" }
    if yname == "mplsLdpEntityGenericLRMax" { return "Mplsldpentitygenericlrmax" }
    if yname == "mplsLdpEntityGenericLabelSpace" { return "Mplsldpentitygenericlabelspace" }
    if yname == "mplsLdpEntityGenericIfIndexOrZero" { return "Mplsldpentitygenericifindexorzero" }
    if yname == "mplsLdpEntityGenericLRStorageType" { return "Mplsldpentitygenericlrstoragetype" }
    if yname == "mplsLdpEntityGenericLRRowStatus" { return "Mplsldpentitygenericlrrowstatus" }
    return ""
}

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetSegmentPath() string {
    return "mplsLdpEntityGenericLREntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentityindex) + "']" + "[mplsLdpEntityGenericLRMin='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentitygenericlrmin) + "']" + "[mplsLdpEntityGenericLRMax='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentitygenericlrmax) + "']"
}

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLdpEntityLdpId"] = mplsldpentitygenericlrentry.Mplsldpentityldpid
    leafs["mplsLdpEntityIndex"] = mplsldpentitygenericlrentry.Mplsldpentityindex
    leafs["mplsLdpEntityGenericLRMin"] = mplsldpentitygenericlrentry.Mplsldpentitygenericlrmin
    leafs["mplsLdpEntityGenericLRMax"] = mplsldpentitygenericlrentry.Mplsldpentitygenericlrmax
    leafs["mplsLdpEntityGenericLabelSpace"] = mplsldpentitygenericlrentry.Mplsldpentitygenericlabelspace
    leafs["mplsLdpEntityGenericIfIndexOrZero"] = mplsldpentitygenericlrentry.Mplsldpentitygenericifindexorzero
    leafs["mplsLdpEntityGenericLRStorageType"] = mplsldpentitygenericlrentry.Mplsldpentitygenericlrstoragetype
    leafs["mplsLdpEntityGenericLRRowStatus"] = mplsldpentitygenericlrentry.Mplsldpentitygenericlrrowstatus
    return leafs
}

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetYangName() string { return "mplsLdpEntityGenericLREntry" }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) SetParent(parent types.Entity) { mplsldpentitygenericlrentry.parent = parent }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetParent() types.Entity { return mplsldpentitygenericlrentry.parent }

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetParentYangName() string { return "mplsLdpEntityGenericLRTable" }

// MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace represents means that the label space type is per Interface.
type MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace string

const (
    MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace_perPlatform MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace = "perPlatform"

    MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace_perInterface MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace = "perInterface"
)

