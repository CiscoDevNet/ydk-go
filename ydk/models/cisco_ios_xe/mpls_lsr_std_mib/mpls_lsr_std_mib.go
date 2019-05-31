// This MIB module contains managed object definitions for
// the Multiprotocol Label Switching (MPLS) Router as
// defined in: Rosen, E., Viswanathan, A., and R.
// Callon, Multiprotocol Label Switching Architecture,
// RFC 3031, January 2001.
// 
// Copyright (C) The Internet Society (2004). The
// initial version of this MIB module was published
// in RFC 3812. For full legal notices see the RFC
// itself or see:
// http://www.ietf.org/copyrights/ianamib.html
package mpls_lsr_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_lsr_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:MPLS-LSR-STD-MIB MPLS-LSR-STD-MIB}", reflect.TypeOf(MPLSLSRSTDMIB{}))
    ydk.RegisterEntity("MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB", reflect.TypeOf(MPLSLSRSTDMIB{}))
}

// MPLSLSRSTDMIB
type MPLSLSRSTDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    MplsLsrObjects MPLSLSRSTDMIB_MplsLsrObjects

    // This table specifies per-interface MPLS capability and associated
    // information.
    MplsInterfaceTable MPLSLSRSTDMIB_MplsInterfaceTable

    // This table contains a description of the incoming MPLS segments (labels) to
    // an LSR and their associated parameters. The index for this table is
    // mplsInSegmentIndex. The index structure of this table is specifically
    // designed to handle many different MPLS implementations that manage their
    // labels both in a distributed and centralized manner. The table is also
    // designed to handle existing MPLS labels as defined in RFC3031 as well as
    // longer ones that may be necessary in the future.  In cases where the label
    // cannot fit into the mplsInSegmentLabel object, the mplsInSegmentLabelPtr
    // will indicate this by being set to the first accessible column in the
    // appropriate extension table's row. In this case an additional table MUST be
    // provided and MUST be indexed by at least the indexes used by this table. In
    // all other cases when the label is represented within the mplsInSegmentLabel
    // object, the mplsInSegmentLabelPtr MUST be set to 0.0. Due to the fact that
    // MPLS labels may not exceed 24 bits, the mplsInSegmentLabelPtr object is
    // only a provision for future-proofing the MIB module. Thus, the definition
    // of any extension tables is beyond the scope of this MIB module.
    MplsInSegmentTable MPLSLSRSTDMIB_MplsInSegmentTable

    // This table contains a representation of the outgoing segments from an LSR.
    MplsOutSegmentTable MPLSLSRSTDMIB_MplsOutSegmentTable

    // This table specifies information for switching between LSP segments.  It
    // supports point-to-point, point-to-multipoint and multipoint-to-point
    // connections.  mplsLabelStackTable specifies the label stack information for
    // a cross-connect LSR and is referred to from mplsXCTable.
    MplsXCTable MPLSLSRSTDMIB_MplsXCTable

    // This table specifies the label stack to be pushed onto a packet, beneath
    // the top label.  Entries into this table are referred to from mplsXCTable.
    MplsLabelStackTable MPLSLSRSTDMIB_MplsLabelStackTable

    // This table specifies the mapping from the mplsInSegmentIndex to the
    // corresponding mplsInSegmentInterface and mplsInSegmentLabel objects. The
    // purpose of this table is to provide the manager with an alternative means
    // by which to locate in-segments.
    MplsInSegmentMapTable MPLSLSRSTDMIB_MplsInSegmentMapTable
}

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSLSRSTDMIB.EntityData.YFilter = mPLSLSRSTDMIB.YFilter
    mPLSLSRSTDMIB.EntityData.YangName = "MPLS-LSR-STD-MIB"
    mPLSLSRSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSLSRSTDMIB.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mPLSLSRSTDMIB.EntityData.SegmentPath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB"
    mPLSLSRSTDMIB.EntityData.AbsolutePath = mPLSLSRSTDMIB.EntityData.SegmentPath
    mPLSLSRSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSLSRSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSLSRSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSLSRSTDMIB.EntityData.Children = types.NewOrderedMap()
    mPLSLSRSTDMIB.EntityData.Children.Append("mplsLsrObjects", types.YChild{"MplsLsrObjects", &mPLSLSRSTDMIB.MplsLsrObjects})
    mPLSLSRSTDMIB.EntityData.Children.Append("mplsInterfaceTable", types.YChild{"MplsInterfaceTable", &mPLSLSRSTDMIB.MplsInterfaceTable})
    mPLSLSRSTDMIB.EntityData.Children.Append("mplsInSegmentTable", types.YChild{"MplsInSegmentTable", &mPLSLSRSTDMIB.MplsInSegmentTable})
    mPLSLSRSTDMIB.EntityData.Children.Append("mplsOutSegmentTable", types.YChild{"MplsOutSegmentTable", &mPLSLSRSTDMIB.MplsOutSegmentTable})
    mPLSLSRSTDMIB.EntityData.Children.Append("mplsXCTable", types.YChild{"MplsXCTable", &mPLSLSRSTDMIB.MplsXCTable})
    mPLSLSRSTDMIB.EntityData.Children.Append("mplsLabelStackTable", types.YChild{"MplsLabelStackTable", &mPLSLSRSTDMIB.MplsLabelStackTable})
    mPLSLSRSTDMIB.EntityData.Children.Append("mplsInSegmentMapTable", types.YChild{"MplsInSegmentMapTable", &mPLSLSRSTDMIB.MplsInSegmentMapTable})
    mPLSLSRSTDMIB.EntityData.Leafs = types.NewOrderedMap()

    mPLSLSRSTDMIB.EntityData.YListKeys = []string {}

    return &(mPLSLSRSTDMIB.EntityData)
}

// MPLSLSRSTDMIB_MplsLsrObjects
type MPLSLSRSTDMIB_MplsLsrObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains the next available value to be used for
    // mplsInSegmentIndex when creating entries in the mplsInSegmentTable. The
    // special value of a string containing the single octet 0x00 indicates that
    // no new entries can be created in this table. Agents not allowing managers
    // to create entries in this table MUST set this object to this special value.
    // The type is string with length: 1..24.
    MplsInSegmentIndexNext interface{}

    // This object contains the next available value to be used for
    // mplsOutSegmentIndex when creating entries in the mplsOutSegmentTable. The
    // special value of a string containing the single octet 0x00 indicates that
    // no new entries can be created in this table. Agents not allowing managers
    // to create entries in this table MUST set this object to this special value.
    // The type is string with length: 1..24.
    MplsOutSegmentIndexNext interface{}

    // This object contains the next available value to be used for mplsXCIndex
    // when creating entries in the mplsXCTable. A special value of the zero
    // length string indicates that no more new entries can be created in the
    // relevant table.  Agents not allowing managers to create entries in this
    // table MUST set this value to the zero length string. The type is string
    // with length: 1..24.
    MplsXCIndexNext interface{}

    // The maximum stack depth supported by this LSR. The type is interface{} with
    // range: 1..2147483647.
    MplsMaxLabelStackDepth interface{}

    // This object contains the next available value to be used for
    // mplsLabelStackIndex when creating entries in the mplsLabelStackTable. The
    // special string containing the single octet 0x00 indicates that no more new
    // entries can be created in the relevant table.  Agents not allowing managers
    // to create entries in this table MUST set this value to the string
    // containing the single octet 0x00. The type is string with length: 1..24.
    MplsLabelStackIndexNext interface{}

    // If this object is set to true(1), then it enables the emission of mplsXCUp
    // and mplsXCDown notifications; otherwise these notifications are not
    // emitted. The type is bool.
    MplsXCNotificationsEnable interface{}
}

func (mplsLsrObjects *MPLSLSRSTDMIB_MplsLsrObjects) GetEntityData() *types.CommonEntityData {
    mplsLsrObjects.EntityData.YFilter = mplsLsrObjects.YFilter
    mplsLsrObjects.EntityData.YangName = "mplsLsrObjects"
    mplsLsrObjects.EntityData.BundleName = "cisco_ios_xe"
    mplsLsrObjects.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsLsrObjects.EntityData.SegmentPath = "mplsLsrObjects"
    mplsLsrObjects.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/" + mplsLsrObjects.EntityData.SegmentPath
    mplsLsrObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLsrObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLsrObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLsrObjects.EntityData.Children = types.NewOrderedMap()
    mplsLsrObjects.EntityData.Leafs = types.NewOrderedMap()
    mplsLsrObjects.EntityData.Leafs.Append("mplsInSegmentIndexNext", types.YLeaf{"MplsInSegmentIndexNext", mplsLsrObjects.MplsInSegmentIndexNext})
    mplsLsrObjects.EntityData.Leafs.Append("mplsOutSegmentIndexNext", types.YLeaf{"MplsOutSegmentIndexNext", mplsLsrObjects.MplsOutSegmentIndexNext})
    mplsLsrObjects.EntityData.Leafs.Append("mplsXCIndexNext", types.YLeaf{"MplsXCIndexNext", mplsLsrObjects.MplsXCIndexNext})
    mplsLsrObjects.EntityData.Leafs.Append("mplsMaxLabelStackDepth", types.YLeaf{"MplsMaxLabelStackDepth", mplsLsrObjects.MplsMaxLabelStackDepth})
    mplsLsrObjects.EntityData.Leafs.Append("mplsLabelStackIndexNext", types.YLeaf{"MplsLabelStackIndexNext", mplsLsrObjects.MplsLabelStackIndexNext})
    mplsLsrObjects.EntityData.Leafs.Append("mplsXCNotificationsEnable", types.YLeaf{"MplsXCNotificationsEnable", mplsLsrObjects.MplsXCNotificationsEnable})

    mplsLsrObjects.EntityData.YListKeys = []string {}

    return &(mplsLsrObjects.EntityData)
}

