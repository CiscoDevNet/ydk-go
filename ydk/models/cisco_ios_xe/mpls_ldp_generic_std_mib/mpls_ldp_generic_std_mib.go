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
    MplsLdpEntityGenericLRTable MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable
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

    mPLSLDPGENERICSTDMIB.EntityData.Children = types.NewOrderedMap()
    mPLSLDPGENERICSTDMIB.EntityData.Children.Append("mplsLdpEntityGenericLRTable", types.YChild{"MplsLdpEntityGenericLRTable", &mPLSLDPGENERICSTDMIB.MplsLdpEntityGenericLRTable})
    mPLSLDPGENERICSTDMIB.EntityData.Leafs = types.NewOrderedMap()

    mPLSLDPGENERICSTDMIB.EntityData.YListKeys = []string {}

    return &(mPLSLDPGENERICSTDMIB.EntityData)
}

// MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable
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
type MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable struct {
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
    // MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry.
    MplsLdpEntityGenericLREntry []*MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry
}

func (mplsLdpEntityGenericLRTable *MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable) GetEntityData() *types.CommonEntityData {
    mplsLdpEntityGenericLRTable.EntityData.YFilter = mplsLdpEntityGenericLRTable.YFilter
    mplsLdpEntityGenericLRTable.EntityData.YangName = "mplsLdpEntityGenericLRTable"
    mplsLdpEntityGenericLRTable.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpEntityGenericLRTable.EntityData.ParentYangName = "MPLS-LDP-GENERIC-STD-MIB"
    mplsLdpEntityGenericLRTable.EntityData.SegmentPath = "mplsLdpEntityGenericLRTable"
    mplsLdpEntityGenericLRTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpEntityGenericLRTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpEntityGenericLRTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpEntityGenericLRTable.EntityData.Children = types.NewOrderedMap()
    mplsLdpEntityGenericLRTable.EntityData.Children.Append("mplsLdpEntityGenericLREntry", types.YChild{"MplsLdpEntityGenericLREntry", nil})
    for i := range mplsLdpEntityGenericLRTable.MplsLdpEntityGenericLREntry {
        mplsLdpEntityGenericLRTable.EntityData.Children.Append(types.GetSegmentPath(mplsLdpEntityGenericLRTable.MplsLdpEntityGenericLREntry[i]), types.YChild{"MplsLdpEntityGenericLREntry", mplsLdpEntityGenericLRTable.MplsLdpEntityGenericLREntry[i]})
    }
    mplsLdpEntityGenericLRTable.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpEntityGenericLRTable.EntityData.YListKeys = []string {}

    return &(mplsLdpEntityGenericLRTable.EntityData)
}

// MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry
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
type MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityLdpId
    MplsLdpEntityLdpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // mpls_ldp_std_mib.MPLSLDPSTDMIB_MplsLdpEntityTable_MplsLdpEntityEntry_MplsLdpEntityIndex
    MplsLdpEntityIndex interface{}

    // This attribute is a key. The minimum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    MplsLdpEntityGenericLRMin interface{}

    // This attribute is a key. The maximum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    MplsLdpEntityGenericLRMax interface{}

    // This value of this object is perPlatform(1), then this means that the label
    // space type is per platform.  If this object is perInterface(2), then this
    // means that the label space type is per Interface. The type is
    // MplsLdpEntityGenericLabelSpace.
    MplsLdpEntityGenericLabelSpace interface{}

    // This value represents either the InterfaceIndex of the 'ifLayer' where
    // these Generic Label would be created,   or 0 (zero).  The value of zero
    // means that the InterfaceIndex is not known.  However, if the InterfaceIndex
    // is known, then it must be represented by this value.  If an InterfaceIndex
    // becomes known, then the network management entity (e.g., SNMP agent)
    // responsible for this object MUST change the value from 0 (zero) to the
    // value of the InterfaceIndex. The type is interface{} with range:
    // 0..2147483647.
    MplsLdpEntityGenericIfIndexOrZero interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent(4)' need not allow write-access to any columnar objects in the
    // row. The type is StorageType.
    MplsLdpEntityGenericLRStorageType interface{}

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
    MplsLdpEntityGenericLRRowStatus interface{}
}

