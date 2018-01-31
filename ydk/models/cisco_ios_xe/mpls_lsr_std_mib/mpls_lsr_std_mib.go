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
    parent types.Entity
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

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetFilter() yfilter.YFilter { return mPLSLSRSTDMIB.YFilter }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) SetFilter(yf yfilter.YFilter) { mPLSLSRSTDMIB.YFilter = yf }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetGoName(yname string) string {
    if yname == "mplsLsrObjects" { return "Mplslsrobjects" }
    if yname == "mplsInterfaceTable" { return "Mplsinterfacetable" }
    if yname == "mplsInSegmentTable" { return "Mplsinsegmenttable" }
    if yname == "mplsOutSegmentTable" { return "Mplsoutsegmenttable" }
    if yname == "mplsXCTable" { return "Mplsxctable" }
    if yname == "mplsLabelStackTable" { return "Mplslabelstacktable" }
    if yname == "mplsInSegmentMapTable" { return "Mplsinsegmentmaptable" }
    return ""
}

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetSegmentPath() string {
    return "MPLS-LSR-STD-MIB:MPLS-LSR-STD-MIB"
}

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLsrObjects" {
        return &mPLSLSRSTDMIB.Mplslsrobjects
    }
    if childYangName == "mplsInterfaceTable" {
        return &mPLSLSRSTDMIB.Mplsinterfacetable
    }
    if childYangName == "mplsInSegmentTable" {
        return &mPLSLSRSTDMIB.Mplsinsegmenttable
    }
    if childYangName == "mplsOutSegmentTable" {
        return &mPLSLSRSTDMIB.Mplsoutsegmenttable
    }
    if childYangName == "mplsXCTable" {
        return &mPLSLSRSTDMIB.Mplsxctable
    }
    if childYangName == "mplsLabelStackTable" {
        return &mPLSLSRSTDMIB.Mplslabelstacktable
    }
    if childYangName == "mplsInSegmentMapTable" {
        return &mPLSLSRSTDMIB.Mplsinsegmentmaptable
    }
    return nil
}

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mplsLsrObjects"] = &mPLSLSRSTDMIB.Mplslsrobjects
    children["mplsInterfaceTable"] = &mPLSLSRSTDMIB.Mplsinterfacetable
    children["mplsInSegmentTable"] = &mPLSLSRSTDMIB.Mplsinsegmenttable
    children["mplsOutSegmentTable"] = &mPLSLSRSTDMIB.Mplsoutsegmenttable
    children["mplsXCTable"] = &mPLSLSRSTDMIB.Mplsxctable
    children["mplsLabelStackTable"] = &mPLSLSRSTDMIB.Mplslabelstacktable
    children["mplsInSegmentMapTable"] = &mPLSLSRSTDMIB.Mplsinsegmentmaptable
    return children
}

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetYangName() string { return "MPLS-LSR-STD-MIB" }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) SetParent(parent types.Entity) { mPLSLSRSTDMIB.parent = parent }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetParent() types.Entity { return mPLSLSRSTDMIB.parent }

func (mPLSLSRSTDMIB *MPLSLSRSTDMIB) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

// MPLSLSRSTDMIB_Mplslsrobjects
type MPLSLSRSTDMIB_Mplslsrobjects struct {
    parent types.Entity
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

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetFilter() yfilter.YFilter { return mplslsrobjects.YFilter }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) SetFilter(yf yfilter.YFilter) { mplslsrobjects.YFilter = yf }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetGoName(yname string) string {
    if yname == "mplsInSegmentIndexNext" { return "Mplsinsegmentindexnext" }
    if yname == "mplsOutSegmentIndexNext" { return "Mplsoutsegmentindexnext" }
    if yname == "mplsXCIndexNext" { return "Mplsxcindexnext" }
    if yname == "mplsMaxLabelStackDepth" { return "Mplsmaxlabelstackdepth" }
    if yname == "mplsLabelStackIndexNext" { return "Mplslabelstackindexnext" }
    if yname == "mplsXCNotificationsEnable" { return "Mplsxcnotificationsenable" }
    return ""
}

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetSegmentPath() string {
    return "mplsLsrObjects"
}

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsInSegmentIndexNext"] = mplslsrobjects.Mplsinsegmentindexnext
    leafs["mplsOutSegmentIndexNext"] = mplslsrobjects.Mplsoutsegmentindexnext
    leafs["mplsXCIndexNext"] = mplslsrobjects.Mplsxcindexnext
    leafs["mplsMaxLabelStackDepth"] = mplslsrobjects.Mplsmaxlabelstackdepth
    leafs["mplsLabelStackIndexNext"] = mplslsrobjects.Mplslabelstackindexnext
    leafs["mplsXCNotificationsEnable"] = mplslsrobjects.Mplsxcnotificationsenable
    return leafs
}

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetBundleName() string { return "cisco_ios_xe" }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetYangName() string { return "mplsLsrObjects" }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) SetParent(parent types.Entity) { mplslsrobjects.parent = parent }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetParent() types.Entity { return mplslsrobjects.parent }