// MPLSLSRSTDMIB_MplsInterfaceTable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSLSRSTDMIB_MplsInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in this table is created automatically by an LSR for every
    // interface capable of supporting MPLS and which is configured to do so. A
    // conceptual row in this table will exist if and only if a corresponding
    // entry in ifTable exists with ifType = mpls(166). If this associated entry
    // in ifTable is operationally disabled (thus removing MPLS capabilities on
    // that interface), the corresponding entry in this table MUST be deleted
    // shortly thereafter. An conceptual row with index 0 is created if the LSR
    // supports per-platform labels. This conceptual row represents the
    // per-platform label space and contains parameters that apply to all
    // interfaces that participate in the per-platform label space. Other
    // conceptual rows in this table represent MPLS interfaces that may
    // participate in either the per-platform or per- interface label spaces, or
    // both.  Implementations that either only support per-platform labels, or
    // have only them configured, may choose to return just the mplsInterfaceEntry
    // of 0 and not return the other rows. This will greatly reduce the number of
    // objects returned. Further information about label space participation of an
    // interface is provided in the DESCRIPTION clause of
    // mplsInterfaceLabelParticipationType. The type is slice of
    // MPLSLSRSTDMIB_MplsInterfaceTable_MplsInterfaceEntry.
    MplsInterfaceEntry []*MPLSLSRSTDMIB_MplsInterfaceTable_MplsInterfaceEntry
}

func (mplsInterfaceTable *MPLSLSRSTDMIB_MplsInterfaceTable) GetEntityData() *types.CommonEntityData {
    mplsInterfaceTable.EntityData.YFilter = mplsInterfaceTable.YFilter
    mplsInterfaceTable.EntityData.YangName = "mplsInterfaceTable"
    mplsInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    mplsInterfaceTable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsInterfaceTable.EntityData.SegmentPath = "mplsInterfaceTable"
    mplsInterfaceTable.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/" + mplsInterfaceTable.EntityData.SegmentPath
    mplsInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInterfaceTable.EntityData.Children = types.NewOrderedMap()
    mplsInterfaceTable.EntityData.Children.Append("mplsInterfaceEntry", types.YChild{"MplsInterfaceEntry", nil})
    for i := range mplsInterfaceTable.MplsInterfaceEntry {
        mplsInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(mplsInterfaceTable.MplsInterfaceEntry[i]), types.YChild{"MplsInterfaceEntry", mplsInterfaceTable.MplsInterfaceEntry[i]})
    }
    mplsInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    mplsInterfaceTable.EntityData.YListKeys = []string {}

    return &(mplsInterfaceTable.EntityData)
}

// MPLSLSRSTDMIB_MplsInterfaceTable_MplsInterfaceEntry
// A conceptual row in this table is created
// automatically by an LSR for every interface capable
// of supporting MPLS and which is configured to do so.
// A conceptual row in this table will exist if and only if
// a corresponding entry in ifTable exists with ifType =
// mpls(166). If this associated entry in ifTable is
// operationally disabled (thus removing MPLS
// capabilities on that interface), the corresponding
// entry in this table MUST be deleted shortly thereafter.
// An conceptual row with index 0 is created if the LSR
// supports per-platform labels. This conceptual row
// represents the per-platform label space and contains
// parameters that apply to all interfaces that participate
// in the per-platform label space. Other conceptual rows
// in this table represent MPLS interfaces that may
// participate in either the per-platform or per-
// interface label spaces, or both.  Implementations
// that either only support per-platform labels,
// or have only them configured, may choose to return
// just the mplsInterfaceEntry of 0 and not return
// the other rows. This will greatly reduce the number
// of objects returned. Further information about label
// space participation of an interface is provided in
// the DESCRIPTION clause of
// mplsInterfaceLabelParticipationType.
type MPLSLSRSTDMIB_MplsInterfaceTable_MplsInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is a unique index for an entry in the
    // MplsInterfaceTable.  A non-zero index for an entry indicates the ifIndex
    // for the corresponding interface entry of the MPLS-layer in the ifTable. The
    // entry with index 0 represents the per-platform label space and contains
    // parameters that apply to all interfaces that participate in the
    // per-platform label space. Other entries defined in this table represent
    // additional MPLS interfaces that may participate in either the per-platform
    // or per-interface label spaces, or both. The type is interface{} with range:
    // 0..2147483647.
    MplsInterfaceIndex interface{}

    // This is the minimum value of an MPLS label that this LSR is willing to
    // receive on this interface. The type is interface{} with range:
    // 0..4294967295.
    MplsInterfaceLabelMinIn interface{}

    // This is the maximum value of an MPLS label that this LSR is willing to
    // receive on this interface. The type is interface{} with range:
    // 0..4294967295.
    MplsInterfaceLabelMaxIn interface{}

    // This is the minimum value of an MPLS label that this LSR is willing to send
    // on this interface. The type is interface{} with range: 0..4294967295.
    MplsInterfaceLabelMinOut interface{}

    // This is the maximum value of an MPLS label that this LSR is willing to send
    // on this interface. The type is interface{} with range: 0..4294967295.
    MplsInterfaceLabelMaxOut interface{}

    // This value indicates the total amount of usable bandwidth on this interface
    // and is specified in kilobits per second (Kbps).  This variable is not
    // applicable when applied to the interface with index 0. When this value
    // cannot be measured, this value should contain the nominal bandwidth. The
    // type is interface{} with range: 0..4294967295. Units are kilobits per
    // second.
    MplsInterfaceTotalBandwidth interface{}

    // This value indicates the total amount of available bandwidth available on
    // this interface and is specified in kilobits per second (Kbps).  This value
    // is calculated as the difference between the amount of bandwidth currently
    // in use and that specified in mplsInterfaceTotalBandwidth.  This variable is
    // not applicable when applied to the interface with index 0. When this value
    // cannot be measured, this value should contain the nominal bandwidth. The
    // type is interface{} with range: 0..4294967295.
    MplsInterfaceAvailableBandwidth interface{}

    // If the value of the mplsInterfaceIndex for this entry is zero, then this
    // entry corresponds to the per-platform label space for all interfaces
    // configured to use that label space. In this case the perPlatform(0) bit
    // MUST be set; the perInterface(1) bit is meaningless and MUST be ignored. 
    // The remainder of this description applies to entries with a non-zero value
    // of mplsInterfaceIndex.  If the perInterface(1) bit is set then the value of
    // mplsInterfaceLabelMinIn, mplsInterfaceLabelMaxIn, mplsInterfaceLabelMinOut,
    // and mplsInterfaceLabelMaxOut for this entry reflect the label ranges for
    // this interface.  If only the perPlatform(0) bit is set, then the value of
    // mplsInterfaceLabelMinIn, mplsInterfaceLabelMaxIn, mplsInterfaceLabelMinOut,
    // and mplsInterfaceLabelMaxOut for this entry MUST be identical to the
    // instance of these objects with index 0.  These objects may only vary from
    // the entry with index 0 if both the perPlatform(0) and perInterface(1) bits
    // are set.  In all cases, at a minimum one of the perPlatform(0) or
    // perInterface(1) bits MUST be set to indicate that at least one label space
    // is in use by this interface. In all cases, agents MUST ensure that label
    // ranges are specified consistently and MUST return an inconsistentValue
    // error when they do not. The type is map[string]bool.
    MplsInterfaceLabelParticipationType interface{}

    // This object counts the number of labels that are in use at this point in
    // time on this interface in the incoming direction. If the interface
    // participates in only the per-platform label space, then the value of the
    // instance of this object MUST be identical to the value of the instance with
    // index 0. If the interface participates in the per-interface label space,
    // then the instance of this object MUST represent the number of per-interface
    // labels that are in use on this interface. The type is interface{} with
    // range: 0..4294967295.
    MplsInterfacePerfInLabelsInUse interface{}

    // This object counts the number of labeled packets that have been received on
    // this interface and which were discarded because there was no matching
    // cross- connect entry. This object MUST count on a per- interface basis
    // regardless of which label space the interface participates in. The type is
    // interface{} with range: 0..4294967295.
    MplsInterfacePerfInLabelLookupFailures interface{}

    // This object counts the number of top-most labels in the outgoing label
    // stacks that are in use at this point in time on this interface. This object
    // MUST count on a per-interface basis regardless of which label space the
    // interface participates in. The type is interface{} with range:
    // 0..4294967295.
    MplsInterfacePerfOutLabelsInUse interface{}

    // This object counts the number of outgoing MPLS packets that required
    // fragmentation before transmission on this interface. This object MUST count
    // on a per-interface basis regardless of which label space the interface
    // participates in. The type is interface{} with range: 0..4294967295.
    MplsInterfacePerfOutFragmentedPkts interface{}
}

