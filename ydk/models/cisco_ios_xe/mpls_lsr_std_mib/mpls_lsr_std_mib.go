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

    
    Mplslsrobjects MPLSLSRSTDMIB_Mplslsrobjects

    // This table specifies per-interface MPLS capability and associated
    // information.
    Mplsinterfacetable MPLSLSRSTDMIB_Mplsinterfacetable

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
    Mplsinsegmenttable MPLSLSRSTDMIB_Mplsinsegmenttable

    // This table contains a representation of the outgoing segments from an LSR.
    Mplsoutsegmenttable MPLSLSRSTDMIB_Mplsoutsegmenttable

    // This table specifies information for switching between LSP segments.  It
    // supports point-to-point, point-to-multipoint and multipoint-to-point
    // connections.  mplsLabelStackTable specifies the label stack information for
    // a cross-connect LSR and is referred to from mplsXCTable.
    Mplsxctable MPLSLSRSTDMIB_Mplsxctable

    // This table specifies the label stack to be pushed onto a packet, beneath
    // the top label.  Entries into this table are referred to from mplsXCTable.
    Mplslabelstacktable MPLSLSRSTDMIB_Mplslabelstacktable

    // This table specifies the mapping from the mplsInSegmentIndex to the
    // corresponding mplsInSegmentInterface and mplsInSegmentLabel objects. The
    // purpose of this table is to provide the manager with an alternative means
    // by which to locate in-segments.
    Mplsinsegmentmaptable MPLSLSRSTDMIB_Mplsinsegmentmaptable
}

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSLSRSTDMIB.EntityData.YFilter = mPLSLSRSTDMIB.YFilter
    mPLSLSRSTDMIB.EntityData.YangName = "MPLS-LSR-STD-MIB"
    mPLSLSRSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSLSRSTDMIB.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mPLSLSRSTDMIB.EntityData.SegmentPath = "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB"
    mPLSLSRSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSLSRSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSLSRSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSLSRSTDMIB.EntityData.Children = make(map[string]types.YChild)
    mPLSLSRSTDMIB.EntityData.Children["mplsLsrObjects"] = types.YChild{"Mplslsrobjects", &mPLSLSRSTDMIB.Mplslsrobjects}
    mPLSLSRSTDMIB.EntityData.Children["mplsInterfaceTable"] = types.YChild{"Mplsinterfacetable", &mPLSLSRSTDMIB.Mplsinterfacetable}
    mPLSLSRSTDMIB.EntityData.Children["mplsInSegmentTable"] = types.YChild{"Mplsinsegmenttable", &mPLSLSRSTDMIB.Mplsinsegmenttable}
    mPLSLSRSTDMIB.EntityData.Children["mplsOutSegmentTable"] = types.YChild{"Mplsoutsegmenttable", &mPLSLSRSTDMIB.Mplsoutsegmenttable}
    mPLSLSRSTDMIB.EntityData.Children["mplsXCTable"] = types.YChild{"Mplsxctable", &mPLSLSRSTDMIB.Mplsxctable}
    mPLSLSRSTDMIB.EntityData.Children["mplsLabelStackTable"] = types.YChild{"Mplslabelstacktable", &mPLSLSRSTDMIB.Mplslabelstacktable}
    mPLSLSRSTDMIB.EntityData.Children["mplsInSegmentMapTable"] = types.YChild{"Mplsinsegmentmaptable", &mPLSLSRSTDMIB.Mplsinsegmentmaptable}
    mPLSLSRSTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mPLSLSRSTDMIB.EntityData)
}

// MPLSLSRSTDMIB_Mplslsrobjects
type MPLSLSRSTDMIB_Mplslsrobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains the next available value to be used for
    // mplsInSegmentIndex when creating entries in the mplsInSegmentTable. The
    // special value of a string containing the single octet 0x00 indicates that
    // no new entries can be created in this table. Agents not allowing managers
    // to create entries in this table MUST set this object to this special value.
    // The type is string with length: 1..24.
    Mplsinsegmentindexnext interface{}

    // This object contains the next available value to be used for
    // mplsOutSegmentIndex when creating entries in the mplsOutSegmentTable. The
    // special value of a string containing the single octet 0x00 indicates that
    // no new entries can be created in this table. Agents not allowing managers
    // to create entries in this table MUST set this object to this special value.
    // The type is string with length: 1..24.
    Mplsoutsegmentindexnext interface{}

    // This object contains the next available value to be used for mplsXCIndex
    // when creating entries in the mplsXCTable. A special value of the zero
    // length string indicates that no more new entries can be created in the
    // relevant table.  Agents not allowing managers to create entries in this
    // table MUST set this value to the zero length string. The type is string
    // with length: 1..24.
    Mplsxcindexnext interface{}

    // The maximum stack depth supported by this LSR. The type is interface{} with
    // range: 1..2147483647.
    Mplsmaxlabelstackdepth interface{}

    // This object contains the next available value to be used for
    // mplsLabelStackIndex when creating entries in the mplsLabelStackTable. The
    // special string containing the single octet 0x00 indicates that no more new
    // entries can be created in the relevant table.  Agents not allowing managers
    // to create entries in this table MUST set this value to the string
    // containing the single octet 0x00. The type is string with length: 1..24.
    Mplslabelstackindexnext interface{}

    // If this object is set to true(1), then it enables the emission of mplsXCUp
    // and mplsXCDown notifications; otherwise these notifications are not
    // emitted. The type is bool.
    Mplsxcnotificationsenable interface{}
}

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetEntityData() *types.CommonEntityData {
    mplslsrobjects.EntityData.YFilter = mplslsrobjects.YFilter
    mplslsrobjects.EntityData.YangName = "mplsLsrObjects"
    mplslsrobjects.EntityData.BundleName = "cisco_ios_xe"
    mplslsrobjects.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplslsrobjects.EntityData.SegmentPath = "mplsLsrObjects"
    mplslsrobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplslsrobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplslsrobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplslsrobjects.EntityData.Children = make(map[string]types.YChild)
    mplslsrobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    mplslsrobjects.EntityData.Leafs["mplsInSegmentIndexNext"] = types.YLeaf{"Mplsinsegmentindexnext", mplslsrobjects.Mplsinsegmentindexnext}
    mplslsrobjects.EntityData.Leafs["mplsOutSegmentIndexNext"] = types.YLeaf{"Mplsoutsegmentindexnext", mplslsrobjects.Mplsoutsegmentindexnext}
    mplslsrobjects.EntityData.Leafs["mplsXCIndexNext"] = types.YLeaf{"Mplsxcindexnext", mplslsrobjects.Mplsxcindexnext}
    mplslsrobjects.EntityData.Leafs["mplsMaxLabelStackDepth"] = types.YLeaf{"Mplsmaxlabelstackdepth", mplslsrobjects.Mplsmaxlabelstackdepth}
    mplslsrobjects.EntityData.Leafs["mplsLabelStackIndexNext"] = types.YLeaf{"Mplslabelstackindexnext", mplslsrobjects.Mplslabelstackindexnext}
    mplslsrobjects.EntityData.Leafs["mplsXCNotificationsEnable"] = types.YLeaf{"Mplsxcnotificationsenable", mplslsrobjects.Mplsxcnotificationsenable}
    return &(mplslsrobjects.EntityData)
}