func (mplslsrobjects *MPLSLSRSTDMIB_Mplslsrobjects) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

// MPLSLSRSTDMIB_Mplsinterfacetable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSLSRSTDMIB_Mplsinterfacetable struct {
    parent types.Entity
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

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetFilter() yfilter.YFilter { return mplsinterfacetable.YFilter }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) SetFilter(yf yfilter.YFilter) { mplsinterfacetable.YFilter = yf }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetGoName(yname string) string {
    if yname == "mplsInterfaceEntry" { return "Mplsinterfaceentry" }
    return ""
}

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetSegmentPath() string {
    return "mplsInterfaceTable"
}

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsInterfaceEntry" {
        for _, c := range mplsinterfacetable.Mplsinterfaceentry {
            if mplsinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry{}
        mplsinterfacetable.Mplsinterfaceentry = append(mplsinterfacetable.Mplsinterfaceentry, child)
        return &mplsinterfacetable.Mplsinterfaceentry[len(mplsinterfacetable.Mplsinterfaceentry)-1]
    }
    return nil
}

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsinterfacetable.Mplsinterfaceentry {
        children[mplsinterfacetable.Mplsinterfaceentry[i].GetSegmentPath()] = &mplsinterfacetable.Mplsinterfaceentry[i]
    }
    return children
}

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetYangName() string { return "mplsInterfaceTable" }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) SetParent(parent types.Entity) { mplsinterfacetable.parent = parent }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetParent() types.Entity { return mplsinterfacetable.parent }

func (mplsinterfacetable *MPLSLSRSTDMIB_Mplsinterfacetable) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

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
    parent types.Entity
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

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetFilter() yfilter.YFilter { return mplsinterfaceentry.YFilter }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) SetFilter(yf yfilter.YFilter) { mplsinterfaceentry.YFilter = yf }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetGoName(yname string) string {
    if yname == "mplsInterfaceIndex" { return "Mplsinterfaceindex" }
    if yname == "mplsInterfaceLabelMinIn" { return "Mplsinterfacelabelminin" }
    if yname == "mplsInterfaceLabelMaxIn" { return "Mplsinterfacelabelmaxin" }
    if yname == "mplsInterfaceLabelMinOut" { return "Mplsinterfacelabelminout" }
    if yname == "mplsInterfaceLabelMaxOut" { return "Mplsinterfacelabelmaxout" }
    if yname == "mplsInterfaceTotalBandwidth" { return "Mplsinterfacetotalbandwidth" }
    if yname == "mplsInterfaceAvailableBandwidth" { return "Mplsinterfaceavailablebandwidth" }
    if yname == "mplsInterfaceLabelParticipationType" { return "Mplsinterfacelabelparticipationtype" }
    if yname == "mplsInterfacePerfInLabelsInUse" { return "Mplsinterfaceperfinlabelsinuse" }
    if yname == "mplsInterfacePerfInLabelLookupFailures" { return "Mplsinterfaceperfinlabellookupfailures" }
    if yname == "mplsInterfacePerfOutLabelsInUse" { return "Mplsinterfaceperfoutlabelsinuse" }
    if yname == "mplsInterfacePerfOutFragmentedPkts" { return "Mplsinterfaceperfoutfragmentedpkts" }
    return ""
}

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetSegmentPath() string {
    return "mplsInterfaceEntry" + "[mplsInterfaceIndex='" + fmt.Sprintf("%v", mplsinterfaceentry.Mplsinterfaceindex) + "']"
}

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsInterfaceIndex"] = mplsinterfaceentry.Mplsinterfaceindex
    leafs["mplsInterfaceLabelMinIn"] = mplsinterfaceentry.Mplsinterfacelabelminin
    leafs["mplsInterfaceLabelMaxIn"] = mplsinterfaceentry.Mplsinterfacelabelmaxin
    leafs["mplsInterfaceLabelMinOut"] = mplsinterfaceentry.Mplsinterfacelabelminout
    leafs["mplsInterfaceLabelMaxOut"] = mplsinterfaceentry.Mplsinterfacelabelmaxout
    leafs["mplsInterfaceTotalBandwidth"] = mplsinterfaceentry.Mplsinterfacetotalbandwidth
    leafs["mplsInterfaceAvailableBandwidth"] = mplsinterfaceentry.Mplsinterfaceavailablebandwidth
    leafs["mplsInterfaceLabelParticipationType"] = mplsinterfaceentry.Mplsinterfacelabelparticipationtype
    leafs["mplsInterfacePerfInLabelsInUse"] = mplsinterfaceentry.Mplsinterfaceperfinlabelsinuse
    leafs["mplsInterfacePerfInLabelLookupFailures"] = mplsinterfaceentry.Mplsinterfaceperfinlabellookupfailures
    leafs["mplsInterfacePerfOutLabelsInUse"] = mplsinterfaceentry.Mplsinterfaceperfoutlabelsinuse
    leafs["mplsInterfacePerfOutFragmentedPkts"] = mplsinterfaceentry.Mplsinterfaceperfoutfragmentedpkts
    return leafs
}

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetYangName() string { return "mplsInterfaceEntry" }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) SetParent(parent types.Entity) { mplsinterfaceentry.parent = parent }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetParent() types.Entity { return mplsinterfaceentry.parent }