func (mplsInterfaceEntry *MPLSLSRSTDMIB_MplsInterfaceTable_MplsInterfaceEntry) GetEntityData() *types.CommonEntityData {
    mplsInterfaceEntry.EntityData.YFilter = mplsInterfaceEntry.YFilter
    mplsInterfaceEntry.EntityData.YangName = "mplsInterfaceEntry"
    mplsInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsInterfaceEntry.EntityData.ParentYangName = "mplsInterfaceTable"
    mplsInterfaceEntry.EntityData.SegmentPath = "mplsInterfaceEntry" + types.AddKeyToken(mplsInterfaceEntry.MplsInterfaceIndex, "mplsInterfaceIndex")
    mplsInterfaceEntry.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/mplsInterfaceTable/" + mplsInterfaceEntry.EntityData.SegmentPath
    mplsInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    mplsInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceIndex", types.YLeaf{"MplsInterfaceIndex", mplsInterfaceEntry.MplsInterfaceIndex})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceLabelMinIn", types.YLeaf{"MplsInterfaceLabelMinIn", mplsInterfaceEntry.MplsInterfaceLabelMinIn})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceLabelMaxIn", types.YLeaf{"MplsInterfaceLabelMaxIn", mplsInterfaceEntry.MplsInterfaceLabelMaxIn})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceLabelMinOut", types.YLeaf{"MplsInterfaceLabelMinOut", mplsInterfaceEntry.MplsInterfaceLabelMinOut})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceLabelMaxOut", types.YLeaf{"MplsInterfaceLabelMaxOut", mplsInterfaceEntry.MplsInterfaceLabelMaxOut})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceTotalBandwidth", types.YLeaf{"MplsInterfaceTotalBandwidth", mplsInterfaceEntry.MplsInterfaceTotalBandwidth})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceAvailableBandwidth", types.YLeaf{"MplsInterfaceAvailableBandwidth", mplsInterfaceEntry.MplsInterfaceAvailableBandwidth})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfaceLabelParticipationType", types.YLeaf{"MplsInterfaceLabelParticipationType", mplsInterfaceEntry.MplsInterfaceLabelParticipationType})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfacePerfInLabelsInUse", types.YLeaf{"MplsInterfacePerfInLabelsInUse", mplsInterfaceEntry.MplsInterfacePerfInLabelsInUse})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfacePerfInLabelLookupFailures", types.YLeaf{"MplsInterfacePerfInLabelLookupFailures", mplsInterfaceEntry.MplsInterfacePerfInLabelLookupFailures})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfacePerfOutLabelsInUse", types.YLeaf{"MplsInterfacePerfOutLabelsInUse", mplsInterfaceEntry.MplsInterfacePerfOutLabelsInUse})
    mplsInterfaceEntry.EntityData.Leafs.Append("mplsInterfacePerfOutFragmentedPkts", types.YLeaf{"MplsInterfacePerfOutFragmentedPkts", mplsInterfaceEntry.MplsInterfacePerfOutFragmentedPkts})

    mplsInterfaceEntry.EntityData.YListKeys = []string {"MplsInterfaceIndex"}

    return &(mplsInterfaceEntry.EntityData)
}

// MPLSLSRSTDMIB_MplsInSegmentTable
// This table contains a description of the incoming MPLS
// segments (labels) to an LSR and their associated parameters.
// The index for this table is mplsInSegmentIndex.
// The index structure of this table is specifically designed
// to handle many different MPLS implementations that manage
// their labels both in a distributed and centralized manner.
// The table is also designed to handle existing MPLS labels
// as defined in RFC3031 as well as longer ones that may
// be necessary in the future.
// 
// In cases where the label cannot fit into the
// mplsInSegmentLabel object, the mplsInSegmentLabelPtr
// will indicate this by being set to the first accessible
// column in the appropriate extension table's row.
// In this case an additional table MUST
// be provided and MUST be indexed by at least the indexes
// used by this table. In all other cases when the label is
// represented within the mplsInSegmentLabel object, the
// mplsInSegmentLabelPtr MUST be set to 0.0. Due to the
// fact that MPLS labels may not exceed 24 bits, the
// mplsInSegmentLabelPtr object is only a provision for
// future-proofing the MIB module. Thus, the definition
// of any extension tables is beyond the scope of this
// MIB module.
type MPLSLSRSTDMIB_MplsInSegmentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents one incoming segment as is represented in
    // an LSR's LFIB. An entry can be created by a network administrator or an
    // SNMP agent, or an MPLS signaling protocol.  The creator of the entry is
    // denoted by mplsInSegmentOwner. The value of mplsInSegmentRowStatus cannot
    // be active(1) unless the ifTable entry corresponding to
    // mplsInSegmentInterface exists.  An entry in this table must match any
    // incoming packets, and indicates an instance of mplsXCEntry based on which
    // forwarding and/or switching actions are taken. The type is slice of
    // MPLSLSRSTDMIB_MplsInSegmentTable_MplsInSegmentEntry.
    MplsInSegmentEntry []*MPLSLSRSTDMIB_MplsInSegmentTable_MplsInSegmentEntry
}

func (mplsInSegmentTable *MPLSLSRSTDMIB_MplsInSegmentTable) GetEntityData() *types.CommonEntityData {
    mplsInSegmentTable.EntityData.YFilter = mplsInSegmentTable.YFilter
    mplsInSegmentTable.EntityData.YangName = "mplsInSegmentTable"
    mplsInSegmentTable.EntityData.BundleName = "cisco_ios_xe"
    mplsInSegmentTable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsInSegmentTable.EntityData.SegmentPath = "mplsInSegmentTable"
    mplsInSegmentTable.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/" + mplsInSegmentTable.EntityData.SegmentPath
    mplsInSegmentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInSegmentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInSegmentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInSegmentTable.EntityData.Children = types.NewOrderedMap()
    mplsInSegmentTable.EntityData.Children.Append("mplsInSegmentEntry", types.YChild{"MplsInSegmentEntry", nil})
    for i := range mplsInSegmentTable.MplsInSegmentEntry {
        mplsInSegmentTable.EntityData.Children.Append(types.GetSegmentPath(mplsInSegmentTable.MplsInSegmentEntry[i]), types.YChild{"MplsInSegmentEntry", mplsInSegmentTable.MplsInSegmentEntry[i]})
    }
    mplsInSegmentTable.EntityData.Leafs = types.NewOrderedMap()

    mplsInSegmentTable.EntityData.YListKeys = []string {}

    return &(mplsInSegmentTable.EntityData)
}

