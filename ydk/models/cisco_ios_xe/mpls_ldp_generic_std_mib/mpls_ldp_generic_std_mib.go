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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The MPLS LDP Entity Generic Label Range (LR) Table.  The purpose of this
    // table is to provide a mechanism for configurating a contiguous range of
    // generic labels, or a 'label range' for LDP Entities.  LDP Entities which
    // use Generic Labels must have at least one entry in this table.  In other
    // words, this table 'extends' the mpldLdpEntityTable for Generic Labels.
    Mplsldpentitygenericlrtable MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable
}

func (mPLSLDPGENERICSTDMIB *MPLSLDPGENERICSTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSLDPGENERICSTDMIB.EntityData.YFilter = mPLSLDPGENERICSTDMIB.YFilter
    mPLSLDPGENERICSTDMIB.EntityData.YangName = "MPLS-LDP-GENERIC-STD-MIB"
    mPLSLDPGENERICSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSLDPGENERICSTDMIB.EntityData.ParentYangName = "MPLS-LDP-GENERIC-STD-MIB"
    mPLSLDPGENERICSTDMIB.EntityData.SegmentPath = "MPLS-LDP-GENERIC-STD-MIB:MPLS-LDP-GENERIC-STD-MIB"
    mPLSLDPGENERICSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSLDPGENERICSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSLDPGENERICSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSLDPGENERICSTDMIB.EntityData.Children = make(map[string]types.YChild)
    mPLSLDPGENERICSTDMIB.EntityData.Children["mplsLdpEntityGenericLRTable"] = types.YChild{"Mplsldpentitygenericlrtable", &mPLSLDPGENERICSTDMIB.Mplsldpentitygenericlrtable}
    mPLSLDPGENERICSTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mPLSLDPGENERICSTDMIB.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (mplsldpentitygenericlrtable *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable) GetEntityData() *types.CommonEntityData {
    mplsldpentitygenericlrtable.EntityData.YFilter = mplsldpentitygenericlrtable.YFilter
    mplsldpentitygenericlrtable.EntityData.YangName = "mplsLdpEntityGenericLRTable"
    mplsldpentitygenericlrtable.EntityData.BundleName = "cisco_ios_xe"
    mplsldpentitygenericlrtable.EntityData.ParentYangName = "MPLS-LDP-GENERIC-STD-MIB"
    mplsldpentitygenericlrtable.EntityData.SegmentPath = "mplsLdpEntityGenericLRTable"
    mplsldpentitygenericlrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpentitygenericlrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpentitygenericlrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpentitygenericlrtable.EntityData.Children = make(map[string]types.YChild)
    mplsldpentitygenericlrtable.EntityData.Children["mplsLdpEntityGenericLREntry"] = types.YChild{"Mplsldpentitygenericlrentry", nil}
    for i := range mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry {
        mplsldpentitygenericlrtable.EntityData.Children[types.GetSegmentPath(&mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry[i])] = types.YChild{"Mplsldpentitygenericlrentry", &mplsldpentitygenericlrtable.Mplsldpentitygenericlrentry[i]}
    }
    mplsldpentitygenericlrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsldpentitygenericlrtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (mplsldpentitygenericlrentry *MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry) GetEntityData() *types.CommonEntityData {
    mplsldpentitygenericlrentry.EntityData.YFilter = mplsldpentitygenericlrentry.YFilter
    mplsldpentitygenericlrentry.EntityData.YangName = "mplsLdpEntityGenericLREntry"
    mplsldpentitygenericlrentry.EntityData.BundleName = "cisco_ios_xe"
    mplsldpentitygenericlrentry.EntityData.ParentYangName = "mplsLdpEntityGenericLRTable"
    mplsldpentitygenericlrentry.EntityData.SegmentPath = "mplsLdpEntityGenericLREntry" + "[mplsLdpEntityLdpId='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentityldpid) + "']" + "[mplsLdpEntityIndex='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentityindex) + "']" + "[mplsLdpEntityGenericLRMin='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentitygenericlrmin) + "']" + "[mplsLdpEntityGenericLRMax='" + fmt.Sprintf("%v", mplsldpentitygenericlrentry.Mplsldpentitygenericlrmax) + "']"
    mplsldpentitygenericlrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsldpentitygenericlrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsldpentitygenericlrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsldpentitygenericlrentry.EntityData.Children = make(map[string]types.YChild)
    mplsldpentitygenericlrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityLdpId"] = types.YLeaf{"Mplsldpentityldpid", mplsldpentitygenericlrentry.Mplsldpentityldpid}
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityIndex"] = types.YLeaf{"Mplsldpentityindex", mplsldpentitygenericlrentry.Mplsldpentityindex}
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityGenericLRMin"] = types.YLeaf{"Mplsldpentitygenericlrmin", mplsldpentitygenericlrentry.Mplsldpentitygenericlrmin}
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityGenericLRMax"] = types.YLeaf{"Mplsldpentitygenericlrmax", mplsldpentitygenericlrentry.Mplsldpentitygenericlrmax}
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityGenericLabelSpace"] = types.YLeaf{"Mplsldpentitygenericlabelspace", mplsldpentitygenericlrentry.Mplsldpentitygenericlabelspace}
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityGenericIfIndexOrZero"] = types.YLeaf{"Mplsldpentitygenericifindexorzero", mplsldpentitygenericlrentry.Mplsldpentitygenericifindexorzero}
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityGenericLRStorageType"] = types.YLeaf{"Mplsldpentitygenericlrstoragetype", mplsldpentitygenericlrentry.Mplsldpentitygenericlrstoragetype}
    mplsldpentitygenericlrentry.EntityData.Leafs["mplsLdpEntityGenericLRRowStatus"] = types.YLeaf{"Mplsldpentitygenericlrrowstatus", mplsldpentitygenericlrentry.Mplsldpentitygenericlrrowstatus}
    return &(mplsldpentitygenericlrentry.EntityData)
}

// MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace represents means that the label space type is per Interface.
type MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace string

const (
    MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace_perPlatform MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace = "perPlatform"

    MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace_perInterface MPLSLDPGENERICSTDMIB_Mplsldpentitygenericlrtable_Mplsldpentitygenericlrentry_Mplsldpentitygenericlabelspace = "perInterface"
)