func (mplsinterfaceentry *MPLSLSRSTDMIB_Mplsinterfacetable_Mplsinterfaceentry) GetParentYangName() string { return "mplsInterfaceTable" }

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
    parent types.Entity
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

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetFilter() yfilter.YFilter { return mplsinsegmenttable.YFilter }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) SetFilter(yf yfilter.YFilter) { mplsinsegmenttable.YFilter = yf }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetGoName(yname string) string {
    if yname == "mplsInSegmentEntry" { return "Mplsinsegmententry" }
    return ""
}

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetSegmentPath() string {
    return "mplsInSegmentTable"
}

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsInSegmentEntry" {
        for _, c := range mplsinsegmenttable.Mplsinsegmententry {
            if mplsinsegmenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry{}
        mplsinsegmenttable.Mplsinsegmententry = append(mplsinsegmenttable.Mplsinsegmententry, child)
        return &mplsinsegmenttable.Mplsinsegmententry[len(mplsinsegmenttable.Mplsinsegmententry)-1]
    }
    return nil
}

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsinsegmenttable.Mplsinsegmententry {
        children[mplsinsegmenttable.Mplsinsegmententry[i].GetSegmentPath()] = &mplsinsegmenttable.Mplsinsegmententry[i]
    }
    return children
}

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetYangName() string { return "mplsInSegmentTable" }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) SetParent(parent types.Entity) { mplsinsegmenttable.parent = parent }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetParent() types.Entity { return mplsinsegmenttable.parent }