// MPLSLSRSTDMIB_MplsInSegmentTable_MplsInSegmentEntry
// An entry in this table represents one incoming
// segment as is represented in an LSR's LFIB.
// An entry can be created by a network
// administrator or an SNMP agent, or an MPLS signaling
// protocol.  The creator of the entry is denoted by
// mplsInSegmentOwner.
// The value of mplsInSegmentRowStatus cannot be active(1)
// unless the ifTable entry corresponding to
// mplsInSegmentInterface exists.  An entry in this table
// must match any incoming packets, and indicates an
// instance of mplsXCEntry based on which forwarding
// and/or switching actions are taken.
type MPLSLSRSTDMIB_MplsInSegmentTable_MplsInSegmentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index for this in-segment. The string
    // containing the single octet 0x00 MUST not be used as an index. The type is
    // string with length: 1..24.
    MplsInSegmentIndex interface{}

    // This object represents the interface index for the incoming MPLS interface.
    // A value of zero represents all interfaces participating in the per-platform
    // label space.  This may only be used in cases where the incoming interface
    // and label are associated with the same mplsXCEntry. Specifically, given a
    // label and any incoming interface pair from the per-platform label space,
    // the outgoing label/interface mapping remains the same. If this is not the
    // case, then individual entries MUST exist that can then be mapped to unique
    // mplsXCEntries. The type is interface{} with range: 0..2147483647.
    MplsInSegmentInterface interface{}

    // If the corresponding instance of mplsInSegmentLabelPtr is zeroDotZero then
    // this object MUST contain the incoming label associated with this
    // in-segment. If not this object SHOULD be zero and MUST be ignored. The type
    // is interface{} with range: 0..4294967295.
    MplsInSegmentLabel interface{}

    // If the label for this segment cannot be represented fully within the
    // mplsInSegmentLabel object, this object MUST point to the first accessible
    // column of a conceptual row in an external table containing the label.  In
    // this case, the mplsInSegmentTopLabel object SHOULD be set to 0 and ignored.
    // This object MUST be set to zeroDotZero otherwise. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    MplsInSegmentLabelPtr interface{}

    // The number of labels to pop from the incoming packet.  Normally only the
    // top label is popped from the packet and used for all switching decisions
    // for that packet.  This is indicated by setting this object to the default
    // value of 1. If an LSR supports popping of more than one label, this object
    // MUST be set to that number. This object cannot be modified if
    // mplsInSegmentRowStatus is active(1). The type is interface{} with range:
    // 1..2147483647.
    MplsInSegmentNPop interface{}

    // The IANA address family [IANAFamily] of packets received on this segment,
    // which is used at an egress LSR to deliver them to the appropriate layer 3
    // entity. A value of other(0) indicates that the family type is either
    // unknown or undefined; this SHOULD NOT be used at an egress LSR. This object
    // cannot be modified if mplsInSegmentRowStatus is active(1). The type is
    // AddressFamilyNumbers.
    MplsInSegmentAddrFamily interface{}

    // Index into mplsXCTable which identifies which cross- connect entry this
    // segment is part of.  The string containing the single octet 0x00 indicates
    // that this entry is not referred to by any cross-connect entry. When a
    // cross-connect entry is created which this in-segment is a part of, this
    // object is automatically updated to reflect the value of mplsXCIndex of that
    // cross-connect entry. The type is string with length: 1..24.
    MplsInSegmentXCIndex interface{}

    // Denotes the entity that created and is responsible for managing this
    // segment. The type is MplsOwner.
    MplsInSegmentOwner interface{}

    // This variable represents a pointer to the traffic parameter specification
    // for this in-segment.  This value may point at an entry in the
    // mplsTunnelResourceTable in the MPLS-TE-STD-MIB (RFC3812) to indicate which
    // traffic parameter settings for this segment if it represents an LSP used
    // for a TE tunnel.  This value may optionally point at an externally defined
    // traffic parameter specification table.  A value of zeroDotZero indicates
    // best-effort treatment.  By having the same value of this object, two or
    // more segments can indicate resource sharing of such things as LSP queue
    // space, etc.  This object cannot be modified if mplsInSegmentRowStatus is
    // active(1).  For entries in this table that are preserved after a re-boot,
    // the agent MUST ensure that their integrity be preserved, or this object
    // should be set to 0.0 if it cannot. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    MplsInSegmentTrafficParamPtr interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table has a row in the active(1) state, no objects in
    // this row can be modified except the mplsInSegmentRowStatus and
    // mplsInSegmentStorageType. The type is RowStatus.
    MplsInSegmentRowStatus interface{}

    // This variable indicates the storage type for this object. The agent MUST
    // ensure that this object's value remains consistent with the associated
    // mplsXCEntry. Conceptual rows having the value 'permanent' need not allow
    // write-access to any columnar objects in the row. The type is StorageType.
    MplsInSegmentStorageType interface{}

    // This value represents the total number of octets received by this segment.
    // It MUST be equal to the least significant 32 bits of
    // mplsInSegmentPerfHCOctets if mplsInSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..4294967295.
    MplsInSegmentPerfOctets interface{}

    // Total number of packets received by this segment. The type is interface{}
    // with range: 0..4294967295.
    MplsInSegmentPerfPackets interface{}

    // The number of errored packets received on this segment. The type is
    // interface{} with range: 0..4294967295.
    MplsInSegmentPerfErrors interface{}

    // The number of labeled packets received on this in- segment, which were
    // chosen to be discarded even though no errors had been detected to prevent
    // their being transmitted.  One possible reason for discarding such a labeled
    // packet could be to free up buffer space. The type is interface{} with
    // range: 0..4294967295.
    MplsInSegmentPerfDiscards interface{}

    // The total number of octets received.  This is the 64 bit version of
    // mplsInSegmentPerfOctets, if mplsInSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..18446744073709551615.
    MplsInSegmentPerfHCOctets interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this segment's Counter32 or Counter64 suffered a discontinuity. If no
    // such discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    MplsInSegmentPerfDiscontinuityTime interface{}
}

func (mplsInSegmentEntry *MPLSLSRSTDMIB_MplsInSegmentTable_MplsInSegmentEntry) GetEntityData() *types.CommonEntityData {
    mplsInSegmentEntry.EntityData.YFilter = mplsInSegmentEntry.YFilter
    mplsInSegmentEntry.EntityData.YangName = "mplsInSegmentEntry"
    mplsInSegmentEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsInSegmentEntry.EntityData.ParentYangName = "mplsInSegmentTable"
    mplsInSegmentEntry.EntityData.SegmentPath = "mplsInSegmentEntry" + types.AddKeyToken(mplsInSegmentEntry.MplsInSegmentIndex, "mplsInSegmentIndex")
    mplsInSegmentEntry.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/mplsInSegmentTable/" + mplsInSegmentEntry.EntityData.SegmentPath
    mplsInSegmentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInSegmentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInSegmentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInSegmentEntry.EntityData.Children = types.NewOrderedMap()
    mplsInSegmentEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentIndex", types.YLeaf{"MplsInSegmentIndex", mplsInSegmentEntry.MplsInSegmentIndex})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentInterface", types.YLeaf{"MplsInSegmentInterface", mplsInSegmentEntry.MplsInSegmentInterface})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentLabel", types.YLeaf{"MplsInSegmentLabel", mplsInSegmentEntry.MplsInSegmentLabel})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentLabelPtr", types.YLeaf{"MplsInSegmentLabelPtr", mplsInSegmentEntry.MplsInSegmentLabelPtr})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentNPop", types.YLeaf{"MplsInSegmentNPop", mplsInSegmentEntry.MplsInSegmentNPop})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentAddrFamily", types.YLeaf{"MplsInSegmentAddrFamily", mplsInSegmentEntry.MplsInSegmentAddrFamily})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentXCIndex", types.YLeaf{"MplsInSegmentXCIndex", mplsInSegmentEntry.MplsInSegmentXCIndex})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentOwner", types.YLeaf{"MplsInSegmentOwner", mplsInSegmentEntry.MplsInSegmentOwner})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentTrafficParamPtr", types.YLeaf{"MplsInSegmentTrafficParamPtr", mplsInSegmentEntry.MplsInSegmentTrafficParamPtr})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentRowStatus", types.YLeaf{"MplsInSegmentRowStatus", mplsInSegmentEntry.MplsInSegmentRowStatus})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentStorageType", types.YLeaf{"MplsInSegmentStorageType", mplsInSegmentEntry.MplsInSegmentStorageType})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentPerfOctets", types.YLeaf{"MplsInSegmentPerfOctets", mplsInSegmentEntry.MplsInSegmentPerfOctets})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentPerfPackets", types.YLeaf{"MplsInSegmentPerfPackets", mplsInSegmentEntry.MplsInSegmentPerfPackets})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentPerfErrors", types.YLeaf{"MplsInSegmentPerfErrors", mplsInSegmentEntry.MplsInSegmentPerfErrors})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentPerfDiscards", types.YLeaf{"MplsInSegmentPerfDiscards", mplsInSegmentEntry.MplsInSegmentPerfDiscards})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentPerfHCOctets", types.YLeaf{"MplsInSegmentPerfHCOctets", mplsInSegmentEntry.MplsInSegmentPerfHCOctets})
    mplsInSegmentEntry.EntityData.Leafs.Append("mplsInSegmentPerfDiscontinuityTime", types.YLeaf{"MplsInSegmentPerfDiscontinuityTime", mplsInSegmentEntry.MplsInSegmentPerfDiscontinuityTime})

    mplsInSegmentEntry.EntityData.YListKeys = []string {"MplsInSegmentIndex"}

    return &(mplsInSegmentEntry.EntityData)
}

// MPLSLSRSTDMIB_MplsOutSegmentTable
// This table contains a representation of the outgoing
// segments from an LSR.
type MPLSLSRSTDMIB_MplsOutSegmentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents one outgoing segment.  An entry can be
    // created by a network administrator, an SNMP agent, or an MPLS signaling
    // protocol.  The object mplsOutSegmentOwner indicates the creator of this
    // entry. The value of mplsOutSegmentRowStatus cannot be active(1) unless the
    // ifTable entry corresponding to mplsOutSegmentInterface exists.  Note that
    // the indexing of this table uses a single, arbitrary index
    // (mplsOutSegmentIndex) to indicate which out-segment (i.e.: label) is being
    // switched to from which in-segment (i.e: label) or in-segments. This is
    // necessary because it is possible to have an equal-cost multi-path situation
    // where two identical out-going labels are assigned to the same cross-connect
    // (i.e.: they go to two different neighboring LSRs); thus, requiring two
    // out-segments. In order to preserve the uniqueness of the references by the
    // mplsXCEntry, an arbitrary integer must be used as the index for this table.
    // The type is slice of MPLSLSRSTDMIB_MplsOutSegmentTable_MplsOutSegmentEntry.
    MplsOutSegmentEntry []*MPLSLSRSTDMIB_MplsOutSegmentTable_MplsOutSegmentEntry
}