// MPLSLSRSTDMIB_Mplsinterfacetable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSLSRSTDMIB_Mplsinterfacetable struct {
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
    // MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry.
    Mplsinterfaceentry []MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry
}

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetEntityData() *types.CommonEntityData {
    mplsinterfacetable.EntityData.YFilter = mplsinterfacetable.YFilter
    mplsinterfacetable.EntityData.YangName = "mplsInterfaceTable"
    mplsinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    mplsinterfacetable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsinterfacetable.EntityData.SegmentPath = "mplsInterfaceTable"
    mplsinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinterfacetable.EntityData.Children = make(map[string]types.YChild)
    mplsinterfacetable.EntityData.Children["mplsInterfaceEntry"] = types.YChild{"Mplsinterfaceentry", nil}
    for i := range mplsinterfacetable.Mplsinterfaceentry {
        mplsinterfacetable.EntityData.Children[types.GetSegmentPath(&mplsinterfacetable.Mplsinterfaceentry[i])] = types.YChild{"Mplsinterfaceentry", &mplsinterfacetable.Mplsinterfaceentry[i]}
    }
    mplsinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsinterfacetable.EntityData)
}

// MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry
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
type MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is a unique index for an entry in the
    // MplsInterfaceTable.  A non-zero index for an entry indicates the ifIndex
    // for the corresponding interface entry of the MPLS-layer in the ifTable. The
    // entry with index 0 represents the per-platform label space and contains
    // parameters that apply to all interfaces that participate in the
    // per-platform label space. Other entries defined in this table represent
    // additional MPLS interfaces that may participate in either the per-platform
    // or per-interface label spaces, or both. The type is interface{} with range:
    // 0..2147483647.
    Mplsinterfaceindex interface{}

    // This is the minimum value of an MPLS label that this LSR is willing to
    // receive on this interface. The type is interface{} with range:
    // 0..4294967295.
    Mplsinterfacelabelminin interface{}

    // This is the maximum value of an MPLS label that this LSR is willing to
    // receive on this interface. The type is interface{} with range:
    // 0..4294967295.
    Mplsinterfacelabelmaxin interface{}

    // This is the minimum value of an MPLS label that this LSR is willing to send
    // on this interface. The type is interface{} with range: 0..4294967295.
    Mplsinterfacelabelminout interface{}

    // This is the maximum value of an MPLS label that this LSR is willing to send
    // on this interface. The type is interface{} with range: 0..4294967295.
    Mplsinterfacelabelmaxout interface{}

    // This value indicates the total amount of usable bandwidth on this interface
    // and is specified in kilobits per second (Kbps).  This variable is not
    // applicable when applied to the interface with index 0. When this value
    // cannot be measured, this value should contain the nominal bandwidth. The
    // type is interface{} with range: 0..4294967295. Units are kilobits per
    // second.
    Mplsinterfacetotalbandwidth interface{}

    // This value indicates the total amount of available bandwidth available on
    // this interface and is specified in kilobits per second (Kbps).  This value
    // is calculated as the difference between the amount of bandwidth currently
    // in use and that specified in mplsInterfaceTotalBandwidth.  This variable is
    // not applicable when applied to the interface with index 0. When this value
    // cannot be measured, this value should contain the nominal bandwidth. The
    // type is interface{} with range: 0..4294967295.
    Mplsinterfaceavailablebandwidth interface{}

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
    Mplsinterfacelabelparticipationtype interface{}

    // This object counts the number of labels that are in use at this point in
    // time on this interface in the incoming direction. If the interface
    // participates in only the per-platform label space, then the value of the
    // instance of this object MUST be identical to the value of the instance with
    // index 0. If the interface participates in the per-interface label space,
    // then the instance of this object MUST represent the number of per-interface
    // labels that are in use on this interface. The type is interface{} with
    // range: 0..4294967295.
    Mplsinterfaceperfinlabelsinuse interface{}

    // This object counts the number of labeled packets that have been received on
    // this interface and which were discarded because there was no matching
    // cross- connect entry. This object MUST count on a per- interface basis
    // regardless of which label space the interface participates in. The type is
    // interface{} with range: 0..4294967295.
    Mplsinterfaceperfinlabellookupfailures interface{}

    // This object counts the number of top-most labels in the outgoing label
    // stacks that are in use at this point in time on this interface. This object
    // MUST count on a per-interface basis regardless of which label space the
    // interface participates in. The type is interface{} with range:
    // 0..4294967295.
    Mplsinterfaceperfoutlabelsinuse interface{}

    // This object counts the number of outgoing MPLS packets that required
    // fragmentation before transmission on this interface. This object MUST count
    // on a per-interface basis regardless of which label space the interface
    // participates in. The type is interface{} with range: 0..4294967295.
    Mplsinterfaceperfoutfragmentedpkts interface{}
}

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetEntityData() *types.CommonEntityData {
    mplsinterfaceentry.EntityData.YFilter = mplsinterfaceentry.YFilter
    mplsinterfaceentry.EntityData.YangName = "mplsInterfaceEntry"
    mplsinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    mplsinterfaceentry.EntityData.ParentYangName = "mplsInterfaceTable"
    mplsinterfaceentry.EntityData.SegmentPath = "mplsInterfaceEntry" + "[mplsInterfaceIndex='" + fmt.Sprintf("%v", mplsinterfaceentry.Mplsinterfaceindex) + "']"
    mplsinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    mplsinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceIndex"] = types.YLeaf{"Mplsinterfaceindex", mplsinterfaceentry.Mplsinterfaceindex}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceLabelMinIn"] = types.YLeaf{"Mplsinterfacelabelminin", mplsinterfaceentry.Mplsinterfacelabelminin}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceLabelMaxIn"] = types.YLeaf{"Mplsinterfacelabelmaxin", mplsinterfaceentry.Mplsinterfacelabelmaxin}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceLabelMinOut"] = types.YLeaf{"Mplsinterfacelabelminout", mplsinterfaceentry.Mplsinterfacelabelminout}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceLabelMaxOut"] = types.YLeaf{"Mplsinterfacelabelmaxout", mplsinterfaceentry.Mplsinterfacelabelmaxout}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceTotalBandwidth"] = types.YLeaf{"Mplsinterfacetotalbandwidth", mplsinterfaceentry.Mplsinterfacetotalbandwidth}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceAvailableBandwidth"] = types.YLeaf{"Mplsinterfaceavailablebandwidth", mplsinterfaceentry.Mplsinterfaceavailablebandwidth}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfaceLabelParticipationType"] = types.YLeaf{"Mplsinterfacelabelparticipationtype", mplsinterfaceentry.Mplsinterfacelabelparticipationtype}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfacePerfInLabelsInUse"] = types.YLeaf{"Mplsinterfaceperfinlabelsinuse", mplsinterfaceentry.Mplsinterfaceperfinlabelsinuse}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfacePerfInLabelLookupFailures"] = types.YLeaf{"Mplsinterfaceperfinlabellookupfailures", mplsinterfaceentry.Mplsinterfaceperfinlabellookupfailures}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfacePerfOutLabelsInUse"] = types.YLeaf{"Mplsinterfaceperfoutlabelsinuse", mplsinterfaceentry.Mplsinterfaceperfoutlabelsinuse}
    mplsinterfaceentry.EntityData.Leafs["mplsInterfacePerfOutFragmentedPkts"] = types.YLeaf{"Mplsinterfaceperfoutfragmentedpkts", mplsinterfaceentry.Mplsinterfaceperfoutfragmentedpkts}
    return &(mplsinterfaceentry.EntityData)
}