func (mplsinsegmenttable *MPLSLSRSTDMIB_Mplsinsegmenttable) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

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
    parent types.Entity
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
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetFilter() yfilter.YFilter { return mplsinsegmententry.YFilter }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) SetFilter(yf yfilter.YFilter) { mplsinsegmententry.YFilter = yf }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetGoName(yname string) string {
    if yname == "mplsInSegmentIndex" { return "Mplsinsegmentindex" }
    if yname == "mplsInSegmentInterface" { return "Mplsinsegmentinterface" }
    if yname == "mplsInSegmentLabel" { return "Mplsinsegmentlabel" }
    if yname == "mplsInSegmentLabelPtr" { return "Mplsinsegmentlabelptr" }
    if yname == "mplsInSegmentNPop" { return "Mplsinsegmentnpop" }
    if yname == "mplsInSegmentAddrFamily" { return "Mplsinsegmentaddrfamily" }
    if yname == "mplsInSegmentXCIndex" { return "Mplsinsegmentxcindex" }
    if yname == "mplsInSegmentOwner" { return "Mplsinsegmentowner" }
    if yname == "mplsInSegmentTrafficParamPtr" { return "Mplsinsegmenttrafficparamptr" }
    if yname == "mplsInSegmentRowStatus" { return "Mplsinsegmentrowstatus" }
    if yname == "mplsInSegmentStorageType" { return "Mplsinsegmentstoragetype" }
    if yname == "mplsInSegmentPerfOctets" { return "Mplsinsegmentperfoctets" }
    if yname == "mplsInSegmentPerfPackets" { return "Mplsinsegmentperfpackets" }
    if yname == "mplsInSegmentPerfErrors" { return "Mplsinsegmentperferrors" }
    if yname == "mplsInSegmentPerfDiscards" { return "Mplsinsegmentperfdiscards" }
    if yname == "mplsInSegmentPerfHCOctets" { return "Mplsinsegmentperfhcoctets" }
    if yname == "mplsInSegmentPerfDiscontinuityTime" { return "Mplsinsegmentperfdiscontinuitytime" }
    return ""
}

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetSegmentPath() string {
    return "mplsInSegmentEntry" + "[mplsInSegmentIndex='" + fmt.Sprintf("%v", mplsinsegmententry.Mplsinsegmentindex) + "']"
}

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsInSegmentIndex"] = mplsinsegmententry.Mplsinsegmentindex
    leafs["mplsInSegmentInterface"] = mplsinsegmententry.Mplsinsegmentinterface
    leafs["mplsInSegmentLabel"] = mplsinsegmententry.Mplsinsegmentlabel
    leafs["mplsInSegmentLabelPtr"] = mplsinsegmententry.Mplsinsegmentlabelptr
    leafs["mplsInSegmentNPop"] = mplsinsegmententry.Mplsinsegmentnpop
    leafs["mplsInSegmentAddrFamily"] = mplsinsegmententry.Mplsinsegmentaddrfamily
    leafs["mplsInSegmentXCIndex"] = mplsinsegmententry.Mplsinsegmentxcindex
    leafs["mplsInSegmentOwner"] = mplsinsegmententry.Mplsinsegmentowner
    leafs["mplsInSegmentTrafficParamPtr"] = mplsinsegmententry.Mplsinsegmenttrafficparamptr
    leafs["mplsInSegmentRowStatus"] = mplsinsegmententry.Mplsinsegmentrowstatus
    leafs["mplsInSegmentStorageType"] = mplsinsegmententry.Mplsinsegmentstoragetype
    leafs["mplsInSegmentPerfOctets"] = mplsinsegmententry.Mplsinsegmentperfoctets
    leafs["mplsInSegmentPerfPackets"] = mplsinsegmententry.Mplsinsegmentperfpackets
    leafs["mplsInSegmentPerfErrors"] = mplsinsegmententry.Mplsinsegmentperferrors
    leafs["mplsInSegmentPerfDiscards"] = mplsinsegmententry.Mplsinsegmentperfdiscards
    leafs["mplsInSegmentPerfHCOctets"] = mplsinsegmententry.Mplsinsegmentperfhcoctets
    leafs["mplsInSegmentPerfDiscontinuityTime"] = mplsinsegmententry.Mplsinsegmentperfdiscontinuitytime
    return leafs
}

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetYangName() string { return "mplsInSegmentEntry" }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) SetParent(parent types.Entity) { mplsinsegmententry.parent = parent }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetParent() types.Entity { return mplsinsegmententry.parent }

func (mplsinsegmententry *MPLSLSRSTDMIB_Mplsinsegmenttable_Mplsinsegmententry) GetParentYangName() string { return "mplsInSegmentTable" }

// MPLSLSRSTDMIB_Mplsoutsegmenttable
// This table contains a representation of the outgoing
// segments from an LSR.
type MPLSLSRSTDMIB_Mplsoutsegmenttable struct {
    parent types.Entity
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

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetFilter() yfilter.YFilter { return mplsoutsegmenttable.YFilter }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) SetFilter(yf yfilter.YFilter) { mplsoutsegmenttable.YFilter = yf }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetGoName(yname string) string {
    if yname == "mplsOutSegmentEntry" { return "Mplsoutsegmententry" }
    return ""
}

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetSegmentPath() string {
    return "mplsOutSegmentTable"
}

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsOutSegmentEntry" {
        for _, c := range mplsoutsegmenttable.Mplsoutsegmententry {
            if mplsoutsegmenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry{}
        mplsoutsegmenttable.Mplsoutsegmententry = append(mplsoutsegmenttable.Mplsoutsegmententry, child)
        return &mplsoutsegmenttable.Mplsoutsegmententry[len(mplsoutsegmenttable.Mplsoutsegmententry)-1]
    }
    return nil
}

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsoutsegmenttable.Mplsoutsegmententry {
        children[mplsoutsegmenttable.Mplsoutsegmententry[i].GetSegmentPath()] = &mplsoutsegmenttable.Mplsoutsegmententry[i]
    }
    return children
}

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetYangName() string { return "mplsOutSegmentTable" }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) SetParent(parent types.Entity) { mplsoutsegmenttable.parent = parent }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetParent() types.Entity { return mplsoutsegmenttable.parent }