func (mplsOutSegmentTable *MPLSLSRSTDMIB_MplsOutSegmentTable) GetEntityData() *types.CommonEntityData {
    mplsOutSegmentTable.EntityData.YFilter = mplsOutSegmentTable.YFilter
    mplsOutSegmentTable.EntityData.YangName = "mplsOutSegmentTable"
    mplsOutSegmentTable.EntityData.BundleName = "cisco_ios_xe"
    mplsOutSegmentTable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsOutSegmentTable.EntityData.SegmentPath = "mplsOutSegmentTable"
    mplsOutSegmentTable.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/" + mplsOutSegmentTable.EntityData.SegmentPath
    mplsOutSegmentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsOutSegmentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsOutSegmentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsOutSegmentTable.EntityData.Children = types.NewOrderedMap()
    mplsOutSegmentTable.EntityData.Children.Append("mplsOutSegmentEntry", types.YChild{"MplsOutSegmentEntry", nil})
    for i := range mplsOutSegmentTable.MplsOutSegmentEntry {
        mplsOutSegmentTable.EntityData.Children.Append(types.GetSegmentPath(mplsOutSegmentTable.MplsOutSegmentEntry[i]), types.YChild{"MplsOutSegmentEntry", mplsOutSegmentTable.MplsOutSegmentEntry[i]})
    }
    mplsOutSegmentTable.EntityData.Leafs = types.NewOrderedMap()

    mplsOutSegmentTable.EntityData.YListKeys = []string {}

    return &(mplsOutSegmentTable.EntityData)
}

// MPLSLSRSTDMIB_MplsOutSegmentTable_MplsOutSegmentEntry
// An entry in this table represents one outgoing
// segment.  An entry can be created by a network
// administrator, an SNMP agent, or an MPLS signaling
// protocol.  The object mplsOutSegmentOwner indicates
// the creator of this entry. The value of
// mplsOutSegmentRowStatus cannot be active(1) unless
// the ifTable entry corresponding to
// mplsOutSegmentInterface exists.
// 
// Note that the indexing of this table uses a single,
// arbitrary index (mplsOutSegmentIndex) to indicate
// which out-segment (i.e.: label) is being switched to
// from which in-segment (i.e: label) or in-segments.
// This is necessary because it is possible to have an
// equal-cost multi-path situation where two identical
// out-going labels are assigned to the same
// cross-connect (i.e.: they go to two different neighboring
// LSRs); thus, requiring two out-segments. In order to
// preserve the uniqueness of the references
// by the mplsXCEntry, an arbitrary integer must be used as
// the index for this table.
type MPLSLSRSTDMIB_MplsOutSegmentTable_MplsOutSegmentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This value contains a unique index for this row.
    // While a value of a string containing the single octet 0x00 is not valid as
    // an index for entries in this table, it can be supplied as a valid value to
    // index the mplsXCTable to represent entries for which no out-segment has
    // been configured or exists. The type is string with length: 1..24.
    MplsOutSegmentIndex interface{}

    // This value must contain the interface index of the outgoing interface. This
    // object cannot be modified if mplsOutSegmentRowStatus is active(1). The
    // mplsOutSegmentRowStatus cannot be set to active(1) until this object is set
    // to a value corresponding to a valid ifEntry. The type is interface{} with
    // range: 0..2147483647.
    MplsOutSegmentInterface interface{}

    // This value indicates whether or not a top label should be pushed onto the
    // outgoing packet's label stack.  The value of this variable MUST be set to
    // true(1) if the outgoing interface does not support pop-and-go (and no label
    // stack remains). For example, on ATM interface, or if the segment represents
    // a tunnel origination.  Note that it is considered an error in the case that
    // mplsOutSegmentPushTopLabel is set to false, but the cross-connect entry
    // which refers to this out-segment has a non-zero mplsLabelStackIndex.  The
    // LSR MUST ensure that this situation does not happen. This object cannot be
    // modified if mplsOutSegmentRowStatus is active(1). The type is bool.
    MplsOutSegmentPushTopLabel interface{}

    // If mplsOutSegmentPushTopLabel is true then this represents the label that
    // should be pushed onto the top of the outgoing packet's label stack.
    // Otherwise this value SHOULD be set to 0 by the management station and MUST
    // be ignored by the agent. This object cannot be modified if
    // mplsOutSegmentRowStatus is active(1). The type is interface{} with range:
    // 0..4294967295.
    MplsOutSegmentTopLabel interface{}

    // If the label for this segment cannot be represented fully within the
    // mplsOutSegmentLabel object, this object MUST point to the first accessible
    // column of a conceptual row in an external table containing the label.  In
    // this case, the mplsOutSegmentTopLabel object SHOULD be set to 0 and
    // ignored. This object MUST be set to zeroDotZero otherwise. The type is
    // string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    MplsOutSegmentTopLabelPtr interface{}

    // Indicates the next hop Internet address type. Only values unknown(0),
    // ipv4(1) or ipv6(2) have to be supported.  A value of unknown(0) is allowed
    // only when the outgoing interface is of type point-to-point. If any other
    // unsupported values are attempted in a set operation, the agent MUST return
    // an inconsistentValue error. The type is InetAddressType.
    MplsOutSegmentNextHopAddrType interface{}

    // The internet address of the next hop. The type of this address is
    // determined by the value of the mplslOutSegmentNextHopAddrType object.  This
    // object cannot be modified if mplsOutSegmentRowStatus is active(1). The type
    // is string with length: 0..255.
    MplsOutSegmentNextHopAddr interface{}

    // Index into mplsXCTable which identifies which cross- connect entry this
    // segment is part of.  A value of the string containing the single octet 0x00
    // indicates that this entry is not referred to by any cross-connect entry. 
    // When a cross-connect entry is created which this out-segment is a part of,
    // this object MUST be updated by the agent to reflect the value of
    // mplsXCIndex of that cross-connect entry. The type is string with length:
    // 1..24.
    MplsOutSegmentXCIndex interface{}

    // Denotes the entity which created and is responsible for managing this
    // segment. The type is MplsOwner.
    MplsOutSegmentOwner interface{}

    // This variable represents a pointer to the traffic parameter specification
    // for this out-segment.  This value may point at an entry in the
    // MplsTunnelResourceEntry in the MPLS-TE-STD-MIB (RFC3812)  RFC Editor:
    // Please fill in RFC number.  to indicate which traffic parameter settings
    // for this segment if it represents an LSP used for a TE tunnel.  This value
    // may optionally point at an externally defined traffic parameter
    // specification table.  A value of zeroDotZero indicates best-effort
    // treatment.  By having the same value of this object, two or more segments
    // can indicate resource sharing of such things as LSP queue space, etc.  This
    // object cannot be modified if mplsOutSegmentRowStatus is active(1). For
    // entries in this table that are preserved after a re-boot, the agent MUST
    // ensure that their integrity be preserved, or this object should be set to
    // 0.0 if it cannot. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    MplsOutSegmentTrafficParamPtr interface{}

    // For creating, modifying, and deleting this row. When a row in this table
    // has a row in the active(1) state, no objects in this row can be modified
    // except the mplsOutSegmentRowStatus or mplsOutSegmentStorageType. The type
    // is RowStatus.
    MplsOutSegmentRowStatus interface{}

    // This variable indicates the storage type for this object. The agent MUST
    // ensure that this object's value remains consistent with the associated
    // mplsXCEntry. Conceptual rows having the value 'permanent' need not allow
    // write-access to any columnar objects in the row. The type is StorageType.
    MplsOutSegmentStorageType interface{}

    // This value contains the total number of octets sent on this segment. It
    // MUST be equal to the least significant 32 bits of
    // mplsOutSegmentPerfHCOctets if mplsOutSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..4294967295.
    MplsOutSegmentPerfOctets interface{}

    // This value contains the total number of packets sent on this segment. The
    // type is interface{} with range: 0..4294967295.
    MplsOutSegmentPerfPackets interface{}

    // Number of packets that could not be sent due to errors on this segment. The
    // type is interface{} with range: 0..4294967295.
    MplsOutSegmentPerfErrors interface{}

    // The number of labeled packets attempted to be transmitted on this
    // out-segment, which were chosen to be discarded even though no errors had
    // been detected to prevent their being transmitted. One possible reason for
    // discarding such a labeled packet could be to free up buffer space. The type
    // is interface{} with range: 0..4294967295.
    MplsOutSegmentPerfDiscards interface{}

    // Total number of octets sent.  This is the 64 bit version of
    // mplsOutSegmentPerfOctets, if mplsOutSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..18446744073709551615.
    MplsOutSegmentPerfHCOctets interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this segment's Counter32 or Counter64 suffered a discontinuity. If no
    // such discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    MplsOutSegmentPerfDiscontinuityTime interface{}
}