func (mplsLdpEntityGenericLREntry *MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry) GetEntityData() *types.CommonEntityData {
    mplsLdpEntityGenericLREntry.EntityData.YFilter = mplsLdpEntityGenericLREntry.YFilter
    mplsLdpEntityGenericLREntry.EntityData.YangName = "mplsLdpEntityGenericLREntry"
    mplsLdpEntityGenericLREntry.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpEntityGenericLREntry.EntityData.ParentYangName = "mplsLdpEntityGenericLRTable"
    mplsLdpEntityGenericLREntry.EntityData.SegmentPath = "mplsLdpEntityGenericLREntry" + types.AddKeyToken(mplsLdpEntityGenericLREntry.MplsLdpEntityLdpId, "mplsLdpEntityLdpId") + types.AddKeyToken(mplsLdpEntityGenericLREntry.MplsLdpEntityIndex, "mplsLdpEntityIndex") + types.AddKeyToken(mplsLdpEntityGenericLREntry.MplsLdpEntityGenericLRMin, "mplsLdpEntityGenericLRMin") + types.AddKeyToken(mplsLdpEntityGenericLREntry.MplsLdpEntityGenericLRMax, "mplsLdpEntityGenericLRMax")
    mplsLdpEntityGenericLREntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpEntityGenericLREntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpEntityGenericLREntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpEntityGenericLREntry.EntityData.Children = types.NewOrderedMap()
    mplsLdpEntityGenericLREntry.EntityData.Leafs = types.NewOrderedMap()
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityLdpId", types.YLeaf{"MplsLdpEntityLdpId", mplsLdpEntityGenericLREntry.MplsLdpEntityLdpId})
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityIndex", types.YLeaf{"MplsLdpEntityIndex", mplsLdpEntityGenericLREntry.MplsLdpEntityIndex})
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityGenericLRMin", types.YLeaf{"MplsLdpEntityGenericLRMin", mplsLdpEntityGenericLREntry.MplsLdpEntityGenericLRMin})
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityGenericLRMax", types.YLeaf{"MplsLdpEntityGenericLRMax", mplsLdpEntityGenericLREntry.MplsLdpEntityGenericLRMax})
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityGenericLabelSpace", types.YLeaf{"MplsLdpEntityGenericLabelSpace", mplsLdpEntityGenericLREntry.MplsLdpEntityGenericLabelSpace})
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityGenericIfIndexOrZero", types.YLeaf{"MplsLdpEntityGenericIfIndexOrZero", mplsLdpEntityGenericLREntry.MplsLdpEntityGenericIfIndexOrZero})
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityGenericLRStorageType", types.YLeaf{"MplsLdpEntityGenericLRStorageType", mplsLdpEntityGenericLREntry.MplsLdpEntityGenericLRStorageType})
    mplsLdpEntityGenericLREntry.EntityData.Leafs.Append("mplsLdpEntityGenericLRRowStatus", types.YLeaf{"MplsLdpEntityGenericLRRowStatus", mplsLdpEntityGenericLREntry.MplsLdpEntityGenericLRRowStatus})

    mplsLdpEntityGenericLREntry.EntityData.YListKeys = []string {"MplsLdpEntityLdpId", "MplsLdpEntityIndex", "MplsLdpEntityGenericLRMin", "MplsLdpEntityGenericLRMax"}

    return &(mplsLdpEntityGenericLREntry.EntityData)
}

// MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry_MplsLdpEntityGenericLabelSpace represents means that the label space type is per Interface.
type MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry_MplsLdpEntityGenericLabelSpace string

const (
    MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry_MplsLdpEntityGenericLabelSpace_perPlatform MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry_MplsLdpEntityGenericLabelSpace = "perPlatform"

    MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry_MplsLdpEntityGenericLabelSpace_perInterface MPLSLDPGENERICSTDMIB_MplsLdpEntityGenericLRTable_MplsLdpEntityGenericLREntry_MplsLdpEntityGenericLabelSpace = "perInterface"
)