func (mplsoutsegmenttable *MPLSLSRSTDMIB_Mplsoutsegmenttable) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetFilter() yfilter.YFilter { return mplsoutsegmententry.YFilter }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) SetFilter(yf yfilter.YFilter) { mplsoutsegmententry.YFilter = yf }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetGoName(yname string) string {
    if yname == "mplsOutSegmentIndex" { return "Mplsoutsegmentindex" }
    if yname == "mplsOutSegmentInterface" { return "Mplsoutsegmentinterface" }
    if yname == "mplsOutSegmentPushTopLabel" { return "Mplsoutsegmentpushtoplabel" }
    if yname == "mplsOutSegmentTopLabel" { return "Mplsoutsegmenttoplabel" }
    if yname == "mplsOutSegmentTopLabelPtr" { return "Mplsoutsegmenttoplabelptr" }
    if yname == "mplsOutSegmentNextHopAddrType" { return "Mplsoutsegmentnexthopaddrtype" }
    if yname == "mplsOutSegmentNextHopAddr" { return "Mplsoutsegmentnexthopaddr" }
    if yname == "mplsOutSegmentXCIndex" { return "Mplsoutsegmentxcindex" }
    if yname == "mplsOutSegmentOwner" { return "Mplsoutsegmentowner" }
    if yname == "mplsOutSegmentTrafficParamPtr" { return "Mplsoutsegmenttrafficparamptr" }
    if yname == "mplsOutSegmentRowStatus" { return "Mplsoutsegmentrowstatus" }
    if yname == "mplsOutSegmentStorageType" { return "Mplsoutsegmentstoragetype" }
    if yname == "mplsOutSegmentPerfOctets" { return "Mplsoutsegmentperfoctets" }
    if yname == "mplsOutSegmentPerfPackets" { return "Mplsoutsegmentperfpackets" }
    if yname == "mplsOutSegmentPerfErrors" { return "Mplsoutsegmentperferrors" }
    if yname == "mplsOutSegmentPerfDiscards" { return "Mplsoutsegmentperfdiscards" }
    if yname == "mplsOutSegmentPerfHCOctets" { return "Mplsoutsegmentperfhcoctets" }
    if yname == "mplsOutSegmentPerfDiscontinuityTime" { return "Mplsoutsegmentperfdiscontinuitytime" }
    return ""
}

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetSegmentPath() string {
    return "mplsOutSegmentEntry" + "[mplsOutSegmentIndex='" + fmt.Sprintf("%v", mplsoutsegmententry.Mplsoutsegmentindex) + "']"
}

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsOutSegmentIndex"] = mplsoutsegmententry.Mplsoutsegmentindex
    leafs["mplsOutSegmentInterface"] = mplsoutsegmententry.Mplsoutsegmentinterface
    leafs["mplsOutSegmentPushTopLabel"] = mplsoutsegmententry.Mplsoutsegmentpushtoplabel
    leafs["mplsOutSegmentTopLabel"] = mplsoutsegmententry.Mplsoutsegmenttoplabel
    leafs["mplsOutSegmentTopLabelPtr"] = mplsoutsegmententry.Mplsoutsegmenttoplabelptr
    leafs["mplsOutSegmentNextHopAddrType"] = mplsoutsegmententry.Mplsoutsegmentnexthopaddrtype
    leafs["mplsOutSegmentNextHopAddr"] = mplsoutsegmententry.Mplsoutsegmentnexthopaddr
    leafs["mplsOutSegmentXCIndex"] = mplsoutsegmententry.Mplsoutsegmentxcindex
    leafs["mplsOutSegmentOwner"] = mplsoutsegmententry.Mplsoutsegmentowner
    leafs["mplsOutSegmentTrafficParamPtr"] = mplsoutsegmententry.Mplsoutsegmenttrafficparamptr
    leafs["mplsOutSegmentRowStatus"] = mplsoutsegmententry.Mplsoutsegmentrowstatus
    leafs["mplsOutSegmentStorageType"] = mplsoutsegmententry.Mplsoutsegmentstoragetype
    leafs["mplsOutSegmentPerfOctets"] = mplsoutsegmententry.Mplsoutsegmentperfoctets
    leafs["mplsOutSegmentPerfPackets"] = mplsoutsegmententry.Mplsoutsegmentperfpackets
    leafs["mplsOutSegmentPerfErrors"] = mplsoutsegmententry.Mplsoutsegmentperferrors
    leafs["mplsOutSegmentPerfDiscards"] = mplsoutsegmententry.Mplsoutsegmentperfdiscards
    leafs["mplsOutSegmentPerfHCOctets"] = mplsoutsegmententry.Mplsoutsegmentperfhcoctets
    leafs["mplsOutSegmentPerfDiscontinuityTime"] = mplsoutsegmententry.Mplsoutsegmentperfdiscontinuitytime
    return leafs
}

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetYangName() string { return "mplsOutSegmentEntry" }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) SetParent(parent types.Entity) { mplsoutsegmententry.parent = parent }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetParent() types.Entity { return mplsoutsegmententry.parent }