func (mplsOutSegmentEntry *MPLSLSRSTDMIB_MplsOutSegmentTable_MplsOutSegmentEntry) GetEntityData() *types.CommonEntityData {
    mplsOutSegmentEntry.EntityData.YFilter = mplsOutSegmentEntry.YFilter
    mplsOutSegmentEntry.EntityData.YangName = "mplsOutSegmentEntry"
    mplsOutSegmentEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsOutSegmentEntry.EntityData.ParentYangName = "mplsOutSegmentTable"
    mplsOutSegmentEntry.EntityData.SegmentPath = "mplsOutSegmentEntry" + types.AddKeyToken(mplsOutSegmentEntry.MplsOutSegmentIndex, "mplsOutSegmentIndex")
    mplsOutSegmentEntry.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/mplsOutSegmentTable/" + mplsOutSegmentEntry.EntityData.SegmentPath
    mplsOutSegmentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsOutSegmentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsOutSegmentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsOutSegmentEntry.EntityData.Children = types.NewOrderedMap()
    mplsOutSegmentEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentIndex", types.YLeaf{"MplsOutSegmentIndex", mplsOutSegmentEntry.MplsOutSegmentIndex})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentInterface", types.YLeaf{"MplsOutSegmentInterface", mplsOutSegmentEntry.MplsOutSegmentInterface})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentPushTopLabel", types.YLeaf{"MplsOutSegmentPushTopLabel", mplsOutSegmentEntry.MplsOutSegmentPushTopLabel})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentTopLabel", types.YLeaf{"MplsOutSegmentTopLabel", mplsOutSegmentEntry.MplsOutSegmentTopLabel})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentTopLabelPtr", types.YLeaf{"MplsOutSegmentTopLabelPtr", mplsOutSegmentEntry.MplsOutSegmentTopLabelPtr})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentNextHopAddrType", types.YLeaf{"MplsOutSegmentNextHopAddrType", mplsOutSegmentEntry.MplsOutSegmentNextHopAddrType})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentNextHopAddr", types.YLeaf{"MplsOutSegmentNextHopAddr", mplsOutSegmentEntry.MplsOutSegmentNextHopAddr})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentXCIndex", types.YLeaf{"MplsOutSegmentXCIndex", mplsOutSegmentEntry.MplsOutSegmentXCIndex})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentOwner", types.YLeaf{"MplsOutSegmentOwner", mplsOutSegmentEntry.MplsOutSegmentOwner})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentTrafficParamPtr", types.YLeaf{"MplsOutSegmentTrafficParamPtr", mplsOutSegmentEntry.MplsOutSegmentTrafficParamPtr})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentRowStatus", types.YLeaf{"MplsOutSegmentRowStatus", mplsOutSegmentEntry.MplsOutSegmentRowStatus})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentStorageType", types.YLeaf{"MplsOutSegmentStorageType", mplsOutSegmentEntry.MplsOutSegmentStorageType})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentPerfOctets", types.YLeaf{"MplsOutSegmentPerfOctets", mplsOutSegmentEntry.MplsOutSegmentPerfOctets})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentPerfPackets", types.YLeaf{"MplsOutSegmentPerfPackets", mplsOutSegmentEntry.MplsOutSegmentPerfPackets})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentPerfErrors", types.YLeaf{"MplsOutSegmentPerfErrors", mplsOutSegmentEntry.MplsOutSegmentPerfErrors})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentPerfDiscards", types.YLeaf{"MplsOutSegmentPerfDiscards", mplsOutSegmentEntry.MplsOutSegmentPerfDiscards})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentPerfHCOctets", types.YLeaf{"MplsOutSegmentPerfHCOctets", mplsOutSegmentEntry.MplsOutSegmentPerfHCOctets})
    mplsOutSegmentEntry.EntityData.Leafs.Append("mplsOutSegmentPerfDiscontinuityTime", types.YLeaf{"MplsOutSegmentPerfDiscontinuityTime", mplsOutSegmentEntry.MplsOutSegmentPerfDiscontinuityTime})

    mplsOutSegmentEntry.EntityData.YListKeys = []string {"MplsOutSegmentIndex"}

    return &(mplsOutSegmentEntry.EntityData)
}

// MPLSLSRSTDMIB_MplsXCTable
// This table specifies information for switching
// between LSP segments.  It supports point-to-point,
// point-to-multipoint and multipoint-to-point
// connections.  mplsLabelStackTable specifies the
// label stack information for a cross-connect LSR and
// is referred to from mplsXCTable.
type MPLSLSRSTDMIB_MplsXCTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents one cross-connect entry.  It is indexed by
    // the following objects:  - cross-connect index mplsXCIndex that uniquely  
    // identifies a group of cross-connect entries  - in-segment index,
    // mplsXCInSegmentIndex  - out-segment index, mplsXCOutSegmentIndex  LSPs
    // originating at this LSR: These are represented by using the special of
    // value of mplsXCInSegmentIndex set to the string containing a single octet
    // 0x00. In this case the mplsXCOutSegmentIndex MUST not be the string
    // containing a single octet 0x00.  LSPs terminating at this LSR: These are
    // represented by using the special value mplsXCOutSegmentIndex set to the
    // string containing a single octet 0x00.  Special labels: Entries indexed by
    // the strings containing the reserved MPLS label values as a single octet
    // 0x00 through 0x0f (inclusive) imply LSPs terminating at this LSR.  Note
    // that situations where LSPs are terminated with incoming label equal to the
    // string containing a single octet 0x00 can be distinguished from LSPs
    // originating at this LSR because the mplsXCOutSegmentIndex equals the string
    // containing the single octet 0x00.  An entry can be created by a network
    // administrator or by an SNMP agent as instructed by an MPLS signaling
    // protocol. The type is slice of MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry.
    MplsXCEntry []*MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry
}

func (mplsXCTable *MPLSLSRSTDMIB_MplsXCTable) GetEntityData() *types.CommonEntityData {
    mplsXCTable.EntityData.YFilter = mplsXCTable.YFilter
    mplsXCTable.EntityData.YangName = "mplsXCTable"
    mplsXCTable.EntityData.BundleName = "cisco_ios_xe"
    mplsXCTable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsXCTable.EntityData.SegmentPath = "mplsXCTable"
    mplsXCTable.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/" + mplsXCTable.EntityData.SegmentPath
    mplsXCTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsXCTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsXCTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsXCTable.EntityData.Children = types.NewOrderedMap()
    mplsXCTable.EntityData.Children.Append("mplsXCEntry", types.YChild{"MplsXCEntry", nil})
    for i := range mplsXCTable.MplsXCEntry {
        mplsXCTable.EntityData.Children.Append(types.GetSegmentPath(mplsXCTable.MplsXCEntry[i]), types.YChild{"MplsXCEntry", mplsXCTable.MplsXCEntry[i]})
    }
    mplsXCTable.EntityData.Leafs = types.NewOrderedMap()

    mplsXCTable.EntityData.YListKeys = []string {}

    return &(mplsXCTable.EntityData)
}

// MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry
// A row in this table represents one cross-connect
// entry.  It is indexed by the following objects:
// 
// - cross-connect index mplsXCIndex that uniquely
//   identifies a group of cross-connect entries
// 
// - in-segment index, mplsXCInSegmentIndex
// 
// - out-segment index, mplsXCOutSegmentIndex
// 
// LSPs originating at this LSR:
// These are represented by using the special
// of value of mplsXCInSegmentIndex set to the
// string containing a single octet 0x00. In
// this case the mplsXCOutSegmentIndex
// MUST not be the string containing a single
// octet 0x00.
// 
// LSPs terminating at this LSR:
// These are represented by using the special value
// mplsXCOutSegmentIndex set to the string containing
// a single octet 0x00.
// 
// Special labels:
// Entries indexed by the strings containing the
// reserved MPLS label values as a single octet 0x00
// through 0x0f (inclusive) imply LSPs terminating at
// this LSR.  Note that situations where LSPs are
// terminated with incoming label equal to the string
// containing a single octet 0x00 can be distinguished
// from LSPs originating at this LSR because the
// mplsXCOutSegmentIndex equals the string containing the
// single octet 0x00.
// 
// An entry can be created by a network administrator
// or by an SNMP agent as instructed by an MPLS
// signaling protocol.
type MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Primary index for the conceptual row identifying a
    // group of cross-connect segments. The string containing a single octet 0x00
    // is an invalid index. The type is string with length: 1..24.
    MplsXCIndex interface{}

    // This attribute is a key. Incoming label index. If this object is set to the
    // string containing a single octet 0x00, this indicates a special case
    // outlined in the table's description above. In this case no corresponding
    // mplsInSegmentEntry shall exist. The type is string with length: 1..24.
    MplsXCInSegmentIndex interface{}

    // This attribute is a key. Index of out-segment for LSPs not terminating on
    // this LSR if not set to the string containing the single octet 0x00. If the
    // segment identified by this entry is terminating, then this object MUST be
    // set to the string containing a single octet 0x00 to indicate that no
    // corresponding mplsOutSegmentEntry shall exist. The type is string with
    // length: 1..24.
    MplsXCOutSegmentIndex interface{}

    // This value identifies the label switched path that this cross-connect entry
    // belongs to. This object cannot be modified if mplsXCRowStatus is active(1)
    // except for this object. The type is string with length: 2..2 | 6..6.
    MplsXCLspId interface{}

    // Primary index into mplsLabelStackTable identifying a stack of labels to be
    // pushed beneath the top label. Note that the top label identified by the
    // out- segment ensures that all the components of a multipoint-to-point
    // connection have the same outgoing label. A value of the string containing
    // the single octet 0x00 indicates that no labels are to be stacked beneath
    // the top label. This object cannot be modified if mplsXCRowStatus is
    // active(1). The type is string with length: 1..24.
    MplsXCLabelStackIndex interface{}

    // Denotes the entity that created and is responsible for managing this
    // cross-connect. The type is MplsOwner.
    MplsXCOwner interface{}

    // For creating, modifying, and deleting this row. When a row in this table
    // has a row in the active(1) state, no objects in this row except this object
    // and the mplsXCStorageType can be modified. . The type is RowStatus.
    MplsXCRowStatus interface{}

    // This variable indicates the storage type for this object. The agent MUST
    // ensure that the associated in and out segments also have the same
    // StorageType value and are restored consistently upon system restart. This
    // value SHOULD be set to permanent(4) if created as a result of a static LSP
    // configuration.  Conceptual rows having the value 'permanent' need not allow
    // write-access to any columnar objects in the row. The type is StorageType.
    MplsXCStorageType interface{}

    // The desired operational status of this segment. The type is
    // MplsXCAdminStatus.
    MplsXCAdminStatus interface{}

    // The actual operational status of this cross- connect. The type is
    // MplsXCOperStatus.
    MplsXCOperStatus interface{}
}