// MPLSLSRSTDMIB_Mplsinsegmenttable
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
type MPLSLSRSTDMIB_Mplsinsegmenttable struct {
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
    // MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry.
    Mplsinsegmententry []MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry
}

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetEntityData() *types.CommonEntityData {
    mplsinsegmenttable.EntityData.YFilter = mplsinsegmenttable.YFilter
    mplsinsegmenttable.EntityData.YangName = "mplsInSegmentTable"
    mplsinsegmenttable.EntityData.BundleName = "cisco_ios_xe"
    mplsinsegmenttable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsinsegmenttable.EntityData.SegmentPath = "mplsInSegmentTable"
    mplsinsegmenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinsegmenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinsegmenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinsegmenttable.EntityData.Children = make(map[string]types.YChild)
    mplsinsegmenttable.EntityData.Children["mplsInSegmentEntry"] = types.YChild{"Mplsinsegmententry", nil}
    for i := range mplsinsegmenttable.Mplsinsegmententry {
        mplsinsegmenttable.EntityData.Children[types.GetSegmentPath(&mplsinsegmenttable.Mplsinsegmententry[i])] = types.YChild{"Mplsinsegmententry", &mplsinsegmenttable.Mplsinsegmententry[i]}
    }
    mplsinsegmenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsinsegmenttable.EntityData)
}

// MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry
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
type MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index for this in-segment. The string
    // containing the single octet 0x00 MUST not be used as an index. The type is
    // string with length: 1..24.
    Mplsinsegmentindex interface{}

    // This object represents the interface index for the incoming MPLS interface.
    // A value of zero represents all interfaces participating in the per-platform
    // label space.  This may only be used in cases where the incoming interface
    // and label are associated with the same mplsXCEntry. Specifically, given a
    // label and any incoming interface pair from the per-platform label space,
    // the outgoing label/interface mapping remains the same. If this is not the
    // case, then individual entries MUST exist that can then be mapped to unique
    // mplsXCEntries. The type is interface{} with range: 0..2147483647.
    Mplsinsegmentinterface interface{}

    // If the corresponding instance of mplsInSegmentLabelPtr is zeroDotZero then
    // this object MUST contain the incoming label associated with this
    // in-segment. If not this object SHOULD be zero and MUST be ignored. The type
    // is interface{} with range: 0..4294967295.
    Mplsinsegmentlabel interface{}

    // If the label for this segment cannot be represented fully within the
    // mplsInSegmentLabel object, this object MUST point to the first accessible
    // column of a conceptual row in an external table containing the label.  In
    // this case, the mplsInSegmentTopLabel object SHOULD be set to 0 and ignored.
    // This object MUST be set to zeroDotZero otherwise. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplsinsegmentlabelptr interface{}

    // The number of labels to pop from the incoming packet.  Normally only the
    // top label is popped from the packet and used for all switching decisions
    // for that packet.  This is indicated by setting this object to the default
    // value of 1. If an LSR supports popping of more than one label, this object
    // MUST be set to that number. This object cannot be modified if
    // mplsInSegmentRowStatus is active(1). The type is interface{} with range:
    // 1..2147483647.
    Mplsinsegmentnpop interface{}

    // The IANA address family [IANAFamily] of packets received on this segment,
    // which is used at an egress LSR to deliver them to the appropriate layer 3
    // entity. A value of other(0) indicates that the family type is either
    // unknown or undefined; this SHOULD NOT be used at an egress LSR. This object
    // cannot be modified if mplsInSegmentRowStatus is active(1). The type is
    // AddressFamilyNumbers.
    Mplsinsegmentaddrfamily interface{}

    // Index into mplsXCTable which identifies which cross- connect entry this
    // segment is part of.  The string containing the single octet 0x00 indicates
    // that this entry is not referred to by any cross-connect entry. When a
    // cross-connect entry is created which this in-segment is a part of, this
    // object is automatically updated to reflect the value of mplsXCIndex of that
    // cross-connect entry. The type is string with length: 1..24.
    Mplsinsegmentxcindex interface{}

    // Denotes the entity that created and is responsible for managing this
    // segment. The type is MplsOwner.
    Mplsinsegmentowner interface{}

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
    Mplsinsegmenttrafficparamptr interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table has a row in the active(1) state, no objects in
    // this row can be modified except the mplsInSegmentRowStatus and
    // mplsInSegmentStorageType. The type is RowStatus.
    Mplsinsegmentrowstatus interface{}

    // This variable indicates the storage type for this object. The agent MUST
    // ensure that this object's value remains consistent with the associated
    // mplsXCEntry. Conceptual rows having the value 'permanent' need not allow
    // write-access to any columnar objects in the row. The type is StorageType.
    Mplsinsegmentstoragetype interface{}

    // This value represents the total number of octets received by this segment.
    // It MUST be equal to the least significant 32 bits of
    // mplsInSegmentPerfHCOctets if mplsInSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..4294967295.
    Mplsinsegmentperfoctets interface{}

    // Total number of packets received by this segment. The type is interface{}
    // with range: 0..4294967295.
    Mplsinsegmentperfpackets interface{}

    // The number of errored packets received on this segment. The type is
    // interface{} with range: 0..4294967295.
    Mplsinsegmentperferrors interface{}

    // The number of labeled packets received on this in- segment, which were
    // chosen to be discarded even though no errors had been detected to prevent
    // their being transmitted.  One possible reason for discarding such a labeled
    // packet could be to free up buffer space. The type is interface{} with
    // range: 0..4294967295.
    Mplsinsegmentperfdiscards interface{}

    // The total number of octets received.  This is the 64 bit version of
    // mplsInSegmentPerfOctets, if mplsInSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..18446744073709551615.
    Mplsinsegmentperfhcoctets interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this segment's Counter32 or Counter64 suffered a discontinuity. If no
    // such discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    Mplsinsegmentperfdiscontinuitytime interface{}
}

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetEntityData() *types.CommonEntityData {
    mplsinsegmententry.EntityData.YFilter = mplsinsegmententry.YFilter
    mplsinsegmententry.EntityData.YangName = "mplsInSegmentEntry"
    mplsinsegmententry.EntityData.BundleName = "cisco_ios_xe"
    mplsinsegmententry.EntityData.ParentYangName = "mplsInSegmentTable"
    mplsinsegmententry.EntityData.SegmentPath = "mplsInSegmentEntry" + "[mplsInSegmentIndex='" + fmt.Sprintf("%v", mplsinsegmententry.Mplsinsegmentindex) + "']"
    mplsinsegmententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinsegmententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinsegmententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinsegmententry.EntityData.Children = make(map[string]types.YChild)
    mplsinsegmententry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentIndex"] = types.YLeaf{"Mplsinsegmentindex", mplsinsegmententry.Mplsinsegmentindex}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentInterface"] = types.YLeaf{"Mplsinsegmentinterface", mplsinsegmententry.Mplsinsegmentinterface}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentLabel"] = types.YLeaf{"Mplsinsegmentlabel", mplsinsegmententry.Mplsinsegmentlabel}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentLabelPtr"] = types.YLeaf{"Mplsinsegmentlabelptr", mplsinsegmententry.Mplsinsegmentlabelptr}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentNPop"] = types.YLeaf{"Mplsinsegmentnpop", mplsinsegmententry.Mplsinsegmentnpop}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentAddrFamily"] = types.YLeaf{"Mplsinsegmentaddrfamily", mplsinsegmententry.Mplsinsegmentaddrfamily}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentXCIndex"] = types.YLeaf{"Mplsinsegmentxcindex", mplsinsegmententry.Mplsinsegmentxcindex}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentOwner"] = types.YLeaf{"Mplsinsegmentowner", mplsinsegmententry.Mplsinsegmentowner}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentTrafficParamPtr"] = types.YLeaf{"Mplsinsegmenttrafficparamptr", mplsinsegmententry.Mplsinsegmenttrafficparamptr}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentRowStatus"] = types.YLeaf{"Mplsinsegmentrowstatus", mplsinsegmententry.Mplsinsegmentrowstatus}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentStorageType"] = types.YLeaf{"Mplsinsegmentstoragetype", mplsinsegmententry.Mplsinsegmentstoragetype}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentPerfOctets"] = types.YLeaf{"Mplsinsegmentperfoctets", mplsinsegmententry.Mplsinsegmentperfoctets}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentPerfPackets"] = types.YLeaf{"Mplsinsegmentperfpackets", mplsinsegmententry.Mplsinsegmentperfpackets}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentPerfErrors"] = types.YLeaf{"Mplsinsegmentperferrors", mplsinsegmententry.Mplsinsegmentperferrors}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentPerfDiscards"] = types.YLeaf{"Mplsinsegmentperfdiscards", mplsinsegmententry.Mplsinsegmentperfdiscards}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentPerfHCOctets"] = types.YLeaf{"Mplsinsegmentperfhcoctets", mplsinsegmententry.Mplsinsegmentperfhcoctets}
    mplsinsegmententry.EntityData.Leafs["mplsInSegmentPerfDiscontinuityTime"] = types.YLeaf{"Mplsinsegmentperfdiscontinuitytime", mplsinsegmententry.Mplsinsegmentperfdiscontinuitytime}
    return &(mplsinsegmententry.EntityData)
}

// MPLSLSRSTDMIB_Mplsoutsegmenttable
// This table contains a representation of the outgoing
// segments from an LSR.
type MPLSLSRSTDMIB_Mplsoutsegmenttable struct {
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
    // The type is slice of MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry.
    Mplsoutsegmententry []MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry
}

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetEntityData() *types.CommonEntityData {
    mplsoutsegmenttable.EntityData.YFilter = mplsoutsegmenttable.YFilter
    mplsoutsegmenttable.EntityData.YangName = "mplsOutSegmentTable"
    mplsoutsegmenttable.EntityData.BundleName = "cisco_ios_xe"
    mplsoutsegmenttable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsoutsegmenttable.EntityData.SegmentPath = "mplsOutSegmentTable"
    mplsoutsegmenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsoutsegmenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsoutsegmenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsoutsegmenttable.EntityData.Children = make(map[string]types.YChild)
    mplsoutsegmenttable.EntityData.Children["mplsOutSegmentEntry"] = types.YChild{"Mplsoutsegmententry", nil}
    for i := range mplsoutsegmenttable.Mplsoutsegmententry {
        mplsoutsegmenttable.EntityData.Children[types.GetSegmentPath(&mplsoutsegmenttable.Mplsoutsegmententry[i])] = types.YChild{"Mplsoutsegmententry", &mplsoutsegmenttable.Mplsoutsegmententry[i]}
    }
    mplsoutsegmenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsoutsegmenttable.EntityData)
}

// MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry
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
type MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This value contains a unique index for this row.
    // While a value of a string containing the single octet 0x00 is not valid as
    // an index for entries in this table, it can be supplied as a valid value to
    // index the mplsXCTable to represent entries for which no out-segment has
    // been configured or exists. The type is string with length: 1..24.
    Mplsoutsegmentindex interface{}

    // This value must contain the interface index of the outgoing interface. This
    // object cannot be modified if mplsOutSegmentRowStatus is active(1). The
    // mplsOutSegmentRowStatus cannot be set to active(1) until this object is set
    // to a value corresponding to a valid ifEntry. The type is interface{} with
    // range: 0..2147483647.
    Mplsoutsegmentinterface interface{}

    // This value indicates whether or not a top label should be pushed onto the
    // outgoing packet's label stack.  The value of this variable MUST be set to
    // true(1) if the outgoing interface does not support pop-and-go (and no label
    // stack remains). For example, on ATM interface, or if the segment represents
    // a tunnel origination.  Note that it is considered an error in the case that
    // mplsOutSegmentPushTopLabel is set to false, but the cross-connect entry
    // which refers to this out-segment has a non-zero mplsLabelStackIndex.  The
    // LSR MUST ensure that this situation does not happen. This object cannot be
    // modified if mplsOutSegmentRowStatus is active(1). The type is bool.
    Mplsoutsegmentpushtoplabel interface{}

    // If mplsOutSegmentPushTopLabel is true then this represents the label that
    // should be pushed onto the top of the outgoing packet's label stack.
    // Otherwise this value SHOULD be set to 0 by the management station and MUST
    // be ignored by the agent. This object cannot be modified if
    // mplsOutSegmentRowStatus is active(1). The type is interface{} with range:
    // 0..4294967295.
    Mplsoutsegmenttoplabel interface{}

    // If the label for this segment cannot be represented fully within the
    // mplsOutSegmentLabel object, this object MUST point to the first accessible
    // column of a conceptual row in an external table containing the label.  In
    // this case, the mplsOutSegmentTopLabel object SHOULD be set to 0 and
    // ignored. This object MUST be set to zeroDotZero otherwise. The type is
    // string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplsoutsegmenttoplabelptr interface{}

    // Indicates the next hop Internet address type. Only values unknown(0),
    // ipv4(1) or ipv6(2) have to be supported.  A value of unknown(0) is allowed
    // only when the outgoing interface is of type point-to-point. If any other
    // unsupported values are attempted in a set operation, the agent MUST return
    // an inconsistentValue error. The type is InetAddressType.
    Mplsoutsegmentnexthopaddrtype interface{}

    // The internet address of the next hop. The type of this address is
    // determined by the value of the mplslOutSegmentNextHopAddrType object.  This
    // object cannot be modified if mplsOutSegmentRowStatus is active(1). The type
    // is string with length: 0..255.
    Mplsoutsegmentnexthopaddr interface{}

    // Index into mplsXCTable which identifies which cross- connect entry this
    // segment is part of.  A value of the string containing the single octet 0x00
    // indicates that this entry is not referred to by any cross-connect entry. 
    // When a cross-connect entry is created which this out-segment is a part of,
    // this object MUST be updated by the agent to reflect the value of
    // mplsXCIndex of that cross-connect entry. The type is string with length:
    // 1..24.
    Mplsoutsegmentxcindex interface{}

    // Denotes the entity which created and is responsible for managing this
    // segment. The type is MplsOwner.
    Mplsoutsegmentowner interface{}

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
    Mplsoutsegmenttrafficparamptr interface{}

    // For creating, modifying, and deleting this row. When a row in this table
    // has a row in the active(1) state, no objects in this row can be modified
    // except the mplsOutSegmentRowStatus or mplsOutSegmentStorageType. The type
    // is RowStatus.
    Mplsoutsegmentrowstatus interface{}

    // This variable indicates the storage type for this object. The agent MUST
    // ensure that this object's value remains consistent with the associated
    // mplsXCEntry. Conceptual rows having the value 'permanent' need not allow
    // write-access to any columnar objects in the row. The type is StorageType.
    Mplsoutsegmentstoragetype interface{}

    // This value contains the total number of octets sent on this segment. It
    // MUST be equal to the least significant 32 bits of
    // mplsOutSegmentPerfHCOctets if mplsOutSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..4294967295.
    Mplsoutsegmentperfoctets interface{}

    // This value contains the total number of packets sent on this segment. The
    // type is interface{} with range: 0..4294967295.
    Mplsoutsegmentperfpackets interface{}

    // Number of packets that could not be sent due to errors on this segment. The
    // type is interface{} with range: 0..4294967295.
    Mplsoutsegmentperferrors interface{}

    // The number of labeled packets attempted to be transmitted on this
    // out-segment, which were chosen to be discarded even though no errors had
    // been detected to prevent their being transmitted. One possible reason for
    // discarding such a labeled packet could be to free up buffer space. The type
    // is interface{} with range: 0..4294967295.
    Mplsoutsegmentperfdiscards interface{}

    // Total number of octets sent.  This is the 64 bit version of
    // mplsOutSegmentPerfOctets, if mplsOutSegmentPerfHCOctets is supported
    // according to the rules spelled out in RFC2863. The type is interface{} with
    // range: 0..18446744073709551615.
    Mplsoutsegmentperfhcoctets interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this segment's Counter32 or Counter64 suffered a discontinuity. If no
    // such discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    Mplsoutsegmentperfdiscontinuitytime interface{}
}

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetEntityData() *types.CommonEntityData {
    mplsoutsegmententry.EntityData.YFilter = mplsoutsegmententry.YFilter
    mplsoutsegmententry.EntityData.YangName = "mplsOutSegmentEntry"
    mplsoutsegmententry.EntityData.BundleName = "cisco_ios_xe"
    mplsoutsegmententry.EntityData.ParentYangName = "mplsOutSegmentTable"
    mplsoutsegmententry.EntityData.SegmentPath = "mplsOutSegmentEntry" + "[mplsOutSegmentIndex='" + fmt.Sprintf("%v", mplsoutsegmententry.Mplsoutsegmentindex) + "']"
    mplsoutsegmententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsoutsegmententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsoutsegmententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsoutsegmententry.EntityData.Children = make(map[string]types.YChild)
    mplsoutsegmententry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentIndex"] = types.YLeaf{"Mplsoutsegmentindex", mplsoutsegmententry.Mplsoutsegmentindex}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentInterface"] = types.YLeaf{"Mplsoutsegmentinterface", mplsoutsegmententry.Mplsoutsegmentinterface}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentPushTopLabel"] = types.YLeaf{"Mplsoutsegmentpushtoplabel", mplsoutsegmententry.Mplsoutsegmentpushtoplabel}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentTopLabel"] = types.YLeaf{"Mplsoutsegmenttoplabel", mplsoutsegmententry.Mplsoutsegmenttoplabel}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentTopLabelPtr"] = types.YLeaf{"Mplsoutsegmenttoplabelptr", mplsoutsegmententry.Mplsoutsegmenttoplabelptr}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentNextHopAddrType"] = types.YLeaf{"Mplsoutsegmentnexthopaddrtype", mplsoutsegmententry.Mplsoutsegmentnexthopaddrtype}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentNextHopAddr"] = types.YLeaf{"Mplsoutsegmentnexthopaddr", mplsoutsegmententry.Mplsoutsegmentnexthopaddr}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentXCIndex"] = types.YLeaf{"Mplsoutsegmentxcindex", mplsoutsegmententry.Mplsoutsegmentxcindex}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentOwner"] = types.YLeaf{"Mplsoutsegmentowner", mplsoutsegmententry.Mplsoutsegmentowner}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentTrafficParamPtr"] = types.YLeaf{"Mplsoutsegmenttrafficparamptr", mplsoutsegmententry.Mplsoutsegmenttrafficparamptr}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentRowStatus"] = types.YLeaf{"Mplsoutsegmentrowstatus", mplsoutsegmententry.Mplsoutsegmentrowstatus}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentStorageType"] = types.YLeaf{"Mplsoutsegmentstoragetype", mplsoutsegmententry.Mplsoutsegmentstoragetype}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentPerfOctets"] = types.YLeaf{"Mplsoutsegmentperfoctets", mplsoutsegmententry.Mplsoutsegmentperfoctets}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentPerfPackets"] = types.YLeaf{"Mplsoutsegmentperfpackets", mplsoutsegmententry.Mplsoutsegmentperfpackets}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentPerfErrors"] = types.YLeaf{"Mplsoutsegmentperferrors", mplsoutsegmententry.Mplsoutsegmentperferrors}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentPerfDiscards"] = types.YLeaf{"Mplsoutsegmentperfdiscards", mplsoutsegmententry.Mplsoutsegmentperfdiscards}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentPerfHCOctets"] = types.YLeaf{"Mplsoutsegmentperfhcoctets", mplsoutsegmententry.Mplsoutsegmentperfhcoctets}
    mplsoutsegmententry.EntityData.Leafs["mplsOutSegmentPerfDiscontinuityTime"] = types.YLeaf{"Mplsoutsegmentperfdiscontinuitytime", mplsoutsegmententry.Mplsoutsegmentperfdiscontinuitytime}
    return &(mplsoutsegmententry.EntityData)
}