func (mplsoutsegmententry *MPLSLSRSTDMIB_Mplsoutsegmenttable_Mplsoutsegmententry) GetParentYangName() string { return "mplsOutSegmentTable" }

// MPLSLSRSTDMIB_Mplsxctable
// This table specifies information for switching
// between LSP segments.  It supports point-to-point,
// point-to-multipoint and multipoint-to-point
// connections.  mplsLabelStackTable specifies the
// label stack information for a cross-connect LSR and
// is referred to from mplsXCTable.
type MPLSLSRSTDMIB_Mplsxctable struct {
    parent types.Entity
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

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetFilter() yfilter.YFilter { return mplsxctable.YFilter }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) SetFilter(yf yfilter.YFilter) { mplsxctable.YFilter = yf }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetGoName(yname string) string {
    if yname == "mplsXCEntry" { return "Mplsxcentry" }
    return ""
}

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetSegmentPath() string {
    return "mplsXCTable"
}

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsXCEntry" {
        for _, c := range mplsxctable.Mplsxcentry {
            if mplsxctable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry{}
        mplsxctable.Mplsxcentry = append(mplsxctable.Mplsxcentry, child)
        return &mplsxctable.Mplsxcentry[len(mplsxctable.Mplsxcentry)-1]
    }
    return nil
}

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsxctable.Mplsxcentry {
        children[mplsxctable.Mplsxcentry[i].GetSegmentPath()] = &mplsxctable.Mplsxcentry[i]
    }
    return children
}

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetYangName() string { return "mplsXCTable" }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) SetParent(parent types.Entity) { mplsxctable.parent = parent }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetParent() types.Entity { return mplsxctable.parent }

func (mplsxctable *MPLSLSRSTDMIB_Mplsxctable) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

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
    parent types.Entity
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

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetFilter() yfilter.YFilter { return mplsxcentry.YFilter }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) SetFilter(yf yfilter.YFilter) { mplsxcentry.YFilter = yf }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetGoName(yname string) string {
    if yname == "mplsXCIndex" { return "Mplsxcindex" }
    if yname == "mplsXCInSegmentIndex" { return "Mplsxcinsegmentindex" }
    if yname == "mplsXCOutSegmentIndex" { return "Mplsxcoutsegmentindex" }
    if yname == "mplsXCLspId" { return "Mplsxclspid" }
    if yname == "mplsXCLabelStackIndex" { return "Mplsxclabelstackindex" }
    if yname == "mplsXCOwner" { return "Mplsxcowner" }
    if yname == "mplsXCRowStatus" { return "Mplsxcrowstatus" }
    if yname == "mplsXCStorageType" { return "Mplsxcstoragetype" }
    if yname == "mplsXCAdminStatus" { return "Mplsxcadminstatus" }
    if yname == "mplsXCOperStatus" { return "Mplsxcoperstatus" }
    return ""
}

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetSegmentPath() string {
    return "mplsXCEntry" + "[mplsXCIndex='" + fmt.Sprintf("%v", mplsxcentry.Mplsxcindex) + "']" + "[mplsXCInSegmentIndex='" + fmt.Sprintf("%v", mplsxcentry.Mplsxcinsegmentindex) + "']" + "[mplsXCOutSegmentIndex='" + fmt.Sprintf("%v", mplsxcentry.Mplsxcoutsegmentindex) + "']"
}

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsXCIndex"] = mplsxcentry.Mplsxcindex
    leafs["mplsXCInSegmentIndex"] = mplsxcentry.Mplsxcinsegmentindex
    leafs["mplsXCOutSegmentIndex"] = mplsxcentry.Mplsxcoutsegmentindex
    leafs["mplsXCLspId"] = mplsxcentry.Mplsxclspid
    leafs["mplsXCLabelStackIndex"] = mplsxcentry.Mplsxclabelstackindex
    leafs["mplsXCOwner"] = mplsxcentry.Mplsxcowner
    leafs["mplsXCRowStatus"] = mplsxcentry.Mplsxcrowstatus
    leafs["mplsXCStorageType"] = mplsxcentry.Mplsxcstoragetype
    leafs["mplsXCAdminStatus"] = mplsxcentry.Mplsxcadminstatus
    leafs["mplsXCOperStatus"] = mplsxcentry.Mplsxcoperstatus
    return leafs
}

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetYangName() string { return "mplsXCEntry" }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) SetParent(parent types.Entity) { mplsxcentry.parent = parent }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetParent() types.Entity { return mplsxcentry.parent }