func (mplsXCEntry *MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry) GetEntityData() *types.CommonEntityData {
    mplsXCEntry.EntityData.YFilter = mplsXCEntry.YFilter
    mplsXCEntry.EntityData.YangName = "mplsXCEntry"
    mplsXCEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsXCEntry.EntityData.ParentYangName = "mplsXCTable"
    mplsXCEntry.EntityData.SegmentPath = "mplsXCEntry" + types.AddKeyToken(mplsXCEntry.MplsXCIndex, "mplsXCIndex") + types.AddKeyToken(mplsXCEntry.MplsXCInSegmentIndex, "mplsXCInSegmentIndex") + types.AddKeyToken(mplsXCEntry.MplsXCOutSegmentIndex, "mplsXCOutSegmentIndex")
    mplsXCEntry.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/mplsXCTable/" + mplsXCEntry.EntityData.SegmentPath
    mplsXCEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsXCEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsXCEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsXCEntry.EntityData.Children = types.NewOrderedMap()
    mplsXCEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsXCEntry.EntityData.Leafs.Append("mplsXCIndex", types.YLeaf{"MplsXCIndex", mplsXCEntry.MplsXCIndex})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCInSegmentIndex", types.YLeaf{"MplsXCInSegmentIndex", mplsXCEntry.MplsXCInSegmentIndex})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCOutSegmentIndex", types.YLeaf{"MplsXCOutSegmentIndex", mplsXCEntry.MplsXCOutSegmentIndex})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCLspId", types.YLeaf{"MplsXCLspId", mplsXCEntry.MplsXCLspId})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCLabelStackIndex", types.YLeaf{"MplsXCLabelStackIndex", mplsXCEntry.MplsXCLabelStackIndex})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCOwner", types.YLeaf{"MplsXCOwner", mplsXCEntry.MplsXCOwner})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCRowStatus", types.YLeaf{"MplsXCRowStatus", mplsXCEntry.MplsXCRowStatus})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCStorageType", types.YLeaf{"MplsXCStorageType", mplsXCEntry.MplsXCStorageType})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCAdminStatus", types.YLeaf{"MplsXCAdminStatus", mplsXCEntry.MplsXCAdminStatus})
    mplsXCEntry.EntityData.Leafs.Append("mplsXCOperStatus", types.YLeaf{"MplsXCOperStatus", mplsXCEntry.MplsXCOperStatus})

    mplsXCEntry.EntityData.YListKeys = []string {"MplsXCIndex", "MplsXCInSegmentIndex", "MplsXCOutSegmentIndex"}

    return &(mplsXCEntry.EntityData)
}

// MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus represents The desired operational status of this segment.
type MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus string

const (
    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus_up MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus = "up"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus_down MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus = "down"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus_testing MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCAdminStatus = "testing"
)

// MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus represents connect.
type MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus string

const (
    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus_up MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus = "up"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus_down MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus = "down"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus_testing MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus = "testing"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus_unknown MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus = "unknown"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus_dormant MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus = "dormant"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus_notPresent MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus = "notPresent"

    MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus_lowerLayerDown MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOperStatus = "lowerLayerDown"
)

// MPLSLSRSTDMIB_MplsLabelStackTable
// This table specifies the label stack to be pushed
// onto a packet, beneath the top label.  Entries into
// this table are referred to from mplsXCTable.
type MPLSLSRSTDMIB_MplsLabelStackTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents one label which is to be pushed onto an
    // outgoing packet, beneath the top label.  An entry can be created by a
    // network administrator or by an SNMP agent as instructed by an MPLS
    // signaling protocol. The type is slice of
    // MPLSLSRSTDMIB_MplsLabelStackTable_MplsLabelStackEntry.
    MplsLabelStackEntry []*MPLSLSRSTDMIB_MplsLabelStackTable_MplsLabelStackEntry
}

func (mplsLabelStackTable *MPLSLSRSTDMIB_MplsLabelStackTable) GetEntityData() *types.CommonEntityData {
    mplsLabelStackTable.EntityData.YFilter = mplsLabelStackTable.YFilter
    mplsLabelStackTable.EntityData.YangName = "mplsLabelStackTable"
    mplsLabelStackTable.EntityData.BundleName = "cisco_ios_xe"
    mplsLabelStackTable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsLabelStackTable.EntityData.SegmentPath = "mplsLabelStackTable"
    mplsLabelStackTable.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/" + mplsLabelStackTable.EntityData.SegmentPath
    mplsLabelStackTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLabelStackTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLabelStackTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLabelStackTable.EntityData.Children = types.NewOrderedMap()
    mplsLabelStackTable.EntityData.Children.Append("mplsLabelStackEntry", types.YChild{"MplsLabelStackEntry", nil})
    for i := range mplsLabelStackTable.MplsLabelStackEntry {
        mplsLabelStackTable.EntityData.Children.Append(types.GetSegmentPath(mplsLabelStackTable.MplsLabelStackEntry[i]), types.YChild{"MplsLabelStackEntry", mplsLabelStackTable.MplsLabelStackEntry[i]})
    }
    mplsLabelStackTable.EntityData.Leafs = types.NewOrderedMap()

    mplsLabelStackTable.EntityData.YListKeys = []string {}

    return &(mplsLabelStackTable.EntityData)
}

// MPLSLSRSTDMIB_MplsLabelStackTable_MplsLabelStackEntry
// An entry in this table represents one label which is
// to be pushed onto an outgoing packet, beneath the
// top label.  An entry can be created by a network
// administrator or by an SNMP agent as instructed by
// an MPLS signaling protocol.
type MPLSLSRSTDMIB_MplsLabelStackTable_MplsLabelStackEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Primary index for this row identifying a stack of
    // labels to be pushed on an outgoing packet, beneath the top label. An index
    // containing the string with a single octet 0x00 MUST not be used. The type
    // is string with length: 1..24.
    MplsLabelStackIndex interface{}

    // This attribute is a key. Secondary index for this row identifying one label
    // of the stack.  Note that an entry with a smaller mplsLabelStackLabelIndex
    // would refer to a label higher up the label stack and would be popped at a
    // downstream LSR before a label represented by a higher
    // mplsLabelStackLabelIndex at a downstream LSR. The type is interface{} with
    // range: 1..2147483647.
    MplsLabelStackLabelIndex interface{}

    // The label to pushed. The type is interface{} with range: 0..4294967295.
    MplsLabelStackLabel interface{}

    // If the label for this segment cannot be represented fully within the
    // mplsLabelStackLabel object, this object MUST point to the first accessible
    // column of a conceptual row in an external table containing the label.  In
    // this case, the mplsLabelStackLabel object SHOULD be set to 0 and ignored.
    // This object MUST be set to zeroDotZero otherwise. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    MplsLabelStackLabelPtr interface{}

    // For creating, modifying, and deleting this row. When a row in this table
    // has a row in the active(1) state, no objects in this row except this object
    // and the mplsLabelStackStorageType can be modified. The type is RowStatus.
    MplsLabelStackRowStatus interface{}

    // This variable indicates the storage type for this object. This object
    // cannot be modified if mplsLabelStackRowStatus is active(1). No objects are
    // required to be writable for rows in this table with this object set to
    // permanent(4).  The agent MUST ensure that all related entries in this table
    // retain the same value for this object.  Agents MUST ensure that the storage
    // type for all entries related to a particular mplsXCEntry retain the same
    // value for this object as the mplsXCEntry's StorageType. The type is
    // StorageType.
    MplsLabelStackStorageType interface{}
}