// MPLSLSRSTDMIB_Mplsxctable
// This table specifies information for switching
// between LSP segments.  It supports point-to-point,
// point-to-multipoint and multipoint-to-point
// connections.  mplsLabelStackTable specifies the
// label stack information for a cross-connect LSR and
// is referred to from mplsXCTable.
type MPLSLSRSTDMIB_Mplsxctable struct {
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
    // protocol. The type is slice of MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry.
    Mplsxcentry []MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry
}

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetEntityData() *types.CommonEntityData {
    mplsxctable.EntityData.YFilter = mplsxctable.YFilter
    mplsxctable.EntityData.YangName = "mplsXCTable"
    mplsxctable.EntityData.BundleName = "cisco_ios_xe"
    mplsxctable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsxctable.EntityData.SegmentPath = "mplsXCTable"
    mplsxctable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsxctable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsxctable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsxctable.EntityData.Children = make(map[string]types.YChild)
    mplsxctable.EntityData.Children["mplsXCEntry"] = types.YChild{"Mplsxcentry", nil}
    for i := range mplsxctable.Mplsxcentry {
        mplsxctable.EntityData.Children[types.GetSegmentPath(&mplsxctable.Mplsxcentry[i])] = types.YChild{"Mplsxcentry", &mplsxctable.Mplsxcentry[i]}
    }
    mplsxctable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsxctable.EntityData)
}

// MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry
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
type MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Primary index for the conceptual row identifying a
    // group of cross-connect segments. The string containing a single octet 0x00
    // is an invalid index. The type is string with length: 1..24.
    Mplsxcindex interface{}

    // This attribute is a key. Incoming label index. If this object is set to the
    // string containing a single octet 0x00, this indicates a special case
    // outlined in the table's description above. In this case no corresponding
    // mplsInSegmentEntry shall exist. The type is string with length: 1..24.
    Mplsxcinsegmentindex interface{}

    // This attribute is a key. Index of out-segment for LSPs not terminating on
    // this LSR if not set to the string containing the single octet 0x00. If the
    // segment identified by this entry is terminating, then this object MUST be
    // set to the string containing a single octet 0x00 to indicate that no
    // corresponding mplsOutSegmentEntry shall exist. The type is string with
    // length: 1..24.
    Mplsxcoutsegmentindex interface{}

    // This value identifies the label switched path that this cross-connect entry
    // belongs to. This object cannot be modified if mplsXCRowStatus is active(1)
    // except for this object. The type is string with length: 2 | 6.
    Mplsxclspid interface{}

    // Primary index into mplsLabelStackTable identifying a stack of labels to be
    // pushed beneath the top label. Note that the top label identified by the
    // out- segment ensures that all the components of a multipoint-to-point
    // connection have the same outgoing label. A value of the string containing
    // the single octet 0x00 indicates that no labels are to be stacked beneath
    // the top label. This object cannot be modified if mplsXCRowStatus is
    // active(1). The type is string with length: 1..24.
    Mplsxclabelstackindex interface{}

    // Denotes the entity that created and is responsible for managing this
    // cross-connect. The type is MplsOwner.
    Mplsxcowner interface{}

    // For creating, modifying, and deleting this row. When a row in this table
    // has a row in the active(1) state, no objects in this row except this object
    // and the mplsXCStorageType can be modified. . The type is RowStatus.
    Mplsxcrowstatus interface{}

    // This variable indicates the storage type for this object. The agent MUST
    // ensure that the associated in and out segments also have the same
    // StorageType value and are restored consistently upon system restart. This
    // value SHOULD be set to permanent(4) if created as a result of a static LSP
    // configuration.  Conceptual rows having the value 'permanent' need not allow
    // write-access to any columnar objects in the row. The type is StorageType.
    Mplsxcstoragetype interface{}

    // The desired operational status of this segment. The type is
    // Mplsxcadminstatus.
    Mplsxcadminstatus interface{}

    // The actual operational status of this cross- connect. The type is
    // Mplsxcoperstatus.
    Mplsxcoperstatus interface{}
}

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetEntityData() *types.CommonEntityData {
    mplsxcentry.EntityData.YFilter = mplsxcentry.YFilter
    mplsxcentry.EntityData.YangName = "mplsXCEntry"
    mplsxcentry.EntityData.BundleName = "cisco_ios_xe"
    mplsxcentry.EntityData.ParentYangName = "mplsXCTable"
    mplsxcentry.EntityData.SegmentPath = "mplsXCEntry" + "[mplsXCIndex='" + fmt.Sprintf("%v", mplsxcentry.Mplsxcindex) + "']" + "[mplsXCInSegmentIndex='" + fmt.Sprintf("%v", mplsxcentry.Mplsxcinsegmentindex) + "']" + "[mplsXCOutSegmentIndex='" + fmt.Sprintf("%v", mplsxcentry.Mplsxcoutsegmentindex) + "']"
    mplsxcentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsxcentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsxcentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsxcentry.EntityData.Children = make(map[string]types.YChild)
    mplsxcentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsxcentry.EntityData.Leafs["mplsXCIndex"] = types.YLeaf{"Mplsxcindex", mplsxcentry.Mplsxcindex}
    mplsxcentry.EntityData.Leafs["mplsXCInSegmentIndex"] = types.YLeaf{"Mplsxcinsegmentindex", mplsxcentry.Mplsxcinsegmentindex}
    mplsxcentry.EntityData.Leafs["mplsXCOutSegmentIndex"] = types.YLeaf{"Mplsxcoutsegmentindex", mplsxcentry.Mplsxcoutsegmentindex}
    mplsxcentry.EntityData.Leafs["mplsXCLspId"] = types.YLeaf{"Mplsxclspid", mplsxcentry.Mplsxclspid}
    mplsxcentry.EntityData.Leafs["mplsXCLabelStackIndex"] = types.YLeaf{"Mplsxclabelstackindex", mplsxcentry.Mplsxclabelstackindex}
    mplsxcentry.EntityData.Leafs["mplsXCOwner"] = types.YLeaf{"Mplsxcowner", mplsxcentry.Mplsxcowner}
    mplsxcentry.EntityData.Leafs["mplsXCRowStatus"] = types.YLeaf{"Mplsxcrowstatus", mplsxcentry.Mplsxcrowstatus}
    mplsxcentry.EntityData.Leafs["mplsXCStorageType"] = types.YLeaf{"Mplsxcstoragetype", mplsxcentry.Mplsxcstoragetype}
    mplsxcentry.EntityData.Leafs["mplsXCAdminStatus"] = types.YLeaf{"Mplsxcadminstatus", mplsxcentry.Mplsxcadminstatus}
    mplsxcentry.EntityData.Leafs["mplsXCOperStatus"] = types.YLeaf{"Mplsxcoperstatus", mplsxcentry.Mplsxcoperstatus}
    return &(mplsxcentry.EntityData)
}

// MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus represents The desired operational status of this segment.
type MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus string

const (
    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus_up MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus = "up"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus_down MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus = "down"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus_testing MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcadminstatus = "testing"
)

// MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus represents connect.
type MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus string

const (
    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus_up MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus = "up"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus_down MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus = "down"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus_testing MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus = "testing"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus_unknown MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus = "unknown"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus_dormant MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus = "dormant"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus_notPresent MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus = "notPresent"

    MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus_lowerLayerDown MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoperstatus = "lowerLayerDown"
)

// MPLSLSRSTDMIB_Mplslabelstacktable
// This table specifies the label stack to be pushed
// onto a packet, beneath the top label.  Entries into
// this table are referred to from mplsXCTable.
type MPLSLSRSTDMIB_Mplslabelstacktable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents one label which is to be pushed onto an
    // outgoing packet, beneath the top label.  An entry can be created by a
    // network administrator or by an SNMP agent as instructed by an MPLS
    // signaling protocol. The type is slice of
    // MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry.
    Mplslabelstackentry []MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry
}

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetEntityData() *types.CommonEntityData {
    mplslabelstacktable.EntityData.YFilter = mplslabelstacktable.YFilter
    mplslabelstacktable.EntityData.YangName = "mplsLabelStackTable"
    mplslabelstacktable.EntityData.BundleName = "cisco_ios_xe"
    mplslabelstacktable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplslabelstacktable.EntityData.SegmentPath = "mplsLabelStackTable"
    mplslabelstacktable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplslabelstacktable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplslabelstacktable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplslabelstacktable.EntityData.Children = make(map[string]types.YChild)
    mplslabelstacktable.EntityData.Children["mplsLabelStackEntry"] = types.YChild{"Mplslabelstackentry", nil}
    for i := range mplslabelstacktable.Mplslabelstackentry {
        mplslabelstacktable.EntityData.Children[types.GetSegmentPath(&mplslabelstacktable.Mplslabelstackentry[i])] = types.YChild{"Mplslabelstackentry", &mplslabelstacktable.Mplslabelstackentry[i]}
    }
    mplslabelstacktable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplslabelstacktable.EntityData)
}

// MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry
// An entry in this table represents one label which is
// to be pushed onto an outgoing packet, beneath the
// top label.  An entry can be created by a network
// administrator or by an SNMP agent as instructed by
// an MPLS signaling protocol.
type MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Primary index for this row identifying a stack of
    // labels to be pushed on an outgoing packet, beneath the top label. An index
    // containing the string with a single octet 0x00 MUST not be used. The type
    // is string with length: 1..24.
    Mplslabelstackindex interface{}

    // This attribute is a key. Secondary index for this row identifying one label
    // of the stack.  Note that an entry with a smaller mplsLabelStackLabelIndex
    // would refer to a label higher up the label stack and would be popped at a
    // downstream LSR before a label represented by a higher
    // mplsLabelStackLabelIndex at a downstream LSR. The type is interface{} with
    // range: 1..2147483647.
    Mplslabelstacklabelindex interface{}

    // The label to pushed. The type is interface{} with range: 0..4294967295.
    Mplslabelstacklabel interface{}

    // If the label for this segment cannot be represented fully within the
    // mplsLabelStackLabel object, this object MUST point to the first accessible
    // column of a conceptual row in an external table containing the label.  In
    // this case, the mplsLabelStackLabel object SHOULD be set to 0 and ignored.
    // This object MUST be set to zeroDotZero otherwise. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplslabelstacklabelptr interface{}

    // For creating, modifying, and deleting this row. When a row in this table
    // has a row in the active(1) state, no objects in this row except this object
    // and the mplsLabelStackStorageType can be modified. The type is RowStatus.
    Mplslabelstackrowstatus interface{}

    // This variable indicates the storage type for this object. This object
    // cannot be modified if mplsLabelStackRowStatus is active(1). No objects are
    // required to be writable for rows in this table with this object set to
    // permanent(4).  The agent MUST ensure that all related entries in this table
    // retain the same value for this object.  Agents MUST ensure that the storage
    // type for all entries related to a particular mplsXCEntry retain the same
    // value for this object as the mplsXCEntry's StorageType. The type is
    // StorageType.
    Mplslabelstackstoragetype interface{}
}

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetEntityData() *types.CommonEntityData {
    mplslabelstackentry.EntityData.YFilter = mplslabelstackentry.YFilter
    mplslabelstackentry.EntityData.YangName = "mplsLabelStackEntry"
    mplslabelstackentry.EntityData.BundleName = "cisco_ios_xe"
    mplslabelstackentry.EntityData.ParentYangName = "mplsLabelStackTable"
    mplslabelstackentry.EntityData.SegmentPath = "mplsLabelStackEntry" + "[mplsLabelStackIndex='" + fmt.Sprintf("%v", mplslabelstackentry.Mplslabelstackindex) + "']" + "[mplsLabelStackLabelIndex='" + fmt.Sprintf("%v", mplslabelstackentry.Mplslabelstacklabelindex) + "']"
    mplslabelstackentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplslabelstackentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplslabelstackentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplslabelstackentry.EntityData.Children = make(map[string]types.YChild)
    mplslabelstackentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplslabelstackentry.EntityData.Leafs["mplsLabelStackIndex"] = types.YLeaf{"Mplslabelstackindex", mplslabelstackentry.Mplslabelstackindex}
    mplslabelstackentry.EntityData.Leafs["mplsLabelStackLabelIndex"] = types.YLeaf{"Mplslabelstacklabelindex", mplslabelstackentry.Mplslabelstacklabelindex}
    mplslabelstackentry.EntityData.Leafs["mplsLabelStackLabel"] = types.YLeaf{"Mplslabelstacklabel", mplslabelstackentry.Mplslabelstacklabel}
    mplslabelstackentry.EntityData.Leafs["mplsLabelStackLabelPtr"] = types.YLeaf{"Mplslabelstacklabelptr", mplslabelstackentry.Mplslabelstacklabelptr}
    mplslabelstackentry.EntityData.Leafs["mplsLabelStackRowStatus"] = types.YLeaf{"Mplslabelstackrowstatus", mplslabelstackentry.Mplslabelstackrowstatus}
    mplslabelstackentry.EntityData.Leafs["mplsLabelStackStorageType"] = types.YLeaf{"Mplslabelstackstoragetype", mplslabelstackentry.Mplslabelstackstoragetype}
    return &(mplslabelstackentry.EntityData)
}