func (mplsxcentry *MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry) GetParentYangName() string { return "mplsXCTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents one label which is to be pushed onto an
    // outgoing packet, beneath the top label.  An entry can be created by a
    // network administrator or by an SNMP agent as instructed by an MPLS
    // signaling protocol. The type is slice of
    // MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry.
    Mplslabelstackentry []MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry
}

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetFilter() yfilter.YFilter { return mplslabelstacktable.YFilter }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) SetFilter(yf yfilter.YFilter) { mplslabelstacktable.YFilter = yf }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetGoName(yname string) string {
    if yname == "mplsLabelStackEntry" { return "Mplslabelstackentry" }
    return ""
}

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetSegmentPath() string {
    return "mplsLabelStackTable"
}

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsLabelStackEntry" {
        for _, c := range mplslabelstacktable.Mplslabelstackentry {
            if mplslabelstacktable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry{}
        mplslabelstacktable.Mplslabelstackentry = append(mplslabelstacktable.Mplslabelstackentry, child)
        return &mplslabelstacktable.Mplslabelstackentry[len(mplslabelstacktable.Mplslabelstackentry)-1]
    }
    return nil
}

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplslabelstacktable.Mplslabelstackentry {
        children[mplslabelstacktable.Mplslabelstackentry[i].GetSegmentPath()] = &mplslabelstacktable.Mplslabelstackentry[i]
    }
    return children
}

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetBundleName() string { return "cisco_ios_xe" }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetYangName() string { return "mplsLabelStackTable" }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) SetParent(parent types.Entity) { mplslabelstacktable.parent = parent }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetParent() types.Entity { return mplslabelstacktable.parent }

func (mplslabelstacktable *MPLSLSRSTDMIB_Mplslabelstacktable) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

// MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry
// An entry in this table represents one label which is
// to be pushed onto an outgoing packet, beneath the
// top label.  An entry can be created by a network
// administrator or by an SNMP agent as instructed by
// an MPLS signaling protocol.
type MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry struct {
    parent types.Entity
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
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetFilter() yfilter.YFilter { return mplslabelstackentry.YFilter }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) SetFilter(yf yfilter.YFilter) { mplslabelstackentry.YFilter = yf }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetGoName(yname string) string {
    if yname == "mplsLabelStackIndex" { return "Mplslabelstackindex" }
    if yname == "mplsLabelStackLabelIndex" { return "Mplslabelstacklabelindex" }
    if yname == "mplsLabelStackLabel" { return "Mplslabelstacklabel" }
    if yname == "mplsLabelStackLabelPtr" { return "Mplslabelstacklabelptr" }
    if yname == "mplsLabelStackRowStatus" { return "Mplslabelstackrowstatus" }
    if yname == "mplsLabelStackStorageType" { return "Mplslabelstackstoragetype" }
    return ""
}

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetSegmentPath() string {
    return "mplsLabelStackEntry" + "[mplsLabelStackIndex='" + fmt.Sprintf("%v", mplslabelstackentry.Mplslabelstackindex) + "']" + "[mplsLabelStackLabelIndex='" + fmt.Sprintf("%v", mplslabelstackentry.Mplslabelstacklabelindex) + "']"
}

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsLabelStackIndex"] = mplslabelstackentry.Mplslabelstackindex
    leafs["mplsLabelStackLabelIndex"] = mplslabelstackentry.Mplslabelstacklabelindex
    leafs["mplsLabelStackLabel"] = mplslabelstackentry.Mplslabelstacklabel
    leafs["mplsLabelStackLabelPtr"] = mplslabelstackentry.Mplslabelstacklabelptr
    leafs["mplsLabelStackRowStatus"] = mplslabelstackentry.Mplslabelstackrowstatus
    leafs["mplsLabelStackStorageType"] = mplslabelstackentry.Mplslabelstackstoragetype
    return leafs
}

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetYangName() string { return "mplsLabelStackEntry" }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) SetParent(parent types.Entity) { mplslabelstackentry.parent = parent }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetParent() types.Entity { return mplslabelstackentry.parent }

func (mplslabelstackentry *MPLSLSRSTDMIB_Mplslabelstacktable_Mplslabelstackentry) GetParentYangName() string { return "mplsLabelStackTable" }