func (mplsLabelStackEntry *MPLSLSRSTDMIB_MplsLabelStackTable_MplsLabelStackEntry) GetEntityData() *types.CommonEntityData {
    mplsLabelStackEntry.EntityData.YFilter = mplsLabelStackEntry.YFilter
    mplsLabelStackEntry.EntityData.YangName = "mplsLabelStackEntry"
    mplsLabelStackEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsLabelStackEntry.EntityData.ParentYangName = "mplsLabelStackTable"
    mplsLabelStackEntry.EntityData.SegmentPath = "mplsLabelStackEntry" + types.AddKeyToken(mplsLabelStackEntry.MplsLabelStackIndex, "mplsLabelStackIndex") + types.AddKeyToken(mplsLabelStackEntry.MplsLabelStackLabelIndex, "mplsLabelStackLabelIndex")
    mplsLabelStackEntry.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/mplsLabelStackTable/" + mplsLabelStackEntry.EntityData.SegmentPath
    mplsLabelStackEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLabelStackEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLabelStackEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLabelStackEntry.EntityData.Children = types.NewOrderedMap()
    mplsLabelStackEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsLabelStackEntry.EntityData.Leafs.Append("mplsLabelStackIndex", types.YLeaf{"MplsLabelStackIndex", mplsLabelStackEntry.MplsLabelStackIndex})
    mplsLabelStackEntry.EntityData.Leafs.Append("mplsLabelStackLabelIndex", types.YLeaf{"MplsLabelStackLabelIndex", mplsLabelStackEntry.MplsLabelStackLabelIndex})
    mplsLabelStackEntry.EntityData.Leafs.Append("mplsLabelStackLabel", types.YLeaf{"MplsLabelStackLabel", mplsLabelStackEntry.MplsLabelStackLabel})
    mplsLabelStackEntry.EntityData.Leafs.Append("mplsLabelStackLabelPtr", types.YLeaf{"MplsLabelStackLabelPtr", mplsLabelStackEntry.MplsLabelStackLabelPtr})
    mplsLabelStackEntry.EntityData.Leafs.Append("mplsLabelStackRowStatus", types.YLeaf{"MplsLabelStackRowStatus", mplsLabelStackEntry.MplsLabelStackRowStatus})
    mplsLabelStackEntry.EntityData.Leafs.Append("mplsLabelStackStorageType", types.YLeaf{"MplsLabelStackStorageType", mplsLabelStackEntry.MplsLabelStackStorageType})

    mplsLabelStackEntry.EntityData.YListKeys = []string {"MplsLabelStackIndex", "MplsLabelStackLabelIndex"}

    return &(mplsLabelStackEntry.EntityData)
}

// MPLSLSRSTDMIB_MplsInSegmentMapTable
// This table specifies the mapping from the
// mplsInSegmentIndex to the corresponding
// mplsInSegmentInterface and mplsInSegmentLabel
// objects. The purpose of this table is to
// provide the manager with an alternative
// means by which to locate in-segments.
type MPLSLSRSTDMIB_MplsInSegmentMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents one interface and incoming label pair. 
    // In cases where the label cannot fit into the mplsInSegmentLabel object, the
    // mplsInSegmentLabelPtr will indicate this by being set to the first
    // accessible column in the appropriate extension table's row, and the
    // mplsInSegmentLabel SHOULD be set to 0. In all other cases when the label is
    // represented within the mplsInSegmentLabel object, the mplsInSegmentLabelPtr
    // MUST be 0.0.  Implementors need to be aware that if the value of the
    // mplsInSegmentMapLabelPtrIndex (an OID) has more that 111 sub-identifiers,
    // then OIDs of column instances in this table will have more than 128
    // sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3.
    // The type is slice of
    // MPLSLSRSTDMIB_MplsInSegmentMapTable_MplsInSegmentMapEntry.
    MplsInSegmentMapEntry []*MPLSLSRSTDMIB_MplsInSegmentMapTable_MplsInSegmentMapEntry
}

func (mplsInSegmentMapTable *MPLSLSRSTDMIB_MplsInSegmentMapTable) GetEntityData() *types.CommonEntityData {
    mplsInSegmentMapTable.EntityData.YFilter = mplsInSegmentMapTable.YFilter
    mplsInSegmentMapTable.EntityData.YangName = "mplsInSegmentMapTable"
    mplsInSegmentMapTable.EntityData.BundleName = "cisco_ios_xe"
    mplsInSegmentMapTable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsInSegmentMapTable.EntityData.SegmentPath = "mplsInSegmentMapTable"
    mplsInSegmentMapTable.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/" + mplsInSegmentMapTable.EntityData.SegmentPath
    mplsInSegmentMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInSegmentMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInSegmentMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInSegmentMapTable.EntityData.Children = types.NewOrderedMap()
    mplsInSegmentMapTable.EntityData.Children.Append("mplsInSegmentMapEntry", types.YChild{"MplsInSegmentMapEntry", nil})
    for i := range mplsInSegmentMapTable.MplsInSegmentMapEntry {
        mplsInSegmentMapTable.EntityData.Children.Append(types.GetSegmentPath(mplsInSegmentMapTable.MplsInSegmentMapEntry[i]), types.YChild{"MplsInSegmentMapEntry", mplsInSegmentMapTable.MplsInSegmentMapEntry[i]})
    }
    mplsInSegmentMapTable.EntityData.Leafs = types.NewOrderedMap()

    mplsInSegmentMapTable.EntityData.YListKeys = []string {}

    return &(mplsInSegmentMapTable.EntityData)
}

// MPLSLSRSTDMIB_MplsInSegmentMapTable_MplsInSegmentMapEntry
// An entry in this table represents one interface
// and incoming label pair.
// 
// In cases where the label cannot fit into the
// mplsInSegmentLabel object, the mplsInSegmentLabelPtr
// will indicate this by being set to the first accessible
// column in the appropriate extension table's row,
// and the mplsInSegmentLabel SHOULD be set to 0.
// In all other cases when the label is
// represented within the mplsInSegmentLabel object, the
// mplsInSegmentLabelPtr MUST be 0.0.
// 
// Implementors need to be aware that if the value of
// the mplsInSegmentMapLabelPtrIndex (an OID) has more
// that 111 sub-identifiers, then OIDs of column
// instances in this table will have more than 128
// sub-identifiers and cannot be accessed using SNMPv1,
// SNMPv2c, or SNMPv3.
type MPLSLSRSTDMIB_MplsInSegmentMapTable_MplsInSegmentMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This index contains the same value as the
    // mplsInSegmentIndex in the mplsInSegmentTable. The type is interface{} with
    // range: 0..2147483647.
    MplsInSegmentMapInterface interface{}

    // This attribute is a key. This index contains the same value as the
    // mplsInSegmentLabel in the mplsInSegmentTable. The type is interface{} with
    // range: 0..4294967295.
    MplsInSegmentMapLabel interface{}

    // This attribute is a key. This index contains the same value as the
    // mplsInSegmentLabelPtr.  If the label for the InSegment cannot be
    // represented fully within the mplsInSegmentLabel object, this index MUST
    // point to the first accessible column of a conceptual row in an external
    // table containing the label.  In this case, the mplsInSegmentTopLabel object
    // SHOULD be set to 0 and ignored. This object MUST be set to zeroDotZero
    // otherwise. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    MplsInSegmentMapLabelPtrIndex interface{}

    // The mplsInSegmentIndex that corresponds to the mplsInSegmentInterface and
    // mplsInSegmentLabel, or the mplsInSegmentInterface and
    // mplsInSegmentLabelPtr, if applicable. The string containing the single
    // octet 0x00 MUST not be returned. The type is string with length: 1..24.
    MplsInSegmentMapIndex interface{}
}

func (mplsInSegmentMapEntry *MPLSLSRSTDMIB_MplsInSegmentMapTable_MplsInSegmentMapEntry) GetEntityData() *types.CommonEntityData {
    mplsInSegmentMapEntry.EntityData.YFilter = mplsInSegmentMapEntry.YFilter
    mplsInSegmentMapEntry.EntityData.YangName = "mplsInSegmentMapEntry"
    mplsInSegmentMapEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsInSegmentMapEntry.EntityData.ParentYangName = "mplsInSegmentMapTable"
    mplsInSegmentMapEntry.EntityData.SegmentPath = "mplsInSegmentMapEntry" + types.AddKeyToken(mplsInSegmentMapEntry.MplsInSegmentMapInterface, "mplsInSegmentMapInterface") + types.AddKeyToken(mplsInSegmentMapEntry.MplsInSegmentMapLabel, "mplsInSegmentMapLabel") + types.AddKeyToken(mplsInSegmentMapEntry.MplsInSegmentMapLabelPtrIndex, "mplsInSegmentMapLabelPtrIndex")
    mplsInSegmentMapEntry.EntityData.AbsolutePath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB/mplsInSegmentMapTable/" + mplsInSegmentMapEntry.EntityData.SegmentPath
    mplsInSegmentMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsInSegmentMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsInSegmentMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsInSegmentMapEntry.EntityData.Children = types.NewOrderedMap()
    mplsInSegmentMapEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsInSegmentMapEntry.EntityData.Leafs.Append("mplsInSegmentMapInterface", types.YLeaf{"MplsInSegmentMapInterface", mplsInSegmentMapEntry.MplsInSegmentMapInterface})
    mplsInSegmentMapEntry.EntityData.Leafs.Append("mplsInSegmentMapLabel", types.YLeaf{"MplsInSegmentMapLabel", mplsInSegmentMapEntry.MplsInSegmentMapLabel})
    mplsInSegmentMapEntry.EntityData.Leafs.Append("mplsInSegmentMapLabelPtrIndex", types.YLeaf{"MplsInSegmentMapLabelPtrIndex", mplsInSegmentMapEntry.MplsInSegmentMapLabelPtrIndex})
    mplsInSegmentMapEntry.EntityData.Leafs.Append("mplsInSegmentMapIndex", types.YLeaf{"MplsInSegmentMapIndex", mplsInSegmentMapEntry.MplsInSegmentMapIndex})

    mplsInSegmentMapEntry.EntityData.YListKeys = []string {"MplsInSegmentMapInterface", "MplsInSegmentMapLabel", "MplsInSegmentMapLabelPtrIndex"}

    return &(mplsInSegmentMapEntry.EntityData)
}