// MPLSLSRSTDMIB_Mplsinsegmentmaptable
// This table specifies the mapping from the
// mplsInSegmentIndex to the corresponding
// mplsInSegmentInterface and mplsInSegmentLabel
// objects. The purpose of this table is to
// provide the manager with an alternative
// means by which to locate in-segments.
type MPLSLSRSTDMIB_Mplsinsegmentmaptable struct {
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
    // MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry.
    Mplsinsegmentmapentry []MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry
}

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetEntityData() *types.CommonEntityData {
    mplsinsegmentmaptable.EntityData.YFilter = mplsinsegmentmaptable.YFilter
    mplsinsegmentmaptable.EntityData.YangName = "mplsInSegmentMapTable"
    mplsinsegmentmaptable.EntityData.BundleName = "cisco_ios_xe"
    mplsinsegmentmaptable.EntityData.ParentYangName = "MPLS-LSR-STD-MIB"
    mplsinsegmentmaptable.EntityData.SegmentPath = "mplsInSegmentMapTable"
    mplsinsegmentmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinsegmentmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinsegmentmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinsegmentmaptable.EntityData.Children = make(map[string]types.YChild)
    mplsinsegmentmaptable.EntityData.Children["mplsInSegmentMapEntry"] = types.YChild{"Mplsinsegmentmapentry", nil}
    for i := range mplsinsegmentmaptable.Mplsinsegmentmapentry {
        mplsinsegmentmaptable.EntityData.Children[types.GetSegmentPath(&mplsinsegmentmaptable.Mplsinsegmentmapentry[i])] = types.YChild{"Mplsinsegmentmapentry", &mplsinsegmentmaptable.Mplsinsegmentmapentry[i]}
    }
    mplsinsegmentmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsinsegmentmaptable.EntityData)
}

// MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry
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
type MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This index contains the same value as the
    // mplsInSegmentIndex in the mplsInSegmentTable. The type is interface{} with
    // range: 0..2147483647.
    Mplsinsegmentmapinterface interface{}

    // This attribute is a key. This index contains the same value as the
    // mplsInSegmentLabel in the mplsInSegmentTable. The type is interface{} with
    // range: 0..4294967295.
    Mplsinsegmentmaplabel interface{}

    // This attribute is a key. This index contains the same value as the
    // mplsInSegmentLabelPtr.  If the label for the InSegment cannot be
    // represented fully within the mplsInSegmentLabel object, this index MUST
    // point to the first accessible column of a conceptual row in an external
    // table containing the label.  In this case, the mplsInSegmentTopLabel object
    // SHOULD be set to 0 and ignored. This object MUST be set to zeroDotZero
    // otherwise. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplsinsegmentmaplabelptrindex interface{}

    // The mplsInSegmentIndex that corresponds to the mplsInSegmentInterface and
    // mplsInSegmentLabel, or the mplsInSegmentInterface and
    // mplsInSegmentLabelPtr, if applicable. The string containing the single
    // octet 0x00 MUST not be returned. The type is string with length: 1..24.
    Mplsinsegmentmapindex interface{}
}

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetEntityData() *types.CommonEntityData {
    mplsinsegmentmapentry.EntityData.YFilter = mplsinsegmentmapentry.YFilter
    mplsinsegmentmapentry.EntityData.YangName = "mplsInSegmentMapEntry"
    mplsinsegmentmapentry.EntityData.BundleName = "cisco_ios_xe"
    mplsinsegmentmapentry.EntityData.ParentYangName = "mplsInSegmentMapTable"
    mplsinsegmentmapentry.EntityData.SegmentPath = "mplsInSegmentMapEntry" + "[mplsInSegmentMapInterface='" + fmt.Sprintf("%v", mplsinsegmentmapentry.Mplsinsegmentmapinterface) + "']" + "[mplsInSegmentMapLabel='" + fmt.Sprintf("%v", mplsinsegmentmapentry.Mplsinsegmentmaplabel) + "']" + "[mplsInSegmentMapLabelPtrIndex='" + fmt.Sprintf("%v", mplsinsegmentmapentry.Mplsinsegmentmaplabelptrindex) + "']"
    mplsinsegmentmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsinsegmentmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsinsegmentmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsinsegmentmapentry.EntityData.Children = make(map[string]types.YChild)
    mplsinsegmentmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsinsegmentmapentry.EntityData.Leafs["mplsInSegmentMapInterface"] = types.YLeaf{"Mplsinsegmentmapinterface", mplsinsegmentmapentry.Mplsinsegmentmapinterface}
    mplsinsegmentmapentry.EntityData.Leafs["mplsInSegmentMapLabel"] = types.YLeaf{"Mplsinsegmentmaplabel", mplsinsegmentmapentry.Mplsinsegmentmaplabel}
    mplsinsegmentmapentry.EntityData.Leafs["mplsInSegmentMapLabelPtrIndex"] = types.YLeaf{"Mplsinsegmentmaplabelptrindex", mplsinsegmentmapentry.Mplsinsegmentmaplabelptrindex}
    mplsinsegmentmapentry.EntityData.Leafs["mplsInSegmentMapIndex"] = types.YLeaf{"Mplsinsegmentmapindex", mplsinsegmentmapentry.Mplsinsegmentmapindex}
    return &(mplsinsegmentmapentry.EntityData)
}