// MPLSLSRSTDMIB_Mplsinsegmentmaptable
// This table specifies the mapping from the
// mplsInSegmentIndex to the corresponding
// mplsInSegmentInterface and mplsInSegmentLabel
// objects. The purpose of this table is to
// provide the manager with an alternative
// means by which to locate in-segments.
type MPLSLSRSTDMIB_Mplsinsegmentmaptable struct {
    parent types.Entity
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

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetFilter() yfilter.YFilter { return mplsinsegmentmaptable.YFilter }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) SetFilter(yf yfilter.YFilter) { mplsinsegmentmaptable.YFilter = yf }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetGoName(yname string) string {
    if yname == "mplsInSegmentMapEntry" { return "Mplsinsegmentmapentry" }
    return ""
}

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetSegmentPath() string {
    return "mplsInSegmentMapTable"
}

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsInSegmentMapEntry" {
        for _, c := range mplsinsegmentmaptable.Mplsinsegmentmapentry {
            if mplsinsegmentmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry{}
        mplsinsegmentmaptable.Mplsinsegmentmapentry = append(mplsinsegmentmaptable.Mplsinsegmentmapentry, child)
        return &mplsinsegmentmaptable.Mplsinsegmentmapentry[len(mplsinsegmentmaptable.Mplsinsegmentmapentry)-1]
    }
    return nil
}

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsinsegmentmaptable.Mplsinsegmentmapentry {
        children[mplsinsegmentmaptable.Mplsinsegmentmapentry[i].GetSegmentPath()] = &mplsinsegmentmaptable.Mplsinsegmentmapentry[i]
    }
    return children
}

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetYangName() string { return "mplsInSegmentMapTable" }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) SetParent(parent types.Entity) { mplsinsegmentmaptable.parent = parent }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetParent() types.Entity { return mplsinsegmentmaptable.parent }

func (mplsinsegmentmaptable *MPLSLSRSTDMIB_Mplsinsegmentmaptable) GetParentYangName() string { return "MPLS-LSR-STD-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Mplsinsegmentmaplabelptrindex interface{}

    // The mplsInSegmentIndex that corresponds to the mplsInSegmentInterface and
    // mplsInSegmentLabel, or the mplsInSegmentInterface and
    // mplsInSegmentLabelPtr, if applicable. The string containing the single
    // octet 0x00 MUST not be returned. The type is string with length: 1..24.
    Mplsinsegmentmapindex interface{}
}

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetFilter() yfilter.YFilter { return mplsinsegmentmapentry.YFilter }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) SetFilter(yf yfilter.YFilter) { mplsinsegmentmapentry.YFilter = yf }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetGoName(yname string) string {
    if yname == "mplsInSegmentMapInterface" { return "Mplsinsegmentmapinterface" }
    if yname == "mplsInSegmentMapLabel" { return "Mplsinsegmentmaplabel" }
    if yname == "mplsInSegmentMapLabelPtrIndex" { return "Mplsinsegmentmaplabelptrindex" }
    if yname == "mplsInSegmentMapIndex" { return "Mplsinsegmentmapindex" }
    return ""
}

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetSegmentPath() string {
    return "mplsInSegmentMapEntry" + "[mplsInSegmentMapInterface='" + fmt.Sprintf("%v", mplsinsegmentmapentry.Mplsinsegmentmapinterface) + "']" + "[mplsInSegmentMapLabel='" + fmt.Sprintf("%v", mplsinsegmentmapentry.Mplsinsegmentmaplabel) + "']" + "[mplsInSegmentMapLabelPtrIndex='" + fmt.Sprintf("%v", mplsinsegmentmapentry.Mplsinsegmentmaplabelptrindex) + "']"
}

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsInSegmentMapInterface"] = mplsinsegmentmapentry.Mplsinsegmentmapinterface
    leafs["mplsInSegmentMapLabel"] = mplsinsegmentmapentry.Mplsinsegmentmaplabel
    leafs["mplsInSegmentMapLabelPtrIndex"] = mplsinsegmentmapentry.Mplsinsegmentmaplabelptrindex
    leafs["mplsInSegmentMapIndex"] = mplsinsegmentmapentry.Mplsinsegmentmapindex
    return leafs
}

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetYangName() string { return "mplsInSegmentMapEntry" }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) SetParent(parent types.Entity) { mplsinsegmentmapentry.parent = parent }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetParent() types.Entity { return mplsinsegmentmapentry.parent }

func (mplsinsegmentmapentry *MPLSLSRSTDMIB_Mplsinsegmentmaptable_Mplsinsegmentmapentry) GetParentYangName() string { return "mplsInSegmentMapTable" }

